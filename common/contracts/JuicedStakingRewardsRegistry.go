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

// JuicedStakingRewardsRegistryMetaData contains all meta data concerning the JuicedStakingRewardsRegistry contract.
var JuicedStakingRewardsRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"canEndorse\",\"type\":\"bool\"}],\"name\":\"ApprovedPoolEndorser\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governance\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovedPoolOwnerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"stakingContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"zapContract\",\"type\":\"address\"}],\"name\":\"DefaultContractsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"stakingPool\",\"type\":\"address\"}],\"name\":\"StakingPoolAdded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingPool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_replaceExistingPool\",\"type\":\"bool\"}],\"name\":\"addStakingPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"approvedPoolOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingToken\",\"type\":\"address\"}],\"name\":\"cloneAndAddStakingPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newStakingPool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isStakingPoolEndorsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"poolEndorsers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"replacedStakingPools\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovedPoolOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingPool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_zapContract\",\"type\":\"address\"}],\"name\":\"setDefaultContracts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setPoolEndorsers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakingPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zapContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// JuicedStakingRewardsRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use JuicedStakingRewardsRegistryMetaData.ABI instead.
var JuicedStakingRewardsRegistryABI = JuicedStakingRewardsRegistryMetaData.ABI

// JuicedStakingRewardsRegistry is an auto generated Go binding around an Ethereum contract.
type JuicedStakingRewardsRegistry struct {
	JuicedStakingRewardsRegistryCaller     // Read-only binding to the contract
	JuicedStakingRewardsRegistryTransactor // Write-only binding to the contract
	JuicedStakingRewardsRegistryFilterer   // Log filterer for contract events
}

// JuicedStakingRewardsRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type JuicedStakingRewardsRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JuicedStakingRewardsRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JuicedStakingRewardsRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JuicedStakingRewardsRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JuicedStakingRewardsRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JuicedStakingRewardsRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JuicedStakingRewardsRegistrySession struct {
	Contract     *JuicedStakingRewardsRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// JuicedStakingRewardsRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JuicedStakingRewardsRegistryCallerSession struct {
	Contract *JuicedStakingRewardsRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// JuicedStakingRewardsRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JuicedStakingRewardsRegistryTransactorSession struct {
	Contract     *JuicedStakingRewardsRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// JuicedStakingRewardsRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type JuicedStakingRewardsRegistryRaw struct {
	Contract *JuicedStakingRewardsRegistry // Generic contract binding to access the raw methods on
}

// JuicedStakingRewardsRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JuicedStakingRewardsRegistryCallerRaw struct {
	Contract *JuicedStakingRewardsRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// JuicedStakingRewardsRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JuicedStakingRewardsRegistryTransactorRaw struct {
	Contract *JuicedStakingRewardsRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJuicedStakingRewardsRegistry creates a new instance of JuicedStakingRewardsRegistry, bound to a specific deployed contract.
func NewJuicedStakingRewardsRegistry(address common.Address, backend bind.ContractBackend) (*JuicedStakingRewardsRegistry, error) {
	contract, err := bindJuicedStakingRewardsRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &JuicedStakingRewardsRegistry{JuicedStakingRewardsRegistryCaller: JuicedStakingRewardsRegistryCaller{contract: contract}, JuicedStakingRewardsRegistryTransactor: JuicedStakingRewardsRegistryTransactor{contract: contract}, JuicedStakingRewardsRegistryFilterer: JuicedStakingRewardsRegistryFilterer{contract: contract}}, nil
}

// NewJuicedStakingRewardsRegistryCaller creates a new read-only instance of JuicedStakingRewardsRegistry, bound to a specific deployed contract.
func NewJuicedStakingRewardsRegistryCaller(address common.Address, caller bind.ContractCaller) (*JuicedStakingRewardsRegistryCaller, error) {
	contract, err := bindJuicedStakingRewardsRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JuicedStakingRewardsRegistryCaller{contract: contract}, nil
}

// NewJuicedStakingRewardsRegistryTransactor creates a new write-only instance of JuicedStakingRewardsRegistry, bound to a specific deployed contract.
func NewJuicedStakingRewardsRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*JuicedStakingRewardsRegistryTransactor, error) {
	contract, err := bindJuicedStakingRewardsRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JuicedStakingRewardsRegistryTransactor{contract: contract}, nil
}

// NewJuicedStakingRewardsRegistryFilterer creates a new log filterer instance of JuicedStakingRewardsRegistry, bound to a specific deployed contract.
func NewJuicedStakingRewardsRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*JuicedStakingRewardsRegistryFilterer, error) {
	contract, err := bindJuicedStakingRewardsRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JuicedStakingRewardsRegistryFilterer{contract: contract}, nil
}

