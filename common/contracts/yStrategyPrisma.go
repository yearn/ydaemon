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

// YStrategyPrismaMetaData contains all meta data concerning the YStrategyPrisma contract.
var YStrategyPrismaMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tradeFactory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_harvestProfitMinInUsdc\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_harvestProfitMaxInUsdc\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_prismaVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_prismaReceiver\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"clone\",\"type\":\"address\"}],\"name\":\"Cloned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyExitEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"triggerState\",\"type\":\"bool\"}],\"name\":\"ForcedHarvestTrigger\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"profit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loss\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debtPayment\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debtOutstanding\",\"type\":\"uint256\"}],\"name\":\"Harvested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"SetDoHealthCheck\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"SetHealthCheck\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"baseFeeOracle\",\"type\":\"address\"}],\"name\":\"UpdatedBaseFeeOracle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"creditThreshold\",\"type\":\"uint256\"}],\"name\":\"UpdatedCreditThreshold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newKeeper\",\"type\":\"address\"}],\"name\":\"UpdatedKeeper\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"UpdatedMaxReportDelay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"}],\"name\":\"UpdatedMetadataURI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"UpdatedMinReportDelay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rewards\",\"type\":\"address\"}],\"name\":\"UpdatedRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newStrategist\",\"type\":\"address\"}],\"name\":\"UpdatedStrategist\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"apiVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balanceOfWant\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseFeeOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimableProfitInUsdc\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimsAreMaxBoosted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_strategist\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewards\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_keeper\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tradeFactory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_harvestProfitMinInUsdc\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_harvestProfitMaxInUsdc\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_prismaVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_prismaReceiver\",\"type\":\"address\"}],\"name\":\"cloneStrategyPrismaConvex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newStrategy\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"convexToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"convexVoter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crv\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"curveVoter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegatedAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"doHealthCheck\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyExit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"estimatedTotalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ethAmount\",\"type\":\"uint256\"}],\"name\":\"ethToWant\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceClaimOnce\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceHarvestTriggerOnce\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"harvest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"harvestProfitMaxInUsdc\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"harvestProfitMinInUsdc\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"callCostinEth\",\"type\":\"uint256\"}],\"name\":\"harvestTrigger\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"healthCheck\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_strategist\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewards\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_keeper\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tradeFactory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_harvestProfitMinInUsdc\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_harvestProfitMaxInUsdc\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_prismaVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_prismaReceiver\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isBaseFeeAcceptable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOriginal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"keeper\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localKeepCRV\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localKeepCVX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localKeepYPrisma\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxReportDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"metadataURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newStrategy\",\"type\":\"address\"}],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minReportDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prismaReceiver\",\"outputs\":[{\"internalType\":\"contractIPrismaReceiver\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prismaVault\",\"outputs\":[{\"internalType\":\"contractIPrismaVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_disableTf\",\"type\":\"bool\"}],\"name\":\"removeTradeFactoryPermissions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_baseFeeOracle\",\"type\":\"address\"}],\"name\":\"setBaseFeeOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_creditThreshold\",\"type\":\"uint256\"}],\"name\":\"setCreditThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_doHealthCheck\",\"type\":\"bool\"}],\"name\":\"setDoHealthCheck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setEmergencyExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_forceClaimOnce\",\"type\":\"bool\"}],\"name\":\"setForceClaimOnce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_forceHarvestTriggerOnce\",\"type\":\"bool\"}],\"name\":\"setForceHarvestTriggerOnce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_harvestProfitMinInUsdc\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_harvestProfitMaxInUsdc\",\"type\":\"uint256\"}],\"name\":\"setHarvestTriggerParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_healthCheck\",\"type\":\"address\"}],\"name\":\"setHealthCheck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_keeper\",\"type\":\"address\"}],\"name\":\"setKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_keepCrv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_keepCvx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_keepYPrisma\",\"type\":\"uint256\"}],\"name\":\"setLocalKeepCrvs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_delay\",\"type\":\"uint256\"}],\"name\":\"setMaxReportDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_metadataURI\",\"type\":\"string\"}],\"name\":\"setMetadataURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_delay\",\"type\":\"uint256\"}],\"name\":\"setMinReportDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewards\",\"type\":\"address\"}],\"name\":\"setRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategist\",\"type\":\"address\"}],\"name\":\"setStrategist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_curveVoter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_convexVoter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_yprismaVoter\",\"type\":\"address\"}],\"name\":\"setVoters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"strategist\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"sweep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"callCostInWei\",\"type\":\"uint256\"}],\"name\":\"tendTrigger\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tradeFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newTradeFactory\",\"type\":\"address\"}],\"name\":\"updateTradeFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"contractVaultAPI\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"want\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountNeeded\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_loss\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"yPrisma\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"yprismaVoter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// YStrategyPrismaABI is the input ABI used to generate the binding from.
// Deprecated: Use YStrategyPrismaMetaData.ABI instead.
var YStrategyPrismaABI = YStrategyPrismaMetaData.ABI

// YStrategyPrisma is an auto generated Go binding around an Ethereum contract.
type YStrategyPrisma struct {
	YStrategyPrismaCaller     // Read-only binding to the contract
	YStrategyPrismaTransactor // Write-only binding to the contract
	YStrategyPrismaFilterer   // Log filterer for contract events
}

// YStrategyPrismaCaller is an auto generated read-only Go binding around an Ethereum contract.
type YStrategyPrismaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YStrategyPrismaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YStrategyPrismaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YStrategyPrismaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YStrategyPrismaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YStrategyPrismaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YStrategyPrismaSession struct {
	Contract     *YStrategyPrisma  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YStrategyPrismaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YStrategyPrismaCallerSession struct {
	Contract *YStrategyPrismaCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// YStrategyPrismaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YStrategyPrismaTransactorSession struct {
	Contract     *YStrategyPrismaTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// YStrategyPrismaRaw is an auto generated low-level Go binding around an Ethereum contract.
type YStrategyPrismaRaw struct {
	Contract *YStrategyPrisma // Generic contract binding to access the raw methods on
}

// YStrategyPrismaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YStrategyPrismaCallerRaw struct {
	Contract *YStrategyPrismaCaller // Generic read-only contract binding to access the raw methods on
}

// YStrategyPrismaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YStrategyPrismaTransactorRaw struct {
	Contract *YStrategyPrismaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYStrategyPrisma creates a new instance of YStrategyPrisma, bound to a specific deployed contract.
func NewYStrategyPrisma(address common.Address, backend bind.ContractBackend) (*YStrategyPrisma, error) {
	contract, err := bindYStrategyPrisma(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YStrategyPrisma{YStrategyPrismaCaller: YStrategyPrismaCaller{contract: contract}, YStrategyPrismaTransactor: YStrategyPrismaTransactor{contract: contract}, YStrategyPrismaFilterer: YStrategyPrismaFilterer{contract: contract}}, nil
}

// NewYStrategyPrismaCaller creates a new read-only instance of YStrategyPrisma, bound to a specific deployed contract.
func NewYStrategyPrismaCaller(address common.Address, caller bind.ContractCaller) (*YStrategyPrismaCaller, error) {
	contract, err := bindYStrategyPrisma(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YStrategyPrismaCaller{contract: contract}, nil
}

// NewYStrategyPrismaTransactor creates a new write-only instance of YStrategyPrisma, bound to a specific deployed contract.
func NewYStrategyPrismaTransactor(address common.Address, transactor bind.ContractTransactor) (*YStrategyPrismaTransactor, error) {
	contract, err := bindYStrategyPrisma(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YStrategyPrismaTransactor{contract: contract}, nil
}

// NewYStrategyPrismaFilterer creates a new log filterer instance of YStrategyPrisma, bound to a specific deployed contract.
func NewYStrategyPrismaFilterer(address common.Address, filterer bind.ContractFilterer) (*YStrategyPrismaFilterer, error) {
	contract, err := bindYStrategyPrisma(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YStrategyPrismaFilterer{contract: contract}, nil
}

// bindYStrategyPrisma binds a generic wrapper to an already deployed contract.
func bindYStrategyPrisma(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := YStrategyPrismaMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YStrategyPrisma *YStrategyPrismaRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YStrategyPrisma.Contract.YStrategyPrismaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YStrategyPrisma *YStrategyPrismaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.YStrategyPrismaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YStrategyPrisma *YStrategyPrismaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.YStrategyPrismaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YStrategyPrisma *YStrategyPrismaCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YStrategyPrisma.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YStrategyPrisma *YStrategyPrismaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YStrategyPrisma *YStrategyPrismaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.contract.Transact(opts, method, params...)
}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() pure returns(string)
func (_YStrategyPrisma *YStrategyPrismaCaller) ApiVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _YStrategyPrisma.contract.Call(opts, &out, "apiVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() pure returns(string)
func (_YStrategyPrisma *YStrategyPrismaSession) ApiVersion() (string, error) {
	return _YStrategyPrisma.Contract.ApiVersion(&_YStrategyPrisma.CallOpts)
}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() pure returns(string)
func (_YStrategyPrisma *YStrategyPrismaCallerSession) ApiVersion() (string, error) {
	return _YStrategyPrisma.Contract.ApiVersion(&_YStrategyPrisma.CallOpts)
}

// BalanceOfWant is a free data retrieval call binding the contract method 0xc1a3d44c.
//
// Solidity: function balanceOfWant() view returns(uint256)
func (_YStrategyPrisma *YStrategyPrismaCaller) BalanceOfWant(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyPrisma.contract.Call(opts, &out, "balanceOfWant")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfWant is a free data retrieval call binding the contract method 0xc1a3d44c.
//
// Solidity: function balanceOfWant() view returns(uint256)
func (_YStrategyPrisma *YStrategyPrismaSession) BalanceOfWant() (*big.Int, error) {
	return _YStrategyPrisma.Contract.BalanceOfWant(&_YStrategyPrisma.CallOpts)
}

// BalanceOfWant is a free data retrieval call binding the contract method 0xc1a3d44c.
//
// Solidity: function balanceOfWant() view returns(uint256)
func (_YStrategyPrisma *YStrategyPrismaCallerSession) BalanceOfWant() (*big.Int, error) {
	return _YStrategyPrisma.Contract.BalanceOfWant(&_YStrategyPrisma.CallOpts)
}

// BaseFeeOracle is a free data retrieval call binding the contract method 0x826cddf6.
//
// Solidity: function baseFeeOracle() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaCaller) BaseFeeOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YStrategyPrisma.contract.Call(opts, &out, "baseFeeOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BaseFeeOracle is a free data retrieval call binding the contract method 0x826cddf6.
//
// Solidity: function baseFeeOracle() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaSession) BaseFeeOracle() (common.Address, error) {
	return _YStrategyPrisma.Contract.BaseFeeOracle(&_YStrategyPrisma.CallOpts)
}

// BaseFeeOracle is a free data retrieval call binding the contract method 0x826cddf6.
//
// Solidity: function baseFeeOracle() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaCallerSession) BaseFeeOracle() (common.Address, error) {
	return _YStrategyPrisma.Contract.BaseFeeOracle(&_YStrategyPrisma.CallOpts)
}

// ClaimableProfitInUsdc is a free data retrieval call binding the contract method 0xb5762114.
//
// Solidity: function claimableProfitInUsdc() view returns(uint256)
func (_YStrategyPrisma *YStrategyPrismaCaller) ClaimableProfitInUsdc(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyPrisma.contract.Call(opts, &out, "claimableProfitInUsdc")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimableProfitInUsdc is a free data retrieval call binding the contract method 0xb5762114.
//
// Solidity: function claimableProfitInUsdc() view returns(uint256)
func (_YStrategyPrisma *YStrategyPrismaSession) ClaimableProfitInUsdc() (*big.Int, error) {
	return _YStrategyPrisma.Contract.ClaimableProfitInUsdc(&_YStrategyPrisma.CallOpts)
}

// ClaimableProfitInUsdc is a free data retrieval call binding the contract method 0xb5762114.
//
// Solidity: function claimableProfitInUsdc() view returns(uint256)
func (_YStrategyPrisma *YStrategyPrismaCallerSession) ClaimableProfitInUsdc() (*big.Int, error) {
	return _YStrategyPrisma.Contract.ClaimableProfitInUsdc(&_YStrategyPrisma.CallOpts)
}

// ClaimsAreMaxBoosted is a free data retrieval call binding the contract method 0xe3a937fe.
//
// Solidity: function claimsAreMaxBoosted() view returns(bool)
func (_YStrategyPrisma *YStrategyPrismaCaller) ClaimsAreMaxBoosted(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _YStrategyPrisma.contract.Call(opts, &out, "claimsAreMaxBoosted")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ClaimsAreMaxBoosted is a free data retrieval call binding the contract method 0xe3a937fe.
//
// Solidity: function claimsAreMaxBoosted() view returns(bool)
func (_YStrategyPrisma *YStrategyPrismaSession) ClaimsAreMaxBoosted() (bool, error) {
	return _YStrategyPrisma.Contract.ClaimsAreMaxBoosted(&_YStrategyPrisma.CallOpts)
}

// ClaimsAreMaxBoosted is a free data retrieval call binding the contract method 0xe3a937fe.
//
// Solidity: function claimsAreMaxBoosted() view returns(bool)
func (_YStrategyPrisma *YStrategyPrismaCallerSession) ClaimsAreMaxBoosted() (bool, error) {
	return _YStrategyPrisma.Contract.ClaimsAreMaxBoosted(&_YStrategyPrisma.CallOpts)
}

// ConvexToken is a free data retrieval call binding the contract method 0xe89133b2.
//
// Solidity: function convexToken() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaCaller) ConvexToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YStrategyPrisma.contract.Call(opts, &out, "convexToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConvexToken is a free data retrieval call binding the contract method 0xe89133b2.
//
// Solidity: function convexToken() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaSession) ConvexToken() (common.Address, error) {
	return _YStrategyPrisma.Contract.ConvexToken(&_YStrategyPrisma.CallOpts)
}

// ConvexToken is a free data retrieval call binding the contract method 0xe89133b2.
//
// Solidity: function convexToken() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaCallerSession) ConvexToken() (common.Address, error) {
	return _YStrategyPrisma.Contract.ConvexToken(&_YStrategyPrisma.CallOpts)
}

// ConvexVoter is a free data retrieval call binding the contract method 0xb084e97b.
//
// Solidity: function convexVoter() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaCaller) ConvexVoter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YStrategyPrisma.contract.Call(opts, &out, "convexVoter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConvexVoter is a free data retrieval call binding the contract method 0xb084e97b.
//
// Solidity: function convexVoter() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaSession) ConvexVoter() (common.Address, error) {
	return _YStrategyPrisma.Contract.ConvexVoter(&_YStrategyPrisma.CallOpts)
}

// ConvexVoter is a free data retrieval call binding the contract method 0xb084e97b.
//
// Solidity: function convexVoter() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaCallerSession) ConvexVoter() (common.Address, error) {
	return _YStrategyPrisma.Contract.ConvexVoter(&_YStrategyPrisma.CallOpts)
}

// CreditThreshold is a free data retrieval call binding the contract method 0xaa5480cf.
//
// Solidity: function creditThreshold() view returns(uint256)
func (_YStrategyPrisma *YStrategyPrismaCaller) CreditThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyPrisma.contract.Call(opts, &out, "creditThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreditThreshold is a free data retrieval call binding the contract method 0xaa5480cf.
//
// Solidity: function creditThreshold() view returns(uint256)
func (_YStrategyPrisma *YStrategyPrismaSession) CreditThreshold() (*big.Int, error) {
	return _YStrategyPrisma.Contract.CreditThreshold(&_YStrategyPrisma.CallOpts)
}

// CreditThreshold is a free data retrieval call binding the contract method 0xaa5480cf.
//
// Solidity: function creditThreshold() view returns(uint256)
func (_YStrategyPrisma *YStrategyPrismaCallerSession) CreditThreshold() (*big.Int, error) {
	return _YStrategyPrisma.Contract.CreditThreshold(&_YStrategyPrisma.CallOpts)
}

// Crv is a free data retrieval call binding the contract method 0x6a4874a1.
//
// Solidity: function crv() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaCaller) Crv(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YStrategyPrisma.contract.Call(opts, &out, "crv")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Crv is a free data retrieval call binding the contract method 0x6a4874a1.
//
// Solidity: function crv() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaSession) Crv() (common.Address, error) {
	return _YStrategyPrisma.Contract.Crv(&_YStrategyPrisma.CallOpts)
}

// Crv is a free data retrieval call binding the contract method 0x6a4874a1.
//
// Solidity: function crv() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaCallerSession) Crv() (common.Address, error) {
	return _YStrategyPrisma.Contract.Crv(&_YStrategyPrisma.CallOpts)
}

// CurveVoter is a free data retrieval call binding the contract method 0x65d2ebbf.
//
// Solidity: function curveVoter() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaCaller) CurveVoter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YStrategyPrisma.contract.Call(opts, &out, "curveVoter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurveVoter is a free data retrieval call binding the contract method 0x65d2ebbf.
//
// Solidity: function curveVoter() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaSession) CurveVoter() (common.Address, error) {
	return _YStrategyPrisma.Contract.CurveVoter(&_YStrategyPrisma.CallOpts)
}

// CurveVoter is a free data retrieval call binding the contract method 0x65d2ebbf.
//
// Solidity: function curveVoter() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaCallerSession) CurveVoter() (common.Address, error) {
	return _YStrategyPrisma.Contract.CurveVoter(&_YStrategyPrisma.CallOpts)
}

