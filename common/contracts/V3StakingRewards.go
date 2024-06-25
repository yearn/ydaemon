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

// V3StakingRewardsMetaData contains all meta data concerning the V3StakingRewards contract.
var V3StakingRewardsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakingToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_zapContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"clone\",\"type\":\"address\"}],\"name\":\"Cloned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Recovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDuration\",\"type\":\"uint256\"}],\"name\":\"RewardsDurationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakedFor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawnFor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_zapContract\",\"type\":\"address\"}],\"name\":\"ZapContractUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardsDistributor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewardsDuration\",\"type\":\"uint256\"}],\"name\":\"addReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"balanceOfUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakingToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_zapContract\",\"type\":\"address\"}],\"name\":\"cloneStakingPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newStakingPool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"}],\"name\":\"earned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pending\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"earnedMulti\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"pending\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"}],\"name\":\"getOneReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"}],\"name\":\"getRewardForDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakingToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_zapContract\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOriginal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isRetired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"}],\"name\":\"lastTimeRewardApplicable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewardAmount\",\"type\":\"uint256\"}],\"name\":\"notifyRewardAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenAmount\",\"type\":\"uint256\"}],\"name\":\"recoverERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardData\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"rewardsDistributor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rewardsDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"periodFinish\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdateTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardPerTokenStored\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"}],\"name\":\"rewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewardTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardTokensLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardsDistributor\",\"type\":\"address\"}],\"name\":\"setRewardsDistributor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewardsDuration\",\"type\":\"uint256\"}],\"name\":\"setRewardsDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_zapContract\",\"type\":\"address\"}],\"name\":\"setZapContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"stakeFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRewardPerTokenPaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_exit\",\"type\":\"bool\"}],\"name\":\"withdrawFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zapContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// V3StakingRewardsABI is the input ABI used to generate the binding from.
// Deprecated: Use V3StakingRewardsMetaData.ABI instead.
var V3StakingRewardsABI = V3StakingRewardsMetaData.ABI

// V3StakingRewards is an auto generated Go binding around an Ethereum contract.
type V3StakingRewards struct {
	V3StakingRewardsCaller     // Read-only binding to the contract
	V3StakingRewardsTransactor // Write-only binding to the contract
	V3StakingRewardsFilterer   // Log filterer for contract events
}

// V3StakingRewardsCaller is an auto generated read-only Go binding around an Ethereum contract.
type V3StakingRewardsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V3StakingRewardsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type V3StakingRewardsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V3StakingRewardsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type V3StakingRewardsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V3StakingRewardsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type V3StakingRewardsSession struct {
	Contract     *V3StakingRewards // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// V3StakingRewardsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type V3StakingRewardsCallerSession struct {
	Contract *V3StakingRewardsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// V3StakingRewardsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type V3StakingRewardsTransactorSession struct {
	Contract     *V3StakingRewardsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// V3StakingRewardsRaw is an auto generated low-level Go binding around an Ethereum contract.
type V3StakingRewardsRaw struct {
	Contract *V3StakingRewards // Generic contract binding to access the raw methods on
}

// V3StakingRewardsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type V3StakingRewardsCallerRaw struct {
	Contract *V3StakingRewardsCaller // Generic read-only contract binding to access the raw methods on
}

// V3StakingRewardsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type V3StakingRewardsTransactorRaw struct {
	Contract *V3StakingRewardsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewV3StakingRewards creates a new instance of V3StakingRewards, bound to a specific deployed contract.
