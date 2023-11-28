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

// HealthCheckAccountantFee is an auto generated low-level Go binding around an user-defined struct.
type HealthCheckAccountantFee struct {
	ManagementFee  uint16
	PerformanceFee uint16
	RefundRatio    uint16
	MaxFee         uint16
	MaxGain        uint16
	MaxLoss        uint16
}

// AccountantMetaData contains all meta data concerning the Accountant contract.
var AccountantMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"defaultManagement\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultPerformance\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultRefund\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultMaxFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultMaxGain\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultMaxLoss\",\"type\":\"uint16\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"name\":\"DistributeRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeManager\",\"type\":\"address\"}],\"name\":\"NewFeeManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"}],\"name\":\"RemovedCustomFeeConfig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"futureFeeManager\",\"type\":\"address\"}],\"name\":\"SetFutureFeeManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"managementFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"performanceFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"refundRatio\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"maxFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"maxGain\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"maxLoss\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structHealthCheckAccountant.Fee\",\"name\":\"custom_config\",\"type\":\"tuple\"}],\"name\":\"UpdateCustomFeeConfig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"managementFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"performanceFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"refundRatio\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"maxFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"maxGain\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"maxLoss\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structHealthCheckAccountant.Fee\",\"name\":\"defaultFeeConfig\",\"type\":\"tuple\"}],\"name\":\"UpdateDefaultFeeConfig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldFeeRecipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newFeeRecipient\",\"type\":\"address\"}],\"name\":\"UpdateFeeRecipient\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxLoss\",\"type\":\"uint256\"}],\"name\":\"UpdateMaxLoss\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"refund\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UpdateRefund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newVaultManager\",\"type\":\"address\"}],\"name\":\"UpdateVaultManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumHealthCheckAccountant.ChangeType\",\"name\":\"change\",\"type\":\"uint8\"}],\"name\":\"VaultChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptFeeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"addVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"customConfig\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"managementFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"performanceFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"refundRatio\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"maxFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"maxGain\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"maxLoss\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultConfig\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"managementFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"performanceFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"refundRatio\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"maxFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"maxGain\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"maxLoss\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"distribute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"distribute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"futureFeeManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"managementFeeThreshold\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxLoss\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"performanceFeeThreshold\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"refund\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"refund\",\"type\":\"bool\"},{\"internalType\":\"uint248\",\"name\":\"amount\",\"type\":\"uint248\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"}],\"name\":\"removeCustomConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"removeVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"loss\",\"type\":\"uint256\"}],\"name\":\"report\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalRefunds\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"customManagement\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"customPerformance\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"customRefund\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"customMaxFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"customMaxGain\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"customMaxLoss\",\"type\":\"uint16\"}],\"name\":\"setCustomConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newFeeRecipient\",\"type\":\"address\"}],\"name\":\"setFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_futureFeeManager\",\"type\":\"address\"}],\"name\":\"setFutureFeeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxLoss\",\"type\":\"uint256\"}],\"name\":\"setMaxLoss\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_refund\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setRefund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newVaultManager\",\"type\":\"address\"}],\"name\":\"setVaultManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"defaultManagement\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultPerformance\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultRefund\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultMaxFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultMaxGain\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultMaxLoss\",\"type\":\"uint16\"}],\"name\":\"updateDefaultConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"}],\"name\":\"useCustomConfig\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"vaults\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawUnderlying\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AccountantABI is the input ABI used to generate the binding from.
// Deprecated: Use AccountantMetaData.ABI instead.
var AccountantABI = AccountantMetaData.ABI

// Accountant is an auto generated Go binding around an Ethereum contract.
type Accountant struct {
	AccountantCaller     // Read-only binding to the contract
	AccountantTransactor // Write-only binding to the contract
	AccountantFilterer   // Log filterer for contract events
}

// AccountantCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountantCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountantTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountantTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountantFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccountantFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountantSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccountantSession struct {
	Contract     *Accountant       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccountantCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccountantCallerSession struct {
	Contract *AccountantCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AccountantTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccountantTransactorSession struct {
	Contract     *AccountantTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AccountantRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccountantRaw struct {
	Contract *Accountant // Generic contract binding to access the raw methods on
}

// AccountantCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccountantCallerRaw struct {
	Contract *AccountantCaller // Generic read-only contract binding to access the raw methods on
}

// AccountantTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccountantTransactorRaw struct {
	Contract *AccountantTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccountant creates a new instance of Accountant, bound to a specific deployed contract.
func NewAccountant(address common.Address, backend bind.ContractBackend) (*Accountant, error) {
	contract, err := bindAccountant(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Accountant{AccountantCaller: AccountantCaller{contract: contract}, AccountantTransactor: AccountantTransactor{contract: contract}, AccountantFilterer: AccountantFilterer{contract: contract}}, nil
}

// NewAccountantCaller creates a new read-only instance of Accountant, bound to a specific deployed contract.
func NewAccountantCaller(address common.Address, caller bind.ContractCaller) (*AccountantCaller, error) {
	contract, err := bindAccountant(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccountantCaller{contract: contract}, nil
}

// NewAccountantTransactor creates a new write-only instance of Accountant, bound to a specific deployed contract.
func NewAccountantTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountantTransactor, error) {
	contract, err := bindAccountant(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccountantTransactor{contract: contract}, nil
}

// NewAccountantFilterer creates a new log filterer instance of Accountant, bound to a specific deployed contract.
func NewAccountantFilterer(address common.Address, filterer bind.ContractFilterer) (*AccountantFilterer, error) {
	contract, err := bindAccountant(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccountantFilterer{contract: contract}, nil
}

// bindAccountant binds a generic wrapper to an already deployed contract.
func bindAccountant(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountantABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Accountant *AccountantRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Accountant.Contract.AccountantCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Accountant *AccountantRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accountant.Contract.AccountantTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Accountant *AccountantRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Accountant.Contract.AccountantTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Accountant *AccountantCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Accountant.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Accountant *AccountantTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accountant.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Accountant *AccountantTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Accountant.Contract.contract.Transact(opts, method, params...)
}

// CustomConfig is a free data retrieval call binding the contract method 0x95ba6e16.
//
// Solidity: function customConfig(address , address ) view returns(uint16 managementFee, uint16 performanceFee, uint16 refundRatio, uint16 maxFee, uint16 maxGain, uint16 maxLoss)
func (_Accountant *AccountantCaller) CustomConfig(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (struct {
	ManagementFee  uint16
	PerformanceFee uint16
	RefundRatio    uint16
	MaxFee         uint16
	MaxGain        uint16
	MaxLoss        uint16
}, error) {
	var out []interface{}
	err := _Accountant.contract.Call(opts, &out, "customConfig", arg0, arg1)

	outstruct := new(struct {
		ManagementFee  uint16
		PerformanceFee uint16
		RefundRatio    uint16
		MaxFee         uint16
		MaxGain        uint16
		MaxLoss        uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ManagementFee = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.PerformanceFee = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.RefundRatio = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.MaxFee = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.MaxGain = *abi.ConvertType(out[4], new(uint16)).(*uint16)
	outstruct.MaxLoss = *abi.ConvertType(out[5], new(uint16)).(*uint16)

	return *outstruct, err

}

// CustomConfig is a free data retrieval call binding the contract method 0x95ba6e16.
//
// Solidity: function customConfig(address , address ) view returns(uint16 managementFee, uint16 performanceFee, uint16 refundRatio, uint16 maxFee, uint16 maxGain, uint16 maxLoss)
func (_Accountant *AccountantSession) CustomConfig(arg0 common.Address, arg1 common.Address) (struct {
	ManagementFee  uint16
	PerformanceFee uint16
	RefundRatio    uint16
	MaxFee         uint16
	MaxGain        uint16
	MaxLoss        uint16
}, error) {
	return _Accountant.Contract.CustomConfig(&_Accountant.CallOpts, arg0, arg1)
}

// CustomConfig is a free data retrieval call binding the contract method 0x95ba6e16.
//
// Solidity: function customConfig(address , address ) view returns(uint16 managementFee, uint16 performanceFee, uint16 refundRatio, uint16 maxFee, uint16 maxGain, uint16 maxLoss)
func (_Accountant *AccountantCallerSession) CustomConfig(arg0 common.Address, arg1 common.Address) (struct {
	ManagementFee  uint16
	PerformanceFee uint16
	RefundRatio    uint16
	MaxFee         uint16
	MaxGain        uint16
	MaxLoss        uint16
}, error) {
	return _Accountant.Contract.CustomConfig(&_Accountant.CallOpts, arg0, arg1)
}

// DefaultConfig is a free data retrieval call binding the contract method 0x9e09ed5f.
//
// Solidity: function defaultConfig() view returns(uint16 managementFee, uint16 performanceFee, uint16 refundRatio, uint16 maxFee, uint16 maxGain, uint16 maxLoss)
func (_Accountant *AccountantCaller) DefaultConfig(opts *bind.CallOpts) (struct {
	ManagementFee  uint16
	PerformanceFee uint16
	RefundRatio    uint16
	MaxFee         uint16
	MaxGain        uint16
	MaxLoss        uint16
}, error) {
	var out []interface{}
	err := _Accountant.contract.Call(opts, &out, "defaultConfig")

	outstruct := new(struct {
		ManagementFee  uint16
		PerformanceFee uint16
		RefundRatio    uint16
		MaxFee         uint16
		MaxGain        uint16
		MaxLoss        uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ManagementFee = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.PerformanceFee = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.RefundRatio = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.MaxFee = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.MaxGain = *abi.ConvertType(out[4], new(uint16)).(*uint16)
	outstruct.MaxLoss = *abi.ConvertType(out[5], new(uint16)).(*uint16)

	return *outstruct, err

}

// DefaultConfig is a free data retrieval call binding the contract method 0x9e09ed5f.
//
// Solidity: function defaultConfig() view returns(uint16 managementFee, uint16 performanceFee, uint16 refundRatio, uint16 maxFee, uint16 maxGain, uint16 maxLoss)
func (_Accountant *AccountantSession) DefaultConfig() (struct {
	ManagementFee  uint16
	PerformanceFee uint16
	RefundRatio    uint16
	MaxFee         uint16
	MaxGain        uint16
	MaxLoss        uint16
}, error) {
	return _Accountant.Contract.DefaultConfig(&_Accountant.CallOpts)
}

// DefaultConfig is a free data retrieval call binding the contract method 0x9e09ed5f.
//
// Solidity: function defaultConfig() view returns(uint16 managementFee, uint16 performanceFee, uint16 refundRatio, uint16 maxFee, uint16 maxGain, uint16 maxLoss)
func (_Accountant *AccountantCallerSession) DefaultConfig() (struct {
	ManagementFee  uint16
	PerformanceFee uint16
	RefundRatio    uint16
	MaxFee         uint16
	MaxGain        uint16
	MaxLoss        uint16
}, error) {
	return _Accountant.Contract.DefaultConfig(&_Accountant.CallOpts)
}

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() view returns(address)
func (_Accountant *AccountantCaller) FeeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Accountant.contract.Call(opts, &out, "feeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() view returns(address)
func (_Accountant *AccountantSession) FeeManager() (common.Address, error) {
	return _Accountant.Contract.FeeManager(&_Accountant.CallOpts)
}

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() view returns(address)
func (_Accountant *AccountantCallerSession) FeeManager() (common.Address, error) {
	return _Accountant.Contract.FeeManager(&_Accountant.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Accountant *AccountantCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Accountant.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Accountant *AccountantSession) FeeRecipient() (common.Address, error) {
	return _Accountant.Contract.FeeRecipient(&_Accountant.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Accountant *AccountantCallerSession) FeeRecipient() (common.Address, error) {
	return _Accountant.Contract.FeeRecipient(&_Accountant.CallOpts)
}

// FutureFeeManager is a free data retrieval call binding the contract method 0x015cf150.
//
// Solidity: function futureFeeManager() view returns(address)
func (_Accountant *AccountantCaller) FutureFeeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Accountant.contract.Call(opts, &out, "futureFeeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureFeeManager is a free data retrieval call binding the contract method 0x015cf150.
//
// Solidity: function futureFeeManager() view returns(address)
func (_Accountant *AccountantSession) FutureFeeManager() (common.Address, error) {
	return _Accountant.Contract.FutureFeeManager(&_Accountant.CallOpts)
}

// FutureFeeManager is a free data retrieval call binding the contract method 0x015cf150.
//
// Solidity: function futureFeeManager() view returns(address)
func (_Accountant *AccountantCallerSession) FutureFeeManager() (common.Address, error) {
	return _Accountant.Contract.FutureFeeManager(&_Accountant.CallOpts)
}

// ManagementFeeThreshold is a free data retrieval call binding the contract method 0x42e313e7.
//
// Solidity: function managementFeeThreshold() pure returns(uint16)
func (_Accountant *AccountantCaller) ManagementFeeThreshold(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Accountant.contract.Call(opts, &out, "managementFeeThreshold")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ManagementFeeThreshold is a free data retrieval call binding the contract method 0x42e313e7.
//
// Solidity: function managementFeeThreshold() pure returns(uint16)
func (_Accountant *AccountantSession) ManagementFeeThreshold() (uint16, error) {
	return _Accountant.Contract.ManagementFeeThreshold(&_Accountant.CallOpts)
}

// ManagementFeeThreshold is a free data retrieval call binding the contract method 0x42e313e7.
//
// Solidity: function managementFeeThreshold() pure returns(uint16)
func (_Accountant *AccountantCallerSession) ManagementFeeThreshold() (uint16, error) {
	return _Accountant.Contract.ManagementFeeThreshold(&_Accountant.CallOpts)
}

// MaxLoss is a free data retrieval call binding the contract method 0x5783fe39.
//
// Solidity: function maxLoss() view returns(uint256)
func (_Accountant *AccountantCaller) MaxLoss(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Accountant.contract.Call(opts, &out, "maxLoss")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxLoss is a free data retrieval call binding the contract method 0x5783fe39.
//
// Solidity: function maxLoss() view returns(uint256)
func (_Accountant *AccountantSession) MaxLoss() (*big.Int, error) {
	return _Accountant.Contract.MaxLoss(&_Accountant.CallOpts)
}

// MaxLoss is a free data retrieval call binding the contract method 0x5783fe39.
//
// Solidity: function maxLoss() view returns(uint256)
func (_Accountant *AccountantCallerSession) MaxLoss() (*big.Int, error) {
	return _Accountant.Contract.MaxLoss(&_Accountant.CallOpts)
}

// PerformanceFeeThreshold is a free data retrieval call binding the contract method 0x4e9e5f14.
//
// Solidity: function performanceFeeThreshold() pure returns(uint16)
func (_Accountant *AccountantCaller) PerformanceFeeThreshold(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Accountant.contract.Call(opts, &out, "performanceFeeThreshold")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// PerformanceFeeThreshold is a free data retrieval call binding the contract method 0x4e9e5f14.
//
// Solidity: function performanceFeeThreshold() pure returns(uint16)
func (_Accountant *AccountantSession) PerformanceFeeThreshold() (uint16, error) {
	return _Accountant.Contract.PerformanceFeeThreshold(&_Accountant.CallOpts)
}

// PerformanceFeeThreshold is a free data retrieval call binding the contract method 0x4e9e5f14.
//
// Solidity: function performanceFeeThreshold() pure returns(uint16)
func (_Accountant *AccountantCallerSession) PerformanceFeeThreshold() (uint16, error) {
	return _Accountant.Contract.PerformanceFeeThreshold(&_Accountant.CallOpts)
}

// Refund is a free data retrieval call binding the contract method 0x5e30b8a6.
//
// Solidity: function refund(address , address ) view returns(bool refund, uint248 amount)
func (_Accountant *AccountantCaller) Refund(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (struct {
	Refund bool
	Amount *big.Int
}, error) {
	var out []interface{}
	err := _Accountant.contract.Call(opts, &out, "refund", arg0, arg1)

	outstruct := new(struct {
		Refund bool
		Amount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Refund = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Refund is a free data retrieval call binding the contract method 0x5e30b8a6.
//
// Solidity: function refund(address , address ) view returns(bool refund, uint248 amount)
func (_Accountant *AccountantSession) Refund(arg0 common.Address, arg1 common.Address) (struct {
	Refund bool
	Amount *big.Int
}, error) {
	return _Accountant.Contract.Refund(&_Accountant.CallOpts, arg0, arg1)
}

// Refund is a free data retrieval call binding the contract method 0x5e30b8a6.
//
// Solidity: function refund(address , address ) view returns(bool refund, uint248 amount)
func (_Accountant *AccountantCallerSession) Refund(arg0 common.Address, arg1 common.Address) (struct {
	Refund bool
	Amount *big.Int
}, error) {
	return _Accountant.Contract.Refund(&_Accountant.CallOpts, arg0, arg1)
}

// UseCustomConfig is a free data retrieval call binding the contract method 0x00d67cc4.
//
// Solidity: function useCustomConfig(address vault, address strategy) view returns(bool)
func (_Accountant *AccountantCaller) UseCustomConfig(opts *bind.CallOpts, vault common.Address, strategy common.Address) (bool, error) {
	var out []interface{}
	err := _Accountant.contract.Call(opts, &out, "useCustomConfig", vault, strategy)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UseCustomConfig is a free data retrieval call binding the contract method 0x00d67cc4.
//
// Solidity: function useCustomConfig(address vault, address strategy) view returns(bool)
func (_Accountant *AccountantSession) UseCustomConfig(vault common.Address, strategy common.Address) (bool, error) {
	return _Accountant.Contract.UseCustomConfig(&_Accountant.CallOpts, vault, strategy)
}

// UseCustomConfig is a free data retrieval call binding the contract method 0x00d67cc4.
//
// Solidity: function useCustomConfig(address vault, address strategy) view returns(bool)
func (_Accountant *AccountantCallerSession) UseCustomConfig(vault common.Address, strategy common.Address) (bool, error) {
	return _Accountant.Contract.UseCustomConfig(&_Accountant.CallOpts, vault, strategy)
}

// VaultManager is a free data retrieval call binding the contract method 0x8a4adf24.
//
// Solidity: function vaultManager() view returns(address)
func (_Accountant *AccountantCaller) VaultManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Accountant.contract.Call(opts, &out, "vaultManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultManager is a free data retrieval call binding the contract method 0x8a4adf24.
//
// Solidity: function vaultManager() view returns(address)
func (_Accountant *AccountantSession) VaultManager() (common.Address, error) {
	return _Accountant.Contract.VaultManager(&_Accountant.CallOpts)
}

// VaultManager is a free data retrieval call binding the contract method 0x8a4adf24.
//
// Solidity: function vaultManager() view returns(address)
func (_Accountant *AccountantCallerSession) VaultManager() (common.Address, error) {
	return _Accountant.Contract.VaultManager(&_Accountant.CallOpts)
}

// Vaults is a free data retrieval call binding the contract method 0xa622ee7c.
//
// Solidity: function vaults(address ) view returns(bool)
func (_Accountant *AccountantCaller) Vaults(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Accountant.contract.Call(opts, &out, "vaults", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Vaults is a free data retrieval call binding the contract method 0xa622ee7c.
//
// Solidity: function vaults(address ) view returns(bool)
func (_Accountant *AccountantSession) Vaults(arg0 common.Address) (bool, error) {
	return _Accountant.Contract.Vaults(&_Accountant.CallOpts, arg0)
}

// Vaults is a free data retrieval call binding the contract method 0xa622ee7c.
//
// Solidity: function vaults(address ) view returns(bool)
func (_Accountant *AccountantCallerSession) Vaults(arg0 common.Address) (bool, error) {
	return _Accountant.Contract.Vaults(&_Accountant.CallOpts, arg0)
}

// AcceptFeeManager is a paid mutator transaction binding the contract method 0xf94c53c7.
//
// Solidity: function acceptFeeManager() returns()
func (_Accountant *AccountantTransactor) AcceptFeeManager(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accountant.contract.Transact(opts, "acceptFeeManager")
}

// AcceptFeeManager is a paid mutator transaction binding the contract method 0xf94c53c7.
//
// Solidity: function acceptFeeManager() returns()
func (_Accountant *AccountantSession) AcceptFeeManager() (*types.Transaction, error) {
	return _Accountant.Contract.AcceptFeeManager(&_Accountant.TransactOpts)
}

// AcceptFeeManager is a paid mutator transaction binding the contract method 0xf94c53c7.
//
// Solidity: function acceptFeeManager() returns()
func (_Accountant *AccountantTransactorSession) AcceptFeeManager() (*types.Transaction, error) {
	return _Accountant.Contract.AcceptFeeManager(&_Accountant.TransactOpts)
}

// AddVault is a paid mutator transaction binding the contract method 0x256b5a02.
//
// Solidity: function addVault(address vault) returns()
func (_Accountant *AccountantTransactor) AddVault(opts *bind.TransactOpts, vault common.Address) (*types.Transaction, error) {
	return _Accountant.contract.Transact(opts, "addVault", vault)
}

// AddVault is a paid mutator transaction binding the contract method 0x256b5a02.
//
// Solidity: function addVault(address vault) returns()
func (_Accountant *AccountantSession) AddVault(vault common.Address) (*types.Transaction, error) {
	return _Accountant.Contract.AddVault(&_Accountant.TransactOpts, vault)
}

// AddVault is a paid mutator transaction binding the contract method 0x256b5a02.
//
// Solidity: function addVault(address vault) returns()
func (_Accountant *AccountantTransactorSession) AddVault(vault common.Address) (*types.Transaction, error) {
	return _Accountant.Contract.AddVault(&_Accountant.TransactOpts, vault)
}

// Distribute is a paid mutator transaction binding the contract method 0x63453ae1.
//
// Solidity: function distribute(address token) returns()
func (_Accountant *AccountantTransactor) Distribute(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Accountant.contract.Transact(opts, "distribute", token)
}

// Distribute is a paid mutator transaction binding the contract method 0x63453ae1.
//
// Solidity: function distribute(address token) returns()
func (_Accountant *AccountantSession) Distribute(token common.Address) (*types.Transaction, error) {
	return _Accountant.Contract.Distribute(&_Accountant.TransactOpts, token)
}

// Distribute is a paid mutator transaction binding the contract method 0x63453ae1.
//
// Solidity: function distribute(address token) returns()
func (_Accountant *AccountantTransactorSession) Distribute(token common.Address) (*types.Transaction, error) {
	return _Accountant.Contract.Distribute(&_Accountant.TransactOpts, token)
}

// Distribute0 is a paid mutator transaction binding the contract method 0xfb932108.
//
// Solidity: function distribute(address token, uint256 amount) returns()
func (_Accountant *AccountantTransactor) Distribute0(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Accountant.contract.Transact(opts, "distribute0", token, amount)
}

// Distribute0 is a paid mutator transaction binding the contract method 0xfb932108.
//
// Solidity: function distribute(address token, uint256 amount) returns()
func (_Accountant *AccountantSession) Distribute0(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Accountant.Contract.Distribute0(&_Accountant.TransactOpts, token, amount)
}

// Distribute0 is a paid mutator transaction binding the contract method 0xfb932108.
//
// Solidity: function distribute(address token, uint256 amount) returns()
func (_Accountant *AccountantTransactorSession) Distribute0(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Accountant.Contract.Distribute0(&_Accountant.TransactOpts, token, amount)
}

// RemoveCustomConfig is a paid mutator transaction binding the contract method 0x006a650b.
//
// Solidity: function removeCustomConfig(address vault, address strategy) returns()
func (_Accountant *AccountantTransactor) RemoveCustomConfig(opts *bind.TransactOpts, vault common.Address, strategy common.Address) (*types.Transaction, error) {
	return _Accountant.contract.Transact(opts, "removeCustomConfig", vault, strategy)
}

// RemoveCustomConfig is a paid mutator transaction binding the contract method 0x006a650b.
//
// Solidity: function removeCustomConfig(address vault, address strategy) returns()
func (_Accountant *AccountantSession) RemoveCustomConfig(vault common.Address, strategy common.Address) (*types.Transaction, error) {
	return _Accountant.Contract.RemoveCustomConfig(&_Accountant.TransactOpts, vault, strategy)
}

// RemoveCustomConfig is a paid mutator transaction binding the contract method 0x006a650b.
//
// Solidity: function removeCustomConfig(address vault, address strategy) returns()
func (_Accountant *AccountantTransactorSession) RemoveCustomConfig(vault common.Address, strategy common.Address) (*types.Transaction, error) {
	return _Accountant.Contract.RemoveCustomConfig(&_Accountant.TransactOpts, vault, strategy)
}

// RemoveVault is a paid mutator transaction binding the contract method 0xceb68c23.
//
// Solidity: function removeVault(address vault) returns()
func (_Accountant *AccountantTransactor) RemoveVault(opts *bind.TransactOpts, vault common.Address) (*types.Transaction, error) {
	return _Accountant.contract.Transact(opts, "removeVault", vault)
}

// RemoveVault is a paid mutator transaction binding the contract method 0xceb68c23.
//
// Solidity: function removeVault(address vault) returns()
func (_Accountant *AccountantSession) RemoveVault(vault common.Address) (*types.Transaction, error) {
	return _Accountant.Contract.RemoveVault(&_Accountant.TransactOpts, vault)
}

// RemoveVault is a paid mutator transaction binding the contract method 0xceb68c23.
//
// Solidity: function removeVault(address vault) returns()
func (_Accountant *AccountantTransactorSession) RemoveVault(vault common.Address) (*types.Transaction, error) {
	return _Accountant.Contract.RemoveVault(&_Accountant.TransactOpts, vault)
}

// Report is a paid mutator transaction binding the contract method 0x921f8a8f.
//
// Solidity: function report(address strategy, uint256 gain, uint256 loss) returns(uint256 totalFees, uint256 totalRefunds)
func (_Accountant *AccountantTransactor) Report(opts *bind.TransactOpts, strategy common.Address, gain *big.Int, loss *big.Int) (*types.Transaction, error) {
	return _Accountant.contract.Transact(opts, "report", strategy, gain, loss)
}

// Report is a paid mutator transaction binding the contract method 0x921f8a8f.
//
// Solidity: function report(address strategy, uint256 gain, uint256 loss) returns(uint256 totalFees, uint256 totalRefunds)
func (_Accountant *AccountantSession) Report(strategy common.Address, gain *big.Int, loss *big.Int) (*types.Transaction, error) {
	return _Accountant.Contract.Report(&_Accountant.TransactOpts, strategy, gain, loss)
}

// Report is a paid mutator transaction binding the contract method 0x921f8a8f.
//
// Solidity: function report(address strategy, uint256 gain, uint256 loss) returns(uint256 totalFees, uint256 totalRefunds)
func (_Accountant *AccountantTransactorSession) Report(strategy common.Address, gain *big.Int, loss *big.Int) (*types.Transaction, error) {
	return _Accountant.Contract.Report(&_Accountant.TransactOpts, strategy, gain, loss)
}

// SetCustomConfig is a paid mutator transaction binding the contract method 0x8274dbe8.
//
// Solidity: function setCustomConfig(address vault, address strategy, uint16 customManagement, uint16 customPerformance, uint16 customRefund, uint16 customMaxFee, uint16 customMaxGain, uint16 customMaxLoss) returns()
func (_Accountant *AccountantTransactor) SetCustomConfig(opts *bind.TransactOpts, vault common.Address, strategy common.Address, customManagement uint16, customPerformance uint16, customRefund uint16, customMaxFee uint16, customMaxGain uint16, customMaxLoss uint16) (*types.Transaction, error) {
	return _Accountant.contract.Transact(opts, "setCustomConfig", vault, strategy, customManagement, customPerformance, customRefund, customMaxFee, customMaxGain, customMaxLoss)
}

// SetCustomConfig is a paid mutator transaction binding the contract method 0x8274dbe8.
//
// Solidity: function setCustomConfig(address vault, address strategy, uint16 customManagement, uint16 customPerformance, uint16 customRefund, uint16 customMaxFee, uint16 customMaxGain, uint16 customMaxLoss) returns()
func (_Accountant *AccountantSession) SetCustomConfig(vault common.Address, strategy common.Address, customManagement uint16, customPerformance uint16, customRefund uint16, customMaxFee uint16, customMaxGain uint16, customMaxLoss uint16) (*types.Transaction, error) {
	return _Accountant.Contract.SetCustomConfig(&_Accountant.TransactOpts, vault, strategy, customManagement, customPerformance, customRefund, customMaxFee, customMaxGain, customMaxLoss)
}

// SetCustomConfig is a paid mutator transaction binding the contract method 0x8274dbe8.
//
// Solidity: function setCustomConfig(address vault, address strategy, uint16 customManagement, uint16 customPerformance, uint16 customRefund, uint16 customMaxFee, uint16 customMaxGain, uint16 customMaxLoss) returns()
func (_Accountant *AccountantTransactorSession) SetCustomConfig(vault common.Address, strategy common.Address, customManagement uint16, customPerformance uint16, customRefund uint16, customMaxFee uint16, customMaxGain uint16, customMaxLoss uint16) (*types.Transaction, error) {
	return _Accountant.Contract.SetCustomConfig(&_Accountant.TransactOpts, vault, strategy, customManagement, customPerformance, customRefund, customMaxFee, customMaxGain, customMaxLoss)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address newFeeRecipient) returns()
func (_Accountant *AccountantTransactor) SetFeeRecipient(opts *bind.TransactOpts, newFeeRecipient common.Address) (*types.Transaction, error) {
	return _Accountant.contract.Transact(opts, "setFeeRecipient", newFeeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address newFeeRecipient) returns()
func (_Accountant *AccountantSession) SetFeeRecipient(newFeeRecipient common.Address) (*types.Transaction, error) {
	return _Accountant.Contract.SetFeeRecipient(&_Accountant.TransactOpts, newFeeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address newFeeRecipient) returns()
func (_Accountant *AccountantTransactorSession) SetFeeRecipient(newFeeRecipient common.Address) (*types.Transaction, error) {
	return _Accountant.Contract.SetFeeRecipient(&_Accountant.TransactOpts, newFeeRecipient)
}

// SetFutureFeeManager is a paid mutator transaction binding the contract method 0x9b3b6955.
//
// Solidity: function setFutureFeeManager(address _futureFeeManager) returns()
func (_Accountant *AccountantTransactor) SetFutureFeeManager(opts *bind.TransactOpts, _futureFeeManager common.Address) (*types.Transaction, error) {
	return _Accountant.contract.Transact(opts, "setFutureFeeManager", _futureFeeManager)
}

// SetFutureFeeManager is a paid mutator transaction binding the contract method 0x9b3b6955.
//
// Solidity: function setFutureFeeManager(address _futureFeeManager) returns()
func (_Accountant *AccountantSession) SetFutureFeeManager(_futureFeeManager common.Address) (*types.Transaction, error) {
	return _Accountant.Contract.SetFutureFeeManager(&_Accountant.TransactOpts, _futureFeeManager)
}

// SetFutureFeeManager is a paid mutator transaction binding the contract method 0x9b3b6955.
//
// Solidity: function setFutureFeeManager(address _futureFeeManager) returns()
func (_Accountant *AccountantTransactorSession) SetFutureFeeManager(_futureFeeManager common.Address) (*types.Transaction, error) {
	return _Accountant.Contract.SetFutureFeeManager(&_Accountant.TransactOpts, _futureFeeManager)
}

// SetMaxLoss is a paid mutator transaction binding the contract method 0x24be6628.
//
// Solidity: function setMaxLoss(uint256 _maxLoss) returns()
func (_Accountant *AccountantTransactor) SetMaxLoss(opts *bind.TransactOpts, _maxLoss *big.Int) (*types.Transaction, error) {
	return _Accountant.contract.Transact(opts, "setMaxLoss", _maxLoss)
}

// SetMaxLoss is a paid mutator transaction binding the contract method 0x24be6628.
//
// Solidity: function setMaxLoss(uint256 _maxLoss) returns()
func (_Accountant *AccountantSession) SetMaxLoss(_maxLoss *big.Int) (*types.Transaction, error) {
	return _Accountant.Contract.SetMaxLoss(&_Accountant.TransactOpts, _maxLoss)
}

// SetMaxLoss is a paid mutator transaction binding the contract method 0x24be6628.
//
// Solidity: function setMaxLoss(uint256 _maxLoss) returns()
func (_Accountant *AccountantTransactorSession) SetMaxLoss(_maxLoss *big.Int) (*types.Transaction, error) {
	return _Accountant.Contract.SetMaxLoss(&_Accountant.TransactOpts, _maxLoss)
}

// SetRefund is a paid mutator transaction binding the contract method 0xd348f47c.
//
// Solidity: function setRefund(address _vault, address _strategy, bool _refund, uint256 _amount) returns()
func (_Accountant *AccountantTransactor) SetRefund(opts *bind.TransactOpts, _vault common.Address, _strategy common.Address, _refund bool, _amount *big.Int) (*types.Transaction, error) {
	return _Accountant.contract.Transact(opts, "setRefund", _vault, _strategy, _refund, _amount)
}

// SetRefund is a paid mutator transaction binding the contract method 0xd348f47c.
//
// Solidity: function setRefund(address _vault, address _strategy, bool _refund, uint256 _amount) returns()
func (_Accountant *AccountantSession) SetRefund(_vault common.Address, _strategy common.Address, _refund bool, _amount *big.Int) (*types.Transaction, error) {
	return _Accountant.Contract.SetRefund(&_Accountant.TransactOpts, _vault, _strategy, _refund, _amount)
}

// SetRefund is a paid mutator transaction binding the contract method 0xd348f47c.
//
// Solidity: function setRefund(address _vault, address _strategy, bool _refund, uint256 _amount) returns()
func (_Accountant *AccountantTransactorSession) SetRefund(_vault common.Address, _strategy common.Address, _refund bool, _amount *big.Int) (*types.Transaction, error) {
	return _Accountant.Contract.SetRefund(&_Accountant.TransactOpts, _vault, _strategy, _refund, _amount)
}

// SetVaultManager is a paid mutator transaction binding the contract method 0xb543503e.
//
// Solidity: function setVaultManager(address newVaultManager) returns()
func (_Accountant *AccountantTransactor) SetVaultManager(opts *bind.TransactOpts, newVaultManager common.Address) (*types.Transaction, error) {
	return _Accountant.contract.Transact(opts, "setVaultManager", newVaultManager)
}

// SetVaultManager is a paid mutator transaction binding the contract method 0xb543503e.
//
// Solidity: function setVaultManager(address newVaultManager) returns()
func (_Accountant *AccountantSession) SetVaultManager(newVaultManager common.Address) (*types.Transaction, error) {
	return _Accountant.Contract.SetVaultManager(&_Accountant.TransactOpts, newVaultManager)
}

// SetVaultManager is a paid mutator transaction binding the contract method 0xb543503e.
//
// Solidity: function setVaultManager(address newVaultManager) returns()
func (_Accountant *AccountantTransactorSession) SetVaultManager(newVaultManager common.Address) (*types.Transaction, error) {
	return _Accountant.Contract.SetVaultManager(&_Accountant.TransactOpts, newVaultManager)
}

// UpdateDefaultConfig is a paid mutator transaction binding the contract method 0xe2a85ce4.
//
// Solidity: function updateDefaultConfig(uint16 defaultManagement, uint16 defaultPerformance, uint16 defaultRefund, uint16 defaultMaxFee, uint16 defaultMaxGain, uint16 defaultMaxLoss) returns()
func (_Accountant *AccountantTransactor) UpdateDefaultConfig(opts *bind.TransactOpts, defaultManagement uint16, defaultPerformance uint16, defaultRefund uint16, defaultMaxFee uint16, defaultMaxGain uint16, defaultMaxLoss uint16) (*types.Transaction, error) {
	return _Accountant.contract.Transact(opts, "updateDefaultConfig", defaultManagement, defaultPerformance, defaultRefund, defaultMaxFee, defaultMaxGain, defaultMaxLoss)
}

// UpdateDefaultConfig is a paid mutator transaction binding the contract method 0xe2a85ce4.
//
// Solidity: function updateDefaultConfig(uint16 defaultManagement, uint16 defaultPerformance, uint16 defaultRefund, uint16 defaultMaxFee, uint16 defaultMaxGain, uint16 defaultMaxLoss) returns()
func (_Accountant *AccountantSession) UpdateDefaultConfig(defaultManagement uint16, defaultPerformance uint16, defaultRefund uint16, defaultMaxFee uint16, defaultMaxGain uint16, defaultMaxLoss uint16) (*types.Transaction, error) {
	return _Accountant.Contract.UpdateDefaultConfig(&_Accountant.TransactOpts, defaultManagement, defaultPerformance, defaultRefund, defaultMaxFee, defaultMaxGain, defaultMaxLoss)
}

// UpdateDefaultConfig is a paid mutator transaction binding the contract method 0xe2a85ce4.
//
// Solidity: function updateDefaultConfig(uint16 defaultManagement, uint16 defaultPerformance, uint16 defaultRefund, uint16 defaultMaxFee, uint16 defaultMaxGain, uint16 defaultMaxLoss) returns()
func (_Accountant *AccountantTransactorSession) UpdateDefaultConfig(defaultManagement uint16, defaultPerformance uint16, defaultRefund uint16, defaultMaxFee uint16, defaultMaxGain uint16, defaultMaxLoss uint16) (*types.Transaction, error) {
	return _Accountant.Contract.UpdateDefaultConfig(&_Accountant.TransactOpts, defaultManagement, defaultPerformance, defaultRefund, defaultMaxFee, defaultMaxGain, defaultMaxLoss)
}

// WithdrawUnderlying is a paid mutator transaction binding the contract method 0xfdb87252.
//
// Solidity: function withdrawUnderlying(address vault, uint256 amount) returns()
func (_Accountant *AccountantTransactor) WithdrawUnderlying(opts *bind.TransactOpts, vault common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Accountant.contract.Transact(opts, "withdrawUnderlying", vault, amount)
}

// WithdrawUnderlying is a paid mutator transaction binding the contract method 0xfdb87252.
//
// Solidity: function withdrawUnderlying(address vault, uint256 amount) returns()
func (_Accountant *AccountantSession) WithdrawUnderlying(vault common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Accountant.Contract.WithdrawUnderlying(&_Accountant.TransactOpts, vault, amount)
}

// WithdrawUnderlying is a paid mutator transaction binding the contract method 0xfdb87252.
//
// Solidity: function withdrawUnderlying(address vault, uint256 amount) returns()
func (_Accountant *AccountantTransactorSession) WithdrawUnderlying(vault common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Accountant.Contract.WithdrawUnderlying(&_Accountant.TransactOpts, vault, amount)
}

// AccountantDistributeRewardsIterator is returned from FilterDistributeRewards and is used to iterate over the raw logs and unpacked data for DistributeRewards events raised by the Accountant contract.
type AccountantDistributeRewardsIterator struct {
	Event *AccountantDistributeRewards // Event containing the contract specifics and raw log

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
func (it *AccountantDistributeRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountantDistributeRewards)
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
		it.Event = new(AccountantDistributeRewards)
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
func (it *AccountantDistributeRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountantDistributeRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountantDistributeRewards represents a DistributeRewards event raised by the Accountant contract.
type AccountantDistributeRewards struct {
	Token   common.Address
	Rewards *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDistributeRewards is a free log retrieval operation binding the contract event 0x962bc326c7b063c70721a63687e0e19450155f93c58eca94769746c0cfb02c5d.
//
// Solidity: event DistributeRewards(address indexed token, uint256 rewards)
func (_Accountant *AccountantFilterer) FilterDistributeRewards(opts *bind.FilterOpts, token []common.Address) (*AccountantDistributeRewardsIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Accountant.contract.FilterLogs(opts, "DistributeRewards", tokenRule)
	if err != nil {
		return nil, err
	}
	return &AccountantDistributeRewardsIterator{contract: _Accountant.contract, event: "DistributeRewards", logs: logs, sub: sub}, nil
}

// WatchDistributeRewards is a free log subscription operation binding the contract event 0x962bc326c7b063c70721a63687e0e19450155f93c58eca94769746c0cfb02c5d.
//
// Solidity: event DistributeRewards(address indexed token, uint256 rewards)
func (_Accountant *AccountantFilterer) WatchDistributeRewards(opts *bind.WatchOpts, sink chan<- *AccountantDistributeRewards, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Accountant.contract.WatchLogs(opts, "DistributeRewards", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountantDistributeRewards)
				if err := _Accountant.contract.UnpackLog(event, "DistributeRewards", log); err != nil {
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

// ParseDistributeRewards is a log parse operation binding the contract event 0x962bc326c7b063c70721a63687e0e19450155f93c58eca94769746c0cfb02c5d.
//
// Solidity: event DistributeRewards(address indexed token, uint256 rewards)
func (_Accountant *AccountantFilterer) ParseDistributeRewards(log types.Log) (*AccountantDistributeRewards, error) {
	event := new(AccountantDistributeRewards)
	if err := _Accountant.contract.UnpackLog(event, "DistributeRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountantNewFeeManagerIterator is returned from FilterNewFeeManager and is used to iterate over the raw logs and unpacked data for NewFeeManager events raised by the Accountant contract.
type AccountantNewFeeManagerIterator struct {
	Event *AccountantNewFeeManager // Event containing the contract specifics and raw log

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
func (it *AccountantNewFeeManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountantNewFeeManager)
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
		it.Event = new(AccountantNewFeeManager)
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
func (it *AccountantNewFeeManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountantNewFeeManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountantNewFeeManager represents a NewFeeManager event raised by the Accountant contract.
type AccountantNewFeeManager struct {
	FeeManager common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewFeeManager is a free log retrieval operation binding the contract event 0x772ddcfc9a0f3b1401c0f60000a81999005d9d593b71bb67707c5f326eb7c94d.
//
// Solidity: event NewFeeManager(address indexed feeManager)
func (_Accountant *AccountantFilterer) FilterNewFeeManager(opts *bind.FilterOpts, feeManager []common.Address) (*AccountantNewFeeManagerIterator, error) {

	var feeManagerRule []interface{}
	for _, feeManagerItem := range feeManager {
		feeManagerRule = append(feeManagerRule, feeManagerItem)
	}

	logs, sub, err := _Accountant.contract.FilterLogs(opts, "NewFeeManager", feeManagerRule)
	if err != nil {
		return nil, err
	}
	return &AccountantNewFeeManagerIterator{contract: _Accountant.contract, event: "NewFeeManager", logs: logs, sub: sub}, nil
}

// WatchNewFeeManager is a free log subscription operation binding the contract event 0x772ddcfc9a0f3b1401c0f60000a81999005d9d593b71bb67707c5f326eb7c94d.
//
// Solidity: event NewFeeManager(address indexed feeManager)
func (_Accountant *AccountantFilterer) WatchNewFeeManager(opts *bind.WatchOpts, sink chan<- *AccountantNewFeeManager, feeManager []common.Address) (event.Subscription, error) {

	var feeManagerRule []interface{}
	for _, feeManagerItem := range feeManager {
		feeManagerRule = append(feeManagerRule, feeManagerItem)
	}

	logs, sub, err := _Accountant.contract.WatchLogs(opts, "NewFeeManager", feeManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountantNewFeeManager)
				if err := _Accountant.contract.UnpackLog(event, "NewFeeManager", log); err != nil {
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

// ParseNewFeeManager is a log parse operation binding the contract event 0x772ddcfc9a0f3b1401c0f60000a81999005d9d593b71bb67707c5f326eb7c94d.
//
// Solidity: event NewFeeManager(address indexed feeManager)
func (_Accountant *AccountantFilterer) ParseNewFeeManager(log types.Log) (*AccountantNewFeeManager, error) {
	event := new(AccountantNewFeeManager)
	if err := _Accountant.contract.UnpackLog(event, "NewFeeManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountantRemovedCustomFeeConfigIterator is returned from FilterRemovedCustomFeeConfig and is used to iterate over the raw logs and unpacked data for RemovedCustomFeeConfig events raised by the Accountant contract.
type AccountantRemovedCustomFeeConfigIterator struct {
	Event *AccountantRemovedCustomFeeConfig // Event containing the contract specifics and raw log

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
func (it *AccountantRemovedCustomFeeConfigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountantRemovedCustomFeeConfig)
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
		it.Event = new(AccountantRemovedCustomFeeConfig)
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
func (it *AccountantRemovedCustomFeeConfigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountantRemovedCustomFeeConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountantRemovedCustomFeeConfig represents a RemovedCustomFeeConfig event raised by the Accountant contract.
type AccountantRemovedCustomFeeConfig struct {
	Vault    common.Address
	Strategy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRemovedCustomFeeConfig is a free log retrieval operation binding the contract event 0xbd3097d5444217271f366657d3e2d4668a90ade03cfefa42f7106de06b89c2c2.
//
// Solidity: event RemovedCustomFeeConfig(address indexed vault, address indexed strategy)
func (_Accountant *AccountantFilterer) FilterRemovedCustomFeeConfig(opts *bind.FilterOpts, vault []common.Address, strategy []common.Address) (*AccountantRemovedCustomFeeConfigIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Accountant.contract.FilterLogs(opts, "RemovedCustomFeeConfig", vaultRule, strategyRule)
	if err != nil {
		return nil, err
	}
	return &AccountantRemovedCustomFeeConfigIterator{contract: _Accountant.contract, event: "RemovedCustomFeeConfig", logs: logs, sub: sub}, nil
}

// WatchRemovedCustomFeeConfig is a free log subscription operation binding the contract event 0xbd3097d5444217271f366657d3e2d4668a90ade03cfefa42f7106de06b89c2c2.
//
// Solidity: event RemovedCustomFeeConfig(address indexed vault, address indexed strategy)
func (_Accountant *AccountantFilterer) WatchRemovedCustomFeeConfig(opts *bind.WatchOpts, sink chan<- *AccountantRemovedCustomFeeConfig, vault []common.Address, strategy []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Accountant.contract.WatchLogs(opts, "RemovedCustomFeeConfig", vaultRule, strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountantRemovedCustomFeeConfig)
				if err := _Accountant.contract.UnpackLog(event, "RemovedCustomFeeConfig", log); err != nil {
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

// ParseRemovedCustomFeeConfig is a log parse operation binding the contract event 0xbd3097d5444217271f366657d3e2d4668a90ade03cfefa42f7106de06b89c2c2.
//
// Solidity: event RemovedCustomFeeConfig(address indexed vault, address indexed strategy)
func (_Accountant *AccountantFilterer) ParseRemovedCustomFeeConfig(log types.Log) (*AccountantRemovedCustomFeeConfig, error) {
	event := new(AccountantRemovedCustomFeeConfig)
	if err := _Accountant.contract.UnpackLog(event, "RemovedCustomFeeConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountantSetFutureFeeManagerIterator is returned from FilterSetFutureFeeManager and is used to iterate over the raw logs and unpacked data for SetFutureFeeManager events raised by the Accountant contract.
type AccountantSetFutureFeeManagerIterator struct {
	Event *AccountantSetFutureFeeManager // Event containing the contract specifics and raw log

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
func (it *AccountantSetFutureFeeManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountantSetFutureFeeManager)
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
		it.Event = new(AccountantSetFutureFeeManager)
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
func (it *AccountantSetFutureFeeManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountantSetFutureFeeManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountantSetFutureFeeManager represents a SetFutureFeeManager event raised by the Accountant contract.
type AccountantSetFutureFeeManager struct {
	FutureFeeManager common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetFutureFeeManager is a free log retrieval operation binding the contract event 0xa839c45565e8a86de41783841928f4acde049c2b7160f0ea4d4698220c5af61b.
//
// Solidity: event SetFutureFeeManager(address indexed futureFeeManager)
func (_Accountant *AccountantFilterer) FilterSetFutureFeeManager(opts *bind.FilterOpts, futureFeeManager []common.Address) (*AccountantSetFutureFeeManagerIterator, error) {

	var futureFeeManagerRule []interface{}
	for _, futureFeeManagerItem := range futureFeeManager {
		futureFeeManagerRule = append(futureFeeManagerRule, futureFeeManagerItem)
	}

	logs, sub, err := _Accountant.contract.FilterLogs(opts, "SetFutureFeeManager", futureFeeManagerRule)
	if err != nil {
		return nil, err
	}
	return &AccountantSetFutureFeeManagerIterator{contract: _Accountant.contract, event: "SetFutureFeeManager", logs: logs, sub: sub}, nil
}

// WatchSetFutureFeeManager is a free log subscription operation binding the contract event 0xa839c45565e8a86de41783841928f4acde049c2b7160f0ea4d4698220c5af61b.
//
// Solidity: event SetFutureFeeManager(address indexed futureFeeManager)
func (_Accountant *AccountantFilterer) WatchSetFutureFeeManager(opts *bind.WatchOpts, sink chan<- *AccountantSetFutureFeeManager, futureFeeManager []common.Address) (event.Subscription, error) {

	var futureFeeManagerRule []interface{}
	for _, futureFeeManagerItem := range futureFeeManager {
		futureFeeManagerRule = append(futureFeeManagerRule, futureFeeManagerItem)
	}

	logs, sub, err := _Accountant.contract.WatchLogs(opts, "SetFutureFeeManager", futureFeeManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountantSetFutureFeeManager)
				if err := _Accountant.contract.UnpackLog(event, "SetFutureFeeManager", log); err != nil {
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

// ParseSetFutureFeeManager is a log parse operation binding the contract event 0xa839c45565e8a86de41783841928f4acde049c2b7160f0ea4d4698220c5af61b.
//
// Solidity: event SetFutureFeeManager(address indexed futureFeeManager)
func (_Accountant *AccountantFilterer) ParseSetFutureFeeManager(log types.Log) (*AccountantSetFutureFeeManager, error) {
	event := new(AccountantSetFutureFeeManager)
	if err := _Accountant.contract.UnpackLog(event, "SetFutureFeeManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountantUpdateCustomFeeConfigIterator is returned from FilterUpdateCustomFeeConfig and is used to iterate over the raw logs and unpacked data for UpdateCustomFeeConfig events raised by the Accountant contract.
type AccountantUpdateCustomFeeConfigIterator struct {
	Event *AccountantUpdateCustomFeeConfig // Event containing the contract specifics and raw log

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
func (it *AccountantUpdateCustomFeeConfigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountantUpdateCustomFeeConfig)
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
		it.Event = new(AccountantUpdateCustomFeeConfig)
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
func (it *AccountantUpdateCustomFeeConfigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountantUpdateCustomFeeConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountantUpdateCustomFeeConfig represents a UpdateCustomFeeConfig event raised by the Accountant contract.
type AccountantUpdateCustomFeeConfig struct {
	Vault        common.Address
	Strategy     common.Address
	CustomConfig HealthCheckAccountantFee
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateCustomFeeConfig is a free log retrieval operation binding the contract event 0xdb9eae8109dc0288158d8a6bf610578cd67a530847095c1aeb641dfd1a4d1e4f.
//
// Solidity: event UpdateCustomFeeConfig(address indexed vault, address indexed strategy, (uint16,uint16,uint16,uint16,uint16,uint16) custom_config)
func (_Accountant *AccountantFilterer) FilterUpdateCustomFeeConfig(opts *bind.FilterOpts, vault []common.Address, strategy []common.Address) (*AccountantUpdateCustomFeeConfigIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Accountant.contract.FilterLogs(opts, "UpdateCustomFeeConfig", vaultRule, strategyRule)
	if err != nil {
		return nil, err
	}
	return &AccountantUpdateCustomFeeConfigIterator{contract: _Accountant.contract, event: "UpdateCustomFeeConfig", logs: logs, sub: sub}, nil
}

// WatchUpdateCustomFeeConfig is a free log subscription operation binding the contract event 0xdb9eae8109dc0288158d8a6bf610578cd67a530847095c1aeb641dfd1a4d1e4f.
//
// Solidity: event UpdateCustomFeeConfig(address indexed vault, address indexed strategy, (uint16,uint16,uint16,uint16,uint16,uint16) custom_config)
func (_Accountant *AccountantFilterer) WatchUpdateCustomFeeConfig(opts *bind.WatchOpts, sink chan<- *AccountantUpdateCustomFeeConfig, vault []common.Address, strategy []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Accountant.contract.WatchLogs(opts, "UpdateCustomFeeConfig", vaultRule, strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountantUpdateCustomFeeConfig)
				if err := _Accountant.contract.UnpackLog(event, "UpdateCustomFeeConfig", log); err != nil {
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

// ParseUpdateCustomFeeConfig is a log parse operation binding the contract event 0xdb9eae8109dc0288158d8a6bf610578cd67a530847095c1aeb641dfd1a4d1e4f.
//
// Solidity: event UpdateCustomFeeConfig(address indexed vault, address indexed strategy, (uint16,uint16,uint16,uint16,uint16,uint16) custom_config)
func (_Accountant *AccountantFilterer) ParseUpdateCustomFeeConfig(log types.Log) (*AccountantUpdateCustomFeeConfig, error) {
	event := new(AccountantUpdateCustomFeeConfig)
	if err := _Accountant.contract.UnpackLog(event, "UpdateCustomFeeConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountantUpdateDefaultFeeConfigIterator is returned from FilterUpdateDefaultFeeConfig and is used to iterate over the raw logs and unpacked data for UpdateDefaultFeeConfig events raised by the Accountant contract.
type AccountantUpdateDefaultFeeConfigIterator struct {
	Event *AccountantUpdateDefaultFeeConfig // Event containing the contract specifics and raw log

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
func (it *AccountantUpdateDefaultFeeConfigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountantUpdateDefaultFeeConfig)
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
		it.Event = new(AccountantUpdateDefaultFeeConfig)
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
func (it *AccountantUpdateDefaultFeeConfigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountantUpdateDefaultFeeConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountantUpdateDefaultFeeConfig represents a UpdateDefaultFeeConfig event raised by the Accountant contract.
type AccountantUpdateDefaultFeeConfig struct {
	DefaultFeeConfig HealthCheckAccountantFee
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUpdateDefaultFeeConfig is a free log retrieval operation binding the contract event 0xf3ff5bcf8ed3012f7176c6f818fefe02724b1948c98c25f8ff0339fc74c11b8c.
//
// Solidity: event UpdateDefaultFeeConfig((uint16,uint16,uint16,uint16,uint16,uint16) defaultFeeConfig)
func (_Accountant *AccountantFilterer) FilterUpdateDefaultFeeConfig(opts *bind.FilterOpts) (*AccountantUpdateDefaultFeeConfigIterator, error) {

	logs, sub, err := _Accountant.contract.FilterLogs(opts, "UpdateDefaultFeeConfig")
	if err != nil {
		return nil, err
	}
	return &AccountantUpdateDefaultFeeConfigIterator{contract: _Accountant.contract, event: "UpdateDefaultFeeConfig", logs: logs, sub: sub}, nil
}

// WatchUpdateDefaultFeeConfig is a free log subscription operation binding the contract event 0xf3ff5bcf8ed3012f7176c6f818fefe02724b1948c98c25f8ff0339fc74c11b8c.
//
// Solidity: event UpdateDefaultFeeConfig((uint16,uint16,uint16,uint16,uint16,uint16) defaultFeeConfig)
func (_Accountant *AccountantFilterer) WatchUpdateDefaultFeeConfig(opts *bind.WatchOpts, sink chan<- *AccountantUpdateDefaultFeeConfig) (event.Subscription, error) {

	logs, sub, err := _Accountant.contract.WatchLogs(opts, "UpdateDefaultFeeConfig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountantUpdateDefaultFeeConfig)
				if err := _Accountant.contract.UnpackLog(event, "UpdateDefaultFeeConfig", log); err != nil {
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

// ParseUpdateDefaultFeeConfig is a log parse operation binding the contract event 0xf3ff5bcf8ed3012f7176c6f818fefe02724b1948c98c25f8ff0339fc74c11b8c.
//
// Solidity: event UpdateDefaultFeeConfig((uint16,uint16,uint16,uint16,uint16,uint16) defaultFeeConfig)
func (_Accountant *AccountantFilterer) ParseUpdateDefaultFeeConfig(log types.Log) (*AccountantUpdateDefaultFeeConfig, error) {
	event := new(AccountantUpdateDefaultFeeConfig)
	if err := _Accountant.contract.UnpackLog(event, "UpdateDefaultFeeConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountantUpdateFeeRecipientIterator is returned from FilterUpdateFeeRecipient and is used to iterate over the raw logs and unpacked data for UpdateFeeRecipient events raised by the Accountant contract.
type AccountantUpdateFeeRecipientIterator struct {
	Event *AccountantUpdateFeeRecipient // Event containing the contract specifics and raw log

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
func (it *AccountantUpdateFeeRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountantUpdateFeeRecipient)
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
		it.Event = new(AccountantUpdateFeeRecipient)
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
func (it *AccountantUpdateFeeRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountantUpdateFeeRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountantUpdateFeeRecipient represents a UpdateFeeRecipient event raised by the Accountant contract.
type AccountantUpdateFeeRecipient struct {
	OldFeeRecipient common.Address
	NewFeeRecipient common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateFeeRecipient is a free log retrieval operation binding the contract event 0xb03f92c8c7ac39710f28b146f754650929499d599a66d51423cbd7f3ceb9aee3.
//
// Solidity: event UpdateFeeRecipient(address indexed oldFeeRecipient, address indexed newFeeRecipient)
func (_Accountant *AccountantFilterer) FilterUpdateFeeRecipient(opts *bind.FilterOpts, oldFeeRecipient []common.Address, newFeeRecipient []common.Address) (*AccountantUpdateFeeRecipientIterator, error) {

	var oldFeeRecipientRule []interface{}
	for _, oldFeeRecipientItem := range oldFeeRecipient {
		oldFeeRecipientRule = append(oldFeeRecipientRule, oldFeeRecipientItem)
	}
	var newFeeRecipientRule []interface{}
	for _, newFeeRecipientItem := range newFeeRecipient {
		newFeeRecipientRule = append(newFeeRecipientRule, newFeeRecipientItem)
	}

	logs, sub, err := _Accountant.contract.FilterLogs(opts, "UpdateFeeRecipient", oldFeeRecipientRule, newFeeRecipientRule)
	if err != nil {
		return nil, err
	}
	return &AccountantUpdateFeeRecipientIterator{contract: _Accountant.contract, event: "UpdateFeeRecipient", logs: logs, sub: sub}, nil
}

// WatchUpdateFeeRecipient is a free log subscription operation binding the contract event 0xb03f92c8c7ac39710f28b146f754650929499d599a66d51423cbd7f3ceb9aee3.
//
// Solidity: event UpdateFeeRecipient(address indexed oldFeeRecipient, address indexed newFeeRecipient)
func (_Accountant *AccountantFilterer) WatchUpdateFeeRecipient(opts *bind.WatchOpts, sink chan<- *AccountantUpdateFeeRecipient, oldFeeRecipient []common.Address, newFeeRecipient []common.Address) (event.Subscription, error) {

	var oldFeeRecipientRule []interface{}
	for _, oldFeeRecipientItem := range oldFeeRecipient {
		oldFeeRecipientRule = append(oldFeeRecipientRule, oldFeeRecipientItem)
	}
	var newFeeRecipientRule []interface{}
	for _, newFeeRecipientItem := range newFeeRecipient {
		newFeeRecipientRule = append(newFeeRecipientRule, newFeeRecipientItem)
	}

	logs, sub, err := _Accountant.contract.WatchLogs(opts, "UpdateFeeRecipient", oldFeeRecipientRule, newFeeRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountantUpdateFeeRecipient)
				if err := _Accountant.contract.UnpackLog(event, "UpdateFeeRecipient", log); err != nil {
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

// ParseUpdateFeeRecipient is a log parse operation binding the contract event 0xb03f92c8c7ac39710f28b146f754650929499d599a66d51423cbd7f3ceb9aee3.
//
// Solidity: event UpdateFeeRecipient(address indexed oldFeeRecipient, address indexed newFeeRecipient)
func (_Accountant *AccountantFilterer) ParseUpdateFeeRecipient(log types.Log) (*AccountantUpdateFeeRecipient, error) {
	event := new(AccountantUpdateFeeRecipient)
	if err := _Accountant.contract.UnpackLog(event, "UpdateFeeRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountantUpdateMaxLossIterator is returned from FilterUpdateMaxLoss and is used to iterate over the raw logs and unpacked data for UpdateMaxLoss events raised by the Accountant contract.
type AccountantUpdateMaxLossIterator struct {
	Event *AccountantUpdateMaxLoss // Event containing the contract specifics and raw log

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
func (it *AccountantUpdateMaxLossIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountantUpdateMaxLoss)
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
		it.Event = new(AccountantUpdateMaxLoss)
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
func (it *AccountantUpdateMaxLossIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountantUpdateMaxLossIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountantUpdateMaxLoss represents a UpdateMaxLoss event raised by the Accountant contract.
type AccountantUpdateMaxLoss struct {
	MaxLoss *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateMaxLoss is a free log retrieval operation binding the contract event 0x18182e268b61d2aada98f23ade23b0ea133d5b0b6712dbfa682dc6da29941c22.
//
// Solidity: event UpdateMaxLoss(uint256 maxLoss)
func (_Accountant *AccountantFilterer) FilterUpdateMaxLoss(opts *bind.FilterOpts) (*AccountantUpdateMaxLossIterator, error) {

	logs, sub, err := _Accountant.contract.FilterLogs(opts, "UpdateMaxLoss")
	if err != nil {
		return nil, err
	}
	return &AccountantUpdateMaxLossIterator{contract: _Accountant.contract, event: "UpdateMaxLoss", logs: logs, sub: sub}, nil
}

// WatchUpdateMaxLoss is a free log subscription operation binding the contract event 0x18182e268b61d2aada98f23ade23b0ea133d5b0b6712dbfa682dc6da29941c22.
//
// Solidity: event UpdateMaxLoss(uint256 maxLoss)
func (_Accountant *AccountantFilterer) WatchUpdateMaxLoss(opts *bind.WatchOpts, sink chan<- *AccountantUpdateMaxLoss) (event.Subscription, error) {

	logs, sub, err := _Accountant.contract.WatchLogs(opts, "UpdateMaxLoss")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountantUpdateMaxLoss)
				if err := _Accountant.contract.UnpackLog(event, "UpdateMaxLoss", log); err != nil {
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

// ParseUpdateMaxLoss is a log parse operation binding the contract event 0x18182e268b61d2aada98f23ade23b0ea133d5b0b6712dbfa682dc6da29941c22.
//
// Solidity: event UpdateMaxLoss(uint256 maxLoss)
func (_Accountant *AccountantFilterer) ParseUpdateMaxLoss(log types.Log) (*AccountantUpdateMaxLoss, error) {
	event := new(AccountantUpdateMaxLoss)
	if err := _Accountant.contract.UnpackLog(event, "UpdateMaxLoss", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountantUpdateRefundIterator is returned from FilterUpdateRefund and is used to iterate over the raw logs and unpacked data for UpdateRefund events raised by the Accountant contract.
type AccountantUpdateRefundIterator struct {
	Event *AccountantUpdateRefund // Event containing the contract specifics and raw log

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
func (it *AccountantUpdateRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountantUpdateRefund)
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
		it.Event = new(AccountantUpdateRefund)
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
func (it *AccountantUpdateRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountantUpdateRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountantUpdateRefund represents a UpdateRefund event raised by the Accountant contract.
type AccountantUpdateRefund struct {
	Vault    common.Address
	Strategy common.Address
	Refund   bool
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateRefund is a free log retrieval operation binding the contract event 0x04818b53c86c78e687f59a12a2352946865db7ea403c60e7248c481d431ef8f8.
//
// Solidity: event UpdateRefund(address indexed vault, address indexed strategy, bool refund, uint256 amount)
func (_Accountant *AccountantFilterer) FilterUpdateRefund(opts *bind.FilterOpts, vault []common.Address, strategy []common.Address) (*AccountantUpdateRefundIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Accountant.contract.FilterLogs(opts, "UpdateRefund", vaultRule, strategyRule)
	if err != nil {
		return nil, err
	}
	return &AccountantUpdateRefundIterator{contract: _Accountant.contract, event: "UpdateRefund", logs: logs, sub: sub}, nil
}

// WatchUpdateRefund is a free log subscription operation binding the contract event 0x04818b53c86c78e687f59a12a2352946865db7ea403c60e7248c481d431ef8f8.
//
// Solidity: event UpdateRefund(address indexed vault, address indexed strategy, bool refund, uint256 amount)
func (_Accountant *AccountantFilterer) WatchUpdateRefund(opts *bind.WatchOpts, sink chan<- *AccountantUpdateRefund, vault []common.Address, strategy []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Accountant.contract.WatchLogs(opts, "UpdateRefund", vaultRule, strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountantUpdateRefund)
				if err := _Accountant.contract.UnpackLog(event, "UpdateRefund", log); err != nil {
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

// ParseUpdateRefund is a log parse operation binding the contract event 0x04818b53c86c78e687f59a12a2352946865db7ea403c60e7248c481d431ef8f8.
//
// Solidity: event UpdateRefund(address indexed vault, address indexed strategy, bool refund, uint256 amount)
func (_Accountant *AccountantFilterer) ParseUpdateRefund(log types.Log) (*AccountantUpdateRefund, error) {
	event := new(AccountantUpdateRefund)
	if err := _Accountant.contract.UnpackLog(event, "UpdateRefund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountantUpdateVaultManagerIterator is returned from FilterUpdateVaultManager and is used to iterate over the raw logs and unpacked data for UpdateVaultManager events raised by the Accountant contract.
type AccountantUpdateVaultManagerIterator struct {
	Event *AccountantUpdateVaultManager // Event containing the contract specifics and raw log

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
func (it *AccountantUpdateVaultManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountantUpdateVaultManager)
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
		it.Event = new(AccountantUpdateVaultManager)
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
func (it *AccountantUpdateVaultManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountantUpdateVaultManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountantUpdateVaultManager represents a UpdateVaultManager event raised by the Accountant contract.
type AccountantUpdateVaultManager struct {
	NewVaultManager common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateVaultManager is a free log retrieval operation binding the contract event 0xda833a9122ed0b27d5c78c99315bb3758f1b77fb240db484a67fd0f286b263e5.
//
// Solidity: event UpdateVaultManager(address indexed newVaultManager)
func (_Accountant *AccountantFilterer) FilterUpdateVaultManager(opts *bind.FilterOpts, newVaultManager []common.Address) (*AccountantUpdateVaultManagerIterator, error) {

	var newVaultManagerRule []interface{}
	for _, newVaultManagerItem := range newVaultManager {
		newVaultManagerRule = append(newVaultManagerRule, newVaultManagerItem)
	}

	logs, sub, err := _Accountant.contract.FilterLogs(opts, "UpdateVaultManager", newVaultManagerRule)
	if err != nil {
		return nil, err
	}
	return &AccountantUpdateVaultManagerIterator{contract: _Accountant.contract, event: "UpdateVaultManager", logs: logs, sub: sub}, nil
}

// WatchUpdateVaultManager is a free log subscription operation binding the contract event 0xda833a9122ed0b27d5c78c99315bb3758f1b77fb240db484a67fd0f286b263e5.
//
// Solidity: event UpdateVaultManager(address indexed newVaultManager)
func (_Accountant *AccountantFilterer) WatchUpdateVaultManager(opts *bind.WatchOpts, sink chan<- *AccountantUpdateVaultManager, newVaultManager []common.Address) (event.Subscription, error) {

	var newVaultManagerRule []interface{}
	for _, newVaultManagerItem := range newVaultManager {
		newVaultManagerRule = append(newVaultManagerRule, newVaultManagerItem)
	}

	logs, sub, err := _Accountant.contract.WatchLogs(opts, "UpdateVaultManager", newVaultManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountantUpdateVaultManager)
				if err := _Accountant.contract.UnpackLog(event, "UpdateVaultManager", log); err != nil {
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

// ParseUpdateVaultManager is a log parse operation binding the contract event 0xda833a9122ed0b27d5c78c99315bb3758f1b77fb240db484a67fd0f286b263e5.
//
// Solidity: event UpdateVaultManager(address indexed newVaultManager)
func (_Accountant *AccountantFilterer) ParseUpdateVaultManager(log types.Log) (*AccountantUpdateVaultManager, error) {
	event := new(AccountantUpdateVaultManager)
	if err := _Accountant.contract.UnpackLog(event, "UpdateVaultManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountantVaultChangedIterator is returned from FilterVaultChanged and is used to iterate over the raw logs and unpacked data for VaultChanged events raised by the Accountant contract.
type AccountantVaultChangedIterator struct {
	Event *AccountantVaultChanged // Event containing the contract specifics and raw log

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
func (it *AccountantVaultChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountantVaultChanged)
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
		it.Event = new(AccountantVaultChanged)
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
func (it *AccountantVaultChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountantVaultChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountantVaultChanged represents a VaultChanged event raised by the Accountant contract.
type AccountantVaultChanged struct {
	Vault  common.Address
	Change uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterVaultChanged is a free log retrieval operation binding the contract event 0xce96c4db32686d3f0011df1abea0ab6c5794b848868dcbece79961fef7e8198d.
//
// Solidity: event VaultChanged(address indexed vault, uint8 change)
func (_Accountant *AccountantFilterer) FilterVaultChanged(opts *bind.FilterOpts, vault []common.Address) (*AccountantVaultChangedIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _Accountant.contract.FilterLogs(opts, "VaultChanged", vaultRule)
	if err != nil {
		return nil, err
	}
	return &AccountantVaultChangedIterator{contract: _Accountant.contract, event: "VaultChanged", logs: logs, sub: sub}, nil
}

// WatchVaultChanged is a free log subscription operation binding the contract event 0xce96c4db32686d3f0011df1abea0ab6c5794b848868dcbece79961fef7e8198d.
//
// Solidity: event VaultChanged(address indexed vault, uint8 change)
func (_Accountant *AccountantFilterer) WatchVaultChanged(opts *bind.WatchOpts, sink chan<- *AccountantVaultChanged, vault []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _Accountant.contract.WatchLogs(opts, "VaultChanged", vaultRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountantVaultChanged)
				if err := _Accountant.contract.UnpackLog(event, "VaultChanged", log); err != nil {
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

// ParseVaultChanged is a log parse operation binding the contract event 0xce96c4db32686d3f0011df1abea0ab6c5794b848868dcbece79961fef7e8198d.
//
// Solidity: event VaultChanged(address indexed vault, uint8 change)
func (_Accountant *AccountantFilterer) ParseVaultChanged(log types.Log) (*AccountantVaultChanged, error) {
	event := new(AccountantVaultChanged)
	if err := _Accountant.contract.UnpackLog(event, "VaultChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