// DelegatedAssets is a free data retrieval call binding the contract method 0x8e6350e2.
//
// Solidity: function delegatedAssets() view returns(uint256)
func (_YStrategyPrisma *YStrategyPrismaCaller) DelegatedAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyPrisma.contract.Call(opts, &out, "delegatedAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegatedAssets is a free data retrieval call binding the contract method 0x8e6350e2.
//
// Solidity: function delegatedAssets() view returns(uint256)
func (_YStrategyPrisma *YStrategyPrismaSession) DelegatedAssets() (*big.Int, error) {
	return _YStrategyPrisma.Contract.DelegatedAssets(&_YStrategyPrisma.CallOpts)
}

// DelegatedAssets is a free data retrieval call binding the contract method 0x8e6350e2.
//
// Solidity: function delegatedAssets() view returns(uint256)
func (_YStrategyPrisma *YStrategyPrismaCallerSession) DelegatedAssets() (*big.Int, error) {
	return _YStrategyPrisma.Contract.DelegatedAssets(&_YStrategyPrisma.CallOpts)
}

// DoHealthCheck is a free data retrieval call binding the contract method 0x6718835f.
//
// Solidity: function doHealthCheck() view returns(bool)
func (_YStrategyPrisma *YStrategyPrismaCaller) DoHealthCheck(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _YStrategyPrisma.contract.Call(opts, &out, "doHealthCheck")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DoHealthCheck is a free data retrieval call binding the contract method 0x6718835f.
//
// Solidity: function doHealthCheck() view returns(bool)
func (_YStrategyPrisma *YStrategyPrismaSession) DoHealthCheck() (bool, error) {
	return _YStrategyPrisma.Contract.DoHealthCheck(&_YStrategyPrisma.CallOpts)
}

// DoHealthCheck is a free data retrieval call binding the contract method 0x6718835f.
//
// Solidity: function doHealthCheck() view returns(bool)
func (_YStrategyPrisma *YStrategyPrismaCallerSession) DoHealthCheck() (bool, error) {
	return _YStrategyPrisma.Contract.DoHealthCheck(&_YStrategyPrisma.CallOpts)
}

// EmergencyExit is a free data retrieval call binding the contract method 0x5641ec03.
//
// Solidity: function emergencyExit() view returns(bool)
func (_YStrategyPrisma *YStrategyPrismaCaller) EmergencyExit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _YStrategyPrisma.contract.Call(opts, &out, "emergencyExit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EmergencyExit is a free data retrieval call binding the contract method 0x5641ec03.
//
// Solidity: function emergencyExit() view returns(bool)
func (_YStrategyPrisma *YStrategyPrismaSession) EmergencyExit() (bool, error) {
	return _YStrategyPrisma.Contract.EmergencyExit(&_YStrategyPrisma.CallOpts)
}

// EmergencyExit is a free data retrieval call binding the contract method 0x5641ec03.
//
// Solidity: function emergencyExit() view returns(bool)
func (_YStrategyPrisma *YStrategyPrismaCallerSession) EmergencyExit() (bool, error) {
	return _YStrategyPrisma.Contract.EmergencyExit(&_YStrategyPrisma.CallOpts)
}

// EstimatedTotalAssets is a free data retrieval call binding the contract method 0xefbb5cb0.
//
// Solidity: function estimatedTotalAssets() view returns(uint256)
func (_YStrategyPrisma *YStrategyPrismaCaller) EstimatedTotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyPrisma.contract.Call(opts, &out, "estimatedTotalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimatedTotalAssets is a free data retrieval call binding the contract method 0xefbb5cb0.
//
// Solidity: function estimatedTotalAssets() view returns(uint256)
func (_YStrategyPrisma *YStrategyPrismaSession) EstimatedTotalAssets() (*big.Int, error) {
	return _YStrategyPrisma.Contract.EstimatedTotalAssets(&_YStrategyPrisma.CallOpts)
}

// EstimatedTotalAssets is a free data retrieval call binding the contract method 0xefbb5cb0.
//
// Solidity: function estimatedTotalAssets() view returns(uint256)
func (_YStrategyPrisma *YStrategyPrismaCallerSession) EstimatedTotalAssets() (*big.Int, error) {
	return _YStrategyPrisma.Contract.EstimatedTotalAssets(&_YStrategyPrisma.CallOpts)
}

// EthToWant is a free data retrieval call binding the contract method 0x780022a0.
//
// Solidity: function ethToWant(uint256 _ethAmount) view returns(uint256)
func (_YStrategyPrisma *YStrategyPrismaCaller) EthToWant(opts *bind.CallOpts, _ethAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyPrisma.contract.Call(opts, &out, "ethToWant", _ethAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EthToWant is a free data retrieval call binding the contract method 0x780022a0.
//
// Solidity: function ethToWant(uint256 _ethAmount) view returns(uint256)
func (_YStrategyPrisma *YStrategyPrismaSession) EthToWant(_ethAmount *big.Int) (*big.Int, error) {
	return _YStrategyPrisma.Contract.EthToWant(&_YStrategyPrisma.CallOpts, _ethAmount)
}

// EthToWant is a free data retrieval call binding the contract method 0x780022a0.
//
// Solidity: function ethToWant(uint256 _ethAmount) view returns(uint256)
func (_YStrategyPrisma *YStrategyPrismaCallerSession) EthToWant(_ethAmount *big.Int) (*big.Int, error) {
	return _YStrategyPrisma.Contract.EthToWant(&_YStrategyPrisma.CallOpts, _ethAmount)
}

// ForceClaimOnce is a free data retrieval call binding the contract method 0x2a99c9df.
//
// Solidity: function forceClaimOnce() view returns(bool)
func (_YStrategyPrisma *YStrategyPrismaCaller) ForceClaimOnce(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _YStrategyPrisma.contract.Call(opts, &out, "forceClaimOnce")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ForceClaimOnce is a free data retrieval call binding the contract method 0x2a99c9df.
//
// Solidity: function forceClaimOnce() view returns(bool)
func (_YStrategyPrisma *YStrategyPrismaSession) ForceClaimOnce() (bool, error) {
	return _YStrategyPrisma.Contract.ForceClaimOnce(&_YStrategyPrisma.CallOpts)
}

// ForceClaimOnce is a free data retrieval call binding the contract method 0x2a99c9df.
//
// Solidity: function forceClaimOnce() view returns(bool)
func (_YStrategyPrisma *YStrategyPrismaCallerSession) ForceClaimOnce() (bool, error) {
	return _YStrategyPrisma.Contract.ForceClaimOnce(&_YStrategyPrisma.CallOpts)
}

// ForceHarvestTriggerOnce is a free data retrieval call binding the contract method 0xa763cf5b.
//
// Solidity: function forceHarvestTriggerOnce() view returns(bool)
func (_YStrategyPrisma *YStrategyPrismaCaller) ForceHarvestTriggerOnce(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _YStrategyPrisma.contract.Call(opts, &out, "forceHarvestTriggerOnce")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ForceHarvestTriggerOnce is a free data retrieval call binding the contract method 0xa763cf5b.
//
// Solidity: function forceHarvestTriggerOnce() view returns(bool)
func (_YStrategyPrisma *YStrategyPrismaSession) ForceHarvestTriggerOnce() (bool, error) {
	return _YStrategyPrisma.Contract.ForceHarvestTriggerOnce(&_YStrategyPrisma.CallOpts)
}

// ForceHarvestTriggerOnce is a free data retrieval call binding the contract method 0xa763cf5b.
//
// Solidity: function forceHarvestTriggerOnce() view returns(bool)
func (_YStrategyPrisma *YStrategyPrismaCallerSession) ForceHarvestTriggerOnce() (bool, error) {
	return _YStrategyPrisma.Contract.ForceHarvestTriggerOnce(&_YStrategyPrisma.CallOpts)
}

// HarvestProfitMaxInUsdc is a free data retrieval call binding the contract method 0xfa4e2df9.
//
// Solidity: function harvestProfitMaxInUsdc() view returns(uint256)
func (_YStrategyPrisma *YStrategyPrismaCaller) HarvestProfitMaxInUsdc(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyPrisma.contract.Call(opts, &out, "harvestProfitMaxInUsdc")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HarvestProfitMaxInUsdc is a free data retrieval call binding the contract method 0xfa4e2df9.
//
// Solidity: function harvestProfitMaxInUsdc() view returns(uint256)
func (_YStrategyPrisma *YStrategyPrismaSession) HarvestProfitMaxInUsdc() (*big.Int, error) {
	return _YStrategyPrisma.Contract.HarvestProfitMaxInUsdc(&_YStrategyPrisma.CallOpts)
}

// HarvestProfitMaxInUsdc is a free data retrieval call binding the contract method 0xfa4e2df9.
//
// Solidity: function harvestProfitMaxInUsdc() view returns(uint256)
func (_YStrategyPrisma *YStrategyPrismaCallerSession) HarvestProfitMaxInUsdc() (*big.Int, error) {
	return _YStrategyPrisma.Contract.HarvestProfitMaxInUsdc(&_YStrategyPrisma.CallOpts)
}

// HarvestProfitMinInUsdc is a free data retrieval call binding the contract method 0x090c4922.
//
// Solidity: function harvestProfitMinInUsdc() view returns(uint256)
func (_YStrategyPrisma *YStrategyPrismaCaller) HarvestProfitMinInUsdc(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyPrisma.contract.Call(opts, &out, "harvestProfitMinInUsdc")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HarvestProfitMinInUsdc is a free data retrieval call binding the contract method 0x090c4922.
//
// Solidity: function harvestProfitMinInUsdc() view returns(uint256)
func (_YStrategyPrisma *YStrategyPrismaSession) HarvestProfitMinInUsdc() (*big.Int, error) {
	return _YStrategyPrisma.Contract.HarvestProfitMinInUsdc(&_YStrategyPrisma.CallOpts)
}

// HarvestProfitMinInUsdc is a free data retrieval call binding the contract method 0x090c4922.
//
// Solidity: function harvestProfitMinInUsdc() view returns(uint256)
func (_YStrategyPrisma *YStrategyPrismaCallerSession) HarvestProfitMinInUsdc() (*big.Int, error) {
	return _YStrategyPrisma.Contract.HarvestProfitMinInUsdc(&_YStrategyPrisma.CallOpts)
}

// HarvestTrigger is a free data retrieval call binding the contract method 0xed882c2b.
//
// Solidity: function harvestTrigger(uint256 callCostinEth) view returns(bool)
func (_YStrategyPrisma *YStrategyPrismaCaller) HarvestTrigger(opts *bind.CallOpts, callCostinEth *big.Int) (bool, error) {
	var out []interface{}
	err := _YStrategyPrisma.contract.Call(opts, &out, "harvestTrigger", callCostinEth)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HarvestTrigger is a free data retrieval call binding the contract method 0xed882c2b.
//
// Solidity: function harvestTrigger(uint256 callCostinEth) view returns(bool)
func (_YStrategyPrisma *YStrategyPrismaSession) HarvestTrigger(callCostinEth *big.Int) (bool, error) {
	return _YStrategyPrisma.Contract.HarvestTrigger(&_YStrategyPrisma.CallOpts, callCostinEth)
}

// HarvestTrigger is a free data retrieval call binding the contract method 0xed882c2b.
//
// Solidity: function harvestTrigger(uint256 callCostinEth) view returns(bool)
func (_YStrategyPrisma *YStrategyPrismaCallerSession) HarvestTrigger(callCostinEth *big.Int) (bool, error) {
	return _YStrategyPrisma.Contract.HarvestTrigger(&_YStrategyPrisma.CallOpts, callCostinEth)
}

// HealthCheck is a free data retrieval call binding the contract method 0xb252720b.
//
// Solidity: function healthCheck() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaCaller) HealthCheck(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YStrategyPrisma.contract.Call(opts, &out, "healthCheck")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HealthCheck is a free data retrieval call binding the contract method 0xb252720b.
//
// Solidity: function healthCheck() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaSession) HealthCheck() (common.Address, error) {
	return _YStrategyPrisma.Contract.HealthCheck(&_YStrategyPrisma.CallOpts)
}

// HealthCheck is a free data retrieval call binding the contract method 0xb252720b.
//
// Solidity: function healthCheck() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaCallerSession) HealthCheck() (common.Address, error) {
	return _YStrategyPrisma.Contract.HealthCheck(&_YStrategyPrisma.CallOpts)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_YStrategyPrisma *YStrategyPrismaCaller) IsActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _YStrategyPrisma.contract.Call(opts, &out, "isActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_YStrategyPrisma *YStrategyPrismaSession) IsActive() (bool, error) {
	return _YStrategyPrisma.Contract.IsActive(&_YStrategyPrisma.CallOpts)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_YStrategyPrisma *YStrategyPrismaCallerSession) IsActive() (bool, error) {
	return _YStrategyPrisma.Contract.IsActive(&_YStrategyPrisma.CallOpts)
}

// IsBaseFeeAcceptable is a free data retrieval call binding the contract method 0x95326e2d.
//
// Solidity: function isBaseFeeAcceptable() view returns(bool)
func (_YStrategyPrisma *YStrategyPrismaCaller) IsBaseFeeAcceptable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _YStrategyPrisma.contract.Call(opts, &out, "isBaseFeeAcceptable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBaseFeeAcceptable is a free data retrieval call binding the contract method 0x95326e2d.
//
// Solidity: function isBaseFeeAcceptable() view returns(bool)
func (_YStrategyPrisma *YStrategyPrismaSession) IsBaseFeeAcceptable() (bool, error) {
	return _YStrategyPrisma.Contract.IsBaseFeeAcceptable(&_YStrategyPrisma.CallOpts)
}

// IsBaseFeeAcceptable is a free data retrieval call binding the contract method 0x95326e2d.
//
// Solidity: function isBaseFeeAcceptable() view returns(bool)
func (_YStrategyPrisma *YStrategyPrismaCallerSession) IsBaseFeeAcceptable() (bool, error) {
	return _YStrategyPrisma.Contract.IsBaseFeeAcceptable(&_YStrategyPrisma.CallOpts)
}

// IsOriginal is a free data retrieval call binding the contract method 0x6f392ce7.
//
// Solidity: function isOriginal() view returns(bool)
func (_YStrategyPrisma *YStrategyPrismaCaller) IsOriginal(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _YStrategyPrisma.contract.Call(opts, &out, "isOriginal")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOriginal is a free data retrieval call binding the contract method 0x6f392ce7.
//
// Solidity: function isOriginal() view returns(bool)
func (_YStrategyPrisma *YStrategyPrismaSession) IsOriginal() (bool, error) {
	return _YStrategyPrisma.Contract.IsOriginal(&_YStrategyPrisma.CallOpts)
}

// IsOriginal is a free data retrieval call binding the contract method 0x6f392ce7.
//
// Solidity: function isOriginal() view returns(bool)
func (_YStrategyPrisma *YStrategyPrismaCallerSession) IsOriginal() (bool, error) {
	return _YStrategyPrisma.Contract.IsOriginal(&_YStrategyPrisma.CallOpts)
}

// Keeper is a free data retrieval call binding the contract method 0xaced1661.
//
// Solidity: function keeper() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaCaller) Keeper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YStrategyPrisma.contract.Call(opts, &out, "keeper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Keeper is a free data retrieval call binding the contract method 0xaced1661.
//
// Solidity: function keeper() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaSession) Keeper() (common.Address, error) {
	return _YStrategyPrisma.Contract.Keeper(&_YStrategyPrisma.CallOpts)
}

// Keeper is a free data retrieval call binding the contract method 0xaced1661.
//
// Solidity: function keeper() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaCallerSession) Keeper() (common.Address, error) {
	return _YStrategyPrisma.Contract.Keeper(&_YStrategyPrisma.CallOpts)
}

// LocalKeepCRV is a free data retrieval call binding the contract method 0x73fd827f.
//
// Solidity: function localKeepCRV() view returns(uint256)
func (_YStrategyPrisma *YStrategyPrismaCaller) LocalKeepCRV(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyPrisma.contract.Call(opts, &out, "localKeepCRV")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LocalKeepCRV is a free data retrieval call binding the contract method 0x73fd827f.
//
// Solidity: function localKeepCRV() view returns(uint256)
func (_YStrategyPrisma *YStrategyPrismaSession) LocalKeepCRV() (*big.Int, error) {
	return _YStrategyPrisma.Contract.LocalKeepCRV(&_YStrategyPrisma.CallOpts)
}

// LocalKeepCRV is a free data retrieval call binding the contract method 0x73fd827f.
//
// Solidity: function localKeepCRV() view returns(uint256)
func (_YStrategyPrisma *YStrategyPrismaCallerSession) LocalKeepCRV() (*big.Int, error) {
	return _YStrategyPrisma.Contract.LocalKeepCRV(&_YStrategyPrisma.CallOpts)
}

// LocalKeepCVX is a free data retrieval call binding the contract method 0x28f30a4c.
//
// Solidity: function localKeepCVX() view returns(uint256)
func (_YStrategyPrisma *YStrategyPrismaCaller) LocalKeepCVX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyPrisma.contract.Call(opts, &out, "localKeepCVX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LocalKeepCVX is a free data retrieval call binding the contract method 0x28f30a4c.
//
// Solidity: function localKeepCVX() view returns(uint256)
func (_YStrategyPrisma *YStrategyPrismaSession) LocalKeepCVX() (*big.Int, error) {
	return _YStrategyPrisma.Contract.LocalKeepCVX(&_YStrategyPrisma.CallOpts)
}

// LocalKeepCVX is a free data retrieval call binding the contract method 0x28f30a4c.
//
// Solidity: function localKeepCVX() view returns(uint256)
func (_YStrategyPrisma *YStrategyPrismaCallerSession) LocalKeepCVX() (*big.Int, error) {
	return _YStrategyPrisma.Contract.LocalKeepCVX(&_YStrategyPrisma.CallOpts)
}

// LocalKeepYPrisma is a free data retrieval call binding the contract method 0x6d208c45.
//
// Solidity: function localKeepYPrisma() view returns(uint256)
func (_YStrategyPrisma *YStrategyPrismaCaller) LocalKeepYPrisma(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyPrisma.contract.Call(opts, &out, "localKeepYPrisma")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LocalKeepYPrisma is a free data retrieval call binding the contract method 0x6d208c45.
//
// Solidity: function localKeepYPrisma() view returns(uint256)
func (_YStrategyPrisma *YStrategyPrismaSession) LocalKeepYPrisma() (*big.Int, error) {
	return _YStrategyPrisma.Contract.LocalKeepYPrisma(&_YStrategyPrisma.CallOpts)
}

// LocalKeepYPrisma is a free data retrieval call binding the contract method 0x6d208c45.
//
// Solidity: function localKeepYPrisma() view returns(uint256)
func (_YStrategyPrisma *YStrategyPrismaCallerSession) LocalKeepYPrisma() (*big.Int, error) {
	return _YStrategyPrisma.Contract.LocalKeepYPrisma(&_YStrategyPrisma.CallOpts)
}

// MaxReportDelay is a free data retrieval call binding the contract method 0x28b7ccf7.
//
// Solidity: function maxReportDelay() view returns(uint256)
func (_YStrategyPrisma *YStrategyPrismaCaller) MaxReportDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyPrisma.contract.Call(opts, &out, "maxReportDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxReportDelay is a free data retrieval call binding the contract method 0x28b7ccf7.
//
// Solidity: function maxReportDelay() view returns(uint256)
func (_YStrategyPrisma *YStrategyPrismaSession) MaxReportDelay() (*big.Int, error) {
	return _YStrategyPrisma.Contract.MaxReportDelay(&_YStrategyPrisma.CallOpts)
}

// MaxReportDelay is a free data retrieval call binding the contract method 0x28b7ccf7.
//
// Solidity: function maxReportDelay() view returns(uint256)
func (_YStrategyPrisma *YStrategyPrismaCallerSession) MaxReportDelay() (*big.Int, error) {
	return _YStrategyPrisma.Contract.MaxReportDelay(&_YStrategyPrisma.CallOpts)
}

// MetadataURI is a free data retrieval call binding the contract method 0x03ee438c.
//
// Solidity: function metadataURI() view returns(string)
func (_YStrategyPrisma *YStrategyPrismaCaller) MetadataURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _YStrategyPrisma.contract.Call(opts, &out, "metadataURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MetadataURI is a free data retrieval call binding the contract method 0x03ee438c.
//
// Solidity: function metadataURI() view returns(string)
func (_YStrategyPrisma *YStrategyPrismaSession) MetadataURI() (string, error) {
	return _YStrategyPrisma.Contract.MetadataURI(&_YStrategyPrisma.CallOpts)
}

// MetadataURI is a free data retrieval call binding the contract method 0x03ee438c.
//
// Solidity: function metadataURI() view returns(string)
func (_YStrategyPrisma *YStrategyPrismaCallerSession) MetadataURI() (string, error) {
	return _YStrategyPrisma.Contract.MetadataURI(&_YStrategyPrisma.CallOpts)
}

// MinReportDelay is a free data retrieval call binding the contract method 0x95e80c50.
//
// Solidity: function minReportDelay() view returns(uint256)
func (_YStrategyPrisma *YStrategyPrismaCaller) MinReportDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyPrisma.contract.Call(opts, &out, "minReportDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinReportDelay is a free data retrieval call binding the contract method 0x95e80c50.
//
// Solidity: function minReportDelay() view returns(uint256)
func (_YStrategyPrisma *YStrategyPrismaSession) MinReportDelay() (*big.Int, error) {
	return _YStrategyPrisma.Contract.MinReportDelay(&_YStrategyPrisma.CallOpts)
}

// MinReportDelay is a free data retrieval call binding the contract method 0x95e80c50.
//
// Solidity: function minReportDelay() view returns(uint256)
func (_YStrategyPrisma *YStrategyPrismaCallerSession) MinReportDelay() (*big.Int, error) {
	return _YStrategyPrisma.Contract.MinReportDelay(&_YStrategyPrisma.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YStrategyPrisma *YStrategyPrismaCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _YStrategyPrisma.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YStrategyPrisma *YStrategyPrismaSession) Name() (string, error) {
	return _YStrategyPrisma.Contract.Name(&_YStrategyPrisma.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YStrategyPrisma *YStrategyPrismaCallerSession) Name() (string, error) {
	return _YStrategyPrisma.Contract.Name(&_YStrategyPrisma.CallOpts)
}

// PrismaReceiver is a free data retrieval call binding the contract method 0xb4ef5af4.
//
// Solidity: function prismaReceiver() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaCaller) PrismaReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YStrategyPrisma.contract.Call(opts, &out, "prismaReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrismaReceiver is a free data retrieval call binding the contract method 0xb4ef5af4.
//
// Solidity: function prismaReceiver() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaSession) PrismaReceiver() (common.Address, error) {
	return _YStrategyPrisma.Contract.PrismaReceiver(&_YStrategyPrisma.CallOpts)
}

// PrismaReceiver is a free data retrieval call binding the contract method 0xb4ef5af4.
//
// Solidity: function prismaReceiver() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaCallerSession) PrismaReceiver() (common.Address, error) {
	return _YStrategyPrisma.Contract.PrismaReceiver(&_YStrategyPrisma.CallOpts)
}

// PrismaVault is a free data retrieval call binding the contract method 0xe59d7653.
//
// Solidity: function prismaVault() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaCaller) PrismaVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YStrategyPrisma.contract.Call(opts, &out, "prismaVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrismaVault is a free data retrieval call binding the contract method 0xe59d7653.
//
// Solidity: function prismaVault() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaSession) PrismaVault() (common.Address, error) {
	return _YStrategyPrisma.Contract.PrismaVault(&_YStrategyPrisma.CallOpts)
}

// PrismaVault is a free data retrieval call binding the contract method 0xe59d7653.
//
// Solidity: function prismaVault() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaCallerSession) PrismaVault() (common.Address, error) {
	return _YStrategyPrisma.Contract.PrismaVault(&_YStrategyPrisma.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x9ec5a894.
//
// Solidity: function rewards() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaCaller) Rewards(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YStrategyPrisma.contract.Call(opts, &out, "rewards")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0x9ec5a894.
//
// Solidity: function rewards() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaSession) Rewards() (common.Address, error) {
	return _YStrategyPrisma.Contract.Rewards(&_YStrategyPrisma.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x9ec5a894.
//
// Solidity: function rewards() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaCallerSession) Rewards() (common.Address, error) {
	return _YStrategyPrisma.Contract.Rewards(&_YStrategyPrisma.CallOpts)
}

// StakedBalance is a free data retrieval call binding the contract method 0x5b9f0016.
//
// Solidity: function stakedBalance() view returns(uint256)
func (_YStrategyPrisma *YStrategyPrismaCaller) StakedBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YStrategyPrisma.contract.Call(opts, &out, "stakedBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakedBalance is a free data retrieval call binding the contract method 0x5b9f0016.
//
// Solidity: function stakedBalance() view returns(uint256)
func (_YStrategyPrisma *YStrategyPrismaSession) StakedBalance() (*big.Int, error) {
	return _YStrategyPrisma.Contract.StakedBalance(&_YStrategyPrisma.CallOpts)
}

// StakedBalance is a free data retrieval call binding the contract method 0x5b9f0016.
//
// Solidity: function stakedBalance() view returns(uint256)
func (_YStrategyPrisma *YStrategyPrismaCallerSession) StakedBalance() (*big.Int, error) {
	return _YStrategyPrisma.Contract.StakedBalance(&_YStrategyPrisma.CallOpts)
}

// Strategist is a free data retrieval call binding the contract method 0x1fe4a686.
//
// Solidity: function strategist() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaCaller) Strategist(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YStrategyPrisma.contract.Call(opts, &out, "strategist")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Strategist is a free data retrieval call binding the contract method 0x1fe4a686.
//
// Solidity: function strategist() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaSession) Strategist() (common.Address, error) {
	return _YStrategyPrisma.Contract.Strategist(&_YStrategyPrisma.CallOpts)
}

// Strategist is a free data retrieval call binding the contract method 0x1fe4a686.
//
// Solidity: function strategist() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaCallerSession) Strategist() (common.Address, error) {
	return _YStrategyPrisma.Contract.Strategist(&_YStrategyPrisma.CallOpts)
}

// TendTrigger is a free data retrieval call binding the contract method 0x650d1880.
//
// Solidity: function tendTrigger(uint256 callCostInWei) view returns(bool)
func (_YStrategyPrisma *YStrategyPrismaCaller) TendTrigger(opts *bind.CallOpts, callCostInWei *big.Int) (bool, error) {
	var out []interface{}
	err := _YStrategyPrisma.contract.Call(opts, &out, "tendTrigger", callCostInWei)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TendTrigger is a free data retrieval call binding the contract method 0x650d1880.
//
// Solidity: function tendTrigger(uint256 callCostInWei) view returns(bool)
func (_YStrategyPrisma *YStrategyPrismaSession) TendTrigger(callCostInWei *big.Int) (bool, error) {
	return _YStrategyPrisma.Contract.TendTrigger(&_YStrategyPrisma.CallOpts, callCostInWei)
}

// TendTrigger is a free data retrieval call binding the contract method 0x650d1880.
//
// Solidity: function tendTrigger(uint256 callCostInWei) view returns(bool)
func (_YStrategyPrisma *YStrategyPrismaCallerSession) TendTrigger(callCostInWei *big.Int) (bool, error) {
	return _YStrategyPrisma.Contract.TendTrigger(&_YStrategyPrisma.CallOpts, callCostInWei)
}

// TradeFactory is a free data retrieval call binding the contract method 0xe5e19b4a.
//
// Solidity: function tradeFactory() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaCaller) TradeFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YStrategyPrisma.contract.Call(opts, &out, "tradeFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TradeFactory is a free data retrieval call binding the contract method 0xe5e19b4a.
//
// Solidity: function tradeFactory() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaSession) TradeFactory() (common.Address, error) {
	return _YStrategyPrisma.Contract.TradeFactory(&_YStrategyPrisma.CallOpts)
}

// TradeFactory is a free data retrieval call binding the contract method 0xe5e19b4a.
//
// Solidity: function tradeFactory() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaCallerSession) TradeFactory() (common.Address, error) {
	return _YStrategyPrisma.Contract.TradeFactory(&_YStrategyPrisma.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YStrategyPrisma.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaSession) Vault() (common.Address, error) {
	return _YStrategyPrisma.Contract.Vault(&_YStrategyPrisma.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaCallerSession) Vault() (common.Address, error) {
	return _YStrategyPrisma.Contract.Vault(&_YStrategyPrisma.CallOpts)
}

// Want is a free data retrieval call binding the contract method 0x1f1fcd51.
//
// Solidity: function want() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaCaller) Want(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YStrategyPrisma.contract.Call(opts, &out, "want")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Want is a free data retrieval call binding the contract method 0x1f1fcd51.
//
// Solidity: function want() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaSession) Want() (common.Address, error) {
	return _YStrategyPrisma.Contract.Want(&_YStrategyPrisma.CallOpts)
}

// Want is a free data retrieval call binding the contract method 0x1f1fcd51.
//
// Solidity: function want() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaCallerSession) Want() (common.Address, error) {
	return _YStrategyPrisma.Contract.Want(&_YStrategyPrisma.CallOpts)
}

// YPrisma is a free data retrieval call binding the contract method 0x3e2ece13.
//
// Solidity: function yPrisma() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaCaller) YPrisma(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YStrategyPrisma.contract.Call(opts, &out, "yPrisma")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// YPrisma is a free data retrieval call binding the contract method 0x3e2ece13.
//
// Solidity: function yPrisma() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaSession) YPrisma() (common.Address, error) {
	return _YStrategyPrisma.Contract.YPrisma(&_YStrategyPrisma.CallOpts)
}

// YPrisma is a free data retrieval call binding the contract method 0x3e2ece13.
//
// Solidity: function yPrisma() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaCallerSession) YPrisma() (common.Address, error) {
	return _YStrategyPrisma.Contract.YPrisma(&_YStrategyPrisma.CallOpts)
}

// YprismaVoter is a free data retrieval call binding the contract method 0x6a727f41.
//
// Solidity: function yprismaVoter() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaCaller) YprismaVoter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YStrategyPrisma.contract.Call(opts, &out, "yprismaVoter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// YprismaVoter is a free data retrieval call binding the contract method 0x6a727f41.
//
// Solidity: function yprismaVoter() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaSession) YprismaVoter() (common.Address, error) {
	return _YStrategyPrisma.Contract.YprismaVoter(&_YStrategyPrisma.CallOpts)
}

// YprismaVoter is a free data retrieval call binding the contract method 0x6a727f41.
//
// Solidity: function yprismaVoter() view returns(address)
func (_YStrategyPrisma *YStrategyPrismaCallerSession) YprismaVoter() (common.Address, error) {
	return _YStrategyPrisma.Contract.YprismaVoter(&_YStrategyPrisma.CallOpts)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() returns()
func (_YStrategyPrisma *YStrategyPrismaTransactor) ClaimRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YStrategyPrisma.contract.Transact(opts, "claimRewards")
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() returns()
func (_YStrategyPrisma *YStrategyPrismaSession) ClaimRewards() (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.ClaimRewards(&_YStrategyPrisma.TransactOpts)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() returns()
func (_YStrategyPrisma *YStrategyPrismaTransactorSession) ClaimRewards() (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.ClaimRewards(&_YStrategyPrisma.TransactOpts)
}

// CloneStrategyPrismaConvex is a paid mutator transaction binding the contract method 0xc66cfed4.
//
// Solidity: function cloneStrategyPrismaConvex(address _vault, address _strategist, address _rewards, address _keeper, address _tradeFactory, uint256 _harvestProfitMinInUsdc, uint256 _harvestProfitMaxInUsdc, address _prismaVault, address _prismaReceiver) returns(address newStrategy)
func (_YStrategyPrisma *YStrategyPrismaTransactor) CloneStrategyPrismaConvex(opts *bind.TransactOpts, _vault common.Address, _strategist common.Address, _rewards common.Address, _keeper common.Address, _tradeFactory common.Address, _harvestProfitMinInUsdc *big.Int, _harvestProfitMaxInUsdc *big.Int, _prismaVault common.Address, _prismaReceiver common.Address) (*types.Transaction, error) {
	return _YStrategyPrisma.contract.Transact(opts, "cloneStrategyPrismaConvex", _vault, _strategist, _rewards, _keeper, _tradeFactory, _harvestProfitMinInUsdc, _harvestProfitMaxInUsdc, _prismaVault, _prismaReceiver)
}

// CloneStrategyPrismaConvex is a paid mutator transaction binding the contract method 0xc66cfed4.
//
// Solidity: function cloneStrategyPrismaConvex(address _vault, address _strategist, address _rewards, address _keeper, address _tradeFactory, uint256 _harvestProfitMinInUsdc, uint256 _harvestProfitMaxInUsdc, address _prismaVault, address _prismaReceiver) returns(address newStrategy)
func (_YStrategyPrisma *YStrategyPrismaSession) CloneStrategyPrismaConvex(_vault common.Address, _strategist common.Address, _rewards common.Address, _keeper common.Address, _tradeFactory common.Address, _harvestProfitMinInUsdc *big.Int, _harvestProfitMaxInUsdc *big.Int, _prismaVault common.Address, _prismaReceiver common.Address) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.CloneStrategyPrismaConvex(&_YStrategyPrisma.TransactOpts, _vault, _strategist, _rewards, _keeper, _tradeFactory, _harvestProfitMinInUsdc, _harvestProfitMaxInUsdc, _prismaVault, _prismaReceiver)
}

// CloneStrategyPrismaConvex is a paid mutator transaction binding the contract method 0xc66cfed4.
//
// Solidity: function cloneStrategyPrismaConvex(address _vault, address _strategist, address _rewards, address _keeper, address _tradeFactory, uint256 _harvestProfitMinInUsdc, uint256 _harvestProfitMaxInUsdc, address _prismaVault, address _prismaReceiver) returns(address newStrategy)
func (_YStrategyPrisma *YStrategyPrismaTransactorSession) CloneStrategyPrismaConvex(_vault common.Address, _strategist common.Address, _rewards common.Address, _keeper common.Address, _tradeFactory common.Address, _harvestProfitMinInUsdc *big.Int, _harvestProfitMaxInUsdc *big.Int, _prismaVault common.Address, _prismaReceiver common.Address) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.CloneStrategyPrismaConvex(&_YStrategyPrisma.TransactOpts, _vault, _strategist, _rewards, _keeper, _tradeFactory, _harvestProfitMinInUsdc, _harvestProfitMaxInUsdc, _prismaVault, _prismaReceiver)
}

// Harvest is a paid mutator transaction binding the contract method 0x4641257d.
//
// Solidity: function harvest() returns()
func (_YStrategyPrisma *YStrategyPrismaTransactor) Harvest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YStrategyPrisma.contract.Transact(opts, "harvest")
}

// Harvest is a paid mutator transaction binding the contract method 0x4641257d.
//
// Solidity: function harvest() returns()
func (_YStrategyPrisma *YStrategyPrismaSession) Harvest() (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.Harvest(&_YStrategyPrisma.TransactOpts)
}

// Harvest is a paid mutator transaction binding the contract method 0x4641257d.
//
// Solidity: function harvest() returns()
func (_YStrategyPrisma *YStrategyPrismaTransactorSession) Harvest() (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.Harvest(&_YStrategyPrisma.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x042c50dd.
//
// Solidity: function initialize(address _vault, address _strategist, address _rewards, address _keeper, address _tradeFactory, uint256 _harvestProfitMinInUsdc, uint256 _harvestProfitMaxInUsdc, address _prismaVault, address _prismaReceiver) returns()
func (_YStrategyPrisma *YStrategyPrismaTransactor) Initialize(opts *bind.TransactOpts, _vault common.Address, _strategist common.Address, _rewards common.Address, _keeper common.Address, _tradeFactory common.Address, _harvestProfitMinInUsdc *big.Int, _harvestProfitMaxInUsdc *big.Int, _prismaVault common.Address, _prismaReceiver common.Address) (*types.Transaction, error) {
	return _YStrategyPrisma.contract.Transact(opts, "initialize", _vault, _strategist, _rewards, _keeper, _tradeFactory, _harvestProfitMinInUsdc, _harvestProfitMaxInUsdc, _prismaVault, _prismaReceiver)
}

// Initialize is a paid mutator transaction binding the contract method 0x042c50dd.
//
// Solidity: function initialize(address _vault, address _strategist, address _rewards, address _keeper, address _tradeFactory, uint256 _harvestProfitMinInUsdc, uint256 _harvestProfitMaxInUsdc, address _prismaVault, address _prismaReceiver) returns()
func (_YStrategyPrisma *YStrategyPrismaSession) Initialize(_vault common.Address, _strategist common.Address, _rewards common.Address, _keeper common.Address, _tradeFactory common.Address, _harvestProfitMinInUsdc *big.Int, _harvestProfitMaxInUsdc *big.Int, _prismaVault common.Address, _prismaReceiver common.Address) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.Initialize(&_YStrategyPrisma.TransactOpts, _vault, _strategist, _rewards, _keeper, _tradeFactory, _harvestProfitMinInUsdc, _harvestProfitMaxInUsdc, _prismaVault, _prismaReceiver)
}

// Initialize is a paid mutator transaction binding the contract method 0x042c50dd.
//
// Solidity: function initialize(address _vault, address _strategist, address _rewards, address _keeper, address _tradeFactory, uint256 _harvestProfitMinInUsdc, uint256 _harvestProfitMaxInUsdc, address _prismaVault, address _prismaReceiver) returns()
func (_YStrategyPrisma *YStrategyPrismaTransactorSession) Initialize(_vault common.Address, _strategist common.Address, _rewards common.Address, _keeper common.Address, _tradeFactory common.Address, _harvestProfitMinInUsdc *big.Int, _harvestProfitMaxInUsdc *big.Int, _prismaVault common.Address, _prismaReceiver common.Address) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.Initialize(&_YStrategyPrisma.TransactOpts, _vault, _strategist, _rewards, _keeper, _tradeFactory, _harvestProfitMinInUsdc, _harvestProfitMaxInUsdc, _prismaVault, _prismaReceiver)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _newStrategy) returns()
func (_YStrategyPrisma *YStrategyPrismaTransactor) Migrate(opts *bind.TransactOpts, _newStrategy common.Address) (*types.Transaction, error) {
	return _YStrategyPrisma.contract.Transact(opts, "migrate", _newStrategy)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _newStrategy) returns()
func (_YStrategyPrisma *YStrategyPrismaSession) Migrate(_newStrategy common.Address) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.Migrate(&_YStrategyPrisma.TransactOpts, _newStrategy)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _newStrategy) returns()
func (_YStrategyPrisma *YStrategyPrismaTransactorSession) Migrate(_newStrategy common.Address) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.Migrate(&_YStrategyPrisma.TransactOpts, _newStrategy)
}

// RemoveTradeFactoryPermissions is a paid mutator transaction binding the contract method 0xe09575a4.
//
// Solidity: function removeTradeFactoryPermissions(bool _disableTf) returns()
func (_YStrategyPrisma *YStrategyPrismaTransactor) RemoveTradeFactoryPermissions(opts *bind.TransactOpts, _disableTf bool) (*types.Transaction, error) {
	return _YStrategyPrisma.contract.Transact(opts, "removeTradeFactoryPermissions", _disableTf)
}

// RemoveTradeFactoryPermissions is a paid mutator transaction binding the contract method 0xe09575a4.
//
// Solidity: function removeTradeFactoryPermissions(bool _disableTf) returns()
func (_YStrategyPrisma *YStrategyPrismaSession) RemoveTradeFactoryPermissions(_disableTf bool) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.RemoveTradeFactoryPermissions(&_YStrategyPrisma.TransactOpts, _disableTf)
}

// RemoveTradeFactoryPermissions is a paid mutator transaction binding the contract method 0xe09575a4.
//
// Solidity: function removeTradeFactoryPermissions(bool _disableTf) returns()
func (_YStrategyPrisma *YStrategyPrismaTransactorSession) RemoveTradeFactoryPermissions(_disableTf bool) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.RemoveTradeFactoryPermissions(&_YStrategyPrisma.TransactOpts, _disableTf)
}

// SetBaseFeeOracle is a paid mutator transaction binding the contract method 0x9f450b5a.
//
// Solidity: function setBaseFeeOracle(address _baseFeeOracle) returns()
func (_YStrategyPrisma *YStrategyPrismaTransactor) SetBaseFeeOracle(opts *bind.TransactOpts, _baseFeeOracle common.Address) (*types.Transaction, error) {
	return _YStrategyPrisma.contract.Transact(opts, "setBaseFeeOracle", _baseFeeOracle)
}

// SetBaseFeeOracle is a paid mutator transaction binding the contract method 0x9f450b5a.
//
// Solidity: function setBaseFeeOracle(address _baseFeeOracle) returns()
func (_YStrategyPrisma *YStrategyPrismaSession) SetBaseFeeOracle(_baseFeeOracle common.Address) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.SetBaseFeeOracle(&_YStrategyPrisma.TransactOpts, _baseFeeOracle)
}

// SetBaseFeeOracle is a paid mutator transaction binding the contract method 0x9f450b5a.
//
// Solidity: function setBaseFeeOracle(address _baseFeeOracle) returns()
func (_YStrategyPrisma *YStrategyPrismaTransactorSession) SetBaseFeeOracle(_baseFeeOracle common.Address) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.SetBaseFeeOracle(&_YStrategyPrisma.TransactOpts, _baseFeeOracle)
}

// SetCreditThreshold is a paid mutator transaction binding the contract method 0xfe2508a6.
//
// Solidity: function setCreditThreshold(uint256 _creditThreshold) returns()
func (_YStrategyPrisma *YStrategyPrismaTransactor) SetCreditThreshold(opts *bind.TransactOpts, _creditThreshold *big.Int) (*types.Transaction, error) {
	return _YStrategyPrisma.contract.Transact(opts, "setCreditThreshold", _creditThreshold)
}

// SetCreditThreshold is a paid mutator transaction binding the contract method 0xfe2508a6.
//
// Solidity: function setCreditThreshold(uint256 _creditThreshold) returns()
func (_YStrategyPrisma *YStrategyPrismaSession) SetCreditThreshold(_creditThreshold *big.Int) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.SetCreditThreshold(&_YStrategyPrisma.TransactOpts, _creditThreshold)
}

// SetCreditThreshold is a paid mutator transaction binding the contract method 0xfe2508a6.
//
// Solidity: function setCreditThreshold(uint256 _creditThreshold) returns()
func (_YStrategyPrisma *YStrategyPrismaTransactorSession) SetCreditThreshold(_creditThreshold *big.Int) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.SetCreditThreshold(&_YStrategyPrisma.TransactOpts, _creditThreshold)
}

// SetDoHealthCheck is a paid mutator transaction binding the contract method 0xac00ff26.
//
// Solidity: function setDoHealthCheck(bool _doHealthCheck) returns()
func (_YStrategyPrisma *YStrategyPrismaTransactor) SetDoHealthCheck(opts *bind.TransactOpts, _doHealthCheck bool) (*types.Transaction, error) {
	return _YStrategyPrisma.contract.Transact(opts, "setDoHealthCheck", _doHealthCheck)
}

// SetDoHealthCheck is a paid mutator transaction binding the contract method 0xac00ff26.
//
// Solidity: function setDoHealthCheck(bool _doHealthCheck) returns()
func (_YStrategyPrisma *YStrategyPrismaSession) SetDoHealthCheck(_doHealthCheck bool) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.SetDoHealthCheck(&_YStrategyPrisma.TransactOpts, _doHealthCheck)
}

