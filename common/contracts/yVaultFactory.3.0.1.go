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

// YFactory301MetaData contains all meta data concerning the YFactory301 contract.
var YFactory301MetaData = &bind.MetaData{
	ABI: "[{\"name\":\"NewVault\",\"inputs\":[{\"name\":\"vault_address\",\"type\":\"address\",\"indexed\":true},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateProtocolFeeBps\",\"inputs\":[{\"name\":\"old_fee_bps\",\"type\":\"uint16\",\"indexed\":false},{\"name\":\"new_fee_bps\",\"type\":\"uint16\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateProtocolFeeRecipient\",\"inputs\":[{\"name\":\"old_fee_recipient\",\"type\":\"address\",\"indexed\":true},{\"name\":\"new_fee_recipient\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateCustomProtocolFee\",\"inputs\":[{\"name\":\"vault\",\"type\":\"address\",\"indexed\":true},{\"name\":\"new_custom_protocol_fee\",\"type\":\"uint16\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemovedCustomProtocolFee\",\"inputs\":[{\"name\":\"vault\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"FactoryShutdown\",\"inputs\":[],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateGovernance\",\"inputs\":[{\"name\":\"governance\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewPendingGovernance\",\"inputs\":[{\"name\":\"pending_governance\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"vault_blueprint\",\"type\":\"address\"},{\"name\":\"governance\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deploy_new_vault\",\"inputs\":[{\"name\":\"asset\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"role_manager\",\"type\":\"address\"},{\"name\":\"profit_max_unlock_time\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"vault_blueprint\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"apiVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"protocol_fee_config\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"components\":[{\"name\":\"fee_bps\",\"type\":\"uint16\"},{\"name\":\"fee_recipient\",\"type\":\"address\"}]}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_protocol_fee_bps\",\"inputs\":[{\"name\":\"new_protocol_fee_bps\",\"type\":\"uint16\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_protocol_fee_recipient\",\"inputs\":[{\"name\":\"new_protocol_fee_recipient\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_custom_protocol_fee_bps\",\"inputs\":[{\"name\":\"vault\",\"type\":\"address\"},{\"name\":\"new_custom_protocol_fee\",\"type\":\"uint16\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_custom_protocol_fee\",\"inputs\":[{\"name\":\"vault\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"shutdown_factory\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_governance\",\"inputs\":[{\"name\":\"new_governance\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"accept_governance\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"shutdown\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"governance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"pending_governance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"default_protocol_fee_config\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"components\":[{\"name\":\"fee_bps\",\"type\":\"uint16\"},{\"name\":\"fee_recipient\",\"type\":\"address\"}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"custom_protocol_fee\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"use_custom_protocol_fee\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]}]",
}

// YFactory301ABI is the input ABI used to generate the binding from.
// Deprecated: Use YFactory301MetaData.ABI instead.
var YFactory301ABI = YFactory301MetaData.ABI

// YFactory301 is an auto generated Go binding around an Ethereum contract.
type YFactory301 struct {
	YFactory301Caller     // Read-only binding to the contract
	YFactory301Transactor // Write-only binding to the contract
	YFactory301Filterer   // Log filterer for contract events
}

// YFactory301Caller is an auto generated read-only Go binding around an Ethereum contract.
type YFactory301Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YFactory301Transactor is an auto generated write-only Go binding around an Ethereum contract.
type YFactory301Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YFactory301Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YFactory301Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YFactory301Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YFactory301Session struct {
	Contract     *YFactory301      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YFactory301CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YFactory301CallerSession struct {
	Contract *YFactory301Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// YFactory301TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YFactory301TransactorSession struct {
	Contract     *YFactory301Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// YFactory301Raw is an auto generated low-level Go binding around an Ethereum contract.
type YFactory301Raw struct {
	Contract *YFactory301 // Generic contract binding to access the raw methods on
}

// YFactory301CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YFactory301CallerRaw struct {
	Contract *YFactory301Caller // Generic read-only contract binding to access the raw methods on
}

// YFactory301TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YFactory301TransactorRaw struct {
	Contract *YFactory301Transactor // Generic write-only contract binding to access the raw methods on
}

