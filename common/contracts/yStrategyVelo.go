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

// IVelodromeRouterRoutes is an auto generated low-level Go binding around an user-defined struct.
type IVelodromeRouterRoutes struct {
	From    common.Address
	To      common.Address
	Stable  bool
	Factory common.Address
}

// YStrategyVeloMetaData contains all meta data concerning the YStrategyVelo contract.
var YStrategyVeloMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gauge\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIVelodromeRouter.Routes[]\",\"name\":\"_veloSwapRouteForToken0\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIVelodromeRouter.Routes[]\",\"name\":\"_veloSwapRouteForToken1\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"clone\",\"type\":\"address\"}],\"name\":\"Cloned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyExitEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"triggerState\",\"type\":\"bool\"}],\"name\":\"ForcedHarvestTrigger\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"profit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loss\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debtPayment\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debtOutstanding\",\"type\":\"uint256\"}],\"name\":\"Harvested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"SetDoHealthCheck\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"SetHealthCheck\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"baseFeeOracle\",\"type\":\"address\"}],\"name\":\"UpdatedBaseFeeOracle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"creditThreshold\",\"type\":\"uint256\"}],\"name\":\"UpdatedCreditThreshold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newKeeper\",\"type\":\"address\"}],\"name\":\"UpdatedKeeper\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"UpdatedMaxReportDelay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"}],\"name\":\"UpdatedMetadataURI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"UpdatedMinReportDelay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rewards\",\"type\":\"address\"}],\"name\":\"UpdatedRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newStrategist\",\"type\":\"address\"}],\"name\":\"UpdatedStrategist\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"apiVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balanceOfWant\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseFeeOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimableProfitInUsdc\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimableRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_strategist\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewards\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_keeper\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gauge\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIVelodromeRouter.Routes[]\",\"name\":\"_veloSwapRouteForToken0\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIVelodromeRouter.Routes[]\",\"name\":\"_veloSwapRouteForToken1\",\"type\":\"tuple[]\"}],\"name\":\"cloneStrategyVelodrome\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newStrategy\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegatedAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"doHealthCheck\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyExit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"estimatedTotalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ethAmount\",\"type\":\"uint256\"}],\"name\":\"ethToWant\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceHarvestTriggerOnce\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gauge\",\"outputs\":[{\"internalType\":\"contractIVelodromeGauge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"harvest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"harvestProfitMaxInUsdc\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"harvestProfitMinInUsdc\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"callCostinEth\",\"type\":\"uint256\"}],\"name\":\"harvestTrigger\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"healthCheck\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_strategist\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewards\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_keeper\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gauge\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIVelodromeRouter.Routes[]\",\"name\":\"_veloSwapRouteForToken0\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIVelodromeRouter.Routes[]\",\"name\":\"_veloSwapRouteForToken1\",\"type\":\"tuple[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isBaseFeeAcceptable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOriginal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isStablePool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"keeper\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localKeepVELO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"manualRewardClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxReportDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"metadataURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newStrategy\",\"type\":\"address\"}],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minReportDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolToken0\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolToken1\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"contractIVelodromeRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_baseFeeOracle\",\"type\":\"address\"}],\"name\":\"setBaseFeeOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_creditThreshold\",\"type\":\"uint256\"}],\"name\":\"setCreditThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_doHealthCheck\",\"type\":\"bool\"}],\"name\":\"setDoHealthCheck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setEmergencyExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_forceHarvestTriggerOnce\",\"type\":\"bool\"}],\"name\":\"setForceHarvestTriggerOnce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_harvestProfitMinInUsdc\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_harvestProfitMaxInUsdc\",\"type\":\"uint256\"}],\"name\":\"setHarvestTriggerParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_healthCheck\",\"type\":\"address\"}],\"name\":\"setHealthCheck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_keeper\",\"type\":\"address\"}],\"name\":\"setKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_keepVelo\",\"type\":\"uint256\"}],\"name\":\"setLocalKeepVelo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_delay\",\"type\":\"uint256\"}],\"name\":\"setMaxReportDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_metadataURI\",\"type\":\"string\"}],\"name\":\"setMetadataURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_delay\",\"type\":\"uint256\"}],\"name\":\"setMinReportDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewards\",\"type\":\"address\"}],\"name\":\"setRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategist\",\"type\":\"address\"}],\"name\":\"setStrategist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIVelodromeRouter.Routes[]\",\"name\":\"_newSwapRouteForToken0\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIVelodromeRouter.Routes[]\",\"name\":\"_newSwapRouteForToken1\",\"type\":\"tuple[]\"}],\"name\":\"setSwapRoutes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_veloVoter\",\"type\":\"address\"}],\"name\":\"setVoter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"strategist\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"swapRouteForToken0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"swapRouteForToken1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"sweep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"callCostInWei\",\"type\":\"uint256\"}],\"name\":\"tendTrigger\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"contractVaultAPI\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"velo\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"veloRouteToToken0\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"veloRouteToToken1\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"veloVoter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"want\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountNeeded\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_loss\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// YStrategyVeloABI is the input ABI used to generate the binding from.
// Deprecated: Use YStrategyVeloMetaData.ABI instead.
var YStrategyVeloABI = YStrategyVeloMetaData.ABI

// YStrategyVelo is an auto generated Go binding around an Ethereum contract.
type YStrategyVelo struct {
	YStrategyVeloCaller     // Read-only binding to the contract
	YStrategyVeloTransactor // Write-only binding to the contract
	YStrategyVeloFilterer   // Log filterer for contract events
}