// bindJuicedStakingRewardsRegistry binds a generic wrapper to an already deployed contract.
func bindJuicedStakingRewardsRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := JuicedStakingRewardsRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JuicedStakingRewardsRegistry.Contract.JuicedStakingRewardsRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JuicedStakingRewardsRegistry.Contract.JuicedStakingRewardsRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JuicedStakingRewardsRegistry.Contract.JuicedStakingRewardsRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JuicedStakingRewardsRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JuicedStakingRewardsRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JuicedStakingRewardsRegistry.Contract.contract.Transact(opts, method, params...)
}

// ApprovedPoolOwner is a free data retrieval call binding the contract method 0x4f2ac26e.
//
// Solidity: function approvedPoolOwner(address ) view returns(bool)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryCaller) ApprovedPoolOwner(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _JuicedStakingRewardsRegistry.contract.Call(opts, &out, "approvedPoolOwner", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ApprovedPoolOwner is a free data retrieval call binding the contract method 0x4f2ac26e.
//
// Solidity: function approvedPoolOwner(address ) view returns(bool)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistrySession) ApprovedPoolOwner(arg0 common.Address) (bool, error) {
	return _JuicedStakingRewardsRegistry.Contract.ApprovedPoolOwner(&_JuicedStakingRewardsRegistry.CallOpts, arg0)
}

// ApprovedPoolOwner is a free data retrieval call binding the contract method 0x4f2ac26e.
//
// Solidity: function approvedPoolOwner(address ) view returns(bool)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryCallerSession) ApprovedPoolOwner(arg0 common.Address) (bool, error) {
	return _JuicedStakingRewardsRegistry.Contract.ApprovedPoolOwner(&_JuicedStakingRewardsRegistry.CallOpts, arg0)
}

// IsStakingPoolEndorsed is a free data retrieval call binding the contract method 0x89884973.
//
// Solidity: function isStakingPoolEndorsed(address ) view returns(bool)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryCaller) IsStakingPoolEndorsed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _JuicedStakingRewardsRegistry.contract.Call(opts, &out, "isStakingPoolEndorsed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStakingPoolEndorsed is a free data retrieval call binding the contract method 0x89884973.
//
// Solidity: function isStakingPoolEndorsed(address ) view returns(bool)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistrySession) IsStakingPoolEndorsed(arg0 common.Address) (bool, error) {
	return _JuicedStakingRewardsRegistry.Contract.IsStakingPoolEndorsed(&_JuicedStakingRewardsRegistry.CallOpts, arg0)
}

// IsStakingPoolEndorsed is a free data retrieval call binding the contract method 0x89884973.
//
// Solidity: function isStakingPoolEndorsed(address ) view returns(bool)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryCallerSession) IsStakingPoolEndorsed(arg0 common.Address) (bool, error) {
	return _JuicedStakingRewardsRegistry.Contract.IsStakingPoolEndorsed(&_JuicedStakingRewardsRegistry.CallOpts, arg0)
}

// NumTokens is a free data retrieval call binding the contract method 0x8e499bcf.
//
// Solidity: function numTokens() view returns(uint256)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryCaller) NumTokens(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _JuicedStakingRewardsRegistry.contract.Call(opts, &out, "numTokens")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumTokens is a free data retrieval call binding the contract method 0x8e499bcf.
//
// Solidity: function numTokens() view returns(uint256)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistrySession) NumTokens() (*big.Int, error) {
	return _JuicedStakingRewardsRegistry.Contract.NumTokens(&_JuicedStakingRewardsRegistry.CallOpts)
}

// NumTokens is a free data retrieval call binding the contract method 0x8e499bcf.
//
// Solidity: function numTokens() view returns(uint256)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryCallerSession) NumTokens() (*big.Int, error) {
	return _JuicedStakingRewardsRegistry.Contract.NumTokens(&_JuicedStakingRewardsRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JuicedStakingRewardsRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistrySession) Owner() (common.Address, error) {
	return _JuicedStakingRewardsRegistry.Contract.Owner(&_JuicedStakingRewardsRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryCallerSession) Owner() (common.Address, error) {
	return _JuicedStakingRewardsRegistry.Contract.Owner(&_JuicedStakingRewardsRegistry.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JuicedStakingRewardsRegistry.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistrySession) PendingOwner() (common.Address, error) {
	return _JuicedStakingRewardsRegistry.Contract.PendingOwner(&_JuicedStakingRewardsRegistry.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryCallerSession) PendingOwner() (common.Address, error) {
	return _JuicedStakingRewardsRegistry.Contract.PendingOwner(&_JuicedStakingRewardsRegistry.CallOpts)
}

// PoolEndorsers is a free data retrieval call binding the contract method 0x4f1fde98.
//
// Solidity: function poolEndorsers(address ) view returns(bool)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryCaller) PoolEndorsers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _JuicedStakingRewardsRegistry.contract.Call(opts, &out, "poolEndorsers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PoolEndorsers is a free data retrieval call binding the contract method 0x4f1fde98.
//
// Solidity: function poolEndorsers(address ) view returns(bool)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistrySession) PoolEndorsers(arg0 common.Address) (bool, error) {
	return _JuicedStakingRewardsRegistry.Contract.PoolEndorsers(&_JuicedStakingRewardsRegistry.CallOpts, arg0)
}

// PoolEndorsers is a free data retrieval call binding the contract method 0x4f1fde98.
//
// Solidity: function poolEndorsers(address ) view returns(bool)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryCallerSession) PoolEndorsers(arg0 common.Address) (bool, error) {
	return _JuicedStakingRewardsRegistry.Contract.PoolEndorsers(&_JuicedStakingRewardsRegistry.CallOpts, arg0)
}

// ReplacedStakingPools is a free data retrieval call binding the contract method 0x403a3f4d.
//
// Solidity: function replacedStakingPools(uint256 ) view returns(address)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryCaller) ReplacedStakingPools(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _JuicedStakingRewardsRegistry.contract.Call(opts, &out, "replacedStakingPools", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReplacedStakingPools is a free data retrieval call binding the contract method 0x403a3f4d.
//
// Solidity: function replacedStakingPools(uint256 ) view returns(address)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistrySession) ReplacedStakingPools(arg0 *big.Int) (common.Address, error) {
	return _JuicedStakingRewardsRegistry.Contract.ReplacedStakingPools(&_JuicedStakingRewardsRegistry.CallOpts, arg0)
}

