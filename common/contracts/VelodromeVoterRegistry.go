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

// YVelodromeVoterRegistryMetaData contains all meta data concerning the YVelodromeVoterRegistry contract.
var YVelodromeVoterRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_forwarder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_factoryRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_v1Factory\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyVotedOrDeposited\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DistributeWindow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FactoryPathNotApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GaugeAlreadyKilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GaugeAlreadyRevived\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"GaugeDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GaugeExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gauge\",\"type\":\"address\"}],\"name\":\"GaugeNotAlive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InactiveManagedNFT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaximumVotingNumberTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonZeroVotes\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAPool\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotApprovedOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEmergencyCouncil\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotGovernor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotMinter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotWhitelistedNFT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotWhitelistedToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SameValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SpecialVotingWindow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnequalLengths\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroBalance\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalWeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Abstained\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DistributeReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"poolFactory\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"votingRewardsFactory\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gaugeFactory\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bribeVotingReward\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeVotingReward\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"GaugeCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"}],\"name\":\"GaugeKilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"}],\"name\":\"GaugeRevived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NotifyReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalWeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Voted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"whitelister\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"_bool\",\"type\":\"bool\"}],\"name\":\"WhitelistNFT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"whitelister\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"_bool\",\"type\":\"bool\"}],\"name\":\"WhitelistToken\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_bribes\",\"type\":\"address[]\"},{\"internalType\":\"address[][]\",\"name\":\"_tokens\",\"type\":\"address[][]\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"claimBribes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_fees\",\"type\":\"address[]\"},{\"internalType\":\"address[][]\",\"name\":\"_tokens\",\"type\":\"address[][]\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"claimFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_gauges\",\"type\":\"address[]\"}],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"createGauge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_mTokenId\",\"type\":\"uint256\"}],\"name\":\"depositManaged\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_gauges\",\"type\":\"address[]\"}],\"name\":\"distribute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_finish\",\"type\":\"uint256\"}],\"name\":\"distribute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyCouncil\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochGovernor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"epochNext\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"epochStart\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"epochVoteEnd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"epochVoteStart\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factoryRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forwarder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"gaugeToBribe\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"gaugeToFees\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"gauges\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isAlive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isGauge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"isWhitelistedNFT\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isWhitelistedToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gauge\",\"type\":\"address\"}],\"name\":\"killGauge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lastVoted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"length\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxVotingNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"notifyRewardAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"poke\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"poolForGauge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolVote\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pools\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"reset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gauge\",\"type\":\"address\"}],\"name\":\"reviveGauge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_council\",\"type\":\"address\"}],\"name\":\"setEmergencyCouncil\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_epochGovernor\",\"type\":\"address\"}],\"name\":\"setEpochGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governor\",\"type\":\"address\"}],\"name\":\"setGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxVotingNum\",\"type\":\"uint256\"}],\"name\":\"setMaxVotingNum\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gauge\",\"type\":\"address\"}],\"name\":\"updateFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"updateFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_gauges\",\"type\":\"address[]\"}],\"name\":\"updateFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"usedWeights\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"v1Factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ve\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_poolVote\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_weights\",\"type\":\"uint256[]\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"votes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"weights\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_bool\",\"type\":\"bool\"}],\"name\":\"whitelistNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_bool\",\"type\":\"bool\"}],\"name\":\"whitelistToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"withdrawManaged\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// YVelodromeVoterRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use YVelodromeVoterRegistryMetaData.ABI instead.
var YVelodromeVoterRegistryABI = YVelodromeVoterRegistryMetaData.ABI

// YVelodromeVoterRegistry is an auto generated Go binding around an Ethereum contract.
type YVelodromeVoterRegistry struct {
	YVelodromeVoterRegistryCaller     // Read-only binding to the contract
	YVelodromeVoterRegistryTransactor // Write-only binding to the contract
	YVelodromeVoterRegistryFilterer   // Log filterer for contract events
}

// YVelodromeVoterRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type YVelodromeVoterRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YVelodromeVoterRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YVelodromeVoterRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YVelodromeVoterRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YVelodromeVoterRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YVelodromeVoterRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YVelodromeVoterRegistrySession struct {
	Contract     *YVelodromeVoterRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// YVelodromeVoterRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YVelodromeVoterRegistryCallerSession struct {
	Contract *YVelodromeVoterRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// YVelodromeVoterRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YVelodromeVoterRegistryTransactorSession struct {
	Contract     *YVelodromeVoterRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// YVelodromeVoterRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type YVelodromeVoterRegistryRaw struct {
	Contract *YVelodromeVoterRegistry // Generic contract binding to access the raw methods on
}

// YVelodromeVoterRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YVelodromeVoterRegistryCallerRaw struct {
	Contract *YVelodromeVoterRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// YVelodromeVoterRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YVelodromeVoterRegistryTransactorRaw struct {
	Contract *YVelodromeVoterRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYVelodromeVoterRegistry creates a new instance of YVelodromeVoterRegistry, bound to a specific deployed contract.
func NewYVelodromeVoterRegistry(address common.Address, backend bind.ContractBackend) (*YVelodromeVoterRegistry, error) {
	contract, err := bindYVelodromeVoterRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YVelodromeVoterRegistry{YVelodromeVoterRegistryCaller: YVelodromeVoterRegistryCaller{contract: contract}, YVelodromeVoterRegistryTransactor: YVelodromeVoterRegistryTransactor{contract: contract}, YVelodromeVoterRegistryFilterer: YVelodromeVoterRegistryFilterer{contract: contract}}, nil
}

// NewYVelodromeVoterRegistryCaller creates a new read-only instance of YVelodromeVoterRegistry, bound to a specific deployed contract.
func NewYVelodromeVoterRegistryCaller(address common.Address, caller bind.ContractCaller) (*YVelodromeVoterRegistryCaller, error) {
	contract, err := bindYVelodromeVoterRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YVelodromeVoterRegistryCaller{contract: contract}, nil
}

// NewYVelodromeVoterRegistryTransactor creates a new write-only instance of YVelodromeVoterRegistry, bound to a specific deployed contract.
func NewYVelodromeVoterRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*YVelodromeVoterRegistryTransactor, error) {
	contract, err := bindYVelodromeVoterRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YVelodromeVoterRegistryTransactor{contract: contract}, nil
}

// NewYVelodromeVoterRegistryFilterer creates a new log filterer instance of YVelodromeVoterRegistry, bound to a specific deployed contract.
func NewYVelodromeVoterRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*YVelodromeVoterRegistryFilterer, error) {
	contract, err := bindYVelodromeVoterRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YVelodromeVoterRegistryFilterer{contract: contract}, nil
}

// bindYVelodromeVoterRegistry binds a generic wrapper to an already deployed contract.
func bindYVelodromeVoterRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(YVelodromeVoterRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YVelodromeVoterRegistry.Contract.YVelodromeVoterRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.YVelodromeVoterRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.YVelodromeVoterRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YVelodromeVoterRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.contract.Transact(opts, method, params...)
}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address ) view returns(uint256)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCaller) Claimable(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YVelodromeVoterRegistry.contract.Call(opts, &out, "claimable", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address ) view returns(uint256)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) Claimable(arg0 common.Address) (*big.Int, error) {
	return _YVelodromeVoterRegistry.Contract.Claimable(&_YVelodromeVoterRegistry.CallOpts, arg0)
}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address ) view returns(uint256)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCallerSession) Claimable(arg0 common.Address) (*big.Int, error) {
	return _YVelodromeVoterRegistry.Contract.Claimable(&_YVelodromeVoterRegistry.CallOpts, arg0)
}

// EmergencyCouncil is a free data retrieval call binding the contract method 0x7778960e.
//
// Solidity: function emergencyCouncil() view returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCaller) EmergencyCouncil(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YVelodromeVoterRegistry.contract.Call(opts, &out, "emergencyCouncil")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EmergencyCouncil is a free data retrieval call binding the contract method 0x7778960e.
//
// Solidity: function emergencyCouncil() view returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) EmergencyCouncil() (common.Address, error) {
	return _YVelodromeVoterRegistry.Contract.EmergencyCouncil(&_YVelodromeVoterRegistry.CallOpts)
}

// EmergencyCouncil is a free data retrieval call binding the contract method 0x7778960e.
//
// Solidity: function emergencyCouncil() view returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCallerSession) EmergencyCouncil() (common.Address, error) {
	return _YVelodromeVoterRegistry.Contract.EmergencyCouncil(&_YVelodromeVoterRegistry.CallOpts)
}

// EpochGovernor is a free data retrieval call binding the contract method 0x3aae971f.
//
// Solidity: function epochGovernor() view returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCaller) EpochGovernor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YVelodromeVoterRegistry.contract.Call(opts, &out, "epochGovernor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EpochGovernor is a free data retrieval call binding the contract method 0x3aae971f.
//
// Solidity: function epochGovernor() view returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) EpochGovernor() (common.Address, error) {
	return _YVelodromeVoterRegistry.Contract.EpochGovernor(&_YVelodromeVoterRegistry.CallOpts)
}

// EpochGovernor is a free data retrieval call binding the contract method 0x3aae971f.
//
// Solidity: function epochGovernor() view returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCallerSession) EpochGovernor() (common.Address, error) {
	return _YVelodromeVoterRegistry.Contract.EpochGovernor(&_YVelodromeVoterRegistry.CallOpts)
}

// EpochNext is a free data retrieval call binding the contract method 0x880e36fc.
//
// Solidity: function epochNext(uint256 _timestamp) pure returns(uint256)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCaller) EpochNext(opts *bind.CallOpts, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YVelodromeVoterRegistry.contract.Call(opts, &out, "epochNext", _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochNext is a free data retrieval call binding the contract method 0x880e36fc.
//
// Solidity: function epochNext(uint256 _timestamp) pure returns(uint256)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) EpochNext(_timestamp *big.Int) (*big.Int, error) {
	return _YVelodromeVoterRegistry.Contract.EpochNext(&_YVelodromeVoterRegistry.CallOpts, _timestamp)
}

