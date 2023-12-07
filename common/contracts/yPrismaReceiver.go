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

// YPrismaReceiverMetaData contains all meta data concerning the YPrismaReceiver contract.
var YPrismaReceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_prisma\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_CRV\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_CVX\",\"type\":\"address\"},{\"internalType\":\"contractIBooster\",\"name\":\"_booster\",\"type\":\"address\"},{\"internalType\":\"contractICurveProxy\",\"name\":\"_proxy\",\"type\":\"address\"},{\"internalType\":\"contractIPrismaVault\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"prismaCore\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lpToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LPTokenDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lpToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LPTokenWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pct\",\"type\":\"uint256\"}],\"name\":\"MaxWeeklyEmissionPctSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allocated\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxAllowed\",\"type\":\"uint256\"}],\"name\":\"MaxWeeklyEmissionsExceeded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prismaAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"crvAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cvxAmount\",\"type\":\"uint256\"}],\"name\":\"RewardClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CRV\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CVX\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRISMA\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRISMA_CORE\",\"outputs\":[{\"internalType\":\"contractIPrismaCore\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"booster\",\"outputs\":[{\"internalType\":\"contractIBooster\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"claimReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"prismaAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"crvAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cvxAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"claimableReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"prismaAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"crvAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cvxAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crvRewards\",\"outputs\":[{\"internalType\":\"contractIBaseRewardPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"curveProxy\",\"outputs\":[{\"internalType\":\"contractICurveProxy\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cvxRewards\",\"outputs\":[{\"internalType\":\"contractIBaseRewardPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositPid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emissionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fetchRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"guardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastCrvBalance\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastCvxBalance\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdate\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lpToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxWeeklyEmissionPct\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"assignedIds\",\"type\":\"uint256[]\"}],\"name\":\"notifyRegisteredId\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodFinish\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pushExcessEmissions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewardIntegral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewardIntegralFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewardRate\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_maxWeeklyEmissionPct\",\"type\":\"uint16\"}],\"name\":\"setMaxWeeklyEmissionPct\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"storedExcessEmissions\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"contractIPrismaVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"claimant\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"vaultClaimReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// YPrismaReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use YPrismaReceiverMetaData.ABI instead.
var YPrismaReceiverABI = YPrismaReceiverMetaData.ABI

// YPrismaReceiver is an auto generated Go binding around an Ethereum contract.
type YPrismaReceiver struct {
	YPrismaReceiverCaller     // Read-only binding to the contract
	YPrismaReceiverTransactor // Write-only binding to the contract
	YPrismaReceiverFilterer   // Log filterer for contract events
}

// YPrismaReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type YPrismaReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YPrismaReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YPrismaReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YPrismaReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YPrismaReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YPrismaReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YPrismaReceiverSession struct {
	Contract     *YPrismaReceiver  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YPrismaReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YPrismaReceiverCallerSession struct {
	Contract *YPrismaReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// YPrismaReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YPrismaReceiverTransactorSession struct {
	Contract     *YPrismaReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// YPrismaReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type YPrismaReceiverRaw struct {
	Contract *YPrismaReceiver // Generic contract binding to access the raw methods on
}

// YPrismaReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YPrismaReceiverCallerRaw struct {
	Contract *YPrismaReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// YPrismaReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YPrismaReceiverTransactorRaw struct {
	Contract *YPrismaReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYPrismaReceiver creates a new instance of YPrismaReceiver, bound to a specific deployed contract.
func NewYPrismaReceiver(address common.Address, backend bind.ContractBackend) (*YPrismaReceiver, error) {
	contract, err := bindYPrismaReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YPrismaReceiver{YPrismaReceiverCaller: YPrismaReceiverCaller{contract: contract}, YPrismaReceiverTransactor: YPrismaReceiverTransactor{contract: contract}, YPrismaReceiverFilterer: YPrismaReceiverFilterer{contract: contract}}, nil
}

// NewYPrismaReceiverCaller creates a new read-only instance of YPrismaReceiver, bound to a specific deployed contract.
func NewYPrismaReceiverCaller(address common.Address, caller bind.ContractCaller) (*YPrismaReceiverCaller, error) {
	contract, err := bindYPrismaReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YPrismaReceiverCaller{contract: contract}, nil
}

// NewYPrismaReceiverTransactor creates a new write-only instance of YPrismaReceiver, bound to a specific deployed contract.
func NewYPrismaReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*YPrismaReceiverTransactor, error) {
	contract, err := bindYPrismaReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YPrismaReceiverTransactor{contract: contract}, nil
}

// NewYPrismaReceiverFilterer creates a new log filterer instance of YPrismaReceiver, bound to a specific deployed contract.
func NewYPrismaReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*YPrismaReceiverFilterer, error) {
	contract, err := bindYPrismaReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YPrismaReceiverFilterer{contract: contract}, nil
}

// bindYPrismaReceiver binds a generic wrapper to an already deployed contract.
func bindYPrismaReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := YPrismaReceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YPrismaReceiver *YPrismaReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YPrismaReceiver.Contract.YPrismaReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YPrismaReceiver *YPrismaReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YPrismaReceiver.Contract.YPrismaReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YPrismaReceiver *YPrismaReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YPrismaReceiver.Contract.YPrismaReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YPrismaReceiver *YPrismaReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YPrismaReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YPrismaReceiver *YPrismaReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YPrismaReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YPrismaReceiver *YPrismaReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YPrismaReceiver.Contract.contract.Transact(opts, method, params...)
}