// SetDoHealthCheck is a paid mutator transaction binding the contract method 0xac00ff26.
//
// Solidity: function setDoHealthCheck(bool _doHealthCheck) returns()
func (_YStrategyPrisma *YStrategyPrismaTransactorSession) SetDoHealthCheck(_doHealthCheck bool) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.SetDoHealthCheck(&_YStrategyPrisma.TransactOpts, _doHealthCheck)
}

// SetEmergencyExit is a paid mutator transaction binding the contract method 0xfcf2d0ad.
//
// Solidity: function setEmergencyExit() returns()
func (_YStrategyPrisma *YStrategyPrismaTransactor) SetEmergencyExit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YStrategyPrisma.contract.Transact(opts, "setEmergencyExit")
}

// SetEmergencyExit is a paid mutator transaction binding the contract method 0xfcf2d0ad.
//
// Solidity: function setEmergencyExit() returns()
func (_YStrategyPrisma *YStrategyPrismaSession) SetEmergencyExit() (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.SetEmergencyExit(&_YStrategyPrisma.TransactOpts)
}

// SetEmergencyExit is a paid mutator transaction binding the contract method 0xfcf2d0ad.
//
// Solidity: function setEmergencyExit() returns()
func (_YStrategyPrisma *YStrategyPrismaTransactorSession) SetEmergencyExit() (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.SetEmergencyExit(&_YStrategyPrisma.TransactOpts)
}