// EpochNext is a free data retrieval call binding the contract method 0x880e36fc.
//
// Solidity: function epochNext(uint256 _timestamp) pure returns(uint256)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCallerSession) EpochNext(_timestamp *big.Int) (*big.Int, error) {
	return _YVelodromeVoterRegistry.Contract.EpochNext(&_YVelodromeVoterRegistry.CallOpts, _timestamp)
}

// EpochStart is a free data retrieval call binding the contract method 0xaa9354a3.
//
// Solidity: function epochStart(uint256 _timestamp) pure returns(uint256)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCaller) EpochStart(opts *bind.CallOpts, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YVelodromeVoterRegistry.contract.Call(opts, &out, "epochStart", _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochStart is a free data retrieval call binding the contract method 0xaa9354a3.
//
// Solidity: function epochStart(uint256 _timestamp) pure returns(uint256)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) EpochStart(_timestamp *big.Int) (*big.Int, error) {
	return _YVelodromeVoterRegistry.Contract.EpochStart(&_YVelodromeVoterRegistry.CallOpts, _timestamp)
}

// EpochStart is a free data retrieval call binding the contract method 0xaa9354a3.
//
// Solidity: function epochStart(uint256 _timestamp) pure returns(uint256)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCallerSession) EpochStart(_timestamp *big.Int) (*big.Int, error) {
	return _YVelodromeVoterRegistry.Contract.EpochStart(&_YVelodromeVoterRegistry.CallOpts, _timestamp)
}

// EpochVoteEnd is a free data retrieval call binding the contract method 0xd58b15d4.
//
// Solidity: function epochVoteEnd(uint256 _timestamp) pure returns(uint256)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCaller) EpochVoteEnd(opts *bind.CallOpts, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YVelodromeVoterRegistry.contract.Call(opts, &out, "epochVoteEnd", _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochVoteEnd is a free data retrieval call binding the contract method 0xd58b15d4.
//
// Solidity: function epochVoteEnd(uint256 _timestamp) pure returns(uint256)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) EpochVoteEnd(_timestamp *big.Int) (*big.Int, error) {
	return _YVelodromeVoterRegistry.Contract.EpochVoteEnd(&_YVelodromeVoterRegistry.CallOpts, _timestamp)
}

// EpochVoteEnd is a free data retrieval call binding the contract method 0xd58b15d4.
//
// Solidity: function epochVoteEnd(uint256 _timestamp) pure returns(uint256)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCallerSession) EpochVoteEnd(_timestamp *big.Int) (*big.Int, error) {
	return _YVelodromeVoterRegistry.Contract.EpochVoteEnd(&_YVelodromeVoterRegistry.CallOpts, _timestamp)
}

// EpochVoteStart is a free data retrieval call binding the contract method 0x39e9f3b6.
//
// Solidity: function epochVoteStart(uint256 _timestamp) pure returns(uint256)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCaller) EpochVoteStart(opts *bind.CallOpts, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YVelodromeVoterRegistry.contract.Call(opts, &out, "epochVoteStart", _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochVoteStart is a free data retrieval call binding the contract method 0x39e9f3b6.
//
// Solidity: function epochVoteStart(uint256 _timestamp) pure returns(uint256)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) EpochVoteStart(_timestamp *big.Int) (*big.Int, error) {
	return _YVelodromeVoterRegistry.Contract.EpochVoteStart(&_YVelodromeVoterRegistry.CallOpts, _timestamp)
}

// EpochVoteStart is a free data retrieval call binding the contract method 0x39e9f3b6.
//
// Solidity: function epochVoteStart(uint256 _timestamp) pure returns(uint256)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCallerSession) EpochVoteStart(_timestamp *big.Int) (*big.Int, error) {
	return _YVelodromeVoterRegistry.Contract.EpochVoteStart(&_YVelodromeVoterRegistry.CallOpts, _timestamp)
}

// FactoryRegistry is a free data retrieval call binding the contract method 0x3bf0c9fb.
//
// Solidity: function factoryRegistry() view returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCaller) FactoryRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YVelodromeVoterRegistry.contract.Call(opts, &out, "factoryRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FactoryRegistry is a free data retrieval call binding the contract method 0x3bf0c9fb.
//
// Solidity: function factoryRegistry() view returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) FactoryRegistry() (common.Address, error) {
	return _YVelodromeVoterRegistry.Contract.FactoryRegistry(&_YVelodromeVoterRegistry.CallOpts)
}

// FactoryRegistry is a free data retrieval call binding the contract method 0x3bf0c9fb.
//
// Solidity: function factoryRegistry() view returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCallerSession) FactoryRegistry() (common.Address, error) {
	return _YVelodromeVoterRegistry.Contract.FactoryRegistry(&_YVelodromeVoterRegistry.CallOpts)
}

// Forwarder is a free data retrieval call binding the contract method 0xf645d4f9.
//
// Solidity: function forwarder() view returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCaller) Forwarder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YVelodromeVoterRegistry.contract.Call(opts, &out, "forwarder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Forwarder is a free data retrieval call binding the contract method 0xf645d4f9.
//
// Solidity: function forwarder() view returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) Forwarder() (common.Address, error) {
	return _YVelodromeVoterRegistry.Contract.Forwarder(&_YVelodromeVoterRegistry.CallOpts)
}

// Forwarder is a free data retrieval call binding the contract method 0xf645d4f9.
//
// Solidity: function forwarder() view returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCallerSession) Forwarder() (common.Address, error) {
	return _YVelodromeVoterRegistry.Contract.Forwarder(&_YVelodromeVoterRegistry.CallOpts)
}

// GaugeToBribe is a free data retrieval call binding the contract method 0x929c8dcd.
//
// Solidity: function gaugeToBribe(address ) view returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCaller) GaugeToBribe(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _YVelodromeVoterRegistry.contract.Call(opts, &out, "gaugeToBribe", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GaugeToBribe is a free data retrieval call binding the contract method 0x929c8dcd.
//
// Solidity: function gaugeToBribe(address ) view returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) GaugeToBribe(arg0 common.Address) (common.Address, error) {
	return _YVelodromeVoterRegistry.Contract.GaugeToBribe(&_YVelodromeVoterRegistry.CallOpts, arg0)
}

// GaugeToBribe is a free data retrieval call binding the contract method 0x929c8dcd.
//
// Solidity: function gaugeToBribe(address ) view returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCallerSession) GaugeToBribe(arg0 common.Address) (common.Address, error) {
	return _YVelodromeVoterRegistry.Contract.GaugeToBribe(&_YVelodromeVoterRegistry.CallOpts, arg0)
}

// GaugeToFees is a free data retrieval call binding the contract method 0xc4f08165.
//
// Solidity: function gaugeToFees(address ) view returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCaller) GaugeToFees(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _YVelodromeVoterRegistry.contract.Call(opts, &out, "gaugeToFees", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GaugeToFees is a free data retrieval call binding the contract method 0xc4f08165.
//
// Solidity: function gaugeToFees(address ) view returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) GaugeToFees(arg0 common.Address) (common.Address, error) {
	return _YVelodromeVoterRegistry.Contract.GaugeToFees(&_YVelodromeVoterRegistry.CallOpts, arg0)
}

// GaugeToFees is a free data retrieval call binding the contract method 0xc4f08165.
//
// Solidity: function gaugeToFees(address ) view returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCallerSession) GaugeToFees(arg0 common.Address) (common.Address, error) {
	return _YVelodromeVoterRegistry.Contract.GaugeToFees(&_YVelodromeVoterRegistry.CallOpts, arg0)
}

// Gauges is a free data retrieval call binding the contract method 0xb9a09fd5.
//
// Solidity: function gauges(address ) view returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCaller) Gauges(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _YVelodromeVoterRegistry.contract.Call(opts, &out, "gauges", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gauges is a free data retrieval call binding the contract method 0xb9a09fd5.
//
// Solidity: function gauges(address ) view returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) Gauges(arg0 common.Address) (common.Address, error) {
	return _YVelodromeVoterRegistry.Contract.Gauges(&_YVelodromeVoterRegistry.CallOpts, arg0)
}

// Gauges is a free data retrieval call binding the contract method 0xb9a09fd5.
//
// Solidity: function gauges(address ) view returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCallerSession) Gauges(arg0 common.Address) (common.Address, error) {
	return _YVelodromeVoterRegistry.Contract.Gauges(&_YVelodromeVoterRegistry.CallOpts, arg0)
}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCaller) Governor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YVelodromeVoterRegistry.contract.Call(opts, &out, "governor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) Governor() (common.Address, error) {
	return _YVelodromeVoterRegistry.Contract.Governor(&_YVelodromeVoterRegistry.CallOpts)
}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCallerSession) Governor() (common.Address, error) {
	return _YVelodromeVoterRegistry.Contract.Governor(&_YVelodromeVoterRegistry.CallOpts)
}

// IsAlive is a free data retrieval call binding the contract method 0x1703e5f9.
//
// Solidity: function isAlive(address ) view returns(bool)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCaller) IsAlive(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _YVelodromeVoterRegistry.contract.Call(opts, &out, "isAlive", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAlive is a free data retrieval call binding the contract method 0x1703e5f9.
//
// Solidity: function isAlive(address ) view returns(bool)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) IsAlive(arg0 common.Address) (bool, error) {
	return _YVelodromeVoterRegistry.Contract.IsAlive(&_YVelodromeVoterRegistry.CallOpts, arg0)
}

// IsAlive is a free data retrieval call binding the contract method 0x1703e5f9.
//
// Solidity: function isAlive(address ) view returns(bool)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCallerSession) IsAlive(arg0 common.Address) (bool, error) {
	return _YVelodromeVoterRegistry.Contract.IsAlive(&_YVelodromeVoterRegistry.CallOpts, arg0)
}

// IsGauge is a free data retrieval call binding the contract method 0xaa79979b.
//
// Solidity: function isGauge(address ) view returns(bool)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCaller) IsGauge(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _YVelodromeVoterRegistry.contract.Call(opts, &out, "isGauge", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGauge is a free data retrieval call binding the contract method 0xaa79979b.
//
// Solidity: function isGauge(address ) view returns(bool)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) IsGauge(arg0 common.Address) (bool, error) {
	return _YVelodromeVoterRegistry.Contract.IsGauge(&_YVelodromeVoterRegistry.CallOpts, arg0)
}