func NewV3StakingRewards(address common.Address, backend bind.ContractBackend) (*V3StakingRewards, error) {
	contract, err := bindV3StakingRewards(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &V3StakingRewards{V3StakingRewardsCaller: V3StakingRewardsCaller{contract: contract}, V3StakingRewardsTransactor: V3StakingRewardsTransactor{contract: contract}, V3StakingRewardsFilterer: V3StakingRewardsFilterer{contract: contract}}, nil
}

// NewV3StakingRewardsCaller creates a new read-only instance of V3StakingRewards, bound to a specific deployed contract.
func NewV3StakingRewardsCaller(address common.Address, caller bind.ContractCaller) (*V3StakingRewardsCaller, error) {
	contract, err := bindV3StakingRewards(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &V3StakingRewardsCaller{contract: contract}, nil
}

// NewV3StakingRewardsTransactor creates a new write-only instance of V3StakingRewards, bound to a specific deployed contract.
func NewV3StakingRewardsTransactor(address common.Address, transactor bind.ContractTransactor) (*V3StakingRewardsTransactor, error) {
	contract, err := bindV3StakingRewards(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &V3StakingRewardsTransactor{contract: contract}, nil
}

// NewV3StakingRewardsFilterer creates a new log filterer instance of V3StakingRewards, bound to a specific deployed contract.
func NewV3StakingRewardsFilterer(address common.Address, filterer bind.ContractFilterer) (*V3StakingRewardsFilterer, error) {
	contract, err := bindV3StakingRewards(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &V3StakingRewardsFilterer{contract: contract}, nil
}

// bindV3StakingRewards binds a generic wrapper to an already deployed contract.
func bindV3StakingRewards(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := V3StakingRewardsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V3StakingRewards *V3StakingRewardsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V3StakingRewards.Contract.V3StakingRewardsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V3StakingRewards *V3StakingRewardsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V3StakingRewards.Contract.V3StakingRewardsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V3StakingRewards *V3StakingRewardsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V3StakingRewards.Contract.V3StakingRewardsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V3StakingRewards *V3StakingRewardsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V3StakingRewards.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V3StakingRewards *V3StakingRewardsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V3StakingRewards.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V3StakingRewards *V3StakingRewardsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V3StakingRewards.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_V3StakingRewards *V3StakingRewardsCaller) BalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _V3StakingRewards.contract.Call(opts, &out, "balanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_V3StakingRewards *V3StakingRewardsSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _V3StakingRewards.Contract.BalanceOf(&_V3StakingRewards.CallOpts, _account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_V3StakingRewards *V3StakingRewardsCallerSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _V3StakingRewards.Contract.BalanceOf(&_V3StakingRewards.CallOpts, _account)
}

// BalanceOfUnderlying is a free data retrieval call binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address _account) view returns(uint256)
func (_V3StakingRewards *V3StakingRewardsCaller) BalanceOfUnderlying(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _V3StakingRewards.contract.Call(opts, &out, "balanceOfUnderlying", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfUnderlying is a free data retrieval call binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address _account) view returns(uint256)
func (_V3StakingRewards *V3StakingRewardsSession) BalanceOfUnderlying(_account common.Address) (*big.Int, error) {
	return _V3StakingRewards.Contract.BalanceOfUnderlying(&_V3StakingRewards.CallOpts, _account)
}

// BalanceOfUnderlying is a free data retrieval call binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address _account) view returns(uint256)
func (_V3StakingRewards *V3StakingRewardsCallerSession) BalanceOfUnderlying(_account common.Address) (*big.Int, error) {
	return _V3StakingRewards.Contract.BalanceOfUnderlying(&_V3StakingRewards.CallOpts, _account)
}

// Earned is a free data retrieval call binding the contract method 0x211dc32d.
//
// Solidity: function earned(address _account, address _rewardsToken) view returns(uint256 pending)
func (_V3StakingRewards *V3StakingRewardsCaller) Earned(opts *bind.CallOpts, _account common.Address, _rewardsToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _V3StakingRewards.contract.Call(opts, &out, "earned", _account, _rewardsToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Earned is a free data retrieval call binding the contract method 0x211dc32d.
//
// Solidity: function earned(address _account, address _rewardsToken) view returns(uint256 pending)
func (_V3StakingRewards *V3StakingRewardsSession) Earned(_account common.Address, _rewardsToken common.Address) (*big.Int, error) {
	return _V3StakingRewards.Contract.Earned(&_V3StakingRewards.CallOpts, _account, _rewardsToken)
}

// Earned is a free data retrieval call binding the contract method 0x211dc32d.
//
// Solidity: function earned(address _account, address _rewardsToken) view returns(uint256 pending)
func (_V3StakingRewards *V3StakingRewardsCallerSession) Earned(_account common.Address, _rewardsToken common.Address) (*big.Int, error) {
	return _V3StakingRewards.Contract.Earned(&_V3StakingRewards.CallOpts, _account, _rewardsToken)
}

// EarnedMulti is a free data retrieval call binding the contract method 0xff342e09.
//
// Solidity: function earnedMulti(address _account) view returns(uint256[] pending)
func (_V3StakingRewards *V3StakingRewardsCaller) EarnedMulti(opts *bind.CallOpts, _account common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _V3StakingRewards.contract.Call(opts, &out, "earnedMulti", _account)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// EarnedMulti is a free data retrieval call binding the contract method 0xff342e09.
//
// Solidity: function earnedMulti(address _account) view returns(uint256[] pending)
func (_V3StakingRewards *V3StakingRewardsSession) EarnedMulti(_account common.Address) ([]*big.Int, error) {
	return _V3StakingRewards.Contract.EarnedMulti(&_V3StakingRewards.CallOpts, _account)
}

// EarnedMulti is a free data retrieval call binding the contract method 0xff342e09.
//
// Solidity: function earnedMulti(address _account) view returns(uint256[] pending)
func (_V3StakingRewards *V3StakingRewardsCallerSession) EarnedMulti(_account common.Address) ([]*big.Int, error) {
	return _V3StakingRewards.Contract.EarnedMulti(&_V3StakingRewards.CallOpts, _account)
}

// GetRewardForDuration is a free data retrieval call binding the contract method 0xbcd11014.
//
// Solidity: function getRewardForDuration(address _rewardsToken) view returns(uint256)
func (_V3StakingRewards *V3StakingRewardsCaller) GetRewardForDuration(opts *bind.CallOpts, _rewardsToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _V3StakingRewards.contract.Call(opts, &out, "getRewardForDuration", _rewardsToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardForDuration is a free data retrieval call binding the contract method 0xbcd11014.
//
// Solidity: function getRewardForDuration(address _rewardsToken) view returns(uint256)
func (_V3StakingRewards *V3StakingRewardsSession) GetRewardForDuration(_rewardsToken common.Address) (*big.Int, error) {
	return _V3StakingRewards.Contract.GetRewardForDuration(&_V3StakingRewards.CallOpts, _rewardsToken)
}

// GetRewardForDuration is a free data retrieval call binding the contract method 0xbcd11014.
//
// Solidity: function getRewardForDuration(address _rewardsToken) view returns(uint256)
func (_V3StakingRewards *V3StakingRewardsCallerSession) GetRewardForDuration(_rewardsToken common.Address) (*big.Int, error) {
	return _V3StakingRewards.Contract.GetRewardForDuration(&_V3StakingRewards.CallOpts, _rewardsToken)
}

// IsOriginal is a free data retrieval call binding the contract method 0x6f392ce7.
//
// Solidity: function isOriginal() view returns(bool)
func (_V3StakingRewards *V3StakingRewardsCaller) IsOriginal(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _V3StakingRewards.contract.Call(opts, &out, "isOriginal")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOriginal is a free data retrieval call binding the contract method 0x6f392ce7.
//
// Solidity: function isOriginal() view returns(bool)
func (_V3StakingRewards *V3StakingRewardsSession) IsOriginal() (bool, error) {
	return _V3StakingRewards.Contract.IsOriginal(&_V3StakingRewards.CallOpts)
}

// IsOriginal is a free data retrieval call binding the contract method 0x6f392ce7.
//
// Solidity: function isOriginal() view returns(bool)
func (_V3StakingRewards *V3StakingRewardsCallerSession) IsOriginal() (bool, error) {
	return _V3StakingRewards.Contract.IsOriginal(&_V3StakingRewards.CallOpts)
}

// IsRetired is a free data retrieval call binding the contract method 0x49343cee.
//
// Solidity: function isRetired() view returns(bool)
func (_V3StakingRewards *V3StakingRewardsCaller) IsRetired(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _V3StakingRewards.contract.Call(opts, &out, "isRetired")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRetired is a free data retrieval call binding the contract method 0x49343cee.
//
// Solidity: function isRetired() view returns(bool)
func (_V3StakingRewards *V3StakingRewardsSession) IsRetired() (bool, error) {
	return _V3StakingRewards.Contract.IsRetired(&_V3StakingRewards.CallOpts)
}

// IsRetired is a free data retrieval call binding the contract method 0x49343cee.
//
// Solidity: function isRetired() view returns(bool)
func (_V3StakingRewards *V3StakingRewardsCallerSession) IsRetired() (bool, error) {
	return _V3StakingRewards.Contract.IsRetired(&_V3StakingRewards.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x638634ee.
//
// Solidity: function lastTimeRewardApplicable(address _rewardsToken) view returns(uint256)
func (_V3StakingRewards *V3StakingRewardsCaller) LastTimeRewardApplicable(opts *bind.CallOpts, _rewardsToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _V3StakingRewards.contract.Call(opts, &out, "lastTimeRewardApplicable", _rewardsToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x638634ee.
//
// Solidity: function lastTimeRewardApplicable(address _rewardsToken) view returns(uint256)
func (_V3StakingRewards *V3StakingRewardsSession) LastTimeRewardApplicable(_rewardsToken common.Address) (*big.Int, error) {
	return _V3StakingRewards.Contract.LastTimeRewardApplicable(&_V3StakingRewards.CallOpts, _rewardsToken)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x638634ee.
//
// Solidity: function lastTimeRewardApplicable(address _rewardsToken) view returns(uint256)
func (_V3StakingRewards *V3StakingRewardsCallerSession) LastTimeRewardApplicable(_rewardsToken common.Address) (*big.Int, error) {
	return _V3StakingRewards.Contract.LastTimeRewardApplicable(&_V3StakingRewards.CallOpts, _rewardsToken)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_V3StakingRewards *V3StakingRewardsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V3StakingRewards.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_V3StakingRewards *V3StakingRewardsSession) Owner() (common.Address, error) {
	return _V3StakingRewards.Contract.Owner(&_V3StakingRewards.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_V3StakingRewards *V3StakingRewardsCallerSession) Owner() (common.Address, error) {
	return _V3StakingRewards.Contract.Owner(&_V3StakingRewards.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_V3StakingRewards *V3StakingRewardsCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _V3StakingRewards.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_V3StakingRewards *V3StakingRewardsSession) Paused() (bool, error) {
	return _V3StakingRewards.Contract.Paused(&_V3StakingRewards.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_V3StakingRewards *V3StakingRewardsCallerSession) Paused() (bool, error) {
	return _V3StakingRewards.Contract.Paused(&_V3StakingRewards.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_V3StakingRewards *V3StakingRewardsCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V3StakingRewards.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_V3StakingRewards *V3StakingRewardsSession) PendingOwner() (common.Address, error) {
	return _V3StakingRewards.Contract.PendingOwner(&_V3StakingRewards.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_V3StakingRewards *V3StakingRewardsCallerSession) PendingOwner() (common.Address, error) {
	return _V3StakingRewards.Contract.PendingOwner(&_V3StakingRewards.CallOpts)
}

// RewardData is a free data retrieval call binding the contract method 0x48e5d9f8.
//
// Solidity: function rewardData(address ) view returns(address rewardsDistributor, uint256 rewardsDuration, uint256 periodFinish, uint256 rewardRate, uint256 lastUpdateTime, uint256 rewardPerTokenStored)
func (_V3StakingRewards *V3StakingRewardsCaller) RewardData(opts *bind.CallOpts, arg0 common.Address) (struct {
	RewardsDistributor   common.Address
	RewardsDuration      *big.Int
	PeriodFinish         *big.Int
	RewardRate           *big.Int
	LastUpdateTime       *big.Int
	RewardPerTokenStored *big.Int
}, error) {
	var out []interface{}
	err := _V3StakingRewards.contract.Call(opts, &out, "rewardData", arg0)

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
func (_V3StakingRewards *V3StakingRewardsSession) RewardData(arg0 common.Address) (struct {
	RewardsDistributor   common.Address
	RewardsDuration      *big.Int
	PeriodFinish         *big.Int
	RewardRate           *big.Int
	LastUpdateTime       *big.Int
	RewardPerTokenStored *big.Int
}, error) {
	return _V3StakingRewards.Contract.RewardData(&_V3StakingRewards.CallOpts, arg0)
}

// RewardData is a free data retrieval call binding the contract method 0x48e5d9f8.
//
// Solidity: function rewardData(address ) view returns(address rewardsDistributor, uint256 rewardsDuration, uint256 periodFinish, uint256 rewardRate, uint256 lastUpdateTime, uint256 rewardPerTokenStored)
func (_V3StakingRewards *V3StakingRewardsCallerSession) RewardData(arg0 common.Address) (struct {
	RewardsDistributor   common.Address
	RewardsDuration      *big.Int
	PeriodFinish         *big.Int
	RewardRate           *big.Int
	LastUpdateTime       *big.Int
	RewardPerTokenStored *big.Int
}, error) {
	return _V3StakingRewards.Contract.RewardData(&_V3StakingRewards.CallOpts, arg0)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xf1229777.
//
// Solidity: function rewardPerToken(address _rewardsToken) view returns(uint256 rewardAmount)
func (_V3StakingRewards *V3StakingRewardsCaller) RewardPerToken(opts *bind.CallOpts, _rewardsToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _V3StakingRewards.contract.Call(opts, &out, "rewardPerToken", _rewardsToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerToken is a free data retrieval call binding the contract method 0xf1229777.
//
// Solidity: function rewardPerToken(address _rewardsToken) view returns(uint256 rewardAmount)
func (_V3StakingRewards *V3StakingRewardsSession) RewardPerToken(_rewardsToken common.Address) (*big.Int, error) {
	return _V3StakingRewards.Contract.RewardPerToken(&_V3StakingRewards.CallOpts, _rewardsToken)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xf1229777.
//
// Solidity: function rewardPerToken(address _rewardsToken) view returns(uint256 rewardAmount)
func (_V3StakingRewards *V3StakingRewardsCallerSession) RewardPerToken(_rewardsToken common.Address) (*big.Int, error) {
	return _V3StakingRewards.Contract.RewardPerToken(&_V3StakingRewards.CallOpts, _rewardsToken)
}

// RewardTokens is a free data retrieval call binding the contract method 0x7bb7bed1.
//
// Solidity: function rewardTokens(uint256 ) view returns(address)
func (_V3StakingRewards *V3StakingRewardsCaller) RewardTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _V3StakingRewards.contract.Call(opts, &out, "rewardTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardTokens is a free data retrieval call binding the contract method 0x7bb7bed1.
//
// Solidity: function rewardTokens(uint256 ) view returns(address)
func (_V3StakingRewards *V3StakingRewardsSession) RewardTokens(arg0 *big.Int) (common.Address, error) {
	return _V3StakingRewards.Contract.RewardTokens(&_V3StakingRewards.CallOpts, arg0)
}

// RewardTokens is a free data retrieval call binding the contract method 0x7bb7bed1.
//
// Solidity: function rewardTokens(uint256 ) view returns(address)
func (_V3StakingRewards *V3StakingRewardsCallerSession) RewardTokens(arg0 *big.Int) (common.Address, error) {
	return _V3StakingRewards.Contract.RewardTokens(&_V3StakingRewards.CallOpts, arg0)
}

// RewardTokensLength is a free data retrieval call binding the contract method 0xbf199e62.
//
// Solidity: function rewardTokensLength() view returns(uint256)
func (_V3StakingRewards *V3StakingRewardsCaller) RewardTokensLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _V3StakingRewards.contract.Call(opts, &out, "rewardTokensLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardTokensLength is a free data retrieval call binding the contract method 0xbf199e62.
//
// Solidity: function rewardTokensLength() view returns(uint256)
func (_V3StakingRewards *V3StakingRewardsSession) RewardTokensLength() (*big.Int, error) {
	return _V3StakingRewards.Contract.RewardTokensLength(&_V3StakingRewards.CallOpts)
}

// RewardTokensLength is a free data retrieval call binding the contract method 0xbf199e62.
//
// Solidity: function rewardTokensLength() view returns(uint256)
func (_V3StakingRewards *V3StakingRewardsCallerSession) RewardTokensLength() (*big.Int, error) {
	return _V3StakingRewards.Contract.RewardTokensLength(&_V3StakingRewards.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0xe70b9e27.
//
// Solidity: function rewards(address , address ) view returns(uint256)
func (_V3StakingRewards *V3StakingRewardsCaller) Rewards(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _V3StakingRewards.contract.Call(opts, &out, "rewards", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0xe70b9e27.
//
// Solidity: function rewards(address , address ) view returns(uint256)
func (_V3StakingRewards *V3StakingRewardsSession) Rewards(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _V3StakingRewards.Contract.Rewards(&_V3StakingRewards.CallOpts, arg0, arg1)
}

// Rewards is a free data retrieval call binding the contract method 0xe70b9e27.
//
// Solidity: function rewards(address , address ) view returns(uint256)
func (_V3StakingRewards *V3StakingRewardsCallerSession) Rewards(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _V3StakingRewards.Contract.Rewards(&_V3StakingRewards.CallOpts, arg0, arg1)
}

// StakerVersion is a free data retrieval call binding the contract method 0x0638ca2a.
//
// Solidity: function stakerVersion() view returns(string)
func (_V3StakingRewards *V3StakingRewardsCaller) StakerVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _V3StakingRewards.contract.Call(opts, &out, "stakerVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// StakerVersion is a free data retrieval call binding the contract method 0x0638ca2a.
//
// Solidity: function stakerVersion() view returns(string)
func (_V3StakingRewards *V3StakingRewardsSession) StakerVersion() (string, error) {
	return _V3StakingRewards.Contract.StakerVersion(&_V3StakingRewards.CallOpts)
}

// StakerVersion is a free data retrieval call binding the contract method 0x0638ca2a.
//
// Solidity: function stakerVersion() view returns(string)
func (_V3StakingRewards *V3StakingRewardsCallerSession) StakerVersion() (string, error) {
	return _V3StakingRewards.Contract.StakerVersion(&_V3StakingRewards.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_V3StakingRewards *V3StakingRewardsCaller) StakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V3StakingRewards.contract.Call(opts, &out, "stakingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_V3StakingRewards *V3StakingRewardsSession) StakingToken() (common.Address, error) {
	return _V3StakingRewards.Contract.StakingToken(&_V3StakingRewards.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_V3StakingRewards *V3StakingRewardsCallerSession) StakingToken() (common.Address, error) {
	return _V3StakingRewards.Contract.StakingToken(&_V3StakingRewards.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_V3StakingRewards *V3StakingRewardsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _V3StakingRewards.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_V3StakingRewards *V3StakingRewardsSession) TotalSupply() (*big.Int, error) {
	return _V3StakingRewards.Contract.TotalSupply(&_V3StakingRewards.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_V3StakingRewards *V3StakingRewardsCallerSession) TotalSupply() (*big.Int, error) {
	return _V3StakingRewards.Contract.TotalSupply(&_V3StakingRewards.CallOpts)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x7035ab98.
//
// Solidity: function userRewardPerTokenPaid(address , address ) view returns(uint256)
func (_V3StakingRewards *V3StakingRewardsCaller) UserRewardPerTokenPaid(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _V3StakingRewards.contract.Call(opts, &out, "userRewardPerTokenPaid", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x7035ab98.
//
// Solidity: function userRewardPerTokenPaid(address , address ) view returns(uint256)
func (_V3StakingRewards *V3StakingRewardsSession) UserRewardPerTokenPaid(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _V3StakingRewards.Contract.UserRewardPerTokenPaid(&_V3StakingRewards.CallOpts, arg0, arg1)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x7035ab98.
//
// Solidity: function userRewardPerTokenPaid(address , address ) view returns(uint256)
func (_V3StakingRewards *V3StakingRewardsCallerSession) UserRewardPerTokenPaid(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _V3StakingRewards.Contract.UserRewardPerTokenPaid(&_V3StakingRewards.CallOpts, arg0, arg1)
}

// ZapContract is a free data retrieval call binding the contract method 0x2549dad9.
//
// Solidity: function zapContract() view returns(address)
func (_V3StakingRewards *V3StakingRewardsCaller) ZapContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V3StakingRewards.contract.Call(opts, &out, "zapContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZapContract is a free data retrieval call binding the contract method 0x2549dad9.
//
// Solidity: function zapContract() view returns(address)
func (_V3StakingRewards *V3StakingRewardsSession) ZapContract() (common.Address, error) {
	return _V3StakingRewards.Contract.ZapContract(&_V3StakingRewards.CallOpts)
}

// ZapContract is a free data retrieval call binding the contract method 0x2549dad9.
//
// Solidity: function zapContract() view returns(address)
func (_V3StakingRewards *V3StakingRewardsCallerSession) ZapContract() (common.Address, error) {
	return _V3StakingRewards.Contract.ZapContract(&_V3StakingRewards.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_V3StakingRewards *V3StakingRewardsTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V3StakingRewards.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_V3StakingRewards *V3StakingRewardsSession) AcceptOwnership() (*types.Transaction, error) {
	return _V3StakingRewards.Contract.AcceptOwnership(&_V3StakingRewards.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_V3StakingRewards *V3StakingRewardsTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _V3StakingRewards.Contract.AcceptOwnership(&_V3StakingRewards.TransactOpts)
}

// AddReward is a paid mutator transaction binding the contract method 0xd0ed26ae.
//
// Solidity: function addReward(address _rewardsToken, address _rewardsDistributor, uint256 _rewardsDuration) returns()
func (_V3StakingRewards *V3StakingRewardsTransactor) AddReward(opts *bind.TransactOpts, _rewardsToken common.Address, _rewardsDistributor common.Address, _rewardsDuration *big.Int) (*types.Transaction, error) {
	return _V3StakingRewards.contract.Transact(opts, "addReward", _rewardsToken, _rewardsDistributor, _rewardsDuration)
}

// AddReward is a paid mutator transaction binding the contract method 0xd0ed26ae.
//
// Solidity: function addReward(address _rewardsToken, address _rewardsDistributor, uint256 _rewardsDuration) returns()
func (_V3StakingRewards *V3StakingRewardsSession) AddReward(_rewardsToken common.Address, _rewardsDistributor common.Address, _rewardsDuration *big.Int) (*types.Transaction, error) {
	return _V3StakingRewards.Contract.AddReward(&_V3StakingRewards.TransactOpts, _rewardsToken, _rewardsDistributor, _rewardsDuration)
}

// AddReward is a paid mutator transaction binding the contract method 0xd0ed26ae.
//
// Solidity: function addReward(address _rewardsToken, address _rewardsDistributor, uint256 _rewardsDuration) returns()
func (_V3StakingRewards *V3StakingRewardsTransactorSession) AddReward(_rewardsToken common.Address, _rewardsDistributor common.Address, _rewardsDuration *big.Int) (*types.Transaction, error) {
	return _V3StakingRewards.Contract.AddReward(&_V3StakingRewards.TransactOpts, _rewardsToken, _rewardsDistributor, _rewardsDuration)
}

// CloneStakingPool is a paid mutator transaction binding the contract method 0x754a0fdc.
//
// Solidity: function cloneStakingPool(address _owner, address _stakingToken, address _zapContract) returns(address newStakingPool)
func (_V3StakingRewards *V3StakingRewardsTransactor) CloneStakingPool(opts *bind.TransactOpts, _owner common.Address, _stakingToken common.Address, _zapContract common.Address) (*types.Transaction, error) {
	return _V3StakingRewards.contract.Transact(opts, "cloneStakingPool", _owner, _stakingToken, _zapContract)
}

// CloneStakingPool is a paid mutator transaction binding the contract method 0x754a0fdc.
//
// Solidity: function cloneStakingPool(address _owner, address _stakingToken, address _zapContract) returns(address newStakingPool)
func (_V3StakingRewards *V3StakingRewardsSession) CloneStakingPool(_owner common.Address, _stakingToken common.Address, _zapContract common.Address) (*types.Transaction, error) {
	return _V3StakingRewards.Contract.CloneStakingPool(&_V3StakingRewards.TransactOpts, _owner, _stakingToken, _zapContract)
}

// CloneStakingPool is a paid mutator transaction binding the contract method 0x754a0fdc.
//
// Solidity: function cloneStakingPool(address _owner, address _stakingToken, address _zapContract) returns(address newStakingPool)
func (_V3StakingRewards *V3StakingRewardsTransactorSession) CloneStakingPool(_owner common.Address, _stakingToken common.Address, _zapContract common.Address) (*types.Transaction, error) {
	return _V3StakingRewards.Contract.CloneStakingPool(&_V3StakingRewards.TransactOpts, _owner, _stakingToken, _zapContract)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_V3StakingRewards *V3StakingRewardsTransactor) Exit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V3StakingRewards.contract.Transact(opts, "exit")
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_V3StakingRewards *V3StakingRewardsSession) Exit() (*types.Transaction, error) {
	return _V3StakingRewards.Contract.Exit(&_V3StakingRewards.TransactOpts)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_V3StakingRewards *V3StakingRewardsTransactorSession) Exit() (*types.Transaction, error) {
	return _V3StakingRewards.Contract.Exit(&_V3StakingRewards.TransactOpts)
}

// GetOneReward is a paid mutator transaction binding the contract method 0x4ef52305.
//
// Solidity: function getOneReward(address _rewardsToken) returns()
func (_V3StakingRewards *V3StakingRewardsTransactor) GetOneReward(opts *bind.TransactOpts, _rewardsToken common.Address) (*types.Transaction, error) {
	return _V3StakingRewards.contract.Transact(opts, "getOneReward", _rewardsToken)
}

// GetOneReward is a paid mutator transaction binding the contract method 0x4ef52305.
//
// Solidity: function getOneReward(address _rewardsToken) returns()
func (_V3StakingRewards *V3StakingRewardsSession) GetOneReward(_rewardsToken common.Address) (*types.Transaction, error) {
	return _V3StakingRewards.Contract.GetOneReward(&_V3StakingRewards.TransactOpts, _rewardsToken)
}

// GetOneReward is a paid mutator transaction binding the contract method 0x4ef52305.
//
// Solidity: function getOneReward(address _rewardsToken) returns()
func (_V3StakingRewards *V3StakingRewardsTransactorSession) GetOneReward(_rewardsToken common.Address) (*types.Transaction, error) {
	return _V3StakingRewards.Contract.GetOneReward(&_V3StakingRewards.TransactOpts, _rewardsToken)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_V3StakingRewards *V3StakingRewardsTransactor) GetReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V3StakingRewards.contract.Transact(opts, "getReward")
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_V3StakingRewards *V3StakingRewardsSession) GetReward() (*types.Transaction, error) {
	return _V3StakingRewards.Contract.GetReward(&_V3StakingRewards.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_V3StakingRewards *V3StakingRewardsTransactorSession) GetReward() (*types.Transaction, error) {
	return _V3StakingRewards.Contract.GetReward(&_V3StakingRewards.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _owner, address _stakingToken, address _zapContract) returns()
func (_V3StakingRewards *V3StakingRewardsTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _stakingToken common.Address, _zapContract common.Address) (*types.Transaction, error) {
	return _V3StakingRewards.contract.Transact(opts, "initialize", _owner, _stakingToken, _zapContract)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _owner, address _stakingToken, address _zapContract) returns()
func (_V3StakingRewards *V3StakingRewardsSession) Initialize(_owner common.Address, _stakingToken common.Address, _zapContract common.Address) (*types.Transaction, error) {
	return _V3StakingRewards.Contract.Initialize(&_V3StakingRewards.TransactOpts, _owner, _stakingToken, _zapContract)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _owner, address _stakingToken, address _zapContract) returns()
func (_V3StakingRewards *V3StakingRewardsTransactorSession) Initialize(_owner common.Address, _stakingToken common.Address, _zapContract common.Address) (*types.Transaction, error) {
	return _V3StakingRewards.Contract.Initialize(&_V3StakingRewards.TransactOpts, _owner, _stakingToken, _zapContract)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0xb66503cf.
//
// Solidity: function notifyRewardAmount(address _rewardsToken, uint256 _rewardAmount) returns()
func (_V3StakingRewards *V3StakingRewardsTransactor) NotifyRewardAmount(opts *bind.TransactOpts, _rewardsToken common.Address, _rewardAmount *big.Int) (*types.Transaction, error) {
	return _V3StakingRewards.contract.Transact(opts, "notifyRewardAmount", _rewardsToken, _rewardAmount)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0xb66503cf.
//
// Solidity: function notifyRewardAmount(address _rewardsToken, uint256 _rewardAmount) returns()
func (_V3StakingRewards *V3StakingRewardsSession) NotifyRewardAmount(_rewardsToken common.Address, _rewardAmount *big.Int) (*types.Transaction, error) {
	return _V3StakingRewards.Contract.NotifyRewardAmount(&_V3StakingRewards.TransactOpts, _rewardsToken, _rewardAmount)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0xb66503cf.
//
// Solidity: function notifyRewardAmount(address _rewardsToken, uint256 _rewardAmount) returns()
func (_V3StakingRewards *V3StakingRewardsTransactorSession) NotifyRewardAmount(_rewardsToken common.Address, _rewardAmount *big.Int) (*types.Transaction, error) {
	return _V3StakingRewards.Contract.NotifyRewardAmount(&_V3StakingRewards.TransactOpts, _rewardsToken, _rewardAmount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address _tokenAddress, uint256 _tokenAmount) returns()
func (_V3StakingRewards *V3StakingRewardsTransactor) RecoverERC20(opts *bind.TransactOpts, _tokenAddress common.Address, _tokenAmount *big.Int) (*types.Transaction, error) {
	return _V3StakingRewards.contract.Transact(opts, "recoverERC20", _tokenAddress, _tokenAmount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address _tokenAddress, uint256 _tokenAmount) returns()
func (_V3StakingRewards *V3StakingRewardsSession) RecoverERC20(_tokenAddress common.Address, _tokenAmount *big.Int) (*types.Transaction, error) {
	return _V3StakingRewards.Contract.RecoverERC20(&_V3StakingRewards.TransactOpts, _tokenAddress, _tokenAmount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address _tokenAddress, uint256 _tokenAmount) returns()
func (_V3StakingRewards *V3StakingRewardsTransactorSession) RecoverERC20(_tokenAddress common.Address, _tokenAmount *big.Int) (*types.Transaction, error) {
	return _V3StakingRewards.Contract.RecoverERC20(&_V3StakingRewards.TransactOpts, _tokenAddress, _tokenAmount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_V3StakingRewards *V3StakingRewardsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V3StakingRewards.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_V3StakingRewards *V3StakingRewardsSession) RenounceOwnership() (*types.Transaction, error) {
	return _V3StakingRewards.Contract.RenounceOwnership(&_V3StakingRewards.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_V3StakingRewards *V3StakingRewardsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _V3StakingRewards.Contract.RenounceOwnership(&_V3StakingRewards.TransactOpts)
}

// SetRewardsDistributor is a paid mutator transaction binding the contract method 0x3f695b45.
//
// Solidity: function setRewardsDistributor(address _rewardsToken, address _rewardsDistributor) returns()
func (_V3StakingRewards *V3StakingRewardsTransactor) SetRewardsDistributor(opts *bind.TransactOpts, _rewardsToken common.Address, _rewardsDistributor common.Address) (*types.Transaction, error) {
	return _V3StakingRewards.contract.Transact(opts, "setRewardsDistributor", _rewardsToken, _rewardsDistributor)
}

// SetRewardsDistributor is a paid mutator transaction binding the contract method 0x3f695b45.
//
// Solidity: function setRewardsDistributor(address _rewardsToken, address _rewardsDistributor) returns()
func (_V3StakingRewards *V3StakingRewardsSession) SetRewardsDistributor(_rewardsToken common.Address, _rewardsDistributor common.Address) (*types.Transaction, error) {
	return _V3StakingRewards.Contract.SetRewardsDistributor(&_V3StakingRewards.TransactOpts, _rewardsToken, _rewardsDistributor)
}

// SetRewardsDistributor is a paid mutator transaction binding the contract method 0x3f695b45.
//
// Solidity: function setRewardsDistributor(address _rewardsToken, address _rewardsDistributor) returns()
func (_V3StakingRewards *V3StakingRewardsTransactorSession) SetRewardsDistributor(_rewardsToken common.Address, _rewardsDistributor common.Address) (*types.Transaction, error) {
	return _V3StakingRewards.Contract.SetRewardsDistributor(&_V3StakingRewards.TransactOpts, _rewardsToken, _rewardsDistributor)
}

// SetRewardsDuration is a paid mutator transaction binding the contract method 0x2378bea6.
//
// Solidity: function setRewardsDuration(address _rewardsToken, uint256 _rewardsDuration) returns()
func (_V3StakingRewards *V3StakingRewardsTransactor) SetRewardsDuration(opts *bind.TransactOpts, _rewardsToken common.Address, _rewardsDuration *big.Int) (*types.Transaction, error) {
	return _V3StakingRewards.contract.Transact(opts, "setRewardsDuration", _rewardsToken, _rewardsDuration)
}

// SetRewardsDuration is a paid mutator transaction binding the contract method 0x2378bea6.
//
// Solidity: function setRewardsDuration(address _rewardsToken, uint256 _rewardsDuration) returns()
func (_V3StakingRewards *V3StakingRewardsSession) SetRewardsDuration(_rewardsToken common.Address, _rewardsDuration *big.Int) (*types.Transaction, error) {
	return _V3StakingRewards.Contract.SetRewardsDuration(&_V3StakingRewards.TransactOpts, _rewardsToken, _rewardsDuration)
}

// SetRewardsDuration is a paid mutator transaction binding the contract method 0x2378bea6.
//
// Solidity: function setRewardsDuration(address _rewardsToken, uint256 _rewardsDuration) returns()
func (_V3StakingRewards *V3StakingRewardsTransactorSession) SetRewardsDuration(_rewardsToken common.Address, _rewardsDuration *big.Int) (*types.Transaction, error) {
	return _V3StakingRewards.Contract.SetRewardsDuration(&_V3StakingRewards.TransactOpts, _rewardsToken, _rewardsDuration)
}

// SetZapContract is a paid mutator transaction binding the contract method 0xfcaa8737.
//
// Solidity: function setZapContract(address _zapContract) returns()
func (_V3StakingRewards *V3StakingRewardsTransactor) SetZapContract(opts *bind.TransactOpts, _zapContract common.Address) (*types.Transaction, error) {
	return _V3StakingRewards.contract.Transact(opts, "setZapContract", _zapContract)
}

// SetZapContract is a paid mutator transaction binding the contract method 0xfcaa8737.
//
// Solidity: function setZapContract(address _zapContract) returns()
func (_V3StakingRewards *V3StakingRewardsSession) SetZapContract(_zapContract common.Address) (*types.Transaction, error) {
	return _V3StakingRewards.Contract.SetZapContract(&_V3StakingRewards.TransactOpts, _zapContract)
}

// SetZapContract is a paid mutator transaction binding the contract method 0xfcaa8737.
//
// Solidity: function setZapContract(address _zapContract) returns()
func (_V3StakingRewards *V3StakingRewardsTransactorSession) SetZapContract(_zapContract common.Address) (*types.Transaction, error) {
	return _V3StakingRewards.Contract.SetZapContract(&_V3StakingRewards.TransactOpts, _zapContract)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns()
func (_V3StakingRewards *V3StakingRewardsTransactor) Stake(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _V3StakingRewards.contract.Transact(opts, "stake", _amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns()
func (_V3StakingRewards *V3StakingRewardsSession) Stake(_amount *big.Int) (*types.Transaction, error) {
	return _V3StakingRewards.Contract.Stake(&_V3StakingRewards.TransactOpts, _amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns()
func (_V3StakingRewards *V3StakingRewardsTransactorSession) Stake(_amount *big.Int) (*types.Transaction, error) {
	return _V3StakingRewards.Contract.Stake(&_V3StakingRewards.TransactOpts, _amount)
}

// StakeFor is a paid mutator transaction binding the contract method 0x2ee40908.
//
// Solidity: function stakeFor(address _recipient, uint256 _amount) returns()
func (_V3StakingRewards *V3StakingRewardsTransactor) StakeFor(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _V3StakingRewards.contract.Transact(opts, "stakeFor", _recipient, _amount)
}

// StakeFor is a paid mutator transaction binding the contract method 0x2ee40908.
//
// Solidity: function stakeFor(address _recipient, uint256 _amount) returns()
func (_V3StakingRewards *V3StakingRewardsSession) StakeFor(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _V3StakingRewards.Contract.StakeFor(&_V3StakingRewards.TransactOpts, _recipient, _amount)
}

// StakeFor is a paid mutator transaction binding the contract method 0x2ee40908.
//
// Solidity: function stakeFor(address _recipient, uint256 _amount) returns()
func (_V3StakingRewards *V3StakingRewardsTransactorSession) StakeFor(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _V3StakingRewards.Contract.StakeFor(&_V3StakingRewards.TransactOpts, _recipient, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_V3StakingRewards *V3StakingRewardsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _V3StakingRewards.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_V3StakingRewards *V3StakingRewardsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _V3StakingRewards.Contract.TransferOwnership(&_V3StakingRewards.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_V3StakingRewards *V3StakingRewardsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _V3StakingRewards.Contract.TransferOwnership(&_V3StakingRewards.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_V3StakingRewards *V3StakingRewardsTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _V3StakingRewards.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_V3StakingRewards *V3StakingRewardsSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _V3StakingRewards.Contract.Withdraw(&_V3StakingRewards.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_V3StakingRewards *V3StakingRewardsTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _V3StakingRewards.Contract.Withdraw(&_V3StakingRewards.TransactOpts, _amount)
}

// WithdrawFor is a paid mutator transaction binding the contract method 0x0e19c699.
//
// Solidity: function withdrawFor(address _recipient, uint256 _amount, bool _exit) returns()
func (_V3StakingRewards *V3StakingRewardsTransactor) WithdrawFor(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int, _exit bool) (*types.Transaction, error) {
	return _V3StakingRewards.contract.Transact(opts, "withdrawFor", _recipient, _amount, _exit)
}

// WithdrawFor is a paid mutator transaction binding the contract method 0x0e19c699.
//
// Solidity: function withdrawFor(address _recipient, uint256 _amount, bool _exit) returns()
func (_V3StakingRewards *V3StakingRewardsSession) WithdrawFor(_recipient common.Address, _amount *big.Int, _exit bool) (*types.Transaction, error) {
	return _V3StakingRewards.Contract.WithdrawFor(&_V3StakingRewards.TransactOpts, _recipient, _amount, _exit)
}

// WithdrawFor is a paid mutator transaction binding the contract method 0x0e19c699.
//
// Solidity: function withdrawFor(address _recipient, uint256 _amount, bool _exit) returns()
func (_V3StakingRewards *V3StakingRewardsTransactorSession) WithdrawFor(_recipient common.Address, _amount *big.Int, _exit bool) (*types.Transaction, error) {
	return _V3StakingRewards.Contract.WithdrawFor(&_V3StakingRewards.TransactOpts, _recipient, _amount, _exit)
}

// V3StakingRewardsClonedIterator is returned from FilterCloned and is used to iterate over the raw logs and unpacked data for Cloned events raised by the V3StakingRewards contract.
type V3StakingRewardsClonedIterator struct {
	Event *V3StakingRewardsCloned // Event containing the contract specifics and raw log

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
func (it *V3StakingRewardsClonedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V3StakingRewardsCloned)
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
		it.Event = new(V3StakingRewardsCloned)
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
func (it *V3StakingRewardsClonedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V3StakingRewardsClonedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V3StakingRewardsCloned represents a Cloned event raised by the V3StakingRewards contract.
type V3StakingRewardsCloned struct {
	Clone common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCloned is a free log retrieval operation binding the contract event 0x783540fb4221a3238720dc7038937d0d79982bcf895274aa6ad179f82cf0d53c.
//
// Solidity: event Cloned(address indexed clone)
func (_V3StakingRewards *V3StakingRewardsFilterer) FilterCloned(opts *bind.FilterOpts, clone []common.Address) (*V3StakingRewardsClonedIterator, error) {

	var cloneRule []interface{}
	for _, cloneItem := range clone {
		cloneRule = append(cloneRule, cloneItem)
	}

	logs, sub, err := _V3StakingRewards.contract.FilterLogs(opts, "Cloned", cloneRule)
	if err != nil {
		return nil, err
	}
	return &V3StakingRewardsClonedIterator{contract: _V3StakingRewards.contract, event: "Cloned", logs: logs, sub: sub}, nil
}

// WatchCloned is a free log subscription operation binding the contract event 0x783540fb4221a3238720dc7038937d0d79982bcf895274aa6ad179f82cf0d53c.
//
// Solidity: event Cloned(address indexed clone)
func (_V3StakingRewards *V3StakingRewardsFilterer) WatchCloned(opts *bind.WatchOpts, sink chan<- *V3StakingRewardsCloned, clone []common.Address) (event.Subscription, error) {

	var cloneRule []interface{}
	for _, cloneItem := range clone {
		cloneRule = append(cloneRule, cloneItem)
	}

	logs, sub, err := _V3StakingRewards.contract.WatchLogs(opts, "Cloned", cloneRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V3StakingRewardsCloned)
				if err := _V3StakingRewards.contract.UnpackLog(event, "Cloned", log); err != nil {
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
func (_V3StakingRewards *V3StakingRewardsFilterer) ParseCloned(log types.Log) (*V3StakingRewardsCloned, error) {
	event := new(V3StakingRewardsCloned)
	if err := _V3StakingRewards.contract.UnpackLog(event, "Cloned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V3StakingRewardsOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the V3StakingRewards contract.
type V3StakingRewardsOwnershipTransferStartedIterator struct {
	Event *V3StakingRewardsOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *V3StakingRewardsOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V3StakingRewardsOwnershipTransferStarted)
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
		it.Event = new(V3StakingRewardsOwnershipTransferStarted)
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
func (it *V3StakingRewardsOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V3StakingRewardsOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V3StakingRewardsOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the V3StakingRewards contract.
type V3StakingRewardsOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_V3StakingRewards *V3StakingRewardsFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*V3StakingRewardsOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _V3StakingRewards.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &V3StakingRewardsOwnershipTransferStartedIterator{contract: _V3StakingRewards.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_V3StakingRewards *V3StakingRewardsFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *V3StakingRewardsOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _V3StakingRewards.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V3StakingRewardsOwnershipTransferStarted)
				if err := _V3StakingRewards.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
func (_V3StakingRewards *V3StakingRewardsFilterer) ParseOwnershipTransferStarted(log types.Log) (*V3StakingRewardsOwnershipTransferStarted, error) {
	event := new(V3StakingRewardsOwnershipTransferStarted)
	if err := _V3StakingRewards.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V3StakingRewardsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the V3StakingRewards contract.
type V3StakingRewardsOwnershipTransferredIterator struct {
	Event *V3StakingRewardsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *V3StakingRewardsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V3StakingRewardsOwnershipTransferred)
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
		it.Event = new(V3StakingRewardsOwnershipTransferred)
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
func (it *V3StakingRewardsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V3StakingRewardsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V3StakingRewardsOwnershipTransferred represents a OwnershipTransferred event raised by the V3StakingRewards contract.
type V3StakingRewardsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_V3StakingRewards *V3StakingRewardsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*V3StakingRewardsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _V3StakingRewards.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &V3StakingRewardsOwnershipTransferredIterator{contract: _V3StakingRewards.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_V3StakingRewards *V3StakingRewardsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *V3StakingRewardsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _V3StakingRewards.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V3StakingRewardsOwnershipTransferred)
				if err := _V3StakingRewards.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_V3StakingRewards *V3StakingRewardsFilterer) ParseOwnershipTransferred(log types.Log) (*V3StakingRewardsOwnershipTransferred, error) {
	event := new(V3StakingRewardsOwnershipTransferred)
	if err := _V3StakingRewards.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V3StakingRewardsPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the V3StakingRewards contract.
type V3StakingRewardsPausedIterator struct {
	Event *V3StakingRewardsPaused // Event containing the contract specifics and raw log

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
func (it *V3StakingRewardsPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V3StakingRewardsPaused)
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
		it.Event = new(V3StakingRewardsPaused)
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
func (it *V3StakingRewardsPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V3StakingRewardsPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V3StakingRewardsPaused represents a Paused event raised by the V3StakingRewards contract.
type V3StakingRewardsPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_V3StakingRewards *V3StakingRewardsFilterer) FilterPaused(opts *bind.FilterOpts) (*V3StakingRewardsPausedIterator, error) {

	logs, sub, err := _V3StakingRewards.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &V3StakingRewardsPausedIterator{contract: _V3StakingRewards.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_V3StakingRewards *V3StakingRewardsFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *V3StakingRewardsPaused) (event.Subscription, error) {

	logs, sub, err := _V3StakingRewards.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V3StakingRewardsPaused)
				if err := _V3StakingRewards.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_V3StakingRewards *V3StakingRewardsFilterer) ParsePaused(log types.Log) (*V3StakingRewardsPaused, error) {
	event := new(V3StakingRewardsPaused)
	if err := _V3StakingRewards.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V3StakingRewardsRecoveredIterator is returned from FilterRecovered and is used to iterate over the raw logs and unpacked data for Recovered events raised by the V3StakingRewards contract.
type V3StakingRewardsRecoveredIterator struct {
	Event *V3StakingRewardsRecovered // Event containing the contract specifics and raw log

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
func (it *V3StakingRewardsRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V3StakingRewardsRecovered)
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
		it.Event = new(V3StakingRewardsRecovered)
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
func (it *V3StakingRewardsRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V3StakingRewardsRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V3StakingRewardsRecovered represents a Recovered event raised by the V3StakingRewards contract.
type V3StakingRewardsRecovered struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRecovered is a free log retrieval operation binding the contract event 0x8c1256b8896378cd5044f80c202f9772b9d77dc85c8a6eb51967210b09bfaa28.
//
// Solidity: event Recovered(address token, uint256 amount)
func (_V3StakingRewards *V3StakingRewardsFilterer) FilterRecovered(opts *bind.FilterOpts) (*V3StakingRewardsRecoveredIterator, error) {

	logs, sub, err := _V3StakingRewards.contract.FilterLogs(opts, "Recovered")
	if err != nil {
		return nil, err
	}
	return &V3StakingRewardsRecoveredIterator{contract: _V3StakingRewards.contract, event: "Recovered", logs: logs, sub: sub}, nil
}

// WatchRecovered is a free log subscription operation binding the contract event 0x8c1256b8896378cd5044f80c202f9772b9d77dc85c8a6eb51967210b09bfaa28.
//
// Solidity: event Recovered(address token, uint256 amount)
func (_V3StakingRewards *V3StakingRewardsFilterer) WatchRecovered(opts *bind.WatchOpts, sink chan<- *V3StakingRewardsRecovered) (event.Subscription, error) {

	logs, sub, err := _V3StakingRewards.contract.WatchLogs(opts, "Recovered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V3StakingRewardsRecovered)
				if err := _V3StakingRewards.contract.UnpackLog(event, "Recovered", log); err != nil {
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
func (_V3StakingRewards *V3StakingRewardsFilterer) ParseRecovered(log types.Log) (*V3StakingRewardsRecovered, error) {
	event := new(V3StakingRewardsRecovered)
	if err := _V3StakingRewards.contract.UnpackLog(event, "Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V3StakingRewardsRewardAddedIterator is returned from FilterRewardAdded and is used to iterate over the raw logs and unpacked data for RewardAdded events raised by the V3StakingRewards contract.
type V3StakingRewardsRewardAddedIterator struct {
	Event *V3StakingRewardsRewardAdded // Event containing the contract specifics and raw log

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
func (it *V3StakingRewardsRewardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V3StakingRewardsRewardAdded)
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
		it.Event = new(V3StakingRewardsRewardAdded)
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
func (it *V3StakingRewardsRewardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V3StakingRewardsRewardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V3StakingRewardsRewardAdded represents a RewardAdded event raised by the V3StakingRewards contract.
type V3StakingRewardsRewardAdded struct {
	RewardToken common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardAdded is a free log retrieval operation binding the contract event 0xac24935fd910bc682b5ccb1a07b718cadf8cf2f6d1404c4f3ddc3662dae40e29.
//
// Solidity: event RewardAdded(address indexed rewardToken, uint256 amount)
func (_V3StakingRewards *V3StakingRewardsFilterer) FilterRewardAdded(opts *bind.FilterOpts, rewardToken []common.Address) (*V3StakingRewardsRewardAddedIterator, error) {

	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}

	logs, sub, err := _V3StakingRewards.contract.FilterLogs(opts, "RewardAdded", rewardTokenRule)
	if err != nil {
		return nil, err
	}
	return &V3StakingRewardsRewardAddedIterator{contract: _V3StakingRewards.contract, event: "RewardAdded", logs: logs, sub: sub}, nil
}

// WatchRewardAdded is a free log subscription operation binding the contract event 0xac24935fd910bc682b5ccb1a07b718cadf8cf2f6d1404c4f3ddc3662dae40e29.
//
// Solidity: event RewardAdded(address indexed rewardToken, uint256 amount)
func (_V3StakingRewards *V3StakingRewardsFilterer) WatchRewardAdded(opts *bind.WatchOpts, sink chan<- *V3StakingRewardsRewardAdded, rewardToken []common.Address) (event.Subscription, error) {

	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}

	logs, sub, err := _V3StakingRewards.contract.WatchLogs(opts, "RewardAdded", rewardTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V3StakingRewardsRewardAdded)
				if err := _V3StakingRewards.contract.UnpackLog(event, "RewardAdded", log); err != nil {
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
func (_V3StakingRewards *V3StakingRewardsFilterer) ParseRewardAdded(log types.Log) (*V3StakingRewardsRewardAdded, error) {
	event := new(V3StakingRewardsRewardAdded)
	if err := _V3StakingRewards.contract.UnpackLog(event, "RewardAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V3StakingRewardsRewardPaidIterator is returned from FilterRewardPaid and is used to iterate over the raw logs and unpacked data for RewardPaid events raised by the V3StakingRewards contract.
type V3StakingRewardsRewardPaidIterator struct {
	Event *V3StakingRewardsRewardPaid // Event containing the contract specifics and raw log

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
func (it *V3StakingRewardsRewardPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V3StakingRewardsRewardPaid)
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
		it.Event = new(V3StakingRewardsRewardPaid)
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
func (it *V3StakingRewardsRewardPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V3StakingRewardsRewardPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V3StakingRewardsRewardPaid represents a RewardPaid event raised by the V3StakingRewards contract.
type V3StakingRewardsRewardPaid struct {
	User        common.Address
	RewardToken common.Address
	Reward      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardPaid is a free log retrieval operation binding the contract event 0x540798df468d7b23d11f156fdb954cb19ad414d150722a7b6d55ba369dea792e.
//
// Solidity: event RewardPaid(address indexed user, address indexed rewardToken, uint256 reward)
func (_V3StakingRewards *V3StakingRewardsFilterer) FilterRewardPaid(opts *bind.FilterOpts, user []common.Address, rewardToken []common.Address) (*V3StakingRewardsRewardPaidIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}

	logs, sub, err := _V3StakingRewards.contract.FilterLogs(opts, "RewardPaid", userRule, rewardTokenRule)
	if err != nil {
		return nil, err
	}
	return &V3StakingRewardsRewardPaidIterator{contract: _V3StakingRewards.contract, event: "RewardPaid", logs: logs, sub: sub}, nil
}

// WatchRewardPaid is a free log subscription operation binding the contract event 0x540798df468d7b23d11f156fdb954cb19ad414d150722a7b6d55ba369dea792e.
//
// Solidity: event RewardPaid(address indexed user, address indexed rewardToken, uint256 reward)
func (_V3StakingRewards *V3StakingRewardsFilterer) WatchRewardPaid(opts *bind.WatchOpts, sink chan<- *V3StakingRewardsRewardPaid, user []common.Address, rewardToken []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}

	logs, sub, err := _V3StakingRewards.contract.WatchLogs(opts, "RewardPaid", userRule, rewardTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V3StakingRewardsRewardPaid)
				if err := _V3StakingRewards.contract.UnpackLog(event, "RewardPaid", log); err != nil {
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
func (_V3StakingRewards *V3StakingRewardsFilterer) ParseRewardPaid(log types.Log) (*V3StakingRewardsRewardPaid, error) {
	event := new(V3StakingRewardsRewardPaid)
	if err := _V3StakingRewards.contract.UnpackLog(event, "RewardPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V3StakingRewardsRewardsDurationUpdatedIterator is returned from FilterRewardsDurationUpdated and is used to iterate over the raw logs and unpacked data for RewardsDurationUpdated events raised by the V3StakingRewards contract.
type V3StakingRewardsRewardsDurationUpdatedIterator struct {
	Event *V3StakingRewardsRewardsDurationUpdated // Event containing the contract specifics and raw log

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
func (it *V3StakingRewardsRewardsDurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V3StakingRewardsRewardsDurationUpdated)
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
		it.Event = new(V3StakingRewardsRewardsDurationUpdated)
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
func (it *V3StakingRewardsRewardsDurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V3StakingRewardsRewardsDurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V3StakingRewardsRewardsDurationUpdated represents a RewardsDurationUpdated event raised by the V3StakingRewards contract.
type V3StakingRewardsRewardsDurationUpdated struct {
	Token       common.Address
	NewDuration *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardsDurationUpdated is a free log retrieval operation binding the contract event 0xad2f86b01ed93b4b3a150d448c61a4f5d8d38075d3c0c64cc0a26fd6e1f49545.
//
// Solidity: event RewardsDurationUpdated(address token, uint256 newDuration)
func (_V3StakingRewards *V3StakingRewardsFilterer) FilterRewardsDurationUpdated(opts *bind.FilterOpts) (*V3StakingRewardsRewardsDurationUpdatedIterator, error) {

	logs, sub, err := _V3StakingRewards.contract.FilterLogs(opts, "RewardsDurationUpdated")
	if err != nil {
		return nil, err
	}
	return &V3StakingRewardsRewardsDurationUpdatedIterator{contract: _V3StakingRewards.contract, event: "RewardsDurationUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardsDurationUpdated is a free log subscription operation binding the contract event 0xad2f86b01ed93b4b3a150d448c61a4f5d8d38075d3c0c64cc0a26fd6e1f49545.
//
// Solidity: event RewardsDurationUpdated(address token, uint256 newDuration)
func (_V3StakingRewards *V3StakingRewardsFilterer) WatchRewardsDurationUpdated(opts *bind.WatchOpts, sink chan<- *V3StakingRewardsRewardsDurationUpdated) (event.Subscription, error) {

	logs, sub, err := _V3StakingRewards.contract.WatchLogs(opts, "RewardsDurationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V3StakingRewardsRewardsDurationUpdated)
				if err := _V3StakingRewards.contract.UnpackLog(event, "RewardsDurationUpdated", log); err != nil {
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
func (_V3StakingRewards *V3StakingRewardsFilterer) ParseRewardsDurationUpdated(log types.Log) (*V3StakingRewardsRewardsDurationUpdated, error) {
	event := new(V3StakingRewardsRewardsDurationUpdated)
	if err := _V3StakingRewards.contract.UnpackLog(event, "RewardsDurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V3StakingRewardsStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the V3StakingRewards contract.
type V3StakingRewardsStakedIterator struct {
	Event *V3StakingRewardsStaked // Event containing the contract specifics and raw log

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
func (it *V3StakingRewardsStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V3StakingRewardsStaked)
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
		it.Event = new(V3StakingRewardsStaked)
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
func (it *V3StakingRewardsStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V3StakingRewardsStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V3StakingRewardsStaked represents a Staked event raised by the V3StakingRewards contract.
type V3StakingRewardsStaked struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_V3StakingRewards *V3StakingRewardsFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address) (*V3StakingRewardsStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _V3StakingRewards.contract.FilterLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return &V3StakingRewardsStakedIterator{contract: _V3StakingRewards.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_V3StakingRewards *V3StakingRewardsFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *V3StakingRewardsStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _V3StakingRewards.contract.WatchLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V3StakingRewardsStaked)
				if err := _V3StakingRewards.contract.UnpackLog(event, "Staked", log); err != nil {
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
func (_V3StakingRewards *V3StakingRewardsFilterer) ParseStaked(log types.Log) (*V3StakingRewardsStaked, error) {
	event := new(V3StakingRewardsStaked)
	if err := _V3StakingRewards.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V3StakingRewardsStakedForIterator is returned from FilterStakedFor and is used to iterate over the raw logs and unpacked data for StakedFor events raised by the V3StakingRewards contract.
type V3StakingRewardsStakedForIterator struct {
	Event *V3StakingRewardsStakedFor // Event containing the contract specifics and raw log

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
func (it *V3StakingRewardsStakedForIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V3StakingRewardsStakedFor)
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
		it.Event = new(V3StakingRewardsStakedFor)
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
func (it *V3StakingRewardsStakedForIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V3StakingRewardsStakedForIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V3StakingRewardsStakedFor represents a StakedFor event raised by the V3StakingRewards contract.
type V3StakingRewardsStakedFor struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakedFor is a free log retrieval operation binding the contract event 0xd185ae938da574e9cd1073962e1972c75ec585ab222b200a88c0abe2bf0cfe67.
//
// Solidity: event StakedFor(address indexed user, uint256 amount)
func (_V3StakingRewards *V3StakingRewardsFilterer) FilterStakedFor(opts *bind.FilterOpts, user []common.Address) (*V3StakingRewardsStakedForIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _V3StakingRewards.contract.FilterLogs(opts, "StakedFor", userRule)
	if err != nil {
		return nil, err
	}
	return &V3StakingRewardsStakedForIterator{contract: _V3StakingRewards.contract, event: "StakedFor", logs: logs, sub: sub}, nil
}

// WatchStakedFor is a free log subscription operation binding the contract event 0xd185ae938da574e9cd1073962e1972c75ec585ab222b200a88c0abe2bf0cfe67.
//
// Solidity: event StakedFor(address indexed user, uint256 amount)
func (_V3StakingRewards *V3StakingRewardsFilterer) WatchStakedFor(opts *bind.WatchOpts, sink chan<- *V3StakingRewardsStakedFor, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _V3StakingRewards.contract.WatchLogs(opts, "StakedFor", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V3StakingRewardsStakedFor)
				if err := _V3StakingRewards.contract.UnpackLog(event, "StakedFor", log); err != nil {
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
func (_V3StakingRewards *V3StakingRewardsFilterer) ParseStakedFor(log types.Log) (*V3StakingRewardsStakedFor, error) {
	event := new(V3StakingRewardsStakedFor)
	if err := _V3StakingRewards.contract.UnpackLog(event, "StakedFor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V3StakingRewardsUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the V3StakingRewards contract.
type V3StakingRewardsUnpausedIterator struct {
	Event *V3StakingRewardsUnpaused // Event containing the contract specifics and raw log

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
func (it *V3StakingRewardsUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V3StakingRewardsUnpaused)
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
		it.Event = new(V3StakingRewardsUnpaused)
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
func (it *V3StakingRewardsUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V3StakingRewardsUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V3StakingRewardsUnpaused represents a Unpaused event raised by the V3StakingRewards contract.
type V3StakingRewardsUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_V3StakingRewards *V3StakingRewardsFilterer) FilterUnpaused(opts *bind.FilterOpts) (*V3StakingRewardsUnpausedIterator, error) {

	logs, sub, err := _V3StakingRewards.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &V3StakingRewardsUnpausedIterator{contract: _V3StakingRewards.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_V3StakingRewards *V3StakingRewardsFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *V3StakingRewardsUnpaused) (event.Subscription, error) {

	logs, sub, err := _V3StakingRewards.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V3StakingRewardsUnpaused)
				if err := _V3StakingRewards.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_V3StakingRewards *V3StakingRewardsFilterer) ParseUnpaused(log types.Log) (*V3StakingRewardsUnpaused, error) {
	event := new(V3StakingRewardsUnpaused)
	if err := _V3StakingRewards.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V3StakingRewardsWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the V3StakingRewards contract.
type V3StakingRewardsWithdrawnIterator struct {
	Event *V3StakingRewardsWithdrawn // Event containing the contract specifics and raw log

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
func (it *V3StakingRewardsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V3StakingRewardsWithdrawn)
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
		it.Event = new(V3StakingRewardsWithdrawn)
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
func (it *V3StakingRewardsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V3StakingRewardsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V3StakingRewardsWithdrawn represents a Withdrawn event raised by the V3StakingRewards contract.
type V3StakingRewardsWithdrawn struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_V3StakingRewards *V3StakingRewardsFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address) (*V3StakingRewardsWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _V3StakingRewards.contract.FilterLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &V3StakingRewardsWithdrawnIterator{contract: _V3StakingRewards.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_V3StakingRewards *V3StakingRewardsFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *V3StakingRewardsWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _V3StakingRewards.contract.WatchLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V3StakingRewardsWithdrawn)
				if err := _V3StakingRewards.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
func (_V3StakingRewards *V3StakingRewardsFilterer) ParseWithdrawn(log types.Log) (*V3StakingRewardsWithdrawn, error) {
	event := new(V3StakingRewardsWithdrawn)
	if err := _V3StakingRewards.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V3StakingRewardsWithdrawnForIterator is returned from FilterWithdrawnFor and is used to iterate over the raw logs and unpacked data for WithdrawnFor events raised by the V3StakingRewards contract.
type V3StakingRewardsWithdrawnForIterator struct {
	Event *V3StakingRewardsWithdrawnFor // Event containing the contract specifics and raw log

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
func (it *V3StakingRewardsWithdrawnForIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V3StakingRewardsWithdrawnFor)
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
		it.Event = new(V3StakingRewardsWithdrawnFor)
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
func (it *V3StakingRewardsWithdrawnForIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V3StakingRewardsWithdrawnForIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V3StakingRewardsWithdrawnFor represents a WithdrawnFor event raised by the V3StakingRewards contract.
type V3StakingRewardsWithdrawnFor struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnFor is a free log retrieval operation binding the contract event 0x04de67fe258683d541c77f7275134560914ac5eb0fe853df87c0946439c2e706.
//
// Solidity: event WithdrawnFor(address indexed user, uint256 amount)
func (_V3StakingRewards *V3StakingRewardsFilterer) FilterWithdrawnFor(opts *bind.FilterOpts, user []common.Address) (*V3StakingRewardsWithdrawnForIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _V3StakingRewards.contract.FilterLogs(opts, "WithdrawnFor", userRule)
	if err != nil {
		return nil, err
	}
	return &V3StakingRewardsWithdrawnForIterator{contract: _V3StakingRewards.contract, event: "WithdrawnFor", logs: logs, sub: sub}, nil
}

// WatchWithdrawnFor is a free log subscription operation binding the contract event 0x04de67fe258683d541c77f7275134560914ac5eb0fe853df87c0946439c2e706.
//
// Solidity: event WithdrawnFor(address indexed user, uint256 amount)
func (_V3StakingRewards *V3StakingRewardsFilterer) WatchWithdrawnFor(opts *bind.WatchOpts, sink chan<- *V3StakingRewardsWithdrawnFor, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _V3StakingRewards.contract.WatchLogs(opts, "WithdrawnFor", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V3StakingRewardsWithdrawnFor)
				if err := _V3StakingRewards.contract.UnpackLog(event, "WithdrawnFor", log); err != nil {
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
func (_V3StakingRewards *V3StakingRewardsFilterer) ParseWithdrawnFor(log types.Log) (*V3StakingRewardsWithdrawnFor, error) {
	event := new(V3StakingRewardsWithdrawnFor)
	if err := _V3StakingRewards.contract.UnpackLog(event, "WithdrawnFor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V3StakingRewardsZapContractUpdatedIterator is returned from FilterZapContractUpdated and is used to iterate over the raw logs and unpacked data for ZapContractUpdated events raised by the V3StakingRewards contract.
type V3StakingRewardsZapContractUpdatedIterator struct {
	Event *V3StakingRewardsZapContractUpdated // Event containing the contract specifics and raw log

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
func (it *V3StakingRewardsZapContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V3StakingRewardsZapContractUpdated)
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
		it.Event = new(V3StakingRewardsZapContractUpdated)
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
func (it *V3StakingRewardsZapContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V3StakingRewardsZapContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V3StakingRewardsZapContractUpdated represents a ZapContractUpdated event raised by the V3StakingRewards contract.
type V3StakingRewardsZapContractUpdated struct {
	ZapContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterZapContractUpdated is a free log retrieval operation binding the contract event 0x0da458581861b4ff3ae5c23a916ec5218bbf4c371497cec29006f5c55d9c9a84.
//
// Solidity: event ZapContractUpdated(address _zapContract)
func (_V3StakingRewards *V3StakingRewardsFilterer) FilterZapContractUpdated(opts *bind.FilterOpts) (*V3StakingRewardsZapContractUpdatedIterator, error) {

	logs, sub, err := _V3StakingRewards.contract.FilterLogs(opts, "ZapContractUpdated")
	if err != nil {
		return nil, err
	}
	return &V3StakingRewardsZapContractUpdatedIterator{contract: _V3StakingRewards.contract, event: "ZapContractUpdated", logs: logs, sub: sub}, nil
}

// WatchZapContractUpdated is a free log subscription operation binding the contract event 0x0da458581861b4ff3ae5c23a916ec5218bbf4c371497cec29006f5c55d9c9a84.
//
// Solidity: event ZapContractUpdated(address _zapContract)
func (_V3StakingRewards *V3StakingRewardsFilterer) WatchZapContractUpdated(opts *bind.WatchOpts, sink chan<- *V3StakingRewardsZapContractUpdated) (event.Subscription, error) {

	logs, sub, err := _V3StakingRewards.contract.WatchLogs(opts, "ZapContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V3StakingRewardsZapContractUpdated)
				if err := _V3StakingRewards.contract.UnpackLog(event, "ZapContractUpdated", log); err != nil {
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
func (_V3StakingRewards *V3StakingRewardsFilterer) ParseZapContractUpdated(log types.Log) (*V3StakingRewardsZapContractUpdated, error) {
	event := new(V3StakingRewardsZapContractUpdated)
	if err := _V3StakingRewards.contract.UnpackLog(event, "ZapContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
