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

// JuicedStakingRewardsMetaData contains all meta data concerning the JuicedStakingRewards contract.
var JuicedStakingRewardsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakingToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_zapContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"clone\",\"type\":\"address\"}],\"name\":\"Cloned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"Ownererance\",\"type\":\"address\"}],\"name\":\"OwnerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Recovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDuration\",\"type\":\"uint256\"}],\"name\":\"RewardsDurationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakedFor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawnFor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_zapContract\",\"type\":\"address\"}],\"name\":\"ZapContractUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardsDistributor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewardsDuration\",\"type\":\"uint256\"}],\"name\":\"addReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakingToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_zapContract\",\"type\":\"address\"}],\"name\":\"cloneStakingPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newStakingPool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"}],\"name\":\"earned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pending\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"}],\"name\":\"getOneReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"}],\"name\":\"getRewardForDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakingToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_zapContract\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOriginal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isRetired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"}],\"name\":\"lastTimeRewardApplicable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewardAmount\",\"type\":\"uint256\"}],\"name\":\"notifyRewardAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenAmount\",\"type\":\"uint256\"}],\"name\":\"recoverERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardData\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"rewardsDistributor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rewardsDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"periodFinish\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdateTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardPerTokenStored\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"}],\"name\":\"rewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewardTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setPendingOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardsDistributor\",\"type\":\"address\"}],\"name\":\"setRewardsDistributor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewardsDuration\",\"type\":\"uint256\"}],\"name\":\"setRewardsDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_zapContract\",\"type\":\"address\"}],\"name\":\"setZapContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"stakeFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRewardPerTokenPaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_exit\",\"type\":\"bool\"}],\"name\":\"withdrawFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zapContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// JuicedStakingRewardsABI is the input ABI used to generate the binding from.
// Deprecated: Use JuicedStakingRewardsMetaData.ABI instead.
var JuicedStakingRewardsABI = JuicedStakingRewardsMetaData.ABI

// JuicedStakingRewards is an auto generated Go binding around an Ethereum contract.
type JuicedStakingRewards struct {
	JuicedStakingRewardsCaller     // Read-only binding to the contract
	JuicedStakingRewardsTransactor // Write-only binding to the contract
	JuicedStakingRewardsFilterer   // Log filterer for contract events
}

// JuicedStakingRewardsCaller is an auto generated read-only Go binding around an Ethereum contract.
type JuicedStakingRewardsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JuicedStakingRewardsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JuicedStakingRewardsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JuicedStakingRewardsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JuicedStakingRewardsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JuicedStakingRewardsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JuicedStakingRewardsSession struct {
	Contract     *JuicedStakingRewards // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// JuicedStakingRewardsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JuicedStakingRewardsCallerSession struct {
	Contract *JuicedStakingRewardsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// JuicedStakingRewardsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JuicedStakingRewardsTransactorSession struct {
	Contract     *JuicedStakingRewardsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// JuicedStakingRewardsRaw is an auto generated low-level Go binding around an Ethereum contract.
type JuicedStakingRewardsRaw struct {
	Contract *JuicedStakingRewards // Generic contract binding to access the raw methods on
}

// JuicedStakingRewardsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JuicedStakingRewardsCallerRaw struct {
	Contract *JuicedStakingRewardsCaller // Generic read-only contract binding to access the raw methods on
}

// JuicedStakingRewardsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JuicedStakingRewardsTransactorRaw struct {
	Contract *JuicedStakingRewardsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJuicedStakingRewards creates a new instance of JuicedStakingRewards, bound to a specific deployed contract.
