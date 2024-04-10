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

// YGaugeRegistryMetaData contains all meta data concerning the YGaugeRegistry contract.
var YGaugeRegistryMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"Register\",\"inputs\":[{\"name\":\"gauge\",\"type\":\"address\",\"indexed\":true},{\"name\":\"idx\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Deregister\",\"inputs\":[{\"name\":\"gauge\",\"type\":\"address\",\"indexed\":true},{\"name\":\"idx\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateIndex\",\"inputs\":[{\"name\":\"old_idx\",\"type\":\"uint256\",\"indexed\":true},{\"name\":\"idx\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SetController\",\"inputs\":[{\"name\":\"controller\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SetFactory\",\"inputs\":[{\"name\":\"factory\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"PendingManagement\",\"inputs\":[{\"name\":\"management\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SetManagement\",\"inputs\":[{\"name\":\"management\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"_controller\",\"type\":\"address\"},{\"name\":\"_factory\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"gauges\",\"inputs\":[{\"name\":\"_idx\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"register\",\"inputs\":[{\"name\":\"_gauge\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deregister\",\"inputs\":[{\"name\":\"_gauge\",\"type\":\"address\"},{\"name\":\"_idx\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"registered\",\"inputs\":[{\"name\":\"_gauge\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_controller\",\"inputs\":[{\"name\":\"_controller\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_factory\",\"inputs\":[{\"name\":\"_factory\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_management\",\"inputs\":[{\"name\":\"_management\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"accept_management\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"management\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"pending_management\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"controller\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"factory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"vault_count\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"vaults\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"vault_gauge_map\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]}]",
}

// YGaugeRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use YGaugeRegistryMetaData.ABI instead.
var YGaugeRegistryABI = YGaugeRegistryMetaData.ABI

// YGaugeRegistry is an auto generated Go binding around an Ethereum contract.
type YGaugeRegistry struct {
	YGaugeRegistryCaller     // Read-only binding to the contract
	YGaugeRegistryTransactor // Write-only binding to the contract
	YGaugeRegistryFilterer   // Log filterer for contract events
}

// YGaugeRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type YGaugeRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YGaugeRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YGaugeRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YGaugeRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YGaugeRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YGaugeRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YGaugeRegistrySession struct {
	Contract     *YGaugeRegistry   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YGaugeRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YGaugeRegistryCallerSession struct {
	Contract *YGaugeRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// YGaugeRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YGaugeRegistryTransactorSession struct {
	Contract     *YGaugeRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// YGaugeRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type YGaugeRegistryRaw struct {
	Contract *YGaugeRegistry // Generic contract binding to access the raw methods on
}

// YGaugeRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YGaugeRegistryCallerRaw struct {
	Contract *YGaugeRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// YGaugeRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YGaugeRegistryTransactorRaw struct {
	Contract *YGaugeRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYGaugeRegistry creates a new instance of YGaugeRegistry, bound to a specific deployed contract.
func NewYGaugeRegistry(address common.Address, backend bind.ContractBackend) (*YGaugeRegistry, error) {
	contract, err := bindYGaugeRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YGaugeRegistry{YGaugeRegistryCaller: YGaugeRegistryCaller{contract: contract}, YGaugeRegistryTransactor: YGaugeRegistryTransactor{contract: contract}, YGaugeRegistryFilterer: YGaugeRegistryFilterer{contract: contract}}, nil
}

// NewYGaugeRegistryCaller creates a new read-only instance of YGaugeRegistry, bound to a specific deployed contract.
func NewYGaugeRegistryCaller(address common.Address, caller bind.ContractCaller) (*YGaugeRegistryCaller, error) {
	contract, err := bindYGaugeRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YGaugeRegistryCaller{contract: contract}, nil
}

// NewYGaugeRegistryTransactor creates a new write-only instance of YGaugeRegistry, bound to a specific deployed contract.
func NewYGaugeRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*YGaugeRegistryTransactor, error) {
	contract, err := bindYGaugeRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YGaugeRegistryTransactor{contract: contract}, nil
}

// NewYGaugeRegistryFilterer creates a new log filterer instance of YGaugeRegistry, bound to a specific deployed contract.
func NewYGaugeRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*YGaugeRegistryFilterer, error) {
	contract, err := bindYGaugeRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YGaugeRegistryFilterer{contract: contract}, nil
}

