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

// ERC4626MetaData contains all meta data concerning the ERC4626 contract.
var ERC4626MetaData = &bind.MetaData{
	ABI: "[{\"name\":\"Rewards\",\"inputs\":[{\"name\":\"pending\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"streaming\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"unlocked\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"delta\",\"type\":\"int256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SetFeeRate\",\"inputs\":[{\"name\":\"fee_rate\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SetHalfTime\",\"inputs\":[{\"name\":\"half_time\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"PendingManagement\",\"inputs\":[{\"name\":\"management\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SetManagement\",\"inputs\":[{\"name\":\"management\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SetTreasury\",\"inputs\":[{\"name\":\"treasury\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Transfer\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Deposit\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true},{\"name\":\"assets\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true},{\"name\":\"assets\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"_asset\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"increaseAllowance\",\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"decreaseAllowance\",\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"totalAssets\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"convertToShares\",\"inputs\":[{\"name\":\"_assets\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"convertToAssets\",\"inputs\":[{\"name\":\"_shares\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"maxDeposit\",\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"previewDeposit\",\"inputs\":[{\"name\":\"_assets\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"_assets\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"_assets\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"maxMint\",\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"previewMint\",\"inputs\":[{\"name\":\"_shares\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"_shares\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"_shares\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"maxWithdraw\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"previewWithdraw\",\"inputs\":[{\"name\":\"_assets\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"_assets\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"_assets\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"_assets\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"},{\"name\":\"_owner\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"maxRedeem\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"previewRedeem\",\"inputs\":[{\"name\":\"_shares\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"redeem\",\"inputs\":[{\"name\":\"_shares\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"redeem\",\"inputs\":[{\"name\":\"_shares\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"redeem\",\"inputs\":[{\"name\":\"_shares\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"},{\"name\":\"_owner\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"update_amounts\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_amounts\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"int256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"vote_weight\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"known\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"rescue\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_performance_fee_rate\",\"inputs\":[{\"name\":\"_fee_rate\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_half_time\",\"inputs\":[{\"name\":\"_half_time\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_management\",\"inputs\":[{\"name\":\"_management\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"accept_management\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_treasury\",\"inputs\":[{\"name\":\"_treasury\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"updated\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"management\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"pending_management\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"performance_fee_rate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"treasury\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"half_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"},{\"name\":\"arg1\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"asset\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]}]",
}

// ERC4626ABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC4626MetaData.ABI instead.
var ERC4626ABI = ERC4626MetaData.ABI

// ERC4626 is an auto generated Go binding around an Ethereum contract.
type ERC4626 struct {
	ERC4626Caller     // Read-only binding to the contract
	ERC4626Transactor // Write-only binding to the contract
	ERC4626Filterer   // Log filterer for contract events
}

