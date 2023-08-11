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

// ConvexBaseStrategyMetaData contains all meta data concerning the ConvexBaseStrategy contract.
var ConvexBaseStrategyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tradeFactory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_harvestProfitMinInUsdc\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_harvestProfitMaxInUsdc\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_booster\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_convexToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"clone\",\"type\":\"address\"}],\"name\":\"Cloned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyExitEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"triggerState\",\"type\":\"bool\"}],\"name\":\"ForcedHarvestTrigger\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"profit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loss\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debtPayment\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debtOutstanding\",\"type\":\"uint256\"}],\"name\":\"Harvested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"SetDoHealthCheck\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"SetHealthCheck\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"baseFeeOracle\",\"type\":\"address\"}],\"name\":\"UpdatedBaseFeeOracle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"creditThreshold\",\"type\":\"uint256\"}],\"name\":\"UpdatedCreditThreshold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newKeeper\",\"type\":\"address\"}],\"name\":\"UpdatedKeeper\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"UpdatedMaxReportDelay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"}],\"name\":\"UpdatedMetadataURI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"UpdatedMinReportDelay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rewards\",\"type\":\"address\"}],\"name\":\"UpdatedRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newStrategist\",\"type\":\"address\"}],\"name\":\"UpdatedStrategist\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"apiVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balanceOfWant\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseFeeOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkEarmark\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimRewards\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimableBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimableProfitInUsdc\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_strategist\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewards\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_keeper\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tradeFactory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_harvestProfitMinInUsdc\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_harvestProfitMaxInUsdc\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_booster\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_convexToken\",\"type\":\"address\"}],\"name\":\"cloneStrategyConvex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newStrategy\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"convexToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"convexVoter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crv\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"curveVoter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cvxWrapper\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegatedAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"doHealthCheck\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uselLocalCRV\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyExit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"estimatedTotalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ethAmount\",\"type\":\"uint256\"}],\"name\":\"ethToWant\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceHarvestTriggerOnce\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"harvest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"harvestProfitMaxInUsdc\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"harvestProfitMinInUsdc\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"callCostinEth\",\"type\":\"uint256\"}],\"name\":\"harvestTrigger\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"healthCheck\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_strategist\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewards\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_keeper\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tradeFactory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_harvestProfitMinInUsdc\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_harvestProfitMaxInUsdc\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_booster\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_convexToken\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isBaseFeeAcceptable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOriginal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"keeper\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localKeepCRV\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localKeepCVX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxReportDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"metadataURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newStrategy\",\"type\":\"address\"}],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minReportDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"needsEarmarkReward\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"needsEarmark\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fraxPid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"id\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_disableTf\",\"type\":\"bool\"}],\"name\":\"removeTradeFactoryPermissions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsContract\",\"outputs\":[{\"internalType\":\"contractIConvexRewards\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewardsTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_baseFeeOracle\",\"type\":\"address\"}],\"name\":\"setBaseFeeOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_claimRewards\",\"type\":\"bool\"}],\"name\":\"setClaimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_creditThreshold\",\"type\":\"uint256\"}],\"name\":\"setCreditThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_doHealthCheck\",\"type\":\"bool\"}],\"name\":\"setDoHealthCheck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setEmergencyExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_forceHarvestTriggerOnce\",\"type\":\"bool\"}],\"name\":\"setForceHarvestTriggerOnce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_harvestProfitMinInUsdc\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_harvestProfitMaxInUsdc\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_checkEarmark\",\"type\":\"bool\"}],\"name\":\"setHarvestTriggerParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_healthCheck\",\"type\":\"address\"}],\"name\":\"setHealthCheck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_keeper\",\"type\":\"address\"}],\"name\":\"setKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_keepCrv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_keepCvx\",\"type\":\"uint256\"}],\"name\":\"setLocalKeepCrvs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_delay\",\"type\":\"uint256\"}],\"name\":\"setMaxReportDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_metadataURI\",\"type\":\"string\"}],\"name\":\"setMetadataURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_delay\",\"type\":\"uint256\"}],\"name\":\"setMinReportDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewards\",\"type\":\"address\"}],\"name\":\"setRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategist\",\"type\":\"address\"}],\"name\":\"setStrategist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_curveVoter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_convexVoter\",\"type\":\"address\"}],\"name\":\"setVoters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"strategist\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"sweep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"callCostInWei\",\"type\":\"uint256\"}],\"name\":\"tendTrigger\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tradeFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newTradeFactory\",\"type\":\"address\"}],\"name\":\"updateTradeFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"contractVaultAPI\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"want\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountNeeded\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_loss\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawToConvexDepositTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ConvexBaseStrategyABI is the input ABI used to generate the binding from.
// Deprecated: Use ConvexBaseStrategyMetaData.ABI instead.
var ConvexBaseStrategyABI = ConvexBaseStrategyMetaData.ABI

// ConvexBaseStrategy is an auto generated Go binding around an Ethereum contract.
type ConvexBaseStrategy struct {
	ConvexBaseStrategyCaller     // Read-only binding to the contract
	ConvexBaseStrategyTransactor // Write-only binding to the contract
	ConvexBaseStrategyFilterer   // Log filterer for contract events
}

// ConvexBaseStrategyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConvexBaseStrategyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConvexBaseStrategyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConvexBaseStrategyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConvexBaseStrategyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConvexBaseStrategyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConvexBaseStrategySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConvexBaseStrategySession struct {
	Contract     *ConvexBaseStrategy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ConvexBaseStrategyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConvexBaseStrategyCallerSession struct {
	Contract *ConvexBaseStrategyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ConvexBaseStrategyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConvexBaseStrategyTransactorSession struct {
	Contract     *ConvexBaseStrategyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ConvexBaseStrategyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConvexBaseStrategyRaw struct {
	Contract *ConvexBaseStrategy // Generic contract binding to access the raw methods on
}

// ConvexBaseStrategyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConvexBaseStrategyCallerRaw struct {
	Contract *ConvexBaseStrategyCaller // Generic read-only contract binding to access the raw methods on
}

// ConvexBaseStrategyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConvexBaseStrategyTransactorRaw struct {
	Contract *ConvexBaseStrategyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConvexBaseStrategy creates a new instance of ConvexBaseStrategy, bound to a specific deployed contract.
func NewConvexBaseStrategy(address common.Address, backend bind.ContractBackend) (*ConvexBaseStrategy, error) {
	contract, err := bindConvexBaseStrategy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ConvexBaseStrategy{ConvexBaseStrategyCaller: ConvexBaseStrategyCaller{contract: contract}, ConvexBaseStrategyTransactor: ConvexBaseStrategyTransactor{contract: contract}, ConvexBaseStrategyFilterer: ConvexBaseStrategyFilterer{contract: contract}}, nil
}

// NewConvexBaseStrategyCaller creates a new read-only instance of ConvexBaseStrategy, bound to a specific deployed contract.
func NewConvexBaseStrategyCaller(address common.Address, caller bind.ContractCaller) (*ConvexBaseStrategyCaller, error) {
	contract, err := bindConvexBaseStrategy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConvexBaseStrategyCaller{contract: contract}, nil
}

// NewConvexBaseStrategyTransactor creates a new write-only instance of ConvexBaseStrategy, bound to a specific deployed contract.
func NewConvexBaseStrategyTransactor(address common.Address, transactor bind.ContractTransactor) (*ConvexBaseStrategyTransactor, error) {
	contract, err := bindConvexBaseStrategy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConvexBaseStrategyTransactor{contract: contract}, nil
}

// NewConvexBaseStrategyFilterer creates a new log filterer instance of ConvexBaseStrategy, bound to a specific deployed contract.
func NewConvexBaseStrategyFilterer(address common.Address, filterer bind.ContractFilterer) (*ConvexBaseStrategyFilterer, error) {
	contract, err := bindConvexBaseStrategy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConvexBaseStrategyFilterer{contract: contract}, nil
}

// bindConvexBaseStrategy binds a generic wrapper to an already deployed contract.
func bindConvexBaseStrategy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ConvexBaseStrategyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConvexBaseStrategy *ConvexBaseStrategyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConvexBaseStrategy.Contract.ConvexBaseStrategyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConvexBaseStrategy *ConvexBaseStrategyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.ConvexBaseStrategyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConvexBaseStrategy *ConvexBaseStrategyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.ConvexBaseStrategyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConvexBaseStrategy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.contract.Transact(opts, method, params...)
}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() pure returns(string)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) ApiVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "apiVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() pure returns(string)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) ApiVersion() (string, error) {
	return _ConvexBaseStrategy.Contract.ApiVersion(&_ConvexBaseStrategy.CallOpts)
}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() pure returns(string)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) ApiVersion() (string, error) {
	return _ConvexBaseStrategy.Contract.ApiVersion(&_ConvexBaseStrategy.CallOpts)
}

// BalanceOfWant is a free data retrieval call binding the contract method 0xc1a3d44c.
//
// Solidity: function balanceOfWant() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) BalanceOfWant(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "balanceOfWant")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfWant is a free data retrieval call binding the contract method 0xc1a3d44c.
//
// Solidity: function balanceOfWant() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) BalanceOfWant() (*big.Int, error) {
	return _ConvexBaseStrategy.Contract.BalanceOfWant(&_ConvexBaseStrategy.CallOpts)
}

// BalanceOfWant is a free data retrieval call binding the contract method 0xc1a3d44c.
//
// Solidity: function balanceOfWant() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) BalanceOfWant() (*big.Int, error) {
	return _ConvexBaseStrategy.Contract.BalanceOfWant(&_ConvexBaseStrategy.CallOpts)
}

// BaseFeeOracle is a free data retrieval call binding the contract method 0x826cddf6.
//
// Solidity: function baseFeeOracle() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) BaseFeeOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "baseFeeOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BaseFeeOracle is a free data retrieval call binding the contract method 0x826cddf6.
//
// Solidity: function baseFeeOracle() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) BaseFeeOracle() (common.Address, error) {
	return _ConvexBaseStrategy.Contract.BaseFeeOracle(&_ConvexBaseStrategy.CallOpts)
}

// BaseFeeOracle is a free data retrieval call binding the contract method 0x826cddf6.
//
// Solidity: function baseFeeOracle() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) BaseFeeOracle() (common.Address, error) {
	return _ConvexBaseStrategy.Contract.BaseFeeOracle(&_ConvexBaseStrategy.CallOpts)
}

// CheckEarmark is a free data retrieval call binding the contract method 0xec2f1050.
//
// Solidity: function checkEarmark() view returns(bool)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) CheckEarmark(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "checkEarmark")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckEarmark is a free data retrieval call binding the contract method 0xec2f1050.
//
// Solidity: function checkEarmark() view returns(bool)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) CheckEarmark() (bool, error) {
	return _ConvexBaseStrategy.Contract.CheckEarmark(&_ConvexBaseStrategy.CallOpts)
}

// CheckEarmark is a free data retrieval call binding the contract method 0xec2f1050.
//
// Solidity: function checkEarmark() view returns(bool)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) CheckEarmark() (bool, error) {
	return _ConvexBaseStrategy.Contract.CheckEarmark(&_ConvexBaseStrategy.CallOpts)
}

// ClaimRewards is a free data retrieval call binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() view returns(bool)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) ClaimRewards(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "claimRewards")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ClaimRewards is a free data retrieval call binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() view returns(bool)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) ClaimRewards() (bool, error) {
	return _ConvexBaseStrategy.Contract.ClaimRewards(&_ConvexBaseStrategy.CallOpts)
}

// ClaimRewards is a free data retrieval call binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() view returns(bool)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) ClaimRewards() (bool, error) {
	return _ConvexBaseStrategy.Contract.ClaimRewards(&_ConvexBaseStrategy.CallOpts)
}

// ClaimableBalance is a free data retrieval call binding the contract method 0xc4f45423.
//
// Solidity: function claimableBalance() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) ClaimableBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "claimableBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimableBalance is a free data retrieval call binding the contract method 0xc4f45423.
//
// Solidity: function claimableBalance() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) ClaimableBalance() (*big.Int, error) {
	return _ConvexBaseStrategy.Contract.ClaimableBalance(&_ConvexBaseStrategy.CallOpts)
}

// ClaimableBalance is a free data retrieval call binding the contract method 0xc4f45423.
//
// Solidity: function claimableBalance() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) ClaimableBalance() (*big.Int, error) {
	return _ConvexBaseStrategy.Contract.ClaimableBalance(&_ConvexBaseStrategy.CallOpts)
}

// ClaimableProfitInUsdc is a free data retrieval call binding the contract method 0xb5762114.
//
// Solidity: function claimableProfitInUsdc() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) ClaimableProfitInUsdc(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "claimableProfitInUsdc")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimableProfitInUsdc is a free data retrieval call binding the contract method 0xb5762114.
//
// Solidity: function claimableProfitInUsdc() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) ClaimableProfitInUsdc() (*big.Int, error) {
	return _ConvexBaseStrategy.Contract.ClaimableProfitInUsdc(&_ConvexBaseStrategy.CallOpts)
}

// ClaimableProfitInUsdc is a free data retrieval call binding the contract method 0xb5762114.
//
// Solidity: function claimableProfitInUsdc() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) ClaimableProfitInUsdc() (*big.Int, error) {
	return _ConvexBaseStrategy.Contract.ClaimableProfitInUsdc(&_ConvexBaseStrategy.CallOpts)
}

// ConvexToken is a free data retrieval call binding the contract method 0xe89133b2.
//
// Solidity: function convexToken() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) ConvexToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "convexToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConvexToken is a free data retrieval call binding the contract method 0xe89133b2.
//
// Solidity: function convexToken() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) ConvexToken() (common.Address, error) {
	return _ConvexBaseStrategy.Contract.ConvexToken(&_ConvexBaseStrategy.CallOpts)
}

// ConvexToken is a free data retrieval call binding the contract method 0xe89133b2.
//
// Solidity: function convexToken() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) ConvexToken() (common.Address, error) {
	return _ConvexBaseStrategy.Contract.ConvexToken(&_ConvexBaseStrategy.CallOpts)
}

// ConvexVoter is a free data retrieval call binding the contract method 0xb084e97b.
//
// Solidity: function convexVoter() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) ConvexVoter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "convexVoter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConvexVoter is a free data retrieval call binding the contract method 0xb084e97b.
//
// Solidity: function convexVoter() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) ConvexVoter() (common.Address, error) {
	return _ConvexBaseStrategy.Contract.ConvexVoter(&_ConvexBaseStrategy.CallOpts)
}