func NewJuicedStakingRewards(address common.Address, backend bind.ContractBackend) (*JuicedStakingRewards, error) {
	contract, err := bindJuicedStakingRewards(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &JuicedStakingRewards{JuicedStakingRewardsCaller: JuicedStakingRewardsCaller{contract: contract}, JuicedStakingRewardsTransactor: JuicedStakingRewardsTransactor{contract: contract}, JuicedStakingRewardsFilterer: JuicedStakingRewardsFilterer{contract: contract}}, nil
}

// NewJuicedStakingRewardsCaller creates a new read-only instance of JuicedStakingRewards, bound to a specific deployed contract.
func NewJuicedStakingRewardsCaller(address common.Address, caller bind.ContractCaller) (*JuicedStakingRewardsCaller, error) {
	contract, err := bindJuicedStakingRewards(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JuicedStakingRewardsCaller{contract: contract}, nil
}

// NewJuicedStakingRewardsTransactor creates a new write-only instance of JuicedStakingRewards, bound to a specific deployed contract.
func NewJuicedStakingRewardsTransactor(address common.Address, transactor bind.ContractTransactor) (*JuicedStakingRewardsTransactor, error) {
	contract, err := bindJuicedStakingRewards(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JuicedStakingRewardsTransactor{contract: contract}, nil
}

// NewJuicedStakingRewardsFilterer creates a new log filterer instance of JuicedStakingRewards, bound to a specific deployed contract.
func NewJuicedStakingRewardsFilterer(address common.Address, filterer bind.ContractFilterer) (*JuicedStakingRewardsFilterer, error) {
	contract, err := bindJuicedStakingRewards(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JuicedStakingRewardsFilterer{contract: contract}, nil
}

// bindJuicedStakingRewards binds a generic wrapper to an already deployed contract.
func bindJuicedStakingRewards(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := JuicedStakingRewardsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JuicedStakingRewards *JuicedStakingRewardsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JuicedStakingRewards.Contract.JuicedStakingRewardsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JuicedStakingRewards *JuicedStakingRewardsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JuicedStakingRewards.Contract.JuicedStakingRewardsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JuicedStakingRewards *JuicedStakingRewardsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JuicedStakingRewards.Contract.JuicedStakingRewardsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JuicedStakingRewards *JuicedStakingRewardsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JuicedStakingRewards.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JuicedStakingRewards *JuicedStakingRewardsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JuicedStakingRewards.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JuicedStakingRewards *JuicedStakingRewardsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JuicedStakingRewards.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_JuicedStakingRewards *JuicedStakingRewardsCaller) BalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _JuicedStakingRewards.contract.Call(opts, &out, "balanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_JuicedStakingRewards *JuicedStakingRewardsSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _JuicedStakingRewards.Contract.BalanceOf(&_JuicedStakingRewards.CallOpts, _account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_JuicedStakingRewards *JuicedStakingRewardsCallerSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _JuicedStakingRewards.Contract.BalanceOf(&_JuicedStakingRewards.CallOpts, _account)
}

// Earned is a free data retrieval call binding the contract method 0x211dc32d.
//
// Solidity: function earned(address _account, address _rewardsToken) view returns(uint256 pending)
func (_JuicedStakingRewards *JuicedStakingRewardsCaller) Earned(opts *bind.CallOpts, _account common.Address, _rewardsToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _JuicedStakingRewards.contract.Call(opts, &out, "earned", _account, _rewardsToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Earned is a free data retrieval call binding the contract method 0x211dc32d.
//
// Solidity: function earned(address _account, address _rewardsToken) view returns(uint256 pending)
func (_JuicedStakingRewards *JuicedStakingRewardsSession) Earned(_account common.Address, _rewardsToken common.Address) (*big.Int, error) {
	return _JuicedStakingRewards.Contract.Earned(&_JuicedStakingRewards.CallOpts, _account, _rewardsToken)
}

// Earned is a free data retrieval call binding the contract method 0x211dc32d.
//
// Solidity: function earned(address _account, address _rewardsToken) view returns(uint256 pending)
func (_JuicedStakingRewards *JuicedStakingRewardsCallerSession) Earned(_account common.Address, _rewardsToken common.Address) (*big.Int, error) {
	return _JuicedStakingRewards.Contract.Earned(&_JuicedStakingRewards.CallOpts, _account, _rewardsToken)
}

// GetRewardForDuration is a free data retrieval call binding the contract method 0xbcd11014.
//
// Solidity: function getRewardForDuration(address _rewardsToken) view returns(uint256)
func (_JuicedStakingRewards *JuicedStakingRewardsCaller) GetRewardForDuration(opts *bind.CallOpts, _rewardsToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _JuicedStakingRewards.contract.Call(opts, &out, "getRewardForDuration", _rewardsToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardForDuration is a free data retrieval call binding the contract method 0xbcd11014.
//
// Solidity: function getRewardForDuration(address _rewardsToken) view returns(uint256)
func (_JuicedStakingRewards *JuicedStakingRewardsSession) GetRewardForDuration(_rewardsToken common.Address) (*big.Int, error) {
	return _JuicedStakingRewards.Contract.GetRewardForDuration(&_JuicedStakingRewards.CallOpts, _rewardsToken)
}

// GetRewardForDuration is a free data retrieval call binding the contract method 0xbcd11014.
//
// Solidity: function getRewardForDuration(address _rewardsToken) view returns(uint256)
func (_JuicedStakingRewards *JuicedStakingRewardsCallerSession) GetRewardForDuration(_rewardsToken common.Address) (*big.Int, error) {
	return _JuicedStakingRewards.Contract.GetRewardForDuration(&_JuicedStakingRewards.CallOpts, _rewardsToken)
}

// IsOriginal is a free data retrieval call binding the contract method 0x6f392ce7.
//
// Solidity: function isOriginal() view returns(bool)
func (_JuicedStakingRewards *JuicedStakingRewardsCaller) IsOriginal(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _JuicedStakingRewards.contract.Call(opts, &out, "isOriginal")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOriginal is a free data retrieval call binding the contract method 0x6f392ce7.
//
// Solidity: function isOriginal() view returns(bool)
func (_JuicedStakingRewards *JuicedStakingRewardsSession) IsOriginal() (bool, error) {
	return _JuicedStakingRewards.Contract.IsOriginal(&_JuicedStakingRewards.CallOpts)
}

// IsOriginal is a free data retrieval call binding the contract method 0x6f392ce7.
//
// Solidity: function isOriginal() view returns(bool)
func (_JuicedStakingRewards *JuicedStakingRewardsCallerSession) IsOriginal() (bool, error) {
	return _JuicedStakingRewards.Contract.IsOriginal(&_JuicedStakingRewards.CallOpts)
}

// IsRetired is a free data retrieval call binding the contract method 0x49343cee.
//
// Solidity: function isRetired() view returns(bool)
func (_JuicedStakingRewards *JuicedStakingRewardsCaller) IsRetired(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _JuicedStakingRewards.contract.Call(opts, &out, "isRetired")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRetired is a free data retrieval call binding the contract method 0x49343cee.
//
// Solidity: function isRetired() view returns(bool)
func (_JuicedStakingRewards *JuicedStakingRewardsSession) IsRetired() (bool, error) {
	return _JuicedStakingRewards.Contract.IsRetired(&_JuicedStakingRewards.CallOpts)
}

// IsRetired is a free data retrieval call binding the contract method 0x49343cee.
//
// Solidity: function isRetired() view returns(bool)
func (_JuicedStakingRewards *JuicedStakingRewardsCallerSession) IsRetired() (bool, error) {
	return _JuicedStakingRewards.Contract.IsRetired(&_JuicedStakingRewards.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x638634ee.
//
// Solidity: function lastTimeRewardApplicable(address _rewardsToken) view returns(uint256)
func (_JuicedStakingRewards *JuicedStakingRewardsCaller) LastTimeRewardApplicable(opts *bind.CallOpts, _rewardsToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _JuicedStakingRewards.contract.Call(opts, &out, "lastTimeRewardApplicable", _rewardsToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x638634ee.
//
// Solidity: function lastTimeRewardApplicable(address _rewardsToken) view returns(uint256)
func (_JuicedStakingRewards *JuicedStakingRewardsSession) LastTimeRewardApplicable(_rewardsToken common.Address) (*big.Int, error) {
	return _JuicedStakingRewards.Contract.LastTimeRewardApplicable(&_JuicedStakingRewards.CallOpts, _rewardsToken)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x638634ee.
//
// Solidity: function lastTimeRewardApplicable(address _rewardsToken) view returns(uint256)
func (_JuicedStakingRewards *JuicedStakingRewardsCallerSession) LastTimeRewardApplicable(_rewardsToken common.Address) (*big.Int, error) {
	return _JuicedStakingRewards.Contract.LastTimeRewardApplicable(&_JuicedStakingRewards.CallOpts, _rewardsToken)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_JuicedStakingRewards *JuicedStakingRewardsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JuicedStakingRewards.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_JuicedStakingRewards *JuicedStakingRewardsSession) Owner() (common.Address, error) {
	return _JuicedStakingRewards.Contract.Owner(&_JuicedStakingRewards.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_JuicedStakingRewards *JuicedStakingRewardsCallerSession) Owner() (common.Address, error) {
	return _JuicedStakingRewards.Contract.Owner(&_JuicedStakingRewards.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_JuicedStakingRewards *JuicedStakingRewardsCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _JuicedStakingRewards.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_JuicedStakingRewards *JuicedStakingRewardsSession) Paused() (bool, error) {
	return _JuicedStakingRewards.Contract.Paused(&_JuicedStakingRewards.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_JuicedStakingRewards *JuicedStakingRewardsCallerSession) Paused() (bool, error) {
	return _JuicedStakingRewards.Contract.Paused(&_JuicedStakingRewards.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_JuicedStakingRewards *JuicedStakingRewardsCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JuicedStakingRewards.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_JuicedStakingRewards *JuicedStakingRewardsSession) PendingOwner() (common.Address, error) {
	return _JuicedStakingRewards.Contract.PendingOwner(&_JuicedStakingRewards.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_JuicedStakingRewards *JuicedStakingRewardsCallerSession) PendingOwner() (common.Address, error) {
	return _JuicedStakingRewards.Contract.PendingOwner(&_JuicedStakingRewards.CallOpts)
}

// RewardData is a free data retrieval call binding the contract method 0x48e5d9f8.
//
// Solidity: function rewardData(address ) view returns(address rewardsDistributor, uint256 rewardsDuration, uint256 periodFinish, uint256 rewardRate, uint256 lastUpdateTime, uint256 rewardPerTokenStored)
func (_JuicedStakingRewards *JuicedStakingRewardsCaller) RewardData(opts *bind.CallOpts, arg0 common.Address) (struct {
	RewardsDistributor   common.Address
	RewardsDuration      *big.Int
	PeriodFinish         *big.Int
	RewardRate           *big.Int
	LastUpdateTime       *big.Int
	RewardPerTokenStored *big.Int
}, error) {
	var out []interface{}
	err := _JuicedStakingRewards.contract.Call(opts, &out, "rewardData", arg0)

	outstruct := new(struct {
		RewardsDistributor   common.Address
		RewardsDuration      *big.Int
		PeriodFinish         *big.Int
		RewardRate           *big.Int
		LastUpdateTime       *big.Int
		RewardPerTokenStored *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RewardsDistributor = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.RewardsDuration = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.PeriodFinish = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.RewardRate = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LastUpdateTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.RewardPerTokenStored = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RewardData is a free data retrieval call binding the contract method 0x48e5d9f8.
//
// Solidity: function rewardData(address ) view returns(address rewardsDistributor, uint256 rewardsDuration, uint256 periodFinish, uint256 rewardRate, uint256 lastUpdateTime, uint256 rewardPerTokenStored)
func (_JuicedStakingRewards *JuicedStakingRewardsSession) RewardData(arg0 common.Address) (struct {
	RewardsDistributor   common.Address
	RewardsDuration      *big.Int
	PeriodFinish         *big.Int
	RewardRate           *big.Int
	LastUpdateTime       *big.Int
	RewardPerTokenStored *big.Int
}, error) {
	return _JuicedStakingRewards.Contract.RewardData(&_JuicedStakingRewards.CallOpts, arg0)
}

// RewardData is a free data retrieval call binding the contract method 0x48e5d9f8.
//
// Solidity: function rewardData(address ) view returns(address rewardsDistributor, uint256 rewardsDuration, uint256 periodFinish, uint256 rewardRate, uint256 lastUpdateTime, uint256 rewardPerTokenStored)
func (_JuicedStakingRewards *JuicedStakingRewardsCallerSession) RewardData(arg0 common.Address) (struct {
	RewardsDistributor   common.Address
	RewardsDuration      *big.Int
	PeriodFinish         *big.Int
	RewardRate           *big.Int
	LastUpdateTime       *big.Int
	RewardPerTokenStored *big.Int
}, error) {
	return _JuicedStakingRewards.Contract.RewardData(&_JuicedStakingRewards.CallOpts, arg0)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xf1229777.
//
// Solidity: function rewardPerToken(address _rewardsToken) view returns(uint256 rewardAmount)
func (_JuicedStakingRewards *JuicedStakingRewardsCaller) RewardPerToken(opts *bind.CallOpts, _rewardsToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _JuicedStakingRewards.contract.Call(opts, &out, "rewardPerToken", _rewardsToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerToken is a free data retrieval call binding the contract method 0xf1229777.
//
// Solidity: function rewardPerToken(address _rewardsToken) view returns(uint256 rewardAmount)
func (_JuicedStakingRewards *JuicedStakingRewardsSession) RewardPerToken(_rewardsToken common.Address) (*big.Int, error) {
	return _JuicedStakingRewards.Contract.RewardPerToken(&_JuicedStakingRewards.CallOpts, _rewardsToken)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xf1229777.
//
// Solidity: function rewardPerToken(address _rewardsToken) view returns(uint256 rewardAmount)
func (_JuicedStakingRewards *JuicedStakingRewardsCallerSession) RewardPerToken(_rewardsToken common.Address) (*big.Int, error) {
	return _JuicedStakingRewards.Contract.RewardPerToken(&_JuicedStakingRewards.CallOpts, _rewardsToken)
}

// RewardTokens is a free data retrieval call binding the contract method 0x7bb7bed1.
//
// Solidity: function rewardTokens(uint256 ) view returns(address)
func (_JuicedStakingRewards *JuicedStakingRewardsCaller) RewardTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _JuicedStakingRewards.contract.Call(opts, &out, "rewardTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardTokens is a free data retrieval call binding the contract method 0x7bb7bed1.
//
// Solidity: function rewardTokens(uint256 ) view returns(address)
func (_JuicedStakingRewards *JuicedStakingRewardsSession) RewardTokens(arg0 *big.Int) (common.Address, error) {
	return _JuicedStakingRewards.Contract.RewardTokens(&_JuicedStakingRewards.CallOpts, arg0)
}

// RewardTokens is a free data retrieval call binding the contract method 0x7bb7bed1.
//
// Solidity: function rewardTokens(uint256 ) view returns(address)
func (_JuicedStakingRewards *JuicedStakingRewardsCallerSession) RewardTokens(arg0 *big.Int) (common.Address, error) {
	return _JuicedStakingRewards.Contract.RewardTokens(&_JuicedStakingRewards.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0xe70b9e27.
//
// Solidity: function rewards(address , address ) view returns(uint256)
func (_JuicedStakingRewards *JuicedStakingRewardsCaller) Rewards(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _JuicedStakingRewards.contract.Call(opts, &out, "rewards", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0xe70b9e27.
//
// Solidity: function rewards(address , address ) view returns(uint256)
func (_JuicedStakingRewards *JuicedStakingRewardsSession) Rewards(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _JuicedStakingRewards.Contract.Rewards(&_JuicedStakingRewards.CallOpts, arg0, arg1)
}

// Rewards is a free data retrieval call binding the contract method 0xe70b9e27.
//
// Solidity: function rewards(address , address ) view returns(uint256)
func (_JuicedStakingRewards *JuicedStakingRewardsCallerSession) Rewards(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _JuicedStakingRewards.Contract.Rewards(&_JuicedStakingRewards.CallOpts, arg0, arg1)
}

// StakerVersion is a free data retrieval call binding the contract method 0x0638ca2a.
//
// Solidity: function stakerVersion() view returns(string)
func (_JuicedStakingRewards *JuicedStakingRewardsCaller) StakerVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _JuicedStakingRewards.contract.Call(opts, &out, "stakerVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// StakerVersion is a free data retrieval call binding the contract method 0x0638ca2a.
//
// Solidity: function stakerVersion() view returns(string)
func (_JuicedStakingRewards *JuicedStakingRewardsSession) StakerVersion() (string, error) {
	return _JuicedStakingRewards.Contract.StakerVersion(&_JuicedStakingRewards.CallOpts)
}

// StakerVersion is a free data retrieval call binding the contract method 0x0638ca2a.
//
// Solidity: function stakerVersion() view returns(string)
func (_JuicedStakingRewards *JuicedStakingRewardsCallerSession) StakerVersion() (string, error) {
	return _JuicedStakingRewards.Contract.StakerVersion(&_JuicedStakingRewards.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_JuicedStakingRewards *JuicedStakingRewardsCaller) StakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JuicedStakingRewards.contract.Call(opts, &out, "stakingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_JuicedStakingRewards *JuicedStakingRewardsSession) StakingToken() (common.Address, error) {
	return _JuicedStakingRewards.Contract.StakingToken(&_JuicedStakingRewards.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_JuicedStakingRewards *JuicedStakingRewardsCallerSession) StakingToken() (common.Address, error) {
	return _JuicedStakingRewards.Contract.StakingToken(&_JuicedStakingRewards.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_JuicedStakingRewards *JuicedStakingRewardsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _JuicedStakingRewards.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_JuicedStakingRewards *JuicedStakingRewardsSession) TotalSupply() (*big.Int, error) {
	return _JuicedStakingRewards.Contract.TotalSupply(&_JuicedStakingRewards.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_JuicedStakingRewards *JuicedStakingRewardsCallerSession) TotalSupply() (*big.Int, error) {
	return _JuicedStakingRewards.Contract.TotalSupply(&_JuicedStakingRewards.CallOpts)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x7035ab98.
//
// Solidity: function userRewardPerTokenPaid(address , address ) view returns(uint256)
func (_JuicedStakingRewards *JuicedStakingRewardsCaller) UserRewardPerTokenPaid(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _JuicedStakingRewards.contract.Call(opts, &out, "userRewardPerTokenPaid", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x7035ab98.
//
// Solidity: function userRewardPerTokenPaid(address , address ) view returns(uint256)
func (_JuicedStakingRewards *JuicedStakingRewardsSession) UserRewardPerTokenPaid(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _JuicedStakingRewards.Contract.UserRewardPerTokenPaid(&_JuicedStakingRewards.CallOpts, arg0, arg1)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x7035ab98.
//
// Solidity: function userRewardPerTokenPaid(address , address ) view returns(uint256)
func (_JuicedStakingRewards *JuicedStakingRewardsCallerSession) UserRewardPerTokenPaid(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _JuicedStakingRewards.Contract.UserRewardPerTokenPaid(&_JuicedStakingRewards.CallOpts, arg0, arg1)
}

// ZapContract is a free data retrieval call binding the contract method 0x2549dad9.
//
// Solidity: function zapContract() view returns(address)
func (_JuicedStakingRewards *JuicedStakingRewardsCaller) ZapContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JuicedStakingRewards.contract.Call(opts, &out, "zapContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZapContract is a free data retrieval call binding the contract method 0x2549dad9.
//
// Solidity: function zapContract() view returns(address)
func (_JuicedStakingRewards *JuicedStakingRewardsSession) ZapContract() (common.Address, error) {
	return _JuicedStakingRewards.Contract.ZapContract(&_JuicedStakingRewards.CallOpts)
}

// ZapContract is a free data retrieval call binding the contract method 0x2549dad9.
//
// Solidity: function zapContract() view returns(address)
func (_JuicedStakingRewards *JuicedStakingRewardsCallerSession) ZapContract() (common.Address, error) {
	return _JuicedStakingRewards.Contract.ZapContract(&_JuicedStakingRewards.CallOpts)
}

// AcceptOwner is a paid mutator transaction binding the contract method 0xebbc4965.
//
// Solidity: function acceptOwner() returns()
func (_JuicedStakingRewards *JuicedStakingRewardsTransactor) AcceptOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JuicedStakingRewards.contract.Transact(opts, "acceptOwner")
}

// AcceptOwner is a paid mutator transaction binding the contract method 0xebbc4965.
//
// Solidity: function acceptOwner() returns()
func (_JuicedStakingRewards *JuicedStakingRewardsSession) AcceptOwner() (*types.Transaction, error) {
	return _JuicedStakingRewards.Contract.AcceptOwner(&_JuicedStakingRewards.TransactOpts)
}

// AcceptOwner is a paid mutator transaction binding the contract method 0xebbc4965.
//
// Solidity: function acceptOwner() returns()
func (_JuicedStakingRewards *JuicedStakingRewardsTransactorSession) AcceptOwner() (*types.Transaction, error) {
	return _JuicedStakingRewards.Contract.AcceptOwner(&_JuicedStakingRewards.TransactOpts)
}

// AddReward is a paid mutator transaction binding the contract method 0xd0ed26ae.
//
// Solidity: function addReward(address _rewardsToken, address _rewardsDistributor, uint256 _rewardsDuration) returns()
func (_JuicedStakingRewards *JuicedStakingRewardsTransactor) AddReward(opts *bind.TransactOpts, _rewardsToken common.Address, _rewardsDistributor common.Address, _rewardsDuration *big.Int) (*types.Transaction, error) {
	return _JuicedStakingRewards.contract.Transact(opts, "addReward", _rewardsToken, _rewardsDistributor, _rewardsDuration)
}

// AddReward is a paid mutator transaction binding the contract method 0xd0ed26ae.
//
// Solidity: function addReward(address _rewardsToken, address _rewardsDistributor, uint256 _rewardsDuration) returns()
func (_JuicedStakingRewards *JuicedStakingRewardsSession) AddReward(_rewardsToken common.Address, _rewardsDistributor common.Address, _rewardsDuration *big.Int) (*types.Transaction, error) {
	return _JuicedStakingRewards.Contract.AddReward(&_JuicedStakingRewards.TransactOpts, _rewardsToken, _rewardsDistributor, _rewardsDuration)
}

// AddReward is a paid mutator transaction binding the contract method 0xd0ed26ae.
//
// Solidity: function addReward(address _rewardsToken, address _rewardsDistributor, uint256 _rewardsDuration) returns()
func (_JuicedStakingRewards *JuicedStakingRewardsTransactorSession) AddReward(_rewardsToken common.Address, _rewardsDistributor common.Address, _rewardsDuration *big.Int) (*types.Transaction, error) {
	return _JuicedStakingRewards.Contract.AddReward(&_JuicedStakingRewards.TransactOpts, _rewardsToken, _rewardsDistributor, _rewardsDuration)
}

// CloneStakingPool is a paid mutator transaction binding the contract method 0x754a0fdc.
//
// Solidity: function cloneStakingPool(address _owner, address _stakingToken, address _zapContract) returns(address newStakingPool)
func (_JuicedStakingRewards *JuicedStakingRewardsTransactor) CloneStakingPool(opts *bind.TransactOpts, _owner common.Address, _stakingToken common.Address, _zapContract common.Address) (*types.Transaction, error) {
	return _JuicedStakingRewards.contract.Transact(opts, "cloneStakingPool", _owner, _stakingToken, _zapContract)
}

// CloneStakingPool is a paid mutator transaction binding the contract method 0x754a0fdc.
//
// Solidity: function cloneStakingPool(address _owner, address _stakingToken, address _zapContract) returns(address newStakingPool)
func (_JuicedStakingRewards *JuicedStakingRewardsSession) CloneStakingPool(_owner common.Address, _stakingToken common.Address, _zapContract common.Address) (*types.Transaction, error) {
	return _JuicedStakingRewards.Contract.CloneStakingPool(&_JuicedStakingRewards.TransactOpts, _owner, _stakingToken, _zapContract)
}

// CloneStakingPool is a paid mutator transaction binding the contract method 0x754a0fdc.
//
// Solidity: function cloneStakingPool(address _owner, address _stakingToken, address _zapContract) returns(address newStakingPool)
func (_JuicedStakingRewards *JuicedStakingRewardsTransactorSession) CloneStakingPool(_owner common.Address, _stakingToken common.Address, _zapContract common.Address) (*types.Transaction, error) {
	return _JuicedStakingRewards.Contract.CloneStakingPool(&_JuicedStakingRewards.TransactOpts, _owner, _stakingToken, _zapContract)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_JuicedStakingRewards *JuicedStakingRewardsTransactor) Exit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JuicedStakingRewards.contract.Transact(opts, "exit")
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_JuicedStakingRewards *JuicedStakingRewardsSession) Exit() (*types.Transaction, error) {
	return _JuicedStakingRewards.Contract.Exit(&_JuicedStakingRewards.TransactOpts)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_JuicedStakingRewards *JuicedStakingRewardsTransactorSession) Exit() (*types.Transaction, error) {
	return _JuicedStakingRewards.Contract.Exit(&_JuicedStakingRewards.TransactOpts)
}

// GetOneReward is a paid mutator transaction binding the contract method 0x4ef52305.
//
// Solidity: function getOneReward(address _rewardsToken) returns()
func (_JuicedStakingRewards *JuicedStakingRewardsTransactor) GetOneReward(opts *bind.TransactOpts, _rewardsToken common.Address) (*types.Transaction, error) {
	return _JuicedStakingRewards.contract.Transact(opts, "getOneReward", _rewardsToken)
}

// GetOneReward is a paid mutator transaction binding the contract method 0x4ef52305.
//
// Solidity: function getOneReward(address _rewardsToken) returns()
func (_JuicedStakingRewards *JuicedStakingRewardsSession) GetOneReward(_rewardsToken common.Address) (*types.Transaction, error) {
	return _JuicedStakingRewards.Contract.GetOneReward(&_JuicedStakingRewards.TransactOpts, _rewardsToken)
}

// GetOneReward is a paid mutator transaction binding the contract method 0x4ef52305.
//
// Solidity: function getOneReward(address _rewardsToken) returns()
func (_JuicedStakingRewards *JuicedStakingRewardsTransactorSession) GetOneReward(_rewardsToken common.Address) (*types.Transaction, error) {
	return _JuicedStakingRewards.Contract.GetOneReward(&_JuicedStakingRewards.TransactOpts, _rewardsToken)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_JuicedStakingRewards *JuicedStakingRewardsTransactor) GetReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JuicedStakingRewards.contract.Transact(opts, "getReward")
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_JuicedStakingRewards *JuicedStakingRewardsSession) GetReward() (*types.Transaction, error) {
	return _JuicedStakingRewards.Contract.GetReward(&_JuicedStakingRewards.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_JuicedStakingRewards *JuicedStakingRewardsTransactorSession) GetReward() (*types.Transaction, error) {
	return _JuicedStakingRewards.Contract.GetReward(&_JuicedStakingRewards.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _owner, address _stakingToken, address _zapContract) returns()
func (_JuicedStakingRewards *JuicedStakingRewardsTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _stakingToken common.Address, _zapContract common.Address) (*types.Transaction, error) {
	return _JuicedStakingRewards.contract.Transact(opts, "initialize", _owner, _stakingToken, _zapContract)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _owner, address _stakingToken, address _zapContract) returns()
func (_JuicedStakingRewards *JuicedStakingRewardsSession) Initialize(_owner common.Address, _stakingToken common.Address, _zapContract common.Address) (*types.Transaction, error) {
	return _JuicedStakingRewards.Contract.Initialize(&_JuicedStakingRewards.TransactOpts, _owner, _stakingToken, _zapContract)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _owner, address _stakingToken, address _zapContract) returns()
func (_JuicedStakingRewards *JuicedStakingRewardsTransactorSession) Initialize(_owner common.Address, _stakingToken common.Address, _zapContract common.Address) (*types.Transaction, error) {
	return _JuicedStakingRewards.Contract.Initialize(&_JuicedStakingRewards.TransactOpts, _owner, _stakingToken, _zapContract)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0xb66503cf.
//
// Solidity: function notifyRewardAmount(address _rewardsToken, uint256 _rewardAmount) returns()
func (_JuicedStakingRewards *JuicedStakingRewardsTransactor) NotifyRewardAmount(opts *bind.TransactOpts, _rewardsToken common.Address, _rewardAmount *big.Int) (*types.Transaction, error) {
	return _JuicedStakingRewards.contract.Transact(opts, "notifyRewardAmount", _rewardsToken, _rewardAmount)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0xb66503cf.
//
// Solidity: function notifyRewardAmount(address _rewardsToken, uint256 _rewardAmount) returns()
func (_JuicedStakingRewards *JuicedStakingRewardsSession) NotifyRewardAmount(_rewardsToken common.Address, _rewardAmount *big.Int) (*types.Transaction, error) {
	return _JuicedStakingRewards.Contract.NotifyRewardAmount(&_JuicedStakingRewards.TransactOpts, _rewardsToken, _rewardAmount)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0xb66503cf.
//
// Solidity: function notifyRewardAmount(address _rewardsToken, uint256 _rewardAmount) returns()
func (_JuicedStakingRewards *JuicedStakingRewardsTransactorSession) NotifyRewardAmount(_rewardsToken common.Address, _rewardAmount *big.Int) (*types.Transaction, error) {
	return _JuicedStakingRewards.Contract.NotifyRewardAmount(&_JuicedStakingRewards.TransactOpts, _rewardsToken, _rewardAmount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address _tokenAddress, uint256 _tokenAmount) returns()
func (_JuicedStakingRewards *JuicedStakingRewardsTransactor) RecoverERC20(opts *bind.TransactOpts, _tokenAddress common.Address, _tokenAmount *big.Int) (*types.Transaction, error) {
	return _JuicedStakingRewards.contract.Transact(opts, "recoverERC20", _tokenAddress, _tokenAmount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address _tokenAddress, uint256 _tokenAmount) returns()
func (_JuicedStakingRewards *JuicedStakingRewardsSession) RecoverERC20(_tokenAddress common.Address, _tokenAmount *big.Int) (*types.Transaction, error) {
	return _JuicedStakingRewards.Contract.RecoverERC20(&_JuicedStakingRewards.TransactOpts, _tokenAddress, _tokenAmount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address _tokenAddress, uint256 _tokenAmount) returns()
func (_JuicedStakingRewards *JuicedStakingRewardsTransactorSession) RecoverERC20(_tokenAddress common.Address, _tokenAmount *big.Int) (*types.Transaction, error) {
	return _JuicedStakingRewards.Contract.RecoverERC20(&_JuicedStakingRewards.TransactOpts, _tokenAddress, _tokenAmount)
}

// SetPendingOwner is a paid mutator transaction binding the contract method 0xc42069ec.
//
// Solidity: function setPendingOwner(address _owner) returns()
func (_JuicedStakingRewards *JuicedStakingRewardsTransactor) SetPendingOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _JuicedStakingRewards.contract.Transact(opts, "setPendingOwner", _owner)
}

// SetPendingOwner is a paid mutator transaction binding the contract method 0xc42069ec.
//
// Solidity: function setPendingOwner(address _owner) returns()
func (_JuicedStakingRewards *JuicedStakingRewardsSession) SetPendingOwner(_owner common.Address) (*types.Transaction, error) {
	return _JuicedStakingRewards.Contract.SetPendingOwner(&_JuicedStakingRewards.TransactOpts, _owner)
}

// SetPendingOwner is a paid mutator transaction binding the contract method 0xc42069ec.
//
// Solidity: function setPendingOwner(address _owner) returns()
func (_JuicedStakingRewards *JuicedStakingRewardsTransactorSession) SetPendingOwner(_owner common.Address) (*types.Transaction, error) {
	return _JuicedStakingRewards.Contract.SetPendingOwner(&_JuicedStakingRewards.TransactOpts, _owner)
}

// SetRewardsDistributor is a paid mutator transaction binding the contract method 0x3f695b45.
//
// Solidity: function setRewardsDistributor(address _rewardsToken, address _rewardsDistributor) returns()
func (_JuicedStakingRewards *JuicedStakingRewardsTransactor) SetRewardsDistributor(opts *bind.TransactOpts, _rewardsToken common.Address, _rewardsDistributor common.Address) (*types.Transaction, error) {
	return _JuicedStakingRewards.contract.Transact(opts, "setRewardsDistributor", _rewardsToken, _rewardsDistributor)
}

// SetRewardsDistributor is a paid mutator transaction binding the contract method 0x3f695b45.
//
// Solidity: function setRewardsDistributor(address _rewardsToken, address _rewardsDistributor) returns()
func (_JuicedStakingRewards *JuicedStakingRewardsSession) SetRewardsDistributor(_rewardsToken common.Address, _rewardsDistributor common.Address) (*types.Transaction, error) {
	return _JuicedStakingRewards.Contract.SetRewardsDistributor(&_JuicedStakingRewards.TransactOpts, _rewardsToken, _rewardsDistributor)
}

// SetRewardsDistributor is a paid mutator transaction binding the contract method 0x3f695b45.
//
// Solidity: function setRewardsDistributor(address _rewardsToken, address _rewardsDistributor) returns()
func (_JuicedStakingRewards *JuicedStakingRewardsTransactorSession) SetRewardsDistributor(_rewardsToken common.Address, _rewardsDistributor common.Address) (*types.Transaction, error) {
	return _JuicedStakingRewards.Contract.SetRewardsDistributor(&_JuicedStakingRewards.TransactOpts, _rewardsToken, _rewardsDistributor)
}

// SetRewardsDuration is a paid mutator transaction binding the contract method 0x2378bea6.
//
// Solidity: function setRewardsDuration(address _rewardsToken, uint256 _rewardsDuration) returns()
func (_JuicedStakingRewards *JuicedStakingRewardsTransactor) SetRewardsDuration(opts *bind.TransactOpts, _rewardsToken common.Address, _rewardsDuration *big.Int) (*types.Transaction, error) {
	return _JuicedStakingRewards.contract.Transact(opts, "setRewardsDuration", _rewardsToken, _rewardsDuration)
}

// SetRewardsDuration is a paid mutator transaction binding the contract method 0x2378bea6.
//
// Solidity: function setRewardsDuration(address _rewardsToken, uint256 _rewardsDuration) returns()
func (_JuicedStakingRewards *JuicedStakingRewardsSession) SetRewardsDuration(_rewardsToken common.Address, _rewardsDuration *big.Int) (*types.Transaction, error) {
	return _JuicedStakingRewards.Contract.SetRewardsDuration(&_JuicedStakingRewards.TransactOpts, _rewardsToken, _rewardsDuration)
}

// SetRewardsDuration is a paid mutator transaction binding the contract method 0x2378bea6.
//
// Solidity: function setRewardsDuration(address _rewardsToken, uint256 _rewardsDuration) returns()
func (_JuicedStakingRewards *JuicedStakingRewardsTransactorSession) SetRewardsDuration(_rewardsToken common.Address, _rewardsDuration *big.Int) (*types.Transaction, error) {
	return _JuicedStakingRewards.Contract.SetRewardsDuration(&_JuicedStakingRewards.TransactOpts, _rewardsToken, _rewardsDuration)
}

// SetZapContract is a paid mutator transaction binding the contract method 0xfcaa8737.
//
// Solidity: function setZapContract(address _zapContract) returns()
func (_JuicedStakingRewards *JuicedStakingRewardsTransactor) SetZapContract(opts *bind.TransactOpts, _zapContract common.Address) (*types.Transaction, error) {
	return _JuicedStakingRewards.contract.Transact(opts, "setZapContract", _zapContract)
}

// SetZapContract is a paid mutator transaction binding the contract method 0xfcaa8737.
//
// Solidity: function setZapContract(address _zapContract) returns()
func (_JuicedStakingRewards *JuicedStakingRewardsSession) SetZapContract(_zapContract common.Address) (*types.Transaction, error) {
	return _JuicedStakingRewards.Contract.SetZapContract(&_JuicedStakingRewards.TransactOpts, _zapContract)
}

// SetZapContract is a paid mutator transaction binding the contract method 0xfcaa8737.
//
// Solidity: function setZapContract(address _zapContract) returns()
func (_JuicedStakingRewards *JuicedStakingRewardsTransactorSession) SetZapContract(_zapContract common.Address) (*types.Transaction, error) {
	return _JuicedStakingRewards.Contract.SetZapContract(&_JuicedStakingRewards.TransactOpts, _zapContract)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns()
func (_JuicedStakingRewards *JuicedStakingRewardsTransactor) Stake(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _JuicedStakingRewards.contract.Transact(opts, "stake", _amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns()
func (_JuicedStakingRewards *JuicedStakingRewardsSession) Stake(_amount *big.Int) (*types.Transaction, error) {
	return _JuicedStakingRewards.Contract.Stake(&_JuicedStakingRewards.TransactOpts, _amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns()
func (_JuicedStakingRewards *JuicedStakingRewardsTransactorSession) Stake(_amount *big.Int) (*types.Transaction, error) {
	return _JuicedStakingRewards.Contract.Stake(&_JuicedStakingRewards.TransactOpts, _amount)
}

// StakeFor is a paid mutator transaction binding the contract method 0x2ee40908.
//
// Solidity: function stakeFor(address _recipient, uint256 _amount) returns()
func (_JuicedStakingRewards *JuicedStakingRewardsTransactor) StakeFor(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _JuicedStakingRewards.contract.Transact(opts, "stakeFor", _recipient, _amount)
}

// StakeFor is a paid mutator transaction binding the contract method 0x2ee40908.
//
// Solidity: function stakeFor(address _recipient, uint256 _amount) returns()
func (_JuicedStakingRewards *JuicedStakingRewardsSession) StakeFor(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _JuicedStakingRewards.Contract.StakeFor(&_JuicedStakingRewards.TransactOpts, _recipient, _amount)
}

// StakeFor is a paid mutator transaction binding the contract method 0x2ee40908.
//
// Solidity: function stakeFor(address _recipient, uint256 _amount) returns()
func (_JuicedStakingRewards *JuicedStakingRewardsTransactorSession) StakeFor(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _JuicedStakingRewards.Contract.StakeFor(&_JuicedStakingRewards.TransactOpts, _recipient, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_JuicedStakingRewards *JuicedStakingRewardsTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _JuicedStakingRewards.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_JuicedStakingRewards *JuicedStakingRewardsSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _JuicedStakingRewards.Contract.Withdraw(&_JuicedStakingRewards.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_JuicedStakingRewards *JuicedStakingRewardsTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _JuicedStakingRewards.Contract.Withdraw(&_JuicedStakingRewards.TransactOpts, _amount)
}

// WithdrawFor is a paid mutator transaction binding the contract method 0x0e19c699.
//
// Solidity: function withdrawFor(address _recipient, uint256 _amount, bool _exit) returns()
func (_JuicedStakingRewards *JuicedStakingRewardsTransactor) WithdrawFor(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int, _exit bool) (*types.Transaction, error) {
	return _JuicedStakingRewards.contract.Transact(opts, "withdrawFor", _recipient, _amount, _exit)
}

// WithdrawFor is a paid mutator transaction binding the contract method 0x0e19c699.
//
// Solidity: function withdrawFor(address _recipient, uint256 _amount, bool _exit) returns()
func (_JuicedStakingRewards *JuicedStakingRewardsSession) WithdrawFor(_recipient common.Address, _amount *big.Int, _exit bool) (*types.Transaction, error) {
	return _JuicedStakingRewards.Contract.WithdrawFor(&_JuicedStakingRewards.TransactOpts, _recipient, _amount, _exit)
}

// WithdrawFor is a paid mutator transaction binding the contract method 0x0e19c699.
//
// Solidity: function withdrawFor(address _recipient, uint256 _amount, bool _exit) returns()
func (_JuicedStakingRewards *JuicedStakingRewardsTransactorSession) WithdrawFor(_recipient common.Address, _amount *big.Int, _exit bool) (*types.Transaction, error) {
	return _JuicedStakingRewards.Contract.WithdrawFor(&_JuicedStakingRewards.TransactOpts, _recipient, _amount, _exit)
}

// JuicedStakingRewardsClonedIterator is returned from FilterCloned and is used to iterate over the raw logs and unpacked data for Cloned events raised by the JuicedStakingRewards contract.
type JuicedStakingRewardsClonedIterator struct {
	Event *JuicedStakingRewardsCloned // Event containing the contract specifics and raw log

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
func (it *JuicedStakingRewardsClonedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JuicedStakingRewardsCloned)
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
		it.Event = new(JuicedStakingRewardsCloned)
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
func (it *JuicedStakingRewardsClonedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JuicedStakingRewardsClonedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JuicedStakingRewardsCloned represents a Cloned event raised by the JuicedStakingRewards contract.
type JuicedStakingRewardsCloned struct {
	Clone common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCloned is a free log retrieval operation binding the contract event 0x783540fb4221a3238720dc7038937d0d79982bcf895274aa6ad179f82cf0d53c.
//
// Solidity: event Cloned(address indexed clone)
func (_JuicedStakingRewards *JuicedStakingRewardsFilterer) FilterCloned(opts *bind.FilterOpts, clone []common.Address) (*JuicedStakingRewardsClonedIterator, error) {

	var cloneRule []interface{}
	for _, cloneItem := range clone {
		cloneRule = append(cloneRule, cloneItem)
	}

	logs, sub, err := _JuicedStakingRewards.contract.FilterLogs(opts, "Cloned", cloneRule)
	if err != nil {
		return nil, err
	}
	return &JuicedStakingRewardsClonedIterator{contract: _JuicedStakingRewards.contract, event: "Cloned", logs: logs, sub: sub}, nil
}

// WatchCloned is a free log subscription operation binding the contract event 0x783540fb4221a3238720dc7038937d0d79982bcf895274aa6ad179f82cf0d53c.
//
// Solidity: event Cloned(address indexed clone)
func (_JuicedStakingRewards *JuicedStakingRewardsFilterer) WatchCloned(opts *bind.WatchOpts, sink chan<- *JuicedStakingRewardsCloned, clone []common.Address) (event.Subscription, error) {

	var cloneRule []interface{}
	for _, cloneItem := range clone {
		cloneRule = append(cloneRule, cloneItem)
	}

	logs, sub, err := _JuicedStakingRewards.contract.WatchLogs(opts, "Cloned", cloneRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JuicedStakingRewardsCloned)
				if err := _JuicedStakingRewards.contract.UnpackLog(event, "Cloned", log); err != nil {
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

// ParseCloned is a log parse operation binding the contract event 0x783540fb4221a3238720dc7038937d0d79982bcf895274aa6ad179f82cf0d53c.
//
// Solidity: event Cloned(address indexed clone)
func (_JuicedStakingRewards *JuicedStakingRewardsFilterer) ParseCloned(log types.Log) (*JuicedStakingRewardsCloned, error) {
	event := new(JuicedStakingRewardsCloned)
	if err := _JuicedStakingRewards.contract.UnpackLog(event, "Cloned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JuicedStakingRewardsOwnerUpdatedIterator is returned from FilterOwnerUpdated and is used to iterate over the raw logs and unpacked data for OwnerUpdated events raised by the JuicedStakingRewards contract.
type JuicedStakingRewardsOwnerUpdatedIterator struct {
	Event *JuicedStakingRewardsOwnerUpdated // Event containing the contract specifics and raw log

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
func (it *JuicedStakingRewardsOwnerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JuicedStakingRewardsOwnerUpdated)
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
		it.Event = new(JuicedStakingRewardsOwnerUpdated)
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
func (it *JuicedStakingRewardsOwnerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JuicedStakingRewardsOwnerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JuicedStakingRewardsOwnerUpdated represents a OwnerUpdated event raised by the JuicedStakingRewards contract.
type JuicedStakingRewardsOwnerUpdated struct {
	Ownererance common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOwnerUpdated is a free log retrieval operation binding the contract event 0x4ffd725fc4a22075e9ec71c59edf9c38cdeb588a91b24fc5b61388c5be41282b.
//
// Solidity: event OwnerUpdated(address indexed Ownererance)
func (_JuicedStakingRewards *JuicedStakingRewardsFilterer) FilterOwnerUpdated(opts *bind.FilterOpts, Ownererance []common.Address) (*JuicedStakingRewardsOwnerUpdatedIterator, error) {

	var OwnereranceRule []interface{}
	for _, OwnereranceItem := range Ownererance {
		OwnereranceRule = append(OwnereranceRule, OwnereranceItem)
	}

	logs, sub, err := _JuicedStakingRewards.contract.FilterLogs(opts, "OwnerUpdated", OwnereranceRule)
	if err != nil {
		return nil, err
	}
	return &JuicedStakingRewardsOwnerUpdatedIterator{contract: _JuicedStakingRewards.contract, event: "OwnerUpdated", logs: logs, sub: sub}, nil
}

// WatchOwnerUpdated is a free log subscription operation binding the contract event 0x4ffd725fc4a22075e9ec71c59edf9c38cdeb588a91b24fc5b61388c5be41282b.
//
// Solidity: event OwnerUpdated(address indexed Ownererance)
func (_JuicedStakingRewards *JuicedStakingRewardsFilterer) WatchOwnerUpdated(opts *bind.WatchOpts, sink chan<- *JuicedStakingRewardsOwnerUpdated, Ownererance []common.Address) (event.Subscription, error) {

	var OwnereranceRule []interface{}
	for _, OwnereranceItem := range Ownererance {
		OwnereranceRule = append(OwnereranceRule, OwnereranceItem)
	}

	logs, sub, err := _JuicedStakingRewards.contract.WatchLogs(opts, "OwnerUpdated", OwnereranceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JuicedStakingRewardsOwnerUpdated)
				if err := _JuicedStakingRewards.contract.UnpackLog(event, "OwnerUpdated", log); err != nil {
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

// ParseOwnerUpdated is a log parse operation binding the contract event 0x4ffd725fc4a22075e9ec71c59edf9c38cdeb588a91b24fc5b61388c5be41282b.
//
// Solidity: event OwnerUpdated(address indexed Ownererance)
func (_JuicedStakingRewards *JuicedStakingRewardsFilterer) ParseOwnerUpdated(log types.Log) (*JuicedStakingRewardsOwnerUpdated, error) {
	event := new(JuicedStakingRewardsOwnerUpdated)
	if err := _JuicedStakingRewards.contract.UnpackLog(event, "OwnerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JuicedStakingRewardsPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the JuicedStakingRewards contract.
type JuicedStakingRewardsPausedIterator struct {
	Event *JuicedStakingRewardsPaused // Event containing the contract specifics and raw log

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
func (it *JuicedStakingRewardsPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JuicedStakingRewardsPaused)
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
		it.Event = new(JuicedStakingRewardsPaused)
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
func (it *JuicedStakingRewardsPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JuicedStakingRewardsPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JuicedStakingRewardsPaused represents a Paused event raised by the JuicedStakingRewards contract.
type JuicedStakingRewardsPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_JuicedStakingRewards *JuicedStakingRewardsFilterer) FilterPaused(opts *bind.FilterOpts) (*JuicedStakingRewardsPausedIterator, error) {

	logs, sub, err := _JuicedStakingRewards.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &JuicedStakingRewardsPausedIterator{contract: _JuicedStakingRewards.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_JuicedStakingRewards *JuicedStakingRewardsFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *JuicedStakingRewardsPaused) (event.Subscription, error) {

	logs, sub, err := _JuicedStakingRewards.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JuicedStakingRewardsPaused)
				if err := _JuicedStakingRewards.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_JuicedStakingRewards *JuicedStakingRewardsFilterer) ParsePaused(log types.Log) (*JuicedStakingRewardsPaused, error) {
	event := new(JuicedStakingRewardsPaused)
	if err := _JuicedStakingRewards.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JuicedStakingRewardsRecoveredIterator is returned from FilterRecovered and is used to iterate over the raw logs and unpacked data for Recovered events raised by the JuicedStakingRewards contract.
type JuicedStakingRewardsRecoveredIterator struct {
	Event *JuicedStakingRewardsRecovered // Event containing the contract specifics and raw log

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
func (it *JuicedStakingRewardsRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JuicedStakingRewardsRecovered)
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
		it.Event = new(JuicedStakingRewardsRecovered)
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
func (it *JuicedStakingRewardsRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JuicedStakingRewardsRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JuicedStakingRewardsRecovered represents a Recovered event raised by the JuicedStakingRewards contract.
type JuicedStakingRewardsRecovered struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRecovered is a free log retrieval operation binding the contract event 0x8c1256b8896378cd5044f80c202f9772b9d77dc85c8a6eb51967210b09bfaa28.
//
// Solidity: event Recovered(address token, uint256 amount)
func (_JuicedStakingRewards *JuicedStakingRewardsFilterer) FilterRecovered(opts *bind.FilterOpts) (*JuicedStakingRewardsRecoveredIterator, error) {

	logs, sub, err := _JuicedStakingRewards.contract.FilterLogs(opts, "Recovered")
	if err != nil {
		return nil, err
	}
	return &JuicedStakingRewardsRecoveredIterator{contract: _JuicedStakingRewards.contract, event: "Recovered", logs: logs, sub: sub}, nil
}

// WatchRecovered is a free log subscription operation binding the contract event 0x8c1256b8896378cd5044f80c202f9772b9d77dc85c8a6eb51967210b09bfaa28.
//
// Solidity: event Recovered(address token, uint256 amount)
func (_JuicedStakingRewards *JuicedStakingRewardsFilterer) WatchRecovered(opts *bind.WatchOpts, sink chan<- *JuicedStakingRewardsRecovered) (event.Subscription, error) {

	logs, sub, err := _JuicedStakingRewards.contract.WatchLogs(opts, "Recovered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JuicedStakingRewardsRecovered)
				if err := _JuicedStakingRewards.contract.UnpackLog(event, "Recovered", log); err != nil {
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
func (_JuicedStakingRewards *JuicedStakingRewardsFilterer) ParseRecovered(log types.Log) (*JuicedStakingRewardsRecovered, error) {
	event := new(JuicedStakingRewardsRecovered)
	if err := _JuicedStakingRewards.contract.UnpackLog(event, "Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JuicedStakingRewardsRewardAddedIterator is returned from FilterRewardAdded and is used to iterate over the raw logs and unpacked data for RewardAdded events raised by the JuicedStakingRewards contract.
type JuicedStakingRewardsRewardAddedIterator struct {
	Event *JuicedStakingRewardsRewardAdded // Event containing the contract specifics and raw log

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
func (it *JuicedStakingRewardsRewardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JuicedStakingRewardsRewardAdded)
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
		it.Event = new(JuicedStakingRewardsRewardAdded)
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
func (it *JuicedStakingRewardsRewardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JuicedStakingRewardsRewardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JuicedStakingRewardsRewardAdded represents a RewardAdded event raised by the JuicedStakingRewards contract.
type JuicedStakingRewardsRewardAdded struct {
	RewardToken common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardAdded is a free log retrieval operation binding the contract event 0xac24935fd910bc682b5ccb1a07b718cadf8cf2f6d1404c4f3ddc3662dae40e29.
//
// Solidity: event RewardAdded(address indexed rewardToken, uint256 amount)
func (_JuicedStakingRewards *JuicedStakingRewardsFilterer) FilterRewardAdded(opts *bind.FilterOpts, rewardToken []common.Address) (*JuicedStakingRewardsRewardAddedIterator, error) {

	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}

	logs, sub, err := _JuicedStakingRewards.contract.FilterLogs(opts, "RewardAdded", rewardTokenRule)
	if err != nil {
		return nil, err
	}
	return &JuicedStakingRewardsRewardAddedIterator{contract: _JuicedStakingRewards.contract, event: "RewardAdded", logs: logs, sub: sub}, nil
}

// WatchRewardAdded is a free log subscription operation binding the contract event 0xac24935fd910bc682b5ccb1a07b718cadf8cf2f6d1404c4f3ddc3662dae40e29.
//
// Solidity: event RewardAdded(address indexed rewardToken, uint256 amount)
func (_JuicedStakingRewards *JuicedStakingRewardsFilterer) WatchRewardAdded(opts *bind.WatchOpts, sink chan<- *JuicedStakingRewardsRewardAdded, rewardToken []common.Address) (event.Subscription, error) {

	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}

	logs, sub, err := _JuicedStakingRewards.contract.WatchLogs(opts, "RewardAdded", rewardTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JuicedStakingRewardsRewardAdded)
				if err := _JuicedStakingRewards.contract.UnpackLog(event, "RewardAdded", log); err != nil {
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

// ParseRewardAdded is a log parse operation binding the contract event 0xac24935fd910bc682b5ccb1a07b718cadf8cf2f6d1404c4f3ddc3662dae40e29.
//
// Solidity: event RewardAdded(address indexed rewardToken, uint256 amount)
func (_JuicedStakingRewards *JuicedStakingRewardsFilterer) ParseRewardAdded(log types.Log) (*JuicedStakingRewardsRewardAdded, error) {
	event := new(JuicedStakingRewardsRewardAdded)
	if err := _JuicedStakingRewards.contract.UnpackLog(event, "RewardAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JuicedStakingRewardsRewardPaidIterator is returned from FilterRewardPaid and is used to iterate over the raw logs and unpacked data for RewardPaid events raised by the JuicedStakingRewards contract.
type JuicedStakingRewardsRewardPaidIterator struct {
	Event *JuicedStakingRewardsRewardPaid // Event containing the contract specifics and raw log

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
func (it *JuicedStakingRewardsRewardPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JuicedStakingRewardsRewardPaid)
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
		it.Event = new(JuicedStakingRewardsRewardPaid)
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
func (it *JuicedStakingRewardsRewardPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JuicedStakingRewardsRewardPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JuicedStakingRewardsRewardPaid represents a RewardPaid event raised by the JuicedStakingRewards contract.
type JuicedStakingRewardsRewardPaid struct {
	User        common.Address
	RewardToken common.Address
	Reward      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardPaid is a free log retrieval operation binding the contract event 0x540798df468d7b23d11f156fdb954cb19ad414d150722a7b6d55ba369dea792e.
//
// Solidity: event RewardPaid(address indexed user, address indexed rewardToken, uint256 reward)
func (_JuicedStakingRewards *JuicedStakingRewardsFilterer) FilterRewardPaid(opts *bind.FilterOpts, user []common.Address, rewardToken []common.Address) (*JuicedStakingRewardsRewardPaidIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}

	logs, sub, err := _JuicedStakingRewards.contract.FilterLogs(opts, "RewardPaid", userRule, rewardTokenRule)
	if err != nil {
		return nil, err
	}
	return &JuicedStakingRewardsRewardPaidIterator{contract: _JuicedStakingRewards.contract, event: "RewardPaid", logs: logs, sub: sub}, nil
}

// WatchRewardPaid is a free log subscription operation binding the contract event 0x540798df468d7b23d11f156fdb954cb19ad414d150722a7b6d55ba369dea792e.
//
// Solidity: event RewardPaid(address indexed user, address indexed rewardToken, uint256 reward)
func (_JuicedStakingRewards *JuicedStakingRewardsFilterer) WatchRewardPaid(opts *bind.WatchOpts, sink chan<- *JuicedStakingRewardsRewardPaid, user []common.Address, rewardToken []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}

	logs, sub, err := _JuicedStakingRewards.contract.WatchLogs(opts, "RewardPaid", userRule, rewardTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JuicedStakingRewardsRewardPaid)
				if err := _JuicedStakingRewards.contract.UnpackLog(event, "RewardPaid", log); err != nil {
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

// ParseRewardPaid is a log parse operation binding the contract event 0x540798df468d7b23d11f156fdb954cb19ad414d150722a7b6d55ba369dea792e.
//
// Solidity: event RewardPaid(address indexed user, address indexed rewardToken, uint256 reward)
func (_JuicedStakingRewards *JuicedStakingRewardsFilterer) ParseRewardPaid(log types.Log) (*JuicedStakingRewardsRewardPaid, error) {
	event := new(JuicedStakingRewardsRewardPaid)
	if err := _JuicedStakingRewards.contract.UnpackLog(event, "RewardPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JuicedStakingRewardsRewardsDurationUpdatedIterator is returned from FilterRewardsDurationUpdated and is used to iterate over the raw logs and unpacked data for RewardsDurationUpdated events raised by the JuicedStakingRewards contract.
type JuicedStakingRewardsRewardsDurationUpdatedIterator struct {
	Event *JuicedStakingRewardsRewardsDurationUpdated // Event containing the contract specifics and raw log

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
func (it *JuicedStakingRewardsRewardsDurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JuicedStakingRewardsRewardsDurationUpdated)
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
		it.Event = new(JuicedStakingRewardsRewardsDurationUpdated)
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
func (it *JuicedStakingRewardsRewardsDurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JuicedStakingRewardsRewardsDurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JuicedStakingRewardsRewardsDurationUpdated represents a RewardsDurationUpdated event raised by the JuicedStakingRewards contract.
type JuicedStakingRewardsRewardsDurationUpdated struct {
	Token       common.Address
	NewDuration *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardsDurationUpdated is a free log retrieval operation binding the contract event 0xad2f86b01ed93b4b3a150d448c61a4f5d8d38075d3c0c64cc0a26fd6e1f49545.
//
// Solidity: event RewardsDurationUpdated(address token, uint256 newDuration)
func (_JuicedStakingRewards *JuicedStakingRewardsFilterer) FilterRewardsDurationUpdated(opts *bind.FilterOpts) (*JuicedStakingRewardsRewardsDurationUpdatedIterator, error) {

	logs, sub, err := _JuicedStakingRewards.contract.FilterLogs(opts, "RewardsDurationUpdated")
	if err != nil {
		return nil, err
	}
	return &JuicedStakingRewardsRewardsDurationUpdatedIterator{contract: _JuicedStakingRewards.contract, event: "RewardsDurationUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardsDurationUpdated is a free log subscription operation binding the contract event 0xad2f86b01ed93b4b3a150d448c61a4f5d8d38075d3c0c64cc0a26fd6e1f49545.
//
// Solidity: event RewardsDurationUpdated(address token, uint256 newDuration)
func (_JuicedStakingRewards *JuicedStakingRewardsFilterer) WatchRewardsDurationUpdated(opts *bind.WatchOpts, sink chan<- *JuicedStakingRewardsRewardsDurationUpdated) (event.Subscription, error) {

	logs, sub, err := _JuicedStakingRewards.contract.WatchLogs(opts, "RewardsDurationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JuicedStakingRewardsRewardsDurationUpdated)
				if err := _JuicedStakingRewards.contract.UnpackLog(event, "RewardsDurationUpdated", log); err != nil {
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

// ParseRewardsDurationUpdated is a log parse operation binding the contract event 0xad2f86b01ed93b4b3a150d448c61a4f5d8d38075d3c0c64cc0a26fd6e1f49545.
//
// Solidity: event RewardsDurationUpdated(address token, uint256 newDuration)
func (_JuicedStakingRewards *JuicedStakingRewardsFilterer) ParseRewardsDurationUpdated(log types.Log) (*JuicedStakingRewardsRewardsDurationUpdated, error) {
	event := new(JuicedStakingRewardsRewardsDurationUpdated)
	if err := _JuicedStakingRewards.contract.UnpackLog(event, "RewardsDurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JuicedStakingRewardsStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the JuicedStakingRewards contract.
type JuicedStakingRewardsStakedIterator struct {
	Event *JuicedStakingRewardsStaked // Event containing the contract specifics and raw log

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
func (it *JuicedStakingRewardsStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JuicedStakingRewardsStaked)
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
		it.Event = new(JuicedStakingRewardsStaked)
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
func (it *JuicedStakingRewardsStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JuicedStakingRewardsStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JuicedStakingRewardsStaked represents a Staked event raised by the JuicedStakingRewards contract.
type JuicedStakingRewardsStaked struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_JuicedStakingRewards *JuicedStakingRewardsFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address) (*JuicedStakingRewardsStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _JuicedStakingRewards.contract.FilterLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return &JuicedStakingRewardsStakedIterator{contract: _JuicedStakingRewards.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_JuicedStakingRewards *JuicedStakingRewardsFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *JuicedStakingRewardsStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _JuicedStakingRewards.contract.WatchLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JuicedStakingRewardsStaked)
				if err := _JuicedStakingRewards.contract.UnpackLog(event, "Staked", log); err != nil {
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
func (_JuicedStakingRewards *JuicedStakingRewardsFilterer) ParseStaked(log types.Log) (*JuicedStakingRewardsStaked, error) {
	event := new(JuicedStakingRewardsStaked)
	if err := _JuicedStakingRewards.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JuicedStakingRewardsStakedForIterator is returned from FilterStakedFor and is used to iterate over the raw logs and unpacked data for StakedFor events raised by the JuicedStakingRewards contract.
type JuicedStakingRewardsStakedForIterator struct {
	Event *JuicedStakingRewardsStakedFor // Event containing the contract specifics and raw log

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
func (it *JuicedStakingRewardsStakedForIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JuicedStakingRewardsStakedFor)
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
		it.Event = new(JuicedStakingRewardsStakedFor)
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
func (it *JuicedStakingRewardsStakedForIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JuicedStakingRewardsStakedForIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JuicedStakingRewardsStakedFor represents a StakedFor event raised by the JuicedStakingRewards contract.
type JuicedStakingRewardsStakedFor struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakedFor is a free log retrieval operation binding the contract event 0xd185ae938da574e9cd1073962e1972c75ec585ab222b200a88c0abe2bf0cfe67.
//
// Solidity: event StakedFor(address indexed user, uint256 amount)
func (_JuicedStakingRewards *JuicedStakingRewardsFilterer) FilterStakedFor(opts *bind.FilterOpts, user []common.Address) (*JuicedStakingRewardsStakedForIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _JuicedStakingRewards.contract.FilterLogs(opts, "StakedFor", userRule)
	if err != nil {
		return nil, err
	}
	return &JuicedStakingRewardsStakedForIterator{contract: _JuicedStakingRewards.contract, event: "StakedFor", logs: logs, sub: sub}, nil
}

// WatchStakedFor is a free log subscription operation binding the contract event 0xd185ae938da574e9cd1073962e1972c75ec585ab222b200a88c0abe2bf0cfe67.
//
// Solidity: event StakedFor(address indexed user, uint256 amount)
func (_JuicedStakingRewards *JuicedStakingRewardsFilterer) WatchStakedFor(opts *bind.WatchOpts, sink chan<- *JuicedStakingRewardsStakedFor, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _JuicedStakingRewards.contract.WatchLogs(opts, "StakedFor", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JuicedStakingRewardsStakedFor)
				if err := _JuicedStakingRewards.contract.UnpackLog(event, "StakedFor", log); err != nil {
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
func (_JuicedStakingRewards *JuicedStakingRewardsFilterer) ParseStakedFor(log types.Log) (*JuicedStakingRewardsStakedFor, error) {
	event := new(JuicedStakingRewardsStakedFor)
	if err := _JuicedStakingRewards.contract.UnpackLog(event, "StakedFor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JuicedStakingRewardsUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the JuicedStakingRewards contract.
type JuicedStakingRewardsUnpausedIterator struct {
	Event *JuicedStakingRewardsUnpaused // Event containing the contract specifics and raw log

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
func (it *JuicedStakingRewardsUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JuicedStakingRewardsUnpaused)
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
		it.Event = new(JuicedStakingRewardsUnpaused)
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
func (it *JuicedStakingRewardsUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JuicedStakingRewardsUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JuicedStakingRewardsUnpaused represents a Unpaused event raised by the JuicedStakingRewards contract.
type JuicedStakingRewardsUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_JuicedStakingRewards *JuicedStakingRewardsFilterer) FilterUnpaused(opts *bind.FilterOpts) (*JuicedStakingRewardsUnpausedIterator, error) {

	logs, sub, err := _JuicedStakingRewards.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &JuicedStakingRewardsUnpausedIterator{contract: _JuicedStakingRewards.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_JuicedStakingRewards *JuicedStakingRewardsFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *JuicedStakingRewardsUnpaused) (event.Subscription, error) {

	logs, sub, err := _JuicedStakingRewards.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JuicedStakingRewardsUnpaused)
				if err := _JuicedStakingRewards.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_JuicedStakingRewards *JuicedStakingRewardsFilterer) ParseUnpaused(log types.Log) (*JuicedStakingRewardsUnpaused, error) {
	event := new(JuicedStakingRewardsUnpaused)
	if err := _JuicedStakingRewards.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JuicedStakingRewardsWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the JuicedStakingRewards contract.
type JuicedStakingRewardsWithdrawnIterator struct {
	Event *JuicedStakingRewardsWithdrawn // Event containing the contract specifics and raw log

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
func (it *JuicedStakingRewardsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JuicedStakingRewardsWithdrawn)
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
		it.Event = new(JuicedStakingRewardsWithdrawn)
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
func (it *JuicedStakingRewardsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JuicedStakingRewardsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JuicedStakingRewardsWithdrawn represents a Withdrawn event raised by the JuicedStakingRewards contract.
type JuicedStakingRewardsWithdrawn struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_JuicedStakingRewards *JuicedStakingRewardsFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address) (*JuicedStakingRewardsWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _JuicedStakingRewards.contract.FilterLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &JuicedStakingRewardsWithdrawnIterator{contract: _JuicedStakingRewards.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_JuicedStakingRewards *JuicedStakingRewardsFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *JuicedStakingRewardsWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _JuicedStakingRewards.contract.WatchLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JuicedStakingRewardsWithdrawn)
				if err := _JuicedStakingRewards.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
func (_JuicedStakingRewards *JuicedStakingRewardsFilterer) ParseWithdrawn(log types.Log) (*JuicedStakingRewardsWithdrawn, error) {
	event := new(JuicedStakingRewardsWithdrawn)
	if err := _JuicedStakingRewards.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JuicedStakingRewardsWithdrawnForIterator is returned from FilterWithdrawnFor and is used to iterate over the raw logs and unpacked data for WithdrawnFor events raised by the JuicedStakingRewards contract.
type JuicedStakingRewardsWithdrawnForIterator struct {
	Event *JuicedStakingRewardsWithdrawnFor // Event containing the contract specifics and raw log

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
func (it *JuicedStakingRewardsWithdrawnForIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JuicedStakingRewardsWithdrawnFor)
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
		it.Event = new(JuicedStakingRewardsWithdrawnFor)
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
func (it *JuicedStakingRewardsWithdrawnForIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JuicedStakingRewardsWithdrawnForIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JuicedStakingRewardsWithdrawnFor represents a WithdrawnFor event raised by the JuicedStakingRewards contract.
type JuicedStakingRewardsWithdrawnFor struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnFor is a free log retrieval operation binding the contract event 0x04de67fe258683d541c77f7275134560914ac5eb0fe853df87c0946439c2e706.
//
// Solidity: event WithdrawnFor(address indexed user, uint256 amount)
func (_JuicedStakingRewards *JuicedStakingRewardsFilterer) FilterWithdrawnFor(opts *bind.FilterOpts, user []common.Address) (*JuicedStakingRewardsWithdrawnForIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _JuicedStakingRewards.contract.FilterLogs(opts, "WithdrawnFor", userRule)
	if err != nil {
		return nil, err
	}
	return &JuicedStakingRewardsWithdrawnForIterator{contract: _JuicedStakingRewards.contract, event: "WithdrawnFor", logs: logs, sub: sub}, nil
}

// WatchWithdrawnFor is a free log subscription operation binding the contract event 0x04de67fe258683d541c77f7275134560914ac5eb0fe853df87c0946439c2e706.
//
// Solidity: event WithdrawnFor(address indexed user, uint256 amount)
func (_JuicedStakingRewards *JuicedStakingRewardsFilterer) WatchWithdrawnFor(opts *bind.WatchOpts, sink chan<- *JuicedStakingRewardsWithdrawnFor, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _JuicedStakingRewards.contract.WatchLogs(opts, "WithdrawnFor", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JuicedStakingRewardsWithdrawnFor)
				if err := _JuicedStakingRewards.contract.UnpackLog(event, "WithdrawnFor", log); err != nil {
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

// ParseWithdrawnFor is a log parse operation binding the contract event 0x04de67fe258683d541c77f7275134560914ac5eb0fe853df87c0946439c2e706.
//
// Solidity: event WithdrawnFor(address indexed user, uint256 amount)
func (_JuicedStakingRewards *JuicedStakingRewardsFilterer) ParseWithdrawnFor(log types.Log) (*JuicedStakingRewardsWithdrawnFor, error) {
	event := new(JuicedStakingRewardsWithdrawnFor)
	if err := _JuicedStakingRewards.contract.UnpackLog(event, "WithdrawnFor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JuicedStakingRewardsZapContractUpdatedIterator is returned from FilterZapContractUpdated and is used to iterate over the raw logs and unpacked data for ZapContractUpdated events raised by the JuicedStakingRewards contract.
type JuicedStakingRewardsZapContractUpdatedIterator struct {
	Event *JuicedStakingRewardsZapContractUpdated // Event containing the contract specifics and raw log

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
func (it *JuicedStakingRewardsZapContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JuicedStakingRewardsZapContractUpdated)
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
		it.Event = new(JuicedStakingRewardsZapContractUpdated)
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
func (it *JuicedStakingRewardsZapContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JuicedStakingRewardsZapContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JuicedStakingRewardsZapContractUpdated represents a ZapContractUpdated event raised by the JuicedStakingRewards contract.
type JuicedStakingRewardsZapContractUpdated struct {
	ZapContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterZapContractUpdated is a free log retrieval operation binding the contract event 0x0da458581861b4ff3ae5c23a916ec5218bbf4c371497cec29006f5c55d9c9a84.
//
// Solidity: event ZapContractUpdated(address _zapContract)
func (_JuicedStakingRewards *JuicedStakingRewardsFilterer) FilterZapContractUpdated(opts *bind.FilterOpts) (*JuicedStakingRewardsZapContractUpdatedIterator, error) {

	logs, sub, err := _JuicedStakingRewards.contract.FilterLogs(opts, "ZapContractUpdated")
	if err != nil {
		return nil, err
	}
	return &JuicedStakingRewardsZapContractUpdatedIterator{contract: _JuicedStakingRewards.contract, event: "ZapContractUpdated", logs: logs, sub: sub}, nil
}

// WatchZapContractUpdated is a free log subscription operation binding the contract event 0x0da458581861b4ff3ae5c23a916ec5218bbf4c371497cec29006f5c55d9c9a84.
//
// Solidity: event ZapContractUpdated(address _zapContract)
func (_JuicedStakingRewards *JuicedStakingRewardsFilterer) WatchZapContractUpdated(opts *bind.WatchOpts, sink chan<- *JuicedStakingRewardsZapContractUpdated) (event.Subscription, error) {

	logs, sub, err := _JuicedStakingRewards.contract.WatchLogs(opts, "ZapContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JuicedStakingRewardsZapContractUpdated)
				if err := _JuicedStakingRewards.contract.UnpackLog(event, "ZapContractUpdated", log); err != nil {
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
func (_JuicedStakingRewards *JuicedStakingRewardsFilterer) ParseZapContractUpdated(log types.Log) (*JuicedStakingRewardsZapContractUpdated, error) {
	event := new(JuicedStakingRewardsZapContractUpdated)
	if err := _JuicedStakingRewards.contract.UnpackLog(event, "ZapContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
