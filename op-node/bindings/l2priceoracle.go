// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// THIS FILE IS EDITED TO REMOVE DUPLICATE "PythStructsPrice" DECLARATION.

// L2PriceOracleMetaData contains all meta data concerning the L2PriceOracle contract.
var L2PriceOracleMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"DEPOSITOR_ACCOUNT\",\"inputs\":[],\"outputs\":[{\"name\":\"addr_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getPriceNoOlderThan\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"age\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"price_\",\"type\":\"tuple\",\"internalType\":\"structPythStructs.Price\",\"components\":[{\"name\":\"price\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"conf\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"expo\",\"type\":\"int32\",\"internalType\":\"int32\"},{\"name\":\"publishTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPriceUnsafe\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"price_\",\"type\":\"tuple\",\"internalType\":\"structPythStructs.Price\",\"components\":[{\"name\":\"price\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"conf\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"expo\",\"type\":\"int32\",\"internalType\":\"int32\"},{\"name\":\"publishTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updatePrices\",\"inputs\":[{\"name\":\"_ids\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"_prices\",\"type\":\"tuple[]\",\"internalType\":\"structPythStructs.Price[]\",\"components\":[{\"name\":\"price\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"conf\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"expo\",\"type\":\"int32\",\"internalType\":\"int32\"},{\"name\":\"publishTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"PriceFeedNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StalePrice\",\"inputs\":[]}]",
}

// L2PriceOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use L2PriceOracleMetaData.ABI instead.
var L2PriceOracleABI = L2PriceOracleMetaData.ABI

// L2PriceOracle is an auto generated Go binding around an Ethereum contract.
type L2PriceOracle struct {
	L2PriceOracleCaller     // Read-only binding to the contract
	L2PriceOracleTransactor // Write-only binding to the contract
	L2PriceOracleFilterer   // Log filterer for contract events
}

// L2PriceOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2PriceOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2PriceOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2PriceOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2PriceOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2PriceOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2PriceOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2PriceOracleSession struct {
	Contract     *L2PriceOracle    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2PriceOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2PriceOracleCallerSession struct {
	Contract *L2PriceOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// L2PriceOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2PriceOracleTransactorSession struct {
	Contract     *L2PriceOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// L2PriceOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2PriceOracleRaw struct {
	Contract *L2PriceOracle // Generic contract binding to access the raw methods on
}

// L2PriceOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2PriceOracleCallerRaw struct {
	Contract *L2PriceOracleCaller // Generic read-only contract binding to access the raw methods on
}

// L2PriceOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2PriceOracleTransactorRaw struct {
	Contract *L2PriceOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2PriceOracle creates a new instance of L2PriceOracle, bound to a specific deployed contract.
func NewL2PriceOracle(address common.Address, backend bind.ContractBackend) (*L2PriceOracle, error) {
	contract, err := bindL2PriceOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2PriceOracle{L2PriceOracleCaller: L2PriceOracleCaller{contract: contract}, L2PriceOracleTransactor: L2PriceOracleTransactor{contract: contract}, L2PriceOracleFilterer: L2PriceOracleFilterer{contract: contract}}, nil
}

// NewL2PriceOracleCaller creates a new read-only instance of L2PriceOracle, bound to a specific deployed contract.
func NewL2PriceOracleCaller(address common.Address, caller bind.ContractCaller) (*L2PriceOracleCaller, error) {
	contract, err := bindL2PriceOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2PriceOracleCaller{contract: contract}, nil
}

// NewL2PriceOracleTransactor creates a new write-only instance of L2PriceOracle, bound to a specific deployed contract.
func NewL2PriceOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*L2PriceOracleTransactor, error) {
	contract, err := bindL2PriceOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2PriceOracleTransactor{contract: contract}, nil
}

// NewL2PriceOracleFilterer creates a new log filterer instance of L2PriceOracle, bound to a specific deployed contract.
func NewL2PriceOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*L2PriceOracleFilterer, error) {
	contract, err := bindL2PriceOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2PriceOracleFilterer{contract: contract}, nil
}

// bindL2PriceOracle binds a generic wrapper to an already deployed contract.
func bindL2PriceOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L2PriceOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2PriceOracle *L2PriceOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2PriceOracle.Contract.L2PriceOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2PriceOracle *L2PriceOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2PriceOracle.Contract.L2PriceOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2PriceOracle *L2PriceOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2PriceOracle.Contract.L2PriceOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2PriceOracle *L2PriceOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2PriceOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2PriceOracle *L2PriceOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2PriceOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2PriceOracle *L2PriceOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2PriceOracle.Contract.contract.Transact(opts, method, params...)
}

