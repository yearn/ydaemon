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

// CVXBoosterMetaData contains all meta data concerning the CVXBooster contract.
var CVXBoosterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"poolid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"poolid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FEE_DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MaxFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lptoken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gauge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stashVersion\",\"type\":\"uint256\"}],\"name\":\"addPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_gauge\",\"type\":\"address\"}],\"name\":\"claimRewards\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crv\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_stake\",\"type\":\"bool\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_stake\",\"type\":\"bool\"}],\"name\":\"depositAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributionAddressId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"earmarkFees\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"earmarkIncentive\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"earmarkRewards\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeDistro\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"gaugeMap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isShutdown\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockFees\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockIncentive\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockRewards\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"lptoken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"crvRewards\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stash\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"shutdown\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardArbitrator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"rewardClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_arb\",\"type\":\"address\"}],\"name\":\"setArbitrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rfactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sfactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tfactory\",\"type\":\"address\"}],\"name\":\"setFactories\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setFeeInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeM\",\"type\":\"address\"}],\"name\":\"setFeeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lockFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stakerFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_callerFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_platform\",\"type\":\"uint256\"}],\"name\":\"setFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"setGaugeRedirect\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolM\",\"type\":\"address\"}],\"name\":\"setPoolManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewards\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakerRewards\",\"type\":\"address\"}],\"name\":\"setRewardContracts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voteDelegate\",\"type\":\"address\"}],\"name\":\"setVoteDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"shutdownPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shutdownSystem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerIncentive\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerRewards\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stashFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_voteId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_votingAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_support\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voteDelegate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_gauge\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_weight\",\"type\":\"uint256[]\"}],\"name\":\"voteGaugeWeight\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voteOwnership\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voteParameter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"withdrawAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"withdrawTo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CVXBoosterABI is the input ABI used to generate the binding from.
// Deprecated: Use CVXBoosterMetaData.ABI instead.
var CVXBoosterABI = CVXBoosterMetaData.ABI

// CVXBooster is an auto generated Go binding around an Ethereum contract.
type CVXBooster struct {
	CVXBoosterCaller     // Read-only binding to the contract
	CVXBoosterTransactor // Write-only binding to the contract
	CVXBoosterFilterer   // Log filterer for contract events
}