// YStrategyVeloCaller is an auto generated read-only Go binding around an Ethereum contract.
type YStrategyVeloCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YStrategyVeloTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YStrategyVeloTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YStrategyVeloFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YStrategyVeloFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YStrategyVeloSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YStrategyVeloSession struct {
	Contract     *YStrategyVelo    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YStrategyVeloCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YStrategyVeloCallerSession struct {
	Contract *YStrategyVeloCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// YStrategyVeloTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YStrategyVeloTransactorSession struct {
	Contract     *YStrategyVeloTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// YStrategyVeloRaw is an auto generated low-level Go binding around an Ethereum contract.
type YStrategyVeloRaw struct {
	Contract *YStrategyVelo // Generic contract binding to access the raw methods on
}

// YStrategyVeloCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YStrategyVeloCallerRaw struct {
	Contract *YStrategyVeloCaller // Generic read-only contract binding to access the raw methods on
}

// YStrategyVeloTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YStrategyVeloTransactorRaw struct {
	Contract *YStrategyVeloTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYStrategyVelo creates a new instance of YStrategyVelo, bound to a specific deployed contract.
func NewYStrategyVelo(address common.Address, backend bind.ContractBackend) (*YStrategyVelo, error) {
	contract, err := bindYStrategyVelo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YStrategyVelo{YStrategyVeloCaller: YStrategyVeloCaller{contract: contract}, YStrategyVeloTransactor: YStrategyVeloTransactor{contract: contract}, YStrategyVeloFilterer: YStrategyVeloFilterer{contract: contract}}, nil
}

// NewYStrategyVeloCaller creates a new read-only instance of YStrategyVelo, bound to a specific deployed contract.
func NewYStrategyVeloCaller(address common.Address, caller bind.ContractCaller) (*YStrategyVeloCaller, error) {
	contract, err := bindYStrategyVelo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YStrategyVeloCaller{contract: contract}, nil
}

// NewYStrategyVeloTransactor creates a new write-only instance of YStrategyVelo, bound to a specific deployed contract.
func NewYStrategyVeloTransactor(address common.Address, transactor bind.ContractTransactor) (*YStrategyVeloTransactor, error) {
	contract, err := bindYStrategyVelo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YStrategyVeloTransactor{contract: contract}, nil
}

// NewYStrategyVeloFilterer creates a new log filterer instance of YStrategyVelo, bound to a specific deployed contract.
func NewYStrategyVeloFilterer(address common.Address, filterer bind.ContractFilterer) (*YStrategyVeloFilterer, error) {
	contract, err := bindYStrategyVelo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YStrategyVeloFilterer{contract: contract}, nil
}

// bindYStrategyVelo binds a generic wrapper to an already deployed contract.
func bindYStrategyVelo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(YStrategyVeloABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YStrategyVelo *YStrategyVeloRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YStrategyVelo.Contract.YStrategyVeloCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YStrategyVelo *YStrategyVeloRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YStrategyVelo.Contract.YStrategyVeloTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YStrategyVelo *YStrategyVeloRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YStrategyVelo.Contract.YStrategyVeloTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YStrategyVelo *YStrategyVeloCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YStrategyVelo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YStrategyVelo *YStrategyVeloTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YStrategyVelo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YStrategyVelo *YStrategyVeloTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YStrategyVelo.Contract.contract.Transact(opts, method, params...)
}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() pure returns(string)
func (_YStrategyVelo *YStrategyVeloCaller) ApiVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _YStrategyVelo.contract.Call(opts, &out, "apiVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() pure returns(string)
func (_YStrategyVelo *YStrategyVeloSession) ApiVersion() (string, error) {
	return _YStrategyVelo.Contract.ApiVersion(&_YStrategyVelo.CallOpts)
}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() pure returns(string)
func (_YStrategyVelo *YStrategyVeloCallerSession) ApiVersion() (string, error) {
	return _YStrategyVelo.Contract.ApiVersion(&_YStrategyVelo.CallOpts)
}

// BalanceOfWant is a free data retrieval call binding the contract method 0xc1a3d44c.
//
// Solidity: function balanceOfWant() view returns(uint256)
func (_YStrategyVelo *YStrategyVeloCaller) BalanceOfWant(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyVelo.contract.Call(opts, &out, "balanceOfWant")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfWant is a free data retrieval call binding the contract method 0xc1a3d44c.
//
// Solidity: function balanceOfWant() view returns(uint256)
func (_YStrategyVelo *YStrategyVeloSession) BalanceOfWant() (*big.Int, error) {
	return _YStrategyVelo.Contract.BalanceOfWant(&_YStrategyVelo.CallOpts)
}

// BalanceOfWant is a free data retrieval call binding the contract method 0xc1a3d44c.
//
// Solidity: function balanceOfWant() view returns(uint256)
func (_YStrategyVelo *YStrategyVeloCallerSession) BalanceOfWant() (*big.Int, error) {
	return _YStrategyVelo.Contract.BalanceOfWant(&_YStrategyVelo.CallOpts)
}

// BaseFeeOracle is a free data retrieval call binding the contract method 0x826cddf6.
//
// Solidity: function baseFeeOracle() view returns(address)
func (_YStrategyVelo *YStrategyVeloCaller) BaseFeeOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YStrategyVelo.contract.Call(opts, &out, "baseFeeOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BaseFeeOracle is a free data retrieval call binding the contract method 0x826cddf6.
//
// Solidity: function baseFeeOracle() view returns(address)
func (_YStrategyVelo *YStrategyVeloSession) BaseFeeOracle() (common.Address, error) {
	return _YStrategyVelo.Contract.BaseFeeOracle(&_YStrategyVelo.CallOpts)
}

// BaseFeeOracle is a free data retrieval call binding the contract method 0x826cddf6.
//
// Solidity: function baseFeeOracle() view returns(address)
func (_YStrategyVelo *YStrategyVeloCallerSession) BaseFeeOracle() (common.Address, error) {
	return _YStrategyVelo.Contract.BaseFeeOracle(&_YStrategyVelo.CallOpts)
}

// ClaimableProfitInUsdc is a free data retrieval call binding the contract method 0xb5762114.
//
// Solidity: function claimableProfitInUsdc() view returns(uint256)
func (_YStrategyVelo *YStrategyVeloCaller) ClaimableProfitInUsdc(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyVelo.contract.Call(opts, &out, "claimableProfitInUsdc")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimableProfitInUsdc is a free data retrieval call binding the contract method 0xb5762114.
//
// Solidity: function claimableProfitInUsdc() view returns(uint256)
func (_YStrategyVelo *YStrategyVeloSession) ClaimableProfitInUsdc() (*big.Int, error) {
	return _YStrategyVelo.Contract.ClaimableProfitInUsdc(&_YStrategyVelo.CallOpts)
}

// ClaimableProfitInUsdc is a free data retrieval call binding the contract method 0xb5762114.
//
// Solidity: function claimableProfitInUsdc() view returns(uint256)
func (_YStrategyVelo *YStrategyVeloCallerSession) ClaimableProfitInUsdc() (*big.Int, error) {
	return _YStrategyVelo.Contract.ClaimableProfitInUsdc(&_YStrategyVelo.CallOpts)
}

// ClaimableRewards is a free data retrieval call binding the contract method 0x6c003a9b.
//
// Solidity: function claimableRewards() view returns(uint256)
func (_YStrategyVelo *YStrategyVeloCaller) ClaimableRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyVelo.contract.Call(opts, &out, "claimableRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimableRewards is a free data retrieval call binding the contract method 0x6c003a9b.
//
// Solidity: function claimableRewards() view returns(uint256)
func (_YStrategyVelo *YStrategyVeloSession) ClaimableRewards() (*big.Int, error) {
	return _YStrategyVelo.Contract.ClaimableRewards(&_YStrategyVelo.CallOpts)
}

// ClaimableRewards is a free data retrieval call binding the contract method 0x6c003a9b.
//
// Solidity: function claimableRewards() view returns(uint256)
func (_YStrategyVelo *YStrategyVeloCallerSession) ClaimableRewards() (*big.Int, error) {
	return _YStrategyVelo.Contract.ClaimableRewards(&_YStrategyVelo.CallOpts)
}

// CreditThreshold is a free data retrieval call binding the contract method 0xaa5480cf.
//
// Solidity: function creditThreshold() view returns(uint256)
func (_YStrategyVelo *YStrategyVeloCaller) CreditThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyVelo.contract.Call(opts, &out, "creditThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreditThreshold is a free data retrieval call binding the contract method 0xaa5480cf.
//
// Solidity: function creditThreshold() view returns(uint256)
func (_YStrategyVelo *YStrategyVeloSession) CreditThreshold() (*big.Int, error) {
	return _YStrategyVelo.Contract.CreditThreshold(&_YStrategyVelo.CallOpts)
}

// CreditThreshold is a free data retrieval call binding the contract method 0xaa5480cf.
//
// Solidity: function creditThreshold() view returns(uint256)
func (_YStrategyVelo *YStrategyVeloCallerSession) CreditThreshold() (*big.Int, error) {
	return _YStrategyVelo.Contract.CreditThreshold(&_YStrategyVelo.CallOpts)
}

// DelegatedAssets is a free data retrieval call binding the contract method 0x8e6350e2.
//
// Solidity: function delegatedAssets() view returns(uint256)
func (_YStrategyVelo *YStrategyVeloCaller) DelegatedAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyVelo.contract.Call(opts, &out, "delegatedAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegatedAssets is a free data retrieval call binding the contract method 0x8e6350e2.
//
// Solidity: function delegatedAssets() view returns(uint256)
func (_YStrategyVelo *YStrategyVeloSession) DelegatedAssets() (*big.Int, error) {
	return _YStrategyVelo.Contract.DelegatedAssets(&_YStrategyVelo.CallOpts)
}

// DelegatedAssets is a free data retrieval call binding the contract method 0x8e6350e2.
//
// Solidity: function delegatedAssets() view returns(uint256)
func (_YStrategyVelo *YStrategyVeloCallerSession) DelegatedAssets() (*big.Int, error) {
	return _YStrategyVelo.Contract.DelegatedAssets(&_YStrategyVelo.CallOpts)
}

// DoHealthCheck is a free data retrieval call binding the contract method 0x6718835f.
//
// Solidity: function doHealthCheck() view returns(bool)
func (_YStrategyVelo *YStrategyVeloCaller) DoHealthCheck(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _YStrategyVelo.contract.Call(opts, &out, "doHealthCheck")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DoHealthCheck is a free data retrieval call binding the contract method 0x6718835f.
//
// Solidity: function doHealthCheck() view returns(bool)
func (_YStrategyVelo *YStrategyVeloSession) DoHealthCheck() (bool, error) {
	return _YStrategyVelo.Contract.DoHealthCheck(&_YStrategyVelo.CallOpts)
}

// DoHealthCheck is a free data retrieval call binding the contract method 0x6718835f.
//
// Solidity: function doHealthCheck() view returns(bool)
func (_YStrategyVelo *YStrategyVeloCallerSession) DoHealthCheck() (bool, error) {
	return _YStrategyVelo.Contract.DoHealthCheck(&_YStrategyVelo.CallOpts)
}

// EmergencyExit is a free data retrieval call binding the contract method 0x5641ec03.
//
// Solidity: function emergencyExit() view returns(bool)
func (_YStrategyVelo *YStrategyVeloCaller) EmergencyExit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _YStrategyVelo.contract.Call(opts, &out, "emergencyExit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EmergencyExit is a free data retrieval call binding the contract method 0x5641ec03.
//
// Solidity: function emergencyExit() view returns(bool)
func (_YStrategyVelo *YStrategyVeloSession) EmergencyExit() (bool, error) {
	return _YStrategyVelo.Contract.EmergencyExit(&_YStrategyVelo.CallOpts)
}

// EmergencyExit is a free data retrieval call binding the contract method 0x5641ec03.
//
// Solidity: function emergencyExit() view returns(bool)
func (_YStrategyVelo *YStrategyVeloCallerSession) EmergencyExit() (bool, error) {
	return _YStrategyVelo.Contract.EmergencyExit(&_YStrategyVelo.CallOpts)
}

// EstimatedTotalAssets is a free data retrieval call binding the contract method 0xefbb5cb0.
//
// Solidity: function estimatedTotalAssets() view returns(uint256)
func (_YStrategyVelo *YStrategyVeloCaller) EstimatedTotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyVelo.contract.Call(opts, &out, "estimatedTotalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimatedTotalAssets is a free data retrieval call binding the contract method 0xefbb5cb0.
//
// Solidity: function estimatedTotalAssets() view returns(uint256)
func (_YStrategyVelo *YStrategyVeloSession) EstimatedTotalAssets() (*big.Int, error) {
	return _YStrategyVelo.Contract.EstimatedTotalAssets(&_YStrategyVelo.CallOpts)
}

// EstimatedTotalAssets is a free data retrieval call binding the contract method 0xefbb5cb0.
//
// Solidity: function estimatedTotalAssets() view returns(uint256)
func (_YStrategyVelo *YStrategyVeloCallerSession) EstimatedTotalAssets() (*big.Int, error) {
	return _YStrategyVelo.Contract.EstimatedTotalAssets(&_YStrategyVelo.CallOpts)
}

// EthToWant is a free data retrieval call binding the contract method 0x780022a0.
//
// Solidity: function ethToWant(uint256 _ethAmount) view returns(uint256)
func (_YStrategyVelo *YStrategyVeloCaller) EthToWant(opts *bind.CallOpts, _ethAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyVelo.contract.Call(opts, &out, "ethToWant", _ethAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EthToWant is a free data retrieval call binding the contract method 0x780022a0.
//
// Solidity: function ethToWant(uint256 _ethAmount) view returns(uint256)
func (_YStrategyVelo *YStrategyVeloSession) EthToWant(_ethAmount *big.Int) (*big.Int, error) {
	return _YStrategyVelo.Contract.EthToWant(&_YStrategyVelo.CallOpts, _ethAmount)
}

// EthToWant is a free data retrieval call binding the contract method 0x780022a0.
//
// Solidity: function ethToWant(uint256 _ethAmount) view returns(uint256)
func (_YStrategyVelo *YStrategyVeloCallerSession) EthToWant(_ethAmount *big.Int) (*big.Int, error) {
	return _YStrategyVelo.Contract.EthToWant(&_YStrategyVelo.CallOpts, _ethAmount)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_YStrategyVelo *YStrategyVeloCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YStrategyVelo.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_YStrategyVelo *YStrategyVeloSession) Factory() (common.Address, error) {
	return _YStrategyVelo.Contract.Factory(&_YStrategyVelo.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_YStrategyVelo *YStrategyVeloCallerSession) Factory() (common.Address, error) {
	return _YStrategyVelo.Contract.Factory(&_YStrategyVelo.CallOpts)
}

// ForceHarvestTriggerOnce is a free data retrieval call binding the contract method 0xa763cf5b.
//
// Solidity: function forceHarvestTriggerOnce() view returns(bool)
func (_YStrategyVelo *YStrategyVeloCaller) ForceHarvestTriggerOnce(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _YStrategyVelo.contract.Call(opts, &out, "forceHarvestTriggerOnce")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ForceHarvestTriggerOnce is a free data retrieval call binding the contract method 0xa763cf5b.
//
// Solidity: function forceHarvestTriggerOnce() view returns(bool)
func (_YStrategyVelo *YStrategyVeloSession) ForceHarvestTriggerOnce() (bool, error) {
	return _YStrategyVelo.Contract.ForceHarvestTriggerOnce(&_YStrategyVelo.CallOpts)
}

// ForceHarvestTriggerOnce is a free data retrieval call binding the contract method 0xa763cf5b.
//
// Solidity: function forceHarvestTriggerOnce() view returns(bool)
func (_YStrategyVelo *YStrategyVeloCallerSession) ForceHarvestTriggerOnce() (bool, error) {
	return _YStrategyVelo.Contract.ForceHarvestTriggerOnce(&_YStrategyVelo.CallOpts)
}

// Gauge is a free data retrieval call binding the contract method 0xa6f19c84.
//
// Solidity: function gauge() view returns(address)
func (_YStrategyVelo *YStrategyVeloCaller) Gauge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YStrategyVelo.contract.Call(opts, &out, "gauge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gauge is a free data retrieval call binding the contract method 0xa6f19c84.
//
// Solidity: function gauge() view returns(address)
func (_YStrategyVelo *YStrategyVeloSession) Gauge() (common.Address, error) {
	return _YStrategyVelo.Contract.Gauge(&_YStrategyVelo.CallOpts)
}

// Gauge is a free data retrieval call binding the contract method 0xa6f19c84.
//
// Solidity: function gauge() view returns(address)
func (_YStrategyVelo *YStrategyVeloCallerSession) Gauge() (common.Address, error) {
	return _YStrategyVelo.Contract.Gauge(&_YStrategyVelo.CallOpts)
}

// HarvestProfitMaxInUsdc is a free data retrieval call binding the contract method 0xfa4e2df9.
//
// Solidity: function harvestProfitMaxInUsdc() view returns(uint256)
func (_YStrategyVelo *YStrategyVeloCaller) HarvestProfitMaxInUsdc(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyVelo.contract.Call(opts, &out, "harvestProfitMaxInUsdc")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HarvestProfitMaxInUsdc is a free data retrieval call binding the contract method 0xfa4e2df9.
//
// Solidity: function harvestProfitMaxInUsdc() view returns(uint256)
func (_YStrategyVelo *YStrategyVeloSession) HarvestProfitMaxInUsdc() (*big.Int, error) {
	return _YStrategyVelo.Contract.HarvestProfitMaxInUsdc(&_YStrategyVelo.CallOpts)
}

// HarvestProfitMaxInUsdc is a free data retrieval call binding the contract method 0xfa4e2df9.
//
// Solidity: function harvestProfitMaxInUsdc() view returns(uint256)
func (_YStrategyVelo *YStrategyVeloCallerSession) HarvestProfitMaxInUsdc() (*big.Int, error) {
	return _YStrategyVelo.Contract.HarvestProfitMaxInUsdc(&_YStrategyVelo.CallOpts)
}

// HarvestProfitMinInUsdc is a free data retrieval call binding the contract method 0x090c4922.
//
// Solidity: function harvestProfitMinInUsdc() view returns(uint256)
func (_YStrategyVelo *YStrategyVeloCaller) HarvestProfitMinInUsdc(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyVelo.contract.Call(opts, &out, "harvestProfitMinInUsdc")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HarvestProfitMinInUsdc is a free data retrieval call binding the contract method 0x090c4922.
//
// Solidity: function harvestProfitMinInUsdc() view returns(uint256)
func (_YStrategyVelo *YStrategyVeloSession) HarvestProfitMinInUsdc() (*big.Int, error) {
	return _YStrategyVelo.Contract.HarvestProfitMinInUsdc(&_YStrategyVelo.CallOpts)
}

// HarvestProfitMinInUsdc is a free data retrieval call binding the contract method 0x090c4922.
//
// Solidity: function harvestProfitMinInUsdc() view returns(uint256)
func (_YStrategyVelo *YStrategyVeloCallerSession) HarvestProfitMinInUsdc() (*big.Int, error) {
	return _YStrategyVelo.Contract.HarvestProfitMinInUsdc(&_YStrategyVelo.CallOpts)
}

// HarvestTrigger is a free data retrieval call binding the contract method 0xed882c2b.
//
// Solidity: function harvestTrigger(uint256 callCostinEth) view returns(bool)
func (_YStrategyVelo *YStrategyVeloCaller) HarvestTrigger(opts *bind.CallOpts, callCostinEth *big.Int) (bool, error) {
	var out []interface{}
	err := _YStrategyVelo.contract.Call(opts, &out, "harvestTrigger", callCostinEth)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HarvestTrigger is a free data retrieval call binding the contract method 0xed882c2b.
//
// Solidity: function harvestTrigger(uint256 callCostinEth) view returns(bool)
func (_YStrategyVelo *YStrategyVeloSession) HarvestTrigger(callCostinEth *big.Int) (bool, error) {
	return _YStrategyVelo.Contract.HarvestTrigger(&_YStrategyVelo.CallOpts, callCostinEth)
}

// HarvestTrigger is a free data retrieval call binding the contract method 0xed882c2b.
//
// Solidity: function harvestTrigger(uint256 callCostinEth) view returns(bool)
func (_YStrategyVelo *YStrategyVeloCallerSession) HarvestTrigger(callCostinEth *big.Int) (bool, error) {
	return _YStrategyVelo.Contract.HarvestTrigger(&_YStrategyVelo.CallOpts, callCostinEth)
}

// HealthCheck is a free data retrieval call binding the contract method 0xb252720b.
//
// Solidity: function healthCheck() view returns(address)
func (_YStrategyVelo *YStrategyVeloCaller) HealthCheck(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YStrategyVelo.contract.Call(opts, &out, "healthCheck")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HealthCheck is a free data retrieval call binding the contract method 0xb252720b.
//
// Solidity: function healthCheck() view returns(address)
func (_YStrategyVelo *YStrategyVeloSession) HealthCheck() (common.Address, error) {
	return _YStrategyVelo.Contract.HealthCheck(&_YStrategyVelo.CallOpts)
}

// HealthCheck is a free data retrieval call binding the contract method 0xb252720b.
//
// Solidity: function healthCheck() view returns(address)
func (_YStrategyVelo *YStrategyVeloCallerSession) HealthCheck() (common.Address, error) {
	return _YStrategyVelo.Contract.HealthCheck(&_YStrategyVelo.CallOpts)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_YStrategyVelo *YStrategyVeloCaller) IsActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _YStrategyVelo.contract.Call(opts, &out, "isActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_YStrategyVelo *YStrategyVeloSession) IsActive() (bool, error) {
	return _YStrategyVelo.Contract.IsActive(&_YStrategyVelo.CallOpts)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_YStrategyVelo *YStrategyVeloCallerSession) IsActive() (bool, error) {
	return _YStrategyVelo.Contract.IsActive(&_YStrategyVelo.CallOpts)
}

// IsBaseFeeAcceptable is a free data retrieval call binding the contract method 0x95326e2d.
//
// Solidity: function isBaseFeeAcceptable() view returns(bool)
func (_YStrategyVelo *YStrategyVeloCaller) IsBaseFeeAcceptable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _YStrategyVelo.contract.Call(opts, &out, "isBaseFeeAcceptable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBaseFeeAcceptable is a free data retrieval call binding the contract method 0x95326e2d.
//
// Solidity: function isBaseFeeAcceptable() view returns(bool)
func (_YStrategyVelo *YStrategyVeloSession) IsBaseFeeAcceptable() (bool, error) {
	return _YStrategyVelo.Contract.IsBaseFeeAcceptable(&_YStrategyVelo.CallOpts)
}

// IsBaseFeeAcceptable is a free data retrieval call binding the contract method 0x95326e2d.
//
// Solidity: function isBaseFeeAcceptable() view returns(bool)
func (_YStrategyVelo *YStrategyVeloCallerSession) IsBaseFeeAcceptable() (bool, error) {
	return _YStrategyVelo.Contract.IsBaseFeeAcceptable(&_YStrategyVelo.CallOpts)
}

// IsOriginal is a free data retrieval call binding the contract method 0x6f392ce7.
//
// Solidity: function isOriginal() view returns(bool)
func (_YStrategyVelo *YStrategyVeloCaller) IsOriginal(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _YStrategyVelo.contract.Call(opts, &out, "isOriginal")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOriginal is a free data retrieval call binding the contract method 0x6f392ce7.
//
// Solidity: function isOriginal() view returns(bool)
func (_YStrategyVelo *YStrategyVeloSession) IsOriginal() (bool, error) {
	return _YStrategyVelo.Contract.IsOriginal(&_YStrategyVelo.CallOpts)
}

// IsOriginal is a free data retrieval call binding the contract method 0x6f392ce7.
//
// Solidity: function isOriginal() view returns(bool)
func (_YStrategyVelo *YStrategyVeloCallerSession) IsOriginal() (bool, error) {
	return _YStrategyVelo.Contract.IsOriginal(&_YStrategyVelo.CallOpts)
}

// IsStablePool is a free data retrieval call binding the contract method 0x5ae29a8c.
//
// Solidity: function isStablePool() view returns(bool)
func (_YStrategyVelo *YStrategyVeloCaller) IsStablePool(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _YStrategyVelo.contract.Call(opts, &out, "isStablePool")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStablePool is a free data retrieval call binding the contract method 0x5ae29a8c.
//
// Solidity: function isStablePool() view returns(bool)
func (_YStrategyVelo *YStrategyVeloSession) IsStablePool() (bool, error) {
	return _YStrategyVelo.Contract.IsStablePool(&_YStrategyVelo.CallOpts)
}

// IsStablePool is a free data retrieval call binding the contract method 0x5ae29a8c.
//
// Solidity: function isStablePool() view returns(bool)
func (_YStrategyVelo *YStrategyVeloCallerSession) IsStablePool() (bool, error) {
	return _YStrategyVelo.Contract.IsStablePool(&_YStrategyVelo.CallOpts)
}

// Keeper is a free data retrieval call binding the contract method 0xaced1661.
//
// Solidity: function keeper() view returns(address)
func (_YStrategyVelo *YStrategyVeloCaller) Keeper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YStrategyVelo.contract.Call(opts, &out, "keeper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Keeper is a free data retrieval call binding the contract method 0xaced1661.
//
// Solidity: function keeper() view returns(address)
func (_YStrategyVelo *YStrategyVeloSession) Keeper() (common.Address, error) {
	return _YStrategyVelo.Contract.Keeper(&_YStrategyVelo.CallOpts)
}

// Keeper is a free data retrieval call binding the contract method 0xaced1661.
//
// Solidity: function keeper() view returns(address)
func (_YStrategyVelo *YStrategyVeloCallerSession) Keeper() (common.Address, error) {
	return _YStrategyVelo.Contract.Keeper(&_YStrategyVelo.CallOpts)
}

// LocalKeepVELO is a free data retrieval call binding the contract method 0xc9d100da.
//
// Solidity: function localKeepVELO() view returns(uint256)
func (_YStrategyVelo *YStrategyVeloCaller) LocalKeepVELO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyVelo.contract.Call(opts, &out, "localKeepVELO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LocalKeepVELO is a free data retrieval call binding the contract method 0xc9d100da.
//
// Solidity: function localKeepVELO() view returns(uint256)
func (_YStrategyVelo *YStrategyVeloSession) LocalKeepVELO() (*big.Int, error) {
	return _YStrategyVelo.Contract.LocalKeepVELO(&_YStrategyVelo.CallOpts)
}

// LocalKeepVELO is a free data retrieval call binding the contract method 0xc9d100da.
//
// Solidity: function localKeepVELO() view returns(uint256)
func (_YStrategyVelo *YStrategyVeloCallerSession) LocalKeepVELO() (*big.Int, error) {
	return _YStrategyVelo.Contract.LocalKeepVELO(&_YStrategyVelo.CallOpts)
}

// MaxReportDelay is a free data retrieval call binding the contract method 0x28b7ccf7.
//
// Solidity: function maxReportDelay() view returns(uint256)
func (_YStrategyVelo *YStrategyVeloCaller) MaxReportDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyVelo.contract.Call(opts, &out, "maxReportDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxReportDelay is a free data retrieval call binding the contract method 0x28b7ccf7.
//
// Solidity: function maxReportDelay() view returns(uint256)
func (_YStrategyVelo *YStrategyVeloSession) MaxReportDelay() (*big.Int, error) {
	return _YStrategyVelo.Contract.MaxReportDelay(&_YStrategyVelo.CallOpts)
}

// MaxReportDelay is a free data retrieval call binding the contract method 0x28b7ccf7.
//
// Solidity: function maxReportDelay() view returns(uint256)
func (_YStrategyVelo *YStrategyVeloCallerSession) MaxReportDelay() (*big.Int, error) {
	return _YStrategyVelo.Contract.MaxReportDelay(&_YStrategyVelo.CallOpts)
}

// MetadataURI is a free data retrieval call binding the contract method 0x03ee438c.
//
// Solidity: function metadataURI() view returns(string)
func (_YStrategyVelo *YStrategyVeloCaller) MetadataURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _YStrategyVelo.contract.Call(opts, &out, "metadataURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MetadataURI is a free data retrieval call binding the contract method 0x03ee438c.
//
// Solidity: function metadataURI() view returns(string)
func (_YStrategyVelo *YStrategyVeloSession) MetadataURI() (string, error) {
	return _YStrategyVelo.Contract.MetadataURI(&_YStrategyVelo.CallOpts)
}

// MetadataURI is a free data retrieval call binding the contract method 0x03ee438c.
//
// Solidity: function metadataURI() view returns(string)
func (_YStrategyVelo *YStrategyVeloCallerSession) MetadataURI() (string, error) {
	return _YStrategyVelo.Contract.MetadataURI(&_YStrategyVelo.CallOpts)
}

// MinReportDelay is a free data retrieval call binding the contract method 0x95e80c50.
//
// Solidity: function minReportDelay() view returns(uint256)
func (_YStrategyVelo *YStrategyVeloCaller) MinReportDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyVelo.contract.Call(opts, &out, "minReportDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinReportDelay is a free data retrieval call binding the contract method 0x95e80c50.
//
// Solidity: function minReportDelay() view returns(uint256)
func (_YStrategyVelo *YStrategyVeloSession) MinReportDelay() (*big.Int, error) {
	return _YStrategyVelo.Contract.MinReportDelay(&_YStrategyVelo.CallOpts)
}

// MinReportDelay is a free data retrieval call binding the contract method 0x95e80c50.
//
// Solidity: function minReportDelay() view returns(uint256)
func (_YStrategyVelo *YStrategyVeloCallerSession) MinReportDelay() (*big.Int, error) {
	return _YStrategyVelo.Contract.MinReportDelay(&_YStrategyVelo.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YStrategyVelo *YStrategyVeloCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _YStrategyVelo.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YStrategyVelo *YStrategyVeloSession) Name() (string, error) {
	return _YStrategyVelo.Contract.Name(&_YStrategyVelo.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YStrategyVelo *YStrategyVeloCallerSession) Name() (string, error) {
	return _YStrategyVelo.Contract.Name(&_YStrategyVelo.CallOpts)
}

// PoolToken0 is a free data retrieval call binding the contract method 0xa73bf6df.
//
// Solidity: function poolToken0() view returns(address)
func (_YStrategyVelo *YStrategyVeloCaller) PoolToken0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YStrategyVelo.contract.Call(opts, &out, "poolToken0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolToken0 is a free data retrieval call binding the contract method 0xa73bf6df.
//
// Solidity: function poolToken0() view returns(address)
func (_YStrategyVelo *YStrategyVeloSession) PoolToken0() (common.Address, error) {
	return _YStrategyVelo.Contract.PoolToken0(&_YStrategyVelo.CallOpts)
}

// PoolToken0 is a free data retrieval call binding the contract method 0xa73bf6df.
//
// Solidity: function poolToken0() view returns(address)
func (_YStrategyVelo *YStrategyVeloCallerSession) PoolToken0() (common.Address, error) {
	return _YStrategyVelo.Contract.PoolToken0(&_YStrategyVelo.CallOpts)
}

// PoolToken1 is a free data retrieval call binding the contract method 0x4758af8a.
//
// Solidity: function poolToken1() view returns(address)
func (_YStrategyVelo *YStrategyVeloCaller) PoolToken1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YStrategyVelo.contract.Call(opts, &out, "poolToken1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolToken1 is a free data retrieval call binding the contract method 0x4758af8a.
//
// Solidity: function poolToken1() view returns(address)
func (_YStrategyVelo *YStrategyVeloSession) PoolToken1() (common.Address, error) {
	return _YStrategyVelo.Contract.PoolToken1(&_YStrategyVelo.CallOpts)
}

// PoolToken1 is a free data retrieval call binding the contract method 0x4758af8a.
//
// Solidity: function poolToken1() view returns(address)
func (_YStrategyVelo *YStrategyVeloCallerSession) PoolToken1() (common.Address, error) {
	return _YStrategyVelo.Contract.PoolToken1(&_YStrategyVelo.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x9ec5a894.
//
// Solidity: function rewards() view returns(address)
func (_YStrategyVelo *YStrategyVeloCaller) Rewards(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YStrategyVelo.contract.Call(opts, &out, "rewards")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0x9ec5a894.
//
// Solidity: function rewards() view returns(address)
func (_YStrategyVelo *YStrategyVeloSession) Rewards() (common.Address, error) {
	return _YStrategyVelo.Contract.Rewards(&_YStrategyVelo.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x9ec5a894.
//
// Solidity: function rewards() view returns(address)
func (_YStrategyVelo *YStrategyVeloCallerSession) Rewards() (common.Address, error) {
	return _YStrategyVelo.Contract.Rewards(&_YStrategyVelo.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_YStrategyVelo *YStrategyVeloCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YStrategyVelo.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_YStrategyVelo *YStrategyVeloSession) Router() (common.Address, error) {
	return _YStrategyVelo.Contract.Router(&_YStrategyVelo.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_YStrategyVelo *YStrategyVeloCallerSession) Router() (common.Address, error) {
	return _YStrategyVelo.Contract.Router(&_YStrategyVelo.CallOpts)
}

// StakedBalance is a free data retrieval call binding the contract method 0x5b9f0016.
//
// Solidity: function stakedBalance() view returns(uint256)
func (_YStrategyVelo *YStrategyVeloCaller) StakedBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyVelo.contract.Call(opts, &out, "stakedBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakedBalance is a free data retrieval call binding the contract method 0x5b9f0016.
//
// Solidity: function stakedBalance() view returns(uint256)
func (_YStrategyVelo *YStrategyVeloSession) StakedBalance() (*big.Int, error) {
	return _YStrategyVelo.Contract.StakedBalance(&_YStrategyVelo.CallOpts)
}

// StakedBalance is a free data retrieval call binding the contract method 0x5b9f0016.
//
// Solidity: function stakedBalance() view returns(uint256)
func (_YStrategyVelo *YStrategyVeloCallerSession) StakedBalance() (*big.Int, error) {
	return _YStrategyVelo.Contract.StakedBalance(&_YStrategyVelo.CallOpts)
}

// Strategist is a free data retrieval call binding the contract method 0x1fe4a686.
//
// Solidity: function strategist() view returns(address)
func (_YStrategyVelo *YStrategyVeloCaller) Strategist(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YStrategyVelo.contract.Call(opts, &out, "strategist")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Strategist is a free data retrieval call binding the contract method 0x1fe4a686.
//
// Solidity: function strategist() view returns(address)
func (_YStrategyVelo *YStrategyVeloSession) Strategist() (common.Address, error) {
	return _YStrategyVelo.Contract.Strategist(&_YStrategyVelo.CallOpts)
}

// Strategist is a free data retrieval call binding the contract method 0x1fe4a686.
//
// Solidity: function strategist() view returns(address)
func (_YStrategyVelo *YStrategyVeloCallerSession) Strategist() (common.Address, error) {
	return _YStrategyVelo.Contract.Strategist(&_YStrategyVelo.CallOpts)
}

// SwapRouteForToken0 is a free data retrieval call binding the contract method 0xc6040a2c.
//
// Solidity: function swapRouteForToken0(uint256 ) view returns(address from, address to, bool stable, address factory)
func (_YStrategyVelo *YStrategyVeloCaller) SwapRouteForToken0(opts *bind.CallOpts, arg0 *big.Int) (struct {
	From    common.Address
	To      common.Address
	Stable  bool
	Factory common.Address
}, error) {
	var out []interface{}
	err := _YStrategyVelo.contract.Call(opts, &out, "swapRouteForToken0", arg0)

	outstruct := new(struct {
		From    common.Address
		To      common.Address
		Stable  bool
		Factory common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.From = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.To = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Stable = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.Factory = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// SwapRouteForToken0 is a free data retrieval call binding the contract method 0xc6040a2c.
//
// Solidity: function swapRouteForToken0(uint256 ) view returns(address from, address to, bool stable, address factory)
func (_YStrategyVelo *YStrategyVeloSession) SwapRouteForToken0(arg0 *big.Int) (struct {
	From    common.Address
	To      common.Address
	Stable  bool
	Factory common.Address
}, error) {
	return _YStrategyVelo.Contract.SwapRouteForToken0(&_YStrategyVelo.CallOpts, arg0)
}

// SwapRouteForToken0 is a free data retrieval call binding the contract method 0xc6040a2c.
//
// Solidity: function swapRouteForToken0(uint256 ) view returns(address from, address to, bool stable, address factory)
func (_YStrategyVelo *YStrategyVeloCallerSession) SwapRouteForToken0(arg0 *big.Int) (struct {
	From    common.Address
	To      common.Address
	Stable  bool
	Factory common.Address
}, error) {
	return _YStrategyVelo.Contract.SwapRouteForToken0(&_YStrategyVelo.CallOpts, arg0)
}

// SwapRouteForToken1 is a free data retrieval call binding the contract method 0xc12f6431.
//
// Solidity: function swapRouteForToken1(uint256 ) view returns(address from, address to, bool stable, address factory)
func (_YStrategyVelo *YStrategyVeloCaller) SwapRouteForToken1(opts *bind.CallOpts, arg0 *big.Int) (struct {
	From    common.Address
	To      common.Address
	Stable  bool
	Factory common.Address
}, error) {
	var out []interface{}
	err := _YStrategyVelo.contract.Call(opts, &out, "swapRouteForToken1", arg0)

	outstruct := new(struct {
		From    common.Address
		To      common.Address
		Stable  bool
		Factory common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.From = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.To = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Stable = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.Factory = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// SwapRouteForToken1 is a free data retrieval call binding the contract method 0xc12f6431.
//
// Solidity: function swapRouteForToken1(uint256 ) view returns(address from, address to, bool stable, address factory)
func (_YStrategyVelo *YStrategyVeloSession) SwapRouteForToken1(arg0 *big.Int) (struct {
	From    common.Address
	To      common.Address
	Stable  bool
	Factory common.Address
}, error) {
	return _YStrategyVelo.Contract.SwapRouteForToken1(&_YStrategyVelo.CallOpts, arg0)
}

// SwapRouteForToken1 is a free data retrieval call binding the contract method 0xc12f6431.
//
// Solidity: function swapRouteForToken1(uint256 ) view returns(address from, address to, bool stable, address factory)
func (_YStrategyVelo *YStrategyVeloCallerSession) SwapRouteForToken1(arg0 *big.Int) (struct {
	From    common.Address
	To      common.Address
	Stable  bool
	Factory common.Address
}, error) {
	return _YStrategyVelo.Contract.SwapRouteForToken1(&_YStrategyVelo.CallOpts, arg0)
}

// TendTrigger is a free data retrieval call binding the contract method 0x650d1880.
//
// Solidity: function tendTrigger(uint256 callCostInWei) view returns(bool)
func (_YStrategyVelo *YStrategyVeloCaller) TendTrigger(opts *bind.CallOpts, callCostInWei *big.Int) (bool, error) {
	var out []interface{}
	err := _YStrategyVelo.contract.Call(opts, &out, "tendTrigger", callCostInWei)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TendTrigger is a free data retrieval call binding the contract method 0x650d1880.
//
// Solidity: function tendTrigger(uint256 callCostInWei) view returns(bool)
func (_YStrategyVelo *YStrategyVeloSession) TendTrigger(callCostInWei *big.Int) (bool, error) {
	return _YStrategyVelo.Contract.TendTrigger(&_YStrategyVelo.CallOpts, callCostInWei)
}

// TendTrigger is a free data retrieval call binding the contract method 0x650d1880.
//
// Solidity: function tendTrigger(uint256 callCostInWei) view returns(bool)
func (_YStrategyVelo *YStrategyVeloCallerSession) TendTrigger(callCostInWei *big.Int) (bool, error) {
	return _YStrategyVelo.Contract.TendTrigger(&_YStrategyVelo.CallOpts, callCostInWei)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_YStrategyVelo *YStrategyVeloCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YStrategyVelo.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_YStrategyVelo *YStrategyVeloSession) Vault() (common.Address, error) {
	return _YStrategyVelo.Contract.Vault(&_YStrategyVelo.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_YStrategyVelo *YStrategyVeloCallerSession) Vault() (common.Address, error) {
	return _YStrategyVelo.Contract.Vault(&_YStrategyVelo.CallOpts)
}

// Velo is a free data retrieval call binding the contract method 0x8c7c53ce.
//
// Solidity: function velo() view returns(address)
func (_YStrategyVelo *YStrategyVeloCaller) Velo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YStrategyVelo.contract.Call(opts, &out, "velo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Velo is a free data retrieval call binding the contract method 0x8c7c53ce.
//
// Solidity: function velo() view returns(address)
func (_YStrategyVelo *YStrategyVeloSession) Velo() (common.Address, error) {
	return _YStrategyVelo.Contract.Velo(&_YStrategyVelo.CallOpts)
}

// Velo is a free data retrieval call binding the contract method 0x8c7c53ce.
//
// Solidity: function velo() view returns(address)
func (_YStrategyVelo *YStrategyVeloCallerSession) Velo() (common.Address, error) {
	return _YStrategyVelo.Contract.Velo(&_YStrategyVelo.CallOpts)
}

// VeloRouteToToken0 is a free data retrieval call binding the contract method 0x143f3dce.
//
// Solidity: function veloRouteToToken0() view returns(address[])
func (_YStrategyVelo *YStrategyVeloCaller) VeloRouteToToken0(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _YStrategyVelo.contract.Call(opts, &out, "veloRouteToToken0")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// VeloRouteToToken0 is a free data retrieval call binding the contract method 0x143f3dce.
//
// Solidity: function veloRouteToToken0() view returns(address[])
func (_YStrategyVelo *YStrategyVeloSession) VeloRouteToToken0() ([]common.Address, error) {
	return _YStrategyVelo.Contract.VeloRouteToToken0(&_YStrategyVelo.CallOpts)
}

// VeloRouteToToken0 is a free data retrieval call binding the contract method 0x143f3dce.
//
// Solidity: function veloRouteToToken0() view returns(address[])
func (_YStrategyVelo *YStrategyVeloCallerSession) VeloRouteToToken0() ([]common.Address, error) {
	return _YStrategyVelo.Contract.VeloRouteToToken0(&_YStrategyVelo.CallOpts)
}

// VeloRouteToToken1 is a free data retrieval call binding the contract method 0xb9fe5c4c.
//
// Solidity: function veloRouteToToken1() view returns(address[])
func (_YStrategyVelo *YStrategyVeloCaller) VeloRouteToToken1(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _YStrategyVelo.contract.Call(opts, &out, "veloRouteToToken1")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// VeloRouteToToken1 is a free data retrieval call binding the contract method 0xb9fe5c4c.
//
// Solidity: function veloRouteToToken1() view returns(address[])
func (_YStrategyVelo *YStrategyVeloSession) VeloRouteToToken1() ([]common.Address, error) {
	return _YStrategyVelo.Contract.VeloRouteToToken1(&_YStrategyVelo.CallOpts)
}

// VeloRouteToToken1 is a free data retrieval call binding the contract method 0xb9fe5c4c.
//
// Solidity: function veloRouteToToken1() view returns(address[])
func (_YStrategyVelo *YStrategyVeloCallerSession) VeloRouteToToken1() ([]common.Address, error) {
	return _YStrategyVelo.Contract.VeloRouteToToken1(&_YStrategyVelo.CallOpts)
}

// VeloVoter is a free data retrieval call binding the contract method 0x5cf6d5ad.
//
// Solidity: function veloVoter() view returns(address)
func (_YStrategyVelo *YStrategyVeloCaller) VeloVoter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YStrategyVelo.contract.Call(opts, &out, "veloVoter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VeloVoter is a free data retrieval call binding the contract method 0x5cf6d5ad.
//
// Solidity: function veloVoter() view returns(address)
func (_YStrategyVelo *YStrategyVeloSession) VeloVoter() (common.Address, error) {
	return _YStrategyVelo.Contract.VeloVoter(&_YStrategyVelo.CallOpts)
}

// VeloVoter is a free data retrieval call binding the contract method 0x5cf6d5ad.
//
// Solidity: function veloVoter() view returns(address)
func (_YStrategyVelo *YStrategyVeloCallerSession) VeloVoter() (common.Address, error) {
	return _YStrategyVelo.Contract.VeloVoter(&_YStrategyVelo.CallOpts)
}

// Want is a free data retrieval call binding the contract method 0x1f1fcd51.
//
// Solidity: function want() view returns(address)
func (_YStrategyVelo *YStrategyVeloCaller) Want(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YStrategyVelo.contract.Call(opts, &out, "want")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Want is a free data retrieval call binding the contract method 0x1f1fcd51.
//
// Solidity: function want() view returns(address)
func (_YStrategyVelo *YStrategyVeloSession) Want() (common.Address, error) {
	return _YStrategyVelo.Contract.Want(&_YStrategyVelo.CallOpts)
}

// Want is a free data retrieval call binding the contract method 0x1f1fcd51.
//
// Solidity: function want() view returns(address)
func (_YStrategyVelo *YStrategyVeloCallerSession) Want() (common.Address, error) {
	return _YStrategyVelo.Contract.Want(&_YStrategyVelo.CallOpts)
}

// CloneStrategyVelodrome is a paid mutator transaction binding the contract method 0x320afc2d.
//
// Solidity: function cloneStrategyVelodrome(address _vault, address _strategist, address _rewards, address _keeper, address _gauge, (address,address,bool,address)[] _veloSwapRouteForToken0, (address,address,bool,address)[] _veloSwapRouteForToken1) returns(address newStrategy)
func (_YStrategyVelo *YStrategyVeloTransactor) CloneStrategyVelodrome(opts *bind.TransactOpts, _vault common.Address, _strategist common.Address, _rewards common.Address, _keeper common.Address, _gauge common.Address, _veloSwapRouteForToken0 []IVelodromeRouterRoutes, _veloSwapRouteForToken1 []IVelodromeRouterRoutes) (*types.Transaction, error) {
	return _YStrategyVelo.contract.Transact(opts, "cloneStrategyVelodrome", _vault, _strategist, _rewards, _keeper, _gauge, _veloSwapRouteForToken0, _veloSwapRouteForToken1)
}

// CloneStrategyVelodrome is a paid mutator transaction binding the contract method 0x320afc2d.
//
// Solidity: function cloneStrategyVelodrome(address _vault, address _strategist, address _rewards, address _keeper, address _gauge, (address,address,bool,address)[] _veloSwapRouteForToken0, (address,address,bool,address)[] _veloSwapRouteForToken1) returns(address newStrategy)
func (_YStrategyVelo *YStrategyVeloSession) CloneStrategyVelodrome(_vault common.Address, _strategist common.Address, _rewards common.Address, _keeper common.Address, _gauge common.Address, _veloSwapRouteForToken0 []IVelodromeRouterRoutes, _veloSwapRouteForToken1 []IVelodromeRouterRoutes) (*types.Transaction, error) {
	return _YStrategyVelo.Contract.CloneStrategyVelodrome(&_YStrategyVelo.TransactOpts, _vault, _strategist, _rewards, _keeper, _gauge, _veloSwapRouteForToken0, _veloSwapRouteForToken1)
}

// CloneStrategyVelodrome is a paid mutator transaction binding the contract method 0x320afc2d.
//
// Solidity: function cloneStrategyVelodrome(address _vault, address _strategist, address _rewards, address _keeper, address _gauge, (address,address,bool,address)[] _veloSwapRouteForToken0, (address,address,bool,address)[] _veloSwapRouteForToken1) returns(address newStrategy)
func (_YStrategyVelo *YStrategyVeloTransactorSession) CloneStrategyVelodrome(_vault common.Address, _strategist common.Address, _rewards common.Address, _keeper common.Address, _gauge common.Address, _veloSwapRouteForToken0 []IVelodromeRouterRoutes, _veloSwapRouteForToken1 []IVelodromeRouterRoutes) (*types.Transaction, error) {
	return _YStrategyVelo.Contract.CloneStrategyVelodrome(&_YStrategyVelo.TransactOpts, _vault, _strategist, _rewards, _keeper, _gauge, _veloSwapRouteForToken0, _veloSwapRouteForToken1)
}

// Harvest is a paid mutator transaction binding the contract method 0x4641257d.
//
// Solidity: function harvest() returns()
func (_YStrategyVelo *YStrategyVeloTransactor) Harvest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YStrategyVelo.contract.Transact(opts, "harvest")
}

// Harvest is a paid mutator transaction binding the contract method 0x4641257d.
//
// Solidity: function harvest() returns()
func (_YStrategyVelo *YStrategyVeloSession) Harvest() (*types.Transaction, error) {
	return _YStrategyVelo.Contract.Harvest(&_YStrategyVelo.TransactOpts)
}

// Harvest is a paid mutator transaction binding the contract method 0x4641257d.
//
// Solidity: function harvest() returns()
func (_YStrategyVelo *YStrategyVeloTransactorSession) Harvest() (*types.Transaction, error) {
	return _YStrategyVelo.Contract.Harvest(&_YStrategyVelo.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xd96d862a.
//
// Solidity: function initialize(address _vault, address _strategist, address _rewards, address _keeper, address _gauge, (address,address,bool,address)[] _veloSwapRouteForToken0, (address,address,bool,address)[] _veloSwapRouteForToken1) returns()
func (_YStrategyVelo *YStrategyVeloTransactor) Initialize(opts *bind.TransactOpts, _vault common.Address, _strategist common.Address, _rewards common.Address, _keeper common.Address, _gauge common.Address, _veloSwapRouteForToken0 []IVelodromeRouterRoutes, _veloSwapRouteForToken1 []IVelodromeRouterRoutes) (*types.Transaction, error) {
	return _YStrategyVelo.contract.Transact(opts, "initialize", _vault, _strategist, _rewards, _keeper, _gauge, _veloSwapRouteForToken0, _veloSwapRouteForToken1)
}

// Initialize is a paid mutator transaction binding the contract method 0xd96d862a.
//
// Solidity: function initialize(address _vault, address _strategist, address _rewards, address _keeper, address _gauge, (address,address,bool,address)[] _veloSwapRouteForToken0, (address,address,bool,address)[] _veloSwapRouteForToken1) returns()
func (_YStrategyVelo *YStrategyVeloSession) Initialize(_vault common.Address, _strategist common.Address, _rewards common.Address, _keeper common.Address, _gauge common.Address, _veloSwapRouteForToken0 []IVelodromeRouterRoutes, _veloSwapRouteForToken1 []IVelodromeRouterRoutes) (*types.Transaction, error) {
	return _YStrategyVelo.Contract.Initialize(&_YStrategyVelo.TransactOpts, _vault, _strategist, _rewards, _keeper, _gauge, _veloSwapRouteForToken0, _veloSwapRouteForToken1)
}

// Initialize is a paid mutator transaction binding the contract method 0xd96d862a.
//
// Solidity: function initialize(address _vault, address _strategist, address _rewards, address _keeper, address _gauge, (address,address,bool,address)[] _veloSwapRouteForToken0, (address,address,bool,address)[] _veloSwapRouteForToken1) returns()
func (_YStrategyVelo *YStrategyVeloTransactorSession) Initialize(_vault common.Address, _strategist common.Address, _rewards common.Address, _keeper common.Address, _gauge common.Address, _veloSwapRouteForToken0 []IVelodromeRouterRoutes, _veloSwapRouteForToken1 []IVelodromeRouterRoutes) (*types.Transaction, error) {
	return _YStrategyVelo.Contract.Initialize(&_YStrategyVelo.TransactOpts, _vault, _strategist, _rewards, _keeper, _gauge, _veloSwapRouteForToken0, _veloSwapRouteForToken1)
}

// ManualRewardClaim is a paid mutator transaction binding the contract method 0xe81c1e79.
//
// Solidity: function manualRewardClaim() returns()
func (_YStrategyVelo *YStrategyVeloTransactor) ManualRewardClaim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YStrategyVelo.contract.Transact(opts, "manualRewardClaim")
}

// ManualRewardClaim is a paid mutator transaction binding the contract method 0xe81c1e79.
//
// Solidity: function manualRewardClaim() returns()
func (_YStrategyVelo *YStrategyVeloSession) ManualRewardClaim() (*types.Transaction, error) {
	return _YStrategyVelo.Contract.ManualRewardClaim(&_YStrategyVelo.TransactOpts)
}

// ManualRewardClaim is a paid mutator transaction binding the contract method 0xe81c1e79.
//
// Solidity: function manualRewardClaim() returns()
func (_YStrategyVelo *YStrategyVeloTransactorSession) ManualRewardClaim() (*types.Transaction, error) {
	return _YStrategyVelo.Contract.ManualRewardClaim(&_YStrategyVelo.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _newStrategy) returns()
func (_YStrategyVelo *YStrategyVeloTransactor) Migrate(opts *bind.TransactOpts, _newStrategy common.Address) (*types.Transaction, error) {
	return _YStrategyVelo.contract.Transact(opts, "migrate", _newStrategy)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _newStrategy) returns()
func (_YStrategyVelo *YStrategyVeloSession) Migrate(_newStrategy common.Address) (*types.Transaction, error) {
	return _YStrategyVelo.Contract.Migrate(&_YStrategyVelo.TransactOpts, _newStrategy)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _newStrategy) returns()
func (_YStrategyVelo *YStrategyVeloTransactorSession) Migrate(_newStrategy common.Address) (*types.Transaction, error) {
	return _YStrategyVelo.Contract.Migrate(&_YStrategyVelo.TransactOpts, _newStrategy)
}

// SetBaseFeeOracle is a paid mutator transaction binding the contract method 0x9f450b5a.
//
// Solidity: function setBaseFeeOracle(address _baseFeeOracle) returns()
func (_YStrategyVelo *YStrategyVeloTransactor) SetBaseFeeOracle(opts *bind.TransactOpts, _baseFeeOracle common.Address) (*types.Transaction, error) {
	return _YStrategyVelo.contract.Transact(opts, "setBaseFeeOracle", _baseFeeOracle)
}

// SetBaseFeeOracle is a paid mutator transaction binding the contract method 0x9f450b5a.
//
// Solidity: function setBaseFeeOracle(address _baseFeeOracle) returns()
func (_YStrategyVelo *YStrategyVeloSession) SetBaseFeeOracle(_baseFeeOracle common.Address) (*types.Transaction, error) {
	return _YStrategyVelo.Contract.SetBaseFeeOracle(&_YStrategyVelo.TransactOpts, _baseFeeOracle)
}

// SetBaseFeeOracle is a paid mutator transaction binding the contract method 0x9f450b5a.
//
// Solidity: function setBaseFeeOracle(address _baseFeeOracle) returns()
func (_YStrategyVelo *YStrategyVeloTransactorSession) SetBaseFeeOracle(_baseFeeOracle common.Address) (*types.Transaction, error) {
	return _YStrategyVelo.Contract.SetBaseFeeOracle(&_YStrategyVelo.TransactOpts, _baseFeeOracle)
}

// SetCreditThreshold is a paid mutator transaction binding the contract method 0xfe2508a6.
//
// Solidity: function setCreditThreshold(uint256 _creditThreshold) returns()
func (_YStrategyVelo *YStrategyVeloTransactor) SetCreditThreshold(opts *bind.TransactOpts, _creditThreshold *big.Int) (*types.Transaction, error) {
	return _YStrategyVelo.contract.Transact(opts, "setCreditThreshold", _creditThreshold)
}

// SetCreditThreshold is a paid mutator transaction binding the contract method 0xfe2508a6.
//
// Solidity: function setCreditThreshold(uint256 _creditThreshold) returns()
func (_YStrategyVelo *YStrategyVeloSession) SetCreditThreshold(_creditThreshold *big.Int) (*types.Transaction, error) {
	return _YStrategyVelo.Contract.SetCreditThreshold(&_YStrategyVelo.TransactOpts, _creditThreshold)
}

// SetCreditThreshold is a paid mutator transaction binding the contract method 0xfe2508a6.
//
// Solidity: function setCreditThreshold(uint256 _creditThreshold) returns()
func (_YStrategyVelo *YStrategyVeloTransactorSession) SetCreditThreshold(_creditThreshold *big.Int) (*types.Transaction, error) {
	return _YStrategyVelo.Contract.SetCreditThreshold(&_YStrategyVelo.TransactOpts, _creditThreshold)
}

// SetDoHealthCheck is a paid mutator transaction binding the contract method 0xac00ff26.
//
// Solidity: function setDoHealthCheck(bool _doHealthCheck) returns()
func (_YStrategyVelo *YStrategyVeloTransactor) SetDoHealthCheck(opts *bind.TransactOpts, _doHealthCheck bool) (*types.Transaction, error) {
	return _YStrategyVelo.contract.Transact(opts, "setDoHealthCheck", _doHealthCheck)
}

// SetDoHealthCheck is a paid mutator transaction binding the contract method 0xac00ff26.
//
// Solidity: function setDoHealthCheck(bool _doHealthCheck) returns()
func (_YStrategyVelo *YStrategyVeloSession) SetDoHealthCheck(_doHealthCheck bool) (*types.Transaction, error) {
	return _YStrategyVelo.Contract.SetDoHealthCheck(&_YStrategyVelo.TransactOpts, _doHealthCheck)
}

// SetDoHealthCheck is a paid mutator transaction binding the contract method 0xac00ff26.
//
// Solidity: function setDoHealthCheck(bool _doHealthCheck) returns()
func (_YStrategyVelo *YStrategyVeloTransactorSession) SetDoHealthCheck(_doHealthCheck bool) (*types.Transaction, error) {
	return _YStrategyVelo.Contract.SetDoHealthCheck(&_YStrategyVelo.TransactOpts, _doHealthCheck)
}

// SetEmergencyExit is a paid mutator transaction binding the contract method 0xfcf2d0ad.
//
// Solidity: function setEmergencyExit() returns()
func (_YStrategyVelo *YStrategyVeloTransactor) SetEmergencyExit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YStrategyVelo.contract.Transact(opts, "setEmergencyExit")
}

// SetEmergencyExit is a paid mutator transaction binding the contract method 0xfcf2d0ad.
//
// Solidity: function setEmergencyExit() returns()
func (_YStrategyVelo *YStrategyVeloSession) SetEmergencyExit() (*types.Transaction, error) {
	return _YStrategyVelo.Contract.SetEmergencyExit(&_YStrategyVelo.TransactOpts)
}

// SetEmergencyExit is a paid mutator transaction binding the contract method 0xfcf2d0ad.
//
// Solidity: function setEmergencyExit() returns()
func (_YStrategyVelo *YStrategyVeloTransactorSession) SetEmergencyExit() (*types.Transaction, error) {
	return _YStrategyVelo.Contract.SetEmergencyExit(&_YStrategyVelo.TransactOpts)
}

// SetForceHarvestTriggerOnce is a paid mutator transaction binding the contract method 0x0ada4dab.
//
// Solidity: function setForceHarvestTriggerOnce(bool _forceHarvestTriggerOnce) returns()
func (_YStrategyVelo *YStrategyVeloTransactor) SetForceHarvestTriggerOnce(opts *bind.TransactOpts, _forceHarvestTriggerOnce bool) (*types.Transaction, error) {
	return _YStrategyVelo.contract.Transact(opts, "setForceHarvestTriggerOnce", _forceHarvestTriggerOnce)
}

// SetForceHarvestTriggerOnce is a paid mutator transaction binding the contract method 0x0ada4dab.
//
// Solidity: function setForceHarvestTriggerOnce(bool _forceHarvestTriggerOnce) returns()
func (_YStrategyVelo *YStrategyVeloSession) SetForceHarvestTriggerOnce(_forceHarvestTriggerOnce bool) (*types.Transaction, error) {
	return _YStrategyVelo.Contract.SetForceHarvestTriggerOnce(&_YStrategyVelo.TransactOpts, _forceHarvestTriggerOnce)
}

// SetForceHarvestTriggerOnce is a paid mutator transaction binding the contract method 0x0ada4dab.
//
// Solidity: function setForceHarvestTriggerOnce(bool _forceHarvestTriggerOnce) returns()
func (_YStrategyVelo *YStrategyVeloTransactorSession) SetForceHarvestTriggerOnce(_forceHarvestTriggerOnce bool) (*types.Transaction, error) {
	return _YStrategyVelo.Contract.SetForceHarvestTriggerOnce(&_YStrategyVelo.TransactOpts, _forceHarvestTriggerOnce)
}

// SetHarvestTriggerParams is a paid mutator transaction binding the contract method 0xee6497f1.
//
// Solidity: function setHarvestTriggerParams(uint256 _harvestProfitMinInUsdc, uint256 _harvestProfitMaxInUsdc) returns()
func (_YStrategyVelo *YStrategyVeloTransactor) SetHarvestTriggerParams(opts *bind.TransactOpts, _harvestProfitMinInUsdc *big.Int, _harvestProfitMaxInUsdc *big.Int) (*types.Transaction, error) {
	return _YStrategyVelo.contract.Transact(opts, "setHarvestTriggerParams", _harvestProfitMinInUsdc, _harvestProfitMaxInUsdc)
}

// SetHarvestTriggerParams is a paid mutator transaction binding the contract method 0xee6497f1.
//
// Solidity: function setHarvestTriggerParams(uint256 _harvestProfitMinInUsdc, uint256 _harvestProfitMaxInUsdc) returns()
func (_YStrategyVelo *YStrategyVeloSession) SetHarvestTriggerParams(_harvestProfitMinInUsdc *big.Int, _harvestProfitMaxInUsdc *big.Int) (*types.Transaction, error) {
	return _YStrategyVelo.Contract.SetHarvestTriggerParams(&_YStrategyVelo.TransactOpts, _harvestProfitMinInUsdc, _harvestProfitMaxInUsdc)
}

// SetHarvestTriggerParams is a paid mutator transaction binding the contract method 0xee6497f1.
//
// Solidity: function setHarvestTriggerParams(uint256 _harvestProfitMinInUsdc, uint256 _harvestProfitMaxInUsdc) returns()
func (_YStrategyVelo *YStrategyVeloTransactorSession) SetHarvestTriggerParams(_harvestProfitMinInUsdc *big.Int, _harvestProfitMaxInUsdc *big.Int) (*types.Transaction, error) {
	return _YStrategyVelo.Contract.SetHarvestTriggerParams(&_YStrategyVelo.TransactOpts, _harvestProfitMinInUsdc, _harvestProfitMaxInUsdc)
}

// SetHealthCheck is a paid mutator transaction binding the contract method 0x11bc8245.
//
// Solidity: function setHealthCheck(address _healthCheck) returns()
func (_YStrategyVelo *YStrategyVeloTransactor) SetHealthCheck(opts *bind.TransactOpts, _healthCheck common.Address) (*types.Transaction, error) {
	return _YStrategyVelo.contract.Transact(opts, "setHealthCheck", _healthCheck)
}

// SetHealthCheck is a paid mutator transaction binding the contract method 0x11bc8245.
//
// Solidity: function setHealthCheck(address _healthCheck) returns()
func (_YStrategyVelo *YStrategyVeloSession) SetHealthCheck(_healthCheck common.Address) (*types.Transaction, error) {
	return _YStrategyVelo.Contract.SetHealthCheck(&_YStrategyVelo.TransactOpts, _healthCheck)
}

// SetHealthCheck is a paid mutator transaction binding the contract method 0x11bc8245.
//
// Solidity: function setHealthCheck(address _healthCheck) returns()
func (_YStrategyVelo *YStrategyVeloTransactorSession) SetHealthCheck(_healthCheck common.Address) (*types.Transaction, error) {
	return _YStrategyVelo.Contract.SetHealthCheck(&_YStrategyVelo.TransactOpts, _healthCheck)
}

// SetKeeper is a paid mutator transaction binding the contract method 0x748747e6.
//
// Solidity: function setKeeper(address _keeper) returns()
func (_YStrategyVelo *YStrategyVeloTransactor) SetKeeper(opts *bind.TransactOpts, _keeper common.Address) (*types.Transaction, error) {
	return _YStrategyVelo.contract.Transact(opts, "setKeeper", _keeper)
}

// SetKeeper is a paid mutator transaction binding the contract method 0x748747e6.
//
// Solidity: function setKeeper(address _keeper) returns()
func (_YStrategyVelo *YStrategyVeloSession) SetKeeper(_keeper common.Address) (*types.Transaction, error) {
	return _YStrategyVelo.Contract.SetKeeper(&_YStrategyVelo.TransactOpts, _keeper)
}

// SetKeeper is a paid mutator transaction binding the contract method 0x748747e6.
//
// Solidity: function setKeeper(address _keeper) returns()
func (_YStrategyVelo *YStrategyVeloTransactorSession) SetKeeper(_keeper common.Address) (*types.Transaction, error) {
	return _YStrategyVelo.Contract.SetKeeper(&_YStrategyVelo.TransactOpts, _keeper)
}

// SetLocalKeepVelo is a paid mutator transaction binding the contract method 0x9a9b23f0.
//
// Solidity: function setLocalKeepVelo(uint256 _keepVelo) returns()
func (_YStrategyVelo *YStrategyVeloTransactor) SetLocalKeepVelo(opts *bind.TransactOpts, _keepVelo *big.Int) (*types.Transaction, error) {
	return _YStrategyVelo.contract.Transact(opts, "setLocalKeepVelo", _keepVelo)
}

// SetLocalKeepVelo is a paid mutator transaction binding the contract method 0x9a9b23f0.
//
// Solidity: function setLocalKeepVelo(uint256 _keepVelo) returns()
func (_YStrategyVelo *YStrategyVeloSession) SetLocalKeepVelo(_keepVelo *big.Int) (*types.Transaction, error) {
	return _YStrategyVelo.Contract.SetLocalKeepVelo(&_YStrategyVelo.TransactOpts, _keepVelo)
}

// SetLocalKeepVelo is a paid mutator transaction binding the contract method 0x9a9b23f0.
//
// Solidity: function setLocalKeepVelo(uint256 _keepVelo) returns()
func (_YStrategyVelo *YStrategyVeloTransactorSession) SetLocalKeepVelo(_keepVelo *big.Int) (*types.Transaction, error) {
	return _YStrategyVelo.Contract.SetLocalKeepVelo(&_YStrategyVelo.TransactOpts, _keepVelo)
}

// SetMaxReportDelay is a paid mutator transaction binding the contract method 0xf017c92f.
//
// Solidity: function setMaxReportDelay(uint256 _delay) returns()
func (_YStrategyVelo *YStrategyVeloTransactor) SetMaxReportDelay(opts *bind.TransactOpts, _delay *big.Int) (*types.Transaction, error) {
	return _YStrategyVelo.contract.Transact(opts, "setMaxReportDelay", _delay)
}

// SetMaxReportDelay is a paid mutator transaction binding the contract method 0xf017c92f.
//
// Solidity: function setMaxReportDelay(uint256 _delay) returns()
func (_YStrategyVelo *YStrategyVeloSession) SetMaxReportDelay(_delay *big.Int) (*types.Transaction, error) {
	return _YStrategyVelo.Contract.SetMaxReportDelay(&_YStrategyVelo.TransactOpts, _delay)
}

// SetMaxReportDelay is a paid mutator transaction binding the contract method 0xf017c92f.
//
// Solidity: function setMaxReportDelay(uint256 _delay) returns()
func (_YStrategyVelo *YStrategyVeloTransactorSession) SetMaxReportDelay(_delay *big.Int) (*types.Transaction, error) {
	return _YStrategyVelo.Contract.SetMaxReportDelay(&_YStrategyVelo.TransactOpts, _delay)
}

// SetMetadataURI is a paid mutator transaction binding the contract method 0x750521f5.
//
// Solidity: function setMetadataURI(string _metadataURI) returns()
func (_YStrategyVelo *YStrategyVeloTransactor) SetMetadataURI(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error) {
	return _YStrategyVelo.contract.Transact(opts, "setMetadataURI", _metadataURI)
}

// SetMetadataURI is a paid mutator transaction binding the contract method 0x750521f5.
//
// Solidity: function setMetadataURI(string _metadataURI) returns()
func (_YStrategyVelo *YStrategyVeloSession) SetMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _YStrategyVelo.Contract.SetMetadataURI(&_YStrategyVelo.TransactOpts, _metadataURI)
}

// SetMetadataURI is a paid mutator transaction binding the contract method 0x750521f5.
//
// Solidity: function setMetadataURI(string _metadataURI) returns()
func (_YStrategyVelo *YStrategyVeloTransactorSession) SetMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _YStrategyVelo.Contract.SetMetadataURI(&_YStrategyVelo.TransactOpts, _metadataURI)
}

// SetMinReportDelay is a paid mutator transaction binding the contract method 0x39a172a8.
//
// Solidity: function setMinReportDelay(uint256 _delay) returns()
func (_YStrategyVelo *YStrategyVeloTransactor) SetMinReportDelay(opts *bind.TransactOpts, _delay *big.Int) (*types.Transaction, error) {
	return _YStrategyVelo.contract.Transact(opts, "setMinReportDelay", _delay)
}

// SetMinReportDelay is a paid mutator transaction binding the contract method 0x39a172a8.
//
// Solidity: function setMinReportDelay(uint256 _delay) returns()
func (_YStrategyVelo *YStrategyVeloSession) SetMinReportDelay(_delay *big.Int) (*types.Transaction, error) {
	return _YStrategyVelo.Contract.SetMinReportDelay(&_YStrategyVelo.TransactOpts, _delay)
}

// SetMinReportDelay is a paid mutator transaction binding the contract method 0x39a172a8.
//
// Solidity: function setMinReportDelay(uint256 _delay) returns()
func (_YStrategyVelo *YStrategyVeloTransactorSession) SetMinReportDelay(_delay *big.Int) (*types.Transaction, error) {
	return _YStrategyVelo.Contract.SetMinReportDelay(&_YStrategyVelo.TransactOpts, _delay)
}

// SetRewards is a paid mutator transaction binding the contract method 0xec38a862.
//
// Solidity: function setRewards(address _rewards) returns()
func (_YStrategyVelo *YStrategyVeloTransactor) SetRewards(opts *bind.TransactOpts, _rewards common.Address) (*types.Transaction, error) {
	return _YStrategyVelo.contract.Transact(opts, "setRewards", _rewards)
}

// SetRewards is a paid mutator transaction binding the contract method 0xec38a862.
//
// Solidity: function setRewards(address _rewards) returns()
func (_YStrategyVelo *YStrategyVeloSession) SetRewards(_rewards common.Address) (*types.Transaction, error) {
	return _YStrategyVelo.Contract.SetRewards(&_YStrategyVelo.TransactOpts, _rewards)
}

// SetRewards is a paid mutator transaction binding the contract method 0xec38a862.
//
// Solidity: function setRewards(address _rewards) returns()
func (_YStrategyVelo *YStrategyVeloTransactorSession) SetRewards(_rewards common.Address) (*types.Transaction, error) {
	return _YStrategyVelo.Contract.SetRewards(&_YStrategyVelo.TransactOpts, _rewards)
}

// SetStrategist is a paid mutator transaction binding the contract method 0xc7b9d530.
//
// Solidity: function setStrategist(address _strategist) returns()
func (_YStrategyVelo *YStrategyVeloTransactor) SetStrategist(opts *bind.TransactOpts, _strategist common.Address) (*types.Transaction, error) {
	return _YStrategyVelo.contract.Transact(opts, "setStrategist", _strategist)
}

// SetStrategist is a paid mutator transaction binding the contract method 0xc7b9d530.
//
// Solidity: function setStrategist(address _strategist) returns()
func (_YStrategyVelo *YStrategyVeloSession) SetStrategist(_strategist common.Address) (*types.Transaction, error) {
	return _YStrategyVelo.Contract.SetStrategist(&_YStrategyVelo.TransactOpts, _strategist)
}

// SetStrategist is a paid mutator transaction binding the contract method 0xc7b9d530.
//
// Solidity: function setStrategist(address _strategist) returns()
func (_YStrategyVelo *YStrategyVeloTransactorSession) SetStrategist(_strategist common.Address) (*types.Transaction, error) {
	return _YStrategyVelo.Contract.SetStrategist(&_YStrategyVelo.TransactOpts, _strategist)
}

// SetSwapRoutes is a paid mutator transaction binding the contract method 0x4a76cbfb.
//
// Solidity: function setSwapRoutes((address,address,bool,address)[] _newSwapRouteForToken0, (address,address,bool,address)[] _newSwapRouteForToken1) returns()
func (_YStrategyVelo *YStrategyVeloTransactor) SetSwapRoutes(opts *bind.TransactOpts, _newSwapRouteForToken0 []IVelodromeRouterRoutes, _newSwapRouteForToken1 []IVelodromeRouterRoutes) (*types.Transaction, error) {
	return _YStrategyVelo.contract.Transact(opts, "setSwapRoutes", _newSwapRouteForToken0, _newSwapRouteForToken1)
}

// SetSwapRoutes is a paid mutator transaction binding the contract method 0x4a76cbfb.
//
// Solidity: function setSwapRoutes((address,address,bool,address)[] _newSwapRouteForToken0, (address,address,bool,address)[] _newSwapRouteForToken1) returns()
func (_YStrategyVelo *YStrategyVeloSession) SetSwapRoutes(_newSwapRouteForToken0 []IVelodromeRouterRoutes, _newSwapRouteForToken1 []IVelodromeRouterRoutes) (*types.Transaction, error) {
	return _YStrategyVelo.Contract.SetSwapRoutes(&_YStrategyVelo.TransactOpts, _newSwapRouteForToken0, _newSwapRouteForToken1)
}

// SetSwapRoutes is a paid mutator transaction binding the contract method 0x4a76cbfb.
//
// Solidity: function setSwapRoutes((address,address,bool,address)[] _newSwapRouteForToken0, (address,address,bool,address)[] _newSwapRouteForToken1) returns()
func (_YStrategyVelo *YStrategyVeloTransactorSession) SetSwapRoutes(_newSwapRouteForToken0 []IVelodromeRouterRoutes, _newSwapRouteForToken1 []IVelodromeRouterRoutes) (*types.Transaction, error) {
	return _YStrategyVelo.Contract.SetSwapRoutes(&_YStrategyVelo.TransactOpts, _newSwapRouteForToken0, _newSwapRouteForToken1)
}

// SetVoter is a paid mutator transaction binding the contract method 0x4bc2a657.
//
// Solidity: function setVoter(address _veloVoter) returns()
func (_YStrategyVelo *YStrategyVeloTransactor) SetVoter(opts *bind.TransactOpts, _veloVoter common.Address) (*types.Transaction, error) {
	return _YStrategyVelo.contract.Transact(opts, "setVoter", _veloVoter)
}

// SetVoter is a paid mutator transaction binding the contract method 0x4bc2a657.
//
// Solidity: function setVoter(address _veloVoter) returns()
func (_YStrategyVelo *YStrategyVeloSession) SetVoter(_veloVoter common.Address) (*types.Transaction, error) {
	return _YStrategyVelo.Contract.SetVoter(&_YStrategyVelo.TransactOpts, _veloVoter)
}

// SetVoter is a paid mutator transaction binding the contract method 0x4bc2a657.
//
// Solidity: function setVoter(address _veloVoter) returns()
func (_YStrategyVelo *YStrategyVeloTransactorSession) SetVoter(_veloVoter common.Address) (*types.Transaction, error) {
	return _YStrategyVelo.Contract.SetVoter(&_YStrategyVelo.TransactOpts, _veloVoter)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address _token) returns()
func (_YStrategyVelo *YStrategyVeloTransactor) Sweep(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _YStrategyVelo.contract.Transact(opts, "sweep", _token)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address _token) returns()
func (_YStrategyVelo *YStrategyVeloSession) Sweep(_token common.Address) (*types.Transaction, error) {
	return _YStrategyVelo.Contract.Sweep(&_YStrategyVelo.TransactOpts, _token)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address _token) returns()
func (_YStrategyVelo *YStrategyVeloTransactorSession) Sweep(_token common.Address) (*types.Transaction, error) {
	return _YStrategyVelo.Contract.Sweep(&_YStrategyVelo.TransactOpts, _token)
}

// Tend is a paid mutator transaction binding the contract method 0x440368a3.
//
// Solidity: function tend() returns()
func (_YStrategyVelo *YStrategyVeloTransactor) Tend(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YStrategyVelo.contract.Transact(opts, "tend")
}

// Tend is a paid mutator transaction binding the contract method 0x440368a3.
//
// Solidity: function tend() returns()
func (_YStrategyVelo *YStrategyVeloSession) Tend() (*types.Transaction, error) {
	return _YStrategyVelo.Contract.Tend(&_YStrategyVelo.TransactOpts)
}

// Tend is a paid mutator transaction binding the contract method 0x440368a3.
//
// Solidity: function tend() returns()
func (_YStrategyVelo *YStrategyVeloTransactorSession) Tend() (*types.Transaction, error) {
	return _YStrategyVelo.Contract.Tend(&_YStrategyVelo.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amountNeeded) returns(uint256 _loss)
func (_YStrategyVelo *YStrategyVeloTransactor) Withdraw(opts *bind.TransactOpts, _amountNeeded *big.Int) (*types.Transaction, error) {
	return _YStrategyVelo.contract.Transact(opts, "withdraw", _amountNeeded)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amountNeeded) returns(uint256 _loss)
func (_YStrategyVelo *YStrategyVeloSession) Withdraw(_amountNeeded *big.Int) (*types.Transaction, error) {
	return _YStrategyVelo.Contract.Withdraw(&_YStrategyVelo.TransactOpts, _amountNeeded)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amountNeeded) returns(uint256 _loss)
func (_YStrategyVelo *YStrategyVeloTransactorSession) Withdraw(_amountNeeded *big.Int) (*types.Transaction, error) {
	return _YStrategyVelo.Contract.Withdraw(&_YStrategyVelo.TransactOpts, _amountNeeded)
}

// YStrategyVeloClonedIterator is returned from FilterCloned and is used to iterate over the raw logs and unpacked data for Cloned events raised by the YStrategyVelo contract.
type YStrategyVeloClonedIterator struct {
	Event *YStrategyVeloCloned // Event containing the contract specifics and raw log

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
func (it *YStrategyVeloClonedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YStrategyVeloCloned)
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
		it.Event = new(YStrategyVeloCloned)
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
func (it *YStrategyVeloClonedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YStrategyVeloClonedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YStrategyVeloCloned represents a Cloned event raised by the YStrategyVelo contract.
type YStrategyVeloCloned struct {
	Clone common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCloned is a free log retrieval operation binding the contract event 0x783540fb4221a3238720dc7038937d0d79982bcf895274aa6ad179f82cf0d53c.
//
// Solidity: event Cloned(address indexed clone)
func (_YStrategyVelo *YStrategyVeloFilterer) FilterCloned(opts *bind.FilterOpts, clone []common.Address) (*YStrategyVeloClonedIterator, error) {

	var cloneRule []interface{}
	for _, cloneItem := range clone {
		cloneRule = append(cloneRule, cloneItem)
	}

	logs, sub, err := _YStrategyVelo.contract.FilterLogs(opts, "Cloned", cloneRule)
	if err != nil {
		return nil, err
	}
	return &YStrategyVeloClonedIterator{contract: _YStrategyVelo.contract, event: "Cloned", logs: logs, sub: sub}, nil
}

// WatchCloned is a free log subscription operation binding the contract event 0x783540fb4221a3238720dc7038937d0d79982bcf895274aa6ad179f82cf0d53c.
//
// Solidity: event Cloned(address indexed clone)
func (_YStrategyVelo *YStrategyVeloFilterer) WatchCloned(opts *bind.WatchOpts, sink chan<- *YStrategyVeloCloned, clone []common.Address) (event.Subscription, error) {

	var cloneRule []interface{}
	for _, cloneItem := range clone {
		cloneRule = append(cloneRule, cloneItem)
	}

	logs, sub, err := _YStrategyVelo.contract.WatchLogs(opts, "Cloned", cloneRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YStrategyVeloCloned)
				if err := _YStrategyVelo.contract.UnpackLog(event, "Cloned", log); err != nil {
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

// ParseCloned is a log parse operation binding the contract event 0x783540fb4221a3238720dc7038937d0d79982bcf895274aa6ad179f82cf0d53c.
//
// Solidity: event Cloned(address indexed clone)
func (_YStrategyVelo *YStrategyVeloFilterer) ParseCloned(log types.Log) (*YStrategyVeloCloned, error) {
	event := new(YStrategyVeloCloned)
	if err := _YStrategyVelo.contract.UnpackLog(event, "Cloned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YStrategyVeloEmergencyExitEnabledIterator is returned from FilterEmergencyExitEnabled and is used to iterate over the raw logs and unpacked data for EmergencyExitEnabled events raised by the YStrategyVelo contract.
type YStrategyVeloEmergencyExitEnabledIterator struct {
	Event *YStrategyVeloEmergencyExitEnabled // Event containing the contract specifics and raw log

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
func (it *YStrategyVeloEmergencyExitEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YStrategyVeloEmergencyExitEnabled)
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
		it.Event = new(YStrategyVeloEmergencyExitEnabled)
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
func (it *YStrategyVeloEmergencyExitEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YStrategyVeloEmergencyExitEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YStrategyVeloEmergencyExitEnabled represents a EmergencyExitEnabled event raised by the YStrategyVelo contract.
type YStrategyVeloEmergencyExitEnabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyExitEnabled is a free log retrieval operation binding the contract event 0x97e963041e952738788b9d4871d854d282065b8f90a464928d6528f2e9a4fd0b.
//
// Solidity: event EmergencyExitEnabled()
func (_YStrategyVelo *YStrategyVeloFilterer) FilterEmergencyExitEnabled(opts *bind.FilterOpts) (*YStrategyVeloEmergencyExitEnabledIterator, error) {

	logs, sub, err := _YStrategyVelo.contract.FilterLogs(opts, "EmergencyExitEnabled")
	if err != nil {
		return nil, err
	}
	return &YStrategyVeloEmergencyExitEnabledIterator{contract: _YStrategyVelo.contract, event: "EmergencyExitEnabled", logs: logs, sub: sub}, nil
}

// WatchEmergencyExitEnabled is a free log subscription operation binding the contract event 0x97e963041e952738788b9d4871d854d282065b8f90a464928d6528f2e9a4fd0b.
//
// Solidity: event EmergencyExitEnabled()
func (_YStrategyVelo *YStrategyVeloFilterer) WatchEmergencyExitEnabled(opts *bind.WatchOpts, sink chan<- *YStrategyVeloEmergencyExitEnabled) (event.Subscription, error) {

	logs, sub, err := _YStrategyVelo.contract.WatchLogs(opts, "EmergencyExitEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YStrategyVeloEmergencyExitEnabled)
				if err := _YStrategyVelo.contract.UnpackLog(event, "EmergencyExitEnabled", log); err != nil {
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

// ParseEmergencyExitEnabled is a log parse operation binding the contract event 0x97e963041e952738788b9d4871d854d282065b8f90a464928d6528f2e9a4fd0b.
//
// Solidity: event EmergencyExitEnabled()
func (_YStrategyVelo *YStrategyVeloFilterer) ParseEmergencyExitEnabled(log types.Log) (*YStrategyVeloEmergencyExitEnabled, error) {
	event := new(YStrategyVeloEmergencyExitEnabled)
	if err := _YStrategyVelo.contract.UnpackLog(event, "EmergencyExitEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YStrategyVeloForcedHarvestTriggerIterator is returned from FilterForcedHarvestTrigger and is used to iterate over the raw logs and unpacked data for ForcedHarvestTrigger events raised by the YStrategyVelo contract.
type YStrategyVeloForcedHarvestTriggerIterator struct {
	Event *YStrategyVeloForcedHarvestTrigger // Event containing the contract specifics and raw log

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
func (it *YStrategyVeloForcedHarvestTriggerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YStrategyVeloForcedHarvestTrigger)
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
		it.Event = new(YStrategyVeloForcedHarvestTrigger)
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
func (it *YStrategyVeloForcedHarvestTriggerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YStrategyVeloForcedHarvestTriggerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YStrategyVeloForcedHarvestTrigger represents a ForcedHarvestTrigger event raised by the YStrategyVelo contract.
type YStrategyVeloForcedHarvestTrigger struct {
	TriggerState bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterForcedHarvestTrigger is a free log retrieval operation binding the contract event 0x6ad28df1b554fa6cacd46ae82fa811748d53798feeb437ddf234bf3083953319.
//
// Solidity: event ForcedHarvestTrigger(bool triggerState)
func (_YStrategyVelo *YStrategyVeloFilterer) FilterForcedHarvestTrigger(opts *bind.FilterOpts) (*YStrategyVeloForcedHarvestTriggerIterator, error) {

	logs, sub, err := _YStrategyVelo.contract.FilterLogs(opts, "ForcedHarvestTrigger")
	if err != nil {
		return nil, err
	}
	return &YStrategyVeloForcedHarvestTriggerIterator{contract: _YStrategyVelo.contract, event: "ForcedHarvestTrigger", logs: logs, sub: sub}, nil
}

// WatchForcedHarvestTrigger is a free log subscription operation binding the contract event 0x6ad28df1b554fa6cacd46ae82fa811748d53798feeb437ddf234bf3083953319.
//
// Solidity: event ForcedHarvestTrigger(bool triggerState)
func (_YStrategyVelo *YStrategyVeloFilterer) WatchForcedHarvestTrigger(opts *bind.WatchOpts, sink chan<- *YStrategyVeloForcedHarvestTrigger) (event.Subscription, error) {

	logs, sub, err := _YStrategyVelo.contract.WatchLogs(opts, "ForcedHarvestTrigger")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YStrategyVeloForcedHarvestTrigger)
				if err := _YStrategyVelo.contract.UnpackLog(event, "ForcedHarvestTrigger", log); err != nil {
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

// ParseForcedHarvestTrigger is a log parse operation binding the contract event 0x6ad28df1b554fa6cacd46ae82fa811748d53798feeb437ddf234bf3083953319.
//
// Solidity: event ForcedHarvestTrigger(bool triggerState)
func (_YStrategyVelo *YStrategyVeloFilterer) ParseForcedHarvestTrigger(log types.Log) (*YStrategyVeloForcedHarvestTrigger, error) {
	event := new(YStrategyVeloForcedHarvestTrigger)
	if err := _YStrategyVelo.contract.UnpackLog(event, "ForcedHarvestTrigger", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YStrategyVeloHarvestedIterator is returned from FilterHarvested and is used to iterate over the raw logs and unpacked data for Harvested events raised by the YStrategyVelo contract.
type YStrategyVeloHarvestedIterator struct {
	Event *YStrategyVeloHarvested // Event containing the contract specifics and raw log

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
func (it *YStrategyVeloHarvestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YStrategyVeloHarvested)
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
		it.Event = new(YStrategyVeloHarvested)
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
func (it *YStrategyVeloHarvestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YStrategyVeloHarvestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YStrategyVeloHarvested represents a Harvested event raised by the YStrategyVelo contract.
type YStrategyVeloHarvested struct {
	Profit          *big.Int
	Loss            *big.Int
	DebtPayment     *big.Int
	DebtOutstanding *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterHarvested is a free log retrieval operation binding the contract event 0x4c0f499ffe6befa0ca7c826b0916cf87bea98de658013e76938489368d60d509.
//
// Solidity: event Harvested(uint256 profit, uint256 loss, uint256 debtPayment, uint256 debtOutstanding)
func (_YStrategyVelo *YStrategyVeloFilterer) FilterHarvested(opts *bind.FilterOpts) (*YStrategyVeloHarvestedIterator, error) {

	logs, sub, err := _YStrategyVelo.contract.FilterLogs(opts, "Harvested")
	if err != nil {
		return nil, err
	}
	return &YStrategyVeloHarvestedIterator{contract: _YStrategyVelo.contract, event: "Harvested", logs: logs, sub: sub}, nil
}

// WatchHarvested is a free log subscription operation binding the contract event 0x4c0f499ffe6befa0ca7c826b0916cf87bea98de658013e76938489368d60d509.
//
// Solidity: event Harvested(uint256 profit, uint256 loss, uint256 debtPayment, uint256 debtOutstanding)
func (_YStrategyVelo *YStrategyVeloFilterer) WatchHarvested(opts *bind.WatchOpts, sink chan<- *YStrategyVeloHarvested) (event.Subscription, error) {

	logs, sub, err := _YStrategyVelo.contract.WatchLogs(opts, "Harvested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YStrategyVeloHarvested)
				if err := _YStrategyVelo.contract.UnpackLog(event, "Harvested", log); err != nil {
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

// ParseHarvested is a log parse operation binding the contract event 0x4c0f499ffe6befa0ca7c826b0916cf87bea98de658013e76938489368d60d509.
//
// Solidity: event Harvested(uint256 profit, uint256 loss, uint256 debtPayment, uint256 debtOutstanding)
func (_YStrategyVelo *YStrategyVeloFilterer) ParseHarvested(log types.Log) (*YStrategyVeloHarvested, error) {
	event := new(YStrategyVeloHarvested)
	if err := _YStrategyVelo.contract.UnpackLog(event, "Harvested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YStrategyVeloSetDoHealthCheckIterator is returned from FilterSetDoHealthCheck and is used to iterate over the raw logs and unpacked data for SetDoHealthCheck events raised by the YStrategyVelo contract.
type YStrategyVeloSetDoHealthCheckIterator struct {
	Event *YStrategyVeloSetDoHealthCheck // Event containing the contract specifics and raw log

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
func (it *YStrategyVeloSetDoHealthCheckIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YStrategyVeloSetDoHealthCheck)
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
		it.Event = new(YStrategyVeloSetDoHealthCheck)
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
func (it *YStrategyVeloSetDoHealthCheckIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YStrategyVeloSetDoHealthCheckIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YStrategyVeloSetDoHealthCheck represents a SetDoHealthCheck event raised by the YStrategyVelo contract.
type YStrategyVeloSetDoHealthCheck struct {
	Arg0 bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetDoHealthCheck is a free log retrieval operation binding the contract event 0xf769f6bf659bbbdabf212d830720ce893eedc57f25ebb8e44edf5b300618a35b.
//
// Solidity: event SetDoHealthCheck(bool arg0)
func (_YStrategyVelo *YStrategyVeloFilterer) FilterSetDoHealthCheck(opts *bind.FilterOpts) (*YStrategyVeloSetDoHealthCheckIterator, error) {

	logs, sub, err := _YStrategyVelo.contract.FilterLogs(opts, "SetDoHealthCheck")
	if err != nil {
		return nil, err
	}
	return &YStrategyVeloSetDoHealthCheckIterator{contract: _YStrategyVelo.contract, event: "SetDoHealthCheck", logs: logs, sub: sub}, nil
}

// WatchSetDoHealthCheck is a free log subscription operation binding the contract event 0xf769f6bf659bbbdabf212d830720ce893eedc57f25ebb8e44edf5b300618a35b.
//
// Solidity: event SetDoHealthCheck(bool arg0)
func (_YStrategyVelo *YStrategyVeloFilterer) WatchSetDoHealthCheck(opts *bind.WatchOpts, sink chan<- *YStrategyVeloSetDoHealthCheck) (event.Subscription, error) {

	logs, sub, err := _YStrategyVelo.contract.WatchLogs(opts, "SetDoHealthCheck")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YStrategyVeloSetDoHealthCheck)
				if err := _YStrategyVelo.contract.UnpackLog(event, "SetDoHealthCheck", log); err != nil {
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

// ParseSetDoHealthCheck is a log parse operation binding the contract event 0xf769f6bf659bbbdabf212d830720ce893eedc57f25ebb8e44edf5b300618a35b.
//
// Solidity: event SetDoHealthCheck(bool arg0)
func (_YStrategyVelo *YStrategyVeloFilterer) ParseSetDoHealthCheck(log types.Log) (*YStrategyVeloSetDoHealthCheck, error) {
	event := new(YStrategyVeloSetDoHealthCheck)
	if err := _YStrategyVelo.contract.UnpackLog(event, "SetDoHealthCheck", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YStrategyVeloSetHealthCheckIterator is returned from FilterSetHealthCheck and is used to iterate over the raw logs and unpacked data for SetHealthCheck events raised by the YStrategyVelo contract.
type YStrategyVeloSetHealthCheckIterator struct {
	Event *YStrategyVeloSetHealthCheck // Event containing the contract specifics and raw log

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
func (it *YStrategyVeloSetHealthCheckIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YStrategyVeloSetHealthCheck)
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
		it.Event = new(YStrategyVeloSetHealthCheck)
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
func (it *YStrategyVeloSetHealthCheckIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YStrategyVeloSetHealthCheckIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YStrategyVeloSetHealthCheck represents a SetHealthCheck event raised by the YStrategyVelo contract.
type YStrategyVeloSetHealthCheck struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetHealthCheck is a free log retrieval operation binding the contract event 0xc8db9c35f716b87af1fbb83f03c78646061931269301fd7ba6dcf189b4cdc2fc.
//
// Solidity: event SetHealthCheck(address arg0)
func (_YStrategyVelo *YStrategyVeloFilterer) FilterSetHealthCheck(opts *bind.FilterOpts) (*YStrategyVeloSetHealthCheckIterator, error) {

	logs, sub, err := _YStrategyVelo.contract.FilterLogs(opts, "SetHealthCheck")
	if err != nil {
		return nil, err
	}
	return &YStrategyVeloSetHealthCheckIterator{contract: _YStrategyVelo.contract, event: "SetHealthCheck", logs: logs, sub: sub}, nil
}

// WatchSetHealthCheck is a free log subscription operation binding the contract event 0xc8db9c35f716b87af1fbb83f03c78646061931269301fd7ba6dcf189b4cdc2fc.
//
// Solidity: event SetHealthCheck(address arg0)
func (_YStrategyVelo *YStrategyVeloFilterer) WatchSetHealthCheck(opts *bind.WatchOpts, sink chan<- *YStrategyVeloSetHealthCheck) (event.Subscription, error) {

	logs, sub, err := _YStrategyVelo.contract.WatchLogs(opts, "SetHealthCheck")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YStrategyVeloSetHealthCheck)
				if err := _YStrategyVelo.contract.UnpackLog(event, "SetHealthCheck", log); err != nil {
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

// ParseSetHealthCheck is a log parse operation binding the contract event 0xc8db9c35f716b87af1fbb83f03c78646061931269301fd7ba6dcf189b4cdc2fc.
//
// Solidity: event SetHealthCheck(address arg0)
func (_YStrategyVelo *YStrategyVeloFilterer) ParseSetHealthCheck(log types.Log) (*YStrategyVeloSetHealthCheck, error) {
	event := new(YStrategyVeloSetHealthCheck)
	if err := _YStrategyVelo.contract.UnpackLog(event, "SetHealthCheck", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YStrategyVeloUpdatedBaseFeeOracleIterator is returned from FilterUpdatedBaseFeeOracle and is used to iterate over the raw logs and unpacked data for UpdatedBaseFeeOracle events raised by the YStrategyVelo contract.
type YStrategyVeloUpdatedBaseFeeOracleIterator struct {
	Event *YStrategyVeloUpdatedBaseFeeOracle // Event containing the contract specifics and raw log

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
func (it *YStrategyVeloUpdatedBaseFeeOracleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YStrategyVeloUpdatedBaseFeeOracle)
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
		it.Event = new(YStrategyVeloUpdatedBaseFeeOracle)
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
func (it *YStrategyVeloUpdatedBaseFeeOracleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YStrategyVeloUpdatedBaseFeeOracleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YStrategyVeloUpdatedBaseFeeOracle represents a UpdatedBaseFeeOracle event raised by the YStrategyVelo contract.
type YStrategyVeloUpdatedBaseFeeOracle struct {
	BaseFeeOracle common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedBaseFeeOracle is a free log retrieval operation binding the contract event 0x711be97287cb9ec921887b9be36e148e1a27c6b158547b22b9704ffc54447a0f.
//
// Solidity: event UpdatedBaseFeeOracle(address baseFeeOracle)
func (_YStrategyVelo *YStrategyVeloFilterer) FilterUpdatedBaseFeeOracle(opts *bind.FilterOpts) (*YStrategyVeloUpdatedBaseFeeOracleIterator, error) {

	logs, sub, err := _YStrategyVelo.contract.FilterLogs(opts, "UpdatedBaseFeeOracle")
	if err != nil {
		return nil, err
	}
	return &YStrategyVeloUpdatedBaseFeeOracleIterator{contract: _YStrategyVelo.contract, event: "UpdatedBaseFeeOracle", logs: logs, sub: sub}, nil
}

// WatchUpdatedBaseFeeOracle is a free log subscription operation binding the contract event 0x711be97287cb9ec921887b9be36e148e1a27c6b158547b22b9704ffc54447a0f.
//
// Solidity: event UpdatedBaseFeeOracle(address baseFeeOracle)
func (_YStrategyVelo *YStrategyVeloFilterer) WatchUpdatedBaseFeeOracle(opts *bind.WatchOpts, sink chan<- *YStrategyVeloUpdatedBaseFeeOracle) (event.Subscription, error) {

	logs, sub, err := _YStrategyVelo.contract.WatchLogs(opts, "UpdatedBaseFeeOracle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YStrategyVeloUpdatedBaseFeeOracle)
				if err := _YStrategyVelo.contract.UnpackLog(event, "UpdatedBaseFeeOracle", log); err != nil {
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

// ParseUpdatedBaseFeeOracle is a log parse operation binding the contract event 0x711be97287cb9ec921887b9be36e148e1a27c6b158547b22b9704ffc54447a0f.
//
// Solidity: event UpdatedBaseFeeOracle(address baseFeeOracle)
func (_YStrategyVelo *YStrategyVeloFilterer) ParseUpdatedBaseFeeOracle(log types.Log) (*YStrategyVeloUpdatedBaseFeeOracle, error) {
	event := new(YStrategyVeloUpdatedBaseFeeOracle)
	if err := _YStrategyVelo.contract.UnpackLog(event, "UpdatedBaseFeeOracle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YStrategyVeloUpdatedCreditThresholdIterator is returned from FilterUpdatedCreditThreshold and is used to iterate over the raw logs and unpacked data for UpdatedCreditThreshold events raised by the YStrategyVelo contract.
type YStrategyVeloUpdatedCreditThresholdIterator struct {
	Event *YStrategyVeloUpdatedCreditThreshold // Event containing the contract specifics and raw log

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
func (it *YStrategyVeloUpdatedCreditThresholdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YStrategyVeloUpdatedCreditThreshold)
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
		it.Event = new(YStrategyVeloUpdatedCreditThreshold)
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
func (it *YStrategyVeloUpdatedCreditThresholdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YStrategyVeloUpdatedCreditThresholdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YStrategyVeloUpdatedCreditThreshold represents a UpdatedCreditThreshold event raised by the YStrategyVelo contract.
type YStrategyVeloUpdatedCreditThreshold struct {
	CreditThreshold *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdatedCreditThreshold is a free log retrieval operation binding the contract event 0xe5ef7832c564a10cbe7b4f1e01ac33a406cb63fcf430a97a9af8616d150af5f3.
//
// Solidity: event UpdatedCreditThreshold(uint256 creditThreshold)
func (_YStrategyVelo *YStrategyVeloFilterer) FilterUpdatedCreditThreshold(opts *bind.FilterOpts) (*YStrategyVeloUpdatedCreditThresholdIterator, error) {

	logs, sub, err := _YStrategyVelo.contract.FilterLogs(opts, "UpdatedCreditThreshold")
	if err != nil {
		return nil, err
	}
	return &YStrategyVeloUpdatedCreditThresholdIterator{contract: _YStrategyVelo.contract, event: "UpdatedCreditThreshold", logs: logs, sub: sub}, nil
}

// WatchUpdatedCreditThreshold is a free log subscription operation binding the contract event 0xe5ef7832c564a10cbe7b4f1e01ac33a406cb63fcf430a97a9af8616d150af5f3.
//
// Solidity: event UpdatedCreditThreshold(uint256 creditThreshold)
func (_YStrategyVelo *YStrategyVeloFilterer) WatchUpdatedCreditThreshold(opts *bind.WatchOpts, sink chan<- *YStrategyVeloUpdatedCreditThreshold) (event.Subscription, error) {

	logs, sub, err := _YStrategyVelo.contract.WatchLogs(opts, "UpdatedCreditThreshold")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YStrategyVeloUpdatedCreditThreshold)
				if err := _YStrategyVelo.contract.UnpackLog(event, "UpdatedCreditThreshold", log); err != nil {
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

// ParseUpdatedCreditThreshold is a log parse operation binding the contract event 0xe5ef7832c564a10cbe7b4f1e01ac33a406cb63fcf430a97a9af8616d150af5f3.
//
// Solidity: event UpdatedCreditThreshold(uint256 creditThreshold)
func (_YStrategyVelo *YStrategyVeloFilterer) ParseUpdatedCreditThreshold(log types.Log) (*YStrategyVeloUpdatedCreditThreshold, error) {
	event := new(YStrategyVeloUpdatedCreditThreshold)
	if err := _YStrategyVelo.contract.UnpackLog(event, "UpdatedCreditThreshold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YStrategyVeloUpdatedKeeperIterator is returned from FilterUpdatedKeeper and is used to iterate over the raw logs and unpacked data for UpdatedKeeper events raised by the YStrategyVelo contract.
type YStrategyVeloUpdatedKeeperIterator struct {
	Event *YStrategyVeloUpdatedKeeper // Event containing the contract specifics and raw log

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
func (it *YStrategyVeloUpdatedKeeperIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YStrategyVeloUpdatedKeeper)
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
		it.Event = new(YStrategyVeloUpdatedKeeper)
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
func (it *YStrategyVeloUpdatedKeeperIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YStrategyVeloUpdatedKeeperIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YStrategyVeloUpdatedKeeper represents a UpdatedKeeper event raised by the YStrategyVelo contract.
type YStrategyVeloUpdatedKeeper struct {
	NewKeeper common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdatedKeeper is a free log retrieval operation binding the contract event 0x2f202ddb4a2e345f6323ed90f8fc8559d770a7abbbeee84dde8aca3351fe7154.
//
// Solidity: event UpdatedKeeper(address newKeeper)
func (_YStrategyVelo *YStrategyVeloFilterer) FilterUpdatedKeeper(opts *bind.FilterOpts) (*YStrategyVeloUpdatedKeeperIterator, error) {

	logs, sub, err := _YStrategyVelo.contract.FilterLogs(opts, "UpdatedKeeper")
	if err != nil {
		return nil, err
	}
	return &YStrategyVeloUpdatedKeeperIterator{contract: _YStrategyVelo.contract, event: "UpdatedKeeper", logs: logs, sub: sub}, nil
}

// WatchUpdatedKeeper is a free log subscription operation binding the contract event 0x2f202ddb4a2e345f6323ed90f8fc8559d770a7abbbeee84dde8aca3351fe7154.
//
// Solidity: event UpdatedKeeper(address newKeeper)
func (_YStrategyVelo *YStrategyVeloFilterer) WatchUpdatedKeeper(opts *bind.WatchOpts, sink chan<- *YStrategyVeloUpdatedKeeper) (event.Subscription, error) {

	logs, sub, err := _YStrategyVelo.contract.WatchLogs(opts, "UpdatedKeeper")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YStrategyVeloUpdatedKeeper)
				if err := _YStrategyVelo.contract.UnpackLog(event, "UpdatedKeeper", log); err != nil {
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

// ParseUpdatedKeeper is a log parse operation binding the contract event 0x2f202ddb4a2e345f6323ed90f8fc8559d770a7abbbeee84dde8aca3351fe7154.
//
// Solidity: event UpdatedKeeper(address newKeeper)
func (_YStrategyVelo *YStrategyVeloFilterer) ParseUpdatedKeeper(log types.Log) (*YStrategyVeloUpdatedKeeper, error) {
	event := new(YStrategyVeloUpdatedKeeper)
	if err := _YStrategyVelo.contract.UnpackLog(event, "UpdatedKeeper", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YStrategyVeloUpdatedMaxReportDelayIterator is returned from FilterUpdatedMaxReportDelay and is used to iterate over the raw logs and unpacked data for UpdatedMaxReportDelay events raised by the YStrategyVelo contract.
type YStrategyVeloUpdatedMaxReportDelayIterator struct {
	Event *YStrategyVeloUpdatedMaxReportDelay // Event containing the contract specifics and raw log

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
func (it *YStrategyVeloUpdatedMaxReportDelayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YStrategyVeloUpdatedMaxReportDelay)
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
		it.Event = new(YStrategyVeloUpdatedMaxReportDelay)
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
func (it *YStrategyVeloUpdatedMaxReportDelayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YStrategyVeloUpdatedMaxReportDelayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YStrategyVeloUpdatedMaxReportDelay represents a UpdatedMaxReportDelay event raised by the YStrategyVelo contract.
type YStrategyVeloUpdatedMaxReportDelay struct {
	Delay *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdatedMaxReportDelay is a free log retrieval operation binding the contract event 0x5430e11864ad7aa9775b07d12657fe52df9aa2ba734355bd8ef8747be2c800c5.
//
// Solidity: event UpdatedMaxReportDelay(uint256 delay)
func (_YStrategyVelo *YStrategyVeloFilterer) FilterUpdatedMaxReportDelay(opts *bind.FilterOpts) (*YStrategyVeloUpdatedMaxReportDelayIterator, error) {

	logs, sub, err := _YStrategyVelo.contract.FilterLogs(opts, "UpdatedMaxReportDelay")
	if err != nil {
		return nil, err
	}
	return &YStrategyVeloUpdatedMaxReportDelayIterator{contract: _YStrategyVelo.contract, event: "UpdatedMaxReportDelay", logs: logs, sub: sub}, nil
}

// WatchUpdatedMaxReportDelay is a free log subscription operation binding the contract event 0x5430e11864ad7aa9775b07d12657fe52df9aa2ba734355bd8ef8747be2c800c5.
//
// Solidity: event UpdatedMaxReportDelay(uint256 delay)
func (_YStrategyVelo *YStrategyVeloFilterer) WatchUpdatedMaxReportDelay(opts *bind.WatchOpts, sink chan<- *YStrategyVeloUpdatedMaxReportDelay) (event.Subscription, error) {

	logs, sub, err := _YStrategyVelo.contract.WatchLogs(opts, "UpdatedMaxReportDelay")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YStrategyVeloUpdatedMaxReportDelay)
				if err := _YStrategyVelo.contract.UnpackLog(event, "UpdatedMaxReportDelay", log); err != nil {
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

// ParseUpdatedMaxReportDelay is a log parse operation binding the contract event 0x5430e11864ad7aa9775b07d12657fe52df9aa2ba734355bd8ef8747be2c800c5.
//
// Solidity: event UpdatedMaxReportDelay(uint256 delay)
func (_YStrategyVelo *YStrategyVeloFilterer) ParseUpdatedMaxReportDelay(log types.Log) (*YStrategyVeloUpdatedMaxReportDelay, error) {
	event := new(YStrategyVeloUpdatedMaxReportDelay)
	if err := _YStrategyVelo.contract.UnpackLog(event, "UpdatedMaxReportDelay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YStrategyVeloUpdatedMetadataURIIterator is returned from FilterUpdatedMetadataURI and is used to iterate over the raw logs and unpacked data for UpdatedMetadataURI events raised by the YStrategyVelo contract.
type YStrategyVeloUpdatedMetadataURIIterator struct {
	Event *YStrategyVeloUpdatedMetadataURI // Event containing the contract specifics and raw log

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
func (it *YStrategyVeloUpdatedMetadataURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YStrategyVeloUpdatedMetadataURI)
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
		it.Event = new(YStrategyVeloUpdatedMetadataURI)
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
func (it *YStrategyVeloUpdatedMetadataURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YStrategyVeloUpdatedMetadataURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YStrategyVeloUpdatedMetadataURI represents a UpdatedMetadataURI event raised by the YStrategyVelo contract.
type YStrategyVeloUpdatedMetadataURI struct {
	MetadataURI string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatedMetadataURI is a free log retrieval operation binding the contract event 0x300e67d5a415b6d015a471d9c7b95dd58f3e8290af965e84e0f845de2996dda6.
//
// Solidity: event UpdatedMetadataURI(string metadataURI)
func (_YStrategyVelo *YStrategyVeloFilterer) FilterUpdatedMetadataURI(opts *bind.FilterOpts) (*YStrategyVeloUpdatedMetadataURIIterator, error) {

	logs, sub, err := _YStrategyVelo.contract.FilterLogs(opts, "UpdatedMetadataURI")
	if err != nil {
		return nil, err
	}
	return &YStrategyVeloUpdatedMetadataURIIterator{contract: _YStrategyVelo.contract, event: "UpdatedMetadataURI", logs: logs, sub: sub}, nil
}

// WatchUpdatedMetadataURI is a free log subscription operation binding the contract event 0x300e67d5a415b6d015a471d9c7b95dd58f3e8290af965e84e0f845de2996dda6.
//
// Solidity: event UpdatedMetadataURI(string metadataURI)
func (_YStrategyVelo *YStrategyVeloFilterer) WatchUpdatedMetadataURI(opts *bind.WatchOpts, sink chan<- *YStrategyVeloUpdatedMetadataURI) (event.Subscription, error) {

	logs, sub, err := _YStrategyVelo.contract.WatchLogs(opts, "UpdatedMetadataURI")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YStrategyVeloUpdatedMetadataURI)
				if err := _YStrategyVelo.contract.UnpackLog(event, "UpdatedMetadataURI", log); err != nil {
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

// ParseUpdatedMetadataURI is a log parse operation binding the contract event 0x300e67d5a415b6d015a471d9c7b95dd58f3e8290af965e84e0f845de2996dda6.
//
// Solidity: event UpdatedMetadataURI(string metadataURI)
func (_YStrategyVelo *YStrategyVeloFilterer) ParseUpdatedMetadataURI(log types.Log) (*YStrategyVeloUpdatedMetadataURI, error) {
	event := new(YStrategyVeloUpdatedMetadataURI)
	if err := _YStrategyVelo.contract.UnpackLog(event, "UpdatedMetadataURI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YStrategyVeloUpdatedMinReportDelayIterator is returned from FilterUpdatedMinReportDelay and is used to iterate over the raw logs and unpacked data for UpdatedMinReportDelay events raised by the YStrategyVelo contract.
type YStrategyVeloUpdatedMinReportDelayIterator struct {
	Event *YStrategyVeloUpdatedMinReportDelay // Event containing the contract specifics and raw log

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
func (it *YStrategyVeloUpdatedMinReportDelayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YStrategyVeloUpdatedMinReportDelay)
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
		it.Event = new(YStrategyVeloUpdatedMinReportDelay)
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
func (it *YStrategyVeloUpdatedMinReportDelayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YStrategyVeloUpdatedMinReportDelayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YStrategyVeloUpdatedMinReportDelay represents a UpdatedMinReportDelay event raised by the YStrategyVelo contract.
type YStrategyVeloUpdatedMinReportDelay struct {
	Delay *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdatedMinReportDelay is a free log retrieval operation binding the contract event 0xbb2c369a0355a34b02ab5fce0643150c87e1c8dfe7c918d465591879f57948b1.
//
// Solidity: event UpdatedMinReportDelay(uint256 delay)
func (_YStrategyVelo *YStrategyVeloFilterer) FilterUpdatedMinReportDelay(opts *bind.FilterOpts) (*YStrategyVeloUpdatedMinReportDelayIterator, error) {

	logs, sub, err := _YStrategyVelo.contract.FilterLogs(opts, "UpdatedMinReportDelay")
	if err != nil {
		return nil, err
	}
	return &YStrategyVeloUpdatedMinReportDelayIterator{contract: _YStrategyVelo.contract, event: "UpdatedMinReportDelay", logs: logs, sub: sub}, nil
}

// WatchUpdatedMinReportDelay is a free log subscription operation binding the contract event 0xbb2c369a0355a34b02ab5fce0643150c87e1c8dfe7c918d465591879f57948b1.
//
// Solidity: event UpdatedMinReportDelay(uint256 delay)
func (_YStrategyVelo *YStrategyVeloFilterer) WatchUpdatedMinReportDelay(opts *bind.WatchOpts, sink chan<- *YStrategyVeloUpdatedMinReportDelay) (event.Subscription, error) {

	logs, sub, err := _YStrategyVelo.contract.WatchLogs(opts, "UpdatedMinReportDelay")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YStrategyVeloUpdatedMinReportDelay)
				if err := _YStrategyVelo.contract.UnpackLog(event, "UpdatedMinReportDelay", log); err != nil {
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

// ParseUpdatedMinReportDelay is a log parse operation binding the contract event 0xbb2c369a0355a34b02ab5fce0643150c87e1c8dfe7c918d465591879f57948b1.
//
// Solidity: event UpdatedMinReportDelay(uint256 delay)
func (_YStrategyVelo *YStrategyVeloFilterer) ParseUpdatedMinReportDelay(log types.Log) (*YStrategyVeloUpdatedMinReportDelay, error) {
	event := new(YStrategyVeloUpdatedMinReportDelay)
	if err := _YStrategyVelo.contract.UnpackLog(event, "UpdatedMinReportDelay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YStrategyVeloUpdatedRewardsIterator is returned from FilterUpdatedRewards and is used to iterate over the raw logs and unpacked data for UpdatedRewards events raised by the YStrategyVelo contract.
type YStrategyVeloUpdatedRewardsIterator struct {
	Event *YStrategyVeloUpdatedRewards // Event containing the contract specifics and raw log

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
func (it *YStrategyVeloUpdatedRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YStrategyVeloUpdatedRewards)
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
		it.Event = new(YStrategyVeloUpdatedRewards)
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
func (it *YStrategyVeloUpdatedRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YStrategyVeloUpdatedRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YStrategyVeloUpdatedRewards represents a UpdatedRewards event raised by the YStrategyVelo contract.
type YStrategyVeloUpdatedRewards struct {
	Rewards common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdatedRewards is a free log retrieval operation binding the contract event 0xafbb66abf8f3b719799940473a4052a3717cdd8e40fb6c8a3faadab316b1a069.
//
// Solidity: event UpdatedRewards(address rewards)
func (_YStrategyVelo *YStrategyVeloFilterer) FilterUpdatedRewards(opts *bind.FilterOpts) (*YStrategyVeloUpdatedRewardsIterator, error) {

	logs, sub, err := _YStrategyVelo.contract.FilterLogs(opts, "UpdatedRewards")
	if err != nil {
		return nil, err
	}
	return &YStrategyVeloUpdatedRewardsIterator{contract: _YStrategyVelo.contract, event: "UpdatedRewards", logs: logs, sub: sub}, nil
}

// WatchUpdatedRewards is a free log subscription operation binding the contract event 0xafbb66abf8f3b719799940473a4052a3717cdd8e40fb6c8a3faadab316b1a069.
//
// Solidity: event UpdatedRewards(address rewards)
func (_YStrategyVelo *YStrategyVeloFilterer) WatchUpdatedRewards(opts *bind.WatchOpts, sink chan<- *YStrategyVeloUpdatedRewards) (event.Subscription, error) {

	logs, sub, err := _YStrategyVelo.contract.WatchLogs(opts, "UpdatedRewards")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YStrategyVeloUpdatedRewards)
				if err := _YStrategyVelo.contract.UnpackLog(event, "UpdatedRewards", log); err != nil {
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

// ParseUpdatedRewards is a log parse operation binding the contract event 0xafbb66abf8f3b719799940473a4052a3717cdd8e40fb6c8a3faadab316b1a069.
//
// Solidity: event UpdatedRewards(address rewards)
func (_YStrategyVelo *YStrategyVeloFilterer) ParseUpdatedRewards(log types.Log) (*YStrategyVeloUpdatedRewards, error) {
	event := new(YStrategyVeloUpdatedRewards)
	if err := _YStrategyVelo.contract.UnpackLog(event, "UpdatedRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YStrategyVeloUpdatedStrategistIterator is returned from FilterUpdatedStrategist and is used to iterate over the raw logs and unpacked data for UpdatedStrategist events raised by the YStrategyVelo contract.
type YStrategyVeloUpdatedStrategistIterator struct {
	Event *YStrategyVeloUpdatedStrategist // Event containing the contract specifics and raw log

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
func (it *YStrategyVeloUpdatedStrategistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YStrategyVeloUpdatedStrategist)
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
		it.Event = new(YStrategyVeloUpdatedStrategist)
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
func (it *YStrategyVeloUpdatedStrategistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YStrategyVeloUpdatedStrategistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YStrategyVeloUpdatedStrategist represents a UpdatedStrategist event raised by the YStrategyVelo contract.
type YStrategyVeloUpdatedStrategist struct {
	NewStrategist common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedStrategist is a free log retrieval operation binding the contract event 0x352ececae6d7d1e6d26bcf2c549dfd55be1637e9b22dc0cf3b71ddb36097a6b4.
//
// Solidity: event UpdatedStrategist(address newStrategist)
func (_YStrategyVelo *YStrategyVeloFilterer) FilterUpdatedStrategist(opts *bind.FilterOpts) (*YStrategyVeloUpdatedStrategistIterator, error) {

	logs, sub, err := _YStrategyVelo.contract.FilterLogs(opts, "UpdatedStrategist")
	if err != nil {
		return nil, err
	}
	return &YStrategyVeloUpdatedStrategistIterator{contract: _YStrategyVelo.contract, event: "UpdatedStrategist", logs: logs, sub: sub}, nil
}

// WatchUpdatedStrategist is a free log subscription operation binding the contract event 0x352ececae6d7d1e6d26bcf2c549dfd55be1637e9b22dc0cf3b71ddb36097a6b4.
//
// Solidity: event UpdatedStrategist(address newStrategist)
func (_YStrategyVelo *YStrategyVeloFilterer) WatchUpdatedStrategist(opts *bind.WatchOpts, sink chan<- *YStrategyVeloUpdatedStrategist) (event.Subscription, error) {

	logs, sub, err := _YStrategyVelo.contract.WatchLogs(opts, "UpdatedStrategist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YStrategyVeloUpdatedStrategist)
				if err := _YStrategyVelo.contract.UnpackLog(event, "UpdatedStrategist", log); err != nil {
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

// ParseUpdatedStrategist is a log parse operation binding the contract event 0x352ececae6d7d1e6d26bcf2c549dfd55be1637e9b22dc0cf3b71ddb36097a6b4.
//
// Solidity: event UpdatedStrategist(address newStrategist)
func (_YStrategyVelo *YStrategyVeloFilterer) ParseUpdatedStrategist(log types.Log) (*YStrategyVeloUpdatedStrategist, error) {
	event := new(YStrategyVeloUpdatedStrategist)
	if err := _YStrategyVelo.contract.UnpackLog(event, "UpdatedStrategist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