// ConvexVoter is a free data retrieval call binding the contract method 0xb084e97b.
//
// Solidity: function convexVoter() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) ConvexVoter() (common.Address, error) {
	return _ConvexBaseStrategy.Contract.ConvexVoter(&_ConvexBaseStrategy.CallOpts)
}

// CreditThreshold is a free data retrieval call binding the contract method 0xaa5480cf.
//
// Solidity: function creditThreshold() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) CreditThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "creditThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreditThreshold is a free data retrieval call binding the contract method 0xaa5480cf.
//
// Solidity: function creditThreshold() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) CreditThreshold() (*big.Int, error) {
	return _ConvexBaseStrategy.Contract.CreditThreshold(&_ConvexBaseStrategy.CallOpts)
}

// CreditThreshold is a free data retrieval call binding the contract method 0xaa5480cf.
//
// Solidity: function creditThreshold() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) CreditThreshold() (*big.Int, error) {
	return _ConvexBaseStrategy.Contract.CreditThreshold(&_ConvexBaseStrategy.CallOpts)
}

// Crv is a free data retrieval call binding the contract method 0x6a4874a1.
//
// Solidity: function crv() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) Crv(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "crv")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Crv is a free data retrieval call binding the contract method 0x6a4874a1.
//
// Solidity: function crv() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) Crv() (common.Address, error) {
	return _ConvexBaseStrategy.Contract.Crv(&_ConvexBaseStrategy.CallOpts)
}

// Crv is a free data retrieval call binding the contract method 0x6a4874a1.
//
// Solidity: function crv() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) Crv() (common.Address, error) {
	return _ConvexBaseStrategy.Contract.Crv(&_ConvexBaseStrategy.CallOpts)
}

// CurveVoter is a free data retrieval call binding the contract method 0x65d2ebbf.
//
// Solidity: function curveVoter() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) CurveVoter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "curveVoter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurveVoter is a free data retrieval call binding the contract method 0x65d2ebbf.
//
// Solidity: function curveVoter() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) CurveVoter() (common.Address, error) {
	return _ConvexBaseStrategy.Contract.CurveVoter(&_ConvexBaseStrategy.CallOpts)
}

// CurveVoter is a free data retrieval call binding the contract method 0x65d2ebbf.
//
// Solidity: function curveVoter() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) CurveVoter() (common.Address, error) {
	return _ConvexBaseStrategy.Contract.CurveVoter(&_ConvexBaseStrategy.CallOpts)
}

// CvxWrapper is a free data retrieval call binding the contract method 0xe000c3a7.
//
// Solidity: function cvxWrapper() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) CvxWrapper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "cvxWrapper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CvxWrapper is a free data retrieval call binding the contract method 0xe000c3a7.
//
// Solidity: function cvxWrapper() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) CvxWrapper() (common.Address, error) {
	return _ConvexBaseStrategy.Contract.CvxWrapper(&_ConvexBaseStrategy.CallOpts)
}

// CvxWrapper is a free data retrieval call binding the contract method 0xe000c3a7.
//
// Solidity: function cvxWrapper() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) CvxWrapper() (common.Address, error) {
	return _ConvexBaseStrategy.Contract.CvxWrapper(&_ConvexBaseStrategy.CallOpts)
}

// DelegatedAssets is a free data retrieval call binding the contract method 0x8e6350e2.
//
// Solidity: function delegatedAssets() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) DelegatedAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "delegatedAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegatedAssets is a free data retrieval call binding the contract method 0x8e6350e2.
//
// Solidity: function delegatedAssets() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) DelegatedAssets() (*big.Int, error) {
	return _ConvexBaseStrategy.Contract.DelegatedAssets(&_ConvexBaseStrategy.CallOpts)
}

// DelegatedAssets is a free data retrieval call binding the contract method 0x8e6350e2.
//
// Solidity: function delegatedAssets() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) DelegatedAssets() (*big.Int, error) {
	return _ConvexBaseStrategy.Contract.DelegatedAssets(&_ConvexBaseStrategy.CallOpts)
}

// DepositContract is a free data retrieval call binding the contract method 0xe94ad65b.
//
// Solidity: function depositContract() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) DepositContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "depositContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DepositContract is a free data retrieval call binding the contract method 0xe94ad65b.
//
// Solidity: function depositContract() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) DepositContract() (common.Address, error) {
	return _ConvexBaseStrategy.Contract.DepositContract(&_ConvexBaseStrategy.CallOpts)
}

// DepositContract is a free data retrieval call binding the contract method 0xe94ad65b.
//
// Solidity: function depositContract() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) DepositContract() (common.Address, error) {
	return _ConvexBaseStrategy.Contract.DepositContract(&_ConvexBaseStrategy.CallOpts)
}

// DoHealthCheck is a free data retrieval call binding the contract method 0x6718835f.
//
// Solidity: function doHealthCheck() view returns(bool)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) DoHealthCheck(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "doHealthCheck")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DoHealthCheck is a free data retrieval call binding the contract method 0x6718835f.
//
// Solidity: function doHealthCheck() view returns(bool)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) DoHealthCheck() (bool, error) {
	return _ConvexBaseStrategy.Contract.DoHealthCheck(&_ConvexBaseStrategy.CallOpts)
}

// DoHealthCheck is a free data retrieval call binding the contract method 0x6718835f.
//
// Solidity: function doHealthCheck() view returns(bool)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) DoHealthCheck() (bool, error) {
	return _ConvexBaseStrategy.Contract.DoHealthCheck(&_ConvexBaseStrategy.CallOpts)
}



// UselLocalCRV is a free data retrieval call binding the contract method 0x6718835f.
//
// Solidity: function uselLocalCRV() view returns(bool)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) UselLocalCRV(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "uselLocalCRV")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UselLocalCRV is a free data retrieval call binding the contract method 0x6718835f.
//
// Solidity: function uselLocalCRV() view returns(bool)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) UselLocalCRV() (bool, error) {
	return _ConvexBaseStrategy.Contract.UselLocalCRV(&_ConvexBaseStrategy.CallOpts)
}

// UselLocalCRV is a free data retrieval call binding the contract method 0x6718835f.
//
// Solidity: function uselLocalCRV() view returns(bool)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) UselLocalCRV() (bool, error) {
	return _ConvexBaseStrategy.Contract.UselLocalCRV(&_ConvexBaseStrategy.CallOpts)
}

// EmergencyExit is a free data retrieval call binding the contract method 0x5641ec03.
//
// Solidity: function emergencyExit() view returns(bool)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) EmergencyExit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "emergencyExit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EmergencyExit is a free data retrieval call binding the contract method 0x5641ec03.
//
// Solidity: function emergencyExit() view returns(bool)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) EmergencyExit() (bool, error) {
	return _ConvexBaseStrategy.Contract.EmergencyExit(&_ConvexBaseStrategy.CallOpts)
}

// EmergencyExit is a free data retrieval call binding the contract method 0x5641ec03.
//
// Solidity: function emergencyExit() view returns(bool)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) EmergencyExit() (bool, error) {
	return _ConvexBaseStrategy.Contract.EmergencyExit(&_ConvexBaseStrategy.CallOpts)
}

// EstimatedTotalAssets is a free data retrieval call binding the contract method 0xefbb5cb0.
//
// Solidity: function estimatedTotalAssets() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) EstimatedTotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "estimatedTotalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimatedTotalAssets is a free data retrieval call binding the contract method 0xefbb5cb0.
//
// Solidity: function estimatedTotalAssets() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) EstimatedTotalAssets() (*big.Int, error) {
	return _ConvexBaseStrategy.Contract.EstimatedTotalAssets(&_ConvexBaseStrategy.CallOpts)
}

// EstimatedTotalAssets is a free data retrieval call binding the contract method 0xefbb5cb0.
//
// Solidity: function estimatedTotalAssets() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) EstimatedTotalAssets() (*big.Int, error) {
	return _ConvexBaseStrategy.Contract.EstimatedTotalAssets(&_ConvexBaseStrategy.CallOpts)
}

// EthToWant is a free data retrieval call binding the contract method 0x780022a0.
//
// Solidity: function ethToWant(uint256 _ethAmount) view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) EthToWant(opts *bind.CallOpts, _ethAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "ethToWant", _ethAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EthToWant is a free data retrieval call binding the contract method 0x780022a0.
//
// Solidity: function ethToWant(uint256 _ethAmount) view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) EthToWant(_ethAmount *big.Int) (*big.Int, error) {
	return _ConvexBaseStrategy.Contract.EthToWant(&_ConvexBaseStrategy.CallOpts, _ethAmount)
}

// EthToWant is a free data retrieval call binding the contract method 0x780022a0.
//
// Solidity: function ethToWant(uint256 _ethAmount) view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) EthToWant(_ethAmount *big.Int) (*big.Int, error) {
	return _ConvexBaseStrategy.Contract.EthToWant(&_ConvexBaseStrategy.CallOpts, _ethAmount)
}

// ForceHarvestTriggerOnce is a free data retrieval call binding the contract method 0xa763cf5b.
//
// Solidity: function forceHarvestTriggerOnce() view returns(bool)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) ForceHarvestTriggerOnce(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "forceHarvestTriggerOnce")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ForceHarvestTriggerOnce is a free data retrieval call binding the contract method 0xa763cf5b.
//
// Solidity: function forceHarvestTriggerOnce() view returns(bool)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) ForceHarvestTriggerOnce() (bool, error) {
	return _ConvexBaseStrategy.Contract.ForceHarvestTriggerOnce(&_ConvexBaseStrategy.CallOpts)
}

// ForceHarvestTriggerOnce is a free data retrieval call binding the contract method 0xa763cf5b.
//
// Solidity: function forceHarvestTriggerOnce() view returns(bool)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) ForceHarvestTriggerOnce() (bool, error) {
	return _ConvexBaseStrategy.Contract.ForceHarvestTriggerOnce(&_ConvexBaseStrategy.CallOpts)
}

// HarvestProfitMaxInUsdc is a free data retrieval call binding the contract method 0xfa4e2df9.
//
// Solidity: function harvestProfitMaxInUsdc() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) HarvestProfitMaxInUsdc(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "harvestProfitMaxInUsdc")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HarvestProfitMaxInUsdc is a free data retrieval call binding the contract method 0xfa4e2df9.
//
// Solidity: function harvestProfitMaxInUsdc() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) HarvestProfitMaxInUsdc() (*big.Int, error) {
	return _ConvexBaseStrategy.Contract.HarvestProfitMaxInUsdc(&_ConvexBaseStrategy.CallOpts)
}

// HarvestProfitMaxInUsdc is a free data retrieval call binding the contract method 0xfa4e2df9.
//
// Solidity: function harvestProfitMaxInUsdc() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) HarvestProfitMaxInUsdc() (*big.Int, error) {
	return _ConvexBaseStrategy.Contract.HarvestProfitMaxInUsdc(&_ConvexBaseStrategy.CallOpts)
}

// HarvestProfitMinInUsdc is a free data retrieval call binding the contract method 0x090c4922.
//
// Solidity: function harvestProfitMinInUsdc() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) HarvestProfitMinInUsdc(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "harvestProfitMinInUsdc")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HarvestProfitMinInUsdc is a free data retrieval call binding the contract method 0x090c4922.
//
// Solidity: function harvestProfitMinInUsdc() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) HarvestProfitMinInUsdc() (*big.Int, error) {
	return _ConvexBaseStrategy.Contract.HarvestProfitMinInUsdc(&_ConvexBaseStrategy.CallOpts)
}

// HarvestProfitMinInUsdc is a free data retrieval call binding the contract method 0x090c4922.
//
// Solidity: function harvestProfitMinInUsdc() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) HarvestProfitMinInUsdc() (*big.Int, error) {
	return _ConvexBaseStrategy.Contract.HarvestProfitMinInUsdc(&_ConvexBaseStrategy.CallOpts)
}

// HarvestTrigger is a free data retrieval call binding the contract method 0xed882c2b.
//
// Solidity: function harvestTrigger(uint256 callCostinEth) view returns(bool)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) HarvestTrigger(opts *bind.CallOpts, callCostinEth *big.Int) (bool, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "harvestTrigger", callCostinEth)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HarvestTrigger is a free data retrieval call binding the contract method 0xed882c2b.
//
// Solidity: function harvestTrigger(uint256 callCostinEth) view returns(bool)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) HarvestTrigger(callCostinEth *big.Int) (bool, error) {
	return _ConvexBaseStrategy.Contract.HarvestTrigger(&_ConvexBaseStrategy.CallOpts, callCostinEth)
}

// HarvestTrigger is a free data retrieval call binding the contract method 0xed882c2b.
//
// Solidity: function harvestTrigger(uint256 callCostinEth) view returns(bool)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) HarvestTrigger(callCostinEth *big.Int) (bool, error) {
	return _ConvexBaseStrategy.Contract.HarvestTrigger(&_ConvexBaseStrategy.CallOpts, callCostinEth)
}

// HealthCheck is a free data retrieval call binding the contract method 0xb252720b.
//
// Solidity: function healthCheck() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) HealthCheck(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "healthCheck")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HealthCheck is a free data retrieval call binding the contract method 0xb252720b.
//
// Solidity: function healthCheck() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) HealthCheck() (common.Address, error) {
	return _ConvexBaseStrategy.Contract.HealthCheck(&_ConvexBaseStrategy.CallOpts)
}

// HealthCheck is a free data retrieval call binding the contract method 0xb252720b.
//
// Solidity: function healthCheck() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) HealthCheck() (common.Address, error) {
	return _ConvexBaseStrategy.Contract.HealthCheck(&_ConvexBaseStrategy.CallOpts)
}


// -------------------------
// CurveGlobal is a free data retrieval call binding the contract method 0xb252720b.
//
// Solidity: function curveGlobal() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) CurveGlobal(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "curveGlobal")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurveGlobal is a free data retrieval call binding the contract method 0xb252720b.
//
// Solidity: function curveGlobal() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) CurveGlobal() (common.Address, error) {
	return _ConvexBaseStrategy.Contract.CurveGlobal(&_ConvexBaseStrategy.CallOpts)
}

// CurveGlobal is a free data retrieval call binding the contract method 0xb252720b.
//
// Solidity: function curveGlobal() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) CurveGlobal() (common.Address, error) {
	return _ConvexBaseStrategy.Contract.CurveGlobal(&_ConvexBaseStrategy.CallOpts)
}
// -------------------------

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) IsActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "isActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) IsActive() (bool, error) {
	return _ConvexBaseStrategy.Contract.IsActive(&_ConvexBaseStrategy.CallOpts)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) IsActive() (bool, error) {
	return _ConvexBaseStrategy.Contract.IsActive(&_ConvexBaseStrategy.CallOpts)
}