// DEPOSITORACCOUNT is a free data retrieval call binding the contract method 0xe591b282.
//
// Solidity: function DEPOSITOR_ACCOUNT() pure returns(address addr_)
func (_L2PriceOracle *L2PriceOracleCaller) DEPOSITORACCOUNT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2PriceOracle.contract.Call(opts, &out, "DEPOSITOR_ACCOUNT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DEPOSITORACCOUNT is a free data retrieval call binding the contract method 0xe591b282.
//
// Solidity: function DEPOSITOR_ACCOUNT() pure returns(address addr_)
func (_L2PriceOracle *L2PriceOracleSession) DEPOSITORACCOUNT() (common.Address, error) {
	return _L2PriceOracle.Contract.DEPOSITORACCOUNT(&_L2PriceOracle.CallOpts)
}

// DEPOSITORACCOUNT is a free data retrieval call binding the contract method 0xe591b282.
//
// Solidity: function DEPOSITOR_ACCOUNT() pure returns(address addr_)
func (_L2PriceOracle *L2PriceOracleCallerSession) DEPOSITORACCOUNT() (common.Address, error) {
	return _L2PriceOracle.Contract.DEPOSITORACCOUNT(&_L2PriceOracle.CallOpts)
}

// GetPriceNoOlderThan is a free data retrieval call binding the contract method 0xa4ae35e0.
//
// Solidity: function getPriceNoOlderThan(bytes32 id, uint256 age) view returns((int64,uint64,int32,uint256) price_)
func (_L2PriceOracle *L2PriceOracleCaller) GetPriceNoOlderThan(opts *bind.CallOpts, id [32]byte, age *big.Int) (PythStructsPrice, error) {
	var out []interface{}
	err := _L2PriceOracle.contract.Call(opts, &out, "getPriceNoOlderThan", id, age)

	if err != nil {
		return *new(PythStructsPrice), err
	}

	out0 := *abi.ConvertType(out[0], new(PythStructsPrice)).(*PythStructsPrice)

	return out0, err

}

// GetPriceNoOlderThan is a free data retrieval call binding the contract method 0xa4ae35e0.
//
// Solidity: function getPriceNoOlderThan(bytes32 id, uint256 age) view returns((int64,uint64,int32,uint256) price_)
func (_L2PriceOracle *L2PriceOracleSession) GetPriceNoOlderThan(id [32]byte, age *big.Int) (PythStructsPrice, error) {
	return _L2PriceOracle.Contract.GetPriceNoOlderThan(&_L2PriceOracle.CallOpts, id, age)
}

// GetPriceNoOlderThan is a free data retrieval call binding the contract method 0xa4ae35e0.
//
// Solidity: function getPriceNoOlderThan(bytes32 id, uint256 age) view returns((int64,uint64,int32,uint256) price_)
func (_L2PriceOracle *L2PriceOracleCallerSession) GetPriceNoOlderThan(id [32]byte, age *big.Int) (PythStructsPrice, error) {
	return _L2PriceOracle.Contract.GetPriceNoOlderThan(&_L2PriceOracle.CallOpts, id, age)
}

// GetPriceUnsafe is a free data retrieval call binding the contract method 0x96834ad3.
//
// Solidity: function getPriceUnsafe(bytes32 id) view returns((int64,uint64,int32,uint256) price_)
func (_L2PriceOracle *L2PriceOracleCaller) GetPriceUnsafe(opts *bind.CallOpts, id [32]byte) (PythStructsPrice, error) {
	var out []interface{}
	err := _L2PriceOracle.contract.Call(opts, &out, "getPriceUnsafe", id)

	if err != nil {
		return *new(PythStructsPrice), err
	}

	out0 := *abi.ConvertType(out[0], new(PythStructsPrice)).(*PythStructsPrice)

	return out0, err

}

// GetPriceUnsafe is a free data retrieval call binding the contract method 0x96834ad3.
//
// Solidity: function getPriceUnsafe(bytes32 id) view returns((int64,uint64,int32,uint256) price_)
func (_L2PriceOracle *L2PriceOracleSession) GetPriceUnsafe(id [32]byte) (PythStructsPrice, error) {
	return _L2PriceOracle.Contract.GetPriceUnsafe(&_L2PriceOracle.CallOpts, id)
}

// GetPriceUnsafe is a free data retrieval call binding the contract method 0x96834ad3.
//
// Solidity: function getPriceUnsafe(bytes32 id) view returns((int64,uint64,int32,uint256) price_)
func (_L2PriceOracle *L2PriceOracleCallerSession) GetPriceUnsafe(id [32]byte) (PythStructsPrice, error) {
	return _L2PriceOracle.Contract.GetPriceUnsafe(&_L2PriceOracle.CallOpts, id)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2PriceOracle *L2PriceOracleCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _L2PriceOracle.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2PriceOracle *L2PriceOracleSession) Version() (string, error) {
	return _L2PriceOracle.Contract.Version(&_L2PriceOracle.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2PriceOracle *L2PriceOracleCallerSession) Version() (string, error) {
	return _L2PriceOracle.Contract.Version(&_L2PriceOracle.CallOpts)
}

// UpdatePrices is a paid mutator transaction binding the contract method 0x419781ae.
//
// Solidity: function updatePrices(bytes32[] _ids, (int64,uint64,int32,uint256)[] _prices) returns()
func (_L2PriceOracle *L2PriceOracleTransactor) UpdatePrices(opts *bind.TransactOpts, _ids [][32]byte, _prices []PythStructsPrice) (*types.Transaction, error) {
	return _L2PriceOracle.contract.Transact(opts, "updatePrices", _ids, _prices)
}

// UpdatePrices is a paid mutator transaction binding the contract method 0x419781ae.
//
// Solidity: function updatePrices(bytes32[] _ids, (int64,uint64,int32,uint256)[] _prices) returns()
func (_L2PriceOracle *L2PriceOracleSession) UpdatePrices(_ids [][32]byte, _prices []PythStructsPrice) (*types.Transaction, error) {
	return _L2PriceOracle.Contract.UpdatePrices(&_L2PriceOracle.TransactOpts, _ids, _prices)
}

// UpdatePrices is a paid mutator transaction binding the contract method 0x419781ae.
//
// Solidity: function updatePrices(bytes32[] _ids, (int64,uint64,int32,uint256)[] _prices) returns()
func (_L2PriceOracle *L2PriceOracleTransactorSession) UpdatePrices(_ids [][32]byte, _prices []PythStructsPrice) (*types.Transaction, error) {
	return _L2PriceOracle.Contract.UpdatePrices(&_L2PriceOracle.TransactOpts, _ids, _prices)
}
