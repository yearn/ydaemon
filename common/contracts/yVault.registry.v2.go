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

// YRegistryV2MetaData contains all meta data concerning the YRegistryV2 contract.
var YRegistryV2MetaData = &bind.MetaData{
	ABI: "[{\"name\":\"NewRelease\",\"inputs\":[{\"name\":\"release_id\",\"type\":\"uint256\",\"indexed\":true},{\"name\":\"template\",\"type\":\"address\",\"indexed\":false},{\"name\":\"api_version\",\"type\":\"string\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewVault\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true},{\"name\":\"vault_id\",\"type\":\"uint256\",\"indexed\":true},{\"name\":\"vault\",\"type\":\"address\",\"indexed\":false},{\"name\":\"api_version\",\"type\":\"string\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewExperimentalVault\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true},{\"name\":\"deployer\",\"type\":\"address\",\"indexed\":true},{\"name\":\"vault\",\"type\":\"address\",\"indexed\":false},{\"name\":\"api_version\",\"type\":\"string\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewGovernance\",\"inputs\":[{\"name\":\"governance\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"VaultTagged\",\"inputs\":[{\"name\":\"vault\",\"type\":\"address\",\"indexed\":false},{\"name\":\"tag\",\"type\":\"string\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"setGovernance\",\"inputs\":[{\"name\":\"governance\",\"type\":\"address\"}],\"outputs\":[],\"gas\":36245},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"acceptGovernance\",\"inputs\":[],\"outputs\":[],\"gas\":37517},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"latestRelease\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"gas\":6831},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"latestVault\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":2587},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"newRelease\",\"inputs\":[{\"name\":\"vault\",\"type\":\"address\"}],\"outputs\":[],\"gas\":82588},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"newVault\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"guardian\",\"type\":\"address\"},{\"name\":\"rewards\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"newVault\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"guardian\",\"type\":\"address\"},{\"name\":\"rewards\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"releaseDelta\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"newExperimentalVault\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"governance\",\"type\":\"address\"},{\"name\":\"guardian\",\"type\":\"address\"},{\"name\":\"rewards\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"newExperimentalVault\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"governance\",\"type\":\"address\"},{\"name\":\"guardian\",\"type\":\"address\"},{\"name\":\"rewards\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"releaseDelta\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"endorseVault\",\"inputs\":[{\"name\":\"vault\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"endorseVault\",\"inputs\":[{\"name\":\"vault\",\"type\":\"address\"},{\"name\":\"releaseDelta\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"setBanksy\",\"inputs\":[{\"name\":\"tagger\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"setBanksy\",\"inputs\":[{\"name\":\"tagger\",\"type\":\"address\"},{\"name\":\"allowed\",\"type\":\"bool\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"tagVault\",\"inputs\":[{\"name\":\"vault\",\"type\":\"address\"},{\"name\":\"tag\",\"type\":\"string\"}],\"outputs\":[],\"gas\":186064},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"numReleases\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":1388},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"releases\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":1533},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"numVaults\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":1663},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"vaults\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"},{\"name\":\"arg1\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":1808},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"tokens\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":1623},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"numTokens\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":1538},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"isRegistered\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":1783},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"governance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":1598},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"pendingGovernance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":1628},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"tags\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"gas\":10229},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"banksy\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":1903}]",
}

// YRegistryV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use YRegistryV2MetaData.ABI instead.
var YRegistryV2ABI = YRegistryV2MetaData.ABI

// YRegistryV2 is an auto generated Go binding around an Ethereum contract.
type YRegistryV2 struct {
	YRegistryV2Caller     // Read-only binding to the contract
	YRegistryV2Transactor // Write-only binding to the contract
	YRegistryV2Filterer   // Log filterer for contract events
}

// YRegistryV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type YRegistryV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YRegistryV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type YRegistryV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YRegistryV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YRegistryV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YRegistryV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YRegistryV2Session struct {
	Contract     *YRegistryV2      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YRegistryV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YRegistryV2CallerSession struct {
	Contract *YRegistryV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// YRegistryV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YRegistryV2TransactorSession struct {
	Contract     *YRegistryV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// YRegistryV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type YRegistryV2Raw struct {
	Contract *YRegistryV2 // Generic contract binding to access the raw methods on
}

// YRegistryV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YRegistryV2CallerRaw struct {
	Contract *YRegistryV2Caller // Generic read-only contract binding to access the raw methods on
}

// YRegistryV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YRegistryV2TransactorRaw struct {
	Contract *YRegistryV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewYRegistryV2 creates a new instance of YRegistryV2, bound to a specific deployed contract.
func NewYRegistryV2(address common.Address, backend bind.ContractBackend) (*YRegistryV2, error) {
	contract, err := bindYRegistryV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YRegistryV2{YRegistryV2Caller: YRegistryV2Caller{contract: contract}, YRegistryV2Transactor: YRegistryV2Transactor{contract: contract}, YRegistryV2Filterer: YRegistryV2Filterer{contract: contract}}, nil
}

// NewYRegistryV2Caller creates a new read-only instance of YRegistryV2, bound to a specific deployed contract.
func NewYRegistryV2Caller(address common.Address, caller bind.ContractCaller) (*YRegistryV2Caller, error) {
	contract, err := bindYRegistryV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YRegistryV2Caller{contract: contract}, nil
}

// NewYRegistryV2Transactor creates a new write-only instance of YRegistryV2, bound to a specific deployed contract.
func NewYRegistryV2Transactor(address common.Address, transactor bind.ContractTransactor) (*YRegistryV2Transactor, error) {
	contract, err := bindYRegistryV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YRegistryV2Transactor{contract: contract}, nil
}

// NewYRegistryV2Filterer creates a new log filterer instance of YRegistryV2, bound to a specific deployed contract.
func NewYRegistryV2Filterer(address common.Address, filterer bind.ContractFilterer) (*YRegistryV2Filterer, error) {
	contract, err := bindYRegistryV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YRegistryV2Filterer{contract: contract}, nil
}