// IsBaseFeeAcceptable is a free data retrieval call binding the contract method 0x95326e2d.
//
// Solidity: function isBaseFeeAcceptable() view returns(bool)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) IsBaseFeeAcceptable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "isBaseFeeAcceptable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBaseFeeAcceptable is a free data retrieval call binding the contract method 0x95326e2d.
//
// Solidity: function isBaseFeeAcceptable() view returns(bool)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) IsBaseFeeAcceptable() (bool, error) {
	return _ConvexBaseStrategy.Contract.IsBaseFeeAcceptable(&_ConvexBaseStrategy.CallOpts)
}

// IsBaseFeeAcceptable is a free data retrieval call binding the contract method 0x95326e2d.
//
// Solidity: function isBaseFeeAcceptable() view returns(bool)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) IsBaseFeeAcceptable() (bool, error) {
	return _ConvexBaseStrategy.Contract.IsBaseFeeAcceptable(&_ConvexBaseStrategy.CallOpts)
}

// IsOriginal is a free data retrieval call binding the contract method 0x6f392ce7.
//
// Solidity: function isOriginal() view returns(bool)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) IsOriginal(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "isOriginal")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOriginal is a free data retrieval call binding the contract method 0x6f392ce7.
//
// Solidity: function isOriginal() view returns(bool)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) IsOriginal() (bool, error) {
	return _ConvexBaseStrategy.Contract.IsOriginal(&_ConvexBaseStrategy.CallOpts)
}

// IsOriginal is a free data retrieval call binding the contract method 0x6f392ce7.
//
// Solidity: function isOriginal() view returns(bool)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) IsOriginal() (bool, error) {
	return _ConvexBaseStrategy.Contract.IsOriginal(&_ConvexBaseStrategy.CallOpts)
}

// Keeper is a free data retrieval call binding the contract method 0xaced1661.
//
// Solidity: function keeper() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) Keeper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "keeper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Keeper is a free data retrieval call binding the contract method 0xaced1661.
//
// Solidity: function keeper() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) Keeper() (common.Address, error) {
	return _ConvexBaseStrategy.Contract.Keeper(&_ConvexBaseStrategy.CallOpts)
}

// Keeper is a free data retrieval call binding the contract method 0xaced1661.
//
// Solidity: function keeper() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) Keeper() (common.Address, error) {
	return _ConvexBaseStrategy.Contract.Keeper(&_ConvexBaseStrategy.CallOpts)
}

// LocalKeepCRV is a free data retrieval call binding the contract method 0x73fd827f.
//
// Solidity: function localKeepCRV() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) LocalKeepCRV(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "localKeepCRV")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LocalKeepCRV is a free data retrieval call binding the contract method 0x73fd827f.
//
// Solidity: function localKeepCRV() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) LocalKeepCRV() (*big.Int, error) {
	return _ConvexBaseStrategy.Contract.LocalKeepCRV(&_ConvexBaseStrategy.CallOpts)
}

// LocalKeepCRV is a free data retrieval call binding the contract method 0x73fd827f.
//
// Solidity: function localKeepCRV() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) LocalKeepCRV() (*big.Int, error) {
	return _ConvexBaseStrategy.Contract.LocalKeepCRV(&_ConvexBaseStrategy.CallOpts)
}



// --------------------------------------------------------------------------------
// LocalCRV is a free data retrieval call binding the contract method 0x73fd827f.
//
// Solidity: function localCRV() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) LocalCRV(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "localCRV")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LocalCRV is a free data retrieval call binding the contract method 0x73fd827f.
//
// Solidity: function localCRV() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) LocalCRV() (*big.Int, error) {
	return _ConvexBaseStrategy.Contract.LocalCRV(&_ConvexBaseStrategy.CallOpts)
}

// LocalCRV is a free data retrieval call binding the contract method 0x73fd827f.
//
// Solidity: function localCRV() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) LocalCRV() (*big.Int, error) {
	return _ConvexBaseStrategy.Contract.LocalCRV(&_ConvexBaseStrategy.CallOpts)
}
// --------------------------------------------------------------------------------

// LocalKeepCVX is a free data retrieval call binding the contract method 0x28f30a4c.
//
// Solidity: function localKeepCVX() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) LocalKeepCVX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "localKeepCVX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LocalKeepCVX is a free data retrieval call binding the contract method 0x28f30a4c.
//
// Solidity: function localKeepCVX() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) LocalKeepCVX() (*big.Int, error) {
	return _ConvexBaseStrategy.Contract.LocalKeepCVX(&_ConvexBaseStrategy.CallOpts)
}

// LocalKeepCVX is a free data retrieval call binding the contract method 0x28f30a4c.
//
// Solidity: function localKeepCVX() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) LocalKeepCVX() (*big.Int, error) {
	return _ConvexBaseStrategy.Contract.LocalKeepCVX(&_ConvexBaseStrategy.CallOpts)
}

// MaxReportDelay is a free data retrieval call binding the contract method 0x28b7ccf7.
//
// Solidity: function maxReportDelay() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) MaxReportDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "maxReportDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxReportDelay is a free data retrieval call binding the contract method 0x28b7ccf7.
//
// Solidity: function maxReportDelay() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) MaxReportDelay() (*big.Int, error) {
	return _ConvexBaseStrategy.Contract.MaxReportDelay(&_ConvexBaseStrategy.CallOpts)
}

// MaxReportDelay is a free data retrieval call binding the contract method 0x28b7ccf7.
//
// Solidity: function maxReportDelay() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) MaxReportDelay() (*big.Int, error) {
	return _ConvexBaseStrategy.Contract.MaxReportDelay(&_ConvexBaseStrategy.CallOpts)
}

// MetadataURI is a free data retrieval call binding the contract method 0x03ee438c.
//
// Solidity: function metadataURI() view returns(string)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) MetadataURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "metadataURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MetadataURI is a free data retrieval call binding the contract method 0x03ee438c.
//
// Solidity: function metadataURI() view returns(string)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) MetadataURI() (string, error) {
	return _ConvexBaseStrategy.Contract.MetadataURI(&_ConvexBaseStrategy.CallOpts)
}

// MetadataURI is a free data retrieval call binding the contract method 0x03ee438c.
//
// Solidity: function metadataURI() view returns(string)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) MetadataURI() (string, error) {
	return _ConvexBaseStrategy.Contract.MetadataURI(&_ConvexBaseStrategy.CallOpts)
}

// MinReportDelay is a free data retrieval call binding the contract method 0x95e80c50.
//
// Solidity: function minReportDelay() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) MinReportDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "minReportDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinReportDelay is a free data retrieval call binding the contract method 0x95e80c50.
//
// Solidity: function minReportDelay() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) MinReportDelay() (*big.Int, error) {
	return _ConvexBaseStrategy.Contract.MinReportDelay(&_ConvexBaseStrategy.CallOpts)
}

// MinReportDelay is a free data retrieval call binding the contract method 0x95e80c50.
//
// Solidity: function minReportDelay() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) MinReportDelay() (*big.Int, error) {
	return _ConvexBaseStrategy.Contract.MinReportDelay(&_ConvexBaseStrategy.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) Name() (string, error) {
	return _ConvexBaseStrategy.Contract.Name(&_ConvexBaseStrategy.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) Name() (string, error) {
	return _ConvexBaseStrategy.Contract.Name(&_ConvexBaseStrategy.CallOpts)
}

// NeedsEarmarkReward is a free data retrieval call binding the contract method 0xf09338df.
//
// Solidity: function needsEarmarkReward() view returns(bool needsEarmark)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) NeedsEarmarkReward(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "needsEarmarkReward")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// NeedsEarmarkReward is a free data retrieval call binding the contract method 0xf09338df.
//
// Solidity: function needsEarmarkReward() view returns(bool needsEarmark)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) NeedsEarmarkReward() (bool, error) {
	return _ConvexBaseStrategy.Contract.NeedsEarmarkReward(&_ConvexBaseStrategy.CallOpts)
}

// NeedsEarmarkReward is a free data retrieval call binding the contract method 0xf09338df.
//
// Solidity: function needsEarmarkReward() view returns(bool needsEarmark)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) NeedsEarmarkReward() (bool, error) {
	return _ConvexBaseStrategy.Contract.NeedsEarmarkReward(&_ConvexBaseStrategy.CallOpts)
}

// Pid is a free data retrieval call binding the contract method 0xf1068454.
//
// Solidity: function pid() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) Pid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "pid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Id is a free data retrieval call binding the contract method 0xf1068454.
//
// Solidity: function id() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) ID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "id")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FraxPid is a free data retrieval call binding the contract method 0x70799be7.
//
// Solidity: function fraxPid() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) FraxPid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "fraxPid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Pid is a free data retrieval call binding the contract method 0xf1068454.
//
// Solidity: function pid() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) Pid() (*big.Int, error) {
	return _ConvexBaseStrategy.Contract.Pid(&_ConvexBaseStrategy.CallOpts)
}

// Pid is a free data retrieval call binding the contract method 0xf1068454.
//
// Solidity: function pid() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) Pid() (*big.Int, error) {
	return _ConvexBaseStrategy.Contract.Pid(&_ConvexBaseStrategy.CallOpts)
}

// ID is a free data retrieval call binding the contract method 0xf1068454.
//
// Solidity: function id() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) ID() (*big.Int, error) {
	return _ConvexBaseStrategy.Contract.ID(&_ConvexBaseStrategy.CallOpts)
}

// ID is a free data retrieval call binding the contract method 0xf1068454.
//
// Solidity: function id() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) ID() (*big.Int, error) {
	return _ConvexBaseStrategy.Contract.ID(&_ConvexBaseStrategy.CallOpts)
}

// FraxPid is a free data retrieval call binding the contract method 0x70799be7.
//
// Solidity: function fraxPid() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) FraxPid() (*big.Int, error) {
	return _ConvexBaseStrategy.Contract.FraxPid(&_ConvexBaseStrategy.CallOpts)
}

// FraxPid is a free data retrieval call binding the contract method 0x70799be7.
//
// Solidity: function fraxPid() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) FraxPid() (*big.Int, error) {
	return _ConvexBaseStrategy.Contract.FraxPid(&_ConvexBaseStrategy.CallOpts)
}


// Rewards is a free data retrieval call binding the contract method 0x9ec5a894.
//
// Solidity: function rewards() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) Rewards(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "rewards")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0x9ec5a894.
//
// Solidity: function rewards() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) Rewards() (common.Address, error) {
	return _ConvexBaseStrategy.Contract.Rewards(&_ConvexBaseStrategy.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x9ec5a894.
//
// Solidity: function rewards() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) Rewards() (common.Address, error) {
	return _ConvexBaseStrategy.Contract.Rewards(&_ConvexBaseStrategy.CallOpts)
}

// RewardsContract is a free data retrieval call binding the contract method 0x220cce97.
//
// Solidity: function rewardsContract() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) RewardsContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "rewardsContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsContract is a free data retrieval call binding the contract method 0x220cce97.
//
// Solidity: function rewardsContract() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) RewardsContract() (common.Address, error) {
	return _ConvexBaseStrategy.Contract.RewardsContract(&_ConvexBaseStrategy.CallOpts)
}

// RewardsContract is a free data retrieval call binding the contract method 0x220cce97.
//
// Solidity: function rewardsContract() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) RewardsContract() (common.Address, error) {
	return _ConvexBaseStrategy.Contract.RewardsContract(&_ConvexBaseStrategy.CallOpts)
}

// RewardsTokens is a free data retrieval call binding the contract method 0xb6d0dcd8.
//
// Solidity: function rewardsTokens(uint256 ) view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) RewardsTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "rewardsTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsTokens is a free data retrieval call binding the contract method 0xb6d0dcd8.
//
// Solidity: function rewardsTokens(uint256 ) view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) RewardsTokens(arg0 *big.Int) (common.Address, error) {
	return _ConvexBaseStrategy.Contract.RewardsTokens(&_ConvexBaseStrategy.CallOpts, arg0)
}

// RewardsTokens is a free data retrieval call binding the contract method 0xb6d0dcd8.
//
// Solidity: function rewardsTokens(uint256 ) view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) RewardsTokens(arg0 *big.Int) (common.Address, error) {
	return _ConvexBaseStrategy.Contract.RewardsTokens(&_ConvexBaseStrategy.CallOpts, arg0)
}

// StakedBalance is a free data retrieval call binding the contract method 0x5b9f0016.
//
// Solidity: function stakedBalance() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) StakedBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "stakedBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakedBalance is a free data retrieval call binding the contract method 0x5b9f0016.
//
// Solidity: function stakedBalance() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) StakedBalance() (*big.Int, error) {
	return _ConvexBaseStrategy.Contract.StakedBalance(&_ConvexBaseStrategy.CallOpts)
}

// StakedBalance is a free data retrieval call binding the contract method 0x5b9f0016.
//
// Solidity: function stakedBalance() view returns(uint256)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) StakedBalance() (*big.Int, error) {
	return _ConvexBaseStrategy.Contract.StakedBalance(&_ConvexBaseStrategy.CallOpts)
}

// Strategist is a free data retrieval call binding the contract method 0x1fe4a686.
//
// Solidity: function strategist() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) Strategist(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "strategist")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Strategist is a free data retrieval call binding the contract method 0x1fe4a686.
//
// Solidity: function strategist() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) Strategist() (common.Address, error) {
	return _ConvexBaseStrategy.Contract.Strategist(&_ConvexBaseStrategy.CallOpts)
}

// Strategist is a free data retrieval call binding the contract method 0x1fe4a686.
//
// Solidity: function strategist() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) Strategist() (common.Address, error) {
	return _ConvexBaseStrategy.Contract.Strategist(&_ConvexBaseStrategy.CallOpts)
}

// TendTrigger is a free data retrieval call binding the contract method 0x650d1880.
//
// Solidity: function tendTrigger(uint256 callCostInWei) view returns(bool)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) TendTrigger(opts *bind.CallOpts, callCostInWei *big.Int) (bool, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "tendTrigger", callCostInWei)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TendTrigger is a free data retrieval call binding the contract method 0x650d1880.
//
// Solidity: function tendTrigger(uint256 callCostInWei) view returns(bool)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) TendTrigger(callCostInWei *big.Int) (bool, error) {
	return _ConvexBaseStrategy.Contract.TendTrigger(&_ConvexBaseStrategy.CallOpts, callCostInWei)
}

// TendTrigger is a free data retrieval call binding the contract method 0x650d1880.
//
// Solidity: function tendTrigger(uint256 callCostInWei) view returns(bool)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) TendTrigger(callCostInWei *big.Int) (bool, error) {
	return _ConvexBaseStrategy.Contract.TendTrigger(&_ConvexBaseStrategy.CallOpts, callCostInWei)
}

