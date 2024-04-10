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

// YVeYFIGaugeMetaData contains all meta data concerning the YVeYFIGauge contract.
var YVeYFIGaugeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BoostedBalanceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Initialize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"RecipientUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentRewards\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lastUpdateTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"periodFinish\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"historicalRewards\",\"type\":\"uint256\"}],\"name\":\"RewardsAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardsQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Sweep\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transfered\",\"type\":\"uint256\"}],\"name\":\"TransferredPenalty\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardPerTokenStored\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lastUpdateTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userRewardPerTokenPaid\",\"type\":\"uint256\"}],\"name\":\"UpdatedRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BOOSTING_FACTOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BOOST_DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRECISION_FACTOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARD_TOKEN\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VEYFI\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VE_YFI_POOL\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"boostedBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"internalType\":\"contractIGaugeController\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assets\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assets\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"earned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReward\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"getReward\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"historicalRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_controller\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_accounts\",\"type\":\"address[]\"}],\"name\":\"kick\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTimeRewardApplicable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"nextBoostedBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodFinish\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assets\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"previewMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assets\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assets\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"recipients\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerTokenStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newController\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"setRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"sweep\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRewardPerTokenPaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_claim\",\"type\":\"bool\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_claim\",\"type\":\"bool\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// YVeYFIGaugeABI is the input ABI used to generate the binding from.
// Deprecated: Use YVeYFIGaugeMetaData.ABI instead.
var YVeYFIGaugeABI = YVeYFIGaugeMetaData.ABI

// YVeYFIGauge is an auto generated Go binding around an Ethereum contract.
type YVeYFIGauge struct {
	YVeYFIGaugeCaller     // Read-only binding to the contract
	YVeYFIGaugeTransactor // Write-only binding to the contract
	YVeYFIGaugeFilterer   // Log filterer for contract events
}

// YVeYFIGaugeCaller is an auto generated read-only Go binding around an Ethereum contract.
type YVeYFIGaugeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YVeYFIGaugeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YVeYFIGaugeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YVeYFIGaugeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YVeYFIGaugeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YVeYFIGaugeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YVeYFIGaugeSession struct {
	Contract     *YVeYFIGauge      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YVeYFIGaugeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YVeYFIGaugeCallerSession struct {
	Contract *YVeYFIGaugeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// YVeYFIGaugeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YVeYFIGaugeTransactorSession struct {
	Contract     *YVeYFIGaugeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// YVeYFIGaugeRaw is an auto generated low-level Go binding around an Ethereum contract.
type YVeYFIGaugeRaw struct {
	Contract *YVeYFIGauge // Generic contract binding to access the raw methods on
}

// YVeYFIGaugeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YVeYFIGaugeCallerRaw struct {
	Contract *YVeYFIGaugeCaller // Generic read-only contract binding to access the raw methods on
}

// YVeYFIGaugeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YVeYFIGaugeTransactorRaw struct {
	Contract *YVeYFIGaugeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYVeYFIGauge creates a new instance of YVeYFIGauge, bound to a specific deployed contract.
func NewYVeYFIGauge(address common.Address, backend bind.ContractBackend) (*YVeYFIGauge, error) {
	contract, err := bindYVeYFIGauge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YVeYFIGauge{YVeYFIGaugeCaller: YVeYFIGaugeCaller{contract: contract}, YVeYFIGaugeTransactor: YVeYFIGaugeTransactor{contract: contract}, YVeYFIGaugeFilterer: YVeYFIGaugeFilterer{contract: contract}}, nil
}

// NewYVeYFIGaugeCaller creates a new read-only instance of YVeYFIGauge, bound to a specific deployed contract.
func NewYVeYFIGaugeCaller(address common.Address, caller bind.ContractCaller) (*YVeYFIGaugeCaller, error) {
	contract, err := bindYVeYFIGauge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YVeYFIGaugeCaller{contract: contract}, nil
}

// NewYVeYFIGaugeTransactor creates a new write-only instance of YVeYFIGauge, bound to a specific deployed contract.
func NewYVeYFIGaugeTransactor(address common.Address, transactor bind.ContractTransactor) (*YVeYFIGaugeTransactor, error) {
	contract, err := bindYVeYFIGauge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YVeYFIGaugeTransactor{contract: contract}, nil
}

// NewYVeYFIGaugeFilterer creates a new log filterer instance of YVeYFIGauge, bound to a specific deployed contract.
func NewYVeYFIGaugeFilterer(address common.Address, filterer bind.ContractFilterer) (*YVeYFIGaugeFilterer, error) {
	contract, err := bindYVeYFIGauge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YVeYFIGaugeFilterer{contract: contract}, nil
}

// bindYVeYFIGauge binds a generic wrapper to an already deployed contract.
func bindYVeYFIGauge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := YVeYFIGaugeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YVeYFIGauge *YVeYFIGaugeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YVeYFIGauge.Contract.YVeYFIGaugeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YVeYFIGauge *YVeYFIGaugeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.YVeYFIGaugeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YVeYFIGauge *YVeYFIGaugeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.YVeYFIGaugeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YVeYFIGauge *YVeYFIGaugeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YVeYFIGauge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YVeYFIGauge *YVeYFIGaugeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YVeYFIGauge *YVeYFIGaugeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.contract.Transact(opts, method, params...)
}

// BOOSTINGFACTOR is a free data retrieval call binding the contract method 0x980091fc.
//
// Solidity: function BOOSTING_FACTOR() view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCaller) BOOSTINGFACTOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YVeYFIGauge.contract.Call(opts, &out, "BOOSTING_FACTOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BOOSTINGFACTOR is a free data retrieval call binding the contract method 0x980091fc.
//
// Solidity: function BOOSTING_FACTOR() view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeSession) BOOSTINGFACTOR() (*big.Int, error) {
	return _YVeYFIGauge.Contract.BOOSTINGFACTOR(&_YVeYFIGauge.CallOpts)
}

// BOOSTINGFACTOR is a free data retrieval call binding the contract method 0x980091fc.
//
// Solidity: function BOOSTING_FACTOR() view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCallerSession) BOOSTINGFACTOR() (*big.Int, error) {
	return _YVeYFIGauge.Contract.BOOSTINGFACTOR(&_YVeYFIGauge.CallOpts)
}

// BOOSTDENOMINATOR is a free data retrieval call binding the contract method 0x1bd32ed3.
//
// Solidity: function BOOST_DENOMINATOR() view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCaller) BOOSTDENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YVeYFIGauge.contract.Call(opts, &out, "BOOST_DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BOOSTDENOMINATOR is a free data retrieval call binding the contract method 0x1bd32ed3.
//
// Solidity: function BOOST_DENOMINATOR() view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeSession) BOOSTDENOMINATOR() (*big.Int, error) {
	return _YVeYFIGauge.Contract.BOOSTDENOMINATOR(&_YVeYFIGauge.CallOpts)
}

// BOOSTDENOMINATOR is a free data retrieval call binding the contract method 0x1bd32ed3.
//
// Solidity: function BOOST_DENOMINATOR() view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCallerSession) BOOSTDENOMINATOR() (*big.Int, error) {
	return _YVeYFIGauge.Contract.BOOSTDENOMINATOR(&_YVeYFIGauge.CallOpts)
}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCaller) PRECISIONFACTOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YVeYFIGauge.contract.Call(opts, &out, "PRECISION_FACTOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeSession) PRECISIONFACTOR() (*big.Int, error) {
	return _YVeYFIGauge.Contract.PRECISIONFACTOR(&_YVeYFIGauge.CallOpts)
}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCallerSession) PRECISIONFACTOR() (*big.Int, error) {
	return _YVeYFIGauge.Contract.PRECISIONFACTOR(&_YVeYFIGauge.CallOpts)
}

