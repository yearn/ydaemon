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

// CurveGaugeMetaData contains all meta data concerning the CurveGauge contract.
var CurveGaugeMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"Deposit\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateLiquidityLimit\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":false},{\"name\":\"original_balance\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"original_supply\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"working_balance\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"working_supply\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CommitOwnership\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ApplyOwnership\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Transfer\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\",\"indexed\":true},{\"name\":\"_to\",\"type\":\"address\",\"indexed\":true},{\"name\":\"_value\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Approval\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"indexed\":true},{\"name\":\"_spender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"_value\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"_lp_token\",\"type\":\"address\"},{\"name\":\"_admin\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":288},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"integrate_checkpoint\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":4560},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"user_checkpoint\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":3123352},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"claimable_tokens\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3038594},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"claimed_reward\",\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_token\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3006},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"claimable_reward\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_reward_token\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":20225},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_rewards_receiver\",\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[],\"gas\":35643},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"claim_rewards\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"claim_rewards\",\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"claim_rewards\",\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"kick\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"outputs\":[],\"gas\":3137443},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_addr\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_claim_rewards\",\"type\":\"bool\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_claim_rewards\",\"type\":\"bool\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":18062446},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":18100396},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":39421},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"increaseAllowance\",\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_added_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":41965},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"decreaseAllowance\",\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtracted_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":41989},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_reward\",\"inputs\":[{\"name\":\"_reward_token\",\"type\":\"address\"},{\"name\":\"_distributor\",\"type\":\"address\"}],\"outputs\":[],\"gas\":113003},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_reward_distributor\",\"inputs\":[{\"name\":\"_reward_token\",\"type\":\"address\"},{\"name\":\"_distributor\",\"type\":\"address\"}],\"outputs\":[],\"gas\":40753},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deposit_reward_token\",\"inputs\":[{\"name\":\"_reward_token\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"outputs\":[],\"gas\":1540169},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_killed\",\"inputs\":[{\"name\":\"_is_killed\",\"type\":\"bool\"}],\"outputs\":[],\"gas\":38115},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"commit_transfer_ownership\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"outputs\":[],\"gas\":40045},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"accept_transfer_ownership\",\"inputs\":[],\"outputs\":[],\"gas\":39990},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"lp_token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3048},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_epoch_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3078},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3323},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3138},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"},{\"name\":\"arg1\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3598},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"gas\":13428},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"gas\":11181},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"working_balances\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3473},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"working_supply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3288},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"period\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"int128\"}],\"gas\":3318},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"period_timestamp\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3393},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"integrate_inv_supply\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3423},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"integrate_inv_supply_of\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3623},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"integrate_checkpoint_of\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3653},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"integrate_fraction\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3683},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"inflation_rate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3498},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"reward_count\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3528},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"reward_tokens\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3603},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"reward_data\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"distributor\",\"type\":\"address\"},{\"name\":\"period_finish\",\"type\":\"uint256\"},{\"name\":\"rate\",\"type\":\"uint256\"},{\"name\":\"last_update\",\"type\":\"uint256\"},{\"name\":\"integral\",\"type\":\"uint256\"}],\"gas\":15033},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"rewards_receiver\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3833},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"reward_integral_for\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"},{\"name\":\"arg1\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":4078},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3678},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3708},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"is_killed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":3738}]",
}

// CurveGaugeABI is the input ABI used to generate the binding from.
// Deprecated: Use CurveGaugeMetaData.ABI instead.
var CurveGaugeABI = CurveGaugeMetaData.ABI

// CurveGauge is an auto generated Go binding around an Ethereum contract.
type CurveGauge struct {
	CurveGaugeCaller     // Read-only binding to the contract
	CurveGaugeTransactor // Write-only binding to the contract
	CurveGaugeFilterer   // Log filterer for contract events
}

// CurveGaugeCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurveGaugeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveGaugeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurveGaugeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveGaugeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurveGaugeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveGaugeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurveGaugeSession struct {
	Contract     *CurveGauge       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CurveGaugeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurveGaugeCallerSession struct {
	Contract *CurveGaugeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CurveGaugeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurveGaugeTransactorSession struct {
	Contract     *CurveGaugeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CurveGaugeRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurveGaugeRaw struct {
	Contract *CurveGauge // Generic contract binding to access the raw methods on
}

// CurveGaugeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurveGaugeCallerRaw struct {
	Contract *CurveGaugeCaller // Generic read-only contract binding to access the raw methods on
}

// CurveGaugeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurveGaugeTransactorRaw struct {
	Contract *CurveGaugeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurveGauge creates a new instance of CurveGauge, bound to a specific deployed contract.
func NewCurveGauge(address common.Address, backend bind.ContractBackend) (*CurveGauge, error) {
	contract, err := bindCurveGauge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CurveGauge{CurveGaugeCaller: CurveGaugeCaller{contract: contract}, CurveGaugeTransactor: CurveGaugeTransactor{contract: contract}, CurveGaugeFilterer: CurveGaugeFilterer{contract: contract}}, nil
}

// NewCurveGaugeCaller creates a new read-only instance of CurveGauge, bound to a specific deployed contract.
func NewCurveGaugeCaller(address common.Address, caller bind.ContractCaller) (*CurveGaugeCaller, error) {
	contract, err := bindCurveGauge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurveGaugeCaller{contract: contract}, nil
}

// NewCurveGaugeTransactor creates a new write-only instance of CurveGauge, bound to a specific deployed contract.
func NewCurveGaugeTransactor(address common.Address, transactor bind.ContractTransactor) (*CurveGaugeTransactor, error) {
	contract, err := bindCurveGauge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurveGaugeTransactor{contract: contract}, nil
}

// NewCurveGaugeFilterer creates a new log filterer instance of CurveGauge, bound to a specific deployed contract.
func NewCurveGaugeFilterer(address common.Address, filterer bind.ContractFilterer) (*CurveGaugeFilterer, error) {
	contract, err := bindCurveGauge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurveGaugeFilterer{contract: contract}, nil
}

// bindCurveGauge binds a generic wrapper to an already deployed contract.
func bindCurveGauge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CurveGaugeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveGauge *CurveGaugeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveGauge.Contract.CurveGaugeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveGauge *CurveGaugeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveGauge.Contract.CurveGaugeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveGauge *CurveGaugeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveGauge.Contract.CurveGaugeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveGauge *CurveGaugeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveGauge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveGauge *CurveGaugeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveGauge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveGauge *CurveGaugeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveGauge.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveGauge *CurveGaugeCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveGauge *CurveGaugeSession) Admin() (common.Address, error) {
	return _CurveGauge.Contract.Admin(&_CurveGauge.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveGauge *CurveGaugeCallerSession) Admin() (common.Address, error) {
	return _CurveGauge.Contract.Admin(&_CurveGauge.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_CurveGauge *CurveGaugeCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_CurveGauge *CurveGaugeSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _CurveGauge.Contract.Allowance(&_CurveGauge.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_CurveGauge *CurveGaugeCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _CurveGauge.Contract.Allowance(&_CurveGauge.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_CurveGauge *CurveGaugeCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_CurveGauge *CurveGaugeSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _CurveGauge.Contract.BalanceOf(&_CurveGauge.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_CurveGauge *CurveGaugeCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _CurveGauge.Contract.BalanceOf(&_CurveGauge.CallOpts, arg0)
}

// ClaimableReward is a free data retrieval call binding the contract method 0x33fd6f74.
//
// Solidity: function claimable_reward(address _user, address _reward_token) view returns(uint256)
func (_CurveGauge *CurveGaugeCaller) ClaimableReward(opts *bind.CallOpts, _user common.Address, _reward_token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "claimable_reward", _user, _reward_token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimableReward is a free data retrieval call binding the contract method 0x33fd6f74.
//
// Solidity: function claimable_reward(address _user, address _reward_token) view returns(uint256)
func (_CurveGauge *CurveGaugeSession) ClaimableReward(_user common.Address, _reward_token common.Address) (*big.Int, error) {
	return _CurveGauge.Contract.ClaimableReward(&_CurveGauge.CallOpts, _user, _reward_token)
}

// ClaimableReward is a free data retrieval call binding the contract method 0x33fd6f74.
//
// Solidity: function claimable_reward(address _user, address _reward_token) view returns(uint256)
func (_CurveGauge *CurveGaugeCallerSession) ClaimableReward(_user common.Address, _reward_token common.Address) (*big.Int, error) {
	return _CurveGauge.Contract.ClaimableReward(&_CurveGauge.CallOpts, _user, _reward_token)
}

// ClaimedReward is a free data retrieval call binding the contract method 0xe77e7437.
//
// Solidity: function claimed_reward(address _addr, address _token) view returns(uint256)
func (_CurveGauge *CurveGaugeCaller) ClaimedReward(opts *bind.CallOpts, _addr common.Address, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "claimed_reward", _addr, _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimedReward is a free data retrieval call binding the contract method 0xe77e7437.
//
// Solidity: function claimed_reward(address _addr, address _token) view returns(uint256)
func (_CurveGauge *CurveGaugeSession) ClaimedReward(_addr common.Address, _token common.Address) (*big.Int, error) {
	return _CurveGauge.Contract.ClaimedReward(&_CurveGauge.CallOpts, _addr, _token)
}

// ClaimedReward is a free data retrieval call binding the contract method 0xe77e7437.
//
// Solidity: function claimed_reward(address _addr, address _token) view returns(uint256)
func (_CurveGauge *CurveGaugeCallerSession) ClaimedReward(_addr common.Address, _token common.Address) (*big.Int, error) {
	return _CurveGauge.Contract.ClaimedReward(&_CurveGauge.CallOpts, _addr, _token)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CurveGauge *CurveGaugeCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CurveGauge *CurveGaugeSession) Decimals() (*big.Int, error) {
	return _CurveGauge.Contract.Decimals(&_CurveGauge.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CurveGauge *CurveGaugeCallerSession) Decimals() (*big.Int, error) {
	return _CurveGauge.Contract.Decimals(&_CurveGauge.CallOpts)
}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_CurveGauge *CurveGaugeCaller) FutureAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "future_admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_CurveGauge *CurveGaugeSession) FutureAdmin() (common.Address, error) {
	return _CurveGauge.Contract.FutureAdmin(&_CurveGauge.CallOpts)
}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_CurveGauge *CurveGaugeCallerSession) FutureAdmin() (common.Address, error) {
	return _CurveGauge.Contract.FutureAdmin(&_CurveGauge.CallOpts)
}

// FutureEpochTime is a free data retrieval call binding the contract method 0xbe5d1be9.
//
// Solidity: function future_epoch_time() view returns(uint256)
func (_CurveGauge *CurveGaugeCaller) FutureEpochTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "future_epoch_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureEpochTime is a free data retrieval call binding the contract method 0xbe5d1be9.
//
// Solidity: function future_epoch_time() view returns(uint256)
func (_CurveGauge *CurveGaugeSession) FutureEpochTime() (*big.Int, error) {
	return _CurveGauge.Contract.FutureEpochTime(&_CurveGauge.CallOpts)
}

// FutureEpochTime is a free data retrieval call binding the contract method 0xbe5d1be9.
//
// Solidity: function future_epoch_time() view returns(uint256)
func (_CurveGauge *CurveGaugeCallerSession) FutureEpochTime() (*big.Int, error) {
	return _CurveGauge.Contract.FutureEpochTime(&_CurveGauge.CallOpts)
}

// InflationRate is a free data retrieval call binding the contract method 0x180692d0.
//
// Solidity: function inflation_rate() view returns(uint256)
func (_CurveGauge *CurveGaugeCaller) InflationRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "inflation_rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InflationRate is a free data retrieval call binding the contract method 0x180692d0.
//
// Solidity: function inflation_rate() view returns(uint256)
func (_CurveGauge *CurveGaugeSession) InflationRate() (*big.Int, error) {
	return _CurveGauge.Contract.InflationRate(&_CurveGauge.CallOpts)
}

// InflationRate is a free data retrieval call binding the contract method 0x180692d0.
//
// Solidity: function inflation_rate() view returns(uint256)
func (_CurveGauge *CurveGaugeCallerSession) InflationRate() (*big.Int, error) {
	return _CurveGauge.Contract.InflationRate(&_CurveGauge.CallOpts)
}

// IntegrateCheckpoint is a free data retrieval call binding the contract method 0xd31f3f6d.
//
// Solidity: function integrate_checkpoint() view returns(uint256)
func (_CurveGauge *CurveGaugeCaller) IntegrateCheckpoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "integrate_checkpoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IntegrateCheckpoint is a free data retrieval call binding the contract method 0xd31f3f6d.
//
// Solidity: function integrate_checkpoint() view returns(uint256)
func (_CurveGauge *CurveGaugeSession) IntegrateCheckpoint() (*big.Int, error) {
	return _CurveGauge.Contract.IntegrateCheckpoint(&_CurveGauge.CallOpts)
}

// IntegrateCheckpoint is a free data retrieval call binding the contract method 0xd31f3f6d.
//
// Solidity: function integrate_checkpoint() view returns(uint256)
func (_CurveGauge *CurveGaugeCallerSession) IntegrateCheckpoint() (*big.Int, error) {
	return _CurveGauge.Contract.IntegrateCheckpoint(&_CurveGauge.CallOpts)
}

// IntegrateCheckpointOf is a free data retrieval call binding the contract method 0x9bd324f2.
//
// Solidity: function integrate_checkpoint_of(address arg0) view returns(uint256)
func (_CurveGauge *CurveGaugeCaller) IntegrateCheckpointOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "integrate_checkpoint_of", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IntegrateCheckpointOf is a free data retrieval call binding the contract method 0x9bd324f2.
//
// Solidity: function integrate_checkpoint_of(address arg0) view returns(uint256)
func (_CurveGauge *CurveGaugeSession) IntegrateCheckpointOf(arg0 common.Address) (*big.Int, error) {
	return _CurveGauge.Contract.IntegrateCheckpointOf(&_CurveGauge.CallOpts, arg0)
}

// IntegrateCheckpointOf is a free data retrieval call binding the contract method 0x9bd324f2.
//
// Solidity: function integrate_checkpoint_of(address arg0) view returns(uint256)
func (_CurveGauge *CurveGaugeCallerSession) IntegrateCheckpointOf(arg0 common.Address) (*big.Int, error) {
	return _CurveGauge.Contract.IntegrateCheckpointOf(&_CurveGauge.CallOpts, arg0)
}

// IntegrateFraction is a free data retrieval call binding the contract method 0x09400707.
//
// Solidity: function integrate_fraction(address arg0) view returns(uint256)
func (_CurveGauge *CurveGaugeCaller) IntegrateFraction(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "integrate_fraction", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IntegrateFraction is a free data retrieval call binding the contract method 0x09400707.
//
// Solidity: function integrate_fraction(address arg0) view returns(uint256)
func (_CurveGauge *CurveGaugeSession) IntegrateFraction(arg0 common.Address) (*big.Int, error) {
	return _CurveGauge.Contract.IntegrateFraction(&_CurveGauge.CallOpts, arg0)
}

// IntegrateFraction is a free data retrieval call binding the contract method 0x09400707.
//
// Solidity: function integrate_fraction(address arg0) view returns(uint256)
func (_CurveGauge *CurveGaugeCallerSession) IntegrateFraction(arg0 common.Address) (*big.Int, error) {
	return _CurveGauge.Contract.IntegrateFraction(&_CurveGauge.CallOpts, arg0)
}

// IntegrateInvSupply is a free data retrieval call binding the contract method 0xfec8ee0c.
//
// Solidity: function integrate_inv_supply(uint256 arg0) view returns(uint256)
func (_CurveGauge *CurveGaugeCaller) IntegrateInvSupply(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "integrate_inv_supply", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IntegrateInvSupply is a free data retrieval call binding the contract method 0xfec8ee0c.
//
// Solidity: function integrate_inv_supply(uint256 arg0) view returns(uint256)
func (_CurveGauge *CurveGaugeSession) IntegrateInvSupply(arg0 *big.Int) (*big.Int, error) {
	return _CurveGauge.Contract.IntegrateInvSupply(&_CurveGauge.CallOpts, arg0)
}

// IntegrateInvSupply is a free data retrieval call binding the contract method 0xfec8ee0c.
//
// Solidity: function integrate_inv_supply(uint256 arg0) view returns(uint256)
func (_CurveGauge *CurveGaugeCallerSession) IntegrateInvSupply(arg0 *big.Int) (*big.Int, error) {
	return _CurveGauge.Contract.IntegrateInvSupply(&_CurveGauge.CallOpts, arg0)
}

// IntegrateInvSupplyOf is a free data retrieval call binding the contract method 0xde263bfa.
//
// Solidity: function integrate_inv_supply_of(address arg0) view returns(uint256)
func (_CurveGauge *CurveGaugeCaller) IntegrateInvSupplyOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "integrate_inv_supply_of", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IntegrateInvSupplyOf is a free data retrieval call binding the contract method 0xde263bfa.
//
// Solidity: function integrate_inv_supply_of(address arg0) view returns(uint256)
func (_CurveGauge *CurveGaugeSession) IntegrateInvSupplyOf(arg0 common.Address) (*big.Int, error) {
	return _CurveGauge.Contract.IntegrateInvSupplyOf(&_CurveGauge.CallOpts, arg0)
}

// IntegrateInvSupplyOf is a free data retrieval call binding the contract method 0xde263bfa.
//
// Solidity: function integrate_inv_supply_of(address arg0) view returns(uint256)
func (_CurveGauge *CurveGaugeCallerSession) IntegrateInvSupplyOf(arg0 common.Address) (*big.Int, error) {
	return _CurveGauge.Contract.IntegrateInvSupplyOf(&_CurveGauge.CallOpts, arg0)
}

// IsKilled is a free data retrieval call binding the contract method 0x9c868ac0.
//
// Solidity: function is_killed() view returns(bool)
func (_CurveGauge *CurveGaugeCaller) IsKilled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "is_killed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsKilled is a free data retrieval call binding the contract method 0x9c868ac0.
//
// Solidity: function is_killed() view returns(bool)
func (_CurveGauge *CurveGaugeSession) IsKilled() (bool, error) {
	return _CurveGauge.Contract.IsKilled(&_CurveGauge.CallOpts)
}

// IsKilled is a free data retrieval call binding the contract method 0x9c868ac0.
//
// Solidity: function is_killed() view returns(bool)
func (_CurveGauge *CurveGaugeCallerSession) IsKilled() (bool, error) {
	return _CurveGauge.Contract.IsKilled(&_CurveGauge.CallOpts)
}

// LpToken is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_CurveGauge *CurveGaugeCaller) LpToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "lp_token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LpToken is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_CurveGauge *CurveGaugeSession) LpToken() (common.Address, error) {
	return _CurveGauge.Contract.LpToken(&_CurveGauge.CallOpts)
}

// LpToken is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_CurveGauge *CurveGaugeCallerSession) LpToken() (common.Address, error) {
	return _CurveGauge.Contract.LpToken(&_CurveGauge.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurveGauge *CurveGaugeCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurveGauge *CurveGaugeSession) Name() (string, error) {
	return _CurveGauge.Contract.Name(&_CurveGauge.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurveGauge *CurveGaugeCallerSession) Name() (string, error) {
	return _CurveGauge.Contract.Name(&_CurveGauge.CallOpts)
}

// Period is a free data retrieval call binding the contract method 0xef78d4fd.
//
// Solidity: function period() view returns(int128)
func (_CurveGauge *CurveGaugeCaller) Period(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "period")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Period is a free data retrieval call binding the contract method 0xef78d4fd.
//
// Solidity: function period() view returns(int128)
func (_CurveGauge *CurveGaugeSession) Period() (*big.Int, error) {
	return _CurveGauge.Contract.Period(&_CurveGauge.CallOpts)
}

// Period is a free data retrieval call binding the contract method 0xef78d4fd.
//
// Solidity: function period() view returns(int128)
func (_CurveGauge *CurveGaugeCallerSession) Period() (*big.Int, error) {
	return _CurveGauge.Contract.Period(&_CurveGauge.CallOpts)
}

// PeriodTimestamp is a free data retrieval call binding the contract method 0x7598108c.
//
// Solidity: function period_timestamp(uint256 arg0) view returns(uint256)
func (_CurveGauge *CurveGaugeCaller) PeriodTimestamp(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "period_timestamp", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeriodTimestamp is a free data retrieval call binding the contract method 0x7598108c.
//
// Solidity: function period_timestamp(uint256 arg0) view returns(uint256)
func (_CurveGauge *CurveGaugeSession) PeriodTimestamp(arg0 *big.Int) (*big.Int, error) {
	return _CurveGauge.Contract.PeriodTimestamp(&_CurveGauge.CallOpts, arg0)
}

// PeriodTimestamp is a free data retrieval call binding the contract method 0x7598108c.
//
// Solidity: function period_timestamp(uint256 arg0) view returns(uint256)
func (_CurveGauge *CurveGaugeCallerSession) PeriodTimestamp(arg0 *big.Int) (*big.Int, error) {
	return _CurveGauge.Contract.PeriodTimestamp(&_CurveGauge.CallOpts, arg0)
}

// RewardCount is a free data retrieval call binding the contract method 0x963c94b9.
//
// Solidity: function reward_count() view returns(uint256)
func (_CurveGauge *CurveGaugeCaller) RewardCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "reward_count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardCount is a free data retrieval call binding the contract method 0x963c94b9.
//
// Solidity: function reward_count() view returns(uint256)
func (_CurveGauge *CurveGaugeSession) RewardCount() (*big.Int, error) {
	return _CurveGauge.Contract.RewardCount(&_CurveGauge.CallOpts)
}

// RewardCount is a free data retrieval call binding the contract method 0x963c94b9.
//
// Solidity: function reward_count() view returns(uint256)
func (_CurveGauge *CurveGaugeCallerSession) RewardCount() (*big.Int, error) {
	return _CurveGauge.Contract.RewardCount(&_CurveGauge.CallOpts)
}

// RewardData is a free data retrieval call binding the contract method 0x48e9c65e.
//
// Solidity: function reward_data(address arg0) view returns(address token, address distributor, uint256 period_finish, uint256 rate, uint256 last_update, uint256 integral)
func (_CurveGauge *CurveGaugeCaller) RewardData(opts *bind.CallOpts, arg0 common.Address) (struct {
	Token        common.Address
	Distributor  common.Address
	PeriodFinish *big.Int
	Rate         *big.Int
	LastUpdate   *big.Int
	Integral     *big.Int
}, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "reward_data", arg0)

	outstruct := new(struct {
		Token        common.Address
		Distributor  common.Address
		PeriodFinish *big.Int
		Rate         *big.Int
		LastUpdate   *big.Int
		Integral     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Distributor = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.PeriodFinish = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Rate = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LastUpdate = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Integral = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RewardData is a free data retrieval call binding the contract method 0x48e9c65e.
//
// Solidity: function reward_data(address arg0) view returns(address token, address distributor, uint256 period_finish, uint256 rate, uint256 last_update, uint256 integral)
func (_CurveGauge *CurveGaugeSession) RewardData(arg0 common.Address) (struct {
	Token        common.Address
	Distributor  common.Address
	PeriodFinish *big.Int
	Rate         *big.Int
	LastUpdate   *big.Int
	Integral     *big.Int
}, error) {
	return _CurveGauge.Contract.RewardData(&_CurveGauge.CallOpts, arg0)
}

// RewardData is a free data retrieval call binding the contract method 0x48e9c65e.
//
// Solidity: function reward_data(address arg0) view returns(address token, address distributor, uint256 period_finish, uint256 rate, uint256 last_update, uint256 integral)
func (_CurveGauge *CurveGaugeCallerSession) RewardData(arg0 common.Address) (struct {
	Token        common.Address
	Distributor  common.Address
	PeriodFinish *big.Int
	Rate         *big.Int
	LastUpdate   *big.Int
	Integral     *big.Int
}, error) {
	return _CurveGauge.Contract.RewardData(&_CurveGauge.CallOpts, arg0)
}

// RewardIntegralFor is a free data retrieval call binding the contract method 0xf05cc058.
//
// Solidity: function reward_integral_for(address arg0, address arg1) view returns(uint256)
func (_CurveGauge *CurveGaugeCaller) RewardIntegralFor(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "reward_integral_for", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardIntegralFor is a free data retrieval call binding the contract method 0xf05cc058.
//
// Solidity: function reward_integral_for(address arg0, address arg1) view returns(uint256)
func (_CurveGauge *CurveGaugeSession) RewardIntegralFor(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _CurveGauge.Contract.RewardIntegralFor(&_CurveGauge.CallOpts, arg0, arg1)
}

// RewardIntegralFor is a free data retrieval call binding the contract method 0xf05cc058.
//
// Solidity: function reward_integral_for(address arg0, address arg1) view returns(uint256)
func (_CurveGauge *CurveGaugeCallerSession) RewardIntegralFor(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _CurveGauge.Contract.RewardIntegralFor(&_CurveGauge.CallOpts, arg0, arg1)
}

// RewardTokens is a free data retrieval call binding the contract method 0x54c49fe9.
//
// Solidity: function reward_tokens(uint256 arg0) view returns(address)
func (_CurveGauge *CurveGaugeCaller) RewardTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "reward_tokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardTokens is a free data retrieval call binding the contract method 0x54c49fe9.
//
// Solidity: function reward_tokens(uint256 arg0) view returns(address)
func (_CurveGauge *CurveGaugeSession) RewardTokens(arg0 *big.Int) (common.Address, error) {
	return _CurveGauge.Contract.RewardTokens(&_CurveGauge.CallOpts, arg0)
}

// RewardTokens is a free data retrieval call binding the contract method 0x54c49fe9.
//
// Solidity: function reward_tokens(uint256 arg0) view returns(address)
func (_CurveGauge *CurveGaugeCallerSession) RewardTokens(arg0 *big.Int) (common.Address, error) {
	return _CurveGauge.Contract.RewardTokens(&_CurveGauge.CallOpts, arg0)
}

// RewardsReceiver is a free data retrieval call binding the contract method 0x01ddabf1.
//
// Solidity: function rewards_receiver(address arg0) view returns(address)
func (_CurveGauge *CurveGaugeCaller) RewardsReceiver(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "rewards_receiver", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsReceiver is a free data retrieval call binding the contract method 0x01ddabf1.
//
// Solidity: function rewards_receiver(address arg0) view returns(address)
func (_CurveGauge *CurveGaugeSession) RewardsReceiver(arg0 common.Address) (common.Address, error) {
	return _CurveGauge.Contract.RewardsReceiver(&_CurveGauge.CallOpts, arg0)
}

// RewardsReceiver is a free data retrieval call binding the contract method 0x01ddabf1.
//
// Solidity: function rewards_receiver(address arg0) view returns(address)
func (_CurveGauge *CurveGaugeCallerSession) RewardsReceiver(arg0 common.Address) (common.Address, error) {
	return _CurveGauge.Contract.RewardsReceiver(&_CurveGauge.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurveGauge *CurveGaugeCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurveGauge *CurveGaugeSession) Symbol() (string, error) {
	return _CurveGauge.Contract.Symbol(&_CurveGauge.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurveGauge *CurveGaugeCallerSession) Symbol() (string, error) {
	return _CurveGauge.Contract.Symbol(&_CurveGauge.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurveGauge *CurveGaugeCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurveGauge *CurveGaugeSession) TotalSupply() (*big.Int, error) {
	return _CurveGauge.Contract.TotalSupply(&_CurveGauge.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurveGauge *CurveGaugeCallerSession) TotalSupply() (*big.Int, error) {
	return _CurveGauge.Contract.TotalSupply(&_CurveGauge.CallOpts)
}

// WorkingBalances is a free data retrieval call binding the contract method 0x13ecb1ca.
//
// Solidity: function working_balances(address arg0) view returns(uint256)
func (_CurveGauge *CurveGaugeCaller) WorkingBalances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "working_balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WorkingBalances is a free data retrieval call binding the contract method 0x13ecb1ca.
//
// Solidity: function working_balances(address arg0) view returns(uint256)
func (_CurveGauge *CurveGaugeSession) WorkingBalances(arg0 common.Address) (*big.Int, error) {
	return _CurveGauge.Contract.WorkingBalances(&_CurveGauge.CallOpts, arg0)
}

// WorkingBalances is a free data retrieval call binding the contract method 0x13ecb1ca.
//
// Solidity: function working_balances(address arg0) view returns(uint256)
func (_CurveGauge *CurveGaugeCallerSession) WorkingBalances(arg0 common.Address) (*big.Int, error) {
	return _CurveGauge.Contract.WorkingBalances(&_CurveGauge.CallOpts, arg0)
}

// WorkingSupply is a free data retrieval call binding the contract method 0x17e28089.
//
// Solidity: function working_supply() view returns(uint256)
func (_CurveGauge *CurveGaugeCaller) WorkingSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "working_supply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WorkingSupply is a free data retrieval call binding the contract method 0x17e28089.
//
// Solidity: function working_supply() view returns(uint256)
func (_CurveGauge *CurveGaugeSession) WorkingSupply() (*big.Int, error) {
	return _CurveGauge.Contract.WorkingSupply(&_CurveGauge.CallOpts)
}

// WorkingSupply is a free data retrieval call binding the contract method 0x17e28089.
//
// Solidity: function working_supply() view returns(uint256)
func (_CurveGauge *CurveGaugeCallerSession) WorkingSupply() (*big.Int, error) {
	return _CurveGauge.Contract.WorkingSupply(&_CurveGauge.CallOpts)
}

// AcceptTransferOwnership is a paid mutator transaction binding the contract method 0xe5ea47b8.
//
// Solidity: function accept_transfer_ownership() returns()
func (_CurveGauge *CurveGaugeTransactor) AcceptTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveGauge.contract.Transact(opts, "accept_transfer_ownership")
}

// AcceptTransferOwnership is a paid mutator transaction binding the contract method 0xe5ea47b8.
//
// Solidity: function accept_transfer_ownership() returns()
func (_CurveGauge *CurveGaugeSession) AcceptTransferOwnership() (*types.Transaction, error) {
	return _CurveGauge.Contract.AcceptTransferOwnership(&_CurveGauge.TransactOpts)
}

// AcceptTransferOwnership is a paid mutator transaction binding the contract method 0xe5ea47b8.
//
// Solidity: function accept_transfer_ownership() returns()
func (_CurveGauge *CurveGaugeTransactorSession) AcceptTransferOwnership() (*types.Transaction, error) {
	return _CurveGauge.Contract.AcceptTransferOwnership(&_CurveGauge.TransactOpts)
}

// AddReward is a paid mutator transaction binding the contract method 0xe8de0d4d.
//
// Solidity: function add_reward(address _reward_token, address _distributor) returns()
func (_CurveGauge *CurveGaugeTransactor) AddReward(opts *bind.TransactOpts, _reward_token common.Address, _distributor common.Address) (*types.Transaction, error) {
	return _CurveGauge.contract.Transact(opts, "add_reward", _reward_token, _distributor)
}

// AddReward is a paid mutator transaction binding the contract method 0xe8de0d4d.
//
// Solidity: function add_reward(address _reward_token, address _distributor) returns()
func (_CurveGauge *CurveGaugeSession) AddReward(_reward_token common.Address, _distributor common.Address) (*types.Transaction, error) {
	return _CurveGauge.Contract.AddReward(&_CurveGauge.TransactOpts, _reward_token, _distributor)
}

// AddReward is a paid mutator transaction binding the contract method 0xe8de0d4d.
//
// Solidity: function add_reward(address _reward_token, address _distributor) returns()
func (_CurveGauge *CurveGaugeTransactorSession) AddReward(_reward_token common.Address, _distributor common.Address) (*types.Transaction, error) {
	return _CurveGauge.Contract.AddReward(&_CurveGauge.TransactOpts, _reward_token, _distributor)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_CurveGauge *CurveGaugeTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveGauge.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_CurveGauge *CurveGaugeSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveGauge.Contract.Approve(&_CurveGauge.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_CurveGauge *CurveGaugeTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveGauge.Contract.Approve(&_CurveGauge.TransactOpts, _spender, _value)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xe6f1daf2.
//
// Solidity: function claim_rewards() returns()
func (_CurveGauge *CurveGaugeTransactor) ClaimRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveGauge.contract.Transact(opts, "claim_rewards")
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xe6f1daf2.
//
// Solidity: function claim_rewards() returns()
func (_CurveGauge *CurveGaugeSession) ClaimRewards() (*types.Transaction, error) {
	return _CurveGauge.Contract.ClaimRewards(&_CurveGauge.TransactOpts)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xe6f1daf2.
//
// Solidity: function claim_rewards() returns()
func (_CurveGauge *CurveGaugeTransactorSession) ClaimRewards() (*types.Transaction, error) {
	return _CurveGauge.Contract.ClaimRewards(&_CurveGauge.TransactOpts)
}

// ClaimRewards0 is a paid mutator transaction binding the contract method 0x84e9bd7e.
//
// Solidity: function claim_rewards(address _addr) returns()
func (_CurveGauge *CurveGaugeTransactor) ClaimRewards0(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _CurveGauge.contract.Transact(opts, "claim_rewards0", _addr)
}

// ClaimRewards0 is a paid mutator transaction binding the contract method 0x84e9bd7e.
//
// Solidity: function claim_rewards(address _addr) returns()
func (_CurveGauge *CurveGaugeSession) ClaimRewards0(_addr common.Address) (*types.Transaction, error) {
	return _CurveGauge.Contract.ClaimRewards0(&_CurveGauge.TransactOpts, _addr)
}

// ClaimRewards0 is a paid mutator transaction binding the contract method 0x84e9bd7e.
//
// Solidity: function claim_rewards(address _addr) returns()
func (_CurveGauge *CurveGaugeTransactorSession) ClaimRewards0(_addr common.Address) (*types.Transaction, error) {
	return _CurveGauge.Contract.ClaimRewards0(&_CurveGauge.TransactOpts, _addr)
}

// ClaimRewards1 is a paid mutator transaction binding the contract method 0x9faceb1b.
//
// Solidity: function claim_rewards(address _addr, address _receiver) returns()
func (_CurveGauge *CurveGaugeTransactor) ClaimRewards1(opts *bind.TransactOpts, _addr common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _CurveGauge.contract.Transact(opts, "claim_rewards1", _addr, _receiver)
}

// ClaimRewards1 is a paid mutator transaction binding the contract method 0x9faceb1b.
//
// Solidity: function claim_rewards(address _addr, address _receiver) returns()
func (_CurveGauge *CurveGaugeSession) ClaimRewards1(_addr common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _CurveGauge.Contract.ClaimRewards1(&_CurveGauge.TransactOpts, _addr, _receiver)
}

// ClaimRewards1 is a paid mutator transaction binding the contract method 0x9faceb1b.
//
// Solidity: function claim_rewards(address _addr, address _receiver) returns()
func (_CurveGauge *CurveGaugeTransactorSession) ClaimRewards1(_addr common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _CurveGauge.Contract.ClaimRewards1(&_CurveGauge.TransactOpts, _addr, _receiver)
}

// ClaimableTokens is a paid mutator transaction binding the contract method 0x33134583.
//
// Solidity: function claimable_tokens(address addr) returns(uint256)
func (_CurveGauge *CurveGaugeTransactor) ClaimableTokens(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _CurveGauge.contract.Transact(opts, "claimable_tokens", addr)
}

// ClaimableTokens is a paid mutator transaction binding the contract method 0x33134583.
//
// Solidity: function claimable_tokens(address addr) returns(uint256)
func (_CurveGauge *CurveGaugeSession) ClaimableTokens(addr common.Address) (*types.Transaction, error) {
	return _CurveGauge.Contract.ClaimableTokens(&_CurveGauge.TransactOpts, addr)
}

// ClaimableTokens is a paid mutator transaction binding the contract method 0x33134583.
//
// Solidity: function claimable_tokens(address addr) returns(uint256)
func (_CurveGauge *CurveGaugeTransactorSession) ClaimableTokens(addr common.Address) (*types.Transaction, error) {
	return _CurveGauge.Contract.ClaimableTokens(&_CurveGauge.TransactOpts, addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address addr) returns()
func (_CurveGauge *CurveGaugeTransactor) CommitTransferOwnership(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _CurveGauge.contract.Transact(opts, "commit_transfer_ownership", addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address addr) returns()
func (_CurveGauge *CurveGaugeSession) CommitTransferOwnership(addr common.Address) (*types.Transaction, error) {
	return _CurveGauge.Contract.CommitTransferOwnership(&_CurveGauge.TransactOpts, addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address addr) returns()
func (_CurveGauge *CurveGaugeTransactorSession) CommitTransferOwnership(addr common.Address) (*types.Transaction, error) {
	return _CurveGauge.Contract.CommitTransferOwnership(&_CurveGauge.TransactOpts, addr)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtracted_value) returns(bool)
func (_CurveGauge *CurveGaugeTransactor) DecreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _subtracted_value *big.Int) (*types.Transaction, error) {
	return _CurveGauge.contract.Transact(opts, "decreaseAllowance", _spender, _subtracted_value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtracted_value) returns(bool)
func (_CurveGauge *CurveGaugeSession) DecreaseAllowance(_spender common.Address, _subtracted_value *big.Int) (*types.Transaction, error) {
	return _CurveGauge.Contract.DecreaseAllowance(&_CurveGauge.TransactOpts, _spender, _subtracted_value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtracted_value) returns(bool)
func (_CurveGauge *CurveGaugeTransactorSession) DecreaseAllowance(_spender common.Address, _subtracted_value *big.Int) (*types.Transaction, error) {
	return _CurveGauge.Contract.DecreaseAllowance(&_CurveGauge.TransactOpts, _spender, _subtracted_value)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _value) returns()
func (_CurveGauge *CurveGaugeTransactor) Deposit(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _CurveGauge.contract.Transact(opts, "deposit", _value)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _value) returns()
func (_CurveGauge *CurveGaugeSession) Deposit(_value *big.Int) (*types.Transaction, error) {
	return _CurveGauge.Contract.Deposit(&_CurveGauge.TransactOpts, _value)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _value) returns()
func (_CurveGauge *CurveGaugeTransactorSession) Deposit(_value *big.Int) (*types.Transaction, error) {
	return _CurveGauge.Contract.Deposit(&_CurveGauge.TransactOpts, _value)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _value, address _addr) returns()
func (_CurveGauge *CurveGaugeTransactor) Deposit0(opts *bind.TransactOpts, _value *big.Int, _addr common.Address) (*types.Transaction, error) {
	return _CurveGauge.contract.Transact(opts, "deposit0", _value, _addr)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _value, address _addr) returns()
func (_CurveGauge *CurveGaugeSession) Deposit0(_value *big.Int, _addr common.Address) (*types.Transaction, error) {
	return _CurveGauge.Contract.Deposit0(&_CurveGauge.TransactOpts, _value, _addr)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _value, address _addr) returns()
func (_CurveGauge *CurveGaugeTransactorSession) Deposit0(_value *big.Int, _addr common.Address) (*types.Transaction, error) {
	return _CurveGauge.Contract.Deposit0(&_CurveGauge.TransactOpts, _value, _addr)
}

// Deposit1 is a paid mutator transaction binding the contract method 0x83df6747.
//
// Solidity: function deposit(uint256 _value, address _addr, bool _claim_rewards) returns()
func (_CurveGauge *CurveGaugeTransactor) Deposit1(opts *bind.TransactOpts, _value *big.Int, _addr common.Address, _claim_rewards bool) (*types.Transaction, error) {
	return _CurveGauge.contract.Transact(opts, "deposit1", _value, _addr, _claim_rewards)
}

// Deposit1 is a paid mutator transaction binding the contract method 0x83df6747.
//
// Solidity: function deposit(uint256 _value, address _addr, bool _claim_rewards) returns()
func (_CurveGauge *CurveGaugeSession) Deposit1(_value *big.Int, _addr common.Address, _claim_rewards bool) (*types.Transaction, error) {
	return _CurveGauge.Contract.Deposit1(&_CurveGauge.TransactOpts, _value, _addr, _claim_rewards)
}

// Deposit1 is a paid mutator transaction binding the contract method 0x83df6747.
//
// Solidity: function deposit(uint256 _value, address _addr, bool _claim_rewards) returns()
func (_CurveGauge *CurveGaugeTransactorSession) Deposit1(_value *big.Int, _addr common.Address, _claim_rewards bool) (*types.Transaction, error) {
	return _CurveGauge.Contract.Deposit1(&_CurveGauge.TransactOpts, _value, _addr, _claim_rewards)
}

// DepositRewardToken is a paid mutator transaction binding the contract method 0x93f7aa67.
//
// Solidity: function deposit_reward_token(address _reward_token, uint256 _amount) returns()
func (_CurveGauge *CurveGaugeTransactor) DepositRewardToken(opts *bind.TransactOpts, _reward_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _CurveGauge.contract.Transact(opts, "deposit_reward_token", _reward_token, _amount)
}

// DepositRewardToken is a paid mutator transaction binding the contract method 0x93f7aa67.
//
// Solidity: function deposit_reward_token(address _reward_token, uint256 _amount) returns()
func (_CurveGauge *CurveGaugeSession) DepositRewardToken(_reward_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _CurveGauge.Contract.DepositRewardToken(&_CurveGauge.TransactOpts, _reward_token, _amount)
}

// DepositRewardToken is a paid mutator transaction binding the contract method 0x93f7aa67.
//
// Solidity: function deposit_reward_token(address _reward_token, uint256 _amount) returns()
func (_CurveGauge *CurveGaugeTransactorSession) DepositRewardToken(_reward_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _CurveGauge.Contract.DepositRewardToken(&_CurveGauge.TransactOpts, _reward_token, _amount)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _added_value) returns(bool)
func (_CurveGauge *CurveGaugeTransactor) IncreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _added_value *big.Int) (*types.Transaction, error) {
	return _CurveGauge.contract.Transact(opts, "increaseAllowance", _spender, _added_value)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _added_value) returns(bool)
func (_CurveGauge *CurveGaugeSession) IncreaseAllowance(_spender common.Address, _added_value *big.Int) (*types.Transaction, error) {
	return _CurveGauge.Contract.IncreaseAllowance(&_CurveGauge.TransactOpts, _spender, _added_value)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _added_value) returns(bool)
func (_CurveGauge *CurveGaugeTransactorSession) IncreaseAllowance(_spender common.Address, _added_value *big.Int) (*types.Transaction, error) {
	return _CurveGauge.Contract.IncreaseAllowance(&_CurveGauge.TransactOpts, _spender, _added_value)
}

// Kick is a paid mutator transaction binding the contract method 0x96c55175.
//
// Solidity: function kick(address addr) returns()
func (_CurveGauge *CurveGaugeTransactor) Kick(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _CurveGauge.contract.Transact(opts, "kick", addr)
}

// Kick is a paid mutator transaction binding the contract method 0x96c55175.
//
// Solidity: function kick(address addr) returns()
func (_CurveGauge *CurveGaugeSession) Kick(addr common.Address) (*types.Transaction, error) {
	return _CurveGauge.Contract.Kick(&_CurveGauge.TransactOpts, addr)
}

// Kick is a paid mutator transaction binding the contract method 0x96c55175.
//
// Solidity: function kick(address addr) returns()
func (_CurveGauge *CurveGaugeTransactorSession) Kick(addr common.Address) (*types.Transaction, error) {
	return _CurveGauge.Contract.Kick(&_CurveGauge.TransactOpts, addr)
}

// SetKilled is a paid mutator transaction binding the contract method 0x90b22997.
//
// Solidity: function set_killed(bool _is_killed) returns()
func (_CurveGauge *CurveGaugeTransactor) SetKilled(opts *bind.TransactOpts, _is_killed bool) (*types.Transaction, error) {
	return _CurveGauge.contract.Transact(opts, "set_killed", _is_killed)
}

// SetKilled is a paid mutator transaction binding the contract method 0x90b22997.
//
// Solidity: function set_killed(bool _is_killed) returns()
func (_CurveGauge *CurveGaugeSession) SetKilled(_is_killed bool) (*types.Transaction, error) {
	return _CurveGauge.Contract.SetKilled(&_CurveGauge.TransactOpts, _is_killed)
}

// SetKilled is a paid mutator transaction binding the contract method 0x90b22997.
//
// Solidity: function set_killed(bool _is_killed) returns()
func (_CurveGauge *CurveGaugeTransactorSession) SetKilled(_is_killed bool) (*types.Transaction, error) {
	return _CurveGauge.Contract.SetKilled(&_CurveGauge.TransactOpts, _is_killed)
}

// SetRewardDistributor is a paid mutator transaction binding the contract method 0x058a3a24.
//
// Solidity: function set_reward_distributor(address _reward_token, address _distributor) returns()
func (_CurveGauge *CurveGaugeTransactor) SetRewardDistributor(opts *bind.TransactOpts, _reward_token common.Address, _distributor common.Address) (*types.Transaction, error) {
	return _CurveGauge.contract.Transact(opts, "set_reward_distributor", _reward_token, _distributor)
}

// SetRewardDistributor is a paid mutator transaction binding the contract method 0x058a3a24.
//
// Solidity: function set_reward_distributor(address _reward_token, address _distributor) returns()
func (_CurveGauge *CurveGaugeSession) SetRewardDistributor(_reward_token common.Address, _distributor common.Address) (*types.Transaction, error) {
	return _CurveGauge.Contract.SetRewardDistributor(&_CurveGauge.TransactOpts, _reward_token, _distributor)
}

// SetRewardDistributor is a paid mutator transaction binding the contract method 0x058a3a24.
//
// Solidity: function set_reward_distributor(address _reward_token, address _distributor) returns()
func (_CurveGauge *CurveGaugeTransactorSession) SetRewardDistributor(_reward_token common.Address, _distributor common.Address) (*types.Transaction, error) {
	return _CurveGauge.Contract.SetRewardDistributor(&_CurveGauge.TransactOpts, _reward_token, _distributor)
}

// SetRewardsReceiver is a paid mutator transaction binding the contract method 0xbdf98116.
//
// Solidity: function set_rewards_receiver(address _receiver) returns()
func (_CurveGauge *CurveGaugeTransactor) SetRewardsReceiver(opts *bind.TransactOpts, _receiver common.Address) (*types.Transaction, error) {
	return _CurveGauge.contract.Transact(opts, "set_rewards_receiver", _receiver)
}

// SetRewardsReceiver is a paid mutator transaction binding the contract method 0xbdf98116.
//
// Solidity: function set_rewards_receiver(address _receiver) returns()
func (_CurveGauge *CurveGaugeSession) SetRewardsReceiver(_receiver common.Address) (*types.Transaction, error) {
	return _CurveGauge.Contract.SetRewardsReceiver(&_CurveGauge.TransactOpts, _receiver)
}

// SetRewardsReceiver is a paid mutator transaction binding the contract method 0xbdf98116.
//
// Solidity: function set_rewards_receiver(address _receiver) returns()
func (_CurveGauge *CurveGaugeTransactorSession) SetRewardsReceiver(_receiver common.Address) (*types.Transaction, error) {
	return _CurveGauge.Contract.SetRewardsReceiver(&_CurveGauge.TransactOpts, _receiver)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_CurveGauge *CurveGaugeTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveGauge.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_CurveGauge *CurveGaugeSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveGauge.Contract.Transfer(&_CurveGauge.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_CurveGauge *CurveGaugeTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveGauge.Contract.Transfer(&_CurveGauge.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_CurveGauge *CurveGaugeTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveGauge.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_CurveGauge *CurveGaugeSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveGauge.Contract.TransferFrom(&_CurveGauge.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_CurveGauge *CurveGaugeTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveGauge.Contract.TransferFrom(&_CurveGauge.TransactOpts, _from, _to, _value)
}

// UserCheckpoint is a paid mutator transaction binding the contract method 0x4b820093.
//
// Solidity: function user_checkpoint(address addr) returns(bool)
func (_CurveGauge *CurveGaugeTransactor) UserCheckpoint(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _CurveGauge.contract.Transact(opts, "user_checkpoint", addr)
}

// UserCheckpoint is a paid mutator transaction binding the contract method 0x4b820093.
//
// Solidity: function user_checkpoint(address addr) returns(bool)
func (_CurveGauge *CurveGaugeSession) UserCheckpoint(addr common.Address) (*types.Transaction, error) {
	return _CurveGauge.Contract.UserCheckpoint(&_CurveGauge.TransactOpts, addr)
}

// UserCheckpoint is a paid mutator transaction binding the contract method 0x4b820093.
//
// Solidity: function user_checkpoint(address addr) returns(bool)
func (_CurveGauge *CurveGaugeTransactorSession) UserCheckpoint(addr common.Address) (*types.Transaction, error) {
	return _CurveGauge.Contract.UserCheckpoint(&_CurveGauge.TransactOpts, addr)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _value) returns()
func (_CurveGauge *CurveGaugeTransactor) Withdraw(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _CurveGauge.contract.Transact(opts, "withdraw", _value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _value) returns()
func (_CurveGauge *CurveGaugeSession) Withdraw(_value *big.Int) (*types.Transaction, error) {
	return _CurveGauge.Contract.Withdraw(&_CurveGauge.TransactOpts, _value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _value) returns()
func (_CurveGauge *CurveGaugeTransactorSession) Withdraw(_value *big.Int) (*types.Transaction, error) {
	return _CurveGauge.Contract.Withdraw(&_CurveGauge.TransactOpts, _value)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x38d07436.
//
// Solidity: function withdraw(uint256 _value, bool _claim_rewards) returns()
func (_CurveGauge *CurveGaugeTransactor) Withdraw0(opts *bind.TransactOpts, _value *big.Int, _claim_rewards bool) (*types.Transaction, error) {
	return _CurveGauge.contract.Transact(opts, "withdraw0", _value, _claim_rewards)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x38d07436.
//
// Solidity: function withdraw(uint256 _value, bool _claim_rewards) returns()
func (_CurveGauge *CurveGaugeSession) Withdraw0(_value *big.Int, _claim_rewards bool) (*types.Transaction, error) {
	return _CurveGauge.Contract.Withdraw0(&_CurveGauge.TransactOpts, _value, _claim_rewards)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x38d07436.
//
// Solidity: function withdraw(uint256 _value, bool _claim_rewards) returns()
func (_CurveGauge *CurveGaugeTransactorSession) Withdraw0(_value *big.Int, _claim_rewards bool) (*types.Transaction, error) {
	return _CurveGauge.Contract.Withdraw0(&_CurveGauge.TransactOpts, _value, _claim_rewards)
}

// CurveGaugeApplyOwnershipIterator is returned from FilterApplyOwnership and is used to iterate over the raw logs and unpacked data for ApplyOwnership events raised by the CurveGauge contract.
type CurveGaugeApplyOwnershipIterator struct {
	Event *CurveGaugeApplyOwnership // Event containing the contract specifics and raw log

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
func (it *CurveGaugeApplyOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveGaugeApplyOwnership)
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
		it.Event = new(CurveGaugeApplyOwnership)
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
func (it *CurveGaugeApplyOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveGaugeApplyOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveGaugeApplyOwnership represents a ApplyOwnership event raised by the CurveGauge contract.
type CurveGaugeApplyOwnership struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterApplyOwnership is a free log retrieval operation binding the contract event 0xebee2d5739011062cb4f14113f3b36bf0ffe3da5c0568f64189d1012a1189105.
//
// Solidity: event ApplyOwnership(address admin)
func (_CurveGauge *CurveGaugeFilterer) FilterApplyOwnership(opts *bind.FilterOpts) (*CurveGaugeApplyOwnershipIterator, error) {

	logs, sub, err := _CurveGauge.contract.FilterLogs(opts, "ApplyOwnership")
	if err != nil {
		return nil, err
	}
	return &CurveGaugeApplyOwnershipIterator{contract: _CurveGauge.contract, event: "ApplyOwnership", logs: logs, sub: sub}, nil
}

// WatchApplyOwnership is a free log subscription operation binding the contract event 0xebee2d5739011062cb4f14113f3b36bf0ffe3da5c0568f64189d1012a1189105.
//
// Solidity: event ApplyOwnership(address admin)
func (_CurveGauge *CurveGaugeFilterer) WatchApplyOwnership(opts *bind.WatchOpts, sink chan<- *CurveGaugeApplyOwnership) (event.Subscription, error) {

	logs, sub, err := _CurveGauge.contract.WatchLogs(opts, "ApplyOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveGaugeApplyOwnership)
				if err := _CurveGauge.contract.UnpackLog(event, "ApplyOwnership", log); err != nil {
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

// ParseApplyOwnership is a log parse operation binding the contract event 0xebee2d5739011062cb4f14113f3b36bf0ffe3da5c0568f64189d1012a1189105.
//
// Solidity: event ApplyOwnership(address admin)
func (_CurveGauge *CurveGaugeFilterer) ParseApplyOwnership(log types.Log) (*CurveGaugeApplyOwnership, error) {
	event := new(CurveGaugeApplyOwnership)
	if err := _CurveGauge.contract.UnpackLog(event, "ApplyOwnership", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveGaugeApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CurveGauge contract.
type CurveGaugeApprovalIterator struct {
	Event *CurveGaugeApproval // Event containing the contract specifics and raw log

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
func (it *CurveGaugeApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveGaugeApproval)
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
		it.Event = new(CurveGaugeApproval)
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
func (it *CurveGaugeApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveGaugeApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveGaugeApproval represents a Approval event raised by the CurveGauge contract.
type CurveGaugeApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_CurveGauge *CurveGaugeFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*CurveGaugeApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _CurveGauge.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &CurveGaugeApprovalIterator{contract: _CurveGauge.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_CurveGauge *CurveGaugeFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CurveGaugeApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _CurveGauge.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveGaugeApproval)
				if err := _CurveGauge.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_CurveGauge *CurveGaugeFilterer) ParseApproval(log types.Log) (*CurveGaugeApproval, error) {
	event := new(CurveGaugeApproval)
	if err := _CurveGauge.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveGaugeCommitOwnershipIterator is returned from FilterCommitOwnership and is used to iterate over the raw logs and unpacked data for CommitOwnership events raised by the CurveGauge contract.
type CurveGaugeCommitOwnershipIterator struct {
	Event *CurveGaugeCommitOwnership // Event containing the contract specifics and raw log

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
func (it *CurveGaugeCommitOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveGaugeCommitOwnership)
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
		it.Event = new(CurveGaugeCommitOwnership)
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
func (it *CurveGaugeCommitOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveGaugeCommitOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveGaugeCommitOwnership represents a CommitOwnership event raised by the CurveGauge contract.
type CurveGaugeCommitOwnership struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCommitOwnership is a free log retrieval operation binding the contract event 0x2f56810a6bf40af059b96d3aea4db54081f378029a518390491093a7b67032e9.
//
// Solidity: event CommitOwnership(address admin)
func (_CurveGauge *CurveGaugeFilterer) FilterCommitOwnership(opts *bind.FilterOpts) (*CurveGaugeCommitOwnershipIterator, error) {

	logs, sub, err := _CurveGauge.contract.FilterLogs(opts, "CommitOwnership")
	if err != nil {
		return nil, err
	}
	return &CurveGaugeCommitOwnershipIterator{contract: _CurveGauge.contract, event: "CommitOwnership", logs: logs, sub: sub}, nil
}

// WatchCommitOwnership is a free log subscription operation binding the contract event 0x2f56810a6bf40af059b96d3aea4db54081f378029a518390491093a7b67032e9.
//
// Solidity: event CommitOwnership(address admin)
func (_CurveGauge *CurveGaugeFilterer) WatchCommitOwnership(opts *bind.WatchOpts, sink chan<- *CurveGaugeCommitOwnership) (event.Subscription, error) {

	logs, sub, err := _CurveGauge.contract.WatchLogs(opts, "CommitOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveGaugeCommitOwnership)
				if err := _CurveGauge.contract.UnpackLog(event, "CommitOwnership", log); err != nil {
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

// ParseCommitOwnership is a log parse operation binding the contract event 0x2f56810a6bf40af059b96d3aea4db54081f378029a518390491093a7b67032e9.
//
// Solidity: event CommitOwnership(address admin)
func (_CurveGauge *CurveGaugeFilterer) ParseCommitOwnership(log types.Log) (*CurveGaugeCommitOwnership, error) {
	event := new(CurveGaugeCommitOwnership)
	if err := _CurveGauge.contract.UnpackLog(event, "CommitOwnership", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveGaugeDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the CurveGauge contract.
type CurveGaugeDepositIterator struct {
	Event *CurveGaugeDeposit // Event containing the contract specifics and raw log

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
func (it *CurveGaugeDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveGaugeDeposit)
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
		it.Event = new(CurveGaugeDeposit)
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
func (it *CurveGaugeDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveGaugeDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveGaugeDeposit represents a Deposit event raised by the CurveGauge contract.
type CurveGaugeDeposit struct {
	Provider common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed provider, uint256 value)
func (_CurveGauge *CurveGaugeFilterer) FilterDeposit(opts *bind.FilterOpts, provider []common.Address) (*CurveGaugeDepositIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurveGauge.contract.FilterLogs(opts, "Deposit", providerRule)
	if err != nil {
		return nil, err
	}
	return &CurveGaugeDepositIterator{contract: _CurveGauge.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed provider, uint256 value)
func (_CurveGauge *CurveGaugeFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *CurveGaugeDeposit, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurveGauge.contract.WatchLogs(opts, "Deposit", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveGaugeDeposit)
				if err := _CurveGauge.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed provider, uint256 value)
func (_CurveGauge *CurveGaugeFilterer) ParseDeposit(log types.Log) (*CurveGaugeDeposit, error) {
	event := new(CurveGaugeDeposit)
	if err := _CurveGauge.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveGaugeTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CurveGauge contract.
type CurveGaugeTransferIterator struct {
	Event *CurveGaugeTransfer // Event containing the contract specifics and raw log

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
func (it *CurveGaugeTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveGaugeTransfer)
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
		it.Event = new(CurveGaugeTransfer)
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
func (it *CurveGaugeTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveGaugeTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveGaugeTransfer represents a Transfer event raised by the CurveGauge contract.
type CurveGaugeTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_CurveGauge *CurveGaugeFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*CurveGaugeTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _CurveGauge.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &CurveGaugeTransferIterator{contract: _CurveGauge.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_CurveGauge *CurveGaugeFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CurveGaugeTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _CurveGauge.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveGaugeTransfer)
				if err := _CurveGauge.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_CurveGauge *CurveGaugeFilterer) ParseTransfer(log types.Log) (*CurveGaugeTransfer, error) {
	event := new(CurveGaugeTransfer)
	if err := _CurveGauge.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveGaugeUpdateLiquidityLimitIterator is returned from FilterUpdateLiquidityLimit and is used to iterate over the raw logs and unpacked data for UpdateLiquidityLimit events raised by the CurveGauge contract.
type CurveGaugeUpdateLiquidityLimitIterator struct {
	Event *CurveGaugeUpdateLiquidityLimit // Event containing the contract specifics and raw log

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
func (it *CurveGaugeUpdateLiquidityLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveGaugeUpdateLiquidityLimit)
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
		it.Event = new(CurveGaugeUpdateLiquidityLimit)
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
func (it *CurveGaugeUpdateLiquidityLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveGaugeUpdateLiquidityLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveGaugeUpdateLiquidityLimit represents a UpdateLiquidityLimit event raised by the CurveGauge contract.
type CurveGaugeUpdateLiquidityLimit struct {
	User            common.Address
	OriginalBalance *big.Int
	OriginalSupply  *big.Int
	WorkingBalance  *big.Int
	WorkingSupply   *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateLiquidityLimit is a free log retrieval operation binding the contract event 0x7ecd84343f76a23d2227290e0288da3251b045541698e575a5515af4f04197a3.
//
// Solidity: event UpdateLiquidityLimit(address user, uint256 original_balance, uint256 original_supply, uint256 working_balance, uint256 working_supply)
func (_CurveGauge *CurveGaugeFilterer) FilterUpdateLiquidityLimit(opts *bind.FilterOpts) (*CurveGaugeUpdateLiquidityLimitIterator, error) {

	logs, sub, err := _CurveGauge.contract.FilterLogs(opts, "UpdateLiquidityLimit")
	if err != nil {
		return nil, err
	}
	return &CurveGaugeUpdateLiquidityLimitIterator{contract: _CurveGauge.contract, event: "UpdateLiquidityLimit", logs: logs, sub: sub}, nil
}

// WatchUpdateLiquidityLimit is a free log subscription operation binding the contract event 0x7ecd84343f76a23d2227290e0288da3251b045541698e575a5515af4f04197a3.
//
// Solidity: event UpdateLiquidityLimit(address user, uint256 original_balance, uint256 original_supply, uint256 working_balance, uint256 working_supply)
func (_CurveGauge *CurveGaugeFilterer) WatchUpdateLiquidityLimit(opts *bind.WatchOpts, sink chan<- *CurveGaugeUpdateLiquidityLimit) (event.Subscription, error) {

	logs, sub, err := _CurveGauge.contract.WatchLogs(opts, "UpdateLiquidityLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveGaugeUpdateLiquidityLimit)
				if err := _CurveGauge.contract.UnpackLog(event, "UpdateLiquidityLimit", log); err != nil {
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

// ParseUpdateLiquidityLimit is a log parse operation binding the contract event 0x7ecd84343f76a23d2227290e0288da3251b045541698e575a5515af4f04197a3.
//
// Solidity: event UpdateLiquidityLimit(address user, uint256 original_balance, uint256 original_supply, uint256 working_balance, uint256 working_supply)
func (_CurveGauge *CurveGaugeFilterer) ParseUpdateLiquidityLimit(log types.Log) (*CurveGaugeUpdateLiquidityLimit, error) {
	event := new(CurveGaugeUpdateLiquidityLimit)
	if err := _CurveGauge.contract.UnpackLog(event, "UpdateLiquidityLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveGaugeWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the CurveGauge contract.
type CurveGaugeWithdrawIterator struct {
	Event *CurveGaugeWithdraw // Event containing the contract specifics and raw log

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
func (it *CurveGaugeWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveGaugeWithdraw)
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
		it.Event = new(CurveGaugeWithdraw)
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
func (it *CurveGaugeWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveGaugeWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveGaugeWithdraw represents a Withdraw event raised by the CurveGauge contract.
type CurveGaugeWithdraw struct {
	Provider common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed provider, uint256 value)
func (_CurveGauge *CurveGaugeFilterer) FilterWithdraw(opts *bind.FilterOpts, provider []common.Address) (*CurveGaugeWithdrawIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurveGauge.contract.FilterLogs(opts, "Withdraw", providerRule)
	if err != nil {
		return nil, err
	}
	return &CurveGaugeWithdrawIterator{contract: _CurveGauge.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed provider, uint256 value)
func (_CurveGauge *CurveGaugeFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *CurveGaugeWithdraw, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurveGauge.contract.WatchLogs(opts, "Withdraw", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveGaugeWithdraw)
				if err := _CurveGauge.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
// Solidity: event Withdraw(address indexed provider, uint256 value)
func (_CurveGauge *CurveGaugeFilterer) ParseWithdraw(log types.Log) (*CurveGaugeWithdraw, error) {
	event := new(CurveGaugeWithdraw)
	if err := _CurveGauge.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