// TradeFactory is a free data retrieval call binding the contract method 0xe5e19b4a.
//
// Solidity: function tradeFactory() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) TradeFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "tradeFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TradeFactory is a free data retrieval call binding the contract method 0xe5e19b4a.
//
// Solidity: function tradeFactory() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) TradeFactory() (common.Address, error) {
	return _ConvexBaseStrategy.Contract.TradeFactory(&_ConvexBaseStrategy.CallOpts)
}

// TradeFactory is a free data retrieval call binding the contract method 0xe5e19b4a.
//
// Solidity: function tradeFactory() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) TradeFactory() (common.Address, error) {
	return _ConvexBaseStrategy.Contract.TradeFactory(&_ConvexBaseStrategy.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) Vault() (common.Address, error) {
	return _ConvexBaseStrategy.Contract.Vault(&_ConvexBaseStrategy.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) Vault() (common.Address, error) {
	return _ConvexBaseStrategy.Contract.Vault(&_ConvexBaseStrategy.CallOpts)
}

// Want is a free data retrieval call binding the contract method 0x1f1fcd51.
//
// Solidity: function want() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategyCaller) Want(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexBaseStrategy.contract.Call(opts, &out, "want")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Want is a free data retrieval call binding the contract method 0x1f1fcd51.
//
// Solidity: function want() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) Want() (common.Address, error) {
	return _ConvexBaseStrategy.Contract.Want(&_ConvexBaseStrategy.CallOpts)
}

// Want is a free data retrieval call binding the contract method 0x1f1fcd51.
//
// Solidity: function want() view returns(address)
func (_ConvexBaseStrategy *ConvexBaseStrategyCallerSession) Want() (common.Address, error) {
	return _ConvexBaseStrategy.Contract.Want(&_ConvexBaseStrategy.CallOpts)
}

// CloneStrategyConvex is a paid mutator transaction binding the contract method 0x73e3e9d2.
//
// Solidity: function cloneStrategyConvex(address _vault, address _strategist, address _rewards, address _keeper, address _tradeFactory, uint256 _pid, uint256 _harvestProfitMinInUsdc, uint256 _harvestProfitMaxInUsdc, address _booster, address _convexToken) returns(address newStrategy)
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactor) CloneStrategyConvex(opts *bind.TransactOpts, _vault common.Address, _strategist common.Address, _rewards common.Address, _keeper common.Address, _tradeFactory common.Address, _pid *big.Int, _harvestProfitMinInUsdc *big.Int, _harvestProfitMaxInUsdc *big.Int, _booster common.Address, _convexToken common.Address) (*types.Transaction, error) {
	return _ConvexBaseStrategy.contract.Transact(opts, "cloneStrategyConvex", _vault, _strategist, _rewards, _keeper, _tradeFactory, _pid, _harvestProfitMinInUsdc, _harvestProfitMaxInUsdc, _booster, _convexToken)
}

// CloneStrategyConvex is a paid mutator transaction binding the contract method 0x73e3e9d2.
//
// Solidity: function cloneStrategyConvex(address _vault, address _strategist, address _rewards, address _keeper, address _tradeFactory, uint256 _pid, uint256 _harvestProfitMinInUsdc, uint256 _harvestProfitMaxInUsdc, address _booster, address _convexToken) returns(address newStrategy)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) CloneStrategyConvex(_vault common.Address, _strategist common.Address, _rewards common.Address, _keeper common.Address, _tradeFactory common.Address, _pid *big.Int, _harvestProfitMinInUsdc *big.Int, _harvestProfitMaxInUsdc *big.Int, _booster common.Address, _convexToken common.Address) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.CloneStrategyConvex(&_ConvexBaseStrategy.TransactOpts, _vault, _strategist, _rewards, _keeper, _tradeFactory, _pid, _harvestProfitMinInUsdc, _harvestProfitMaxInUsdc, _booster, _convexToken)
}

// CloneStrategyConvex is a paid mutator transaction binding the contract method 0x73e3e9d2.
//
// Solidity: function cloneStrategyConvex(address _vault, address _strategist, address _rewards, address _keeper, address _tradeFactory, uint256 _pid, uint256 _harvestProfitMinInUsdc, uint256 _harvestProfitMaxInUsdc, address _booster, address _convexToken) returns(address newStrategy)
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactorSession) CloneStrategyConvex(_vault common.Address, _strategist common.Address, _rewards common.Address, _keeper common.Address, _tradeFactory common.Address, _pid *big.Int, _harvestProfitMinInUsdc *big.Int, _harvestProfitMaxInUsdc *big.Int, _booster common.Address, _convexToken common.Address) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.CloneStrategyConvex(&_ConvexBaseStrategy.TransactOpts, _vault, _strategist, _rewards, _keeper, _tradeFactory, _pid, _harvestProfitMinInUsdc, _harvestProfitMaxInUsdc, _booster, _convexToken)
}

// Harvest is a paid mutator transaction binding the contract method 0x4641257d.
//
// Solidity: function harvest() returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactor) Harvest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConvexBaseStrategy.contract.Transact(opts, "harvest")
}

// Harvest is a paid mutator transaction binding the contract method 0x4641257d.
//
// Solidity: function harvest() returns()
func (_ConvexBaseStrategy *ConvexBaseStrategySession) Harvest() (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.Harvest(&_ConvexBaseStrategy.TransactOpts)
}

// Harvest is a paid mutator transaction binding the contract method 0x4641257d.
//
// Solidity: function harvest() returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactorSession) Harvest() (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.Harvest(&_ConvexBaseStrategy.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x51773853.
//
// Solidity: function initialize(address _vault, address _strategist, address _rewards, address _keeper, address _tradeFactory, uint256 _pid, uint256 _harvestProfitMinInUsdc, uint256 _harvestProfitMaxInUsdc, address _booster, address _convexToken) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactor) Initialize(opts *bind.TransactOpts, _vault common.Address, _strategist common.Address, _rewards common.Address, _keeper common.Address, _tradeFactory common.Address, _pid *big.Int, _harvestProfitMinInUsdc *big.Int, _harvestProfitMaxInUsdc *big.Int, _booster common.Address, _convexToken common.Address) (*types.Transaction, error) {
	return _ConvexBaseStrategy.contract.Transact(opts, "initialize", _vault, _strategist, _rewards, _keeper, _tradeFactory, _pid, _harvestProfitMinInUsdc, _harvestProfitMaxInUsdc, _booster, _convexToken)
}

// Initialize is a paid mutator transaction binding the contract method 0x51773853.
//
// Solidity: function initialize(address _vault, address _strategist, address _rewards, address _keeper, address _tradeFactory, uint256 _pid, uint256 _harvestProfitMinInUsdc, uint256 _harvestProfitMaxInUsdc, address _booster, address _convexToken) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategySession) Initialize(_vault common.Address, _strategist common.Address, _rewards common.Address, _keeper common.Address, _tradeFactory common.Address, _pid *big.Int, _harvestProfitMinInUsdc *big.Int, _harvestProfitMaxInUsdc *big.Int, _booster common.Address, _convexToken common.Address) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.Initialize(&_ConvexBaseStrategy.TransactOpts, _vault, _strategist, _rewards, _keeper, _tradeFactory, _pid, _harvestProfitMinInUsdc, _harvestProfitMaxInUsdc, _booster, _convexToken)
}

// Initialize is a paid mutator transaction binding the contract method 0x51773853.
//
// Solidity: function initialize(address _vault, address _strategist, address _rewards, address _keeper, address _tradeFactory, uint256 _pid, uint256 _harvestProfitMinInUsdc, uint256 _harvestProfitMaxInUsdc, address _booster, address _convexToken) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactorSession) Initialize(_vault common.Address, _strategist common.Address, _rewards common.Address, _keeper common.Address, _tradeFactory common.Address, _pid *big.Int, _harvestProfitMinInUsdc *big.Int, _harvestProfitMaxInUsdc *big.Int, _booster common.Address, _convexToken common.Address) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.Initialize(&_ConvexBaseStrategy.TransactOpts, _vault, _strategist, _rewards, _keeper, _tradeFactory, _pid, _harvestProfitMinInUsdc, _harvestProfitMaxInUsdc, _booster, _convexToken)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _newStrategy) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactor) Migrate(opts *bind.TransactOpts, _newStrategy common.Address) (*types.Transaction, error) {
	return _ConvexBaseStrategy.contract.Transact(opts, "migrate", _newStrategy)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _newStrategy) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategySession) Migrate(_newStrategy common.Address) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.Migrate(&_ConvexBaseStrategy.TransactOpts, _newStrategy)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _newStrategy) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactorSession) Migrate(_newStrategy common.Address) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.Migrate(&_ConvexBaseStrategy.TransactOpts, _newStrategy)
}

// RemoveTradeFactoryPermissions is a paid mutator transaction binding the contract method 0xe09575a4.
//
// Solidity: function removeTradeFactoryPermissions(bool _disableTf) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactor) RemoveTradeFactoryPermissions(opts *bind.TransactOpts, _disableTf bool) (*types.Transaction, error) {
	return _ConvexBaseStrategy.contract.Transact(opts, "removeTradeFactoryPermissions", _disableTf)
}

// RemoveTradeFactoryPermissions is a paid mutator transaction binding the contract method 0xe09575a4.
//
// Solidity: function removeTradeFactoryPermissions(bool _disableTf) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategySession) RemoveTradeFactoryPermissions(_disableTf bool) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.RemoveTradeFactoryPermissions(&_ConvexBaseStrategy.TransactOpts, _disableTf)
}

// RemoveTradeFactoryPermissions is a paid mutator transaction binding the contract method 0xe09575a4.
//
// Solidity: function removeTradeFactoryPermissions(bool _disableTf) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactorSession) RemoveTradeFactoryPermissions(_disableTf bool) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.RemoveTradeFactoryPermissions(&_ConvexBaseStrategy.TransactOpts, _disableTf)
}

// SetBaseFeeOracle is a paid mutator transaction binding the contract method 0x9f450b5a.
//
// Solidity: function setBaseFeeOracle(address _baseFeeOracle) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactor) SetBaseFeeOracle(opts *bind.TransactOpts, _baseFeeOracle common.Address) (*types.Transaction, error) {
	return _ConvexBaseStrategy.contract.Transact(opts, "setBaseFeeOracle", _baseFeeOracle)
}

// SetBaseFeeOracle is a paid mutator transaction binding the contract method 0x9f450b5a.
//
// Solidity: function setBaseFeeOracle(address _baseFeeOracle) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategySession) SetBaseFeeOracle(_baseFeeOracle common.Address) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.SetBaseFeeOracle(&_ConvexBaseStrategy.TransactOpts, _baseFeeOracle)
}

// SetBaseFeeOracle is a paid mutator transaction binding the contract method 0x9f450b5a.
//
// Solidity: function setBaseFeeOracle(address _baseFeeOracle) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactorSession) SetBaseFeeOracle(_baseFeeOracle common.Address) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.SetBaseFeeOracle(&_ConvexBaseStrategy.TransactOpts, _baseFeeOracle)
}

// SetClaimRewards is a paid mutator transaction binding the contract method 0xa98f9296.
//
// Solidity: function setClaimRewards(bool _claimRewards) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactor) SetClaimRewards(opts *bind.TransactOpts, _claimRewards bool) (*types.Transaction, error) {
	return _ConvexBaseStrategy.contract.Transact(opts, "setClaimRewards", _claimRewards)
}

// SetClaimRewards is a paid mutator transaction binding the contract method 0xa98f9296.
//
// Solidity: function setClaimRewards(bool _claimRewards) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategySession) SetClaimRewards(_claimRewards bool) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.SetClaimRewards(&_ConvexBaseStrategy.TransactOpts, _claimRewards)
}

// SetClaimRewards is a paid mutator transaction binding the contract method 0xa98f9296.
//
// Solidity: function setClaimRewards(bool _claimRewards) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactorSession) SetClaimRewards(_claimRewards bool) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.SetClaimRewards(&_ConvexBaseStrategy.TransactOpts, _claimRewards)
}

// SetCreditThreshold is a paid mutator transaction binding the contract method 0xfe2508a6.
//
// Solidity: function setCreditThreshold(uint256 _creditThreshold) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactor) SetCreditThreshold(opts *bind.TransactOpts, _creditThreshold *big.Int) (*types.Transaction, error) {
	return _ConvexBaseStrategy.contract.Transact(opts, "setCreditThreshold", _creditThreshold)
}

// SetCreditThreshold is a paid mutator transaction binding the contract method 0xfe2508a6.
//
// Solidity: function setCreditThreshold(uint256 _creditThreshold) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategySession) SetCreditThreshold(_creditThreshold *big.Int) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.SetCreditThreshold(&_ConvexBaseStrategy.TransactOpts, _creditThreshold)
}

// SetCreditThreshold is a paid mutator transaction binding the contract method 0xfe2508a6.
//
// Solidity: function setCreditThreshold(uint256 _creditThreshold) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactorSession) SetCreditThreshold(_creditThreshold *big.Int) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.SetCreditThreshold(&_ConvexBaseStrategy.TransactOpts, _creditThreshold)
}

// SetDoHealthCheck is a paid mutator transaction binding the contract method 0xac00ff26.
//
// Solidity: function setDoHealthCheck(bool _doHealthCheck) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactor) SetDoHealthCheck(opts *bind.TransactOpts, _doHealthCheck bool) (*types.Transaction, error) {
	return _ConvexBaseStrategy.contract.Transact(opts, "setDoHealthCheck", _doHealthCheck)
}

// SetDoHealthCheck is a paid mutator transaction binding the contract method 0xac00ff26.
//
// Solidity: function setDoHealthCheck(bool _doHealthCheck) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategySession) SetDoHealthCheck(_doHealthCheck bool) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.SetDoHealthCheck(&_ConvexBaseStrategy.TransactOpts, _doHealthCheck)
}

// SetDoHealthCheck is a paid mutator transaction binding the contract method 0xac00ff26.
//
// Solidity: function setDoHealthCheck(bool _doHealthCheck) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactorSession) SetDoHealthCheck(_doHealthCheck bool) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.SetDoHealthCheck(&_ConvexBaseStrategy.TransactOpts, _doHealthCheck)
}

// SetEmergencyExit is a paid mutator transaction binding the contract method 0xfcf2d0ad.
//
// Solidity: function setEmergencyExit() returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactor) SetEmergencyExit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConvexBaseStrategy.contract.Transact(opts, "setEmergencyExit")
}

