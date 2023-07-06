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

// StrategyParams is an auto generated low-level Go binding around an user-defined struct.
type StrategyParams struct {
	PerformanceFee *big.Int
	Activation     *big.Int
	DebtRatio      *big.Int
	RateLimit      *big.Int
	LastReport     *big.Int
	TotalDebt      *big.Int
	TotalGain      *big.Int
	TotalLoss      *big.Int
}

// BaseStrategyMetaData contains all meta data concerning the BaseStrategy contract.
var BaseStrategyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyExitEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"profit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loss\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debtPayment\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debtOutstanding\",\"type\":\"uint256\"}],\"name\":\"Harvested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debtThreshold\",\"type\":\"uint256\"}],\"name\":\"UpdatedDebtThreshold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newKeeper\",\"type\":\"address\"}],\"name\":\"UpdatedKeeper\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"profitFactor\",\"type\":\"uint256\"}],\"name\":\"UpdatedProfitFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"UpdatedReportDelay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rewards\",\"type\":\"address\"}],\"name\":\"UpdatedRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newStrategist\",\"type\":\"address\"}],\"name\":\"UpdatedStrategist\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"apiVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"debtThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegatedAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"doHealthCheck\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyExit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"estimatedTotalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"harvest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"callCost\",\"type\":\"uint256\"}],\"name\":\"harvestTrigger\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"healthCheck\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"keeper\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxReportDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newStrategy\",\"type\":\"address\"}],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"profitFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_debtThreshold\",\"type\":\"uint256\"}],\"name\":\"setDebtThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_doHealthCheck\",\"type\":\"bool\"}],\"name\":\"setDoHealthCheck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setEmergencyExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_healthCheck\",\"type\":\"address\"}],\"name\":\"setHealthCheck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_keeper\",\"type\":\"address\"}],\"name\":\"setKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_delay\",\"type\":\"uint256\"}],\"name\":\"setMaxReportDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_profitFactor\",\"type\":\"uint256\"}],\"name\":\"setProfitFactor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewards\",\"type\":\"address\"}],\"name\":\"setRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategist\",\"type\":\"address\"}],\"name\":\"setStrategist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"strategist\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"sweep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"callCost\",\"type\":\"uint256\"}],\"name\":\"tendTrigger\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"contractVaultAPI\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"want\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountNeeded\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_loss\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"25829410": "apiVersion()",
		"1d12f28b": "debtThreshold()",
		"8e6350e2": "delegatedAssets()",
		"6718835f": "doHealthCheck()",
		"5641ec03": "emergencyExit()",
		"efbb5cb0": "estimatedTotalAssets()",
		"4641257d": "harvest()",
		"ed882c2b": "harvestTrigger(uint256)",
		"b252720b": "healthCheck()",
		"22f3e2d4": "isActive()",
		"aced1661": "keeper()",
		"28b7ccf7": "maxReportDelay()",
		"ce5494bb": "migrate(address)",
		"06fdde03": "name()",
		"8cdfe166": "profitFactor()",
		"9ec5a894": "rewards()",
		"0f969b87": "setDebtThreshold(uint256)",
		"ac00ff26": "setDoHealthCheck(bool)",
		"fcf2d0ad": "setEmergencyExit()",
		"11bc8245": "setHealthCheck(address)",
		"748747e6": "setKeeper(address)",
		"f017c92f": "setMaxReportDelay(uint256)",
		"91397ab4": "setProfitFactor(uint256)",
		"ec38a862": "setRewards(address)",
		"c7b9d530": "setStrategist(address)",
		"1fe4a686": "strategist()",
		"01681a62": "sweep(address)",
		"440368a3": "tend()",
		"650d1880": "tendTrigger(uint256)",
		"fbfa77cf": "vault()",
		"1f1fcd51": "want()",
		"2e1a7d4d": "withdraw(uint256)",
	},
}

// BaseStrategyABI is the input ABI used to generate the binding from.
// Deprecated: Use BaseStrategyMetaData.ABI instead.
var BaseStrategyABI = BaseStrategyMetaData.ABI

// Deprecated: Use BaseStrategyMetaData.Sigs instead.
// BaseStrategyFuncSigs maps the 4-byte function signature to its string representation.
var BaseStrategyFuncSigs = BaseStrategyMetaData.Sigs

// BaseStrategy is an auto generated Go binding around an Ethereum contract.
type BaseStrategy struct {
	BaseStrategyCaller     // Read-only binding to the contract
	BaseStrategyTransactor // Write-only binding to the contract
	BaseStrategyFilterer   // Log filterer for contract events
}

// BaseStrategyCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaseStrategyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseStrategyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaseStrategyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseStrategyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaseStrategyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseStrategySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaseStrategySession struct {
	Contract     *BaseStrategy     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaseStrategyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaseStrategyCallerSession struct {
	Contract *BaseStrategyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BaseStrategyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaseStrategyTransactorSession struct {
	Contract     *BaseStrategyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BaseStrategyRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaseStrategyRaw struct {
	Contract *BaseStrategy // Generic contract binding to access the raw methods on
}

// BaseStrategyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaseStrategyCallerRaw struct {
	Contract *BaseStrategyCaller // Generic read-only contract binding to access the raw methods on
}

// BaseStrategyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaseStrategyTransactorRaw struct {
	Contract *BaseStrategyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaseStrategy creates a new instance of BaseStrategy, bound to a specific deployed contract.
func NewBaseStrategy(address common.Address, backend bind.ContractBackend) (*BaseStrategy, error) {
	contract, err := bindBaseStrategy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaseStrategy{BaseStrategyCaller: BaseStrategyCaller{contract: contract}, BaseStrategyTransactor: BaseStrategyTransactor{contract: contract}, BaseStrategyFilterer: BaseStrategyFilterer{contract: contract}}, nil
}

// NewBaseStrategyCaller creates a new read-only instance of BaseStrategy, bound to a specific deployed contract.
func NewBaseStrategyCaller(address common.Address, caller bind.ContractCaller) (*BaseStrategyCaller, error) {
	contract, err := bindBaseStrategy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaseStrategyCaller{contract: contract}, nil
}

// NewBaseStrategyTransactor creates a new write-only instance of BaseStrategy, bound to a specific deployed contract.
func NewBaseStrategyTransactor(address common.Address, transactor bind.ContractTransactor) (*BaseStrategyTransactor, error) {
	contract, err := bindBaseStrategy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaseStrategyTransactor{contract: contract}, nil
}

// NewBaseStrategyFilterer creates a new log filterer instance of BaseStrategy, bound to a specific deployed contract.
func NewBaseStrategyFilterer(address common.Address, filterer bind.ContractFilterer) (*BaseStrategyFilterer, error) {
	contract, err := bindBaseStrategy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaseStrategyFilterer{contract: contract}, nil
}

// bindBaseStrategy binds a generic wrapper to an already deployed contract.
func bindBaseStrategy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BaseStrategyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseStrategy *BaseStrategyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseStrategy.Contract.BaseStrategyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseStrategy *BaseStrategyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseStrategy.Contract.BaseStrategyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseStrategy *BaseStrategyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseStrategy.Contract.BaseStrategyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseStrategy *BaseStrategyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseStrategy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseStrategy *BaseStrategyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseStrategy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseStrategy *BaseStrategyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseStrategy.Contract.contract.Transact(opts, method, params...)
}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() pure returns(string)
func (_BaseStrategy *BaseStrategyCaller) ApiVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BaseStrategy.contract.Call(opts, &out, "apiVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() pure returns(string)
func (_BaseStrategy *BaseStrategySession) ApiVersion() (string, error) {
	return _BaseStrategy.Contract.ApiVersion(&_BaseStrategy.CallOpts)
}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() pure returns(string)
func (_BaseStrategy *BaseStrategyCallerSession) ApiVersion() (string, error) {
	return _BaseStrategy.Contract.ApiVersion(&_BaseStrategy.CallOpts)
}

// DebtThreshold is a free data retrieval call binding the contract method 0x1d12f28b.
//
// Solidity: function debtThreshold() view returns(uint256)
func (_BaseStrategy *BaseStrategyCaller) DebtThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BaseStrategy.contract.Call(opts, &out, "debtThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtThreshold is a free data retrieval call binding the contract method 0x1d12f28b.
//
// Solidity: function debtThreshold() view returns(uint256)
func (_BaseStrategy *BaseStrategySession) DebtThreshold() (*big.Int, error) {
	return _BaseStrategy.Contract.DebtThreshold(&_BaseStrategy.CallOpts)
}

// DebtThreshold is a free data retrieval call binding the contract method 0x1d12f28b.
//
// Solidity: function debtThreshold() view returns(uint256)
func (_BaseStrategy *BaseStrategyCallerSession) DebtThreshold() (*big.Int, error) {
	return _BaseStrategy.Contract.DebtThreshold(&_BaseStrategy.CallOpts)
}

// DelegatedAssets is a free data retrieval call binding the contract method 0x8e6350e2.
//
// Solidity: function delegatedAssets() view returns(uint256)
func (_BaseStrategy *BaseStrategyCaller) DelegatedAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BaseStrategy.contract.Call(opts, &out, "delegatedAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegatedAssets is a free data retrieval call binding the contract method 0x8e6350e2.
//
// Solidity: function delegatedAssets() view returns(uint256)
func (_BaseStrategy *BaseStrategySession) DelegatedAssets() (*big.Int, error) {
	return _BaseStrategy.Contract.DelegatedAssets(&_BaseStrategy.CallOpts)
}

// DelegatedAssets is a free data retrieval call binding the contract method 0x8e6350e2.
//
// Solidity: function delegatedAssets() view returns(uint256)
func (_BaseStrategy *BaseStrategyCallerSession) DelegatedAssets() (*big.Int, error) {
	return _BaseStrategy.Contract.DelegatedAssets(&_BaseStrategy.CallOpts)
}

// DoHealthCheck is a free data retrieval call binding the contract method 0x6718835f.
//
// Solidity: function doHealthCheck() view returns(bool)
func (_BaseStrategy *BaseStrategyCaller) DoHealthCheck(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BaseStrategy.contract.Call(opts, &out, "doHealthCheck")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DoHealthCheck is a free data retrieval call binding the contract method 0x6718835f.
//
// Solidity: function doHealthCheck() view returns(bool)
func (_BaseStrategy *BaseStrategySession) DoHealthCheck() (bool, error) {
	return _BaseStrategy.Contract.DoHealthCheck(&_BaseStrategy.CallOpts)
}

// DoHealthCheck is a free data retrieval call binding the contract method 0x6718835f.
//
// Solidity: function doHealthCheck() view returns(bool)
func (_BaseStrategy *BaseStrategyCallerSession) DoHealthCheck() (bool, error) {
	return _BaseStrategy.Contract.DoHealthCheck(&_BaseStrategy.CallOpts)
}

// EmergencyExit is a free data retrieval call binding the contract method 0x5641ec03.
//
// Solidity: function emergencyExit() view returns(bool)
func (_BaseStrategy *BaseStrategyCaller) EmergencyExit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BaseStrategy.contract.Call(opts, &out, "emergencyExit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EmergencyExit is a free data retrieval call binding the contract method 0x5641ec03.
//
// Solidity: function emergencyExit() view returns(bool)
func (_BaseStrategy *BaseStrategySession) EmergencyExit() (bool, error) {
	return _BaseStrategy.Contract.EmergencyExit(&_BaseStrategy.CallOpts)
}

// EmergencyExit is a free data retrieval call binding the contract method 0x5641ec03.
//
// Solidity: function emergencyExit() view returns(bool)
func (_BaseStrategy *BaseStrategyCallerSession) EmergencyExit() (bool, error) {
	return _BaseStrategy.Contract.EmergencyExit(&_BaseStrategy.CallOpts)
}

// EstimatedTotalAssets is a free data retrieval call binding the contract method 0xefbb5cb0.
//
// Solidity: function estimatedTotalAssets() view returns(uint256)
func (_BaseStrategy *BaseStrategyCaller) EstimatedTotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BaseStrategy.contract.Call(opts, &out, "estimatedTotalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimatedTotalAssets is a free data retrieval call binding the contract method 0xefbb5cb0.
//
// Solidity: function estimatedTotalAssets() view returns(uint256)
func (_BaseStrategy *BaseStrategySession) EstimatedTotalAssets() (*big.Int, error) {
	return _BaseStrategy.Contract.EstimatedTotalAssets(&_BaseStrategy.CallOpts)
}

// EstimatedTotalAssets is a free data retrieval call binding the contract method 0xefbb5cb0.
//
// Solidity: function estimatedTotalAssets() view returns(uint256)
func (_BaseStrategy *BaseStrategyCallerSession) EstimatedTotalAssets() (*big.Int, error) {
	return _BaseStrategy.Contract.EstimatedTotalAssets(&_BaseStrategy.CallOpts)
}

// HarvestTrigger is a free data retrieval call binding the contract method 0xed882c2b.
//
// Solidity: function harvestTrigger(uint256 callCost) view returns(bool)
func (_BaseStrategy *BaseStrategyCaller) HarvestTrigger(opts *bind.CallOpts, callCost *big.Int) (bool, error) {
	var out []interface{}
	err := _BaseStrategy.contract.Call(opts, &out, "harvestTrigger", callCost)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HarvestTrigger is a free data retrieval call binding the contract method 0xed882c2b.
//
// Solidity: function harvestTrigger(uint256 callCost) view returns(bool)
func (_BaseStrategy *BaseStrategySession) HarvestTrigger(callCost *big.Int) (bool, error) {
	return _BaseStrategy.Contract.HarvestTrigger(&_BaseStrategy.CallOpts, callCost)
}

// HarvestTrigger is a free data retrieval call binding the contract method 0xed882c2b.
//
// Solidity: function harvestTrigger(uint256 callCost) view returns(bool)
func (_BaseStrategy *BaseStrategyCallerSession) HarvestTrigger(callCost *big.Int) (bool, error) {
	return _BaseStrategy.Contract.HarvestTrigger(&_BaseStrategy.CallOpts, callCost)
}

// HealthCheck is a free data retrieval call binding the contract method 0xb252720b.
//
// Solidity: function healthCheck() view returns(address)
func (_BaseStrategy *BaseStrategyCaller) HealthCheck(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseStrategy.contract.Call(opts, &out, "healthCheck")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HealthCheck is a free data retrieval call binding the contract method 0xb252720b.
//
// Solidity: function healthCheck() view returns(address)
func (_BaseStrategy *BaseStrategySession) HealthCheck() (common.Address, error) {
	return _BaseStrategy.Contract.HealthCheck(&_BaseStrategy.CallOpts)
}

// HealthCheck is a free data retrieval call binding the contract method 0xb252720b.
//
// Solidity: function healthCheck() view returns(address)
func (_BaseStrategy *BaseStrategyCallerSession) HealthCheck() (common.Address, error) {
	return _BaseStrategy.Contract.HealthCheck(&_BaseStrategy.CallOpts)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_BaseStrategy *BaseStrategyCaller) IsActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BaseStrategy.contract.Call(opts, &out, "isActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_BaseStrategy *BaseStrategySession) IsActive() (bool, error) {
	return _BaseStrategy.Contract.IsActive(&_BaseStrategy.CallOpts)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_BaseStrategy *BaseStrategyCallerSession) IsActive() (bool, error) {
	return _BaseStrategy.Contract.IsActive(&_BaseStrategy.CallOpts)
}

// Keeper is a free data retrieval call binding the contract method 0xaced1661.
//
// Solidity: function keeper() view returns(address)
func (_BaseStrategy *BaseStrategyCaller) Keeper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseStrategy.contract.Call(opts, &out, "keeper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Keeper is a free data retrieval call binding the contract method 0xaced1661.
//
// Solidity: function keeper() view returns(address)
func (_BaseStrategy *BaseStrategySession) Keeper() (common.Address, error) {
	return _BaseStrategy.Contract.Keeper(&_BaseStrategy.CallOpts)
}

// Keeper is a free data retrieval call binding the contract method 0xaced1661.
//
// Solidity: function keeper() view returns(address)
func (_BaseStrategy *BaseStrategyCallerSession) Keeper() (common.Address, error) {
	return _BaseStrategy.Contract.Keeper(&_BaseStrategy.CallOpts)
}

// MaxReportDelay is a free data retrieval call binding the contract method 0x28b7ccf7.
//
// Solidity: function maxReportDelay() view returns(uint256)
func (_BaseStrategy *BaseStrategyCaller) MaxReportDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BaseStrategy.contract.Call(opts, &out, "maxReportDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxReportDelay is a free data retrieval call binding the contract method 0x28b7ccf7.
//
// Solidity: function maxReportDelay() view returns(uint256)
func (_BaseStrategy *BaseStrategySession) MaxReportDelay() (*big.Int, error) {
	return _BaseStrategy.Contract.MaxReportDelay(&_BaseStrategy.CallOpts)
}

// MaxReportDelay is a free data retrieval call binding the contract method 0x28b7ccf7.
//
// Solidity: function maxReportDelay() view returns(uint256)
func (_BaseStrategy *BaseStrategyCallerSession) MaxReportDelay() (*big.Int, error) {
	return _BaseStrategy.Contract.MaxReportDelay(&_BaseStrategy.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BaseStrategy *BaseStrategyCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BaseStrategy.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BaseStrategy *BaseStrategySession) Name() (string, error) {
	return _BaseStrategy.Contract.Name(&_BaseStrategy.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BaseStrategy *BaseStrategyCallerSession) Name() (string, error) {
	return _BaseStrategy.Contract.Name(&_BaseStrategy.CallOpts)
}

// ProfitFactor is a free data retrieval call binding the contract method 0x8cdfe166.
//
// Solidity: function profitFactor() view returns(uint256)
func (_BaseStrategy *BaseStrategyCaller) ProfitFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BaseStrategy.contract.Call(opts, &out, "profitFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProfitFactor is a free data retrieval call binding the contract method 0x8cdfe166.
//
// Solidity: function profitFactor() view returns(uint256)
func (_BaseStrategy *BaseStrategySession) ProfitFactor() (*big.Int, error) {
	return _BaseStrategy.Contract.ProfitFactor(&_BaseStrategy.CallOpts)
}

// ProfitFactor is a free data retrieval call binding the contract method 0x8cdfe166.
//
// Solidity: function profitFactor() view returns(uint256)
func (_BaseStrategy *BaseStrategyCallerSession) ProfitFactor() (*big.Int, error) {
	return _BaseStrategy.Contract.ProfitFactor(&_BaseStrategy.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x9ec5a894.
//
// Solidity: function rewards() view returns(address)
func (_BaseStrategy *BaseStrategyCaller) Rewards(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseStrategy.contract.Call(opts, &out, "rewards")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0x9ec5a894.
//
// Solidity: function rewards() view returns(address)
func (_BaseStrategy *BaseStrategySession) Rewards() (common.Address, error) {
	return _BaseStrategy.Contract.Rewards(&_BaseStrategy.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x9ec5a894.
//
// Solidity: function rewards() view returns(address)
func (_BaseStrategy *BaseStrategyCallerSession) Rewards() (common.Address, error) {
	return _BaseStrategy.Contract.Rewards(&_BaseStrategy.CallOpts)
}

// Strategist is a free data retrieval call binding the contract method 0x1fe4a686.
//
// Solidity: function strategist() view returns(address)
func (_BaseStrategy *BaseStrategyCaller) Strategist(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseStrategy.contract.Call(opts, &out, "strategist")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Strategist is a free data retrieval call binding the contract method 0x1fe4a686.
//
// Solidity: function strategist() view returns(address)
func (_BaseStrategy *BaseStrategySession) Strategist() (common.Address, error) {
	return _BaseStrategy.Contract.Strategist(&_BaseStrategy.CallOpts)
}

// Strategist is a free data retrieval call binding the contract method 0x1fe4a686.
//
// Solidity: function strategist() view returns(address)
func (_BaseStrategy *BaseStrategyCallerSession) Strategist() (common.Address, error) {
	return _BaseStrategy.Contract.Strategist(&_BaseStrategy.CallOpts)
}

// TendTrigger is a free data retrieval call binding the contract method 0x650d1880.
//
// Solidity: function tendTrigger(uint256 callCost) view returns(bool)
func (_BaseStrategy *BaseStrategyCaller) TendTrigger(opts *bind.CallOpts, callCost *big.Int) (bool, error) {
	var out []interface{}
	err := _BaseStrategy.contract.Call(opts, &out, "tendTrigger", callCost)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TendTrigger is a free data retrieval call binding the contract method 0x650d1880.
//
// Solidity: function tendTrigger(uint256 callCost) view returns(bool)
func (_BaseStrategy *BaseStrategySession) TendTrigger(callCost *big.Int) (bool, error) {
	return _BaseStrategy.Contract.TendTrigger(&_BaseStrategy.CallOpts, callCost)
}

// TendTrigger is a free data retrieval call binding the contract method 0x650d1880.
//
// Solidity: function tendTrigger(uint256 callCost) view returns(bool)
func (_BaseStrategy *BaseStrategyCallerSession) TendTrigger(callCost *big.Int) (bool, error) {
	return _BaseStrategy.Contract.TendTrigger(&_BaseStrategy.CallOpts, callCost)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_BaseStrategy *BaseStrategyCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseStrategy.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_BaseStrategy *BaseStrategySession) Vault() (common.Address, error) {
	return _BaseStrategy.Contract.Vault(&_BaseStrategy.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_BaseStrategy *BaseStrategyCallerSession) Vault() (common.Address, error) {
	return _BaseStrategy.Contract.Vault(&_BaseStrategy.CallOpts)
}

// Want is a free data retrieval call binding the contract method 0x1f1fcd51.
//
// Solidity: function want() view returns(address)
func (_BaseStrategy *BaseStrategyCaller) Want(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseStrategy.contract.Call(opts, &out, "want")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Want is a free data retrieval call binding the contract method 0x1f1fcd51.
//
// Solidity: function want() view returns(address)
func (_BaseStrategy *BaseStrategySession) Want() (common.Address, error) {
	return _BaseStrategy.Contract.Want(&_BaseStrategy.CallOpts)
}

// Want is a free data retrieval call binding the contract method 0x1f1fcd51.
//
// Solidity: function want() view returns(address)
func (_BaseStrategy *BaseStrategyCallerSession) Want() (common.Address, error) {
	return _BaseStrategy.Contract.Want(&_BaseStrategy.CallOpts)
}

// Harvest is a paid mutator transaction binding the contract method 0x4641257d.
//
// Solidity: function harvest() returns()
func (_BaseStrategy *BaseStrategyTransactor) Harvest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseStrategy.contract.Transact(opts, "harvest")
}

// Harvest is a paid mutator transaction binding the contract method 0x4641257d.
//
// Solidity: function harvest() returns()
func (_BaseStrategy *BaseStrategySession) Harvest() (*types.Transaction, error) {
	return _BaseStrategy.Contract.Harvest(&_BaseStrategy.TransactOpts)
}

// Harvest is a paid mutator transaction binding the contract method 0x4641257d.
//
// Solidity: function harvest() returns()
func (_BaseStrategy *BaseStrategyTransactorSession) Harvest() (*types.Transaction, error) {
	return _BaseStrategy.Contract.Harvest(&_BaseStrategy.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _newStrategy) returns()
func (_BaseStrategy *BaseStrategyTransactor) Migrate(opts *bind.TransactOpts, _newStrategy common.Address) (*types.Transaction, error) {
	return _BaseStrategy.contract.Transact(opts, "migrate", _newStrategy)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _newStrategy) returns()
func (_BaseStrategy *BaseStrategySession) Migrate(_newStrategy common.Address) (*types.Transaction, error) {
	return _BaseStrategy.Contract.Migrate(&_BaseStrategy.TransactOpts, _newStrategy)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _newStrategy) returns()
func (_BaseStrategy *BaseStrategyTransactorSession) Migrate(_newStrategy common.Address) (*types.Transaction, error) {
	return _BaseStrategy.Contract.Migrate(&_BaseStrategy.TransactOpts, _newStrategy)
}

// SetDebtThreshold is a paid mutator transaction binding the contract method 0x0f969b87.
//
// Solidity: function setDebtThreshold(uint256 _debtThreshold) returns()
func (_BaseStrategy *BaseStrategyTransactor) SetDebtThreshold(opts *bind.TransactOpts, _debtThreshold *big.Int) (*types.Transaction, error) {
	return _BaseStrategy.contract.Transact(opts, "setDebtThreshold", _debtThreshold)
}

// SetDebtThreshold is a paid mutator transaction binding the contract method 0x0f969b87.
//
// Solidity: function setDebtThreshold(uint256 _debtThreshold) returns()
func (_BaseStrategy *BaseStrategySession) SetDebtThreshold(_debtThreshold *big.Int) (*types.Transaction, error) {
	return _BaseStrategy.Contract.SetDebtThreshold(&_BaseStrategy.TransactOpts, _debtThreshold)
}

// SetDebtThreshold is a paid mutator transaction binding the contract method 0x0f969b87.
//
// Solidity: function setDebtThreshold(uint256 _debtThreshold) returns()
func (_BaseStrategy *BaseStrategyTransactorSession) SetDebtThreshold(_debtThreshold *big.Int) (*types.Transaction, error) {
	return _BaseStrategy.Contract.SetDebtThreshold(&_BaseStrategy.TransactOpts, _debtThreshold)
}

// SetDoHealthCheck is a paid mutator transaction binding the contract method 0xac00ff26.
//
// Solidity: function setDoHealthCheck(bool _doHealthCheck) returns()
func (_BaseStrategy *BaseStrategyTransactor) SetDoHealthCheck(opts *bind.TransactOpts, _doHealthCheck bool) (*types.Transaction, error) {
	return _BaseStrategy.contract.Transact(opts, "setDoHealthCheck", _doHealthCheck)
}

// SetDoHealthCheck is a paid mutator transaction binding the contract method 0xac00ff26.
//
// Solidity: function setDoHealthCheck(bool _doHealthCheck) returns()
func (_BaseStrategy *BaseStrategySession) SetDoHealthCheck(_doHealthCheck bool) (*types.Transaction, error) {
	return _BaseStrategy.Contract.SetDoHealthCheck(&_BaseStrategy.TransactOpts, _doHealthCheck)
}

// SetDoHealthCheck is a paid mutator transaction binding the contract method 0xac00ff26.
//
// Solidity: function setDoHealthCheck(bool _doHealthCheck) returns()
func (_BaseStrategy *BaseStrategyTransactorSession) SetDoHealthCheck(_doHealthCheck bool) (*types.Transaction, error) {
	return _BaseStrategy.Contract.SetDoHealthCheck(&_BaseStrategy.TransactOpts, _doHealthCheck)
}

// SetEmergencyExit is a paid mutator transaction binding the contract method 0xfcf2d0ad.
//
// Solidity: function setEmergencyExit() returns()
func (_BaseStrategy *BaseStrategyTransactor) SetEmergencyExit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseStrategy.contract.Transact(opts, "setEmergencyExit")
}

// SetEmergencyExit is a paid mutator transaction binding the contract method 0xfcf2d0ad.
//
// Solidity: function setEmergencyExit() returns()
func (_BaseStrategy *BaseStrategySession) SetEmergencyExit() (*types.Transaction, error) {
	return _BaseStrategy.Contract.SetEmergencyExit(&_BaseStrategy.TransactOpts)
}

// SetEmergencyExit is a paid mutator transaction binding the contract method 0xfcf2d0ad.
//
// Solidity: function setEmergencyExit() returns()
func (_BaseStrategy *BaseStrategyTransactorSession) SetEmergencyExit() (*types.Transaction, error) {
	return _BaseStrategy.Contract.SetEmergencyExit(&_BaseStrategy.TransactOpts)
}

// SetHealthCheck is a paid mutator transaction binding the contract method 0x11bc8245.
//
// Solidity: function setHealthCheck(address _healthCheck) returns()
func (_BaseStrategy *BaseStrategyTransactor) SetHealthCheck(opts *bind.TransactOpts, _healthCheck common.Address) (*types.Transaction, error) {
	return _BaseStrategy.contract.Transact(opts, "setHealthCheck", _healthCheck)
}

// SetHealthCheck is a paid mutator transaction binding the contract method 0x11bc8245.
//
// Solidity: function setHealthCheck(address _healthCheck) returns()
func (_BaseStrategy *BaseStrategySession) SetHealthCheck(_healthCheck common.Address) (*types.Transaction, error) {
	return _BaseStrategy.Contract.SetHealthCheck(&_BaseStrategy.TransactOpts, _healthCheck)
}

// SetHealthCheck is a paid mutator transaction binding the contract method 0x11bc8245.
//
// Solidity: function setHealthCheck(address _healthCheck) returns()
func (_BaseStrategy *BaseStrategyTransactorSession) SetHealthCheck(_healthCheck common.Address) (*types.Transaction, error) {
	return _BaseStrategy.Contract.SetHealthCheck(&_BaseStrategy.TransactOpts, _healthCheck)
}

// SetKeeper is a paid mutator transaction binding the contract method 0x748747e6.
//
// Solidity: function setKeeper(address _keeper) returns()
func (_BaseStrategy *BaseStrategyTransactor) SetKeeper(opts *bind.TransactOpts, _keeper common.Address) (*types.Transaction, error) {
	return _BaseStrategy.contract.Transact(opts, "setKeeper", _keeper)
}

// SetKeeper is a paid mutator transaction binding the contract method 0x748747e6.
//
// Solidity: function setKeeper(address _keeper) returns()
func (_BaseStrategy *BaseStrategySession) SetKeeper(_keeper common.Address) (*types.Transaction, error) {
	return _BaseStrategy.Contract.SetKeeper(&_BaseStrategy.TransactOpts, _keeper)
}

// SetKeeper is a paid mutator transaction binding the contract method 0x748747e6.
//
// Solidity: function setKeeper(address _keeper) returns()
func (_BaseStrategy *BaseStrategyTransactorSession) SetKeeper(_keeper common.Address) (*types.Transaction, error) {
	return _BaseStrategy.Contract.SetKeeper(&_BaseStrategy.TransactOpts, _keeper)
}

// SetMaxReportDelay is a paid mutator transaction binding the contract method 0xf017c92f.
//
// Solidity: function setMaxReportDelay(uint256 _delay) returns()
func (_BaseStrategy *BaseStrategyTransactor) SetMaxReportDelay(opts *bind.TransactOpts, _delay *big.Int) (*types.Transaction, error) {
	return _BaseStrategy.contract.Transact(opts, "setMaxReportDelay", _delay)
}

// SetMaxReportDelay is a paid mutator transaction binding the contract method 0xf017c92f.
//
// Solidity: function setMaxReportDelay(uint256 _delay) returns()
func (_BaseStrategy *BaseStrategySession) SetMaxReportDelay(_delay *big.Int) (*types.Transaction, error) {
	return _BaseStrategy.Contract.SetMaxReportDelay(&_BaseStrategy.TransactOpts, _delay)
}

// SetMaxReportDelay is a paid mutator transaction binding the contract method 0xf017c92f.
//
// Solidity: function setMaxReportDelay(uint256 _delay) returns()
func (_BaseStrategy *BaseStrategyTransactorSession) SetMaxReportDelay(_delay *big.Int) (*types.Transaction, error) {
	return _BaseStrategy.Contract.SetMaxReportDelay(&_BaseStrategy.TransactOpts, _delay)
}

// SetProfitFactor is a paid mutator transaction binding the contract method 0x91397ab4.
//
// Solidity: function setProfitFactor(uint256 _profitFactor) returns()
func (_BaseStrategy *BaseStrategyTransactor) SetProfitFactor(opts *bind.TransactOpts, _profitFactor *big.Int) (*types.Transaction, error) {
	return _BaseStrategy.contract.Transact(opts, "setProfitFactor", _profitFactor)
}

// SetProfitFactor is a paid mutator transaction binding the contract method 0x91397ab4.
//
// Solidity: function setProfitFactor(uint256 _profitFactor) returns()
func (_BaseStrategy *BaseStrategySession) SetProfitFactor(_profitFactor *big.Int) (*types.Transaction, error) {
	return _BaseStrategy.Contract.SetProfitFactor(&_BaseStrategy.TransactOpts, _profitFactor)
}

// SetProfitFactor is a paid mutator transaction binding the contract method 0x91397ab4.
//
// Solidity: function setProfitFactor(uint256 _profitFactor) returns()
func (_BaseStrategy *BaseStrategyTransactorSession) SetProfitFactor(_profitFactor *big.Int) (*types.Transaction, error) {
	return _BaseStrategy.Contract.SetProfitFactor(&_BaseStrategy.TransactOpts, _profitFactor)
}

// SetRewards is a paid mutator transaction binding the contract method 0xec38a862.
//
// Solidity: function setRewards(address _rewards) returns()
func (_BaseStrategy *BaseStrategyTransactor) SetRewards(opts *bind.TransactOpts, _rewards common.Address) (*types.Transaction, error) {
	return _BaseStrategy.contract.Transact(opts, "setRewards", _rewards)
}

// SetRewards is a paid mutator transaction binding the contract method 0xec38a862.
//
// Solidity: function setRewards(address _rewards) returns()
func (_BaseStrategy *BaseStrategySession) SetRewards(_rewards common.Address) (*types.Transaction, error) {
	return _BaseStrategy.Contract.SetRewards(&_BaseStrategy.TransactOpts, _rewards)
}

// SetRewards is a paid mutator transaction binding the contract method 0xec38a862.
//
// Solidity: function setRewards(address _rewards) returns()
func (_BaseStrategy *BaseStrategyTransactorSession) SetRewards(_rewards common.Address) (*types.Transaction, error) {
	return _BaseStrategy.Contract.SetRewards(&_BaseStrategy.TransactOpts, _rewards)
}

// SetStrategist is a paid mutator transaction binding the contract method 0xc7b9d530.
//
// Solidity: function setStrategist(address _strategist) returns()
func (_BaseStrategy *BaseStrategyTransactor) SetStrategist(opts *bind.TransactOpts, _strategist common.Address) (*types.Transaction, error) {
	return _BaseStrategy.contract.Transact(opts, "setStrategist", _strategist)
}

// SetStrategist is a paid mutator transaction binding the contract method 0xc7b9d530.
//
// Solidity: function setStrategist(address _strategist) returns()
func (_BaseStrategy *BaseStrategySession) SetStrategist(_strategist common.Address) (*types.Transaction, error) {
	return _BaseStrategy.Contract.SetStrategist(&_BaseStrategy.TransactOpts, _strategist)
}

// SetStrategist is a paid mutator transaction binding the contract method 0xc7b9d530.
//
// Solidity: function setStrategist(address _strategist) returns()
func (_BaseStrategy *BaseStrategyTransactorSession) SetStrategist(_strategist common.Address) (*types.Transaction, error) {
	return _BaseStrategy.Contract.SetStrategist(&_BaseStrategy.TransactOpts, _strategist)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address _token) returns()
func (_BaseStrategy *BaseStrategyTransactor) Sweep(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _BaseStrategy.contract.Transact(opts, "sweep", _token)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address _token) returns()
func (_BaseStrategy *BaseStrategySession) Sweep(_token common.Address) (*types.Transaction, error) {
	return _BaseStrategy.Contract.Sweep(&_BaseStrategy.TransactOpts, _token)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address _token) returns()
func (_BaseStrategy *BaseStrategyTransactorSession) Sweep(_token common.Address) (*types.Transaction, error) {
	return _BaseStrategy.Contract.Sweep(&_BaseStrategy.TransactOpts, _token)
}

// Tend is a paid mutator transaction binding the contract method 0x440368a3.
//
// Solidity: function tend() returns()
func (_BaseStrategy *BaseStrategyTransactor) Tend(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseStrategy.contract.Transact(opts, "tend")
}

// Tend is a paid mutator transaction binding the contract method 0x440368a3.
//
// Solidity: function tend() returns()
func (_BaseStrategy *BaseStrategySession) Tend() (*types.Transaction, error) {
	return _BaseStrategy.Contract.Tend(&_BaseStrategy.TransactOpts)
}

// Tend is a paid mutator transaction binding the contract method 0x440368a3.
//
// Solidity: function tend() returns()
func (_BaseStrategy *BaseStrategyTransactorSession) Tend() (*types.Transaction, error) {
	return _BaseStrategy.Contract.Tend(&_BaseStrategy.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amountNeeded) returns(uint256 _loss)
func (_BaseStrategy *BaseStrategyTransactor) Withdraw(opts *bind.TransactOpts, _amountNeeded *big.Int) (*types.Transaction, error) {
	return _BaseStrategy.contract.Transact(opts, "withdraw", _amountNeeded)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amountNeeded) returns(uint256 _loss)
func (_BaseStrategy *BaseStrategySession) Withdraw(_amountNeeded *big.Int) (*types.Transaction, error) {
	return _BaseStrategy.Contract.Withdraw(&_BaseStrategy.TransactOpts, _amountNeeded)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amountNeeded) returns(uint256 _loss)
func (_BaseStrategy *BaseStrategyTransactorSession) Withdraw(_amountNeeded *big.Int) (*types.Transaction, error) {
	return _BaseStrategy.Contract.Withdraw(&_BaseStrategy.TransactOpts, _amountNeeded)
}

// BaseStrategyEmergencyExitEnabledIterator is returned from FilterEmergencyExitEnabled and is used to iterate over the raw logs and unpacked data for EmergencyExitEnabled events raised by the BaseStrategy contract.
type BaseStrategyEmergencyExitEnabledIterator struct {
	Event *BaseStrategyEmergencyExitEnabled // Event containing the contract specifics and raw log

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
func (it *BaseStrategyEmergencyExitEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseStrategyEmergencyExitEnabled)
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
		it.Event = new(BaseStrategyEmergencyExitEnabled)
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
func (it *BaseStrategyEmergencyExitEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseStrategyEmergencyExitEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseStrategyEmergencyExitEnabled represents a EmergencyExitEnabled event raised by the BaseStrategy contract.
type BaseStrategyEmergencyExitEnabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyExitEnabled is a free log retrieval operation binding the contract event 0x97e963041e952738788b9d4871d854d282065b8f90a464928d6528f2e9a4fd0b.
//
// Solidity: event EmergencyExitEnabled()
func (_BaseStrategy *BaseStrategyFilterer) FilterEmergencyExitEnabled(opts *bind.FilterOpts) (*BaseStrategyEmergencyExitEnabledIterator, error) {

	logs, sub, err := _BaseStrategy.contract.FilterLogs(opts, "EmergencyExitEnabled")
	if err != nil {
		return nil, err
	}
	return &BaseStrategyEmergencyExitEnabledIterator{contract: _BaseStrategy.contract, event: "EmergencyExitEnabled", logs: logs, sub: sub}, nil
}

// WatchEmergencyExitEnabled is a free log subscription operation binding the contract event 0x97e963041e952738788b9d4871d854d282065b8f90a464928d6528f2e9a4fd0b.
//
// Solidity: event EmergencyExitEnabled()
func (_BaseStrategy *BaseStrategyFilterer) WatchEmergencyExitEnabled(opts *bind.WatchOpts, sink chan<- *BaseStrategyEmergencyExitEnabled) (event.Subscription, error) {

	logs, sub, err := _BaseStrategy.contract.WatchLogs(opts, "EmergencyExitEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseStrategyEmergencyExitEnabled)
				if err := _BaseStrategy.contract.UnpackLog(event, "EmergencyExitEnabled", log); err != nil {
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

// ParseEmergencyExitEnabled is a log parse operation binding the contract event 0x97e963041e952738788b9d4871d854d282065b8f90a464928d6528f2e9a4fd0b.
//
// Solidity: event EmergencyExitEnabled()
func (_BaseStrategy *BaseStrategyFilterer) ParseEmergencyExitEnabled(log types.Log) (*BaseStrategyEmergencyExitEnabled, error) {
	event := new(BaseStrategyEmergencyExitEnabled)
	if err := _BaseStrategy.contract.UnpackLog(event, "EmergencyExitEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseStrategyHarvestedIterator is returned from FilterHarvested and is used to iterate over the raw logs and unpacked data for Harvested events raised by the BaseStrategy contract.
type BaseStrategyHarvestedIterator struct {
	Event *BaseStrategyHarvested // Event containing the contract specifics and raw log

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
func (it *BaseStrategyHarvestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseStrategyHarvested)
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
		it.Event = new(BaseStrategyHarvested)
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
func (it *BaseStrategyHarvestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseStrategyHarvestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseStrategyHarvested represents a Harvested event raised by the BaseStrategy contract.
type BaseStrategyHarvested struct {
	Profit          *big.Int
	Loss            *big.Int
	DebtPayment     *big.Int
	DebtOutstanding *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterHarvested is a free log retrieval operation binding the contract event 0x4c0f499ffe6befa0ca7c826b0916cf87bea98de658013e76938489368d60d509.
//
// Solidity: event Harvested(uint256 profit, uint256 loss, uint256 debtPayment, uint256 debtOutstanding)
func (_BaseStrategy *BaseStrategyFilterer) FilterHarvested(opts *bind.FilterOpts) (*BaseStrategyHarvestedIterator, error) {

	logs, sub, err := _BaseStrategy.contract.FilterLogs(opts, "Harvested")
	if err != nil {
		return nil, err
	}
	return &BaseStrategyHarvestedIterator{contract: _BaseStrategy.contract, event: "Harvested", logs: logs, sub: sub}, nil
}

// WatchHarvested is a free log subscription operation binding the contract event 0x4c0f499ffe6befa0ca7c826b0916cf87bea98de658013e76938489368d60d509.
//
// Solidity: event Harvested(uint256 profit, uint256 loss, uint256 debtPayment, uint256 debtOutstanding)
func (_BaseStrategy *BaseStrategyFilterer) WatchHarvested(opts *bind.WatchOpts, sink chan<- *BaseStrategyHarvested) (event.Subscription, error) {

	logs, sub, err := _BaseStrategy.contract.WatchLogs(opts, "Harvested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseStrategyHarvested)
				if err := _BaseStrategy.contract.UnpackLog(event, "Harvested", log); err != nil {
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

// ParseHarvested is a log parse operation binding the contract event 0x4c0f499ffe6befa0ca7c826b0916cf87bea98de658013e76938489368d60d509.
//
// Solidity: event Harvested(uint256 profit, uint256 loss, uint256 debtPayment, uint256 debtOutstanding)
func (_BaseStrategy *BaseStrategyFilterer) ParseHarvested(log types.Log) (*BaseStrategyHarvested, error) {
	event := new(BaseStrategyHarvested)
	if err := _BaseStrategy.contract.UnpackLog(event, "Harvested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseStrategyUpdatedDebtThresholdIterator is returned from FilterUpdatedDebtThreshold and is used to iterate over the raw logs and unpacked data for UpdatedDebtThreshold events raised by the BaseStrategy contract.
type BaseStrategyUpdatedDebtThresholdIterator struct {
	Event *BaseStrategyUpdatedDebtThreshold // Event containing the contract specifics and raw log

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
func (it *BaseStrategyUpdatedDebtThresholdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseStrategyUpdatedDebtThreshold)
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
		it.Event = new(BaseStrategyUpdatedDebtThreshold)
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
func (it *BaseStrategyUpdatedDebtThresholdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseStrategyUpdatedDebtThresholdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseStrategyUpdatedDebtThreshold represents a UpdatedDebtThreshold event raised by the BaseStrategy contract.
type BaseStrategyUpdatedDebtThreshold struct {
	DebtThreshold *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedDebtThreshold is a free log retrieval operation binding the contract event 0xa68ba126373d04c004c5748c300c9fca12bd444b3d4332e261f3bd2bac4a8600.
//
// Solidity: event UpdatedDebtThreshold(uint256 debtThreshold)
func (_BaseStrategy *BaseStrategyFilterer) FilterUpdatedDebtThreshold(opts *bind.FilterOpts) (*BaseStrategyUpdatedDebtThresholdIterator, error) {

	logs, sub, err := _BaseStrategy.contract.FilterLogs(opts, "UpdatedDebtThreshold")
	if err != nil {
		return nil, err
	}
	return &BaseStrategyUpdatedDebtThresholdIterator{contract: _BaseStrategy.contract, event: "UpdatedDebtThreshold", logs: logs, sub: sub}, nil
}

// WatchUpdatedDebtThreshold is a free log subscription operation binding the contract event 0xa68ba126373d04c004c5748c300c9fca12bd444b3d4332e261f3bd2bac4a8600.
//
// Solidity: event UpdatedDebtThreshold(uint256 debtThreshold)
func (_BaseStrategy *BaseStrategyFilterer) WatchUpdatedDebtThreshold(opts *bind.WatchOpts, sink chan<- *BaseStrategyUpdatedDebtThreshold) (event.Subscription, error) {

	logs, sub, err := _BaseStrategy.contract.WatchLogs(opts, "UpdatedDebtThreshold")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseStrategyUpdatedDebtThreshold)
				if err := _BaseStrategy.contract.UnpackLog(event, "UpdatedDebtThreshold", log); err != nil {
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

// ParseUpdatedDebtThreshold is a log parse operation binding the contract event 0xa68ba126373d04c004c5748c300c9fca12bd444b3d4332e261f3bd2bac4a8600.
//
// Solidity: event UpdatedDebtThreshold(uint256 debtThreshold)
func (_BaseStrategy *BaseStrategyFilterer) ParseUpdatedDebtThreshold(log types.Log) (*BaseStrategyUpdatedDebtThreshold, error) {
	event := new(BaseStrategyUpdatedDebtThreshold)
	if err := _BaseStrategy.contract.UnpackLog(event, "UpdatedDebtThreshold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseStrategyUpdatedKeeperIterator is returned from FilterUpdatedKeeper and is used to iterate over the raw logs and unpacked data for UpdatedKeeper events raised by the BaseStrategy contract.
type BaseStrategyUpdatedKeeperIterator struct {
	Event *BaseStrategyUpdatedKeeper // Event containing the contract specifics and raw log

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
func (it *BaseStrategyUpdatedKeeperIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseStrategyUpdatedKeeper)
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
		it.Event = new(BaseStrategyUpdatedKeeper)
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
func (it *BaseStrategyUpdatedKeeperIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseStrategyUpdatedKeeperIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseStrategyUpdatedKeeper represents a UpdatedKeeper event raised by the BaseStrategy contract.
type BaseStrategyUpdatedKeeper struct {
	NewKeeper common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdatedKeeper is a free log retrieval operation binding the contract event 0x2f202ddb4a2e345f6323ed90f8fc8559d770a7abbbeee84dde8aca3351fe7154.
//
// Solidity: event UpdatedKeeper(address newKeeper)
func (_BaseStrategy *BaseStrategyFilterer) FilterUpdatedKeeper(opts *bind.FilterOpts) (*BaseStrategyUpdatedKeeperIterator, error) {

	logs, sub, err := _BaseStrategy.contract.FilterLogs(opts, "UpdatedKeeper")
	if err != nil {
		return nil, err
	}
	return &BaseStrategyUpdatedKeeperIterator{contract: _BaseStrategy.contract, event: "UpdatedKeeper", logs: logs, sub: sub}, nil
}

// WatchUpdatedKeeper is a free log subscription operation binding the contract event 0x2f202ddb4a2e345f6323ed90f8fc8559d770a7abbbeee84dde8aca3351fe7154.
//
// Solidity: event UpdatedKeeper(address newKeeper)
func (_BaseStrategy *BaseStrategyFilterer) WatchUpdatedKeeper(opts *bind.WatchOpts, sink chan<- *BaseStrategyUpdatedKeeper) (event.Subscription, error) {

	logs, sub, err := _BaseStrategy.contract.WatchLogs(opts, "UpdatedKeeper")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseStrategyUpdatedKeeper)
				if err := _BaseStrategy.contract.UnpackLog(event, "UpdatedKeeper", log); err != nil {
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

// ParseUpdatedKeeper is a log parse operation binding the contract event 0x2f202ddb4a2e345f6323ed90f8fc8559d770a7abbbeee84dde8aca3351fe7154.
//
// Solidity: event UpdatedKeeper(address newKeeper)
func (_BaseStrategy *BaseStrategyFilterer) ParseUpdatedKeeper(log types.Log) (*BaseStrategyUpdatedKeeper, error) {
	event := new(BaseStrategyUpdatedKeeper)
	if err := _BaseStrategy.contract.UnpackLog(event, "UpdatedKeeper", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseStrategyUpdatedProfitFactorIterator is returned from FilterUpdatedProfitFactor and is used to iterate over the raw logs and unpacked data for UpdatedProfitFactor events raised by the BaseStrategy contract.
type BaseStrategyUpdatedProfitFactorIterator struct {
	Event *BaseStrategyUpdatedProfitFactor // Event containing the contract specifics and raw log

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
func (it *BaseStrategyUpdatedProfitFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseStrategyUpdatedProfitFactor)
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
		it.Event = new(BaseStrategyUpdatedProfitFactor)
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
func (it *BaseStrategyUpdatedProfitFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseStrategyUpdatedProfitFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseStrategyUpdatedProfitFactor represents a UpdatedProfitFactor event raised by the BaseStrategy contract.
type BaseStrategyUpdatedProfitFactor struct {
	ProfitFactor *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdatedProfitFactor is a free log retrieval operation binding the contract event 0xd94596337df4c2f0f44d30a7fc5db1c7bb60d9aca4185ed77c6fd96eb45ec298.
//
// Solidity: event UpdatedProfitFactor(uint256 profitFactor)
func (_BaseStrategy *BaseStrategyFilterer) FilterUpdatedProfitFactor(opts *bind.FilterOpts) (*BaseStrategyUpdatedProfitFactorIterator, error) {

	logs, sub, err := _BaseStrategy.contract.FilterLogs(opts, "UpdatedProfitFactor")
	if err != nil {
		return nil, err
	}
	return &BaseStrategyUpdatedProfitFactorIterator{contract: _BaseStrategy.contract, event: "UpdatedProfitFactor", logs: logs, sub: sub}, nil
}

// WatchUpdatedProfitFactor is a free log subscription operation binding the contract event 0xd94596337df4c2f0f44d30a7fc5db1c7bb60d9aca4185ed77c6fd96eb45ec298.
//
// Solidity: event UpdatedProfitFactor(uint256 profitFactor)
func (_BaseStrategy *BaseStrategyFilterer) WatchUpdatedProfitFactor(opts *bind.WatchOpts, sink chan<- *BaseStrategyUpdatedProfitFactor) (event.Subscription, error) {

	logs, sub, err := _BaseStrategy.contract.WatchLogs(opts, "UpdatedProfitFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseStrategyUpdatedProfitFactor)
				if err := _BaseStrategy.contract.UnpackLog(event, "UpdatedProfitFactor", log); err != nil {
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

// ParseUpdatedProfitFactor is a log parse operation binding the contract event 0xd94596337df4c2f0f44d30a7fc5db1c7bb60d9aca4185ed77c6fd96eb45ec298.
//
// Solidity: event UpdatedProfitFactor(uint256 profitFactor)
func (_BaseStrategy *BaseStrategyFilterer) ParseUpdatedProfitFactor(log types.Log) (*BaseStrategyUpdatedProfitFactor, error) {
	event := new(BaseStrategyUpdatedProfitFactor)
	if err := _BaseStrategy.contract.UnpackLog(event, "UpdatedProfitFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseStrategyUpdatedReportDelayIterator is returned from FilterUpdatedReportDelay and is used to iterate over the raw logs and unpacked data for UpdatedReportDelay events raised by the BaseStrategy contract.
type BaseStrategyUpdatedReportDelayIterator struct {
	Event *BaseStrategyUpdatedReportDelay // Event containing the contract specifics and raw log

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
func (it *BaseStrategyUpdatedReportDelayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseStrategyUpdatedReportDelay)
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
		it.Event = new(BaseStrategyUpdatedReportDelay)
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
func (it *BaseStrategyUpdatedReportDelayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseStrategyUpdatedReportDelayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseStrategyUpdatedReportDelay represents a UpdatedReportDelay event raised by the BaseStrategy contract.
type BaseStrategyUpdatedReportDelay struct {
	Delay *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdatedReportDelay is a free log retrieval operation binding the contract event 0x4aaf232568bff365c53cad69bdb6e83014e79df80216ceba8ee01769723dfd68.
//
// Solidity: event UpdatedReportDelay(uint256 delay)
func (_BaseStrategy *BaseStrategyFilterer) FilterUpdatedReportDelay(opts *bind.FilterOpts) (*BaseStrategyUpdatedReportDelayIterator, error) {

	logs, sub, err := _BaseStrategy.contract.FilterLogs(opts, "UpdatedReportDelay")
	if err != nil {
		return nil, err
	}
	return &BaseStrategyUpdatedReportDelayIterator{contract: _BaseStrategy.contract, event: "UpdatedReportDelay", logs: logs, sub: sub}, nil
}

// WatchUpdatedReportDelay is a free log subscription operation binding the contract event 0x4aaf232568bff365c53cad69bdb6e83014e79df80216ceba8ee01769723dfd68.
//
// Solidity: event UpdatedReportDelay(uint256 delay)
func (_BaseStrategy *BaseStrategyFilterer) WatchUpdatedReportDelay(opts *bind.WatchOpts, sink chan<- *BaseStrategyUpdatedReportDelay) (event.Subscription, error) {

	logs, sub, err := _BaseStrategy.contract.WatchLogs(opts, "UpdatedReportDelay")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseStrategyUpdatedReportDelay)
				if err := _BaseStrategy.contract.UnpackLog(event, "UpdatedReportDelay", log); err != nil {
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

// ParseUpdatedReportDelay is a log parse operation binding the contract event 0x4aaf232568bff365c53cad69bdb6e83014e79df80216ceba8ee01769723dfd68.
//
// Solidity: event UpdatedReportDelay(uint256 delay)
func (_BaseStrategy *BaseStrategyFilterer) ParseUpdatedReportDelay(log types.Log) (*BaseStrategyUpdatedReportDelay, error) {
	event := new(BaseStrategyUpdatedReportDelay)
	if err := _BaseStrategy.contract.UnpackLog(event, "UpdatedReportDelay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseStrategyUpdatedRewardsIterator is returned from FilterUpdatedRewards and is used to iterate over the raw logs and unpacked data for UpdatedRewards events raised by the BaseStrategy contract.
type BaseStrategyUpdatedRewardsIterator struct {
	Event *BaseStrategyUpdatedRewards // Event containing the contract specifics and raw log

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
func (it *BaseStrategyUpdatedRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseStrategyUpdatedRewards)
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
		it.Event = new(BaseStrategyUpdatedRewards)
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
func (it *BaseStrategyUpdatedRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseStrategyUpdatedRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseStrategyUpdatedRewards represents a UpdatedRewards event raised by the BaseStrategy contract.
type BaseStrategyUpdatedRewards struct {
	Rewards common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdatedRewards is a free log retrieval operation binding the contract event 0xafbb66abf8f3b719799940473a4052a3717cdd8e40fb6c8a3faadab316b1a069.
//
// Solidity: event UpdatedRewards(address rewards)
func (_BaseStrategy *BaseStrategyFilterer) FilterUpdatedRewards(opts *bind.FilterOpts) (*BaseStrategyUpdatedRewardsIterator, error) {

	logs, sub, err := _BaseStrategy.contract.FilterLogs(opts, "UpdatedRewards")
	if err != nil {
		return nil, err
	}
	return &BaseStrategyUpdatedRewardsIterator{contract: _BaseStrategy.contract, event: "UpdatedRewards", logs: logs, sub: sub}, nil
}

// WatchUpdatedRewards is a free log subscription operation binding the contract event 0xafbb66abf8f3b719799940473a4052a3717cdd8e40fb6c8a3faadab316b1a069.
//
// Solidity: event UpdatedRewards(address rewards)
func (_BaseStrategy *BaseStrategyFilterer) WatchUpdatedRewards(opts *bind.WatchOpts, sink chan<- *BaseStrategyUpdatedRewards) (event.Subscription, error) {

	logs, sub, err := _BaseStrategy.contract.WatchLogs(opts, "UpdatedRewards")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseStrategyUpdatedRewards)
				if err := _BaseStrategy.contract.UnpackLog(event, "UpdatedRewards", log); err != nil {
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

// ParseUpdatedRewards is a log parse operation binding the contract event 0xafbb66abf8f3b719799940473a4052a3717cdd8e40fb6c8a3faadab316b1a069.
//
// Solidity: event UpdatedRewards(address rewards)
func (_BaseStrategy *BaseStrategyFilterer) ParseUpdatedRewards(log types.Log) (*BaseStrategyUpdatedRewards, error) {
	event := new(BaseStrategyUpdatedRewards)
	if err := _BaseStrategy.contract.UnpackLog(event, "UpdatedRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseStrategyUpdatedStrategistIterator is returned from FilterUpdatedStrategist and is used to iterate over the raw logs and unpacked data for UpdatedStrategist events raised by the BaseStrategy contract.
type BaseStrategyUpdatedStrategistIterator struct {
	Event *BaseStrategyUpdatedStrategist // Event containing the contract specifics and raw log

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
func (it *BaseStrategyUpdatedStrategistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseStrategyUpdatedStrategist)
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
		it.Event = new(BaseStrategyUpdatedStrategist)
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
func (it *BaseStrategyUpdatedStrategistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseStrategyUpdatedStrategistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseStrategyUpdatedStrategist represents a UpdatedStrategist event raised by the BaseStrategy contract.
type BaseStrategyUpdatedStrategist struct {
	NewStrategist common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedStrategist is a free log retrieval operation binding the contract event 0x352ececae6d7d1e6d26bcf2c549dfd55be1637e9b22dc0cf3b71ddb36097a6b4.
//
// Solidity: event UpdatedStrategist(address newStrategist)
func (_BaseStrategy *BaseStrategyFilterer) FilterUpdatedStrategist(opts *bind.FilterOpts) (*BaseStrategyUpdatedStrategistIterator, error) {

	logs, sub, err := _BaseStrategy.contract.FilterLogs(opts, "UpdatedStrategist")
	if err != nil {
		return nil, err
	}
	return &BaseStrategyUpdatedStrategistIterator{contract: _BaseStrategy.contract, event: "UpdatedStrategist", logs: logs, sub: sub}, nil
}

// WatchUpdatedStrategist is a free log subscription operation binding the contract event 0x352ececae6d7d1e6d26bcf2c549dfd55be1637e9b22dc0cf3b71ddb36097a6b4.
//
// Solidity: event UpdatedStrategist(address newStrategist)
func (_BaseStrategy *BaseStrategyFilterer) WatchUpdatedStrategist(opts *bind.WatchOpts, sink chan<- *BaseStrategyUpdatedStrategist) (event.Subscription, error) {

	logs, sub, err := _BaseStrategy.contract.WatchLogs(opts, "UpdatedStrategist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseStrategyUpdatedStrategist)
				if err := _BaseStrategy.contract.UnpackLog(event, "UpdatedStrategist", log); err != nil {
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

// ParseUpdatedStrategist is a log parse operation binding the contract event 0x352ececae6d7d1e6d26bcf2c549dfd55be1637e9b22dc0cf3b71ddb36097a6b4.
//
// Solidity: event UpdatedStrategist(address newStrategist)
func (_BaseStrategy *BaseStrategyFilterer) ParseUpdatedStrategist(log types.Log) (*BaseStrategyUpdatedStrategist, error) {
	event := new(BaseStrategyUpdatedStrategist)
	if err := _BaseStrategy.contract.UnpackLog(event, "UpdatedStrategist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HealthCheckMetaData contains all meta data concerning the HealthCheck contract.
var HealthCheckMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"loss\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debtPayment\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debtOutstanding\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDebt\",\"type\":\"uint256\"}],\"name\":\"check\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"c70fa00b": "check(uint256,uint256,uint256,uint256,uint256)",
	},
}

// HealthCheckABI is the input ABI used to generate the binding from.
// Deprecated: Use HealthCheckMetaData.ABI instead.
var HealthCheckABI = HealthCheckMetaData.ABI

// Deprecated: Use HealthCheckMetaData.Sigs instead.
// HealthCheckFuncSigs maps the 4-byte function signature to its string representation.
var HealthCheckFuncSigs = HealthCheckMetaData.Sigs

// HealthCheck is an auto generated Go binding around an Ethereum contract.
type HealthCheck struct {
	HealthCheckCaller     // Read-only binding to the contract
	HealthCheckTransactor // Write-only binding to the contract
	HealthCheckFilterer   // Log filterer for contract events
}

// HealthCheckCaller is an auto generated read-only Go binding around an Ethereum contract.
type HealthCheckCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HealthCheckTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HealthCheckTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HealthCheckFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HealthCheckFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HealthCheckSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HealthCheckSession struct {
	Contract     *HealthCheck      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HealthCheckCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HealthCheckCallerSession struct {
	Contract *HealthCheckCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// HealthCheckTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HealthCheckTransactorSession struct {
	Contract     *HealthCheckTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// HealthCheckRaw is an auto generated low-level Go binding around an Ethereum contract.
type HealthCheckRaw struct {
	Contract *HealthCheck // Generic contract binding to access the raw methods on
}

// HealthCheckCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HealthCheckCallerRaw struct {
	Contract *HealthCheckCaller // Generic read-only contract binding to access the raw methods on
}

// HealthCheckTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HealthCheckTransactorRaw struct {
	Contract *HealthCheckTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHealthCheck creates a new instance of HealthCheck, bound to a specific deployed contract.
func NewHealthCheck(address common.Address, backend bind.ContractBackend) (*HealthCheck, error) {
	contract, err := bindHealthCheck(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HealthCheck{HealthCheckCaller: HealthCheckCaller{contract: contract}, HealthCheckTransactor: HealthCheckTransactor{contract: contract}, HealthCheckFilterer: HealthCheckFilterer{contract: contract}}, nil
}

// NewHealthCheckCaller creates a new read-only instance of HealthCheck, bound to a specific deployed contract.
func NewHealthCheckCaller(address common.Address, caller bind.ContractCaller) (*HealthCheckCaller, error) {
	contract, err := bindHealthCheck(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HealthCheckCaller{contract: contract}, nil
}

// NewHealthCheckTransactor creates a new write-only instance of HealthCheck, bound to a specific deployed contract.
func NewHealthCheckTransactor(address common.Address, transactor bind.ContractTransactor) (*HealthCheckTransactor, error) {
	contract, err := bindHealthCheck(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HealthCheckTransactor{contract: contract}, nil
}

// NewHealthCheckFilterer creates a new log filterer instance of HealthCheck, bound to a specific deployed contract.
func NewHealthCheckFilterer(address common.Address, filterer bind.ContractFilterer) (*HealthCheckFilterer, error) {
	contract, err := bindHealthCheck(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HealthCheckFilterer{contract: contract}, nil
}

// bindHealthCheck binds a generic wrapper to an already deployed contract.
func bindHealthCheck(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HealthCheckABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HealthCheck *HealthCheckRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HealthCheck.Contract.HealthCheckCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HealthCheck *HealthCheckRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HealthCheck.Contract.HealthCheckTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HealthCheck *HealthCheckRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HealthCheck.Contract.HealthCheckTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HealthCheck *HealthCheckCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HealthCheck.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HealthCheck *HealthCheckTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HealthCheck.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HealthCheck *HealthCheckTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HealthCheck.Contract.contract.Transact(opts, method, params...)
}

// Check is a free data retrieval call binding the contract method 0xc70fa00b.
//
// Solidity: function check(uint256 profit, uint256 loss, uint256 debtPayment, uint256 debtOutstanding, uint256 totalDebt) view returns(bool)
func (_HealthCheck *HealthCheckCaller) Check(opts *bind.CallOpts, profit *big.Int, loss *big.Int, debtPayment *big.Int, debtOutstanding *big.Int, totalDebt *big.Int) (bool, error) {
	var out []interface{}
	err := _HealthCheck.contract.Call(opts, &out, "check", profit, loss, debtPayment, debtOutstanding, totalDebt)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Check is a free data retrieval call binding the contract method 0xc70fa00b.
//
// Solidity: function check(uint256 profit, uint256 loss, uint256 debtPayment, uint256 debtOutstanding, uint256 totalDebt) view returns(bool)
func (_HealthCheck *HealthCheckSession) Check(profit *big.Int, loss *big.Int, debtPayment *big.Int, debtOutstanding *big.Int, totalDebt *big.Int) (bool, error) {
	return _HealthCheck.Contract.Check(&_HealthCheck.CallOpts, profit, loss, debtPayment, debtOutstanding, totalDebt)
}

// Check is a free data retrieval call binding the contract method 0xc70fa00b.
//
// Solidity: function check(uint256 profit, uint256 loss, uint256 debtPayment, uint256 debtOutstanding, uint256 totalDebt) view returns(bool)
func (_HealthCheck *HealthCheckCallerSession) Check(profit *big.Int, loss *big.Int, debtPayment *big.Int, debtOutstanding *big.Int, totalDebt *big.Int) (bool, error) {
	return _HealthCheck.Contract.Check(&_HealthCheck.CallOpts, profit, loss, debtPayment, debtOutstanding, totalDebt)
}

// IBaseFeeMetaData contains all meta data concerning the IBaseFee contract.
var IBaseFeeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"isCurrentBaseFeeAcceptable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"34a9e75c": "isCurrentBaseFeeAcceptable()",
	},
}

// IBaseFeeABI is the input ABI used to generate the binding from.
// Deprecated: Use IBaseFeeMetaData.ABI instead.
var IBaseFeeABI = IBaseFeeMetaData.ABI

// Deprecated: Use IBaseFeeMetaData.Sigs instead.
// IBaseFeeFuncSigs maps the 4-byte function signature to its string representation.
var IBaseFeeFuncSigs = IBaseFeeMetaData.Sigs

// IBaseFee is an auto generated Go binding around an Ethereum contract.
type IBaseFee struct {
	IBaseFeeCaller     // Read-only binding to the contract
	IBaseFeeTransactor // Write-only binding to the contract
	IBaseFeeFilterer   // Log filterer for contract events
}

// IBaseFeeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBaseFeeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBaseFeeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBaseFeeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBaseFeeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBaseFeeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBaseFeeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBaseFeeSession struct {
	Contract     *IBaseFee         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBaseFeeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBaseFeeCallerSession struct {
	Contract *IBaseFeeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IBaseFeeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBaseFeeTransactorSession struct {
	Contract     *IBaseFeeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IBaseFeeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBaseFeeRaw struct {
	Contract *IBaseFee // Generic contract binding to access the raw methods on
}

// IBaseFeeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBaseFeeCallerRaw struct {
	Contract *IBaseFeeCaller // Generic read-only contract binding to access the raw methods on
}

// IBaseFeeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBaseFeeTransactorRaw struct {
	Contract *IBaseFeeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBaseFee creates a new instance of IBaseFee, bound to a specific deployed contract.
func NewIBaseFee(address common.Address, backend bind.ContractBackend) (*IBaseFee, error) {
	contract, err := bindIBaseFee(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBaseFee{IBaseFeeCaller: IBaseFeeCaller{contract: contract}, IBaseFeeTransactor: IBaseFeeTransactor{contract: contract}, IBaseFeeFilterer: IBaseFeeFilterer{contract: contract}}, nil
}

// NewIBaseFeeCaller creates a new read-only instance of IBaseFee, bound to a specific deployed contract.
func NewIBaseFeeCaller(address common.Address, caller bind.ContractCaller) (*IBaseFeeCaller, error) {
	contract, err := bindIBaseFee(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBaseFeeCaller{contract: contract}, nil
}

// NewIBaseFeeTransactor creates a new write-only instance of IBaseFee, bound to a specific deployed contract.
func NewIBaseFeeTransactor(address common.Address, transactor bind.ContractTransactor) (*IBaseFeeTransactor, error) {
	contract, err := bindIBaseFee(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBaseFeeTransactor{contract: contract}, nil
}

// NewIBaseFeeFilterer creates a new log filterer instance of IBaseFee, bound to a specific deployed contract.
func NewIBaseFeeFilterer(address common.Address, filterer bind.ContractFilterer) (*IBaseFeeFilterer, error) {
	contract, err := bindIBaseFee(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBaseFeeFilterer{contract: contract}, nil
}

// bindIBaseFee binds a generic wrapper to an already deployed contract.
func bindIBaseFee(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IBaseFeeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBaseFee *IBaseFeeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBaseFee.Contract.IBaseFeeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBaseFee *IBaseFeeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBaseFee.Contract.IBaseFeeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBaseFee *IBaseFeeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBaseFee.Contract.IBaseFeeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBaseFee *IBaseFeeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBaseFee.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBaseFee *IBaseFeeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBaseFee.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBaseFee *IBaseFeeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBaseFee.Contract.contract.Transact(opts, method, params...)
}

// IsCurrentBaseFeeAcceptable is a free data retrieval call binding the contract method 0x34a9e75c.
//
// Solidity: function isCurrentBaseFeeAcceptable() view returns(bool)
func (_IBaseFee *IBaseFeeCaller) IsCurrentBaseFeeAcceptable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IBaseFee.contract.Call(opts, &out, "isCurrentBaseFeeAcceptable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCurrentBaseFeeAcceptable is a free data retrieval call binding the contract method 0x34a9e75c.
//
// Solidity: function isCurrentBaseFeeAcceptable() view returns(bool)
func (_IBaseFee *IBaseFeeSession) IsCurrentBaseFeeAcceptable() (bool, error) {
	return _IBaseFee.Contract.IsCurrentBaseFeeAcceptable(&_IBaseFee.CallOpts)
}

// IsCurrentBaseFeeAcceptable is a free data retrieval call binding the contract method 0x34a9e75c.
//
// Solidity: function isCurrentBaseFeeAcceptable() view returns(bool)
func (_IBaseFee *IBaseFeeCallerSession) IsCurrentBaseFeeAcceptable() (bool, error) {
	return _IBaseFee.Contract.IsCurrentBaseFeeAcceptable(&_IBaseFee.CallOpts)
}

// IConvexDepositMetaData contains all meta data concerning the IConvexDeposit contract.
var IConvexDepositMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_stake\",\"type\":\"bool\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"43a0d066": "deposit(uint256,uint256,bool)",
		"1526fe27": "poolInfo(uint256)",
		"441a3e70": "withdraw(uint256,uint256)",
	},
}

// IConvexDepositABI is the input ABI used to generate the binding from.
// Deprecated: Use IConvexDepositMetaData.ABI instead.
var IConvexDepositABI = IConvexDepositMetaData.ABI

// Deprecated: Use IConvexDepositMetaData.Sigs instead.
// IConvexDepositFuncSigs maps the 4-byte function signature to its string representation.
var IConvexDepositFuncSigs = IConvexDepositMetaData.Sigs

// IConvexDeposit is an auto generated Go binding around an Ethereum contract.
type IConvexDeposit struct {
	IConvexDepositCaller     // Read-only binding to the contract
	IConvexDepositTransactor // Write-only binding to the contract
	IConvexDepositFilterer   // Log filterer for contract events
}

// IConvexDepositCaller is an auto generated read-only Go binding around an Ethereum contract.
type IConvexDepositCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IConvexDepositTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IConvexDepositTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IConvexDepositFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IConvexDepositFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IConvexDepositSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IConvexDepositSession struct {
	Contract     *IConvexDeposit   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IConvexDepositCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IConvexDepositCallerSession struct {
	Contract *IConvexDepositCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IConvexDepositTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IConvexDepositTransactorSession struct {
	Contract     *IConvexDepositTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IConvexDepositRaw is an auto generated low-level Go binding around an Ethereum contract.
type IConvexDepositRaw struct {
	Contract *IConvexDeposit // Generic contract binding to access the raw methods on
}

// IConvexDepositCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IConvexDepositCallerRaw struct {
	Contract *IConvexDepositCaller // Generic read-only contract binding to access the raw methods on
}

// IConvexDepositTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IConvexDepositTransactorRaw struct {
	Contract *IConvexDepositTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIConvexDeposit creates a new instance of IConvexDeposit, bound to a specific deployed contract.
func NewIConvexDeposit(address common.Address, backend bind.ContractBackend) (*IConvexDeposit, error) {
	contract, err := bindIConvexDeposit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IConvexDeposit{IConvexDepositCaller: IConvexDepositCaller{contract: contract}, IConvexDepositTransactor: IConvexDepositTransactor{contract: contract}, IConvexDepositFilterer: IConvexDepositFilterer{contract: contract}}, nil
}

// NewIConvexDepositCaller creates a new read-only instance of IConvexDeposit, bound to a specific deployed contract.
func NewIConvexDepositCaller(address common.Address, caller bind.ContractCaller) (*IConvexDepositCaller, error) {
	contract, err := bindIConvexDeposit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IConvexDepositCaller{contract: contract}, nil
}

// NewIConvexDepositTransactor creates a new write-only instance of IConvexDeposit, bound to a specific deployed contract.
func NewIConvexDepositTransactor(address common.Address, transactor bind.ContractTransactor) (*IConvexDepositTransactor, error) {
	contract, err := bindIConvexDeposit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IConvexDepositTransactor{contract: contract}, nil
}

// NewIConvexDepositFilterer creates a new log filterer instance of IConvexDeposit, bound to a specific deployed contract.
func NewIConvexDepositFilterer(address common.Address, filterer bind.ContractFilterer) (*IConvexDepositFilterer, error) {
	contract, err := bindIConvexDeposit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IConvexDepositFilterer{contract: contract}, nil
}

// bindIConvexDeposit binds a generic wrapper to an already deployed contract.
func bindIConvexDeposit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IConvexDepositABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IConvexDeposit *IConvexDepositRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IConvexDeposit.Contract.IConvexDepositCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IConvexDeposit *IConvexDepositRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IConvexDeposit.Contract.IConvexDepositTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IConvexDeposit *IConvexDepositRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IConvexDeposit.Contract.IConvexDepositTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IConvexDeposit *IConvexDepositCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IConvexDeposit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IConvexDeposit *IConvexDepositTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IConvexDeposit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IConvexDeposit *IConvexDepositTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IConvexDeposit.Contract.contract.Transact(opts, method, params...)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address, address, address, address, address, bool)
func (_IConvexDeposit *IConvexDepositCaller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (common.Address, common.Address, common.Address, common.Address, common.Address, bool, error) {
	var out []interface{}
	err := _IConvexDeposit.contract.Call(opts, &out, "poolInfo", arg0)

	if err != nil {
		return *new(common.Address), *new(common.Address), *new(common.Address), *new(common.Address), *new(common.Address), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	out3 := *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	out4 := *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	out5 := *abi.ConvertType(out[5], new(bool)).(*bool)

	return out0, out1, out2, out3, out4, out5, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address, address, address, address, address, bool)
func (_IConvexDeposit *IConvexDepositSession) PoolInfo(arg0 *big.Int) (common.Address, common.Address, common.Address, common.Address, common.Address, bool, error) {
	return _IConvexDeposit.Contract.PoolInfo(&_IConvexDeposit.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address, address, address, address, address, bool)
func (_IConvexDeposit *IConvexDepositCallerSession) PoolInfo(arg0 *big.Int) (common.Address, common.Address, common.Address, common.Address, common.Address, bool, error) {
	return _IConvexDeposit.Contract.PoolInfo(&_IConvexDeposit.CallOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0x43a0d066.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount, bool _stake) returns(bool)
func (_IConvexDeposit *IConvexDepositTransactor) Deposit(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int, _stake bool) (*types.Transaction, error) {
	return _IConvexDeposit.contract.Transact(opts, "deposit", _pid, _amount, _stake)
}

// Deposit is a paid mutator transaction binding the contract method 0x43a0d066.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount, bool _stake) returns(bool)
func (_IConvexDeposit *IConvexDepositSession) Deposit(_pid *big.Int, _amount *big.Int, _stake bool) (*types.Transaction, error) {
	return _IConvexDeposit.Contract.Deposit(&_IConvexDeposit.TransactOpts, _pid, _amount, _stake)
}

// Deposit is a paid mutator transaction binding the contract method 0x43a0d066.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount, bool _stake) returns(bool)
func (_IConvexDeposit *IConvexDepositTransactorSession) Deposit(_pid *big.Int, _amount *big.Int, _stake bool) (*types.Transaction, error) {
	return _IConvexDeposit.Contract.Deposit(&_IConvexDeposit.TransactOpts, _pid, _amount, _stake)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns(bool)
func (_IConvexDeposit *IConvexDepositTransactor) Withdraw(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _IConvexDeposit.contract.Transact(opts, "withdraw", _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns(bool)
func (_IConvexDeposit *IConvexDepositSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _IConvexDeposit.Contract.Withdraw(&_IConvexDeposit.TransactOpts, _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns(bool)
func (_IConvexDeposit *IConvexDepositTransactorSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _IConvexDeposit.Contract.Withdraw(&_IConvexDeposit.TransactOpts, _pid, _amount)
}

// IConvexRewardsMetaData contains all meta data concerning the IConvexRewards contract.
var IConvexRewardsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"earned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_reward\",\"type\":\"uint256\"}],\"name\":\"extraRewards\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"extraRewardsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_claimExtras\",\"type\":\"bool\"}],\"name\":\"getReward\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodFinish\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_claim\",\"type\":\"bool\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_claim\",\"type\":\"bool\"}],\"name\":\"withdrawAndUnwrap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"70a08231": "balanceOf(address)",
		"008cc262": "earned(address)",
		"40c35446": "extraRewards(uint256)",
		"d55a23f4": "extraRewardsLength()",
		"7050ccd9": "getReward(address,bool)",
		"ebe2b12b": "periodFinish()",
		"f7c618c1": "rewardToken()",
		"a694fc3a": "stake(uint256)",
		"38d07436": "withdraw(uint256,bool)",
		"c32e7202": "withdrawAndUnwrap(uint256,bool)",
	},
}

// IConvexRewardsABI is the input ABI used to generate the binding from.
// Deprecated: Use IConvexRewardsMetaData.ABI instead.
var IConvexRewardsABI = IConvexRewardsMetaData.ABI

// Deprecated: Use IConvexRewardsMetaData.Sigs instead.
// IConvexRewardsFuncSigs maps the 4-byte function signature to its string representation.
var IConvexRewardsFuncSigs = IConvexRewardsMetaData.Sigs

// IConvexRewards is an auto generated Go binding around an Ethereum contract.
type IConvexRewards struct {
	IConvexRewardsCaller     // Read-only binding to the contract
	IConvexRewardsTransactor // Write-only binding to the contract
	IConvexRewardsFilterer   // Log filterer for contract events
}

// IConvexRewardsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IConvexRewardsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IConvexRewardsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IConvexRewardsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IConvexRewardsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IConvexRewardsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IConvexRewardsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IConvexRewardsSession struct {
	Contract     *IConvexRewards   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IConvexRewardsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IConvexRewardsCallerSession struct {
	Contract *IConvexRewardsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IConvexRewardsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IConvexRewardsTransactorSession struct {
	Contract     *IConvexRewardsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IConvexRewardsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IConvexRewardsRaw struct {
	Contract *IConvexRewards // Generic contract binding to access the raw methods on
}

// IConvexRewardsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IConvexRewardsCallerRaw struct {
	Contract *IConvexRewardsCaller // Generic read-only contract binding to access the raw methods on
}

// IConvexRewardsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IConvexRewardsTransactorRaw struct {
	Contract *IConvexRewardsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIConvexRewards creates a new instance of IConvexRewards, bound to a specific deployed contract.
func NewIConvexRewards(address common.Address, backend bind.ContractBackend) (*IConvexRewards, error) {
	contract, err := bindIConvexRewards(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IConvexRewards{IConvexRewardsCaller: IConvexRewardsCaller{contract: contract}, IConvexRewardsTransactor: IConvexRewardsTransactor{contract: contract}, IConvexRewardsFilterer: IConvexRewardsFilterer{contract: contract}}, nil
}

// NewIConvexRewardsCaller creates a new read-only instance of IConvexRewards, bound to a specific deployed contract.
func NewIConvexRewardsCaller(address common.Address, caller bind.ContractCaller) (*IConvexRewardsCaller, error) {
	contract, err := bindIConvexRewards(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IConvexRewardsCaller{contract: contract}, nil
}

// NewIConvexRewardsTransactor creates a new write-only instance of IConvexRewards, bound to a specific deployed contract.
func NewIConvexRewardsTransactor(address common.Address, transactor bind.ContractTransactor) (*IConvexRewardsTransactor, error) {
	contract, err := bindIConvexRewards(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IConvexRewardsTransactor{contract: contract}, nil
}

// NewIConvexRewardsFilterer creates a new log filterer instance of IConvexRewards, bound to a specific deployed contract.
func NewIConvexRewardsFilterer(address common.Address, filterer bind.ContractFilterer) (*IConvexRewardsFilterer, error) {
	contract, err := bindIConvexRewards(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IConvexRewardsFilterer{contract: contract}, nil
}

// bindIConvexRewards binds a generic wrapper to an already deployed contract.
func bindIConvexRewards(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IConvexRewardsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IConvexRewards *IConvexRewardsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IConvexRewards.Contract.IConvexRewardsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IConvexRewards *IConvexRewardsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IConvexRewards.Contract.IConvexRewardsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IConvexRewards *IConvexRewardsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IConvexRewards.Contract.IConvexRewardsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IConvexRewards *IConvexRewardsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IConvexRewards.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IConvexRewards *IConvexRewardsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IConvexRewards.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IConvexRewards *IConvexRewardsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IConvexRewards.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IConvexRewards *IConvexRewardsCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IConvexRewards.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IConvexRewards *IConvexRewardsSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IConvexRewards.Contract.BalanceOf(&_IConvexRewards.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IConvexRewards *IConvexRewardsCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IConvexRewards.Contract.BalanceOf(&_IConvexRewards.CallOpts, account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_IConvexRewards *IConvexRewardsCaller) Earned(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IConvexRewards.contract.Call(opts, &out, "earned", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_IConvexRewards *IConvexRewardsSession) Earned(account common.Address) (*big.Int, error) {
	return _IConvexRewards.Contract.Earned(&_IConvexRewards.CallOpts, account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_IConvexRewards *IConvexRewardsCallerSession) Earned(account common.Address) (*big.Int, error) {
	return _IConvexRewards.Contract.Earned(&_IConvexRewards.CallOpts, account)
}

// ExtraRewards is a free data retrieval call binding the contract method 0x40c35446.
//
// Solidity: function extraRewards(uint256 _reward) view returns(address)
func (_IConvexRewards *IConvexRewardsCaller) ExtraRewards(opts *bind.CallOpts, _reward *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IConvexRewards.contract.Call(opts, &out, "extraRewards", _reward)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExtraRewards is a free data retrieval call binding the contract method 0x40c35446.
//
// Solidity: function extraRewards(uint256 _reward) view returns(address)
func (_IConvexRewards *IConvexRewardsSession) ExtraRewards(_reward *big.Int) (common.Address, error) {
	return _IConvexRewards.Contract.ExtraRewards(&_IConvexRewards.CallOpts, _reward)
}

// ExtraRewards is a free data retrieval call binding the contract method 0x40c35446.
//
// Solidity: function extraRewards(uint256 _reward) view returns(address)
func (_IConvexRewards *IConvexRewardsCallerSession) ExtraRewards(_reward *big.Int) (common.Address, error) {
	return _IConvexRewards.Contract.ExtraRewards(&_IConvexRewards.CallOpts, _reward)
}

// ExtraRewardsLength is a free data retrieval call binding the contract method 0xd55a23f4.
//
// Solidity: function extraRewardsLength() view returns(uint256)
func (_IConvexRewards *IConvexRewardsCaller) ExtraRewardsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IConvexRewards.contract.Call(opts, &out, "extraRewardsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExtraRewardsLength is a free data retrieval call binding the contract method 0xd55a23f4.
//
// Solidity: function extraRewardsLength() view returns(uint256)
func (_IConvexRewards *IConvexRewardsSession) ExtraRewardsLength() (*big.Int, error) {
	return _IConvexRewards.Contract.ExtraRewardsLength(&_IConvexRewards.CallOpts)
}

// ExtraRewardsLength is a free data retrieval call binding the contract method 0xd55a23f4.
//
// Solidity: function extraRewardsLength() view returns(uint256)
func (_IConvexRewards *IConvexRewardsCallerSession) ExtraRewardsLength() (*big.Int, error) {
	return _IConvexRewards.Contract.ExtraRewardsLength(&_IConvexRewards.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_IConvexRewards *IConvexRewardsCaller) PeriodFinish(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IConvexRewards.contract.Call(opts, &out, "periodFinish")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_IConvexRewards *IConvexRewardsSession) PeriodFinish() (*big.Int, error) {
	return _IConvexRewards.Contract.PeriodFinish(&_IConvexRewards.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_IConvexRewards *IConvexRewardsCallerSession) PeriodFinish() (*big.Int, error) {
	return _IConvexRewards.Contract.PeriodFinish(&_IConvexRewards.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_IConvexRewards *IConvexRewardsCaller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IConvexRewards.contract.Call(opts, &out, "rewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_IConvexRewards *IConvexRewardsSession) RewardToken() (common.Address, error) {
	return _IConvexRewards.Contract.RewardToken(&_IConvexRewards.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_IConvexRewards *IConvexRewardsCallerSession) RewardToken() (common.Address, error) {
	return _IConvexRewards.Contract.RewardToken(&_IConvexRewards.CallOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x7050ccd9.
//
// Solidity: function getReward(address _account, bool _claimExtras) returns(bool)
func (_IConvexRewards *IConvexRewardsTransactor) GetReward(opts *bind.TransactOpts, _account common.Address, _claimExtras bool) (*types.Transaction, error) {
	return _IConvexRewards.contract.Transact(opts, "getReward", _account, _claimExtras)
}

// GetReward is a paid mutator transaction binding the contract method 0x7050ccd9.
//
// Solidity: function getReward(address _account, bool _claimExtras) returns(bool)
func (_IConvexRewards *IConvexRewardsSession) GetReward(_account common.Address, _claimExtras bool) (*types.Transaction, error) {
	return _IConvexRewards.Contract.GetReward(&_IConvexRewards.TransactOpts, _account, _claimExtras)
}

// GetReward is a paid mutator transaction binding the contract method 0x7050ccd9.
//
// Solidity: function getReward(address _account, bool _claimExtras) returns(bool)
func (_IConvexRewards *IConvexRewardsTransactorSession) GetReward(_account common.Address, _claimExtras bool) (*types.Transaction, error) {
	return _IConvexRewards.Contract.GetReward(&_IConvexRewards.TransactOpts, _account, _claimExtras)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns(bool)
func (_IConvexRewards *IConvexRewardsTransactor) Stake(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _IConvexRewards.contract.Transact(opts, "stake", _amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns(bool)
func (_IConvexRewards *IConvexRewardsSession) Stake(_amount *big.Int) (*types.Transaction, error) {
	return _IConvexRewards.Contract.Stake(&_IConvexRewards.TransactOpts, _amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns(bool)
func (_IConvexRewards *IConvexRewardsTransactorSession) Stake(_amount *big.Int) (*types.Transaction, error) {
	return _IConvexRewards.Contract.Stake(&_IConvexRewards.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x38d07436.
//
// Solidity: function withdraw(uint256 _amount, bool _claim) returns(bool)
func (_IConvexRewards *IConvexRewardsTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int, _claim bool) (*types.Transaction, error) {
	return _IConvexRewards.contract.Transact(opts, "withdraw", _amount, _claim)
}

// Withdraw is a paid mutator transaction binding the contract method 0x38d07436.
//
// Solidity: function withdraw(uint256 _amount, bool _claim) returns(bool)
func (_IConvexRewards *IConvexRewardsSession) Withdraw(_amount *big.Int, _claim bool) (*types.Transaction, error) {
	return _IConvexRewards.Contract.Withdraw(&_IConvexRewards.TransactOpts, _amount, _claim)
}

// Withdraw is a paid mutator transaction binding the contract method 0x38d07436.
//
// Solidity: function withdraw(uint256 _amount, bool _claim) returns(bool)
func (_IConvexRewards *IConvexRewardsTransactorSession) Withdraw(_amount *big.Int, _claim bool) (*types.Transaction, error) {
	return _IConvexRewards.Contract.Withdraw(&_IConvexRewards.TransactOpts, _amount, _claim)
}

// WithdrawAndUnwrap is a paid mutator transaction binding the contract method 0xc32e7202.
//
// Solidity: function withdrawAndUnwrap(uint256 _amount, bool _claim) returns(bool)
func (_IConvexRewards *IConvexRewardsTransactor) WithdrawAndUnwrap(opts *bind.TransactOpts, _amount *big.Int, _claim bool) (*types.Transaction, error) {
	return _IConvexRewards.contract.Transact(opts, "withdrawAndUnwrap", _amount, _claim)
}

// WithdrawAndUnwrap is a paid mutator transaction binding the contract method 0xc32e7202.
//
// Solidity: function withdrawAndUnwrap(uint256 _amount, bool _claim) returns(bool)
func (_IConvexRewards *IConvexRewardsSession) WithdrawAndUnwrap(_amount *big.Int, _claim bool) (*types.Transaction, error) {
	return _IConvexRewards.Contract.WithdrawAndUnwrap(&_IConvexRewards.TransactOpts, _amount, _claim)
}

// WithdrawAndUnwrap is a paid mutator transaction binding the contract method 0xc32e7202.
//
// Solidity: function withdrawAndUnwrap(uint256 _amount, bool _claim) returns(bool)
func (_IConvexRewards *IConvexRewardsTransactorSession) WithdrawAndUnwrap(_amount *big.Int, _claim bool) (*types.Transaction, error) {
	return _IConvexRewards.Contract.WithdrawAndUnwrap(&_IConvexRewards.TransactOpts, _amount, _claim)
}

// ICurveFiMetaData contains all meta data concerning the ICurveFi contract.
var ICurveFiMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"amounts\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256\",\"name\":\"min_mint_amount\",\"type\":\"uint256\"}],\"name\":\"add_liquidity\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"amounts\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"min_mint_amount\",\"type\":\"uint256\"}],\"name\":\"add_liquidity\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[3]\",\"name\":\"amounts\",\"type\":\"uint256[3]\"},{\"internalType\":\"uint256\",\"name\":\"min_mint_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_use_underlying\",\"type\":\"bool\"}],\"name\":\"add_liquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256[4]\",\"name\":\"amounts\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256\",\"name\":\"min_mint_amount\",\"type\":\"uint256\"}],\"name\":\"add_liquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[3]\",\"name\":\"amounts\",\"type\":\"uint256[3]\"},{\"internalType\":\"uint256\",\"name\":\"min_mint_amount\",\"type\":\"uint256\"}],\"name\":\"add_liquidity\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"amounts\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256\",\"name\":\"min_mint_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_use_underlying\",\"type\":\"bool\"}],\"name\":\"add_liquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"amounts\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"min_mint_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_use_underlying\",\"type\":\"bool\"}],\"name\":\"add_liquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[3]\",\"name\":\"_amounts\",\"type\":\"uint256[3]\"},{\"internalType\":\"bool\",\"name\":\"_is_deposit\",\"type\":\"bool\"}],\"name\":\"calc_token_amount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"},{\"internalType\":\"uint256[4]\",\"name\":\"_amounts\",\"type\":\"uint256[4]\"},{\"internalType\":\"bool\",\"name\":\"_is_deposit\",\"type\":\"bool\"}],\"name\":\"calc_token_amount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"_amounts\",\"type\":\"uint256[4]\"},{\"internalType\":\"bool\",\"name\":\"_is_deposit\",\"type\":\"bool\"}],\"name\":\"calc_token_amount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"_amounts\",\"type\":\"uint256[2]\"},{\"internalType\":\"bool\",\"name\":\"_is_deposit\",\"type\":\"bool\"}],\"name\":\"calc_token_amount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"}],\"name\":\"calc_withdraw_one_coin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"coins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_from_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_min_to_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"use_eth\",\"type\":\"bool\"}],\"name\":\"exchange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"from\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"to\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"_from_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_min_to_amount\",\"type\":\"uint256\"}],\"name\":\"exchange\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"from\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"to\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"_from_amount\",\"type\":\"uint256\"}],\"name\":\"get_dy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_virtual_price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price_oracle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[2]\",\"name\":\"amounts\",\"type\":\"uint256[2]\"}],\"name\":\"remove_liquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"amounts\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"max_burn_amount\",\"type\":\"uint256\"}],\"name\":\"remove_liquidity_imbalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_token_amount\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"min_amount\",\"type\":\"uint256\"}],\"name\":\"remove_liquidity_one_coin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"384e03db": "add_liquidity(address,uint256[4],uint256)",
		"0b4c7e4d": "add_liquidity(uint256[2],uint256)",
		"ee22be23": "add_liquidity(uint256[2],uint256,bool)",
		"4515cef3": "add_liquidity(uint256[3],uint256)",
		"2b6e993a": "add_liquidity(uint256[3],uint256,bool)",
		"029b2f34": "add_liquidity(uint256[4],uint256)",
		"dc3a2d81": "add_liquidity(uint256[4],uint256,bool)",
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"4903b0d1": "balances(uint256)",
		"861cdef0": "calc_token_amount(address,uint256[4],bool)",
		"ed8e84f3": "calc_token_amount(uint256[2],bool)",
		"3883e119": "calc_token_amount(uint256[3],bool)",
		"cf701ff7": "calc_token_amount(uint256[4],bool)",
		"cc2b27d7": "calc_withdraw_one_coin(uint256,int128)",
		"c6610657": "coins(uint256)",
		"3df02124": "exchange(int128,int128,uint256,uint256)",
		"394747c5": "exchange(uint256,uint256,uint256,uint256,bool)",
		"5e0d443f": "get_dy(int128,int128,uint256)",
		"bb7b8b80": "get_virtual_price()",
		"86fc88d3": "price_oracle()",
		"5b36389c": "remove_liquidity(uint256,uint256[2])",
		"e3103273": "remove_liquidity_imbalance(uint256[2],uint256)",
		"1a4d01d2": "remove_liquidity_one_coin(uint256,int128,uint256)",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
}

// ICurveFiABI is the input ABI used to generate the binding from.
// Deprecated: Use ICurveFiMetaData.ABI instead.
var ICurveFiABI = ICurveFiMetaData.ABI

// Deprecated: Use ICurveFiMetaData.Sigs instead.
// ICurveFiFuncSigs maps the 4-byte function signature to its string representation.
var ICurveFiFuncSigs = ICurveFiMetaData.Sigs

// ICurveFi is an auto generated Go binding around an Ethereum contract.
type ICurveFi struct {
	ICurveFiCaller     // Read-only binding to the contract
	ICurveFiTransactor // Write-only binding to the contract
	ICurveFiFilterer   // Log filterer for contract events
}

// ICurveFiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICurveFiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICurveFiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICurveFiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICurveFiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICurveFiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICurveFiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICurveFiSession struct {
	Contract     *ICurveFi         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ICurveFiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICurveFiCallerSession struct {
	Contract *ICurveFiCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ICurveFiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICurveFiTransactorSession struct {
	Contract     *ICurveFiTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ICurveFiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICurveFiRaw struct {
	Contract *ICurveFi // Generic contract binding to access the raw methods on
}

// ICurveFiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICurveFiCallerRaw struct {
	Contract *ICurveFiCaller // Generic read-only contract binding to access the raw methods on
}

// ICurveFiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICurveFiTransactorRaw struct {
	Contract *ICurveFiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICurveFi creates a new instance of ICurveFi, bound to a specific deployed contract.
func NewICurveFi(address common.Address, backend bind.ContractBackend) (*ICurveFi, error) {
	contract, err := bindICurveFi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICurveFi{ICurveFiCaller: ICurveFiCaller{contract: contract}, ICurveFiTransactor: ICurveFiTransactor{contract: contract}, ICurveFiFilterer: ICurveFiFilterer{contract: contract}}, nil
}

// NewICurveFiCaller creates a new read-only instance of ICurveFi, bound to a specific deployed contract.
func NewICurveFiCaller(address common.Address, caller bind.ContractCaller) (*ICurveFiCaller, error) {
	contract, err := bindICurveFi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICurveFiCaller{contract: contract}, nil
}

// NewICurveFiTransactor creates a new write-only instance of ICurveFi, bound to a specific deployed contract.
func NewICurveFiTransactor(address common.Address, transactor bind.ContractTransactor) (*ICurveFiTransactor, error) {
	contract, err := bindICurveFi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICurveFiTransactor{contract: contract}, nil
}

// NewICurveFiFilterer creates a new log filterer instance of ICurveFi, bound to a specific deployed contract.
func NewICurveFiFilterer(address common.Address, filterer bind.ContractFilterer) (*ICurveFiFilterer, error) {
	contract, err := bindICurveFi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICurveFiFilterer{contract: contract}, nil
}

// bindICurveFi binds a generic wrapper to an already deployed contract.
func bindICurveFi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ICurveFiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICurveFi *ICurveFiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICurveFi.Contract.ICurveFiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICurveFi *ICurveFiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICurveFi.Contract.ICurveFiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICurveFi *ICurveFiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICurveFi.Contract.ICurveFiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICurveFi *ICurveFiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICurveFi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICurveFi *ICurveFiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICurveFi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICurveFi *ICurveFiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICurveFi.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ICurveFi *ICurveFiCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ICurveFi.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ICurveFi *ICurveFiSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ICurveFi.Contract.Allowance(&_ICurveFi.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ICurveFi *ICurveFiCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ICurveFi.Contract.Allowance(&_ICurveFi.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ICurveFi *ICurveFiCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ICurveFi.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ICurveFi *ICurveFiSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ICurveFi.Contract.BalanceOf(&_ICurveFi.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ICurveFi *ICurveFiCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ICurveFi.Contract.BalanceOf(&_ICurveFi.CallOpts, account)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 ) view returns(uint256)
func (_ICurveFi *ICurveFiCaller) Balances(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ICurveFi.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 ) view returns(uint256)
func (_ICurveFi *ICurveFiSession) Balances(arg0 *big.Int) (*big.Int, error) {
	return _ICurveFi.Contract.Balances(&_ICurveFi.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 ) view returns(uint256)
func (_ICurveFi *ICurveFiCallerSession) Balances(arg0 *big.Int) (*big.Int, error) {
	return _ICurveFi.Contract.Balances(&_ICurveFi.CallOpts, arg0)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x3883e119.
//
// Solidity: function calc_token_amount(uint256[3] _amounts, bool _is_deposit) view returns(uint256)
func (_ICurveFi *ICurveFiCaller) CalcTokenAmount(opts *bind.CallOpts, _amounts [3]*big.Int, _is_deposit bool) (*big.Int, error) {
	var out []interface{}
	err := _ICurveFi.contract.Call(opts, &out, "calc_token_amount", _amounts, _is_deposit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x3883e119.
//
// Solidity: function calc_token_amount(uint256[3] _amounts, bool _is_deposit) view returns(uint256)
func (_ICurveFi *ICurveFiSession) CalcTokenAmount(_amounts [3]*big.Int, _is_deposit bool) (*big.Int, error) {
	return _ICurveFi.Contract.CalcTokenAmount(&_ICurveFi.CallOpts, _amounts, _is_deposit)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x3883e119.
//
// Solidity: function calc_token_amount(uint256[3] _amounts, bool _is_deposit) view returns(uint256)
func (_ICurveFi *ICurveFiCallerSession) CalcTokenAmount(_amounts [3]*big.Int, _is_deposit bool) (*big.Int, error) {
	return _ICurveFi.Contract.CalcTokenAmount(&_ICurveFi.CallOpts, _amounts, _is_deposit)
}

// CalcTokenAmount0 is a free data retrieval call binding the contract method 0x861cdef0.
//
// Solidity: function calc_token_amount(address _pool, uint256[4] _amounts, bool _is_deposit) view returns(uint256)
func (_ICurveFi *ICurveFiCaller) CalcTokenAmount0(opts *bind.CallOpts, _pool common.Address, _amounts [4]*big.Int, _is_deposit bool) (*big.Int, error) {
	var out []interface{}
	err := _ICurveFi.contract.Call(opts, &out, "calc_token_amount0", _pool, _amounts, _is_deposit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcTokenAmount0 is a free data retrieval call binding the contract method 0x861cdef0.
//
// Solidity: function calc_token_amount(address _pool, uint256[4] _amounts, bool _is_deposit) view returns(uint256)
func (_ICurveFi *ICurveFiSession) CalcTokenAmount0(_pool common.Address, _amounts [4]*big.Int, _is_deposit bool) (*big.Int, error) {
	return _ICurveFi.Contract.CalcTokenAmount0(&_ICurveFi.CallOpts, _pool, _amounts, _is_deposit)
}

// CalcTokenAmount0 is a free data retrieval call binding the contract method 0x861cdef0.
//
// Solidity: function calc_token_amount(address _pool, uint256[4] _amounts, bool _is_deposit) view returns(uint256)
func (_ICurveFi *ICurveFiCallerSession) CalcTokenAmount0(_pool common.Address, _amounts [4]*big.Int, _is_deposit bool) (*big.Int, error) {
	return _ICurveFi.Contract.CalcTokenAmount0(&_ICurveFi.CallOpts, _pool, _amounts, _is_deposit)
}

// CalcTokenAmount1 is a free data retrieval call binding the contract method 0xcf701ff7.
//
// Solidity: function calc_token_amount(uint256[4] _amounts, bool _is_deposit) view returns(uint256)
func (_ICurveFi *ICurveFiCaller) CalcTokenAmount1(opts *bind.CallOpts, _amounts [4]*big.Int, _is_deposit bool) (*big.Int, error) {
	var out []interface{}
	err := _ICurveFi.contract.Call(opts, &out, "calc_token_amount1", _amounts, _is_deposit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcTokenAmount1 is a free data retrieval call binding the contract method 0xcf701ff7.
//
// Solidity: function calc_token_amount(uint256[4] _amounts, bool _is_deposit) view returns(uint256)
func (_ICurveFi *ICurveFiSession) CalcTokenAmount1(_amounts [4]*big.Int, _is_deposit bool) (*big.Int, error) {
	return _ICurveFi.Contract.CalcTokenAmount1(&_ICurveFi.CallOpts, _amounts, _is_deposit)
}

// CalcTokenAmount1 is a free data retrieval call binding the contract method 0xcf701ff7.
//
// Solidity: function calc_token_amount(uint256[4] _amounts, bool _is_deposit) view returns(uint256)
func (_ICurveFi *ICurveFiCallerSession) CalcTokenAmount1(_amounts [4]*big.Int, _is_deposit bool) (*big.Int, error) {
	return _ICurveFi.Contract.CalcTokenAmount1(&_ICurveFi.CallOpts, _amounts, _is_deposit)
}

// CalcTokenAmount2 is a free data retrieval call binding the contract method 0xed8e84f3.
//
// Solidity: function calc_token_amount(uint256[2] _amounts, bool _is_deposit) view returns(uint256)
func (_ICurveFi *ICurveFiCaller) CalcTokenAmount2(opts *bind.CallOpts, _amounts [2]*big.Int, _is_deposit bool) (*big.Int, error) {
	var out []interface{}
	err := _ICurveFi.contract.Call(opts, &out, "calc_token_amount2", _amounts, _is_deposit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcTokenAmount2 is a free data retrieval call binding the contract method 0xed8e84f3.
//
// Solidity: function calc_token_amount(uint256[2] _amounts, bool _is_deposit) view returns(uint256)
func (_ICurveFi *ICurveFiSession) CalcTokenAmount2(_amounts [2]*big.Int, _is_deposit bool) (*big.Int, error) {
	return _ICurveFi.Contract.CalcTokenAmount2(&_ICurveFi.CallOpts, _amounts, _is_deposit)
}

// CalcTokenAmount2 is a free data retrieval call binding the contract method 0xed8e84f3.
//
// Solidity: function calc_token_amount(uint256[2] _amounts, bool _is_deposit) view returns(uint256)
func (_ICurveFi *ICurveFiCallerSession) CalcTokenAmount2(_amounts [2]*big.Int, _is_deposit bool) (*big.Int, error) {
	return _ICurveFi.Contract.CalcTokenAmount2(&_ICurveFi.CallOpts, _amounts, _is_deposit)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 amount, int128 i) view returns(uint256)
func (_ICurveFi *ICurveFiCaller) CalcWithdrawOneCoin(opts *bind.CallOpts, amount *big.Int, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ICurveFi.contract.Call(opts, &out, "calc_withdraw_one_coin", amount, i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 amount, int128 i) view returns(uint256)
func (_ICurveFi *ICurveFiSession) CalcWithdrawOneCoin(amount *big.Int, i *big.Int) (*big.Int, error) {
	return _ICurveFi.Contract.CalcWithdrawOneCoin(&_ICurveFi.CallOpts, amount, i)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 amount, int128 i) view returns(uint256)
func (_ICurveFi *ICurveFiCallerSession) CalcWithdrawOneCoin(amount *big.Int, i *big.Int) (*big.Int, error) {
	return _ICurveFi.Contract.CalcWithdrawOneCoin(&_ICurveFi.CallOpts, amount, i)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 ) view returns(address)
func (_ICurveFi *ICurveFiCaller) Coins(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ICurveFi.contract.Call(opts, &out, "coins", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 ) view returns(address)
func (_ICurveFi *ICurveFiSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _ICurveFi.Contract.Coins(&_ICurveFi.CallOpts, arg0)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 ) view returns(address)
func (_ICurveFi *ICurveFiCallerSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _ICurveFi.Contract.Coins(&_ICurveFi.CallOpts, arg0)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 from, int128 to, uint256 _from_amount) view returns(uint256)
func (_ICurveFi *ICurveFiCaller) GetDy(opts *bind.CallOpts, from *big.Int, to *big.Int, _from_amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ICurveFi.contract.Call(opts, &out, "get_dy", from, to, _from_amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 from, int128 to, uint256 _from_amount) view returns(uint256)
func (_ICurveFi *ICurveFiSession) GetDy(from *big.Int, to *big.Int, _from_amount *big.Int) (*big.Int, error) {
	return _ICurveFi.Contract.GetDy(&_ICurveFi.CallOpts, from, to, _from_amount)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 from, int128 to, uint256 _from_amount) view returns(uint256)
func (_ICurveFi *ICurveFiCallerSession) GetDy(from *big.Int, to *big.Int, _from_amount *big.Int) (*big.Int, error) {
	return _ICurveFi.Contract.GetDy(&_ICurveFi.CallOpts, from, to, _from_amount)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_ICurveFi *ICurveFiCaller) GetVirtualPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ICurveFi.contract.Call(opts, &out, "get_virtual_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_ICurveFi *ICurveFiSession) GetVirtualPrice() (*big.Int, error) {
	return _ICurveFi.Contract.GetVirtualPrice(&_ICurveFi.CallOpts)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_ICurveFi *ICurveFiCallerSession) GetVirtualPrice() (*big.Int, error) {
	return _ICurveFi.Contract.GetVirtualPrice(&_ICurveFi.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x86fc88d3.
//
// Solidity: function price_oracle() view returns(uint256)
func (_ICurveFi *ICurveFiCaller) PriceOracle(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ICurveFi.contract.Call(opts, &out, "price_oracle")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceOracle is a free data retrieval call binding the contract method 0x86fc88d3.
//
// Solidity: function price_oracle() view returns(uint256)
func (_ICurveFi *ICurveFiSession) PriceOracle() (*big.Int, error) {
	return _ICurveFi.Contract.PriceOracle(&_ICurveFi.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x86fc88d3.
//
// Solidity: function price_oracle() view returns(uint256)
func (_ICurveFi *ICurveFiCallerSession) PriceOracle() (*big.Int, error) {
	return _ICurveFi.Contract.PriceOracle(&_ICurveFi.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ICurveFi *ICurveFiCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ICurveFi.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ICurveFi *ICurveFiSession) TotalSupply() (*big.Int, error) {
	return _ICurveFi.Contract.TotalSupply(&_ICurveFi.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ICurveFi *ICurveFiCallerSession) TotalSupply() (*big.Int, error) {
	return _ICurveFi.Contract.TotalSupply(&_ICurveFi.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x029b2f34.
//
// Solidity: function add_liquidity(uint256[4] amounts, uint256 min_mint_amount) payable returns()
func (_ICurveFi *ICurveFiTransactor) AddLiquidity(opts *bind.TransactOpts, amounts [4]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _ICurveFi.contract.Transact(opts, "add_liquidity", amounts, min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x029b2f34.
//
// Solidity: function add_liquidity(uint256[4] amounts, uint256 min_mint_amount) payable returns()
func (_ICurveFi *ICurveFiSession) AddLiquidity(amounts [4]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _ICurveFi.Contract.AddLiquidity(&_ICurveFi.TransactOpts, amounts, min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x029b2f34.
//
// Solidity: function add_liquidity(uint256[4] amounts, uint256 min_mint_amount) payable returns()
func (_ICurveFi *ICurveFiTransactorSession) AddLiquidity(amounts [4]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _ICurveFi.Contract.AddLiquidity(&_ICurveFi.TransactOpts, amounts, min_mint_amount)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount) payable returns()
func (_ICurveFi *ICurveFiTransactor) AddLiquidity0(opts *bind.TransactOpts, amounts [2]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _ICurveFi.contract.Transact(opts, "add_liquidity0", amounts, min_mint_amount)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount) payable returns()
func (_ICurveFi *ICurveFiSession) AddLiquidity0(amounts [2]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _ICurveFi.Contract.AddLiquidity0(&_ICurveFi.TransactOpts, amounts, min_mint_amount)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount) payable returns()
func (_ICurveFi *ICurveFiTransactorSession) AddLiquidity0(amounts [2]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _ICurveFi.Contract.AddLiquidity0(&_ICurveFi.TransactOpts, amounts, min_mint_amount)
}

// AddLiquidity1 is a paid mutator transaction binding the contract method 0x2b6e993a.
//
// Solidity: function add_liquidity(uint256[3] amounts, uint256 min_mint_amount, bool _use_underlying) payable returns(uint256)
func (_ICurveFi *ICurveFiTransactor) AddLiquidity1(opts *bind.TransactOpts, amounts [3]*big.Int, min_mint_amount *big.Int, _use_underlying bool) (*types.Transaction, error) {
	return _ICurveFi.contract.Transact(opts, "add_liquidity1", amounts, min_mint_amount, _use_underlying)
}

// AddLiquidity1 is a paid mutator transaction binding the contract method 0x2b6e993a.
//
// Solidity: function add_liquidity(uint256[3] amounts, uint256 min_mint_amount, bool _use_underlying) payable returns(uint256)
func (_ICurveFi *ICurveFiSession) AddLiquidity1(amounts [3]*big.Int, min_mint_amount *big.Int, _use_underlying bool) (*types.Transaction, error) {
	return _ICurveFi.Contract.AddLiquidity1(&_ICurveFi.TransactOpts, amounts, min_mint_amount, _use_underlying)
}

// AddLiquidity1 is a paid mutator transaction binding the contract method 0x2b6e993a.
//
// Solidity: function add_liquidity(uint256[3] amounts, uint256 min_mint_amount, bool _use_underlying) payable returns(uint256)
func (_ICurveFi *ICurveFiTransactorSession) AddLiquidity1(amounts [3]*big.Int, min_mint_amount *big.Int, _use_underlying bool) (*types.Transaction, error) {
	return _ICurveFi.Contract.AddLiquidity1(&_ICurveFi.TransactOpts, amounts, min_mint_amount, _use_underlying)
}

// AddLiquidity2 is a paid mutator transaction binding the contract method 0x384e03db.
//
// Solidity: function add_liquidity(address pool, uint256[4] amounts, uint256 min_mint_amount) returns()
func (_ICurveFi *ICurveFiTransactor) AddLiquidity2(opts *bind.TransactOpts, pool common.Address, amounts [4]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _ICurveFi.contract.Transact(opts, "add_liquidity2", pool, amounts, min_mint_amount)
}

// AddLiquidity2 is a paid mutator transaction binding the contract method 0x384e03db.
//
// Solidity: function add_liquidity(address pool, uint256[4] amounts, uint256 min_mint_amount) returns()
func (_ICurveFi *ICurveFiSession) AddLiquidity2(pool common.Address, amounts [4]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _ICurveFi.Contract.AddLiquidity2(&_ICurveFi.TransactOpts, pool, amounts, min_mint_amount)
}

// AddLiquidity2 is a paid mutator transaction binding the contract method 0x384e03db.
//
// Solidity: function add_liquidity(address pool, uint256[4] amounts, uint256 min_mint_amount) returns()
func (_ICurveFi *ICurveFiTransactorSession) AddLiquidity2(pool common.Address, amounts [4]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _ICurveFi.Contract.AddLiquidity2(&_ICurveFi.TransactOpts, pool, amounts, min_mint_amount)
}

// AddLiquidity3 is a paid mutator transaction binding the contract method 0x4515cef3.
//
// Solidity: function add_liquidity(uint256[3] amounts, uint256 min_mint_amount) payable returns()
func (_ICurveFi *ICurveFiTransactor) AddLiquidity3(opts *bind.TransactOpts, amounts [3]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _ICurveFi.contract.Transact(opts, "add_liquidity3", amounts, min_mint_amount)
}

// AddLiquidity3 is a paid mutator transaction binding the contract method 0x4515cef3.
//
// Solidity: function add_liquidity(uint256[3] amounts, uint256 min_mint_amount) payable returns()
func (_ICurveFi *ICurveFiSession) AddLiquidity3(amounts [3]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _ICurveFi.Contract.AddLiquidity3(&_ICurveFi.TransactOpts, amounts, min_mint_amount)
}

// AddLiquidity3 is a paid mutator transaction binding the contract method 0x4515cef3.
//
// Solidity: function add_liquidity(uint256[3] amounts, uint256 min_mint_amount) payable returns()
func (_ICurveFi *ICurveFiTransactorSession) AddLiquidity3(amounts [3]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _ICurveFi.Contract.AddLiquidity3(&_ICurveFi.TransactOpts, amounts, min_mint_amount)
}

// AddLiquidity4 is a paid mutator transaction binding the contract method 0xdc3a2d81.
//
// Solidity: function add_liquidity(uint256[4] amounts, uint256 min_mint_amount, bool _use_underlying) payable returns(uint256)
func (_ICurveFi *ICurveFiTransactor) AddLiquidity4(opts *bind.TransactOpts, amounts [4]*big.Int, min_mint_amount *big.Int, _use_underlying bool) (*types.Transaction, error) {
	return _ICurveFi.contract.Transact(opts, "add_liquidity4", amounts, min_mint_amount, _use_underlying)
}

// AddLiquidity4 is a paid mutator transaction binding the contract method 0xdc3a2d81.
//
// Solidity: function add_liquidity(uint256[4] amounts, uint256 min_mint_amount, bool _use_underlying) payable returns(uint256)
func (_ICurveFi *ICurveFiSession) AddLiquidity4(amounts [4]*big.Int, min_mint_amount *big.Int, _use_underlying bool) (*types.Transaction, error) {
	return _ICurveFi.Contract.AddLiquidity4(&_ICurveFi.TransactOpts, amounts, min_mint_amount, _use_underlying)
}

// AddLiquidity4 is a paid mutator transaction binding the contract method 0xdc3a2d81.
//
// Solidity: function add_liquidity(uint256[4] amounts, uint256 min_mint_amount, bool _use_underlying) payable returns(uint256)
func (_ICurveFi *ICurveFiTransactorSession) AddLiquidity4(amounts [4]*big.Int, min_mint_amount *big.Int, _use_underlying bool) (*types.Transaction, error) {
	return _ICurveFi.Contract.AddLiquidity4(&_ICurveFi.TransactOpts, amounts, min_mint_amount, _use_underlying)
}

// AddLiquidity5 is a paid mutator transaction binding the contract method 0xee22be23.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount, bool _use_underlying) payable returns(uint256)
func (_ICurveFi *ICurveFiTransactor) AddLiquidity5(opts *bind.TransactOpts, amounts [2]*big.Int, min_mint_amount *big.Int, _use_underlying bool) (*types.Transaction, error) {
	return _ICurveFi.contract.Transact(opts, "add_liquidity5", amounts, min_mint_amount, _use_underlying)
}

// AddLiquidity5 is a paid mutator transaction binding the contract method 0xee22be23.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount, bool _use_underlying) payable returns(uint256)
func (_ICurveFi *ICurveFiSession) AddLiquidity5(amounts [2]*big.Int, min_mint_amount *big.Int, _use_underlying bool) (*types.Transaction, error) {
	return _ICurveFi.Contract.AddLiquidity5(&_ICurveFi.TransactOpts, amounts, min_mint_amount, _use_underlying)
}

// AddLiquidity5 is a paid mutator transaction binding the contract method 0xee22be23.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount, bool _use_underlying) payable returns(uint256)
func (_ICurveFi *ICurveFiTransactorSession) AddLiquidity5(amounts [2]*big.Int, min_mint_amount *big.Int, _use_underlying bool) (*types.Transaction, error) {
	return _ICurveFi.Contract.AddLiquidity5(&_ICurveFi.TransactOpts, amounts, min_mint_amount, _use_underlying)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ICurveFi *ICurveFiTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ICurveFi.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ICurveFi *ICurveFiSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ICurveFi.Contract.Approve(&_ICurveFi.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ICurveFi *ICurveFiTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ICurveFi.Contract.Approve(&_ICurveFi.TransactOpts, spender, amount)
}

// Exchange is a paid mutator transaction binding the contract method 0x394747c5.
//
// Solidity: function exchange(uint256 from, uint256 to, uint256 _from_amount, uint256 _min_to_amount, bool use_eth) returns()
func (_ICurveFi *ICurveFiTransactor) Exchange(opts *bind.TransactOpts, from *big.Int, to *big.Int, _from_amount *big.Int, _min_to_amount *big.Int, use_eth bool) (*types.Transaction, error) {
	return _ICurveFi.contract.Transact(opts, "exchange", from, to, _from_amount, _min_to_amount, use_eth)
}

// Exchange is a paid mutator transaction binding the contract method 0x394747c5.
//
// Solidity: function exchange(uint256 from, uint256 to, uint256 _from_amount, uint256 _min_to_amount, bool use_eth) returns()
func (_ICurveFi *ICurveFiSession) Exchange(from *big.Int, to *big.Int, _from_amount *big.Int, _min_to_amount *big.Int, use_eth bool) (*types.Transaction, error) {
	return _ICurveFi.Contract.Exchange(&_ICurveFi.TransactOpts, from, to, _from_amount, _min_to_amount, use_eth)
}

// Exchange is a paid mutator transaction binding the contract method 0x394747c5.
//
// Solidity: function exchange(uint256 from, uint256 to, uint256 _from_amount, uint256 _min_to_amount, bool use_eth) returns()
func (_ICurveFi *ICurveFiTransactorSession) Exchange(from *big.Int, to *big.Int, _from_amount *big.Int, _min_to_amount *big.Int, use_eth bool) (*types.Transaction, error) {
	return _ICurveFi.Contract.Exchange(&_ICurveFi.TransactOpts, from, to, _from_amount, _min_to_amount, use_eth)
}

// Exchange0 is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 from, int128 to, uint256 _from_amount, uint256 _min_to_amount) payable returns(uint256)
func (_ICurveFi *ICurveFiTransactor) Exchange0(opts *bind.TransactOpts, from *big.Int, to *big.Int, _from_amount *big.Int, _min_to_amount *big.Int) (*types.Transaction, error) {
	return _ICurveFi.contract.Transact(opts, "exchange0", from, to, _from_amount, _min_to_amount)
}

// Exchange0 is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 from, int128 to, uint256 _from_amount, uint256 _min_to_amount) payable returns(uint256)
func (_ICurveFi *ICurveFiSession) Exchange0(from *big.Int, to *big.Int, _from_amount *big.Int, _min_to_amount *big.Int) (*types.Transaction, error) {
	return _ICurveFi.Contract.Exchange0(&_ICurveFi.TransactOpts, from, to, _from_amount, _min_to_amount)
}

// Exchange0 is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 from, int128 to, uint256 _from_amount, uint256 _min_to_amount) payable returns(uint256)
func (_ICurveFi *ICurveFiTransactorSession) Exchange0(from *big.Int, to *big.Int, _from_amount *big.Int, _min_to_amount *big.Int) (*types.Transaction, error) {
	return _ICurveFi.Contract.Exchange0(&_ICurveFi.TransactOpts, from, to, _from_amount, _min_to_amount)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] amounts) returns()
func (_ICurveFi *ICurveFiTransactor) RemoveLiquidity(opts *bind.TransactOpts, _amount *big.Int, amounts [2]*big.Int) (*types.Transaction, error) {
	return _ICurveFi.contract.Transact(opts, "remove_liquidity", _amount, amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] amounts) returns()
func (_ICurveFi *ICurveFiSession) RemoveLiquidity(_amount *big.Int, amounts [2]*big.Int) (*types.Transaction, error) {
	return _ICurveFi.Contract.RemoveLiquidity(&_ICurveFi.TransactOpts, _amount, amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] amounts) returns()
func (_ICurveFi *ICurveFiTransactorSession) RemoveLiquidity(_amount *big.Int, amounts [2]*big.Int) (*types.Transaction, error) {
	return _ICurveFi.Contract.RemoveLiquidity(&_ICurveFi.TransactOpts, _amount, amounts)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0xe3103273.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] amounts, uint256 max_burn_amount) returns()
func (_ICurveFi *ICurveFiTransactor) RemoveLiquidityImbalance(opts *bind.TransactOpts, amounts [2]*big.Int, max_burn_amount *big.Int) (*types.Transaction, error) {
	return _ICurveFi.contract.Transact(opts, "remove_liquidity_imbalance", amounts, max_burn_amount)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0xe3103273.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] amounts, uint256 max_burn_amount) returns()
func (_ICurveFi *ICurveFiSession) RemoveLiquidityImbalance(amounts [2]*big.Int, max_burn_amount *big.Int) (*types.Transaction, error) {
	return _ICurveFi.Contract.RemoveLiquidityImbalance(&_ICurveFi.TransactOpts, amounts, max_burn_amount)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0xe3103273.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] amounts, uint256 max_burn_amount) returns()
func (_ICurveFi *ICurveFiTransactorSession) RemoveLiquidityImbalance(amounts [2]*big.Int, max_burn_amount *big.Int) (*types.Transaction, error) {
	return _ICurveFi.Contract.RemoveLiquidityImbalance(&_ICurveFi.TransactOpts, amounts, max_burn_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _token_amount, int128 i, uint256 min_amount) returns()
func (_ICurveFi *ICurveFiTransactor) RemoveLiquidityOneCoin(opts *bind.TransactOpts, _token_amount *big.Int, i *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _ICurveFi.contract.Transact(opts, "remove_liquidity_one_coin", _token_amount, i, min_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _token_amount, int128 i, uint256 min_amount) returns()
func (_ICurveFi *ICurveFiSession) RemoveLiquidityOneCoin(_token_amount *big.Int, i *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _ICurveFi.Contract.RemoveLiquidityOneCoin(&_ICurveFi.TransactOpts, _token_amount, i, min_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _token_amount, int128 i, uint256 min_amount) returns()
func (_ICurveFi *ICurveFiTransactorSession) RemoveLiquidityOneCoin(_token_amount *big.Int, i *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _ICurveFi.Contract.RemoveLiquidityOneCoin(&_ICurveFi.TransactOpts, _token_amount, i, min_amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ICurveFi *ICurveFiTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ICurveFi.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ICurveFi *ICurveFiSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ICurveFi.Contract.Transfer(&_ICurveFi.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ICurveFi *ICurveFiTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ICurveFi.Contract.Transfer(&_ICurveFi.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ICurveFi *ICurveFiTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ICurveFi.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ICurveFi *ICurveFiSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ICurveFi.Contract.TransferFrom(&_ICurveFi.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ICurveFi *ICurveFiTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ICurveFi.Contract.TransferFrom(&_ICurveFi.TransactOpts, sender, recipient, amount)
}

// ICurveFiApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ICurveFi contract.
type ICurveFiApprovalIterator struct {
	Event *ICurveFiApproval // Event containing the contract specifics and raw log

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
func (it *ICurveFiApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICurveFiApproval)
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
		it.Event = new(ICurveFiApproval)
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
func (it *ICurveFiApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICurveFiApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICurveFiApproval represents a Approval event raised by the ICurveFi contract.
type ICurveFiApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ICurveFi *ICurveFiFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ICurveFiApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ICurveFi.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ICurveFiApprovalIterator{contract: _ICurveFi.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ICurveFi *ICurveFiFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ICurveFiApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ICurveFi.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICurveFiApproval)
				if err := _ICurveFi.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ICurveFi *ICurveFiFilterer) ParseApproval(log types.Log) (*ICurveFiApproval, error) {
	event := new(ICurveFiApproval)
	if err := _ICurveFi.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICurveFiTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ICurveFi contract.
type ICurveFiTransferIterator struct {
	Event *ICurveFiTransfer // Event containing the contract specifics and raw log

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
func (it *ICurveFiTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICurveFiTransfer)
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
		it.Event = new(ICurveFiTransfer)
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
func (it *ICurveFiTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICurveFiTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICurveFiTransfer represents a Transfer event raised by the ICurveFi contract.
type ICurveFiTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ICurveFi *ICurveFiFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ICurveFiTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ICurveFi.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ICurveFiTransferIterator{contract: _ICurveFi.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ICurveFi *ICurveFiFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ICurveFiTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ICurveFi.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICurveFiTransfer)
				if err := _ICurveFi.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ICurveFi *ICurveFiFilterer) ParseTransfer(log types.Log) (*ICurveFiTransfer, error) {
	event := new(ICurveFiTransfer)
	if err := _ICurveFi.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOracleMetaData contains all meta data concerning the IOracle contract.
var IOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"latestAnswer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"50d25bcd": "latestAnswer()",
	},
}

// IOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use IOracleMetaData.ABI instead.
var IOracleABI = IOracleMetaData.ABI

// Deprecated: Use IOracleMetaData.Sigs instead.
// IOracleFuncSigs maps the 4-byte function signature to its string representation.
var IOracleFuncSigs = IOracleMetaData.Sigs

// IOracle is an auto generated Go binding around an Ethereum contract.
type IOracle struct {
	IOracleCaller     // Read-only binding to the contract
	IOracleTransactor // Write-only binding to the contract
	IOracleFilterer   // Log filterer for contract events
}

// IOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type IOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IOracleSession struct {
	Contract     *IOracle          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IOracleCallerSession struct {
	Contract *IOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IOracleTransactorSession struct {
	Contract     *IOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type IOracleRaw struct {
	Contract *IOracle // Generic contract binding to access the raw methods on
}

// IOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IOracleCallerRaw struct {
	Contract *IOracleCaller // Generic read-only contract binding to access the raw methods on
}

// IOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IOracleTransactorRaw struct {
	Contract *IOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIOracle creates a new instance of IOracle, bound to a specific deployed contract.
func NewIOracle(address common.Address, backend bind.ContractBackend) (*IOracle, error) {
	contract, err := bindIOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IOracle{IOracleCaller: IOracleCaller{contract: contract}, IOracleTransactor: IOracleTransactor{contract: contract}, IOracleFilterer: IOracleFilterer{contract: contract}}, nil
}

// NewIOracleCaller creates a new read-only instance of IOracle, bound to a specific deployed contract.
func NewIOracleCaller(address common.Address, caller bind.ContractCaller) (*IOracleCaller, error) {
	contract, err := bindIOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IOracleCaller{contract: contract}, nil
}

// NewIOracleTransactor creates a new write-only instance of IOracle, bound to a specific deployed contract.
func NewIOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*IOracleTransactor, error) {
	contract, err := bindIOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IOracleTransactor{contract: contract}, nil
}

// NewIOracleFilterer creates a new log filterer instance of IOracle, bound to a specific deployed contract.
func NewIOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*IOracleFilterer, error) {
	contract, err := bindIOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IOracleFilterer{contract: contract}, nil
}

// bindIOracle binds a generic wrapper to an already deployed contract.
func bindIOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOracle *IOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOracle.Contract.IOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOracle *IOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOracle.Contract.IOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOracle *IOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOracle.Contract.IOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOracle *IOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOracle *IOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOracle *IOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOracle.Contract.contract.Transact(opts, method, params...)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(uint256)
func (_IOracle *IOracleCaller) LatestAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOracle.contract.Call(opts, &out, "latestAnswer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(uint256)
func (_IOracle *IOracleSession) LatestAnswer() (*big.Int, error) {
	return _IOracle.Contract.LatestAnswer(&_IOracle.CallOpts)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(uint256)
func (_IOracle *IOracleCallerSession) LatestAnswer() (*big.Int, error) {
	return _IOracle.Contract.LatestAnswer(&_IOracle.CallOpts)
}

// IUniswapV2Router01MetaData contains all meta data concerning the IUniswapV2Router01 contract.
var IUniswapV2Router01MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountADesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsIn\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityETHWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapETHForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactETHForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"ad5c4648": "WETH()",
		"e8e33700": "addLiquidity(address,address,uint256,uint256,uint256,uint256,address,uint256)",
		"f305d719": "addLiquidityETH(address,uint256,uint256,uint256,address,uint256)",
		"c45a0155": "factory()",
		"85f8c259": "getAmountIn(uint256,uint256,uint256)",
		"054d50d4": "getAmountOut(uint256,uint256,uint256)",
		"1f00ca74": "getAmountsIn(uint256,address[])",
		"d06ca61f": "getAmountsOut(uint256,address[])",
		"ad615dec": "quote(uint256,uint256,uint256)",
		"baa2abde": "removeLiquidity(address,address,uint256,uint256,uint256,address,uint256)",
		"02751cec": "removeLiquidityETH(address,uint256,uint256,uint256,address,uint256)",
		"ded9382a": "removeLiquidityETHWithPermit(address,uint256,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32)",
		"2195995c": "removeLiquidityWithPermit(address,address,uint256,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32)",
		"fb3bdb41": "swapETHForExactTokens(uint256,address[],address,uint256)",
		"7ff36ab5": "swapExactETHForTokens(uint256,address[],address,uint256)",
		"18cbafe5": "swapExactTokensForETH(uint256,uint256,address[],address,uint256)",
		"38ed1739": "swapExactTokensForTokens(uint256,uint256,address[],address,uint256)",
		"4a25d94a": "swapTokensForExactETH(uint256,uint256,address[],address,uint256)",
		"8803dbee": "swapTokensForExactTokens(uint256,uint256,address[],address,uint256)",
	},
}

// IUniswapV2Router01ABI is the input ABI used to generate the binding from.
// Deprecated: Use IUniswapV2Router01MetaData.ABI instead.
var IUniswapV2Router01ABI = IUniswapV2Router01MetaData.ABI

// Deprecated: Use IUniswapV2Router01MetaData.Sigs instead.
// IUniswapV2Router01FuncSigs maps the 4-byte function signature to its string representation.
var IUniswapV2Router01FuncSigs = IUniswapV2Router01MetaData.Sigs

// IUniswapV2Router01 is an auto generated Go binding around an Ethereum contract.
type IUniswapV2Router01 struct {
	IUniswapV2Router01Caller     // Read-only binding to the contract
	IUniswapV2Router01Transactor // Write-only binding to the contract
	IUniswapV2Router01Filterer   // Log filterer for contract events
}

// IUniswapV2Router01Caller is an auto generated read-only Go binding around an Ethereum contract.
type IUniswapV2Router01Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV2Router01Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IUniswapV2Router01Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV2Router01Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IUniswapV2Router01Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV2Router01Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IUniswapV2Router01Session struct {
	Contract     *IUniswapV2Router01 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IUniswapV2Router01CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IUniswapV2Router01CallerSession struct {
	Contract *IUniswapV2Router01Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IUniswapV2Router01TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IUniswapV2Router01TransactorSession struct {
	Contract     *IUniswapV2Router01Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IUniswapV2Router01Raw is an auto generated low-level Go binding around an Ethereum contract.
type IUniswapV2Router01Raw struct {
	Contract *IUniswapV2Router01 // Generic contract binding to access the raw methods on
}

// IUniswapV2Router01CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IUniswapV2Router01CallerRaw struct {
	Contract *IUniswapV2Router01Caller // Generic read-only contract binding to access the raw methods on
}

// IUniswapV2Router01TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IUniswapV2Router01TransactorRaw struct {
	Contract *IUniswapV2Router01Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIUniswapV2Router01 creates a new instance of IUniswapV2Router01, bound to a specific deployed contract.
func NewIUniswapV2Router01(address common.Address, backend bind.ContractBackend) (*IUniswapV2Router01, error) {
	contract, err := bindIUniswapV2Router01(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IUniswapV2Router01{IUniswapV2Router01Caller: IUniswapV2Router01Caller{contract: contract}, IUniswapV2Router01Transactor: IUniswapV2Router01Transactor{contract: contract}, IUniswapV2Router01Filterer: IUniswapV2Router01Filterer{contract: contract}}, nil
}

// NewIUniswapV2Router01Caller creates a new read-only instance of IUniswapV2Router01, bound to a specific deployed contract.
func NewIUniswapV2Router01Caller(address common.Address, caller bind.ContractCaller) (*IUniswapV2Router01Caller, error) {
	contract, err := bindIUniswapV2Router01(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapV2Router01Caller{contract: contract}, nil
}

// NewIUniswapV2Router01Transactor creates a new write-only instance of IUniswapV2Router01, bound to a specific deployed contract.
func NewIUniswapV2Router01Transactor(address common.Address, transactor bind.ContractTransactor) (*IUniswapV2Router01Transactor, error) {
	contract, err := bindIUniswapV2Router01(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapV2Router01Transactor{contract: contract}, nil
}

// NewIUniswapV2Router01Filterer creates a new log filterer instance of IUniswapV2Router01, bound to a specific deployed contract.
func NewIUniswapV2Router01Filterer(address common.Address, filterer bind.ContractFilterer) (*IUniswapV2Router01Filterer, error) {
	contract, err := bindIUniswapV2Router01(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IUniswapV2Router01Filterer{contract: contract}, nil
}

// bindIUniswapV2Router01 binds a generic wrapper to an already deployed contract.
func bindIUniswapV2Router01(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IUniswapV2Router01ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapV2Router01 *IUniswapV2Router01Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapV2Router01.Contract.IUniswapV2Router01Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapV2Router01 *IUniswapV2Router01Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapV2Router01.Contract.IUniswapV2Router01Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapV2Router01 *IUniswapV2Router01Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapV2Router01.Contract.IUniswapV2Router01Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapV2Router01 *IUniswapV2Router01CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapV2Router01.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapV2Router01 *IUniswapV2Router01TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapV2Router01.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapV2Router01 *IUniswapV2Router01TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapV2Router01.Contract.contract.Transact(opts, method, params...)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() pure returns(address)
func (_IUniswapV2Router01 *IUniswapV2Router01Caller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IUniswapV2Router01.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() pure returns(address)
func (_IUniswapV2Router01 *IUniswapV2Router01Session) WETH() (common.Address, error) {
	return _IUniswapV2Router01.Contract.WETH(&_IUniswapV2Router01.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() pure returns(address)
func (_IUniswapV2Router01 *IUniswapV2Router01CallerSession) WETH() (common.Address, error) {
	return _IUniswapV2Router01.Contract.WETH(&_IUniswapV2Router01.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_IUniswapV2Router01 *IUniswapV2Router01Caller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IUniswapV2Router01.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_IUniswapV2Router01 *IUniswapV2Router01Session) Factory() (common.Address, error) {
	return _IUniswapV2Router01.Contract.Factory(&_IUniswapV2Router01.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_IUniswapV2Router01 *IUniswapV2Router01CallerSession) Factory() (common.Address, error) {
	return _IUniswapV2Router01.Contract.Factory(&_IUniswapV2Router01.CallOpts)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_IUniswapV2Router01 *IUniswapV2Router01Caller) GetAmountIn(opts *bind.CallOpts, amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IUniswapV2Router01.contract.Call(opts, &out, "getAmountIn", amountOut, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_IUniswapV2Router01 *IUniswapV2Router01Session) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _IUniswapV2Router01.Contract.GetAmountIn(&_IUniswapV2Router01.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_IUniswapV2Router01 *IUniswapV2Router01CallerSession) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _IUniswapV2Router01.Contract.GetAmountIn(&_IUniswapV2Router01.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_IUniswapV2Router01 *IUniswapV2Router01Caller) GetAmountOut(opts *bind.CallOpts, amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IUniswapV2Router01.contract.Call(opts, &out, "getAmountOut", amountIn, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_IUniswapV2Router01 *IUniswapV2Router01Session) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _IUniswapV2Router01.Contract.GetAmountOut(&_IUniswapV2Router01.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_IUniswapV2Router01 *IUniswapV2Router01CallerSession) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _IUniswapV2Router01.Contract.GetAmountOut(&_IUniswapV2Router01.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_IUniswapV2Router01 *IUniswapV2Router01Caller) GetAmountsIn(opts *bind.CallOpts, amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _IUniswapV2Router01.contract.Call(opts, &out, "getAmountsIn", amountOut, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_IUniswapV2Router01 *IUniswapV2Router01Session) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return _IUniswapV2Router01.Contract.GetAmountsIn(&_IUniswapV2Router01.CallOpts, amountOut, path)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_IUniswapV2Router01 *IUniswapV2Router01CallerSession) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return _IUniswapV2Router01.Contract.GetAmountsIn(&_IUniswapV2Router01.CallOpts, amountOut, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_IUniswapV2Router01 *IUniswapV2Router01Caller) GetAmountsOut(opts *bind.CallOpts, amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _IUniswapV2Router01.contract.Call(opts, &out, "getAmountsOut", amountIn, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_IUniswapV2Router01 *IUniswapV2Router01Session) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _IUniswapV2Router01.Contract.GetAmountsOut(&_IUniswapV2Router01.CallOpts, amountIn, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_IUniswapV2Router01 *IUniswapV2Router01CallerSession) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _IUniswapV2Router01.Contract.GetAmountsOut(&_IUniswapV2Router01.CallOpts, amountIn, path)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_IUniswapV2Router01 *IUniswapV2Router01Caller) Quote(opts *bind.CallOpts, amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IUniswapV2Router01.contract.Call(opts, &out, "quote", amountA, reserveA, reserveB)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_IUniswapV2Router01 *IUniswapV2Router01Session) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _IUniswapV2Router01.Contract.Quote(&_IUniswapV2Router01.CallOpts, amountA, reserveA, reserveB)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_IUniswapV2Router01 *IUniswapV2Router01CallerSession) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _IUniswapV2Router01.Contract.Quote(&_IUniswapV2Router01.CallOpts, amountA, reserveA, reserveB)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_IUniswapV2Router01 *IUniswapV2Router01Transactor) AddLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router01.contract.Transact(opts, "addLiquidity", tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_IUniswapV2Router01 *IUniswapV2Router01Session) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router01.Contract.AddLiquidity(&_IUniswapV2Router01.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_IUniswapV2Router01 *IUniswapV2Router01TransactorSession) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router01.Contract.AddLiquidity(&_IUniswapV2Router01.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xf305d719.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_IUniswapV2Router01 *IUniswapV2Router01Transactor) AddLiquidityETH(opts *bind.TransactOpts, token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router01.contract.Transact(opts, "addLiquidityETH", token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xf305d719.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_IUniswapV2Router01 *IUniswapV2Router01Session) AddLiquidityETH(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router01.Contract.AddLiquidityETH(&_IUniswapV2Router01.TransactOpts, token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xf305d719.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_IUniswapV2Router01 *IUniswapV2Router01TransactorSession) AddLiquidityETH(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router01.Contract.AddLiquidityETH(&_IUniswapV2Router01.TransactOpts, token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_IUniswapV2Router01 *IUniswapV2Router01Transactor) RemoveLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router01.contract.Transact(opts, "removeLiquidity", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_IUniswapV2Router01 *IUniswapV2Router01Session) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router01.Contract.RemoveLiquidity(&_IUniswapV2Router01.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_IUniswapV2Router01 *IUniswapV2Router01TransactorSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router01.Contract.RemoveLiquidity(&_IUniswapV2Router01.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0x02751cec.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_IUniswapV2Router01 *IUniswapV2Router01Transactor) RemoveLiquidityETH(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router01.contract.Transact(opts, "removeLiquidityETH", token, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0x02751cec.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_IUniswapV2Router01 *IUniswapV2Router01Session) RemoveLiquidityETH(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router01.Contract.RemoveLiquidityETH(&_IUniswapV2Router01.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0x02751cec.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_IUniswapV2Router01 *IUniswapV2Router01TransactorSession) RemoveLiquidityETH(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router01.Contract.RemoveLiquidityETH(&_IUniswapV2Router01.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETHWithPermit is a paid mutator transaction binding the contract method 0xded9382a.
//
// Solidity: function removeLiquidityETHWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountETH)
func (_IUniswapV2Router01 *IUniswapV2Router01Transactor) RemoveLiquidityETHWithPermit(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IUniswapV2Router01.contract.Transact(opts, "removeLiquidityETHWithPermit", token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityETHWithPermit is a paid mutator transaction binding the contract method 0xded9382a.
//
// Solidity: function removeLiquidityETHWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountETH)
func (_IUniswapV2Router01 *IUniswapV2Router01Session) RemoveLiquidityETHWithPermit(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IUniswapV2Router01.Contract.RemoveLiquidityETHWithPermit(&_IUniswapV2Router01.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityETHWithPermit is a paid mutator transaction binding the contract method 0xded9382a.
//
// Solidity: function removeLiquidityETHWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountETH)
func (_IUniswapV2Router01 *IUniswapV2Router01TransactorSession) RemoveLiquidityETHWithPermit(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IUniswapV2Router01.Contract.RemoveLiquidityETHWithPermit(&_IUniswapV2Router01.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_IUniswapV2Router01 *IUniswapV2Router01Transactor) RemoveLiquidityWithPermit(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IUniswapV2Router01.contract.Transact(opts, "removeLiquidityWithPermit", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_IUniswapV2Router01 *IUniswapV2Router01Session) RemoveLiquidityWithPermit(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IUniswapV2Router01.Contract.RemoveLiquidityWithPermit(&_IUniswapV2Router01.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_IUniswapV2Router01 *IUniswapV2Router01TransactorSession) RemoveLiquidityWithPermit(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IUniswapV2Router01.Contract.RemoveLiquidityWithPermit(&_IUniswapV2Router01.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0xfb3bdb41.
//
// Solidity: function swapETHForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IUniswapV2Router01 *IUniswapV2Router01Transactor) SwapETHForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router01.contract.Transact(opts, "swapETHForExactTokens", amountOut, path, to, deadline)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0xfb3bdb41.
//
// Solidity: function swapETHForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IUniswapV2Router01 *IUniswapV2Router01Session) SwapETHForExactTokens(amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router01.Contract.SwapETHForExactTokens(&_IUniswapV2Router01.TransactOpts, amountOut, path, to, deadline)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0xfb3bdb41.
//
// Solidity: function swapETHForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IUniswapV2Router01 *IUniswapV2Router01TransactorSession) SwapETHForExactTokens(amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router01.Contract.SwapETHForExactTokens(&_IUniswapV2Router01.TransactOpts, amountOut, path, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IUniswapV2Router01 *IUniswapV2Router01Transactor) SwapExactETHForTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router01.contract.Transact(opts, "swapExactETHForTokens", amountOutMin, path, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IUniswapV2Router01 *IUniswapV2Router01Session) SwapExactETHForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router01.Contract.SwapExactETHForTokens(&_IUniswapV2Router01.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IUniswapV2Router01 *IUniswapV2Router01TransactorSession) SwapExactETHForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router01.Contract.SwapExactETHForTokens(&_IUniswapV2Router01.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IUniswapV2Router01 *IUniswapV2Router01Transactor) SwapExactTokensForETH(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router01.contract.Transact(opts, "swapExactTokensForETH", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IUniswapV2Router01 *IUniswapV2Router01Session) SwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router01.Contract.SwapExactTokensForETH(&_IUniswapV2Router01.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IUniswapV2Router01 *IUniswapV2Router01TransactorSession) SwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router01.Contract.SwapExactTokensForETH(&_IUniswapV2Router01.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IUniswapV2Router01 *IUniswapV2Router01Transactor) SwapExactTokensForTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router01.contract.Transact(opts, "swapExactTokensForTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IUniswapV2Router01 *IUniswapV2Router01Session) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router01.Contract.SwapExactTokensForTokens(&_IUniswapV2Router01.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IUniswapV2Router01 *IUniswapV2Router01TransactorSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router01.Contract.SwapExactTokensForTokens(&_IUniswapV2Router01.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0x4a25d94a.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IUniswapV2Router01 *IUniswapV2Router01Transactor) SwapTokensForExactETH(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router01.contract.Transact(opts, "swapTokensForExactETH", amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0x4a25d94a.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IUniswapV2Router01 *IUniswapV2Router01Session) SwapTokensForExactETH(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router01.Contract.SwapTokensForExactETH(&_IUniswapV2Router01.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0x4a25d94a.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IUniswapV2Router01 *IUniswapV2Router01TransactorSession) SwapTokensForExactETH(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router01.Contract.SwapTokensForExactETH(&_IUniswapV2Router01.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IUniswapV2Router01 *IUniswapV2Router01Transactor) SwapTokensForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router01.contract.Transact(opts, "swapTokensForExactTokens", amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IUniswapV2Router01 *IUniswapV2Router01Session) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router01.Contract.SwapTokensForExactTokens(&_IUniswapV2Router01.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IUniswapV2Router01 *IUniswapV2Router01TransactorSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router01.Contract.SwapTokensForExactTokens(&_IUniswapV2Router01.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// IUniswapV2Router02MetaData contains all meta data concerning the IUniswapV2Router02 contract.
var IUniswapV2Router02MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountADesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsIn\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityETHSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityETHWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityETHWithPermitSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapETHForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactETHForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactETHForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForETHSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"ad5c4648": "WETH()",
		"e8e33700": "addLiquidity(address,address,uint256,uint256,uint256,uint256,address,uint256)",
		"f305d719": "addLiquidityETH(address,uint256,uint256,uint256,address,uint256)",
		"c45a0155": "factory()",
		"85f8c259": "getAmountIn(uint256,uint256,uint256)",
		"054d50d4": "getAmountOut(uint256,uint256,uint256)",
		"1f00ca74": "getAmountsIn(uint256,address[])",
		"d06ca61f": "getAmountsOut(uint256,address[])",
		"ad615dec": "quote(uint256,uint256,uint256)",
		"baa2abde": "removeLiquidity(address,address,uint256,uint256,uint256,address,uint256)",
		"02751cec": "removeLiquidityETH(address,uint256,uint256,uint256,address,uint256)",
		"af2979eb": "removeLiquidityETHSupportingFeeOnTransferTokens(address,uint256,uint256,uint256,address,uint256)",
		"ded9382a": "removeLiquidityETHWithPermit(address,uint256,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32)",
		"5b0d5984": "removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(address,uint256,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32)",
		"2195995c": "removeLiquidityWithPermit(address,address,uint256,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32)",
		"fb3bdb41": "swapETHForExactTokens(uint256,address[],address,uint256)",
		"7ff36ab5": "swapExactETHForTokens(uint256,address[],address,uint256)",
		"b6f9de95": "swapExactETHForTokensSupportingFeeOnTransferTokens(uint256,address[],address,uint256)",
		"18cbafe5": "swapExactTokensForETH(uint256,uint256,address[],address,uint256)",
		"791ac947": "swapExactTokensForETHSupportingFeeOnTransferTokens(uint256,uint256,address[],address,uint256)",
		"38ed1739": "swapExactTokensForTokens(uint256,uint256,address[],address,uint256)",
		"5c11d795": "swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256,uint256,address[],address,uint256)",
		"4a25d94a": "swapTokensForExactETH(uint256,uint256,address[],address,uint256)",
		"8803dbee": "swapTokensForExactTokens(uint256,uint256,address[],address,uint256)",
	},
}

// IUniswapV2Router02ABI is the input ABI used to generate the binding from.
// Deprecated: Use IUniswapV2Router02MetaData.ABI instead.
var IUniswapV2Router02ABI = IUniswapV2Router02MetaData.ABI

// Deprecated: Use IUniswapV2Router02MetaData.Sigs instead.
// IUniswapV2Router02FuncSigs maps the 4-byte function signature to its string representation.
var IUniswapV2Router02FuncSigs = IUniswapV2Router02MetaData.Sigs

// IUniswapV2Router02 is an auto generated Go binding around an Ethereum contract.
type IUniswapV2Router02 struct {
	IUniswapV2Router02Caller     // Read-only binding to the contract
	IUniswapV2Router02Transactor // Write-only binding to the contract
	IUniswapV2Router02Filterer   // Log filterer for contract events
}

// IUniswapV2Router02Caller is an auto generated read-only Go binding around an Ethereum contract.
type IUniswapV2Router02Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV2Router02Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IUniswapV2Router02Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV2Router02Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IUniswapV2Router02Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV2Router02Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IUniswapV2Router02Session struct {
	Contract     *IUniswapV2Router02 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IUniswapV2Router02CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IUniswapV2Router02CallerSession struct {
	Contract *IUniswapV2Router02Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IUniswapV2Router02TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IUniswapV2Router02TransactorSession struct {
	Contract     *IUniswapV2Router02Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IUniswapV2Router02Raw is an auto generated low-level Go binding around an Ethereum contract.
type IUniswapV2Router02Raw struct {
	Contract *IUniswapV2Router02 // Generic contract binding to access the raw methods on
}

// IUniswapV2Router02CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IUniswapV2Router02CallerRaw struct {
	Contract *IUniswapV2Router02Caller // Generic read-only contract binding to access the raw methods on
}

// IUniswapV2Router02TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IUniswapV2Router02TransactorRaw struct {
	Contract *IUniswapV2Router02Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIUniswapV2Router02 creates a new instance of IUniswapV2Router02, bound to a specific deployed contract.
func NewIUniswapV2Router02(address common.Address, backend bind.ContractBackend) (*IUniswapV2Router02, error) {
	contract, err := bindIUniswapV2Router02(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IUniswapV2Router02{IUniswapV2Router02Caller: IUniswapV2Router02Caller{contract: contract}, IUniswapV2Router02Transactor: IUniswapV2Router02Transactor{contract: contract}, IUniswapV2Router02Filterer: IUniswapV2Router02Filterer{contract: contract}}, nil
}

// NewIUniswapV2Router02Caller creates a new read-only instance of IUniswapV2Router02, bound to a specific deployed contract.
func NewIUniswapV2Router02Caller(address common.Address, caller bind.ContractCaller) (*IUniswapV2Router02Caller, error) {
	contract, err := bindIUniswapV2Router02(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapV2Router02Caller{contract: contract}, nil
}

// NewIUniswapV2Router02Transactor creates a new write-only instance of IUniswapV2Router02, bound to a specific deployed contract.
func NewIUniswapV2Router02Transactor(address common.Address, transactor bind.ContractTransactor) (*IUniswapV2Router02Transactor, error) {
	contract, err := bindIUniswapV2Router02(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapV2Router02Transactor{contract: contract}, nil
}

// NewIUniswapV2Router02Filterer creates a new log filterer instance of IUniswapV2Router02, bound to a specific deployed contract.
func NewIUniswapV2Router02Filterer(address common.Address, filterer bind.ContractFilterer) (*IUniswapV2Router02Filterer, error) {
	contract, err := bindIUniswapV2Router02(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IUniswapV2Router02Filterer{contract: contract}, nil
}

// bindIUniswapV2Router02 binds a generic wrapper to an already deployed contract.
func bindIUniswapV2Router02(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IUniswapV2Router02ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapV2Router02 *IUniswapV2Router02Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapV2Router02.Contract.IUniswapV2Router02Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapV2Router02 *IUniswapV2Router02Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapV2Router02.Contract.IUniswapV2Router02Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapV2Router02 *IUniswapV2Router02Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapV2Router02.Contract.IUniswapV2Router02Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapV2Router02 *IUniswapV2Router02CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapV2Router02.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapV2Router02 *IUniswapV2Router02TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapV2Router02.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapV2Router02 *IUniswapV2Router02TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapV2Router02.Contract.contract.Transact(opts, method, params...)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() pure returns(address)
func (_IUniswapV2Router02 *IUniswapV2Router02Caller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IUniswapV2Router02.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() pure returns(address)
func (_IUniswapV2Router02 *IUniswapV2Router02Session) WETH() (common.Address, error) {
	return _IUniswapV2Router02.Contract.WETH(&_IUniswapV2Router02.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() pure returns(address)
func (_IUniswapV2Router02 *IUniswapV2Router02CallerSession) WETH() (common.Address, error) {
	return _IUniswapV2Router02.Contract.WETH(&_IUniswapV2Router02.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_IUniswapV2Router02 *IUniswapV2Router02Caller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IUniswapV2Router02.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_IUniswapV2Router02 *IUniswapV2Router02Session) Factory() (common.Address, error) {
	return _IUniswapV2Router02.Contract.Factory(&_IUniswapV2Router02.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_IUniswapV2Router02 *IUniswapV2Router02CallerSession) Factory() (common.Address, error) {
	return _IUniswapV2Router02.Contract.Factory(&_IUniswapV2Router02.CallOpts)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_IUniswapV2Router02 *IUniswapV2Router02Caller) GetAmountIn(opts *bind.CallOpts, amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IUniswapV2Router02.contract.Call(opts, &out, "getAmountIn", amountOut, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_IUniswapV2Router02 *IUniswapV2Router02Session) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _IUniswapV2Router02.Contract.GetAmountIn(&_IUniswapV2Router02.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_IUniswapV2Router02 *IUniswapV2Router02CallerSession) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _IUniswapV2Router02.Contract.GetAmountIn(&_IUniswapV2Router02.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_IUniswapV2Router02 *IUniswapV2Router02Caller) GetAmountOut(opts *bind.CallOpts, amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IUniswapV2Router02.contract.Call(opts, &out, "getAmountOut", amountIn, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_IUniswapV2Router02 *IUniswapV2Router02Session) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _IUniswapV2Router02.Contract.GetAmountOut(&_IUniswapV2Router02.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_IUniswapV2Router02 *IUniswapV2Router02CallerSession) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _IUniswapV2Router02.Contract.GetAmountOut(&_IUniswapV2Router02.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_IUniswapV2Router02 *IUniswapV2Router02Caller) GetAmountsIn(opts *bind.CallOpts, amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _IUniswapV2Router02.contract.Call(opts, &out, "getAmountsIn", amountOut, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_IUniswapV2Router02 *IUniswapV2Router02Session) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return _IUniswapV2Router02.Contract.GetAmountsIn(&_IUniswapV2Router02.CallOpts, amountOut, path)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_IUniswapV2Router02 *IUniswapV2Router02CallerSession) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return _IUniswapV2Router02.Contract.GetAmountsIn(&_IUniswapV2Router02.CallOpts, amountOut, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_IUniswapV2Router02 *IUniswapV2Router02Caller) GetAmountsOut(opts *bind.CallOpts, amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _IUniswapV2Router02.contract.Call(opts, &out, "getAmountsOut", amountIn, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_IUniswapV2Router02 *IUniswapV2Router02Session) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _IUniswapV2Router02.Contract.GetAmountsOut(&_IUniswapV2Router02.CallOpts, amountIn, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_IUniswapV2Router02 *IUniswapV2Router02CallerSession) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _IUniswapV2Router02.Contract.GetAmountsOut(&_IUniswapV2Router02.CallOpts, amountIn, path)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_IUniswapV2Router02 *IUniswapV2Router02Caller) Quote(opts *bind.CallOpts, amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IUniswapV2Router02.contract.Call(opts, &out, "quote", amountA, reserveA, reserveB)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_IUniswapV2Router02 *IUniswapV2Router02Session) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _IUniswapV2Router02.Contract.Quote(&_IUniswapV2Router02.CallOpts, amountA, reserveA, reserveB)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_IUniswapV2Router02 *IUniswapV2Router02CallerSession) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _IUniswapV2Router02.Contract.Quote(&_IUniswapV2Router02.CallOpts, amountA, reserveA, reserveB)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_IUniswapV2Router02 *IUniswapV2Router02Transactor) AddLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router02.contract.Transact(opts, "addLiquidity", tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_IUniswapV2Router02 *IUniswapV2Router02Session) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router02.Contract.AddLiquidity(&_IUniswapV2Router02.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_IUniswapV2Router02 *IUniswapV2Router02TransactorSession) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router02.Contract.AddLiquidity(&_IUniswapV2Router02.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xf305d719.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_IUniswapV2Router02 *IUniswapV2Router02Transactor) AddLiquidityETH(opts *bind.TransactOpts, token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router02.contract.Transact(opts, "addLiquidityETH", token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xf305d719.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_IUniswapV2Router02 *IUniswapV2Router02Session) AddLiquidityETH(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router02.Contract.AddLiquidityETH(&_IUniswapV2Router02.TransactOpts, token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xf305d719.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_IUniswapV2Router02 *IUniswapV2Router02TransactorSession) AddLiquidityETH(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router02.Contract.AddLiquidityETH(&_IUniswapV2Router02.TransactOpts, token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_IUniswapV2Router02 *IUniswapV2Router02Transactor) RemoveLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router02.contract.Transact(opts, "removeLiquidity", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_IUniswapV2Router02 *IUniswapV2Router02Session) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router02.Contract.RemoveLiquidity(&_IUniswapV2Router02.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_IUniswapV2Router02 *IUniswapV2Router02TransactorSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router02.Contract.RemoveLiquidity(&_IUniswapV2Router02.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0x02751cec.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_IUniswapV2Router02 *IUniswapV2Router02Transactor) RemoveLiquidityETH(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router02.contract.Transact(opts, "removeLiquidityETH", token, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0x02751cec.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_IUniswapV2Router02 *IUniswapV2Router02Session) RemoveLiquidityETH(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router02.Contract.RemoveLiquidityETH(&_IUniswapV2Router02.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0x02751cec.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_IUniswapV2Router02 *IUniswapV2Router02TransactorSession) RemoveLiquidityETH(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router02.Contract.RemoveLiquidityETH(&_IUniswapV2Router02.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xaf2979eb.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountETH)
func (_IUniswapV2Router02 *IUniswapV2Router02Transactor) RemoveLiquidityETHSupportingFeeOnTransferTokens(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router02.contract.Transact(opts, "removeLiquidityETHSupportingFeeOnTransferTokens", token, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xaf2979eb.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountETH)
func (_IUniswapV2Router02 *IUniswapV2Router02Session) RemoveLiquidityETHSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router02.Contract.RemoveLiquidityETHSupportingFeeOnTransferTokens(&_IUniswapV2Router02.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xaf2979eb.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountETH)
func (_IUniswapV2Router02 *IUniswapV2Router02TransactorSession) RemoveLiquidityETHSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router02.Contract.RemoveLiquidityETHSupportingFeeOnTransferTokens(&_IUniswapV2Router02.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETHWithPermit is a paid mutator transaction binding the contract method 0xded9382a.
//
// Solidity: function removeLiquidityETHWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountETH)
func (_IUniswapV2Router02 *IUniswapV2Router02Transactor) RemoveLiquidityETHWithPermit(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IUniswapV2Router02.contract.Transact(opts, "removeLiquidityETHWithPermit", token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityETHWithPermit is a paid mutator transaction binding the contract method 0xded9382a.
//
// Solidity: function removeLiquidityETHWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountETH)
func (_IUniswapV2Router02 *IUniswapV2Router02Session) RemoveLiquidityETHWithPermit(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IUniswapV2Router02.Contract.RemoveLiquidityETHWithPermit(&_IUniswapV2Router02.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityETHWithPermit is a paid mutator transaction binding the contract method 0xded9382a.
//
// Solidity: function removeLiquidityETHWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountETH)
func (_IUniswapV2Router02 *IUniswapV2Router02TransactorSession) RemoveLiquidityETHWithPermit(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IUniswapV2Router02.Contract.RemoveLiquidityETHWithPermit(&_IUniswapV2Router02.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5b0d5984.
//
// Solidity: function removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountETH)
func (_IUniswapV2Router02 *IUniswapV2Router02Transactor) RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IUniswapV2Router02.contract.Transact(opts, "removeLiquidityETHWithPermitSupportingFeeOnTransferTokens", token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5b0d5984.
//
// Solidity: function removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountETH)
func (_IUniswapV2Router02 *IUniswapV2Router02Session) RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IUniswapV2Router02.Contract.RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(&_IUniswapV2Router02.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5b0d5984.
//
// Solidity: function removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountETH)
func (_IUniswapV2Router02 *IUniswapV2Router02TransactorSession) RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IUniswapV2Router02.Contract.RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(&_IUniswapV2Router02.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_IUniswapV2Router02 *IUniswapV2Router02Transactor) RemoveLiquidityWithPermit(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IUniswapV2Router02.contract.Transact(opts, "removeLiquidityWithPermit", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_IUniswapV2Router02 *IUniswapV2Router02Session) RemoveLiquidityWithPermit(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IUniswapV2Router02.Contract.RemoveLiquidityWithPermit(&_IUniswapV2Router02.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_IUniswapV2Router02 *IUniswapV2Router02TransactorSession) RemoveLiquidityWithPermit(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IUniswapV2Router02.Contract.RemoveLiquidityWithPermit(&_IUniswapV2Router02.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0xfb3bdb41.
//
// Solidity: function swapETHForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IUniswapV2Router02 *IUniswapV2Router02Transactor) SwapETHForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router02.contract.Transact(opts, "swapETHForExactTokens", amountOut, path, to, deadline)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0xfb3bdb41.
//
// Solidity: function swapETHForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IUniswapV2Router02 *IUniswapV2Router02Session) SwapETHForExactTokens(amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router02.Contract.SwapETHForExactTokens(&_IUniswapV2Router02.TransactOpts, amountOut, path, to, deadline)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0xfb3bdb41.
//
// Solidity: function swapETHForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IUniswapV2Router02 *IUniswapV2Router02TransactorSession) SwapETHForExactTokens(amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router02.Contract.SwapETHForExactTokens(&_IUniswapV2Router02.TransactOpts, amountOut, path, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IUniswapV2Router02 *IUniswapV2Router02Transactor) SwapExactETHForTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router02.contract.Transact(opts, "swapExactETHForTokens", amountOutMin, path, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IUniswapV2Router02 *IUniswapV2Router02Session) SwapExactETHForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router02.Contract.SwapExactETHForTokens(&_IUniswapV2Router02.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_IUniswapV2Router02 *IUniswapV2Router02TransactorSession) SwapExactETHForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router02.Contract.SwapExactETHForTokens(&_IUniswapV2Router02.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb6f9de95.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_IUniswapV2Router02 *IUniswapV2Router02Transactor) SwapExactETHForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router02.contract.Transact(opts, "swapExactETHForTokensSupportingFeeOnTransferTokens", amountOutMin, path, to, deadline)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb6f9de95.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_IUniswapV2Router02 *IUniswapV2Router02Session) SwapExactETHForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router02.Contract.SwapExactETHForTokensSupportingFeeOnTransferTokens(&_IUniswapV2Router02.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb6f9de95.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_IUniswapV2Router02 *IUniswapV2Router02TransactorSession) SwapExactETHForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router02.Contract.SwapExactETHForTokensSupportingFeeOnTransferTokens(&_IUniswapV2Router02.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IUniswapV2Router02 *IUniswapV2Router02Transactor) SwapExactTokensForETH(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router02.contract.Transact(opts, "swapExactTokensForETH", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IUniswapV2Router02 *IUniswapV2Router02Session) SwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router02.Contract.SwapExactTokensForETH(&_IUniswapV2Router02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IUniswapV2Router02 *IUniswapV2Router02TransactorSession) SwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router02.Contract.SwapExactTokensForETH(&_IUniswapV2Router02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x791ac947.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_IUniswapV2Router02 *IUniswapV2Router02Transactor) SwapExactTokensForETHSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router02.contract.Transact(opts, "swapExactTokensForETHSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x791ac947.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_IUniswapV2Router02 *IUniswapV2Router02Session) SwapExactTokensForETHSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router02.Contract.SwapExactTokensForETHSupportingFeeOnTransferTokens(&_IUniswapV2Router02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x791ac947.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_IUniswapV2Router02 *IUniswapV2Router02TransactorSession) SwapExactTokensForETHSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router02.Contract.SwapExactTokensForETHSupportingFeeOnTransferTokens(&_IUniswapV2Router02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IUniswapV2Router02 *IUniswapV2Router02Transactor) SwapExactTokensForTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router02.contract.Transact(opts, "swapExactTokensForTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IUniswapV2Router02 *IUniswapV2Router02Session) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router02.Contract.SwapExactTokensForTokens(&_IUniswapV2Router02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IUniswapV2Router02 *IUniswapV2Router02TransactorSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router02.Contract.SwapExactTokensForTokens(&_IUniswapV2Router02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_IUniswapV2Router02 *IUniswapV2Router02Transactor) SwapExactTokensForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router02.contract.Transact(opts, "swapExactTokensForTokensSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_IUniswapV2Router02 *IUniswapV2Router02Session) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router02.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_IUniswapV2Router02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_IUniswapV2Router02 *IUniswapV2Router02TransactorSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router02.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_IUniswapV2Router02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0x4a25d94a.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IUniswapV2Router02 *IUniswapV2Router02Transactor) SwapTokensForExactETH(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router02.contract.Transact(opts, "swapTokensForExactETH", amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0x4a25d94a.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IUniswapV2Router02 *IUniswapV2Router02Session) SwapTokensForExactETH(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router02.Contract.SwapTokensForExactETH(&_IUniswapV2Router02.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0x4a25d94a.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IUniswapV2Router02 *IUniswapV2Router02TransactorSession) SwapTokensForExactETH(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router02.Contract.SwapTokensForExactETH(&_IUniswapV2Router02.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IUniswapV2Router02 *IUniswapV2Router02Transactor) SwapTokensForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router02.contract.Transact(opts, "swapTokensForExactTokens", amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IUniswapV2Router02 *IUniswapV2Router02Session) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router02.Contract.SwapTokensForExactTokens(&_IUniswapV2Router02.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IUniswapV2Router02 *IUniswapV2Router02TransactorSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Router02.Contract.SwapTokensForExactTokens(&_IUniswapV2Router02.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// IWethMetaData contains all meta data concerning the IWeth contract.
var IWethMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"2e1a7d4d": "withdraw(uint256)",
	},
}

// IWethABI is the input ABI used to generate the binding from.
// Deprecated: Use IWethMetaData.ABI instead.
var IWethABI = IWethMetaData.ABI

// Deprecated: Use IWethMetaData.Sigs instead.
// IWethFuncSigs maps the 4-byte function signature to its string representation.
var IWethFuncSigs = IWethMetaData.Sigs

// IWeth is an auto generated Go binding around an Ethereum contract.
type IWeth struct {
	IWethCaller     // Read-only binding to the contract
	IWethTransactor // Write-only binding to the contract
	IWethFilterer   // Log filterer for contract events
}

// IWethCaller is an auto generated read-only Go binding around an Ethereum contract.
type IWethCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWethTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IWethTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWethFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IWethFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWethSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IWethSession struct {
	Contract     *IWeth            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWethCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IWethCallerSession struct {
	Contract *IWethCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IWethTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IWethTransactorSession struct {
	Contract     *IWethTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWethRaw is an auto generated low-level Go binding around an Ethereum contract.
type IWethRaw struct {
	Contract *IWeth // Generic contract binding to access the raw methods on
}

// IWethCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IWethCallerRaw struct {
	Contract *IWethCaller // Generic read-only contract binding to access the raw methods on
}

// IWethTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IWethTransactorRaw struct {
	Contract *IWethTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIWeth creates a new instance of IWeth, bound to a specific deployed contract.
func NewIWeth(address common.Address, backend bind.ContractBackend) (*IWeth, error) {
	contract, err := bindIWeth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IWeth{IWethCaller: IWethCaller{contract: contract}, IWethTransactor: IWethTransactor{contract: contract}, IWethFilterer: IWethFilterer{contract: contract}}, nil
}

// NewIWethCaller creates a new read-only instance of IWeth, bound to a specific deployed contract.
func NewIWethCaller(address common.Address, caller bind.ContractCaller) (*IWethCaller, error) {
	contract, err := bindIWeth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IWethCaller{contract: contract}, nil
}

// NewIWethTransactor creates a new write-only instance of IWeth, bound to a specific deployed contract.
func NewIWethTransactor(address common.Address, transactor bind.ContractTransactor) (*IWethTransactor, error) {
	contract, err := bindIWeth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IWethTransactor{contract: contract}, nil
}

// NewIWethFilterer creates a new log filterer instance of IWeth, bound to a specific deployed contract.
func NewIWethFilterer(address common.Address, filterer bind.ContractFilterer) (*IWethFilterer, error) {
	contract, err := bindIWeth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IWethFilterer{contract: contract}, nil
}

// bindIWeth binds a generic wrapper to an already deployed contract.
func bindIWeth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IWethABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWeth *IWethRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWeth.Contract.IWethCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWeth *IWethRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWeth.Contract.IWethTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWeth *IWethRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWeth.Contract.IWethTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWeth *IWethCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWeth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWeth *IWethTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWeth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWeth *IWethTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWeth.Contract.contract.Transact(opts, method, params...)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_IWeth *IWethTransactor) Withdraw(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _IWeth.contract.Transact(opts, "withdraw", wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_IWeth *IWethSession) Withdraw(wad *big.Int) (*types.Transaction, error) {
	return _IWeth.Contract.Withdraw(&_IWeth.TransactOpts, wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_IWeth *IWethTransactorSession) Withdraw(wad *big.Int) (*types.Transaction, error) {
	return _IWeth.Contract.Withdraw(&_IWeth.TransactOpts, wad)
}

// MathMetaData contains all meta data concerning the Math contract.
var MathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220bb99db4d47a714da9d030ee870f0bfe17aa178d941c2cae75fd7424b246169e264736f6c634300060c0033",
}

// MathABI is the input ABI used to generate the binding from.
// Deprecated: Use MathMetaData.ABI instead.
var MathABI = MathMetaData.ABI

// MathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MathMetaData.Bin instead.
var MathBin = MathMetaData.Bin

// DeployMath deploys a new Ethereum contract, binding an instance of Math to it.
func DeployMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Math, error) {
	parsed, err := MathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// Math is an auto generated Go binding around an Ethereum contract.
type Math struct {
	MathCaller     // Read-only binding to the contract
	MathTransactor // Write-only binding to the contract
	MathFilterer   // Log filterer for contract events
}

// MathCaller is an auto generated read-only Go binding around an Ethereum contract.
type MathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MathSession struct {
	Contract     *Math             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MathCallerSession struct {
	Contract *MathCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MathTransactorSession struct {
	Contract     *MathTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathRaw is an auto generated low-level Go binding around an Ethereum contract.
type MathRaw struct {
	Contract *Math // Generic contract binding to access the raw methods on
}

// MathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MathCallerRaw struct {
	Contract *MathCaller // Generic read-only contract binding to access the raw methods on
}

// MathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MathTransactorRaw struct {
	Contract *MathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMath creates a new instance of Math, bound to a specific deployed contract.
func NewMath(address common.Address, backend bind.ContractBackend) (*Math, error) {
	contract, err := bindMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// NewMathCaller creates a new read-only instance of Math, bound to a specific deployed contract.
func NewMathCaller(address common.Address, caller bind.ContractCaller) (*MathCaller, error) {
	contract, err := bindMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MathCaller{contract: contract}, nil
}

// NewMathTransactor creates a new write-only instance of Math, bound to a specific deployed contract.
func NewMathTransactor(address common.Address, transactor bind.ContractTransactor) (*MathTransactor, error) {
	contract, err := bindMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MathTransactor{contract: contract}, nil
}

// NewMathFilterer creates a new log filterer instance of Math, bound to a specific deployed contract.
func NewMathFilterer(address common.Address, filterer bind.ContractFilterer) (*MathFilterer, error) {
	contract, err := bindMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MathFilterer{contract: contract}, nil
}

// bindMath binds a generic wrapper to an already deployed contract.
func bindMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Math.Contract.MathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Math.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.contract.Transact(opts, method, params...)
}

// SafeERC20MetaData contains all meta data concerning the SafeERC20 contract.
var SafeERC20MetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220b606594df0ff8d91edb4764202d3985fa22865008dd1f0266616e73caaf93fef64736f6c634300060c0033",
}

// SafeERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeERC20MetaData.ABI instead.
var SafeERC20ABI = SafeERC20MetaData.ABI

// SafeERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeERC20MetaData.Bin instead.
var SafeERC20Bin = SafeERC20MetaData.Bin

// DeploySafeERC20 deploys a new Ethereum contract, binding an instance of SafeERC20 to it.
func DeploySafeERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeERC20, error) {
	parsed, err := SafeERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// SafeERC20 is an auto generated Go binding around an Ethereum contract.
type SafeERC20 struct {
	SafeERC20Caller     // Read-only binding to the contract
	SafeERC20Transactor // Write-only binding to the contract
	SafeERC20Filterer   // Log filterer for contract events
}

// SafeERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type SafeERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeERC20Session struct {
	Contract     *SafeERC20        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeERC20CallerSession struct {
	Contract *SafeERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SafeERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeERC20TransactorSession struct {
	Contract     *SafeERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SafeERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type SafeERC20Raw struct {
	Contract *SafeERC20 // Generic contract binding to access the raw methods on
}

// SafeERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeERC20CallerRaw struct {
	Contract *SafeERC20Caller // Generic read-only contract binding to access the raw methods on
}

// SafeERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeERC20TransactorRaw struct {
	Contract *SafeERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeERC20 creates a new instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20(address common.Address, backend bind.ContractBackend) (*SafeERC20, error) {
	contract, err := bindSafeERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// NewSafeERC20Caller creates a new read-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Caller(address common.Address, caller bind.ContractCaller) (*SafeERC20Caller, error) {
	contract, err := bindSafeERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Caller{contract: contract}, nil
}

// NewSafeERC20Transactor creates a new write-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*SafeERC20Transactor, error) {
	contract, err := bindSafeERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Transactor{contract: contract}, nil
}

// NewSafeERC20Filterer creates a new log filterer instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*SafeERC20Filterer, error) {
	contract, err := bindSafeERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Filterer{contract: contract}, nil
}

// bindSafeERC20 binds a generic wrapper to an already deployed contract.
func bindSafeERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.SafeERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transact(opts, method, params...)
}

// SafeMathMetaData contains all meta data concerning the SafeMath contract.
var SafeMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220aefe8a3dca908e5c24f7f459c229f95186a0ebc0d21df22fa19e9e688be2118c64736f6c634300060c0033",
}

// SafeMathABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeMathMetaData.ABI instead.
var SafeMathABI = SafeMathMetaData.ABI

// SafeMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeMathMetaData.Bin instead.
var SafeMathBin = SafeMathMetaData.Bin

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := SafeMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// StrategyBaseMetaData contains all meta data concerning the StrategyBase contract.
var StrategyBaseMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyExitEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"profit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loss\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debtPayment\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debtOutstanding\",\"type\":\"uint256\"}],\"name\":\"Harvested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debtThreshold\",\"type\":\"uint256\"}],\"name\":\"UpdatedDebtThreshold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newKeeper\",\"type\":\"address\"}],\"name\":\"UpdatedKeeper\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"profitFactor\",\"type\":\"uint256\"}],\"name\":\"UpdatedProfitFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"UpdatedReportDelay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rewards\",\"type\":\"address\"}],\"name\":\"UpdatedRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newStrategist\",\"type\":\"address\"}],\"name\":\"UpdatedStrategist\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"apiVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balanceOfWant\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimRewards\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimableBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"debtThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegatedAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"doHealthCheck\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyExit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"estimatedTotalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"harvest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"harvestProfitMax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"harvestProfitMin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"callCost\",\"type\":\"uint256\"}],\"name\":\"harvestTrigger\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"healthCheck\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"keepCRV\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"keepCrvPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"keepCVX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"keepCVXDestination\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"keeper\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxReportDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newStrategy\",\"type\":\"address\"}],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"profitFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsContract\",\"outputs\":[{\"internalType\":\"contractIConvexRewards\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_claimRewards\",\"type\":\"bool\"}],\"name\":\"setClaimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_debtThreshold\",\"type\":\"uint256\"}],\"name\":\"setDebtThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_doHealthCheck\",\"type\":\"bool\"}],\"name\":\"setDoHealthCheck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setEmergencyExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_forceHarvestTriggerOnce\",\"type\":\"bool\"}],\"name\":\"setForceHarvestTriggerOnce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_healthCheck\",\"type\":\"address\"}],\"name\":\"setHealthCheck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_keepCRV\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_keepCVX\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_keepCVXDestination\",\"type\":\"address\"}],\"name\":\"setKeep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_keeper\",\"type\":\"address\"}],\"name\":\"setKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_delay\",\"type\":\"uint256\"}],\"name\":\"setMaxReportDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_profitFactor\",\"type\":\"uint256\"}],\"name\":\"setProfitFactor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewards\",\"type\":\"address\"}],\"name\":\"setRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategist\",\"type\":\"address\"}],\"name\":\"setStrategist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"strategist\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"sweep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"callCost\",\"type\":\"uint256\"}],\"name\":\"tendTrigger\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"contractVaultAPI\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"virtualRewardsPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"want\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountNeeded\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_loss\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawToConvexDepositTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"25829410": "apiVersion()",
		"c1a3d44c": "balanceOfWant()",
		"372500ab": "claimRewards()",
		"c4f45423": "claimableBalance()",
		"aa5480cf": "creditThreshold()",
		"1d12f28b": "debtThreshold()",
		"8e6350e2": "delegatedAssets()",
		"6718835f": "doHealthCheck()",
		"5641ec03": "emergencyExit()",
		"efbb5cb0": "estimatedTotalAssets()",
		"4641257d": "harvest()",
		"3b7c6e2f": "harvestProfitMax()",
		"5fbeb25f": "harvestProfitMin()",
		"ed882c2b": "harvestTrigger(uint256)",
		"b252720b": "healthCheck()",
		"22f3e2d4": "isActive()",
		"7fef901a": "keepCRV()",
		"4b31217e": "keepCVX()",
		"ef86b23c": "keepCVXDestination()",
		"aced1661": "keeper()",
		"28b7ccf7": "maxReportDelay()",
		"ce5494bb": "migrate(address)",
		"06fdde03": "name()",
		"f1068454": "pid()",
		"8cdfe166": "profitFactor()",
		"9ec5a894": "rewards()",
		"220cce97": "rewardsContract()",
		"a98f9296": "setClaimRewards(bool)",
		"0f969b87": "setDebtThreshold(uint256)",
		"ac00ff26": "setDoHealthCheck(bool)",
		"fcf2d0ad": "setEmergencyExit()",
		"0ada4dab": "setForceHarvestTriggerOnce(bool)",
		"11bc8245": "setHealthCheck(address)",
		"1111fe1c": "setKeep(uint256,uint256,address)",
		"748747e6": "setKeeper(address)",
		"f017c92f": "setMaxReportDelay(uint256)",
		"91397ab4": "setProfitFactor(uint256)",
		"ec38a862": "setRewards(address)",
		"c7b9d530": "setStrategist(address)",
		"5b9f0016": "stakedBalance()",
		"1fe4a686": "strategist()",
		"01681a62": "sweep(address)",
		"440368a3": "tend()",
		"650d1880": "tendTrigger(uint256)",
		"fbfa77cf": "vault()",
		"ba28e59c": "virtualRewardsPool()",
		"1f1fcd51": "want()",
		"2e1a7d4d": "withdraw(uint256)",
		"34659dc5": "withdrawToConvexDepositTokens()",
	},
}

// StrategyBaseABI is the input ABI used to generate the binding from.
// Deprecated: Use StrategyBaseMetaData.ABI instead.
var StrategyBaseABI = StrategyBaseMetaData.ABI

// Deprecated: Use StrategyBaseMetaData.Sigs instead.
// StrategyBaseFuncSigs maps the 4-byte function signature to its string representation.
var StrategyBaseFuncSigs = StrategyBaseMetaData.Sigs

// StrategyBase is an auto generated Go binding around an Ethereum contract.
type StrategyBase struct {
	StrategyBaseCaller     // Read-only binding to the contract
	StrategyBaseTransactor // Write-only binding to the contract
	StrategyBaseFilterer   // Log filterer for contract events
}

// StrategyBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type StrategyBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StrategyBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StrategyBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StrategyBaseSession struct {
	Contract     *StrategyBase     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StrategyBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StrategyBaseCallerSession struct {
	Contract *StrategyBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// StrategyBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StrategyBaseTransactorSession struct {
	Contract     *StrategyBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StrategyBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type StrategyBaseRaw struct {
	Contract *StrategyBase // Generic contract binding to access the raw methods on
}

// StrategyBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StrategyBaseCallerRaw struct {
	Contract *StrategyBaseCaller // Generic read-only contract binding to access the raw methods on
}

// StrategyBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StrategyBaseTransactorRaw struct {
	Contract *StrategyBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStrategyBase creates a new instance of StrategyBase, bound to a specific deployed contract.
func NewStrategyBase(address common.Address, backend bind.ContractBackend) (*StrategyBase, error) {
	contract, err := bindStrategyBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StrategyBase{StrategyBaseCaller: StrategyBaseCaller{contract: contract}, StrategyBaseTransactor: StrategyBaseTransactor{contract: contract}, StrategyBaseFilterer: StrategyBaseFilterer{contract: contract}}, nil
}

// NewStrategyBaseCaller creates a new read-only instance of StrategyBase, bound to a specific deployed contract.
func NewStrategyBaseCaller(address common.Address, caller bind.ContractCaller) (*StrategyBaseCaller, error) {
	contract, err := bindStrategyBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StrategyBaseCaller{contract: contract}, nil
}

// NewStrategyBaseTransactor creates a new write-only instance of StrategyBase, bound to a specific deployed contract.
func NewStrategyBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*StrategyBaseTransactor, error) {
	contract, err := bindStrategyBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StrategyBaseTransactor{contract: contract}, nil
}

// NewStrategyBaseFilterer creates a new log filterer instance of StrategyBase, bound to a specific deployed contract.
func NewStrategyBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*StrategyBaseFilterer, error) {
	contract, err := bindStrategyBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StrategyBaseFilterer{contract: contract}, nil
}

// bindStrategyBase binds a generic wrapper to an already deployed contract.
func bindStrategyBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StrategyBaseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StrategyBase *StrategyBaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StrategyBase.Contract.StrategyBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StrategyBase *StrategyBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyBase.Contract.StrategyBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StrategyBase *StrategyBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StrategyBase.Contract.StrategyBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StrategyBase *StrategyBaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StrategyBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StrategyBase *StrategyBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StrategyBase *StrategyBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StrategyBase.Contract.contract.Transact(opts, method, params...)
}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() pure returns(string)
func (_StrategyBase *StrategyBaseCaller) ApiVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "apiVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() pure returns(string)
func (_StrategyBase *StrategyBaseSession) ApiVersion() (string, error) {
	return _StrategyBase.Contract.ApiVersion(&_StrategyBase.CallOpts)
}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() pure returns(string)
func (_StrategyBase *StrategyBaseCallerSession) ApiVersion() (string, error) {
	return _StrategyBase.Contract.ApiVersion(&_StrategyBase.CallOpts)
}

// BalanceOfWant is a free data retrieval call binding the contract method 0xc1a3d44c.
//
// Solidity: function balanceOfWant() view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) BalanceOfWant(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "balanceOfWant")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfWant is a free data retrieval call binding the contract method 0xc1a3d44c.
//
// Solidity: function balanceOfWant() view returns(uint256)
func (_StrategyBase *StrategyBaseSession) BalanceOfWant() (*big.Int, error) {
	return _StrategyBase.Contract.BalanceOfWant(&_StrategyBase.CallOpts)
}

// BalanceOfWant is a free data retrieval call binding the contract method 0xc1a3d44c.
//
// Solidity: function balanceOfWant() view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) BalanceOfWant() (*big.Int, error) {
	return _StrategyBase.Contract.BalanceOfWant(&_StrategyBase.CallOpts)
}

// ClaimRewards is a free data retrieval call binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() view returns(bool)
func (_StrategyBase *StrategyBaseCaller) ClaimRewards(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "claimRewards")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ClaimRewards is a free data retrieval call binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() view returns(bool)
func (_StrategyBase *StrategyBaseSession) ClaimRewards() (bool, error) {
	return _StrategyBase.Contract.ClaimRewards(&_StrategyBase.CallOpts)
}

// ClaimRewards is a free data retrieval call binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() view returns(bool)
func (_StrategyBase *StrategyBaseCallerSession) ClaimRewards() (bool, error) {
	return _StrategyBase.Contract.ClaimRewards(&_StrategyBase.CallOpts)
}

// ClaimableBalance is a free data retrieval call binding the contract method 0xc4f45423.
//
// Solidity: function claimableBalance() view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) ClaimableBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "claimableBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimableBalance is a free data retrieval call binding the contract method 0xc4f45423.
//
// Solidity: function claimableBalance() view returns(uint256)
func (_StrategyBase *StrategyBaseSession) ClaimableBalance() (*big.Int, error) {
	return _StrategyBase.Contract.ClaimableBalance(&_StrategyBase.CallOpts)
}

// ClaimableBalance is a free data retrieval call binding the contract method 0xc4f45423.
//
// Solidity: function claimableBalance() view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) ClaimableBalance() (*big.Int, error) {
	return _StrategyBase.Contract.ClaimableBalance(&_StrategyBase.CallOpts)
}

// CreditThreshold is a free data retrieval call binding the contract method 0xaa5480cf.
//
// Solidity: function creditThreshold() view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) CreditThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "creditThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreditThreshold is a free data retrieval call binding the contract method 0xaa5480cf.
//
// Solidity: function creditThreshold() view returns(uint256)
func (_StrategyBase *StrategyBaseSession) CreditThreshold() (*big.Int, error) {
	return _StrategyBase.Contract.CreditThreshold(&_StrategyBase.CallOpts)
}

// CreditThreshold is a free data retrieval call binding the contract method 0xaa5480cf.
//
// Solidity: function creditThreshold() view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) CreditThreshold() (*big.Int, error) {
	return _StrategyBase.Contract.CreditThreshold(&_StrategyBase.CallOpts)
}

// DebtThreshold is a free data retrieval call binding the contract method 0x1d12f28b.
//
// Solidity: function debtThreshold() view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) DebtThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "debtThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtThreshold is a free data retrieval call binding the contract method 0x1d12f28b.
//
// Solidity: function debtThreshold() view returns(uint256)
func (_StrategyBase *StrategyBaseSession) DebtThreshold() (*big.Int, error) {
	return _StrategyBase.Contract.DebtThreshold(&_StrategyBase.CallOpts)
}

// DebtThreshold is a free data retrieval call binding the contract method 0x1d12f28b.
//
// Solidity: function debtThreshold() view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) DebtThreshold() (*big.Int, error) {
	return _StrategyBase.Contract.DebtThreshold(&_StrategyBase.CallOpts)
}

// DelegatedAssets is a free data retrieval call binding the contract method 0x8e6350e2.
//
// Solidity: function delegatedAssets() view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) DelegatedAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "delegatedAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegatedAssets is a free data retrieval call binding the contract method 0x8e6350e2.
//
// Solidity: function delegatedAssets() view returns(uint256)
func (_StrategyBase *StrategyBaseSession) DelegatedAssets() (*big.Int, error) {
	return _StrategyBase.Contract.DelegatedAssets(&_StrategyBase.CallOpts)
}

// DelegatedAssets is a free data retrieval call binding the contract method 0x8e6350e2.
//
// Solidity: function delegatedAssets() view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) DelegatedAssets() (*big.Int, error) {
	return _StrategyBase.Contract.DelegatedAssets(&_StrategyBase.CallOpts)
}

// DoHealthCheck is a free data retrieval call binding the contract method 0x6718835f.
//
// Solidity: function doHealthCheck() view returns(bool)
func (_StrategyBase *StrategyBaseCaller) DoHealthCheck(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "doHealthCheck")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DoHealthCheck is a free data retrieval call binding the contract method 0x6718835f.
//
// Solidity: function doHealthCheck() view returns(bool)
func (_StrategyBase *StrategyBaseSession) DoHealthCheck() (bool, error) {
	return _StrategyBase.Contract.DoHealthCheck(&_StrategyBase.CallOpts)
}

// DoHealthCheck is a free data retrieval call binding the contract method 0x6718835f.
//
// Solidity: function doHealthCheck() view returns(bool)
func (_StrategyBase *StrategyBaseCallerSession) DoHealthCheck() (bool, error) {
	return _StrategyBase.Contract.DoHealthCheck(&_StrategyBase.CallOpts)
}

// EmergencyExit is a free data retrieval call binding the contract method 0x5641ec03.
//
// Solidity: function emergencyExit() view returns(bool)
func (_StrategyBase *StrategyBaseCaller) EmergencyExit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "emergencyExit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EmergencyExit is a free data retrieval call binding the contract method 0x5641ec03.
//
// Solidity: function emergencyExit() view returns(bool)
func (_StrategyBase *StrategyBaseSession) EmergencyExit() (bool, error) {
	return _StrategyBase.Contract.EmergencyExit(&_StrategyBase.CallOpts)
}

// EmergencyExit is a free data retrieval call binding the contract method 0x5641ec03.
//
// Solidity: function emergencyExit() view returns(bool)
func (_StrategyBase *StrategyBaseCallerSession) EmergencyExit() (bool, error) {
	return _StrategyBase.Contract.EmergencyExit(&_StrategyBase.CallOpts)
}

// EstimatedTotalAssets is a free data retrieval call binding the contract method 0xefbb5cb0.
//
// Solidity: function estimatedTotalAssets() view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) EstimatedTotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "estimatedTotalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimatedTotalAssets is a free data retrieval call binding the contract method 0xefbb5cb0.
//
// Solidity: function estimatedTotalAssets() view returns(uint256)
func (_StrategyBase *StrategyBaseSession) EstimatedTotalAssets() (*big.Int, error) {
	return _StrategyBase.Contract.EstimatedTotalAssets(&_StrategyBase.CallOpts)
}

// EstimatedTotalAssets is a free data retrieval call binding the contract method 0xefbb5cb0.
//
// Solidity: function estimatedTotalAssets() view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) EstimatedTotalAssets() (*big.Int, error) {
	return _StrategyBase.Contract.EstimatedTotalAssets(&_StrategyBase.CallOpts)
}

// HarvestProfitMax is a free data retrieval call binding the contract method 0x3b7c6e2f.
//
// Solidity: function harvestProfitMax() view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) HarvestProfitMax(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "harvestProfitMax")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HarvestProfitMax is a free data retrieval call binding the contract method 0x3b7c6e2f.
//
// Solidity: function harvestProfitMax() view returns(uint256)
func (_StrategyBase *StrategyBaseSession) HarvestProfitMax() (*big.Int, error) {
	return _StrategyBase.Contract.HarvestProfitMax(&_StrategyBase.CallOpts)
}

// HarvestProfitMax is a free data retrieval call binding the contract method 0x3b7c6e2f.
//
// Solidity: function harvestProfitMax() view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) HarvestProfitMax() (*big.Int, error) {
	return _StrategyBase.Contract.HarvestProfitMax(&_StrategyBase.CallOpts)
}

// HarvestProfitMin is a free data retrieval call binding the contract method 0x5fbeb25f.
//
// Solidity: function harvestProfitMin() view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) HarvestProfitMin(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "harvestProfitMin")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HarvestProfitMin is a free data retrieval call binding the contract method 0x5fbeb25f.
//
// Solidity: function harvestProfitMin() view returns(uint256)
func (_StrategyBase *StrategyBaseSession) HarvestProfitMin() (*big.Int, error) {
	return _StrategyBase.Contract.HarvestProfitMin(&_StrategyBase.CallOpts)
}

// HarvestProfitMin is a free data retrieval call binding the contract method 0x5fbeb25f.
//
// Solidity: function harvestProfitMin() view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) HarvestProfitMin() (*big.Int, error) {
	return _StrategyBase.Contract.HarvestProfitMin(&_StrategyBase.CallOpts)
}

// HarvestTrigger is a free data retrieval call binding the contract method 0xed882c2b.
//
// Solidity: function harvestTrigger(uint256 callCost) view returns(bool)
func (_StrategyBase *StrategyBaseCaller) HarvestTrigger(opts *bind.CallOpts, callCost *big.Int) (bool, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "harvestTrigger", callCost)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HarvestTrigger is a free data retrieval call binding the contract method 0xed882c2b.
//
// Solidity: function harvestTrigger(uint256 callCost) view returns(bool)
func (_StrategyBase *StrategyBaseSession) HarvestTrigger(callCost *big.Int) (bool, error) {
	return _StrategyBase.Contract.HarvestTrigger(&_StrategyBase.CallOpts, callCost)
}

// HarvestTrigger is a free data retrieval call binding the contract method 0xed882c2b.
//
// Solidity: function harvestTrigger(uint256 callCost) view returns(bool)
func (_StrategyBase *StrategyBaseCallerSession) HarvestTrigger(callCost *big.Int) (bool, error) {
	return _StrategyBase.Contract.HarvestTrigger(&_StrategyBase.CallOpts, callCost)
}

// HealthCheck is a free data retrieval call binding the contract method 0xb252720b.
//
// Solidity: function healthCheck() view returns(address)
func (_StrategyBase *StrategyBaseCaller) HealthCheck(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "healthCheck")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HealthCheck is a free data retrieval call binding the contract method 0xb252720b.
//
// Solidity: function healthCheck() view returns(address)
func (_StrategyBase *StrategyBaseSession) HealthCheck() (common.Address, error) {
	return _StrategyBase.Contract.HealthCheck(&_StrategyBase.CallOpts)
}

// HealthCheck is a free data retrieval call binding the contract method 0xb252720b.
//
// Solidity: function healthCheck() view returns(address)
func (_StrategyBase *StrategyBaseCallerSession) HealthCheck() (common.Address, error) {
	return _StrategyBase.Contract.HealthCheck(&_StrategyBase.CallOpts)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_StrategyBase *StrategyBaseCaller) IsActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "isActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_StrategyBase *StrategyBaseSession) IsActive() (bool, error) {
	return _StrategyBase.Contract.IsActive(&_StrategyBase.CallOpts)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_StrategyBase *StrategyBaseCallerSession) IsActive() (bool, error) {
	return _StrategyBase.Contract.IsActive(&_StrategyBase.CallOpts)
}

// KeepCRV is a free data retrieval call binding the contract method 0x7fef901a.
//
// Solidity: function keepCRV() view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) KeepCRV(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "keepCRV")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KeepCRV is a free data retrieval call binding the contract method 0x7fef901a.
//
// Solidity: function keepCRV() view returns(uint256)
func (_StrategyBase *StrategyBaseSession) KeepCRV() (*big.Int, error) {
	return _StrategyBase.Contract.KeepCRV(&_StrategyBase.CallOpts)
}

// KeepCRV is a free data retrieval call binding the contract method 0x7fef901a.
//
// Solidity: function keepCRV() view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) KeepCRV() (*big.Int, error) {
	return _StrategyBase.Contract.KeepCRV(&_StrategyBase.CallOpts)
}

// KeepCrvPercent is a free data retrieval call binding the contract method 0x7fef901a.
//
// Solidity: function keepCrvPercent() view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) KeepCrvPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "keepCrvPercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// KeepCrvPercent is a free data retrieval call binding the contract method 0x7fef901a.
//
// Solidity: function keepCrvPercent() view returns(uint256)
func (_StrategyBase *StrategyBaseSession) KeepCrvPercent() (*big.Int, error) {
	return _StrategyBase.Contract.KeepCrvPercent(&_StrategyBase.CallOpts)
}

// KeepCrvPercent is a free data retrieval call binding the contract method 0x7fef901a.
//
// Solidity: function keepCrvPercent() view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) KeepCrvPercent() (*big.Int, error) {
	return _StrategyBase.Contract.KeepCrvPercent(&_StrategyBase.CallOpts)
}

// KeepCVX is a free data retrieval call binding the contract method 0x4b31217e.
//
// Solidity: function keepCVX() view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) KeepCVX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "keepCVX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KeepCVX is a free data retrieval call binding the contract method 0x4b31217e.
//
// Solidity: function keepCVX() view returns(uint256)
func (_StrategyBase *StrategyBaseSession) KeepCVX() (*big.Int, error) {
	return _StrategyBase.Contract.KeepCVX(&_StrategyBase.CallOpts)
}

// KeepCVX is a free data retrieval call binding the contract method 0x4b31217e.
//
// Solidity: function keepCVX() view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) KeepCVX() (*big.Int, error) {
	return _StrategyBase.Contract.KeepCVX(&_StrategyBase.CallOpts)
}

// KeepCVXDestination is a free data retrieval call binding the contract method 0xef86b23c.
//
// Solidity: function keepCVXDestination() view returns(address)
func (_StrategyBase *StrategyBaseCaller) KeepCVXDestination(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "keepCVXDestination")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KeepCVXDestination is a free data retrieval call binding the contract method 0xef86b23c.
//
// Solidity: function keepCVXDestination() view returns(address)
func (_StrategyBase *StrategyBaseSession) KeepCVXDestination() (common.Address, error) {
	return _StrategyBase.Contract.KeepCVXDestination(&_StrategyBase.CallOpts)
}

// KeepCVXDestination is a free data retrieval call binding the contract method 0xef86b23c.
//
// Solidity: function keepCVXDestination() view returns(address)
func (_StrategyBase *StrategyBaseCallerSession) KeepCVXDestination() (common.Address, error) {
	return _StrategyBase.Contract.KeepCVXDestination(&_StrategyBase.CallOpts)
}

// Keeper is a free data retrieval call binding the contract method 0xaced1661.
//
// Solidity: function keeper() view returns(address)
func (_StrategyBase *StrategyBaseCaller) Keeper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "keeper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Keeper is a free data retrieval call binding the contract method 0xaced1661.
//
// Solidity: function keeper() view returns(address)
func (_StrategyBase *StrategyBaseSession) Keeper() (common.Address, error) {
	return _StrategyBase.Contract.Keeper(&_StrategyBase.CallOpts)
}

// Keeper is a free data retrieval call binding the contract method 0xaced1661.
//
// Solidity: function keeper() view returns(address)
func (_StrategyBase *StrategyBaseCallerSession) Keeper() (common.Address, error) {
	return _StrategyBase.Contract.Keeper(&_StrategyBase.CallOpts)
}

// MaxReportDelay is a free data retrieval call binding the contract method 0x28b7ccf7.
//
// Solidity: function maxReportDelay() view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) MaxReportDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "maxReportDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxReportDelay is a free data retrieval call binding the contract method 0x28b7ccf7.
//
// Solidity: function maxReportDelay() view returns(uint256)
func (_StrategyBase *StrategyBaseSession) MaxReportDelay() (*big.Int, error) {
	return _StrategyBase.Contract.MaxReportDelay(&_StrategyBase.CallOpts)
}

// MaxReportDelay is a free data retrieval call binding the contract method 0x28b7ccf7.
//
// Solidity: function maxReportDelay() view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) MaxReportDelay() (*big.Int, error) {
	return _StrategyBase.Contract.MaxReportDelay(&_StrategyBase.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StrategyBase *StrategyBaseCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StrategyBase *StrategyBaseSession) Name() (string, error) {
	return _StrategyBase.Contract.Name(&_StrategyBase.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StrategyBase *StrategyBaseCallerSession) Name() (string, error) {
	return _StrategyBase.Contract.Name(&_StrategyBase.CallOpts)
}

// Pid is a free data retrieval call binding the contract method 0xf1068454.
//
// Solidity: function pid() view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) Pid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "pid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Pid is a free data retrieval call binding the contract method 0xf1068454.
//
// Solidity: function pid() view returns(uint256)
func (_StrategyBase *StrategyBaseSession) Pid() (*big.Int, error) {
	return _StrategyBase.Contract.Pid(&_StrategyBase.CallOpts)
}

// Pid is a free data retrieval call binding the contract method 0xf1068454.
//
// Solidity: function pid() view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) Pid() (*big.Int, error) {
	return _StrategyBase.Contract.Pid(&_StrategyBase.CallOpts)
}

// ProfitFactor is a free data retrieval call binding the contract method 0x8cdfe166.
//
// Solidity: function profitFactor() view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) ProfitFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "profitFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProfitFactor is a free data retrieval call binding the contract method 0x8cdfe166.
//
// Solidity: function profitFactor() view returns(uint256)
func (_StrategyBase *StrategyBaseSession) ProfitFactor() (*big.Int, error) {
	return _StrategyBase.Contract.ProfitFactor(&_StrategyBase.CallOpts)
}

// ProfitFactor is a free data retrieval call binding the contract method 0x8cdfe166.
//
// Solidity: function profitFactor() view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) ProfitFactor() (*big.Int, error) {
	return _StrategyBase.Contract.ProfitFactor(&_StrategyBase.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x9ec5a894.
//
// Solidity: function rewards() view returns(address)
func (_StrategyBase *StrategyBaseCaller) Rewards(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "rewards")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0x9ec5a894.
//
// Solidity: function rewards() view returns(address)
func (_StrategyBase *StrategyBaseSession) Rewards() (common.Address, error) {
	return _StrategyBase.Contract.Rewards(&_StrategyBase.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x9ec5a894.
//
// Solidity: function rewards() view returns(address)
func (_StrategyBase *StrategyBaseCallerSession) Rewards() (common.Address, error) {
	return _StrategyBase.Contract.Rewards(&_StrategyBase.CallOpts)
}

// RewardsContract is a free data retrieval call binding the contract method 0x220cce97.
//
// Solidity: function rewardsContract() view returns(address)
func (_StrategyBase *StrategyBaseCaller) RewardsContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "rewardsContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsContract is a free data retrieval call binding the contract method 0x220cce97.
//
// Solidity: function rewardsContract() view returns(address)
func (_StrategyBase *StrategyBaseSession) RewardsContract() (common.Address, error) {
	return _StrategyBase.Contract.RewardsContract(&_StrategyBase.CallOpts)
}

// RewardsContract is a free data retrieval call binding the contract method 0x220cce97.
//
// Solidity: function rewardsContract() view returns(address)
func (_StrategyBase *StrategyBaseCallerSession) RewardsContract() (common.Address, error) {
	return _StrategyBase.Contract.RewardsContract(&_StrategyBase.CallOpts)
}

// StakedBalance is a free data retrieval call binding the contract method 0x5b9f0016.
//
// Solidity: function stakedBalance() view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) StakedBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "stakedBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakedBalance is a free data retrieval call binding the contract method 0x5b9f0016.
//
// Solidity: function stakedBalance() view returns(uint256)
func (_StrategyBase *StrategyBaseSession) StakedBalance() (*big.Int, error) {
	return _StrategyBase.Contract.StakedBalance(&_StrategyBase.CallOpts)
}

// StakedBalance is a free data retrieval call binding the contract method 0x5b9f0016.
//
// Solidity: function stakedBalance() view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) StakedBalance() (*big.Int, error) {
	return _StrategyBase.Contract.StakedBalance(&_StrategyBase.CallOpts)
}

// Strategist is a free data retrieval call binding the contract method 0x1fe4a686.
//
// Solidity: function strategist() view returns(address)
func (_StrategyBase *StrategyBaseCaller) Strategist(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "strategist")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Strategist is a free data retrieval call binding the contract method 0x1fe4a686.
//
// Solidity: function strategist() view returns(address)
func (_StrategyBase *StrategyBaseSession) Strategist() (common.Address, error) {
	return _StrategyBase.Contract.Strategist(&_StrategyBase.CallOpts)
}

// Strategist is a free data retrieval call binding the contract method 0x1fe4a686.
//
// Solidity: function strategist() view returns(address)
func (_StrategyBase *StrategyBaseCallerSession) Strategist() (common.Address, error) {
	return _StrategyBase.Contract.Strategist(&_StrategyBase.CallOpts)
}

// TendTrigger is a free data retrieval call binding the contract method 0x650d1880.
//
// Solidity: function tendTrigger(uint256 callCost) view returns(bool)
func (_StrategyBase *StrategyBaseCaller) TendTrigger(opts *bind.CallOpts, callCost *big.Int) (bool, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "tendTrigger", callCost)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TendTrigger is a free data retrieval call binding the contract method 0x650d1880.
//
// Solidity: function tendTrigger(uint256 callCost) view returns(bool)
func (_StrategyBase *StrategyBaseSession) TendTrigger(callCost *big.Int) (bool, error) {
	return _StrategyBase.Contract.TendTrigger(&_StrategyBase.CallOpts, callCost)
}

// TendTrigger is a free data retrieval call binding the contract method 0x650d1880.
//
// Solidity: function tendTrigger(uint256 callCost) view returns(bool)
func (_StrategyBase *StrategyBaseCallerSession) TendTrigger(callCost *big.Int) (bool, error) {
	return _StrategyBase.Contract.TendTrigger(&_StrategyBase.CallOpts, callCost)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_StrategyBase *StrategyBaseCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_StrategyBase *StrategyBaseSession) Vault() (common.Address, error) {
	return _StrategyBase.Contract.Vault(&_StrategyBase.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_StrategyBase *StrategyBaseCallerSession) Vault() (common.Address, error) {
	return _StrategyBase.Contract.Vault(&_StrategyBase.CallOpts)
}

// VirtualRewardsPool is a free data retrieval call binding the contract method 0xba28e59c.
//
// Solidity: function virtualRewardsPool() view returns(address)
func (_StrategyBase *StrategyBaseCaller) VirtualRewardsPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "virtualRewardsPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VirtualRewardsPool is a free data retrieval call binding the contract method 0xba28e59c.
//
// Solidity: function virtualRewardsPool() view returns(address)
func (_StrategyBase *StrategyBaseSession) VirtualRewardsPool() (common.Address, error) {
	return _StrategyBase.Contract.VirtualRewardsPool(&_StrategyBase.CallOpts)
}

// VirtualRewardsPool is a free data retrieval call binding the contract method 0xba28e59c.
//
// Solidity: function virtualRewardsPool() view returns(address)
func (_StrategyBase *StrategyBaseCallerSession) VirtualRewardsPool() (common.Address, error) {
	return _StrategyBase.Contract.VirtualRewardsPool(&_StrategyBase.CallOpts)
}

// Want is a free data retrieval call binding the contract method 0x1f1fcd51.
//
// Solidity: function want() view returns(address)
func (_StrategyBase *StrategyBaseCaller) Want(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "want")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Want is a free data retrieval call binding the contract method 0x1f1fcd51.
//
// Solidity: function want() view returns(address)
func (_StrategyBase *StrategyBaseSession) Want() (common.Address, error) {
	return _StrategyBase.Contract.Want(&_StrategyBase.CallOpts)
}

// Want is a free data retrieval call binding the contract method 0x1f1fcd51.
//
// Solidity: function want() view returns(address)
func (_StrategyBase *StrategyBaseCallerSession) Want() (common.Address, error) {
	return _StrategyBase.Contract.Want(&_StrategyBase.CallOpts)
}

// Harvest is a paid mutator transaction binding the contract method 0x4641257d.
//
// Solidity: function harvest() returns()
func (_StrategyBase *StrategyBaseTransactor) Harvest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "harvest")
}

// Harvest is a paid mutator transaction binding the contract method 0x4641257d.
//
// Solidity: function harvest() returns()
func (_StrategyBase *StrategyBaseSession) Harvest() (*types.Transaction, error) {
	return _StrategyBase.Contract.Harvest(&_StrategyBase.TransactOpts)
}

// Harvest is a paid mutator transaction binding the contract method 0x4641257d.
//
// Solidity: function harvest() returns()
func (_StrategyBase *StrategyBaseTransactorSession) Harvest() (*types.Transaction, error) {
	return _StrategyBase.Contract.Harvest(&_StrategyBase.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _newStrategy) returns()
func (_StrategyBase *StrategyBaseTransactor) Migrate(opts *bind.TransactOpts, _newStrategy common.Address) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "migrate", _newStrategy)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _newStrategy) returns()
func (_StrategyBase *StrategyBaseSession) Migrate(_newStrategy common.Address) (*types.Transaction, error) {
	return _StrategyBase.Contract.Migrate(&_StrategyBase.TransactOpts, _newStrategy)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _newStrategy) returns()
func (_StrategyBase *StrategyBaseTransactorSession) Migrate(_newStrategy common.Address) (*types.Transaction, error) {
	return _StrategyBase.Contract.Migrate(&_StrategyBase.TransactOpts, _newStrategy)
}

// SetClaimRewards is a paid mutator transaction binding the contract method 0xa98f9296.
//
// Solidity: function setClaimRewards(bool _claimRewards) returns()
func (_StrategyBase *StrategyBaseTransactor) SetClaimRewards(opts *bind.TransactOpts, _claimRewards bool) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "setClaimRewards", _claimRewards)
}

// SetClaimRewards is a paid mutator transaction binding the contract method 0xa98f9296.
//
// Solidity: function setClaimRewards(bool _claimRewards) returns()
func (_StrategyBase *StrategyBaseSession) SetClaimRewards(_claimRewards bool) (*types.Transaction, error) {
	return _StrategyBase.Contract.SetClaimRewards(&_StrategyBase.TransactOpts, _claimRewards)
}

// SetClaimRewards is a paid mutator transaction binding the contract method 0xa98f9296.
//
// Solidity: function setClaimRewards(bool _claimRewards) returns()
func (_StrategyBase *StrategyBaseTransactorSession) SetClaimRewards(_claimRewards bool) (*types.Transaction, error) {
	return _StrategyBase.Contract.SetClaimRewards(&_StrategyBase.TransactOpts, _claimRewards)
}

// SetDebtThreshold is a paid mutator transaction binding the contract method 0x0f969b87.
//
// Solidity: function setDebtThreshold(uint256 _debtThreshold) returns()
func (_StrategyBase *StrategyBaseTransactor) SetDebtThreshold(opts *bind.TransactOpts, _debtThreshold *big.Int) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "setDebtThreshold", _debtThreshold)
}

// SetDebtThreshold is a paid mutator transaction binding the contract method 0x0f969b87.
//
// Solidity: function setDebtThreshold(uint256 _debtThreshold) returns()
func (_StrategyBase *StrategyBaseSession) SetDebtThreshold(_debtThreshold *big.Int) (*types.Transaction, error) {
	return _StrategyBase.Contract.SetDebtThreshold(&_StrategyBase.TransactOpts, _debtThreshold)
}

// SetDebtThreshold is a paid mutator transaction binding the contract method 0x0f969b87.
//
// Solidity: function setDebtThreshold(uint256 _debtThreshold) returns()
func (_StrategyBase *StrategyBaseTransactorSession) SetDebtThreshold(_debtThreshold *big.Int) (*types.Transaction, error) {
	return _StrategyBase.Contract.SetDebtThreshold(&_StrategyBase.TransactOpts, _debtThreshold)
}

// SetDoHealthCheck is a paid mutator transaction binding the contract method 0xac00ff26.
//
// Solidity: function setDoHealthCheck(bool _doHealthCheck) returns()
func (_StrategyBase *StrategyBaseTransactor) SetDoHealthCheck(opts *bind.TransactOpts, _doHealthCheck bool) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "setDoHealthCheck", _doHealthCheck)
}

// SetDoHealthCheck is a paid mutator transaction binding the contract method 0xac00ff26.
//
// Solidity: function setDoHealthCheck(bool _doHealthCheck) returns()
func (_StrategyBase *StrategyBaseSession) SetDoHealthCheck(_doHealthCheck bool) (*types.Transaction, error) {
	return _StrategyBase.Contract.SetDoHealthCheck(&_StrategyBase.TransactOpts, _doHealthCheck)
}

// SetDoHealthCheck is a paid mutator transaction binding the contract method 0xac00ff26.
//
// Solidity: function setDoHealthCheck(bool _doHealthCheck) returns()
func (_StrategyBase *StrategyBaseTransactorSession) SetDoHealthCheck(_doHealthCheck bool) (*types.Transaction, error) {
	return _StrategyBase.Contract.SetDoHealthCheck(&_StrategyBase.TransactOpts, _doHealthCheck)
}

// SetEmergencyExit is a paid mutator transaction binding the contract method 0xfcf2d0ad.
//
// Solidity: function setEmergencyExit() returns()
func (_StrategyBase *StrategyBaseTransactor) SetEmergencyExit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "setEmergencyExit")
}

// SetEmergencyExit is a paid mutator transaction binding the contract method 0xfcf2d0ad.
//
// Solidity: function setEmergencyExit() returns()
func (_StrategyBase *StrategyBaseSession) SetEmergencyExit() (*types.Transaction, error) {
	return _StrategyBase.Contract.SetEmergencyExit(&_StrategyBase.TransactOpts)
}

// SetEmergencyExit is a paid mutator transaction binding the contract method 0xfcf2d0ad.
//
// Solidity: function setEmergencyExit() returns()
func (_StrategyBase *StrategyBaseTransactorSession) SetEmergencyExit() (*types.Transaction, error) {
	return _StrategyBase.Contract.SetEmergencyExit(&_StrategyBase.TransactOpts)
}

// SetForceHarvestTriggerOnce is a paid mutator transaction binding the contract method 0x0ada4dab.
//
// Solidity: function setForceHarvestTriggerOnce(bool _forceHarvestTriggerOnce) returns()
func (_StrategyBase *StrategyBaseTransactor) SetForceHarvestTriggerOnce(opts *bind.TransactOpts, _forceHarvestTriggerOnce bool) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "setForceHarvestTriggerOnce", _forceHarvestTriggerOnce)
}

// SetForceHarvestTriggerOnce is a paid mutator transaction binding the contract method 0x0ada4dab.
//
// Solidity: function setForceHarvestTriggerOnce(bool _forceHarvestTriggerOnce) returns()
func (_StrategyBase *StrategyBaseSession) SetForceHarvestTriggerOnce(_forceHarvestTriggerOnce bool) (*types.Transaction, error) {
	return _StrategyBase.Contract.SetForceHarvestTriggerOnce(&_StrategyBase.TransactOpts, _forceHarvestTriggerOnce)
}

// SetForceHarvestTriggerOnce is a paid mutator transaction binding the contract method 0x0ada4dab.
//
// Solidity: function setForceHarvestTriggerOnce(bool _forceHarvestTriggerOnce) returns()
func (_StrategyBase *StrategyBaseTransactorSession) SetForceHarvestTriggerOnce(_forceHarvestTriggerOnce bool) (*types.Transaction, error) {
	return _StrategyBase.Contract.SetForceHarvestTriggerOnce(&_StrategyBase.TransactOpts, _forceHarvestTriggerOnce)
}

// SetHealthCheck is a paid mutator transaction binding the contract method 0x11bc8245.
//
// Solidity: function setHealthCheck(address _healthCheck) returns()
func (_StrategyBase *StrategyBaseTransactor) SetHealthCheck(opts *bind.TransactOpts, _healthCheck common.Address) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "setHealthCheck", _healthCheck)
}

// SetHealthCheck is a paid mutator transaction binding the contract method 0x11bc8245.
//
// Solidity: function setHealthCheck(address _healthCheck) returns()
func (_StrategyBase *StrategyBaseSession) SetHealthCheck(_healthCheck common.Address) (*types.Transaction, error) {
	return _StrategyBase.Contract.SetHealthCheck(&_StrategyBase.TransactOpts, _healthCheck)
}

// SetHealthCheck is a paid mutator transaction binding the contract method 0x11bc8245.
//
// Solidity: function setHealthCheck(address _healthCheck) returns()
func (_StrategyBase *StrategyBaseTransactorSession) SetHealthCheck(_healthCheck common.Address) (*types.Transaction, error) {
	return _StrategyBase.Contract.SetHealthCheck(&_StrategyBase.TransactOpts, _healthCheck)
}

// SetKeep is a paid mutator transaction binding the contract method 0x1111fe1c.
//
// Solidity: function setKeep(uint256 _keepCRV, uint256 _keepCVX, address _keepCVXDestination) returns()
func (_StrategyBase *StrategyBaseTransactor) SetKeep(opts *bind.TransactOpts, _keepCRV *big.Int, _keepCVX *big.Int, _keepCVXDestination common.Address) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "setKeep", _keepCRV, _keepCVX, _keepCVXDestination)
}

// SetKeep is a paid mutator transaction binding the contract method 0x1111fe1c.
//
// Solidity: function setKeep(uint256 _keepCRV, uint256 _keepCVX, address _keepCVXDestination) returns()
func (_StrategyBase *StrategyBaseSession) SetKeep(_keepCRV *big.Int, _keepCVX *big.Int, _keepCVXDestination common.Address) (*types.Transaction, error) {
	return _StrategyBase.Contract.SetKeep(&_StrategyBase.TransactOpts, _keepCRV, _keepCVX, _keepCVXDestination)
}

// SetKeep is a paid mutator transaction binding the contract method 0x1111fe1c.
//
// Solidity: function setKeep(uint256 _keepCRV, uint256 _keepCVX, address _keepCVXDestination) returns()
func (_StrategyBase *StrategyBaseTransactorSession) SetKeep(_keepCRV *big.Int, _keepCVX *big.Int, _keepCVXDestination common.Address) (*types.Transaction, error) {
	return _StrategyBase.Contract.SetKeep(&_StrategyBase.TransactOpts, _keepCRV, _keepCVX, _keepCVXDestination)
}

// SetKeeper is a paid mutator transaction binding the contract method 0x748747e6.
//
// Solidity: function setKeeper(address _keeper) returns()
func (_StrategyBase *StrategyBaseTransactor) SetKeeper(opts *bind.TransactOpts, _keeper common.Address) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "setKeeper", _keeper)
}

// SetKeeper is a paid mutator transaction binding the contract method 0x748747e6.
//
// Solidity: function setKeeper(address _keeper) returns()
func (_StrategyBase *StrategyBaseSession) SetKeeper(_keeper common.Address) (*types.Transaction, error) {
	return _StrategyBase.Contract.SetKeeper(&_StrategyBase.TransactOpts, _keeper)
}

// SetKeeper is a paid mutator transaction binding the contract method 0x748747e6.
//
// Solidity: function setKeeper(address _keeper) returns()
func (_StrategyBase *StrategyBaseTransactorSession) SetKeeper(_keeper common.Address) (*types.Transaction, error) {
	return _StrategyBase.Contract.SetKeeper(&_StrategyBase.TransactOpts, _keeper)
}

// SetMaxReportDelay is a paid mutator transaction binding the contract method 0xf017c92f.
//
// Solidity: function setMaxReportDelay(uint256 _delay) returns()
func (_StrategyBase *StrategyBaseTransactor) SetMaxReportDelay(opts *bind.TransactOpts, _delay *big.Int) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "setMaxReportDelay", _delay)
}

// SetMaxReportDelay is a paid mutator transaction binding the contract method 0xf017c92f.
//
// Solidity: function setMaxReportDelay(uint256 _delay) returns()
func (_StrategyBase *StrategyBaseSession) SetMaxReportDelay(_delay *big.Int) (*types.Transaction, error) {
	return _StrategyBase.Contract.SetMaxReportDelay(&_StrategyBase.TransactOpts, _delay)
}

// SetMaxReportDelay is a paid mutator transaction binding the contract method 0xf017c92f.
//
// Solidity: function setMaxReportDelay(uint256 _delay) returns()
func (_StrategyBase *StrategyBaseTransactorSession) SetMaxReportDelay(_delay *big.Int) (*types.Transaction, error) {
	return _StrategyBase.Contract.SetMaxReportDelay(&_StrategyBase.TransactOpts, _delay)
}

// SetProfitFactor is a paid mutator transaction binding the contract method 0x91397ab4.
//
// Solidity: function setProfitFactor(uint256 _profitFactor) returns()
func (_StrategyBase *StrategyBaseTransactor) SetProfitFactor(opts *bind.TransactOpts, _profitFactor *big.Int) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "setProfitFactor", _profitFactor)
}

// SetProfitFactor is a paid mutator transaction binding the contract method 0x91397ab4.
//
// Solidity: function setProfitFactor(uint256 _profitFactor) returns()
func (_StrategyBase *StrategyBaseSession) SetProfitFactor(_profitFactor *big.Int) (*types.Transaction, error) {
	return _StrategyBase.Contract.SetProfitFactor(&_StrategyBase.TransactOpts, _profitFactor)
}

// SetProfitFactor is a paid mutator transaction binding the contract method 0x91397ab4.
//
// Solidity: function setProfitFactor(uint256 _profitFactor) returns()
func (_StrategyBase *StrategyBaseTransactorSession) SetProfitFactor(_profitFactor *big.Int) (*types.Transaction, error) {
	return _StrategyBase.Contract.SetProfitFactor(&_StrategyBase.TransactOpts, _profitFactor)
}

// SetRewards is a paid mutator transaction binding the contract method 0xec38a862.
//
// Solidity: function setRewards(address _rewards) returns()
func (_StrategyBase *StrategyBaseTransactor) SetRewards(opts *bind.TransactOpts, _rewards common.Address) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "setRewards", _rewards)
}

// SetRewards is a paid mutator transaction binding the contract method 0xec38a862.
//
// Solidity: function setRewards(address _rewards) returns()
func (_StrategyBase *StrategyBaseSession) SetRewards(_rewards common.Address) (*types.Transaction, error) {
	return _StrategyBase.Contract.SetRewards(&_StrategyBase.TransactOpts, _rewards)
}

// SetRewards is a paid mutator transaction binding the contract method 0xec38a862.
//
// Solidity: function setRewards(address _rewards) returns()
func (_StrategyBase *StrategyBaseTransactorSession) SetRewards(_rewards common.Address) (*types.Transaction, error) {
	return _StrategyBase.Contract.SetRewards(&_StrategyBase.TransactOpts, _rewards)
}

// SetStrategist is a paid mutator transaction binding the contract method 0xc7b9d530.
//
// Solidity: function setStrategist(address _strategist) returns()
func (_StrategyBase *StrategyBaseTransactor) SetStrategist(opts *bind.TransactOpts, _strategist common.Address) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "setStrategist", _strategist)
}

// SetStrategist is a paid mutator transaction binding the contract method 0xc7b9d530.
//
// Solidity: function setStrategist(address _strategist) returns()
func (_StrategyBase *StrategyBaseSession) SetStrategist(_strategist common.Address) (*types.Transaction, error) {
	return _StrategyBase.Contract.SetStrategist(&_StrategyBase.TransactOpts, _strategist)
}

// SetStrategist is a paid mutator transaction binding the contract method 0xc7b9d530.
//
// Solidity: function setStrategist(address _strategist) returns()
func (_StrategyBase *StrategyBaseTransactorSession) SetStrategist(_strategist common.Address) (*types.Transaction, error) {
	return _StrategyBase.Contract.SetStrategist(&_StrategyBase.TransactOpts, _strategist)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address _token) returns()
func (_StrategyBase *StrategyBaseTransactor) Sweep(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "sweep", _token)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address _token) returns()
func (_StrategyBase *StrategyBaseSession) Sweep(_token common.Address) (*types.Transaction, error) {
	return _StrategyBase.Contract.Sweep(&_StrategyBase.TransactOpts, _token)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address _token) returns()
func (_StrategyBase *StrategyBaseTransactorSession) Sweep(_token common.Address) (*types.Transaction, error) {
	return _StrategyBase.Contract.Sweep(&_StrategyBase.TransactOpts, _token)
}

// Tend is a paid mutator transaction binding the contract method 0x440368a3.
//
// Solidity: function tend() returns()
func (_StrategyBase *StrategyBaseTransactor) Tend(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "tend")
}

// Tend is a paid mutator transaction binding the contract method 0x440368a3.
//
// Solidity: function tend() returns()
func (_StrategyBase *StrategyBaseSession) Tend() (*types.Transaction, error) {
	return _StrategyBase.Contract.Tend(&_StrategyBase.TransactOpts)
}

// Tend is a paid mutator transaction binding the contract method 0x440368a3.
//
// Solidity: function tend() returns()
func (_StrategyBase *StrategyBaseTransactorSession) Tend() (*types.Transaction, error) {
	return _StrategyBase.Contract.Tend(&_StrategyBase.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amountNeeded) returns(uint256 _loss)
func (_StrategyBase *StrategyBaseTransactor) Withdraw(opts *bind.TransactOpts, _amountNeeded *big.Int) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "withdraw", _amountNeeded)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amountNeeded) returns(uint256 _loss)
func (_StrategyBase *StrategyBaseSession) Withdraw(_amountNeeded *big.Int) (*types.Transaction, error) {
	return _StrategyBase.Contract.Withdraw(&_StrategyBase.TransactOpts, _amountNeeded)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amountNeeded) returns(uint256 _loss)
func (_StrategyBase *StrategyBaseTransactorSession) Withdraw(_amountNeeded *big.Int) (*types.Transaction, error) {
	return _StrategyBase.Contract.Withdraw(&_StrategyBase.TransactOpts, _amountNeeded)
}

// WithdrawToConvexDepositTokens is a paid mutator transaction binding the contract method 0x34659dc5.
//
// Solidity: function withdrawToConvexDepositTokens() returns()
func (_StrategyBase *StrategyBaseTransactor) WithdrawToConvexDepositTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "withdrawToConvexDepositTokens")
}

// WithdrawToConvexDepositTokens is a paid mutator transaction binding the contract method 0x34659dc5.
//
// Solidity: function withdrawToConvexDepositTokens() returns()
func (_StrategyBase *StrategyBaseSession) WithdrawToConvexDepositTokens() (*types.Transaction, error) {
	return _StrategyBase.Contract.WithdrawToConvexDepositTokens(&_StrategyBase.TransactOpts)
}

// WithdrawToConvexDepositTokens is a paid mutator transaction binding the contract method 0x34659dc5.
//
// Solidity: function withdrawToConvexDepositTokens() returns()
func (_StrategyBase *StrategyBaseTransactorSession) WithdrawToConvexDepositTokens() (*types.Transaction, error) {
	return _StrategyBase.Contract.WithdrawToConvexDepositTokens(&_StrategyBase.TransactOpts)
}

// StrategyBaseEmergencyExitEnabledIterator is returned from FilterEmergencyExitEnabled and is used to iterate over the raw logs and unpacked data for EmergencyExitEnabled events raised by the StrategyBase contract.
type StrategyBaseEmergencyExitEnabledIterator struct {
	Event *StrategyBaseEmergencyExitEnabled // Event containing the contract specifics and raw log

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
func (it *StrategyBaseEmergencyExitEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyBaseEmergencyExitEnabled)
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
		it.Event = new(StrategyBaseEmergencyExitEnabled)
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
func (it *StrategyBaseEmergencyExitEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyBaseEmergencyExitEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyBaseEmergencyExitEnabled represents a EmergencyExitEnabled event raised by the StrategyBase contract.
type StrategyBaseEmergencyExitEnabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyExitEnabled is a free log retrieval operation binding the contract event 0x97e963041e952738788b9d4871d854d282065b8f90a464928d6528f2e9a4fd0b.
//
// Solidity: event EmergencyExitEnabled()
func (_StrategyBase *StrategyBaseFilterer) FilterEmergencyExitEnabled(opts *bind.FilterOpts) (*StrategyBaseEmergencyExitEnabledIterator, error) {

	logs, sub, err := _StrategyBase.contract.FilterLogs(opts, "EmergencyExitEnabled")
	if err != nil {
		return nil, err
	}
	return &StrategyBaseEmergencyExitEnabledIterator{contract: _StrategyBase.contract, event: "EmergencyExitEnabled", logs: logs, sub: sub}, nil
}

// WatchEmergencyExitEnabled is a free log subscription operation binding the contract event 0x97e963041e952738788b9d4871d854d282065b8f90a464928d6528f2e9a4fd0b.
//
// Solidity: event EmergencyExitEnabled()
func (_StrategyBase *StrategyBaseFilterer) WatchEmergencyExitEnabled(opts *bind.WatchOpts, sink chan<- *StrategyBaseEmergencyExitEnabled) (event.Subscription, error) {

	logs, sub, err := _StrategyBase.contract.WatchLogs(opts, "EmergencyExitEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyBaseEmergencyExitEnabled)
				if err := _StrategyBase.contract.UnpackLog(event, "EmergencyExitEnabled", log); err != nil {
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

// ParseEmergencyExitEnabled is a log parse operation binding the contract event 0x97e963041e952738788b9d4871d854d282065b8f90a464928d6528f2e9a4fd0b.
//
// Solidity: event EmergencyExitEnabled()
func (_StrategyBase *StrategyBaseFilterer) ParseEmergencyExitEnabled(log types.Log) (*StrategyBaseEmergencyExitEnabled, error) {
	event := new(StrategyBaseEmergencyExitEnabled)
	if err := _StrategyBase.contract.UnpackLog(event, "EmergencyExitEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyBaseHarvestedIterator is returned from FilterHarvested and is used to iterate over the raw logs and unpacked data for Harvested events raised by the StrategyBase contract.
type StrategyBaseHarvestedIterator struct {
	Event *StrategyBaseHarvested // Event containing the contract specifics and raw log

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
func (it *StrategyBaseHarvestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyBaseHarvested)
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
		it.Event = new(StrategyBaseHarvested)
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
func (it *StrategyBaseHarvestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyBaseHarvestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyBaseHarvested represents a Harvested event raised by the StrategyBase contract.
type StrategyBaseHarvested struct {
	Profit          *big.Int
	Loss            *big.Int
	DebtPayment     *big.Int
	DebtOutstanding *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterHarvested is a free log retrieval operation binding the contract event 0x4c0f499ffe6befa0ca7c826b0916cf87bea98de658013e76938489368d60d509.
//
// Solidity: event Harvested(uint256 profit, uint256 loss, uint256 debtPayment, uint256 debtOutstanding)
func (_StrategyBase *StrategyBaseFilterer) FilterHarvested(opts *bind.FilterOpts) (*StrategyBaseHarvestedIterator, error) {

	logs, sub, err := _StrategyBase.contract.FilterLogs(opts, "Harvested")
	if err != nil {
		return nil, err
	}
	return &StrategyBaseHarvestedIterator{contract: _StrategyBase.contract, event: "Harvested", logs: logs, sub: sub}, nil
}

// WatchHarvested is a free log subscription operation binding the contract event 0x4c0f499ffe6befa0ca7c826b0916cf87bea98de658013e76938489368d60d509.
//
// Solidity: event Harvested(uint256 profit, uint256 loss, uint256 debtPayment, uint256 debtOutstanding)
func (_StrategyBase *StrategyBaseFilterer) WatchHarvested(opts *bind.WatchOpts, sink chan<- *StrategyBaseHarvested) (event.Subscription, error) {

	logs, sub, err := _StrategyBase.contract.WatchLogs(opts, "Harvested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyBaseHarvested)
				if err := _StrategyBase.contract.UnpackLog(event, "Harvested", log); err != nil {
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

// ParseHarvested is a log parse operation binding the contract event 0x4c0f499ffe6befa0ca7c826b0916cf87bea98de658013e76938489368d60d509.
//
// Solidity: event Harvested(uint256 profit, uint256 loss, uint256 debtPayment, uint256 debtOutstanding)
func (_StrategyBase *StrategyBaseFilterer) ParseHarvested(log types.Log) (*StrategyBaseHarvested, error) {
	event := new(StrategyBaseHarvested)
	if err := _StrategyBase.contract.UnpackLog(event, "Harvested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyBaseUpdatedDebtThresholdIterator is returned from FilterUpdatedDebtThreshold and is used to iterate over the raw logs and unpacked data for UpdatedDebtThreshold events raised by the StrategyBase contract.
type StrategyBaseUpdatedDebtThresholdIterator struct {
	Event *StrategyBaseUpdatedDebtThreshold // Event containing the contract specifics and raw log

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
func (it *StrategyBaseUpdatedDebtThresholdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyBaseUpdatedDebtThreshold)
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
		it.Event = new(StrategyBaseUpdatedDebtThreshold)
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
func (it *StrategyBaseUpdatedDebtThresholdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyBaseUpdatedDebtThresholdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyBaseUpdatedDebtThreshold represents a UpdatedDebtThreshold event raised by the StrategyBase contract.
type StrategyBaseUpdatedDebtThreshold struct {
	DebtThreshold *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedDebtThreshold is a free log retrieval operation binding the contract event 0xa68ba126373d04c004c5748c300c9fca12bd444b3d4332e261f3bd2bac4a8600.
//
// Solidity: event UpdatedDebtThreshold(uint256 debtThreshold)
func (_StrategyBase *StrategyBaseFilterer) FilterUpdatedDebtThreshold(opts *bind.FilterOpts) (*StrategyBaseUpdatedDebtThresholdIterator, error) {

	logs, sub, err := _StrategyBase.contract.FilterLogs(opts, "UpdatedDebtThreshold")
	if err != nil {
		return nil, err
	}
	return &StrategyBaseUpdatedDebtThresholdIterator{contract: _StrategyBase.contract, event: "UpdatedDebtThreshold", logs: logs, sub: sub}, nil
}

// WatchUpdatedDebtThreshold is a free log subscription operation binding the contract event 0xa68ba126373d04c004c5748c300c9fca12bd444b3d4332e261f3bd2bac4a8600.
//
// Solidity: event UpdatedDebtThreshold(uint256 debtThreshold)
func (_StrategyBase *StrategyBaseFilterer) WatchUpdatedDebtThreshold(opts *bind.WatchOpts, sink chan<- *StrategyBaseUpdatedDebtThreshold) (event.Subscription, error) {

	logs, sub, err := _StrategyBase.contract.WatchLogs(opts, "UpdatedDebtThreshold")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyBaseUpdatedDebtThreshold)
				if err := _StrategyBase.contract.UnpackLog(event, "UpdatedDebtThreshold", log); err != nil {
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

// ParseUpdatedDebtThreshold is a log parse operation binding the contract event 0xa68ba126373d04c004c5748c300c9fca12bd444b3d4332e261f3bd2bac4a8600.
//
// Solidity: event UpdatedDebtThreshold(uint256 debtThreshold)
func (_StrategyBase *StrategyBaseFilterer) ParseUpdatedDebtThreshold(log types.Log) (*StrategyBaseUpdatedDebtThreshold, error) {
	event := new(StrategyBaseUpdatedDebtThreshold)
	if err := _StrategyBase.contract.UnpackLog(event, "UpdatedDebtThreshold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyBaseUpdatedKeeperIterator is returned from FilterUpdatedKeeper and is used to iterate over the raw logs and unpacked data for UpdatedKeeper events raised by the StrategyBase contract.
type StrategyBaseUpdatedKeeperIterator struct {
	Event *StrategyBaseUpdatedKeeper // Event containing the contract specifics and raw log

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
func (it *StrategyBaseUpdatedKeeperIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyBaseUpdatedKeeper)
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
		it.Event = new(StrategyBaseUpdatedKeeper)
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
func (it *StrategyBaseUpdatedKeeperIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyBaseUpdatedKeeperIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyBaseUpdatedKeeper represents a UpdatedKeeper event raised by the StrategyBase contract.
type StrategyBaseUpdatedKeeper struct {
	NewKeeper common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdatedKeeper is a free log retrieval operation binding the contract event 0x2f202ddb4a2e345f6323ed90f8fc8559d770a7abbbeee84dde8aca3351fe7154.
//
// Solidity: event UpdatedKeeper(address newKeeper)
func (_StrategyBase *StrategyBaseFilterer) FilterUpdatedKeeper(opts *bind.FilterOpts) (*StrategyBaseUpdatedKeeperIterator, error) {

	logs, sub, err := _StrategyBase.contract.FilterLogs(opts, "UpdatedKeeper")
	if err != nil {
		return nil, err
	}
	return &StrategyBaseUpdatedKeeperIterator{contract: _StrategyBase.contract, event: "UpdatedKeeper", logs: logs, sub: sub}, nil
}

// WatchUpdatedKeeper is a free log subscription operation binding the contract event 0x2f202ddb4a2e345f6323ed90f8fc8559d770a7abbbeee84dde8aca3351fe7154.
//
// Solidity: event UpdatedKeeper(address newKeeper)
func (_StrategyBase *StrategyBaseFilterer) WatchUpdatedKeeper(opts *bind.WatchOpts, sink chan<- *StrategyBaseUpdatedKeeper) (event.Subscription, error) {

	logs, sub, err := _StrategyBase.contract.WatchLogs(opts, "UpdatedKeeper")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyBaseUpdatedKeeper)
				if err := _StrategyBase.contract.UnpackLog(event, "UpdatedKeeper", log); err != nil {
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

// ParseUpdatedKeeper is a log parse operation binding the contract event 0x2f202ddb4a2e345f6323ed90f8fc8559d770a7abbbeee84dde8aca3351fe7154.
//
// Solidity: event UpdatedKeeper(address newKeeper)
func (_StrategyBase *StrategyBaseFilterer) ParseUpdatedKeeper(log types.Log) (*StrategyBaseUpdatedKeeper, error) {
	event := new(StrategyBaseUpdatedKeeper)
	if err := _StrategyBase.contract.UnpackLog(event, "UpdatedKeeper", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyBaseUpdatedProfitFactorIterator is returned from FilterUpdatedProfitFactor and is used to iterate over the raw logs and unpacked data for UpdatedProfitFactor events raised by the StrategyBase contract.
type StrategyBaseUpdatedProfitFactorIterator struct {
	Event *StrategyBaseUpdatedProfitFactor // Event containing the contract specifics and raw log

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
func (it *StrategyBaseUpdatedProfitFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyBaseUpdatedProfitFactor)
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
		it.Event = new(StrategyBaseUpdatedProfitFactor)
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
func (it *StrategyBaseUpdatedProfitFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyBaseUpdatedProfitFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyBaseUpdatedProfitFactor represents a UpdatedProfitFactor event raised by the StrategyBase contract.
type StrategyBaseUpdatedProfitFactor struct {
	ProfitFactor *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdatedProfitFactor is a free log retrieval operation binding the contract event 0xd94596337df4c2f0f44d30a7fc5db1c7bb60d9aca4185ed77c6fd96eb45ec298.
//
// Solidity: event UpdatedProfitFactor(uint256 profitFactor)
func (_StrategyBase *StrategyBaseFilterer) FilterUpdatedProfitFactor(opts *bind.FilterOpts) (*StrategyBaseUpdatedProfitFactorIterator, error) {

	logs, sub, err := _StrategyBase.contract.FilterLogs(opts, "UpdatedProfitFactor")
	if err != nil {
		return nil, err
	}
	return &StrategyBaseUpdatedProfitFactorIterator{contract: _StrategyBase.contract, event: "UpdatedProfitFactor", logs: logs, sub: sub}, nil
}

// WatchUpdatedProfitFactor is a free log subscription operation binding the contract event 0xd94596337df4c2f0f44d30a7fc5db1c7bb60d9aca4185ed77c6fd96eb45ec298.
//
// Solidity: event UpdatedProfitFactor(uint256 profitFactor)
func (_StrategyBase *StrategyBaseFilterer) WatchUpdatedProfitFactor(opts *bind.WatchOpts, sink chan<- *StrategyBaseUpdatedProfitFactor) (event.Subscription, error) {

	logs, sub, err := _StrategyBase.contract.WatchLogs(opts, "UpdatedProfitFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyBaseUpdatedProfitFactor)
				if err := _StrategyBase.contract.UnpackLog(event, "UpdatedProfitFactor", log); err != nil {
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

// ParseUpdatedProfitFactor is a log parse operation binding the contract event 0xd94596337df4c2f0f44d30a7fc5db1c7bb60d9aca4185ed77c6fd96eb45ec298.
//
// Solidity: event UpdatedProfitFactor(uint256 profitFactor)
func (_StrategyBase *StrategyBaseFilterer) ParseUpdatedProfitFactor(log types.Log) (*StrategyBaseUpdatedProfitFactor, error) {
	event := new(StrategyBaseUpdatedProfitFactor)
	if err := _StrategyBase.contract.UnpackLog(event, "UpdatedProfitFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyBaseUpdatedReportDelayIterator is returned from FilterUpdatedReportDelay and is used to iterate over the raw logs and unpacked data for UpdatedReportDelay events raised by the StrategyBase contract.
type StrategyBaseUpdatedReportDelayIterator struct {
	Event *StrategyBaseUpdatedReportDelay // Event containing the contract specifics and raw log

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
func (it *StrategyBaseUpdatedReportDelayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyBaseUpdatedReportDelay)
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
		it.Event = new(StrategyBaseUpdatedReportDelay)
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
func (it *StrategyBaseUpdatedReportDelayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyBaseUpdatedReportDelayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyBaseUpdatedReportDelay represents a UpdatedReportDelay event raised by the StrategyBase contract.
type StrategyBaseUpdatedReportDelay struct {
	Delay *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdatedReportDelay is a free log retrieval operation binding the contract event 0x4aaf232568bff365c53cad69bdb6e83014e79df80216ceba8ee01769723dfd68.
//
// Solidity: event UpdatedReportDelay(uint256 delay)
func (_StrategyBase *StrategyBaseFilterer) FilterUpdatedReportDelay(opts *bind.FilterOpts) (*StrategyBaseUpdatedReportDelayIterator, error) {

	logs, sub, err := _StrategyBase.contract.FilterLogs(opts, "UpdatedReportDelay")
	if err != nil {
		return nil, err
	}
	return &StrategyBaseUpdatedReportDelayIterator{contract: _StrategyBase.contract, event: "UpdatedReportDelay", logs: logs, sub: sub}, nil
}

// WatchUpdatedReportDelay is a free log subscription operation binding the contract event 0x4aaf232568bff365c53cad69bdb6e83014e79df80216ceba8ee01769723dfd68.
//
// Solidity: event UpdatedReportDelay(uint256 delay)
func (_StrategyBase *StrategyBaseFilterer) WatchUpdatedReportDelay(opts *bind.WatchOpts, sink chan<- *StrategyBaseUpdatedReportDelay) (event.Subscription, error) {

	logs, sub, err := _StrategyBase.contract.WatchLogs(opts, "UpdatedReportDelay")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyBaseUpdatedReportDelay)
				if err := _StrategyBase.contract.UnpackLog(event, "UpdatedReportDelay", log); err != nil {
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

// ParseUpdatedReportDelay is a log parse operation binding the contract event 0x4aaf232568bff365c53cad69bdb6e83014e79df80216ceba8ee01769723dfd68.
//
// Solidity: event UpdatedReportDelay(uint256 delay)
func (_StrategyBase *StrategyBaseFilterer) ParseUpdatedReportDelay(log types.Log) (*StrategyBaseUpdatedReportDelay, error) {
	event := new(StrategyBaseUpdatedReportDelay)
	if err := _StrategyBase.contract.UnpackLog(event, "UpdatedReportDelay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyBaseUpdatedRewardsIterator is returned from FilterUpdatedRewards and is used to iterate over the raw logs and unpacked data for UpdatedRewards events raised by the StrategyBase contract.
type StrategyBaseUpdatedRewardsIterator struct {
	Event *StrategyBaseUpdatedRewards // Event containing the contract specifics and raw log

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
func (it *StrategyBaseUpdatedRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyBaseUpdatedRewards)
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
		it.Event = new(StrategyBaseUpdatedRewards)
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
func (it *StrategyBaseUpdatedRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyBaseUpdatedRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyBaseUpdatedRewards represents a UpdatedRewards event raised by the StrategyBase contract.
type StrategyBaseUpdatedRewards struct {
	Rewards common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdatedRewards is a free log retrieval operation binding the contract event 0xafbb66abf8f3b719799940473a4052a3717cdd8e40fb6c8a3faadab316b1a069.
//
// Solidity: event UpdatedRewards(address rewards)
func (_StrategyBase *StrategyBaseFilterer) FilterUpdatedRewards(opts *bind.FilterOpts) (*StrategyBaseUpdatedRewardsIterator, error) {

	logs, sub, err := _StrategyBase.contract.FilterLogs(opts, "UpdatedRewards")
	if err != nil {
		return nil, err
	}
	return &StrategyBaseUpdatedRewardsIterator{contract: _StrategyBase.contract, event: "UpdatedRewards", logs: logs, sub: sub}, nil
}

// WatchUpdatedRewards is a free log subscription operation binding the contract event 0xafbb66abf8f3b719799940473a4052a3717cdd8e40fb6c8a3faadab316b1a069.
//
// Solidity: event UpdatedRewards(address rewards)
func (_StrategyBase *StrategyBaseFilterer) WatchUpdatedRewards(opts *bind.WatchOpts, sink chan<- *StrategyBaseUpdatedRewards) (event.Subscription, error) {

	logs, sub, err := _StrategyBase.contract.WatchLogs(opts, "UpdatedRewards")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyBaseUpdatedRewards)
				if err := _StrategyBase.contract.UnpackLog(event, "UpdatedRewards", log); err != nil {
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

// ParseUpdatedRewards is a log parse operation binding the contract event 0xafbb66abf8f3b719799940473a4052a3717cdd8e40fb6c8a3faadab316b1a069.
//
// Solidity: event UpdatedRewards(address rewards)
func (_StrategyBase *StrategyBaseFilterer) ParseUpdatedRewards(log types.Log) (*StrategyBaseUpdatedRewards, error) {
	event := new(StrategyBaseUpdatedRewards)
	if err := _StrategyBase.contract.UnpackLog(event, "UpdatedRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyBaseUpdatedStrategistIterator is returned from FilterUpdatedStrategist and is used to iterate over the raw logs and unpacked data for UpdatedStrategist events raised by the StrategyBase contract.
type StrategyBaseUpdatedStrategistIterator struct {
	Event *StrategyBaseUpdatedStrategist // Event containing the contract specifics and raw log

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
func (it *StrategyBaseUpdatedStrategistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyBaseUpdatedStrategist)
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
		it.Event = new(StrategyBaseUpdatedStrategist)
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
func (it *StrategyBaseUpdatedStrategistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyBaseUpdatedStrategistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyBaseUpdatedStrategist represents a UpdatedStrategist event raised by the StrategyBase contract.
type StrategyBaseUpdatedStrategist struct {
	NewStrategist common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedStrategist is a free log retrieval operation binding the contract event 0x352ececae6d7d1e6d26bcf2c549dfd55be1637e9b22dc0cf3b71ddb36097a6b4.
//
// Solidity: event UpdatedStrategist(address newStrategist)
func (_StrategyBase *StrategyBaseFilterer) FilterUpdatedStrategist(opts *bind.FilterOpts) (*StrategyBaseUpdatedStrategistIterator, error) {

	logs, sub, err := _StrategyBase.contract.FilterLogs(opts, "UpdatedStrategist")
	if err != nil {
		return nil, err
	}
	return &StrategyBaseUpdatedStrategistIterator{contract: _StrategyBase.contract, event: "UpdatedStrategist", logs: logs, sub: sub}, nil
}

// WatchUpdatedStrategist is a free log subscription operation binding the contract event 0x352ececae6d7d1e6d26bcf2c549dfd55be1637e9b22dc0cf3b71ddb36097a6b4.
//
// Solidity: event UpdatedStrategist(address newStrategist)
func (_StrategyBase *StrategyBaseFilterer) WatchUpdatedStrategist(opts *bind.WatchOpts, sink chan<- *StrategyBaseUpdatedStrategist) (event.Subscription, error) {

	logs, sub, err := _StrategyBase.contract.WatchLogs(opts, "UpdatedStrategist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyBaseUpdatedStrategist)
				if err := _StrategyBase.contract.UnpackLog(event, "UpdatedStrategist", log); err != nil {
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

// ParseUpdatedStrategist is a log parse operation binding the contract event 0x352ececae6d7d1e6d26bcf2c549dfd55be1637e9b22dc0cf3b71ddb36097a6b4.
//
// Solidity: event UpdatedStrategist(address newStrategist)
func (_StrategyBase *StrategyBaseFilterer) ParseUpdatedStrategist(log types.Log) (*StrategyBaseUpdatedStrategist, error) {
	event := new(StrategyBaseUpdatedStrategist)
	if err := _StrategyBase.contract.UnpackLog(event, "UpdatedStrategist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyConvexstETHMetaData contains all meta data concerning the StrategyConvexstETH contract.
var StrategyConvexstETHMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_curvePool\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyExitEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"profit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loss\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debtPayment\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debtOutstanding\",\"type\":\"uint256\"}],\"name\":\"Harvested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debtThreshold\",\"type\":\"uint256\"}],\"name\":\"UpdatedDebtThreshold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newKeeper\",\"type\":\"address\"}],\"name\":\"UpdatedKeeper\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"profitFactor\",\"type\":\"uint256\"}],\"name\":\"UpdatedProfitFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"UpdatedReportDelay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rewards\",\"type\":\"address\"}],\"name\":\"UpdatedRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newStrategist\",\"type\":\"address\"}],\"name\":\"UpdatedStrategist\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"apiVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balanceOfWant\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkEarmark\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimRewards\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimableBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimableProfitInUsdt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"curve\",\"outputs\":[{\"internalType\":\"contractICurveFi\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"debtThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegatedAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"doHealthCheck\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyExit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"estimatedTotalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"harvest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"harvestProfitMax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"harvestProfitMin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"callCostinEth\",\"type\":\"uint256\"}],\"name\":\"harvestTrigger\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasRewards\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"healthCheck\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"keepCRV\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"keepCrvPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"keepCVX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"keepCVXDestination\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"keeper\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxReportDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newStrategy\",\"type\":\"address\"}],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"needsEarmarkReward\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"needsEarmark\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"profitFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsContract\",\"outputs\":[{\"internalType\":\"contractIConvexRewards\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_claimRewards\",\"type\":\"bool\"}],\"name\":\"setClaimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_debtThreshold\",\"type\":\"uint256\"}],\"name\":\"setDebtThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_doHealthCheck\",\"type\":\"bool\"}],\"name\":\"setDoHealthCheck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setEmergencyExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_forceHarvestTriggerOnce\",\"type\":\"bool\"}],\"name\":\"setForceHarvestTriggerOnce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_harvestProfitMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_harvestProfitMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_creditThreshold\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_checkEarmark\",\"type\":\"bool\"}],\"name\":\"setHarvestTriggerParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_healthCheck\",\"type\":\"address\"}],\"name\":\"setHealthCheck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_keepCRV\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_keepCVX\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_keepCVXDestination\",\"type\":\"address\"}],\"name\":\"setKeep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_keeper\",\"type\":\"address\"}],\"name\":\"setKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_delay\",\"type\":\"uint256\"}],\"name\":\"setMaxReportDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_profitFactor\",\"type\":\"uint256\"}],\"name\":\"setProfitFactor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewards\",\"type\":\"address\"}],\"name\":\"setRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategist\",\"type\":\"address\"}],\"name\":\"setStrategist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"strategist\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"sweep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"callCost\",\"type\":\"uint256\"}],\"name\":\"tendTrigger\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_hasRewards\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_rewardsIndex\",\"type\":\"uint256\"}],\"name\":\"updateRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"contractVaultAPI\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"virtualRewardsPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"want\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountNeeded\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_loss\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawToConvexDepositTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Sigs: map[string]string{
		"25829410": "apiVersion()",
		"c1a3d44c": "balanceOfWant()",
		"ec2f1050": "checkEarmark()",
		"372500ab": "claimRewards()",
		"c4f45423": "claimableBalance()",
		"06cfb3c0": "claimableProfitInUsdt()",
		"aa5480cf": "creditThreshold()",
		"7165485d": "curve()",
		"1d12f28b": "debtThreshold()",
		"8e6350e2": "delegatedAssets()",
		"6718835f": "doHealthCheck()",
		"5641ec03": "emergencyExit()",
		"efbb5cb0": "estimatedTotalAssets()",
		"4641257d": "harvest()",
		"3b7c6e2f": "harvestProfitMax()",
		"5fbeb25f": "harvestProfitMin()",
		"ed882c2b": "harvestTrigger(uint256)",
		"ecf04e15": "hasRewards()",
		"b252720b": "healthCheck()",
		"22f3e2d4": "isActive()",
		"7fef901a": "keepCRV()",
		"4b31217e": "keepCVX()",
		"ef86b23c": "keepCVXDestination()",
		"aced1661": "keeper()",
		"28b7ccf7": "maxReportDelay()",
		"ce5494bb": "migrate(address)",
		"06fdde03": "name()",
		"f09338df": "needsEarmarkReward()",
		"f1068454": "pid()",
		"8cdfe166": "profitFactor()",
		"9ec5a894": "rewards()",
		"220cce97": "rewardsContract()",
		"d1af0c7d": "rewardsToken()",
		"a98f9296": "setClaimRewards(bool)",
		"0f969b87": "setDebtThreshold(uint256)",
		"ac00ff26": "setDoHealthCheck(bool)",
		"fcf2d0ad": "setEmergencyExit()",
		"0ada4dab": "setForceHarvestTriggerOnce(bool)",
		"b4d48fd4": "setHarvestTriggerParams(uint256,uint256,uint256,bool)",
		"11bc8245": "setHealthCheck(address)",
		"1111fe1c": "setKeep(uint256,uint256,address)",
		"748747e6": "setKeeper(address)",
		"f017c92f": "setMaxReportDelay(uint256)",
		"91397ab4": "setProfitFactor(uint256)",
		"ec38a862": "setRewards(address)",
		"c7b9d530": "setStrategist(address)",
		"5b9f0016": "stakedBalance()",
		"1fe4a686": "strategist()",
		"01681a62": "sweep(address)",
		"440368a3": "tend()",
		"650d1880": "tendTrigger(uint256)",
		"dbad1ab1": "updateRewards(bool,uint256)",
		"fbfa77cf": "vault()",
		"ba28e59c": "virtualRewardsPool()",
		"1f1fcd51": "want()",
		"2e1a7d4d": "withdraw(uint256)",
		"34659dc5": "withdrawToConvexDepositTokens()",
	},
	Bin: "0x608060405262015180600655606460075560006008556017805460ff191660011790553480156200002f57600080fd5b5060405162004c4838038062004c488339810160408190526200005291620007c1565b838062000062813380806200007c565b50620000729050838383620002b0565b505050506200095c565b6005546001600160a01b031615620000b15760405162461bcd60e51b8152600401620000a890620008c2565b60405180910390fd5b600180546001600160a01b0319166001600160a01b03868116919091179182905560408051637e062a3560e11b81529051929091169163fc0c546a91600480820192602092909190829003018186803b1580156200010e57600080fd5b505afa15801562000123573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000149919062000710565b600580546001600160a01b0319166001600160a01b03928316179081905560405163095ea7b360e01b815291169063095ea7b3906200019190879060001990600401620008a9565b602060405180830381600087803b158015620001ac57600080fd5b505af1158015620001c1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001e7919062000887565b50600280546001600160a01b038086166001600160a01b03199283161790925560038054858416908316179081905560048054858516931692909217825560015460405163095ea7b360e01b81529084169363095ea7b3936200025393909116916000199101620008a9565b602060405180830381600087803b1580156200026e57600080fd5b505af115801562000283573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620002a9919062000887565b5050505050565b60145461010090046001600160a01b031615620002cc57600080fd5b621baf8060065560008054610100600160a81b03191674ddcea799ff1699e98edf118e0629a974df7df01200179055640df8475800600f55641bf08eb00060105568d8d726b7177a8000006011556103e8600c55600e80547393a62da5a14c80f265dabc077fcee437b1a0efde6001600160a01b031990911617905560055460405163095ea7b360e01b81526001600160a01b039091169063095ea7b390620003929073f403c135812408bfbe8713b5a23a04b3d48aae319060001990600401620008a9565b602060405180830381600087803b158015620003ad57600080fd5b505af1158015620003c2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620003e8919062000887565b5060405163095ea7b360e01b8152734e3fbd56cd56c3e72c1403e103b45db9da5b9d2b9063095ea7b3906200043a9073b576491f1e6e5e62f1d8f26062ee822b40b0e0d49060001990600401620008a9565b602060405180830381600087803b1580156200045557600080fd5b505af11580156200046a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000490919062000887565b5060405163095ea7b360e01b815273d533a949740bb3306d119cc777fa900ba034cd529063095ea7b390620004e290738301ae4fc9c624d1d396cbdaa1ed877821d7c5119060001990600401620008a9565b602060405180830381600087803b158015620004fd57600080fd5b505af115801562000512573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000538919062000887565b5060148054610100600160a81b0319166101006001600160a01b03851602179055600b839055604051631526fe2760e01b8152600090819073f403c135812408bfbe8713b5a23a04b3d48aae3190631526fe27906200059c908890600401620008f9565b60c06040518083038186803b158015620005b557600080fd5b505afa158015620005ca573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620005f0919062000735565b5050600980546001600160a01b0380841661010002610100600160a81b031990921691909117909155600554949650909450808616931692909214915062000639905057600080fd5b82516200064e90601390602086019062000656565b505050505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200069957805160ff1916838001178555620006c9565b82800160010185558215620006c9579182015b82811115620006c9578251825591602001919060010190620006ac565b50620006d7929150620006db565b5090565b5b80821115620006d75760008155600101620006dc565b80516001600160a01b03811681146200070a57600080fd5b92915050565b60006020828403121562000722578081fd5b6200072e8383620006f2565b9392505050565b60008060008060008060c087890312156200074e578182fd5b6200075a8888620006f2565b95506200076b8860208901620006f2565b94506200077c8860408901620006f2565b93506200078d8860608901620006f2565b92506200079e8860808901620006f2565b915060a08701518015158114620007b3578182fd5b809150509295509295509295565b60008060008060808587031215620007d7578384fd5b620007e38686620006f2565b935060208501519250620007fb8660408701620006f2565b60608601519092506001600160401b038082111562000818578283fd5b818701915087601f8301126200082c578283fd5b8151818111156200083b578384fd5b62000850601f8201601f191660200162000902565b915080825288602082850101111562000867578384fd5b6200087a81602084016020860162000929565b5094979396509194505050565b60006020828403121562000899578081fd5b815180151581146200072e578182fd5b6001600160a01b03929092168252602082015260400190565b6020808252601c908201527f537472617465677920616c726561647920696e697469616c697a656400000000604082015260600190565b90815260200190565b6040518181016001600160401b03811182821017156200092157600080fd5b604052919050565b60005b83811015620009465781810151838201526020016200092c565b8381111562000956576000848401525b50505050565b6142dc806200096c6000396000f3fe60806040526004361061037a5760003560e01c80637fef901a116101d1578063c7b9d53011610102578063ed882c2b116100a0578063f09338df1161006f578063f09338df146108e4578063f1068454146108f9578063fbfa77cf1461090e578063fcf2d0ad1461092357610381565b8063ed882c2b1461087a578063ef86b23c1461089a578063efbb5cb0146108af578063f017c92f146108c457610381565b8063dbad1ab1116100dc578063dbad1ab114610810578063ec2f105014610830578063ec38a86214610845578063ecf04e151461086557610381565b8063c7b9d530146107bb578063ce5494bb146107db578063d1af0c7d146107fb57610381565b8063ac00ff261161016f578063b4d48fd411610149578063b4d48fd41461075c578063ba28e59c1461077c578063c1a3d44c14610791578063c4f45423146107a657610381565b8063ac00ff2614610712578063aced166114610732578063b252720b1461074757610381565b806391397ab4116101ab57806391397ab4146106a85780639ec5a894146106c8578063a98f9296146106dd578063aa5480cf146106fd57610381565b80637fef901a146106695780638cdfe1661461067e5780638e6350e21461069357610381565b80632e1a7d4d116102ab5780635641ec0311610249578063650d188011610223578063650d1880146105ff5780636718835f1461061f5780637165485d14610634578063748747e61461064957610381565b80635641ec03146105c05780635b9f0016146105d55780635fbeb25f146105ea57610381565b80633b7c6e2f116102855780633b7c6e2f1461056c578063440368a3146105815780634641257d146105965780634b31217e146105ab57610381565b80632e1a7d4d1461052257806334659dc514610542578063372500ab1461055757610381565b80631d12f28b11610318578063220cce97116102f2578063220cce97146104c157806322f3e2d4146104d657806325829410146104f857806328b7ccf71461050d57610381565b80631d12f28b146104755780631f1fcd511461048a5780631fe4a686146104ac57610381565b80630ada4dab116103545780630ada4dab146103f55780630f969b87146104155780631111fe1c1461043557806311bc82451461045557610381565b806301681a621461038657806306cfb3c0146103a857806306fdde03146103d357610381565b3661038157005b600080fd5b34801561039257600080fd5b506103a66103a1366004613b65565b610938565b005b3480156103b457600080fd5b506103bd610b42565b6040516103ca91906140be565b60405180910390f35b3480156103df57600080fd5b506103e8611064565b6040516103ca9190613e7b565b34801561040157600080fd5b506103a6610410366004613c32565b6110fa565b34801561042157600080fd5b506103a6610430366004613d07565b61115a565b34801561044157600080fd5b506103a6610450366004613d37565b6111e7565b34801561046157600080fd5b506103a6610470366004613b65565b611266565b34801561048157600080fd5b506103bd6112c6565b34801561049657600080fd5b5061049f6112cc565b6040516103ca9190613dcb565b3480156104b857600080fd5b5061049f6112db565b3480156104cd57600080fd5b5061049f6112ea565b3480156104e257600080fd5b506104eb6112fe565b6040516103ca9190613e4b565b34801561050457600080fd5b506103e86113a0565b34801561051957600080fd5b506103bd6113bf565b34801561052e57600080fd5b506103bd61053d366004613d07565b6113c5565b34801561054e57600080fd5b506103a66114b4565b34801561056357600080fd5b506104eb6115a5565b34801561057857600080fd5b506103bd6115ae565b34801561058d57600080fd5b506103a66115b4565b3480156105a257600080fd5b506103a661169b565b3480156105b757600080fd5b506103bd611a34565b3480156105cc57600080fd5b506104eb611a3a565b3480156105e157600080fd5b506103bd611a43565b3480156105f657600080fd5b506103bd611ac9565b34801561060b57600080fd5b506104eb61061a366004613d07565b611acf565b34801561062b57600080fd5b506104eb611ad7565b34801561064057600080fd5b5061049f611ae0565b34801561065557600080fd5b506103a6610664366004613b65565b611af4565b34801561067557600080fd5b506103bd611b9f565b34801561068a57600080fd5b506103bd611ba5565b34801561069f57600080fd5b506103bd611bab565b3480156106b457600080fd5b506103a66106c3366004613d07565b611bb0565b3480156106d457600080fd5b5061049f611c32565b3480156106e957600080fd5b506103a66106f8366004613c32565b611c41565b34801561070957600080fd5b506103bd611ca1565b34801561071e57600080fd5b506103a661072d366004613c32565b611ca7565b34801561073e57600080fd5b5061049f611cf2565b34801561075357600080fd5b5061049f611d01565b34801561076857600080fd5b506103a6610777366004613d6f565b611d15565b34801561078857600080fd5b5061049f611d8f565b34801561079d57600080fd5b506103bd611d9e565b3480156107b257600080fd5b506103bd611dcf565b3480156107c757600080fd5b506103a66107d6366004613b65565b611e03565b3480156107e757600080fd5b506103a66107f6366004613b65565b611eae565b34801561080757600080fd5b5061049f61202a565b34801561081c57600080fd5b506103a661082b366004613c6a565b612039565b34801561083c57600080fd5b506104eb612397565b34801561085157600080fd5b506103a6610860366004613b65565b6123a7565b34801561087157600080fd5b506104eb61242f565b34801561088657600080fd5b506104eb610895366004613d07565b61243f565b3480156108a657600080fd5b5061049f612635565b3480156108bb57600080fd5b506103bd612644565b3480156108d057600080fd5b506103a66108df366004613d07565b612659565b3480156108f057600080fd5b506104eb6126db565b34801561090557600080fd5b506103bd612817565b34801561091a57600080fd5b5061049f61281d565b34801561092f57600080fd5b506103a661282c565b610940612910565b6001600160a01b0316336001600160a01b0316146109795760405162461bcd60e51b81526004016109709061402b565b60405180910390fd5b6005546001600160a01b03828116911614156109a75760405162461bcd60e51b815260040161097090613ed3565b6001546001600160a01b03828116911614156109d55760405162461bcd60e51b815260040161097090613fd3565b60606109df61298d565b905060005b8151811015610a3a578181815181106109f957fe5b60200260200101516001600160a01b0316836001600160a01b03161415610a325760405162461bcd60e51b81526004016109709061409a565b6001016109e4565b50816001600160a01b031663a9059cbb610a52612910565b6040516370a0823160e01b81526001600160a01b038616906370a0823190610a7e903090600401613dcb565b60206040518083038186803b158015610a9657600080fd5b505afa158015610aaa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ace9190613d1f565b6040518363ffffffff1660e01b8152600401610aeb929190613dfa565b602060405180830381600087803b158015610b0557600080fd5b505af1158015610b19573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b3d9190613c4e565b505050565b6000806103e8905060006a52b7d2dcc80cd2e40000009050600069152d02c7e14af680000090506000734e3fbd56cd56c3e72c1403e103b45db9da5b9d2b6001600160a01b03166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b158015610bb857600080fd5b505afa158015610bcc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bf09190613d1f565b9050600080610bff8385612992565b90506000610c0b611dcf565b905086821015610c56576000610c2188846129dd565b9050610c3788610c318484612a1f565b90612992565b93506000610c4588876129dd565b905080851115610c53578094505b50505b6000735f4ec3df9cbd43714fe2740f5e3616155c5b841990506000610ce76064836001600160a01b03166350d25bcd6040518163ffffffff1660e01b815260040160206040518083038186803b158015610caf57600080fd5b505afa158015610cc3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c319190613d1f565b90506000610d86670de0b6b3a7640000610c3184738301ae4fc9c624d1d396cbdaa1ed877821d7c5116001600160a01b03166386fc88d36040518163ffffffff1660e01b815260040160206040518083038186803b158015610d4857600080fd5b505afa158015610d5c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d809190613d1f565b90612a1f565b90506000610de7670de0b6b3a7640000610c318573b576491f1e6e5e62f1d8f26062ee822b40b0e0d46001600160a01b03166386fc88d36040518163ffffffff1660e01b815260040160206040518083038186803b158015610d4857600080fd5b90506000610e01670de0b6b3a7640000610c318589612a1f565b90506000610e1b670de0b6b3a7640000610c31858c612a1f565b601554909150600090600160a01b900460ff161561103c5760408051600380825260808201909252606091602082018380368337505060155482519293506001600160a01b031691839150600090610e6f57fe5b60200260200101906001600160a01b031690816001600160a01b03168152505073c02aaa39b223fe8d0a0e5c4f27ead9083c756cc281600181518110610eb157fe5b60200260200101906001600160a01b031690816001600160a01b03168152505073dac17f958d2ee523a2206206994597c13d831ec781600281518110610ef357fe5b6001600160a01b039283166020918202929092010152600a546040516246613160e11b81526000929190911690628cc26290610f33903090600401613dcb565b60206040518083038186803b158015610f4b57600080fd5b505afa158015610f5f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f839190613d1f565b905080156110395760405163d06ca61f60e01b815260609073d9e1ce17f2641f24ae83637ab66a2cca9c378b9f9063d06ca61f90610fc790859087906004016140c7565b60006040518083038186803b158015610fdf57600080fd5b505afa158015610ff3573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261101b9190810190613b9d565b90508060018251038151811061102d57fe5b60200260200101519350505b50505b6110508161104a8585612a59565b90612a59565b9e5050505050505050505050505050505b90565b60138054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156110f05780601f106110c5576101008083540402835291602001916110f0565b820191906000526020600020905b8154815290600101906020018083116110d357829003601f168201915b5050505050905090565b6002546001600160a01b031633148061112b5750611116612910565b6001600160a01b0316336001600160a01b0316145b6111475760405162461bcd60e51b81526004016109709061402b565b6012805460ff1916911515919091179055565b6002546001600160a01b031633148061118b5750611176612910565b6001600160a01b0316336001600160a01b0316145b6111a75760405162461bcd60e51b81526004016109709061402b565b60088190556040517fa68ba126373d04c004c5748c300c9fca12bd444b3d4332e261f3bd2bac4a8600906111dc9083906140be565b60405180910390a150565b6111ef612910565b6001600160a01b0316336001600160a01b03161461121f5760405162461bcd60e51b81526004016109709061402b565b612710831115801561123357506127108211155b61123c57600080fd5b600c92909255600d55600e80546001600160a01b0319166001600160a01b03909216919091179055565b61126e612910565b6001600160a01b0316336001600160a01b03161461129e5760405162461bcd60e51b81526004016109709061402b565b600080546001600160a01b0390921661010002610100600160a81b0319909216919091179055565b60085481565b6005546001600160a01b031681565b6002546001600160a01b031681565b60095461010090046001600160a01b031681565b6001546040516339ebf82360e01b815260009182916001600160a01b03909116906339ebf82390611333903090600401613dcb565b6101006040518083038186803b15801561134c57600080fd5b505afa158015611360573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113849190613c95565b60400151118061139b57506000611399612644565b115b905090565b6040805180820190915260058152640302e332e360dc1b602082015290565b60065481565b6001546000906001600160a01b031633146113f25760405162461bcd60e51b815260040161097090613fb3565b60006113fd83612a7e565b9250905061140b8183612a59565b83146114295760405162461bcd60e51b815260040161097090613f4f565b60055460405163a9059cbb60e01b81526001600160a01b039091169063a9059cbb9061145b9033908590600401613dfa565b602060405180830381600087803b15801561147557600080fd5b505af1158015611489573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114ad9190613c4e565b5050919050565b6002546001600160a01b03163314806114e557506114d0612910565b6001600160a01b0316336001600160a01b0316145b6115015760405162461bcd60e51b81526004016109709061402b565b600061150b611a43565b905080156115a257600954601454604051631c683a1b60e11b81526101009092046001600160a01b0316916338d074369161154e91859160ff169060040161411d565b602060405180830381600087803b15801561156857600080fd5b505af115801561157c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115a09190613c4e565b505b50565b60145460ff1681565b60105481565b6004546001600160a01b03163314806115d757506002546001600160a01b031633145b806115fa57506115e5612910565b6001600160a01b0316336001600160a01b0316145b6116165760405162461bcd60e51b81526004016109709061402b565b6001546040805163bf3759b560e01b81529051611699926001600160a01b03169163bf3759b5916004808301926020929190829003018186803b15801561165c57600080fd5b505afa158015611670573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116949190613d1f565b612b8b565b565b6004546001600160a01b03163314806116be57506002546001600160a01b031633145b806116e157506116cc612910565b6001600160a01b0316336001600160a01b0316145b6116fd5760405162461bcd60e51b81526004016109709061402b565b6000806000600160009054906101000a90046001600160a01b03166001600160a01b031663bf3759b56040518163ffffffff1660e01b815260040160206040518083038186803b15801561175057600080fd5b505afa158015611764573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117889190613d1f565b60095490915060009060ff16156117de5760006117a3612644565b90506117bc8382116117b557836117b7565b815b612a7e565b94509150828211156117d8576117d282846129dd565b94508291505b506117ef565b6117e782612beb565b919550935090505b6001546040516339ebf82360e01b81526000916001600160a01b0316906339ebf82390611820903090600401613dcb565b6101006040518083038186803b15801561183957600080fd5b505afa15801561184d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118719190613c95565b60a001516001546040516328766ebf60e21b81529192506001600160a01b03169063a1d9bafc906118aa908890889087906004016141b8565b602060405180830381600087803b1580156118c457600080fd5b505af11580156118d8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118fc9190613d1f565b9250611906613256565b61190f83612b8b565b60005460ff168015611930575060005461010090046001600160a01b031615155b156119e25760005460405163c70fa00b60e01b81526101009091046001600160a01b03169063c70fa00b9061197190889088908790899088906004016141e9565b60206040518083038186803b15801561198957600080fd5b505afa15801561199d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119c19190613c4e565b6119dd5760405162461bcd60e51b815260040161097090613f29565b6119f0565b6000805460ff191660011790555b7f4c0f499ffe6befa0ca7c826b0916cf87bea98de658013e76938489368d60d50985858486604051611a2594939291906141ce565b60405180910390a15050505050565b600d5481565b60095460ff1681565b6009546040516370a0823160e01b815260009161010090046001600160a01b0316906370a0823190611a79903090600401613dcb565b60206040518083038186803b158015611a9157600080fd5b505afa158015611aa5573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061139b9190613d1f565b600f5481565b60005b919050565b60005460ff1681565b60145461010090046001600160a01b031681565b6002546001600160a01b0316331480611b255750611b10612910565b6001600160a01b0316336001600160a01b0316145b611b415760405162461bcd60e51b81526004016109709061402b565b6001600160a01b038116611b5457600080fd5b600480546001600160a01b0319166001600160a01b0383161790556040517f2f202ddb4a2e345f6323ed90f8fc8559d770a7abbbeee84dde8aca3351fe7154906111dc908390613dcb565b600c5481565b60075481565b600090565b6002546001600160a01b0316331480611be15750611bcc612910565b6001600160a01b0316336001600160a01b0316145b611bfd5760405162461bcd60e51b81526004016109709061402b565b60078190556040517fd94596337df4c2f0f44d30a7fc5db1c7bb60d9aca4185ed77c6fd96eb45ec298906111dc9083906140be565b6003546001600160a01b031681565b6002546001600160a01b0316331480611c725750611c5d612910565b6001600160a01b0316336001600160a01b0316145b611c8e5760405162461bcd60e51b81526004016109709061402b565b6014805460ff1916911515919091179055565b60115481565b611caf612910565b6001600160a01b0316336001600160a01b031614611cdf5760405162461bcd60e51b81526004016109709061402b565b6000805460ff1916911515919091179055565b6004546001600160a01b031681565b60005461010090046001600160a01b031681565b6002546001600160a01b0316331480611d465750611d31612910565b6001600160a01b0316336001600160a01b0316145b611d625760405162461bcd60e51b81526004016109709061402b565b600f9390935560109190915560115560148054911515600160a81b0260ff60a81b19909216919091179055565b600a546001600160a01b031681565b6005546040516370a0823160e01b81526000916001600160a01b0316906370a0823190611a79903090600401613dcb565b6009546040516246613160e11b815260009161010090046001600160a01b031690628cc26290611a79903090600401613dcb565b6002546001600160a01b0316331480611e345750611e1f612910565b6001600160a01b0316336001600160a01b0316145b611e505760405162461bcd60e51b81526004016109709061402b565b6001600160a01b038116611e6357600080fd5b600280546001600160a01b0319166001600160a01b0383161790556040517f352ececae6d7d1e6d26bcf2c549dfd55be1637e9b22dc0cf3b71ddb36097a6b4906111dc908390613dcb565b6001546001600160a01b0316331480611edf5750611eca612910565b6001600160a01b0316336001600160a01b0316145b611ee857600080fd5b6001546040805163fbfa77cf60e01b815290516001600160a01b039283169284169163fbfa77cf916004808301926020929190829003018186803b158015611f2f57600080fd5b505afa158015611f43573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f679190613b81565b6001600160a01b031614611f7a57600080fd5b611f8381613315565b6005546040516370a0823160e01b81526001600160a01b039091169063a9059cbb90839083906370a0823190611fbd903090600401613dcb565b60206040518083038186803b158015611fd557600080fd5b505afa158015611fe9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061200d9190613d1f565b6040518363ffffffff1660e01b815260040161154e929190613dfa565b6015546001600160a01b031681565b612041612910565b6001600160a01b0316336001600160a01b0316146120715760405162461bcd60e51b81526004016109709061402b565b6015546001600160a01b0316158015906120aa57506015546001600160a01b0316734e3fbd56cd56c3e72c1403e103b45db9da5b9d2b14155b1561214a5760155460405163095ea7b360e01b81526001600160a01b039091169063095ea7b3906120f69073d9e1ce17f2641f24ae83637ab66a2cca9c378b9f90600090600401613dfa565b602060405180830381600087803b15801561211057600080fd5b505af1158015612124573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121489190613c4e565b505b8161217457601580546001600160a81b0319169055600a80546001600160a01b03191690556115a0565b600954604051632061aa2360e11b81526101009091046001600160a01b0316906340c35446906121a89084906004016140be565b60206040518083038186803b1580156121c057600080fd5b505afa1580156121d4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121f89190613b81565b600a80546001600160a01b0319166001600160a01b0392831617908190556040805163f7c618c160e01b81529051600093929092169163f7c618c191600480820192602092909190829003018186803b15801561225457600080fd5b505afa158015612268573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061228c9190613b81565b601580546001600160a01b0319166001600160a01b03838116919091179182905560405163095ea7b360e01b8152929350169063095ea7b3906122eb9073d9e1ce17f2641f24ae83637ab66a2cca9c378b9f9060001990600401613dfa565b602060405180830381600087803b15801561230557600080fd5b505af1158015612319573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061233d9190613c4e565b50604080518082019091526015546001600160a01b0316815273c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2602082015261237e906016906002613aa0565b50506015805460ff60a01b1916600160a01b1790555050565b601454600160a81b900460ff1681565b6002546001600160a01b031633146123d15760405162461bcd60e51b815260040161097090613eae565b6001600160a01b0381166123e457600080fd5b600380546001600160a01b0319166001600160a01b0383161790556040517fafbb66abf8f3b719799940473a4052a3717cdd8e40fb6c8a3faadab316b1a069906111dc908390613dcb565b601554600160a01b900460ff1681565b60006124496112fe565b61245557506000611ad2565b601454600160a81b900460ff161561247c5761246f6126db565b1561247c57506000611ad2565b6000612486610b42565b905060105481111561249c576001915050611ad2565b6124a4613508565b6124b2576000915050611ad2565b60125460ff16156124c7576001915050611ad2565b600f548111156124db576001915050611ad2565b6124e3613b01565b6001546040516339ebf82360e01b81526001600160a01b03909116906339ebf82390612513903090600401613dcb565b6101006040518083038186803b15801561252c57600080fd5b505afa158015612540573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125649190613c95565b90506006546125808260800151426129dd90919063ffffffff16565b111561259157600192505050611ad2565b601154600160009054906101000a90046001600160a01b03166001600160a01b031663112c1f9b6040518163ffffffff1660e01b815260040160206040518083038186803b1580156125e257600080fd5b505afa1580156125f6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061261a9190613d1f565b111561262b57600192505050611ad2565b5060009392505050565b600e546001600160a01b031681565b600061139b612651611a43565b61104a611d9e565b6002546001600160a01b031633148061268a5750612675612910565b6001600160a01b0316336001600160a01b0316145b6126a65760405162461bcd60e51b81526004016109709061402b565b60068190556040517f4aaf232568bff365c53cad69bdb6e83014e79df80216ceba8ee01769723dfd68906111dc9083906140be565b600080600960019054906101000a90046001600160a01b03166001600160a01b031663ebe2b12b6040518163ffffffff1660e01b815260040160206040518083038186803b15801561272c57600080fd5b505afa158015612740573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127649190613d1f565b905042811015612778576001915050611061565b601554600160a01b900460ff161561281357600a546040805163ebe2b12b60e01b815290516000926001600160a01b03169163ebe2b12b916004808301926020929190829003018186803b1580156127cf57600080fd5b505afa1580156127e3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128079190613d1f565b42119250611061915050565b5090565b600b5481565b6001546001600160a01b031681565b6002546001600160a01b031633148061285d5750612848612910565b6001600160a01b0316336001600160a01b0316145b6128795760405162461bcd60e51b81526004016109709061402b565b6009805460ff19166001908117909155546040805163507257cd60e11b815290516001600160a01b039092169163a0e4af9a9160048082019260009290919082900301818387803b1580156128cd57600080fd5b505af11580156128e1573d6000803e3d6000fd5b50506040517f97e963041e952738788b9d4871d854d282065b8f90a464928d6528f2e9a4fd0b925060009150a1565b60015460408051635aa6e67560e01b815290516000926001600160a01b031691635aa6e675916004808301926020929190829003018186803b15801561295557600080fd5b505afa158015612969573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061139b9190613b81565b606090565b60006129d483836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f00000000000081525061358f565b90505b92915050565b60006129d483836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506135c6565b600082612a2e575060006129d7565b82820282848281612a3b57fe5b04146129d45760405162461bcd60e51b815260040161097090613f72565b6000828201838110156129d45760405162461bcd60e51b815260040161097090613ef2565b6000806000612a8b611d9e565b905080841115612b77576000612a9f611a43565b90508015612b4c5760095461010090046001600160a01b031663c32e7202612ad083612acb89876129dd565b6135f2565b6014546040516001600160e01b031960e085901b168152612af8929160ff169060040161411d565b602060405180830381600087803b158015612b1257600080fd5b505af1158015612b26573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b4a9190613c4e565b505b6000612b56611d9e565b9050612b6286826135f2565b9450612b6e86866129dd565b93505050612b84565b8360009250925050612b86565b505b915091565b60095460ff1615612b9b576115a2565b6000612ba5611d9e565b905080156115a057600b546040516321d0683360e11b815273f403c135812408bfbe8713b5a23a04b3d48aae31916343a0d06691610aeb919085906001906004016141a0565b600954604051637050ccd960e01b81526000918291829161010090046001600160a01b031690637050ccd990612c28903090600190600401613ddf565b602060405180830381600087803b158015612c4257600080fd5b505af1158015612c56573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c7a9190613c4e565b506040516370a0823160e01b815260009073d533a949740bb3306d119cc777fa900ba034cd52906370a0823190612cb5903090600401613dcb565b60206040518083038186803b158015612ccd57600080fd5b505afa158015612ce1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d059190613d1f565b6040516370a0823160e01b8152909150600090734e3fbd56cd56c3e72c1403e103b45db9da5b9d2b906370a0823190612d42903090600401613dcb565b60206040518083038186803b158015612d5a57600080fd5b505afa158015612d6e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d929190613d1f565b90506000612db1612710610c31600c5486612a1f90919063ffffffff16565b90508015612e7657612dec73d533a949740bb3306d119cc777fa900ba034cd5273f147b8125d2ef93fb6965db97d6746952a13393483613608565b6040516370a0823160e01b815273d533a949740bb3306d119cc777fa900ba034cd52906370a0823190612e23903090600401613dcb565b60206040518083038186803b158015612e3b57600080fd5b505afa158015612e4f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e739190613d1f565b92505b6000612e93612710610c31600d5486612a1f90919063ffffffff16565b90508015612f5157600e54612ec790734e3fbd56cd56c3e72c1403e103b45db9da5b9d2b906001600160a01b031683613608565b6040516370a0823160e01b8152734e3fbd56cd56c3e72c1403e103b45db9da5b9d2b906370a0823190612efe903090600401613dcb565b60206040518083038186803b158015612f1657600080fd5b505afa158015612f2a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f4e9190613d1f565b92505b601554600160a01b900460ff1615612ff7576015546040516370a0823160e01b81526000916001600160a01b0316906370a0823190612f94903090600401613dcb565b60206040518083038186803b158015612fac57600080fd5b505afa158015612fc0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612fe49190613d1f565b90508015612ff557612ff58161365e565b505b61300184846136f5565b478015613086576014546040805180820182528381526000602082018190529151630b4c7e4d60e01b81526101009093046001600160a01b031692630b4c7e4d92859261305392909190600401613e13565b6000604051808303818588803b15801561306c57600080fd5b505af1158015613080573d6000803e3d6000fd5b50505050505b8815613155576000613096611a43565b9050801561313a5760095461010090046001600160a01b031663c32e72026130be838d6135f2565b6014546040516001600160e01b031960e085901b1681526130e6929160ff169060040161411d565b602060405180830381600087803b15801561310057600080fd5b505af1158015613114573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906131389190613c4e565b505b6000613144611d9e565b90506131508b826135f2565b975050505b600061315f612644565b6001546040516339ebf82360e01b81529192506000916001600160a01b03909116906339ebf82390613195903090600401613dcb565b6101006040518083038186803b1580156131ae57600080fd5b505afa1580156131c2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906131e69190613c95565b60a0015190508082111561322f576131fe82826129dd565b9950600061320a611d9e565b9050806132178c8b612a59565b11156132295761322683612a7e565b50505b5061323c565b61323981836129dd565b98505b50506012805460ff19169055509597949650929450505050565b6001546040516370a0823160e01b81526000916001600160a01b0316906370a0823190613287903090600401613dcb565b60206040518083038186803b15801561329f57600080fd5b505afa1580156132b3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906132d79190613d1f565b905080156115a25760015460035460405163a9059cbb60e01b81526001600160a01b039283169263a9059cbb9261154e929116908590600401613dfa565b600061331f611a43565b905080156133b657600954601454604051636197390160e11b81526101009092046001600160a01b03169163c32e72029161336291859160ff169060040161411d565b602060405180830381600087803b15801561337c57600080fd5b505af1158015613390573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906133b49190613c4e565b505b6040516370a0823160e01b815261345f90839073d533a949740bb3306d119cc777fa900ba034cd52906370a08231906133f3903090600401613dcb565b60206040518083038186803b15801561340b57600080fd5b505afa15801561341f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906134439190613d1f565b73d533a949740bb3306d119cc777fa900ba034cd529190613608565b6040516370a0823160e01b81526115a0908390734e3fbd56cd56c3e72c1403e103b45db9da5b9d2b906370a082319061349c903090600401613dcb565b60206040518083038186803b1580156134b457600080fd5b505afa1580156134c8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906134ec9190613d1f565b734e3fbd56cd56c3e72c1403e103b45db9da5b9d2b9190613608565b600073b5e1cacb567d98faadb60a1fd4820720141f064f6001600160a01b03166334a9e75c6040518163ffffffff1660e01b815260040160206040518083038186803b15801561355757600080fd5b505afa15801561356b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061139b9190613c4e565b600081836135b05760405162461bcd60e51b81526004016109709190613e7b565b5060008385816135bc57fe5b0495945050505050565b600081848411156135ea5760405162461bcd60e51b81526004016109709190613e7b565b505050900390565b600081831061360157816129d4565b5090919050565b610b3d8363a9059cbb60e01b8484604051602401613627929190613dfa565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b0319909316929092179091526138fd565b6040516338ed173960e01b815273d9e1ce17f2641f24ae83637ab66a2cca9c378b9f906338ed17399061369f9084906000906016903090429060040161412d565b600060405180830381600087803b1580156136b957600080fd5b505af11580156136cd573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526115a09190810190613b9d565b67016345785d8a00008111156137795760405163394747c560e01b815273b576491f1e6e5e62f1d8f26062ee822b40b0e0d49063394747c59061374690600190600090869082908190600401613e56565b600060405180830381600087803b15801561376057600080fd5b505af1158015613774573d6000803e3d6000fd5b505050505b67016345785d8a00008211156137fd5760405163394747c560e01b8152738301ae4fc9c624d1d396cbdaa1ed877821d7c5119063394747c5906137ca90600190600090879082908190600401613e56565b600060405180830381600087803b1580156137e457600080fd5b505af11580156137f8573d6000803e3d6000fd5b505050505b6040516370a0823160e01b815260009073c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2906370a0823190613837903090600401613dcb565b60206040518083038186803b15801561384f57600080fd5b505afa158015613863573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906138879190613d1f565b90508015610b3d57604051632e1a7d4d60e01b815273c02aaa39b223fe8d0a0e5c4f27ead9083c756cc290632e1a7d4d906138c69084906004016140be565b600060405180830381600087803b1580156138e057600080fd5b505af11580156138f4573d6000803e3d6000fd5b50505050505050565b6060613952826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b031661398c9092919063ffffffff16565b805190915015610b3d57808060200190518101906139709190613c4e565b610b3d5760405162461bcd60e51b815260040161097090614050565b606061399b84846000856139a3565b949350505050565b60606139ae85613a67565b6139ca5760405162461bcd60e51b815260040161097090613ff4565b60006060866001600160a01b031685876040516139e79190613daf565b60006040518083038185875af1925050503d8060008114613a24576040519150601f19603f3d011682016040523d82523d6000602084013e613a29565b606091505b50915091508115613a3d57915061399b9050565b805115613a4d5780518082602001fd5b8360405162461bcd60e51b81526004016109709190613e7b565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47081811480159061399b575050151592915050565b828054828255906000526020600020908101928215613af5579160200282015b82811115613af557825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190613ac0565b50612813929150613b46565b60405180610100016040528060008152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b5b808211156128135780546001600160a01b0319168155600101613b47565b600060208284031215613b76578081fd5b81356129d481614283565b600060208284031215613b92578081fd5b81516129d481614283565b60006020808385031215613baf578182fd5b825167ffffffffffffffff811115613bc5578283fd5b8301601f81018513613bd5578283fd5b8051613be8613be382614233565b61420c565b8181528381019083850185840285018601891015613c04578687fd5b8694505b83851015613c26578051835260019490940193918501918501613c08565b50979650505050505050565b600060208284031215613c43578081fd5b81356129d481614298565b600060208284031215613c5f578081fd5b81516129d481614298565b60008060408385031215613c7c578081fd5b8235613c8781614298565b946020939093013593505050565b6000610100808385031215613ca8578182fd5b613cb18161420c565b9050825181526020830151602082015260408301516040820152606083015160608201526080830151608082015260a083015160a082015260c083015160c082015260e083015160e08201528091505092915050565b600060208284031215613d18578081fd5b5035919050565b600060208284031215613d30578081fd5b5051919050565b600080600060608486031215613d4b578081fd5b83359250602084013591506040840135613d6481614283565b809150509250925092565b60008060008060808587031215613d84578081fd5b8435935060208501359250604085013591506060850135613da481614298565b939692955090935050565b60008251613dc1818460208701614253565b9190910192915050565b6001600160a01b0391909116815260200190565b6001600160a01b039290921682521515602082015260400190565b6001600160a01b03929092168252602082015260400190565b60608101818460005b6002811015613e3b578151835260209283019290910190600101613e1c565b5050508260408301529392505050565b901515815260200190565b9485526020850193909352604084019190915260608301521515608082015260a00190565b6000602082528251806020840152613e9a816040850160208701614253565b601f01601f19169190910160400192915050565b6020808252600b908201526a085cdd1c985d1959da5cdd60aa1b604082015260600190565b602080825260059082015264085dd85b9d60da1b604082015260600190565b6020808252601b908201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604082015260600190565b6020808252600c908201526b216865616c7468636865636b60a01b604082015260600190565b60208082526009908201526821776974686472617760b81b604082015260600190565b60208082526021908201527f536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f6040820152607760f81b606082015260800190565b602080825260069082015265085d985d5b1d60d21b604082015260600190565b6020808252600790820152662173686172657360c81b604082015260600190565b6020808252601d908201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604082015260600190565b6020808252600b908201526a08585d5d1a1bdc9a5e995960aa1b604082015260600190565b6020808252602a908201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6040820152691bdd081cdd58d8d9595960b21b606082015260800190565b6020808252600a9082015269085c1c9bdd1958dd195960b21b604082015260600190565b90815260200190565b60006040820184835260206040818501528185518084526060860191508287019350845b818110156141105784516001600160a01b0316835293830193918301916001016140eb565b5090979650505050505050565b9182521515602082015260400190565b600060a082018783526020878185015260a0604085015281875480845260c0860191508885528285209350845b8181101561417f5784546001600160a01b03168352600194850194928401920161415a565b50506001600160a01b03969096166060850152505050608001529392505050565b92835260208301919091521515604082015260600190565b9283526020830191909152604082015260600190565b93845260208401929092526040830152606082015260800190565b948552602085019390935260408401919091526060830152608082015260a00190565b60405181810167ffffffffffffffff8111828210171561422b57600080fd5b604052919050565b600067ffffffffffffffff821115614249578081fd5b5060209081020190565b60005b8381101561426e578181015183820152602001614256565b8381111561427d576000848401525b50505050565b6001600160a01b03811681146115a257600080fd5b80151581146115a257600080fdfea26469706673582212203d4307536ab44b913d9d9d491502c5a3ec7a3d593ceb372c0b73139dd488814264736f6c634300060c0033",
}

// StrategyConvexstETHABI is the input ABI used to generate the binding from.
// Deprecated: Use StrategyConvexstETHMetaData.ABI instead.
var StrategyConvexstETHABI = StrategyConvexstETHMetaData.ABI

// Deprecated: Use StrategyConvexstETHMetaData.Sigs instead.
// StrategyConvexstETHFuncSigs maps the 4-byte function signature to its string representation.
var StrategyConvexstETHFuncSigs = StrategyConvexstETHMetaData.Sigs

// StrategyConvexstETHBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StrategyConvexstETHMetaData.Bin instead.
var StrategyConvexstETHBin = StrategyConvexstETHMetaData.Bin

// DeployStrategyConvexstETH deploys a new Ethereum contract, binding an instance of StrategyConvexstETH to it.
func DeployStrategyConvexstETH(auth *bind.TransactOpts, backend bind.ContractBackend, _vault common.Address, _pid *big.Int, _curvePool common.Address, _name string) (common.Address, *types.Transaction, *StrategyConvexstETH, error) {
	parsed, err := StrategyConvexstETHMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StrategyConvexstETHBin), backend, _vault, _pid, _curvePool, _name)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StrategyConvexstETH{StrategyConvexstETHCaller: StrategyConvexstETHCaller{contract: contract}, StrategyConvexstETHTransactor: StrategyConvexstETHTransactor{contract: contract}, StrategyConvexstETHFilterer: StrategyConvexstETHFilterer{contract: contract}}, nil
}

// StrategyConvexstETH is an auto generated Go binding around an Ethereum contract.
type StrategyConvexstETH struct {
	StrategyConvexstETHCaller     // Read-only binding to the contract
	StrategyConvexstETHTransactor // Write-only binding to the contract
	StrategyConvexstETHFilterer   // Log filterer for contract events
}

// StrategyConvexstETHCaller is an auto generated read-only Go binding around an Ethereum contract.
type StrategyConvexstETHCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyConvexstETHTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StrategyConvexstETHTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyConvexstETHFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StrategyConvexstETHFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyConvexstETHSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StrategyConvexstETHSession struct {
	Contract     *StrategyConvexstETH // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// StrategyConvexstETHCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StrategyConvexstETHCallerSession struct {
	Contract *StrategyConvexstETHCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// StrategyConvexstETHTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StrategyConvexstETHTransactorSession struct {
	Contract     *StrategyConvexstETHTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// StrategyConvexstETHRaw is an auto generated low-level Go binding around an Ethereum contract.
type StrategyConvexstETHRaw struct {
	Contract *StrategyConvexstETH // Generic contract binding to access the raw methods on
}

// StrategyConvexstETHCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StrategyConvexstETHCallerRaw struct {
	Contract *StrategyConvexstETHCaller // Generic read-only contract binding to access the raw methods on
}

// StrategyConvexstETHTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StrategyConvexstETHTransactorRaw struct {
	Contract *StrategyConvexstETHTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStrategyConvexstETH creates a new instance of StrategyConvexstETH, bound to a specific deployed contract.
func NewStrategyConvexstETH(address common.Address, backend bind.ContractBackend) (*StrategyConvexstETH, error) {
	contract, err := bindStrategyConvexstETH(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StrategyConvexstETH{StrategyConvexstETHCaller: StrategyConvexstETHCaller{contract: contract}, StrategyConvexstETHTransactor: StrategyConvexstETHTransactor{contract: contract}, StrategyConvexstETHFilterer: StrategyConvexstETHFilterer{contract: contract}}, nil
}

// NewStrategyConvexstETHCaller creates a new read-only instance of StrategyConvexstETH, bound to a specific deployed contract.
func NewStrategyConvexstETHCaller(address common.Address, caller bind.ContractCaller) (*StrategyConvexstETHCaller, error) {
	contract, err := bindStrategyConvexstETH(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StrategyConvexstETHCaller{contract: contract}, nil
}

// NewStrategyConvexstETHTransactor creates a new write-only instance of StrategyConvexstETH, bound to a specific deployed contract.
func NewStrategyConvexstETHTransactor(address common.Address, transactor bind.ContractTransactor) (*StrategyConvexstETHTransactor, error) {
	contract, err := bindStrategyConvexstETH(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StrategyConvexstETHTransactor{contract: contract}, nil
}

// NewStrategyConvexstETHFilterer creates a new log filterer instance of StrategyConvexstETH, bound to a specific deployed contract.
func NewStrategyConvexstETHFilterer(address common.Address, filterer bind.ContractFilterer) (*StrategyConvexstETHFilterer, error) {
	contract, err := bindStrategyConvexstETH(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StrategyConvexstETHFilterer{contract: contract}, nil
}

// bindStrategyConvexstETH binds a generic wrapper to an already deployed contract.
func bindStrategyConvexstETH(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StrategyConvexstETHABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StrategyConvexstETH *StrategyConvexstETHRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StrategyConvexstETH.Contract.StrategyConvexstETHCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StrategyConvexstETH *StrategyConvexstETHRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.StrategyConvexstETHTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StrategyConvexstETH *StrategyConvexstETHRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.StrategyConvexstETHTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StrategyConvexstETH *StrategyConvexstETHCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StrategyConvexstETH.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StrategyConvexstETH *StrategyConvexstETHTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StrategyConvexstETH *StrategyConvexstETHTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.contract.Transact(opts, method, params...)
}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() pure returns(string)
func (_StrategyConvexstETH *StrategyConvexstETHCaller) ApiVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StrategyConvexstETH.contract.Call(opts, &out, "apiVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() pure returns(string)
func (_StrategyConvexstETH *StrategyConvexstETHSession) ApiVersion() (string, error) {
	return _StrategyConvexstETH.Contract.ApiVersion(&_StrategyConvexstETH.CallOpts)
}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() pure returns(string)
func (_StrategyConvexstETH *StrategyConvexstETHCallerSession) ApiVersion() (string, error) {
	return _StrategyConvexstETH.Contract.ApiVersion(&_StrategyConvexstETH.CallOpts)
}

// BalanceOfWant is a free data retrieval call binding the contract method 0xc1a3d44c.
//
// Solidity: function balanceOfWant() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHCaller) BalanceOfWant(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyConvexstETH.contract.Call(opts, &out, "balanceOfWant")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfWant is a free data retrieval call binding the contract method 0xc1a3d44c.
//
// Solidity: function balanceOfWant() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHSession) BalanceOfWant() (*big.Int, error) {
	return _StrategyConvexstETH.Contract.BalanceOfWant(&_StrategyConvexstETH.CallOpts)
}

// BalanceOfWant is a free data retrieval call binding the contract method 0xc1a3d44c.
//
// Solidity: function balanceOfWant() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHCallerSession) BalanceOfWant() (*big.Int, error) {
	return _StrategyConvexstETH.Contract.BalanceOfWant(&_StrategyConvexstETH.CallOpts)
}

// CheckEarmark is a free data retrieval call binding the contract method 0xec2f1050.
//
// Solidity: function checkEarmark() view returns(bool)
func (_StrategyConvexstETH *StrategyConvexstETHCaller) CheckEarmark(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StrategyConvexstETH.contract.Call(opts, &out, "checkEarmark")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckEarmark is a free data retrieval call binding the contract method 0xec2f1050.
//
// Solidity: function checkEarmark() view returns(bool)
func (_StrategyConvexstETH *StrategyConvexstETHSession) CheckEarmark() (bool, error) {
	return _StrategyConvexstETH.Contract.CheckEarmark(&_StrategyConvexstETH.CallOpts)
}

// CheckEarmark is a free data retrieval call binding the contract method 0xec2f1050.
//
// Solidity: function checkEarmark() view returns(bool)
func (_StrategyConvexstETH *StrategyConvexstETHCallerSession) CheckEarmark() (bool, error) {
	return _StrategyConvexstETH.Contract.CheckEarmark(&_StrategyConvexstETH.CallOpts)
}

// ClaimRewards is a free data retrieval call binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() view returns(bool)
func (_StrategyConvexstETH *StrategyConvexstETHCaller) ClaimRewards(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StrategyConvexstETH.contract.Call(opts, &out, "claimRewards")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ClaimRewards is a free data retrieval call binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() view returns(bool)
func (_StrategyConvexstETH *StrategyConvexstETHSession) ClaimRewards() (bool, error) {
	return _StrategyConvexstETH.Contract.ClaimRewards(&_StrategyConvexstETH.CallOpts)
}

// ClaimRewards is a free data retrieval call binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() view returns(bool)
func (_StrategyConvexstETH *StrategyConvexstETHCallerSession) ClaimRewards() (bool, error) {
	return _StrategyConvexstETH.Contract.ClaimRewards(&_StrategyConvexstETH.CallOpts)
}

// ClaimableBalance is a free data retrieval call binding the contract method 0xc4f45423.
//
// Solidity: function claimableBalance() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHCaller) ClaimableBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyConvexstETH.contract.Call(opts, &out, "claimableBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimableBalance is a free data retrieval call binding the contract method 0xc4f45423.
//
// Solidity: function claimableBalance() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHSession) ClaimableBalance() (*big.Int, error) {
	return _StrategyConvexstETH.Contract.ClaimableBalance(&_StrategyConvexstETH.CallOpts)
}

// ClaimableBalance is a free data retrieval call binding the contract method 0xc4f45423.
//
// Solidity: function claimableBalance() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHCallerSession) ClaimableBalance() (*big.Int, error) {
	return _StrategyConvexstETH.Contract.ClaimableBalance(&_StrategyConvexstETH.CallOpts)
}

// ClaimableProfitInUsdt is a free data retrieval call binding the contract method 0x06cfb3c0.
//
// Solidity: function claimableProfitInUsdt() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHCaller) ClaimableProfitInUsdt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyConvexstETH.contract.Call(opts, &out, "claimableProfitInUsdt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimableProfitInUsdt is a free data retrieval call binding the contract method 0x06cfb3c0.
//
// Solidity: function claimableProfitInUsdt() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHSession) ClaimableProfitInUsdt() (*big.Int, error) {
	return _StrategyConvexstETH.Contract.ClaimableProfitInUsdt(&_StrategyConvexstETH.CallOpts)
}

// ClaimableProfitInUsdt is a free data retrieval call binding the contract method 0x06cfb3c0.
//
// Solidity: function claimableProfitInUsdt() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHCallerSession) ClaimableProfitInUsdt() (*big.Int, error) {
	return _StrategyConvexstETH.Contract.ClaimableProfitInUsdt(&_StrategyConvexstETH.CallOpts)
}

// CreditThreshold is a free data retrieval call binding the contract method 0xaa5480cf.
//
// Solidity: function creditThreshold() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHCaller) CreditThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyConvexstETH.contract.Call(opts, &out, "creditThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreditThreshold is a free data retrieval call binding the contract method 0xaa5480cf.
//
// Solidity: function creditThreshold() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHSession) CreditThreshold() (*big.Int, error) {
	return _StrategyConvexstETH.Contract.CreditThreshold(&_StrategyConvexstETH.CallOpts)
}

// CreditThreshold is a free data retrieval call binding the contract method 0xaa5480cf.
//
// Solidity: function creditThreshold() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHCallerSession) CreditThreshold() (*big.Int, error) {
	return _StrategyConvexstETH.Contract.CreditThreshold(&_StrategyConvexstETH.CallOpts)
}

// Curve is a free data retrieval call binding the contract method 0x7165485d.
//
// Solidity: function curve() view returns(address)
func (_StrategyConvexstETH *StrategyConvexstETHCaller) Curve(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyConvexstETH.contract.Call(opts, &out, "curve")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Curve is a free data retrieval call binding the contract method 0x7165485d.
//
// Solidity: function curve() view returns(address)
func (_StrategyConvexstETH *StrategyConvexstETHSession) Curve() (common.Address, error) {
	return _StrategyConvexstETH.Contract.Curve(&_StrategyConvexstETH.CallOpts)
}

// Curve is a free data retrieval call binding the contract method 0x7165485d.
//
// Solidity: function curve() view returns(address)
func (_StrategyConvexstETH *StrategyConvexstETHCallerSession) Curve() (common.Address, error) {
	return _StrategyConvexstETH.Contract.Curve(&_StrategyConvexstETH.CallOpts)
}

// DebtThreshold is a free data retrieval call binding the contract method 0x1d12f28b.
//
// Solidity: function debtThreshold() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHCaller) DebtThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyConvexstETH.contract.Call(opts, &out, "debtThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtThreshold is a free data retrieval call binding the contract method 0x1d12f28b.
//
// Solidity: function debtThreshold() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHSession) DebtThreshold() (*big.Int, error) {
	return _StrategyConvexstETH.Contract.DebtThreshold(&_StrategyConvexstETH.CallOpts)
}

// DebtThreshold is a free data retrieval call binding the contract method 0x1d12f28b.
//
// Solidity: function debtThreshold() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHCallerSession) DebtThreshold() (*big.Int, error) {
	return _StrategyConvexstETH.Contract.DebtThreshold(&_StrategyConvexstETH.CallOpts)
}

// DelegatedAssets is a free data retrieval call binding the contract method 0x8e6350e2.
//
// Solidity: function delegatedAssets() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHCaller) DelegatedAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyConvexstETH.contract.Call(opts, &out, "delegatedAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegatedAssets is a free data retrieval call binding the contract method 0x8e6350e2.
//
// Solidity: function delegatedAssets() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHSession) DelegatedAssets() (*big.Int, error) {
	return _StrategyConvexstETH.Contract.DelegatedAssets(&_StrategyConvexstETH.CallOpts)
}

// DelegatedAssets is a free data retrieval call binding the contract method 0x8e6350e2.
//
// Solidity: function delegatedAssets() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHCallerSession) DelegatedAssets() (*big.Int, error) {
	return _StrategyConvexstETH.Contract.DelegatedAssets(&_StrategyConvexstETH.CallOpts)
}

// DoHealthCheck is a free data retrieval call binding the contract method 0x6718835f.
//
// Solidity: function doHealthCheck() view returns(bool)
func (_StrategyConvexstETH *StrategyConvexstETHCaller) DoHealthCheck(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StrategyConvexstETH.contract.Call(opts, &out, "doHealthCheck")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DoHealthCheck is a free data retrieval call binding the contract method 0x6718835f.
//
// Solidity: function doHealthCheck() view returns(bool)
func (_StrategyConvexstETH *StrategyConvexstETHSession) DoHealthCheck() (bool, error) {
	return _StrategyConvexstETH.Contract.DoHealthCheck(&_StrategyConvexstETH.CallOpts)
}

// DoHealthCheck is a free data retrieval call binding the contract method 0x6718835f.
//
// Solidity: function doHealthCheck() view returns(bool)
func (_StrategyConvexstETH *StrategyConvexstETHCallerSession) DoHealthCheck() (bool, error) {
	return _StrategyConvexstETH.Contract.DoHealthCheck(&_StrategyConvexstETH.CallOpts)
}

// EmergencyExit is a free data retrieval call binding the contract method 0x5641ec03.
//
// Solidity: function emergencyExit() view returns(bool)
func (_StrategyConvexstETH *StrategyConvexstETHCaller) EmergencyExit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StrategyConvexstETH.contract.Call(opts, &out, "emergencyExit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EmergencyExit is a free data retrieval call binding the contract method 0x5641ec03.
//
// Solidity: function emergencyExit() view returns(bool)
func (_StrategyConvexstETH *StrategyConvexstETHSession) EmergencyExit() (bool, error) {
	return _StrategyConvexstETH.Contract.EmergencyExit(&_StrategyConvexstETH.CallOpts)
}

// EmergencyExit is a free data retrieval call binding the contract method 0x5641ec03.
//
// Solidity: function emergencyExit() view returns(bool)
func (_StrategyConvexstETH *StrategyConvexstETHCallerSession) EmergencyExit() (bool, error) {
	return _StrategyConvexstETH.Contract.EmergencyExit(&_StrategyConvexstETH.CallOpts)
}

// EstimatedTotalAssets is a free data retrieval call binding the contract method 0xefbb5cb0.
//
// Solidity: function estimatedTotalAssets() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHCaller) EstimatedTotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyConvexstETH.contract.Call(opts, &out, "estimatedTotalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimatedTotalAssets is a free data retrieval call binding the contract method 0xefbb5cb0.
//
// Solidity: function estimatedTotalAssets() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHSession) EstimatedTotalAssets() (*big.Int, error) {
	return _StrategyConvexstETH.Contract.EstimatedTotalAssets(&_StrategyConvexstETH.CallOpts)
}

// EstimatedTotalAssets is a free data retrieval call binding the contract method 0xefbb5cb0.
//
// Solidity: function estimatedTotalAssets() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHCallerSession) EstimatedTotalAssets() (*big.Int, error) {
	return _StrategyConvexstETH.Contract.EstimatedTotalAssets(&_StrategyConvexstETH.CallOpts)
}

// HarvestProfitMax is a free data retrieval call binding the contract method 0x3b7c6e2f.
//
// Solidity: function harvestProfitMax() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHCaller) HarvestProfitMax(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyConvexstETH.contract.Call(opts, &out, "harvestProfitMax")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HarvestProfitMax is a free data retrieval call binding the contract method 0x3b7c6e2f.
//
// Solidity: function harvestProfitMax() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHSession) HarvestProfitMax() (*big.Int, error) {
	return _StrategyConvexstETH.Contract.HarvestProfitMax(&_StrategyConvexstETH.CallOpts)
}

// HarvestProfitMax is a free data retrieval call binding the contract method 0x3b7c6e2f.
//
// Solidity: function harvestProfitMax() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHCallerSession) HarvestProfitMax() (*big.Int, error) {
	return _StrategyConvexstETH.Contract.HarvestProfitMax(&_StrategyConvexstETH.CallOpts)
}

// HarvestProfitMin is a free data retrieval call binding the contract method 0x5fbeb25f.
//
// Solidity: function harvestProfitMin() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHCaller) HarvestProfitMin(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyConvexstETH.contract.Call(opts, &out, "harvestProfitMin")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HarvestProfitMin is a free data retrieval call binding the contract method 0x5fbeb25f.
//
// Solidity: function harvestProfitMin() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHSession) HarvestProfitMin() (*big.Int, error) {
	return _StrategyConvexstETH.Contract.HarvestProfitMin(&_StrategyConvexstETH.CallOpts)
}

// HarvestProfitMin is a free data retrieval call binding the contract method 0x5fbeb25f.
//
// Solidity: function harvestProfitMin() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHCallerSession) HarvestProfitMin() (*big.Int, error) {
	return _StrategyConvexstETH.Contract.HarvestProfitMin(&_StrategyConvexstETH.CallOpts)
}

// HarvestTrigger is a free data retrieval call binding the contract method 0xed882c2b.
//
// Solidity: function harvestTrigger(uint256 callCostinEth) view returns(bool)
func (_StrategyConvexstETH *StrategyConvexstETHCaller) HarvestTrigger(opts *bind.CallOpts, callCostinEth *big.Int) (bool, error) {
	var out []interface{}
	err := _StrategyConvexstETH.contract.Call(opts, &out, "harvestTrigger", callCostinEth)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HarvestTrigger is a free data retrieval call binding the contract method 0xed882c2b.
//
// Solidity: function harvestTrigger(uint256 callCostinEth) view returns(bool)
func (_StrategyConvexstETH *StrategyConvexstETHSession) HarvestTrigger(callCostinEth *big.Int) (bool, error) {
	return _StrategyConvexstETH.Contract.HarvestTrigger(&_StrategyConvexstETH.CallOpts, callCostinEth)
}

// HarvestTrigger is a free data retrieval call binding the contract method 0xed882c2b.
//
// Solidity: function harvestTrigger(uint256 callCostinEth) view returns(bool)
func (_StrategyConvexstETH *StrategyConvexstETHCallerSession) HarvestTrigger(callCostinEth *big.Int) (bool, error) {
	return _StrategyConvexstETH.Contract.HarvestTrigger(&_StrategyConvexstETH.CallOpts, callCostinEth)
}

// HasRewards is a free data retrieval call binding the contract method 0xecf04e15.
//
// Solidity: function hasRewards() view returns(bool)
func (_StrategyConvexstETH *StrategyConvexstETHCaller) HasRewards(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StrategyConvexstETH.contract.Call(opts, &out, "hasRewards")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRewards is a free data retrieval call binding the contract method 0xecf04e15.
//
// Solidity: function hasRewards() view returns(bool)
func (_StrategyConvexstETH *StrategyConvexstETHSession) HasRewards() (bool, error) {
	return _StrategyConvexstETH.Contract.HasRewards(&_StrategyConvexstETH.CallOpts)
}

// HasRewards is a free data retrieval call binding the contract method 0xecf04e15.
//
// Solidity: function hasRewards() view returns(bool)
func (_StrategyConvexstETH *StrategyConvexstETHCallerSession) HasRewards() (bool, error) {
	return _StrategyConvexstETH.Contract.HasRewards(&_StrategyConvexstETH.CallOpts)
}

// HealthCheck is a free data retrieval call binding the contract method 0xb252720b.
//
// Solidity: function healthCheck() view returns(address)
func (_StrategyConvexstETH *StrategyConvexstETHCaller) HealthCheck(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyConvexstETH.contract.Call(opts, &out, "healthCheck")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HealthCheck is a free data retrieval call binding the contract method 0xb252720b.
//
// Solidity: function healthCheck() view returns(address)
func (_StrategyConvexstETH *StrategyConvexstETHSession) HealthCheck() (common.Address, error) {
	return _StrategyConvexstETH.Contract.HealthCheck(&_StrategyConvexstETH.CallOpts)
}

// HealthCheck is a free data retrieval call binding the contract method 0xb252720b.
//
// Solidity: function healthCheck() view returns(address)
func (_StrategyConvexstETH *StrategyConvexstETHCallerSession) HealthCheck() (common.Address, error) {
	return _StrategyConvexstETH.Contract.HealthCheck(&_StrategyConvexstETH.CallOpts)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_StrategyConvexstETH *StrategyConvexstETHCaller) IsActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StrategyConvexstETH.contract.Call(opts, &out, "isActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_StrategyConvexstETH *StrategyConvexstETHSession) IsActive() (bool, error) {
	return _StrategyConvexstETH.Contract.IsActive(&_StrategyConvexstETH.CallOpts)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_StrategyConvexstETH *StrategyConvexstETHCallerSession) IsActive() (bool, error) {
	return _StrategyConvexstETH.Contract.IsActive(&_StrategyConvexstETH.CallOpts)
}

// KeepCRV is a free data retrieval call binding the contract method 0x7fef901a.
//
// Solidity: function keepCRV() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHCaller) KeepCRV(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyConvexstETH.contract.Call(opts, &out, "keepCRV")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KeepCRV is a free data retrieval call binding the contract method 0x7fef901a.
//
// Solidity: function keepCRV() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHSession) KeepCRV() (*big.Int, error) {
	return _StrategyConvexstETH.Contract.KeepCRV(&_StrategyConvexstETH.CallOpts)
}

// KeepCRV is a free data retrieval call binding the contract method 0x7fef901a.
//
// Solidity: function keepCRV() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHCallerSession) KeepCRV() (*big.Int, error) {
	return _StrategyConvexstETH.Contract.KeepCRV(&_StrategyConvexstETH.CallOpts)
}

// KeepCrvPercent is a free data retrieval call binding the contract method 0x7fef901a.
//
// Solidity: function keepCrvPercent() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHCaller) KeepCrvPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyConvexstETH.contract.Call(opts, &out, "keepCrvPercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KeepCrvPercent is a free data retrieval call binding the contract method 0x7fef901a.
//
// Solidity: function keepCrvPercent() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHSession) KeepCrvPercent() (*big.Int, error) {
	return _StrategyConvexstETH.Contract.KeepCRV(&_StrategyConvexstETH.CallOpts)
}

// KeepCrvPercent is a free data retrieval call binding the contract method 0x7fef901a.
//
// Solidity: function keepCrvPercent() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHCallerSession) KeepCrvPercent() (*big.Int, error) {
	return _StrategyConvexstETH.Contract.KeepCRV(&_StrategyConvexstETH.CallOpts)
}


// KeepCVX is a free data retrieval call binding the contract method 0x4b31217e.
//
// Solidity: function keepCVX() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHCaller) KeepCVX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyConvexstETH.contract.Call(opts, &out, "keepCVX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KeepCVX is a free data retrieval call binding the contract method 0x4b31217e.
//
// Solidity: function keepCVX() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHSession) KeepCVX() (*big.Int, error) {
	return _StrategyConvexstETH.Contract.KeepCVX(&_StrategyConvexstETH.CallOpts)
}

// KeepCVX is a free data retrieval call binding the contract method 0x4b31217e.
//
// Solidity: function keepCVX() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHCallerSession) KeepCVX() (*big.Int, error) {
	return _StrategyConvexstETH.Contract.KeepCVX(&_StrategyConvexstETH.CallOpts)
}

// KeepCVXDestination is a free data retrieval call binding the contract method 0xef86b23c.
//
// Solidity: function keepCVXDestination() view returns(address)
func (_StrategyConvexstETH *StrategyConvexstETHCaller) KeepCVXDestination(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyConvexstETH.contract.Call(opts, &out, "keepCVXDestination")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KeepCVXDestination is a free data retrieval call binding the contract method 0xef86b23c.
//
// Solidity: function keepCVXDestination() view returns(address)
func (_StrategyConvexstETH *StrategyConvexstETHSession) KeepCVXDestination() (common.Address, error) {
	return _StrategyConvexstETH.Contract.KeepCVXDestination(&_StrategyConvexstETH.CallOpts)
}

// KeepCVXDestination is a free data retrieval call binding the contract method 0xef86b23c.
//
// Solidity: function keepCVXDestination() view returns(address)
func (_StrategyConvexstETH *StrategyConvexstETHCallerSession) KeepCVXDestination() (common.Address, error) {
	return _StrategyConvexstETH.Contract.KeepCVXDestination(&_StrategyConvexstETH.CallOpts)
}

// Keeper is a free data retrieval call binding the contract method 0xaced1661.
//
// Solidity: function keeper() view returns(address)
func (_StrategyConvexstETH *StrategyConvexstETHCaller) Keeper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyConvexstETH.contract.Call(opts, &out, "keeper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Keeper is a free data retrieval call binding the contract method 0xaced1661.
//
// Solidity: function keeper() view returns(address)
func (_StrategyConvexstETH *StrategyConvexstETHSession) Keeper() (common.Address, error) {
	return _StrategyConvexstETH.Contract.Keeper(&_StrategyConvexstETH.CallOpts)
}

// Keeper is a free data retrieval call binding the contract method 0xaced1661.
//
// Solidity: function keeper() view returns(address)
func (_StrategyConvexstETH *StrategyConvexstETHCallerSession) Keeper() (common.Address, error) {
	return _StrategyConvexstETH.Contract.Keeper(&_StrategyConvexstETH.CallOpts)
}

// MaxReportDelay is a free data retrieval call binding the contract method 0x28b7ccf7.
//
// Solidity: function maxReportDelay() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHCaller) MaxReportDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyConvexstETH.contract.Call(opts, &out, "maxReportDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxReportDelay is a free data retrieval call binding the contract method 0x28b7ccf7.
//
// Solidity: function maxReportDelay() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHSession) MaxReportDelay() (*big.Int, error) {
	return _StrategyConvexstETH.Contract.MaxReportDelay(&_StrategyConvexstETH.CallOpts)
}

// MaxReportDelay is a free data retrieval call binding the contract method 0x28b7ccf7.
//
// Solidity: function maxReportDelay() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHCallerSession) MaxReportDelay() (*big.Int, error) {
	return _StrategyConvexstETH.Contract.MaxReportDelay(&_StrategyConvexstETH.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StrategyConvexstETH *StrategyConvexstETHCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StrategyConvexstETH.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StrategyConvexstETH *StrategyConvexstETHSession) Name() (string, error) {
	return _StrategyConvexstETH.Contract.Name(&_StrategyConvexstETH.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StrategyConvexstETH *StrategyConvexstETHCallerSession) Name() (string, error) {
	return _StrategyConvexstETH.Contract.Name(&_StrategyConvexstETH.CallOpts)
}

// NeedsEarmarkReward is a free data retrieval call binding the contract method 0xf09338df.
//
// Solidity: function needsEarmarkReward() view returns(bool needsEarmark)
func (_StrategyConvexstETH *StrategyConvexstETHCaller) NeedsEarmarkReward(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StrategyConvexstETH.contract.Call(opts, &out, "needsEarmarkReward")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// NeedsEarmarkReward is a free data retrieval call binding the contract method 0xf09338df.
//
// Solidity: function needsEarmarkReward() view returns(bool needsEarmark)
func (_StrategyConvexstETH *StrategyConvexstETHSession) NeedsEarmarkReward() (bool, error) {
	return _StrategyConvexstETH.Contract.NeedsEarmarkReward(&_StrategyConvexstETH.CallOpts)
}

// NeedsEarmarkReward is a free data retrieval call binding the contract method 0xf09338df.
//
// Solidity: function needsEarmarkReward() view returns(bool needsEarmark)
func (_StrategyConvexstETH *StrategyConvexstETHCallerSession) NeedsEarmarkReward() (bool, error) {
	return _StrategyConvexstETH.Contract.NeedsEarmarkReward(&_StrategyConvexstETH.CallOpts)
}

// Pid is a free data retrieval call binding the contract method 0xf1068454.
//
// Solidity: function pid() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHCaller) Pid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyConvexstETH.contract.Call(opts, &out, "pid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Pid is a free data retrieval call binding the contract method 0xf1068454.
//
// Solidity: function pid() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHSession) Pid() (*big.Int, error) {
	return _StrategyConvexstETH.Contract.Pid(&_StrategyConvexstETH.CallOpts)
}

// Pid is a free data retrieval call binding the contract method 0xf1068454.
//
// Solidity: function pid() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHCallerSession) Pid() (*big.Int, error) {
	return _StrategyConvexstETH.Contract.Pid(&_StrategyConvexstETH.CallOpts)
}

// ProfitFactor is a free data retrieval call binding the contract method 0x8cdfe166.
//
// Solidity: function profitFactor() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHCaller) ProfitFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyConvexstETH.contract.Call(opts, &out, "profitFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProfitFactor is a free data retrieval call binding the contract method 0x8cdfe166.
//
// Solidity: function profitFactor() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHSession) ProfitFactor() (*big.Int, error) {
	return _StrategyConvexstETH.Contract.ProfitFactor(&_StrategyConvexstETH.CallOpts)
}

// ProfitFactor is a free data retrieval call binding the contract method 0x8cdfe166.
//
// Solidity: function profitFactor() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHCallerSession) ProfitFactor() (*big.Int, error) {
	return _StrategyConvexstETH.Contract.ProfitFactor(&_StrategyConvexstETH.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x9ec5a894.
//
// Solidity: function rewards() view returns(address)
func (_StrategyConvexstETH *StrategyConvexstETHCaller) Rewards(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyConvexstETH.contract.Call(opts, &out, "rewards")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0x9ec5a894.
//
// Solidity: function rewards() view returns(address)
func (_StrategyConvexstETH *StrategyConvexstETHSession) Rewards() (common.Address, error) {
	return _StrategyConvexstETH.Contract.Rewards(&_StrategyConvexstETH.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x9ec5a894.
//
// Solidity: function rewards() view returns(address)
func (_StrategyConvexstETH *StrategyConvexstETHCallerSession) Rewards() (common.Address, error) {
	return _StrategyConvexstETH.Contract.Rewards(&_StrategyConvexstETH.CallOpts)
}

// RewardsContract is a free data retrieval call binding the contract method 0x220cce97.
//
// Solidity: function rewardsContract() view returns(address)
func (_StrategyConvexstETH *StrategyConvexstETHCaller) RewardsContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyConvexstETH.contract.Call(opts, &out, "rewardsContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsContract is a free data retrieval call binding the contract method 0x220cce97.
//
// Solidity: function rewardsContract() view returns(address)
func (_StrategyConvexstETH *StrategyConvexstETHSession) RewardsContract() (common.Address, error) {
	return _StrategyConvexstETH.Contract.RewardsContract(&_StrategyConvexstETH.CallOpts)
}

// RewardsContract is a free data retrieval call binding the contract method 0x220cce97.
//
// Solidity: function rewardsContract() view returns(address)
func (_StrategyConvexstETH *StrategyConvexstETHCallerSession) RewardsContract() (common.Address, error) {
	return _StrategyConvexstETH.Contract.RewardsContract(&_StrategyConvexstETH.CallOpts)
}

// RewardsToken is a free data retrieval call binding the contract method 0xd1af0c7d.
//
// Solidity: function rewardsToken() view returns(address)
func (_StrategyConvexstETH *StrategyConvexstETHCaller) RewardsToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyConvexstETH.contract.Call(opts, &out, "rewardsToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsToken is a free data retrieval call binding the contract method 0xd1af0c7d.
//
// Solidity: function rewardsToken() view returns(address)
func (_StrategyConvexstETH *StrategyConvexstETHSession) RewardsToken() (common.Address, error) {
	return _StrategyConvexstETH.Contract.RewardsToken(&_StrategyConvexstETH.CallOpts)
}

// RewardsToken is a free data retrieval call binding the contract method 0xd1af0c7d.
//
// Solidity: function rewardsToken() view returns(address)
func (_StrategyConvexstETH *StrategyConvexstETHCallerSession) RewardsToken() (common.Address, error) {
	return _StrategyConvexstETH.Contract.RewardsToken(&_StrategyConvexstETH.CallOpts)
}

// StakedBalance is a free data retrieval call binding the contract method 0x5b9f0016.
//
// Solidity: function stakedBalance() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHCaller) StakedBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyConvexstETH.contract.Call(opts, &out, "stakedBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakedBalance is a free data retrieval call binding the contract method 0x5b9f0016.
//
// Solidity: function stakedBalance() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHSession) StakedBalance() (*big.Int, error) {
	return _StrategyConvexstETH.Contract.StakedBalance(&_StrategyConvexstETH.CallOpts)
}

// StakedBalance is a free data retrieval call binding the contract method 0x5b9f0016.
//
// Solidity: function stakedBalance() view returns(uint256)
func (_StrategyConvexstETH *StrategyConvexstETHCallerSession) StakedBalance() (*big.Int, error) {
	return _StrategyConvexstETH.Contract.StakedBalance(&_StrategyConvexstETH.CallOpts)
}

// Strategist is a free data retrieval call binding the contract method 0x1fe4a686.
//
// Solidity: function strategist() view returns(address)
func (_StrategyConvexstETH *StrategyConvexstETHCaller) Strategist(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyConvexstETH.contract.Call(opts, &out, "strategist")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Strategist is a free data retrieval call binding the contract method 0x1fe4a686.
//
// Solidity: function strategist() view returns(address)
func (_StrategyConvexstETH *StrategyConvexstETHSession) Strategist() (common.Address, error) {
	return _StrategyConvexstETH.Contract.Strategist(&_StrategyConvexstETH.CallOpts)
}

// Strategist is a free data retrieval call binding the contract method 0x1fe4a686.
//
// Solidity: function strategist() view returns(address)
func (_StrategyConvexstETH *StrategyConvexstETHCallerSession) Strategist() (common.Address, error) {
	return _StrategyConvexstETH.Contract.Strategist(&_StrategyConvexstETH.CallOpts)
}

// TendTrigger is a free data retrieval call binding the contract method 0x650d1880.
//
// Solidity: function tendTrigger(uint256 callCost) view returns(bool)
func (_StrategyConvexstETH *StrategyConvexstETHCaller) TendTrigger(opts *bind.CallOpts, callCost *big.Int) (bool, error) {
	var out []interface{}
	err := _StrategyConvexstETH.contract.Call(opts, &out, "tendTrigger", callCost)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TendTrigger is a free data retrieval call binding the contract method 0x650d1880.
//
// Solidity: function tendTrigger(uint256 callCost) view returns(bool)
func (_StrategyConvexstETH *StrategyConvexstETHSession) TendTrigger(callCost *big.Int) (bool, error) {
	return _StrategyConvexstETH.Contract.TendTrigger(&_StrategyConvexstETH.CallOpts, callCost)
}

// TendTrigger is a free data retrieval call binding the contract method 0x650d1880.
//
// Solidity: function tendTrigger(uint256 callCost) view returns(bool)
func (_StrategyConvexstETH *StrategyConvexstETHCallerSession) TendTrigger(callCost *big.Int) (bool, error) {
	return _StrategyConvexstETH.Contract.TendTrigger(&_StrategyConvexstETH.CallOpts, callCost)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_StrategyConvexstETH *StrategyConvexstETHCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyConvexstETH.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_StrategyConvexstETH *StrategyConvexstETHSession) Vault() (common.Address, error) {
	return _StrategyConvexstETH.Contract.Vault(&_StrategyConvexstETH.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_StrategyConvexstETH *StrategyConvexstETHCallerSession) Vault() (common.Address, error) {
	return _StrategyConvexstETH.Contract.Vault(&_StrategyConvexstETH.CallOpts)
}

// VirtualRewardsPool is a free data retrieval call binding the contract method 0xba28e59c.
//
// Solidity: function virtualRewardsPool() view returns(address)
func (_StrategyConvexstETH *StrategyConvexstETHCaller) VirtualRewardsPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyConvexstETH.contract.Call(opts, &out, "virtualRewardsPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VirtualRewardsPool is a free data retrieval call binding the contract method 0xba28e59c.
//
// Solidity: function virtualRewardsPool() view returns(address)
func (_StrategyConvexstETH *StrategyConvexstETHSession) VirtualRewardsPool() (common.Address, error) {
	return _StrategyConvexstETH.Contract.VirtualRewardsPool(&_StrategyConvexstETH.CallOpts)
}

// VirtualRewardsPool is a free data retrieval call binding the contract method 0xba28e59c.
//
// Solidity: function virtualRewardsPool() view returns(address)
func (_StrategyConvexstETH *StrategyConvexstETHCallerSession) VirtualRewardsPool() (common.Address, error) {
	return _StrategyConvexstETH.Contract.VirtualRewardsPool(&_StrategyConvexstETH.CallOpts)
}

// Want is a free data retrieval call binding the contract method 0x1f1fcd51.
//
// Solidity: function want() view returns(address)
func (_StrategyConvexstETH *StrategyConvexstETHCaller) Want(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyConvexstETH.contract.Call(opts, &out, "want")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Want is a free data retrieval call binding the contract method 0x1f1fcd51.
//
// Solidity: function want() view returns(address)
func (_StrategyConvexstETH *StrategyConvexstETHSession) Want() (common.Address, error) {
	return _StrategyConvexstETH.Contract.Want(&_StrategyConvexstETH.CallOpts)
}

// Want is a free data retrieval call binding the contract method 0x1f1fcd51.
//
// Solidity: function want() view returns(address)
func (_StrategyConvexstETH *StrategyConvexstETHCallerSession) Want() (common.Address, error) {
	return _StrategyConvexstETH.Contract.Want(&_StrategyConvexstETH.CallOpts)
}

// Harvest is a paid mutator transaction binding the contract method 0x4641257d.
//
// Solidity: function harvest() returns()
func (_StrategyConvexstETH *StrategyConvexstETHTransactor) Harvest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyConvexstETH.contract.Transact(opts, "harvest")
}

// Harvest is a paid mutator transaction binding the contract method 0x4641257d.
//
// Solidity: function harvest() returns()
func (_StrategyConvexstETH *StrategyConvexstETHSession) Harvest() (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.Harvest(&_StrategyConvexstETH.TransactOpts)
}

// Harvest is a paid mutator transaction binding the contract method 0x4641257d.
//
// Solidity: function harvest() returns()
func (_StrategyConvexstETH *StrategyConvexstETHTransactorSession) Harvest() (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.Harvest(&_StrategyConvexstETH.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _newStrategy) returns()
func (_StrategyConvexstETH *StrategyConvexstETHTransactor) Migrate(opts *bind.TransactOpts, _newStrategy common.Address) (*types.Transaction, error) {
	return _StrategyConvexstETH.contract.Transact(opts, "migrate", _newStrategy)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _newStrategy) returns()
func (_StrategyConvexstETH *StrategyConvexstETHSession) Migrate(_newStrategy common.Address) (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.Migrate(&_StrategyConvexstETH.TransactOpts, _newStrategy)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _newStrategy) returns()
func (_StrategyConvexstETH *StrategyConvexstETHTransactorSession) Migrate(_newStrategy common.Address) (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.Migrate(&_StrategyConvexstETH.TransactOpts, _newStrategy)
}

// SetClaimRewards is a paid mutator transaction binding the contract method 0xa98f9296.
//
// Solidity: function setClaimRewards(bool _claimRewards) returns()
func (_StrategyConvexstETH *StrategyConvexstETHTransactor) SetClaimRewards(opts *bind.TransactOpts, _claimRewards bool) (*types.Transaction, error) {
	return _StrategyConvexstETH.contract.Transact(opts, "setClaimRewards", _claimRewards)
}

// SetClaimRewards is a paid mutator transaction binding the contract method 0xa98f9296.
//
// Solidity: function setClaimRewards(bool _claimRewards) returns()
func (_StrategyConvexstETH *StrategyConvexstETHSession) SetClaimRewards(_claimRewards bool) (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.SetClaimRewards(&_StrategyConvexstETH.TransactOpts, _claimRewards)
}

// SetClaimRewards is a paid mutator transaction binding the contract method 0xa98f9296.
//
// Solidity: function setClaimRewards(bool _claimRewards) returns()
func (_StrategyConvexstETH *StrategyConvexstETHTransactorSession) SetClaimRewards(_claimRewards bool) (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.SetClaimRewards(&_StrategyConvexstETH.TransactOpts, _claimRewards)
}

// SetDebtThreshold is a paid mutator transaction binding the contract method 0x0f969b87.
//
// Solidity: function setDebtThreshold(uint256 _debtThreshold) returns()
func (_StrategyConvexstETH *StrategyConvexstETHTransactor) SetDebtThreshold(opts *bind.TransactOpts, _debtThreshold *big.Int) (*types.Transaction, error) {
	return _StrategyConvexstETH.contract.Transact(opts, "setDebtThreshold", _debtThreshold)
}

// SetDebtThreshold is a paid mutator transaction binding the contract method 0x0f969b87.
//
// Solidity: function setDebtThreshold(uint256 _debtThreshold) returns()
func (_StrategyConvexstETH *StrategyConvexstETHSession) SetDebtThreshold(_debtThreshold *big.Int) (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.SetDebtThreshold(&_StrategyConvexstETH.TransactOpts, _debtThreshold)
}

// SetDebtThreshold is a paid mutator transaction binding the contract method 0x0f969b87.
//
// Solidity: function setDebtThreshold(uint256 _debtThreshold) returns()
func (_StrategyConvexstETH *StrategyConvexstETHTransactorSession) SetDebtThreshold(_debtThreshold *big.Int) (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.SetDebtThreshold(&_StrategyConvexstETH.TransactOpts, _debtThreshold)
}

// SetDoHealthCheck is a paid mutator transaction binding the contract method 0xac00ff26.
//
// Solidity: function setDoHealthCheck(bool _doHealthCheck) returns()
func (_StrategyConvexstETH *StrategyConvexstETHTransactor) SetDoHealthCheck(opts *bind.TransactOpts, _doHealthCheck bool) (*types.Transaction, error) {
	return _StrategyConvexstETH.contract.Transact(opts, "setDoHealthCheck", _doHealthCheck)
}

// SetDoHealthCheck is a paid mutator transaction binding the contract method 0xac00ff26.
//
// Solidity: function setDoHealthCheck(bool _doHealthCheck) returns()
func (_StrategyConvexstETH *StrategyConvexstETHSession) SetDoHealthCheck(_doHealthCheck bool) (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.SetDoHealthCheck(&_StrategyConvexstETH.TransactOpts, _doHealthCheck)
}

// SetDoHealthCheck is a paid mutator transaction binding the contract method 0xac00ff26.
//
// Solidity: function setDoHealthCheck(bool _doHealthCheck) returns()
func (_StrategyConvexstETH *StrategyConvexstETHTransactorSession) SetDoHealthCheck(_doHealthCheck bool) (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.SetDoHealthCheck(&_StrategyConvexstETH.TransactOpts, _doHealthCheck)
}

// SetEmergencyExit is a paid mutator transaction binding the contract method 0xfcf2d0ad.
//
// Solidity: function setEmergencyExit() returns()
func (_StrategyConvexstETH *StrategyConvexstETHTransactor) SetEmergencyExit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyConvexstETH.contract.Transact(opts, "setEmergencyExit")
}

// SetEmergencyExit is a paid mutator transaction binding the contract method 0xfcf2d0ad.
//
// Solidity: function setEmergencyExit() returns()
func (_StrategyConvexstETH *StrategyConvexstETHSession) SetEmergencyExit() (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.SetEmergencyExit(&_StrategyConvexstETH.TransactOpts)
}

// SetEmergencyExit is a paid mutator transaction binding the contract method 0xfcf2d0ad.
//
// Solidity: function setEmergencyExit() returns()
func (_StrategyConvexstETH *StrategyConvexstETHTransactorSession) SetEmergencyExit() (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.SetEmergencyExit(&_StrategyConvexstETH.TransactOpts)
}

// SetForceHarvestTriggerOnce is a paid mutator transaction binding the contract method 0x0ada4dab.
//
// Solidity: function setForceHarvestTriggerOnce(bool _forceHarvestTriggerOnce) returns()
func (_StrategyConvexstETH *StrategyConvexstETHTransactor) SetForceHarvestTriggerOnce(opts *bind.TransactOpts, _forceHarvestTriggerOnce bool) (*types.Transaction, error) {
	return _StrategyConvexstETH.contract.Transact(opts, "setForceHarvestTriggerOnce", _forceHarvestTriggerOnce)
}

// SetForceHarvestTriggerOnce is a paid mutator transaction binding the contract method 0x0ada4dab.
//
// Solidity: function setForceHarvestTriggerOnce(bool _forceHarvestTriggerOnce) returns()
func (_StrategyConvexstETH *StrategyConvexstETHSession) SetForceHarvestTriggerOnce(_forceHarvestTriggerOnce bool) (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.SetForceHarvestTriggerOnce(&_StrategyConvexstETH.TransactOpts, _forceHarvestTriggerOnce)
}

// SetForceHarvestTriggerOnce is a paid mutator transaction binding the contract method 0x0ada4dab.
//
// Solidity: function setForceHarvestTriggerOnce(bool _forceHarvestTriggerOnce) returns()
func (_StrategyConvexstETH *StrategyConvexstETHTransactorSession) SetForceHarvestTriggerOnce(_forceHarvestTriggerOnce bool) (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.SetForceHarvestTriggerOnce(&_StrategyConvexstETH.TransactOpts, _forceHarvestTriggerOnce)
}

// SetHarvestTriggerParams is a paid mutator transaction binding the contract method 0xb4d48fd4.
//
// Solidity: function setHarvestTriggerParams(uint256 _harvestProfitMin, uint256 _harvestProfitMax, uint256 _creditThreshold, bool _checkEarmark) returns()
func (_StrategyConvexstETH *StrategyConvexstETHTransactor) SetHarvestTriggerParams(opts *bind.TransactOpts, _harvestProfitMin *big.Int, _harvestProfitMax *big.Int, _creditThreshold *big.Int, _checkEarmark bool) (*types.Transaction, error) {
	return _StrategyConvexstETH.contract.Transact(opts, "setHarvestTriggerParams", _harvestProfitMin, _harvestProfitMax, _creditThreshold, _checkEarmark)
}

// SetHarvestTriggerParams is a paid mutator transaction binding the contract method 0xb4d48fd4.
//
// Solidity: function setHarvestTriggerParams(uint256 _harvestProfitMin, uint256 _harvestProfitMax, uint256 _creditThreshold, bool _checkEarmark) returns()
func (_StrategyConvexstETH *StrategyConvexstETHSession) SetHarvestTriggerParams(_harvestProfitMin *big.Int, _harvestProfitMax *big.Int, _creditThreshold *big.Int, _checkEarmark bool) (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.SetHarvestTriggerParams(&_StrategyConvexstETH.TransactOpts, _harvestProfitMin, _harvestProfitMax, _creditThreshold, _checkEarmark)
}

// SetHarvestTriggerParams is a paid mutator transaction binding the contract method 0xb4d48fd4.
//
// Solidity: function setHarvestTriggerParams(uint256 _harvestProfitMin, uint256 _harvestProfitMax, uint256 _creditThreshold, bool _checkEarmark) returns()
func (_StrategyConvexstETH *StrategyConvexstETHTransactorSession) SetHarvestTriggerParams(_harvestProfitMin *big.Int, _harvestProfitMax *big.Int, _creditThreshold *big.Int, _checkEarmark bool) (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.SetHarvestTriggerParams(&_StrategyConvexstETH.TransactOpts, _harvestProfitMin, _harvestProfitMax, _creditThreshold, _checkEarmark)
}

// SetHealthCheck is a paid mutator transaction binding the contract method 0x11bc8245.
//
// Solidity: function setHealthCheck(address _healthCheck) returns()
func (_StrategyConvexstETH *StrategyConvexstETHTransactor) SetHealthCheck(opts *bind.TransactOpts, _healthCheck common.Address) (*types.Transaction, error) {
	return _StrategyConvexstETH.contract.Transact(opts, "setHealthCheck", _healthCheck)
}

// SetHealthCheck is a paid mutator transaction binding the contract method 0x11bc8245.
//
// Solidity: function setHealthCheck(address _healthCheck) returns()
func (_StrategyConvexstETH *StrategyConvexstETHSession) SetHealthCheck(_healthCheck common.Address) (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.SetHealthCheck(&_StrategyConvexstETH.TransactOpts, _healthCheck)
}

// SetHealthCheck is a paid mutator transaction binding the contract method 0x11bc8245.
//
// Solidity: function setHealthCheck(address _healthCheck) returns()
func (_StrategyConvexstETH *StrategyConvexstETHTransactorSession) SetHealthCheck(_healthCheck common.Address) (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.SetHealthCheck(&_StrategyConvexstETH.TransactOpts, _healthCheck)
}

// SetKeep is a paid mutator transaction binding the contract method 0x1111fe1c.
//
// Solidity: function setKeep(uint256 _keepCRV, uint256 _keepCVX, address _keepCVXDestination) returns()
func (_StrategyConvexstETH *StrategyConvexstETHTransactor) SetKeep(opts *bind.TransactOpts, _keepCRV *big.Int, _keepCVX *big.Int, _keepCVXDestination common.Address) (*types.Transaction, error) {
	return _StrategyConvexstETH.contract.Transact(opts, "setKeep", _keepCRV, _keepCVX, _keepCVXDestination)
}

// SetKeep is a paid mutator transaction binding the contract method 0x1111fe1c.
//
// Solidity: function setKeep(uint256 _keepCRV, uint256 _keepCVX, address _keepCVXDestination) returns()
func (_StrategyConvexstETH *StrategyConvexstETHSession) SetKeep(_keepCRV *big.Int, _keepCVX *big.Int, _keepCVXDestination common.Address) (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.SetKeep(&_StrategyConvexstETH.TransactOpts, _keepCRV, _keepCVX, _keepCVXDestination)
}

// SetKeep is a paid mutator transaction binding the contract method 0x1111fe1c.
//
// Solidity: function setKeep(uint256 _keepCRV, uint256 _keepCVX, address _keepCVXDestination) returns()
func (_StrategyConvexstETH *StrategyConvexstETHTransactorSession) SetKeep(_keepCRV *big.Int, _keepCVX *big.Int, _keepCVXDestination common.Address) (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.SetKeep(&_StrategyConvexstETH.TransactOpts, _keepCRV, _keepCVX, _keepCVXDestination)
}

// SetKeeper is a paid mutator transaction binding the contract method 0x748747e6.
//
// Solidity: function setKeeper(address _keeper) returns()
func (_StrategyConvexstETH *StrategyConvexstETHTransactor) SetKeeper(opts *bind.TransactOpts, _keeper common.Address) (*types.Transaction, error) {
	return _StrategyConvexstETH.contract.Transact(opts, "setKeeper", _keeper)
}

// SetKeeper is a paid mutator transaction binding the contract method 0x748747e6.
//
// Solidity: function setKeeper(address _keeper) returns()
func (_StrategyConvexstETH *StrategyConvexstETHSession) SetKeeper(_keeper common.Address) (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.SetKeeper(&_StrategyConvexstETH.TransactOpts, _keeper)
}

// SetKeeper is a paid mutator transaction binding the contract method 0x748747e6.
//
// Solidity: function setKeeper(address _keeper) returns()
func (_StrategyConvexstETH *StrategyConvexstETHTransactorSession) SetKeeper(_keeper common.Address) (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.SetKeeper(&_StrategyConvexstETH.TransactOpts, _keeper)
}

// SetMaxReportDelay is a paid mutator transaction binding the contract method 0xf017c92f.
//
// Solidity: function setMaxReportDelay(uint256 _delay) returns()
func (_StrategyConvexstETH *StrategyConvexstETHTransactor) SetMaxReportDelay(opts *bind.TransactOpts, _delay *big.Int) (*types.Transaction, error) {
	return _StrategyConvexstETH.contract.Transact(opts, "setMaxReportDelay", _delay)
}

// SetMaxReportDelay is a paid mutator transaction binding the contract method 0xf017c92f.
//
// Solidity: function setMaxReportDelay(uint256 _delay) returns()
func (_StrategyConvexstETH *StrategyConvexstETHSession) SetMaxReportDelay(_delay *big.Int) (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.SetMaxReportDelay(&_StrategyConvexstETH.TransactOpts, _delay)
}

// SetMaxReportDelay is a paid mutator transaction binding the contract method 0xf017c92f.
//
// Solidity: function setMaxReportDelay(uint256 _delay) returns()
func (_StrategyConvexstETH *StrategyConvexstETHTransactorSession) SetMaxReportDelay(_delay *big.Int) (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.SetMaxReportDelay(&_StrategyConvexstETH.TransactOpts, _delay)
}

// SetProfitFactor is a paid mutator transaction binding the contract method 0x91397ab4.
//
// Solidity: function setProfitFactor(uint256 _profitFactor) returns()
func (_StrategyConvexstETH *StrategyConvexstETHTransactor) SetProfitFactor(opts *bind.TransactOpts, _profitFactor *big.Int) (*types.Transaction, error) {
	return _StrategyConvexstETH.contract.Transact(opts, "setProfitFactor", _profitFactor)
}

// SetProfitFactor is a paid mutator transaction binding the contract method 0x91397ab4.
//
// Solidity: function setProfitFactor(uint256 _profitFactor) returns()
func (_StrategyConvexstETH *StrategyConvexstETHSession) SetProfitFactor(_profitFactor *big.Int) (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.SetProfitFactor(&_StrategyConvexstETH.TransactOpts, _profitFactor)
}

// SetProfitFactor is a paid mutator transaction binding the contract method 0x91397ab4.
//
// Solidity: function setProfitFactor(uint256 _profitFactor) returns()
func (_StrategyConvexstETH *StrategyConvexstETHTransactorSession) SetProfitFactor(_profitFactor *big.Int) (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.SetProfitFactor(&_StrategyConvexstETH.TransactOpts, _profitFactor)
}

// SetRewards is a paid mutator transaction binding the contract method 0xec38a862.
//
// Solidity: function setRewards(address _rewards) returns()
func (_StrategyConvexstETH *StrategyConvexstETHTransactor) SetRewards(opts *bind.TransactOpts, _rewards common.Address) (*types.Transaction, error) {
	return _StrategyConvexstETH.contract.Transact(opts, "setRewards", _rewards)
}

// SetRewards is a paid mutator transaction binding the contract method 0xec38a862.
//
// Solidity: function setRewards(address _rewards) returns()
func (_StrategyConvexstETH *StrategyConvexstETHSession) SetRewards(_rewards common.Address) (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.SetRewards(&_StrategyConvexstETH.TransactOpts, _rewards)
}

// SetRewards is a paid mutator transaction binding the contract method 0xec38a862.
//
// Solidity: function setRewards(address _rewards) returns()
func (_StrategyConvexstETH *StrategyConvexstETHTransactorSession) SetRewards(_rewards common.Address) (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.SetRewards(&_StrategyConvexstETH.TransactOpts, _rewards)
}

// SetStrategist is a paid mutator transaction binding the contract method 0xc7b9d530.
//
// Solidity: function setStrategist(address _strategist) returns()
func (_StrategyConvexstETH *StrategyConvexstETHTransactor) SetStrategist(opts *bind.TransactOpts, _strategist common.Address) (*types.Transaction, error) {
	return _StrategyConvexstETH.contract.Transact(opts, "setStrategist", _strategist)
}

// SetStrategist is a paid mutator transaction binding the contract method 0xc7b9d530.
//
// Solidity: function setStrategist(address _strategist) returns()
func (_StrategyConvexstETH *StrategyConvexstETHSession) SetStrategist(_strategist common.Address) (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.SetStrategist(&_StrategyConvexstETH.TransactOpts, _strategist)
}

// SetStrategist is a paid mutator transaction binding the contract method 0xc7b9d530.
//
// Solidity: function setStrategist(address _strategist) returns()
func (_StrategyConvexstETH *StrategyConvexstETHTransactorSession) SetStrategist(_strategist common.Address) (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.SetStrategist(&_StrategyConvexstETH.TransactOpts, _strategist)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address _token) returns()
func (_StrategyConvexstETH *StrategyConvexstETHTransactor) Sweep(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _StrategyConvexstETH.contract.Transact(opts, "sweep", _token)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address _token) returns()
func (_StrategyConvexstETH *StrategyConvexstETHSession) Sweep(_token common.Address) (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.Sweep(&_StrategyConvexstETH.TransactOpts, _token)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address _token) returns()
func (_StrategyConvexstETH *StrategyConvexstETHTransactorSession) Sweep(_token common.Address) (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.Sweep(&_StrategyConvexstETH.TransactOpts, _token)
}

// Tend is a paid mutator transaction binding the contract method 0x440368a3.
//
// Solidity: function tend() returns()
func (_StrategyConvexstETH *StrategyConvexstETHTransactor) Tend(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyConvexstETH.contract.Transact(opts, "tend")
}

// Tend is a paid mutator transaction binding the contract method 0x440368a3.
//
// Solidity: function tend() returns()
func (_StrategyConvexstETH *StrategyConvexstETHSession) Tend() (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.Tend(&_StrategyConvexstETH.TransactOpts)
}

// Tend is a paid mutator transaction binding the contract method 0x440368a3.
//
// Solidity: function tend() returns()
func (_StrategyConvexstETH *StrategyConvexstETHTransactorSession) Tend() (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.Tend(&_StrategyConvexstETH.TransactOpts)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0xdbad1ab1.
//
// Solidity: function updateRewards(bool _hasRewards, uint256 _rewardsIndex) returns()
func (_StrategyConvexstETH *StrategyConvexstETHTransactor) UpdateRewards(opts *bind.TransactOpts, _hasRewards bool, _rewardsIndex *big.Int) (*types.Transaction, error) {
	return _StrategyConvexstETH.contract.Transact(opts, "updateRewards", _hasRewards, _rewardsIndex)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0xdbad1ab1.
//
// Solidity: function updateRewards(bool _hasRewards, uint256 _rewardsIndex) returns()
func (_StrategyConvexstETH *StrategyConvexstETHSession) UpdateRewards(_hasRewards bool, _rewardsIndex *big.Int) (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.UpdateRewards(&_StrategyConvexstETH.TransactOpts, _hasRewards, _rewardsIndex)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0xdbad1ab1.
//
// Solidity: function updateRewards(bool _hasRewards, uint256 _rewardsIndex) returns()
func (_StrategyConvexstETH *StrategyConvexstETHTransactorSession) UpdateRewards(_hasRewards bool, _rewardsIndex *big.Int) (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.UpdateRewards(&_StrategyConvexstETH.TransactOpts, _hasRewards, _rewardsIndex)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amountNeeded) returns(uint256 _loss)
func (_StrategyConvexstETH *StrategyConvexstETHTransactor) Withdraw(opts *bind.TransactOpts, _amountNeeded *big.Int) (*types.Transaction, error) {
	return _StrategyConvexstETH.contract.Transact(opts, "withdraw", _amountNeeded)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amountNeeded) returns(uint256 _loss)
func (_StrategyConvexstETH *StrategyConvexstETHSession) Withdraw(_amountNeeded *big.Int) (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.Withdraw(&_StrategyConvexstETH.TransactOpts, _amountNeeded)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amountNeeded) returns(uint256 _loss)
func (_StrategyConvexstETH *StrategyConvexstETHTransactorSession) Withdraw(_amountNeeded *big.Int) (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.Withdraw(&_StrategyConvexstETH.TransactOpts, _amountNeeded)
}

// WithdrawToConvexDepositTokens is a paid mutator transaction binding the contract method 0x34659dc5.
//
// Solidity: function withdrawToConvexDepositTokens() returns()
func (_StrategyConvexstETH *StrategyConvexstETHTransactor) WithdrawToConvexDepositTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyConvexstETH.contract.Transact(opts, "withdrawToConvexDepositTokens")
}

// WithdrawToConvexDepositTokens is a paid mutator transaction binding the contract method 0x34659dc5.
//
// Solidity: function withdrawToConvexDepositTokens() returns()
func (_StrategyConvexstETH *StrategyConvexstETHSession) WithdrawToConvexDepositTokens() (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.WithdrawToConvexDepositTokens(&_StrategyConvexstETH.TransactOpts)
}

// WithdrawToConvexDepositTokens is a paid mutator transaction binding the contract method 0x34659dc5.
//
// Solidity: function withdrawToConvexDepositTokens() returns()
func (_StrategyConvexstETH *StrategyConvexstETHTransactorSession) WithdrawToConvexDepositTokens() (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.WithdrawToConvexDepositTokens(&_StrategyConvexstETH.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StrategyConvexstETH *StrategyConvexstETHTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyConvexstETH.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StrategyConvexstETH *StrategyConvexstETHSession) Receive() (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.Receive(&_StrategyConvexstETH.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StrategyConvexstETH *StrategyConvexstETHTransactorSession) Receive() (*types.Transaction, error) {
	return _StrategyConvexstETH.Contract.Receive(&_StrategyConvexstETH.TransactOpts)
}

// StrategyConvexstETHEmergencyExitEnabledIterator is returned from FilterEmergencyExitEnabled and is used to iterate over the raw logs and unpacked data for EmergencyExitEnabled events raised by the StrategyConvexstETH contract.
type StrategyConvexstETHEmergencyExitEnabledIterator struct {
	Event *StrategyConvexstETHEmergencyExitEnabled // Event containing the contract specifics and raw log

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
func (it *StrategyConvexstETHEmergencyExitEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyConvexstETHEmergencyExitEnabled)
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
		it.Event = new(StrategyConvexstETHEmergencyExitEnabled)
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
func (it *StrategyConvexstETHEmergencyExitEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyConvexstETHEmergencyExitEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyConvexstETHEmergencyExitEnabled represents a EmergencyExitEnabled event raised by the StrategyConvexstETH contract.
type StrategyConvexstETHEmergencyExitEnabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyExitEnabled is a free log retrieval operation binding the contract event 0x97e963041e952738788b9d4871d854d282065b8f90a464928d6528f2e9a4fd0b.
//
// Solidity: event EmergencyExitEnabled()
func (_StrategyConvexstETH *StrategyConvexstETHFilterer) FilterEmergencyExitEnabled(opts *bind.FilterOpts) (*StrategyConvexstETHEmergencyExitEnabledIterator, error) {

	logs, sub, err := _StrategyConvexstETH.contract.FilterLogs(opts, "EmergencyExitEnabled")
	if err != nil {
		return nil, err
	}
	return &StrategyConvexstETHEmergencyExitEnabledIterator{contract: _StrategyConvexstETH.contract, event: "EmergencyExitEnabled", logs: logs, sub: sub}, nil
}

// WatchEmergencyExitEnabled is a free log subscription operation binding the contract event 0x97e963041e952738788b9d4871d854d282065b8f90a464928d6528f2e9a4fd0b.
//
// Solidity: event EmergencyExitEnabled()
func (_StrategyConvexstETH *StrategyConvexstETHFilterer) WatchEmergencyExitEnabled(opts *bind.WatchOpts, sink chan<- *StrategyConvexstETHEmergencyExitEnabled) (event.Subscription, error) {

	logs, sub, err := _StrategyConvexstETH.contract.WatchLogs(opts, "EmergencyExitEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyConvexstETHEmergencyExitEnabled)
				if err := _StrategyConvexstETH.contract.UnpackLog(event, "EmergencyExitEnabled", log); err != nil {
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

// ParseEmergencyExitEnabled is a log parse operation binding the contract event 0x97e963041e952738788b9d4871d854d282065b8f90a464928d6528f2e9a4fd0b.
//
// Solidity: event EmergencyExitEnabled()
func (_StrategyConvexstETH *StrategyConvexstETHFilterer) ParseEmergencyExitEnabled(log types.Log) (*StrategyConvexstETHEmergencyExitEnabled, error) {
	event := new(StrategyConvexstETHEmergencyExitEnabled)
	if err := _StrategyConvexstETH.contract.UnpackLog(event, "EmergencyExitEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyConvexstETHHarvestedIterator is returned from FilterHarvested and is used to iterate over the raw logs and unpacked data for Harvested events raised by the StrategyConvexstETH contract.
type StrategyConvexstETHHarvestedIterator struct {
	Event *StrategyConvexstETHHarvested // Event containing the contract specifics and raw log

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
func (it *StrategyConvexstETHHarvestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyConvexstETHHarvested)
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
		it.Event = new(StrategyConvexstETHHarvested)
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
func (it *StrategyConvexstETHHarvestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyConvexstETHHarvestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyConvexstETHHarvested represents a Harvested event raised by the StrategyConvexstETH contract.
type StrategyConvexstETHHarvested struct {
	Profit          *big.Int
	Loss            *big.Int
	DebtPayment     *big.Int
	DebtOutstanding *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterHarvested is a free log retrieval operation binding the contract event 0x4c0f499ffe6befa0ca7c826b0916cf87bea98de658013e76938489368d60d509.
//
// Solidity: event Harvested(uint256 profit, uint256 loss, uint256 debtPayment, uint256 debtOutstanding)
func (_StrategyConvexstETH *StrategyConvexstETHFilterer) FilterHarvested(opts *bind.FilterOpts) (*StrategyConvexstETHHarvestedIterator, error) {

	logs, sub, err := _StrategyConvexstETH.contract.FilterLogs(opts, "Harvested")
	if err != nil {
		return nil, err
	}
	return &StrategyConvexstETHHarvestedIterator{contract: _StrategyConvexstETH.contract, event: "Harvested", logs: logs, sub: sub}, nil
}

// WatchHarvested is a free log subscription operation binding the contract event 0x4c0f499ffe6befa0ca7c826b0916cf87bea98de658013e76938489368d60d509.
//
// Solidity: event Harvested(uint256 profit, uint256 loss, uint256 debtPayment, uint256 debtOutstanding)
func (_StrategyConvexstETH *StrategyConvexstETHFilterer) WatchHarvested(opts *bind.WatchOpts, sink chan<- *StrategyConvexstETHHarvested) (event.Subscription, error) {

	logs, sub, err := _StrategyConvexstETH.contract.WatchLogs(opts, "Harvested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyConvexstETHHarvested)
				if err := _StrategyConvexstETH.contract.UnpackLog(event, "Harvested", log); err != nil {
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

// ParseHarvested is a log parse operation binding the contract event 0x4c0f499ffe6befa0ca7c826b0916cf87bea98de658013e76938489368d60d509.
//
// Solidity: event Harvested(uint256 profit, uint256 loss, uint256 debtPayment, uint256 debtOutstanding)
func (_StrategyConvexstETH *StrategyConvexstETHFilterer) ParseHarvested(log types.Log) (*StrategyConvexstETHHarvested, error) {
	event := new(StrategyConvexstETHHarvested)
	if err := _StrategyConvexstETH.contract.UnpackLog(event, "Harvested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyConvexstETHUpdatedDebtThresholdIterator is returned from FilterUpdatedDebtThreshold and is used to iterate over the raw logs and unpacked data for UpdatedDebtThreshold events raised by the StrategyConvexstETH contract.
type StrategyConvexstETHUpdatedDebtThresholdIterator struct {
	Event *StrategyConvexstETHUpdatedDebtThreshold // Event containing the contract specifics and raw log

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
func (it *StrategyConvexstETHUpdatedDebtThresholdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyConvexstETHUpdatedDebtThreshold)
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
		it.Event = new(StrategyConvexstETHUpdatedDebtThreshold)
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
func (it *StrategyConvexstETHUpdatedDebtThresholdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyConvexstETHUpdatedDebtThresholdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyConvexstETHUpdatedDebtThreshold represents a UpdatedDebtThreshold event raised by the StrategyConvexstETH contract.
type StrategyConvexstETHUpdatedDebtThreshold struct {
	DebtThreshold *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedDebtThreshold is a free log retrieval operation binding the contract event 0xa68ba126373d04c004c5748c300c9fca12bd444b3d4332e261f3bd2bac4a8600.
//
// Solidity: event UpdatedDebtThreshold(uint256 debtThreshold)
func (_StrategyConvexstETH *StrategyConvexstETHFilterer) FilterUpdatedDebtThreshold(opts *bind.FilterOpts) (*StrategyConvexstETHUpdatedDebtThresholdIterator, error) {

	logs, sub, err := _StrategyConvexstETH.contract.FilterLogs(opts, "UpdatedDebtThreshold")
	if err != nil {
		return nil, err
	}
	return &StrategyConvexstETHUpdatedDebtThresholdIterator{contract: _StrategyConvexstETH.contract, event: "UpdatedDebtThreshold", logs: logs, sub: sub}, nil
}

// WatchUpdatedDebtThreshold is a free log subscription operation binding the contract event 0xa68ba126373d04c004c5748c300c9fca12bd444b3d4332e261f3bd2bac4a8600.
//
// Solidity: event UpdatedDebtThreshold(uint256 debtThreshold)
func (_StrategyConvexstETH *StrategyConvexstETHFilterer) WatchUpdatedDebtThreshold(opts *bind.WatchOpts, sink chan<- *StrategyConvexstETHUpdatedDebtThreshold) (event.Subscription, error) {

	logs, sub, err := _StrategyConvexstETH.contract.WatchLogs(opts, "UpdatedDebtThreshold")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyConvexstETHUpdatedDebtThreshold)
				if err := _StrategyConvexstETH.contract.UnpackLog(event, "UpdatedDebtThreshold", log); err != nil {
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

// ParseUpdatedDebtThreshold is a log parse operation binding the contract event 0xa68ba126373d04c004c5748c300c9fca12bd444b3d4332e261f3bd2bac4a8600.
//
// Solidity: event UpdatedDebtThreshold(uint256 debtThreshold)
func (_StrategyConvexstETH *StrategyConvexstETHFilterer) ParseUpdatedDebtThreshold(log types.Log) (*StrategyConvexstETHUpdatedDebtThreshold, error) {
	event := new(StrategyConvexstETHUpdatedDebtThreshold)
	if err := _StrategyConvexstETH.contract.UnpackLog(event, "UpdatedDebtThreshold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyConvexstETHUpdatedKeeperIterator is returned from FilterUpdatedKeeper and is used to iterate over the raw logs and unpacked data for UpdatedKeeper events raised by the StrategyConvexstETH contract.
type StrategyConvexstETHUpdatedKeeperIterator struct {
	Event *StrategyConvexstETHUpdatedKeeper // Event containing the contract specifics and raw log

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
func (it *StrategyConvexstETHUpdatedKeeperIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyConvexstETHUpdatedKeeper)
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
		it.Event = new(StrategyConvexstETHUpdatedKeeper)
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
func (it *StrategyConvexstETHUpdatedKeeperIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyConvexstETHUpdatedKeeperIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyConvexstETHUpdatedKeeper represents a UpdatedKeeper event raised by the StrategyConvexstETH contract.
type StrategyConvexstETHUpdatedKeeper struct {
	NewKeeper common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdatedKeeper is a free log retrieval operation binding the contract event 0x2f202ddb4a2e345f6323ed90f8fc8559d770a7abbbeee84dde8aca3351fe7154.
//
// Solidity: event UpdatedKeeper(address newKeeper)
func (_StrategyConvexstETH *StrategyConvexstETHFilterer) FilterUpdatedKeeper(opts *bind.FilterOpts) (*StrategyConvexstETHUpdatedKeeperIterator, error) {

	logs, sub, err := _StrategyConvexstETH.contract.FilterLogs(opts, "UpdatedKeeper")
	if err != nil {
		return nil, err
	}
	return &StrategyConvexstETHUpdatedKeeperIterator{contract: _StrategyConvexstETH.contract, event: "UpdatedKeeper", logs: logs, sub: sub}, nil
}

// WatchUpdatedKeeper is a free log subscription operation binding the contract event 0x2f202ddb4a2e345f6323ed90f8fc8559d770a7abbbeee84dde8aca3351fe7154.
//
// Solidity: event UpdatedKeeper(address newKeeper)
func (_StrategyConvexstETH *StrategyConvexstETHFilterer) WatchUpdatedKeeper(opts *bind.WatchOpts, sink chan<- *StrategyConvexstETHUpdatedKeeper) (event.Subscription, error) {

	logs, sub, err := _StrategyConvexstETH.contract.WatchLogs(opts, "UpdatedKeeper")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyConvexstETHUpdatedKeeper)
				if err := _StrategyConvexstETH.contract.UnpackLog(event, "UpdatedKeeper", log); err != nil {
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

// ParseUpdatedKeeper is a log parse operation binding the contract event 0x2f202ddb4a2e345f6323ed90f8fc8559d770a7abbbeee84dde8aca3351fe7154.
//
// Solidity: event UpdatedKeeper(address newKeeper)
func (_StrategyConvexstETH *StrategyConvexstETHFilterer) ParseUpdatedKeeper(log types.Log) (*StrategyConvexstETHUpdatedKeeper, error) {
	event := new(StrategyConvexstETHUpdatedKeeper)
	if err := _StrategyConvexstETH.contract.UnpackLog(event, "UpdatedKeeper", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyConvexstETHUpdatedProfitFactorIterator is returned from FilterUpdatedProfitFactor and is used to iterate over the raw logs and unpacked data for UpdatedProfitFactor events raised by the StrategyConvexstETH contract.
type StrategyConvexstETHUpdatedProfitFactorIterator struct {
	Event *StrategyConvexstETHUpdatedProfitFactor // Event containing the contract specifics and raw log

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
func (it *StrategyConvexstETHUpdatedProfitFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyConvexstETHUpdatedProfitFactor)
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
		it.Event = new(StrategyConvexstETHUpdatedProfitFactor)
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
func (it *StrategyConvexstETHUpdatedProfitFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyConvexstETHUpdatedProfitFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyConvexstETHUpdatedProfitFactor represents a UpdatedProfitFactor event raised by the StrategyConvexstETH contract.
type StrategyConvexstETHUpdatedProfitFactor struct {
	ProfitFactor *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdatedProfitFactor is a free log retrieval operation binding the contract event 0xd94596337df4c2f0f44d30a7fc5db1c7bb60d9aca4185ed77c6fd96eb45ec298.
//
// Solidity: event UpdatedProfitFactor(uint256 profitFactor)
func (_StrategyConvexstETH *StrategyConvexstETHFilterer) FilterUpdatedProfitFactor(opts *bind.FilterOpts) (*StrategyConvexstETHUpdatedProfitFactorIterator, error) {

	logs, sub, err := _StrategyConvexstETH.contract.FilterLogs(opts, "UpdatedProfitFactor")
	if err != nil {
		return nil, err
	}
	return &StrategyConvexstETHUpdatedProfitFactorIterator{contract: _StrategyConvexstETH.contract, event: "UpdatedProfitFactor", logs: logs, sub: sub}, nil
}

// WatchUpdatedProfitFactor is a free log subscription operation binding the contract event 0xd94596337df4c2f0f44d30a7fc5db1c7bb60d9aca4185ed77c6fd96eb45ec298.
//
// Solidity: event UpdatedProfitFactor(uint256 profitFactor)
func (_StrategyConvexstETH *StrategyConvexstETHFilterer) WatchUpdatedProfitFactor(opts *bind.WatchOpts, sink chan<- *StrategyConvexstETHUpdatedProfitFactor) (event.Subscription, error) {

	logs, sub, err := _StrategyConvexstETH.contract.WatchLogs(opts, "UpdatedProfitFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyConvexstETHUpdatedProfitFactor)
				if err := _StrategyConvexstETH.contract.UnpackLog(event, "UpdatedProfitFactor", log); err != nil {
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

// ParseUpdatedProfitFactor is a log parse operation binding the contract event 0xd94596337df4c2f0f44d30a7fc5db1c7bb60d9aca4185ed77c6fd96eb45ec298.
//
// Solidity: event UpdatedProfitFactor(uint256 profitFactor)
func (_StrategyConvexstETH *StrategyConvexstETHFilterer) ParseUpdatedProfitFactor(log types.Log) (*StrategyConvexstETHUpdatedProfitFactor, error) {
	event := new(StrategyConvexstETHUpdatedProfitFactor)
	if err := _StrategyConvexstETH.contract.UnpackLog(event, "UpdatedProfitFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyConvexstETHUpdatedReportDelayIterator is returned from FilterUpdatedReportDelay and is used to iterate over the raw logs and unpacked data for UpdatedReportDelay events raised by the StrategyConvexstETH contract.
type StrategyConvexstETHUpdatedReportDelayIterator struct {
	Event *StrategyConvexstETHUpdatedReportDelay // Event containing the contract specifics and raw log

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
func (it *StrategyConvexstETHUpdatedReportDelayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyConvexstETHUpdatedReportDelay)
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
		it.Event = new(StrategyConvexstETHUpdatedReportDelay)
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
func (it *StrategyConvexstETHUpdatedReportDelayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyConvexstETHUpdatedReportDelayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyConvexstETHUpdatedReportDelay represents a UpdatedReportDelay event raised by the StrategyConvexstETH contract.
type StrategyConvexstETHUpdatedReportDelay struct {
	Delay *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdatedReportDelay is a free log retrieval operation binding the contract event 0x4aaf232568bff365c53cad69bdb6e83014e79df80216ceba8ee01769723dfd68.
//
// Solidity: event UpdatedReportDelay(uint256 delay)
func (_StrategyConvexstETH *StrategyConvexstETHFilterer) FilterUpdatedReportDelay(opts *bind.FilterOpts) (*StrategyConvexstETHUpdatedReportDelayIterator, error) {

	logs, sub, err := _StrategyConvexstETH.contract.FilterLogs(opts, "UpdatedReportDelay")
	if err != nil {
		return nil, err
	}
	return &StrategyConvexstETHUpdatedReportDelayIterator{contract: _StrategyConvexstETH.contract, event: "UpdatedReportDelay", logs: logs, sub: sub}, nil
}

// WatchUpdatedReportDelay is a free log subscription operation binding the contract event 0x4aaf232568bff365c53cad69bdb6e83014e79df80216ceba8ee01769723dfd68.
//
// Solidity: event UpdatedReportDelay(uint256 delay)
func (_StrategyConvexstETH *StrategyConvexstETHFilterer) WatchUpdatedReportDelay(opts *bind.WatchOpts, sink chan<- *StrategyConvexstETHUpdatedReportDelay) (event.Subscription, error) {

	logs, sub, err := _StrategyConvexstETH.contract.WatchLogs(opts, "UpdatedReportDelay")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyConvexstETHUpdatedReportDelay)
				if err := _StrategyConvexstETH.contract.UnpackLog(event, "UpdatedReportDelay", log); err != nil {
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

// ParseUpdatedReportDelay is a log parse operation binding the contract event 0x4aaf232568bff365c53cad69bdb6e83014e79df80216ceba8ee01769723dfd68.
//
// Solidity: event UpdatedReportDelay(uint256 delay)
func (_StrategyConvexstETH *StrategyConvexstETHFilterer) ParseUpdatedReportDelay(log types.Log) (*StrategyConvexstETHUpdatedReportDelay, error) {
	event := new(StrategyConvexstETHUpdatedReportDelay)
	if err := _StrategyConvexstETH.contract.UnpackLog(event, "UpdatedReportDelay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyConvexstETHUpdatedRewardsIterator is returned from FilterUpdatedRewards and is used to iterate over the raw logs and unpacked data for UpdatedRewards events raised by the StrategyConvexstETH contract.
type StrategyConvexstETHUpdatedRewardsIterator struct {
	Event *StrategyConvexstETHUpdatedRewards // Event containing the contract specifics and raw log

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
func (it *StrategyConvexstETHUpdatedRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyConvexstETHUpdatedRewards)
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
		it.Event = new(StrategyConvexstETHUpdatedRewards)
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
func (it *StrategyConvexstETHUpdatedRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyConvexstETHUpdatedRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyConvexstETHUpdatedRewards represents a UpdatedRewards event raised by the StrategyConvexstETH contract.
type StrategyConvexstETHUpdatedRewards struct {
	Rewards common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdatedRewards is a free log retrieval operation binding the contract event 0xafbb66abf8f3b719799940473a4052a3717cdd8e40fb6c8a3faadab316b1a069.
//
// Solidity: event UpdatedRewards(address rewards)
func (_StrategyConvexstETH *StrategyConvexstETHFilterer) FilterUpdatedRewards(opts *bind.FilterOpts) (*StrategyConvexstETHUpdatedRewardsIterator, error) {

	logs, sub, err := _StrategyConvexstETH.contract.FilterLogs(opts, "UpdatedRewards")
	if err != nil {
		return nil, err
	}
	return &StrategyConvexstETHUpdatedRewardsIterator{contract: _StrategyConvexstETH.contract, event: "UpdatedRewards", logs: logs, sub: sub}, nil
}

// WatchUpdatedRewards is a free log subscription operation binding the contract event 0xafbb66abf8f3b719799940473a4052a3717cdd8e40fb6c8a3faadab316b1a069.
//
// Solidity: event UpdatedRewards(address rewards)
func (_StrategyConvexstETH *StrategyConvexstETHFilterer) WatchUpdatedRewards(opts *bind.WatchOpts, sink chan<- *StrategyConvexstETHUpdatedRewards) (event.Subscription, error) {

	logs, sub, err := _StrategyConvexstETH.contract.WatchLogs(opts, "UpdatedRewards")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyConvexstETHUpdatedRewards)
				if err := _StrategyConvexstETH.contract.UnpackLog(event, "UpdatedRewards", log); err != nil {
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

// ParseUpdatedRewards is a log parse operation binding the contract event 0xafbb66abf8f3b719799940473a4052a3717cdd8e40fb6c8a3faadab316b1a069.
//
// Solidity: event UpdatedRewards(address rewards)
func (_StrategyConvexstETH *StrategyConvexstETHFilterer) ParseUpdatedRewards(log types.Log) (*StrategyConvexstETHUpdatedRewards, error) {
	event := new(StrategyConvexstETHUpdatedRewards)
	if err := _StrategyConvexstETH.contract.UnpackLog(event, "UpdatedRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyConvexstETHUpdatedStrategistIterator is returned from FilterUpdatedStrategist and is used to iterate over the raw logs and unpacked data for UpdatedStrategist events raised by the StrategyConvexstETH contract.
type StrategyConvexstETHUpdatedStrategistIterator struct {
	Event *StrategyConvexstETHUpdatedStrategist // Event containing the contract specifics and raw log

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
func (it *StrategyConvexstETHUpdatedStrategistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyConvexstETHUpdatedStrategist)
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
		it.Event = new(StrategyConvexstETHUpdatedStrategist)
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
func (it *StrategyConvexstETHUpdatedStrategistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyConvexstETHUpdatedStrategistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyConvexstETHUpdatedStrategist represents a UpdatedStrategist event raised by the StrategyConvexstETH contract.
type StrategyConvexstETHUpdatedStrategist struct {
	NewStrategist common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedStrategist is a free log retrieval operation binding the contract event 0x352ececae6d7d1e6d26bcf2c549dfd55be1637e9b22dc0cf3b71ddb36097a6b4.
//
// Solidity: event UpdatedStrategist(address newStrategist)
func (_StrategyConvexstETH *StrategyConvexstETHFilterer) FilterUpdatedStrategist(opts *bind.FilterOpts) (*StrategyConvexstETHUpdatedStrategistIterator, error) {

	logs, sub, err := _StrategyConvexstETH.contract.FilterLogs(opts, "UpdatedStrategist")
	if err != nil {
		return nil, err
	}
	return &StrategyConvexstETHUpdatedStrategistIterator{contract: _StrategyConvexstETH.contract, event: "UpdatedStrategist", logs: logs, sub: sub}, nil
}

// WatchUpdatedStrategist is a free log subscription operation binding the contract event 0x352ececae6d7d1e6d26bcf2c549dfd55be1637e9b22dc0cf3b71ddb36097a6b4.
//
// Solidity: event UpdatedStrategist(address newStrategist)
func (_StrategyConvexstETH *StrategyConvexstETHFilterer) WatchUpdatedStrategist(opts *bind.WatchOpts, sink chan<- *StrategyConvexstETHUpdatedStrategist) (event.Subscription, error) {

	logs, sub, err := _StrategyConvexstETH.contract.WatchLogs(opts, "UpdatedStrategist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyConvexstETHUpdatedStrategist)
				if err := _StrategyConvexstETH.contract.UnpackLog(event, "UpdatedStrategist", log); err != nil {
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

// ParseUpdatedStrategist is a log parse operation binding the contract event 0x352ececae6d7d1e6d26bcf2c549dfd55be1637e9b22dc0cf3b71ddb36097a6b4.
//
// Solidity: event UpdatedStrategist(address newStrategist)
func (_StrategyConvexstETH *StrategyConvexstETHFilterer) ParseUpdatedStrategist(log types.Log) (*StrategyConvexstETHUpdatedStrategist, error) {
	event := new(StrategyConvexstETHUpdatedStrategist)
	if err := _StrategyConvexstETH.contract.UnpackLog(event, "UpdatedStrategist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultAPIMetaData contains all meta data concerning the VaultAPI contract.
var VaultAPIMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"apiVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditAvailable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"debtOutstanding\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expectedReturn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_loss\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_debtPayment\",\"type\":\"uint256\"}],\"name\":\"report\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revokeStrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"}],\"name\":\"strategies\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"performanceFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"activation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debtRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rateLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastReport\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalGain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalLoss\",\"type\":\"uint256\"}],\"internalType\":\"structStrategyParams\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"25829410": "apiVersion()",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"112c1f9b": "creditAvailable()",
		"bf3759b5": "debtOutstanding()",
		"d3406abd": "expectedReturn()",
		"5aa6e675": "governance()",
		"a1d9bafc": "report(uint256,uint256,uint256)",
		"a0e4af9a": "revokeStrategy()",
		"39ebf823": "strategies(address)",
		"fc0c546a": "token()",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
		"00f714ce": "withdraw(uint256,address)",
	},
}

// VaultAPIABI is the input ABI used to generate the binding from.
// Deprecated: Use VaultAPIMetaData.ABI instead.
var VaultAPIABI = VaultAPIMetaData.ABI

// Deprecated: Use VaultAPIMetaData.Sigs instead.
// VaultAPIFuncSigs maps the 4-byte function signature to its string representation.
var VaultAPIFuncSigs = VaultAPIMetaData.Sigs

// VaultAPI is an auto generated Go binding around an Ethereum contract.
type VaultAPI struct {
	VaultAPICaller     // Read-only binding to the contract
	VaultAPITransactor // Write-only binding to the contract
	VaultAPIFilterer   // Log filterer for contract events
}

// VaultAPICaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultAPICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultAPITransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultAPITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultAPIFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultAPIFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultAPISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultAPISession struct {
	Contract     *VaultAPI         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultAPICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultAPICallerSession struct {
	Contract *VaultAPICaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// VaultAPITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultAPITransactorSession struct {
	Contract     *VaultAPITransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// VaultAPIRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultAPIRaw struct {
	Contract *VaultAPI // Generic contract binding to access the raw methods on
}

// VaultAPICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultAPICallerRaw struct {
	Contract *VaultAPICaller // Generic read-only contract binding to access the raw methods on
}

// VaultAPITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultAPITransactorRaw struct {
	Contract *VaultAPITransactor // Generic write-only contract binding to access the raw methods on
}

// NewVaultAPI creates a new instance of VaultAPI, bound to a specific deployed contract.
func NewVaultAPI(address common.Address, backend bind.ContractBackend) (*VaultAPI, error) {
	contract, err := bindVaultAPI(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VaultAPI{VaultAPICaller: VaultAPICaller{contract: contract}, VaultAPITransactor: VaultAPITransactor{contract: contract}, VaultAPIFilterer: VaultAPIFilterer{contract: contract}}, nil
}

// NewVaultAPICaller creates a new read-only instance of VaultAPI, bound to a specific deployed contract.
func NewVaultAPICaller(address common.Address, caller bind.ContractCaller) (*VaultAPICaller, error) {
	contract, err := bindVaultAPI(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultAPICaller{contract: contract}, nil
}

// NewVaultAPITransactor creates a new write-only instance of VaultAPI, bound to a specific deployed contract.
func NewVaultAPITransactor(address common.Address, transactor bind.ContractTransactor) (*VaultAPITransactor, error) {
	contract, err := bindVaultAPI(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultAPITransactor{contract: contract}, nil
}

// NewVaultAPIFilterer creates a new log filterer instance of VaultAPI, bound to a specific deployed contract.
func NewVaultAPIFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultAPIFilterer, error) {
	contract, err := bindVaultAPI(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultAPIFilterer{contract: contract}, nil
}

// bindVaultAPI binds a generic wrapper to an already deployed contract.
func bindVaultAPI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultAPIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultAPI *VaultAPIRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultAPI.Contract.VaultAPICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultAPI *VaultAPIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultAPI.Contract.VaultAPITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultAPI *VaultAPIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultAPI.Contract.VaultAPITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultAPI *VaultAPICallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultAPI.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultAPI *VaultAPITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultAPI.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultAPI *VaultAPITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultAPI.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_VaultAPI *VaultAPICaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultAPI.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_VaultAPI *VaultAPISession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _VaultAPI.Contract.Allowance(&_VaultAPI.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_VaultAPI *VaultAPICallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _VaultAPI.Contract.Allowance(&_VaultAPI.CallOpts, owner, spender)
}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() view returns(string)
func (_VaultAPI *VaultAPICaller) ApiVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VaultAPI.contract.Call(opts, &out, "apiVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() view returns(string)
func (_VaultAPI *VaultAPISession) ApiVersion() (string, error) {
	return _VaultAPI.Contract.ApiVersion(&_VaultAPI.CallOpts)
}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() view returns(string)
func (_VaultAPI *VaultAPICallerSession) ApiVersion() (string, error) {
	return _VaultAPI.Contract.ApiVersion(&_VaultAPI.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_VaultAPI *VaultAPICaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultAPI.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_VaultAPI *VaultAPISession) BalanceOf(account common.Address) (*big.Int, error) {
	return _VaultAPI.Contract.BalanceOf(&_VaultAPI.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_VaultAPI *VaultAPICallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _VaultAPI.Contract.BalanceOf(&_VaultAPI.CallOpts, account)
}

// CreditAvailable is a free data retrieval call binding the contract method 0x112c1f9b.
//
// Solidity: function creditAvailable() view returns(uint256)
func (_VaultAPI *VaultAPICaller) CreditAvailable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultAPI.contract.Call(opts, &out, "creditAvailable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreditAvailable is a free data retrieval call binding the contract method 0x112c1f9b.
//
// Solidity: function creditAvailable() view returns(uint256)
func (_VaultAPI *VaultAPISession) CreditAvailable() (*big.Int, error) {
	return _VaultAPI.Contract.CreditAvailable(&_VaultAPI.CallOpts)
}

// CreditAvailable is a free data retrieval call binding the contract method 0x112c1f9b.
//
// Solidity: function creditAvailable() view returns(uint256)
func (_VaultAPI *VaultAPICallerSession) CreditAvailable() (*big.Int, error) {
	return _VaultAPI.Contract.CreditAvailable(&_VaultAPI.CallOpts)
}

// DebtOutstanding is a free data retrieval call binding the contract method 0xbf3759b5.
//
// Solidity: function debtOutstanding() view returns(uint256)
func (_VaultAPI *VaultAPICaller) DebtOutstanding(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultAPI.contract.Call(opts, &out, "debtOutstanding")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtOutstanding is a free data retrieval call binding the contract method 0xbf3759b5.
//
// Solidity: function debtOutstanding() view returns(uint256)
func (_VaultAPI *VaultAPISession) DebtOutstanding() (*big.Int, error) {
	return _VaultAPI.Contract.DebtOutstanding(&_VaultAPI.CallOpts)
}

// DebtOutstanding is a free data retrieval call binding the contract method 0xbf3759b5.
//
// Solidity: function debtOutstanding() view returns(uint256)
func (_VaultAPI *VaultAPICallerSession) DebtOutstanding() (*big.Int, error) {
	return _VaultAPI.Contract.DebtOutstanding(&_VaultAPI.CallOpts)
}

// ExpectedReturn is a free data retrieval call binding the contract method 0xd3406abd.
//
// Solidity: function expectedReturn() view returns(uint256)
func (_VaultAPI *VaultAPICaller) ExpectedReturn(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultAPI.contract.Call(opts, &out, "expectedReturn")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpectedReturn is a free data retrieval call binding the contract method 0xd3406abd.
//
// Solidity: function expectedReturn() view returns(uint256)
func (_VaultAPI *VaultAPISession) ExpectedReturn() (*big.Int, error) {
	return _VaultAPI.Contract.ExpectedReturn(&_VaultAPI.CallOpts)
}

// ExpectedReturn is a free data retrieval call binding the contract method 0xd3406abd.
//
// Solidity: function expectedReturn() view returns(uint256)
func (_VaultAPI *VaultAPICallerSession) ExpectedReturn() (*big.Int, error) {
	return _VaultAPI.Contract.ExpectedReturn(&_VaultAPI.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_VaultAPI *VaultAPICaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultAPI.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_VaultAPI *VaultAPISession) Governance() (common.Address, error) {
	return _VaultAPI.Contract.Governance(&_VaultAPI.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_VaultAPI *VaultAPICallerSession) Governance() (common.Address, error) {
	return _VaultAPI.Contract.Governance(&_VaultAPI.CallOpts)
}

// Strategies is a free data retrieval call binding the contract method 0x39ebf823.
//
// Solidity: function strategies(address _strategy) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_VaultAPI *VaultAPICaller) Strategies(opts *bind.CallOpts, _strategy common.Address) (StrategyParams, error) {
	var out []interface{}
	err := _VaultAPI.contract.Call(opts, &out, "strategies", _strategy)

	if err != nil {
		return *new(StrategyParams), err
	}

	out0 := *abi.ConvertType(out[0], new(StrategyParams)).(*StrategyParams)

	return out0, err

}

// Strategies is a free data retrieval call binding the contract method 0x39ebf823.
//
// Solidity: function strategies(address _strategy) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_VaultAPI *VaultAPISession) Strategies(_strategy common.Address) (StrategyParams, error) {
	return _VaultAPI.Contract.Strategies(&_VaultAPI.CallOpts, _strategy)
}

// Strategies is a free data retrieval call binding the contract method 0x39ebf823.
//
// Solidity: function strategies(address _strategy) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_VaultAPI *VaultAPICallerSession) Strategies(_strategy common.Address) (StrategyParams, error) {
	return _VaultAPI.Contract.Strategies(&_VaultAPI.CallOpts, _strategy)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_VaultAPI *VaultAPICaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultAPI.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_VaultAPI *VaultAPISession) Token() (common.Address, error) {
	return _VaultAPI.Contract.Token(&_VaultAPI.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_VaultAPI *VaultAPICallerSession) Token() (common.Address, error) {
	return _VaultAPI.Contract.Token(&_VaultAPI.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VaultAPI *VaultAPICaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultAPI.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VaultAPI *VaultAPISession) TotalSupply() (*big.Int, error) {
	return _VaultAPI.Contract.TotalSupply(&_VaultAPI.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VaultAPI *VaultAPICallerSession) TotalSupply() (*big.Int, error) {
	return _VaultAPI.Contract.TotalSupply(&_VaultAPI.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_VaultAPI *VaultAPITransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultAPI.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_VaultAPI *VaultAPISession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultAPI.Contract.Approve(&_VaultAPI.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_VaultAPI *VaultAPITransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultAPI.Contract.Approve(&_VaultAPI.TransactOpts, spender, amount)
}

// Report is a paid mutator transaction binding the contract method 0xa1d9bafc.
//
// Solidity: function report(uint256 _gain, uint256 _loss, uint256 _debtPayment) returns(uint256)
func (_VaultAPI *VaultAPITransactor) Report(opts *bind.TransactOpts, _gain *big.Int, _loss *big.Int, _debtPayment *big.Int) (*types.Transaction, error) {
	return _VaultAPI.contract.Transact(opts, "report", _gain, _loss, _debtPayment)
}

// Report is a paid mutator transaction binding the contract method 0xa1d9bafc.
//
// Solidity: function report(uint256 _gain, uint256 _loss, uint256 _debtPayment) returns(uint256)
func (_VaultAPI *VaultAPISession) Report(_gain *big.Int, _loss *big.Int, _debtPayment *big.Int) (*types.Transaction, error) {
	return _VaultAPI.Contract.Report(&_VaultAPI.TransactOpts, _gain, _loss, _debtPayment)
}

// Report is a paid mutator transaction binding the contract method 0xa1d9bafc.
//
// Solidity: function report(uint256 _gain, uint256 _loss, uint256 _debtPayment) returns(uint256)
func (_VaultAPI *VaultAPITransactorSession) Report(_gain *big.Int, _loss *big.Int, _debtPayment *big.Int) (*types.Transaction, error) {
	return _VaultAPI.Contract.Report(&_VaultAPI.TransactOpts, _gain, _loss, _debtPayment)
}

// RevokeStrategy is a paid mutator transaction binding the contract method 0xa0e4af9a.
//
// Solidity: function revokeStrategy() returns()
func (_VaultAPI *VaultAPITransactor) RevokeStrategy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultAPI.contract.Transact(opts, "revokeStrategy")
}

// RevokeStrategy is a paid mutator transaction binding the contract method 0xa0e4af9a.
//
// Solidity: function revokeStrategy() returns()
func (_VaultAPI *VaultAPISession) RevokeStrategy() (*types.Transaction, error) {
	return _VaultAPI.Contract.RevokeStrategy(&_VaultAPI.TransactOpts)
}

// RevokeStrategy is a paid mutator transaction binding the contract method 0xa0e4af9a.
//
// Solidity: function revokeStrategy() returns()
func (_VaultAPI *VaultAPITransactorSession) RevokeStrategy() (*types.Transaction, error) {
	return _VaultAPI.Contract.RevokeStrategy(&_VaultAPI.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_VaultAPI *VaultAPITransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultAPI.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_VaultAPI *VaultAPISession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultAPI.Contract.Transfer(&_VaultAPI.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_VaultAPI *VaultAPITransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultAPI.Contract.Transfer(&_VaultAPI.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_VaultAPI *VaultAPITransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultAPI.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_VaultAPI *VaultAPISession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultAPI.Contract.TransferFrom(&_VaultAPI.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_VaultAPI *VaultAPITransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultAPI.Contract.TransferFrom(&_VaultAPI.TransactOpts, sender, recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 shares, address recipient) returns()
func (_VaultAPI *VaultAPITransactor) Withdraw(opts *bind.TransactOpts, shares *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _VaultAPI.contract.Transact(opts, "withdraw", shares, recipient)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 shares, address recipient) returns()
func (_VaultAPI *VaultAPISession) Withdraw(shares *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _VaultAPI.Contract.Withdraw(&_VaultAPI.TransactOpts, shares, recipient)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 shares, address recipient) returns()
func (_VaultAPI *VaultAPITransactorSession) Withdraw(shares *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _VaultAPI.Contract.Withdraw(&_VaultAPI.TransactOpts, shares, recipient)
}

// VaultAPIApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the VaultAPI contract.
type VaultAPIApprovalIterator struct {
	Event *VaultAPIApproval // Event containing the contract specifics and raw log

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
func (it *VaultAPIApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultAPIApproval)
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
		it.Event = new(VaultAPIApproval)
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
func (it *VaultAPIApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultAPIApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultAPIApproval represents a Approval event raised by the VaultAPI contract.
type VaultAPIApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_VaultAPI *VaultAPIFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*VaultAPIApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _VaultAPI.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &VaultAPIApprovalIterator{contract: _VaultAPI.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_VaultAPI *VaultAPIFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *VaultAPIApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _VaultAPI.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultAPIApproval)
				if err := _VaultAPI.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_VaultAPI *VaultAPIFilterer) ParseApproval(log types.Log) (*VaultAPIApproval, error) {
	event := new(VaultAPIApproval)
	if err := _VaultAPI.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultAPITransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the VaultAPI contract.
type VaultAPITransferIterator struct {
	Event *VaultAPITransfer // Event containing the contract specifics and raw log

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
func (it *VaultAPITransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultAPITransfer)
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
		it.Event = new(VaultAPITransfer)
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
func (it *VaultAPITransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultAPITransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultAPITransfer represents a Transfer event raised by the VaultAPI contract.
type VaultAPITransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_VaultAPI *VaultAPIFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VaultAPITransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VaultAPI.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VaultAPITransferIterator{contract: _VaultAPI.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_VaultAPI *VaultAPIFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *VaultAPITransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VaultAPI.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultAPITransfer)
				if err := _VaultAPI.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_VaultAPI *VaultAPIFilterer) ParseTransfer(log types.Log) (*VaultAPITransfer, error) {
	event := new(VaultAPITransfer)
	if err := _VaultAPI.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
