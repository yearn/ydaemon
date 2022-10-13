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

// Yregistryv1MetaData contains all meta data concerning the Yregistryv1 contract.
var Yregistryv1MetaData = &bind.MetaData{
	ABI: "[{\"name\":\"NewRelease\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"release_id\",\"indexed\":true},{\"type\":\"address\",\"name\":\"template\",\"indexed\":false},{\"type\":\"string\",\"name\":\"api_version\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewVault\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"deployment_id\",\"indexed\":true},{\"type\":\"address\",\"name\":\"vault\",\"indexed\":false},{\"type\":\"string\",\"name\":\"api_version\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewExperimentalVault\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"indexed\":true},{\"type\":\"address\",\"name\":\"deployer\",\"indexed\":true},{\"type\":\"address\",\"name\":\"vault\",\"indexed\":false},{\"type\":\"string\",\"name\":\"api_version\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewGovernance\",\"inputs\":[{\"type\":\"address\",\"name\":\"governance\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"VaultTagged\",\"inputs\":[{\"type\":\"address\",\"name\":\"vault\",\"indexed\":false},{\"type\":\"string\",\"name\":\"tag\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"name\":\"setGovernance\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"governance\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":36218},{\"name\":\"acceptGovernance\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37490},{\"name\":\"latestRelease\",\"outputs\":[{\"type\":\"string\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":6804},{\"name\":\"latestVault\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"token\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2560},{\"name\":\"newRelease\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"vault\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":166808},{\"name\":\"newVault\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"token\"},{\"type\":\"address\",\"name\":\"guardian\"},{\"type\":\"address\",\"name\":\"rewards\"},{\"type\":\"string\",\"name\":\"name\"},{\"type\":\"string\",\"name\":\"symbol\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":136989},{\"name\":\"newExperimentalVault\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"token\"},{\"type\":\"address\",\"name\":\"governance\"},{\"type\":\"address\",\"name\":\"guardian\"},{\"type\":\"address\",\"name\":\"rewards\"},{\"type\":\"string\",\"name\":\"name\"},{\"type\":\"string\",\"name\":\"symbol\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":61714},{\"name\":\"endorseVault\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"vault\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":88812},{\"name\":\"setBanksy\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"tagger\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"setBanksy\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"tagger\"},{\"type\":\"bool\",\"name\":\"allowed\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"tagVault\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"vault\"},{\"type\":\"string\",\"name\":\"tag\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":186127},{\"name\":\"nextRelease\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1451},{\"name\":\"releases\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1596},{\"name\":\"nextDeployment\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1726},{\"name\":\"vaults\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"},{\"type\":\"uint256\",\"name\":\"arg1\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1871},{\"name\":\"governance\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1571},{\"name\":\"tags\",\"outputs\":[{\"type\":\"string\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":10172},{\"name\":\"banksy\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1846}]",
}

// Yregistryv1ABI is the input ABI used to generate the binding from.
// Deprecated: Use Yregistryv1MetaData.ABI instead.
var Yregistryv1ABI = Yregistryv1MetaData.ABI

// Yregistryv1 is an auto generated Go binding around an Ethereum contract.
type Yregistryv1 struct {
	Yregistryv1Caller     // Read-only binding to the contract
	Yregistryv1Transactor // Write-only binding to the contract
	Yregistryv1Filterer   // Log filterer for contract events
}

// Yregistryv1Caller is an auto generated read-only Go binding around an Ethereum contract.
type Yregistryv1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Yregistryv1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Yregistryv1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Yregistryv1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Yregistryv1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Yregistryv1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Yregistryv1Session struct {
	Contract     *Yregistryv1      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Yregistryv1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Yregistryv1CallerSession struct {
	Contract *Yregistryv1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// Yregistryv1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Yregistryv1TransactorSession struct {
	Contract     *Yregistryv1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// Yregistryv1Raw is an auto generated low-level Go binding around an Ethereum contract.
type Yregistryv1Raw struct {
	Contract *Yregistryv1 // Generic contract binding to access the raw methods on
}

// Yregistryv1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Yregistryv1CallerRaw struct {
	Contract *Yregistryv1Caller // Generic read-only contract binding to access the raw methods on
}

// Yregistryv1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Yregistryv1TransactorRaw struct {
	Contract *Yregistryv1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewYregistryv1 creates a new instance of Yregistryv1, bound to a specific deployed contract.
func NewYregistryv1(address common.Address, backend bind.ContractBackend) (*Yregistryv1, error) {
	contract, err := bindYregistryv1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Yregistryv1{Yregistryv1Caller: Yregistryv1Caller{contract: contract}, Yregistryv1Transactor: Yregistryv1Transactor{contract: contract}, Yregistryv1Filterer: Yregistryv1Filterer{contract: contract}}, nil
}

// NewYregistryv1Caller creates a new read-only instance of Yregistryv1, bound to a specific deployed contract.
func NewYregistryv1Caller(address common.Address, caller bind.ContractCaller) (*Yregistryv1Caller, error) {
	contract, err := bindYregistryv1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Yregistryv1Caller{contract: contract}, nil
}

// NewYregistryv1Transactor creates a new write-only instance of Yregistryv1, bound to a specific deployed contract.
func NewYregistryv1Transactor(address common.Address, transactor bind.ContractTransactor) (*Yregistryv1Transactor, error) {
	contract, err := bindYregistryv1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Yregistryv1Transactor{contract: contract}, nil
}

// NewYregistryv1Filterer creates a new log filterer instance of Yregistryv1, bound to a specific deployed contract.
func NewYregistryv1Filterer(address common.Address, filterer bind.ContractFilterer) (*Yregistryv1Filterer, error) {
	contract, err := bindYregistryv1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Yregistryv1Filterer{contract: contract}, nil
}

// bindYregistryv1 binds a generic wrapper to an already deployed contract.
func bindYregistryv1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Yregistryv1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Yregistryv1 *Yregistryv1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Yregistryv1.Contract.Yregistryv1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Yregistryv1 *Yregistryv1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yregistryv1.Contract.Yregistryv1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Yregistryv1 *Yregistryv1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Yregistryv1.Contract.Yregistryv1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Yregistryv1 *Yregistryv1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Yregistryv1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Yregistryv1 *Yregistryv1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yregistryv1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Yregistryv1 *Yregistryv1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Yregistryv1.Contract.contract.Transact(opts, method, params...)
}

