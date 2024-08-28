package derive

import (
	"strings"

	"github.com/ethereum-optimism/optimism/op-node/bindings"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var (
	L1PriceOracleABI, _ = abi.JSON(strings.NewReader(bindings.L1PriceOracleABI))
)

func FetchPricesUpdatedLog(receipts []*types.Receipt, oracleContractAddr common.Address) *types.Log {
	// We iterate through receipts backwards since we want the most recent
	for i := len(receipts) - 1; i >= 0; i-- {
		rec := receipts[i]
		if rec.Status != types.ReceiptStatusSuccessful {
			continue
		}
		// We iterate through logs backwards since we want the most recent
		for j := len(rec.Logs) - 1; j >= 0; j-- {
			log := rec.Logs[j]
			if log.Address == oracleContractAddr && len(log.Topics) > 0 && log.Topics[0] == L1PriceOracleABI.Events["PricesUpdated"].ID {
				return log
			}
		}
	}
	return nil
}

func ParsePricesUpdatedLog(log *types.Log) (*bindings.L1PriceOraclePricesUpdated, error) {
	contract := bind.NewBoundContract(common.Address{}, L1PriceOracleABI, nil, nil, nil)
	event := new(bindings.L1PriceOraclePricesUpdated)
	if err := contract.UnpackLog(event, "PricesUpdated", *log); err != nil {
		return nil, err
	}
	event.Raw = *log
	return event, nil
}
