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

// YOptimismStakingRewardRegistryMetaData contains all meta data concerning the YOptimismStakingRewardRegistry contract.
var YOptimismStakingRewardRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"canEndorse\",\"type\":\"bool\"}],\"name\":\"ApprovedPoolEndorser\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governance\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovedPoolOwnerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakingPool\",\"type\":\"address\"}],\"name\":\"StakingPoolAdded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingPool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_replaceExistingPool\",\"type\":\"bool\"}],\"name\":\"addStakingPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"approvedPoolOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isStakingPoolEndorsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"poolEndorsers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovedPoolOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setPoolEndorsers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakingPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// YOptimismStakingRewardRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use YOptimismStakingRewardRegistryMetaData.ABI instead.
var YOptimismStakingRewardRegistryABI = YOptimismStakingRewardRegistryMetaData.ABI

// YOptimismStakingRewardRegistry is an auto generated Go binding around an Ethereum contract.
type YOptimismStakingRewardRegistry struct {
	YOptimismStakingRewardRegistryCaller     // Read-only binding to the contract
	YOptimismStakingRewardRegistryTransactor // Write-only binding to the contract
	YOptimismStakingRewardRegistryFilterer   // Log filterer for contract events
}

// YOptimismStakingRewardRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type YOptimismStakingRewardRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YOptimismStakingRewardRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YOptimismStakingRewardRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YOptimismStakingRewardRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YOptimismStakingRewardRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YOptimismStakingRewardRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YOptimismStakingRewardRegistrySession struct {
	Contract     *YOptimismStakingRewardRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                   // Call options to use throughout this session
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// YOptimismStakingRewardRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YOptimismStakingRewardRegistryCallerSession struct {
	Contract *YOptimismStakingRewardRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                         // Call options to use throughout this session
}

// YOptimismStakingRewardRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YOptimismStakingRewardRegistryTransactorSession struct {
	Contract     *YOptimismStakingRewardRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                         // Transaction auth options to use throughout this session
}

// YOptimismStakingRewardRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type YOptimismStakingRewardRegistryRaw struct {
	Contract *YOptimismStakingRewardRegistry // Generic contract binding to access the raw methods on
}

// YOptimismStakingRewardRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YOptimismStakingRewardRegistryCallerRaw struct {
	Contract *YOptimismStakingRewardRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// YOptimismStakingRewardRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YOptimismStakingRewardRegistryTransactorRaw struct {
	Contract *YOptimismStakingRewardRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYOptimismStakingRewardRegistry creates a new instance of YOptimismStakingRewardRegistry, bound to a specific deployed contract.
func NewYOptimismStakingRewardRegistry(address common.Address, backend bind.ContractBackend) (*YOptimismStakingRewardRegistry, error) {
	contract, err := bindYOptimismStakingRewardRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YOptimismStakingRewardRegistry{YOptimismStakingRewardRegistryCaller: YOptimismStakingRewardRegistryCaller{contract: contract}, YOptimismStakingRewardRegistryTransactor: YOptimismStakingRewardRegistryTransactor{contract: contract}, YOptimismStakingRewardRegistryFilterer: YOptimismStakingRewardRegistryFilterer{contract: contract}}, nil
}

// NewYOptimismStakingRewardRegistryCaller creates a new read-only instance of YOptimismStakingRewardRegistry, bound to a specific deployed contract.
func NewYOptimismStakingRewardRegistryCaller(address common.Address, caller bind.ContractCaller) (*YOptimismStakingRewardRegistryCaller, error) {
	contract, err := bindYOptimismStakingRewardRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YOptimismStakingRewardRegistryCaller{contract: contract}, nil
}

// NewYOptimismStakingRewardRegistryTransactor creates a new write-only instance of YOptimismStakingRewardRegistry, bound to a specific deployed contract.
func NewYOptimismStakingRewardRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*YOptimismStakingRewardRegistryTransactor, error) {
	contract, err := bindYOptimismStakingRewardRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YOptimismStakingRewardRegistryTransactor{contract: contract}, nil
}

