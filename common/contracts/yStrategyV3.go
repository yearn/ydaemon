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

// YStrategyV3MetaData contains all meta data concerning the YStrategyV3 contract.
var YStrategyV3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"apiVersion\",\"type\":\"string\"}],\"name\":\"NewTokenizedStrategy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"profit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loss\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"performanceFees\",\"type\":\"uint256\"}],\"name\":\"Reported\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"StrategyShutdown\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newKeeper\",\"type\":\"address\"}],\"name\":\"UpdateKeeper\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newManagement\",\"type\":\"address\"}],\"name\":\"UpdateManagement\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPendingManagement\",\"type\":\"address\"}],\"name\":\"UpdatePendingManagement\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newPerformanceFee\",\"type\":\"uint16\"}],\"name\":\"UpdatePerformanceFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPerformanceFeeRecipient\",\"type\":\"address\"}],\"name\":\"UpdatePerformanceFeeRecipient\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newProfitMaxUnlockTime\",\"type\":\"uint256\"}],\"name\":\"UpdateProfitMaxUnlockTime\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FACTORY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_FEE\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_FEE\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptManagement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"apiVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fullProfitUnlockDate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_management\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_performanceFeeRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_keeper\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"isKeeperOrManagement\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"isManagement\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isShutdown\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"keeper\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastReport\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"management\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"maxMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxMint\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxRedeem\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxWithdraw\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingManagement\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"performanceFee\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"performanceFeeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pricePerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"profitMaxUnlockTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"profitUnlockingRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxLoss\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"report\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"profit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"loss\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_keeper\",\"type\":\"address\"}],\"name\":\"setKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_management\",\"type\":\"address\"}],\"name\":\"setPendingManagement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_performanceFee\",\"type\":\"uint16\"}],\"name\":\"setPerformanceFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_performanceFeeRecipient\",\"type\":\"address\"}],\"name\":\"setPerformanceFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_profitMaxUnlockTime\",\"type\":\"uint256\"}],\"name\":\"setProfitMaxUnlockTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shutdownStrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalIdle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxLoss\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// YStrategyV3ABI is the input ABI used to generate the binding from.
// Deprecated: Use YStrategyV3MetaData.ABI instead.
var YStrategyV3ABI = YStrategyV3MetaData.ABI

// YStrategyV3 is an auto generated Go binding around an Ethereum contract.
type YStrategyV3 struct {
	YStrategyV3Caller     // Read-only binding to the contract
	YStrategyV3Transactor // Write-only binding to the contract
	YStrategyV3Filterer   // Log filterer for contract events
}

