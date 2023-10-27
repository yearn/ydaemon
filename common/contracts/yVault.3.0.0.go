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

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	Activation  *big.Int
	LastReport  *big.Int
	CurrentDebt *big.Int
	MaxDebt     *big.Int
}

// Yvault300MetaData contains all meta data concerning the Yvault300 contract.
var Yvault300MetaData = &bind.MetaData{
	ABI: "[{\"name\":\"Deposit\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true},{\"name\":\"assets\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true},{\"name\":\"assets\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Transfer\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StrategyChanged\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":true},{\"name\":\"change_type\",\"type\":\"uint256\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StrategyReported\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":true},{\"name\":\"gain\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"loss\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"current_debt\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"protocol_fees\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"total_fees\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"total_refunds\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"DebtUpdated\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":true},{\"name\":\"current_debt\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"new_debt\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RoleSet\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true},{\"name\":\"role\",\"type\":\"uint256\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RoleStatusChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"uint256\",\"indexed\":true},{\"name\":\"status\",\"type\":\"uint256\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateRoleManager\",\"inputs\":[{\"name\":\"role_manager\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateAccountant\",\"inputs\":[{\"name\":\"accountant\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateDepositLimitModule\",\"inputs\":[{\"name\":\"deposit_limit_module\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateWithdrawLimitModule\",\"inputs\":[{\"name\":\"withdraw_limit_module\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateDefaultQueue\",\"inputs\":[{\"name\":\"new_default_queue\",\"type\":\"address[]\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateUseDefaultQueue\",\"inputs\":[{\"name\":\"use_default_queue\",\"type\":\"bool\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdatedMaxDebtForStrategy\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":true},{\"name\":\"new_debt\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateDepositLimit\",\"inputs\":[{\"name\":\"deposit_limit\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateMinimumTotalIdle\",\"inputs\":[{\"name\":\"minimum_total_idle\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateProfitMaxUnlockTime\",\"inputs\":[{\"name\":\"profit_max_unlock_time\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"DebtPurchased\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":true},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Shutdown\",\"inputs\":[],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"asset\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"role_manager\",\"type\":\"address\"},{\"name\":\"profit_max_unlock_time\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_accountant\",\"inputs\":[{\"name\":\"new_accountant\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_default_queue\",\"inputs\":[{\"name\":\"new_default_queue\",\"type\":\"address[]\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_use_default_queue\",\"inputs\":[{\"name\":\"use_default_queue\",\"type\":\"bool\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_deposit_limit\",\"inputs\":[{\"name\":\"deposit_limit\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_deposit_limit_module\",\"inputs\":[{\"name\":\"deposit_limit_module\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_withdraw_limit_module\",\"inputs\":[{\"name\":\"withdraw_limit_module\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_minimum_total_idle\",\"inputs\":[{\"name\":\"minimum_total_idle\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"setProfitMaxUnlockTime\",\"inputs\":[{\"name\":\"new_profit_max_unlock_time\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_role\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"role\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_role\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"role\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_role\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"role\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_open_role\",\"inputs\":[{\"name\":\"role\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"close_open_role\",\"inputs\":[{\"name\":\"role\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transfer_role_manager\",\"inputs\":[{\"name\":\"role_manager\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"accept_role_manager\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"isShutdown\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"unlockedShares\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"pricePerShare\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_default_queue\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"process_report\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"buy_debt\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_strategy\",\"inputs\":[{\"name\":\"new_strategy\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"revoke_strategy\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"force_revoke_strategy\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"update_max_debt_for_strategy\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"},{\"name\":\"new_max_debt\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"update_debt\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"},{\"name\":\"target_debt\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"shutdown_vault\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"owner\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"max_loss\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"max_loss\",\"type\":\"uint256\"},{\"name\":\"strategies\",\"type\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"redeem\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"owner\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"redeem\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"max_loss\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"redeem\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"max_loss\",\"type\":\"uint256\"},{\"name\":\"strategies\",\"type\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"increaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"decreaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"permit\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"asset\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"totalAssets\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"totalIdle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"totalDebt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"convertToShares\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"previewDeposit\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"previewMint\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"convertToAssets\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"maxDeposit\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"maxMint\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"maxWithdraw\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"maxWithdraw\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"max_loss\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"maxWithdraw\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"max_loss\",\"type\":\"uint256\"},{\"name\":\"strategies\",\"type\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"maxRedeem\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"maxRedeem\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"max_loss\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"maxRedeem\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"max_loss\",\"type\":\"uint256\"},{\"name\":\"strategies\",\"type\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"previewWithdraw\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"previewRedeem\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"apiVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"assess_share_of_unrealised_losses\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"},{\"name\":\"assets_needed\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"profitMaxUnlockTime\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"fullProfitUnlockDate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"profitUnlockingRate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"lastProfitUpdate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"DOMAIN_SEPARATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"FACTORY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"strategies\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"components\":[{\"name\":\"activation\",\"type\":\"uint256\"},{\"name\":\"last_report\",\"type\":\"uint256\"},{\"name\":\"current_debt\",\"type\":\"uint256\"},{\"name\":\"max_debt\",\"type\":\"uint256\"}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"default_queue\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"use_default_queue\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"},{\"name\":\"arg1\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"total_supply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"minimum_total_idle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"deposit_limit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"accountant\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"deposit_limit_module\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"withdraw_limit_module\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"roles\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"open_roles\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"role_manager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_role_manager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]}]",
}

// Yvault300ABI is the input ABI used to generate the binding from.
// Deprecated: Use Yvault300MetaData.ABI instead.
var Yvault300ABI = Yvault300MetaData.ABI

// Yvault300 is an auto generated Go binding around an Ethereum contract.
type Yvault300 struct {
	Yvault300Caller     // Read-only binding to the contract
	Yvault300Transactor // Write-only binding to the contract
	Yvault300Filterer   // Log filterer for contract events
}

// Yvault300Caller is an auto generated read-only Go binding around an Ethereum contract.
type Yvault300Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Yvault300Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Yvault300Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Yvault300Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Yvault300Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Yvault300Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Yvault300Session struct {
	Contract     *Yvault300        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Yvault300CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Yvault300CallerSession struct {
	Contract *Yvault300Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// Yvault300TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Yvault300TransactorSession struct {
	Contract     *Yvault300Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// Yvault300Raw is an auto generated low-level Go binding around an Ethereum contract.
type Yvault300Raw struct {
	Contract *Yvault300 // Generic contract binding to access the raw methods on
}

// Yvault300CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Yvault300CallerRaw struct {
	Contract *Yvault300Caller // Generic read-only contract binding to access the raw methods on
}

// Yvault300TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Yvault300TransactorRaw struct {
	Contract *Yvault300Transactor // Generic write-only contract binding to access the raw methods on
}