// SetEmergencyExit is a paid mutator transaction binding the contract method 0xfcf2d0ad.
//
// Solidity: function setEmergencyExit() returns()
func (_ConvexBaseStrategy *ConvexBaseStrategySession) SetEmergencyExit() (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.SetEmergencyExit(&_ConvexBaseStrategy.TransactOpts)
}

// SetEmergencyExit is a paid mutator transaction binding the contract method 0xfcf2d0ad.
//
// Solidity: function setEmergencyExit() returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactorSession) SetEmergencyExit() (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.SetEmergencyExit(&_ConvexBaseStrategy.TransactOpts)
}

// SetForceHarvestTriggerOnce is a paid mutator transaction binding the contract method 0x0ada4dab.
//
// Solidity: function setForceHarvestTriggerOnce(bool _forceHarvestTriggerOnce) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactor) SetForceHarvestTriggerOnce(opts *bind.TransactOpts, _forceHarvestTriggerOnce bool) (*types.Transaction, error) {
	return _ConvexBaseStrategy.contract.Transact(opts, "setForceHarvestTriggerOnce", _forceHarvestTriggerOnce)
}

// SetForceHarvestTriggerOnce is a paid mutator transaction binding the contract method 0x0ada4dab.
//
// Solidity: function setForceHarvestTriggerOnce(bool _forceHarvestTriggerOnce) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategySession) SetForceHarvestTriggerOnce(_forceHarvestTriggerOnce bool) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.SetForceHarvestTriggerOnce(&_ConvexBaseStrategy.TransactOpts, _forceHarvestTriggerOnce)
}

// SetForceHarvestTriggerOnce is a paid mutator transaction binding the contract method 0x0ada4dab.
//
// Solidity: function setForceHarvestTriggerOnce(bool _forceHarvestTriggerOnce) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactorSession) SetForceHarvestTriggerOnce(_forceHarvestTriggerOnce bool) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.SetForceHarvestTriggerOnce(&_ConvexBaseStrategy.TransactOpts, _forceHarvestTriggerOnce)
}

// SetHarvestTriggerParams is a paid mutator transaction binding the contract method 0xfe060a63.
//
// Solidity: function setHarvestTriggerParams(uint256 _harvestProfitMinInUsdc, uint256 _harvestProfitMaxInUsdc, bool _checkEarmark) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactor) SetHarvestTriggerParams(opts *bind.TransactOpts, _harvestProfitMinInUsdc *big.Int, _harvestProfitMaxInUsdc *big.Int, _checkEarmark bool) (*types.Transaction, error) {
	return _ConvexBaseStrategy.contract.Transact(opts, "setHarvestTriggerParams", _harvestProfitMinInUsdc, _harvestProfitMaxInUsdc, _checkEarmark)
}

// SetHarvestTriggerParams is a paid mutator transaction binding the contract method 0xfe060a63.
//
// Solidity: function setHarvestTriggerParams(uint256 _harvestProfitMinInUsdc, uint256 _harvestProfitMaxInUsdc, bool _checkEarmark) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategySession) SetHarvestTriggerParams(_harvestProfitMinInUsdc *big.Int, _harvestProfitMaxInUsdc *big.Int, _checkEarmark bool) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.SetHarvestTriggerParams(&_ConvexBaseStrategy.TransactOpts, _harvestProfitMinInUsdc, _harvestProfitMaxInUsdc, _checkEarmark)
}

// SetHarvestTriggerParams is a paid mutator transaction binding the contract method 0xfe060a63.
//
// Solidity: function setHarvestTriggerParams(uint256 _harvestProfitMinInUsdc, uint256 _harvestProfitMaxInUsdc, bool _checkEarmark) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactorSession) SetHarvestTriggerParams(_harvestProfitMinInUsdc *big.Int, _harvestProfitMaxInUsdc *big.Int, _checkEarmark bool) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.SetHarvestTriggerParams(&_ConvexBaseStrategy.TransactOpts, _harvestProfitMinInUsdc, _harvestProfitMaxInUsdc, _checkEarmark)
}

// SetHealthCheck is a paid mutator transaction binding the contract method 0x11bc8245.
//
// Solidity: function setHealthCheck(address _healthCheck) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactor) SetHealthCheck(opts *bind.TransactOpts, _healthCheck common.Address) (*types.Transaction, error) {
	return _ConvexBaseStrategy.contract.Transact(opts, "setHealthCheck", _healthCheck)
}

// SetHealthCheck is a paid mutator transaction binding the contract method 0x11bc8245.
//
// Solidity: function setHealthCheck(address _healthCheck) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategySession) SetHealthCheck(_healthCheck common.Address) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.SetHealthCheck(&_ConvexBaseStrategy.TransactOpts, _healthCheck)
}

// SetHealthCheck is a paid mutator transaction binding the contract method 0x11bc8245.
//
// Solidity: function setHealthCheck(address _healthCheck) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactorSession) SetHealthCheck(_healthCheck common.Address) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.SetHealthCheck(&_ConvexBaseStrategy.TransactOpts, _healthCheck)
}

// SetKeeper is a paid mutator transaction binding the contract method 0x748747e6.
//
// Solidity: function setKeeper(address _keeper) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactor) SetKeeper(opts *bind.TransactOpts, _keeper common.Address) (*types.Transaction, error) {
	return _ConvexBaseStrategy.contract.Transact(opts, "setKeeper", _keeper)
}

// SetKeeper is a paid mutator transaction binding the contract method 0x748747e6.
//
// Solidity: function setKeeper(address _keeper) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategySession) SetKeeper(_keeper common.Address) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.SetKeeper(&_ConvexBaseStrategy.TransactOpts, _keeper)
}

// SetKeeper is a paid mutator transaction binding the contract method 0x748747e6.
//
// Solidity: function setKeeper(address _keeper) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactorSession) SetKeeper(_keeper common.Address) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.SetKeeper(&_ConvexBaseStrategy.TransactOpts, _keeper)
}

// SetLocalKeepCrvs is a paid mutator transaction binding the contract method 0xfc26ab47.
//
// Solidity: function setLocalKeepCrvs(uint256 _keepCrv, uint256 _keepCvx) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactor) SetLocalKeepCrvs(opts *bind.TransactOpts, _keepCrv *big.Int, _keepCvx *big.Int) (*types.Transaction, error) {
	return _ConvexBaseStrategy.contract.Transact(opts, "setLocalKeepCrvs", _keepCrv, _keepCvx)
}

// SetLocalKeepCrvs is a paid mutator transaction binding the contract method 0xfc26ab47.
//
// Solidity: function setLocalKeepCrvs(uint256 _keepCrv, uint256 _keepCvx) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategySession) SetLocalKeepCrvs(_keepCrv *big.Int, _keepCvx *big.Int) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.SetLocalKeepCrvs(&_ConvexBaseStrategy.TransactOpts, _keepCrv, _keepCvx)
}

// SetLocalKeepCrvs is a paid mutator transaction binding the contract method 0xfc26ab47.
//
// Solidity: function setLocalKeepCrvs(uint256 _keepCrv, uint256 _keepCvx) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactorSession) SetLocalKeepCrvs(_keepCrv *big.Int, _keepCvx *big.Int) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.SetLocalKeepCrvs(&_ConvexBaseStrategy.TransactOpts, _keepCrv, _keepCvx)
}

// SetMaxReportDelay is a paid mutator transaction binding the contract method 0xf017c92f.
//
// Solidity: function setMaxReportDelay(uint256 _delay) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactor) SetMaxReportDelay(opts *bind.TransactOpts, _delay *big.Int) (*types.Transaction, error) {
	return _ConvexBaseStrategy.contract.Transact(opts, "setMaxReportDelay", _delay)
}

// SetMaxReportDelay is a paid mutator transaction binding the contract method 0xf017c92f.
//
// Solidity: function setMaxReportDelay(uint256 _delay) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategySession) SetMaxReportDelay(_delay *big.Int) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.SetMaxReportDelay(&_ConvexBaseStrategy.TransactOpts, _delay)
}

// SetMaxReportDelay is a paid mutator transaction binding the contract method 0xf017c92f.
//
// Solidity: function setMaxReportDelay(uint256 _delay) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactorSession) SetMaxReportDelay(_delay *big.Int) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.SetMaxReportDelay(&_ConvexBaseStrategy.TransactOpts, _delay)
}

// SetMetadataURI is a paid mutator transaction binding the contract method 0x750521f5.
//
// Solidity: function setMetadataURI(string _metadataURI) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactor) SetMetadataURI(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error) {
	return _ConvexBaseStrategy.contract.Transact(opts, "setMetadataURI", _metadataURI)
}

// SetMetadataURI is a paid mutator transaction binding the contract method 0x750521f5.
//
// Solidity: function setMetadataURI(string _metadataURI) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategySession) SetMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.SetMetadataURI(&_ConvexBaseStrategy.TransactOpts, _metadataURI)
}

// SetMetadataURI is a paid mutator transaction binding the contract method 0x750521f5.
//
// Solidity: function setMetadataURI(string _metadataURI) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactorSession) SetMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.SetMetadataURI(&_ConvexBaseStrategy.TransactOpts, _metadataURI)
}

// SetMinReportDelay is a paid mutator transaction binding the contract method 0x39a172a8.
//
// Solidity: function setMinReportDelay(uint256 _delay) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactor) SetMinReportDelay(opts *bind.TransactOpts, _delay *big.Int) (*types.Transaction, error) {
	return _ConvexBaseStrategy.contract.Transact(opts, "setMinReportDelay", _delay)
}

// SetMinReportDelay is a paid mutator transaction binding the contract method 0x39a172a8.
//
// Solidity: function setMinReportDelay(uint256 _delay) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategySession) SetMinReportDelay(_delay *big.Int) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.SetMinReportDelay(&_ConvexBaseStrategy.TransactOpts, _delay)
}

// SetMinReportDelay is a paid mutator transaction binding the contract method 0x39a172a8.
//
// Solidity: function setMinReportDelay(uint256 _delay) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactorSession) SetMinReportDelay(_delay *big.Int) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.SetMinReportDelay(&_ConvexBaseStrategy.TransactOpts, _delay)
}

// SetRewards is a paid mutator transaction binding the contract method 0xec38a862.
//
// Solidity: function setRewards(address _rewards) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactor) SetRewards(opts *bind.TransactOpts, _rewards common.Address) (*types.Transaction, error) {
	return _ConvexBaseStrategy.contract.Transact(opts, "setRewards", _rewards)
}

// SetRewards is a paid mutator transaction binding the contract method 0xec38a862.
//
// Solidity: function setRewards(address _rewards) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategySession) SetRewards(_rewards common.Address) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.SetRewards(&_ConvexBaseStrategy.TransactOpts, _rewards)
}

// SetRewards is a paid mutator transaction binding the contract method 0xec38a862.
//
// Solidity: function setRewards(address _rewards) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactorSession) SetRewards(_rewards common.Address) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.SetRewards(&_ConvexBaseStrategy.TransactOpts, _rewards)
}

// SetStrategist is a paid mutator transaction binding the contract method 0xc7b9d530.
//
// Solidity: function setStrategist(address _strategist) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactor) SetStrategist(opts *bind.TransactOpts, _strategist common.Address) (*types.Transaction, error) {
	return _ConvexBaseStrategy.contract.Transact(opts, "setStrategist", _strategist)
}

// SetStrategist is a paid mutator transaction binding the contract method 0xc7b9d530.
//
// Solidity: function setStrategist(address _strategist) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategySession) SetStrategist(_strategist common.Address) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.SetStrategist(&_ConvexBaseStrategy.TransactOpts, _strategist)
}

// SetStrategist is a paid mutator transaction binding the contract method 0xc7b9d530.
//
// Solidity: function setStrategist(address _strategist) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactorSession) SetStrategist(_strategist common.Address) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.SetStrategist(&_ConvexBaseStrategy.TransactOpts, _strategist)
}

// SetVoters is a paid mutator transaction binding the contract method 0x742ca1e2.
//
// Solidity: function setVoters(address _curveVoter, address _convexVoter) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactor) SetVoters(opts *bind.TransactOpts, _curveVoter common.Address, _convexVoter common.Address) (*types.Transaction, error) {
	return _ConvexBaseStrategy.contract.Transact(opts, "setVoters", _curveVoter, _convexVoter)
}

// SetVoters is a paid mutator transaction binding the contract method 0x742ca1e2.
//
// Solidity: function setVoters(address _curveVoter, address _convexVoter) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategySession) SetVoters(_curveVoter common.Address, _convexVoter common.Address) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.SetVoters(&_ConvexBaseStrategy.TransactOpts, _curveVoter, _convexVoter)
}

// SetVoters is a paid mutator transaction binding the contract method 0x742ca1e2.
//
// Solidity: function setVoters(address _curveVoter, address _convexVoter) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactorSession) SetVoters(_curveVoter common.Address, _convexVoter common.Address) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.SetVoters(&_ConvexBaseStrategy.TransactOpts, _curveVoter, _convexVoter)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address _token) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactor) Sweep(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _ConvexBaseStrategy.contract.Transact(opts, "sweep", _token)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address _token) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategySession) Sweep(_token common.Address) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.Sweep(&_ConvexBaseStrategy.TransactOpts, _token)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address _token) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactorSession) Sweep(_token common.Address) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.Sweep(&_ConvexBaseStrategy.TransactOpts, _token)
}

// Tend is a paid mutator transaction binding the contract method 0x440368a3.
//
// Solidity: function tend() returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactor) Tend(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConvexBaseStrategy.contract.Transact(opts, "tend")
}

// Tend is a paid mutator transaction binding the contract method 0x440368a3.
//
// Solidity: function tend() returns()
func (_ConvexBaseStrategy *ConvexBaseStrategySession) Tend() (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.Tend(&_ConvexBaseStrategy.TransactOpts)
}

// Tend is a paid mutator transaction binding the contract method 0x440368a3.
//
// Solidity: function tend() returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactorSession) Tend() (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.Tend(&_ConvexBaseStrategy.TransactOpts)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x3e158b0c.
//
// Solidity: function updateRewards() returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactor) UpdateRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConvexBaseStrategy.contract.Transact(opts, "updateRewards")
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x3e158b0c.
//
// Solidity: function updateRewards() returns()
func (_ConvexBaseStrategy *ConvexBaseStrategySession) UpdateRewards() (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.UpdateRewards(&_ConvexBaseStrategy.TransactOpts)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x3e158b0c.
//
// Solidity: function updateRewards() returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactorSession) UpdateRewards() (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.UpdateRewards(&_ConvexBaseStrategy.TransactOpts)
}

