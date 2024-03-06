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

// YRegistryGammaMetaData contains all meta data concerning the YRegistryGamma contract.
var YRegistryGammaMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_management\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_peformanceFeeRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_keeper\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"NewGammaLPCompounder\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"assetToStrategy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"getStrategyByAsset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"}],\"name\":\"isDeployedStrategy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"keeper\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"management\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_PID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_NATIVE\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"newGammaLPCompounder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_PID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_NATIVE\",\"type\":\"address\"},{\"internalType\":\"address[][2]\",\"name\":\"_midRouteNativeToToken0Token1\",\"type\":\"address[][2]\"},{\"internalType\":\"address[]\",\"name\":\"_rewards\",\"type\":\"address[]\"},{\"internalType\":\"address[][]\",\"name\":\"_midRouteRewardToNative\",\"type\":\"address[][]\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"newGammaLPCompounder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"performanceFeeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_keeper\",\"type\":\"address\"}],\"name\":\"setKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_management\",\"type\":\"address\"}],\"name\":\"setManagement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_performanceFeeRecipient\",\"type\":\"address\"}],\"name\":\"setPerformanceFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"}],\"name\":\"setStrategyByAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// YRegistryGammaABI is the input ABI used to generate the binding from.
// Deprecated: Use YRegistryGammaMetaData.ABI instead.
var YRegistryGammaABI = YRegistryGammaMetaData.ABI

// YRegistryGamma is an auto generated Go binding around an Ethereum contract.
type YRegistryGamma struct {
	YRegistryGammaCaller     // Read-only binding to the contract
	YRegistryGammaTransactor // Write-only binding to the contract
	YRegistryGammaFilterer   // Log filterer for contract events
}