// SetForceClaimOnce is a paid mutator transaction binding the contract method 0x70f7e3fb.
//
// Solidity: function setForceClaimOnce(bool _forceClaimOnce) returns()
func (_YStrategyPrisma *YStrategyPrismaTransactor) SetForceClaimOnce(opts *bind.TransactOpts, _forceClaimOnce bool) (*types.Transaction, error) {
	return _YStrategyPrisma.contract.Transact(opts, "setForceClaimOnce", _forceClaimOnce)
}

// SetForceClaimOnce is a paid mutator transaction binding the contract method 0x70f7e3fb.
//
// Solidity: function setForceClaimOnce(bool _forceClaimOnce) returns()
func (_YStrategyPrisma *YStrategyPrismaSession) SetForceClaimOnce(_forceClaimOnce bool) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.SetForceClaimOnce(&_YStrategyPrisma.TransactOpts, _forceClaimOnce)
}

// SetForceClaimOnce is a paid mutator transaction binding the contract method 0x70f7e3fb.
//
// Solidity: function setForceClaimOnce(bool _forceClaimOnce) returns()
func (_YStrategyPrisma *YStrategyPrismaTransactorSession) SetForceClaimOnce(_forceClaimOnce bool) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.SetForceClaimOnce(&_YStrategyPrisma.TransactOpts, _forceClaimOnce)
}