// bindYGaugeRegistry binds a generic wrapper to an already deployed contract.
func bindYGaugeRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := YGaugeRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YGaugeRegistry *YGaugeRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YGaugeRegistry.Contract.YGaugeRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YGaugeRegistry *YGaugeRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YGaugeRegistry.Contract.YGaugeRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YGaugeRegistry *YGaugeRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YGaugeRegistry.Contract.YGaugeRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YGaugeRegistry *YGaugeRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YGaugeRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YGaugeRegistry *YGaugeRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YGaugeRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YGaugeRegistry *YGaugeRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YGaugeRegistry.Contract.contract.Transact(opts, method, params...)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_YGaugeRegistry *YGaugeRegistryCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YGaugeRegistry.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_YGaugeRegistry *YGaugeRegistrySession) Controller() (common.Address, error) {
	return _YGaugeRegistry.Contract.Controller(&_YGaugeRegistry.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_YGaugeRegistry *YGaugeRegistryCallerSession) Controller() (common.Address, error) {
	return _YGaugeRegistry.Contract.Controller(&_YGaugeRegistry.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_YGaugeRegistry *YGaugeRegistryCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YGaugeRegistry.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_YGaugeRegistry *YGaugeRegistrySession) Factory() (common.Address, error) {
	return _YGaugeRegistry.Contract.Factory(&_YGaugeRegistry.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_YGaugeRegistry *YGaugeRegistryCallerSession) Factory() (common.Address, error) {
	return _YGaugeRegistry.Contract.Factory(&_YGaugeRegistry.CallOpts)
}

// Gauges is a free data retrieval call binding the contract method 0xb0539187.
//
// Solidity: function gauges(uint256 _idx) view returns(address)
func (_YGaugeRegistry *YGaugeRegistryCaller) Gauges(opts *bind.CallOpts, _idx *big.Int) (common.Address, error) {
	var out []interface{}
	err := _YGaugeRegistry.contract.Call(opts, &out, "gauges", _idx)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gauges is a free data retrieval call binding the contract method 0xb0539187.
//
// Solidity: function gauges(uint256 _idx) view returns(address)
func (_YGaugeRegistry *YGaugeRegistrySession) Gauges(_idx *big.Int) (common.Address, error) {
	return _YGaugeRegistry.Contract.Gauges(&_YGaugeRegistry.CallOpts, _idx)
}

// Gauges is a free data retrieval call binding the contract method 0xb0539187.
//
// Solidity: function gauges(uint256 _idx) view returns(address)
func (_YGaugeRegistry *YGaugeRegistryCallerSession) Gauges(_idx *big.Int) (common.Address, error) {
	return _YGaugeRegistry.Contract.Gauges(&_YGaugeRegistry.CallOpts, _idx)
}

// Management is a free data retrieval call binding the contract method 0x88a8d602.
//
// Solidity: function management() view returns(address)
func (_YGaugeRegistry *YGaugeRegistryCaller) Management(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YGaugeRegistry.contract.Call(opts, &out, "management")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Management is a free data retrieval call binding the contract method 0x88a8d602.
//
// Solidity: function management() view returns(address)
func (_YGaugeRegistry *YGaugeRegistrySession) Management() (common.Address, error) {
	return _YGaugeRegistry.Contract.Management(&_YGaugeRegistry.CallOpts)
}

// Management is a free data retrieval call binding the contract method 0x88a8d602.
//
// Solidity: function management() view returns(address)
func (_YGaugeRegistry *YGaugeRegistryCallerSession) Management() (common.Address, error) {
	return _YGaugeRegistry.Contract.Management(&_YGaugeRegistry.CallOpts)
}

// PendingManagement is a free data retrieval call binding the contract method 0x770817ec.
//
// Solidity: function pending_management() view returns(address)
func (_YGaugeRegistry *YGaugeRegistryCaller) PendingManagement(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YGaugeRegistry.contract.Call(opts, &out, "pending_management")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingManagement is a free data retrieval call binding the contract method 0x770817ec.
//
// Solidity: function pending_management() view returns(address)
func (_YGaugeRegistry *YGaugeRegistrySession) PendingManagement() (common.Address, error) {
	return _YGaugeRegistry.Contract.PendingManagement(&_YGaugeRegistry.CallOpts)
}

// PendingManagement is a free data retrieval call binding the contract method 0x770817ec.
//
// Solidity: function pending_management() view returns(address)
func (_YGaugeRegistry *YGaugeRegistryCallerSession) PendingManagement() (common.Address, error) {
	return _YGaugeRegistry.Contract.PendingManagement(&_YGaugeRegistry.CallOpts)
}

// Registered is a free data retrieval call binding the contract method 0xb2dd5c07.
//
// Solidity: function registered(address _gauge) view returns(bool)
func (_YGaugeRegistry *YGaugeRegistryCaller) Registered(opts *bind.CallOpts, _gauge common.Address) (bool, error) {
	var out []interface{}
	err := _YGaugeRegistry.contract.Call(opts, &out, "registered", _gauge)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Registered is a free data retrieval call binding the contract method 0xb2dd5c07.
//
// Solidity: function registered(address _gauge) view returns(bool)
func (_YGaugeRegistry *YGaugeRegistrySession) Registered(_gauge common.Address) (bool, error) {
	return _YGaugeRegistry.Contract.Registered(&_YGaugeRegistry.CallOpts, _gauge)
}

// Registered is a free data retrieval call binding the contract method 0xb2dd5c07.
//
// Solidity: function registered(address _gauge) view returns(bool)
func (_YGaugeRegistry *YGaugeRegistryCallerSession) Registered(_gauge common.Address) (bool, error) {
	return _YGaugeRegistry.Contract.Registered(&_YGaugeRegistry.CallOpts, _gauge)
}

// VaultCount is a free data retrieval call binding the contract method 0xed20cb30.
//
// Solidity: function vault_count() view returns(uint256)
func (_YGaugeRegistry *YGaugeRegistryCaller) VaultCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YGaugeRegistry.contract.Call(opts, &out, "vault_count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VaultCount is a free data retrieval call binding the contract method 0xed20cb30.
//
// Solidity: function vault_count() view returns(uint256)
func (_YGaugeRegistry *YGaugeRegistrySession) VaultCount() (*big.Int, error) {
	return _YGaugeRegistry.Contract.VaultCount(&_YGaugeRegistry.CallOpts)
}

// VaultCount is a free data retrieval call binding the contract method 0xed20cb30.
//
// Solidity: function vault_count() view returns(uint256)
func (_YGaugeRegistry *YGaugeRegistryCallerSession) VaultCount() (*big.Int, error) {
	return _YGaugeRegistry.Contract.VaultCount(&_YGaugeRegistry.CallOpts)
}

// VaultGaugeMap is a free data retrieval call binding the contract method 0xdfe3d284.
//
// Solidity: function vault_gauge_map(address arg0) view returns(address)
func (_YGaugeRegistry *YGaugeRegistryCaller) VaultGaugeMap(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _YGaugeRegistry.contract.Call(opts, &out, "vault_gauge_map", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultGaugeMap is a free data retrieval call binding the contract method 0xdfe3d284.
//
// Solidity: function vault_gauge_map(address arg0) view returns(address)
func (_YGaugeRegistry *YGaugeRegistrySession) VaultGaugeMap(arg0 common.Address) (common.Address, error) {
	return _YGaugeRegistry.Contract.VaultGaugeMap(&_YGaugeRegistry.CallOpts, arg0)
}

// VaultGaugeMap is a free data retrieval call binding the contract method 0xdfe3d284.
//
// Solidity: function vault_gauge_map(address arg0) view returns(address)
func (_YGaugeRegistry *YGaugeRegistryCallerSession) VaultGaugeMap(arg0 common.Address) (common.Address, error) {
	return _YGaugeRegistry.Contract.VaultGaugeMap(&_YGaugeRegistry.CallOpts, arg0)
}

// Vaults is a free data retrieval call binding the contract method 0x8c64ea4a.
//
// Solidity: function vaults(uint256 arg0) view returns(address)
func (_YGaugeRegistry *YGaugeRegistryCaller) Vaults(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _YGaugeRegistry.contract.Call(opts, &out, "vaults", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vaults is a free data retrieval call binding the contract method 0x8c64ea4a.
//
// Solidity: function vaults(uint256 arg0) view returns(address)
func (_YGaugeRegistry *YGaugeRegistrySession) Vaults(arg0 *big.Int) (common.Address, error) {
	return _YGaugeRegistry.Contract.Vaults(&_YGaugeRegistry.CallOpts, arg0)
}

// Vaults is a free data retrieval call binding the contract method 0x8c64ea4a.
//
// Solidity: function vaults(uint256 arg0) view returns(address)
func (_YGaugeRegistry *YGaugeRegistryCallerSession) Vaults(arg0 *big.Int) (common.Address, error) {
	return _YGaugeRegistry.Contract.Vaults(&_YGaugeRegistry.CallOpts, arg0)
}

// AcceptManagement is a paid mutator transaction binding the contract method 0x759be10c.
//
// Solidity: function accept_management() returns()
func (_YGaugeRegistry *YGaugeRegistryTransactor) AcceptManagement(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YGaugeRegistry.contract.Transact(opts, "accept_management")
}

// AcceptManagement is a paid mutator transaction binding the contract method 0x759be10c.
//
// Solidity: function accept_management() returns()
func (_YGaugeRegistry *YGaugeRegistrySession) AcceptManagement() (*types.Transaction, error) {
	return _YGaugeRegistry.Contract.AcceptManagement(&_YGaugeRegistry.TransactOpts)
}

// AcceptManagement is a paid mutator transaction binding the contract method 0x759be10c.
//
// Solidity: function accept_management() returns()
func (_YGaugeRegistry *YGaugeRegistryTransactorSession) AcceptManagement() (*types.Transaction, error) {
	return _YGaugeRegistry.Contract.AcceptManagement(&_YGaugeRegistry.TransactOpts)
}

// Deregister is a paid mutator transaction binding the contract method 0xeef6c450.
//
// Solidity: function deregister(address _gauge, uint256 _idx) returns()
func (_YGaugeRegistry *YGaugeRegistryTransactor) Deregister(opts *bind.TransactOpts, _gauge common.Address, _idx *big.Int) (*types.Transaction, error) {
	return _YGaugeRegistry.contract.Transact(opts, "deregister", _gauge, _idx)
}

// Deregister is a paid mutator transaction binding the contract method 0xeef6c450.
//
// Solidity: function deregister(address _gauge, uint256 _idx) returns()
func (_YGaugeRegistry *YGaugeRegistrySession) Deregister(_gauge common.Address, _idx *big.Int) (*types.Transaction, error) {
	return _YGaugeRegistry.Contract.Deregister(&_YGaugeRegistry.TransactOpts, _gauge, _idx)
}

// Deregister is a paid mutator transaction binding the contract method 0xeef6c450.
//
// Solidity: function deregister(address _gauge, uint256 _idx) returns()
func (_YGaugeRegistry *YGaugeRegistryTransactorSession) Deregister(_gauge common.Address, _idx *big.Int) (*types.Transaction, error) {
	return _YGaugeRegistry.Contract.Deregister(&_YGaugeRegistry.TransactOpts, _gauge, _idx)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address _gauge) returns(uint256)
func (_YGaugeRegistry *YGaugeRegistryTransactor) Register(opts *bind.TransactOpts, _gauge common.Address) (*types.Transaction, error) {
	return _YGaugeRegistry.contract.Transact(opts, "register", _gauge)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address _gauge) returns(uint256)
func (_YGaugeRegistry *YGaugeRegistrySession) Register(_gauge common.Address) (*types.Transaction, error) {
	return _YGaugeRegistry.Contract.Register(&_YGaugeRegistry.TransactOpts, _gauge)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address _gauge) returns(uint256)
func (_YGaugeRegistry *YGaugeRegistryTransactorSession) Register(_gauge common.Address) (*types.Transaction, error) {
	return _YGaugeRegistry.Contract.Register(&_YGaugeRegistry.TransactOpts, _gauge)
}

// SetController is a paid mutator transaction binding the contract method 0x91b10ffa.
//
// Solidity: function set_controller(address _controller) returns()
func (_YGaugeRegistry *YGaugeRegistryTransactor) SetController(opts *bind.TransactOpts, _controller common.Address) (*types.Transaction, error) {
	return _YGaugeRegistry.contract.Transact(opts, "set_controller", _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x91b10ffa.
//
// Solidity: function set_controller(address _controller) returns()
func (_YGaugeRegistry *YGaugeRegistrySession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _YGaugeRegistry.Contract.SetController(&_YGaugeRegistry.TransactOpts, _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x91b10ffa.
//
// Solidity: function set_controller(address _controller) returns()
func (_YGaugeRegistry *YGaugeRegistryTransactorSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _YGaugeRegistry.Contract.SetController(&_YGaugeRegistry.TransactOpts, _controller)
}

// SetFactory is a paid mutator transaction binding the contract method 0xee2af3fb.
//
// Solidity: function set_factory(address _factory) returns()
func (_YGaugeRegistry *YGaugeRegistryTransactor) SetFactory(opts *bind.TransactOpts, _factory common.Address) (*types.Transaction, error) {
	return _YGaugeRegistry.contract.Transact(opts, "set_factory", _factory)
}

// SetFactory is a paid mutator transaction binding the contract method 0xee2af3fb.
//
// Solidity: function set_factory(address _factory) returns()
func (_YGaugeRegistry *YGaugeRegistrySession) SetFactory(_factory common.Address) (*types.Transaction, error) {
	return _YGaugeRegistry.Contract.SetFactory(&_YGaugeRegistry.TransactOpts, _factory)
}

// SetFactory is a paid mutator transaction binding the contract method 0xee2af3fb.
//
// Solidity: function set_factory(address _factory) returns()
func (_YGaugeRegistry *YGaugeRegistryTransactorSession) SetFactory(_factory common.Address) (*types.Transaction, error) {
	return _YGaugeRegistry.Contract.SetFactory(&_YGaugeRegistry.TransactOpts, _factory)
}

// SetManagement is a paid mutator transaction binding the contract method 0xfd066ecc.
//
// Solidity: function set_management(address _management) returns()
func (_YGaugeRegistry *YGaugeRegistryTransactor) SetManagement(opts *bind.TransactOpts, _management common.Address) (*types.Transaction, error) {
	return _YGaugeRegistry.contract.Transact(opts, "set_management", _management)
}

// SetManagement is a paid mutator transaction binding the contract method 0xfd066ecc.
//
// Solidity: function set_management(address _management) returns()
func (_YGaugeRegistry *YGaugeRegistrySession) SetManagement(_management common.Address) (*types.Transaction, error) {
	return _YGaugeRegistry.Contract.SetManagement(&_YGaugeRegistry.TransactOpts, _management)
}

// SetManagement is a paid mutator transaction binding the contract method 0xfd066ecc.
//
// Solidity: function set_management(address _management) returns()
func (_YGaugeRegistry *YGaugeRegistryTransactorSession) SetManagement(_management common.Address) (*types.Transaction, error) {
	return _YGaugeRegistry.Contract.SetManagement(&_YGaugeRegistry.TransactOpts, _management)
}

// YGaugeRegistryDeregisterIterator is returned from FilterDeregister and is used to iterate over the raw logs and unpacked data for Deregister events raised by the YGaugeRegistry contract.
type YGaugeRegistryDeregisterIterator struct {
	Event *YGaugeRegistryDeregister // Event containing the contract specifics and raw log

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
func (it *YGaugeRegistryDeregisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YGaugeRegistryDeregister)
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
		it.Event = new(YGaugeRegistryDeregister)
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
func (it *YGaugeRegistryDeregisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YGaugeRegistryDeregisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YGaugeRegistryDeregister represents a Deregister event raised by the YGaugeRegistry contract.
type YGaugeRegistryDeregister struct {
	Gauge common.Address
	Idx   *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDeregister is a free log retrieval operation binding the contract event 0xf974f9b9976eab09e5a66d08a91feff28233bd865bdc3124015b69640bb31f45.
//
// Solidity: event Deregister(address indexed gauge, uint256 idx)
func (_YGaugeRegistry *YGaugeRegistryFilterer) FilterDeregister(opts *bind.FilterOpts, gauge []common.Address) (*YGaugeRegistryDeregisterIterator, error) {

	var gaugeRule []interface{}
	for _, gaugeItem := range gauge {
		gaugeRule = append(gaugeRule, gaugeItem)
	}

	logs, sub, err := _YGaugeRegistry.contract.FilterLogs(opts, "Deregister", gaugeRule)
	if err != nil {
		return nil, err
	}
	return &YGaugeRegistryDeregisterIterator{contract: _YGaugeRegistry.contract, event: "Deregister", logs: logs, sub: sub}, nil
}

// WatchDeregister is a free log subscription operation binding the contract event 0xf974f9b9976eab09e5a66d08a91feff28233bd865bdc3124015b69640bb31f45.
//
// Solidity: event Deregister(address indexed gauge, uint256 idx)
func (_YGaugeRegistry *YGaugeRegistryFilterer) WatchDeregister(opts *bind.WatchOpts, sink chan<- *YGaugeRegistryDeregister, gauge []common.Address) (event.Subscription, error) {

	var gaugeRule []interface{}
	for _, gaugeItem := range gauge {
		gaugeRule = append(gaugeRule, gaugeItem)
	}

	logs, sub, err := _YGaugeRegistry.contract.WatchLogs(opts, "Deregister", gaugeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YGaugeRegistryDeregister)
				if err := _YGaugeRegistry.contract.UnpackLog(event, "Deregister", log); err != nil {
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

// ParseDeregister is a log parse operation binding the contract event 0xf974f9b9976eab09e5a66d08a91feff28233bd865bdc3124015b69640bb31f45.
//
// Solidity: event Deregister(address indexed gauge, uint256 idx)
func (_YGaugeRegistry *YGaugeRegistryFilterer) ParseDeregister(log types.Log) (*YGaugeRegistryDeregister, error) {
	event := new(YGaugeRegistryDeregister)
	if err := _YGaugeRegistry.contract.UnpackLog(event, "Deregister", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YGaugeRegistryPendingManagementIterator is returned from FilterPendingManagement and is used to iterate over the raw logs and unpacked data for PendingManagement events raised by the YGaugeRegistry contract.
type YGaugeRegistryPendingManagementIterator struct {
	Event *YGaugeRegistryPendingManagement // Event containing the contract specifics and raw log

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
func (it *YGaugeRegistryPendingManagementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YGaugeRegistryPendingManagement)
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
		it.Event = new(YGaugeRegistryPendingManagement)
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
func (it *YGaugeRegistryPendingManagementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YGaugeRegistryPendingManagementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YGaugeRegistryPendingManagement represents a PendingManagement event raised by the YGaugeRegistry contract.
type YGaugeRegistryPendingManagement struct {
	Management common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPendingManagement is a free log retrieval operation binding the contract event 0xe7b5cc087e6e47e33e86bdfe4720b7e849891938b18ff6e0c3f92299de79e60c.
//
// Solidity: event PendingManagement(address indexed management)
func (_YGaugeRegistry *YGaugeRegistryFilterer) FilterPendingManagement(opts *bind.FilterOpts, management []common.Address) (*YGaugeRegistryPendingManagementIterator, error) {

	var managementRule []interface{}
	for _, managementItem := range management {
		managementRule = append(managementRule, managementItem)
	}

	logs, sub, err := _YGaugeRegistry.contract.FilterLogs(opts, "PendingManagement", managementRule)
	if err != nil {
		return nil, err
	}
	return &YGaugeRegistryPendingManagementIterator{contract: _YGaugeRegistry.contract, event: "PendingManagement", logs: logs, sub: sub}, nil
}

// WatchPendingManagement is a free log subscription operation binding the contract event 0xe7b5cc087e6e47e33e86bdfe4720b7e849891938b18ff6e0c3f92299de79e60c.
//
// Solidity: event PendingManagement(address indexed management)
func (_YGaugeRegistry *YGaugeRegistryFilterer) WatchPendingManagement(opts *bind.WatchOpts, sink chan<- *YGaugeRegistryPendingManagement, management []common.Address) (event.Subscription, error) {

	var managementRule []interface{}
	for _, managementItem := range management {
		managementRule = append(managementRule, managementItem)
	}

	logs, sub, err := _YGaugeRegistry.contract.WatchLogs(opts, "PendingManagement", managementRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YGaugeRegistryPendingManagement)
				if err := _YGaugeRegistry.contract.UnpackLog(event, "PendingManagement", log); err != nil {
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

// ParsePendingManagement is a log parse operation binding the contract event 0xe7b5cc087e6e47e33e86bdfe4720b7e849891938b18ff6e0c3f92299de79e60c.
//
// Solidity: event PendingManagement(address indexed management)
func (_YGaugeRegistry *YGaugeRegistryFilterer) ParsePendingManagement(log types.Log) (*YGaugeRegistryPendingManagement, error) {
	event := new(YGaugeRegistryPendingManagement)
	if err := _YGaugeRegistry.contract.UnpackLog(event, "PendingManagement", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YGaugeRegistryRegisterIterator is returned from FilterRegister and is used to iterate over the raw logs and unpacked data for Register events raised by the YGaugeRegistry contract.
type YGaugeRegistryRegisterIterator struct {
	Event *YGaugeRegistryRegister // Event containing the contract specifics and raw log

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
func (it *YGaugeRegistryRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YGaugeRegistryRegister)
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
		it.Event = new(YGaugeRegistryRegister)
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
func (it *YGaugeRegistryRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YGaugeRegistryRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YGaugeRegistryRegister represents a Register event raised by the YGaugeRegistry contract.
type YGaugeRegistryRegister struct {
	Gauge common.Address
	Idx   *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRegister is a free log retrieval operation binding the contract event 0x007dc6ab80cc84c043b7b8d4fcafc802187470087f7ea7fccd2e17aecd0256a1.
//
// Solidity: event Register(address indexed gauge, uint256 idx)
func (_YGaugeRegistry *YGaugeRegistryFilterer) FilterRegister(opts *bind.FilterOpts, gauge []common.Address) (*YGaugeRegistryRegisterIterator, error) {

	var gaugeRule []interface{}
	for _, gaugeItem := range gauge {
		gaugeRule = append(gaugeRule, gaugeItem)
	}

	logs, sub, err := _YGaugeRegistry.contract.FilterLogs(opts, "Register", gaugeRule)
	if err != nil {
		return nil, err
	}
	return &YGaugeRegistryRegisterIterator{contract: _YGaugeRegistry.contract, event: "Register", logs: logs, sub: sub}, nil
}

// WatchRegister is a free log subscription operation binding the contract event 0x007dc6ab80cc84c043b7b8d4fcafc802187470087f7ea7fccd2e17aecd0256a1.
//
// Solidity: event Register(address indexed gauge, uint256 idx)
func (_YGaugeRegistry *YGaugeRegistryFilterer) WatchRegister(opts *bind.WatchOpts, sink chan<- *YGaugeRegistryRegister, gauge []common.Address) (event.Subscription, error) {

	var gaugeRule []interface{}
	for _, gaugeItem := range gauge {
		gaugeRule = append(gaugeRule, gaugeItem)
	}

	logs, sub, err := _YGaugeRegistry.contract.WatchLogs(opts, "Register", gaugeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YGaugeRegistryRegister)
				if err := _YGaugeRegistry.contract.UnpackLog(event, "Register", log); err != nil {
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

// ParseRegister is a log parse operation binding the contract event 0x007dc6ab80cc84c043b7b8d4fcafc802187470087f7ea7fccd2e17aecd0256a1.
//
// Solidity: event Register(address indexed gauge, uint256 idx)
func (_YGaugeRegistry *YGaugeRegistryFilterer) ParseRegister(log types.Log) (*YGaugeRegistryRegister, error) {
	event := new(YGaugeRegistryRegister)
	if err := _YGaugeRegistry.contract.UnpackLog(event, "Register", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YGaugeRegistrySetControllerIterator is returned from FilterSetController and is used to iterate over the raw logs and unpacked data for SetController events raised by the YGaugeRegistry contract.
type YGaugeRegistrySetControllerIterator struct {
	Event *YGaugeRegistrySetController // Event containing the contract specifics and raw log

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
func (it *YGaugeRegistrySetControllerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YGaugeRegistrySetController)
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
		it.Event = new(YGaugeRegistrySetController)
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
func (it *YGaugeRegistrySetControllerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YGaugeRegistrySetControllerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YGaugeRegistrySetController represents a SetController event raised by the YGaugeRegistry contract.
type YGaugeRegistrySetController struct {
	Controller common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetController is a free log retrieval operation binding the contract event 0x4ff638452bbf33c012645d18ae6f05515ff5f2d1dfb0cece8cbf018c60903f70.
//
// Solidity: event SetController(address controller)
func (_YGaugeRegistry *YGaugeRegistryFilterer) FilterSetController(opts *bind.FilterOpts) (*YGaugeRegistrySetControllerIterator, error) {

	logs, sub, err := _YGaugeRegistry.contract.FilterLogs(opts, "SetController")
	if err != nil {
		return nil, err
	}
	return &YGaugeRegistrySetControllerIterator{contract: _YGaugeRegistry.contract, event: "SetController", logs: logs, sub: sub}, nil
}

// WatchSetController is a free log subscription operation binding the contract event 0x4ff638452bbf33c012645d18ae6f05515ff5f2d1dfb0cece8cbf018c60903f70.
//
// Solidity: event SetController(address controller)
func (_YGaugeRegistry *YGaugeRegistryFilterer) WatchSetController(opts *bind.WatchOpts, sink chan<- *YGaugeRegistrySetController) (event.Subscription, error) {

	logs, sub, err := _YGaugeRegistry.contract.WatchLogs(opts, "SetController")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YGaugeRegistrySetController)
				if err := _YGaugeRegistry.contract.UnpackLog(event, "SetController", log); err != nil {
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

// ParseSetController is a log parse operation binding the contract event 0x4ff638452bbf33c012645d18ae6f05515ff5f2d1dfb0cece8cbf018c60903f70.
//
// Solidity: event SetController(address controller)
func (_YGaugeRegistry *YGaugeRegistryFilterer) ParseSetController(log types.Log) (*YGaugeRegistrySetController, error) {
	event := new(YGaugeRegistrySetController)
	if err := _YGaugeRegistry.contract.UnpackLog(event, "SetController", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YGaugeRegistrySetFactoryIterator is returned from FilterSetFactory and is used to iterate over the raw logs and unpacked data for SetFactory events raised by the YGaugeRegistry contract.
type YGaugeRegistrySetFactoryIterator struct {
	Event *YGaugeRegistrySetFactory // Event containing the contract specifics and raw log

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
func (it *YGaugeRegistrySetFactoryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YGaugeRegistrySetFactory)
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
		it.Event = new(YGaugeRegistrySetFactory)
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
func (it *YGaugeRegistrySetFactoryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YGaugeRegistrySetFactoryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YGaugeRegistrySetFactory represents a SetFactory event raised by the YGaugeRegistry contract.
type YGaugeRegistrySetFactory struct {
	Factory common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetFactory is a free log retrieval operation binding the contract event 0x1c893ef9379093af30f458b9e74d2aba13c499660b68dec5e29af7b199c188b9.
//
// Solidity: event SetFactory(address factory)
func (_YGaugeRegistry *YGaugeRegistryFilterer) FilterSetFactory(opts *bind.FilterOpts) (*YGaugeRegistrySetFactoryIterator, error) {

	logs, sub, err := _YGaugeRegistry.contract.FilterLogs(opts, "SetFactory")
	if err != nil {
		return nil, err
	}
	return &YGaugeRegistrySetFactoryIterator{contract: _YGaugeRegistry.contract, event: "SetFactory", logs: logs, sub: sub}, nil
}

// WatchSetFactory is a free log subscription operation binding the contract event 0x1c893ef9379093af30f458b9e74d2aba13c499660b68dec5e29af7b199c188b9.
//
// Solidity: event SetFactory(address factory)
func (_YGaugeRegistry *YGaugeRegistryFilterer) WatchSetFactory(opts *bind.WatchOpts, sink chan<- *YGaugeRegistrySetFactory) (event.Subscription, error) {

	logs, sub, err := _YGaugeRegistry.contract.WatchLogs(opts, "SetFactory")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YGaugeRegistrySetFactory)
				if err := _YGaugeRegistry.contract.UnpackLog(event, "SetFactory", log); err != nil {
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

// ParseSetFactory is a log parse operation binding the contract event 0x1c893ef9379093af30f458b9e74d2aba13c499660b68dec5e29af7b199c188b9.
//
// Solidity: event SetFactory(address factory)
func (_YGaugeRegistry *YGaugeRegistryFilterer) ParseSetFactory(log types.Log) (*YGaugeRegistrySetFactory, error) {
	event := new(YGaugeRegistrySetFactory)
	if err := _YGaugeRegistry.contract.UnpackLog(event, "SetFactory", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YGaugeRegistrySetManagementIterator is returned from FilterSetManagement and is used to iterate over the raw logs and unpacked data for SetManagement events raised by the YGaugeRegistry contract.
type YGaugeRegistrySetManagementIterator struct {
	Event *YGaugeRegistrySetManagement // Event containing the contract specifics and raw log

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
func (it *YGaugeRegistrySetManagementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YGaugeRegistrySetManagement)
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
		it.Event = new(YGaugeRegistrySetManagement)
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
func (it *YGaugeRegistrySetManagementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YGaugeRegistrySetManagementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YGaugeRegistrySetManagement represents a SetManagement event raised by the YGaugeRegistry contract.
type YGaugeRegistrySetManagement struct {
	Management common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetManagement is a free log retrieval operation binding the contract event 0xafe23f9e1f603b288748a507d5a993957e9f14313a5889d5a070851299939d59.
//
// Solidity: event SetManagement(address indexed management)
func (_YGaugeRegistry *YGaugeRegistryFilterer) FilterSetManagement(opts *bind.FilterOpts, management []common.Address) (*YGaugeRegistrySetManagementIterator, error) {

	var managementRule []interface{}
	for _, managementItem := range management {
		managementRule = append(managementRule, managementItem)
	}

	logs, sub, err := _YGaugeRegistry.contract.FilterLogs(opts, "SetManagement", managementRule)
	if err != nil {
		return nil, err
	}
	return &YGaugeRegistrySetManagementIterator{contract: _YGaugeRegistry.contract, event: "SetManagement", logs: logs, sub: sub}, nil
}

// WatchSetManagement is a free log subscription operation binding the contract event 0xafe23f9e1f603b288748a507d5a993957e9f14313a5889d5a070851299939d59.
//
// Solidity: event SetManagement(address indexed management)
func (_YGaugeRegistry *YGaugeRegistryFilterer) WatchSetManagement(opts *bind.WatchOpts, sink chan<- *YGaugeRegistrySetManagement, management []common.Address) (event.Subscription, error) {

	var managementRule []interface{}
	for _, managementItem := range management {
		managementRule = append(managementRule, managementItem)
	}

	logs, sub, err := _YGaugeRegistry.contract.WatchLogs(opts, "SetManagement", managementRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YGaugeRegistrySetManagement)
				if err := _YGaugeRegistry.contract.UnpackLog(event, "SetManagement", log); err != nil {
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

// ParseSetManagement is a log parse operation binding the contract event 0xafe23f9e1f603b288748a507d5a993957e9f14313a5889d5a070851299939d59.
//
// Solidity: event SetManagement(address indexed management)
func (_YGaugeRegistry *YGaugeRegistryFilterer) ParseSetManagement(log types.Log) (*YGaugeRegistrySetManagement, error) {
	event := new(YGaugeRegistrySetManagement)
	if err := _YGaugeRegistry.contract.UnpackLog(event, "SetManagement", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YGaugeRegistryUpdateIndexIterator is returned from FilterUpdateIndex and is used to iterate over the raw logs and unpacked data for UpdateIndex events raised by the YGaugeRegistry contract.
type YGaugeRegistryUpdateIndexIterator struct {
	Event *YGaugeRegistryUpdateIndex // Event containing the contract specifics and raw log

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
func (it *YGaugeRegistryUpdateIndexIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YGaugeRegistryUpdateIndex)
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
		it.Event = new(YGaugeRegistryUpdateIndex)
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
func (it *YGaugeRegistryUpdateIndexIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YGaugeRegistryUpdateIndexIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YGaugeRegistryUpdateIndex represents a UpdateIndex event raised by the YGaugeRegistry contract.
type YGaugeRegistryUpdateIndex struct {
	OldIdx *big.Int
	Idx    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUpdateIndex is a free log retrieval operation binding the contract event 0xc9d12a7659bfe5207e3439ef54f3218cb66f9fe3ec60af9da1f1c8878533cb46.
//
// Solidity: event UpdateIndex(uint256 indexed old_idx, uint256 idx)
func (_YGaugeRegistry *YGaugeRegistryFilterer) FilterUpdateIndex(opts *bind.FilterOpts, old_idx []*big.Int) (*YGaugeRegistryUpdateIndexIterator, error) {

	var old_idxRule []interface{}
	for _, old_idxItem := range old_idx {
		old_idxRule = append(old_idxRule, old_idxItem)
	}

	logs, sub, err := _YGaugeRegistry.contract.FilterLogs(opts, "UpdateIndex", old_idxRule)
	if err != nil {
		return nil, err
	}
	return &YGaugeRegistryUpdateIndexIterator{contract: _YGaugeRegistry.contract, event: "UpdateIndex", logs: logs, sub: sub}, nil
}

// WatchUpdateIndex is a free log subscription operation binding the contract event 0xc9d12a7659bfe5207e3439ef54f3218cb66f9fe3ec60af9da1f1c8878533cb46.
//
// Solidity: event UpdateIndex(uint256 indexed old_idx, uint256 idx)
func (_YGaugeRegistry *YGaugeRegistryFilterer) WatchUpdateIndex(opts *bind.WatchOpts, sink chan<- *YGaugeRegistryUpdateIndex, old_idx []*big.Int) (event.Subscription, error) {

	var old_idxRule []interface{}
	for _, old_idxItem := range old_idx {
		old_idxRule = append(old_idxRule, old_idxItem)
	}

	logs, sub, err := _YGaugeRegistry.contract.WatchLogs(opts, "UpdateIndex", old_idxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YGaugeRegistryUpdateIndex)
				if err := _YGaugeRegistry.contract.UnpackLog(event, "UpdateIndex", log); err != nil {
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

// ParseUpdateIndex is a log parse operation binding the contract event 0xc9d12a7659bfe5207e3439ef54f3218cb66f9fe3ec60af9da1f1c8878533cb46.
//
// Solidity: event UpdateIndex(uint256 indexed old_idx, uint256 idx)
func (_YGaugeRegistry *YGaugeRegistryFilterer) ParseUpdateIndex(log types.Log) (*YGaugeRegistryUpdateIndex, error) {
	event := new(YGaugeRegistryUpdateIndex)
	if err := _YGaugeRegistry.contract.UnpackLog(event, "UpdateIndex", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