// YRegistryGammaCaller is an auto generated read-only Go binding around an Ethereum contract.
type YRegistryGammaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YRegistryGammaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YRegistryGammaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YRegistryGammaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YRegistryGammaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YRegistryGammaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YRegistryGammaSession struct {
	Contract     *YRegistryGamma   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YRegistryGammaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YRegistryGammaCallerSession struct {
	Contract *YRegistryGammaCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// YRegistryGammaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YRegistryGammaTransactorSession struct {
	Contract     *YRegistryGammaTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// YRegistryGammaRaw is an auto generated low-level Go binding around an Ethereum contract.
type YRegistryGammaRaw struct {
	Contract *YRegistryGamma // Generic contract binding to access the raw methods on
}

// YRegistryGammaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YRegistryGammaCallerRaw struct {
	Contract *YRegistryGammaCaller // Generic read-only contract binding to access the raw methods on
}

// YRegistryGammaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YRegistryGammaTransactorRaw struct {
	Contract *YRegistryGammaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYRegistryGamma creates a new instance of YRegistryGamma, bound to a specific deployed contract.
func NewYRegistryGamma(address common.Address, backend bind.ContractBackend) (*YRegistryGamma, error) {
	contract, err := bindYRegistryGamma(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YRegistryGamma{YRegistryGammaCaller: YRegistryGammaCaller{contract: contract}, YRegistryGammaTransactor: YRegistryGammaTransactor{contract: contract}, YRegistryGammaFilterer: YRegistryGammaFilterer{contract: contract}}, nil
}

// NewYRegistryGammaCaller creates a new read-only instance of YRegistryGamma, bound to a specific deployed contract.
func NewYRegistryGammaCaller(address common.Address, caller bind.ContractCaller) (*YRegistryGammaCaller, error) {
	contract, err := bindYRegistryGamma(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YRegistryGammaCaller{contract: contract}, nil
}

// NewYRegistryGammaTransactor creates a new write-only instance of YRegistryGamma, bound to a specific deployed contract.
func NewYRegistryGammaTransactor(address common.Address, transactor bind.ContractTransactor) (*YRegistryGammaTransactor, error) {
	contract, err := bindYRegistryGamma(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YRegistryGammaTransactor{contract: contract}, nil
}

// NewYRegistryGammaFilterer creates a new log filterer instance of YRegistryGamma, bound to a specific deployed contract.
func NewYRegistryGammaFilterer(address common.Address, filterer bind.ContractFilterer) (*YRegistryGammaFilterer, error) {
	contract, err := bindYRegistryGamma(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YRegistryGammaFilterer{contract: contract}, nil
}

// bindYRegistryGamma binds a generic wrapper to an already deployed contract.
func bindYRegistryGamma(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := YRegistryGammaMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YRegistryGamma *YRegistryGammaRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YRegistryGamma.Contract.YRegistryGammaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YRegistryGamma *YRegistryGammaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YRegistryGamma.Contract.YRegistryGammaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YRegistryGamma *YRegistryGammaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YRegistryGamma.Contract.YRegistryGammaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YRegistryGamma *YRegistryGammaCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YRegistryGamma.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YRegistryGamma *YRegistryGammaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YRegistryGamma.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YRegistryGamma *YRegistryGammaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YRegistryGamma.Contract.contract.Transact(opts, method, params...)
}

// AssetToStrategy is a free data retrieval call binding the contract method 0x4a3e6544.
//
// Solidity: function assetToStrategy(address ) view returns(address)
func (_YRegistryGamma *YRegistryGammaCaller) AssetToStrategy(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _YRegistryGamma.contract.Call(opts, &out, "assetToStrategy", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AssetToStrategy is a free data retrieval call binding the contract method 0x4a3e6544.
//
// Solidity: function assetToStrategy(address ) view returns(address)
func (_YRegistryGamma *YRegistryGammaSession) AssetToStrategy(arg0 common.Address) (common.Address, error) {
	return _YRegistryGamma.Contract.AssetToStrategy(&_YRegistryGamma.CallOpts, arg0)
}

// AssetToStrategy is a free data retrieval call binding the contract method 0x4a3e6544.
//
// Solidity: function assetToStrategy(address ) view returns(address)
func (_YRegistryGamma *YRegistryGammaCallerSession) AssetToStrategy(arg0 common.Address) (common.Address, error) {
	return _YRegistryGamma.Contract.AssetToStrategy(&_YRegistryGamma.CallOpts, arg0)
}

// GetStrategyByAsset is a free data retrieval call binding the contract method 0xdb90b1e0.
//
// Solidity: function getStrategyByAsset(address _asset) view returns(address)
func (_YRegistryGamma *YRegistryGammaCaller) GetStrategyByAsset(opts *bind.CallOpts, _asset common.Address) (common.Address, error) {
	var out []interface{}
	err := _YRegistryGamma.contract.Call(opts, &out, "getStrategyByAsset", _asset)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStrategyByAsset is a free data retrieval call binding the contract method 0xdb90b1e0.
//
// Solidity: function getStrategyByAsset(address _asset) view returns(address)
func (_YRegistryGamma *YRegistryGammaSession) GetStrategyByAsset(_asset common.Address) (common.Address, error) {
	return _YRegistryGamma.Contract.GetStrategyByAsset(&_YRegistryGamma.CallOpts, _asset)
}

// GetStrategyByAsset is a free data retrieval call binding the contract method 0xdb90b1e0.
//
// Solidity: function getStrategyByAsset(address _asset) view returns(address)
func (_YRegistryGamma *YRegistryGammaCallerSession) GetStrategyByAsset(_asset common.Address) (common.Address, error) {
	return _YRegistryGamma.Contract.GetStrategyByAsset(&_YRegistryGamma.CallOpts, _asset)
}

// IsDeployedStrategy is a free data retrieval call binding the contract method 0x0d004424.
//
// Solidity: function isDeployedStrategy(address _strategy) view returns(bool)
func (_YRegistryGamma *YRegistryGammaCaller) IsDeployedStrategy(opts *bind.CallOpts, _strategy common.Address) (bool, error) {
	var out []interface{}
	err := _YRegistryGamma.contract.Call(opts, &out, "isDeployedStrategy", _strategy)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDeployedStrategy is a free data retrieval call binding the contract method 0x0d004424.
//
// Solidity: function isDeployedStrategy(address _strategy) view returns(bool)
func (_YRegistryGamma *YRegistryGammaSession) IsDeployedStrategy(_strategy common.Address) (bool, error) {
	return _YRegistryGamma.Contract.IsDeployedStrategy(&_YRegistryGamma.CallOpts, _strategy)
}

// IsDeployedStrategy is a free data retrieval call binding the contract method 0x0d004424.
//
// Solidity: function isDeployedStrategy(address _strategy) view returns(bool)
func (_YRegistryGamma *YRegistryGammaCallerSession) IsDeployedStrategy(_strategy common.Address) (bool, error) {
	return _YRegistryGamma.Contract.IsDeployedStrategy(&_YRegistryGamma.CallOpts, _strategy)
}

// Keeper is a free data retrieval call binding the contract method 0xaced1661.
//
// Solidity: function keeper() view returns(address)
func (_YRegistryGamma *YRegistryGammaCaller) Keeper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YRegistryGamma.contract.Call(opts, &out, "keeper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Keeper is a free data retrieval call binding the contract method 0xaced1661.
//
// Solidity: function keeper() view returns(address)
func (_YRegistryGamma *YRegistryGammaSession) Keeper() (common.Address, error) {
	return _YRegistryGamma.Contract.Keeper(&_YRegistryGamma.CallOpts)
}

// Keeper is a free data retrieval call binding the contract method 0xaced1661.
//
// Solidity: function keeper() view returns(address)
func (_YRegistryGamma *YRegistryGammaCallerSession) Keeper() (common.Address, error) {
	return _YRegistryGamma.Contract.Keeper(&_YRegistryGamma.CallOpts)
}

// Management is a free data retrieval call binding the contract method 0x88a8d602.
//
// Solidity: function management() view returns(address)
func (_YRegistryGamma *YRegistryGammaCaller) Management(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YRegistryGamma.contract.Call(opts, &out, "management")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Management is a free data retrieval call binding the contract method 0x88a8d602.
//
// Solidity: function management() view returns(address)
func (_YRegistryGamma *YRegistryGammaSession) Management() (common.Address, error) {
	return _YRegistryGamma.Contract.Management(&_YRegistryGamma.CallOpts)
}

// Management is a free data retrieval call binding the contract method 0x88a8d602.
//
// Solidity: function management() view returns(address)
func (_YRegistryGamma *YRegistryGammaCallerSession) Management() (common.Address, error) {
	return _YRegistryGamma.Contract.Management(&_YRegistryGamma.CallOpts)
}

// PerformanceFeeRecipient is a free data retrieval call binding the contract method 0xed27f7c9.
//
// Solidity: function performanceFeeRecipient() view returns(address)
func (_YRegistryGamma *YRegistryGammaCaller) PerformanceFeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YRegistryGamma.contract.Call(opts, &out, "performanceFeeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PerformanceFeeRecipient is a free data retrieval call binding the contract method 0xed27f7c9.
//
// Solidity: function performanceFeeRecipient() view returns(address)
func (_YRegistryGamma *YRegistryGammaSession) PerformanceFeeRecipient() (common.Address, error) {
	return _YRegistryGamma.Contract.PerformanceFeeRecipient(&_YRegistryGamma.CallOpts)
}

// PerformanceFeeRecipient is a free data retrieval call binding the contract method 0xed27f7c9.
//
// Solidity: function performanceFeeRecipient() view returns(address)
func (_YRegistryGamma *YRegistryGammaCallerSession) PerformanceFeeRecipient() (common.Address, error) {
	return _YRegistryGamma.Contract.PerformanceFeeRecipient(&_YRegistryGamma.CallOpts)
}

// NewGammaLPCompounder is a paid mutator transaction binding the contract method 0xb8d7dc9c.
//
// Solidity: function newGammaLPCompounder(address _asset, uint256 _PID, address _NATIVE, string _name) returns(address)
func (_YRegistryGamma *YRegistryGammaTransactor) NewGammaLPCompounder(opts *bind.TransactOpts, _asset common.Address, _PID *big.Int, _NATIVE common.Address, _name string) (*types.Transaction, error) {
	return _YRegistryGamma.contract.Transact(opts, "newGammaLPCompounder", _asset, _PID, _NATIVE, _name)
}

// NewGammaLPCompounder is a paid mutator transaction binding the contract method 0xb8d7dc9c.
//
// Solidity: function newGammaLPCompounder(address _asset, uint256 _PID, address _NATIVE, string _name) returns(address)
func (_YRegistryGamma *YRegistryGammaSession) NewGammaLPCompounder(_asset common.Address, _PID *big.Int, _NATIVE common.Address, _name string) (*types.Transaction, error) {
	return _YRegistryGamma.Contract.NewGammaLPCompounder(&_YRegistryGamma.TransactOpts, _asset, _PID, _NATIVE, _name)
}

// NewGammaLPCompounder is a paid mutator transaction binding the contract method 0xb8d7dc9c.
//
// Solidity: function newGammaLPCompounder(address _asset, uint256 _PID, address _NATIVE, string _name) returns(address)
func (_YRegistryGamma *YRegistryGammaTransactorSession) NewGammaLPCompounder(_asset common.Address, _PID *big.Int, _NATIVE common.Address, _name string) (*types.Transaction, error) {
	return _YRegistryGamma.Contract.NewGammaLPCompounder(&_YRegistryGamma.TransactOpts, _asset, _PID, _NATIVE, _name)
}

// NewGammaLPCompounder0 is a paid mutator transaction binding the contract method 0xe27a5281.
//
// Solidity: function newGammaLPCompounder(address _asset, uint256 _PID, address _NATIVE, address[][2] _midRouteNativeToToken0Token1, address[] _rewards, address[][] _midRouteRewardToNative, string _name) returns(address)
func (_YRegistryGamma *YRegistryGammaTransactor) NewGammaLPCompounder0(opts *bind.TransactOpts, _asset common.Address, _PID *big.Int, _NATIVE common.Address, _midRouteNativeToToken0Token1 [2][]common.Address, _rewards []common.Address, _midRouteRewardToNative [][]common.Address, _name string) (*types.Transaction, error) {
	return _YRegistryGamma.contract.Transact(opts, "newGammaLPCompounder0", _asset, _PID, _NATIVE, _midRouteNativeToToken0Token1, _rewards, _midRouteRewardToNative, _name)
}

// NewGammaLPCompounder0 is a paid mutator transaction binding the contract method 0xe27a5281.
//
// Solidity: function newGammaLPCompounder(address _asset, uint256 _PID, address _NATIVE, address[][2] _midRouteNativeToToken0Token1, address[] _rewards, address[][] _midRouteRewardToNative, string _name) returns(address)
func (_YRegistryGamma *YRegistryGammaSession) NewGammaLPCompounder0(_asset common.Address, _PID *big.Int, _NATIVE common.Address, _midRouteNativeToToken0Token1 [2][]common.Address, _rewards []common.Address, _midRouteRewardToNative [][]common.Address, _name string) (*types.Transaction, error) {
	return _YRegistryGamma.Contract.NewGammaLPCompounder0(&_YRegistryGamma.TransactOpts, _asset, _PID, _NATIVE, _midRouteNativeToToken0Token1, _rewards, _midRouteRewardToNative, _name)
}

// NewGammaLPCompounder0 is a paid mutator transaction binding the contract method 0xe27a5281.
//
// Solidity: function newGammaLPCompounder(address _asset, uint256 _PID, address _NATIVE, address[][2] _midRouteNativeToToken0Token1, address[] _rewards, address[][] _midRouteRewardToNative, string _name) returns(address)
func (_YRegistryGamma *YRegistryGammaTransactorSession) NewGammaLPCompounder0(_asset common.Address, _PID *big.Int, _NATIVE common.Address, _midRouteNativeToToken0Token1 [2][]common.Address, _rewards []common.Address, _midRouteRewardToNative [][]common.Address, _name string) (*types.Transaction, error) {
	return _YRegistryGamma.Contract.NewGammaLPCompounder0(&_YRegistryGamma.TransactOpts, _asset, _PID, _NATIVE, _midRouteNativeToToken0Token1, _rewards, _midRouteRewardToNative, _name)
}

// SetKeeper is a paid mutator transaction binding the contract method 0x748747e6.
//
// Solidity: function setKeeper(address _keeper) returns()
func (_YRegistryGamma *YRegistryGammaTransactor) SetKeeper(opts *bind.TransactOpts, _keeper common.Address) (*types.Transaction, error) {
	return _YRegistryGamma.contract.Transact(opts, "setKeeper", _keeper)
}

// SetKeeper is a paid mutator transaction binding the contract method 0x748747e6.
//
// Solidity: function setKeeper(address _keeper) returns()
func (_YRegistryGamma *YRegistryGammaSession) SetKeeper(_keeper common.Address) (*types.Transaction, error) {
	return _YRegistryGamma.Contract.SetKeeper(&_YRegistryGamma.TransactOpts, _keeper)
}

// SetKeeper is a paid mutator transaction binding the contract method 0x748747e6.
//
// Solidity: function setKeeper(address _keeper) returns()
func (_YRegistryGamma *YRegistryGammaTransactorSession) SetKeeper(_keeper common.Address) (*types.Transaction, error) {
	return _YRegistryGamma.Contract.SetKeeper(&_YRegistryGamma.TransactOpts, _keeper)
}

// SetManagement is a paid mutator transaction binding the contract method 0xd4a22bde.
//
// Solidity: function setManagement(address _management) returns()
func (_YRegistryGamma *YRegistryGammaTransactor) SetManagement(opts *bind.TransactOpts, _management common.Address) (*types.Transaction, error) {
	return _YRegistryGamma.contract.Transact(opts, "setManagement", _management)
}

// SetManagement is a paid mutator transaction binding the contract method 0xd4a22bde.
//
// Solidity: function setManagement(address _management) returns()
func (_YRegistryGamma *YRegistryGammaSession) SetManagement(_management common.Address) (*types.Transaction, error) {
	return _YRegistryGamma.Contract.SetManagement(&_YRegistryGamma.TransactOpts, _management)
}

// SetManagement is a paid mutator transaction binding the contract method 0xd4a22bde.
//
// Solidity: function setManagement(address _management) returns()
func (_YRegistryGamma *YRegistryGammaTransactorSession) SetManagement(_management common.Address) (*types.Transaction, error) {
	return _YRegistryGamma.Contract.SetManagement(&_YRegistryGamma.TransactOpts, _management)
}

// SetPerformanceFeeRecipient is a paid mutator transaction binding the contract method 0x6a5f1aa2.
//
// Solidity: function setPerformanceFeeRecipient(address _performanceFeeRecipient) returns()
func (_YRegistryGamma *YRegistryGammaTransactor) SetPerformanceFeeRecipient(opts *bind.TransactOpts, _performanceFeeRecipient common.Address) (*types.Transaction, error) {
	return _YRegistryGamma.contract.Transact(opts, "setPerformanceFeeRecipient", _performanceFeeRecipient)
}

// SetPerformanceFeeRecipient is a paid mutator transaction binding the contract method 0x6a5f1aa2.
//
// Solidity: function setPerformanceFeeRecipient(address _performanceFeeRecipient) returns()
func (_YRegistryGamma *YRegistryGammaSession) SetPerformanceFeeRecipient(_performanceFeeRecipient common.Address) (*types.Transaction, error) {
	return _YRegistryGamma.Contract.SetPerformanceFeeRecipient(&_YRegistryGamma.TransactOpts, _performanceFeeRecipient)
}

// SetPerformanceFeeRecipient is a paid mutator transaction binding the contract method 0x6a5f1aa2.
//
// Solidity: function setPerformanceFeeRecipient(address _performanceFeeRecipient) returns()
func (_YRegistryGamma *YRegistryGammaTransactorSession) SetPerformanceFeeRecipient(_performanceFeeRecipient common.Address) (*types.Transaction, error) {
	return _YRegistryGamma.Contract.SetPerformanceFeeRecipient(&_YRegistryGamma.TransactOpts, _performanceFeeRecipient)
}

// SetStrategyByAsset is a paid mutator transaction binding the contract method 0x97256094.
//
// Solidity: function setStrategyByAsset(address _asset, address _strategy) returns()
func (_YRegistryGamma *YRegistryGammaTransactor) SetStrategyByAsset(opts *bind.TransactOpts, _asset common.Address, _strategy common.Address) (*types.Transaction, error) {
	return _YRegistryGamma.contract.Transact(opts, "setStrategyByAsset", _asset, _strategy)
}

// SetStrategyByAsset is a paid mutator transaction binding the contract method 0x97256094.
//
// Solidity: function setStrategyByAsset(address _asset, address _strategy) returns()
func (_YRegistryGamma *YRegistryGammaSession) SetStrategyByAsset(_asset common.Address, _strategy common.Address) (*types.Transaction, error) {
	return _YRegistryGamma.Contract.SetStrategyByAsset(&_YRegistryGamma.TransactOpts, _asset, _strategy)
}

// SetStrategyByAsset is a paid mutator transaction binding the contract method 0x97256094.
//
// Solidity: function setStrategyByAsset(address _asset, address _strategy) returns()
func (_YRegistryGamma *YRegistryGammaTransactorSession) SetStrategyByAsset(_asset common.Address, _strategy common.Address) (*types.Transaction, error) {
	return _YRegistryGamma.Contract.SetStrategyByAsset(&_YRegistryGamma.TransactOpts, _asset, _strategy)
}

// YRegistryGammaNewGammaLPCompounderIterator is returned from FilterNewGammaLPCompounder and is used to iterate over the raw logs and unpacked data for NewGammaLPCompounder events raised by the YRegistryGamma contract.
type YRegistryGammaNewGammaLPCompounderIterator struct {
	Event *YRegistryGammaNewGammaLPCompounder // Event containing the contract specifics and raw log

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
func (it *YRegistryGammaNewGammaLPCompounderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YRegistryGammaNewGammaLPCompounder)
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
		it.Event = new(YRegistryGammaNewGammaLPCompounder)
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
func (it *YRegistryGammaNewGammaLPCompounderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YRegistryGammaNewGammaLPCompounderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YRegistryGammaNewGammaLPCompounder represents a NewGammaLPCompounder event raised by the YRegistryGamma contract.
type YRegistryGammaNewGammaLPCompounder struct {
	Strategy common.Address
	Asset    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewGammaLPCompounder is a free log retrieval operation binding the contract event 0x44dd64e5186bc028c5485c2eeb4994c4c116dcf53075a5da69ad4debfc67565b.
//
// Solidity: event NewGammaLPCompounder(address indexed strategy, address indexed asset)
func (_YRegistryGamma *YRegistryGammaFilterer) FilterNewGammaLPCompounder(opts *bind.FilterOpts, strategy []common.Address, asset []common.Address) (*YRegistryGammaNewGammaLPCompounderIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _YRegistryGamma.contract.FilterLogs(opts, "NewGammaLPCompounder", strategyRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &YRegistryGammaNewGammaLPCompounderIterator{contract: _YRegistryGamma.contract, event: "NewGammaLPCompounder", logs: logs, sub: sub}, nil
}

// WatchNewGammaLPCompounder is a free log subscription operation binding the contract event 0x44dd64e5186bc028c5485c2eeb4994c4c116dcf53075a5da69ad4debfc67565b.
//
// Solidity: event NewGammaLPCompounder(address indexed strategy, address indexed asset)
func (_YRegistryGamma *YRegistryGammaFilterer) WatchNewGammaLPCompounder(opts *bind.WatchOpts, sink chan<- *YRegistryGammaNewGammaLPCompounder, strategy []common.Address, asset []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _YRegistryGamma.contract.WatchLogs(opts, "NewGammaLPCompounder", strategyRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YRegistryGammaNewGammaLPCompounder)
				if err := _YRegistryGamma.contract.UnpackLog(event, "NewGammaLPCompounder", log); err != nil {
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

// ParseNewGammaLPCompounder is a log parse operation binding the contract event 0x44dd64e5186bc028c5485c2eeb4994c4c116dcf53075a5da69ad4debfc67565b.
//
// Solidity: event NewGammaLPCompounder(address indexed strategy, address indexed asset)
func (_YRegistryGamma *YRegistryGammaFilterer) ParseNewGammaLPCompounder(log types.Log) (*YRegistryGammaNewGammaLPCompounder, error) {
	event := new(YRegistryGammaNewGammaLPCompounder)
	if err := _YRegistryGamma.contract.UnpackLog(event, "NewGammaLPCompounder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
