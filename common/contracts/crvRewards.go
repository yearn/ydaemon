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

// CrvRewardsMetaData contains all meta data concerning the CrvRewards contract.
var CrvRewardsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pid_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakingToken_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardToken_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardManager_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reward\",\"type\":\"address\"}],\"name\":\"addExtraReward\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clearExtraRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"donate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"duration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"earned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"extraRewards\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"extraRewardsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReward\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_claimExtras\",\"type\":\"bool\"}],\"name\":\"getReward\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"historicalRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTimeRewardApplicable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newRewardRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodFinish\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewards\",\"type\":\"uint256\"}],\"name\":\"queueNewRewards\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queuedRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerTokenStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_for\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"stakeFor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRewardPerTokenPaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"claim\",\"type\":\"bool\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"claim\",\"type\":\"bool\"}],\"name\":\"withdrawAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"claim\",\"type\":\"bool\"}],\"name\":\"withdrawAllAndUnwrap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"claim\",\"type\":\"bool\"}],\"name\":\"withdrawAndUnwrap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CrvRewardsABI is the input ABI used to generate the binding from.
// Deprecated: Use CrvRewardsMetaData.ABI instead.
var CrvRewardsABI = CrvRewardsMetaData.ABI

// CrvRewards is an auto generated Go binding around an Ethereum contract.
type CrvRewards struct {
	CrvRewardsCaller     // Read-only binding to the contract
	CrvRewardsTransactor // Write-only binding to the contract
	CrvRewardsFilterer   // Log filterer for contract events
}

// CrvRewardsCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrvRewardsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrvRewardsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrvRewardsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrvRewardsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrvRewardsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrvRewardsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrvRewardsSession struct {
	Contract     *CrvRewards       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrvRewardsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrvRewardsCallerSession struct {
	Contract *CrvRewardsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CrvRewardsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrvRewardsTransactorSession struct {
	Contract     *CrvRewardsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CrvRewardsRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrvRewardsRaw struct {
	Contract *CrvRewards // Generic contract binding to access the raw methods on
}

// CrvRewardsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrvRewardsCallerRaw struct {
	Contract *CrvRewardsCaller // Generic read-only contract binding to access the raw methods on
}

// CrvRewardsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrvRewardsTransactorRaw struct {
	Contract *CrvRewardsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrvRewards creates a new instance of CrvRewards, bound to a specific deployed contract.
func NewCrvRewards(address common.Address, backend bind.ContractBackend) (*CrvRewards, error) {
	contract, err := bindCrvRewards(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrvRewards{CrvRewardsCaller: CrvRewardsCaller{contract: contract}, CrvRewardsTransactor: CrvRewardsTransactor{contract: contract}, CrvRewardsFilterer: CrvRewardsFilterer{contract: contract}}, nil
}

// NewCrvRewardsCaller creates a new read-only instance of CrvRewards, bound to a specific deployed contract.
func NewCrvRewardsCaller(address common.Address, caller bind.ContractCaller) (*CrvRewardsCaller, error) {
	contract, err := bindCrvRewards(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrvRewardsCaller{contract: contract}, nil
}

// NewCrvRewardsTransactor creates a new write-only instance of CrvRewards, bound to a specific deployed contract.
func NewCrvRewardsTransactor(address common.Address, transactor bind.ContractTransactor) (*CrvRewardsTransactor, error) {
	contract, err := bindCrvRewards(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrvRewardsTransactor{contract: contract}, nil
}

// NewCrvRewardsFilterer creates a new log filterer instance of CrvRewards, bound to a specific deployed contract.
func NewCrvRewardsFilterer(address common.Address, filterer bind.ContractFilterer) (*CrvRewardsFilterer, error) {
	contract, err := bindCrvRewards(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrvRewardsFilterer{contract: contract}, nil
}

// bindCrvRewards binds a generic wrapper to an already deployed contract.
func bindCrvRewards(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CrvRewardsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrvRewards *CrvRewardsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrvRewards.Contract.CrvRewardsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrvRewards *CrvRewardsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrvRewards.Contract.CrvRewardsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrvRewards *CrvRewardsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrvRewards.Contract.CrvRewardsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrvRewards *CrvRewardsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrvRewards.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrvRewards *CrvRewardsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrvRewards.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrvRewards *CrvRewardsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrvRewards.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CrvRewards *CrvRewardsCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrvRewards.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CrvRewards *CrvRewardsSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _CrvRewards.Contract.BalanceOf(&_CrvRewards.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CrvRewards *CrvRewardsCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _CrvRewards.Contract.BalanceOf(&_CrvRewards.CallOpts, account)
}

// CurrentRewards is a free data retrieval call binding the contract method 0x901a7d53.
//
// Solidity: function currentRewards() view returns(uint256)
func (_CrvRewards *CrvRewardsCaller) CurrentRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvRewards.contract.Call(opts, &out, "currentRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentRewards is a free data retrieval call binding the contract method 0x901a7d53.
//
// Solidity: function currentRewards() view returns(uint256)
func (_CrvRewards *CrvRewardsSession) CurrentRewards() (*big.Int, error) {
	return _CrvRewards.Contract.CurrentRewards(&_CrvRewards.CallOpts)
}

// CurrentRewards is a free data retrieval call binding the contract method 0x901a7d53.
//
// Solidity: function currentRewards() view returns(uint256)
func (_CrvRewards *CrvRewardsCallerSession) CurrentRewards() (*big.Int, error) {
	return _CrvRewards.Contract.CurrentRewards(&_CrvRewards.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_CrvRewards *CrvRewardsCaller) Duration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvRewards.contract.Call(opts, &out, "duration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_CrvRewards *CrvRewardsSession) Duration() (*big.Int, error) {
	return _CrvRewards.Contract.Duration(&_CrvRewards.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_CrvRewards *CrvRewardsCallerSession) Duration() (*big.Int, error) {
	return _CrvRewards.Contract.Duration(&_CrvRewards.CallOpts)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_CrvRewards *CrvRewardsCaller) Earned(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrvRewards.contract.Call(opts, &out, "earned", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_CrvRewards *CrvRewardsSession) Earned(account common.Address) (*big.Int, error) {
	return _CrvRewards.Contract.Earned(&_CrvRewards.CallOpts, account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_CrvRewards *CrvRewardsCallerSession) Earned(account common.Address) (*big.Int, error) {
	return _CrvRewards.Contract.Earned(&_CrvRewards.CallOpts, account)
}

// ExtraRewards is a free data retrieval call binding the contract method 0x40c35446.
//
// Solidity: function extraRewards(uint256 ) view returns(address)
func (_CrvRewards *CrvRewardsCaller) ExtraRewards(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CrvRewards.contract.Call(opts, &out, "extraRewards", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExtraRewards is a free data retrieval call binding the contract method 0x40c35446.
//
// Solidity: function extraRewards(uint256 ) view returns(address)
func (_CrvRewards *CrvRewardsSession) ExtraRewards(arg0 *big.Int) (common.Address, error) {
	return _CrvRewards.Contract.ExtraRewards(&_CrvRewards.CallOpts, arg0)
}

// ExtraRewards is a free data retrieval call binding the contract method 0x40c35446.
//
// Solidity: function extraRewards(uint256 ) view returns(address)
func (_CrvRewards *CrvRewardsCallerSession) ExtraRewards(arg0 *big.Int) (common.Address, error) {
	return _CrvRewards.Contract.ExtraRewards(&_CrvRewards.CallOpts, arg0)
}

// ExtraRewardsLength is a free data retrieval call binding the contract method 0xd55a23f4.
//
// Solidity: function extraRewardsLength() view returns(uint256)
func (_CrvRewards *CrvRewardsCaller) ExtraRewardsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvRewards.contract.Call(opts, &out, "extraRewardsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExtraRewardsLength is a free data retrieval call binding the contract method 0xd55a23f4.
//
// Solidity: function extraRewardsLength() view returns(uint256)
func (_CrvRewards *CrvRewardsSession) ExtraRewardsLength() (*big.Int, error) {
	return _CrvRewards.Contract.ExtraRewardsLength(&_CrvRewards.CallOpts)
}

// ExtraRewardsLength is a free data retrieval call binding the contract method 0xd55a23f4.
//
// Solidity: function extraRewardsLength() view returns(uint256)
func (_CrvRewards *CrvRewardsCallerSession) ExtraRewardsLength() (*big.Int, error) {
	return _CrvRewards.Contract.ExtraRewardsLength(&_CrvRewards.CallOpts)
}

// HistoricalRewards is a free data retrieval call binding the contract method 0x262d3d6d.
//
// Solidity: function historicalRewards() view returns(uint256)
func (_CrvRewards *CrvRewardsCaller) HistoricalRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvRewards.contract.Call(opts, &out, "historicalRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HistoricalRewards is a free data retrieval call binding the contract method 0x262d3d6d.
//
// Solidity: function historicalRewards() view returns(uint256)
func (_CrvRewards *CrvRewardsSession) HistoricalRewards() (*big.Int, error) {
	return _CrvRewards.Contract.HistoricalRewards(&_CrvRewards.CallOpts)
}

// HistoricalRewards is a free data retrieval call binding the contract method 0x262d3d6d.
//
// Solidity: function historicalRewards() view returns(uint256)
func (_CrvRewards *CrvRewardsCallerSession) HistoricalRewards() (*big.Int, error) {
	return _CrvRewards.Contract.HistoricalRewards(&_CrvRewards.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_CrvRewards *CrvRewardsCaller) LastTimeRewardApplicable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvRewards.contract.Call(opts, &out, "lastTimeRewardApplicable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_CrvRewards *CrvRewardsSession) LastTimeRewardApplicable() (*big.Int, error) {
	return _CrvRewards.Contract.LastTimeRewardApplicable(&_CrvRewards.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_CrvRewards *CrvRewardsCallerSession) LastTimeRewardApplicable() (*big.Int, error) {
	return _CrvRewards.Contract.LastTimeRewardApplicable(&_CrvRewards.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_CrvRewards *CrvRewardsCaller) LastUpdateTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvRewards.contract.Call(opts, &out, "lastUpdateTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_CrvRewards *CrvRewardsSession) LastUpdateTime() (*big.Int, error) {
	return _CrvRewards.Contract.LastUpdateTime(&_CrvRewards.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_CrvRewards *CrvRewardsCallerSession) LastUpdateTime() (*big.Int, error) {
	return _CrvRewards.Contract.LastUpdateTime(&_CrvRewards.CallOpts)
}

// NewRewardRatio is a free data retrieval call binding the contract method 0x6c8bcee8.
//
// Solidity: function newRewardRatio() view returns(uint256)
func (_CrvRewards *CrvRewardsCaller) NewRewardRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvRewards.contract.Call(opts, &out, "newRewardRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NewRewardRatio is a free data retrieval call binding the contract method 0x6c8bcee8.
//
// Solidity: function newRewardRatio() view returns(uint256)
func (_CrvRewards *CrvRewardsSession) NewRewardRatio() (*big.Int, error) {
	return _CrvRewards.Contract.NewRewardRatio(&_CrvRewards.CallOpts)
}

// NewRewardRatio is a free data retrieval call binding the contract method 0x6c8bcee8.
//
// Solidity: function newRewardRatio() view returns(uint256)
func (_CrvRewards *CrvRewardsCallerSession) NewRewardRatio() (*big.Int, error) {
	return _CrvRewards.Contract.NewRewardRatio(&_CrvRewards.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_CrvRewards *CrvRewardsCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrvRewards.contract.Call(opts, &out, "operator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_CrvRewards *CrvRewardsSession) Operator() (common.Address, error) {
	return _CrvRewards.Contract.Operator(&_CrvRewards.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_CrvRewards *CrvRewardsCallerSession) Operator() (common.Address, error) {
	return _CrvRewards.Contract.Operator(&_CrvRewards.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_CrvRewards *CrvRewardsCaller) PeriodFinish(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvRewards.contract.Call(opts, &out, "periodFinish")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_CrvRewards *CrvRewardsSession) PeriodFinish() (*big.Int, error) {
	return _CrvRewards.Contract.PeriodFinish(&_CrvRewards.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_CrvRewards *CrvRewardsCallerSession) PeriodFinish() (*big.Int, error) {
	return _CrvRewards.Contract.PeriodFinish(&_CrvRewards.CallOpts)
}

// Pid is a free data retrieval call binding the contract method 0xf1068454.
//
// Solidity: function pid() view returns(uint256)
func (_CrvRewards *CrvRewardsCaller) Pid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvRewards.contract.Call(opts, &out, "pid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Pid is a free data retrieval call binding the contract method 0xf1068454.
//
// Solidity: function pid() view returns(uint256)
func (_CrvRewards *CrvRewardsSession) Pid() (*big.Int, error) {
	return _CrvRewards.Contract.Pid(&_CrvRewards.CallOpts)
}

// Pid is a free data retrieval call binding the contract method 0xf1068454.
//
// Solidity: function pid() view returns(uint256)
func (_CrvRewards *CrvRewardsCallerSession) Pid() (*big.Int, error) {
	return _CrvRewards.Contract.Pid(&_CrvRewards.CallOpts)
}

// QueuedRewards is a free data retrieval call binding the contract method 0x63d38c3b.
//
// Solidity: function queuedRewards() view returns(uint256)
func (_CrvRewards *CrvRewardsCaller) QueuedRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvRewards.contract.Call(opts, &out, "queuedRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueuedRewards is a free data retrieval call binding the contract method 0x63d38c3b.
//
// Solidity: function queuedRewards() view returns(uint256)
func (_CrvRewards *CrvRewardsSession) QueuedRewards() (*big.Int, error) {
	return _CrvRewards.Contract.QueuedRewards(&_CrvRewards.CallOpts)
}

// QueuedRewards is a free data retrieval call binding the contract method 0x63d38c3b.
//
// Solidity: function queuedRewards() view returns(uint256)
func (_CrvRewards *CrvRewardsCallerSession) QueuedRewards() (*big.Int, error) {
	return _CrvRewards.Contract.QueuedRewards(&_CrvRewards.CallOpts)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_CrvRewards *CrvRewardsCaller) RewardManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrvRewards.contract.Call(opts, &out, "rewardManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_CrvRewards *CrvRewardsSession) RewardManager() (common.Address, error) {
	return _CrvRewards.Contract.RewardManager(&_CrvRewards.CallOpts)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_CrvRewards *CrvRewardsCallerSession) RewardManager() (common.Address, error) {
	return _CrvRewards.Contract.RewardManager(&_CrvRewards.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_CrvRewards *CrvRewardsCaller) RewardPerToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvRewards.contract.Call(opts, &out, "rewardPerToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_CrvRewards *CrvRewardsSession) RewardPerToken() (*big.Int, error) {
	return _CrvRewards.Contract.RewardPerToken(&_CrvRewards.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_CrvRewards *CrvRewardsCallerSession) RewardPerToken() (*big.Int, error) {
	return _CrvRewards.Contract.RewardPerToken(&_CrvRewards.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_CrvRewards *CrvRewardsCaller) RewardPerTokenStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvRewards.contract.Call(opts, &out, "rewardPerTokenStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_CrvRewards *CrvRewardsSession) RewardPerTokenStored() (*big.Int, error) {
	return _CrvRewards.Contract.RewardPerTokenStored(&_CrvRewards.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_CrvRewards *CrvRewardsCallerSession) RewardPerTokenStored() (*big.Int, error) {
	return _CrvRewards.Contract.RewardPerTokenStored(&_CrvRewards.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_CrvRewards *CrvRewardsCaller) RewardRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvRewards.contract.Call(opts, &out, "rewardRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_CrvRewards *CrvRewardsSession) RewardRate() (*big.Int, error) {
	return _CrvRewards.Contract.RewardRate(&_CrvRewards.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_CrvRewards *CrvRewardsCallerSession) RewardRate() (*big.Int, error) {
	return _CrvRewards.Contract.RewardRate(&_CrvRewards.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_CrvRewards *CrvRewardsCaller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrvRewards.contract.Call(opts, &out, "rewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_CrvRewards *CrvRewardsSession) RewardToken() (common.Address, error) {
	return _CrvRewards.Contract.RewardToken(&_CrvRewards.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_CrvRewards *CrvRewardsCallerSession) RewardToken() (common.Address, error) {
	return _CrvRewards.Contract.RewardToken(&_CrvRewards.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_CrvRewards *CrvRewardsCaller) Rewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrvRewards.contract.Call(opts, &out, "rewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_CrvRewards *CrvRewardsSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _CrvRewards.Contract.Rewards(&_CrvRewards.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_CrvRewards *CrvRewardsCallerSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _CrvRewards.Contract.Rewards(&_CrvRewards.CallOpts, arg0)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_CrvRewards *CrvRewardsCaller) StakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrvRewards.contract.Call(opts, &out, "stakingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_CrvRewards *CrvRewardsSession) StakingToken() (common.Address, error) {
	return _CrvRewards.Contract.StakingToken(&_CrvRewards.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_CrvRewards *CrvRewardsCallerSession) StakingToken() (common.Address, error) {
	return _CrvRewards.Contract.StakingToken(&_CrvRewards.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CrvRewards *CrvRewardsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvRewards.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CrvRewards *CrvRewardsSession) TotalSupply() (*big.Int, error) {
	return _CrvRewards.Contract.TotalSupply(&_CrvRewards.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CrvRewards *CrvRewardsCallerSession) TotalSupply() (*big.Int, error) {
	return _CrvRewards.Contract.TotalSupply(&_CrvRewards.CallOpts)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_CrvRewards *CrvRewardsCaller) UserRewardPerTokenPaid(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrvRewards.contract.Call(opts, &out, "userRewardPerTokenPaid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_CrvRewards *CrvRewardsSession) UserRewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _CrvRewards.Contract.UserRewardPerTokenPaid(&_CrvRewards.CallOpts, arg0)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_CrvRewards *CrvRewardsCallerSession) UserRewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _CrvRewards.Contract.UserRewardPerTokenPaid(&_CrvRewards.CallOpts, arg0)
}

// AddExtraReward is a paid mutator transaction binding the contract method 0x5e43c47b.
//
// Solidity: function addExtraReward(address _reward) returns(bool)
func (_CrvRewards *CrvRewardsTransactor) AddExtraReward(opts *bind.TransactOpts, _reward common.Address) (*types.Transaction, error) {
	return _CrvRewards.contract.Transact(opts, "addExtraReward", _reward)
}

// AddExtraReward is a paid mutator transaction binding the contract method 0x5e43c47b.
//
// Solidity: function addExtraReward(address _reward) returns(bool)
func (_CrvRewards *CrvRewardsSession) AddExtraReward(_reward common.Address) (*types.Transaction, error) {
	return _CrvRewards.Contract.AddExtraReward(&_CrvRewards.TransactOpts, _reward)
}

// AddExtraReward is a paid mutator transaction binding the contract method 0x5e43c47b.
//
// Solidity: function addExtraReward(address _reward) returns(bool)
func (_CrvRewards *CrvRewardsTransactorSession) AddExtraReward(_reward common.Address) (*types.Transaction, error) {
	return _CrvRewards.Contract.AddExtraReward(&_CrvRewards.TransactOpts, _reward)
}

// ClearExtraRewards is a paid mutator transaction binding the contract method 0x0569d388.
//
// Solidity: function clearExtraRewards() returns()
func (_CrvRewards *CrvRewardsTransactor) ClearExtraRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrvRewards.contract.Transact(opts, "clearExtraRewards")
}

// ClearExtraRewards is a paid mutator transaction binding the contract method 0x0569d388.
//
// Solidity: function clearExtraRewards() returns()
func (_CrvRewards *CrvRewardsSession) ClearExtraRewards() (*types.Transaction, error) {
	return _CrvRewards.Contract.ClearExtraRewards(&_CrvRewards.TransactOpts)
}

// ClearExtraRewards is a paid mutator transaction binding the contract method 0x0569d388.
//
// Solidity: function clearExtraRewards() returns()
func (_CrvRewards *CrvRewardsTransactorSession) ClearExtraRewards() (*types.Transaction, error) {
	return _CrvRewards.Contract.ClearExtraRewards(&_CrvRewards.TransactOpts)
}

// Donate is a paid mutator transaction binding the contract method 0xf14faf6f.
//
// Solidity: function donate(uint256 _amount) returns(bool)
func (_CrvRewards *CrvRewardsTransactor) Donate(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _CrvRewards.contract.Transact(opts, "donate", _amount)
}

// Donate is a paid mutator transaction binding the contract method 0xf14faf6f.
//
// Solidity: function donate(uint256 _amount) returns(bool)
func (_CrvRewards *CrvRewardsSession) Donate(_amount *big.Int) (*types.Transaction, error) {
	return _CrvRewards.Contract.Donate(&_CrvRewards.TransactOpts, _amount)
}

// Donate is a paid mutator transaction binding the contract method 0xf14faf6f.
//
// Solidity: function donate(uint256 _amount) returns(bool)
func (_CrvRewards *CrvRewardsTransactorSession) Donate(_amount *big.Int) (*types.Transaction, error) {
	return _CrvRewards.Contract.Donate(&_CrvRewards.TransactOpts, _amount)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns(bool)
func (_CrvRewards *CrvRewardsTransactor) GetReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrvRewards.contract.Transact(opts, "getReward")
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns(bool)
func (_CrvRewards *CrvRewardsSession) GetReward() (*types.Transaction, error) {
	return _CrvRewards.Contract.GetReward(&_CrvRewards.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns(bool)
func (_CrvRewards *CrvRewardsTransactorSession) GetReward() (*types.Transaction, error) {
	return _CrvRewards.Contract.GetReward(&_CrvRewards.TransactOpts)
}

// GetReward0 is a paid mutator transaction binding the contract method 0x7050ccd9.
//
// Solidity: function getReward(address _account, bool _claimExtras) returns(bool)
func (_CrvRewards *CrvRewardsTransactor) GetReward0(opts *bind.TransactOpts, _account common.Address, _claimExtras bool) (*types.Transaction, error) {
	return _CrvRewards.contract.Transact(opts, "getReward0", _account, _claimExtras)
}

// GetReward0 is a paid mutator transaction binding the contract method 0x7050ccd9.
//
// Solidity: function getReward(address _account, bool _claimExtras) returns(bool)
func (_CrvRewards *CrvRewardsSession) GetReward0(_account common.Address, _claimExtras bool) (*types.Transaction, error) {
	return _CrvRewards.Contract.GetReward0(&_CrvRewards.TransactOpts, _account, _claimExtras)
}

// GetReward0 is a paid mutator transaction binding the contract method 0x7050ccd9.
//
// Solidity: function getReward(address _account, bool _claimExtras) returns(bool)
func (_CrvRewards *CrvRewardsTransactorSession) GetReward0(_account common.Address, _claimExtras bool) (*types.Transaction, error) {
	return _CrvRewards.Contract.GetReward0(&_CrvRewards.TransactOpts, _account, _claimExtras)
}

// QueueNewRewards is a paid mutator transaction binding the contract method 0x590a41f5.
//
// Solidity: function queueNewRewards(uint256 _rewards) returns(bool)
func (_CrvRewards *CrvRewardsTransactor) QueueNewRewards(opts *bind.TransactOpts, _rewards *big.Int) (*types.Transaction, error) {
	return _CrvRewards.contract.Transact(opts, "queueNewRewards", _rewards)
}

// QueueNewRewards is a paid mutator transaction binding the contract method 0x590a41f5.
//
// Solidity: function queueNewRewards(uint256 _rewards) returns(bool)
func (_CrvRewards *CrvRewardsSession) QueueNewRewards(_rewards *big.Int) (*types.Transaction, error) {
	return _CrvRewards.Contract.QueueNewRewards(&_CrvRewards.TransactOpts, _rewards)
}

// QueueNewRewards is a paid mutator transaction binding the contract method 0x590a41f5.
//
// Solidity: function queueNewRewards(uint256 _rewards) returns(bool)
func (_CrvRewards *CrvRewardsTransactorSession) QueueNewRewards(_rewards *big.Int) (*types.Transaction, error) {
	return _CrvRewards.Contract.QueueNewRewards(&_CrvRewards.TransactOpts, _rewards)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns(bool)
func (_CrvRewards *CrvRewardsTransactor) Stake(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _CrvRewards.contract.Transact(opts, "stake", _amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns(bool)
func (_CrvRewards *CrvRewardsSession) Stake(_amount *big.Int) (*types.Transaction, error) {
	return _CrvRewards.Contract.Stake(&_CrvRewards.TransactOpts, _amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns(bool)
func (_CrvRewards *CrvRewardsTransactorSession) Stake(_amount *big.Int) (*types.Transaction, error) {
	return _CrvRewards.Contract.Stake(&_CrvRewards.TransactOpts, _amount)
}

// StakeAll is a paid mutator transaction binding the contract method 0x8dcb4061.
//
// Solidity: function stakeAll() returns(bool)
func (_CrvRewards *CrvRewardsTransactor) StakeAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrvRewards.contract.Transact(opts, "stakeAll")
}

// StakeAll is a paid mutator transaction binding the contract method 0x8dcb4061.
//
// Solidity: function stakeAll() returns(bool)
func (_CrvRewards *CrvRewardsSession) StakeAll() (*types.Transaction, error) {
	return _CrvRewards.Contract.StakeAll(&_CrvRewards.TransactOpts)
}

// StakeAll is a paid mutator transaction binding the contract method 0x8dcb4061.
//
// Solidity: function stakeAll() returns(bool)
func (_CrvRewards *CrvRewardsTransactorSession) StakeAll() (*types.Transaction, error) {
	return _CrvRewards.Contract.StakeAll(&_CrvRewards.TransactOpts)
}

// StakeFor is a paid mutator transaction binding the contract method 0x2ee40908.
//
// Solidity: function stakeFor(address _for, uint256 _amount) returns(bool)
func (_CrvRewards *CrvRewardsTransactor) StakeFor(opts *bind.TransactOpts, _for common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _CrvRewards.contract.Transact(opts, "stakeFor", _for, _amount)
}

// StakeFor is a paid mutator transaction binding the contract method 0x2ee40908.
//
// Solidity: function stakeFor(address _for, uint256 _amount) returns(bool)
func (_CrvRewards *CrvRewardsSession) StakeFor(_for common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _CrvRewards.Contract.StakeFor(&_CrvRewards.TransactOpts, _for, _amount)
}

// StakeFor is a paid mutator transaction binding the contract method 0x2ee40908.
//
// Solidity: function stakeFor(address _for, uint256 _amount) returns(bool)
func (_CrvRewards *CrvRewardsTransactorSession) StakeFor(_for common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _CrvRewards.Contract.StakeFor(&_CrvRewards.TransactOpts, _for, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x38d07436.
//
// Solidity: function withdraw(uint256 amount, bool claim) returns(bool)
func (_CrvRewards *CrvRewardsTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int, claim bool) (*types.Transaction, error) {
	return _CrvRewards.contract.Transact(opts, "withdraw", amount, claim)
}

// Withdraw is a paid mutator transaction binding the contract method 0x38d07436.
//
// Solidity: function withdraw(uint256 amount, bool claim) returns(bool)
func (_CrvRewards *CrvRewardsSession) Withdraw(amount *big.Int, claim bool) (*types.Transaction, error) {
	return _CrvRewards.Contract.Withdraw(&_CrvRewards.TransactOpts, amount, claim)
}

// Withdraw is a paid mutator transaction binding the contract method 0x38d07436.
//
// Solidity: function withdraw(uint256 amount, bool claim) returns(bool)
func (_CrvRewards *CrvRewardsTransactorSession) Withdraw(amount *big.Int, claim bool) (*types.Transaction, error) {
	return _CrvRewards.Contract.Withdraw(&_CrvRewards.TransactOpts, amount, claim)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x1c1c6fe5.
//
// Solidity: function withdrawAll(bool claim) returns()
func (_CrvRewards *CrvRewardsTransactor) WithdrawAll(opts *bind.TransactOpts, claim bool) (*types.Transaction, error) {
	return _CrvRewards.contract.Transact(opts, "withdrawAll", claim)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x1c1c6fe5.
//
// Solidity: function withdrawAll(bool claim) returns()
func (_CrvRewards *CrvRewardsSession) WithdrawAll(claim bool) (*types.Transaction, error) {
	return _CrvRewards.Contract.WithdrawAll(&_CrvRewards.TransactOpts, claim)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x1c1c6fe5.
//
// Solidity: function withdrawAll(bool claim) returns()
func (_CrvRewards *CrvRewardsTransactorSession) WithdrawAll(claim bool) (*types.Transaction, error) {
	return _CrvRewards.Contract.WithdrawAll(&_CrvRewards.TransactOpts, claim)
}

// WithdrawAllAndUnwrap is a paid mutator transaction binding the contract method 0x49f039a2.
//
// Solidity: function withdrawAllAndUnwrap(bool claim) returns()
func (_CrvRewards *CrvRewardsTransactor) WithdrawAllAndUnwrap(opts *bind.TransactOpts, claim bool) (*types.Transaction, error) {
	return _CrvRewards.contract.Transact(opts, "withdrawAllAndUnwrap", claim)
}

// WithdrawAllAndUnwrap is a paid mutator transaction binding the contract method 0x49f039a2.
//
// Solidity: function withdrawAllAndUnwrap(bool claim) returns()
func (_CrvRewards *CrvRewardsSession) WithdrawAllAndUnwrap(claim bool) (*types.Transaction, error) {
	return _CrvRewards.Contract.WithdrawAllAndUnwrap(&_CrvRewards.TransactOpts, claim)
}

// WithdrawAllAndUnwrap is a paid mutator transaction binding the contract method 0x49f039a2.
//
// Solidity: function withdrawAllAndUnwrap(bool claim) returns()
func (_CrvRewards *CrvRewardsTransactorSession) WithdrawAllAndUnwrap(claim bool) (*types.Transaction, error) {
	return _CrvRewards.Contract.WithdrawAllAndUnwrap(&_CrvRewards.TransactOpts, claim)
}

// WithdrawAndUnwrap is a paid mutator transaction binding the contract method 0xc32e7202.
//
// Solidity: function withdrawAndUnwrap(uint256 amount, bool claim) returns(bool)
func (_CrvRewards *CrvRewardsTransactor) WithdrawAndUnwrap(opts *bind.TransactOpts, amount *big.Int, claim bool) (*types.Transaction, error) {
	return _CrvRewards.contract.Transact(opts, "withdrawAndUnwrap", amount, claim)
}

// WithdrawAndUnwrap is a paid mutator transaction binding the contract method 0xc32e7202.
//
// Solidity: function withdrawAndUnwrap(uint256 amount, bool claim) returns(bool)
func (_CrvRewards *CrvRewardsSession) WithdrawAndUnwrap(amount *big.Int, claim bool) (*types.Transaction, error) {
	return _CrvRewards.Contract.WithdrawAndUnwrap(&_CrvRewards.TransactOpts, amount, claim)
}

// WithdrawAndUnwrap is a paid mutator transaction binding the contract method 0xc32e7202.
//
// Solidity: function withdrawAndUnwrap(uint256 amount, bool claim) returns(bool)
func (_CrvRewards *CrvRewardsTransactorSession) WithdrawAndUnwrap(amount *big.Int, claim bool) (*types.Transaction, error) {
	return _CrvRewards.Contract.WithdrawAndUnwrap(&_CrvRewards.TransactOpts, amount, claim)
}

// CrvRewardsRewardAddedIterator is returned from FilterRewardAdded and is used to iterate over the raw logs and unpacked data for RewardAdded events raised by the CrvRewards contract.
type CrvRewardsRewardAddedIterator struct {
	Event *CrvRewardsRewardAdded // Event containing the contract specifics and raw log

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
func (it *CrvRewardsRewardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvRewardsRewardAdded)
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
		it.Event = new(CrvRewardsRewardAdded)
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
func (it *CrvRewardsRewardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvRewardsRewardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvRewardsRewardAdded represents a RewardAdded event raised by the CrvRewards contract.
type CrvRewardsRewardAdded struct {
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardAdded is a free log retrieval operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_CrvRewards *CrvRewardsFilterer) FilterRewardAdded(opts *bind.FilterOpts) (*CrvRewardsRewardAddedIterator, error) {

	logs, sub, err := _CrvRewards.contract.FilterLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return &CrvRewardsRewardAddedIterator{contract: _CrvRewards.contract, event: "RewardAdded", logs: logs, sub: sub}, nil
}

// WatchRewardAdded is a free log subscription operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_CrvRewards *CrvRewardsFilterer) WatchRewardAdded(opts *bind.WatchOpts, sink chan<- *CrvRewardsRewardAdded) (event.Subscription, error) {

	logs, sub, err := _CrvRewards.contract.WatchLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvRewardsRewardAdded)
				if err := _CrvRewards.contract.UnpackLog(event, "RewardAdded", log); err != nil {
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
func (_CrvRewards *CrvRewardsFilterer) ParseRewardAdded(log types.Log) (*CrvRewardsRewardAdded, error) {
	event := new(CrvRewardsRewardAdded)
	if err := _CrvRewards.contract.UnpackLog(event, "RewardAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrvRewardsRewardPaidIterator is returned from FilterRewardPaid and is used to iterate over the raw logs and unpacked data for RewardPaid events raised by the CrvRewards contract.
type CrvRewardsRewardPaidIterator struct {
	Event *CrvRewardsRewardPaid // Event containing the contract specifics and raw log

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
func (it *CrvRewardsRewardPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvRewardsRewardPaid)
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
		it.Event = new(CrvRewardsRewardPaid)
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
func (it *CrvRewardsRewardPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvRewardsRewardPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvRewardsRewardPaid represents a RewardPaid event raised by the CrvRewards contract.
type CrvRewardsRewardPaid struct {
	User   common.Address
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardPaid is a free log retrieval operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_CrvRewards *CrvRewardsFilterer) FilterRewardPaid(opts *bind.FilterOpts, user []common.Address) (*CrvRewardsRewardPaidIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _CrvRewards.contract.FilterLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return &CrvRewardsRewardPaidIterator{contract: _CrvRewards.contract, event: "RewardPaid", logs: logs, sub: sub}, nil
}

// WatchRewardPaid is a free log subscription operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_CrvRewards *CrvRewardsFilterer) WatchRewardPaid(opts *bind.WatchOpts, sink chan<- *CrvRewardsRewardPaid, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _CrvRewards.contract.WatchLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvRewardsRewardPaid)
				if err := _CrvRewards.contract.UnpackLog(event, "RewardPaid", log); err != nil {
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
func (_CrvRewards *CrvRewardsFilterer) ParseRewardPaid(log types.Log) (*CrvRewardsRewardPaid, error) {
	event := new(CrvRewardsRewardPaid)
	if err := _CrvRewards.contract.UnpackLog(event, "RewardPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrvRewardsStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the CrvRewards contract.
type CrvRewardsStakedIterator struct {
	Event *CrvRewardsStaked // Event containing the contract specifics and raw log

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
func (it *CrvRewardsStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvRewardsStaked)
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
		it.Event = new(CrvRewardsStaked)
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
func (it *CrvRewardsStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvRewardsStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvRewardsStaked represents a Staked event raised by the CrvRewards contract.
type CrvRewardsStaked struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_CrvRewards *CrvRewardsFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address) (*CrvRewardsStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _CrvRewards.contract.FilterLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return &CrvRewardsStakedIterator{contract: _CrvRewards.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_CrvRewards *CrvRewardsFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *CrvRewardsStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _CrvRewards.contract.WatchLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvRewardsStaked)
				if err := _CrvRewards.contract.UnpackLog(event, "Staked", log); err != nil {
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
func (_CrvRewards *CrvRewardsFilterer) ParseStaked(log types.Log) (*CrvRewardsStaked, error) {
	event := new(CrvRewardsStaked)
	if err := _CrvRewards.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrvRewardsWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the CrvRewards contract.
type CrvRewardsWithdrawnIterator struct {
	Event *CrvRewardsWithdrawn // Event containing the contract specifics and raw log

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
func (it *CrvRewardsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvRewardsWithdrawn)
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
		it.Event = new(CrvRewardsWithdrawn)
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
func (it *CrvRewardsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvRewardsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvRewardsWithdrawn represents a Withdrawn event raised by the CrvRewards contract.
type CrvRewardsWithdrawn struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_CrvRewards *CrvRewardsFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address) (*CrvRewardsWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _CrvRewards.contract.FilterLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &CrvRewardsWithdrawnIterator{contract: _CrvRewards.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_CrvRewards *CrvRewardsFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *CrvRewardsWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _CrvRewards.contract.WatchLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvRewardsWithdrawn)
				if err := _CrvRewards.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
func (_CrvRewards *CrvRewardsFilterer) ParseWithdrawn(log types.Log) (*CrvRewardsWithdrawn, error) {
	event := new(CrvRewardsWithdrawn)
	if err := _CrvRewards.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