// REWARDTOKEN is a free data retrieval call binding the contract method 0x99248ea7.
//
// Solidity: function REWARD_TOKEN() view returns(address)
func (_YVeYFIGauge *YVeYFIGaugeCaller) REWARDTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YVeYFIGauge.contract.Call(opts, &out, "REWARD_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// REWARDTOKEN is a free data retrieval call binding the contract method 0x99248ea7.
//
// Solidity: function REWARD_TOKEN() view returns(address)
func (_YVeYFIGauge *YVeYFIGaugeSession) REWARDTOKEN() (common.Address, error) {
	return _YVeYFIGauge.Contract.REWARDTOKEN(&_YVeYFIGauge.CallOpts)
}

// REWARDTOKEN is a free data retrieval call binding the contract method 0x99248ea7.
//
// Solidity: function REWARD_TOKEN() view returns(address)
func (_YVeYFIGauge *YVeYFIGaugeCallerSession) REWARDTOKEN() (common.Address, error) {
	return _YVeYFIGauge.Contract.REWARDTOKEN(&_YVeYFIGauge.CallOpts)
}

// VEYFI is a free data retrieval call binding the contract method 0x7d2f791d.
//
// Solidity: function VEYFI() view returns(address)
func (_YVeYFIGauge *YVeYFIGaugeCaller) VEYFI(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YVeYFIGauge.contract.Call(opts, &out, "VEYFI")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VEYFI is a free data retrieval call binding the contract method 0x7d2f791d.
//
// Solidity: function VEYFI() view returns(address)
func (_YVeYFIGauge *YVeYFIGaugeSession) VEYFI() (common.Address, error) {
	return _YVeYFIGauge.Contract.VEYFI(&_YVeYFIGauge.CallOpts)
}

// VEYFI is a free data retrieval call binding the contract method 0x7d2f791d.
//
// Solidity: function VEYFI() view returns(address)
func (_YVeYFIGauge *YVeYFIGaugeCallerSession) VEYFI() (common.Address, error) {
	return _YVeYFIGauge.Contract.VEYFI(&_YVeYFIGauge.CallOpts)
}

// VEYFIPOOL is a free data retrieval call binding the contract method 0xb5387c78.
//
// Solidity: function VE_YFI_POOL() view returns(address)
func (_YVeYFIGauge *YVeYFIGaugeCaller) VEYFIPOOL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YVeYFIGauge.contract.Call(opts, &out, "VE_YFI_POOL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VEYFIPOOL is a free data retrieval call binding the contract method 0xb5387c78.
//
// Solidity: function VE_YFI_POOL() view returns(address)
func (_YVeYFIGauge *YVeYFIGaugeSession) VEYFIPOOL() (common.Address, error) {
	return _YVeYFIGauge.Contract.VEYFIPOOL(&_YVeYFIGauge.CallOpts)
}

// VEYFIPOOL is a free data retrieval call binding the contract method 0xb5387c78.
//
// Solidity: function VE_YFI_POOL() view returns(address)
func (_YVeYFIGauge *YVeYFIGaugeCallerSession) VEYFIPOOL() (common.Address, error) {
	return _YVeYFIGauge.Contract.VEYFIPOOL(&_YVeYFIGauge.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YVeYFIGauge.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _YVeYFIGauge.Contract.Allowance(&_YVeYFIGauge.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _YVeYFIGauge.Contract.Allowance(&_YVeYFIGauge.CallOpts, owner, spender)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_YVeYFIGauge *YVeYFIGaugeCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YVeYFIGauge.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_YVeYFIGauge *YVeYFIGaugeSession) Asset() (common.Address, error) {
	return _YVeYFIGauge.Contract.Asset(&_YVeYFIGauge.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_YVeYFIGauge *YVeYFIGaugeCallerSession) Asset() (common.Address, error) {
	return _YVeYFIGauge.Contract.Asset(&_YVeYFIGauge.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YVeYFIGauge.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _YVeYFIGauge.Contract.BalanceOf(&_YVeYFIGauge.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _YVeYFIGauge.Contract.BalanceOf(&_YVeYFIGauge.CallOpts, account)
}

// BoostedBalanceOf is a free data retrieval call binding the contract method 0x1beabcd2.
//
// Solidity: function boostedBalanceOf(address _account) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCaller) BoostedBalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YVeYFIGauge.contract.Call(opts, &out, "boostedBalanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BoostedBalanceOf is a free data retrieval call binding the contract method 0x1beabcd2.
//
// Solidity: function boostedBalanceOf(address _account) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeSession) BoostedBalanceOf(_account common.Address) (*big.Int, error) {
	return _YVeYFIGauge.Contract.BoostedBalanceOf(&_YVeYFIGauge.CallOpts, _account)
}

// BoostedBalanceOf is a free data retrieval call binding the contract method 0x1beabcd2.
//
// Solidity: function boostedBalanceOf(address _account) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCallerSession) BoostedBalanceOf(_account common.Address) (*big.Int, error) {
	return _YVeYFIGauge.Contract.BoostedBalanceOf(&_YVeYFIGauge.CallOpts, _account)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_YVeYFIGauge *YVeYFIGaugeCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YVeYFIGauge.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_YVeYFIGauge *YVeYFIGaugeSession) Controller() (common.Address, error) {
	return _YVeYFIGauge.Contract.Controller(&_YVeYFIGauge.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_YVeYFIGauge *YVeYFIGaugeCallerSession) Controller() (common.Address, error) {
	return _YVeYFIGauge.Contract.Controller(&_YVeYFIGauge.CallOpts)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 _shares) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCaller) ConvertToAssets(opts *bind.CallOpts, _shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YVeYFIGauge.contract.Call(opts, &out, "convertToAssets", _shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 _shares) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeSession) ConvertToAssets(_shares *big.Int) (*big.Int, error) {
	return _YVeYFIGauge.Contract.ConvertToAssets(&_YVeYFIGauge.CallOpts, _shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 _shares) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCallerSession) ConvertToAssets(_shares *big.Int) (*big.Int, error) {
	return _YVeYFIGauge.Contract.ConvertToAssets(&_YVeYFIGauge.CallOpts, _shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 _assets) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCaller) ConvertToShares(opts *bind.CallOpts, _assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YVeYFIGauge.contract.Call(opts, &out, "convertToShares", _assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 _assets) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeSession) ConvertToShares(_assets *big.Int) (*big.Int, error) {
	return _YVeYFIGauge.Contract.ConvertToShares(&_YVeYFIGauge.CallOpts, _assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 _assets) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCallerSession) ConvertToShares(_assets *big.Int) (*big.Int, error) {
	return _YVeYFIGauge.Contract.ConvertToShares(&_YVeYFIGauge.CallOpts, _assets)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_YVeYFIGauge *YVeYFIGaugeCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _YVeYFIGauge.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_YVeYFIGauge *YVeYFIGaugeSession) Decimals() (uint8, error) {
	return _YVeYFIGauge.Contract.Decimals(&_YVeYFIGauge.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_YVeYFIGauge *YVeYFIGaugeCallerSession) Decimals() (uint8, error) {
	return _YVeYFIGauge.Contract.Decimals(&_YVeYFIGauge.CallOpts)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address _account) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCaller) Earned(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YVeYFIGauge.contract.Call(opts, &out, "earned", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address _account) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeSession) Earned(_account common.Address) (*big.Int, error) {
	return _YVeYFIGauge.Contract.Earned(&_YVeYFIGauge.CallOpts, _account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address _account) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCallerSession) Earned(_account common.Address) (*big.Int, error) {
	return _YVeYFIGauge.Contract.Earned(&_YVeYFIGauge.CallOpts, _account)
}

// HistoricalRewards is a free data retrieval call binding the contract method 0x262d3d6d.
//
// Solidity: function historicalRewards() view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCaller) HistoricalRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YVeYFIGauge.contract.Call(opts, &out, "historicalRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HistoricalRewards is a free data retrieval call binding the contract method 0x262d3d6d.
//
// Solidity: function historicalRewards() view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeSession) HistoricalRewards() (*big.Int, error) {
	return _YVeYFIGauge.Contract.HistoricalRewards(&_YVeYFIGauge.CallOpts)
}

// HistoricalRewards is a free data retrieval call binding the contract method 0x262d3d6d.
//
// Solidity: function historicalRewards() view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCallerSession) HistoricalRewards() (*big.Int, error) {
	return _YVeYFIGauge.Contract.HistoricalRewards(&_YVeYFIGauge.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCaller) LastTimeRewardApplicable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YVeYFIGauge.contract.Call(opts, &out, "lastTimeRewardApplicable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeSession) LastTimeRewardApplicable() (*big.Int, error) {
	return _YVeYFIGauge.Contract.LastTimeRewardApplicable(&_YVeYFIGauge.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCallerSession) LastTimeRewardApplicable() (*big.Int, error) {
	return _YVeYFIGauge.Contract.LastTimeRewardApplicable(&_YVeYFIGauge.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCaller) LastUpdateTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YVeYFIGauge.contract.Call(opts, &out, "lastUpdateTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeSession) LastUpdateTime() (*big.Int, error) {
	return _YVeYFIGauge.Contract.LastUpdateTime(&_YVeYFIGauge.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCallerSession) LastUpdateTime() (*big.Int, error) {
	return _YVeYFIGauge.Contract.LastUpdateTime(&_YVeYFIGauge.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCaller) MaxDeposit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YVeYFIGauge.contract.Call(opts, &out, "maxDeposit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _YVeYFIGauge.Contract.MaxDeposit(&_YVeYFIGauge.CallOpts, arg0)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCallerSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _YVeYFIGauge.Contract.MaxDeposit(&_YVeYFIGauge.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCaller) MaxMint(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YVeYFIGauge.contract.Call(opts, &out, "maxMint", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _YVeYFIGauge.Contract.MaxMint(&_YVeYFIGauge.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCallerSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _YVeYFIGauge.Contract.MaxMint(&_YVeYFIGauge.CallOpts, arg0)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address _owner) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCaller) MaxRedeem(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YVeYFIGauge.contract.Call(opts, &out, "maxRedeem", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address _owner) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeSession) MaxRedeem(_owner common.Address) (*big.Int, error) {
	return _YVeYFIGauge.Contract.MaxRedeem(&_YVeYFIGauge.CallOpts, _owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address _owner) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCallerSession) MaxRedeem(_owner common.Address) (*big.Int, error) {
	return _YVeYFIGauge.Contract.MaxRedeem(&_YVeYFIGauge.CallOpts, _owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address _owner) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCaller) MaxWithdraw(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YVeYFIGauge.contract.Call(opts, &out, "maxWithdraw", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address _owner) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeSession) MaxWithdraw(_owner common.Address) (*big.Int, error) {
	return _YVeYFIGauge.Contract.MaxWithdraw(&_YVeYFIGauge.CallOpts, _owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address _owner) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCallerSession) MaxWithdraw(_owner common.Address) (*big.Int, error) {
	return _YVeYFIGauge.Contract.MaxWithdraw(&_YVeYFIGauge.CallOpts, _owner)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YVeYFIGauge *YVeYFIGaugeCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _YVeYFIGauge.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YVeYFIGauge *YVeYFIGaugeSession) Name() (string, error) {
	return _YVeYFIGauge.Contract.Name(&_YVeYFIGauge.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YVeYFIGauge *YVeYFIGaugeCallerSession) Name() (string, error) {
	return _YVeYFIGauge.Contract.Name(&_YVeYFIGauge.CallOpts)
}

// NextBoostedBalanceOf is a free data retrieval call binding the contract method 0xc67ffb4e.
//
// Solidity: function nextBoostedBalanceOf(address _account) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCaller) NextBoostedBalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YVeYFIGauge.contract.Call(opts, &out, "nextBoostedBalanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextBoostedBalanceOf is a free data retrieval call binding the contract method 0xc67ffb4e.
//
// Solidity: function nextBoostedBalanceOf(address _account) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeSession) NextBoostedBalanceOf(_account common.Address) (*big.Int, error) {
	return _YVeYFIGauge.Contract.NextBoostedBalanceOf(&_YVeYFIGauge.CallOpts, _account)
}

// NextBoostedBalanceOf is a free data retrieval call binding the contract method 0xc67ffb4e.
//
// Solidity: function nextBoostedBalanceOf(address _account) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCallerSession) NextBoostedBalanceOf(_account common.Address) (*big.Int, error) {
	return _YVeYFIGauge.Contract.NextBoostedBalanceOf(&_YVeYFIGauge.CallOpts, _account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_YVeYFIGauge *YVeYFIGaugeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YVeYFIGauge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_YVeYFIGauge *YVeYFIGaugeSession) Owner() (common.Address, error) {
	return _YVeYFIGauge.Contract.Owner(&_YVeYFIGauge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_YVeYFIGauge *YVeYFIGaugeCallerSession) Owner() (common.Address, error) {
	return _YVeYFIGauge.Contract.Owner(&_YVeYFIGauge.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCaller) PeriodFinish(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YVeYFIGauge.contract.Call(opts, &out, "periodFinish")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeSession) PeriodFinish() (*big.Int, error) {
	return _YVeYFIGauge.Contract.PeriodFinish(&_YVeYFIGauge.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCallerSession) PeriodFinish() (*big.Int, error) {
	return _YVeYFIGauge.Contract.PeriodFinish(&_YVeYFIGauge.CallOpts)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 _assets) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCaller) PreviewDeposit(opts *bind.CallOpts, _assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YVeYFIGauge.contract.Call(opts, &out, "previewDeposit", _assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 _assets) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeSession) PreviewDeposit(_assets *big.Int) (*big.Int, error) {
	return _YVeYFIGauge.Contract.PreviewDeposit(&_YVeYFIGauge.CallOpts, _assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 _assets) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCallerSession) PreviewDeposit(_assets *big.Int) (*big.Int, error) {
	return _YVeYFIGauge.Contract.PreviewDeposit(&_YVeYFIGauge.CallOpts, _assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 _shares) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCaller) PreviewMint(opts *bind.CallOpts, _shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YVeYFIGauge.contract.Call(opts, &out, "previewMint", _shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 _shares) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeSession) PreviewMint(_shares *big.Int) (*big.Int, error) {
	return _YVeYFIGauge.Contract.PreviewMint(&_YVeYFIGauge.CallOpts, _shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 _shares) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCallerSession) PreviewMint(_shares *big.Int) (*big.Int, error) {
	return _YVeYFIGauge.Contract.PreviewMint(&_YVeYFIGauge.CallOpts, _shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 _assets) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCaller) PreviewRedeem(opts *bind.CallOpts, _assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YVeYFIGauge.contract.Call(opts, &out, "previewRedeem", _assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 _assets) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeSession) PreviewRedeem(_assets *big.Int) (*big.Int, error) {
	return _YVeYFIGauge.Contract.PreviewRedeem(&_YVeYFIGauge.CallOpts, _assets)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 _assets) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCallerSession) PreviewRedeem(_assets *big.Int) (*big.Int, error) {
	return _YVeYFIGauge.Contract.PreviewRedeem(&_YVeYFIGauge.CallOpts, _assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 _assets) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCaller) PreviewWithdraw(opts *bind.CallOpts, _assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YVeYFIGauge.contract.Call(opts, &out, "previewWithdraw", _assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 _assets) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeSession) PreviewWithdraw(_assets *big.Int) (*big.Int, error) {
	return _YVeYFIGauge.Contract.PreviewWithdraw(&_YVeYFIGauge.CallOpts, _assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 _assets) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCallerSession) PreviewWithdraw(_assets *big.Int) (*big.Int, error) {
	return _YVeYFIGauge.Contract.PreviewWithdraw(&_YVeYFIGauge.CallOpts, _assets)
}

// Recipients is a free data retrieval call binding the contract method 0xeb820312.
//
// Solidity: function recipients(address ) view returns(address)
func (_YVeYFIGauge *YVeYFIGaugeCaller) Recipients(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _YVeYFIGauge.contract.Call(opts, &out, "recipients", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Recipients is a free data retrieval call binding the contract method 0xeb820312.
//
// Solidity: function recipients(address ) view returns(address)
func (_YVeYFIGauge *YVeYFIGaugeSession) Recipients(arg0 common.Address) (common.Address, error) {
	return _YVeYFIGauge.Contract.Recipients(&_YVeYFIGauge.CallOpts, arg0)
}

// Recipients is a free data retrieval call binding the contract method 0xeb820312.
//
// Solidity: function recipients(address ) view returns(address)
func (_YVeYFIGauge *YVeYFIGaugeCallerSession) Recipients(arg0 common.Address) (common.Address, error) {
	return _YVeYFIGauge.Contract.Recipients(&_YVeYFIGauge.CallOpts, arg0)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCaller) RewardPerToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YVeYFIGauge.contract.Call(opts, &out, "rewardPerToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeSession) RewardPerToken() (*big.Int, error) {
	return _YVeYFIGauge.Contract.RewardPerToken(&_YVeYFIGauge.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCallerSession) RewardPerToken() (*big.Int, error) {
	return _YVeYFIGauge.Contract.RewardPerToken(&_YVeYFIGauge.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCaller) RewardPerTokenStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YVeYFIGauge.contract.Call(opts, &out, "rewardPerTokenStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeSession) RewardPerTokenStored() (*big.Int, error) {
	return _YVeYFIGauge.Contract.RewardPerTokenStored(&_YVeYFIGauge.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCallerSession) RewardPerTokenStored() (*big.Int, error) {
	return _YVeYFIGauge.Contract.RewardPerTokenStored(&_YVeYFIGauge.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCaller) RewardRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YVeYFIGauge.contract.Call(opts, &out, "rewardRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeSession) RewardRate() (*big.Int, error) {
	return _YVeYFIGauge.Contract.RewardRate(&_YVeYFIGauge.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCallerSession) RewardRate() (*big.Int, error) {
	return _YVeYFIGauge.Contract.RewardRate(&_YVeYFIGauge.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCaller) Rewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YVeYFIGauge.contract.Call(opts, &out, "rewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _YVeYFIGauge.Contract.Rewards(&_YVeYFIGauge.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCallerSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _YVeYFIGauge.Contract.Rewards(&_YVeYFIGauge.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_YVeYFIGauge *YVeYFIGaugeCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _YVeYFIGauge.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_YVeYFIGauge *YVeYFIGaugeSession) Symbol() (string, error) {
	return _YVeYFIGauge.Contract.Symbol(&_YVeYFIGauge.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_YVeYFIGauge *YVeYFIGaugeCallerSession) Symbol() (string, error) {
	return _YVeYFIGauge.Contract.Symbol(&_YVeYFIGauge.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YVeYFIGauge.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeSession) TotalAssets() (*big.Int, error) {
	return _YVeYFIGauge.Contract.TotalAssets(&_YVeYFIGauge.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCallerSession) TotalAssets() (*big.Int, error) {
	return _YVeYFIGauge.Contract.TotalAssets(&_YVeYFIGauge.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YVeYFIGauge.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeSession) TotalSupply() (*big.Int, error) {
	return _YVeYFIGauge.Contract.TotalSupply(&_YVeYFIGauge.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCallerSession) TotalSupply() (*big.Int, error) {
	return _YVeYFIGauge.Contract.TotalSupply(&_YVeYFIGauge.CallOpts)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCaller) UserRewardPerTokenPaid(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YVeYFIGauge.contract.Call(opts, &out, "userRewardPerTokenPaid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeSession) UserRewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _YVeYFIGauge.Contract.UserRewardPerTokenPaid(&_YVeYFIGauge.CallOpts, arg0)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeCallerSession) UserRewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _YVeYFIGauge.Contract.UserRewardPerTokenPaid(&_YVeYFIGauge.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_YVeYFIGauge *YVeYFIGaugeTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YVeYFIGauge.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_YVeYFIGauge *YVeYFIGaugeSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.Approve(&_YVeYFIGauge.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_YVeYFIGauge *YVeYFIGaugeTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.Approve(&_YVeYFIGauge.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_YVeYFIGauge *YVeYFIGaugeTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _YVeYFIGauge.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_YVeYFIGauge *YVeYFIGaugeSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.DecreaseAllowance(&_YVeYFIGauge.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_YVeYFIGauge *YVeYFIGaugeTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.DecreaseAllowance(&_YVeYFIGauge.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _assets, address _receiver) returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeTransactor) Deposit(opts *bind.TransactOpts, _assets *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _YVeYFIGauge.contract.Transact(opts, "deposit", _assets, _receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _assets, address _receiver) returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeSession) Deposit(_assets *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.Deposit(&_YVeYFIGauge.TransactOpts, _assets, _receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _assets, address _receiver) returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeTransactorSession) Deposit(_assets *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.Deposit(&_YVeYFIGauge.TransactOpts, _assets, _receiver)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _assets) returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeTransactor) Deposit0(opts *bind.TransactOpts, _assets *big.Int) (*types.Transaction, error) {
	return _YVeYFIGauge.contract.Transact(opts, "deposit0", _assets)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _assets) returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeSession) Deposit0(_assets *big.Int) (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.Deposit0(&_YVeYFIGauge.TransactOpts, _assets)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _assets) returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeTransactorSession) Deposit0(_assets *big.Int) (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.Deposit0(&_YVeYFIGauge.TransactOpts, _assets)
}

// Deposit1 is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeTransactor) Deposit1(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YVeYFIGauge.contract.Transact(opts, "deposit1")
}

// Deposit1 is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeSession) Deposit1() (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.Deposit1(&_YVeYFIGauge.TransactOpts)
}

// Deposit1 is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeTransactorSession) Deposit1() (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.Deposit1(&_YVeYFIGauge.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns(bool)
func (_YVeYFIGauge *YVeYFIGaugeTransactor) GetReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YVeYFIGauge.contract.Transact(opts, "getReward")
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns(bool)
func (_YVeYFIGauge *YVeYFIGaugeSession) GetReward() (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.GetReward(&_YVeYFIGauge.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns(bool)
func (_YVeYFIGauge *YVeYFIGaugeTransactorSession) GetReward() (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.GetReward(&_YVeYFIGauge.TransactOpts)
}

// GetReward0 is a paid mutator transaction binding the contract method 0xc00007b0.
//
// Solidity: function getReward(address _account) returns(bool)
func (_YVeYFIGauge *YVeYFIGaugeTransactor) GetReward0(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _YVeYFIGauge.contract.Transact(opts, "getReward0", _account)
}

// GetReward0 is a paid mutator transaction binding the contract method 0xc00007b0.
//
// Solidity: function getReward(address _account) returns(bool)
func (_YVeYFIGauge *YVeYFIGaugeSession) GetReward0(_account common.Address) (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.GetReward0(&_YVeYFIGauge.TransactOpts, _account)
}

// GetReward0 is a paid mutator transaction binding the contract method 0xc00007b0.
//
// Solidity: function getReward(address _account) returns(bool)
func (_YVeYFIGauge *YVeYFIGaugeTransactorSession) GetReward0(_account common.Address) (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.GetReward0(&_YVeYFIGauge.TransactOpts, _account)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_YVeYFIGauge *YVeYFIGaugeTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _YVeYFIGauge.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_YVeYFIGauge *YVeYFIGaugeSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.IncreaseAllowance(&_YVeYFIGauge.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_YVeYFIGauge *YVeYFIGaugeTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.IncreaseAllowance(&_YVeYFIGauge.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x246581f7.
//
// Solidity: function initialize(address _asset, address _owner, address _controller, bytes _data) returns()
func (_YVeYFIGauge *YVeYFIGaugeTransactor) Initialize(opts *bind.TransactOpts, _asset common.Address, _owner common.Address, _controller common.Address, _data []byte) (*types.Transaction, error) {
	return _YVeYFIGauge.contract.Transact(opts, "initialize", _asset, _owner, _controller, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0x246581f7.
//
// Solidity: function initialize(address _asset, address _owner, address _controller, bytes _data) returns()
func (_YVeYFIGauge *YVeYFIGaugeSession) Initialize(_asset common.Address, _owner common.Address, _controller common.Address, _data []byte) (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.Initialize(&_YVeYFIGauge.TransactOpts, _asset, _owner, _controller, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0x246581f7.
//
// Solidity: function initialize(address _asset, address _owner, address _controller, bytes _data) returns()
func (_YVeYFIGauge *YVeYFIGaugeTransactorSession) Initialize(_asset common.Address, _owner common.Address, _controller common.Address, _data []byte) (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.Initialize(&_YVeYFIGauge.TransactOpts, _asset, _owner, _controller, _data)
}

// Kick is a paid mutator transaction binding the contract method 0x1530e6d8.
//
// Solidity: function kick(address[] _accounts) returns()
func (_YVeYFIGauge *YVeYFIGaugeTransactor) Kick(opts *bind.TransactOpts, _accounts []common.Address) (*types.Transaction, error) {
	return _YVeYFIGauge.contract.Transact(opts, "kick", _accounts)
}

// Kick is a paid mutator transaction binding the contract method 0x1530e6d8.
//
// Solidity: function kick(address[] _accounts) returns()
func (_YVeYFIGauge *YVeYFIGaugeSession) Kick(_accounts []common.Address) (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.Kick(&_YVeYFIGauge.TransactOpts, _accounts)
}

// Kick is a paid mutator transaction binding the contract method 0x1530e6d8.
//
// Solidity: function kick(address[] _accounts) returns()
func (_YVeYFIGauge *YVeYFIGaugeTransactorSession) Kick(_accounts []common.Address) (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.Kick(&_YVeYFIGauge.TransactOpts, _accounts)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 _shares, address _receiver) returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeTransactor) Mint(opts *bind.TransactOpts, _shares *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _YVeYFIGauge.contract.Transact(opts, "mint", _shares, _receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 _shares, address _receiver) returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeSession) Mint(_shares *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.Mint(&_YVeYFIGauge.TransactOpts, _shares, _receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 _shares, address _receiver) returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeTransactorSession) Mint(_shares *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.Mint(&_YVeYFIGauge.TransactOpts, _shares, _receiver)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 _assets, address _receiver, address _owner) returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeTransactor) Redeem(opts *bind.TransactOpts, _assets *big.Int, _receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _YVeYFIGauge.contract.Transact(opts, "redeem", _assets, _receiver, _owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 _assets, address _receiver, address _owner) returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeSession) Redeem(_assets *big.Int, _receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.Redeem(&_YVeYFIGauge.TransactOpts, _assets, _receiver, _owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 _assets, address _receiver, address _owner) returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeTransactorSession) Redeem(_assets *big.Int, _receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.Redeem(&_YVeYFIGauge.TransactOpts, _assets, _receiver, _owner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_YVeYFIGauge *YVeYFIGaugeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YVeYFIGauge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_YVeYFIGauge *YVeYFIGaugeSession) RenounceOwnership() (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.RenounceOwnership(&_YVeYFIGauge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_YVeYFIGauge *YVeYFIGaugeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.RenounceOwnership(&_YVeYFIGauge.TransactOpts)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _newController) returns()
func (_YVeYFIGauge *YVeYFIGaugeTransactor) SetController(opts *bind.TransactOpts, _newController common.Address) (*types.Transaction, error) {
	return _YVeYFIGauge.contract.Transact(opts, "setController", _newController)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _newController) returns()
func (_YVeYFIGauge *YVeYFIGaugeSession) SetController(_newController common.Address) (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.SetController(&_YVeYFIGauge.TransactOpts, _newController)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _newController) returns()
func (_YVeYFIGauge *YVeYFIGaugeTransactorSession) SetController(_newController common.Address) (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.SetController(&_YVeYFIGauge.TransactOpts, _newController)
}

// SetRecipient is a paid mutator transaction binding the contract method 0x3bbed4a0.
//
// Solidity: function setRecipient(address _recipient) returns()
func (_YVeYFIGauge *YVeYFIGaugeTransactor) SetRecipient(opts *bind.TransactOpts, _recipient common.Address) (*types.Transaction, error) {
	return _YVeYFIGauge.contract.Transact(opts, "setRecipient", _recipient)
}

// SetRecipient is a paid mutator transaction binding the contract method 0x3bbed4a0.
//
// Solidity: function setRecipient(address _recipient) returns()
func (_YVeYFIGauge *YVeYFIGaugeSession) SetRecipient(_recipient common.Address) (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.SetRecipient(&_YVeYFIGauge.TransactOpts, _recipient)
}

// SetRecipient is a paid mutator transaction binding the contract method 0x3bbed4a0.
//
// Solidity: function setRecipient(address _recipient) returns()
func (_YVeYFIGauge *YVeYFIGaugeTransactorSession) SetRecipient(_recipient common.Address) (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.SetRecipient(&_YVeYFIGauge.TransactOpts, _recipient)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address _token) returns(bool)
func (_YVeYFIGauge *YVeYFIGaugeTransactor) Sweep(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _YVeYFIGauge.contract.Transact(opts, "sweep", _token)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address _token) returns(bool)
func (_YVeYFIGauge *YVeYFIGaugeSession) Sweep(_token common.Address) (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.Sweep(&_YVeYFIGauge.TransactOpts, _token)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address _token) returns(bool)
func (_YVeYFIGauge *YVeYFIGaugeTransactorSession) Sweep(_token common.Address) (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.Sweep(&_YVeYFIGauge.TransactOpts, _token)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_YVeYFIGauge *YVeYFIGaugeTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YVeYFIGauge.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_YVeYFIGauge *YVeYFIGaugeSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.Transfer(&_YVeYFIGauge.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_YVeYFIGauge *YVeYFIGaugeTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.Transfer(&_YVeYFIGauge.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_YVeYFIGauge *YVeYFIGaugeTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YVeYFIGauge.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_YVeYFIGauge *YVeYFIGaugeSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.TransferFrom(&_YVeYFIGauge.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_YVeYFIGauge *YVeYFIGaugeTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.TransferFrom(&_YVeYFIGauge.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_YVeYFIGauge *YVeYFIGaugeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _YVeYFIGauge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_YVeYFIGauge *YVeYFIGaugeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.TransferOwnership(&_YVeYFIGauge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_YVeYFIGauge *YVeYFIGaugeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.TransferOwnership(&_YVeYFIGauge.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YVeYFIGauge.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeSession) Withdraw() (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.Withdraw(&_YVeYFIGauge.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeTransactorSession) Withdraw() (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.Withdraw(&_YVeYFIGauge.TransactOpts)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xa810a54c.
//
// Solidity: function withdraw(bool _claim) returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeTransactor) Withdraw0(opts *bind.TransactOpts, _claim bool) (*types.Transaction, error) {
	return _YVeYFIGauge.contract.Transact(opts, "withdraw0", _claim)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xa810a54c.
//
// Solidity: function withdraw(bool _claim) returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeSession) Withdraw0(_claim bool) (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.Withdraw0(&_YVeYFIGauge.TransactOpts, _claim)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xa810a54c.
//
// Solidity: function withdraw(bool _claim) returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeTransactorSession) Withdraw0(_claim bool) (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.Withdraw0(&_YVeYFIGauge.TransactOpts, _claim)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 _assets, address _receiver, address _owner) returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeTransactor) Withdraw1(opts *bind.TransactOpts, _assets *big.Int, _receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _YVeYFIGauge.contract.Transact(opts, "withdraw1", _assets, _receiver, _owner)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 _assets, address _receiver, address _owner) returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeSession) Withdraw1(_assets *big.Int, _receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.Withdraw1(&_YVeYFIGauge.TransactOpts, _assets, _receiver, _owner)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 _assets, address _receiver, address _owner) returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeTransactorSession) Withdraw1(_assets *big.Int, _receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.Withdraw1(&_YVeYFIGauge.TransactOpts, _assets, _receiver, _owner)
}

// Withdraw2 is a paid mutator transaction binding the contract method 0xd045f2c4.
//
// Solidity: function withdraw(uint256 _assets, address _receiver, address _owner, bool _claim) returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeTransactor) Withdraw2(opts *bind.TransactOpts, _assets *big.Int, _receiver common.Address, _owner common.Address, _claim bool) (*types.Transaction, error) {
	return _YVeYFIGauge.contract.Transact(opts, "withdraw2", _assets, _receiver, _owner, _claim)
}

// Withdraw2 is a paid mutator transaction binding the contract method 0xd045f2c4.
//
// Solidity: function withdraw(uint256 _assets, address _receiver, address _owner, bool _claim) returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeSession) Withdraw2(_assets *big.Int, _receiver common.Address, _owner common.Address, _claim bool) (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.Withdraw2(&_YVeYFIGauge.TransactOpts, _assets, _receiver, _owner, _claim)
}

// Withdraw2 is a paid mutator transaction binding the contract method 0xd045f2c4.
//
// Solidity: function withdraw(uint256 _assets, address _receiver, address _owner, bool _claim) returns(uint256)
func (_YVeYFIGauge *YVeYFIGaugeTransactorSession) Withdraw2(_assets *big.Int, _receiver common.Address, _owner common.Address, _claim bool) (*types.Transaction, error) {
	return _YVeYFIGauge.Contract.Withdraw2(&_YVeYFIGauge.TransactOpts, _assets, _receiver, _owner, _claim)
}

// YVeYFIGaugeApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the YVeYFIGauge contract.
type YVeYFIGaugeApprovalIterator struct {
	Event *YVeYFIGaugeApproval // Event containing the contract specifics and raw log

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
func (it *YVeYFIGaugeApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YVeYFIGaugeApproval)
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
		it.Event = new(YVeYFIGaugeApproval)
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
func (it *YVeYFIGaugeApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YVeYFIGaugeApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YVeYFIGaugeApproval represents a Approval event raised by the YVeYFIGauge contract.
type YVeYFIGaugeApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_YVeYFIGauge *YVeYFIGaugeFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*YVeYFIGaugeApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _YVeYFIGauge.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &YVeYFIGaugeApprovalIterator{contract: _YVeYFIGauge.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_YVeYFIGauge *YVeYFIGaugeFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *YVeYFIGaugeApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _YVeYFIGauge.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YVeYFIGaugeApproval)
				if err := _YVeYFIGauge.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_YVeYFIGauge *YVeYFIGaugeFilterer) ParseApproval(log types.Log) (*YVeYFIGaugeApproval, error) {
	event := new(YVeYFIGaugeApproval)
	if err := _YVeYFIGauge.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YVeYFIGaugeBoostedBalanceUpdatedIterator is returned from FilterBoostedBalanceUpdated and is used to iterate over the raw logs and unpacked data for BoostedBalanceUpdated events raised by the YVeYFIGauge contract.
type YVeYFIGaugeBoostedBalanceUpdatedIterator struct {
	Event *YVeYFIGaugeBoostedBalanceUpdated // Event containing the contract specifics and raw log

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
func (it *YVeYFIGaugeBoostedBalanceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YVeYFIGaugeBoostedBalanceUpdated)
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
		it.Event = new(YVeYFIGaugeBoostedBalanceUpdated)
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
func (it *YVeYFIGaugeBoostedBalanceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YVeYFIGaugeBoostedBalanceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YVeYFIGaugeBoostedBalanceUpdated represents a BoostedBalanceUpdated event raised by the YVeYFIGauge contract.
type YVeYFIGaugeBoostedBalanceUpdated struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBoostedBalanceUpdated is a free log retrieval operation binding the contract event 0x291ff844d30f85bb011aca3bfccedead238b6ed2e4b283504e3c2231d134524b.
//
// Solidity: event BoostedBalanceUpdated(address indexed account, uint256 amount)
func (_YVeYFIGauge *YVeYFIGaugeFilterer) FilterBoostedBalanceUpdated(opts *bind.FilterOpts, account []common.Address) (*YVeYFIGaugeBoostedBalanceUpdatedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _YVeYFIGauge.contract.FilterLogs(opts, "BoostedBalanceUpdated", accountRule)
	if err != nil {
		return nil, err
	}
	return &YVeYFIGaugeBoostedBalanceUpdatedIterator{contract: _YVeYFIGauge.contract, event: "BoostedBalanceUpdated", logs: logs, sub: sub}, nil
}

// WatchBoostedBalanceUpdated is a free log subscription operation binding the contract event 0x291ff844d30f85bb011aca3bfccedead238b6ed2e4b283504e3c2231d134524b.
//
// Solidity: event BoostedBalanceUpdated(address indexed account, uint256 amount)
func (_YVeYFIGauge *YVeYFIGaugeFilterer) WatchBoostedBalanceUpdated(opts *bind.WatchOpts, sink chan<- *YVeYFIGaugeBoostedBalanceUpdated, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _YVeYFIGauge.contract.WatchLogs(opts, "BoostedBalanceUpdated", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YVeYFIGaugeBoostedBalanceUpdated)
				if err := _YVeYFIGauge.contract.UnpackLog(event, "BoostedBalanceUpdated", log); err != nil {
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

// ParseBoostedBalanceUpdated is a log parse operation binding the contract event 0x291ff844d30f85bb011aca3bfccedead238b6ed2e4b283504e3c2231d134524b.
//
// Solidity: event BoostedBalanceUpdated(address indexed account, uint256 amount)
func (_YVeYFIGauge *YVeYFIGaugeFilterer) ParseBoostedBalanceUpdated(log types.Log) (*YVeYFIGaugeBoostedBalanceUpdated, error) {
	event := new(YVeYFIGaugeBoostedBalanceUpdated)
	if err := _YVeYFIGauge.contract.UnpackLog(event, "BoostedBalanceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YVeYFIGaugeDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the YVeYFIGauge contract.
type YVeYFIGaugeDepositIterator struct {
	Event *YVeYFIGaugeDeposit // Event containing the contract specifics and raw log

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
func (it *YVeYFIGaugeDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YVeYFIGaugeDeposit)
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
		it.Event = new(YVeYFIGaugeDeposit)
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
func (it *YVeYFIGaugeDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YVeYFIGaugeDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YVeYFIGaugeDeposit represents a Deposit event raised by the YVeYFIGauge contract.
type YVeYFIGaugeDeposit struct {
	Caller common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed caller, address indexed owner, uint256 assets, uint256 shares)
func (_YVeYFIGauge *YVeYFIGaugeFilterer) FilterDeposit(opts *bind.FilterOpts, caller []common.Address, owner []common.Address) (*YVeYFIGaugeDepositIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _YVeYFIGauge.contract.FilterLogs(opts, "Deposit", callerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &YVeYFIGaugeDepositIterator{contract: _YVeYFIGauge.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed caller, address indexed owner, uint256 assets, uint256 shares)
func (_YVeYFIGauge *YVeYFIGaugeFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *YVeYFIGaugeDeposit, caller []common.Address, owner []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _YVeYFIGauge.contract.WatchLogs(opts, "Deposit", callerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YVeYFIGaugeDeposit)
				if err := _YVeYFIGauge.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_YVeYFIGauge *YVeYFIGaugeFilterer) ParseDeposit(log types.Log) (*YVeYFIGaugeDeposit, error) {
	event := new(YVeYFIGaugeDeposit)
	if err := _YVeYFIGauge.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YVeYFIGaugeInitializeIterator is returned from FilterInitialize and is used to iterate over the raw logs and unpacked data for Initialize events raised by the YVeYFIGauge contract.
type YVeYFIGaugeInitializeIterator struct {
	Event *YVeYFIGaugeInitialize // Event containing the contract specifics and raw log

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
func (it *YVeYFIGaugeInitializeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YVeYFIGaugeInitialize)
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
		it.Event = new(YVeYFIGaugeInitialize)
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
func (it *YVeYFIGaugeInitializeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YVeYFIGaugeInitializeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YVeYFIGaugeInitialize represents a Initialize event raised by the YVeYFIGauge contract.
type YVeYFIGaugeInitialize struct {
	Asset common.Address
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterInitialize is a free log retrieval operation binding the contract event 0xdc90fed0326ba91706deeac7eb34ac9f8b680734f9d782864dc29704d23bed6a.
//
// Solidity: event Initialize(address indexed asset, address indexed owner)
func (_YVeYFIGauge *YVeYFIGaugeFilterer) FilterInitialize(opts *bind.FilterOpts, asset []common.Address, owner []common.Address) (*YVeYFIGaugeInitializeIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _YVeYFIGauge.contract.FilterLogs(opts, "Initialize", assetRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &YVeYFIGaugeInitializeIterator{contract: _YVeYFIGauge.contract, event: "Initialize", logs: logs, sub: sub}, nil
}

// WatchInitialize is a free log subscription operation binding the contract event 0xdc90fed0326ba91706deeac7eb34ac9f8b680734f9d782864dc29704d23bed6a.
//
// Solidity: event Initialize(address indexed asset, address indexed owner)
func (_YVeYFIGauge *YVeYFIGaugeFilterer) WatchInitialize(opts *bind.WatchOpts, sink chan<- *YVeYFIGaugeInitialize, asset []common.Address, owner []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _YVeYFIGauge.contract.WatchLogs(opts, "Initialize", assetRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YVeYFIGaugeInitialize)
				if err := _YVeYFIGauge.contract.UnpackLog(event, "Initialize", log); err != nil {
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

// ParseInitialize is a log parse operation binding the contract event 0xdc90fed0326ba91706deeac7eb34ac9f8b680734f9d782864dc29704d23bed6a.
//
// Solidity: event Initialize(address indexed asset, address indexed owner)
func (_YVeYFIGauge *YVeYFIGaugeFilterer) ParseInitialize(log types.Log) (*YVeYFIGaugeInitialize, error) {
	event := new(YVeYFIGaugeInitialize)
	if err := _YVeYFIGauge.contract.UnpackLog(event, "Initialize", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YVeYFIGaugeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the YVeYFIGauge contract.
type YVeYFIGaugeInitializedIterator struct {
	Event *YVeYFIGaugeInitialized // Event containing the contract specifics and raw log

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
func (it *YVeYFIGaugeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YVeYFIGaugeInitialized)
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
		it.Event = new(YVeYFIGaugeInitialized)
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
func (it *YVeYFIGaugeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YVeYFIGaugeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YVeYFIGaugeInitialized represents a Initialized event raised by the YVeYFIGauge contract.
type YVeYFIGaugeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_YVeYFIGauge *YVeYFIGaugeFilterer) FilterInitialized(opts *bind.FilterOpts) (*YVeYFIGaugeInitializedIterator, error) {

	logs, sub, err := _YVeYFIGauge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &YVeYFIGaugeInitializedIterator{contract: _YVeYFIGauge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_YVeYFIGauge *YVeYFIGaugeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *YVeYFIGaugeInitialized) (event.Subscription, error) {

	logs, sub, err := _YVeYFIGauge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YVeYFIGaugeInitialized)
				if err := _YVeYFIGauge.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_YVeYFIGauge *YVeYFIGaugeFilterer) ParseInitialized(log types.Log) (*YVeYFIGaugeInitialized, error) {
	event := new(YVeYFIGaugeInitialized)
	if err := _YVeYFIGauge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YVeYFIGaugeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the YVeYFIGauge contract.
type YVeYFIGaugeOwnershipTransferredIterator struct {
	Event *YVeYFIGaugeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *YVeYFIGaugeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YVeYFIGaugeOwnershipTransferred)
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
		it.Event = new(YVeYFIGaugeOwnershipTransferred)
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
func (it *YVeYFIGaugeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YVeYFIGaugeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YVeYFIGaugeOwnershipTransferred represents a OwnershipTransferred event raised by the YVeYFIGauge contract.
type YVeYFIGaugeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_YVeYFIGauge *YVeYFIGaugeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*YVeYFIGaugeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _YVeYFIGauge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &YVeYFIGaugeOwnershipTransferredIterator{contract: _YVeYFIGauge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_YVeYFIGauge *YVeYFIGaugeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *YVeYFIGaugeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _YVeYFIGauge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YVeYFIGaugeOwnershipTransferred)
				if err := _YVeYFIGauge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_YVeYFIGauge *YVeYFIGaugeFilterer) ParseOwnershipTransferred(log types.Log) (*YVeYFIGaugeOwnershipTransferred, error) {
	event := new(YVeYFIGaugeOwnershipTransferred)
	if err := _YVeYFIGauge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YVeYFIGaugeRecipientUpdatedIterator is returned from FilterRecipientUpdated and is used to iterate over the raw logs and unpacked data for RecipientUpdated events raised by the YVeYFIGauge contract.
type YVeYFIGaugeRecipientUpdatedIterator struct {
	Event *YVeYFIGaugeRecipientUpdated // Event containing the contract specifics and raw log

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
func (it *YVeYFIGaugeRecipientUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YVeYFIGaugeRecipientUpdated)
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
		it.Event = new(YVeYFIGaugeRecipientUpdated)
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
func (it *YVeYFIGaugeRecipientUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YVeYFIGaugeRecipientUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YVeYFIGaugeRecipientUpdated represents a RecipientUpdated event raised by the YVeYFIGauge contract.
type YVeYFIGaugeRecipientUpdated struct {
	Account   common.Address
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRecipientUpdated is a free log retrieval operation binding the contract event 0x62e69886a5df0ba8ffcacbfc1388754e7abd9bde24b036354c561f1acd4e4593.
//
// Solidity: event RecipientUpdated(address indexed account, address indexed recipient)
func (_YVeYFIGauge *YVeYFIGaugeFilterer) FilterRecipientUpdated(opts *bind.FilterOpts, account []common.Address, recipient []common.Address) (*YVeYFIGaugeRecipientUpdatedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _YVeYFIGauge.contract.FilterLogs(opts, "RecipientUpdated", accountRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &YVeYFIGaugeRecipientUpdatedIterator{contract: _YVeYFIGauge.contract, event: "RecipientUpdated", logs: logs, sub: sub}, nil
}

// WatchRecipientUpdated is a free log subscription operation binding the contract event 0x62e69886a5df0ba8ffcacbfc1388754e7abd9bde24b036354c561f1acd4e4593.
//
// Solidity: event RecipientUpdated(address indexed account, address indexed recipient)
func (_YVeYFIGauge *YVeYFIGaugeFilterer) WatchRecipientUpdated(opts *bind.WatchOpts, sink chan<- *YVeYFIGaugeRecipientUpdated, account []common.Address, recipient []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _YVeYFIGauge.contract.WatchLogs(opts, "RecipientUpdated", accountRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YVeYFIGaugeRecipientUpdated)
				if err := _YVeYFIGauge.contract.UnpackLog(event, "RecipientUpdated", log); err != nil {
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

// ParseRecipientUpdated is a log parse operation binding the contract event 0x62e69886a5df0ba8ffcacbfc1388754e7abd9bde24b036354c561f1acd4e4593.
//
// Solidity: event RecipientUpdated(address indexed account, address indexed recipient)
func (_YVeYFIGauge *YVeYFIGaugeFilterer) ParseRecipientUpdated(log types.Log) (*YVeYFIGaugeRecipientUpdated, error) {
	event := new(YVeYFIGaugeRecipientUpdated)
	if err := _YVeYFIGauge.contract.UnpackLog(event, "RecipientUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YVeYFIGaugeRewardPaidIterator is returned from FilterRewardPaid and is used to iterate over the raw logs and unpacked data for RewardPaid events raised by the YVeYFIGauge contract.
type YVeYFIGaugeRewardPaidIterator struct {
	Event *YVeYFIGaugeRewardPaid // Event containing the contract specifics and raw log

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
func (it *YVeYFIGaugeRewardPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YVeYFIGaugeRewardPaid)
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
		it.Event = new(YVeYFIGaugeRewardPaid)
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
func (it *YVeYFIGaugeRewardPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YVeYFIGaugeRewardPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YVeYFIGaugeRewardPaid represents a RewardPaid event raised by the YVeYFIGauge contract.
type YVeYFIGaugeRewardPaid struct {
	User   common.Address
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardPaid is a free log retrieval operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_YVeYFIGauge *YVeYFIGaugeFilterer) FilterRewardPaid(opts *bind.FilterOpts, user []common.Address) (*YVeYFIGaugeRewardPaidIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _YVeYFIGauge.contract.FilterLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return &YVeYFIGaugeRewardPaidIterator{contract: _YVeYFIGauge.contract, event: "RewardPaid", logs: logs, sub: sub}, nil
}

// WatchRewardPaid is a free log subscription operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_YVeYFIGauge *YVeYFIGaugeFilterer) WatchRewardPaid(opts *bind.WatchOpts, sink chan<- *YVeYFIGaugeRewardPaid, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _YVeYFIGauge.contract.WatchLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YVeYFIGaugeRewardPaid)
				if err := _YVeYFIGauge.contract.UnpackLog(event, "RewardPaid", log); err != nil {
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

// ParseRewardPaid is a log parse operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_YVeYFIGauge *YVeYFIGaugeFilterer) ParseRewardPaid(log types.Log) (*YVeYFIGaugeRewardPaid, error) {
	event := new(YVeYFIGaugeRewardPaid)
	if err := _YVeYFIGauge.contract.UnpackLog(event, "RewardPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YVeYFIGaugeRewardsAddedIterator is returned from FilterRewardsAdded and is used to iterate over the raw logs and unpacked data for RewardsAdded events raised by the YVeYFIGauge contract.
type YVeYFIGaugeRewardsAddedIterator struct {
	Event *YVeYFIGaugeRewardsAdded // Event containing the contract specifics and raw log

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
func (it *YVeYFIGaugeRewardsAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YVeYFIGaugeRewardsAdded)
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
		it.Event = new(YVeYFIGaugeRewardsAdded)
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
func (it *YVeYFIGaugeRewardsAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YVeYFIGaugeRewardsAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YVeYFIGaugeRewardsAdded represents a RewardsAdded event raised by the YVeYFIGauge contract.
type YVeYFIGaugeRewardsAdded struct {
	CurrentRewards    *big.Int
	LastUpdateTime    *big.Int
	PeriodFinish      *big.Int
	RewardRate        *big.Int
	HistoricalRewards *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRewardsAdded is a free log retrieval operation binding the contract event 0x944ffd3678415a15cbfef07dd7d9f20cdc6f36d12588a4ba7e8eb440f32c61be.
//
// Solidity: event RewardsAdded(uint256 currentRewards, uint256 lastUpdateTime, uint256 periodFinish, uint256 rewardRate, uint256 historicalRewards)
func (_YVeYFIGauge *YVeYFIGaugeFilterer) FilterRewardsAdded(opts *bind.FilterOpts) (*YVeYFIGaugeRewardsAddedIterator, error) {

	logs, sub, err := _YVeYFIGauge.contract.FilterLogs(opts, "RewardsAdded")
	if err != nil {
		return nil, err
	}
	return &YVeYFIGaugeRewardsAddedIterator{contract: _YVeYFIGauge.contract, event: "RewardsAdded", logs: logs, sub: sub}, nil
}

// WatchRewardsAdded is a free log subscription operation binding the contract event 0x944ffd3678415a15cbfef07dd7d9f20cdc6f36d12588a4ba7e8eb440f32c61be.
//
// Solidity: event RewardsAdded(uint256 currentRewards, uint256 lastUpdateTime, uint256 periodFinish, uint256 rewardRate, uint256 historicalRewards)
func (_YVeYFIGauge *YVeYFIGaugeFilterer) WatchRewardsAdded(opts *bind.WatchOpts, sink chan<- *YVeYFIGaugeRewardsAdded) (event.Subscription, error) {

	logs, sub, err := _YVeYFIGauge.contract.WatchLogs(opts, "RewardsAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YVeYFIGaugeRewardsAdded)
				if err := _YVeYFIGauge.contract.UnpackLog(event, "RewardsAdded", log); err != nil {
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

// ParseRewardsAdded is a log parse operation binding the contract event 0x944ffd3678415a15cbfef07dd7d9f20cdc6f36d12588a4ba7e8eb440f32c61be.
//
// Solidity: event RewardsAdded(uint256 currentRewards, uint256 lastUpdateTime, uint256 periodFinish, uint256 rewardRate, uint256 historicalRewards)
func (_YVeYFIGauge *YVeYFIGaugeFilterer) ParseRewardsAdded(log types.Log) (*YVeYFIGaugeRewardsAdded, error) {
	event := new(YVeYFIGaugeRewardsAdded)
	if err := _YVeYFIGauge.contract.UnpackLog(event, "RewardsAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YVeYFIGaugeRewardsQueuedIterator is returned from FilterRewardsQueued and is used to iterate over the raw logs and unpacked data for RewardsQueued events raised by the YVeYFIGauge contract.
type YVeYFIGaugeRewardsQueuedIterator struct {
	Event *YVeYFIGaugeRewardsQueued // Event containing the contract specifics and raw log

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
func (it *YVeYFIGaugeRewardsQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YVeYFIGaugeRewardsQueued)
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
		it.Event = new(YVeYFIGaugeRewardsQueued)
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
func (it *YVeYFIGaugeRewardsQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YVeYFIGaugeRewardsQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YVeYFIGaugeRewardsQueued represents a RewardsQueued event raised by the YVeYFIGauge contract.
type YVeYFIGaugeRewardsQueued struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardsQueued is a free log retrieval operation binding the contract event 0x1c88aa9a39b1a6357a85c97a3bd4e2b0738e74c68b92928276bc85f495b2450b.
//
// Solidity: event RewardsQueued(address indexed from, uint256 amount)
func (_YVeYFIGauge *YVeYFIGaugeFilterer) FilterRewardsQueued(opts *bind.FilterOpts, from []common.Address) (*YVeYFIGaugeRewardsQueuedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _YVeYFIGauge.contract.FilterLogs(opts, "RewardsQueued", fromRule)
	if err != nil {
		return nil, err
	}
	return &YVeYFIGaugeRewardsQueuedIterator{contract: _YVeYFIGauge.contract, event: "RewardsQueued", logs: logs, sub: sub}, nil
}

// WatchRewardsQueued is a free log subscription operation binding the contract event 0x1c88aa9a39b1a6357a85c97a3bd4e2b0738e74c68b92928276bc85f495b2450b.
//
// Solidity: event RewardsQueued(address indexed from, uint256 amount)
func (_YVeYFIGauge *YVeYFIGaugeFilterer) WatchRewardsQueued(opts *bind.WatchOpts, sink chan<- *YVeYFIGaugeRewardsQueued, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _YVeYFIGauge.contract.WatchLogs(opts, "RewardsQueued", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YVeYFIGaugeRewardsQueued)
				if err := _YVeYFIGauge.contract.UnpackLog(event, "RewardsQueued", log); err != nil {
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

// ParseRewardsQueued is a log parse operation binding the contract event 0x1c88aa9a39b1a6357a85c97a3bd4e2b0738e74c68b92928276bc85f495b2450b.
//
// Solidity: event RewardsQueued(address indexed from, uint256 amount)
func (_YVeYFIGauge *YVeYFIGaugeFilterer) ParseRewardsQueued(log types.Log) (*YVeYFIGaugeRewardsQueued, error) {
	event := new(YVeYFIGaugeRewardsQueued)
	if err := _YVeYFIGauge.contract.UnpackLog(event, "RewardsQueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YVeYFIGaugeSweepIterator is returned from FilterSweep and is used to iterate over the raw logs and unpacked data for Sweep events raised by the YVeYFIGauge contract.
type YVeYFIGaugeSweepIterator struct {
	Event *YVeYFIGaugeSweep // Event containing the contract specifics and raw log

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
func (it *YVeYFIGaugeSweepIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YVeYFIGaugeSweep)
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
		it.Event = new(YVeYFIGaugeSweep)
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
func (it *YVeYFIGaugeSweepIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YVeYFIGaugeSweepIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YVeYFIGaugeSweep represents a Sweep event raised by the YVeYFIGauge contract.
type YVeYFIGaugeSweep struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSweep is a free log retrieval operation binding the contract event 0xab2246061d7b0dd3631d037e3f6da75782ae489eeb9f6af878a4b25df9b07c77.
//
// Solidity: event Sweep(address indexed token, uint256 amount)
func (_YVeYFIGauge *YVeYFIGaugeFilterer) FilterSweep(opts *bind.FilterOpts, token []common.Address) (*YVeYFIGaugeSweepIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _YVeYFIGauge.contract.FilterLogs(opts, "Sweep", tokenRule)
	if err != nil {
		return nil, err
	}
	return &YVeYFIGaugeSweepIterator{contract: _YVeYFIGauge.contract, event: "Sweep", logs: logs, sub: sub}, nil
}

// WatchSweep is a free log subscription operation binding the contract event 0xab2246061d7b0dd3631d037e3f6da75782ae489eeb9f6af878a4b25df9b07c77.
//
// Solidity: event Sweep(address indexed token, uint256 amount)
func (_YVeYFIGauge *YVeYFIGaugeFilterer) WatchSweep(opts *bind.WatchOpts, sink chan<- *YVeYFIGaugeSweep, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _YVeYFIGauge.contract.WatchLogs(opts, "Sweep", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YVeYFIGaugeSweep)
				if err := _YVeYFIGauge.contract.UnpackLog(event, "Sweep", log); err != nil {
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

// ParseSweep is a log parse operation binding the contract event 0xab2246061d7b0dd3631d037e3f6da75782ae489eeb9f6af878a4b25df9b07c77.
//
// Solidity: event Sweep(address indexed token, uint256 amount)
func (_YVeYFIGauge *YVeYFIGaugeFilterer) ParseSweep(log types.Log) (*YVeYFIGaugeSweep, error) {
	event := new(YVeYFIGaugeSweep)
	if err := _YVeYFIGauge.contract.UnpackLog(event, "Sweep", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YVeYFIGaugeTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the YVeYFIGauge contract.
type YVeYFIGaugeTransferIterator struct {
	Event *YVeYFIGaugeTransfer // Event containing the contract specifics and raw log

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
func (it *YVeYFIGaugeTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YVeYFIGaugeTransfer)
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
		it.Event = new(YVeYFIGaugeTransfer)
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
func (it *YVeYFIGaugeTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YVeYFIGaugeTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YVeYFIGaugeTransfer represents a Transfer event raised by the YVeYFIGauge contract.
type YVeYFIGaugeTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_YVeYFIGauge *YVeYFIGaugeFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*YVeYFIGaugeTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _YVeYFIGauge.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &YVeYFIGaugeTransferIterator{contract: _YVeYFIGauge.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_YVeYFIGauge *YVeYFIGaugeFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *YVeYFIGaugeTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _YVeYFIGauge.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YVeYFIGaugeTransfer)
				if err := _YVeYFIGauge.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_YVeYFIGauge *YVeYFIGaugeFilterer) ParseTransfer(log types.Log) (*YVeYFIGaugeTransfer, error) {
	event := new(YVeYFIGaugeTransfer)
	if err := _YVeYFIGauge.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YVeYFIGaugeTransferredPenaltyIterator is returned from FilterTransferredPenalty and is used to iterate over the raw logs and unpacked data for TransferredPenalty events raised by the YVeYFIGauge contract.
type YVeYFIGaugeTransferredPenaltyIterator struct {
	Event *YVeYFIGaugeTransferredPenalty // Event containing the contract specifics and raw log

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
func (it *YVeYFIGaugeTransferredPenaltyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YVeYFIGaugeTransferredPenalty)
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
		it.Event = new(YVeYFIGaugeTransferredPenalty)
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
func (it *YVeYFIGaugeTransferredPenaltyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YVeYFIGaugeTransferredPenaltyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YVeYFIGaugeTransferredPenalty represents a TransferredPenalty event raised by the YVeYFIGauge contract.
type YVeYFIGaugeTransferredPenalty struct {
	Account    common.Address
	Transfered *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTransferredPenalty is a free log retrieval operation binding the contract event 0xfdcc759119f4a689ba608afdccb078153573a5a615700713ebb84704609694cc.
//
// Solidity: event TransferredPenalty(address indexed account, uint256 transfered)
func (_YVeYFIGauge *YVeYFIGaugeFilterer) FilterTransferredPenalty(opts *bind.FilterOpts, account []common.Address) (*YVeYFIGaugeTransferredPenaltyIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _YVeYFIGauge.contract.FilterLogs(opts, "TransferredPenalty", accountRule)
	if err != nil {
		return nil, err
	}
	return &YVeYFIGaugeTransferredPenaltyIterator{contract: _YVeYFIGauge.contract, event: "TransferredPenalty", logs: logs, sub: sub}, nil
}

// WatchTransferredPenalty is a free log subscription operation binding the contract event 0xfdcc759119f4a689ba608afdccb078153573a5a615700713ebb84704609694cc.
//
// Solidity: event TransferredPenalty(address indexed account, uint256 transfered)
func (_YVeYFIGauge *YVeYFIGaugeFilterer) WatchTransferredPenalty(opts *bind.WatchOpts, sink chan<- *YVeYFIGaugeTransferredPenalty, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _YVeYFIGauge.contract.WatchLogs(opts, "TransferredPenalty", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YVeYFIGaugeTransferredPenalty)
				if err := _YVeYFIGauge.contract.UnpackLog(event, "TransferredPenalty", log); err != nil {
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

// ParseTransferredPenalty is a log parse operation binding the contract event 0xfdcc759119f4a689ba608afdccb078153573a5a615700713ebb84704609694cc.
//
// Solidity: event TransferredPenalty(address indexed account, uint256 transfered)
func (_YVeYFIGauge *YVeYFIGaugeFilterer) ParseTransferredPenalty(log types.Log) (*YVeYFIGaugeTransferredPenalty, error) {
	event := new(YVeYFIGaugeTransferredPenalty)
	if err := _YVeYFIGauge.contract.UnpackLog(event, "TransferredPenalty", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YVeYFIGaugeUpdatedRewardsIterator is returned from FilterUpdatedRewards and is used to iterate over the raw logs and unpacked data for UpdatedRewards events raised by the YVeYFIGauge contract.
type YVeYFIGaugeUpdatedRewardsIterator struct {
	Event *YVeYFIGaugeUpdatedRewards // Event containing the contract specifics and raw log

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
func (it *YVeYFIGaugeUpdatedRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YVeYFIGaugeUpdatedRewards)
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
		it.Event = new(YVeYFIGaugeUpdatedRewards)
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
func (it *YVeYFIGaugeUpdatedRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YVeYFIGaugeUpdatedRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YVeYFIGaugeUpdatedRewards represents a UpdatedRewards event raised by the YVeYFIGauge contract.
type YVeYFIGaugeUpdatedRewards struct {
	Account                common.Address
	RewardPerTokenStored   *big.Int
	LastUpdateTime         *big.Int
	Rewards                *big.Int
	UserRewardPerTokenPaid *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterUpdatedRewards is a free log retrieval operation binding the contract event 0xfbe590c835e1c07f8e971c36021d1be46f43f7b0b6dc5413dbd5753590569d58.
//
// Solidity: event UpdatedRewards(address indexed account, uint256 rewardPerTokenStored, uint256 lastUpdateTime, uint256 rewards, uint256 userRewardPerTokenPaid)
func (_YVeYFIGauge *YVeYFIGaugeFilterer) FilterUpdatedRewards(opts *bind.FilterOpts, account []common.Address) (*YVeYFIGaugeUpdatedRewardsIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _YVeYFIGauge.contract.FilterLogs(opts, "UpdatedRewards", accountRule)
	if err != nil {
		return nil, err
	}
	return &YVeYFIGaugeUpdatedRewardsIterator{contract: _YVeYFIGauge.contract, event: "UpdatedRewards", logs: logs, sub: sub}, nil
}

// WatchUpdatedRewards is a free log subscription operation binding the contract event 0xfbe590c835e1c07f8e971c36021d1be46f43f7b0b6dc5413dbd5753590569d58.
//
// Solidity: event UpdatedRewards(address indexed account, uint256 rewardPerTokenStored, uint256 lastUpdateTime, uint256 rewards, uint256 userRewardPerTokenPaid)
func (_YVeYFIGauge *YVeYFIGaugeFilterer) WatchUpdatedRewards(opts *bind.WatchOpts, sink chan<- *YVeYFIGaugeUpdatedRewards, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _YVeYFIGauge.contract.WatchLogs(opts, "UpdatedRewards", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YVeYFIGaugeUpdatedRewards)
				if err := _YVeYFIGauge.contract.UnpackLog(event, "UpdatedRewards", log); err != nil {
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

// ParseUpdatedRewards is a log parse operation binding the contract event 0xfbe590c835e1c07f8e971c36021d1be46f43f7b0b6dc5413dbd5753590569d58.
//
// Solidity: event UpdatedRewards(address indexed account, uint256 rewardPerTokenStored, uint256 lastUpdateTime, uint256 rewards, uint256 userRewardPerTokenPaid)
func (_YVeYFIGauge *YVeYFIGaugeFilterer) ParseUpdatedRewards(log types.Log) (*YVeYFIGaugeUpdatedRewards, error) {
	event := new(YVeYFIGaugeUpdatedRewards)
	if err := _YVeYFIGauge.contract.UnpackLog(event, "UpdatedRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YVeYFIGaugeWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the YVeYFIGauge contract.
type YVeYFIGaugeWithdrawIterator struct {
	Event *YVeYFIGaugeWithdraw // Event containing the contract specifics and raw log

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
func (it *YVeYFIGaugeWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YVeYFIGaugeWithdraw)
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
		it.Event = new(YVeYFIGaugeWithdraw)
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
func (it *YVeYFIGaugeWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YVeYFIGaugeWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YVeYFIGaugeWithdraw represents a Withdraw event raised by the YVeYFIGauge contract.
type YVeYFIGaugeWithdraw struct {
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
func (_YVeYFIGauge *YVeYFIGaugeFilterer) FilterWithdraw(opts *bind.FilterOpts, caller []common.Address, receiver []common.Address, owner []common.Address) (*YVeYFIGaugeWithdrawIterator, error) {

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

	logs, sub, err := _YVeYFIGauge.contract.FilterLogs(opts, "Withdraw", callerRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &YVeYFIGaugeWithdrawIterator{contract: _YVeYFIGauge.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed caller, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_YVeYFIGauge *YVeYFIGaugeFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *YVeYFIGaugeWithdraw, caller []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _YVeYFIGauge.contract.WatchLogs(opts, "Withdraw", callerRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YVeYFIGaugeWithdraw)
				if err := _YVeYFIGauge.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_YVeYFIGauge *YVeYFIGaugeFilterer) ParseWithdraw(log types.Log) (*YVeYFIGaugeWithdraw, error) {
	event := new(YVeYFIGaugeWithdraw)
	if err := _YVeYFIGauge.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