// NewYOptimismStakingRewardRegistryFilterer creates a new log filterer instance of YOptimismStakingRewardRegistry, bound to a specific deployed contract.
func NewYOptimismStakingRewardRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*YOptimismStakingRewardRegistryFilterer, error) {
	contract, err := bindYOptimismStakingRewardRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YOptimismStakingRewardRegistryFilterer{contract: contract}, nil
}

// bindYOptimismStakingRewardRegistry binds a generic wrapper to an already deployed contract.
func bindYOptimismStakingRewardRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(YOptimismStakingRewardRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YOptimismStakingRewardRegistry.Contract.YOptimismStakingRewardRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YOptimismStakingRewardRegistry.Contract.YOptimismStakingRewardRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YOptimismStakingRewardRegistry.Contract.YOptimismStakingRewardRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YOptimismStakingRewardRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YOptimismStakingRewardRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YOptimismStakingRewardRegistry.Contract.contract.Transact(opts, method, params...)
}

// ApprovedPoolOwner is a free data retrieval call binding the contract method 0x4f2ac26e.
//
// Solidity: function approvedPoolOwner(address ) view returns(bool)
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistryCaller) ApprovedPoolOwner(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _YOptimismStakingRewardRegistry.contract.Call(opts, &out, "approvedPoolOwner", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ApprovedPoolOwner is a free data retrieval call binding the contract method 0x4f2ac26e.
//
// Solidity: function approvedPoolOwner(address ) view returns(bool)
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistrySession) ApprovedPoolOwner(arg0 common.Address) (bool, error) {
	return _YOptimismStakingRewardRegistry.Contract.ApprovedPoolOwner(&_YOptimismStakingRewardRegistry.CallOpts, arg0)
}

// ApprovedPoolOwner is a free data retrieval call binding the contract method 0x4f2ac26e.
//
// Solidity: function approvedPoolOwner(address ) view returns(bool)
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistryCallerSession) ApprovedPoolOwner(arg0 common.Address) (bool, error) {
	return _YOptimismStakingRewardRegistry.Contract.ApprovedPoolOwner(&_YOptimismStakingRewardRegistry.CallOpts, arg0)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address ) view returns(bool)
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistryCaller) IsRegistered(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _YOptimismStakingRewardRegistry.contract.Call(opts, &out, "isRegistered", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address ) view returns(bool)
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistrySession) IsRegistered(arg0 common.Address) (bool, error) {
	return _YOptimismStakingRewardRegistry.Contract.IsRegistered(&_YOptimismStakingRewardRegistry.CallOpts, arg0)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address ) view returns(bool)
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistryCallerSession) IsRegistered(arg0 common.Address) (bool, error) {
	return _YOptimismStakingRewardRegistry.Contract.IsRegistered(&_YOptimismStakingRewardRegistry.CallOpts, arg0)
}

// IsStakingPoolEndorsed is a free data retrieval call binding the contract method 0x89884973.
//
// Solidity: function isStakingPoolEndorsed(address ) view returns(bool)
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistryCaller) IsStakingPoolEndorsed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _YOptimismStakingRewardRegistry.contract.Call(opts, &out, "isStakingPoolEndorsed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStakingPoolEndorsed is a free data retrieval call binding the contract method 0x89884973.
//
// Solidity: function isStakingPoolEndorsed(address ) view returns(bool)
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistrySession) IsStakingPoolEndorsed(arg0 common.Address) (bool, error) {
	return _YOptimismStakingRewardRegistry.Contract.IsStakingPoolEndorsed(&_YOptimismStakingRewardRegistry.CallOpts, arg0)
}

// IsStakingPoolEndorsed is a free data retrieval call binding the contract method 0x89884973.
//
// Solidity: function isStakingPoolEndorsed(address ) view returns(bool)
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistryCallerSession) IsStakingPoolEndorsed(arg0 common.Address) (bool, error) {
	return _YOptimismStakingRewardRegistry.Contract.IsStakingPoolEndorsed(&_YOptimismStakingRewardRegistry.CallOpts, arg0)
}

