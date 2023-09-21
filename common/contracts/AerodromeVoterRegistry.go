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

// AerodromeVoterRegistryMetaData contains all meta data concerning the AerodromeVoterRegistry contract.
var AerodromeVoterRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_forwarder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_factoryRegistry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyVotedOrDeposited\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DistributeWindow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FactoryPathNotApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GaugeAlreadyKilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GaugeAlreadyRevived\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"GaugeDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GaugeExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gauge\",\"type\":\"address\"}],\"name\":\"GaugeNotAlive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InactiveManagedNFT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaximumVotingNumberTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonZeroVotes\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAPool\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotApprovedOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEmergencyCouncil\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotGovernor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotMinter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotWhitelistedNFT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotWhitelistedToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SameValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SpecialVotingWindow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnequalLengths\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroBalance\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalWeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Abstained\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DistributeReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"poolFactory\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"votingRewardsFactory\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gaugeFactory\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bribeVotingReward\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeVotingReward\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"GaugeCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"}],\"name\":\"GaugeKilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"}],\"name\":\"GaugeRevived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NotifyReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalWeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Voted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"whitelister\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"_bool\",\"type\":\"bool\"}],\"name\":\"WhitelistNFT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"whitelister\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"_bool\",\"type\":\"bool\"}],\"name\":\"WhitelistToken\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_bribes\",\"type\":\"address[]\"},{\"internalType\":\"address[][]\",\"name\":\"_tokens\",\"type\":\"address[][]\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"claimBribes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_fees\",\"type\":\"address[]\"},{\"internalType\":\"address[][]\",\"name\":\"_tokens\",\"type\":\"address[][]\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"claimFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_gauges\",\"type\":\"address[]\"}],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"createGauge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_mTokenId\",\"type\":\"uint256\"}],\"name\":\"depositManaged\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_gauges\",\"type\":\"address[]\"}],\"name\":\"distribute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_finish\",\"type\":\"uint256\"}],\"name\":\"distribute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyCouncil\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochGovernor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"epochNext\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"epochStart\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"epochVoteEnd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"epochVoteStart\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factoryRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forwarder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"gaugeToBribe\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"gaugeToFees\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"gauges\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isAlive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isGauge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"isWhitelistedNFT\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isWhitelistedToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gauge\",\"type\":\"address\"}],\"name\":\"killGauge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lastVoted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"length\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxVotingNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"notifyRewardAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"poke\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"poolForGauge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolVote\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pools\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"reset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gauge\",\"type\":\"address\"}],\"name\":\"reviveGauge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_council\",\"type\":\"address\"}],\"name\":\"setEmergencyCouncil\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_epochGovernor\",\"type\":\"address\"}],\"name\":\"setEpochGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governor\",\"type\":\"address\"}],\"name\":\"setGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxVotingNum\",\"type\":\"uint256\"}],\"name\":\"setMaxVotingNum\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gauge\",\"type\":\"address\"}],\"name\":\"updateFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"updateFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_gauges\",\"type\":\"address[]\"}],\"name\":\"updateFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"usedWeights\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ve\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_poolVote\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_weights\",\"type\":\"uint256[]\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"votes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"weights\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_bool\",\"type\":\"bool\"}],\"name\":\"whitelistNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_bool\",\"type\":\"bool\"}],\"name\":\"whitelistToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"withdrawManaged\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AerodromeVoterRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use AerodromeVoterRegistryMetaData.ABI instead.
var AerodromeVoterRegistryABI = AerodromeVoterRegistryMetaData.ABI

// AerodromeVoterRegistry is an auto generated Go binding around an Ethereum contract.
type AerodromeVoterRegistry struct {
	AerodromeVoterRegistryCaller     // Read-only binding to the contract
	AerodromeVoterRegistryTransactor // Write-only binding to the contract
	AerodromeVoterRegistryFilterer   // Log filterer for contract events
}

// AerodromeVoterRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AerodromeVoterRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AerodromeVoterRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AerodromeVoterRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AerodromeVoterRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AerodromeVoterRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AerodromeVoterRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AerodromeVoterRegistrySession struct {
	Contract     *AerodromeVoterRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AerodromeVoterRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AerodromeVoterRegistryCallerSession struct {
	Contract *AerodromeVoterRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// AerodromeVoterRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AerodromeVoterRegistryTransactorSession struct {
	Contract     *AerodromeVoterRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// AerodromeVoterRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type AerodromeVoterRegistryRaw struct {
	Contract *AerodromeVoterRegistry // Generic contract binding to access the raw methods on
}

// AerodromeVoterRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AerodromeVoterRegistryCallerRaw struct {
	Contract *AerodromeVoterRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// AerodromeVoterRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AerodromeVoterRegistryTransactorRaw struct {
	Contract *AerodromeVoterRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAerodromeVoterRegistry creates a new instance of AerodromeVoterRegistry, bound to a specific deployed contract.
func NewAerodromeVoterRegistry(address common.Address, backend bind.ContractBackend) (*AerodromeVoterRegistry, error) {
	contract, err := bindAerodromeVoterRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AerodromeVoterRegistry{AerodromeVoterRegistryCaller: AerodromeVoterRegistryCaller{contract: contract}, AerodromeVoterRegistryTransactor: AerodromeVoterRegistryTransactor{contract: contract}, AerodromeVoterRegistryFilterer: AerodromeVoterRegistryFilterer{contract: contract}}, nil
}

// NewAerodromeVoterRegistryCaller creates a new read-only instance of AerodromeVoterRegistry, bound to a specific deployed contract.
func NewAerodromeVoterRegistryCaller(address common.Address, caller bind.ContractCaller) (*AerodromeVoterRegistryCaller, error) {
	contract, err := bindAerodromeVoterRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AerodromeVoterRegistryCaller{contract: contract}, nil
}

// NewAerodromeVoterRegistryTransactor creates a new write-only instance of AerodromeVoterRegistry, bound to a specific deployed contract.
func NewAerodromeVoterRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*AerodromeVoterRegistryTransactor, error) {
	contract, err := bindAerodromeVoterRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AerodromeVoterRegistryTransactor{contract: contract}, nil
}

// NewAerodromeVoterRegistryFilterer creates a new log filterer instance of AerodromeVoterRegistry, bound to a specific deployed contract.
func NewAerodromeVoterRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*AerodromeVoterRegistryFilterer, error) {
	contract, err := bindAerodromeVoterRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AerodromeVoterRegistryFilterer{contract: contract}, nil
}

// bindAerodromeVoterRegistry binds a generic wrapper to an already deployed contract.
func bindAerodromeVoterRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AerodromeVoterRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AerodromeVoterRegistry *AerodromeVoterRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AerodromeVoterRegistry.Contract.AerodromeVoterRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AerodromeVoterRegistry *AerodromeVoterRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.AerodromeVoterRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AerodromeVoterRegistry *AerodromeVoterRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.AerodromeVoterRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AerodromeVoterRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.contract.Transact(opts, method, params...)
}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address ) view returns(uint256)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCaller) Claimable(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AerodromeVoterRegistry.contract.Call(opts, &out, "claimable", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address ) view returns(uint256)
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) Claimable(arg0 common.Address) (*big.Int, error) {
	return _AerodromeVoterRegistry.Contract.Claimable(&_AerodromeVoterRegistry.CallOpts, arg0)
}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address ) view returns(uint256)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCallerSession) Claimable(arg0 common.Address) (*big.Int, error) {
	return _AerodromeVoterRegistry.Contract.Claimable(&_AerodromeVoterRegistry.CallOpts, arg0)
}

// EmergencyCouncil is a free data retrieval call binding the contract method 0x7778960e.
//
// Solidity: function emergencyCouncil() view returns(address)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCaller) EmergencyCouncil(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AerodromeVoterRegistry.contract.Call(opts, &out, "emergencyCouncil")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EmergencyCouncil is a free data retrieval call binding the contract method 0x7778960e.
//
// Solidity: function emergencyCouncil() view returns(address)
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) EmergencyCouncil() (common.Address, error) {
	return _AerodromeVoterRegistry.Contract.EmergencyCouncil(&_AerodromeVoterRegistry.CallOpts)
}

// EmergencyCouncil is a free data retrieval call binding the contract method 0x7778960e.
//
// Solidity: function emergencyCouncil() view returns(address)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCallerSession) EmergencyCouncil() (common.Address, error) {
	return _AerodromeVoterRegistry.Contract.EmergencyCouncil(&_AerodromeVoterRegistry.CallOpts)
}

// EpochGovernor is a free data retrieval call binding the contract method 0x3aae971f.
//
// Solidity: function epochGovernor() view returns(address)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCaller) EpochGovernor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AerodromeVoterRegistry.contract.Call(opts, &out, "epochGovernor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EpochGovernor is a free data retrieval call binding the contract method 0x3aae971f.
//
// Solidity: function epochGovernor() view returns(address)
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) EpochGovernor() (common.Address, error) {
	return _AerodromeVoterRegistry.Contract.EpochGovernor(&_AerodromeVoterRegistry.CallOpts)
}

// EpochGovernor is a free data retrieval call binding the contract method 0x3aae971f.
//
// Solidity: function epochGovernor() view returns(address)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCallerSession) EpochGovernor() (common.Address, error) {
	return _AerodromeVoterRegistry.Contract.EpochGovernor(&_AerodromeVoterRegistry.CallOpts)
}

// EpochNext is a free data retrieval call binding the contract method 0x880e36fc.
//
// Solidity: function epochNext(uint256 _timestamp) pure returns(uint256)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCaller) EpochNext(opts *bind.CallOpts, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AerodromeVoterRegistry.contract.Call(opts, &out, "epochNext", _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochNext is a free data retrieval call binding the contract method 0x880e36fc.
//
// Solidity: function epochNext(uint256 _timestamp) pure returns(uint256)
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) EpochNext(_timestamp *big.Int) (*big.Int, error) {
	return _AerodromeVoterRegistry.Contract.EpochNext(&_AerodromeVoterRegistry.CallOpts, _timestamp)
}

// EpochNext is a free data retrieval call binding the contract method 0x880e36fc.
//
// Solidity: function epochNext(uint256 _timestamp) pure returns(uint256)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCallerSession) EpochNext(_timestamp *big.Int) (*big.Int, error) {
	return _AerodromeVoterRegistry.Contract.EpochNext(&_AerodromeVoterRegistry.CallOpts, _timestamp)
}

// EpochStart is a free data retrieval call binding the contract method 0xaa9354a3.
//
// Solidity: function epochStart(uint256 _timestamp) pure returns(uint256)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCaller) EpochStart(opts *bind.CallOpts, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AerodromeVoterRegistry.contract.Call(opts, &out, "epochStart", _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochStart is a free data retrieval call binding the contract method 0xaa9354a3.
//
// Solidity: function epochStart(uint256 _timestamp) pure returns(uint256)
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) EpochStart(_timestamp *big.Int) (*big.Int, error) {
	return _AerodromeVoterRegistry.Contract.EpochStart(&_AerodromeVoterRegistry.CallOpts, _timestamp)
}