// Banksy is a free data retrieval call binding the contract method 0xee711ed5.
//
// Solidity: function banksy(address arg0) view returns(bool)
func (_Yregistryv1 *Yregistryv1Caller) Banksy(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Yregistryv1.contract.Call(opts, &out, "banksy", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Banksy is a free data retrieval call binding the contract method 0xee711ed5.
//
// Solidity: function banksy(address arg0) view returns(bool)
func (_Yregistryv1 *Yregistryv1Session) Banksy(arg0 common.Address) (bool, error) {
	return _Yregistryv1.Contract.Banksy(&_Yregistryv1.CallOpts, arg0)
}

// Banksy is a free data retrieval call binding the contract method 0xee711ed5.
//
// Solidity: function banksy(address arg0) view returns(bool)
func (_Yregistryv1 *Yregistryv1CallerSession) Banksy(arg0 common.Address) (bool, error) {
	return _Yregistryv1.Contract.Banksy(&_Yregistryv1.CallOpts, arg0)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Yregistryv1 *Yregistryv1Caller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yregistryv1.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Yregistryv1 *Yregistryv1Session) Governance() (common.Address, error) {
	return _Yregistryv1.Contract.Governance(&_Yregistryv1.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Yregistryv1 *Yregistryv1CallerSession) Governance() (common.Address, error) {
	return _Yregistryv1.Contract.Governance(&_Yregistryv1.CallOpts)
}

// LatestRelease is a free data retrieval call binding the contract method 0x7be0ca5e.
//
// Solidity: function latestRelease() view returns(string)
func (_Yregistryv1 *Yregistryv1Caller) LatestRelease(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Yregistryv1.contract.Call(opts, &out, "latestRelease")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LatestRelease is a free data retrieval call binding the contract method 0x7be0ca5e.
//
// Solidity: function latestRelease() view returns(string)
func (_Yregistryv1 *Yregistryv1Session) LatestRelease() (string, error) {
	return _Yregistryv1.Contract.LatestRelease(&_Yregistryv1.CallOpts)
}

// LatestRelease is a free data retrieval call binding the contract method 0x7be0ca5e.
//
// Solidity: function latestRelease() view returns(string)
func (_Yregistryv1 *Yregistryv1CallerSession) LatestRelease() (string, error) {
	return _Yregistryv1.Contract.LatestRelease(&_Yregistryv1.CallOpts)
}

// LatestVault is a free data retrieval call binding the contract method 0xe177dc70.
//
// Solidity: function latestVault(address token) view returns(address)
func (_Yregistryv1 *Yregistryv1Caller) LatestVault(opts *bind.CallOpts, token common.Address) (common.Address, error) {
	var out []interface{}
	err := _Yregistryv1.contract.Call(opts, &out, "latestVault", token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LatestVault is a free data retrieval call binding the contract method 0xe177dc70.
//
// Solidity: function latestVault(address token) view returns(address)
func (_Yregistryv1 *Yregistryv1Session) LatestVault(token common.Address) (common.Address, error) {
	return _Yregistryv1.Contract.LatestVault(&_Yregistryv1.CallOpts, token)
}

// LatestVault is a free data retrieval call binding the contract method 0xe177dc70.
//
// Solidity: function latestVault(address token) view returns(address)
func (_Yregistryv1 *Yregistryv1CallerSession) LatestVault(token common.Address) (common.Address, error) {
	return _Yregistryv1.Contract.LatestVault(&_Yregistryv1.CallOpts, token)
}

// NextDeployment is a free data retrieval call binding the contract method 0xba7ee48c.
//
// Solidity: function nextDeployment(address arg0) view returns(uint256)
func (_Yregistryv1 *Yregistryv1Caller) NextDeployment(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yregistryv1.contract.Call(opts, &out, "nextDeployment", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextDeployment is a free data retrieval call binding the contract method 0xba7ee48c.
//
// Solidity: function nextDeployment(address arg0) view returns(uint256)
func (_Yregistryv1 *Yregistryv1Session) NextDeployment(arg0 common.Address) (*big.Int, error) {
	return _Yregistryv1.Contract.NextDeployment(&_Yregistryv1.CallOpts, arg0)
}

// NextDeployment is a free data retrieval call binding the contract method 0xba7ee48c.
//
// Solidity: function nextDeployment(address arg0) view returns(uint256)
func (_Yregistryv1 *Yregistryv1CallerSession) NextDeployment(arg0 common.Address) (*big.Int, error) {
	return _Yregistryv1.Contract.NextDeployment(&_Yregistryv1.CallOpts, arg0)
}

// NextRelease is a free data retrieval call binding the contract method 0xa1571902.
//
// Solidity: function nextRelease() view returns(uint256)
func (_Yregistryv1 *Yregistryv1Caller) NextRelease(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yregistryv1.contract.Call(opts, &out, "nextRelease")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextRelease is a free data retrieval call binding the contract method 0xa1571902.
//
// Solidity: function nextRelease() view returns(uint256)
func (_Yregistryv1 *Yregistryv1Session) NextRelease() (*big.Int, error) {
	return _Yregistryv1.Contract.NextRelease(&_Yregistryv1.CallOpts)
}

// NextRelease is a free data retrieval call binding the contract method 0xa1571902.
//
// Solidity: function nextRelease() view returns(uint256)
func (_Yregistryv1 *Yregistryv1CallerSession) NextRelease() (*big.Int, error) {
	return _Yregistryv1.Contract.NextRelease(&_Yregistryv1.CallOpts)
}

// Releases is a free data retrieval call binding the contract method 0xb6a9f40f.
//
// Solidity: function releases(uint256 arg0) view returns(address)
func (_Yregistryv1 *Yregistryv1Caller) Releases(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Yregistryv1.contract.Call(opts, &out, "releases", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Releases is a free data retrieval call binding the contract method 0xb6a9f40f.
//
// Solidity: function releases(uint256 arg0) view returns(address)
func (_Yregistryv1 *Yregistryv1Session) Releases(arg0 *big.Int) (common.Address, error) {
	return _Yregistryv1.Contract.Releases(&_Yregistryv1.CallOpts, arg0)
}

// Releases is a free data retrieval call binding the contract method 0xb6a9f40f.
//
// Solidity: function releases(uint256 arg0) view returns(address)
func (_Yregistryv1 *Yregistryv1CallerSession) Releases(arg0 *big.Int) (common.Address, error) {
	return _Yregistryv1.Contract.Releases(&_Yregistryv1.CallOpts, arg0)
}

// Tags is a free data retrieval call binding the contract method 0x58b8f842.
//
// Solidity: function tags(address arg0) view returns(string)
func (_Yregistryv1 *Yregistryv1Caller) Tags(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _Yregistryv1.contract.Call(opts, &out, "tags", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Tags is a free data retrieval call binding the contract method 0x58b8f842.
//
// Solidity: function tags(address arg0) view returns(string)
func (_Yregistryv1 *Yregistryv1Session) Tags(arg0 common.Address) (string, error) {
	return _Yregistryv1.Contract.Tags(&_Yregistryv1.CallOpts, arg0)
}

// Tags is a free data retrieval call binding the contract method 0x58b8f842.
//
// Solidity: function tags(address arg0) view returns(string)
func (_Yregistryv1 *Yregistryv1CallerSession) Tags(arg0 common.Address) (string, error) {
	return _Yregistryv1.Contract.Tags(&_Yregistryv1.CallOpts, arg0)
}

// Vaults is a free data retrieval call binding the contract method 0x7bbfc69e.
//
// Solidity: function vaults(address arg0, uint256 arg1) view returns(address)
func (_Yregistryv1 *Yregistryv1Caller) Vaults(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Yregistryv1.contract.Call(opts, &out, "vaults", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vaults is a free data retrieval call binding the contract method 0x7bbfc69e.
//
// Solidity: function vaults(address arg0, uint256 arg1) view returns(address)
func (_Yregistryv1 *Yregistryv1Session) Vaults(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Yregistryv1.Contract.Vaults(&_Yregistryv1.CallOpts, arg0, arg1)
}

// Vaults is a free data retrieval call binding the contract method 0x7bbfc69e.
//
// Solidity: function vaults(address arg0, uint256 arg1) view returns(address)
func (_Yregistryv1 *Yregistryv1CallerSession) Vaults(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Yregistryv1.Contract.Vaults(&_Yregistryv1.CallOpts, arg0, arg1)
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_Yregistryv1 *Yregistryv1Transactor) AcceptGovernance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yregistryv1.contract.Transact(opts, "acceptGovernance")
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_Yregistryv1 *Yregistryv1Session) AcceptGovernance() (*types.Transaction, error) {
	return _Yregistryv1.Contract.AcceptGovernance(&_Yregistryv1.TransactOpts)
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_Yregistryv1 *Yregistryv1TransactorSession) AcceptGovernance() (*types.Transaction, error) {
	return _Yregistryv1.Contract.AcceptGovernance(&_Yregistryv1.TransactOpts)
}

// EndorseVault is a paid mutator transaction binding the contract method 0x29b2e0c6.
//
// Solidity: function endorseVault(address vault) returns()
func (_Yregistryv1 *Yregistryv1Transactor) EndorseVault(opts *bind.TransactOpts, vault common.Address) (*types.Transaction, error) {
	return _Yregistryv1.contract.Transact(opts, "endorseVault", vault)
}

// EndorseVault is a paid mutator transaction binding the contract method 0x29b2e0c6.
//
// Solidity: function endorseVault(address vault) returns()
func (_Yregistryv1 *Yregistryv1Session) EndorseVault(vault common.Address) (*types.Transaction, error) {
	return _Yregistryv1.Contract.EndorseVault(&_Yregistryv1.TransactOpts, vault)
}

// EndorseVault is a paid mutator transaction binding the contract method 0x29b2e0c6.
//
// Solidity: function endorseVault(address vault) returns()
func (_Yregistryv1 *Yregistryv1TransactorSession) EndorseVault(vault common.Address) (*types.Transaction, error) {
	return _Yregistryv1.Contract.EndorseVault(&_Yregistryv1.TransactOpts, vault)
}

// NewExperimentalVault is a paid mutator transaction binding the contract method 0x5b73aa0d.
//
// Solidity: function newExperimentalVault(address token, address governance, address guardian, address rewards, string name, string symbol) returns(address)
func (_Yregistryv1 *Yregistryv1Transactor) NewExperimentalVault(opts *bind.TransactOpts, token common.Address, governance common.Address, guardian common.Address, rewards common.Address, name string, symbol string) (*types.Transaction, error) {
	return _Yregistryv1.contract.Transact(opts, "newExperimentalVault", token, governance, guardian, rewards, name, symbol)
}

// NewExperimentalVault is a paid mutator transaction binding the contract method 0x5b73aa0d.
//
// Solidity: function newExperimentalVault(address token, address governance, address guardian, address rewards, string name, string symbol) returns(address)
func (_Yregistryv1 *Yregistryv1Session) NewExperimentalVault(token common.Address, governance common.Address, guardian common.Address, rewards common.Address, name string, symbol string) (*types.Transaction, error) {
	return _Yregistryv1.Contract.NewExperimentalVault(&_Yregistryv1.TransactOpts, token, governance, guardian, rewards, name, symbol)
}

// NewExperimentalVault is a paid mutator transaction binding the contract method 0x5b73aa0d.
//
// Solidity: function newExperimentalVault(address token, address governance, address guardian, address rewards, string name, string symbol) returns(address)
func (_Yregistryv1 *Yregistryv1TransactorSession) NewExperimentalVault(token common.Address, governance common.Address, guardian common.Address, rewards common.Address, name string, symbol string) (*types.Transaction, error) {
	return _Yregistryv1.Contract.NewExperimentalVault(&_Yregistryv1.TransactOpts, token, governance, guardian, rewards, name, symbol)
}

// NewRelease is a paid mutator transaction binding the contract method 0x33990d4b.
//
// Solidity: function newRelease(address vault) returns()
func (_Yregistryv1 *Yregistryv1Transactor) NewRelease(opts *bind.TransactOpts, vault common.Address) (*types.Transaction, error) {
	return _Yregistryv1.contract.Transact(opts, "newRelease", vault)
}

// NewRelease is a paid mutator transaction binding the contract method 0x33990d4b.
//
// Solidity: function newRelease(address vault) returns()
func (_Yregistryv1 *Yregistryv1Session) NewRelease(vault common.Address) (*types.Transaction, error) {
	return _Yregistryv1.Contract.NewRelease(&_Yregistryv1.TransactOpts, vault)
}

// NewRelease is a paid mutator transaction binding the contract method 0x33990d4b.
//
// Solidity: function newRelease(address vault) returns()
func (_Yregistryv1 *Yregistryv1TransactorSession) NewRelease(vault common.Address) (*types.Transaction, error) {
	return _Yregistryv1.Contract.NewRelease(&_Yregistryv1.TransactOpts, vault)
}

// NewVault is a paid mutator transaction binding the contract method 0x108ca11e.
//
// Solidity: function newVault(address token, address guardian, address rewards, string name, string symbol) returns(address)
func (_Yregistryv1 *Yregistryv1Transactor) NewVault(opts *bind.TransactOpts, token common.Address, guardian common.Address, rewards common.Address, name string, symbol string) (*types.Transaction, error) {
	return _Yregistryv1.contract.Transact(opts, "newVault", token, guardian, rewards, name, symbol)
}

// NewVault is a paid mutator transaction binding the contract method 0x108ca11e.
//
// Solidity: function newVault(address token, address guardian, address rewards, string name, string symbol) returns(address)
func (_Yregistryv1 *Yregistryv1Session) NewVault(token common.Address, guardian common.Address, rewards common.Address, name string, symbol string) (*types.Transaction, error) {
	return _Yregistryv1.Contract.NewVault(&_Yregistryv1.TransactOpts, token, guardian, rewards, name, symbol)
}

// NewVault is a paid mutator transaction binding the contract method 0x108ca11e.
//
// Solidity: function newVault(address token, address guardian, address rewards, string name, string symbol) returns(address)
func (_Yregistryv1 *Yregistryv1TransactorSession) NewVault(token common.Address, guardian common.Address, rewards common.Address, name string, symbol string) (*types.Transaction, error) {
	return _Yregistryv1.Contract.NewVault(&_Yregistryv1.TransactOpts, token, guardian, rewards, name, symbol)
}

// SetBanksy is a paid mutator transaction binding the contract method 0x2cad8f9f.
//
// Solidity: function setBanksy(address tagger) returns()
func (_Yregistryv1 *Yregistryv1Transactor) SetBanksy(opts *bind.TransactOpts, tagger common.Address) (*types.Transaction, error) {
	return _Yregistryv1.contract.Transact(opts, "setBanksy", tagger)
}

// SetBanksy is a paid mutator transaction binding the contract method 0x2cad8f9f.
//
// Solidity: function setBanksy(address tagger) returns()
func (_Yregistryv1 *Yregistryv1Session) SetBanksy(tagger common.Address) (*types.Transaction, error) {
	return _Yregistryv1.Contract.SetBanksy(&_Yregistryv1.TransactOpts, tagger)
}

// SetBanksy is a paid mutator transaction binding the contract method 0x2cad8f9f.
//
// Solidity: function setBanksy(address tagger) returns()
func (_Yregistryv1 *Yregistryv1TransactorSession) SetBanksy(tagger common.Address) (*types.Transaction, error) {
	return _Yregistryv1.Contract.SetBanksy(&_Yregistryv1.TransactOpts, tagger)
}

// SetBanksy0 is a paid mutator transaction binding the contract method 0x5e05f6af.
//
// Solidity: function setBanksy(address tagger, bool allowed) returns()
func (_Yregistryv1 *Yregistryv1Transactor) SetBanksy0(opts *bind.TransactOpts, tagger common.Address, allowed bool) (*types.Transaction, error) {
	return _Yregistryv1.contract.Transact(opts, "setBanksy0", tagger, allowed)
}

// SetBanksy0 is a paid mutator transaction binding the contract method 0x5e05f6af.
//
// Solidity: function setBanksy(address tagger, bool allowed) returns()
func (_Yregistryv1 *Yregistryv1Session) SetBanksy0(tagger common.Address, allowed bool) (*types.Transaction, error) {
	return _Yregistryv1.Contract.SetBanksy0(&_Yregistryv1.TransactOpts, tagger, allowed)
}

// SetBanksy0 is a paid mutator transaction binding the contract method 0x5e05f6af.
//
// Solidity: function setBanksy(address tagger, bool allowed) returns()
func (_Yregistryv1 *Yregistryv1TransactorSession) SetBanksy0(tagger common.Address, allowed bool) (*types.Transaction, error) {
	return _Yregistryv1.Contract.SetBanksy0(&_Yregistryv1.TransactOpts, tagger, allowed)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address governance) returns()
func (_Yregistryv1 *Yregistryv1Transactor) SetGovernance(opts *bind.TransactOpts, governance common.Address) (*types.Transaction, error) {
	return _Yregistryv1.contract.Transact(opts, "setGovernance", governance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address governance) returns()
func (_Yregistryv1 *Yregistryv1Session) SetGovernance(governance common.Address) (*types.Transaction, error) {
	return _Yregistryv1.Contract.SetGovernance(&_Yregistryv1.TransactOpts, governance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address governance) returns()
func (_Yregistryv1 *Yregistryv1TransactorSession) SetGovernance(governance common.Address) (*types.Transaction, error) {
	return _Yregistryv1.Contract.SetGovernance(&_Yregistryv1.TransactOpts, governance)
}

// TagVault is a paid mutator transaction binding the contract method 0x60bd68f8.
//
// Solidity: function tagVault(address vault, string tag) returns()
func (_Yregistryv1 *Yregistryv1Transactor) TagVault(opts *bind.TransactOpts, vault common.Address, tag string) (*types.Transaction, error) {
	return _Yregistryv1.contract.Transact(opts, "tagVault", vault, tag)
}

// TagVault is a paid mutator transaction binding the contract method 0x60bd68f8.
//
// Solidity: function tagVault(address vault, string tag) returns()
func (_Yregistryv1 *Yregistryv1Session) TagVault(vault common.Address, tag string) (*types.Transaction, error) {
	return _Yregistryv1.Contract.TagVault(&_Yregistryv1.TransactOpts, vault, tag)
}

// TagVault is a paid mutator transaction binding the contract method 0x60bd68f8.
//
// Solidity: function tagVault(address vault, string tag) returns()
func (_Yregistryv1 *Yregistryv1TransactorSession) TagVault(vault common.Address, tag string) (*types.Transaction, error) {
	return _Yregistryv1.Contract.TagVault(&_Yregistryv1.TransactOpts, vault, tag)
}

// Yregistryv1NewExperimentalVaultIterator is returned from FilterNewExperimentalVault and is used to iterate over the raw logs and unpacked data for NewExperimentalVault events raised by the Yregistryv1 contract.
type Yregistryv1NewExperimentalVaultIterator struct {
	Event *Yregistryv1NewExperimentalVault // Event containing the contract specifics and raw log

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
func (it *Yregistryv1NewExperimentalVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yregistryv1NewExperimentalVault)
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
		it.Event = new(Yregistryv1NewExperimentalVault)
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
func (it *Yregistryv1NewExperimentalVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yregistryv1NewExperimentalVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yregistryv1NewExperimentalVault represents a NewExperimentalVault event raised by the Yregistryv1 contract.
type Yregistryv1NewExperimentalVault struct {
	Token      common.Address
	Deployer   common.Address
	Vault      common.Address
	ApiVersion string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewExperimentalVault is a free log retrieval operation binding the contract event 0x57a9cdc2a05e05f66e76769bdbe88e21ec45d9ee0f97d4cb60395d4c75dcbcda.
//
// Solidity: event NewExperimentalVault(address indexed token, address indexed deployer, address vault, string api_version)
func (_Yregistryv1 *Yregistryv1Filterer) FilterNewExperimentalVault(opts *bind.FilterOpts, token []common.Address, deployer []common.Address) (*Yregistryv1NewExperimentalVaultIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var deployerRule []interface{}
	for _, deployerItem := range deployer {
		deployerRule = append(deployerRule, deployerItem)
	}

	logs, sub, err := _Yregistryv1.contract.FilterLogs(opts, "NewExperimentalVault", tokenRule, deployerRule)
	if err != nil {
		return nil, err
	}
	return &Yregistryv1NewExperimentalVaultIterator{contract: _Yregistryv1.contract, event: "NewExperimentalVault", logs: logs, sub: sub}, nil
}

// WatchNewExperimentalVault is a free log subscription operation binding the contract event 0x57a9cdc2a05e05f66e76769bdbe88e21ec45d9ee0f97d4cb60395d4c75dcbcda.
//
// Solidity: event NewExperimentalVault(address indexed token, address indexed deployer, address vault, string api_version)
func (_Yregistryv1 *Yregistryv1Filterer) WatchNewExperimentalVault(opts *bind.WatchOpts, sink chan<- *Yregistryv1NewExperimentalVault, token []common.Address, deployer []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var deployerRule []interface{}
	for _, deployerItem := range deployer {
		deployerRule = append(deployerRule, deployerItem)
	}

	logs, sub, err := _Yregistryv1.contract.WatchLogs(opts, "NewExperimentalVault", tokenRule, deployerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yregistryv1NewExperimentalVault)
				if err := _Yregistryv1.contract.UnpackLog(event, "NewExperimentalVault", log); err != nil {
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
func (_Yregistryv1 *Yregistryv1Filterer) ParseNewExperimentalVault(log types.Log) (*Yregistryv1NewExperimentalVault, error) {
	event := new(Yregistryv1NewExperimentalVault)
	if err := _Yregistryv1.contract.UnpackLog(event, "NewExperimentalVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yregistryv1NewGovernanceIterator is returned from FilterNewGovernance and is used to iterate over the raw logs and unpacked data for NewGovernance events raised by the Yregistryv1 contract.
type Yregistryv1NewGovernanceIterator struct {
	Event *Yregistryv1NewGovernance // Event containing the contract specifics and raw log

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
func (it *Yregistryv1NewGovernanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yregistryv1NewGovernance)
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
		it.Event = new(Yregistryv1NewGovernance)
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
func (it *Yregistryv1NewGovernanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yregistryv1NewGovernanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yregistryv1NewGovernance represents a NewGovernance event raised by the Yregistryv1 contract.
type Yregistryv1NewGovernance struct {
	Governance common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewGovernance is a free log retrieval operation binding the contract event 0x4f386975ea1c2f7cf1845b08bee00626fbb624ecad16254d63d9bb9ba86526de.
//
// Solidity: event NewGovernance(address governance)
func (_Yregistryv1 *Yregistryv1Filterer) FilterNewGovernance(opts *bind.FilterOpts) (*Yregistryv1NewGovernanceIterator, error) {

	logs, sub, err := _Yregistryv1.contract.FilterLogs(opts, "NewGovernance")
	if err != nil {
		return nil, err
	}
	return &Yregistryv1NewGovernanceIterator{contract: _Yregistryv1.contract, event: "NewGovernance", logs: logs, sub: sub}, nil
}

// WatchNewGovernance is a free log subscription operation binding the contract event 0x4f386975ea1c2f7cf1845b08bee00626fbb624ecad16254d63d9bb9ba86526de.
//
// Solidity: event NewGovernance(address governance)
func (_Yregistryv1 *Yregistryv1Filterer) WatchNewGovernance(opts *bind.WatchOpts, sink chan<- *Yregistryv1NewGovernance) (event.Subscription, error) {

	logs, sub, err := _Yregistryv1.contract.WatchLogs(opts, "NewGovernance")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yregistryv1NewGovernance)
				if err := _Yregistryv1.contract.UnpackLog(event, "NewGovernance", log); err != nil {
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
func (_Yregistryv1 *Yregistryv1Filterer) ParseNewGovernance(log types.Log) (*Yregistryv1NewGovernance, error) {
	event := new(Yregistryv1NewGovernance)
	if err := _Yregistryv1.contract.UnpackLog(event, "NewGovernance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yregistryv1NewReleaseIterator is returned from FilterNewRelease and is used to iterate over the raw logs and unpacked data for NewRelease events raised by the Yregistryv1 contract.
type Yregistryv1NewReleaseIterator struct {
	Event *Yregistryv1NewRelease // Event containing the contract specifics and raw log

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
func (it *Yregistryv1NewReleaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yregistryv1NewRelease)
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
		it.Event = new(Yregistryv1NewRelease)
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
func (it *Yregistryv1NewReleaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yregistryv1NewReleaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yregistryv1NewRelease represents a NewRelease event raised by the Yregistryv1 contract.
type Yregistryv1NewRelease struct {
	ReleaseId  *big.Int
	Template   common.Address
	ApiVersion string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewRelease is a free log retrieval operation binding the contract event 0xa6fbd216b6734f34092f1be6b7247e1551a6d4f2d5000c53721cfdc119a5b7cf.
//
// Solidity: event NewRelease(uint256 indexed release_id, address template, string api_version)
func (_Yregistryv1 *Yregistryv1Filterer) FilterNewRelease(opts *bind.FilterOpts, release_id []*big.Int) (*Yregistryv1NewReleaseIterator, error) {

	var release_idRule []interface{}
	for _, release_idItem := range release_id {
		release_idRule = append(release_idRule, release_idItem)
	}

	logs, sub, err := _Yregistryv1.contract.FilterLogs(opts, "NewRelease", release_idRule)
	if err != nil {
		return nil, err
	}
	return &Yregistryv1NewReleaseIterator{contract: _Yregistryv1.contract, event: "NewRelease", logs: logs, sub: sub}, nil
}

// WatchNewRelease is a free log subscription operation binding the contract event 0xa6fbd216b6734f34092f1be6b7247e1551a6d4f2d5000c53721cfdc119a5b7cf.
//
// Solidity: event NewRelease(uint256 indexed release_id, address template, string api_version)
func (_Yregistryv1 *Yregistryv1Filterer) WatchNewRelease(opts *bind.WatchOpts, sink chan<- *Yregistryv1NewRelease, release_id []*big.Int) (event.Subscription, error) {

	var release_idRule []interface{}
	for _, release_idItem := range release_id {
		release_idRule = append(release_idRule, release_idItem)
	}

	logs, sub, err := _Yregistryv1.contract.WatchLogs(opts, "NewRelease", release_idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yregistryv1NewRelease)
				if err := _Yregistryv1.contract.UnpackLog(event, "NewRelease", log); err != nil {
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
func (_Yregistryv1 *Yregistryv1Filterer) ParseNewRelease(log types.Log) (*Yregistryv1NewRelease, error) {
	event := new(Yregistryv1NewRelease)
	if err := _Yregistryv1.contract.UnpackLog(event, "NewRelease", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yregistryv1NewVaultIterator is returned from FilterNewVault and is used to iterate over the raw logs and unpacked data for NewVault events raised by the Yregistryv1 contract.
type Yregistryv1NewVaultIterator struct {
	Event *Yregistryv1NewVault // Event containing the contract specifics and raw log

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
func (it *Yregistryv1NewVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yregistryv1NewVault)
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
		it.Event = new(Yregistryv1NewVault)
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
func (it *Yregistryv1NewVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yregistryv1NewVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yregistryv1NewVault represents a NewVault event raised by the Yregistryv1 contract.
type Yregistryv1NewVault struct {
	Token        common.Address
	DeploymentId *big.Int
	Vault        common.Address
	ApiVersion   string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewVault is a free log retrieval operation binding the contract event 0xce089905ba4a4d622553bcb5646fd23e895c256f0376eee04e99e61cec1dc7e8.
//
// Solidity: event NewVault(address indexed token, uint256 indexed deployment_id, address vault, string api_version)
func (_Yregistryv1 *Yregistryv1Filterer) FilterNewVault(opts *bind.FilterOpts, token []common.Address, deployment_id []*big.Int) (*Yregistryv1NewVaultIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var deployment_idRule []interface{}
	for _, deployment_idItem := range deployment_id {
		deployment_idRule = append(deployment_idRule, deployment_idItem)
	}

	logs, sub, err := _Yregistryv1.contract.FilterLogs(opts, "NewVault", tokenRule, deployment_idRule)
	if err != nil {
		return nil, err
	}
	return &Yregistryv1NewVaultIterator{contract: _Yregistryv1.contract, event: "NewVault", logs: logs, sub: sub}, nil
}

// WatchNewVault is a free log subscription operation binding the contract event 0xce089905ba4a4d622553bcb5646fd23e895c256f0376eee04e99e61cec1dc7e8.
//
// Solidity: event NewVault(address indexed token, uint256 indexed deployment_id, address vault, string api_version)
func (_Yregistryv1 *Yregistryv1Filterer) WatchNewVault(opts *bind.WatchOpts, sink chan<- *Yregistryv1NewVault, token []common.Address, deployment_id []*big.Int) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var deployment_idRule []interface{}
	for _, deployment_idItem := range deployment_id {
		deployment_idRule = append(deployment_idRule, deployment_idItem)
	}

	logs, sub, err := _Yregistryv1.contract.WatchLogs(opts, "NewVault", tokenRule, deployment_idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yregistryv1NewVault)
				if err := _Yregistryv1.contract.UnpackLog(event, "NewVault", log); err != nil {
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
// Solidity: event NewVault(address indexed token, uint256 indexed deployment_id, address vault, string api_version)
func (_Yregistryv1 *Yregistryv1Filterer) ParseNewVault(log types.Log) (*Yregistryv1NewVault, error) {
	event := new(Yregistryv1NewVault)
	if err := _Yregistryv1.contract.UnpackLog(event, "NewVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yregistryv1VaultTaggedIterator is returned from FilterVaultTagged and is used to iterate over the raw logs and unpacked data for VaultTagged events raised by the Yregistryv1 contract.
type Yregistryv1VaultTaggedIterator struct {
	Event *Yregistryv1VaultTagged // Event containing the contract specifics and raw log

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
func (it *Yregistryv1VaultTaggedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yregistryv1VaultTagged)
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
		it.Event = new(Yregistryv1VaultTagged)
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
func (it *Yregistryv1VaultTaggedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yregistryv1VaultTaggedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yregistryv1VaultTagged represents a VaultTagged event raised by the Yregistryv1 contract.
type Yregistryv1VaultTagged struct {
	Vault common.Address
	Tag   string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterVaultTagged is a free log retrieval operation binding the contract event 0x07bd58794c2ca0ae152f7719eb5f02c654476de972cadec0e8191ae8be42096d.
//
// Solidity: event VaultTagged(address vault, string tag)
func (_Yregistryv1 *Yregistryv1Filterer) FilterVaultTagged(opts *bind.FilterOpts) (*Yregistryv1VaultTaggedIterator, error) {

	logs, sub, err := _Yregistryv1.contract.FilterLogs(opts, "VaultTagged")
	if err != nil {
		return nil, err
	}
	return &Yregistryv1VaultTaggedIterator{contract: _Yregistryv1.contract, event: "VaultTagged", logs: logs, sub: sub}, nil
}

// WatchVaultTagged is a free log subscription operation binding the contract event 0x07bd58794c2ca0ae152f7719eb5f02c654476de972cadec0e8191ae8be42096d.
//
// Solidity: event VaultTagged(address vault, string tag)
func (_Yregistryv1 *Yregistryv1Filterer) WatchVaultTagged(opts *bind.WatchOpts, sink chan<- *Yregistryv1VaultTagged) (event.Subscription, error) {

	logs, sub, err := _Yregistryv1.contract.WatchLogs(opts, "VaultTagged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yregistryv1VaultTagged)
				if err := _Yregistryv1.contract.UnpackLog(event, "VaultTagged", log); err != nil {
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
func (_Yregistryv1 *Yregistryv1Filterer) ParseVaultTagged(log types.Log) (*Yregistryv1VaultTagged, error) {
	event := new(Yregistryv1VaultTagged)
	if err := _Yregistryv1.contract.UnpackLog(event, "VaultTagged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