// ReplacedStakingPools is a free data retrieval call binding the contract method 0x403a3f4d.
//
// Solidity: function replacedStakingPools(uint256 ) view returns(address)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryCallerSession) ReplacedStakingPools(arg0 *big.Int) (common.Address, error) {
	return _JuicedStakingRewardsRegistry.Contract.ReplacedStakingPools(&_JuicedStakingRewardsRegistry.CallOpts, arg0)
}

// StakingContract is a free data retrieval call binding the contract method 0xee99205c.
//
// Solidity: function stakingContract() view returns(address)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryCaller) StakingContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JuicedStakingRewardsRegistry.contract.Call(opts, &out, "stakingContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingContract is a free data retrieval call binding the contract method 0xee99205c.
//
// Solidity: function stakingContract() view returns(address)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistrySession) StakingContract() (common.Address, error) {
	return _JuicedStakingRewardsRegistry.Contract.StakingContract(&_JuicedStakingRewardsRegistry.CallOpts)
}

// StakingContract is a free data retrieval call binding the contract method 0xee99205c.
//
// Solidity: function stakingContract() view returns(address)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryCallerSession) StakingContract() (common.Address, error) {
	return _JuicedStakingRewardsRegistry.Contract.StakingContract(&_JuicedStakingRewardsRegistry.CallOpts)
}

// StakingPool is a free data retrieval call binding the contract method 0x8dddb3e5.
//
// Solidity: function stakingPool(address ) view returns(address)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryCaller) StakingPool(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _JuicedStakingRewardsRegistry.contract.Call(opts, &out, "stakingPool", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingPool is a free data retrieval call binding the contract method 0x8dddb3e5.
//
// Solidity: function stakingPool(address ) view returns(address)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistrySession) StakingPool(arg0 common.Address) (common.Address, error) {
	return _JuicedStakingRewardsRegistry.Contract.StakingPool(&_JuicedStakingRewardsRegistry.CallOpts, arg0)
}

// StakingPool is a free data retrieval call binding the contract method 0x8dddb3e5.
//
// Solidity: function stakingPool(address ) view returns(address)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryCallerSession) StakingPool(arg0 common.Address) (common.Address, error) {
	return _JuicedStakingRewardsRegistry.Contract.StakingPool(&_JuicedStakingRewardsRegistry.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryCaller) Tokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _JuicedStakingRewardsRegistry.contract.Call(opts, &out, "tokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistrySession) Tokens(arg0 *big.Int) (common.Address, error) {
	return _JuicedStakingRewardsRegistry.Contract.Tokens(&_JuicedStakingRewardsRegistry.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryCallerSession) Tokens(arg0 *big.Int) (common.Address, error) {
	return _JuicedStakingRewardsRegistry.Contract.Tokens(&_JuicedStakingRewardsRegistry.CallOpts, arg0)
}

// ZapContract is a free data retrieval call binding the contract method 0x2549dad9.
//
// Solidity: function zapContract() view returns(address)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryCaller) ZapContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JuicedStakingRewardsRegistry.contract.Call(opts, &out, "zapContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZapContract is a free data retrieval call binding the contract method 0x2549dad9.
//
// Solidity: function zapContract() view returns(address)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistrySession) ZapContract() (common.Address, error) {
	return _JuicedStakingRewardsRegistry.Contract.ZapContract(&_JuicedStakingRewardsRegistry.CallOpts)
}