// EpochStart is a free data retrieval call binding the contract method 0xaa9354a3.
//
// Solidity: function epochStart(uint256 _timestamp) pure returns(uint256)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCallerSession) EpochStart(_timestamp *big.Int) (*big.Int, error) {
	return _AerodromeVoterRegistry.Contract.EpochStart(&_AerodromeVoterRegistry.CallOpts, _timestamp)
}

// EpochVoteEnd is a free data retrieval call binding the contract method 0xd58b15d4.
//
// Solidity: function epochVoteEnd(uint256 _timestamp) pure returns(uint256)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCaller) EpochVoteEnd(opts *bind.CallOpts, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AerodromeVoterRegistry.contract.Call(opts, &out, "epochVoteEnd", _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochVoteEnd is a free data retrieval call binding the contract method 0xd58b15d4.
//
// Solidity: function epochVoteEnd(uint256 _timestamp) pure returns(uint256)
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) EpochVoteEnd(_timestamp *big.Int) (*big.Int, error) {
	return _AerodromeVoterRegistry.Contract.EpochVoteEnd(&_AerodromeVoterRegistry.CallOpts, _timestamp)
}

// EpochVoteEnd is a free data retrieval call binding the contract method 0xd58b15d4.
//
// Solidity: function epochVoteEnd(uint256 _timestamp) pure returns(uint256)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCallerSession) EpochVoteEnd(_timestamp *big.Int) (*big.Int, error) {
	return _AerodromeVoterRegistry.Contract.EpochVoteEnd(&_AerodromeVoterRegistry.CallOpts, _timestamp)
}

// EpochVoteStart is a free data retrieval call binding the contract method 0x39e9f3b6.
//
// Solidity: function epochVoteStart(uint256 _timestamp) pure returns(uint256)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCaller) EpochVoteStart(opts *bind.CallOpts, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AerodromeVoterRegistry.contract.Call(opts, &out, "epochVoteStart", _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochVoteStart is a free data retrieval call binding the contract method 0x39e9f3b6.
//
// Solidity: function epochVoteStart(uint256 _timestamp) pure returns(uint256)
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) EpochVoteStart(_timestamp *big.Int) (*big.Int, error) {
	return _AerodromeVoterRegistry.Contract.EpochVoteStart(&_AerodromeVoterRegistry.CallOpts, _timestamp)
}

// EpochVoteStart is a free data retrieval call binding the contract method 0x39e9f3b6.
//
// Solidity: function epochVoteStart(uint256 _timestamp) pure returns(uint256)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCallerSession) EpochVoteStart(_timestamp *big.Int) (*big.Int, error) {
	return _AerodromeVoterRegistry.Contract.EpochVoteStart(&_AerodromeVoterRegistry.CallOpts, _timestamp)
}

// FactoryRegistry is a free data retrieval call binding the contract method 0x3bf0c9fb.
//
// Solidity: function factoryRegistry() view returns(address)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCaller) FactoryRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AerodromeVoterRegistry.contract.Call(opts, &out, "factoryRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FactoryRegistry is a free data retrieval call binding the contract method 0x3bf0c9fb.
//
// Solidity: function factoryRegistry() view returns(address)
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) FactoryRegistry() (common.Address, error) {
	return _AerodromeVoterRegistry.Contract.FactoryRegistry(&_AerodromeVoterRegistry.CallOpts)
}

// FactoryRegistry is a free data retrieval call binding the contract method 0x3bf0c9fb.
//
// Solidity: function factoryRegistry() view returns(address)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCallerSession) FactoryRegistry() (common.Address, error) {
	return _AerodromeVoterRegistry.Contract.FactoryRegistry(&_AerodromeVoterRegistry.CallOpts)
}

// Forwarder is a free data retrieval call binding the contract method 0xf645d4f9.
//
// Solidity: function forwarder() view returns(address)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCaller) Forwarder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AerodromeVoterRegistry.contract.Call(opts, &out, "forwarder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Forwarder is a free data retrieval call binding the contract method 0xf645d4f9.
//
// Solidity: function forwarder() view returns(address)
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) Forwarder() (common.Address, error) {
	return _AerodromeVoterRegistry.Contract.Forwarder(&_AerodromeVoterRegistry.CallOpts)
}

// Forwarder is a free data retrieval call binding the contract method 0xf645d4f9.
//
// Solidity: function forwarder() view returns(address)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCallerSession) Forwarder() (common.Address, error) {
	return _AerodromeVoterRegistry.Contract.Forwarder(&_AerodromeVoterRegistry.CallOpts)
}

// GaugeToBribe is a free data retrieval call binding the contract method 0x929c8dcd.
//
// Solidity: function gaugeToBribe(address ) view returns(address)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCaller) GaugeToBribe(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _AerodromeVoterRegistry.contract.Call(opts, &out, "gaugeToBribe", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GaugeToBribe is a free data retrieval call binding the contract method 0x929c8dcd.
//
// Solidity: function gaugeToBribe(address ) view returns(address)
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) GaugeToBribe(arg0 common.Address) (common.Address, error) {
	return _AerodromeVoterRegistry.Contract.GaugeToBribe(&_AerodromeVoterRegistry.CallOpts, arg0)
}

// GaugeToBribe is a free data retrieval call binding the contract method 0x929c8dcd.
//
// Solidity: function gaugeToBribe(address ) view returns(address)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCallerSession) GaugeToBribe(arg0 common.Address) (common.Address, error) {
	return _AerodromeVoterRegistry.Contract.GaugeToBribe(&_AerodromeVoterRegistry.CallOpts, arg0)
}

// GaugeToFees is a free data retrieval call binding the contract method 0xc4f08165.
//
// Solidity: function gaugeToFees(address ) view returns(address)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCaller) GaugeToFees(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _AerodromeVoterRegistry.contract.Call(opts, &out, "gaugeToFees", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GaugeToFees is a free data retrieval call binding the contract method 0xc4f08165.
//
// Solidity: function gaugeToFees(address ) view returns(address)
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) GaugeToFees(arg0 common.Address) (common.Address, error) {
	return _AerodromeVoterRegistry.Contract.GaugeToFees(&_AerodromeVoterRegistry.CallOpts, arg0)
}

// GaugeToFees is a free data retrieval call binding the contract method 0xc4f08165.
//
// Solidity: function gaugeToFees(address ) view returns(address)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCallerSession) GaugeToFees(arg0 common.Address) (common.Address, error) {
	return _AerodromeVoterRegistry.Contract.GaugeToFees(&_AerodromeVoterRegistry.CallOpts, arg0)
}

// Gauges is a free data retrieval call binding the contract method 0xb9a09fd5.
//
// Solidity: function gauges(address ) view returns(address)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCaller) Gauges(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _AerodromeVoterRegistry.contract.Call(opts, &out, "gauges", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gauges is a free data retrieval call binding the contract method 0xb9a09fd5.
//
// Solidity: function gauges(address ) view returns(address)
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) Gauges(arg0 common.Address) (common.Address, error) {
	return _AerodromeVoterRegistry.Contract.Gauges(&_AerodromeVoterRegistry.CallOpts, arg0)
}

// Gauges is a free data retrieval call binding the contract method 0xb9a09fd5.
//
// Solidity: function gauges(address ) view returns(address)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCallerSession) Gauges(arg0 common.Address) (common.Address, error) {
	return _AerodromeVoterRegistry.Contract.Gauges(&_AerodromeVoterRegistry.CallOpts, arg0)
}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(address)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCaller) Governor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AerodromeVoterRegistry.contract.Call(opts, &out, "governor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(address)
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) Governor() (common.Address, error) {
	return _AerodromeVoterRegistry.Contract.Governor(&_AerodromeVoterRegistry.CallOpts)
}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(address)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCallerSession) Governor() (common.Address, error) {
	return _AerodromeVoterRegistry.Contract.Governor(&_AerodromeVoterRegistry.CallOpts)
}

// IsAlive is a free data retrieval call binding the contract method 0x1703e5f9.
//
// Solidity: function isAlive(address ) view returns(bool)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCaller) IsAlive(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _AerodromeVoterRegistry.contract.Call(opts, &out, "isAlive", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAlive is a free data retrieval call binding the contract method 0x1703e5f9.
//
// Solidity: function isAlive(address ) view returns(bool)
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) IsAlive(arg0 common.Address) (bool, error) {
	return _AerodromeVoterRegistry.Contract.IsAlive(&_AerodromeVoterRegistry.CallOpts, arg0)
}

// IsAlive is a free data retrieval call binding the contract method 0x1703e5f9.
//
// Solidity: function isAlive(address ) view returns(bool)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCallerSession) IsAlive(arg0 common.Address) (bool, error) {
	return _AerodromeVoterRegistry.Contract.IsAlive(&_AerodromeVoterRegistry.CallOpts, arg0)
}

// IsGauge is a free data retrieval call binding the contract method 0xaa79979b.
//
// Solidity: function isGauge(address ) view returns(bool)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCaller) IsGauge(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _AerodromeVoterRegistry.contract.Call(opts, &out, "isGauge", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGauge is a free data retrieval call binding the contract method 0xaa79979b.
//
// Solidity: function isGauge(address ) view returns(bool)
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) IsGauge(arg0 common.Address) (bool, error) {
	return _AerodromeVoterRegistry.Contract.IsGauge(&_AerodromeVoterRegistry.CallOpts, arg0)
}

// IsGauge is a free data retrieval call binding the contract method 0xaa79979b.
//
// Solidity: function isGauge(address ) view returns(bool)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCallerSession) IsGauge(arg0 common.Address) (bool, error) {
	return _AerodromeVoterRegistry.Contract.IsGauge(&_AerodromeVoterRegistry.CallOpts, arg0)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _AerodromeVoterRegistry.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _AerodromeVoterRegistry.Contract.IsTrustedForwarder(&_AerodromeVoterRegistry.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _AerodromeVoterRegistry.Contract.IsTrustedForwarder(&_AerodromeVoterRegistry.CallOpts, forwarder)
}

// IsWhitelistedNFT is a free data retrieval call binding the contract method 0xd4e2616f.
//
// Solidity: function isWhitelistedNFT(uint256 ) view returns(bool)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCaller) IsWhitelistedNFT(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _AerodromeVoterRegistry.contract.Call(opts, &out, "isWhitelistedNFT", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelistedNFT is a free data retrieval call binding the contract method 0xd4e2616f.
//
// Solidity: function isWhitelistedNFT(uint256 ) view returns(bool)
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) IsWhitelistedNFT(arg0 *big.Int) (bool, error) {
	return _AerodromeVoterRegistry.Contract.IsWhitelistedNFT(&_AerodromeVoterRegistry.CallOpts, arg0)
}

