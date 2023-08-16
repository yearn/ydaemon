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

// YVelodromePoolMetaData contains all meta data concerning the YVelodromePool contract.
var YVelodromePoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_forwarder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakingToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feesVotingReward\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isPool\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"NotAlive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAuthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotVoter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RewardRateTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroRewardRate\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimed0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimed1\",\"type\":\"uint256\"}],\"name\":\"ClaimFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ClaimRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NotifyReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"earned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fees0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fees1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feesVotingReward\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"getReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTimeRewardApplicable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"left\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"notifyRewardAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodFinish\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerTokenStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewardRateByEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRewardPerTokenPaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// YVelodromePoolABI is the input ABI used to generate the binding from.
// Deprecated: Use YVelodromePoolMetaData.ABI instead.
var YVelodromePoolABI = YVelodromePoolMetaData.ABI

// YVelodromePool is an auto generated Go binding around an Ethereum contract.
type YVelodromePool struct {
	YVelodromePoolCaller     // Read-only binding to the contract
	YVelodromePoolTransactor // Write-only binding to the contract
	YVelodromePoolFilterer   // Log filterer for contract events
}

// YVelodromePoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type YVelodromePoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YVelodromePoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YVelodromePoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YVelodromePoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YVelodromePoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YVelodromePoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YVelodromePoolSession struct {
	Contract     *YVelodromePool   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YVelodromePoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YVelodromePoolCallerSession struct {
	Contract *YVelodromePoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// YVelodromePoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YVelodromePoolTransactorSession struct {
	Contract     *YVelodromePoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// YVelodromePoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type YVelodromePoolRaw struct {
	Contract *YVelodromePool // Generic contract binding to access the raw methods on
}

// YVelodromePoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YVelodromePoolCallerRaw struct {
	Contract *YVelodromePoolCaller // Generic read-only contract binding to access the raw methods on
}

// YVelodromePoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YVelodromePoolTransactorRaw struct {
	Contract *YVelodromePoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYVelodromePool creates a new instance of YVelodromePool, bound to a specific deployed contract.
func NewYVelodromePool(address common.Address, backend bind.ContractBackend) (*YVelodromePool, error) {
	contract, err := bindYVelodromePool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YVelodromePool{YVelodromePoolCaller: YVelodromePoolCaller{contract: contract}, YVelodromePoolTransactor: YVelodromePoolTransactor{contract: contract}, YVelodromePoolFilterer: YVelodromePoolFilterer{contract: contract}}, nil
}

// NewYVelodromePoolCaller creates a new read-only instance of YVelodromePool, bound to a specific deployed contract.
func NewYVelodromePoolCaller(address common.Address, caller bind.ContractCaller) (*YVelodromePoolCaller, error) {
	contract, err := bindYVelodromePool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YVelodromePoolCaller{contract: contract}, nil
}

// NewYVelodromePoolTransactor creates a new write-only instance of YVelodromePool, bound to a specific deployed contract.
func NewYVelodromePoolTransactor(address common.Address, transactor bind.ContractTransactor) (*YVelodromePoolTransactor, error) {
	contract, err := bindYVelodromePool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YVelodromePoolTransactor{contract: contract}, nil
}

// NewYVelodromePoolFilterer creates a new log filterer instance of YVelodromePool, bound to a specific deployed contract.
func NewYVelodromePoolFilterer(address common.Address, filterer bind.ContractFilterer) (*YVelodromePoolFilterer, error) {
	contract, err := bindYVelodromePool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YVelodromePoolFilterer{contract: contract}, nil
}