// YStrategyV3Caller is an auto generated read-only Go binding around an Ethereum contract.
type YStrategyV3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YStrategyV3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type YStrategyV3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YStrategyV3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YStrategyV3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YStrategyV3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YStrategyV3Session struct {
	Contract     *YStrategyV3      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YStrategyV3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YStrategyV3CallerSession struct {
	Contract *YStrategyV3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// YStrategyV3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YStrategyV3TransactorSession struct {
	Contract     *YStrategyV3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// YStrategyV3Raw is an auto generated low-level Go binding around an Ethereum contract.
type YStrategyV3Raw struct {
	Contract *YStrategyV3 // Generic contract binding to access the raw methods on
}

// YStrategyV3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YStrategyV3CallerRaw struct {
	Contract *YStrategyV3Caller // Generic read-only contract binding to access the raw methods on
}

// YStrategyV3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YStrategyV3TransactorRaw struct {
	Contract *YStrategyV3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewYStrategyV3 creates a new instance of YStrategyV3, bound to a specific deployed contract.
func NewYStrategyV3(address common.Address, backend bind.ContractBackend) (*YStrategyV3, error) {
	contract, err := bindYStrategyV3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YStrategyV3{YStrategyV3Caller: YStrategyV3Caller{contract: contract}, YStrategyV3Transactor: YStrategyV3Transactor{contract: contract}, YStrategyV3Filterer: YStrategyV3Filterer{contract: contract}}, nil
}

// NewYStrategyV3Caller creates a new read-only instance of YStrategyV3, bound to a specific deployed contract.
func NewYStrategyV3Caller(address common.Address, caller bind.ContractCaller) (*YStrategyV3Caller, error) {
	contract, err := bindYStrategyV3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YStrategyV3Caller{contract: contract}, nil
}

// NewYStrategyV3Transactor creates a new write-only instance of YStrategyV3, bound to a specific deployed contract.
func NewYStrategyV3Transactor(address common.Address, transactor bind.ContractTransactor) (*YStrategyV3Transactor, error) {
	contract, err := bindYStrategyV3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YStrategyV3Transactor{contract: contract}, nil
}

// NewYStrategyV3Filterer creates a new log filterer instance of YStrategyV3, bound to a specific deployed contract.
func NewYStrategyV3Filterer(address common.Address, filterer bind.ContractFilterer) (*YStrategyV3Filterer, error) {
	contract, err := bindYStrategyV3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YStrategyV3Filterer{contract: contract}, nil
}

// bindYStrategyV3 binds a generic wrapper to an already deployed contract.
func bindYStrategyV3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(YStrategyV3ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YStrategyV3 *YStrategyV3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YStrategyV3.Contract.YStrategyV3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YStrategyV3 *YStrategyV3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YStrategyV3.Contract.YStrategyV3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YStrategyV3 *YStrategyV3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YStrategyV3.Contract.YStrategyV3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YStrategyV3 *YStrategyV3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YStrategyV3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YStrategyV3 *YStrategyV3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YStrategyV3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YStrategyV3 *YStrategyV3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YStrategyV3.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_YStrategyV3 *YStrategyV3Caller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _YStrategyV3.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_YStrategyV3 *YStrategyV3Session) DOMAINSEPARATOR() ([32]byte, error) {
	return _YStrategyV3.Contract.DOMAINSEPARATOR(&_YStrategyV3.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_YStrategyV3 *YStrategyV3CallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _YStrategyV3.Contract.DOMAINSEPARATOR(&_YStrategyV3.CallOpts)
}

// FACTORY is a free data retrieval call binding the contract method 0x2dd31000.
//
// Solidity: function FACTORY() view returns(address)
func (_YStrategyV3 *YStrategyV3Caller) FACTORY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YStrategyV3.contract.Call(opts, &out, "FACTORY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FACTORY is a free data retrieval call binding the contract method 0x2dd31000.
//
// Solidity: function FACTORY() view returns(address)
func (_YStrategyV3 *YStrategyV3Session) FACTORY() (common.Address, error) {
	return _YStrategyV3.Contract.FACTORY(&_YStrategyV3.CallOpts)
}

// FACTORY is a free data retrieval call binding the contract method 0x2dd31000.
//
// Solidity: function FACTORY() view returns(address)
func (_YStrategyV3 *YStrategyV3CallerSession) FACTORY() (common.Address, error) {
	return _YStrategyV3.Contract.FACTORY(&_YStrategyV3.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint16)
func (_YStrategyV3 *YStrategyV3Caller) MAXFEE(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _YStrategyV3.contract.Call(opts, &out, "MAX_FEE")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint16)
func (_YStrategyV3 *YStrategyV3Session) MAXFEE() (uint16, error) {
	return _YStrategyV3.Contract.MAXFEE(&_YStrategyV3.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint16)
func (_YStrategyV3 *YStrategyV3CallerSession) MAXFEE() (uint16, error) {
	return _YStrategyV3.Contract.MAXFEE(&_YStrategyV3.CallOpts)
}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint16)
func (_YStrategyV3 *YStrategyV3Caller) MINFEE(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _YStrategyV3.contract.Call(opts, &out, "MIN_FEE")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint16)
func (_YStrategyV3 *YStrategyV3Session) MINFEE() (uint16, error) {
	return _YStrategyV3.Contract.MINFEE(&_YStrategyV3.CallOpts)
}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint16)
func (_YStrategyV3 *YStrategyV3CallerSession) MINFEE() (uint16, error) {
	return _YStrategyV3.Contract.MINFEE(&_YStrategyV3.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_YStrategyV3 *YStrategyV3Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyV3.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_YStrategyV3 *YStrategyV3Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _YStrategyV3.Contract.Allowance(&_YStrategyV3.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_YStrategyV3 *YStrategyV3CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _YStrategyV3.Contract.Allowance(&_YStrategyV3.CallOpts, owner, spender)
}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() pure returns(string)
func (_YStrategyV3 *YStrategyV3Caller) ApiVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _YStrategyV3.contract.Call(opts, &out, "apiVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() pure returns(string)
func (_YStrategyV3 *YStrategyV3Session) ApiVersion() (string, error) {
	return _YStrategyV3.Contract.ApiVersion(&_YStrategyV3.CallOpts)
}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() pure returns(string)
func (_YStrategyV3 *YStrategyV3CallerSession) ApiVersion() (string, error) {
	return _YStrategyV3.Contract.ApiVersion(&_YStrategyV3.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_YStrategyV3 *YStrategyV3Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyV3.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_YStrategyV3 *YStrategyV3Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _YStrategyV3.Contract.BalanceOf(&_YStrategyV3.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_YStrategyV3 *YStrategyV3CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _YStrategyV3.Contract.BalanceOf(&_YStrategyV3.CallOpts, account)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_YStrategyV3 *YStrategyV3Caller) ConvertToAssets(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyV3.contract.Call(opts, &out, "convertToAssets", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_YStrategyV3 *YStrategyV3Session) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _YStrategyV3.Contract.ConvertToAssets(&_YStrategyV3.CallOpts, shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_YStrategyV3 *YStrategyV3CallerSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _YStrategyV3.Contract.ConvertToAssets(&_YStrategyV3.CallOpts, shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_YStrategyV3 *YStrategyV3Caller) ConvertToShares(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyV3.contract.Call(opts, &out, "convertToShares", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_YStrategyV3 *YStrategyV3Session) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _YStrategyV3.Contract.ConvertToShares(&_YStrategyV3.CallOpts, assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_YStrategyV3 *YStrategyV3CallerSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _YStrategyV3.Contract.ConvertToShares(&_YStrategyV3.CallOpts, assets)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_YStrategyV3 *YStrategyV3Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _YStrategyV3.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_YStrategyV3 *YStrategyV3Session) Decimals() (uint8, error) {
	return _YStrategyV3.Contract.Decimals(&_YStrategyV3.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_YStrategyV3 *YStrategyV3CallerSession) Decimals() (uint8, error) {
	return _YStrategyV3.Contract.Decimals(&_YStrategyV3.CallOpts)
}

// FullProfitUnlockDate is a free data retrieval call binding the contract method 0x2d632692.
//
// Solidity: function fullProfitUnlockDate() view returns(uint256)
func (_YStrategyV3 *YStrategyV3Caller) FullProfitUnlockDate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyV3.contract.Call(opts, &out, "fullProfitUnlockDate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FullProfitUnlockDate is a free data retrieval call binding the contract method 0x2d632692.
//
// Solidity: function fullProfitUnlockDate() view returns(uint256)
func (_YStrategyV3 *YStrategyV3Session) FullProfitUnlockDate() (*big.Int, error) {
	return _YStrategyV3.Contract.FullProfitUnlockDate(&_YStrategyV3.CallOpts)
}

// FullProfitUnlockDate is a free data retrieval call binding the contract method 0x2d632692.
//
// Solidity: function fullProfitUnlockDate() view returns(uint256)
func (_YStrategyV3 *YStrategyV3CallerSession) FullProfitUnlockDate() (*big.Int, error) {
	return _YStrategyV3.Contract.FullProfitUnlockDate(&_YStrategyV3.CallOpts)
}

// IsKeeperOrManagement is a free data retrieval call binding the contract method 0x1d3b7227.
//
// Solidity: function isKeeperOrManagement(address _sender) view returns()
func (_YStrategyV3 *YStrategyV3Caller) IsKeeperOrManagement(opts *bind.CallOpts, _sender common.Address) error {
	var out []interface{}
	err := _YStrategyV3.contract.Call(opts, &out, "isKeeperOrManagement", _sender)

	if err != nil {
		return err
	}

	return err

}

// IsKeeperOrManagement is a free data retrieval call binding the contract method 0x1d3b7227.
//
// Solidity: function isKeeperOrManagement(address _sender) view returns()
func (_YStrategyV3 *YStrategyV3Session) IsKeeperOrManagement(_sender common.Address) error {
	return _YStrategyV3.Contract.IsKeeperOrManagement(&_YStrategyV3.CallOpts, _sender)
}

// IsKeeperOrManagement is a free data retrieval call binding the contract method 0x1d3b7227.
//
// Solidity: function isKeeperOrManagement(address _sender) view returns()
func (_YStrategyV3 *YStrategyV3CallerSession) IsKeeperOrManagement(_sender common.Address) error {
	return _YStrategyV3.Contract.IsKeeperOrManagement(&_YStrategyV3.CallOpts, _sender)
}

// IsManagement is a free data retrieval call binding the contract method 0xec0c7e28.
//
// Solidity: function isManagement(address _sender) view returns()
func (_YStrategyV3 *YStrategyV3Caller) IsManagement(opts *bind.CallOpts, _sender common.Address) error {
	var out []interface{}
	err := _YStrategyV3.contract.Call(opts, &out, "isManagement", _sender)

	if err != nil {
		return err
	}

	return err

}

// IsManagement is a free data retrieval call binding the contract method 0xec0c7e28.
//
// Solidity: function isManagement(address _sender) view returns()
func (_YStrategyV3 *YStrategyV3Session) IsManagement(_sender common.Address) error {
	return _YStrategyV3.Contract.IsManagement(&_YStrategyV3.CallOpts, _sender)
}

// IsManagement is a free data retrieval call binding the contract method 0xec0c7e28.
//
// Solidity: function isManagement(address _sender) view returns()
func (_YStrategyV3 *YStrategyV3CallerSession) IsManagement(_sender common.Address) error {
	return _YStrategyV3.Contract.IsManagement(&_YStrategyV3.CallOpts, _sender)
}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_YStrategyV3 *YStrategyV3Caller) IsShutdown(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _YStrategyV3.contract.Call(opts, &out, "isShutdown")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_YStrategyV3 *YStrategyV3Session) IsShutdown() (bool, error) {
	return _YStrategyV3.Contract.IsShutdown(&_YStrategyV3.CallOpts)
}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_YStrategyV3 *YStrategyV3CallerSession) IsShutdown() (bool, error) {
	return _YStrategyV3.Contract.IsShutdown(&_YStrategyV3.CallOpts)
}

// Keeper is a free data retrieval call binding the contract method 0xaced1661.
//
// Solidity: function keeper() view returns(address)
func (_YStrategyV3 *YStrategyV3Caller) Keeper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YStrategyV3.contract.Call(opts, &out, "keeper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Keeper is a free data retrieval call binding the contract method 0xaced1661.
//
// Solidity: function keeper() view returns(address)
func (_YStrategyV3 *YStrategyV3Session) Keeper() (common.Address, error) {
	return _YStrategyV3.Contract.Keeper(&_YStrategyV3.CallOpts)
}

// Keeper is a free data retrieval call binding the contract method 0xaced1661.
//
// Solidity: function keeper() view returns(address)
func (_YStrategyV3 *YStrategyV3CallerSession) Keeper() (common.Address, error) {
	return _YStrategyV3.Contract.Keeper(&_YStrategyV3.CallOpts)
}

// LastReport is a free data retrieval call binding the contract method 0xc3535b52.
//
// Solidity: function lastReport() view returns(uint256)
func (_YStrategyV3 *YStrategyV3Caller) LastReport(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyV3.contract.Call(opts, &out, "lastReport")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastReport is a free data retrieval call binding the contract method 0xc3535b52.
//
// Solidity: function lastReport() view returns(uint256)
func (_YStrategyV3 *YStrategyV3Session) LastReport() (*big.Int, error) {
	return _YStrategyV3.Contract.LastReport(&_YStrategyV3.CallOpts)
}

// LastReport is a free data retrieval call binding the contract method 0xc3535b52.
//
// Solidity: function lastReport() view returns(uint256)
func (_YStrategyV3 *YStrategyV3CallerSession) LastReport() (*big.Int, error) {
	return _YStrategyV3.Contract.LastReport(&_YStrategyV3.CallOpts)
}

// Management is a free data retrieval call binding the contract method 0x88a8d602.
//
// Solidity: function management() view returns(address)
func (_YStrategyV3 *YStrategyV3Caller) Management(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YStrategyV3.contract.Call(opts, &out, "management")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Management is a free data retrieval call binding the contract method 0x88a8d602.
//
// Solidity: function management() view returns(address)
func (_YStrategyV3 *YStrategyV3Session) Management() (common.Address, error) {
	return _YStrategyV3.Contract.Management(&_YStrategyV3.CallOpts)
}

// Management is a free data retrieval call binding the contract method 0x88a8d602.
//
// Solidity: function management() view returns(address)
func (_YStrategyV3 *YStrategyV3CallerSession) Management() (common.Address, error) {
	return _YStrategyV3.Contract.Management(&_YStrategyV3.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address _owner) view returns(uint256)
func (_YStrategyV3 *YStrategyV3Caller) MaxDeposit(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyV3.contract.Call(opts, &out, "maxDeposit", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address _owner) view returns(uint256)
func (_YStrategyV3 *YStrategyV3Session) MaxDeposit(_owner common.Address) (*big.Int, error) {
	return _YStrategyV3.Contract.MaxDeposit(&_YStrategyV3.CallOpts, _owner)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address _owner) view returns(uint256)
func (_YStrategyV3 *YStrategyV3CallerSession) MaxDeposit(_owner common.Address) (*big.Int, error) {
	return _YStrategyV3.Contract.MaxDeposit(&_YStrategyV3.CallOpts, _owner)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address _owner) view returns(uint256 _maxMint)
func (_YStrategyV3 *YStrategyV3Caller) MaxMint(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyV3.contract.Call(opts, &out, "maxMint", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address _owner) view returns(uint256 _maxMint)
func (_YStrategyV3 *YStrategyV3Session) MaxMint(_owner common.Address) (*big.Int, error) {
	return _YStrategyV3.Contract.MaxMint(&_YStrategyV3.CallOpts, _owner)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address _owner) view returns(uint256 _maxMint)
func (_YStrategyV3 *YStrategyV3CallerSession) MaxMint(_owner common.Address) (*big.Int, error) {
	return _YStrategyV3.Contract.MaxMint(&_YStrategyV3.CallOpts, _owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address _owner) view returns(uint256 _maxRedeem)
func (_YStrategyV3 *YStrategyV3Caller) MaxRedeem(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyV3.contract.Call(opts, &out, "maxRedeem", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address _owner) view returns(uint256 _maxRedeem)
func (_YStrategyV3 *YStrategyV3Session) MaxRedeem(_owner common.Address) (*big.Int, error) {
	return _YStrategyV3.Contract.MaxRedeem(&_YStrategyV3.CallOpts, _owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address _owner) view returns(uint256 _maxRedeem)
func (_YStrategyV3 *YStrategyV3CallerSession) MaxRedeem(_owner common.Address) (*big.Int, error) {
	return _YStrategyV3.Contract.MaxRedeem(&_YStrategyV3.CallOpts, _owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address _owner) view returns(uint256 _maxWithdraw)
func (_YStrategyV3 *YStrategyV3Caller) MaxWithdraw(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyV3.contract.Call(opts, &out, "maxWithdraw", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address _owner) view returns(uint256 _maxWithdraw)
func (_YStrategyV3 *YStrategyV3Session) MaxWithdraw(_owner common.Address) (*big.Int, error) {
	return _YStrategyV3.Contract.MaxWithdraw(&_YStrategyV3.CallOpts, _owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address _owner) view returns(uint256 _maxWithdraw)
func (_YStrategyV3 *YStrategyV3CallerSession) MaxWithdraw(_owner common.Address) (*big.Int, error) {
	return _YStrategyV3.Contract.MaxWithdraw(&_YStrategyV3.CallOpts, _owner)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YStrategyV3 *YStrategyV3Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _YStrategyV3.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YStrategyV3 *YStrategyV3Session) Name() (string, error) {
	return _YStrategyV3.Contract.Name(&_YStrategyV3.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YStrategyV3 *YStrategyV3CallerSession) Name() (string, error) {
	return _YStrategyV3.Contract.Name(&_YStrategyV3.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address _owner) view returns(uint256)
func (_YStrategyV3 *YStrategyV3Caller) Nonces(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyV3.contract.Call(opts, &out, "nonces", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address _owner) view returns(uint256)
func (_YStrategyV3 *YStrategyV3Session) Nonces(_owner common.Address) (*big.Int, error) {
	return _YStrategyV3.Contract.Nonces(&_YStrategyV3.CallOpts, _owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address _owner) view returns(uint256)
func (_YStrategyV3 *YStrategyV3CallerSession) Nonces(_owner common.Address) (*big.Int, error) {
	return _YStrategyV3.Contract.Nonces(&_YStrategyV3.CallOpts, _owner)
}

// PendingManagement is a free data retrieval call binding the contract method 0x0b68f46f.
//
// Solidity: function pendingManagement() view returns(address)
func (_YStrategyV3 *YStrategyV3Caller) PendingManagement(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YStrategyV3.contract.Call(opts, &out, "pendingManagement")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingManagement is a free data retrieval call binding the contract method 0x0b68f46f.
//
// Solidity: function pendingManagement() view returns(address)
func (_YStrategyV3 *YStrategyV3Session) PendingManagement() (common.Address, error) {
	return _YStrategyV3.Contract.PendingManagement(&_YStrategyV3.CallOpts)
}

// PendingManagement is a free data retrieval call binding the contract method 0x0b68f46f.
//
// Solidity: function pendingManagement() view returns(address)
func (_YStrategyV3 *YStrategyV3CallerSession) PendingManagement() (common.Address, error) {
	return _YStrategyV3.Contract.PendingManagement(&_YStrategyV3.CallOpts)
}

// PerformanceFee is a free data retrieval call binding the contract method 0x87788782.
//
// Solidity: function performanceFee() view returns(uint16)
func (_YStrategyV3 *YStrategyV3Caller) PerformanceFee(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _YStrategyV3.contract.Call(opts, &out, "performanceFee")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// PerformanceFee is a free data retrieval call binding the contract method 0x87788782.
//
// Solidity: function performanceFee() view returns(uint16)
func (_YStrategyV3 *YStrategyV3Session) PerformanceFee() (uint16, error) {
	return _YStrategyV3.Contract.PerformanceFee(&_YStrategyV3.CallOpts)
}

// PerformanceFee is a free data retrieval call binding the contract method 0x87788782.
//
// Solidity: function performanceFee() view returns(uint16)
func (_YStrategyV3 *YStrategyV3CallerSession) PerformanceFee() (uint16, error) {
	return _YStrategyV3.Contract.PerformanceFee(&_YStrategyV3.CallOpts)
}

// PerformanceFeeRecipient is a free data retrieval call binding the contract method 0xed27f7c9.
//
// Solidity: function performanceFeeRecipient() view returns(address)
func (_YStrategyV3 *YStrategyV3Caller) PerformanceFeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YStrategyV3.contract.Call(opts, &out, "performanceFeeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PerformanceFeeRecipient is a free data retrieval call binding the contract method 0xed27f7c9.
//
// Solidity: function performanceFeeRecipient() view returns(address)
func (_YStrategyV3 *YStrategyV3Session) PerformanceFeeRecipient() (common.Address, error) {
	return _YStrategyV3.Contract.PerformanceFeeRecipient(&_YStrategyV3.CallOpts)
}

// PerformanceFeeRecipient is a free data retrieval call binding the contract method 0xed27f7c9.
//
// Solidity: function performanceFeeRecipient() view returns(address)
func (_YStrategyV3 *YStrategyV3CallerSession) PerformanceFeeRecipient() (common.Address, error) {
	return _YStrategyV3.Contract.PerformanceFeeRecipient(&_YStrategyV3.CallOpts)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_YStrategyV3 *YStrategyV3Caller) PreviewDeposit(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyV3.contract.Call(opts, &out, "previewDeposit", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_YStrategyV3 *YStrategyV3Session) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _YStrategyV3.Contract.PreviewDeposit(&_YStrategyV3.CallOpts, assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_YStrategyV3 *YStrategyV3CallerSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _YStrategyV3.Contract.PreviewDeposit(&_YStrategyV3.CallOpts, assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_YStrategyV3 *YStrategyV3Caller) PreviewMint(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyV3.contract.Call(opts, &out, "previewMint", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_YStrategyV3 *YStrategyV3Session) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _YStrategyV3.Contract.PreviewMint(&_YStrategyV3.CallOpts, shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_YStrategyV3 *YStrategyV3CallerSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _YStrategyV3.Contract.PreviewMint(&_YStrategyV3.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_YStrategyV3 *YStrategyV3Caller) PreviewRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyV3.contract.Call(opts, &out, "previewRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_YStrategyV3 *YStrategyV3Session) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _YStrategyV3.Contract.PreviewRedeem(&_YStrategyV3.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_YStrategyV3 *YStrategyV3CallerSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _YStrategyV3.Contract.PreviewRedeem(&_YStrategyV3.CallOpts, shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_YStrategyV3 *YStrategyV3Caller) PreviewWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyV3.contract.Call(opts, &out, "previewWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_YStrategyV3 *YStrategyV3Session) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _YStrategyV3.Contract.PreviewWithdraw(&_YStrategyV3.CallOpts, assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_YStrategyV3 *YStrategyV3CallerSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _YStrategyV3.Contract.PreviewWithdraw(&_YStrategyV3.CallOpts, assets)
}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_YStrategyV3 *YStrategyV3Caller) PricePerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyV3.contract.Call(opts, &out, "pricePerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_YStrategyV3 *YStrategyV3Session) PricePerShare() (*big.Int, error) {
	return _YStrategyV3.Contract.PricePerShare(&_YStrategyV3.CallOpts)
}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_YStrategyV3 *YStrategyV3CallerSession) PricePerShare() (*big.Int, error) {
	return _YStrategyV3.Contract.PricePerShare(&_YStrategyV3.CallOpts)
}

// ProfitMaxUnlockTime is a free data retrieval call binding the contract method 0x0952864e.
//
// Solidity: function profitMaxUnlockTime() view returns(uint256)
func (_YStrategyV3 *YStrategyV3Caller) ProfitMaxUnlockTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyV3.contract.Call(opts, &out, "profitMaxUnlockTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProfitMaxUnlockTime is a free data retrieval call binding the contract method 0x0952864e.
//
// Solidity: function profitMaxUnlockTime() view returns(uint256)
func (_YStrategyV3 *YStrategyV3Session) ProfitMaxUnlockTime() (*big.Int, error) {
	return _YStrategyV3.Contract.ProfitMaxUnlockTime(&_YStrategyV3.CallOpts)
}

// ProfitMaxUnlockTime is a free data retrieval call binding the contract method 0x0952864e.
//
// Solidity: function profitMaxUnlockTime() view returns(uint256)
func (_YStrategyV3 *YStrategyV3CallerSession) ProfitMaxUnlockTime() (*big.Int, error) {
	return _YStrategyV3.Contract.ProfitMaxUnlockTime(&_YStrategyV3.CallOpts)
}

// ProfitUnlockingRate is a free data retrieval call binding the contract method 0x5141eebb.
//
// Solidity: function profitUnlockingRate() view returns(uint256)
func (_YStrategyV3 *YStrategyV3Caller) ProfitUnlockingRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyV3.contract.Call(opts, &out, "profitUnlockingRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProfitUnlockingRate is a free data retrieval call binding the contract method 0x5141eebb.
//
// Solidity: function profitUnlockingRate() view returns(uint256)
func (_YStrategyV3 *YStrategyV3Session) ProfitUnlockingRate() (*big.Int, error) {
	return _YStrategyV3.Contract.ProfitUnlockingRate(&_YStrategyV3.CallOpts)
}

// ProfitUnlockingRate is a free data retrieval call binding the contract method 0x5141eebb.
//
// Solidity: function profitUnlockingRate() view returns(uint256)
func (_YStrategyV3 *YStrategyV3CallerSession) ProfitUnlockingRate() (*big.Int, error) {
	return _YStrategyV3.Contract.ProfitUnlockingRate(&_YStrategyV3.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_YStrategyV3 *YStrategyV3Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _YStrategyV3.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_YStrategyV3 *YStrategyV3Session) Symbol() (string, error) {
	return _YStrategyV3.Contract.Symbol(&_YStrategyV3.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_YStrategyV3 *YStrategyV3CallerSession) Symbol() (string, error) {
	return _YStrategyV3.Contract.Symbol(&_YStrategyV3.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_YStrategyV3 *YStrategyV3Caller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyV3.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_YStrategyV3 *YStrategyV3Session) TotalAssets() (*big.Int, error) {
	return _YStrategyV3.Contract.TotalAssets(&_YStrategyV3.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_YStrategyV3 *YStrategyV3CallerSession) TotalAssets() (*big.Int, error) {
	return _YStrategyV3.Contract.TotalAssets(&_YStrategyV3.CallOpts)
}

// TotalDebt is a free data retrieval call binding the contract method 0xfc7b9c18.
//
// Solidity: function totalDebt() view returns(uint256)
func (_YStrategyV3 *YStrategyV3Caller) TotalDebt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyV3.contract.Call(opts, &out, "totalDebt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDebt is a free data retrieval call binding the contract method 0xfc7b9c18.
//
// Solidity: function totalDebt() view returns(uint256)
func (_YStrategyV3 *YStrategyV3Session) TotalDebt() (*big.Int, error) {
	return _YStrategyV3.Contract.TotalDebt(&_YStrategyV3.CallOpts)
}

// TotalDebt is a free data retrieval call binding the contract method 0xfc7b9c18.
//
// Solidity: function totalDebt() view returns(uint256)
func (_YStrategyV3 *YStrategyV3CallerSession) TotalDebt() (*big.Int, error) {
	return _YStrategyV3.Contract.TotalDebt(&_YStrategyV3.CallOpts)
}

// TotalIdle is a free data retrieval call binding the contract method 0x9aa7df94.
//
// Solidity: function totalIdle() view returns(uint256)
func (_YStrategyV3 *YStrategyV3Caller) TotalIdle(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyV3.contract.Call(opts, &out, "totalIdle")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalIdle is a free data retrieval call binding the contract method 0x9aa7df94.
//
// Solidity: function totalIdle() view returns(uint256)
func (_YStrategyV3 *YStrategyV3Session) TotalIdle() (*big.Int, error) {
	return _YStrategyV3.Contract.TotalIdle(&_YStrategyV3.CallOpts)
}

// TotalIdle is a free data retrieval call binding the contract method 0x9aa7df94.
//
// Solidity: function totalIdle() view returns(uint256)
func (_YStrategyV3 *YStrategyV3CallerSession) TotalIdle() (*big.Int, error) {
	return _YStrategyV3.Contract.TotalIdle(&_YStrategyV3.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_YStrategyV3 *YStrategyV3Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyV3.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_YStrategyV3 *YStrategyV3Session) TotalSupply() (*big.Int, error) {
	return _YStrategyV3.Contract.TotalSupply(&_YStrategyV3.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_YStrategyV3 *YStrategyV3CallerSession) TotalSupply() (*big.Int, error) {
	return _YStrategyV3.Contract.TotalSupply(&_YStrategyV3.CallOpts)
}

// AcceptManagement is a paid mutator transaction binding the contract method 0xc8c2fe6c.
//
// Solidity: function acceptManagement() returns()
func (_YStrategyV3 *YStrategyV3Transactor) AcceptManagement(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YStrategyV3.contract.Transact(opts, "acceptManagement")
}

// AcceptManagement is a paid mutator transaction binding the contract method 0xc8c2fe6c.
//
// Solidity: function acceptManagement() returns()
func (_YStrategyV3 *YStrategyV3Session) AcceptManagement() (*types.Transaction, error) {
	return _YStrategyV3.Contract.AcceptManagement(&_YStrategyV3.TransactOpts)
}

// AcceptManagement is a paid mutator transaction binding the contract method 0xc8c2fe6c.
//
// Solidity: function acceptManagement() returns()
func (_YStrategyV3 *YStrategyV3TransactorSession) AcceptManagement() (*types.Transaction, error) {
	return _YStrategyV3.Contract.AcceptManagement(&_YStrategyV3.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_YStrategyV3 *YStrategyV3Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YStrategyV3.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_YStrategyV3 *YStrategyV3Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YStrategyV3.Contract.Approve(&_YStrategyV3.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_YStrategyV3 *YStrategyV3TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YStrategyV3.Contract.Approve(&_YStrategyV3.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_YStrategyV3 *YStrategyV3Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _YStrategyV3.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_YStrategyV3 *YStrategyV3Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _YStrategyV3.Contract.DecreaseAllowance(&_YStrategyV3.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_YStrategyV3 *YStrategyV3TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _YStrategyV3.Contract.DecreaseAllowance(&_YStrategyV3.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_YStrategyV3 *YStrategyV3Transactor) Deposit(opts *bind.TransactOpts, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _YStrategyV3.contract.Transact(opts, "deposit", assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_YStrategyV3 *YStrategyV3Session) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _YStrategyV3.Contract.Deposit(&_YStrategyV3.TransactOpts, assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_YStrategyV3 *YStrategyV3TransactorSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _YStrategyV3.Contract.Deposit(&_YStrategyV3.TransactOpts, assets, receiver)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _amount) returns()
func (_YStrategyV3 *YStrategyV3Transactor) EmergencyWithdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _YStrategyV3.contract.Transact(opts, "emergencyWithdraw", _amount)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _amount) returns()
func (_YStrategyV3 *YStrategyV3Session) EmergencyWithdraw(_amount *big.Int) (*types.Transaction, error) {
	return _YStrategyV3.Contract.EmergencyWithdraw(&_YStrategyV3.TransactOpts, _amount)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _amount) returns()
func (_YStrategyV3 *YStrategyV3TransactorSession) EmergencyWithdraw(_amount *big.Int) (*types.Transaction, error) {
	return _YStrategyV3.Contract.EmergencyWithdraw(&_YStrategyV3.TransactOpts, _amount)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_YStrategyV3 *YStrategyV3Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _YStrategyV3.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_YStrategyV3 *YStrategyV3Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _YStrategyV3.Contract.IncreaseAllowance(&_YStrategyV3.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_YStrategyV3 *YStrategyV3TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _YStrategyV3.Contract.IncreaseAllowance(&_YStrategyV3.TransactOpts, spender, addedValue)
}

// Init is a paid mutator transaction binding the contract method 0x2ecfe315.
//
// Solidity: function init(address _asset, string _name, address _management, address _performanceFeeRecipient, address _keeper) returns()
func (_YStrategyV3 *YStrategyV3Transactor) Init(opts *bind.TransactOpts, _asset common.Address, _name string, _management common.Address, _performanceFeeRecipient common.Address, _keeper common.Address) (*types.Transaction, error) {
	return _YStrategyV3.contract.Transact(opts, "init", _asset, _name, _management, _performanceFeeRecipient, _keeper)
}

// Init is a paid mutator transaction binding the contract method 0x2ecfe315.
//
// Solidity: function init(address _asset, string _name, address _management, address _performanceFeeRecipient, address _keeper) returns()
func (_YStrategyV3 *YStrategyV3Session) Init(_asset common.Address, _name string, _management common.Address, _performanceFeeRecipient common.Address, _keeper common.Address) (*types.Transaction, error) {
	return _YStrategyV3.Contract.Init(&_YStrategyV3.TransactOpts, _asset, _name, _management, _performanceFeeRecipient, _keeper)
}

// Init is a paid mutator transaction binding the contract method 0x2ecfe315.
//
// Solidity: function init(address _asset, string _name, address _management, address _performanceFeeRecipient, address _keeper) returns()
func (_YStrategyV3 *YStrategyV3TransactorSession) Init(_asset common.Address, _name string, _management common.Address, _performanceFeeRecipient common.Address, _keeper common.Address) (*types.Transaction, error) {
	return _YStrategyV3.Contract.Init(&_YStrategyV3.TransactOpts, _asset, _name, _management, _performanceFeeRecipient, _keeper)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_YStrategyV3 *YStrategyV3Transactor) Mint(opts *bind.TransactOpts, shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _YStrategyV3.contract.Transact(opts, "mint", shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_YStrategyV3 *YStrategyV3Session) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _YStrategyV3.Contract.Mint(&_YStrategyV3.TransactOpts, shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_YStrategyV3 *YStrategyV3TransactorSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _YStrategyV3.Contract.Mint(&_YStrategyV3.TransactOpts, shares, receiver)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_YStrategyV3 *YStrategyV3Transactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _YStrategyV3.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_YStrategyV3 *YStrategyV3Session) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _YStrategyV3.Contract.Permit(&_YStrategyV3.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_YStrategyV3 *YStrategyV3TransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _YStrategyV3.Contract.Permit(&_YStrategyV3.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Redeem is a paid mutator transaction binding the contract method 0x9f40a7b3.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner, uint256 maxLoss) returns(uint256)
func (_YStrategyV3 *YStrategyV3Transactor) Redeem(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address, maxLoss *big.Int) (*types.Transaction, error) {
	return _YStrategyV3.contract.Transact(opts, "redeem", shares, receiver, owner, maxLoss)
}

// Redeem is a paid mutator transaction binding the contract method 0x9f40a7b3.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner, uint256 maxLoss) returns(uint256)
func (_YStrategyV3 *YStrategyV3Session) Redeem(shares *big.Int, receiver common.Address, owner common.Address, maxLoss *big.Int) (*types.Transaction, error) {
	return _YStrategyV3.Contract.Redeem(&_YStrategyV3.TransactOpts, shares, receiver, owner, maxLoss)
}

// Redeem is a paid mutator transaction binding the contract method 0x9f40a7b3.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner, uint256 maxLoss) returns(uint256)
func (_YStrategyV3 *YStrategyV3TransactorSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address, maxLoss *big.Int) (*types.Transaction, error) {
	return _YStrategyV3.Contract.Redeem(&_YStrategyV3.TransactOpts, shares, receiver, owner, maxLoss)
}

// Redeem0 is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_YStrategyV3 *YStrategyV3Transactor) Redeem0(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _YStrategyV3.contract.Transact(opts, "redeem0", shares, receiver, owner)
}

// Redeem0 is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_YStrategyV3 *YStrategyV3Session) Redeem0(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _YStrategyV3.Contract.Redeem0(&_YStrategyV3.TransactOpts, shares, receiver, owner)
}

// Redeem0 is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_YStrategyV3 *YStrategyV3TransactorSession) Redeem0(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _YStrategyV3.Contract.Redeem0(&_YStrategyV3.TransactOpts, shares, receiver, owner)
}

// Report is a paid mutator transaction binding the contract method 0x2606a10b.
//
// Solidity: function report() returns(uint256 profit, uint256 loss)
func (_YStrategyV3 *YStrategyV3Transactor) Report(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YStrategyV3.contract.Transact(opts, "report")
}

// Report is a paid mutator transaction binding the contract method 0x2606a10b.
//
// Solidity: function report() returns(uint256 profit, uint256 loss)
func (_YStrategyV3 *YStrategyV3Session) Report() (*types.Transaction, error) {
	return _YStrategyV3.Contract.Report(&_YStrategyV3.TransactOpts)
}

// Report is a paid mutator transaction binding the contract method 0x2606a10b.
//
// Solidity: function report() returns(uint256 profit, uint256 loss)
func (_YStrategyV3 *YStrategyV3TransactorSession) Report() (*types.Transaction, error) {
	return _YStrategyV3.Contract.Report(&_YStrategyV3.TransactOpts)
}

// SetKeeper is a paid mutator transaction binding the contract method 0x748747e6.
//
// Solidity: function setKeeper(address _keeper) returns()
func (_YStrategyV3 *YStrategyV3Transactor) SetKeeper(opts *bind.TransactOpts, _keeper common.Address) (*types.Transaction, error) {
	return _YStrategyV3.contract.Transact(opts, "setKeeper", _keeper)
}

// SetKeeper is a paid mutator transaction binding the contract method 0x748747e6.
//
// Solidity: function setKeeper(address _keeper) returns()
func (_YStrategyV3 *YStrategyV3Session) SetKeeper(_keeper common.Address) (*types.Transaction, error) {
	return _YStrategyV3.Contract.SetKeeper(&_YStrategyV3.TransactOpts, _keeper)
}

// SetKeeper is a paid mutator transaction binding the contract method 0x748747e6.
//
// Solidity: function setKeeper(address _keeper) returns()
func (_YStrategyV3 *YStrategyV3TransactorSession) SetKeeper(_keeper common.Address) (*types.Transaction, error) {
	return _YStrategyV3.Contract.SetKeeper(&_YStrategyV3.TransactOpts, _keeper)
}

// SetPendingManagement is a paid mutator transaction binding the contract method 0xf629b790.
//
// Solidity: function setPendingManagement(address _management) returns()
func (_YStrategyV3 *YStrategyV3Transactor) SetPendingManagement(opts *bind.TransactOpts, _management common.Address) (*types.Transaction, error) {
	return _YStrategyV3.contract.Transact(opts, "setPendingManagement", _management)
}

// SetPendingManagement is a paid mutator transaction binding the contract method 0xf629b790.
//
// Solidity: function setPendingManagement(address _management) returns()
func (_YStrategyV3 *YStrategyV3Session) SetPendingManagement(_management common.Address) (*types.Transaction, error) {
	return _YStrategyV3.Contract.SetPendingManagement(&_YStrategyV3.TransactOpts, _management)
}

// SetPendingManagement is a paid mutator transaction binding the contract method 0xf629b790.
//
// Solidity: function setPendingManagement(address _management) returns()
func (_YStrategyV3 *YStrategyV3TransactorSession) SetPendingManagement(_management common.Address) (*types.Transaction, error) {
	return _YStrategyV3.Contract.SetPendingManagement(&_YStrategyV3.TransactOpts, _management)
}

// SetPerformanceFee is a paid mutator transaction binding the contract method 0xaa290e6d.
//
// Solidity: function setPerformanceFee(uint16 _performanceFee) returns()
func (_YStrategyV3 *YStrategyV3Transactor) SetPerformanceFee(opts *bind.TransactOpts, _performanceFee uint16) (*types.Transaction, error) {
	return _YStrategyV3.contract.Transact(opts, "setPerformanceFee", _performanceFee)
}

// SetPerformanceFee is a paid mutator transaction binding the contract method 0xaa290e6d.
//
// Solidity: function setPerformanceFee(uint16 _performanceFee) returns()
func (_YStrategyV3 *YStrategyV3Session) SetPerformanceFee(_performanceFee uint16) (*types.Transaction, error) {
	return _YStrategyV3.Contract.SetPerformanceFee(&_YStrategyV3.TransactOpts, _performanceFee)
}

// SetPerformanceFee is a paid mutator transaction binding the contract method 0xaa290e6d.
//
// Solidity: function setPerformanceFee(uint16 _performanceFee) returns()
func (_YStrategyV3 *YStrategyV3TransactorSession) SetPerformanceFee(_performanceFee uint16) (*types.Transaction, error) {
	return _YStrategyV3.Contract.SetPerformanceFee(&_YStrategyV3.TransactOpts, _performanceFee)
}

// SetPerformanceFeeRecipient is a paid mutator transaction binding the contract method 0x6a5f1aa2.
//
// Solidity: function setPerformanceFeeRecipient(address _performanceFeeRecipient) returns()
func (_YStrategyV3 *YStrategyV3Transactor) SetPerformanceFeeRecipient(opts *bind.TransactOpts, _performanceFeeRecipient common.Address) (*types.Transaction, error) {
	return _YStrategyV3.contract.Transact(opts, "setPerformanceFeeRecipient", _performanceFeeRecipient)
}

// SetPerformanceFeeRecipient is a paid mutator transaction binding the contract method 0x6a5f1aa2.
//
// Solidity: function setPerformanceFeeRecipient(address _performanceFeeRecipient) returns()
func (_YStrategyV3 *YStrategyV3Session) SetPerformanceFeeRecipient(_performanceFeeRecipient common.Address) (*types.Transaction, error) {
	return _YStrategyV3.Contract.SetPerformanceFeeRecipient(&_YStrategyV3.TransactOpts, _performanceFeeRecipient)
}

// SetPerformanceFeeRecipient is a paid mutator transaction binding the contract method 0x6a5f1aa2.
//
// Solidity: function setPerformanceFeeRecipient(address _performanceFeeRecipient) returns()
func (_YStrategyV3 *YStrategyV3TransactorSession) SetPerformanceFeeRecipient(_performanceFeeRecipient common.Address) (*types.Transaction, error) {
	return _YStrategyV3.Contract.SetPerformanceFeeRecipient(&_YStrategyV3.TransactOpts, _performanceFeeRecipient)
}

// SetProfitMaxUnlockTime is a paid mutator transaction binding the contract method 0xdf69b22a.
//
// Solidity: function setProfitMaxUnlockTime(uint256 _profitMaxUnlockTime) returns()
func (_YStrategyV3 *YStrategyV3Transactor) SetProfitMaxUnlockTime(opts *bind.TransactOpts, _profitMaxUnlockTime *big.Int) (*types.Transaction, error) {
	return _YStrategyV3.contract.Transact(opts, "setProfitMaxUnlockTime", _profitMaxUnlockTime)
}

// SetProfitMaxUnlockTime is a paid mutator transaction binding the contract method 0xdf69b22a.
//
// Solidity: function setProfitMaxUnlockTime(uint256 _profitMaxUnlockTime) returns()
func (_YStrategyV3 *YStrategyV3Session) SetProfitMaxUnlockTime(_profitMaxUnlockTime *big.Int) (*types.Transaction, error) {
	return _YStrategyV3.Contract.SetProfitMaxUnlockTime(&_YStrategyV3.TransactOpts, _profitMaxUnlockTime)
}

// SetProfitMaxUnlockTime is a paid mutator transaction binding the contract method 0xdf69b22a.
//
// Solidity: function setProfitMaxUnlockTime(uint256 _profitMaxUnlockTime) returns()
func (_YStrategyV3 *YStrategyV3TransactorSession) SetProfitMaxUnlockTime(_profitMaxUnlockTime *big.Int) (*types.Transaction, error) {
	return _YStrategyV3.Contract.SetProfitMaxUnlockTime(&_YStrategyV3.TransactOpts, _profitMaxUnlockTime)
}

// ShutdownStrategy is a paid mutator transaction binding the contract method 0xbe8f1668.
//
// Solidity: function shutdownStrategy() returns()
func (_YStrategyV3 *YStrategyV3Transactor) ShutdownStrategy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YStrategyV3.contract.Transact(opts, "shutdownStrategy")
}

// ShutdownStrategy is a paid mutator transaction binding the contract method 0xbe8f1668.
//
// Solidity: function shutdownStrategy() returns()
func (_YStrategyV3 *YStrategyV3Session) ShutdownStrategy() (*types.Transaction, error) {
	return _YStrategyV3.Contract.ShutdownStrategy(&_YStrategyV3.TransactOpts)
}

// ShutdownStrategy is a paid mutator transaction binding the contract method 0xbe8f1668.
//
// Solidity: function shutdownStrategy() returns()
func (_YStrategyV3 *YStrategyV3TransactorSession) ShutdownStrategy() (*types.Transaction, error) {
	return _YStrategyV3.Contract.ShutdownStrategy(&_YStrategyV3.TransactOpts)
}

// Tend is a paid mutator transaction binding the contract method 0x440368a3.
//
// Solidity: function tend() returns()
func (_YStrategyV3 *YStrategyV3Transactor) Tend(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YStrategyV3.contract.Transact(opts, "tend")
}

// Tend is a paid mutator transaction binding the contract method 0x440368a3.
//
// Solidity: function tend() returns()
func (_YStrategyV3 *YStrategyV3Session) Tend() (*types.Transaction, error) {
	return _YStrategyV3.Contract.Tend(&_YStrategyV3.TransactOpts)
}

// Tend is a paid mutator transaction binding the contract method 0x440368a3.
//
// Solidity: function tend() returns()
func (_YStrategyV3 *YStrategyV3TransactorSession) Tend() (*types.Transaction, error) {
	return _YStrategyV3.Contract.Tend(&_YStrategyV3.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_YStrategyV3 *YStrategyV3Transactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YStrategyV3.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_YStrategyV3 *YStrategyV3Session) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YStrategyV3.Contract.Transfer(&_YStrategyV3.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_YStrategyV3 *YStrategyV3TransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YStrategyV3.Contract.Transfer(&_YStrategyV3.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_YStrategyV3 *YStrategyV3Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YStrategyV3.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_YStrategyV3 *YStrategyV3Session) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YStrategyV3.Contract.TransferFrom(&_YStrategyV3.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_YStrategyV3 *YStrategyV3TransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YStrategyV3.Contract.TransferFrom(&_YStrategyV3.TransactOpts, from, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa318c1a4.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner, uint256 maxLoss) returns(uint256 shares)
func (_YStrategyV3 *YStrategyV3Transactor) Withdraw(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address, maxLoss *big.Int) (*types.Transaction, error) {
	return _YStrategyV3.contract.Transact(opts, "withdraw", assets, receiver, owner, maxLoss)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa318c1a4.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner, uint256 maxLoss) returns(uint256 shares)
func (_YStrategyV3 *YStrategyV3Session) Withdraw(assets *big.Int, receiver common.Address, owner common.Address, maxLoss *big.Int) (*types.Transaction, error) {
	return _YStrategyV3.Contract.Withdraw(&_YStrategyV3.TransactOpts, assets, receiver, owner, maxLoss)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa318c1a4.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner, uint256 maxLoss) returns(uint256 shares)
func (_YStrategyV3 *YStrategyV3TransactorSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address, maxLoss *big.Int) (*types.Transaction, error) {
	return _YStrategyV3.Contract.Withdraw(&_YStrategyV3.TransactOpts, assets, receiver, owner, maxLoss)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_YStrategyV3 *YStrategyV3Transactor) Withdraw0(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _YStrategyV3.contract.Transact(opts, "withdraw0", assets, receiver, owner)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_YStrategyV3 *YStrategyV3Session) Withdraw0(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _YStrategyV3.Contract.Withdraw0(&_YStrategyV3.TransactOpts, assets, receiver, owner)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_YStrategyV3 *YStrategyV3TransactorSession) Withdraw0(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _YStrategyV3.Contract.Withdraw0(&_YStrategyV3.TransactOpts, assets, receiver, owner)
}

// YStrategyV3ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the YStrategyV3 contract.
type YStrategyV3ApprovalIterator struct {
	Event *YStrategyV3Approval // Event containing the contract specifics and raw log

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
func (it *YStrategyV3ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YStrategyV3Approval)
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
		it.Event = new(YStrategyV3Approval)
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
func (it *YStrategyV3ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YStrategyV3ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YStrategyV3Approval represents a Approval event raised by the YStrategyV3 contract.
type YStrategyV3Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_YStrategyV3 *YStrategyV3Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*YStrategyV3ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _YStrategyV3.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &YStrategyV3ApprovalIterator{contract: _YStrategyV3.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_YStrategyV3 *YStrategyV3Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *YStrategyV3Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _YStrategyV3.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YStrategyV3Approval)
				if err := _YStrategyV3.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_YStrategyV3 *YStrategyV3Filterer) ParseApproval(log types.Log) (*YStrategyV3Approval, error) {
	event := new(YStrategyV3Approval)
	if err := _YStrategyV3.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YStrategyV3DepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the YStrategyV3 contract.
type YStrategyV3DepositIterator struct {
	Event *YStrategyV3Deposit // Event containing the contract specifics and raw log

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
func (it *YStrategyV3DepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YStrategyV3Deposit)
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
		it.Event = new(YStrategyV3Deposit)
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
func (it *YStrategyV3DepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YStrategyV3DepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YStrategyV3Deposit represents a Deposit event raised by the YStrategyV3 contract.
type YStrategyV3Deposit struct {
	Caller common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed caller, address indexed owner, uint256 assets, uint256 shares)
func (_YStrategyV3 *YStrategyV3Filterer) FilterDeposit(opts *bind.FilterOpts, caller []common.Address, owner []common.Address) (*YStrategyV3DepositIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _YStrategyV3.contract.FilterLogs(opts, "Deposit", callerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &YStrategyV3DepositIterator{contract: _YStrategyV3.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed caller, address indexed owner, uint256 assets, uint256 shares)
func (_YStrategyV3 *YStrategyV3Filterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *YStrategyV3Deposit, caller []common.Address, owner []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _YStrategyV3.contract.WatchLogs(opts, "Deposit", callerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YStrategyV3Deposit)
				if err := _YStrategyV3.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed caller, address indexed owner, uint256 assets, uint256 shares)
func (_YStrategyV3 *YStrategyV3Filterer) ParseDeposit(log types.Log) (*YStrategyV3Deposit, error) {
	event := new(YStrategyV3Deposit)
	if err := _YStrategyV3.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YStrategyV3NewTokenizedStrategyIterator is returned from FilterNewTokenizedStrategy and is used to iterate over the raw logs and unpacked data for NewTokenizedStrategy events raised by the YStrategyV3 contract.
type YStrategyV3NewTokenizedStrategyIterator struct {
	Event *YStrategyV3NewTokenizedStrategy // Event containing the contract specifics and raw log

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
func (it *YStrategyV3NewTokenizedStrategyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YStrategyV3NewTokenizedStrategy)
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
		it.Event = new(YStrategyV3NewTokenizedStrategy)
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
func (it *YStrategyV3NewTokenizedStrategyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YStrategyV3NewTokenizedStrategyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YStrategyV3NewTokenizedStrategy represents a NewTokenizedStrategy event raised by the YStrategyV3 contract.
type YStrategyV3NewTokenizedStrategy struct {
	Strategy   common.Address
	Asset      common.Address
	ApiVersion string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewTokenizedStrategy is a free log retrieval operation binding the contract event 0xfb1616746b8474b6b7c67f2fe5ada156ed24774d0efe8bfe529cf537ba173330.
//
// Solidity: event NewTokenizedStrategy(address indexed strategy, address indexed asset, string apiVersion)
func (_YStrategyV3 *YStrategyV3Filterer) FilterNewTokenizedStrategy(opts *bind.FilterOpts, strategy []common.Address, asset []common.Address) (*YStrategyV3NewTokenizedStrategyIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _YStrategyV3.contract.FilterLogs(opts, "NewTokenizedStrategy", strategyRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &YStrategyV3NewTokenizedStrategyIterator{contract: _YStrategyV3.contract, event: "NewTokenizedStrategy", logs: logs, sub: sub}, nil
}

// WatchNewTokenizedStrategy is a free log subscription operation binding the contract event 0xfb1616746b8474b6b7c67f2fe5ada156ed24774d0efe8bfe529cf537ba173330.
//
// Solidity: event NewTokenizedStrategy(address indexed strategy, address indexed asset, string apiVersion)
func (_YStrategyV3 *YStrategyV3Filterer) WatchNewTokenizedStrategy(opts *bind.WatchOpts, sink chan<- *YStrategyV3NewTokenizedStrategy, strategy []common.Address, asset []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _YStrategyV3.contract.WatchLogs(opts, "NewTokenizedStrategy", strategyRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YStrategyV3NewTokenizedStrategy)
				if err := _YStrategyV3.contract.UnpackLog(event, "NewTokenizedStrategy", log); err != nil {
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

// ParseNewTokenizedStrategy is a log parse operation binding the contract event 0xfb1616746b8474b6b7c67f2fe5ada156ed24774d0efe8bfe529cf537ba173330.
//
// Solidity: event NewTokenizedStrategy(address indexed strategy, address indexed asset, string apiVersion)
func (_YStrategyV3 *YStrategyV3Filterer) ParseNewTokenizedStrategy(log types.Log) (*YStrategyV3NewTokenizedStrategy, error) {
	event := new(YStrategyV3NewTokenizedStrategy)
	if err := _YStrategyV3.contract.UnpackLog(event, "NewTokenizedStrategy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YStrategyV3ReportedIterator is returned from FilterReported and is used to iterate over the raw logs and unpacked data for Reported events raised by the YStrategyV3 contract.
type YStrategyV3ReportedIterator struct {
	Event *YStrategyV3Reported // Event containing the contract specifics and raw log

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
func (it *YStrategyV3ReportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YStrategyV3Reported)
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
		it.Event = new(YStrategyV3Reported)
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
func (it *YStrategyV3ReportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YStrategyV3ReportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YStrategyV3Reported represents a Reported event raised by the YStrategyV3 contract.
type YStrategyV3Reported struct {
	Profit          *big.Int
	Loss            *big.Int
	ProtocolFees    *big.Int
	PerformanceFees *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterReported is a free log retrieval operation binding the contract event 0xecdd072e4d5bd913a75a37f02daedcea7e2dc0281f9942c0063cfd1cfe5c4c4f.
//
// Solidity: event Reported(uint256 profit, uint256 loss, uint256 protocolFees, uint256 performanceFees)
func (_YStrategyV3 *YStrategyV3Filterer) FilterReported(opts *bind.FilterOpts) (*YStrategyV3ReportedIterator, error) {

	logs, sub, err := _YStrategyV3.contract.FilterLogs(opts, "Reported")
	if err != nil {
		return nil, err
	}
	return &YStrategyV3ReportedIterator{contract: _YStrategyV3.contract, event: "Reported", logs: logs, sub: sub}, nil
}

// WatchReported is a free log subscription operation binding the contract event 0xecdd072e4d5bd913a75a37f02daedcea7e2dc0281f9942c0063cfd1cfe5c4c4f.
//
// Solidity: event Reported(uint256 profit, uint256 loss, uint256 protocolFees, uint256 performanceFees)
func (_YStrategyV3 *YStrategyV3Filterer) WatchReported(opts *bind.WatchOpts, sink chan<- *YStrategyV3Reported) (event.Subscription, error) {

	logs, sub, err := _YStrategyV3.contract.WatchLogs(opts, "Reported")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YStrategyV3Reported)
				if err := _YStrategyV3.contract.UnpackLog(event, "Reported", log); err != nil {
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

// ParseReported is a log parse operation binding the contract event 0xecdd072e4d5bd913a75a37f02daedcea7e2dc0281f9942c0063cfd1cfe5c4c4f.
//
// Solidity: event Reported(uint256 profit, uint256 loss, uint256 protocolFees, uint256 performanceFees)
func (_YStrategyV3 *YStrategyV3Filterer) ParseReported(log types.Log) (*YStrategyV3Reported, error) {
	event := new(YStrategyV3Reported)
	if err := _YStrategyV3.contract.UnpackLog(event, "Reported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YStrategyV3StrategyShutdownIterator is returned from FilterStrategyShutdown and is used to iterate over the raw logs and unpacked data for StrategyShutdown events raised by the YStrategyV3 contract.
type YStrategyV3StrategyShutdownIterator struct {
	Event *YStrategyV3StrategyShutdown // Event containing the contract specifics and raw log

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
func (it *YStrategyV3StrategyShutdownIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YStrategyV3StrategyShutdown)
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
		it.Event = new(YStrategyV3StrategyShutdown)
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
func (it *YStrategyV3StrategyShutdownIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YStrategyV3StrategyShutdownIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YStrategyV3StrategyShutdown represents a StrategyShutdown event raised by the YStrategyV3 contract.
type YStrategyV3StrategyShutdown struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStrategyShutdown is a free log retrieval operation binding the contract event 0xfc1249757a7f27c510c8173c55d03ba442e0d33d9223e06ceb416feac8c7693f.
//
// Solidity: event StrategyShutdown()
func (_YStrategyV3 *YStrategyV3Filterer) FilterStrategyShutdown(opts *bind.FilterOpts) (*YStrategyV3StrategyShutdownIterator, error) {

	logs, sub, err := _YStrategyV3.contract.FilterLogs(opts, "StrategyShutdown")
	if err != nil {
		return nil, err
	}
	return &YStrategyV3StrategyShutdownIterator{contract: _YStrategyV3.contract, event: "StrategyShutdown", logs: logs, sub: sub}, nil
}

// WatchStrategyShutdown is a free log subscription operation binding the contract event 0xfc1249757a7f27c510c8173c55d03ba442e0d33d9223e06ceb416feac8c7693f.
//
// Solidity: event StrategyShutdown()
func (_YStrategyV3 *YStrategyV3Filterer) WatchStrategyShutdown(opts *bind.WatchOpts, sink chan<- *YStrategyV3StrategyShutdown) (event.Subscription, error) {

	logs, sub, err := _YStrategyV3.contract.WatchLogs(opts, "StrategyShutdown")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YStrategyV3StrategyShutdown)
				if err := _YStrategyV3.contract.UnpackLog(event, "StrategyShutdown", log); err != nil {
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

// ParseStrategyShutdown is a log parse operation binding the contract event 0xfc1249757a7f27c510c8173c55d03ba442e0d33d9223e06ceb416feac8c7693f.
//
// Solidity: event StrategyShutdown()
func (_YStrategyV3 *YStrategyV3Filterer) ParseStrategyShutdown(log types.Log) (*YStrategyV3StrategyShutdown, error) {
	event := new(YStrategyV3StrategyShutdown)
	if err := _YStrategyV3.contract.UnpackLog(event, "StrategyShutdown", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YStrategyV3TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the YStrategyV3 contract.
type YStrategyV3TransferIterator struct {
	Event *YStrategyV3Transfer // Event containing the contract specifics and raw log

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
func (it *YStrategyV3TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YStrategyV3Transfer)
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
		it.Event = new(YStrategyV3Transfer)
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
func (it *YStrategyV3TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YStrategyV3TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YStrategyV3Transfer represents a Transfer event raised by the YStrategyV3 contract.
type YStrategyV3Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_YStrategyV3 *YStrategyV3Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*YStrategyV3TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _YStrategyV3.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &YStrategyV3TransferIterator{contract: _YStrategyV3.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_YStrategyV3 *YStrategyV3Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *YStrategyV3Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _YStrategyV3.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YStrategyV3Transfer)
				if err := _YStrategyV3.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_YStrategyV3 *YStrategyV3Filterer) ParseTransfer(log types.Log) (*YStrategyV3Transfer, error) {
	event := new(YStrategyV3Transfer)
	if err := _YStrategyV3.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YStrategyV3UpdateKeeperIterator is returned from FilterUpdateKeeper and is used to iterate over the raw logs and unpacked data for UpdateKeeper events raised by the YStrategyV3 contract.
type YStrategyV3UpdateKeeperIterator struct {
	Event *YStrategyV3UpdateKeeper // Event containing the contract specifics and raw log

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
func (it *YStrategyV3UpdateKeeperIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YStrategyV3UpdateKeeper)
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
		it.Event = new(YStrategyV3UpdateKeeper)
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
func (it *YStrategyV3UpdateKeeperIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YStrategyV3UpdateKeeperIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YStrategyV3UpdateKeeper represents a UpdateKeeper event raised by the YStrategyV3 contract.
type YStrategyV3UpdateKeeper struct {
	NewKeeper common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateKeeper is a free log retrieval operation binding the contract event 0xd7f49e282c36d417b290d4181a56943f6d670aaa2987c0d40e60d39919c68882.
//
// Solidity: event UpdateKeeper(address indexed newKeeper)
func (_YStrategyV3 *YStrategyV3Filterer) FilterUpdateKeeper(opts *bind.FilterOpts, newKeeper []common.Address) (*YStrategyV3UpdateKeeperIterator, error) {

	var newKeeperRule []interface{}
	for _, newKeeperItem := range newKeeper {
		newKeeperRule = append(newKeeperRule, newKeeperItem)
	}

	logs, sub, err := _YStrategyV3.contract.FilterLogs(opts, "UpdateKeeper", newKeeperRule)
	if err != nil {
		return nil, err
	}
	return &YStrategyV3UpdateKeeperIterator{contract: _YStrategyV3.contract, event: "UpdateKeeper", logs: logs, sub: sub}, nil
}

// WatchUpdateKeeper is a free log subscription operation binding the contract event 0xd7f49e282c36d417b290d4181a56943f6d670aaa2987c0d40e60d39919c68882.
//
// Solidity: event UpdateKeeper(address indexed newKeeper)
func (_YStrategyV3 *YStrategyV3Filterer) WatchUpdateKeeper(opts *bind.WatchOpts, sink chan<- *YStrategyV3UpdateKeeper, newKeeper []common.Address) (event.Subscription, error) {

	var newKeeperRule []interface{}
	for _, newKeeperItem := range newKeeper {
		newKeeperRule = append(newKeeperRule, newKeeperItem)
	}

	logs, sub, err := _YStrategyV3.contract.WatchLogs(opts, "UpdateKeeper", newKeeperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YStrategyV3UpdateKeeper)
				if err := _YStrategyV3.contract.UnpackLog(event, "UpdateKeeper", log); err != nil {
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

// ParseUpdateKeeper is a log parse operation binding the contract event 0xd7f49e282c36d417b290d4181a56943f6d670aaa2987c0d40e60d39919c68882.
//
// Solidity: event UpdateKeeper(address indexed newKeeper)
func (_YStrategyV3 *YStrategyV3Filterer) ParseUpdateKeeper(log types.Log) (*YStrategyV3UpdateKeeper, error) {
	event := new(YStrategyV3UpdateKeeper)
	if err := _YStrategyV3.contract.UnpackLog(event, "UpdateKeeper", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YStrategyV3UpdateManagementIterator is returned from FilterUpdateManagement and is used to iterate over the raw logs and unpacked data for UpdateManagement events raised by the YStrategyV3 contract.
type YStrategyV3UpdateManagementIterator struct {
	Event *YStrategyV3UpdateManagement // Event containing the contract specifics and raw log

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
func (it *YStrategyV3UpdateManagementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YStrategyV3UpdateManagement)
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
		it.Event = new(YStrategyV3UpdateManagement)
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
func (it *YStrategyV3UpdateManagementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YStrategyV3UpdateManagementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YStrategyV3UpdateManagement represents a UpdateManagement event raised by the YStrategyV3 contract.
type YStrategyV3UpdateManagement struct {
	NewManagement common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdateManagement is a free log retrieval operation binding the contract event 0xff54978127edd34aec0f9061fb3b155fbe0ededdfa881ee3e0d541d3a1eef438.
//
// Solidity: event UpdateManagement(address indexed newManagement)
func (_YStrategyV3 *YStrategyV3Filterer) FilterUpdateManagement(opts *bind.FilterOpts, newManagement []common.Address) (*YStrategyV3UpdateManagementIterator, error) {

	var newManagementRule []interface{}
	for _, newManagementItem := range newManagement {
		newManagementRule = append(newManagementRule, newManagementItem)
	}

	logs, sub, err := _YStrategyV3.contract.FilterLogs(opts, "UpdateManagement", newManagementRule)
	if err != nil {
		return nil, err
	}
	return &YStrategyV3UpdateManagementIterator{contract: _YStrategyV3.contract, event: "UpdateManagement", logs: logs, sub: sub}, nil
}

// WatchUpdateManagement is a free log subscription operation binding the contract event 0xff54978127edd34aec0f9061fb3b155fbe0ededdfa881ee3e0d541d3a1eef438.
//
// Solidity: event UpdateManagement(address indexed newManagement)
func (_YStrategyV3 *YStrategyV3Filterer) WatchUpdateManagement(opts *bind.WatchOpts, sink chan<- *YStrategyV3UpdateManagement, newManagement []common.Address) (event.Subscription, error) {

	var newManagementRule []interface{}
	for _, newManagementItem := range newManagement {
		newManagementRule = append(newManagementRule, newManagementItem)
	}

	logs, sub, err := _YStrategyV3.contract.WatchLogs(opts, "UpdateManagement", newManagementRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YStrategyV3UpdateManagement)
				if err := _YStrategyV3.contract.UnpackLog(event, "UpdateManagement", log); err != nil {
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

// ParseUpdateManagement is a log parse operation binding the contract event 0xff54978127edd34aec0f9061fb3b155fbe0ededdfa881ee3e0d541d3a1eef438.
//
// Solidity: event UpdateManagement(address indexed newManagement)
func (_YStrategyV3 *YStrategyV3Filterer) ParseUpdateManagement(log types.Log) (*YStrategyV3UpdateManagement, error) {
	event := new(YStrategyV3UpdateManagement)
	if err := _YStrategyV3.contract.UnpackLog(event, "UpdateManagement", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YStrategyV3UpdatePendingManagementIterator is returned from FilterUpdatePendingManagement and is used to iterate over the raw logs and unpacked data for UpdatePendingManagement events raised by the YStrategyV3 contract.
type YStrategyV3UpdatePendingManagementIterator struct {
	Event *YStrategyV3UpdatePendingManagement // Event containing the contract specifics and raw log

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
func (it *YStrategyV3UpdatePendingManagementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YStrategyV3UpdatePendingManagement)
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
		it.Event = new(YStrategyV3UpdatePendingManagement)
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
func (it *YStrategyV3UpdatePendingManagementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YStrategyV3UpdatePendingManagementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YStrategyV3UpdatePendingManagement represents a UpdatePendingManagement event raised by the YStrategyV3 contract.
type YStrategyV3UpdatePendingManagement struct {
	NewPendingManagement common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdatePendingManagement is a free log retrieval operation binding the contract event 0xd74668a8c80a07cc56d7c3318a06439eaa815e740d97dcd83487e1fc75076b8b.
//
// Solidity: event UpdatePendingManagement(address indexed newPendingManagement)
func (_YStrategyV3 *YStrategyV3Filterer) FilterUpdatePendingManagement(opts *bind.FilterOpts, newPendingManagement []common.Address) (*YStrategyV3UpdatePendingManagementIterator, error) {

	var newPendingManagementRule []interface{}
	for _, newPendingManagementItem := range newPendingManagement {
		newPendingManagementRule = append(newPendingManagementRule, newPendingManagementItem)
	}

	logs, sub, err := _YStrategyV3.contract.FilterLogs(opts, "UpdatePendingManagement", newPendingManagementRule)
	if err != nil {
		return nil, err
	}
	return &YStrategyV3UpdatePendingManagementIterator{contract: _YStrategyV3.contract, event: "UpdatePendingManagement", logs: logs, sub: sub}, nil
}

// WatchUpdatePendingManagement is a free log subscription operation binding the contract event 0xd74668a8c80a07cc56d7c3318a06439eaa815e740d97dcd83487e1fc75076b8b.
//
// Solidity: event UpdatePendingManagement(address indexed newPendingManagement)
func (_YStrategyV3 *YStrategyV3Filterer) WatchUpdatePendingManagement(opts *bind.WatchOpts, sink chan<- *YStrategyV3UpdatePendingManagement, newPendingManagement []common.Address) (event.Subscription, error) {

	var newPendingManagementRule []interface{}
	for _, newPendingManagementItem := range newPendingManagement {
		newPendingManagementRule = append(newPendingManagementRule, newPendingManagementItem)
	}

	logs, sub, err := _YStrategyV3.contract.WatchLogs(opts, "UpdatePendingManagement", newPendingManagementRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YStrategyV3UpdatePendingManagement)
				if err := _YStrategyV3.contract.UnpackLog(event, "UpdatePendingManagement", log); err != nil {
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

// ParseUpdatePendingManagement is a log parse operation binding the contract event 0xd74668a8c80a07cc56d7c3318a06439eaa815e740d97dcd83487e1fc75076b8b.
//
// Solidity: event UpdatePendingManagement(address indexed newPendingManagement)
func (_YStrategyV3 *YStrategyV3Filterer) ParseUpdatePendingManagement(log types.Log) (*YStrategyV3UpdatePendingManagement, error) {
	event := new(YStrategyV3UpdatePendingManagement)
	if err := _YStrategyV3.contract.UnpackLog(event, "UpdatePendingManagement", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YStrategyV3UpdatePerformanceFeeIterator is returned from FilterUpdatePerformanceFee and is used to iterate over the raw logs and unpacked data for UpdatePerformanceFee events raised by the YStrategyV3 contract.
type YStrategyV3UpdatePerformanceFeeIterator struct {
	Event *YStrategyV3UpdatePerformanceFee // Event containing the contract specifics and raw log

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
func (it *YStrategyV3UpdatePerformanceFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YStrategyV3UpdatePerformanceFee)
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
		it.Event = new(YStrategyV3UpdatePerformanceFee)
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
func (it *YStrategyV3UpdatePerformanceFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YStrategyV3UpdatePerformanceFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YStrategyV3UpdatePerformanceFee represents a UpdatePerformanceFee event raised by the YStrategyV3 contract.
type YStrategyV3UpdatePerformanceFee struct {
	NewPerformanceFee uint16
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUpdatePerformanceFee is a free log retrieval operation binding the contract event 0xdc843735e683348ec21c302ffff45462399c5c46f75f67b0a1a5395c53599753.
//
// Solidity: event UpdatePerformanceFee(uint16 newPerformanceFee)
func (_YStrategyV3 *YStrategyV3Filterer) FilterUpdatePerformanceFee(opts *bind.FilterOpts) (*YStrategyV3UpdatePerformanceFeeIterator, error) {

	logs, sub, err := _YStrategyV3.contract.FilterLogs(opts, "UpdatePerformanceFee")
	if err != nil {
		return nil, err
	}
	return &YStrategyV3UpdatePerformanceFeeIterator{contract: _YStrategyV3.contract, event: "UpdatePerformanceFee", logs: logs, sub: sub}, nil
}

// WatchUpdatePerformanceFee is a free log subscription operation binding the contract event 0xdc843735e683348ec21c302ffff45462399c5c46f75f67b0a1a5395c53599753.
//
// Solidity: event UpdatePerformanceFee(uint16 newPerformanceFee)
func (_YStrategyV3 *YStrategyV3Filterer) WatchUpdatePerformanceFee(opts *bind.WatchOpts, sink chan<- *YStrategyV3UpdatePerformanceFee) (event.Subscription, error) {

	logs, sub, err := _YStrategyV3.contract.WatchLogs(opts, "UpdatePerformanceFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YStrategyV3UpdatePerformanceFee)
				if err := _YStrategyV3.contract.UnpackLog(event, "UpdatePerformanceFee", log); err != nil {
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

// ParseUpdatePerformanceFee is a log parse operation binding the contract event 0xdc843735e683348ec21c302ffff45462399c5c46f75f67b0a1a5395c53599753.
//
// Solidity: event UpdatePerformanceFee(uint16 newPerformanceFee)
func (_YStrategyV3 *YStrategyV3Filterer) ParseUpdatePerformanceFee(log types.Log) (*YStrategyV3UpdatePerformanceFee, error) {
	event := new(YStrategyV3UpdatePerformanceFee)
	if err := _YStrategyV3.contract.UnpackLog(event, "UpdatePerformanceFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YStrategyV3UpdatePerformanceFeeRecipientIterator is returned from FilterUpdatePerformanceFeeRecipient and is used to iterate over the raw logs and unpacked data for UpdatePerformanceFeeRecipient events raised by the YStrategyV3 contract.
type YStrategyV3UpdatePerformanceFeeRecipientIterator struct {
	Event *YStrategyV3UpdatePerformanceFeeRecipient // Event containing the contract specifics and raw log

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
func (it *YStrategyV3UpdatePerformanceFeeRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YStrategyV3UpdatePerformanceFeeRecipient)
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
		it.Event = new(YStrategyV3UpdatePerformanceFeeRecipient)
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
func (it *YStrategyV3UpdatePerformanceFeeRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YStrategyV3UpdatePerformanceFeeRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YStrategyV3UpdatePerformanceFeeRecipient represents a UpdatePerformanceFeeRecipient event raised by the YStrategyV3 contract.
type YStrategyV3UpdatePerformanceFeeRecipient struct {
	NewPerformanceFeeRecipient common.Address
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterUpdatePerformanceFeeRecipient is a free log retrieval operation binding the contract event 0x9ebbf695dd251e855d9d15a146a72f5f654dc6f8630fbc11212f27e0c88ba11a.
//
// Solidity: event UpdatePerformanceFeeRecipient(address indexed newPerformanceFeeRecipient)
func (_YStrategyV3 *YStrategyV3Filterer) FilterUpdatePerformanceFeeRecipient(opts *bind.FilterOpts, newPerformanceFeeRecipient []common.Address) (*YStrategyV3UpdatePerformanceFeeRecipientIterator, error) {

	var newPerformanceFeeRecipientRule []interface{}
	for _, newPerformanceFeeRecipientItem := range newPerformanceFeeRecipient {
		newPerformanceFeeRecipientRule = append(newPerformanceFeeRecipientRule, newPerformanceFeeRecipientItem)
	}

	logs, sub, err := _YStrategyV3.contract.FilterLogs(opts, "UpdatePerformanceFeeRecipient", newPerformanceFeeRecipientRule)
	if err != nil {
		return nil, err
	}
	return &YStrategyV3UpdatePerformanceFeeRecipientIterator{contract: _YStrategyV3.contract, event: "UpdatePerformanceFeeRecipient", logs: logs, sub: sub}, nil
}

// WatchUpdatePerformanceFeeRecipient is a free log subscription operation binding the contract event 0x9ebbf695dd251e855d9d15a146a72f5f654dc6f8630fbc11212f27e0c88ba11a.
//
// Solidity: event UpdatePerformanceFeeRecipient(address indexed newPerformanceFeeRecipient)
func (_YStrategyV3 *YStrategyV3Filterer) WatchUpdatePerformanceFeeRecipient(opts *bind.WatchOpts, sink chan<- *YStrategyV3UpdatePerformanceFeeRecipient, newPerformanceFeeRecipient []common.Address) (event.Subscription, error) {

	var newPerformanceFeeRecipientRule []interface{}
	for _, newPerformanceFeeRecipientItem := range newPerformanceFeeRecipient {
		newPerformanceFeeRecipientRule = append(newPerformanceFeeRecipientRule, newPerformanceFeeRecipientItem)
	}

	logs, sub, err := _YStrategyV3.contract.WatchLogs(opts, "UpdatePerformanceFeeRecipient", newPerformanceFeeRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YStrategyV3UpdatePerformanceFeeRecipient)
				if err := _YStrategyV3.contract.UnpackLog(event, "UpdatePerformanceFeeRecipient", log); err != nil {
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

// ParseUpdatePerformanceFeeRecipient is a log parse operation binding the contract event 0x9ebbf695dd251e855d9d15a146a72f5f654dc6f8630fbc11212f27e0c88ba11a.
//
// Solidity: event UpdatePerformanceFeeRecipient(address indexed newPerformanceFeeRecipient)
func (_YStrategyV3 *YStrategyV3Filterer) ParseUpdatePerformanceFeeRecipient(log types.Log) (*YStrategyV3UpdatePerformanceFeeRecipient, error) {
	event := new(YStrategyV3UpdatePerformanceFeeRecipient)
	if err := _YStrategyV3.contract.UnpackLog(event, "UpdatePerformanceFeeRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YStrategyV3UpdateProfitMaxUnlockTimeIterator is returned from FilterUpdateProfitMaxUnlockTime and is used to iterate over the raw logs and unpacked data for UpdateProfitMaxUnlockTime events raised by the YStrategyV3 contract.
type YStrategyV3UpdateProfitMaxUnlockTimeIterator struct {
	Event *YStrategyV3UpdateProfitMaxUnlockTime // Event containing the contract specifics and raw log

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
func (it *YStrategyV3UpdateProfitMaxUnlockTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YStrategyV3UpdateProfitMaxUnlockTime)
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
		it.Event = new(YStrategyV3UpdateProfitMaxUnlockTime)
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
func (it *YStrategyV3UpdateProfitMaxUnlockTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YStrategyV3UpdateProfitMaxUnlockTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YStrategyV3UpdateProfitMaxUnlockTime represents a UpdateProfitMaxUnlockTime event raised by the YStrategyV3 contract.
type YStrategyV3UpdateProfitMaxUnlockTime struct {
	NewProfitMaxUnlockTime *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterUpdateProfitMaxUnlockTime is a free log retrieval operation binding the contract event 0xf361aed463da6fa20358e45c6209f1d3e16d4eca706e6eab0b0aeb338729c77a.
//
// Solidity: event UpdateProfitMaxUnlockTime(uint256 newProfitMaxUnlockTime)
func (_YStrategyV3 *YStrategyV3Filterer) FilterUpdateProfitMaxUnlockTime(opts *bind.FilterOpts) (*YStrategyV3UpdateProfitMaxUnlockTimeIterator, error) {

	logs, sub, err := _YStrategyV3.contract.FilterLogs(opts, "UpdateProfitMaxUnlockTime")
	if err != nil {
		return nil, err
	}
	return &YStrategyV3UpdateProfitMaxUnlockTimeIterator{contract: _YStrategyV3.contract, event: "UpdateProfitMaxUnlockTime", logs: logs, sub: sub}, nil
}

// WatchUpdateProfitMaxUnlockTime is a free log subscription operation binding the contract event 0xf361aed463da6fa20358e45c6209f1d3e16d4eca706e6eab0b0aeb338729c77a.
//
// Solidity: event UpdateProfitMaxUnlockTime(uint256 newProfitMaxUnlockTime)
func (_YStrategyV3 *YStrategyV3Filterer) WatchUpdateProfitMaxUnlockTime(opts *bind.WatchOpts, sink chan<- *YStrategyV3UpdateProfitMaxUnlockTime) (event.Subscription, error) {

	logs, sub, err := _YStrategyV3.contract.WatchLogs(opts, "UpdateProfitMaxUnlockTime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YStrategyV3UpdateProfitMaxUnlockTime)
				if err := _YStrategyV3.contract.UnpackLog(event, "UpdateProfitMaxUnlockTime", log); err != nil {
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

// ParseUpdateProfitMaxUnlockTime is a log parse operation binding the contract event 0xf361aed463da6fa20358e45c6209f1d3e16d4eca706e6eab0b0aeb338729c77a.
//
// Solidity: event UpdateProfitMaxUnlockTime(uint256 newProfitMaxUnlockTime)
func (_YStrategyV3 *YStrategyV3Filterer) ParseUpdateProfitMaxUnlockTime(log types.Log) (*YStrategyV3UpdateProfitMaxUnlockTime, error) {
	event := new(YStrategyV3UpdateProfitMaxUnlockTime)
	if err := _YStrategyV3.contract.UnpackLog(event, "UpdateProfitMaxUnlockTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YStrategyV3WithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the YStrategyV3 contract.
type YStrategyV3WithdrawIterator struct {
	Event *YStrategyV3Withdraw // Event containing the contract specifics and raw log

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
func (it *YStrategyV3WithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YStrategyV3Withdraw)
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
		it.Event = new(YStrategyV3Withdraw)
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
func (it *YStrategyV3WithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YStrategyV3WithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YStrategyV3Withdraw represents a Withdraw event raised by the YStrategyV3 contract.
type YStrategyV3Withdraw struct {
	Caller   common.Address
	Receiver common.Address
	Owner    common.Address
	Assets   *big.Int
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed caller, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_YStrategyV3 *YStrategyV3Filterer) FilterWithdraw(opts *bind.FilterOpts, caller []common.Address, receiver []common.Address, owner []common.Address) (*YStrategyV3WithdrawIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _YStrategyV3.contract.FilterLogs(opts, "Withdraw", callerRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &YStrategyV3WithdrawIterator{contract: _YStrategyV3.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed caller, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_YStrategyV3 *YStrategyV3Filterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *YStrategyV3Withdraw, caller []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _YStrategyV3.contract.WatchLogs(opts, "Withdraw", callerRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YStrategyV3Withdraw)
				if err := _YStrategyV3.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed caller, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_YStrategyV3 *YStrategyV3Filterer) ParseWithdraw(log types.Log) (*YStrategyV3Withdraw, error) {
	event := new(YStrategyV3Withdraw)
	if err := _YStrategyV3.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
