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

// FeeStruct0 is an auto generated low-level Go binding around an user-defined struct.
type FeeStruct0 struct {
	FeeBps       uint16
	FeeRecipient common.Address
}

// YRegistryV5MetaData contains all meta data concerning the YRegistryV5 contract.
var YRegistryV5MetaData = &bind.MetaData{
	ABI: "[{\"name\":\"NewVault\",\"inputs\":[{\"name\":\"vault_address\",\"type\":\"address\",\"indexed\":true},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateProtocolFeeBps\",\"inputs\":[{\"name\":\"old_fee_bps\",\"type\":\"uint16\",\"indexed\":false},{\"name\":\"new_fee_bps\",\"type\":\"uint16\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateProtocolFeeRecipient\",\"inputs\":[{\"name\":\"old_fee_recipient\",\"type\":\"address\",\"indexed\":true},{\"name\":\"new_fee_recipient\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateCustomProtocolFee\",\"inputs\":[{\"name\":\"vault\",\"type\":\"address\",\"indexed\":true},{\"name\":\"new_custom_protocol_fee\",\"type\":\"uint16\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemovedCustomProtocolFee\",\"inputs\":[{\"name\":\"vault\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"FactoryShutdown\",\"inputs\":[],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateGovernance\",\"inputs\":[{\"name\":\"governance\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewPendingGovernance\",\"inputs\":[{\"name\":\"pending_governance\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"vault_blueprint\",\"type\":\"address\"},{\"name\":\"governance\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deploy_new_vault\",\"inputs\":[{\"name\":\"asset\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"role_manager\",\"type\":\"address\"},{\"name\":\"profit_max_unlock_time\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"vault_blueprint\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"apiVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"protocol_fee_config\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"components\":[{\"name\":\"fee_bps\",\"type\":\"uint16\"},{\"name\":\"fee_recipient\",\"type\":\"address\"}]}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_protocol_fee_bps\",\"inputs\":[{\"name\":\"new_protocol_fee_bps\",\"type\":\"uint16\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_protocol_fee_recipient\",\"inputs\":[{\"name\":\"new_protocol_fee_recipient\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_custom_protocol_fee_bps\",\"inputs\":[{\"name\":\"vault\",\"type\":\"address\"},{\"name\":\"new_custom_protocol_fee\",\"type\":\"uint16\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_custom_protocol_fee\",\"inputs\":[{\"name\":\"vault\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"shutdown_factory\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_governance\",\"inputs\":[{\"name\":\"new_governance\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"accept_governance\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"shutdown\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"governance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"pending_governance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"default_protocol_fee_config\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"components\":[{\"name\":\"fee_bps\",\"type\":\"uint16\"},{\"name\":\"fee_recipient\",\"type\":\"address\"}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"custom_protocol_fee\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"use_custom_protocol_fee\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]}]",
}

// YRegistryV5ABI is the input ABI used to generate the binding from.
// Deprecated: Use YRegistryV5MetaData.ABI instead.
var YRegistryV5ABI = YRegistryV5MetaData.ABI

// YRegistryV5 is an auto generated Go binding around an Ethereum contract.
type YRegistryV5 struct {
	YRegistryV5Caller     // Read-only binding to the contract
	YRegistryV5Transactor // Write-only binding to the contract
	YRegistryV5Filterer   // Log filterer for contract events
}

// YRegistryV5Caller is an auto generated read-only Go binding around an Ethereum contract.
type YRegistryV5Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YRegistryV5Transactor is an auto generated write-only Go binding around an Ethereum contract.
type YRegistryV5Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YRegistryV5Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YRegistryV5Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YRegistryV5Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YRegistryV5Session struct {
	Contract     *YRegistryV5      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YRegistryV5CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YRegistryV5CallerSession struct {
	Contract *YRegistryV5Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// YRegistryV5TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YRegistryV5TransactorSession struct {
	Contract     *YRegistryV5Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// YRegistryV5Raw is an auto generated low-level Go binding around an Ethereum contract.
type YRegistryV5Raw struct {
	Contract *YRegistryV5 // Generic contract binding to access the raw methods on
}

// YRegistryV5CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YRegistryV5CallerRaw struct {
	Contract *YRegistryV5Caller // Generic read-only contract binding to access the raw methods on
}

// YRegistryV5TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YRegistryV5TransactorRaw struct {
	Contract *YRegistryV5Transactor // Generic write-only contract binding to access the raw methods on
}

