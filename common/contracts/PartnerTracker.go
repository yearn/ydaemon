// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// YPartnerTrackerMetaData contains all meta data concerning the YPartnerTracker contract.
var YPartnerTrackerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"partnerId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountAdded\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalDeposited\",\"type\":\"uint256\"}],\"name\":\"ReferredBalanceIncreased\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"partnerId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"partnerId\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"referredBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// YPartnerTrackerABI is the input ABI used to generate the binding from.
// Deprecated: Use YPartnerTrackerMetaData.ABI instead.
var YPartnerTrackerABI = YPartnerTrackerMetaData.ABI

// YPartnerTracker is an auto generated Go binding around an Ethereum contract.
type YPartnerTracker struct {
	YPartnerTrackerCaller     // Read-only binding to the contract
	YPartnerTrackerTransactor // Write-only binding to the contract
	YPartnerTrackerFilterer   // Log filterer for contract events
}

// YPartnerTrackerCaller is an auto generated read-only Go binding around an Ethereum contract.
type YPartnerTrackerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YPartnerTrackerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YPartnerTrackerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YPartnerTrackerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YPartnerTrackerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YPartnerTrackerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YPartnerTrackerSession struct {
	Contract     *YPartnerTracker  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YPartnerTrackerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YPartnerTrackerCallerSession struct {
	Contract *YPartnerTrackerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// YPartnerTrackerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YPartnerTrackerTransactorSession struct {
	Contract     *YPartnerTrackerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// YPartnerTrackerRaw is an auto generated low-level Go binding around an Ethereum contract.
type YPartnerTrackerRaw struct {
	Contract *YPartnerTracker // Generic contract binding to access the raw methods on
}

// YPartnerTrackerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YPartnerTrackerCallerRaw struct {
	Contract *YPartnerTrackerCaller // Generic read-only contract binding to access the raw methods on
}

// YPartnerTrackerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YPartnerTrackerTransactorRaw struct {
	Contract *YPartnerTrackerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYPartnerTracker creates a new instance of YPartnerTracker, bound to a specific deployed contract.
func NewYPartnerTracker(address common.Address, backend bind.ContractBackend) (*YPartnerTracker, error) {
	contract, err := bindYPartnerTracker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YPartnerTracker{YPartnerTrackerCaller: YPartnerTrackerCaller{contract: contract}, YPartnerTrackerTransactor: YPartnerTrackerTransactor{contract: contract}, YPartnerTrackerFilterer: YPartnerTrackerFilterer{contract: contract}}, nil
}

// NewYPartnerTrackerCaller creates a new read-only instance of YPartnerTracker, bound to a specific deployed contract.
func NewYPartnerTrackerCaller(address common.Address, caller bind.ContractCaller) (*YPartnerTrackerCaller, error) {
	contract, err := bindYPartnerTracker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YPartnerTrackerCaller{contract: contract}, nil
}

// NewYPartnerTrackerTransactor creates a new write-only instance of YPartnerTracker, bound to a specific deployed contract.
func NewYPartnerTrackerTransactor(address common.Address, transactor bind.ContractTransactor) (*YPartnerTrackerTransactor, error) {
	contract, err := bindYPartnerTracker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YPartnerTrackerTransactor{contract: contract}, nil
}

// NewYPartnerTrackerFilterer creates a new log filterer instance of YPartnerTracker, bound to a specific deployed contract.
func NewYPartnerTrackerFilterer(address common.Address, filterer bind.ContractFilterer) (*YPartnerTrackerFilterer, error) {
	contract, err := bindYPartnerTracker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YPartnerTrackerFilterer{contract: contract}, nil
}

// bindYPartnerTracker binds a generic wrapper to an already deployed contract.
func bindYPartnerTracker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(YPartnerTrackerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YPartnerTracker *YPartnerTrackerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YPartnerTracker.Contract.YPartnerTrackerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YPartnerTracker *YPartnerTrackerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YPartnerTracker.Contract.YPartnerTrackerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YPartnerTracker *YPartnerTrackerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YPartnerTracker.Contract.YPartnerTrackerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YPartnerTracker *YPartnerTrackerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YPartnerTracker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YPartnerTracker *YPartnerTrackerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YPartnerTracker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YPartnerTracker *YPartnerTrackerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YPartnerTracker.Contract.contract.Transact(opts, method, params...)
}

// ReferredBalance is a free data retrieval call binding the contract method 0x9e79e259.
//
// Solidity: function referredBalance(address , address , address ) view returns(uint256)
func (_YPartnerTracker *YPartnerTrackerCaller) ReferredBalance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YPartnerTracker.contract.Call(opts, &out, "referredBalance", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReferredBalance is a free data retrieval call binding the contract method 0x9e79e259.
//
// Solidity: function referredBalance(address , address , address ) view returns(uint256)
func (_YPartnerTracker *YPartnerTrackerSession) ReferredBalance(arg0 common.Address, arg1 common.Address, arg2 common.Address) (*big.Int, error) {
	return _YPartnerTracker.Contract.ReferredBalance(&_YPartnerTracker.CallOpts, arg0, arg1, arg2)
}

// ReferredBalance is a free data retrieval call binding the contract method 0x9e79e259.
//
// Solidity: function referredBalance(address , address , address ) view returns(uint256)
func (_YPartnerTracker *YPartnerTrackerCallerSession) ReferredBalance(arg0 common.Address, arg1 common.Address, arg2 common.Address) (*big.Int, error) {
	return _YPartnerTracker.Contract.ReferredBalance(&_YPartnerTracker.CallOpts, arg0, arg1, arg2)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(address vault, address partnerId, uint256 amount) returns(uint256)
func (_YPartnerTracker *YPartnerTrackerTransactor) Deposit(opts *bind.TransactOpts, vault common.Address, partnerId common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YPartnerTracker.contract.Transact(opts, "deposit", vault, partnerId, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(address vault, address partnerId, uint256 amount) returns(uint256)
func (_YPartnerTracker *YPartnerTrackerSession) Deposit(vault common.Address, partnerId common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YPartnerTracker.Contract.Deposit(&_YPartnerTracker.TransactOpts, vault, partnerId, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(address vault, address partnerId, uint256 amount) returns(uint256)
func (_YPartnerTracker *YPartnerTrackerTransactorSession) Deposit(vault common.Address, partnerId common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YPartnerTracker.Contract.Deposit(&_YPartnerTracker.TransactOpts, vault, partnerId, amount)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xf9609f08.
//
// Solidity: function deposit(address vault, address partnerId) returns(uint256)
func (_YPartnerTracker *YPartnerTrackerTransactor) Deposit0(opts *bind.TransactOpts, vault common.Address, partnerId common.Address) (*types.Transaction, error) {
	return _YPartnerTracker.contract.Transact(opts, "deposit0", vault, partnerId)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xf9609f08.
//
// Solidity: function deposit(address vault, address partnerId) returns(uint256)
func (_YPartnerTracker *YPartnerTrackerSession) Deposit0(vault common.Address, partnerId common.Address) (*types.Transaction, error) {
	return _YPartnerTracker.Contract.Deposit0(&_YPartnerTracker.TransactOpts, vault, partnerId)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xf9609f08.
//
// Solidity: function deposit(address vault, address partnerId) returns(uint256)
func (_YPartnerTracker *YPartnerTrackerTransactorSession) Deposit0(vault common.Address, partnerId common.Address) (*types.Transaction, error) {
	return _YPartnerTracker.Contract.Deposit0(&_YPartnerTracker.TransactOpts, vault, partnerId)
}

// YPartnerTrackerReferredBalanceIncreasedIterator is returned from FilterReferredBalanceIncreased and is used to iterate over the raw logs and unpacked data for ReferredBalanceIncreased events raised by the YPartnerTracker contract.
type YPartnerTrackerReferredBalanceIncreasedIterator struct {
	Event *YPartnerTrackerReferredBalanceIncreased // Event containing the contract specifics and raw log

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
func (it *YPartnerTrackerReferredBalanceIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YPartnerTrackerReferredBalanceIncreased)
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
		it.Event = new(YPartnerTrackerReferredBalanceIncreased)
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
func (it *YPartnerTrackerReferredBalanceIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YPartnerTrackerReferredBalanceIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YPartnerTrackerReferredBalanceIncreased represents a ReferredBalanceIncreased event raised by the YPartnerTracker contract.
type YPartnerTrackerReferredBalanceIncreased struct {
	PartnerId      common.Address
	Vault          common.Address
	Depositer      common.Address
	AmountAdded    *big.Int
	TotalDeposited *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterReferredBalanceIncreased is a free log retrieval operation binding the contract event 0xdb11012627bcb857ec915e92e4c5a65c80098ea77290a7b32d59a8ef2b3ea41b.
//
// Solidity: event ReferredBalanceIncreased(address indexed partnerId, address indexed vault, address indexed depositer, uint256 amountAdded, uint256 totalDeposited)
func (_YPartnerTracker *YPartnerTrackerFilterer) FilterReferredBalanceIncreased(opts *bind.FilterOpts, partnerId []common.Address, vault []common.Address, depositer []common.Address) (*YPartnerTrackerReferredBalanceIncreasedIterator, error) {

	var partnerIdRule []interface{}
	for _, partnerIdItem := range partnerId {
		partnerIdRule = append(partnerIdRule, partnerIdItem)
	}
	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var depositerRule []interface{}
	for _, depositerItem := range depositer {
		depositerRule = append(depositerRule, depositerItem)
	}

	logs, sub, err := _YPartnerTracker.contract.FilterLogs(opts, "ReferredBalanceIncreased", partnerIdRule, vaultRule, depositerRule)
	if err != nil {
		return nil, err
	}
	return &YPartnerTrackerReferredBalanceIncreasedIterator{contract: _YPartnerTracker.contract, event: "ReferredBalanceIncreased", logs: logs, sub: sub}, nil
}

// WatchReferredBalanceIncreased is a free log subscription operation binding the contract event 0xdb11012627bcb857ec915e92e4c5a65c80098ea77290a7b32d59a8ef2b3ea41b.
//
// Solidity: event ReferredBalanceIncreased(address indexed partnerId, address indexed vault, address indexed depositer, uint256 amountAdded, uint256 totalDeposited)
func (_YPartnerTracker *YPartnerTrackerFilterer) WatchReferredBalanceIncreased(opts *bind.WatchOpts, sink chan<- *YPartnerTrackerReferredBalanceIncreased, partnerId []common.Address, vault []common.Address, depositer []common.Address) (event.Subscription, error) {

	var partnerIdRule []interface{}
	for _, partnerIdItem := range partnerId {
		partnerIdRule = append(partnerIdRule, partnerIdItem)
	}
	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var depositerRule []interface{}
	for _, depositerItem := range depositer {
		depositerRule = append(depositerRule, depositerItem)
	}

	logs, sub, err := _YPartnerTracker.contract.WatchLogs(opts, "ReferredBalanceIncreased", partnerIdRule, vaultRule, depositerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YPartnerTrackerReferredBalanceIncreased)
				if err := _YPartnerTracker.contract.UnpackLog(event, "ReferredBalanceIncreased", log); err != nil {
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

// ParseReferredBalanceIncreased is a log parse operation binding the contract event 0xdb11012627bcb857ec915e92e4c5a65c80098ea77290a7b32d59a8ef2b3ea41b.
//
// Solidity: event ReferredBalanceIncreased(address indexed partnerId, address indexed vault, address indexed depositer, uint256 amountAdded, uint256 totalDeposited)
func (_YPartnerTracker *YPartnerTrackerFilterer) ParseReferredBalanceIncreased(log types.Log) (*YPartnerTrackerReferredBalanceIncreased, error) {
	event := new(YPartnerTrackerReferredBalanceIncreased)
	if err := _YPartnerTracker.contract.UnpackLog(event, "ReferredBalanceIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