// ZapContract is a free data retrieval call binding the contract method 0x2549dad9.
//
// Solidity: function zapContract() view returns(address)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryCallerSession) ZapContract() (common.Address, error) {
	return _JuicedStakingRewardsRegistry.Contract.ZapContract(&_JuicedStakingRewardsRegistry.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JuicedStakingRewardsRegistry.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistrySession) AcceptOwnership() (*types.Transaction, error) {
	return _JuicedStakingRewardsRegistry.Contract.AcceptOwnership(&_JuicedStakingRewardsRegistry.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _JuicedStakingRewardsRegistry.Contract.AcceptOwnership(&_JuicedStakingRewardsRegistry.TransactOpts)
}

// AddStakingPool is a paid mutator transaction binding the contract method 0xaa9f079c.
//
// Solidity: function addStakingPool(address _stakingPool, address _token, bool _replaceExistingPool) returns()
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryTransactor) AddStakingPool(opts *bind.TransactOpts, _stakingPool common.Address, _token common.Address, _replaceExistingPool bool) (*types.Transaction, error) {
	return _JuicedStakingRewardsRegistry.contract.Transact(opts, "addStakingPool", _stakingPool, _token, _replaceExistingPool)
}

// AddStakingPool is a paid mutator transaction binding the contract method 0xaa9f079c.
//
// Solidity: function addStakingPool(address _stakingPool, address _token, bool _replaceExistingPool) returns()
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistrySession) AddStakingPool(_stakingPool common.Address, _token common.Address, _replaceExistingPool bool) (*types.Transaction, error) {
	return _JuicedStakingRewardsRegistry.Contract.AddStakingPool(&_JuicedStakingRewardsRegistry.TransactOpts, _stakingPool, _token, _replaceExistingPool)
}

// AddStakingPool is a paid mutator transaction binding the contract method 0xaa9f079c.
//
// Solidity: function addStakingPool(address _stakingPool, address _token, bool _replaceExistingPool) returns()
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryTransactorSession) AddStakingPool(_stakingPool common.Address, _token common.Address, _replaceExistingPool bool) (*types.Transaction, error) {
	return _JuicedStakingRewardsRegistry.Contract.AddStakingPool(&_JuicedStakingRewardsRegistry.TransactOpts, _stakingPool, _token, _replaceExistingPool)
}

// CloneAndAddStakingPool is a paid mutator transaction binding the contract method 0xf58422ea.
//
// Solidity: function cloneAndAddStakingPool(address _stakingToken) returns(address newStakingPool)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryTransactor) CloneAndAddStakingPool(opts *bind.TransactOpts, _stakingToken common.Address) (*types.Transaction, error) {
	return _JuicedStakingRewardsRegistry.contract.Transact(opts, "cloneAndAddStakingPool", _stakingToken)
}

// CloneAndAddStakingPool is a paid mutator transaction binding the contract method 0xf58422ea.
//
// Solidity: function cloneAndAddStakingPool(address _stakingToken) returns(address newStakingPool)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistrySession) CloneAndAddStakingPool(_stakingToken common.Address) (*types.Transaction, error) {
	return _JuicedStakingRewardsRegistry.Contract.CloneAndAddStakingPool(&_JuicedStakingRewardsRegistry.TransactOpts, _stakingToken)
}

// CloneAndAddStakingPool is a paid mutator transaction binding the contract method 0xf58422ea.
//
// Solidity: function cloneAndAddStakingPool(address _stakingToken) returns(address newStakingPool)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryTransactorSession) CloneAndAddStakingPool(_stakingToken common.Address) (*types.Transaction, error) {
	return _JuicedStakingRewardsRegistry.Contract.CloneAndAddStakingPool(&_JuicedStakingRewardsRegistry.TransactOpts, _stakingToken)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JuicedStakingRewardsRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _JuicedStakingRewardsRegistry.Contract.RenounceOwnership(&_JuicedStakingRewardsRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _JuicedStakingRewardsRegistry.Contract.RenounceOwnership(&_JuicedStakingRewardsRegistry.TransactOpts)
}

// SetApprovedPoolOwner is a paid mutator transaction binding the contract method 0xaa9fd6ef.
//
// Solidity: function setApprovedPoolOwner(address _addr, bool _approved) returns()
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryTransactor) SetApprovedPoolOwner(opts *bind.TransactOpts, _addr common.Address, _approved bool) (*types.Transaction, error) {
	return _JuicedStakingRewardsRegistry.contract.Transact(opts, "setApprovedPoolOwner", _addr, _approved)
}

// SetApprovedPoolOwner is a paid mutator transaction binding the contract method 0xaa9fd6ef.
//
// Solidity: function setApprovedPoolOwner(address _addr, bool _approved) returns()
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistrySession) SetApprovedPoolOwner(_addr common.Address, _approved bool) (*types.Transaction, error) {
	return _JuicedStakingRewardsRegistry.Contract.SetApprovedPoolOwner(&_JuicedStakingRewardsRegistry.TransactOpts, _addr, _approved)
}

