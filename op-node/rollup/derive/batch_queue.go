package derive

import (
	"context"
	"fmt"
	"io"
	"sort"

	"github.com/ethereum-optimism/optimism/op-node/eth"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

type L1ReceiptsFetcher interface {
	Fetch(ctx context.Context, blockHash common.Hash) (eth.L1Info, types.Transactions, types.Receipts, error)
}

type BatchQueueOutput interface {
	OpenOrigin(eth.L1BlockRef)
	CloseOrigin()
	AddBatch(batch *BatchData)
	SafeL2Head() eth.L2BlockRef
}

type BatchWithL1InclusionBlock struct {
	L1InclusionBlock eth.L1BlockRef
	Batch            *BatchData
}

// BatchQueue contains a set of batches for every L1 block.
// L1 blocks are contiguous and this does not support reorgs.
type BatchQueue struct {
	log                     log.Logger
	originOpen              bool // true if the last origin expects more batches
	highestL1InclusionBlock eth.L1BlockRef
	config                  *rollup.Config
	dl                      L1ReceiptsFetcher
	next                    BatchQueueOutput

	l1Blocks []eth.L1BlockRef

	// epoch number that the batch queue is currently processing
	// IMPT: This must be updated in lockstep with the window
	epoch uint64

	// All batches with the same L2 block number. Batches are ordered by when they are seen.
	// Do a linear scan over the batches rather than deeply nested maps.
	// Note: Only a single batch with the same tuple (block number, timestamp, epoch) is allowed.
	batchesByNumber map[uint64][]*BatchWithL1InclusionBlock
}

// NewBatchQueue creates a BatchQueue, which should be Reset(origin) before use.
func NewBatchQueue(log log.Logger, cfg *rollup.Config, dl L1ReceiptsFetcher, next BatchQueueOutput) *BatchQueue {
	return &BatchQueue{
		log:             log,
		config:          cfg,
		dl:              dl,
		next:            next,
		batchesByNumber: make(map[uint64][]*BatchWithL1InclusionBlock),
	}
}

func (bq *BatchQueue) CurrentOrigin() eth.L1BlockRef {
	last := bq.highestL1InclusionBlock
	if len(bq.l1Blocks) != 0 {
		last = bq.l1Blocks[len(bq.l1Blocks)-1]
	}
	return last
}

func (bq *BatchQueue) OpenOrigin(origin eth.L1BlockRef) error {
	if bq.originOpen {
		panic("double open")
	}
	bq.log.Warn("open origin", "origin", origin)
	// TODO: WTF, why can we assume we have current origin when opening an origin?
	parent := bq.CurrentOrigin()
	if parent.Hash != origin.ParentHash {
		return fmt.Errorf("cannot process L1 reorg from %s to %s (parent %s)", parent, origin.ID(), origin.ParentID())
	}
	bq.l1Blocks = append(bq.l1Blocks, origin)
	bq.originOpen = true
	bq.next.OpenOrigin(origin)
	return nil
}

func (bq *BatchQueue) CloseOrigin() {
	if !bq.originOpen {
		panic("double close")
	}
	bq.log.Warn("close origin", "origin", bq.CurrentOrigin())
	bq.originOpen = false
	bq.next.CloseOrigin()
}

func (bq *BatchQueue) IsOriginOpen() bool {
	return bq.originOpen
}

func (bq *BatchQueue) Step(ctx context.Context) error {
	batches := bq.deriveBatches(bq.next.SafeL2Head())
	if len(batches) == 0 {
		return io.EOF
	}

	for _, batch := range batches {
		if uint64(batch.Batch.Timestamp) <= bq.next.SafeL2Head().Time {
			// drop attributes if we are still progressing towards the next stage
			// (after a reset rolled us back a full sequence window)
			continue
		}
		bq.next.AddBatch(batch.Batch)
	}
	return nil
}

func (bq *BatchQueue) ResetStep(ctx context.Context, l1Fetcher L1Fetcher) error {
	// Reset such that the highestL1InclusionBlock is the same as the l2SafeHeadOrigin - the sequence window size
	bq.batchesByNumber = make(map[uint64][]*BatchWithL1InclusionBlock)
	bq.l1Blocks = bq.l1Blocks[:0]

	startNumber := bq.next.SafeL2Head().L1Origin.Number
	if startNumber < bq.config.SeqWindowSize {
		startNumber = 0
	} else {
		startNumber -= bq.config.SeqWindowSize
	}
	l1BlockStart, err := l1Fetcher.L1BlockRefByNumber(ctx, startNumber)
	if err != nil {
		return err
	}

	bq.highestL1InclusionBlock = l1BlockStart
	bq.log.Info("found reset origin for batch queue", "origin", bq.highestL1InclusionBlock)
	bq.l1Blocks = append(bq.l1Blocks, bq.highestL1InclusionBlock)
	bq.originOpen = true
	bq.epoch = bq.highestL1InclusionBlock.Number
	return io.EOF
}

func (bq *BatchQueue) AddBatch(batch *BatchData) error {
	if !bq.originOpen {
		panic("write batch while closed")
	}
	bq.log.Warn("add batch", "origin", bq.CurrentOrigin(), "tx_count", len(batch.Transactions))
	if len(bq.l1Blocks) == 0 {
		return fmt.Errorf("cannot add batch with timestamp %d, no origin was prepared", batch.Timestamp)
	}
	if batch.BlockNumber <= bq.next.SafeL2Head().Number {
		bq.log.Warn("Old batch is getting included")
		return nil
	}
	data := BatchWithL1InclusionBlock{
		L1InclusionBlock: bq.CurrentOrigin(),
		Batch:            batch,
	}
	batches, ok := bq.batchesByNumber[batch.BlockNumber]
	// Filter complete duplicates. This step is not strictly needed as we always append, but it is nice to avoid lots
	// of spam. If we enforced a tighter relationship between block number & timestamp we could limit to a single
	// blocknumber & epoch pair.
	// TODO: Just log instead of error?
	// TODO: am getting duplicate batches here, not sure why
	if ok {
		for _, b := range batches {
			if b.Batch.BlockNumber == batch.BlockNumber &&
				b.Batch.Timestamp == batch.Timestamp &&
				b.Batch.Epoch == batch.Epoch {
				// return errors.New("duplicate batch")
				bq.log.Warn("duplicate batch", "batch_number", batch.BlockNumber, "batch_epoch", batch.Epoch, "batch_timestamp", batch.Timestamp, "txs", len(batch.Transactions))
				return nil
			}
		}
	} else {
		bq.log.Warn("First seen batch", "batch_number", batch.BlockNumber, "batch_epoch", batch.Epoch, "batch_timestamp", batch.Timestamp, "txs", len(batch.Transactions))

	}
	// May have duplicate block numbers or individual fields, but have limited complete duplicates
	bq.batchesByNumber[batch.BlockNumber] = append(batches, &data)
	return nil
}

// validExtension determines if a batch follows the previous attributes
func validExtension(cfg *rollup.Config, batch *BatchWithL1InclusionBlock, prevNumber, prevTime, prevEpoch uint64) bool {
	if batch.Batch.BlockNumber != prevNumber+1 {
		return false
	}
	if batch.Batch.Timestamp != prevTime+cfg.BlockTime {
		return false
	}
	if batch.Batch.Epoch != rollup.Epoch(prevEpoch) && batch.Batch.Epoch != rollup.Epoch(prevEpoch+1) {
		return false
	}
	// TODO (VERY IMPORTANT): Filter this batch out if the origin of the batch is too far past the epoch
	// TODO (ALSO VERY IMPT): Get this equality check correct
	// TODO: Also do this check when ingesting batches.
	if uint64(batch.Batch.Epoch) >= batch.L1InclusionBlock.Number+cfg.SeqWindowSize {
		return false
	}
	return true
}

// deriveBatches pulls a single batch eagerly or a collection of batches if it is the end of
// the sequencing window.
func (bq *BatchQueue) deriveBatches(l2SafeHead eth.L2BlockRef) []*BatchWithL1InclusionBlock {

	// Decide if need to fill out empty batches & process an epoch at once
	// If not, just return a single batch
	if bq.CurrentOrigin().Number >= bq.epoch+bq.config.SeqWindowSize {
		// 2a. Gather all batches. First sort by number and then by first seen.
		var bns []uint64
		for n := range bq.batchesByNumber {
			bns = append(bns, n)
		}
		sort.Slice(bns, func(i, j int) bool { return bns[i] < bns[j] })

		var batches []*BatchWithL1InclusionBlock
		for _, n := range bns {
			for _, batch := range bq.batchesByNumber[n] {
				// Filter out batches that were submitted too late.
				// TODO: Another place to check the equality symbol.
				if batch.L1InclusionBlock.Number-bq.CurrentOrigin().Number >= bq.config.SeqWindowSize {
					continue
				}
				// Pre filter batches in the correct epoch
				if batch.Batch.Epoch == rollup.Epoch(bq.epoch) {
					batches = append(batches, batch)
				}
			}
		}

		// 2b. Determine the valid time window
		l1OriginTime := bq.l1Blocks[0].Time
		nextL1BlockTime := bq.l1Blocks[1].Time // Safe b/c the epoch is the L1 Block nubmer of the first block in L1Blocks
		minL2Time := l2SafeHead.Time + bq.config.BlockTime
		maxL2Time := l1OriginTime + bq.config.MaxSequencerDrift
		if minL2Time+bq.config.BlockTime > maxL2Time {
			maxL2Time = minL2Time + bq.config.BlockTime
		}
		// Filter + Fill batches
		batches = FilterBatchesV2(bq.config, rollup.Epoch(bq.epoch), minL2Time, maxL2Time, batches)
		bq.log.Warn("filtered batches", "len", len(batches), "l1Origin", bq.l1Blocks[0], "nextL1Block", bq.l1Blocks[1])
		batches = FillMissingBatchesV2(batches, bq.epoch, bq.config.BlockTime, minL2Time, nextL1BlockTime, l2SafeHead.Number)

		// Advance an epoch after filling all batches.
		bq.epoch += 1
		bq.l1Blocks = bq.l1Blocks[1:]

		return batches

	} else {
		var ret []*BatchWithL1InclusionBlock
		next := bq.tryPopNextBatch(l2SafeHead)
		if next != nil {
			ret = append(ret, next)
		}
		return ret
	}

}

// tryPopNextBatch tries to get the next batch from the batch queue using an eager approach.
func (bq *BatchQueue) tryPopNextBatch(l2SafeHead eth.L2BlockRef) *BatchWithL1InclusionBlock {
	// We require at least 1 L1 blocks to look at.
	if len(bq.l1Blocks) == 0 {
		return nil
	}
	batches, ok := bq.batchesByNumber[l2SafeHead.Number+1]
	// No more batches found. Rely on another function to fill missing batches for us.
	if !ok {
		return nil
	}

	// Find the first batch saved for this number.
	// Note that we expect the number of batches for the same block number to be small (frequently just 1 ).
	for _, batch := range batches {
		l1OriginTime := bq.l1Blocks[0].Time

		// If this batch advances the epoch, check it's validity against the next L1 Origin
		if batch.Batch.Epoch != rollup.Epoch(l2SafeHead.L1Origin.Number) {
			// With only 1 l1Block we cannot look at the next L1 Origin.
			// Note: This means that we are unable to determine validity of a batch
			// without more information. In this case we should bail out until we have
			// more information otherwise the eager algorithm may diverge from a non-eager
			// algorithm.
			if len(bq.l1Blocks) < 2 {
				return nil
			}
			l1OriginTime = bq.l1Blocks[1].Time
		}

		// Timestamp bounds
		minL2Time := l2SafeHead.Time + bq.config.BlockTime
		maxL2Time := l1OriginTime + bq.config.MaxSequencerDrift
		if minL2Time+bq.config.BlockTime > maxL2Time {
			maxL2Time = minL2Time + bq.config.BlockTime
		}

		// Note: Don't check epoch here, check it in `validExtension`
		// Mainly check tx validity + timestamp bounds
		if err := ValidBatch(batch.Batch, bq.config, batch.Batch.Epoch, minL2Time, maxL2Time); err != nil {
			break
		}

		// We have a valid batch, no make sure that it builds off the previous L2 block
		if validExtension(bq.config, batch, l2SafeHead.Number, l2SafeHead.Number, l2SafeHead.L1Origin.Number) {
			// Advance the epoch if needed
			if l2SafeHead.L1Origin.Number != uint64(batch.Batch.Epoch) {
				bq.l1Blocks = bq.l1Blocks[1:]
				bq.epoch = uint64(batch.Batch.Epoch)
			}
			// Don't leak data in the map
			delete(bq.batchesByNumber, batch.Batch.BlockNumber)

			// We have found the fist valid batch.
			return batch
		}
	}

	return nil
}
