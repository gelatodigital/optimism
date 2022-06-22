package derive

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/ethereum-optimism/optimism/op-node/eth"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/log"
)

type PayloadQueueOutput interface {
	AddSafeAttributes(attributes *eth.PayloadAttributes)
	SafeL2Head() eth.L2BlockRef
}

type PayloadQueue struct {
	log       log.Logger
	config    *rollup.Config
	dl        L1ReceiptsFetcher
	next      PayloadQueueOutput
	resetting bool

	l1Origins []eth.L1BlockRef
	batches   []*BatchData
}

func NewPayloadQueue(log log.Logger, cfg *rollup.Config, l1Fetcher L1ReceiptsFetcher, next PayloadQueueOutput) *PayloadQueue {
	return &PayloadQueue{
		log:    log,
		config: cfg,
		dl:     l1Fetcher,
		next:   next,
	}
}

func (pq *PayloadQueue) AddBatch(batch *BatchData) {
	pq.log.Warn("Received next batch", "batch_number", batch.BlockNumber, "batch_epoch", batch.Epoch, "batch_timestamp", batch.Timestamp, "txs", len(batch.Transactions))
	pq.batches = append(pq.batches, batch)
}

func (pq *PayloadQueue) OpenOrigin(origin eth.L1BlockRef) {
	pq.l1Origins = append(pq.l1Origins, origin)
	pq.log.Warn("Got next L1 origin", "origin", origin, "first_origin", pq.l1Origins[0])
}

func (pq *PayloadQueue) CloseOrigin() {
	// IDK WTF this does
}

func (pq *PayloadQueue) Step(ctx context.Context) error {
	attr, err := pq.DeriveL2Inputs(ctx, pq.next.SafeL2Head())
	if err != nil {
		return err
	}
	// TODO: Reset step here
	pq.next.AddSafeAttributes(attr)
	return nil
}

func (pq *PayloadQueue) ResetStep(ctx context.Context, l1Fetcher L1Fetcher) error {
	// Reset such that the highestL1InclusionBlock is the same as the l2SafeHeadOrigin - the sequence window size
	pq.batches = pq.batches[:0]
	pq.l1Origins = pq.l1Origins[:0]

	startNumber := pq.next.SafeL2Head().L1Origin.Number
	if startNumber < pq.config.SeqWindowSize {
		startNumber = 0
	} else {
		startNumber -= pq.config.SeqWindowSize
	}
	l1BlockStart, err := l1Fetcher.L1BlockRefByNumber(ctx, startNumber)
	if err != nil {
		return err
	}

	pq.log.Info("found reset origin for payload queue", "origin", l1BlockStart)
	pq.l1Origins = append(pq.l1Origins, l1BlockStart)
	// pq.originOpen = true
	return io.EOF
}

func (pq *PayloadQueue) SafeL2Head() eth.L2BlockRef {
	return pq.next.SafeL2Head()
}

func (pq *PayloadQueue) DeriveL2Inputs(ctx context.Context, l2SafeHead eth.L2BlockRef) (*eth.PayloadAttributes, error) {
	if len(pq.l1Origins) == 0 {
		return nil, io.EOF
	}
	if len(pq.batches) == 0 {
		return nil, io.EOF
	}
	batch := pq.batches[0]
	// TODO: Don't slice until later?
	pq.batches = pq.batches[1:]

	seqNumber := l2SafeHead.SequenceNumber + 1
	l1Origin := pq.l1Origins[0]

	// Check if we need to advance an epoch & update local state
	if l1Origin.Number != uint64(batch.Epoch) {
		pq.log.Info("advancing epoch in the payload queue")
		seqNumber = 0
		// Not enough saved origins to advance an epoch here.
		if len(pq.l1Origins) < 2 {
			return nil, io.EOF
		}
		l1Origin = pq.l1Origins[1]
		pq.l1Origins = pq.l1Origins[1:]
	}

	fetchCtx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()
	l1Info, _, receipts, err := pq.dl.Fetch(fetchCtx, l1Origin.Hash)
	if err != nil {
		pq.log.Error("failed to fetch L1 block info", "l1Origin", l1Origin, "err", err)
		return nil, err
	}

	// Fill in deposits if we are the first block of the epoch
	var deposits []hexutil.Bytes
	if seqNumber == 0 {
		fetchCtx, cancel = context.WithTimeout(ctx, 20*time.Second)
		defer cancel()

		var errs []error
		deposits, errs = DeriveDeposits(receipts, pq.config.DepositContractAddress)
		for _, err := range errs {
			pq.log.Error("Failed to derive a deposit", "l1OriginHash", l1Origin.Hash, "err", err)
		}
		if len(errs) != 0 {
			// TODO: Multierror here
			return nil, fmt.Errorf("failed to derive some deposits: %v", errs)
		}
	}

	var txns []eth.Data
	l1InfoTx, err := L1InfoDepositBytes(seqNumber, l1Info)
	if err != nil {
		return nil, fmt.Errorf("failed to create l1InfoTx: %w", err)
	}
	txns = append(txns, l1InfoTx)
	if seqNumber == 0 {
		txns = append(txns, deposits...)
	}
	txns = append(txns, batch.Transactions...)
	attrs := &eth.PayloadAttributes{
		Timestamp:             hexutil.Uint64(batch.Timestamp),
		PrevRandao:            eth.Bytes32(l1Info.MixDigest()),
		SuggestedFeeRecipient: pq.config.FeeRecipientAddress,
		Transactions:          txns,
		// we are verifying, not sequencing, we've got all transactions and do not pull from the tx-pool
		// (that would make the block derivation non-deterministic)
		NoTxPool: true,
	}
	pq.log.Warn("generated attributes in payload queue", "tx_count", len(txns), "timestamp", batch.Timestamp)

	return attrs, nil
}