// IsWhitelistedNFT is a free data retrieval call binding the contract method 0xd4e2616f.
//
// Solidity: function isWhitelistedNFT(uint256 ) view returns(bool)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCallerSession) IsWhitelistedNFT(arg0 *big.Int) (bool, error) {
	return _AerodromeVoterRegistry.Contract.IsWhitelistedNFT(&_AerodromeVoterRegistry.CallOpts, arg0)
}

// IsWhitelistedToken is a free data retrieval call binding the contract method 0xab37f486.
//
// Solidity: function isWhitelistedToken(address ) view returns(bool)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCaller) IsWhitelistedToken(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _AerodromeVoterRegistry.contract.Call(opts, &out, "isWhitelistedToken", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelistedToken is a free data retrieval call binding the contract method 0xab37f486.
//
// Solidity: function isWhitelistedToken(address ) view returns(bool)
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) IsWhitelistedToken(arg0 common.Address) (bool, error) {
	return _AerodromeVoterRegistry.Contract.IsWhitelistedToken(&_AerodromeVoterRegistry.CallOpts, arg0)
}

// IsWhitelistedToken is a free data retrieval call binding the contract method 0xab37f486.
//
// Solidity: function isWhitelistedToken(address ) view returns(bool)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCallerSession) IsWhitelistedToken(arg0 common.Address) (bool, error) {
	return _AerodromeVoterRegistry.Contract.IsWhitelistedToken(&_AerodromeVoterRegistry.CallOpts, arg0)
}

// LastVoted is a free data retrieval call binding the contract method 0xf3594be0.
//
// Solidity: function lastVoted(uint256 ) view returns(uint256)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCaller) LastVoted(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AerodromeVoterRegistry.contract.Call(opts, &out, "lastVoted", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastVoted is a free data retrieval call binding the contract method 0xf3594be0.
//
// Solidity: function lastVoted(uint256 ) view returns(uint256)
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) LastVoted(arg0 *big.Int) (*big.Int, error) {
	return _AerodromeVoterRegistry.Contract.LastVoted(&_AerodromeVoterRegistry.CallOpts, arg0)
}

// LastVoted is a free data retrieval call binding the contract method 0xf3594be0.
//
// Solidity: function lastVoted(uint256 ) view returns(uint256)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCallerSession) LastVoted(arg0 *big.Int) (*big.Int, error) {
	return _AerodromeVoterRegistry.Contract.LastVoted(&_AerodromeVoterRegistry.CallOpts, arg0)
}

// Length is a free data retrieval call binding the contract method 0x1f7b6d32.
//
// Solidity: function length() view returns(uint256)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCaller) Length(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AerodromeVoterRegistry.contract.Call(opts, &out, "length")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Length is a free data retrieval call binding the contract method 0x1f7b6d32.
//
// Solidity: function length() view returns(uint256)
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) Length() (*big.Int, error) {
	return _AerodromeVoterRegistry.Contract.Length(&_AerodromeVoterRegistry.CallOpts)
}

// Length is a free data retrieval call binding the contract method 0x1f7b6d32.
//
// Solidity: function length() view returns(uint256)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCallerSession) Length() (*big.Int, error) {
	return _AerodromeVoterRegistry.Contract.Length(&_AerodromeVoterRegistry.CallOpts)
}

// MaxVotingNum is a free data retrieval call binding the contract method 0xe8b3fd57.
//
// Solidity: function maxVotingNum() view returns(uint256)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCaller) MaxVotingNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AerodromeVoterRegistry.contract.Call(opts, &out, "maxVotingNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxVotingNum is a free data retrieval call binding the contract method 0xe8b3fd57.
//
// Solidity: function maxVotingNum() view returns(uint256)
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) MaxVotingNum() (*big.Int, error) {
	return _AerodromeVoterRegistry.Contract.MaxVotingNum(&_AerodromeVoterRegistry.CallOpts)
}

// MaxVotingNum is a free data retrieval call binding the contract method 0xe8b3fd57.
//
// Solidity: function maxVotingNum() view returns(uint256)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCallerSession) MaxVotingNum() (*big.Int, error) {
	return _AerodromeVoterRegistry.Contract.MaxVotingNum(&_AerodromeVoterRegistry.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AerodromeVoterRegistry.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) Minter() (common.Address, error) {
	return _AerodromeVoterRegistry.Contract.Minter(&_AerodromeVoterRegistry.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCallerSession) Minter() (common.Address, error) {
	return _AerodromeVoterRegistry.Contract.Minter(&_AerodromeVoterRegistry.CallOpts)
}

// PoolForGauge is a free data retrieval call binding the contract method 0x06d6a1b2.
//
// Solidity: function poolForGauge(address ) view returns(address)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCaller) PoolForGauge(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _AerodromeVoterRegistry.contract.Call(opts, &out, "poolForGauge", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolForGauge is a free data retrieval call binding the contract method 0x06d6a1b2.
//
// Solidity: function poolForGauge(address ) view returns(address)
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) PoolForGauge(arg0 common.Address) (common.Address, error) {
	return _AerodromeVoterRegistry.Contract.PoolForGauge(&_AerodromeVoterRegistry.CallOpts, arg0)
}

// PoolForGauge is a free data retrieval call binding the contract method 0x06d6a1b2.
//
// Solidity: function poolForGauge(address ) view returns(address)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCallerSession) PoolForGauge(arg0 common.Address) (common.Address, error) {
	return _AerodromeVoterRegistry.Contract.PoolForGauge(&_AerodromeVoterRegistry.CallOpts, arg0)
}

// PoolVote is a free data retrieval call binding the contract method 0xa86a366d.
//
// Solidity: function poolVote(uint256 , uint256 ) view returns(address)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCaller) PoolVote(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AerodromeVoterRegistry.contract.Call(opts, &out, "poolVote", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolVote is a free data retrieval call binding the contract method 0xa86a366d.
//
// Solidity: function poolVote(uint256 , uint256 ) view returns(address)
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) PoolVote(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _AerodromeVoterRegistry.Contract.PoolVote(&_AerodromeVoterRegistry.CallOpts, arg0, arg1)
}

// PoolVote is a free data retrieval call binding the contract method 0xa86a366d.
//
// Solidity: function poolVote(uint256 , uint256 ) view returns(address)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCallerSession) PoolVote(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _AerodromeVoterRegistry.Contract.PoolVote(&_AerodromeVoterRegistry.CallOpts, arg0, arg1)
}

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(address)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCaller) Pools(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AerodromeVoterRegistry.contract.Call(opts, &out, "pools", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(address)
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) Pools(arg0 *big.Int) (common.Address, error) {
	return _AerodromeVoterRegistry.Contract.Pools(&_AerodromeVoterRegistry.CallOpts, arg0)
}

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(address)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCallerSession) Pools(arg0 *big.Int) (common.Address, error) {
	return _AerodromeVoterRegistry.Contract.Pools(&_AerodromeVoterRegistry.CallOpts, arg0)
}

// TotalWeight is a free data retrieval call binding the contract method 0x96c82e57.
//
// Solidity: function totalWeight() view returns(uint256)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCaller) TotalWeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AerodromeVoterRegistry.contract.Call(opts, &out, "totalWeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalWeight is a free data retrieval call binding the contract method 0x96c82e57.
//
// Solidity: function totalWeight() view returns(uint256)
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) TotalWeight() (*big.Int, error) {
	return _AerodromeVoterRegistry.Contract.TotalWeight(&_AerodromeVoterRegistry.CallOpts)
}

// TotalWeight is a free data retrieval call binding the contract method 0x96c82e57.
//
// Solidity: function totalWeight() view returns(uint256)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCallerSession) TotalWeight() (*big.Int, error) {
	return _AerodromeVoterRegistry.Contract.TotalWeight(&_AerodromeVoterRegistry.CallOpts)
}

// UsedWeights is a free data retrieval call binding the contract method 0x79e93824.
//
// Solidity: function usedWeights(uint256 ) view returns(uint256)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCaller) UsedWeights(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AerodromeVoterRegistry.contract.Call(opts, &out, "usedWeights", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsedWeights is a free data retrieval call binding the contract method 0x79e93824.
//
// Solidity: function usedWeights(uint256 ) view returns(uint256)
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) UsedWeights(arg0 *big.Int) (*big.Int, error) {
	return _AerodromeVoterRegistry.Contract.UsedWeights(&_AerodromeVoterRegistry.CallOpts, arg0)
}

// UsedWeights is a free data retrieval call binding the contract method 0x79e93824.
//
// Solidity: function usedWeights(uint256 ) view returns(uint256)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCallerSession) UsedWeights(arg0 *big.Int) (*big.Int, error) {
	return _AerodromeVoterRegistry.Contract.UsedWeights(&_AerodromeVoterRegistry.CallOpts, arg0)
}

// Ve is a free data retrieval call binding the contract method 0x1f850716.
//
// Solidity: function ve() view returns(address)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCaller) Ve(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AerodromeVoterRegistry.contract.Call(opts, &out, "ve")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ve is a free data retrieval call binding the contract method 0x1f850716.
//
// Solidity: function ve() view returns(address)
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) Ve() (common.Address, error) {
	return _AerodromeVoterRegistry.Contract.Ve(&_AerodromeVoterRegistry.CallOpts)
}

// Ve is a free data retrieval call binding the contract method 0x1f850716.
//
// Solidity: function ve() view returns(address)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCallerSession) Ve() (common.Address, error) {
	return _AerodromeVoterRegistry.Contract.Ve(&_AerodromeVoterRegistry.CallOpts)
}

// Votes is a free data retrieval call binding the contract method 0xd23254b4.
//
// Solidity: function votes(uint256 , address ) view returns(uint256)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCaller) Votes(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AerodromeVoterRegistry.contract.Call(opts, &out, "votes", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Votes is a free data retrieval call binding the contract method 0xd23254b4.
//
// Solidity: function votes(uint256 , address ) view returns(uint256)
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) Votes(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _AerodromeVoterRegistry.Contract.Votes(&_AerodromeVoterRegistry.CallOpts, arg0, arg1)
}

// Votes is a free data retrieval call binding the contract method 0xd23254b4.
//
// Solidity: function votes(uint256 , address ) view returns(uint256)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCallerSession) Votes(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _AerodromeVoterRegistry.Contract.Votes(&_AerodromeVoterRegistry.CallOpts, arg0, arg1)
}

// Weights is a free data retrieval call binding the contract method 0xa7cac846.
//
// Solidity: function weights(address ) view returns(uint256)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCaller) Weights(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AerodromeVoterRegistry.contract.Call(opts, &out, "weights", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Weights is a free data retrieval call binding the contract method 0xa7cac846.
//
// Solidity: function weights(address ) view returns(uint256)
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) Weights(arg0 common.Address) (*big.Int, error) {
	return _AerodromeVoterRegistry.Contract.Weights(&_AerodromeVoterRegistry.CallOpts, arg0)
}

