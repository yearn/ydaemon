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

// YVaultsV3APROracleMetaData contains all meta data concerning the YVaultsV3APROracle contract.
var YVaultsV3APROracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"getCurrentApr\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"apr\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"_delta\",\"type\":\"int256\"}],\"name\":\"getExpectedApr\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"apr\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"_debtChange\",\"type\":\"int256\"}],\"name\":\"getStrategyApr\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"oracles\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"}],\"name\":\"weightedApr\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// YVaultsV3APROracleABI is the input ABI used to generate the binding from.
// Deprecated: Use YVaultsV3APROracleMetaData.ABI instead.
var YVaultsV3APROracleABI = YVaultsV3APROracleMetaData.ABI

// YVaultsV3APROracle is an auto generated Go binding around an Ethereum contract.
type YVaultsV3APROracle struct {
	YVaultsV3APROracleCaller     // Read-only binding to the contract
	YVaultsV3APROracleTransactor // Write-only binding to the contract
	YVaultsV3APROracleFilterer   // Log filterer for contract events
}

// YVaultsV3APROracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type YVaultsV3APROracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YVaultsV3APROracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YVaultsV3APROracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YVaultsV3APROracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YVaultsV3APROracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YVaultsV3APROracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YVaultsV3APROracleSession struct {
	Contract     *YVaultsV3APROracle // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// YVaultsV3APROracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YVaultsV3APROracleCallerSession struct {
	Contract *YVaultsV3APROracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// YVaultsV3APROracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YVaultsV3APROracleTransactorSession struct {
	Contract     *YVaultsV3APROracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// YVaultsV3APROracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type YVaultsV3APROracleRaw struct {
	Contract *YVaultsV3APROracle // Generic contract binding to access the raw methods on
}

// YVaultsV3APROracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YVaultsV3APROracleCallerRaw struct {
	Contract *YVaultsV3APROracleCaller // Generic read-only contract binding to access the raw methods on
}

// YVaultsV3APROracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YVaultsV3APROracleTransactorRaw struct {
	Contract *YVaultsV3APROracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYVaultsV3APROracle creates a new instance of YVaultsV3APROracle, bound to a specific deployed contract.
func NewYVaultsV3APROracle(address common.Address, backend bind.ContractBackend) (*YVaultsV3APROracle, error) {
	contract, err := bindYVaultsV3APROracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YVaultsV3APROracle{YVaultsV3APROracleCaller: YVaultsV3APROracleCaller{contract: contract}, YVaultsV3APROracleTransactor: YVaultsV3APROracleTransactor{contract: contract}, YVaultsV3APROracleFilterer: YVaultsV3APROracleFilterer{contract: contract}}, nil
}

// NewYVaultsV3APROracleCaller creates a new read-only instance of YVaultsV3APROracle, bound to a specific deployed contract.
func NewYVaultsV3APROracleCaller(address common.Address, caller bind.ContractCaller) (*YVaultsV3APROracleCaller, error) {
	contract, err := bindYVaultsV3APROracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YVaultsV3APROracleCaller{contract: contract}, nil
}

// NewYVaultsV3APROracleTransactor creates a new write-only instance of YVaultsV3APROracle, bound to a specific deployed contract.
func NewYVaultsV3APROracleTransactor(address common.Address, transactor bind.ContractTransactor) (*YVaultsV3APROracleTransactor, error) {
	contract, err := bindYVaultsV3APROracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YVaultsV3APROracleTransactor{contract: contract}, nil
}

// NewYVaultsV3APROracleFilterer creates a new log filterer instance of YVaultsV3APROracle, bound to a specific deployed contract.
func NewYVaultsV3APROracleFilterer(address common.Address, filterer bind.ContractFilterer) (*YVaultsV3APROracleFilterer, error) {
	contract, err := bindYVaultsV3APROracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YVaultsV3APROracleFilterer{contract: contract}, nil
}

