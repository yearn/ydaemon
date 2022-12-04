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

// ATokenV1MetaData contains all meta data concerning the ATokenV1 contract.
var ATokenV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractLendingPoolAddressesProvider\",\"name\":\"_addressesProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_underlyingAsset\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_underlyingAssetDecimals\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromBalanceIncrease\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toBalanceIncrease\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toIndex\",\"type\":\"uint256\"}],\"name\":\"BalanceTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromBalanceIncrease\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromIndex\",\"type\":\"uint256\"}],\"name\":\"BurnOnLiquidation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"InterestRedirectionAllowanceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_redirectedBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromBalanceIncrease\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromIndex\",\"type\":\"uint256\"}],\"name\":\"InterestStreamRedirected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromBalanceIncrease\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromIndex\",\"type\":\"uint256\"}],\"name\":\"MintOnDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromBalanceIncrease\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromIndex\",\"type\":\"uint256\"}],\"name\":\"Redeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_targetAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_targetBalanceIncrease\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_targetIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_redirectedBalanceAdded\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_redirectedBalanceRemoved\",\"type\":\"uint256\"}],\"name\":\"RedirectedBalanceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"UINT_MAX_VALUE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"allowInterestRedirectionTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burnOnLiquidation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getInterestRedirectionAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getRedirectedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"isTransferAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mintOnDeposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"principalBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"redirectInterestStream\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"redirectInterestStreamOf\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferOnLiquidation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"underlyingAssetAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ATokenV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use ATokenV1MetaData.ABI instead.
var ATokenV1ABI = ATokenV1MetaData.ABI

// ATokenV1 is an auto generated Go binding around an Ethereum contract.
type ATokenV1 struct {
	ATokenV1Caller     // Read-only binding to the contract
	ATokenV1Transactor // Write-only binding to the contract
	ATokenV1Filterer   // Log filterer for contract events
}

// ATokenV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type ATokenV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ATokenV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ATokenV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ATokenV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ATokenV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ATokenV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ATokenV1Session struct {
	Contract     *ATokenV1         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ATokenV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ATokenV1CallerSession struct {
	Contract *ATokenV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ATokenV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ATokenV1TransactorSession struct {
	Contract     *ATokenV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ATokenV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type ATokenV1Raw struct {
	Contract *ATokenV1 // Generic contract binding to access the raw methods on
}

// ATokenV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ATokenV1CallerRaw struct {
	Contract *ATokenV1Caller // Generic read-only contract binding to access the raw methods on
}

// ATokenV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ATokenV1TransactorRaw struct {
	Contract *ATokenV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewATokenV1 creates a new instance of ATokenV1, bound to a specific deployed contract.
func NewATokenV1(address common.Address, backend bind.ContractBackend) (*ATokenV1, error) {
	contract, err := bindATokenV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ATokenV1{ATokenV1Caller: ATokenV1Caller{contract: contract}, ATokenV1Transactor: ATokenV1Transactor{contract: contract}, ATokenV1Filterer: ATokenV1Filterer{contract: contract}}, nil
}

// NewATokenV1Caller creates a new read-only instance of ATokenV1, bound to a specific deployed contract.
func NewATokenV1Caller(address common.Address, caller bind.ContractCaller) (*ATokenV1Caller, error) {
	contract, err := bindATokenV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ATokenV1Caller{contract: contract}, nil
}

// NewATokenV1Transactor creates a new write-only instance of ATokenV1, bound to a specific deployed contract.
func NewATokenV1Transactor(address common.Address, transactor bind.ContractTransactor) (*ATokenV1Transactor, error) {
	contract, err := bindATokenV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ATokenV1Transactor{contract: contract}, nil
}

// NewATokenV1Filterer creates a new log filterer instance of ATokenV1, bound to a specific deployed contract.
func NewATokenV1Filterer(address common.Address, filterer bind.ContractFilterer) (*ATokenV1Filterer, error) {
	contract, err := bindATokenV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ATokenV1Filterer{contract: contract}, nil
}

// bindATokenV1 binds a generic wrapper to an already deployed contract.
func bindATokenV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ATokenV1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ATokenV1 *ATokenV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ATokenV1.Contract.ATokenV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ATokenV1 *ATokenV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ATokenV1.Contract.ATokenV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ATokenV1 *ATokenV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ATokenV1.Contract.ATokenV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ATokenV1 *ATokenV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ATokenV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ATokenV1 *ATokenV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ATokenV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ATokenV1 *ATokenV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ATokenV1.Contract.contract.Transact(opts, method, params...)
}