// NewYvault300 creates a new instance of Yvault300, bound to a specific deployed contract.
func NewYvault300(address common.Address, backend bind.ContractBackend) (*Yvault300, error) {
	contract, err := bindYvault300(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Yvault300{Yvault300Caller: Yvault300Caller{contract: contract}, Yvault300Transactor: Yvault300Transactor{contract: contract}, Yvault300Filterer: Yvault300Filterer{contract: contract}}, nil
}

// NewYvault300Caller creates a new read-only instance of Yvault300, bound to a specific deployed contract.
func NewYvault300Caller(address common.Address, caller bind.ContractCaller) (*Yvault300Caller, error) {
	contract, err := bindYvault300(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Yvault300Caller{contract: contract}, nil
}

// NewYvault300Transactor creates a new write-only instance of Yvault300, bound to a specific deployed contract.
func NewYvault300Transactor(address common.Address, transactor bind.ContractTransactor) (*Yvault300Transactor, error) {
	contract, err := bindYvault300(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Yvault300Transactor{contract: contract}, nil
}

// NewYvault300Filterer creates a new log filterer instance of Yvault300, bound to a specific deployed contract.
func NewYvault300Filterer(address common.Address, filterer bind.ContractFilterer) (*Yvault300Filterer, error) {
	contract, err := bindYvault300(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Yvault300Filterer{contract: contract}, nil
}

// bindYvault300 binds a generic wrapper to an already deployed contract.
func bindYvault300(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Yvault300ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Yvault300 *Yvault300Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Yvault300.Contract.Yvault300Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Yvault300 *Yvault300Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yvault300.Contract.Yvault300Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Yvault300 *Yvault300Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Yvault300.Contract.Yvault300Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Yvault300 *Yvault300CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Yvault300.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Yvault300 *Yvault300TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yvault300.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Yvault300 *Yvault300TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Yvault300.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Yvault300 *Yvault300Caller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Yvault300 *Yvault300Session) DOMAINSEPARATOR() ([32]byte, error) {
	return _Yvault300.Contract.DOMAINSEPARATOR(&_Yvault300.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Yvault300 *Yvault300CallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Yvault300.Contract.DOMAINSEPARATOR(&_Yvault300.CallOpts)
}

// FACTORY is a free data retrieval call binding the contract method 0x2dd31000.
//
// Solidity: function FACTORY() view returns(address)
func (_Yvault300 *Yvault300Caller) FACTORY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "FACTORY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FACTORY is a free data retrieval call binding the contract method 0x2dd31000.
//
// Solidity: function FACTORY() view returns(address)
func (_Yvault300 *Yvault300Session) FACTORY() (common.Address, error) {
	return _Yvault300.Contract.FACTORY(&_Yvault300.CallOpts)
}

// FACTORY is a free data retrieval call binding the contract method 0x2dd31000.
//
// Solidity: function FACTORY() view returns(address)
func (_Yvault300 *Yvault300CallerSession) FACTORY() (common.Address, error) {
	return _Yvault300.Contract.FACTORY(&_Yvault300.CallOpts)
}

// Accountant is a free data retrieval call binding the contract method 0x4fb3ccc5.
//
// Solidity: function accountant() view returns(address)
func (_Yvault300 *Yvault300Caller) Accountant(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "accountant")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Accountant is a free data retrieval call binding the contract method 0x4fb3ccc5.
//
// Solidity: function accountant() view returns(address)
func (_Yvault300 *Yvault300Session) Accountant() (common.Address, error) {
	return _Yvault300.Contract.Accountant(&_Yvault300.CallOpts)
}

// Accountant is a free data retrieval call binding the contract method 0x4fb3ccc5.
//
// Solidity: function accountant() view returns(address)
func (_Yvault300 *Yvault300CallerSession) Accountant() (common.Address, error) {
	return _Yvault300.Contract.Accountant(&_Yvault300.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_Yvault300 *Yvault300Caller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_Yvault300 *Yvault300Session) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Yvault300.Contract.Allowance(&_Yvault300.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_Yvault300 *Yvault300CallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Yvault300.Contract.Allowance(&_Yvault300.CallOpts, arg0, arg1)
}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() view returns(string)
func (_Yvault300 *Yvault300Caller) ApiVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "apiVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() view returns(string)
func (_Yvault300 *Yvault300Session) ApiVersion() (string, error) {
	return _Yvault300.Contract.ApiVersion(&_Yvault300.CallOpts)
}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() view returns(string)
func (_Yvault300 *Yvault300CallerSession) ApiVersion() (string, error) {
	return _Yvault300.Contract.ApiVersion(&_Yvault300.CallOpts)
}

// AssessShareOfUnrealisedLosses is a free data retrieval call binding the contract method 0x66d3ae57.
//
// Solidity: function assess_share_of_unrealised_losses(address strategy, uint256 assets_needed) view returns(uint256)
func (_Yvault300 *Yvault300Caller) AssessShareOfUnrealisedLosses(opts *bind.CallOpts, strategy common.Address, assets_needed *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "assess_share_of_unrealised_losses", strategy, assets_needed)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AssessShareOfUnrealisedLosses is a free data retrieval call binding the contract method 0x66d3ae57.
//
// Solidity: function assess_share_of_unrealised_losses(address strategy, uint256 assets_needed) view returns(uint256)
func (_Yvault300 *Yvault300Session) AssessShareOfUnrealisedLosses(strategy common.Address, assets_needed *big.Int) (*big.Int, error) {
	return _Yvault300.Contract.AssessShareOfUnrealisedLosses(&_Yvault300.CallOpts, strategy, assets_needed)
}

// AssessShareOfUnrealisedLosses is a free data retrieval call binding the contract method 0x66d3ae57.
//
// Solidity: function assess_share_of_unrealised_losses(address strategy, uint256 assets_needed) view returns(uint256)
func (_Yvault300 *Yvault300CallerSession) AssessShareOfUnrealisedLosses(strategy common.Address, assets_needed *big.Int) (*big.Int, error) {
	return _Yvault300.Contract.AssessShareOfUnrealisedLosses(&_Yvault300.CallOpts, strategy, assets_needed)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_Yvault300 *Yvault300Caller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_Yvault300 *Yvault300Session) Asset() (common.Address, error) {
	return _Yvault300.Contract.Asset(&_Yvault300.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_Yvault300 *Yvault300CallerSession) Asset() (common.Address, error) {
	return _Yvault300.Contract.Asset(&_Yvault300.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address addr) view returns(uint256)
func (_Yvault300 *Yvault300Caller) BalanceOf(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "balanceOf", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address addr) view returns(uint256)
func (_Yvault300 *Yvault300Session) BalanceOf(addr common.Address) (*big.Int, error) {
	return _Yvault300.Contract.BalanceOf(&_Yvault300.CallOpts, addr)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address addr) view returns(uint256)
func (_Yvault300 *Yvault300CallerSession) BalanceOf(addr common.Address) (*big.Int, error) {
	return _Yvault300.Contract.BalanceOf(&_Yvault300.CallOpts, addr)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_Yvault300 *Yvault300Caller) ConvertToAssets(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "convertToAssets", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_Yvault300 *Yvault300Session) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _Yvault300.Contract.ConvertToAssets(&_Yvault300.CallOpts, shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_Yvault300 *Yvault300CallerSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _Yvault300.Contract.ConvertToAssets(&_Yvault300.CallOpts, shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_Yvault300 *Yvault300Caller) ConvertToShares(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "convertToShares", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_Yvault300 *Yvault300Session) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _Yvault300.Contract.ConvertToShares(&_Yvault300.CallOpts, assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_Yvault300 *Yvault300CallerSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _Yvault300.Contract.ConvertToShares(&_Yvault300.CallOpts, assets)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Yvault300 *Yvault300Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Yvault300 *Yvault300Session) Decimals() (uint8, error) {
	return _Yvault300.Contract.Decimals(&_Yvault300.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Yvault300 *Yvault300CallerSession) Decimals() (uint8, error) {
	return _Yvault300.Contract.Decimals(&_Yvault300.CallOpts)
}

// DefaultQueue is a free data retrieval call binding the contract method 0x8bf03b9e.
//
// Solidity: function default_queue(uint256 arg0) view returns(address)
func (_Yvault300 *Yvault300Caller) DefaultQueue(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "default_queue", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultQueue is a free data retrieval call binding the contract method 0x8bf03b9e.
//
// Solidity: function default_queue(uint256 arg0) view returns(address)
func (_Yvault300 *Yvault300Session) DefaultQueue(arg0 *big.Int) (common.Address, error) {
	return _Yvault300.Contract.DefaultQueue(&_Yvault300.CallOpts, arg0)
}

// DefaultQueue is a free data retrieval call binding the contract method 0x8bf03b9e.
//
// Solidity: function default_queue(uint256 arg0) view returns(address)
func (_Yvault300 *Yvault300CallerSession) DefaultQueue(arg0 *big.Int) (common.Address, error) {
	return _Yvault300.Contract.DefaultQueue(&_Yvault300.CallOpts, arg0)
}

// DepositLimit is a free data retrieval call binding the contract method 0xe46a5797.
//
// Solidity: function deposit_limit() view returns(uint256)
func (_Yvault300 *Yvault300Caller) DepositLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "deposit_limit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositLimit is a free data retrieval call binding the contract method 0xe46a5797.
//
// Solidity: function deposit_limit() view returns(uint256)
func (_Yvault300 *Yvault300Session) DepositLimit() (*big.Int, error) {
	return _Yvault300.Contract.DepositLimit(&_Yvault300.CallOpts)
}

// DepositLimit is a free data retrieval call binding the contract method 0xe46a5797.
//
// Solidity: function deposit_limit() view returns(uint256)
func (_Yvault300 *Yvault300CallerSession) DepositLimit() (*big.Int, error) {
	return _Yvault300.Contract.DepositLimit(&_Yvault300.CallOpts)
}

// DepositLimitModule is a free data retrieval call binding the contract method 0x61c2ccf4.
//
// Solidity: function deposit_limit_module() view returns(address)
func (_Yvault300 *Yvault300Caller) DepositLimitModule(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "deposit_limit_module")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DepositLimitModule is a free data retrieval call binding the contract method 0x61c2ccf4.
//
// Solidity: function deposit_limit_module() view returns(address)
func (_Yvault300 *Yvault300Session) DepositLimitModule() (common.Address, error) {
	return _Yvault300.Contract.DepositLimitModule(&_Yvault300.CallOpts)
}

// DepositLimitModule is a free data retrieval call binding the contract method 0x61c2ccf4.
//
// Solidity: function deposit_limit_module() view returns(address)
func (_Yvault300 *Yvault300CallerSession) DepositLimitModule() (common.Address, error) {
	return _Yvault300.Contract.DepositLimitModule(&_Yvault300.CallOpts)
}

// FullProfitUnlockDate is a free data retrieval call binding the contract method 0x2d632692.
//
// Solidity: function fullProfitUnlockDate() view returns(uint256)
func (_Yvault300 *Yvault300Caller) FullProfitUnlockDate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "fullProfitUnlockDate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FullProfitUnlockDate is a free data retrieval call binding the contract method 0x2d632692.
//
// Solidity: function fullProfitUnlockDate() view returns(uint256)
func (_Yvault300 *Yvault300Session) FullProfitUnlockDate() (*big.Int, error) {
	return _Yvault300.Contract.FullProfitUnlockDate(&_Yvault300.CallOpts)
}

// FullProfitUnlockDate is a free data retrieval call binding the contract method 0x2d632692.
//
// Solidity: function fullProfitUnlockDate() view returns(uint256)
func (_Yvault300 *Yvault300CallerSession) FullProfitUnlockDate() (*big.Int, error) {
	return _Yvault300.Contract.FullProfitUnlockDate(&_Yvault300.CallOpts)
}

// FutureRoleManager is a free data retrieval call binding the contract method 0x9a98f418.
//
// Solidity: function future_role_manager() view returns(address)
func (_Yvault300 *Yvault300Caller) FutureRoleManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "future_role_manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureRoleManager is a free data retrieval call binding the contract method 0x9a98f418.
//
// Solidity: function future_role_manager() view returns(address)
func (_Yvault300 *Yvault300Session) FutureRoleManager() (common.Address, error) {
	return _Yvault300.Contract.FutureRoleManager(&_Yvault300.CallOpts)
}

// FutureRoleManager is a free data retrieval call binding the contract method 0x9a98f418.
//
// Solidity: function future_role_manager() view returns(address)
func (_Yvault300 *Yvault300CallerSession) FutureRoleManager() (common.Address, error) {
	return _Yvault300.Contract.FutureRoleManager(&_Yvault300.CallOpts)
}

// GetDefaultQueue is a free data retrieval call binding the contract method 0xa9bbf1cc.
//
// Solidity: function get_default_queue() view returns(address[])
func (_Yvault300 *Yvault300Caller) GetDefaultQueue(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "get_default_queue")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetDefaultQueue is a free data retrieval call binding the contract method 0xa9bbf1cc.
//
// Solidity: function get_default_queue() view returns(address[])
func (_Yvault300 *Yvault300Session) GetDefaultQueue() ([]common.Address, error) {
	return _Yvault300.Contract.GetDefaultQueue(&_Yvault300.CallOpts)
}

// GetDefaultQueue is a free data retrieval call binding the contract method 0xa9bbf1cc.
//
// Solidity: function get_default_queue() view returns(address[])
func (_Yvault300 *Yvault300CallerSession) GetDefaultQueue() ([]common.Address, error) {
	return _Yvault300.Contract.GetDefaultQueue(&_Yvault300.CallOpts)
}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_Yvault300 *Yvault300Caller) IsShutdown(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "isShutdown")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_Yvault300 *Yvault300Session) IsShutdown() (bool, error) {
	return _Yvault300.Contract.IsShutdown(&_Yvault300.CallOpts)
}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_Yvault300 *Yvault300CallerSession) IsShutdown() (bool, error) {
	return _Yvault300.Contract.IsShutdown(&_Yvault300.CallOpts)
}

// LastProfitUpdate is a free data retrieval call binding the contract method 0x8afca8f0.
//
// Solidity: function lastProfitUpdate() view returns(uint256)
func (_Yvault300 *Yvault300Caller) LastProfitUpdate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "lastProfitUpdate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastProfitUpdate is a free data retrieval call binding the contract method 0x8afca8f0.
//
// Solidity: function lastProfitUpdate() view returns(uint256)
func (_Yvault300 *Yvault300Session) LastProfitUpdate() (*big.Int, error) {
	return _Yvault300.Contract.LastProfitUpdate(&_Yvault300.CallOpts)
}

// LastProfitUpdate is a free data retrieval call binding the contract method 0x8afca8f0.
//
// Solidity: function lastProfitUpdate() view returns(uint256)
func (_Yvault300 *Yvault300CallerSession) LastProfitUpdate() (*big.Int, error) {
	return _Yvault300.Contract.LastProfitUpdate(&_Yvault300.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address receiver) view returns(uint256)
func (_Yvault300 *Yvault300Caller) MaxDeposit(opts *bind.CallOpts, receiver common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "maxDeposit", receiver)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address receiver) view returns(uint256)
func (_Yvault300 *Yvault300Session) MaxDeposit(receiver common.Address) (*big.Int, error) {
	return _Yvault300.Contract.MaxDeposit(&_Yvault300.CallOpts, receiver)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address receiver) view returns(uint256)
func (_Yvault300 *Yvault300CallerSession) MaxDeposit(receiver common.Address) (*big.Int, error) {
	return _Yvault300.Contract.MaxDeposit(&_Yvault300.CallOpts, receiver)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address receiver) view returns(uint256)
func (_Yvault300 *Yvault300Caller) MaxMint(opts *bind.CallOpts, receiver common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "maxMint", receiver)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address receiver) view returns(uint256)
func (_Yvault300 *Yvault300Session) MaxMint(receiver common.Address) (*big.Int, error) {
	return _Yvault300.Contract.MaxMint(&_Yvault300.CallOpts, receiver)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address receiver) view returns(uint256)
func (_Yvault300 *Yvault300CallerSession) MaxMint(receiver common.Address) (*big.Int, error) {
	return _Yvault300.Contract.MaxMint(&_Yvault300.CallOpts, receiver)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_Yvault300 *Yvault300Caller) MaxRedeem(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "maxRedeem", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_Yvault300 *Yvault300Session) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _Yvault300.Contract.MaxRedeem(&_Yvault300.CallOpts, owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_Yvault300 *Yvault300CallerSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _Yvault300.Contract.MaxRedeem(&_Yvault300.CallOpts, owner)
}

// MaxRedeem0 is a free data retrieval call binding the contract method 0x4abe4137.
//
// Solidity: function maxRedeem(address owner, uint256 max_loss) view returns(uint256)
func (_Yvault300 *Yvault300Caller) MaxRedeem0(opts *bind.CallOpts, owner common.Address, max_loss *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "maxRedeem0", owner, max_loss)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem0 is a free data retrieval call binding the contract method 0x4abe4137.
//
// Solidity: function maxRedeem(address owner, uint256 max_loss) view returns(uint256)
func (_Yvault300 *Yvault300Session) MaxRedeem0(owner common.Address, max_loss *big.Int) (*big.Int, error) {
	return _Yvault300.Contract.MaxRedeem0(&_Yvault300.CallOpts, owner, max_loss)
}

// MaxRedeem0 is a free data retrieval call binding the contract method 0x4abe4137.
//
// Solidity: function maxRedeem(address owner, uint256 max_loss) view returns(uint256)
func (_Yvault300 *Yvault300CallerSession) MaxRedeem0(owner common.Address, max_loss *big.Int) (*big.Int, error) {
	return _Yvault300.Contract.MaxRedeem0(&_Yvault300.CallOpts, owner, max_loss)
}

// MaxRedeem1 is a free data retrieval call binding the contract method 0x34b5fab6.
//
// Solidity: function maxRedeem(address owner, uint256 max_loss, address[] strategies) view returns(uint256)
func (_Yvault300 *Yvault300Caller) MaxRedeem1(opts *bind.CallOpts, owner common.Address, max_loss *big.Int, strategies []common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "maxRedeem1", owner, max_loss, strategies)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem1 is a free data retrieval call binding the contract method 0x34b5fab6.
//
// Solidity: function maxRedeem(address owner, uint256 max_loss, address[] strategies) view returns(uint256)
func (_Yvault300 *Yvault300Session) MaxRedeem1(owner common.Address, max_loss *big.Int, strategies []common.Address) (*big.Int, error) {
	return _Yvault300.Contract.MaxRedeem1(&_Yvault300.CallOpts, owner, max_loss, strategies)
}

// MaxRedeem1 is a free data retrieval call binding the contract method 0x34b5fab6.
//
// Solidity: function maxRedeem(address owner, uint256 max_loss, address[] strategies) view returns(uint256)
func (_Yvault300 *Yvault300CallerSession) MaxRedeem1(owner common.Address, max_loss *big.Int, strategies []common.Address) (*big.Int, error) {
	return _Yvault300.Contract.MaxRedeem1(&_Yvault300.CallOpts, owner, max_loss, strategies)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_Yvault300 *Yvault300Caller) MaxWithdraw(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "maxWithdraw", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_Yvault300 *Yvault300Session) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _Yvault300.Contract.MaxWithdraw(&_Yvault300.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_Yvault300 *Yvault300CallerSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _Yvault300.Contract.MaxWithdraw(&_Yvault300.CallOpts, owner)
}

// MaxWithdraw0 is a free data retrieval call binding the contract method 0x85b68756.
//
// Solidity: function maxWithdraw(address owner, uint256 max_loss) view returns(uint256)
func (_Yvault300 *Yvault300Caller) MaxWithdraw0(opts *bind.CallOpts, owner common.Address, max_loss *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "maxWithdraw0", owner, max_loss)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw0 is a free data retrieval call binding the contract method 0x85b68756.
//
// Solidity: function maxWithdraw(address owner, uint256 max_loss) view returns(uint256)
func (_Yvault300 *Yvault300Session) MaxWithdraw0(owner common.Address, max_loss *big.Int) (*big.Int, error) {
	return _Yvault300.Contract.MaxWithdraw0(&_Yvault300.CallOpts, owner, max_loss)
}

// MaxWithdraw0 is a free data retrieval call binding the contract method 0x85b68756.
//
// Solidity: function maxWithdraw(address owner, uint256 max_loss) view returns(uint256)
func (_Yvault300 *Yvault300CallerSession) MaxWithdraw0(owner common.Address, max_loss *big.Int) (*big.Int, error) {
	return _Yvault300.Contract.MaxWithdraw0(&_Yvault300.CallOpts, owner, max_loss)
}

// MaxWithdraw1 is a free data retrieval call binding the contract method 0x65cb6765.
//
// Solidity: function maxWithdraw(address owner, uint256 max_loss, address[] strategies) view returns(uint256)
func (_Yvault300 *Yvault300Caller) MaxWithdraw1(opts *bind.CallOpts, owner common.Address, max_loss *big.Int, strategies []common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "maxWithdraw1", owner, max_loss, strategies)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw1 is a free data retrieval call binding the contract method 0x65cb6765.
//
// Solidity: function maxWithdraw(address owner, uint256 max_loss, address[] strategies) view returns(uint256)
func (_Yvault300 *Yvault300Session) MaxWithdraw1(owner common.Address, max_loss *big.Int, strategies []common.Address) (*big.Int, error) {
	return _Yvault300.Contract.MaxWithdraw1(&_Yvault300.CallOpts, owner, max_loss, strategies)
}

// MaxWithdraw1 is a free data retrieval call binding the contract method 0x65cb6765.
//
// Solidity: function maxWithdraw(address owner, uint256 max_loss, address[] strategies) view returns(uint256)
func (_Yvault300 *Yvault300CallerSession) MaxWithdraw1(owner common.Address, max_loss *big.Int, strategies []common.Address) (*big.Int, error) {
	return _Yvault300.Contract.MaxWithdraw1(&_Yvault300.CallOpts, owner, max_loss, strategies)
}

// MinimumTotalIdle is a free data retrieval call binding the contract method 0x356d6409.
//
// Solidity: function minimum_total_idle() view returns(uint256)
func (_Yvault300 *Yvault300Caller) MinimumTotalIdle(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "minimum_total_idle")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumTotalIdle is a free data retrieval call binding the contract method 0x356d6409.
//
// Solidity: function minimum_total_idle() view returns(uint256)
func (_Yvault300 *Yvault300Session) MinimumTotalIdle() (*big.Int, error) {
	return _Yvault300.Contract.MinimumTotalIdle(&_Yvault300.CallOpts)
}

// MinimumTotalIdle is a free data retrieval call binding the contract method 0x356d6409.
//
// Solidity: function minimum_total_idle() view returns(uint256)
func (_Yvault300 *Yvault300CallerSession) MinimumTotalIdle() (*big.Int, error) {
	return _Yvault300.Contract.MinimumTotalIdle(&_Yvault300.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Yvault300 *Yvault300Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Yvault300 *Yvault300Session) Name() (string, error) {
	return _Yvault300.Contract.Name(&_Yvault300.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Yvault300 *Yvault300CallerSession) Name() (string, error) {
	return _Yvault300.Contract.Name(&_Yvault300.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_Yvault300 *Yvault300Caller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_Yvault300 *Yvault300Session) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Yvault300.Contract.Nonces(&_Yvault300.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_Yvault300 *Yvault300CallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Yvault300.Contract.Nonces(&_Yvault300.CallOpts, arg0)
}

// OpenRoles is a free data retrieval call binding the contract method 0xf3789e45.
//
// Solidity: function open_roles(uint256 arg0) view returns(bool)
func (_Yvault300 *Yvault300Caller) OpenRoles(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "open_roles", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OpenRoles is a free data retrieval call binding the contract method 0xf3789e45.
//
// Solidity: function open_roles(uint256 arg0) view returns(bool)
func (_Yvault300 *Yvault300Session) OpenRoles(arg0 *big.Int) (bool, error) {
	return _Yvault300.Contract.OpenRoles(&_Yvault300.CallOpts, arg0)
}

// OpenRoles is a free data retrieval call binding the contract method 0xf3789e45.
//
// Solidity: function open_roles(uint256 arg0) view returns(bool)
func (_Yvault300 *Yvault300CallerSession) OpenRoles(arg0 *big.Int) (bool, error) {
	return _Yvault300.Contract.OpenRoles(&_Yvault300.CallOpts, arg0)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_Yvault300 *Yvault300Caller) PreviewDeposit(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "previewDeposit", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_Yvault300 *Yvault300Session) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _Yvault300.Contract.PreviewDeposit(&_Yvault300.CallOpts, assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_Yvault300 *Yvault300CallerSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _Yvault300.Contract.PreviewDeposit(&_Yvault300.CallOpts, assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_Yvault300 *Yvault300Caller) PreviewMint(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "previewMint", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_Yvault300 *Yvault300Session) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _Yvault300.Contract.PreviewMint(&_Yvault300.CallOpts, shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_Yvault300 *Yvault300CallerSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _Yvault300.Contract.PreviewMint(&_Yvault300.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_Yvault300 *Yvault300Caller) PreviewRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "previewRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_Yvault300 *Yvault300Session) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _Yvault300.Contract.PreviewRedeem(&_Yvault300.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_Yvault300 *Yvault300CallerSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _Yvault300.Contract.PreviewRedeem(&_Yvault300.CallOpts, shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_Yvault300 *Yvault300Caller) PreviewWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "previewWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_Yvault300 *Yvault300Session) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _Yvault300.Contract.PreviewWithdraw(&_Yvault300.CallOpts, assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_Yvault300 *Yvault300CallerSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _Yvault300.Contract.PreviewWithdraw(&_Yvault300.CallOpts, assets)
}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_Yvault300 *Yvault300Caller) PricePerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "pricePerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_Yvault300 *Yvault300Session) PricePerShare() (*big.Int, error) {
	return _Yvault300.Contract.PricePerShare(&_Yvault300.CallOpts)
}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_Yvault300 *Yvault300CallerSession) PricePerShare() (*big.Int, error) {
	return _Yvault300.Contract.PricePerShare(&_Yvault300.CallOpts)
}

// ProfitMaxUnlockTime is a free data retrieval call binding the contract method 0x0952864e.
//
// Solidity: function profitMaxUnlockTime() view returns(uint256)
func (_Yvault300 *Yvault300Caller) ProfitMaxUnlockTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "profitMaxUnlockTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProfitMaxUnlockTime is a free data retrieval call binding the contract method 0x0952864e.
//
// Solidity: function profitMaxUnlockTime() view returns(uint256)
func (_Yvault300 *Yvault300Session) ProfitMaxUnlockTime() (*big.Int, error) {
	return _Yvault300.Contract.ProfitMaxUnlockTime(&_Yvault300.CallOpts)
}

// ProfitMaxUnlockTime is a free data retrieval call binding the contract method 0x0952864e.
//
// Solidity: function profitMaxUnlockTime() view returns(uint256)
func (_Yvault300 *Yvault300CallerSession) ProfitMaxUnlockTime() (*big.Int, error) {
	return _Yvault300.Contract.ProfitMaxUnlockTime(&_Yvault300.CallOpts)
}

// ProfitUnlockingRate is a free data retrieval call binding the contract method 0x5141eebb.
//
// Solidity: function profitUnlockingRate() view returns(uint256)
func (_Yvault300 *Yvault300Caller) ProfitUnlockingRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "profitUnlockingRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProfitUnlockingRate is a free data retrieval call binding the contract method 0x5141eebb.
//
// Solidity: function profitUnlockingRate() view returns(uint256)
func (_Yvault300 *Yvault300Session) ProfitUnlockingRate() (*big.Int, error) {
	return _Yvault300.Contract.ProfitUnlockingRate(&_Yvault300.CallOpts)
}

// ProfitUnlockingRate is a free data retrieval call binding the contract method 0x5141eebb.
//
// Solidity: function profitUnlockingRate() view returns(uint256)
func (_Yvault300 *Yvault300CallerSession) ProfitUnlockingRate() (*big.Int, error) {
	return _Yvault300.Contract.ProfitUnlockingRate(&_Yvault300.CallOpts)
}

// RoleManager is a free data retrieval call binding the contract method 0x79b98917.
//
// Solidity: function role_manager() view returns(address)
func (_Yvault300 *Yvault300Caller) RoleManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "role_manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoleManager is a free data retrieval call binding the contract method 0x79b98917.
//
// Solidity: function role_manager() view returns(address)
func (_Yvault300 *Yvault300Session) RoleManager() (common.Address, error) {
	return _Yvault300.Contract.RoleManager(&_Yvault300.CallOpts)
}

// RoleManager is a free data retrieval call binding the contract method 0x79b98917.
//
// Solidity: function role_manager() view returns(address)
func (_Yvault300 *Yvault300CallerSession) RoleManager() (common.Address, error) {
	return _Yvault300.Contract.RoleManager(&_Yvault300.CallOpts)
}

// Roles is a free data retrieval call binding the contract method 0x99374642.
//
// Solidity: function roles(address arg0) view returns(uint256)
func (_Yvault300 *Yvault300Caller) Roles(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "roles", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Roles is a free data retrieval call binding the contract method 0x99374642.
//
// Solidity: function roles(address arg0) view returns(uint256)
func (_Yvault300 *Yvault300Session) Roles(arg0 common.Address) (*big.Int, error) {
	return _Yvault300.Contract.Roles(&_Yvault300.CallOpts, arg0)
}

// Roles is a free data retrieval call binding the contract method 0x99374642.
//
// Solidity: function roles(address arg0) view returns(uint256)
func (_Yvault300 *Yvault300CallerSession) Roles(arg0 common.Address) (*big.Int, error) {
	return _Yvault300.Contract.Roles(&_Yvault300.CallOpts, arg0)
}

// Strategies is a free data retrieval call binding the contract method 0x39ebf823.
//
// Solidity: function strategies(address arg0) view returns((uint256,uint256,uint256,uint256))
func (_Yvault300 *Yvault300Caller) Strategies(opts *bind.CallOpts, arg0 common.Address) (Struct0, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "strategies", arg0)

	if err != nil {
		return *new(Struct0), err
	}

	out0 := *abi.ConvertType(out[0], new(Struct0)).(*Struct0)

	return out0, err

}

// Strategies is a free data retrieval call binding the contract method 0x39ebf823.
//
// Solidity: function strategies(address arg0) view returns((uint256,uint256,uint256,uint256))
func (_Yvault300 *Yvault300Session) Strategies(arg0 common.Address) (Struct0, error) {
	return _Yvault300.Contract.Strategies(&_Yvault300.CallOpts, arg0)
}

// Strategies is a free data retrieval call binding the contract method 0x39ebf823.
//
// Solidity: function strategies(address arg0) view returns((uint256,uint256,uint256,uint256))
func (_Yvault300 *Yvault300CallerSession) Strategies(arg0 common.Address) (Struct0, error) {
	return _Yvault300.Contract.Strategies(&_Yvault300.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Yvault300 *Yvault300Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Yvault300 *Yvault300Session) Symbol() (string, error) {
	return _Yvault300.Contract.Symbol(&_Yvault300.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Yvault300 *Yvault300CallerSession) Symbol() (string, error) {
	return _Yvault300.Contract.Symbol(&_Yvault300.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_Yvault300 *Yvault300Caller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_Yvault300 *Yvault300Session) TotalAssets() (*big.Int, error) {
	return _Yvault300.Contract.TotalAssets(&_Yvault300.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_Yvault300 *Yvault300CallerSession) TotalAssets() (*big.Int, error) {
	return _Yvault300.Contract.TotalAssets(&_Yvault300.CallOpts)
}

// TotalDebt is a free data retrieval call binding the contract method 0xfc7b9c18.
//
// Solidity: function totalDebt() view returns(uint256)
func (_Yvault300 *Yvault300Caller) TotalDebt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "totalDebt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDebt is a free data retrieval call binding the contract method 0xfc7b9c18.
//
// Solidity: function totalDebt() view returns(uint256)
func (_Yvault300 *Yvault300Session) TotalDebt() (*big.Int, error) {
	return _Yvault300.Contract.TotalDebt(&_Yvault300.CallOpts)
}

// TotalDebt is a free data retrieval call binding the contract method 0xfc7b9c18.
//
// Solidity: function totalDebt() view returns(uint256)
func (_Yvault300 *Yvault300CallerSession) TotalDebt() (*big.Int, error) {
	return _Yvault300.Contract.TotalDebt(&_Yvault300.CallOpts)
}

// TotalIdle is a free data retrieval call binding the contract method 0x9aa7df94.
//
// Solidity: function totalIdle() view returns(uint256)
func (_Yvault300 *Yvault300Caller) TotalIdle(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "totalIdle")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalIdle is a free data retrieval call binding the contract method 0x9aa7df94.
//
// Solidity: function totalIdle() view returns(uint256)
func (_Yvault300 *Yvault300Session) TotalIdle() (*big.Int, error) {
	return _Yvault300.Contract.TotalIdle(&_Yvault300.CallOpts)
}

// TotalIdle is a free data retrieval call binding the contract method 0x9aa7df94.
//
// Solidity: function totalIdle() view returns(uint256)
func (_Yvault300 *Yvault300CallerSession) TotalIdle() (*big.Int, error) {
	return _Yvault300.Contract.TotalIdle(&_Yvault300.CallOpts)
}

// OtherTotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Yvault300 *Yvault300Caller) OtherTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OtherTotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Yvault300 *Yvault300Session) OtherTotalSupply() (*big.Int, error) {
	return _Yvault300.Contract.OtherTotalSupply(&_Yvault300.CallOpts)
}

// OtherTotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Yvault300 *Yvault300CallerSession) OtherTotalSupply() (*big.Int, error) {
	return _Yvault300.Contract.OtherTotalSupply(&_Yvault300.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x3940e9ee.
//
// Solidity: function total_supply() view returns(uint256)
func (_Yvault300 *Yvault300Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "total_supply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x3940e9ee.
//
// Solidity: function total_supply() view returns(uint256)
func (_Yvault300 *Yvault300Session) TotalSupply() (*big.Int, error) {
	return _Yvault300.Contract.TotalSupply(&_Yvault300.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x3940e9ee.
//
// Solidity: function total_supply() view returns(uint256)
func (_Yvault300 *Yvault300CallerSession) TotalSupply() (*big.Int, error) {
	return _Yvault300.Contract.TotalSupply(&_Yvault300.CallOpts)
}

// UnlockedShares is a free data retrieval call binding the contract method 0xd9a0e97a.
//
// Solidity: function unlockedShares() view returns(uint256)
func (_Yvault300 *Yvault300Caller) UnlockedShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "unlockedShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnlockedShares is a free data retrieval call binding the contract method 0xd9a0e97a.
//
// Solidity: function unlockedShares() view returns(uint256)
func (_Yvault300 *Yvault300Session) UnlockedShares() (*big.Int, error) {
	return _Yvault300.Contract.UnlockedShares(&_Yvault300.CallOpts)
}

// UnlockedShares is a free data retrieval call binding the contract method 0xd9a0e97a.
//
// Solidity: function unlockedShares() view returns(uint256)
func (_Yvault300 *Yvault300CallerSession) UnlockedShares() (*big.Int, error) {
	return _Yvault300.Contract.UnlockedShares(&_Yvault300.CallOpts)
}

// UseDefaultQueue is a free data retrieval call binding the contract method 0x1e56558d.
//
// Solidity: function use_default_queue() view returns(bool)
func (_Yvault300 *Yvault300Caller) UseDefaultQueue(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "use_default_queue")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UseDefaultQueue is a free data retrieval call binding the contract method 0x1e56558d.
//
// Solidity: function use_default_queue() view returns(bool)
func (_Yvault300 *Yvault300Session) UseDefaultQueue() (bool, error) {
	return _Yvault300.Contract.UseDefaultQueue(&_Yvault300.CallOpts)
}

// UseDefaultQueue is a free data retrieval call binding the contract method 0x1e56558d.
//
// Solidity: function use_default_queue() view returns(bool)
func (_Yvault300 *Yvault300CallerSession) UseDefaultQueue() (bool, error) {
	return _Yvault300.Contract.UseDefaultQueue(&_Yvault300.CallOpts)
}

// WithdrawLimitModule is a free data retrieval call binding the contract method 0xf5ba68f3.
//
// Solidity: function withdraw_limit_module() view returns(address)
func (_Yvault300 *Yvault300Caller) WithdrawLimitModule(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yvault300.contract.Call(opts, &out, "withdraw_limit_module")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WithdrawLimitModule is a free data retrieval call binding the contract method 0xf5ba68f3.
//
// Solidity: function withdraw_limit_module() view returns(address)
func (_Yvault300 *Yvault300Session) WithdrawLimitModule() (common.Address, error) {
	return _Yvault300.Contract.WithdrawLimitModule(&_Yvault300.CallOpts)
}

// WithdrawLimitModule is a free data retrieval call binding the contract method 0xf5ba68f3.
//
// Solidity: function withdraw_limit_module() view returns(address)
func (_Yvault300 *Yvault300CallerSession) WithdrawLimitModule() (common.Address, error) {
	return _Yvault300.Contract.WithdrawLimitModule(&_Yvault300.CallOpts)
}

// AcceptRoleManager is a paid mutator transaction binding the contract method 0xf776bf1f.
//
// Solidity: function accept_role_manager() returns()
func (_Yvault300 *Yvault300Transactor) AcceptRoleManager(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yvault300.contract.Transact(opts, "accept_role_manager")
}

// AcceptRoleManager is a paid mutator transaction binding the contract method 0xf776bf1f.
//
// Solidity: function accept_role_manager() returns()
func (_Yvault300 *Yvault300Session) AcceptRoleManager() (*types.Transaction, error) {
	return _Yvault300.Contract.AcceptRoleManager(&_Yvault300.TransactOpts)
}

// AcceptRoleManager is a paid mutator transaction binding the contract method 0xf776bf1f.
//
// Solidity: function accept_role_manager() returns()
func (_Yvault300 *Yvault300TransactorSession) AcceptRoleManager() (*types.Transaction, error) {
	return _Yvault300.Contract.AcceptRoleManager(&_Yvault300.TransactOpts)
}

// AddRole is a paid mutator transaction binding the contract method 0xa97cefa2.
//
// Solidity: function add_role(address account, uint256 role) returns()
func (_Yvault300 *Yvault300Transactor) AddRole(opts *bind.TransactOpts, account common.Address, role *big.Int) (*types.Transaction, error) {
	return _Yvault300.contract.Transact(opts, "add_role", account, role)
}

// AddRole is a paid mutator transaction binding the contract method 0xa97cefa2.
//
// Solidity: function add_role(address account, uint256 role) returns()
func (_Yvault300 *Yvault300Session) AddRole(account common.Address, role *big.Int) (*types.Transaction, error) {
	return _Yvault300.Contract.AddRole(&_Yvault300.TransactOpts, account, role)
}

// AddRole is a paid mutator transaction binding the contract method 0xa97cefa2.
//
// Solidity: function add_role(address account, uint256 role) returns()
func (_Yvault300 *Yvault300TransactorSession) AddRole(account common.Address, role *big.Int) (*types.Transaction, error) {
	return _Yvault300.Contract.AddRole(&_Yvault300.TransactOpts, account, role)
}

// AddStrategy is a paid mutator transaction binding the contract method 0xde7aeb41.
//
// Solidity: function add_strategy(address new_strategy) returns()
func (_Yvault300 *Yvault300Transactor) AddStrategy(opts *bind.TransactOpts, new_strategy common.Address) (*types.Transaction, error) {
	return _Yvault300.contract.Transact(opts, "add_strategy", new_strategy)
}

// AddStrategy is a paid mutator transaction binding the contract method 0xde7aeb41.
//
// Solidity: function add_strategy(address new_strategy) returns()
func (_Yvault300 *Yvault300Session) AddStrategy(new_strategy common.Address) (*types.Transaction, error) {
	return _Yvault300.Contract.AddStrategy(&_Yvault300.TransactOpts, new_strategy)
}

// AddStrategy is a paid mutator transaction binding the contract method 0xde7aeb41.
//
// Solidity: function add_strategy(address new_strategy) returns()
func (_Yvault300 *Yvault300TransactorSession) AddStrategy(new_strategy common.Address) (*types.Transaction, error) {
	return _Yvault300.Contract.AddStrategy(&_Yvault300.TransactOpts, new_strategy)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Yvault300 *Yvault300Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault300.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Yvault300 *Yvault300Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault300.Contract.Approve(&_Yvault300.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Yvault300 *Yvault300TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault300.Contract.Approve(&_Yvault300.TransactOpts, spender, amount)
}

// BuyDebt is a paid mutator transaction binding the contract method 0xe5e91818.
//
// Solidity: function buy_debt(address strategy, uint256 amount) returns()
func (_Yvault300 *Yvault300Transactor) BuyDebt(opts *bind.TransactOpts, strategy common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault300.contract.Transact(opts, "buy_debt", strategy, amount)
}

// BuyDebt is a paid mutator transaction binding the contract method 0xe5e91818.
//
// Solidity: function buy_debt(address strategy, uint256 amount) returns()
func (_Yvault300 *Yvault300Session) BuyDebt(strategy common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault300.Contract.BuyDebt(&_Yvault300.TransactOpts, strategy, amount)
}

// BuyDebt is a paid mutator transaction binding the contract method 0xe5e91818.
//
// Solidity: function buy_debt(address strategy, uint256 amount) returns()
func (_Yvault300 *Yvault300TransactorSession) BuyDebt(strategy common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault300.Contract.BuyDebt(&_Yvault300.TransactOpts, strategy, amount)
}

// CloseOpenRole is a paid mutator transaction binding the contract method 0xd52fe498.
//
// Solidity: function close_open_role(uint256 role) returns()
func (_Yvault300 *Yvault300Transactor) CloseOpenRole(opts *bind.TransactOpts, role *big.Int) (*types.Transaction, error) {
	return _Yvault300.contract.Transact(opts, "close_open_role", role)
}

// CloseOpenRole is a paid mutator transaction binding the contract method 0xd52fe498.
//
// Solidity: function close_open_role(uint256 role) returns()
func (_Yvault300 *Yvault300Session) CloseOpenRole(role *big.Int) (*types.Transaction, error) {
	return _Yvault300.Contract.CloseOpenRole(&_Yvault300.TransactOpts, role)
}

// CloseOpenRole is a paid mutator transaction binding the contract method 0xd52fe498.
//
// Solidity: function close_open_role(uint256 role) returns()
func (_Yvault300 *Yvault300TransactorSession) CloseOpenRole(role *big.Int) (*types.Transaction, error) {
	return _Yvault300.Contract.CloseOpenRole(&_Yvault300.TransactOpts, role)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 amount) returns(bool)
func (_Yvault300 *Yvault300Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault300.contract.Transact(opts, "decreaseAllowance", spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 amount) returns(bool)
func (_Yvault300 *Yvault300Session) DecreaseAllowance(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault300.Contract.DecreaseAllowance(&_Yvault300.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 amount) returns(bool)
func (_Yvault300 *Yvault300TransactorSession) DecreaseAllowance(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault300.Contract.DecreaseAllowance(&_Yvault300.TransactOpts, spender, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_Yvault300 *Yvault300Transactor) Deposit(opts *bind.TransactOpts, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Yvault300.contract.Transact(opts, "deposit", assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_Yvault300 *Yvault300Session) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Yvault300.Contract.Deposit(&_Yvault300.TransactOpts, assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_Yvault300 *Yvault300TransactorSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Yvault300.Contract.Deposit(&_Yvault300.TransactOpts, assets, receiver)
}

// ForceRevokeStrategy is a paid mutator transaction binding the contract method 0xfd129e63.
//
// Solidity: function force_revoke_strategy(address strategy) returns()
func (_Yvault300 *Yvault300Transactor) ForceRevokeStrategy(opts *bind.TransactOpts, strategy common.Address) (*types.Transaction, error) {
	return _Yvault300.contract.Transact(opts, "force_revoke_strategy", strategy)
}

// ForceRevokeStrategy is a paid mutator transaction binding the contract method 0xfd129e63.
//
// Solidity: function force_revoke_strategy(address strategy) returns()
func (_Yvault300 *Yvault300Session) ForceRevokeStrategy(strategy common.Address) (*types.Transaction, error) {
	return _Yvault300.Contract.ForceRevokeStrategy(&_Yvault300.TransactOpts, strategy)
}

// ForceRevokeStrategy is a paid mutator transaction binding the contract method 0xfd129e63.
//
// Solidity: function force_revoke_strategy(address strategy) returns()
func (_Yvault300 *Yvault300TransactorSession) ForceRevokeStrategy(strategy common.Address) (*types.Transaction, error) {
	return _Yvault300.Contract.ForceRevokeStrategy(&_Yvault300.TransactOpts, strategy)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 amount) returns(bool)
func (_Yvault300 *Yvault300Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault300.contract.Transact(opts, "increaseAllowance", spender, amount)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 amount) returns(bool)
func (_Yvault300 *Yvault300Session) IncreaseAllowance(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault300.Contract.IncreaseAllowance(&_Yvault300.TransactOpts, spender, amount)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 amount) returns(bool)
func (_Yvault300 *Yvault300TransactorSession) IncreaseAllowance(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault300.Contract.IncreaseAllowance(&_Yvault300.TransactOpts, spender, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_Yvault300 *Yvault300Transactor) Mint(opts *bind.TransactOpts, shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Yvault300.contract.Transact(opts, "mint", shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_Yvault300 *Yvault300Session) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Yvault300.Contract.Mint(&_Yvault300.TransactOpts, shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_Yvault300 *Yvault300TransactorSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Yvault300.Contract.Mint(&_Yvault300.TransactOpts, shares, receiver)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns(bool)
func (_Yvault300 *Yvault300Transactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Yvault300.contract.Transact(opts, "permit", owner, spender, amount, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns(bool)
func (_Yvault300 *Yvault300Session) Permit(owner common.Address, spender common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Yvault300.Contract.Permit(&_Yvault300.TransactOpts, owner, spender, amount, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns(bool)
func (_Yvault300 *Yvault300TransactorSession) Permit(owner common.Address, spender common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Yvault300.Contract.Permit(&_Yvault300.TransactOpts, owner, spender, amount, deadline, v, r, s)
}

// ProcessReport is a paid mutator transaction binding the contract method 0x6ec2b8d4.
//
// Solidity: function process_report(address strategy) returns(uint256, uint256)
func (_Yvault300 *Yvault300Transactor) ProcessReport(opts *bind.TransactOpts, strategy common.Address) (*types.Transaction, error) {
	return _Yvault300.contract.Transact(opts, "process_report", strategy)
}

// ProcessReport is a paid mutator transaction binding the contract method 0x6ec2b8d4.
//
// Solidity: function process_report(address strategy) returns(uint256, uint256)
func (_Yvault300 *Yvault300Session) ProcessReport(strategy common.Address) (*types.Transaction, error) {
	return _Yvault300.Contract.ProcessReport(&_Yvault300.TransactOpts, strategy)
}

// ProcessReport is a paid mutator transaction binding the contract method 0x6ec2b8d4.
//
// Solidity: function process_report(address strategy) returns(uint256, uint256)
func (_Yvault300 *Yvault300TransactorSession) ProcessReport(strategy common.Address) (*types.Transaction, error) {
	return _Yvault300.Contract.ProcessReport(&_Yvault300.TransactOpts, strategy)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_Yvault300 *Yvault300Transactor) Redeem(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Yvault300.contract.Transact(opts, "redeem", shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_Yvault300 *Yvault300Session) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Yvault300.Contract.Redeem(&_Yvault300.TransactOpts, shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_Yvault300 *Yvault300TransactorSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Yvault300.Contract.Redeem(&_Yvault300.TransactOpts, shares, receiver, owner)
}

// Redeem0 is a paid mutator transaction binding the contract method 0x9f40a7b3.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner, uint256 max_loss) returns(uint256)
func (_Yvault300 *Yvault300Transactor) Redeem0(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int) (*types.Transaction, error) {
	return _Yvault300.contract.Transact(opts, "redeem0", shares, receiver, owner, max_loss)
}

// Redeem0 is a paid mutator transaction binding the contract method 0x9f40a7b3.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner, uint256 max_loss) returns(uint256)
func (_Yvault300 *Yvault300Session) Redeem0(shares *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int) (*types.Transaction, error) {
	return _Yvault300.Contract.Redeem0(&_Yvault300.TransactOpts, shares, receiver, owner, max_loss)
}

// Redeem0 is a paid mutator transaction binding the contract method 0x9f40a7b3.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner, uint256 max_loss) returns(uint256)
func (_Yvault300 *Yvault300TransactorSession) Redeem0(shares *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int) (*types.Transaction, error) {
	return _Yvault300.Contract.Redeem0(&_Yvault300.TransactOpts, shares, receiver, owner, max_loss)
}

// Redeem1 is a paid mutator transaction binding the contract method 0x06580f2d.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner, uint256 max_loss, address[] strategies) returns(uint256)
func (_Yvault300 *Yvault300Transactor) Redeem1(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int, strategies []common.Address) (*types.Transaction, error) {
	return _Yvault300.contract.Transact(opts, "redeem1", shares, receiver, owner, max_loss, strategies)
}

// Redeem1 is a paid mutator transaction binding the contract method 0x06580f2d.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner, uint256 max_loss, address[] strategies) returns(uint256)
func (_Yvault300 *Yvault300Session) Redeem1(shares *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int, strategies []common.Address) (*types.Transaction, error) {
	return _Yvault300.Contract.Redeem1(&_Yvault300.TransactOpts, shares, receiver, owner, max_loss, strategies)
}

// Redeem1 is a paid mutator transaction binding the contract method 0x06580f2d.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner, uint256 max_loss, address[] strategies) returns(uint256)
func (_Yvault300 *Yvault300TransactorSession) Redeem1(shares *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int, strategies []common.Address) (*types.Transaction, error) {
	return _Yvault300.Contract.Redeem1(&_Yvault300.TransactOpts, shares, receiver, owner, max_loss, strategies)
}

// RemoveRole is a paid mutator transaction binding the contract method 0xe2bf56dd.
//
// Solidity: function remove_role(address account, uint256 role) returns()
func (_Yvault300 *Yvault300Transactor) RemoveRole(opts *bind.TransactOpts, account common.Address, role *big.Int) (*types.Transaction, error) {
	return _Yvault300.contract.Transact(opts, "remove_role", account, role)
}

// RemoveRole is a paid mutator transaction binding the contract method 0xe2bf56dd.
//
// Solidity: function remove_role(address account, uint256 role) returns()
func (_Yvault300 *Yvault300Session) RemoveRole(account common.Address, role *big.Int) (*types.Transaction, error) {
	return _Yvault300.Contract.RemoveRole(&_Yvault300.TransactOpts, account, role)
}

// RemoveRole is a paid mutator transaction binding the contract method 0xe2bf56dd.
//
// Solidity: function remove_role(address account, uint256 role) returns()
func (_Yvault300 *Yvault300TransactorSession) RemoveRole(account common.Address, role *big.Int) (*types.Transaction, error) {
	return _Yvault300.Contract.RemoveRole(&_Yvault300.TransactOpts, account, role)
}

// RevokeStrategy is a paid mutator transaction binding the contract method 0x577db316.
//
// Solidity: function revoke_strategy(address strategy) returns()
func (_Yvault300 *Yvault300Transactor) RevokeStrategy(opts *bind.TransactOpts, strategy common.Address) (*types.Transaction, error) {
	return _Yvault300.contract.Transact(opts, "revoke_strategy", strategy)
}

// RevokeStrategy is a paid mutator transaction binding the contract method 0x577db316.
//
// Solidity: function revoke_strategy(address strategy) returns()
func (_Yvault300 *Yvault300Session) RevokeStrategy(strategy common.Address) (*types.Transaction, error) {
	return _Yvault300.Contract.RevokeStrategy(&_Yvault300.TransactOpts, strategy)
}

// RevokeStrategy is a paid mutator transaction binding the contract method 0x577db316.
//
// Solidity: function revoke_strategy(address strategy) returns()
func (_Yvault300 *Yvault300TransactorSession) RevokeStrategy(strategy common.Address) (*types.Transaction, error) {
	return _Yvault300.Contract.RevokeStrategy(&_Yvault300.TransactOpts, strategy)
}

// SetProfitMaxUnlockTime is a paid mutator transaction binding the contract method 0xdf69b22a.
//
// Solidity: function setProfitMaxUnlockTime(uint256 new_profit_max_unlock_time) returns()
func (_Yvault300 *Yvault300Transactor) SetProfitMaxUnlockTime(opts *bind.TransactOpts, new_profit_max_unlock_time *big.Int) (*types.Transaction, error) {
	return _Yvault300.contract.Transact(opts, "setProfitMaxUnlockTime", new_profit_max_unlock_time)
}

// SetProfitMaxUnlockTime is a paid mutator transaction binding the contract method 0xdf69b22a.
//
// Solidity: function setProfitMaxUnlockTime(uint256 new_profit_max_unlock_time) returns()
func (_Yvault300 *Yvault300Session) SetProfitMaxUnlockTime(new_profit_max_unlock_time *big.Int) (*types.Transaction, error) {
	return _Yvault300.Contract.SetProfitMaxUnlockTime(&_Yvault300.TransactOpts, new_profit_max_unlock_time)
}

// SetProfitMaxUnlockTime is a paid mutator transaction binding the contract method 0xdf69b22a.
//
// Solidity: function setProfitMaxUnlockTime(uint256 new_profit_max_unlock_time) returns()
func (_Yvault300 *Yvault300TransactorSession) SetProfitMaxUnlockTime(new_profit_max_unlock_time *big.Int) (*types.Transaction, error) {
	return _Yvault300.Contract.SetProfitMaxUnlockTime(&_Yvault300.TransactOpts, new_profit_max_unlock_time)
}

// SetAccountant is a paid mutator transaction binding the contract method 0x71da8a8d.
//
// Solidity: function set_accountant(address new_accountant) returns()
func (_Yvault300 *Yvault300Transactor) SetAccountant(opts *bind.TransactOpts, new_accountant common.Address) (*types.Transaction, error) {
	return _Yvault300.contract.Transact(opts, "set_accountant", new_accountant)
}

// SetAccountant is a paid mutator transaction binding the contract method 0x71da8a8d.
//
// Solidity: function set_accountant(address new_accountant) returns()
func (_Yvault300 *Yvault300Session) SetAccountant(new_accountant common.Address) (*types.Transaction, error) {
	return _Yvault300.Contract.SetAccountant(&_Yvault300.TransactOpts, new_accountant)
}

// SetAccountant is a paid mutator transaction binding the contract method 0x71da8a8d.
//
// Solidity: function set_accountant(address new_accountant) returns()
func (_Yvault300 *Yvault300TransactorSession) SetAccountant(new_accountant common.Address) (*types.Transaction, error) {
	return _Yvault300.Contract.SetAccountant(&_Yvault300.TransactOpts, new_accountant)
}

// SetDefaultQueue is a paid mutator transaction binding the contract method 0x2d9caa4e.
//
// Solidity: function set_default_queue(address[] new_default_queue) returns()
func (_Yvault300 *Yvault300Transactor) SetDefaultQueue(opts *bind.TransactOpts, new_default_queue []common.Address) (*types.Transaction, error) {
	return _Yvault300.contract.Transact(opts, "set_default_queue", new_default_queue)
}

// SetDefaultQueue is a paid mutator transaction binding the contract method 0x2d9caa4e.
//
// Solidity: function set_default_queue(address[] new_default_queue) returns()
func (_Yvault300 *Yvault300Session) SetDefaultQueue(new_default_queue []common.Address) (*types.Transaction, error) {
	return _Yvault300.Contract.SetDefaultQueue(&_Yvault300.TransactOpts, new_default_queue)
}

// SetDefaultQueue is a paid mutator transaction binding the contract method 0x2d9caa4e.
//
// Solidity: function set_default_queue(address[] new_default_queue) returns()
func (_Yvault300 *Yvault300TransactorSession) SetDefaultQueue(new_default_queue []common.Address) (*types.Transaction, error) {
	return _Yvault300.Contract.SetDefaultQueue(&_Yvault300.TransactOpts, new_default_queue)
}

// SetDepositLimit is a paid mutator transaction binding the contract method 0x6fe01d1e.
//
// Solidity: function set_deposit_limit(uint256 deposit_limit) returns()
func (_Yvault300 *Yvault300Transactor) SetDepositLimit(opts *bind.TransactOpts, deposit_limit *big.Int) (*types.Transaction, error) {
	return _Yvault300.contract.Transact(opts, "set_deposit_limit", deposit_limit)
}

// SetDepositLimit is a paid mutator transaction binding the contract method 0x6fe01d1e.
//
// Solidity: function set_deposit_limit(uint256 deposit_limit) returns()
func (_Yvault300 *Yvault300Session) SetDepositLimit(deposit_limit *big.Int) (*types.Transaction, error) {
	return _Yvault300.Contract.SetDepositLimit(&_Yvault300.TransactOpts, deposit_limit)
}

// SetDepositLimit is a paid mutator transaction binding the contract method 0x6fe01d1e.
//
// Solidity: function set_deposit_limit(uint256 deposit_limit) returns()
func (_Yvault300 *Yvault300TransactorSession) SetDepositLimit(deposit_limit *big.Int) (*types.Transaction, error) {
	return _Yvault300.Contract.SetDepositLimit(&_Yvault300.TransactOpts, deposit_limit)
}

// SetDepositLimitModule is a paid mutator transaction binding the contract method 0xbb435466.
//
// Solidity: function set_deposit_limit_module(address deposit_limit_module) returns()
func (_Yvault300 *Yvault300Transactor) SetDepositLimitModule(opts *bind.TransactOpts, deposit_limit_module common.Address) (*types.Transaction, error) {
	return _Yvault300.contract.Transact(opts, "set_deposit_limit_module", deposit_limit_module)
}

// SetDepositLimitModule is a paid mutator transaction binding the contract method 0xbb435466.
//
// Solidity: function set_deposit_limit_module(address deposit_limit_module) returns()
func (_Yvault300 *Yvault300Session) SetDepositLimitModule(deposit_limit_module common.Address) (*types.Transaction, error) {
	return _Yvault300.Contract.SetDepositLimitModule(&_Yvault300.TransactOpts, deposit_limit_module)
}

// SetDepositLimitModule is a paid mutator transaction binding the contract method 0xbb435466.
//
// Solidity: function set_deposit_limit_module(address deposit_limit_module) returns()
func (_Yvault300 *Yvault300TransactorSession) SetDepositLimitModule(deposit_limit_module common.Address) (*types.Transaction, error) {
	return _Yvault300.Contract.SetDepositLimitModule(&_Yvault300.TransactOpts, deposit_limit_module)
}

// SetMinimumTotalIdle is a paid mutator transaction binding the contract method 0xbdd81c01.
//
// Solidity: function set_minimum_total_idle(uint256 minimum_total_idle) returns()
func (_Yvault300 *Yvault300Transactor) SetMinimumTotalIdle(opts *bind.TransactOpts, minimum_total_idle *big.Int) (*types.Transaction, error) {
	return _Yvault300.contract.Transact(opts, "set_minimum_total_idle", minimum_total_idle)
}

// SetMinimumTotalIdle is a paid mutator transaction binding the contract method 0xbdd81c01.
//
// Solidity: function set_minimum_total_idle(uint256 minimum_total_idle) returns()
func (_Yvault300 *Yvault300Session) SetMinimumTotalIdle(minimum_total_idle *big.Int) (*types.Transaction, error) {
	return _Yvault300.Contract.SetMinimumTotalIdle(&_Yvault300.TransactOpts, minimum_total_idle)
}

// SetMinimumTotalIdle is a paid mutator transaction binding the contract method 0xbdd81c01.
//
// Solidity: function set_minimum_total_idle(uint256 minimum_total_idle) returns()
func (_Yvault300 *Yvault300TransactorSession) SetMinimumTotalIdle(minimum_total_idle *big.Int) (*types.Transaction, error) {
	return _Yvault300.Contract.SetMinimumTotalIdle(&_Yvault300.TransactOpts, minimum_total_idle)
}

// SetOpenRole is a paid mutator transaction binding the contract method 0x0b10dd80.
//
// Solidity: function set_open_role(uint256 role) returns()
func (_Yvault300 *Yvault300Transactor) SetOpenRole(opts *bind.TransactOpts, role *big.Int) (*types.Transaction, error) {
	return _Yvault300.contract.Transact(opts, "set_open_role", role)
}

// SetOpenRole is a paid mutator transaction binding the contract method 0x0b10dd80.
//
// Solidity: function set_open_role(uint256 role) returns()
func (_Yvault300 *Yvault300Session) SetOpenRole(role *big.Int) (*types.Transaction, error) {
	return _Yvault300.Contract.SetOpenRole(&_Yvault300.TransactOpts, role)
}

// SetOpenRole is a paid mutator transaction binding the contract method 0x0b10dd80.
//
// Solidity: function set_open_role(uint256 role) returns()
func (_Yvault300 *Yvault300TransactorSession) SetOpenRole(role *big.Int) (*types.Transaction, error) {
	return _Yvault300.Contract.SetOpenRole(&_Yvault300.TransactOpts, role)
}

// SetRole is a paid mutator transaction binding the contract method 0x2cf7fd85.
//
// Solidity: function set_role(address account, uint256 role) returns()
func (_Yvault300 *Yvault300Transactor) SetRole(opts *bind.TransactOpts, account common.Address, role *big.Int) (*types.Transaction, error) {
	return _Yvault300.contract.Transact(opts, "set_role", account, role)
}

// SetRole is a paid mutator transaction binding the contract method 0x2cf7fd85.
//
// Solidity: function set_role(address account, uint256 role) returns()
func (_Yvault300 *Yvault300Session) SetRole(account common.Address, role *big.Int) (*types.Transaction, error) {
	return _Yvault300.Contract.SetRole(&_Yvault300.TransactOpts, account, role)
}

// SetRole is a paid mutator transaction binding the contract method 0x2cf7fd85.
//
// Solidity: function set_role(address account, uint256 role) returns()
func (_Yvault300 *Yvault300TransactorSession) SetRole(account common.Address, role *big.Int) (*types.Transaction, error) {
	return _Yvault300.Contract.SetRole(&_Yvault300.TransactOpts, account, role)
}

// SetUseDefaultQueue is a paid mutator transaction binding the contract method 0x29c8a33b.
//
// Solidity: function set_use_default_queue(bool use_default_queue) returns()
func (_Yvault300 *Yvault300Transactor) SetUseDefaultQueue(opts *bind.TransactOpts, use_default_queue bool) (*types.Transaction, error) {
	return _Yvault300.contract.Transact(opts, "set_use_default_queue", use_default_queue)
}

// SetUseDefaultQueue is a paid mutator transaction binding the contract method 0x29c8a33b.
//
// Solidity: function set_use_default_queue(bool use_default_queue) returns()
func (_Yvault300 *Yvault300Session) SetUseDefaultQueue(use_default_queue bool) (*types.Transaction, error) {
	return _Yvault300.Contract.SetUseDefaultQueue(&_Yvault300.TransactOpts, use_default_queue)
}

// SetUseDefaultQueue is a paid mutator transaction binding the contract method 0x29c8a33b.
//
// Solidity: function set_use_default_queue(bool use_default_queue) returns()
func (_Yvault300 *Yvault300TransactorSession) SetUseDefaultQueue(use_default_queue bool) (*types.Transaction, error) {
	return _Yvault300.Contract.SetUseDefaultQueue(&_Yvault300.TransactOpts, use_default_queue)
}

// SetWithdrawLimitModule is a paid mutator transaction binding the contract method 0x7b675894.
//
// Solidity: function set_withdraw_limit_module(address withdraw_limit_module) returns()
func (_Yvault300 *Yvault300Transactor) SetWithdrawLimitModule(opts *bind.TransactOpts, withdraw_limit_module common.Address) (*types.Transaction, error) {
	return _Yvault300.contract.Transact(opts, "set_withdraw_limit_module", withdraw_limit_module)
}

// SetWithdrawLimitModule is a paid mutator transaction binding the contract method 0x7b675894.
//
// Solidity: function set_withdraw_limit_module(address withdraw_limit_module) returns()
func (_Yvault300 *Yvault300Session) SetWithdrawLimitModule(withdraw_limit_module common.Address) (*types.Transaction, error) {
	return _Yvault300.Contract.SetWithdrawLimitModule(&_Yvault300.TransactOpts, withdraw_limit_module)
}

// SetWithdrawLimitModule is a paid mutator transaction binding the contract method 0x7b675894.
//
// Solidity: function set_withdraw_limit_module(address withdraw_limit_module) returns()
func (_Yvault300 *Yvault300TransactorSession) SetWithdrawLimitModule(withdraw_limit_module common.Address) (*types.Transaction, error) {
	return _Yvault300.Contract.SetWithdrawLimitModule(&_Yvault300.TransactOpts, withdraw_limit_module)
}

// ShutdownVault is a paid mutator transaction binding the contract method 0x36a55450.
//
// Solidity: function shutdown_vault() returns()
func (_Yvault300 *Yvault300Transactor) ShutdownVault(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yvault300.contract.Transact(opts, "shutdown_vault")
}

// ShutdownVault is a paid mutator transaction binding the contract method 0x36a55450.
//
// Solidity: function shutdown_vault() returns()
func (_Yvault300 *Yvault300Session) ShutdownVault() (*types.Transaction, error) {
	return _Yvault300.Contract.ShutdownVault(&_Yvault300.TransactOpts)
}

// ShutdownVault is a paid mutator transaction binding the contract method 0x36a55450.
//
// Solidity: function shutdown_vault() returns()
func (_Yvault300 *Yvault300TransactorSession) ShutdownVault() (*types.Transaction, error) {
	return _Yvault300.Contract.ShutdownVault(&_Yvault300.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 amount) returns(bool)
func (_Yvault300 *Yvault300Transactor) Transfer(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault300.contract.Transact(opts, "transfer", receiver, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 amount) returns(bool)
func (_Yvault300 *Yvault300Session) Transfer(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault300.Contract.Transfer(&_Yvault300.TransactOpts, receiver, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 amount) returns(bool)
func (_Yvault300 *Yvault300TransactorSession) Transfer(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault300.Contract.Transfer(&_Yvault300.TransactOpts, receiver, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address receiver, uint256 amount) returns(bool)
func (_Yvault300 *Yvault300Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault300.contract.Transact(opts, "transferFrom", sender, receiver, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address receiver, uint256 amount) returns(bool)
func (_Yvault300 *Yvault300Session) TransferFrom(sender common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault300.Contract.TransferFrom(&_Yvault300.TransactOpts, sender, receiver, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address receiver, uint256 amount) returns(bool)
func (_Yvault300 *Yvault300TransactorSession) TransferFrom(sender common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yvault300.Contract.TransferFrom(&_Yvault300.TransactOpts, sender, receiver, amount)
}

// TransferRoleManager is a paid mutator transaction binding the contract method 0xef54cefd.
//
// Solidity: function transfer_role_manager(address role_manager) returns()
func (_Yvault300 *Yvault300Transactor) TransferRoleManager(opts *bind.TransactOpts, role_manager common.Address) (*types.Transaction, error) {
	return _Yvault300.contract.Transact(opts, "transfer_role_manager", role_manager)
}

// TransferRoleManager is a paid mutator transaction binding the contract method 0xef54cefd.
//
// Solidity: function transfer_role_manager(address role_manager) returns()
func (_Yvault300 *Yvault300Session) TransferRoleManager(role_manager common.Address) (*types.Transaction, error) {
	return _Yvault300.Contract.TransferRoleManager(&_Yvault300.TransactOpts, role_manager)
}

// TransferRoleManager is a paid mutator transaction binding the contract method 0xef54cefd.
//
// Solidity: function transfer_role_manager(address role_manager) returns()
func (_Yvault300 *Yvault300TransactorSession) TransferRoleManager(role_manager common.Address) (*types.Transaction, error) {
	return _Yvault300.Contract.TransferRoleManager(&_Yvault300.TransactOpts, role_manager)
}

// UpdateDebt is a paid mutator transaction binding the contract method 0x0aeebf55.
//
// Solidity: function update_debt(address strategy, uint256 target_debt) returns(uint256)
func (_Yvault300 *Yvault300Transactor) UpdateDebt(opts *bind.TransactOpts, strategy common.Address, target_debt *big.Int) (*types.Transaction, error) {
	return _Yvault300.contract.Transact(opts, "update_debt", strategy, target_debt)
}

// UpdateDebt is a paid mutator transaction binding the contract method 0x0aeebf55.
//
// Solidity: function update_debt(address strategy, uint256 target_debt) returns(uint256)
func (_Yvault300 *Yvault300Session) UpdateDebt(strategy common.Address, target_debt *big.Int) (*types.Transaction, error) {
	return _Yvault300.Contract.UpdateDebt(&_Yvault300.TransactOpts, strategy, target_debt)
}

// UpdateDebt is a paid mutator transaction binding the contract method 0x0aeebf55.
//
// Solidity: function update_debt(address strategy, uint256 target_debt) returns(uint256)
func (_Yvault300 *Yvault300TransactorSession) UpdateDebt(strategy common.Address, target_debt *big.Int) (*types.Transaction, error) {
	return _Yvault300.Contract.UpdateDebt(&_Yvault300.TransactOpts, strategy, target_debt)
}

// UpdateMaxDebtForStrategy is a paid mutator transaction binding the contract method 0xb9ddcd68.
//
// Solidity: function update_max_debt_for_strategy(address strategy, uint256 new_max_debt) returns()
func (_Yvault300 *Yvault300Transactor) UpdateMaxDebtForStrategy(opts *bind.TransactOpts, strategy common.Address, new_max_debt *big.Int) (*types.Transaction, error) {
	return _Yvault300.contract.Transact(opts, "update_max_debt_for_strategy", strategy, new_max_debt)
}

// UpdateMaxDebtForStrategy is a paid mutator transaction binding the contract method 0xb9ddcd68.
//
// Solidity: function update_max_debt_for_strategy(address strategy, uint256 new_max_debt) returns()
func (_Yvault300 *Yvault300Session) UpdateMaxDebtForStrategy(strategy common.Address, new_max_debt *big.Int) (*types.Transaction, error) {
	return _Yvault300.Contract.UpdateMaxDebtForStrategy(&_Yvault300.TransactOpts, strategy, new_max_debt)
}

// UpdateMaxDebtForStrategy is a paid mutator transaction binding the contract method 0xb9ddcd68.
//
// Solidity: function update_max_debt_for_strategy(address strategy, uint256 new_max_debt) returns()
func (_Yvault300 *Yvault300TransactorSession) UpdateMaxDebtForStrategy(strategy common.Address, new_max_debt *big.Int) (*types.Transaction, error) {
	return _Yvault300.Contract.UpdateMaxDebtForStrategy(&_Yvault300.TransactOpts, strategy, new_max_debt)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_Yvault300 *Yvault300Transactor) Withdraw(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Yvault300.contract.Transact(opts, "withdraw", assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_Yvault300 *Yvault300Session) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Yvault300.Contract.Withdraw(&_Yvault300.TransactOpts, assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_Yvault300 *Yvault300TransactorSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Yvault300.Contract.Withdraw(&_Yvault300.TransactOpts, assets, receiver, owner)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xa318c1a4.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner, uint256 max_loss) returns(uint256)
func (_Yvault300 *Yvault300Transactor) Withdraw0(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int) (*types.Transaction, error) {
	return _Yvault300.contract.Transact(opts, "withdraw0", assets, receiver, owner, max_loss)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xa318c1a4.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner, uint256 max_loss) returns(uint256)
func (_Yvault300 *Yvault300Session) Withdraw0(assets *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int) (*types.Transaction, error) {
	return _Yvault300.Contract.Withdraw0(&_Yvault300.TransactOpts, assets, receiver, owner, max_loss)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xa318c1a4.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner, uint256 max_loss) returns(uint256)
func (_Yvault300 *Yvault300TransactorSession) Withdraw0(assets *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int) (*types.Transaction, error) {
	return _Yvault300.Contract.Withdraw0(&_Yvault300.TransactOpts, assets, receiver, owner, max_loss)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0xd81a09f6.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner, uint256 max_loss, address[] strategies) returns(uint256)
func (_Yvault300 *Yvault300Transactor) Withdraw1(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int, strategies []common.Address) (*types.Transaction, error) {
	return _Yvault300.contract.Transact(opts, "withdraw1", assets, receiver, owner, max_loss, strategies)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0xd81a09f6.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner, uint256 max_loss, address[] strategies) returns(uint256)
func (_Yvault300 *Yvault300Session) Withdraw1(assets *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int, strategies []common.Address) (*types.Transaction, error) {
	return _Yvault300.Contract.Withdraw1(&_Yvault300.TransactOpts, assets, receiver, owner, max_loss, strategies)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0xd81a09f6.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner, uint256 max_loss, address[] strategies) returns(uint256)
func (_Yvault300 *Yvault300TransactorSession) Withdraw1(assets *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int, strategies []common.Address) (*types.Transaction, error) {
	return _Yvault300.Contract.Withdraw1(&_Yvault300.TransactOpts, assets, receiver, owner, max_loss, strategies)
}

// Yvault300ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Yvault300 contract.
type Yvault300ApprovalIterator struct {
	Event *Yvault300Approval // Event containing the contract specifics and raw log

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
func (it *Yvault300ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault300Approval)
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
		it.Event = new(Yvault300Approval)
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
func (it *Yvault300ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault300ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault300Approval represents a Approval event raised by the Yvault300 contract.
type Yvault300Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Yvault300 *Yvault300Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Yvault300ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Yvault300.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Yvault300ApprovalIterator{contract: _Yvault300.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Yvault300 *Yvault300Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Yvault300Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Yvault300.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault300Approval)
				if err := _Yvault300.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Yvault300 *Yvault300Filterer) ParseApproval(log types.Log) (*Yvault300Approval, error) {
	event := new(Yvault300Approval)
	if err := _Yvault300.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault300DebtPurchasedIterator is returned from FilterDebtPurchased and is used to iterate over the raw logs and unpacked data for DebtPurchased events raised by the Yvault300 contract.
type Yvault300DebtPurchasedIterator struct {
	Event *Yvault300DebtPurchased // Event containing the contract specifics and raw log

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
func (it *Yvault300DebtPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault300DebtPurchased)
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
		it.Event = new(Yvault300DebtPurchased)
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
func (it *Yvault300DebtPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault300DebtPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault300DebtPurchased represents a DebtPurchased event raised by the Yvault300 contract.
type Yvault300DebtPurchased struct {
	Strategy common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDebtPurchased is a free log retrieval operation binding the contract event 0xe94e7f88819f66c19b097748cb754149f63b1a176ed425dee1f1ee933e6d09b0.
//
// Solidity: event DebtPurchased(address indexed strategy, uint256 amount)
func (_Yvault300 *Yvault300Filterer) FilterDebtPurchased(opts *bind.FilterOpts, strategy []common.Address) (*Yvault300DebtPurchasedIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault300.contract.FilterLogs(opts, "DebtPurchased", strategyRule)
	if err != nil {
		return nil, err
	}
	return &Yvault300DebtPurchasedIterator{contract: _Yvault300.contract, event: "DebtPurchased", logs: logs, sub: sub}, nil
}

// WatchDebtPurchased is a free log subscription operation binding the contract event 0xe94e7f88819f66c19b097748cb754149f63b1a176ed425dee1f1ee933e6d09b0.
//
// Solidity: event DebtPurchased(address indexed strategy, uint256 amount)
func (_Yvault300 *Yvault300Filterer) WatchDebtPurchased(opts *bind.WatchOpts, sink chan<- *Yvault300DebtPurchased, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault300.contract.WatchLogs(opts, "DebtPurchased", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault300DebtPurchased)
				if err := _Yvault300.contract.UnpackLog(event, "DebtPurchased", log); err != nil {
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

// ParseDebtPurchased is a log parse operation binding the contract event 0xe94e7f88819f66c19b097748cb754149f63b1a176ed425dee1f1ee933e6d09b0.
//
// Solidity: event DebtPurchased(address indexed strategy, uint256 amount)
func (_Yvault300 *Yvault300Filterer) ParseDebtPurchased(log types.Log) (*Yvault300DebtPurchased, error) {
	event := new(Yvault300DebtPurchased)
	if err := _Yvault300.contract.UnpackLog(event, "DebtPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault300DebtUpdatedIterator is returned from FilterDebtUpdated and is used to iterate over the raw logs and unpacked data for DebtUpdated events raised by the Yvault300 contract.
type Yvault300DebtUpdatedIterator struct {
	Event *Yvault300DebtUpdated // Event containing the contract specifics and raw log

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
func (it *Yvault300DebtUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault300DebtUpdated)
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
		it.Event = new(Yvault300DebtUpdated)
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
func (it *Yvault300DebtUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault300DebtUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault300DebtUpdated represents a DebtUpdated event raised by the Yvault300 contract.
type Yvault300DebtUpdated struct {
	Strategy    common.Address
	CurrentDebt *big.Int
	NewDebt     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDebtUpdated is a free log retrieval operation binding the contract event 0x5e2b8821ad6e0e26207e0cb4d242d07eeb1cbb1cfd853e645bdcd27cc5484f95.
//
// Solidity: event DebtUpdated(address indexed strategy, uint256 current_debt, uint256 new_debt)
func (_Yvault300 *Yvault300Filterer) FilterDebtUpdated(opts *bind.FilterOpts, strategy []common.Address) (*Yvault300DebtUpdatedIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault300.contract.FilterLogs(opts, "DebtUpdated", strategyRule)
	if err != nil {
		return nil, err
	}
	return &Yvault300DebtUpdatedIterator{contract: _Yvault300.contract, event: "DebtUpdated", logs: logs, sub: sub}, nil
}

// WatchDebtUpdated is a free log subscription operation binding the contract event 0x5e2b8821ad6e0e26207e0cb4d242d07eeb1cbb1cfd853e645bdcd27cc5484f95.
//
// Solidity: event DebtUpdated(address indexed strategy, uint256 current_debt, uint256 new_debt)
func (_Yvault300 *Yvault300Filterer) WatchDebtUpdated(opts *bind.WatchOpts, sink chan<- *Yvault300DebtUpdated, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault300.contract.WatchLogs(opts, "DebtUpdated", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault300DebtUpdated)
				if err := _Yvault300.contract.UnpackLog(event, "DebtUpdated", log); err != nil {
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

// ParseDebtUpdated is a log parse operation binding the contract event 0x5e2b8821ad6e0e26207e0cb4d242d07eeb1cbb1cfd853e645bdcd27cc5484f95.
//
// Solidity: event DebtUpdated(address indexed strategy, uint256 current_debt, uint256 new_debt)
func (_Yvault300 *Yvault300Filterer) ParseDebtUpdated(log types.Log) (*Yvault300DebtUpdated, error) {
	event := new(Yvault300DebtUpdated)
	if err := _Yvault300.contract.UnpackLog(event, "DebtUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault300DepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Yvault300 contract.
type Yvault300DepositIterator struct {
	Event *Yvault300Deposit // Event containing the contract specifics and raw log

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
func (it *Yvault300DepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault300Deposit)
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
		it.Event = new(Yvault300Deposit)
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
func (it *Yvault300DepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault300DepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault300Deposit represents a Deposit event raised by the Yvault300 contract.
type Yvault300Deposit struct {
	Sender common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_Yvault300 *Yvault300Filterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, owner []common.Address) (*Yvault300DepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Yvault300.contract.FilterLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &Yvault300DepositIterator{contract: _Yvault300.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_Yvault300 *Yvault300Filterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *Yvault300Deposit, sender []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Yvault300.contract.WatchLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault300Deposit)
				if err := _Yvault300.contract.UnpackLog(event, "Deposit", log); err != nil {
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
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_Yvault300 *Yvault300Filterer) ParseDeposit(log types.Log) (*Yvault300Deposit, error) {
	event := new(Yvault300Deposit)
	if err := _Yvault300.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault300RoleSetIterator is returned from FilterRoleSet and is used to iterate over the raw logs and unpacked data for RoleSet events raised by the Yvault300 contract.
type Yvault300RoleSetIterator struct {
	Event *Yvault300RoleSet // Event containing the contract specifics and raw log

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
func (it *Yvault300RoleSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault300RoleSet)
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
		it.Event = new(Yvault300RoleSet)
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
func (it *Yvault300RoleSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault300RoleSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault300RoleSet represents a RoleSet event raised by the Yvault300 contract.
type Yvault300RoleSet struct {
	Account common.Address
	Role    *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleSet is a free log retrieval operation binding the contract event 0x78557646b1d8efa2cd49740d66df5aca39eb610ca8ca0e1ccac08979b6b2c46e.
//
// Solidity: event RoleSet(address indexed account, uint256 indexed role)
func (_Yvault300 *Yvault300Filterer) FilterRoleSet(opts *bind.FilterOpts, account []common.Address, role []*big.Int) (*Yvault300RoleSetIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _Yvault300.contract.FilterLogs(opts, "RoleSet", accountRule, roleRule)
	if err != nil {
		return nil, err
	}
	return &Yvault300RoleSetIterator{contract: _Yvault300.contract, event: "RoleSet", logs: logs, sub: sub}, nil
}

// WatchRoleSet is a free log subscription operation binding the contract event 0x78557646b1d8efa2cd49740d66df5aca39eb610ca8ca0e1ccac08979b6b2c46e.
//
// Solidity: event RoleSet(address indexed account, uint256 indexed role)
func (_Yvault300 *Yvault300Filterer) WatchRoleSet(opts *bind.WatchOpts, sink chan<- *Yvault300RoleSet, account []common.Address, role []*big.Int) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _Yvault300.contract.WatchLogs(opts, "RoleSet", accountRule, roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault300RoleSet)
				if err := _Yvault300.contract.UnpackLog(event, "RoleSet", log); err != nil {
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

// ParseRoleSet is a log parse operation binding the contract event 0x78557646b1d8efa2cd49740d66df5aca39eb610ca8ca0e1ccac08979b6b2c46e.
//
// Solidity: event RoleSet(address indexed account, uint256 indexed role)
func (_Yvault300 *Yvault300Filterer) ParseRoleSet(log types.Log) (*Yvault300RoleSet, error) {
	event := new(Yvault300RoleSet)
	if err := _Yvault300.contract.UnpackLog(event, "RoleSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault300RoleStatusChangedIterator is returned from FilterRoleStatusChanged and is used to iterate over the raw logs and unpacked data for RoleStatusChanged events raised by the Yvault300 contract.
type Yvault300RoleStatusChangedIterator struct {
	Event *Yvault300RoleStatusChanged // Event containing the contract specifics and raw log

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
func (it *Yvault300RoleStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault300RoleStatusChanged)
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
		it.Event = new(Yvault300RoleStatusChanged)
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
func (it *Yvault300RoleStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault300RoleStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault300RoleStatusChanged represents a RoleStatusChanged event raised by the Yvault300 contract.
type Yvault300RoleStatusChanged struct {
	Role   *big.Int
	Status *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRoleStatusChanged is a free log retrieval operation binding the contract event 0xfe075e51fb76b038a5d44dd2e56b16e6c928e35c0f3cc237312ad09bbca5aee5.
//
// Solidity: event RoleStatusChanged(uint256 indexed role, uint256 indexed status)
func (_Yvault300 *Yvault300Filterer) FilterRoleStatusChanged(opts *bind.FilterOpts, role []*big.Int, status []*big.Int) (*Yvault300RoleStatusChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _Yvault300.contract.FilterLogs(opts, "RoleStatusChanged", roleRule, statusRule)
	if err != nil {
		return nil, err
	}
	return &Yvault300RoleStatusChangedIterator{contract: _Yvault300.contract, event: "RoleStatusChanged", logs: logs, sub: sub}, nil
}

// WatchRoleStatusChanged is a free log subscription operation binding the contract event 0xfe075e51fb76b038a5d44dd2e56b16e6c928e35c0f3cc237312ad09bbca5aee5.
//
// Solidity: event RoleStatusChanged(uint256 indexed role, uint256 indexed status)
func (_Yvault300 *Yvault300Filterer) WatchRoleStatusChanged(opts *bind.WatchOpts, sink chan<- *Yvault300RoleStatusChanged, role []*big.Int, status []*big.Int) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _Yvault300.contract.WatchLogs(opts, "RoleStatusChanged", roleRule, statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault300RoleStatusChanged)
				if err := _Yvault300.contract.UnpackLog(event, "RoleStatusChanged", log); err != nil {
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

// ParseRoleStatusChanged is a log parse operation binding the contract event 0xfe075e51fb76b038a5d44dd2e56b16e6c928e35c0f3cc237312ad09bbca5aee5.
//
// Solidity: event RoleStatusChanged(uint256 indexed role, uint256 indexed status)
func (_Yvault300 *Yvault300Filterer) ParseRoleStatusChanged(log types.Log) (*Yvault300RoleStatusChanged, error) {
	event := new(Yvault300RoleStatusChanged)
	if err := _Yvault300.contract.UnpackLog(event, "RoleStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault300ShutdownIterator is returned from FilterShutdown and is used to iterate over the raw logs and unpacked data for Shutdown events raised by the Yvault300 contract.
type Yvault300ShutdownIterator struct {
	Event *Yvault300Shutdown // Event containing the contract specifics and raw log

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
func (it *Yvault300ShutdownIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault300Shutdown)
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
		it.Event = new(Yvault300Shutdown)
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
func (it *Yvault300ShutdownIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault300ShutdownIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault300Shutdown represents a Shutdown event raised by the Yvault300 contract.
type Yvault300Shutdown struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterShutdown is a free log retrieval operation binding the contract event 0x4426aa1fb73e391071491fcfe21a88b5c38a0a0333a1f6e77161470439704cf8.
//
// Solidity: event Shutdown()
func (_Yvault300 *Yvault300Filterer) FilterShutdown(opts *bind.FilterOpts) (*Yvault300ShutdownIterator, error) {

	logs, sub, err := _Yvault300.contract.FilterLogs(opts, "Shutdown")
	if err != nil {
		return nil, err
	}
	return &Yvault300ShutdownIterator{contract: _Yvault300.contract, event: "Shutdown", logs: logs, sub: sub}, nil
}

// WatchShutdown is a free log subscription operation binding the contract event 0x4426aa1fb73e391071491fcfe21a88b5c38a0a0333a1f6e77161470439704cf8.
//
// Solidity: event Shutdown()
func (_Yvault300 *Yvault300Filterer) WatchShutdown(opts *bind.WatchOpts, sink chan<- *Yvault300Shutdown) (event.Subscription, error) {

	logs, sub, err := _Yvault300.contract.WatchLogs(opts, "Shutdown")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault300Shutdown)
				if err := _Yvault300.contract.UnpackLog(event, "Shutdown", log); err != nil {
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

// ParseShutdown is a log parse operation binding the contract event 0x4426aa1fb73e391071491fcfe21a88b5c38a0a0333a1f6e77161470439704cf8.
//
// Solidity: event Shutdown()
func (_Yvault300 *Yvault300Filterer) ParseShutdown(log types.Log) (*Yvault300Shutdown, error) {
	event := new(Yvault300Shutdown)
	if err := _Yvault300.contract.UnpackLog(event, "Shutdown", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault300StrategyChangedIterator is returned from FilterStrategyChanged and is used to iterate over the raw logs and unpacked data for StrategyChanged events raised by the Yvault300 contract.
type Yvault300StrategyChangedIterator struct {
	Event *Yvault300StrategyChanged // Event containing the contract specifics and raw log

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
func (it *Yvault300StrategyChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault300StrategyChanged)
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
		it.Event = new(Yvault300StrategyChanged)
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
func (it *Yvault300StrategyChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault300StrategyChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault300StrategyChanged represents a StrategyChanged event raised by the Yvault300 contract.
type Yvault300StrategyChanged struct {
	Strategy   common.Address
	ChangeType *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterStrategyChanged is a free log retrieval operation binding the contract event 0xde8ff765a5c5dad48d27bc9faa99836fb81f3b07c9dc62cfe005475d6b83a2ca.
//
// Solidity: event StrategyChanged(address indexed strategy, uint256 indexed change_type)
func (_Yvault300 *Yvault300Filterer) FilterStrategyChanged(opts *bind.FilterOpts, strategy []common.Address, change_type []*big.Int) (*Yvault300StrategyChangedIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}
	var change_typeRule []interface{}
	for _, change_typeItem := range change_type {
		change_typeRule = append(change_typeRule, change_typeItem)
	}

	logs, sub, err := _Yvault300.contract.FilterLogs(opts, "StrategyChanged", strategyRule, change_typeRule)
	if err != nil {
		return nil, err
	}
	return &Yvault300StrategyChangedIterator{contract: _Yvault300.contract, event: "StrategyChanged", logs: logs, sub: sub}, nil
}

// WatchStrategyChanged is a free log subscription operation binding the contract event 0xde8ff765a5c5dad48d27bc9faa99836fb81f3b07c9dc62cfe005475d6b83a2ca.
//
// Solidity: event StrategyChanged(address indexed strategy, uint256 indexed change_type)
func (_Yvault300 *Yvault300Filterer) WatchStrategyChanged(opts *bind.WatchOpts, sink chan<- *Yvault300StrategyChanged, strategy []common.Address, change_type []*big.Int) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}
	var change_typeRule []interface{}
	for _, change_typeItem := range change_type {
		change_typeRule = append(change_typeRule, change_typeItem)
	}

	logs, sub, err := _Yvault300.contract.WatchLogs(opts, "StrategyChanged", strategyRule, change_typeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault300StrategyChanged)
				if err := _Yvault300.contract.UnpackLog(event, "StrategyChanged", log); err != nil {
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

// ParseStrategyChanged is a log parse operation binding the contract event 0xde8ff765a5c5dad48d27bc9faa99836fb81f3b07c9dc62cfe005475d6b83a2ca.
//
// Solidity: event StrategyChanged(address indexed strategy, uint256 indexed change_type)
func (_Yvault300 *Yvault300Filterer) ParseStrategyChanged(log types.Log) (*Yvault300StrategyChanged, error) {
	event := new(Yvault300StrategyChanged)
	if err := _Yvault300.contract.UnpackLog(event, "StrategyChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault300StrategyReportedIterator is returned from FilterStrategyReported and is used to iterate over the raw logs and unpacked data for StrategyReported events raised by the Yvault300 contract.
type Yvault300StrategyReportedIterator struct {
	Event *Yvault300StrategyReported // Event containing the contract specifics and raw log

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
func (it *Yvault300StrategyReportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault300StrategyReported)
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
		it.Event = new(Yvault300StrategyReported)
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
func (it *Yvault300StrategyReportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault300StrategyReportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault300StrategyReported represents a StrategyReported event raised by the Yvault300 contract.
type Yvault300StrategyReported struct {
	Strategy     common.Address
	Gain         *big.Int
	Loss         *big.Int
	CurrentDebt  *big.Int
	ProtocolFees *big.Int
	TotalFees    *big.Int
	TotalRefunds *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStrategyReported is a free log retrieval operation binding the contract event 0x7f2ad1d3ba35276f35ef140f83e3e0f17b23064fd710113d3f7a5ab30d267811.
//
// Solidity: event StrategyReported(address indexed strategy, uint256 gain, uint256 loss, uint256 current_debt, uint256 protocol_fees, uint256 total_fees, uint256 total_refunds)
func (_Yvault300 *Yvault300Filterer) FilterStrategyReported(opts *bind.FilterOpts, strategy []common.Address) (*Yvault300StrategyReportedIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault300.contract.FilterLogs(opts, "StrategyReported", strategyRule)
	if err != nil {
		return nil, err
	}
	return &Yvault300StrategyReportedIterator{contract: _Yvault300.contract, event: "StrategyReported", logs: logs, sub: sub}, nil
}

// WatchStrategyReported is a free log subscription operation binding the contract event 0x7f2ad1d3ba35276f35ef140f83e3e0f17b23064fd710113d3f7a5ab30d267811.
//
// Solidity: event StrategyReported(address indexed strategy, uint256 gain, uint256 loss, uint256 current_debt, uint256 protocol_fees, uint256 total_fees, uint256 total_refunds)
func (_Yvault300 *Yvault300Filterer) WatchStrategyReported(opts *bind.WatchOpts, sink chan<- *Yvault300StrategyReported, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault300.contract.WatchLogs(opts, "StrategyReported", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault300StrategyReported)
				if err := _Yvault300.contract.UnpackLog(event, "StrategyReported", log); err != nil {
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

// ParseStrategyReported is a log parse operation binding the contract event 0x7f2ad1d3ba35276f35ef140f83e3e0f17b23064fd710113d3f7a5ab30d267811.
//
// Solidity: event StrategyReported(address indexed strategy, uint256 gain, uint256 loss, uint256 current_debt, uint256 protocol_fees, uint256 total_fees, uint256 total_refunds)
func (_Yvault300 *Yvault300Filterer) ParseStrategyReported(log types.Log) (*Yvault300StrategyReported, error) {
	event := new(Yvault300StrategyReported)
	if err := _Yvault300.contract.UnpackLog(event, "StrategyReported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault300TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Yvault300 contract.
type Yvault300TransferIterator struct {
	Event *Yvault300Transfer // Event containing the contract specifics and raw log

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
func (it *Yvault300TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault300Transfer)
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
		it.Event = new(Yvault300Transfer)
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
func (it *Yvault300TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault300TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault300Transfer represents a Transfer event raised by the Yvault300 contract.
type Yvault300Transfer struct {
	Sender   common.Address
	Receiver common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_Yvault300 *Yvault300Filterer) FilterTransfer(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*Yvault300TransferIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Yvault300.contract.FilterLogs(opts, "Transfer", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &Yvault300TransferIterator{contract: _Yvault300.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_Yvault300 *Yvault300Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Yvault300Transfer, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Yvault300.contract.WatchLogs(opts, "Transfer", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault300Transfer)
				if err := _Yvault300.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Yvault300 *Yvault300Filterer) ParseTransfer(log types.Log) (*Yvault300Transfer, error) {
	event := new(Yvault300Transfer)
	if err := _Yvault300.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault300UpdateAccountantIterator is returned from FilterUpdateAccountant and is used to iterate over the raw logs and unpacked data for UpdateAccountant events raised by the Yvault300 contract.
type Yvault300UpdateAccountantIterator struct {
	Event *Yvault300UpdateAccountant // Event containing the contract specifics and raw log

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
func (it *Yvault300UpdateAccountantIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault300UpdateAccountant)
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
		it.Event = new(Yvault300UpdateAccountant)
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
func (it *Yvault300UpdateAccountantIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault300UpdateAccountantIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault300UpdateAccountant represents a UpdateAccountant event raised by the Yvault300 contract.
type Yvault300UpdateAccountant struct {
	Accountant common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateAccountant is a free log retrieval operation binding the contract event 0x28709a2dab2a5d5e8688e96159011151c51644ab21839a8a45b449634d7c8b2b.
//
// Solidity: event UpdateAccountant(address indexed accountant)
func (_Yvault300 *Yvault300Filterer) FilterUpdateAccountant(opts *bind.FilterOpts, accountant []common.Address) (*Yvault300UpdateAccountantIterator, error) {

	var accountantRule []interface{}
	for _, accountantItem := range accountant {
		accountantRule = append(accountantRule, accountantItem)
	}

	logs, sub, err := _Yvault300.contract.FilterLogs(opts, "UpdateAccountant", accountantRule)
	if err != nil {
		return nil, err
	}
	return &Yvault300UpdateAccountantIterator{contract: _Yvault300.contract, event: "UpdateAccountant", logs: logs, sub: sub}, nil
}

// WatchUpdateAccountant is a free log subscription operation binding the contract event 0x28709a2dab2a5d5e8688e96159011151c51644ab21839a8a45b449634d7c8b2b.
//
// Solidity: event UpdateAccountant(address indexed accountant)
func (_Yvault300 *Yvault300Filterer) WatchUpdateAccountant(opts *bind.WatchOpts, sink chan<- *Yvault300UpdateAccountant, accountant []common.Address) (event.Subscription, error) {

	var accountantRule []interface{}
	for _, accountantItem := range accountant {
		accountantRule = append(accountantRule, accountantItem)
	}

	logs, sub, err := _Yvault300.contract.WatchLogs(opts, "UpdateAccountant", accountantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault300UpdateAccountant)
				if err := _Yvault300.contract.UnpackLog(event, "UpdateAccountant", log); err != nil {
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

// ParseUpdateAccountant is a log parse operation binding the contract event 0x28709a2dab2a5d5e8688e96159011151c51644ab21839a8a45b449634d7c8b2b.
//
// Solidity: event UpdateAccountant(address indexed accountant)
func (_Yvault300 *Yvault300Filterer) ParseUpdateAccountant(log types.Log) (*Yvault300UpdateAccountant, error) {
	event := new(Yvault300UpdateAccountant)
	if err := _Yvault300.contract.UnpackLog(event, "UpdateAccountant", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault300UpdateDefaultQueueIterator is returned from FilterUpdateDefaultQueue and is used to iterate over the raw logs and unpacked data for UpdateDefaultQueue events raised by the Yvault300 contract.
type Yvault300UpdateDefaultQueueIterator struct {
	Event *Yvault300UpdateDefaultQueue // Event containing the contract specifics and raw log

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
func (it *Yvault300UpdateDefaultQueueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault300UpdateDefaultQueue)
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
		it.Event = new(Yvault300UpdateDefaultQueue)
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
func (it *Yvault300UpdateDefaultQueueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault300UpdateDefaultQueueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault300UpdateDefaultQueue represents a UpdateDefaultQueue event raised by the Yvault300 contract.
type Yvault300UpdateDefaultQueue struct {
	NewDefaultQueue []common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateDefaultQueue is a free log retrieval operation binding the contract event 0x0bc0cb8c5ccee13e6a2fd26a699f57ad7ff6e454e6aae97ec41cd2eb9ebd63a5.
//
// Solidity: event UpdateDefaultQueue(address[] new_default_queue)
func (_Yvault300 *Yvault300Filterer) FilterUpdateDefaultQueue(opts *bind.FilterOpts) (*Yvault300UpdateDefaultQueueIterator, error) {

	logs, sub, err := _Yvault300.contract.FilterLogs(opts, "UpdateDefaultQueue")
	if err != nil {
		return nil, err
	}
	return &Yvault300UpdateDefaultQueueIterator{contract: _Yvault300.contract, event: "UpdateDefaultQueue", logs: logs, sub: sub}, nil
}

// WatchUpdateDefaultQueue is a free log subscription operation binding the contract event 0x0bc0cb8c5ccee13e6a2fd26a699f57ad7ff6e454e6aae97ec41cd2eb9ebd63a5.
//
// Solidity: event UpdateDefaultQueue(address[] new_default_queue)
func (_Yvault300 *Yvault300Filterer) WatchUpdateDefaultQueue(opts *bind.WatchOpts, sink chan<- *Yvault300UpdateDefaultQueue) (event.Subscription, error) {

	logs, sub, err := _Yvault300.contract.WatchLogs(opts, "UpdateDefaultQueue")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault300UpdateDefaultQueue)
				if err := _Yvault300.contract.UnpackLog(event, "UpdateDefaultQueue", log); err != nil {
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

// ParseUpdateDefaultQueue is a log parse operation binding the contract event 0x0bc0cb8c5ccee13e6a2fd26a699f57ad7ff6e454e6aae97ec41cd2eb9ebd63a5.
//
// Solidity: event UpdateDefaultQueue(address[] new_default_queue)
func (_Yvault300 *Yvault300Filterer) ParseUpdateDefaultQueue(log types.Log) (*Yvault300UpdateDefaultQueue, error) {
	event := new(Yvault300UpdateDefaultQueue)
	if err := _Yvault300.contract.UnpackLog(event, "UpdateDefaultQueue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault300UpdateDepositLimitIterator is returned from FilterUpdateDepositLimit and is used to iterate over the raw logs and unpacked data for UpdateDepositLimit events raised by the Yvault300 contract.
type Yvault300UpdateDepositLimitIterator struct {
	Event *Yvault300UpdateDepositLimit // Event containing the contract specifics and raw log

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
func (it *Yvault300UpdateDepositLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault300UpdateDepositLimit)
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
		it.Event = new(Yvault300UpdateDepositLimit)
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
func (it *Yvault300UpdateDepositLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault300UpdateDepositLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault300UpdateDepositLimit represents a UpdateDepositLimit event raised by the Yvault300 contract.
type Yvault300UpdateDepositLimit struct {
	DepositLimit *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateDepositLimit is a free log retrieval operation binding the contract event 0xae565aab888bca5e19e25a13db7b0c9144305bf55cb0f3f4d724f730e5acdd62.
//
// Solidity: event UpdateDepositLimit(uint256 deposit_limit)
func (_Yvault300 *Yvault300Filterer) FilterUpdateDepositLimit(opts *bind.FilterOpts) (*Yvault300UpdateDepositLimitIterator, error) {

	logs, sub, err := _Yvault300.contract.FilterLogs(opts, "UpdateDepositLimit")
	if err != nil {
		return nil, err
	}
	return &Yvault300UpdateDepositLimitIterator{contract: _Yvault300.contract, event: "UpdateDepositLimit", logs: logs, sub: sub}, nil
}

// WatchUpdateDepositLimit is a free log subscription operation binding the contract event 0xae565aab888bca5e19e25a13db7b0c9144305bf55cb0f3f4d724f730e5acdd62.
//
// Solidity: event UpdateDepositLimit(uint256 deposit_limit)
func (_Yvault300 *Yvault300Filterer) WatchUpdateDepositLimit(opts *bind.WatchOpts, sink chan<- *Yvault300UpdateDepositLimit) (event.Subscription, error) {

	logs, sub, err := _Yvault300.contract.WatchLogs(opts, "UpdateDepositLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault300UpdateDepositLimit)
				if err := _Yvault300.contract.UnpackLog(event, "UpdateDepositLimit", log); err != nil {
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
// Solidity: event UpdateDepositLimit(uint256 deposit_limit)
func (_Yvault300 *Yvault300Filterer) ParseUpdateDepositLimit(log types.Log) (*Yvault300UpdateDepositLimit, error) {
	event := new(Yvault300UpdateDepositLimit)
	if err := _Yvault300.contract.UnpackLog(event, "UpdateDepositLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault300UpdateDepositLimitModuleIterator is returned from FilterUpdateDepositLimitModule and is used to iterate over the raw logs and unpacked data for UpdateDepositLimitModule events raised by the Yvault300 contract.
type Yvault300UpdateDepositLimitModuleIterator struct {
	Event *Yvault300UpdateDepositLimitModule // Event containing the contract specifics and raw log

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
func (it *Yvault300UpdateDepositLimitModuleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault300UpdateDepositLimitModule)
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
		it.Event = new(Yvault300UpdateDepositLimitModule)
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
func (it *Yvault300UpdateDepositLimitModuleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault300UpdateDepositLimitModuleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault300UpdateDepositLimitModule represents a UpdateDepositLimitModule event raised by the Yvault300 contract.
type Yvault300UpdateDepositLimitModule struct {
	DepositLimitModule common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterUpdateDepositLimitModule is a free log retrieval operation binding the contract event 0x777d215db24fb9fee4ded85c66b422abd7162a1caa6ed3ec4c031f6cd29ada52.
//
// Solidity: event UpdateDepositLimitModule(address indexed deposit_limit_module)
func (_Yvault300 *Yvault300Filterer) FilterUpdateDepositLimitModule(opts *bind.FilterOpts, deposit_limit_module []common.Address) (*Yvault300UpdateDepositLimitModuleIterator, error) {

	var deposit_limit_moduleRule []interface{}
	for _, deposit_limit_moduleItem := range deposit_limit_module {
		deposit_limit_moduleRule = append(deposit_limit_moduleRule, deposit_limit_moduleItem)
	}

	logs, sub, err := _Yvault300.contract.FilterLogs(opts, "UpdateDepositLimitModule", deposit_limit_moduleRule)
	if err != nil {
		return nil, err
	}
	return &Yvault300UpdateDepositLimitModuleIterator{contract: _Yvault300.contract, event: "UpdateDepositLimitModule", logs: logs, sub: sub}, nil
}

// WatchUpdateDepositLimitModule is a free log subscription operation binding the contract event 0x777d215db24fb9fee4ded85c66b422abd7162a1caa6ed3ec4c031f6cd29ada52.
//
// Solidity: event UpdateDepositLimitModule(address indexed deposit_limit_module)
func (_Yvault300 *Yvault300Filterer) WatchUpdateDepositLimitModule(opts *bind.WatchOpts, sink chan<- *Yvault300UpdateDepositLimitModule, deposit_limit_module []common.Address) (event.Subscription, error) {

	var deposit_limit_moduleRule []interface{}
	for _, deposit_limit_moduleItem := range deposit_limit_module {
		deposit_limit_moduleRule = append(deposit_limit_moduleRule, deposit_limit_moduleItem)
	}

	logs, sub, err := _Yvault300.contract.WatchLogs(opts, "UpdateDepositLimitModule", deposit_limit_moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault300UpdateDepositLimitModule)
				if err := _Yvault300.contract.UnpackLog(event, "UpdateDepositLimitModule", log); err != nil {
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

// ParseUpdateDepositLimitModule is a log parse operation binding the contract event 0x777d215db24fb9fee4ded85c66b422abd7162a1caa6ed3ec4c031f6cd29ada52.
//
// Solidity: event UpdateDepositLimitModule(address indexed deposit_limit_module)
func (_Yvault300 *Yvault300Filterer) ParseUpdateDepositLimitModule(log types.Log) (*Yvault300UpdateDepositLimitModule, error) {
	event := new(Yvault300UpdateDepositLimitModule)
	if err := _Yvault300.contract.UnpackLog(event, "UpdateDepositLimitModule", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault300UpdateMinimumTotalIdleIterator is returned from FilterUpdateMinimumTotalIdle and is used to iterate over the raw logs and unpacked data for UpdateMinimumTotalIdle events raised by the Yvault300 contract.
type Yvault300UpdateMinimumTotalIdleIterator struct {
	Event *Yvault300UpdateMinimumTotalIdle // Event containing the contract specifics and raw log

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
func (it *Yvault300UpdateMinimumTotalIdleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault300UpdateMinimumTotalIdle)
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
		it.Event = new(Yvault300UpdateMinimumTotalIdle)
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
func (it *Yvault300UpdateMinimumTotalIdleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault300UpdateMinimumTotalIdleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault300UpdateMinimumTotalIdle represents a UpdateMinimumTotalIdle event raised by the Yvault300 contract.
type Yvault300UpdateMinimumTotalIdle struct {
	MinimumTotalIdle *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUpdateMinimumTotalIdle is a free log retrieval operation binding the contract event 0x01a4494beed88920b88742cc58f2744e198f55ff192635a1fbabc6be8ffade81.
//
// Solidity: event UpdateMinimumTotalIdle(uint256 minimum_total_idle)
func (_Yvault300 *Yvault300Filterer) FilterUpdateMinimumTotalIdle(opts *bind.FilterOpts) (*Yvault300UpdateMinimumTotalIdleIterator, error) {

	logs, sub, err := _Yvault300.contract.FilterLogs(opts, "UpdateMinimumTotalIdle")
	if err != nil {
		return nil, err
	}
	return &Yvault300UpdateMinimumTotalIdleIterator{contract: _Yvault300.contract, event: "UpdateMinimumTotalIdle", logs: logs, sub: sub}, nil
}

// WatchUpdateMinimumTotalIdle is a free log subscription operation binding the contract event 0x01a4494beed88920b88742cc58f2744e198f55ff192635a1fbabc6be8ffade81.
//
// Solidity: event UpdateMinimumTotalIdle(uint256 minimum_total_idle)
func (_Yvault300 *Yvault300Filterer) WatchUpdateMinimumTotalIdle(opts *bind.WatchOpts, sink chan<- *Yvault300UpdateMinimumTotalIdle) (event.Subscription, error) {

	logs, sub, err := _Yvault300.contract.WatchLogs(opts, "UpdateMinimumTotalIdle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault300UpdateMinimumTotalIdle)
				if err := _Yvault300.contract.UnpackLog(event, "UpdateMinimumTotalIdle", log); err != nil {
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

// ParseUpdateMinimumTotalIdle is a log parse operation binding the contract event 0x01a4494beed88920b88742cc58f2744e198f55ff192635a1fbabc6be8ffade81.
//
// Solidity: event UpdateMinimumTotalIdle(uint256 minimum_total_idle)
func (_Yvault300 *Yvault300Filterer) ParseUpdateMinimumTotalIdle(log types.Log) (*Yvault300UpdateMinimumTotalIdle, error) {
	event := new(Yvault300UpdateMinimumTotalIdle)
	if err := _Yvault300.contract.UnpackLog(event, "UpdateMinimumTotalIdle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault300UpdateProfitMaxUnlockTimeIterator is returned from FilterUpdateProfitMaxUnlockTime and is used to iterate over the raw logs and unpacked data for UpdateProfitMaxUnlockTime events raised by the Yvault300 contract.
type Yvault300UpdateProfitMaxUnlockTimeIterator struct {
	Event *Yvault300UpdateProfitMaxUnlockTime // Event containing the contract specifics and raw log

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
func (it *Yvault300UpdateProfitMaxUnlockTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault300UpdateProfitMaxUnlockTime)
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
		it.Event = new(Yvault300UpdateProfitMaxUnlockTime)
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
func (it *Yvault300UpdateProfitMaxUnlockTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault300UpdateProfitMaxUnlockTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault300UpdateProfitMaxUnlockTime represents a UpdateProfitMaxUnlockTime event raised by the Yvault300 contract.
type Yvault300UpdateProfitMaxUnlockTime struct {
	ProfitMaxUnlockTime *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterUpdateProfitMaxUnlockTime is a free log retrieval operation binding the contract event 0xf361aed463da6fa20358e45c6209f1d3e16d4eca706e6eab0b0aeb338729c77a.
//
// Solidity: event UpdateProfitMaxUnlockTime(uint256 profit_max_unlock_time)
func (_Yvault300 *Yvault300Filterer) FilterUpdateProfitMaxUnlockTime(opts *bind.FilterOpts) (*Yvault300UpdateProfitMaxUnlockTimeIterator, error) {

	logs, sub, err := _Yvault300.contract.FilterLogs(opts, "UpdateProfitMaxUnlockTime")
	if err != nil {
		return nil, err
	}
	return &Yvault300UpdateProfitMaxUnlockTimeIterator{contract: _Yvault300.contract, event: "UpdateProfitMaxUnlockTime", logs: logs, sub: sub}, nil
}

// WatchUpdateProfitMaxUnlockTime is a free log subscription operation binding the contract event 0xf361aed463da6fa20358e45c6209f1d3e16d4eca706e6eab0b0aeb338729c77a.
//
// Solidity: event UpdateProfitMaxUnlockTime(uint256 profit_max_unlock_time)
func (_Yvault300 *Yvault300Filterer) WatchUpdateProfitMaxUnlockTime(opts *bind.WatchOpts, sink chan<- *Yvault300UpdateProfitMaxUnlockTime) (event.Subscription, error) {

	logs, sub, err := _Yvault300.contract.WatchLogs(opts, "UpdateProfitMaxUnlockTime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault300UpdateProfitMaxUnlockTime)
				if err := _Yvault300.contract.UnpackLog(event, "UpdateProfitMaxUnlockTime", log); err != nil {
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

// ParseUpdateProfitMaxUnlockTime is a log parse operation binding the contract event 0xf361aed463da6fa20358e45c6209f1d3e16d4eca706e6eab0b0aeb338729c77a.
//
// Solidity: event UpdateProfitMaxUnlockTime(uint256 profit_max_unlock_time)
func (_Yvault300 *Yvault300Filterer) ParseUpdateProfitMaxUnlockTime(log types.Log) (*Yvault300UpdateProfitMaxUnlockTime, error) {
	event := new(Yvault300UpdateProfitMaxUnlockTime)
	if err := _Yvault300.contract.UnpackLog(event, "UpdateProfitMaxUnlockTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault300UpdateRoleManagerIterator is returned from FilterUpdateRoleManager and is used to iterate over the raw logs and unpacked data for UpdateRoleManager events raised by the Yvault300 contract.
type Yvault300UpdateRoleManagerIterator struct {
	Event *Yvault300UpdateRoleManager // Event containing the contract specifics and raw log

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
func (it *Yvault300UpdateRoleManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault300UpdateRoleManager)
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
		it.Event = new(Yvault300UpdateRoleManager)
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
func (it *Yvault300UpdateRoleManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault300UpdateRoleManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault300UpdateRoleManager represents a UpdateRoleManager event raised by the Yvault300 contract.
type Yvault300UpdateRoleManager struct {
	RoleManager common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdateRoleManager is a free log retrieval operation binding the contract event 0xce93baa0b608a7d420822b6b90cfcccb70574363ba4fd26ef5ac17dd465016c4.
//
// Solidity: event UpdateRoleManager(address indexed role_manager)
func (_Yvault300 *Yvault300Filterer) FilterUpdateRoleManager(opts *bind.FilterOpts, role_manager []common.Address) (*Yvault300UpdateRoleManagerIterator, error) {

	var role_managerRule []interface{}
	for _, role_managerItem := range role_manager {
		role_managerRule = append(role_managerRule, role_managerItem)
	}

	logs, sub, err := _Yvault300.contract.FilterLogs(opts, "UpdateRoleManager", role_managerRule)
	if err != nil {
		return nil, err
	}
	return &Yvault300UpdateRoleManagerIterator{contract: _Yvault300.contract, event: "UpdateRoleManager", logs: logs, sub: sub}, nil
}

// WatchUpdateRoleManager is a free log subscription operation binding the contract event 0xce93baa0b608a7d420822b6b90cfcccb70574363ba4fd26ef5ac17dd465016c4.
//
// Solidity: event UpdateRoleManager(address indexed role_manager)
func (_Yvault300 *Yvault300Filterer) WatchUpdateRoleManager(opts *bind.WatchOpts, sink chan<- *Yvault300UpdateRoleManager, role_manager []common.Address) (event.Subscription, error) {

	var role_managerRule []interface{}
	for _, role_managerItem := range role_manager {
		role_managerRule = append(role_managerRule, role_managerItem)
	}

	logs, sub, err := _Yvault300.contract.WatchLogs(opts, "UpdateRoleManager", role_managerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault300UpdateRoleManager)
				if err := _Yvault300.contract.UnpackLog(event, "UpdateRoleManager", log); err != nil {
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

// ParseUpdateRoleManager is a log parse operation binding the contract event 0xce93baa0b608a7d420822b6b90cfcccb70574363ba4fd26ef5ac17dd465016c4.
//
// Solidity: event UpdateRoleManager(address indexed role_manager)
func (_Yvault300 *Yvault300Filterer) ParseUpdateRoleManager(log types.Log) (*Yvault300UpdateRoleManager, error) {
	event := new(Yvault300UpdateRoleManager)
	if err := _Yvault300.contract.UnpackLog(event, "UpdateRoleManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault300UpdateUseDefaultQueueIterator is returned from FilterUpdateUseDefaultQueue and is used to iterate over the raw logs and unpacked data for UpdateUseDefaultQueue events raised by the Yvault300 contract.
type Yvault300UpdateUseDefaultQueueIterator struct {
	Event *Yvault300UpdateUseDefaultQueue // Event containing the contract specifics and raw log

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
func (it *Yvault300UpdateUseDefaultQueueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault300UpdateUseDefaultQueue)
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
		it.Event = new(Yvault300UpdateUseDefaultQueue)
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
func (it *Yvault300UpdateUseDefaultQueueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault300UpdateUseDefaultQueueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault300UpdateUseDefaultQueue represents a UpdateUseDefaultQueue event raised by the Yvault300 contract.
type Yvault300UpdateUseDefaultQueue struct {
	UseDefaultQueue bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateUseDefaultQueue is a free log retrieval operation binding the contract event 0x1f88e73ebc721f227812938fe07a069ec1f7136aafacb397ed460bd15dee13f1.
//
// Solidity: event UpdateUseDefaultQueue(bool use_default_queue)
func (_Yvault300 *Yvault300Filterer) FilterUpdateUseDefaultQueue(opts *bind.FilterOpts) (*Yvault300UpdateUseDefaultQueueIterator, error) {

	logs, sub, err := _Yvault300.contract.FilterLogs(opts, "UpdateUseDefaultQueue")
	if err != nil {
		return nil, err
	}
	return &Yvault300UpdateUseDefaultQueueIterator{contract: _Yvault300.contract, event: "UpdateUseDefaultQueue", logs: logs, sub: sub}, nil
}

// WatchUpdateUseDefaultQueue is a free log subscription operation binding the contract event 0x1f88e73ebc721f227812938fe07a069ec1f7136aafacb397ed460bd15dee13f1.
//
// Solidity: event UpdateUseDefaultQueue(bool use_default_queue)
func (_Yvault300 *Yvault300Filterer) WatchUpdateUseDefaultQueue(opts *bind.WatchOpts, sink chan<- *Yvault300UpdateUseDefaultQueue) (event.Subscription, error) {

	logs, sub, err := _Yvault300.contract.WatchLogs(opts, "UpdateUseDefaultQueue")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault300UpdateUseDefaultQueue)
				if err := _Yvault300.contract.UnpackLog(event, "UpdateUseDefaultQueue", log); err != nil {
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

// ParseUpdateUseDefaultQueue is a log parse operation binding the contract event 0x1f88e73ebc721f227812938fe07a069ec1f7136aafacb397ed460bd15dee13f1.
//
// Solidity: event UpdateUseDefaultQueue(bool use_default_queue)
func (_Yvault300 *Yvault300Filterer) ParseUpdateUseDefaultQueue(log types.Log) (*Yvault300UpdateUseDefaultQueue, error) {
	event := new(Yvault300UpdateUseDefaultQueue)
	if err := _Yvault300.contract.UnpackLog(event, "UpdateUseDefaultQueue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault300UpdateWithdrawLimitModuleIterator is returned from FilterUpdateWithdrawLimitModule and is used to iterate over the raw logs and unpacked data for UpdateWithdrawLimitModule events raised by the Yvault300 contract.
type Yvault300UpdateWithdrawLimitModuleIterator struct {
	Event *Yvault300UpdateWithdrawLimitModule // Event containing the contract specifics and raw log

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
func (it *Yvault300UpdateWithdrawLimitModuleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault300UpdateWithdrawLimitModule)
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
		it.Event = new(Yvault300UpdateWithdrawLimitModule)
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
func (it *Yvault300UpdateWithdrawLimitModuleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault300UpdateWithdrawLimitModuleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault300UpdateWithdrawLimitModule represents a UpdateWithdrawLimitModule event raised by the Yvault300 contract.
type Yvault300UpdateWithdrawLimitModule struct {
	WithdrawLimitModule common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterUpdateWithdrawLimitModule is a free log retrieval operation binding the contract event 0xce6e3f8beda82a13c441d76efd4a6335760f219f38c60502e6680060874e109d.
//
// Solidity: event UpdateWithdrawLimitModule(address indexed withdraw_limit_module)
func (_Yvault300 *Yvault300Filterer) FilterUpdateWithdrawLimitModule(opts *bind.FilterOpts, withdraw_limit_module []common.Address) (*Yvault300UpdateWithdrawLimitModuleIterator, error) {

	var withdraw_limit_moduleRule []interface{}
	for _, withdraw_limit_moduleItem := range withdraw_limit_module {
		withdraw_limit_moduleRule = append(withdraw_limit_moduleRule, withdraw_limit_moduleItem)
	}

	logs, sub, err := _Yvault300.contract.FilterLogs(opts, "UpdateWithdrawLimitModule", withdraw_limit_moduleRule)
	if err != nil {
		return nil, err
	}
	return &Yvault300UpdateWithdrawLimitModuleIterator{contract: _Yvault300.contract, event: "UpdateWithdrawLimitModule", logs: logs, sub: sub}, nil
}

// WatchUpdateWithdrawLimitModule is a free log subscription operation binding the contract event 0xce6e3f8beda82a13c441d76efd4a6335760f219f38c60502e6680060874e109d.
//
// Solidity: event UpdateWithdrawLimitModule(address indexed withdraw_limit_module)
func (_Yvault300 *Yvault300Filterer) WatchUpdateWithdrawLimitModule(opts *bind.WatchOpts, sink chan<- *Yvault300UpdateWithdrawLimitModule, withdraw_limit_module []common.Address) (event.Subscription, error) {

	var withdraw_limit_moduleRule []interface{}
	for _, withdraw_limit_moduleItem := range withdraw_limit_module {
		withdraw_limit_moduleRule = append(withdraw_limit_moduleRule, withdraw_limit_moduleItem)
	}

	logs, sub, err := _Yvault300.contract.WatchLogs(opts, "UpdateWithdrawLimitModule", withdraw_limit_moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault300UpdateWithdrawLimitModule)
				if err := _Yvault300.contract.UnpackLog(event, "UpdateWithdrawLimitModule", log); err != nil {
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

// ParseUpdateWithdrawLimitModule is a log parse operation binding the contract event 0xce6e3f8beda82a13c441d76efd4a6335760f219f38c60502e6680060874e109d.
//
// Solidity: event UpdateWithdrawLimitModule(address indexed withdraw_limit_module)
func (_Yvault300 *Yvault300Filterer) ParseUpdateWithdrawLimitModule(log types.Log) (*Yvault300UpdateWithdrawLimitModule, error) {
	event := new(Yvault300UpdateWithdrawLimitModule)
	if err := _Yvault300.contract.UnpackLog(event, "UpdateWithdrawLimitModule", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault300UpdatedMaxDebtForStrategyIterator is returned from FilterUpdatedMaxDebtForStrategy and is used to iterate over the raw logs and unpacked data for UpdatedMaxDebtForStrategy events raised by the Yvault300 contract.
type Yvault300UpdatedMaxDebtForStrategyIterator struct {
	Event *Yvault300UpdatedMaxDebtForStrategy // Event containing the contract specifics and raw log

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
func (it *Yvault300UpdatedMaxDebtForStrategyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault300UpdatedMaxDebtForStrategy)
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
		it.Event = new(Yvault300UpdatedMaxDebtForStrategy)
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
func (it *Yvault300UpdatedMaxDebtForStrategyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault300UpdatedMaxDebtForStrategyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault300UpdatedMaxDebtForStrategy represents a UpdatedMaxDebtForStrategy event raised by the Yvault300 contract.
type Yvault300UpdatedMaxDebtForStrategy struct {
	Sender   common.Address
	Strategy common.Address
	NewDebt  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdatedMaxDebtForStrategy is a free log retrieval operation binding the contract event 0xb3eef2123fec1523a6bbc90aceb203000154c1a4974335fe06b544c7534d4b89.
//
// Solidity: event UpdatedMaxDebtForStrategy(address indexed sender, address indexed strategy, uint256 new_debt)
func (_Yvault300 *Yvault300Filterer) FilterUpdatedMaxDebtForStrategy(opts *bind.FilterOpts, sender []common.Address, strategy []common.Address) (*Yvault300UpdatedMaxDebtForStrategyIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault300.contract.FilterLogs(opts, "UpdatedMaxDebtForStrategy", senderRule, strategyRule)
	if err != nil {
		return nil, err
	}
	return &Yvault300UpdatedMaxDebtForStrategyIterator{contract: _Yvault300.contract, event: "UpdatedMaxDebtForStrategy", logs: logs, sub: sub}, nil
}

// WatchUpdatedMaxDebtForStrategy is a free log subscription operation binding the contract event 0xb3eef2123fec1523a6bbc90aceb203000154c1a4974335fe06b544c7534d4b89.
//
// Solidity: event UpdatedMaxDebtForStrategy(address indexed sender, address indexed strategy, uint256 new_debt)
func (_Yvault300 *Yvault300Filterer) WatchUpdatedMaxDebtForStrategy(opts *bind.WatchOpts, sink chan<- *Yvault300UpdatedMaxDebtForStrategy, sender []common.Address, strategy []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault300.contract.WatchLogs(opts, "UpdatedMaxDebtForStrategy", senderRule, strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault300UpdatedMaxDebtForStrategy)
				if err := _Yvault300.contract.UnpackLog(event, "UpdatedMaxDebtForStrategy", log); err != nil {
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

// ParseUpdatedMaxDebtForStrategy is a log parse operation binding the contract event 0xb3eef2123fec1523a6bbc90aceb203000154c1a4974335fe06b544c7534d4b89.
//
// Solidity: event UpdatedMaxDebtForStrategy(address indexed sender, address indexed strategy, uint256 new_debt)
func (_Yvault300 *Yvault300Filterer) ParseUpdatedMaxDebtForStrategy(log types.Log) (*Yvault300UpdatedMaxDebtForStrategy, error) {
	event := new(Yvault300UpdatedMaxDebtForStrategy)
	if err := _Yvault300.contract.UnpackLog(event, "UpdatedMaxDebtForStrategy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault300WithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Yvault300 contract.
type Yvault300WithdrawIterator struct {
	Event *Yvault300Withdraw // Event containing the contract specifics and raw log

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
func (it *Yvault300WithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault300Withdraw)
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
		it.Event = new(Yvault300Withdraw)
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
func (it *Yvault300WithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault300WithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault300Withdraw represents a Withdraw event raised by the Yvault300 contract.
type Yvault300Withdraw struct {
	Sender   common.Address
	Receiver common.Address
	Owner    common.Address
	Assets   *big.Int
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_Yvault300 *Yvault300Filterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address, owner []common.Address) (*Yvault300WithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Yvault300.contract.FilterLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &Yvault300WithdrawIterator{contract: _Yvault300.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_Yvault300 *Yvault300Filterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *Yvault300Withdraw, sender []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Yvault300.contract.WatchLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault300Withdraw)
				if err := _Yvault300.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_Yvault300 *Yvault300Filterer) ParseWithdraw(log types.Log) (*Yvault300Withdraw, error) {
	event := new(Yvault300Withdraw)
	if err := _Yvault300.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