// SetApprovedPoolOwner is a paid mutator transaction binding the contract method 0xaa9fd6ef.
//
// Solidity: function setApprovedPoolOwner(address _addr, bool _approved) returns()
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryTransactorSession) SetApprovedPoolOwner(_addr common.Address, _approved bool) (*types.Transaction, error) {
	return _JuicedStakingRewardsRegistry.Contract.SetApprovedPoolOwner(&_JuicedStakingRewardsRegistry.TransactOpts, _addr, _approved)
}

// SetDefaultContracts is a paid mutator transaction binding the contract method 0x0c689126.
//
// Solidity: function setDefaultContracts(address _stakingPool, address _zapContract) returns()
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryTransactor) SetDefaultContracts(opts *bind.TransactOpts, _stakingPool common.Address, _zapContract common.Address) (*types.Transaction, error) {
	return _JuicedStakingRewardsRegistry.contract.Transact(opts, "setDefaultContracts", _stakingPool, _zapContract)
}

// SetDefaultContracts is a paid mutator transaction binding the contract method 0x0c689126.
//
// Solidity: function setDefaultContracts(address _stakingPool, address _zapContract) returns()
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistrySession) SetDefaultContracts(_stakingPool common.Address, _zapContract common.Address) (*types.Transaction, error) {
	return _JuicedStakingRewardsRegistry.Contract.SetDefaultContracts(&_JuicedStakingRewardsRegistry.TransactOpts, _stakingPool, _zapContract)
}

// SetDefaultContracts is a paid mutator transaction binding the contract method 0x0c689126.
//
// Solidity: function setDefaultContracts(address _stakingPool, address _zapContract) returns()
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryTransactorSession) SetDefaultContracts(_stakingPool common.Address, _zapContract common.Address) (*types.Transaction, error) {
	return _JuicedStakingRewardsRegistry.Contract.SetDefaultContracts(&_JuicedStakingRewardsRegistry.TransactOpts, _stakingPool, _zapContract)
}

// SetPoolEndorsers is a paid mutator transaction binding the contract method 0x5ea38625.
//
// Solidity: function setPoolEndorsers(address _addr, bool _approved) returns()
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryTransactor) SetPoolEndorsers(opts *bind.TransactOpts, _addr common.Address, _approved bool) (*types.Transaction, error) {
	return _JuicedStakingRewardsRegistry.contract.Transact(opts, "setPoolEndorsers", _addr, _approved)
}

// SetPoolEndorsers is a paid mutator transaction binding the contract method 0x5ea38625.
//
// Solidity: function setPoolEndorsers(address _addr, bool _approved) returns()
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistrySession) SetPoolEndorsers(_addr common.Address, _approved bool) (*types.Transaction, error) {
	return _JuicedStakingRewardsRegistry.Contract.SetPoolEndorsers(&_JuicedStakingRewardsRegistry.TransactOpts, _addr, _approved)
}