// NumTokens is a free data retrieval call binding the contract method 0x8e499bcf.
//
// Solidity: function numTokens() view returns(uint256)
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistryCaller) NumTokens(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YOptimismStakingRewardRegistry.contract.Call(opts, &out, "numTokens")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumTokens is a free data retrieval call binding the contract method 0x8e499bcf.
//
// Solidity: function numTokens() view returns(uint256)
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistrySession) NumTokens() (*big.Int, error) {
	return _YOptimismStakingRewardRegistry.Contract.NumTokens(&_YOptimismStakingRewardRegistry.CallOpts)
}

// NumTokens is a free data retrieval call binding the contract method 0x8e499bcf.
//
// Solidity: function numTokens() view returns(uint256)
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistryCallerSession) NumTokens() (*big.Int, error) {
	return _YOptimismStakingRewardRegistry.Contract.NumTokens(&_YOptimismStakingRewardRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YOptimismStakingRewardRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistrySession) Owner() (common.Address, error) {
	return _YOptimismStakingRewardRegistry.Contract.Owner(&_YOptimismStakingRewardRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistryCallerSession) Owner() (common.Address, error) {
	return _YOptimismStakingRewardRegistry.Contract.Owner(&_YOptimismStakingRewardRegistry.CallOpts)
}

// PoolEndorsers is a free data retrieval call binding the contract method 0x4f1fde98.
//
// Solidity: function poolEndorsers(address ) view returns(bool)
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistryCaller) PoolEndorsers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _YOptimismStakingRewardRegistry.contract.Call(opts, &out, "poolEndorsers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PoolEndorsers is a free data retrieval call binding the contract method 0x4f1fde98.
//
// Solidity: function poolEndorsers(address ) view returns(bool)
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistrySession) PoolEndorsers(arg0 common.Address) (bool, error) {
	return _YOptimismStakingRewardRegistry.Contract.PoolEndorsers(&_YOptimismStakingRewardRegistry.CallOpts, arg0)
}

// PoolEndorsers is a free data retrieval call binding the contract method 0x4f1fde98.
//
// Solidity: function poolEndorsers(address ) view returns(bool)
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistryCallerSession) PoolEndorsers(arg0 common.Address) (bool, error) {
	return _YOptimismStakingRewardRegistry.Contract.PoolEndorsers(&_YOptimismStakingRewardRegistry.CallOpts, arg0)
}

// StakingPool is a free data retrieval call binding the contract method 0x8dddb3e5.
//
// Solidity: function stakingPool(address ) view returns(address)
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistryCaller) StakingPool(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _YOptimismStakingRewardRegistry.contract.Call(opts, &out, "stakingPool", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingPool is a free data retrieval call binding the contract method 0x8dddb3e5.
//
// Solidity: function stakingPool(address ) view returns(address)
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistrySession) StakingPool(arg0 common.Address) (common.Address, error) {
	return _YOptimismStakingRewardRegistry.Contract.StakingPool(&_YOptimismStakingRewardRegistry.CallOpts, arg0)
}

// StakingPool is a free data retrieval call binding the contract method 0x8dddb3e5.
//
// Solidity: function stakingPool(address ) view returns(address)
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistryCallerSession) StakingPool(arg0 common.Address) (common.Address, error) {
	return _YOptimismStakingRewardRegistry.Contract.StakingPool(&_YOptimismStakingRewardRegistry.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistryCaller) Tokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _YOptimismStakingRewardRegistry.contract.Call(opts, &out, "tokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistrySession) Tokens(arg0 *big.Int) (common.Address, error) {
	return _YOptimismStakingRewardRegistry.Contract.Tokens(&_YOptimismStakingRewardRegistry.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistryCallerSession) Tokens(arg0 *big.Int) (common.Address, error) {
	return _YOptimismStakingRewardRegistry.Contract.Tokens(&_YOptimismStakingRewardRegistry.CallOpts, arg0)
}

// AddStakingPool is a paid mutator transaction binding the contract method 0xaa9f079c.
//
// Solidity: function addStakingPool(address _stakingPool, address _token, bool _replaceExistingPool) returns()
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistryTransactor) AddStakingPool(opts *bind.TransactOpts, _stakingPool common.Address, _token common.Address, _replaceExistingPool bool) (*types.Transaction, error) {
	return _YOptimismStakingRewardRegistry.contract.Transact(opts, "addStakingPool", _stakingPool, _token, _replaceExistingPool)
}

// AddStakingPool is a paid mutator transaction binding the contract method 0xaa9f079c.
//
// Solidity: function addStakingPool(address _stakingPool, address _token, bool _replaceExistingPool) returns()
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistrySession) AddStakingPool(_stakingPool common.Address, _token common.Address, _replaceExistingPool bool) (*types.Transaction, error) {
	return _YOptimismStakingRewardRegistry.Contract.AddStakingPool(&_YOptimismStakingRewardRegistry.TransactOpts, _stakingPool, _token, _replaceExistingPool)
}

