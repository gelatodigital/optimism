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

// PythStructsPrice is an auto generated low-level Go binding around an user-defined struct.
type PythStructsPrice struct {
	Price       int64
	Conf        uint64
	Expo        int32
	PublishTime *big.Int
}

// L1PriceOracleMetaData contains all meta data concerning the L1PriceOracle contract.
var L1PriceOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"ids\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"price\",\"type\":\"int64\"},{\"internalType\":\"uint64\",\"name\":\"conf\",\"type\":\"uint64\"},{\"internalType\":\"int32\",\"name\":\"expo\",\"type\":\"int32\"},{\"internalType\":\"uint256\",\"name\":\"publishTime\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structPythStructs.Price[]\",\"name\":\"prices\",\"type\":\"tuple[]\"}],\"name\":\"PricesUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"dedicatedMsgSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pyth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dedicatedMsgSender\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pyth\",\"outputs\":[{\"internalType\":\"contractIPyth\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_ids\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_updateData\",\"type\":\"bytes[]\"}],\"name\":\"updatePrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// L1PriceOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use L1PriceOracleMetaData.ABI instead.
var L1PriceOracleABI = L1PriceOracleMetaData.ABI

// L1PriceOracle is an auto generated Go binding around an Ethereum contract.
type L1PriceOracle struct {
	L1PriceOracleCaller     // Read-only binding to the contract
	L1PriceOracleTransactor // Write-only binding to the contract
	L1PriceOracleFilterer   // Log filterer for contract events
}

// L1PriceOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1PriceOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1PriceOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1PriceOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1PriceOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L1PriceOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1PriceOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L1PriceOracleSession struct {
	Contract     *L1PriceOracle    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L1PriceOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L1PriceOracleCallerSession struct {
	Contract *L1PriceOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// L1PriceOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L1PriceOracleTransactorSession struct {
	Contract     *L1PriceOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// L1PriceOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type L1PriceOracleRaw struct {
	Contract *L1PriceOracle // Generic contract binding to access the raw methods on
}

// L1PriceOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L1PriceOracleCallerRaw struct {
	Contract *L1PriceOracleCaller // Generic read-only contract binding to access the raw methods on
}

// L1PriceOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L1PriceOracleTransactorRaw struct {
	Contract *L1PriceOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL1PriceOracle creates a new instance of L1PriceOracle, bound to a specific deployed contract.
func NewL1PriceOracle(address common.Address, backend bind.ContractBackend) (*L1PriceOracle, error) {
	contract, err := bindL1PriceOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L1PriceOracle{L1PriceOracleCaller: L1PriceOracleCaller{contract: contract}, L1PriceOracleTransactor: L1PriceOracleTransactor{contract: contract}, L1PriceOracleFilterer: L1PriceOracleFilterer{contract: contract}}, nil
}

// NewL1PriceOracleCaller creates a new read-only instance of L1PriceOracle, bound to a specific deployed contract.
func NewL1PriceOracleCaller(address common.Address, caller bind.ContractCaller) (*L1PriceOracleCaller, error) {
	contract, err := bindL1PriceOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1PriceOracleCaller{contract: contract}, nil
}

// NewL1PriceOracleTransactor creates a new write-only instance of L1PriceOracle, bound to a specific deployed contract.
func NewL1PriceOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*L1PriceOracleTransactor, error) {
	contract, err := bindL1PriceOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1PriceOracleTransactor{contract: contract}, nil
}

// NewL1PriceOracleFilterer creates a new log filterer instance of L1PriceOracle, bound to a specific deployed contract.
func NewL1PriceOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*L1PriceOracleFilterer, error) {
	contract, err := bindL1PriceOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L1PriceOracleFilterer{contract: contract}, nil
}

// bindL1PriceOracle binds a generic wrapper to an already deployed contract.
func bindL1PriceOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L1PriceOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1PriceOracle *L1PriceOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1PriceOracle.Contract.L1PriceOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1PriceOracle *L1PriceOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1PriceOracle.Contract.L1PriceOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1PriceOracle *L1PriceOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1PriceOracle.Contract.L1PriceOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1PriceOracle *L1PriceOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1PriceOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1PriceOracle *L1PriceOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1PriceOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1PriceOracle *L1PriceOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1PriceOracle.Contract.contract.Transact(opts, method, params...)
}

