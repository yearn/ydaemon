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

// Yvault034MetaData contains all meta data concerning the Yvault034 contract.
var Yvault034MetaData = &bind.MetaData{
	ABI: "[{\"name\":\"Transfer\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StrategyAdded\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":true},{\"name\":\"debtRatio\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"minDebtPerHarvest\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"maxDebtPerHarvest\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"performanceFee\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StrategyReported\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":true},{\"name\":\"gain\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"loss\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"debtPaid\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"totalGain\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"totalLoss\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"totalDebt\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"debtAdded\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"debtRatio\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateGovernance\",\"inputs\":[{\"name\":\"governance\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateManagement\",\"inputs\":[{\"name\":\"management\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateGuestList\",\"inputs\":[{\"name\":\"guestList\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateRewards\",\"inputs\":[{\"name\":\"rewards\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateDepositLimit\",\"inputs\":[{\"name\":\"depositLimit\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdatePerformanceFee\",\"inputs\":[{\"name\":\"performanceFee\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateManagementFee\",\"inputs\":[{\"name\":\"managementFee\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateGuardian\",\"inputs\":[{\"name\":\"guardian\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"EmergencyShutdown\",\"inputs\":[{\"name\":\"active\",\"type\":\"bool\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateWithdrawalQueue\",\"inputs\":[{\"name\":\"queue\",\"type\":\"address[20]\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StrategyUpdateDebtRatio\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":true},{\"name\":\"debtRatio\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StrategyUpdateMinDebtPerHarvest\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":true},{\"name\":\"minDebtPerHarvest\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StrategyUpdateMaxDebtPerHarvest\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":true},{\"name\":\"maxDebtPerHarvest\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StrategyUpdatePerformanceFee\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":true},{\"name\":\"performanceFee\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StrategyMigrated\",\"inputs\":[{\"name\":\"oldVersion\",\"type\":\"address\",\"indexed\":true},{\"name\":\"newVersion\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StrategyRevoked\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StrategyRemovedFromQueue\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StrategyAddedToQueue\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"governance\",\"type\":\"address\"},{\"name\":\"rewards\",\"type\":\"address\"},{\"name\":\"nameOverride\",\"type\":\"string\"},{\"name\":\"symbolOverride\",\"type\":\"string\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"governance\",\"type\":\"address\"},{\"name\":\"rewards\",\"type\":\"address\"},{\"name\":\"nameOverride\",\"type\":\"string\"},{\"name\":\"symbolOverride\",\"type\":\"string\"},{\"name\":\"guardian\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"pure\",\"type\":\"function\",\"name\":\"apiVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"gas\":4546},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"setName\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"outputs\":[],\"gas\":107044},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"setSymbol\",\"inputs\":[{\"name\":\"symbol\",\"type\":\"string\"}],\"outputs\":[],\"gas\":71894},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"setGovernance\",\"inputs\":[{\"name\":\"governance\",\"type\":\"address\"}],\"outputs\":[],\"gas\":36365},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"acceptGovernance\",\"inputs\":[],\"outputs\":[],\"gas\":37637},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"setManagement\",\"inputs\":[{\"name\":\"management\",\"type\":\"address\"}],\"outputs\":[],\"gas\":37775},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"setGuestList\",\"inputs\":[{\"name\":\"guestList\",\"type\":\"address\"}],\"outputs\":[],\"gas\":37805},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"setRewards\",\"inputs\":[{\"name\":\"rewards\",\"type\":\"address\"}],\"outputs\":[],\"gas\":37835},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"setLockedProfitDegration\",\"inputs\":[{\"name\":\"degration\",\"type\":\"uint256\"}],\"outputs\":[],\"gas\":36519},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"setDepositLimit\",\"inputs\":[{\"name\":\"limit\",\"type\":\"uint256\"}],\"outputs\":[],\"gas\":37795},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"setPerformanceFee\",\"inputs\":[{\"name\":\"fee\",\"type\":\"uint256\"}],\"outputs\":[],\"gas\":37929},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"setManagementFee\",\"inputs\":[{\"name\":\"fee\",\"type\":\"uint256\"}],\"outputs\":[],\"gas\":37959},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"setGuardian\",\"inputs\":[{\"name\":\"guardian\",\"type\":\"address\"}],\"outputs\":[],\"gas\":39203},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"setEmergencyShutdown\",\"inputs\":[{\"name\":\"active\",\"type\":\"bool\"}],\"outputs\":[],\"gas\":39274},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"setWithdrawalQueue\",\"inputs\":[{\"name\":\"queue\",\"type\":\"address[20]\"}],\"outputs\":[],\"gas\":763950},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":76768},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":116531},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":38271},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"increaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":40312},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"decreaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":40336},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"permit\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":81264},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"totalAssets\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":4098},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"maxAvailableShares\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":366010},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"maxShares\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"maxShares\",\"type\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"maxShares\",\"type\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"maxLoss\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"pricePerShare\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":17041},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"addStrategy\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"},{\"name\":\"debtRatio\",\"type\":\"uint256\"},{\"name\":\"minDebtPerHarvest\",\"type\":\"uint256\"},{\"name\":\"maxDebtPerHarvest\",\"type\":\"uint256\"},{\"name\":\"performanceFee\",\"type\":\"uint256\"}],\"outputs\":[],\"gas\":1485796},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"updateStrategyDebtRatio\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"},{\"name\":\"debtRatio\",\"type\":\"uint256\"}],\"outputs\":[],\"gas\":115193},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"updateStrategyMinDebtPerHarvest\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"},{\"name\":\"minDebtPerHarvest\",\"type\":\"uint256\"}],\"outputs\":[],\"gas\":42441},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"updateStrategyMaxDebtPerHarvest\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"},{\"name\":\"maxDebtPerHarvest\",\"type\":\"uint256\"}],\"outputs\":[],\"gas\":42471},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"updateStrategyPerformanceFee\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"},{\"name\":\"performanceFee\",\"type\":\"uint256\"}],\"outputs\":[],\"gas\":41251},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"migrateStrategy\",\"inputs\":[{\"name\":\"oldVersion\",\"type\":\"address\"},{\"name\":\"newVersion\",\"type\":\"address\"}],\"outputs\":[],\"gas\":1141468},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"revokeStrategy\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"revokeStrategy\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"addStrategyToQueue\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"}],\"outputs\":[],\"gas\":1199804},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"removeStrategyFromQueue\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"}],\"outputs\":[],\"gas\":23088703},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"debtOutstanding\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"debtOutstanding\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"creditAvailable\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"creditAvailable\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"availableDepositLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":9551},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"expectedReturn\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"expectedReturn\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"report\",\"inputs\":[{\"name\":\"gain\",\"type\":\"uint256\"},{\"name\":\"loss\",\"type\":\"uint256\"},{\"name\":\"_debtPayment\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":1116120},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"sweep\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"sweep\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"gas\":8750},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"gas\":7803},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":2408},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":2653},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"},{\"name\":\"arg1\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":2898},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":2498},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":2528},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"governance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":2558},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"management\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":2588},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"guardian\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":2618},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"guestList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":2648},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"strategies\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"performanceFee\",\"type\":\"uint256\"},{\"name\":\"activation\",\"type\":\"uint256\"},{\"name\":\"debtRatio\",\"type\":\"uint256\"},{\"name\":\"minDebtPerHarvest\",\"type\":\"uint256\"},{\"name\":\"maxDebtPerHarvest\",\"type\":\"uint256\"},{\"name\":\"lastReport\",\"type\":\"uint256\"},{\"name\":\"totalDebt\",\"type\":\"uint256\"},{\"name\":\"totalGain\",\"type\":\"uint256\"},{\"name\":\"totalLoss\",\"type\":\"uint256\"}],\"gas\":11001},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"withdrawalQueue\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":2817},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"emergencyShutdown\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":2738},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"depositLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":2768},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"debtRatio\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":2798},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"totalDebt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":2828},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"delegatedAssets\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":2858},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"lastReport\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":2888},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"activation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":2918},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"lockedProfit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":2948},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"lockedProfitDegration\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":2978},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"rewards\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3008},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"managementFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3038},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"performanceFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3068},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3313},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"DOMAIN_SEPARATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"gas\":3128}]",
}

// Yvault034ABI is the input ABI used to generate the binding from.
// Deprecated: Use Yvault034MetaData.ABI instead.
var Yvault034ABI = Yvault034MetaData.ABI

// Yvault034 is an auto generated Go binding around an Ethereum contract.
type Yvault034 struct {
	Yvault034Caller     // Read-only binding to the contract
	Yvault034Transactor // Write-only binding to the contract
	Yvault034Filterer   // Log filterer for contract events
}

// Yvault034Caller is an auto generated read-only Go binding around an Ethereum contract.
type Yvault034Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Yvault034Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Yvault034Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Yvault034Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Yvault034Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Yvault034Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Yvault034Session struct {
	Contract     *Yvault034        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Yvault034CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Yvault034CallerSession struct {
	Contract *Yvault034Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// Yvault034TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Yvault034TransactorSession struct {
	Contract     *Yvault034Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// Yvault034Raw is an auto generated low-level Go binding around an Ethereum contract.
type Yvault034Raw struct {
	Contract *Yvault034 // Generic contract binding to access the raw methods on
}

// Yvault034CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Yvault034CallerRaw struct {
	Contract *Yvault034Caller // Generic read-only contract binding to access the raw methods on
}

// Yvault034TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Yvault034TransactorRaw struct {
	Contract *Yvault034Transactor // Generic write-only contract binding to access the raw methods on
}