// Weights is a free data retrieval call binding the contract method 0xa7cac846.
//
// Solidity: function weights(address ) view returns(uint256)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryCallerSession) Weights(arg0 common.Address) (*big.Int, error) {
	return _AerodromeVoterRegistry.Contract.Weights(&_AerodromeVoterRegistry.CallOpts, arg0)
}

// ClaimBribes is a paid mutator transaction binding the contract method 0x7715ee75.
//
// Solidity: function claimBribes(address[] _bribes, address[][] _tokens, uint256 _tokenId) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactor) ClaimBribes(opts *bind.TransactOpts, _bribes []common.Address, _tokens [][]common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.contract.Transact(opts, "claimBribes", _bribes, _tokens, _tokenId)
}

// ClaimBribes is a paid mutator transaction binding the contract method 0x7715ee75.
//
// Solidity: function claimBribes(address[] _bribes, address[][] _tokens, uint256 _tokenId) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) ClaimBribes(_bribes []common.Address, _tokens [][]common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.ClaimBribes(&_AerodromeVoterRegistry.TransactOpts, _bribes, _tokens, _tokenId)
}

// ClaimBribes is a paid mutator transaction binding the contract method 0x7715ee75.
//
// Solidity: function claimBribes(address[] _bribes, address[][] _tokens, uint256 _tokenId) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactorSession) ClaimBribes(_bribes []common.Address, _tokens [][]common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.ClaimBribes(&_AerodromeVoterRegistry.TransactOpts, _bribes, _tokens, _tokenId)
}

// ClaimFees is a paid mutator transaction binding the contract method 0x666256aa.
//
// Solidity: function claimFees(address[] _fees, address[][] _tokens, uint256 _tokenId) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactor) ClaimFees(opts *bind.TransactOpts, _fees []common.Address, _tokens [][]common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.contract.Transact(opts, "claimFees", _fees, _tokens, _tokenId)
}

// ClaimFees is a paid mutator transaction binding the contract method 0x666256aa.
//
// Solidity: function claimFees(address[] _fees, address[][] _tokens, uint256 _tokenId) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) ClaimFees(_fees []common.Address, _tokens [][]common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.ClaimFees(&_AerodromeVoterRegistry.TransactOpts, _fees, _tokens, _tokenId)
}

// ClaimFees is a paid mutator transaction binding the contract method 0x666256aa.
//
// Solidity: function claimFees(address[] _fees, address[][] _tokens, uint256 _tokenId) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactorSession) ClaimFees(_fees []common.Address, _tokens [][]common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.ClaimFees(&_AerodromeVoterRegistry.TransactOpts, _fees, _tokens, _tokenId)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xf9f031df.
//
// Solidity: function claimRewards(address[] _gauges) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactor) ClaimRewards(opts *bind.TransactOpts, _gauges []common.Address) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.contract.Transact(opts, "claimRewards", _gauges)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xf9f031df.
//
// Solidity: function claimRewards(address[] _gauges) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) ClaimRewards(_gauges []common.Address) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.ClaimRewards(&_AerodromeVoterRegistry.TransactOpts, _gauges)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xf9f031df.
//
// Solidity: function claimRewards(address[] _gauges) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactorSession) ClaimRewards(_gauges []common.Address) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.ClaimRewards(&_AerodromeVoterRegistry.TransactOpts, _gauges)
}

// CreateGauge is a paid mutator transaction binding the contract method 0x794cea3c.
//
// Solidity: function createGauge(address _poolFactory, address _pool) returns(address)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactor) CreateGauge(opts *bind.TransactOpts, _poolFactory common.Address, _pool common.Address) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.contract.Transact(opts, "createGauge", _poolFactory, _pool)
}

// CreateGauge is a paid mutator transaction binding the contract method 0x794cea3c.
//
// Solidity: function createGauge(address _poolFactory, address _pool) returns(address)
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) CreateGauge(_poolFactory common.Address, _pool common.Address) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.CreateGauge(&_AerodromeVoterRegistry.TransactOpts, _poolFactory, _pool)
}

// CreateGauge is a paid mutator transaction binding the contract method 0x794cea3c.
//
// Solidity: function createGauge(address _poolFactory, address _pool) returns(address)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactorSession) CreateGauge(_poolFactory common.Address, _pool common.Address) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.CreateGauge(&_AerodromeVoterRegistry.TransactOpts, _poolFactory, _pool)
}

// DepositManaged is a paid mutator transaction binding the contract method 0xe0c11f9a.
//
// Solidity: function depositManaged(uint256 _tokenId, uint256 _mTokenId) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactor) DepositManaged(opts *bind.TransactOpts, _tokenId *big.Int, _mTokenId *big.Int) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.contract.Transact(opts, "depositManaged", _tokenId, _mTokenId)
}

// DepositManaged is a paid mutator transaction binding the contract method 0xe0c11f9a.
//
// Solidity: function depositManaged(uint256 _tokenId, uint256 _mTokenId) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) DepositManaged(_tokenId *big.Int, _mTokenId *big.Int) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.DepositManaged(&_AerodromeVoterRegistry.TransactOpts, _tokenId, _mTokenId)
}

// DepositManaged is a paid mutator transaction binding the contract method 0xe0c11f9a.
//
// Solidity: function depositManaged(uint256 _tokenId, uint256 _mTokenId) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactorSession) DepositManaged(_tokenId *big.Int, _mTokenId *big.Int) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.DepositManaged(&_AerodromeVoterRegistry.TransactOpts, _tokenId, _mTokenId)
}

// Distribute is a paid mutator transaction binding the contract method 0x6138889b.
//
// Solidity: function distribute(address[] _gauges) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactor) Distribute(opts *bind.TransactOpts, _gauges []common.Address) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.contract.Transact(opts, "distribute", _gauges)
}

// Distribute is a paid mutator transaction binding the contract method 0x6138889b.
//
// Solidity: function distribute(address[] _gauges) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) Distribute(_gauges []common.Address) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.Distribute(&_AerodromeVoterRegistry.TransactOpts, _gauges)
}

// Distribute is a paid mutator transaction binding the contract method 0x6138889b.
//
// Solidity: function distribute(address[] _gauges) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactorSession) Distribute(_gauges []common.Address) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.Distribute(&_AerodromeVoterRegistry.TransactOpts, _gauges)
}

// Distribute0 is a paid mutator transaction binding the contract method 0x7625391a.
//
// Solidity: function distribute(uint256 _start, uint256 _finish) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactor) Distribute0(opts *bind.TransactOpts, _start *big.Int, _finish *big.Int) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.contract.Transact(opts, "distribute0", _start, _finish)
}

// Distribute0 is a paid mutator transaction binding the contract method 0x7625391a.
//
// Solidity: function distribute(uint256 _start, uint256 _finish) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) Distribute0(_start *big.Int, _finish *big.Int) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.Distribute0(&_AerodromeVoterRegistry.TransactOpts, _start, _finish)
}

// Distribute0 is a paid mutator transaction binding the contract method 0x7625391a.
//
// Solidity: function distribute(uint256 _start, uint256 _finish) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactorSession) Distribute0(_start *big.Int, _finish *big.Int) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.Distribute0(&_AerodromeVoterRegistry.TransactOpts, _start, _finish)
}

// Initialize is a paid mutator transaction binding the contract method 0x462d0b2e.
//
// Solidity: function initialize(address[] _tokens, address _minter) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactor) Initialize(opts *bind.TransactOpts, _tokens []common.Address, _minter common.Address) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.contract.Transact(opts, "initialize", _tokens, _minter)
}

// Initialize is a paid mutator transaction binding the contract method 0x462d0b2e.
//
// Solidity: function initialize(address[] _tokens, address _minter) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) Initialize(_tokens []common.Address, _minter common.Address) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.Initialize(&_AerodromeVoterRegistry.TransactOpts, _tokens, _minter)
}

// Initialize is a paid mutator transaction binding the contract method 0x462d0b2e.
//
// Solidity: function initialize(address[] _tokens, address _minter) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactorSession) Initialize(_tokens []common.Address, _minter common.Address) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.Initialize(&_AerodromeVoterRegistry.TransactOpts, _tokens, _minter)
}

// KillGauge is a paid mutator transaction binding the contract method 0x992a7933.
//
// Solidity: function killGauge(address _gauge) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactor) KillGauge(opts *bind.TransactOpts, _gauge common.Address) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.contract.Transact(opts, "killGauge", _gauge)
}

// KillGauge is a paid mutator transaction binding the contract method 0x992a7933.
//
// Solidity: function killGauge(address _gauge) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) KillGauge(_gauge common.Address) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.KillGauge(&_AerodromeVoterRegistry.TransactOpts, _gauge)
}

// KillGauge is a paid mutator transaction binding the contract method 0x992a7933.
//
// Solidity: function killGauge(address _gauge) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactorSession) KillGauge(_gauge common.Address) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.KillGauge(&_AerodromeVoterRegistry.TransactOpts, _gauge)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 _amount) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactor) NotifyRewardAmount(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.contract.Transact(opts, "notifyRewardAmount", _amount)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 _amount) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) NotifyRewardAmount(_amount *big.Int) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.NotifyRewardAmount(&_AerodromeVoterRegistry.TransactOpts, _amount)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 _amount) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactorSession) NotifyRewardAmount(_amount *big.Int) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.NotifyRewardAmount(&_AerodromeVoterRegistry.TransactOpts, _amount)
}

// Poke is a paid mutator transaction binding the contract method 0x32145f90.
//
// Solidity: function poke(uint256 _tokenId) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactor) Poke(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.contract.Transact(opts, "poke", _tokenId)
}

// Poke is a paid mutator transaction binding the contract method 0x32145f90.
//
// Solidity: function poke(uint256 _tokenId) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) Poke(_tokenId *big.Int) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.Poke(&_AerodromeVoterRegistry.TransactOpts, _tokenId)
}

// Poke is a paid mutator transaction binding the contract method 0x32145f90.
//
// Solidity: function poke(uint256 _tokenId) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactorSession) Poke(_tokenId *big.Int) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.Poke(&_AerodromeVoterRegistry.TransactOpts, _tokenId)
}

// Reset is a paid mutator transaction binding the contract method 0x310bd74b.
//
// Solidity: function reset(uint256 _tokenId) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactor) Reset(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.contract.Transact(opts, "reset", _tokenId)
}

// Reset is a paid mutator transaction binding the contract method 0x310bd74b.
//
// Solidity: function reset(uint256 _tokenId) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) Reset(_tokenId *big.Int) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.Reset(&_AerodromeVoterRegistry.TransactOpts, _tokenId)
}