// SetPoolEndorsers is a paid mutator transaction binding the contract method 0x5ea38625.
//
// Solidity: function setPoolEndorsers(address _addr, bool _approved) returns()
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryTransactorSession) SetPoolEndorsers(_addr common.Address, _approved bool) (*types.Transaction, error) {
	return _JuicedStakingRewardsRegistry.Contract.SetPoolEndorsers(&_JuicedStakingRewardsRegistry.TransactOpts, _addr, _approved)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _JuicedStakingRewardsRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _JuicedStakingRewardsRegistry.Contract.TransferOwnership(&_JuicedStakingRewardsRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _JuicedStakingRewardsRegistry.Contract.TransferOwnership(&_JuicedStakingRewardsRegistry.TransactOpts, newOwner)
}

// JuicedStakingRewardsRegistryApprovedPoolEndorserIterator is returned from FilterApprovedPoolEndorser and is used to iterate over the raw logs and unpacked data for ApprovedPoolEndorser events raised by the JuicedStakingRewardsRegistry contract.
type JuicedStakingRewardsRegistryApprovedPoolEndorserIterator struct {
	Event *JuicedStakingRewardsRegistryApprovedPoolEndorser // Event containing the contract specifics and raw log

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
func (it *JuicedStakingRewardsRegistryApprovedPoolEndorserIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JuicedStakingRewardsRegistryApprovedPoolEndorser)
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
		it.Event = new(JuicedStakingRewardsRegistryApprovedPoolEndorser)
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
func (it *JuicedStakingRewardsRegistryApprovedPoolEndorserIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JuicedStakingRewardsRegistryApprovedPoolEndorserIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JuicedStakingRewardsRegistryApprovedPoolEndorser represents a ApprovedPoolEndorser event raised by the JuicedStakingRewardsRegistry contract.
type JuicedStakingRewardsRegistryApprovedPoolEndorser struct {
	Account    common.Address
	CanEndorse bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterApprovedPoolEndorser is a free log retrieval operation binding the contract event 0x971fecdddf4686fb97d914426cefb86d567c8e5081e48f0346d9ea0c931cac69.
//
// Solidity: event ApprovedPoolEndorser(address account, bool canEndorse)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryFilterer) FilterApprovedPoolEndorser(opts *bind.FilterOpts) (*JuicedStakingRewardsRegistryApprovedPoolEndorserIterator, error) {

	logs, sub, err := _JuicedStakingRewardsRegistry.contract.FilterLogs(opts, "ApprovedPoolEndorser")
	if err != nil {
		return nil, err
	}
	return &JuicedStakingRewardsRegistryApprovedPoolEndorserIterator{contract: _JuicedStakingRewardsRegistry.contract, event: "ApprovedPoolEndorser", logs: logs, sub: sub}, nil
}

// WatchApprovedPoolEndorser is a free log subscription operation binding the contract event 0x971fecdddf4686fb97d914426cefb86d567c8e5081e48f0346d9ea0c931cac69.
//
// Solidity: event ApprovedPoolEndorser(address account, bool canEndorse)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryFilterer) WatchApprovedPoolEndorser(opts *bind.WatchOpts, sink chan<- *JuicedStakingRewardsRegistryApprovedPoolEndorser) (event.Subscription, error) {

	logs, sub, err := _JuicedStakingRewardsRegistry.contract.WatchLogs(opts, "ApprovedPoolEndorser")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JuicedStakingRewardsRegistryApprovedPoolEndorser)
				if err := _JuicedStakingRewardsRegistry.contract.UnpackLog(event, "ApprovedPoolEndorser", log); err != nil {
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
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryFilterer) ParseApprovedPoolEndorser(log types.Log) (*JuicedStakingRewardsRegistryApprovedPoolEndorser, error) {
	event := new(JuicedStakingRewardsRegistryApprovedPoolEndorser)
	if err := _JuicedStakingRewardsRegistry.contract.UnpackLog(event, "ApprovedPoolEndorser", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JuicedStakingRewardsRegistryApprovedPoolOwnerUpdatedIterator is returned from FilterApprovedPoolOwnerUpdated and is used to iterate over the raw logs and unpacked data for ApprovedPoolOwnerUpdated events raised by the JuicedStakingRewardsRegistry contract.
type JuicedStakingRewardsRegistryApprovedPoolOwnerUpdatedIterator struct {
	Event *JuicedStakingRewardsRegistryApprovedPoolOwnerUpdated // Event containing the contract specifics and raw log

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
func (it *JuicedStakingRewardsRegistryApprovedPoolOwnerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JuicedStakingRewardsRegistryApprovedPoolOwnerUpdated)
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
		it.Event = new(JuicedStakingRewardsRegistryApprovedPoolOwnerUpdated)
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
func (it *JuicedStakingRewardsRegistryApprovedPoolOwnerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JuicedStakingRewardsRegistryApprovedPoolOwnerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JuicedStakingRewardsRegistryApprovedPoolOwnerUpdated represents a ApprovedPoolOwnerUpdated event raised by the JuicedStakingRewardsRegistry contract.
type JuicedStakingRewardsRegistryApprovedPoolOwnerUpdated struct {
	Governance common.Address
	Approved   bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterApprovedPoolOwnerUpdated is a free log retrieval operation binding the contract event 0xb3b57e8caeb8200bd5bdd2e4c7a421b0f9683a9e9a376d8a9c6c29883f7a7292.
//
// Solidity: event ApprovedPoolOwnerUpdated(address governance, bool approved)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryFilterer) FilterApprovedPoolOwnerUpdated(opts *bind.FilterOpts) (*JuicedStakingRewardsRegistryApprovedPoolOwnerUpdatedIterator, error) {

	logs, sub, err := _JuicedStakingRewardsRegistry.contract.FilterLogs(opts, "ApprovedPoolOwnerUpdated")
	if err != nil {
		return nil, err
	}
	return &JuicedStakingRewardsRegistryApprovedPoolOwnerUpdatedIterator{contract: _JuicedStakingRewardsRegistry.contract, event: "ApprovedPoolOwnerUpdated", logs: logs, sub: sub}, nil
}

// WatchApprovedPoolOwnerUpdated is a free log subscription operation binding the contract event 0xb3b57e8caeb8200bd5bdd2e4c7a421b0f9683a9e9a376d8a9c6c29883f7a7292.
//
// Solidity: event ApprovedPoolOwnerUpdated(address governance, bool approved)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryFilterer) WatchApprovedPoolOwnerUpdated(opts *bind.WatchOpts, sink chan<- *JuicedStakingRewardsRegistryApprovedPoolOwnerUpdated) (event.Subscription, error) {

	logs, sub, err := _JuicedStakingRewardsRegistry.contract.WatchLogs(opts, "ApprovedPoolOwnerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JuicedStakingRewardsRegistryApprovedPoolOwnerUpdated)
				if err := _JuicedStakingRewardsRegistry.contract.UnpackLog(event, "ApprovedPoolOwnerUpdated", log); err != nil {
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
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryFilterer) ParseApprovedPoolOwnerUpdated(log types.Log) (*JuicedStakingRewardsRegistryApprovedPoolOwnerUpdated, error) {
	event := new(JuicedStakingRewardsRegistryApprovedPoolOwnerUpdated)
	if err := _JuicedStakingRewardsRegistry.contract.UnpackLog(event, "ApprovedPoolOwnerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JuicedStakingRewardsRegistryDefaultContractsUpdatedIterator is returned from FilterDefaultContractsUpdated and is used to iterate over the raw logs and unpacked data for DefaultContractsUpdated events raised by the JuicedStakingRewardsRegistry contract.
type JuicedStakingRewardsRegistryDefaultContractsUpdatedIterator struct {
	Event *JuicedStakingRewardsRegistryDefaultContractsUpdated // Event containing the contract specifics and raw log

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
func (it *JuicedStakingRewardsRegistryDefaultContractsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JuicedStakingRewardsRegistryDefaultContractsUpdated)
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
		it.Event = new(JuicedStakingRewardsRegistryDefaultContractsUpdated)
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
func (it *JuicedStakingRewardsRegistryDefaultContractsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JuicedStakingRewardsRegistryDefaultContractsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JuicedStakingRewardsRegistryDefaultContractsUpdated represents a DefaultContractsUpdated event raised by the JuicedStakingRewardsRegistry contract.
type JuicedStakingRewardsRegistryDefaultContractsUpdated struct {
	StakingContract common.Address
	ZapContract     common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDefaultContractsUpdated is a free log retrieval operation binding the contract event 0xb5be6a994a6c1e2ba32dd4dca027469b4c1d3eaacb8aba9d5815dd2d947e6c40.
//
// Solidity: event DefaultContractsUpdated(address stakingContract, address zapContract)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryFilterer) FilterDefaultContractsUpdated(opts *bind.FilterOpts) (*JuicedStakingRewardsRegistryDefaultContractsUpdatedIterator, error) {

	logs, sub, err := _JuicedStakingRewardsRegistry.contract.FilterLogs(opts, "DefaultContractsUpdated")
	if err != nil {
		return nil, err
	}
	return &JuicedStakingRewardsRegistryDefaultContractsUpdatedIterator{contract: _JuicedStakingRewardsRegistry.contract, event: "DefaultContractsUpdated", logs: logs, sub: sub}, nil
}

// WatchDefaultContractsUpdated is a free log subscription operation binding the contract event 0xb5be6a994a6c1e2ba32dd4dca027469b4c1d3eaacb8aba9d5815dd2d947e6c40.
//
// Solidity: event DefaultContractsUpdated(address stakingContract, address zapContract)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryFilterer) WatchDefaultContractsUpdated(opts *bind.WatchOpts, sink chan<- *JuicedStakingRewardsRegistryDefaultContractsUpdated) (event.Subscription, error) {

	logs, sub, err := _JuicedStakingRewardsRegistry.contract.WatchLogs(opts, "DefaultContractsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JuicedStakingRewardsRegistryDefaultContractsUpdated)
				if err := _JuicedStakingRewardsRegistry.contract.UnpackLog(event, "DefaultContractsUpdated", log); err != nil {
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

// ParseDefaultContractsUpdated is a log parse operation binding the contract event 0xb5be6a994a6c1e2ba32dd4dca027469b4c1d3eaacb8aba9d5815dd2d947e6c40.
//
// Solidity: event DefaultContractsUpdated(address stakingContract, address zapContract)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryFilterer) ParseDefaultContractsUpdated(log types.Log) (*JuicedStakingRewardsRegistryDefaultContractsUpdated, error) {
	event := new(JuicedStakingRewardsRegistryDefaultContractsUpdated)
	if err := _JuicedStakingRewardsRegistry.contract.UnpackLog(event, "DefaultContractsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JuicedStakingRewardsRegistryOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the JuicedStakingRewardsRegistry contract.
type JuicedStakingRewardsRegistryOwnershipTransferStartedIterator struct {
	Event *JuicedStakingRewardsRegistryOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *JuicedStakingRewardsRegistryOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JuicedStakingRewardsRegistryOwnershipTransferStarted)
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
		it.Event = new(JuicedStakingRewardsRegistryOwnershipTransferStarted)
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
func (it *JuicedStakingRewardsRegistryOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JuicedStakingRewardsRegistryOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JuicedStakingRewardsRegistryOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the JuicedStakingRewardsRegistry contract.
type JuicedStakingRewardsRegistryOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*JuicedStakingRewardsRegistryOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _JuicedStakingRewardsRegistry.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &JuicedStakingRewardsRegistryOwnershipTransferStartedIterator{contract: _JuicedStakingRewardsRegistry.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *JuicedStakingRewardsRegistryOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _JuicedStakingRewardsRegistry.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JuicedStakingRewardsRegistryOwnershipTransferStarted)
				if err := _JuicedStakingRewardsRegistry.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryFilterer) ParseOwnershipTransferStarted(log types.Log) (*JuicedStakingRewardsRegistryOwnershipTransferStarted, error) {
	event := new(JuicedStakingRewardsRegistryOwnershipTransferStarted)
	if err := _JuicedStakingRewardsRegistry.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JuicedStakingRewardsRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the JuicedStakingRewardsRegistry contract.
type JuicedStakingRewardsRegistryOwnershipTransferredIterator struct {
	Event *JuicedStakingRewardsRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *JuicedStakingRewardsRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JuicedStakingRewardsRegistryOwnershipTransferred)
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
		it.Event = new(JuicedStakingRewardsRegistryOwnershipTransferred)
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
func (it *JuicedStakingRewardsRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JuicedStakingRewardsRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JuicedStakingRewardsRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the JuicedStakingRewardsRegistry contract.
type JuicedStakingRewardsRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*JuicedStakingRewardsRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _JuicedStakingRewardsRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &JuicedStakingRewardsRegistryOwnershipTransferredIterator{contract: _JuicedStakingRewardsRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *JuicedStakingRewardsRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _JuicedStakingRewardsRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JuicedStakingRewardsRegistryOwnershipTransferred)
				if err := _JuicedStakingRewardsRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*JuicedStakingRewardsRegistryOwnershipTransferred, error) {
	event := new(JuicedStakingRewardsRegistryOwnershipTransferred)
	if err := _JuicedStakingRewardsRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JuicedStakingRewardsRegistryStakingPoolAddedIterator is returned from FilterStakingPoolAdded and is used to iterate over the raw logs and unpacked data for StakingPoolAdded events raised by the JuicedStakingRewardsRegistry contract.
type JuicedStakingRewardsRegistryStakingPoolAddedIterator struct {
	Event *JuicedStakingRewardsRegistryStakingPoolAdded // Event containing the contract specifics and raw log

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
func (it *JuicedStakingRewardsRegistryStakingPoolAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JuicedStakingRewardsRegistryStakingPoolAdded)
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
		it.Event = new(JuicedStakingRewardsRegistryStakingPoolAdded)
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
func (it *JuicedStakingRewardsRegistryStakingPoolAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JuicedStakingRewardsRegistryStakingPoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JuicedStakingRewardsRegistryStakingPoolAdded represents a StakingPoolAdded event raised by the JuicedStakingRewardsRegistry contract.
type JuicedStakingRewardsRegistryStakingPoolAdded struct {
	Token       common.Address
	StakingPool common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStakingPoolAdded is a free log retrieval operation binding the contract event 0xd58b22ec3b77fb836c6ae1bba270411b0fa4961ff8423444de63e7f804826f74.
//
// Solidity: event StakingPoolAdded(address indexed token, address stakingPool)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryFilterer) FilterStakingPoolAdded(opts *bind.FilterOpts, token []common.Address) (*JuicedStakingRewardsRegistryStakingPoolAddedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _JuicedStakingRewardsRegistry.contract.FilterLogs(opts, "StakingPoolAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return &JuicedStakingRewardsRegistryStakingPoolAddedIterator{contract: _JuicedStakingRewardsRegistry.contract, event: "StakingPoolAdded", logs: logs, sub: sub}, nil
}

// WatchStakingPoolAdded is a free log subscription operation binding the contract event 0xd58b22ec3b77fb836c6ae1bba270411b0fa4961ff8423444de63e7f804826f74.
//
// Solidity: event StakingPoolAdded(address indexed token, address stakingPool)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryFilterer) WatchStakingPoolAdded(opts *bind.WatchOpts, sink chan<- *JuicedStakingRewardsRegistryStakingPoolAdded, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _JuicedStakingRewardsRegistry.contract.WatchLogs(opts, "StakingPoolAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JuicedStakingRewardsRegistryStakingPoolAdded)
				if err := _JuicedStakingRewardsRegistry.contract.UnpackLog(event, "StakingPoolAdded", log); err != nil {
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
// Solidity: event StakingPoolAdded(address indexed token, address stakingPool)
func (_JuicedStakingRewardsRegistry *JuicedStakingRewardsRegistryFilterer) ParseStakingPoolAdded(log types.Log) (*JuicedStakingRewardsRegistryStakingPoolAdded, error) {
	event := new(JuicedStakingRewardsRegistryStakingPoolAdded)
	if err := _JuicedStakingRewardsRegistry.contract.UnpackLog(event, "StakingPoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