// CVXBoosterCaller is an auto generated read-only Go binding around an Ethereum contract.
type CVXBoosterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CVXBoosterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CVXBoosterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CVXBoosterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CVXBoosterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CVXBoosterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CVXBoosterSession struct {
	Contract     *CVXBooster       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CVXBoosterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CVXBoosterCallerSession struct {
	Contract *CVXBoosterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CVXBoosterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CVXBoosterTransactorSession struct {
	Contract     *CVXBoosterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CVXBoosterRaw is an auto generated low-level Go binding around an Ethereum contract.
type CVXBoosterRaw struct {
	Contract *CVXBooster // Generic contract binding to access the raw methods on
}

// CVXBoosterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CVXBoosterCallerRaw struct {
	Contract *CVXBoosterCaller // Generic read-only contract binding to access the raw methods on
}

// CVXBoosterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CVXBoosterTransactorRaw struct {
	Contract *CVXBoosterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCVXBooster creates a new instance of CVXBooster, bound to a specific deployed contract.
func NewCVXBooster(address common.Address, backend bind.ContractBackend) (*CVXBooster, error) {
	contract, err := bindCVXBooster(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CVXBooster{CVXBoosterCaller: CVXBoosterCaller{contract: contract}, CVXBoosterTransactor: CVXBoosterTransactor{contract: contract}, CVXBoosterFilterer: CVXBoosterFilterer{contract: contract}}, nil
}

// NewCVXBoosterCaller creates a new read-only instance of CVXBooster, bound to a specific deployed contract.
func NewCVXBoosterCaller(address common.Address, caller bind.ContractCaller) (*CVXBoosterCaller, error) {
	contract, err := bindCVXBooster(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CVXBoosterCaller{contract: contract}, nil
}

// NewCVXBoosterTransactor creates a new write-only instance of CVXBooster, bound to a specific deployed contract.
func NewCVXBoosterTransactor(address common.Address, transactor bind.ContractTransactor) (*CVXBoosterTransactor, error) {
	contract, err := bindCVXBooster(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CVXBoosterTransactor{contract: contract}, nil
}

// NewCVXBoosterFilterer creates a new log filterer instance of CVXBooster, bound to a specific deployed contract.
func NewCVXBoosterFilterer(address common.Address, filterer bind.ContractFilterer) (*CVXBoosterFilterer, error) {
	contract, err := bindCVXBooster(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CVXBoosterFilterer{contract: contract}, nil
}

// bindCVXBooster binds a generic wrapper to an already deployed contract.
func bindCVXBooster(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CVXBoosterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CVXBooster *CVXBoosterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CVXBooster.Contract.CVXBoosterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CVXBooster *CVXBoosterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CVXBooster.Contract.CVXBoosterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CVXBooster *CVXBoosterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CVXBooster.Contract.CVXBoosterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CVXBooster *CVXBoosterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CVXBooster.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CVXBooster *CVXBoosterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CVXBooster.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CVXBooster *CVXBoosterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CVXBooster.Contract.contract.Transact(opts, method, params...)
}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint256)
func (_CVXBooster *CVXBoosterCaller) FEEDENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CVXBooster.contract.Call(opts, &out, "FEE_DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint256)
func (_CVXBooster *CVXBoosterSession) FEEDENOMINATOR() (*big.Int, error) {
	return _CVXBooster.Contract.FEEDENOMINATOR(&_CVXBooster.CallOpts)
}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint256)
func (_CVXBooster *CVXBoosterCallerSession) FEEDENOMINATOR() (*big.Int, error) {
	return _CVXBooster.Contract.FEEDENOMINATOR(&_CVXBooster.CallOpts)
}

// MaxFees is a free data retrieval call binding the contract method 0x7303df9a.
//
// Solidity: function MaxFees() view returns(uint256)
func (_CVXBooster *CVXBoosterCaller) MaxFees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CVXBooster.contract.Call(opts, &out, "MaxFees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxFees is a free data retrieval call binding the contract method 0x7303df9a.
//
// Solidity: function MaxFees() view returns(uint256)
func (_CVXBooster *CVXBoosterSession) MaxFees() (*big.Int, error) {
	return _CVXBooster.Contract.MaxFees(&_CVXBooster.CallOpts)
}

// MaxFees is a free data retrieval call binding the contract method 0x7303df9a.
//
// Solidity: function MaxFees() view returns(uint256)
func (_CVXBooster *CVXBoosterCallerSession) MaxFees() (*big.Int, error) {
	return _CVXBooster.Contract.MaxFees(&_CVXBooster.CallOpts)
}

// Crv is a free data retrieval call binding the contract method 0x6a4874a1.
//
// Solidity: function crv() view returns(address)
func (_CVXBooster *CVXBoosterCaller) Crv(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CVXBooster.contract.Call(opts, &out, "crv")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Crv is a free data retrieval call binding the contract method 0x6a4874a1.
//
// Solidity: function crv() view returns(address)
func (_CVXBooster *CVXBoosterSession) Crv() (common.Address, error) {
	return _CVXBooster.Contract.Crv(&_CVXBooster.CallOpts)
}

// Crv is a free data retrieval call binding the contract method 0x6a4874a1.
//
// Solidity: function crv() view returns(address)
func (_CVXBooster *CVXBoosterCallerSession) Crv() (common.Address, error) {
	return _CVXBooster.Contract.Crv(&_CVXBooster.CallOpts)
}

// DistributionAddressId is a free data retrieval call binding the contract method 0x93e846a0.
//
// Solidity: function distributionAddressId() view returns(uint256)
func (_CVXBooster *CVXBoosterCaller) DistributionAddressId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CVXBooster.contract.Call(opts, &out, "distributionAddressId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DistributionAddressId is a free data retrieval call binding the contract method 0x93e846a0.
//
// Solidity: function distributionAddressId() view returns(uint256)
func (_CVXBooster *CVXBoosterSession) DistributionAddressId() (*big.Int, error) {
	return _CVXBooster.Contract.DistributionAddressId(&_CVXBooster.CallOpts)
}

// DistributionAddressId is a free data retrieval call binding the contract method 0x93e846a0.
//
// Solidity: function distributionAddressId() view returns(uint256)
func (_CVXBooster *CVXBoosterCallerSession) DistributionAddressId() (*big.Int, error) {
	return _CVXBooster.Contract.DistributionAddressId(&_CVXBooster.CallOpts)
}

// EarmarkIncentive is a free data retrieval call binding the contract method 0x3a088cd2.
//
// Solidity: function earmarkIncentive() view returns(uint256)
func (_CVXBooster *CVXBoosterCaller) EarmarkIncentive(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CVXBooster.contract.Call(opts, &out, "earmarkIncentive")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EarmarkIncentive is a free data retrieval call binding the contract method 0x3a088cd2.
//
// Solidity: function earmarkIncentive() view returns(uint256)
func (_CVXBooster *CVXBoosterSession) EarmarkIncentive() (*big.Int, error) {
	return _CVXBooster.Contract.EarmarkIncentive(&_CVXBooster.CallOpts)
}

// EarmarkIncentive is a free data retrieval call binding the contract method 0x3a088cd2.
//
// Solidity: function earmarkIncentive() view returns(uint256)
func (_CVXBooster *CVXBoosterCallerSession) EarmarkIncentive() (*big.Int, error) {
	return _CVXBooster.Contract.EarmarkIncentive(&_CVXBooster.CallOpts)
}

// FeeDistro is a free data retrieval call binding the contract method 0xd6a0f530.
//
// Solidity: function feeDistro() view returns(address)
func (_CVXBooster *CVXBoosterCaller) FeeDistro(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CVXBooster.contract.Call(opts, &out, "feeDistro")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeDistro is a free data retrieval call binding the contract method 0xd6a0f530.
//
// Solidity: function feeDistro() view returns(address)
func (_CVXBooster *CVXBoosterSession) FeeDistro() (common.Address, error) {
	return _CVXBooster.Contract.FeeDistro(&_CVXBooster.CallOpts)
}

// FeeDistro is a free data retrieval call binding the contract method 0xd6a0f530.
//
// Solidity: function feeDistro() view returns(address)
func (_CVXBooster *CVXBoosterCallerSession) FeeDistro() (common.Address, error) {
	return _CVXBooster.Contract.FeeDistro(&_CVXBooster.CallOpts)
}

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() view returns(address)
func (_CVXBooster *CVXBoosterCaller) FeeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CVXBooster.contract.Call(opts, &out, "feeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() view returns(address)
func (_CVXBooster *CVXBoosterSession) FeeManager() (common.Address, error) {
	return _CVXBooster.Contract.FeeManager(&_CVXBooster.CallOpts)
}

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() view returns(address)
func (_CVXBooster *CVXBoosterCallerSession) FeeManager() (common.Address, error) {
	return _CVXBooster.Contract.FeeManager(&_CVXBooster.CallOpts)
}

// FeeToken is a free data retrieval call binding the contract method 0x647846a5.
//
// Solidity: function feeToken() view returns(address)
func (_CVXBooster *CVXBoosterCaller) FeeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CVXBooster.contract.Call(opts, &out, "feeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeToken is a free data retrieval call binding the contract method 0x647846a5.
//
// Solidity: function feeToken() view returns(address)
func (_CVXBooster *CVXBoosterSession) FeeToken() (common.Address, error) {
	return _CVXBooster.Contract.FeeToken(&_CVXBooster.CallOpts)
}

// FeeToken is a free data retrieval call binding the contract method 0x647846a5.
//
// Solidity: function feeToken() view returns(address)
func (_CVXBooster *CVXBoosterCallerSession) FeeToken() (common.Address, error) {
	return _CVXBooster.Contract.FeeToken(&_CVXBooster.CallOpts)
}

// GaugeMap is a free data retrieval call binding the contract method 0xcb0d5b52.
//
// Solidity: function gaugeMap(address ) view returns(bool)
func (_CVXBooster *CVXBoosterCaller) GaugeMap(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _CVXBooster.contract.Call(opts, &out, "gaugeMap", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GaugeMap is a free data retrieval call binding the contract method 0xcb0d5b52.
//
// Solidity: function gaugeMap(address ) view returns(bool)
func (_CVXBooster *CVXBoosterSession) GaugeMap(arg0 common.Address) (bool, error) {
	return _CVXBooster.Contract.GaugeMap(&_CVXBooster.CallOpts, arg0)
}

// GaugeMap is a free data retrieval call binding the contract method 0xcb0d5b52.
//
// Solidity: function gaugeMap(address ) view returns(bool)
func (_CVXBooster *CVXBoosterCallerSession) GaugeMap(arg0 common.Address) (bool, error) {
	return _CVXBooster.Contract.GaugeMap(&_CVXBooster.CallOpts, arg0)
}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_CVXBooster *CVXBoosterCaller) IsShutdown(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CVXBooster.contract.Call(opts, &out, "isShutdown")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_CVXBooster *CVXBoosterSession) IsShutdown() (bool, error) {
	return _CVXBooster.Contract.IsShutdown(&_CVXBooster.CallOpts)
}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_CVXBooster *CVXBoosterCallerSession) IsShutdown() (bool, error) {
	return _CVXBooster.Contract.IsShutdown(&_CVXBooster.CallOpts)
}

// LockFees is a free data retrieval call binding the contract method 0xab366292.
//
// Solidity: function lockFees() view returns(address)
func (_CVXBooster *CVXBoosterCaller) LockFees(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CVXBooster.contract.Call(opts, &out, "lockFees")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LockFees is a free data retrieval call binding the contract method 0xab366292.
//
// Solidity: function lockFees() view returns(address)
func (_CVXBooster *CVXBoosterSession) LockFees() (common.Address, error) {
	return _CVXBooster.Contract.LockFees(&_CVXBooster.CallOpts)
}

// LockFees is a free data retrieval call binding the contract method 0xab366292.
//
// Solidity: function lockFees() view returns(address)
func (_CVXBooster *CVXBoosterCallerSession) LockFees() (common.Address, error) {
	return _CVXBooster.Contract.LockFees(&_CVXBooster.CallOpts)
}

// LockIncentive is a free data retrieval call binding the contract method 0x50940618.
//
// Solidity: function lockIncentive() view returns(uint256)
func (_CVXBooster *CVXBoosterCaller) LockIncentive(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CVXBooster.contract.Call(opts, &out, "lockIncentive")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockIncentive is a free data retrieval call binding the contract method 0x50940618.
//
// Solidity: function lockIncentive() view returns(uint256)
func (_CVXBooster *CVXBoosterSession) LockIncentive() (*big.Int, error) {
	return _CVXBooster.Contract.LockIncentive(&_CVXBooster.CallOpts)
}

// LockIncentive is a free data retrieval call binding the contract method 0x50940618.
//
// Solidity: function lockIncentive() view returns(uint256)
func (_CVXBooster *CVXBoosterCallerSession) LockIncentive() (*big.Int, error) {
	return _CVXBooster.Contract.LockIncentive(&_CVXBooster.CallOpts)
}

// LockRewards is a free data retrieval call binding the contract method 0x376d771a.
//
// Solidity: function lockRewards() view returns(address)
func (_CVXBooster *CVXBoosterCaller) LockRewards(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CVXBooster.contract.Call(opts, &out, "lockRewards")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LockRewards is a free data retrieval call binding the contract method 0x376d771a.
//
// Solidity: function lockRewards() view returns(address)
func (_CVXBooster *CVXBoosterSession) LockRewards() (common.Address, error) {
	return _CVXBooster.Contract.LockRewards(&_CVXBooster.CallOpts)
}

// LockRewards is a free data retrieval call binding the contract method 0x376d771a.
//
// Solidity: function lockRewards() view returns(address)
func (_CVXBooster *CVXBoosterCallerSession) LockRewards() (common.Address, error) {
	return _CVXBooster.Contract.LockRewards(&_CVXBooster.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_CVXBooster *CVXBoosterCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CVXBooster.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_CVXBooster *CVXBoosterSession) Minter() (common.Address, error) {
	return _CVXBooster.Contract.Minter(&_CVXBooster.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_CVXBooster *CVXBoosterCallerSession) Minter() (common.Address, error) {
	return _CVXBooster.Contract.Minter(&_CVXBooster.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CVXBooster *CVXBoosterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CVXBooster.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CVXBooster *CVXBoosterSession) Owner() (common.Address, error) {
	return _CVXBooster.Contract.Owner(&_CVXBooster.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CVXBooster *CVXBoosterCallerSession) Owner() (common.Address, error) {
	return _CVXBooster.Contract.Owner(&_CVXBooster.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_CVXBooster *CVXBoosterCaller) PlatformFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CVXBooster.contract.Call(opts, &out, "platformFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_CVXBooster *CVXBoosterSession) PlatformFee() (*big.Int, error) {
	return _CVXBooster.Contract.PlatformFee(&_CVXBooster.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_CVXBooster *CVXBoosterCallerSession) PlatformFee() (*big.Int, error) {
	return _CVXBooster.Contract.PlatformFee(&_CVXBooster.CallOpts)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lptoken, address token, address gauge, address crvRewards, address stash, bool shutdown)
func (_CVXBooster *CVXBoosterCaller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Lptoken    common.Address
	Token      common.Address
	Gauge      common.Address
	CrvRewards common.Address
	Stash      common.Address
	Shutdown   bool
}, error) {
	var out []interface{}
	err := _CVXBooster.contract.Call(opts, &out, "poolInfo", arg0)

	outstruct := new(struct {
		Lptoken    common.Address
		Token      common.Address
		Gauge      common.Address
		CrvRewards common.Address
		Stash      common.Address
		Shutdown   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Lptoken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Token = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Gauge = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.CrvRewards = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Stash = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Shutdown = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lptoken, address token, address gauge, address crvRewards, address stash, bool shutdown)
func (_CVXBooster *CVXBoosterSession) PoolInfo(arg0 *big.Int) (struct {
	Lptoken    common.Address
	Token      common.Address
	Gauge      common.Address
	CrvRewards common.Address
	Stash      common.Address
	Shutdown   bool
}, error) {
	return _CVXBooster.Contract.PoolInfo(&_CVXBooster.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lptoken, address token, address gauge, address crvRewards, address stash, bool shutdown)
func (_CVXBooster *CVXBoosterCallerSession) PoolInfo(arg0 *big.Int) (struct {
	Lptoken    common.Address
	Token      common.Address
	Gauge      common.Address
	CrvRewards common.Address
	Stash      common.Address
	Shutdown   bool
}, error) {
	return _CVXBooster.Contract.PoolInfo(&_CVXBooster.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_CVXBooster *CVXBoosterCaller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CVXBooster.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_CVXBooster *CVXBoosterSession) PoolLength() (*big.Int, error) {
	return _CVXBooster.Contract.PoolLength(&_CVXBooster.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_CVXBooster *CVXBoosterCallerSession) PoolLength() (*big.Int, error) {
	return _CVXBooster.Contract.PoolLength(&_CVXBooster.CallOpts)
}

// PoolManager is a free data retrieval call binding the contract method 0xdc4c90d3.
//
// Solidity: function poolManager() view returns(address)
func (_CVXBooster *CVXBoosterCaller) PoolManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CVXBooster.contract.Call(opts, &out, "poolManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolManager is a free data retrieval call binding the contract method 0xdc4c90d3.
//
// Solidity: function poolManager() view returns(address)
func (_CVXBooster *CVXBoosterSession) PoolManager() (common.Address, error) {
	return _CVXBooster.Contract.PoolManager(&_CVXBooster.CallOpts)
}

// PoolManager is a free data retrieval call binding the contract method 0xdc4c90d3.
//
// Solidity: function poolManager() view returns(address)
func (_CVXBooster *CVXBoosterCallerSession) PoolManager() (common.Address, error) {
	return _CVXBooster.Contract.PoolManager(&_CVXBooster.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_CVXBooster *CVXBoosterCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CVXBooster.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_CVXBooster *CVXBoosterSession) Registry() (common.Address, error) {
	return _CVXBooster.Contract.Registry(&_CVXBooster.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_CVXBooster *CVXBoosterCallerSession) Registry() (common.Address, error) {
	return _CVXBooster.Contract.Registry(&_CVXBooster.CallOpts)
}

// RewardArbitrator is a free data retrieval call binding the contract method 0x043b684a.
//
// Solidity: function rewardArbitrator() view returns(address)
func (_CVXBooster *CVXBoosterCaller) RewardArbitrator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CVXBooster.contract.Call(opts, &out, "rewardArbitrator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardArbitrator is a free data retrieval call binding the contract method 0x043b684a.
//
// Solidity: function rewardArbitrator() view returns(address)
func (_CVXBooster *CVXBoosterSession) RewardArbitrator() (common.Address, error) {
	return _CVXBooster.Contract.RewardArbitrator(&_CVXBooster.CallOpts)
}

// RewardArbitrator is a free data retrieval call binding the contract method 0x043b684a.
//
// Solidity: function rewardArbitrator() view returns(address)
func (_CVXBooster *CVXBoosterCallerSession) RewardArbitrator() (common.Address, error) {
	return _CVXBooster.Contract.RewardArbitrator(&_CVXBooster.CallOpts)
}

// RewardFactory is a free data retrieval call binding the contract method 0x245e4bf0.
//
// Solidity: function rewardFactory() view returns(address)
func (_CVXBooster *CVXBoosterCaller) RewardFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CVXBooster.contract.Call(opts, &out, "rewardFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardFactory is a free data retrieval call binding the contract method 0x245e4bf0.
//
// Solidity: function rewardFactory() view returns(address)
func (_CVXBooster *CVXBoosterSession) RewardFactory() (common.Address, error) {
	return _CVXBooster.Contract.RewardFactory(&_CVXBooster.CallOpts)
}

// RewardFactory is a free data retrieval call binding the contract method 0x245e4bf0.
//
// Solidity: function rewardFactory() view returns(address)
func (_CVXBooster *CVXBoosterCallerSession) RewardFactory() (common.Address, error) {
	return _CVXBooster.Contract.RewardFactory(&_CVXBooster.CallOpts)
}

// Staker is a free data retrieval call binding the contract method 0x5ebaf1db.
//
// Solidity: function staker() view returns(address)
func (_CVXBooster *CVXBoosterCaller) Staker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CVXBooster.contract.Call(opts, &out, "staker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Staker is a free data retrieval call binding the contract method 0x5ebaf1db.
//
// Solidity: function staker() view returns(address)
func (_CVXBooster *CVXBoosterSession) Staker() (common.Address, error) {
	return _CVXBooster.Contract.Staker(&_CVXBooster.CallOpts)
}

// Staker is a free data retrieval call binding the contract method 0x5ebaf1db.
//
// Solidity: function staker() view returns(address)
func (_CVXBooster *CVXBoosterCallerSession) Staker() (common.Address, error) {
	return _CVXBooster.Contract.Staker(&_CVXBooster.CallOpts)
}

// StakerIncentive is a free data retrieval call binding the contract method 0x62d28ac7.
//
// Solidity: function stakerIncentive() view returns(uint256)
func (_CVXBooster *CVXBoosterCaller) StakerIncentive(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CVXBooster.contract.Call(opts, &out, "stakerIncentive")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerIncentive is a free data retrieval call binding the contract method 0x62d28ac7.
//
// Solidity: function stakerIncentive() view returns(uint256)
func (_CVXBooster *CVXBoosterSession) StakerIncentive() (*big.Int, error) {
	return _CVXBooster.Contract.StakerIncentive(&_CVXBooster.CallOpts)
}

// StakerIncentive is a free data retrieval call binding the contract method 0x62d28ac7.
//
// Solidity: function stakerIncentive() view returns(uint256)
func (_CVXBooster *CVXBoosterCallerSession) StakerIncentive() (*big.Int, error) {
	return _CVXBooster.Contract.StakerIncentive(&_CVXBooster.CallOpts)
}

// StakerRewards is a free data retrieval call binding the contract method 0xcfb9cfba.
//
// Solidity: function stakerRewards() view returns(address)
func (_CVXBooster *CVXBoosterCaller) StakerRewards(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CVXBooster.contract.Call(opts, &out, "stakerRewards")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakerRewards is a free data retrieval call binding the contract method 0xcfb9cfba.
//
// Solidity: function stakerRewards() view returns(address)
func (_CVXBooster *CVXBoosterSession) StakerRewards() (common.Address, error) {
	return _CVXBooster.Contract.StakerRewards(&_CVXBooster.CallOpts)
}

// StakerRewards is a free data retrieval call binding the contract method 0xcfb9cfba.
//
// Solidity: function stakerRewards() view returns(address)
func (_CVXBooster *CVXBoosterCallerSession) StakerRewards() (common.Address, error) {
	return _CVXBooster.Contract.StakerRewards(&_CVXBooster.CallOpts)
}

// StashFactory is a free data retrieval call binding the contract method 0x068eb19e.
//
// Solidity: function stashFactory() view returns(address)
func (_CVXBooster *CVXBoosterCaller) StashFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CVXBooster.contract.Call(opts, &out, "stashFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StashFactory is a free data retrieval call binding the contract method 0x068eb19e.
//
// Solidity: function stashFactory() view returns(address)
func (_CVXBooster *CVXBoosterSession) StashFactory() (common.Address, error) {
	return _CVXBooster.Contract.StashFactory(&_CVXBooster.CallOpts)
}

// StashFactory is a free data retrieval call binding the contract method 0x068eb19e.
//
// Solidity: function stashFactory() view returns(address)
func (_CVXBooster *CVXBoosterCallerSession) StashFactory() (common.Address, error) {
	return _CVXBooster.Contract.StashFactory(&_CVXBooster.CallOpts)
}

// TokenFactory is a free data retrieval call binding the contract method 0xe77772fe.
//
// Solidity: function tokenFactory() view returns(address)
func (_CVXBooster *CVXBoosterCaller) TokenFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CVXBooster.contract.Call(opts, &out, "tokenFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenFactory is a free data retrieval call binding the contract method 0xe77772fe.
//
// Solidity: function tokenFactory() view returns(address)
func (_CVXBooster *CVXBoosterSession) TokenFactory() (common.Address, error) {
	return _CVXBooster.Contract.TokenFactory(&_CVXBooster.CallOpts)
}

// TokenFactory is a free data retrieval call binding the contract method 0xe77772fe.
//
// Solidity: function tokenFactory() view returns(address)
func (_CVXBooster *CVXBoosterCallerSession) TokenFactory() (common.Address, error) {
	return _CVXBooster.Contract.TokenFactory(&_CVXBooster.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_CVXBooster *CVXBoosterCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CVXBooster.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_CVXBooster *CVXBoosterSession) Treasury() (common.Address, error) {
	return _CVXBooster.Contract.Treasury(&_CVXBooster.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_CVXBooster *CVXBoosterCallerSession) Treasury() (common.Address, error) {
	return _CVXBooster.Contract.Treasury(&_CVXBooster.CallOpts)
}

// VoteDelegate is a free data retrieval call binding the contract method 0x9f00332b.
//
// Solidity: function voteDelegate() view returns(address)
func (_CVXBooster *CVXBoosterCaller) VoteDelegate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CVXBooster.contract.Call(opts, &out, "voteDelegate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VoteDelegate is a free data retrieval call binding the contract method 0x9f00332b.
//
// Solidity: function voteDelegate() view returns(address)
func (_CVXBooster *CVXBoosterSession) VoteDelegate() (common.Address, error) {
	return _CVXBooster.Contract.VoteDelegate(&_CVXBooster.CallOpts)
}

// VoteDelegate is a free data retrieval call binding the contract method 0x9f00332b.
//
// Solidity: function voteDelegate() view returns(address)
func (_CVXBooster *CVXBoosterCallerSession) VoteDelegate() (common.Address, error) {
	return _CVXBooster.Contract.VoteDelegate(&_CVXBooster.CallOpts)
}

// VoteOwnership is a free data retrieval call binding the contract method 0xa386a080.
//
// Solidity: function voteOwnership() view returns(address)
func (_CVXBooster *CVXBoosterCaller) VoteOwnership(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CVXBooster.contract.Call(opts, &out, "voteOwnership")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VoteOwnership is a free data retrieval call binding the contract method 0xa386a080.
//
// Solidity: function voteOwnership() view returns(address)
func (_CVXBooster *CVXBoosterSession) VoteOwnership() (common.Address, error) {
	return _CVXBooster.Contract.VoteOwnership(&_CVXBooster.CallOpts)
}

// VoteOwnership is a free data retrieval call binding the contract method 0xa386a080.
//
// Solidity: function voteOwnership() view returns(address)
func (_CVXBooster *CVXBoosterCallerSession) VoteOwnership() (common.Address, error) {
	return _CVXBooster.Contract.VoteOwnership(&_CVXBooster.CallOpts)
}

// VoteParameter is a free data retrieval call binding the contract method 0xb42eda71.
//
// Solidity: function voteParameter() view returns(address)
func (_CVXBooster *CVXBoosterCaller) VoteParameter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CVXBooster.contract.Call(opts, &out, "voteParameter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VoteParameter is a free data retrieval call binding the contract method 0xb42eda71.
//
// Solidity: function voteParameter() view returns(address)
func (_CVXBooster *CVXBoosterSession) VoteParameter() (common.Address, error) {
	return _CVXBooster.Contract.VoteParameter(&_CVXBooster.CallOpts)
}

// VoteParameter is a free data retrieval call binding the contract method 0xb42eda71.
//
// Solidity: function voteParameter() view returns(address)
func (_CVXBooster *CVXBoosterCallerSession) VoteParameter() (common.Address, error) {
	return _CVXBooster.Contract.VoteParameter(&_CVXBooster.CallOpts)
}

// AddPool is a paid mutator transaction binding the contract method 0x7e29d6c2.
//
// Solidity: function addPool(address _lptoken, address _gauge, uint256 _stashVersion) returns(bool)
func (_CVXBooster *CVXBoosterTransactor) AddPool(opts *bind.TransactOpts, _lptoken common.Address, _gauge common.Address, _stashVersion *big.Int) (*types.Transaction, error) {
	return _CVXBooster.contract.Transact(opts, "addPool", _lptoken, _gauge, _stashVersion)
}

// AddPool is a paid mutator transaction binding the contract method 0x7e29d6c2.
//
// Solidity: function addPool(address _lptoken, address _gauge, uint256 _stashVersion) returns(bool)
func (_CVXBooster *CVXBoosterSession) AddPool(_lptoken common.Address, _gauge common.Address, _stashVersion *big.Int) (*types.Transaction, error) {
	return _CVXBooster.Contract.AddPool(&_CVXBooster.TransactOpts, _lptoken, _gauge, _stashVersion)
}

// AddPool is a paid mutator transaction binding the contract method 0x7e29d6c2.
//
// Solidity: function addPool(address _lptoken, address _gauge, uint256 _stashVersion) returns(bool)
func (_CVXBooster *CVXBoosterTransactorSession) AddPool(_lptoken common.Address, _gauge common.Address, _stashVersion *big.Int) (*types.Transaction, error) {
	return _CVXBooster.Contract.AddPool(&_CVXBooster.TransactOpts, _lptoken, _gauge, _stashVersion)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x6c7b69cb.
//
// Solidity: function claimRewards(uint256 _pid, address _gauge) returns(bool)
func (_CVXBooster *CVXBoosterTransactor) ClaimRewards(opts *bind.TransactOpts, _pid *big.Int, _gauge common.Address) (*types.Transaction, error) {
	return _CVXBooster.contract.Transact(opts, "claimRewards", _pid, _gauge)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x6c7b69cb.
//
// Solidity: function claimRewards(uint256 _pid, address _gauge) returns(bool)
func (_CVXBooster *CVXBoosterSession) ClaimRewards(_pid *big.Int, _gauge common.Address) (*types.Transaction, error) {
	return _CVXBooster.Contract.ClaimRewards(&_CVXBooster.TransactOpts, _pid, _gauge)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x6c7b69cb.
//
// Solidity: function claimRewards(uint256 _pid, address _gauge) returns(bool)
func (_CVXBooster *CVXBoosterTransactorSession) ClaimRewards(_pid *big.Int, _gauge common.Address) (*types.Transaction, error) {
	return _CVXBooster.Contract.ClaimRewards(&_CVXBooster.TransactOpts, _pid, _gauge)
}

// Deposit is a paid mutator transaction binding the contract method 0x43a0d066.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount, bool _stake) returns(bool)
func (_CVXBooster *CVXBoosterTransactor) Deposit(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int, _stake bool) (*types.Transaction, error) {
	return _CVXBooster.contract.Transact(opts, "deposit", _pid, _amount, _stake)
}

// Deposit is a paid mutator transaction binding the contract method 0x43a0d066.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount, bool _stake) returns(bool)
func (_CVXBooster *CVXBoosterSession) Deposit(_pid *big.Int, _amount *big.Int, _stake bool) (*types.Transaction, error) {
	return _CVXBooster.Contract.Deposit(&_CVXBooster.TransactOpts, _pid, _amount, _stake)
}

// Deposit is a paid mutator transaction binding the contract method 0x43a0d066.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount, bool _stake) returns(bool)
func (_CVXBooster *CVXBoosterTransactorSession) Deposit(_pid *big.Int, _amount *big.Int, _stake bool) (*types.Transaction, error) {
	return _CVXBooster.Contract.Deposit(&_CVXBooster.TransactOpts, _pid, _amount, _stake)
}

// DepositAll is a paid mutator transaction binding the contract method 0x60759fce.
//
// Solidity: function depositAll(uint256 _pid, bool _stake) returns(bool)
func (_CVXBooster *CVXBoosterTransactor) DepositAll(opts *bind.TransactOpts, _pid *big.Int, _stake bool) (*types.Transaction, error) {
	return _CVXBooster.contract.Transact(opts, "depositAll", _pid, _stake)
}

// DepositAll is a paid mutator transaction binding the contract method 0x60759fce.
//
// Solidity: function depositAll(uint256 _pid, bool _stake) returns(bool)
func (_CVXBooster *CVXBoosterSession) DepositAll(_pid *big.Int, _stake bool) (*types.Transaction, error) {
	return _CVXBooster.Contract.DepositAll(&_CVXBooster.TransactOpts, _pid, _stake)
}

// DepositAll is a paid mutator transaction binding the contract method 0x60759fce.
//
// Solidity: function depositAll(uint256 _pid, bool _stake) returns(bool)
func (_CVXBooster *CVXBoosterTransactorSession) DepositAll(_pid *big.Int, _stake bool) (*types.Transaction, error) {
	return _CVXBooster.Contract.DepositAll(&_CVXBooster.TransactOpts, _pid, _stake)
}

// EarmarkFees is a paid mutator transaction binding the contract method 0x22230b96.
//
// Solidity: function earmarkFees() returns(bool)
func (_CVXBooster *CVXBoosterTransactor) EarmarkFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CVXBooster.contract.Transact(opts, "earmarkFees")
}

// EarmarkFees is a paid mutator transaction binding the contract method 0x22230b96.
//
// Solidity: function earmarkFees() returns(bool)
func (_CVXBooster *CVXBoosterSession) EarmarkFees() (*types.Transaction, error) {
	return _CVXBooster.Contract.EarmarkFees(&_CVXBooster.TransactOpts)
}

// EarmarkFees is a paid mutator transaction binding the contract method 0x22230b96.
//
// Solidity: function earmarkFees() returns(bool)
func (_CVXBooster *CVXBoosterTransactorSession) EarmarkFees() (*types.Transaction, error) {
	return _CVXBooster.Contract.EarmarkFees(&_CVXBooster.TransactOpts)
}

// EarmarkRewards is a paid mutator transaction binding the contract method 0xcc956f3f.
//
// Solidity: function earmarkRewards(uint256 _pid) returns(bool)
func (_CVXBooster *CVXBoosterTransactor) EarmarkRewards(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _CVXBooster.contract.Transact(opts, "earmarkRewards", _pid)
}

// EarmarkRewards is a paid mutator transaction binding the contract method 0xcc956f3f.
//
// Solidity: function earmarkRewards(uint256 _pid) returns(bool)
func (_CVXBooster *CVXBoosterSession) EarmarkRewards(_pid *big.Int) (*types.Transaction, error) {
	return _CVXBooster.Contract.EarmarkRewards(&_CVXBooster.TransactOpts, _pid)
}

// EarmarkRewards is a paid mutator transaction binding the contract method 0xcc956f3f.
//
// Solidity: function earmarkRewards(uint256 _pid) returns(bool)
func (_CVXBooster *CVXBoosterTransactorSession) EarmarkRewards(_pid *big.Int) (*types.Transaction, error) {
	return _CVXBooster.Contract.EarmarkRewards(&_CVXBooster.TransactOpts, _pid)
}

// RewardClaimed is a paid mutator transaction binding the contract method 0x71192b17.
//
// Solidity: function rewardClaimed(uint256 _pid, address _address, uint256 _amount) returns(bool)
func (_CVXBooster *CVXBoosterTransactor) RewardClaimed(opts *bind.TransactOpts, _pid *big.Int, _address common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _CVXBooster.contract.Transact(opts, "rewardClaimed", _pid, _address, _amount)
}

// RewardClaimed is a paid mutator transaction binding the contract method 0x71192b17.
//
// Solidity: function rewardClaimed(uint256 _pid, address _address, uint256 _amount) returns(bool)
func (_CVXBooster *CVXBoosterSession) RewardClaimed(_pid *big.Int, _address common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _CVXBooster.Contract.RewardClaimed(&_CVXBooster.TransactOpts, _pid, _address, _amount)
}

// RewardClaimed is a paid mutator transaction binding the contract method 0x71192b17.
//
// Solidity: function rewardClaimed(uint256 _pid, address _address, uint256 _amount) returns(bool)
func (_CVXBooster *CVXBoosterTransactorSession) RewardClaimed(_pid *big.Int, _address common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _CVXBooster.Contract.RewardClaimed(&_CVXBooster.TransactOpts, _pid, _address, _amount)
}

// SetArbitrator is a paid mutator transaction binding the contract method 0xb0eefabe.
//
// Solidity: function setArbitrator(address _arb) returns()
func (_CVXBooster *CVXBoosterTransactor) SetArbitrator(opts *bind.TransactOpts, _arb common.Address) (*types.Transaction, error) {
	return _CVXBooster.contract.Transact(opts, "setArbitrator", _arb)
}

// SetArbitrator is a paid mutator transaction binding the contract method 0xb0eefabe.
//
// Solidity: function setArbitrator(address _arb) returns()
func (_CVXBooster *CVXBoosterSession) SetArbitrator(_arb common.Address) (*types.Transaction, error) {
	return _CVXBooster.Contract.SetArbitrator(&_CVXBooster.TransactOpts, _arb)
}

// SetArbitrator is a paid mutator transaction binding the contract method 0xb0eefabe.
//
// Solidity: function setArbitrator(address _arb) returns()
func (_CVXBooster *CVXBoosterTransactorSession) SetArbitrator(_arb common.Address) (*types.Transaction, error) {
	return _CVXBooster.Contract.SetArbitrator(&_CVXBooster.TransactOpts, _arb)
}

// SetFactories is a paid mutator transaction binding the contract method 0x7bd3b995.
//
// Solidity: function setFactories(address _rfactory, address _sfactory, address _tfactory) returns()
func (_CVXBooster *CVXBoosterTransactor) SetFactories(opts *bind.TransactOpts, _rfactory common.Address, _sfactory common.Address, _tfactory common.Address) (*types.Transaction, error) {
	return _CVXBooster.contract.Transact(opts, "setFactories", _rfactory, _sfactory, _tfactory)
}

// SetFactories is a paid mutator transaction binding the contract method 0x7bd3b995.
//
// Solidity: function setFactories(address _rfactory, address _sfactory, address _tfactory) returns()
func (_CVXBooster *CVXBoosterSession) SetFactories(_rfactory common.Address, _sfactory common.Address, _tfactory common.Address) (*types.Transaction, error) {
	return _CVXBooster.Contract.SetFactories(&_CVXBooster.TransactOpts, _rfactory, _sfactory, _tfactory)
}

// SetFactories is a paid mutator transaction binding the contract method 0x7bd3b995.
//
// Solidity: function setFactories(address _rfactory, address _sfactory, address _tfactory) returns()
func (_CVXBooster *CVXBoosterTransactorSession) SetFactories(_rfactory common.Address, _sfactory common.Address, _tfactory common.Address) (*types.Transaction, error) {
	return _CVXBooster.Contract.SetFactories(&_CVXBooster.TransactOpts, _rfactory, _sfactory, _tfactory)
}

// SetFeeInfo is a paid mutator transaction binding the contract method 0x5a4ae5ca.
//
// Solidity: function setFeeInfo() returns()
func (_CVXBooster *CVXBoosterTransactor) SetFeeInfo(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CVXBooster.contract.Transact(opts, "setFeeInfo")
}

// SetFeeInfo is a paid mutator transaction binding the contract method 0x5a4ae5ca.
//
// Solidity: function setFeeInfo() returns()
func (_CVXBooster *CVXBoosterSession) SetFeeInfo() (*types.Transaction, error) {
	return _CVXBooster.Contract.SetFeeInfo(&_CVXBooster.TransactOpts)
}

// SetFeeInfo is a paid mutator transaction binding the contract method 0x5a4ae5ca.
//
// Solidity: function setFeeInfo() returns()
func (_CVXBooster *CVXBoosterTransactorSession) SetFeeInfo() (*types.Transaction, error) {
	return _CVXBooster.Contract.SetFeeInfo(&_CVXBooster.TransactOpts)
}

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address _feeM) returns()
func (_CVXBooster *CVXBoosterTransactor) SetFeeManager(opts *bind.TransactOpts, _feeM common.Address) (*types.Transaction, error) {
	return _CVXBooster.contract.Transact(opts, "setFeeManager", _feeM)
}

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address _feeM) returns()
func (_CVXBooster *CVXBoosterSession) SetFeeManager(_feeM common.Address) (*types.Transaction, error) {
	return _CVXBooster.Contract.SetFeeManager(&_CVXBooster.TransactOpts, _feeM)
}

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address _feeM) returns()
func (_CVXBooster *CVXBoosterTransactorSession) SetFeeManager(_feeM common.Address) (*types.Transaction, error) {
	return _CVXBooster.Contract.SetFeeManager(&_CVXBooster.TransactOpts, _feeM)
}

// SetFees is a paid mutator transaction binding the contract method 0x6fcba377.
//
// Solidity: function setFees(uint256 _lockFees, uint256 _stakerFees, uint256 _callerFees, uint256 _platform) returns()
func (_CVXBooster *CVXBoosterTransactor) SetFees(opts *bind.TransactOpts, _lockFees *big.Int, _stakerFees *big.Int, _callerFees *big.Int, _platform *big.Int) (*types.Transaction, error) {
	return _CVXBooster.contract.Transact(opts, "setFees", _lockFees, _stakerFees, _callerFees, _platform)
}

// SetFees is a paid mutator transaction binding the contract method 0x6fcba377.
//
// Solidity: function setFees(uint256 _lockFees, uint256 _stakerFees, uint256 _callerFees, uint256 _platform) returns()
func (_CVXBooster *CVXBoosterSession) SetFees(_lockFees *big.Int, _stakerFees *big.Int, _callerFees *big.Int, _platform *big.Int) (*types.Transaction, error) {
	return _CVXBooster.Contract.SetFees(&_CVXBooster.TransactOpts, _lockFees, _stakerFees, _callerFees, _platform)
}

// SetFees is a paid mutator transaction binding the contract method 0x6fcba377.
//
// Solidity: function setFees(uint256 _lockFees, uint256 _stakerFees, uint256 _callerFees, uint256 _platform) returns()
func (_CVXBooster *CVXBoosterTransactorSession) SetFees(_lockFees *big.Int, _stakerFees *big.Int, _callerFees *big.Int, _platform *big.Int) (*types.Transaction, error) {
	return _CVXBooster.Contract.SetFees(&_CVXBooster.TransactOpts, _lockFees, _stakerFees, _callerFees, _platform)
}

// SetGaugeRedirect is a paid mutator transaction binding the contract method 0x9123d404.
//
// Solidity: function setGaugeRedirect(uint256 _pid) returns(bool)
func (_CVXBooster *CVXBoosterTransactor) SetGaugeRedirect(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _CVXBooster.contract.Transact(opts, "setGaugeRedirect", _pid)
}

// SetGaugeRedirect is a paid mutator transaction binding the contract method 0x9123d404.
//
// Solidity: function setGaugeRedirect(uint256 _pid) returns(bool)
func (_CVXBooster *CVXBoosterSession) SetGaugeRedirect(_pid *big.Int) (*types.Transaction, error) {
	return _CVXBooster.Contract.SetGaugeRedirect(&_CVXBooster.TransactOpts, _pid)
}

// SetGaugeRedirect is a paid mutator transaction binding the contract method 0x9123d404.
//
// Solidity: function setGaugeRedirect(uint256 _pid) returns(bool)
func (_CVXBooster *CVXBoosterTransactorSession) SetGaugeRedirect(_pid *big.Int) (*types.Transaction, error) {
	return _CVXBooster.Contract.SetGaugeRedirect(&_CVXBooster.TransactOpts, _pid)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_CVXBooster *CVXBoosterTransactor) SetOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _CVXBooster.contract.Transact(opts, "setOwner", _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_CVXBooster *CVXBoosterSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _CVXBooster.Contract.SetOwner(&_CVXBooster.TransactOpts, _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_CVXBooster *CVXBoosterTransactorSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _CVXBooster.Contract.SetOwner(&_CVXBooster.TransactOpts, _owner)
}

// SetPoolManager is a paid mutator transaction binding the contract method 0x7aef6715.
//
// Solidity: function setPoolManager(address _poolM) returns()
func (_CVXBooster *CVXBoosterTransactor) SetPoolManager(opts *bind.TransactOpts, _poolM common.Address) (*types.Transaction, error) {
	return _CVXBooster.contract.Transact(opts, "setPoolManager", _poolM)
}

// SetPoolManager is a paid mutator transaction binding the contract method 0x7aef6715.
//
// Solidity: function setPoolManager(address _poolM) returns()
func (_CVXBooster *CVXBoosterSession) SetPoolManager(_poolM common.Address) (*types.Transaction, error) {
	return _CVXBooster.Contract.SetPoolManager(&_CVXBooster.TransactOpts, _poolM)
}

// SetPoolManager is a paid mutator transaction binding the contract method 0x7aef6715.
//
// Solidity: function setPoolManager(address _poolM) returns()
func (_CVXBooster *CVXBoosterTransactorSession) SetPoolManager(_poolM common.Address) (*types.Transaction, error) {
	return _CVXBooster.Contract.SetPoolManager(&_CVXBooster.TransactOpts, _poolM)
}

// SetRewardContracts is a paid mutator transaction binding the contract method 0x95539a1d.
//
// Solidity: function setRewardContracts(address _rewards, address _stakerRewards) returns()
func (_CVXBooster *CVXBoosterTransactor) SetRewardContracts(opts *bind.TransactOpts, _rewards common.Address, _stakerRewards common.Address) (*types.Transaction, error) {
	return _CVXBooster.contract.Transact(opts, "setRewardContracts", _rewards, _stakerRewards)
}

// SetRewardContracts is a paid mutator transaction binding the contract method 0x95539a1d.
//
// Solidity: function setRewardContracts(address _rewards, address _stakerRewards) returns()
func (_CVXBooster *CVXBoosterSession) SetRewardContracts(_rewards common.Address, _stakerRewards common.Address) (*types.Transaction, error) {
	return _CVXBooster.Contract.SetRewardContracts(&_CVXBooster.TransactOpts, _rewards, _stakerRewards)
}

// SetRewardContracts is a paid mutator transaction binding the contract method 0x95539a1d.
//
// Solidity: function setRewardContracts(address _rewards, address _stakerRewards) returns()
func (_CVXBooster *CVXBoosterTransactorSession) SetRewardContracts(_rewards common.Address, _stakerRewards common.Address) (*types.Transaction, error) {
	return _CVXBooster.Contract.SetRewardContracts(&_CVXBooster.TransactOpts, _rewards, _stakerRewards)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _treasury) returns()
func (_CVXBooster *CVXBoosterTransactor) SetTreasury(opts *bind.TransactOpts, _treasury common.Address) (*types.Transaction, error) {
	return _CVXBooster.contract.Transact(opts, "setTreasury", _treasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _treasury) returns()
func (_CVXBooster *CVXBoosterSession) SetTreasury(_treasury common.Address) (*types.Transaction, error) {
	return _CVXBooster.Contract.SetTreasury(&_CVXBooster.TransactOpts, _treasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _treasury) returns()
func (_CVXBooster *CVXBoosterTransactorSession) SetTreasury(_treasury common.Address) (*types.Transaction, error) {
	return _CVXBooster.Contract.SetTreasury(&_CVXBooster.TransactOpts, _treasury)
}

// SetVoteDelegate is a paid mutator transaction binding the contract method 0x74874323.
//
// Solidity: function setVoteDelegate(address _voteDelegate) returns()
func (_CVXBooster *CVXBoosterTransactor) SetVoteDelegate(opts *bind.TransactOpts, _voteDelegate common.Address) (*types.Transaction, error) {
	return _CVXBooster.contract.Transact(opts, "setVoteDelegate", _voteDelegate)
}

// SetVoteDelegate is a paid mutator transaction binding the contract method 0x74874323.
//
// Solidity: function setVoteDelegate(address _voteDelegate) returns()
func (_CVXBooster *CVXBoosterSession) SetVoteDelegate(_voteDelegate common.Address) (*types.Transaction, error) {
	return _CVXBooster.Contract.SetVoteDelegate(&_CVXBooster.TransactOpts, _voteDelegate)
}

// SetVoteDelegate is a paid mutator transaction binding the contract method 0x74874323.
//
// Solidity: function setVoteDelegate(address _voteDelegate) returns()
func (_CVXBooster *CVXBoosterTransactorSession) SetVoteDelegate(_voteDelegate common.Address) (*types.Transaction, error) {
	return _CVXBooster.Contract.SetVoteDelegate(&_CVXBooster.TransactOpts, _voteDelegate)
}

// ShutdownPool is a paid mutator transaction binding the contract method 0x60cafe84.
//
// Solidity: function shutdownPool(uint256 _pid) returns(bool)
func (_CVXBooster *CVXBoosterTransactor) ShutdownPool(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _CVXBooster.contract.Transact(opts, "shutdownPool", _pid)
}

// ShutdownPool is a paid mutator transaction binding the contract method 0x60cafe84.
//
// Solidity: function shutdownPool(uint256 _pid) returns(bool)
func (_CVXBooster *CVXBoosterSession) ShutdownPool(_pid *big.Int) (*types.Transaction, error) {
	return _CVXBooster.Contract.ShutdownPool(&_CVXBooster.TransactOpts, _pid)
}

// ShutdownPool is a paid mutator transaction binding the contract method 0x60cafe84.
//
// Solidity: function shutdownPool(uint256 _pid) returns(bool)
func (_CVXBooster *CVXBoosterTransactorSession) ShutdownPool(_pid *big.Int) (*types.Transaction, error) {
	return _CVXBooster.Contract.ShutdownPool(&_CVXBooster.TransactOpts, _pid)
}

// ShutdownSystem is a paid mutator transaction binding the contract method 0x354af919.
//
// Solidity: function shutdownSystem() returns()
func (_CVXBooster *CVXBoosterTransactor) ShutdownSystem(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CVXBooster.contract.Transact(opts, "shutdownSystem")
}

// ShutdownSystem is a paid mutator transaction binding the contract method 0x354af919.
//
// Solidity: function shutdownSystem() returns()
func (_CVXBooster *CVXBoosterSession) ShutdownSystem() (*types.Transaction, error) {
	return _CVXBooster.Contract.ShutdownSystem(&_CVXBooster.TransactOpts)
}

// ShutdownSystem is a paid mutator transaction binding the contract method 0x354af919.
//
// Solidity: function shutdownSystem() returns()
func (_CVXBooster *CVXBoosterTransactorSession) ShutdownSystem() (*types.Transaction, error) {
	return _CVXBooster.Contract.ShutdownSystem(&_CVXBooster.TransactOpts)
}

// Vote is a paid mutator transaction binding the contract method 0xe2cdd42a.
//
// Solidity: function vote(uint256 _voteId, address _votingAddress, bool _support) returns(bool)
func (_CVXBooster *CVXBoosterTransactor) Vote(opts *bind.TransactOpts, _voteId *big.Int, _votingAddress common.Address, _support bool) (*types.Transaction, error) {
	return _CVXBooster.contract.Transact(opts, "vote", _voteId, _votingAddress, _support)
}

// Vote is a paid mutator transaction binding the contract method 0xe2cdd42a.
//
// Solidity: function vote(uint256 _voteId, address _votingAddress, bool _support) returns(bool)
func (_CVXBooster *CVXBoosterSession) Vote(_voteId *big.Int, _votingAddress common.Address, _support bool) (*types.Transaction, error) {
	return _CVXBooster.Contract.Vote(&_CVXBooster.TransactOpts, _voteId, _votingAddress, _support)
}

// Vote is a paid mutator transaction binding the contract method 0xe2cdd42a.
//
// Solidity: function vote(uint256 _voteId, address _votingAddress, bool _support) returns(bool)
func (_CVXBooster *CVXBoosterTransactorSession) Vote(_voteId *big.Int, _votingAddress common.Address, _support bool) (*types.Transaction, error) {
	return _CVXBooster.Contract.Vote(&_CVXBooster.TransactOpts, _voteId, _votingAddress, _support)
}

// VoteGaugeWeight is a paid mutator transaction binding the contract method 0xbfad96ba.
//
// Solidity: function voteGaugeWeight(address[] _gauge, uint256[] _weight) returns(bool)
func (_CVXBooster *CVXBoosterTransactor) VoteGaugeWeight(opts *bind.TransactOpts, _gauge []common.Address, _weight []*big.Int) (*types.Transaction, error) {
	return _CVXBooster.contract.Transact(opts, "voteGaugeWeight", _gauge, _weight)
}

// VoteGaugeWeight is a paid mutator transaction binding the contract method 0xbfad96ba.
//
// Solidity: function voteGaugeWeight(address[] _gauge, uint256[] _weight) returns(bool)
func (_CVXBooster *CVXBoosterSession) VoteGaugeWeight(_gauge []common.Address, _weight []*big.Int) (*types.Transaction, error) {
	return _CVXBooster.Contract.VoteGaugeWeight(&_CVXBooster.TransactOpts, _gauge, _weight)
}

// VoteGaugeWeight is a paid mutator transaction binding the contract method 0xbfad96ba.
//
// Solidity: function voteGaugeWeight(address[] _gauge, uint256[] _weight) returns(bool)
func (_CVXBooster *CVXBoosterTransactorSession) VoteGaugeWeight(_gauge []common.Address, _weight []*big.Int) (*types.Transaction, error) {
	return _CVXBooster.Contract.VoteGaugeWeight(&_CVXBooster.TransactOpts, _gauge, _weight)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns(bool)
func (_CVXBooster *CVXBoosterTransactor) Withdraw(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _CVXBooster.contract.Transact(opts, "withdraw", _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns(bool)
func (_CVXBooster *CVXBoosterSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _CVXBooster.Contract.Withdraw(&_CVXBooster.TransactOpts, _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns(bool)
func (_CVXBooster *CVXBoosterTransactorSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _CVXBooster.Contract.Withdraw(&_CVXBooster.TransactOpts, _pid, _amount)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x958e2d31.
//
// Solidity: function withdrawAll(uint256 _pid) returns(bool)
func (_CVXBooster *CVXBoosterTransactor) WithdrawAll(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _CVXBooster.contract.Transact(opts, "withdrawAll", _pid)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x958e2d31.
//
// Solidity: function withdrawAll(uint256 _pid) returns(bool)
func (_CVXBooster *CVXBoosterSession) WithdrawAll(_pid *big.Int) (*types.Transaction, error) {
	return _CVXBooster.Contract.WithdrawAll(&_CVXBooster.TransactOpts, _pid)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x958e2d31.
//
// Solidity: function withdrawAll(uint256 _pid) returns(bool)
func (_CVXBooster *CVXBoosterTransactorSession) WithdrawAll(_pid *big.Int) (*types.Transaction, error) {
	return _CVXBooster.Contract.WithdrawAll(&_CVXBooster.TransactOpts, _pid)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x14cd70e4.
//
// Solidity: function withdrawTo(uint256 _pid, uint256 _amount, address _to) returns(bool)
func (_CVXBooster *CVXBoosterTransactor) WithdrawTo(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _CVXBooster.contract.Transact(opts, "withdrawTo", _pid, _amount, _to)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x14cd70e4.
//
// Solidity: function withdrawTo(uint256 _pid, uint256 _amount, address _to) returns(bool)
func (_CVXBooster *CVXBoosterSession) WithdrawTo(_pid *big.Int, _amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _CVXBooster.Contract.WithdrawTo(&_CVXBooster.TransactOpts, _pid, _amount, _to)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x14cd70e4.
//
// Solidity: function withdrawTo(uint256 _pid, uint256 _amount, address _to) returns(bool)
func (_CVXBooster *CVXBoosterTransactorSession) WithdrawTo(_pid *big.Int, _amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _CVXBooster.Contract.WithdrawTo(&_CVXBooster.TransactOpts, _pid, _amount, _to)
}

// CVXBoosterDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the CVXBooster contract.
type CVXBoosterDepositedIterator struct {
	Event *CVXBoosterDeposited // Event containing the contract specifics and raw log

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
func (it *CVXBoosterDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CVXBoosterDeposited)
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
		it.Event = new(CVXBoosterDeposited)
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
func (it *CVXBoosterDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CVXBoosterDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CVXBoosterDeposited represents a Deposited event raised by the CVXBooster contract.
type CVXBoosterDeposited struct {
	User   common.Address
	Poolid *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x73a19dd210f1a7f902193214c0ee91dd35ee5b4d920cba8d519eca65a7b488ca.
//
// Solidity: event Deposited(address indexed user, uint256 indexed poolid, uint256 amount)
func (_CVXBooster *CVXBoosterFilterer) FilterDeposited(opts *bind.FilterOpts, user []common.Address, poolid []*big.Int) (*CVXBoosterDepositedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolidRule []interface{}
	for _, poolidItem := range poolid {
		poolidRule = append(poolidRule, poolidItem)
	}

	logs, sub, err := _CVXBooster.contract.FilterLogs(opts, "Deposited", userRule, poolidRule)
	if err != nil {
		return nil, err
	}
	return &CVXBoosterDepositedIterator{contract: _CVXBooster.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x73a19dd210f1a7f902193214c0ee91dd35ee5b4d920cba8d519eca65a7b488ca.
//
// Solidity: event Deposited(address indexed user, uint256 indexed poolid, uint256 amount)
func (_CVXBooster *CVXBoosterFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *CVXBoosterDeposited, user []common.Address, poolid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolidRule []interface{}
	for _, poolidItem := range poolid {
		poolidRule = append(poolidRule, poolidItem)
	}

	logs, sub, err := _CVXBooster.contract.WatchLogs(opts, "Deposited", userRule, poolidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CVXBoosterDeposited)
				if err := _CVXBooster.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x73a19dd210f1a7f902193214c0ee91dd35ee5b4d920cba8d519eca65a7b488ca.
//
// Solidity: event Deposited(address indexed user, uint256 indexed poolid, uint256 amount)
func (_CVXBooster *CVXBoosterFilterer) ParseDeposited(log types.Log) (*CVXBoosterDeposited, error) {
	event := new(CVXBoosterDeposited)
	if err := _CVXBooster.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CVXBoosterWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the CVXBooster contract.
type CVXBoosterWithdrawnIterator struct {
	Event *CVXBoosterWithdrawn // Event containing the contract specifics and raw log

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
func (it *CVXBoosterWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CVXBoosterWithdrawn)
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
		it.Event = new(CVXBoosterWithdrawn)
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
func (it *CVXBoosterWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CVXBoosterWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CVXBoosterWithdrawn represents a Withdrawn event raised by the CVXBooster contract.
type CVXBoosterWithdrawn struct {
	User   common.Address
	Poolid *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: event Withdrawn(address indexed user, uint256 indexed poolid, uint256 amount)
func (_CVXBooster *CVXBoosterFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address, poolid []*big.Int) (*CVXBoosterWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolidRule []interface{}
	for _, poolidItem := range poolid {
		poolidRule = append(poolidRule, poolidItem)
	}

	logs, sub, err := _CVXBooster.contract.FilterLogs(opts, "Withdrawn", userRule, poolidRule)
	if err != nil {
		return nil, err
	}
	return &CVXBoosterWithdrawnIterator{contract: _CVXBooster.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: event Withdrawn(address indexed user, uint256 indexed poolid, uint256 amount)
func (_CVXBooster *CVXBoosterFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *CVXBoosterWithdrawn, user []common.Address, poolid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolidRule []interface{}
	for _, poolidItem := range poolid {
		poolidRule = append(poolidRule, poolidItem)
	}

	logs, sub, err := _CVXBooster.contract.WatchLogs(opts, "Withdrawn", userRule, poolidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CVXBoosterWithdrawn)
				if err := _CVXBooster.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: event Withdrawn(address indexed user, uint256 indexed poolid, uint256 amount)
func (_CVXBooster *CVXBoosterFilterer) ParseWithdrawn(log types.Log) (*CVXBoosterWithdrawn, error) {
	event := new(CVXBoosterWithdrawn)
	if err := _CVXBooster.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