// IsGauge is a free data retrieval call binding the contract method 0xaa79979b.
//
// Solidity: function isGauge(address ) view returns(bool)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCallerSession) IsGauge(arg0 common.Address) (bool, error) {
	return _YVelodromeVoterRegistry.Contract.IsGauge(&_YVelodromeVoterRegistry.CallOpts, arg0)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _YVelodromeVoterRegistry.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _YVelodromeVoterRegistry.Contract.IsTrustedForwarder(&_YVelodromeVoterRegistry.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _YVelodromeVoterRegistry.Contract.IsTrustedForwarder(&_YVelodromeVoterRegistry.CallOpts, forwarder)
}

// IsWhitelistedNFT is a free data retrieval call binding the contract method 0xd4e2616f.
//
// Solidity: function isWhitelistedNFT(uint256 ) view returns(bool)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCaller) IsWhitelistedNFT(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _YVelodromeVoterRegistry.contract.Call(opts, &out, "isWhitelistedNFT", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelistedNFT is a free data retrieval call binding the contract method 0xd4e2616f.
//
// Solidity: function isWhitelistedNFT(uint256 ) view returns(bool)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) IsWhitelistedNFT(arg0 *big.Int) (bool, error) {
	return _YVelodromeVoterRegistry.Contract.IsWhitelistedNFT(&_YVelodromeVoterRegistry.CallOpts, arg0)
}

// IsWhitelistedNFT is a free data retrieval call binding the contract method 0xd4e2616f.
//
// Solidity: function isWhitelistedNFT(uint256 ) view returns(bool)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCallerSession) IsWhitelistedNFT(arg0 *big.Int) (bool, error) {
	return _YVelodromeVoterRegistry.Contract.IsWhitelistedNFT(&_YVelodromeVoterRegistry.CallOpts, arg0)
}

// IsWhitelistedToken is a free data retrieval call binding the contract method 0xab37f486.
//
// Solidity: function isWhitelistedToken(address ) view returns(bool)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCaller) IsWhitelistedToken(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _YVelodromeVoterRegistry.contract.Call(opts, &out, "isWhitelistedToken", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelistedToken is a free data retrieval call binding the contract method 0xab37f486.
//
// Solidity: function isWhitelistedToken(address ) view returns(bool)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) IsWhitelistedToken(arg0 common.Address) (bool, error) {
	return _YVelodromeVoterRegistry.Contract.IsWhitelistedToken(&_YVelodromeVoterRegistry.CallOpts, arg0)
}

// IsWhitelistedToken is a free data retrieval call binding the contract method 0xab37f486.
//
// Solidity: function isWhitelistedToken(address ) view returns(bool)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCallerSession) IsWhitelistedToken(arg0 common.Address) (bool, error) {
	return _YVelodromeVoterRegistry.Contract.IsWhitelistedToken(&_YVelodromeVoterRegistry.CallOpts, arg0)
}

// LastVoted is a free data retrieval call binding the contract method 0xf3594be0.
//
// Solidity: function lastVoted(uint256 ) view returns(uint256)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCaller) LastVoted(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YVelodromeVoterRegistry.contract.Call(opts, &out, "lastVoted", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastVoted is a free data retrieval call binding the contract method 0xf3594be0.
//
// Solidity: function lastVoted(uint256 ) view returns(uint256)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) LastVoted(arg0 *big.Int) (*big.Int, error) {
	return _YVelodromeVoterRegistry.Contract.LastVoted(&_YVelodromeVoterRegistry.CallOpts, arg0)
}

// LastVoted is a free data retrieval call binding the contract method 0xf3594be0.
//
// Solidity: function lastVoted(uint256 ) view returns(uint256)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCallerSession) LastVoted(arg0 *big.Int) (*big.Int, error) {
	return _YVelodromeVoterRegistry.Contract.LastVoted(&_YVelodromeVoterRegistry.CallOpts, arg0)
}

// Length is a free data retrieval call binding the contract method 0x1f7b6d32.
//
// Solidity: function length() view returns(uint256)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCaller) Length(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YVelodromeVoterRegistry.contract.Call(opts, &out, "length")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Length is a free data retrieval call binding the contract method 0x1f7b6d32.
//
// Solidity: function length() view returns(uint256)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) Length() (*big.Int, error) {
	return _YVelodromeVoterRegistry.Contract.Length(&_YVelodromeVoterRegistry.CallOpts)
}

// Length is a free data retrieval call binding the contract method 0x1f7b6d32.
//
// Solidity: function length() view returns(uint256)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCallerSession) Length() (*big.Int, error) {
	return _YVelodromeVoterRegistry.Contract.Length(&_YVelodromeVoterRegistry.CallOpts)
}

// MaxVotingNum is a free data retrieval call binding the contract method 0xe8b3fd57.
//
// Solidity: function maxVotingNum() view returns(uint256)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCaller) MaxVotingNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YVelodromeVoterRegistry.contract.Call(opts, &out, "maxVotingNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxVotingNum is a free data retrieval call binding the contract method 0xe8b3fd57.
//
// Solidity: function maxVotingNum() view returns(uint256)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) MaxVotingNum() (*big.Int, error) {
	return _YVelodromeVoterRegistry.Contract.MaxVotingNum(&_YVelodromeVoterRegistry.CallOpts)
}

// MaxVotingNum is a free data retrieval call binding the contract method 0xe8b3fd57.
//
// Solidity: function maxVotingNum() view returns(uint256)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCallerSession) MaxVotingNum() (*big.Int, error) {
	return _YVelodromeVoterRegistry.Contract.MaxVotingNum(&_YVelodromeVoterRegistry.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YVelodromeVoterRegistry.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) Minter() (common.Address, error) {
	return _YVelodromeVoterRegistry.Contract.Minter(&_YVelodromeVoterRegistry.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCallerSession) Minter() (common.Address, error) {
	return _YVelodromeVoterRegistry.Contract.Minter(&_YVelodromeVoterRegistry.CallOpts)
}

// PoolForGauge is a free data retrieval call binding the contract method 0x06d6a1b2.
//
// Solidity: function poolForGauge(address ) view returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCaller) PoolForGauge(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _YVelodromeVoterRegistry.contract.Call(opts, &out, "poolForGauge", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolForGauge is a free data retrieval call binding the contract method 0x06d6a1b2.
//
// Solidity: function poolForGauge(address ) view returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) PoolForGauge(arg0 common.Address) (common.Address, error) {
	return _YVelodromeVoterRegistry.Contract.PoolForGauge(&_YVelodromeVoterRegistry.CallOpts, arg0)
}

// PoolForGauge is a free data retrieval call binding the contract method 0x06d6a1b2.
//
// Solidity: function poolForGauge(address ) view returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCallerSession) PoolForGauge(arg0 common.Address) (common.Address, error) {
	return _YVelodromeVoterRegistry.Contract.PoolForGauge(&_YVelodromeVoterRegistry.CallOpts, arg0)
}

// PoolVote is a free data retrieval call binding the contract method 0xa86a366d.
//
// Solidity: function poolVote(uint256 , uint256 ) view returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCaller) PoolVote(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _YVelodromeVoterRegistry.contract.Call(opts, &out, "poolVote", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolVote is a free data retrieval call binding the contract method 0xa86a366d.
//
// Solidity: function poolVote(uint256 , uint256 ) view returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) PoolVote(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _YVelodromeVoterRegistry.Contract.PoolVote(&_YVelodromeVoterRegistry.CallOpts, arg0, arg1)
}

// PoolVote is a free data retrieval call binding the contract method 0xa86a366d.
//
// Solidity: function poolVote(uint256 , uint256 ) view returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCallerSession) PoolVote(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _YVelodromeVoterRegistry.Contract.PoolVote(&_YVelodromeVoterRegistry.CallOpts, arg0, arg1)
}

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCaller) Pools(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _YVelodromeVoterRegistry.contract.Call(opts, &out, "pools", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) Pools(arg0 *big.Int) (common.Address, error) {
	return _YVelodromeVoterRegistry.Contract.Pools(&_YVelodromeVoterRegistry.CallOpts, arg0)
}

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCallerSession) Pools(arg0 *big.Int) (common.Address, error) {
	return _YVelodromeVoterRegistry.Contract.Pools(&_YVelodromeVoterRegistry.CallOpts, arg0)
}

// TotalWeight is a free data retrieval call binding the contract method 0x96c82e57.
//
// Solidity: function totalWeight() view returns(uint256)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCaller) TotalWeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YVelodromeVoterRegistry.contract.Call(opts, &out, "totalWeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalWeight is a free data retrieval call binding the contract method 0x96c82e57.
//
// Solidity: function totalWeight() view returns(uint256)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) TotalWeight() (*big.Int, error) {
	return _YVelodromeVoterRegistry.Contract.TotalWeight(&_YVelodromeVoterRegistry.CallOpts)
}

// TotalWeight is a free data retrieval call binding the contract method 0x96c82e57.
//
// Solidity: function totalWeight() view returns(uint256)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCallerSession) TotalWeight() (*big.Int, error) {
	return _YVelodromeVoterRegistry.Contract.TotalWeight(&_YVelodromeVoterRegistry.CallOpts)
}

// UsedWeights is a free data retrieval call binding the contract method 0x79e93824.
//
// Solidity: function usedWeights(uint256 ) view returns(uint256)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCaller) UsedWeights(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YVelodromeVoterRegistry.contract.Call(opts, &out, "usedWeights", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsedWeights is a free data retrieval call binding the contract method 0x79e93824.
//
// Solidity: function usedWeights(uint256 ) view returns(uint256)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) UsedWeights(arg0 *big.Int) (*big.Int, error) {
	return _YVelodromeVoterRegistry.Contract.UsedWeights(&_YVelodromeVoterRegistry.CallOpts, arg0)
}

// UsedWeights is a free data retrieval call binding the contract method 0x79e93824.
//
// Solidity: function usedWeights(uint256 ) view returns(uint256)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCallerSession) UsedWeights(arg0 *big.Int) (*big.Int, error) {
	return _YVelodromeVoterRegistry.Contract.UsedWeights(&_YVelodromeVoterRegistry.CallOpts, arg0)
}