// CRV is a free data retrieval call binding the contract method 0x945c9142.
//
// Solidity: function CRV() view returns(address)
func (_YPrismaReceiver *YPrismaReceiverCaller) CRV(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YPrismaReceiver.contract.Call(opts, &out, "CRV")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CRV is a free data retrieval call binding the contract method 0x945c9142.
//
// Solidity: function CRV() view returns(address)
func (_YPrismaReceiver *YPrismaReceiverSession) CRV() (common.Address, error) {
	return _YPrismaReceiver.Contract.CRV(&_YPrismaReceiver.CallOpts)
}

// CRV is a free data retrieval call binding the contract method 0x945c9142.
//
// Solidity: function CRV() view returns(address)
func (_YPrismaReceiver *YPrismaReceiverCallerSession) CRV() (common.Address, error) {
	return _YPrismaReceiver.Contract.CRV(&_YPrismaReceiver.CallOpts)
}

// CVX is a free data retrieval call binding the contract method 0x759cb53b.
//
// Solidity: function CVX() view returns(address)
func (_YPrismaReceiver *YPrismaReceiverCaller) CVX(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YPrismaReceiver.contract.Call(opts, &out, "CVX")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CVX is a free data retrieval call binding the contract method 0x759cb53b.
//
// Solidity: function CVX() view returns(address)
func (_YPrismaReceiver *YPrismaReceiverSession) CVX() (common.Address, error) {
	return _YPrismaReceiver.Contract.CVX(&_YPrismaReceiver.CallOpts)
}

// CVX is a free data retrieval call binding the contract method 0x759cb53b.
//
// Solidity: function CVX() view returns(address)
func (_YPrismaReceiver *YPrismaReceiverCallerSession) CVX() (common.Address, error) {
	return _YPrismaReceiver.Contract.CVX(&_YPrismaReceiver.CallOpts)
}

// PRISMA is a free data retrieval call binding the contract method 0x9b9b1c51.
//
// Solidity: function PRISMA() view returns(address)
func (_YPrismaReceiver *YPrismaReceiverCaller) PRISMA(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YPrismaReceiver.contract.Call(opts, &out, "PRISMA")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PRISMA is a free data retrieval call binding the contract method 0x9b9b1c51.
//
// Solidity: function PRISMA() view returns(address)
func (_YPrismaReceiver *YPrismaReceiverSession) PRISMA() (common.Address, error) {
	return _YPrismaReceiver.Contract.PRISMA(&_YPrismaReceiver.CallOpts)
}

// PRISMA is a free data retrieval call binding the contract method 0x9b9b1c51.
//
// Solidity: function PRISMA() view returns(address)
func (_YPrismaReceiver *YPrismaReceiverCallerSession) PRISMA() (common.Address, error) {
	return _YPrismaReceiver.Contract.PRISMA(&_YPrismaReceiver.CallOpts)
}

// PRISMACORE is a free data retrieval call binding the contract method 0xcc9641a8.
//
// Solidity: function PRISMA_CORE() view returns(address)
func (_YPrismaReceiver *YPrismaReceiverCaller) PRISMACORE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YPrismaReceiver.contract.Call(opts, &out, "PRISMA_CORE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PRISMACORE is a free data retrieval call binding the contract method 0xcc9641a8.
//
// Solidity: function PRISMA_CORE() view returns(address)
func (_YPrismaReceiver *YPrismaReceiverSession) PRISMACORE() (common.Address, error) {
	return _YPrismaReceiver.Contract.PRISMACORE(&_YPrismaReceiver.CallOpts)
}

// PRISMACORE is a free data retrieval call binding the contract method 0xcc9641a8.
//
// Solidity: function PRISMA_CORE() view returns(address)
func (_YPrismaReceiver *YPrismaReceiverCallerSession) PRISMACORE() (common.Address, error) {
	return _YPrismaReceiver.Contract.PRISMACORE(&_YPrismaReceiver.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_YPrismaReceiver *YPrismaReceiverCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YPrismaReceiver.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_YPrismaReceiver *YPrismaReceiverSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _YPrismaReceiver.Contract.Allowance(&_YPrismaReceiver.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_YPrismaReceiver *YPrismaReceiverCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _YPrismaReceiver.Contract.Allowance(&_YPrismaReceiver.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_YPrismaReceiver *YPrismaReceiverCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YPrismaReceiver.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_YPrismaReceiver *YPrismaReceiverSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _YPrismaReceiver.Contract.BalanceOf(&_YPrismaReceiver.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_YPrismaReceiver *YPrismaReceiverCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _YPrismaReceiver.Contract.BalanceOf(&_YPrismaReceiver.CallOpts, arg0)
}

// Booster is a free data retrieval call binding the contract method 0xc6def076.
//
// Solidity: function booster() view returns(address)
func (_YPrismaReceiver *YPrismaReceiverCaller) Booster(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YPrismaReceiver.contract.Call(opts, &out, "booster")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Booster is a free data retrieval call binding the contract method 0xc6def076.
//
// Solidity: function booster() view returns(address)
func (_YPrismaReceiver *YPrismaReceiverSession) Booster() (common.Address, error) {
	return _YPrismaReceiver.Contract.Booster(&_YPrismaReceiver.CallOpts)
}

// Booster is a free data retrieval call binding the contract method 0xc6def076.
//
// Solidity: function booster() view returns(address)
func (_YPrismaReceiver *YPrismaReceiverCallerSession) Booster() (common.Address, error) {
	return _YPrismaReceiver.Contract.Booster(&_YPrismaReceiver.CallOpts)
}

// ClaimableReward is a free data retrieval call binding the contract method 0xe9503425.
//
// Solidity: function claimableReward(address account) view returns(uint256 prismaAmount, uint256 crvAmount, uint256 cvxAmount)
func (_YPrismaReceiver *YPrismaReceiverCaller) ClaimableReward(opts *bind.CallOpts, account common.Address) (struct {
	PrismaAmount *big.Int
	CrvAmount    *big.Int
	CvxAmount    *big.Int
}, error) {
	var out []interface{}
	err := _YPrismaReceiver.contract.Call(opts, &out, "claimableReward", account)

	outstruct := new(struct {
		PrismaAmount *big.Int
		CrvAmount    *big.Int
		CvxAmount    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PrismaAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CrvAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CvxAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ClaimableReward is a free data retrieval call binding the contract method 0xe9503425.
//
// Solidity: function claimableReward(address account) view returns(uint256 prismaAmount, uint256 crvAmount, uint256 cvxAmount)
func (_YPrismaReceiver *YPrismaReceiverSession) ClaimableReward(account common.Address) (struct {
	PrismaAmount *big.Int
	CrvAmount    *big.Int
	CvxAmount    *big.Int
}, error) {
	return _YPrismaReceiver.Contract.ClaimableReward(&_YPrismaReceiver.CallOpts, account)
}

// ClaimableReward is a free data retrieval call binding the contract method 0xe9503425.
//
// Solidity: function claimableReward(address account) view returns(uint256 prismaAmount, uint256 crvAmount, uint256 cvxAmount)
func (_YPrismaReceiver *YPrismaReceiverCallerSession) ClaimableReward(account common.Address) (struct {
	PrismaAmount *big.Int
	CrvAmount    *big.Int
	CvxAmount    *big.Int
}, error) {
	return _YPrismaReceiver.Contract.ClaimableReward(&_YPrismaReceiver.CallOpts, account)
}

// CrvRewards is a free data retrieval call binding the contract method 0xb6bff295.
//
// Solidity: function crvRewards() view returns(address)
func (_YPrismaReceiver *YPrismaReceiverCaller) CrvRewards(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YPrismaReceiver.contract.Call(opts, &out, "crvRewards")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CrvRewards is a free data retrieval call binding the contract method 0xb6bff295.
//
// Solidity: function crvRewards() view returns(address)
func (_YPrismaReceiver *YPrismaReceiverSession) CrvRewards() (common.Address, error) {
	return _YPrismaReceiver.Contract.CrvRewards(&_YPrismaReceiver.CallOpts)
}

// CrvRewards is a free data retrieval call binding the contract method 0xb6bff295.
//
// Solidity: function crvRewards() view returns(address)
func (_YPrismaReceiver *YPrismaReceiverCallerSession) CrvRewards() (common.Address, error) {
	return _YPrismaReceiver.Contract.CrvRewards(&_YPrismaReceiver.CallOpts)
}

// CurveProxy is a free data retrieval call binding the contract method 0xb008c38d.
//
// Solidity: function curveProxy() view returns(address)
func (_YPrismaReceiver *YPrismaReceiverCaller) CurveProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YPrismaReceiver.contract.Call(opts, &out, "curveProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurveProxy is a free data retrieval call binding the contract method 0xb008c38d.
//
// Solidity: function curveProxy() view returns(address)
func (_YPrismaReceiver *YPrismaReceiverSession) CurveProxy() (common.Address, error) {
	return _YPrismaReceiver.Contract.CurveProxy(&_YPrismaReceiver.CallOpts)
}

// CurveProxy is a free data retrieval call binding the contract method 0xb008c38d.
//
// Solidity: function curveProxy() view returns(address)
func (_YPrismaReceiver *YPrismaReceiverCallerSession) CurveProxy() (common.Address, error) {
	return _YPrismaReceiver.Contract.CurveProxy(&_YPrismaReceiver.CallOpts)
}

// CvxRewards is a free data retrieval call binding the contract method 0xaa5ccb90.
//
// Solidity: function cvxRewards() view returns(address)
func (_YPrismaReceiver *YPrismaReceiverCaller) CvxRewards(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YPrismaReceiver.contract.Call(opts, &out, "cvxRewards")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CvxRewards is a free data retrieval call binding the contract method 0xaa5ccb90.
//
// Solidity: function cvxRewards() view returns(address)
func (_YPrismaReceiver *YPrismaReceiverSession) CvxRewards() (common.Address, error) {
	return _YPrismaReceiver.Contract.CvxRewards(&_YPrismaReceiver.CallOpts)
}

// CvxRewards is a free data retrieval call binding the contract method 0xaa5ccb90.
//
// Solidity: function cvxRewards() view returns(address)
func (_YPrismaReceiver *YPrismaReceiverCallerSession) CvxRewards() (common.Address, error) {
	return _YPrismaReceiver.Contract.CvxRewards(&_YPrismaReceiver.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_YPrismaReceiver *YPrismaReceiverCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YPrismaReceiver.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_YPrismaReceiver *YPrismaReceiverSession) Decimals() (*big.Int, error) {
	return _YPrismaReceiver.Contract.Decimals(&_YPrismaReceiver.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_YPrismaReceiver *YPrismaReceiverCallerSession) Decimals() (*big.Int, error) {
	return _YPrismaReceiver.Contract.Decimals(&_YPrismaReceiver.CallOpts)
}

// DepositPid is a free data retrieval call binding the contract method 0x89f4e7ae.
//
// Solidity: function depositPid() view returns(uint256)
func (_YPrismaReceiver *YPrismaReceiverCaller) DepositPid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YPrismaReceiver.contract.Call(opts, &out, "depositPid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositPid is a free data retrieval call binding the contract method 0x89f4e7ae.
//
// Solidity: function depositPid() view returns(uint256)
func (_YPrismaReceiver *YPrismaReceiverSession) DepositPid() (*big.Int, error) {
	return _YPrismaReceiver.Contract.DepositPid(&_YPrismaReceiver.CallOpts)
}

// DepositPid is a free data retrieval call binding the contract method 0x89f4e7ae.
//
// Solidity: function depositPid() view returns(uint256)
func (_YPrismaReceiver *YPrismaReceiverCallerSession) DepositPid() (*big.Int, error) {
	return _YPrismaReceiver.Contract.DepositPid(&_YPrismaReceiver.CallOpts)
}

// EmissionId is a free data retrieval call binding the contract method 0x47636371.
//
// Solidity: function emissionId() view returns(uint256)
func (_YPrismaReceiver *YPrismaReceiverCaller) EmissionId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YPrismaReceiver.contract.Call(opts, &out, "emissionId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EmissionId is a free data retrieval call binding the contract method 0x47636371.
//
// Solidity: function emissionId() view returns(uint256)
func (_YPrismaReceiver *YPrismaReceiverSession) EmissionId() (*big.Int, error) {
	return _YPrismaReceiver.Contract.EmissionId(&_YPrismaReceiver.CallOpts)
}

// EmissionId is a free data retrieval call binding the contract method 0x47636371.
//
// Solidity: function emissionId() view returns(uint256)
func (_YPrismaReceiver *YPrismaReceiverCallerSession) EmissionId() (*big.Int, error) {
	return _YPrismaReceiver.Contract.EmissionId(&_YPrismaReceiver.CallOpts)
}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_YPrismaReceiver *YPrismaReceiverCaller) Guardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YPrismaReceiver.contract.Call(opts, &out, "guardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_YPrismaReceiver *YPrismaReceiverSession) Guardian() (common.Address, error) {
	return _YPrismaReceiver.Contract.Guardian(&_YPrismaReceiver.CallOpts)
}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_YPrismaReceiver *YPrismaReceiverCallerSession) Guardian() (common.Address, error) {
	return _YPrismaReceiver.Contract.Guardian(&_YPrismaReceiver.CallOpts)
}

// LastCrvBalance is a free data retrieval call binding the contract method 0xd838b129.
//
// Solidity: function lastCrvBalance() view returns(uint128)
func (_YPrismaReceiver *YPrismaReceiverCaller) LastCrvBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YPrismaReceiver.contract.Call(opts, &out, "lastCrvBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastCrvBalance is a free data retrieval call binding the contract method 0xd838b129.
//
// Solidity: function lastCrvBalance() view returns(uint128)
func (_YPrismaReceiver *YPrismaReceiverSession) LastCrvBalance() (*big.Int, error) {
	return _YPrismaReceiver.Contract.LastCrvBalance(&_YPrismaReceiver.CallOpts)
}

// LastCrvBalance is a free data retrieval call binding the contract method 0xd838b129.
//
// Solidity: function lastCrvBalance() view returns(uint128)
func (_YPrismaReceiver *YPrismaReceiverCallerSession) LastCrvBalance() (*big.Int, error) {
	return _YPrismaReceiver.Contract.LastCrvBalance(&_YPrismaReceiver.CallOpts)
}

// LastCvxBalance is a free data retrieval call binding the contract method 0x561542b6.
//
// Solidity: function lastCvxBalance() view returns(uint128)
func (_YPrismaReceiver *YPrismaReceiverCaller) LastCvxBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YPrismaReceiver.contract.Call(opts, &out, "lastCvxBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastCvxBalance is a free data retrieval call binding the contract method 0x561542b6.
//
// Solidity: function lastCvxBalance() view returns(uint128)
func (_YPrismaReceiver *YPrismaReceiverSession) LastCvxBalance() (*big.Int, error) {
	return _YPrismaReceiver.Contract.LastCvxBalance(&_YPrismaReceiver.CallOpts)
}

// LastCvxBalance is a free data retrieval call binding the contract method 0x561542b6.
//
// Solidity: function lastCvxBalance() view returns(uint128)
func (_YPrismaReceiver *YPrismaReceiverCallerSession) LastCvxBalance() (*big.Int, error) {
	return _YPrismaReceiver.Contract.LastCvxBalance(&_YPrismaReceiver.CallOpts)
}

// LastUpdate is a free data retrieval call binding the contract method 0xc0463711.
//
// Solidity: function lastUpdate() view returns(uint32)
func (_YPrismaReceiver *YPrismaReceiverCaller) LastUpdate(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _YPrismaReceiver.contract.Call(opts, &out, "lastUpdate")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LastUpdate is a free data retrieval call binding the contract method 0xc0463711.
//
// Solidity: function lastUpdate() view returns(uint32)
func (_YPrismaReceiver *YPrismaReceiverSession) LastUpdate() (uint32, error) {
	return _YPrismaReceiver.Contract.LastUpdate(&_YPrismaReceiver.CallOpts)
}

// LastUpdate is a free data retrieval call binding the contract method 0xc0463711.
//
// Solidity: function lastUpdate() view returns(uint32)
func (_YPrismaReceiver *YPrismaReceiverCallerSession) LastUpdate() (uint32, error) {
	return _YPrismaReceiver.Contract.LastUpdate(&_YPrismaReceiver.CallOpts)
}

// LpToken is a free data retrieval call binding the contract method 0x5fcbd285.
//
// Solidity: function lpToken() view returns(address)
func (_YPrismaReceiver *YPrismaReceiverCaller) LpToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YPrismaReceiver.contract.Call(opts, &out, "lpToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LpToken is a free data retrieval call binding the contract method 0x5fcbd285.
//
// Solidity: function lpToken() view returns(address)
func (_YPrismaReceiver *YPrismaReceiverSession) LpToken() (common.Address, error) {
	return _YPrismaReceiver.Contract.LpToken(&_YPrismaReceiver.CallOpts)
}

// LpToken is a free data retrieval call binding the contract method 0x5fcbd285.
//
// Solidity: function lpToken() view returns(address)
func (_YPrismaReceiver *YPrismaReceiverCallerSession) LpToken() (common.Address, error) {
	return _YPrismaReceiver.Contract.LpToken(&_YPrismaReceiver.CallOpts)
}

// MaxWeeklyEmissionPct is a free data retrieval call binding the contract method 0x25dcafec.
//
// Solidity: function maxWeeklyEmissionPct() view returns(uint16)
func (_YPrismaReceiver *YPrismaReceiverCaller) MaxWeeklyEmissionPct(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _YPrismaReceiver.contract.Call(opts, &out, "maxWeeklyEmissionPct")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MaxWeeklyEmissionPct is a free data retrieval call binding the contract method 0x25dcafec.
//
// Solidity: function maxWeeklyEmissionPct() view returns(uint16)
func (_YPrismaReceiver *YPrismaReceiverSession) MaxWeeklyEmissionPct() (uint16, error) {
	return _YPrismaReceiver.Contract.MaxWeeklyEmissionPct(&_YPrismaReceiver.CallOpts)
}

// MaxWeeklyEmissionPct is a free data retrieval call binding the contract method 0x25dcafec.
//
// Solidity: function maxWeeklyEmissionPct() view returns(uint16)
func (_YPrismaReceiver *YPrismaReceiverCallerSession) MaxWeeklyEmissionPct() (uint16, error) {
	return _YPrismaReceiver.Contract.MaxWeeklyEmissionPct(&_YPrismaReceiver.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YPrismaReceiver *YPrismaReceiverCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _YPrismaReceiver.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YPrismaReceiver *YPrismaReceiverSession) Name() (string, error) {
	return _YPrismaReceiver.Contract.Name(&_YPrismaReceiver.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YPrismaReceiver *YPrismaReceiverCallerSession) Name() (string, error) {
	return _YPrismaReceiver.Contract.Name(&_YPrismaReceiver.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_YPrismaReceiver *YPrismaReceiverCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YPrismaReceiver.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_YPrismaReceiver *YPrismaReceiverSession) Owner() (common.Address, error) {
	return _YPrismaReceiver.Contract.Owner(&_YPrismaReceiver.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_YPrismaReceiver *YPrismaReceiverCallerSession) Owner() (common.Address, error) {
	return _YPrismaReceiver.Contract.Owner(&_YPrismaReceiver.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint32)
func (_YPrismaReceiver *YPrismaReceiverCaller) PeriodFinish(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _YPrismaReceiver.contract.Call(opts, &out, "periodFinish")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint32)
func (_YPrismaReceiver *YPrismaReceiverSession) PeriodFinish() (uint32, error) {
	return _YPrismaReceiver.Contract.PeriodFinish(&_YPrismaReceiver.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint32)
func (_YPrismaReceiver *YPrismaReceiverCallerSession) PeriodFinish() (uint32, error) {
	return _YPrismaReceiver.Contract.PeriodFinish(&_YPrismaReceiver.CallOpts)
}

// RewardIntegral is a free data retrieval call binding the contract method 0xb0ff51eb.
//
// Solidity: function rewardIntegral(uint256 ) view returns(uint256)
func (_YPrismaReceiver *YPrismaReceiverCaller) RewardIntegral(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YPrismaReceiver.contract.Call(opts, &out, "rewardIntegral", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardIntegral is a free data retrieval call binding the contract method 0xb0ff51eb.
//
// Solidity: function rewardIntegral(uint256 ) view returns(uint256)
func (_YPrismaReceiver *YPrismaReceiverSession) RewardIntegral(arg0 *big.Int) (*big.Int, error) {
	return _YPrismaReceiver.Contract.RewardIntegral(&_YPrismaReceiver.CallOpts, arg0)
}

// RewardIntegral is a free data retrieval call binding the contract method 0xb0ff51eb.
//
// Solidity: function rewardIntegral(uint256 ) view returns(uint256)
func (_YPrismaReceiver *YPrismaReceiverCallerSession) RewardIntegral(arg0 *big.Int) (*big.Int, error) {
	return _YPrismaReceiver.Contract.RewardIntegral(&_YPrismaReceiver.CallOpts, arg0)
}

// RewardIntegralFor is a free data retrieval call binding the contract method 0xec8f3fae.
//
// Solidity: function rewardIntegralFor(address , uint256 ) view returns(uint256)
func (_YPrismaReceiver *YPrismaReceiverCaller) RewardIntegralFor(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YPrismaReceiver.contract.Call(opts, &out, "rewardIntegralFor", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardIntegralFor is a free data retrieval call binding the contract method 0xec8f3fae.
//
// Solidity: function rewardIntegralFor(address , uint256 ) view returns(uint256)
func (_YPrismaReceiver *YPrismaReceiverSession) RewardIntegralFor(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _YPrismaReceiver.Contract.RewardIntegralFor(&_YPrismaReceiver.CallOpts, arg0, arg1)
}

// RewardIntegralFor is a free data retrieval call binding the contract method 0xec8f3fae.
//
// Solidity: function rewardIntegralFor(address , uint256 ) view returns(uint256)
func (_YPrismaReceiver *YPrismaReceiverCallerSession) RewardIntegralFor(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _YPrismaReceiver.Contract.RewardIntegralFor(&_YPrismaReceiver.CallOpts, arg0, arg1)
}

// RewardRate is a free data retrieval call binding the contract method 0xcea01962.
//
// Solidity: function rewardRate(uint256 ) view returns(uint128)
func (_YPrismaReceiver *YPrismaReceiverCaller) RewardRate(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YPrismaReceiver.contract.Call(opts, &out, "rewardRate", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRate is a free data retrieval call binding the contract method 0xcea01962.
//
// Solidity: function rewardRate(uint256 ) view returns(uint128)
func (_YPrismaReceiver *YPrismaReceiverSession) RewardRate(arg0 *big.Int) (*big.Int, error) {
	return _YPrismaReceiver.Contract.RewardRate(&_YPrismaReceiver.CallOpts, arg0)
}

// RewardRate is a free data retrieval call binding the contract method 0xcea01962.
//
// Solidity: function rewardRate(uint256 ) view returns(uint128)
func (_YPrismaReceiver *YPrismaReceiverCallerSession) RewardRate(arg0 *big.Int) (*big.Int, error) {
	return _YPrismaReceiver.Contract.RewardRate(&_YPrismaReceiver.CallOpts, arg0)
}

// StoredExcessEmissions is a free data retrieval call binding the contract method 0xebb7401b.
//
// Solidity: function storedExcessEmissions() view returns(uint128)
func (_YPrismaReceiver *YPrismaReceiverCaller) StoredExcessEmissions(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YPrismaReceiver.contract.Call(opts, &out, "storedExcessEmissions")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StoredExcessEmissions is a free data retrieval call binding the contract method 0xebb7401b.
//
// Solidity: function storedExcessEmissions() view returns(uint128)
func (_YPrismaReceiver *YPrismaReceiverSession) StoredExcessEmissions() (*big.Int, error) {
	return _YPrismaReceiver.Contract.StoredExcessEmissions(&_YPrismaReceiver.CallOpts)
}

// StoredExcessEmissions is a free data retrieval call binding the contract method 0xebb7401b.
//
// Solidity: function storedExcessEmissions() view returns(uint128)
func (_YPrismaReceiver *YPrismaReceiverCallerSession) StoredExcessEmissions() (*big.Int, error) {
	return _YPrismaReceiver.Contract.StoredExcessEmissions(&_YPrismaReceiver.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_YPrismaReceiver *YPrismaReceiverCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _YPrismaReceiver.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_YPrismaReceiver *YPrismaReceiverSession) Symbol() (string, error) {
	return _YPrismaReceiver.Contract.Symbol(&_YPrismaReceiver.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_YPrismaReceiver *YPrismaReceiverCallerSession) Symbol() (string, error) {
	return _YPrismaReceiver.Contract.Symbol(&_YPrismaReceiver.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_YPrismaReceiver *YPrismaReceiverCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YPrismaReceiver.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_YPrismaReceiver *YPrismaReceiverSession) TotalSupply() (*big.Int, error) {
	return _YPrismaReceiver.Contract.TotalSupply(&_YPrismaReceiver.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_YPrismaReceiver *YPrismaReceiverCallerSession) TotalSupply() (*big.Int, error) {
	return _YPrismaReceiver.Contract.TotalSupply(&_YPrismaReceiver.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_YPrismaReceiver *YPrismaReceiverCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YPrismaReceiver.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_YPrismaReceiver *YPrismaReceiverSession) Vault() (common.Address, error) {
	return _YPrismaReceiver.Contract.Vault(&_YPrismaReceiver.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_YPrismaReceiver *YPrismaReceiverCallerSession) Vault() (common.Address, error) {
	return _YPrismaReceiver.Contract.Vault(&_YPrismaReceiver.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_YPrismaReceiver *YPrismaReceiverTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _YPrismaReceiver.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_YPrismaReceiver *YPrismaReceiverSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _YPrismaReceiver.Contract.Approve(&_YPrismaReceiver.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_YPrismaReceiver *YPrismaReceiverTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _YPrismaReceiver.Contract.Approve(&_YPrismaReceiver.TransactOpts, _spender, _value)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address receiver) returns(uint256 prismaAmount, uint256 crvAmount, uint256 cvxAmount)
func (_YPrismaReceiver *YPrismaReceiverTransactor) ClaimReward(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _YPrismaReceiver.contract.Transact(opts, "claimReward", receiver)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address receiver) returns(uint256 prismaAmount, uint256 crvAmount, uint256 cvxAmount)
func (_YPrismaReceiver *YPrismaReceiverSession) ClaimReward(receiver common.Address) (*types.Transaction, error) {
	return _YPrismaReceiver.Contract.ClaimReward(&_YPrismaReceiver.TransactOpts, receiver)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address receiver) returns(uint256 prismaAmount, uint256 crvAmount, uint256 cvxAmount)
func (_YPrismaReceiver *YPrismaReceiverTransactorSession) ClaimReward(receiver common.Address) (*types.Transaction, error) {
	return _YPrismaReceiver.Contract.ClaimReward(&_YPrismaReceiver.TransactOpts, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address receiver, uint256 amount) returns(bool)
func (_YPrismaReceiver *YPrismaReceiverTransactor) Deposit(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YPrismaReceiver.contract.Transact(opts, "deposit", receiver, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address receiver, uint256 amount) returns(bool)
func (_YPrismaReceiver *YPrismaReceiverSession) Deposit(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YPrismaReceiver.Contract.Deposit(&_YPrismaReceiver.TransactOpts, receiver, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address receiver, uint256 amount) returns(bool)
func (_YPrismaReceiver *YPrismaReceiverTransactorSession) Deposit(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YPrismaReceiver.Contract.Deposit(&_YPrismaReceiver.TransactOpts, receiver, amount)
}

// FetchRewards is a paid mutator transaction binding the contract method 0xed21aa89.
//
// Solidity: function fetchRewards() returns()
func (_YPrismaReceiver *YPrismaReceiverTransactor) FetchRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YPrismaReceiver.contract.Transact(opts, "fetchRewards")
}

// FetchRewards is a paid mutator transaction binding the contract method 0xed21aa89.
//
// Solidity: function fetchRewards() returns()
func (_YPrismaReceiver *YPrismaReceiverSession) FetchRewards() (*types.Transaction, error) {
	return _YPrismaReceiver.Contract.FetchRewards(&_YPrismaReceiver.TransactOpts)
}

// FetchRewards is a paid mutator transaction binding the contract method 0xed21aa89.
//
// Solidity: function fetchRewards() returns()
func (_YPrismaReceiver *YPrismaReceiverTransactorSession) FetchRewards() (*types.Transaction, error) {
	return _YPrismaReceiver.Contract.FetchRewards(&_YPrismaReceiver.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 pid) returns()
func (_YPrismaReceiver *YPrismaReceiverTransactor) Initialize(opts *bind.TransactOpts, pid *big.Int) (*types.Transaction, error) {
	return _YPrismaReceiver.contract.Transact(opts, "initialize", pid)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 pid) returns()
func (_YPrismaReceiver *YPrismaReceiverSession) Initialize(pid *big.Int) (*types.Transaction, error) {
	return _YPrismaReceiver.Contract.Initialize(&_YPrismaReceiver.TransactOpts, pid)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 pid) returns()
func (_YPrismaReceiver *YPrismaReceiverTransactorSession) Initialize(pid *big.Int) (*types.Transaction, error) {
	return _YPrismaReceiver.Contract.Initialize(&_YPrismaReceiver.TransactOpts, pid)
}

// NotifyRegisteredId is a paid mutator transaction binding the contract method 0x5bbe8aad.
//
// Solidity: function notifyRegisteredId(uint256[] assignedIds) returns(bool)
func (_YPrismaReceiver *YPrismaReceiverTransactor) NotifyRegisteredId(opts *bind.TransactOpts, assignedIds []*big.Int) (*types.Transaction, error) {
	return _YPrismaReceiver.contract.Transact(opts, "notifyRegisteredId", assignedIds)
}

// NotifyRegisteredId is a paid mutator transaction binding the contract method 0x5bbe8aad.
//
// Solidity: function notifyRegisteredId(uint256[] assignedIds) returns(bool)
func (_YPrismaReceiver *YPrismaReceiverSession) NotifyRegisteredId(assignedIds []*big.Int) (*types.Transaction, error) {
	return _YPrismaReceiver.Contract.NotifyRegisteredId(&_YPrismaReceiver.TransactOpts, assignedIds)
}

// NotifyRegisteredId is a paid mutator transaction binding the contract method 0x5bbe8aad.
//
// Solidity: function notifyRegisteredId(uint256[] assignedIds) returns(bool)
func (_YPrismaReceiver *YPrismaReceiverTransactorSession) NotifyRegisteredId(assignedIds []*big.Int) (*types.Transaction, error) {
	return _YPrismaReceiver.Contract.NotifyRegisteredId(&_YPrismaReceiver.TransactOpts, assignedIds)
}

// PushExcessEmissions is a paid mutator transaction binding the contract method 0x36c6e318.
//
// Solidity: function pushExcessEmissions() returns()
func (_YPrismaReceiver *YPrismaReceiverTransactor) PushExcessEmissions(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YPrismaReceiver.contract.Transact(opts, "pushExcessEmissions")
}

// PushExcessEmissions is a paid mutator transaction binding the contract method 0x36c6e318.
//
// Solidity: function pushExcessEmissions() returns()
func (_YPrismaReceiver *YPrismaReceiverSession) PushExcessEmissions() (*types.Transaction, error) {
	return _YPrismaReceiver.Contract.PushExcessEmissions(&_YPrismaReceiver.TransactOpts)
}

// PushExcessEmissions is a paid mutator transaction binding the contract method 0x36c6e318.
//
// Solidity: function pushExcessEmissions() returns()
func (_YPrismaReceiver *YPrismaReceiverTransactorSession) PushExcessEmissions() (*types.Transaction, error) {
	return _YPrismaReceiver.Contract.PushExcessEmissions(&_YPrismaReceiver.TransactOpts)
}

// SetMaxWeeklyEmissionPct is a paid mutator transaction binding the contract method 0xf60c5eba.
//
// Solidity: function setMaxWeeklyEmissionPct(uint16 _maxWeeklyEmissionPct) returns(bool)
func (_YPrismaReceiver *YPrismaReceiverTransactor) SetMaxWeeklyEmissionPct(opts *bind.TransactOpts, _maxWeeklyEmissionPct uint16) (*types.Transaction, error) {
	return _YPrismaReceiver.contract.Transact(opts, "setMaxWeeklyEmissionPct", _maxWeeklyEmissionPct)
}

// SetMaxWeeklyEmissionPct is a paid mutator transaction binding the contract method 0xf60c5eba.
//
// Solidity: function setMaxWeeklyEmissionPct(uint16 _maxWeeklyEmissionPct) returns(bool)
func (_YPrismaReceiver *YPrismaReceiverSession) SetMaxWeeklyEmissionPct(_maxWeeklyEmissionPct uint16) (*types.Transaction, error) {
	return _YPrismaReceiver.Contract.SetMaxWeeklyEmissionPct(&_YPrismaReceiver.TransactOpts, _maxWeeklyEmissionPct)
}

// SetMaxWeeklyEmissionPct is a paid mutator transaction binding the contract method 0xf60c5eba.
//
// Solidity: function setMaxWeeklyEmissionPct(uint16 _maxWeeklyEmissionPct) returns(bool)
func (_YPrismaReceiver *YPrismaReceiverTransactorSession) SetMaxWeeklyEmissionPct(_maxWeeklyEmissionPct uint16) (*types.Transaction, error) {
	return _YPrismaReceiver.Contract.SetMaxWeeklyEmissionPct(&_YPrismaReceiver.TransactOpts, _maxWeeklyEmissionPct)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_YPrismaReceiver *YPrismaReceiverTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _YPrismaReceiver.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_YPrismaReceiver *YPrismaReceiverSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _YPrismaReceiver.Contract.Transfer(&_YPrismaReceiver.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_YPrismaReceiver *YPrismaReceiverTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _YPrismaReceiver.Contract.Transfer(&_YPrismaReceiver.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_YPrismaReceiver *YPrismaReceiverTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _YPrismaReceiver.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_YPrismaReceiver *YPrismaReceiverSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _YPrismaReceiver.Contract.TransferFrom(&_YPrismaReceiver.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_YPrismaReceiver *YPrismaReceiverTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _YPrismaReceiver.Contract.TransferFrom(&_YPrismaReceiver.TransactOpts, _from, _to, _value)
}

// VaultClaimReward is a paid mutator transaction binding the contract method 0xf0582038.
//
// Solidity: function vaultClaimReward(address claimant, address receiver) returns(uint256)
func (_YPrismaReceiver *YPrismaReceiverTransactor) VaultClaimReward(opts *bind.TransactOpts, claimant common.Address, receiver common.Address) (*types.Transaction, error) {
	return _YPrismaReceiver.contract.Transact(opts, "vaultClaimReward", claimant, receiver)
}

// VaultClaimReward is a paid mutator transaction binding the contract method 0xf0582038.
//
// Solidity: function vaultClaimReward(address claimant, address receiver) returns(uint256)
func (_YPrismaReceiver *YPrismaReceiverSession) VaultClaimReward(claimant common.Address, receiver common.Address) (*types.Transaction, error) {
	return _YPrismaReceiver.Contract.VaultClaimReward(&_YPrismaReceiver.TransactOpts, claimant, receiver)
}

// VaultClaimReward is a paid mutator transaction binding the contract method 0xf0582038.
//
// Solidity: function vaultClaimReward(address claimant, address receiver) returns(uint256)
func (_YPrismaReceiver *YPrismaReceiverTransactorSession) VaultClaimReward(claimant common.Address, receiver common.Address) (*types.Transaction, error) {
	return _YPrismaReceiver.Contract.VaultClaimReward(&_YPrismaReceiver.TransactOpts, claimant, receiver)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address receiver, uint256 amount) returns(bool)
func (_YPrismaReceiver *YPrismaReceiverTransactor) Withdraw(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YPrismaReceiver.contract.Transact(opts, "withdraw", receiver, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address receiver, uint256 amount) returns(bool)
func (_YPrismaReceiver *YPrismaReceiverSession) Withdraw(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YPrismaReceiver.Contract.Withdraw(&_YPrismaReceiver.TransactOpts, receiver, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address receiver, uint256 amount) returns(bool)
func (_YPrismaReceiver *YPrismaReceiverTransactorSession) Withdraw(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YPrismaReceiver.Contract.Withdraw(&_YPrismaReceiver.TransactOpts, receiver, amount)
}

// YPrismaReceiverApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the YPrismaReceiver contract.
type YPrismaReceiverApprovalIterator struct {
	Event *YPrismaReceiverApproval // Event containing the contract specifics and raw log

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
func (it *YPrismaReceiverApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YPrismaReceiverApproval)
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
		it.Event = new(YPrismaReceiverApproval)
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
func (it *YPrismaReceiverApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YPrismaReceiverApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YPrismaReceiverApproval represents a Approval event raised by the YPrismaReceiver contract.
type YPrismaReceiverApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_YPrismaReceiver *YPrismaReceiverFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*YPrismaReceiverApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _YPrismaReceiver.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &YPrismaReceiverApprovalIterator{contract: _YPrismaReceiver.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_YPrismaReceiver *YPrismaReceiverFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *YPrismaReceiverApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _YPrismaReceiver.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YPrismaReceiverApproval)
				if err := _YPrismaReceiver.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_YPrismaReceiver *YPrismaReceiverFilterer) ParseApproval(log types.Log) (*YPrismaReceiverApproval, error) {
	event := new(YPrismaReceiverApproval)
	if err := _YPrismaReceiver.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YPrismaReceiverLPTokenDepositedIterator is returned from FilterLPTokenDeposited and is used to iterate over the raw logs and unpacked data for LPTokenDeposited events raised by the YPrismaReceiver contract.
type YPrismaReceiverLPTokenDepositedIterator struct {
	Event *YPrismaReceiverLPTokenDeposited // Event containing the contract specifics and raw log

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
func (it *YPrismaReceiverLPTokenDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YPrismaReceiverLPTokenDeposited)
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
		it.Event = new(YPrismaReceiverLPTokenDeposited)
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
func (it *YPrismaReceiverLPTokenDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YPrismaReceiverLPTokenDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YPrismaReceiverLPTokenDeposited represents a LPTokenDeposited event raised by the YPrismaReceiver contract.
type YPrismaReceiverLPTokenDeposited struct {
	LpToken  common.Address
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLPTokenDeposited is a free log retrieval operation binding the contract event 0x26372de6decc0ca1770f6d75300a8df13c11cf3a9b592071a221f546f201810d.
//
// Solidity: event LPTokenDeposited(address indexed lpToken, address indexed receiver, uint256 amount)
func (_YPrismaReceiver *YPrismaReceiverFilterer) FilterLPTokenDeposited(opts *bind.FilterOpts, lpToken []common.Address, receiver []common.Address) (*YPrismaReceiverLPTokenDepositedIterator, error) {

	var lpTokenRule []interface{}
	for _, lpTokenItem := range lpToken {
		lpTokenRule = append(lpTokenRule, lpTokenItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _YPrismaReceiver.contract.FilterLogs(opts, "LPTokenDeposited", lpTokenRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &YPrismaReceiverLPTokenDepositedIterator{contract: _YPrismaReceiver.contract, event: "LPTokenDeposited", logs: logs, sub: sub}, nil
}

// WatchLPTokenDeposited is a free log subscription operation binding the contract event 0x26372de6decc0ca1770f6d75300a8df13c11cf3a9b592071a221f546f201810d.
//
// Solidity: event LPTokenDeposited(address indexed lpToken, address indexed receiver, uint256 amount)
func (_YPrismaReceiver *YPrismaReceiverFilterer) WatchLPTokenDeposited(opts *bind.WatchOpts, sink chan<- *YPrismaReceiverLPTokenDeposited, lpToken []common.Address, receiver []common.Address) (event.Subscription, error) {

	var lpTokenRule []interface{}
	for _, lpTokenItem := range lpToken {
		lpTokenRule = append(lpTokenRule, lpTokenItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _YPrismaReceiver.contract.WatchLogs(opts, "LPTokenDeposited", lpTokenRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YPrismaReceiverLPTokenDeposited)
				if err := _YPrismaReceiver.contract.UnpackLog(event, "LPTokenDeposited", log); err != nil {
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

// ParseLPTokenDeposited is a log parse operation binding the contract event 0x26372de6decc0ca1770f6d75300a8df13c11cf3a9b592071a221f546f201810d.
//
// Solidity: event LPTokenDeposited(address indexed lpToken, address indexed receiver, uint256 amount)
func (_YPrismaReceiver *YPrismaReceiverFilterer) ParseLPTokenDeposited(log types.Log) (*YPrismaReceiverLPTokenDeposited, error) {
	event := new(YPrismaReceiverLPTokenDeposited)
	if err := _YPrismaReceiver.contract.UnpackLog(event, "LPTokenDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YPrismaReceiverLPTokenWithdrawnIterator is returned from FilterLPTokenWithdrawn and is used to iterate over the raw logs and unpacked data for LPTokenWithdrawn events raised by the YPrismaReceiver contract.
type YPrismaReceiverLPTokenWithdrawnIterator struct {
	Event *YPrismaReceiverLPTokenWithdrawn // Event containing the contract specifics and raw log

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
func (it *YPrismaReceiverLPTokenWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YPrismaReceiverLPTokenWithdrawn)
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
		it.Event = new(YPrismaReceiverLPTokenWithdrawn)
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
func (it *YPrismaReceiverLPTokenWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YPrismaReceiverLPTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YPrismaReceiverLPTokenWithdrawn represents a LPTokenWithdrawn event raised by the YPrismaReceiver contract.
type YPrismaReceiverLPTokenWithdrawn struct {
	LpToken  common.Address
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLPTokenWithdrawn is a free log retrieval operation binding the contract event 0x2d7689b3add04b23136b575399b9bb324f6705d949a2cf734850634803751dd3.
//
// Solidity: event LPTokenWithdrawn(address indexed lpToken, address indexed receiver, uint256 amount)
func (_YPrismaReceiver *YPrismaReceiverFilterer) FilterLPTokenWithdrawn(opts *bind.FilterOpts, lpToken []common.Address, receiver []common.Address) (*YPrismaReceiverLPTokenWithdrawnIterator, error) {

	var lpTokenRule []interface{}
	for _, lpTokenItem := range lpToken {
		lpTokenRule = append(lpTokenRule, lpTokenItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _YPrismaReceiver.contract.FilterLogs(opts, "LPTokenWithdrawn", lpTokenRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &YPrismaReceiverLPTokenWithdrawnIterator{contract: _YPrismaReceiver.contract, event: "LPTokenWithdrawn", logs: logs, sub: sub}, nil
}

// WatchLPTokenWithdrawn is a free log subscription operation binding the contract event 0x2d7689b3add04b23136b575399b9bb324f6705d949a2cf734850634803751dd3.
//
// Solidity: event LPTokenWithdrawn(address indexed lpToken, address indexed receiver, uint256 amount)
func (_YPrismaReceiver *YPrismaReceiverFilterer) WatchLPTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *YPrismaReceiverLPTokenWithdrawn, lpToken []common.Address, receiver []common.Address) (event.Subscription, error) {

	var lpTokenRule []interface{}
	for _, lpTokenItem := range lpToken {
		lpTokenRule = append(lpTokenRule, lpTokenItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _YPrismaReceiver.contract.WatchLogs(opts, "LPTokenWithdrawn", lpTokenRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YPrismaReceiverLPTokenWithdrawn)
				if err := _YPrismaReceiver.contract.UnpackLog(event, "LPTokenWithdrawn", log); err != nil {
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

// ParseLPTokenWithdrawn is a log parse operation binding the contract event 0x2d7689b3add04b23136b575399b9bb324f6705d949a2cf734850634803751dd3.
//
// Solidity: event LPTokenWithdrawn(address indexed lpToken, address indexed receiver, uint256 amount)
func (_YPrismaReceiver *YPrismaReceiverFilterer) ParseLPTokenWithdrawn(log types.Log) (*YPrismaReceiverLPTokenWithdrawn, error) {
	event := new(YPrismaReceiverLPTokenWithdrawn)
	if err := _YPrismaReceiver.contract.UnpackLog(event, "LPTokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YPrismaReceiverMaxWeeklyEmissionPctSetIterator is returned from FilterMaxWeeklyEmissionPctSet and is used to iterate over the raw logs and unpacked data for MaxWeeklyEmissionPctSet events raised by the YPrismaReceiver contract.
type YPrismaReceiverMaxWeeklyEmissionPctSetIterator struct {
	Event *YPrismaReceiverMaxWeeklyEmissionPctSet // Event containing the contract specifics and raw log

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
func (it *YPrismaReceiverMaxWeeklyEmissionPctSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YPrismaReceiverMaxWeeklyEmissionPctSet)
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
		it.Event = new(YPrismaReceiverMaxWeeklyEmissionPctSet)
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
func (it *YPrismaReceiverMaxWeeklyEmissionPctSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YPrismaReceiverMaxWeeklyEmissionPctSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YPrismaReceiverMaxWeeklyEmissionPctSet represents a MaxWeeklyEmissionPctSet event raised by the YPrismaReceiver contract.
type YPrismaReceiverMaxWeeklyEmissionPctSet struct {
	Pct *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMaxWeeklyEmissionPctSet is a free log retrieval operation binding the contract event 0xa4aeb2ea39782dbbd2a2796104f0207098cc3ef580a9ecacfa3ba33c0f1ddcaa.
//
// Solidity: event MaxWeeklyEmissionPctSet(uint256 pct)
func (_YPrismaReceiver *YPrismaReceiverFilterer) FilterMaxWeeklyEmissionPctSet(opts *bind.FilterOpts) (*YPrismaReceiverMaxWeeklyEmissionPctSetIterator, error) {

	logs, sub, err := _YPrismaReceiver.contract.FilterLogs(opts, "MaxWeeklyEmissionPctSet")
	if err != nil {
		return nil, err
	}
	return &YPrismaReceiverMaxWeeklyEmissionPctSetIterator{contract: _YPrismaReceiver.contract, event: "MaxWeeklyEmissionPctSet", logs: logs, sub: sub}, nil
}

// WatchMaxWeeklyEmissionPctSet is a free log subscription operation binding the contract event 0xa4aeb2ea39782dbbd2a2796104f0207098cc3ef580a9ecacfa3ba33c0f1ddcaa.
//
// Solidity: event MaxWeeklyEmissionPctSet(uint256 pct)
func (_YPrismaReceiver *YPrismaReceiverFilterer) WatchMaxWeeklyEmissionPctSet(opts *bind.WatchOpts, sink chan<- *YPrismaReceiverMaxWeeklyEmissionPctSet) (event.Subscription, error) {

	logs, sub, err := _YPrismaReceiver.contract.WatchLogs(opts, "MaxWeeklyEmissionPctSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YPrismaReceiverMaxWeeklyEmissionPctSet)
				if err := _YPrismaReceiver.contract.UnpackLog(event, "MaxWeeklyEmissionPctSet", log); err != nil {
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

// ParseMaxWeeklyEmissionPctSet is a log parse operation binding the contract event 0xa4aeb2ea39782dbbd2a2796104f0207098cc3ef580a9ecacfa3ba33c0f1ddcaa.
//
// Solidity: event MaxWeeklyEmissionPctSet(uint256 pct)
func (_YPrismaReceiver *YPrismaReceiverFilterer) ParseMaxWeeklyEmissionPctSet(log types.Log) (*YPrismaReceiverMaxWeeklyEmissionPctSet, error) {
	event := new(YPrismaReceiverMaxWeeklyEmissionPctSet)
	if err := _YPrismaReceiver.contract.UnpackLog(event, "MaxWeeklyEmissionPctSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YPrismaReceiverMaxWeeklyEmissionsExceededIterator is returned from FilterMaxWeeklyEmissionsExceeded and is used to iterate over the raw logs and unpacked data for MaxWeeklyEmissionsExceeded events raised by the YPrismaReceiver contract.
type YPrismaReceiverMaxWeeklyEmissionsExceededIterator struct {
	Event *YPrismaReceiverMaxWeeklyEmissionsExceeded // Event containing the contract specifics and raw log

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
func (it *YPrismaReceiverMaxWeeklyEmissionsExceededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YPrismaReceiverMaxWeeklyEmissionsExceeded)
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
		it.Event = new(YPrismaReceiverMaxWeeklyEmissionsExceeded)
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
func (it *YPrismaReceiverMaxWeeklyEmissionsExceededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YPrismaReceiverMaxWeeklyEmissionsExceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YPrismaReceiverMaxWeeklyEmissionsExceeded represents a MaxWeeklyEmissionsExceeded event raised by the YPrismaReceiver contract.
type YPrismaReceiverMaxWeeklyEmissionsExceeded struct {
	Allocated  *big.Int
	MaxAllowed *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMaxWeeklyEmissionsExceeded is a free log retrieval operation binding the contract event 0xbfbf172885c7567a7d54f21f5384146de64bed6c0a60b10d97a813ff9a712cd1.
//
// Solidity: event MaxWeeklyEmissionsExceeded(uint256 allocated, uint256 maxAllowed)
func (_YPrismaReceiver *YPrismaReceiverFilterer) FilterMaxWeeklyEmissionsExceeded(opts *bind.FilterOpts) (*YPrismaReceiverMaxWeeklyEmissionsExceededIterator, error) {

	logs, sub, err := _YPrismaReceiver.contract.FilterLogs(opts, "MaxWeeklyEmissionsExceeded")
	if err != nil {
		return nil, err
	}
	return &YPrismaReceiverMaxWeeklyEmissionsExceededIterator{contract: _YPrismaReceiver.contract, event: "MaxWeeklyEmissionsExceeded", logs: logs, sub: sub}, nil
}

// WatchMaxWeeklyEmissionsExceeded is a free log subscription operation binding the contract event 0xbfbf172885c7567a7d54f21f5384146de64bed6c0a60b10d97a813ff9a712cd1.
//
// Solidity: event MaxWeeklyEmissionsExceeded(uint256 allocated, uint256 maxAllowed)
func (_YPrismaReceiver *YPrismaReceiverFilterer) WatchMaxWeeklyEmissionsExceeded(opts *bind.WatchOpts, sink chan<- *YPrismaReceiverMaxWeeklyEmissionsExceeded) (event.Subscription, error) {

	logs, sub, err := _YPrismaReceiver.contract.WatchLogs(opts, "MaxWeeklyEmissionsExceeded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YPrismaReceiverMaxWeeklyEmissionsExceeded)
				if err := _YPrismaReceiver.contract.UnpackLog(event, "MaxWeeklyEmissionsExceeded", log); err != nil {
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

// ParseMaxWeeklyEmissionsExceeded is a log parse operation binding the contract event 0xbfbf172885c7567a7d54f21f5384146de64bed6c0a60b10d97a813ff9a712cd1.
//
// Solidity: event MaxWeeklyEmissionsExceeded(uint256 allocated, uint256 maxAllowed)
func (_YPrismaReceiver *YPrismaReceiverFilterer) ParseMaxWeeklyEmissionsExceeded(log types.Log) (*YPrismaReceiverMaxWeeklyEmissionsExceeded, error) {
	event := new(YPrismaReceiverMaxWeeklyEmissionsExceeded)
	if err := _YPrismaReceiver.contract.UnpackLog(event, "MaxWeeklyEmissionsExceeded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YPrismaReceiverRewardClaimedIterator is returned from FilterRewardClaimed and is used to iterate over the raw logs and unpacked data for RewardClaimed events raised by the YPrismaReceiver contract.
type YPrismaReceiverRewardClaimedIterator struct {
	Event *YPrismaReceiverRewardClaimed // Event containing the contract specifics and raw log

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
func (it *YPrismaReceiverRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YPrismaReceiverRewardClaimed)
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
		it.Event = new(YPrismaReceiverRewardClaimed)
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
func (it *YPrismaReceiverRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YPrismaReceiverRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YPrismaReceiverRewardClaimed represents a RewardClaimed event raised by the YPrismaReceiver contract.
type YPrismaReceiverRewardClaimed struct {
	Receiver     common.Address
	PrismaAmount *big.Int
	CrvAmount    *big.Int
	CvxAmount    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRewardClaimed is a free log retrieval operation binding the contract event 0x812be816db82c66cd18ca8457005cd84689642d8ac4d38599cc6af444a2dc72a.
//
// Solidity: event RewardClaimed(address indexed receiver, uint256 prismaAmount, uint256 crvAmount, uint256 cvxAmount)
func (_YPrismaReceiver *YPrismaReceiverFilterer) FilterRewardClaimed(opts *bind.FilterOpts, receiver []common.Address) (*YPrismaReceiverRewardClaimedIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _YPrismaReceiver.contract.FilterLogs(opts, "RewardClaimed", receiverRule)
	if err != nil {
		return nil, err
	}
	return &YPrismaReceiverRewardClaimedIterator{contract: _YPrismaReceiver.contract, event: "RewardClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardClaimed is a free log subscription operation binding the contract event 0x812be816db82c66cd18ca8457005cd84689642d8ac4d38599cc6af444a2dc72a.
//
// Solidity: event RewardClaimed(address indexed receiver, uint256 prismaAmount, uint256 crvAmount, uint256 cvxAmount)
func (_YPrismaReceiver *YPrismaReceiverFilterer) WatchRewardClaimed(opts *bind.WatchOpts, sink chan<- *YPrismaReceiverRewardClaimed, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _YPrismaReceiver.contract.WatchLogs(opts, "RewardClaimed", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YPrismaReceiverRewardClaimed)
				if err := _YPrismaReceiver.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
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

// ParseRewardClaimed is a log parse operation binding the contract event 0x812be816db82c66cd18ca8457005cd84689642d8ac4d38599cc6af444a2dc72a.
//
// Solidity: event RewardClaimed(address indexed receiver, uint256 prismaAmount, uint256 crvAmount, uint256 cvxAmount)
func (_YPrismaReceiver *YPrismaReceiverFilterer) ParseRewardClaimed(log types.Log) (*YPrismaReceiverRewardClaimed, error) {
	event := new(YPrismaReceiverRewardClaimed)
	if err := _YPrismaReceiver.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YPrismaReceiverTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the YPrismaReceiver contract.
type YPrismaReceiverTransferIterator struct {
	Event *YPrismaReceiverTransfer // Event containing the contract specifics and raw log

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
func (it *YPrismaReceiverTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YPrismaReceiverTransfer)
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
		it.Event = new(YPrismaReceiverTransfer)
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
func (it *YPrismaReceiverTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YPrismaReceiverTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YPrismaReceiverTransfer represents a Transfer event raised by the YPrismaReceiver contract.
type YPrismaReceiverTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_YPrismaReceiver *YPrismaReceiverFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*YPrismaReceiverTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _YPrismaReceiver.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &YPrismaReceiverTransferIterator{contract: _YPrismaReceiver.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_YPrismaReceiver *YPrismaReceiverFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *YPrismaReceiverTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _YPrismaReceiver.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YPrismaReceiverTransfer)
				if err := _YPrismaReceiver.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_YPrismaReceiver *YPrismaReceiverFilterer) ParseTransfer(log types.Log) (*YPrismaReceiverTransfer, error) {
	event := new(YPrismaReceiverTransfer)
	if err := _YPrismaReceiver.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
