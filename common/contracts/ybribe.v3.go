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

// YBribeV3MetaData contains all meta data concerning the YBribeV3 contract.
var YBribeV3MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"Blacklisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ChangeOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"ClearRewardRecipient\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"FeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reward_token\",\"type\":\"address\"}],\"name\":\"NewTokenReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bias\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blacklisted_bias\",\"type\":\"uint256\"}],\"name\":\"PeriodUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"RemovedFromBlacklist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"briber\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reward_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"RewardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reward_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"SetRewardRecipient\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_gauges_per_reward\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_rewards_in_gauge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_rewards_per_gauge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accept_owner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"active_period\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reward_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"add_reward_amount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"add_to_blacklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reward_token\",\"type\":\"address\"}],\"name\":\"claim_reward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reward_token\",\"type\":\"address\"}],\"name\":\"claim_reward_for\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_users\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_gauges\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_reward_tokens\",\"type\":\"address[]\"}],\"name\":\"claim_reward_for_many\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reward_token\",\"type\":\"address\"}],\"name\":\"claimable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claims_per_gauge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clear_recipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"current_period\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee_percent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee_recipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"}],\"name\":\"gauges_per_reward\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_blacklist\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_blacklist\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"}],\"name\":\"get_blacklisted_bias\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"address_to_check\",\"type\":\"address\"}],\"name\":\"is_blacklisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"last_user_claim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"next_claim_time\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pending_owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"remove_from_blacklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"reward_per_gauge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"reward_per_token\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"reward_recipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"}],\"name\":\"rewards_per_gauge\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_percent\",\"type\":\"uint256\"}],\"name\":\"set_fee_percent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"set_fee_recipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_new_owner\",\"type\":\"address\"}],\"name\":\"set_owner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"set_recipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// YBribeV3ABI is the input ABI used to generate the binding from.
// Deprecated: Use YBribeV3MetaData.ABI instead.
var YBribeV3ABI = YBribeV3MetaData.ABI

// YBribeV3 is an auto generated Go binding around an Ethereum contract.
type YBribeV3 struct {
	YBribeV3Caller     // Read-only binding to the contract
	YBribeV3Transactor // Write-only binding to the contract
	YBribeV3Filterer   // Log filterer for contract events
}