// NewYFactory301 creates a new instance of YFactory301, bound to a specific deployed contract.
func NewYFactory301(address common.Address, backend bind.ContractBackend) (*YFactory301, error) {
	contract, err := bindYFactory301(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YFactory301{YFactory301Caller: YFactory301Caller{contract: contract}, YFactory301Transactor: YFactory301Transactor{contract: contract}, YFactory301Filterer: YFactory301Filterer{contract: contract}}, nil
}

// NewYFactory301Caller creates a new read-only instance of YFactory301, bound to a specific deployed contract.
func NewYFactory301Caller(address common.Address, caller bind.ContractCaller) (*YFactory301Caller, error) {
	contract, err := bindYFactory301(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YFactory301Caller{contract: contract}, nil
}

// NewYFactory301Transactor creates a new write-only instance of YFactory301, bound to a specific deployed contract.
func NewYFactory301Transactor(address common.Address, transactor bind.ContractTransactor) (*YFactory301Transactor, error) {
	contract, err := bindYFactory301(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YFactory301Transactor{contract: contract}, nil
}

// NewYFactory301Filterer creates a new log filterer instance of YFactory301, bound to a specific deployed contract.
func NewYFactory301Filterer(address common.Address, filterer bind.ContractFilterer) (*YFactory301Filterer, error) {
	contract, err := bindYFactory301(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YFactory301Filterer{contract: contract}, nil
}

// bindYFactory301 binds a generic wrapper to an already deployed contract.
func bindYFactory301(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(YFactory301ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YFactory301 *YFactory301Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YFactory301.Contract.YFactory301Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YFactory301 *YFactory301Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YFactory301.Contract.YFactory301Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YFactory301 *YFactory301Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YFactory301.Contract.YFactory301Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YFactory301 *YFactory301CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YFactory301.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YFactory301 *YFactory301TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YFactory301.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YFactory301 *YFactory301TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YFactory301.Contract.contract.Transact(opts, method, params...)
}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() view returns(string)
func (_YFactory301 *YFactory301Caller) ApiVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _YFactory301.contract.Call(opts, &out, "apiVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() view returns(string)
func (_YFactory301 *YFactory301Session) ApiVersion() (string, error) {
	return _YFactory301.Contract.ApiVersion(&_YFactory301.CallOpts)
}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() view returns(string)
func (_YFactory301 *YFactory301CallerSession) ApiVersion() (string, error) {
	return _YFactory301.Contract.ApiVersion(&_YFactory301.CallOpts)
}

// CustomProtocolFee is a free data retrieval call binding the contract method 0xcbe28663.
//
// Solidity: function custom_protocol_fee(address arg0) view returns(uint16)
func (_YFactory301 *YFactory301Caller) CustomProtocolFee(opts *bind.CallOpts, arg0 common.Address) (uint16, error) {
	var out []interface{}
	err := _YFactory301.contract.Call(opts, &out, "custom_protocol_fee", arg0)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// CustomProtocolFee is a free data retrieval call binding the contract method 0xcbe28663.
//
// Solidity: function custom_protocol_fee(address arg0) view returns(uint16)
func (_YFactory301 *YFactory301Session) CustomProtocolFee(arg0 common.Address) (uint16, error) {
	return _YFactory301.Contract.CustomProtocolFee(&_YFactory301.CallOpts, arg0)
}

// CustomProtocolFee is a free data retrieval call binding the contract method 0xcbe28663.
//
// Solidity: function custom_protocol_fee(address arg0) view returns(uint16)
func (_YFactory301 *YFactory301CallerSession) CustomProtocolFee(arg0 common.Address) (uint16, error) {
	return _YFactory301.Contract.CustomProtocolFee(&_YFactory301.CallOpts, arg0)
}

// DefaultProtocolFeeConfig is a free data retrieval call binding the contract method 0x97ad2ecc.
//
// Solidity: function default_protocol_fee_config() view returns((uint16,address))
func (_YFactory301 *YFactory301Caller) DefaultProtocolFeeConfig(opts *bind.CallOpts) (Struct0, error) {
	var out []interface{}
	err := _YFactory301.contract.Call(opts, &out, "default_protocol_fee_config")

	if err != nil {
		return *new(Struct0), err
	}

	out0 := *abi.ConvertType(out[0], new(Struct0)).(*Struct0)

	return out0, err

}

// DefaultProtocolFeeConfig is a free data retrieval call binding the contract method 0x97ad2ecc.
//
// Solidity: function default_protocol_fee_config() view returns((uint16,address))
func (_YFactory301 *YFactory301Session) DefaultProtocolFeeConfig() (Struct0, error) {
	return _YFactory301.Contract.DefaultProtocolFeeConfig(&_YFactory301.CallOpts)
}

// DefaultProtocolFeeConfig is a free data retrieval call binding the contract method 0x97ad2ecc.
//
// Solidity: function default_protocol_fee_config() view returns((uint16,address))
func (_YFactory301 *YFactory301CallerSession) DefaultProtocolFeeConfig() (Struct0, error) {
	return _YFactory301.Contract.DefaultProtocolFeeConfig(&_YFactory301.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_YFactory301 *YFactory301Caller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YFactory301.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_YFactory301 *YFactory301Session) Governance() (common.Address, error) {
	return _YFactory301.Contract.Governance(&_YFactory301.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_YFactory301 *YFactory301CallerSession) Governance() (common.Address, error) {
	return _YFactory301.Contract.Governance(&_YFactory301.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YFactory301 *YFactory301Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _YFactory301.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YFactory301 *YFactory301Session) Name() (string, error) {
	return _YFactory301.Contract.Name(&_YFactory301.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YFactory301 *YFactory301CallerSession) Name() (string, error) {
	return _YFactory301.Contract.Name(&_YFactory301.CallOpts)
}

// PendingGovernance is a free data retrieval call binding the contract method 0xc66eb0a2.
//
// Solidity: function pending_governance() view returns(address)
func (_YFactory301 *YFactory301Caller) PendingGovernance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YFactory301.contract.Call(opts, &out, "pending_governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingGovernance is a free data retrieval call binding the contract method 0xc66eb0a2.
//
// Solidity: function pending_governance() view returns(address)
func (_YFactory301 *YFactory301Session) PendingGovernance() (common.Address, error) {
	return _YFactory301.Contract.PendingGovernance(&_YFactory301.CallOpts)
}

// PendingGovernance is a free data retrieval call binding the contract method 0xc66eb0a2.
//
// Solidity: function pending_governance() view returns(address)
func (_YFactory301 *YFactory301CallerSession) PendingGovernance() (common.Address, error) {
	return _YFactory301.Contract.PendingGovernance(&_YFactory301.CallOpts)
}

// ProtocolFeeConfig is a free data retrieval call binding the contract method 0x5153b199.
//
// Solidity: function protocol_fee_config() view returns((uint16,address))
func (_YFactory301 *YFactory301Caller) ProtocolFeeConfig(opts *bind.CallOpts) (Struct0, error) {
	var out []interface{}
	err := _YFactory301.contract.Call(opts, &out, "protocol_fee_config")

	if err != nil {
		return *new(Struct0), err
	}

	out0 := *abi.ConvertType(out[0], new(Struct0)).(*Struct0)

	return out0, err

}

// ProtocolFeeConfig is a free data retrieval call binding the contract method 0x5153b199.
//
// Solidity: function protocol_fee_config() view returns((uint16,address))
func (_YFactory301 *YFactory301Session) ProtocolFeeConfig() (Struct0, error) {
	return _YFactory301.Contract.ProtocolFeeConfig(&_YFactory301.CallOpts)
}

// ProtocolFeeConfig is a free data retrieval call binding the contract method 0x5153b199.
//
// Solidity: function protocol_fee_config() view returns((uint16,address))
func (_YFactory301 *YFactory301CallerSession) ProtocolFeeConfig() (Struct0, error) {
	return _YFactory301.Contract.ProtocolFeeConfig(&_YFactory301.CallOpts)
}

// Shutdown is a free data retrieval call binding the contract method 0xfc0e74d1.
//
// Solidity: function shutdown() view returns(bool)
func (_YFactory301 *YFactory301Caller) Shutdown(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _YFactory301.contract.Call(opts, &out, "shutdown")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Shutdown is a free data retrieval call binding the contract method 0xfc0e74d1.
//
// Solidity: function shutdown() view returns(bool)
func (_YFactory301 *YFactory301Session) Shutdown() (bool, error) {
	return _YFactory301.Contract.Shutdown(&_YFactory301.CallOpts)
}

// Shutdown is a free data retrieval call binding the contract method 0xfc0e74d1.
//
// Solidity: function shutdown() view returns(bool)
func (_YFactory301 *YFactory301CallerSession) Shutdown() (bool, error) {
	return _YFactory301.Contract.Shutdown(&_YFactory301.CallOpts)
}

// UseCustomProtocolFee is a free data retrieval call binding the contract method 0xe94860d8.
//
// Solidity: function use_custom_protocol_fee(address arg0) view returns(bool)
func (_YFactory301 *YFactory301Caller) UseCustomProtocolFee(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _YFactory301.contract.Call(opts, &out, "use_custom_protocol_fee", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UseCustomProtocolFee is a free data retrieval call binding the contract method 0xe94860d8.
//
// Solidity: function use_custom_protocol_fee(address arg0) view returns(bool)
func (_YFactory301 *YFactory301Session) UseCustomProtocolFee(arg0 common.Address) (bool, error) {
	return _YFactory301.Contract.UseCustomProtocolFee(&_YFactory301.CallOpts, arg0)
}

// UseCustomProtocolFee is a free data retrieval call binding the contract method 0xe94860d8.
//
// Solidity: function use_custom_protocol_fee(address arg0) view returns(bool)
func (_YFactory301 *YFactory301CallerSession) UseCustomProtocolFee(arg0 common.Address) (bool, error) {
	return _YFactory301.Contract.UseCustomProtocolFee(&_YFactory301.CallOpts, arg0)
}

// VaultBlueprint is a free data retrieval call binding the contract method 0x05f16b0f.
//
// Solidity: function vault_blueprint() view returns(address)
func (_YFactory301 *YFactory301Caller) VaultBlueprint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YFactory301.contract.Call(opts, &out, "vault_blueprint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultBlueprint is a free data retrieval call binding the contract method 0x05f16b0f.
//
// Solidity: function vault_blueprint() view returns(address)
func (_YFactory301 *YFactory301Session) VaultBlueprint() (common.Address, error) {
	return _YFactory301.Contract.VaultBlueprint(&_YFactory301.CallOpts)
}

// VaultBlueprint is a free data retrieval call binding the contract method 0x05f16b0f.
//
// Solidity: function vault_blueprint() view returns(address)
func (_YFactory301 *YFactory301CallerSession) VaultBlueprint() (common.Address, error) {
	return _YFactory301.Contract.VaultBlueprint(&_YFactory301.CallOpts)
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0xa7dbff3e.
//
// Solidity: function accept_governance() returns()
func (_YFactory301 *YFactory301Transactor) AcceptGovernance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YFactory301.contract.Transact(opts, "accept_governance")
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0xa7dbff3e.
//
// Solidity: function accept_governance() returns()
func (_YFactory301 *YFactory301Session) AcceptGovernance() (*types.Transaction, error) {
	return _YFactory301.Contract.AcceptGovernance(&_YFactory301.TransactOpts)
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0xa7dbff3e.
//
// Solidity: function accept_governance() returns()
func (_YFactory301 *YFactory301TransactorSession) AcceptGovernance() (*types.Transaction, error) {
	return _YFactory301.Contract.AcceptGovernance(&_YFactory301.TransactOpts)
}

// DeployNewVault is a paid mutator transaction binding the contract method 0xb4aeee77.
//
// Solidity: function deploy_new_vault(address asset, string name, string symbol, address role_manager, uint256 profit_max_unlock_time) returns(address)
func (_YFactory301 *YFactory301Transactor) DeployNewVault(opts *bind.TransactOpts, asset common.Address, name string, symbol string, role_manager common.Address, profit_max_unlock_time *big.Int) (*types.Transaction, error) {
	return _YFactory301.contract.Transact(opts, "deploy_new_vault", asset, name, symbol, role_manager, profit_max_unlock_time)
}

// DeployNewVault is a paid mutator transaction binding the contract method 0xb4aeee77.
//
// Solidity: function deploy_new_vault(address asset, string name, string symbol, address role_manager, uint256 profit_max_unlock_time) returns(address)
func (_YFactory301 *YFactory301Session) DeployNewVault(asset common.Address, name string, symbol string, role_manager common.Address, profit_max_unlock_time *big.Int) (*types.Transaction, error) {
	return _YFactory301.Contract.DeployNewVault(&_YFactory301.TransactOpts, asset, name, symbol, role_manager, profit_max_unlock_time)
}

// DeployNewVault is a paid mutator transaction binding the contract method 0xb4aeee77.
//
// Solidity: function deploy_new_vault(address asset, string name, string symbol, address role_manager, uint256 profit_max_unlock_time) returns(address)
func (_YFactory301 *YFactory301TransactorSession) DeployNewVault(asset common.Address, name string, symbol string, role_manager common.Address, profit_max_unlock_time *big.Int) (*types.Transaction, error) {
	return _YFactory301.Contract.DeployNewVault(&_YFactory301.TransactOpts, asset, name, symbol, role_manager, profit_max_unlock_time)
}

// RemoveCustomProtocolFee is a paid mutator transaction binding the contract method 0x11a3a434.
//
// Solidity: function remove_custom_protocol_fee(address vault) returns()
func (_YFactory301 *YFactory301Transactor) RemoveCustomProtocolFee(opts *bind.TransactOpts, vault common.Address) (*types.Transaction, error) {
	return _YFactory301.contract.Transact(opts, "remove_custom_protocol_fee", vault)
}

// RemoveCustomProtocolFee is a paid mutator transaction binding the contract method 0x11a3a434.
//
// Solidity: function remove_custom_protocol_fee(address vault) returns()
func (_YFactory301 *YFactory301Session) RemoveCustomProtocolFee(vault common.Address) (*types.Transaction, error) {
	return _YFactory301.Contract.RemoveCustomProtocolFee(&_YFactory301.TransactOpts, vault)
}

// RemoveCustomProtocolFee is a paid mutator transaction binding the contract method 0x11a3a434.
//
// Solidity: function remove_custom_protocol_fee(address vault) returns()
func (_YFactory301 *YFactory301TransactorSession) RemoveCustomProtocolFee(vault common.Address) (*types.Transaction, error) {
	return _YFactory301.Contract.RemoveCustomProtocolFee(&_YFactory301.TransactOpts, vault)
}

// SetCustomProtocolFeeBps is a paid mutator transaction binding the contract method 0xb5a71e07.
//
// Solidity: function set_custom_protocol_fee_bps(address vault, uint16 new_custom_protocol_fee) returns()
func (_YFactory301 *YFactory301Transactor) SetCustomProtocolFeeBps(opts *bind.TransactOpts, vault common.Address, new_custom_protocol_fee uint16) (*types.Transaction, error) {
	return _YFactory301.contract.Transact(opts, "set_custom_protocol_fee_bps", vault, new_custom_protocol_fee)
}

// SetCustomProtocolFeeBps is a paid mutator transaction binding the contract method 0xb5a71e07.
//
// Solidity: function set_custom_protocol_fee_bps(address vault, uint16 new_custom_protocol_fee) returns()
func (_YFactory301 *YFactory301Session) SetCustomProtocolFeeBps(vault common.Address, new_custom_protocol_fee uint16) (*types.Transaction, error) {
	return _YFactory301.Contract.SetCustomProtocolFeeBps(&_YFactory301.TransactOpts, vault, new_custom_protocol_fee)
}

// SetCustomProtocolFeeBps is a paid mutator transaction binding the contract method 0xb5a71e07.
//
// Solidity: function set_custom_protocol_fee_bps(address vault, uint16 new_custom_protocol_fee) returns()
func (_YFactory301 *YFactory301TransactorSession) SetCustomProtocolFeeBps(vault common.Address, new_custom_protocol_fee uint16) (*types.Transaction, error) {
	return _YFactory301.Contract.SetCustomProtocolFeeBps(&_YFactory301.TransactOpts, vault, new_custom_protocol_fee)
}

// SetGovernance is a paid mutator transaction binding the contract method 0x070313fa.
//
// Solidity: function set_governance(address new_governance) returns()
func (_YFactory301 *YFactory301Transactor) SetGovernance(opts *bind.TransactOpts, new_governance common.Address) (*types.Transaction, error) {
	return _YFactory301.contract.Transact(opts, "set_governance", new_governance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0x070313fa.
//
// Solidity: function set_governance(address new_governance) returns()
func (_YFactory301 *YFactory301Session) SetGovernance(new_governance common.Address) (*types.Transaction, error) {
	return _YFactory301.Contract.SetGovernance(&_YFactory301.TransactOpts, new_governance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0x070313fa.
//
// Solidity: function set_governance(address new_governance) returns()
func (_YFactory301 *YFactory301TransactorSession) SetGovernance(new_governance common.Address) (*types.Transaction, error) {
	return _YFactory301.Contract.SetGovernance(&_YFactory301.TransactOpts, new_governance)
}

// SetProtocolFeeBps is a paid mutator transaction binding the contract method 0x62fbf603.
//
// Solidity: function set_protocol_fee_bps(uint16 new_protocol_fee_bps) returns()
func (_YFactory301 *YFactory301Transactor) SetProtocolFeeBps(opts *bind.TransactOpts, new_protocol_fee_bps uint16) (*types.Transaction, error) {
	return _YFactory301.contract.Transact(opts, "set_protocol_fee_bps", new_protocol_fee_bps)
}

// SetProtocolFeeBps is a paid mutator transaction binding the contract method 0x62fbf603.
//
// Solidity: function set_protocol_fee_bps(uint16 new_protocol_fee_bps) returns()
func (_YFactory301 *YFactory301Session) SetProtocolFeeBps(new_protocol_fee_bps uint16) (*types.Transaction, error) {
	return _YFactory301.Contract.SetProtocolFeeBps(&_YFactory301.TransactOpts, new_protocol_fee_bps)
}

// SetProtocolFeeBps is a paid mutator transaction binding the contract method 0x62fbf603.
//
// Solidity: function set_protocol_fee_bps(uint16 new_protocol_fee_bps) returns()
func (_YFactory301 *YFactory301TransactorSession) SetProtocolFeeBps(new_protocol_fee_bps uint16) (*types.Transaction, error) {
	return _YFactory301.Contract.SetProtocolFeeBps(&_YFactory301.TransactOpts, new_protocol_fee_bps)
}

// SetProtocolFeeRecipient is a paid mutator transaction binding the contract method 0xf8ebccea.
//
// Solidity: function set_protocol_fee_recipient(address new_protocol_fee_recipient) returns()
func (_YFactory301 *YFactory301Transactor) SetProtocolFeeRecipient(opts *bind.TransactOpts, new_protocol_fee_recipient common.Address) (*types.Transaction, error) {
	return _YFactory301.contract.Transact(opts, "set_protocol_fee_recipient", new_protocol_fee_recipient)
}

// SetProtocolFeeRecipient is a paid mutator transaction binding the contract method 0xf8ebccea.
//
// Solidity: function set_protocol_fee_recipient(address new_protocol_fee_recipient) returns()
func (_YFactory301 *YFactory301Session) SetProtocolFeeRecipient(new_protocol_fee_recipient common.Address) (*types.Transaction, error) {
	return _YFactory301.Contract.SetProtocolFeeRecipient(&_YFactory301.TransactOpts, new_protocol_fee_recipient)
}

// SetProtocolFeeRecipient is a paid mutator transaction binding the contract method 0xf8ebccea.
//
// Solidity: function set_protocol_fee_recipient(address new_protocol_fee_recipient) returns()
func (_YFactory301 *YFactory301TransactorSession) SetProtocolFeeRecipient(new_protocol_fee_recipient common.Address) (*types.Transaction, error) {
	return _YFactory301.Contract.SetProtocolFeeRecipient(&_YFactory301.TransactOpts, new_protocol_fee_recipient)
}

// ShutdownFactory is a paid mutator transaction binding the contract method 0x365adba6.
//
// Solidity: function shutdown_factory() returns()
func (_YFactory301 *YFactory301Transactor) ShutdownFactory(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YFactory301.contract.Transact(opts, "shutdown_factory")
}

// ShutdownFactory is a paid mutator transaction binding the contract method 0x365adba6.
//
// Solidity: function shutdown_factory() returns()
func (_YFactory301 *YFactory301Session) ShutdownFactory() (*types.Transaction, error) {
	return _YFactory301.Contract.ShutdownFactory(&_YFactory301.TransactOpts)
}

// ShutdownFactory is a paid mutator transaction binding the contract method 0x365adba6.
//
// Solidity: function shutdown_factory() returns()
func (_YFactory301 *YFactory301TransactorSession) ShutdownFactory() (*types.Transaction, error) {
	return _YFactory301.Contract.ShutdownFactory(&_YFactory301.TransactOpts)
}

// YFactory301FactoryShutdownIterator is returned from FilterFactoryShutdown and is used to iterate over the raw logs and unpacked data for FactoryShutdown events raised by the YFactory301 contract.
type YFactory301FactoryShutdownIterator struct {
	Event *YFactory301FactoryShutdown // Event containing the contract specifics and raw log

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
func (it *YFactory301FactoryShutdownIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YFactory301FactoryShutdown)
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
		it.Event = new(YFactory301FactoryShutdown)
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
func (it *YFactory301FactoryShutdownIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YFactory301FactoryShutdownIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YFactory301FactoryShutdown represents a FactoryShutdown event raised by the YFactory301 contract.
type YFactory301FactoryShutdown struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFactoryShutdown is a free log retrieval operation binding the contract event 0xc643193a97fc0e18d69c95e1c034b91f51fa164ba8ea68dfb6dd98568b9bc96b.
//
// Solidity: event FactoryShutdown()
func (_YFactory301 *YFactory301Filterer) FilterFactoryShutdown(opts *bind.FilterOpts) (*YFactory301FactoryShutdownIterator, error) {

	logs, sub, err := _YFactory301.contract.FilterLogs(opts, "FactoryShutdown")
	if err != nil {
		return nil, err
	}
	return &YFactory301FactoryShutdownIterator{contract: _YFactory301.contract, event: "FactoryShutdown", logs: logs, sub: sub}, nil
}

// WatchFactoryShutdown is a free log subscription operation binding the contract event 0xc643193a97fc0e18d69c95e1c034b91f51fa164ba8ea68dfb6dd98568b9bc96b.
//
// Solidity: event FactoryShutdown()
func (_YFactory301 *YFactory301Filterer) WatchFactoryShutdown(opts *bind.WatchOpts, sink chan<- *YFactory301FactoryShutdown) (event.Subscription, error) {

	logs, sub, err := _YFactory301.contract.WatchLogs(opts, "FactoryShutdown")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YFactory301FactoryShutdown)
				if err := _YFactory301.contract.UnpackLog(event, "FactoryShutdown", log); err != nil {
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
func (_YFactory301 *YFactory301Filterer) ParseFactoryShutdown(log types.Log) (*YFactory301FactoryShutdown, error) {
	event := new(YFactory301FactoryShutdown)
	if err := _YFactory301.contract.UnpackLog(event, "FactoryShutdown", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YFactory301NewPendingGovernanceIterator is returned from FilterNewPendingGovernance and is used to iterate over the raw logs and unpacked data for NewPendingGovernance events raised by the YFactory301 contract.
type YFactory301NewPendingGovernanceIterator struct {
	Event *YFactory301NewPendingGovernance // Event containing the contract specifics and raw log

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
func (it *YFactory301NewPendingGovernanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YFactory301NewPendingGovernance)
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
		it.Event = new(YFactory301NewPendingGovernance)
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
func (it *YFactory301NewPendingGovernanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YFactory301NewPendingGovernanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YFactory301NewPendingGovernance represents a NewPendingGovernance event raised by the YFactory301 contract.
type YFactory301NewPendingGovernance struct {
	PendingGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewPendingGovernance is a free log retrieval operation binding the contract event 0x90ad4c550d25bd23af61db38d1ff8671b89edaaa0bca0fc36bac5084ecc120bd.
//
// Solidity: event NewPendingGovernance(address indexed pending_governance)
func (_YFactory301 *YFactory301Filterer) FilterNewPendingGovernance(opts *bind.FilterOpts, pending_governance []common.Address) (*YFactory301NewPendingGovernanceIterator, error) {

	var pending_governanceRule []interface{}
	for _, pending_governanceItem := range pending_governance {
		pending_governanceRule = append(pending_governanceRule, pending_governanceItem)
	}

	logs, sub, err := _YFactory301.contract.FilterLogs(opts, "NewPendingGovernance", pending_governanceRule)
	if err != nil {
		return nil, err
	}
	return &YFactory301NewPendingGovernanceIterator{contract: _YFactory301.contract, event: "NewPendingGovernance", logs: logs, sub: sub}, nil
}

// WatchNewPendingGovernance is a free log subscription operation binding the contract event 0x90ad4c550d25bd23af61db38d1ff8671b89edaaa0bca0fc36bac5084ecc120bd.
//
// Solidity: event NewPendingGovernance(address indexed pending_governance)
func (_YFactory301 *YFactory301Filterer) WatchNewPendingGovernance(opts *bind.WatchOpts, sink chan<- *YFactory301NewPendingGovernance, pending_governance []common.Address) (event.Subscription, error) {

	var pending_governanceRule []interface{}
	for _, pending_governanceItem := range pending_governance {
		pending_governanceRule = append(pending_governanceRule, pending_governanceItem)
	}

	logs, sub, err := _YFactory301.contract.WatchLogs(opts, "NewPendingGovernance", pending_governanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YFactory301NewPendingGovernance)
				if err := _YFactory301.contract.UnpackLog(event, "NewPendingGovernance", log); err != nil {
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
func (_YFactory301 *YFactory301Filterer) ParseNewPendingGovernance(log types.Log) (*YFactory301NewPendingGovernance, error) {
	event := new(YFactory301NewPendingGovernance)
	if err := _YFactory301.contract.UnpackLog(event, "NewPendingGovernance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YFactory301NewVaultIterator is returned from FilterNewVault and is used to iterate over the raw logs and unpacked data for NewVault events raised by the YFactory301 contract.
type YFactory301NewVaultIterator struct {
	Event *YFactory301NewVault // Event containing the contract specifics and raw log

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
func (it *YFactory301NewVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YFactory301NewVault)
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
		it.Event = new(YFactory301NewVault)
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
func (it *YFactory301NewVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YFactory301NewVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YFactory301NewVault represents a NewVault event raised by the YFactory301 contract.
type YFactory301NewVault struct {
	VaultAddress common.Address
	Asset        common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewVault is a free log retrieval operation binding the contract event 0x4241302c393c713e690702c4a45a57e93cef59aa8c6e2358495853b3420551d8.
//
// Solidity: event NewVault(address indexed vault_address, address indexed asset)
func (_YFactory301 *YFactory301Filterer) FilterNewVault(opts *bind.FilterOpts, vault_address []common.Address, asset []common.Address) (*YFactory301NewVaultIterator, error) {

	var vault_addressRule []interface{}
	for _, vault_addressItem := range vault_address {
		vault_addressRule = append(vault_addressRule, vault_addressItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _YFactory301.contract.FilterLogs(opts, "NewVault", vault_addressRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &YFactory301NewVaultIterator{contract: _YFactory301.contract, event: "NewVault", logs: logs, sub: sub}, nil
}

// WatchNewVault is a free log subscription operation binding the contract event 0x4241302c393c713e690702c4a45a57e93cef59aa8c6e2358495853b3420551d8.
//
// Solidity: event NewVault(address indexed vault_address, address indexed asset)
func (_YFactory301 *YFactory301Filterer) WatchNewVault(opts *bind.WatchOpts, sink chan<- *YFactory301NewVault, vault_address []common.Address, asset []common.Address) (event.Subscription, error) {

	var vault_addressRule []interface{}
	for _, vault_addressItem := range vault_address {
		vault_addressRule = append(vault_addressRule, vault_addressItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _YFactory301.contract.WatchLogs(opts, "NewVault", vault_addressRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YFactory301NewVault)
				if err := _YFactory301.contract.UnpackLog(event, "NewVault", log); err != nil {
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
func (_YFactory301 *YFactory301Filterer) ParseNewVault(log types.Log) (*YFactory301NewVault, error) {
	event := new(YFactory301NewVault)
	if err := _YFactory301.contract.UnpackLog(event, "NewVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YFactory301RemovedCustomProtocolFeeIterator is returned from FilterRemovedCustomProtocolFee and is used to iterate over the raw logs and unpacked data for RemovedCustomProtocolFee events raised by the YFactory301 contract.
type YFactory301RemovedCustomProtocolFeeIterator struct {
	Event *YFactory301RemovedCustomProtocolFee // Event containing the contract specifics and raw log

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
func (it *YFactory301RemovedCustomProtocolFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YFactory301RemovedCustomProtocolFee)
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
		it.Event = new(YFactory301RemovedCustomProtocolFee)
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
func (it *YFactory301RemovedCustomProtocolFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YFactory301RemovedCustomProtocolFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YFactory301RemovedCustomProtocolFee represents a RemovedCustomProtocolFee event raised by the YFactory301 contract.
type YFactory301RemovedCustomProtocolFee struct {
	Vault common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRemovedCustomProtocolFee is a free log retrieval operation binding the contract event 0x39612c4f13d7a058dece05cf6730e3322fd9a11d6f055a5eacdde49191f45f1f.
//
// Solidity: event RemovedCustomProtocolFee(address indexed vault)
func (_YFactory301 *YFactory301Filterer) FilterRemovedCustomProtocolFee(opts *bind.FilterOpts, vault []common.Address) (*YFactory301RemovedCustomProtocolFeeIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _YFactory301.contract.FilterLogs(opts, "RemovedCustomProtocolFee", vaultRule)
	if err != nil {
		return nil, err
	}
	return &YFactory301RemovedCustomProtocolFeeIterator{contract: _YFactory301.contract, event: "RemovedCustomProtocolFee", logs: logs, sub: sub}, nil
}

// WatchRemovedCustomProtocolFee is a free log subscription operation binding the contract event 0x39612c4f13d7a058dece05cf6730e3322fd9a11d6f055a5eacdde49191f45f1f.
//
// Solidity: event RemovedCustomProtocolFee(address indexed vault)
func (_YFactory301 *YFactory301Filterer) WatchRemovedCustomProtocolFee(opts *bind.WatchOpts, sink chan<- *YFactory301RemovedCustomProtocolFee, vault []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _YFactory301.contract.WatchLogs(opts, "RemovedCustomProtocolFee", vaultRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YFactory301RemovedCustomProtocolFee)
				if err := _YFactory301.contract.UnpackLog(event, "RemovedCustomProtocolFee", log); err != nil {
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
func (_YFactory301 *YFactory301Filterer) ParseRemovedCustomProtocolFee(log types.Log) (*YFactory301RemovedCustomProtocolFee, error) {
	event := new(YFactory301RemovedCustomProtocolFee)
	if err := _YFactory301.contract.UnpackLog(event, "RemovedCustomProtocolFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YFactory301UpdateCustomProtocolFeeIterator is returned from FilterUpdateCustomProtocolFee and is used to iterate over the raw logs and unpacked data for UpdateCustomProtocolFee events raised by the YFactory301 contract.
type YFactory301UpdateCustomProtocolFeeIterator struct {
	Event *YFactory301UpdateCustomProtocolFee // Event containing the contract specifics and raw log

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
func (it *YFactory301UpdateCustomProtocolFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YFactory301UpdateCustomProtocolFee)
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
		it.Event = new(YFactory301UpdateCustomProtocolFee)
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
func (it *YFactory301UpdateCustomProtocolFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YFactory301UpdateCustomProtocolFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YFactory301UpdateCustomProtocolFee represents a UpdateCustomProtocolFee event raised by the YFactory301 contract.
type YFactory301UpdateCustomProtocolFee struct {
	Vault                common.Address
	NewCustomProtocolFee uint16
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdateCustomProtocolFee is a free log retrieval operation binding the contract event 0x96d6cc624354ffe5a7207dc2dcc152e58e23ac8df9c96943f3cfb10ea4c140ac.
//
// Solidity: event UpdateCustomProtocolFee(address indexed vault, uint16 new_custom_protocol_fee)
func (_YFactory301 *YFactory301Filterer) FilterUpdateCustomProtocolFee(opts *bind.FilterOpts, vault []common.Address) (*YFactory301UpdateCustomProtocolFeeIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _YFactory301.contract.FilterLogs(opts, "UpdateCustomProtocolFee", vaultRule)
	if err != nil {
		return nil, err
	}
	return &YFactory301UpdateCustomProtocolFeeIterator{contract: _YFactory301.contract, event: "UpdateCustomProtocolFee", logs: logs, sub: sub}, nil
}

// WatchUpdateCustomProtocolFee is a free log subscription operation binding the contract event 0x96d6cc624354ffe5a7207dc2dcc152e58e23ac8df9c96943f3cfb10ea4c140ac.
//
// Solidity: event UpdateCustomProtocolFee(address indexed vault, uint16 new_custom_protocol_fee)
func (_YFactory301 *YFactory301Filterer) WatchUpdateCustomProtocolFee(opts *bind.WatchOpts, sink chan<- *YFactory301UpdateCustomProtocolFee, vault []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _YFactory301.contract.WatchLogs(opts, "UpdateCustomProtocolFee", vaultRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YFactory301UpdateCustomProtocolFee)
				if err := _YFactory301.contract.UnpackLog(event, "UpdateCustomProtocolFee", log); err != nil {
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
func (_YFactory301 *YFactory301Filterer) ParseUpdateCustomProtocolFee(log types.Log) (*YFactory301UpdateCustomProtocolFee, error) {
	event := new(YFactory301UpdateCustomProtocolFee)
	if err := _YFactory301.contract.UnpackLog(event, "UpdateCustomProtocolFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YFactory301UpdateGovernanceIterator is returned from FilterUpdateGovernance and is used to iterate over the raw logs and unpacked data for UpdateGovernance events raised by the YFactory301 contract.
type YFactory301UpdateGovernanceIterator struct {
	Event *YFactory301UpdateGovernance // Event containing the contract specifics and raw log

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
func (it *YFactory301UpdateGovernanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YFactory301UpdateGovernance)
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
		it.Event = new(YFactory301UpdateGovernance)
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
func (it *YFactory301UpdateGovernanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YFactory301UpdateGovernanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YFactory301UpdateGovernance represents a UpdateGovernance event raised by the YFactory301 contract.
type YFactory301UpdateGovernance struct {
	Governance common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateGovernance is a free log retrieval operation binding the contract event 0x8d55d160c0009eb3d739442df0a3ca089ed64378bfac017e7ddad463f9815b87.
//
// Solidity: event UpdateGovernance(address indexed governance)
func (_YFactory301 *YFactory301Filterer) FilterUpdateGovernance(opts *bind.FilterOpts, governance []common.Address) (*YFactory301UpdateGovernanceIterator, error) {

	var governanceRule []interface{}
	for _, governanceItem := range governance {
		governanceRule = append(governanceRule, governanceItem)
	}

	logs, sub, err := _YFactory301.contract.FilterLogs(opts, "UpdateGovernance", governanceRule)
	if err != nil {
		return nil, err
	}
	return &YFactory301UpdateGovernanceIterator{contract: _YFactory301.contract, event: "UpdateGovernance", logs: logs, sub: sub}, nil
}

// WatchUpdateGovernance is a free log subscription operation binding the contract event 0x8d55d160c0009eb3d739442df0a3ca089ed64378bfac017e7ddad463f9815b87.
//
// Solidity: event UpdateGovernance(address indexed governance)
func (_YFactory301 *YFactory301Filterer) WatchUpdateGovernance(opts *bind.WatchOpts, sink chan<- *YFactory301UpdateGovernance, governance []common.Address) (event.Subscription, error) {

	var governanceRule []interface{}
	for _, governanceItem := range governance {
		governanceRule = append(governanceRule, governanceItem)
	}

	logs, sub, err := _YFactory301.contract.WatchLogs(opts, "UpdateGovernance", governanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YFactory301UpdateGovernance)
				if err := _YFactory301.contract.UnpackLog(event, "UpdateGovernance", log); err != nil {
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
func (_YFactory301 *YFactory301Filterer) ParseUpdateGovernance(log types.Log) (*YFactory301UpdateGovernance, error) {
	event := new(YFactory301UpdateGovernance)
	if err := _YFactory301.contract.UnpackLog(event, "UpdateGovernance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YFactory301UpdateProtocolFeeBpsIterator is returned from FilterUpdateProtocolFeeBps and is used to iterate over the raw logs and unpacked data for UpdateProtocolFeeBps events raised by the YFactory301 contract.
type YFactory301UpdateProtocolFeeBpsIterator struct {
	Event *YFactory301UpdateProtocolFeeBps // Event containing the contract specifics and raw log

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
func (it *YFactory301UpdateProtocolFeeBpsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YFactory301UpdateProtocolFeeBps)
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
		it.Event = new(YFactory301UpdateProtocolFeeBps)
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
func (it *YFactory301UpdateProtocolFeeBpsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YFactory301UpdateProtocolFeeBpsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YFactory301UpdateProtocolFeeBps represents a UpdateProtocolFeeBps event raised by the YFactory301 contract.
type YFactory301UpdateProtocolFeeBps struct {
	OldFeeBps uint16
	NewFeeBps uint16
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateProtocolFeeBps is a free log retrieval operation binding the contract event 0x678d2b2fe79c193f6c2c18d7515e339afcbd73fcfb360b1d0fbadae07342e051.
//
// Solidity: event UpdateProtocolFeeBps(uint16 old_fee_bps, uint16 new_fee_bps)
func (_YFactory301 *YFactory301Filterer) FilterUpdateProtocolFeeBps(opts *bind.FilterOpts) (*YFactory301UpdateProtocolFeeBpsIterator, error) {

	logs, sub, err := _YFactory301.contract.FilterLogs(opts, "UpdateProtocolFeeBps")
	if err != nil {
		return nil, err
	}
	return &YFactory301UpdateProtocolFeeBpsIterator{contract: _YFactory301.contract, event: "UpdateProtocolFeeBps", logs: logs, sub: sub}, nil
}

// WatchUpdateProtocolFeeBps is a free log subscription operation binding the contract event 0x678d2b2fe79c193f6c2c18d7515e339afcbd73fcfb360b1d0fbadae07342e051.
//
// Solidity: event UpdateProtocolFeeBps(uint16 old_fee_bps, uint16 new_fee_bps)
func (_YFactory301 *YFactory301Filterer) WatchUpdateProtocolFeeBps(opts *bind.WatchOpts, sink chan<- *YFactory301UpdateProtocolFeeBps) (event.Subscription, error) {

	logs, sub, err := _YFactory301.contract.WatchLogs(opts, "UpdateProtocolFeeBps")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YFactory301UpdateProtocolFeeBps)
				if err := _YFactory301.contract.UnpackLog(event, "UpdateProtocolFeeBps", log); err != nil {
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
func (_YFactory301 *YFactory301Filterer) ParseUpdateProtocolFeeBps(log types.Log) (*YFactory301UpdateProtocolFeeBps, error) {
	event := new(YFactory301UpdateProtocolFeeBps)
	if err := _YFactory301.contract.UnpackLog(event, "UpdateProtocolFeeBps", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YFactory301UpdateProtocolFeeRecipientIterator is returned from FilterUpdateProtocolFeeRecipient and is used to iterate over the raw logs and unpacked data for UpdateProtocolFeeRecipient events raised by the YFactory301 contract.
type YFactory301UpdateProtocolFeeRecipientIterator struct {
	Event *YFactory301UpdateProtocolFeeRecipient // Event containing the contract specifics and raw log

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
func (it *YFactory301UpdateProtocolFeeRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YFactory301UpdateProtocolFeeRecipient)
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
		it.Event = new(YFactory301UpdateProtocolFeeRecipient)
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
func (it *YFactory301UpdateProtocolFeeRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YFactory301UpdateProtocolFeeRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YFactory301UpdateProtocolFeeRecipient represents a UpdateProtocolFeeRecipient event raised by the YFactory301 contract.
type YFactory301UpdateProtocolFeeRecipient struct {
	OldFeeRecipient common.Address
	NewFeeRecipient common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateProtocolFeeRecipient is a free log retrieval operation binding the contract event 0x6af4e38beb02e4b110090dd85c5adfb341e2278b905068773762fe4666e5db7a.
//
// Solidity: event UpdateProtocolFeeRecipient(address indexed old_fee_recipient, address indexed new_fee_recipient)
func (_YFactory301 *YFactory301Filterer) FilterUpdateProtocolFeeRecipient(opts *bind.FilterOpts, old_fee_recipient []common.Address, new_fee_recipient []common.Address) (*YFactory301UpdateProtocolFeeRecipientIterator, error) {

	var old_fee_recipientRule []interface{}
	for _, old_fee_recipientItem := range old_fee_recipient {
		old_fee_recipientRule = append(old_fee_recipientRule, old_fee_recipientItem)
	}
	var new_fee_recipientRule []interface{}
	for _, new_fee_recipientItem := range new_fee_recipient {
		new_fee_recipientRule = append(new_fee_recipientRule, new_fee_recipientItem)
	}

	logs, sub, err := _YFactory301.contract.FilterLogs(opts, "UpdateProtocolFeeRecipient", old_fee_recipientRule, new_fee_recipientRule)
	if err != nil {
		return nil, err
	}
	return &YFactory301UpdateProtocolFeeRecipientIterator{contract: _YFactory301.contract, event: "UpdateProtocolFeeRecipient", logs: logs, sub: sub}, nil
}

// WatchUpdateProtocolFeeRecipient is a free log subscription operation binding the contract event 0x6af4e38beb02e4b110090dd85c5adfb341e2278b905068773762fe4666e5db7a.
//
// Solidity: event UpdateProtocolFeeRecipient(address indexed old_fee_recipient, address indexed new_fee_recipient)
func (_YFactory301 *YFactory301Filterer) WatchUpdateProtocolFeeRecipient(opts *bind.WatchOpts, sink chan<- *YFactory301UpdateProtocolFeeRecipient, old_fee_recipient []common.Address, new_fee_recipient []common.Address) (event.Subscription, error) {

	var old_fee_recipientRule []interface{}
	for _, old_fee_recipientItem := range old_fee_recipient {
		old_fee_recipientRule = append(old_fee_recipientRule, old_fee_recipientItem)
	}
	var new_fee_recipientRule []interface{}
	for _, new_fee_recipientItem := range new_fee_recipient {
		new_fee_recipientRule = append(new_fee_recipientRule, new_fee_recipientItem)
	}

	logs, sub, err := _YFactory301.contract.WatchLogs(opts, "UpdateProtocolFeeRecipient", old_fee_recipientRule, new_fee_recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YFactory301UpdateProtocolFeeRecipient)
				if err := _YFactory301.contract.UnpackLog(event, "UpdateProtocolFeeRecipient", log); err != nil {
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
func (_YFactory301 *YFactory301Filterer) ParseUpdateProtocolFeeRecipient(log types.Log) (*YFactory301UpdateProtocolFeeRecipient, error) {
	event := new(YFactory301UpdateProtocolFeeRecipient)
	if err := _YFactory301.contract.UnpackLog(event, "UpdateProtocolFeeRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