// bindYRegistryV2 binds a generic wrapper to an already deployed contract.
func bindYRegistryV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(YRegistryV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YRegistryV2 *YRegistryV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YRegistryV2.Contract.YRegistryV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YRegistryV2 *YRegistryV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YRegistryV2.Contract.YRegistryV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YRegistryV2 *YRegistryV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YRegistryV2.Contract.YRegistryV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YRegistryV2 *YRegistryV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YRegistryV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YRegistryV2 *YRegistryV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YRegistryV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YRegistryV2 *YRegistryV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YRegistryV2.Contract.contract.Transact(opts, method, params...)
}

// Banksy is a free data retrieval call binding the contract method 0xee711ed5.
//
// Solidity: function banksy(address arg0) view returns(bool)
func (_YRegistryV2 *YRegistryV2Caller) Banksy(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _YRegistryV2.contract.Call(opts, &out, "banksy", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Banksy is a free data retrieval call binding the contract method 0xee711ed5.
//
// Solidity: function banksy(address arg0) view returns(bool)
func (_YRegistryV2 *YRegistryV2Session) Banksy(arg0 common.Address) (bool, error) {
	return _YRegistryV2.Contract.Banksy(&_YRegistryV2.CallOpts, arg0)
}

// Banksy is a free data retrieval call binding the contract method 0xee711ed5.
//
// Solidity: function banksy(address arg0) view returns(bool)
func (_YRegistryV2 *YRegistryV2CallerSession) Banksy(arg0 common.Address) (bool, error) {
	return _YRegistryV2.Contract.Banksy(&_YRegistryV2.CallOpts, arg0)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_YRegistryV2 *YRegistryV2Caller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YRegistryV2.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_YRegistryV2 *YRegistryV2Session) Governance() (common.Address, error) {
	return _YRegistryV2.Contract.Governance(&_YRegistryV2.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_YRegistryV2 *YRegistryV2CallerSession) Governance() (common.Address, error) {
	return _YRegistryV2.Contract.Governance(&_YRegistryV2.CallOpts)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address arg0) view returns(bool)
func (_YRegistryV2 *YRegistryV2Caller) IsRegistered(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _YRegistryV2.contract.Call(opts, &out, "isRegistered", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address arg0) view returns(bool)
func (_YRegistryV2 *YRegistryV2Session) IsRegistered(arg0 common.Address) (bool, error) {
	return _YRegistryV2.Contract.IsRegistered(&_YRegistryV2.CallOpts, arg0)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address arg0) view returns(bool)
func (_YRegistryV2 *YRegistryV2CallerSession) IsRegistered(arg0 common.Address) (bool, error) {
	return _YRegistryV2.Contract.IsRegistered(&_YRegistryV2.CallOpts, arg0)
}

// LatestRelease is a free data retrieval call binding the contract method 0x7be0ca5e.
//
// Solidity: function latestRelease() view returns(string)
func (_YRegistryV2 *YRegistryV2Caller) LatestRelease(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _YRegistryV2.contract.Call(opts, &out, "latestRelease")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LatestRelease is a free data retrieval call binding the contract method 0x7be0ca5e.
//
// Solidity: function latestRelease() view returns(string)
func (_YRegistryV2 *YRegistryV2Session) LatestRelease() (string, error) {
	return _YRegistryV2.Contract.LatestRelease(&_YRegistryV2.CallOpts)
}

// LatestRelease is a free data retrieval call binding the contract method 0x7be0ca5e.
//
// Solidity: function latestRelease() view returns(string)
func (_YRegistryV2 *YRegistryV2CallerSession) LatestRelease() (string, error) {
	return _YRegistryV2.Contract.LatestRelease(&_YRegistryV2.CallOpts)
}

// LatestVault is a free data retrieval call binding the contract method 0xe177dc70.
//
// Solidity: function latestVault(address token) view returns(address)
func (_YRegistryV2 *YRegistryV2Caller) LatestVault(opts *bind.CallOpts, token common.Address) (common.Address, error) {
	var out []interface{}
	err := _YRegistryV2.contract.Call(opts, &out, "latestVault", token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LatestVault is a free data retrieval call binding the contract method 0xe177dc70.
//
// Solidity: function latestVault(address token) view returns(address)
func (_YRegistryV2 *YRegistryV2Session) LatestVault(token common.Address) (common.Address, error) {
	return _YRegistryV2.Contract.LatestVault(&_YRegistryV2.CallOpts, token)
}

// LatestVault is a free data retrieval call binding the contract method 0xe177dc70.
//
// Solidity: function latestVault(address token) view returns(address)
func (_YRegistryV2 *YRegistryV2CallerSession) LatestVault(token common.Address) (common.Address, error) {
	return _YRegistryV2.Contract.LatestVault(&_YRegistryV2.CallOpts, token)
}

// NumReleases is a free data retrieval call binding the contract method 0x56e0a94b.
//
// Solidity: function numReleases() view returns(uint256)
func (_YRegistryV2 *YRegistryV2Caller) NumReleases(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YRegistryV2.contract.Call(opts, &out, "numReleases")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumReleases is a free data retrieval call binding the contract method 0x56e0a94b.
//
// Solidity: function numReleases() view returns(uint256)
func (_YRegistryV2 *YRegistryV2Session) NumReleases() (*big.Int, error) {
	return _YRegistryV2.Contract.NumReleases(&_YRegistryV2.CallOpts)
}

// NumReleases is a free data retrieval call binding the contract method 0x56e0a94b.
//
// Solidity: function numReleases() view returns(uint256)
func (_YRegistryV2 *YRegistryV2CallerSession) NumReleases() (*big.Int, error) {
	return _YRegistryV2.Contract.NumReleases(&_YRegistryV2.CallOpts)
}

// NumTokens is a free data retrieval call binding the contract method 0x8e499bcf.
//
// Solidity: function numTokens() view returns(uint256)
func (_YRegistryV2 *YRegistryV2Caller) NumTokens(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YRegistryV2.contract.Call(opts, &out, "numTokens")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumTokens is a free data retrieval call binding the contract method 0x8e499bcf.
//
// Solidity: function numTokens() view returns(uint256)
func (_YRegistryV2 *YRegistryV2Session) NumTokens() (*big.Int, error) {
	return _YRegistryV2.Contract.NumTokens(&_YRegistryV2.CallOpts)
}

// NumTokens is a free data retrieval call binding the contract method 0x8e499bcf.
//
// Solidity: function numTokens() view returns(uint256)
func (_YRegistryV2 *YRegistryV2CallerSession) NumTokens() (*big.Int, error) {
	return _YRegistryV2.Contract.NumTokens(&_YRegistryV2.CallOpts)
}

// NumVaults is a free data retrieval call binding the contract method 0xf9c7bba5.
//
// Solidity: function numVaults(address arg0) view returns(uint256)
func (_YRegistryV2 *YRegistryV2Caller) NumVaults(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YRegistryV2.contract.Call(opts, &out, "numVaults", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumVaults is a free data retrieval call binding the contract method 0xf9c7bba5.
//
// Solidity: function numVaults(address arg0) view returns(uint256)
func (_YRegistryV2 *YRegistryV2Session) NumVaults(arg0 common.Address) (*big.Int, error) {
	return _YRegistryV2.Contract.NumVaults(&_YRegistryV2.CallOpts, arg0)
}

// NumVaults is a free data retrieval call binding the contract method 0xf9c7bba5.
//
// Solidity: function numVaults(address arg0) view returns(uint256)
func (_YRegistryV2 *YRegistryV2CallerSession) NumVaults(arg0 common.Address) (*big.Int, error) {
	return _YRegistryV2.Contract.NumVaults(&_YRegistryV2.CallOpts, arg0)
}

// PendingGovernance is a free data retrieval call binding the contract method 0xf39c38a0.
//
// Solidity: function pendingGovernance() view returns(address)
func (_YRegistryV2 *YRegistryV2Caller) PendingGovernance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YRegistryV2.contract.Call(opts, &out, "pendingGovernance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingGovernance is a free data retrieval call binding the contract method 0xf39c38a0.
//
// Solidity: function pendingGovernance() view returns(address)
func (_YRegistryV2 *YRegistryV2Session) PendingGovernance() (common.Address, error) {
	return _YRegistryV2.Contract.PendingGovernance(&_YRegistryV2.CallOpts)
}

// PendingGovernance is a free data retrieval call binding the contract method 0xf39c38a0.
//
// Solidity: function pendingGovernance() view returns(address)
func (_YRegistryV2 *YRegistryV2CallerSession) PendingGovernance() (common.Address, error) {
	return _YRegistryV2.Contract.PendingGovernance(&_YRegistryV2.CallOpts)
}

// Releases is a free data retrieval call binding the contract method 0xb6a9f40f.
//
// Solidity: function releases(uint256 arg0) view returns(address)
func (_YRegistryV2 *YRegistryV2Caller) Releases(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _YRegistryV2.contract.Call(opts, &out, "releases", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Releases is a free data retrieval call binding the contract method 0xb6a9f40f.
//
// Solidity: function releases(uint256 arg0) view returns(address)
func (_YRegistryV2 *YRegistryV2Session) Releases(arg0 *big.Int) (common.Address, error) {
	return _YRegistryV2.Contract.Releases(&_YRegistryV2.CallOpts, arg0)
}

// Releases is a free data retrieval call binding the contract method 0xb6a9f40f.
//
// Solidity: function releases(uint256 arg0) view returns(address)
func (_YRegistryV2 *YRegistryV2CallerSession) Releases(arg0 *big.Int) (common.Address, error) {
	return _YRegistryV2.Contract.Releases(&_YRegistryV2.CallOpts, arg0)
}

// Tags is a free data retrieval call binding the contract method 0x58b8f842.
//
// Solidity: function tags(address arg0) view returns(string)
func (_YRegistryV2 *YRegistryV2Caller) Tags(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _YRegistryV2.contract.Call(opts, &out, "tags", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Tags is a free data retrieval call binding the contract method 0x58b8f842.
//
// Solidity: function tags(address arg0) view returns(string)
func (_YRegistryV2 *YRegistryV2Session) Tags(arg0 common.Address) (string, error) {
	return _YRegistryV2.Contract.Tags(&_YRegistryV2.CallOpts, arg0)
}

// Tags is a free data retrieval call binding the contract method 0x58b8f842.
//
// Solidity: function tags(address arg0) view returns(string)
func (_YRegistryV2 *YRegistryV2CallerSession) Tags(arg0 common.Address) (string, error) {
	return _YRegistryV2.Contract.Tags(&_YRegistryV2.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 arg0) view returns(address)
func (_YRegistryV2 *YRegistryV2Caller) Tokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _YRegistryV2.contract.Call(opts, &out, "tokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 arg0) view returns(address)
func (_YRegistryV2 *YRegistryV2Session) Tokens(arg0 *big.Int) (common.Address, error) {
	return _YRegistryV2.Contract.Tokens(&_YRegistryV2.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 arg0) view returns(address)
func (_YRegistryV2 *YRegistryV2CallerSession) Tokens(arg0 *big.Int) (common.Address, error) {
	return _YRegistryV2.Contract.Tokens(&_YRegistryV2.CallOpts, arg0)
}

// Vaults is a free data retrieval call binding the contract method 0x7bbfc69e.
//
// Solidity: function vaults(address arg0, uint256 arg1) view returns(address)
func (_YRegistryV2 *YRegistryV2Caller) Vaults(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _YRegistryV2.contract.Call(opts, &out, "vaults", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vaults is a free data retrieval call binding the contract method 0x7bbfc69e.
//
// Solidity: function vaults(address arg0, uint256 arg1) view returns(address)
func (_YRegistryV2 *YRegistryV2Session) Vaults(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _YRegistryV2.Contract.Vaults(&_YRegistryV2.CallOpts, arg0, arg1)
}

// Vaults is a free data retrieval call binding the contract method 0x7bbfc69e.
//
// Solidity: function vaults(address arg0, uint256 arg1) view returns(address)
func (_YRegistryV2 *YRegistryV2CallerSession) Vaults(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _YRegistryV2.Contract.Vaults(&_YRegistryV2.CallOpts, arg0, arg1)
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_YRegistryV2 *YRegistryV2Transactor) AcceptGovernance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YRegistryV2.contract.Transact(opts, "acceptGovernance")
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_YRegistryV2 *YRegistryV2Session) AcceptGovernance() (*types.Transaction, error) {
	return _YRegistryV2.Contract.AcceptGovernance(&_YRegistryV2.TransactOpts)
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_YRegistryV2 *YRegistryV2TransactorSession) AcceptGovernance() (*types.Transaction, error) {
	return _YRegistryV2.Contract.AcceptGovernance(&_YRegistryV2.TransactOpts)
}

// EndorseVault is a paid mutator transaction binding the contract method 0x29b2e0c6.
//
// Solidity: function endorseVault(address vault) returns()
func (_YRegistryV2 *YRegistryV2Transactor) EndorseVault(opts *bind.TransactOpts, vault common.Address) (*types.Transaction, error) {
	return _YRegistryV2.contract.Transact(opts, "endorseVault", vault)
}

// EndorseVault is a paid mutator transaction binding the contract method 0x29b2e0c6.
//
// Solidity: function endorseVault(address vault) returns()
func (_YRegistryV2 *YRegistryV2Session) EndorseVault(vault common.Address) (*types.Transaction, error) {
	return _YRegistryV2.Contract.EndorseVault(&_YRegistryV2.TransactOpts, vault)
}

// EndorseVault is a paid mutator transaction binding the contract method 0x29b2e0c6.
//
// Solidity: function endorseVault(address vault) returns()
func (_YRegistryV2 *YRegistryV2TransactorSession) EndorseVault(vault common.Address) (*types.Transaction, error) {
	return _YRegistryV2.Contract.EndorseVault(&_YRegistryV2.TransactOpts, vault)
}

// EndorseVault0 is a paid mutator transaction binding the contract method 0xb366a35c.
//
// Solidity: function endorseVault(address vault, uint256 releaseDelta) returns()
func (_YRegistryV2 *YRegistryV2Transactor) EndorseVault0(opts *bind.TransactOpts, vault common.Address, releaseDelta *big.Int) (*types.Transaction, error) {
	return _YRegistryV2.contract.Transact(opts, "endorseVault0", vault, releaseDelta)
}

// EndorseVault0 is a paid mutator transaction binding the contract method 0xb366a35c.
//
// Solidity: function endorseVault(address vault, uint256 releaseDelta) returns()
func (_YRegistryV2 *YRegistryV2Session) EndorseVault0(vault common.Address, releaseDelta *big.Int) (*types.Transaction, error) {
	return _YRegistryV2.Contract.EndorseVault0(&_YRegistryV2.TransactOpts, vault, releaseDelta)
}

// EndorseVault0 is a paid mutator transaction binding the contract method 0xb366a35c.
//
// Solidity: function endorseVault(address vault, uint256 releaseDelta) returns()
func (_YRegistryV2 *YRegistryV2TransactorSession) EndorseVault0(vault common.Address, releaseDelta *big.Int) (*types.Transaction, error) {
	return _YRegistryV2.Contract.EndorseVault0(&_YRegistryV2.TransactOpts, vault, releaseDelta)
}

// NewExperimentalVault is a paid mutator transaction binding the contract method 0x5b73aa0d.
//
// Solidity: function newExperimentalVault(address token, address governance, address guardian, address rewards, string name, string symbol) returns(address)
func (_YRegistryV2 *YRegistryV2Transactor) NewExperimentalVault(opts *bind.TransactOpts, token common.Address, governance common.Address, guardian common.Address, rewards common.Address, name string, symbol string) (*types.Transaction, error) {
	return _YRegistryV2.contract.Transact(opts, "newExperimentalVault", token, governance, guardian, rewards, name, symbol)
}

// NewExperimentalVault is a paid mutator transaction binding the contract method 0x5b73aa0d.
//
// Solidity: function newExperimentalVault(address token, address governance, address guardian, address rewards, string name, string symbol) returns(address)
func (_YRegistryV2 *YRegistryV2Session) NewExperimentalVault(token common.Address, governance common.Address, guardian common.Address, rewards common.Address, name string, symbol string) (*types.Transaction, error) {
	return _YRegistryV2.Contract.NewExperimentalVault(&_YRegistryV2.TransactOpts, token, governance, guardian, rewards, name, symbol)
}

// NewExperimentalVault is a paid mutator transaction binding the contract method 0x5b73aa0d.
//
// Solidity: function newExperimentalVault(address token, address governance, address guardian, address rewards, string name, string symbol) returns(address)
func (_YRegistryV2 *YRegistryV2TransactorSession) NewExperimentalVault(token common.Address, governance common.Address, guardian common.Address, rewards common.Address, name string, symbol string) (*types.Transaction, error) {
	return _YRegistryV2.Contract.NewExperimentalVault(&_YRegistryV2.TransactOpts, token, governance, guardian, rewards, name, symbol)
}

// NewExperimentalVault0 is a paid mutator transaction binding the contract method 0x5bd4b0f2.
//
// Solidity: function newExperimentalVault(address token, address governance, address guardian, address rewards, string name, string symbol, uint256 releaseDelta) returns(address)
func (_YRegistryV2 *YRegistryV2Transactor) NewExperimentalVault0(opts *bind.TransactOpts, token common.Address, governance common.Address, guardian common.Address, rewards common.Address, name string, symbol string, releaseDelta *big.Int) (*types.Transaction, error) {
	return _YRegistryV2.contract.Transact(opts, "newExperimentalVault0", token, governance, guardian, rewards, name, symbol, releaseDelta)
}

// NewExperimentalVault0 is a paid mutator transaction binding the contract method 0x5bd4b0f2.
//
// Solidity: function newExperimentalVault(address token, address governance, address guardian, address rewards, string name, string symbol, uint256 releaseDelta) returns(address)
func (_YRegistryV2 *YRegistryV2Session) NewExperimentalVault0(token common.Address, governance common.Address, guardian common.Address, rewards common.Address, name string, symbol string, releaseDelta *big.Int) (*types.Transaction, error) {
	return _YRegistryV2.Contract.NewExperimentalVault0(&_YRegistryV2.TransactOpts, token, governance, guardian, rewards, name, symbol, releaseDelta)
}

// NewExperimentalVault0 is a paid mutator transaction binding the contract method 0x5bd4b0f2.
//
// Solidity: function newExperimentalVault(address token, address governance, address guardian, address rewards, string name, string symbol, uint256 releaseDelta) returns(address)
func (_YRegistryV2 *YRegistryV2TransactorSession) NewExperimentalVault0(token common.Address, governance common.Address, guardian common.Address, rewards common.Address, name string, symbol string, releaseDelta *big.Int) (*types.Transaction, error) {
	return _YRegistryV2.Contract.NewExperimentalVault0(&_YRegistryV2.TransactOpts, token, governance, guardian, rewards, name, symbol, releaseDelta)
}

// NewRelease is a paid mutator transaction binding the contract method 0x33990d4b.
//
// Solidity: function newRelease(address vault) returns()
func (_YRegistryV2 *YRegistryV2Transactor) NewRelease(opts *bind.TransactOpts, vault common.Address) (*types.Transaction, error) {
	return _YRegistryV2.contract.Transact(opts, "newRelease", vault)
}

// NewRelease is a paid mutator transaction binding the contract method 0x33990d4b.
//
// Solidity: function newRelease(address vault) returns()
func (_YRegistryV2 *YRegistryV2Session) NewRelease(vault common.Address) (*types.Transaction, error) {
	return _YRegistryV2.Contract.NewRelease(&_YRegistryV2.TransactOpts, vault)
}

// NewRelease is a paid mutator transaction binding the contract method 0x33990d4b.
//
// Solidity: function newRelease(address vault) returns()
func (_YRegistryV2 *YRegistryV2TransactorSession) NewRelease(vault common.Address) (*types.Transaction, error) {
	return _YRegistryV2.Contract.NewRelease(&_YRegistryV2.TransactOpts, vault)
}

// NewVault is a paid mutator transaction binding the contract method 0x108ca11e.
//
// Solidity: function newVault(address token, address guardian, address rewards, string name, string symbol) returns(address)
func (_YRegistryV2 *YRegistryV2Transactor) NewVault(opts *bind.TransactOpts, token common.Address, guardian common.Address, rewards common.Address, name string, symbol string) (*types.Transaction, error) {
	return _YRegistryV2.contract.Transact(opts, "newVault", token, guardian, rewards, name, symbol)
}

// NewVault is a paid mutator transaction binding the contract method 0x108ca11e.
//
// Solidity: function newVault(address token, address guardian, address rewards, string name, string symbol) returns(address)
func (_YRegistryV2 *YRegistryV2Session) NewVault(token common.Address, guardian common.Address, rewards common.Address, name string, symbol string) (*types.Transaction, error) {
	return _YRegistryV2.Contract.NewVault(&_YRegistryV2.TransactOpts, token, guardian, rewards, name, symbol)
}

// NewVault is a paid mutator transaction binding the contract method 0x108ca11e.
//
// Solidity: function newVault(address token, address guardian, address rewards, string name, string symbol) returns(address)
func (_YRegistryV2 *YRegistryV2TransactorSession) NewVault(token common.Address, guardian common.Address, rewards common.Address, name string, symbol string) (*types.Transaction, error) {
	return _YRegistryV2.Contract.NewVault(&_YRegistryV2.TransactOpts, token, guardian, rewards, name, symbol)
}

// NewVault0 is a paid mutator transaction binding the contract method 0xb0b40fce.
//
// Solidity: function newVault(address token, address guardian, address rewards, string name, string symbol, uint256 releaseDelta) returns(address)
func (_YRegistryV2 *YRegistryV2Transactor) NewVault0(opts *bind.TransactOpts, token common.Address, guardian common.Address, rewards common.Address, name string, symbol string, releaseDelta *big.Int) (*types.Transaction, error) {
	return _YRegistryV2.contract.Transact(opts, "newVault0", token, guardian, rewards, name, symbol, releaseDelta)
}

// NewVault0 is a paid mutator transaction binding the contract method 0xb0b40fce.
//
// Solidity: function newVault(address token, address guardian, address rewards, string name, string symbol, uint256 releaseDelta) returns(address)
func (_YRegistryV2 *YRegistryV2Session) NewVault0(token common.Address, guardian common.Address, rewards common.Address, name string, symbol string, releaseDelta *big.Int) (*types.Transaction, error) {
	return _YRegistryV2.Contract.NewVault0(&_YRegistryV2.TransactOpts, token, guardian, rewards, name, symbol, releaseDelta)
}

// NewVault0 is a paid mutator transaction binding the contract method 0xb0b40fce.
//
// Solidity: function newVault(address token, address guardian, address rewards, string name, string symbol, uint256 releaseDelta) returns(address)
func (_YRegistryV2 *YRegistryV2TransactorSession) NewVault0(token common.Address, guardian common.Address, rewards common.Address, name string, symbol string, releaseDelta *big.Int) (*types.Transaction, error) {
	return _YRegistryV2.Contract.NewVault0(&_YRegistryV2.TransactOpts, token, guardian, rewards, name, symbol, releaseDelta)
}

// SetBanksy is a paid mutator transaction binding the contract method 0x2cad8f9f.
//
// Solidity: function setBanksy(address tagger) returns()
func (_YRegistryV2 *YRegistryV2Transactor) SetBanksy(opts *bind.TransactOpts, tagger common.Address) (*types.Transaction, error) {
	return _YRegistryV2.contract.Transact(opts, "setBanksy", tagger)
}

// SetBanksy is a paid mutator transaction binding the contract method 0x2cad8f9f.
//
// Solidity: function setBanksy(address tagger) returns()
func (_YRegistryV2 *YRegistryV2Session) SetBanksy(tagger common.Address) (*types.Transaction, error) {
	return _YRegistryV2.Contract.SetBanksy(&_YRegistryV2.TransactOpts, tagger)
}

// SetBanksy is a paid mutator transaction binding the contract method 0x2cad8f9f.
//
// Solidity: function setBanksy(address tagger) returns()
func (_YRegistryV2 *YRegistryV2TransactorSession) SetBanksy(tagger common.Address) (*types.Transaction, error) {
	return _YRegistryV2.Contract.SetBanksy(&_YRegistryV2.TransactOpts, tagger)
}

// SetBanksy0 is a paid mutator transaction binding the contract method 0x5e05f6af.
//
// Solidity: function setBanksy(address tagger, bool allowed) returns()
func (_YRegistryV2 *YRegistryV2Transactor) SetBanksy0(opts *bind.TransactOpts, tagger common.Address, allowed bool) (*types.Transaction, error) {
	return _YRegistryV2.contract.Transact(opts, "setBanksy0", tagger, allowed)
}

// SetBanksy0 is a paid mutator transaction binding the contract method 0x5e05f6af.
//
// Solidity: function setBanksy(address tagger, bool allowed) returns()
func (_YRegistryV2 *YRegistryV2Session) SetBanksy0(tagger common.Address, allowed bool) (*types.Transaction, error) {
	return _YRegistryV2.Contract.SetBanksy0(&_YRegistryV2.TransactOpts, tagger, allowed)
}

// SetBanksy0 is a paid mutator transaction binding the contract method 0x5e05f6af.
//
// Solidity: function setBanksy(address tagger, bool allowed) returns()
func (_YRegistryV2 *YRegistryV2TransactorSession) SetBanksy0(tagger common.Address, allowed bool) (*types.Transaction, error) {
	return _YRegistryV2.Contract.SetBanksy0(&_YRegistryV2.TransactOpts, tagger, allowed)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address governance) returns()
func (_YRegistryV2 *YRegistryV2Transactor) SetGovernance(opts *bind.TransactOpts, governance common.Address) (*types.Transaction, error) {
	return _YRegistryV2.contract.Transact(opts, "setGovernance", governance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address governance) returns()
func (_YRegistryV2 *YRegistryV2Session) SetGovernance(governance common.Address) (*types.Transaction, error) {
	return _YRegistryV2.Contract.SetGovernance(&_YRegistryV2.TransactOpts, governance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address governance) returns()
func (_YRegistryV2 *YRegistryV2TransactorSession) SetGovernance(governance common.Address) (*types.Transaction, error) {
	return _YRegistryV2.Contract.SetGovernance(&_YRegistryV2.TransactOpts, governance)
}

// TagVault is a paid mutator transaction binding the contract method 0x60bd68f8.
//
// Solidity: function tagVault(address vault, string tag) returns()
func (_YRegistryV2 *YRegistryV2Transactor) TagVault(opts *bind.TransactOpts, vault common.Address, tag string) (*types.Transaction, error) {
	return _YRegistryV2.contract.Transact(opts, "tagVault", vault, tag)
}

// TagVault is a paid mutator transaction binding the contract method 0x60bd68f8.
//
// Solidity: function tagVault(address vault, string tag) returns()
func (_YRegistryV2 *YRegistryV2Session) TagVault(vault common.Address, tag string) (*types.Transaction, error) {
	return _YRegistryV2.Contract.TagVault(&_YRegistryV2.TransactOpts, vault, tag)
}

// TagVault is a paid mutator transaction binding the contract method 0x60bd68f8.
//
// Solidity: function tagVault(address vault, string tag) returns()
func (_YRegistryV2 *YRegistryV2TransactorSession) TagVault(vault common.Address, tag string) (*types.Transaction, error) {
	return _YRegistryV2.Contract.TagVault(&_YRegistryV2.TransactOpts, vault, tag)
}

// YRegistryV2NewExperimentalVaultIterator is returned from FilterNewExperimentalVault and is used to iterate over the raw logs and unpacked data for NewExperimentalVault events raised by the YRegistryV2 contract.
type YRegistryV2NewExperimentalVaultIterator struct {
	Event *YRegistryV2NewExperimentalVault // Event containing the contract specifics and raw log

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
func (it *YRegistryV2NewExperimentalVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YRegistryV2NewExperimentalVault)
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
		it.Event = new(YRegistryV2NewExperimentalVault)
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
func (it *YRegistryV2NewExperimentalVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YRegistryV2NewExperimentalVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YRegistryV2NewExperimentalVault represents a NewExperimentalVault event raised by the YRegistryV2 contract.
type YRegistryV2NewExperimentalVault struct {
	Token      common.Address
	Deployer   common.Address
	Vault      common.Address
	ApiVersion string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewExperimentalVault is a free log retrieval operation binding the contract event 0x57a9cdc2a05e05f66e76769bdbe88e21ec45d9ee0f97d4cb60395d4c75dcbcda.
//
// Solidity: event NewExperimentalVault(address indexed token, address indexed deployer, address vault, string api_version)
func (_YRegistryV2 *YRegistryV2Filterer) FilterNewExperimentalVault(opts *bind.FilterOpts, token []common.Address, deployer []common.Address) (*YRegistryV2NewExperimentalVaultIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var deployerRule []interface{}
	for _, deployerItem := range deployer {
		deployerRule = append(deployerRule, deployerItem)
	}

	logs, sub, err := _YRegistryV2.contract.FilterLogs(opts, "NewExperimentalVault", tokenRule, deployerRule)
	if err != nil {
		return nil, err
	}
	return &YRegistryV2NewExperimentalVaultIterator{contract: _YRegistryV2.contract, event: "NewExperimentalVault", logs: logs, sub: sub}, nil
}

// WatchNewExperimentalVault is a free log subscription operation binding the contract event 0x57a9cdc2a05e05f66e76769bdbe88e21ec45d9ee0f97d4cb60395d4c75dcbcda.
//
// Solidity: event NewExperimentalVault(address indexed token, address indexed deployer, address vault, string api_version)
func (_YRegistryV2 *YRegistryV2Filterer) WatchNewExperimentalVault(opts *bind.WatchOpts, sink chan<- *YRegistryV2NewExperimentalVault, token []common.Address, deployer []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var deployerRule []interface{}
	for _, deployerItem := range deployer {
		deployerRule = append(deployerRule, deployerItem)
	}

	logs, sub, err := _YRegistryV2.contract.WatchLogs(opts, "NewExperimentalVault", tokenRule, deployerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YRegistryV2NewExperimentalVault)
				if err := _YRegistryV2.contract.UnpackLog(event, "NewExperimentalVault", log); err != nil {
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

// ParseNewExperimentalVault is a log parse operation binding the contract event 0x57a9cdc2a05e05f66e76769bdbe88e21ec45d9ee0f97d4cb60395d4c75dcbcda.
//
// Solidity: event NewExperimentalVault(address indexed token, address indexed deployer, address vault, string api_version)
func (_YRegistryV2 *YRegistryV2Filterer) ParseNewExperimentalVault(log types.Log) (*YRegistryV2NewExperimentalVault, error) {
	event := new(YRegistryV2NewExperimentalVault)
	if err := _YRegistryV2.contract.UnpackLog(event, "NewExperimentalVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YRegistryV2NewGovernanceIterator is returned from FilterNewGovernance and is used to iterate over the raw logs and unpacked data for NewGovernance events raised by the YRegistryV2 contract.
type YRegistryV2NewGovernanceIterator struct {
	Event *YRegistryV2NewGovernance // Event containing the contract specifics and raw log

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
func (it *YRegistryV2NewGovernanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YRegistryV2NewGovernance)
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
		it.Event = new(YRegistryV2NewGovernance)
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
func (it *YRegistryV2NewGovernanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YRegistryV2NewGovernanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YRegistryV2NewGovernance represents a NewGovernance event raised by the YRegistryV2 contract.
type YRegistryV2NewGovernance struct {
	Governance common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewGovernance is a free log retrieval operation binding the contract event 0x4f386975ea1c2f7cf1845b08bee00626fbb624ecad16254d63d9bb9ba86526de.
//
// Solidity: event NewGovernance(address governance)
func (_YRegistryV2 *YRegistryV2Filterer) FilterNewGovernance(opts *bind.FilterOpts) (*YRegistryV2NewGovernanceIterator, error) {

	logs, sub, err := _YRegistryV2.contract.FilterLogs(opts, "NewGovernance")
	if err != nil {
		return nil, err
	}
	return &YRegistryV2NewGovernanceIterator{contract: _YRegistryV2.contract, event: "NewGovernance", logs: logs, sub: sub}, nil
}

// WatchNewGovernance is a free log subscription operation binding the contract event 0x4f386975ea1c2f7cf1845b08bee00626fbb624ecad16254d63d9bb9ba86526de.
//
// Solidity: event NewGovernance(address governance)
func (_YRegistryV2 *YRegistryV2Filterer) WatchNewGovernance(opts *bind.WatchOpts, sink chan<- *YRegistryV2NewGovernance) (event.Subscription, error) {

	logs, sub, err := _YRegistryV2.contract.WatchLogs(opts, "NewGovernance")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YRegistryV2NewGovernance)
				if err := _YRegistryV2.contract.UnpackLog(event, "NewGovernance", log); err != nil {
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

// ParseNewGovernance is a log parse operation binding the contract event 0x4f386975ea1c2f7cf1845b08bee00626fbb624ecad16254d63d9bb9ba86526de.
//
// Solidity: event NewGovernance(address governance)
func (_YRegistryV2 *YRegistryV2Filterer) ParseNewGovernance(log types.Log) (*YRegistryV2NewGovernance, error) {
	event := new(YRegistryV2NewGovernance)
	if err := _YRegistryV2.contract.UnpackLog(event, "NewGovernance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YRegistryV2NewReleaseIterator is returned from FilterNewRelease and is used to iterate over the raw logs and unpacked data for NewRelease events raised by the YRegistryV2 contract.
type YRegistryV2NewReleaseIterator struct {
	Event *YRegistryV2NewRelease // Event containing the contract specifics and raw log

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
func (it *YRegistryV2NewReleaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YRegistryV2NewRelease)
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
		it.Event = new(YRegistryV2NewRelease)
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
func (it *YRegistryV2NewReleaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YRegistryV2NewReleaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YRegistryV2NewRelease represents a NewRelease event raised by the YRegistryV2 contract.
type YRegistryV2NewRelease struct {
	ReleaseId  *big.Int
	Template   common.Address
	ApiVersion string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewRelease is a free log retrieval operation binding the contract event 0xa6fbd216b6734f34092f1be6b7247e1551a6d4f2d5000c53721cfdc119a5b7cf.
//
// Solidity: event NewRelease(uint256 indexed release_id, address template, string api_version)
func (_YRegistryV2 *YRegistryV2Filterer) FilterNewRelease(opts *bind.FilterOpts, release_id []*big.Int) (*YRegistryV2NewReleaseIterator, error) {

	var release_idRule []interface{}
	for _, release_idItem := range release_id {
		release_idRule = append(release_idRule, release_idItem)
	}

	logs, sub, err := _YRegistryV2.contract.FilterLogs(opts, "NewRelease", release_idRule)
	if err != nil {
		return nil, err
	}
	return &YRegistryV2NewReleaseIterator{contract: _YRegistryV2.contract, event: "NewRelease", logs: logs, sub: sub}, nil
}

// WatchNewRelease is a free log subscription operation binding the contract event 0xa6fbd216b6734f34092f1be6b7247e1551a6d4f2d5000c53721cfdc119a5b7cf.
//
// Solidity: event NewRelease(uint256 indexed release_id, address template, string api_version)
func (_YRegistryV2 *YRegistryV2Filterer) WatchNewRelease(opts *bind.WatchOpts, sink chan<- *YRegistryV2NewRelease, release_id []*big.Int) (event.Subscription, error) {

	var release_idRule []interface{}
	for _, release_idItem := range release_id {
		release_idRule = append(release_idRule, release_idItem)
	}

	logs, sub, err := _YRegistryV2.contract.WatchLogs(opts, "NewRelease", release_idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YRegistryV2NewRelease)
				if err := _YRegistryV2.contract.UnpackLog(event, "NewRelease", log); err != nil {
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

// ParseNewRelease is a log parse operation binding the contract event 0xa6fbd216b6734f34092f1be6b7247e1551a6d4f2d5000c53721cfdc119a5b7cf.
//
// Solidity: event NewRelease(uint256 indexed release_id, address template, string api_version)
func (_YRegistryV2 *YRegistryV2Filterer) ParseNewRelease(log types.Log) (*YRegistryV2NewRelease, error) {
	event := new(YRegistryV2NewRelease)
	if err := _YRegistryV2.contract.UnpackLog(event, "NewRelease", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YRegistryV2NewVaultIterator is returned from FilterNewVault and is used to iterate over the raw logs and unpacked data for NewVault events raised by the YRegistryV2 contract.
type YRegistryV2NewVaultIterator struct {
	Event *YRegistryV2NewVault // Event containing the contract specifics and raw log

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
func (it *YRegistryV2NewVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YRegistryV2NewVault)
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
		it.Event = new(YRegistryV2NewVault)
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
func (it *YRegistryV2NewVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YRegistryV2NewVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YRegistryV2NewVault represents a NewVault event raised by the YRegistryV2 contract.
type YRegistryV2NewVault struct {
	Token      common.Address
	VaultId    *big.Int
	Vault      common.Address
	ApiVersion string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewVault is a free log retrieval operation binding the contract event 0xce089905ba4a4d622553bcb5646fd23e895c256f0376eee04e99e61cec1dc7e8.
//
// Solidity: event NewVault(address indexed token, uint256 indexed vault_id, address vault, string api_version)
func (_YRegistryV2 *YRegistryV2Filterer) FilterNewVault(opts *bind.FilterOpts, token []common.Address, vault_id []*big.Int) (*YRegistryV2NewVaultIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var vault_idRule []interface{}
	for _, vault_idItem := range vault_id {
		vault_idRule = append(vault_idRule, vault_idItem)
	}

	logs, sub, err := _YRegistryV2.contract.FilterLogs(opts, "NewVault", tokenRule, vault_idRule)
	if err != nil {
		return nil, err
	}
	return &YRegistryV2NewVaultIterator{contract: _YRegistryV2.contract, event: "NewVault", logs: logs, sub: sub}, nil
}

// WatchNewVault is a free log subscription operation binding the contract event 0xce089905ba4a4d622553bcb5646fd23e895c256f0376eee04e99e61cec1dc7e8.
//
// Solidity: event NewVault(address indexed token, uint256 indexed vault_id, address vault, string api_version)
func (_YRegistryV2 *YRegistryV2Filterer) WatchNewVault(opts *bind.WatchOpts, sink chan<- *YRegistryV2NewVault, token []common.Address, vault_id []*big.Int) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var vault_idRule []interface{}
	for _, vault_idItem := range vault_id {
		vault_idRule = append(vault_idRule, vault_idItem)
	}

	logs, sub, err := _YRegistryV2.contract.WatchLogs(opts, "NewVault", tokenRule, vault_idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YRegistryV2NewVault)
				if err := _YRegistryV2.contract.UnpackLog(event, "NewVault", log); err != nil {
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

// ParseNewVault is a log parse operation binding the contract event 0xce089905ba4a4d622553bcb5646fd23e895c256f0376eee04e99e61cec1dc7e8.
//
// Solidity: event NewVault(address indexed token, uint256 indexed vault_id, address vault, string api_version)
func (_YRegistryV2 *YRegistryV2Filterer) ParseNewVault(log types.Log) (*YRegistryV2NewVault, error) {
	event := new(YRegistryV2NewVault)
	if err := _YRegistryV2.contract.UnpackLog(event, "NewVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YRegistryV2VaultTaggedIterator is returned from FilterVaultTagged and is used to iterate over the raw logs and unpacked data for VaultTagged events raised by the YRegistryV2 contract.
type YRegistryV2VaultTaggedIterator struct {
	Event *YRegistryV2VaultTagged // Event containing the contract specifics and raw log

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
func (it *YRegistryV2VaultTaggedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YRegistryV2VaultTagged)
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
		it.Event = new(YRegistryV2VaultTagged)
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
func (it *YRegistryV2VaultTaggedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YRegistryV2VaultTaggedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YRegistryV2VaultTagged represents a VaultTagged event raised by the YRegistryV2 contract.
type YRegistryV2VaultTagged struct {
	Vault common.Address
	Tag   string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterVaultTagged is a free log retrieval operation binding the contract event 0x07bd58794c2ca0ae152f7719eb5f02c654476de972cadec0e8191ae8be42096d.
//
// Solidity: event VaultTagged(address vault, string tag)
func (_YRegistryV2 *YRegistryV2Filterer) FilterVaultTagged(opts *bind.FilterOpts) (*YRegistryV2VaultTaggedIterator, error) {

	logs, sub, err := _YRegistryV2.contract.FilterLogs(opts, "VaultTagged")
	if err != nil {
		return nil, err
	}
	return &YRegistryV2VaultTaggedIterator{contract: _YRegistryV2.contract, event: "VaultTagged", logs: logs, sub: sub}, nil
}

// WatchVaultTagged is a free log subscription operation binding the contract event 0x07bd58794c2ca0ae152f7719eb5f02c654476de972cadec0e8191ae8be42096d.
//
// Solidity: event VaultTagged(address vault, string tag)
func (_YRegistryV2 *YRegistryV2Filterer) WatchVaultTagged(opts *bind.WatchOpts, sink chan<- *YRegistryV2VaultTagged) (event.Subscription, error) {

	logs, sub, err := _YRegistryV2.contract.WatchLogs(opts, "VaultTagged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YRegistryV2VaultTagged)
				if err := _YRegistryV2.contract.UnpackLog(event, "VaultTagged", log); err != nil {
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

// ParseVaultTagged is a log parse operation binding the contract event 0x07bd58794c2ca0ae152f7719eb5f02c654476de972cadec0e8191ae8be42096d.
//
// Solidity: event VaultTagged(address vault, string tag)
func (_YRegistryV2 *YRegistryV2Filterer) ParseVaultTagged(log types.Log) (*YRegistryV2VaultTagged, error) {
	event := new(YRegistryV2VaultTagged)
	if err := _YRegistryV2.contract.UnpackLog(event, "VaultTagged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