// bindYVelodromePool binds a generic wrapper to an already deployed contract.
func bindYVelodromePool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(YVelodromePoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YVelodromePool *YVelodromePoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YVelodromePool.Contract.YVelodromePoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YVelodromePool *YVelodromePoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YVelodromePool.Contract.YVelodromePoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YVelodromePool *YVelodromePoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YVelodromePool.Contract.YVelodromePoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YVelodromePool *YVelodromePoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YVelodromePool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YVelodromePool *YVelodromePoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YVelodromePool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YVelodromePool *YVelodromePoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YVelodromePool.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_YVelodromePool *YVelodromePoolCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YVelodromePool.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_YVelodromePool *YVelodromePoolSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _YVelodromePool.Contract.BalanceOf(&_YVelodromePool.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_YVelodromePool *YVelodromePoolCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _YVelodromePool.Contract.BalanceOf(&_YVelodromePool.CallOpts, arg0)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address _account) view returns(uint256)
func (_YVelodromePool *YVelodromePoolCaller) Earned(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YVelodromePool.contract.Call(opts, &out, "earned", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address _account) view returns(uint256)
func (_YVelodromePool *YVelodromePoolSession) Earned(_account common.Address) (*big.Int, error) {
	return _YVelodromePool.Contract.Earned(&_YVelodromePool.CallOpts, _account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address _account) view returns(uint256)
func (_YVelodromePool *YVelodromePoolCallerSession) Earned(_account common.Address) (*big.Int, error) {
	return _YVelodromePool.Contract.Earned(&_YVelodromePool.CallOpts, _account)
}

// Fees0 is a free data retrieval call binding the contract method 0x93f1c442.
//
// Solidity: function fees0() view returns(uint256)
func (_YVelodromePool *YVelodromePoolCaller) Fees0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YVelodromePool.contract.Call(opts, &out, "fees0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fees0 is a free data retrieval call binding the contract method 0x93f1c442.
//
// Solidity: function fees0() view returns(uint256)
func (_YVelodromePool *YVelodromePoolSession) Fees0() (*big.Int, error) {
	return _YVelodromePool.Contract.Fees0(&_YVelodromePool.CallOpts)
}

// Fees0 is a free data retrieval call binding the contract method 0x93f1c442.
//
// Solidity: function fees0() view returns(uint256)
func (_YVelodromePool *YVelodromePoolCallerSession) Fees0() (*big.Int, error) {
	return _YVelodromePool.Contract.Fees0(&_YVelodromePool.CallOpts)
}

// Fees1 is a free data retrieval call binding the contract method 0x4c02a21c.
//
// Solidity: function fees1() view returns(uint256)
func (_YVelodromePool *YVelodromePoolCaller) Fees1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YVelodromePool.contract.Call(opts, &out, "fees1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fees1 is a free data retrieval call binding the contract method 0x4c02a21c.
//
// Solidity: function fees1() view returns(uint256)
func (_YVelodromePool *YVelodromePoolSession) Fees1() (*big.Int, error) {
	return _YVelodromePool.Contract.Fees1(&_YVelodromePool.CallOpts)
}

// Fees1 is a free data retrieval call binding the contract method 0x4c02a21c.
//
// Solidity: function fees1() view returns(uint256)
func (_YVelodromePool *YVelodromePoolCallerSession) Fees1() (*big.Int, error) {
	return _YVelodromePool.Contract.Fees1(&_YVelodromePool.CallOpts)
}

// FeesVotingReward is a free data retrieval call binding the contract method 0x0fe2f711.
//
// Solidity: function feesVotingReward() view returns(address)
func (_YVelodromePool *YVelodromePoolCaller) FeesVotingReward(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YVelodromePool.contract.Call(opts, &out, "feesVotingReward")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeesVotingReward is a free data retrieval call binding the contract method 0x0fe2f711.
//
// Solidity: function feesVotingReward() view returns(address)
func (_YVelodromePool *YVelodromePoolSession) FeesVotingReward() (common.Address, error) {
	return _YVelodromePool.Contract.FeesVotingReward(&_YVelodromePool.CallOpts)
}

// FeesVotingReward is a free data retrieval call binding the contract method 0x0fe2f711.
//
// Solidity: function feesVotingReward() view returns(address)
func (_YVelodromePool *YVelodromePoolCallerSession) FeesVotingReward() (common.Address, error) {
	return _YVelodromePool.Contract.FeesVotingReward(&_YVelodromePool.CallOpts)
}

// IsPool is a free data retrieval call binding the contract method 0xe2e1c6db.
//
// Solidity: function isPool() view returns(bool)
func (_YVelodromePool *YVelodromePoolCaller) IsPool(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _YVelodromePool.contract.Call(opts, &out, "isPool")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPool is a free data retrieval call binding the contract method 0xe2e1c6db.
//
// Solidity: function isPool() view returns(bool)
func (_YVelodromePool *YVelodromePoolSession) IsPool() (bool, error) {
	return _YVelodromePool.Contract.IsPool(&_YVelodromePool.CallOpts)
}

// IsPool is a free data retrieval call binding the contract method 0xe2e1c6db.
//
// Solidity: function isPool() view returns(bool)
func (_YVelodromePool *YVelodromePoolCallerSession) IsPool() (bool, error) {
	return _YVelodromePool.Contract.IsPool(&_YVelodromePool.CallOpts)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_YVelodromePool *YVelodromePoolCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _YVelodromePool.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_YVelodromePool *YVelodromePoolSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _YVelodromePool.Contract.IsTrustedForwarder(&_YVelodromePool.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_YVelodromePool *YVelodromePoolCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _YVelodromePool.Contract.IsTrustedForwarder(&_YVelodromePool.CallOpts, forwarder)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_YVelodromePool *YVelodromePoolCaller) LastTimeRewardApplicable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YVelodromePool.contract.Call(opts, &out, "lastTimeRewardApplicable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_YVelodromePool *YVelodromePoolSession) LastTimeRewardApplicable() (*big.Int, error) {
	return _YVelodromePool.Contract.LastTimeRewardApplicable(&_YVelodromePool.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_YVelodromePool *YVelodromePoolCallerSession) LastTimeRewardApplicable() (*big.Int, error) {
	return _YVelodromePool.Contract.LastTimeRewardApplicable(&_YVelodromePool.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_YVelodromePool *YVelodromePoolCaller) LastUpdateTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YVelodromePool.contract.Call(opts, &out, "lastUpdateTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_YVelodromePool *YVelodromePoolSession) LastUpdateTime() (*big.Int, error) {
	return _YVelodromePool.Contract.LastUpdateTime(&_YVelodromePool.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_YVelodromePool *YVelodromePoolCallerSession) LastUpdateTime() (*big.Int, error) {
	return _YVelodromePool.Contract.LastUpdateTime(&_YVelodromePool.CallOpts)
}

// Left is a free data retrieval call binding the contract method 0x16e64048.
//
// Solidity: function left() view returns(uint256)
func (_YVelodromePool *YVelodromePoolCaller) Left(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YVelodromePool.contract.Call(opts, &out, "left")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Left is a free data retrieval call binding the contract method 0x16e64048.
//
// Solidity: function left() view returns(uint256)
func (_YVelodromePool *YVelodromePoolSession) Left() (*big.Int, error) {
	return _YVelodromePool.Contract.Left(&_YVelodromePool.CallOpts)
}

// Left is a free data retrieval call binding the contract method 0x16e64048.
//
// Solidity: function left() view returns(uint256)
func (_YVelodromePool *YVelodromePoolCallerSession) Left() (*big.Int, error) {
	return _YVelodromePool.Contract.Left(&_YVelodromePool.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_YVelodromePool *YVelodromePoolCaller) PeriodFinish(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YVelodromePool.contract.Call(opts, &out, "periodFinish")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_YVelodromePool *YVelodromePoolSession) PeriodFinish() (*big.Int, error) {
	return _YVelodromePool.Contract.PeriodFinish(&_YVelodromePool.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_YVelodromePool *YVelodromePoolCallerSession) PeriodFinish() (*big.Int, error) {
	return _YVelodromePool.Contract.PeriodFinish(&_YVelodromePool.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_YVelodromePool *YVelodromePoolCaller) RewardPerToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YVelodromePool.contract.Call(opts, &out, "rewardPerToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_YVelodromePool *YVelodromePoolSession) RewardPerToken() (*big.Int, error) {
	return _YVelodromePool.Contract.RewardPerToken(&_YVelodromePool.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_YVelodromePool *YVelodromePoolCallerSession) RewardPerToken() (*big.Int, error) {
	return _YVelodromePool.Contract.RewardPerToken(&_YVelodromePool.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_YVelodromePool *YVelodromePoolCaller) RewardPerTokenStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YVelodromePool.contract.Call(opts, &out, "rewardPerTokenStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_YVelodromePool *YVelodromePoolSession) RewardPerTokenStored() (*big.Int, error) {
	return _YVelodromePool.Contract.RewardPerTokenStored(&_YVelodromePool.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_YVelodromePool *YVelodromePoolCallerSession) RewardPerTokenStored() (*big.Int, error) {
	return _YVelodromePool.Contract.RewardPerTokenStored(&_YVelodromePool.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_YVelodromePool *YVelodromePoolCaller) RewardRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YVelodromePool.contract.Call(opts, &out, "rewardRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_YVelodromePool *YVelodromePoolSession) RewardRate() (*big.Int, error) {
	return _YVelodromePool.Contract.RewardRate(&_YVelodromePool.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_YVelodromePool *YVelodromePoolCallerSession) RewardRate() (*big.Int, error) {
	return _YVelodromePool.Contract.RewardRate(&_YVelodromePool.CallOpts)
}

// RewardRateByEpoch is a free data retrieval call binding the contract method 0x94af5b63.
//
// Solidity: function rewardRateByEpoch(uint256 ) view returns(uint256)
func (_YVelodromePool *YVelodromePoolCaller) RewardRateByEpoch(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YVelodromePool.contract.Call(opts, &out, "rewardRateByEpoch", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRateByEpoch is a free data retrieval call binding the contract method 0x94af5b63.
//
// Solidity: function rewardRateByEpoch(uint256 ) view returns(uint256)
func (_YVelodromePool *YVelodromePoolSession) RewardRateByEpoch(arg0 *big.Int) (*big.Int, error) {
	return _YVelodromePool.Contract.RewardRateByEpoch(&_YVelodromePool.CallOpts, arg0)
}

// RewardRateByEpoch is a free data retrieval call binding the contract method 0x94af5b63.
//
// Solidity: function rewardRateByEpoch(uint256 ) view returns(uint256)
func (_YVelodromePool *YVelodromePoolCallerSession) RewardRateByEpoch(arg0 *big.Int) (*big.Int, error) {
	return _YVelodromePool.Contract.RewardRateByEpoch(&_YVelodromePool.CallOpts, arg0)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_YVelodromePool *YVelodromePoolCaller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YVelodromePool.contract.Call(opts, &out, "rewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_YVelodromePool *YVelodromePoolSession) RewardToken() (common.Address, error) {
	return _YVelodromePool.Contract.RewardToken(&_YVelodromePool.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_YVelodromePool *YVelodromePoolCallerSession) RewardToken() (common.Address, error) {
	return _YVelodromePool.Contract.RewardToken(&_YVelodromePool.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_YVelodromePool *YVelodromePoolCaller) Rewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YVelodromePool.contract.Call(opts, &out, "rewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_YVelodromePool *YVelodromePoolSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _YVelodromePool.Contract.Rewards(&_YVelodromePool.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_YVelodromePool *YVelodromePoolCallerSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _YVelodromePool.Contract.Rewards(&_YVelodromePool.CallOpts, arg0)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_YVelodromePool *YVelodromePoolCaller) StakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YVelodromePool.contract.Call(opts, &out, "stakingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_YVelodromePool *YVelodromePoolSession) StakingToken() (common.Address, error) {
	return _YVelodromePool.Contract.StakingToken(&_YVelodromePool.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_YVelodromePool *YVelodromePoolCallerSession) StakingToken() (common.Address, error) {
	return _YVelodromePool.Contract.StakingToken(&_YVelodromePool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_YVelodromePool *YVelodromePoolCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YVelodromePool.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_YVelodromePool *YVelodromePoolSession) TotalSupply() (*big.Int, error) {
	return _YVelodromePool.Contract.TotalSupply(&_YVelodromePool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_YVelodromePool *YVelodromePoolCallerSession) TotalSupply() (*big.Int, error) {
	return _YVelodromePool.Contract.TotalSupply(&_YVelodromePool.CallOpts)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_YVelodromePool *YVelodromePoolCaller) UserRewardPerTokenPaid(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YVelodromePool.contract.Call(opts, &out, "userRewardPerTokenPaid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_YVelodromePool *YVelodromePoolSession) UserRewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _YVelodromePool.Contract.UserRewardPerTokenPaid(&_YVelodromePool.CallOpts, arg0)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_YVelodromePool *YVelodromePoolCallerSession) UserRewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _YVelodromePool.Contract.UserRewardPerTokenPaid(&_YVelodromePool.CallOpts, arg0)
}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_YVelodromePool *YVelodromePoolCaller) Voter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YVelodromePool.contract.Call(opts, &out, "voter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_YVelodromePool *YVelodromePoolSession) Voter() (common.Address, error) {
	return _YVelodromePool.Contract.Voter(&_YVelodromePool.CallOpts)
}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_YVelodromePool *YVelodromePoolCallerSession) Voter() (common.Address, error) {
	return _YVelodromePool.Contract.Voter(&_YVelodromePool.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _amount, address _recipient) returns()
func (_YVelodromePool *YVelodromePoolTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _YVelodromePool.contract.Transact(opts, "deposit", _amount, _recipient)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _amount, address _recipient) returns()
func (_YVelodromePool *YVelodromePoolSession) Deposit(_amount *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _YVelodromePool.Contract.Deposit(&_YVelodromePool.TransactOpts, _amount, _recipient)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _amount, address _recipient) returns()
func (_YVelodromePool *YVelodromePoolTransactorSession) Deposit(_amount *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _YVelodromePool.Contract.Deposit(&_YVelodromePool.TransactOpts, _amount, _recipient)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_YVelodromePool *YVelodromePoolTransactor) Deposit0(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _YVelodromePool.contract.Transact(opts, "deposit0", _amount)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_YVelodromePool *YVelodromePoolSession) Deposit0(_amount *big.Int) (*types.Transaction, error) {
	return _YVelodromePool.Contract.Deposit0(&_YVelodromePool.TransactOpts, _amount)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_YVelodromePool *YVelodromePoolTransactorSession) Deposit0(_amount *big.Int) (*types.Transaction, error) {
	return _YVelodromePool.Contract.Deposit0(&_YVelodromePool.TransactOpts, _amount)
}

// GetReward is a paid mutator transaction binding the contract method 0xc00007b0.
//
// Solidity: function getReward(address _account) returns()
func (_YVelodromePool *YVelodromePoolTransactor) GetReward(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _YVelodromePool.contract.Transact(opts, "getReward", _account)
}

// GetReward is a paid mutator transaction binding the contract method 0xc00007b0.
//
// Solidity: function getReward(address _account) returns()
func (_YVelodromePool *YVelodromePoolSession) GetReward(_account common.Address) (*types.Transaction, error) {
	return _YVelodromePool.Contract.GetReward(&_YVelodromePool.TransactOpts, _account)
}

// GetReward is a paid mutator transaction binding the contract method 0xc00007b0.
//
// Solidity: function getReward(address _account) returns()
func (_YVelodromePool *YVelodromePoolTransactorSession) GetReward(_account common.Address) (*types.Transaction, error) {
	return _YVelodromePool.Contract.GetReward(&_YVelodromePool.TransactOpts, _account)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 _amount) returns()
func (_YVelodromePool *YVelodromePoolTransactor) NotifyRewardAmount(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _YVelodromePool.contract.Transact(opts, "notifyRewardAmount", _amount)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 _amount) returns()
func (_YVelodromePool *YVelodromePoolSession) NotifyRewardAmount(_amount *big.Int) (*types.Transaction, error) {
	return _YVelodromePool.Contract.NotifyRewardAmount(&_YVelodromePool.TransactOpts, _amount)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 _amount) returns()
func (_YVelodromePool *YVelodromePoolTransactorSession) NotifyRewardAmount(_amount *big.Int) (*types.Transaction, error) {
	return _YVelodromePool.Contract.NotifyRewardAmount(&_YVelodromePool.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_YVelodromePool *YVelodromePoolTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _YVelodromePool.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_YVelodromePool *YVelodromePoolSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _YVelodromePool.Contract.Withdraw(&_YVelodromePool.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_YVelodromePool *YVelodromePoolTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _YVelodromePool.Contract.Withdraw(&_YVelodromePool.TransactOpts, _amount)
}

// YVelodromePoolClaimFeesIterator is returned from FilterClaimFees and is used to iterate over the raw logs and unpacked data for ClaimFees events raised by the YVelodromePool contract.
type YVelodromePoolClaimFeesIterator struct {
	Event *YVelodromePoolClaimFees // Event containing the contract specifics and raw log

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
func (it *YVelodromePoolClaimFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YVelodromePoolClaimFees)
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
		it.Event = new(YVelodromePoolClaimFees)
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
func (it *YVelodromePoolClaimFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YVelodromePoolClaimFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YVelodromePoolClaimFees represents a ClaimFees event raised by the YVelodromePool contract.
type YVelodromePoolClaimFees struct {
	From     common.Address
	Claimed0 *big.Int
	Claimed1 *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterClaimFees is a free log retrieval operation binding the contract event 0xbc567d6cbad26368064baa0ab5a757be46aae4d70f707f9203d9d9b6c8ccbfa3.
//
// Solidity: event ClaimFees(address indexed from, uint256 claimed0, uint256 claimed1)
func (_YVelodromePool *YVelodromePoolFilterer) FilterClaimFees(opts *bind.FilterOpts, from []common.Address) (*YVelodromePoolClaimFeesIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _YVelodromePool.contract.FilterLogs(opts, "ClaimFees", fromRule)
	if err != nil {
		return nil, err
	}
	return &YVelodromePoolClaimFeesIterator{contract: _YVelodromePool.contract, event: "ClaimFees", logs: logs, sub: sub}, nil
}

// WatchClaimFees is a free log subscription operation binding the contract event 0xbc567d6cbad26368064baa0ab5a757be46aae4d70f707f9203d9d9b6c8ccbfa3.
//
// Solidity: event ClaimFees(address indexed from, uint256 claimed0, uint256 claimed1)
func (_YVelodromePool *YVelodromePoolFilterer) WatchClaimFees(opts *bind.WatchOpts, sink chan<- *YVelodromePoolClaimFees, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _YVelodromePool.contract.WatchLogs(opts, "ClaimFees", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YVelodromePoolClaimFees)
				if err := _YVelodromePool.contract.UnpackLog(event, "ClaimFees", log); err != nil {
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

// ParseClaimFees is a log parse operation binding the contract event 0xbc567d6cbad26368064baa0ab5a757be46aae4d70f707f9203d9d9b6c8ccbfa3.
//
// Solidity: event ClaimFees(address indexed from, uint256 claimed0, uint256 claimed1)
func (_YVelodromePool *YVelodromePoolFilterer) ParseClaimFees(log types.Log) (*YVelodromePoolClaimFees, error) {
	event := new(YVelodromePoolClaimFees)
	if err := _YVelodromePool.contract.UnpackLog(event, "ClaimFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YVelodromePoolClaimRewardsIterator is returned from FilterClaimRewards and is used to iterate over the raw logs and unpacked data for ClaimRewards events raised by the YVelodromePool contract.
type YVelodromePoolClaimRewardsIterator struct {
	Event *YVelodromePoolClaimRewards // Event containing the contract specifics and raw log

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
func (it *YVelodromePoolClaimRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YVelodromePoolClaimRewards)
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
		it.Event = new(YVelodromePoolClaimRewards)
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
func (it *YVelodromePoolClaimRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YVelodromePoolClaimRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YVelodromePoolClaimRewards represents a ClaimRewards event raised by the YVelodromePool contract.
type YVelodromePoolClaimRewards struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaimRewards is a free log retrieval operation binding the contract event 0x1f89f96333d3133000ee447473151fa9606543368f02271c9d95ae14f13bcc67.
//
// Solidity: event ClaimRewards(address indexed from, uint256 amount)
func (_YVelodromePool *YVelodromePoolFilterer) FilterClaimRewards(opts *bind.FilterOpts, from []common.Address) (*YVelodromePoolClaimRewardsIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _YVelodromePool.contract.FilterLogs(opts, "ClaimRewards", fromRule)
	if err != nil {
		return nil, err
	}
	return &YVelodromePoolClaimRewardsIterator{contract: _YVelodromePool.contract, event: "ClaimRewards", logs: logs, sub: sub}, nil
}

// WatchClaimRewards is a free log subscription operation binding the contract event 0x1f89f96333d3133000ee447473151fa9606543368f02271c9d95ae14f13bcc67.
//
// Solidity: event ClaimRewards(address indexed from, uint256 amount)
func (_YVelodromePool *YVelodromePoolFilterer) WatchClaimRewards(opts *bind.WatchOpts, sink chan<- *YVelodromePoolClaimRewards, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _YVelodromePool.contract.WatchLogs(opts, "ClaimRewards", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YVelodromePoolClaimRewards)
				if err := _YVelodromePool.contract.UnpackLog(event, "ClaimRewards", log); err != nil {
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

// ParseClaimRewards is a log parse operation binding the contract event 0x1f89f96333d3133000ee447473151fa9606543368f02271c9d95ae14f13bcc67.
//
// Solidity: event ClaimRewards(address indexed from, uint256 amount)
func (_YVelodromePool *YVelodromePoolFilterer) ParseClaimRewards(log types.Log) (*YVelodromePoolClaimRewards, error) {
	event := new(YVelodromePoolClaimRewards)
	if err := _YVelodromePool.contract.UnpackLog(event, "ClaimRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YVelodromePoolDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the YVelodromePool contract.
type YVelodromePoolDepositIterator struct {
	Event *YVelodromePoolDeposit // Event containing the contract specifics and raw log

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
func (it *YVelodromePoolDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YVelodromePoolDeposit)
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
		it.Event = new(YVelodromePoolDeposit)
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
func (it *YVelodromePoolDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YVelodromePoolDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YVelodromePoolDeposit represents a Deposit event raised by the YVelodromePool contract.
type YVelodromePoolDeposit struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed from, address indexed to, uint256 amount)
func (_YVelodromePool *YVelodromePoolFilterer) FilterDeposit(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*YVelodromePoolDepositIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _YVelodromePool.contract.FilterLogs(opts, "Deposit", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &YVelodromePoolDepositIterator{contract: _YVelodromePool.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed from, address indexed to, uint256 amount)
func (_YVelodromePool *YVelodromePoolFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *YVelodromePoolDeposit, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _YVelodromePool.contract.WatchLogs(opts, "Deposit", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YVelodromePoolDeposit)
				if err := _YVelodromePool.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed from, address indexed to, uint256 amount)
func (_YVelodromePool *YVelodromePoolFilterer) ParseDeposit(log types.Log) (*YVelodromePoolDeposit, error) {
	event := new(YVelodromePoolDeposit)
	if err := _YVelodromePool.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YVelodromePoolNotifyRewardIterator is returned from FilterNotifyReward and is used to iterate over the raw logs and unpacked data for NotifyReward events raised by the YVelodromePool contract.
type YVelodromePoolNotifyRewardIterator struct {
	Event *YVelodromePoolNotifyReward // Event containing the contract specifics and raw log

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
func (it *YVelodromePoolNotifyRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YVelodromePoolNotifyReward)
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
		it.Event = new(YVelodromePoolNotifyReward)
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
func (it *YVelodromePoolNotifyRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YVelodromePoolNotifyRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YVelodromePoolNotifyReward represents a NotifyReward event raised by the YVelodromePool contract.
type YVelodromePoolNotifyReward struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNotifyReward is a free log retrieval operation binding the contract event 0x095667752957714306e1a6ad83495404412df6fdb932fca6dc849a7ee910d4c1.
//
// Solidity: event NotifyReward(address indexed from, uint256 amount)
func (_YVelodromePool *YVelodromePoolFilterer) FilterNotifyReward(opts *bind.FilterOpts, from []common.Address) (*YVelodromePoolNotifyRewardIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _YVelodromePool.contract.FilterLogs(opts, "NotifyReward", fromRule)
	if err != nil {
		return nil, err
	}
	return &YVelodromePoolNotifyRewardIterator{contract: _YVelodromePool.contract, event: "NotifyReward", logs: logs, sub: sub}, nil
}

// WatchNotifyReward is a free log subscription operation binding the contract event 0x095667752957714306e1a6ad83495404412df6fdb932fca6dc849a7ee910d4c1.
//
// Solidity: event NotifyReward(address indexed from, uint256 amount)
func (_YVelodromePool *YVelodromePoolFilterer) WatchNotifyReward(opts *bind.WatchOpts, sink chan<- *YVelodromePoolNotifyReward, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _YVelodromePool.contract.WatchLogs(opts, "NotifyReward", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YVelodromePoolNotifyReward)
				if err := _YVelodromePool.contract.UnpackLog(event, "NotifyReward", log); err != nil {
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

// ParseNotifyReward is a log parse operation binding the contract event 0x095667752957714306e1a6ad83495404412df6fdb932fca6dc849a7ee910d4c1.
//
// Solidity: event NotifyReward(address indexed from, uint256 amount)
func (_YVelodromePool *YVelodromePoolFilterer) ParseNotifyReward(log types.Log) (*YVelodromePoolNotifyReward, error) {
	event := new(YVelodromePoolNotifyReward)
	if err := _YVelodromePool.contract.UnpackLog(event, "NotifyReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YVelodromePoolWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the YVelodromePool contract.
type YVelodromePoolWithdrawIterator struct {
	Event *YVelodromePoolWithdraw // Event containing the contract specifics and raw log

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
func (it *YVelodromePoolWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YVelodromePoolWithdraw)
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
		it.Event = new(YVelodromePoolWithdraw)
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
func (it *YVelodromePoolWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YVelodromePoolWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YVelodromePoolWithdraw represents a Withdraw event raised by the YVelodromePool contract.
type YVelodromePoolWithdraw struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed from, uint256 amount)
func (_YVelodromePool *YVelodromePoolFilterer) FilterWithdraw(opts *bind.FilterOpts, from []common.Address) (*YVelodromePoolWithdrawIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _YVelodromePool.contract.FilterLogs(opts, "Withdraw", fromRule)
	if err != nil {
		return nil, err
	}
	return &YVelodromePoolWithdrawIterator{contract: _YVelodromePool.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed from, uint256 amount)
func (_YVelodromePool *YVelodromePoolFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *YVelodromePoolWithdraw, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _YVelodromePool.contract.WatchLogs(opts, "Withdraw", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YVelodromePoolWithdraw)
				if err := _YVelodromePool.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed from, uint256 amount)
func (_YVelodromePool *YVelodromePoolFilterer) ParseWithdraw(log types.Log) (*YVelodromePoolWithdraw, error) {
	event := new(YVelodromePoolWithdraw)
	if err := _YVelodromePool.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