// V1Factory is a free data retrieval call binding the contract method 0x8083f7bb.
//
// Solidity: function v1Factory() view returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCaller) V1Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YVelodromeVoterRegistry.contract.Call(opts, &out, "v1Factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// V1Factory is a free data retrieval call binding the contract method 0x8083f7bb.
//
// Solidity: function v1Factory() view returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) V1Factory() (common.Address, error) {
	return _YVelodromeVoterRegistry.Contract.V1Factory(&_YVelodromeVoterRegistry.CallOpts)
}

// V1Factory is a free data retrieval call binding the contract method 0x8083f7bb.
//
// Solidity: function v1Factory() view returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCallerSession) V1Factory() (common.Address, error) {
	return _YVelodromeVoterRegistry.Contract.V1Factory(&_YVelodromeVoterRegistry.CallOpts)
}

// Ve is a free data retrieval call binding the contract method 0x1f850716.
//
// Solidity: function ve() view returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCaller) Ve(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YVelodromeVoterRegistry.contract.Call(opts, &out, "ve")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ve is a free data retrieval call binding the contract method 0x1f850716.
//
// Solidity: function ve() view returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) Ve() (common.Address, error) {
	return _YVelodromeVoterRegistry.Contract.Ve(&_YVelodromeVoterRegistry.CallOpts)
}

// Ve is a free data retrieval call binding the contract method 0x1f850716.
//
// Solidity: function ve() view returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCallerSession) Ve() (common.Address, error) {
	return _YVelodromeVoterRegistry.Contract.Ve(&_YVelodromeVoterRegistry.CallOpts)
}

// Votes is a free data retrieval call binding the contract method 0xd23254b4.
//
// Solidity: function votes(uint256 , address ) view returns(uint256)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCaller) Votes(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YVelodromeVoterRegistry.contract.Call(opts, &out, "votes", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Votes is a free data retrieval call binding the contract method 0xd23254b4.
//
// Solidity: function votes(uint256 , address ) view returns(uint256)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) Votes(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _YVelodromeVoterRegistry.Contract.Votes(&_YVelodromeVoterRegistry.CallOpts, arg0, arg1)
}

// Votes is a free data retrieval call binding the contract method 0xd23254b4.
//
// Solidity: function votes(uint256 , address ) view returns(uint256)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCallerSession) Votes(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _YVelodromeVoterRegistry.Contract.Votes(&_YVelodromeVoterRegistry.CallOpts, arg0, arg1)
}

// Weights is a free data retrieval call binding the contract method 0xa7cac846.
//
// Solidity: function weights(address ) view returns(uint256)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCaller) Weights(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YVelodromeVoterRegistry.contract.Call(opts, &out, "weights", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Weights is a free data retrieval call binding the contract method 0xa7cac846.
//
// Solidity: function weights(address ) view returns(uint256)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) Weights(arg0 common.Address) (*big.Int, error) {
	return _YVelodromeVoterRegistry.Contract.Weights(&_YVelodromeVoterRegistry.CallOpts, arg0)
}

// Weights is a free data retrieval call binding the contract method 0xa7cac846.
//
// Solidity: function weights(address ) view returns(uint256)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryCallerSession) Weights(arg0 common.Address) (*big.Int, error) {
	return _YVelodromeVoterRegistry.Contract.Weights(&_YVelodromeVoterRegistry.CallOpts, arg0)
}

// ClaimBribes is a paid mutator transaction binding the contract method 0x7715ee75.
//
// Solidity: function claimBribes(address[] _bribes, address[][] _tokens, uint256 _tokenId) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactor) ClaimBribes(opts *bind.TransactOpts, _bribes []common.Address, _tokens [][]common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.contract.Transact(opts, "claimBribes", _bribes, _tokens, _tokenId)
}

// ClaimBribes is a paid mutator transaction binding the contract method 0x7715ee75.
//
// Solidity: function claimBribes(address[] _bribes, address[][] _tokens, uint256 _tokenId) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) ClaimBribes(_bribes []common.Address, _tokens [][]common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.ClaimBribes(&_YVelodromeVoterRegistry.TransactOpts, _bribes, _tokens, _tokenId)
}

// ClaimBribes is a paid mutator transaction binding the contract method 0x7715ee75.
//
// Solidity: function claimBribes(address[] _bribes, address[][] _tokens, uint256 _tokenId) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactorSession) ClaimBribes(_bribes []common.Address, _tokens [][]common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.ClaimBribes(&_YVelodromeVoterRegistry.TransactOpts, _bribes, _tokens, _tokenId)
}

// ClaimFees is a paid mutator transaction binding the contract method 0x666256aa.
//
// Solidity: function claimFees(address[] _fees, address[][] _tokens, uint256 _tokenId) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactor) ClaimFees(opts *bind.TransactOpts, _fees []common.Address, _tokens [][]common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.contract.Transact(opts, "claimFees", _fees, _tokens, _tokenId)
}

// ClaimFees is a paid mutator transaction binding the contract method 0x666256aa.
//
// Solidity: function claimFees(address[] _fees, address[][] _tokens, uint256 _tokenId) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) ClaimFees(_fees []common.Address, _tokens [][]common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.ClaimFees(&_YVelodromeVoterRegistry.TransactOpts, _fees, _tokens, _tokenId)
}

// ClaimFees is a paid mutator transaction binding the contract method 0x666256aa.
//
// Solidity: function claimFees(address[] _fees, address[][] _tokens, uint256 _tokenId) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactorSession) ClaimFees(_fees []common.Address, _tokens [][]common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.ClaimFees(&_YVelodromeVoterRegistry.TransactOpts, _fees, _tokens, _tokenId)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xf9f031df.
//
// Solidity: function claimRewards(address[] _gauges) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactor) ClaimRewards(opts *bind.TransactOpts, _gauges []common.Address) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.contract.Transact(opts, "claimRewards", _gauges)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xf9f031df.
//
// Solidity: function claimRewards(address[] _gauges) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) ClaimRewards(_gauges []common.Address) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.ClaimRewards(&_YVelodromeVoterRegistry.TransactOpts, _gauges)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xf9f031df.
//
// Solidity: function claimRewards(address[] _gauges) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactorSession) ClaimRewards(_gauges []common.Address) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.ClaimRewards(&_YVelodromeVoterRegistry.TransactOpts, _gauges)
}

// CreateGauge is a paid mutator transaction binding the contract method 0x794cea3c.
//
// Solidity: function createGauge(address _poolFactory, address _pool) returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactor) CreateGauge(opts *bind.TransactOpts, _poolFactory common.Address, _pool common.Address) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.contract.Transact(opts, "createGauge", _poolFactory, _pool)
}

// CreateGauge is a paid mutator transaction binding the contract method 0x794cea3c.
//
// Solidity: function createGauge(address _poolFactory, address _pool) returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) CreateGauge(_poolFactory common.Address, _pool common.Address) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.CreateGauge(&_YVelodromeVoterRegistry.TransactOpts, _poolFactory, _pool)
}

// CreateGauge is a paid mutator transaction binding the contract method 0x794cea3c.
//
// Solidity: function createGauge(address _poolFactory, address _pool) returns(address)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactorSession) CreateGauge(_poolFactory common.Address, _pool common.Address) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.CreateGauge(&_YVelodromeVoterRegistry.TransactOpts, _poolFactory, _pool)
}

// DepositManaged is a paid mutator transaction binding the contract method 0xe0c11f9a.
//
// Solidity: function depositManaged(uint256 _tokenId, uint256 _mTokenId) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactor) DepositManaged(opts *bind.TransactOpts, _tokenId *big.Int, _mTokenId *big.Int) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.contract.Transact(opts, "depositManaged", _tokenId, _mTokenId)
}

// DepositManaged is a paid mutator transaction binding the contract method 0xe0c11f9a.
//
// Solidity: function depositManaged(uint256 _tokenId, uint256 _mTokenId) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) DepositManaged(_tokenId *big.Int, _mTokenId *big.Int) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.DepositManaged(&_YVelodromeVoterRegistry.TransactOpts, _tokenId, _mTokenId)
}

// DepositManaged is a paid mutator transaction binding the contract method 0xe0c11f9a.
//
// Solidity: function depositManaged(uint256 _tokenId, uint256 _mTokenId) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactorSession) DepositManaged(_tokenId *big.Int, _mTokenId *big.Int) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.DepositManaged(&_YVelodromeVoterRegistry.TransactOpts, _tokenId, _mTokenId)
}

// Distribute is a paid mutator transaction binding the contract method 0x6138889b.
//
// Solidity: function distribute(address[] _gauges) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactor) Distribute(opts *bind.TransactOpts, _gauges []common.Address) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.contract.Transact(opts, "distribute", _gauges)
}

// Distribute is a paid mutator transaction binding the contract method 0x6138889b.
//
// Solidity: function distribute(address[] _gauges) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) Distribute(_gauges []common.Address) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.Distribute(&_YVelodromeVoterRegistry.TransactOpts, _gauges)
}

// Distribute is a paid mutator transaction binding the contract method 0x6138889b.
//
// Solidity: function distribute(address[] _gauges) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactorSession) Distribute(_gauges []common.Address) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.Distribute(&_YVelodromeVoterRegistry.TransactOpts, _gauges)
}

// Distribute0 is a paid mutator transaction binding the contract method 0x7625391a.
//
// Solidity: function distribute(uint256 _start, uint256 _finish) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactor) Distribute0(opts *bind.TransactOpts, _start *big.Int, _finish *big.Int) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.contract.Transact(opts, "distribute0", _start, _finish)
}

// Distribute0 is a paid mutator transaction binding the contract method 0x7625391a.
//
// Solidity: function distribute(uint256 _start, uint256 _finish) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) Distribute0(_start *big.Int, _finish *big.Int) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.Distribute0(&_YVelodromeVoterRegistry.TransactOpts, _start, _finish)
}

// Distribute0 is a paid mutator transaction binding the contract method 0x7625391a.
//
// Solidity: function distribute(uint256 _start, uint256 _finish) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactorSession) Distribute0(_start *big.Int, _finish *big.Int) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.Distribute0(&_YVelodromeVoterRegistry.TransactOpts, _start, _finish)
}