// DedicatedMsgSender is a free data retrieval call binding the contract method 0x28f150eb.
//
// Solidity: function dedicatedMsgSender() view returns(address)
func (_L1PriceOracle *L1PriceOracleCaller) DedicatedMsgSender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1PriceOracle.contract.Call(opts, &out, "dedicatedMsgSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DedicatedMsgSender is a free data retrieval call binding the contract method 0x28f150eb.
//
// Solidity: function dedicatedMsgSender() view returns(address)
func (_L1PriceOracle *L1PriceOracleSession) DedicatedMsgSender() (common.Address, error) {
	return _L1PriceOracle.Contract.DedicatedMsgSender(&_L1PriceOracle.CallOpts)
}

// DedicatedMsgSender is a free data retrieval call binding the contract method 0x28f150eb.
//
// Solidity: function dedicatedMsgSender() view returns(address)
func (_L1PriceOracle *L1PriceOracleCallerSession) DedicatedMsgSender() (common.Address, error) {
	return _L1PriceOracle.Contract.DedicatedMsgSender(&_L1PriceOracle.CallOpts)
}

// Pyth is a free data retrieval call binding the contract method 0xf98d06f0.
//
// Solidity: function pyth() view returns(address)
func (_L1PriceOracle *L1PriceOracleCaller) Pyth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1PriceOracle.contract.Call(opts, &out, "pyth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pyth is a free data retrieval call binding the contract method 0xf98d06f0.
//
// Solidity: function pyth() view returns(address)
func (_L1PriceOracle *L1PriceOracleSession) Pyth() (common.Address, error) {
	return _L1PriceOracle.Contract.Pyth(&_L1PriceOracle.CallOpts)
}

// Pyth is a free data retrieval call binding the contract method 0xf98d06f0.
//
// Solidity: function pyth() view returns(address)
func (_L1PriceOracle *L1PriceOracleCallerSession) Pyth() (common.Address, error) {
	return _L1PriceOracle.Contract.Pyth(&_L1PriceOracle.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_L1PriceOracle *L1PriceOracleTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1PriceOracle.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_L1PriceOracle *L1PriceOracleSession) Deposit() (*types.Transaction, error) {
	return _L1PriceOracle.Contract.Deposit(&_L1PriceOracle.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_L1PriceOracle *L1PriceOracleTransactorSession) Deposit() (*types.Transaction, error) {
	return _L1PriceOracle.Contract.Deposit(&_L1PriceOracle.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _pyth, address _dedicatedMsgSender) returns()
func (_L1PriceOracle *L1PriceOracleTransactor) Initialize(opts *bind.TransactOpts, _pyth common.Address, _dedicatedMsgSender common.Address) (*types.Transaction, error) {
	return _L1PriceOracle.contract.Transact(opts, "initialize", _pyth, _dedicatedMsgSender)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _pyth, address _dedicatedMsgSender) returns()
func (_L1PriceOracle *L1PriceOracleSession) Initialize(_pyth common.Address, _dedicatedMsgSender common.Address) (*types.Transaction, error) {
	return _L1PriceOracle.Contract.Initialize(&_L1PriceOracle.TransactOpts, _pyth, _dedicatedMsgSender)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _pyth, address _dedicatedMsgSender) returns()
func (_L1PriceOracle *L1PriceOracleTransactorSession) Initialize(_pyth common.Address, _dedicatedMsgSender common.Address) (*types.Transaction, error) {
	return _L1PriceOracle.Contract.Initialize(&_L1PriceOracle.TransactOpts, _pyth, _dedicatedMsgSender)
}

// UpdatePrices is a paid mutator transaction binding the contract method 0x0aa9adbc.
//
// Solidity: function updatePrices(bytes32[] _ids, bytes[] _updateData) returns()
func (_L1PriceOracle *L1PriceOracleTransactor) UpdatePrices(opts *bind.TransactOpts, _ids [][32]byte, _updateData [][]byte) (*types.Transaction, error) {
	return _L1PriceOracle.contract.Transact(opts, "updatePrices", _ids, _updateData)
}

// UpdatePrices is a paid mutator transaction binding the contract method 0x0aa9adbc.
//
// Solidity: function updatePrices(bytes32[] _ids, bytes[] _updateData) returns()
func (_L1PriceOracle *L1PriceOracleSession) UpdatePrices(_ids [][32]byte, _updateData [][]byte) (*types.Transaction, error) {
	return _L1PriceOracle.Contract.UpdatePrices(&_L1PriceOracle.TransactOpts, _ids, _updateData)
}

// UpdatePrices is a paid mutator transaction binding the contract method 0x0aa9adbc.
//
// Solidity: function updatePrices(bytes32[] _ids, bytes[] _updateData) returns()
func (_L1PriceOracle *L1PriceOracleTransactorSession) UpdatePrices(_ids [][32]byte, _updateData [][]byte) (*types.Transaction, error) {
	return _L1PriceOracle.Contract.UpdatePrices(&_L1PriceOracle.TransactOpts, _ids, _updateData)
}

// L1PriceOracleInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L1PriceOracle contract.
type L1PriceOracleInitializedIterator struct {
	Event *L1PriceOracleInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *L1PriceOracleInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PriceOracleInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(L1PriceOracleInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *L1PriceOracleInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PriceOracleInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PriceOracleInitialized represents a Initialized event raised by the L1PriceOracle contract.
type L1PriceOracleInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_L1PriceOracle *L1PriceOracleFilterer) FilterInitialized(opts *bind.FilterOpts) (*L1PriceOracleInitializedIterator, error) {

	logs, sub, err := _L1PriceOracle.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L1PriceOracleInitializedIterator{contract: _L1PriceOracle.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_L1PriceOracle *L1PriceOracleFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L1PriceOracleInitialized) (event.Subscription, error) {

	logs, sub, err := _L1PriceOracle.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PriceOracleInitialized)
				if err := _L1PriceOracle.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_L1PriceOracle *L1PriceOracleFilterer) ParseInitialized(log types.Log) (*L1PriceOracleInitialized, error) {
	event := new(L1PriceOracleInitialized)
	if err := _L1PriceOracle.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PriceOraclePricesUpdatedIterator is returned from FilterPricesUpdated and is used to iterate over the raw logs and unpacked data for PricesUpdated events raised by the L1PriceOracle contract.
type L1PriceOraclePricesUpdatedIterator struct {
	Event *L1PriceOraclePricesUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *L1PriceOraclePricesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PriceOraclePricesUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(L1PriceOraclePricesUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *L1PriceOraclePricesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PriceOraclePricesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PriceOraclePricesUpdated represents a PricesUpdated event raised by the L1PriceOracle contract.
type L1PriceOraclePricesUpdated struct {
	Ids    [][32]byte
	Prices []PythStructsPrice
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPricesUpdated is a free log retrieval operation binding the contract event 0xcf1044103f148b58a40745f90b72a8e9d119c3207bb5aba5918798a7ecc2bac1.
//
// Solidity: event PricesUpdated(bytes32[] ids, (int64,uint64,int32,uint256)[] prices)
func (_L1PriceOracle *L1PriceOracleFilterer) FilterPricesUpdated(opts *bind.FilterOpts) (*L1PriceOraclePricesUpdatedIterator, error) {

	logs, sub, err := _L1PriceOracle.contract.FilterLogs(opts, "PricesUpdated")
	if err != nil {
		return nil, err
	}
	return &L1PriceOraclePricesUpdatedIterator{contract: _L1PriceOracle.contract, event: "PricesUpdated", logs: logs, sub: sub}, nil
}

// WatchPricesUpdated is a free log subscription operation binding the contract event 0xcf1044103f148b58a40745f90b72a8e9d119c3207bb5aba5918798a7ecc2bac1.
//
// Solidity: event PricesUpdated(bytes32[] ids, (int64,uint64,int32,uint256)[] prices)
func (_L1PriceOracle *L1PriceOracleFilterer) WatchPricesUpdated(opts *bind.WatchOpts, sink chan<- *L1PriceOraclePricesUpdated) (event.Subscription, error) {

	logs, sub, err := _L1PriceOracle.contract.WatchLogs(opts, "PricesUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PriceOraclePricesUpdated)
				if err := _L1PriceOracle.contract.UnpackLog(event, "PricesUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePricesUpdated is a log parse operation binding the contract event 0xcf1044103f148b58a40745f90b72a8e9d119c3207bb5aba5918798a7ecc2bac1.
//
// Solidity: event PricesUpdated(bytes32[] ids, (int64,uint64,int32,uint256)[] prices)
func (_L1PriceOracle *L1PriceOracleFilterer) ParsePricesUpdated(log types.Log) (*L1PriceOraclePricesUpdated, error) {
	event := new(L1PriceOraclePricesUpdated)
	if err := _L1PriceOracle.contract.UnpackLog(event, "PricesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
