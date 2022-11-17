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

// YRegistryV3MetaData contains all meta data concerning the YRegistryV3 contract.
var YRegistryV3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_releaseRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_legacyRegistry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"existingVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newVault\",\"type\":\"address\"}],\"name\":\"EndorseVaultWithSameVersion\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"GovernanceMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidVaultType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllowedToEndorse\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"vaultType\",\"type\":\"uint256\"}],\"name\":\"VaultAlreadyEndorsed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"v1\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"v2\",\"type\":\"string\"}],\"name\":\"VersionMissmatch\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"canEndorse\",\"type\":\"bool\"}],\"name\":\"ApprovedVaultEndorser\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governance\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovedVaultOwnerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"apiVersion\",\"type\":\"string\"}],\"name\":\"NewVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newRegistry\",\"type\":\"address\"}],\"name\":\"ReleaseRegistryUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"approvedVaultsOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"endorseVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_releaseDelta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_type\",\"type\":\"uint256\"}],\"name\":\"endorseVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_releaseDelta\",\"type\":\"uint256\"}],\"name\":\"endorseVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isVaultEndorsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"latestVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_type\",\"type\":\"uint256\"}],\"name\":\"latestVaultOfType\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_guardian\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewards\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_releaseDelta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_type\",\"type\":\"uint256\"}],\"name\":\"newVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_guardian\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewards\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_releaseDelta\",\"type\":\"uint256\"}],\"name\":\"newVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"numVaults\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"releaseRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovedVaultsOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setVaultEndorsers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newRegistry\",\"type\":\"address\"}],\"name\":\"updateReleaseRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"vaultEndorsers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"vaultType\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"vaults\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// YRegistryV3ABI is the input ABI used to generate the binding from.
// Deprecated: Use YRegistryV3MetaData.ABI instead.
var YRegistryV3ABI = YRegistryV3MetaData.ABI

// YRegistryV3 is an auto generated Go binding around an Ethereum contract.
type YRegistryV3 struct {
	YRegistryV3Caller     // Read-only binding to the contract
	YRegistryV3Transactor // Write-only binding to the contract
	YRegistryV3Filterer   // Log filterer for contract events
}