// Initialize is a paid mutator transaction binding the contract method 0x462d0b2e.
//
// Solidity: function initialize(address[] _tokens, address _minter) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactor) Initialize(opts *bind.TransactOpts, _tokens []common.Address, _minter common.Address) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.contract.Transact(opts, "initialize", _tokens, _minter)
}

// Initialize is a paid mutator transaction binding the contract method 0x462d0b2e.
//
// Solidity: function initialize(address[] _tokens, address _minter) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) Initialize(_tokens []common.Address, _minter common.Address) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.Initialize(&_YVelodromeVoterRegistry.TransactOpts, _tokens, _minter)
}

// Initialize is a paid mutator transaction binding the contract method 0x462d0b2e.
//
// Solidity: function initialize(address[] _tokens, address _minter) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactorSession) Initialize(_tokens []common.Address, _minter common.Address) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.Initialize(&_YVelodromeVoterRegistry.TransactOpts, _tokens, _minter)
}

// KillGauge is a paid mutator transaction binding the contract method 0x992a7933.
//
// Solidity: function killGauge(address _gauge) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactor) KillGauge(opts *bind.TransactOpts, _gauge common.Address) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.contract.Transact(opts, "killGauge", _gauge)
}

// KillGauge is a paid mutator transaction binding the contract method 0x992a7933.
//
// Solidity: function killGauge(address _gauge) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) KillGauge(_gauge common.Address) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.KillGauge(&_YVelodromeVoterRegistry.TransactOpts, _gauge)
}

// KillGauge is a paid mutator transaction binding the contract method 0x992a7933.
//
// Solidity: function killGauge(address _gauge) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactorSession) KillGauge(_gauge common.Address) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.KillGauge(&_YVelodromeVoterRegistry.TransactOpts, _gauge)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 _amount) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactor) NotifyRewardAmount(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.contract.Transact(opts, "notifyRewardAmount", _amount)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 _amount) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) NotifyRewardAmount(_amount *big.Int) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.NotifyRewardAmount(&_YVelodromeVoterRegistry.TransactOpts, _amount)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 _amount) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactorSession) NotifyRewardAmount(_amount *big.Int) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.NotifyRewardAmount(&_YVelodromeVoterRegistry.TransactOpts, _amount)
}

// Poke is a paid mutator transaction binding the contract method 0x32145f90.
//
// Solidity: function poke(uint256 _tokenId) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactor) Poke(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.contract.Transact(opts, "poke", _tokenId)
}

// Poke is a paid mutator transaction binding the contract method 0x32145f90.
//
// Solidity: function poke(uint256 _tokenId) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) Poke(_tokenId *big.Int) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.Poke(&_YVelodromeVoterRegistry.TransactOpts, _tokenId)
}

// Poke is a paid mutator transaction binding the contract method 0x32145f90.
//
// Solidity: function poke(uint256 _tokenId) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactorSession) Poke(_tokenId *big.Int) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.Poke(&_YVelodromeVoterRegistry.TransactOpts, _tokenId)
}

// Reset is a paid mutator transaction binding the contract method 0x310bd74b.
//
// Solidity: function reset(uint256 _tokenId) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactor) Reset(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.contract.Transact(opts, "reset", _tokenId)
}

// Reset is a paid mutator transaction binding the contract method 0x310bd74b.
//
// Solidity: function reset(uint256 _tokenId) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) Reset(_tokenId *big.Int) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.Reset(&_YVelodromeVoterRegistry.TransactOpts, _tokenId)
}

// Reset is a paid mutator transaction binding the contract method 0x310bd74b.
//
// Solidity: function reset(uint256 _tokenId) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactorSession) Reset(_tokenId *big.Int) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.Reset(&_YVelodromeVoterRegistry.TransactOpts, _tokenId)
}

// ReviveGauge is a paid mutator transaction binding the contract method 0x9f06247b.
//
// Solidity: function reviveGauge(address _gauge) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactor) ReviveGauge(opts *bind.TransactOpts, _gauge common.Address) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.contract.Transact(opts, "reviveGauge", _gauge)
}

// ReviveGauge is a paid mutator transaction binding the contract method 0x9f06247b.
//
// Solidity: function reviveGauge(address _gauge) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) ReviveGauge(_gauge common.Address) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.ReviveGauge(&_YVelodromeVoterRegistry.TransactOpts, _gauge)
}

// ReviveGauge is a paid mutator transaction binding the contract method 0x9f06247b.
//
// Solidity: function reviveGauge(address _gauge) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactorSession) ReviveGauge(_gauge common.Address) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.ReviveGauge(&_YVelodromeVoterRegistry.TransactOpts, _gauge)
}

// SetEmergencyCouncil is a paid mutator transaction binding the contract method 0xe586875f.
//
// Solidity: function setEmergencyCouncil(address _council) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactor) SetEmergencyCouncil(opts *bind.TransactOpts, _council common.Address) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.contract.Transact(opts, "setEmergencyCouncil", _council)
}

// SetEmergencyCouncil is a paid mutator transaction binding the contract method 0xe586875f.
//
// Solidity: function setEmergencyCouncil(address _council) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) SetEmergencyCouncil(_council common.Address) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.SetEmergencyCouncil(&_YVelodromeVoterRegistry.TransactOpts, _council)
}

// SetEmergencyCouncil is a paid mutator transaction binding the contract method 0xe586875f.
//
// Solidity: function setEmergencyCouncil(address _council) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactorSession) SetEmergencyCouncil(_council common.Address) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.SetEmergencyCouncil(&_YVelodromeVoterRegistry.TransactOpts, _council)
}

// SetEpochGovernor is a paid mutator transaction binding the contract method 0x598d521b.
//
// Solidity: function setEpochGovernor(address _epochGovernor) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactor) SetEpochGovernor(opts *bind.TransactOpts, _epochGovernor common.Address) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.contract.Transact(opts, "setEpochGovernor", _epochGovernor)
}

// SetEpochGovernor is a paid mutator transaction binding the contract method 0x598d521b.
//
// Solidity: function setEpochGovernor(address _epochGovernor) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) SetEpochGovernor(_epochGovernor common.Address) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.SetEpochGovernor(&_YVelodromeVoterRegistry.TransactOpts, _epochGovernor)
}

// SetEpochGovernor is a paid mutator transaction binding the contract method 0x598d521b.
//
// Solidity: function setEpochGovernor(address _epochGovernor) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactorSession) SetEpochGovernor(_epochGovernor common.Address) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.SetEpochGovernor(&_YVelodromeVoterRegistry.TransactOpts, _epochGovernor)
}

// SetGovernor is a paid mutator transaction binding the contract method 0xc42cf535.
//
// Solidity: function setGovernor(address _governor) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactor) SetGovernor(opts *bind.TransactOpts, _governor common.Address) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.contract.Transact(opts, "setGovernor", _governor)
}

// SetGovernor is a paid mutator transaction binding the contract method 0xc42cf535.
//
// Solidity: function setGovernor(address _governor) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) SetGovernor(_governor common.Address) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.SetGovernor(&_YVelodromeVoterRegistry.TransactOpts, _governor)
}

// SetGovernor is a paid mutator transaction binding the contract method 0xc42cf535.
//
// Solidity: function setGovernor(address _governor) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactorSession) SetGovernor(_governor common.Address) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.SetGovernor(&_YVelodromeVoterRegistry.TransactOpts, _governor)
}

// SetMaxVotingNum is a paid mutator transaction binding the contract method 0x30331b2f.
//
// Solidity: function setMaxVotingNum(uint256 _maxVotingNum) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactor) SetMaxVotingNum(opts *bind.TransactOpts, _maxVotingNum *big.Int) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.contract.Transact(opts, "setMaxVotingNum", _maxVotingNum)
}

// SetMaxVotingNum is a paid mutator transaction binding the contract method 0x30331b2f.
//
// Solidity: function setMaxVotingNum(uint256 _maxVotingNum) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) SetMaxVotingNum(_maxVotingNum *big.Int) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.SetMaxVotingNum(&_YVelodromeVoterRegistry.TransactOpts, _maxVotingNum)
}

// SetMaxVotingNum is a paid mutator transaction binding the contract method 0x30331b2f.
//
// Solidity: function setMaxVotingNum(uint256 _maxVotingNum) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactorSession) SetMaxVotingNum(_maxVotingNum *big.Int) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.SetMaxVotingNum(&_YVelodromeVoterRegistry.TransactOpts, _maxVotingNum)
}

// UpdateFor is a paid mutator transaction binding the contract method 0x0e0a5968.
//
// Solidity: function updateFor(address _gauge) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactor) UpdateFor(opts *bind.TransactOpts, _gauge common.Address) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.contract.Transact(opts, "updateFor", _gauge)
}

// UpdateFor is a paid mutator transaction binding the contract method 0x0e0a5968.
//
// Solidity: function updateFor(address _gauge) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) UpdateFor(_gauge common.Address) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.UpdateFor(&_YVelodromeVoterRegistry.TransactOpts, _gauge)
}

// UpdateFor is a paid mutator transaction binding the contract method 0x0e0a5968.
//
// Solidity: function updateFor(address _gauge) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactorSession) UpdateFor(_gauge common.Address) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.UpdateFor(&_YVelodromeVoterRegistry.TransactOpts, _gauge)
}

// UpdateFor0 is a paid mutator transaction binding the contract method 0xc9ff6f4d.
//
// Solidity: function updateFor(uint256 start, uint256 end) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactor) UpdateFor0(opts *bind.TransactOpts, start *big.Int, end *big.Int) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.contract.Transact(opts, "updateFor0", start, end)
}

// UpdateFor0 is a paid mutator transaction binding the contract method 0xc9ff6f4d.
//
// Solidity: function updateFor(uint256 start, uint256 end) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) UpdateFor0(start *big.Int, end *big.Int) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.UpdateFor0(&_YVelodromeVoterRegistry.TransactOpts, start, end)
}

// UpdateFor0 is a paid mutator transaction binding the contract method 0xc9ff6f4d.
//
// Solidity: function updateFor(uint256 start, uint256 end) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactorSession) UpdateFor0(start *big.Int, end *big.Int) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.UpdateFor0(&_YVelodromeVoterRegistry.TransactOpts, start, end)
}

