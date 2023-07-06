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
	_ = abi.ConvertType
)

// CRVUSDStruct0 is an auto generated low-level Go binding around an user-defined struct.
type CRVUSDStruct0 struct {
	Pool      common.Address
	IsInverse bool
}

// CrvUSDMetaData contains all meta data concerning the CrvUSD contract.
var CrvUSDMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"AddPricePair\",\"inputs\":[{\"name\":\"n\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"pool\",\"type\":\"address\",\"indexed\":false},{\"name\":\"is_inverse\",\"type\":\"bool\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemovePricePair\",\"inputs\":[{\"name\":\"n\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"MovePricePair\",\"inputs\":[{\"name\":\"n_from\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"n_to\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SetAdmin\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"stablecoin\",\"type\":\"address\"},{\"name\":\"sigma\",\"type\":\"uint256\"},{\"name\":\"admin\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_admin\",\"inputs\":[{\"name\":\"_admin\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"sigma\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"stablecoin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_price_pair\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_price_pair\",\"inputs\":[{\"name\":\"n\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"price\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"price_pairs\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"components\":[{\"name\":\"pool\",\"type\":\"address\"},{\"name\":\"is_inverse\",\"type\":\"bool\"}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]}]",
}

// CrvUSDABI is the input ABI used to generate the binding from.
// Deprecated: Use CrvUSDMetaData.ABI instead.
var CrvUSDABI = CrvUSDMetaData.ABI

// CrvUSD is an auto generated Go binding around an Ethereum contract.
type CrvUSD struct {
	CrvUSDCaller     // Read-only binding to the contract
	CrvUSDTransactor // Write-only binding to the contract
	CrvUSDFilterer   // Log filterer for contract events
}

// CrvUSDCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrvUSDCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrvUSDTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrvUSDTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrvUSDFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrvUSDFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrvUSDSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrvUSDSession struct {
	Contract     *CrvUSD           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrvUSDCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrvUSDCallerSession struct {
	Contract *CrvUSDCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CrvUSDTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrvUSDTransactorSession struct {
	Contract     *CrvUSDTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrvUSDRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrvUSDRaw struct {
	Contract *CrvUSD // Generic contract binding to access the raw methods on
}

// CrvUSDCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrvUSDCallerRaw struct {
	Contract *CrvUSDCaller // Generic read-only contract binding to access the raw methods on
}

// CrvUSDTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrvUSDTransactorRaw struct {
	Contract *CrvUSDTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrvUSD creates a new instance of CrvUSD, bound to a specific deployed contract.
func NewCrvUSD(address common.Address, backend bind.ContractBackend) (*CrvUSD, error) {
	contract, err := bindCrvUSD(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrvUSD{CrvUSDCaller: CrvUSDCaller{contract: contract}, CrvUSDTransactor: CrvUSDTransactor{contract: contract}, CrvUSDFilterer: CrvUSDFilterer{contract: contract}}, nil
}

// NewCrvUSDCaller creates a new read-only instance of CrvUSD, bound to a specific deployed contract.
func NewCrvUSDCaller(address common.Address, caller bind.ContractCaller) (*CrvUSDCaller, error) {
	contract, err := bindCrvUSD(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrvUSDCaller{contract: contract}, nil
}

// NewCrvUSDTransactor creates a new write-only instance of CrvUSD, bound to a specific deployed contract.
func NewCrvUSDTransactor(address common.Address, transactor bind.ContractTransactor) (*CrvUSDTransactor, error) {
	contract, err := bindCrvUSD(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrvUSDTransactor{contract: contract}, nil
}

// NewCrvUSDFilterer creates a new log filterer instance of CrvUSD, bound to a specific deployed contract.
func NewCrvUSDFilterer(address common.Address, filterer bind.ContractFilterer) (*CrvUSDFilterer, error) {
	contract, err := bindCrvUSD(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrvUSDFilterer{contract: contract}, nil
}

// bindCrvUSD binds a generic wrapper to an already deployed contract.
func bindCrvUSD(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CrvUSDMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrvUSD *CrvUSDRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrvUSD.Contract.CrvUSDCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrvUSD *CrvUSDRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrvUSD.Contract.CrvUSDTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrvUSD *CrvUSDRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrvUSD.Contract.CrvUSDTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrvUSD *CrvUSDCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrvUSD.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrvUSD *CrvUSDTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrvUSD.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrvUSD *CrvUSDTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrvUSD.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CrvUSD *CrvUSDCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrvUSD.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CrvUSD *CrvUSDSession) Admin() (common.Address, error) {
	return _CrvUSD.Contract.Admin(&_CrvUSD.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CrvUSD *CrvUSDCallerSession) Admin() (common.Address, error) {
	return _CrvUSD.Contract.Admin(&_CrvUSD.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_CrvUSD *CrvUSDCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvUSD.contract.Call(opts, &out, "price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_CrvUSD *CrvUSDSession) Price() (*big.Int, error) {
	return _CrvUSD.Contract.Price(&_CrvUSD.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_CrvUSD *CrvUSDCallerSession) Price() (*big.Int, error) {
	return _CrvUSD.Contract.Price(&_CrvUSD.CallOpts)
}

// PricePairs is a free data retrieval call binding the contract method 0xba5feb37.
//
// Solidity: function price_pairs(uint256 arg0) view returns((address,bool))
func (_CrvUSD *CrvUSDCaller) PricePairs(opts *bind.CallOpts, arg0 *big.Int) (CRVUSDStruct0, error) {
	var out []interface{}
	err := _CrvUSD.contract.Call(opts, &out, "price_pairs", arg0)

	if err != nil {
		return *new(CRVUSDStruct0), err
	}

	out0 := *abi.ConvertType(out[0], new(CRVUSDStruct0)).(*CRVUSDStruct0)

	return out0, err

}

// PricePairs is a free data retrieval call binding the contract method 0xba5feb37.
//
// Solidity: function price_pairs(uint256 arg0) view returns((address,bool))
func (_CrvUSD *CrvUSDSession) PricePairs(arg0 *big.Int) (CRVUSDStruct0, error) {
	return _CrvUSD.Contract.PricePairs(&_CrvUSD.CallOpts, arg0)
}

// PricePairs is a free data retrieval call binding the contract method 0xba5feb37.
//
// Solidity: function price_pairs(uint256 arg0) view returns((address,bool))
func (_CrvUSD *CrvUSDCallerSession) PricePairs(arg0 *big.Int) (CRVUSDStruct0, error) {
	return _CrvUSD.Contract.PricePairs(&_CrvUSD.CallOpts, arg0)
}

// Sigma is a free data retrieval call binding the contract method 0xafdf31cd.
//
// Solidity: function sigma() view returns(uint256)
func (_CrvUSD *CrvUSDCaller) Sigma(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvUSD.contract.Call(opts, &out, "sigma")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Sigma is a free data retrieval call binding the contract method 0xafdf31cd.
//
// Solidity: function sigma() view returns(uint256)
func (_CrvUSD *CrvUSDSession) Sigma() (*big.Int, error) {
	return _CrvUSD.Contract.Sigma(&_CrvUSD.CallOpts)
}

// Sigma is a free data retrieval call binding the contract method 0xafdf31cd.
//
// Solidity: function sigma() view returns(uint256)
func (_CrvUSD *CrvUSDCallerSession) Sigma() (*big.Int, error) {
	return _CrvUSD.Contract.Sigma(&_CrvUSD.CallOpts)
}

// Stablecoin is a free data retrieval call binding the contract method 0xe9cbd822.
//
// Solidity: function stablecoin() view returns(address)
func (_CrvUSD *CrvUSDCaller) Stablecoin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrvUSD.contract.Call(opts, &out, "stablecoin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Stablecoin is a free data retrieval call binding the contract method 0xe9cbd822.
//
// Solidity: function stablecoin() view returns(address)
func (_CrvUSD *CrvUSDSession) Stablecoin() (common.Address, error) {
	return _CrvUSD.Contract.Stablecoin(&_CrvUSD.CallOpts)
}

// Stablecoin is a free data retrieval call binding the contract method 0xe9cbd822.
//
// Solidity: function stablecoin() view returns(address)
func (_CrvUSD *CrvUSDCallerSession) Stablecoin() (common.Address, error) {
	return _CrvUSD.Contract.Stablecoin(&_CrvUSD.CallOpts)
}

// AddPricePair is a paid mutator transaction binding the contract method 0xa51ec1dd.
//
// Solidity: function add_price_pair(address _pool) returns()
func (_CrvUSD *CrvUSDTransactor) AddPricePair(opts *bind.TransactOpts, _pool common.Address) (*types.Transaction, error) {
	return _CrvUSD.contract.Transact(opts, "add_price_pair", _pool)
}

// AddPricePair is a paid mutator transaction binding the contract method 0xa51ec1dd.
//
// Solidity: function add_price_pair(address _pool) returns()
func (_CrvUSD *CrvUSDSession) AddPricePair(_pool common.Address) (*types.Transaction, error) {
	return _CrvUSD.Contract.AddPricePair(&_CrvUSD.TransactOpts, _pool)
}

// AddPricePair is a paid mutator transaction binding the contract method 0xa51ec1dd.
//
// Solidity: function add_price_pair(address _pool) returns()
func (_CrvUSD *CrvUSDTransactorSession) AddPricePair(_pool common.Address) (*types.Transaction, error) {
	return _CrvUSD.Contract.AddPricePair(&_CrvUSD.TransactOpts, _pool)
}

// RemovePricePair is a paid mutator transaction binding the contract method 0xb3162fdb.
//
// Solidity: function remove_price_pair(uint256 n) returns()
func (_CrvUSD *CrvUSDTransactor) RemovePricePair(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _CrvUSD.contract.Transact(opts, "remove_price_pair", n)
}

// RemovePricePair is a paid mutator transaction binding the contract method 0xb3162fdb.
//
// Solidity: function remove_price_pair(uint256 n) returns()
func (_CrvUSD *CrvUSDSession) RemovePricePair(n *big.Int) (*types.Transaction, error) {
	return _CrvUSD.Contract.RemovePricePair(&_CrvUSD.TransactOpts, n)
}

// RemovePricePair is a paid mutator transaction binding the contract method 0xb3162fdb.
//
// Solidity: function remove_price_pair(uint256 n) returns()
func (_CrvUSD *CrvUSDTransactorSession) RemovePricePair(n *big.Int) (*types.Transaction, error) {
	return _CrvUSD.Contract.RemovePricePair(&_CrvUSD.TransactOpts, n)
}

// SetAdmin is a paid mutator transaction binding the contract method 0xe9333fab.
//
// Solidity: function set_admin(address _admin) returns()
func (_CrvUSD *CrvUSDTransactor) SetAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _CrvUSD.contract.Transact(opts, "set_admin", _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0xe9333fab.
//
// Solidity: function set_admin(address _admin) returns()
func (_CrvUSD *CrvUSDSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _CrvUSD.Contract.SetAdmin(&_CrvUSD.TransactOpts, _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0xe9333fab.
//
// Solidity: function set_admin(address _admin) returns()
func (_CrvUSD *CrvUSDTransactorSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _CrvUSD.Contract.SetAdmin(&_CrvUSD.TransactOpts, _admin)
}

// CrvUSDAddPricePairIterator is returned from FilterAddPricePair and is used to iterate over the raw logs and unpacked data for AddPricePair events raised by the CrvUSD contract.
type CrvUSDAddPricePairIterator struct {
	Event *CrvUSDAddPricePair // Event containing the contract specifics and raw log

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
func (it *CrvUSDAddPricePairIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvUSDAddPricePair)
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
		it.Event = new(CrvUSDAddPricePair)
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
func (it *CrvUSDAddPricePairIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvUSDAddPricePairIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvUSDAddPricePair represents a AddPricePair event raised by the CrvUSD contract.
type CrvUSDAddPricePair struct {
	N         *big.Int
	Pool      common.Address
	IsInverse bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAddPricePair is a free log retrieval operation binding the contract event 0xfa41d19543baad0c0257c3430e1b3cecfa57f753cdc21330b9ec0862cde1b258.
//
// Solidity: event AddPricePair(uint256 n, address pool, bool is_inverse)
func (_CrvUSD *CrvUSDFilterer) FilterAddPricePair(opts *bind.FilterOpts) (*CrvUSDAddPricePairIterator, error) {

	logs, sub, err := _CrvUSD.contract.FilterLogs(opts, "AddPricePair")
	if err != nil {
		return nil, err
	}
	return &CrvUSDAddPricePairIterator{contract: _CrvUSD.contract, event: "AddPricePair", logs: logs, sub: sub}, nil
}

// WatchAddPricePair is a free log subscription operation binding the contract event 0xfa41d19543baad0c0257c3430e1b3cecfa57f753cdc21330b9ec0862cde1b258.
//
// Solidity: event AddPricePair(uint256 n, address pool, bool is_inverse)
func (_CrvUSD *CrvUSDFilterer) WatchAddPricePair(opts *bind.WatchOpts, sink chan<- *CrvUSDAddPricePair) (event.Subscription, error) {

	logs, sub, err := _CrvUSD.contract.WatchLogs(opts, "AddPricePair")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvUSDAddPricePair)
				if err := _CrvUSD.contract.UnpackLog(event, "AddPricePair", log); err != nil {
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

// ParseAddPricePair is a log parse operation binding the contract event 0xfa41d19543baad0c0257c3430e1b3cecfa57f753cdc21330b9ec0862cde1b258.
//
// Solidity: event AddPricePair(uint256 n, address pool, bool is_inverse)
func (_CrvUSD *CrvUSDFilterer) ParseAddPricePair(log types.Log) (*CrvUSDAddPricePair, error) {
	event := new(CrvUSDAddPricePair)
	if err := _CrvUSD.contract.UnpackLog(event, "AddPricePair", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrvUSDMovePricePairIterator is returned from FilterMovePricePair and is used to iterate over the raw logs and unpacked data for MovePricePair events raised by the CrvUSD contract.
type CrvUSDMovePricePairIterator struct {
	Event *CrvUSDMovePricePair // Event containing the contract specifics and raw log

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
func (it *CrvUSDMovePricePairIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvUSDMovePricePair)
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
		it.Event = new(CrvUSDMovePricePair)
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
func (it *CrvUSDMovePricePairIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvUSDMovePricePairIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvUSDMovePricePair represents a MovePricePair event raised by the CrvUSD contract.
type CrvUSDMovePricePair struct {
	NFrom *big.Int
	NTo   *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMovePricePair is a free log retrieval operation binding the contract event 0x3be85492dc07bf3888ee5d022674f0eced07fa5845d47b118257a8a0616d97ec.
//
// Solidity: event MovePricePair(uint256 n_from, uint256 n_to)
func (_CrvUSD *CrvUSDFilterer) FilterMovePricePair(opts *bind.FilterOpts) (*CrvUSDMovePricePairIterator, error) {

	logs, sub, err := _CrvUSD.contract.FilterLogs(opts, "MovePricePair")
	if err != nil {
		return nil, err
	}
	return &CrvUSDMovePricePairIterator{contract: _CrvUSD.contract, event: "MovePricePair", logs: logs, sub: sub}, nil
}

// WatchMovePricePair is a free log subscription operation binding the contract event 0x3be85492dc07bf3888ee5d022674f0eced07fa5845d47b118257a8a0616d97ec.
//
// Solidity: event MovePricePair(uint256 n_from, uint256 n_to)
func (_CrvUSD *CrvUSDFilterer) WatchMovePricePair(opts *bind.WatchOpts, sink chan<- *CrvUSDMovePricePair) (event.Subscription, error) {

	logs, sub, err := _CrvUSD.contract.WatchLogs(opts, "MovePricePair")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvUSDMovePricePair)
				if err := _CrvUSD.contract.UnpackLog(event, "MovePricePair", log); err != nil {
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

// ParseMovePricePair is a log parse operation binding the contract event 0x3be85492dc07bf3888ee5d022674f0eced07fa5845d47b118257a8a0616d97ec.
//
// Solidity: event MovePricePair(uint256 n_from, uint256 n_to)
func (_CrvUSD *CrvUSDFilterer) ParseMovePricePair(log types.Log) (*CrvUSDMovePricePair, error) {
	event := new(CrvUSDMovePricePair)
	if err := _CrvUSD.contract.UnpackLog(event, "MovePricePair", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrvUSDRemovePricePairIterator is returned from FilterRemovePricePair and is used to iterate over the raw logs and unpacked data for RemovePricePair events raised by the CrvUSD contract.
type CrvUSDRemovePricePairIterator struct {
	Event *CrvUSDRemovePricePair // Event containing the contract specifics and raw log

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
func (it *CrvUSDRemovePricePairIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvUSDRemovePricePair)
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
		it.Event = new(CrvUSDRemovePricePair)
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
func (it *CrvUSDRemovePricePairIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvUSDRemovePricePairIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvUSDRemovePricePair represents a RemovePricePair event raised by the CrvUSD contract.
type CrvUSDRemovePricePair struct {
	N   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRemovePricePair is a free log retrieval operation binding the contract event 0x017592f2f16e82cccce60102865c737270289c308f34ff88e754d5e99ea0bae1.
//
// Solidity: event RemovePricePair(uint256 n)
func (_CrvUSD *CrvUSDFilterer) FilterRemovePricePair(opts *bind.FilterOpts) (*CrvUSDRemovePricePairIterator, error) {

	logs, sub, err := _CrvUSD.contract.FilterLogs(opts, "RemovePricePair")
	if err != nil {
		return nil, err
	}
	return &CrvUSDRemovePricePairIterator{contract: _CrvUSD.contract, event: "RemovePricePair", logs: logs, sub: sub}, nil
}

// WatchRemovePricePair is a free log subscription operation binding the contract event 0x017592f2f16e82cccce60102865c737270289c308f34ff88e754d5e99ea0bae1.
//
// Solidity: event RemovePricePair(uint256 n)
func (_CrvUSD *CrvUSDFilterer) WatchRemovePricePair(opts *bind.WatchOpts, sink chan<- *CrvUSDRemovePricePair) (event.Subscription, error) {

	logs, sub, err := _CrvUSD.contract.WatchLogs(opts, "RemovePricePair")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvUSDRemovePricePair)
				if err := _CrvUSD.contract.UnpackLog(event, "RemovePricePair", log); err != nil {
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

// ParseRemovePricePair is a log parse operation binding the contract event 0x017592f2f16e82cccce60102865c737270289c308f34ff88e754d5e99ea0bae1.
//
// Solidity: event RemovePricePair(uint256 n)
func (_CrvUSD *CrvUSDFilterer) ParseRemovePricePair(log types.Log) (*CrvUSDRemovePricePair, error) {
	event := new(CrvUSDRemovePricePair)
	if err := _CrvUSD.contract.UnpackLog(event, "RemovePricePair", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrvUSDSetAdminIterator is returned from FilterSetAdmin and is used to iterate over the raw logs and unpacked data for SetAdmin events raised by the CrvUSD contract.
type CrvUSDSetAdminIterator struct {
	Event *CrvUSDSetAdmin // Event containing the contract specifics and raw log

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
func (it *CrvUSDSetAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvUSDSetAdmin)
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
		it.Event = new(CrvUSDSetAdmin)
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
func (it *CrvUSDSetAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvUSDSetAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvUSDSetAdmin represents a SetAdmin event raised by the CrvUSD contract.
type CrvUSDSetAdmin struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetAdmin is a free log retrieval operation binding the contract event 0x5a272403b402d892977df56625f4164ccaf70ca3863991c43ecfe76a6905b0a1.
//
// Solidity: event SetAdmin(address admin)
func (_CrvUSD *CrvUSDFilterer) FilterSetAdmin(opts *bind.FilterOpts) (*CrvUSDSetAdminIterator, error) {

	logs, sub, err := _CrvUSD.contract.FilterLogs(opts, "SetAdmin")
	if err != nil {
		return nil, err
	}
	return &CrvUSDSetAdminIterator{contract: _CrvUSD.contract, event: "SetAdmin", logs: logs, sub: sub}, nil
}

// WatchSetAdmin is a free log subscription operation binding the contract event 0x5a272403b402d892977df56625f4164ccaf70ca3863991c43ecfe76a6905b0a1.
//
// Solidity: event SetAdmin(address admin)
func (_CrvUSD *CrvUSDFilterer) WatchSetAdmin(opts *bind.WatchOpts, sink chan<- *CrvUSDSetAdmin) (event.Subscription, error) {

	logs, sub, err := _CrvUSD.contract.WatchLogs(opts, "SetAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvUSDSetAdmin)
				if err := _CrvUSD.contract.UnpackLog(event, "SetAdmin", log); err != nil {
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

// ParseSetAdmin is a log parse operation binding the contract event 0x5a272403b402d892977df56625f4164ccaf70ca3863991c43ecfe76a6905b0a1.
//
// Solidity: event SetAdmin(address admin)
func (_CrvUSD *CrvUSDFilterer) ParseSetAdmin(log types.Log) (*CrvUSDSetAdmin, error) {
	event := new(CrvUSDSetAdmin)
	if err := _CrvUSD.contract.UnpackLog(event, "SetAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