// AddStakingPool is a paid mutator transaction binding the contract method 0xaa9f079c.
//
// Solidity: function addStakingPool(address _stakingPool, address _token, bool _replaceExistingPool) returns()
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistryTransactorSession) AddStakingPool(_stakingPool common.Address, _token common.Address, _replaceExistingPool bool) (*types.Transaction, error) {
	return _YOptimismStakingRewardRegistry.Contract.AddStakingPool(&_YOptimismStakingRewardRegistry.TransactOpts, _stakingPool, _token, _replaceExistingPool)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YOptimismStakingRewardRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _YOptimismStakingRewardRegistry.Contract.RenounceOwnership(&_YOptimismStakingRewardRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _YOptimismStakingRewardRegistry.Contract.RenounceOwnership(&_YOptimismStakingRewardRegistry.TransactOpts)
}

// SetApprovedPoolOwner is a paid mutator transaction binding the contract method 0xaa9fd6ef.
//
// Solidity: function setApprovedPoolOwner(address _addr, bool _approved) returns()
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistryTransactor) SetApprovedPoolOwner(opts *bind.TransactOpts, _addr common.Address, _approved bool) (*types.Transaction, error) {
	return _YOptimismStakingRewardRegistry.contract.Transact(opts, "setApprovedPoolOwner", _addr, _approved)
}

// SetApprovedPoolOwner is a paid mutator transaction binding the contract method 0xaa9fd6ef.
//
// Solidity: function setApprovedPoolOwner(address _addr, bool _approved) returns()
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistrySession) SetApprovedPoolOwner(_addr common.Address, _approved bool) (*types.Transaction, error) {
	return _YOptimismStakingRewardRegistry.Contract.SetApprovedPoolOwner(&_YOptimismStakingRewardRegistry.TransactOpts, _addr, _approved)
}

// SetApprovedPoolOwner is a paid mutator transaction binding the contract method 0xaa9fd6ef.
//
// Solidity: function setApprovedPoolOwner(address _addr, bool _approved) returns()
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistryTransactorSession) SetApprovedPoolOwner(_addr common.Address, _approved bool) (*types.Transaction, error) {
	return _YOptimismStakingRewardRegistry.Contract.SetApprovedPoolOwner(&_YOptimismStakingRewardRegistry.TransactOpts, _addr, _approved)
}

// SetPoolEndorsers is a paid mutator transaction binding the contract method 0x5ea38625.
//
// Solidity: function setPoolEndorsers(address _addr, bool _approved) returns()
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistryTransactor) SetPoolEndorsers(opts *bind.TransactOpts, _addr common.Address, _approved bool) (*types.Transaction, error) {
	return _YOptimismStakingRewardRegistry.contract.Transact(opts, "setPoolEndorsers", _addr, _approved)
}

// SetPoolEndorsers is a paid mutator transaction binding the contract method 0x5ea38625.
//
// Solidity: function setPoolEndorsers(address _addr, bool _approved) returns()
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistrySession) SetPoolEndorsers(_addr common.Address, _approved bool) (*types.Transaction, error) {
	return _YOptimismStakingRewardRegistry.Contract.SetPoolEndorsers(&_YOptimismStakingRewardRegistry.TransactOpts, _addr, _approved)
}

