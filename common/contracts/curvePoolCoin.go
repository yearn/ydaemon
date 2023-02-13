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

// CurvePoolCoinMetaData contains all meta data concerning the CurvePoolCoin contract.
var CurvePoolCoinMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"Approval\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"indexed\":true},{\"name\":\"_spender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"_value\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Transfer\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\",\"indexed\":true},{\"name\":\"_to\",\"type\":\"address\",\"indexed\":true},{\"name\":\"_value\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":78632},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":116616},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":39151},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"permit\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_deadline\",\"type\":\"uint256\"},{\"name\":\"_v\",\"type\":\"uint8\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":102221},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"increaseAllowance\",\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_added_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":41711},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"decreaseAllowance\",\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtracted_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":41737},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":80902},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"mint_relative\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"frac\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":81224},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"burnFrom\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":80954},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"gas\":630},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"gas\":6287},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[],\"gas\":259459},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"gas\":13109},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"gas\":10868},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"DOMAIN_SEPARATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"gas\":2880},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3176},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"},{\"name\":\"arg1\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3472},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":2970},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"minter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3000},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3296}]",
}

// CurvePoolCoinABI is the input ABI used to generate the binding from.
// Deprecated: Use CurvePoolCoinMetaData.ABI instead.
var CurvePoolCoinABI = CurvePoolCoinMetaData.ABI

// CurvePoolCoin is an auto generated Go binding around an Ethereum contract.
type CurvePoolCoin struct {
	CurvePoolCoinCaller     // Read-only binding to the contract
	CurvePoolCoinTransactor // Write-only binding to the contract
	CurvePoolCoinFilterer   // Log filterer for contract events
}

// CurvePoolCoinCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurvePoolCoinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurvePoolCoinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurvePoolCoinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurvePoolCoinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurvePoolCoinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurvePoolCoinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurvePoolCoinSession struct {
	Contract     *CurvePoolCoin    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CurvePoolCoinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurvePoolCoinCallerSession struct {
	Contract *CurvePoolCoinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// CurvePoolCoinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurvePoolCoinTransactorSession struct {
	Contract     *CurvePoolCoinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// CurvePoolCoinRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurvePoolCoinRaw struct {
	Contract *CurvePoolCoin // Generic contract binding to access the raw methods on
}

// CurvePoolCoinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurvePoolCoinCallerRaw struct {
	Contract *CurvePoolCoinCaller // Generic read-only contract binding to access the raw methods on
}

// CurvePoolCoinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurvePoolCoinTransactorRaw struct {
	Contract *CurvePoolCoinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurvePoolCoin creates a new instance of CurvePoolCoin, bound to a specific deployed contract.
func NewCurvePoolCoin(address common.Address, backend bind.ContractBackend) (*CurvePoolCoin, error) {
	contract, err := bindCurvePoolCoin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CurvePoolCoin{CurvePoolCoinCaller: CurvePoolCoinCaller{contract: contract}, CurvePoolCoinTransactor: CurvePoolCoinTransactor{contract: contract}, CurvePoolCoinFilterer: CurvePoolCoinFilterer{contract: contract}}, nil
}

// NewCurvePoolCoinCaller creates a new read-only instance of CurvePoolCoin, bound to a specific deployed contract.
func NewCurvePoolCoinCaller(address common.Address, caller bind.ContractCaller) (*CurvePoolCoinCaller, error) {
	contract, err := bindCurvePoolCoin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurvePoolCoinCaller{contract: contract}, nil
}

// NewCurvePoolCoinTransactor creates a new write-only instance of CurvePoolCoin, bound to a specific deployed contract.
func NewCurvePoolCoinTransactor(address common.Address, transactor bind.ContractTransactor) (*CurvePoolCoinTransactor, error) {
	contract, err := bindCurvePoolCoin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurvePoolCoinTransactor{contract: contract}, nil
}

// NewCurvePoolCoinFilterer creates a new log filterer instance of CurvePoolCoin, bound to a specific deployed contract.
func NewCurvePoolCoinFilterer(address common.Address, filterer bind.ContractFilterer) (*CurvePoolCoinFilterer, error) {
	contract, err := bindCurvePoolCoin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurvePoolCoinFilterer{contract: contract}, nil
}