// UpdateFor1 is a paid mutator transaction binding the contract method 0xd560b0d7.
//
// Solidity: function updateFor(address[] _gauges) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactor) UpdateFor1(opts *bind.TransactOpts, _gauges []common.Address) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.contract.Transact(opts, "updateFor1", _gauges)
}

// UpdateFor1 is a paid mutator transaction binding the contract method 0xd560b0d7.
//
// Solidity: function updateFor(address[] _gauges) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) UpdateFor1(_gauges []common.Address) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.UpdateFor1(&_YVelodromeVoterRegistry.TransactOpts, _gauges)
}

// UpdateFor1 is a paid mutator transaction binding the contract method 0xd560b0d7.
//
// Solidity: function updateFor(address[] _gauges) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactorSession) UpdateFor1(_gauges []common.Address) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.UpdateFor1(&_YVelodromeVoterRegistry.TransactOpts, _gauges)
}

// Vote is a paid mutator transaction binding the contract method 0x7ac09bf7.
//
// Solidity: function vote(uint256 _tokenId, address[] _poolVote, uint256[] _weights) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactor) Vote(opts *bind.TransactOpts, _tokenId *big.Int, _poolVote []common.Address, _weights []*big.Int) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.contract.Transact(opts, "vote", _tokenId, _poolVote, _weights)
}

// Vote is a paid mutator transaction binding the contract method 0x7ac09bf7.
//
// Solidity: function vote(uint256 _tokenId, address[] _poolVote, uint256[] _weights) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) Vote(_tokenId *big.Int, _poolVote []common.Address, _weights []*big.Int) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.Vote(&_YVelodromeVoterRegistry.TransactOpts, _tokenId, _poolVote, _weights)
}

// Vote is a paid mutator transaction binding the contract method 0x7ac09bf7.
//
// Solidity: function vote(uint256 _tokenId, address[] _poolVote, uint256[] _weights) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactorSession) Vote(_tokenId *big.Int, _poolVote []common.Address, _weights []*big.Int) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.Vote(&_YVelodromeVoterRegistry.TransactOpts, _tokenId, _poolVote, _weights)
}

// WhitelistNFT is a paid mutator transaction binding the contract method 0xe2819d5c.
//
// Solidity: function whitelistNFT(uint256 _tokenId, bool _bool) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactor) WhitelistNFT(opts *bind.TransactOpts, _tokenId *big.Int, _bool bool) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.contract.Transact(opts, "whitelistNFT", _tokenId, _bool)
}

// WhitelistNFT is a paid mutator transaction binding the contract method 0xe2819d5c.
//
// Solidity: function whitelistNFT(uint256 _tokenId, bool _bool) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) WhitelistNFT(_tokenId *big.Int, _bool bool) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.WhitelistNFT(&_YVelodromeVoterRegistry.TransactOpts, _tokenId, _bool)
}

// WhitelistNFT is a paid mutator transaction binding the contract method 0xe2819d5c.
//
// Solidity: function whitelistNFT(uint256 _tokenId, bool _bool) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactorSession) WhitelistNFT(_tokenId *big.Int, _bool bool) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.WhitelistNFT(&_YVelodromeVoterRegistry.TransactOpts, _tokenId, _bool)
}

// WhitelistToken is a paid mutator transaction binding the contract method 0x0ffb1d8b.
//
// Solidity: function whitelistToken(address _token, bool _bool) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactor) WhitelistToken(opts *bind.TransactOpts, _token common.Address, _bool bool) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.contract.Transact(opts, "whitelistToken", _token, _bool)
}

// WhitelistToken is a paid mutator transaction binding the contract method 0x0ffb1d8b.
//
// Solidity: function whitelistToken(address _token, bool _bool) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) WhitelistToken(_token common.Address, _bool bool) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.WhitelistToken(&_YVelodromeVoterRegistry.TransactOpts, _token, _bool)
}

// WhitelistToken is a paid mutator transaction binding the contract method 0x0ffb1d8b.
//
// Solidity: function whitelistToken(address _token, bool _bool) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactorSession) WhitelistToken(_token common.Address, _bool bool) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.WhitelistToken(&_YVelodromeVoterRegistry.TransactOpts, _token, _bool)
}

// WithdrawManaged is a paid mutator transaction binding the contract method 0x370fb5fa.
//
// Solidity: function withdrawManaged(uint256 _tokenId) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactor) WithdrawManaged(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.contract.Transact(opts, "withdrawManaged", _tokenId)
}

// WithdrawManaged is a paid mutator transaction binding the contract method 0x370fb5fa.
//
// Solidity: function withdrawManaged(uint256 _tokenId) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistrySession) WithdrawManaged(_tokenId *big.Int) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.WithdrawManaged(&_YVelodromeVoterRegistry.TransactOpts, _tokenId)
}

// WithdrawManaged is a paid mutator transaction binding the contract method 0x370fb5fa.
//
// Solidity: function withdrawManaged(uint256 _tokenId) returns()
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryTransactorSession) WithdrawManaged(_tokenId *big.Int) (*types.Transaction, error) {
	return _YVelodromeVoterRegistry.Contract.WithdrawManaged(&_YVelodromeVoterRegistry.TransactOpts, _tokenId)
}