// SetPoolEndorsers is a paid mutator transaction binding the contract method 0x5ea38625.
//
// Solidity: function setPoolEndorsers(address _addr, bool _approved) returns()
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistryTransactorSession) SetPoolEndorsers(_addr common.Address, _approved bool) (*types.Transaction, error) {
	return _YOptimismStakingRewardRegistry.Contract.SetPoolEndorsers(&_YOptimismStakingRewardRegistry.TransactOpts, _addr, _approved)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _YOptimismStakingRewardRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _YOptimismStakingRewardRegistry.Contract.TransferOwnership(&_YOptimismStakingRewardRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _YOptimismStakingRewardRegistry.Contract.TransferOwnership(&_YOptimismStakingRewardRegistry.TransactOpts, newOwner)
}

// YOptimismStakingRewardRegistryApprovedPoolEndorserIterator is returned from FilterApprovedPoolEndorser and is used to iterate over the raw logs and unpacked data for ApprovedPoolEndorser events raised by the YOptimismStakingRewardRegistry contract.
type YOptimismStakingRewardRegistryApprovedPoolEndorserIterator struct {
	Event *YOptimismStakingRewardRegistryApprovedPoolEndorser // Event containing the contract specifics and raw log

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
func (it *YOptimismStakingRewardRegistryApprovedPoolEndorserIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YOptimismStakingRewardRegistryApprovedPoolEndorser)
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
		it.Event = new(YOptimismStakingRewardRegistryApprovedPoolEndorser)
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
func (it *YOptimismStakingRewardRegistryApprovedPoolEndorserIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YOptimismStakingRewardRegistryApprovedPoolEndorserIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YOptimismStakingRewardRegistryApprovedPoolEndorser represents a ApprovedPoolEndorser event raised by the YOptimismStakingRewardRegistry contract.
type YOptimismStakingRewardRegistryApprovedPoolEndorser struct {
	Account    common.Address
	CanEndorse bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterApprovedPoolEndorser is a free log retrieval operation binding the contract event 0x971fecdddf4686fb97d914426cefb86d567c8e5081e48f0346d9ea0c931cac69.
//
// Solidity: event ApprovedPoolEndorser(address account, bool canEndorse)
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistryFilterer) FilterApprovedPoolEndorser(opts *bind.FilterOpts) (*YOptimismStakingRewardRegistryApprovedPoolEndorserIterator, error) {

	logs, sub, err := _YOptimismStakingRewardRegistry.contract.FilterLogs(opts, "ApprovedPoolEndorser")
	if err != nil {
		return nil, err
	}
	return &YOptimismStakingRewardRegistryApprovedPoolEndorserIterator{contract: _YOptimismStakingRewardRegistry.contract, event: "ApprovedPoolEndorser", logs: logs, sub: sub}, nil
}

// WatchApprovedPoolEndorser is a free log subscription operation binding the contract event 0x971fecdddf4686fb97d914426cefb86d567c8e5081e48f0346d9ea0c931cac69.
//
// Solidity: event ApprovedPoolEndorser(address account, bool canEndorse)
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistryFilterer) WatchApprovedPoolEndorser(opts *bind.WatchOpts, sink chan<- *YOptimismStakingRewardRegistryApprovedPoolEndorser) (event.Subscription, error) {

	logs, sub, err := _YOptimismStakingRewardRegistry.contract.WatchLogs(opts, "ApprovedPoolEndorser")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YOptimismStakingRewardRegistryApprovedPoolEndorser)
				if err := _YOptimismStakingRewardRegistry.contract.UnpackLog(event, "ApprovedPoolEndorser", log); err != nil {
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

// ParseApprovedPoolEndorser is a log parse operation binding the contract event 0x971fecdddf4686fb97d914426cefb86d567c8e5081e48f0346d9ea0c931cac69.
//
// Solidity: event ApprovedPoolEndorser(address account, bool canEndorse)
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistryFilterer) ParseApprovedPoolEndorser(log types.Log) (*YOptimismStakingRewardRegistryApprovedPoolEndorser, error) {
	event := new(YOptimismStakingRewardRegistryApprovedPoolEndorser)
	if err := _YOptimismStakingRewardRegistry.contract.UnpackLog(event, "ApprovedPoolEndorser", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YOptimismStakingRewardRegistryApprovedPoolOwnerUpdatedIterator is returned from FilterApprovedPoolOwnerUpdated and is used to iterate over the raw logs and unpacked data for ApprovedPoolOwnerUpdated events raised by the YOptimismStakingRewardRegistry contract.
type YOptimismStakingRewardRegistryApprovedPoolOwnerUpdatedIterator struct {
	Event *YOptimismStakingRewardRegistryApprovedPoolOwnerUpdated // Event containing the contract specifics and raw log

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
func (it *YOptimismStakingRewardRegistryApprovedPoolOwnerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YOptimismStakingRewardRegistryApprovedPoolOwnerUpdated)
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
		it.Event = new(YOptimismStakingRewardRegistryApprovedPoolOwnerUpdated)
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
func (it *YOptimismStakingRewardRegistryApprovedPoolOwnerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YOptimismStakingRewardRegistryApprovedPoolOwnerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YOptimismStakingRewardRegistryApprovedPoolOwnerUpdated represents a ApprovedPoolOwnerUpdated event raised by the YOptimismStakingRewardRegistry contract.
type YOptimismStakingRewardRegistryApprovedPoolOwnerUpdated struct {
	Governance common.Address
	Approved   bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterApprovedPoolOwnerUpdated is a free log retrieval operation binding the contract event 0xb3b57e8caeb8200bd5bdd2e4c7a421b0f9683a9e9a376d8a9c6c29883f7a7292.
//
// Solidity: event ApprovedPoolOwnerUpdated(address governance, bool approved)
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistryFilterer) FilterApprovedPoolOwnerUpdated(opts *bind.FilterOpts) (*YOptimismStakingRewardRegistryApprovedPoolOwnerUpdatedIterator, error) {

	logs, sub, err := _YOptimismStakingRewardRegistry.contract.FilterLogs(opts, "ApprovedPoolOwnerUpdated")
	if err != nil {
		return nil, err
	}
	return &YOptimismStakingRewardRegistryApprovedPoolOwnerUpdatedIterator{contract: _YOptimismStakingRewardRegistry.contract, event: "ApprovedPoolOwnerUpdated", logs: logs, sub: sub}, nil
}

// WatchApprovedPoolOwnerUpdated is a free log subscription operation binding the contract event 0xb3b57e8caeb8200bd5bdd2e4c7a421b0f9683a9e9a376d8a9c6c29883f7a7292.
//
// Solidity: event ApprovedPoolOwnerUpdated(address governance, bool approved)
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistryFilterer) WatchApprovedPoolOwnerUpdated(opts *bind.WatchOpts, sink chan<- *YOptimismStakingRewardRegistryApprovedPoolOwnerUpdated) (event.Subscription, error) {

	logs, sub, err := _YOptimismStakingRewardRegistry.contract.WatchLogs(opts, "ApprovedPoolOwnerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YOptimismStakingRewardRegistryApprovedPoolOwnerUpdated)
				if err := _YOptimismStakingRewardRegistry.contract.UnpackLog(event, "ApprovedPoolOwnerUpdated", log); err != nil {
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

// ParseApprovedPoolOwnerUpdated is a log parse operation binding the contract event 0xb3b57e8caeb8200bd5bdd2e4c7a421b0f9683a9e9a376d8a9c6c29883f7a7292.
//
// Solidity: event ApprovedPoolOwnerUpdated(address governance, bool approved)
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistryFilterer) ParseApprovedPoolOwnerUpdated(log types.Log) (*YOptimismStakingRewardRegistryApprovedPoolOwnerUpdated, error) {
	event := new(YOptimismStakingRewardRegistryApprovedPoolOwnerUpdated)
	if err := _YOptimismStakingRewardRegistry.contract.UnpackLog(event, "ApprovedPoolOwnerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YOptimismStakingRewardRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the YOptimismStakingRewardRegistry contract.
type YOptimismStakingRewardRegistryOwnershipTransferredIterator struct {
	Event *YOptimismStakingRewardRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *YOptimismStakingRewardRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YOptimismStakingRewardRegistryOwnershipTransferred)
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
		it.Event = new(YOptimismStakingRewardRegistryOwnershipTransferred)
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
func (it *YOptimismStakingRewardRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YOptimismStakingRewardRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YOptimismStakingRewardRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the YOptimismStakingRewardRegistry contract.
type YOptimismStakingRewardRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*YOptimismStakingRewardRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _YOptimismStakingRewardRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &YOptimismStakingRewardRegistryOwnershipTransferredIterator{contract: _YOptimismStakingRewardRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *YOptimismStakingRewardRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _YOptimismStakingRewardRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YOptimismStakingRewardRegistryOwnershipTransferred)
				if err := _YOptimismStakingRewardRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*YOptimismStakingRewardRegistryOwnershipTransferred, error) {
	event := new(YOptimismStakingRewardRegistryOwnershipTransferred)
	if err := _YOptimismStakingRewardRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YOptimismStakingRewardRegistryStakingPoolAddedIterator is returned from FilterStakingPoolAdded and is used to iterate over the raw logs and unpacked data for StakingPoolAdded events raised by the YOptimismStakingRewardRegistry contract.
type YOptimismStakingRewardRegistryStakingPoolAddedIterator struct {
	Event *YOptimismStakingRewardRegistryStakingPoolAdded // Event containing the contract specifics and raw log

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
func (it *YOptimismStakingRewardRegistryStakingPoolAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YOptimismStakingRewardRegistryStakingPoolAdded)
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
		it.Event = new(YOptimismStakingRewardRegistryStakingPoolAdded)
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
func (it *YOptimismStakingRewardRegistryStakingPoolAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YOptimismStakingRewardRegistryStakingPoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YOptimismStakingRewardRegistryStakingPoolAdded represents a StakingPoolAdded event raised by the YOptimismStakingRewardRegistry contract.
type YOptimismStakingRewardRegistryStakingPoolAdded struct {
	Token       common.Address
	StakingPool common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStakingPoolAdded is a free log retrieval operation binding the contract event 0xd58b22ec3b77fb836c6ae1bba270411b0fa4961ff8423444de63e7f804826f74.
//
// Solidity: event StakingPoolAdded(address indexed token, address indexed stakingPool)
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistryFilterer) FilterStakingPoolAdded(opts *bind.FilterOpts, token []common.Address, stakingPool []common.Address) (*YOptimismStakingRewardRegistryStakingPoolAddedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var stakingPoolRule []interface{}
	for _, stakingPoolItem := range stakingPool {
		stakingPoolRule = append(stakingPoolRule, stakingPoolItem)
	}

	logs, sub, err := _YOptimismStakingRewardRegistry.contract.FilterLogs(opts, "StakingPoolAdded", tokenRule, stakingPoolRule)
	if err != nil {
		return nil, err
	}
	return &YOptimismStakingRewardRegistryStakingPoolAddedIterator{contract: _YOptimismStakingRewardRegistry.contract, event: "StakingPoolAdded", logs: logs, sub: sub}, nil
}

// WatchStakingPoolAdded is a free log subscription operation binding the contract event 0xd58b22ec3b77fb836c6ae1bba270411b0fa4961ff8423444de63e7f804826f74.
//
// Solidity: event StakingPoolAdded(address indexed token, address indexed stakingPool)
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistryFilterer) WatchStakingPoolAdded(opts *bind.WatchOpts, sink chan<- *YOptimismStakingRewardRegistryStakingPoolAdded, token []common.Address, stakingPool []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var stakingPoolRule []interface{}
	for _, stakingPoolItem := range stakingPool {
		stakingPoolRule = append(stakingPoolRule, stakingPoolItem)
	}

	logs, sub, err := _YOptimismStakingRewardRegistry.contract.WatchLogs(opts, "StakingPoolAdded", tokenRule, stakingPoolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YOptimismStakingRewardRegistryStakingPoolAdded)
				if err := _YOptimismStakingRewardRegistry.contract.UnpackLog(event, "StakingPoolAdded", log); err != nil {
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

// ParseStakingPoolAdded is a log parse operation binding the contract event 0xd58b22ec3b77fb836c6ae1bba270411b0fa4961ff8423444de63e7f804826f74.
//
// Solidity: event StakingPoolAdded(address indexed token, address indexed stakingPool)
func (_YOptimismStakingRewardRegistry *YOptimismStakingRewardRegistryFilterer) ParseStakingPoolAdded(log types.Log) (*YOptimismStakingRewardRegistryStakingPoolAdded, error) {
	event := new(YOptimismStakingRewardRegistryStakingPoolAdded)
	if err := _YOptimismStakingRewardRegistry.contract.UnpackLog(event, "StakingPoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
