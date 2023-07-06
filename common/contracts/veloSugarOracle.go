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

// VeloSugarOracleMetaData contains all meta data concerning the VeloSugarOracle contract.
var VeloSugarOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_initcodeHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"src_len\",\"type\":\"uint8\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"connectors\",\"type\":\"address[]\"}],\"name\":\"getManyRatesWithConnectors\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"rates\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"connectors\",\"type\":\"address[]\"}],\"name\":\"getRateWithConnectors\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initcodeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VeloSugarOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use VeloSugarOracleMetaData.ABI instead.
var VeloSugarOracleABI = VeloSugarOracleMetaData.ABI

// VeloSugarOracle is an auto generated Go binding around an Ethereum contract.
type VeloSugarOracle struct {
	VeloSugarOracleCaller     // Read-only binding to the contract
	VeloSugarOracleTransactor // Write-only binding to the contract
	VeloSugarOracleFilterer   // Log filterer for contract events
}

// VeloSugarOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type VeloSugarOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VeloSugarOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VeloSugarOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VeloSugarOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VeloSugarOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VeloSugarOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VeloSugarOracleSession struct {
	Contract     *VeloSugarOracle  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VeloSugarOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VeloSugarOracleCallerSession struct {
	Contract *VeloSugarOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// VeloSugarOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VeloSugarOracleTransactorSession struct {
	Contract     *VeloSugarOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// VeloSugarOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type VeloSugarOracleRaw struct {
	Contract *VeloSugarOracle // Generic contract binding to access the raw methods on
}

// VeloSugarOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VeloSugarOracleCallerRaw struct {
	Contract *VeloSugarOracleCaller // Generic read-only contract binding to access the raw methods on
}

// VeloSugarOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VeloSugarOracleTransactorRaw struct {
	Contract *VeloSugarOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVeloSugarOracle creates a new instance of VeloSugarOracle, bound to a specific deployed contract.
func NewVeloSugarOracle(address common.Address, backend bind.ContractBackend) (*VeloSugarOracle, error) {
	contract, err := bindVeloSugarOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VeloSugarOracle{VeloSugarOracleCaller: VeloSugarOracleCaller{contract: contract}, VeloSugarOracleTransactor: VeloSugarOracleTransactor{contract: contract}, VeloSugarOracleFilterer: VeloSugarOracleFilterer{contract: contract}}, nil
}

// NewVeloSugarOracleCaller creates a new read-only instance of VeloSugarOracle, bound to a specific deployed contract.
func NewVeloSugarOracleCaller(address common.Address, caller bind.ContractCaller) (*VeloSugarOracleCaller, error) {
	contract, err := bindVeloSugarOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VeloSugarOracleCaller{contract: contract}, nil
}

// NewVeloSugarOracleTransactor creates a new write-only instance of VeloSugarOracle, bound to a specific deployed contract.
func NewVeloSugarOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*VeloSugarOracleTransactor, error) {
	contract, err := bindVeloSugarOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VeloSugarOracleTransactor{contract: contract}, nil
}

// NewVeloSugarOracleFilterer creates a new log filterer instance of VeloSugarOracle, bound to a specific deployed contract.
func NewVeloSugarOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*VeloSugarOracleFilterer, error) {
	contract, err := bindVeloSugarOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VeloSugarOracleFilterer{contract: contract}, nil
}

