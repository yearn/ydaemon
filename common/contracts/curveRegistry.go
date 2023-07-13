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

// CurveRegistryMetaData contains all meta data concerning the CurveRegistry contract.
var CurveRegistryMetaData = &bind.MetaData{
	ABI: "[{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_virtual_price_from_lp_token\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]}]",
}

// CurveRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use CurveRegistryMetaData.ABI instead.
var CurveRegistryABI = CurveRegistryMetaData.ABI

// CurveRegistry is an auto generated Go binding around an Ethereum contract.
type CurveRegistry struct {
	CurveRegistryCaller     // Read-only binding to the contract
	CurveRegistryTransactor // Write-only binding to the contract
	CurveRegistryFilterer   // Log filterer for contract events
}

// CurveRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurveRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurveRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurveRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurveRegistrySession struct {
	Contract     *CurveRegistry    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CurveRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurveRegistryCallerSession struct {
	Contract *CurveRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// CurveRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurveRegistryTransactorSession struct {
	Contract     *CurveRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// CurveRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurveRegistryRaw struct {
	Contract *CurveRegistry // Generic contract binding to access the raw methods on
}

// CurveRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurveRegistryCallerRaw struct {
	Contract *CurveRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// CurveRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurveRegistryTransactorRaw struct {
	Contract *CurveRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurveRegistry creates a new instance of CurveRegistry, bound to a specific deployed contract.
func NewCurveRegistry(address common.Address, backend bind.ContractBackend) (*CurveRegistry, error) {
	contract, err := bindCurveRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CurveRegistry{CurveRegistryCaller: CurveRegistryCaller{contract: contract}, CurveRegistryTransactor: CurveRegistryTransactor{contract: contract}, CurveRegistryFilterer: CurveRegistryFilterer{contract: contract}}, nil
}

// NewCurveRegistryCaller creates a new read-only instance of CurveRegistry, bound to a specific deployed contract.
func NewCurveRegistryCaller(address common.Address, caller bind.ContractCaller) (*CurveRegistryCaller, error) {
	contract, err := bindCurveRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurveRegistryCaller{contract: contract}, nil
}

// NewCurveRegistryTransactor creates a new write-only instance of CurveRegistry, bound to a specific deployed contract.
func NewCurveRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*CurveRegistryTransactor, error) {
	contract, err := bindCurveRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurveRegistryTransactor{contract: contract}, nil
}

// NewCurveRegistryFilterer creates a new log filterer instance of CurveRegistry, bound to a specific deployed contract.
func NewCurveRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*CurveRegistryFilterer, error) {
	contract, err := bindCurveRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurveRegistryFilterer{contract: contract}, nil
}

// bindCurveRegistry binds a generic wrapper to an already deployed contract.
func bindCurveRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CurveRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveRegistry *CurveRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveRegistry.Contract.CurveRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveRegistry *CurveRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveRegistry.Contract.CurveRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveRegistry *CurveRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveRegistry.Contract.CurveRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveRegistry *CurveRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveRegistry *CurveRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveRegistry *CurveRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetVirtualPriceFromLpToken is a free data retrieval call binding the contract method 0xc5b7074a.
//
// Solidity: function get_virtual_price_from_lp_token(address _token) view returns(uint256)
func (_CurveRegistry *CurveRegistryCaller) GetVirtualPriceFromLpToken(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveRegistry.contract.Call(opts, &out, "get_virtual_price_from_lp_token", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPriceFromLpToken is a free data retrieval call binding the contract method 0xc5b7074a.
//
// Solidity: function get_virtual_price_from_lp_token(address _token) view returns(uint256)
func (_CurveRegistry *CurveRegistrySession) GetVirtualPriceFromLpToken(_token common.Address) (*big.Int, error) {
	return _CurveRegistry.Contract.GetVirtualPriceFromLpToken(&_CurveRegistry.CallOpts, _token)
}

// GetVirtualPriceFromLpToken is a free data retrieval call binding the contract method 0xc5b7074a.
//
// Solidity: function get_virtual_price_from_lp_token(address _token) view returns(uint256)
func (_CurveRegistry *CurveRegistryCallerSession) GetVirtualPriceFromLpToken(_token common.Address) (*big.Int, error) {
	return _CurveRegistry.Contract.GetVirtualPriceFromLpToken(&_CurveRegistry.CallOpts, _token)
}
