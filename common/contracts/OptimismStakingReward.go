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

// YOptimismStakingRewardMetaData contains all meta data concerning the YOptimismStakingReward contract.
var YOptimismStakingRewardMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"canEndorse\",\"type\":\"bool\"}],\"name\":\"ApprovedPoolEndorser\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governance\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovedPoolOwnerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakingPool\",\"type\":\"address\"}],\"name\":\"StakingPoolAdded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingPool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_replaceExistingPool\",\"type\":\"bool\"}],\"name\":\"addStakingPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"approvedPoolOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isStakingPoolEndorsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"poolEndorsers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovedPoolOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setPoolEndorsers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakingPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// YOptimismStakingRewardABI is the input ABI used to generate the binding from.
// Deprecated: Use YOptimismStakingRewardMetaData.ABI instead.
var YOptimismStakingRewardABI = YOptimismStakingRewardMetaData.ABI

// YOptimismStakingReward is an auto generated Go binding around an Ethereum contract.
type YOptimismStakingReward struct {
	YOptimismStakingRewardCaller     // Read-only binding to the contract
	YOptimismStakingRewardTransactor // Write-only binding to the contract
	YOptimismStakingRewardFilterer   // Log filterer for contract events
}

// YOptimismStakingRewardCaller is an auto generated read-only Go binding around an Ethereum contract.
type YOptimismStakingRewardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YOptimismStakingRewardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YOptimismStakingRewardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YOptimismStakingRewardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YOptimismStakingRewardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YOptimismStakingRewardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YOptimismStakingRewardSession struct {
	Contract     *YOptimismStakingReward // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// YOptimismStakingRewardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YOptimismStakingRewardCallerSession struct {
	Contract *YOptimismStakingRewardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// YOptimismStakingRewardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YOptimismStakingRewardTransactorSession struct {
	Contract     *YOptimismStakingRewardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// YOptimismStakingRewardRaw is an auto generated low-level Go binding around an Ethereum contract.
type YOptimismStakingRewardRaw struct {
	Contract *YOptimismStakingReward // Generic contract binding to access the raw methods on
}

// YOptimismStakingRewardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YOptimismStakingRewardCallerRaw struct {
	Contract *YOptimismStakingRewardCaller // Generic read-only contract binding to access the raw methods on
}

// YOptimismStakingRewardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YOptimismStakingRewardTransactorRaw struct {
	Contract *YOptimismStakingRewardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYOptimismStakingReward creates a new instance of YOptimismStakingReward, bound to a specific deployed contract.