// YVelodromeVoterRegistryAbstainedIterator is returned from FilterAbstained and is used to iterate over the raw logs and unpacked data for Abstained events raised by the YVelodromeVoterRegistry contract.
type YVelodromeVoterRegistryAbstainedIterator struct {
	Event *YVelodromeVoterRegistryAbstained // Event containing the contract specifics and raw log

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
func (it *YVelodromeVoterRegistryAbstainedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YVelodromeVoterRegistryAbstained)
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
		it.Event = new(YVelodromeVoterRegistryAbstained)
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
func (it *YVelodromeVoterRegistryAbstainedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YVelodromeVoterRegistryAbstainedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YVelodromeVoterRegistryAbstained represents a Abstained event raised by the YVelodromeVoterRegistry contract.
type YVelodromeVoterRegistryAbstained struct {
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
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryFilterer) FilterAbstained(opts *bind.FilterOpts, voter []common.Address, pool []common.Address, tokenId []*big.Int) (*YVelodromeVoterRegistryAbstainedIterator, error) {

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

	logs, sub, err := _YVelodromeVoterRegistry.contract.FilterLogs(opts, "Abstained", voterRule, poolRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &YVelodromeVoterRegistryAbstainedIterator{contract: _YVelodromeVoterRegistry.contract, event: "Abstained", logs: logs, sub: sub}, nil
}

// WatchAbstained is a free log subscription operation binding the contract event 0xadab630928b1d46214641293704a312ee7ad87e03ae14a7fd95e7308b93998df.
//
// Solidity: event Abstained(address indexed voter, address indexed pool, uint256 indexed tokenId, uint256 weight, uint256 totalWeight, uint256 timestamp)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryFilterer) WatchAbstained(opts *bind.WatchOpts, sink chan<- *YVelodromeVoterRegistryAbstained, voter []common.Address, pool []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _YVelodromeVoterRegistry.contract.WatchLogs(opts, "Abstained", voterRule, poolRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YVelodromeVoterRegistryAbstained)
				if err := _YVelodromeVoterRegistry.contract.UnpackLog(event, "Abstained", log); err != nil {
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
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryFilterer) ParseAbstained(log types.Log) (*YVelodromeVoterRegistryAbstained, error) {
	event := new(YVelodromeVoterRegistryAbstained)
	if err := _YVelodromeVoterRegistry.contract.UnpackLog(event, "Abstained", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YVelodromeVoterRegistryDistributeRewardIterator is returned from FilterDistributeReward and is used to iterate over the raw logs and unpacked data for DistributeReward events raised by the YVelodromeVoterRegistry contract.
type YVelodromeVoterRegistryDistributeRewardIterator struct {
	Event *YVelodromeVoterRegistryDistributeReward // Event containing the contract specifics and raw log

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
func (it *YVelodromeVoterRegistryDistributeRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YVelodromeVoterRegistryDistributeReward)
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
		it.Event = new(YVelodromeVoterRegistryDistributeReward)
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
func (it *YVelodromeVoterRegistryDistributeRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YVelodromeVoterRegistryDistributeRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YVelodromeVoterRegistryDistributeReward represents a DistributeReward event raised by the YVelodromeVoterRegistry contract.
type YVelodromeVoterRegistryDistributeReward struct {
	Sender common.Address
	Gauge  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDistributeReward is a free log retrieval operation binding the contract event 0x4fa9693cae526341d334e2862ca2413b2e503f1266255f9e0869fb36e6d89b17.
//
// Solidity: event DistributeReward(address indexed sender, address indexed gauge, uint256 amount)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryFilterer) FilterDistributeReward(opts *bind.FilterOpts, sender []common.Address, gauge []common.Address) (*YVelodromeVoterRegistryDistributeRewardIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var gaugeRule []interface{}
	for _, gaugeItem := range gauge {
		gaugeRule = append(gaugeRule, gaugeItem)
	}

	logs, sub, err := _YVelodromeVoterRegistry.contract.FilterLogs(opts, "DistributeReward", senderRule, gaugeRule)
	if err != nil {
		return nil, err
	}
	return &YVelodromeVoterRegistryDistributeRewardIterator{contract: _YVelodromeVoterRegistry.contract, event: "DistributeReward", logs: logs, sub: sub}, nil
}

// WatchDistributeReward is a free log subscription operation binding the contract event 0x4fa9693cae526341d334e2862ca2413b2e503f1266255f9e0869fb36e6d89b17.
//
// Solidity: event DistributeReward(address indexed sender, address indexed gauge, uint256 amount)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryFilterer) WatchDistributeReward(opts *bind.WatchOpts, sink chan<- *YVelodromeVoterRegistryDistributeReward, sender []common.Address, gauge []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var gaugeRule []interface{}
	for _, gaugeItem := range gauge {
		gaugeRule = append(gaugeRule, gaugeItem)
	}

	logs, sub, err := _YVelodromeVoterRegistry.contract.WatchLogs(opts, "DistributeReward", senderRule, gaugeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YVelodromeVoterRegistryDistributeReward)
				if err := _YVelodromeVoterRegistry.contract.UnpackLog(event, "DistributeReward", log); err != nil {
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
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryFilterer) ParseDistributeReward(log types.Log) (*YVelodromeVoterRegistryDistributeReward, error) {
	event := new(YVelodromeVoterRegistryDistributeReward)
	if err := _YVelodromeVoterRegistry.contract.UnpackLog(event, "DistributeReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YVelodromeVoterRegistryGaugeCreatedIterator is returned from FilterGaugeCreated and is used to iterate over the raw logs and unpacked data for GaugeCreated events raised by the YVelodromeVoterRegistry contract.
type YVelodromeVoterRegistryGaugeCreatedIterator struct {
	Event *YVelodromeVoterRegistryGaugeCreated // Event containing the contract specifics and raw log

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
func (it *YVelodromeVoterRegistryGaugeCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YVelodromeVoterRegistryGaugeCreated)
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
		it.Event = new(YVelodromeVoterRegistryGaugeCreated)
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
func (it *YVelodromeVoterRegistryGaugeCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YVelodromeVoterRegistryGaugeCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YVelodromeVoterRegistryGaugeCreated represents a GaugeCreated event raised by the YVelodromeVoterRegistry contract.
type YVelodromeVoterRegistryGaugeCreated struct {
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
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryFilterer) FilterGaugeCreated(opts *bind.FilterOpts, poolFactory []common.Address, votingRewardsFactory []common.Address, gaugeFactory []common.Address) (*YVelodromeVoterRegistryGaugeCreatedIterator, error) {

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

	logs, sub, err := _YVelodromeVoterRegistry.contract.FilterLogs(opts, "GaugeCreated", poolFactoryRule, votingRewardsFactoryRule, gaugeFactoryRule)
	if err != nil {
		return nil, err
	}
	return &YVelodromeVoterRegistryGaugeCreatedIterator{contract: _YVelodromeVoterRegistry.contract, event: "GaugeCreated", logs: logs, sub: sub}, nil
}

// WatchGaugeCreated is a free log subscription operation binding the contract event 0xef9f7d1ffff3b249c6b9bf2528499e935f7d96bb6d6ec4e7da504d1d3c6279e1.
//
// Solidity: event GaugeCreated(address indexed poolFactory, address indexed votingRewardsFactory, address indexed gaugeFactory, address pool, address bribeVotingReward, address feeVotingReward, address gauge, address creator)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryFilterer) WatchGaugeCreated(opts *bind.WatchOpts, sink chan<- *YVelodromeVoterRegistryGaugeCreated, poolFactory []common.Address, votingRewardsFactory []common.Address, gaugeFactory []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _YVelodromeVoterRegistry.contract.WatchLogs(opts, "GaugeCreated", poolFactoryRule, votingRewardsFactoryRule, gaugeFactoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YVelodromeVoterRegistryGaugeCreated)
				if err := _YVelodromeVoterRegistry.contract.UnpackLog(event, "GaugeCreated", log); err != nil {
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
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryFilterer) ParseGaugeCreated(log types.Log) (*YVelodromeVoterRegistryGaugeCreated, error) {
	event := new(YVelodromeVoterRegistryGaugeCreated)
	if err := _YVelodromeVoterRegistry.contract.UnpackLog(event, "GaugeCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YVelodromeVoterRegistryGaugeKilledIterator is returned from FilterGaugeKilled and is used to iterate over the raw logs and unpacked data for GaugeKilled events raised by the YVelodromeVoterRegistry contract.
type YVelodromeVoterRegistryGaugeKilledIterator struct {
	Event *YVelodromeVoterRegistryGaugeKilled // Event containing the contract specifics and raw log

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
func (it *YVelodromeVoterRegistryGaugeKilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YVelodromeVoterRegistryGaugeKilled)
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
		it.Event = new(YVelodromeVoterRegistryGaugeKilled)
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
func (it *YVelodromeVoterRegistryGaugeKilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YVelodromeVoterRegistryGaugeKilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YVelodromeVoterRegistryGaugeKilled represents a GaugeKilled event raised by the YVelodromeVoterRegistry contract.
type YVelodromeVoterRegistryGaugeKilled struct {
	Gauge common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGaugeKilled is a free log retrieval operation binding the contract event 0x04a5d3f5d80d22d9345acc80618f4a4e7e663cf9e1aed23b57d975acec002ba7.
//
// Solidity: event GaugeKilled(address indexed gauge)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryFilterer) FilterGaugeKilled(opts *bind.FilterOpts, gauge []common.Address) (*YVelodromeVoterRegistryGaugeKilledIterator, error) {

	var gaugeRule []interface{}
	for _, gaugeItem := range gauge {
		gaugeRule = append(gaugeRule, gaugeItem)
	}

	logs, sub, err := _YVelodromeVoterRegistry.contract.FilterLogs(opts, "GaugeKilled", gaugeRule)
	if err != nil {
		return nil, err
	}
	return &YVelodromeVoterRegistryGaugeKilledIterator{contract: _YVelodromeVoterRegistry.contract, event: "GaugeKilled", logs: logs, sub: sub}, nil
}

// WatchGaugeKilled is a free log subscription operation binding the contract event 0x04a5d3f5d80d22d9345acc80618f4a4e7e663cf9e1aed23b57d975acec002ba7.
//
// Solidity: event GaugeKilled(address indexed gauge)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryFilterer) WatchGaugeKilled(opts *bind.WatchOpts, sink chan<- *YVelodromeVoterRegistryGaugeKilled, gauge []common.Address) (event.Subscription, error) {

	var gaugeRule []interface{}
	for _, gaugeItem := range gauge {
		gaugeRule = append(gaugeRule, gaugeItem)
	}

	logs, sub, err := _YVelodromeVoterRegistry.contract.WatchLogs(opts, "GaugeKilled", gaugeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YVelodromeVoterRegistryGaugeKilled)
				if err := _YVelodromeVoterRegistry.contract.UnpackLog(event, "GaugeKilled", log); err != nil {
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
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryFilterer) ParseGaugeKilled(log types.Log) (*YVelodromeVoterRegistryGaugeKilled, error) {
	event := new(YVelodromeVoterRegistryGaugeKilled)
	if err := _YVelodromeVoterRegistry.contract.UnpackLog(event, "GaugeKilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YVelodromeVoterRegistryGaugeRevivedIterator is returned from FilterGaugeRevived and is used to iterate over the raw logs and unpacked data for GaugeRevived events raised by the YVelodromeVoterRegistry contract.
type YVelodromeVoterRegistryGaugeRevivedIterator struct {
	Event *YVelodromeVoterRegistryGaugeRevived // Event containing the contract specifics and raw log

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
func (it *YVelodromeVoterRegistryGaugeRevivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YVelodromeVoterRegistryGaugeRevived)
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
		it.Event = new(YVelodromeVoterRegistryGaugeRevived)
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
func (it *YVelodromeVoterRegistryGaugeRevivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YVelodromeVoterRegistryGaugeRevivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YVelodromeVoterRegistryGaugeRevived represents a GaugeRevived event raised by the YVelodromeVoterRegistry contract.
type YVelodromeVoterRegistryGaugeRevived struct {
	Gauge common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGaugeRevived is a free log retrieval operation binding the contract event 0xed18e9faa3dccfd8aa45f69c4de40546b2ca9cccc4538a2323531656516db1aa.
//
// Solidity: event GaugeRevived(address indexed gauge)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryFilterer) FilterGaugeRevived(opts *bind.FilterOpts, gauge []common.Address) (*YVelodromeVoterRegistryGaugeRevivedIterator, error) {

	var gaugeRule []interface{}
	for _, gaugeItem := range gauge {
		gaugeRule = append(gaugeRule, gaugeItem)
	}

	logs, sub, err := _YVelodromeVoterRegistry.contract.FilterLogs(opts, "GaugeRevived", gaugeRule)
	if err != nil {
		return nil, err
	}
	return &YVelodromeVoterRegistryGaugeRevivedIterator{contract: _YVelodromeVoterRegistry.contract, event: "GaugeRevived", logs: logs, sub: sub}, nil
}

// WatchGaugeRevived is a free log subscription operation binding the contract event 0xed18e9faa3dccfd8aa45f69c4de40546b2ca9cccc4538a2323531656516db1aa.
//
// Solidity: event GaugeRevived(address indexed gauge)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryFilterer) WatchGaugeRevived(opts *bind.WatchOpts, sink chan<- *YVelodromeVoterRegistryGaugeRevived, gauge []common.Address) (event.Subscription, error) {

	var gaugeRule []interface{}
	for _, gaugeItem := range gauge {
		gaugeRule = append(gaugeRule, gaugeItem)
	}

	logs, sub, err := _YVelodromeVoterRegistry.contract.WatchLogs(opts, "GaugeRevived", gaugeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YVelodromeVoterRegistryGaugeRevived)
				if err := _YVelodromeVoterRegistry.contract.UnpackLog(event, "GaugeRevived", log); err != nil {
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
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryFilterer) ParseGaugeRevived(log types.Log) (*YVelodromeVoterRegistryGaugeRevived, error) {
	event := new(YVelodromeVoterRegistryGaugeRevived)
	if err := _YVelodromeVoterRegistry.contract.UnpackLog(event, "GaugeRevived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YVelodromeVoterRegistryNotifyRewardIterator is returned from FilterNotifyReward and is used to iterate over the raw logs and unpacked data for NotifyReward events raised by the YVelodromeVoterRegistry contract.
type YVelodromeVoterRegistryNotifyRewardIterator struct {
	Event *YVelodromeVoterRegistryNotifyReward // Event containing the contract specifics and raw log

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
func (it *YVelodromeVoterRegistryNotifyRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YVelodromeVoterRegistryNotifyReward)
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
		it.Event = new(YVelodromeVoterRegistryNotifyReward)
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
func (it *YVelodromeVoterRegistryNotifyRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YVelodromeVoterRegistryNotifyRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YVelodromeVoterRegistryNotifyReward represents a NotifyReward event raised by the YVelodromeVoterRegistry contract.
type YVelodromeVoterRegistryNotifyReward struct {
	Sender common.Address
	Reward common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNotifyReward is a free log retrieval operation binding the contract event 0xf70d5c697de7ea828df48e5c4573cb2194c659f1901f70110c52b066dcf50826.
//
// Solidity: event NotifyReward(address indexed sender, address indexed reward, uint256 amount)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryFilterer) FilterNotifyReward(opts *bind.FilterOpts, sender []common.Address, reward []common.Address) (*YVelodromeVoterRegistryNotifyRewardIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}

	logs, sub, err := _YVelodromeVoterRegistry.contract.FilterLogs(opts, "NotifyReward", senderRule, rewardRule)
	if err != nil {
		return nil, err
	}
	return &YVelodromeVoterRegistryNotifyRewardIterator{contract: _YVelodromeVoterRegistry.contract, event: "NotifyReward", logs: logs, sub: sub}, nil
}

// WatchNotifyReward is a free log subscription operation binding the contract event 0xf70d5c697de7ea828df48e5c4573cb2194c659f1901f70110c52b066dcf50826.
//
// Solidity: event NotifyReward(address indexed sender, address indexed reward, uint256 amount)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryFilterer) WatchNotifyReward(opts *bind.WatchOpts, sink chan<- *YVelodromeVoterRegistryNotifyReward, sender []common.Address, reward []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}

	logs, sub, err := _YVelodromeVoterRegistry.contract.WatchLogs(opts, "NotifyReward", senderRule, rewardRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YVelodromeVoterRegistryNotifyReward)
				if err := _YVelodromeVoterRegistry.contract.UnpackLog(event, "NotifyReward", log); err != nil {
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
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryFilterer) ParseNotifyReward(log types.Log) (*YVelodromeVoterRegistryNotifyReward, error) {
	event := new(YVelodromeVoterRegistryNotifyReward)
	if err := _YVelodromeVoterRegistry.contract.UnpackLog(event, "NotifyReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YVelodromeVoterRegistryVotedIterator is returned from FilterVoted and is used to iterate over the raw logs and unpacked data for Voted events raised by the YVelodromeVoterRegistry contract.
type YVelodromeVoterRegistryVotedIterator struct {
	Event *YVelodromeVoterRegistryVoted // Event containing the contract specifics and raw log

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
func (it *YVelodromeVoterRegistryVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YVelodromeVoterRegistryVoted)
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
		it.Event = new(YVelodromeVoterRegistryVoted)
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
func (it *YVelodromeVoterRegistryVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YVelodromeVoterRegistryVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YVelodromeVoterRegistryVoted represents a Voted event raised by the YVelodromeVoterRegistry contract.
type YVelodromeVoterRegistryVoted struct {
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
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryFilterer) FilterVoted(opts *bind.FilterOpts, voter []common.Address, pool []common.Address, tokenId []*big.Int) (*YVelodromeVoterRegistryVotedIterator, error) {

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

	logs, sub, err := _YVelodromeVoterRegistry.contract.FilterLogs(opts, "Voted", voterRule, poolRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &YVelodromeVoterRegistryVotedIterator{contract: _YVelodromeVoterRegistry.contract, event: "Voted", logs: logs, sub: sub}, nil
}

// WatchVoted is a free log subscription operation binding the contract event 0x452d440efc30dfa14a0ef803ccb55936af860ec6a6960ed27f129bef913f296a.
//
// Solidity: event Voted(address indexed voter, address indexed pool, uint256 indexed tokenId, uint256 weight, uint256 totalWeight, uint256 timestamp)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryFilterer) WatchVoted(opts *bind.WatchOpts, sink chan<- *YVelodromeVoterRegistryVoted, voter []common.Address, pool []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _YVelodromeVoterRegistry.contract.WatchLogs(opts, "Voted", voterRule, poolRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YVelodromeVoterRegistryVoted)
				if err := _YVelodromeVoterRegistry.contract.UnpackLog(event, "Voted", log); err != nil {
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
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryFilterer) ParseVoted(log types.Log) (*YVelodromeVoterRegistryVoted, error) {
	event := new(YVelodromeVoterRegistryVoted)
	if err := _YVelodromeVoterRegistry.contract.UnpackLog(event, "Voted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YVelodromeVoterRegistryWhitelistNFTIterator is returned from FilterWhitelistNFT and is used to iterate over the raw logs and unpacked data for WhitelistNFT events raised by the YVelodromeVoterRegistry contract.
type YVelodromeVoterRegistryWhitelistNFTIterator struct {
	Event *YVelodromeVoterRegistryWhitelistNFT // Event containing the contract specifics and raw log

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
func (it *YVelodromeVoterRegistryWhitelistNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YVelodromeVoterRegistryWhitelistNFT)
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
		it.Event = new(YVelodromeVoterRegistryWhitelistNFT)
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
func (it *YVelodromeVoterRegistryWhitelistNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YVelodromeVoterRegistryWhitelistNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YVelodromeVoterRegistryWhitelistNFT represents a WhitelistNFT event raised by the YVelodromeVoterRegistry contract.
type YVelodromeVoterRegistryWhitelistNFT struct {
	Whitelister common.Address
	TokenId     *big.Int
	Bool        bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWhitelistNFT is a free log retrieval operation binding the contract event 0x8a6ff732c8641e1e34d771e1f8b1673e988c1abdfb694ebdf6c910a5e3d0d853.
//
// Solidity: event WhitelistNFT(address indexed whitelister, uint256 indexed tokenId, bool indexed _bool)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryFilterer) FilterWhitelistNFT(opts *bind.FilterOpts, whitelister []common.Address, tokenId []*big.Int, _bool []bool) (*YVelodromeVoterRegistryWhitelistNFTIterator, error) {

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

	logs, sub, err := _YVelodromeVoterRegistry.contract.FilterLogs(opts, "WhitelistNFT", whitelisterRule, tokenIdRule, _boolRule)
	if err != nil {
		return nil, err
	}
	return &YVelodromeVoterRegistryWhitelistNFTIterator{contract: _YVelodromeVoterRegistry.contract, event: "WhitelistNFT", logs: logs, sub: sub}, nil
}

// WatchWhitelistNFT is a free log subscription operation binding the contract event 0x8a6ff732c8641e1e34d771e1f8b1673e988c1abdfb694ebdf6c910a5e3d0d853.
//
// Solidity: event WhitelistNFT(address indexed whitelister, uint256 indexed tokenId, bool indexed _bool)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryFilterer) WatchWhitelistNFT(opts *bind.WatchOpts, sink chan<- *YVelodromeVoterRegistryWhitelistNFT, whitelister []common.Address, tokenId []*big.Int, _bool []bool) (event.Subscription, error) {

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

	logs, sub, err := _YVelodromeVoterRegistry.contract.WatchLogs(opts, "WhitelistNFT", whitelisterRule, tokenIdRule, _boolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YVelodromeVoterRegistryWhitelistNFT)
				if err := _YVelodromeVoterRegistry.contract.UnpackLog(event, "WhitelistNFT", log); err != nil {
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
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryFilterer) ParseWhitelistNFT(log types.Log) (*YVelodromeVoterRegistryWhitelistNFT, error) {
	event := new(YVelodromeVoterRegistryWhitelistNFT)
	if err := _YVelodromeVoterRegistry.contract.UnpackLog(event, "WhitelistNFT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YVelodromeVoterRegistryWhitelistTokenIterator is returned from FilterWhitelistToken and is used to iterate over the raw logs and unpacked data for WhitelistToken events raised by the YVelodromeVoterRegistry contract.
type YVelodromeVoterRegistryWhitelistTokenIterator struct {
	Event *YVelodromeVoterRegistryWhitelistToken // Event containing the contract specifics and raw log

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
func (it *YVelodromeVoterRegistryWhitelistTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YVelodromeVoterRegistryWhitelistToken)
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
		it.Event = new(YVelodromeVoterRegistryWhitelistToken)
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
func (it *YVelodromeVoterRegistryWhitelistTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YVelodromeVoterRegistryWhitelistTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YVelodromeVoterRegistryWhitelistToken represents a WhitelistToken event raised by the YVelodromeVoterRegistry contract.
type YVelodromeVoterRegistryWhitelistToken struct {
	Whitelister common.Address
	Token       common.Address
	Bool        bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWhitelistToken is a free log retrieval operation binding the contract event 0x44948130cf88523dbc150908a47dd6332c33a01a3869d7f2fa78e51d5a5f9c57.
//
// Solidity: event WhitelistToken(address indexed whitelister, address indexed token, bool indexed _bool)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryFilterer) FilterWhitelistToken(opts *bind.FilterOpts, whitelister []common.Address, token []common.Address, _bool []bool) (*YVelodromeVoterRegistryWhitelistTokenIterator, error) {

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

	logs, sub, err := _YVelodromeVoterRegistry.contract.FilterLogs(opts, "WhitelistToken", whitelisterRule, tokenRule, _boolRule)
	if err != nil {
		return nil, err
	}
	return &YVelodromeVoterRegistryWhitelistTokenIterator{contract: _YVelodromeVoterRegistry.contract, event: "WhitelistToken", logs: logs, sub: sub}, nil
}

// WatchWhitelistToken is a free log subscription operation binding the contract event 0x44948130cf88523dbc150908a47dd6332c33a01a3869d7f2fa78e51d5a5f9c57.
//
// Solidity: event WhitelistToken(address indexed whitelister, address indexed token, bool indexed _bool)
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryFilterer) WatchWhitelistToken(opts *bind.WatchOpts, sink chan<- *YVelodromeVoterRegistryWhitelistToken, whitelister []common.Address, token []common.Address, _bool []bool) (event.Subscription, error) {

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

	logs, sub, err := _YVelodromeVoterRegistry.contract.WatchLogs(opts, "WhitelistToken", whitelisterRule, tokenRule, _boolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YVelodromeVoterRegistryWhitelistToken)
				if err := _YVelodromeVoterRegistry.contract.UnpackLog(event, "WhitelistToken", log); err != nil {
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
func (_YVelodromeVoterRegistry *YVelodromeVoterRegistryFilterer) ParseWhitelistToken(log types.Log) (*YVelodromeVoterRegistryWhitelistToken, error) {
	event := new(YVelodromeVoterRegistryWhitelistToken)
	if err := _YVelodromeVoterRegistry.contract.UnpackLog(event, "WhitelistToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
