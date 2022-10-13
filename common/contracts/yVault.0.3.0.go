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

// Yvault030MetaData contains all meta data concerning the Yvault030 contract.
var Yvault030MetaData = &bind.MetaData{
	ABI: "[{\"name\":\"Transfer\",\"inputs\":[{\"type\":\"address\",\"name\":\"sender\",\"indexed\":true},{\"type\":\"address\",\"name\":\"receiver\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Approval\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"indexed\":true},{\"type\":\"address\",\"name\":\"spender\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StrategyAdded\",\"inputs\":[{\"type\":\"address\",\"name\":\"strategy\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"debtRatio\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"rateLimit\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"performanceFee\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StrategyReported\",\"inputs\":[{\"type\":\"address\",\"name\":\"strategy\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"gain\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"loss\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"totalGain\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"totalLoss\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"totalDebt\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"debtAdded\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"debtRatio\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateGovernance\",\"inputs\":[{\"type\":\"address\",\"name\":\"governance\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateManagement\",\"inputs\":[{\"type\":\"address\",\"name\":\"management\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateGuestList\",\"inputs\":[{\"type\":\"address\",\"name\":\"guestList\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateRewards\",\"inputs\":[{\"type\":\"address\",\"name\":\"rewards\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateDepositLimit\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"depositLimit\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdatePerformanceFee\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"performanceFee\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateManagementFee\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"managementFee\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateGuardian\",\"inputs\":[{\"type\":\"address\",\"name\":\"guardian\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"EmergencyShutdown\",\"inputs\":[{\"type\":\"bool\",\"name\":\"active\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateWithdrawalQueue\",\"inputs\":[{\"type\":\"address[20]\",\"name\":\"queue\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StrategyUpdateDebtRatio\",\"inputs\":[{\"type\":\"address\",\"name\":\"strategy\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"debtRatio\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StrategyUpdateRateLimit\",\"inputs\":[{\"type\":\"address\",\"name\":\"strategy\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"rateLimit\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StrategyUpdatePerformanceFee\",\"inputs\":[{\"type\":\"address\",\"name\":\"strategy\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"performanceFee\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StrategyMigrated\",\"inputs\":[{\"type\":\"address\",\"name\":\"oldVersion\",\"indexed\":true},{\"type\":\"address\",\"name\":\"newVersion\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StrategyRevoked\",\"inputs\":[{\"type\":\"address\",\"name\":\"strategy\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StrategyRemovedFromQueue\",\"inputs\":[{\"type\":\"address\",\"name\":\"strategy\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StrategyAddedToQueue\",\"inputs\":[{\"type\":\"address\",\"name\":\"strategy\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"initialize\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"token\"},{\"type\":\"address\",\"name\":\"governance\"},{\"type\":\"address\",\"name\":\"rewards\"},{\"type\":\"string\",\"name\":\"nameOverride\"},{\"type\":\"string\",\"name\":\"symbolOverride\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"initialize\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"token\"},{\"type\":\"address\",\"name\":\"governance\"},{\"type\":\"address\",\"name\":\"rewards\"},{\"type\":\"string\",\"name\":\"nameOverride\"},{\"type\":\"string\",\"name\":\"symbolOverride\"},{\"type\":\"address\",\"name\":\"guardian\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"apiVersion\",\"outputs\":[{\"type\":\"string\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\",\"gas\":4519},{\"name\":\"setName\",\"outputs\":[],\"inputs\":[{\"type\":\"string\",\"name\":\"name\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":107017},{\"name\":\"setSymbol\",\"outputs\":[],\"inputs\":[{\"type\":\"string\",\"name\":\"symbol\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":71867},{\"name\":\"setGovernance\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"governance\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":36338},{\"name\":\"acceptGovernance\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37610},{\"name\":\"setManagement\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"management\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37748},{\"name\":\"setGuestList\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"guestList\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37778},{\"name\":\"setRewards\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"rewards\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37808},{\"name\":\"setDepositLimit\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"limit\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37738},{\"name\":\"setPerformanceFee\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"fee\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37872},{\"name\":\"setManagementFee\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"fee\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37902},{\"name\":\"setGuardian\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"guardian\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":39146},{\"name\":\"setEmergencyShutdown\",\"outputs\":[],\"inputs\":[{\"type\":\"bool\",\"name\":\"active\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":39217},{\"name\":\"setWithdrawalQueue\",\"outputs\":[],\"inputs\":[{\"type\":\"address[20]\",\"name\":\"queue\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":763893},{\"name\":\"transfer\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"receiver\"},{\"type\":\"uint256\",\"name\":\"amount\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":76733},{\"name\":\"transferFrom\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"sender\"},{\"type\":\"address\",\"name\":\"receiver\"},{\"type\":\"uint256\",\"name\":\"amount\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":116496},{\"name\":\"approve\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"spender\"},{\"type\":\"uint256\",\"name\":\"amount\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":38244},{\"name\":\"increaseAllowance\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"spender\"},{\"type\":\"uint256\",\"name\":\"amount\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":40285},{\"name\":\"decreaseAllowance\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"spender\"},{\"type\":\"uint256\",\"name\":\"amount\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":40309},{\"name\":\"permit\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"owner\"},{\"type\":\"address\",\"name\":\"spender\"},{\"type\":\"uint256\",\"name\":\"amount\"},{\"type\":\"uint256\",\"name\":\"expiry\"},{\"type\":\"bytes\",\"name\":\"signature\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":81237},{\"name\":\"totalAssets\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":4123},{\"name\":\"deposit\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"deposit\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_amount\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"deposit\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_amount\"},{\"type\":\"address\",\"name\":\"recipient\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"maxAvailableShares\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":364171},{\"name\":\"withdraw\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"withdraw\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"maxShares\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"withdraw\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"maxShares\"},{\"type\":\"address\",\"name\":\"recipient\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"withdraw\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"maxShares\"},{\"type\":\"address\",\"name\":\"recipient\"},{\"type\":\"uint256\",\"name\":\"maxLoss\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"pricePerShare\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":12412},{\"name\":\"addStrategy\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"strategy\"},{\"type\":\"uint256\",\"name\":\"debtRatio\"},{\"type\":\"uint256\",\"name\":\"rateLimit\"},{\"type\":\"uint256\",\"name\":\"performanceFee\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":1450351},{\"name\":\"updateStrategyDebtRatio\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"strategy\"},{\"type\":\"uint256\",\"name\":\"debtRatio\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":115316},{\"name\":\"updateStrategyRateLimit\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"strategy\"},{\"type\":\"uint256\",\"name\":\"rateLimit\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":41467},{\"name\":\"updateStrategyPerformanceFee\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"strategy\"},{\"type\":\"uint256\",\"name\":\"performanceFee\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":41344},{\"name\":\"migrateStrategy\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"oldVersion\"},{\"type\":\"address\",\"name\":\"newVersion\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":1105801},{\"name\":\"revokeStrategy\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"revokeStrategy\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"strategy\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"addStrategyToQueue\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"strategy\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":1196920},{\"name\":\"removeStrategyFromQueue\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"strategy\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":23091666},{\"name\":\"debtOutstanding\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"name\":\"debtOutstanding\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"strategy\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"name\":\"creditAvailable\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"name\":\"creditAvailable\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"strategy\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"name\":\"availableDepositLimit\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":9808},{\"name\":\"expectedReturn\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"name\":\"expectedReturn\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"strategy\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"name\":\"report\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"gain\"},{\"type\":\"uint256\",\"name\":\"loss\"},{\"type\":\"uint256\",\"name\":\"_debtPayment\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":937520},{\"name\":\"sweep\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"token\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"sweep\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"token\"},{\"type\":\"uint256\",\"name\":\"amount\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"name\",\"outputs\":[{\"type\":\"string\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":9053},{\"name\":\"symbol\",\"outputs\":[{\"type\":\"string\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":8106},{\"name\":\"decimals\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2711},{\"name\":\"balanceOf\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2956},{\"name\":\"allowance\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"},{\"type\":\"address\",\"name\":\"arg1\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3201},{\"name\":\"totalSupply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2801},{\"name\":\"token\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2831},{\"name\":\"governance\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2861},{\"name\":\"management\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2891},{\"name\":\"guardian\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2921},{\"name\":\"guestList\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2951},{\"name\":\"strategies\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"performanceFee\"},{\"type\":\"uint256\",\"name\":\"activation\"},{\"type\":\"uint256\",\"name\":\"debtRatio\"},{\"type\":\"uint256\",\"name\":\"rateLimit\"},{\"type\":\"uint256\",\"name\":\"lastReport\"},{\"type\":\"uint256\",\"name\":\"totalDebt\"},{\"type\":\"uint256\",\"name\":\"totalGain\"},{\"type\":\"uint256\",\"name\":\"totalLoss\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":10322},{\"name\":\"withdrawalQueue\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3120},{\"name\":\"emergencyShutdown\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3041},{\"name\":\"depositLimit\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3071},{\"name\":\"debtRatio\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3101},{\"name\":\"totalDebt\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3131},{\"name\":\"lastReport\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3161},{\"name\":\"activation\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3191},{\"name\":\"rewards\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3221},{\"name\":\"managementFee\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3251},{\"name\":\"performanceFee\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3281},{\"name\":\"nonces\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3526},{\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3341}]",
}

// Yvault030ABI is the input ABI used to generate the binding from.
// Deprecated: Use Yvault030MetaData.ABI instead.
var Yvault030ABI = Yvault030MetaData.ABI

// Yvault030 is an auto generated Go binding around an Ethereum contract.
type Yvault030 struct {
	Yvault030Caller     // Read-only binding to the contract
	Yvault030Transactor // Write-only binding to the contract
	Yvault030Filterer   // Log filterer for contract events
}

// Yvault030Caller is an auto generated read-only Go binding around an Ethereum contract.
type Yvault030Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Yvault030Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Yvault030Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Yvault030Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Yvault030Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Yvault030Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Yvault030Session struct {
	Contract     *Yvault030        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Yvault030CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Yvault030CallerSession struct {
	Contract *Yvault030Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// Yvault030TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Yvault030TransactorSession struct {
	Contract     *Yvault030Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// Yvault030Raw is an auto generated low-level Go binding around an Ethereum contract.
type Yvault030Raw struct {
	Contract *Yvault030 // Generic contract binding to access the raw methods on
}

// Yvault030CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Yvault030CallerRaw struct {
	Contract *Yvault030Caller // Generic read-only contract binding to access the raw methods on
}

// Yvault030TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Yvault030TransactorRaw struct {
	Contract *Yvault030Transactor // Generic write-only contract binding to access the raw methods on
}