// YBribeV3Caller is an auto generated read-only Go binding around an Ethereum contract.
type YBribeV3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YBribeV3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type YBribeV3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YBribeV3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YBribeV3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YBribeV3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YBribeV3Session struct {
	Contract     *YBribeV3         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YBribeV3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YBribeV3CallerSession struct {
	Contract *YBribeV3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// YBribeV3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YBribeV3TransactorSession struct {
	Contract     *YBribeV3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// YBribeV3Raw is an auto generated low-level Go binding around an Ethereum contract.
type YBribeV3Raw struct {
	Contract *YBribeV3 // Generic contract binding to access the raw methods on
}

// YBribeV3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YBribeV3CallerRaw struct {
	Contract *YBribeV3Caller // Generic read-only contract binding to access the raw methods on
}

// YBribeV3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YBribeV3TransactorRaw struct {
	Contract *YBribeV3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewYBribeV3 creates a new instance of YBribeV3, bound to a specific deployed contract.
func NewYBribeV3(address common.Address, backend bind.ContractBackend) (*YBribeV3, error) {
	contract, err := bindYBribeV3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YBribeV3{YBribeV3Caller: YBribeV3Caller{contract: contract}, YBribeV3Transactor: YBribeV3Transactor{contract: contract}, YBribeV3Filterer: YBribeV3Filterer{contract: contract}}, nil
}

// NewYBribeV3Caller creates a new read-only instance of YBribeV3, bound to a specific deployed contract.
func NewYBribeV3Caller(address common.Address, caller bind.ContractCaller) (*YBribeV3Caller, error) {
	contract, err := bindYBribeV3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YBribeV3Caller{contract: contract}, nil
}

// NewYBribeV3Transactor creates a new write-only instance of YBribeV3, bound to a specific deployed contract.
func NewYBribeV3Transactor(address common.Address, transactor bind.ContractTransactor) (*YBribeV3Transactor, error) {
	contract, err := bindYBribeV3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YBribeV3Transactor{contract: contract}, nil
}

// NewYBribeV3Filterer creates a new log filterer instance of YBribeV3, bound to a specific deployed contract.
func NewYBribeV3Filterer(address common.Address, filterer bind.ContractFilterer) (*YBribeV3Filterer, error) {
	contract, err := bindYBribeV3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YBribeV3Filterer{contract: contract}, nil
}

// bindYBribeV3 binds a generic wrapper to an already deployed contract.
func bindYBribeV3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(YBribeV3ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YBribeV3 *YBribeV3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YBribeV3.Contract.YBribeV3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YBribeV3 *YBribeV3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YBribeV3.Contract.YBribeV3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YBribeV3 *YBribeV3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YBribeV3.Contract.YBribeV3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YBribeV3 *YBribeV3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YBribeV3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YBribeV3 *YBribeV3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YBribeV3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YBribeV3 *YBribeV3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YBribeV3.Contract.contract.Transact(opts, method, params...)
}

// GaugesPerReward is a free data retrieval call binding the contract method 0xbdd48147.
//
// Solidity: function _gauges_per_reward(address , uint256 ) view returns(address)
func (_YBribeV3 *YBribeV3Caller) GaugesPerReward(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _YBribeV3.contract.Call(opts, &out, "_gauges_per_reward", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GaugesPerReward is a free data retrieval call binding the contract method 0xbdd48147.
//
// Solidity: function _gauges_per_reward(address , uint256 ) view returns(address)
func (_YBribeV3 *YBribeV3Session) GaugesPerReward(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _YBribeV3.Contract.GaugesPerReward(&_YBribeV3.CallOpts, arg0, arg1)
}

// GaugesPerReward is a free data retrieval call binding the contract method 0xbdd48147.
//
// Solidity: function _gauges_per_reward(address , uint256 ) view returns(address)
func (_YBribeV3 *YBribeV3CallerSession) GaugesPerReward(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _YBribeV3.Contract.GaugesPerReward(&_YBribeV3.CallOpts, arg0, arg1)
}

// RewardsInGauge is a free data retrieval call binding the contract method 0xed2b5dc8.
//
// Solidity: function _rewards_in_gauge(address , address ) view returns(bool)
func (_YBribeV3 *YBribeV3Caller) RewardsInGauge(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _YBribeV3.contract.Call(opts, &out, "_rewards_in_gauge", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RewardsInGauge is a free data retrieval call binding the contract method 0xed2b5dc8.
//
// Solidity: function _rewards_in_gauge(address , address ) view returns(bool)
func (_YBribeV3 *YBribeV3Session) RewardsInGauge(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _YBribeV3.Contract.RewardsInGauge(&_YBribeV3.CallOpts, arg0, arg1)
}

// RewardsInGauge is a free data retrieval call binding the contract method 0xed2b5dc8.
//
// Solidity: function _rewards_in_gauge(address , address ) view returns(bool)
func (_YBribeV3 *YBribeV3CallerSession) RewardsInGauge(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _YBribeV3.Contract.RewardsInGauge(&_YBribeV3.CallOpts, arg0, arg1)
}

// YbribeRewardsPerGauge is a free data retrieval call binding the contract method 0xc0ca1fca.
//
// Solidity: function _rewards_per_gauge(address , uint256 ) view returns(address)
func (_YBribeV3 *YBribeV3Caller) YbribeRewardsPerGauge(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _YBribeV3.contract.Call(opts, &out, "_rewards_per_gauge", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// YbribeRewardsPerGauge is a free data retrieval call binding the contract method 0xc0ca1fca.
//
// Solidity: function _rewards_per_gauge(address , uint256 ) view returns(address)
func (_YBribeV3 *YBribeV3Session) YbribeRewardsPerGauge(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _YBribeV3.Contract.YbribeRewardsPerGauge(&_YBribeV3.CallOpts, arg0, arg1)
}

// YbribeRewardsPerGauge is a free data retrieval call binding the contract method 0xc0ca1fca.
//
// Solidity: function _rewards_per_gauge(address , uint256 ) view returns(address)
func (_YBribeV3 *YBribeV3CallerSession) YbribeRewardsPerGauge(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _YBribeV3.Contract.YbribeRewardsPerGauge(&_YBribeV3.CallOpts, arg0, arg1)
}

// ActivePeriod is a free data retrieval call binding the contract method 0x2de24ac7.
//
// Solidity: function active_period(address , address ) view returns(uint256)
func (_YBribeV3 *YBribeV3Caller) ActivePeriod(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YBribeV3.contract.Call(opts, &out, "active_period", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActivePeriod is a free data retrieval call binding the contract method 0x2de24ac7.
//
// Solidity: function active_period(address , address ) view returns(uint256)
func (_YBribeV3 *YBribeV3Session) ActivePeriod(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _YBribeV3.Contract.ActivePeriod(&_YBribeV3.CallOpts, arg0, arg1)
}

// ActivePeriod is a free data retrieval call binding the contract method 0x2de24ac7.
//
// Solidity: function active_period(address , address ) view returns(uint256)
func (_YBribeV3 *YBribeV3CallerSession) ActivePeriod(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _YBribeV3.Contract.ActivePeriod(&_YBribeV3.CallOpts, arg0, arg1)
}

// Claimable is a free data retrieval call binding the contract method 0x5fa135d9.
//
// Solidity: function claimable(address user, address gauge, address reward_token) view returns(uint256)
func (_YBribeV3 *YBribeV3Caller) Claimable(opts *bind.CallOpts, user common.Address, gauge common.Address, reward_token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YBribeV3.contract.Call(opts, &out, "claimable", user, gauge, reward_token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Claimable is a free data retrieval call binding the contract method 0x5fa135d9.
//
// Solidity: function claimable(address user, address gauge, address reward_token) view returns(uint256)
func (_YBribeV3 *YBribeV3Session) Claimable(user common.Address, gauge common.Address, reward_token common.Address) (*big.Int, error) {
	return _YBribeV3.Contract.Claimable(&_YBribeV3.CallOpts, user, gauge, reward_token)
}

// Claimable is a free data retrieval call binding the contract method 0x5fa135d9.
//
// Solidity: function claimable(address user, address gauge, address reward_token) view returns(uint256)
func (_YBribeV3 *YBribeV3CallerSession) Claimable(user common.Address, gauge common.Address, reward_token common.Address) (*big.Int, error) {
	return _YBribeV3.Contract.Claimable(&_YBribeV3.CallOpts, user, gauge, reward_token)
}

// ClaimsPerGauge is a free data retrieval call binding the contract method 0xf703789c.
//
// Solidity: function claims_per_gauge(address , address ) view returns(uint256)
func (_YBribeV3 *YBribeV3Caller) ClaimsPerGauge(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YBribeV3.contract.Call(opts, &out, "claims_per_gauge", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimsPerGauge is a free data retrieval call binding the contract method 0xf703789c.
//
// Solidity: function claims_per_gauge(address , address ) view returns(uint256)
func (_YBribeV3 *YBribeV3Session) ClaimsPerGauge(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _YBribeV3.Contract.ClaimsPerGauge(&_YBribeV3.CallOpts, arg0, arg1)
}

// ClaimsPerGauge is a free data retrieval call binding the contract method 0xf703789c.
//
// Solidity: function claims_per_gauge(address , address ) view returns(uint256)
func (_YBribeV3 *YBribeV3CallerSession) ClaimsPerGauge(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _YBribeV3.Contract.ClaimsPerGauge(&_YBribeV3.CallOpts, arg0, arg1)
}

// CurrentPeriod is a free data retrieval call binding the contract method 0x431dc4b6.
//
// Solidity: function current_period() view returns(uint256)
func (_YBribeV3 *YBribeV3Caller) CurrentPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YBribeV3.contract.Call(opts, &out, "current_period")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentPeriod is a free data retrieval call binding the contract method 0x431dc4b6.
//
// Solidity: function current_period() view returns(uint256)
func (_YBribeV3 *YBribeV3Session) CurrentPeriod() (*big.Int, error) {
	return _YBribeV3.Contract.CurrentPeriod(&_YBribeV3.CallOpts)
}

// CurrentPeriod is a free data retrieval call binding the contract method 0x431dc4b6.
//
// Solidity: function current_period() view returns(uint256)
func (_YBribeV3 *YBribeV3CallerSession) CurrentPeriod() (*big.Int, error) {
	return _YBribeV3.Contract.CurrentPeriod(&_YBribeV3.CallOpts)
}

// FeePercent is a free data retrieval call binding the contract method 0x4ff2cc44.
//
// Solidity: function fee_percent() view returns(uint256)
func (_YBribeV3 *YBribeV3Caller) FeePercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YBribeV3.contract.Call(opts, &out, "fee_percent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeePercent is a free data retrieval call binding the contract method 0x4ff2cc44.
//
// Solidity: function fee_percent() view returns(uint256)
func (_YBribeV3 *YBribeV3Session) FeePercent() (*big.Int, error) {
	return _YBribeV3.Contract.FeePercent(&_YBribeV3.CallOpts)
}

// FeePercent is a free data retrieval call binding the contract method 0x4ff2cc44.
//
// Solidity: function fee_percent() view returns(uint256)
func (_YBribeV3 *YBribeV3CallerSession) FeePercent() (*big.Int, error) {
	return _YBribeV3.Contract.FeePercent(&_YBribeV3.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x758f969a.
//
// Solidity: function fee_recipient() view returns(address)
func (_YBribeV3 *YBribeV3Caller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YBribeV3.contract.Call(opts, &out, "fee_recipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x758f969a.
//
// Solidity: function fee_recipient() view returns(address)
func (_YBribeV3 *YBribeV3Session) FeeRecipient() (common.Address, error) {
	return _YBribeV3.Contract.FeeRecipient(&_YBribeV3.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x758f969a.
//
// Solidity: function fee_recipient() view returns(address)
func (_YBribeV3 *YBribeV3CallerSession) FeeRecipient() (common.Address, error) {
	return _YBribeV3.Contract.FeeRecipient(&_YBribeV3.CallOpts)
}

// YbribeGaugesPerReward is a free data retrieval call binding the contract method 0x32be6791.
//
// Solidity: function gauges_per_reward(address reward) view returns(address[])
func (_YBribeV3 *YBribeV3Caller) YbribeGaugesPerReward(opts *bind.CallOpts, reward common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _YBribeV3.contract.Call(opts, &out, "gauges_per_reward", reward)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// YbribeGaugesPerReward is a free data retrieval call binding the contract method 0x32be6791.
//
// Solidity: function gauges_per_reward(address reward) view returns(address[])
func (_YBribeV3 *YBribeV3Session) YbribeGaugesPerReward(reward common.Address) ([]common.Address, error) {
	return _YBribeV3.Contract.YbribeGaugesPerReward(&_YBribeV3.CallOpts, reward)
}

// YbribeGaugesPerReward is a free data retrieval call binding the contract method 0x32be6791.
//
// Solidity: function gauges_per_reward(address reward) view returns(address[])
func (_YBribeV3 *YBribeV3CallerSession) YbribeGaugesPerReward(reward common.Address) ([]common.Address, error) {
	return _YBribeV3.Contract.YbribeGaugesPerReward(&_YBribeV3.CallOpts, reward)
}

// GetBlacklist is a free data retrieval call binding the contract method 0xee2f892e.
//
// Solidity: function get_blacklist() view returns(address[] _blacklist)
func (_YBribeV3 *YBribeV3Caller) GetBlacklist(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _YBribeV3.contract.Call(opts, &out, "get_blacklist")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetBlacklist is a free data retrieval call binding the contract method 0xee2f892e.
//
// Solidity: function get_blacklist() view returns(address[] _blacklist)
func (_YBribeV3 *YBribeV3Session) GetBlacklist() ([]common.Address, error) {
	return _YBribeV3.Contract.GetBlacklist(&_YBribeV3.CallOpts)
}

// GetBlacklist is a free data retrieval call binding the contract method 0xee2f892e.
//
// Solidity: function get_blacklist() view returns(address[] _blacklist)
func (_YBribeV3 *YBribeV3CallerSession) GetBlacklist() ([]common.Address, error) {
	return _YBribeV3.Contract.GetBlacklist(&_YBribeV3.CallOpts)
}

// GetBlacklistedBias is a free data retrieval call binding the contract method 0x6885f4ee.
//
// Solidity: function get_blacklisted_bias(address gauge) view returns(uint256)
func (_YBribeV3 *YBribeV3Caller) GetBlacklistedBias(opts *bind.CallOpts, gauge common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YBribeV3.contract.Call(opts, &out, "get_blacklisted_bias", gauge)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlacklistedBias is a free data retrieval call binding the contract method 0x6885f4ee.
//
// Solidity: function get_blacklisted_bias(address gauge) view returns(uint256)
func (_YBribeV3 *YBribeV3Session) GetBlacklistedBias(gauge common.Address) (*big.Int, error) {
	return _YBribeV3.Contract.GetBlacklistedBias(&_YBribeV3.CallOpts, gauge)
}

// GetBlacklistedBias is a free data retrieval call binding the contract method 0x6885f4ee.
//
// Solidity: function get_blacklisted_bias(address gauge) view returns(uint256)
func (_YBribeV3 *YBribeV3CallerSession) GetBlacklistedBias(gauge common.Address) (*big.Int, error) {
	return _YBribeV3.Contract.GetBlacklistedBias(&_YBribeV3.CallOpts, gauge)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0x1794e95a.
//
// Solidity: function is_blacklisted(address address_to_check) view returns(bool)
func (_YBribeV3 *YBribeV3Caller) IsBlacklisted(opts *bind.CallOpts, address_to_check common.Address) (bool, error) {
	var out []interface{}
	err := _YBribeV3.contract.Call(opts, &out, "is_blacklisted", address_to_check)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlacklisted is a free data retrieval call binding the contract method 0x1794e95a.
//
// Solidity: function is_blacklisted(address address_to_check) view returns(bool)
func (_YBribeV3 *YBribeV3Session) IsBlacklisted(address_to_check common.Address) (bool, error) {
	return _YBribeV3.Contract.IsBlacklisted(&_YBribeV3.CallOpts, address_to_check)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0x1794e95a.
//
// Solidity: function is_blacklisted(address address_to_check) view returns(bool)
func (_YBribeV3 *YBribeV3CallerSession) IsBlacklisted(address_to_check common.Address) (bool, error) {
	return _YBribeV3.Contract.IsBlacklisted(&_YBribeV3.CallOpts, address_to_check)
}

// LastUserClaim is a free data retrieval call binding the contract method 0xe9773ad0.
//
// Solidity: function last_user_claim(address , address , address ) view returns(uint256)
func (_YBribeV3 *YBribeV3Caller) LastUserClaim(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YBribeV3.contract.Call(opts, &out, "last_user_claim", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUserClaim is a free data retrieval call binding the contract method 0xe9773ad0.
//
// Solidity: function last_user_claim(address , address , address ) view returns(uint256)
func (_YBribeV3 *YBribeV3Session) LastUserClaim(arg0 common.Address, arg1 common.Address, arg2 common.Address) (*big.Int, error) {
	return _YBribeV3.Contract.LastUserClaim(&_YBribeV3.CallOpts, arg0, arg1, arg2)
}

// LastUserClaim is a free data retrieval call binding the contract method 0xe9773ad0.
//
// Solidity: function last_user_claim(address , address , address ) view returns(uint256)
func (_YBribeV3 *YBribeV3CallerSession) LastUserClaim(arg0 common.Address, arg1 common.Address, arg2 common.Address) (*big.Int, error) {
	return _YBribeV3.Contract.LastUserClaim(&_YBribeV3.CallOpts, arg0, arg1, arg2)
}

// NextClaimTime is a free data retrieval call binding the contract method 0xff40e773.
//
// Solidity: function next_claim_time(address ) view returns(uint256)
func (_YBribeV3 *YBribeV3Caller) NextClaimTime(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YBribeV3.contract.Call(opts, &out, "next_claim_time", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextClaimTime is a free data retrieval call binding the contract method 0xff40e773.
//
// Solidity: function next_claim_time(address ) view returns(uint256)
func (_YBribeV3 *YBribeV3Session) NextClaimTime(arg0 common.Address) (*big.Int, error) {
	return _YBribeV3.Contract.NextClaimTime(&_YBribeV3.CallOpts, arg0)
}

// NextClaimTime is a free data retrieval call binding the contract method 0xff40e773.
//
// Solidity: function next_claim_time(address ) view returns(uint256)
func (_YBribeV3 *YBribeV3CallerSession) NextClaimTime(arg0 common.Address) (*big.Int, error) {
	return _YBribeV3.Contract.NextClaimTime(&_YBribeV3.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_YBribeV3 *YBribeV3Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YBribeV3.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_YBribeV3 *YBribeV3Session) Owner() (common.Address, error) {
	return _YBribeV3.Contract.Owner(&_YBribeV3.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_YBribeV3 *YBribeV3CallerSession) Owner() (common.Address, error) {
	return _YBribeV3.Contract.Owner(&_YBribeV3.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0x7f4ec4c3.
//
// Solidity: function pending_owner() view returns(address)
func (_YBribeV3 *YBribeV3Caller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YBribeV3.contract.Call(opts, &out, "pending_owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0x7f4ec4c3.
//
// Solidity: function pending_owner() view returns(address)
func (_YBribeV3 *YBribeV3Session) PendingOwner() (common.Address, error) {
	return _YBribeV3.Contract.PendingOwner(&_YBribeV3.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0x7f4ec4c3.
//
// Solidity: function pending_owner() view returns(address)
func (_YBribeV3 *YBribeV3CallerSession) PendingOwner() (common.Address, error) {
	return _YBribeV3.Contract.PendingOwner(&_YBribeV3.CallOpts)
}

// RewardPerGauge is a free data retrieval call binding the contract method 0x6938ae40.
//
// Solidity: function reward_per_gauge(address , address ) view returns(uint256)
func (_YBribeV3 *YBribeV3Caller) RewardPerGauge(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YBribeV3.contract.Call(opts, &out, "reward_per_gauge", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerGauge is a free data retrieval call binding the contract method 0x6938ae40.
//
// Solidity: function reward_per_gauge(address , address ) view returns(uint256)
func (_YBribeV3 *YBribeV3Session) RewardPerGauge(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _YBribeV3.Contract.RewardPerGauge(&_YBribeV3.CallOpts, arg0, arg1)
}

// RewardPerGauge is a free data retrieval call binding the contract method 0x6938ae40.
//
// Solidity: function reward_per_gauge(address , address ) view returns(uint256)
func (_YBribeV3 *YBribeV3CallerSession) RewardPerGauge(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _YBribeV3.Contract.RewardPerGauge(&_YBribeV3.CallOpts, arg0, arg1)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xa70a61c2.
//
// Solidity: function reward_per_token(address , address ) view returns(uint256)
func (_YBribeV3 *YBribeV3Caller) RewardPerToken(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YBribeV3.contract.Call(opts, &out, "reward_per_token", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerToken is a free data retrieval call binding the contract method 0xa70a61c2.
//
// Solidity: function reward_per_token(address , address ) view returns(uint256)
func (_YBribeV3 *YBribeV3Session) RewardPerToken(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _YBribeV3.Contract.RewardPerToken(&_YBribeV3.CallOpts, arg0, arg1)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xa70a61c2.
//
// Solidity: function reward_per_token(address , address ) view returns(uint256)
func (_YBribeV3 *YBribeV3CallerSession) RewardPerToken(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _YBribeV3.Contract.RewardPerToken(&_YBribeV3.CallOpts, arg0, arg1)
}

// RewardRecipient is a free data retrieval call binding the contract method 0x72b1f4e3.
//
// Solidity: function reward_recipient(address ) view returns(address)
func (_YBribeV3 *YBribeV3Caller) RewardRecipient(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _YBribeV3.contract.Call(opts, &out, "reward_recipient", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardRecipient is a free data retrieval call binding the contract method 0x72b1f4e3.
//
// Solidity: function reward_recipient(address ) view returns(address)
func (_YBribeV3 *YBribeV3Session) RewardRecipient(arg0 common.Address) (common.Address, error) {
	return _YBribeV3.Contract.RewardRecipient(&_YBribeV3.CallOpts, arg0)
}

// RewardRecipient is a free data retrieval call binding the contract method 0x72b1f4e3.
//
// Solidity: function reward_recipient(address ) view returns(address)
func (_YBribeV3 *YBribeV3CallerSession) RewardRecipient(arg0 common.Address) (common.Address, error) {
	return _YBribeV3.Contract.RewardRecipient(&_YBribeV3.CallOpts, arg0)
}

// RewardsPerGauge is a free data retrieval call binding the contract method 0x5dff2e13.
//
// Solidity: function rewards_per_gauge(address gauge) view returns(address[])
func (_YBribeV3 *YBribeV3Caller) RewardsPerGauge(opts *bind.CallOpts, gauge common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _YBribeV3.contract.Call(opts, &out, "rewards_per_gauge", gauge)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// RewardsPerGauge is a free data retrieval call binding the contract method 0x5dff2e13.
//
// Solidity: function rewards_per_gauge(address gauge) view returns(address[])
func (_YBribeV3 *YBribeV3Session) RewardsPerGauge(gauge common.Address) ([]common.Address, error) {
	return _YBribeV3.Contract.RewardsPerGauge(&_YBribeV3.CallOpts, gauge)
}

// RewardsPerGauge is a free data retrieval call binding the contract method 0x5dff2e13.
//
// Solidity: function rewards_per_gauge(address gauge) view returns(address[])
func (_YBribeV3 *YBribeV3CallerSession) RewardsPerGauge(gauge common.Address) ([]common.Address, error) {
	return _YBribeV3.Contract.RewardsPerGauge(&_YBribeV3.CallOpts, gauge)
}

// AcceptOwner is a paid mutator transaction binding the contract method 0xb2c1e1de.
//
// Solidity: function accept_owner() returns()
func (_YBribeV3 *YBribeV3Transactor) AcceptOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YBribeV3.contract.Transact(opts, "accept_owner")
}

// AcceptOwner is a paid mutator transaction binding the contract method 0xb2c1e1de.
//
// Solidity: function accept_owner() returns()
func (_YBribeV3 *YBribeV3Session) AcceptOwner() (*types.Transaction, error) {
	return _YBribeV3.Contract.AcceptOwner(&_YBribeV3.TransactOpts)
}

// AcceptOwner is a paid mutator transaction binding the contract method 0xb2c1e1de.
//
// Solidity: function accept_owner() returns()
func (_YBribeV3 *YBribeV3TransactorSession) AcceptOwner() (*types.Transaction, error) {
	return _YBribeV3.Contract.AcceptOwner(&_YBribeV3.TransactOpts)
}

// AddRewardAmount is a paid mutator transaction binding the contract method 0xdf406967.
//
// Solidity: function add_reward_amount(address gauge, address reward_token, uint256 amount) returns(bool)
func (_YBribeV3 *YBribeV3Transactor) AddRewardAmount(opts *bind.TransactOpts, gauge common.Address, reward_token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YBribeV3.contract.Transact(opts, "add_reward_amount", gauge, reward_token, amount)
}

// AddRewardAmount is a paid mutator transaction binding the contract method 0xdf406967.
//
// Solidity: function add_reward_amount(address gauge, address reward_token, uint256 amount) returns(bool)
func (_YBribeV3 *YBribeV3Session) AddRewardAmount(gauge common.Address, reward_token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YBribeV3.Contract.AddRewardAmount(&_YBribeV3.TransactOpts, gauge, reward_token, amount)
}

// AddRewardAmount is a paid mutator transaction binding the contract method 0xdf406967.
//
// Solidity: function add_reward_amount(address gauge, address reward_token, uint256 amount) returns(bool)
func (_YBribeV3 *YBribeV3TransactorSession) AddRewardAmount(gauge common.Address, reward_token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YBribeV3.Contract.AddRewardAmount(&_YBribeV3.TransactOpts, gauge, reward_token, amount)
}

// AddToBlacklist is a paid mutator transaction binding the contract method 0xb0b34c23.
//
// Solidity: function add_to_blacklist(address _user) returns()
func (_YBribeV3 *YBribeV3Transactor) AddToBlacklist(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _YBribeV3.contract.Transact(opts, "add_to_blacklist", _user)
}

// AddToBlacklist is a paid mutator transaction binding the contract method 0xb0b34c23.
//
// Solidity: function add_to_blacklist(address _user) returns()
func (_YBribeV3 *YBribeV3Session) AddToBlacklist(_user common.Address) (*types.Transaction, error) {
	return _YBribeV3.Contract.AddToBlacklist(&_YBribeV3.TransactOpts, _user)
}

// AddToBlacklist is a paid mutator transaction binding the contract method 0xb0b34c23.
//
// Solidity: function add_to_blacklist(address _user) returns()
func (_YBribeV3 *YBribeV3TransactorSession) AddToBlacklist(_user common.Address) (*types.Transaction, error) {
	return _YBribeV3.Contract.AddToBlacklist(&_YBribeV3.TransactOpts, _user)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x750d4926.
//
// Solidity: function claim_reward(address gauge, address reward_token) returns(uint256)
func (_YBribeV3 *YBribeV3Transactor) ClaimReward(opts *bind.TransactOpts, gauge common.Address, reward_token common.Address) (*types.Transaction, error) {
	return _YBribeV3.contract.Transact(opts, "claim_reward", gauge, reward_token)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x750d4926.
//
// Solidity: function claim_reward(address gauge, address reward_token) returns(uint256)
func (_YBribeV3 *YBribeV3Session) ClaimReward(gauge common.Address, reward_token common.Address) (*types.Transaction, error) {
	return _YBribeV3.Contract.ClaimReward(&_YBribeV3.TransactOpts, gauge, reward_token)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x750d4926.
//
// Solidity: function claim_reward(address gauge, address reward_token) returns(uint256)
func (_YBribeV3 *YBribeV3TransactorSession) ClaimReward(gauge common.Address, reward_token common.Address) (*types.Transaction, error) {
	return _YBribeV3.Contract.ClaimReward(&_YBribeV3.TransactOpts, gauge, reward_token)
}

// ClaimRewardFor is a paid mutator transaction binding the contract method 0xa7a914b5.
//
// Solidity: function claim_reward_for(address user, address gauge, address reward_token) returns(uint256)
func (_YBribeV3 *YBribeV3Transactor) ClaimRewardFor(opts *bind.TransactOpts, user common.Address, gauge common.Address, reward_token common.Address) (*types.Transaction, error) {
	return _YBribeV3.contract.Transact(opts, "claim_reward_for", user, gauge, reward_token)
}

// ClaimRewardFor is a paid mutator transaction binding the contract method 0xa7a914b5.
//
// Solidity: function claim_reward_for(address user, address gauge, address reward_token) returns(uint256)
func (_YBribeV3 *YBribeV3Session) ClaimRewardFor(user common.Address, gauge common.Address, reward_token common.Address) (*types.Transaction, error) {
	return _YBribeV3.Contract.ClaimRewardFor(&_YBribeV3.TransactOpts, user, gauge, reward_token)
}

// ClaimRewardFor is a paid mutator transaction binding the contract method 0xa7a914b5.
//
// Solidity: function claim_reward_for(address user, address gauge, address reward_token) returns(uint256)
func (_YBribeV3 *YBribeV3TransactorSession) ClaimRewardFor(user common.Address, gauge common.Address, reward_token common.Address) (*types.Transaction, error) {
	return _YBribeV3.Contract.ClaimRewardFor(&_YBribeV3.TransactOpts, user, gauge, reward_token)
}

// ClaimRewardForMany is a paid mutator transaction binding the contract method 0xb46d4f2e.
//
// Solidity: function claim_reward_for_many(address[] _users, address[] _gauges, address[] _reward_tokens) returns(uint256[] amounts)
func (_YBribeV3 *YBribeV3Transactor) ClaimRewardForMany(opts *bind.TransactOpts, _users []common.Address, _gauges []common.Address, _reward_tokens []common.Address) (*types.Transaction, error) {
	return _YBribeV3.contract.Transact(opts, "claim_reward_for_many", _users, _gauges, _reward_tokens)
}

// ClaimRewardForMany is a paid mutator transaction binding the contract method 0xb46d4f2e.
//
// Solidity: function claim_reward_for_many(address[] _users, address[] _gauges, address[] _reward_tokens) returns(uint256[] amounts)
func (_YBribeV3 *YBribeV3Session) ClaimRewardForMany(_users []common.Address, _gauges []common.Address, _reward_tokens []common.Address) (*types.Transaction, error) {
	return _YBribeV3.Contract.ClaimRewardForMany(&_YBribeV3.TransactOpts, _users, _gauges, _reward_tokens)
}

// ClaimRewardForMany is a paid mutator transaction binding the contract method 0xb46d4f2e.
//
// Solidity: function claim_reward_for_many(address[] _users, address[] _gauges, address[] _reward_tokens) returns(uint256[] amounts)
func (_YBribeV3 *YBribeV3TransactorSession) ClaimRewardForMany(_users []common.Address, _gauges []common.Address, _reward_tokens []common.Address) (*types.Transaction, error) {
	return _YBribeV3.Contract.ClaimRewardForMany(&_YBribeV3.TransactOpts, _users, _gauges, _reward_tokens)
}

// ClearRecipient is a paid mutator transaction binding the contract method 0xa770a359.
//
// Solidity: function clear_recipient() returns()
func (_YBribeV3 *YBribeV3Transactor) ClearRecipient(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YBribeV3.contract.Transact(opts, "clear_recipient")
}

// ClearRecipient is a paid mutator transaction binding the contract method 0xa770a359.
//
// Solidity: function clear_recipient() returns()
func (_YBribeV3 *YBribeV3Session) ClearRecipient() (*types.Transaction, error) {
	return _YBribeV3.Contract.ClearRecipient(&_YBribeV3.TransactOpts)
}

// ClearRecipient is a paid mutator transaction binding the contract method 0xa770a359.
//
// Solidity: function clear_recipient() returns()
func (_YBribeV3 *YBribeV3TransactorSession) ClearRecipient() (*types.Transaction, error) {
	return _YBribeV3.Contract.ClearRecipient(&_YBribeV3.TransactOpts)
}

// RemoveFromBlacklist is a paid mutator transaction binding the contract method 0xbfcc3d1c.
//
// Solidity: function remove_from_blacklist(address _user) returns()
func (_YBribeV3 *YBribeV3Transactor) RemoveFromBlacklist(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _YBribeV3.contract.Transact(opts, "remove_from_blacklist", _user)
}

// RemoveFromBlacklist is a paid mutator transaction binding the contract method 0xbfcc3d1c.
//
// Solidity: function remove_from_blacklist(address _user) returns()
func (_YBribeV3 *YBribeV3Session) RemoveFromBlacklist(_user common.Address) (*types.Transaction, error) {
	return _YBribeV3.Contract.RemoveFromBlacklist(&_YBribeV3.TransactOpts, _user)
}

// RemoveFromBlacklist is a paid mutator transaction binding the contract method 0xbfcc3d1c.
//
// Solidity: function remove_from_blacklist(address _user) returns()
func (_YBribeV3 *YBribeV3TransactorSession) RemoveFromBlacklist(_user common.Address) (*types.Transaction, error) {
	return _YBribeV3.Contract.RemoveFromBlacklist(&_YBribeV3.TransactOpts, _user)
}

// SetFeePercent is a paid mutator transaction binding the contract method 0xca3f7e1e.
//
// Solidity: function set_fee_percent(uint256 _percent) returns()
func (_YBribeV3 *YBribeV3Transactor) SetFeePercent(opts *bind.TransactOpts, _percent *big.Int) (*types.Transaction, error) {
	return _YBribeV3.contract.Transact(opts, "set_fee_percent", _percent)
}

// SetFeePercent is a paid mutator transaction binding the contract method 0xca3f7e1e.
//
// Solidity: function set_fee_percent(uint256 _percent) returns()
func (_YBribeV3 *YBribeV3Session) SetFeePercent(_percent *big.Int) (*types.Transaction, error) {
	return _YBribeV3.Contract.SetFeePercent(&_YBribeV3.TransactOpts, _percent)
}

// SetFeePercent is a paid mutator transaction binding the contract method 0xca3f7e1e.
//
// Solidity: function set_fee_percent(uint256 _percent) returns()
func (_YBribeV3 *YBribeV3TransactorSession) SetFeePercent(_percent *big.Int) (*types.Transaction, error) {
	return _YBribeV3.Contract.SetFeePercent(&_YBribeV3.TransactOpts, _percent)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0x30cc317b.
//
// Solidity: function set_fee_recipient(address _recipient) returns()
func (_YBribeV3 *YBribeV3Transactor) SetFeeRecipient(opts *bind.TransactOpts, _recipient common.Address) (*types.Transaction, error) {
	return _YBribeV3.contract.Transact(opts, "set_fee_recipient", _recipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0x30cc317b.
//
// Solidity: function set_fee_recipient(address _recipient) returns()
func (_YBribeV3 *YBribeV3Session) SetFeeRecipient(_recipient common.Address) (*types.Transaction, error) {
	return _YBribeV3.Contract.SetFeeRecipient(&_YBribeV3.TransactOpts, _recipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0x30cc317b.
//
// Solidity: function set_fee_recipient(address _recipient) returns()
func (_YBribeV3 *YBribeV3TransactorSession) SetFeeRecipient(_recipient common.Address) (*types.Transaction, error) {
	return _YBribeV3.Contract.SetFeeRecipient(&_YBribeV3.TransactOpts, _recipient)
}

// SetOwner is a paid mutator transaction binding the contract method 0x7cb97b2b.
//
// Solidity: function set_owner(address _new_owner) returns()
func (_YBribeV3 *YBribeV3Transactor) SetOwner(opts *bind.TransactOpts, _new_owner common.Address) (*types.Transaction, error) {
	return _YBribeV3.contract.Transact(opts, "set_owner", _new_owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x7cb97b2b.
//
// Solidity: function set_owner(address _new_owner) returns()
func (_YBribeV3 *YBribeV3Session) SetOwner(_new_owner common.Address) (*types.Transaction, error) {
	return _YBribeV3.Contract.SetOwner(&_YBribeV3.TransactOpts, _new_owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x7cb97b2b.
//
// Solidity: function set_owner(address _new_owner) returns()
func (_YBribeV3 *YBribeV3TransactorSession) SetOwner(_new_owner common.Address) (*types.Transaction, error) {
	return _YBribeV3.Contract.SetOwner(&_YBribeV3.TransactOpts, _new_owner)
}

// SetRecipient is a paid mutator transaction binding the contract method 0xf21a1d11.
//
// Solidity: function set_recipient(address _recipient) returns()
func (_YBribeV3 *YBribeV3Transactor) SetRecipient(opts *bind.TransactOpts, _recipient common.Address) (*types.Transaction, error) {
	return _YBribeV3.contract.Transact(opts, "set_recipient", _recipient)
}

// SetRecipient is a paid mutator transaction binding the contract method 0xf21a1d11.
//
// Solidity: function set_recipient(address _recipient) returns()
func (_YBribeV3 *YBribeV3Session) SetRecipient(_recipient common.Address) (*types.Transaction, error) {
	return _YBribeV3.Contract.SetRecipient(&_YBribeV3.TransactOpts, _recipient)
}

// SetRecipient is a paid mutator transaction binding the contract method 0xf21a1d11.
//
// Solidity: function set_recipient(address _recipient) returns()
func (_YBribeV3 *YBribeV3TransactorSession) SetRecipient(_recipient common.Address) (*types.Transaction, error) {
	return _YBribeV3.Contract.SetRecipient(&_YBribeV3.TransactOpts, _recipient)
}

// YBribeV3BlacklistedIterator is returned from FilterBlacklisted and is used to iterate over the raw logs and unpacked data for Blacklisted events raised by the YBribeV3 contract.
type YBribeV3BlacklistedIterator struct {
	Event *YBribeV3Blacklisted // Event containing the contract specifics and raw log

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
func (it *YBribeV3BlacklistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBribeV3Blacklisted)
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
		it.Event = new(YBribeV3Blacklisted)
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
func (it *YBribeV3BlacklistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBribeV3BlacklistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBribeV3Blacklisted represents a Blacklisted event raised by the YBribeV3 contract.
type YBribeV3Blacklisted struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterBlacklisted is a free log retrieval operation binding the contract event 0xffa4e6181777692565cf28528fc88fd1516ea86b56da075235fa575af6a4b855.
//
// Solidity: event Blacklisted(address indexed user)
func (_YBribeV3 *YBribeV3Filterer) FilterBlacklisted(opts *bind.FilterOpts, user []common.Address) (*YBribeV3BlacklistedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _YBribeV3.contract.FilterLogs(opts, "Blacklisted", userRule)
	if err != nil {
		return nil, err
	}
	return &YBribeV3BlacklistedIterator{contract: _YBribeV3.contract, event: "Blacklisted", logs: logs, sub: sub}, nil
}

// WatchBlacklisted is a free log subscription operation binding the contract event 0xffa4e6181777692565cf28528fc88fd1516ea86b56da075235fa575af6a4b855.
//
// Solidity: event Blacklisted(address indexed user)
func (_YBribeV3 *YBribeV3Filterer) WatchBlacklisted(opts *bind.WatchOpts, sink chan<- *YBribeV3Blacklisted, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _YBribeV3.contract.WatchLogs(opts, "Blacklisted", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBribeV3Blacklisted)
				if err := _YBribeV3.contract.UnpackLog(event, "Blacklisted", log); err != nil {
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

// ParseBlacklisted is a log parse operation binding the contract event 0xffa4e6181777692565cf28528fc88fd1516ea86b56da075235fa575af6a4b855.
//
// Solidity: event Blacklisted(address indexed user)
func (_YBribeV3 *YBribeV3Filterer) ParseBlacklisted(log types.Log) (*YBribeV3Blacklisted, error) {
	event := new(YBribeV3Blacklisted)
	if err := _YBribeV3.contract.UnpackLog(event, "Blacklisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBribeV3ChangeOwnerIterator is returned from FilterChangeOwner and is used to iterate over the raw logs and unpacked data for ChangeOwner events raised by the YBribeV3 contract.
type YBribeV3ChangeOwnerIterator struct {
	Event *YBribeV3ChangeOwner // Event containing the contract specifics and raw log

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
func (it *YBribeV3ChangeOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBribeV3ChangeOwner)
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
		it.Event = new(YBribeV3ChangeOwner)
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
func (it *YBribeV3ChangeOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBribeV3ChangeOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBribeV3ChangeOwner represents a ChangeOwner event raised by the YBribeV3 contract.
type YBribeV3ChangeOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterChangeOwner is a free log retrieval operation binding the contract event 0xf285329298fd841af46eb83bbe90d1ebe2951c975a65b19a02f965f842ee69c5.
//
// Solidity: event ChangeOwner(address owner)
func (_YBribeV3 *YBribeV3Filterer) FilterChangeOwner(opts *bind.FilterOpts) (*YBribeV3ChangeOwnerIterator, error) {

	logs, sub, err := _YBribeV3.contract.FilterLogs(opts, "ChangeOwner")
	if err != nil {
		return nil, err
	}
	return &YBribeV3ChangeOwnerIterator{contract: _YBribeV3.contract, event: "ChangeOwner", logs: logs, sub: sub}, nil
}

// WatchChangeOwner is a free log subscription operation binding the contract event 0xf285329298fd841af46eb83bbe90d1ebe2951c975a65b19a02f965f842ee69c5.
//
// Solidity: event ChangeOwner(address owner)
func (_YBribeV3 *YBribeV3Filterer) WatchChangeOwner(opts *bind.WatchOpts, sink chan<- *YBribeV3ChangeOwner) (event.Subscription, error) {

	logs, sub, err := _YBribeV3.contract.WatchLogs(opts, "ChangeOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBribeV3ChangeOwner)
				if err := _YBribeV3.contract.UnpackLog(event, "ChangeOwner", log); err != nil {
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

// ParseChangeOwner is a log parse operation binding the contract event 0xf285329298fd841af46eb83bbe90d1ebe2951c975a65b19a02f965f842ee69c5.
//
// Solidity: event ChangeOwner(address owner)
func (_YBribeV3 *YBribeV3Filterer) ParseChangeOwner(log types.Log) (*YBribeV3ChangeOwner, error) {
	event := new(YBribeV3ChangeOwner)
	if err := _YBribeV3.contract.UnpackLog(event, "ChangeOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBribeV3ClearRewardRecipientIterator is returned from FilterClearRewardRecipient and is used to iterate over the raw logs and unpacked data for ClearRewardRecipient events raised by the YBribeV3 contract.
type YBribeV3ClearRewardRecipientIterator struct {
	Event *YBribeV3ClearRewardRecipient // Event containing the contract specifics and raw log

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
func (it *YBribeV3ClearRewardRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBribeV3ClearRewardRecipient)
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
		it.Event = new(YBribeV3ClearRewardRecipient)
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
func (it *YBribeV3ClearRewardRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBribeV3ClearRewardRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBribeV3ClearRewardRecipient represents a ClearRewardRecipient event raised by the YBribeV3 contract.
type YBribeV3ClearRewardRecipient struct {
	User      common.Address
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClearRewardRecipient is a free log retrieval operation binding the contract event 0x058f076b8b7f4d71ee98cf4a857fed97696b4cdf8f46838077710800628321e5.
//
// Solidity: event ClearRewardRecipient(address indexed user, address recipient)
func (_YBribeV3 *YBribeV3Filterer) FilterClearRewardRecipient(opts *bind.FilterOpts, user []common.Address) (*YBribeV3ClearRewardRecipientIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _YBribeV3.contract.FilterLogs(opts, "ClearRewardRecipient", userRule)
	if err != nil {
		return nil, err
	}
	return &YBribeV3ClearRewardRecipientIterator{contract: _YBribeV3.contract, event: "ClearRewardRecipient", logs: logs, sub: sub}, nil
}

// WatchClearRewardRecipient is a free log subscription operation binding the contract event 0x058f076b8b7f4d71ee98cf4a857fed97696b4cdf8f46838077710800628321e5.
//
// Solidity: event ClearRewardRecipient(address indexed user, address recipient)
func (_YBribeV3 *YBribeV3Filterer) WatchClearRewardRecipient(opts *bind.WatchOpts, sink chan<- *YBribeV3ClearRewardRecipient, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _YBribeV3.contract.WatchLogs(opts, "ClearRewardRecipient", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBribeV3ClearRewardRecipient)
				if err := _YBribeV3.contract.UnpackLog(event, "ClearRewardRecipient", log); err != nil {
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

// ParseClearRewardRecipient is a log parse operation binding the contract event 0x058f076b8b7f4d71ee98cf4a857fed97696b4cdf8f46838077710800628321e5.
//
// Solidity: event ClearRewardRecipient(address indexed user, address recipient)
func (_YBribeV3 *YBribeV3Filterer) ParseClearRewardRecipient(log types.Log) (*YBribeV3ClearRewardRecipient, error) {
	event := new(YBribeV3ClearRewardRecipient)
	if err := _YBribeV3.contract.UnpackLog(event, "ClearRewardRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBribeV3FeeUpdatedIterator is returned from FilterFeeUpdated and is used to iterate over the raw logs and unpacked data for FeeUpdated events raised by the YBribeV3 contract.
type YBribeV3FeeUpdatedIterator struct {
	Event *YBribeV3FeeUpdated // Event containing the contract specifics and raw log

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
func (it *YBribeV3FeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBribeV3FeeUpdated)
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
		it.Event = new(YBribeV3FeeUpdated)
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
func (it *YBribeV3FeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBribeV3FeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBribeV3FeeUpdated represents a FeeUpdated event raised by the YBribeV3 contract.
type YBribeV3FeeUpdated struct {
	Fee *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFeeUpdated is a free log retrieval operation binding the contract event 0x8c4d35e54a3f2ef1134138fd8ea3daee6a3c89e10d2665996babdf70261e2c76.
//
// Solidity: event FeeUpdated(uint256 fee)
func (_YBribeV3 *YBribeV3Filterer) FilterFeeUpdated(opts *bind.FilterOpts) (*YBribeV3FeeUpdatedIterator, error) {

	logs, sub, err := _YBribeV3.contract.FilterLogs(opts, "FeeUpdated")
	if err != nil {
		return nil, err
	}
	return &YBribeV3FeeUpdatedIterator{contract: _YBribeV3.contract, event: "FeeUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeUpdated is a free log subscription operation binding the contract event 0x8c4d35e54a3f2ef1134138fd8ea3daee6a3c89e10d2665996babdf70261e2c76.
//
// Solidity: event FeeUpdated(uint256 fee)
func (_YBribeV3 *YBribeV3Filterer) WatchFeeUpdated(opts *bind.WatchOpts, sink chan<- *YBribeV3FeeUpdated) (event.Subscription, error) {

	logs, sub, err := _YBribeV3.contract.WatchLogs(opts, "FeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBribeV3FeeUpdated)
				if err := _YBribeV3.contract.UnpackLog(event, "FeeUpdated", log); err != nil {
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

// ParseFeeUpdated is a log parse operation binding the contract event 0x8c4d35e54a3f2ef1134138fd8ea3daee6a3c89e10d2665996babdf70261e2c76.
//
// Solidity: event FeeUpdated(uint256 fee)
func (_YBribeV3 *YBribeV3Filterer) ParseFeeUpdated(log types.Log) (*YBribeV3FeeUpdated, error) {
	event := new(YBribeV3FeeUpdated)
	if err := _YBribeV3.contract.UnpackLog(event, "FeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBribeV3NewTokenRewardIterator is returned from FilterNewTokenReward and is used to iterate over the raw logs and unpacked data for NewTokenReward events raised by the YBribeV3 contract.
type YBribeV3NewTokenRewardIterator struct {
	Event *YBribeV3NewTokenReward // Event containing the contract specifics and raw log

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
func (it *YBribeV3NewTokenRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBribeV3NewTokenReward)
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
		it.Event = new(YBribeV3NewTokenReward)
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
func (it *YBribeV3NewTokenRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBribeV3NewTokenRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBribeV3NewTokenReward represents a NewTokenReward event raised by the YBribeV3 contract.
type YBribeV3NewTokenReward struct {
	Gauge       common.Address
	RewardToken common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewTokenReward is a free log retrieval operation binding the contract event 0xc3da16a0185e98874c12feef929e38c56e0abe7e4d6b62fca1c9b9ad2edd9f69.
//
// Solidity: event NewTokenReward(address indexed gauge, address indexed reward_token)
func (_YBribeV3 *YBribeV3Filterer) FilterNewTokenReward(opts *bind.FilterOpts, gauge []common.Address, reward_token []common.Address) (*YBribeV3NewTokenRewardIterator, error) {

	var gaugeRule []interface{}
	for _, gaugeItem := range gauge {
		gaugeRule = append(gaugeRule, gaugeItem)
	}
	var reward_tokenRule []interface{}
	for _, reward_tokenItem := range reward_token {
		reward_tokenRule = append(reward_tokenRule, reward_tokenItem)
	}

	logs, sub, err := _YBribeV3.contract.FilterLogs(opts, "NewTokenReward", gaugeRule, reward_tokenRule)
	if err != nil {
		return nil, err
	}
	return &YBribeV3NewTokenRewardIterator{contract: _YBribeV3.contract, event: "NewTokenReward", logs: logs, sub: sub}, nil
}

// WatchNewTokenReward is a free log subscription operation binding the contract event 0xc3da16a0185e98874c12feef929e38c56e0abe7e4d6b62fca1c9b9ad2edd9f69.
//
// Solidity: event NewTokenReward(address indexed gauge, address indexed reward_token)
func (_YBribeV3 *YBribeV3Filterer) WatchNewTokenReward(opts *bind.WatchOpts, sink chan<- *YBribeV3NewTokenReward, gauge []common.Address, reward_token []common.Address) (event.Subscription, error) {

	var gaugeRule []interface{}
	for _, gaugeItem := range gauge {
		gaugeRule = append(gaugeRule, gaugeItem)
	}
	var reward_tokenRule []interface{}
	for _, reward_tokenItem := range reward_token {
		reward_tokenRule = append(reward_tokenRule, reward_tokenItem)
	}

	logs, sub, err := _YBribeV3.contract.WatchLogs(opts, "NewTokenReward", gaugeRule, reward_tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBribeV3NewTokenReward)
				if err := _YBribeV3.contract.UnpackLog(event, "NewTokenReward", log); err != nil {
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

// ParseNewTokenReward is a log parse operation binding the contract event 0xc3da16a0185e98874c12feef929e38c56e0abe7e4d6b62fca1c9b9ad2edd9f69.
//
// Solidity: event NewTokenReward(address indexed gauge, address indexed reward_token)
func (_YBribeV3 *YBribeV3Filterer) ParseNewTokenReward(log types.Log) (*YBribeV3NewTokenReward, error) {
	event := new(YBribeV3NewTokenReward)
	if err := _YBribeV3.contract.UnpackLog(event, "NewTokenReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBribeV3PeriodUpdatedIterator is returned from FilterPeriodUpdated and is used to iterate over the raw logs and unpacked data for PeriodUpdated events raised by the YBribeV3 contract.
type YBribeV3PeriodUpdatedIterator struct {
	Event *YBribeV3PeriodUpdated // Event containing the contract specifics and raw log

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
func (it *YBribeV3PeriodUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBribeV3PeriodUpdated)
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
		it.Event = new(YBribeV3PeriodUpdated)
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
func (it *YBribeV3PeriodUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBribeV3PeriodUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBribeV3PeriodUpdated represents a PeriodUpdated event raised by the YBribeV3 contract.
type YBribeV3PeriodUpdated struct {
	Gauge           common.Address
	Period          *big.Int
	Bias            *big.Int
	BlacklistedBias *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPeriodUpdated is a free log retrieval operation binding the contract event 0xbd8342205f3471f235c1fc5537dec6e898313fe7c31c1ab75ecaa09feb7580b3.
//
// Solidity: event PeriodUpdated(address indexed gauge, uint256 indexed period, uint256 bias, uint256 blacklisted_bias)
func (_YBribeV3 *YBribeV3Filterer) FilterPeriodUpdated(opts *bind.FilterOpts, gauge []common.Address, period []*big.Int) (*YBribeV3PeriodUpdatedIterator, error) {

	var gaugeRule []interface{}
	for _, gaugeItem := range gauge {
		gaugeRule = append(gaugeRule, gaugeItem)
	}
	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _YBribeV3.contract.FilterLogs(opts, "PeriodUpdated", gaugeRule, periodRule)
	if err != nil {
		return nil, err
	}
	return &YBribeV3PeriodUpdatedIterator{contract: _YBribeV3.contract, event: "PeriodUpdated", logs: logs, sub: sub}, nil
}

// WatchPeriodUpdated is a free log subscription operation binding the contract event 0xbd8342205f3471f235c1fc5537dec6e898313fe7c31c1ab75ecaa09feb7580b3.
//
// Solidity: event PeriodUpdated(address indexed gauge, uint256 indexed period, uint256 bias, uint256 blacklisted_bias)
func (_YBribeV3 *YBribeV3Filterer) WatchPeriodUpdated(opts *bind.WatchOpts, sink chan<- *YBribeV3PeriodUpdated, gauge []common.Address, period []*big.Int) (event.Subscription, error) {

	var gaugeRule []interface{}
	for _, gaugeItem := range gauge {
		gaugeRule = append(gaugeRule, gaugeItem)
	}
	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _YBribeV3.contract.WatchLogs(opts, "PeriodUpdated", gaugeRule, periodRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBribeV3PeriodUpdated)
				if err := _YBribeV3.contract.UnpackLog(event, "PeriodUpdated", log); err != nil {
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

// ParsePeriodUpdated is a log parse operation binding the contract event 0xbd8342205f3471f235c1fc5537dec6e898313fe7c31c1ab75ecaa09feb7580b3.
//
// Solidity: event PeriodUpdated(address indexed gauge, uint256 indexed period, uint256 bias, uint256 blacklisted_bias)
func (_YBribeV3 *YBribeV3Filterer) ParsePeriodUpdated(log types.Log) (*YBribeV3PeriodUpdated, error) {
	event := new(YBribeV3PeriodUpdated)
	if err := _YBribeV3.contract.UnpackLog(event, "PeriodUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBribeV3RemovedFromBlacklistIterator is returned from FilterRemovedFromBlacklist and is used to iterate over the raw logs and unpacked data for RemovedFromBlacklist events raised by the YBribeV3 contract.
type YBribeV3RemovedFromBlacklistIterator struct {
	Event *YBribeV3RemovedFromBlacklist // Event containing the contract specifics and raw log

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
func (it *YBribeV3RemovedFromBlacklistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBribeV3RemovedFromBlacklist)
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
		it.Event = new(YBribeV3RemovedFromBlacklist)
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
func (it *YBribeV3RemovedFromBlacklistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBribeV3RemovedFromBlacklistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBribeV3RemovedFromBlacklist represents a RemovedFromBlacklist event raised by the YBribeV3 contract.
type YBribeV3RemovedFromBlacklist struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRemovedFromBlacklist is a free log retrieval operation binding the contract event 0x2b6bf71b58b3583add364b3d9060ebf8019650f65f5be35f5464b9cb3e4ba2d4.
//
// Solidity: event RemovedFromBlacklist(address indexed user)
func (_YBribeV3 *YBribeV3Filterer) FilterRemovedFromBlacklist(opts *bind.FilterOpts, user []common.Address) (*YBribeV3RemovedFromBlacklistIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _YBribeV3.contract.FilterLogs(opts, "RemovedFromBlacklist", userRule)
	if err != nil {
		return nil, err
	}
	return &YBribeV3RemovedFromBlacklistIterator{contract: _YBribeV3.contract, event: "RemovedFromBlacklist", logs: logs, sub: sub}, nil
}

// WatchRemovedFromBlacklist is a free log subscription operation binding the contract event 0x2b6bf71b58b3583add364b3d9060ebf8019650f65f5be35f5464b9cb3e4ba2d4.
//
// Solidity: event RemovedFromBlacklist(address indexed user)
func (_YBribeV3 *YBribeV3Filterer) WatchRemovedFromBlacklist(opts *bind.WatchOpts, sink chan<- *YBribeV3RemovedFromBlacklist, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _YBribeV3.contract.WatchLogs(opts, "RemovedFromBlacklist", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBribeV3RemovedFromBlacklist)
				if err := _YBribeV3.contract.UnpackLog(event, "RemovedFromBlacklist", log); err != nil {
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

// ParseRemovedFromBlacklist is a log parse operation binding the contract event 0x2b6bf71b58b3583add364b3d9060ebf8019650f65f5be35f5464b9cb3e4ba2d4.
//
// Solidity: event RemovedFromBlacklist(address indexed user)
func (_YBribeV3 *YBribeV3Filterer) ParseRemovedFromBlacklist(log types.Log) (*YBribeV3RemovedFromBlacklist, error) {
	event := new(YBribeV3RemovedFromBlacklist)
	if err := _YBribeV3.contract.UnpackLog(event, "RemovedFromBlacklist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBribeV3RewardAddedIterator is returned from FilterRewardAdded and is used to iterate over the raw logs and unpacked data for RewardAdded events raised by the YBribeV3 contract.
type YBribeV3RewardAddedIterator struct {
	Event *YBribeV3RewardAdded // Event containing the contract specifics and raw log

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
func (it *YBribeV3RewardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBribeV3RewardAdded)
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
		it.Event = new(YBribeV3RewardAdded)
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
func (it *YBribeV3RewardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBribeV3RewardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBribeV3RewardAdded represents a RewardAdded event raised by the YBribeV3 contract.
type YBribeV3RewardAdded struct {
	Briber      common.Address
	Gauge       common.Address
	RewardToken common.Address
	Amount      *big.Int
	Fee         *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardAdded is a free log retrieval operation binding the contract event 0xbfdc09d345164585641da75a2ba1b1adc60bf86cdf108cc974ec9f4776d7f787.
//
// Solidity: event RewardAdded(address indexed briber, address indexed gauge, address indexed reward_token, uint256 amount, uint256 fee)
func (_YBribeV3 *YBribeV3Filterer) FilterRewardAdded(opts *bind.FilterOpts, briber []common.Address, gauge []common.Address, reward_token []common.Address) (*YBribeV3RewardAddedIterator, error) {

	var briberRule []interface{}
	for _, briberItem := range briber {
		briberRule = append(briberRule, briberItem)
	}
	var gaugeRule []interface{}
	for _, gaugeItem := range gauge {
		gaugeRule = append(gaugeRule, gaugeItem)
	}
	var reward_tokenRule []interface{}
	for _, reward_tokenItem := range reward_token {
		reward_tokenRule = append(reward_tokenRule, reward_tokenItem)
	}

	logs, sub, err := _YBribeV3.contract.FilterLogs(opts, "RewardAdded", briberRule, gaugeRule, reward_tokenRule)
	if err != nil {
		return nil, err
	}
	return &YBribeV3RewardAddedIterator{contract: _YBribeV3.contract, event: "RewardAdded", logs: logs, sub: sub}, nil
}

// WatchRewardAdded is a free log subscription operation binding the contract event 0xbfdc09d345164585641da75a2ba1b1adc60bf86cdf108cc974ec9f4776d7f787.
//
// Solidity: event RewardAdded(address indexed briber, address indexed gauge, address indexed reward_token, uint256 amount, uint256 fee)
func (_YBribeV3 *YBribeV3Filterer) WatchRewardAdded(opts *bind.WatchOpts, sink chan<- *YBribeV3RewardAdded, briber []common.Address, gauge []common.Address, reward_token []common.Address) (event.Subscription, error) {

	var briberRule []interface{}
	for _, briberItem := range briber {
		briberRule = append(briberRule, briberItem)
	}
	var gaugeRule []interface{}
	for _, gaugeItem := range gauge {
		gaugeRule = append(gaugeRule, gaugeItem)
	}
	var reward_tokenRule []interface{}
	for _, reward_tokenItem := range reward_token {
		reward_tokenRule = append(reward_tokenRule, reward_tokenItem)
	}

	logs, sub, err := _YBribeV3.contract.WatchLogs(opts, "RewardAdded", briberRule, gaugeRule, reward_tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBribeV3RewardAdded)
				if err := _YBribeV3.contract.UnpackLog(event, "RewardAdded", log); err != nil {
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

// ParseRewardAdded is a log parse operation binding the contract event 0xbfdc09d345164585641da75a2ba1b1adc60bf86cdf108cc974ec9f4776d7f787.
//
// Solidity: event RewardAdded(address indexed briber, address indexed gauge, address indexed reward_token, uint256 amount, uint256 fee)
func (_YBribeV3 *YBribeV3Filterer) ParseRewardAdded(log types.Log) (*YBribeV3RewardAdded, error) {
	event := new(YBribeV3RewardAdded)
	if err := _YBribeV3.contract.UnpackLog(event, "RewardAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBribeV3RewardClaimedIterator is returned from FilterRewardClaimed and is used to iterate over the raw logs and unpacked data for RewardClaimed events raised by the YBribeV3 contract.
type YBribeV3RewardClaimedIterator struct {
	Event *YBribeV3RewardClaimed // Event containing the contract specifics and raw log

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
func (it *YBribeV3RewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBribeV3RewardClaimed)
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
		it.Event = new(YBribeV3RewardClaimed)
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
func (it *YBribeV3RewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBribeV3RewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBribeV3RewardClaimed represents a RewardClaimed event raised by the YBribeV3 contract.
type YBribeV3RewardClaimed struct {
	User        common.Address
	Gauge       common.Address
	RewardToken common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardClaimed is a free log retrieval operation binding the contract event 0x2422cac5e23c46c890fdcf42d0c64757409df6832174df639337558f09d99c68.
//
// Solidity: event RewardClaimed(address indexed user, address indexed gauge, address indexed reward_token, uint256 amount)
func (_YBribeV3 *YBribeV3Filterer) FilterRewardClaimed(opts *bind.FilterOpts, user []common.Address, gauge []common.Address, reward_token []common.Address) (*YBribeV3RewardClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var gaugeRule []interface{}
	for _, gaugeItem := range gauge {
		gaugeRule = append(gaugeRule, gaugeItem)
	}
	var reward_tokenRule []interface{}
	for _, reward_tokenItem := range reward_token {
		reward_tokenRule = append(reward_tokenRule, reward_tokenItem)
	}

	logs, sub, err := _YBribeV3.contract.FilterLogs(opts, "RewardClaimed", userRule, gaugeRule, reward_tokenRule)
	if err != nil {
		return nil, err
	}
	return &YBribeV3RewardClaimedIterator{contract: _YBribeV3.contract, event: "RewardClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardClaimed is a free log subscription operation binding the contract event 0x2422cac5e23c46c890fdcf42d0c64757409df6832174df639337558f09d99c68.
//
// Solidity: event RewardClaimed(address indexed user, address indexed gauge, address indexed reward_token, uint256 amount)
func (_YBribeV3 *YBribeV3Filterer) WatchRewardClaimed(opts *bind.WatchOpts, sink chan<- *YBribeV3RewardClaimed, user []common.Address, gauge []common.Address, reward_token []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var gaugeRule []interface{}
	for _, gaugeItem := range gauge {
		gaugeRule = append(gaugeRule, gaugeItem)
	}
	var reward_tokenRule []interface{}
	for _, reward_tokenItem := range reward_token {
		reward_tokenRule = append(reward_tokenRule, reward_tokenItem)
	}

	logs, sub, err := _YBribeV3.contract.WatchLogs(opts, "RewardClaimed", userRule, gaugeRule, reward_tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBribeV3RewardClaimed)
				if err := _YBribeV3.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
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

// ParseRewardClaimed is a log parse operation binding the contract event 0x2422cac5e23c46c890fdcf42d0c64757409df6832174df639337558f09d99c68.
//
// Solidity: event RewardClaimed(address indexed user, address indexed gauge, address indexed reward_token, uint256 amount)
func (_YBribeV3 *YBribeV3Filterer) ParseRewardClaimed(log types.Log) (*YBribeV3RewardClaimed, error) {
	event := new(YBribeV3RewardClaimed)
	if err := _YBribeV3.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBribeV3SetRewardRecipientIterator is returned from FilterSetRewardRecipient and is used to iterate over the raw logs and unpacked data for SetRewardRecipient events raised by the YBribeV3 contract.
type YBribeV3SetRewardRecipientIterator struct {
	Event *YBribeV3SetRewardRecipient // Event containing the contract specifics and raw log

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
func (it *YBribeV3SetRewardRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBribeV3SetRewardRecipient)
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
		it.Event = new(YBribeV3SetRewardRecipient)
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
func (it *YBribeV3SetRewardRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBribeV3SetRewardRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBribeV3SetRewardRecipient represents a SetRewardRecipient event raised by the YBribeV3 contract.
type YBribeV3SetRewardRecipient struct {
	User      common.Address
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetRewardRecipient is a free log retrieval operation binding the contract event 0xc6b66e0e282673c442421e1c6b89458b7631f26f5dcd0b2b216c45831ca1d7d5.
//
// Solidity: event SetRewardRecipient(address indexed user, address recipient)
func (_YBribeV3 *YBribeV3Filterer) FilterSetRewardRecipient(opts *bind.FilterOpts, user []common.Address) (*YBribeV3SetRewardRecipientIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _YBribeV3.contract.FilterLogs(opts, "SetRewardRecipient", userRule)
	if err != nil {
		return nil, err
	}
	return &YBribeV3SetRewardRecipientIterator{contract: _YBribeV3.contract, event: "SetRewardRecipient", logs: logs, sub: sub}, nil
}

// WatchSetRewardRecipient is a free log subscription operation binding the contract event 0xc6b66e0e282673c442421e1c6b89458b7631f26f5dcd0b2b216c45831ca1d7d5.
//
// Solidity: event SetRewardRecipient(address indexed user, address recipient)
func (_YBribeV3 *YBribeV3Filterer) WatchSetRewardRecipient(opts *bind.WatchOpts, sink chan<- *YBribeV3SetRewardRecipient, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _YBribeV3.contract.WatchLogs(opts, "SetRewardRecipient", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBribeV3SetRewardRecipient)
				if err := _YBribeV3.contract.UnpackLog(event, "SetRewardRecipient", log); err != nil {
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

// ParseSetRewardRecipient is a log parse operation binding the contract event 0xc6b66e0e282673c442421e1c6b89458b7631f26f5dcd0b2b216c45831ca1d7d5.
//
// Solidity: event SetRewardRecipient(address indexed user, address recipient)
func (_YBribeV3 *YBribeV3Filterer) ParseSetRewardRecipient(log types.Log) (*YBribeV3SetRewardRecipient, error) {
	event := new(YBribeV3SetRewardRecipient)
	if err := _YBribeV3.contract.UnpackLog(event, "SetRewardRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
