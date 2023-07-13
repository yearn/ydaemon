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

// DeploylessMulticall3Call is an auto generated low-level Go binding around an user-defined struct.
type DeploylessMulticall3Call struct {
	Target       common.Address
	AllowFailure bool
	CallData     []byte
}

// DeploylessMulticall3MetaData contains all meta data concerning the DeploylessMulticall3 contract.
var DeploylessMulticall3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowFailure\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structDeploylessMulticall3.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]",
}

// DeploylessMulticall3ABI is the input ABI used to generate the binding from.
// Deprecated: Use DeploylessMulticall3MetaData.ABI instead.
var DeploylessMulticall3ABI = DeploylessMulticall3MetaData.ABI

// DeploylessMulticall3 is an auto generated Go binding around an Ethereum contract.
type DeploylessMulticall3 struct {
	DeploylessMulticall3Caller     // Read-only binding to the contract
	DeploylessMulticall3Transactor // Write-only binding to the contract
	DeploylessMulticall3Filterer   // Log filterer for contract events
}

// DeploylessMulticall3Caller is an auto generated read-only Go binding around an Ethereum contract.
type DeploylessMulticall3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeploylessMulticall3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type DeploylessMulticall3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeploylessMulticall3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DeploylessMulticall3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeploylessMulticall3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DeploylessMulticall3Session struct {
	Contract     *DeploylessMulticall3 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DeploylessMulticall3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DeploylessMulticall3CallerSession struct {
	Contract *DeploylessMulticall3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// DeploylessMulticall3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DeploylessMulticall3TransactorSession struct {
	Contract     *DeploylessMulticall3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// DeploylessMulticall3Raw is an auto generated low-level Go binding around an Ethereum contract.
type DeploylessMulticall3Raw struct {
	Contract *DeploylessMulticall3 // Generic contract binding to access the raw methods on
}

// DeploylessMulticall3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DeploylessMulticall3CallerRaw struct {
	Contract *DeploylessMulticall3Caller // Generic read-only contract binding to access the raw methods on
}

// DeploylessMulticall3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DeploylessMulticall3TransactorRaw struct {
	Contract *DeploylessMulticall3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewDeploylessMulticall3 creates a new instance of DeploylessMulticall3, bound to a specific deployed contract.
func NewDeploylessMulticall3(address common.Address, backend bind.ContractBackend) (*DeploylessMulticall3, error) {
	contract, err := bindDeploylessMulticall3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DeploylessMulticall3{DeploylessMulticall3Caller: DeploylessMulticall3Caller{contract: contract}, DeploylessMulticall3Transactor: DeploylessMulticall3Transactor{contract: contract}, DeploylessMulticall3Filterer: DeploylessMulticall3Filterer{contract: contract}}, nil
}

// NewDeploylessMulticall3Caller creates a new read-only instance of DeploylessMulticall3, bound to a specific deployed contract.
func NewDeploylessMulticall3Caller(address common.Address, caller bind.ContractCaller) (*DeploylessMulticall3Caller, error) {
	contract, err := bindDeploylessMulticall3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DeploylessMulticall3Caller{contract: contract}, nil
}

// NewDeploylessMulticall3Transactor creates a new write-only instance of DeploylessMulticall3, bound to a specific deployed contract.
func NewDeploylessMulticall3Transactor(address common.Address, transactor bind.ContractTransactor) (*DeploylessMulticall3Transactor, error) {
	contract, err := bindDeploylessMulticall3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DeploylessMulticall3Transactor{contract: contract}, nil
}

// NewDeploylessMulticall3Filterer creates a new log filterer instance of DeploylessMulticall3, bound to a specific deployed contract.
func NewDeploylessMulticall3Filterer(address common.Address, filterer bind.ContractFilterer) (*DeploylessMulticall3Filterer, error) {
	contract, err := bindDeploylessMulticall3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DeploylessMulticall3Filterer{contract: contract}, nil
}

// bindDeploylessMulticall3 binds a generic wrapper to an already deployed contract.
func bindDeploylessMulticall3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DeploylessMulticall3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DeploylessMulticall3 *DeploylessMulticall3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DeploylessMulticall3.Contract.DeploylessMulticall3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DeploylessMulticall3 *DeploylessMulticall3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeploylessMulticall3.Contract.DeploylessMulticall3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DeploylessMulticall3 *DeploylessMulticall3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DeploylessMulticall3.Contract.DeploylessMulticall3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DeploylessMulticall3 *DeploylessMulticall3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DeploylessMulticall3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DeploylessMulticall3 *DeploylessMulticall3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeploylessMulticall3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DeploylessMulticall3 *DeploylessMulticall3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DeploylessMulticall3.Contract.contract.Transact(opts, method, params...)
}