// ERC4626Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC4626Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC4626Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC4626Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC4626Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC4626Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC4626Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC4626Session struct {
	Contract     *ERC4626          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC4626CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC4626CallerSession struct {
	Contract *ERC4626Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ERC4626TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC4626TransactorSession struct {
	Contract     *ERC4626Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ERC4626Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC4626Raw struct {
	Contract *ERC4626 // Generic contract binding to access the raw methods on
}

// ERC4626CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC4626CallerRaw struct {
	Contract *ERC4626Caller // Generic read-only contract binding to access the raw methods on
}

// ERC4626TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC4626TransactorRaw struct {
	Contract *ERC4626Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC4626 creates a new instance of ERC4626, bound to a specific deployed contract.
func NewERC4626(address common.Address, backend bind.ContractBackend) (*ERC4626, error) {
	contract, err := bindERC4626(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC4626{ERC4626Caller: ERC4626Caller{contract: contract}, ERC4626Transactor: ERC4626Transactor{contract: contract}, ERC4626Filterer: ERC4626Filterer{contract: contract}}, nil
}

// NewERC4626Caller creates a new read-only instance of ERC4626, bound to a specific deployed contract.
func NewERC4626Caller(address common.Address, caller bind.ContractCaller) (*ERC4626Caller, error) {
	contract, err := bindERC4626(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC4626Caller{contract: contract}, nil
}

// NewERC4626Transactor creates a new write-only instance of ERC4626, bound to a specific deployed contract.
func NewERC4626Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC4626Transactor, error) {
	contract, err := bindERC4626(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC4626Transactor{contract: contract}, nil
}

// NewERC4626Filterer creates a new log filterer instance of ERC4626, bound to a specific deployed contract.
func NewERC4626Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC4626Filterer, error) {
	contract, err := bindERC4626(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC4626Filterer{contract: contract}, nil
}

// bindERC4626 binds a generic wrapper to an already deployed contract.
func bindERC4626(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC4626ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC4626 *ERC4626Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC4626.Contract.ERC4626Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC4626 *ERC4626Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC4626.Contract.ERC4626Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC4626 *ERC4626Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC4626.Contract.ERC4626Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC4626 *ERC4626CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC4626.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC4626 *ERC4626TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC4626.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC4626 *ERC4626TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC4626.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_ERC4626 *ERC4626Caller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_ERC4626 *ERC4626Session) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _ERC4626.Contract.Allowance(&_ERC4626.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_ERC4626 *ERC4626CallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _ERC4626.Contract.Allowance(&_ERC4626.CallOpts, arg0, arg1)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_ERC4626 *ERC4626Caller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC4626.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_ERC4626 *ERC4626Session) Asset() (common.Address, error) {
	return _ERC4626.Contract.Asset(&_ERC4626.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_ERC4626 *ERC4626CallerSession) Asset() (common.Address, error) {
	return _ERC4626.Contract.Asset(&_ERC4626.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_ERC4626 *ERC4626Caller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_ERC4626 *ERC4626Session) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _ERC4626.Contract.BalanceOf(&_ERC4626.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_ERC4626 *ERC4626CallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _ERC4626.Contract.BalanceOf(&_ERC4626.CallOpts, arg0)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 _shares) view returns(uint256)
func (_ERC4626 *ERC4626Caller) ConvertToAssets(opts *bind.CallOpts, _shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626.contract.Call(opts, &out, "convertToAssets", _shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 _shares) view returns(uint256)
func (_ERC4626 *ERC4626Session) ConvertToAssets(_shares *big.Int) (*big.Int, error) {
	return _ERC4626.Contract.ConvertToAssets(&_ERC4626.CallOpts, _shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 _shares) view returns(uint256)
func (_ERC4626 *ERC4626CallerSession) ConvertToAssets(_shares *big.Int) (*big.Int, error) {
	return _ERC4626.Contract.ConvertToAssets(&_ERC4626.CallOpts, _shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 _assets) view returns(uint256)
func (_ERC4626 *ERC4626Caller) ConvertToShares(opts *bind.CallOpts, _assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626.contract.Call(opts, &out, "convertToShares", _assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 _assets) view returns(uint256)
func (_ERC4626 *ERC4626Session) ConvertToShares(_assets *big.Int) (*big.Int, error) {
	return _ERC4626.Contract.ConvertToShares(&_ERC4626.CallOpts, _assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 _assets) view returns(uint256)
func (_ERC4626 *ERC4626CallerSession) ConvertToShares(_assets *big.Int) (*big.Int, error) {
	return _ERC4626.Contract.ConvertToShares(&_ERC4626.CallOpts, _assets)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC4626 *ERC4626Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC4626.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC4626 *ERC4626Session) Decimals() (uint8, error) {
	return _ERC4626.Contract.Decimals(&_ERC4626.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC4626 *ERC4626CallerSession) Decimals() (uint8, error) {
	return _ERC4626.Contract.Decimals(&_ERC4626.CallOpts)
}

// GetAmounts is a free data retrieval call binding the contract method 0xccaf0c30.
//
// Solidity: function get_amounts() view returns(uint256, uint256, uint256, uint256, int256)
func (_ERC4626 *ERC4626Caller) GetAmounts(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ERC4626.contract.Call(opts, &out, "get_amounts")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, err

}

// GetAmounts is a free data retrieval call binding the contract method 0xccaf0c30.
//
// Solidity: function get_amounts() view returns(uint256, uint256, uint256, uint256, int256)
func (_ERC4626 *ERC4626Session) GetAmounts() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _ERC4626.Contract.GetAmounts(&_ERC4626.CallOpts)
}

// GetAmounts is a free data retrieval call binding the contract method 0xccaf0c30.
//
// Solidity: function get_amounts() view returns(uint256, uint256, uint256, uint256, int256)
func (_ERC4626 *ERC4626CallerSession) GetAmounts() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _ERC4626.Contract.GetAmounts(&_ERC4626.CallOpts)
}

// HalfTime is a free data retrieval call binding the contract method 0xc5c91533.
//
// Solidity: function half_time() view returns(uint256)
func (_ERC4626 *ERC4626Caller) HalfTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626.contract.Call(opts, &out, "half_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HalfTime is a free data retrieval call binding the contract method 0xc5c91533.
//
// Solidity: function half_time() view returns(uint256)
func (_ERC4626 *ERC4626Session) HalfTime() (*big.Int, error) {
	return _ERC4626.Contract.HalfTime(&_ERC4626.CallOpts)
}

// HalfTime is a free data retrieval call binding the contract method 0xc5c91533.
//
// Solidity: function half_time() view returns(uint256)
func (_ERC4626 *ERC4626CallerSession) HalfTime() (*big.Int, error) {
	return _ERC4626.Contract.HalfTime(&_ERC4626.CallOpts)
}

// Known is a free data retrieval call binding the contract method 0xc445c938.
//
// Solidity: function known() view returns(uint256)
func (_ERC4626 *ERC4626Caller) Known(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626.contract.Call(opts, &out, "known")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Known is a free data retrieval call binding the contract method 0xc445c938.
//
// Solidity: function known() view returns(uint256)
func (_ERC4626 *ERC4626Session) Known() (*big.Int, error) {
	return _ERC4626.Contract.Known(&_ERC4626.CallOpts)
}

// Known is a free data retrieval call binding the contract method 0xc445c938.
//
// Solidity: function known() view returns(uint256)
func (_ERC4626 *ERC4626CallerSession) Known() (*big.Int, error) {
	return _ERC4626.Contract.Known(&_ERC4626.CallOpts)
}

// Management is a free data retrieval call binding the contract method 0x88a8d602.
//
// Solidity: function management() view returns(address)
func (_ERC4626 *ERC4626Caller) Management(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC4626.contract.Call(opts, &out, "management")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Management is a free data retrieval call binding the contract method 0x88a8d602.
//
// Solidity: function management() view returns(address)
func (_ERC4626 *ERC4626Session) Management() (common.Address, error) {
	return _ERC4626.Contract.Management(&_ERC4626.CallOpts)
}

// Management is a free data retrieval call binding the contract method 0x88a8d602.
//
// Solidity: function management() view returns(address)
func (_ERC4626 *ERC4626CallerSession) Management() (common.Address, error) {
	return _ERC4626.Contract.Management(&_ERC4626.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address _receiver) view returns(uint256)
func (_ERC4626 *ERC4626Caller) MaxDeposit(opts *bind.CallOpts, _receiver common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626.contract.Call(opts, &out, "maxDeposit", _receiver)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address _receiver) view returns(uint256)
func (_ERC4626 *ERC4626Session) MaxDeposit(_receiver common.Address) (*big.Int, error) {
	return _ERC4626.Contract.MaxDeposit(&_ERC4626.CallOpts, _receiver)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address _receiver) view returns(uint256)
func (_ERC4626 *ERC4626CallerSession) MaxDeposit(_receiver common.Address) (*big.Int, error) {
	return _ERC4626.Contract.MaxDeposit(&_ERC4626.CallOpts, _receiver)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address _receiver) view returns(uint256)
func (_ERC4626 *ERC4626Caller) MaxMint(opts *bind.CallOpts, _receiver common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626.contract.Call(opts, &out, "maxMint", _receiver)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address _receiver) view returns(uint256)
func (_ERC4626 *ERC4626Session) MaxMint(_receiver common.Address) (*big.Int, error) {
	return _ERC4626.Contract.MaxMint(&_ERC4626.CallOpts, _receiver)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address _receiver) view returns(uint256)
func (_ERC4626 *ERC4626CallerSession) MaxMint(_receiver common.Address) (*big.Int, error) {
	return _ERC4626.Contract.MaxMint(&_ERC4626.CallOpts, _receiver)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address _owner) view returns(uint256)
func (_ERC4626 *ERC4626Caller) MaxRedeem(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626.contract.Call(opts, &out, "maxRedeem", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address _owner) view returns(uint256)
func (_ERC4626 *ERC4626Session) MaxRedeem(_owner common.Address) (*big.Int, error) {
	return _ERC4626.Contract.MaxRedeem(&_ERC4626.CallOpts, _owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address _owner) view returns(uint256)
func (_ERC4626 *ERC4626CallerSession) MaxRedeem(_owner common.Address) (*big.Int, error) {
	return _ERC4626.Contract.MaxRedeem(&_ERC4626.CallOpts, _owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address _owner) view returns(uint256)
func (_ERC4626 *ERC4626Caller) MaxWithdraw(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626.contract.Call(opts, &out, "maxWithdraw", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address _owner) view returns(uint256)
func (_ERC4626 *ERC4626Session) MaxWithdraw(_owner common.Address) (*big.Int, error) {
	return _ERC4626.Contract.MaxWithdraw(&_ERC4626.CallOpts, _owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address _owner) view returns(uint256)
func (_ERC4626 *ERC4626CallerSession) MaxWithdraw(_owner common.Address) (*big.Int, error) {
	return _ERC4626.Contract.MaxWithdraw(&_ERC4626.CallOpts, _owner)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC4626 *ERC4626Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC4626.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC4626 *ERC4626Session) Name() (string, error) {
	return _ERC4626.Contract.Name(&_ERC4626.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC4626 *ERC4626CallerSession) Name() (string, error) {
	return _ERC4626.Contract.Name(&_ERC4626.CallOpts)
}

// PendingManagement is a free data retrieval call binding the contract method 0x770817ec.
//
// Solidity: function pending_management() view returns(address)
func (_ERC4626 *ERC4626Caller) PendingManagement(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC4626.contract.Call(opts, &out, "pending_management")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingManagement is a free data retrieval call binding the contract method 0x770817ec.
//
// Solidity: function pending_management() view returns(address)
func (_ERC4626 *ERC4626Session) PendingManagement() (common.Address, error) {
	return _ERC4626.Contract.PendingManagement(&_ERC4626.CallOpts)
}

// PendingManagement is a free data retrieval call binding the contract method 0x770817ec.
//
// Solidity: function pending_management() view returns(address)
func (_ERC4626 *ERC4626CallerSession) PendingManagement() (common.Address, error) {
	return _ERC4626.Contract.PendingManagement(&_ERC4626.CallOpts)
}

// PerformanceFeeRate is a free data retrieval call binding the contract method 0x18f7b782.
//
// Solidity: function performance_fee_rate() view returns(uint256)
func (_ERC4626 *ERC4626Caller) PerformanceFeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626.contract.Call(opts, &out, "performance_fee_rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PerformanceFeeRate is a free data retrieval call binding the contract method 0x18f7b782.
//
// Solidity: function performance_fee_rate() view returns(uint256)
func (_ERC4626 *ERC4626Session) PerformanceFeeRate() (*big.Int, error) {
	return _ERC4626.Contract.PerformanceFeeRate(&_ERC4626.CallOpts)
}

// PerformanceFeeRate is a free data retrieval call binding the contract method 0x18f7b782.
//
// Solidity: function performance_fee_rate() view returns(uint256)
func (_ERC4626 *ERC4626CallerSession) PerformanceFeeRate() (*big.Int, error) {
	return _ERC4626.Contract.PerformanceFeeRate(&_ERC4626.CallOpts)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 _assets) view returns(uint256)
func (_ERC4626 *ERC4626Caller) PreviewDeposit(opts *bind.CallOpts, _assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626.contract.Call(opts, &out, "previewDeposit", _assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 _assets) view returns(uint256)
func (_ERC4626 *ERC4626Session) PreviewDeposit(_assets *big.Int) (*big.Int, error) {
	return _ERC4626.Contract.PreviewDeposit(&_ERC4626.CallOpts, _assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 _assets) view returns(uint256)
func (_ERC4626 *ERC4626CallerSession) PreviewDeposit(_assets *big.Int) (*big.Int, error) {
	return _ERC4626.Contract.PreviewDeposit(&_ERC4626.CallOpts, _assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 _shares) view returns(uint256)
func (_ERC4626 *ERC4626Caller) PreviewMint(opts *bind.CallOpts, _shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626.contract.Call(opts, &out, "previewMint", _shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 _shares) view returns(uint256)
func (_ERC4626 *ERC4626Session) PreviewMint(_shares *big.Int) (*big.Int, error) {
	return _ERC4626.Contract.PreviewMint(&_ERC4626.CallOpts, _shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 _shares) view returns(uint256)
func (_ERC4626 *ERC4626CallerSession) PreviewMint(_shares *big.Int) (*big.Int, error) {
	return _ERC4626.Contract.PreviewMint(&_ERC4626.CallOpts, _shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 _shares) view returns(uint256)
func (_ERC4626 *ERC4626Caller) PreviewRedeem(opts *bind.CallOpts, _shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626.contract.Call(opts, &out, "previewRedeem", _shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 _shares) view returns(uint256)
func (_ERC4626 *ERC4626Session) PreviewRedeem(_shares *big.Int) (*big.Int, error) {
	return _ERC4626.Contract.PreviewRedeem(&_ERC4626.CallOpts, _shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 _shares) view returns(uint256)
func (_ERC4626 *ERC4626CallerSession) PreviewRedeem(_shares *big.Int) (*big.Int, error) {
	return _ERC4626.Contract.PreviewRedeem(&_ERC4626.CallOpts, _shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 _assets) view returns(uint256)
func (_ERC4626 *ERC4626Caller) PreviewWithdraw(opts *bind.CallOpts, _assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626.contract.Call(opts, &out, "previewWithdraw", _assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 _assets) view returns(uint256)
func (_ERC4626 *ERC4626Session) PreviewWithdraw(_assets *big.Int) (*big.Int, error) {
	return _ERC4626.Contract.PreviewWithdraw(&_ERC4626.CallOpts, _assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 _assets) view returns(uint256)
func (_ERC4626 *ERC4626CallerSession) PreviewWithdraw(_assets *big.Int) (*big.Int, error) {
	return _ERC4626.Contract.PreviewWithdraw(&_ERC4626.CallOpts, _assets)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC4626 *ERC4626Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC4626.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC4626 *ERC4626Session) Symbol() (string, error) {
	return _ERC4626.Contract.Symbol(&_ERC4626.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC4626 *ERC4626CallerSession) Symbol() (string, error) {
	return _ERC4626.Contract.Symbol(&_ERC4626.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_ERC4626 *ERC4626Caller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_ERC4626 *ERC4626Session) TotalAssets() (*big.Int, error) {
	return _ERC4626.Contract.TotalAssets(&_ERC4626.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_ERC4626 *ERC4626CallerSession) TotalAssets() (*big.Int, error) {
	return _ERC4626.Contract.TotalAssets(&_ERC4626.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC4626 *ERC4626Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC4626 *ERC4626Session) TotalSupply() (*big.Int, error) {
	return _ERC4626.Contract.TotalSupply(&_ERC4626.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC4626 *ERC4626CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC4626.Contract.TotalSupply(&_ERC4626.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_ERC4626 *ERC4626Caller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC4626.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_ERC4626 *ERC4626Session) Treasury() (common.Address, error) {
	return _ERC4626.Contract.Treasury(&_ERC4626.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_ERC4626 *ERC4626CallerSession) Treasury() (common.Address, error) {
	return _ERC4626.Contract.Treasury(&_ERC4626.CallOpts)
}

// Updated is a free data retrieval call binding the contract method 0x7b2aab03.
//
// Solidity: function updated() view returns(uint256)
func (_ERC4626 *ERC4626Caller) Updated(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626.contract.Call(opts, &out, "updated")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Updated is a free data retrieval call binding the contract method 0x7b2aab03.
//
// Solidity: function updated() view returns(uint256)
func (_ERC4626 *ERC4626Session) Updated() (*big.Int, error) {
	return _ERC4626.Contract.Updated(&_ERC4626.CallOpts)
}

// Updated is a free data retrieval call binding the contract method 0x7b2aab03.
//
// Solidity: function updated() view returns(uint256)
func (_ERC4626 *ERC4626CallerSession) Updated() (*big.Int, error) {
	return _ERC4626.Contract.Updated(&_ERC4626.CallOpts)
}

// VoteWeight is a free data retrieval call binding the contract method 0xf49ec310.
//
// Solidity: function vote_weight(address _account) view returns(uint256)
func (_ERC4626 *ERC4626Caller) VoteWeight(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626.contract.Call(opts, &out, "vote_weight", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VoteWeight is a free data retrieval call binding the contract method 0xf49ec310.
//
// Solidity: function vote_weight(address _account) view returns(uint256)
func (_ERC4626 *ERC4626Session) VoteWeight(_account common.Address) (*big.Int, error) {
	return _ERC4626.Contract.VoteWeight(&_ERC4626.CallOpts, _account)
}

// VoteWeight is a free data retrieval call binding the contract method 0xf49ec310.
//
// Solidity: function vote_weight(address _account) view returns(uint256)
func (_ERC4626 *ERC4626CallerSession) VoteWeight(_account common.Address) (*big.Int, error) {
	return _ERC4626.Contract.VoteWeight(&_ERC4626.CallOpts, _account)
}

// AcceptManagement is a paid mutator transaction binding the contract method 0x759be10c.
//
// Solidity: function accept_management() returns()
func (_ERC4626 *ERC4626Transactor) AcceptManagement(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC4626.contract.Transact(opts, "accept_management")
}

// AcceptManagement is a paid mutator transaction binding the contract method 0x759be10c.
//
// Solidity: function accept_management() returns()
func (_ERC4626 *ERC4626Session) AcceptManagement() (*types.Transaction, error) {
	return _ERC4626.Contract.AcceptManagement(&_ERC4626.TransactOpts)
}

// AcceptManagement is a paid mutator transaction binding the contract method 0x759be10c.
//
// Solidity: function accept_management() returns()
func (_ERC4626 *ERC4626TransactorSession) AcceptManagement() (*types.Transaction, error) {
	return _ERC4626.Contract.AcceptManagement(&_ERC4626.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_ERC4626 *ERC4626Transactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC4626.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_ERC4626 *ERC4626Session) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC4626.Contract.Approve(&_ERC4626.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_ERC4626 *ERC4626TransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC4626.Contract.Approve(&_ERC4626.TransactOpts, _spender, _value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _value) returns(bool)
func (_ERC4626 *ERC4626Transactor) DecreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC4626.contract.Transact(opts, "decreaseAllowance", _spender, _value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _value) returns(bool)
func (_ERC4626 *ERC4626Session) DecreaseAllowance(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC4626.Contract.DecreaseAllowance(&_ERC4626.TransactOpts, _spender, _value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _value) returns(bool)
func (_ERC4626 *ERC4626TransactorSession) DecreaseAllowance(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC4626.Contract.DecreaseAllowance(&_ERC4626.TransactOpts, _spender, _value)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _assets) returns(uint256)
func (_ERC4626 *ERC4626Transactor) Deposit(opts *bind.TransactOpts, _assets *big.Int) (*types.Transaction, error) {
	return _ERC4626.contract.Transact(opts, "deposit", _assets)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _assets) returns(uint256)
func (_ERC4626 *ERC4626Session) Deposit(_assets *big.Int) (*types.Transaction, error) {
	return _ERC4626.Contract.Deposit(&_ERC4626.TransactOpts, _assets)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _assets) returns(uint256)
func (_ERC4626 *ERC4626TransactorSession) Deposit(_assets *big.Int) (*types.Transaction, error) {
	return _ERC4626.Contract.Deposit(&_ERC4626.TransactOpts, _assets)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _assets, address _receiver) returns(uint256)
func (_ERC4626 *ERC4626Transactor) Deposit0(opts *bind.TransactOpts, _assets *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _ERC4626.contract.Transact(opts, "deposit0", _assets, _receiver)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _assets, address _receiver) returns(uint256)
func (_ERC4626 *ERC4626Session) Deposit0(_assets *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _ERC4626.Contract.Deposit0(&_ERC4626.TransactOpts, _assets, _receiver)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _assets, address _receiver) returns(uint256)
func (_ERC4626 *ERC4626TransactorSession) Deposit0(_assets *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _ERC4626.Contract.Deposit0(&_ERC4626.TransactOpts, _assets, _receiver)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _value) returns(bool)
func (_ERC4626 *ERC4626Transactor) IncreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC4626.contract.Transact(opts, "increaseAllowance", _spender, _value)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _value) returns(bool)
func (_ERC4626 *ERC4626Session) IncreaseAllowance(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC4626.Contract.IncreaseAllowance(&_ERC4626.TransactOpts, _spender, _value)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _value) returns(bool)
func (_ERC4626 *ERC4626TransactorSession) IncreaseAllowance(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC4626.Contract.IncreaseAllowance(&_ERC4626.TransactOpts, _spender, _value)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 _shares) returns(uint256)
func (_ERC4626 *ERC4626Transactor) Mint(opts *bind.TransactOpts, _shares *big.Int) (*types.Transaction, error) {
	return _ERC4626.contract.Transact(opts, "mint", _shares)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 _shares) returns(uint256)
func (_ERC4626 *ERC4626Session) Mint(_shares *big.Int) (*types.Transaction, error) {
	return _ERC4626.Contract.Mint(&_ERC4626.TransactOpts, _shares)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 _shares) returns(uint256)
func (_ERC4626 *ERC4626TransactorSession) Mint(_shares *big.Int) (*types.Transaction, error) {
	return _ERC4626.Contract.Mint(&_ERC4626.TransactOpts, _shares)
}

// Mint0 is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 _shares, address _receiver) returns(uint256)
func (_ERC4626 *ERC4626Transactor) Mint0(opts *bind.TransactOpts, _shares *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _ERC4626.contract.Transact(opts, "mint0", _shares, _receiver)
}

// Mint0 is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 _shares, address _receiver) returns(uint256)
func (_ERC4626 *ERC4626Session) Mint0(_shares *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _ERC4626.Contract.Mint0(&_ERC4626.TransactOpts, _shares, _receiver)
}

// Mint0 is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 _shares, address _receiver) returns(uint256)
func (_ERC4626 *ERC4626TransactorSession) Mint0(_shares *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _ERC4626.Contract.Mint0(&_ERC4626.TransactOpts, _shares, _receiver)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 _shares) returns(uint256)
func (_ERC4626 *ERC4626Transactor) Redeem(opts *bind.TransactOpts, _shares *big.Int) (*types.Transaction, error) {
	return _ERC4626.contract.Transact(opts, "redeem", _shares)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 _shares) returns(uint256)
func (_ERC4626 *ERC4626Session) Redeem(_shares *big.Int) (*types.Transaction, error) {
	return _ERC4626.Contract.Redeem(&_ERC4626.TransactOpts, _shares)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 _shares) returns(uint256)
func (_ERC4626 *ERC4626TransactorSession) Redeem(_shares *big.Int) (*types.Transaction, error) {
	return _ERC4626.Contract.Redeem(&_ERC4626.TransactOpts, _shares)
}

// Redeem0 is a paid mutator transaction binding the contract method 0x7bde82f2.
//
// Solidity: function redeem(uint256 _shares, address _receiver) returns(uint256)
func (_ERC4626 *ERC4626Transactor) Redeem0(opts *bind.TransactOpts, _shares *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _ERC4626.contract.Transact(opts, "redeem0", _shares, _receiver)
}

// Redeem0 is a paid mutator transaction binding the contract method 0x7bde82f2.
//
// Solidity: function redeem(uint256 _shares, address _receiver) returns(uint256)
func (_ERC4626 *ERC4626Session) Redeem0(_shares *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _ERC4626.Contract.Redeem0(&_ERC4626.TransactOpts, _shares, _receiver)
}

// Redeem0 is a paid mutator transaction binding the contract method 0x7bde82f2.
//
// Solidity: function redeem(uint256 _shares, address _receiver) returns(uint256)
func (_ERC4626 *ERC4626TransactorSession) Redeem0(_shares *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _ERC4626.Contract.Redeem0(&_ERC4626.TransactOpts, _shares, _receiver)
}

// Redeem1 is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 _shares, address _receiver, address _owner) returns(uint256)
func (_ERC4626 *ERC4626Transactor) Redeem1(opts *bind.TransactOpts, _shares *big.Int, _receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _ERC4626.contract.Transact(opts, "redeem1", _shares, _receiver, _owner)
}

// Redeem1 is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 _shares, address _receiver, address _owner) returns(uint256)
func (_ERC4626 *ERC4626Session) Redeem1(_shares *big.Int, _receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _ERC4626.Contract.Redeem1(&_ERC4626.TransactOpts, _shares, _receiver, _owner)
}

// Redeem1 is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 _shares, address _receiver, address _owner) returns(uint256)
func (_ERC4626 *ERC4626TransactorSession) Redeem1(_shares *big.Int, _receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _ERC4626.Contract.Redeem1(&_ERC4626.TransactOpts, _shares, _receiver, _owner)
}

// Rescue is a paid mutator transaction binding the contract method 0x4fdf5d1d.
//
// Solidity: function rescue(address _token, address _receiver) returns()
func (_ERC4626 *ERC4626Transactor) Rescue(opts *bind.TransactOpts, _token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _ERC4626.contract.Transact(opts, "rescue", _token, _receiver)
}

// Rescue is a paid mutator transaction binding the contract method 0x4fdf5d1d.
//
// Solidity: function rescue(address _token, address _receiver) returns()
func (_ERC4626 *ERC4626Session) Rescue(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _ERC4626.Contract.Rescue(&_ERC4626.TransactOpts, _token, _receiver)
}

// Rescue is a paid mutator transaction binding the contract method 0x4fdf5d1d.
//
// Solidity: function rescue(address _token, address _receiver) returns()
func (_ERC4626 *ERC4626TransactorSession) Rescue(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _ERC4626.Contract.Rescue(&_ERC4626.TransactOpts, _token, _receiver)
}

// SetHalfTime is a paid mutator transaction binding the contract method 0x1bf93f86.
//
// Solidity: function set_half_time(uint256 _half_time) returns()
func (_ERC4626 *ERC4626Transactor) SetHalfTime(opts *bind.TransactOpts, _half_time *big.Int) (*types.Transaction, error) {
	return _ERC4626.contract.Transact(opts, "set_half_time", _half_time)
}

// SetHalfTime is a paid mutator transaction binding the contract method 0x1bf93f86.
//
// Solidity: function set_half_time(uint256 _half_time) returns()
func (_ERC4626 *ERC4626Session) SetHalfTime(_half_time *big.Int) (*types.Transaction, error) {
	return _ERC4626.Contract.SetHalfTime(&_ERC4626.TransactOpts, _half_time)
}

// SetHalfTime is a paid mutator transaction binding the contract method 0x1bf93f86.
//
// Solidity: function set_half_time(uint256 _half_time) returns()
func (_ERC4626 *ERC4626TransactorSession) SetHalfTime(_half_time *big.Int) (*types.Transaction, error) {
	return _ERC4626.Contract.SetHalfTime(&_ERC4626.TransactOpts, _half_time)
}

// SetManagement is a paid mutator transaction binding the contract method 0xfd066ecc.
//
// Solidity: function set_management(address _management) returns()
func (_ERC4626 *ERC4626Transactor) SetManagement(opts *bind.TransactOpts, _management common.Address) (*types.Transaction, error) {
	return _ERC4626.contract.Transact(opts, "set_management", _management)
}

// SetManagement is a paid mutator transaction binding the contract method 0xfd066ecc.
//
// Solidity: function set_management(address _management) returns()
func (_ERC4626 *ERC4626Session) SetManagement(_management common.Address) (*types.Transaction, error) {
	return _ERC4626.Contract.SetManagement(&_ERC4626.TransactOpts, _management)
}

// SetManagement is a paid mutator transaction binding the contract method 0xfd066ecc.
//
// Solidity: function set_management(address _management) returns()
func (_ERC4626 *ERC4626TransactorSession) SetManagement(_management common.Address) (*types.Transaction, error) {
	return _ERC4626.Contract.SetManagement(&_ERC4626.TransactOpts, _management)
}

// SetPerformanceFeeRate is a paid mutator transaction binding the contract method 0x3047ce9d.
//
// Solidity: function set_performance_fee_rate(uint256 _fee_rate) returns()
func (_ERC4626 *ERC4626Transactor) SetPerformanceFeeRate(opts *bind.TransactOpts, _fee_rate *big.Int) (*types.Transaction, error) {
	return _ERC4626.contract.Transact(opts, "set_performance_fee_rate", _fee_rate)
}

// SetPerformanceFeeRate is a paid mutator transaction binding the contract method 0x3047ce9d.
//
// Solidity: function set_performance_fee_rate(uint256 _fee_rate) returns()
func (_ERC4626 *ERC4626Session) SetPerformanceFeeRate(_fee_rate *big.Int) (*types.Transaction, error) {
	return _ERC4626.Contract.SetPerformanceFeeRate(&_ERC4626.TransactOpts, _fee_rate)
}

// SetPerformanceFeeRate is a paid mutator transaction binding the contract method 0x3047ce9d.
//
// Solidity: function set_performance_fee_rate(uint256 _fee_rate) returns()
func (_ERC4626 *ERC4626TransactorSession) SetPerformanceFeeRate(_fee_rate *big.Int) (*types.Transaction, error) {
	return _ERC4626.Contract.SetPerformanceFeeRate(&_ERC4626.TransactOpts, _fee_rate)
}

// SetTreasury is a paid mutator transaction binding the contract method 0x30bcd67b.
//
// Solidity: function set_treasury(address _treasury) returns()
func (_ERC4626 *ERC4626Transactor) SetTreasury(opts *bind.TransactOpts, _treasury common.Address) (*types.Transaction, error) {
	return _ERC4626.contract.Transact(opts, "set_treasury", _treasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0x30bcd67b.
//
// Solidity: function set_treasury(address _treasury) returns()
func (_ERC4626 *ERC4626Session) SetTreasury(_treasury common.Address) (*types.Transaction, error) {
	return _ERC4626.Contract.SetTreasury(&_ERC4626.TransactOpts, _treasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0x30bcd67b.
//
// Solidity: function set_treasury(address _treasury) returns()
func (_ERC4626 *ERC4626TransactorSession) SetTreasury(_treasury common.Address) (*types.Transaction, error) {
	return _ERC4626.Contract.SetTreasury(&_ERC4626.TransactOpts, _treasury)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_ERC4626 *ERC4626Transactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC4626.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_ERC4626 *ERC4626Session) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC4626.Contract.Transfer(&_ERC4626.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_ERC4626 *ERC4626TransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC4626.Contract.Transfer(&_ERC4626.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_ERC4626 *ERC4626Transactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC4626.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_ERC4626 *ERC4626Session) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC4626.Contract.TransferFrom(&_ERC4626.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_ERC4626 *ERC4626TransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC4626.Contract.TransferFrom(&_ERC4626.TransactOpts, _from, _to, _value)
}

// UpdateAmounts is a paid mutator transaction binding the contract method 0x3ee352fc.
//
// Solidity: function update_amounts() returns(uint256, uint256, uint256)
func (_ERC4626 *ERC4626Transactor) UpdateAmounts(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC4626.contract.Transact(opts, "update_amounts")
}

// UpdateAmounts is a paid mutator transaction binding the contract method 0x3ee352fc.
//
// Solidity: function update_amounts() returns(uint256, uint256, uint256)
func (_ERC4626 *ERC4626Session) UpdateAmounts() (*types.Transaction, error) {
	return _ERC4626.Contract.UpdateAmounts(&_ERC4626.TransactOpts)
}

// UpdateAmounts is a paid mutator transaction binding the contract method 0x3ee352fc.
//
// Solidity: function update_amounts() returns(uint256, uint256, uint256)
func (_ERC4626 *ERC4626TransactorSession) UpdateAmounts() (*types.Transaction, error) {
	return _ERC4626.Contract.UpdateAmounts(&_ERC4626.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _assets) returns(uint256)
func (_ERC4626 *ERC4626Transactor) Withdraw(opts *bind.TransactOpts, _assets *big.Int) (*types.Transaction, error) {
	return _ERC4626.contract.Transact(opts, "withdraw", _assets)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _assets) returns(uint256)
func (_ERC4626 *ERC4626Session) Withdraw(_assets *big.Int) (*types.Transaction, error) {
	return _ERC4626.Contract.Withdraw(&_ERC4626.TransactOpts, _assets)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _assets) returns(uint256)
func (_ERC4626 *ERC4626TransactorSession) Withdraw(_assets *big.Int) (*types.Transaction, error) {
	return _ERC4626.Contract.Withdraw(&_ERC4626.TransactOpts, _assets)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 _assets, address _receiver) returns(uint256)
func (_ERC4626 *ERC4626Transactor) Withdraw0(opts *bind.TransactOpts, _assets *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _ERC4626.contract.Transact(opts, "withdraw0", _assets, _receiver)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 _assets, address _receiver) returns(uint256)
func (_ERC4626 *ERC4626Session) Withdraw0(_assets *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _ERC4626.Contract.Withdraw0(&_ERC4626.TransactOpts, _assets, _receiver)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 _assets, address _receiver) returns(uint256)
func (_ERC4626 *ERC4626TransactorSession) Withdraw0(_assets *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _ERC4626.Contract.Withdraw0(&_ERC4626.TransactOpts, _assets, _receiver)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 _assets, address _receiver, address _owner) returns(uint256)
func (_ERC4626 *ERC4626Transactor) Withdraw1(opts *bind.TransactOpts, _assets *big.Int, _receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _ERC4626.contract.Transact(opts, "withdraw1", _assets, _receiver, _owner)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 _assets, address _receiver, address _owner) returns(uint256)
func (_ERC4626 *ERC4626Session) Withdraw1(_assets *big.Int, _receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _ERC4626.Contract.Withdraw1(&_ERC4626.TransactOpts, _assets, _receiver, _owner)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 _assets, address _receiver, address _owner) returns(uint256)
func (_ERC4626 *ERC4626TransactorSession) Withdraw1(_assets *big.Int, _receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _ERC4626.Contract.Withdraw1(&_ERC4626.TransactOpts, _assets, _receiver, _owner)
}

// ERC4626ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC4626 contract.
type ERC4626ApprovalIterator struct {
	Event *ERC4626Approval // Event containing the contract specifics and raw log

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
func (it *ERC4626ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC4626Approval)
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
		it.Event = new(ERC4626Approval)
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
func (it *ERC4626ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC4626ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC4626Approval represents a Approval event raised by the ERC4626 contract.
type ERC4626Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC4626 *ERC4626Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC4626ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC4626.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC4626ApprovalIterator{contract: _ERC4626.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC4626 *ERC4626Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC4626Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC4626.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC4626Approval)
				if err := _ERC4626.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC4626 *ERC4626Filterer) ParseApproval(log types.Log) (*ERC4626Approval, error) {
	event := new(ERC4626Approval)
	if err := _ERC4626.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC4626DepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the ERC4626 contract.
type ERC4626DepositIterator struct {
	Event *ERC4626Deposit // Event containing the contract specifics and raw log

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
func (it *ERC4626DepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC4626Deposit)
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
		it.Event = new(ERC4626Deposit)
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
func (it *ERC4626DepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC4626DepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC4626Deposit represents a Deposit event raised by the ERC4626 contract.
type ERC4626Deposit struct {
	Sender common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_ERC4626 *ERC4626Filterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, owner []common.Address) (*ERC4626DepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ERC4626.contract.FilterLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ERC4626DepositIterator{contract: _ERC4626.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_ERC4626 *ERC4626Filterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ERC4626Deposit, sender []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ERC4626.contract.WatchLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC4626Deposit)
				if err := _ERC4626.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_ERC4626 *ERC4626Filterer) ParseDeposit(log types.Log) (*ERC4626Deposit, error) {
	event := new(ERC4626Deposit)
	if err := _ERC4626.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC4626PendingManagementIterator is returned from FilterPendingManagement and is used to iterate over the raw logs and unpacked data for PendingManagement events raised by the ERC4626 contract.
type ERC4626PendingManagementIterator struct {
	Event *ERC4626PendingManagement // Event containing the contract specifics and raw log

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
func (it *ERC4626PendingManagementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC4626PendingManagement)
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
		it.Event = new(ERC4626PendingManagement)
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
func (it *ERC4626PendingManagementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC4626PendingManagementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC4626PendingManagement represents a PendingManagement event raised by the ERC4626 contract.
type ERC4626PendingManagement struct {
	Management common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPendingManagement is a free log retrieval operation binding the contract event 0xe7b5cc087e6e47e33e86bdfe4720b7e849891938b18ff6e0c3f92299de79e60c.
//
// Solidity: event PendingManagement(address indexed management)
func (_ERC4626 *ERC4626Filterer) FilterPendingManagement(opts *bind.FilterOpts, management []common.Address) (*ERC4626PendingManagementIterator, error) {

	var managementRule []interface{}
	for _, managementItem := range management {
		managementRule = append(managementRule, managementItem)
	}

	logs, sub, err := _ERC4626.contract.FilterLogs(opts, "PendingManagement", managementRule)
	if err != nil {
		return nil, err
	}
	return &ERC4626PendingManagementIterator{contract: _ERC4626.contract, event: "PendingManagement", logs: logs, sub: sub}, nil
}

// WatchPendingManagement is a free log subscription operation binding the contract event 0xe7b5cc087e6e47e33e86bdfe4720b7e849891938b18ff6e0c3f92299de79e60c.
//
// Solidity: event PendingManagement(address indexed management)
func (_ERC4626 *ERC4626Filterer) WatchPendingManagement(opts *bind.WatchOpts, sink chan<- *ERC4626PendingManagement, management []common.Address) (event.Subscription, error) {

	var managementRule []interface{}
	for _, managementItem := range management {
		managementRule = append(managementRule, managementItem)
	}

	logs, sub, err := _ERC4626.contract.WatchLogs(opts, "PendingManagement", managementRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC4626PendingManagement)
				if err := _ERC4626.contract.UnpackLog(event, "PendingManagement", log); err != nil {
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

// ParsePendingManagement is a log parse operation binding the contract event 0xe7b5cc087e6e47e33e86bdfe4720b7e849891938b18ff6e0c3f92299de79e60c.
//
// Solidity: event PendingManagement(address indexed management)
func (_ERC4626 *ERC4626Filterer) ParsePendingManagement(log types.Log) (*ERC4626PendingManagement, error) {
	event := new(ERC4626PendingManagement)
	if err := _ERC4626.contract.UnpackLog(event, "PendingManagement", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC4626RewardsIterator is returned from FilterRewards and is used to iterate over the raw logs and unpacked data for Rewards events raised by the ERC4626 contract.
type ERC4626RewardsIterator struct {
	Event *ERC4626Rewards // Event containing the contract specifics and raw log

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
func (it *ERC4626RewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC4626Rewards)
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
		it.Event = new(ERC4626Rewards)
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
func (it *ERC4626RewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC4626RewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC4626Rewards represents a Rewards event raised by the ERC4626 contract.
type ERC4626Rewards struct {
	Pending   *big.Int
	Streaming *big.Int
	Unlocked  *big.Int
	Delta     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRewards is a free log retrieval operation binding the contract event 0xabf5b9a9abf080b67be5bfa614a827845afc366da18173be4db7bf2761f2517e.
//
// Solidity: event Rewards(uint256 pending, uint256 streaming, uint256 unlocked, int256 delta)
func (_ERC4626 *ERC4626Filterer) FilterRewards(opts *bind.FilterOpts) (*ERC4626RewardsIterator, error) {

	logs, sub, err := _ERC4626.contract.FilterLogs(opts, "Rewards")
	if err != nil {
		return nil, err
	}
	return &ERC4626RewardsIterator{contract: _ERC4626.contract, event: "Rewards", logs: logs, sub: sub}, nil
}

// WatchRewards is a free log subscription operation binding the contract event 0xabf5b9a9abf080b67be5bfa614a827845afc366da18173be4db7bf2761f2517e.
//
// Solidity: event Rewards(uint256 pending, uint256 streaming, uint256 unlocked, int256 delta)
func (_ERC4626 *ERC4626Filterer) WatchRewards(opts *bind.WatchOpts, sink chan<- *ERC4626Rewards) (event.Subscription, error) {

	logs, sub, err := _ERC4626.contract.WatchLogs(opts, "Rewards")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC4626Rewards)
				if err := _ERC4626.contract.UnpackLog(event, "Rewards", log); err != nil {
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

// ParseRewards is a log parse operation binding the contract event 0xabf5b9a9abf080b67be5bfa614a827845afc366da18173be4db7bf2761f2517e.
//
// Solidity: event Rewards(uint256 pending, uint256 streaming, uint256 unlocked, int256 delta)
func (_ERC4626 *ERC4626Filterer) ParseRewards(log types.Log) (*ERC4626Rewards, error) {
	event := new(ERC4626Rewards)
	if err := _ERC4626.contract.UnpackLog(event, "Rewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC4626SetFeeRateIterator is returned from FilterSetFeeRate and is used to iterate over the raw logs and unpacked data for SetFeeRate events raised by the ERC4626 contract.
type ERC4626SetFeeRateIterator struct {
	Event *ERC4626SetFeeRate // Event containing the contract specifics and raw log

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
func (it *ERC4626SetFeeRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC4626SetFeeRate)
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
		it.Event = new(ERC4626SetFeeRate)
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
func (it *ERC4626SetFeeRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC4626SetFeeRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC4626SetFeeRate represents a SetFeeRate event raised by the ERC4626 contract.
type ERC4626SetFeeRate struct {
	FeeRate *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetFeeRate is a free log retrieval operation binding the contract event 0x6717373928cccf59cc9912055cfa8db86e7085b95c94c15862b121114aa333be.
//
// Solidity: event SetFeeRate(uint256 fee_rate)
func (_ERC4626 *ERC4626Filterer) FilterSetFeeRate(opts *bind.FilterOpts) (*ERC4626SetFeeRateIterator, error) {

	logs, sub, err := _ERC4626.contract.FilterLogs(opts, "SetFeeRate")
	if err != nil {
		return nil, err
	}
	return &ERC4626SetFeeRateIterator{contract: _ERC4626.contract, event: "SetFeeRate", logs: logs, sub: sub}, nil
}

// WatchSetFeeRate is a free log subscription operation binding the contract event 0x6717373928cccf59cc9912055cfa8db86e7085b95c94c15862b121114aa333be.
//
// Solidity: event SetFeeRate(uint256 fee_rate)
func (_ERC4626 *ERC4626Filterer) WatchSetFeeRate(opts *bind.WatchOpts, sink chan<- *ERC4626SetFeeRate) (event.Subscription, error) {

	logs, sub, err := _ERC4626.contract.WatchLogs(opts, "SetFeeRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC4626SetFeeRate)
				if err := _ERC4626.contract.UnpackLog(event, "SetFeeRate", log); err != nil {
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

// ParseSetFeeRate is a log parse operation binding the contract event 0x6717373928cccf59cc9912055cfa8db86e7085b95c94c15862b121114aa333be.
//
// Solidity: event SetFeeRate(uint256 fee_rate)
func (_ERC4626 *ERC4626Filterer) ParseSetFeeRate(log types.Log) (*ERC4626SetFeeRate, error) {
	event := new(ERC4626SetFeeRate)
	if err := _ERC4626.contract.UnpackLog(event, "SetFeeRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC4626SetHalfTimeIterator is returned from FilterSetHalfTime and is used to iterate over the raw logs and unpacked data for SetHalfTime events raised by the ERC4626 contract.
type ERC4626SetHalfTimeIterator struct {
	Event *ERC4626SetHalfTime // Event containing the contract specifics and raw log

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
func (it *ERC4626SetHalfTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC4626SetHalfTime)
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
		it.Event = new(ERC4626SetHalfTime)
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
func (it *ERC4626SetHalfTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC4626SetHalfTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC4626SetHalfTime represents a SetHalfTime event raised by the ERC4626 contract.
type ERC4626SetHalfTime struct {
	HalfTime *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetHalfTime is a free log retrieval operation binding the contract event 0xa1eced997ca08ff4a574f4f14ce5dfcb37e8bde7f90a609f48f0c86edf1224ee.
//
// Solidity: event SetHalfTime(uint256 half_time)
func (_ERC4626 *ERC4626Filterer) FilterSetHalfTime(opts *bind.FilterOpts) (*ERC4626SetHalfTimeIterator, error) {

	logs, sub, err := _ERC4626.contract.FilterLogs(opts, "SetHalfTime")
	if err != nil {
		return nil, err
	}
	return &ERC4626SetHalfTimeIterator{contract: _ERC4626.contract, event: "SetHalfTime", logs: logs, sub: sub}, nil
}

// WatchSetHalfTime is a free log subscription operation binding the contract event 0xa1eced997ca08ff4a574f4f14ce5dfcb37e8bde7f90a609f48f0c86edf1224ee.
//
// Solidity: event SetHalfTime(uint256 half_time)
func (_ERC4626 *ERC4626Filterer) WatchSetHalfTime(opts *bind.WatchOpts, sink chan<- *ERC4626SetHalfTime) (event.Subscription, error) {

	logs, sub, err := _ERC4626.contract.WatchLogs(opts, "SetHalfTime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC4626SetHalfTime)
				if err := _ERC4626.contract.UnpackLog(event, "SetHalfTime", log); err != nil {
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

// ParseSetHalfTime is a log parse operation binding the contract event 0xa1eced997ca08ff4a574f4f14ce5dfcb37e8bde7f90a609f48f0c86edf1224ee.
//
// Solidity: event SetHalfTime(uint256 half_time)
func (_ERC4626 *ERC4626Filterer) ParseSetHalfTime(log types.Log) (*ERC4626SetHalfTime, error) {
	event := new(ERC4626SetHalfTime)
	if err := _ERC4626.contract.UnpackLog(event, "SetHalfTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC4626SetManagementIterator is returned from FilterSetManagement and is used to iterate over the raw logs and unpacked data for SetManagement events raised by the ERC4626 contract.
type ERC4626SetManagementIterator struct {
	Event *ERC4626SetManagement // Event containing the contract specifics and raw log

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
func (it *ERC4626SetManagementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC4626SetManagement)
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
		it.Event = new(ERC4626SetManagement)
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
func (it *ERC4626SetManagementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC4626SetManagementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC4626SetManagement represents a SetManagement event raised by the ERC4626 contract.
type ERC4626SetManagement struct {
	Management common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetManagement is a free log retrieval operation binding the contract event 0xafe23f9e1f603b288748a507d5a993957e9f14313a5889d5a070851299939d59.
//
// Solidity: event SetManagement(address indexed management)
func (_ERC4626 *ERC4626Filterer) FilterSetManagement(opts *bind.FilterOpts, management []common.Address) (*ERC4626SetManagementIterator, error) {

	var managementRule []interface{}
	for _, managementItem := range management {
		managementRule = append(managementRule, managementItem)
	}

	logs, sub, err := _ERC4626.contract.FilterLogs(opts, "SetManagement", managementRule)
	if err != nil {
		return nil, err
	}
	return &ERC4626SetManagementIterator{contract: _ERC4626.contract, event: "SetManagement", logs: logs, sub: sub}, nil
}

// WatchSetManagement is a free log subscription operation binding the contract event 0xafe23f9e1f603b288748a507d5a993957e9f14313a5889d5a070851299939d59.
//
// Solidity: event SetManagement(address indexed management)
func (_ERC4626 *ERC4626Filterer) WatchSetManagement(opts *bind.WatchOpts, sink chan<- *ERC4626SetManagement, management []common.Address) (event.Subscription, error) {

	var managementRule []interface{}
	for _, managementItem := range management {
		managementRule = append(managementRule, managementItem)
	}

	logs, sub, err := _ERC4626.contract.WatchLogs(opts, "SetManagement", managementRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC4626SetManagement)
				if err := _ERC4626.contract.UnpackLog(event, "SetManagement", log); err != nil {
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

// ParseSetManagement is a log parse operation binding the contract event 0xafe23f9e1f603b288748a507d5a993957e9f14313a5889d5a070851299939d59.
//
// Solidity: event SetManagement(address indexed management)
func (_ERC4626 *ERC4626Filterer) ParseSetManagement(log types.Log) (*ERC4626SetManagement, error) {
	event := new(ERC4626SetManagement)
	if err := _ERC4626.contract.UnpackLog(event, "SetManagement", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC4626SetTreasuryIterator is returned from FilterSetTreasury and is used to iterate over the raw logs and unpacked data for SetTreasury events raised by the ERC4626 contract.
type ERC4626SetTreasuryIterator struct {
	Event *ERC4626SetTreasury // Event containing the contract specifics and raw log

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
func (it *ERC4626SetTreasuryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC4626SetTreasury)
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
		it.Event = new(ERC4626SetTreasury)
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
func (it *ERC4626SetTreasuryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC4626SetTreasuryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC4626SetTreasury represents a SetTreasury event raised by the ERC4626 contract.
type ERC4626SetTreasury struct {
	Treasury common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetTreasury is a free log retrieval operation binding the contract event 0xcb7ef3e545f5cdb893f5c568ba710fe08f336375a2d9fd66e161033f8fc09ef3.
//
// Solidity: event SetTreasury(address indexed treasury)
func (_ERC4626 *ERC4626Filterer) FilterSetTreasury(opts *bind.FilterOpts, treasury []common.Address) (*ERC4626SetTreasuryIterator, error) {

	var treasuryRule []interface{}
	for _, treasuryItem := range treasury {
		treasuryRule = append(treasuryRule, treasuryItem)
	}

	logs, sub, err := _ERC4626.contract.FilterLogs(opts, "SetTreasury", treasuryRule)
	if err != nil {
		return nil, err
	}
	return &ERC4626SetTreasuryIterator{contract: _ERC4626.contract, event: "SetTreasury", logs: logs, sub: sub}, nil
}

// WatchSetTreasury is a free log subscription operation binding the contract event 0xcb7ef3e545f5cdb893f5c568ba710fe08f336375a2d9fd66e161033f8fc09ef3.
//
// Solidity: event SetTreasury(address indexed treasury)
func (_ERC4626 *ERC4626Filterer) WatchSetTreasury(opts *bind.WatchOpts, sink chan<- *ERC4626SetTreasury, treasury []common.Address) (event.Subscription, error) {

	var treasuryRule []interface{}
	for _, treasuryItem := range treasury {
		treasuryRule = append(treasuryRule, treasuryItem)
	}

	logs, sub, err := _ERC4626.contract.WatchLogs(opts, "SetTreasury", treasuryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC4626SetTreasury)
				if err := _ERC4626.contract.UnpackLog(event, "SetTreasury", log); err != nil {
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

// ParseSetTreasury is a log parse operation binding the contract event 0xcb7ef3e545f5cdb893f5c568ba710fe08f336375a2d9fd66e161033f8fc09ef3.
//
// Solidity: event SetTreasury(address indexed treasury)
func (_ERC4626 *ERC4626Filterer) ParseSetTreasury(log types.Log) (*ERC4626SetTreasury, error) {
	event := new(ERC4626SetTreasury)
	if err := _ERC4626.contract.UnpackLog(event, "SetTreasury", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC4626TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC4626 contract.
type ERC4626TransferIterator struct {
	Event *ERC4626Transfer // Event containing the contract specifics and raw log

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
func (it *ERC4626TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC4626Transfer)
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
		it.Event = new(ERC4626Transfer)
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
func (it *ERC4626TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC4626TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC4626Transfer represents a Transfer event raised by the ERC4626 contract.
type ERC4626Transfer struct {
	Sender   common.Address
	Receiver common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_ERC4626 *ERC4626Filterer) FilterTransfer(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*ERC4626TransferIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ERC4626.contract.FilterLogs(opts, "Transfer", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ERC4626TransferIterator{contract: _ERC4626.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_ERC4626 *ERC4626Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC4626Transfer, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ERC4626.contract.WatchLogs(opts, "Transfer", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC4626Transfer)
				if err := _ERC4626.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_ERC4626 *ERC4626Filterer) ParseTransfer(log types.Log) (*ERC4626Transfer, error) {
	event := new(ERC4626Transfer)
	if err := _ERC4626.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC4626WithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the ERC4626 contract.
type ERC4626WithdrawIterator struct {
	Event *ERC4626Withdraw // Event containing the contract specifics and raw log

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
func (it *ERC4626WithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC4626Withdraw)
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
		it.Event = new(ERC4626Withdraw)
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
func (it *ERC4626WithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC4626WithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC4626Withdraw represents a Withdraw event raised by the ERC4626 contract.
type ERC4626Withdraw struct {
	Sender   common.Address
	Receiver common.Address
	Owner    common.Address
	Assets   *big.Int
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_ERC4626 *ERC4626Filterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address, owner []common.Address) (*ERC4626WithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ERC4626.contract.FilterLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ERC4626WithdrawIterator{contract: _ERC4626.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_ERC4626 *ERC4626Filterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *ERC4626Withdraw, sender []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ERC4626.contract.WatchLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC4626Withdraw)
				if err := _ERC4626.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_ERC4626 *ERC4626Filterer) ParseWithdraw(log types.Log) (*ERC4626Withdraw, error) {
	event := new(ERC4626Withdraw)
	if err := _ERC4626.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