// bindCurvePoolCoin binds a generic wrapper to an already deployed contract.
func bindCurvePoolCoin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CurvePoolCoinABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurvePoolCoin *CurvePoolCoinRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurvePoolCoin.Contract.CurvePoolCoinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurvePoolCoin *CurvePoolCoinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurvePoolCoin.Contract.CurvePoolCoinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurvePoolCoin *CurvePoolCoinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurvePoolCoin.Contract.CurvePoolCoinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurvePoolCoin *CurvePoolCoinCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurvePoolCoin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurvePoolCoin *CurvePoolCoinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurvePoolCoin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurvePoolCoin *CurvePoolCoinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurvePoolCoin.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_CurvePoolCoin *CurvePoolCoinCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CurvePoolCoin.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_CurvePoolCoin *CurvePoolCoinSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _CurvePoolCoin.Contract.DOMAINSEPARATOR(&_CurvePoolCoin.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_CurvePoolCoin *CurvePoolCoinCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _CurvePoolCoin.Contract.DOMAINSEPARATOR(&_CurvePoolCoin.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_CurvePoolCoin *CurvePoolCoinCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolCoin.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_CurvePoolCoin *CurvePoolCoinSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _CurvePoolCoin.Contract.Allowance(&_CurvePoolCoin.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_CurvePoolCoin *CurvePoolCoinCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _CurvePoolCoin.Contract.Allowance(&_CurvePoolCoin.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_CurvePoolCoin *CurvePoolCoinCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolCoin.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_CurvePoolCoin *CurvePoolCoinSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _CurvePoolCoin.Contract.BalanceOf(&_CurvePoolCoin.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_CurvePoolCoin *CurvePoolCoinCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _CurvePoolCoin.Contract.BalanceOf(&_CurvePoolCoin.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CurvePoolCoin *CurvePoolCoinCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CurvePoolCoin.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CurvePoolCoin *CurvePoolCoinSession) Decimals() (uint8, error) {
	return _CurvePoolCoin.Contract.Decimals(&_CurvePoolCoin.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CurvePoolCoin *CurvePoolCoinCallerSession) Decimals() (uint8, error) {
	return _CurvePoolCoin.Contract.Decimals(&_CurvePoolCoin.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_CurvePoolCoin *CurvePoolCoinCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurvePoolCoin.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_CurvePoolCoin *CurvePoolCoinSession) Minter() (common.Address, error) {
	return _CurvePoolCoin.Contract.Minter(&_CurvePoolCoin.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_CurvePoolCoin *CurvePoolCoinCallerSession) Minter() (common.Address, error) {
	return _CurvePoolCoin.Contract.Minter(&_CurvePoolCoin.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurvePoolCoin *CurvePoolCoinCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CurvePoolCoin.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurvePoolCoin *CurvePoolCoinSession) Name() (string, error) {
	return _CurvePoolCoin.Contract.Name(&_CurvePoolCoin.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurvePoolCoin *CurvePoolCoinCallerSession) Name() (string, error) {
	return _CurvePoolCoin.Contract.Name(&_CurvePoolCoin.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_CurvePoolCoin *CurvePoolCoinCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolCoin.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_CurvePoolCoin *CurvePoolCoinSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _CurvePoolCoin.Contract.Nonces(&_CurvePoolCoin.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_CurvePoolCoin *CurvePoolCoinCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _CurvePoolCoin.Contract.Nonces(&_CurvePoolCoin.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurvePoolCoin *CurvePoolCoinCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CurvePoolCoin.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurvePoolCoin *CurvePoolCoinSession) Symbol() (string, error) {
	return _CurvePoolCoin.Contract.Symbol(&_CurvePoolCoin.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurvePoolCoin *CurvePoolCoinCallerSession) Symbol() (string, error) {
	return _CurvePoolCoin.Contract.Symbol(&_CurvePoolCoin.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurvePoolCoin *CurvePoolCoinCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolCoin.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurvePoolCoin *CurvePoolCoinSession) TotalSupply() (*big.Int, error) {
	return _CurvePoolCoin.Contract.TotalSupply(&_CurvePoolCoin.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurvePoolCoin *CurvePoolCoinCallerSession) TotalSupply() (*big.Int, error) {
	return _CurvePoolCoin.Contract.TotalSupply(&_CurvePoolCoin.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_CurvePoolCoin *CurvePoolCoinCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CurvePoolCoin.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_CurvePoolCoin *CurvePoolCoinSession) Version() (string, error) {
	return _CurvePoolCoin.Contract.Version(&_CurvePoolCoin.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_CurvePoolCoin *CurvePoolCoinCallerSession) Version() (string, error) {
	return _CurvePoolCoin.Contract.Version(&_CurvePoolCoin.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_CurvePoolCoin *CurvePoolCoinTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurvePoolCoin.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_CurvePoolCoin *CurvePoolCoinSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurvePoolCoin.Contract.Approve(&_CurvePoolCoin.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_CurvePoolCoin *CurvePoolCoinTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurvePoolCoin.Contract.Approve(&_CurvePoolCoin.TransactOpts, _spender, _value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address _to, uint256 _value) returns(bool)
func (_CurvePoolCoin *CurvePoolCoinTransactor) BurnFrom(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurvePoolCoin.contract.Transact(opts, "burnFrom", _to, _value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address _to, uint256 _value) returns(bool)
func (_CurvePoolCoin *CurvePoolCoinSession) BurnFrom(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurvePoolCoin.Contract.BurnFrom(&_CurvePoolCoin.TransactOpts, _to, _value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address _to, uint256 _value) returns(bool)
func (_CurvePoolCoin *CurvePoolCoinTransactorSession) BurnFrom(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurvePoolCoin.Contract.BurnFrom(&_CurvePoolCoin.TransactOpts, _to, _value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtracted_value) returns(bool)
func (_CurvePoolCoin *CurvePoolCoinTransactor) DecreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _subtracted_value *big.Int) (*types.Transaction, error) {
	return _CurvePoolCoin.contract.Transact(opts, "decreaseAllowance", _spender, _subtracted_value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtracted_value) returns(bool)
func (_CurvePoolCoin *CurvePoolCoinSession) DecreaseAllowance(_spender common.Address, _subtracted_value *big.Int) (*types.Transaction, error) {
	return _CurvePoolCoin.Contract.DecreaseAllowance(&_CurvePoolCoin.TransactOpts, _spender, _subtracted_value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtracted_value) returns(bool)
func (_CurvePoolCoin *CurvePoolCoinTransactorSession) DecreaseAllowance(_spender common.Address, _subtracted_value *big.Int) (*types.Transaction, error) {
	return _CurvePoolCoin.Contract.DecreaseAllowance(&_CurvePoolCoin.TransactOpts, _spender, _subtracted_value)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _added_value) returns(bool)
func (_CurvePoolCoin *CurvePoolCoinTransactor) IncreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _added_value *big.Int) (*types.Transaction, error) {
	return _CurvePoolCoin.contract.Transact(opts, "increaseAllowance", _spender, _added_value)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _added_value) returns(bool)
func (_CurvePoolCoin *CurvePoolCoinSession) IncreaseAllowance(_spender common.Address, _added_value *big.Int) (*types.Transaction, error) {
	return _CurvePoolCoin.Contract.IncreaseAllowance(&_CurvePoolCoin.TransactOpts, _spender, _added_value)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _added_value) returns(bool)
func (_CurvePoolCoin *CurvePoolCoinTransactorSession) IncreaseAllowance(_spender common.Address, _added_value *big.Int) (*types.Transaction, error) {
	return _CurvePoolCoin.Contract.IncreaseAllowance(&_CurvePoolCoin.TransactOpts, _spender, _added_value)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string _name, string _symbol, address _pool) returns()
func (_CurvePoolCoin *CurvePoolCoinTransactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string, _pool common.Address) (*types.Transaction, error) {
	return _CurvePoolCoin.contract.Transact(opts, "initialize", _name, _symbol, _pool)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string _name, string _symbol, address _pool) returns()
func (_CurvePoolCoin *CurvePoolCoinSession) Initialize(_name string, _symbol string, _pool common.Address) (*types.Transaction, error) {
	return _CurvePoolCoin.Contract.Initialize(&_CurvePoolCoin.TransactOpts, _name, _symbol, _pool)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string _name, string _symbol, address _pool) returns()
func (_CurvePoolCoin *CurvePoolCoinTransactorSession) Initialize(_name string, _symbol string, _pool common.Address) (*types.Transaction, error) {
	return _CurvePoolCoin.Contract.Initialize(&_CurvePoolCoin.TransactOpts, _name, _symbol, _pool)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _value) returns(bool)
func (_CurvePoolCoin *CurvePoolCoinTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurvePoolCoin.contract.Transact(opts, "mint", _to, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _value) returns(bool)
func (_CurvePoolCoin *CurvePoolCoinSession) Mint(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurvePoolCoin.Contract.Mint(&_CurvePoolCoin.TransactOpts, _to, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _value) returns(bool)
func (_CurvePoolCoin *CurvePoolCoinTransactorSession) Mint(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurvePoolCoin.Contract.Mint(&_CurvePoolCoin.TransactOpts, _to, _value)
}

// MintRelative is a paid mutator transaction binding the contract method 0x6962f845.
//
// Solidity: function mint_relative(address _to, uint256 frac) returns(uint256)
func (_CurvePoolCoin *CurvePoolCoinTransactor) MintRelative(opts *bind.TransactOpts, _to common.Address, frac *big.Int) (*types.Transaction, error) {
	return _CurvePoolCoin.contract.Transact(opts, "mint_relative", _to, frac)
}

// MintRelative is a paid mutator transaction binding the contract method 0x6962f845.
//
// Solidity: function mint_relative(address _to, uint256 frac) returns(uint256)
func (_CurvePoolCoin *CurvePoolCoinSession) MintRelative(_to common.Address, frac *big.Int) (*types.Transaction, error) {
	return _CurvePoolCoin.Contract.MintRelative(&_CurvePoolCoin.TransactOpts, _to, frac)
}

// MintRelative is a paid mutator transaction binding the contract method 0x6962f845.
//
// Solidity: function mint_relative(address _to, uint256 frac) returns(uint256)
func (_CurvePoolCoin *CurvePoolCoinTransactorSession) MintRelative(_to common.Address, frac *big.Int) (*types.Transaction, error) {
	return _CurvePoolCoin.Contract.MintRelative(&_CurvePoolCoin.TransactOpts, _to, frac)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(bool)
func (_CurvePoolCoin *CurvePoolCoinTransactor) Permit(opts *bind.TransactOpts, _owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _CurvePoolCoin.contract.Transact(opts, "permit", _owner, _spender, _value, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(bool)
func (_CurvePoolCoin *CurvePoolCoinSession) Permit(_owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _CurvePoolCoin.Contract.Permit(&_CurvePoolCoin.TransactOpts, _owner, _spender, _value, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(bool)
func (_CurvePoolCoin *CurvePoolCoinTransactorSession) Permit(_owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _CurvePoolCoin.Contract.Permit(&_CurvePoolCoin.TransactOpts, _owner, _spender, _value, _deadline, _v, _r, _s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_CurvePoolCoin *CurvePoolCoinTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurvePoolCoin.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_CurvePoolCoin *CurvePoolCoinSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurvePoolCoin.Contract.Transfer(&_CurvePoolCoin.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_CurvePoolCoin *CurvePoolCoinTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurvePoolCoin.Contract.Transfer(&_CurvePoolCoin.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_CurvePoolCoin *CurvePoolCoinTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurvePoolCoin.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_CurvePoolCoin *CurvePoolCoinSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurvePoolCoin.Contract.TransferFrom(&_CurvePoolCoin.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_CurvePoolCoin *CurvePoolCoinTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurvePoolCoin.Contract.TransferFrom(&_CurvePoolCoin.TransactOpts, _from, _to, _value)
}

// CurvePoolCoinApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CurvePoolCoin contract.
type CurvePoolCoinApprovalIterator struct {
	Event *CurvePoolCoinApproval // Event containing the contract specifics and raw log

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
func (it *CurvePoolCoinApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolCoinApproval)
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
		it.Event = new(CurvePoolCoinApproval)
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
func (it *CurvePoolCoinApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolCoinApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolCoinApproval represents a Approval event raised by the CurvePoolCoin contract.
type CurvePoolCoinApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_CurvePoolCoin *CurvePoolCoinFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*CurvePoolCoinApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _CurvePoolCoin.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &CurvePoolCoinApprovalIterator{contract: _CurvePoolCoin.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_CurvePoolCoin *CurvePoolCoinFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CurvePoolCoinApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _CurvePoolCoin.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolCoinApproval)
				if err := _CurvePoolCoin.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_CurvePoolCoin *CurvePoolCoinFilterer) ParseApproval(log types.Log) (*CurvePoolCoinApproval, error) {
	event := new(CurvePoolCoinApproval)
	if err := _CurvePoolCoin.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolCoinTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CurvePoolCoin contract.
type CurvePoolCoinTransferIterator struct {
	Event *CurvePoolCoinTransfer // Event containing the contract specifics and raw log

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
func (it *CurvePoolCoinTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolCoinTransfer)
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
		it.Event = new(CurvePoolCoinTransfer)
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
func (it *CurvePoolCoinTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolCoinTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolCoinTransfer represents a Transfer event raised by the CurvePoolCoin contract.
type CurvePoolCoinTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_CurvePoolCoin *CurvePoolCoinFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*CurvePoolCoinTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _CurvePoolCoin.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &CurvePoolCoinTransferIterator{contract: _CurvePoolCoin.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_CurvePoolCoin *CurvePoolCoinFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CurvePoolCoinTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _CurvePoolCoin.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolCoinTransfer)
				if err := _CurvePoolCoin.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_CurvePoolCoin *CurvePoolCoinFilterer) ParseTransfer(log types.Log) (*CurvePoolCoinTransfer, error) {
	event := new(CurvePoolCoinTransfer)
	if err := _CurvePoolCoin.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