func NewYOptimismStakingReward(address common.Address, backend bind.ContractBackend) (*YOptimismStakingReward, error) {
	contract, err := bindYOptimismStakingReward(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YOptimismStakingReward{YOptimismStakingRewardCaller: YOptimismStakingRewardCaller{contract: contract}, YOptimismStakingRewardTransactor: YOptimismStakingRewardTransactor{contract: contract}, YOptimismStakingRewardFilterer: YOptimismStakingRewardFilterer{contract: contract}}, nil
}

// NewYOptimismStakingRewardCaller creates a new read-only instance of YOptimismStakingReward, bound to a specific deployed contract.
func NewYOptimismStakingRewardCaller(address common.Address, caller bind.ContractCaller) (*YOptimismStakingRewardCaller, error) {
	contract, err := bindYOptimismStakingReward(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YOptimismStakingRewardCaller{contract: contract}, nil
}

// NewYOptimismStakingRewardTransactor creates a new write-only instance of YOptimismStakingReward, bound to a specific deployed contract.
func NewYOptimismStakingRewardTransactor(address common.Address, transactor bind.ContractTransactor) (*YOptimismStakingRewardTransactor, error) {
	contract, err := bindYOptimismStakingReward(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YOptimismStakingRewardTransactor{contract: contract}, nil
}

// NewYOptimismStakingRewardFilterer creates a new log filterer instance of YOptimismStakingReward, bound to a specific deployed contract.
func NewYOptimismStakingRewardFilterer(address common.Address, filterer bind.ContractFilterer) (*YOptimismStakingRewardFilterer, error) {
	contract, err := bindYOptimismStakingReward(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YOptimismStakingRewardFilterer{contract: contract}, nil
}

// bindYOptimismStakingReward binds a generic wrapper to an already deployed contract.
func bindYOptimismStakingReward(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(YOptimismStakingRewardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YOptimismStakingReward *YOptimismStakingRewardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YOptimismStakingReward.Contract.YOptimismStakingRewardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YOptimismStakingReward *YOptimismStakingRewardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YOptimismStakingReward.Contract.YOptimismStakingRewardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YOptimismStakingReward *YOptimismStakingRewardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YOptimismStakingReward.Contract.YOptimismStakingRewardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YOptimismStakingReward *YOptimismStakingRewardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YOptimismStakingReward.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YOptimismStakingReward *YOptimismStakingRewardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YOptimismStakingReward.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YOptimismStakingReward *YOptimismStakingRewardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YOptimismStakingReward.Contract.contract.Transact(opts, method, params...)
}

// ApprovedPoolOwner is a free data retrieval call binding the contract method 0x4f2ac26e.
//
// Solidity: function approvedPoolOwner(address ) view returns(bool)
func (_YOptimismStakingReward *YOptimismStakingRewardCaller) ApprovedPoolOwner(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _YOptimismStakingReward.contract.Call(opts, &out, "approvedPoolOwner", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ApprovedPoolOwner is a free data retrieval call binding the contract method 0x4f2ac26e.
//
// Solidity: function approvedPoolOwner(address ) view returns(bool)
func (_YOptimismStakingReward *YOptimismStakingRewardSession) ApprovedPoolOwner(arg0 common.Address) (bool, error) {
	return _YOptimismStakingReward.Contract.ApprovedPoolOwner(&_YOptimismStakingReward.CallOpts, arg0)
}

// ApprovedPoolOwner is a free data retrieval call binding the contract method 0x4f2ac26e.
//
// Solidity: function approvedPoolOwner(address ) view returns(bool)
func (_YOptimismStakingReward *YOptimismStakingRewardCallerSession) ApprovedPoolOwner(arg0 common.Address) (bool, error) {
	return _YOptimismStakingReward.Contract.ApprovedPoolOwner(&_YOptimismStakingReward.CallOpts, arg0)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address ) view returns(bool)
func (_YOptimismStakingReward *YOptimismStakingRewardCaller) IsRegistered(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _YOptimismStakingReward.contract.Call(opts, &out, "isRegistered", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address ) view returns(bool)
func (_YOptimismStakingReward *YOptimismStakingRewardSession) IsRegistered(arg0 common.Address) (bool, error) {
	return _YOptimismStakingReward.Contract.IsRegistered(&_YOptimismStakingReward.CallOpts, arg0)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address ) view returns(bool)
func (_YOptimismStakingReward *YOptimismStakingRewardCallerSession) IsRegistered(arg0 common.Address) (bool, error) {
	return _YOptimismStakingReward.Contract.IsRegistered(&_YOptimismStakingReward.CallOpts, arg0)
}

// IsStakingPoolEndorsed is a free data retrieval call binding the contract method 0x89884973.
//
// Solidity: function isStakingPoolEndorsed(address ) view returns(bool)
func (_YOptimismStakingReward *YOptimismStakingRewardCaller) IsStakingPoolEndorsed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _YOptimismStakingReward.contract.Call(opts, &out, "isStakingPoolEndorsed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStakingPoolEndorsed is a free data retrieval call binding the contract method 0x89884973.
//
// Solidity: function isStakingPoolEndorsed(address ) view returns(bool)
func (_YOptimismStakingReward *YOptimismStakingRewardSession) IsStakingPoolEndorsed(arg0 common.Address) (bool, error) {
	return _YOptimismStakingReward.Contract.IsStakingPoolEndorsed(&_YOptimismStakingReward.CallOpts, arg0)
}

// IsStakingPoolEndorsed is a free data retrieval call binding the contract method 0x89884973.
//
// Solidity: function isStakingPoolEndorsed(address ) view returns(bool)
func (_YOptimismStakingReward *YOptimismStakingRewardCallerSession) IsStakingPoolEndorsed(arg0 common.Address) (bool, error) {
	return _YOptimismStakingReward.Contract.IsStakingPoolEndorsed(&_YOptimismStakingReward.CallOpts, arg0)
}

// NumTokens is a free data retrieval call binding the contract method 0x8e499bcf.
//
// Solidity: function numTokens() view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardCaller) NumTokens(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YOptimismStakingReward.contract.Call(opts, &out, "numTokens")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumTokens is a free data retrieval call binding the contract method 0x8e499bcf.
//
// Solidity: function numTokens() view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardSession) NumTokens() (*big.Int, error) {
	return _YOptimismStakingReward.Contract.NumTokens(&_YOptimismStakingReward.CallOpts)
}

// NumTokens is a free data retrieval call binding the contract method 0x8e499bcf.
//
// Solidity: function numTokens() view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardCallerSession) NumTokens() (*big.Int, error) {
	return _YOptimismStakingReward.Contract.NumTokens(&_YOptimismStakingReward.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_YOptimismStakingReward *YOptimismStakingRewardCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YOptimismStakingReward.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_YOptimismStakingReward *YOptimismStakingRewardSession) Owner() (common.Address, error) {
	return _YOptimismStakingReward.Contract.Owner(&_YOptimismStakingReward.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_YOptimismStakingReward *YOptimismStakingRewardCallerSession) Owner() (common.Address, error) {
	return _YOptimismStakingReward.Contract.Owner(&_YOptimismStakingReward.CallOpts)
}

// PoolEndorsers is a free data retrieval call binding the contract method 0x4f1fde98.
//
// Solidity: function poolEndorsers(address ) view returns(bool)
func (_YOptimismStakingReward *YOptimismStakingRewardCaller) PoolEndorsers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _YOptimismStakingReward.contract.Call(opts, &out, "poolEndorsers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PoolEndorsers is a free data retrieval call binding the contract method 0x4f1fde98.
//
// Solidity: function poolEndorsers(address ) view returns(bool)
func (_YOptimismStakingReward *YOptimismStakingRewardSession) PoolEndorsers(arg0 common.Address) (bool, error) {
	return _YOptimismStakingReward.Contract.PoolEndorsers(&_YOptimismStakingReward.CallOpts, arg0)
}

// PoolEndorsers is a free data retrieval call binding the contract method 0x4f1fde98.
//
// Solidity: function poolEndorsers(address ) view returns(bool)
func (_YOptimismStakingReward *YOptimismStakingRewardCallerSession) PoolEndorsers(arg0 common.Address) (bool, error) {
	return _YOptimismStakingReward.Contract.PoolEndorsers(&_YOptimismStakingReward.CallOpts, arg0)
}

// StakingPool is a free data retrieval call binding the contract method 0x8dddb3e5.
//
// Solidity: function stakingPool(address ) view returns(address)
func (_YOptimismStakingReward *YOptimismStakingRewardCaller) StakingPool(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _YOptimismStakingReward.contract.Call(opts, &out, "stakingPool", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingPool is a free data retrieval call binding the contract method 0x8dddb3e5.
//
// Solidity: function stakingPool(address ) view returns(address)
func (_YOptimismStakingReward *YOptimismStakingRewardSession) StakingPool(arg0 common.Address) (common.Address, error) {
	return _YOptimismStakingReward.Contract.StakingPool(&_YOptimismStakingReward.CallOpts, arg0)
}

// StakingPool is a free data retrieval call binding the contract method 0x8dddb3e5.
//
// Solidity: function stakingPool(address ) view returns(address)
func (_YOptimismStakingReward *YOptimismStakingRewardCallerSession) StakingPool(arg0 common.Address) (common.Address, error) {
	return _YOptimismStakingReward.Contract.StakingPool(&_YOptimismStakingReward.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_YOptimismStakingReward *YOptimismStakingRewardCaller) Tokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _YOptimismStakingReward.contract.Call(opts, &out, "tokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_YOptimismStakingReward *YOptimismStakingRewardSession) Tokens(arg0 *big.Int) (common.Address, error) {
	return _YOptimismStakingReward.Contract.Tokens(&_YOptimismStakingReward.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_YOptimismStakingReward *YOptimismStakingRewardCallerSession) Tokens(arg0 *big.Int) (common.Address, error) {
	return _YOptimismStakingReward.Contract.Tokens(&_YOptimismStakingReward.CallOpts, arg0)
}

// AddStakingPool is a paid mutator transaction binding the contract method 0xaa9f079c.
//
// Solidity: function addStakingPool(address _stakingPool, address _token, bool _replaceExistingPool) returns()
func (_YOptimismStakingReward *YOptimismStakingRewardTransactor) AddStakingPool(opts *bind.TransactOpts, _stakingPool common.Address, _token common.Address, _replaceExistingPool bool) (*types.Transaction, error) {
	return _YOptimismStakingReward.contract.Transact(opts, "addStakingPool", _stakingPool, _token, _replaceExistingPool)
}

// AddStakingPool is a paid mutator transaction binding the contract method 0xaa9f079c.
//
// Solidity: function addStakingPool(address _stakingPool, address _token, bool _replaceExistingPool) returns()
func (_YOptimismStakingReward *YOptimismStakingRewardSession) AddStakingPool(_stakingPool common.Address, _token common.Address, _replaceExistingPool bool) (*types.Transaction, error) {
	return _YOptimismStakingReward.Contract.AddStakingPool(&_YOptimismStakingReward.TransactOpts, _stakingPool, _token, _replaceExistingPool)
}

// AddStakingPool is a paid mutator transaction binding the contract method 0xaa9f079c.
//
// Solidity: function addStakingPool(address _stakingPool, address _token, bool _replaceExistingPool) returns()
func (_YOptimismStakingReward *YOptimismStakingRewardTransactorSession) AddStakingPool(_stakingPool common.Address, _token common.Address, _replaceExistingPool bool) (*types.Transaction, error) {
	return _YOptimismStakingReward.Contract.AddStakingPool(&_YOptimismStakingReward.TransactOpts, _stakingPool, _token, _replaceExistingPool)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_YOptimismStakingReward *YOptimismStakingRewardTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YOptimismStakingReward.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_YOptimismStakingReward *YOptimismStakingRewardSession) RenounceOwnership() (*types.Transaction, error) {
	return _YOptimismStakingReward.Contract.RenounceOwnership(&_YOptimismStakingReward.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_YOptimismStakingReward *YOptimismStakingRewardTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _YOptimismStakingReward.Contract.RenounceOwnership(&_YOptimismStakingReward.TransactOpts)
}

// SetApprovedPoolOwner is a paid mutator transaction binding the contract method 0xaa9fd6ef.
//
// Solidity: function setApprovedPoolOwner(address _addr, bool _approved) returns()
func (_YOptimismStakingReward *YOptimismStakingRewardTransactor) SetApprovedPoolOwner(opts *bind.TransactOpts, _addr common.Address, _approved bool) (*types.Transaction, error) {
	return _YOptimismStakingReward.contract.Transact(opts, "setApprovedPoolOwner", _addr, _approved)
}

// SetApprovedPoolOwner is a paid mutator transaction binding the contract method 0xaa9fd6ef.
//
// Solidity: function setApprovedPoolOwner(address _addr, bool _approved) returns()
func (_YOptimismStakingReward *YOptimismStakingRewardSession) SetApprovedPoolOwner(_addr common.Address, _approved bool) (*types.Transaction, error) {
	return _YOptimismStakingReward.Contract.SetApprovedPoolOwner(&_YOptimismStakingReward.TransactOpts, _addr, _approved)
}

// SetApprovedPoolOwner is a paid mutator transaction binding the contract method 0xaa9fd6ef.
//
// Solidity: function setApprovedPoolOwner(address _addr, bool _approved) returns()
func (_YOptimismStakingReward *YOptimismStakingRewardTransactorSession) SetApprovedPoolOwner(_addr common.Address, _approved bool) (*types.Transaction, error) {
	return _YOptimismStakingReward.Contract.SetApprovedPoolOwner(&_YOptimismStakingReward.TransactOpts, _addr, _approved)
}

// SetPoolEndorsers is a paid mutator transaction binding the contract method 0x5ea38625.
//
// Solidity: function setPoolEndorsers(address _addr, bool _approved) returns()
func (_YOptimismStakingReward *YOptimismStakingRewardTransactor) SetPoolEndorsers(opts *bind.TransactOpts, _addr common.Address, _approved bool) (*types.Transaction, error) {
	return _YOptimismStakingReward.contract.Transact(opts, "setPoolEndorsers", _addr, _approved)
}

// SetPoolEndorsers is a paid mutator transaction binding the contract method 0x5ea38625.
//
// Solidity: function setPoolEndorsers(address _addr, bool _approved) returns()
func (_YOptimismStakingReward *YOptimismStakingRewardSession) SetPoolEndorsers(_addr common.Address, _approved bool) (*types.Transaction, error) {
	return _YOptimismStakingReward.Contract.SetPoolEndorsers(&_YOptimismStakingReward.TransactOpts, _addr, _approved)
}

// SetPoolEndorsers is a paid mutator transaction binding the contract method 0x5ea38625.
//
// Solidity: function setPoolEndorsers(address _addr, bool _approved) returns()
func (_YOptimismStakingReward *YOptimismStakingRewardTransactorSession) SetPoolEndorsers(_addr common.Address, _approved bool) (*types.Transaction, error) {
	return _YOptimismStakingReward.Contract.SetPoolEndorsers(&_YOptimismStakingReward.TransactOpts, _addr, _approved)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_YOptimismStakingReward *YOptimismStakingRewardTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _YOptimismStakingReward.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_YOptimismStakingReward *YOptimismStakingRewardSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _YOptimismStakingReward.Contract.TransferOwnership(&_YOptimismStakingReward.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_YOptimismStakingReward *YOptimismStakingRewardTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _YOptimismStakingReward.Contract.TransferOwnership(&_YOptimismStakingReward.TransactOpts, newOwner)
}

// YOptimismStakingRewardApprovedPoolEndorserIterator is returned from FilterApprovedPoolEndorser and is used to iterate over the raw logs and unpacked data for ApprovedPoolEndorser events raised by the YOptimismStakingReward contract.
type YOptimismStakingRewardApprovedPoolEndorserIterator struct {
	Event *YOptimismStakingRewardApprovedPoolEndorser // Event containing the contract specifics and raw log

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
func (it *YOptimismStakingRewardApprovedPoolEndorserIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YOptimismStakingRewardApprovedPoolEndorser)
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
		it.Event = new(YOptimismStakingRewardApprovedPoolEndorser)
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
func (it *YOptimismStakingRewardApprovedPoolEndorserIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YOptimismStakingRewardApprovedPoolEndorserIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YOptimismStakingRewardApprovedPoolEndorser represents a ApprovedPoolEndorser event raised by the YOptimismStakingReward contract.
type YOptimismStakingRewardApprovedPoolEndorser struct {
	Account    common.Address
	CanEndorse bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterApprovedPoolEndorser is a free log retrieval operation binding the contract event 0x971fecdddf4686fb97d914426cefb86d567c8e5081e48f0346d9ea0c931cac69.
//
// Solidity: event ApprovedPoolEndorser(address account, bool canEndorse)
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) FilterApprovedPoolEndorser(opts *bind.FilterOpts) (*YOptimismStakingRewardApprovedPoolEndorserIterator, error) {

	logs, sub, err := _YOptimismStakingReward.contract.FilterLogs(opts, "ApprovedPoolEndorser")
	if err != nil {
		return nil, err
	}
	return &YOptimismStakingRewardApprovedPoolEndorserIterator{contract: _YOptimismStakingReward.contract, event: "ApprovedPoolEndorser", logs: logs, sub: sub}, nil
}

// WatchApprovedPoolEndorser is a free log subscription operation binding the contract event 0x971fecdddf4686fb97d914426cefb86d567c8e5081e48f0346d9ea0c931cac69.
//
// Solidity: event ApprovedPoolEndorser(address account, bool canEndorse)
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) WatchApprovedPoolEndorser(opts *bind.WatchOpts, sink chan<- *YOptimismStakingRewardApprovedPoolEndorser) (event.Subscription, error) {

	logs, sub, err := _YOptimismStakingReward.contract.WatchLogs(opts, "ApprovedPoolEndorser")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YOptimismStakingRewardApprovedPoolEndorser)
				if err := _YOptimismStakingReward.contract.UnpackLog(event, "ApprovedPoolEndorser", log); err != nil {
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
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) ParseApprovedPoolEndorser(log types.Log) (*YOptimismStakingRewardApprovedPoolEndorser, error) {
	event := new(YOptimismStakingRewardApprovedPoolEndorser)
	if err := _YOptimismStakingReward.contract.UnpackLog(event, "ApprovedPoolEndorser", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YOptimismStakingRewardApprovedPoolOwnerUpdatedIterator is returned from FilterApprovedPoolOwnerUpdated and is used to iterate over the raw logs and unpacked data for ApprovedPoolOwnerUpdated events raised by the YOptimismStakingReward contract.
type YOptimismStakingRewardApprovedPoolOwnerUpdatedIterator struct {
	Event *YOptimismStakingRewardApprovedPoolOwnerUpdated // Event containing the contract specifics and raw log

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
func (it *YOptimismStakingRewardApprovedPoolOwnerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YOptimismStakingRewardApprovedPoolOwnerUpdated)
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
		it.Event = new(YOptimismStakingRewardApprovedPoolOwnerUpdated)
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
func (it *YOptimismStakingRewardApprovedPoolOwnerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YOptimismStakingRewardApprovedPoolOwnerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YOptimismStakingRewardApprovedPoolOwnerUpdated represents a ApprovedPoolOwnerUpdated event raised by the YOptimismStakingReward contract.
type YOptimismStakingRewardApprovedPoolOwnerUpdated struct {
	Governance common.Address
	Approved   bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterApprovedPoolOwnerUpdated is a free log retrieval operation binding the contract event 0xb3b57e8caeb8200bd5bdd2e4c7a421b0f9683a9e9a376d8a9c6c29883f7a7292.
//
// Solidity: event ApprovedPoolOwnerUpdated(address governance, bool approved)
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) FilterApprovedPoolOwnerUpdated(opts *bind.FilterOpts) (*YOptimismStakingRewardApprovedPoolOwnerUpdatedIterator, error) {

	logs, sub, err := _YOptimismStakingReward.contract.FilterLogs(opts, "ApprovedPoolOwnerUpdated")
	if err != nil {
		return nil, err
	}
	return &YOptimismStakingRewardApprovedPoolOwnerUpdatedIterator{contract: _YOptimismStakingReward.contract, event: "ApprovedPoolOwnerUpdated", logs: logs, sub: sub}, nil
}

// WatchApprovedPoolOwnerUpdated is a free log subscription operation binding the contract event 0xb3b57e8caeb8200bd5bdd2e4c7a421b0f9683a9e9a376d8a9c6c29883f7a7292.
//
// Solidity: event ApprovedPoolOwnerUpdated(address governance, bool approved)
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) WatchApprovedPoolOwnerUpdated(opts *bind.WatchOpts, sink chan<- *YOptimismStakingRewardApprovedPoolOwnerUpdated) (event.Subscription, error) {

	logs, sub, err := _YOptimismStakingReward.contract.WatchLogs(opts, "ApprovedPoolOwnerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YOptimismStakingRewardApprovedPoolOwnerUpdated)
				if err := _YOptimismStakingReward.contract.UnpackLog(event, "ApprovedPoolOwnerUpdated", log); err != nil {
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
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) ParseApprovedPoolOwnerUpdated(log types.Log) (*YOptimismStakingRewardApprovedPoolOwnerUpdated, error) {
	event := new(YOptimismStakingRewardApprovedPoolOwnerUpdated)
	if err := _YOptimismStakingReward.contract.UnpackLog(event, "ApprovedPoolOwnerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YOptimismStakingRewardOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the YOptimismStakingReward contract.
type YOptimismStakingRewardOwnershipTransferredIterator struct {
	Event *YOptimismStakingRewardOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *YOptimismStakingRewardOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YOptimismStakingRewardOwnershipTransferred)
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
		it.Event = new(YOptimismStakingRewardOwnershipTransferred)
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
func (it *YOptimismStakingRewardOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YOptimismStakingRewardOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YOptimismStakingRewardOwnershipTransferred represents a OwnershipTransferred event raised by the YOptimismStakingReward contract.
type YOptimismStakingRewardOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*YOptimismStakingRewardOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _YOptimismStakingReward.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &YOptimismStakingRewardOwnershipTransferredIterator{contract: _YOptimismStakingReward.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *YOptimismStakingRewardOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _YOptimismStakingReward.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YOptimismStakingRewardOwnershipTransferred)
				if err := _YOptimismStakingReward.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) ParseOwnershipTransferred(log types.Log) (*YOptimismStakingRewardOwnershipTransferred, error) {
	event := new(YOptimismStakingRewardOwnershipTransferred)
	if err := _YOptimismStakingReward.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YOptimismStakingRewardStakingPoolAddedIterator is returned from FilterStakingPoolAdded and is used to iterate over the raw logs and unpacked data for StakingPoolAdded events raised by the YOptimismStakingReward contract.
type YOptimismStakingRewardStakingPoolAddedIterator struct {
	Event *YOptimismStakingRewardStakingPoolAdded // Event containing the contract specifics and raw log

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
func (it *YOptimismStakingRewardStakingPoolAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YOptimismStakingRewardStakingPoolAdded)
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
		it.Event = new(YOptimismStakingRewardStakingPoolAdded)
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
func (it *YOptimismStakingRewardStakingPoolAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YOptimismStakingRewardStakingPoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YOptimismStakingRewardStakingPoolAdded represents a StakingPoolAdded event raised by the YOptimismStakingReward contract.
type YOptimismStakingRewardStakingPoolAdded struct {
	Token       common.Address
	StakingPool common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStakingPoolAdded is a free log retrieval operation binding the contract event 0xd58b22ec3b77fb836c6ae1bba270411b0fa4961ff8423444de63e7f804826f74.
//
// Solidity: event StakingPoolAdded(address indexed token, address indexed stakingPool)
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) FilterStakingPoolAdded(opts *bind.FilterOpts, token []common.Address, stakingPool []common.Address) (*YOptimismStakingRewardStakingPoolAddedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var stakingPoolRule []interface{}
	for _, stakingPoolItem := range stakingPool {
		stakingPoolRule = append(stakingPoolRule, stakingPoolItem)
	}

	logs, sub, err := _YOptimismStakingReward.contract.FilterLogs(opts, "StakingPoolAdded", tokenRule, stakingPoolRule)
	if err != nil {
		return nil, err
	}
	return &YOptimismStakingRewardStakingPoolAddedIterator{contract: _YOptimismStakingReward.contract, event: "StakingPoolAdded", logs: logs, sub: sub}, nil
}

// WatchStakingPoolAdded is a free log subscription operation binding the contract event 0xd58b22ec3b77fb836c6ae1bba270411b0fa4961ff8423444de63e7f804826f74.
//
// Solidity: event StakingPoolAdded(address indexed token, address indexed stakingPool)
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) WatchStakingPoolAdded(opts *bind.WatchOpts, sink chan<- *YOptimismStakingRewardStakingPoolAdded, token []common.Address, stakingPool []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var stakingPoolRule []interface{}
	for _, stakingPoolItem := range stakingPool {
		stakingPoolRule = append(stakingPoolRule, stakingPoolItem)
	}

	logs, sub, err := _YOptimismStakingReward.contract.WatchLogs(opts, "StakingPoolAdded", tokenRule, stakingPoolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YOptimismStakingRewardStakingPoolAdded)
				if err := _YOptimismStakingReward.contract.UnpackLog(event, "StakingPoolAdded", log); err != nil {
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
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) ParseStakingPoolAdded(log types.Log) (*YOptimismStakingRewardStakingPoolAdded, error) {
	event := new(YOptimismStakingRewardStakingPoolAdded)
	if err := _YOptimismStakingReward.contract.UnpackLog(event, "StakingPoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