// SetForceHarvestTriggerOnce is a paid mutator transaction binding the contract method 0x0ada4dab.
//
// Solidity: function setForceHarvestTriggerOnce(bool _forceHarvestTriggerOnce) returns()
func (_YStrategyPrisma *YStrategyPrismaTransactor) SetForceHarvestTriggerOnce(opts *bind.TransactOpts, _forceHarvestTriggerOnce bool) (*types.Transaction, error) {
	return _YStrategyPrisma.contract.Transact(opts, "setForceHarvestTriggerOnce", _forceHarvestTriggerOnce)
}

// SetForceHarvestTriggerOnce is a paid mutator transaction binding the contract method 0x0ada4dab.
//
// Solidity: function setForceHarvestTriggerOnce(bool _forceHarvestTriggerOnce) returns()
func (_YStrategyPrisma *YStrategyPrismaSession) SetForceHarvestTriggerOnce(_forceHarvestTriggerOnce bool) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.SetForceHarvestTriggerOnce(&_YStrategyPrisma.TransactOpts, _forceHarvestTriggerOnce)
}

// SetForceHarvestTriggerOnce is a paid mutator transaction binding the contract method 0x0ada4dab.
//
// Solidity: function setForceHarvestTriggerOnce(bool _forceHarvestTriggerOnce) returns()
func (_YStrategyPrisma *YStrategyPrismaTransactorSession) SetForceHarvestTriggerOnce(_forceHarvestTriggerOnce bool) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.SetForceHarvestTriggerOnce(&_YStrategyPrisma.TransactOpts, _forceHarvestTriggerOnce)
}

// SetHarvestTriggerParams is a paid mutator transaction binding the contract method 0xee6497f1.
//
// Solidity: function setHarvestTriggerParams(uint256 _harvestProfitMinInUsdc, uint256 _harvestProfitMaxInUsdc) returns()
func (_YStrategyPrisma *YStrategyPrismaTransactor) SetHarvestTriggerParams(opts *bind.TransactOpts, _harvestProfitMinInUsdc *big.Int, _harvestProfitMaxInUsdc *big.Int) (*types.Transaction, error) {
	return _YStrategyPrisma.contract.Transact(opts, "setHarvestTriggerParams", _harvestProfitMinInUsdc, _harvestProfitMaxInUsdc)
}

// SetHarvestTriggerParams is a paid mutator transaction binding the contract method 0xee6497f1.
//
// Solidity: function setHarvestTriggerParams(uint256 _harvestProfitMinInUsdc, uint256 _harvestProfitMaxInUsdc) returns()
func (_YStrategyPrisma *YStrategyPrismaSession) SetHarvestTriggerParams(_harvestProfitMinInUsdc *big.Int, _harvestProfitMaxInUsdc *big.Int) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.SetHarvestTriggerParams(&_YStrategyPrisma.TransactOpts, _harvestProfitMinInUsdc, _harvestProfitMaxInUsdc)
}

// SetHarvestTriggerParams is a paid mutator transaction binding the contract method 0xee6497f1.
//
// Solidity: function setHarvestTriggerParams(uint256 _harvestProfitMinInUsdc, uint256 _harvestProfitMaxInUsdc) returns()
func (_YStrategyPrisma *YStrategyPrismaTransactorSession) SetHarvestTriggerParams(_harvestProfitMinInUsdc *big.Int, _harvestProfitMaxInUsdc *big.Int) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.SetHarvestTriggerParams(&_YStrategyPrisma.TransactOpts, _harvestProfitMinInUsdc, _harvestProfitMaxInUsdc)
}

// SetHealthCheck is a paid mutator transaction binding the contract method 0x11bc8245.
//
// Solidity: function setHealthCheck(address _healthCheck) returns()
func (_YStrategyPrisma *YStrategyPrismaTransactor) SetHealthCheck(opts *bind.TransactOpts, _healthCheck common.Address) (*types.Transaction, error) {
	return _YStrategyPrisma.contract.Transact(opts, "setHealthCheck", _healthCheck)
}

// SetHealthCheck is a paid mutator transaction binding the contract method 0x11bc8245.
//
// Solidity: function setHealthCheck(address _healthCheck) returns()
func (_YStrategyPrisma *YStrategyPrismaSession) SetHealthCheck(_healthCheck common.Address) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.SetHealthCheck(&_YStrategyPrisma.TransactOpts, _healthCheck)
}

// SetHealthCheck is a paid mutator transaction binding the contract method 0x11bc8245.
//
// Solidity: function setHealthCheck(address _healthCheck) returns()
func (_YStrategyPrisma *YStrategyPrismaTransactorSession) SetHealthCheck(_healthCheck common.Address) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.SetHealthCheck(&_YStrategyPrisma.TransactOpts, _healthCheck)
}

// SetKeeper is a paid mutator transaction binding the contract method 0x748747e6.
//
// Solidity: function setKeeper(address _keeper) returns()
func (_YStrategyPrisma *YStrategyPrismaTransactor) SetKeeper(opts *bind.TransactOpts, _keeper common.Address) (*types.Transaction, error) {
	return _YStrategyPrisma.contract.Transact(opts, "setKeeper", _keeper)
}

// SetKeeper is a paid mutator transaction binding the contract method 0x748747e6.
//
// Solidity: function setKeeper(address _keeper) returns()
func (_YStrategyPrisma *YStrategyPrismaSession) SetKeeper(_keeper common.Address) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.SetKeeper(&_YStrategyPrisma.TransactOpts, _keeper)
}

// SetKeeper is a paid mutator transaction binding the contract method 0x748747e6.
//
// Solidity: function setKeeper(address _keeper) returns()
func (_YStrategyPrisma *YStrategyPrismaTransactorSession) SetKeeper(_keeper common.Address) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.SetKeeper(&_YStrategyPrisma.TransactOpts, _keeper)
}

// SetLocalKeepCrvs is a paid mutator transaction binding the contract method 0x58aea918.
//
// Solidity: function setLocalKeepCrvs(uint256 _keepCrv, uint256 _keepCvx, uint256 _keepYPrisma) returns()
func (_YStrategyPrisma *YStrategyPrismaTransactor) SetLocalKeepCrvs(opts *bind.TransactOpts, _keepCrv *big.Int, _keepCvx *big.Int, _keepYPrisma *big.Int) (*types.Transaction, error) {
	return _YStrategyPrisma.contract.Transact(opts, "setLocalKeepCrvs", _keepCrv, _keepCvx, _keepYPrisma)
}

// SetLocalKeepCrvs is a paid mutator transaction binding the contract method 0x58aea918.
//
// Solidity: function setLocalKeepCrvs(uint256 _keepCrv, uint256 _keepCvx, uint256 _keepYPrisma) returns()
func (_YStrategyPrisma *YStrategyPrismaSession) SetLocalKeepCrvs(_keepCrv *big.Int, _keepCvx *big.Int, _keepYPrisma *big.Int) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.SetLocalKeepCrvs(&_YStrategyPrisma.TransactOpts, _keepCrv, _keepCvx, _keepYPrisma)
}

// SetLocalKeepCrvs is a paid mutator transaction binding the contract method 0x58aea918.
//
// Solidity: function setLocalKeepCrvs(uint256 _keepCrv, uint256 _keepCvx, uint256 _keepYPrisma) returns()
func (_YStrategyPrisma *YStrategyPrismaTransactorSession) SetLocalKeepCrvs(_keepCrv *big.Int, _keepCvx *big.Int, _keepYPrisma *big.Int) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.SetLocalKeepCrvs(&_YStrategyPrisma.TransactOpts, _keepCrv, _keepCvx, _keepYPrisma)
}

// SetMaxReportDelay is a paid mutator transaction binding the contract method 0xf017c92f.
//
// Solidity: function setMaxReportDelay(uint256 _delay) returns()
func (_YStrategyPrisma *YStrategyPrismaTransactor) SetMaxReportDelay(opts *bind.TransactOpts, _delay *big.Int) (*types.Transaction, error) {
	return _YStrategyPrisma.contract.Transact(opts, "setMaxReportDelay", _delay)
}

// SetMaxReportDelay is a paid mutator transaction binding the contract method 0xf017c92f.
//
// Solidity: function setMaxReportDelay(uint256 _delay) returns()
func (_YStrategyPrisma *YStrategyPrismaSession) SetMaxReportDelay(_delay *big.Int) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.SetMaxReportDelay(&_YStrategyPrisma.TransactOpts, _delay)
}

// SetMaxReportDelay is a paid mutator transaction binding the contract method 0xf017c92f.
//
// Solidity: function setMaxReportDelay(uint256 _delay) returns()
func (_YStrategyPrisma *YStrategyPrismaTransactorSession) SetMaxReportDelay(_delay *big.Int) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.SetMaxReportDelay(&_YStrategyPrisma.TransactOpts, _delay)
}

// SetMetadataURI is a paid mutator transaction binding the contract method 0x750521f5.
//
// Solidity: function setMetadataURI(string _metadataURI) returns()
func (_YStrategyPrisma *YStrategyPrismaTransactor) SetMetadataURI(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error) {
	return _YStrategyPrisma.contract.Transact(opts, "setMetadataURI", _metadataURI)
}

// SetMetadataURI is a paid mutator transaction binding the contract method 0x750521f5.
//
// Solidity: function setMetadataURI(string _metadataURI) returns()
func (_YStrategyPrisma *YStrategyPrismaSession) SetMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.SetMetadataURI(&_YStrategyPrisma.TransactOpts, _metadataURI)
}

// SetMetadataURI is a paid mutator transaction binding the contract method 0x750521f5.
//
// Solidity: function setMetadataURI(string _metadataURI) returns()
func (_YStrategyPrisma *YStrategyPrismaTransactorSession) SetMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.SetMetadataURI(&_YStrategyPrisma.TransactOpts, _metadataURI)
}

// SetMinReportDelay is a paid mutator transaction binding the contract method 0x39a172a8.
//
// Solidity: function setMinReportDelay(uint256 _delay) returns()
func (_YStrategyPrisma *YStrategyPrismaTransactor) SetMinReportDelay(opts *bind.TransactOpts, _delay *big.Int) (*types.Transaction, error) {
	return _YStrategyPrisma.contract.Transact(opts, "setMinReportDelay", _delay)
}

// SetMinReportDelay is a paid mutator transaction binding the contract method 0x39a172a8.
//
// Solidity: function setMinReportDelay(uint256 _delay) returns()
func (_YStrategyPrisma *YStrategyPrismaSession) SetMinReportDelay(_delay *big.Int) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.SetMinReportDelay(&_YStrategyPrisma.TransactOpts, _delay)
}

// SetMinReportDelay is a paid mutator transaction binding the contract method 0x39a172a8.
//
// Solidity: function setMinReportDelay(uint256 _delay) returns()
func (_YStrategyPrisma *YStrategyPrismaTransactorSession) SetMinReportDelay(_delay *big.Int) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.SetMinReportDelay(&_YStrategyPrisma.TransactOpts, _delay)
}

// SetRewards is a paid mutator transaction binding the contract method 0xec38a862.
//
// Solidity: function setRewards(address _rewards) returns()
func (_YStrategyPrisma *YStrategyPrismaTransactor) SetRewards(opts *bind.TransactOpts, _rewards common.Address) (*types.Transaction, error) {
	return _YStrategyPrisma.contract.Transact(opts, "setRewards", _rewards)
}

// SetRewards is a paid mutator transaction binding the contract method 0xec38a862.
//
// Solidity: function setRewards(address _rewards) returns()
func (_YStrategyPrisma *YStrategyPrismaSession) SetRewards(_rewards common.Address) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.SetRewards(&_YStrategyPrisma.TransactOpts, _rewards)
}