// Reset is a paid mutator transaction binding the contract method 0x310bd74b.
//
// Solidity: function reset(uint256 _tokenId) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactorSession) Reset(_tokenId *big.Int) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.Reset(&_AerodromeVoterRegistry.TransactOpts, _tokenId)
}

// ReviveGauge is a paid mutator transaction binding the contract method 0x9f06247b.
//
// Solidity: function reviveGauge(address _gauge) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactor) ReviveGauge(opts *bind.TransactOpts, _gauge common.Address) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.contract.Transact(opts, "reviveGauge", _gauge)
}

// ReviveGauge is a paid mutator transaction binding the contract method 0x9f06247b.
//
// Solidity: function reviveGauge(address _gauge) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) ReviveGauge(_gauge common.Address) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.ReviveGauge(&_AerodromeVoterRegistry.TransactOpts, _gauge)
}

// ReviveGauge is a paid mutator transaction binding the contract method 0x9f06247b.
//
// Solidity: function reviveGauge(address _gauge) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactorSession) ReviveGauge(_gauge common.Address) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.ReviveGauge(&_AerodromeVoterRegistry.TransactOpts, _gauge)
}

// SetEmergencyCouncil is a paid mutator transaction binding the contract method 0xe586875f.
//
// Solidity: function setEmergencyCouncil(address _council) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactor) SetEmergencyCouncil(opts *bind.TransactOpts, _council common.Address) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.contract.Transact(opts, "setEmergencyCouncil", _council)
}

// SetEmergencyCouncil is a paid mutator transaction binding the contract method 0xe586875f.
//
// Solidity: function setEmergencyCouncil(address _council) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) SetEmergencyCouncil(_council common.Address) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.SetEmergencyCouncil(&_AerodromeVoterRegistry.TransactOpts, _council)
}

// SetEmergencyCouncil is a paid mutator transaction binding the contract method 0xe586875f.
//
// Solidity: function setEmergencyCouncil(address _council) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactorSession) SetEmergencyCouncil(_council common.Address) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.SetEmergencyCouncil(&_AerodromeVoterRegistry.TransactOpts, _council)
}

// SetEpochGovernor is a paid mutator transaction binding the contract method 0x598d521b.
//
// Solidity: function setEpochGovernor(address _epochGovernor) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactor) SetEpochGovernor(opts *bind.TransactOpts, _epochGovernor common.Address) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.contract.Transact(opts, "setEpochGovernor", _epochGovernor)
}

// SetEpochGovernor is a paid mutator transaction binding the contract method 0x598d521b.
//
// Solidity: function setEpochGovernor(address _epochGovernor) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) SetEpochGovernor(_epochGovernor common.Address) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.SetEpochGovernor(&_AerodromeVoterRegistry.TransactOpts, _epochGovernor)
}

// SetEpochGovernor is a paid mutator transaction binding the contract method 0x598d521b.
//
// Solidity: function setEpochGovernor(address _epochGovernor) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactorSession) SetEpochGovernor(_epochGovernor common.Address) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.SetEpochGovernor(&_AerodromeVoterRegistry.TransactOpts, _epochGovernor)
}

// SetGovernor is a paid mutator transaction binding the contract method 0xc42cf535.
//
// Solidity: function setGovernor(address _governor) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactor) SetGovernor(opts *bind.TransactOpts, _governor common.Address) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.contract.Transact(opts, "setGovernor", _governor)
}

// SetGovernor is a paid mutator transaction binding the contract method 0xc42cf535.
//
// Solidity: function setGovernor(address _governor) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) SetGovernor(_governor common.Address) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.SetGovernor(&_AerodromeVoterRegistry.TransactOpts, _governor)
}

// SetGovernor is a paid mutator transaction binding the contract method 0xc42cf535.
//
// Solidity: function setGovernor(address _governor) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactorSession) SetGovernor(_governor common.Address) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.SetGovernor(&_AerodromeVoterRegistry.TransactOpts, _governor)
}

// SetMaxVotingNum is a paid mutator transaction binding the contract method 0x30331b2f.
//
// Solidity: function setMaxVotingNum(uint256 _maxVotingNum) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactor) SetMaxVotingNum(opts *bind.TransactOpts, _maxVotingNum *big.Int) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.contract.Transact(opts, "setMaxVotingNum", _maxVotingNum)
}

// SetMaxVotingNum is a paid mutator transaction binding the contract method 0x30331b2f.
//
// Solidity: function setMaxVotingNum(uint256 _maxVotingNum) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) SetMaxVotingNum(_maxVotingNum *big.Int) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.SetMaxVotingNum(&_AerodromeVoterRegistry.TransactOpts, _maxVotingNum)
}

// SetMaxVotingNum is a paid mutator transaction binding the contract method 0x30331b2f.
//
// Solidity: function setMaxVotingNum(uint256 _maxVotingNum) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactorSession) SetMaxVotingNum(_maxVotingNum *big.Int) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.SetMaxVotingNum(&_AerodromeVoterRegistry.TransactOpts, _maxVotingNum)
}

// UpdateFor is a paid mutator transaction binding the contract method 0x0e0a5968.
//
// Solidity: function updateFor(address _gauge) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactor) UpdateFor(opts *bind.TransactOpts, _gauge common.Address) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.contract.Transact(opts, "updateFor", _gauge)
}

// UpdateFor is a paid mutator transaction binding the contract method 0x0e0a5968.
//
// Solidity: function updateFor(address _gauge) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) UpdateFor(_gauge common.Address) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.UpdateFor(&_AerodromeVoterRegistry.TransactOpts, _gauge)
}

// UpdateFor is a paid mutator transaction binding the contract method 0x0e0a5968.
//
// Solidity: function updateFor(address _gauge) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactorSession) UpdateFor(_gauge common.Address) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.UpdateFor(&_AerodromeVoterRegistry.TransactOpts, _gauge)
}

// UpdateFor0 is a paid mutator transaction binding the contract method 0xc9ff6f4d.
//
// Solidity: function updateFor(uint256 start, uint256 end) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactor) UpdateFor0(opts *bind.TransactOpts, start *big.Int, end *big.Int) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.contract.Transact(opts, "updateFor0", start, end)
}

// UpdateFor0 is a paid mutator transaction binding the contract method 0xc9ff6f4d.
//
// Solidity: function updateFor(uint256 start, uint256 end) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) UpdateFor0(start *big.Int, end *big.Int) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.UpdateFor0(&_AerodromeVoterRegistry.TransactOpts, start, end)
}

// UpdateFor0 is a paid mutator transaction binding the contract method 0xc9ff6f4d.
//
// Solidity: function updateFor(uint256 start, uint256 end) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactorSession) UpdateFor0(start *big.Int, end *big.Int) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.UpdateFor0(&_AerodromeVoterRegistry.TransactOpts, start, end)
}

// UpdateFor1 is a paid mutator transaction binding the contract method 0xd560b0d7.
//
// Solidity: function updateFor(address[] _gauges) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactor) UpdateFor1(opts *bind.TransactOpts, _gauges []common.Address) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.contract.Transact(opts, "updateFor1", _gauges)
}

// UpdateFor1 is a paid mutator transaction binding the contract method 0xd560b0d7.
//
// Solidity: function updateFor(address[] _gauges) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) UpdateFor1(_gauges []common.Address) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.UpdateFor1(&_AerodromeVoterRegistry.TransactOpts, _gauges)
}

// UpdateFor1 is a paid mutator transaction binding the contract method 0xd560b0d7.
//
// Solidity: function updateFor(address[] _gauges) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactorSession) UpdateFor1(_gauges []common.Address) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.UpdateFor1(&_AerodromeVoterRegistry.TransactOpts, _gauges)
}

// Vote is a paid mutator transaction binding the contract method 0x7ac09bf7.
//
// Solidity: function vote(uint256 _tokenId, address[] _poolVote, uint256[] _weights) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactor) Vote(opts *bind.TransactOpts, _tokenId *big.Int, _poolVote []common.Address, _weights []*big.Int) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.contract.Transact(opts, "vote", _tokenId, _poolVote, _weights)
}

// Vote is a paid mutator transaction binding the contract method 0x7ac09bf7.
//
// Solidity: function vote(uint256 _tokenId, address[] _poolVote, uint256[] _weights) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) Vote(_tokenId *big.Int, _poolVote []common.Address, _weights []*big.Int) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.Vote(&_AerodromeVoterRegistry.TransactOpts, _tokenId, _poolVote, _weights)
}

// Vote is a paid mutator transaction binding the contract method 0x7ac09bf7.
//
// Solidity: function vote(uint256 _tokenId, address[] _poolVote, uint256[] _weights) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactorSession) Vote(_tokenId *big.Int, _poolVote []common.Address, _weights []*big.Int) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.Vote(&_AerodromeVoterRegistry.TransactOpts, _tokenId, _poolVote, _weights)
}

// WhitelistNFT is a paid mutator transaction binding the contract method 0xe2819d5c.
//
// Solidity: function whitelistNFT(uint256 _tokenId, bool _bool) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactor) WhitelistNFT(opts *bind.TransactOpts, _tokenId *big.Int, _bool bool) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.contract.Transact(opts, "whitelistNFT", _tokenId, _bool)
}

// WhitelistNFT is a paid mutator transaction binding the contract method 0xe2819d5c.
//
// Solidity: function whitelistNFT(uint256 _tokenId, bool _bool) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) WhitelistNFT(_tokenId *big.Int, _bool bool) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.WhitelistNFT(&_AerodromeVoterRegistry.TransactOpts, _tokenId, _bool)
}

// WhitelistNFT is a paid mutator transaction binding the contract method 0xe2819d5c.
//
// Solidity: function whitelistNFT(uint256 _tokenId, bool _bool) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactorSession) WhitelistNFT(_tokenId *big.Int, _bool bool) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.WhitelistNFT(&_AerodromeVoterRegistry.TransactOpts, _tokenId, _bool)
}

// WhitelistToken is a paid mutator transaction binding the contract method 0x0ffb1d8b.
//
// Solidity: function whitelistToken(address _token, bool _bool) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactor) WhitelistToken(opts *bind.TransactOpts, _token common.Address, _bool bool) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.contract.Transact(opts, "whitelistToken", _token, _bool)
}

// WhitelistToken is a paid mutator transaction binding the contract method 0x0ffb1d8b.
//
// Solidity: function whitelistToken(address _token, bool _bool) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) WhitelistToken(_token common.Address, _bool bool) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.WhitelistToken(&_AerodromeVoterRegistry.TransactOpts, _token, _bool)
}

// WhitelistToken is a paid mutator transaction binding the contract method 0x0ffb1d8b.
//
// Solidity: function whitelistToken(address _token, bool _bool) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactorSession) WhitelistToken(_token common.Address, _bool bool) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.WhitelistToken(&_AerodromeVoterRegistry.TransactOpts, _token, _bool)
}