// bindYVaultsV3APROracle binds a generic wrapper to an already deployed contract.
func bindYVaultsV3APROracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := YVaultsV3APROracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YVaultsV3APROracle *YVaultsV3APROracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YVaultsV3APROracle.Contract.YVaultsV3APROracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YVaultsV3APROracle *YVaultsV3APROracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YVaultsV3APROracle.Contract.YVaultsV3APROracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YVaultsV3APROracle *YVaultsV3APROracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YVaultsV3APROracle.Contract.YVaultsV3APROracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YVaultsV3APROracle *YVaultsV3APROracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YVaultsV3APROracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YVaultsV3APROracle *YVaultsV3APROracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YVaultsV3APROracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YVaultsV3APROracle *YVaultsV3APROracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YVaultsV3APROracle.Contract.contract.Transact(opts, method, params...)
}

// GetCurrentApr is a free data retrieval call binding the contract method 0x59d8703d.
//
// Solidity: function getCurrentApr(address _vault) view returns(uint256 apr)
func (_YVaultsV3APROracle *YVaultsV3APROracleCaller) GetCurrentApr(opts *bind.CallOpts, _vault common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YVaultsV3APROracle.contract.Call(opts, &out, "getCurrentApr", _vault)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentApr is a free data retrieval call binding the contract method 0x59d8703d.
//
// Solidity: function getCurrentApr(address _vault) view returns(uint256 apr)
func (_YVaultsV3APROracle *YVaultsV3APROracleSession) GetCurrentApr(_vault common.Address) (*big.Int, error) {
	return _YVaultsV3APROracle.Contract.GetCurrentApr(&_YVaultsV3APROracle.CallOpts, _vault)
}

// GetCurrentApr is a free data retrieval call binding the contract method 0x59d8703d.
//
// Solidity: function getCurrentApr(address _vault) view returns(uint256 apr)
func (_YVaultsV3APROracle *YVaultsV3APROracleCallerSession) GetCurrentApr(_vault common.Address) (*big.Int, error) {
	return _YVaultsV3APROracle.Contract.GetCurrentApr(&_YVaultsV3APROracle.CallOpts, _vault)
}

// GetExpectedApr is a free data retrieval call binding the contract method 0x8c236964.
//
// Solidity: function getExpectedApr(address _vault, int256 _delta) view returns(uint256 apr)
func (_YVaultsV3APROracle *YVaultsV3APROracleCaller) GetExpectedApr(opts *bind.CallOpts, _vault common.Address, _delta *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YVaultsV3APROracle.contract.Call(opts, &out, "getExpectedApr", _vault, _delta)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExpectedApr is a free data retrieval call binding the contract method 0x8c236964.
//
// Solidity: function getExpectedApr(address _vault, int256 _delta) view returns(uint256 apr)
func (_YVaultsV3APROracle *YVaultsV3APROracleSession) GetExpectedApr(_vault common.Address, _delta *big.Int) (*big.Int, error) {
	return _YVaultsV3APROracle.Contract.GetExpectedApr(&_YVaultsV3APROracle.CallOpts, _vault, _delta)
}

// GetExpectedApr is a free data retrieval call binding the contract method 0x8c236964.
//
// Solidity: function getExpectedApr(address _vault, int256 _delta) view returns(uint256 apr)
func (_YVaultsV3APROracle *YVaultsV3APROracleCallerSession) GetExpectedApr(_vault common.Address, _delta *big.Int) (*big.Int, error) {
	return _YVaultsV3APROracle.Contract.GetExpectedApr(&_YVaultsV3APROracle.CallOpts, _vault, _delta)
}

// GetStrategyApr is a free data retrieval call binding the contract method 0x4d060e36.
//
// Solidity: function getStrategyApr(address _strategy, int256 _debtChange) view returns(uint256)
func (_YVaultsV3APROracle *YVaultsV3APROracleCaller) GetStrategyApr(opts *bind.CallOpts, _strategy common.Address, _debtChange *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YVaultsV3APROracle.contract.Call(opts, &out, "getStrategyApr", _strategy, _debtChange)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStrategyApr is a free data retrieval call binding the contract method 0x4d060e36.
//
// Solidity: function getStrategyApr(address _strategy, int256 _debtChange) view returns(uint256)
func (_YVaultsV3APROracle *YVaultsV3APROracleSession) GetStrategyApr(_strategy common.Address, _debtChange *big.Int) (*big.Int, error) {
	return _YVaultsV3APROracle.Contract.GetStrategyApr(&_YVaultsV3APROracle.CallOpts, _strategy, _debtChange)
}

// GetStrategyApr is a free data retrieval call binding the contract method 0x4d060e36.
//
// Solidity: function getStrategyApr(address _strategy, int256 _debtChange) view returns(uint256)
func (_YVaultsV3APROracle *YVaultsV3APROracleCallerSession) GetStrategyApr(_strategy common.Address, _debtChange *big.Int) (*big.Int, error) {
	return _YVaultsV3APROracle.Contract.GetStrategyApr(&_YVaultsV3APROracle.CallOpts, _strategy, _debtChange)
}

// Oracles is a free data retrieval call binding the contract method 0xaddd5099.
//
// Solidity: function oracles(address ) view returns(address)
func (_YVaultsV3APROracle *YVaultsV3APROracleCaller) Oracles(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _YVaultsV3APROracle.contract.Call(opts, &out, "oracles", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracles is a free data retrieval call binding the contract method 0xaddd5099.
//
// Solidity: function oracles(address ) view returns(address)
func (_YVaultsV3APROracle *YVaultsV3APROracleSession) Oracles(arg0 common.Address) (common.Address, error) {
	return _YVaultsV3APROracle.Contract.Oracles(&_YVaultsV3APROracle.CallOpts, arg0)
}

// Oracles is a free data retrieval call binding the contract method 0xaddd5099.
//
// Solidity: function oracles(address ) view returns(address)
func (_YVaultsV3APROracle *YVaultsV3APROracleCallerSession) Oracles(arg0 common.Address) (common.Address, error) {
	return _YVaultsV3APROracle.Contract.Oracles(&_YVaultsV3APROracle.CallOpts, arg0)
}

// WeightedApr is a free data retrieval call binding the contract method 0xf753dd5c.
//
// Solidity: function weightedApr(address _strategy) view returns(uint256)
func (_YVaultsV3APROracle *YVaultsV3APROracleCaller) WeightedApr(opts *bind.CallOpts, _strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YVaultsV3APROracle.contract.Call(opts, &out, "weightedApr", _strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeightedApr is a free data retrieval call binding the contract method 0xf753dd5c.
//
// Solidity: function weightedApr(address _strategy) view returns(uint256)
func (_YVaultsV3APROracle *YVaultsV3APROracleSession) WeightedApr(_strategy common.Address) (*big.Int, error) {
	return _YVaultsV3APROracle.Contract.WeightedApr(&_YVaultsV3APROracle.CallOpts, _strategy)
}

// WeightedApr is a free data retrieval call binding the contract method 0xf753dd5c.
//
// Solidity: function weightedApr(address _strategy) view returns(uint256)
func (_YVaultsV3APROracle *YVaultsV3APROracleCallerSession) WeightedApr(_strategy common.Address) (*big.Int, error) {
	return _YVaultsV3APROracle.Contract.WeightedApr(&_YVaultsV3APROracle.CallOpts, _strategy)
}

// SetOracle is a paid mutator transaction binding the contract method 0x5c38eb3a.
//
// Solidity: function setOracle(address _strategy, address _oracle) returns()
func (_YVaultsV3APROracle *YVaultsV3APROracleTransactor) SetOracle(opts *bind.TransactOpts, _strategy common.Address, _oracle common.Address) (*types.Transaction, error) {
	return _YVaultsV3APROracle.contract.Transact(opts, "setOracle", _strategy, _oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x5c38eb3a.
//
// Solidity: function setOracle(address _strategy, address _oracle) returns()
func (_YVaultsV3APROracle *YVaultsV3APROracleSession) SetOracle(_strategy common.Address, _oracle common.Address) (*types.Transaction, error) {
	return _YVaultsV3APROracle.Contract.SetOracle(&_YVaultsV3APROracle.TransactOpts, _strategy, _oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x5c38eb3a.
//
// Solidity: function setOracle(address _strategy, address _oracle) returns()
func (_YVaultsV3APROracle *YVaultsV3APROracleTransactorSession) SetOracle(_strategy common.Address, _oracle common.Address) (*types.Transaction, error) {
	return _YVaultsV3APROracle.Contract.SetOracle(&_YVaultsV3APROracle.TransactOpts, _strategy, _oracle)
}