// SetRewards is a paid mutator transaction binding the contract method 0xec38a862.
//
// Solidity: function setRewards(address _rewards) returns()
func (_YStrategyPrisma *YStrategyPrismaTransactorSession) SetRewards(_rewards common.Address) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.SetRewards(&_YStrategyPrisma.TransactOpts, _rewards)
}

// SetStrategist is a paid mutator transaction binding the contract method 0xc7b9d530.
//
// Solidity: function setStrategist(address _strategist) returns()
func (_YStrategyPrisma *YStrategyPrismaTransactor) SetStrategist(opts *bind.TransactOpts, _strategist common.Address) (*types.Transaction, error) {
	return _YStrategyPrisma.contract.Transact(opts, "setStrategist", _strategist)
}

// SetStrategist is a paid mutator transaction binding the contract method 0xc7b9d530.
//
// Solidity: function setStrategist(address _strategist) returns()
func (_YStrategyPrisma *YStrategyPrismaSession) SetStrategist(_strategist common.Address) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.SetStrategist(&_YStrategyPrisma.TransactOpts, _strategist)
}

// SetStrategist is a paid mutator transaction binding the contract method 0xc7b9d530.
//
// Solidity: function setStrategist(address _strategist) returns()
func (_YStrategyPrisma *YStrategyPrismaTransactorSession) SetStrategist(_strategist common.Address) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.SetStrategist(&_YStrategyPrisma.TransactOpts, _strategist)
}

// SetVoters is a paid mutator transaction binding the contract method 0x97f72f8f.
//
// Solidity: function setVoters(address _curveVoter, address _convexVoter, address _yprismaVoter) returns()
func (_YStrategyPrisma *YStrategyPrismaTransactor) SetVoters(opts *bind.TransactOpts, _curveVoter common.Address, _convexVoter common.Address, _yprismaVoter common.Address) (*types.Transaction, error) {
	return _YStrategyPrisma.contract.Transact(opts, "setVoters", _curveVoter, _convexVoter, _yprismaVoter)
}

// SetVoters is a paid mutator transaction binding the contract method 0x97f72f8f.
//
// Solidity: function setVoters(address _curveVoter, address _convexVoter, address _yprismaVoter) returns()
func (_YStrategyPrisma *YStrategyPrismaSession) SetVoters(_curveVoter common.Address, _convexVoter common.Address, _yprismaVoter common.Address) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.SetVoters(&_YStrategyPrisma.TransactOpts, _curveVoter, _convexVoter, _yprismaVoter)
}

// SetVoters is a paid mutator transaction binding the contract method 0x97f72f8f.
//
// Solidity: function setVoters(address _curveVoter, address _convexVoter, address _yprismaVoter) returns()
func (_YStrategyPrisma *YStrategyPrismaTransactorSession) SetVoters(_curveVoter common.Address, _convexVoter common.Address, _yprismaVoter common.Address) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.SetVoters(&_YStrategyPrisma.TransactOpts, _curveVoter, _convexVoter, _yprismaVoter)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address _token) returns()
func (_YStrategyPrisma *YStrategyPrismaTransactor) Sweep(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _YStrategyPrisma.contract.Transact(opts, "sweep", _token)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address _token) returns()
func (_YStrategyPrisma *YStrategyPrismaSession) Sweep(_token common.Address) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.Sweep(&_YStrategyPrisma.TransactOpts, _token)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address _token) returns()
func (_YStrategyPrisma *YStrategyPrismaTransactorSession) Sweep(_token common.Address) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.Sweep(&_YStrategyPrisma.TransactOpts, _token)
}

// Tend is a paid mutator transaction binding the contract method 0x440368a3.
//
// Solidity: function tend() returns()
func (_YStrategyPrisma *YStrategyPrismaTransactor) Tend(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YStrategyPrisma.contract.Transact(opts, "tend")
}

// Tend is a paid mutator transaction binding the contract method 0x440368a3.
//
// Solidity: function tend() returns()
func (_YStrategyPrisma *YStrategyPrismaSession) Tend() (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.Tend(&_YStrategyPrisma.TransactOpts)
}

// Tend is a paid mutator transaction binding the contract method 0x440368a3.
//
// Solidity: function tend() returns()
func (_YStrategyPrisma *YStrategyPrismaTransactorSession) Tend() (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.Tend(&_YStrategyPrisma.TransactOpts)
}

// UpdateTradeFactory is a paid mutator transaction binding the contract method 0xd8c658c2.
//
// Solidity: function updateTradeFactory(address _newTradeFactory) returns()
func (_YStrategyPrisma *YStrategyPrismaTransactor) UpdateTradeFactory(opts *bind.TransactOpts, _newTradeFactory common.Address) (*types.Transaction, error) {
	return _YStrategyPrisma.contract.Transact(opts, "updateTradeFactory", _newTradeFactory)
}

// UpdateTradeFactory is a paid mutator transaction binding the contract method 0xd8c658c2.
//
// Solidity: function updateTradeFactory(address _newTradeFactory) returns()
func (_YStrategyPrisma *YStrategyPrismaSession) UpdateTradeFactory(_newTradeFactory common.Address) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.UpdateTradeFactory(&_YStrategyPrisma.TransactOpts, _newTradeFactory)
}

// UpdateTradeFactory is a paid mutator transaction binding the contract method 0xd8c658c2.
//
// Solidity: function updateTradeFactory(address _newTradeFactory) returns()
func (_YStrategyPrisma *YStrategyPrismaTransactorSession) UpdateTradeFactory(_newTradeFactory common.Address) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.UpdateTradeFactory(&_YStrategyPrisma.TransactOpts, _newTradeFactory)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amountNeeded) returns(uint256 _loss)
func (_YStrategyPrisma *YStrategyPrismaTransactor) Withdraw(opts *bind.TransactOpts, _amountNeeded *big.Int) (*types.Transaction, error) {
	return _YStrategyPrisma.contract.Transact(opts, "withdraw", _amountNeeded)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amountNeeded) returns(uint256 _loss)
func (_YStrategyPrisma *YStrategyPrismaSession) Withdraw(_amountNeeded *big.Int) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.Withdraw(&_YStrategyPrisma.TransactOpts, _amountNeeded)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amountNeeded) returns(uint256 _loss)
func (_YStrategyPrisma *YStrategyPrismaTransactorSession) Withdraw(_amountNeeded *big.Int) (*types.Transaction, error) {
	return _YStrategyPrisma.Contract.Withdraw(&_YStrategyPrisma.TransactOpts, _amountNeeded)
}

