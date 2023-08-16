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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardsDistribution\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakingToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_zapContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"}],\"name\":\"PauseChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Recovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDuration\",\"type\":\"uint256\"}],\"name\":\"RewardsDurationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakedFor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_zapContract\",\"type\":\"address\"}],\"name\":\"ZapContractUpdated\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"earned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"exit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getReward\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRewardForDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isRetired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastPauseTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastTimeRewardApplicable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastUpdateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"notifyRewardAmount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"periodFinish\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"recoverERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardPerTokenStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardsDistribution\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardsDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardsToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_paused\",\"type\":\"bool\"}],\"name\":\"setPaused\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardsDistribution\",\"type\":\"address\"}],\"name\":\"setRewardsDistribution\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardsDuration\",\"type\":\"uint256\"}],\"name\":\"setRewardsDuration\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_zapContract\",\"type\":\"address\"}],\"name\":\"setZapContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stakeFor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakingToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRewardPerTokenPaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"zapContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YOptimismStakingReward.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _YOptimismStakingReward.Contract.BalanceOf(&_YOptimismStakingReward.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _YOptimismStakingReward.Contract.BalanceOf(&_YOptimismStakingReward.CallOpts, account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardCaller) Earned(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YOptimismStakingReward.contract.Call(opts, &out, "earned", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardSession) Earned(account common.Address) (*big.Int, error) {
	return _YOptimismStakingReward.Contract.Earned(&_YOptimismStakingReward.CallOpts, account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardCallerSession) Earned(account common.Address) (*big.Int, error) {
	return _YOptimismStakingReward.Contract.Earned(&_YOptimismStakingReward.CallOpts, account)
}

// GetRewardForDuration is a free data retrieval call binding the contract method 0x1c1f78eb.
//
// Solidity: function getRewardForDuration() view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardCaller) GetRewardForDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YOptimismStakingReward.contract.Call(opts, &out, "getRewardForDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardForDuration is a free data retrieval call binding the contract method 0x1c1f78eb.
//
// Solidity: function getRewardForDuration() view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardSession) GetRewardForDuration() (*big.Int, error) {
	return _YOptimismStakingReward.Contract.GetRewardForDuration(&_YOptimismStakingReward.CallOpts)
}

// GetRewardForDuration is a free data retrieval call binding the contract method 0x1c1f78eb.
//
// Solidity: function getRewardForDuration() view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardCallerSession) GetRewardForDuration() (*big.Int, error) {
	return _YOptimismStakingReward.Contract.GetRewardForDuration(&_YOptimismStakingReward.CallOpts)
}

// IsRetired is a free data retrieval call binding the contract method 0x49343cee.
//
// Solidity: function isRetired() view returns(bool)
func (_YOptimismStakingReward *YOptimismStakingRewardCaller) IsRetired(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _YOptimismStakingReward.contract.Call(opts, &out, "isRetired")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRetired is a free data retrieval call binding the contract method 0x49343cee.
//
// Solidity: function isRetired() view returns(bool)
func (_YOptimismStakingReward *YOptimismStakingRewardSession) IsRetired() (bool, error) {
	return _YOptimismStakingReward.Contract.IsRetired(&_YOptimismStakingReward.CallOpts)
}

// IsRetired is a free data retrieval call binding the contract method 0x49343cee.
//
// Solidity: function isRetired() view returns(bool)
func (_YOptimismStakingReward *YOptimismStakingRewardCallerSession) IsRetired() (bool, error) {
	return _YOptimismStakingReward.Contract.IsRetired(&_YOptimismStakingReward.CallOpts)
}

// LastPauseTime is a free data retrieval call binding the contract method 0x91b4ded9.
//
// Solidity: function lastPauseTime() view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardCaller) LastPauseTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YOptimismStakingReward.contract.Call(opts, &out, "lastPauseTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastPauseTime is a free data retrieval call binding the contract method 0x91b4ded9.
//
// Solidity: function lastPauseTime() view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardSession) LastPauseTime() (*big.Int, error) {
	return _YOptimismStakingReward.Contract.LastPauseTime(&_YOptimismStakingReward.CallOpts)
}

// LastPauseTime is a free data retrieval call binding the contract method 0x91b4ded9.
//
// Solidity: function lastPauseTime() view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardCallerSession) LastPauseTime() (*big.Int, error) {
	return _YOptimismStakingReward.Contract.LastPauseTime(&_YOptimismStakingReward.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardCaller) LastTimeRewardApplicable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YOptimismStakingReward.contract.Call(opts, &out, "lastTimeRewardApplicable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardSession) LastTimeRewardApplicable() (*big.Int, error) {
	return _YOptimismStakingReward.Contract.LastTimeRewardApplicable(&_YOptimismStakingReward.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardCallerSession) LastTimeRewardApplicable() (*big.Int, error) {
	return _YOptimismStakingReward.Contract.LastTimeRewardApplicable(&_YOptimismStakingReward.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardCaller) LastUpdateTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YOptimismStakingReward.contract.Call(opts, &out, "lastUpdateTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardSession) LastUpdateTime() (*big.Int, error) {
	return _YOptimismStakingReward.Contract.LastUpdateTime(&_YOptimismStakingReward.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardCallerSession) LastUpdateTime() (*big.Int, error) {
	return _YOptimismStakingReward.Contract.LastUpdateTime(&_YOptimismStakingReward.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_YOptimismStakingReward *YOptimismStakingRewardCaller) NominatedOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YOptimismStakingReward.contract.Call(opts, &out, "nominatedOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_YOptimismStakingReward *YOptimismStakingRewardSession) NominatedOwner() (common.Address, error) {
	return _YOptimismStakingReward.Contract.NominatedOwner(&_YOptimismStakingReward.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_YOptimismStakingReward *YOptimismStakingRewardCallerSession) NominatedOwner() (common.Address, error) {
	return _YOptimismStakingReward.Contract.NominatedOwner(&_YOptimismStakingReward.CallOpts)
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

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_YOptimismStakingReward *YOptimismStakingRewardCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _YOptimismStakingReward.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_YOptimismStakingReward *YOptimismStakingRewardSession) Paused() (bool, error) {
	return _YOptimismStakingReward.Contract.Paused(&_YOptimismStakingReward.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_YOptimismStakingReward *YOptimismStakingRewardCallerSession) Paused() (bool, error) {
	return _YOptimismStakingReward.Contract.Paused(&_YOptimismStakingReward.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardCaller) PeriodFinish(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YOptimismStakingReward.contract.Call(opts, &out, "periodFinish")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardSession) PeriodFinish() (*big.Int, error) {
	return _YOptimismStakingReward.Contract.PeriodFinish(&_YOptimismStakingReward.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardCallerSession) PeriodFinish() (*big.Int, error) {
	return _YOptimismStakingReward.Contract.PeriodFinish(&_YOptimismStakingReward.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardCaller) RewardPerToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YOptimismStakingReward.contract.Call(opts, &out, "rewardPerToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardSession) RewardPerToken() (*big.Int, error) {
	return _YOptimismStakingReward.Contract.RewardPerToken(&_YOptimismStakingReward.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardCallerSession) RewardPerToken() (*big.Int, error) {
	return _YOptimismStakingReward.Contract.RewardPerToken(&_YOptimismStakingReward.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardCaller) RewardPerTokenStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YOptimismStakingReward.contract.Call(opts, &out, "rewardPerTokenStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardSession) RewardPerTokenStored() (*big.Int, error) {
	return _YOptimismStakingReward.Contract.RewardPerTokenStored(&_YOptimismStakingReward.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardCallerSession) RewardPerTokenStored() (*big.Int, error) {
	return _YOptimismStakingReward.Contract.RewardPerTokenStored(&_YOptimismStakingReward.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardCaller) RewardRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YOptimismStakingReward.contract.Call(opts, &out, "rewardRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardSession) RewardRate() (*big.Int, error) {
	return _YOptimismStakingReward.Contract.RewardRate(&_YOptimismStakingReward.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardCallerSession) RewardRate() (*big.Int, error) {
	return _YOptimismStakingReward.Contract.RewardRate(&_YOptimismStakingReward.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardCaller) Rewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YOptimismStakingReward.contract.Call(opts, &out, "rewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _YOptimismStakingReward.Contract.Rewards(&_YOptimismStakingReward.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardCallerSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _YOptimismStakingReward.Contract.Rewards(&_YOptimismStakingReward.CallOpts, arg0)
}

// RewardsDistribution is a free data retrieval call binding the contract method 0x3fc6df6e.
//
// Solidity: function rewardsDistribution() view returns(address)
func (_YOptimismStakingReward *YOptimismStakingRewardCaller) RewardsDistribution(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YOptimismStakingReward.contract.Call(opts, &out, "rewardsDistribution")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsDistribution is a free data retrieval call binding the contract method 0x3fc6df6e.
//
// Solidity: function rewardsDistribution() view returns(address)
func (_YOptimismStakingReward *YOptimismStakingRewardSession) RewardsDistribution() (common.Address, error) {
	return _YOptimismStakingReward.Contract.RewardsDistribution(&_YOptimismStakingReward.CallOpts)
}

// RewardsDistribution is a free data retrieval call binding the contract method 0x3fc6df6e.
//
// Solidity: function rewardsDistribution() view returns(address)
func (_YOptimismStakingReward *YOptimismStakingRewardCallerSession) RewardsDistribution() (common.Address, error) {
	return _YOptimismStakingReward.Contract.RewardsDistribution(&_YOptimismStakingReward.CallOpts)
}

// RewardsDuration is a free data retrieval call binding the contract method 0x386a9525.
//
// Solidity: function rewardsDuration() view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardCaller) RewardsDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YOptimismStakingReward.contract.Call(opts, &out, "rewardsDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardsDuration is a free data retrieval call binding the contract method 0x386a9525.
//
// Solidity: function rewardsDuration() view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardSession) RewardsDuration() (*big.Int, error) {
	return _YOptimismStakingReward.Contract.RewardsDuration(&_YOptimismStakingReward.CallOpts)
}

// RewardsDuration is a free data retrieval call binding the contract method 0x386a9525.
//
// Solidity: function rewardsDuration() view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardCallerSession) RewardsDuration() (*big.Int, error) {
	return _YOptimismStakingReward.Contract.RewardsDuration(&_YOptimismStakingReward.CallOpts)
}

// RewardsToken is a free data retrieval call binding the contract method 0xd1af0c7d.
//
// Solidity: function rewardsToken() view returns(address)
func (_YOptimismStakingReward *YOptimismStakingRewardCaller) RewardsToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YOptimismStakingReward.contract.Call(opts, &out, "rewardsToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsToken is a free data retrieval call binding the contract method 0xd1af0c7d.
//
// Solidity: function rewardsToken() view returns(address)
func (_YOptimismStakingReward *YOptimismStakingRewardSession) RewardsToken() (common.Address, error) {
	return _YOptimismStakingReward.Contract.RewardsToken(&_YOptimismStakingReward.CallOpts)
}

// RewardsToken is a free data retrieval call binding the contract method 0xd1af0c7d.
//
// Solidity: function rewardsToken() view returns(address)
func (_YOptimismStakingReward *YOptimismStakingRewardCallerSession) RewardsToken() (common.Address, error) {
	return _YOptimismStakingReward.Contract.RewardsToken(&_YOptimismStakingReward.CallOpts)
}


// RewardsToken is a free data retrieval call binding the contract method 0xd1af0c7d.
//
// Solidity: function rewardToken() view returns(address)
func (_YOptimismStakingReward *YOptimismStakingRewardCaller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YOptimismStakingReward.contract.Call(opts, &out, "rewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardToken is a free data retrieval call binding the contract method 0xd1af0c7d.
//
// Solidity: function rewardToken() view returns(address)
func (_YOptimismStakingReward *YOptimismStakingRewardSession) RewardToken() (common.Address, error) {
	return _YOptimismStakingReward.Contract.RewardToken(&_YOptimismStakingReward.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xd1af0c7d.
//
// Solidity: function rewardToken() view returns(address)
func (_YOptimismStakingReward *YOptimismStakingRewardCallerSession) RewardToken() (common.Address, error) {
	return _YOptimismStakingReward.Contract.RewardToken(&_YOptimismStakingReward.CallOpts)
}


// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_YOptimismStakingReward *YOptimismStakingRewardCaller) StakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YOptimismStakingReward.contract.Call(opts, &out, "stakingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_YOptimismStakingReward *YOptimismStakingRewardSession) StakingToken() (common.Address, error) {
	return _YOptimismStakingReward.Contract.StakingToken(&_YOptimismStakingReward.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_YOptimismStakingReward *YOptimismStakingRewardCallerSession) StakingToken() (common.Address, error) {
	return _YOptimismStakingReward.Contract.StakingToken(&_YOptimismStakingReward.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YOptimismStakingReward.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardSession) TotalSupply() (*big.Int, error) {
	return _YOptimismStakingReward.Contract.TotalSupply(&_YOptimismStakingReward.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardCallerSession) TotalSupply() (*big.Int, error) {
	return _YOptimismStakingReward.Contract.TotalSupply(&_YOptimismStakingReward.CallOpts)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardCaller) UserRewardPerTokenPaid(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YOptimismStakingReward.contract.Call(opts, &out, "userRewardPerTokenPaid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardSession) UserRewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _YOptimismStakingReward.Contract.UserRewardPerTokenPaid(&_YOptimismStakingReward.CallOpts, arg0)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_YOptimismStakingReward *YOptimismStakingRewardCallerSession) UserRewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _YOptimismStakingReward.Contract.UserRewardPerTokenPaid(&_YOptimismStakingReward.CallOpts, arg0)
}

// ZapContract is a free data retrieval call binding the contract method 0x2549dad9.
//
// Solidity: function zapContract() view returns(address)
func (_YOptimismStakingReward *YOptimismStakingRewardCaller) ZapContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YOptimismStakingReward.contract.Call(opts, &out, "zapContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZapContract is a free data retrieval call binding the contract method 0x2549dad9.
//
// Solidity: function zapContract() view returns(address)
func (_YOptimismStakingReward *YOptimismStakingRewardSession) ZapContract() (common.Address, error) {
	return _YOptimismStakingReward.Contract.ZapContract(&_YOptimismStakingReward.CallOpts)
}

// ZapContract is a free data retrieval call binding the contract method 0x2549dad9.
//
// Solidity: function zapContract() view returns(address)
func (_YOptimismStakingReward *YOptimismStakingRewardCallerSession) ZapContract() (common.Address, error) {
	return _YOptimismStakingReward.Contract.ZapContract(&_YOptimismStakingReward.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_YOptimismStakingReward *YOptimismStakingRewardTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YOptimismStakingReward.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_YOptimismStakingReward *YOptimismStakingRewardSession) AcceptOwnership() (*types.Transaction, error) {
	return _YOptimismStakingReward.Contract.AcceptOwnership(&_YOptimismStakingReward.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_YOptimismStakingReward *YOptimismStakingRewardTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _YOptimismStakingReward.Contract.AcceptOwnership(&_YOptimismStakingReward.TransactOpts)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_YOptimismStakingReward *YOptimismStakingRewardTransactor) Exit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YOptimismStakingReward.contract.Transact(opts, "exit")
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_YOptimismStakingReward *YOptimismStakingRewardSession) Exit() (*types.Transaction, error) {
	return _YOptimismStakingReward.Contract.Exit(&_YOptimismStakingReward.TransactOpts)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_YOptimismStakingReward *YOptimismStakingRewardTransactorSession) Exit() (*types.Transaction, error) {
	return _YOptimismStakingReward.Contract.Exit(&_YOptimismStakingReward.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_YOptimismStakingReward *YOptimismStakingRewardTransactor) GetReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YOptimismStakingReward.contract.Transact(opts, "getReward")
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_YOptimismStakingReward *YOptimismStakingRewardSession) GetReward() (*types.Transaction, error) {
	return _YOptimismStakingReward.Contract.GetReward(&_YOptimismStakingReward.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_YOptimismStakingReward *YOptimismStakingRewardTransactorSession) GetReward() (*types.Transaction, error) {
	return _YOptimismStakingReward.Contract.GetReward(&_YOptimismStakingReward.TransactOpts)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_YOptimismStakingReward *YOptimismStakingRewardTransactor) NominateNewOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _YOptimismStakingReward.contract.Transact(opts, "nominateNewOwner", _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_YOptimismStakingReward *YOptimismStakingRewardSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _YOptimismStakingReward.Contract.NominateNewOwner(&_YOptimismStakingReward.TransactOpts, _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_YOptimismStakingReward *YOptimismStakingRewardTransactorSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _YOptimismStakingReward.Contract.NominateNewOwner(&_YOptimismStakingReward.TransactOpts, _owner)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 reward) returns()
func (_YOptimismStakingReward *YOptimismStakingRewardTransactor) NotifyRewardAmount(opts *bind.TransactOpts, reward *big.Int) (*types.Transaction, error) {
	return _YOptimismStakingReward.contract.Transact(opts, "notifyRewardAmount", reward)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 reward) returns()
func (_YOptimismStakingReward *YOptimismStakingRewardSession) NotifyRewardAmount(reward *big.Int) (*types.Transaction, error) {
	return _YOptimismStakingReward.Contract.NotifyRewardAmount(&_YOptimismStakingReward.TransactOpts, reward)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 reward) returns()
func (_YOptimismStakingReward *YOptimismStakingRewardTransactorSession) NotifyRewardAmount(reward *big.Int) (*types.Transaction, error) {
	return _YOptimismStakingReward.Contract.NotifyRewardAmount(&_YOptimismStakingReward.TransactOpts, reward)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address tokenAddress, uint256 tokenAmount) returns()
func (_YOptimismStakingReward *YOptimismStakingRewardTransactor) RecoverERC20(opts *bind.TransactOpts, tokenAddress common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _YOptimismStakingReward.contract.Transact(opts, "recoverERC20", tokenAddress, tokenAmount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address tokenAddress, uint256 tokenAmount) returns()
func (_YOptimismStakingReward *YOptimismStakingRewardSession) RecoverERC20(tokenAddress common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _YOptimismStakingReward.Contract.RecoverERC20(&_YOptimismStakingReward.TransactOpts, tokenAddress, tokenAmount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address tokenAddress, uint256 tokenAmount) returns()
func (_YOptimismStakingReward *YOptimismStakingRewardTransactorSession) RecoverERC20(tokenAddress common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _YOptimismStakingReward.Contract.RecoverERC20(&_YOptimismStakingReward.TransactOpts, tokenAddress, tokenAmount)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool _paused) returns()
func (_YOptimismStakingReward *YOptimismStakingRewardTransactor) SetPaused(opts *bind.TransactOpts, _paused bool) (*types.Transaction, error) {
	return _YOptimismStakingReward.contract.Transact(opts, "setPaused", _paused)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool _paused) returns()
func (_YOptimismStakingReward *YOptimismStakingRewardSession) SetPaused(_paused bool) (*types.Transaction, error) {
	return _YOptimismStakingReward.Contract.SetPaused(&_YOptimismStakingReward.TransactOpts, _paused)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool _paused) returns()
func (_YOptimismStakingReward *YOptimismStakingRewardTransactorSession) SetPaused(_paused bool) (*types.Transaction, error) {
	return _YOptimismStakingReward.Contract.SetPaused(&_YOptimismStakingReward.TransactOpts, _paused)
}

// SetRewardsDistribution is a paid mutator transaction binding the contract method 0x19762143.
//
// Solidity: function setRewardsDistribution(address _rewardsDistribution) returns()
func (_YOptimismStakingReward *YOptimismStakingRewardTransactor) SetRewardsDistribution(opts *bind.TransactOpts, _rewardsDistribution common.Address) (*types.Transaction, error) {
	return _YOptimismStakingReward.contract.Transact(opts, "setRewardsDistribution", _rewardsDistribution)
}

// SetRewardsDistribution is a paid mutator transaction binding the contract method 0x19762143.
//
// Solidity: function setRewardsDistribution(address _rewardsDistribution) returns()
func (_YOptimismStakingReward *YOptimismStakingRewardSession) SetRewardsDistribution(_rewardsDistribution common.Address) (*types.Transaction, error) {
	return _YOptimismStakingReward.Contract.SetRewardsDistribution(&_YOptimismStakingReward.TransactOpts, _rewardsDistribution)
}

// SetRewardsDistribution is a paid mutator transaction binding the contract method 0x19762143.
//
// Solidity: function setRewardsDistribution(address _rewardsDistribution) returns()
func (_YOptimismStakingReward *YOptimismStakingRewardTransactorSession) SetRewardsDistribution(_rewardsDistribution common.Address) (*types.Transaction, error) {
	return _YOptimismStakingReward.Contract.SetRewardsDistribution(&_YOptimismStakingReward.TransactOpts, _rewardsDistribution)
}

// SetRewardsDuration is a paid mutator transaction binding the contract method 0xcc1a378f.
//
// Solidity: function setRewardsDuration(uint256 _rewardsDuration) returns()
func (_YOptimismStakingReward *YOptimismStakingRewardTransactor) SetRewardsDuration(opts *bind.TransactOpts, _rewardsDuration *big.Int) (*types.Transaction, error) {
	return _YOptimismStakingReward.contract.Transact(opts, "setRewardsDuration", _rewardsDuration)
}

// SetRewardsDuration is a paid mutator transaction binding the contract method 0xcc1a378f.
//
// Solidity: function setRewardsDuration(uint256 _rewardsDuration) returns()
func (_YOptimismStakingReward *YOptimismStakingRewardSession) SetRewardsDuration(_rewardsDuration *big.Int) (*types.Transaction, error) {
	return _YOptimismStakingReward.Contract.SetRewardsDuration(&_YOptimismStakingReward.TransactOpts, _rewardsDuration)
}

// SetRewardsDuration is a paid mutator transaction binding the contract method 0xcc1a378f.
//
// Solidity: function setRewardsDuration(uint256 _rewardsDuration) returns()
func (_YOptimismStakingReward *YOptimismStakingRewardTransactorSession) SetRewardsDuration(_rewardsDuration *big.Int) (*types.Transaction, error) {
	return _YOptimismStakingReward.Contract.SetRewardsDuration(&_YOptimismStakingReward.TransactOpts, _rewardsDuration)
}

// SetZapContract is a paid mutator transaction binding the contract method 0xfcaa8737.
//
// Solidity: function setZapContract(address _zapContract) returns()
func (_YOptimismStakingReward *YOptimismStakingRewardTransactor) SetZapContract(opts *bind.TransactOpts, _zapContract common.Address) (*types.Transaction, error) {
	return _YOptimismStakingReward.contract.Transact(opts, "setZapContract", _zapContract)
}

// SetZapContract is a paid mutator transaction binding the contract method 0xfcaa8737.
//
// Solidity: function setZapContract(address _zapContract) returns()
func (_YOptimismStakingReward *YOptimismStakingRewardSession) SetZapContract(_zapContract common.Address) (*types.Transaction, error) {
	return _YOptimismStakingReward.Contract.SetZapContract(&_YOptimismStakingReward.TransactOpts, _zapContract)
}

// SetZapContract is a paid mutator transaction binding the contract method 0xfcaa8737.
//
// Solidity: function setZapContract(address _zapContract) returns()
func (_YOptimismStakingReward *YOptimismStakingRewardTransactorSession) SetZapContract(_zapContract common.Address) (*types.Transaction, error) {
	return _YOptimismStakingReward.Contract.SetZapContract(&_YOptimismStakingReward.TransactOpts, _zapContract)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_YOptimismStakingReward *YOptimismStakingRewardTransactor) Stake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _YOptimismStakingReward.contract.Transact(opts, "stake", amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_YOptimismStakingReward *YOptimismStakingRewardSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _YOptimismStakingReward.Contract.Stake(&_YOptimismStakingReward.TransactOpts, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_YOptimismStakingReward *YOptimismStakingRewardTransactorSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _YOptimismStakingReward.Contract.Stake(&_YOptimismStakingReward.TransactOpts, amount)
}

// StakeFor is a paid mutator transaction binding the contract method 0x2ee40908.
//
// Solidity: function stakeFor(address recipient, uint256 amount) returns()
func (_YOptimismStakingReward *YOptimismStakingRewardTransactor) StakeFor(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YOptimismStakingReward.contract.Transact(opts, "stakeFor", recipient, amount)
}

// StakeFor is a paid mutator transaction binding the contract method 0x2ee40908.
//
// Solidity: function stakeFor(address recipient, uint256 amount) returns()
func (_YOptimismStakingReward *YOptimismStakingRewardSession) StakeFor(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YOptimismStakingReward.Contract.StakeFor(&_YOptimismStakingReward.TransactOpts, recipient, amount)
}

// StakeFor is a paid mutator transaction binding the contract method 0x2ee40908.
//
// Solidity: function stakeFor(address recipient, uint256 amount) returns()
func (_YOptimismStakingReward *YOptimismStakingRewardTransactorSession) StakeFor(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YOptimismStakingReward.Contract.StakeFor(&_YOptimismStakingReward.TransactOpts, recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_YOptimismStakingReward *YOptimismStakingRewardTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _YOptimismStakingReward.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_YOptimismStakingReward *YOptimismStakingRewardSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _YOptimismStakingReward.Contract.Withdraw(&_YOptimismStakingReward.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_YOptimismStakingReward *YOptimismStakingRewardTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _YOptimismStakingReward.Contract.Withdraw(&_YOptimismStakingReward.TransactOpts, amount)
}

// YOptimismStakingRewardOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the YOptimismStakingReward contract.
type YOptimismStakingRewardOwnerChangedIterator struct {
	Event *YOptimismStakingRewardOwnerChanged // Event containing the contract specifics and raw log

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
func (it *YOptimismStakingRewardOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YOptimismStakingRewardOwnerChanged)
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
		it.Event = new(YOptimismStakingRewardOwnerChanged)
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
func (it *YOptimismStakingRewardOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YOptimismStakingRewardOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YOptimismStakingRewardOwnerChanged represents a OwnerChanged event raised by the YOptimismStakingReward contract.
type YOptimismStakingRewardOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*YOptimismStakingRewardOwnerChangedIterator, error) {

	logs, sub, err := _YOptimismStakingReward.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &YOptimismStakingRewardOwnerChangedIterator{contract: _YOptimismStakingReward.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *YOptimismStakingRewardOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _YOptimismStakingReward.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YOptimismStakingRewardOwnerChanged)
				if err := _YOptimismStakingReward.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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

// ParseOwnerChanged is a log parse operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) ParseOwnerChanged(log types.Log) (*YOptimismStakingRewardOwnerChanged, error) {
	event := new(YOptimismStakingRewardOwnerChanged)
	if err := _YOptimismStakingReward.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YOptimismStakingRewardOwnerNominatedIterator is returned from FilterOwnerNominated and is used to iterate over the raw logs and unpacked data for OwnerNominated events raised by the YOptimismStakingReward contract.
type YOptimismStakingRewardOwnerNominatedIterator struct {
	Event *YOptimismStakingRewardOwnerNominated // Event containing the contract specifics and raw log

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
func (it *YOptimismStakingRewardOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YOptimismStakingRewardOwnerNominated)
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
		it.Event = new(YOptimismStakingRewardOwnerNominated)
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
func (it *YOptimismStakingRewardOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YOptimismStakingRewardOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YOptimismStakingRewardOwnerNominated represents a OwnerNominated event raised by the YOptimismStakingReward contract.
type YOptimismStakingRewardOwnerNominated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerNominated is a free log retrieval operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) FilterOwnerNominated(opts *bind.FilterOpts) (*YOptimismStakingRewardOwnerNominatedIterator, error) {

	logs, sub, err := _YOptimismStakingReward.contract.FilterLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return &YOptimismStakingRewardOwnerNominatedIterator{contract: _YOptimismStakingReward.contract, event: "OwnerNominated", logs: logs, sub: sub}, nil
}

// WatchOwnerNominated is a free log subscription operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) WatchOwnerNominated(opts *bind.WatchOpts, sink chan<- *YOptimismStakingRewardOwnerNominated) (event.Subscription, error) {

	logs, sub, err := _YOptimismStakingReward.contract.WatchLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YOptimismStakingRewardOwnerNominated)
				if err := _YOptimismStakingReward.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
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

// ParseOwnerNominated is a log parse operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) ParseOwnerNominated(log types.Log) (*YOptimismStakingRewardOwnerNominated, error) {
	event := new(YOptimismStakingRewardOwnerNominated)
	if err := _YOptimismStakingReward.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YOptimismStakingRewardPauseChangedIterator is returned from FilterPauseChanged and is used to iterate over the raw logs and unpacked data for PauseChanged events raised by the YOptimismStakingReward contract.
type YOptimismStakingRewardPauseChangedIterator struct {
	Event *YOptimismStakingRewardPauseChanged // Event containing the contract specifics and raw log

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
func (it *YOptimismStakingRewardPauseChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YOptimismStakingRewardPauseChanged)
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
		it.Event = new(YOptimismStakingRewardPauseChanged)
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
func (it *YOptimismStakingRewardPauseChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YOptimismStakingRewardPauseChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YOptimismStakingRewardPauseChanged represents a PauseChanged event raised by the YOptimismStakingReward contract.
type YOptimismStakingRewardPauseChanged struct {
	IsPaused bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPauseChanged is a free log retrieval operation binding the contract event 0x8fb6c181ee25a520cf3dd6565006ef91229fcfe5a989566c2a3b8c115570cec5.
//
// Solidity: event PauseChanged(bool isPaused)
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) FilterPauseChanged(opts *bind.FilterOpts) (*YOptimismStakingRewardPauseChangedIterator, error) {

	logs, sub, err := _YOptimismStakingReward.contract.FilterLogs(opts, "PauseChanged")
	if err != nil {
		return nil, err
	}
	return &YOptimismStakingRewardPauseChangedIterator{contract: _YOptimismStakingReward.contract, event: "PauseChanged", logs: logs, sub: sub}, nil
}

// WatchPauseChanged is a free log subscription operation binding the contract event 0x8fb6c181ee25a520cf3dd6565006ef91229fcfe5a989566c2a3b8c115570cec5.
//
// Solidity: event PauseChanged(bool isPaused)
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) WatchPauseChanged(opts *bind.WatchOpts, sink chan<- *YOptimismStakingRewardPauseChanged) (event.Subscription, error) {

	logs, sub, err := _YOptimismStakingReward.contract.WatchLogs(opts, "PauseChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YOptimismStakingRewardPauseChanged)
				if err := _YOptimismStakingReward.contract.UnpackLog(event, "PauseChanged", log); err != nil {
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

// ParsePauseChanged is a log parse operation binding the contract event 0x8fb6c181ee25a520cf3dd6565006ef91229fcfe5a989566c2a3b8c115570cec5.
//
// Solidity: event PauseChanged(bool isPaused)
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) ParsePauseChanged(log types.Log) (*YOptimismStakingRewardPauseChanged, error) {
	event := new(YOptimismStakingRewardPauseChanged)
	if err := _YOptimismStakingReward.contract.UnpackLog(event, "PauseChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YOptimismStakingRewardRecoveredIterator is returned from FilterRecovered and is used to iterate over the raw logs and unpacked data for Recovered events raised by the YOptimismStakingReward contract.
type YOptimismStakingRewardRecoveredIterator struct {
	Event *YOptimismStakingRewardRecovered // Event containing the contract specifics and raw log

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
func (it *YOptimismStakingRewardRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YOptimismStakingRewardRecovered)
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
		it.Event = new(YOptimismStakingRewardRecovered)
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
func (it *YOptimismStakingRewardRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YOptimismStakingRewardRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YOptimismStakingRewardRecovered represents a Recovered event raised by the YOptimismStakingReward contract.
type YOptimismStakingRewardRecovered struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRecovered is a free log retrieval operation binding the contract event 0x8c1256b8896378cd5044f80c202f9772b9d77dc85c8a6eb51967210b09bfaa28.
//
// Solidity: event Recovered(address token, uint256 amount)
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) FilterRecovered(opts *bind.FilterOpts) (*YOptimismStakingRewardRecoveredIterator, error) {

	logs, sub, err := _YOptimismStakingReward.contract.FilterLogs(opts, "Recovered")
	if err != nil {
		return nil, err
	}
	return &YOptimismStakingRewardRecoveredIterator{contract: _YOptimismStakingReward.contract, event: "Recovered", logs: logs, sub: sub}, nil
}

// WatchRecovered is a free log subscription operation binding the contract event 0x8c1256b8896378cd5044f80c202f9772b9d77dc85c8a6eb51967210b09bfaa28.
//
// Solidity: event Recovered(address token, uint256 amount)
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) WatchRecovered(opts *bind.WatchOpts, sink chan<- *YOptimismStakingRewardRecovered) (event.Subscription, error) {

	logs, sub, err := _YOptimismStakingReward.contract.WatchLogs(opts, "Recovered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YOptimismStakingRewardRecovered)
				if err := _YOptimismStakingReward.contract.UnpackLog(event, "Recovered", log); err != nil {
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

// ParseRecovered is a log parse operation binding the contract event 0x8c1256b8896378cd5044f80c202f9772b9d77dc85c8a6eb51967210b09bfaa28.
//
// Solidity: event Recovered(address token, uint256 amount)
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) ParseRecovered(log types.Log) (*YOptimismStakingRewardRecovered, error) {
	event := new(YOptimismStakingRewardRecovered)
	if err := _YOptimismStakingReward.contract.UnpackLog(event, "Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YOptimismStakingRewardRewardAddedIterator is returned from FilterRewardAdded and is used to iterate over the raw logs and unpacked data for RewardAdded events raised by the YOptimismStakingReward contract.
type YOptimismStakingRewardRewardAddedIterator struct {
	Event *YOptimismStakingRewardRewardAdded // Event containing the contract specifics and raw log

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
func (it *YOptimismStakingRewardRewardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YOptimismStakingRewardRewardAdded)
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
		it.Event = new(YOptimismStakingRewardRewardAdded)
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
func (it *YOptimismStakingRewardRewardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YOptimismStakingRewardRewardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YOptimismStakingRewardRewardAdded represents a RewardAdded event raised by the YOptimismStakingReward contract.
type YOptimismStakingRewardRewardAdded struct {
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardAdded is a free log retrieval operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) FilterRewardAdded(opts *bind.FilterOpts) (*YOptimismStakingRewardRewardAddedIterator, error) {

	logs, sub, err := _YOptimismStakingReward.contract.FilterLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return &YOptimismStakingRewardRewardAddedIterator{contract: _YOptimismStakingReward.contract, event: "RewardAdded", logs: logs, sub: sub}, nil
}

// WatchRewardAdded is a free log subscription operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) WatchRewardAdded(opts *bind.WatchOpts, sink chan<- *YOptimismStakingRewardRewardAdded) (event.Subscription, error) {

	logs, sub, err := _YOptimismStakingReward.contract.WatchLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YOptimismStakingRewardRewardAdded)
				if err := _YOptimismStakingReward.contract.UnpackLog(event, "RewardAdded", log); err != nil {
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

// ParseRewardAdded is a log parse operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) ParseRewardAdded(log types.Log) (*YOptimismStakingRewardRewardAdded, error) {
	event := new(YOptimismStakingRewardRewardAdded)
	if err := _YOptimismStakingReward.contract.UnpackLog(event, "RewardAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YOptimismStakingRewardRewardPaidIterator is returned from FilterRewardPaid and is used to iterate over the raw logs and unpacked data for RewardPaid events raised by the YOptimismStakingReward contract.
type YOptimismStakingRewardRewardPaidIterator struct {
	Event *YOptimismStakingRewardRewardPaid // Event containing the contract specifics and raw log

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
func (it *YOptimismStakingRewardRewardPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YOptimismStakingRewardRewardPaid)
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
		it.Event = new(YOptimismStakingRewardRewardPaid)
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
func (it *YOptimismStakingRewardRewardPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YOptimismStakingRewardRewardPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YOptimismStakingRewardRewardPaid represents a RewardPaid event raised by the YOptimismStakingReward contract.
type YOptimismStakingRewardRewardPaid struct {
	User   common.Address
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardPaid is a free log retrieval operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) FilterRewardPaid(opts *bind.FilterOpts, user []common.Address) (*YOptimismStakingRewardRewardPaidIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _YOptimismStakingReward.contract.FilterLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return &YOptimismStakingRewardRewardPaidIterator{contract: _YOptimismStakingReward.contract, event: "RewardPaid", logs: logs, sub: sub}, nil
}

// WatchRewardPaid is a free log subscription operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) WatchRewardPaid(opts *bind.WatchOpts, sink chan<- *YOptimismStakingRewardRewardPaid, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _YOptimismStakingReward.contract.WatchLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YOptimismStakingRewardRewardPaid)
				if err := _YOptimismStakingReward.contract.UnpackLog(event, "RewardPaid", log); err != nil {
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

// ParseRewardPaid is a log parse operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) ParseRewardPaid(log types.Log) (*YOptimismStakingRewardRewardPaid, error) {
	event := new(YOptimismStakingRewardRewardPaid)
	if err := _YOptimismStakingReward.contract.UnpackLog(event, "RewardPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YOptimismStakingRewardRewardsDurationUpdatedIterator is returned from FilterRewardsDurationUpdated and is used to iterate over the raw logs and unpacked data for RewardsDurationUpdated events raised by the YOptimismStakingReward contract.
type YOptimismStakingRewardRewardsDurationUpdatedIterator struct {
	Event *YOptimismStakingRewardRewardsDurationUpdated // Event containing the contract specifics and raw log

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
func (it *YOptimismStakingRewardRewardsDurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YOptimismStakingRewardRewardsDurationUpdated)
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
		it.Event = new(YOptimismStakingRewardRewardsDurationUpdated)
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
func (it *YOptimismStakingRewardRewardsDurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YOptimismStakingRewardRewardsDurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YOptimismStakingRewardRewardsDurationUpdated represents a RewardsDurationUpdated event raised by the YOptimismStakingReward contract.
type YOptimismStakingRewardRewardsDurationUpdated struct {
	NewDuration *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardsDurationUpdated is a free log retrieval operation binding the contract event 0xfb46ca5a5e06d4540d6387b930a7c978bce0db5f449ec6b3f5d07c6e1d44f2d3.
//
// Solidity: event RewardsDurationUpdated(uint256 newDuration)
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) FilterRewardsDurationUpdated(opts *bind.FilterOpts) (*YOptimismStakingRewardRewardsDurationUpdatedIterator, error) {

	logs, sub, err := _YOptimismStakingReward.contract.FilterLogs(opts, "RewardsDurationUpdated")
	if err != nil {
		return nil, err
	}
	return &YOptimismStakingRewardRewardsDurationUpdatedIterator{contract: _YOptimismStakingReward.contract, event: "RewardsDurationUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardsDurationUpdated is a free log subscription operation binding the contract event 0xfb46ca5a5e06d4540d6387b930a7c978bce0db5f449ec6b3f5d07c6e1d44f2d3.
//
// Solidity: event RewardsDurationUpdated(uint256 newDuration)
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) WatchRewardsDurationUpdated(opts *bind.WatchOpts, sink chan<- *YOptimismStakingRewardRewardsDurationUpdated) (event.Subscription, error) {

	logs, sub, err := _YOptimismStakingReward.contract.WatchLogs(opts, "RewardsDurationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YOptimismStakingRewardRewardsDurationUpdated)
				if err := _YOptimismStakingReward.contract.UnpackLog(event, "RewardsDurationUpdated", log); err != nil {
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

// ParseRewardsDurationUpdated is a log parse operation binding the contract event 0xfb46ca5a5e06d4540d6387b930a7c978bce0db5f449ec6b3f5d07c6e1d44f2d3.
//
// Solidity: event RewardsDurationUpdated(uint256 newDuration)
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) ParseRewardsDurationUpdated(log types.Log) (*YOptimismStakingRewardRewardsDurationUpdated, error) {
	event := new(YOptimismStakingRewardRewardsDurationUpdated)
	if err := _YOptimismStakingReward.contract.UnpackLog(event, "RewardsDurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YOptimismStakingRewardStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the YOptimismStakingReward contract.
type YOptimismStakingRewardStakedIterator struct {
	Event *YOptimismStakingRewardStaked // Event containing the contract specifics and raw log

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
func (it *YOptimismStakingRewardStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YOptimismStakingRewardStaked)
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
		it.Event = new(YOptimismStakingRewardStaked)
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
func (it *YOptimismStakingRewardStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YOptimismStakingRewardStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YOptimismStakingRewardStaked represents a Staked event raised by the YOptimismStakingReward contract.
type YOptimismStakingRewardStaked struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address) (*YOptimismStakingRewardStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _YOptimismStakingReward.contract.FilterLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return &YOptimismStakingRewardStakedIterator{contract: _YOptimismStakingReward.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *YOptimismStakingRewardStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _YOptimismStakingReward.contract.WatchLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YOptimismStakingRewardStaked)
				if err := _YOptimismStakingReward.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) ParseStaked(log types.Log) (*YOptimismStakingRewardStaked, error) {
	event := new(YOptimismStakingRewardStaked)
	if err := _YOptimismStakingReward.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YOptimismStakingRewardStakedForIterator is returned from FilterStakedFor and is used to iterate over the raw logs and unpacked data for StakedFor events raised by the YOptimismStakingReward contract.
type YOptimismStakingRewardStakedForIterator struct {
	Event *YOptimismStakingRewardStakedFor // Event containing the contract specifics and raw log

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
func (it *YOptimismStakingRewardStakedForIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YOptimismStakingRewardStakedFor)
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
		it.Event = new(YOptimismStakingRewardStakedFor)
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
func (it *YOptimismStakingRewardStakedForIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YOptimismStakingRewardStakedForIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YOptimismStakingRewardStakedFor represents a StakedFor event raised by the YOptimismStakingReward contract.
type YOptimismStakingRewardStakedFor struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakedFor is a free log retrieval operation binding the contract event 0xd185ae938da574e9cd1073962e1972c75ec585ab222b200a88c0abe2bf0cfe67.
//
// Solidity: event StakedFor(address indexed user, uint256 amount)
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) FilterStakedFor(opts *bind.FilterOpts, user []common.Address) (*YOptimismStakingRewardStakedForIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _YOptimismStakingReward.contract.FilterLogs(opts, "StakedFor", userRule)
	if err != nil {
		return nil, err
	}
	return &YOptimismStakingRewardStakedForIterator{contract: _YOptimismStakingReward.contract, event: "StakedFor", logs: logs, sub: sub}, nil
}

// WatchStakedFor is a free log subscription operation binding the contract event 0xd185ae938da574e9cd1073962e1972c75ec585ab222b200a88c0abe2bf0cfe67.
//
// Solidity: event StakedFor(address indexed user, uint256 amount)
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) WatchStakedFor(opts *bind.WatchOpts, sink chan<- *YOptimismStakingRewardStakedFor, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _YOptimismStakingReward.contract.WatchLogs(opts, "StakedFor", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YOptimismStakingRewardStakedFor)
				if err := _YOptimismStakingReward.contract.UnpackLog(event, "StakedFor", log); err != nil {
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

// ParseStakedFor is a log parse operation binding the contract event 0xd185ae938da574e9cd1073962e1972c75ec585ab222b200a88c0abe2bf0cfe67.
//
// Solidity: event StakedFor(address indexed user, uint256 amount)
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) ParseStakedFor(log types.Log) (*YOptimismStakingRewardStakedFor, error) {
	event := new(YOptimismStakingRewardStakedFor)
	if err := _YOptimismStakingReward.contract.UnpackLog(event, "StakedFor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YOptimismStakingRewardWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the YOptimismStakingReward contract.
type YOptimismStakingRewardWithdrawnIterator struct {
	Event *YOptimismStakingRewardWithdrawn // Event containing the contract specifics and raw log

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
func (it *YOptimismStakingRewardWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YOptimismStakingRewardWithdrawn)
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
		it.Event = new(YOptimismStakingRewardWithdrawn)
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
func (it *YOptimismStakingRewardWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YOptimismStakingRewardWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YOptimismStakingRewardWithdrawn represents a Withdrawn event raised by the YOptimismStakingReward contract.
type YOptimismStakingRewardWithdrawn struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address) (*YOptimismStakingRewardWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _YOptimismStakingReward.contract.FilterLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &YOptimismStakingRewardWithdrawnIterator{contract: _YOptimismStakingReward.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *YOptimismStakingRewardWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _YOptimismStakingReward.contract.WatchLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YOptimismStakingRewardWithdrawn)
				if err := _YOptimismStakingReward.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) ParseWithdrawn(log types.Log) (*YOptimismStakingRewardWithdrawn, error) {
	event := new(YOptimismStakingRewardWithdrawn)
	if err := _YOptimismStakingReward.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YOptimismStakingRewardZapContractUpdatedIterator is returned from FilterZapContractUpdated and is used to iterate over the raw logs and unpacked data for ZapContractUpdated events raised by the YOptimismStakingReward contract.
type YOptimismStakingRewardZapContractUpdatedIterator struct {
	Event *YOptimismStakingRewardZapContractUpdated // Event containing the contract specifics and raw log

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
func (it *YOptimismStakingRewardZapContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YOptimismStakingRewardZapContractUpdated)
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
		it.Event = new(YOptimismStakingRewardZapContractUpdated)
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
func (it *YOptimismStakingRewardZapContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YOptimismStakingRewardZapContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YOptimismStakingRewardZapContractUpdated represents a ZapContractUpdated event raised by the YOptimismStakingReward contract.
type YOptimismStakingRewardZapContractUpdated struct {
	ZapContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterZapContractUpdated is a free log retrieval operation binding the contract event 0x0da458581861b4ff3ae5c23a916ec5218bbf4c371497cec29006f5c55d9c9a84.
//
// Solidity: event ZapContractUpdated(address _zapContract)
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) FilterZapContractUpdated(opts *bind.FilterOpts) (*YOptimismStakingRewardZapContractUpdatedIterator, error) {

	logs, sub, err := _YOptimismStakingReward.contract.FilterLogs(opts, "ZapContractUpdated")
	if err != nil {
		return nil, err
	}
	return &YOptimismStakingRewardZapContractUpdatedIterator{contract: _YOptimismStakingReward.contract, event: "ZapContractUpdated", logs: logs, sub: sub}, nil
}

// WatchZapContractUpdated is a free log subscription operation binding the contract event 0x0da458581861b4ff3ae5c23a916ec5218bbf4c371497cec29006f5c55d9c9a84.
//
// Solidity: event ZapContractUpdated(address _zapContract)
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) WatchZapContractUpdated(opts *bind.WatchOpts, sink chan<- *YOptimismStakingRewardZapContractUpdated) (event.Subscription, error) {

	logs, sub, err := _YOptimismStakingReward.contract.WatchLogs(opts, "ZapContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YOptimismStakingRewardZapContractUpdated)
				if err := _YOptimismStakingReward.contract.UnpackLog(event, "ZapContractUpdated", log); err != nil {
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

// ParseZapContractUpdated is a log parse operation binding the contract event 0x0da458581861b4ff3ae5c23a916ec5218bbf4c371497cec29006f5c55d9c9a84.
//
// Solidity: event ZapContractUpdated(address _zapContract)
func (_YOptimismStakingReward *YOptimismStakingRewardFilterer) ParseZapContractUpdated(log types.Log) (*YOptimismStakingRewardZapContractUpdated, error) {
	event := new(YOptimismStakingRewardZapContractUpdated)
	if err := _YOptimismStakingReward.contract.UnpackLog(event, "ZapContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