// NewYvault030 creates a new instance of Yvault030, bound to a specific deployed contract.
func NewYvault030(address common.Address, backend bind.ContractBackend) (*Yvault030, error) {
	contract, err := bindYvault030(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Yvault030{Yvault030Caller: Yvault030Caller{contract: contract}, Yvault030Transactor: Yvault030Transactor{contract: contract}, Yvault030Filterer: Yvault030Filterer{contract: contract}}, nil
}

// NewYvault030Caller creates a new read-only instance of Yvault030, bound to a specific deployed contract.
func NewYvault030Caller(address common.Address, caller bind.ContractCaller) (*Yvault030Caller, error) {
	contract, err := bindYvault030(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Yvault030Caller{contract: contract}, nil
}

// NewYvault030Transactor creates a new write-only instance of Yvault030, bound to a specific deployed contract.
func NewYvault030Transactor(address common.Address, transactor bind.ContractTransactor) (*Yvault030Transactor, error) {
	contract, err := bindYvault030(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Yvault030Transactor{contract: contract}, nil
}

// NewYvault030Filterer creates a new log filterer instance of Yvault030, bound to a specific deployed contract.
func NewYvault030Filterer(address common.Address, filterer bind.ContractFilterer) (*Yvault030Filterer, error) {
	contract, err := bindYvault030(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Yvault030Filterer{contract: contract}, nil
}

// bindYvault030 binds a generic wrapper to an already deployed contract.
func bindYvault030(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Yvault030ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Yvault030 *Yvault030Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Yvault030.Contract.Yvault030Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Yvault030 *Yvault030Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yvault030.Contract.Yvault030Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Yvault030 *Yvault030Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Yvault030.Contract.Yvault030Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Yvault030 *Yvault030CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Yvault030.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Yvault030 *Yvault030TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yvault030.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Yvault030 *Yvault030TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Yvault030.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Yvault030 *Yvault030Caller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Yvault030.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Yvault030 *Yvault030Session) DOMAINSEPARATOR() ([32]byte, error) {
	return _Yvault030.Contract.DOMAINSEPARATOR(&_Yvault030.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Yvault030 *Yvault030CallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Yvault030.Contract.DOMAINSEPARATOR(&_Yvault030.CallOpts)
}

// Activation is a free data retrieval call binding the contract method 0x3629c8de.
//
// Solidity: function activation() view returns(uint256)
func (_Yvault030 *Yvault030Caller) Activation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault030.contract.Call(opts, &out, "activation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Activation is a free data retrieval call binding the contract method 0x3629c8de.
//
// Solidity: function activation() view returns(uint256)
func (_Yvault030 *Yvault030Session) Activation() (*big.Int, error) {
	return _Yvault030.Contract.Activation(&_Yvault030.CallOpts)
}

// Activation is a free data retrieval call binding the contract method 0x3629c8de.
//
// Solidity: function activation() view returns(uint256)
func (_Yvault030 *Yvault030CallerSession) Activation() (*big.Int, error) {
	return _Yvault030.Contract.Activation(&_Yvault030.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_Yvault030 *Yvault030Caller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yvault030.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_Yvault030 *Yvault030Session) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Yvault030.Contract.Allowance(&_Yvault030.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_Yvault030 *Yvault030CallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Yvault030.Contract.Allowance(&_Yvault030.CallOpts, arg0, arg1)
}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() pure returns(string)
func (_Yvault030 *Yvault030Caller) ApiVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Yvault030.contract.Call(opts, &out, "apiVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() pure returns(string)
func (_Yvault030 *Yvault030Session) ApiVersion() (string, error) {
	return _Yvault030.Contract.ApiVersion(&_Yvault030.CallOpts)
}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() pure returns(string)
func (_Yvault030 *Yvault030CallerSession) ApiVersion() (string, error) {
	return _Yvault030.Contract.ApiVersion(&_Yvault030.CallOpts)
}

// AvailableDepositLimit is a free data retrieval call binding the contract method 0x153c27c4.
//
// Solidity: function availableDepositLimit() view returns(uint256)
func (_Yvault030 *Yvault030Caller) AvailableDepositLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault030.contract.Call(opts, &out, "availableDepositLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailableDepositLimit is a free data retrieval call binding the contract method 0x153c27c4.
//
// Solidity: function availableDepositLimit() view returns(uint256)
func (_Yvault030 *Yvault030Session) AvailableDepositLimit() (*big.Int, error) {
	return _Yvault030.Contract.AvailableDepositLimit(&_Yvault030.CallOpts)
}

// AvailableDepositLimit is a free data retrieval call binding the contract method 0x153c27c4.
//
// Solidity: function availableDepositLimit() view returns(uint256)
func (_Yvault030 *Yvault030CallerSession) AvailableDepositLimit() (*big.Int, error) {
	return _Yvault030.Contract.AvailableDepositLimit(&_Yvault030.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_Yvault030 *Yvault030Caller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yvault030.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_Yvault030 *Yvault030Session) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Yvault030.Contract.BalanceOf(&_Yvault030.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_Yvault030 *Yvault030CallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Yvault030.Contract.BalanceOf(&_Yvault030.CallOpts, arg0)
}

// CreditAvailable is a free data retrieval call binding the contract method 0x112c1f9b.
//
// Solidity: function creditAvailable() view returns(uint256)
func (_Yvault030 *Yvault030Caller) CreditAvailable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault030.contract.Call(opts, &out, "creditAvailable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreditAvailable is a free data retrieval call binding the contract method 0x112c1f9b.
//
// Solidity: function creditAvailable() view returns(uint256)
func (_Yvault030 *Yvault030Session) CreditAvailable() (*big.Int, error) {
	return _Yvault030.Contract.CreditAvailable(&_Yvault030.CallOpts)
}

// CreditAvailable is a free data retrieval call binding the contract method 0x112c1f9b.
//
// Solidity: function creditAvailable() view returns(uint256)
func (_Yvault030 *Yvault030CallerSession) CreditAvailable() (*big.Int, error) {
	return _Yvault030.Contract.CreditAvailable(&_Yvault030.CallOpts)
}

// CreditAvailable0 is a free data retrieval call binding the contract method 0xd7648013.
//
// Solidity: function creditAvailable(address strategy) view returns(uint256)
func (_Yvault030 *Yvault030Caller) CreditAvailable0(opts *bind.CallOpts, strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yvault030.contract.Call(opts, &out, "creditAvailable0", strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreditAvailable0 is a free data retrieval call binding the contract method 0xd7648013.
//
// Solidity: function creditAvailable(address strategy) view returns(uint256)
func (_Yvault030 *Yvault030Session) CreditAvailable0(strategy common.Address) (*big.Int, error) {
	return _Yvault030.Contract.CreditAvailable0(&_Yvault030.CallOpts, strategy)
}

// CreditAvailable0 is a free data retrieval call binding the contract method 0xd7648013.
//
// Solidity: function creditAvailable(address strategy) view returns(uint256)
func (_Yvault030 *Yvault030CallerSession) CreditAvailable0(strategy common.Address) (*big.Int, error) {
	return _Yvault030.Contract.CreditAvailable0(&_Yvault030.CallOpts, strategy)
}

// DebtOutstanding is a free data retrieval call binding the contract method 0xbf3759b5.
//
// Solidity: function debtOutstanding() view returns(uint256)
func (_Yvault030 *Yvault030Caller) DebtOutstanding(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault030.contract.Call(opts, &out, "debtOutstanding")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtOutstanding is a free data retrieval call binding the contract method 0xbf3759b5.
//
// Solidity: function debtOutstanding() view returns(uint256)
func (_Yvault030 *Yvault030Session) DebtOutstanding() (*big.Int, error) {
	return _Yvault030.Contract.DebtOutstanding(&_Yvault030.CallOpts)
}

// DebtOutstanding is a free data retrieval call binding the contract method 0xbf3759b5.
//
// Solidity: function debtOutstanding() view returns(uint256)
func (_Yvault030 *Yvault030CallerSession) DebtOutstanding() (*big.Int, error) {
	return _Yvault030.Contract.DebtOutstanding(&_Yvault030.CallOpts)
}

// DebtOutstanding0 is a free data retrieval call binding the contract method 0xbdcf36bb.
//
// Solidity: function debtOutstanding(address strategy) view returns(uint256)
func (_Yvault030 *Yvault030Caller) DebtOutstanding0(opts *bind.CallOpts, strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yvault030.contract.Call(opts, &out, "debtOutstanding0", strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtOutstanding0 is a free data retrieval call binding the contract method 0xbdcf36bb.
//
// Solidity: function debtOutstanding(address strategy) view returns(uint256)
func (_Yvault030 *Yvault030Session) DebtOutstanding0(strategy common.Address) (*big.Int, error) {
	return _Yvault030.Contract.DebtOutstanding0(&_Yvault030.CallOpts, strategy)
}

// DebtOutstanding0 is a free data retrieval call binding the contract method 0xbdcf36bb.
//
// Solidity: function debtOutstanding(address strategy) view returns(uint256)
func (_Yvault030 *Yvault030CallerSession) DebtOutstanding0(strategy common.Address) (*big.Int, error) {
	return _Yvault030.Contract.DebtOutstanding0(&_Yvault030.CallOpts, strategy)
}

// DebtRatio is a free data retrieval call binding the contract method 0xcea55f57.
//
// Solidity: function debtRatio() view returns(uint256)
func (_Yvault030 *Yvault030Caller) DebtRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault030.contract.Call(opts, &out, "debtRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtRatio is a free data retrieval call binding the contract method 0xcea55f57.
//
// Solidity: function debtRatio() view returns(uint256)
func (_Yvault030 *Yvault030Session) DebtRatio() (*big.Int, error) {
	return _Yvault030.Contract.DebtRatio(&_Yvault030.CallOpts)
}

// DebtRatio is a free data retrieval call binding the contract method 0xcea55f57.
//
// Solidity: function debtRatio() view returns(uint256)
func (_Yvault030 *Yvault030CallerSession) DebtRatio() (*big.Int, error) {
	return _Yvault030.Contract.DebtRatio(&_Yvault030.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Yvault030 *Yvault030Caller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault030.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Yvault030 *Yvault030Session) Decimals() (*big.Int, error) {
	return _Yvault030.Contract.Decimals(&_Yvault030.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Yvault030 *Yvault030CallerSession) Decimals() (*big.Int, error) {
	return _Yvault030.Contract.Decimals(&_Yvault030.CallOpts)
}

// DepositLimit is a free data retrieval call binding the contract method 0xecf70858.
//
// Solidity: function depositLimit() view returns(uint256)
func (_Yvault030 *Yvault030Caller) DepositLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault030.contract.Call(opts, &out, "depositLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositLimit is a free data retrieval call binding the contract method 0xecf70858.
//
// Solidity: function depositLimit() view returns(uint256)
func (_Yvault030 *Yvault030Session) DepositLimit() (*big.Int, error) {
	return _Yvault030.Contract.DepositLimit(&_Yvault030.CallOpts)
}

// DepositLimit is a free data retrieval call binding the contract method 0xecf70858.
//
// Solidity: function depositLimit() view returns(uint256)
func (_Yvault030 *Yvault030CallerSession) DepositLimit() (*big.Int, error) {
	return _Yvault030.Contract.DepositLimit(&_Yvault030.CallOpts)
}

// EmergencyShutdown is a free data retrieval call binding the contract method 0x3403c2fc.
//
// Solidity: function emergencyShutdown() view returns(bool)
func (_Yvault030 *Yvault030Caller) EmergencyShutdown(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Yvault030.contract.Call(opts, &out, "emergencyShutdown")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EmergencyShutdown is a free data retrieval call binding the contract method 0x3403c2fc.
//
// Solidity: function emergencyShutdown() view returns(bool)
func (_Yvault030 *Yvault030Session) EmergencyShutdown() (bool, error) {
	return _Yvault030.Contract.EmergencyShutdown(&_Yvault030.CallOpts)
}

// EmergencyShutdown is a free data retrieval call binding the contract method 0x3403c2fc.
//
// Solidity: function emergencyShutdown() view returns(bool)
func (_Yvault030 *Yvault030CallerSession) EmergencyShutdown() (bool, error) {
	return _Yvault030.Contract.EmergencyShutdown(&_Yvault030.CallOpts)
}

// ExpectedReturn is a free data retrieval call binding the contract method 0xd3406abd.
//
// Solidity: function expectedReturn() view returns(uint256)
func (_Yvault030 *Yvault030Caller) ExpectedReturn(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault030.contract.Call(opts, &out, "expectedReturn")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpectedReturn is a free data retrieval call binding the contract method 0xd3406abd.
//
// Solidity: function expectedReturn() view returns(uint256)
func (_Yvault030 *Yvault030Session) ExpectedReturn() (*big.Int, error) {
	return _Yvault030.Contract.ExpectedReturn(&_Yvault030.CallOpts)
}

// ExpectedReturn is a free data retrieval call binding the contract method 0xd3406abd.
//
// Solidity: function expectedReturn() view returns(uint256)
func (_Yvault030 *Yvault030CallerSession) ExpectedReturn() (*big.Int, error) {
	return _Yvault030.Contract.ExpectedReturn(&_Yvault030.CallOpts)
}

// ExpectedReturn0 is a free data retrieval call binding the contract method 0x33586b67.
//
// Solidity: function expectedReturn(address strategy) view returns(uint256)
func (_Yvault030 *Yvault030Caller) ExpectedReturn0(opts *bind.CallOpts, strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yvault030.contract.Call(opts, &out, "expectedReturn0", strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpectedReturn0 is a free data retrieval call binding the contract method 0x33586b67.
//
// Solidity: function expectedReturn(address strategy) view returns(uint256)
func (_Yvault030 *Yvault030Session) ExpectedReturn0(strategy common.Address) (*big.Int, error) {
	return _Yvault030.Contract.ExpectedReturn0(&_Yvault030.CallOpts, strategy)
}

// ExpectedReturn0 is a free data retrieval call binding the contract method 0x33586b67.
//
// Solidity: function expectedReturn(address strategy) view returns(uint256)
func (_Yvault030 *Yvault030CallerSession) ExpectedReturn0(strategy common.Address) (*big.Int, error) {
	return _Yvault030.Contract.ExpectedReturn0(&_Yvault030.CallOpts, strategy)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Yvault030 *Yvault030Caller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yvault030.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Yvault030 *Yvault030Session) Governance() (common.Address, error) {
	return _Yvault030.Contract.Governance(&_Yvault030.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Yvault030 *Yvault030CallerSession) Governance() (common.Address, error) {
	return _Yvault030.Contract.Governance(&_Yvault030.CallOpts)
}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_Yvault030 *Yvault030Caller) Guardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yvault030.contract.Call(opts, &out, "guardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_Yvault030 *Yvault030Session) Guardian() (common.Address, error) {
	return _Yvault030.Contract.Guardian(&_Yvault030.CallOpts)
}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_Yvault030 *Yvault030CallerSession) Guardian() (common.Address, error) {
	return _Yvault030.Contract.Guardian(&_Yvault030.CallOpts)
}

// GuestList is a free data retrieval call binding the contract method 0x46d55875.
//
// Solidity: function guestList() view returns(address)
func (_Yvault030 *Yvault030Caller) GuestList(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yvault030.contract.Call(opts, &out, "guestList")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GuestList is a free data retrieval call binding the contract method 0x46d55875.
//
// Solidity: function guestList() view returns(address)
func (_Yvault030 *Yvault030Session) GuestList() (common.Address, error) {
	return _Yvault030.Contract.GuestList(&_Yvault030.CallOpts)
}

// GuestList is a free data retrieval call binding the contract method 0x46d55875.
//
// Solidity: function guestList() view returns(address)
func (_Yvault030 *Yvault030CallerSession) GuestList() (common.Address, error) {
	return _Yvault030.Contract.GuestList(&_Yvault030.CallOpts)
}

// LastReport is a free data retrieval call binding the contract method 0xc3535b52.
//
// Solidity: function lastReport() view returns(uint256)
func (_Yvault030 *Yvault030Caller) LastReport(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault030.contract.Call(opts, &out, "lastReport")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastReport is a free data retrieval call binding the contract method 0xc3535b52.
//
// Solidity: function lastReport() view returns(uint256)
func (_Yvault030 *Yvault030Session) LastReport() (*big.Int, error) {
	return _Yvault030.Contract.LastReport(&_Yvault030.CallOpts)
}

// LastReport is a free data retrieval call binding the contract method 0xc3535b52.
//
// Solidity: function lastReport() view returns(uint256)
func (_Yvault030 *Yvault030CallerSession) LastReport() (*big.Int, error) {
	return _Yvault030.Contract.LastReport(&_Yvault030.CallOpts)
}

// Management is a free data retrieval call binding the contract method 0x88a8d602.
//
// Solidity: function management() view returns(address)
func (_Yvault030 *Yvault030Caller) Management(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yvault030.contract.Call(opts, &out, "management")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Management is a free data retrieval call binding the contract method 0x88a8d602.
//
// Solidity: function management() view returns(address)
func (_Yvault030 *Yvault030Session) Management() (common.Address, error) {
	return _Yvault030.Contract.Management(&_Yvault030.CallOpts)
}

// Management is a free data retrieval call binding the contract method 0x88a8d602.
//
// Solidity: function management() view returns(address)
func (_Yvault030 *Yvault030CallerSession) Management() (common.Address, error) {
	return _Yvault030.Contract.Management(&_Yvault030.CallOpts)
}

// ManagementFee is a free data retrieval call binding the contract method 0xa6f7f5d6.
//
// Solidity: function managementFee() view returns(uint256)
func (_Yvault030 *Yvault030Caller) ManagementFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault030.contract.Call(opts, &out, "managementFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ManagementFee is a free data retrieval call binding the contract method 0xa6f7f5d6.
//
// Solidity: function managementFee() view returns(uint256)
func (_Yvault030 *Yvault030Session) ManagementFee() (*big.Int, error) {
	return _Yvault030.Contract.ManagementFee(&_Yvault030.CallOpts)
}

// ManagementFee is a free data retrieval call binding the contract method 0xa6f7f5d6.
//
// Solidity: function managementFee() view returns(uint256)
func (_Yvault030 *Yvault030CallerSession) ManagementFee() (*big.Int, error) {
	return _Yvault030.Contract.ManagementFee(&_Yvault030.CallOpts)
}

// MaxAvailableShares is a free data retrieval call binding the contract method 0x75de2902.
//
// Solidity: function maxAvailableShares() view returns(uint256)
func (_Yvault030 *Yvault030Caller) MaxAvailableShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault030.contract.Call(opts, &out, "maxAvailableShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAvailableShares is a free data retrieval call binding the contract method 0x75de2902.
//
// Solidity: function maxAvailableShares() view returns(uint256)
func (_Yvault030 *Yvault030Session) MaxAvailableShares() (*big.Int, error) {
	return _Yvault030.Contract.MaxAvailableShares(&_Yvault030.CallOpts)
}

// MaxAvailableShares is a free data retrieval call binding the contract method 0x75de2902.
//
// Solidity: function maxAvailableShares() view returns(uint256)
func (_Yvault030 *Yvault030CallerSession) MaxAvailableShares() (*big.Int, error) {
	return _Yvault030.Contract.MaxAvailableShares(&_Yvault030.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Yvault030 *Yvault030Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Yvault030.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Yvault030 *Yvault030Session) Name() (string, error) {
	return _Yvault030.Contract.Name(&_Yvault030.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Yvault030 *Yvault030CallerSession) Name() (string, error) {
	return _Yvault030.Contract.Name(&_Yvault030.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_Yvault030 *Yvault030Caller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yvault030.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_Yvault030 *Yvault030Session) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Yvault030.Contract.Nonces(&_Yvault030.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_Yvault030 *Yvault030CallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Yvault030.Contract.Nonces(&_Yvault030.CallOpts, arg0)
}

// PerformanceFee is a free data retrieval call binding the contract method 0x87788782.
//
// Solidity: function performanceFee() view returns(uint256)
func (_Yvault030 *Yvault030Caller) PerformanceFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault030.contract.Call(opts, &out, "performanceFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PerformanceFee is a free data retrieval call binding the contract method 0x87788782.
//
// Solidity: function performanceFee() view returns(uint256)
func (_Yvault030 *Yvault030Session) PerformanceFee() (*big.Int, error) {
	return _Yvault030.Contract.PerformanceFee(&_Yvault030.CallOpts)
}

// PerformanceFee is a free data retrieval call binding the contract method 0x87788782.
//
// Solidity: function performanceFee() view returns(uint256)
func (_Yvault030 *Yvault030CallerSession) PerformanceFee() (*big.Int, error) {
	return _Yvault030.Contract.PerformanceFee(&_Yvault030.CallOpts)
}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_Yvault030 *Yvault030Caller) PricePerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault030.contract.Call(opts, &out, "pricePerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_Yvault030 *Yvault030Session) PricePerShare() (*big.Int, error) {
	return _Yvault030.Contract.PricePerShare(&_Yvault030.CallOpts)
}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_Yvault030 *Yvault030CallerSession) PricePerShare() (*big.Int, error) {
	return _Yvault030.Contract.PricePerShare(&_Yvault030.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x9ec5a894.
//
// Solidity: function rewards() view returns(address)
func (_Yvault030 *Yvault030Caller) Rewards(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yvault030.contract.Call(opts, &out, "rewards")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0x9ec5a894.
//
// Solidity: function rewards() view returns(address)
func (_Yvault030 *Yvault030Session) Rewards() (common.Address, error) {
	return _Yvault030.Contract.Rewards(&_Yvault030.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x9ec5a894.
//
// Solidity: function rewards() view returns(address)
func (_Yvault030 *Yvault030CallerSession) Rewards() (common.Address, error) {
	return _Yvault030.Contract.Rewards(&_Yvault030.CallOpts)
}

// Strategies is a free data retrieval call binding the contract method 0x39ebf823.
//
// Solidity: function strategies(address arg0) view returns(uint256 performanceFee, uint256 activation, uint256 debtRatio, uint256 rateLimit, uint256 lastReport, uint256 totalDebt, uint256 totalGain, uint256 totalLoss)
func (_Yvault030 *Yvault030Caller) Strategies(opts *bind.CallOpts, arg0 common.Address) (struct {
	PerformanceFee *big.Int
	Activation     *big.Int
	DebtRatio      *big.Int
	RateLimit      *big.Int
	LastReport     *big.Int
	TotalDebt      *big.Int
	TotalGain      *big.Int
	TotalLoss      *big.Int
}, error) {
	var out []interface{}
	err := _Yvault030.contract.Call(opts, &out, "strategies", arg0)

	outstruct := new(struct {
		PerformanceFee *big.Int
		Activation     *big.Int
		DebtRatio      *big.Int
		RateLimit      *big.Int
		LastReport     *big.Int
		TotalDebt      *big.Int
		TotalGain      *big.Int
		TotalLoss      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PerformanceFee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Activation = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.DebtRatio = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.RateLimit = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LastReport = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.TotalDebt = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.TotalGain = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.TotalLoss = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Strategies is a free data retrieval call binding the contract method 0x39ebf823.
//
// Solidity: function strategies(address arg0) view returns(uint256 performanceFee, uint256 activation, uint256 debtRatio, uint256 rateLimit, uint256 lastReport, uint256 totalDebt, uint256 totalGain, uint256 totalLoss)
func (_Yvault030 *Yvault030Session) Strategies(arg0 common.Address) (struct {
	PerformanceFee *big.Int
	Activation     *big.Int
	DebtRatio      *big.Int
	RateLimit      *big.Int
	LastReport     *big.Int
	TotalDebt      *big.Int
	TotalGain      *big.Int
	TotalLoss      *big.Int
}, error) {
	return _Yvault030.Contract.Strategies(&_Yvault030.CallOpts, arg0)
}

// Strategies is a free data retrieval call binding the contract method 0x39ebf823.
//
// Solidity: function strategies(address arg0) view returns(uint256 performanceFee, uint256 activation, uint256 debtRatio, uint256 rateLimit, uint256 lastReport, uint256 totalDebt, uint256 totalGain, uint256 totalLoss)
func (_Yvault030 *Yvault030CallerSession) Strategies(arg0 common.Address) (struct {
	PerformanceFee *big.Int
	Activation     *big.Int
	DebtRatio      *big.Int
	RateLimit      *big.Int
	LastReport     *big.Int
	TotalDebt      *big.Int
	TotalGain      *big.Int
	TotalLoss      *big.Int
}, error) {
	return _Yvault030.Contract.Strategies(&_Yvault030.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Yvault030 *Yvault030Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Yvault030.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Yvault030 *Yvault030Session) Symbol() (string, error) {
	return _Yvault030.Contract.Symbol(&_Yvault030.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Yvault030 *Yvault030CallerSession) Symbol() (string, error) {
	return _Yvault030.Contract.Symbol(&_Yvault030.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Yvault030 *Yvault030Caller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yvault030.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Yvault030 *Yvault030Session) Token() (common.Address, error) {
	return _Yvault030.Contract.Token(&_Yvault030.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Yvault030 *Yvault030CallerSession) Token() (common.Address, error) {
	return _Yvault030.Contract.Token(&_Yvault030.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_Yvault030 *Yvault030Caller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault030.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_Yvault030 *Yvault030Session) TotalAssets() (*big.Int, error) {
	return _Yvault030.Contract.TotalAssets(&_Yvault030.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_Yvault030 *Yvault030CallerSession) TotalAssets() (*big.Int, error) {
	return _Yvault030.Contract.TotalAssets(&_Yvault030.CallOpts)
}

// TotalDebt is a free data retrieval call binding the contract method 0xfc7b9c18.
//
// Solidity: function totalDebt() view returns(uint256)
func (_Yvault030 *Yvault030Caller) TotalDebt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault030.contract.Call(opts, &out, "totalDebt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDebt is a free data retrieval call binding the contract method 0xfc7b9c18.
//
// Solidity: function totalDebt() view returns(uint256)
func (_Yvault030 *Yvault030Session) TotalDebt() (*big.Int, error) {
	return _Yvault030.Contract.TotalDebt(&_Yvault030.CallOpts)
}

// TotalDebt is a free data retrieval call binding the contract method 0xfc7b9c18.
//
// Solidity: function totalDebt() view returns(uint256)
func (_Yvault030 *Yvault030CallerSession) TotalDebt() (*big.Int, error) {
	return _Yvault030.Contract.TotalDebt(&_Yvault030.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Yvault030 *Yvault030Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault030.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Yvault030 *Yvault030Session) TotalSupply() (*big.Int, error) {
	return _Yvault030.Contract.TotalSupply(&_Yvault030.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Yvault030 *Yvault030CallerSession) TotalSupply() (*big.Int, error) {
	return _Yvault030.Contract.TotalSupply(&_Yvault030.CallOpts)
}

// WithdrawalQueue is a free data retrieval call binding the contract method 0xc822adda.
//
// Solidity: function withdrawalQueue(uint256 arg0) view returns(address)
func (_Yvault030 *Yvault030Caller) WithdrawalQueue(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Yvault030.contract.Call(opts, &out, "withdrawalQueue", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WithdrawalQueue is a free data retrieval call binding the contract method 0xc822adda.
//
// Solidity: function withdrawalQueue(uint256 arg0) view returns(address)
func (_Yvault030 *Yvault030Session) WithdrawalQueue(arg0 *big.Int) (common.Address, error) {
	return _Yvault030.Contract.WithdrawalQueue(&_Yvault030.CallOpts, arg0)
}

// WithdrawalQueue is a free data retrieval call binding the contract method 0xc822adda.
//
// Solidity: function withdrawalQueue(uint256 arg0) view returns(address)
func (_Yvault030 *Yvault030CallerSession) WithdrawalQueue(arg0 *big.Int) (common.Address, error) {
	return _Yvault030.Contract.WithdrawalQueue(&_Yvault030.CallOpts, arg0)
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_Yvault030 *Yvault030Transactor) AcceptGovernance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yvault030.contract.Transact(opts, "acceptGovernance")
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_Yvault030 *Yvault030Session) AcceptGovernance() (*types.Transaction, error) {
	return _Yvault030.Contract.AcceptGovernance(&_Yvault030.TransactOpts)
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_Yvault030 *Yvault030TransactorSession) AcceptGovernance() (*types.Transaction, error) {
	return _Yvault030.Contract.AcceptGovernance(&_Yvault030.TransactOpts)
}

// AddStrategy is a paid mutator transaction binding the contract method 0x0dd21b6c.
//
// Solidity: function addStrategy(address strategy, uint256 debtRatio, uint256 rateLimit, uint256 performanceFee) returns()
func (_Yvault030 *Yvault030Transactor) AddStrategy(opts *bind.TransactOpts, strategy common.Address, debtRatio *big.Int, rateLimit *big.Int, performanceFee *big.Int) (*types.Transaction, error) {
	return _Yvault030.contract.Transact(opts, "addStrategy", strategy, debtRatio, rateLimit, performanceFee)
}

// AddStrategy is a paid mutator transaction binding the contract method 0x0dd21b6c.
//
// Solidity: function addStrategy(address strategy, uint256 debtRatio, uint256 rateLimit, uint256 performanceFee) returns()
func (_Yvault030 *Yvault030Session) AddStrategy(strategy common.Address, debtRatio *big.Int, rateLimit *big.Int, performanceFee *big.Int) (*types.Transaction, error) {
	return _Yvault030.Contract.AddStrategy(&_Yvault030.TransactOpts, strategy, debtRatio, rateLimit, performanceFee)
}

// AddStrategy is a paid mutator transaction binding the contract method 0x0dd21b6c.
//
// Solidity: function addStrategy(address strategy, uint256 debtRatio, uint256 rateLimit, uint256 performanceFee) returns()
func (_Yvault030 *Yvault030TransactorSession) AddStrategy(strategy common.Address, debtRatio *big.Int, rateLimit *big.Int, performanceFee *big.Int) (*types.Transaction, error) {
	return _Yvault030.Contract.AddStrategy(&_Yvault030.TransactOpts, strategy, debtRatio, rateLimit, performanceFee)
}

// AddStrategyToQueue is a paid mutator transaction binding the contract method 0xf76e4caa.
//
// Solidity: function addStrategyToQueue(address strategy) returns()
func (_Yvault030 *Yvault030Transactor) AddStrategyToQueue(opts *bind.TransactOpts, strategy common.Address) (*types.Transaction, error) {
	return _Yvault030.contract.Transact(opts, "addStrategyToQueue", strategy)
}

// AddStrategyToQueue is a paid mutator transaction binding the contract method 0xf76e4caa.
//
// Solidity: function addStrategyToQueue(address strategy) returns()
func (_Yvault030 *Yvault030Session) AddStrategyToQueue(strategy common.Address) (*types.Transaction, error) {
	return _Yvault030.Contract.AddStrategyToQueue(&_Yvault030.TransactOpts, strategy)
}

// AddStrategyToQueue is a paid mutator transaction binding the contract method 0xf76e4caa.
//
// Solidity: function addStrategyToQueue(address strategy) returns()
func (_Yvault030 *Yvault030TransactorSession) AddStrategyToQueue(strategy common.Address) (*types.Transaction, error) {
	return _Yvault030.Contract.AddStrategyToQueue(&_Yvault030.TransactOpts, strategy)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Yvault030 *Yvault030Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault030.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Yvault030 *Yvault030Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault030.Contract.Approve(&_Yvault030.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Yvault030 *Yvault030TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault030.Contract.Approve(&_Yvault030.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 amount) returns(bool)
func (_Yvault030 *Yvault030Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault030.contract.Transact(opts, "decreaseAllowance", spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 amount) returns(bool)
func (_Yvault030 *Yvault030Session) DecreaseAllowance(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault030.Contract.DecreaseAllowance(&_Yvault030.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 amount) returns(bool)
func (_Yvault030 *Yvault030TransactorSession) DecreaseAllowance(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault030.Contract.DecreaseAllowance(&_Yvault030.TransactOpts, spender, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(uint256)
func (_Yvault030 *Yvault030Transactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yvault030.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(uint256)
func (_Yvault030 *Yvault030Session) Deposit() (*types.Transaction, error) {
	return _Yvault030.Contract.Deposit(&_Yvault030.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(uint256)
func (_Yvault030 *Yvault030TransactorSession) Deposit() (*types.Transaction, error) {
	return _Yvault030.Contract.Deposit(&_Yvault030.TransactOpts)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns(uint256)
func (_Yvault030 *Yvault030Transactor) Deposit0(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Yvault030.contract.Transact(opts, "deposit0", _amount)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns(uint256)
func (_Yvault030 *Yvault030Session) Deposit0(_amount *big.Int) (*types.Transaction, error) {
	return _Yvault030.Contract.Deposit0(&_Yvault030.TransactOpts, _amount)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns(uint256)
func (_Yvault030 *Yvault030TransactorSession) Deposit0(_amount *big.Int) (*types.Transaction, error) {
	return _Yvault030.Contract.Deposit0(&_Yvault030.TransactOpts, _amount)
}

// Deposit1 is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _amount, address recipient) returns(uint256)
func (_Yvault030 *Yvault030Transactor) Deposit1(opts *bind.TransactOpts, _amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Yvault030.contract.Transact(opts, "deposit1", _amount, recipient)
}

// Deposit1 is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _amount, address recipient) returns(uint256)
func (_Yvault030 *Yvault030Session) Deposit1(_amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Yvault030.Contract.Deposit1(&_Yvault030.TransactOpts, _amount, recipient)
}

// Deposit1 is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _amount, address recipient) returns(uint256)
func (_Yvault030 *Yvault030TransactorSession) Deposit1(_amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Yvault030.Contract.Deposit1(&_Yvault030.TransactOpts, _amount, recipient)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 amount) returns(bool)
func (_Yvault030 *Yvault030Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault030.contract.Transact(opts, "increaseAllowance", spender, amount)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 amount) returns(bool)
func (_Yvault030 *Yvault030Session) IncreaseAllowance(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault030.Contract.IncreaseAllowance(&_Yvault030.TransactOpts, spender, amount)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 amount) returns(bool)
func (_Yvault030 *Yvault030TransactorSession) IncreaseAllowance(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault030.Contract.IncreaseAllowance(&_Yvault030.TransactOpts, spender, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x83b43589.
//
// Solidity: function initialize(address token, address governance, address rewards, string nameOverride, string symbolOverride) returns()
func (_Yvault030 *Yvault030Transactor) Initialize(opts *bind.TransactOpts, token common.Address, governance common.Address, rewards common.Address, nameOverride string, symbolOverride string) (*types.Transaction, error) {
	return _Yvault030.contract.Transact(opts, "initialize", token, governance, rewards, nameOverride, symbolOverride)
}

// Initialize is a paid mutator transaction binding the contract method 0x83b43589.
//
// Solidity: function initialize(address token, address governance, address rewards, string nameOverride, string symbolOverride) returns()
func (_Yvault030 *Yvault030Session) Initialize(token common.Address, governance common.Address, rewards common.Address, nameOverride string, symbolOverride string) (*types.Transaction, error) {
	return _Yvault030.Contract.Initialize(&_Yvault030.TransactOpts, token, governance, rewards, nameOverride, symbolOverride)
}

// Initialize is a paid mutator transaction binding the contract method 0x83b43589.
//
// Solidity: function initialize(address token, address governance, address rewards, string nameOverride, string symbolOverride) returns()
func (_Yvault030 *Yvault030TransactorSession) Initialize(token common.Address, governance common.Address, rewards common.Address, nameOverride string, symbolOverride string) (*types.Transaction, error) {
	return _Yvault030.Contract.Initialize(&_Yvault030.TransactOpts, token, governance, rewards, nameOverride, symbolOverride)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xa5b81fdf.
//
// Solidity: function initialize(address token, address governance, address rewards, string nameOverride, string symbolOverride, address guardian) returns()
func (_Yvault030 *Yvault030Transactor) Initialize0(opts *bind.TransactOpts, token common.Address, governance common.Address, rewards common.Address, nameOverride string, symbolOverride string, guardian common.Address) (*types.Transaction, error) {
	return _Yvault030.contract.Transact(opts, "initialize0", token, governance, rewards, nameOverride, symbolOverride, guardian)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xa5b81fdf.
//
// Solidity: function initialize(address token, address governance, address rewards, string nameOverride, string symbolOverride, address guardian) returns()
func (_Yvault030 *Yvault030Session) Initialize0(token common.Address, governance common.Address, rewards common.Address, nameOverride string, symbolOverride string, guardian common.Address) (*types.Transaction, error) {
	return _Yvault030.Contract.Initialize0(&_Yvault030.TransactOpts, token, governance, rewards, nameOverride, symbolOverride, guardian)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xa5b81fdf.
//
// Solidity: function initialize(address token, address governance, address rewards, string nameOverride, string symbolOverride, address guardian) returns()
func (_Yvault030 *Yvault030TransactorSession) Initialize0(token common.Address, governance common.Address, rewards common.Address, nameOverride string, symbolOverride string, guardian common.Address) (*types.Transaction, error) {
	return _Yvault030.Contract.Initialize0(&_Yvault030.TransactOpts, token, governance, rewards, nameOverride, symbolOverride, guardian)
}

// MigrateStrategy is a paid mutator transaction binding the contract method 0x6cb56d19.
//
// Solidity: function migrateStrategy(address oldVersion, address newVersion) returns()
func (_Yvault030 *Yvault030Transactor) MigrateStrategy(opts *bind.TransactOpts, oldVersion common.Address, newVersion common.Address) (*types.Transaction, error) {
	return _Yvault030.contract.Transact(opts, "migrateStrategy", oldVersion, newVersion)
}

// MigrateStrategy is a paid mutator transaction binding the contract method 0x6cb56d19.
//
// Solidity: function migrateStrategy(address oldVersion, address newVersion) returns()
func (_Yvault030 *Yvault030Session) MigrateStrategy(oldVersion common.Address, newVersion common.Address) (*types.Transaction, error) {
	return _Yvault030.Contract.MigrateStrategy(&_Yvault030.TransactOpts, oldVersion, newVersion)
}

// MigrateStrategy is a paid mutator transaction binding the contract method 0x6cb56d19.
//
// Solidity: function migrateStrategy(address oldVersion, address newVersion) returns()
func (_Yvault030 *Yvault030TransactorSession) MigrateStrategy(oldVersion common.Address, newVersion common.Address) (*types.Transaction, error) {
	return _Yvault030.Contract.MigrateStrategy(&_Yvault030.TransactOpts, oldVersion, newVersion)
}

// Permit is a paid mutator transaction binding the contract method 0x9fd5a6cf.
//
// Solidity: function permit(address owner, address spender, uint256 amount, uint256 expiry, bytes signature) returns(bool)
func (_Yvault030 *Yvault030Transactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, amount *big.Int, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _Yvault030.contract.Transact(opts, "permit", owner, spender, amount, expiry, signature)
}

// Permit is a paid mutator transaction binding the contract method 0x9fd5a6cf.
//
// Solidity: function permit(address owner, address spender, uint256 amount, uint256 expiry, bytes signature) returns(bool)
func (_Yvault030 *Yvault030Session) Permit(owner common.Address, spender common.Address, amount *big.Int, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _Yvault030.Contract.Permit(&_Yvault030.TransactOpts, owner, spender, amount, expiry, signature)
}

// Permit is a paid mutator transaction binding the contract method 0x9fd5a6cf.
//
// Solidity: function permit(address owner, address spender, uint256 amount, uint256 expiry, bytes signature) returns(bool)
func (_Yvault030 *Yvault030TransactorSession) Permit(owner common.Address, spender common.Address, amount *big.Int, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _Yvault030.Contract.Permit(&_Yvault030.TransactOpts, owner, spender, amount, expiry, signature)
}

// RemoveStrategyFromQueue is a paid mutator transaction binding the contract method 0xb22439f5.
//
// Solidity: function removeStrategyFromQueue(address strategy) returns()
func (_Yvault030 *Yvault030Transactor) RemoveStrategyFromQueue(opts *bind.TransactOpts, strategy common.Address) (*types.Transaction, error) {
	return _Yvault030.contract.Transact(opts, "removeStrategyFromQueue", strategy)
}

// RemoveStrategyFromQueue is a paid mutator transaction binding the contract method 0xb22439f5.
//
// Solidity: function removeStrategyFromQueue(address strategy) returns()
func (_Yvault030 *Yvault030Session) RemoveStrategyFromQueue(strategy common.Address) (*types.Transaction, error) {
	return _Yvault030.Contract.RemoveStrategyFromQueue(&_Yvault030.TransactOpts, strategy)
}

// RemoveStrategyFromQueue is a paid mutator transaction binding the contract method 0xb22439f5.
//
// Solidity: function removeStrategyFromQueue(address strategy) returns()
func (_Yvault030 *Yvault030TransactorSession) RemoveStrategyFromQueue(strategy common.Address) (*types.Transaction, error) {
	return _Yvault030.Contract.RemoveStrategyFromQueue(&_Yvault030.TransactOpts, strategy)
}

// Report is a paid mutator transaction binding the contract method 0xa1d9bafc.
//
// Solidity: function report(uint256 gain, uint256 loss, uint256 _debtPayment) returns(uint256)
func (_Yvault030 *Yvault030Transactor) Report(opts *bind.TransactOpts, gain *big.Int, loss *big.Int, _debtPayment *big.Int) (*types.Transaction, error) {
	return _Yvault030.contract.Transact(opts, "report", gain, loss, _debtPayment)
}

// Report is a paid mutator transaction binding the contract method 0xa1d9bafc.
//
// Solidity: function report(uint256 gain, uint256 loss, uint256 _debtPayment) returns(uint256)
func (_Yvault030 *Yvault030Session) Report(gain *big.Int, loss *big.Int, _debtPayment *big.Int) (*types.Transaction, error) {
	return _Yvault030.Contract.Report(&_Yvault030.TransactOpts, gain, loss, _debtPayment)
}

// Report is a paid mutator transaction binding the contract method 0xa1d9bafc.
//
// Solidity: function report(uint256 gain, uint256 loss, uint256 _debtPayment) returns(uint256)
func (_Yvault030 *Yvault030TransactorSession) Report(gain *big.Int, loss *big.Int, _debtPayment *big.Int) (*types.Transaction, error) {
	return _Yvault030.Contract.Report(&_Yvault030.TransactOpts, gain, loss, _debtPayment)
}

// RevokeStrategy is a paid mutator transaction binding the contract method 0xa0e4af9a.
//
// Solidity: function revokeStrategy() returns()
func (_Yvault030 *Yvault030Transactor) RevokeStrategy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yvault030.contract.Transact(opts, "revokeStrategy")
}

// RevokeStrategy is a paid mutator transaction binding the contract method 0xa0e4af9a.
//
// Solidity: function revokeStrategy() returns()
func (_Yvault030 *Yvault030Session) RevokeStrategy() (*types.Transaction, error) {
	return _Yvault030.Contract.RevokeStrategy(&_Yvault030.TransactOpts)
}

// RevokeStrategy is a paid mutator transaction binding the contract method 0xa0e4af9a.
//
// Solidity: function revokeStrategy() returns()
func (_Yvault030 *Yvault030TransactorSession) RevokeStrategy() (*types.Transaction, error) {
	return _Yvault030.Contract.RevokeStrategy(&_Yvault030.TransactOpts)
}

// RevokeStrategy0 is a paid mutator transaction binding the contract method 0xbb994d48.
//
// Solidity: function revokeStrategy(address strategy) returns()
func (_Yvault030 *Yvault030Transactor) RevokeStrategy0(opts *bind.TransactOpts, strategy common.Address) (*types.Transaction, error) {
	return _Yvault030.contract.Transact(opts, "revokeStrategy0", strategy)
}

// RevokeStrategy0 is a paid mutator transaction binding the contract method 0xbb994d48.
//
// Solidity: function revokeStrategy(address strategy) returns()
func (_Yvault030 *Yvault030Session) RevokeStrategy0(strategy common.Address) (*types.Transaction, error) {
	return _Yvault030.Contract.RevokeStrategy0(&_Yvault030.TransactOpts, strategy)
}

// RevokeStrategy0 is a paid mutator transaction binding the contract method 0xbb994d48.
//
// Solidity: function revokeStrategy(address strategy) returns()
func (_Yvault030 *Yvault030TransactorSession) RevokeStrategy0(strategy common.Address) (*types.Transaction, error) {
	return _Yvault030.Contract.RevokeStrategy0(&_Yvault030.TransactOpts, strategy)
}

// SetDepositLimit is a paid mutator transaction binding the contract method 0xbdc8144b.
//
// Solidity: function setDepositLimit(uint256 limit) returns()
func (_Yvault030 *Yvault030Transactor) SetDepositLimit(opts *bind.TransactOpts, limit *big.Int) (*types.Transaction, error) {
	return _Yvault030.contract.Transact(opts, "setDepositLimit", limit)
}

// SetDepositLimit is a paid mutator transaction binding the contract method 0xbdc8144b.
//
// Solidity: function setDepositLimit(uint256 limit) returns()
func (_Yvault030 *Yvault030Session) SetDepositLimit(limit *big.Int) (*types.Transaction, error) {
	return _Yvault030.Contract.SetDepositLimit(&_Yvault030.TransactOpts, limit)
}

// SetDepositLimit is a paid mutator transaction binding the contract method 0xbdc8144b.
//
// Solidity: function setDepositLimit(uint256 limit) returns()
func (_Yvault030 *Yvault030TransactorSession) SetDepositLimit(limit *big.Int) (*types.Transaction, error) {
	return _Yvault030.Contract.SetDepositLimit(&_Yvault030.TransactOpts, limit)
}

// SetEmergencyShutdown is a paid mutator transaction binding the contract method 0x14c64402.
//
// Solidity: function setEmergencyShutdown(bool active) returns()
func (_Yvault030 *Yvault030Transactor) SetEmergencyShutdown(opts *bind.TransactOpts, active bool) (*types.Transaction, error) {
	return _Yvault030.contract.Transact(opts, "setEmergencyShutdown", active)
}

// SetEmergencyShutdown is a paid mutator transaction binding the contract method 0x14c64402.
//
// Solidity: function setEmergencyShutdown(bool active) returns()
func (_Yvault030 *Yvault030Session) SetEmergencyShutdown(active bool) (*types.Transaction, error) {
	return _Yvault030.Contract.SetEmergencyShutdown(&_Yvault030.TransactOpts, active)
}

// SetEmergencyShutdown is a paid mutator transaction binding the contract method 0x14c64402.
//
// Solidity: function setEmergencyShutdown(bool active) returns()
func (_Yvault030 *Yvault030TransactorSession) SetEmergencyShutdown(active bool) (*types.Transaction, error) {
	return _Yvault030.Contract.SetEmergencyShutdown(&_Yvault030.TransactOpts, active)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address governance) returns()
func (_Yvault030 *Yvault030Transactor) SetGovernance(opts *bind.TransactOpts, governance common.Address) (*types.Transaction, error) {
	return _Yvault030.contract.Transact(opts, "setGovernance", governance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address governance) returns()
func (_Yvault030 *Yvault030Session) SetGovernance(governance common.Address) (*types.Transaction, error) {
	return _Yvault030.Contract.SetGovernance(&_Yvault030.TransactOpts, governance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address governance) returns()
func (_Yvault030 *Yvault030TransactorSession) SetGovernance(governance common.Address) (*types.Transaction, error) {
	return _Yvault030.Contract.SetGovernance(&_Yvault030.TransactOpts, governance)
}

// SetGuardian is a paid mutator transaction binding the contract method 0x8a0dac4a.
//
// Solidity: function setGuardian(address guardian) returns()
func (_Yvault030 *Yvault030Transactor) SetGuardian(opts *bind.TransactOpts, guardian common.Address) (*types.Transaction, error) {
	return _Yvault030.contract.Transact(opts, "setGuardian", guardian)
}

// SetGuardian is a paid mutator transaction binding the contract method 0x8a0dac4a.
//
// Solidity: function setGuardian(address guardian) returns()
func (_Yvault030 *Yvault030Session) SetGuardian(guardian common.Address) (*types.Transaction, error) {
	return _Yvault030.Contract.SetGuardian(&_Yvault030.TransactOpts, guardian)
}

// SetGuardian is a paid mutator transaction binding the contract method 0x8a0dac4a.
//
// Solidity: function setGuardian(address guardian) returns()
func (_Yvault030 *Yvault030TransactorSession) SetGuardian(guardian common.Address) (*types.Transaction, error) {
	return _Yvault030.Contract.SetGuardian(&_Yvault030.TransactOpts, guardian)
}

// SetGuestList is a paid mutator transaction binding the contract method 0x0b5b78eb.
//
// Solidity: function setGuestList(address guestList) returns()
func (_Yvault030 *Yvault030Transactor) SetGuestList(opts *bind.TransactOpts, guestList common.Address) (*types.Transaction, error) {
	return _Yvault030.contract.Transact(opts, "setGuestList", guestList)
}

// SetGuestList is a paid mutator transaction binding the contract method 0x0b5b78eb.
//
// Solidity: function setGuestList(address guestList) returns()
func (_Yvault030 *Yvault030Session) SetGuestList(guestList common.Address) (*types.Transaction, error) {
	return _Yvault030.Contract.SetGuestList(&_Yvault030.TransactOpts, guestList)
}

// SetGuestList is a paid mutator transaction binding the contract method 0x0b5b78eb.
//
// Solidity: function setGuestList(address guestList) returns()
func (_Yvault030 *Yvault030TransactorSession) SetGuestList(guestList common.Address) (*types.Transaction, error) {
	return _Yvault030.Contract.SetGuestList(&_Yvault030.TransactOpts, guestList)
}

// SetManagement is a paid mutator transaction binding the contract method 0xd4a22bde.
//
// Solidity: function setManagement(address management) returns()
func (_Yvault030 *Yvault030Transactor) SetManagement(opts *bind.TransactOpts, management common.Address) (*types.Transaction, error) {
	return _Yvault030.contract.Transact(opts, "setManagement", management)
}

// SetManagement is a paid mutator transaction binding the contract method 0xd4a22bde.
//
// Solidity: function setManagement(address management) returns()
func (_Yvault030 *Yvault030Session) SetManagement(management common.Address) (*types.Transaction, error) {
	return _Yvault030.Contract.SetManagement(&_Yvault030.TransactOpts, management)
}

// SetManagement is a paid mutator transaction binding the contract method 0xd4a22bde.
//
// Solidity: function setManagement(address management) returns()
func (_Yvault030 *Yvault030TransactorSession) SetManagement(management common.Address) (*types.Transaction, error) {
	return _Yvault030.Contract.SetManagement(&_Yvault030.TransactOpts, management)
}

// SetManagementFee is a paid mutator transaction binding the contract method 0xfe56e232.
//
// Solidity: function setManagementFee(uint256 fee) returns()
func (_Yvault030 *Yvault030Transactor) SetManagementFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _Yvault030.contract.Transact(opts, "setManagementFee", fee)
}

// SetManagementFee is a paid mutator transaction binding the contract method 0xfe56e232.
//
// Solidity: function setManagementFee(uint256 fee) returns()
func (_Yvault030 *Yvault030Session) SetManagementFee(fee *big.Int) (*types.Transaction, error) {
	return _Yvault030.Contract.SetManagementFee(&_Yvault030.TransactOpts, fee)
}

// SetManagementFee is a paid mutator transaction binding the contract method 0xfe56e232.
//
// Solidity: function setManagementFee(uint256 fee) returns()
func (_Yvault030 *Yvault030TransactorSession) SetManagementFee(fee *big.Int) (*types.Transaction, error) {
	return _Yvault030.Contract.SetManagementFee(&_Yvault030.TransactOpts, fee)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name) returns()
func (_Yvault030 *Yvault030Transactor) SetName(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _Yvault030.contract.Transact(opts, "setName", name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name) returns()
func (_Yvault030 *Yvault030Session) SetName(name string) (*types.Transaction, error) {
	return _Yvault030.Contract.SetName(&_Yvault030.TransactOpts, name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name) returns()
func (_Yvault030 *Yvault030TransactorSession) SetName(name string) (*types.Transaction, error) {
	return _Yvault030.Contract.SetName(&_Yvault030.TransactOpts, name)
}

// SetPerformanceFee is a paid mutator transaction binding the contract method 0x70897b23.
//
// Solidity: function setPerformanceFee(uint256 fee) returns()
func (_Yvault030 *Yvault030Transactor) SetPerformanceFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _Yvault030.contract.Transact(opts, "setPerformanceFee", fee)
}

// SetPerformanceFee is a paid mutator transaction binding the contract method 0x70897b23.
//
// Solidity: function setPerformanceFee(uint256 fee) returns()
func (_Yvault030 *Yvault030Session) SetPerformanceFee(fee *big.Int) (*types.Transaction, error) {
	return _Yvault030.Contract.SetPerformanceFee(&_Yvault030.TransactOpts, fee)
}

// SetPerformanceFee is a paid mutator transaction binding the contract method 0x70897b23.
//
// Solidity: function setPerformanceFee(uint256 fee) returns()
func (_Yvault030 *Yvault030TransactorSession) SetPerformanceFee(fee *big.Int) (*types.Transaction, error) {
	return _Yvault030.Contract.SetPerformanceFee(&_Yvault030.TransactOpts, fee)
}

// SetRewards is a paid mutator transaction binding the contract method 0xec38a862.
//
// Solidity: function setRewards(address rewards) returns()
func (_Yvault030 *Yvault030Transactor) SetRewards(opts *bind.TransactOpts, rewards common.Address) (*types.Transaction, error) {
	return _Yvault030.contract.Transact(opts, "setRewards", rewards)
}

// SetRewards is a paid mutator transaction binding the contract method 0xec38a862.
//
// Solidity: function setRewards(address rewards) returns()
func (_Yvault030 *Yvault030Session) SetRewards(rewards common.Address) (*types.Transaction, error) {
	return _Yvault030.Contract.SetRewards(&_Yvault030.TransactOpts, rewards)
}

// SetRewards is a paid mutator transaction binding the contract method 0xec38a862.
//
// Solidity: function setRewards(address rewards) returns()
func (_Yvault030 *Yvault030TransactorSession) SetRewards(rewards common.Address) (*types.Transaction, error) {
	return _Yvault030.Contract.SetRewards(&_Yvault030.TransactOpts, rewards)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string symbol) returns()
func (_Yvault030 *Yvault030Transactor) SetSymbol(opts *bind.TransactOpts, symbol string) (*types.Transaction, error) {
	return _Yvault030.contract.Transact(opts, "setSymbol", symbol)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string symbol) returns()
func (_Yvault030 *Yvault030Session) SetSymbol(symbol string) (*types.Transaction, error) {
	return _Yvault030.Contract.SetSymbol(&_Yvault030.TransactOpts, symbol)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string symbol) returns()
func (_Yvault030 *Yvault030TransactorSession) SetSymbol(symbol string) (*types.Transaction, error) {
	return _Yvault030.Contract.SetSymbol(&_Yvault030.TransactOpts, symbol)
}

// SetWithdrawalQueue is a paid mutator transaction binding the contract method 0x94148415.
//
// Solidity: function setWithdrawalQueue(address[20] queue) returns()
func (_Yvault030 *Yvault030Transactor) SetWithdrawalQueue(opts *bind.TransactOpts, queue [20]common.Address) (*types.Transaction, error) {
	return _Yvault030.contract.Transact(opts, "setWithdrawalQueue", queue)
}

// SetWithdrawalQueue is a paid mutator transaction binding the contract method 0x94148415.
//
// Solidity: function setWithdrawalQueue(address[20] queue) returns()
func (_Yvault030 *Yvault030Session) SetWithdrawalQueue(queue [20]common.Address) (*types.Transaction, error) {
	return _Yvault030.Contract.SetWithdrawalQueue(&_Yvault030.TransactOpts, queue)
}

// SetWithdrawalQueue is a paid mutator transaction binding the contract method 0x94148415.
//
// Solidity: function setWithdrawalQueue(address[20] queue) returns()
func (_Yvault030 *Yvault030TransactorSession) SetWithdrawalQueue(queue [20]common.Address) (*types.Transaction, error) {
	return _Yvault030.Contract.SetWithdrawalQueue(&_Yvault030.TransactOpts, queue)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address token) returns()
func (_Yvault030 *Yvault030Transactor) Sweep(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Yvault030.contract.Transact(opts, "sweep", token)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address token) returns()
func (_Yvault030 *Yvault030Session) Sweep(token common.Address) (*types.Transaction, error) {
	return _Yvault030.Contract.Sweep(&_Yvault030.TransactOpts, token)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address token) returns()
func (_Yvault030 *Yvault030TransactorSession) Sweep(token common.Address) (*types.Transaction, error) {
	return _Yvault030.Contract.Sweep(&_Yvault030.TransactOpts, token)
}

// Sweep0 is a paid mutator transaction binding the contract method 0x6ea056a9.
//
// Solidity: function sweep(address token, uint256 amount) returns()
func (_Yvault030 *Yvault030Transactor) Sweep0(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault030.contract.Transact(opts, "sweep0", token, amount)
}

// Sweep0 is a paid mutator transaction binding the contract method 0x6ea056a9.
//
// Solidity: function sweep(address token, uint256 amount) returns()
func (_Yvault030 *Yvault030Session) Sweep0(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault030.Contract.Sweep0(&_Yvault030.TransactOpts, token, amount)
}

// Sweep0 is a paid mutator transaction binding the contract method 0x6ea056a9.
//
// Solidity: function sweep(address token, uint256 amount) returns()
func (_Yvault030 *Yvault030TransactorSession) Sweep0(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault030.Contract.Sweep0(&_Yvault030.TransactOpts, token, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 amount) returns(bool)
func (_Yvault030 *Yvault030Transactor) Transfer(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault030.contract.Transact(opts, "transfer", receiver, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 amount) returns(bool)
func (_Yvault030 *Yvault030Session) Transfer(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault030.Contract.Transfer(&_Yvault030.TransactOpts, receiver, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 amount) returns(bool)
func (_Yvault030 *Yvault030TransactorSession) Transfer(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault030.Contract.Transfer(&_Yvault030.TransactOpts, receiver, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address receiver, uint256 amount) returns(bool)
func (_Yvault030 *Yvault030Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault030.contract.Transact(opts, "transferFrom", sender, receiver, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address receiver, uint256 amount) returns(bool)
func (_Yvault030 *Yvault030Session) TransferFrom(sender common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault030.Contract.TransferFrom(&_Yvault030.TransactOpts, sender, receiver, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address receiver, uint256 amount) returns(bool)
func (_Yvault030 *Yvault030TransactorSession) TransferFrom(sender common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault030.Contract.TransferFrom(&_Yvault030.TransactOpts, sender, receiver, amount)
}

// UpdateStrategyDebtRatio is a paid mutator transaction binding the contract method 0x7c6a4f24.
//
// Solidity: function updateStrategyDebtRatio(address strategy, uint256 debtRatio) returns()
func (_Yvault030 *Yvault030Transactor) UpdateStrategyDebtRatio(opts *bind.TransactOpts, strategy common.Address, debtRatio *big.Int) (*types.Transaction, error) {
	return _Yvault030.contract.Transact(opts, "updateStrategyDebtRatio", strategy, debtRatio)
}

// UpdateStrategyDebtRatio is a paid mutator transaction binding the contract method 0x7c6a4f24.
//
// Solidity: function updateStrategyDebtRatio(address strategy, uint256 debtRatio) returns()
func (_Yvault030 *Yvault030Session) UpdateStrategyDebtRatio(strategy common.Address, debtRatio *big.Int) (*types.Transaction, error) {
	return _Yvault030.Contract.UpdateStrategyDebtRatio(&_Yvault030.TransactOpts, strategy, debtRatio)
}

// UpdateStrategyDebtRatio is a paid mutator transaction binding the contract method 0x7c6a4f24.
//
// Solidity: function updateStrategyDebtRatio(address strategy, uint256 debtRatio) returns()
func (_Yvault030 *Yvault030TransactorSession) UpdateStrategyDebtRatio(strategy common.Address, debtRatio *big.Int) (*types.Transaction, error) {
	return _Yvault030.Contract.UpdateStrategyDebtRatio(&_Yvault030.TransactOpts, strategy, debtRatio)
}

// UpdateStrategyPerformanceFee is a paid mutator transaction binding the contract method 0xd0194ed6.
//
// Solidity: function updateStrategyPerformanceFee(address strategy, uint256 performanceFee) returns()
func (_Yvault030 *Yvault030Transactor) UpdateStrategyPerformanceFee(opts *bind.TransactOpts, strategy common.Address, performanceFee *big.Int) (*types.Transaction, error) {
	return _Yvault030.contract.Transact(opts, "updateStrategyPerformanceFee", strategy, performanceFee)
}

// UpdateStrategyPerformanceFee is a paid mutator transaction binding the contract method 0xd0194ed6.
//
// Solidity: function updateStrategyPerformanceFee(address strategy, uint256 performanceFee) returns()
func (_Yvault030 *Yvault030Session) UpdateStrategyPerformanceFee(strategy common.Address, performanceFee *big.Int) (*types.Transaction, error) {
	return _Yvault030.Contract.UpdateStrategyPerformanceFee(&_Yvault030.TransactOpts, strategy, performanceFee)
}

// UpdateStrategyPerformanceFee is a paid mutator transaction binding the contract method 0xd0194ed6.
//
// Solidity: function updateStrategyPerformanceFee(address strategy, uint256 performanceFee) returns()
func (_Yvault030 *Yvault030TransactorSession) UpdateStrategyPerformanceFee(strategy common.Address, performanceFee *big.Int) (*types.Transaction, error) {
	return _Yvault030.Contract.UpdateStrategyPerformanceFee(&_Yvault030.TransactOpts, strategy, performanceFee)
}

// UpdateStrategyRateLimit is a paid mutator transaction binding the contract method 0x62fdbc9f.
//
// Solidity: function updateStrategyRateLimit(address strategy, uint256 rateLimit) returns()
func (_Yvault030 *Yvault030Transactor) UpdateStrategyRateLimit(opts *bind.TransactOpts, strategy common.Address, rateLimit *big.Int) (*types.Transaction, error) {
	return _Yvault030.contract.Transact(opts, "updateStrategyRateLimit", strategy, rateLimit)
}

// UpdateStrategyRateLimit is a paid mutator transaction binding the contract method 0x62fdbc9f.
//
// Solidity: function updateStrategyRateLimit(address strategy, uint256 rateLimit) returns()
func (_Yvault030 *Yvault030Session) UpdateStrategyRateLimit(strategy common.Address, rateLimit *big.Int) (*types.Transaction, error) {
	return _Yvault030.Contract.UpdateStrategyRateLimit(&_Yvault030.TransactOpts, strategy, rateLimit)
}

// UpdateStrategyRateLimit is a paid mutator transaction binding the contract method 0x62fdbc9f.
//
// Solidity: function updateStrategyRateLimit(address strategy, uint256 rateLimit) returns()
func (_Yvault030 *Yvault030TransactorSession) UpdateStrategyRateLimit(strategy common.Address, rateLimit *big.Int) (*types.Transaction, error) {
	return _Yvault030.Contract.UpdateStrategyRateLimit(&_Yvault030.TransactOpts, strategy, rateLimit)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(uint256)
func (_Yvault030 *Yvault030Transactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yvault030.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(uint256)
func (_Yvault030 *Yvault030Session) Withdraw() (*types.Transaction, error) {
	return _Yvault030.Contract.Withdraw(&_Yvault030.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(uint256)
func (_Yvault030 *Yvault030TransactorSession) Withdraw() (*types.Transaction, error) {
	return _Yvault030.Contract.Withdraw(&_Yvault030.TransactOpts)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 maxShares) returns(uint256)
func (_Yvault030 *Yvault030Transactor) Withdraw0(opts *bind.TransactOpts, maxShares *big.Int) (*types.Transaction, error) {
	return _Yvault030.contract.Transact(opts, "withdraw0", maxShares)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 maxShares) returns(uint256)
func (_Yvault030 *Yvault030Session) Withdraw0(maxShares *big.Int) (*types.Transaction, error) {
	return _Yvault030.Contract.Withdraw0(&_Yvault030.TransactOpts, maxShares)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 maxShares) returns(uint256)
func (_Yvault030 *Yvault030TransactorSession) Withdraw0(maxShares *big.Int) (*types.Transaction, error) {
	return _Yvault030.Contract.Withdraw0(&_Yvault030.TransactOpts, maxShares)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 maxShares, address recipient) returns(uint256)
func (_Yvault030 *Yvault030Transactor) Withdraw1(opts *bind.TransactOpts, maxShares *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Yvault030.contract.Transact(opts, "withdraw1", maxShares, recipient)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 maxShares, address recipient) returns(uint256)
func (_Yvault030 *Yvault030Session) Withdraw1(maxShares *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Yvault030.Contract.Withdraw1(&_Yvault030.TransactOpts, maxShares, recipient)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 maxShares, address recipient) returns(uint256)
func (_Yvault030 *Yvault030TransactorSession) Withdraw1(maxShares *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Yvault030.Contract.Withdraw1(&_Yvault030.TransactOpts, maxShares, recipient)
}

// Withdraw2 is a paid mutator transaction binding the contract method 0xe63697c8.
//
// Solidity: function withdraw(uint256 maxShares, address recipient, uint256 maxLoss) returns(uint256)
func (_Yvault030 *Yvault030Transactor) Withdraw2(opts *bind.TransactOpts, maxShares *big.Int, recipient common.Address, maxLoss *big.Int) (*types.Transaction, error) {
	return _Yvault030.contract.Transact(opts, "withdraw2", maxShares, recipient, maxLoss)
}

// Withdraw2 is a paid mutator transaction binding the contract method 0xe63697c8.
//
// Solidity: function withdraw(uint256 maxShares, address recipient, uint256 maxLoss) returns(uint256)
func (_Yvault030 *Yvault030Session) Withdraw2(maxShares *big.Int, recipient common.Address, maxLoss *big.Int) (*types.Transaction, error) {
	return _Yvault030.Contract.Withdraw2(&_Yvault030.TransactOpts, maxShares, recipient, maxLoss)
}

// Withdraw2 is a paid mutator transaction binding the contract method 0xe63697c8.
//
// Solidity: function withdraw(uint256 maxShares, address recipient, uint256 maxLoss) returns(uint256)
func (_Yvault030 *Yvault030TransactorSession) Withdraw2(maxShares *big.Int, recipient common.Address, maxLoss *big.Int) (*types.Transaction, error) {
	return _Yvault030.Contract.Withdraw2(&_Yvault030.TransactOpts, maxShares, recipient, maxLoss)
}

// Yvault030ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Yvault030 contract.
type Yvault030ApprovalIterator struct {
	Event *Yvault030Approval // Event containing the contract specifics and raw log

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
func (it *Yvault030ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault030Approval)
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
		it.Event = new(Yvault030Approval)
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
func (it *Yvault030ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault030ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault030Approval represents a Approval event raised by the Yvault030 contract.
type Yvault030Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Yvault030 *Yvault030Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Yvault030ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Yvault030.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Yvault030ApprovalIterator{contract: _Yvault030.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Yvault030 *Yvault030Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Yvault030Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Yvault030.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault030Approval)
				if err := _Yvault030.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Yvault030 *Yvault030Filterer) ParseApproval(log types.Log) (*Yvault030Approval, error) {
	event := new(Yvault030Approval)
	if err := _Yvault030.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault030EmergencyShutdownIterator is returned from FilterEmergencyShutdown and is used to iterate over the raw logs and unpacked data for EmergencyShutdown events raised by the Yvault030 contract.
type Yvault030EmergencyShutdownIterator struct {
	Event *Yvault030EmergencyShutdown // Event containing the contract specifics and raw log

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
func (it *Yvault030EmergencyShutdownIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault030EmergencyShutdown)
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
		it.Event = new(Yvault030EmergencyShutdown)
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
func (it *Yvault030EmergencyShutdownIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault030EmergencyShutdownIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault030EmergencyShutdown represents a EmergencyShutdown event raised by the Yvault030 contract.
type Yvault030EmergencyShutdown struct {
	Active bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyShutdown is a free log retrieval operation binding the contract event 0xba40372a3a724dca3c57156128ef1e896724b65b37a17f190b1ad5de68f3a4f3.
//
// Solidity: event EmergencyShutdown(bool active)
func (_Yvault030 *Yvault030Filterer) FilterEmergencyShutdown(opts *bind.FilterOpts) (*Yvault030EmergencyShutdownIterator, error) {

	logs, sub, err := _Yvault030.contract.FilterLogs(opts, "EmergencyShutdown")
	if err != nil {
		return nil, err
	}
	return &Yvault030EmergencyShutdownIterator{contract: _Yvault030.contract, event: "EmergencyShutdown", logs: logs, sub: sub}, nil
}

// WatchEmergencyShutdown is a free log subscription operation binding the contract event 0xba40372a3a724dca3c57156128ef1e896724b65b37a17f190b1ad5de68f3a4f3.
//
// Solidity: event EmergencyShutdown(bool active)
func (_Yvault030 *Yvault030Filterer) WatchEmergencyShutdown(opts *bind.WatchOpts, sink chan<- *Yvault030EmergencyShutdown) (event.Subscription, error) {

	logs, sub, err := _Yvault030.contract.WatchLogs(opts, "EmergencyShutdown")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault030EmergencyShutdown)
				if err := _Yvault030.contract.UnpackLog(event, "EmergencyShutdown", log); err != nil {
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
func (_Yvault030 *Yvault030Filterer) ParseEmergencyShutdown(log types.Log) (*Yvault030EmergencyShutdown, error) {
	event := new(Yvault030EmergencyShutdown)
	if err := _Yvault030.contract.UnpackLog(event, "EmergencyShutdown", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault030StrategyAddedIterator is returned from FilterStrategyAdded and is used to iterate over the raw logs and unpacked data for StrategyAdded events raised by the Yvault030 contract.
type Yvault030StrategyAddedIterator struct {
	Event *Yvault030StrategyAdded // Event containing the contract specifics and raw log

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
func (it *Yvault030StrategyAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault030StrategyAdded)
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
		it.Event = new(Yvault030StrategyAdded)
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
func (it *Yvault030StrategyAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault030StrategyAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault030StrategyAdded represents a StrategyAdded event raised by the Yvault030 contract.
type Yvault030StrategyAdded struct {
	Strategy       common.Address
	DebtRatio      *big.Int
	RateLimit      *big.Int
	PerformanceFee *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterStrategyAdded is a free log retrieval operation binding the contract event 0x5ec27a4fa537fc86d0d17d84e0ee3172c9d253c78cc4ab5c69ee99c5f7084f51.
//
// Solidity: event StrategyAdded(address indexed strategy, uint256 debtRatio, uint256 rateLimit, uint256 performanceFee)
func (_Yvault030 *Yvault030Filterer) FilterStrategyAdded(opts *bind.FilterOpts, strategy []common.Address) (*Yvault030StrategyAddedIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault030.contract.FilterLogs(opts, "StrategyAdded", strategyRule)
	if err != nil {
		return nil, err
	}
	return &Yvault030StrategyAddedIterator{contract: _Yvault030.contract, event: "StrategyAdded", logs: logs, sub: sub}, nil
}

// WatchStrategyAdded is a free log subscription operation binding the contract event 0x5ec27a4fa537fc86d0d17d84e0ee3172c9d253c78cc4ab5c69ee99c5f7084f51.
//
// Solidity: event StrategyAdded(address indexed strategy, uint256 debtRatio, uint256 rateLimit, uint256 performanceFee)
func (_Yvault030 *Yvault030Filterer) WatchStrategyAdded(opts *bind.WatchOpts, sink chan<- *Yvault030StrategyAdded, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault030.contract.WatchLogs(opts, "StrategyAdded", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault030StrategyAdded)
				if err := _Yvault030.contract.UnpackLog(event, "StrategyAdded", log); err != nil {
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

// ParseStrategyAdded is a log parse operation binding the contract event 0x5ec27a4fa537fc86d0d17d84e0ee3172c9d253c78cc4ab5c69ee99c5f7084f51.
//
// Solidity: event StrategyAdded(address indexed strategy, uint256 debtRatio, uint256 rateLimit, uint256 performanceFee)
func (_Yvault030 *Yvault030Filterer) ParseStrategyAdded(log types.Log) (*Yvault030StrategyAdded, error) {
	event := new(Yvault030StrategyAdded)
	if err := _Yvault030.contract.UnpackLog(event, "StrategyAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault030StrategyAddedToQueueIterator is returned from FilterStrategyAddedToQueue and is used to iterate over the raw logs and unpacked data for StrategyAddedToQueue events raised by the Yvault030 contract.
type Yvault030StrategyAddedToQueueIterator struct {
	Event *Yvault030StrategyAddedToQueue // Event containing the contract specifics and raw log

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
func (it *Yvault030StrategyAddedToQueueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault030StrategyAddedToQueue)
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
		it.Event = new(Yvault030StrategyAddedToQueue)
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
func (it *Yvault030StrategyAddedToQueueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault030StrategyAddedToQueueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault030StrategyAddedToQueue represents a StrategyAddedToQueue event raised by the Yvault030 contract.
type Yvault030StrategyAddedToQueue struct {
	Strategy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyAddedToQueue is a free log retrieval operation binding the contract event 0xa8727d412c6fa1e2497d6d6f275e2d9fe4d9318d5b793632e60ad9d38ee8f1fa.
//
// Solidity: event StrategyAddedToQueue(address indexed strategy)
func (_Yvault030 *Yvault030Filterer) FilterStrategyAddedToQueue(opts *bind.FilterOpts, strategy []common.Address) (*Yvault030StrategyAddedToQueueIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault030.contract.FilterLogs(opts, "StrategyAddedToQueue", strategyRule)
	if err != nil {
		return nil, err
	}
	return &Yvault030StrategyAddedToQueueIterator{contract: _Yvault030.contract, event: "StrategyAddedToQueue", logs: logs, sub: sub}, nil
}

// WatchStrategyAddedToQueue is a free log subscription operation binding the contract event 0xa8727d412c6fa1e2497d6d6f275e2d9fe4d9318d5b793632e60ad9d38ee8f1fa.
//
// Solidity: event StrategyAddedToQueue(address indexed strategy)
func (_Yvault030 *Yvault030Filterer) WatchStrategyAddedToQueue(opts *bind.WatchOpts, sink chan<- *Yvault030StrategyAddedToQueue, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault030.contract.WatchLogs(opts, "StrategyAddedToQueue", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault030StrategyAddedToQueue)
				if err := _Yvault030.contract.UnpackLog(event, "StrategyAddedToQueue", log); err != nil {
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
func (_Yvault030 *Yvault030Filterer) ParseStrategyAddedToQueue(log types.Log) (*Yvault030StrategyAddedToQueue, error) {
	event := new(Yvault030StrategyAddedToQueue)
	if err := _Yvault030.contract.UnpackLog(event, "StrategyAddedToQueue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault030StrategyMigratedIterator is returned from FilterStrategyMigrated and is used to iterate over the raw logs and unpacked data for StrategyMigrated events raised by the Yvault030 contract.
type Yvault030StrategyMigratedIterator struct {
	Event *Yvault030StrategyMigrated // Event containing the contract specifics and raw log

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
func (it *Yvault030StrategyMigratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault030StrategyMigrated)
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
		it.Event = new(Yvault030StrategyMigrated)
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
func (it *Yvault030StrategyMigratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault030StrategyMigratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault030StrategyMigrated represents a StrategyMigrated event raised by the Yvault030 contract.
type Yvault030StrategyMigrated struct {
	OldVersion common.Address
	NewVersion common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterStrategyMigrated is a free log retrieval operation binding the contract event 0x100b69bb6b504e1252e36b375233158edee64d071b399e2f81473a695fd1b021.
//
// Solidity: event StrategyMigrated(address indexed oldVersion, address indexed newVersion)
func (_Yvault030 *Yvault030Filterer) FilterStrategyMigrated(opts *bind.FilterOpts, oldVersion []common.Address, newVersion []common.Address) (*Yvault030StrategyMigratedIterator, error) {

	var oldVersionRule []interface{}
	for _, oldVersionItem := range oldVersion {
		oldVersionRule = append(oldVersionRule, oldVersionItem)
	}
	var newVersionRule []interface{}
	for _, newVersionItem := range newVersion {
		newVersionRule = append(newVersionRule, newVersionItem)
	}

	logs, sub, err := _Yvault030.contract.FilterLogs(opts, "StrategyMigrated", oldVersionRule, newVersionRule)
	if err != nil {
		return nil, err
	}
	return &Yvault030StrategyMigratedIterator{contract: _Yvault030.contract, event: "StrategyMigrated", logs: logs, sub: sub}, nil
}

// WatchStrategyMigrated is a free log subscription operation binding the contract event 0x100b69bb6b504e1252e36b375233158edee64d071b399e2f81473a695fd1b021.
//
// Solidity: event StrategyMigrated(address indexed oldVersion, address indexed newVersion)
func (_Yvault030 *Yvault030Filterer) WatchStrategyMigrated(opts *bind.WatchOpts, sink chan<- *Yvault030StrategyMigrated, oldVersion []common.Address, newVersion []common.Address) (event.Subscription, error) {

	var oldVersionRule []interface{}
	for _, oldVersionItem := range oldVersion {
		oldVersionRule = append(oldVersionRule, oldVersionItem)
	}
	var newVersionRule []interface{}
	for _, newVersionItem := range newVersion {
		newVersionRule = append(newVersionRule, newVersionItem)
	}

	logs, sub, err := _Yvault030.contract.WatchLogs(opts, "StrategyMigrated", oldVersionRule, newVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault030StrategyMigrated)
				if err := _Yvault030.contract.UnpackLog(event, "StrategyMigrated", log); err != nil {
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
func (_Yvault030 *Yvault030Filterer) ParseStrategyMigrated(log types.Log) (*Yvault030StrategyMigrated, error) {
	event := new(Yvault030StrategyMigrated)
	if err := _Yvault030.contract.UnpackLog(event, "StrategyMigrated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault030StrategyRemovedFromQueueIterator is returned from FilterStrategyRemovedFromQueue and is used to iterate over the raw logs and unpacked data for StrategyRemovedFromQueue events raised by the Yvault030 contract.
type Yvault030StrategyRemovedFromQueueIterator struct {
	Event *Yvault030StrategyRemovedFromQueue // Event containing the contract specifics and raw log

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
func (it *Yvault030StrategyRemovedFromQueueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault030StrategyRemovedFromQueue)
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
		it.Event = new(Yvault030StrategyRemovedFromQueue)
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
func (it *Yvault030StrategyRemovedFromQueueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault030StrategyRemovedFromQueueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault030StrategyRemovedFromQueue represents a StrategyRemovedFromQueue event raised by the Yvault030 contract.
type Yvault030StrategyRemovedFromQueue struct {
	Strategy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyRemovedFromQueue is a free log retrieval operation binding the contract event 0x8e1ec3c16d6a67ea8effe2ac7adef9c2de0bc0dc47c49cdf18f6a8b0048085be.
//
// Solidity: event StrategyRemovedFromQueue(address indexed strategy)
func (_Yvault030 *Yvault030Filterer) FilterStrategyRemovedFromQueue(opts *bind.FilterOpts, strategy []common.Address) (*Yvault030StrategyRemovedFromQueueIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault030.contract.FilterLogs(opts, "StrategyRemovedFromQueue", strategyRule)
	if err != nil {
		return nil, err
	}
	return &Yvault030StrategyRemovedFromQueueIterator{contract: _Yvault030.contract, event: "StrategyRemovedFromQueue", logs: logs, sub: sub}, nil
}

// WatchStrategyRemovedFromQueue is a free log subscription operation binding the contract event 0x8e1ec3c16d6a67ea8effe2ac7adef9c2de0bc0dc47c49cdf18f6a8b0048085be.
//
// Solidity: event StrategyRemovedFromQueue(address indexed strategy)
func (_Yvault030 *Yvault030Filterer) WatchStrategyRemovedFromQueue(opts *bind.WatchOpts, sink chan<- *Yvault030StrategyRemovedFromQueue, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault030.contract.WatchLogs(opts, "StrategyRemovedFromQueue", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault030StrategyRemovedFromQueue)
				if err := _Yvault030.contract.UnpackLog(event, "StrategyRemovedFromQueue", log); err != nil {
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
func (_Yvault030 *Yvault030Filterer) ParseStrategyRemovedFromQueue(log types.Log) (*Yvault030StrategyRemovedFromQueue, error) {
	event := new(Yvault030StrategyRemovedFromQueue)
	if err := _Yvault030.contract.UnpackLog(event, "StrategyRemovedFromQueue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault030StrategyReportedIterator is returned from FilterStrategyReported and is used to iterate over the raw logs and unpacked data for StrategyReported events raised by the Yvault030 contract.
type Yvault030StrategyReportedIterator struct {
	Event *Yvault030StrategyReported // Event containing the contract specifics and raw log

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
func (it *Yvault030StrategyReportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault030StrategyReported)
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
		it.Event = new(Yvault030StrategyReported)
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
func (it *Yvault030StrategyReportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault030StrategyReportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault030StrategyReported represents a StrategyReported event raised by the Yvault030 contract.
type Yvault030StrategyReported struct {
	Strategy  common.Address
	Gain      *big.Int
	Loss      *big.Int
	TotalGain *big.Int
	TotalLoss *big.Int
	TotalDebt *big.Int
	DebtAdded *big.Int
	DebtRatio *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStrategyReported is a free log retrieval operation binding the contract event 0x2fb611faf48b1d1b91edbba34cee10c6357adee410540e4a8f7a82b6b38673e4.
//
// Solidity: event StrategyReported(address indexed strategy, uint256 gain, uint256 loss, uint256 totalGain, uint256 totalLoss, uint256 totalDebt, uint256 debtAdded, uint256 debtRatio)
func (_Yvault030 *Yvault030Filterer) FilterStrategyReported(opts *bind.FilterOpts, strategy []common.Address) (*Yvault030StrategyReportedIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault030.contract.FilterLogs(opts, "StrategyReported", strategyRule)
	if err != nil {
		return nil, err
	}
	return &Yvault030StrategyReportedIterator{contract: _Yvault030.contract, event: "StrategyReported", logs: logs, sub: sub}, nil
}

// WatchStrategyReported is a free log subscription operation binding the contract event 0x2fb611faf48b1d1b91edbba34cee10c6357adee410540e4a8f7a82b6b38673e4.
//
// Solidity: event StrategyReported(address indexed strategy, uint256 gain, uint256 loss, uint256 totalGain, uint256 totalLoss, uint256 totalDebt, uint256 debtAdded, uint256 debtRatio)
func (_Yvault030 *Yvault030Filterer) WatchStrategyReported(opts *bind.WatchOpts, sink chan<- *Yvault030StrategyReported, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault030.contract.WatchLogs(opts, "StrategyReported", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault030StrategyReported)
				if err := _Yvault030.contract.UnpackLog(event, "StrategyReported", log); err != nil {
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

// ParseStrategyReported is a log parse operation binding the contract event 0x2fb611faf48b1d1b91edbba34cee10c6357adee410540e4a8f7a82b6b38673e4.
//
// Solidity: event StrategyReported(address indexed strategy, uint256 gain, uint256 loss, uint256 totalGain, uint256 totalLoss, uint256 totalDebt, uint256 debtAdded, uint256 debtRatio)
func (_Yvault030 *Yvault030Filterer) ParseStrategyReported(log types.Log) (*Yvault030StrategyReported, error) {
	event := new(Yvault030StrategyReported)
	if err := _Yvault030.contract.UnpackLog(event, "StrategyReported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault030StrategyRevokedIterator is returned from FilterStrategyRevoked and is used to iterate over the raw logs and unpacked data for StrategyRevoked events raised by the Yvault030 contract.
type Yvault030StrategyRevokedIterator struct {
	Event *Yvault030StrategyRevoked // Event containing the contract specifics and raw log

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
func (it *Yvault030StrategyRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault030StrategyRevoked)
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
		it.Event = new(Yvault030StrategyRevoked)
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
func (it *Yvault030StrategyRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault030StrategyRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault030StrategyRevoked represents a StrategyRevoked event raised by the Yvault030 contract.
type Yvault030StrategyRevoked struct {
	Strategy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyRevoked is a free log retrieval operation binding the contract event 0x4201c688d84c01154d321afa0c72f1bffe9eef53005c9de9d035074e71e9b32a.
//
// Solidity: event StrategyRevoked(address indexed strategy)
func (_Yvault030 *Yvault030Filterer) FilterStrategyRevoked(opts *bind.FilterOpts, strategy []common.Address) (*Yvault030StrategyRevokedIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault030.contract.FilterLogs(opts, "StrategyRevoked", strategyRule)
	if err != nil {
		return nil, err
	}
	return &Yvault030StrategyRevokedIterator{contract: _Yvault030.contract, event: "StrategyRevoked", logs: logs, sub: sub}, nil
}

// WatchStrategyRevoked is a free log subscription operation binding the contract event 0x4201c688d84c01154d321afa0c72f1bffe9eef53005c9de9d035074e71e9b32a.
//
// Solidity: event StrategyRevoked(address indexed strategy)
func (_Yvault030 *Yvault030Filterer) WatchStrategyRevoked(opts *bind.WatchOpts, sink chan<- *Yvault030StrategyRevoked, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault030.contract.WatchLogs(opts, "StrategyRevoked", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault030StrategyRevoked)
				if err := _Yvault030.contract.UnpackLog(event, "StrategyRevoked", log); err != nil {
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
func (_Yvault030 *Yvault030Filterer) ParseStrategyRevoked(log types.Log) (*Yvault030StrategyRevoked, error) {
	event := new(Yvault030StrategyRevoked)
	if err := _Yvault030.contract.UnpackLog(event, "StrategyRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault030StrategyUpdateDebtRatioIterator is returned from FilterStrategyUpdateDebtRatio and is used to iterate over the raw logs and unpacked data for StrategyUpdateDebtRatio events raised by the Yvault030 contract.
type Yvault030StrategyUpdateDebtRatioIterator struct {
	Event *Yvault030StrategyUpdateDebtRatio // Event containing the contract specifics and raw log

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
func (it *Yvault030StrategyUpdateDebtRatioIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault030StrategyUpdateDebtRatio)
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
		it.Event = new(Yvault030StrategyUpdateDebtRatio)
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
func (it *Yvault030StrategyUpdateDebtRatioIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault030StrategyUpdateDebtRatioIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault030StrategyUpdateDebtRatio represents a StrategyUpdateDebtRatio event raised by the Yvault030 contract.
type Yvault030StrategyUpdateDebtRatio struct {
	Strategy  common.Address
	DebtRatio *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStrategyUpdateDebtRatio is a free log retrieval operation binding the contract event 0xbda9398315c83ccef012bcaa318a2ff7b680f36429d36597bd4bc25ac11ead59.
//
// Solidity: event StrategyUpdateDebtRatio(address indexed strategy, uint256 debtRatio)
func (_Yvault030 *Yvault030Filterer) FilterStrategyUpdateDebtRatio(opts *bind.FilterOpts, strategy []common.Address) (*Yvault030StrategyUpdateDebtRatioIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault030.contract.FilterLogs(opts, "StrategyUpdateDebtRatio", strategyRule)
	if err != nil {
		return nil, err
	}
	return &Yvault030StrategyUpdateDebtRatioIterator{contract: _Yvault030.contract, event: "StrategyUpdateDebtRatio", logs: logs, sub: sub}, nil
}

// WatchStrategyUpdateDebtRatio is a free log subscription operation binding the contract event 0xbda9398315c83ccef012bcaa318a2ff7b680f36429d36597bd4bc25ac11ead59.
//
// Solidity: event StrategyUpdateDebtRatio(address indexed strategy, uint256 debtRatio)
func (_Yvault030 *Yvault030Filterer) WatchStrategyUpdateDebtRatio(opts *bind.WatchOpts, sink chan<- *Yvault030StrategyUpdateDebtRatio, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault030.contract.WatchLogs(opts, "StrategyUpdateDebtRatio", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault030StrategyUpdateDebtRatio)
				if err := _Yvault030.contract.UnpackLog(event, "StrategyUpdateDebtRatio", log); err != nil {
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
func (_Yvault030 *Yvault030Filterer) ParseStrategyUpdateDebtRatio(log types.Log) (*Yvault030StrategyUpdateDebtRatio, error) {
	event := new(Yvault030StrategyUpdateDebtRatio)
	if err := _Yvault030.contract.UnpackLog(event, "StrategyUpdateDebtRatio", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault030StrategyUpdatePerformanceFeeIterator is returned from FilterStrategyUpdatePerformanceFee and is used to iterate over the raw logs and unpacked data for StrategyUpdatePerformanceFee events raised by the Yvault030 contract.
type Yvault030StrategyUpdatePerformanceFeeIterator struct {
	Event *Yvault030StrategyUpdatePerformanceFee // Event containing the contract specifics and raw log

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
func (it *Yvault030StrategyUpdatePerformanceFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault030StrategyUpdatePerformanceFee)
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
		it.Event = new(Yvault030StrategyUpdatePerformanceFee)
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
func (it *Yvault030StrategyUpdatePerformanceFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault030StrategyUpdatePerformanceFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault030StrategyUpdatePerformanceFee represents a StrategyUpdatePerformanceFee event raised by the Yvault030 contract.
type Yvault030StrategyUpdatePerformanceFee struct {
	Strategy       common.Address
	PerformanceFee *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterStrategyUpdatePerformanceFee is a free log retrieval operation binding the contract event 0xe57488a65fa53066d4c25bac90db47dda4e5de3025ac12bf76ff07211cf7f39e.
//
// Solidity: event StrategyUpdatePerformanceFee(address indexed strategy, uint256 performanceFee)
func (_Yvault030 *Yvault030Filterer) FilterStrategyUpdatePerformanceFee(opts *bind.FilterOpts, strategy []common.Address) (*Yvault030StrategyUpdatePerformanceFeeIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault030.contract.FilterLogs(opts, "StrategyUpdatePerformanceFee", strategyRule)
	if err != nil {
		return nil, err
	}
	return &Yvault030StrategyUpdatePerformanceFeeIterator{contract: _Yvault030.contract, event: "StrategyUpdatePerformanceFee", logs: logs, sub: sub}, nil
}

// WatchStrategyUpdatePerformanceFee is a free log subscription operation binding the contract event 0xe57488a65fa53066d4c25bac90db47dda4e5de3025ac12bf76ff07211cf7f39e.
//
// Solidity: event StrategyUpdatePerformanceFee(address indexed strategy, uint256 performanceFee)
func (_Yvault030 *Yvault030Filterer) WatchStrategyUpdatePerformanceFee(opts *bind.WatchOpts, sink chan<- *Yvault030StrategyUpdatePerformanceFee, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault030.contract.WatchLogs(opts, "StrategyUpdatePerformanceFee", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault030StrategyUpdatePerformanceFee)
				if err := _Yvault030.contract.UnpackLog(event, "StrategyUpdatePerformanceFee", log); err != nil {
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
func (_Yvault030 *Yvault030Filterer) ParseStrategyUpdatePerformanceFee(log types.Log) (*Yvault030StrategyUpdatePerformanceFee, error) {
	event := new(Yvault030StrategyUpdatePerformanceFee)
	if err := _Yvault030.contract.UnpackLog(event, "StrategyUpdatePerformanceFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault030StrategyUpdateRateLimitIterator is returned from FilterStrategyUpdateRateLimit and is used to iterate over the raw logs and unpacked data for StrategyUpdateRateLimit events raised by the Yvault030 contract.
type Yvault030StrategyUpdateRateLimitIterator struct {
	Event *Yvault030StrategyUpdateRateLimit // Event containing the contract specifics and raw log

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
func (it *Yvault030StrategyUpdateRateLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault030StrategyUpdateRateLimit)
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
		it.Event = new(Yvault030StrategyUpdateRateLimit)
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
func (it *Yvault030StrategyUpdateRateLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault030StrategyUpdateRateLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault030StrategyUpdateRateLimit represents a StrategyUpdateRateLimit event raised by the Yvault030 contract.
type Yvault030StrategyUpdateRateLimit struct {
	Strategy  common.Address
	RateLimit *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStrategyUpdateRateLimit is a free log retrieval operation binding the contract event 0xfc0e145ae9ec0b117fcf52a2ac0fa309d2896e9ff19e6a789d1c20f425b243ea.
//
// Solidity: event StrategyUpdateRateLimit(address indexed strategy, uint256 rateLimit)
func (_Yvault030 *Yvault030Filterer) FilterStrategyUpdateRateLimit(opts *bind.FilterOpts, strategy []common.Address) (*Yvault030StrategyUpdateRateLimitIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault030.contract.FilterLogs(opts, "StrategyUpdateRateLimit", strategyRule)
	if err != nil {
		return nil, err
	}
	return &Yvault030StrategyUpdateRateLimitIterator{contract: _Yvault030.contract, event: "StrategyUpdateRateLimit", logs: logs, sub: sub}, nil
}

// WatchStrategyUpdateRateLimit is a free log subscription operation binding the contract event 0xfc0e145ae9ec0b117fcf52a2ac0fa309d2896e9ff19e6a789d1c20f425b243ea.
//
// Solidity: event StrategyUpdateRateLimit(address indexed strategy, uint256 rateLimit)
func (_Yvault030 *Yvault030Filterer) WatchStrategyUpdateRateLimit(opts *bind.WatchOpts, sink chan<- *Yvault030StrategyUpdateRateLimit, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault030.contract.WatchLogs(opts, "StrategyUpdateRateLimit", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault030StrategyUpdateRateLimit)
				if err := _Yvault030.contract.UnpackLog(event, "StrategyUpdateRateLimit", log); err != nil {
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

// ParseStrategyUpdateRateLimit is a log parse operation binding the contract event 0xfc0e145ae9ec0b117fcf52a2ac0fa309d2896e9ff19e6a789d1c20f425b243ea.
//
// Solidity: event StrategyUpdateRateLimit(address indexed strategy, uint256 rateLimit)
func (_Yvault030 *Yvault030Filterer) ParseStrategyUpdateRateLimit(log types.Log) (*Yvault030StrategyUpdateRateLimit, error) {
	event := new(Yvault030StrategyUpdateRateLimit)
	if err := _Yvault030.contract.UnpackLog(event, "StrategyUpdateRateLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault030TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Yvault030 contract.
type Yvault030TransferIterator struct {
	Event *Yvault030Transfer // Event containing the contract specifics and raw log

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
func (it *Yvault030TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault030Transfer)
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
		it.Event = new(Yvault030Transfer)
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
func (it *Yvault030TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault030TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault030Transfer represents a Transfer event raised by the Yvault030 contract.
type Yvault030Transfer struct {
	Sender   common.Address
	Receiver common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_Yvault030 *Yvault030Filterer) FilterTransfer(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*Yvault030TransferIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Yvault030.contract.FilterLogs(opts, "Transfer", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &Yvault030TransferIterator{contract: _Yvault030.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_Yvault030 *Yvault030Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Yvault030Transfer, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Yvault030.contract.WatchLogs(opts, "Transfer", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault030Transfer)
				if err := _Yvault030.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Yvault030 *Yvault030Filterer) ParseTransfer(log types.Log) (*Yvault030Transfer, error) {
	event := new(Yvault030Transfer)
	if err := _Yvault030.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault030UpdateDepositLimitIterator is returned from FilterUpdateDepositLimit and is used to iterate over the raw logs and unpacked data for UpdateDepositLimit events raised by the Yvault030 contract.
type Yvault030UpdateDepositLimitIterator struct {
	Event *Yvault030UpdateDepositLimit // Event containing the contract specifics and raw log

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
func (it *Yvault030UpdateDepositLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault030UpdateDepositLimit)
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
		it.Event = new(Yvault030UpdateDepositLimit)
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
func (it *Yvault030UpdateDepositLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault030UpdateDepositLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault030UpdateDepositLimit represents a UpdateDepositLimit event raised by the Yvault030 contract.
type Yvault030UpdateDepositLimit struct {
	DepositLimit *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateDepositLimit is a free log retrieval operation binding the contract event 0xae565aab888bca5e19e25a13db7b0c9144305bf55cb0f3f4d724f730e5acdd62.
//
// Solidity: event UpdateDepositLimit(uint256 depositLimit)
func (_Yvault030 *Yvault030Filterer) FilterUpdateDepositLimit(opts *bind.FilterOpts) (*Yvault030UpdateDepositLimitIterator, error) {

	logs, sub, err := _Yvault030.contract.FilterLogs(opts, "UpdateDepositLimit")
	if err != nil {
		return nil, err
	}
	return &Yvault030UpdateDepositLimitIterator{contract: _Yvault030.contract, event: "UpdateDepositLimit", logs: logs, sub: sub}, nil
}

// WatchUpdateDepositLimit is a free log subscription operation binding the contract event 0xae565aab888bca5e19e25a13db7b0c9144305bf55cb0f3f4d724f730e5acdd62.
//
// Solidity: event UpdateDepositLimit(uint256 depositLimit)
func (_Yvault030 *Yvault030Filterer) WatchUpdateDepositLimit(opts *bind.WatchOpts, sink chan<- *Yvault030UpdateDepositLimit) (event.Subscription, error) {

	logs, sub, err := _Yvault030.contract.WatchLogs(opts, "UpdateDepositLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault030UpdateDepositLimit)
				if err := _Yvault030.contract.UnpackLog(event, "UpdateDepositLimit", log); err != nil {
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
func (_Yvault030 *Yvault030Filterer) ParseUpdateDepositLimit(log types.Log) (*Yvault030UpdateDepositLimit, error) {
	event := new(Yvault030UpdateDepositLimit)
	if err := _Yvault030.contract.UnpackLog(event, "UpdateDepositLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault030UpdateGovernanceIterator is returned from FilterUpdateGovernance and is used to iterate over the raw logs and unpacked data for UpdateGovernance events raised by the Yvault030 contract.
type Yvault030UpdateGovernanceIterator struct {
	Event *Yvault030UpdateGovernance // Event containing the contract specifics and raw log

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
func (it *Yvault030UpdateGovernanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault030UpdateGovernance)
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
		it.Event = new(Yvault030UpdateGovernance)
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
func (it *Yvault030UpdateGovernanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault030UpdateGovernanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault030UpdateGovernance represents a UpdateGovernance event raised by the Yvault030 contract.
type Yvault030UpdateGovernance struct {
	Governance common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateGovernance is a free log retrieval operation binding the contract event 0x8d55d160c0009eb3d739442df0a3ca089ed64378bfac017e7ddad463f9815b87.
//
// Solidity: event UpdateGovernance(address governance)
func (_Yvault030 *Yvault030Filterer) FilterUpdateGovernance(opts *bind.FilterOpts) (*Yvault030UpdateGovernanceIterator, error) {

	logs, sub, err := _Yvault030.contract.FilterLogs(opts, "UpdateGovernance")
	if err != nil {
		return nil, err
	}
	return &Yvault030UpdateGovernanceIterator{contract: _Yvault030.contract, event: "UpdateGovernance", logs: logs, sub: sub}, nil
}

// WatchUpdateGovernance is a free log subscription operation binding the contract event 0x8d55d160c0009eb3d739442df0a3ca089ed64378bfac017e7ddad463f9815b87.
//
// Solidity: event UpdateGovernance(address governance)
func (_Yvault030 *Yvault030Filterer) WatchUpdateGovernance(opts *bind.WatchOpts, sink chan<- *Yvault030UpdateGovernance) (event.Subscription, error) {

	logs, sub, err := _Yvault030.contract.WatchLogs(opts, "UpdateGovernance")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault030UpdateGovernance)
				if err := _Yvault030.contract.UnpackLog(event, "UpdateGovernance", log); err != nil {
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
func (_Yvault030 *Yvault030Filterer) ParseUpdateGovernance(log types.Log) (*Yvault030UpdateGovernance, error) {
	event := new(Yvault030UpdateGovernance)
	if err := _Yvault030.contract.UnpackLog(event, "UpdateGovernance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault030UpdateGuardianIterator is returned from FilterUpdateGuardian and is used to iterate over the raw logs and unpacked data for UpdateGuardian events raised by the Yvault030 contract.
type Yvault030UpdateGuardianIterator struct {
	Event *Yvault030UpdateGuardian // Event containing the contract specifics and raw log

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
func (it *Yvault030UpdateGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault030UpdateGuardian)
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
		it.Event = new(Yvault030UpdateGuardian)
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
func (it *Yvault030UpdateGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault030UpdateGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault030UpdateGuardian represents a UpdateGuardian event raised by the Yvault030 contract.
type Yvault030UpdateGuardian struct {
	Guardian common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateGuardian is a free log retrieval operation binding the contract event 0x837b9ad138a0a1839a9637afce5306a5c13e23eb63365686843a5319a243609c.
//
// Solidity: event UpdateGuardian(address guardian)
func (_Yvault030 *Yvault030Filterer) FilterUpdateGuardian(opts *bind.FilterOpts) (*Yvault030UpdateGuardianIterator, error) {

	logs, sub, err := _Yvault030.contract.FilterLogs(opts, "UpdateGuardian")
	if err != nil {
		return nil, err
	}
	return &Yvault030UpdateGuardianIterator{contract: _Yvault030.contract, event: "UpdateGuardian", logs: logs, sub: sub}, nil
}

// WatchUpdateGuardian is a free log subscription operation binding the contract event 0x837b9ad138a0a1839a9637afce5306a5c13e23eb63365686843a5319a243609c.
//
// Solidity: event UpdateGuardian(address guardian)
func (_Yvault030 *Yvault030Filterer) WatchUpdateGuardian(opts *bind.WatchOpts, sink chan<- *Yvault030UpdateGuardian) (event.Subscription, error) {

	logs, sub, err := _Yvault030.contract.WatchLogs(opts, "UpdateGuardian")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault030UpdateGuardian)
				if err := _Yvault030.contract.UnpackLog(event, "UpdateGuardian", log); err != nil {
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
func (_Yvault030 *Yvault030Filterer) ParseUpdateGuardian(log types.Log) (*Yvault030UpdateGuardian, error) {
	event := new(Yvault030UpdateGuardian)
	if err := _Yvault030.contract.UnpackLog(event, "UpdateGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault030UpdateGuestListIterator is returned from FilterUpdateGuestList and is used to iterate over the raw logs and unpacked data for UpdateGuestList events raised by the Yvault030 contract.
type Yvault030UpdateGuestListIterator struct {
	Event *Yvault030UpdateGuestList // Event containing the contract specifics and raw log

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
func (it *Yvault030UpdateGuestListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault030UpdateGuestList)
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
		it.Event = new(Yvault030UpdateGuestList)
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
func (it *Yvault030UpdateGuestListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault030UpdateGuestListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault030UpdateGuestList represents a UpdateGuestList event raised by the Yvault030 contract.
type Yvault030UpdateGuestList struct {
	GuestList common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateGuestList is a free log retrieval operation binding the contract event 0x6d674c311329fb38bbc96dc33d2aad03b9bf9fcfdd8f5e5054fda291a5b3c1f8.
//
// Solidity: event UpdateGuestList(address guestList)
func (_Yvault030 *Yvault030Filterer) FilterUpdateGuestList(opts *bind.FilterOpts) (*Yvault030UpdateGuestListIterator, error) {

	logs, sub, err := _Yvault030.contract.FilterLogs(opts, "UpdateGuestList")
	if err != nil {
		return nil, err
	}
	return &Yvault030UpdateGuestListIterator{contract: _Yvault030.contract, event: "UpdateGuestList", logs: logs, sub: sub}, nil
}

// WatchUpdateGuestList is a free log subscription operation binding the contract event 0x6d674c311329fb38bbc96dc33d2aad03b9bf9fcfdd8f5e5054fda291a5b3c1f8.
//
// Solidity: event UpdateGuestList(address guestList)
func (_Yvault030 *Yvault030Filterer) WatchUpdateGuestList(opts *bind.WatchOpts, sink chan<- *Yvault030UpdateGuestList) (event.Subscription, error) {

	logs, sub, err := _Yvault030.contract.WatchLogs(opts, "UpdateGuestList")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault030UpdateGuestList)
				if err := _Yvault030.contract.UnpackLog(event, "UpdateGuestList", log); err != nil {
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
func (_Yvault030 *Yvault030Filterer) ParseUpdateGuestList(log types.Log) (*Yvault030UpdateGuestList, error) {
	event := new(Yvault030UpdateGuestList)
	if err := _Yvault030.contract.UnpackLog(event, "UpdateGuestList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault030UpdateManagementIterator is returned from FilterUpdateManagement and is used to iterate over the raw logs and unpacked data for UpdateManagement events raised by the Yvault030 contract.
type Yvault030UpdateManagementIterator struct {
	Event *Yvault030UpdateManagement // Event containing the contract specifics and raw log

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
func (it *Yvault030UpdateManagementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault030UpdateManagement)
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
		it.Event = new(Yvault030UpdateManagement)
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
func (it *Yvault030UpdateManagementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault030UpdateManagementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault030UpdateManagement represents a UpdateManagement event raised by the Yvault030 contract.
type Yvault030UpdateManagement struct {
	Management common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateManagement is a free log retrieval operation binding the contract event 0xff54978127edd34aec0f9061fb3b155fbe0ededdfa881ee3e0d541d3a1eef438.
//
// Solidity: event UpdateManagement(address management)
func (_Yvault030 *Yvault030Filterer) FilterUpdateManagement(opts *bind.FilterOpts) (*Yvault030UpdateManagementIterator, error) {

	logs, sub, err := _Yvault030.contract.FilterLogs(opts, "UpdateManagement")
	if err != nil {
		return nil, err
	}
	return &Yvault030UpdateManagementIterator{contract: _Yvault030.contract, event: "UpdateManagement", logs: logs, sub: sub}, nil
}

// WatchUpdateManagement is a free log subscription operation binding the contract event 0xff54978127edd34aec0f9061fb3b155fbe0ededdfa881ee3e0d541d3a1eef438.
//
// Solidity: event UpdateManagement(address management)
func (_Yvault030 *Yvault030Filterer) WatchUpdateManagement(opts *bind.WatchOpts, sink chan<- *Yvault030UpdateManagement) (event.Subscription, error) {

	logs, sub, err := _Yvault030.contract.WatchLogs(opts, "UpdateManagement")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault030UpdateManagement)
				if err := _Yvault030.contract.UnpackLog(event, "UpdateManagement", log); err != nil {
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
func (_Yvault030 *Yvault030Filterer) ParseUpdateManagement(log types.Log) (*Yvault030UpdateManagement, error) {
	event := new(Yvault030UpdateManagement)
	if err := _Yvault030.contract.UnpackLog(event, "UpdateManagement", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault030UpdateManagementFeeIterator is returned from FilterUpdateManagementFee and is used to iterate over the raw logs and unpacked data for UpdateManagementFee events raised by the Yvault030 contract.
type Yvault030UpdateManagementFeeIterator struct {
	Event *Yvault030UpdateManagementFee // Event containing the contract specifics and raw log

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
func (it *Yvault030UpdateManagementFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault030UpdateManagementFee)
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
		it.Event = new(Yvault030UpdateManagementFee)
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
func (it *Yvault030UpdateManagementFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault030UpdateManagementFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault030UpdateManagementFee represents a UpdateManagementFee event raised by the Yvault030 contract.
type Yvault030UpdateManagementFee struct {
	ManagementFee *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdateManagementFee is a free log retrieval operation binding the contract event 0x7a7883b0074f96e2c7fab65eb25abf624c488761a5db889e3bb84855dcc6daaf.
//
// Solidity: event UpdateManagementFee(uint256 managementFee)
func (_Yvault030 *Yvault030Filterer) FilterUpdateManagementFee(opts *bind.FilterOpts) (*Yvault030UpdateManagementFeeIterator, error) {

	logs, sub, err := _Yvault030.contract.FilterLogs(opts, "UpdateManagementFee")
	if err != nil {
		return nil, err
	}
	return &Yvault030UpdateManagementFeeIterator{contract: _Yvault030.contract, event: "UpdateManagementFee", logs: logs, sub: sub}, nil
}

// WatchUpdateManagementFee is a free log subscription operation binding the contract event 0x7a7883b0074f96e2c7fab65eb25abf624c488761a5db889e3bb84855dcc6daaf.
//
// Solidity: event UpdateManagementFee(uint256 managementFee)
func (_Yvault030 *Yvault030Filterer) WatchUpdateManagementFee(opts *bind.WatchOpts, sink chan<- *Yvault030UpdateManagementFee) (event.Subscription, error) {

	logs, sub, err := _Yvault030.contract.WatchLogs(opts, "UpdateManagementFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault030UpdateManagementFee)
				if err := _Yvault030.contract.UnpackLog(event, "UpdateManagementFee", log); err != nil {
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
func (_Yvault030 *Yvault030Filterer) ParseUpdateManagementFee(log types.Log) (*Yvault030UpdateManagementFee, error) {
	event := new(Yvault030UpdateManagementFee)
	if err := _Yvault030.contract.UnpackLog(event, "UpdateManagementFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault030UpdatePerformanceFeeIterator is returned from FilterUpdatePerformanceFee and is used to iterate over the raw logs and unpacked data for UpdatePerformanceFee events raised by the Yvault030 contract.
type Yvault030UpdatePerformanceFeeIterator struct {
	Event *Yvault030UpdatePerformanceFee // Event containing the contract specifics and raw log

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
func (it *Yvault030UpdatePerformanceFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault030UpdatePerformanceFee)
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
		it.Event = new(Yvault030UpdatePerformanceFee)
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
func (it *Yvault030UpdatePerformanceFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault030UpdatePerformanceFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault030UpdatePerformanceFee represents a UpdatePerformanceFee event raised by the Yvault030 contract.
type Yvault030UpdatePerformanceFee struct {
	PerformanceFee *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpdatePerformanceFee is a free log retrieval operation binding the contract event 0x0810a1c261ca2c0cd86a0152c51c43ba9dc329639d2349f98140891b2ea798eb.
//
// Solidity: event UpdatePerformanceFee(uint256 performanceFee)
func (_Yvault030 *Yvault030Filterer) FilterUpdatePerformanceFee(opts *bind.FilterOpts) (*Yvault030UpdatePerformanceFeeIterator, error) {

	logs, sub, err := _Yvault030.contract.FilterLogs(opts, "UpdatePerformanceFee")
	if err != nil {
		return nil, err
	}
	return &Yvault030UpdatePerformanceFeeIterator{contract: _Yvault030.contract, event: "UpdatePerformanceFee", logs: logs, sub: sub}, nil
}

// WatchUpdatePerformanceFee is a free log subscription operation binding the contract event 0x0810a1c261ca2c0cd86a0152c51c43ba9dc329639d2349f98140891b2ea798eb.
//
// Solidity: event UpdatePerformanceFee(uint256 performanceFee)
func (_Yvault030 *Yvault030Filterer) WatchUpdatePerformanceFee(opts *bind.WatchOpts, sink chan<- *Yvault030UpdatePerformanceFee) (event.Subscription, error) {

	logs, sub, err := _Yvault030.contract.WatchLogs(opts, "UpdatePerformanceFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault030UpdatePerformanceFee)
				if err := _Yvault030.contract.UnpackLog(event, "UpdatePerformanceFee", log); err != nil {
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
func (_Yvault030 *Yvault030Filterer) ParseUpdatePerformanceFee(log types.Log) (*Yvault030UpdatePerformanceFee, error) {
	event := new(Yvault030UpdatePerformanceFee)
	if err := _Yvault030.contract.UnpackLog(event, "UpdatePerformanceFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault030UpdateRewardsIterator is returned from FilterUpdateRewards and is used to iterate over the raw logs and unpacked data for UpdateRewards events raised by the Yvault030 contract.
type Yvault030UpdateRewardsIterator struct {
	Event *Yvault030UpdateRewards // Event containing the contract specifics and raw log

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
func (it *Yvault030UpdateRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault030UpdateRewards)
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
		it.Event = new(Yvault030UpdateRewards)
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
func (it *Yvault030UpdateRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault030UpdateRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault030UpdateRewards represents a UpdateRewards event raised by the Yvault030 contract.
type Yvault030UpdateRewards struct {
	Rewards common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateRewards is a free log retrieval operation binding the contract event 0xdf3c41a916aecbf42361a147f8348c242662c3ce20ecef30e826b80642477a3d.
//
// Solidity: event UpdateRewards(address rewards)
func (_Yvault030 *Yvault030Filterer) FilterUpdateRewards(opts *bind.FilterOpts) (*Yvault030UpdateRewardsIterator, error) {

	logs, sub, err := _Yvault030.contract.FilterLogs(opts, "UpdateRewards")
	if err != nil {
		return nil, err
	}
	return &Yvault030UpdateRewardsIterator{contract: _Yvault030.contract, event: "UpdateRewards", logs: logs, sub: sub}, nil
}

// WatchUpdateRewards is a free log subscription operation binding the contract event 0xdf3c41a916aecbf42361a147f8348c242662c3ce20ecef30e826b80642477a3d.
//
// Solidity: event UpdateRewards(address rewards)
func (_Yvault030 *Yvault030Filterer) WatchUpdateRewards(opts *bind.WatchOpts, sink chan<- *Yvault030UpdateRewards) (event.Subscription, error) {

	logs, sub, err := _Yvault030.contract.WatchLogs(opts, "UpdateRewards")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault030UpdateRewards)
				if err := _Yvault030.contract.UnpackLog(event, "UpdateRewards", log); err != nil {
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
func (_Yvault030 *Yvault030Filterer) ParseUpdateRewards(log types.Log) (*Yvault030UpdateRewards, error) {
	event := new(Yvault030UpdateRewards)
	if err := _Yvault030.contract.UnpackLog(event, "UpdateRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault030UpdateWithdrawalQueueIterator is returned from FilterUpdateWithdrawalQueue and is used to iterate over the raw logs and unpacked data for UpdateWithdrawalQueue events raised by the Yvault030 contract.
type Yvault030UpdateWithdrawalQueueIterator struct {
	Event *Yvault030UpdateWithdrawalQueue // Event containing the contract specifics and raw log

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
func (it *Yvault030UpdateWithdrawalQueueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault030UpdateWithdrawalQueue)
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
		it.Event = new(Yvault030UpdateWithdrawalQueue)
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
func (it *Yvault030UpdateWithdrawalQueueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault030UpdateWithdrawalQueueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault030UpdateWithdrawalQueue represents a UpdateWithdrawalQueue event raised by the Yvault030 contract.
type Yvault030UpdateWithdrawalQueue struct {
	Queue [20]common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdateWithdrawalQueue is a free log retrieval operation binding the contract event 0x695ac3ac73f08f2002284ffe563cefe798ee2878a5e04219522e2e99eb89d168.
//
// Solidity: event UpdateWithdrawalQueue(address[20] queue)
func (_Yvault030 *Yvault030Filterer) FilterUpdateWithdrawalQueue(opts *bind.FilterOpts) (*Yvault030UpdateWithdrawalQueueIterator, error) {

	logs, sub, err := _Yvault030.contract.FilterLogs(opts, "UpdateWithdrawalQueue")
	if err != nil {
		return nil, err
	}
	return &Yvault030UpdateWithdrawalQueueIterator{contract: _Yvault030.contract, event: "UpdateWithdrawalQueue", logs: logs, sub: sub}, nil
}

// WatchUpdateWithdrawalQueue is a free log subscription operation binding the contract event 0x695ac3ac73f08f2002284ffe563cefe798ee2878a5e04219522e2e99eb89d168.
//
// Solidity: event UpdateWithdrawalQueue(address[20] queue)
func (_Yvault030 *Yvault030Filterer) WatchUpdateWithdrawalQueue(opts *bind.WatchOpts, sink chan<- *Yvault030UpdateWithdrawalQueue) (event.Subscription, error) {

	logs, sub, err := _Yvault030.contract.WatchLogs(opts, "UpdateWithdrawalQueue")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault030UpdateWithdrawalQueue)
				if err := _Yvault030.contract.UnpackLog(event, "UpdateWithdrawalQueue", log); err != nil {
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
func (_Yvault030 *Yvault030Filterer) ParseUpdateWithdrawalQueue(log types.Log) (*Yvault030UpdateWithdrawalQueue, error) {
	event := new(Yvault030UpdateWithdrawalQueue)
	if err := _Yvault030.contract.UnpackLog(event, "UpdateWithdrawalQueue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