// UINTMAXVALUE is a free data retrieval call binding the contract method 0xd0fc81d2.
//
// Solidity: function UINT_MAX_VALUE() view returns(uint256)
func (_ATokenV1 *ATokenV1Caller) UINTMAXVALUE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ATokenV1.contract.Call(opts, &out, "UINT_MAX_VALUE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UINTMAXVALUE is a free data retrieval call binding the contract method 0xd0fc81d2.
//
// Solidity: function UINT_MAX_VALUE() view returns(uint256)
func (_ATokenV1 *ATokenV1Session) UINTMAXVALUE() (*big.Int, error) {
	return _ATokenV1.Contract.UINTMAXVALUE(&_ATokenV1.CallOpts)
}

// UINTMAXVALUE is a free data retrieval call binding the contract method 0xd0fc81d2.
//
// Solidity: function UINT_MAX_VALUE() view returns(uint256)
func (_ATokenV1 *ATokenV1CallerSession) UINTMAXVALUE() (*big.Int, error) {
	return _ATokenV1.Contract.UINTMAXVALUE(&_ATokenV1.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ATokenV1 *ATokenV1Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ATokenV1.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ATokenV1 *ATokenV1Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ATokenV1.Contract.Allowance(&_ATokenV1.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ATokenV1 *ATokenV1CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ATokenV1.Contract.Allowance(&_ATokenV1.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_ATokenV1 *ATokenV1Caller) BalanceOf(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ATokenV1.contract.Call(opts, &out, "balanceOf", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_ATokenV1 *ATokenV1Session) BalanceOf(_user common.Address) (*big.Int, error) {
	return _ATokenV1.Contract.BalanceOf(&_ATokenV1.CallOpts, _user)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_ATokenV1 *ATokenV1CallerSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _ATokenV1.Contract.BalanceOf(&_ATokenV1.CallOpts, _user)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ATokenV1 *ATokenV1Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ATokenV1.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ATokenV1 *ATokenV1Session) Decimals() (uint8, error) {
	return _ATokenV1.Contract.Decimals(&_ATokenV1.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ATokenV1 *ATokenV1CallerSession) Decimals() (uint8, error) {
	return _ATokenV1.Contract.Decimals(&_ATokenV1.CallOpts)
}

// GetInterestRedirectionAddress is a free data retrieval call binding the contract method 0x445e8010.
//
// Solidity: function getInterestRedirectionAddress(address _user) view returns(address)
func (_ATokenV1 *ATokenV1Caller) GetInterestRedirectionAddress(opts *bind.CallOpts, _user common.Address) (common.Address, error) {
	var out []interface{}
	err := _ATokenV1.contract.Call(opts, &out, "getInterestRedirectionAddress", _user)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetInterestRedirectionAddress is a free data retrieval call binding the contract method 0x445e8010.
//
// Solidity: function getInterestRedirectionAddress(address _user) view returns(address)
func (_ATokenV1 *ATokenV1Session) GetInterestRedirectionAddress(_user common.Address) (common.Address, error) {
	return _ATokenV1.Contract.GetInterestRedirectionAddress(&_ATokenV1.CallOpts, _user)
}

// GetInterestRedirectionAddress is a free data retrieval call binding the contract method 0x445e8010.
//
// Solidity: function getInterestRedirectionAddress(address _user) view returns(address)
func (_ATokenV1 *ATokenV1CallerSession) GetInterestRedirectionAddress(_user common.Address) (common.Address, error) {
	return _ATokenV1.Contract.GetInterestRedirectionAddress(&_ATokenV1.CallOpts, _user)
}

// GetRedirectedBalance is a free data retrieval call binding the contract method 0x1d51e7cf.
//
// Solidity: function getRedirectedBalance(address _user) view returns(uint256)
func (_ATokenV1 *ATokenV1Caller) GetRedirectedBalance(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ATokenV1.contract.Call(opts, &out, "getRedirectedBalance", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRedirectedBalance is a free data retrieval call binding the contract method 0x1d51e7cf.
//
// Solidity: function getRedirectedBalance(address _user) view returns(uint256)
func (_ATokenV1 *ATokenV1Session) GetRedirectedBalance(_user common.Address) (*big.Int, error) {
	return _ATokenV1.Contract.GetRedirectedBalance(&_ATokenV1.CallOpts, _user)
}

// GetRedirectedBalance is a free data retrieval call binding the contract method 0x1d51e7cf.
//
// Solidity: function getRedirectedBalance(address _user) view returns(uint256)
func (_ATokenV1 *ATokenV1CallerSession) GetRedirectedBalance(_user common.Address) (*big.Int, error) {
	return _ATokenV1.Contract.GetRedirectedBalance(&_ATokenV1.CallOpts, _user)
}

// GetUserIndex is a free data retrieval call binding the contract method 0xee9907a4.
//
// Solidity: function getUserIndex(address _user) view returns(uint256)
func (_ATokenV1 *ATokenV1Caller) GetUserIndex(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ATokenV1.contract.Call(opts, &out, "getUserIndex", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserIndex is a free data retrieval call binding the contract method 0xee9907a4.
//
// Solidity: function getUserIndex(address _user) view returns(uint256)
func (_ATokenV1 *ATokenV1Session) GetUserIndex(_user common.Address) (*big.Int, error) {
	return _ATokenV1.Contract.GetUserIndex(&_ATokenV1.CallOpts, _user)
}

// GetUserIndex is a free data retrieval call binding the contract method 0xee9907a4.
//
// Solidity: function getUserIndex(address _user) view returns(uint256)
func (_ATokenV1 *ATokenV1CallerSession) GetUserIndex(_user common.Address) (*big.Int, error) {
	return _ATokenV1.Contract.GetUserIndex(&_ATokenV1.CallOpts, _user)
}

// IsTransferAllowed is a free data retrieval call binding the contract method 0x5eae177c.
//
// Solidity: function isTransferAllowed(address _user, uint256 _amount) view returns(bool)
func (_ATokenV1 *ATokenV1Caller) IsTransferAllowed(opts *bind.CallOpts, _user common.Address, _amount *big.Int) (bool, error) {
	var out []interface{}
	err := _ATokenV1.contract.Call(opts, &out, "isTransferAllowed", _user, _amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTransferAllowed is a free data retrieval call binding the contract method 0x5eae177c.
//
// Solidity: function isTransferAllowed(address _user, uint256 _amount) view returns(bool)
func (_ATokenV1 *ATokenV1Session) IsTransferAllowed(_user common.Address, _amount *big.Int) (bool, error) {
	return _ATokenV1.Contract.IsTransferAllowed(&_ATokenV1.CallOpts, _user, _amount)
}

// IsTransferAllowed is a free data retrieval call binding the contract method 0x5eae177c.
//
// Solidity: function isTransferAllowed(address _user, uint256 _amount) view returns(bool)
func (_ATokenV1 *ATokenV1CallerSession) IsTransferAllowed(_user common.Address, _amount *big.Int) (bool, error) {
	return _ATokenV1.Contract.IsTransferAllowed(&_ATokenV1.CallOpts, _user, _amount)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ATokenV1 *ATokenV1Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ATokenV1.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ATokenV1 *ATokenV1Session) Name() (string, error) {
	return _ATokenV1.Contract.Name(&_ATokenV1.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ATokenV1 *ATokenV1CallerSession) Name() (string, error) {
	return _ATokenV1.Contract.Name(&_ATokenV1.CallOpts)
}

// PrincipalBalanceOf is a free data retrieval call binding the contract method 0xc634dfaa.
//
// Solidity: function principalBalanceOf(address _user) view returns(uint256)
func (_ATokenV1 *ATokenV1Caller) PrincipalBalanceOf(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ATokenV1.contract.Call(opts, &out, "principalBalanceOf", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PrincipalBalanceOf is a free data retrieval call binding the contract method 0xc634dfaa.
//
// Solidity: function principalBalanceOf(address _user) view returns(uint256)
func (_ATokenV1 *ATokenV1Session) PrincipalBalanceOf(_user common.Address) (*big.Int, error) {
	return _ATokenV1.Contract.PrincipalBalanceOf(&_ATokenV1.CallOpts, _user)
}

// PrincipalBalanceOf is a free data retrieval call binding the contract method 0xc634dfaa.
//
// Solidity: function principalBalanceOf(address _user) view returns(uint256)
func (_ATokenV1 *ATokenV1CallerSession) PrincipalBalanceOf(_user common.Address) (*big.Int, error) {
	return _ATokenV1.Contract.PrincipalBalanceOf(&_ATokenV1.CallOpts, _user)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ATokenV1 *ATokenV1Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ATokenV1.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ATokenV1 *ATokenV1Session) Symbol() (string, error) {
	return _ATokenV1.Contract.Symbol(&_ATokenV1.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ATokenV1 *ATokenV1CallerSession) Symbol() (string, error) {
	return _ATokenV1.Contract.Symbol(&_ATokenV1.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ATokenV1 *ATokenV1Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ATokenV1.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ATokenV1 *ATokenV1Session) TotalSupply() (*big.Int, error) {
	return _ATokenV1.Contract.TotalSupply(&_ATokenV1.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ATokenV1 *ATokenV1CallerSession) TotalSupply() (*big.Int, error) {
	return _ATokenV1.Contract.TotalSupply(&_ATokenV1.CallOpts)
}

// UnderlyingAssetAddress is a free data retrieval call binding the contract method 0x89d1a0fc.
//
// Solidity: function underlyingAssetAddress() view returns(address)
func (_ATokenV1 *ATokenV1Caller) UnderlyingAssetAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ATokenV1.contract.Call(opts, &out, "underlyingAssetAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingAssetAddress is a free data retrieval call binding the contract method 0x89d1a0fc.
//
// Solidity: function underlyingAssetAddress() view returns(address)
func (_ATokenV1 *ATokenV1Session) UnderlyingAssetAddress() (common.Address, error) {
	return _ATokenV1.Contract.UnderlyingAssetAddress(&_ATokenV1.CallOpts)
}

// UnderlyingAssetAddress is a free data retrieval call binding the contract method 0x89d1a0fc.
//
// Solidity: function underlyingAssetAddress() view returns(address)
func (_ATokenV1 *ATokenV1CallerSession) UnderlyingAssetAddress() (common.Address, error) {
	return _ATokenV1.Contract.UnderlyingAssetAddress(&_ATokenV1.CallOpts)
}

// AllowInterestRedirectionTo is a paid mutator transaction binding the contract method 0x12c87c2d.
//
// Solidity: function allowInterestRedirectionTo(address _to) returns()
func (_ATokenV1 *ATokenV1Transactor) AllowInterestRedirectionTo(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _ATokenV1.contract.Transact(opts, "allowInterestRedirectionTo", _to)
}

// AllowInterestRedirectionTo is a paid mutator transaction binding the contract method 0x12c87c2d.
//
// Solidity: function allowInterestRedirectionTo(address _to) returns()
func (_ATokenV1 *ATokenV1Session) AllowInterestRedirectionTo(_to common.Address) (*types.Transaction, error) {
	return _ATokenV1.Contract.AllowInterestRedirectionTo(&_ATokenV1.TransactOpts, _to)
}

// AllowInterestRedirectionTo is a paid mutator transaction binding the contract method 0x12c87c2d.
//
// Solidity: function allowInterestRedirectionTo(address _to) returns()
func (_ATokenV1 *ATokenV1TransactorSession) AllowInterestRedirectionTo(_to common.Address) (*types.Transaction, error) {
	return _ATokenV1.Contract.AllowInterestRedirectionTo(&_ATokenV1.TransactOpts, _to)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ATokenV1 *ATokenV1Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ATokenV1.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ATokenV1 *ATokenV1Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ATokenV1.Contract.Approve(&_ATokenV1.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ATokenV1 *ATokenV1TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ATokenV1.Contract.Approve(&_ATokenV1.TransactOpts, spender, value)
}

// BurnOnLiquidation is a paid mutator transaction binding the contract method 0x3edb7cb8.
//
// Solidity: function burnOnLiquidation(address _account, uint256 _value) returns()
func (_ATokenV1 *ATokenV1Transactor) BurnOnLiquidation(opts *bind.TransactOpts, _account common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ATokenV1.contract.Transact(opts, "burnOnLiquidation", _account, _value)
}

// BurnOnLiquidation is a paid mutator transaction binding the contract method 0x3edb7cb8.
//
// Solidity: function burnOnLiquidation(address _account, uint256 _value) returns()
func (_ATokenV1 *ATokenV1Session) BurnOnLiquidation(_account common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ATokenV1.Contract.BurnOnLiquidation(&_ATokenV1.TransactOpts, _account, _value)
}

// BurnOnLiquidation is a paid mutator transaction binding the contract method 0x3edb7cb8.
//
// Solidity: function burnOnLiquidation(address _account, uint256 _value) returns()
func (_ATokenV1 *ATokenV1TransactorSession) BurnOnLiquidation(_account common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ATokenV1.Contract.BurnOnLiquidation(&_ATokenV1.TransactOpts, _account, _value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ATokenV1 *ATokenV1Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ATokenV1.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ATokenV1 *ATokenV1Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ATokenV1.Contract.DecreaseAllowance(&_ATokenV1.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ATokenV1 *ATokenV1TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ATokenV1.Contract.DecreaseAllowance(&_ATokenV1.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ATokenV1 *ATokenV1Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ATokenV1.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ATokenV1 *ATokenV1Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ATokenV1.Contract.IncreaseAllowance(&_ATokenV1.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ATokenV1 *ATokenV1TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ATokenV1.Contract.IncreaseAllowance(&_ATokenV1.TransactOpts, spender, addedValue)
}

// MintOnDeposit is a paid mutator transaction binding the contract method 0x94362e8b.
//
// Solidity: function mintOnDeposit(address _account, uint256 _amount) returns()
func (_ATokenV1 *ATokenV1Transactor) MintOnDeposit(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ATokenV1.contract.Transact(opts, "mintOnDeposit", _account, _amount)
}

// MintOnDeposit is a paid mutator transaction binding the contract method 0x94362e8b.
//
// Solidity: function mintOnDeposit(address _account, uint256 _amount) returns()
func (_ATokenV1 *ATokenV1Session) MintOnDeposit(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ATokenV1.Contract.MintOnDeposit(&_ATokenV1.TransactOpts, _account, _amount)
}

// MintOnDeposit is a paid mutator transaction binding the contract method 0x94362e8b.
//
// Solidity: function mintOnDeposit(address _account, uint256 _amount) returns()
func (_ATokenV1 *ATokenV1TransactorSession) MintOnDeposit(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ATokenV1.Contract.MintOnDeposit(&_ATokenV1.TransactOpts, _account, _amount)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 _amount) returns()
func (_ATokenV1 *ATokenV1Transactor) Redeem(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _ATokenV1.contract.Transact(opts, "redeem", _amount)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 _amount) returns()
func (_ATokenV1 *ATokenV1Session) Redeem(_amount *big.Int) (*types.Transaction, error) {
	return _ATokenV1.Contract.Redeem(&_ATokenV1.TransactOpts, _amount)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 _amount) returns()
func (_ATokenV1 *ATokenV1TransactorSession) Redeem(_amount *big.Int) (*types.Transaction, error) {
	return _ATokenV1.Contract.Redeem(&_ATokenV1.TransactOpts, _amount)
}

// RedirectInterestStream is a paid mutator transaction binding the contract method 0x0e49072d.
//
// Solidity: function redirectInterestStream(address _to) returns()
func (_ATokenV1 *ATokenV1Transactor) RedirectInterestStream(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _ATokenV1.contract.Transact(opts, "redirectInterestStream", _to)
}

// RedirectInterestStream is a paid mutator transaction binding the contract method 0x0e49072d.
//
// Solidity: function redirectInterestStream(address _to) returns()
func (_ATokenV1 *ATokenV1Session) RedirectInterestStream(_to common.Address) (*types.Transaction, error) {
	return _ATokenV1.Contract.RedirectInterestStream(&_ATokenV1.TransactOpts, _to)
}

// RedirectInterestStream is a paid mutator transaction binding the contract method 0x0e49072d.
//
// Solidity: function redirectInterestStream(address _to) returns()
func (_ATokenV1 *ATokenV1TransactorSession) RedirectInterestStream(_to common.Address) (*types.Transaction, error) {
	return _ATokenV1.Contract.RedirectInterestStream(&_ATokenV1.TransactOpts, _to)
}

// RedirectInterestStreamOf is a paid mutator transaction binding the contract method 0x325a9b13.
//
// Solidity: function redirectInterestStreamOf(address _from, address _to) returns()
func (_ATokenV1 *ATokenV1Transactor) RedirectInterestStreamOf(opts *bind.TransactOpts, _from common.Address, _to common.Address) (*types.Transaction, error) {
	return _ATokenV1.contract.Transact(opts, "redirectInterestStreamOf", _from, _to)
}

// RedirectInterestStreamOf is a paid mutator transaction binding the contract method 0x325a9b13.
//
// Solidity: function redirectInterestStreamOf(address _from, address _to) returns()
func (_ATokenV1 *ATokenV1Session) RedirectInterestStreamOf(_from common.Address, _to common.Address) (*types.Transaction, error) {
	return _ATokenV1.Contract.RedirectInterestStreamOf(&_ATokenV1.TransactOpts, _from, _to)
}

// RedirectInterestStreamOf is a paid mutator transaction binding the contract method 0x325a9b13.
//
// Solidity: function redirectInterestStreamOf(address _from, address _to) returns()
func (_ATokenV1 *ATokenV1TransactorSession) RedirectInterestStreamOf(_from common.Address, _to common.Address) (*types.Transaction, error) {
	return _ATokenV1.Contract.RedirectInterestStreamOf(&_ATokenV1.TransactOpts, _from, _to)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ATokenV1 *ATokenV1Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ATokenV1.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ATokenV1 *ATokenV1Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ATokenV1.Contract.Transfer(&_ATokenV1.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ATokenV1 *ATokenV1TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ATokenV1.Contract.Transfer(&_ATokenV1.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ATokenV1 *ATokenV1Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ATokenV1.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ATokenV1 *ATokenV1Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ATokenV1.Contract.TransferFrom(&_ATokenV1.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ATokenV1 *ATokenV1TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ATokenV1.Contract.TransferFrom(&_ATokenV1.TransactOpts, sender, recipient, amount)
}

// TransferOnLiquidation is a paid mutator transaction binding the contract method 0xf866c319.
//
// Solidity: function transferOnLiquidation(address _from, address _to, uint256 _value) returns()
func (_ATokenV1 *ATokenV1Transactor) TransferOnLiquidation(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ATokenV1.contract.Transact(opts, "transferOnLiquidation", _from, _to, _value)
}

// TransferOnLiquidation is a paid mutator transaction binding the contract method 0xf866c319.
//
// Solidity: function transferOnLiquidation(address _from, address _to, uint256 _value) returns()
func (_ATokenV1 *ATokenV1Session) TransferOnLiquidation(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ATokenV1.Contract.TransferOnLiquidation(&_ATokenV1.TransactOpts, _from, _to, _value)
}

// TransferOnLiquidation is a paid mutator transaction binding the contract method 0xf866c319.
//
// Solidity: function transferOnLiquidation(address _from, address _to, uint256 _value) returns()
func (_ATokenV1 *ATokenV1TransactorSession) TransferOnLiquidation(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ATokenV1.Contract.TransferOnLiquidation(&_ATokenV1.TransactOpts, _from, _to, _value)
}

// ATokenV1ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ATokenV1 contract.
type ATokenV1ApprovalIterator struct {
	Event *ATokenV1Approval // Event containing the contract specifics and raw log

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
func (it *ATokenV1ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ATokenV1Approval)
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
		it.Event = new(ATokenV1Approval)
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
func (it *ATokenV1ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ATokenV1ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ATokenV1Approval represents a Approval event raised by the ATokenV1 contract.
type ATokenV1Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ATokenV1 *ATokenV1Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ATokenV1ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ATokenV1.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ATokenV1ApprovalIterator{contract: _ATokenV1.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ATokenV1 *ATokenV1Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ATokenV1Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ATokenV1.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ATokenV1Approval)
				if err := _ATokenV1.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ATokenV1 *ATokenV1Filterer) ParseApproval(log types.Log) (*ATokenV1Approval, error) {
	event := new(ATokenV1Approval)
	if err := _ATokenV1.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ATokenV1BalanceTransferIterator is returned from FilterBalanceTransfer and is used to iterate over the raw logs and unpacked data for BalanceTransfer events raised by the ATokenV1 contract.
type ATokenV1BalanceTransferIterator struct {
	Event *ATokenV1BalanceTransfer // Event containing the contract specifics and raw log

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
func (it *ATokenV1BalanceTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ATokenV1BalanceTransfer)
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
		it.Event = new(ATokenV1BalanceTransfer)
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
func (it *ATokenV1BalanceTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ATokenV1BalanceTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ATokenV1BalanceTransfer represents a BalanceTransfer event raised by the ATokenV1 contract.
type ATokenV1BalanceTransfer struct {
	From                common.Address
	To                  common.Address
	Value               *big.Int
	FromBalanceIncrease *big.Int
	ToBalanceIncrease   *big.Int
	FromIndex           *big.Int
	ToIndex             *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterBalanceTransfer is a free log retrieval operation binding the contract event 0x89a178eaa27e0cd201bd795ca8ff716ac0df9618494510ca79771cfc66ffcde8.
//
// Solidity: event BalanceTransfer(address indexed _from, address indexed _to, uint256 _value, uint256 _fromBalanceIncrease, uint256 _toBalanceIncrease, uint256 _fromIndex, uint256 _toIndex)
func (_ATokenV1 *ATokenV1Filterer) FilterBalanceTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*ATokenV1BalanceTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ATokenV1.contract.FilterLogs(opts, "BalanceTransfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &ATokenV1BalanceTransferIterator{contract: _ATokenV1.contract, event: "BalanceTransfer", logs: logs, sub: sub}, nil
}

// WatchBalanceTransfer is a free log subscription operation binding the contract event 0x89a178eaa27e0cd201bd795ca8ff716ac0df9618494510ca79771cfc66ffcde8.
//
// Solidity: event BalanceTransfer(address indexed _from, address indexed _to, uint256 _value, uint256 _fromBalanceIncrease, uint256 _toBalanceIncrease, uint256 _fromIndex, uint256 _toIndex)
func (_ATokenV1 *ATokenV1Filterer) WatchBalanceTransfer(opts *bind.WatchOpts, sink chan<- *ATokenV1BalanceTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ATokenV1.contract.WatchLogs(opts, "BalanceTransfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ATokenV1BalanceTransfer)
				if err := _ATokenV1.contract.UnpackLog(event, "BalanceTransfer", log); err != nil {
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

// ParseBalanceTransfer is a log parse operation binding the contract event 0x89a178eaa27e0cd201bd795ca8ff716ac0df9618494510ca79771cfc66ffcde8.
//
// Solidity: event BalanceTransfer(address indexed _from, address indexed _to, uint256 _value, uint256 _fromBalanceIncrease, uint256 _toBalanceIncrease, uint256 _fromIndex, uint256 _toIndex)
func (_ATokenV1 *ATokenV1Filterer) ParseBalanceTransfer(log types.Log) (*ATokenV1BalanceTransfer, error) {
	event := new(ATokenV1BalanceTransfer)
	if err := _ATokenV1.contract.UnpackLog(event, "BalanceTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ATokenV1BurnOnLiquidationIterator is returned from FilterBurnOnLiquidation and is used to iterate over the raw logs and unpacked data for BurnOnLiquidation events raised by the ATokenV1 contract.
type ATokenV1BurnOnLiquidationIterator struct {
	Event *ATokenV1BurnOnLiquidation // Event containing the contract specifics and raw log

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
func (it *ATokenV1BurnOnLiquidationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ATokenV1BurnOnLiquidation)
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
		it.Event = new(ATokenV1BurnOnLiquidation)
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
func (it *ATokenV1BurnOnLiquidationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ATokenV1BurnOnLiquidationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ATokenV1BurnOnLiquidation represents a BurnOnLiquidation event raised by the ATokenV1 contract.
type ATokenV1BurnOnLiquidation struct {
	From                common.Address
	Value               *big.Int
	FromBalanceIncrease *big.Int
	FromIndex           *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterBurnOnLiquidation is a free log retrieval operation binding the contract event 0x90e5d3d68ec162d9c7de393037a3ede04dd44f68e051bf2ace4a73c299dbc4db.
//
// Solidity: event BurnOnLiquidation(address indexed _from, uint256 _value, uint256 _fromBalanceIncrease, uint256 _fromIndex)
func (_ATokenV1 *ATokenV1Filterer) FilterBurnOnLiquidation(opts *bind.FilterOpts, _from []common.Address) (*ATokenV1BurnOnLiquidationIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _ATokenV1.contract.FilterLogs(opts, "BurnOnLiquidation", _fromRule)
	if err != nil {
		return nil, err
	}
	return &ATokenV1BurnOnLiquidationIterator{contract: _ATokenV1.contract, event: "BurnOnLiquidation", logs: logs, sub: sub}, nil
}

// WatchBurnOnLiquidation is a free log subscription operation binding the contract event 0x90e5d3d68ec162d9c7de393037a3ede04dd44f68e051bf2ace4a73c299dbc4db.
//
// Solidity: event BurnOnLiquidation(address indexed _from, uint256 _value, uint256 _fromBalanceIncrease, uint256 _fromIndex)
func (_ATokenV1 *ATokenV1Filterer) WatchBurnOnLiquidation(opts *bind.WatchOpts, sink chan<- *ATokenV1BurnOnLiquidation, _from []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _ATokenV1.contract.WatchLogs(opts, "BurnOnLiquidation", _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ATokenV1BurnOnLiquidation)
				if err := _ATokenV1.contract.UnpackLog(event, "BurnOnLiquidation", log); err != nil {
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

// ParseBurnOnLiquidation is a log parse operation binding the contract event 0x90e5d3d68ec162d9c7de393037a3ede04dd44f68e051bf2ace4a73c299dbc4db.
//
// Solidity: event BurnOnLiquidation(address indexed _from, uint256 _value, uint256 _fromBalanceIncrease, uint256 _fromIndex)
func (_ATokenV1 *ATokenV1Filterer) ParseBurnOnLiquidation(log types.Log) (*ATokenV1BurnOnLiquidation, error) {
	event := new(ATokenV1BurnOnLiquidation)
	if err := _ATokenV1.contract.UnpackLog(event, "BurnOnLiquidation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ATokenV1InterestRedirectionAllowanceChangedIterator is returned from FilterInterestRedirectionAllowanceChanged and is used to iterate over the raw logs and unpacked data for InterestRedirectionAllowanceChanged events raised by the ATokenV1 contract.
type ATokenV1InterestRedirectionAllowanceChangedIterator struct {
	Event *ATokenV1InterestRedirectionAllowanceChanged // Event containing the contract specifics and raw log

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
func (it *ATokenV1InterestRedirectionAllowanceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ATokenV1InterestRedirectionAllowanceChanged)
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
		it.Event = new(ATokenV1InterestRedirectionAllowanceChanged)
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
func (it *ATokenV1InterestRedirectionAllowanceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ATokenV1InterestRedirectionAllowanceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ATokenV1InterestRedirectionAllowanceChanged represents a InterestRedirectionAllowanceChanged event raised by the ATokenV1 contract.
type ATokenV1InterestRedirectionAllowanceChanged struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterInterestRedirectionAllowanceChanged is a free log retrieval operation binding the contract event 0xc2d6a42a9d5273283f73009a07aacfb043f2f91173a8d9d33b504afe898db08b.
//
// Solidity: event InterestRedirectionAllowanceChanged(address indexed _from, address indexed _to)
func (_ATokenV1 *ATokenV1Filterer) FilterInterestRedirectionAllowanceChanged(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*ATokenV1InterestRedirectionAllowanceChangedIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ATokenV1.contract.FilterLogs(opts, "InterestRedirectionAllowanceChanged", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &ATokenV1InterestRedirectionAllowanceChangedIterator{contract: _ATokenV1.contract, event: "InterestRedirectionAllowanceChanged", logs: logs, sub: sub}, nil
}

// WatchInterestRedirectionAllowanceChanged is a free log subscription operation binding the contract event 0xc2d6a42a9d5273283f73009a07aacfb043f2f91173a8d9d33b504afe898db08b.
//
// Solidity: event InterestRedirectionAllowanceChanged(address indexed _from, address indexed _to)
func (_ATokenV1 *ATokenV1Filterer) WatchInterestRedirectionAllowanceChanged(opts *bind.WatchOpts, sink chan<- *ATokenV1InterestRedirectionAllowanceChanged, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ATokenV1.contract.WatchLogs(opts, "InterestRedirectionAllowanceChanged", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ATokenV1InterestRedirectionAllowanceChanged)
				if err := _ATokenV1.contract.UnpackLog(event, "InterestRedirectionAllowanceChanged", log); err != nil {
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

// ParseInterestRedirectionAllowanceChanged is a log parse operation binding the contract event 0xc2d6a42a9d5273283f73009a07aacfb043f2f91173a8d9d33b504afe898db08b.
//
// Solidity: event InterestRedirectionAllowanceChanged(address indexed _from, address indexed _to)
func (_ATokenV1 *ATokenV1Filterer) ParseInterestRedirectionAllowanceChanged(log types.Log) (*ATokenV1InterestRedirectionAllowanceChanged, error) {
	event := new(ATokenV1InterestRedirectionAllowanceChanged)
	if err := _ATokenV1.contract.UnpackLog(event, "InterestRedirectionAllowanceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ATokenV1InterestStreamRedirectedIterator is returned from FilterInterestStreamRedirected and is used to iterate over the raw logs and unpacked data for InterestStreamRedirected events raised by the ATokenV1 contract.
type ATokenV1InterestStreamRedirectedIterator struct {
	Event *ATokenV1InterestStreamRedirected // Event containing the contract specifics and raw log

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
func (it *ATokenV1InterestStreamRedirectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ATokenV1InterestStreamRedirected)
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
		it.Event = new(ATokenV1InterestStreamRedirected)
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
func (it *ATokenV1InterestStreamRedirectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ATokenV1InterestStreamRedirectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ATokenV1InterestStreamRedirected represents a InterestStreamRedirected event raised by the ATokenV1 contract.
type ATokenV1InterestStreamRedirected struct {
	From                common.Address
	To                  common.Address
	RedirectedBalance   *big.Int
	FromBalanceIncrease *big.Int
	FromIndex           *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterInterestStreamRedirected is a free log retrieval operation binding the contract event 0x5e3cad45b1fe24159d1cb39788d82d0f69cc15770aa96fba1d3d1a7348735594.
//
// Solidity: event InterestStreamRedirected(address indexed _from, address indexed _to, uint256 _redirectedBalance, uint256 _fromBalanceIncrease, uint256 _fromIndex)
func (_ATokenV1 *ATokenV1Filterer) FilterInterestStreamRedirected(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*ATokenV1InterestStreamRedirectedIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ATokenV1.contract.FilterLogs(opts, "InterestStreamRedirected", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &ATokenV1InterestStreamRedirectedIterator{contract: _ATokenV1.contract, event: "InterestStreamRedirected", logs: logs, sub: sub}, nil
}

// WatchInterestStreamRedirected is a free log subscription operation binding the contract event 0x5e3cad45b1fe24159d1cb39788d82d0f69cc15770aa96fba1d3d1a7348735594.
//
// Solidity: event InterestStreamRedirected(address indexed _from, address indexed _to, uint256 _redirectedBalance, uint256 _fromBalanceIncrease, uint256 _fromIndex)
func (_ATokenV1 *ATokenV1Filterer) WatchInterestStreamRedirected(opts *bind.WatchOpts, sink chan<- *ATokenV1InterestStreamRedirected, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ATokenV1.contract.WatchLogs(opts, "InterestStreamRedirected", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ATokenV1InterestStreamRedirected)
				if err := _ATokenV1.contract.UnpackLog(event, "InterestStreamRedirected", log); err != nil {
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

// ParseInterestStreamRedirected is a log parse operation binding the contract event 0x5e3cad45b1fe24159d1cb39788d82d0f69cc15770aa96fba1d3d1a7348735594.
//
// Solidity: event InterestStreamRedirected(address indexed _from, address indexed _to, uint256 _redirectedBalance, uint256 _fromBalanceIncrease, uint256 _fromIndex)
func (_ATokenV1 *ATokenV1Filterer) ParseInterestStreamRedirected(log types.Log) (*ATokenV1InterestStreamRedirected, error) {
	event := new(ATokenV1InterestStreamRedirected)
	if err := _ATokenV1.contract.UnpackLog(event, "InterestStreamRedirected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ATokenV1MintOnDepositIterator is returned from FilterMintOnDeposit and is used to iterate over the raw logs and unpacked data for MintOnDeposit events raised by the ATokenV1 contract.
type ATokenV1MintOnDepositIterator struct {
	Event *ATokenV1MintOnDeposit // Event containing the contract specifics and raw log

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
func (it *ATokenV1MintOnDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ATokenV1MintOnDeposit)
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
		it.Event = new(ATokenV1MintOnDeposit)
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
func (it *ATokenV1MintOnDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ATokenV1MintOnDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ATokenV1MintOnDeposit represents a MintOnDeposit event raised by the ATokenV1 contract.
type ATokenV1MintOnDeposit struct {
	From                common.Address
	Value               *big.Int
	FromBalanceIncrease *big.Int
	FromIndex           *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterMintOnDeposit is a free log retrieval operation binding the contract event 0xbe7799898ca2d813ff902b487c1b434ab45b47edd8f38b77ad5e99aae8341b7a.
//
// Solidity: event MintOnDeposit(address indexed _from, uint256 _value, uint256 _fromBalanceIncrease, uint256 _fromIndex)
func (_ATokenV1 *ATokenV1Filterer) FilterMintOnDeposit(opts *bind.FilterOpts, _from []common.Address) (*ATokenV1MintOnDepositIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _ATokenV1.contract.FilterLogs(opts, "MintOnDeposit", _fromRule)
	if err != nil {
		return nil, err
	}
	return &ATokenV1MintOnDepositIterator{contract: _ATokenV1.contract, event: "MintOnDeposit", logs: logs, sub: sub}, nil
}

// WatchMintOnDeposit is a free log subscription operation binding the contract event 0xbe7799898ca2d813ff902b487c1b434ab45b47edd8f38b77ad5e99aae8341b7a.
//
// Solidity: event MintOnDeposit(address indexed _from, uint256 _value, uint256 _fromBalanceIncrease, uint256 _fromIndex)
func (_ATokenV1 *ATokenV1Filterer) WatchMintOnDeposit(opts *bind.WatchOpts, sink chan<- *ATokenV1MintOnDeposit, _from []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _ATokenV1.contract.WatchLogs(opts, "MintOnDeposit", _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ATokenV1MintOnDeposit)
				if err := _ATokenV1.contract.UnpackLog(event, "MintOnDeposit", log); err != nil {
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

// ParseMintOnDeposit is a log parse operation binding the contract event 0xbe7799898ca2d813ff902b487c1b434ab45b47edd8f38b77ad5e99aae8341b7a.
//
// Solidity: event MintOnDeposit(address indexed _from, uint256 _value, uint256 _fromBalanceIncrease, uint256 _fromIndex)
func (_ATokenV1 *ATokenV1Filterer) ParseMintOnDeposit(log types.Log) (*ATokenV1MintOnDeposit, error) {
	event := new(ATokenV1MintOnDeposit)
	if err := _ATokenV1.contract.UnpackLog(event, "MintOnDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ATokenV1RedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the ATokenV1 contract.
type ATokenV1RedeemIterator struct {
	Event *ATokenV1Redeem // Event containing the contract specifics and raw log

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
func (it *ATokenV1RedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ATokenV1Redeem)
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
		it.Event = new(ATokenV1Redeem)
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
func (it *ATokenV1RedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ATokenV1RedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ATokenV1Redeem represents a Redeem event raised by the ATokenV1 contract.
type ATokenV1Redeem struct {
	From                common.Address
	Value               *big.Int
	FromBalanceIncrease *big.Int
	FromIndex           *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0xbd5034ffbd47e4e72a94baa2cdb74c6fad73cb3bcdc13036b72ec8306f5a7646.
//
// Solidity: event Redeem(address indexed _from, uint256 _value, uint256 _fromBalanceIncrease, uint256 _fromIndex)
func (_ATokenV1 *ATokenV1Filterer) FilterRedeem(opts *bind.FilterOpts, _from []common.Address) (*ATokenV1RedeemIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _ATokenV1.contract.FilterLogs(opts, "Redeem", _fromRule)
	if err != nil {
		return nil, err
	}
	return &ATokenV1RedeemIterator{contract: _ATokenV1.contract, event: "Redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0xbd5034ffbd47e4e72a94baa2cdb74c6fad73cb3bcdc13036b72ec8306f5a7646.
//
// Solidity: event Redeem(address indexed _from, uint256 _value, uint256 _fromBalanceIncrease, uint256 _fromIndex)
func (_ATokenV1 *ATokenV1Filterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *ATokenV1Redeem, _from []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _ATokenV1.contract.WatchLogs(opts, "Redeem", _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ATokenV1Redeem)
				if err := _ATokenV1.contract.UnpackLog(event, "Redeem", log); err != nil {
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

// ParseRedeem is a log parse operation binding the contract event 0xbd5034ffbd47e4e72a94baa2cdb74c6fad73cb3bcdc13036b72ec8306f5a7646.
//
// Solidity: event Redeem(address indexed _from, uint256 _value, uint256 _fromBalanceIncrease, uint256 _fromIndex)
func (_ATokenV1 *ATokenV1Filterer) ParseRedeem(log types.Log) (*ATokenV1Redeem, error) {
	event := new(ATokenV1Redeem)
	if err := _ATokenV1.contract.UnpackLog(event, "Redeem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ATokenV1RedirectedBalanceUpdatedIterator is returned from FilterRedirectedBalanceUpdated and is used to iterate over the raw logs and unpacked data for RedirectedBalanceUpdated events raised by the ATokenV1 contract.
type ATokenV1RedirectedBalanceUpdatedIterator struct {
	Event *ATokenV1RedirectedBalanceUpdated // Event containing the contract specifics and raw log

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
func (it *ATokenV1RedirectedBalanceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ATokenV1RedirectedBalanceUpdated)
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
		it.Event = new(ATokenV1RedirectedBalanceUpdated)
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
func (it *ATokenV1RedirectedBalanceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ATokenV1RedirectedBalanceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ATokenV1RedirectedBalanceUpdated represents a RedirectedBalanceUpdated event raised by the ATokenV1 contract.
type ATokenV1RedirectedBalanceUpdated struct {
	TargetAddress            common.Address
	TargetBalanceIncrease    *big.Int
	TargetIndex              *big.Int
	RedirectedBalanceAdded   *big.Int
	RedirectedBalanceRemoved *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterRedirectedBalanceUpdated is a free log retrieval operation binding the contract event 0x70ff8cf632603e2bfd1afb7e4061acce53d95356b1be9726b99fa22ba733b01f.
//
// Solidity: event RedirectedBalanceUpdated(address indexed _targetAddress, uint256 _targetBalanceIncrease, uint256 _targetIndex, uint256 _redirectedBalanceAdded, uint256 _redirectedBalanceRemoved)
func (_ATokenV1 *ATokenV1Filterer) FilterRedirectedBalanceUpdated(opts *bind.FilterOpts, _targetAddress []common.Address) (*ATokenV1RedirectedBalanceUpdatedIterator, error) {

	var _targetAddressRule []interface{}
	for _, _targetAddressItem := range _targetAddress {
		_targetAddressRule = append(_targetAddressRule, _targetAddressItem)
	}

	logs, sub, err := _ATokenV1.contract.FilterLogs(opts, "RedirectedBalanceUpdated", _targetAddressRule)
	if err != nil {
		return nil, err
	}
	return &ATokenV1RedirectedBalanceUpdatedIterator{contract: _ATokenV1.contract, event: "RedirectedBalanceUpdated", logs: logs, sub: sub}, nil
}

// WatchRedirectedBalanceUpdated is a free log subscription operation binding the contract event 0x70ff8cf632603e2bfd1afb7e4061acce53d95356b1be9726b99fa22ba733b01f.
//
// Solidity: event RedirectedBalanceUpdated(address indexed _targetAddress, uint256 _targetBalanceIncrease, uint256 _targetIndex, uint256 _redirectedBalanceAdded, uint256 _redirectedBalanceRemoved)
func (_ATokenV1 *ATokenV1Filterer) WatchRedirectedBalanceUpdated(opts *bind.WatchOpts, sink chan<- *ATokenV1RedirectedBalanceUpdated, _targetAddress []common.Address) (event.Subscription, error) {

	var _targetAddressRule []interface{}
	for _, _targetAddressItem := range _targetAddress {
		_targetAddressRule = append(_targetAddressRule, _targetAddressItem)
	}

	logs, sub, err := _ATokenV1.contract.WatchLogs(opts, "RedirectedBalanceUpdated", _targetAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ATokenV1RedirectedBalanceUpdated)
				if err := _ATokenV1.contract.UnpackLog(event, "RedirectedBalanceUpdated", log); err != nil {
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

// ParseRedirectedBalanceUpdated is a log parse operation binding the contract event 0x70ff8cf632603e2bfd1afb7e4061acce53d95356b1be9726b99fa22ba733b01f.
//
// Solidity: event RedirectedBalanceUpdated(address indexed _targetAddress, uint256 _targetBalanceIncrease, uint256 _targetIndex, uint256 _redirectedBalanceAdded, uint256 _redirectedBalanceRemoved)
func (_ATokenV1 *ATokenV1Filterer) ParseRedirectedBalanceUpdated(log types.Log) (*ATokenV1RedirectedBalanceUpdated, error) {
	event := new(ATokenV1RedirectedBalanceUpdated)
	if err := _ATokenV1.contract.UnpackLog(event, "RedirectedBalanceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ATokenV1TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ATokenV1 contract.
type ATokenV1TransferIterator struct {
	Event *ATokenV1Transfer // Event containing the contract specifics and raw log

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
func (it *ATokenV1TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ATokenV1Transfer)
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
		it.Event = new(ATokenV1Transfer)
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
func (it *ATokenV1TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ATokenV1TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ATokenV1Transfer represents a Transfer event raised by the ATokenV1 contract.
type ATokenV1Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ATokenV1 *ATokenV1Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ATokenV1TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ATokenV1.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ATokenV1TransferIterator{contract: _ATokenV1.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ATokenV1 *ATokenV1Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ATokenV1Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ATokenV1.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ATokenV1Transfer)
				if err := _ATokenV1.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ATokenV1 *ATokenV1Filterer) ParseTransfer(log types.Log) (*ATokenV1Transfer, error) {
	event := new(ATokenV1Transfer)
	if err := _ATokenV1.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