// UpdateTradeFactory is a paid mutator transaction binding the contract method 0xd8c658c2.
//
// Solidity: function updateTradeFactory(address _newTradeFactory) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactor) UpdateTradeFactory(opts *bind.TransactOpts, _newTradeFactory common.Address) (*types.Transaction, error) {
	return _ConvexBaseStrategy.contract.Transact(opts, "updateTradeFactory", _newTradeFactory)
}

// UpdateTradeFactory is a paid mutator transaction binding the contract method 0xd8c658c2.
//
// Solidity: function updateTradeFactory(address _newTradeFactory) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategySession) UpdateTradeFactory(_newTradeFactory common.Address) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.UpdateTradeFactory(&_ConvexBaseStrategy.TransactOpts, _newTradeFactory)
}

// UpdateTradeFactory is a paid mutator transaction binding the contract method 0xd8c658c2.
//
// Solidity: function updateTradeFactory(address _newTradeFactory) returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactorSession) UpdateTradeFactory(_newTradeFactory common.Address) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.UpdateTradeFactory(&_ConvexBaseStrategy.TransactOpts, _newTradeFactory)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amountNeeded) returns(uint256 _loss)
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactor) Withdraw(opts *bind.TransactOpts, _amountNeeded *big.Int) (*types.Transaction, error) {
	return _ConvexBaseStrategy.contract.Transact(opts, "withdraw", _amountNeeded)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amountNeeded) returns(uint256 _loss)
func (_ConvexBaseStrategy *ConvexBaseStrategySession) Withdraw(_amountNeeded *big.Int) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.Withdraw(&_ConvexBaseStrategy.TransactOpts, _amountNeeded)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amountNeeded) returns(uint256 _loss)
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactorSession) Withdraw(_amountNeeded *big.Int) (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.Withdraw(&_ConvexBaseStrategy.TransactOpts, _amountNeeded)
}

// WithdrawToConvexDepositTokens is a paid mutator transaction binding the contract method 0x34659dc5.
//
// Solidity: function withdrawToConvexDepositTokens() returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactor) WithdrawToConvexDepositTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConvexBaseStrategy.contract.Transact(opts, "withdrawToConvexDepositTokens")
}

// WithdrawToConvexDepositTokens is a paid mutator transaction binding the contract method 0x34659dc5.
//
// Solidity: function withdrawToConvexDepositTokens() returns()
func (_ConvexBaseStrategy *ConvexBaseStrategySession) WithdrawToConvexDepositTokens() (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.WithdrawToConvexDepositTokens(&_ConvexBaseStrategy.TransactOpts)
}

// WithdrawToConvexDepositTokens is a paid mutator transaction binding the contract method 0x34659dc5.
//
// Solidity: function withdrawToConvexDepositTokens() returns()
func (_ConvexBaseStrategy *ConvexBaseStrategyTransactorSession) WithdrawToConvexDepositTokens() (*types.Transaction, error) {
	return _ConvexBaseStrategy.Contract.WithdrawToConvexDepositTokens(&_ConvexBaseStrategy.TransactOpts)
}