// YStrategyPrismaClonedIterator is returned from FilterCloned and is used to iterate over the raw logs and unpacked data for Cloned events raised by the YStrategyPrisma contract.
type YStrategyPrismaClonedIterator struct {
	Event *YStrategyPrismaCloned // Event containing the contract specifics and raw log

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
func (it *YStrategyPrismaClonedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YStrategyPrismaCloned)
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
		it.Event = new(YStrategyPrismaCloned)
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
func (it *YStrategyPrismaClonedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YStrategyPrismaClonedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YStrategyPrismaCloned represents a Cloned event raised by the YStrategyPrisma contract.
type YStrategyPrismaCloned struct {
	Clone common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCloned is a free log retrieval operation binding the contract event 0x783540fb4221a3238720dc7038937d0d79982bcf895274aa6ad179f82cf0d53c.
//
// Solidity: event Cloned(address indexed clone)
func (_YStrategyPrisma *YStrategyPrismaFilterer) FilterCloned(opts *bind.FilterOpts, clone []common.Address) (*YStrategyPrismaClonedIterator, error) {

	var cloneRule []interface{}
	for _, cloneItem := range clone {
		cloneRule = append(cloneRule, cloneItem)
	}

	logs, sub, err := _YStrategyPrisma.contract.FilterLogs(opts, "Cloned", cloneRule)
	if err != nil {
		return nil, err
	}
	return &YStrategyPrismaClonedIterator{contract: _YStrategyPrisma.contract, event: "Cloned", logs: logs, sub: sub}, nil
}

// WatchCloned is a free log subscription operation binding the contract event 0x783540fb4221a3238720dc7038937d0d79982bcf895274aa6ad179f82cf0d53c.
//
// Solidity: event Cloned(address indexed clone)
func (_YStrategyPrisma *YStrategyPrismaFilterer) WatchCloned(opts *bind.WatchOpts, sink chan<- *YStrategyPrismaCloned, clone []common.Address) (event.Subscription, error) {

	var cloneRule []interface{}
	for _, cloneItem := range clone {
		cloneRule = append(cloneRule, cloneItem)
	}

	logs, sub, err := _YStrategyPrisma.contract.WatchLogs(opts, "Cloned", cloneRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YStrategyPrismaCloned)
				if err := _YStrategyPrisma.contract.UnpackLog(event, "Cloned", log); err != nil {
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
func (_YStrategyPrisma *YStrategyPrismaFilterer) ParseCloned(log types.Log) (*YStrategyPrismaCloned, error) {
	event := new(YStrategyPrismaCloned)
	if err := _YStrategyPrisma.contract.UnpackLog(event, "Cloned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YStrategyPrismaEmergencyExitEnabledIterator is returned from FilterEmergencyExitEnabled and is used to iterate over the raw logs and unpacked data for EmergencyExitEnabled events raised by the YStrategyPrisma contract.
type YStrategyPrismaEmergencyExitEnabledIterator struct {
	Event *YStrategyPrismaEmergencyExitEnabled // Event containing the contract specifics and raw log

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
func (it *YStrategyPrismaEmergencyExitEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YStrategyPrismaEmergencyExitEnabled)
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
		it.Event = new(YStrategyPrismaEmergencyExitEnabled)
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
func (it *YStrategyPrismaEmergencyExitEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YStrategyPrismaEmergencyExitEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YStrategyPrismaEmergencyExitEnabled represents a EmergencyExitEnabled event raised by the YStrategyPrisma contract.
type YStrategyPrismaEmergencyExitEnabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyExitEnabled is a free log retrieval operation binding the contract event 0x97e963041e952738788b9d4871d854d282065b8f90a464928d6528f2e9a4fd0b.
//
// Solidity: event EmergencyExitEnabled()
func (_YStrategyPrisma *YStrategyPrismaFilterer) FilterEmergencyExitEnabled(opts *bind.FilterOpts) (*YStrategyPrismaEmergencyExitEnabledIterator, error) {

	logs, sub, err := _YStrategyPrisma.contract.FilterLogs(opts, "EmergencyExitEnabled")
	if err != nil {
		return nil, err
	}
	return &YStrategyPrismaEmergencyExitEnabledIterator{contract: _YStrategyPrisma.contract, event: "EmergencyExitEnabled", logs: logs, sub: sub}, nil
}

// WatchEmergencyExitEnabled is a free log subscription operation binding the contract event 0x97e963041e952738788b9d4871d854d282065b8f90a464928d6528f2e9a4fd0b.
//
// Solidity: event EmergencyExitEnabled()
func (_YStrategyPrisma *YStrategyPrismaFilterer) WatchEmergencyExitEnabled(opts *bind.WatchOpts, sink chan<- *YStrategyPrismaEmergencyExitEnabled) (event.Subscription, error) {

	logs, sub, err := _YStrategyPrisma.contract.WatchLogs(opts, "EmergencyExitEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YStrategyPrismaEmergencyExitEnabled)
				if err := _YStrategyPrisma.contract.UnpackLog(event, "EmergencyExitEnabled", log); err != nil {
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
func (_YStrategyPrisma *YStrategyPrismaFilterer) ParseEmergencyExitEnabled(log types.Log) (*YStrategyPrismaEmergencyExitEnabled, error) {
	event := new(YStrategyPrismaEmergencyExitEnabled)
	if err := _YStrategyPrisma.contract.UnpackLog(event, "EmergencyExitEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YStrategyPrismaForcedHarvestTriggerIterator is returned from FilterForcedHarvestTrigger and is used to iterate over the raw logs and unpacked data for ForcedHarvestTrigger events raised by the YStrategyPrisma contract.
type YStrategyPrismaForcedHarvestTriggerIterator struct {
	Event *YStrategyPrismaForcedHarvestTrigger // Event containing the contract specifics and raw log

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
func (it *YStrategyPrismaForcedHarvestTriggerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YStrategyPrismaForcedHarvestTrigger)
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
		it.Event = new(YStrategyPrismaForcedHarvestTrigger)
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
func (it *YStrategyPrismaForcedHarvestTriggerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YStrategyPrismaForcedHarvestTriggerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YStrategyPrismaForcedHarvestTrigger represents a ForcedHarvestTrigger event raised by the YStrategyPrisma contract.
type YStrategyPrismaForcedHarvestTrigger struct {
	TriggerState bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterForcedHarvestTrigger is a free log retrieval operation binding the contract event 0x6ad28df1b554fa6cacd46ae82fa811748d53798feeb437ddf234bf3083953319.
//
// Solidity: event ForcedHarvestTrigger(bool triggerState)
func (_YStrategyPrisma *YStrategyPrismaFilterer) FilterForcedHarvestTrigger(opts *bind.FilterOpts) (*YStrategyPrismaForcedHarvestTriggerIterator, error) {

	logs, sub, err := _YStrategyPrisma.contract.FilterLogs(opts, "ForcedHarvestTrigger")
	if err != nil {
		return nil, err
	}
	return &YStrategyPrismaForcedHarvestTriggerIterator{contract: _YStrategyPrisma.contract, event: "ForcedHarvestTrigger", logs: logs, sub: sub}, nil
}

// WatchForcedHarvestTrigger is a free log subscription operation binding the contract event 0x6ad28df1b554fa6cacd46ae82fa811748d53798feeb437ddf234bf3083953319.
//
// Solidity: event ForcedHarvestTrigger(bool triggerState)
func (_YStrategyPrisma *YStrategyPrismaFilterer) WatchForcedHarvestTrigger(opts *bind.WatchOpts, sink chan<- *YStrategyPrismaForcedHarvestTrigger) (event.Subscription, error) {

	logs, sub, err := _YStrategyPrisma.contract.WatchLogs(opts, "ForcedHarvestTrigger")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YStrategyPrismaForcedHarvestTrigger)
				if err := _YStrategyPrisma.contract.UnpackLog(event, "ForcedHarvestTrigger", log); err != nil {
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
func (_YStrategyPrisma *YStrategyPrismaFilterer) ParseForcedHarvestTrigger(log types.Log) (*YStrategyPrismaForcedHarvestTrigger, error) {
	event := new(YStrategyPrismaForcedHarvestTrigger)
	if err := _YStrategyPrisma.contract.UnpackLog(event, "ForcedHarvestTrigger", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YStrategyPrismaHarvestedIterator is returned from FilterHarvested and is used to iterate over the raw logs and unpacked data for Harvested events raised by the YStrategyPrisma contract.
type YStrategyPrismaHarvestedIterator struct {
	Event *YStrategyPrismaHarvested // Event containing the contract specifics and raw log

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
func (it *YStrategyPrismaHarvestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YStrategyPrismaHarvested)
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
		it.Event = new(YStrategyPrismaHarvested)
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
func (it *YStrategyPrismaHarvestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YStrategyPrismaHarvestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YStrategyPrismaHarvested represents a Harvested event raised by the YStrategyPrisma contract.
type YStrategyPrismaHarvested struct {
	Profit          *big.Int
	Loss            *big.Int
	DebtPayment     *big.Int
	DebtOutstanding *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterHarvested is a free log retrieval operation binding the contract event 0x4c0f499ffe6befa0ca7c826b0916cf87bea98de658013e76938489368d60d509.
//
// Solidity: event Harvested(uint256 profit, uint256 loss, uint256 debtPayment, uint256 debtOutstanding)
func (_YStrategyPrisma *YStrategyPrismaFilterer) FilterHarvested(opts *bind.FilterOpts) (*YStrategyPrismaHarvestedIterator, error) {

	logs, sub, err := _YStrategyPrisma.contract.FilterLogs(opts, "Harvested")
	if err != nil {
		return nil, err
	}
	return &YStrategyPrismaHarvestedIterator{contract: _YStrategyPrisma.contract, event: "Harvested", logs: logs, sub: sub}, nil
}

// WatchHarvested is a free log subscription operation binding the contract event 0x4c0f499ffe6befa0ca7c826b0916cf87bea98de658013e76938489368d60d509.
//
// Solidity: event Harvested(uint256 profit, uint256 loss, uint256 debtPayment, uint256 debtOutstanding)
func (_YStrategyPrisma *YStrategyPrismaFilterer) WatchHarvested(opts *bind.WatchOpts, sink chan<- *YStrategyPrismaHarvested) (event.Subscription, error) {

	logs, sub, err := _YStrategyPrisma.contract.WatchLogs(opts, "Harvested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YStrategyPrismaHarvested)
				if err := _YStrategyPrisma.contract.UnpackLog(event, "Harvested", log); err != nil {
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
func (_YStrategyPrisma *YStrategyPrismaFilterer) ParseHarvested(log types.Log) (*YStrategyPrismaHarvested, error) {
	event := new(YStrategyPrismaHarvested)
	if err := _YStrategyPrisma.contract.UnpackLog(event, "Harvested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YStrategyPrismaSetDoHealthCheckIterator is returned from FilterSetDoHealthCheck and is used to iterate over the raw logs and unpacked data for SetDoHealthCheck events raised by the YStrategyPrisma contract.
type YStrategyPrismaSetDoHealthCheckIterator struct {
	Event *YStrategyPrismaSetDoHealthCheck // Event containing the contract specifics and raw log

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
func (it *YStrategyPrismaSetDoHealthCheckIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YStrategyPrismaSetDoHealthCheck)
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
		it.Event = new(YStrategyPrismaSetDoHealthCheck)
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
func (it *YStrategyPrismaSetDoHealthCheckIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YStrategyPrismaSetDoHealthCheckIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YStrategyPrismaSetDoHealthCheck represents a SetDoHealthCheck event raised by the YStrategyPrisma contract.
type YStrategyPrismaSetDoHealthCheck struct {
	Arg0 bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetDoHealthCheck is a free log retrieval operation binding the contract event 0xf769f6bf659bbbdabf212d830720ce893eedc57f25ebb8e44edf5b300618a35b.
//
// Solidity: event SetDoHealthCheck(bool arg0)
func (_YStrategyPrisma *YStrategyPrismaFilterer) FilterSetDoHealthCheck(opts *bind.FilterOpts) (*YStrategyPrismaSetDoHealthCheckIterator, error) {

	logs, sub, err := _YStrategyPrisma.contract.FilterLogs(opts, "SetDoHealthCheck")
	if err != nil {
		return nil, err
	}
	return &YStrategyPrismaSetDoHealthCheckIterator{contract: _YStrategyPrisma.contract, event: "SetDoHealthCheck", logs: logs, sub: sub}, nil
}

// WatchSetDoHealthCheck is a free log subscription operation binding the contract event 0xf769f6bf659bbbdabf212d830720ce893eedc57f25ebb8e44edf5b300618a35b.
//
// Solidity: event SetDoHealthCheck(bool arg0)
func (_YStrategyPrisma *YStrategyPrismaFilterer) WatchSetDoHealthCheck(opts *bind.WatchOpts, sink chan<- *YStrategyPrismaSetDoHealthCheck) (event.Subscription, error) {

	logs, sub, err := _YStrategyPrisma.contract.WatchLogs(opts, "SetDoHealthCheck")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YStrategyPrismaSetDoHealthCheck)
				if err := _YStrategyPrisma.contract.UnpackLog(event, "SetDoHealthCheck", log); err != nil {
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
func (_YStrategyPrisma *YStrategyPrismaFilterer) ParseSetDoHealthCheck(log types.Log) (*YStrategyPrismaSetDoHealthCheck, error) {
	event := new(YStrategyPrismaSetDoHealthCheck)
	if err := _YStrategyPrisma.contract.UnpackLog(event, "SetDoHealthCheck", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YStrategyPrismaSetHealthCheckIterator is returned from FilterSetHealthCheck and is used to iterate over the raw logs and unpacked data for SetHealthCheck events raised by the YStrategyPrisma contract.
type YStrategyPrismaSetHealthCheckIterator struct {
	Event *YStrategyPrismaSetHealthCheck // Event containing the contract specifics and raw log

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
func (it *YStrategyPrismaSetHealthCheckIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YStrategyPrismaSetHealthCheck)
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
		it.Event = new(YStrategyPrismaSetHealthCheck)
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
func (it *YStrategyPrismaSetHealthCheckIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YStrategyPrismaSetHealthCheckIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YStrategyPrismaSetHealthCheck represents a SetHealthCheck event raised by the YStrategyPrisma contract.
type YStrategyPrismaSetHealthCheck struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetHealthCheck is a free log retrieval operation binding the contract event 0xc8db9c35f716b87af1fbb83f03c78646061931269301fd7ba6dcf189b4cdc2fc.
//
// Solidity: event SetHealthCheck(address arg0)
func (_YStrategyPrisma *YStrategyPrismaFilterer) FilterSetHealthCheck(opts *bind.FilterOpts) (*YStrategyPrismaSetHealthCheckIterator, error) {

	logs, sub, err := _YStrategyPrisma.contract.FilterLogs(opts, "SetHealthCheck")
	if err != nil {
		return nil, err
	}
	return &YStrategyPrismaSetHealthCheckIterator{contract: _YStrategyPrisma.contract, event: "SetHealthCheck", logs: logs, sub: sub}, nil
}

// WatchSetHealthCheck is a free log subscription operation binding the contract event 0xc8db9c35f716b87af1fbb83f03c78646061931269301fd7ba6dcf189b4cdc2fc.
//
// Solidity: event SetHealthCheck(address arg0)
func (_YStrategyPrisma *YStrategyPrismaFilterer) WatchSetHealthCheck(opts *bind.WatchOpts, sink chan<- *YStrategyPrismaSetHealthCheck) (event.Subscription, error) {

	logs, sub, err := _YStrategyPrisma.contract.WatchLogs(opts, "SetHealthCheck")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YStrategyPrismaSetHealthCheck)
				if err := _YStrategyPrisma.contract.UnpackLog(event, "SetHealthCheck", log); err != nil {
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
func (_YStrategyPrisma *YStrategyPrismaFilterer) ParseSetHealthCheck(log types.Log) (*YStrategyPrismaSetHealthCheck, error) {
	event := new(YStrategyPrismaSetHealthCheck)
	if err := _YStrategyPrisma.contract.UnpackLog(event, "SetHealthCheck", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YStrategyPrismaUpdatedBaseFeeOracleIterator is returned from FilterUpdatedBaseFeeOracle and is used to iterate over the raw logs and unpacked data for UpdatedBaseFeeOracle events raised by the YStrategyPrisma contract.
type YStrategyPrismaUpdatedBaseFeeOracleIterator struct {
	Event *YStrategyPrismaUpdatedBaseFeeOracle // Event containing the contract specifics and raw log

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
func (it *YStrategyPrismaUpdatedBaseFeeOracleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YStrategyPrismaUpdatedBaseFeeOracle)
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
		it.Event = new(YStrategyPrismaUpdatedBaseFeeOracle)
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
func (it *YStrategyPrismaUpdatedBaseFeeOracleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YStrategyPrismaUpdatedBaseFeeOracleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YStrategyPrismaUpdatedBaseFeeOracle represents a UpdatedBaseFeeOracle event raised by the YStrategyPrisma contract.
type YStrategyPrismaUpdatedBaseFeeOracle struct {
	BaseFeeOracle common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedBaseFeeOracle is a free log retrieval operation binding the contract event 0x711be97287cb9ec921887b9be36e148e1a27c6b158547b22b9704ffc54447a0f.
//
// Solidity: event UpdatedBaseFeeOracle(address baseFeeOracle)
func (_YStrategyPrisma *YStrategyPrismaFilterer) FilterUpdatedBaseFeeOracle(opts *bind.FilterOpts) (*YStrategyPrismaUpdatedBaseFeeOracleIterator, error) {

	logs, sub, err := _YStrategyPrisma.contract.FilterLogs(opts, "UpdatedBaseFeeOracle")
	if err != nil {
		return nil, err
	}
	return &YStrategyPrismaUpdatedBaseFeeOracleIterator{contract: _YStrategyPrisma.contract, event: "UpdatedBaseFeeOracle", logs: logs, sub: sub}, nil
}

// WatchUpdatedBaseFeeOracle is a free log subscription operation binding the contract event 0x711be97287cb9ec921887b9be36e148e1a27c6b158547b22b9704ffc54447a0f.
//
// Solidity: event UpdatedBaseFeeOracle(address baseFeeOracle)
func (_YStrategyPrisma *YStrategyPrismaFilterer) WatchUpdatedBaseFeeOracle(opts *bind.WatchOpts, sink chan<- *YStrategyPrismaUpdatedBaseFeeOracle) (event.Subscription, error) {

	logs, sub, err := _YStrategyPrisma.contract.WatchLogs(opts, "UpdatedBaseFeeOracle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YStrategyPrismaUpdatedBaseFeeOracle)
				if err := _YStrategyPrisma.contract.UnpackLog(event, "UpdatedBaseFeeOracle", log); err != nil {
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
func (_YStrategyPrisma *YStrategyPrismaFilterer) ParseUpdatedBaseFeeOracle(log types.Log) (*YStrategyPrismaUpdatedBaseFeeOracle, error) {
	event := new(YStrategyPrismaUpdatedBaseFeeOracle)
	if err := _YStrategyPrisma.contract.UnpackLog(event, "UpdatedBaseFeeOracle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YStrategyPrismaUpdatedCreditThresholdIterator is returned from FilterUpdatedCreditThreshold and is used to iterate over the raw logs and unpacked data for UpdatedCreditThreshold events raised by the YStrategyPrisma contract.
type YStrategyPrismaUpdatedCreditThresholdIterator struct {
	Event *YStrategyPrismaUpdatedCreditThreshold // Event containing the contract specifics and raw log

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
func (it *YStrategyPrismaUpdatedCreditThresholdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YStrategyPrismaUpdatedCreditThreshold)
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
		it.Event = new(YStrategyPrismaUpdatedCreditThreshold)
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
func (it *YStrategyPrismaUpdatedCreditThresholdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YStrategyPrismaUpdatedCreditThresholdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YStrategyPrismaUpdatedCreditThreshold represents a UpdatedCreditThreshold event raised by the YStrategyPrisma contract.
type YStrategyPrismaUpdatedCreditThreshold struct {
	CreditThreshold *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdatedCreditThreshold is a free log retrieval operation binding the contract event 0xe5ef7832c564a10cbe7b4f1e01ac33a406cb63fcf430a97a9af8616d150af5f3.
//
// Solidity: event UpdatedCreditThreshold(uint256 creditThreshold)
func (_YStrategyPrisma *YStrategyPrismaFilterer) FilterUpdatedCreditThreshold(opts *bind.FilterOpts) (*YStrategyPrismaUpdatedCreditThresholdIterator, error) {

	logs, sub, err := _YStrategyPrisma.contract.FilterLogs(opts, "UpdatedCreditThreshold")
	if err != nil {
		return nil, err
	}
	return &YStrategyPrismaUpdatedCreditThresholdIterator{contract: _YStrategyPrisma.contract, event: "UpdatedCreditThreshold", logs: logs, sub: sub}, nil
}

// WatchUpdatedCreditThreshold is a free log subscription operation binding the contract event 0xe5ef7832c564a10cbe7b4f1e01ac33a406cb63fcf430a97a9af8616d150af5f3.
//
// Solidity: event UpdatedCreditThreshold(uint256 creditThreshold)
func (_YStrategyPrisma *YStrategyPrismaFilterer) WatchUpdatedCreditThreshold(opts *bind.WatchOpts, sink chan<- *YStrategyPrismaUpdatedCreditThreshold) (event.Subscription, error) {

	logs, sub, err := _YStrategyPrisma.contract.WatchLogs(opts, "UpdatedCreditThreshold")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YStrategyPrismaUpdatedCreditThreshold)
				if err := _YStrategyPrisma.contract.UnpackLog(event, "UpdatedCreditThreshold", log); err != nil {
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
func (_YStrategyPrisma *YStrategyPrismaFilterer) ParseUpdatedCreditThreshold(log types.Log) (*YStrategyPrismaUpdatedCreditThreshold, error) {
	event := new(YStrategyPrismaUpdatedCreditThreshold)
	if err := _YStrategyPrisma.contract.UnpackLog(event, "UpdatedCreditThreshold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YStrategyPrismaUpdatedKeeperIterator is returned from FilterUpdatedKeeper and is used to iterate over the raw logs and unpacked data for UpdatedKeeper events raised by the YStrategyPrisma contract.
type YStrategyPrismaUpdatedKeeperIterator struct {
	Event *YStrategyPrismaUpdatedKeeper // Event containing the contract specifics and raw log

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
func (it *YStrategyPrismaUpdatedKeeperIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YStrategyPrismaUpdatedKeeper)
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
		it.Event = new(YStrategyPrismaUpdatedKeeper)
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
func (it *YStrategyPrismaUpdatedKeeperIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YStrategyPrismaUpdatedKeeperIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YStrategyPrismaUpdatedKeeper represents a UpdatedKeeper event raised by the YStrategyPrisma contract.
type YStrategyPrismaUpdatedKeeper struct {
	NewKeeper common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdatedKeeper is a free log retrieval operation binding the contract event 0x2f202ddb4a2e345f6323ed90f8fc8559d770a7abbbeee84dde8aca3351fe7154.
//
// Solidity: event UpdatedKeeper(address newKeeper)
func (_YStrategyPrisma *YStrategyPrismaFilterer) FilterUpdatedKeeper(opts *bind.FilterOpts) (*YStrategyPrismaUpdatedKeeperIterator, error) {

	logs, sub, err := _YStrategyPrisma.contract.FilterLogs(opts, "UpdatedKeeper")
	if err != nil {
		return nil, err
	}
	return &YStrategyPrismaUpdatedKeeperIterator{contract: _YStrategyPrisma.contract, event: "UpdatedKeeper", logs: logs, sub: sub}, nil
}

// WatchUpdatedKeeper is a free log subscription operation binding the contract event 0x2f202ddb4a2e345f6323ed90f8fc8559d770a7abbbeee84dde8aca3351fe7154.
//
// Solidity: event UpdatedKeeper(address newKeeper)
func (_YStrategyPrisma *YStrategyPrismaFilterer) WatchUpdatedKeeper(opts *bind.WatchOpts, sink chan<- *YStrategyPrismaUpdatedKeeper) (event.Subscription, error) {

	logs, sub, err := _YStrategyPrisma.contract.WatchLogs(opts, "UpdatedKeeper")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YStrategyPrismaUpdatedKeeper)
				if err := _YStrategyPrisma.contract.UnpackLog(event, "UpdatedKeeper", log); err != nil {
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
func (_YStrategyPrisma *YStrategyPrismaFilterer) ParseUpdatedKeeper(log types.Log) (*YStrategyPrismaUpdatedKeeper, error) {
	event := new(YStrategyPrismaUpdatedKeeper)
	if err := _YStrategyPrisma.contract.UnpackLog(event, "UpdatedKeeper", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YStrategyPrismaUpdatedMaxReportDelayIterator is returned from FilterUpdatedMaxReportDelay and is used to iterate over the raw logs and unpacked data for UpdatedMaxReportDelay events raised by the YStrategyPrisma contract.
type YStrategyPrismaUpdatedMaxReportDelayIterator struct {
	Event *YStrategyPrismaUpdatedMaxReportDelay // Event containing the contract specifics and raw log

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
func (it *YStrategyPrismaUpdatedMaxReportDelayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YStrategyPrismaUpdatedMaxReportDelay)
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
		it.Event = new(YStrategyPrismaUpdatedMaxReportDelay)
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
func (it *YStrategyPrismaUpdatedMaxReportDelayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YStrategyPrismaUpdatedMaxReportDelayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YStrategyPrismaUpdatedMaxReportDelay represents a UpdatedMaxReportDelay event raised by the YStrategyPrisma contract.
type YStrategyPrismaUpdatedMaxReportDelay struct {
	Delay *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdatedMaxReportDelay is a free log retrieval operation binding the contract event 0x5430e11864ad7aa9775b07d12657fe52df9aa2ba734355bd8ef8747be2c800c5.
//
// Solidity: event UpdatedMaxReportDelay(uint256 delay)
func (_YStrategyPrisma *YStrategyPrismaFilterer) FilterUpdatedMaxReportDelay(opts *bind.FilterOpts) (*YStrategyPrismaUpdatedMaxReportDelayIterator, error) {

	logs, sub, err := _YStrategyPrisma.contract.FilterLogs(opts, "UpdatedMaxReportDelay")
	if err != nil {
		return nil, err
	}
	return &YStrategyPrismaUpdatedMaxReportDelayIterator{contract: _YStrategyPrisma.contract, event: "UpdatedMaxReportDelay", logs: logs, sub: sub}, nil
}

// WatchUpdatedMaxReportDelay is a free log subscription operation binding the contract event 0x5430e11864ad7aa9775b07d12657fe52df9aa2ba734355bd8ef8747be2c800c5.
//
// Solidity: event UpdatedMaxReportDelay(uint256 delay)
func (_YStrategyPrisma *YStrategyPrismaFilterer) WatchUpdatedMaxReportDelay(opts *bind.WatchOpts, sink chan<- *YStrategyPrismaUpdatedMaxReportDelay) (event.Subscription, error) {

	logs, sub, err := _YStrategyPrisma.contract.WatchLogs(opts, "UpdatedMaxReportDelay")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YStrategyPrismaUpdatedMaxReportDelay)
				if err := _YStrategyPrisma.contract.UnpackLog(event, "UpdatedMaxReportDelay", log); err != nil {
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
func (_YStrategyPrisma *YStrategyPrismaFilterer) ParseUpdatedMaxReportDelay(log types.Log) (*YStrategyPrismaUpdatedMaxReportDelay, error) {
	event := new(YStrategyPrismaUpdatedMaxReportDelay)
	if err := _YStrategyPrisma.contract.UnpackLog(event, "UpdatedMaxReportDelay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YStrategyPrismaUpdatedMetadataURIIterator is returned from FilterUpdatedMetadataURI and is used to iterate over the raw logs and unpacked data for UpdatedMetadataURI events raised by the YStrategyPrisma contract.
type YStrategyPrismaUpdatedMetadataURIIterator struct {
	Event *YStrategyPrismaUpdatedMetadataURI // Event containing the contract specifics and raw log

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
func (it *YStrategyPrismaUpdatedMetadataURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YStrategyPrismaUpdatedMetadataURI)
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
		it.Event = new(YStrategyPrismaUpdatedMetadataURI)
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
func (it *YStrategyPrismaUpdatedMetadataURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YStrategyPrismaUpdatedMetadataURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YStrategyPrismaUpdatedMetadataURI represents a UpdatedMetadataURI event raised by the YStrategyPrisma contract.
type YStrategyPrismaUpdatedMetadataURI struct {
	MetadataURI string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatedMetadataURI is a free log retrieval operation binding the contract event 0x300e67d5a415b6d015a471d9c7b95dd58f3e8290af965e84e0f845de2996dda6.
//
// Solidity: event UpdatedMetadataURI(string metadataURI)
func (_YStrategyPrisma *YStrategyPrismaFilterer) FilterUpdatedMetadataURI(opts *bind.FilterOpts) (*YStrategyPrismaUpdatedMetadataURIIterator, error) {

	logs, sub, err := _YStrategyPrisma.contract.FilterLogs(opts, "UpdatedMetadataURI")
	if err != nil {
		return nil, err
	}
	return &YStrategyPrismaUpdatedMetadataURIIterator{contract: _YStrategyPrisma.contract, event: "UpdatedMetadataURI", logs: logs, sub: sub}, nil
}

// WatchUpdatedMetadataURI is a free log subscription operation binding the contract event 0x300e67d5a415b6d015a471d9c7b95dd58f3e8290af965e84e0f845de2996dda6.
//
// Solidity: event UpdatedMetadataURI(string metadataURI)
func (_YStrategyPrisma *YStrategyPrismaFilterer) WatchUpdatedMetadataURI(opts *bind.WatchOpts, sink chan<- *YStrategyPrismaUpdatedMetadataURI) (event.Subscription, error) {

	logs, sub, err := _YStrategyPrisma.contract.WatchLogs(opts, "UpdatedMetadataURI")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YStrategyPrismaUpdatedMetadataURI)
				if err := _YStrategyPrisma.contract.UnpackLog(event, "UpdatedMetadataURI", log); err != nil {
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
func (_YStrategyPrisma *YStrategyPrismaFilterer) ParseUpdatedMetadataURI(log types.Log) (*YStrategyPrismaUpdatedMetadataURI, error) {
	event := new(YStrategyPrismaUpdatedMetadataURI)
	if err := _YStrategyPrisma.contract.UnpackLog(event, "UpdatedMetadataURI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YStrategyPrismaUpdatedMinReportDelayIterator is returned from FilterUpdatedMinReportDelay and is used to iterate over the raw logs and unpacked data for UpdatedMinReportDelay events raised by the YStrategyPrisma contract.
type YStrategyPrismaUpdatedMinReportDelayIterator struct {
	Event *YStrategyPrismaUpdatedMinReportDelay // Event containing the contract specifics and raw log

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
func (it *YStrategyPrismaUpdatedMinReportDelayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YStrategyPrismaUpdatedMinReportDelay)
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
		it.Event = new(YStrategyPrismaUpdatedMinReportDelay)
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
func (it *YStrategyPrismaUpdatedMinReportDelayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YStrategyPrismaUpdatedMinReportDelayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YStrategyPrismaUpdatedMinReportDelay represents a UpdatedMinReportDelay event raised by the YStrategyPrisma contract.
type YStrategyPrismaUpdatedMinReportDelay struct {
	Delay *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdatedMinReportDelay is a free log retrieval operation binding the contract event 0xbb2c369a0355a34b02ab5fce0643150c87e1c8dfe7c918d465591879f57948b1.
//
// Solidity: event UpdatedMinReportDelay(uint256 delay)
func (_YStrategyPrisma *YStrategyPrismaFilterer) FilterUpdatedMinReportDelay(opts *bind.FilterOpts) (*YStrategyPrismaUpdatedMinReportDelayIterator, error) {

	logs, sub, err := _YStrategyPrisma.contract.FilterLogs(opts, "UpdatedMinReportDelay")
	if err != nil {
		return nil, err
	}
	return &YStrategyPrismaUpdatedMinReportDelayIterator{contract: _YStrategyPrisma.contract, event: "UpdatedMinReportDelay", logs: logs, sub: sub}, nil
}

// WatchUpdatedMinReportDelay is a free log subscription operation binding the contract event 0xbb2c369a0355a34b02ab5fce0643150c87e1c8dfe7c918d465591879f57948b1.
//
// Solidity: event UpdatedMinReportDelay(uint256 delay)
func (_YStrategyPrisma *YStrategyPrismaFilterer) WatchUpdatedMinReportDelay(opts *bind.WatchOpts, sink chan<- *YStrategyPrismaUpdatedMinReportDelay) (event.Subscription, error) {

	logs, sub, err := _YStrategyPrisma.contract.WatchLogs(opts, "UpdatedMinReportDelay")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YStrategyPrismaUpdatedMinReportDelay)
				if err := _YStrategyPrisma.contract.UnpackLog(event, "UpdatedMinReportDelay", log); err != nil {
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
func (_YStrategyPrisma *YStrategyPrismaFilterer) ParseUpdatedMinReportDelay(log types.Log) (*YStrategyPrismaUpdatedMinReportDelay, error) {
	event := new(YStrategyPrismaUpdatedMinReportDelay)
	if err := _YStrategyPrisma.contract.UnpackLog(event, "UpdatedMinReportDelay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YStrategyPrismaUpdatedRewardsIterator is returned from FilterUpdatedRewards and is used to iterate over the raw logs and unpacked data for UpdatedRewards events raised by the YStrategyPrisma contract.
type YStrategyPrismaUpdatedRewardsIterator struct {
	Event *YStrategyPrismaUpdatedRewards // Event containing the contract specifics and raw log

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
func (it *YStrategyPrismaUpdatedRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YStrategyPrismaUpdatedRewards)
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
		it.Event = new(YStrategyPrismaUpdatedRewards)
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
func (it *YStrategyPrismaUpdatedRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YStrategyPrismaUpdatedRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YStrategyPrismaUpdatedRewards represents a UpdatedRewards event raised by the YStrategyPrisma contract.
type YStrategyPrismaUpdatedRewards struct {
	Rewards common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdatedRewards is a free log retrieval operation binding the contract event 0xafbb66abf8f3b719799940473a4052a3717cdd8e40fb6c8a3faadab316b1a069.
//
// Solidity: event UpdatedRewards(address rewards)
func (_YStrategyPrisma *YStrategyPrismaFilterer) FilterUpdatedRewards(opts *bind.FilterOpts) (*YStrategyPrismaUpdatedRewardsIterator, error) {

	logs, sub, err := _YStrategyPrisma.contract.FilterLogs(opts, "UpdatedRewards")
	if err != nil {
		return nil, err
	}
	return &YStrategyPrismaUpdatedRewardsIterator{contract: _YStrategyPrisma.contract, event: "UpdatedRewards", logs: logs, sub: sub}, nil
}

// WatchUpdatedRewards is a free log subscription operation binding the contract event 0xafbb66abf8f3b719799940473a4052a3717cdd8e40fb6c8a3faadab316b1a069.
//
// Solidity: event UpdatedRewards(address rewards)
func (_YStrategyPrisma *YStrategyPrismaFilterer) WatchUpdatedRewards(opts *bind.WatchOpts, sink chan<- *YStrategyPrismaUpdatedRewards) (event.Subscription, error) {

	logs, sub, err := _YStrategyPrisma.contract.WatchLogs(opts, "UpdatedRewards")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YStrategyPrismaUpdatedRewards)
				if err := _YStrategyPrisma.contract.UnpackLog(event, "UpdatedRewards", log); err != nil {
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
func (_YStrategyPrisma *YStrategyPrismaFilterer) ParseUpdatedRewards(log types.Log) (*YStrategyPrismaUpdatedRewards, error) {
	event := new(YStrategyPrismaUpdatedRewards)
	if err := _YStrategyPrisma.contract.UnpackLog(event, "UpdatedRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YStrategyPrismaUpdatedStrategistIterator is returned from FilterUpdatedStrategist and is used to iterate over the raw logs and unpacked data for UpdatedStrategist events raised by the YStrategyPrisma contract.
type YStrategyPrismaUpdatedStrategistIterator struct {
	Event *YStrategyPrismaUpdatedStrategist // Event containing the contract specifics and raw log

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
func (it *YStrategyPrismaUpdatedStrategistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YStrategyPrismaUpdatedStrategist)
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
		it.Event = new(YStrategyPrismaUpdatedStrategist)
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
func (it *YStrategyPrismaUpdatedStrategistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YStrategyPrismaUpdatedStrategistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YStrategyPrismaUpdatedStrategist represents a UpdatedStrategist event raised by the YStrategyPrisma contract.
type YStrategyPrismaUpdatedStrategist struct {
	NewStrategist common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedStrategist is a free log retrieval operation binding the contract event 0x352ececae6d7d1e6d26bcf2c549dfd55be1637e9b22dc0cf3b71ddb36097a6b4.
//
// Solidity: event UpdatedStrategist(address newStrategist)
func (_YStrategyPrisma *YStrategyPrismaFilterer) FilterUpdatedStrategist(opts *bind.FilterOpts) (*YStrategyPrismaUpdatedStrategistIterator, error) {

	logs, sub, err := _YStrategyPrisma.contract.FilterLogs(opts, "UpdatedStrategist")
	if err != nil {
		return nil, err
	}
	return &YStrategyPrismaUpdatedStrategistIterator{contract: _YStrategyPrisma.contract, event: "UpdatedStrategist", logs: logs, sub: sub}, nil
}

// WatchUpdatedStrategist is a free log subscription operation binding the contract event 0x352ececae6d7d1e6d26bcf2c549dfd55be1637e9b22dc0cf3b71ddb36097a6b4.
//
// Solidity: event UpdatedStrategist(address newStrategist)
func (_YStrategyPrisma *YStrategyPrismaFilterer) WatchUpdatedStrategist(opts *bind.WatchOpts, sink chan<- *YStrategyPrismaUpdatedStrategist) (event.Subscription, error) {

	logs, sub, err := _YStrategyPrisma.contract.WatchLogs(opts, "UpdatedStrategist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YStrategyPrismaUpdatedStrategist)
				if err := _YStrategyPrisma.contract.UnpackLog(event, "UpdatedStrategist", log); err != nil {
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
func (_YStrategyPrisma *YStrategyPrismaFilterer) ParseUpdatedStrategist(log types.Log) (*YStrategyPrismaUpdatedStrategist, error) {
	event := new(YStrategyPrismaUpdatedStrategist)
	if err := _YStrategyPrisma.contract.UnpackLog(event, "UpdatedStrategist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