// YRegistryV3Caller is an auto generated read-only Go binding around an Ethereum contract.
type YRegistryV3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YRegistryV3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type YRegistryV3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YRegistryV3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YRegistryV3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YRegistryV3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YRegistryV3Session struct {
	Contract     *YRegistryV3      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YRegistryV3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YRegistryV3CallerSession struct {
	Contract *YRegistryV3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// YRegistryV3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YRegistryV3TransactorSession struct {
	Contract     *YRegistryV3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// YRegistryV3Raw is an auto generated low-level Go binding around an Ethereum contract.
type YRegistryV3Raw struct {
	Contract *YRegistryV3 // Generic contract binding to access the raw methods on
}

// YRegistryV3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YRegistryV3CallerRaw struct {
	Contract *YRegistryV3Caller // Generic read-only contract binding to access the raw methods on
}

// YRegistryV3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YRegistryV3TransactorRaw struct {
	Contract *YRegistryV3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewYRegistryV3 creates a new instance of YRegistryV3, bound to a specific deployed contract.
func NewYRegistryV3(address common.Address, backend bind.ContractBackend) (*YRegistryV3, error) {
	contract, err := bindYRegistryV3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YRegistryV3{YRegistryV3Caller: YRegistryV3Caller{contract: contract}, YRegistryV3Transactor: YRegistryV3Transactor{contract: contract}, YRegistryV3Filterer: YRegistryV3Filterer{contract: contract}}, nil
}

// NewYRegistryV3Caller creates a new read-only instance of YRegistryV3, bound to a specific deployed contract.
func NewYRegistryV3Caller(address common.Address, caller bind.ContractCaller) (*YRegistryV3Caller, error) {
	contract, err := bindYRegistryV3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YRegistryV3Caller{contract: contract}, nil
}

// NewYRegistryV3Transactor creates a new write-only instance of YRegistryV3, bound to a specific deployed contract.
func NewYRegistryV3Transactor(address common.Address, transactor bind.ContractTransactor) (*YRegistryV3Transactor, error) {
	contract, err := bindYRegistryV3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YRegistryV3Transactor{contract: contract}, nil
}

// NewYRegistryV3Filterer creates a new log filterer instance of YRegistryV3, bound to a specific deployed contract.
func NewYRegistryV3Filterer(address common.Address, filterer bind.ContractFilterer) (*YRegistryV3Filterer, error) {
	contract, err := bindYRegistryV3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YRegistryV3Filterer{contract: contract}, nil
}

// bindYRegistryV3 binds a generic wrapper to an already deployed contract.
func bindYRegistryV3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(YRegistryV3ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YRegistryV3 *YRegistryV3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YRegistryV3.Contract.YRegistryV3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YRegistryV3 *YRegistryV3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YRegistryV3.Contract.YRegistryV3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YRegistryV3 *YRegistryV3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YRegistryV3.Contract.YRegistryV3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YRegistryV3 *YRegistryV3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YRegistryV3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YRegistryV3 *YRegistryV3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YRegistryV3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YRegistryV3 *YRegistryV3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YRegistryV3.Contract.contract.Transact(opts, method, params...)
}

// ApprovedVaultsOwner is a free data retrieval call binding the contract method 0x2b017c8b.
//
// Solidity: function approvedVaultsOwner(address ) view returns(bool)
func (_YRegistryV3 *YRegistryV3Caller) ApprovedVaultsOwner(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _YRegistryV3.contract.Call(opts, &out, "approvedVaultsOwner", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ApprovedVaultsOwner is a free data retrieval call binding the contract method 0x2b017c8b.
//
// Solidity: function approvedVaultsOwner(address ) view returns(bool)
func (_YRegistryV3 *YRegistryV3Session) ApprovedVaultsOwner(arg0 common.Address) (bool, error) {
	return _YRegistryV3.Contract.ApprovedVaultsOwner(&_YRegistryV3.CallOpts, arg0)
}

// ApprovedVaultsOwner is a free data retrieval call binding the contract method 0x2b017c8b.
//
// Solidity: function approvedVaultsOwner(address ) view returns(bool)
func (_YRegistryV3 *YRegistryV3CallerSession) ApprovedVaultsOwner(arg0 common.Address) (bool, error) {
	return _YRegistryV3.Contract.ApprovedVaultsOwner(&_YRegistryV3.CallOpts, arg0)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address ) view returns(bool)
func (_YRegistryV3 *YRegistryV3Caller) IsRegistered(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _YRegistryV3.contract.Call(opts, &out, "isRegistered", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address ) view returns(bool)
func (_YRegistryV3 *YRegistryV3Session) IsRegistered(arg0 common.Address) (bool, error) {
	return _YRegistryV3.Contract.IsRegistered(&_YRegistryV3.CallOpts, arg0)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address ) view returns(bool)
func (_YRegistryV3 *YRegistryV3CallerSession) IsRegistered(arg0 common.Address) (bool, error) {
	return _YRegistryV3.Contract.IsRegistered(&_YRegistryV3.CallOpts, arg0)
}

// IsVaultEndorsed is a free data retrieval call binding the contract method 0xa0dd0bcf.
//
// Solidity: function isVaultEndorsed(address ) view returns(bool)
func (_YRegistryV3 *YRegistryV3Caller) IsVaultEndorsed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _YRegistryV3.contract.Call(opts, &out, "isVaultEndorsed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVaultEndorsed is a free data retrieval call binding the contract method 0xa0dd0bcf.
//
// Solidity: function isVaultEndorsed(address ) view returns(bool)
func (_YRegistryV3 *YRegistryV3Session) IsVaultEndorsed(arg0 common.Address) (bool, error) {
	return _YRegistryV3.Contract.IsVaultEndorsed(&_YRegistryV3.CallOpts, arg0)
}

// IsVaultEndorsed is a free data retrieval call binding the contract method 0xa0dd0bcf.
//
// Solidity: function isVaultEndorsed(address ) view returns(bool)
func (_YRegistryV3 *YRegistryV3CallerSession) IsVaultEndorsed(arg0 common.Address) (bool, error) {
	return _YRegistryV3.Contract.IsVaultEndorsed(&_YRegistryV3.CallOpts, arg0)
}

// LatestVault is a free data retrieval call binding the contract method 0xe177dc70.
//
// Solidity: function latestVault(address _token) view returns(address)
func (_YRegistryV3 *YRegistryV3Caller) LatestVault(opts *bind.CallOpts, _token common.Address) (common.Address, error) {
	var out []interface{}
	err := _YRegistryV3.contract.Call(opts, &out, "latestVault", _token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LatestVault is a free data retrieval call binding the contract method 0xe177dc70.
//
// Solidity: function latestVault(address _token) view returns(address)
func (_YRegistryV3 *YRegistryV3Session) LatestVault(_token common.Address) (common.Address, error) {
	return _YRegistryV3.Contract.LatestVault(&_YRegistryV3.CallOpts, _token)
}

// LatestVault is a free data retrieval call binding the contract method 0xe177dc70.
//
// Solidity: function latestVault(address _token) view returns(address)
func (_YRegistryV3 *YRegistryV3CallerSession) LatestVault(_token common.Address) (common.Address, error) {
	return _YRegistryV3.Contract.LatestVault(&_YRegistryV3.CallOpts, _token)
}

// LatestVaultOfType is a free data retrieval call binding the contract method 0x9de312f0.
//
// Solidity: function latestVaultOfType(address _token, uint256 _type) view returns(address)
func (_YRegistryV3 *YRegistryV3Caller) LatestVaultOfType(opts *bind.CallOpts, _token common.Address, _type *big.Int) (common.Address, error) {
	var out []interface{}
	err := _YRegistryV3.contract.Call(opts, &out, "latestVaultOfType", _token, _type)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LatestVaultOfType is a free data retrieval call binding the contract method 0x9de312f0.
//
// Solidity: function latestVaultOfType(address _token, uint256 _type) view returns(address)
func (_YRegistryV3 *YRegistryV3Session) LatestVaultOfType(_token common.Address, _type *big.Int) (common.Address, error) {
	return _YRegistryV3.Contract.LatestVaultOfType(&_YRegistryV3.CallOpts, _token, _type)
}

// LatestVaultOfType is a free data retrieval call binding the contract method 0x9de312f0.
//
// Solidity: function latestVaultOfType(address _token, uint256 _type) view returns(address)
func (_YRegistryV3 *YRegistryV3CallerSession) LatestVaultOfType(_token common.Address, _type *big.Int) (common.Address, error) {
	return _YRegistryV3.Contract.LatestVaultOfType(&_YRegistryV3.CallOpts, _token, _type)
}

// NumTokens is a free data retrieval call binding the contract method 0x8e499bcf.
//
// Solidity: function numTokens() view returns(uint256)
func (_YRegistryV3 *YRegistryV3Caller) NumTokens(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YRegistryV3.contract.Call(opts, &out, "numTokens")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumTokens is a free data retrieval call binding the contract method 0x8e499bcf.
//
// Solidity: function numTokens() view returns(uint256)
func (_YRegistryV3 *YRegistryV3Session) NumTokens() (*big.Int, error) {
	return _YRegistryV3.Contract.NumTokens(&_YRegistryV3.CallOpts)
}

// NumTokens is a free data retrieval call binding the contract method 0x8e499bcf.
//
// Solidity: function numTokens() view returns(uint256)
func (_YRegistryV3 *YRegistryV3CallerSession) NumTokens() (*big.Int, error) {
	return _YRegistryV3.Contract.NumTokens(&_YRegistryV3.CallOpts)
}

// NumVaults is a free data retrieval call binding the contract method 0xf9c7bba5.
//
// Solidity: function numVaults(address _token) view returns(uint256)
func (_YRegistryV3 *YRegistryV3Caller) NumVaults(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YRegistryV3.contract.Call(opts, &out, "numVaults", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumVaults is a free data retrieval call binding the contract method 0xf9c7bba5.
//
// Solidity: function numVaults(address _token) view returns(uint256)
func (_YRegistryV3 *YRegistryV3Session) NumVaults(_token common.Address) (*big.Int, error) {
	return _YRegistryV3.Contract.NumVaults(&_YRegistryV3.CallOpts, _token)
}

// NumVaults is a free data retrieval call binding the contract method 0xf9c7bba5.
//
// Solidity: function numVaults(address _token) view returns(uint256)
func (_YRegistryV3 *YRegistryV3CallerSession) NumVaults(_token common.Address) (*big.Int, error) {
	return _YRegistryV3.Contract.NumVaults(&_YRegistryV3.CallOpts, _token)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_YRegistryV3 *YRegistryV3Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YRegistryV3.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_YRegistryV3 *YRegistryV3Session) Owner() (common.Address, error) {
	return _YRegistryV3.Contract.Owner(&_YRegistryV3.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_YRegistryV3 *YRegistryV3CallerSession) Owner() (common.Address, error) {
	return _YRegistryV3.Contract.Owner(&_YRegistryV3.CallOpts)
}

// ReleaseRegistry is a free data retrieval call binding the contract method 0x19ee073e.
//
// Solidity: function releaseRegistry() view returns(address)
func (_YRegistryV3 *YRegistryV3Caller) ReleaseRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YRegistryV3.contract.Call(opts, &out, "releaseRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReleaseRegistry is a free data retrieval call binding the contract method 0x19ee073e.
//
// Solidity: function releaseRegistry() view returns(address)
func (_YRegistryV3 *YRegistryV3Session) ReleaseRegistry() (common.Address, error) {
	return _YRegistryV3.Contract.ReleaseRegistry(&_YRegistryV3.CallOpts)
}

// ReleaseRegistry is a free data retrieval call binding the contract method 0x19ee073e.
//
// Solidity: function releaseRegistry() view returns(address)
func (_YRegistryV3 *YRegistryV3CallerSession) ReleaseRegistry() (common.Address, error) {
	return _YRegistryV3.Contract.ReleaseRegistry(&_YRegistryV3.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_YRegistryV3 *YRegistryV3Caller) Tokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _YRegistryV3.contract.Call(opts, &out, "tokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_YRegistryV3 *YRegistryV3Session) Tokens(arg0 *big.Int) (common.Address, error) {
	return _YRegistryV3.Contract.Tokens(&_YRegistryV3.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_YRegistryV3 *YRegistryV3CallerSession) Tokens(arg0 *big.Int) (common.Address, error) {
	return _YRegistryV3.Contract.Tokens(&_YRegistryV3.CallOpts, arg0)
}

// VaultEndorsers is a free data retrieval call binding the contract method 0x6fc74f9e.
//
// Solidity: function vaultEndorsers(address ) view returns(bool)
func (_YRegistryV3 *YRegistryV3Caller) VaultEndorsers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _YRegistryV3.contract.Call(opts, &out, "vaultEndorsers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VaultEndorsers is a free data retrieval call binding the contract method 0x6fc74f9e.
//
// Solidity: function vaultEndorsers(address ) view returns(bool)
func (_YRegistryV3 *YRegistryV3Session) VaultEndorsers(arg0 common.Address) (bool, error) {
	return _YRegistryV3.Contract.VaultEndorsers(&_YRegistryV3.CallOpts, arg0)
}

// VaultEndorsers is a free data retrieval call binding the contract method 0x6fc74f9e.
//
// Solidity: function vaultEndorsers(address ) view returns(bool)
func (_YRegistryV3 *YRegistryV3CallerSession) VaultEndorsers(arg0 common.Address) (bool, error) {
	return _YRegistryV3.Contract.VaultEndorsers(&_YRegistryV3.CallOpts, arg0)
}

// VaultType is a free data retrieval call binding the contract method 0xc9b8f055.
//
// Solidity: function vaultType(address ) view returns(uint256)
func (_YRegistryV3 *YRegistryV3Caller) VaultType(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YRegistryV3.contract.Call(opts, &out, "vaultType", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VaultType is a free data retrieval call binding the contract method 0xc9b8f055.
//
// Solidity: function vaultType(address ) view returns(uint256)
func (_YRegistryV3 *YRegistryV3Session) VaultType(arg0 common.Address) (*big.Int, error) {
	return _YRegistryV3.Contract.VaultType(&_YRegistryV3.CallOpts, arg0)
}

// VaultType is a free data retrieval call binding the contract method 0xc9b8f055.
//
// Solidity: function vaultType(address ) view returns(uint256)
func (_YRegistryV3 *YRegistryV3CallerSession) VaultType(arg0 common.Address) (*big.Int, error) {
	return _YRegistryV3.Contract.VaultType(&_YRegistryV3.CallOpts, arg0)
}

// Vaults is a free data retrieval call binding the contract method 0x7bbfc69e.
//
// Solidity: function vaults(address , uint256 ) view returns(address)
func (_YRegistryV3 *YRegistryV3Caller) Vaults(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _YRegistryV3.contract.Call(opts, &out, "vaults", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vaults is a free data retrieval call binding the contract method 0x7bbfc69e.
//
// Solidity: function vaults(address , uint256 ) view returns(address)
func (_YRegistryV3 *YRegistryV3Session) Vaults(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _YRegistryV3.Contract.Vaults(&_YRegistryV3.CallOpts, arg0, arg1)
}

// Vaults is a free data retrieval call binding the contract method 0x7bbfc69e.
//
// Solidity: function vaults(address , uint256 ) view returns(address)
func (_YRegistryV3 *YRegistryV3CallerSession) Vaults(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _YRegistryV3.Contract.Vaults(&_YRegistryV3.CallOpts, arg0, arg1)
}

// EndorseVault is a paid mutator transaction binding the contract method 0x29b2e0c6.
//
// Solidity: function endorseVault(address _vault) returns()
func (_YRegistryV3 *YRegistryV3Transactor) EndorseVault(opts *bind.TransactOpts, _vault common.Address) (*types.Transaction, error) {
	return _YRegistryV3.contract.Transact(opts, "endorseVault", _vault)
}

// EndorseVault is a paid mutator transaction binding the contract method 0x29b2e0c6.
//
// Solidity: function endorseVault(address _vault) returns()
func (_YRegistryV3 *YRegistryV3Session) EndorseVault(_vault common.Address) (*types.Transaction, error) {
	return _YRegistryV3.Contract.EndorseVault(&_YRegistryV3.TransactOpts, _vault)
}

// EndorseVault is a paid mutator transaction binding the contract method 0x29b2e0c6.
//
// Solidity: function endorseVault(address _vault) returns()
func (_YRegistryV3 *YRegistryV3TransactorSession) EndorseVault(_vault common.Address) (*types.Transaction, error) {
	return _YRegistryV3.Contract.EndorseVault(&_YRegistryV3.TransactOpts, _vault)
}

// EndorseVault0 is a paid mutator transaction binding the contract method 0x931074ba.
//
// Solidity: function endorseVault(address _vault, uint256 _releaseDelta, uint256 _type) returns()
func (_YRegistryV3 *YRegistryV3Transactor) EndorseVault0(opts *bind.TransactOpts, _vault common.Address, _releaseDelta *big.Int, _type *big.Int) (*types.Transaction, error) {
	return _YRegistryV3.contract.Transact(opts, "endorseVault0", _vault, _releaseDelta, _type)
}

// EndorseVault0 is a paid mutator transaction binding the contract method 0x931074ba.
//
// Solidity: function endorseVault(address _vault, uint256 _releaseDelta, uint256 _type) returns()
func (_YRegistryV3 *YRegistryV3Session) EndorseVault0(_vault common.Address, _releaseDelta *big.Int, _type *big.Int) (*types.Transaction, error) {
	return _YRegistryV3.Contract.EndorseVault0(&_YRegistryV3.TransactOpts, _vault, _releaseDelta, _type)
}

// EndorseVault0 is a paid mutator transaction binding the contract method 0x931074ba.
//
// Solidity: function endorseVault(address _vault, uint256 _releaseDelta, uint256 _type) returns()
func (_YRegistryV3 *YRegistryV3TransactorSession) EndorseVault0(_vault common.Address, _releaseDelta *big.Int, _type *big.Int) (*types.Transaction, error) {
	return _YRegistryV3.Contract.EndorseVault0(&_YRegistryV3.TransactOpts, _vault, _releaseDelta, _type)
}

// EndorseVault1 is a paid mutator transaction binding the contract method 0xb366a35c.
//
// Solidity: function endorseVault(address _vault, uint256 _releaseDelta) returns()
func (_YRegistryV3 *YRegistryV3Transactor) EndorseVault1(opts *bind.TransactOpts, _vault common.Address, _releaseDelta *big.Int) (*types.Transaction, error) {
	return _YRegistryV3.contract.Transact(opts, "endorseVault1", _vault, _releaseDelta)
}

// EndorseVault1 is a paid mutator transaction binding the contract method 0xb366a35c.
//
// Solidity: function endorseVault(address _vault, uint256 _releaseDelta) returns()
func (_YRegistryV3 *YRegistryV3Session) EndorseVault1(_vault common.Address, _releaseDelta *big.Int) (*types.Transaction, error) {
	return _YRegistryV3.Contract.EndorseVault1(&_YRegistryV3.TransactOpts, _vault, _releaseDelta)
}

// EndorseVault1 is a paid mutator transaction binding the contract method 0xb366a35c.
//
// Solidity: function endorseVault(address _vault, uint256 _releaseDelta) returns()
func (_YRegistryV3 *YRegistryV3TransactorSession) EndorseVault1(_vault common.Address, _releaseDelta *big.Int) (*types.Transaction, error) {
	return _YRegistryV3.Contract.EndorseVault1(&_YRegistryV3.TransactOpts, _vault, _releaseDelta)
}

// NewVault is a paid mutator transaction binding the contract method 0x125d4336.
//
// Solidity: function newVault(address _token, address _governance, address _guardian, address _rewards, string _name, string _symbol, uint256 _releaseDelta, uint256 _type) returns(address)
func (_YRegistryV3 *YRegistryV3Transactor) NewVault(opts *bind.TransactOpts, _token common.Address, _governance common.Address, _guardian common.Address, _rewards common.Address, _name string, _symbol string, _releaseDelta *big.Int, _type *big.Int) (*types.Transaction, error) {
	return _YRegistryV3.contract.Transact(opts, "newVault", _token, _governance, _guardian, _rewards, _name, _symbol, _releaseDelta, _type)
}

// NewVault is a paid mutator transaction binding the contract method 0x125d4336.
//
// Solidity: function newVault(address _token, address _governance, address _guardian, address _rewards, string _name, string _symbol, uint256 _releaseDelta, uint256 _type) returns(address)
func (_YRegistryV3 *YRegistryV3Session) NewVault(_token common.Address, _governance common.Address, _guardian common.Address, _rewards common.Address, _name string, _symbol string, _releaseDelta *big.Int, _type *big.Int) (*types.Transaction, error) {
	return _YRegistryV3.Contract.NewVault(&_YRegistryV3.TransactOpts, _token, _governance, _guardian, _rewards, _name, _symbol, _releaseDelta, _type)
}

// NewVault is a paid mutator transaction binding the contract method 0x125d4336.
//
// Solidity: function newVault(address _token, address _governance, address _guardian, address _rewards, string _name, string _symbol, uint256 _releaseDelta, uint256 _type) returns(address)
func (_YRegistryV3 *YRegistryV3TransactorSession) NewVault(_token common.Address, _governance common.Address, _guardian common.Address, _rewards common.Address, _name string, _symbol string, _releaseDelta *big.Int, _type *big.Int) (*types.Transaction, error) {
	return _YRegistryV3.Contract.NewVault(&_YRegistryV3.TransactOpts, _token, _governance, _guardian, _rewards, _name, _symbol, _releaseDelta, _type)
}

// NewVault0 is a paid mutator transaction binding the contract method 0xb0b40fce.
//
// Solidity: function newVault(address _token, address _guardian, address _rewards, string _name, string _symbol, uint256 _releaseDelta) returns(address)
func (_YRegistryV3 *YRegistryV3Transactor) NewVault0(opts *bind.TransactOpts, _token common.Address, _guardian common.Address, _rewards common.Address, _name string, _symbol string, _releaseDelta *big.Int) (*types.Transaction, error) {
	return _YRegistryV3.contract.Transact(opts, "newVault0", _token, _guardian, _rewards, _name, _symbol, _releaseDelta)
}

// NewVault0 is a paid mutator transaction binding the contract method 0xb0b40fce.
//
// Solidity: function newVault(address _token, address _guardian, address _rewards, string _name, string _symbol, uint256 _releaseDelta) returns(address)
func (_YRegistryV3 *YRegistryV3Session) NewVault0(_token common.Address, _guardian common.Address, _rewards common.Address, _name string, _symbol string, _releaseDelta *big.Int) (*types.Transaction, error) {
	return _YRegistryV3.Contract.NewVault0(&_YRegistryV3.TransactOpts, _token, _guardian, _rewards, _name, _symbol, _releaseDelta)
}

// NewVault0 is a paid mutator transaction binding the contract method 0xb0b40fce.
//
// Solidity: function newVault(address _token, address _guardian, address _rewards, string _name, string _symbol, uint256 _releaseDelta) returns(address)
func (_YRegistryV3 *YRegistryV3TransactorSession) NewVault0(_token common.Address, _guardian common.Address, _rewards common.Address, _name string, _symbol string, _releaseDelta *big.Int) (*types.Transaction, error) {
	return _YRegistryV3.Contract.NewVault0(&_YRegistryV3.TransactOpts, _token, _guardian, _rewards, _name, _symbol, _releaseDelta)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_YRegistryV3 *YRegistryV3Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YRegistryV3.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_YRegistryV3 *YRegistryV3Session) RenounceOwnership() (*types.Transaction, error) {
	return _YRegistryV3.Contract.RenounceOwnership(&_YRegistryV3.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_YRegistryV3 *YRegistryV3TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _YRegistryV3.Contract.RenounceOwnership(&_YRegistryV3.TransactOpts)
}

// SetApprovedVaultsOwner is a paid mutator transaction binding the contract method 0x51f69a6e.
//
// Solidity: function setApprovedVaultsOwner(address _addr, bool _approved) returns()
func (_YRegistryV3 *YRegistryV3Transactor) SetApprovedVaultsOwner(opts *bind.TransactOpts, _addr common.Address, _approved bool) (*types.Transaction, error) {
	return _YRegistryV3.contract.Transact(opts, "setApprovedVaultsOwner", _addr, _approved)
}

// SetApprovedVaultsOwner is a paid mutator transaction binding the contract method 0x51f69a6e.
//
// Solidity: function setApprovedVaultsOwner(address _addr, bool _approved) returns()
func (_YRegistryV3 *YRegistryV3Session) SetApprovedVaultsOwner(_addr common.Address, _approved bool) (*types.Transaction, error) {
	return _YRegistryV3.Contract.SetApprovedVaultsOwner(&_YRegistryV3.TransactOpts, _addr, _approved)
}

// SetApprovedVaultsOwner is a paid mutator transaction binding the contract method 0x51f69a6e.
//
// Solidity: function setApprovedVaultsOwner(address _addr, bool _approved) returns()
func (_YRegistryV3 *YRegistryV3TransactorSession) SetApprovedVaultsOwner(_addr common.Address, _approved bool) (*types.Transaction, error) {
	return _YRegistryV3.Contract.SetApprovedVaultsOwner(&_YRegistryV3.TransactOpts, _addr, _approved)
}

// SetVaultEndorsers is a paid mutator transaction binding the contract method 0xdbf42461.
//
// Solidity: function setVaultEndorsers(address _addr, bool _approved) returns()
func (_YRegistryV3 *YRegistryV3Transactor) SetVaultEndorsers(opts *bind.TransactOpts, _addr common.Address, _approved bool) (*types.Transaction, error) {
	return _YRegistryV3.contract.Transact(opts, "setVaultEndorsers", _addr, _approved)
}

// SetVaultEndorsers is a paid mutator transaction binding the contract method 0xdbf42461.
//
// Solidity: function setVaultEndorsers(address _addr, bool _approved) returns()
func (_YRegistryV3 *YRegistryV3Session) SetVaultEndorsers(_addr common.Address, _approved bool) (*types.Transaction, error) {
	return _YRegistryV3.Contract.SetVaultEndorsers(&_YRegistryV3.TransactOpts, _addr, _approved)
}

// SetVaultEndorsers is a paid mutator transaction binding the contract method 0xdbf42461.
//
// Solidity: function setVaultEndorsers(address _addr, bool _approved) returns()
func (_YRegistryV3 *YRegistryV3TransactorSession) SetVaultEndorsers(_addr common.Address, _approved bool) (*types.Transaction, error) {
	return _YRegistryV3.Contract.SetVaultEndorsers(&_YRegistryV3.TransactOpts, _addr, _approved)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_YRegistryV3 *YRegistryV3Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _YRegistryV3.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_YRegistryV3 *YRegistryV3Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _YRegistryV3.Contract.TransferOwnership(&_YRegistryV3.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_YRegistryV3 *YRegistryV3TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _YRegistryV3.Contract.TransferOwnership(&_YRegistryV3.TransactOpts, newOwner)
}

// UpdateReleaseRegistry is a paid mutator transaction binding the contract method 0x8cd473c2.
//
// Solidity: function updateReleaseRegistry(address _newRegistry) returns()
func (_YRegistryV3 *YRegistryV3Transactor) UpdateReleaseRegistry(opts *bind.TransactOpts, _newRegistry common.Address) (*types.Transaction, error) {
	return _YRegistryV3.contract.Transact(opts, "updateReleaseRegistry", _newRegistry)
}

// UpdateReleaseRegistry is a paid mutator transaction binding the contract method 0x8cd473c2.
//
// Solidity: function updateReleaseRegistry(address _newRegistry) returns()
func (_YRegistryV3 *YRegistryV3Session) UpdateReleaseRegistry(_newRegistry common.Address) (*types.Transaction, error) {
	return _YRegistryV3.Contract.UpdateReleaseRegistry(&_YRegistryV3.TransactOpts, _newRegistry)
}

// UpdateReleaseRegistry is a paid mutator transaction binding the contract method 0x8cd473c2.
//
// Solidity: function updateReleaseRegistry(address _newRegistry) returns()
func (_YRegistryV3 *YRegistryV3TransactorSession) UpdateReleaseRegistry(_newRegistry common.Address) (*types.Transaction, error) {
	return _YRegistryV3.Contract.UpdateReleaseRegistry(&_YRegistryV3.TransactOpts, _newRegistry)
}

// YRegistryV3ApprovedVaultEndorserIterator is returned from FilterApprovedVaultEndorser and is used to iterate over the raw logs and unpacked data for ApprovedVaultEndorser events raised by the YRegistryV3 contract.
type YRegistryV3ApprovedVaultEndorserIterator struct {
	Event *YRegistryV3ApprovedVaultEndorser // Event containing the contract specifics and raw log

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
func (it *YRegistryV3ApprovedVaultEndorserIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YRegistryV3ApprovedVaultEndorser)
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
		it.Event = new(YRegistryV3ApprovedVaultEndorser)
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
func (it *YRegistryV3ApprovedVaultEndorserIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YRegistryV3ApprovedVaultEndorserIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YRegistryV3ApprovedVaultEndorser represents a ApprovedVaultEndorser event raised by the YRegistryV3 contract.
type YRegistryV3ApprovedVaultEndorser struct {
	Account    common.Address
	CanEndorse bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterApprovedVaultEndorser is a free log retrieval operation binding the contract event 0x0432d20aa3457fdf75aad100b0473a9a5a3f29275b066abbeb243dc7e0db4a48.
//
// Solidity: event ApprovedVaultEndorser(address account, bool canEndorse)
func (_YRegistryV3 *YRegistryV3Filterer) FilterApprovedVaultEndorser(opts *bind.FilterOpts) (*YRegistryV3ApprovedVaultEndorserIterator, error) {

	logs, sub, err := _YRegistryV3.contract.FilterLogs(opts, "ApprovedVaultEndorser")
	if err != nil {
		return nil, err
	}
	return &YRegistryV3ApprovedVaultEndorserIterator{contract: _YRegistryV3.contract, event: "ApprovedVaultEndorser", logs: logs, sub: sub}, nil
}

// WatchApprovedVaultEndorser is a free log subscription operation binding the contract event 0x0432d20aa3457fdf75aad100b0473a9a5a3f29275b066abbeb243dc7e0db4a48.
//
// Solidity: event ApprovedVaultEndorser(address account, bool canEndorse)
func (_YRegistryV3 *YRegistryV3Filterer) WatchApprovedVaultEndorser(opts *bind.WatchOpts, sink chan<- *YRegistryV3ApprovedVaultEndorser) (event.Subscription, error) {

	logs, sub, err := _YRegistryV3.contract.WatchLogs(opts, "ApprovedVaultEndorser")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YRegistryV3ApprovedVaultEndorser)
				if err := _YRegistryV3.contract.UnpackLog(event, "ApprovedVaultEndorser", log); err != nil {
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

// ParseApprovedVaultEndorser is a log parse operation binding the contract event 0x0432d20aa3457fdf75aad100b0473a9a5a3f29275b066abbeb243dc7e0db4a48.
//
// Solidity: event ApprovedVaultEndorser(address account, bool canEndorse)
func (_YRegistryV3 *YRegistryV3Filterer) ParseApprovedVaultEndorser(log types.Log) (*YRegistryV3ApprovedVaultEndorser, error) {
	event := new(YRegistryV3ApprovedVaultEndorser)
	if err := _YRegistryV3.contract.UnpackLog(event, "ApprovedVaultEndorser", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YRegistryV3ApprovedVaultOwnerUpdatedIterator is returned from FilterApprovedVaultOwnerUpdated and is used to iterate over the raw logs and unpacked data for ApprovedVaultOwnerUpdated events raised by the YRegistryV3 contract.
type YRegistryV3ApprovedVaultOwnerUpdatedIterator struct {
	Event *YRegistryV3ApprovedVaultOwnerUpdated // Event containing the contract specifics and raw log

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
func (it *YRegistryV3ApprovedVaultOwnerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YRegistryV3ApprovedVaultOwnerUpdated)
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
		it.Event = new(YRegistryV3ApprovedVaultOwnerUpdated)
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
func (it *YRegistryV3ApprovedVaultOwnerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YRegistryV3ApprovedVaultOwnerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YRegistryV3ApprovedVaultOwnerUpdated represents a ApprovedVaultOwnerUpdated event raised by the YRegistryV3 contract.
type YRegistryV3ApprovedVaultOwnerUpdated struct {
	Governance common.Address
	Approved   bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterApprovedVaultOwnerUpdated is a free log retrieval operation binding the contract event 0xc9219e6871f1efb578009a697784e88c6375187771bda2c8899b8ead4bbea822.
//
// Solidity: event ApprovedVaultOwnerUpdated(address governance, bool approved)
func (_YRegistryV3 *YRegistryV3Filterer) FilterApprovedVaultOwnerUpdated(opts *bind.FilterOpts) (*YRegistryV3ApprovedVaultOwnerUpdatedIterator, error) {

	logs, sub, err := _YRegistryV3.contract.FilterLogs(opts, "ApprovedVaultOwnerUpdated")
	if err != nil {
		return nil, err
	}
	return &YRegistryV3ApprovedVaultOwnerUpdatedIterator{contract: _YRegistryV3.contract, event: "ApprovedVaultOwnerUpdated", logs: logs, sub: sub}, nil
}

// WatchApprovedVaultOwnerUpdated is a free log subscription operation binding the contract event 0xc9219e6871f1efb578009a697784e88c6375187771bda2c8899b8ead4bbea822.
//
// Solidity: event ApprovedVaultOwnerUpdated(address governance, bool approved)
func (_YRegistryV3 *YRegistryV3Filterer) WatchApprovedVaultOwnerUpdated(opts *bind.WatchOpts, sink chan<- *YRegistryV3ApprovedVaultOwnerUpdated) (event.Subscription, error) {

	logs, sub, err := _YRegistryV3.contract.WatchLogs(opts, "ApprovedVaultOwnerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YRegistryV3ApprovedVaultOwnerUpdated)
				if err := _YRegistryV3.contract.UnpackLog(event, "ApprovedVaultOwnerUpdated", log); err != nil {
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

// ParseApprovedVaultOwnerUpdated is a log parse operation binding the contract event 0xc9219e6871f1efb578009a697784e88c6375187771bda2c8899b8ead4bbea822.
//
// Solidity: event ApprovedVaultOwnerUpdated(address governance, bool approved)
func (_YRegistryV3 *YRegistryV3Filterer) ParseApprovedVaultOwnerUpdated(log types.Log) (*YRegistryV3ApprovedVaultOwnerUpdated, error) {
	event := new(YRegistryV3ApprovedVaultOwnerUpdated)
	if err := _YRegistryV3.contract.UnpackLog(event, "ApprovedVaultOwnerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YRegistryV3NewVaultIterator is returned from FilterNewVault and is used to iterate over the raw logs and unpacked data for NewVault events raised by the YRegistryV3 contract.
type YRegistryV3NewVaultIterator struct {
	Event *YRegistryV3NewVault // Event containing the contract specifics and raw log

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
func (it *YRegistryV3NewVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YRegistryV3NewVault)
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
		it.Event = new(YRegistryV3NewVault)
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
func (it *YRegistryV3NewVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YRegistryV3NewVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YRegistryV3NewVault represents a NewVault event raised by the YRegistryV3 contract.
type YRegistryV3NewVault struct {
	Token      common.Address
	VaultId    *big.Int
	VaultType  *big.Int
	Vault      common.Address
	ApiVersion string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewVault is a free log retrieval operation binding the contract event 0x665814fc0711cd098dc48648ad2f63460a3cb758a6aca058c168d7a9067d9f38.
//
// Solidity: event NewVault(address indexed token, uint256 indexed vaultId, uint256 vaultType, address vault, string apiVersion)
func (_YRegistryV3 *YRegistryV3Filterer) FilterNewVault(opts *bind.FilterOpts, token []common.Address, vaultId []*big.Int) (*YRegistryV3NewVaultIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var vaultIdRule []interface{}
	for _, vaultIdItem := range vaultId {
		vaultIdRule = append(vaultIdRule, vaultIdItem)
	}

	logs, sub, err := _YRegistryV3.contract.FilterLogs(opts, "NewVault", tokenRule, vaultIdRule)
	if err != nil {
		return nil, err
	}
	return &YRegistryV3NewVaultIterator{contract: _YRegistryV3.contract, event: "NewVault", logs: logs, sub: sub}, nil
}

// WatchNewVault is a free log subscription operation binding the contract event 0x665814fc0711cd098dc48648ad2f63460a3cb758a6aca058c168d7a9067d9f38.
//
// Solidity: event NewVault(address indexed token, uint256 indexed vaultId, uint256 vaultType, address vault, string apiVersion)
func (_YRegistryV3 *YRegistryV3Filterer) WatchNewVault(opts *bind.WatchOpts, sink chan<- *YRegistryV3NewVault, token []common.Address, vaultId []*big.Int) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var vaultIdRule []interface{}
	for _, vaultIdItem := range vaultId {
		vaultIdRule = append(vaultIdRule, vaultIdItem)
	}

	logs, sub, err := _YRegistryV3.contract.WatchLogs(opts, "NewVault", tokenRule, vaultIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YRegistryV3NewVault)
				if err := _YRegistryV3.contract.UnpackLog(event, "NewVault", log); err != nil {
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

// ParseNewVault is a log parse operation binding the contract event 0x665814fc0711cd098dc48648ad2f63460a3cb758a6aca058c168d7a9067d9f38.
//
// Solidity: event NewVault(address indexed token, uint256 indexed vaultId, uint256 vaultType, address vault, string apiVersion)
func (_YRegistryV3 *YRegistryV3Filterer) ParseNewVault(log types.Log) (*YRegistryV3NewVault, error) {
	event := new(YRegistryV3NewVault)
	if err := _YRegistryV3.contract.UnpackLog(event, "NewVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YRegistryV3OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the YRegistryV3 contract.
type YRegistryV3OwnershipTransferredIterator struct {
	Event *YRegistryV3OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *YRegistryV3OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YRegistryV3OwnershipTransferred)
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
		it.Event = new(YRegistryV3OwnershipTransferred)
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
func (it *YRegistryV3OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YRegistryV3OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YRegistryV3OwnershipTransferred represents a OwnershipTransferred event raised by the YRegistryV3 contract.
type YRegistryV3OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_YRegistryV3 *YRegistryV3Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*YRegistryV3OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _YRegistryV3.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &YRegistryV3OwnershipTransferredIterator{contract: _YRegistryV3.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_YRegistryV3 *YRegistryV3Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *YRegistryV3OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _YRegistryV3.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YRegistryV3OwnershipTransferred)
				if err := _YRegistryV3.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_YRegistryV3 *YRegistryV3Filterer) ParseOwnershipTransferred(log types.Log) (*YRegistryV3OwnershipTransferred, error) {
	event := new(YRegistryV3OwnershipTransferred)
	if err := _YRegistryV3.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YRegistryV3ReleaseRegistryUpdatedIterator is returned from FilterReleaseRegistryUpdated and is used to iterate over the raw logs and unpacked data for ReleaseRegistryUpdated events raised by the YRegistryV3 contract.
type YRegistryV3ReleaseRegistryUpdatedIterator struct {
	Event *YRegistryV3ReleaseRegistryUpdated // Event containing the contract specifics and raw log

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
func (it *YRegistryV3ReleaseRegistryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YRegistryV3ReleaseRegistryUpdated)
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
		it.Event = new(YRegistryV3ReleaseRegistryUpdated)
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
func (it *YRegistryV3ReleaseRegistryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YRegistryV3ReleaseRegistryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YRegistryV3ReleaseRegistryUpdated represents a ReleaseRegistryUpdated event raised by the YRegistryV3 contract.
type YRegistryV3ReleaseRegistryUpdated struct {
	NewRegistry common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReleaseRegistryUpdated is a free log retrieval operation binding the contract event 0x65104c4699f20a240a0ce199d616669f9518c0469b60b2a8c904f7cd4eb471fd.
//
// Solidity: event ReleaseRegistryUpdated(address newRegistry)
func (_YRegistryV3 *YRegistryV3Filterer) FilterReleaseRegistryUpdated(opts *bind.FilterOpts) (*YRegistryV3ReleaseRegistryUpdatedIterator, error) {

	logs, sub, err := _YRegistryV3.contract.FilterLogs(opts, "ReleaseRegistryUpdated")
	if err != nil {
		return nil, err
	}
	return &YRegistryV3ReleaseRegistryUpdatedIterator{contract: _YRegistryV3.contract, event: "ReleaseRegistryUpdated", logs: logs, sub: sub}, nil
}

// WatchReleaseRegistryUpdated is a free log subscription operation binding the contract event 0x65104c4699f20a240a0ce199d616669f9518c0469b60b2a8c904f7cd4eb471fd.
//
// Solidity: event ReleaseRegistryUpdated(address newRegistry)
func (_YRegistryV3 *YRegistryV3Filterer) WatchReleaseRegistryUpdated(opts *bind.WatchOpts, sink chan<- *YRegistryV3ReleaseRegistryUpdated) (event.Subscription, error) {

	logs, sub, err := _YRegistryV3.contract.WatchLogs(opts, "ReleaseRegistryUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YRegistryV3ReleaseRegistryUpdated)
				if err := _YRegistryV3.contract.UnpackLog(event, "ReleaseRegistryUpdated", log); err != nil {
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

// ParseReleaseRegistryUpdated is a log parse operation binding the contract event 0x65104c4699f20a240a0ce199d616669f9518c0469b60b2a8c904f7cd4eb471fd.
//
// Solidity: event ReleaseRegistryUpdated(address newRegistry)
func (_YRegistryV3 *YRegistryV3Filterer) ParseReleaseRegistryUpdated(log types.Log) (*YRegistryV3ReleaseRegistryUpdated, error) {
	event := new(YRegistryV3ReleaseRegistryUpdated)
	if err := _YRegistryV3.contract.UnpackLog(event, "ReleaseRegistryUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
