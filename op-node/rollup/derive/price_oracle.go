package derive

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum-optimism/optimism/op-node/bindings"
	"github.com/ethereum-optimism/optimism/op-service/predeploys"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var (
	L2PriceOracleABI, _  = abi.JSON(strings.NewReader(bindings.L2PriceOracleABI))
	L2PriceOracleAddress = predeploys.L2PriceOracleAddr
)

func PriceOracleDeposit(event *bindings.L1PriceOraclePricesUpdated) (*types.DepositTx, error) {
	data, err := L2PriceOracleABI.Pack("updatePrices", event.Ids, event.Prices)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal price oracle update: %w", err)
	}
	// We only update prices in the first block of every epoch
	// Hence, sequencer number is always zero
	source := L1InfoDepositSource{
		L1BlockHash: event.Raw.BlockHash,
		SeqNumber:   0,
	}
	out := &types.DepositTx{
		SourceHash:          source.SourceHash(),
		From:                L1InfoDepositerAddress,
		To:                  &L2PriceOracleAddress,
		Mint:                nil,
		Value:               big.NewInt(0),
		Gas:                 RegolithSystemTxGas,
		IsSystemTransaction: false,
		Data:                data,
	}
	return out, nil
}

func PriceOracleDepositBytes(receipts []*types.Receipt, l1PriceOracleAddr common.Address) ([]byte, error) {
	zeroAddr := common.Address{}
	if l1PriceOracleAddr == zeroAddr {
		return nil, nil
	}
	log := FetchPricesUpdatedLog(receipts, l1PriceOracleAddr)
	if log == nil {
		return nil, fmt.Errorf("failed to find price oracle update log")
	}
	event, err := ParsePricesUpdatedLog(log)
	if err != nil {
		return nil, fmt.Errorf("failed to parse price oracle update log: %w", err)
	}
	dep, err := PriceOracleDeposit(event)
	if err != nil {
		return nil, fmt.Errorf("failed to create price oracle tx: %w", err)
	}
	tx := types.NewTx(dep)
	opaqueTx, err := tx.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("failed to encode price oracle tx: %w", err)
	}
	return opaqueTx, nil
}