// ConvexBaseStrategyClonedIterator is returned from FilterCloned and is used to iterate over the raw logs and unpacked data for Cloned events raised by the ConvexBaseStrategy contract.
type ConvexBaseStrategyClonedIterator struct {
	Event *ConvexBaseStrategyCloned // Event containing the contract specifics and raw log

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
func (it *ConvexBaseStrategyClonedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConvexBaseStrategyCloned)
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
		it.Event = new(ConvexBaseStrategyCloned)
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
func (it *ConvexBaseStrategyClonedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConvexBaseStrategyClonedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConvexBaseStrategyCloned represents a Cloned event raised by the ConvexBaseStrategy contract.
type ConvexBaseStrategyCloned struct {
	Clone common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCloned is a free log retrieval operation binding the contract event 0x783540fb4221a3238720dc7038937d0d79982bcf895274aa6ad179f82cf0d53c.
//
// Solidity: event Cloned(address indexed clone)
func (_ConvexBaseStrategy *ConvexBaseStrategyFilterer) FilterCloned(opts *bind.FilterOpts, clone []common.Address) (*ConvexBaseStrategyClonedIterator, error) {

	var cloneRule []interface{}
	for _, cloneItem := range clone {
		cloneRule = append(cloneRule, cloneItem)
	}

	logs, sub, err := _ConvexBaseStrategy.contract.FilterLogs(opts, "Cloned", cloneRule)
	if err != nil {
		return nil, err
	}
	return &ConvexBaseStrategyClonedIterator{contract: _ConvexBaseStrategy.contract, event: "Cloned", logs: logs, sub: sub}, nil
}

// WatchCloned is a free log subscription operation binding the contract event 0x783540fb4221a3238720dc7038937d0d79982bcf895274aa6ad179f82cf0d53c.
//
// Solidity: event Cloned(address indexed clone)
func (_ConvexBaseStrategy *ConvexBaseStrategyFilterer) WatchCloned(opts *bind.WatchOpts, sink chan<- *ConvexBaseStrategyCloned, clone []common.Address) (event.Subscription, error) {

	var cloneRule []interface{}
	for _, cloneItem := range clone {
		cloneRule = append(cloneRule, cloneItem)
	}

	logs, sub, err := _ConvexBaseStrategy.contract.WatchLogs(opts, "Cloned", cloneRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConvexBaseStrategyCloned)
				if err := _ConvexBaseStrategy.contract.UnpackLog(event, "Cloned", log); err != nil {
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
func (_ConvexBaseStrategy *ConvexBaseStrategyFilterer) ParseCloned(log types.Log) (*ConvexBaseStrategyCloned, error) {
	event := new(ConvexBaseStrategyCloned)
	if err := _ConvexBaseStrategy.contract.UnpackLog(event, "Cloned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConvexBaseStrategyEmergencyExitEnabledIterator is returned from FilterEmergencyExitEnabled and is used to iterate over the raw logs and unpacked data for EmergencyExitEnabled events raised by the ConvexBaseStrategy contract.
type ConvexBaseStrategyEmergencyExitEnabledIterator struct {
	Event *ConvexBaseStrategyEmergencyExitEnabled // Event containing the contract specifics and raw log

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
func (it *ConvexBaseStrategyEmergencyExitEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConvexBaseStrategyEmergencyExitEnabled)
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
		it.Event = new(ConvexBaseStrategyEmergencyExitEnabled)
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
func (it *ConvexBaseStrategyEmergencyExitEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConvexBaseStrategyEmergencyExitEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConvexBaseStrategyEmergencyExitEnabled represents a EmergencyExitEnabled event raised by the ConvexBaseStrategy contract.
type ConvexBaseStrategyEmergencyExitEnabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyExitEnabled is a free log retrieval operation binding the contract event 0x97e963041e952738788b9d4871d854d282065b8f90a464928d6528f2e9a4fd0b.
//
// Solidity: event EmergencyExitEnabled()
func (_ConvexBaseStrategy *ConvexBaseStrategyFilterer) FilterEmergencyExitEnabled(opts *bind.FilterOpts) (*ConvexBaseStrategyEmergencyExitEnabledIterator, error) {

	logs, sub, err := _ConvexBaseStrategy.contract.FilterLogs(opts, "EmergencyExitEnabled")
	if err != nil {
		return nil, err
	}
	return &ConvexBaseStrategyEmergencyExitEnabledIterator{contract: _ConvexBaseStrategy.contract, event: "EmergencyExitEnabled", logs: logs, sub: sub}, nil
}

// WatchEmergencyExitEnabled is a free log subscription operation binding the contract event 0x97e963041e952738788b9d4871d854d282065b8f90a464928d6528f2e9a4fd0b.
//
// Solidity: event EmergencyExitEnabled()
func (_ConvexBaseStrategy *ConvexBaseStrategyFilterer) WatchEmergencyExitEnabled(opts *bind.WatchOpts, sink chan<- *ConvexBaseStrategyEmergencyExitEnabled) (event.Subscription, error) {

	logs, sub, err := _ConvexBaseStrategy.contract.WatchLogs(opts, "EmergencyExitEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConvexBaseStrategyEmergencyExitEnabled)
				if err := _ConvexBaseStrategy.contract.UnpackLog(event, "EmergencyExitEnabled", log); err != nil {
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
func (_ConvexBaseStrategy *ConvexBaseStrategyFilterer) ParseEmergencyExitEnabled(log types.Log) (*ConvexBaseStrategyEmergencyExitEnabled, error) {
	event := new(ConvexBaseStrategyEmergencyExitEnabled)
	if err := _ConvexBaseStrategy.contract.UnpackLog(event, "EmergencyExitEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConvexBaseStrategyForcedHarvestTriggerIterator is returned from FilterForcedHarvestTrigger and is used to iterate over the raw logs and unpacked data for ForcedHarvestTrigger events raised by the ConvexBaseStrategy contract.
type ConvexBaseStrategyForcedHarvestTriggerIterator struct {
	Event *ConvexBaseStrategyForcedHarvestTrigger // Event containing the contract specifics and raw log

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
func (it *ConvexBaseStrategyForcedHarvestTriggerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConvexBaseStrategyForcedHarvestTrigger)
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
		it.Event = new(ConvexBaseStrategyForcedHarvestTrigger)
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
func (it *ConvexBaseStrategyForcedHarvestTriggerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConvexBaseStrategyForcedHarvestTriggerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConvexBaseStrategyForcedHarvestTrigger represents a ForcedHarvestTrigger event raised by the ConvexBaseStrategy contract.
type ConvexBaseStrategyForcedHarvestTrigger struct {
	TriggerState bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterForcedHarvestTrigger is a free log retrieval operation binding the contract event 0x6ad28df1b554fa6cacd46ae82fa811748d53798feeb437ddf234bf3083953319.
//
// Solidity: event ForcedHarvestTrigger(bool triggerState)
func (_ConvexBaseStrategy *ConvexBaseStrategyFilterer) FilterForcedHarvestTrigger(opts *bind.FilterOpts) (*ConvexBaseStrategyForcedHarvestTriggerIterator, error) {

	logs, sub, err := _ConvexBaseStrategy.contract.FilterLogs(opts, "ForcedHarvestTrigger")
	if err != nil {
		return nil, err
	}
	return &ConvexBaseStrategyForcedHarvestTriggerIterator{contract: _ConvexBaseStrategy.contract, event: "ForcedHarvestTrigger", logs: logs, sub: sub}, nil
}

// WatchForcedHarvestTrigger is a free log subscription operation binding the contract event 0x6ad28df1b554fa6cacd46ae82fa811748d53798feeb437ddf234bf3083953319.
//
// Solidity: event ForcedHarvestTrigger(bool triggerState)
func (_ConvexBaseStrategy *ConvexBaseStrategyFilterer) WatchForcedHarvestTrigger(opts *bind.WatchOpts, sink chan<- *ConvexBaseStrategyForcedHarvestTrigger) (event.Subscription, error) {

	logs, sub, err := _ConvexBaseStrategy.contract.WatchLogs(opts, "ForcedHarvestTrigger")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConvexBaseStrategyForcedHarvestTrigger)
				if err := _ConvexBaseStrategy.contract.UnpackLog(event, "ForcedHarvestTrigger", log); err != nil {
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
func (_ConvexBaseStrategy *ConvexBaseStrategyFilterer) ParseForcedHarvestTrigger(log types.Log) (*ConvexBaseStrategyForcedHarvestTrigger, error) {
	event := new(ConvexBaseStrategyForcedHarvestTrigger)
	if err := _ConvexBaseStrategy.contract.UnpackLog(event, "ForcedHarvestTrigger", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConvexBaseStrategyHarvestedIterator is returned from FilterHarvested and is used to iterate over the raw logs and unpacked data for Harvested events raised by the ConvexBaseStrategy contract.
type ConvexBaseStrategyHarvestedIterator struct {
	Event *ConvexBaseStrategyHarvested // Event containing the contract specifics and raw log

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
func (it *ConvexBaseStrategyHarvestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConvexBaseStrategyHarvested)
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
		it.Event = new(ConvexBaseStrategyHarvested)
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
func (it *ConvexBaseStrategyHarvestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConvexBaseStrategyHarvestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConvexBaseStrategyHarvested represents a Harvested event raised by the ConvexBaseStrategy contract.
type ConvexBaseStrategyHarvested struct {
	Profit          *big.Int
	Loss            *big.Int
	DebtPayment     *big.Int
	DebtOutstanding *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterHarvested is a free log retrieval operation binding the contract event 0x4c0f499ffe6befa0ca7c826b0916cf87bea98de658013e76938489368d60d509.
//
// Solidity: event Harvested(uint256 profit, uint256 loss, uint256 debtPayment, uint256 debtOutstanding)
func (_ConvexBaseStrategy *ConvexBaseStrategyFilterer) FilterHarvested(opts *bind.FilterOpts) (*ConvexBaseStrategyHarvestedIterator, error) {

	logs, sub, err := _ConvexBaseStrategy.contract.FilterLogs(opts, "Harvested")
	if err != nil {
		return nil, err
	}
	return &ConvexBaseStrategyHarvestedIterator{contract: _ConvexBaseStrategy.contract, event: "Harvested", logs: logs, sub: sub}, nil
}

// WatchHarvested is a free log subscription operation binding the contract event 0x4c0f499ffe6befa0ca7c826b0916cf87bea98de658013e76938489368d60d509.
//
// Solidity: event Harvested(uint256 profit, uint256 loss, uint256 debtPayment, uint256 debtOutstanding)
func (_ConvexBaseStrategy *ConvexBaseStrategyFilterer) WatchHarvested(opts *bind.WatchOpts, sink chan<- *ConvexBaseStrategyHarvested) (event.Subscription, error) {

	logs, sub, err := _ConvexBaseStrategy.contract.WatchLogs(opts, "Harvested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConvexBaseStrategyHarvested)
				if err := _ConvexBaseStrategy.contract.UnpackLog(event, "Harvested", log); err != nil {
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
func (_ConvexBaseStrategy *ConvexBaseStrategyFilterer) ParseHarvested(log types.Log) (*ConvexBaseStrategyHarvested, error) {
	event := new(ConvexBaseStrategyHarvested)
	if err := _ConvexBaseStrategy.contract.UnpackLog(event, "Harvested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConvexBaseStrategySetDoHealthCheckIterator is returned from FilterSetDoHealthCheck and is used to iterate over the raw logs and unpacked data for SetDoHealthCheck events raised by the ConvexBaseStrategy contract.
type ConvexBaseStrategySetDoHealthCheckIterator struct {
	Event *ConvexBaseStrategySetDoHealthCheck // Event containing the contract specifics and raw log

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
func (it *ConvexBaseStrategySetDoHealthCheckIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConvexBaseStrategySetDoHealthCheck)
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
		it.Event = new(ConvexBaseStrategySetDoHealthCheck)
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
func (it *ConvexBaseStrategySetDoHealthCheckIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConvexBaseStrategySetDoHealthCheckIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConvexBaseStrategySetDoHealthCheck represents a SetDoHealthCheck event raised by the ConvexBaseStrategy contract.
type ConvexBaseStrategySetDoHealthCheck struct {
	Arg0 bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetDoHealthCheck is a free log retrieval operation binding the contract event 0xf769f6bf659bbbdabf212d830720ce893eedc57f25ebb8e44edf5b300618a35b.
//
// Solidity: event SetDoHealthCheck(bool arg0)
func (_ConvexBaseStrategy *ConvexBaseStrategyFilterer) FilterSetDoHealthCheck(opts *bind.FilterOpts) (*ConvexBaseStrategySetDoHealthCheckIterator, error) {

	logs, sub, err := _ConvexBaseStrategy.contract.FilterLogs(opts, "SetDoHealthCheck")
	if err != nil {
		return nil, err
	}
	return &ConvexBaseStrategySetDoHealthCheckIterator{contract: _ConvexBaseStrategy.contract, event: "SetDoHealthCheck", logs: logs, sub: sub}, nil
}

// WatchSetDoHealthCheck is a free log subscription operation binding the contract event 0xf769f6bf659bbbdabf212d830720ce893eedc57f25ebb8e44edf5b300618a35b.
//
// Solidity: event SetDoHealthCheck(bool arg0)
func (_ConvexBaseStrategy *ConvexBaseStrategyFilterer) WatchSetDoHealthCheck(opts *bind.WatchOpts, sink chan<- *ConvexBaseStrategySetDoHealthCheck) (event.Subscription, error) {

	logs, sub, err := _ConvexBaseStrategy.contract.WatchLogs(opts, "SetDoHealthCheck")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConvexBaseStrategySetDoHealthCheck)
				if err := _ConvexBaseStrategy.contract.UnpackLog(event, "SetDoHealthCheck", log); err != nil {
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
func (_ConvexBaseStrategy *ConvexBaseStrategyFilterer) ParseSetDoHealthCheck(log types.Log) (*ConvexBaseStrategySetDoHealthCheck, error) {
	event := new(ConvexBaseStrategySetDoHealthCheck)
	if err := _ConvexBaseStrategy.contract.UnpackLog(event, "SetDoHealthCheck", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConvexBaseStrategySetHealthCheckIterator is returned from FilterSetHealthCheck and is used to iterate over the raw logs and unpacked data for SetHealthCheck events raised by the ConvexBaseStrategy contract.
type ConvexBaseStrategySetHealthCheckIterator struct {
	Event *ConvexBaseStrategySetHealthCheck // Event containing the contract specifics and raw log

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
func (it *ConvexBaseStrategySetHealthCheckIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConvexBaseStrategySetHealthCheck)
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
		it.Event = new(ConvexBaseStrategySetHealthCheck)
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
func (it *ConvexBaseStrategySetHealthCheckIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConvexBaseStrategySetHealthCheckIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConvexBaseStrategySetHealthCheck represents a SetHealthCheck event raised by the ConvexBaseStrategy contract.
type ConvexBaseStrategySetHealthCheck struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetHealthCheck is a free log retrieval operation binding the contract event 0xc8db9c35f716b87af1fbb83f03c78646061931269301fd7ba6dcf189b4cdc2fc.
//
// Solidity: event SetHealthCheck(address arg0)
func (_ConvexBaseStrategy *ConvexBaseStrategyFilterer) FilterSetHealthCheck(opts *bind.FilterOpts) (*ConvexBaseStrategySetHealthCheckIterator, error) {

	logs, sub, err := _ConvexBaseStrategy.contract.FilterLogs(opts, "SetHealthCheck")
	if err != nil {
		return nil, err
	}
	return &ConvexBaseStrategySetHealthCheckIterator{contract: _ConvexBaseStrategy.contract, event: "SetHealthCheck", logs: logs, sub: sub}, nil
}

// WatchSetHealthCheck is a free log subscription operation binding the contract event 0xc8db9c35f716b87af1fbb83f03c78646061931269301fd7ba6dcf189b4cdc2fc.
//
// Solidity: event SetHealthCheck(address arg0)
func (_ConvexBaseStrategy *ConvexBaseStrategyFilterer) WatchSetHealthCheck(opts *bind.WatchOpts, sink chan<- *ConvexBaseStrategySetHealthCheck) (event.Subscription, error) {

	logs, sub, err := _ConvexBaseStrategy.contract.WatchLogs(opts, "SetHealthCheck")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConvexBaseStrategySetHealthCheck)
				if err := _ConvexBaseStrategy.contract.UnpackLog(event, "SetHealthCheck", log); err != nil {
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
func (_ConvexBaseStrategy *ConvexBaseStrategyFilterer) ParseSetHealthCheck(log types.Log) (*ConvexBaseStrategySetHealthCheck, error) {
	event := new(ConvexBaseStrategySetHealthCheck)
	if err := _ConvexBaseStrategy.contract.UnpackLog(event, "SetHealthCheck", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConvexBaseStrategyUpdatedBaseFeeOracleIterator is returned from FilterUpdatedBaseFeeOracle and is used to iterate over the raw logs and unpacked data for UpdatedBaseFeeOracle events raised by the ConvexBaseStrategy contract.
type ConvexBaseStrategyUpdatedBaseFeeOracleIterator struct {
	Event *ConvexBaseStrategyUpdatedBaseFeeOracle // Event containing the contract specifics and raw log

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
func (it *ConvexBaseStrategyUpdatedBaseFeeOracleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConvexBaseStrategyUpdatedBaseFeeOracle)
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
		it.Event = new(ConvexBaseStrategyUpdatedBaseFeeOracle)
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
func (it *ConvexBaseStrategyUpdatedBaseFeeOracleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConvexBaseStrategyUpdatedBaseFeeOracleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConvexBaseStrategyUpdatedBaseFeeOracle represents a UpdatedBaseFeeOracle event raised by the ConvexBaseStrategy contract.
type ConvexBaseStrategyUpdatedBaseFeeOracle struct {
	BaseFeeOracle common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedBaseFeeOracle is a free log retrieval operation binding the contract event 0x711be97287cb9ec921887b9be36e148e1a27c6b158547b22b9704ffc54447a0f.
//
// Solidity: event UpdatedBaseFeeOracle(address baseFeeOracle)
func (_ConvexBaseStrategy *ConvexBaseStrategyFilterer) FilterUpdatedBaseFeeOracle(opts *bind.FilterOpts) (*ConvexBaseStrategyUpdatedBaseFeeOracleIterator, error) {

	logs, sub, err := _ConvexBaseStrategy.contract.FilterLogs(opts, "UpdatedBaseFeeOracle")
	if err != nil {
		return nil, err
	}
	return &ConvexBaseStrategyUpdatedBaseFeeOracleIterator{contract: _ConvexBaseStrategy.contract, event: "UpdatedBaseFeeOracle", logs: logs, sub: sub}, nil
}

// WatchUpdatedBaseFeeOracle is a free log subscription operation binding the contract event 0x711be97287cb9ec921887b9be36e148e1a27c6b158547b22b9704ffc54447a0f.
//
// Solidity: event UpdatedBaseFeeOracle(address baseFeeOracle)
func (_ConvexBaseStrategy *ConvexBaseStrategyFilterer) WatchUpdatedBaseFeeOracle(opts *bind.WatchOpts, sink chan<- *ConvexBaseStrategyUpdatedBaseFeeOracle) (event.Subscription, error) {

	logs, sub, err := _ConvexBaseStrategy.contract.WatchLogs(opts, "UpdatedBaseFeeOracle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConvexBaseStrategyUpdatedBaseFeeOracle)
				if err := _ConvexBaseStrategy.contract.UnpackLog(event, "UpdatedBaseFeeOracle", log); err != nil {
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
func (_ConvexBaseStrategy *ConvexBaseStrategyFilterer) ParseUpdatedBaseFeeOracle(log types.Log) (*ConvexBaseStrategyUpdatedBaseFeeOracle, error) {
	event := new(ConvexBaseStrategyUpdatedBaseFeeOracle)
	if err := _ConvexBaseStrategy.contract.UnpackLog(event, "UpdatedBaseFeeOracle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConvexBaseStrategyUpdatedCreditThresholdIterator is returned from FilterUpdatedCreditThreshold and is used to iterate over the raw logs and unpacked data for UpdatedCreditThreshold events raised by the ConvexBaseStrategy contract.
type ConvexBaseStrategyUpdatedCreditThresholdIterator struct {
	Event *ConvexBaseStrategyUpdatedCreditThreshold // Event containing the contract specifics and raw log

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
func (it *ConvexBaseStrategyUpdatedCreditThresholdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConvexBaseStrategyUpdatedCreditThreshold)
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
		it.Event = new(ConvexBaseStrategyUpdatedCreditThreshold)
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
func (it *ConvexBaseStrategyUpdatedCreditThresholdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConvexBaseStrategyUpdatedCreditThresholdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConvexBaseStrategyUpdatedCreditThreshold represents a UpdatedCreditThreshold event raised by the ConvexBaseStrategy contract.
type ConvexBaseStrategyUpdatedCreditThreshold struct {
	CreditThreshold *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdatedCreditThreshold is a free log retrieval operation binding the contract event 0xe5ef7832c564a10cbe7b4f1e01ac33a406cb63fcf430a97a9af8616d150af5f3.
//
// Solidity: event UpdatedCreditThreshold(uint256 creditThreshold)
func (_ConvexBaseStrategy *ConvexBaseStrategyFilterer) FilterUpdatedCreditThreshold(opts *bind.FilterOpts) (*ConvexBaseStrategyUpdatedCreditThresholdIterator, error) {

	logs, sub, err := _ConvexBaseStrategy.contract.FilterLogs(opts, "UpdatedCreditThreshold")
	if err != nil {
		return nil, err
	}
	return &ConvexBaseStrategyUpdatedCreditThresholdIterator{contract: _ConvexBaseStrategy.contract, event: "UpdatedCreditThreshold", logs: logs, sub: sub}, nil
}

// WatchUpdatedCreditThreshold is a free log subscription operation binding the contract event 0xe5ef7832c564a10cbe7b4f1e01ac33a406cb63fcf430a97a9af8616d150af5f3.
//
// Solidity: event UpdatedCreditThreshold(uint256 creditThreshold)
func (_ConvexBaseStrategy *ConvexBaseStrategyFilterer) WatchUpdatedCreditThreshold(opts *bind.WatchOpts, sink chan<- *ConvexBaseStrategyUpdatedCreditThreshold) (event.Subscription, error) {

	logs, sub, err := _ConvexBaseStrategy.contract.WatchLogs(opts, "UpdatedCreditThreshold")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConvexBaseStrategyUpdatedCreditThreshold)
				if err := _ConvexBaseStrategy.contract.UnpackLog(event, "UpdatedCreditThreshold", log); err != nil {
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
func (_ConvexBaseStrategy *ConvexBaseStrategyFilterer) ParseUpdatedCreditThreshold(log types.Log) (*ConvexBaseStrategyUpdatedCreditThreshold, error) {
	event := new(ConvexBaseStrategyUpdatedCreditThreshold)
	if err := _ConvexBaseStrategy.contract.UnpackLog(event, "UpdatedCreditThreshold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConvexBaseStrategyUpdatedKeeperIterator is returned from FilterUpdatedKeeper and is used to iterate over the raw logs and unpacked data for UpdatedKeeper events raised by the ConvexBaseStrategy contract.
type ConvexBaseStrategyUpdatedKeeperIterator struct {
	Event *ConvexBaseStrategyUpdatedKeeper // Event containing the contract specifics and raw log

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
func (it *ConvexBaseStrategyUpdatedKeeperIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConvexBaseStrategyUpdatedKeeper)
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
		it.Event = new(ConvexBaseStrategyUpdatedKeeper)
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
func (it *ConvexBaseStrategyUpdatedKeeperIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConvexBaseStrategyUpdatedKeeperIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConvexBaseStrategyUpdatedKeeper represents a UpdatedKeeper event raised by the ConvexBaseStrategy contract.
type ConvexBaseStrategyUpdatedKeeper struct {
	NewKeeper common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdatedKeeper is a free log retrieval operation binding the contract event 0x2f202ddb4a2e345f6323ed90f8fc8559d770a7abbbeee84dde8aca3351fe7154.
//
// Solidity: event UpdatedKeeper(address newKeeper)
func (_ConvexBaseStrategy *ConvexBaseStrategyFilterer) FilterUpdatedKeeper(opts *bind.FilterOpts) (*ConvexBaseStrategyUpdatedKeeperIterator, error) {

	logs, sub, err := _ConvexBaseStrategy.contract.FilterLogs(opts, "UpdatedKeeper")
	if err != nil {
		return nil, err
	}
	return &ConvexBaseStrategyUpdatedKeeperIterator{contract: _ConvexBaseStrategy.contract, event: "UpdatedKeeper", logs: logs, sub: sub}, nil
}

// WatchUpdatedKeeper is a free log subscription operation binding the contract event 0x2f202ddb4a2e345f6323ed90f8fc8559d770a7abbbeee84dde8aca3351fe7154.
//
// Solidity: event UpdatedKeeper(address newKeeper)
func (_ConvexBaseStrategy *ConvexBaseStrategyFilterer) WatchUpdatedKeeper(opts *bind.WatchOpts, sink chan<- *ConvexBaseStrategyUpdatedKeeper) (event.Subscription, error) {

	logs, sub, err := _ConvexBaseStrategy.contract.WatchLogs(opts, "UpdatedKeeper")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConvexBaseStrategyUpdatedKeeper)
				if err := _ConvexBaseStrategy.contract.UnpackLog(event, "UpdatedKeeper", log); err != nil {
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
func (_ConvexBaseStrategy *ConvexBaseStrategyFilterer) ParseUpdatedKeeper(log types.Log) (*ConvexBaseStrategyUpdatedKeeper, error) {
	event := new(ConvexBaseStrategyUpdatedKeeper)
	if err := _ConvexBaseStrategy.contract.UnpackLog(event, "UpdatedKeeper", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConvexBaseStrategyUpdatedMaxReportDelayIterator is returned from FilterUpdatedMaxReportDelay and is used to iterate over the raw logs and unpacked data for UpdatedMaxReportDelay events raised by the ConvexBaseStrategy contract.
type ConvexBaseStrategyUpdatedMaxReportDelayIterator struct {
	Event *ConvexBaseStrategyUpdatedMaxReportDelay // Event containing the contract specifics and raw log

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
func (it *ConvexBaseStrategyUpdatedMaxReportDelayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConvexBaseStrategyUpdatedMaxReportDelay)
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
		it.Event = new(ConvexBaseStrategyUpdatedMaxReportDelay)
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
func (it *ConvexBaseStrategyUpdatedMaxReportDelayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConvexBaseStrategyUpdatedMaxReportDelayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConvexBaseStrategyUpdatedMaxReportDelay represents a UpdatedMaxReportDelay event raised by the ConvexBaseStrategy contract.
type ConvexBaseStrategyUpdatedMaxReportDelay struct {
	Delay *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdatedMaxReportDelay is a free log retrieval operation binding the contract event 0x5430e11864ad7aa9775b07d12657fe52df9aa2ba734355bd8ef8747be2c800c5.
//
// Solidity: event UpdatedMaxReportDelay(uint256 delay)
func (_ConvexBaseStrategy *ConvexBaseStrategyFilterer) FilterUpdatedMaxReportDelay(opts *bind.FilterOpts) (*ConvexBaseStrategyUpdatedMaxReportDelayIterator, error) {

	logs, sub, err := _ConvexBaseStrategy.contract.FilterLogs(opts, "UpdatedMaxReportDelay")
	if err != nil {
		return nil, err
	}
	return &ConvexBaseStrategyUpdatedMaxReportDelayIterator{contract: _ConvexBaseStrategy.contract, event: "UpdatedMaxReportDelay", logs: logs, sub: sub}, nil
}

// WatchUpdatedMaxReportDelay is a free log subscription operation binding the contract event 0x5430e11864ad7aa9775b07d12657fe52df9aa2ba734355bd8ef8747be2c800c5.
//
// Solidity: event UpdatedMaxReportDelay(uint256 delay)
func (_ConvexBaseStrategy *ConvexBaseStrategyFilterer) WatchUpdatedMaxReportDelay(opts *bind.WatchOpts, sink chan<- *ConvexBaseStrategyUpdatedMaxReportDelay) (event.Subscription, error) {

	logs, sub, err := _ConvexBaseStrategy.contract.WatchLogs(opts, "UpdatedMaxReportDelay")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConvexBaseStrategyUpdatedMaxReportDelay)
				if err := _ConvexBaseStrategy.contract.UnpackLog(event, "UpdatedMaxReportDelay", log); err != nil {
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
func (_ConvexBaseStrategy *ConvexBaseStrategyFilterer) ParseUpdatedMaxReportDelay(log types.Log) (*ConvexBaseStrategyUpdatedMaxReportDelay, error) {
	event := new(ConvexBaseStrategyUpdatedMaxReportDelay)
	if err := _ConvexBaseStrategy.contract.UnpackLog(event, "UpdatedMaxReportDelay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConvexBaseStrategyUpdatedMetadataURIIterator is returned from FilterUpdatedMetadataURI and is used to iterate over the raw logs and unpacked data for UpdatedMetadataURI events raised by the ConvexBaseStrategy contract.
type ConvexBaseStrategyUpdatedMetadataURIIterator struct {
	Event *ConvexBaseStrategyUpdatedMetadataURI // Event containing the contract specifics and raw log

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
func (it *ConvexBaseStrategyUpdatedMetadataURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConvexBaseStrategyUpdatedMetadataURI)
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
		it.Event = new(ConvexBaseStrategyUpdatedMetadataURI)
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
func (it *ConvexBaseStrategyUpdatedMetadataURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConvexBaseStrategyUpdatedMetadataURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConvexBaseStrategyUpdatedMetadataURI represents a UpdatedMetadataURI event raised by the ConvexBaseStrategy contract.
type ConvexBaseStrategyUpdatedMetadataURI struct {
	MetadataURI string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatedMetadataURI is a free log retrieval operation binding the contract event 0x300e67d5a415b6d015a471d9c7b95dd58f3e8290af965e84e0f845de2996dda6.
//
// Solidity: event UpdatedMetadataURI(string metadataURI)
func (_ConvexBaseStrategy *ConvexBaseStrategyFilterer) FilterUpdatedMetadataURI(opts *bind.FilterOpts) (*ConvexBaseStrategyUpdatedMetadataURIIterator, error) {

	logs, sub, err := _ConvexBaseStrategy.contract.FilterLogs(opts, "UpdatedMetadataURI")
	if err != nil {
		return nil, err
	}
	return &ConvexBaseStrategyUpdatedMetadataURIIterator{contract: _ConvexBaseStrategy.contract, event: "UpdatedMetadataURI", logs: logs, sub: sub}, nil
}

// WatchUpdatedMetadataURI is a free log subscription operation binding the contract event 0x300e67d5a415b6d015a471d9c7b95dd58f3e8290af965e84e0f845de2996dda6.
//
// Solidity: event UpdatedMetadataURI(string metadataURI)
func (_ConvexBaseStrategy *ConvexBaseStrategyFilterer) WatchUpdatedMetadataURI(opts *bind.WatchOpts, sink chan<- *ConvexBaseStrategyUpdatedMetadataURI) (event.Subscription, error) {

	logs, sub, err := _ConvexBaseStrategy.contract.WatchLogs(opts, "UpdatedMetadataURI")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConvexBaseStrategyUpdatedMetadataURI)
				if err := _ConvexBaseStrategy.contract.UnpackLog(event, "UpdatedMetadataURI", log); err != nil {
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
func (_ConvexBaseStrategy *ConvexBaseStrategyFilterer) ParseUpdatedMetadataURI(log types.Log) (*ConvexBaseStrategyUpdatedMetadataURI, error) {
	event := new(ConvexBaseStrategyUpdatedMetadataURI)
	if err := _ConvexBaseStrategy.contract.UnpackLog(event, "UpdatedMetadataURI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConvexBaseStrategyUpdatedMinReportDelayIterator is returned from FilterUpdatedMinReportDelay and is used to iterate over the raw logs and unpacked data for UpdatedMinReportDelay events raised by the ConvexBaseStrategy contract.
type ConvexBaseStrategyUpdatedMinReportDelayIterator struct {
	Event *ConvexBaseStrategyUpdatedMinReportDelay // Event containing the contract specifics and raw log

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
func (it *ConvexBaseStrategyUpdatedMinReportDelayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConvexBaseStrategyUpdatedMinReportDelay)
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
		it.Event = new(ConvexBaseStrategyUpdatedMinReportDelay)
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
func (it *ConvexBaseStrategyUpdatedMinReportDelayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConvexBaseStrategyUpdatedMinReportDelayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConvexBaseStrategyUpdatedMinReportDelay represents a UpdatedMinReportDelay event raised by the ConvexBaseStrategy contract.
type ConvexBaseStrategyUpdatedMinReportDelay struct {
	Delay *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdatedMinReportDelay is a free log retrieval operation binding the contract event 0xbb2c369a0355a34b02ab5fce0643150c87e1c8dfe7c918d465591879f57948b1.
//
// Solidity: event UpdatedMinReportDelay(uint256 delay)
func (_ConvexBaseStrategy *ConvexBaseStrategyFilterer) FilterUpdatedMinReportDelay(opts *bind.FilterOpts) (*ConvexBaseStrategyUpdatedMinReportDelayIterator, error) {

	logs, sub, err := _ConvexBaseStrategy.contract.FilterLogs(opts, "UpdatedMinReportDelay")
	if err != nil {
		return nil, err
	}
	return &ConvexBaseStrategyUpdatedMinReportDelayIterator{contract: _ConvexBaseStrategy.contract, event: "UpdatedMinReportDelay", logs: logs, sub: sub}, nil
}

// WatchUpdatedMinReportDelay is a free log subscription operation binding the contract event 0xbb2c369a0355a34b02ab5fce0643150c87e1c8dfe7c918d465591879f57948b1.
//
// Solidity: event UpdatedMinReportDelay(uint256 delay)
func (_ConvexBaseStrategy *ConvexBaseStrategyFilterer) WatchUpdatedMinReportDelay(opts *bind.WatchOpts, sink chan<- *ConvexBaseStrategyUpdatedMinReportDelay) (event.Subscription, error) {

	logs, sub, err := _ConvexBaseStrategy.contract.WatchLogs(opts, "UpdatedMinReportDelay")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConvexBaseStrategyUpdatedMinReportDelay)
				if err := _ConvexBaseStrategy.contract.UnpackLog(event, "UpdatedMinReportDelay", log); err != nil {
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
func (_ConvexBaseStrategy *ConvexBaseStrategyFilterer) ParseUpdatedMinReportDelay(log types.Log) (*ConvexBaseStrategyUpdatedMinReportDelay, error) {
	event := new(ConvexBaseStrategyUpdatedMinReportDelay)
	if err := _ConvexBaseStrategy.contract.UnpackLog(event, "UpdatedMinReportDelay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConvexBaseStrategyUpdatedRewardsIterator is returned from FilterUpdatedRewards and is used to iterate over the raw logs and unpacked data for UpdatedRewards events raised by the ConvexBaseStrategy contract.
type ConvexBaseStrategyUpdatedRewardsIterator struct {
	Event *ConvexBaseStrategyUpdatedRewards // Event containing the contract specifics and raw log

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
func (it *ConvexBaseStrategyUpdatedRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConvexBaseStrategyUpdatedRewards)
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
		it.Event = new(ConvexBaseStrategyUpdatedRewards)
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
func (it *ConvexBaseStrategyUpdatedRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConvexBaseStrategyUpdatedRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConvexBaseStrategyUpdatedRewards represents a UpdatedRewards event raised by the ConvexBaseStrategy contract.
type ConvexBaseStrategyUpdatedRewards struct {
	Rewards common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdatedRewards is a free log retrieval operation binding the contract event 0xafbb66abf8f3b719799940473a4052a3717cdd8e40fb6c8a3faadab316b1a069.
//
// Solidity: event UpdatedRewards(address rewards)
func (_ConvexBaseStrategy *ConvexBaseStrategyFilterer) FilterUpdatedRewards(opts *bind.FilterOpts) (*ConvexBaseStrategyUpdatedRewardsIterator, error) {

	logs, sub, err := _ConvexBaseStrategy.contract.FilterLogs(opts, "UpdatedRewards")
	if err != nil {
		return nil, err
	}
	return &ConvexBaseStrategyUpdatedRewardsIterator{contract: _ConvexBaseStrategy.contract, event: "UpdatedRewards", logs: logs, sub: sub}, nil
}

// WatchUpdatedRewards is a free log subscription operation binding the contract event 0xafbb66abf8f3b719799940473a4052a3717cdd8e40fb6c8a3faadab316b1a069.
//
// Solidity: event UpdatedRewards(address rewards)
func (_ConvexBaseStrategy *ConvexBaseStrategyFilterer) WatchUpdatedRewards(opts *bind.WatchOpts, sink chan<- *ConvexBaseStrategyUpdatedRewards) (event.Subscription, error) {

	logs, sub, err := _ConvexBaseStrategy.contract.WatchLogs(opts, "UpdatedRewards")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConvexBaseStrategyUpdatedRewards)
				if err := _ConvexBaseStrategy.contract.UnpackLog(event, "UpdatedRewards", log); err != nil {
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
func (_ConvexBaseStrategy *ConvexBaseStrategyFilterer) ParseUpdatedRewards(log types.Log) (*ConvexBaseStrategyUpdatedRewards, error) {
	event := new(ConvexBaseStrategyUpdatedRewards)
	if err := _ConvexBaseStrategy.contract.UnpackLog(event, "UpdatedRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConvexBaseStrategyUpdatedStrategistIterator is returned from FilterUpdatedStrategist and is used to iterate over the raw logs and unpacked data for UpdatedStrategist events raised by the ConvexBaseStrategy contract.
type ConvexBaseStrategyUpdatedStrategistIterator struct {
	Event *ConvexBaseStrategyUpdatedStrategist // Event containing the contract specifics and raw log

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
func (it *ConvexBaseStrategyUpdatedStrategistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConvexBaseStrategyUpdatedStrategist)
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
		it.Event = new(ConvexBaseStrategyUpdatedStrategist)
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
func (it *ConvexBaseStrategyUpdatedStrategistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConvexBaseStrategyUpdatedStrategistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConvexBaseStrategyUpdatedStrategist represents a UpdatedStrategist event raised by the ConvexBaseStrategy contract.
type ConvexBaseStrategyUpdatedStrategist struct {
	NewStrategist common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedStrategist is a free log retrieval operation binding the contract event 0x352ececae6d7d1e6d26bcf2c549dfd55be1637e9b22dc0cf3b71ddb36097a6b4.
//
// Solidity: event UpdatedStrategist(address newStrategist)
func (_ConvexBaseStrategy *ConvexBaseStrategyFilterer) FilterUpdatedStrategist(opts *bind.FilterOpts) (*ConvexBaseStrategyUpdatedStrategistIterator, error) {

	logs, sub, err := _ConvexBaseStrategy.contract.FilterLogs(opts, "UpdatedStrategist")
	if err != nil {
		return nil, err
	}
	return &ConvexBaseStrategyUpdatedStrategistIterator{contract: _ConvexBaseStrategy.contract, event: "UpdatedStrategist", logs: logs, sub: sub}, nil
}

// WatchUpdatedStrategist is a free log subscription operation binding the contract event 0x352ececae6d7d1e6d26bcf2c549dfd55be1637e9b22dc0cf3b71ddb36097a6b4.
//
// Solidity: event UpdatedStrategist(address newStrategist)
func (_ConvexBaseStrategy *ConvexBaseStrategyFilterer) WatchUpdatedStrategist(opts *bind.WatchOpts, sink chan<- *ConvexBaseStrategyUpdatedStrategist) (event.Subscription, error) {

	logs, sub, err := _ConvexBaseStrategy.contract.WatchLogs(opts, "UpdatedStrategist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConvexBaseStrategyUpdatedStrategist)
				if err := _ConvexBaseStrategy.contract.UnpackLog(event, "UpdatedStrategist", log); err != nil {
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
func (_ConvexBaseStrategy *ConvexBaseStrategyFilterer) ParseUpdatedStrategist(log types.Log) (*ConvexBaseStrategyUpdatedStrategist, error) {
	event := new(ConvexBaseStrategyUpdatedStrategist)
	if err := _ConvexBaseStrategy.contract.UnpackLog(event, "UpdatedStrategist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