// WithdrawManaged is a paid mutator transaction binding the contract method 0x370fb5fa.
//
// Solidity: function withdrawManaged(uint256 _tokenId) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactor) WithdrawManaged(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.contract.Transact(opts, "withdrawManaged", _tokenId)
}

// WithdrawManaged is a paid mutator transaction binding the contract method 0x370fb5fa.
//
// Solidity: function withdrawManaged(uint256 _tokenId) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistrySession) WithdrawManaged(_tokenId *big.Int) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.WithdrawManaged(&_AerodromeVoterRegistry.TransactOpts, _tokenId)
}

// WithdrawManaged is a paid mutator transaction binding the contract method 0x370fb5fa.
//
// Solidity: function withdrawManaged(uint256 _tokenId) returns()
func (_AerodromeVoterRegistry *AerodromeVoterRegistryTransactorSession) WithdrawManaged(_tokenId *big.Int) (*types.Transaction, error) {
	return _AerodromeVoterRegistry.Contract.WithdrawManaged(&_AerodromeVoterRegistry.TransactOpts, _tokenId)
}

// AerodromeVoterRegistryAbstainedIterator is returned from FilterAbstained and is used to iterate over the raw logs and unpacked data for Abstained events raised by the AerodromeVoterRegistry contract.
type AerodromeVoterRegistryAbstainedIterator struct {
	Event *AerodromeVoterRegistryAbstained // Event containing the contract specifics and raw log

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
func (it *AerodromeVoterRegistryAbstainedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AerodromeVoterRegistryAbstained)
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
		it.Event = new(AerodromeVoterRegistryAbstained)
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
func (it *AerodromeVoterRegistryAbstainedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AerodromeVoterRegistryAbstainedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AerodromeVoterRegistryAbstained represents a Abstained event raised by the AerodromeVoterRegistry contract.
type AerodromeVoterRegistryAbstained struct {
	Voter       common.Address
	Pool        common.Address
	TokenId     *big.Int
	Weight      *big.Int
	TotalWeight *big.Int
	Timestamp   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAbstained is a free log retrieval operation binding the contract event 0xadab630928b1d46214641293704a312ee7ad87e03ae14a7fd95e7308b93998df.
//
// Solidity: event Abstained(address indexed voter, address indexed pool, uint256 indexed tokenId, uint256 weight, uint256 totalWeight, uint256 timestamp)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryFilterer) FilterAbstained(opts *bind.FilterOpts, voter []common.Address, pool []common.Address, tokenId []*big.Int) (*AerodromeVoterRegistryAbstainedIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _AerodromeVoterRegistry.contract.FilterLogs(opts, "Abstained", voterRule, poolRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &AerodromeVoterRegistryAbstainedIterator{contract: _AerodromeVoterRegistry.contract, event: "Abstained", logs: logs, sub: sub}, nil
}

// WatchAbstained is a free log subscription operation binding the contract event 0xadab630928b1d46214641293704a312ee7ad87e03ae14a7fd95e7308b93998df.
//
// Solidity: event Abstained(address indexed voter, address indexed pool, uint256 indexed tokenId, uint256 weight, uint256 totalWeight, uint256 timestamp)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryFilterer) WatchAbstained(opts *bind.WatchOpts, sink chan<- *AerodromeVoterRegistryAbstained, voter []common.Address, pool []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _AerodromeVoterRegistry.contract.WatchLogs(opts, "Abstained", voterRule, poolRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AerodromeVoterRegistryAbstained)
				if err := _AerodromeVoterRegistry.contract.UnpackLog(event, "Abstained", log); err != nil {
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

// ParseAbstained is a log parse operation binding the contract event 0xadab630928b1d46214641293704a312ee7ad87e03ae14a7fd95e7308b93998df.
//
// Solidity: event Abstained(address indexed voter, address indexed pool, uint256 indexed tokenId, uint256 weight, uint256 totalWeight, uint256 timestamp)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryFilterer) ParseAbstained(log types.Log) (*AerodromeVoterRegistryAbstained, error) {
	event := new(AerodromeVoterRegistryAbstained)
	if err := _AerodromeVoterRegistry.contract.UnpackLog(event, "Abstained", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AerodromeVoterRegistryDistributeRewardIterator is returned from FilterDistributeReward and is used to iterate over the raw logs and unpacked data for DistributeReward events raised by the AerodromeVoterRegistry contract.
type AerodromeVoterRegistryDistributeRewardIterator struct {
	Event *AerodromeVoterRegistryDistributeReward // Event containing the contract specifics and raw log

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
func (it *AerodromeVoterRegistryDistributeRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AerodromeVoterRegistryDistributeReward)
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
		it.Event = new(AerodromeVoterRegistryDistributeReward)
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
func (it *AerodromeVoterRegistryDistributeRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AerodromeVoterRegistryDistributeRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AerodromeVoterRegistryDistributeReward represents a DistributeReward event raised by the AerodromeVoterRegistry contract.
type AerodromeVoterRegistryDistributeReward struct {
	Sender common.Address
	Gauge  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDistributeReward is a free log retrieval operation binding the contract event 0x4fa9693cae526341d334e2862ca2413b2e503f1266255f9e0869fb36e6d89b17.
//
// Solidity: event DistributeReward(address indexed sender, address indexed gauge, uint256 amount)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryFilterer) FilterDistributeReward(opts *bind.FilterOpts, sender []common.Address, gauge []common.Address) (*AerodromeVoterRegistryDistributeRewardIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var gaugeRule []interface{}
	for _, gaugeItem := range gauge {
		gaugeRule = append(gaugeRule, gaugeItem)
	}

	logs, sub, err := _AerodromeVoterRegistry.contract.FilterLogs(opts, "DistributeReward", senderRule, gaugeRule)
	if err != nil {
		return nil, err
	}
	return &AerodromeVoterRegistryDistributeRewardIterator{contract: _AerodromeVoterRegistry.contract, event: "DistributeReward", logs: logs, sub: sub}, nil
}

// WatchDistributeReward is a free log subscription operation binding the contract event 0x4fa9693cae526341d334e2862ca2413b2e503f1266255f9e0869fb36e6d89b17.
//
// Solidity: event DistributeReward(address indexed sender, address indexed gauge, uint256 amount)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryFilterer) WatchDistributeReward(opts *bind.WatchOpts, sink chan<- *AerodromeVoterRegistryDistributeReward, sender []common.Address, gauge []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var gaugeRule []interface{}
	for _, gaugeItem := range gauge {
		gaugeRule = append(gaugeRule, gaugeItem)
	}

	logs, sub, err := _AerodromeVoterRegistry.contract.WatchLogs(opts, "DistributeReward", senderRule, gaugeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AerodromeVoterRegistryDistributeReward)
				if err := _AerodromeVoterRegistry.contract.UnpackLog(event, "DistributeReward", log); err != nil {
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

// ParseDistributeReward is a log parse operation binding the contract event 0x4fa9693cae526341d334e2862ca2413b2e503f1266255f9e0869fb36e6d89b17.
//
// Solidity: event DistributeReward(address indexed sender, address indexed gauge, uint256 amount)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryFilterer) ParseDistributeReward(log types.Log) (*AerodromeVoterRegistryDistributeReward, error) {
	event := new(AerodromeVoterRegistryDistributeReward)
	if err := _AerodromeVoterRegistry.contract.UnpackLog(event, "DistributeReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AerodromeVoterRegistryGaugeCreatedIterator is returned from FilterGaugeCreated and is used to iterate over the raw logs and unpacked data for GaugeCreated events raised by the AerodromeVoterRegistry contract.
type AerodromeVoterRegistryGaugeCreatedIterator struct {
	Event *AerodromeVoterRegistryGaugeCreated // Event containing the contract specifics and raw log

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
func (it *AerodromeVoterRegistryGaugeCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AerodromeVoterRegistryGaugeCreated)
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
		it.Event = new(AerodromeVoterRegistryGaugeCreated)
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
func (it *AerodromeVoterRegistryGaugeCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AerodromeVoterRegistryGaugeCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AerodromeVoterRegistryGaugeCreated represents a GaugeCreated event raised by the AerodromeVoterRegistry contract.
type AerodromeVoterRegistryGaugeCreated struct {
	PoolFactory          common.Address
	VotingRewardsFactory common.Address
	GaugeFactory         common.Address
	Pool                 common.Address
	BribeVotingReward    common.Address
	FeeVotingReward      common.Address
	Gauge                common.Address
	Creator              common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterGaugeCreated is a free log retrieval operation binding the contract event 0xef9f7d1ffff3b249c6b9bf2528499e935f7d96bb6d6ec4e7da504d1d3c6279e1.
//
// Solidity: event GaugeCreated(address indexed poolFactory, address indexed votingRewardsFactory, address indexed gaugeFactory, address pool, address bribeVotingReward, address feeVotingReward, address gauge, address creator)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryFilterer) FilterGaugeCreated(opts *bind.FilterOpts, poolFactory []common.Address, votingRewardsFactory []common.Address, gaugeFactory []common.Address) (*AerodromeVoterRegistryGaugeCreatedIterator, error) {

	var poolFactoryRule []interface{}
	for _, poolFactoryItem := range poolFactory {
		poolFactoryRule = append(poolFactoryRule, poolFactoryItem)
	}
	var votingRewardsFactoryRule []interface{}
	for _, votingRewardsFactoryItem := range votingRewardsFactory {
		votingRewardsFactoryRule = append(votingRewardsFactoryRule, votingRewardsFactoryItem)
	}
	var gaugeFactoryRule []interface{}
	for _, gaugeFactoryItem := range gaugeFactory {
		gaugeFactoryRule = append(gaugeFactoryRule, gaugeFactoryItem)
	}

	logs, sub, err := _AerodromeVoterRegistry.contract.FilterLogs(opts, "GaugeCreated", poolFactoryRule, votingRewardsFactoryRule, gaugeFactoryRule)
	if err != nil {
		return nil, err
	}
	return &AerodromeVoterRegistryGaugeCreatedIterator{contract: _AerodromeVoterRegistry.contract, event: "GaugeCreated", logs: logs, sub: sub}, nil
}

// WatchGaugeCreated is a free log subscription operation binding the contract event 0xef9f7d1ffff3b249c6b9bf2528499e935f7d96bb6d6ec4e7da504d1d3c6279e1.
//
// Solidity: event GaugeCreated(address indexed poolFactory, address indexed votingRewardsFactory, address indexed gaugeFactory, address pool, address bribeVotingReward, address feeVotingReward, address gauge, address creator)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryFilterer) WatchGaugeCreated(opts *bind.WatchOpts, sink chan<- *AerodromeVoterRegistryGaugeCreated, poolFactory []common.Address, votingRewardsFactory []common.Address, gaugeFactory []common.Address) (event.Subscription, error) {

	var poolFactoryRule []interface{}
	for _, poolFactoryItem := range poolFactory {
		poolFactoryRule = append(poolFactoryRule, poolFactoryItem)
	}
	var votingRewardsFactoryRule []interface{}
	for _, votingRewardsFactoryItem := range votingRewardsFactory {
		votingRewardsFactoryRule = append(votingRewardsFactoryRule, votingRewardsFactoryItem)
	}
	var gaugeFactoryRule []interface{}
	for _, gaugeFactoryItem := range gaugeFactory {
		gaugeFactoryRule = append(gaugeFactoryRule, gaugeFactoryItem)
	}

	logs, sub, err := _AerodromeVoterRegistry.contract.WatchLogs(opts, "GaugeCreated", poolFactoryRule, votingRewardsFactoryRule, gaugeFactoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AerodromeVoterRegistryGaugeCreated)
				if err := _AerodromeVoterRegistry.contract.UnpackLog(event, "GaugeCreated", log); err != nil {
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

// ParseGaugeCreated is a log parse operation binding the contract event 0xef9f7d1ffff3b249c6b9bf2528499e935f7d96bb6d6ec4e7da504d1d3c6279e1.
//
// Solidity: event GaugeCreated(address indexed poolFactory, address indexed votingRewardsFactory, address indexed gaugeFactory, address pool, address bribeVotingReward, address feeVotingReward, address gauge, address creator)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryFilterer) ParseGaugeCreated(log types.Log) (*AerodromeVoterRegistryGaugeCreated, error) {
	event := new(AerodromeVoterRegistryGaugeCreated)
	if err := _AerodromeVoterRegistry.contract.UnpackLog(event, "GaugeCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AerodromeVoterRegistryGaugeKilledIterator is returned from FilterGaugeKilled and is used to iterate over the raw logs and unpacked data for GaugeKilled events raised by the AerodromeVoterRegistry contract.
type AerodromeVoterRegistryGaugeKilledIterator struct {
	Event *AerodromeVoterRegistryGaugeKilled // Event containing the contract specifics and raw log

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
func (it *AerodromeVoterRegistryGaugeKilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AerodromeVoterRegistryGaugeKilled)
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
		it.Event = new(AerodromeVoterRegistryGaugeKilled)
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
func (it *AerodromeVoterRegistryGaugeKilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AerodromeVoterRegistryGaugeKilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AerodromeVoterRegistryGaugeKilled represents a GaugeKilled event raised by the AerodromeVoterRegistry contract.
type AerodromeVoterRegistryGaugeKilled struct {
	Gauge common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGaugeKilled is a free log retrieval operation binding the contract event 0x04a5d3f5d80d22d9345acc80618f4a4e7e663cf9e1aed23b57d975acec002ba7.
//
// Solidity: event GaugeKilled(address indexed gauge)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryFilterer) FilterGaugeKilled(opts *bind.FilterOpts, gauge []common.Address) (*AerodromeVoterRegistryGaugeKilledIterator, error) {

	var gaugeRule []interface{}
	for _, gaugeItem := range gauge {
		gaugeRule = append(gaugeRule, gaugeItem)
	}

	logs, sub, err := _AerodromeVoterRegistry.contract.FilterLogs(opts, "GaugeKilled", gaugeRule)
	if err != nil {
		return nil, err
	}
	return &AerodromeVoterRegistryGaugeKilledIterator{contract: _AerodromeVoterRegistry.contract, event: "GaugeKilled", logs: logs, sub: sub}, nil
}

// WatchGaugeKilled is a free log subscription operation binding the contract event 0x04a5d3f5d80d22d9345acc80618f4a4e7e663cf9e1aed23b57d975acec002ba7.
//
// Solidity: event GaugeKilled(address indexed gauge)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryFilterer) WatchGaugeKilled(opts *bind.WatchOpts, sink chan<- *AerodromeVoterRegistryGaugeKilled, gauge []common.Address) (event.Subscription, error) {

	var gaugeRule []interface{}
	for _, gaugeItem := range gauge {
		gaugeRule = append(gaugeRule, gaugeItem)
	}

	logs, sub, err := _AerodromeVoterRegistry.contract.WatchLogs(opts, "GaugeKilled", gaugeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AerodromeVoterRegistryGaugeKilled)
				if err := _AerodromeVoterRegistry.contract.UnpackLog(event, "GaugeKilled", log); err != nil {
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

// ParseGaugeKilled is a log parse operation binding the contract event 0x04a5d3f5d80d22d9345acc80618f4a4e7e663cf9e1aed23b57d975acec002ba7.
//
// Solidity: event GaugeKilled(address indexed gauge)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryFilterer) ParseGaugeKilled(log types.Log) (*AerodromeVoterRegistryGaugeKilled, error) {
	event := new(AerodromeVoterRegistryGaugeKilled)
	if err := _AerodromeVoterRegistry.contract.UnpackLog(event, "GaugeKilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AerodromeVoterRegistryGaugeRevivedIterator is returned from FilterGaugeRevived and is used to iterate over the raw logs and unpacked data for GaugeRevived events raised by the AerodromeVoterRegistry contract.
type AerodromeVoterRegistryGaugeRevivedIterator struct {
	Event *AerodromeVoterRegistryGaugeRevived // Event containing the contract specifics and raw log

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
func (it *AerodromeVoterRegistryGaugeRevivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AerodromeVoterRegistryGaugeRevived)
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
		it.Event = new(AerodromeVoterRegistryGaugeRevived)
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
func (it *AerodromeVoterRegistryGaugeRevivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AerodromeVoterRegistryGaugeRevivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AerodromeVoterRegistryGaugeRevived represents a GaugeRevived event raised by the AerodromeVoterRegistry contract.
type AerodromeVoterRegistryGaugeRevived struct {
	Gauge common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGaugeRevived is a free log retrieval operation binding the contract event 0xed18e9faa3dccfd8aa45f69c4de40546b2ca9cccc4538a2323531656516db1aa.
//
// Solidity: event GaugeRevived(address indexed gauge)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryFilterer) FilterGaugeRevived(opts *bind.FilterOpts, gauge []common.Address) (*AerodromeVoterRegistryGaugeRevivedIterator, error) {

	var gaugeRule []interface{}
	for _, gaugeItem := range gauge {
		gaugeRule = append(gaugeRule, gaugeItem)
	}

	logs, sub, err := _AerodromeVoterRegistry.contract.FilterLogs(opts, "GaugeRevived", gaugeRule)
	if err != nil {
		return nil, err
	}
	return &AerodromeVoterRegistryGaugeRevivedIterator{contract: _AerodromeVoterRegistry.contract, event: "GaugeRevived", logs: logs, sub: sub}, nil
}

// WatchGaugeRevived is a free log subscription operation binding the contract event 0xed18e9faa3dccfd8aa45f69c4de40546b2ca9cccc4538a2323531656516db1aa.
//
// Solidity: event GaugeRevived(address indexed gauge)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryFilterer) WatchGaugeRevived(opts *bind.WatchOpts, sink chan<- *AerodromeVoterRegistryGaugeRevived, gauge []common.Address) (event.Subscription, error) {

	var gaugeRule []interface{}
	for _, gaugeItem := range gauge {
		gaugeRule = append(gaugeRule, gaugeItem)
	}

	logs, sub, err := _AerodromeVoterRegistry.contract.WatchLogs(opts, "GaugeRevived", gaugeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AerodromeVoterRegistryGaugeRevived)
				if err := _AerodromeVoterRegistry.contract.UnpackLog(event, "GaugeRevived", log); err != nil {
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

// ParseGaugeRevived is a log parse operation binding the contract event 0xed18e9faa3dccfd8aa45f69c4de40546b2ca9cccc4538a2323531656516db1aa.
//
// Solidity: event GaugeRevived(address indexed gauge)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryFilterer) ParseGaugeRevived(log types.Log) (*AerodromeVoterRegistryGaugeRevived, error) {
	event := new(AerodromeVoterRegistryGaugeRevived)
	if err := _AerodromeVoterRegistry.contract.UnpackLog(event, "GaugeRevived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AerodromeVoterRegistryNotifyRewardIterator is returned from FilterNotifyReward and is used to iterate over the raw logs and unpacked data for NotifyReward events raised by the AerodromeVoterRegistry contract.
type AerodromeVoterRegistryNotifyRewardIterator struct {
	Event *AerodromeVoterRegistryNotifyReward // Event containing the contract specifics and raw log

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
func (it *AerodromeVoterRegistryNotifyRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AerodromeVoterRegistryNotifyReward)
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
		it.Event = new(AerodromeVoterRegistryNotifyReward)
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
func (it *AerodromeVoterRegistryNotifyRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AerodromeVoterRegistryNotifyRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AerodromeVoterRegistryNotifyReward represents a NotifyReward event raised by the AerodromeVoterRegistry contract.
type AerodromeVoterRegistryNotifyReward struct {
	Sender common.Address
	Reward common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNotifyReward is a free log retrieval operation binding the contract event 0xf70d5c697de7ea828df48e5c4573cb2194c659f1901f70110c52b066dcf50826.
//
// Solidity: event NotifyReward(address indexed sender, address indexed reward, uint256 amount)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryFilterer) FilterNotifyReward(opts *bind.FilterOpts, sender []common.Address, reward []common.Address) (*AerodromeVoterRegistryNotifyRewardIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}

	logs, sub, err := _AerodromeVoterRegistry.contract.FilterLogs(opts, "NotifyReward", senderRule, rewardRule)
	if err != nil {
		return nil, err
	}
	return &AerodromeVoterRegistryNotifyRewardIterator{contract: _AerodromeVoterRegistry.contract, event: "NotifyReward", logs: logs, sub: sub}, nil
}

// WatchNotifyReward is a free log subscription operation binding the contract event 0xf70d5c697de7ea828df48e5c4573cb2194c659f1901f70110c52b066dcf50826.
//
// Solidity: event NotifyReward(address indexed sender, address indexed reward, uint256 amount)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryFilterer) WatchNotifyReward(opts *bind.WatchOpts, sink chan<- *AerodromeVoterRegistryNotifyReward, sender []common.Address, reward []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}

	logs, sub, err := _AerodromeVoterRegistry.contract.WatchLogs(opts, "NotifyReward", senderRule, rewardRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AerodromeVoterRegistryNotifyReward)
				if err := _AerodromeVoterRegistry.contract.UnpackLog(event, "NotifyReward", log); err != nil {
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

// ParseNotifyReward is a log parse operation binding the contract event 0xf70d5c697de7ea828df48e5c4573cb2194c659f1901f70110c52b066dcf50826.
//
// Solidity: event NotifyReward(address indexed sender, address indexed reward, uint256 amount)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryFilterer) ParseNotifyReward(log types.Log) (*AerodromeVoterRegistryNotifyReward, error) {
	event := new(AerodromeVoterRegistryNotifyReward)
	if err := _AerodromeVoterRegistry.contract.UnpackLog(event, "NotifyReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AerodromeVoterRegistryVotedIterator is returned from FilterVoted and is used to iterate over the raw logs and unpacked data for Voted events raised by the AerodromeVoterRegistry contract.
type AerodromeVoterRegistryVotedIterator struct {
	Event *AerodromeVoterRegistryVoted // Event containing the contract specifics and raw log

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
func (it *AerodromeVoterRegistryVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AerodromeVoterRegistryVoted)
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
		it.Event = new(AerodromeVoterRegistryVoted)
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
func (it *AerodromeVoterRegistryVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AerodromeVoterRegistryVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AerodromeVoterRegistryVoted represents a Voted event raised by the AerodromeVoterRegistry contract.
type AerodromeVoterRegistryVoted struct {
	Voter       common.Address
	Pool        common.Address
	TokenId     *big.Int
	Weight      *big.Int
	TotalWeight *big.Int
	Timestamp   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVoted is a free log retrieval operation binding the contract event 0x452d440efc30dfa14a0ef803ccb55936af860ec6a6960ed27f129bef913f296a.
//
// Solidity: event Voted(address indexed voter, address indexed pool, uint256 indexed tokenId, uint256 weight, uint256 totalWeight, uint256 timestamp)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryFilterer) FilterVoted(opts *bind.FilterOpts, voter []common.Address, pool []common.Address, tokenId []*big.Int) (*AerodromeVoterRegistryVotedIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _AerodromeVoterRegistry.contract.FilterLogs(opts, "Voted", voterRule, poolRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &AerodromeVoterRegistryVotedIterator{contract: _AerodromeVoterRegistry.contract, event: "Voted", logs: logs, sub: sub}, nil
}

// WatchVoted is a free log subscription operation binding the contract event 0x452d440efc30dfa14a0ef803ccb55936af860ec6a6960ed27f129bef913f296a.
//
// Solidity: event Voted(address indexed voter, address indexed pool, uint256 indexed tokenId, uint256 weight, uint256 totalWeight, uint256 timestamp)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryFilterer) WatchVoted(opts *bind.WatchOpts, sink chan<- *AerodromeVoterRegistryVoted, voter []common.Address, pool []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _AerodromeVoterRegistry.contract.WatchLogs(opts, "Voted", voterRule, poolRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AerodromeVoterRegistryVoted)
				if err := _AerodromeVoterRegistry.contract.UnpackLog(event, "Voted", log); err != nil {
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

// ParseVoted is a log parse operation binding the contract event 0x452d440efc30dfa14a0ef803ccb55936af860ec6a6960ed27f129bef913f296a.
//
// Solidity: event Voted(address indexed voter, address indexed pool, uint256 indexed tokenId, uint256 weight, uint256 totalWeight, uint256 timestamp)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryFilterer) ParseVoted(log types.Log) (*AerodromeVoterRegistryVoted, error) {
	event := new(AerodromeVoterRegistryVoted)
	if err := _AerodromeVoterRegistry.contract.UnpackLog(event, "Voted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AerodromeVoterRegistryWhitelistNFTIterator is returned from FilterWhitelistNFT and is used to iterate over the raw logs and unpacked data for WhitelistNFT events raised by the AerodromeVoterRegistry contract.
type AerodromeVoterRegistryWhitelistNFTIterator struct {
	Event *AerodromeVoterRegistryWhitelistNFT // Event containing the contract specifics and raw log

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
func (it *AerodromeVoterRegistryWhitelistNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AerodromeVoterRegistryWhitelistNFT)
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
		it.Event = new(AerodromeVoterRegistryWhitelistNFT)
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
func (it *AerodromeVoterRegistryWhitelistNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AerodromeVoterRegistryWhitelistNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AerodromeVoterRegistryWhitelistNFT represents a WhitelistNFT event raised by the AerodromeVoterRegistry contract.
type AerodromeVoterRegistryWhitelistNFT struct {
	Whitelister common.Address
	TokenId     *big.Int
	Bool        bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWhitelistNFT is a free log retrieval operation binding the contract event 0x8a6ff732c8641e1e34d771e1f8b1673e988c1abdfb694ebdf6c910a5e3d0d853.
//
// Solidity: event WhitelistNFT(address indexed whitelister, uint256 indexed tokenId, bool indexed _bool)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryFilterer) FilterWhitelistNFT(opts *bind.FilterOpts, whitelister []common.Address, tokenId []*big.Int, _bool []bool) (*AerodromeVoterRegistryWhitelistNFTIterator, error) {

	var whitelisterRule []interface{}
	for _, whitelisterItem := range whitelister {
		whitelisterRule = append(whitelisterRule, whitelisterItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var _boolRule []interface{}
	for _, _boolItem := range _bool {
		_boolRule = append(_boolRule, _boolItem)
	}

	logs, sub, err := _AerodromeVoterRegistry.contract.FilterLogs(opts, "WhitelistNFT", whitelisterRule, tokenIdRule, _boolRule)
	if err != nil {
		return nil, err
	}
	return &AerodromeVoterRegistryWhitelistNFTIterator{contract: _AerodromeVoterRegistry.contract, event: "WhitelistNFT", logs: logs, sub: sub}, nil
}

// WatchWhitelistNFT is a free log subscription operation binding the contract event 0x8a6ff732c8641e1e34d771e1f8b1673e988c1abdfb694ebdf6c910a5e3d0d853.
//
// Solidity: event WhitelistNFT(address indexed whitelister, uint256 indexed tokenId, bool indexed _bool)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryFilterer) WatchWhitelistNFT(opts *bind.WatchOpts, sink chan<- *AerodromeVoterRegistryWhitelistNFT, whitelister []common.Address, tokenId []*big.Int, _bool []bool) (event.Subscription, error) {

	var whitelisterRule []interface{}
	for _, whitelisterItem := range whitelister {
		whitelisterRule = append(whitelisterRule, whitelisterItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var _boolRule []interface{}
	for _, _boolItem := range _bool {
		_boolRule = append(_boolRule, _boolItem)
	}

	logs, sub, err := _AerodromeVoterRegistry.contract.WatchLogs(opts, "WhitelistNFT", whitelisterRule, tokenIdRule, _boolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AerodromeVoterRegistryWhitelistNFT)
				if err := _AerodromeVoterRegistry.contract.UnpackLog(event, "WhitelistNFT", log); err != nil {
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

// ParseWhitelistNFT is a log parse operation binding the contract event 0x8a6ff732c8641e1e34d771e1f8b1673e988c1abdfb694ebdf6c910a5e3d0d853.
//
// Solidity: event WhitelistNFT(address indexed whitelister, uint256 indexed tokenId, bool indexed _bool)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryFilterer) ParseWhitelistNFT(log types.Log) (*AerodromeVoterRegistryWhitelistNFT, error) {
	event := new(AerodromeVoterRegistryWhitelistNFT)
	if err := _AerodromeVoterRegistry.contract.UnpackLog(event, "WhitelistNFT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AerodromeVoterRegistryWhitelistTokenIterator is returned from FilterWhitelistToken and is used to iterate over the raw logs and unpacked data for WhitelistToken events raised by the AerodromeVoterRegistry contract.
type AerodromeVoterRegistryWhitelistTokenIterator struct {
	Event *AerodromeVoterRegistryWhitelistToken // Event containing the contract specifics and raw log

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
func (it *AerodromeVoterRegistryWhitelistTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AerodromeVoterRegistryWhitelistToken)
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
		it.Event = new(AerodromeVoterRegistryWhitelistToken)
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
func (it *AerodromeVoterRegistryWhitelistTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AerodromeVoterRegistryWhitelistTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AerodromeVoterRegistryWhitelistToken represents a WhitelistToken event raised by the AerodromeVoterRegistry contract.
type AerodromeVoterRegistryWhitelistToken struct {
	Whitelister common.Address
	Token       common.Address
	Bool        bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWhitelistToken is a free log retrieval operation binding the contract event 0x44948130cf88523dbc150908a47dd6332c33a01a3869d7f2fa78e51d5a5f9c57.
//
// Solidity: event WhitelistToken(address indexed whitelister, address indexed token, bool indexed _bool)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryFilterer) FilterWhitelistToken(opts *bind.FilterOpts, whitelister []common.Address, token []common.Address, _bool []bool) (*AerodromeVoterRegistryWhitelistTokenIterator, error) {

	var whitelisterRule []interface{}
	for _, whitelisterItem := range whitelister {
		whitelisterRule = append(whitelisterRule, whitelisterItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var _boolRule []interface{}
	for _, _boolItem := range _bool {
		_boolRule = append(_boolRule, _boolItem)
	}

	logs, sub, err := _AerodromeVoterRegistry.contract.FilterLogs(opts, "WhitelistToken", whitelisterRule, tokenRule, _boolRule)
	if err != nil {
		return nil, err
	}
	return &AerodromeVoterRegistryWhitelistTokenIterator{contract: _AerodromeVoterRegistry.contract, event: "WhitelistToken", logs: logs, sub: sub}, nil
}

// WatchWhitelistToken is a free log subscription operation binding the contract event 0x44948130cf88523dbc150908a47dd6332c33a01a3869d7f2fa78e51d5a5f9c57.
//
// Solidity: event WhitelistToken(address indexed whitelister, address indexed token, bool indexed _bool)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryFilterer) WatchWhitelistToken(opts *bind.WatchOpts, sink chan<- *AerodromeVoterRegistryWhitelistToken, whitelister []common.Address, token []common.Address, _bool []bool) (event.Subscription, error) {

	var whitelisterRule []interface{}
	for _, whitelisterItem := range whitelister {
		whitelisterRule = append(whitelisterRule, whitelisterItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var _boolRule []interface{}
	for _, _boolItem := range _bool {
		_boolRule = append(_boolRule, _boolItem)
	}

	logs, sub, err := _AerodromeVoterRegistry.contract.WatchLogs(opts, "WhitelistToken", whitelisterRule, tokenRule, _boolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AerodromeVoterRegistryWhitelistToken)
				if err := _AerodromeVoterRegistry.contract.UnpackLog(event, "WhitelistToken", log); err != nil {
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

// ParseWhitelistToken is a log parse operation binding the contract event 0x44948130cf88523dbc150908a47dd6332c33a01a3869d7f2fa78e51d5a5f9c57.
//
// Solidity: event WhitelistToken(address indexed whitelister, address indexed token, bool indexed _bool)
func (_AerodromeVoterRegistry *AerodromeVoterRegistryFilterer) ParseWhitelistToken(log types.Log) (*AerodromeVoterRegistryWhitelistToken, error) {
	event := new(AerodromeVoterRegistryWhitelistToken)
	if err := _AerodromeVoterRegistry.contract.UnpackLog(event, "WhitelistToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