// NewYvault034 creates a new instance of Yvault034, bound to a specific deployed contract.
func NewYvault034(address common.Address, backend bind.ContractBackend) (*Yvault034, error) {
	contract, err := bindYvault034(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Yvault034{Yvault034Caller: Yvault034Caller{contract: contract}, Yvault034Transactor: Yvault034Transactor{contract: contract}, Yvault034Filterer: Yvault034Filterer{contract: contract}}, nil
}

// NewYvault034Caller creates a new read-only instance of Yvault034, bound to a specific deployed contract.
func NewYvault034Caller(address common.Address, caller bind.ContractCaller) (*Yvault034Caller, error) {
	contract, err := bindYvault034(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Yvault034Caller{contract: contract}, nil
}

// NewYvault034Transactor creates a new write-only instance of Yvault034, bound to a specific deployed contract.
func NewYvault034Transactor(address common.Address, transactor bind.ContractTransactor) (*Yvault034Transactor, error) {
	contract, err := bindYvault034(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Yvault034Transactor{contract: contract}, nil
}

// NewYvault034Filterer creates a new log filterer instance of Yvault034, bound to a specific deployed contract.
func NewYvault034Filterer(address common.Address, filterer bind.ContractFilterer) (*Yvault034Filterer, error) {
	contract, err := bindYvault034(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Yvault034Filterer{contract: contract}, nil
}

// bindYvault034 binds a generic wrapper to an already deployed contract.
func bindYvault034(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Yvault034ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Yvault034 *Yvault034Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Yvault034.Contract.Yvault034Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Yvault034 *Yvault034Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yvault034.Contract.Yvault034Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Yvault034 *Yvault034Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Yvault034.Contract.Yvault034Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Yvault034 *Yvault034CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Yvault034.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Yvault034 *Yvault034TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yvault034.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Yvault034 *Yvault034TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Yvault034.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Yvault034 *Yvault034Caller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Yvault034.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Yvault034 *Yvault034Session) DOMAINSEPARATOR() ([32]byte, error) {
	return _Yvault034.Contract.DOMAINSEPARATOR(&_Yvault034.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Yvault034 *Yvault034CallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Yvault034.Contract.DOMAINSEPARATOR(&_Yvault034.CallOpts)
}

// Activation is a free data retrieval call binding the contract method 0x3629c8de.
//
// Solidity: function activation() view returns(uint256)
func (_Yvault034 *Yvault034Caller) Activation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault034.contract.Call(opts, &out, "activation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Activation is a free data retrieval call binding the contract method 0x3629c8de.
//
// Solidity: function activation() view returns(uint256)
func (_Yvault034 *Yvault034Session) Activation() (*big.Int, error) {
	return _Yvault034.Contract.Activation(&_Yvault034.CallOpts)
}

// Activation is a free data retrieval call binding the contract method 0x3629c8de.
//
// Solidity: function activation() view returns(uint256)
func (_Yvault034 *Yvault034CallerSession) Activation() (*big.Int, error) {
	return _Yvault034.Contract.Activation(&_Yvault034.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_Yvault034 *Yvault034Caller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yvault034.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_Yvault034 *Yvault034Session) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Yvault034.Contract.Allowance(&_Yvault034.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_Yvault034 *Yvault034CallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Yvault034.Contract.Allowance(&_Yvault034.CallOpts, arg0, arg1)
}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() pure returns(string)
func (_Yvault034 *Yvault034Caller) ApiVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Yvault034.contract.Call(opts, &out, "apiVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() pure returns(string)
func (_Yvault034 *Yvault034Session) ApiVersion() (string, error) {
	return _Yvault034.Contract.ApiVersion(&_Yvault034.CallOpts)
}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() pure returns(string)
func (_Yvault034 *Yvault034CallerSession) ApiVersion() (string, error) {
	return _Yvault034.Contract.ApiVersion(&_Yvault034.CallOpts)
}

// AvailableDepositLimit is a free data retrieval call binding the contract method 0x153c27c4.
//
// Solidity: function availableDepositLimit() view returns(uint256)
func (_Yvault034 *Yvault034Caller) AvailableDepositLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault034.contract.Call(opts, &out, "availableDepositLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailableDepositLimit is a free data retrieval call binding the contract method 0x153c27c4.
//
// Solidity: function availableDepositLimit() view returns(uint256)
func (_Yvault034 *Yvault034Session) AvailableDepositLimit() (*big.Int, error) {
	return _Yvault034.Contract.AvailableDepositLimit(&_Yvault034.CallOpts)
}

// AvailableDepositLimit is a free data retrieval call binding the contract method 0x153c27c4.
//
// Solidity: function availableDepositLimit() view returns(uint256)
func (_Yvault034 *Yvault034CallerSession) AvailableDepositLimit() (*big.Int, error) {
	return _Yvault034.Contract.AvailableDepositLimit(&_Yvault034.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_Yvault034 *Yvault034Caller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yvault034.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_Yvault034 *Yvault034Session) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Yvault034.Contract.BalanceOf(&_Yvault034.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_Yvault034 *Yvault034CallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Yvault034.Contract.BalanceOf(&_Yvault034.CallOpts, arg0)
}

// CreditAvailable is a free data retrieval call binding the contract method 0x112c1f9b.
//
// Solidity: function creditAvailable() view returns(uint256)
func (_Yvault034 *Yvault034Caller) CreditAvailable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault034.contract.Call(opts, &out, "creditAvailable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreditAvailable is a free data retrieval call binding the contract method 0x112c1f9b.
//
// Solidity: function creditAvailable() view returns(uint256)
func (_Yvault034 *Yvault034Session) CreditAvailable() (*big.Int, error) {
	return _Yvault034.Contract.CreditAvailable(&_Yvault034.CallOpts)
}

// CreditAvailable is a free data retrieval call binding the contract method 0x112c1f9b.
//
// Solidity: function creditAvailable() view returns(uint256)
func (_Yvault034 *Yvault034CallerSession) CreditAvailable() (*big.Int, error) {
	return _Yvault034.Contract.CreditAvailable(&_Yvault034.CallOpts)
}

// CreditAvailable0 is a free data retrieval call binding the contract method 0xd7648013.
//
// Solidity: function creditAvailable(address strategy) view returns(uint256)
func (_Yvault034 *Yvault034Caller) CreditAvailable0(opts *bind.CallOpts, strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yvault034.contract.Call(opts, &out, "creditAvailable0", strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreditAvailable0 is a free data retrieval call binding the contract method 0xd7648013.
//
// Solidity: function creditAvailable(address strategy) view returns(uint256)
func (_Yvault034 *Yvault034Session) CreditAvailable0(strategy common.Address) (*big.Int, error) {
	return _Yvault034.Contract.CreditAvailable0(&_Yvault034.CallOpts, strategy)
}

// CreditAvailable0 is a free data retrieval call binding the contract method 0xd7648013.
//
// Solidity: function creditAvailable(address strategy) view returns(uint256)
func (_Yvault034 *Yvault034CallerSession) CreditAvailable0(strategy common.Address) (*big.Int, error) {
	return _Yvault034.Contract.CreditAvailable0(&_Yvault034.CallOpts, strategy)
}

// DebtOutstanding is a free data retrieval call binding the contract method 0xbf3759b5.
//
// Solidity: function debtOutstanding() view returns(uint256)
func (_Yvault034 *Yvault034Caller) DebtOutstanding(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault034.contract.Call(opts, &out, "debtOutstanding")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtOutstanding is a free data retrieval call binding the contract method 0xbf3759b5.
//
// Solidity: function debtOutstanding() view returns(uint256)
func (_Yvault034 *Yvault034Session) DebtOutstanding() (*big.Int, error) {
	return _Yvault034.Contract.DebtOutstanding(&_Yvault034.CallOpts)
}

// DebtOutstanding is a free data retrieval call binding the contract method 0xbf3759b5.
//
// Solidity: function debtOutstanding() view returns(uint256)
func (_Yvault034 *Yvault034CallerSession) DebtOutstanding() (*big.Int, error) {
	return _Yvault034.Contract.DebtOutstanding(&_Yvault034.CallOpts)
}

// DebtOutstanding0 is a free data retrieval call binding the contract method 0xbdcf36bb.
//
// Solidity: function debtOutstanding(address strategy) view returns(uint256)
func (_Yvault034 *Yvault034Caller) DebtOutstanding0(opts *bind.CallOpts, strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yvault034.contract.Call(opts, &out, "debtOutstanding0", strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtOutstanding0 is a free data retrieval call binding the contract method 0xbdcf36bb.
//
// Solidity: function debtOutstanding(address strategy) view returns(uint256)
func (_Yvault034 *Yvault034Session) DebtOutstanding0(strategy common.Address) (*big.Int, error) {
	return _Yvault034.Contract.DebtOutstanding0(&_Yvault034.CallOpts, strategy)
}

// DebtOutstanding0 is a free data retrieval call binding the contract method 0xbdcf36bb.
//
// Solidity: function debtOutstanding(address strategy) view returns(uint256)
func (_Yvault034 *Yvault034CallerSession) DebtOutstanding0(strategy common.Address) (*big.Int, error) {
	return _Yvault034.Contract.DebtOutstanding0(&_Yvault034.CallOpts, strategy)
}

// DebtRatio is a free data retrieval call binding the contract method 0xcea55f57.
//
// Solidity: function debtRatio() view returns(uint256)
func (_Yvault034 *Yvault034Caller) DebtRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault034.contract.Call(opts, &out, "debtRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtRatio is a free data retrieval call binding the contract method 0xcea55f57.
//
// Solidity: function debtRatio() view returns(uint256)
func (_Yvault034 *Yvault034Session) DebtRatio() (*big.Int, error) {
	return _Yvault034.Contract.DebtRatio(&_Yvault034.CallOpts)
}

// DebtRatio is a free data retrieval call binding the contract method 0xcea55f57.
//
// Solidity: function debtRatio() view returns(uint256)
func (_Yvault034 *Yvault034CallerSession) DebtRatio() (*big.Int, error) {
	return _Yvault034.Contract.DebtRatio(&_Yvault034.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Yvault034 *Yvault034Caller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault034.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Yvault034 *Yvault034Session) Decimals() (*big.Int, error) {
	return _Yvault034.Contract.Decimals(&_Yvault034.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Yvault034 *Yvault034CallerSession) Decimals() (*big.Int, error) {
	return _Yvault034.Contract.Decimals(&_Yvault034.CallOpts)
}

// DelegatedAssets is a free data retrieval call binding the contract method 0x8e6350e2.
//
// Solidity: function delegatedAssets() view returns(uint256)
func (_Yvault034 *Yvault034Caller) DelegatedAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault034.contract.Call(opts, &out, "delegatedAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegatedAssets is a free data retrieval call binding the contract method 0x8e6350e2.
//
// Solidity: function delegatedAssets() view returns(uint256)
func (_Yvault034 *Yvault034Session) DelegatedAssets() (*big.Int, error) {
	return _Yvault034.Contract.DelegatedAssets(&_Yvault034.CallOpts)
}

// DelegatedAssets is a free data retrieval call binding the contract method 0x8e6350e2.
//
// Solidity: function delegatedAssets() view returns(uint256)
func (_Yvault034 *Yvault034CallerSession) DelegatedAssets() (*big.Int, error) {
	return _Yvault034.Contract.DelegatedAssets(&_Yvault034.CallOpts)
}

// DepositLimit is a free data retrieval call binding the contract method 0xecf70858.
//
// Solidity: function depositLimit() view returns(uint256)
func (_Yvault034 *Yvault034Caller) DepositLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault034.contract.Call(opts, &out, "depositLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositLimit is a free data retrieval call binding the contract method 0xecf70858.
//
// Solidity: function depositLimit() view returns(uint256)
func (_Yvault034 *Yvault034Session) DepositLimit() (*big.Int, error) {
	return _Yvault034.Contract.DepositLimit(&_Yvault034.CallOpts)
}

// DepositLimit is a free data retrieval call binding the contract method 0xecf70858.
//
// Solidity: function depositLimit() view returns(uint256)
func (_Yvault034 *Yvault034CallerSession) DepositLimit() (*big.Int, error) {
	return _Yvault034.Contract.DepositLimit(&_Yvault034.CallOpts)
}

// EmergencyShutdown is a free data retrieval call binding the contract method 0x3403c2fc.
//
// Solidity: function emergencyShutdown() view returns(bool)
func (_Yvault034 *Yvault034Caller) EmergencyShutdown(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Yvault034.contract.Call(opts, &out, "emergencyShutdown")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EmergencyShutdown is a free data retrieval call binding the contract method 0x3403c2fc.
//
// Solidity: function emergencyShutdown() view returns(bool)
func (_Yvault034 *Yvault034Session) EmergencyShutdown() (bool, error) {
	return _Yvault034.Contract.EmergencyShutdown(&_Yvault034.CallOpts)
}

// EmergencyShutdown is a free data retrieval call binding the contract method 0x3403c2fc.
//
// Solidity: function emergencyShutdown() view returns(bool)
func (_Yvault034 *Yvault034CallerSession) EmergencyShutdown() (bool, error) {
	return _Yvault034.Contract.EmergencyShutdown(&_Yvault034.CallOpts)
}

// ExpectedReturn is a free data retrieval call binding the contract method 0xd3406abd.
//
// Solidity: function expectedReturn() view returns(uint256)
func (_Yvault034 *Yvault034Caller) ExpectedReturn(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault034.contract.Call(opts, &out, "expectedReturn")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpectedReturn is a free data retrieval call binding the contract method 0xd3406abd.
//
// Solidity: function expectedReturn() view returns(uint256)
func (_Yvault034 *Yvault034Session) ExpectedReturn() (*big.Int, error) {
	return _Yvault034.Contract.ExpectedReturn(&_Yvault034.CallOpts)
}

// ExpectedReturn is a free data retrieval call binding the contract method 0xd3406abd.
//
// Solidity: function expectedReturn() view returns(uint256)
func (_Yvault034 *Yvault034CallerSession) ExpectedReturn() (*big.Int, error) {
	return _Yvault034.Contract.ExpectedReturn(&_Yvault034.CallOpts)
}

// ExpectedReturn0 is a free data retrieval call binding the contract method 0x33586b67.
//
// Solidity: function expectedReturn(address strategy) view returns(uint256)
func (_Yvault034 *Yvault034Caller) ExpectedReturn0(opts *bind.CallOpts, strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yvault034.contract.Call(opts, &out, "expectedReturn0", strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpectedReturn0 is a free data retrieval call binding the contract method 0x33586b67.
//
// Solidity: function expectedReturn(address strategy) view returns(uint256)
func (_Yvault034 *Yvault034Session) ExpectedReturn0(strategy common.Address) (*big.Int, error) {
	return _Yvault034.Contract.ExpectedReturn0(&_Yvault034.CallOpts, strategy)
}

// ExpectedReturn0 is a free data retrieval call binding the contract method 0x33586b67.
//
// Solidity: function expectedReturn(address strategy) view returns(uint256)
func (_Yvault034 *Yvault034CallerSession) ExpectedReturn0(strategy common.Address) (*big.Int, error) {
	return _Yvault034.Contract.ExpectedReturn0(&_Yvault034.CallOpts, strategy)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Yvault034 *Yvault034Caller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yvault034.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Yvault034 *Yvault034Session) Governance() (common.Address, error) {
	return _Yvault034.Contract.Governance(&_Yvault034.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Yvault034 *Yvault034CallerSession) Governance() (common.Address, error) {
	return _Yvault034.Contract.Governance(&_Yvault034.CallOpts)
}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_Yvault034 *Yvault034Caller) Guardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yvault034.contract.Call(opts, &out, "guardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_Yvault034 *Yvault034Session) Guardian() (common.Address, error) {
	return _Yvault034.Contract.Guardian(&_Yvault034.CallOpts)
}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_Yvault034 *Yvault034CallerSession) Guardian() (common.Address, error) {
	return _Yvault034.Contract.Guardian(&_Yvault034.CallOpts)
}

// GuestList is a free data retrieval call binding the contract method 0x46d55875.
//
// Solidity: function guestList() view returns(address)
func (_Yvault034 *Yvault034Caller) GuestList(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yvault034.contract.Call(opts, &out, "guestList")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GuestList is a free data retrieval call binding the contract method 0x46d55875.
//
// Solidity: function guestList() view returns(address)
func (_Yvault034 *Yvault034Session) GuestList() (common.Address, error) {
	return _Yvault034.Contract.GuestList(&_Yvault034.CallOpts)
}

// GuestList is a free data retrieval call binding the contract method 0x46d55875.
//
// Solidity: function guestList() view returns(address)
func (_Yvault034 *Yvault034CallerSession) GuestList() (common.Address, error) {
	return _Yvault034.Contract.GuestList(&_Yvault034.CallOpts)
}

// LastReport is a free data retrieval call binding the contract method 0xc3535b52.
//
// Solidity: function lastReport() view returns(uint256)
func (_Yvault034 *Yvault034Caller) LastReport(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault034.contract.Call(opts, &out, "lastReport")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastReport is a free data retrieval call binding the contract method 0xc3535b52.
//
// Solidity: function lastReport() view returns(uint256)
func (_Yvault034 *Yvault034Session) LastReport() (*big.Int, error) {
	return _Yvault034.Contract.LastReport(&_Yvault034.CallOpts)
}

// LastReport is a free data retrieval call binding the contract method 0xc3535b52.
//
// Solidity: function lastReport() view returns(uint256)
func (_Yvault034 *Yvault034CallerSession) LastReport() (*big.Int, error) {
	return _Yvault034.Contract.LastReport(&_Yvault034.CallOpts)
}

// LockedProfit is a free data retrieval call binding the contract method 0x44b81396.
//
// Solidity: function lockedProfit() view returns(uint256)
func (_Yvault034 *Yvault034Caller) LockedProfit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault034.contract.Call(opts, &out, "lockedProfit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedProfit is a free data retrieval call binding the contract method 0x44b81396.
//
// Solidity: function lockedProfit() view returns(uint256)
func (_Yvault034 *Yvault034Session) LockedProfit() (*big.Int, error) {
	return _Yvault034.Contract.LockedProfit(&_Yvault034.CallOpts)
}

// LockedProfit is a free data retrieval call binding the contract method 0x44b81396.
//
// Solidity: function lockedProfit() view returns(uint256)
func (_Yvault034 *Yvault034CallerSession) LockedProfit() (*big.Int, error) {
	return _Yvault034.Contract.LockedProfit(&_Yvault034.CallOpts)
}

// LockedProfitDegration is a free data retrieval call binding the contract method 0x2140254d.
//
// Solidity: function lockedProfitDegration() view returns(uint256)
func (_Yvault034 *Yvault034Caller) LockedProfitDegration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault034.contract.Call(opts, &out, "lockedProfitDegration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedProfitDegration is a free data retrieval call binding the contract method 0x2140254d.
//
// Solidity: function lockedProfitDegration() view returns(uint256)
func (_Yvault034 *Yvault034Session) LockedProfitDegration() (*big.Int, error) {
	return _Yvault034.Contract.LockedProfitDegration(&_Yvault034.CallOpts)
}

// LockedProfitDegration is a free data retrieval call binding the contract method 0x2140254d.
//
// Solidity: function lockedProfitDegration() view returns(uint256)
func (_Yvault034 *Yvault034CallerSession) LockedProfitDegration() (*big.Int, error) {
	return _Yvault034.Contract.LockedProfitDegration(&_Yvault034.CallOpts)
}

// Management is a free data retrieval call binding the contract method 0x88a8d602.
//
// Solidity: function management() view returns(address)
func (_Yvault034 *Yvault034Caller) Management(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yvault034.contract.Call(opts, &out, "management")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Management is a free data retrieval call binding the contract method 0x88a8d602.
//
// Solidity: function management() view returns(address)
func (_Yvault034 *Yvault034Session) Management() (common.Address, error) {
	return _Yvault034.Contract.Management(&_Yvault034.CallOpts)
}

// Management is a free data retrieval call binding the contract method 0x88a8d602.
//
// Solidity: function management() view returns(address)
func (_Yvault034 *Yvault034CallerSession) Management() (common.Address, error) {
	return _Yvault034.Contract.Management(&_Yvault034.CallOpts)
}

// ManagementFee is a free data retrieval call binding the contract method 0xa6f7f5d6.
//
// Solidity: function managementFee() view returns(uint256)
func (_Yvault034 *Yvault034Caller) ManagementFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault034.contract.Call(opts, &out, "managementFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ManagementFee is a free data retrieval call binding the contract method 0xa6f7f5d6.
//
// Solidity: function managementFee() view returns(uint256)
func (_Yvault034 *Yvault034Session) ManagementFee() (*big.Int, error) {
	return _Yvault034.Contract.ManagementFee(&_Yvault034.CallOpts)
}

// ManagementFee is a free data retrieval call binding the contract method 0xa6f7f5d6.
//
// Solidity: function managementFee() view returns(uint256)
func (_Yvault034 *Yvault034CallerSession) ManagementFee() (*big.Int, error) {
	return _Yvault034.Contract.ManagementFee(&_Yvault034.CallOpts)
}

// MaxAvailableShares is a free data retrieval call binding the contract method 0x75de2902.
//
// Solidity: function maxAvailableShares() view returns(uint256)
func (_Yvault034 *Yvault034Caller) MaxAvailableShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault034.contract.Call(opts, &out, "maxAvailableShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAvailableShares is a free data retrieval call binding the contract method 0x75de2902.
//
// Solidity: function maxAvailableShares() view returns(uint256)
func (_Yvault034 *Yvault034Session) MaxAvailableShares() (*big.Int, error) {
	return _Yvault034.Contract.MaxAvailableShares(&_Yvault034.CallOpts)
}

// MaxAvailableShares is a free data retrieval call binding the contract method 0x75de2902.
//
// Solidity: function maxAvailableShares() view returns(uint256)
func (_Yvault034 *Yvault034CallerSession) MaxAvailableShares() (*big.Int, error) {
	return _Yvault034.Contract.MaxAvailableShares(&_Yvault034.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Yvault034 *Yvault034Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Yvault034.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Yvault034 *Yvault034Session) Name() (string, error) {
	return _Yvault034.Contract.Name(&_Yvault034.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Yvault034 *Yvault034CallerSession) Name() (string, error) {
	return _Yvault034.Contract.Name(&_Yvault034.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_Yvault034 *Yvault034Caller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yvault034.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_Yvault034 *Yvault034Session) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Yvault034.Contract.Nonces(&_Yvault034.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_Yvault034 *Yvault034CallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Yvault034.Contract.Nonces(&_Yvault034.CallOpts, arg0)
}

// PerformanceFee is a free data retrieval call binding the contract method 0x87788782.
//
// Solidity: function performanceFee() view returns(uint256)
func (_Yvault034 *Yvault034Caller) PerformanceFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault034.contract.Call(opts, &out, "performanceFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PerformanceFee is a free data retrieval call binding the contract method 0x87788782.
//
// Solidity: function performanceFee() view returns(uint256)
func (_Yvault034 *Yvault034Session) PerformanceFee() (*big.Int, error) {
	return _Yvault034.Contract.PerformanceFee(&_Yvault034.CallOpts)
}

// PerformanceFee is a free data retrieval call binding the contract method 0x87788782.
//
// Solidity: function performanceFee() view returns(uint256)
func (_Yvault034 *Yvault034CallerSession) PerformanceFee() (*big.Int, error) {
	return _Yvault034.Contract.PerformanceFee(&_Yvault034.CallOpts)
}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_Yvault034 *Yvault034Caller) PricePerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault034.contract.Call(opts, &out, "pricePerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_Yvault034 *Yvault034Session) PricePerShare() (*big.Int, error) {
	return _Yvault034.Contract.PricePerShare(&_Yvault034.CallOpts)
}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_Yvault034 *Yvault034CallerSession) PricePerShare() (*big.Int, error) {
	return _Yvault034.Contract.PricePerShare(&_Yvault034.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x9ec5a894.
//
// Solidity: function rewards() view returns(address)
func (_Yvault034 *Yvault034Caller) Rewards(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yvault034.contract.Call(opts, &out, "rewards")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0x9ec5a894.
//
// Solidity: function rewards() view returns(address)
func (_Yvault034 *Yvault034Session) Rewards() (common.Address, error) {
	return _Yvault034.Contract.Rewards(&_Yvault034.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x9ec5a894.
//
// Solidity: function rewards() view returns(address)
func (_Yvault034 *Yvault034CallerSession) Rewards() (common.Address, error) {
	return _Yvault034.Contract.Rewards(&_Yvault034.CallOpts)
}

// Strategies is a free data retrieval call binding the contract method 0x39ebf823.
//
// Solidity: function strategies(address arg0) view returns(uint256 performanceFee, uint256 activation, uint256 debtRatio, uint256 minDebtPerHarvest, uint256 maxDebtPerHarvest, uint256 lastReport, uint256 totalDebt, uint256 totalGain, uint256 totalLoss)
func (_Yvault034 *Yvault034Caller) Strategies(opts *bind.CallOpts, arg0 common.Address) (struct {
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
	err := _Yvault034.contract.Call(opts, &out, "strategies", arg0)

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
func (_Yvault034 *Yvault034Session) Strategies(arg0 common.Address) (struct {
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
	return _Yvault034.Contract.Strategies(&_Yvault034.CallOpts, arg0)
}

// Strategies is a free data retrieval call binding the contract method 0x39ebf823.
//
// Solidity: function strategies(address arg0) view returns(uint256 performanceFee, uint256 activation, uint256 debtRatio, uint256 minDebtPerHarvest, uint256 maxDebtPerHarvest, uint256 lastReport, uint256 totalDebt, uint256 totalGain, uint256 totalLoss)
func (_Yvault034 *Yvault034CallerSession) Strategies(arg0 common.Address) (struct {
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
	return _Yvault034.Contract.Strategies(&_Yvault034.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Yvault034 *Yvault034Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Yvault034.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Yvault034 *Yvault034Session) Symbol() (string, error) {
	return _Yvault034.Contract.Symbol(&_Yvault034.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Yvault034 *Yvault034CallerSession) Symbol() (string, error) {
	return _Yvault034.Contract.Symbol(&_Yvault034.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Yvault034 *Yvault034Caller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yvault034.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Yvault034 *Yvault034Session) Token() (common.Address, error) {
	return _Yvault034.Contract.Token(&_Yvault034.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Yvault034 *Yvault034CallerSession) Token() (common.Address, error) {
	return _Yvault034.Contract.Token(&_Yvault034.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_Yvault034 *Yvault034Caller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault034.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_Yvault034 *Yvault034Session) TotalAssets() (*big.Int, error) {
	return _Yvault034.Contract.TotalAssets(&_Yvault034.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_Yvault034 *Yvault034CallerSession) TotalAssets() (*big.Int, error) {
	return _Yvault034.Contract.TotalAssets(&_Yvault034.CallOpts)
}

// TotalDebt is a free data retrieval call binding the contract method 0xfc7b9c18.
//
// Solidity: function totalDebt() view returns(uint256)
func (_Yvault034 *Yvault034Caller) TotalDebt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault034.contract.Call(opts, &out, "totalDebt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDebt is a free data retrieval call binding the contract method 0xfc7b9c18.
//
// Solidity: function totalDebt() view returns(uint256)
func (_Yvault034 *Yvault034Session) TotalDebt() (*big.Int, error) {
	return _Yvault034.Contract.TotalDebt(&_Yvault034.CallOpts)
}

// TotalDebt is a free data retrieval call binding the contract method 0xfc7b9c18.
//
// Solidity: function totalDebt() view returns(uint256)
func (_Yvault034 *Yvault034CallerSession) TotalDebt() (*big.Int, error) {
	return _Yvault034.Contract.TotalDebt(&_Yvault034.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Yvault034 *Yvault034Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault034.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Yvault034 *Yvault034Session) TotalSupply() (*big.Int, error) {
	return _Yvault034.Contract.TotalSupply(&_Yvault034.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Yvault034 *Yvault034CallerSession) TotalSupply() (*big.Int, error) {
	return _Yvault034.Contract.TotalSupply(&_Yvault034.CallOpts)
}

// WithdrawalQueue is a free data retrieval call binding the contract method 0xc822adda.
//
// Solidity: function withdrawalQueue(uint256 arg0) view returns(address)
func (_Yvault034 *Yvault034Caller) WithdrawalQueue(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Yvault034.contract.Call(opts, &out, "withdrawalQueue", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WithdrawalQueue is a free data retrieval call binding the contract method 0xc822adda.
//
// Solidity: function withdrawalQueue(uint256 arg0) view returns(address)
func (_Yvault034 *Yvault034Session) WithdrawalQueue(arg0 *big.Int) (common.Address, error) {
	return _Yvault034.Contract.WithdrawalQueue(&_Yvault034.CallOpts, arg0)
}

// WithdrawalQueue is a free data retrieval call binding the contract method 0xc822adda.
//
// Solidity: function withdrawalQueue(uint256 arg0) view returns(address)
func (_Yvault034 *Yvault034CallerSession) WithdrawalQueue(arg0 *big.Int) (common.Address, error) {
	return _Yvault034.Contract.WithdrawalQueue(&_Yvault034.CallOpts, arg0)
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_Yvault034 *Yvault034Transactor) AcceptGovernance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yvault034.contract.Transact(opts, "acceptGovernance")
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_Yvault034 *Yvault034Session) AcceptGovernance() (*types.Transaction, error) {
	return _Yvault034.Contract.AcceptGovernance(&_Yvault034.TransactOpts)
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_Yvault034 *Yvault034TransactorSession) AcceptGovernance() (*types.Transaction, error) {
	return _Yvault034.Contract.AcceptGovernance(&_Yvault034.TransactOpts)
}

// AddStrategy is a paid mutator transaction binding the contract method 0x14b4e26e.
//
// Solidity: function addStrategy(address strategy, uint256 debtRatio, uint256 minDebtPerHarvest, uint256 maxDebtPerHarvest, uint256 performanceFee) returns()
func (_Yvault034 *Yvault034Transactor) AddStrategy(opts *bind.TransactOpts, strategy common.Address, debtRatio *big.Int, minDebtPerHarvest *big.Int, maxDebtPerHarvest *big.Int, performanceFee *big.Int) (*types.Transaction, error) {
	return _Yvault034.contract.Transact(opts, "addStrategy", strategy, debtRatio, minDebtPerHarvest, maxDebtPerHarvest, performanceFee)
}

// AddStrategy is a paid mutator transaction binding the contract method 0x14b4e26e.
//
// Solidity: function addStrategy(address strategy, uint256 debtRatio, uint256 minDebtPerHarvest, uint256 maxDebtPerHarvest, uint256 performanceFee) returns()
func (_Yvault034 *Yvault034Session) AddStrategy(strategy common.Address, debtRatio *big.Int, minDebtPerHarvest *big.Int, maxDebtPerHarvest *big.Int, performanceFee *big.Int) (*types.Transaction, error) {
	return _Yvault034.Contract.AddStrategy(&_Yvault034.TransactOpts, strategy, debtRatio, minDebtPerHarvest, maxDebtPerHarvest, performanceFee)
}

// AddStrategy is a paid mutator transaction binding the contract method 0x14b4e26e.
//
// Solidity: function addStrategy(address strategy, uint256 debtRatio, uint256 minDebtPerHarvest, uint256 maxDebtPerHarvest, uint256 performanceFee) returns()
func (_Yvault034 *Yvault034TransactorSession) AddStrategy(strategy common.Address, debtRatio *big.Int, minDebtPerHarvest *big.Int, maxDebtPerHarvest *big.Int, performanceFee *big.Int) (*types.Transaction, error) {
	return _Yvault034.Contract.AddStrategy(&_Yvault034.TransactOpts, strategy, debtRatio, minDebtPerHarvest, maxDebtPerHarvest, performanceFee)
}

// AddStrategyToQueue is a paid mutator transaction binding the contract method 0xf76e4caa.
//
// Solidity: function addStrategyToQueue(address strategy) returns()
func (_Yvault034 *Yvault034Transactor) AddStrategyToQueue(opts *bind.TransactOpts, strategy common.Address) (*types.Transaction, error) {
	return _Yvault034.contract.Transact(opts, "addStrategyToQueue", strategy)
}

// AddStrategyToQueue is a paid mutator transaction binding the contract method 0xf76e4caa.
//
// Solidity: function addStrategyToQueue(address strategy) returns()
func (_Yvault034 *Yvault034Session) AddStrategyToQueue(strategy common.Address) (*types.Transaction, error) {
	return _Yvault034.Contract.AddStrategyToQueue(&_Yvault034.TransactOpts, strategy)
}

// AddStrategyToQueue is a paid mutator transaction binding the contract method 0xf76e4caa.
//
// Solidity: function addStrategyToQueue(address strategy) returns()
func (_Yvault034 *Yvault034TransactorSession) AddStrategyToQueue(strategy common.Address) (*types.Transaction, error) {
	return _Yvault034.Contract.AddStrategyToQueue(&_Yvault034.TransactOpts, strategy)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Yvault034 *Yvault034Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault034.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Yvault034 *Yvault034Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault034.Contract.Approve(&_Yvault034.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Yvault034 *Yvault034TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault034.Contract.Approve(&_Yvault034.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 amount) returns(bool)
func (_Yvault034 *Yvault034Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault034.contract.Transact(opts, "decreaseAllowance", spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 amount) returns(bool)
func (_Yvault034 *Yvault034Session) DecreaseAllowance(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault034.Contract.DecreaseAllowance(&_Yvault034.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 amount) returns(bool)
func (_Yvault034 *Yvault034TransactorSession) DecreaseAllowance(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault034.Contract.DecreaseAllowance(&_Yvault034.TransactOpts, spender, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(uint256)
func (_Yvault034 *Yvault034Transactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yvault034.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(uint256)
func (_Yvault034 *Yvault034Session) Deposit() (*types.Transaction, error) {
	return _Yvault034.Contract.Deposit(&_Yvault034.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(uint256)
func (_Yvault034 *Yvault034TransactorSession) Deposit() (*types.Transaction, error) {
	return _Yvault034.Contract.Deposit(&_Yvault034.TransactOpts)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns(uint256)
func (_Yvault034 *Yvault034Transactor) Deposit0(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Yvault034.contract.Transact(opts, "deposit0", _amount)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns(uint256)
func (_Yvault034 *Yvault034Session) Deposit0(_amount *big.Int) (*types.Transaction, error) {
	return _Yvault034.Contract.Deposit0(&_Yvault034.TransactOpts, _amount)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns(uint256)
func (_Yvault034 *Yvault034TransactorSession) Deposit0(_amount *big.Int) (*types.Transaction, error) {
	return _Yvault034.Contract.Deposit0(&_Yvault034.TransactOpts, _amount)
}

// Deposit1 is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _amount, address recipient) returns(uint256)
func (_Yvault034 *Yvault034Transactor) Deposit1(opts *bind.TransactOpts, _amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Yvault034.contract.Transact(opts, "deposit1", _amount, recipient)
}

// Deposit1 is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _amount, address recipient) returns(uint256)
func (_Yvault034 *Yvault034Session) Deposit1(_amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Yvault034.Contract.Deposit1(&_Yvault034.TransactOpts, _amount, recipient)
}

// Deposit1 is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _amount, address recipient) returns(uint256)
func (_Yvault034 *Yvault034TransactorSession) Deposit1(_amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Yvault034.Contract.Deposit1(&_Yvault034.TransactOpts, _amount, recipient)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 amount) returns(bool)
func (_Yvault034 *Yvault034Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault034.contract.Transact(opts, "increaseAllowance", spender, amount)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 amount) returns(bool)
func (_Yvault034 *Yvault034Session) IncreaseAllowance(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault034.Contract.IncreaseAllowance(&_Yvault034.TransactOpts, spender, amount)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 amount) returns(bool)
func (_Yvault034 *Yvault034TransactorSession) IncreaseAllowance(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault034.Contract.IncreaseAllowance(&_Yvault034.TransactOpts, spender, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x83b43589.
//
// Solidity: function initialize(address token, address governance, address rewards, string nameOverride, string symbolOverride) returns()
func (_Yvault034 *Yvault034Transactor) Initialize(opts *bind.TransactOpts, token common.Address, governance common.Address, rewards common.Address, nameOverride string, symbolOverride string) (*types.Transaction, error) {
	return _Yvault034.contract.Transact(opts, "initialize", token, governance, rewards, nameOverride, symbolOverride)
}

// Initialize is a paid mutator transaction binding the contract method 0x83b43589.
//
// Solidity: function initialize(address token, address governance, address rewards, string nameOverride, string symbolOverride) returns()
func (_Yvault034 *Yvault034Session) Initialize(token common.Address, governance common.Address, rewards common.Address, nameOverride string, symbolOverride string) (*types.Transaction, error) {
	return _Yvault034.Contract.Initialize(&_Yvault034.TransactOpts, token, governance, rewards, nameOverride, symbolOverride)
}

// Initialize is a paid mutator transaction binding the contract method 0x83b43589.
//
// Solidity: function initialize(address token, address governance, address rewards, string nameOverride, string symbolOverride) returns()
func (_Yvault034 *Yvault034TransactorSession) Initialize(token common.Address, governance common.Address, rewards common.Address, nameOverride string, symbolOverride string) (*types.Transaction, error) {
	return _Yvault034.Contract.Initialize(&_Yvault034.TransactOpts, token, governance, rewards, nameOverride, symbolOverride)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xa5b81fdf.
//
// Solidity: function initialize(address token, address governance, address rewards, string nameOverride, string symbolOverride, address guardian) returns()
func (_Yvault034 *Yvault034Transactor) Initialize0(opts *bind.TransactOpts, token common.Address, governance common.Address, rewards common.Address, nameOverride string, symbolOverride string, guardian common.Address) (*types.Transaction, error) {
	return _Yvault034.contract.Transact(opts, "initialize0", token, governance, rewards, nameOverride, symbolOverride, guardian)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xa5b81fdf.
//
// Solidity: function initialize(address token, address governance, address rewards, string nameOverride, string symbolOverride, address guardian) returns()
func (_Yvault034 *Yvault034Session) Initialize0(token common.Address, governance common.Address, rewards common.Address, nameOverride string, symbolOverride string, guardian common.Address) (*types.Transaction, error) {
	return _Yvault034.Contract.Initialize0(&_Yvault034.TransactOpts, token, governance, rewards, nameOverride, symbolOverride, guardian)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xa5b81fdf.
//
// Solidity: function initialize(address token, address governance, address rewards, string nameOverride, string symbolOverride, address guardian) returns()
func (_Yvault034 *Yvault034TransactorSession) Initialize0(token common.Address, governance common.Address, rewards common.Address, nameOverride string, symbolOverride string, guardian common.Address) (*types.Transaction, error) {
	return _Yvault034.Contract.Initialize0(&_Yvault034.TransactOpts, token, governance, rewards, nameOverride, symbolOverride, guardian)
}

// MigrateStrategy is a paid mutator transaction binding the contract method 0x6cb56d19.
//
// Solidity: function migrateStrategy(address oldVersion, address newVersion) returns()
func (_Yvault034 *Yvault034Transactor) MigrateStrategy(opts *bind.TransactOpts, oldVersion common.Address, newVersion common.Address) (*types.Transaction, error) {
	return _Yvault034.contract.Transact(opts, "migrateStrategy", oldVersion, newVersion)
}

// MigrateStrategy is a paid mutator transaction binding the contract method 0x6cb56d19.
//
// Solidity: function migrateStrategy(address oldVersion, address newVersion) returns()
func (_Yvault034 *Yvault034Session) MigrateStrategy(oldVersion common.Address, newVersion common.Address) (*types.Transaction, error) {
	return _Yvault034.Contract.MigrateStrategy(&_Yvault034.TransactOpts, oldVersion, newVersion)
}

// MigrateStrategy is a paid mutator transaction binding the contract method 0x6cb56d19.
//
// Solidity: function migrateStrategy(address oldVersion, address newVersion) returns()
func (_Yvault034 *Yvault034TransactorSession) MigrateStrategy(oldVersion common.Address, newVersion common.Address) (*types.Transaction, error) {
	return _Yvault034.Contract.MigrateStrategy(&_Yvault034.TransactOpts, oldVersion, newVersion)
}

// Permit is a paid mutator transaction binding the contract method 0x9fd5a6cf.
//
// Solidity: function permit(address owner, address spender, uint256 amount, uint256 expiry, bytes signature) returns(bool)
func (_Yvault034 *Yvault034Transactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, amount *big.Int, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _Yvault034.contract.Transact(opts, "permit", owner, spender, amount, expiry, signature)
}

// Permit is a paid mutator transaction binding the contract method 0x9fd5a6cf.
//
// Solidity: function permit(address owner, address spender, uint256 amount, uint256 expiry, bytes signature) returns(bool)
func (_Yvault034 *Yvault034Session) Permit(owner common.Address, spender common.Address, amount *big.Int, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _Yvault034.Contract.Permit(&_Yvault034.TransactOpts, owner, spender, amount, expiry, signature)
}

// Permit is a paid mutator transaction binding the contract method 0x9fd5a6cf.
//
// Solidity: function permit(address owner, address spender, uint256 amount, uint256 expiry, bytes signature) returns(bool)
func (_Yvault034 *Yvault034TransactorSession) Permit(owner common.Address, spender common.Address, amount *big.Int, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _Yvault034.Contract.Permit(&_Yvault034.TransactOpts, owner, spender, amount, expiry, signature)
}

// RemoveStrategyFromQueue is a paid mutator transaction binding the contract method 0xb22439f5.
//
// Solidity: function removeStrategyFromQueue(address strategy) returns()
func (_Yvault034 *Yvault034Transactor) RemoveStrategyFromQueue(opts *bind.TransactOpts, strategy common.Address) (*types.Transaction, error) {
	return _Yvault034.contract.Transact(opts, "removeStrategyFromQueue", strategy)
}

// RemoveStrategyFromQueue is a paid mutator transaction binding the contract method 0xb22439f5.
//
// Solidity: function removeStrategyFromQueue(address strategy) returns()
func (_Yvault034 *Yvault034Session) RemoveStrategyFromQueue(strategy common.Address) (*types.Transaction, error) {
	return _Yvault034.Contract.RemoveStrategyFromQueue(&_Yvault034.TransactOpts, strategy)
}

// RemoveStrategyFromQueue is a paid mutator transaction binding the contract method 0xb22439f5.
//
// Solidity: function removeStrategyFromQueue(address strategy) returns()
func (_Yvault034 *Yvault034TransactorSession) RemoveStrategyFromQueue(strategy common.Address) (*types.Transaction, error) {
	return _Yvault034.Contract.RemoveStrategyFromQueue(&_Yvault034.TransactOpts, strategy)
}

// Report is a paid mutator transaction binding the contract method 0xa1d9bafc.
//
// Solidity: function report(uint256 gain, uint256 loss, uint256 _debtPayment) returns(uint256)
func (_Yvault034 *Yvault034Transactor) Report(opts *bind.TransactOpts, gain *big.Int, loss *big.Int, _debtPayment *big.Int) (*types.Transaction, error) {
	return _Yvault034.contract.Transact(opts, "report", gain, loss, _debtPayment)
}

// Report is a paid mutator transaction binding the contract method 0xa1d9bafc.
//
// Solidity: function report(uint256 gain, uint256 loss, uint256 _debtPayment) returns(uint256)
func (_Yvault034 *Yvault034Session) Report(gain *big.Int, loss *big.Int, _debtPayment *big.Int) (*types.Transaction, error) {
	return _Yvault034.Contract.Report(&_Yvault034.TransactOpts, gain, loss, _debtPayment)
}

// Report is a paid mutator transaction binding the contract method 0xa1d9bafc.
//
// Solidity: function report(uint256 gain, uint256 loss, uint256 _debtPayment) returns(uint256)
func (_Yvault034 *Yvault034TransactorSession) Report(gain *big.Int, loss *big.Int, _debtPayment *big.Int) (*types.Transaction, error) {
	return _Yvault034.Contract.Report(&_Yvault034.TransactOpts, gain, loss, _debtPayment)
}

// RevokeStrategy is a paid mutator transaction binding the contract method 0xa0e4af9a.
//
// Solidity: function revokeStrategy() returns()
func (_Yvault034 *Yvault034Transactor) RevokeStrategy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yvault034.contract.Transact(opts, "revokeStrategy")
}

// RevokeStrategy is a paid mutator transaction binding the contract method 0xa0e4af9a.
//
// Solidity: function revokeStrategy() returns()
func (_Yvault034 *Yvault034Session) RevokeStrategy() (*types.Transaction, error) {
	return _Yvault034.Contract.RevokeStrategy(&_Yvault034.TransactOpts)
}

// RevokeStrategy is a paid mutator transaction binding the contract method 0xa0e4af9a.
//
// Solidity: function revokeStrategy() returns()
func (_Yvault034 *Yvault034TransactorSession) RevokeStrategy() (*types.Transaction, error) {
	return _Yvault034.Contract.RevokeStrategy(&_Yvault034.TransactOpts)
}

// RevokeStrategy0 is a paid mutator transaction binding the contract method 0xbb994d48.
//
// Solidity: function revokeStrategy(address strategy) returns()
func (_Yvault034 *Yvault034Transactor) RevokeStrategy0(opts *bind.TransactOpts, strategy common.Address) (*types.Transaction, error) {
	return _Yvault034.contract.Transact(opts, "revokeStrategy0", strategy)
}

// RevokeStrategy0 is a paid mutator transaction binding the contract method 0xbb994d48.
//
// Solidity: function revokeStrategy(address strategy) returns()
func (_Yvault034 *Yvault034Session) RevokeStrategy0(strategy common.Address) (*types.Transaction, error) {
	return _Yvault034.Contract.RevokeStrategy0(&_Yvault034.TransactOpts, strategy)
}

// RevokeStrategy0 is a paid mutator transaction binding the contract method 0xbb994d48.
//
// Solidity: function revokeStrategy(address strategy) returns()
func (_Yvault034 *Yvault034TransactorSession) RevokeStrategy0(strategy common.Address) (*types.Transaction, error) {
	return _Yvault034.Contract.RevokeStrategy0(&_Yvault034.TransactOpts, strategy)
}

// SetDepositLimit is a paid mutator transaction binding the contract method 0xbdc8144b.
//
// Solidity: function setDepositLimit(uint256 limit) returns()
func (_Yvault034 *Yvault034Transactor) SetDepositLimit(opts *bind.TransactOpts, limit *big.Int) (*types.Transaction, error) {
	return _Yvault034.contract.Transact(opts, "setDepositLimit", limit)
}

// SetDepositLimit is a paid mutator transaction binding the contract method 0xbdc8144b.
//
// Solidity: function setDepositLimit(uint256 limit) returns()
func (_Yvault034 *Yvault034Session) SetDepositLimit(limit *big.Int) (*types.Transaction, error) {
	return _Yvault034.Contract.SetDepositLimit(&_Yvault034.TransactOpts, limit)
}

// SetDepositLimit is a paid mutator transaction binding the contract method 0xbdc8144b.
//
// Solidity: function setDepositLimit(uint256 limit) returns()
func (_Yvault034 *Yvault034TransactorSession) SetDepositLimit(limit *big.Int) (*types.Transaction, error) {
	return _Yvault034.Contract.SetDepositLimit(&_Yvault034.TransactOpts, limit)
}

// SetEmergencyShutdown is a paid mutator transaction binding the contract method 0x14c64402.
//
// Solidity: function setEmergencyShutdown(bool active) returns()
func (_Yvault034 *Yvault034Transactor) SetEmergencyShutdown(opts *bind.TransactOpts, active bool) (*types.Transaction, error) {
	return _Yvault034.contract.Transact(opts, "setEmergencyShutdown", active)
}

// SetEmergencyShutdown is a paid mutator transaction binding the contract method 0x14c64402.
//
// Solidity: function setEmergencyShutdown(bool active) returns()
func (_Yvault034 *Yvault034Session) SetEmergencyShutdown(active bool) (*types.Transaction, error) {
	return _Yvault034.Contract.SetEmergencyShutdown(&_Yvault034.TransactOpts, active)
}

// SetEmergencyShutdown is a paid mutator transaction binding the contract method 0x14c64402.
//
// Solidity: function setEmergencyShutdown(bool active) returns()
func (_Yvault034 *Yvault034TransactorSession) SetEmergencyShutdown(active bool) (*types.Transaction, error) {
	return _Yvault034.Contract.SetEmergencyShutdown(&_Yvault034.TransactOpts, active)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address governance) returns()
func (_Yvault034 *Yvault034Transactor) SetGovernance(opts *bind.TransactOpts, governance common.Address) (*types.Transaction, error) {
	return _Yvault034.contract.Transact(opts, "setGovernance", governance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address governance) returns()
func (_Yvault034 *Yvault034Session) SetGovernance(governance common.Address) (*types.Transaction, error) {
	return _Yvault034.Contract.SetGovernance(&_Yvault034.TransactOpts, governance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address governance) returns()
func (_Yvault034 *Yvault034TransactorSession) SetGovernance(governance common.Address) (*types.Transaction, error) {
	return _Yvault034.Contract.SetGovernance(&_Yvault034.TransactOpts, governance)
}

// SetGuardian is a paid mutator transaction binding the contract method 0x8a0dac4a.
//
// Solidity: function setGuardian(address guardian) returns()
func (_Yvault034 *Yvault034Transactor) SetGuardian(opts *bind.TransactOpts, guardian common.Address) (*types.Transaction, error) {
	return _Yvault034.contract.Transact(opts, "setGuardian", guardian)
}

// SetGuardian is a paid mutator transaction binding the contract method 0x8a0dac4a.
//
// Solidity: function setGuardian(address guardian) returns()
func (_Yvault034 *Yvault034Session) SetGuardian(guardian common.Address) (*types.Transaction, error) {
	return _Yvault034.Contract.SetGuardian(&_Yvault034.TransactOpts, guardian)
}

// SetGuardian is a paid mutator transaction binding the contract method 0x8a0dac4a.
//
// Solidity: function setGuardian(address guardian) returns()
func (_Yvault034 *Yvault034TransactorSession) SetGuardian(guardian common.Address) (*types.Transaction, error) {
	return _Yvault034.Contract.SetGuardian(&_Yvault034.TransactOpts, guardian)
}

// SetGuestList is a paid mutator transaction binding the contract method 0x0b5b78eb.
//
// Solidity: function setGuestList(address guestList) returns()
func (_Yvault034 *Yvault034Transactor) SetGuestList(opts *bind.TransactOpts, guestList common.Address) (*types.Transaction, error) {
	return _Yvault034.contract.Transact(opts, "setGuestList", guestList)
}

// SetGuestList is a paid mutator transaction binding the contract method 0x0b5b78eb.
//
// Solidity: function setGuestList(address guestList) returns()
func (_Yvault034 *Yvault034Session) SetGuestList(guestList common.Address) (*types.Transaction, error) {
	return _Yvault034.Contract.SetGuestList(&_Yvault034.TransactOpts, guestList)
}

// SetGuestList is a paid mutator transaction binding the contract method 0x0b5b78eb.
//
// Solidity: function setGuestList(address guestList) returns()
func (_Yvault034 *Yvault034TransactorSession) SetGuestList(guestList common.Address) (*types.Transaction, error) {
	return _Yvault034.Contract.SetGuestList(&_Yvault034.TransactOpts, guestList)
}

// SetLockedProfitDegration is a paid mutator transaction binding the contract method 0x8402a84f.
//
// Solidity: function setLockedProfitDegration(uint256 degration) returns()
func (_Yvault034 *Yvault034Transactor) SetLockedProfitDegration(opts *bind.TransactOpts, degration *big.Int) (*types.Transaction, error) {
	return _Yvault034.contract.Transact(opts, "setLockedProfitDegration", degration)
}

// SetLockedProfitDegration is a paid mutator transaction binding the contract method 0x8402a84f.
//
// Solidity: function setLockedProfitDegration(uint256 degration) returns()
func (_Yvault034 *Yvault034Session) SetLockedProfitDegration(degration *big.Int) (*types.Transaction, error) {
	return _Yvault034.Contract.SetLockedProfitDegration(&_Yvault034.TransactOpts, degration)
}

// SetLockedProfitDegration is a paid mutator transaction binding the contract method 0x8402a84f.
//
// Solidity: function setLockedProfitDegration(uint256 degration) returns()
func (_Yvault034 *Yvault034TransactorSession) SetLockedProfitDegration(degration *big.Int) (*types.Transaction, error) {
	return _Yvault034.Contract.SetLockedProfitDegration(&_Yvault034.TransactOpts, degration)
}

// SetManagement is a paid mutator transaction binding the contract method 0xd4a22bde.
//
// Solidity: function setManagement(address management) returns()
func (_Yvault034 *Yvault034Transactor) SetManagement(opts *bind.TransactOpts, management common.Address) (*types.Transaction, error) {
	return _Yvault034.contract.Transact(opts, "setManagement", management)
}

// SetManagement is a paid mutator transaction binding the contract method 0xd4a22bde.
//
// Solidity: function setManagement(address management) returns()
func (_Yvault034 *Yvault034Session) SetManagement(management common.Address) (*types.Transaction, error) {
	return _Yvault034.Contract.SetManagement(&_Yvault034.TransactOpts, management)
}

// SetManagement is a paid mutator transaction binding the contract method 0xd4a22bde.
//
// Solidity: function setManagement(address management) returns()
func (_Yvault034 *Yvault034TransactorSession) SetManagement(management common.Address) (*types.Transaction, error) {
	return _Yvault034.Contract.SetManagement(&_Yvault034.TransactOpts, management)
}

// SetManagementFee is a paid mutator transaction binding the contract method 0xfe56e232.
//
// Solidity: function setManagementFee(uint256 fee) returns()
func (_Yvault034 *Yvault034Transactor) SetManagementFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _Yvault034.contract.Transact(opts, "setManagementFee", fee)
}

// SetManagementFee is a paid mutator transaction binding the contract method 0xfe56e232.
//
// Solidity: function setManagementFee(uint256 fee) returns()
func (_Yvault034 *Yvault034Session) SetManagementFee(fee *big.Int) (*types.Transaction, error) {
	return _Yvault034.Contract.SetManagementFee(&_Yvault034.TransactOpts, fee)
}

// SetManagementFee is a paid mutator transaction binding the contract method 0xfe56e232.
//
// Solidity: function setManagementFee(uint256 fee) returns()
func (_Yvault034 *Yvault034TransactorSession) SetManagementFee(fee *big.Int) (*types.Transaction, error) {
	return _Yvault034.Contract.SetManagementFee(&_Yvault034.TransactOpts, fee)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name) returns()
func (_Yvault034 *Yvault034Transactor) SetName(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _Yvault034.contract.Transact(opts, "setName", name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name) returns()
func (_Yvault034 *Yvault034Session) SetName(name string) (*types.Transaction, error) {
	return _Yvault034.Contract.SetName(&_Yvault034.TransactOpts, name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name) returns()
func (_Yvault034 *Yvault034TransactorSession) SetName(name string) (*types.Transaction, error) {
	return _Yvault034.Contract.SetName(&_Yvault034.TransactOpts, name)
}

// SetPerformanceFee is a paid mutator transaction binding the contract method 0x70897b23.
//
// Solidity: function setPerformanceFee(uint256 fee) returns()
func (_Yvault034 *Yvault034Transactor) SetPerformanceFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _Yvault034.contract.Transact(opts, "setPerformanceFee", fee)
}

// SetPerformanceFee is a paid mutator transaction binding the contract method 0x70897b23.
//
// Solidity: function setPerformanceFee(uint256 fee) returns()
func (_Yvault034 *Yvault034Session) SetPerformanceFee(fee *big.Int) (*types.Transaction, error) {
	return _Yvault034.Contract.SetPerformanceFee(&_Yvault034.TransactOpts, fee)
}

// SetPerformanceFee is a paid mutator transaction binding the contract method 0x70897b23.
//
// Solidity: function setPerformanceFee(uint256 fee) returns()
func (_Yvault034 *Yvault034TransactorSession) SetPerformanceFee(fee *big.Int) (*types.Transaction, error) {
	return _Yvault034.Contract.SetPerformanceFee(&_Yvault034.TransactOpts, fee)
}

// SetRewards is a paid mutator transaction binding the contract method 0xec38a862.
//
// Solidity: function setRewards(address rewards) returns()
func (_Yvault034 *Yvault034Transactor) SetRewards(opts *bind.TransactOpts, rewards common.Address) (*types.Transaction, error) {
	return _Yvault034.contract.Transact(opts, "setRewards", rewards)
}

// SetRewards is a paid mutator transaction binding the contract method 0xec38a862.
//
// Solidity: function setRewards(address rewards) returns()
func (_Yvault034 *Yvault034Session) SetRewards(rewards common.Address) (*types.Transaction, error) {
	return _Yvault034.Contract.SetRewards(&_Yvault034.TransactOpts, rewards)
}

// SetRewards is a paid mutator transaction binding the contract method 0xec38a862.
//
// Solidity: function setRewards(address rewards) returns()
func (_Yvault034 *Yvault034TransactorSession) SetRewards(rewards common.Address) (*types.Transaction, error) {
	return _Yvault034.Contract.SetRewards(&_Yvault034.TransactOpts, rewards)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string symbol) returns()
func (_Yvault034 *Yvault034Transactor) SetSymbol(opts *bind.TransactOpts, symbol string) (*types.Transaction, error) {
	return _Yvault034.contract.Transact(opts, "setSymbol", symbol)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string symbol) returns()
func (_Yvault034 *Yvault034Session) SetSymbol(symbol string) (*types.Transaction, error) {
	return _Yvault034.Contract.SetSymbol(&_Yvault034.TransactOpts, symbol)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string symbol) returns()
func (_Yvault034 *Yvault034TransactorSession) SetSymbol(symbol string) (*types.Transaction, error) {
	return _Yvault034.Contract.SetSymbol(&_Yvault034.TransactOpts, symbol)
}

// SetWithdrawalQueue is a paid mutator transaction binding the contract method 0x94148415.
//
// Solidity: function setWithdrawalQueue(address[20] queue) returns()
func (_Yvault034 *Yvault034Transactor) SetWithdrawalQueue(opts *bind.TransactOpts, queue [20]common.Address) (*types.Transaction, error) {
	return _Yvault034.contract.Transact(opts, "setWithdrawalQueue", queue)
}

// SetWithdrawalQueue is a paid mutator transaction binding the contract method 0x94148415.
//
// Solidity: function setWithdrawalQueue(address[20] queue) returns()
func (_Yvault034 *Yvault034Session) SetWithdrawalQueue(queue [20]common.Address) (*types.Transaction, error) {
	return _Yvault034.Contract.SetWithdrawalQueue(&_Yvault034.TransactOpts, queue)
}

// SetWithdrawalQueue is a paid mutator transaction binding the contract method 0x94148415.
//
// Solidity: function setWithdrawalQueue(address[20] queue) returns()
func (_Yvault034 *Yvault034TransactorSession) SetWithdrawalQueue(queue [20]common.Address) (*types.Transaction, error) {
	return _Yvault034.Contract.SetWithdrawalQueue(&_Yvault034.TransactOpts, queue)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address token) returns()
func (_Yvault034 *Yvault034Transactor) Sweep(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Yvault034.contract.Transact(opts, "sweep", token)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address token) returns()
func (_Yvault034 *Yvault034Session) Sweep(token common.Address) (*types.Transaction, error) {
	return _Yvault034.Contract.Sweep(&_Yvault034.TransactOpts, token)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address token) returns()
func (_Yvault034 *Yvault034TransactorSession) Sweep(token common.Address) (*types.Transaction, error) {
	return _Yvault034.Contract.Sweep(&_Yvault034.TransactOpts, token)
}

// Sweep0 is a paid mutator transaction binding the contract method 0x6ea056a9.
//
// Solidity: function sweep(address token, uint256 amount) returns()
func (_Yvault034 *Yvault034Transactor) Sweep0(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault034.contract.Transact(opts, "sweep0", token, amount)
}

// Sweep0 is a paid mutator transaction binding the contract method 0x6ea056a9.
//
// Solidity: function sweep(address token, uint256 amount) returns()
func (_Yvault034 *Yvault034Session) Sweep0(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault034.Contract.Sweep0(&_Yvault034.TransactOpts, token, amount)
}

// Sweep0 is a paid mutator transaction binding the contract method 0x6ea056a9.
//
// Solidity: function sweep(address token, uint256 amount) returns()
func (_Yvault034 *Yvault034TransactorSession) Sweep0(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault034.Contract.Sweep0(&_Yvault034.TransactOpts, token, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 amount) returns(bool)
func (_Yvault034 *Yvault034Transactor) Transfer(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault034.contract.Transact(opts, "transfer", receiver, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 amount) returns(bool)
func (_Yvault034 *Yvault034Session) Transfer(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault034.Contract.Transfer(&_Yvault034.TransactOpts, receiver, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 amount) returns(bool)
func (_Yvault034 *Yvault034TransactorSession) Transfer(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault034.Contract.Transfer(&_Yvault034.TransactOpts, receiver, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address receiver, uint256 amount) returns(bool)
func (_Yvault034 *Yvault034Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault034.contract.Transact(opts, "transferFrom", sender, receiver, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address receiver, uint256 amount) returns(bool)
func (_Yvault034 *Yvault034Session) TransferFrom(sender common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault034.Contract.TransferFrom(&_Yvault034.TransactOpts, sender, receiver, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address receiver, uint256 amount) returns(bool)
func (_Yvault034 *Yvault034TransactorSession) TransferFrom(sender common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault034.Contract.TransferFrom(&_Yvault034.TransactOpts, sender, receiver, amount)
}

// UpdateStrategyDebtRatio is a paid mutator transaction binding the contract method 0x7c6a4f24.
//
// Solidity: function updateStrategyDebtRatio(address strategy, uint256 debtRatio) returns()
func (_Yvault034 *Yvault034Transactor) UpdateStrategyDebtRatio(opts *bind.TransactOpts, strategy common.Address, debtRatio *big.Int) (*types.Transaction, error) {
	return _Yvault034.contract.Transact(opts, "updateStrategyDebtRatio", strategy, debtRatio)
}

// UpdateStrategyDebtRatio is a paid mutator transaction binding the contract method 0x7c6a4f24.
//
// Solidity: function updateStrategyDebtRatio(address strategy, uint256 debtRatio) returns()
func (_Yvault034 *Yvault034Session) UpdateStrategyDebtRatio(strategy common.Address, debtRatio *big.Int) (*types.Transaction, error) {
	return _Yvault034.Contract.UpdateStrategyDebtRatio(&_Yvault034.TransactOpts, strategy, debtRatio)
}

// UpdateStrategyDebtRatio is a paid mutator transaction binding the contract method 0x7c6a4f24.
//
// Solidity: function updateStrategyDebtRatio(address strategy, uint256 debtRatio) returns()
func (_Yvault034 *Yvault034TransactorSession) UpdateStrategyDebtRatio(strategy common.Address, debtRatio *big.Int) (*types.Transaction, error) {
	return _Yvault034.Contract.UpdateStrategyDebtRatio(&_Yvault034.TransactOpts, strategy, debtRatio)
}

// UpdateStrategyMaxDebtPerHarvest is a paid mutator transaction binding the contract method 0x4757a156.
//
// Solidity: function updateStrategyMaxDebtPerHarvest(address strategy, uint256 maxDebtPerHarvest) returns()
func (_Yvault034 *Yvault034Transactor) UpdateStrategyMaxDebtPerHarvest(opts *bind.TransactOpts, strategy common.Address, maxDebtPerHarvest *big.Int) (*types.Transaction, error) {
	return _Yvault034.contract.Transact(opts, "updateStrategyMaxDebtPerHarvest", strategy, maxDebtPerHarvest)
}

// UpdateStrategyMaxDebtPerHarvest is a paid mutator transaction binding the contract method 0x4757a156.
//
// Solidity: function updateStrategyMaxDebtPerHarvest(address strategy, uint256 maxDebtPerHarvest) returns()
func (_Yvault034 *Yvault034Session) UpdateStrategyMaxDebtPerHarvest(strategy common.Address, maxDebtPerHarvest *big.Int) (*types.Transaction, error) {
	return _Yvault034.Contract.UpdateStrategyMaxDebtPerHarvest(&_Yvault034.TransactOpts, strategy, maxDebtPerHarvest)
}

// UpdateStrategyMaxDebtPerHarvest is a paid mutator transaction binding the contract method 0x4757a156.
//
// Solidity: function updateStrategyMaxDebtPerHarvest(address strategy, uint256 maxDebtPerHarvest) returns()
func (_Yvault034 *Yvault034TransactorSession) UpdateStrategyMaxDebtPerHarvest(strategy common.Address, maxDebtPerHarvest *big.Int) (*types.Transaction, error) {
	return _Yvault034.Contract.UpdateStrategyMaxDebtPerHarvest(&_Yvault034.TransactOpts, strategy, maxDebtPerHarvest)
}

// UpdateStrategyMinDebtPerHarvest is a paid mutator transaction binding the contract method 0xe722befe.
//
// Solidity: function updateStrategyMinDebtPerHarvest(address strategy, uint256 minDebtPerHarvest) returns()
func (_Yvault034 *Yvault034Transactor) UpdateStrategyMinDebtPerHarvest(opts *bind.TransactOpts, strategy common.Address, minDebtPerHarvest *big.Int) (*types.Transaction, error) {
	return _Yvault034.contract.Transact(opts, "updateStrategyMinDebtPerHarvest", strategy, minDebtPerHarvest)
}

// UpdateStrategyMinDebtPerHarvest is a paid mutator transaction binding the contract method 0xe722befe.
//
// Solidity: function updateStrategyMinDebtPerHarvest(address strategy, uint256 minDebtPerHarvest) returns()
func (_Yvault034 *Yvault034Session) UpdateStrategyMinDebtPerHarvest(strategy common.Address, minDebtPerHarvest *big.Int) (*types.Transaction, error) {
	return _Yvault034.Contract.UpdateStrategyMinDebtPerHarvest(&_Yvault034.TransactOpts, strategy, minDebtPerHarvest)
}

// UpdateStrategyMinDebtPerHarvest is a paid mutator transaction binding the contract method 0xe722befe.
//
// Solidity: function updateStrategyMinDebtPerHarvest(address strategy, uint256 minDebtPerHarvest) returns()
func (_Yvault034 *Yvault034TransactorSession) UpdateStrategyMinDebtPerHarvest(strategy common.Address, minDebtPerHarvest *big.Int) (*types.Transaction, error) {
	return _Yvault034.Contract.UpdateStrategyMinDebtPerHarvest(&_Yvault034.TransactOpts, strategy, minDebtPerHarvest)
}

// UpdateStrategyPerformanceFee is a paid mutator transaction binding the contract method 0xd0194ed6.
//
// Solidity: function updateStrategyPerformanceFee(address strategy, uint256 performanceFee) returns()
func (_Yvault034 *Yvault034Transactor) UpdateStrategyPerformanceFee(opts *bind.TransactOpts, strategy common.Address, performanceFee *big.Int) (*types.Transaction, error) {
	return _Yvault034.contract.Transact(opts, "updateStrategyPerformanceFee", strategy, performanceFee)
}

// UpdateStrategyPerformanceFee is a paid mutator transaction binding the contract method 0xd0194ed6.
//
// Solidity: function updateStrategyPerformanceFee(address strategy, uint256 performanceFee) returns()
func (_Yvault034 *Yvault034Session) UpdateStrategyPerformanceFee(strategy common.Address, performanceFee *big.Int) (*types.Transaction, error) {
	return _Yvault034.Contract.UpdateStrategyPerformanceFee(&_Yvault034.TransactOpts, strategy, performanceFee)
}

// UpdateStrategyPerformanceFee is a paid mutator transaction binding the contract method 0xd0194ed6.
//
// Solidity: function updateStrategyPerformanceFee(address strategy, uint256 performanceFee) returns()
func (_Yvault034 *Yvault034TransactorSession) UpdateStrategyPerformanceFee(strategy common.Address, performanceFee *big.Int) (*types.Transaction, error) {
	return _Yvault034.Contract.UpdateStrategyPerformanceFee(&_Yvault034.TransactOpts, strategy, performanceFee)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(uint256)
func (_Yvault034 *Yvault034Transactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yvault034.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(uint256)
func (_Yvault034 *Yvault034Session) Withdraw() (*types.Transaction, error) {
	return _Yvault034.Contract.Withdraw(&_Yvault034.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(uint256)
func (_Yvault034 *Yvault034TransactorSession) Withdraw() (*types.Transaction, error) {
	return _Yvault034.Contract.Withdraw(&_Yvault034.TransactOpts)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 maxShares) returns(uint256)
func (_Yvault034 *Yvault034Transactor) Withdraw0(opts *bind.TransactOpts, maxShares *big.Int) (*types.Transaction, error) {
	return _Yvault034.contract.Transact(opts, "withdraw0", maxShares)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 maxShares) returns(uint256)
func (_Yvault034 *Yvault034Session) Withdraw0(maxShares *big.Int) (*types.Transaction, error) {
	return _Yvault034.Contract.Withdraw0(&_Yvault034.TransactOpts, maxShares)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 maxShares) returns(uint256)
func (_Yvault034 *Yvault034TransactorSession) Withdraw0(maxShares *big.Int) (*types.Transaction, error) {
	return _Yvault034.Contract.Withdraw0(&_Yvault034.TransactOpts, maxShares)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 maxShares, address recipient) returns(uint256)
func (_Yvault034 *Yvault034Transactor) Withdraw1(opts *bind.TransactOpts, maxShares *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Yvault034.contract.Transact(opts, "withdraw1", maxShares, recipient)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 maxShares, address recipient) returns(uint256)
func (_Yvault034 *Yvault034Session) Withdraw1(maxShares *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Yvault034.Contract.Withdraw1(&_Yvault034.TransactOpts, maxShares, recipient)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 maxShares, address recipient) returns(uint256)
func (_Yvault034 *Yvault034TransactorSession) Withdraw1(maxShares *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Yvault034.Contract.Withdraw1(&_Yvault034.TransactOpts, maxShares, recipient)
}

// Withdraw2 is a paid mutator transaction binding the contract method 0xe63697c8.
//
// Solidity: function withdraw(uint256 maxShares, address recipient, uint256 maxLoss) returns(uint256)
func (_Yvault034 *Yvault034Transactor) Withdraw2(opts *bind.TransactOpts, maxShares *big.Int, recipient common.Address, maxLoss *big.Int) (*types.Transaction, error) {
	return _Yvault034.contract.Transact(opts, "withdraw2", maxShares, recipient, maxLoss)
}

// Withdraw2 is a paid mutator transaction binding the contract method 0xe63697c8.
//
// Solidity: function withdraw(uint256 maxShares, address recipient, uint256 maxLoss) returns(uint256)
func (_Yvault034 *Yvault034Session) Withdraw2(maxShares *big.Int, recipient common.Address, maxLoss *big.Int) (*types.Transaction, error) {
	return _Yvault034.Contract.Withdraw2(&_Yvault034.TransactOpts, maxShares, recipient, maxLoss)
}

// Withdraw2 is a paid mutator transaction binding the contract method 0xe63697c8.
//
// Solidity: function withdraw(uint256 maxShares, address recipient, uint256 maxLoss) returns(uint256)
func (_Yvault034 *Yvault034TransactorSession) Withdraw2(maxShares *big.Int, recipient common.Address, maxLoss *big.Int) (*types.Transaction, error) {
	return _Yvault034.Contract.Withdraw2(&_Yvault034.TransactOpts, maxShares, recipient, maxLoss)
}

// Yvault034ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Yvault034 contract.
type Yvault034ApprovalIterator struct {
	Event *Yvault034Approval // Event containing the contract specifics and raw log

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
func (it *Yvault034ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault034Approval)
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
		it.Event = new(Yvault034Approval)
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
func (it *Yvault034ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault034ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault034Approval represents a Approval event raised by the Yvault034 contract.
type Yvault034Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Yvault034 *Yvault034Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Yvault034ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Yvault034.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Yvault034ApprovalIterator{contract: _Yvault034.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Yvault034 *Yvault034Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Yvault034Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Yvault034.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault034Approval)
				if err := _Yvault034.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Yvault034 *Yvault034Filterer) ParseApproval(log types.Log) (*Yvault034Approval, error) {
	event := new(Yvault034Approval)
	if err := _Yvault034.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault034EmergencyShutdownIterator is returned from FilterEmergencyShutdown and is used to iterate over the raw logs and unpacked data for EmergencyShutdown events raised by the Yvault034 contract.
type Yvault034EmergencyShutdownIterator struct {
	Event *Yvault034EmergencyShutdown // Event containing the contract specifics and raw log

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
func (it *Yvault034EmergencyShutdownIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault034EmergencyShutdown)
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
		it.Event = new(Yvault034EmergencyShutdown)
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
func (it *Yvault034EmergencyShutdownIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault034EmergencyShutdownIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault034EmergencyShutdown represents a EmergencyShutdown event raised by the Yvault034 contract.
type Yvault034EmergencyShutdown struct {
	Active bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyShutdown is a free log retrieval operation binding the contract event 0xba40372a3a724dca3c57156128ef1e896724b65b37a17f190b1ad5de68f3a4f3.
//
// Solidity: event EmergencyShutdown(bool active)
func (_Yvault034 *Yvault034Filterer) FilterEmergencyShutdown(opts *bind.FilterOpts) (*Yvault034EmergencyShutdownIterator, error) {

	logs, sub, err := _Yvault034.contract.FilterLogs(opts, "EmergencyShutdown")
	if err != nil {
		return nil, err
	}
	return &Yvault034EmergencyShutdownIterator{contract: _Yvault034.contract, event: "EmergencyShutdown", logs: logs, sub: sub}, nil
}

// WatchEmergencyShutdown is a free log subscription operation binding the contract event 0xba40372a3a724dca3c57156128ef1e896724b65b37a17f190b1ad5de68f3a4f3.
//
// Solidity: event EmergencyShutdown(bool active)
func (_Yvault034 *Yvault034Filterer) WatchEmergencyShutdown(opts *bind.WatchOpts, sink chan<- *Yvault034EmergencyShutdown) (event.Subscription, error) {

	logs, sub, err := _Yvault034.contract.WatchLogs(opts, "EmergencyShutdown")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault034EmergencyShutdown)
				if err := _Yvault034.contract.UnpackLog(event, "EmergencyShutdown", log); err != nil {
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
func (_Yvault034 *Yvault034Filterer) ParseEmergencyShutdown(log types.Log) (*Yvault034EmergencyShutdown, error) {
	event := new(Yvault034EmergencyShutdown)
	if err := _Yvault034.contract.UnpackLog(event, "EmergencyShutdown", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault034StrategyAddedIterator is returned from FilterStrategyAdded and is used to iterate over the raw logs and unpacked data for StrategyAdded events raised by the Yvault034 contract.
type Yvault034StrategyAddedIterator struct {
	Event *Yvault034StrategyAdded // Event containing the contract specifics and raw log

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
func (it *Yvault034StrategyAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault034StrategyAdded)
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
		it.Event = new(Yvault034StrategyAdded)
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
func (it *Yvault034StrategyAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault034StrategyAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault034StrategyAdded represents a StrategyAdded event raised by the Yvault034 contract.
type Yvault034StrategyAdded struct {
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
func (_Yvault034 *Yvault034Filterer) FilterStrategyAdded(opts *bind.FilterOpts, strategy []common.Address) (*Yvault034StrategyAddedIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault034.contract.FilterLogs(opts, "StrategyAdded", strategyRule)
	if err != nil {
		return nil, err
	}
	return &Yvault034StrategyAddedIterator{contract: _Yvault034.contract, event: "StrategyAdded", logs: logs, sub: sub}, nil
}

// WatchStrategyAdded is a free log subscription operation binding the contract event 0x5a6abd2af9fe6c0554fa08649e2d86e4393ff19dc304d072d38d295c9291d4dc.
//
// Solidity: event StrategyAdded(address indexed strategy, uint256 debtRatio, uint256 minDebtPerHarvest, uint256 maxDebtPerHarvest, uint256 performanceFee)
func (_Yvault034 *Yvault034Filterer) WatchStrategyAdded(opts *bind.WatchOpts, sink chan<- *Yvault034StrategyAdded, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault034.contract.WatchLogs(opts, "StrategyAdded", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault034StrategyAdded)
				if err := _Yvault034.contract.UnpackLog(event, "StrategyAdded", log); err != nil {
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
func (_Yvault034 *Yvault034Filterer) ParseStrategyAdded(log types.Log) (*Yvault034StrategyAdded, error) {
	event := new(Yvault034StrategyAdded)
	if err := _Yvault034.contract.UnpackLog(event, "StrategyAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault034StrategyAddedToQueueIterator is returned from FilterStrategyAddedToQueue and is used to iterate over the raw logs and unpacked data for StrategyAddedToQueue events raised by the Yvault034 contract.
type Yvault034StrategyAddedToQueueIterator struct {
	Event *Yvault034StrategyAddedToQueue // Event containing the contract specifics and raw log

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
func (it *Yvault034StrategyAddedToQueueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault034StrategyAddedToQueue)
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
		it.Event = new(Yvault034StrategyAddedToQueue)
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
func (it *Yvault034StrategyAddedToQueueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault034StrategyAddedToQueueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault034StrategyAddedToQueue represents a StrategyAddedToQueue event raised by the Yvault034 contract.
type Yvault034StrategyAddedToQueue struct {
	Strategy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyAddedToQueue is a free log retrieval operation binding the contract event 0xa8727d412c6fa1e2497d6d6f275e2d9fe4d9318d5b793632e60ad9d38ee8f1fa.
//
// Solidity: event StrategyAddedToQueue(address indexed strategy)
func (_Yvault034 *Yvault034Filterer) FilterStrategyAddedToQueue(opts *bind.FilterOpts, strategy []common.Address) (*Yvault034StrategyAddedToQueueIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault034.contract.FilterLogs(opts, "StrategyAddedToQueue", strategyRule)
	if err != nil {
		return nil, err
	}
	return &Yvault034StrategyAddedToQueueIterator{contract: _Yvault034.contract, event: "StrategyAddedToQueue", logs: logs, sub: sub}, nil
}

// WatchStrategyAddedToQueue is a free log subscription operation binding the contract event 0xa8727d412c6fa1e2497d6d6f275e2d9fe4d9318d5b793632e60ad9d38ee8f1fa.
//
// Solidity: event StrategyAddedToQueue(address indexed strategy)
func (_Yvault034 *Yvault034Filterer) WatchStrategyAddedToQueue(opts *bind.WatchOpts, sink chan<- *Yvault034StrategyAddedToQueue, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault034.contract.WatchLogs(opts, "StrategyAddedToQueue", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault034StrategyAddedToQueue)
				if err := _Yvault034.contract.UnpackLog(event, "StrategyAddedToQueue", log); err != nil {
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
func (_Yvault034 *Yvault034Filterer) ParseStrategyAddedToQueue(log types.Log) (*Yvault034StrategyAddedToQueue, error) {
	event := new(Yvault034StrategyAddedToQueue)
	if err := _Yvault034.contract.UnpackLog(event, "StrategyAddedToQueue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault034StrategyMigratedIterator is returned from FilterStrategyMigrated and is used to iterate over the raw logs and unpacked data for StrategyMigrated events raised by the Yvault034 contract.
type Yvault034StrategyMigratedIterator struct {
	Event *Yvault034StrategyMigrated // Event containing the contract specifics and raw log

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
func (it *Yvault034StrategyMigratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault034StrategyMigrated)
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
		it.Event = new(Yvault034StrategyMigrated)
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
func (it *Yvault034StrategyMigratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault034StrategyMigratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault034StrategyMigrated represents a StrategyMigrated event raised by the Yvault034 contract.
type Yvault034StrategyMigrated struct {
	OldVersion common.Address
	NewVersion common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterStrategyMigrated is a free log retrieval operation binding the contract event 0x100b69bb6b504e1252e36b375233158edee64d071b399e2f81473a695fd1b021.
//
// Solidity: event StrategyMigrated(address indexed oldVersion, address indexed newVersion)
func (_Yvault034 *Yvault034Filterer) FilterStrategyMigrated(opts *bind.FilterOpts, oldVersion []common.Address, newVersion []common.Address) (*Yvault034StrategyMigratedIterator, error) {

	var oldVersionRule []interface{}
	for _, oldVersionItem := range oldVersion {
		oldVersionRule = append(oldVersionRule, oldVersionItem)
	}
	var newVersionRule []interface{}
	for _, newVersionItem := range newVersion {
		newVersionRule = append(newVersionRule, newVersionItem)
	}

	logs, sub, err := _Yvault034.contract.FilterLogs(opts, "StrategyMigrated", oldVersionRule, newVersionRule)
	if err != nil {
		return nil, err
	}
	return &Yvault034StrategyMigratedIterator{contract: _Yvault034.contract, event: "StrategyMigrated", logs: logs, sub: sub}, nil
}

// WatchStrategyMigrated is a free log subscription operation binding the contract event 0x100b69bb6b504e1252e36b375233158edee64d071b399e2f81473a695fd1b021.
//
// Solidity: event StrategyMigrated(address indexed oldVersion, address indexed newVersion)
func (_Yvault034 *Yvault034Filterer) WatchStrategyMigrated(opts *bind.WatchOpts, sink chan<- *Yvault034StrategyMigrated, oldVersion []common.Address, newVersion []common.Address) (event.Subscription, error) {

	var oldVersionRule []interface{}
	for _, oldVersionItem := range oldVersion {
		oldVersionRule = append(oldVersionRule, oldVersionItem)
	}
	var newVersionRule []interface{}
	for _, newVersionItem := range newVersion {
		newVersionRule = append(newVersionRule, newVersionItem)
	}

	logs, sub, err := _Yvault034.contract.WatchLogs(opts, "StrategyMigrated", oldVersionRule, newVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault034StrategyMigrated)
				if err := _Yvault034.contract.UnpackLog(event, "StrategyMigrated", log); err != nil {
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
func (_Yvault034 *Yvault034Filterer) ParseStrategyMigrated(log types.Log) (*Yvault034StrategyMigrated, error) {
	event := new(Yvault034StrategyMigrated)
	if err := _Yvault034.contract.UnpackLog(event, "StrategyMigrated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault034StrategyRemovedFromQueueIterator is returned from FilterStrategyRemovedFromQueue and is used to iterate over the raw logs and unpacked data for StrategyRemovedFromQueue events raised by the Yvault034 contract.
type Yvault034StrategyRemovedFromQueueIterator struct {
	Event *Yvault034StrategyRemovedFromQueue // Event containing the contract specifics and raw log

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
func (it *Yvault034StrategyRemovedFromQueueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault034StrategyRemovedFromQueue)
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
		it.Event = new(Yvault034StrategyRemovedFromQueue)
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
func (it *Yvault034StrategyRemovedFromQueueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault034StrategyRemovedFromQueueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault034StrategyRemovedFromQueue represents a StrategyRemovedFromQueue event raised by the Yvault034 contract.
type Yvault034StrategyRemovedFromQueue struct {
	Strategy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyRemovedFromQueue is a free log retrieval operation binding the contract event 0x8e1ec3c16d6a67ea8effe2ac7adef9c2de0bc0dc47c49cdf18f6a8b0048085be.
//
// Solidity: event StrategyRemovedFromQueue(address indexed strategy)
func (_Yvault034 *Yvault034Filterer) FilterStrategyRemovedFromQueue(opts *bind.FilterOpts, strategy []common.Address) (*Yvault034StrategyRemovedFromQueueIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault034.contract.FilterLogs(opts, "StrategyRemovedFromQueue", strategyRule)
	if err != nil {
		return nil, err
	}
	return &Yvault034StrategyRemovedFromQueueIterator{contract: _Yvault034.contract, event: "StrategyRemovedFromQueue", logs: logs, sub: sub}, nil
}

// WatchStrategyRemovedFromQueue is a free log subscription operation binding the contract event 0x8e1ec3c16d6a67ea8effe2ac7adef9c2de0bc0dc47c49cdf18f6a8b0048085be.
//
// Solidity: event StrategyRemovedFromQueue(address indexed strategy)
func (_Yvault034 *Yvault034Filterer) WatchStrategyRemovedFromQueue(opts *bind.WatchOpts, sink chan<- *Yvault034StrategyRemovedFromQueue, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault034.contract.WatchLogs(opts, "StrategyRemovedFromQueue", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault034StrategyRemovedFromQueue)
				if err := _Yvault034.contract.UnpackLog(event, "StrategyRemovedFromQueue", log); err != nil {
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
func (_Yvault034 *Yvault034Filterer) ParseStrategyRemovedFromQueue(log types.Log) (*Yvault034StrategyRemovedFromQueue, error) {
	event := new(Yvault034StrategyRemovedFromQueue)
	if err := _Yvault034.contract.UnpackLog(event, "StrategyRemovedFromQueue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault034StrategyReportedIterator is returned from FilterStrategyReported and is used to iterate over the raw logs and unpacked data for StrategyReported events raised by the Yvault034 contract.
type Yvault034StrategyReportedIterator struct {
	Event *Yvault034StrategyReported // Event containing the contract specifics and raw log

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
func (it *Yvault034StrategyReportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault034StrategyReported)
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
		it.Event = new(Yvault034StrategyReported)
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
func (it *Yvault034StrategyReportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault034StrategyReportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault034StrategyReported represents a StrategyReported event raised by the Yvault034 contract.
type Yvault034StrategyReported struct {
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
func (_Yvault034 *Yvault034Filterer) FilterStrategyReported(opts *bind.FilterOpts, strategy []common.Address) (*Yvault034StrategyReportedIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault034.contract.FilterLogs(opts, "StrategyReported", strategyRule)
	if err != nil {
		return nil, err
	}
	return &Yvault034StrategyReportedIterator{contract: _Yvault034.contract, event: "StrategyReported", logs: logs, sub: sub}, nil
}

// WatchStrategyReported is a free log subscription operation binding the contract event 0x67f96d2854a335a4cadb49f84fd3ca6f990744ddb3feceeb4b349d2d53d32ad3.
//
// Solidity: event StrategyReported(address indexed strategy, uint256 gain, uint256 loss, uint256 debtPaid, uint256 totalGain, uint256 totalLoss, uint256 totalDebt, uint256 debtAdded, uint256 debtRatio)
func (_Yvault034 *Yvault034Filterer) WatchStrategyReported(opts *bind.WatchOpts, sink chan<- *Yvault034StrategyReported, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault034.contract.WatchLogs(opts, "StrategyReported", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault034StrategyReported)
				if err := _Yvault034.contract.UnpackLog(event, "StrategyReported", log); err != nil {
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
func (_Yvault034 *Yvault034Filterer) ParseStrategyReported(log types.Log) (*Yvault034StrategyReported, error) {
	event := new(Yvault034StrategyReported)
	if err := _Yvault034.contract.UnpackLog(event, "StrategyReported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault034StrategyRevokedIterator is returned from FilterStrategyRevoked and is used to iterate over the raw logs and unpacked data for StrategyRevoked events raised by the Yvault034 contract.
type Yvault034StrategyRevokedIterator struct {
	Event *Yvault034StrategyRevoked // Event containing the contract specifics and raw log

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
func (it *Yvault034StrategyRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault034StrategyRevoked)
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
		it.Event = new(Yvault034StrategyRevoked)
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
func (it *Yvault034StrategyRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault034StrategyRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault034StrategyRevoked represents a StrategyRevoked event raised by the Yvault034 contract.
type Yvault034StrategyRevoked struct {
	Strategy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyRevoked is a free log retrieval operation binding the contract event 0x4201c688d84c01154d321afa0c72f1bffe9eef53005c9de9d035074e71e9b32a.
//
// Solidity: event StrategyRevoked(address indexed strategy)
func (_Yvault034 *Yvault034Filterer) FilterStrategyRevoked(opts *bind.FilterOpts, strategy []common.Address) (*Yvault034StrategyRevokedIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault034.contract.FilterLogs(opts, "StrategyRevoked", strategyRule)
	if err != nil {
		return nil, err
	}
	return &Yvault034StrategyRevokedIterator{contract: _Yvault034.contract, event: "StrategyRevoked", logs: logs, sub: sub}, nil
}

// WatchStrategyRevoked is a free log subscription operation binding the contract event 0x4201c688d84c01154d321afa0c72f1bffe9eef53005c9de9d035074e71e9b32a.
//
// Solidity: event StrategyRevoked(address indexed strategy)
func (_Yvault034 *Yvault034Filterer) WatchStrategyRevoked(opts *bind.WatchOpts, sink chan<- *Yvault034StrategyRevoked, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault034.contract.WatchLogs(opts, "StrategyRevoked", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault034StrategyRevoked)
				if err := _Yvault034.contract.UnpackLog(event, "StrategyRevoked", log); err != nil {
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
func (_Yvault034 *Yvault034Filterer) ParseStrategyRevoked(log types.Log) (*Yvault034StrategyRevoked, error) {
	event := new(Yvault034StrategyRevoked)
	if err := _Yvault034.contract.UnpackLog(event, "StrategyRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault034StrategyUpdateDebtRatioIterator is returned from FilterStrategyUpdateDebtRatio and is used to iterate over the raw logs and unpacked data for StrategyUpdateDebtRatio events raised by the Yvault034 contract.
type Yvault034StrategyUpdateDebtRatioIterator struct {
	Event *Yvault034StrategyUpdateDebtRatio // Event containing the contract specifics and raw log

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
func (it *Yvault034StrategyUpdateDebtRatioIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault034StrategyUpdateDebtRatio)
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
		it.Event = new(Yvault034StrategyUpdateDebtRatio)
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
func (it *Yvault034StrategyUpdateDebtRatioIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault034StrategyUpdateDebtRatioIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault034StrategyUpdateDebtRatio represents a StrategyUpdateDebtRatio event raised by the Yvault034 contract.
type Yvault034StrategyUpdateDebtRatio struct {
	Strategy  common.Address
	DebtRatio *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStrategyUpdateDebtRatio is a free log retrieval operation binding the contract event 0xbda9398315c83ccef012bcaa318a2ff7b680f36429d36597bd4bc25ac11ead59.
//
// Solidity: event StrategyUpdateDebtRatio(address indexed strategy, uint256 debtRatio)
func (_Yvault034 *Yvault034Filterer) FilterStrategyUpdateDebtRatio(opts *bind.FilterOpts, strategy []common.Address) (*Yvault034StrategyUpdateDebtRatioIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault034.contract.FilterLogs(opts, "StrategyUpdateDebtRatio", strategyRule)
	if err != nil {
		return nil, err
	}
	return &Yvault034StrategyUpdateDebtRatioIterator{contract: _Yvault034.contract, event: "StrategyUpdateDebtRatio", logs: logs, sub: sub}, nil
}

// WatchStrategyUpdateDebtRatio is a free log subscription operation binding the contract event 0xbda9398315c83ccef012bcaa318a2ff7b680f36429d36597bd4bc25ac11ead59.
//
// Solidity: event StrategyUpdateDebtRatio(address indexed strategy, uint256 debtRatio)
func (_Yvault034 *Yvault034Filterer) WatchStrategyUpdateDebtRatio(opts *bind.WatchOpts, sink chan<- *Yvault034StrategyUpdateDebtRatio, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault034.contract.WatchLogs(opts, "StrategyUpdateDebtRatio", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault034StrategyUpdateDebtRatio)
				if err := _Yvault034.contract.UnpackLog(event, "StrategyUpdateDebtRatio", log); err != nil {
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
func (_Yvault034 *Yvault034Filterer) ParseStrategyUpdateDebtRatio(log types.Log) (*Yvault034StrategyUpdateDebtRatio, error) {
	event := new(Yvault034StrategyUpdateDebtRatio)
	if err := _Yvault034.contract.UnpackLog(event, "StrategyUpdateDebtRatio", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault034StrategyUpdateMaxDebtPerHarvestIterator is returned from FilterStrategyUpdateMaxDebtPerHarvest and is used to iterate over the raw logs and unpacked data for StrategyUpdateMaxDebtPerHarvest events raised by the Yvault034 contract.
type Yvault034StrategyUpdateMaxDebtPerHarvestIterator struct {
	Event *Yvault034StrategyUpdateMaxDebtPerHarvest // Event containing the contract specifics and raw log

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
func (it *Yvault034StrategyUpdateMaxDebtPerHarvestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault034StrategyUpdateMaxDebtPerHarvest)
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
		it.Event = new(Yvault034StrategyUpdateMaxDebtPerHarvest)
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
func (it *Yvault034StrategyUpdateMaxDebtPerHarvestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault034StrategyUpdateMaxDebtPerHarvestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault034StrategyUpdateMaxDebtPerHarvest represents a StrategyUpdateMaxDebtPerHarvest event raised by the Yvault034 contract.
type Yvault034StrategyUpdateMaxDebtPerHarvest struct {
	Strategy          common.Address
	MaxDebtPerHarvest *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterStrategyUpdateMaxDebtPerHarvest is a free log retrieval operation binding the contract event 0x1796a8e0760e2de5b72e7bf64fccb7666c48ceab94cb6cae7cb7eff4b6f641ab.
//
// Solidity: event StrategyUpdateMaxDebtPerHarvest(address indexed strategy, uint256 maxDebtPerHarvest)
func (_Yvault034 *Yvault034Filterer) FilterStrategyUpdateMaxDebtPerHarvest(opts *bind.FilterOpts, strategy []common.Address) (*Yvault034StrategyUpdateMaxDebtPerHarvestIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault034.contract.FilterLogs(opts, "StrategyUpdateMaxDebtPerHarvest", strategyRule)
	if err != nil {
		return nil, err
	}
	return &Yvault034StrategyUpdateMaxDebtPerHarvestIterator{contract: _Yvault034.contract, event: "StrategyUpdateMaxDebtPerHarvest", logs: logs, sub: sub}, nil
}

// WatchStrategyUpdateMaxDebtPerHarvest is a free log subscription operation binding the contract event 0x1796a8e0760e2de5b72e7bf64fccb7666c48ceab94cb6cae7cb7eff4b6f641ab.
//
// Solidity: event StrategyUpdateMaxDebtPerHarvest(address indexed strategy, uint256 maxDebtPerHarvest)
func (_Yvault034 *Yvault034Filterer) WatchStrategyUpdateMaxDebtPerHarvest(opts *bind.WatchOpts, sink chan<- *Yvault034StrategyUpdateMaxDebtPerHarvest, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault034.contract.WatchLogs(opts, "StrategyUpdateMaxDebtPerHarvest", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault034StrategyUpdateMaxDebtPerHarvest)
				if err := _Yvault034.contract.UnpackLog(event, "StrategyUpdateMaxDebtPerHarvest", log); err != nil {
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
func (_Yvault034 *Yvault034Filterer) ParseStrategyUpdateMaxDebtPerHarvest(log types.Log) (*Yvault034StrategyUpdateMaxDebtPerHarvest, error) {
	event := new(Yvault034StrategyUpdateMaxDebtPerHarvest)
	if err := _Yvault034.contract.UnpackLog(event, "StrategyUpdateMaxDebtPerHarvest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault034StrategyUpdateMinDebtPerHarvestIterator is returned from FilterStrategyUpdateMinDebtPerHarvest and is used to iterate over the raw logs and unpacked data for StrategyUpdateMinDebtPerHarvest events raised by the Yvault034 contract.
type Yvault034StrategyUpdateMinDebtPerHarvestIterator struct {
	Event *Yvault034StrategyUpdateMinDebtPerHarvest // Event containing the contract specifics and raw log

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
func (it *Yvault034StrategyUpdateMinDebtPerHarvestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault034StrategyUpdateMinDebtPerHarvest)
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
		it.Event = new(Yvault034StrategyUpdateMinDebtPerHarvest)
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
func (it *Yvault034StrategyUpdateMinDebtPerHarvestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault034StrategyUpdateMinDebtPerHarvestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault034StrategyUpdateMinDebtPerHarvest represents a StrategyUpdateMinDebtPerHarvest event raised by the Yvault034 contract.
type Yvault034StrategyUpdateMinDebtPerHarvest struct {
	Strategy          common.Address
	MinDebtPerHarvest *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterStrategyUpdateMinDebtPerHarvest is a free log retrieval operation binding the contract event 0x0b728ad785976532c4aaadde09b1cba5f262a7090e83c62d2377bc405678b29c.
//
// Solidity: event StrategyUpdateMinDebtPerHarvest(address indexed strategy, uint256 minDebtPerHarvest)
func (_Yvault034 *Yvault034Filterer) FilterStrategyUpdateMinDebtPerHarvest(opts *bind.FilterOpts, strategy []common.Address) (*Yvault034StrategyUpdateMinDebtPerHarvestIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault034.contract.FilterLogs(opts, "StrategyUpdateMinDebtPerHarvest", strategyRule)
	if err != nil {
		return nil, err
	}
	return &Yvault034StrategyUpdateMinDebtPerHarvestIterator{contract: _Yvault034.contract, event: "StrategyUpdateMinDebtPerHarvest", logs: logs, sub: sub}, nil
}

// WatchStrategyUpdateMinDebtPerHarvest is a free log subscription operation binding the contract event 0x0b728ad785976532c4aaadde09b1cba5f262a7090e83c62d2377bc405678b29c.
//
// Solidity: event StrategyUpdateMinDebtPerHarvest(address indexed strategy, uint256 minDebtPerHarvest)
func (_Yvault034 *Yvault034Filterer) WatchStrategyUpdateMinDebtPerHarvest(opts *bind.WatchOpts, sink chan<- *Yvault034StrategyUpdateMinDebtPerHarvest, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault034.contract.WatchLogs(opts, "StrategyUpdateMinDebtPerHarvest", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault034StrategyUpdateMinDebtPerHarvest)
				if err := _Yvault034.contract.UnpackLog(event, "StrategyUpdateMinDebtPerHarvest", log); err != nil {
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
func (_Yvault034 *Yvault034Filterer) ParseStrategyUpdateMinDebtPerHarvest(log types.Log) (*Yvault034StrategyUpdateMinDebtPerHarvest, error) {
	event := new(Yvault034StrategyUpdateMinDebtPerHarvest)
	if err := _Yvault034.contract.UnpackLog(event, "StrategyUpdateMinDebtPerHarvest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault034StrategyUpdatePerformanceFeeIterator is returned from FilterStrategyUpdatePerformanceFee and is used to iterate over the raw logs and unpacked data for StrategyUpdatePerformanceFee events raised by the Yvault034 contract.
type Yvault034StrategyUpdatePerformanceFeeIterator struct {
	Event *Yvault034StrategyUpdatePerformanceFee // Event containing the contract specifics and raw log

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
func (it *Yvault034StrategyUpdatePerformanceFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault034StrategyUpdatePerformanceFee)
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
		it.Event = new(Yvault034StrategyUpdatePerformanceFee)
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
func (it *Yvault034StrategyUpdatePerformanceFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault034StrategyUpdatePerformanceFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault034StrategyUpdatePerformanceFee represents a StrategyUpdatePerformanceFee event raised by the Yvault034 contract.
type Yvault034StrategyUpdatePerformanceFee struct {
	Strategy       common.Address
	PerformanceFee *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterStrategyUpdatePerformanceFee is a free log retrieval operation binding the contract event 0xe57488a65fa53066d4c25bac90db47dda4e5de3025ac12bf76ff07211cf7f39e.
//
// Solidity: event StrategyUpdatePerformanceFee(address indexed strategy, uint256 performanceFee)
func (_Yvault034 *Yvault034Filterer) FilterStrategyUpdatePerformanceFee(opts *bind.FilterOpts, strategy []common.Address) (*Yvault034StrategyUpdatePerformanceFeeIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault034.contract.FilterLogs(opts, "StrategyUpdatePerformanceFee", strategyRule)
	if err != nil {
		return nil, err
	}
	return &Yvault034StrategyUpdatePerformanceFeeIterator{contract: _Yvault034.contract, event: "StrategyUpdatePerformanceFee", logs: logs, sub: sub}, nil
}

// WatchStrategyUpdatePerformanceFee is a free log subscription operation binding the contract event 0xe57488a65fa53066d4c25bac90db47dda4e5de3025ac12bf76ff07211cf7f39e.
//
// Solidity: event StrategyUpdatePerformanceFee(address indexed strategy, uint256 performanceFee)
func (_Yvault034 *Yvault034Filterer) WatchStrategyUpdatePerformanceFee(opts *bind.WatchOpts, sink chan<- *Yvault034StrategyUpdatePerformanceFee, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault034.contract.WatchLogs(opts, "StrategyUpdatePerformanceFee", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault034StrategyUpdatePerformanceFee)
				if err := _Yvault034.contract.UnpackLog(event, "StrategyUpdatePerformanceFee", log); err != nil {
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
func (_Yvault034 *Yvault034Filterer) ParseStrategyUpdatePerformanceFee(log types.Log) (*Yvault034StrategyUpdatePerformanceFee, error) {
	event := new(Yvault034StrategyUpdatePerformanceFee)
	if err := _Yvault034.contract.UnpackLog(event, "StrategyUpdatePerformanceFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault034TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Yvault034 contract.
type Yvault034TransferIterator struct {
	Event *Yvault034Transfer // Event containing the contract specifics and raw log

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
func (it *Yvault034TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault034Transfer)
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
		it.Event = new(Yvault034Transfer)
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
func (it *Yvault034TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault034TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault034Transfer represents a Transfer event raised by the Yvault034 contract.
type Yvault034Transfer struct {
	Sender   common.Address
	Receiver common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_Yvault034 *Yvault034Filterer) FilterTransfer(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*Yvault034TransferIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Yvault034.contract.FilterLogs(opts, "Transfer", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &Yvault034TransferIterator{contract: _Yvault034.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_Yvault034 *Yvault034Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Yvault034Transfer, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Yvault034.contract.WatchLogs(opts, "Transfer", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault034Transfer)
				if err := _Yvault034.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Yvault034 *Yvault034Filterer) ParseTransfer(log types.Log) (*Yvault034Transfer, error) {
	event := new(Yvault034Transfer)
	if err := _Yvault034.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault034UpdateDepositLimitIterator is returned from FilterUpdateDepositLimit and is used to iterate over the raw logs and unpacked data for UpdateDepositLimit events raised by the Yvault034 contract.
type Yvault034UpdateDepositLimitIterator struct {
	Event *Yvault034UpdateDepositLimit // Event containing the contract specifics and raw log

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
func (it *Yvault034UpdateDepositLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault034UpdateDepositLimit)
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
		it.Event = new(Yvault034UpdateDepositLimit)
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
func (it *Yvault034UpdateDepositLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault034UpdateDepositLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault034UpdateDepositLimit represents a UpdateDepositLimit event raised by the Yvault034 contract.
type Yvault034UpdateDepositLimit struct {
	DepositLimit *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateDepositLimit is a free log retrieval operation binding the contract event 0xae565aab888bca5e19e25a13db7b0c9144305bf55cb0f3f4d724f730e5acdd62.
//
// Solidity: event UpdateDepositLimit(uint256 depositLimit)
func (_Yvault034 *Yvault034Filterer) FilterUpdateDepositLimit(opts *bind.FilterOpts) (*Yvault034UpdateDepositLimitIterator, error) {

	logs, sub, err := _Yvault034.contract.FilterLogs(opts, "UpdateDepositLimit")
	if err != nil {
		return nil, err
	}
	return &Yvault034UpdateDepositLimitIterator{contract: _Yvault034.contract, event: "UpdateDepositLimit", logs: logs, sub: sub}, nil
}

// WatchUpdateDepositLimit is a free log subscription operation binding the contract event 0xae565aab888bca5e19e25a13db7b0c9144305bf55cb0f3f4d724f730e5acdd62.
//
// Solidity: event UpdateDepositLimit(uint256 depositLimit)
func (_Yvault034 *Yvault034Filterer) WatchUpdateDepositLimit(opts *bind.WatchOpts, sink chan<- *Yvault034UpdateDepositLimit) (event.Subscription, error) {

	logs, sub, err := _Yvault034.contract.WatchLogs(opts, "UpdateDepositLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault034UpdateDepositLimit)
				if err := _Yvault034.contract.UnpackLog(event, "UpdateDepositLimit", log); err != nil {
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
func (_Yvault034 *Yvault034Filterer) ParseUpdateDepositLimit(log types.Log) (*Yvault034UpdateDepositLimit, error) {
	event := new(Yvault034UpdateDepositLimit)
	if err := _Yvault034.contract.UnpackLog(event, "UpdateDepositLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault034UpdateGovernanceIterator is returned from FilterUpdateGovernance and is used to iterate over the raw logs and unpacked data for UpdateGovernance events raised by the Yvault034 contract.
type Yvault034UpdateGovernanceIterator struct {
	Event *Yvault034UpdateGovernance // Event containing the contract specifics and raw log

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
func (it *Yvault034UpdateGovernanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault034UpdateGovernance)
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
		it.Event = new(Yvault034UpdateGovernance)
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
func (it *Yvault034UpdateGovernanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault034UpdateGovernanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault034UpdateGovernance represents a UpdateGovernance event raised by the Yvault034 contract.
type Yvault034UpdateGovernance struct {
	Governance common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateGovernance is a free log retrieval operation binding the contract event 0x8d55d160c0009eb3d739442df0a3ca089ed64378bfac017e7ddad463f9815b87.
//
// Solidity: event UpdateGovernance(address governance)
func (_Yvault034 *Yvault034Filterer) FilterUpdateGovernance(opts *bind.FilterOpts) (*Yvault034UpdateGovernanceIterator, error) {

	logs, sub, err := _Yvault034.contract.FilterLogs(opts, "UpdateGovernance")
	if err != nil {
		return nil, err
	}
	return &Yvault034UpdateGovernanceIterator{contract: _Yvault034.contract, event: "UpdateGovernance", logs: logs, sub: sub}, nil
}

// WatchUpdateGovernance is a free log subscription operation binding the contract event 0x8d55d160c0009eb3d739442df0a3ca089ed64378bfac017e7ddad463f9815b87.
//
// Solidity: event UpdateGovernance(address governance)
func (_Yvault034 *Yvault034Filterer) WatchUpdateGovernance(opts *bind.WatchOpts, sink chan<- *Yvault034UpdateGovernance) (event.Subscription, error) {

	logs, sub, err := _Yvault034.contract.WatchLogs(opts, "UpdateGovernance")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault034UpdateGovernance)
				if err := _Yvault034.contract.UnpackLog(event, "UpdateGovernance", log); err != nil {
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
func (_Yvault034 *Yvault034Filterer) ParseUpdateGovernance(log types.Log) (*Yvault034UpdateGovernance, error) {
	event := new(Yvault034UpdateGovernance)
	if err := _Yvault034.contract.UnpackLog(event, "UpdateGovernance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault034UpdateGuardianIterator is returned from FilterUpdateGuardian and is used to iterate over the raw logs and unpacked data for UpdateGuardian events raised by the Yvault034 contract.
type Yvault034UpdateGuardianIterator struct {
	Event *Yvault034UpdateGuardian // Event containing the contract specifics and raw log

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
func (it *Yvault034UpdateGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault034UpdateGuardian)
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
		it.Event = new(Yvault034UpdateGuardian)
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
func (it *Yvault034UpdateGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault034UpdateGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault034UpdateGuardian represents a UpdateGuardian event raised by the Yvault034 contract.
type Yvault034UpdateGuardian struct {
	Guardian common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateGuardian is a free log retrieval operation binding the contract event 0x837b9ad138a0a1839a9637afce5306a5c13e23eb63365686843a5319a243609c.
//
// Solidity: event UpdateGuardian(address guardian)
func (_Yvault034 *Yvault034Filterer) FilterUpdateGuardian(opts *bind.FilterOpts) (*Yvault034UpdateGuardianIterator, error) {

	logs, sub, err := _Yvault034.contract.FilterLogs(opts, "UpdateGuardian")
	if err != nil {
		return nil, err
	}
	return &Yvault034UpdateGuardianIterator{contract: _Yvault034.contract, event: "UpdateGuardian", logs: logs, sub: sub}, nil
}

// WatchUpdateGuardian is a free log subscription operation binding the contract event 0x837b9ad138a0a1839a9637afce5306a5c13e23eb63365686843a5319a243609c.
//
// Solidity: event UpdateGuardian(address guardian)
func (_Yvault034 *Yvault034Filterer) WatchUpdateGuardian(opts *bind.WatchOpts, sink chan<- *Yvault034UpdateGuardian) (event.Subscription, error) {

	logs, sub, err := _Yvault034.contract.WatchLogs(opts, "UpdateGuardian")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault034UpdateGuardian)
				if err := _Yvault034.contract.UnpackLog(event, "UpdateGuardian", log); err != nil {
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
func (_Yvault034 *Yvault034Filterer) ParseUpdateGuardian(log types.Log) (*Yvault034UpdateGuardian, error) {
	event := new(Yvault034UpdateGuardian)
	if err := _Yvault034.contract.UnpackLog(event, "UpdateGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault034UpdateGuestListIterator is returned from FilterUpdateGuestList and is used to iterate over the raw logs and unpacked data for UpdateGuestList events raised by the Yvault034 contract.
type Yvault034UpdateGuestListIterator struct {
	Event *Yvault034UpdateGuestList // Event containing the contract specifics and raw log

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
func (it *Yvault034UpdateGuestListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault034UpdateGuestList)
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
		it.Event = new(Yvault034UpdateGuestList)
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
func (it *Yvault034UpdateGuestListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault034UpdateGuestListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault034UpdateGuestList represents a UpdateGuestList event raised by the Yvault034 contract.
type Yvault034UpdateGuestList struct {
	GuestList common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateGuestList is a free log retrieval operation binding the contract event 0x6d674c311329fb38bbc96dc33d2aad03b9bf9fcfdd8f5e5054fda291a5b3c1f8.
//
// Solidity: event UpdateGuestList(address guestList)
func (_Yvault034 *Yvault034Filterer) FilterUpdateGuestList(opts *bind.FilterOpts) (*Yvault034UpdateGuestListIterator, error) {

	logs, sub, err := _Yvault034.contract.FilterLogs(opts, "UpdateGuestList")
	if err != nil {
		return nil, err
	}
	return &Yvault034UpdateGuestListIterator{contract: _Yvault034.contract, event: "UpdateGuestList", logs: logs, sub: sub}, nil
}

// WatchUpdateGuestList is a free log subscription operation binding the contract event 0x6d674c311329fb38bbc96dc33d2aad03b9bf9fcfdd8f5e5054fda291a5b3c1f8.
//
// Solidity: event UpdateGuestList(address guestList)
func (_Yvault034 *Yvault034Filterer) WatchUpdateGuestList(opts *bind.WatchOpts, sink chan<- *Yvault034UpdateGuestList) (event.Subscription, error) {

	logs, sub, err := _Yvault034.contract.WatchLogs(opts, "UpdateGuestList")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault034UpdateGuestList)
				if err := _Yvault034.contract.UnpackLog(event, "UpdateGuestList", log); err != nil {
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
func (_Yvault034 *Yvault034Filterer) ParseUpdateGuestList(log types.Log) (*Yvault034UpdateGuestList, error) {
	event := new(Yvault034UpdateGuestList)
	if err := _Yvault034.contract.UnpackLog(event, "UpdateGuestList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault034UpdateManagementIterator is returned from FilterUpdateManagement and is used to iterate over the raw logs and unpacked data for UpdateManagement events raised by the Yvault034 contract.
type Yvault034UpdateManagementIterator struct {
	Event *Yvault034UpdateManagement // Event containing the contract specifics and raw log

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
func (it *Yvault034UpdateManagementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault034UpdateManagement)
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
		it.Event = new(Yvault034UpdateManagement)
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
func (it *Yvault034UpdateManagementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault034UpdateManagementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault034UpdateManagement represents a UpdateManagement event raised by the Yvault034 contract.
type Yvault034UpdateManagement struct {
	Management common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateManagement is a free log retrieval operation binding the contract event 0xff54978127edd34aec0f9061fb3b155fbe0ededdfa881ee3e0d541d3a1eef438.
//
// Solidity: event UpdateManagement(address management)
func (_Yvault034 *Yvault034Filterer) FilterUpdateManagement(opts *bind.FilterOpts) (*Yvault034UpdateManagementIterator, error) {

	logs, sub, err := _Yvault034.contract.FilterLogs(opts, "UpdateManagement")
	if err != nil {
		return nil, err
	}
	return &Yvault034UpdateManagementIterator{contract: _Yvault034.contract, event: "UpdateManagement", logs: logs, sub: sub}, nil
}

// WatchUpdateManagement is a free log subscription operation binding the contract event 0xff54978127edd34aec0f9061fb3b155fbe0ededdfa881ee3e0d541d3a1eef438.
//
// Solidity: event UpdateManagement(address management)
func (_Yvault034 *Yvault034Filterer) WatchUpdateManagement(opts *bind.WatchOpts, sink chan<- *Yvault034UpdateManagement) (event.Subscription, error) {

	logs, sub, err := _Yvault034.contract.WatchLogs(opts, "UpdateManagement")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault034UpdateManagement)
				if err := _Yvault034.contract.UnpackLog(event, "UpdateManagement", log); err != nil {
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
func (_Yvault034 *Yvault034Filterer) ParseUpdateManagement(log types.Log) (*Yvault034UpdateManagement, error) {
	event := new(Yvault034UpdateManagement)
	if err := _Yvault034.contract.UnpackLog(event, "UpdateManagement", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault034UpdateManagementFeeIterator is returned from FilterUpdateManagementFee and is used to iterate over the raw logs and unpacked data for UpdateManagementFee events raised by the Yvault034 contract.
type Yvault034UpdateManagementFeeIterator struct {
	Event *Yvault034UpdateManagementFee // Event containing the contract specifics and raw log

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
func (it *Yvault034UpdateManagementFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault034UpdateManagementFee)
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
		it.Event = new(Yvault034UpdateManagementFee)
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
func (it *Yvault034UpdateManagementFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault034UpdateManagementFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault034UpdateManagementFee represents a UpdateManagementFee event raised by the Yvault034 contract.
type Yvault034UpdateManagementFee struct {
	ManagementFee *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdateManagementFee is a free log retrieval operation binding the contract event 0x7a7883b0074f96e2c7fab65eb25abf624c488761a5db889e3bb84855dcc6daaf.
//
// Solidity: event UpdateManagementFee(uint256 managementFee)
func (_Yvault034 *Yvault034Filterer) FilterUpdateManagementFee(opts *bind.FilterOpts) (*Yvault034UpdateManagementFeeIterator, error) {

	logs, sub, err := _Yvault034.contract.FilterLogs(opts, "UpdateManagementFee")
	if err != nil {
		return nil, err
	}
	return &Yvault034UpdateManagementFeeIterator{contract: _Yvault034.contract, event: "UpdateManagementFee", logs: logs, sub: sub}, nil
}

// WatchUpdateManagementFee is a free log subscription operation binding the contract event 0x7a7883b0074f96e2c7fab65eb25abf624c488761a5db889e3bb84855dcc6daaf.
//
// Solidity: event UpdateManagementFee(uint256 managementFee)
func (_Yvault034 *Yvault034Filterer) WatchUpdateManagementFee(opts *bind.WatchOpts, sink chan<- *Yvault034UpdateManagementFee) (event.Subscription, error) {

	logs, sub, err := _Yvault034.contract.WatchLogs(opts, "UpdateManagementFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault034UpdateManagementFee)
				if err := _Yvault034.contract.UnpackLog(event, "UpdateManagementFee", log); err != nil {
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
func (_Yvault034 *Yvault034Filterer) ParseUpdateManagementFee(log types.Log) (*Yvault034UpdateManagementFee, error) {
	event := new(Yvault034UpdateManagementFee)
	if err := _Yvault034.contract.UnpackLog(event, "UpdateManagementFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault034UpdatePerformanceFeeIterator is returned from FilterUpdatePerformanceFee and is used to iterate over the raw logs and unpacked data for UpdatePerformanceFee events raised by the Yvault034 contract.
type Yvault034UpdatePerformanceFeeIterator struct {
	Event *Yvault034UpdatePerformanceFee // Event containing the contract specifics and raw log

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
func (it *Yvault034UpdatePerformanceFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault034UpdatePerformanceFee)
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
		it.Event = new(Yvault034UpdatePerformanceFee)
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
func (it *Yvault034UpdatePerformanceFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault034UpdatePerformanceFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault034UpdatePerformanceFee represents a UpdatePerformanceFee event raised by the Yvault034 contract.
type Yvault034UpdatePerformanceFee struct {
	PerformanceFee *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpdatePerformanceFee is a free log retrieval operation binding the contract event 0x0810a1c261ca2c0cd86a0152c51c43ba9dc329639d2349f98140891b2ea798eb.
//
// Solidity: event UpdatePerformanceFee(uint256 performanceFee)
func (_Yvault034 *Yvault034Filterer) FilterUpdatePerformanceFee(opts *bind.FilterOpts) (*Yvault034UpdatePerformanceFeeIterator, error) {

	logs, sub, err := _Yvault034.contract.FilterLogs(opts, "UpdatePerformanceFee")
	if err != nil {
		return nil, err
	}
	return &Yvault034UpdatePerformanceFeeIterator{contract: _Yvault034.contract, event: "UpdatePerformanceFee", logs: logs, sub: sub}, nil
}

// WatchUpdatePerformanceFee is a free log subscription operation binding the contract event 0x0810a1c261ca2c0cd86a0152c51c43ba9dc329639d2349f98140891b2ea798eb.
//
// Solidity: event UpdatePerformanceFee(uint256 performanceFee)
func (_Yvault034 *Yvault034Filterer) WatchUpdatePerformanceFee(opts *bind.WatchOpts, sink chan<- *Yvault034UpdatePerformanceFee) (event.Subscription, error) {

	logs, sub, err := _Yvault034.contract.WatchLogs(opts, "UpdatePerformanceFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault034UpdatePerformanceFee)
				if err := _Yvault034.contract.UnpackLog(event, "UpdatePerformanceFee", log); err != nil {
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
func (_Yvault034 *Yvault034Filterer) ParseUpdatePerformanceFee(log types.Log) (*Yvault034UpdatePerformanceFee, error) {
	event := new(Yvault034UpdatePerformanceFee)
	if err := _Yvault034.contract.UnpackLog(event, "UpdatePerformanceFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault034UpdateRewardsIterator is returned from FilterUpdateRewards and is used to iterate over the raw logs and unpacked data for UpdateRewards events raised by the Yvault034 contract.
type Yvault034UpdateRewardsIterator struct {
	Event *Yvault034UpdateRewards // Event containing the contract specifics and raw log

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
func (it *Yvault034UpdateRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault034UpdateRewards)
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
		it.Event = new(Yvault034UpdateRewards)
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
func (it *Yvault034UpdateRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault034UpdateRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault034UpdateRewards represents a UpdateRewards event raised by the Yvault034 contract.
type Yvault034UpdateRewards struct {
	Rewards common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateRewards is a free log retrieval operation binding the contract event 0xdf3c41a916aecbf42361a147f8348c242662c3ce20ecef30e826b80642477a3d.
//
// Solidity: event UpdateRewards(address rewards)
func (_Yvault034 *Yvault034Filterer) FilterUpdateRewards(opts *bind.FilterOpts) (*Yvault034UpdateRewardsIterator, error) {

	logs, sub, err := _Yvault034.contract.FilterLogs(opts, "UpdateRewards")
	if err != nil {
		return nil, err
	}
	return &Yvault034UpdateRewardsIterator{contract: _Yvault034.contract, event: "UpdateRewards", logs: logs, sub: sub}, nil
}

// WatchUpdateRewards is a free log subscription operation binding the contract event 0xdf3c41a916aecbf42361a147f8348c242662c3ce20ecef30e826b80642477a3d.
//
// Solidity: event UpdateRewards(address rewards)
func (_Yvault034 *Yvault034Filterer) WatchUpdateRewards(opts *bind.WatchOpts, sink chan<- *Yvault034UpdateRewards) (event.Subscription, error) {

	logs, sub, err := _Yvault034.contract.WatchLogs(opts, "UpdateRewards")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault034UpdateRewards)
				if err := _Yvault034.contract.UnpackLog(event, "UpdateRewards", log); err != nil {
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
func (_Yvault034 *Yvault034Filterer) ParseUpdateRewards(log types.Log) (*Yvault034UpdateRewards, error) {
	event := new(Yvault034UpdateRewards)
	if err := _Yvault034.contract.UnpackLog(event, "UpdateRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault034UpdateWithdrawalQueueIterator is returned from FilterUpdateWithdrawalQueue and is used to iterate over the raw logs and unpacked data for UpdateWithdrawalQueue events raised by the Yvault034 contract.
type Yvault034UpdateWithdrawalQueueIterator struct {
	Event *Yvault034UpdateWithdrawalQueue // Event containing the contract specifics and raw log

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
func (it *Yvault034UpdateWithdrawalQueueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault034UpdateWithdrawalQueue)
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
		it.Event = new(Yvault034UpdateWithdrawalQueue)
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
func (it *Yvault034UpdateWithdrawalQueueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault034UpdateWithdrawalQueueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault034UpdateWithdrawalQueue represents a UpdateWithdrawalQueue event raised by the Yvault034 contract.
type Yvault034UpdateWithdrawalQueue struct {
	Queue [20]common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdateWithdrawalQueue is a free log retrieval operation binding the contract event 0x695ac3ac73f08f2002284ffe563cefe798ee2878a5e04219522e2e99eb89d168.
//
// Solidity: event UpdateWithdrawalQueue(address[20] queue)
func (_Yvault034 *Yvault034Filterer) FilterUpdateWithdrawalQueue(opts *bind.FilterOpts) (*Yvault034UpdateWithdrawalQueueIterator, error) {

	logs, sub, err := _Yvault034.contract.FilterLogs(opts, "UpdateWithdrawalQueue")
	if err != nil {
		return nil, err
	}
	return &Yvault034UpdateWithdrawalQueueIterator{contract: _Yvault034.contract, event: "UpdateWithdrawalQueue", logs: logs, sub: sub}, nil
}

// WatchUpdateWithdrawalQueue is a free log subscription operation binding the contract event 0x695ac3ac73f08f2002284ffe563cefe798ee2878a5e04219522e2e99eb89d168.
//
// Solidity: event UpdateWithdrawalQueue(address[20] queue)
func (_Yvault034 *Yvault034Filterer) WatchUpdateWithdrawalQueue(opts *bind.WatchOpts, sink chan<- *Yvault034UpdateWithdrawalQueue) (event.Subscription, error) {

	logs, sub, err := _Yvault034.contract.WatchLogs(opts, "UpdateWithdrawalQueue")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault034UpdateWithdrawalQueue)
				if err := _Yvault034.contract.UnpackLog(event, "UpdateWithdrawalQueue", log); err != nil {
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
func (_Yvault034 *Yvault034Filterer) ParseUpdateWithdrawalQueue(log types.Log) (*Yvault034UpdateWithdrawalQueue, error) {
	event := new(Yvault034UpdateWithdrawalQueue)
	if err := _Yvault034.contract.UnpackLog(event, "UpdateWithdrawalQueue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
