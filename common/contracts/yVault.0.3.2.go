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

// Yvault032MetaData contains all meta data concerning the Yvault032 contract.
var Yvault032MetaData = &bind.MetaData{
	ABI: "[{\"name\":\"Transfer\",\"inputs\":[{\"type\":\"address\",\"name\":\"sender\",\"indexed\":true},{\"type\":\"address\",\"name\":\"receiver\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Approval\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"indexed\":true},{\"type\":\"address\",\"name\":\"spender\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StrategyAdded\",\"inputs\":[{\"type\":\"address\",\"name\":\"strategy\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"debtRatio\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"minDebtPerHarvest\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"maxDebtPerHarvest\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"performanceFee\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StrategyReported\",\"inputs\":[{\"type\":\"address\",\"name\":\"strategy\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"gain\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"loss\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"debtPaid\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"totalGain\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"totalLoss\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"totalDebt\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"debtAdded\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"debtRatio\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateGovernance\",\"inputs\":[{\"type\":\"address\",\"name\":\"governance\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateManagement\",\"inputs\":[{\"type\":\"address\",\"name\":\"management\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateGuestList\",\"inputs\":[{\"type\":\"address\",\"name\":\"guestList\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateRewards\",\"inputs\":[{\"type\":\"address\",\"name\":\"rewards\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateDepositLimit\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"depositLimit\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdatePerformanceFee\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"performanceFee\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateManagementFee\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"managementFee\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateGuardian\",\"inputs\":[{\"type\":\"address\",\"name\":\"guardian\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"EmergencyShutdown\",\"inputs\":[{\"type\":\"bool\",\"name\":\"active\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateWithdrawalQueue\",\"inputs\":[{\"type\":\"address[20]\",\"name\":\"queue\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StrategyUpdateDebtRatio\",\"inputs\":[{\"type\":\"address\",\"name\":\"strategy\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"debtRatio\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StrategyUpdateMinDebtPerHarvest\",\"inputs\":[{\"type\":\"address\",\"name\":\"strategy\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"minDebtPerHarvest\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StrategyUpdateMaxDebtPerHarvest\",\"inputs\":[{\"type\":\"address\",\"name\":\"strategy\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"maxDebtPerHarvest\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StrategyUpdatePerformanceFee\",\"inputs\":[{\"type\":\"address\",\"name\":\"strategy\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"performanceFee\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StrategyMigrated\",\"inputs\":[{\"type\":\"address\",\"name\":\"oldVersion\",\"indexed\":true},{\"type\":\"address\",\"name\":\"newVersion\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StrategyRevoked\",\"inputs\":[{\"type\":\"address\",\"name\":\"strategy\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StrategyRemovedFromQueue\",\"inputs\":[{\"type\":\"address\",\"name\":\"strategy\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StrategyAddedToQueue\",\"inputs\":[{\"type\":\"address\",\"name\":\"strategy\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"initialize\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"token\"},{\"type\":\"address\",\"name\":\"governance\"},{\"type\":\"address\",\"name\":\"rewards\"},{\"type\":\"string\",\"name\":\"nameOverride\"},{\"type\":\"string\",\"name\":\"symbolOverride\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"initialize\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"token\"},{\"type\":\"address\",\"name\":\"governance\"},{\"type\":\"address\",\"name\":\"rewards\"},{\"type\":\"string\",\"name\":\"nameOverride\"},{\"type\":\"string\",\"name\":\"symbolOverride\"},{\"type\":\"address\",\"name\":\"guardian\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"apiVersion\",\"outputs\":[{\"type\":\"string\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\",\"gas\":4519},{\"name\":\"setName\",\"outputs\":[],\"inputs\":[{\"type\":\"string\",\"name\":\"name\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":107017},{\"name\":\"setSymbol\",\"outputs\":[],\"inputs\":[{\"type\":\"string\",\"name\":\"symbol\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":71867},{\"name\":\"setGovernance\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"governance\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":36338},{\"name\":\"acceptGovernance\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37610},{\"name\":\"setManagement\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"management\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37748},{\"name\":\"setGuestList\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"guestList\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37778},{\"name\":\"setRewards\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"rewards\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37808},{\"name\":\"setLockedProfitDegration\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"degration\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":36516},{\"name\":\"setDepositLimit\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"limit\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37768},{\"name\":\"setPerformanceFee\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"fee\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37902},{\"name\":\"setManagementFee\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"fee\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37932},{\"name\":\"setGuardian\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"guardian\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":39176},{\"name\":\"setEmergencyShutdown\",\"outputs\":[],\"inputs\":[{\"type\":\"bool\",\"name\":\"active\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":39247},{\"name\":\"setWithdrawalQueue\",\"outputs\":[],\"inputs\":[{\"type\":\"address[20]\",\"name\":\"queue\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":763923},{\"name\":\"transfer\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"receiver\"},{\"type\":\"uint256\",\"name\":\"amount\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":76913},{\"name\":\"transferFrom\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"sender\"},{\"type\":\"address\",\"name\":\"receiver\"},{\"type\":\"uint256\",\"name\":\"amount\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":116676},{\"name\":\"approve\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"spender\"},{\"type\":\"uint256\",\"name\":\"amount\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":38334},{\"name\":\"increaseAllowance\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"spender\"},{\"type\":\"uint256\",\"name\":\"amount\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":40375},{\"name\":\"decreaseAllowance\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"spender\"},{\"type\":\"uint256\",\"name\":\"amount\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":40399},{\"name\":\"permit\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"owner\"},{\"type\":\"address\",\"name\":\"spender\"},{\"type\":\"uint256\",\"name\":\"amount\"},{\"type\":\"uint256\",\"name\":\"expiry\"},{\"type\":\"bytes\",\"name\":\"signature\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":81327},{\"name\":\"totalAssets\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":4303},{\"name\":\"deposit\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"deposit\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_amount\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"deposit\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_amount\"},{\"type\":\"address\",\"name\":\"recipient\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"maxAvailableShares\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":379843},{\"name\":\"withdraw\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"withdraw\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"maxShares\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"withdraw\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"maxShares\"},{\"type\":\"address\",\"name\":\"recipient\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"withdraw\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"maxShares\"},{\"type\":\"address\",\"name\":\"recipient\"},{\"type\":\"uint256\",\"name\":\"maxLoss\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"pricePerShare\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":17509},{\"name\":\"addStrategy\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"strategy\"},{\"type\":\"uint256\",\"name\":\"debtRatio\"},{\"type\":\"uint256\",\"name\":\"minDebtPerHarvest\"},{\"type\":\"uint256\",\"name\":\"maxDebtPerHarvest\"},{\"type\":\"uint256\",\"name\":\"performanceFee\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":1486241},{\"name\":\"updateStrategyDebtRatio\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"strategy\"},{\"type\":\"uint256\",\"name\":\"debtRatio\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":115406},{\"name\":\"updateStrategyMinDebtPerHarvest\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"strategy\"},{\"type\":\"uint256\",\"name\":\"minDebtPerHarvest\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":42654},{\"name\":\"updateStrategyMaxDebtPerHarvest\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"strategy\"},{\"type\":\"uint256\",\"name\":\"maxDebtPerHarvest\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":42684},{\"name\":\"updateStrategyPerformanceFee\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"strategy\"},{\"type\":\"uint256\",\"name\":\"performanceFee\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":41464},{\"name\":\"migrateStrategy\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"oldVersion\"},{\"type\":\"address\",\"name\":\"newVersion\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":1141973},{\"name\":\"revokeStrategy\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"revokeStrategy\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"strategy\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"addStrategyToQueue\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"strategy\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":1197130},{\"name\":\"removeStrategyFromQueue\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"strategy\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":23093586},{\"name\":\"debtOutstanding\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"name\":\"debtOutstanding\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"strategy\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"name\":\"creditAvailable\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"name\":\"creditAvailable\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"strategy\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"name\":\"availableDepositLimit\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":10108},{\"name\":\"expectedReturn\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"name\":\"expectedReturn\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"strategy\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"name\":\"report\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"gain\"},{\"type\":\"uint256\",\"name\":\"loss\"},{\"type\":\"uint256\",\"name\":\"_debtPayment\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":1009335},{\"name\":\"sweep\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"token\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"sweep\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"token\"},{\"type\":\"uint256\",\"name\":\"amount\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"name\",\"outputs\":[{\"type\":\"string\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":9143},{\"name\":\"symbol\",\"outputs\":[{\"type\":\"string\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":8196},{\"name\":\"decimals\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2801},{\"name\":\"balanceOf\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3046},{\"name\":\"allowance\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"},{\"type\":\"address\",\"name\":\"arg1\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3291},{\"name\":\"totalSupply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2891},{\"name\":\"token\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2921},{\"name\":\"governance\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2951},{\"name\":\"management\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2981},{\"name\":\"guardian\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3011},{\"name\":\"guestList\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3041},{\"name\":\"strategies\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"performanceFee\"},{\"type\":\"uint256\",\"name\":\"activation\"},{\"type\":\"uint256\",\"name\":\"debtRatio\"},{\"type\":\"uint256\",\"name\":\"minDebtPerHarvest\"},{\"type\":\"uint256\",\"name\":\"maxDebtPerHarvest\"},{\"type\":\"uint256\",\"name\":\"lastReport\"},{\"type\":\"uint256\",\"name\":\"totalDebt\"},{\"type\":\"uint256\",\"name\":\"totalGain\"},{\"type\":\"uint256\",\"name\":\"totalLoss\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":11394},{\"name\":\"withdrawalQueue\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3210},{\"name\":\"emergencyShutdown\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3131},{\"name\":\"depositLimit\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3161},{\"name\":\"debtRatio\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3191},{\"name\":\"totalDebt\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3221},{\"name\":\"lastReport\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3251},{\"name\":\"activation\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3281},{\"name\":\"lockedProfit\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3311},{\"name\":\"lockedProfitDegration\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3341},{\"name\":\"rewards\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3371},{\"name\":\"managementFee\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3401},{\"name\":\"performanceFee\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3431},{\"name\":\"nonces\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3676},{\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3491}]",
}

// Yvault032ABI is the input ABI used to generate the binding from.
// Deprecated: Use Yvault032MetaData.ABI instead.
var Yvault032ABI = Yvault032MetaData.ABI

// Yvault032 is an auto generated Go binding around an Ethereum contract.
type Yvault032 struct {
	Yvault032Caller     // Read-only binding to the contract
	Yvault032Transactor // Write-only binding to the contract
	Yvault032Filterer   // Log filterer for contract events
}

// Yvault032Caller is an auto generated read-only Go binding around an Ethereum contract.
type Yvault032Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Yvault032Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Yvault032Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Yvault032Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Yvault032Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Yvault032Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Yvault032Session struct {
	Contract     *Yvault032        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Yvault032CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Yvault032CallerSession struct {
	Contract *Yvault032Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// Yvault032TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Yvault032TransactorSession struct {
	Contract     *Yvault032Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// Yvault032Raw is an auto generated low-level Go binding around an Ethereum contract.
type Yvault032Raw struct {
	Contract *Yvault032 // Generic contract binding to access the raw methods on
}

// Yvault032CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Yvault032CallerRaw struct {
	Contract *Yvault032Caller // Generic read-only contract binding to access the raw methods on
}

// Yvault032TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Yvault032TransactorRaw struct {
	Contract *Yvault032Transactor // Generic write-only contract binding to access the raw methods on
}