// bindVeloSugarOracle binds a generic wrapper to an already deployed contract.
func bindVeloSugarOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VeloSugarOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VeloSugarOracle *VeloSugarOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VeloSugarOracle.Contract.VeloSugarOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VeloSugarOracle *VeloSugarOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VeloSugarOracle.Contract.VeloSugarOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VeloSugarOracle *VeloSugarOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VeloSugarOracle.Contract.VeloSugarOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VeloSugarOracle *VeloSugarOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VeloSugarOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VeloSugarOracle *VeloSugarOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VeloSugarOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VeloSugarOracle *VeloSugarOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VeloSugarOracle.Contract.contract.Transact(opts, method, params...)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_VeloSugarOracle *VeloSugarOracleCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VeloSugarOracle.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_VeloSugarOracle *VeloSugarOracleSession) Factory() (common.Address, error) {
	return _VeloSugarOracle.Contract.Factory(&_VeloSugarOracle.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_VeloSugarOracle *VeloSugarOracleCallerSession) Factory() (common.Address, error) {
	return _VeloSugarOracle.Contract.Factory(&_VeloSugarOracle.CallOpts)
}

// GetManyRatesWithConnectors is a free data retrieval call binding the contract method 0xfe6b9b4c.
//
// Solidity: function getManyRatesWithConnectors(uint8 src_len, address[] connectors) view returns(uint256[] rates)
func (_VeloSugarOracle *VeloSugarOracleCaller) GetManyRatesWithConnectors(opts *bind.CallOpts, src_len uint8, connectors []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _VeloSugarOracle.contract.Call(opts, &out, "getManyRatesWithConnectors", src_len, connectors)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetManyRatesWithConnectors is a free data retrieval call binding the contract method 0xfe6b9b4c.
//
// Solidity: function getManyRatesWithConnectors(uint8 src_len, address[] connectors) view returns(uint256[] rates)
func (_VeloSugarOracle *VeloSugarOracleSession) GetManyRatesWithConnectors(src_len uint8, connectors []common.Address) ([]*big.Int, error) {
	return _VeloSugarOracle.Contract.GetManyRatesWithConnectors(&_VeloSugarOracle.CallOpts, src_len, connectors)
}

// GetManyRatesWithConnectors is a free data retrieval call binding the contract method 0xfe6b9b4c.
//
// Solidity: function getManyRatesWithConnectors(uint8 src_len, address[] connectors) view returns(uint256[] rates)
func (_VeloSugarOracle *VeloSugarOracleCallerSession) GetManyRatesWithConnectors(src_len uint8, connectors []common.Address) ([]*big.Int, error) {
	return _VeloSugarOracle.Contract.GetManyRatesWithConnectors(&_VeloSugarOracle.CallOpts, src_len, connectors)
}

// GetRateWithConnectors is a free data retrieval call binding the contract method 0x3e8c05bd.
//
// Solidity: function getRateWithConnectors(address[] connectors) view returns(uint256 rate)
func (_VeloSugarOracle *VeloSugarOracleCaller) GetRateWithConnectors(opts *bind.CallOpts, connectors []common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VeloSugarOracle.contract.Call(opts, &out, "getRateWithConnectors", connectors)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRateWithConnectors is a free data retrieval call binding the contract method 0x3e8c05bd.
//
// Solidity: function getRateWithConnectors(address[] connectors) view returns(uint256 rate)
func (_VeloSugarOracle *VeloSugarOracleSession) GetRateWithConnectors(connectors []common.Address) (*big.Int, error) {
	return _VeloSugarOracle.Contract.GetRateWithConnectors(&_VeloSugarOracle.CallOpts, connectors)
}

// GetRateWithConnectors is a free data retrieval call binding the contract method 0x3e8c05bd.
//
// Solidity: function getRateWithConnectors(address[] connectors) view returns(uint256 rate)
func (_VeloSugarOracle *VeloSugarOracleCallerSession) GetRateWithConnectors(connectors []common.Address) (*big.Int, error) {
	return _VeloSugarOracle.Contract.GetRateWithConnectors(&_VeloSugarOracle.CallOpts, connectors)
}

// InitcodeHash is a free data retrieval call binding the contract method 0x5a4fb9a8.
//
// Solidity: function initcodeHash() view returns(bytes32)
func (_VeloSugarOracle *VeloSugarOracleCaller) InitcodeHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VeloSugarOracle.contract.Call(opts, &out, "initcodeHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// InitcodeHash is a free data retrieval call binding the contract method 0x5a4fb9a8.
//
// Solidity: function initcodeHash() view returns(bytes32)
func (_VeloSugarOracle *VeloSugarOracleSession) InitcodeHash() ([32]byte, error) {
	return _VeloSugarOracle.Contract.InitcodeHash(&_VeloSugarOracle.CallOpts)
}

// InitcodeHash is a free data retrieval call binding the contract method 0x5a4fb9a8.
//
// Solidity: function initcodeHash() view returns(bytes32)
func (_VeloSugarOracle *VeloSugarOracleCallerSession) InitcodeHash() ([32]byte, error) {
	return _VeloSugarOracle.Contract.InitcodeHash(&_VeloSugarOracle.CallOpts)
}