// NewYRegistryV5 creates a new instance of YRegistryV5, bound to a specific deployed contract.
func NewYRegistryV5(address common.Address, backend bind.ContractBackend) (*YRegistryV5, error) {
	contract, err := bindYRegistryV5(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YRegistryV5{YRegistryV5Caller: YRegistryV5Caller{contract: contract}, YRegistryV5Transactor: YRegistryV5Transactor{contract: contract}, YRegistryV5Filterer: YRegistryV5Filterer{contract: contract}}, nil
}

// NewYRegistryV5Caller creates a new read-only instance of YRegistryV5, bound to a specific deployed contract.
func NewYRegistryV5Caller(address common.Address, caller bind.ContractCaller) (*YRegistryV5Caller, error) {
	contract, err := bindYRegistryV5(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YRegistryV5Caller{contract: contract}, nil
}

// NewYRegistryV5Transactor creates a new write-only instance of YRegistryV5, bound to a specific deployed contract.
func NewYRegistryV5Transactor(address common.Address, transactor bind.ContractTransactor) (*YRegistryV5Transactor, error) {
	contract, err := bindYRegistryV5(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YRegistryV5Transactor{contract: contract}, nil
}

// NewYRegistryV5Filterer creates a new log filterer instance of YRegistryV5, bound to a specific deployed contract.
func NewYRegistryV5Filterer(address common.Address, filterer bind.ContractFilterer) (*YRegistryV5Filterer, error) {
	contract, err := bindYRegistryV5(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YRegistryV5Filterer{contract: contract}, nil
}

// bindYRegistryV5 binds a generic wrapper to an already deployed contract.
func bindYRegistryV5(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(YRegistryV5ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YRegistryV5 *YRegistryV5Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YRegistryV5.Contract.YRegistryV5Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YRegistryV5 *YRegistryV5Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YRegistryV5.Contract.YRegistryV5Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YRegistryV5 *YRegistryV5Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YRegistryV5.Contract.YRegistryV5Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YRegistryV5 *YRegistryV5CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YRegistryV5.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YRegistryV5 *YRegistryV5TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YRegistryV5.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YRegistryV5 *YRegistryV5TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YRegistryV5.Contract.contract.Transact(opts, method, params...)
}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() view returns(string)
func (_YRegistryV5 *YRegistryV5Caller) ApiVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _YRegistryV5.contract.Call(opts, &out, "apiVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() view returns(string)
func (_YRegistryV5 *YRegistryV5Session) ApiVersion() (string, error) {
	return _YRegistryV5.Contract.ApiVersion(&_YRegistryV5.CallOpts)
}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() view returns(string)
func (_YRegistryV5 *YRegistryV5CallerSession) ApiVersion() (string, error) {
	return _YRegistryV5.Contract.ApiVersion(&_YRegistryV5.CallOpts)
}

// CustomProtocolFee is a free data retrieval call binding the contract method 0xcbe28663.
//
// Solidity: function custom_protocol_fee(address arg0) view returns(uint16)
func (_YRegistryV5 *YRegistryV5Caller) CustomProtocolFee(opts *bind.CallOpts, arg0 common.Address) (uint16, error) {
	var out []interface{}
	err := _YRegistryV5.contract.Call(opts, &out, "custom_protocol_fee", arg0)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// CustomProtocolFee is a free data retrieval call binding the contract method 0xcbe28663.
//
// Solidity: function custom_protocol_fee(address arg0) view returns(uint16)
func (_YRegistryV5 *YRegistryV5Session) CustomProtocolFee(arg0 common.Address) (uint16, error) {
	return _YRegistryV5.Contract.CustomProtocolFee(&_YRegistryV5.CallOpts, arg0)
}

// CustomProtocolFee is a free data retrieval call binding the contract method 0xcbe28663.
//
// Solidity: function custom_protocol_fee(address arg0) view returns(uint16)
func (_YRegistryV5 *YRegistryV5CallerSession) CustomProtocolFee(arg0 common.Address) (uint16, error) {
	return _YRegistryV5.Contract.CustomProtocolFee(&_YRegistryV5.CallOpts, arg0)
}

// DefaultProtocolFeeConfig is a free data retrieval call binding the contract method 0x97ad2ecc.
//
// Solidity: function default_protocol_fee_config() view returns((uint16,address))
func (_YRegistryV5 *YRegistryV5Caller) DefaultProtocolFeeConfig(opts *bind.CallOpts) (FeeStruct0, error) {
	var out []interface{}
	err := _YRegistryV5.contract.Call(opts, &out, "default_protocol_fee_config")

	if err != nil {
		return *new(FeeStruct0), err
	}

	out0 := *abi.ConvertType(out[0], new(FeeStruct0)).(*FeeStruct0)

	return out0, err

}

// DefaultProtocolFeeConfig is a free data retrieval call binding the contract method 0x97ad2ecc.
//
// Solidity: function default_protocol_fee_config() view returns((uint16,address))
func (_YRegistryV5 *YRegistryV5Session) DefaultProtocolFeeConfig() (FeeStruct0, error) {
	return _YRegistryV5.Contract.DefaultProtocolFeeConfig(&_YRegistryV5.CallOpts)
}

// DefaultProtocolFeeConfig is a free data retrieval call binding the contract method 0x97ad2ecc.
//
// Solidity: function default_protocol_fee_config() view returns((uint16,address))
func (_YRegistryV5 *YRegistryV5CallerSession) DefaultProtocolFeeConfig() (FeeStruct0, error) {
	return _YRegistryV5.Contract.DefaultProtocolFeeConfig(&_YRegistryV5.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_YRegistryV5 *YRegistryV5Caller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YRegistryV5.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_YRegistryV5 *YRegistryV5Session) Governance() (common.Address, error) {
	return _YRegistryV5.Contract.Governance(&_YRegistryV5.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_YRegistryV5 *YRegistryV5CallerSession) Governance() (common.Address, error) {
	return _YRegistryV5.Contract.Governance(&_YRegistryV5.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YRegistryV5 *YRegistryV5Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _YRegistryV5.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YRegistryV5 *YRegistryV5Session) Name() (string, error) {
	return _YRegistryV5.Contract.Name(&_YRegistryV5.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YRegistryV5 *YRegistryV5CallerSession) Name() (string, error) {
	return _YRegistryV5.Contract.Name(&_YRegistryV5.CallOpts)
}

// PendingGovernance is a free data retrieval call binding the contract method 0xc66eb0a2.
//
// Solidity: function pending_governance() view returns(address)
func (_YRegistryV5 *YRegistryV5Caller) PendingGovernance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YRegistryV5.contract.Call(opts, &out, "pending_governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingGovernance is a free data retrieval call binding the contract method 0xc66eb0a2.
//
// Solidity: function pending_governance() view returns(address)
func (_YRegistryV5 *YRegistryV5Session) PendingGovernance() (common.Address, error) {
	return _YRegistryV5.Contract.PendingGovernance(&_YRegistryV5.CallOpts)
}

// PendingGovernance is a free data retrieval call binding the contract method 0xc66eb0a2.
//
// Solidity: function pending_governance() view returns(address)
func (_YRegistryV5 *YRegistryV5CallerSession) PendingGovernance() (common.Address, error) {
	return _YRegistryV5.Contract.PendingGovernance(&_YRegistryV5.CallOpts)
}

// ProtocolFeeConfig is a free data retrieval call binding the contract method 0x5153b199.
//
// Solidity: function protocol_fee_config() view returns((uint16,address))
func (_YRegistryV5 *YRegistryV5Caller) ProtocolFeeConfig(opts *bind.CallOpts) (FeeStruct0, error) {
	var out []interface{}
	err := _YRegistryV5.contract.Call(opts, &out, "protocol_fee_config")

	if err != nil {
		return *new(FeeStruct0), err
	}

	out0 := *abi.ConvertType(out[0], new(FeeStruct0)).(*FeeStruct0)

	return out0, err

}

// ProtocolFeeConfig is a free data retrieval call binding the contract method 0x5153b199.
//
// Solidity: function protocol_fee_config() view returns((uint16,address))
func (_YRegistryV5 *YRegistryV5Session) ProtocolFeeConfig() (FeeStruct0, error) {
	return _YRegistryV5.Contract.ProtocolFeeConfig(&_YRegistryV5.CallOpts)
}

// ProtocolFeeConfig is a free data retrieval call binding the contract method 0x5153b199.
//
// Solidity: function protocol_fee_config() view returns((uint16,address))
func (_YRegistryV5 *YRegistryV5CallerSession) ProtocolFeeConfig() (FeeStruct0, error) {
	return _YRegistryV5.Contract.ProtocolFeeConfig(&_YRegistryV5.CallOpts)
}

// Shutdown is a free data retrieval call binding the contract method 0xfc0e74d1.
//
// Solidity: function shutdown() view returns(bool)
func (_YRegistryV5 *YRegistryV5Caller) Shutdown(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _YRegistryV5.contract.Call(opts, &out, "shutdown")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Shutdown is a free data retrieval call binding the contract method 0xfc0e74d1.
//
// Solidity: function shutdown() view returns(bool)
func (_YRegistryV5 *YRegistryV5Session) Shutdown() (bool, error) {
	return _YRegistryV5.Contract.Shutdown(&_YRegistryV5.CallOpts)
}

// Shutdown is a free data retrieval call binding the contract method 0xfc0e74d1.
//
// Solidity: function shutdown() view returns(bool)
func (_YRegistryV5 *YRegistryV5CallerSession) Shutdown() (bool, error) {
	return _YRegistryV5.Contract.Shutdown(&_YRegistryV5.CallOpts)
}

// UseCustomProtocolFee is a free data retrieval call binding the contract method 0xe94860d8.
//
// Solidity: function use_custom_protocol_fee(address arg0) view returns(bool)
func (_YRegistryV5 *YRegistryV5Caller) UseCustomProtocolFee(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _YRegistryV5.contract.Call(opts, &out, "use_custom_protocol_fee", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UseCustomProtocolFee is a free data retrieval call binding the contract method 0xe94860d8.
//
// Solidity: function use_custom_protocol_fee(address arg0) view returns(bool)
func (_YRegistryV5 *YRegistryV5Session) UseCustomProtocolFee(arg0 common.Address) (bool, error) {
	return _YRegistryV5.Contract.UseCustomProtocolFee(&_YRegistryV5.CallOpts, arg0)
}

// UseCustomProtocolFee is a free data retrieval call binding the contract method 0xe94860d8.
//
// Solidity: function use_custom_protocol_fee(address arg0) view returns(bool)
func (_YRegistryV5 *YRegistryV5CallerSession) UseCustomProtocolFee(arg0 common.Address) (bool, error) {
	return _YRegistryV5.Contract.UseCustomProtocolFee(&_YRegistryV5.CallOpts, arg0)
}

// VaultBlueprint is a free data retrieval call binding the contract method 0x05f16b0f.
//
// Solidity: function vault_blueprint() view returns(address)
func (_YRegistryV5 *YRegistryV5Caller) VaultBlueprint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YRegistryV5.contract.Call(opts, &out, "vault_blueprint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultBlueprint is a free data retrieval call binding the contract method 0x05f16b0f.
//
// Solidity: function vault_blueprint() view returns(address)
func (_YRegistryV5 *YRegistryV5Session) VaultBlueprint() (common.Address, error) {
	return _YRegistryV5.Contract.VaultBlueprint(&_YRegistryV5.CallOpts)
}

// VaultBlueprint is a free data retrieval call binding the contract method 0x05f16b0f.
//
// Solidity: function vault_blueprint() view returns(address)
func (_YRegistryV5 *YRegistryV5CallerSession) VaultBlueprint() (common.Address, error) {
	return _YRegistryV5.Contract.VaultBlueprint(&_YRegistryV5.CallOpts)
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0xa7dbff3e.
//
// Solidity: function accept_governance() returns()
func (_YRegistryV5 *YRegistryV5Transactor) AcceptGovernance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YRegistryV5.contract.Transact(opts, "accept_governance")
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0xa7dbff3e.
//
// Solidity: function accept_governance() returns()
func (_YRegistryV5 *YRegistryV5Session) AcceptGovernance() (*types.Transaction, error) {
	return _YRegistryV5.Contract.AcceptGovernance(&_YRegistryV5.TransactOpts)
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0xa7dbff3e.
//
// Solidity: function accept_governance() returns()
func (_YRegistryV5 *YRegistryV5TransactorSession) AcceptGovernance() (*types.Transaction, error) {
	return _YRegistryV5.Contract.AcceptGovernance(&_YRegistryV5.TransactOpts)
}

// DeployNewVault is a paid mutator transaction binding the contract method 0xb4aeee77.
//
// Solidity: function deploy_new_vault(address asset, string name, string symbol, address role_manager, uint256 profit_max_unlock_time) returns(address)
func (_YRegistryV5 *YRegistryV5Transactor) DeployNewVault(opts *bind.TransactOpts, asset common.Address, name string, symbol string, role_manager common.Address, profit_max_unlock_time *big.Int) (*types.Transaction, error) {
	return _YRegistryV5.contract.Transact(opts, "deploy_new_vault", asset, name, symbol, role_manager, profit_max_unlock_time)
}

// DeployNewVault is a paid mutator transaction binding the contract method 0xb4aeee77.
//
// Solidity: function deploy_new_vault(address asset, string name, string symbol, address role_manager, uint256 profit_max_unlock_time) returns(address)
func (_YRegistryV5 *YRegistryV5Session) DeployNewVault(asset common.Address, name string, symbol string, role_manager common.Address, profit_max_unlock_time *big.Int) (*types.Transaction, error) {
	return _YRegistryV5.Contract.DeployNewVault(&_YRegistryV5.TransactOpts, asset, name, symbol, role_manager, profit_max_unlock_time)
}

// DeployNewVault is a paid mutator transaction binding the contract method 0xb4aeee77.
//
// Solidity: function deploy_new_vault(address asset, string name, string symbol, address role_manager, uint256 profit_max_unlock_time) returns(address)
func (_YRegistryV5 *YRegistryV5TransactorSession) DeployNewVault(asset common.Address, name string, symbol string, role_manager common.Address, profit_max_unlock_time *big.Int) (*types.Transaction, error) {
	return _YRegistryV5.Contract.DeployNewVault(&_YRegistryV5.TransactOpts, asset, name, symbol, role_manager, profit_max_unlock_time)
}

// RemoveCustomProtocolFee is a paid mutator transaction binding the contract method 0x11a3a434.
//
// Solidity: function remove_custom_protocol_fee(address vault) returns()
func (_YRegistryV5 *YRegistryV5Transactor) RemoveCustomProtocolFee(opts *bind.TransactOpts, vault common.Address) (*types.Transaction, error) {
	return _YRegistryV5.contract.Transact(opts, "remove_custom_protocol_fee", vault)
}

// RemoveCustomProtocolFee is a paid mutator transaction binding the contract method 0x11a3a434.
//
// Solidity: function remove_custom_protocol_fee(address vault) returns()
func (_YRegistryV5 *YRegistryV5Session) RemoveCustomProtocolFee(vault common.Address) (*types.Transaction, error) {
	return _YRegistryV5.Contract.RemoveCustomProtocolFee(&_YRegistryV5.TransactOpts, vault)
}

// RemoveCustomProtocolFee is a paid mutator transaction binding the contract method 0x11a3a434.
//
// Solidity: function remove_custom_protocol_fee(address vault) returns()
func (_YRegistryV5 *YRegistryV5TransactorSession) RemoveCustomProtocolFee(vault common.Address) (*types.Transaction, error) {
	return _YRegistryV5.Contract.RemoveCustomProtocolFee(&_YRegistryV5.TransactOpts, vault)
}

// SetCustomProtocolFeeBps is a paid mutator transaction binding the contract method 0xb5a71e07.
//
// Solidity: function set_custom_protocol_fee_bps(address vault, uint16 new_custom_protocol_fee) returns()
func (_YRegistryV5 *YRegistryV5Transactor) SetCustomProtocolFeeBps(opts *bind.TransactOpts, vault common.Address, new_custom_protocol_fee uint16) (*types.Transaction, error) {
	return _YRegistryV5.contract.Transact(opts, "set_custom_protocol_fee_bps", vault, new_custom_protocol_fee)
}

// SetCustomProtocolFeeBps is a paid mutator transaction binding the contract method 0xb5a71e07.
//
// Solidity: function set_custom_protocol_fee_bps(address vault, uint16 new_custom_protocol_fee) returns()
func (_YRegistryV5 *YRegistryV5Session) SetCustomProtocolFeeBps(vault common.Address, new_custom_protocol_fee uint16) (*types.Transaction, error) {
	return _YRegistryV5.Contract.SetCustomProtocolFeeBps(&_YRegistryV5.TransactOpts, vault, new_custom_protocol_fee)
}

// SetCustomProtocolFeeBps is a paid mutator transaction binding the contract method 0xb5a71e07.
//
// Solidity: function set_custom_protocol_fee_bps(address vault, uint16 new_custom_protocol_fee) returns()
func (_YRegistryV5 *YRegistryV5TransactorSession) SetCustomProtocolFeeBps(vault common.Address, new_custom_protocol_fee uint16) (*types.Transaction, error) {
	return _YRegistryV5.Contract.SetCustomProtocolFeeBps(&_YRegistryV5.TransactOpts, vault, new_custom_protocol_fee)
}

// SetGovernance is a paid mutator transaction binding the contract method 0x070313fa.
//
// Solidity: function set_governance(address new_governance) returns()
func (_YRegistryV5 *YRegistryV5Transactor) SetGovernance(opts *bind.TransactOpts, new_governance common.Address) (*types.Transaction, error) {
	return _YRegistryV5.contract.Transact(opts, "set_governance", new_governance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0x070313fa.
//
// Solidity: function set_governance(address new_governance) returns()
func (_YRegistryV5 *YRegistryV5Session) SetGovernance(new_governance common.Address) (*types.Transaction, error) {
	return _YRegistryV5.Contract.SetGovernance(&_YRegistryV5.TransactOpts, new_governance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0x070313fa.
//
// Solidity: function set_governance(address new_governance) returns()
func (_YRegistryV5 *YRegistryV5TransactorSession) SetGovernance(new_governance common.Address) (*types.Transaction, error) {
	return _YRegistryV5.Contract.SetGovernance(&_YRegistryV5.TransactOpts, new_governance)
}

// SetProtocolFeeBps is a paid mutator transaction binding the contract method 0x62fbf603.
//
// Solidity: function set_protocol_fee_bps(uint16 new_protocol_fee_bps) returns()
func (_YRegistryV5 *YRegistryV5Transactor) SetProtocolFeeBps(opts *bind.TransactOpts, new_protocol_fee_bps uint16) (*types.Transaction, error) {
	return _YRegistryV5.contract.Transact(opts, "set_protocol_fee_bps", new_protocol_fee_bps)
}

// SetProtocolFeeBps is a paid mutator transaction binding the contract method 0x62fbf603.
//
// Solidity: function set_protocol_fee_bps(uint16 new_protocol_fee_bps) returns()
func (_YRegistryV5 *YRegistryV5Session) SetProtocolFeeBps(new_protocol_fee_bps uint16) (*types.Transaction, error) {
	return _YRegistryV5.Contract.SetProtocolFeeBps(&_YRegistryV5.TransactOpts, new_protocol_fee_bps)
}

// SetProtocolFeeBps is a paid mutator transaction binding the contract method 0x62fbf603.
//
// Solidity: function set_protocol_fee_bps(uint16 new_protocol_fee_bps) returns()
func (_YRegistryV5 *YRegistryV5TransactorSession) SetProtocolFeeBps(new_protocol_fee_bps uint16) (*types.Transaction, error) {
	return _YRegistryV5.Contract.SetProtocolFeeBps(&_YRegistryV5.TransactOpts, new_protocol_fee_bps)
}

// SetProtocolFeeRecipient is a paid mutator transaction binding the contract method 0xf8ebccea.
//
// Solidity: function set_protocol_fee_recipient(address new_protocol_fee_recipient) returns()
func (_YRegistryV5 *YRegistryV5Transactor) SetProtocolFeeRecipient(opts *bind.TransactOpts, new_protocol_fee_recipient common.Address) (*types.Transaction, error) {
	return _YRegistryV5.contract.Transact(opts, "set_protocol_fee_recipient", new_protocol_fee_recipient)
}

// SetProtocolFeeRecipient is a paid mutator transaction binding the contract method 0xf8ebccea.
//
// Solidity: function set_protocol_fee_recipient(address new_protocol_fee_recipient) returns()
func (_YRegistryV5 *YRegistryV5Session) SetProtocolFeeRecipient(new_protocol_fee_recipient common.Address) (*types.Transaction, error) {
	return _YRegistryV5.Contract.SetProtocolFeeRecipient(&_YRegistryV5.TransactOpts, new_protocol_fee_recipient)
}

// SetProtocolFeeRecipient is a paid mutator transaction binding the contract method 0xf8ebccea.
//
// Solidity: function set_protocol_fee_recipient(address new_protocol_fee_recipient) returns()
func (_YRegistryV5 *YRegistryV5TransactorSession) SetProtocolFeeRecipient(new_protocol_fee_recipient common.Address) (*types.Transaction, error) {
	return _YRegistryV5.Contract.SetProtocolFeeRecipient(&_YRegistryV5.TransactOpts, new_protocol_fee_recipient)
}

// ShutdownFactory is a paid mutator transaction binding the contract method 0x365adba6.
//
// Solidity: function shutdown_factory() returns()
func (_YRegistryV5 *YRegistryV5Transactor) ShutdownFactory(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YRegistryV5.contract.Transact(opts, "shutdown_factory")
}

// ShutdownFactory is a paid mutator transaction binding the contract method 0x365adba6.
//
// Solidity: function shutdown_factory() returns()
func (_YRegistryV5 *YRegistryV5Session) ShutdownFactory() (*types.Transaction, error) {
	return _YRegistryV5.Contract.ShutdownFactory(&_YRegistryV5.TransactOpts)
}

// ShutdownFactory is a paid mutator transaction binding the contract method 0x365adba6.
//
// Solidity: function shutdown_factory() returns()
func (_YRegistryV5 *YRegistryV5TransactorSession) ShutdownFactory() (*types.Transaction, error) {
	return _YRegistryV5.Contract.ShutdownFactory(&_YRegistryV5.TransactOpts)
}

// YRegistryV5FactoryShutdownIterator is returned from FilterFactoryShutdown and is used to iterate over the raw logs and unpacked data for FactoryShutdown events raised by the YRegistryV5 contract.
type YRegistryV5FactoryShutdownIterator struct {
	Event *YRegistryV5FactoryShutdown // Event containing the contract specifics and raw log

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
func (it *YRegistryV5FactoryShutdownIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YRegistryV5FactoryShutdown)
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
		it.Event = new(YRegistryV5FactoryShutdown)
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
func (it *YRegistryV5FactoryShutdownIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YRegistryV5FactoryShutdownIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YRegistryV5FactoryShutdown represents a FactoryShutdown event raised by the YRegistryV5 contract.
type YRegistryV5FactoryShutdown struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFactoryShutdown is a free log retrieval operation binding the contract event 0xc643193a97fc0e18d69c95e1c034b91f51fa164ba8ea68dfb6dd98568b9bc96b.
//
// Solidity: event FactoryShutdown()
func (_YRegistryV5 *YRegistryV5Filterer) FilterFactoryShutdown(opts *bind.FilterOpts) (*YRegistryV5FactoryShutdownIterator, error) {

	logs, sub, err := _YRegistryV5.contract.FilterLogs(opts, "FactoryShutdown")
	if err != nil {
		return nil, err
	}
	return &YRegistryV5FactoryShutdownIterator{contract: _YRegistryV5.contract, event: "FactoryShutdown", logs: logs, sub: sub}, nil
}

// WatchFactoryShutdown is a free log subscription operation binding the contract event 0xc643193a97fc0e18d69c95e1c034b91f51fa164ba8ea68dfb6dd98568b9bc96b.
//
// Solidity: event FactoryShutdown()
func (_YRegistryV5 *YRegistryV5Filterer) WatchFactoryShutdown(opts *bind.WatchOpts, sink chan<- *YRegistryV5FactoryShutdown) (event.Subscription, error) {

	logs, sub, err := _YRegistryV5.contract.WatchLogs(opts, "FactoryShutdown")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YRegistryV5FactoryShutdown)
				if err := _YRegistryV5.contract.UnpackLog(event, "FactoryShutdown", log); err != nil {
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

// ParseFactoryShutdown is a log parse operation binding the contract event 0xc643193a97fc0e18d69c95e1c034b91f51fa164ba8ea68dfb6dd98568b9bc96b.
//
// Solidity: event FactoryShutdown()
func (_YRegistryV5 *YRegistryV5Filterer) ParseFactoryShutdown(log types.Log) (*YRegistryV5FactoryShutdown, error) {
	event := new(YRegistryV5FactoryShutdown)
	if err := _YRegistryV5.contract.UnpackLog(event, "FactoryShutdown", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YRegistryV5NewPendingGovernanceIterator is returned from FilterNewPendingGovernance and is used to iterate over the raw logs and unpacked data for NewPendingGovernance events raised by the YRegistryV5 contract.
type YRegistryV5NewPendingGovernanceIterator struct {
	Event *YRegistryV5NewPendingGovernance // Event containing the contract specifics and raw log

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
func (it *YRegistryV5NewPendingGovernanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YRegistryV5NewPendingGovernance)
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
		it.Event = new(YRegistryV5NewPendingGovernance)
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
func (it *YRegistryV5NewPendingGovernanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YRegistryV5NewPendingGovernanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YRegistryV5NewPendingGovernance represents a NewPendingGovernance event raised by the YRegistryV5 contract.
type YRegistryV5NewPendingGovernance struct {
	PendingGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewPendingGovernance is a free log retrieval operation binding the contract event 0x90ad4c550d25bd23af61db38d1ff8671b89edaaa0bca0fc36bac5084ecc120bd.
//
// Solidity: event NewPendingGovernance(address indexed pending_governance)
func (_YRegistryV5 *YRegistryV5Filterer) FilterNewPendingGovernance(opts *bind.FilterOpts, pending_governance []common.Address) (*YRegistryV5NewPendingGovernanceIterator, error) {

	var pending_governanceRule []interface{}
	for _, pending_governanceItem := range pending_governance {
		pending_governanceRule = append(pending_governanceRule, pending_governanceItem)
	}

	logs, sub, err := _YRegistryV5.contract.FilterLogs(opts, "NewPendingGovernance", pending_governanceRule)
	if err != nil {
		return nil, err
	}
	return &YRegistryV5NewPendingGovernanceIterator{contract: _YRegistryV5.contract, event: "NewPendingGovernance", logs: logs, sub: sub}, nil
}

// WatchNewPendingGovernance is a free log subscription operation binding the contract event 0x90ad4c550d25bd23af61db38d1ff8671b89edaaa0bca0fc36bac5084ecc120bd.
//
// Solidity: event NewPendingGovernance(address indexed pending_governance)
func (_YRegistryV5 *YRegistryV5Filterer) WatchNewPendingGovernance(opts *bind.WatchOpts, sink chan<- *YRegistryV5NewPendingGovernance, pending_governance []common.Address) (event.Subscription, error) {

	var pending_governanceRule []interface{}
	for _, pending_governanceItem := range pending_governance {
		pending_governanceRule = append(pending_governanceRule, pending_governanceItem)
	}

	logs, sub, err := _YRegistryV5.contract.WatchLogs(opts, "NewPendingGovernance", pending_governanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YRegistryV5NewPendingGovernance)
				if err := _YRegistryV5.contract.UnpackLog(event, "NewPendingGovernance", log); err != nil {
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

// ParseNewPendingGovernance is a log parse operation binding the contract event 0x90ad4c550d25bd23af61db38d1ff8671b89edaaa0bca0fc36bac5084ecc120bd.
//
// Solidity: event NewPendingGovernance(address indexed pending_governance)
func (_YRegistryV5 *YRegistryV5Filterer) ParseNewPendingGovernance(log types.Log) (*YRegistryV5NewPendingGovernance, error) {
	event := new(YRegistryV5NewPendingGovernance)
	if err := _YRegistryV5.contract.UnpackLog(event, "NewPendingGovernance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YRegistryV5NewVaultIterator is returned from FilterNewVault and is used to iterate over the raw logs and unpacked data for NewVault events raised by the YRegistryV5 contract.
type YRegistryV5NewVaultIterator struct {
	Event *YRegistryV5NewVault // Event containing the contract specifics and raw log

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
func (it *YRegistryV5NewVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YRegistryV5NewVault)
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
		it.Event = new(YRegistryV5NewVault)
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
func (it *YRegistryV5NewVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YRegistryV5NewVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YRegistryV5NewVault represents a NewVault event raised by the YRegistryV5 contract.
type YRegistryV5NewVault struct {
	VaultAddress common.Address
	Asset        common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewVault is a free log retrieval operation binding the contract event 0x4241302c393c713e690702c4a45a57e93cef59aa8c6e2358495853b3420551d8.
//
// Solidity: event NewVault(address indexed vault_address, address indexed asset)
func (_YRegistryV5 *YRegistryV5Filterer) FilterNewVault(opts *bind.FilterOpts, vault_address []common.Address, asset []common.Address) (*YRegistryV5NewVaultIterator, error) {

	var vault_addressRule []interface{}
	for _, vault_addressItem := range vault_address {
		vault_addressRule = append(vault_addressRule, vault_addressItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _YRegistryV5.contract.FilterLogs(opts, "NewVault", vault_addressRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &YRegistryV5NewVaultIterator{contract: _YRegistryV5.contract, event: "NewVault", logs: logs, sub: sub}, nil
}

// WatchNewVault is a free log subscription operation binding the contract event 0x4241302c393c713e690702c4a45a57e93cef59aa8c6e2358495853b3420551d8.
//
// Solidity: event NewVault(address indexed vault_address, address indexed asset)
func (_YRegistryV5 *YRegistryV5Filterer) WatchNewVault(opts *bind.WatchOpts, sink chan<- *YRegistryV5NewVault, vault_address []common.Address, asset []common.Address) (event.Subscription, error) {

	var vault_addressRule []interface{}
	for _, vault_addressItem := range vault_address {
		vault_addressRule = append(vault_addressRule, vault_addressItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _YRegistryV5.contract.WatchLogs(opts, "NewVault", vault_addressRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YRegistryV5NewVault)
				if err := _YRegistryV5.contract.UnpackLog(event, "NewVault", log); err != nil {
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

// ParseNewVault is a log parse operation binding the contract event 0x4241302c393c713e690702c4a45a57e93cef59aa8c6e2358495853b3420551d8.
//
// Solidity: event NewVault(address indexed vault_address, address indexed asset)
func (_YRegistryV5 *YRegistryV5Filterer) ParseNewVault(log types.Log) (*YRegistryV5NewVault, error) {
	event := new(YRegistryV5NewVault)
	if err := _YRegistryV5.contract.UnpackLog(event, "NewVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YRegistryV5RemovedCustomProtocolFeeIterator is returned from FilterRemovedCustomProtocolFee and is used to iterate over the raw logs and unpacked data for RemovedCustomProtocolFee events raised by the YRegistryV5 contract.
type YRegistryV5RemovedCustomProtocolFeeIterator struct {
	Event *YRegistryV5RemovedCustomProtocolFee // Event containing the contract specifics and raw log

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
func (it *YRegistryV5RemovedCustomProtocolFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YRegistryV5RemovedCustomProtocolFee)
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
		it.Event = new(YRegistryV5RemovedCustomProtocolFee)
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
func (it *YRegistryV5RemovedCustomProtocolFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YRegistryV5RemovedCustomProtocolFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YRegistryV5RemovedCustomProtocolFee represents a RemovedCustomProtocolFee event raised by the YRegistryV5 contract.
type YRegistryV5RemovedCustomProtocolFee struct {
	Vault common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRemovedCustomProtocolFee is a free log retrieval operation binding the contract event 0x39612c4f13d7a058dece05cf6730e3322fd9a11d6f055a5eacdde49191f45f1f.
//
// Solidity: event RemovedCustomProtocolFee(address indexed vault)
func (_YRegistryV5 *YRegistryV5Filterer) FilterRemovedCustomProtocolFee(opts *bind.FilterOpts, vault []common.Address) (*YRegistryV5RemovedCustomProtocolFeeIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _YRegistryV5.contract.FilterLogs(opts, "RemovedCustomProtocolFee", vaultRule)
	if err != nil {
		return nil, err
	}
	return &YRegistryV5RemovedCustomProtocolFeeIterator{contract: _YRegistryV5.contract, event: "RemovedCustomProtocolFee", logs: logs, sub: sub}, nil
}

// WatchRemovedCustomProtocolFee is a free log subscription operation binding the contract event 0x39612c4f13d7a058dece05cf6730e3322fd9a11d6f055a5eacdde49191f45f1f.
//
// Solidity: event RemovedCustomProtocolFee(address indexed vault)
func (_YRegistryV5 *YRegistryV5Filterer) WatchRemovedCustomProtocolFee(opts *bind.WatchOpts, sink chan<- *YRegistryV5RemovedCustomProtocolFee, vault []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _YRegistryV5.contract.WatchLogs(opts, "RemovedCustomProtocolFee", vaultRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YRegistryV5RemovedCustomProtocolFee)
				if err := _YRegistryV5.contract.UnpackLog(event, "RemovedCustomProtocolFee", log); err != nil {
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

// ParseRemovedCustomProtocolFee is a log parse operation binding the contract event 0x39612c4f13d7a058dece05cf6730e3322fd9a11d6f055a5eacdde49191f45f1f.
//
// Solidity: event RemovedCustomProtocolFee(address indexed vault)
func (_YRegistryV5 *YRegistryV5Filterer) ParseRemovedCustomProtocolFee(log types.Log) (*YRegistryV5RemovedCustomProtocolFee, error) {
	event := new(YRegistryV5RemovedCustomProtocolFee)
	if err := _YRegistryV5.contract.UnpackLog(event, "RemovedCustomProtocolFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YRegistryV5UpdateCustomProtocolFeeIterator is returned from FilterUpdateCustomProtocolFee and is used to iterate over the raw logs and unpacked data for UpdateCustomProtocolFee events raised by the YRegistryV5 contract.
type YRegistryV5UpdateCustomProtocolFeeIterator struct {
	Event *YRegistryV5UpdateCustomProtocolFee // Event containing the contract specifics and raw log

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
func (it *YRegistryV5UpdateCustomProtocolFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YRegistryV5UpdateCustomProtocolFee)
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
		it.Event = new(YRegistryV5UpdateCustomProtocolFee)
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
func (it *YRegistryV5UpdateCustomProtocolFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YRegistryV5UpdateCustomProtocolFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YRegistryV5UpdateCustomProtocolFee represents a UpdateCustomProtocolFee event raised by the YRegistryV5 contract.
type YRegistryV5UpdateCustomProtocolFee struct {
	Vault                common.Address
	NewCustomProtocolFee uint16
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdateCustomProtocolFee is a free log retrieval operation binding the contract event 0x96d6cc624354ffe5a7207dc2dcc152e58e23ac8df9c96943f3cfb10ea4c140ac.
//
// Solidity: event UpdateCustomProtocolFee(address indexed vault, uint16 new_custom_protocol_fee)
func (_YRegistryV5 *YRegistryV5Filterer) FilterUpdateCustomProtocolFee(opts *bind.FilterOpts, vault []common.Address) (*YRegistryV5UpdateCustomProtocolFeeIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _YRegistryV5.contract.FilterLogs(opts, "UpdateCustomProtocolFee", vaultRule)
	if err != nil {
		return nil, err
	}
	return &YRegistryV5UpdateCustomProtocolFeeIterator{contract: _YRegistryV5.contract, event: "UpdateCustomProtocolFee", logs: logs, sub: sub}, nil
}

// WatchUpdateCustomProtocolFee is a free log subscription operation binding the contract event 0x96d6cc624354ffe5a7207dc2dcc152e58e23ac8df9c96943f3cfb10ea4c140ac.
//
// Solidity: event UpdateCustomProtocolFee(address indexed vault, uint16 new_custom_protocol_fee)
func (_YRegistryV5 *YRegistryV5Filterer) WatchUpdateCustomProtocolFee(opts *bind.WatchOpts, sink chan<- *YRegistryV5UpdateCustomProtocolFee, vault []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _YRegistryV5.contract.WatchLogs(opts, "UpdateCustomProtocolFee", vaultRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YRegistryV5UpdateCustomProtocolFee)
				if err := _YRegistryV5.contract.UnpackLog(event, "UpdateCustomProtocolFee", log); err != nil {
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

// ParseUpdateCustomProtocolFee is a log parse operation binding the contract event 0x96d6cc624354ffe5a7207dc2dcc152e58e23ac8df9c96943f3cfb10ea4c140ac.
//
// Solidity: event UpdateCustomProtocolFee(address indexed vault, uint16 new_custom_protocol_fee)
func (_YRegistryV5 *YRegistryV5Filterer) ParseUpdateCustomProtocolFee(log types.Log) (*YRegistryV5UpdateCustomProtocolFee, error) {
	event := new(YRegistryV5UpdateCustomProtocolFee)
	if err := _YRegistryV5.contract.UnpackLog(event, "UpdateCustomProtocolFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YRegistryV5UpdateGovernanceIterator is returned from FilterUpdateGovernance and is used to iterate over the raw logs and unpacked data for UpdateGovernance events raised by the YRegistryV5 contract.
type YRegistryV5UpdateGovernanceIterator struct {
	Event *YRegistryV5UpdateGovernance // Event containing the contract specifics and raw log

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
func (it *YRegistryV5UpdateGovernanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YRegistryV5UpdateGovernance)
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
		it.Event = new(YRegistryV5UpdateGovernance)
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
func (it *YRegistryV5UpdateGovernanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YRegistryV5UpdateGovernanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YRegistryV5UpdateGovernance represents a UpdateGovernance event raised by the YRegistryV5 contract.
type YRegistryV5UpdateGovernance struct {
	Governance common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateGovernance is a free log retrieval operation binding the contract event 0x8d55d160c0009eb3d739442df0a3ca089ed64378bfac017e7ddad463f9815b87.
//
// Solidity: event UpdateGovernance(address indexed governance)
func (_YRegistryV5 *YRegistryV5Filterer) FilterUpdateGovernance(opts *bind.FilterOpts, governance []common.Address) (*YRegistryV5UpdateGovernanceIterator, error) {

	var governanceRule []interface{}
	for _, governanceItem := range governance {
		governanceRule = append(governanceRule, governanceItem)
	}

	logs, sub, err := _YRegistryV5.contract.FilterLogs(opts, "UpdateGovernance", governanceRule)
	if err != nil {
		return nil, err
	}
	return &YRegistryV5UpdateGovernanceIterator{contract: _YRegistryV5.contract, event: "UpdateGovernance", logs: logs, sub: sub}, nil
}

// WatchUpdateGovernance is a free log subscription operation binding the contract event 0x8d55d160c0009eb3d739442df0a3ca089ed64378bfac017e7ddad463f9815b87.
//
// Solidity: event UpdateGovernance(address indexed governance)
func (_YRegistryV5 *YRegistryV5Filterer) WatchUpdateGovernance(opts *bind.WatchOpts, sink chan<- *YRegistryV5UpdateGovernance, governance []common.Address) (event.Subscription, error) {

	var governanceRule []interface{}
	for _, governanceItem := range governance {
		governanceRule = append(governanceRule, governanceItem)
	}

	logs, sub, err := _YRegistryV5.contract.WatchLogs(opts, "UpdateGovernance", governanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YRegistryV5UpdateGovernance)
				if err := _YRegistryV5.contract.UnpackLog(event, "UpdateGovernance", log); err != nil {
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

// ParseUpdateGovernance is a log parse operation binding the contract event 0x8d55d160c0009eb3d739442df0a3ca089ed64378bfac017e7ddad463f9815b87.
//
// Solidity: event UpdateGovernance(address indexed governance)
func (_YRegistryV5 *YRegistryV5Filterer) ParseUpdateGovernance(log types.Log) (*YRegistryV5UpdateGovernance, error) {
	event := new(YRegistryV5UpdateGovernance)
	if err := _YRegistryV5.contract.UnpackLog(event, "UpdateGovernance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YRegistryV5UpdateProtocolFeeBpsIterator is returned from FilterUpdateProtocolFeeBps and is used to iterate over the raw logs and unpacked data for UpdateProtocolFeeBps events raised by the YRegistryV5 contract.
type YRegistryV5UpdateProtocolFeeBpsIterator struct {
	Event *YRegistryV5UpdateProtocolFeeBps // Event containing the contract specifics and raw log

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
func (it *YRegistryV5UpdateProtocolFeeBpsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YRegistryV5UpdateProtocolFeeBps)
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
		it.Event = new(YRegistryV5UpdateProtocolFeeBps)
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
func (it *YRegistryV5UpdateProtocolFeeBpsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YRegistryV5UpdateProtocolFeeBpsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YRegistryV5UpdateProtocolFeeBps represents a UpdateProtocolFeeBps event raised by the YRegistryV5 contract.
type YRegistryV5UpdateProtocolFeeBps struct {
	OldFeeBps uint16
	NewFeeBps uint16
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateProtocolFeeBps is a free log retrieval operation binding the contract event 0x678d2b2fe79c193f6c2c18d7515e339afcbd73fcfb360b1d0fbadae07342e051.
//
// Solidity: event UpdateProtocolFeeBps(uint16 old_fee_bps, uint16 new_fee_bps)
func (_YRegistryV5 *YRegistryV5Filterer) FilterUpdateProtocolFeeBps(opts *bind.FilterOpts) (*YRegistryV5UpdateProtocolFeeBpsIterator, error) {

	logs, sub, err := _YRegistryV5.contract.FilterLogs(opts, "UpdateProtocolFeeBps")
	if err != nil {
		return nil, err
	}
	return &YRegistryV5UpdateProtocolFeeBpsIterator{contract: _YRegistryV5.contract, event: "UpdateProtocolFeeBps", logs: logs, sub: sub}, nil
}

// WatchUpdateProtocolFeeBps is a free log subscription operation binding the contract event 0x678d2b2fe79c193f6c2c18d7515e339afcbd73fcfb360b1d0fbadae07342e051.
//
// Solidity: event UpdateProtocolFeeBps(uint16 old_fee_bps, uint16 new_fee_bps)
func (_YRegistryV5 *YRegistryV5Filterer) WatchUpdateProtocolFeeBps(opts *bind.WatchOpts, sink chan<- *YRegistryV5UpdateProtocolFeeBps) (event.Subscription, error) {

	logs, sub, err := _YRegistryV5.contract.WatchLogs(opts, "UpdateProtocolFeeBps")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YRegistryV5UpdateProtocolFeeBps)
				if err := _YRegistryV5.contract.UnpackLog(event, "UpdateProtocolFeeBps", log); err != nil {
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

// ParseUpdateProtocolFeeBps is a log parse operation binding the contract event 0x678d2b2fe79c193f6c2c18d7515e339afcbd73fcfb360b1d0fbadae07342e051.
//
// Solidity: event UpdateProtocolFeeBps(uint16 old_fee_bps, uint16 new_fee_bps)
func (_YRegistryV5 *YRegistryV5Filterer) ParseUpdateProtocolFeeBps(log types.Log) (*YRegistryV5UpdateProtocolFeeBps, error) {
	event := new(YRegistryV5UpdateProtocolFeeBps)
	if err := _YRegistryV5.contract.UnpackLog(event, "UpdateProtocolFeeBps", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YRegistryV5UpdateProtocolFeeRecipientIterator is returned from FilterUpdateProtocolFeeRecipient and is used to iterate over the raw logs and unpacked data for UpdateProtocolFeeRecipient events raised by the YRegistryV5 contract.
type YRegistryV5UpdateProtocolFeeRecipientIterator struct {
	Event *YRegistryV5UpdateProtocolFeeRecipient // Event containing the contract specifics and raw log

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
func (it *YRegistryV5UpdateProtocolFeeRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YRegistryV5UpdateProtocolFeeRecipient)
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
		it.Event = new(YRegistryV5UpdateProtocolFeeRecipient)
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
func (it *YRegistryV5UpdateProtocolFeeRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YRegistryV5UpdateProtocolFeeRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YRegistryV5UpdateProtocolFeeRecipient represents a UpdateProtocolFeeRecipient event raised by the YRegistryV5 contract.
type YRegistryV5UpdateProtocolFeeRecipient struct {
	OldFeeRecipient common.Address
	NewFeeRecipient common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateProtocolFeeRecipient is a free log retrieval operation binding the contract event 0x6af4e38beb02e4b110090dd85c5adfb341e2278b905068773762fe4666e5db7a.
//
// Solidity: event UpdateProtocolFeeRecipient(address indexed old_fee_recipient, address indexed new_fee_recipient)
func (_YRegistryV5 *YRegistryV5Filterer) FilterUpdateProtocolFeeRecipient(opts *bind.FilterOpts, old_fee_recipient []common.Address, new_fee_recipient []common.Address) (*YRegistryV5UpdateProtocolFeeRecipientIterator, error) {

	var old_fee_recipientRule []interface{}
	for _, old_fee_recipientItem := range old_fee_recipient {
		old_fee_recipientRule = append(old_fee_recipientRule, old_fee_recipientItem)
	}
	var new_fee_recipientRule []interface{}
	for _, new_fee_recipientItem := range new_fee_recipient {
		new_fee_recipientRule = append(new_fee_recipientRule, new_fee_recipientItem)
	}

	logs, sub, err := _YRegistryV5.contract.FilterLogs(opts, "UpdateProtocolFeeRecipient", old_fee_recipientRule, new_fee_recipientRule)
	if err != nil {
		return nil, err
	}
	return &YRegistryV5UpdateProtocolFeeRecipientIterator{contract: _YRegistryV5.contract, event: "UpdateProtocolFeeRecipient", logs: logs, sub: sub}, nil
}

// WatchUpdateProtocolFeeRecipient is a free log subscription operation binding the contract event 0x6af4e38beb02e4b110090dd85c5adfb341e2278b905068773762fe4666e5db7a.
//
// Solidity: event UpdateProtocolFeeRecipient(address indexed old_fee_recipient, address indexed new_fee_recipient)
func (_YRegistryV5 *YRegistryV5Filterer) WatchUpdateProtocolFeeRecipient(opts *bind.WatchOpts, sink chan<- *YRegistryV5UpdateProtocolFeeRecipient, old_fee_recipient []common.Address, new_fee_recipient []common.Address) (event.Subscription, error) {

	var old_fee_recipientRule []interface{}
	for _, old_fee_recipientItem := range old_fee_recipient {
		old_fee_recipientRule = append(old_fee_recipientRule, old_fee_recipientItem)
	}
	var new_fee_recipientRule []interface{}
	for _, new_fee_recipientItem := range new_fee_recipient {
		new_fee_recipientRule = append(new_fee_recipientRule, new_fee_recipientItem)
	}

	logs, sub, err := _YRegistryV5.contract.WatchLogs(opts, "UpdateProtocolFeeRecipient", old_fee_recipientRule, new_fee_recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YRegistryV5UpdateProtocolFeeRecipient)
				if err := _YRegistryV5.contract.UnpackLog(event, "UpdateProtocolFeeRecipient", log); err != nil {
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

// ParseUpdateProtocolFeeRecipient is a log parse operation binding the contract event 0x6af4e38beb02e4b110090dd85c5adfb341e2278b905068773762fe4666e5db7a.
//
// Solidity: event UpdateProtocolFeeRecipient(address indexed old_fee_recipient, address indexed new_fee_recipient)
func (_YRegistryV5 *YRegistryV5Filterer) ParseUpdateProtocolFeeRecipient(log types.Log) (*YRegistryV5UpdateProtocolFeeRecipient, error) {
	event := new(YRegistryV5UpdateProtocolFeeRecipient)
	if err := _YRegistryV5.contract.UnpackLog(event, "UpdateProtocolFeeRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