// NewYvault032 creates a new instance of Yvault032, bound to a specific deployed contract.
func NewYvault032(address common.Address, backend bind.ContractBackend) (*Yvault032, error) {
	contract, err := bindYvault032(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Yvault032{Yvault032Caller: Yvault032Caller{contract: contract}, Yvault032Transactor: Yvault032Transactor{contract: contract}, Yvault032Filterer: Yvault032Filterer{contract: contract}}, nil
}

// NewYvault032Caller creates a new read-only instance of Yvault032, bound to a specific deployed contract.
func NewYvault032Caller(address common.Address, caller bind.ContractCaller) (*Yvault032Caller, error) {
	contract, err := bindYvault032(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Yvault032Caller{contract: contract}, nil
}

// NewYvault032Transactor creates a new write-only instance of Yvault032, bound to a specific deployed contract.
func NewYvault032Transactor(address common.Address, transactor bind.ContractTransactor) (*Yvault032Transactor, error) {
	contract, err := bindYvault032(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Yvault032Transactor{contract: contract}, nil
}

// NewYvault032Filterer creates a new log filterer instance of Yvault032, bound to a specific deployed contract.
func NewYvault032Filterer(address common.Address, filterer bind.ContractFilterer) (*Yvault032Filterer, error) {
	contract, err := bindYvault032(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Yvault032Filterer{contract: contract}, nil
}

// bindYvault032 binds a generic wrapper to an already deployed contract.
func bindYvault032(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Yvault032ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Yvault032 *Yvault032Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Yvault032.Contract.Yvault032Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Yvault032 *Yvault032Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yvault032.Contract.Yvault032Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Yvault032 *Yvault032Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Yvault032.Contract.Yvault032Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Yvault032 *Yvault032CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Yvault032.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Yvault032 *Yvault032TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yvault032.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Yvault032 *Yvault032TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Yvault032.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Yvault032 *Yvault032Caller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Yvault032.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Yvault032 *Yvault032Session) DOMAINSEPARATOR() ([32]byte, error) {
	return _Yvault032.Contract.DOMAINSEPARATOR(&_Yvault032.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Yvault032 *Yvault032CallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Yvault032.Contract.DOMAINSEPARATOR(&_Yvault032.CallOpts)
}

// Activation is a free data retrieval call binding the contract method 0x3629c8de.
//
// Solidity: function activation() view returns(uint256)
func (_Yvault032 *Yvault032Caller) Activation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault032.contract.Call(opts, &out, "activation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Activation is a free data retrieval call binding the contract method 0x3629c8de.
//
// Solidity: function activation() view returns(uint256)
func (_Yvault032 *Yvault032Session) Activation() (*big.Int, error) {
	return _Yvault032.Contract.Activation(&_Yvault032.CallOpts)
}

// Activation is a free data retrieval call binding the contract method 0x3629c8de.
//
// Solidity: function activation() view returns(uint256)
func (_Yvault032 *Yvault032CallerSession) Activation() (*big.Int, error) {
	return _Yvault032.Contract.Activation(&_Yvault032.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_Yvault032 *Yvault032Caller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yvault032.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_Yvault032 *Yvault032Session) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Yvault032.Contract.Allowance(&_Yvault032.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_Yvault032 *Yvault032CallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Yvault032.Contract.Allowance(&_Yvault032.CallOpts, arg0, arg1)
}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() pure returns(string)
func (_Yvault032 *Yvault032Caller) ApiVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Yvault032.contract.Call(opts, &out, "apiVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() pure returns(string)
func (_Yvault032 *Yvault032Session) ApiVersion() (string, error) {
	return _Yvault032.Contract.ApiVersion(&_Yvault032.CallOpts)
}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() pure returns(string)
func (_Yvault032 *Yvault032CallerSession) ApiVersion() (string, error) {
	return _Yvault032.Contract.ApiVersion(&_Yvault032.CallOpts)
}

// AvailableDepositLimit is a free data retrieval call binding the contract method 0x153c27c4.
//
// Solidity: function availableDepositLimit() view returns(uint256)
func (_Yvault032 *Yvault032Caller) AvailableDepositLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault032.contract.Call(opts, &out, "availableDepositLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailableDepositLimit is a free data retrieval call binding the contract method 0x153c27c4.
//
// Solidity: function availableDepositLimit() view returns(uint256)
func (_Yvault032 *Yvault032Session) AvailableDepositLimit() (*big.Int, error) {
	return _Yvault032.Contract.AvailableDepositLimit(&_Yvault032.CallOpts)
}

// AvailableDepositLimit is a free data retrieval call binding the contract method 0x153c27c4.
//
// Solidity: function availableDepositLimit() view returns(uint256)
func (_Yvault032 *Yvault032CallerSession) AvailableDepositLimit() (*big.Int, error) {
	return _Yvault032.Contract.AvailableDepositLimit(&_Yvault032.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_Yvault032 *Yvault032Caller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yvault032.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_Yvault032 *Yvault032Session) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Yvault032.Contract.BalanceOf(&_Yvault032.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_Yvault032 *Yvault032CallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Yvault032.Contract.BalanceOf(&_Yvault032.CallOpts, arg0)
}

// CreditAvailable is a free data retrieval call binding the contract method 0x112c1f9b.
//
// Solidity: function creditAvailable() view returns(uint256)
func (_Yvault032 *Yvault032Caller) CreditAvailable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault032.contract.Call(opts, &out, "creditAvailable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreditAvailable is a free data retrieval call binding the contract method 0x112c1f9b.
//
// Solidity: function creditAvailable() view returns(uint256)
func (_Yvault032 *Yvault032Session) CreditAvailable() (*big.Int, error) {
	return _Yvault032.Contract.CreditAvailable(&_Yvault032.CallOpts)
}

// CreditAvailable is a free data retrieval call binding the contract method 0x112c1f9b.
//
// Solidity: function creditAvailable() view returns(uint256)
func (_Yvault032 *Yvault032CallerSession) CreditAvailable() (*big.Int, error) {
	return _Yvault032.Contract.CreditAvailable(&_Yvault032.CallOpts)
}

// CreditAvailable0 is a free data retrieval call binding the contract method 0xd7648013.
//
// Solidity: function creditAvailable(address strategy) view returns(uint256)
func (_Yvault032 *Yvault032Caller) CreditAvailable0(opts *bind.CallOpts, strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yvault032.contract.Call(opts, &out, "creditAvailable0", strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreditAvailable0 is a free data retrieval call binding the contract method 0xd7648013.
//
// Solidity: function creditAvailable(address strategy) view returns(uint256)
func (_Yvault032 *Yvault032Session) CreditAvailable0(strategy common.Address) (*big.Int, error) {
	return _Yvault032.Contract.CreditAvailable0(&_Yvault032.CallOpts, strategy)
}

// CreditAvailable0 is a free data retrieval call binding the contract method 0xd7648013.
//
// Solidity: function creditAvailable(address strategy) view returns(uint256)
func (_Yvault032 *Yvault032CallerSession) CreditAvailable0(strategy common.Address) (*big.Int, error) {
	return _Yvault032.Contract.CreditAvailable0(&_Yvault032.CallOpts, strategy)
}

// DebtOutstanding is a free data retrieval call binding the contract method 0xbf3759b5.
//
// Solidity: function debtOutstanding() view returns(uint256)
func (_Yvault032 *Yvault032Caller) DebtOutstanding(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault032.contract.Call(opts, &out, "debtOutstanding")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtOutstanding is a free data retrieval call binding the contract method 0xbf3759b5.
//
// Solidity: function debtOutstanding() view returns(uint256)
func (_Yvault032 *Yvault032Session) DebtOutstanding() (*big.Int, error) {
	return _Yvault032.Contract.DebtOutstanding(&_Yvault032.CallOpts)
}

// DebtOutstanding is a free data retrieval call binding the contract method 0xbf3759b5.
//
// Solidity: function debtOutstanding() view returns(uint256)
func (_Yvault032 *Yvault032CallerSession) DebtOutstanding() (*big.Int, error) {
	return _Yvault032.Contract.DebtOutstanding(&_Yvault032.CallOpts)
}

// DebtOutstanding0 is a free data retrieval call binding the contract method 0xbdcf36bb.
//
// Solidity: function debtOutstanding(address strategy) view returns(uint256)
func (_Yvault032 *Yvault032Caller) DebtOutstanding0(opts *bind.CallOpts, strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yvault032.contract.Call(opts, &out, "debtOutstanding0", strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtOutstanding0 is a free data retrieval call binding the contract method 0xbdcf36bb.
//
// Solidity: function debtOutstanding(address strategy) view returns(uint256)
func (_Yvault032 *Yvault032Session) DebtOutstanding0(strategy common.Address) (*big.Int, error) {
	return _Yvault032.Contract.DebtOutstanding0(&_Yvault032.CallOpts, strategy)
}

// DebtOutstanding0 is a free data retrieval call binding the contract method 0xbdcf36bb.
//
// Solidity: function debtOutstanding(address strategy) view returns(uint256)
func (_Yvault032 *Yvault032CallerSession) DebtOutstanding0(strategy common.Address) (*big.Int, error) {
	return _Yvault032.Contract.DebtOutstanding0(&_Yvault032.CallOpts, strategy)
}

// DebtRatio is a free data retrieval call binding the contract method 0xcea55f57.
//
// Solidity: function debtRatio() view returns(uint256)
func (_Yvault032 *Yvault032Caller) DebtRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault032.contract.Call(opts, &out, "debtRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtRatio is a free data retrieval call binding the contract method 0xcea55f57.
//
// Solidity: function debtRatio() view returns(uint256)
func (_Yvault032 *Yvault032Session) DebtRatio() (*big.Int, error) {
	return _Yvault032.Contract.DebtRatio(&_Yvault032.CallOpts)
}

// DebtRatio is a free data retrieval call binding the contract method 0xcea55f57.
//
// Solidity: function debtRatio() view returns(uint256)
func (_Yvault032 *Yvault032CallerSession) DebtRatio() (*big.Int, error) {
	return _Yvault032.Contract.DebtRatio(&_Yvault032.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Yvault032 *Yvault032Caller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault032.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Yvault032 *Yvault032Session) Decimals() (*big.Int, error) {
	return _Yvault032.Contract.Decimals(&_Yvault032.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Yvault032 *Yvault032CallerSession) Decimals() (*big.Int, error) {
	return _Yvault032.Contract.Decimals(&_Yvault032.CallOpts)
}

// DepositLimit is a free data retrieval call binding the contract method 0xecf70858.
//
// Solidity: function depositLimit() view returns(uint256)
func (_Yvault032 *Yvault032Caller) DepositLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault032.contract.Call(opts, &out, "depositLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositLimit is a free data retrieval call binding the contract method 0xecf70858.
//
// Solidity: function depositLimit() view returns(uint256)
func (_Yvault032 *Yvault032Session) DepositLimit() (*big.Int, error) {
	return _Yvault032.Contract.DepositLimit(&_Yvault032.CallOpts)
}

// DepositLimit is a free data retrieval call binding the contract method 0xecf70858.
//
// Solidity: function depositLimit() view returns(uint256)
func (_Yvault032 *Yvault032CallerSession) DepositLimit() (*big.Int, error) {
	return _Yvault032.Contract.DepositLimit(&_Yvault032.CallOpts)
}

// EmergencyShutdown is a free data retrieval call binding the contract method 0x3403c2fc.
//
// Solidity: function emergencyShutdown() view returns(bool)
func (_Yvault032 *Yvault032Caller) EmergencyShutdown(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Yvault032.contract.Call(opts, &out, "emergencyShutdown")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EmergencyShutdown is a free data retrieval call binding the contract method 0x3403c2fc.
//
// Solidity: function emergencyShutdown() view returns(bool)
func (_Yvault032 *Yvault032Session) EmergencyShutdown() (bool, error) {
	return _Yvault032.Contract.EmergencyShutdown(&_Yvault032.CallOpts)
}

// EmergencyShutdown is a free data retrieval call binding the contract method 0x3403c2fc.
//
// Solidity: function emergencyShutdown() view returns(bool)
func (_Yvault032 *Yvault032CallerSession) EmergencyShutdown() (bool, error) {
	return _Yvault032.Contract.EmergencyShutdown(&_Yvault032.CallOpts)
}

// ExpectedReturn is a free data retrieval call binding the contract method 0xd3406abd.
//
// Solidity: function expectedReturn() view returns(uint256)
func (_Yvault032 *Yvault032Caller) ExpectedReturn(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault032.contract.Call(opts, &out, "expectedReturn")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpectedReturn is a free data retrieval call binding the contract method 0xd3406abd.
//
// Solidity: function expectedReturn() view returns(uint256)
func (_Yvault032 *Yvault032Session) ExpectedReturn() (*big.Int, error) {
	return _Yvault032.Contract.ExpectedReturn(&_Yvault032.CallOpts)
}

// ExpectedReturn is a free data retrieval call binding the contract method 0xd3406abd.
//
// Solidity: function expectedReturn() view returns(uint256)
func (_Yvault032 *Yvault032CallerSession) ExpectedReturn() (*big.Int, error) {
	return _Yvault032.Contract.ExpectedReturn(&_Yvault032.CallOpts)
}

// ExpectedReturn0 is a free data retrieval call binding the contract method 0x33586b67.
//
// Solidity: function expectedReturn(address strategy) view returns(uint256)
func (_Yvault032 *Yvault032Caller) ExpectedReturn0(opts *bind.CallOpts, strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yvault032.contract.Call(opts, &out, "expectedReturn0", strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpectedReturn0 is a free data retrieval call binding the contract method 0x33586b67.
//
// Solidity: function expectedReturn(address strategy) view returns(uint256)
func (_Yvault032 *Yvault032Session) ExpectedReturn0(strategy common.Address) (*big.Int, error) {
	return _Yvault032.Contract.ExpectedReturn0(&_Yvault032.CallOpts, strategy)
}

// ExpectedReturn0 is a free data retrieval call binding the contract method 0x33586b67.
//
// Solidity: function expectedReturn(address strategy) view returns(uint256)
func (_Yvault032 *Yvault032CallerSession) ExpectedReturn0(strategy common.Address) (*big.Int, error) {
	return _Yvault032.Contract.ExpectedReturn0(&_Yvault032.CallOpts, strategy)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Yvault032 *Yvault032Caller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yvault032.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Yvault032 *Yvault032Session) Governance() (common.Address, error) {
	return _Yvault032.Contract.Governance(&_Yvault032.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Yvault032 *Yvault032CallerSession) Governance() (common.Address, error) {
	return _Yvault032.Contract.Governance(&_Yvault032.CallOpts)
}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_Yvault032 *Yvault032Caller) Guardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yvault032.contract.Call(opts, &out, "guardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_Yvault032 *Yvault032Session) Guardian() (common.Address, error) {
	return _Yvault032.Contract.Guardian(&_Yvault032.CallOpts)
}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_Yvault032 *Yvault032CallerSession) Guardian() (common.Address, error) {
	return _Yvault032.Contract.Guardian(&_Yvault032.CallOpts)
}

// GuestList is a free data retrieval call binding the contract method 0x46d55875.
//
// Solidity: function guestList() view returns(address)
func (_Yvault032 *Yvault032Caller) GuestList(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yvault032.contract.Call(opts, &out, "guestList")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GuestList is a free data retrieval call binding the contract method 0x46d55875.
//
// Solidity: function guestList() view returns(address)
func (_Yvault032 *Yvault032Session) GuestList() (common.Address, error) {
	return _Yvault032.Contract.GuestList(&_Yvault032.CallOpts)
}

// GuestList is a free data retrieval call binding the contract method 0x46d55875.
//
// Solidity: function guestList() view returns(address)
func (_Yvault032 *Yvault032CallerSession) GuestList() (common.Address, error) {
	return _Yvault032.Contract.GuestList(&_Yvault032.CallOpts)
}

// LastReport is a free data retrieval call binding the contract method 0xc3535b52.
//
// Solidity: function lastReport() view returns(uint256)
func (_Yvault032 *Yvault032Caller) LastReport(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault032.contract.Call(opts, &out, "lastReport")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastReport is a free data retrieval call binding the contract method 0xc3535b52.
//
// Solidity: function lastReport() view returns(uint256)
func (_Yvault032 *Yvault032Session) LastReport() (*big.Int, error) {
	return _Yvault032.Contract.LastReport(&_Yvault032.CallOpts)
}

// LastReport is a free data retrieval call binding the contract method 0xc3535b52.
//
// Solidity: function lastReport() view returns(uint256)
func (_Yvault032 *Yvault032CallerSession) LastReport() (*big.Int, error) {
	return _Yvault032.Contract.LastReport(&_Yvault032.CallOpts)
}

// LockedProfit is a free data retrieval call binding the contract method 0x44b81396.
//
// Solidity: function lockedProfit() view returns(uint256)
func (_Yvault032 *Yvault032Caller) LockedProfit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault032.contract.Call(opts, &out, "lockedProfit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedProfit is a free data retrieval call binding the contract method 0x44b81396.
//
// Solidity: function lockedProfit() view returns(uint256)
func (_Yvault032 *Yvault032Session) LockedProfit() (*big.Int, error) {
	return _Yvault032.Contract.LockedProfit(&_Yvault032.CallOpts)
}

// LockedProfit is a free data retrieval call binding the contract method 0x44b81396.
//
// Solidity: function lockedProfit() view returns(uint256)
func (_Yvault032 *Yvault032CallerSession) LockedProfit() (*big.Int, error) {
	return _Yvault032.Contract.LockedProfit(&_Yvault032.CallOpts)
}

// LockedProfitDegration is a free data retrieval call binding the contract method 0x2140254d.
//
// Solidity: function lockedProfitDegration() view returns(uint256)
func (_Yvault032 *Yvault032Caller) LockedProfitDegration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault032.contract.Call(opts, &out, "lockedProfitDegration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedProfitDegration is a free data retrieval call binding the contract method 0x2140254d.
//
// Solidity: function lockedProfitDegration() view returns(uint256)
func (_Yvault032 *Yvault032Session) LockedProfitDegration() (*big.Int, error) {
	return _Yvault032.Contract.LockedProfitDegration(&_Yvault032.CallOpts)
}

// LockedProfitDegration is a free data retrieval call binding the contract method 0x2140254d.
//
// Solidity: function lockedProfitDegration() view returns(uint256)
func (_Yvault032 *Yvault032CallerSession) LockedProfitDegration() (*big.Int, error) {
	return _Yvault032.Contract.LockedProfitDegration(&_Yvault032.CallOpts)
}

// Management is a free data retrieval call binding the contract method 0x88a8d602.
//
// Solidity: function management() view returns(address)
func (_Yvault032 *Yvault032Caller) Management(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yvault032.contract.Call(opts, &out, "management")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Management is a free data retrieval call binding the contract method 0x88a8d602.
//
// Solidity: function management() view returns(address)
func (_Yvault032 *Yvault032Session) Management() (common.Address, error) {
	return _Yvault032.Contract.Management(&_Yvault032.CallOpts)
}

// Management is a free data retrieval call binding the contract method 0x88a8d602.
//
// Solidity: function management() view returns(address)
func (_Yvault032 *Yvault032CallerSession) Management() (common.Address, error) {
	return _Yvault032.Contract.Management(&_Yvault032.CallOpts)
}

// ManagementFee is a free data retrieval call binding the contract method 0xa6f7f5d6.
//
// Solidity: function managementFee() view returns(uint256)
func (_Yvault032 *Yvault032Caller) ManagementFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault032.contract.Call(opts, &out, "managementFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ManagementFee is a free data retrieval call binding the contract method 0xa6f7f5d6.
//
// Solidity: function managementFee() view returns(uint256)
func (_Yvault032 *Yvault032Session) ManagementFee() (*big.Int, error) {
	return _Yvault032.Contract.ManagementFee(&_Yvault032.CallOpts)
}

// ManagementFee is a free data retrieval call binding the contract method 0xa6f7f5d6.
//
// Solidity: function managementFee() view returns(uint256)
func (_Yvault032 *Yvault032CallerSession) ManagementFee() (*big.Int, error) {
	return _Yvault032.Contract.ManagementFee(&_Yvault032.CallOpts)
}

// MaxAvailableShares is a free data retrieval call binding the contract method 0x75de2902.
//
// Solidity: function maxAvailableShares() view returns(uint256)
func (_Yvault032 *Yvault032Caller) MaxAvailableShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault032.contract.Call(opts, &out, "maxAvailableShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAvailableShares is a free data retrieval call binding the contract method 0x75de2902.
//
// Solidity: function maxAvailableShares() view returns(uint256)
func (_Yvault032 *Yvault032Session) MaxAvailableShares() (*big.Int, error) {
	return _Yvault032.Contract.MaxAvailableShares(&_Yvault032.CallOpts)
}

// MaxAvailableShares is a free data retrieval call binding the contract method 0x75de2902.
//
// Solidity: function maxAvailableShares() view returns(uint256)
func (_Yvault032 *Yvault032CallerSession) MaxAvailableShares() (*big.Int, error) {
	return _Yvault032.Contract.MaxAvailableShares(&_Yvault032.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Yvault032 *Yvault032Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Yvault032.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Yvault032 *Yvault032Session) Name() (string, error) {
	return _Yvault032.Contract.Name(&_Yvault032.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Yvault032 *Yvault032CallerSession) Name() (string, error) {
	return _Yvault032.Contract.Name(&_Yvault032.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_Yvault032 *Yvault032Caller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yvault032.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_Yvault032 *Yvault032Session) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Yvault032.Contract.Nonces(&_Yvault032.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_Yvault032 *Yvault032CallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Yvault032.Contract.Nonces(&_Yvault032.CallOpts, arg0)
}

// PerformanceFee is a free data retrieval call binding the contract method 0x87788782.
//
// Solidity: function performanceFee() view returns(uint256)
func (_Yvault032 *Yvault032Caller) PerformanceFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault032.contract.Call(opts, &out, "performanceFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PerformanceFee is a free data retrieval call binding the contract method 0x87788782.
//
// Solidity: function performanceFee() view returns(uint256)
func (_Yvault032 *Yvault032Session) PerformanceFee() (*big.Int, error) {
	return _Yvault032.Contract.PerformanceFee(&_Yvault032.CallOpts)
}

// PerformanceFee is a free data retrieval call binding the contract method 0x87788782.
//
// Solidity: function performanceFee() view returns(uint256)
func (_Yvault032 *Yvault032CallerSession) PerformanceFee() (*big.Int, error) {
	return _Yvault032.Contract.PerformanceFee(&_Yvault032.CallOpts)
}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_Yvault032 *Yvault032Caller) PricePerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault032.contract.Call(opts, &out, "pricePerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_Yvault032 *Yvault032Session) PricePerShare() (*big.Int, error) {
	return _Yvault032.Contract.PricePerShare(&_Yvault032.CallOpts)
}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_Yvault032 *Yvault032CallerSession) PricePerShare() (*big.Int, error) {
	return _Yvault032.Contract.PricePerShare(&_Yvault032.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x9ec5a894.
//
// Solidity: function rewards() view returns(address)
func (_Yvault032 *Yvault032Caller) Rewards(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yvault032.contract.Call(opts, &out, "rewards")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0x9ec5a894.
//
// Solidity: function rewards() view returns(address)
func (_Yvault032 *Yvault032Session) Rewards() (common.Address, error) {
	return _Yvault032.Contract.Rewards(&_Yvault032.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x9ec5a894.
//
// Solidity: function rewards() view returns(address)
func (_Yvault032 *Yvault032CallerSession) Rewards() (common.Address, error) {
	return _Yvault032.Contract.Rewards(&_Yvault032.CallOpts)
}

// Strategies is a free data retrieval call binding the contract method 0x39ebf823.
//
// Solidity: function strategies(address arg0) view returns(uint256 performanceFee, uint256 activation, uint256 debtRatio, uint256 minDebtPerHarvest, uint256 maxDebtPerHarvest, uint256 lastReport, uint256 totalDebt, uint256 totalGain, uint256 totalLoss)
func (_Yvault032 *Yvault032Caller) Strategies(opts *bind.CallOpts, arg0 common.Address) (struct {
	PerformanceFee    *big.Int
	Activation        *big.Int
	DebtRatio         *big.Int
	MinDebtPerHarvest *big.Int
	MaxDebtPerHarvest *big.Int
	LastReport        *big.Int
	TotalDebt         *big.Int
	TotalGain         *big.Int
	TotalLoss         *big.Int
}, error) {
	var out []interface{}
	err := _Yvault032.contract.Call(opts, &out, "strategies", arg0)

	outstruct := new(struct {
		PerformanceFee    *big.Int
		Activation        *big.Int
		DebtRatio         *big.Int
		MinDebtPerHarvest *big.Int
		MaxDebtPerHarvest *big.Int
		LastReport        *big.Int
		TotalDebt         *big.Int
		TotalGain         *big.Int
		TotalLoss         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PerformanceFee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Activation = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.DebtRatio = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MinDebtPerHarvest = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.MaxDebtPerHarvest = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.LastReport = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.TotalDebt = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.TotalGain = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.TotalLoss = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Strategies is a free data retrieval call binding the contract method 0x39ebf823.
//
// Solidity: function strategies(address arg0) view returns(uint256 performanceFee, uint256 activation, uint256 debtRatio, uint256 minDebtPerHarvest, uint256 maxDebtPerHarvest, uint256 lastReport, uint256 totalDebt, uint256 totalGain, uint256 totalLoss)
func (_Yvault032 *Yvault032Session) Strategies(arg0 common.Address) (struct {
	PerformanceFee    *big.Int
	Activation        *big.Int
	DebtRatio         *big.Int
	MinDebtPerHarvest *big.Int
	MaxDebtPerHarvest *big.Int
	LastReport        *big.Int
	TotalDebt         *big.Int
	TotalGain         *big.Int
	TotalLoss         *big.Int
}, error) {
	return _Yvault032.Contract.Strategies(&_Yvault032.CallOpts, arg0)
}

// Strategies is a free data retrieval call binding the contract method 0x39ebf823.
//
// Solidity: function strategies(address arg0) view returns(uint256 performanceFee, uint256 activation, uint256 debtRatio, uint256 minDebtPerHarvest, uint256 maxDebtPerHarvest, uint256 lastReport, uint256 totalDebt, uint256 totalGain, uint256 totalLoss)
func (_Yvault032 *Yvault032CallerSession) Strategies(arg0 common.Address) (struct {
	PerformanceFee    *big.Int
	Activation        *big.Int
	DebtRatio         *big.Int
	MinDebtPerHarvest *big.Int
	MaxDebtPerHarvest *big.Int
	LastReport        *big.Int
	TotalDebt         *big.Int
	TotalGain         *big.Int
	TotalLoss         *big.Int
}, error) {
	return _Yvault032.Contract.Strategies(&_Yvault032.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Yvault032 *Yvault032Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Yvault032.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Yvault032 *Yvault032Session) Symbol() (string, error) {
	return _Yvault032.Contract.Symbol(&_Yvault032.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Yvault032 *Yvault032CallerSession) Symbol() (string, error) {
	return _Yvault032.Contract.Symbol(&_Yvault032.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Yvault032 *Yvault032Caller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yvault032.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Yvault032 *Yvault032Session) Token() (common.Address, error) {
	return _Yvault032.Contract.Token(&_Yvault032.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Yvault032 *Yvault032CallerSession) Token() (common.Address, error) {
	return _Yvault032.Contract.Token(&_Yvault032.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_Yvault032 *Yvault032Caller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault032.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_Yvault032 *Yvault032Session) TotalAssets() (*big.Int, error) {
	return _Yvault032.Contract.TotalAssets(&_Yvault032.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_Yvault032 *Yvault032CallerSession) TotalAssets() (*big.Int, error) {
	return _Yvault032.Contract.TotalAssets(&_Yvault032.CallOpts)
}

// TotalDebt is a free data retrieval call binding the contract method 0xfc7b9c18.
//
// Solidity: function totalDebt() view returns(uint256)
func (_Yvault032 *Yvault032Caller) TotalDebt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault032.contract.Call(opts, &out, "totalDebt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDebt is a free data retrieval call binding the contract method 0xfc7b9c18.
//
// Solidity: function totalDebt() view returns(uint256)
func (_Yvault032 *Yvault032Session) TotalDebt() (*big.Int, error) {
	return _Yvault032.Contract.TotalDebt(&_Yvault032.CallOpts)
}

// TotalDebt is a free data retrieval call binding the contract method 0xfc7b9c18.
//
// Solidity: function totalDebt() view returns(uint256)
func (_Yvault032 *Yvault032CallerSession) TotalDebt() (*big.Int, error) {
	return _Yvault032.Contract.TotalDebt(&_Yvault032.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Yvault032 *Yvault032Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault032.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Yvault032 *Yvault032Session) TotalSupply() (*big.Int, error) {
	return _Yvault032.Contract.TotalSupply(&_Yvault032.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Yvault032 *Yvault032CallerSession) TotalSupply() (*big.Int, error) {
	return _Yvault032.Contract.TotalSupply(&_Yvault032.CallOpts)
}

// WithdrawalQueue is a free data retrieval call binding the contract method 0xc822adda.
//
// Solidity: function withdrawalQueue(uint256 arg0) view returns(address)
func (_Yvault032 *Yvault032Caller) WithdrawalQueue(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Yvault032.contract.Call(opts, &out, "withdrawalQueue", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WithdrawalQueue is a free data retrieval call binding the contract method 0xc822adda.
//
// Solidity: function withdrawalQueue(uint256 arg0) view returns(address)
func (_Yvault032 *Yvault032Session) WithdrawalQueue(arg0 *big.Int) (common.Address, error) {
	return _Yvault032.Contract.WithdrawalQueue(&_Yvault032.CallOpts, arg0)
}

// WithdrawalQueue is a free data retrieval call binding the contract method 0xc822adda.
//
// Solidity: function withdrawalQueue(uint256 arg0) view returns(address)
func (_Yvault032 *Yvault032CallerSession) WithdrawalQueue(arg0 *big.Int) (common.Address, error) {
	return _Yvault032.Contract.WithdrawalQueue(&_Yvault032.CallOpts, arg0)
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_Yvault032 *Yvault032Transactor) AcceptGovernance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yvault032.contract.Transact(opts, "acceptGovernance")
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_Yvault032 *Yvault032Session) AcceptGovernance() (*types.Transaction, error) {
	return _Yvault032.Contract.AcceptGovernance(&_Yvault032.TransactOpts)
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_Yvault032 *Yvault032TransactorSession) AcceptGovernance() (*types.Transaction, error) {
	return _Yvault032.Contract.AcceptGovernance(&_Yvault032.TransactOpts)
}

// AddStrategy is a paid mutator transaction binding the contract method 0x14b4e26e.
//
// Solidity: function addStrategy(address strategy, uint256 debtRatio, uint256 minDebtPerHarvest, uint256 maxDebtPerHarvest, uint256 performanceFee) returns()
func (_Yvault032 *Yvault032Transactor) AddStrategy(opts *bind.TransactOpts, strategy common.Address, debtRatio *big.Int, minDebtPerHarvest *big.Int, maxDebtPerHarvest *big.Int, performanceFee *big.Int) (*types.Transaction, error) {
	return _Yvault032.contract.Transact(opts, "addStrategy", strategy, debtRatio, minDebtPerHarvest, maxDebtPerHarvest, performanceFee)
}

// AddStrategy is a paid mutator transaction binding the contract method 0x14b4e26e.
//
// Solidity: function addStrategy(address strategy, uint256 debtRatio, uint256 minDebtPerHarvest, uint256 maxDebtPerHarvest, uint256 performanceFee) returns()
func (_Yvault032 *Yvault032Session) AddStrategy(strategy common.Address, debtRatio *big.Int, minDebtPerHarvest *big.Int, maxDebtPerHarvest *big.Int, performanceFee *big.Int) (*types.Transaction, error) {
	return _Yvault032.Contract.AddStrategy(&_Yvault032.TransactOpts, strategy, debtRatio, minDebtPerHarvest, maxDebtPerHarvest, performanceFee)
}

// AddStrategy is a paid mutator transaction binding the contract method 0x14b4e26e.
//
// Solidity: function addStrategy(address strategy, uint256 debtRatio, uint256 minDebtPerHarvest, uint256 maxDebtPerHarvest, uint256 performanceFee) returns()
func (_Yvault032 *Yvault032TransactorSession) AddStrategy(strategy common.Address, debtRatio *big.Int, minDebtPerHarvest *big.Int, maxDebtPerHarvest *big.Int, performanceFee *big.Int) (*types.Transaction, error) {
	return _Yvault032.Contract.AddStrategy(&_Yvault032.TransactOpts, strategy, debtRatio, minDebtPerHarvest, maxDebtPerHarvest, performanceFee)
}

// AddStrategyToQueue is a paid mutator transaction binding the contract method 0xf76e4caa.
//
// Solidity: function addStrategyToQueue(address strategy) returns()
func (_Yvault032 *Yvault032Transactor) AddStrategyToQueue(opts *bind.TransactOpts, strategy common.Address) (*types.Transaction, error) {
	return _Yvault032.contract.Transact(opts, "addStrategyToQueue", strategy)
}

// AddStrategyToQueue is a paid mutator transaction binding the contract method 0xf76e4caa.
//
// Solidity: function addStrategyToQueue(address strategy) returns()
func (_Yvault032 *Yvault032Session) AddStrategyToQueue(strategy common.Address) (*types.Transaction, error) {
	return _Yvault032.Contract.AddStrategyToQueue(&_Yvault032.TransactOpts, strategy)
}

// AddStrategyToQueue is a paid mutator transaction binding the contract method 0xf76e4caa.
//
// Solidity: function addStrategyToQueue(address strategy) returns()
func (_Yvault032 *Yvault032TransactorSession) AddStrategyToQueue(strategy common.Address) (*types.Transaction, error) {
	return _Yvault032.Contract.AddStrategyToQueue(&_Yvault032.TransactOpts, strategy)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Yvault032 *Yvault032Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault032.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Yvault032 *Yvault032Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault032.Contract.Approve(&_Yvault032.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Yvault032 *Yvault032TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault032.Contract.Approve(&_Yvault032.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 amount) returns(bool)
func (_Yvault032 *Yvault032Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault032.contract.Transact(opts, "decreaseAllowance", spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 amount) returns(bool)
func (_Yvault032 *Yvault032Session) DecreaseAllowance(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault032.Contract.DecreaseAllowance(&_Yvault032.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 amount) returns(bool)
func (_Yvault032 *Yvault032TransactorSession) DecreaseAllowance(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault032.Contract.DecreaseAllowance(&_Yvault032.TransactOpts, spender, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(uint256)
func (_Yvault032 *Yvault032Transactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yvault032.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(uint256)
func (_Yvault032 *Yvault032Session) Deposit() (*types.Transaction, error) {
	return _Yvault032.Contract.Deposit(&_Yvault032.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(uint256)
func (_Yvault032 *Yvault032TransactorSession) Deposit() (*types.Transaction, error) {
	return _Yvault032.Contract.Deposit(&_Yvault032.TransactOpts)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns(uint256)
func (_Yvault032 *Yvault032Transactor) Deposit0(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Yvault032.contract.Transact(opts, "deposit0", _amount)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns(uint256)
func (_Yvault032 *Yvault032Session) Deposit0(_amount *big.Int) (*types.Transaction, error) {
	return _Yvault032.Contract.Deposit0(&_Yvault032.TransactOpts, _amount)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns(uint256)
func (_Yvault032 *Yvault032TransactorSession) Deposit0(_amount *big.Int) (*types.Transaction, error) {
	return _Yvault032.Contract.Deposit0(&_Yvault032.TransactOpts, _amount)
}

// Deposit1 is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _amount, address recipient) returns(uint256)
func (_Yvault032 *Yvault032Transactor) Deposit1(opts *bind.TransactOpts, _amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Yvault032.contract.Transact(opts, "deposit1", _amount, recipient)
}

// Deposit1 is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _amount, address recipient) returns(uint256)
func (_Yvault032 *Yvault032Session) Deposit1(_amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Yvault032.Contract.Deposit1(&_Yvault032.TransactOpts, _amount, recipient)
}

// Deposit1 is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _amount, address recipient) returns(uint256)
func (_Yvault032 *Yvault032TransactorSession) Deposit1(_amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Yvault032.Contract.Deposit1(&_Yvault032.TransactOpts, _amount, recipient)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 amount) returns(bool)
func (_Yvault032 *Yvault032Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault032.contract.Transact(opts, "increaseAllowance", spender, amount)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 amount) returns(bool)
func (_Yvault032 *Yvault032Session) IncreaseAllowance(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault032.Contract.IncreaseAllowance(&_Yvault032.TransactOpts, spender, amount)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 amount) returns(bool)
func (_Yvault032 *Yvault032TransactorSession) IncreaseAllowance(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault032.Contract.IncreaseAllowance(&_Yvault032.TransactOpts, spender, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x83b43589.
//
// Solidity: function initialize(address token, address governance, address rewards, string nameOverride, string symbolOverride) returns()
func (_Yvault032 *Yvault032Transactor) Initialize(opts *bind.TransactOpts, token common.Address, governance common.Address, rewards common.Address, nameOverride string, symbolOverride string) (*types.Transaction, error) {
	return _Yvault032.contract.Transact(opts, "initialize", token, governance, rewards, nameOverride, symbolOverride)
}

// Initialize is a paid mutator transaction binding the contract method 0x83b43589.
//
// Solidity: function initialize(address token, address governance, address rewards, string nameOverride, string symbolOverride) returns()
func (_Yvault032 *Yvault032Session) Initialize(token common.Address, governance common.Address, rewards common.Address, nameOverride string, symbolOverride string) (*types.Transaction, error) {
	return _Yvault032.Contract.Initialize(&_Yvault032.TransactOpts, token, governance, rewards, nameOverride, symbolOverride)
}

// Initialize is a paid mutator transaction binding the contract method 0x83b43589.
//
// Solidity: function initialize(address token, address governance, address rewards, string nameOverride, string symbolOverride) returns()
func (_Yvault032 *Yvault032TransactorSession) Initialize(token common.Address, governance common.Address, rewards common.Address, nameOverride string, symbolOverride string) (*types.Transaction, error) {
	return _Yvault032.Contract.Initialize(&_Yvault032.TransactOpts, token, governance, rewards, nameOverride, symbolOverride)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xa5b81fdf.
//
// Solidity: function initialize(address token, address governance, address rewards, string nameOverride, string symbolOverride, address guardian) returns()
func (_Yvault032 *Yvault032Transactor) Initialize0(opts *bind.TransactOpts, token common.Address, governance common.Address, rewards common.Address, nameOverride string, symbolOverride string, guardian common.Address) (*types.Transaction, error) {
	return _Yvault032.contract.Transact(opts, "initialize0", token, governance, rewards, nameOverride, symbolOverride, guardian)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xa5b81fdf.
//
// Solidity: function initialize(address token, address governance, address rewards, string nameOverride, string symbolOverride, address guardian) returns()
func (_Yvault032 *Yvault032Session) Initialize0(token common.Address, governance common.Address, rewards common.Address, nameOverride string, symbolOverride string, guardian common.Address) (*types.Transaction, error) {
	return _Yvault032.Contract.Initialize0(&_Yvault032.TransactOpts, token, governance, rewards, nameOverride, symbolOverride, guardian)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xa5b81fdf.
//
// Solidity: function initialize(address token, address governance, address rewards, string nameOverride, string symbolOverride, address guardian) returns()
func (_Yvault032 *Yvault032TransactorSession) Initialize0(token common.Address, governance common.Address, rewards common.Address, nameOverride string, symbolOverride string, guardian common.Address) (*types.Transaction, error) {
	return _Yvault032.Contract.Initialize0(&_Yvault032.TransactOpts, token, governance, rewards, nameOverride, symbolOverride, guardian)
}

// MigrateStrategy is a paid mutator transaction binding the contract method 0x6cb56d19.
//
// Solidity: function migrateStrategy(address oldVersion, address newVersion) returns()
func (_Yvault032 *Yvault032Transactor) MigrateStrategy(opts *bind.TransactOpts, oldVersion common.Address, newVersion common.Address) (*types.Transaction, error) {
	return _Yvault032.contract.Transact(opts, "migrateStrategy", oldVersion, newVersion)
}

// MigrateStrategy is a paid mutator transaction binding the contract method 0x6cb56d19.
//
// Solidity: function migrateStrategy(address oldVersion, address newVersion) returns()
func (_Yvault032 *Yvault032Session) MigrateStrategy(oldVersion common.Address, newVersion common.Address) (*types.Transaction, error) {
	return _Yvault032.Contract.MigrateStrategy(&_Yvault032.TransactOpts, oldVersion, newVersion)
}

// MigrateStrategy is a paid mutator transaction binding the contract method 0x6cb56d19.
//
// Solidity: function migrateStrategy(address oldVersion, address newVersion) returns()
func (_Yvault032 *Yvault032TransactorSession) MigrateStrategy(oldVersion common.Address, newVersion common.Address) (*types.Transaction, error) {
	return _Yvault032.Contract.MigrateStrategy(&_Yvault032.TransactOpts, oldVersion, newVersion)
}

// Permit is a paid mutator transaction binding the contract method 0x9fd5a6cf.
//
// Solidity: function permit(address owner, address spender, uint256 amount, uint256 expiry, bytes signature) returns(bool)
func (_Yvault032 *Yvault032Transactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, amount *big.Int, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _Yvault032.contract.Transact(opts, "permit", owner, spender, amount, expiry, signature)
}

// Permit is a paid mutator transaction binding the contract method 0x9fd5a6cf.
//
// Solidity: function permit(address owner, address spender, uint256 amount, uint256 expiry, bytes signature) returns(bool)
func (_Yvault032 *Yvault032Session) Permit(owner common.Address, spender common.Address, amount *big.Int, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _Yvault032.Contract.Permit(&_Yvault032.TransactOpts, owner, spender, amount, expiry, signature)
}

// Permit is a paid mutator transaction binding the contract method 0x9fd5a6cf.
//
// Solidity: function permit(address owner, address spender, uint256 amount, uint256 expiry, bytes signature) returns(bool)
func (_Yvault032 *Yvault032TransactorSession) Permit(owner common.Address, spender common.Address, amount *big.Int, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _Yvault032.Contract.Permit(&_Yvault032.TransactOpts, owner, spender, amount, expiry, signature)
}

// RemoveStrategyFromQueue is a paid mutator transaction binding the contract method 0xb22439f5.
//
// Solidity: function removeStrategyFromQueue(address strategy) returns()
func (_Yvault032 *Yvault032Transactor) RemoveStrategyFromQueue(opts *bind.TransactOpts, strategy common.Address) (*types.Transaction, error) {
	return _Yvault032.contract.Transact(opts, "removeStrategyFromQueue", strategy)
}

// RemoveStrategyFromQueue is a paid mutator transaction binding the contract method 0xb22439f5.
//
// Solidity: function removeStrategyFromQueue(address strategy) returns()
func (_Yvault032 *Yvault032Session) RemoveStrategyFromQueue(strategy common.Address) (*types.Transaction, error) {
	return _Yvault032.Contract.RemoveStrategyFromQueue(&_Yvault032.TransactOpts, strategy)
}

// RemoveStrategyFromQueue is a paid mutator transaction binding the contract method 0xb22439f5.
//
// Solidity: function removeStrategyFromQueue(address strategy) returns()
func (_Yvault032 *Yvault032TransactorSession) RemoveStrategyFromQueue(strategy common.Address) (*types.Transaction, error) {
	return _Yvault032.Contract.RemoveStrategyFromQueue(&_Yvault032.TransactOpts, strategy)
}

// Report is a paid mutator transaction binding the contract method 0xa1d9bafc.
//
// Solidity: function report(uint256 gain, uint256 loss, uint256 _debtPayment) returns(uint256)
func (_Yvault032 *Yvault032Transactor) Report(opts *bind.TransactOpts, gain *big.Int, loss *big.Int, _debtPayment *big.Int) (*types.Transaction, error) {
	return _Yvault032.contract.Transact(opts, "report", gain, loss, _debtPayment)
}

// Report is a paid mutator transaction binding the contract method 0xa1d9bafc.
//
// Solidity: function report(uint256 gain, uint256 loss, uint256 _debtPayment) returns(uint256)
func (_Yvault032 *Yvault032Session) Report(gain *big.Int, loss *big.Int, _debtPayment *big.Int) (*types.Transaction, error) {
	return _Yvault032.Contract.Report(&_Yvault032.TransactOpts, gain, loss, _debtPayment)
}

// Report is a paid mutator transaction binding the contract method 0xa1d9bafc.
//
// Solidity: function report(uint256 gain, uint256 loss, uint256 _debtPayment) returns(uint256)
func (_Yvault032 *Yvault032TransactorSession) Report(gain *big.Int, loss *big.Int, _debtPayment *big.Int) (*types.Transaction, error) {
	return _Yvault032.Contract.Report(&_Yvault032.TransactOpts, gain, loss, _debtPayment)
}

// RevokeStrategy is a paid mutator transaction binding the contract method 0xa0e4af9a.
//
// Solidity: function revokeStrategy() returns()
func (_Yvault032 *Yvault032Transactor) RevokeStrategy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yvault032.contract.Transact(opts, "revokeStrategy")
}

// RevokeStrategy is a paid mutator transaction binding the contract method 0xa0e4af9a.
//
// Solidity: function revokeStrategy() returns()
func (_Yvault032 *Yvault032Session) RevokeStrategy() (*types.Transaction, error) {
	return _Yvault032.Contract.RevokeStrategy(&_Yvault032.TransactOpts)
}

// RevokeStrategy is a paid mutator transaction binding the contract method 0xa0e4af9a.
//
// Solidity: function revokeStrategy() returns()
func (_Yvault032 *Yvault032TransactorSession) RevokeStrategy() (*types.Transaction, error) {
	return _Yvault032.Contract.RevokeStrategy(&_Yvault032.TransactOpts)
}

// RevokeStrategy0 is a paid mutator transaction binding the contract method 0xbb994d48.
//
// Solidity: function revokeStrategy(address strategy) returns()
func (_Yvault032 *Yvault032Transactor) RevokeStrategy0(opts *bind.TransactOpts, strategy common.Address) (*types.Transaction, error) {
	return _Yvault032.contract.Transact(opts, "revokeStrategy0", strategy)
}

// RevokeStrategy0 is a paid mutator transaction binding the contract method 0xbb994d48.
//
// Solidity: function revokeStrategy(address strategy) returns()
func (_Yvault032 *Yvault032Session) RevokeStrategy0(strategy common.Address) (*types.Transaction, error) {
	return _Yvault032.Contract.RevokeStrategy0(&_Yvault032.TransactOpts, strategy)
}

// RevokeStrategy0 is a paid mutator transaction binding the contract method 0xbb994d48.
//
// Solidity: function revokeStrategy(address strategy) returns()
func (_Yvault032 *Yvault032TransactorSession) RevokeStrategy0(strategy common.Address) (*types.Transaction, error) {
	return _Yvault032.Contract.RevokeStrategy0(&_Yvault032.TransactOpts, strategy)
}

// SetDepositLimit is a paid mutator transaction binding the contract method 0xbdc8144b.
//
// Solidity: function setDepositLimit(uint256 limit) returns()
func (_Yvault032 *Yvault032Transactor) SetDepositLimit(opts *bind.TransactOpts, limit *big.Int) (*types.Transaction, error) {
	return _Yvault032.contract.Transact(opts, "setDepositLimit", limit)
}

// SetDepositLimit is a paid mutator transaction binding the contract method 0xbdc8144b.
//
// Solidity: function setDepositLimit(uint256 limit) returns()
func (_Yvault032 *Yvault032Session) SetDepositLimit(limit *big.Int) (*types.Transaction, error) {
	return _Yvault032.Contract.SetDepositLimit(&_Yvault032.TransactOpts, limit)
}

// SetDepositLimit is a paid mutator transaction binding the contract method 0xbdc8144b.
//
// Solidity: function setDepositLimit(uint256 limit) returns()
func (_Yvault032 *Yvault032TransactorSession) SetDepositLimit(limit *big.Int) (*types.Transaction, error) {
	return _Yvault032.Contract.SetDepositLimit(&_Yvault032.TransactOpts, limit)
}

// SetEmergencyShutdown is a paid mutator transaction binding the contract method 0x14c64402.
//
// Solidity: function setEmergencyShutdown(bool active) returns()
func (_Yvault032 *Yvault032Transactor) SetEmergencyShutdown(opts *bind.TransactOpts, active bool) (*types.Transaction, error) {
	return _Yvault032.contract.Transact(opts, "setEmergencyShutdown", active)
}

// SetEmergencyShutdown is a paid mutator transaction binding the contract method 0x14c64402.
//
// Solidity: function setEmergencyShutdown(bool active) returns()
func (_Yvault032 *Yvault032Session) SetEmergencyShutdown(active bool) (*types.Transaction, error) {
	return _Yvault032.Contract.SetEmergencyShutdown(&_Yvault032.TransactOpts, active)
}

// SetEmergencyShutdown is a paid mutator transaction binding the contract method 0x14c64402.
//
// Solidity: function setEmergencyShutdown(bool active) returns()
func (_Yvault032 *Yvault032TransactorSession) SetEmergencyShutdown(active bool) (*types.Transaction, error) {
	return _Yvault032.Contract.SetEmergencyShutdown(&_Yvault032.TransactOpts, active)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address governance) returns()
func (_Yvault032 *Yvault032Transactor) SetGovernance(opts *bind.TransactOpts, governance common.Address) (*types.Transaction, error) {
	return _Yvault032.contract.Transact(opts, "setGovernance", governance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address governance) returns()
func (_Yvault032 *Yvault032Session) SetGovernance(governance common.Address) (*types.Transaction, error) {
	return _Yvault032.Contract.SetGovernance(&_Yvault032.TransactOpts, governance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address governance) returns()
func (_Yvault032 *Yvault032TransactorSession) SetGovernance(governance common.Address) (*types.Transaction, error) {
	return _Yvault032.Contract.SetGovernance(&_Yvault032.TransactOpts, governance)
}

// SetGuardian is a paid mutator transaction binding the contract method 0x8a0dac4a.
//
// Solidity: function setGuardian(address guardian) returns()
func (_Yvault032 *Yvault032Transactor) SetGuardian(opts *bind.TransactOpts, guardian common.Address) (*types.Transaction, error) {
	return _Yvault032.contract.Transact(opts, "setGuardian", guardian)
}

// SetGuardian is a paid mutator transaction binding the contract method 0x8a0dac4a.
//
// Solidity: function setGuardian(address guardian) returns()
func (_Yvault032 *Yvault032Session) SetGuardian(guardian common.Address) (*types.Transaction, error) {
	return _Yvault032.Contract.SetGuardian(&_Yvault032.TransactOpts, guardian)
}

// SetGuardian is a paid mutator transaction binding the contract method 0x8a0dac4a.
//
// Solidity: function setGuardian(address guardian) returns()
func (_Yvault032 *Yvault032TransactorSession) SetGuardian(guardian common.Address) (*types.Transaction, error) {
	return _Yvault032.Contract.SetGuardian(&_Yvault032.TransactOpts, guardian)
}

// SetGuestList is a paid mutator transaction binding the contract method 0x0b5b78eb.
//
// Solidity: function setGuestList(address guestList) returns()
func (_Yvault032 *Yvault032Transactor) SetGuestList(opts *bind.TransactOpts, guestList common.Address) (*types.Transaction, error) {
	return _Yvault032.contract.Transact(opts, "setGuestList", guestList)
}

// SetGuestList is a paid mutator transaction binding the contract method 0x0b5b78eb.
//
// Solidity: function setGuestList(address guestList) returns()
func (_Yvault032 *Yvault032Session) SetGuestList(guestList common.Address) (*types.Transaction, error) {
	return _Yvault032.Contract.SetGuestList(&_Yvault032.TransactOpts, guestList)
}

// SetGuestList is a paid mutator transaction binding the contract method 0x0b5b78eb.
//
// Solidity: function setGuestList(address guestList) returns()
func (_Yvault032 *Yvault032TransactorSession) SetGuestList(guestList common.Address) (*types.Transaction, error) {
	return _Yvault032.Contract.SetGuestList(&_Yvault032.TransactOpts, guestList)
}

// SetLockedProfitDegration is a paid mutator transaction binding the contract method 0x8402a84f.
//
// Solidity: function setLockedProfitDegration(uint256 degration) returns()
func (_Yvault032 *Yvault032Transactor) SetLockedProfitDegration(opts *bind.TransactOpts, degration *big.Int) (*types.Transaction, error) {
	return _Yvault032.contract.Transact(opts, "setLockedProfitDegration", degration)
}

// SetLockedProfitDegration is a paid mutator transaction binding the contract method 0x8402a84f.
//
// Solidity: function setLockedProfitDegration(uint256 degration) returns()
func (_Yvault032 *Yvault032Session) SetLockedProfitDegration(degration *big.Int) (*types.Transaction, error) {
	return _Yvault032.Contract.SetLockedProfitDegration(&_Yvault032.TransactOpts, degration)
}

// SetLockedProfitDegration is a paid mutator transaction binding the contract method 0x8402a84f.
//
// Solidity: function setLockedProfitDegration(uint256 degration) returns()
func (_Yvault032 *Yvault032TransactorSession) SetLockedProfitDegration(degration *big.Int) (*types.Transaction, error) {
	return _Yvault032.Contract.SetLockedProfitDegration(&_Yvault032.TransactOpts, degration)
}

// SetManagement is a paid mutator transaction binding the contract method 0xd4a22bde.
//
// Solidity: function setManagement(address management) returns()
func (_Yvault032 *Yvault032Transactor) SetManagement(opts *bind.TransactOpts, management common.Address) (*types.Transaction, error) {
	return _Yvault032.contract.Transact(opts, "setManagement", management)
}

// SetManagement is a paid mutator transaction binding the contract method 0xd4a22bde.
//
// Solidity: function setManagement(address management) returns()
func (_Yvault032 *Yvault032Session) SetManagement(management common.Address) (*types.Transaction, error) {
	return _Yvault032.Contract.SetManagement(&_Yvault032.TransactOpts, management)
}

// SetManagement is a paid mutator transaction binding the contract method 0xd4a22bde.
//
// Solidity: function setManagement(address management) returns()
func (_Yvault032 *Yvault032TransactorSession) SetManagement(management common.Address) (*types.Transaction, error) {
	return _Yvault032.Contract.SetManagement(&_Yvault032.TransactOpts, management)
}

// SetManagementFee is a paid mutator transaction binding the contract method 0xfe56e232.
//
// Solidity: function setManagementFee(uint256 fee) returns()
func (_Yvault032 *Yvault032Transactor) SetManagementFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _Yvault032.contract.Transact(opts, "setManagementFee", fee)
}

// SetManagementFee is a paid mutator transaction binding the contract method 0xfe56e232.
//
// Solidity: function setManagementFee(uint256 fee) returns()
func (_Yvault032 *Yvault032Session) SetManagementFee(fee *big.Int) (*types.Transaction, error) {
	return _Yvault032.Contract.SetManagementFee(&_Yvault032.TransactOpts, fee)
}

// SetManagementFee is a paid mutator transaction binding the contract method 0xfe56e232.
//
// Solidity: function setManagementFee(uint256 fee) returns()
func (_Yvault032 *Yvault032TransactorSession) SetManagementFee(fee *big.Int) (*types.Transaction, error) {
	return _Yvault032.Contract.SetManagementFee(&_Yvault032.TransactOpts, fee)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name) returns()
func (_Yvault032 *Yvault032Transactor) SetName(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _Yvault032.contract.Transact(opts, "setName", name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name) returns()
func (_Yvault032 *Yvault032Session) SetName(name string) (*types.Transaction, error) {
	return _Yvault032.Contract.SetName(&_Yvault032.TransactOpts, name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name) returns()
func (_Yvault032 *Yvault032TransactorSession) SetName(name string) (*types.Transaction, error) {
	return _Yvault032.Contract.SetName(&_Yvault032.TransactOpts, name)
}

// SetPerformanceFee is a paid mutator transaction binding the contract method 0x70897b23.
//
// Solidity: function setPerformanceFee(uint256 fee) returns()
func (_Yvault032 *Yvault032Transactor) SetPerformanceFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _Yvault032.contract.Transact(opts, "setPerformanceFee", fee)
}

// SetPerformanceFee is a paid mutator transaction binding the contract method 0x70897b23.
//
// Solidity: function setPerformanceFee(uint256 fee) returns()
func (_Yvault032 *Yvault032Session) SetPerformanceFee(fee *big.Int) (*types.Transaction, error) {
	return _Yvault032.Contract.SetPerformanceFee(&_Yvault032.TransactOpts, fee)
}

// SetPerformanceFee is a paid mutator transaction binding the contract method 0x70897b23.
//
// Solidity: function setPerformanceFee(uint256 fee) returns()
func (_Yvault032 *Yvault032TransactorSession) SetPerformanceFee(fee *big.Int) (*types.Transaction, error) {
	return _Yvault032.Contract.SetPerformanceFee(&_Yvault032.TransactOpts, fee)
}

// SetRewards is a paid mutator transaction binding the contract method 0xec38a862.
//
// Solidity: function setRewards(address rewards) returns()
func (_Yvault032 *Yvault032Transactor) SetRewards(opts *bind.TransactOpts, rewards common.Address) (*types.Transaction, error) {
	return _Yvault032.contract.Transact(opts, "setRewards", rewards)
}

// SetRewards is a paid mutator transaction binding the contract method 0xec38a862.
//
// Solidity: function setRewards(address rewards) returns()
func (_Yvault032 *Yvault032Session) SetRewards(rewards common.Address) (*types.Transaction, error) {
	return _Yvault032.Contract.SetRewards(&_Yvault032.TransactOpts, rewards)
}

// SetRewards is a paid mutator transaction binding the contract method 0xec38a862.
//
// Solidity: function setRewards(address rewards) returns()
func (_Yvault032 *Yvault032TransactorSession) SetRewards(rewards common.Address) (*types.Transaction, error) {
	return _Yvault032.Contract.SetRewards(&_Yvault032.TransactOpts, rewards)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string symbol) returns()
func (_Yvault032 *Yvault032Transactor) SetSymbol(opts *bind.TransactOpts, symbol string) (*types.Transaction, error) {
	return _Yvault032.contract.Transact(opts, "setSymbol", symbol)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string symbol) returns()
func (_Yvault032 *Yvault032Session) SetSymbol(symbol string) (*types.Transaction, error) {
	return _Yvault032.Contract.SetSymbol(&_Yvault032.TransactOpts, symbol)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string symbol) returns()
func (_Yvault032 *Yvault032TransactorSession) SetSymbol(symbol string) (*types.Transaction, error) {
	return _Yvault032.Contract.SetSymbol(&_Yvault032.TransactOpts, symbol)
}

// SetWithdrawalQueue is a paid mutator transaction binding the contract method 0x94148415.
//
// Solidity: function setWithdrawalQueue(address[20] queue) returns()
func (_Yvault032 *Yvault032Transactor) SetWithdrawalQueue(opts *bind.TransactOpts, queue [20]common.Address) (*types.Transaction, error) {
	return _Yvault032.contract.Transact(opts, "setWithdrawalQueue", queue)
}

// SetWithdrawalQueue is a paid mutator transaction binding the contract method 0x94148415.
//
// Solidity: function setWithdrawalQueue(address[20] queue) returns()
func (_Yvault032 *Yvault032Session) SetWithdrawalQueue(queue [20]common.Address) (*types.Transaction, error) {
	return _Yvault032.Contract.SetWithdrawalQueue(&_Yvault032.TransactOpts, queue)
}

// SetWithdrawalQueue is a paid mutator transaction binding the contract method 0x94148415.
//
// Solidity: function setWithdrawalQueue(address[20] queue) returns()
func (_Yvault032 *Yvault032TransactorSession) SetWithdrawalQueue(queue [20]common.Address) (*types.Transaction, error) {
	return _Yvault032.Contract.SetWithdrawalQueue(&_Yvault032.TransactOpts, queue)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address token) returns()
func (_Yvault032 *Yvault032Transactor) Sweep(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Yvault032.contract.Transact(opts, "sweep", token)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address token) returns()
func (_Yvault032 *Yvault032Session) Sweep(token common.Address) (*types.Transaction, error) {
	return _Yvault032.Contract.Sweep(&_Yvault032.TransactOpts, token)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address token) returns()
func (_Yvault032 *Yvault032TransactorSession) Sweep(token common.Address) (*types.Transaction, error) {
	return _Yvault032.Contract.Sweep(&_Yvault032.TransactOpts, token)
}

// Sweep0 is a paid mutator transaction binding the contract method 0x6ea056a9.
//
// Solidity: function sweep(address token, uint256 amount) returns()
func (_Yvault032 *Yvault032Transactor) Sweep0(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault032.contract.Transact(opts, "sweep0", token, amount)
}

// Sweep0 is a paid mutator transaction binding the contract method 0x6ea056a9.
//
// Solidity: function sweep(address token, uint256 amount) returns()
func (_Yvault032 *Yvault032Session) Sweep0(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault032.Contract.Sweep0(&_Yvault032.TransactOpts, token, amount)
}

// Sweep0 is a paid mutator transaction binding the contract method 0x6ea056a9.
//
// Solidity: function sweep(address token, uint256 amount) returns()
func (_Yvault032 *Yvault032TransactorSession) Sweep0(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault032.Contract.Sweep0(&_Yvault032.TransactOpts, token, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 amount) returns(bool)
func (_Yvault032 *Yvault032Transactor) Transfer(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault032.contract.Transact(opts, "transfer", receiver, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 amount) returns(bool)
func (_Yvault032 *Yvault032Session) Transfer(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault032.Contract.Transfer(&_Yvault032.TransactOpts, receiver, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 amount) returns(bool)
func (_Yvault032 *Yvault032TransactorSession) Transfer(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault032.Contract.Transfer(&_Yvault032.TransactOpts, receiver, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address receiver, uint256 amount) returns(bool)
func (_Yvault032 *Yvault032Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault032.contract.Transact(opts, "transferFrom", sender, receiver, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address receiver, uint256 amount) returns(bool)
func (_Yvault032 *Yvault032Session) TransferFrom(sender common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault032.Contract.TransferFrom(&_Yvault032.TransactOpts, sender, receiver, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address receiver, uint256 amount) returns(bool)
func (_Yvault032 *Yvault032TransactorSession) TransferFrom(sender common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault032.Contract.TransferFrom(&_Yvault032.TransactOpts, sender, receiver, amount)
}

// UpdateStrategyDebtRatio is a paid mutator transaction binding the contract method 0x7c6a4f24.
//
// Solidity: function updateStrategyDebtRatio(address strategy, uint256 debtRatio) returns()
func (_Yvault032 *Yvault032Transactor) UpdateStrategyDebtRatio(opts *bind.TransactOpts, strategy common.Address, debtRatio *big.Int) (*types.Transaction, error) {
	return _Yvault032.contract.Transact(opts, "updateStrategyDebtRatio", strategy, debtRatio)
}

// UpdateStrategyDebtRatio is a paid mutator transaction binding the contract method 0x7c6a4f24.
//
// Solidity: function updateStrategyDebtRatio(address strategy, uint256 debtRatio) returns()
func (_Yvault032 *Yvault032Session) UpdateStrategyDebtRatio(strategy common.Address, debtRatio *big.Int) (*types.Transaction, error) {
	return _Yvault032.Contract.UpdateStrategyDebtRatio(&_Yvault032.TransactOpts, strategy, debtRatio)
}

// UpdateStrategyDebtRatio is a paid mutator transaction binding the contract method 0x7c6a4f24.
//
// Solidity: function updateStrategyDebtRatio(address strategy, uint256 debtRatio) returns()
func (_Yvault032 *Yvault032TransactorSession) UpdateStrategyDebtRatio(strategy common.Address, debtRatio *big.Int) (*types.Transaction, error) {
	return _Yvault032.Contract.UpdateStrategyDebtRatio(&_Yvault032.TransactOpts, strategy, debtRatio)
}

// UpdateStrategyMaxDebtPerHarvest is a paid mutator transaction binding the contract method 0x4757a156.
//
// Solidity: function updateStrategyMaxDebtPerHarvest(address strategy, uint256 maxDebtPerHarvest) returns()
func (_Yvault032 *Yvault032Transactor) UpdateStrategyMaxDebtPerHarvest(opts *bind.TransactOpts, strategy common.Address, maxDebtPerHarvest *big.Int) (*types.Transaction, error) {
	return _Yvault032.contract.Transact(opts, "updateStrategyMaxDebtPerHarvest", strategy, maxDebtPerHarvest)
}

// UpdateStrategyMaxDebtPerHarvest is a paid mutator transaction binding the contract method 0x4757a156.
//
// Solidity: function updateStrategyMaxDebtPerHarvest(address strategy, uint256 maxDebtPerHarvest) returns()
func (_Yvault032 *Yvault032Session) UpdateStrategyMaxDebtPerHarvest(strategy common.Address, maxDebtPerHarvest *big.Int) (*types.Transaction, error) {
	return _Yvault032.Contract.UpdateStrategyMaxDebtPerHarvest(&_Yvault032.TransactOpts, strategy, maxDebtPerHarvest)
}

// UpdateStrategyMaxDebtPerHarvest is a paid mutator transaction binding the contract method 0x4757a156.
//
// Solidity: function updateStrategyMaxDebtPerHarvest(address strategy, uint256 maxDebtPerHarvest) returns()
func (_Yvault032 *Yvault032TransactorSession) UpdateStrategyMaxDebtPerHarvest(strategy common.Address, maxDebtPerHarvest *big.Int) (*types.Transaction, error) {
	return _Yvault032.Contract.UpdateStrategyMaxDebtPerHarvest(&_Yvault032.TransactOpts, strategy, maxDebtPerHarvest)
}

// UpdateStrategyMinDebtPerHarvest is a paid mutator transaction binding the contract method 0xe722befe.
//
// Solidity: function updateStrategyMinDebtPerHarvest(address strategy, uint256 minDebtPerHarvest) returns()
func (_Yvault032 *Yvault032Transactor) UpdateStrategyMinDebtPerHarvest(opts *bind.TransactOpts, strategy common.Address, minDebtPerHarvest *big.Int) (*types.Transaction, error) {
	return _Yvault032.contract.Transact(opts, "updateStrategyMinDebtPerHarvest", strategy, minDebtPerHarvest)
}

// UpdateStrategyMinDebtPerHarvest is a paid mutator transaction binding the contract method 0xe722befe.
//
// Solidity: function updateStrategyMinDebtPerHarvest(address strategy, uint256 minDebtPerHarvest) returns()
func (_Yvault032 *Yvault032Session) UpdateStrategyMinDebtPerHarvest(strategy common.Address, minDebtPerHarvest *big.Int) (*types.Transaction, error) {
	return _Yvault032.Contract.UpdateStrategyMinDebtPerHarvest(&_Yvault032.TransactOpts, strategy, minDebtPerHarvest)
}

// UpdateStrategyMinDebtPerHarvest is a paid mutator transaction binding the contract method 0xe722befe.
//
// Solidity: function updateStrategyMinDebtPerHarvest(address strategy, uint256 minDebtPerHarvest) returns()
func (_Yvault032 *Yvault032TransactorSession) UpdateStrategyMinDebtPerHarvest(strategy common.Address, minDebtPerHarvest *big.Int) (*types.Transaction, error) {
	return _Yvault032.Contract.UpdateStrategyMinDebtPerHarvest(&_Yvault032.TransactOpts, strategy, minDebtPerHarvest)
}

// UpdateStrategyPerformanceFee is a paid mutator transaction binding the contract method 0xd0194ed6.
//
// Solidity: function updateStrategyPerformanceFee(address strategy, uint256 performanceFee) returns()
func (_Yvault032 *Yvault032Transactor) UpdateStrategyPerformanceFee(opts *bind.TransactOpts, strategy common.Address, performanceFee *big.Int) (*types.Transaction, error) {
	return _Yvault032.contract.Transact(opts, "updateStrategyPerformanceFee", strategy, performanceFee)
}

// UpdateStrategyPerformanceFee is a paid mutator transaction binding the contract method 0xd0194ed6.
//
// Solidity: function updateStrategyPerformanceFee(address strategy, uint256 performanceFee) returns()
func (_Yvault032 *Yvault032Session) UpdateStrategyPerformanceFee(strategy common.Address, performanceFee *big.Int) (*types.Transaction, error) {
	return _Yvault032.Contract.UpdateStrategyPerformanceFee(&_Yvault032.TransactOpts, strategy, performanceFee)
}

// UpdateStrategyPerformanceFee is a paid mutator transaction binding the contract method 0xd0194ed6.
//
// Solidity: function updateStrategyPerformanceFee(address strategy, uint256 performanceFee) returns()
func (_Yvault032 *Yvault032TransactorSession) UpdateStrategyPerformanceFee(strategy common.Address, performanceFee *big.Int) (*types.Transaction, error) {
	return _Yvault032.Contract.UpdateStrategyPerformanceFee(&_Yvault032.TransactOpts, strategy, performanceFee)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(uint256)
func (_Yvault032 *Yvault032Transactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yvault032.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(uint256)
func (_Yvault032 *Yvault032Session) Withdraw() (*types.Transaction, error) {
	return _Yvault032.Contract.Withdraw(&_Yvault032.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(uint256)
func (_Yvault032 *Yvault032TransactorSession) Withdraw() (*types.Transaction, error) {
	return _Yvault032.Contract.Withdraw(&_Yvault032.TransactOpts)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 maxShares) returns(uint256)
func (_Yvault032 *Yvault032Transactor) Withdraw0(opts *bind.TransactOpts, maxShares *big.Int) (*types.Transaction, error) {
	return _Yvault032.contract.Transact(opts, "withdraw0", maxShares)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 maxShares) returns(uint256)
func (_Yvault032 *Yvault032Session) Withdraw0(maxShares *big.Int) (*types.Transaction, error) {
	return _Yvault032.Contract.Withdraw0(&_Yvault032.TransactOpts, maxShares)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 maxShares) returns(uint256)
func (_Yvault032 *Yvault032TransactorSession) Withdraw0(maxShares *big.Int) (*types.Transaction, error) {
	return _Yvault032.Contract.Withdraw0(&_Yvault032.TransactOpts, maxShares)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 maxShares, address recipient) returns(uint256)
func (_Yvault032 *Yvault032Transactor) Withdraw1(opts *bind.TransactOpts, maxShares *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Yvault032.contract.Transact(opts, "withdraw1", maxShares, recipient)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 maxShares, address recipient) returns(uint256)
func (_Yvault032 *Yvault032Session) Withdraw1(maxShares *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Yvault032.Contract.Withdraw1(&_Yvault032.TransactOpts, maxShares, recipient)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 maxShares, address recipient) returns(uint256)
func (_Yvault032 *Yvault032TransactorSession) Withdraw1(maxShares *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Yvault032.Contract.Withdraw1(&_Yvault032.TransactOpts, maxShares, recipient)
}

// Withdraw2 is a paid mutator transaction binding the contract method 0xe63697c8.
//
// Solidity: function withdraw(uint256 maxShares, address recipient, uint256 maxLoss) returns(uint256)
func (_Yvault032 *Yvault032Transactor) Withdraw2(opts *bind.TransactOpts, maxShares *big.Int, recipient common.Address, maxLoss *big.Int) (*types.Transaction, error) {
	return _Yvault032.contract.Transact(opts, "withdraw2", maxShares, recipient, maxLoss)
}

// Withdraw2 is a paid mutator transaction binding the contract method 0xe63697c8.
//
// Solidity: function withdraw(uint256 maxShares, address recipient, uint256 maxLoss) returns(uint256)
func (_Yvault032 *Yvault032Session) Withdraw2(maxShares *big.Int, recipient common.Address, maxLoss *big.Int) (*types.Transaction, error) {
	return _Yvault032.Contract.Withdraw2(&_Yvault032.TransactOpts, maxShares, recipient, maxLoss)
}

// Withdraw2 is a paid mutator transaction binding the contract method 0xe63697c8.
//
// Solidity: function withdraw(uint256 maxShares, address recipient, uint256 maxLoss) returns(uint256)
func (_Yvault032 *Yvault032TransactorSession) Withdraw2(maxShares *big.Int, recipient common.Address, maxLoss *big.Int) (*types.Transaction, error) {
	return _Yvault032.Contract.Withdraw2(&_Yvault032.TransactOpts, maxShares, recipient, maxLoss)
}

// Yvault032ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Yvault032 contract.
type Yvault032ApprovalIterator struct {
	Event *Yvault032Approval // Event containing the contract specifics and raw log

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
func (it *Yvault032ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault032Approval)
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
		it.Event = new(Yvault032Approval)
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
func (it *Yvault032ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault032ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault032Approval represents a Approval event raised by the Yvault032 contract.
type Yvault032Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Yvault032 *Yvault032Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Yvault032ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Yvault032.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Yvault032ApprovalIterator{contract: _Yvault032.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Yvault032 *Yvault032Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Yvault032Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Yvault032.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault032Approval)
				if err := _Yvault032.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Yvault032 *Yvault032Filterer) ParseApproval(log types.Log) (*Yvault032Approval, error) {
	event := new(Yvault032Approval)
	if err := _Yvault032.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault032EmergencyShutdownIterator is returned from FilterEmergencyShutdown and is used to iterate over the raw logs and unpacked data for EmergencyShutdown events raised by the Yvault032 contract.
type Yvault032EmergencyShutdownIterator struct {
	Event *Yvault032EmergencyShutdown // Event containing the contract specifics and raw log

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
func (it *Yvault032EmergencyShutdownIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault032EmergencyShutdown)
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
		it.Event = new(Yvault032EmergencyShutdown)
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
func (it *Yvault032EmergencyShutdownIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault032EmergencyShutdownIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault032EmergencyShutdown represents a EmergencyShutdown event raised by the Yvault032 contract.
type Yvault032EmergencyShutdown struct {
	Active bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyShutdown is a free log retrieval operation binding the contract event 0xba40372a3a724dca3c57156128ef1e896724b65b37a17f190b1ad5de68f3a4f3.
//
// Solidity: event EmergencyShutdown(bool active)
func (_Yvault032 *Yvault032Filterer) FilterEmergencyShutdown(opts *bind.FilterOpts) (*Yvault032EmergencyShutdownIterator, error) {

	logs, sub, err := _Yvault032.contract.FilterLogs(opts, "EmergencyShutdown")
	if err != nil {
		return nil, err
	}
	return &Yvault032EmergencyShutdownIterator{contract: _Yvault032.contract, event: "EmergencyShutdown", logs: logs, sub: sub}, nil
}

// WatchEmergencyShutdown is a free log subscription operation binding the contract event 0xba40372a3a724dca3c57156128ef1e896724b65b37a17f190b1ad5de68f3a4f3.
//
// Solidity: event EmergencyShutdown(bool active)
func (_Yvault032 *Yvault032Filterer) WatchEmergencyShutdown(opts *bind.WatchOpts, sink chan<- *Yvault032EmergencyShutdown) (event.Subscription, error) {

	logs, sub, err := _Yvault032.contract.WatchLogs(opts, "EmergencyShutdown")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault032EmergencyShutdown)
				if err := _Yvault032.contract.UnpackLog(event, "EmergencyShutdown", log); err != nil {
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

// ParseEmergencyShutdown is a log parse operation binding the contract event 0xba40372a3a724dca3c57156128ef1e896724b65b37a17f190b1ad5de68f3a4f3.
//
// Solidity: event EmergencyShutdown(bool active)
func (_Yvault032 *Yvault032Filterer) ParseEmergencyShutdown(log types.Log) (*Yvault032EmergencyShutdown, error) {
	event := new(Yvault032EmergencyShutdown)
	if err := _Yvault032.contract.UnpackLog(event, "EmergencyShutdown", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault032StrategyAddedIterator is returned from FilterStrategyAdded and is used to iterate over the raw logs and unpacked data for StrategyAdded events raised by the Yvault032 contract.
type Yvault032StrategyAddedIterator struct {
	Event *Yvault032StrategyAdded // Event containing the contract specifics and raw log

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
func (it *Yvault032StrategyAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault032StrategyAdded)
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
		it.Event = new(Yvault032StrategyAdded)
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
func (it *Yvault032StrategyAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault032StrategyAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault032StrategyAdded represents a StrategyAdded event raised by the Yvault032 contract.
type Yvault032StrategyAdded struct {
	Strategy          common.Address
	DebtRatio         *big.Int
	MinDebtPerHarvest *big.Int
	MaxDebtPerHarvest *big.Int
	PerformanceFee    *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterStrategyAdded is a free log retrieval operation binding the contract event 0x5a6abd2af9fe6c0554fa08649e2d86e4393ff19dc304d072d38d295c9291d4dc.
//
// Solidity: event StrategyAdded(address indexed strategy, uint256 debtRatio, uint256 minDebtPerHarvest, uint256 maxDebtPerHarvest, uint256 performanceFee)
func (_Yvault032 *Yvault032Filterer) FilterStrategyAdded(opts *bind.FilterOpts, strategy []common.Address) (*Yvault032StrategyAddedIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault032.contract.FilterLogs(opts, "StrategyAdded", strategyRule)
	if err != nil {
		return nil, err
	}
	return &Yvault032StrategyAddedIterator{contract: _Yvault032.contract, event: "StrategyAdded", logs: logs, sub: sub}, nil
}

// WatchStrategyAdded is a free log subscription operation binding the contract event 0x5a6abd2af9fe6c0554fa08649e2d86e4393ff19dc304d072d38d295c9291d4dc.
//
// Solidity: event StrategyAdded(address indexed strategy, uint256 debtRatio, uint256 minDebtPerHarvest, uint256 maxDebtPerHarvest, uint256 performanceFee)
func (_Yvault032 *Yvault032Filterer) WatchStrategyAdded(opts *bind.WatchOpts, sink chan<- *Yvault032StrategyAdded, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault032.contract.WatchLogs(opts, "StrategyAdded", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault032StrategyAdded)
				if err := _Yvault032.contract.UnpackLog(event, "StrategyAdded", log); err != nil {
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

// ParseStrategyAdded is a log parse operation binding the contract event 0x5a6abd2af9fe6c0554fa08649e2d86e4393ff19dc304d072d38d295c9291d4dc.
//
// Solidity: event StrategyAdded(address indexed strategy, uint256 debtRatio, uint256 minDebtPerHarvest, uint256 maxDebtPerHarvest, uint256 performanceFee)
func (_Yvault032 *Yvault032Filterer) ParseStrategyAdded(log types.Log) (*Yvault032StrategyAdded, error) {
	event := new(Yvault032StrategyAdded)
	if err := _Yvault032.contract.UnpackLog(event, "StrategyAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault032StrategyAddedToQueueIterator is returned from FilterStrategyAddedToQueue and is used to iterate over the raw logs and unpacked data for StrategyAddedToQueue events raised by the Yvault032 contract.
type Yvault032StrategyAddedToQueueIterator struct {
	Event *Yvault032StrategyAddedToQueue // Event containing the contract specifics and raw log

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
func (it *Yvault032StrategyAddedToQueueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault032StrategyAddedToQueue)
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
		it.Event = new(Yvault032StrategyAddedToQueue)
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
func (it *Yvault032StrategyAddedToQueueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault032StrategyAddedToQueueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault032StrategyAddedToQueue represents a StrategyAddedToQueue event raised by the Yvault032 contract.
type Yvault032StrategyAddedToQueue struct {
	Strategy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyAddedToQueue is a free log retrieval operation binding the contract event 0xa8727d412c6fa1e2497d6d6f275e2d9fe4d9318d5b793632e60ad9d38ee8f1fa.
//
// Solidity: event StrategyAddedToQueue(address indexed strategy)
func (_Yvault032 *Yvault032Filterer) FilterStrategyAddedToQueue(opts *bind.FilterOpts, strategy []common.Address) (*Yvault032StrategyAddedToQueueIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault032.contract.FilterLogs(opts, "StrategyAddedToQueue", strategyRule)
	if err != nil {
		return nil, err
	}
	return &Yvault032StrategyAddedToQueueIterator{contract: _Yvault032.contract, event: "StrategyAddedToQueue", logs: logs, sub: sub}, nil
}

// WatchStrategyAddedToQueue is a free log subscription operation binding the contract event 0xa8727d412c6fa1e2497d6d6f275e2d9fe4d9318d5b793632e60ad9d38ee8f1fa.
//
// Solidity: event StrategyAddedToQueue(address indexed strategy)
func (_Yvault032 *Yvault032Filterer) WatchStrategyAddedToQueue(opts *bind.WatchOpts, sink chan<- *Yvault032StrategyAddedToQueue, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault032.contract.WatchLogs(opts, "StrategyAddedToQueue", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault032StrategyAddedToQueue)
				if err := _Yvault032.contract.UnpackLog(event, "StrategyAddedToQueue", log); err != nil {
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

// ParseStrategyAddedToQueue is a log parse operation binding the contract event 0xa8727d412c6fa1e2497d6d6f275e2d9fe4d9318d5b793632e60ad9d38ee8f1fa.
//
// Solidity: event StrategyAddedToQueue(address indexed strategy)
func (_Yvault032 *Yvault032Filterer) ParseStrategyAddedToQueue(log types.Log) (*Yvault032StrategyAddedToQueue, error) {
	event := new(Yvault032StrategyAddedToQueue)
	if err := _Yvault032.contract.UnpackLog(event, "StrategyAddedToQueue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault032StrategyMigratedIterator is returned from FilterStrategyMigrated and is used to iterate over the raw logs and unpacked data for StrategyMigrated events raised by the Yvault032 contract.
type Yvault032StrategyMigratedIterator struct {
	Event *Yvault032StrategyMigrated // Event containing the contract specifics and raw log

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
func (it *Yvault032StrategyMigratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault032StrategyMigrated)
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
		it.Event = new(Yvault032StrategyMigrated)
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
func (it *Yvault032StrategyMigratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault032StrategyMigratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault032StrategyMigrated represents a StrategyMigrated event raised by the Yvault032 contract.
type Yvault032StrategyMigrated struct {
	OldVersion common.Address
	NewVersion common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterStrategyMigrated is a free log retrieval operation binding the contract event 0x100b69bb6b504e1252e36b375233158edee64d071b399e2f81473a695fd1b021.
//
// Solidity: event StrategyMigrated(address indexed oldVersion, address indexed newVersion)
func (_Yvault032 *Yvault032Filterer) FilterStrategyMigrated(opts *bind.FilterOpts, oldVersion []common.Address, newVersion []common.Address) (*Yvault032StrategyMigratedIterator, error) {

	var oldVersionRule []interface{}
	for _, oldVersionItem := range oldVersion {
		oldVersionRule = append(oldVersionRule, oldVersionItem)
	}
	var newVersionRule []interface{}
	for _, newVersionItem := range newVersion {
		newVersionRule = append(newVersionRule, newVersionItem)
	}

	logs, sub, err := _Yvault032.contract.FilterLogs(opts, "StrategyMigrated", oldVersionRule, newVersionRule)
	if err != nil {
		return nil, err
	}
	return &Yvault032StrategyMigratedIterator{contract: _Yvault032.contract, event: "StrategyMigrated", logs: logs, sub: sub}, nil
}

// WatchStrategyMigrated is a free log subscription operation binding the contract event 0x100b69bb6b504e1252e36b375233158edee64d071b399e2f81473a695fd1b021.
//
// Solidity: event StrategyMigrated(address indexed oldVersion, address indexed newVersion)
func (_Yvault032 *Yvault032Filterer) WatchStrategyMigrated(opts *bind.WatchOpts, sink chan<- *Yvault032StrategyMigrated, oldVersion []common.Address, newVersion []common.Address) (event.Subscription, error) {

	var oldVersionRule []interface{}
	for _, oldVersionItem := range oldVersion {
		oldVersionRule = append(oldVersionRule, oldVersionItem)
	}
	var newVersionRule []interface{}
	for _, newVersionItem := range newVersion {
		newVersionRule = append(newVersionRule, newVersionItem)
	}

	logs, sub, err := _Yvault032.contract.WatchLogs(opts, "StrategyMigrated", oldVersionRule, newVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault032StrategyMigrated)
				if err := _Yvault032.contract.UnpackLog(event, "StrategyMigrated", log); err != nil {
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

// ParseStrategyMigrated is a log parse operation binding the contract event 0x100b69bb6b504e1252e36b375233158edee64d071b399e2f81473a695fd1b021.
//
// Solidity: event StrategyMigrated(address indexed oldVersion, address indexed newVersion)
func (_Yvault032 *Yvault032Filterer) ParseStrategyMigrated(log types.Log) (*Yvault032StrategyMigrated, error) {
	event := new(Yvault032StrategyMigrated)
	if err := _Yvault032.contract.UnpackLog(event, "StrategyMigrated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault032StrategyRemovedFromQueueIterator is returned from FilterStrategyRemovedFromQueue and is used to iterate over the raw logs and unpacked data for StrategyRemovedFromQueue events raised by the Yvault032 contract.
type Yvault032StrategyRemovedFromQueueIterator struct {
	Event *Yvault032StrategyRemovedFromQueue // Event containing the contract specifics and raw log

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
func (it *Yvault032StrategyRemovedFromQueueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault032StrategyRemovedFromQueue)
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
		it.Event = new(Yvault032StrategyRemovedFromQueue)
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
func (it *Yvault032StrategyRemovedFromQueueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault032StrategyRemovedFromQueueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault032StrategyRemovedFromQueue represents a StrategyRemovedFromQueue event raised by the Yvault032 contract.
type Yvault032StrategyRemovedFromQueue struct {
	Strategy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyRemovedFromQueue is a free log retrieval operation binding the contract event 0x8e1ec3c16d6a67ea8effe2ac7adef9c2de0bc0dc47c49cdf18f6a8b0048085be.
//
// Solidity: event StrategyRemovedFromQueue(address indexed strategy)
func (_Yvault032 *Yvault032Filterer) FilterStrategyRemovedFromQueue(opts *bind.FilterOpts, strategy []common.Address) (*Yvault032StrategyRemovedFromQueueIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault032.contract.FilterLogs(opts, "StrategyRemovedFromQueue", strategyRule)
	if err != nil {
		return nil, err
	}
	return &Yvault032StrategyRemovedFromQueueIterator{contract: _Yvault032.contract, event: "StrategyRemovedFromQueue", logs: logs, sub: sub}, nil
}

// WatchStrategyRemovedFromQueue is a free log subscription operation binding the contract event 0x8e1ec3c16d6a67ea8effe2ac7adef9c2de0bc0dc47c49cdf18f6a8b0048085be.
//
// Solidity: event StrategyRemovedFromQueue(address indexed strategy)
func (_Yvault032 *Yvault032Filterer) WatchStrategyRemovedFromQueue(opts *bind.WatchOpts, sink chan<- *Yvault032StrategyRemovedFromQueue, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault032.contract.WatchLogs(opts, "StrategyRemovedFromQueue", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault032StrategyRemovedFromQueue)
				if err := _Yvault032.contract.UnpackLog(event, "StrategyRemovedFromQueue", log); err != nil {
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

// ParseStrategyRemovedFromQueue is a log parse operation binding the contract event 0x8e1ec3c16d6a67ea8effe2ac7adef9c2de0bc0dc47c49cdf18f6a8b0048085be.
//
// Solidity: event StrategyRemovedFromQueue(address indexed strategy)
func (_Yvault032 *Yvault032Filterer) ParseStrategyRemovedFromQueue(log types.Log) (*Yvault032StrategyRemovedFromQueue, error) {
	event := new(Yvault032StrategyRemovedFromQueue)
	if err := _Yvault032.contract.UnpackLog(event, "StrategyRemovedFromQueue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault032StrategyReportedIterator is returned from FilterStrategyReported and is used to iterate over the raw logs and unpacked data for StrategyReported events raised by the Yvault032 contract.
type Yvault032StrategyReportedIterator struct {
	Event *Yvault032StrategyReported // Event containing the contract specifics and raw log

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
func (it *Yvault032StrategyReportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault032StrategyReported)
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
		it.Event = new(Yvault032StrategyReported)
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
func (it *Yvault032StrategyReportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault032StrategyReportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault032StrategyReported represents a StrategyReported event raised by the Yvault032 contract.
type Yvault032StrategyReported struct {
	Strategy  common.Address
	Gain      *big.Int
	Loss      *big.Int
	DebtPaid  *big.Int
	TotalGain *big.Int
	TotalLoss *big.Int
	TotalDebt *big.Int
	DebtAdded *big.Int
	DebtRatio *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStrategyReported is a free log retrieval operation binding the contract event 0x67f96d2854a335a4cadb49f84fd3ca6f990744ddb3feceeb4b349d2d53d32ad3.
//
// Solidity: event StrategyReported(address indexed strategy, uint256 gain, uint256 loss, uint256 debtPaid, uint256 totalGain, uint256 totalLoss, uint256 totalDebt, uint256 debtAdded, uint256 debtRatio)
func (_Yvault032 *Yvault032Filterer) FilterStrategyReported(opts *bind.FilterOpts, strategy []common.Address) (*Yvault032StrategyReportedIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault032.contract.FilterLogs(opts, "StrategyReported", strategyRule)
	if err != nil {
		return nil, err
	}
	return &Yvault032StrategyReportedIterator{contract: _Yvault032.contract, event: "StrategyReported", logs: logs, sub: sub}, nil
}

// WatchStrategyReported is a free log subscription operation binding the contract event 0x67f96d2854a335a4cadb49f84fd3ca6f990744ddb3feceeb4b349d2d53d32ad3.
//
// Solidity: event StrategyReported(address indexed strategy, uint256 gain, uint256 loss, uint256 debtPaid, uint256 totalGain, uint256 totalLoss, uint256 totalDebt, uint256 debtAdded, uint256 debtRatio)
func (_Yvault032 *Yvault032Filterer) WatchStrategyReported(opts *bind.WatchOpts, sink chan<- *Yvault032StrategyReported, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault032.contract.WatchLogs(opts, "StrategyReported", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault032StrategyReported)
				if err := _Yvault032.contract.UnpackLog(event, "StrategyReported", log); err != nil {
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

// ParseStrategyReported is a log parse operation binding the contract event 0x67f96d2854a335a4cadb49f84fd3ca6f990744ddb3feceeb4b349d2d53d32ad3.
//
// Solidity: event StrategyReported(address indexed strategy, uint256 gain, uint256 loss, uint256 debtPaid, uint256 totalGain, uint256 totalLoss, uint256 totalDebt, uint256 debtAdded, uint256 debtRatio)
func (_Yvault032 *Yvault032Filterer) ParseStrategyReported(log types.Log) (*Yvault032StrategyReported, error) {
	event := new(Yvault032StrategyReported)
	if err := _Yvault032.contract.UnpackLog(event, "StrategyReported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault032StrategyRevokedIterator is returned from FilterStrategyRevoked and is used to iterate over the raw logs and unpacked data for StrategyRevoked events raised by the Yvault032 contract.
type Yvault032StrategyRevokedIterator struct {
	Event *Yvault032StrategyRevoked // Event containing the contract specifics and raw log

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
func (it *Yvault032StrategyRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault032StrategyRevoked)
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
		it.Event = new(Yvault032StrategyRevoked)
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
func (it *Yvault032StrategyRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault032StrategyRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault032StrategyRevoked represents a StrategyRevoked event raised by the Yvault032 contract.
type Yvault032StrategyRevoked struct {
	Strategy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyRevoked is a free log retrieval operation binding the contract event 0x4201c688d84c01154d321afa0c72f1bffe9eef53005c9de9d035074e71e9b32a.
//
// Solidity: event StrategyRevoked(address indexed strategy)
func (_Yvault032 *Yvault032Filterer) FilterStrategyRevoked(opts *bind.FilterOpts, strategy []common.Address) (*Yvault032StrategyRevokedIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault032.contract.FilterLogs(opts, "StrategyRevoked", strategyRule)
	if err != nil {
		return nil, err
	}
	return &Yvault032StrategyRevokedIterator{contract: _Yvault032.contract, event: "StrategyRevoked", logs: logs, sub: sub}, nil
}

// WatchStrategyRevoked is a free log subscription operation binding the contract event 0x4201c688d84c01154d321afa0c72f1bffe9eef53005c9de9d035074e71e9b32a.
//
// Solidity: event StrategyRevoked(address indexed strategy)
func (_Yvault032 *Yvault032Filterer) WatchStrategyRevoked(opts *bind.WatchOpts, sink chan<- *Yvault032StrategyRevoked, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault032.contract.WatchLogs(opts, "StrategyRevoked", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault032StrategyRevoked)
				if err := _Yvault032.contract.UnpackLog(event, "StrategyRevoked", log); err != nil {
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

// ParseStrategyRevoked is a log parse operation binding the contract event 0x4201c688d84c01154d321afa0c72f1bffe9eef53005c9de9d035074e71e9b32a.
//
// Solidity: event StrategyRevoked(address indexed strategy)
func (_Yvault032 *Yvault032Filterer) ParseStrategyRevoked(log types.Log) (*Yvault032StrategyRevoked, error) {
	event := new(Yvault032StrategyRevoked)
	if err := _Yvault032.contract.UnpackLog(event, "StrategyRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault032StrategyUpdateDebtRatioIterator is returned from FilterStrategyUpdateDebtRatio and is used to iterate over the raw logs and unpacked data for StrategyUpdateDebtRatio events raised by the Yvault032 contract.
type Yvault032StrategyUpdateDebtRatioIterator struct {
	Event *Yvault032StrategyUpdateDebtRatio // Event containing the contract specifics and raw log

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
func (it *Yvault032StrategyUpdateDebtRatioIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault032StrategyUpdateDebtRatio)
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
		it.Event = new(Yvault032StrategyUpdateDebtRatio)
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
func (it *Yvault032StrategyUpdateDebtRatioIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault032StrategyUpdateDebtRatioIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault032StrategyUpdateDebtRatio represents a StrategyUpdateDebtRatio event raised by the Yvault032 contract.
type Yvault032StrategyUpdateDebtRatio struct {
	Strategy  common.Address
	DebtRatio *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStrategyUpdateDebtRatio is a free log retrieval operation binding the contract event 0xbda9398315c83ccef012bcaa318a2ff7b680f36429d36597bd4bc25ac11ead59.
//
// Solidity: event StrategyUpdateDebtRatio(address indexed strategy, uint256 debtRatio)
func (_Yvault032 *Yvault032Filterer) FilterStrategyUpdateDebtRatio(opts *bind.FilterOpts, strategy []common.Address) (*Yvault032StrategyUpdateDebtRatioIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault032.contract.FilterLogs(opts, "StrategyUpdateDebtRatio", strategyRule)
	if err != nil {
		return nil, err
	}
	return &Yvault032StrategyUpdateDebtRatioIterator{contract: _Yvault032.contract, event: "StrategyUpdateDebtRatio", logs: logs, sub: sub}, nil
}

// WatchStrategyUpdateDebtRatio is a free log subscription operation binding the contract event 0xbda9398315c83ccef012bcaa318a2ff7b680f36429d36597bd4bc25ac11ead59.
//
// Solidity: event StrategyUpdateDebtRatio(address indexed strategy, uint256 debtRatio)
func (_Yvault032 *Yvault032Filterer) WatchStrategyUpdateDebtRatio(opts *bind.WatchOpts, sink chan<- *Yvault032StrategyUpdateDebtRatio, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault032.contract.WatchLogs(opts, "StrategyUpdateDebtRatio", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault032StrategyUpdateDebtRatio)
				if err := _Yvault032.contract.UnpackLog(event, "StrategyUpdateDebtRatio", log); err != nil {
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

// ParseStrategyUpdateDebtRatio is a log parse operation binding the contract event 0xbda9398315c83ccef012bcaa318a2ff7b680f36429d36597bd4bc25ac11ead59.
//
// Solidity: event StrategyUpdateDebtRatio(address indexed strategy, uint256 debtRatio)
func (_Yvault032 *Yvault032Filterer) ParseStrategyUpdateDebtRatio(log types.Log) (*Yvault032StrategyUpdateDebtRatio, error) {
	event := new(Yvault032StrategyUpdateDebtRatio)
	if err := _Yvault032.contract.UnpackLog(event, "StrategyUpdateDebtRatio", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault032StrategyUpdateMaxDebtPerHarvestIterator is returned from FilterStrategyUpdateMaxDebtPerHarvest and is used to iterate over the raw logs and unpacked data for StrategyUpdateMaxDebtPerHarvest events raised by the Yvault032 contract.
type Yvault032StrategyUpdateMaxDebtPerHarvestIterator struct {
	Event *Yvault032StrategyUpdateMaxDebtPerHarvest // Event containing the contract specifics and raw log

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
func (it *Yvault032StrategyUpdateMaxDebtPerHarvestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault032StrategyUpdateMaxDebtPerHarvest)
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
		it.Event = new(Yvault032StrategyUpdateMaxDebtPerHarvest)
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
func (it *Yvault032StrategyUpdateMaxDebtPerHarvestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault032StrategyUpdateMaxDebtPerHarvestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault032StrategyUpdateMaxDebtPerHarvest represents a StrategyUpdateMaxDebtPerHarvest event raised by the Yvault032 contract.
type Yvault032StrategyUpdateMaxDebtPerHarvest struct {
	Strategy          common.Address
	MaxDebtPerHarvest *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterStrategyUpdateMaxDebtPerHarvest is a free log retrieval operation binding the contract event 0x1796a8e0760e2de5b72e7bf64fccb7666c48ceab94cb6cae7cb7eff4b6f641ab.
//
// Solidity: event StrategyUpdateMaxDebtPerHarvest(address indexed strategy, uint256 maxDebtPerHarvest)
func (_Yvault032 *Yvault032Filterer) FilterStrategyUpdateMaxDebtPerHarvest(opts *bind.FilterOpts, strategy []common.Address) (*Yvault032StrategyUpdateMaxDebtPerHarvestIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault032.contract.FilterLogs(opts, "StrategyUpdateMaxDebtPerHarvest", strategyRule)
	if err != nil {
		return nil, err
	}
	return &Yvault032StrategyUpdateMaxDebtPerHarvestIterator{contract: _Yvault032.contract, event: "StrategyUpdateMaxDebtPerHarvest", logs: logs, sub: sub}, nil
}

// WatchStrategyUpdateMaxDebtPerHarvest is a free log subscription operation binding the contract event 0x1796a8e0760e2de5b72e7bf64fccb7666c48ceab94cb6cae7cb7eff4b6f641ab.
//
// Solidity: event StrategyUpdateMaxDebtPerHarvest(address indexed strategy, uint256 maxDebtPerHarvest)
func (_Yvault032 *Yvault032Filterer) WatchStrategyUpdateMaxDebtPerHarvest(opts *bind.WatchOpts, sink chan<- *Yvault032StrategyUpdateMaxDebtPerHarvest, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault032.contract.WatchLogs(opts, "StrategyUpdateMaxDebtPerHarvest", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault032StrategyUpdateMaxDebtPerHarvest)
				if err := _Yvault032.contract.UnpackLog(event, "StrategyUpdateMaxDebtPerHarvest", log); err != nil {
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

// ParseStrategyUpdateMaxDebtPerHarvest is a log parse operation binding the contract event 0x1796a8e0760e2de5b72e7bf64fccb7666c48ceab94cb6cae7cb7eff4b6f641ab.
//
// Solidity: event StrategyUpdateMaxDebtPerHarvest(address indexed strategy, uint256 maxDebtPerHarvest)
func (_Yvault032 *Yvault032Filterer) ParseStrategyUpdateMaxDebtPerHarvest(log types.Log) (*Yvault032StrategyUpdateMaxDebtPerHarvest, error) {
	event := new(Yvault032StrategyUpdateMaxDebtPerHarvest)
	if err := _Yvault032.contract.UnpackLog(event, "StrategyUpdateMaxDebtPerHarvest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault032StrategyUpdateMinDebtPerHarvestIterator is returned from FilterStrategyUpdateMinDebtPerHarvest and is used to iterate over the raw logs and unpacked data for StrategyUpdateMinDebtPerHarvest events raised by the Yvault032 contract.
type Yvault032StrategyUpdateMinDebtPerHarvestIterator struct {
	Event *Yvault032StrategyUpdateMinDebtPerHarvest // Event containing the contract specifics and raw log

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
func (it *Yvault032StrategyUpdateMinDebtPerHarvestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault032StrategyUpdateMinDebtPerHarvest)
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
		it.Event = new(Yvault032StrategyUpdateMinDebtPerHarvest)
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
func (it *Yvault032StrategyUpdateMinDebtPerHarvestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault032StrategyUpdateMinDebtPerHarvestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault032StrategyUpdateMinDebtPerHarvest represents a StrategyUpdateMinDebtPerHarvest event raised by the Yvault032 contract.
type Yvault032StrategyUpdateMinDebtPerHarvest struct {
	Strategy          common.Address
	MinDebtPerHarvest *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterStrategyUpdateMinDebtPerHarvest is a free log retrieval operation binding the contract event 0x0b728ad785976532c4aaadde09b1cba5f262a7090e83c62d2377bc405678b29c.
//
// Solidity: event StrategyUpdateMinDebtPerHarvest(address indexed strategy, uint256 minDebtPerHarvest)
func (_Yvault032 *Yvault032Filterer) FilterStrategyUpdateMinDebtPerHarvest(opts *bind.FilterOpts, strategy []common.Address) (*Yvault032StrategyUpdateMinDebtPerHarvestIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault032.contract.FilterLogs(opts, "StrategyUpdateMinDebtPerHarvest", strategyRule)
	if err != nil {
		return nil, err
	}
	return &Yvault032StrategyUpdateMinDebtPerHarvestIterator{contract: _Yvault032.contract, event: "StrategyUpdateMinDebtPerHarvest", logs: logs, sub: sub}, nil
}

// WatchStrategyUpdateMinDebtPerHarvest is a free log subscription operation binding the contract event 0x0b728ad785976532c4aaadde09b1cba5f262a7090e83c62d2377bc405678b29c.
//
// Solidity: event StrategyUpdateMinDebtPerHarvest(address indexed strategy, uint256 minDebtPerHarvest)
func (_Yvault032 *Yvault032Filterer) WatchStrategyUpdateMinDebtPerHarvest(opts *bind.WatchOpts, sink chan<- *Yvault032StrategyUpdateMinDebtPerHarvest, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault032.contract.WatchLogs(opts, "StrategyUpdateMinDebtPerHarvest", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault032StrategyUpdateMinDebtPerHarvest)
				if err := _Yvault032.contract.UnpackLog(event, "StrategyUpdateMinDebtPerHarvest", log); err != nil {
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

// ParseStrategyUpdateMinDebtPerHarvest is a log parse operation binding the contract event 0x0b728ad785976532c4aaadde09b1cba5f262a7090e83c62d2377bc405678b29c.
//
// Solidity: event StrategyUpdateMinDebtPerHarvest(address indexed strategy, uint256 minDebtPerHarvest)
func (_Yvault032 *Yvault032Filterer) ParseStrategyUpdateMinDebtPerHarvest(log types.Log) (*Yvault032StrategyUpdateMinDebtPerHarvest, error) {
	event := new(Yvault032StrategyUpdateMinDebtPerHarvest)
	if err := _Yvault032.contract.UnpackLog(event, "StrategyUpdateMinDebtPerHarvest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault032StrategyUpdatePerformanceFeeIterator is returned from FilterStrategyUpdatePerformanceFee and is used to iterate over the raw logs and unpacked data for StrategyUpdatePerformanceFee events raised by the Yvault032 contract.
type Yvault032StrategyUpdatePerformanceFeeIterator struct {
	Event *Yvault032StrategyUpdatePerformanceFee // Event containing the contract specifics and raw log

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
func (it *Yvault032StrategyUpdatePerformanceFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault032StrategyUpdatePerformanceFee)
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
		it.Event = new(Yvault032StrategyUpdatePerformanceFee)
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
func (it *Yvault032StrategyUpdatePerformanceFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault032StrategyUpdatePerformanceFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault032StrategyUpdatePerformanceFee represents a StrategyUpdatePerformanceFee event raised by the Yvault032 contract.
type Yvault032StrategyUpdatePerformanceFee struct {
	Strategy       common.Address
	PerformanceFee *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterStrategyUpdatePerformanceFee is a free log retrieval operation binding the contract event 0xe57488a65fa53066d4c25bac90db47dda4e5de3025ac12bf76ff07211cf7f39e.
//
// Solidity: event StrategyUpdatePerformanceFee(address indexed strategy, uint256 performanceFee)
func (_Yvault032 *Yvault032Filterer) FilterStrategyUpdatePerformanceFee(opts *bind.FilterOpts, strategy []common.Address) (*Yvault032StrategyUpdatePerformanceFeeIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault032.contract.FilterLogs(opts, "StrategyUpdatePerformanceFee", strategyRule)
	if err != nil {
		return nil, err
	}
	return &Yvault032StrategyUpdatePerformanceFeeIterator{contract: _Yvault032.contract, event: "StrategyUpdatePerformanceFee", logs: logs, sub: sub}, nil
}

// WatchStrategyUpdatePerformanceFee is a free log subscription operation binding the contract event 0xe57488a65fa53066d4c25bac90db47dda4e5de3025ac12bf76ff07211cf7f39e.
//
// Solidity: event StrategyUpdatePerformanceFee(address indexed strategy, uint256 performanceFee)
func (_Yvault032 *Yvault032Filterer) WatchStrategyUpdatePerformanceFee(opts *bind.WatchOpts, sink chan<- *Yvault032StrategyUpdatePerformanceFee, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault032.contract.WatchLogs(opts, "StrategyUpdatePerformanceFee", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault032StrategyUpdatePerformanceFee)
				if err := _Yvault032.contract.UnpackLog(event, "StrategyUpdatePerformanceFee", log); err != nil {
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

// ParseStrategyUpdatePerformanceFee is a log parse operation binding the contract event 0xe57488a65fa53066d4c25bac90db47dda4e5de3025ac12bf76ff07211cf7f39e.
//
// Solidity: event StrategyUpdatePerformanceFee(address indexed strategy, uint256 performanceFee)
func (_Yvault032 *Yvault032Filterer) ParseStrategyUpdatePerformanceFee(log types.Log) (*Yvault032StrategyUpdatePerformanceFee, error) {
	event := new(Yvault032StrategyUpdatePerformanceFee)
	if err := _Yvault032.contract.UnpackLog(event, "StrategyUpdatePerformanceFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault032TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Yvault032 contract.
type Yvault032TransferIterator struct {
	Event *Yvault032Transfer // Event containing the contract specifics and raw log

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
func (it *Yvault032TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault032Transfer)
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
		it.Event = new(Yvault032Transfer)
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
func (it *Yvault032TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault032TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault032Transfer represents a Transfer event raised by the Yvault032 contract.
type Yvault032Transfer struct {
	Sender   common.Address
	Receiver common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_Yvault032 *Yvault032Filterer) FilterTransfer(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*Yvault032TransferIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Yvault032.contract.FilterLogs(opts, "Transfer", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &Yvault032TransferIterator{contract: _Yvault032.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_Yvault032 *Yvault032Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Yvault032Transfer, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Yvault032.contract.WatchLogs(opts, "Transfer", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault032Transfer)
				if err := _Yvault032.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_Yvault032 *Yvault032Filterer) ParseTransfer(log types.Log) (*Yvault032Transfer, error) {
	event := new(Yvault032Transfer)
	if err := _Yvault032.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault032UpdateDepositLimitIterator is returned from FilterUpdateDepositLimit and is used to iterate over the raw logs and unpacked data for UpdateDepositLimit events raised by the Yvault032 contract.
type Yvault032UpdateDepositLimitIterator struct {
	Event *Yvault032UpdateDepositLimit // Event containing the contract specifics and raw log

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
func (it *Yvault032UpdateDepositLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault032UpdateDepositLimit)
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
		it.Event = new(Yvault032UpdateDepositLimit)
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
func (it *Yvault032UpdateDepositLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault032UpdateDepositLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault032UpdateDepositLimit represents a UpdateDepositLimit event raised by the Yvault032 contract.
type Yvault032UpdateDepositLimit struct {
	DepositLimit *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateDepositLimit is a free log retrieval operation binding the contract event 0xae565aab888bca5e19e25a13db7b0c9144305bf55cb0f3f4d724f730e5acdd62.
//
// Solidity: event UpdateDepositLimit(uint256 depositLimit)
func (_Yvault032 *Yvault032Filterer) FilterUpdateDepositLimit(opts *bind.FilterOpts) (*Yvault032UpdateDepositLimitIterator, error) {

	logs, sub, err := _Yvault032.contract.FilterLogs(opts, "UpdateDepositLimit")
	if err != nil {
		return nil, err
	}
	return &Yvault032UpdateDepositLimitIterator{contract: _Yvault032.contract, event: "UpdateDepositLimit", logs: logs, sub: sub}, nil
}

// WatchUpdateDepositLimit is a free log subscription operation binding the contract event 0xae565aab888bca5e19e25a13db7b0c9144305bf55cb0f3f4d724f730e5acdd62.
//
// Solidity: event UpdateDepositLimit(uint256 depositLimit)
func (_Yvault032 *Yvault032Filterer) WatchUpdateDepositLimit(opts *bind.WatchOpts, sink chan<- *Yvault032UpdateDepositLimit) (event.Subscription, error) {

	logs, sub, err := _Yvault032.contract.WatchLogs(opts, "UpdateDepositLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault032UpdateDepositLimit)
				if err := _Yvault032.contract.UnpackLog(event, "UpdateDepositLimit", log); err != nil {
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

// ParseUpdateDepositLimit is a log parse operation binding the contract event 0xae565aab888bca5e19e25a13db7b0c9144305bf55cb0f3f4d724f730e5acdd62.
//
// Solidity: event UpdateDepositLimit(uint256 depositLimit)
func (_Yvault032 *Yvault032Filterer) ParseUpdateDepositLimit(log types.Log) (*Yvault032UpdateDepositLimit, error) {
	event := new(Yvault032UpdateDepositLimit)
	if err := _Yvault032.contract.UnpackLog(event, "UpdateDepositLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault032UpdateGovernanceIterator is returned from FilterUpdateGovernance and is used to iterate over the raw logs and unpacked data for UpdateGovernance events raised by the Yvault032 contract.
type Yvault032UpdateGovernanceIterator struct {
	Event *Yvault032UpdateGovernance // Event containing the contract specifics and raw log

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
func (it *Yvault032UpdateGovernanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault032UpdateGovernance)
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
		it.Event = new(Yvault032UpdateGovernance)
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
func (it *Yvault032UpdateGovernanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault032UpdateGovernanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault032UpdateGovernance represents a UpdateGovernance event raised by the Yvault032 contract.
type Yvault032UpdateGovernance struct {
	Governance common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateGovernance is a free log retrieval operation binding the contract event 0x8d55d160c0009eb3d739442df0a3ca089ed64378bfac017e7ddad463f9815b87.
//
// Solidity: event UpdateGovernance(address governance)
func (_Yvault032 *Yvault032Filterer) FilterUpdateGovernance(opts *bind.FilterOpts) (*Yvault032UpdateGovernanceIterator, error) {

	logs, sub, err := _Yvault032.contract.FilterLogs(opts, "UpdateGovernance")
	if err != nil {
		return nil, err
	}
	return &Yvault032UpdateGovernanceIterator{contract: _Yvault032.contract, event: "UpdateGovernance", logs: logs, sub: sub}, nil
}

// WatchUpdateGovernance is a free log subscription operation binding the contract event 0x8d55d160c0009eb3d739442df0a3ca089ed64378bfac017e7ddad463f9815b87.
//
// Solidity: event UpdateGovernance(address governance)
func (_Yvault032 *Yvault032Filterer) WatchUpdateGovernance(opts *bind.WatchOpts, sink chan<- *Yvault032UpdateGovernance) (event.Subscription, error) {

	logs, sub, err := _Yvault032.contract.WatchLogs(opts, "UpdateGovernance")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault032UpdateGovernance)
				if err := _Yvault032.contract.UnpackLog(event, "UpdateGovernance", log); err != nil {
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
// Solidity: event UpdateGovernance(address governance)
func (_Yvault032 *Yvault032Filterer) ParseUpdateGovernance(log types.Log) (*Yvault032UpdateGovernance, error) {
	event := new(Yvault032UpdateGovernance)
	if err := _Yvault032.contract.UnpackLog(event, "UpdateGovernance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault032UpdateGuardianIterator is returned from FilterUpdateGuardian and is used to iterate over the raw logs and unpacked data for UpdateGuardian events raised by the Yvault032 contract.
type Yvault032UpdateGuardianIterator struct {
	Event *Yvault032UpdateGuardian // Event containing the contract specifics and raw log

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
func (it *Yvault032UpdateGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault032UpdateGuardian)
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
		it.Event = new(Yvault032UpdateGuardian)
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
func (it *Yvault032UpdateGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault032UpdateGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault032UpdateGuardian represents a UpdateGuardian event raised by the Yvault032 contract.
type Yvault032UpdateGuardian struct {
	Guardian common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateGuardian is a free log retrieval operation binding the contract event 0x837b9ad138a0a1839a9637afce5306a5c13e23eb63365686843a5319a243609c.
//
// Solidity: event UpdateGuardian(address guardian)
func (_Yvault032 *Yvault032Filterer) FilterUpdateGuardian(opts *bind.FilterOpts) (*Yvault032UpdateGuardianIterator, error) {

	logs, sub, err := _Yvault032.contract.FilterLogs(opts, "UpdateGuardian")
	if err != nil {
		return nil, err
	}
	return &Yvault032UpdateGuardianIterator{contract: _Yvault032.contract, event: "UpdateGuardian", logs: logs, sub: sub}, nil
}

// WatchUpdateGuardian is a free log subscription operation binding the contract event 0x837b9ad138a0a1839a9637afce5306a5c13e23eb63365686843a5319a243609c.
//
// Solidity: event UpdateGuardian(address guardian)
func (_Yvault032 *Yvault032Filterer) WatchUpdateGuardian(opts *bind.WatchOpts, sink chan<- *Yvault032UpdateGuardian) (event.Subscription, error) {

	logs, sub, err := _Yvault032.contract.WatchLogs(opts, "UpdateGuardian")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault032UpdateGuardian)
				if err := _Yvault032.contract.UnpackLog(event, "UpdateGuardian", log); err != nil {
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

// ParseUpdateGuardian is a log parse operation binding the contract event 0x837b9ad138a0a1839a9637afce5306a5c13e23eb63365686843a5319a243609c.
//
// Solidity: event UpdateGuardian(address guardian)
func (_Yvault032 *Yvault032Filterer) ParseUpdateGuardian(log types.Log) (*Yvault032UpdateGuardian, error) {
	event := new(Yvault032UpdateGuardian)
	if err := _Yvault032.contract.UnpackLog(event, "UpdateGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault032UpdateGuestListIterator is returned from FilterUpdateGuestList and is used to iterate over the raw logs and unpacked data for UpdateGuestList events raised by the Yvault032 contract.
type Yvault032UpdateGuestListIterator struct {
	Event *Yvault032UpdateGuestList // Event containing the contract specifics and raw log

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
func (it *Yvault032UpdateGuestListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault032UpdateGuestList)
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
		it.Event = new(Yvault032UpdateGuestList)
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
func (it *Yvault032UpdateGuestListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault032UpdateGuestListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault032UpdateGuestList represents a UpdateGuestList event raised by the Yvault032 contract.
type Yvault032UpdateGuestList struct {
	GuestList common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateGuestList is a free log retrieval operation binding the contract event 0x6d674c311329fb38bbc96dc33d2aad03b9bf9fcfdd8f5e5054fda291a5b3c1f8.
//
// Solidity: event UpdateGuestList(address guestList)
func (_Yvault032 *Yvault032Filterer) FilterUpdateGuestList(opts *bind.FilterOpts) (*Yvault032UpdateGuestListIterator, error) {

	logs, sub, err := _Yvault032.contract.FilterLogs(opts, "UpdateGuestList")
	if err != nil {
		return nil, err
	}
	return &Yvault032UpdateGuestListIterator{contract: _Yvault032.contract, event: "UpdateGuestList", logs: logs, sub: sub}, nil
}

// WatchUpdateGuestList is a free log subscription operation binding the contract event 0x6d674c311329fb38bbc96dc33d2aad03b9bf9fcfdd8f5e5054fda291a5b3c1f8.
//
// Solidity: event UpdateGuestList(address guestList)
func (_Yvault032 *Yvault032Filterer) WatchUpdateGuestList(opts *bind.WatchOpts, sink chan<- *Yvault032UpdateGuestList) (event.Subscription, error) {

	logs, sub, err := _Yvault032.contract.WatchLogs(opts, "UpdateGuestList")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault032UpdateGuestList)
				if err := _Yvault032.contract.UnpackLog(event, "UpdateGuestList", log); err != nil {
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

// ParseUpdateGuestList is a log parse operation binding the contract event 0x6d674c311329fb38bbc96dc33d2aad03b9bf9fcfdd8f5e5054fda291a5b3c1f8.
//
// Solidity: event UpdateGuestList(address guestList)
func (_Yvault032 *Yvault032Filterer) ParseUpdateGuestList(log types.Log) (*Yvault032UpdateGuestList, error) {
	event := new(Yvault032UpdateGuestList)
	if err := _Yvault032.contract.UnpackLog(event, "UpdateGuestList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault032UpdateManagementIterator is returned from FilterUpdateManagement and is used to iterate over the raw logs and unpacked data for UpdateManagement events raised by the Yvault032 contract.
type Yvault032UpdateManagementIterator struct {
	Event *Yvault032UpdateManagement // Event containing the contract specifics and raw log

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
func (it *Yvault032UpdateManagementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault032UpdateManagement)
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
		it.Event = new(Yvault032UpdateManagement)
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
func (it *Yvault032UpdateManagementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault032UpdateManagementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault032UpdateManagement represents a UpdateManagement event raised by the Yvault032 contract.
type Yvault032UpdateManagement struct {
	Management common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateManagement is a free log retrieval operation binding the contract event 0xff54978127edd34aec0f9061fb3b155fbe0ededdfa881ee3e0d541d3a1eef438.
//
// Solidity: event UpdateManagement(address management)
func (_Yvault032 *Yvault032Filterer) FilterUpdateManagement(opts *bind.FilterOpts) (*Yvault032UpdateManagementIterator, error) {

	logs, sub, err := _Yvault032.contract.FilterLogs(opts, "UpdateManagement")
	if err != nil {
		return nil, err
	}
	return &Yvault032UpdateManagementIterator{contract: _Yvault032.contract, event: "UpdateManagement", logs: logs, sub: sub}, nil
}

// WatchUpdateManagement is a free log subscription operation binding the contract event 0xff54978127edd34aec0f9061fb3b155fbe0ededdfa881ee3e0d541d3a1eef438.
//
// Solidity: event UpdateManagement(address management)
func (_Yvault032 *Yvault032Filterer) WatchUpdateManagement(opts *bind.WatchOpts, sink chan<- *Yvault032UpdateManagement) (event.Subscription, error) {

	logs, sub, err := _Yvault032.contract.WatchLogs(opts, "UpdateManagement")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault032UpdateManagement)
				if err := _Yvault032.contract.UnpackLog(event, "UpdateManagement", log); err != nil {
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
// Solidity: event UpdateManagement(address management)
func (_Yvault032 *Yvault032Filterer) ParseUpdateManagement(log types.Log) (*Yvault032UpdateManagement, error) {
	event := new(Yvault032UpdateManagement)
	if err := _Yvault032.contract.UnpackLog(event, "UpdateManagement", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault032UpdateManagementFeeIterator is returned from FilterUpdateManagementFee and is used to iterate over the raw logs and unpacked data for UpdateManagementFee events raised by the Yvault032 contract.
type Yvault032UpdateManagementFeeIterator struct {
	Event *Yvault032UpdateManagementFee // Event containing the contract specifics and raw log

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
func (it *Yvault032UpdateManagementFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault032UpdateManagementFee)
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
		it.Event = new(Yvault032UpdateManagementFee)
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
func (it *Yvault032UpdateManagementFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault032UpdateManagementFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault032UpdateManagementFee represents a UpdateManagementFee event raised by the Yvault032 contract.
type Yvault032UpdateManagementFee struct {
	ManagementFee *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdateManagementFee is a free log retrieval operation binding the contract event 0x7a7883b0074f96e2c7fab65eb25abf624c488761a5db889e3bb84855dcc6daaf.
//
// Solidity: event UpdateManagementFee(uint256 managementFee)
func (_Yvault032 *Yvault032Filterer) FilterUpdateManagementFee(opts *bind.FilterOpts) (*Yvault032UpdateManagementFeeIterator, error) {

	logs, sub, err := _Yvault032.contract.FilterLogs(opts, "UpdateManagementFee")
	if err != nil {
		return nil, err
	}
	return &Yvault032UpdateManagementFeeIterator{contract: _Yvault032.contract, event: "UpdateManagementFee", logs: logs, sub: sub}, nil
}

// WatchUpdateManagementFee is a free log subscription operation binding the contract event 0x7a7883b0074f96e2c7fab65eb25abf624c488761a5db889e3bb84855dcc6daaf.
//
// Solidity: event UpdateManagementFee(uint256 managementFee)
func (_Yvault032 *Yvault032Filterer) WatchUpdateManagementFee(opts *bind.WatchOpts, sink chan<- *Yvault032UpdateManagementFee) (event.Subscription, error) {

	logs, sub, err := _Yvault032.contract.WatchLogs(opts, "UpdateManagementFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault032UpdateManagementFee)
				if err := _Yvault032.contract.UnpackLog(event, "UpdateManagementFee", log); err != nil {
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

// ParseUpdateManagementFee is a log parse operation binding the contract event 0x7a7883b0074f96e2c7fab65eb25abf624c488761a5db889e3bb84855dcc6daaf.
//
// Solidity: event UpdateManagementFee(uint256 managementFee)
func (_Yvault032 *Yvault032Filterer) ParseUpdateManagementFee(log types.Log) (*Yvault032UpdateManagementFee, error) {
	event := new(Yvault032UpdateManagementFee)
	if err := _Yvault032.contract.UnpackLog(event, "UpdateManagementFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault032UpdatePerformanceFeeIterator is returned from FilterUpdatePerformanceFee and is used to iterate over the raw logs and unpacked data for UpdatePerformanceFee events raised by the Yvault032 contract.
type Yvault032UpdatePerformanceFeeIterator struct {
	Event *Yvault032UpdatePerformanceFee // Event containing the contract specifics and raw log

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
func (it *Yvault032UpdatePerformanceFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault032UpdatePerformanceFee)
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
		it.Event = new(Yvault032UpdatePerformanceFee)
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
func (it *Yvault032UpdatePerformanceFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault032UpdatePerformanceFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault032UpdatePerformanceFee represents a UpdatePerformanceFee event raised by the Yvault032 contract.
type Yvault032UpdatePerformanceFee struct {
	PerformanceFee *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpdatePerformanceFee is a free log retrieval operation binding the contract event 0x0810a1c261ca2c0cd86a0152c51c43ba9dc329639d2349f98140891b2ea798eb.
//
// Solidity: event UpdatePerformanceFee(uint256 performanceFee)
func (_Yvault032 *Yvault032Filterer) FilterUpdatePerformanceFee(opts *bind.FilterOpts) (*Yvault032UpdatePerformanceFeeIterator, error) {

	logs, sub, err := _Yvault032.contract.FilterLogs(opts, "UpdatePerformanceFee")
	if err != nil {
		return nil, err
	}
	return &Yvault032UpdatePerformanceFeeIterator{contract: _Yvault032.contract, event: "UpdatePerformanceFee", logs: logs, sub: sub}, nil
}

// WatchUpdatePerformanceFee is a free log subscription operation binding the contract event 0x0810a1c261ca2c0cd86a0152c51c43ba9dc329639d2349f98140891b2ea798eb.
//
// Solidity: event UpdatePerformanceFee(uint256 performanceFee)
func (_Yvault032 *Yvault032Filterer) WatchUpdatePerformanceFee(opts *bind.WatchOpts, sink chan<- *Yvault032UpdatePerformanceFee) (event.Subscription, error) {

	logs, sub, err := _Yvault032.contract.WatchLogs(opts, "UpdatePerformanceFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault032UpdatePerformanceFee)
				if err := _Yvault032.contract.UnpackLog(event, "UpdatePerformanceFee", log); err != nil {
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

// ParseUpdatePerformanceFee is a log parse operation binding the contract event 0x0810a1c261ca2c0cd86a0152c51c43ba9dc329639d2349f98140891b2ea798eb.
//
// Solidity: event UpdatePerformanceFee(uint256 performanceFee)
func (_Yvault032 *Yvault032Filterer) ParseUpdatePerformanceFee(log types.Log) (*Yvault032UpdatePerformanceFee, error) {
	event := new(Yvault032UpdatePerformanceFee)
	if err := _Yvault032.contract.UnpackLog(event, "UpdatePerformanceFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault032UpdateRewardsIterator is returned from FilterUpdateRewards and is used to iterate over the raw logs and unpacked data for UpdateRewards events raised by the Yvault032 contract.
type Yvault032UpdateRewardsIterator struct {
	Event *Yvault032UpdateRewards // Event containing the contract specifics and raw log

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
func (it *Yvault032UpdateRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault032UpdateRewards)
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
		it.Event = new(Yvault032UpdateRewards)
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
func (it *Yvault032UpdateRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault032UpdateRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault032UpdateRewards represents a UpdateRewards event raised by the Yvault032 contract.
type Yvault032UpdateRewards struct {
	Rewards common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateRewards is a free log retrieval operation binding the contract event 0xdf3c41a916aecbf42361a147f8348c242662c3ce20ecef30e826b80642477a3d.
//
// Solidity: event UpdateRewards(address rewards)
func (_Yvault032 *Yvault032Filterer) FilterUpdateRewards(opts *bind.FilterOpts) (*Yvault032UpdateRewardsIterator, error) {

	logs, sub, err := _Yvault032.contract.FilterLogs(opts, "UpdateRewards")
	if err != nil {
		return nil, err
	}
	return &Yvault032UpdateRewardsIterator{contract: _Yvault032.contract, event: "UpdateRewards", logs: logs, sub: sub}, nil
}

// WatchUpdateRewards is a free log subscription operation binding the contract event 0xdf3c41a916aecbf42361a147f8348c242662c3ce20ecef30e826b80642477a3d.
//
// Solidity: event UpdateRewards(address rewards)
func (_Yvault032 *Yvault032Filterer) WatchUpdateRewards(opts *bind.WatchOpts, sink chan<- *Yvault032UpdateRewards) (event.Subscription, error) {

	logs, sub, err := _Yvault032.contract.WatchLogs(opts, "UpdateRewards")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault032UpdateRewards)
				if err := _Yvault032.contract.UnpackLog(event, "UpdateRewards", log); err != nil {
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

// ParseUpdateRewards is a log parse operation binding the contract event 0xdf3c41a916aecbf42361a147f8348c242662c3ce20ecef30e826b80642477a3d.
//
// Solidity: event UpdateRewards(address rewards)
func (_Yvault032 *Yvault032Filterer) ParseUpdateRewards(log types.Log) (*Yvault032UpdateRewards, error) {
	event := new(Yvault032UpdateRewards)
	if err := _Yvault032.contract.UnpackLog(event, "UpdateRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault032UpdateWithdrawalQueueIterator is returned from FilterUpdateWithdrawalQueue and is used to iterate over the raw logs and unpacked data for UpdateWithdrawalQueue events raised by the Yvault032 contract.
type Yvault032UpdateWithdrawalQueueIterator struct {
	Event *Yvault032UpdateWithdrawalQueue // Event containing the contract specifics and raw log

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
func (it *Yvault032UpdateWithdrawalQueueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault032UpdateWithdrawalQueue)
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
		it.Event = new(Yvault032UpdateWithdrawalQueue)
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
func (it *Yvault032UpdateWithdrawalQueueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault032UpdateWithdrawalQueueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault032UpdateWithdrawalQueue represents a UpdateWithdrawalQueue event raised by the Yvault032 contract.
type Yvault032UpdateWithdrawalQueue struct {
	Queue [20]common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdateWithdrawalQueue is a free log retrieval operation binding the contract event 0x695ac3ac73f08f2002284ffe563cefe798ee2878a5e04219522e2e99eb89d168.
//
// Solidity: event UpdateWithdrawalQueue(address[20] queue)
func (_Yvault032 *Yvault032Filterer) FilterUpdateWithdrawalQueue(opts *bind.FilterOpts) (*Yvault032UpdateWithdrawalQueueIterator, error) {

	logs, sub, err := _Yvault032.contract.FilterLogs(opts, "UpdateWithdrawalQueue")
	if err != nil {
		return nil, err
	}
	return &Yvault032UpdateWithdrawalQueueIterator{contract: _Yvault032.contract, event: "UpdateWithdrawalQueue", logs: logs, sub: sub}, nil
}

// WatchUpdateWithdrawalQueue is a free log subscription operation binding the contract event 0x695ac3ac73f08f2002284ffe563cefe798ee2878a5e04219522e2e99eb89d168.
//
// Solidity: event UpdateWithdrawalQueue(address[20] queue)
func (_Yvault032 *Yvault032Filterer) WatchUpdateWithdrawalQueue(opts *bind.WatchOpts, sink chan<- *Yvault032UpdateWithdrawalQueue) (event.Subscription, error) {

	logs, sub, err := _Yvault032.contract.WatchLogs(opts, "UpdateWithdrawalQueue")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault032UpdateWithdrawalQueue)
				if err := _Yvault032.contract.UnpackLog(event, "UpdateWithdrawalQueue", log); err != nil {
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

// ParseUpdateWithdrawalQueue is a log parse operation binding the contract event 0x695ac3ac73f08f2002284ffe563cefe798ee2878a5e04219522e2e99eb89d168.
//
// Solidity: event UpdateWithdrawalQueue(address[20] queue)
func (_Yvault032 *Yvault032Filterer) ParseUpdateWithdrawalQueue(log types.Log) (*Yvault032UpdateWithdrawalQueue, error) {
	event := new(Yvault032UpdateWithdrawalQueue)
	if err := _Yvault032.contract.UnpackLog(event, "UpdateWithdrawalQueue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
