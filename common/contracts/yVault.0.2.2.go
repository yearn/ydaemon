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

// Yvault022MetaData contains all meta data concerning the Yvault022 contract.
var Yvault022MetaData = &bind.MetaData{
	ABI: "[{\"name\":\"Transfer\",\"inputs\":[{\"type\":\"address\",\"name\":\"sender\",\"indexed\":true},{\"type\":\"address\",\"name\":\"receiver\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Approval\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"indexed\":true},{\"type\":\"address\",\"name\":\"spender\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StrategyAdded\",\"inputs\":[{\"type\":\"address\",\"name\":\"strategy\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"debtLimit\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"rateLimit\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"performanceFee\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StrategyReported\",\"inputs\":[{\"type\":\"address\",\"name\":\"strategy\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"gain\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"loss\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"totalGain\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"totalLoss\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"totalDebt\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"debtAdded\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"debtLimit\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_token\"},{\"type\":\"address\",\"name\":\"_governance\"},{\"type\":\"address\",\"name\":\"_rewards\"},{\"type\":\"string\",\"name\":\"_nameOverride\"},{\"type\":\"string\",\"name\":\"_symbolOverride\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"name\":\"apiVersion\",\"outputs\":[{\"type\":\"string\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\",\"gas\":4489},{\"name\":\"setName\",\"outputs\":[],\"inputs\":[{\"type\":\"string\",\"name\":\"_name\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":106987},{\"name\":\"setSymbol\",\"outputs\":[],\"inputs\":[{\"type\":\"string\",\"name\":\"_symbol\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":71837},{\"name\":\"setGovernance\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_governance\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":36308},{\"name\":\"acceptGovernance\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":36234},{\"name\":\"setGuestList\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_guestList\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":36368},{\"name\":\"setRewards\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_rewards\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":36398},{\"name\":\"setDepositLimit\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_limit\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":36328},{\"name\":\"setPerformanceFee\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_fee\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":36358},{\"name\":\"setManagementFee\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_fee\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":36388},{\"name\":\"setGuardian\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_guardian\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37745},{\"name\":\"setEmergencyShutdown\",\"outputs\":[],\"inputs\":[{\"type\":\"bool\",\"name\":\"_active\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37775},{\"name\":\"setWithdrawalQueue\",\"outputs\":[],\"inputs\":[{\"type\":\"address[20]\",\"name\":\"_queue\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":750044},{\"name\":\"transfer\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_to\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":76619},{\"name\":\"transferFrom\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_from\"},{\"type\":\"address\",\"name\":\"_to\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":116382},{\"name\":\"approve\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_spender\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":38184},{\"name\":\"increaseAllowance\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_spender\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":40225},{\"name\":\"decreaseAllowance\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_spender\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":40249},{\"name\":\"permit\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"owner\"},{\"type\":\"address\",\"name\":\"spender\"},{\"type\":\"uint256\",\"name\":\"amount\"},{\"type\":\"uint256\",\"name\":\"expiry\"},{\"type\":\"bytes\",\"name\":\"signature\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":81177},{\"name\":\"totalAssets\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":4003},{\"name\":\"balanceSheetOfStrategy\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_strategy\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2508},{\"name\":\"totalBalanceSheet\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address[40]\",\"name\":\"_strategies\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":77066},{\"name\":\"deposit\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"deposit\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_amount\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"deposit\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_amount\"},{\"type\":\"address\",\"name\":\"_recipient\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"maxAvailableShares\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":359791},{\"name\":\"withdraw\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"withdraw\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_shares\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"withdraw\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_shares\"},{\"type\":\"address\",\"name\":\"_recipient\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"pricePerShare\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":12352},{\"name\":\"addStrategy\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_strategy\"},{\"type\":\"uint256\",\"name\":\"_debtLimit\"},{\"type\":\"uint256\",\"name\":\"_rateLimit\"},{\"type\":\"uint256\",\"name\":\"_performanceFee\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":1445752},{\"name\":\"updateStrategyDebtLimit\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_strategy\"},{\"type\":\"uint256\",\"name\":\"_debtLimit\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":111496},{\"name\":\"updateStrategyRateLimit\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_strategy\"},{\"type\":\"uint256\",\"name\":\"_rateLimit\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":38548},{\"name\":\"updateStrategyPerformanceFee\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_strategy\"},{\"type\":\"uint256\",\"name\":\"_performanceFee\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":38572},{\"name\":\"migrateStrategy\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_oldVersion\"},{\"type\":\"address\",\"name\":\"_newVersion\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":1178418},{\"name\":\"revokeStrategy\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"revokeStrategy\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_strategy\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"addStrategyToQueue\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_strategy\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":1194595},{\"name\":\"removeStrategyFromQueue\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_strategy\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":23068248},{\"name\":\"debtOutstanding\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"name\":\"debtOutstanding\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_strategy\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"name\":\"creditAvailable\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"name\":\"creditAvailable\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_strategy\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"name\":\"availableDepositLimit\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":9688},{\"name\":\"expectedReturn\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"name\":\"expectedReturn\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_strategy\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"name\":\"report\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_gain\"},{\"type\":\"uint256\",\"name\":\"_loss\"},{\"type\":\"uint256\",\"name\":\"_debtPayment\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":919553},{\"name\":\"sweep\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_token\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"sweep\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_token\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"name\",\"outputs\":[{\"type\":\"string\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":9053},{\"name\":\"symbol\",\"outputs\":[{\"type\":\"string\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":8106},{\"name\":\"decimals\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2711},{\"name\":\"balanceOf\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2956},{\"name\":\"allowance\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"},{\"type\":\"address\",\"name\":\"arg1\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3201},{\"name\":\"totalSupply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2801},{\"name\":\"token\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2831},{\"name\":\"governance\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2861},{\"name\":\"guardian\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2891},{\"name\":\"guestList\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2921},{\"name\":\"strategies\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"performanceFee\"},{\"type\":\"uint256\",\"name\":\"activation\"},{\"type\":\"uint256\",\"name\":\"debtLimit\"},{\"type\":\"uint256\",\"name\":\"rateLimit\"},{\"type\":\"uint256\",\"name\":\"lastReport\"},{\"type\":\"uint256\",\"name\":\"totalDebt\"},{\"type\":\"uint256\",\"name\":\"totalGain\"},{\"type\":\"uint256\",\"name\":\"totalLoss\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":10292},{\"name\":\"withdrawalQueue\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3090},{\"name\":\"emergencyShutdown\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3011},{\"name\":\"depositLimit\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3041},{\"name\":\"debtLimit\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3071},{\"name\":\"totalDebt\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3101},{\"name\":\"lastReport\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3131},{\"name\":\"activation\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3161},{\"name\":\"rewards\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3191},{\"name\":\"managementFee\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3221},{\"name\":\"performanceFee\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3251},{\"name\":\"nonces\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3496},{\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3311}]",
}

// Yvault022ABI is the input ABI used to generate the binding from.
// Deprecated: Use Yvault022MetaData.ABI instead.
var Yvault022ABI = Yvault022MetaData.ABI

// Yvault022 is an auto generated Go binding around an Ethereum contract.
type Yvault022 struct {
	Yvault022Caller     // Read-only binding to the contract
	Yvault022Transactor // Write-only binding to the contract
	Yvault022Filterer   // Log filterer for contract events
}

// Yvault022Caller is an auto generated read-only Go binding around an Ethereum contract.
type Yvault022Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Yvault022Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Yvault022Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Yvault022Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Yvault022Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Yvault022Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Yvault022Session struct {
	Contract     *Yvault022        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Yvault022CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Yvault022CallerSession struct {
	Contract *Yvault022Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// Yvault022TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Yvault022TransactorSession struct {
	Contract     *Yvault022Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// Yvault022Raw is an auto generated low-level Go binding around an Ethereum contract.
type Yvault022Raw struct {
	Contract *Yvault022 // Generic contract binding to access the raw methods on
}

// Yvault022CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Yvault022CallerRaw struct {
	Contract *Yvault022Caller // Generic read-only contract binding to access the raw methods on
}

// Yvault022TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Yvault022TransactorRaw struct {
	Contract *Yvault022Transactor // Generic write-only contract binding to access the raw methods on
}

// NewYvault022 creates a new instance of Yvault022, bound to a specific deployed contract.
func NewYvault022(address common.Address, backend bind.ContractBackend) (*Yvault022, error) {
	contract, err := bindYvault022(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Yvault022{Yvault022Caller: Yvault022Caller{contract: contract}, Yvault022Transactor: Yvault022Transactor{contract: contract}, Yvault022Filterer: Yvault022Filterer{contract: contract}}, nil
}

// NewYvault022Caller creates a new read-only instance of Yvault022, bound to a specific deployed contract.
func NewYvault022Caller(address common.Address, caller bind.ContractCaller) (*Yvault022Caller, error) {
	contract, err := bindYvault022(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Yvault022Caller{contract: contract}, nil
}

// NewYvault022Transactor creates a new write-only instance of Yvault022, bound to a specific deployed contract.
func NewYvault022Transactor(address common.Address, transactor bind.ContractTransactor) (*Yvault022Transactor, error) {
	contract, err := bindYvault022(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Yvault022Transactor{contract: contract}, nil
}

// NewYvault022Filterer creates a new log filterer instance of Yvault022, bound to a specific deployed contract.
func NewYvault022Filterer(address common.Address, filterer bind.ContractFilterer) (*Yvault022Filterer, error) {
	contract, err := bindYvault022(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Yvault022Filterer{contract: contract}, nil
}

// bindYvault022 binds a generic wrapper to an already deployed contract.
func bindYvault022(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Yvault022ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Yvault022 *Yvault022Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Yvault022.Contract.Yvault022Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Yvault022 *Yvault022Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yvault022.Contract.Yvault022Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Yvault022 *Yvault022Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Yvault022.Contract.Yvault022Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Yvault022 *Yvault022CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Yvault022.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Yvault022 *Yvault022TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yvault022.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Yvault022 *Yvault022TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Yvault022.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Yvault022 *Yvault022Caller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Yvault022.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Yvault022 *Yvault022Session) DOMAINSEPARATOR() ([32]byte, error) {
	return _Yvault022.Contract.DOMAINSEPARATOR(&_Yvault022.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Yvault022 *Yvault022CallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Yvault022.Contract.DOMAINSEPARATOR(&_Yvault022.CallOpts)
}

// Activation is a free data retrieval call binding the contract method 0x3629c8de.
//
// Solidity: function activation() view returns(uint256)
func (_Yvault022 *Yvault022Caller) Activation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault022.contract.Call(opts, &out, "activation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Activation is a free data retrieval call binding the contract method 0x3629c8de.
//
// Solidity: function activation() view returns(uint256)
func (_Yvault022 *Yvault022Session) Activation() (*big.Int, error) {
	return _Yvault022.Contract.Activation(&_Yvault022.CallOpts)
}

// Activation is a free data retrieval call binding the contract method 0x3629c8de.
//
// Solidity: function activation() view returns(uint256)
func (_Yvault022 *Yvault022CallerSession) Activation() (*big.Int, error) {
	return _Yvault022.Contract.Activation(&_Yvault022.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_Yvault022 *Yvault022Caller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yvault022.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_Yvault022 *Yvault022Session) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Yvault022.Contract.Allowance(&_Yvault022.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_Yvault022 *Yvault022CallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Yvault022.Contract.Allowance(&_Yvault022.CallOpts, arg0, arg1)
}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() pure returns(string)
func (_Yvault022 *Yvault022Caller) ApiVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Yvault022.contract.Call(opts, &out, "apiVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() pure returns(string)
func (_Yvault022 *Yvault022Session) ApiVersion() (string, error) {
	return _Yvault022.Contract.ApiVersion(&_Yvault022.CallOpts)
}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() pure returns(string)
func (_Yvault022 *Yvault022CallerSession) ApiVersion() (string, error) {
	return _Yvault022.Contract.ApiVersion(&_Yvault022.CallOpts)
}

// AvailableDepositLimit is a free data retrieval call binding the contract method 0x153c27c4.
//
// Solidity: function availableDepositLimit() view returns(uint256)
func (_Yvault022 *Yvault022Caller) AvailableDepositLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault022.contract.Call(opts, &out, "availableDepositLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailableDepositLimit is a free data retrieval call binding the contract method 0x153c27c4.
//
// Solidity: function availableDepositLimit() view returns(uint256)
func (_Yvault022 *Yvault022Session) AvailableDepositLimit() (*big.Int, error) {
	return _Yvault022.Contract.AvailableDepositLimit(&_Yvault022.CallOpts)
}

// AvailableDepositLimit is a free data retrieval call binding the contract method 0x153c27c4.
//
// Solidity: function availableDepositLimit() view returns(uint256)
func (_Yvault022 *Yvault022CallerSession) AvailableDepositLimit() (*big.Int, error) {
	return _Yvault022.Contract.AvailableDepositLimit(&_Yvault022.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_Yvault022 *Yvault022Caller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yvault022.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_Yvault022 *Yvault022Session) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Yvault022.Contract.BalanceOf(&_Yvault022.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_Yvault022 *Yvault022CallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Yvault022.Contract.BalanceOf(&_Yvault022.CallOpts, arg0)
}

// BalanceSheetOfStrategy is a free data retrieval call binding the contract method 0x5ac22080.
//
// Solidity: function balanceSheetOfStrategy(address _strategy) view returns(uint256)
func (_Yvault022 *Yvault022Caller) BalanceSheetOfStrategy(opts *bind.CallOpts, _strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yvault022.contract.Call(opts, &out, "balanceSheetOfStrategy", _strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceSheetOfStrategy is a free data retrieval call binding the contract method 0x5ac22080.
//
// Solidity: function balanceSheetOfStrategy(address _strategy) view returns(uint256)
func (_Yvault022 *Yvault022Session) BalanceSheetOfStrategy(_strategy common.Address) (*big.Int, error) {
	return _Yvault022.Contract.BalanceSheetOfStrategy(&_Yvault022.CallOpts, _strategy)
}

// BalanceSheetOfStrategy is a free data retrieval call binding the contract method 0x5ac22080.
//
// Solidity: function balanceSheetOfStrategy(address _strategy) view returns(uint256)
func (_Yvault022 *Yvault022CallerSession) BalanceSheetOfStrategy(_strategy common.Address) (*big.Int, error) {
	return _Yvault022.Contract.BalanceSheetOfStrategy(&_Yvault022.CallOpts, _strategy)
}

// CreditAvailable is a free data retrieval call binding the contract method 0x112c1f9b.
//
// Solidity: function creditAvailable() view returns(uint256)
func (_Yvault022 *Yvault022Caller) CreditAvailable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault022.contract.Call(opts, &out, "creditAvailable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreditAvailable is a free data retrieval call binding the contract method 0x112c1f9b.
//
// Solidity: function creditAvailable() view returns(uint256)
func (_Yvault022 *Yvault022Session) CreditAvailable() (*big.Int, error) {
	return _Yvault022.Contract.CreditAvailable(&_Yvault022.CallOpts)
}

// CreditAvailable is a free data retrieval call binding the contract method 0x112c1f9b.
//
// Solidity: function creditAvailable() view returns(uint256)
func (_Yvault022 *Yvault022CallerSession) CreditAvailable() (*big.Int, error) {
	return _Yvault022.Contract.CreditAvailable(&_Yvault022.CallOpts)
}

// CreditAvailable0 is a free data retrieval call binding the contract method 0xd7648013.
//
// Solidity: function creditAvailable(address _strategy) view returns(uint256)
func (_Yvault022 *Yvault022Caller) CreditAvailable0(opts *bind.CallOpts, _strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yvault022.contract.Call(opts, &out, "creditAvailable0", _strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreditAvailable0 is a free data retrieval call binding the contract method 0xd7648013.
//
// Solidity: function creditAvailable(address _strategy) view returns(uint256)
func (_Yvault022 *Yvault022Session) CreditAvailable0(_strategy common.Address) (*big.Int, error) {
	return _Yvault022.Contract.CreditAvailable0(&_Yvault022.CallOpts, _strategy)
}

// CreditAvailable0 is a free data retrieval call binding the contract method 0xd7648013.
//
// Solidity: function creditAvailable(address _strategy) view returns(uint256)
func (_Yvault022 *Yvault022CallerSession) CreditAvailable0(_strategy common.Address) (*big.Int, error) {
	return _Yvault022.Contract.CreditAvailable0(&_Yvault022.CallOpts, _strategy)
}

// DebtLimit is a free data retrieval call binding the contract method 0x18a1c4b6.
//
// Solidity: function debtLimit() view returns(uint256)
func (_Yvault022 *Yvault022Caller) DebtLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault022.contract.Call(opts, &out, "debtLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtLimit is a free data retrieval call binding the contract method 0x18a1c4b6.
//
// Solidity: function debtLimit() view returns(uint256)
func (_Yvault022 *Yvault022Session) DebtLimit() (*big.Int, error) {
	return _Yvault022.Contract.DebtLimit(&_Yvault022.CallOpts)
}

// DebtLimit is a free data retrieval call binding the contract method 0x18a1c4b6.
//
// Solidity: function debtLimit() view returns(uint256)
func (_Yvault022 *Yvault022CallerSession) DebtLimit() (*big.Int, error) {
	return _Yvault022.Contract.DebtLimit(&_Yvault022.CallOpts)
}

// DebtOutstanding is a free data retrieval call binding the contract method 0xbf3759b5.
//
// Solidity: function debtOutstanding() view returns(uint256)
func (_Yvault022 *Yvault022Caller) DebtOutstanding(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault022.contract.Call(opts, &out, "debtOutstanding")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtOutstanding is a free data retrieval call binding the contract method 0xbf3759b5.
//
// Solidity: function debtOutstanding() view returns(uint256)
func (_Yvault022 *Yvault022Session) DebtOutstanding() (*big.Int, error) {
	return _Yvault022.Contract.DebtOutstanding(&_Yvault022.CallOpts)
}

// DebtOutstanding is a free data retrieval call binding the contract method 0xbf3759b5.
//
// Solidity: function debtOutstanding() view returns(uint256)
func (_Yvault022 *Yvault022CallerSession) DebtOutstanding() (*big.Int, error) {
	return _Yvault022.Contract.DebtOutstanding(&_Yvault022.CallOpts)
}

// DebtOutstanding0 is a free data retrieval call binding the contract method 0xbdcf36bb.
//
// Solidity: function debtOutstanding(address _strategy) view returns(uint256)
func (_Yvault022 *Yvault022Caller) DebtOutstanding0(opts *bind.CallOpts, _strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yvault022.contract.Call(opts, &out, "debtOutstanding0", _strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtOutstanding0 is a free data retrieval call binding the contract method 0xbdcf36bb.
//
// Solidity: function debtOutstanding(address _strategy) view returns(uint256)
func (_Yvault022 *Yvault022Session) DebtOutstanding0(_strategy common.Address) (*big.Int, error) {
	return _Yvault022.Contract.DebtOutstanding0(&_Yvault022.CallOpts, _strategy)
}

// DebtOutstanding0 is a free data retrieval call binding the contract method 0xbdcf36bb.
//
// Solidity: function debtOutstanding(address _strategy) view returns(uint256)
func (_Yvault022 *Yvault022CallerSession) DebtOutstanding0(_strategy common.Address) (*big.Int, error) {
	return _Yvault022.Contract.DebtOutstanding0(&_Yvault022.CallOpts, _strategy)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Yvault022 *Yvault022Caller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault022.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Yvault022 *Yvault022Session) Decimals() (*big.Int, error) {
	return _Yvault022.Contract.Decimals(&_Yvault022.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Yvault022 *Yvault022CallerSession) Decimals() (*big.Int, error) {
	return _Yvault022.Contract.Decimals(&_Yvault022.CallOpts)
}

// DepositLimit is a free data retrieval call binding the contract method 0xecf70858.
//
// Solidity: function depositLimit() view returns(uint256)
func (_Yvault022 *Yvault022Caller) DepositLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault022.contract.Call(opts, &out, "depositLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositLimit is a free data retrieval call binding the contract method 0xecf70858.
//
// Solidity: function depositLimit() view returns(uint256)
func (_Yvault022 *Yvault022Session) DepositLimit() (*big.Int, error) {
	return _Yvault022.Contract.DepositLimit(&_Yvault022.CallOpts)
}

// DepositLimit is a free data retrieval call binding the contract method 0xecf70858.
//
// Solidity: function depositLimit() view returns(uint256)
func (_Yvault022 *Yvault022CallerSession) DepositLimit() (*big.Int, error) {
	return _Yvault022.Contract.DepositLimit(&_Yvault022.CallOpts)
}

// EmergencyShutdown is a free data retrieval call binding the contract method 0x3403c2fc.
//
// Solidity: function emergencyShutdown() view returns(bool)
func (_Yvault022 *Yvault022Caller) EmergencyShutdown(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Yvault022.contract.Call(opts, &out, "emergencyShutdown")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EmergencyShutdown is a free data retrieval call binding the contract method 0x3403c2fc.
//
// Solidity: function emergencyShutdown() view returns(bool)
func (_Yvault022 *Yvault022Session) EmergencyShutdown() (bool, error) {
	return _Yvault022.Contract.EmergencyShutdown(&_Yvault022.CallOpts)
}

// EmergencyShutdown is a free data retrieval call binding the contract method 0x3403c2fc.
//
// Solidity: function emergencyShutdown() view returns(bool)
func (_Yvault022 *Yvault022CallerSession) EmergencyShutdown() (bool, error) {
	return _Yvault022.Contract.EmergencyShutdown(&_Yvault022.CallOpts)
}

// ExpectedReturn is a free data retrieval call binding the contract method 0xd3406abd.
//
// Solidity: function expectedReturn() view returns(uint256)
func (_Yvault022 *Yvault022Caller) ExpectedReturn(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault022.contract.Call(opts, &out, "expectedReturn")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpectedReturn is a free data retrieval call binding the contract method 0xd3406abd.
//
// Solidity: function expectedReturn() view returns(uint256)
func (_Yvault022 *Yvault022Session) ExpectedReturn() (*big.Int, error) {
	return _Yvault022.Contract.ExpectedReturn(&_Yvault022.CallOpts)
}

// ExpectedReturn is a free data retrieval call binding the contract method 0xd3406abd.
//
// Solidity: function expectedReturn() view returns(uint256)
func (_Yvault022 *Yvault022CallerSession) ExpectedReturn() (*big.Int, error) {
	return _Yvault022.Contract.ExpectedReturn(&_Yvault022.CallOpts)
}

// ExpectedReturn0 is a free data retrieval call binding the contract method 0x33586b67.
//
// Solidity: function expectedReturn(address _strategy) view returns(uint256)
func (_Yvault022 *Yvault022Caller) ExpectedReturn0(opts *bind.CallOpts, _strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yvault022.contract.Call(opts, &out, "expectedReturn0", _strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpectedReturn0 is a free data retrieval call binding the contract method 0x33586b67.
//
// Solidity: function expectedReturn(address _strategy) view returns(uint256)
func (_Yvault022 *Yvault022Session) ExpectedReturn0(_strategy common.Address) (*big.Int, error) {
	return _Yvault022.Contract.ExpectedReturn0(&_Yvault022.CallOpts, _strategy)
}

// ExpectedReturn0 is a free data retrieval call binding the contract method 0x33586b67.
//
// Solidity: function expectedReturn(address _strategy) view returns(uint256)
func (_Yvault022 *Yvault022CallerSession) ExpectedReturn0(_strategy common.Address) (*big.Int, error) {
	return _Yvault022.Contract.ExpectedReturn0(&_Yvault022.CallOpts, _strategy)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Yvault022 *Yvault022Caller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yvault022.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Yvault022 *Yvault022Session) Governance() (common.Address, error) {
	return _Yvault022.Contract.Governance(&_Yvault022.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Yvault022 *Yvault022CallerSession) Governance() (common.Address, error) {
	return _Yvault022.Contract.Governance(&_Yvault022.CallOpts)
}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_Yvault022 *Yvault022Caller) Guardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yvault022.contract.Call(opts, &out, "guardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_Yvault022 *Yvault022Session) Guardian() (common.Address, error) {
	return _Yvault022.Contract.Guardian(&_Yvault022.CallOpts)
}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_Yvault022 *Yvault022CallerSession) Guardian() (common.Address, error) {
	return _Yvault022.Contract.Guardian(&_Yvault022.CallOpts)
}

// GuestList is a free data retrieval call binding the contract method 0x46d55875.
//
// Solidity: function guestList() view returns(address)
func (_Yvault022 *Yvault022Caller) GuestList(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yvault022.contract.Call(opts, &out, "guestList")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GuestList is a free data retrieval call binding the contract method 0x46d55875.
//
// Solidity: function guestList() view returns(address)
func (_Yvault022 *Yvault022Session) GuestList() (common.Address, error) {
	return _Yvault022.Contract.GuestList(&_Yvault022.CallOpts)
}

// GuestList is a free data retrieval call binding the contract method 0x46d55875.
//
// Solidity: function guestList() view returns(address)
func (_Yvault022 *Yvault022CallerSession) GuestList() (common.Address, error) {
	return _Yvault022.Contract.GuestList(&_Yvault022.CallOpts)
}

// LastReport is a free data retrieval call binding the contract method 0xc3535b52.
//
// Solidity: function lastReport() view returns(uint256)
func (_Yvault022 *Yvault022Caller) LastReport(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault022.contract.Call(opts, &out, "lastReport")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastReport is a free data retrieval call binding the contract method 0xc3535b52.
//
// Solidity: function lastReport() view returns(uint256)
func (_Yvault022 *Yvault022Session) LastReport() (*big.Int, error) {
	return _Yvault022.Contract.LastReport(&_Yvault022.CallOpts)
}

// LastReport is a free data retrieval call binding the contract method 0xc3535b52.
//
// Solidity: function lastReport() view returns(uint256)
func (_Yvault022 *Yvault022CallerSession) LastReport() (*big.Int, error) {
	return _Yvault022.Contract.LastReport(&_Yvault022.CallOpts)
}

// ManagementFee is a free data retrieval call binding the contract method 0xa6f7f5d6.
//
// Solidity: function managementFee() view returns(uint256)
func (_Yvault022 *Yvault022Caller) ManagementFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault022.contract.Call(opts, &out, "managementFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ManagementFee is a free data retrieval call binding the contract method 0xa6f7f5d6.
//
// Solidity: function managementFee() view returns(uint256)
func (_Yvault022 *Yvault022Session) ManagementFee() (*big.Int, error) {
	return _Yvault022.Contract.ManagementFee(&_Yvault022.CallOpts)
}

// ManagementFee is a free data retrieval call binding the contract method 0xa6f7f5d6.
//
// Solidity: function managementFee() view returns(uint256)
func (_Yvault022 *Yvault022CallerSession) ManagementFee() (*big.Int, error) {
	return _Yvault022.Contract.ManagementFee(&_Yvault022.CallOpts)
}

// MaxAvailableShares is a free data retrieval call binding the contract method 0x75de2902.
//
// Solidity: function maxAvailableShares() view returns(uint256)
func (_Yvault022 *Yvault022Caller) MaxAvailableShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault022.contract.Call(opts, &out, "maxAvailableShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAvailableShares is a free data retrieval call binding the contract method 0x75de2902.
//
// Solidity: function maxAvailableShares() view returns(uint256)
func (_Yvault022 *Yvault022Session) MaxAvailableShares() (*big.Int, error) {
	return _Yvault022.Contract.MaxAvailableShares(&_Yvault022.CallOpts)
}

// MaxAvailableShares is a free data retrieval call binding the contract method 0x75de2902.
//
// Solidity: function maxAvailableShares() view returns(uint256)
func (_Yvault022 *Yvault022CallerSession) MaxAvailableShares() (*big.Int, error) {
	return _Yvault022.Contract.MaxAvailableShares(&_Yvault022.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Yvault022 *Yvault022Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Yvault022.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Yvault022 *Yvault022Session) Name() (string, error) {
	return _Yvault022.Contract.Name(&_Yvault022.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Yvault022 *Yvault022CallerSession) Name() (string, error) {
	return _Yvault022.Contract.Name(&_Yvault022.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_Yvault022 *Yvault022Caller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yvault022.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_Yvault022 *Yvault022Session) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Yvault022.Contract.Nonces(&_Yvault022.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_Yvault022 *Yvault022CallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Yvault022.Contract.Nonces(&_Yvault022.CallOpts, arg0)
}

// PerformanceFee is a free data retrieval call binding the contract method 0x87788782.
//
// Solidity: function performanceFee() view returns(uint256)
func (_Yvault022 *Yvault022Caller) PerformanceFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault022.contract.Call(opts, &out, "performanceFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PerformanceFee is a free data retrieval call binding the contract method 0x87788782.
//
// Solidity: function performanceFee() view returns(uint256)
func (_Yvault022 *Yvault022Session) PerformanceFee() (*big.Int, error) {
	return _Yvault022.Contract.PerformanceFee(&_Yvault022.CallOpts)
}

// PerformanceFee is a free data retrieval call binding the contract method 0x87788782.
//
// Solidity: function performanceFee() view returns(uint256)
func (_Yvault022 *Yvault022CallerSession) PerformanceFee() (*big.Int, error) {
	return _Yvault022.Contract.PerformanceFee(&_Yvault022.CallOpts)
}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_Yvault022 *Yvault022Caller) PricePerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault022.contract.Call(opts, &out, "pricePerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_Yvault022 *Yvault022Session) PricePerShare() (*big.Int, error) {
	return _Yvault022.Contract.PricePerShare(&_Yvault022.CallOpts)
}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_Yvault022 *Yvault022CallerSession) PricePerShare() (*big.Int, error) {
	return _Yvault022.Contract.PricePerShare(&_Yvault022.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x9ec5a894.
//
// Solidity: function rewards() view returns(address)
func (_Yvault022 *Yvault022Caller) Rewards(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yvault022.contract.Call(opts, &out, "rewards")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0x9ec5a894.
//
// Solidity: function rewards() view returns(address)
func (_Yvault022 *Yvault022Session) Rewards() (common.Address, error) {
	return _Yvault022.Contract.Rewards(&_Yvault022.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x9ec5a894.
//
// Solidity: function rewards() view returns(address)
func (_Yvault022 *Yvault022CallerSession) Rewards() (common.Address, error) {
	return _Yvault022.Contract.Rewards(&_Yvault022.CallOpts)
}

// Strategies is a free data retrieval call binding the contract method 0x39ebf823.
//
// Solidity: function strategies(address arg0) view returns(uint256 performanceFee, uint256 activation, uint256 debtLimit, uint256 rateLimit, uint256 lastReport, uint256 totalDebt, uint256 totalGain, uint256 totalLoss)
func (_Yvault022 *Yvault022Caller) Strategies(opts *bind.CallOpts, arg0 common.Address) (struct {
	PerformanceFee *big.Int
	Activation     *big.Int
	DebtLimit      *big.Int
	RateLimit      *big.Int
	LastReport     *big.Int
	TotalDebt      *big.Int
	TotalGain      *big.Int
	TotalLoss      *big.Int
}, error) {
	var out []interface{}
	err := _Yvault022.contract.Call(opts, &out, "strategies", arg0)

	outstruct := new(struct {
		PerformanceFee *big.Int
		Activation     *big.Int
		DebtLimit      *big.Int
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
	outstruct.DebtLimit = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.RateLimit = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LastReport = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.TotalDebt = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.TotalGain = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.TotalLoss = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Strategies is a free data retrieval call binding the contract method 0x39ebf823.
//
// Solidity: function strategies(address arg0) view returns(uint256 performanceFee, uint256 activation, uint256 debtLimit, uint256 rateLimit, uint256 lastReport, uint256 totalDebt, uint256 totalGain, uint256 totalLoss)
func (_Yvault022 *Yvault022Session) Strategies(arg0 common.Address) (struct {
	PerformanceFee *big.Int
	Activation     *big.Int
	DebtLimit      *big.Int
	RateLimit      *big.Int
	LastReport     *big.Int
	TotalDebt      *big.Int
	TotalGain      *big.Int
	TotalLoss      *big.Int
}, error) {
	return _Yvault022.Contract.Strategies(&_Yvault022.CallOpts, arg0)
}

// Strategies is a free data retrieval call binding the contract method 0x39ebf823.
//
// Solidity: function strategies(address arg0) view returns(uint256 performanceFee, uint256 activation, uint256 debtLimit, uint256 rateLimit, uint256 lastReport, uint256 totalDebt, uint256 totalGain, uint256 totalLoss)
func (_Yvault022 *Yvault022CallerSession) Strategies(arg0 common.Address) (struct {
	PerformanceFee *big.Int
	Activation     *big.Int
	DebtLimit      *big.Int
	RateLimit      *big.Int
	LastReport     *big.Int
	TotalDebt      *big.Int
	TotalGain      *big.Int
	TotalLoss      *big.Int
}, error) {
	return _Yvault022.Contract.Strategies(&_Yvault022.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Yvault022 *Yvault022Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Yvault022.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Yvault022 *Yvault022Session) Symbol() (string, error) {
	return _Yvault022.Contract.Symbol(&_Yvault022.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Yvault022 *Yvault022CallerSession) Symbol() (string, error) {
	return _Yvault022.Contract.Symbol(&_Yvault022.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Yvault022 *Yvault022Caller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yvault022.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Yvault022 *Yvault022Session) Token() (common.Address, error) {
	return _Yvault022.Contract.Token(&_Yvault022.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Yvault022 *Yvault022CallerSession) Token() (common.Address, error) {
	return _Yvault022.Contract.Token(&_Yvault022.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_Yvault022 *Yvault022Caller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault022.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_Yvault022 *Yvault022Session) TotalAssets() (*big.Int, error) {
	return _Yvault022.Contract.TotalAssets(&_Yvault022.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_Yvault022 *Yvault022CallerSession) TotalAssets() (*big.Int, error) {
	return _Yvault022.Contract.TotalAssets(&_Yvault022.CallOpts)
}

// TotalBalanceSheet is a free data retrieval call binding the contract method 0x1d324976.
//
// Solidity: function totalBalanceSheet(address[40] _strategies) view returns(uint256)
func (_Yvault022 *Yvault022Caller) TotalBalanceSheet(opts *bind.CallOpts, _strategies [40]common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yvault022.contract.Call(opts, &out, "totalBalanceSheet", _strategies)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBalanceSheet is a free data retrieval call binding the contract method 0x1d324976.
//
// Solidity: function totalBalanceSheet(address[40] _strategies) view returns(uint256)
func (_Yvault022 *Yvault022Session) TotalBalanceSheet(_strategies [40]common.Address) (*big.Int, error) {
	return _Yvault022.Contract.TotalBalanceSheet(&_Yvault022.CallOpts, _strategies)
}

// TotalBalanceSheet is a free data retrieval call binding the contract method 0x1d324976.
//
// Solidity: function totalBalanceSheet(address[40] _strategies) view returns(uint256)
func (_Yvault022 *Yvault022CallerSession) TotalBalanceSheet(_strategies [40]common.Address) (*big.Int, error) {
	return _Yvault022.Contract.TotalBalanceSheet(&_Yvault022.CallOpts, _strategies)
}

// TotalDebt is a free data retrieval call binding the contract method 0xfc7b9c18.
//
// Solidity: function totalDebt() view returns(uint256)
func (_Yvault022 *Yvault022Caller) TotalDebt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault022.contract.Call(opts, &out, "totalDebt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDebt is a free data retrieval call binding the contract method 0xfc7b9c18.
//
// Solidity: function totalDebt() view returns(uint256)
func (_Yvault022 *Yvault022Session) TotalDebt() (*big.Int, error) {
	return _Yvault022.Contract.TotalDebt(&_Yvault022.CallOpts)
}

// TotalDebt is a free data retrieval call binding the contract method 0xfc7b9c18.
//
// Solidity: function totalDebt() view returns(uint256)
func (_Yvault022 *Yvault022CallerSession) TotalDebt() (*big.Int, error) {
	return _Yvault022.Contract.TotalDebt(&_Yvault022.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Yvault022 *Yvault022Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yvault022.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Yvault022 *Yvault022Session) TotalSupply() (*big.Int, error) {
	return _Yvault022.Contract.TotalSupply(&_Yvault022.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Yvault022 *Yvault022CallerSession) TotalSupply() (*big.Int, error) {
	return _Yvault022.Contract.TotalSupply(&_Yvault022.CallOpts)
}

// WithdrawalQueue is a free data retrieval call binding the contract method 0xc822adda.
//
// Solidity: function withdrawalQueue(uint256 arg0) view returns(address)
func (_Yvault022 *Yvault022Caller) WithdrawalQueue(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Yvault022.contract.Call(opts, &out, "withdrawalQueue", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WithdrawalQueue is a free data retrieval call binding the contract method 0xc822adda.
//
// Solidity: function withdrawalQueue(uint256 arg0) view returns(address)
func (_Yvault022 *Yvault022Session) WithdrawalQueue(arg0 *big.Int) (common.Address, error) {
	return _Yvault022.Contract.WithdrawalQueue(&_Yvault022.CallOpts, arg0)
}

// WithdrawalQueue is a free data retrieval call binding the contract method 0xc822adda.
//
// Solidity: function withdrawalQueue(uint256 arg0) view returns(address)
func (_Yvault022 *Yvault022CallerSession) WithdrawalQueue(arg0 *big.Int) (common.Address, error) {
	return _Yvault022.Contract.WithdrawalQueue(&_Yvault022.CallOpts, arg0)
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_Yvault022 *Yvault022Transactor) AcceptGovernance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yvault022.contract.Transact(opts, "acceptGovernance")
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_Yvault022 *Yvault022Session) AcceptGovernance() (*types.Transaction, error) {
	return _Yvault022.Contract.AcceptGovernance(&_Yvault022.TransactOpts)
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_Yvault022 *Yvault022TransactorSession) AcceptGovernance() (*types.Transaction, error) {
	return _Yvault022.Contract.AcceptGovernance(&_Yvault022.TransactOpts)
}

// AddStrategy is a paid mutator transaction binding the contract method 0x0dd21b6c.
//
// Solidity: function addStrategy(address _strategy, uint256 _debtLimit, uint256 _rateLimit, uint256 _performanceFee) returns()
func (_Yvault022 *Yvault022Transactor) AddStrategy(opts *bind.TransactOpts, _strategy common.Address, _debtLimit *big.Int, _rateLimit *big.Int, _performanceFee *big.Int) (*types.Transaction, error) {
	return _Yvault022.contract.Transact(opts, "addStrategy", _strategy, _debtLimit, _rateLimit, _performanceFee)
}

// AddStrategy is a paid mutator transaction binding the contract method 0x0dd21b6c.
//
// Solidity: function addStrategy(address _strategy, uint256 _debtLimit, uint256 _rateLimit, uint256 _performanceFee) returns()
func (_Yvault022 *Yvault022Session) AddStrategy(_strategy common.Address, _debtLimit *big.Int, _rateLimit *big.Int, _performanceFee *big.Int) (*types.Transaction, error) {
	return _Yvault022.Contract.AddStrategy(&_Yvault022.TransactOpts, _strategy, _debtLimit, _rateLimit, _performanceFee)
}

// AddStrategy is a paid mutator transaction binding the contract method 0x0dd21b6c.
//
// Solidity: function addStrategy(address _strategy, uint256 _debtLimit, uint256 _rateLimit, uint256 _performanceFee) returns()
func (_Yvault022 *Yvault022TransactorSession) AddStrategy(_strategy common.Address, _debtLimit *big.Int, _rateLimit *big.Int, _performanceFee *big.Int) (*types.Transaction, error) {
	return _Yvault022.Contract.AddStrategy(&_Yvault022.TransactOpts, _strategy, _debtLimit, _rateLimit, _performanceFee)
}

// AddStrategyToQueue is a paid mutator transaction binding the contract method 0xf76e4caa.
//
// Solidity: function addStrategyToQueue(address _strategy) returns()
func (_Yvault022 *Yvault022Transactor) AddStrategyToQueue(opts *bind.TransactOpts, _strategy common.Address) (*types.Transaction, error) {
	return _Yvault022.contract.Transact(opts, "addStrategyToQueue", _strategy)
}

// AddStrategyToQueue is a paid mutator transaction binding the contract method 0xf76e4caa.
//
// Solidity: function addStrategyToQueue(address _strategy) returns()
func (_Yvault022 *Yvault022Session) AddStrategyToQueue(_strategy common.Address) (*types.Transaction, error) {
	return _Yvault022.Contract.AddStrategyToQueue(&_Yvault022.TransactOpts, _strategy)
}

// AddStrategyToQueue is a paid mutator transaction binding the contract method 0xf76e4caa.
//
// Solidity: function addStrategyToQueue(address _strategy) returns()
func (_Yvault022 *Yvault022TransactorSession) AddStrategyToQueue(_strategy common.Address) (*types.Transaction, error) {
	return _Yvault022.Contract.AddStrategyToQueue(&_Yvault022.TransactOpts, _strategy)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Yvault022 *Yvault022Transactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Yvault022.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Yvault022 *Yvault022Session) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Yvault022.Contract.Approve(&_Yvault022.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Yvault022 *Yvault022TransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Yvault022.Contract.Approve(&_Yvault022.TransactOpts, _spender, _value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _value) returns(bool)
func (_Yvault022 *Yvault022Transactor) DecreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Yvault022.contract.Transact(opts, "decreaseAllowance", _spender, _value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _value) returns(bool)
func (_Yvault022 *Yvault022Session) DecreaseAllowance(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Yvault022.Contract.DecreaseAllowance(&_Yvault022.TransactOpts, _spender, _value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _value) returns(bool)
func (_Yvault022 *Yvault022TransactorSession) DecreaseAllowance(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Yvault022.Contract.DecreaseAllowance(&_Yvault022.TransactOpts, _spender, _value)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(uint256)
func (_Yvault022 *Yvault022Transactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yvault022.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(uint256)
func (_Yvault022 *Yvault022Session) Deposit() (*types.Transaction, error) {
	return _Yvault022.Contract.Deposit(&_Yvault022.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(uint256)
func (_Yvault022 *Yvault022TransactorSession) Deposit() (*types.Transaction, error) {
	return _Yvault022.Contract.Deposit(&_Yvault022.TransactOpts)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns(uint256)
func (_Yvault022 *Yvault022Transactor) Deposit0(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Yvault022.contract.Transact(opts, "deposit0", _amount)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns(uint256)
func (_Yvault022 *Yvault022Session) Deposit0(_amount *big.Int) (*types.Transaction, error) {
	return _Yvault022.Contract.Deposit0(&_Yvault022.TransactOpts, _amount)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns(uint256)
func (_Yvault022 *Yvault022TransactorSession) Deposit0(_amount *big.Int) (*types.Transaction, error) {
	return _Yvault022.Contract.Deposit0(&_Yvault022.TransactOpts, _amount)
}

// Deposit1 is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _amount, address _recipient) returns(uint256)
func (_Yvault022 *Yvault022Transactor) Deposit1(opts *bind.TransactOpts, _amount *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _Yvault022.contract.Transact(opts, "deposit1", _amount, _recipient)
}

// Deposit1 is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _amount, address _recipient) returns(uint256)
func (_Yvault022 *Yvault022Session) Deposit1(_amount *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _Yvault022.Contract.Deposit1(&_Yvault022.TransactOpts, _amount, _recipient)
}

// Deposit1 is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _amount, address _recipient) returns(uint256)
func (_Yvault022 *Yvault022TransactorSession) Deposit1(_amount *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _Yvault022.Contract.Deposit1(&_Yvault022.TransactOpts, _amount, _recipient)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _value) returns(bool)
func (_Yvault022 *Yvault022Transactor) IncreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Yvault022.contract.Transact(opts, "increaseAllowance", _spender, _value)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _value) returns(bool)
func (_Yvault022 *Yvault022Session) IncreaseAllowance(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Yvault022.Contract.IncreaseAllowance(&_Yvault022.TransactOpts, _spender, _value)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _value) returns(bool)
func (_Yvault022 *Yvault022TransactorSession) IncreaseAllowance(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Yvault022.Contract.IncreaseAllowance(&_Yvault022.TransactOpts, _spender, _value)
}

// MigrateStrategy is a paid mutator transaction binding the contract method 0x6cb56d19.
//
// Solidity: function migrateStrategy(address _oldVersion, address _newVersion) returns()
func (_Yvault022 *Yvault022Transactor) MigrateStrategy(opts *bind.TransactOpts, _oldVersion common.Address, _newVersion common.Address) (*types.Transaction, error) {
	return _Yvault022.contract.Transact(opts, "migrateStrategy", _oldVersion, _newVersion)
}

// MigrateStrategy is a paid mutator transaction binding the contract method 0x6cb56d19.
//
// Solidity: function migrateStrategy(address _oldVersion, address _newVersion) returns()
func (_Yvault022 *Yvault022Session) MigrateStrategy(_oldVersion common.Address, _newVersion common.Address) (*types.Transaction, error) {
	return _Yvault022.Contract.MigrateStrategy(&_Yvault022.TransactOpts, _oldVersion, _newVersion)
}

// MigrateStrategy is a paid mutator transaction binding the contract method 0x6cb56d19.
//
// Solidity: function migrateStrategy(address _oldVersion, address _newVersion) returns()
func (_Yvault022 *Yvault022TransactorSession) MigrateStrategy(_oldVersion common.Address, _newVersion common.Address) (*types.Transaction, error) {
	return _Yvault022.Contract.MigrateStrategy(&_Yvault022.TransactOpts, _oldVersion, _newVersion)
}

// Permit is a paid mutator transaction binding the contract method 0x9fd5a6cf.
//
// Solidity: function permit(address owner, address spender, uint256 amount, uint256 expiry, bytes signature) returns(bool)
func (_Yvault022 *Yvault022Transactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, amount *big.Int, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _Yvault022.contract.Transact(opts, "permit", owner, spender, amount, expiry, signature)
}

// Permit is a paid mutator transaction binding the contract method 0x9fd5a6cf.
//
// Solidity: function permit(address owner, address spender, uint256 amount, uint256 expiry, bytes signature) returns(bool)
func (_Yvault022 *Yvault022Session) Permit(owner common.Address, spender common.Address, amount *big.Int, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _Yvault022.Contract.Permit(&_Yvault022.TransactOpts, owner, spender, amount, expiry, signature)
}

// Permit is a paid mutator transaction binding the contract method 0x9fd5a6cf.
//
// Solidity: function permit(address owner, address spender, uint256 amount, uint256 expiry, bytes signature) returns(bool)
func (_Yvault022 *Yvault022TransactorSession) Permit(owner common.Address, spender common.Address, amount *big.Int, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _Yvault022.Contract.Permit(&_Yvault022.TransactOpts, owner, spender, amount, expiry, signature)
}

// RemoveStrategyFromQueue is a paid mutator transaction binding the contract method 0xb22439f5.
//
// Solidity: function removeStrategyFromQueue(address _strategy) returns()
func (_Yvault022 *Yvault022Transactor) RemoveStrategyFromQueue(opts *bind.TransactOpts, _strategy common.Address) (*types.Transaction, error) {
	return _Yvault022.contract.Transact(opts, "removeStrategyFromQueue", _strategy)
}

// RemoveStrategyFromQueue is a paid mutator transaction binding the contract method 0xb22439f5.
//
// Solidity: function removeStrategyFromQueue(address _strategy) returns()
func (_Yvault022 *Yvault022Session) RemoveStrategyFromQueue(_strategy common.Address) (*types.Transaction, error) {
	return _Yvault022.Contract.RemoveStrategyFromQueue(&_Yvault022.TransactOpts, _strategy)
}

// RemoveStrategyFromQueue is a paid mutator transaction binding the contract method 0xb22439f5.
//
// Solidity: function removeStrategyFromQueue(address _strategy) returns()
func (_Yvault022 *Yvault022TransactorSession) RemoveStrategyFromQueue(_strategy common.Address) (*types.Transaction, error) {
	return _Yvault022.Contract.RemoveStrategyFromQueue(&_Yvault022.TransactOpts, _strategy)
}

// Report is a paid mutator transaction binding the contract method 0xa1d9bafc.
//
// Solidity: function report(uint256 _gain, uint256 _loss, uint256 _debtPayment) returns(uint256)
func (_Yvault022 *Yvault022Transactor) Report(opts *bind.TransactOpts, _gain *big.Int, _loss *big.Int, _debtPayment *big.Int) (*types.Transaction, error) {
	return _Yvault022.contract.Transact(opts, "report", _gain, _loss, _debtPayment)
}

// Report is a paid mutator transaction binding the contract method 0xa1d9bafc.
//
// Solidity: function report(uint256 _gain, uint256 _loss, uint256 _debtPayment) returns(uint256)
func (_Yvault022 *Yvault022Session) Report(_gain *big.Int, _loss *big.Int, _debtPayment *big.Int) (*types.Transaction, error) {
	return _Yvault022.Contract.Report(&_Yvault022.TransactOpts, _gain, _loss, _debtPayment)
}

// Report is a paid mutator transaction binding the contract method 0xa1d9bafc.
//
// Solidity: function report(uint256 _gain, uint256 _loss, uint256 _debtPayment) returns(uint256)
func (_Yvault022 *Yvault022TransactorSession) Report(_gain *big.Int, _loss *big.Int, _debtPayment *big.Int) (*types.Transaction, error) {
	return _Yvault022.Contract.Report(&_Yvault022.TransactOpts, _gain, _loss, _debtPayment)
}

// RevokeStrategy is a paid mutator transaction binding the contract method 0xa0e4af9a.
//
// Solidity: function revokeStrategy() returns()
func (_Yvault022 *Yvault022Transactor) RevokeStrategy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yvault022.contract.Transact(opts, "revokeStrategy")
}

// RevokeStrategy is a paid mutator transaction binding the contract method 0xa0e4af9a.
//
// Solidity: function revokeStrategy() returns()
func (_Yvault022 *Yvault022Session) RevokeStrategy() (*types.Transaction, error) {
	return _Yvault022.Contract.RevokeStrategy(&_Yvault022.TransactOpts)
}

// RevokeStrategy is a paid mutator transaction binding the contract method 0xa0e4af9a.
//
// Solidity: function revokeStrategy() returns()
func (_Yvault022 *Yvault022TransactorSession) RevokeStrategy() (*types.Transaction, error) {
	return _Yvault022.Contract.RevokeStrategy(&_Yvault022.TransactOpts)
}

// RevokeStrategy0 is a paid mutator transaction binding the contract method 0xbb994d48.
//
// Solidity: function revokeStrategy(address _strategy) returns()
func (_Yvault022 *Yvault022Transactor) RevokeStrategy0(opts *bind.TransactOpts, _strategy common.Address) (*types.Transaction, error) {
	return _Yvault022.contract.Transact(opts, "revokeStrategy0", _strategy)
}

// RevokeStrategy0 is a paid mutator transaction binding the contract method 0xbb994d48.
//
// Solidity: function revokeStrategy(address _strategy) returns()
func (_Yvault022 *Yvault022Session) RevokeStrategy0(_strategy common.Address) (*types.Transaction, error) {
	return _Yvault022.Contract.RevokeStrategy0(&_Yvault022.TransactOpts, _strategy)
}

// RevokeStrategy0 is a paid mutator transaction binding the contract method 0xbb994d48.
//
// Solidity: function revokeStrategy(address _strategy) returns()
func (_Yvault022 *Yvault022TransactorSession) RevokeStrategy0(_strategy common.Address) (*types.Transaction, error) {
	return _Yvault022.Contract.RevokeStrategy0(&_Yvault022.TransactOpts, _strategy)
}

// SetDepositLimit is a paid mutator transaction binding the contract method 0xbdc8144b.
//
// Solidity: function setDepositLimit(uint256 _limit) returns()
func (_Yvault022 *Yvault022Transactor) SetDepositLimit(opts *bind.TransactOpts, _limit *big.Int) (*types.Transaction, error) {
	return _Yvault022.contract.Transact(opts, "setDepositLimit", _limit)
}

// SetDepositLimit is a paid mutator transaction binding the contract method 0xbdc8144b.
//
// Solidity: function setDepositLimit(uint256 _limit) returns()
func (_Yvault022 *Yvault022Session) SetDepositLimit(_limit *big.Int) (*types.Transaction, error) {
	return _Yvault022.Contract.SetDepositLimit(&_Yvault022.TransactOpts, _limit)
}

// SetDepositLimit is a paid mutator transaction binding the contract method 0xbdc8144b.
//
// Solidity: function setDepositLimit(uint256 _limit) returns()
func (_Yvault022 *Yvault022TransactorSession) SetDepositLimit(_limit *big.Int) (*types.Transaction, error) {
	return _Yvault022.Contract.SetDepositLimit(&_Yvault022.TransactOpts, _limit)
}

// SetEmergencyShutdown is a paid mutator transaction binding the contract method 0x14c64402.
//
// Solidity: function setEmergencyShutdown(bool _active) returns()
func (_Yvault022 *Yvault022Transactor) SetEmergencyShutdown(opts *bind.TransactOpts, _active bool) (*types.Transaction, error) {
	return _Yvault022.contract.Transact(opts, "setEmergencyShutdown", _active)
}

// SetEmergencyShutdown is a paid mutator transaction binding the contract method 0x14c64402.
//
// Solidity: function setEmergencyShutdown(bool _active) returns()
func (_Yvault022 *Yvault022Session) SetEmergencyShutdown(_active bool) (*types.Transaction, error) {
	return _Yvault022.Contract.SetEmergencyShutdown(&_Yvault022.TransactOpts, _active)
}

// SetEmergencyShutdown is a paid mutator transaction binding the contract method 0x14c64402.
//
// Solidity: function setEmergencyShutdown(bool _active) returns()
func (_Yvault022 *Yvault022TransactorSession) SetEmergencyShutdown(_active bool) (*types.Transaction, error) {
	return _Yvault022.Contract.SetEmergencyShutdown(&_Yvault022.TransactOpts, _active)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address _governance) returns()
func (_Yvault022 *Yvault022Transactor) SetGovernance(opts *bind.TransactOpts, _governance common.Address) (*types.Transaction, error) {
	return _Yvault022.contract.Transact(opts, "setGovernance", _governance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address _governance) returns()
func (_Yvault022 *Yvault022Session) SetGovernance(_governance common.Address) (*types.Transaction, error) {
	return _Yvault022.Contract.SetGovernance(&_Yvault022.TransactOpts, _governance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address _governance) returns()
func (_Yvault022 *Yvault022TransactorSession) SetGovernance(_governance common.Address) (*types.Transaction, error) {
	return _Yvault022.Contract.SetGovernance(&_Yvault022.TransactOpts, _governance)
}

// SetGuardian is a paid mutator transaction binding the contract method 0x8a0dac4a.
//
// Solidity: function setGuardian(address _guardian) returns()
func (_Yvault022 *Yvault022Transactor) SetGuardian(opts *bind.TransactOpts, _guardian common.Address) (*types.Transaction, error) {
	return _Yvault022.contract.Transact(opts, "setGuardian", _guardian)
}

// SetGuardian is a paid mutator transaction binding the contract method 0x8a0dac4a.
//
// Solidity: function setGuardian(address _guardian) returns()
func (_Yvault022 *Yvault022Session) SetGuardian(_guardian common.Address) (*types.Transaction, error) {
	return _Yvault022.Contract.SetGuardian(&_Yvault022.TransactOpts, _guardian)
}

// SetGuardian is a paid mutator transaction binding the contract method 0x8a0dac4a.
//
// Solidity: function setGuardian(address _guardian) returns()
func (_Yvault022 *Yvault022TransactorSession) SetGuardian(_guardian common.Address) (*types.Transaction, error) {
	return _Yvault022.Contract.SetGuardian(&_Yvault022.TransactOpts, _guardian)
}

// SetGuestList is a paid mutator transaction binding the contract method 0x0b5b78eb.
//
// Solidity: function setGuestList(address _guestList) returns()
func (_Yvault022 *Yvault022Transactor) SetGuestList(opts *bind.TransactOpts, _guestList common.Address) (*types.Transaction, error) {
	return _Yvault022.contract.Transact(opts, "setGuestList", _guestList)
}

// SetGuestList is a paid mutator transaction binding the contract method 0x0b5b78eb.
//
// Solidity: function setGuestList(address _guestList) returns()
func (_Yvault022 *Yvault022Session) SetGuestList(_guestList common.Address) (*types.Transaction, error) {
	return _Yvault022.Contract.SetGuestList(&_Yvault022.TransactOpts, _guestList)
}

// SetGuestList is a paid mutator transaction binding the contract method 0x0b5b78eb.
//
// Solidity: function setGuestList(address _guestList) returns()
func (_Yvault022 *Yvault022TransactorSession) SetGuestList(_guestList common.Address) (*types.Transaction, error) {
	return _Yvault022.Contract.SetGuestList(&_Yvault022.TransactOpts, _guestList)
}

// SetManagementFee is a paid mutator transaction binding the contract method 0xfe56e232.
//
// Solidity: function setManagementFee(uint256 _fee) returns()
func (_Yvault022 *Yvault022Transactor) SetManagementFee(opts *bind.TransactOpts, _fee *big.Int) (*types.Transaction, error) {
	return _Yvault022.contract.Transact(opts, "setManagementFee", _fee)
}

// SetManagementFee is a paid mutator transaction binding the contract method 0xfe56e232.
//
// Solidity: function setManagementFee(uint256 _fee) returns()
func (_Yvault022 *Yvault022Session) SetManagementFee(_fee *big.Int) (*types.Transaction, error) {
	return _Yvault022.Contract.SetManagementFee(&_Yvault022.TransactOpts, _fee)
}

// SetManagementFee is a paid mutator transaction binding the contract method 0xfe56e232.
//
// Solidity: function setManagementFee(uint256 _fee) returns()
func (_Yvault022 *Yvault022TransactorSession) SetManagementFee(_fee *big.Int) (*types.Transaction, error) {
	return _Yvault022.Contract.SetManagementFee(&_Yvault022.TransactOpts, _fee)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string _name) returns()
func (_Yvault022 *Yvault022Transactor) SetName(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _Yvault022.contract.Transact(opts, "setName", _name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string _name) returns()
func (_Yvault022 *Yvault022Session) SetName(_name string) (*types.Transaction, error) {
	return _Yvault022.Contract.SetName(&_Yvault022.TransactOpts, _name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string _name) returns()
func (_Yvault022 *Yvault022TransactorSession) SetName(_name string) (*types.Transaction, error) {
	return _Yvault022.Contract.SetName(&_Yvault022.TransactOpts, _name)
}

// SetPerformanceFee is a paid mutator transaction binding the contract method 0x70897b23.
//
// Solidity: function setPerformanceFee(uint256 _fee) returns()
func (_Yvault022 *Yvault022Transactor) SetPerformanceFee(opts *bind.TransactOpts, _fee *big.Int) (*types.Transaction, error) {
	return _Yvault022.contract.Transact(opts, "setPerformanceFee", _fee)
}

// SetPerformanceFee is a paid mutator transaction binding the contract method 0x70897b23.
//
// Solidity: function setPerformanceFee(uint256 _fee) returns()
func (_Yvault022 *Yvault022Session) SetPerformanceFee(_fee *big.Int) (*types.Transaction, error) {
	return _Yvault022.Contract.SetPerformanceFee(&_Yvault022.TransactOpts, _fee)
}

// SetPerformanceFee is a paid mutator transaction binding the contract method 0x70897b23.
//
// Solidity: function setPerformanceFee(uint256 _fee) returns()
func (_Yvault022 *Yvault022TransactorSession) SetPerformanceFee(_fee *big.Int) (*types.Transaction, error) {
	return _Yvault022.Contract.SetPerformanceFee(&_Yvault022.TransactOpts, _fee)
}

// SetRewards is a paid mutator transaction binding the contract method 0xec38a862.
//
// Solidity: function setRewards(address _rewards) returns()
func (_Yvault022 *Yvault022Transactor) SetRewards(opts *bind.TransactOpts, _rewards common.Address) (*types.Transaction, error) {
	return _Yvault022.contract.Transact(opts, "setRewards", _rewards)
}

// SetRewards is a paid mutator transaction binding the contract method 0xec38a862.
//
// Solidity: function setRewards(address _rewards) returns()
func (_Yvault022 *Yvault022Session) SetRewards(_rewards common.Address) (*types.Transaction, error) {
	return _Yvault022.Contract.SetRewards(&_Yvault022.TransactOpts, _rewards)
}

// SetRewards is a paid mutator transaction binding the contract method 0xec38a862.
//
// Solidity: function setRewards(address _rewards) returns()
func (_Yvault022 *Yvault022TransactorSession) SetRewards(_rewards common.Address) (*types.Transaction, error) {
	return _Yvault022.Contract.SetRewards(&_Yvault022.TransactOpts, _rewards)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string _symbol) returns()
func (_Yvault022 *Yvault022Transactor) SetSymbol(opts *bind.TransactOpts, _symbol string) (*types.Transaction, error) {
	return _Yvault022.contract.Transact(opts, "setSymbol", _symbol)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string _symbol) returns()
func (_Yvault022 *Yvault022Session) SetSymbol(_symbol string) (*types.Transaction, error) {
	return _Yvault022.Contract.SetSymbol(&_Yvault022.TransactOpts, _symbol)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string _symbol) returns()
func (_Yvault022 *Yvault022TransactorSession) SetSymbol(_symbol string) (*types.Transaction, error) {
	return _Yvault022.Contract.SetSymbol(&_Yvault022.TransactOpts, _symbol)
}

// SetWithdrawalQueue is a paid mutator transaction binding the contract method 0x94148415.
//
// Solidity: function setWithdrawalQueue(address[20] _queue) returns()
func (_Yvault022 *Yvault022Transactor) SetWithdrawalQueue(opts *bind.TransactOpts, _queue [20]common.Address) (*types.Transaction, error) {
	return _Yvault022.contract.Transact(opts, "setWithdrawalQueue", _queue)
}

// SetWithdrawalQueue is a paid mutator transaction binding the contract method 0x94148415.
//
// Solidity: function setWithdrawalQueue(address[20] _queue) returns()
func (_Yvault022 *Yvault022Session) SetWithdrawalQueue(_queue [20]common.Address) (*types.Transaction, error) {
	return _Yvault022.Contract.SetWithdrawalQueue(&_Yvault022.TransactOpts, _queue)
}

// SetWithdrawalQueue is a paid mutator transaction binding the contract method 0x94148415.
//
// Solidity: function setWithdrawalQueue(address[20] _queue) returns()
func (_Yvault022 *Yvault022TransactorSession) SetWithdrawalQueue(_queue [20]common.Address) (*types.Transaction, error) {
	return _Yvault022.Contract.SetWithdrawalQueue(&_Yvault022.TransactOpts, _queue)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address _token) returns()
func (_Yvault022 *Yvault022Transactor) Sweep(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Yvault022.contract.Transact(opts, "sweep", _token)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address _token) returns()
func (_Yvault022 *Yvault022Session) Sweep(_token common.Address) (*types.Transaction, error) {
	return _Yvault022.Contract.Sweep(&_Yvault022.TransactOpts, _token)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address _token) returns()
func (_Yvault022 *Yvault022TransactorSession) Sweep(_token common.Address) (*types.Transaction, error) {
	return _Yvault022.Contract.Sweep(&_Yvault022.TransactOpts, _token)
}

// Sweep0 is a paid mutator transaction binding the contract method 0x6ea056a9.
//
// Solidity: function sweep(address _token, uint256 _value) returns()
func (_Yvault022 *Yvault022Transactor) Sweep0(opts *bind.TransactOpts, _token common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Yvault022.contract.Transact(opts, "sweep0", _token, _value)
}

// Sweep0 is a paid mutator transaction binding the contract method 0x6ea056a9.
//
// Solidity: function sweep(address _token, uint256 _value) returns()
func (_Yvault022 *Yvault022Session) Sweep0(_token common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Yvault022.Contract.Sweep0(&_Yvault022.TransactOpts, _token, _value)
}

// Sweep0 is a paid mutator transaction binding the contract method 0x6ea056a9.
//
// Solidity: function sweep(address _token, uint256 _value) returns()
func (_Yvault022 *Yvault022TransactorSession) Sweep0(_token common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Yvault022.Contract.Sweep0(&_Yvault022.TransactOpts, _token, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Yvault022 *Yvault022Transactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Yvault022.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Yvault022 *Yvault022Session) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Yvault022.Contract.Transfer(&_Yvault022.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Yvault022 *Yvault022TransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Yvault022.Contract.Transfer(&_Yvault022.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Yvault022 *Yvault022Transactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Yvault022.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Yvault022 *Yvault022Session) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Yvault022.Contract.TransferFrom(&_Yvault022.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Yvault022 *Yvault022TransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Yvault022.Contract.TransferFrom(&_Yvault022.TransactOpts, _from, _to, _value)
}

// UpdateStrategyDebtLimit is a paid mutator transaction binding the contract method 0xcd7d8f4f.
//
// Solidity: function updateStrategyDebtLimit(address _strategy, uint256 _debtLimit) returns()
func (_Yvault022 *Yvault022Transactor) UpdateStrategyDebtLimit(opts *bind.TransactOpts, _strategy common.Address, _debtLimit *big.Int) (*types.Transaction, error) {
	return _Yvault022.contract.Transact(opts, "updateStrategyDebtLimit", _strategy, _debtLimit)
}

// UpdateStrategyDebtLimit is a paid mutator transaction binding the contract method 0xcd7d8f4f.
//
// Solidity: function updateStrategyDebtLimit(address _strategy, uint256 _debtLimit) returns()
func (_Yvault022 *Yvault022Session) UpdateStrategyDebtLimit(_strategy common.Address, _debtLimit *big.Int) (*types.Transaction, error) {
	return _Yvault022.Contract.UpdateStrategyDebtLimit(&_Yvault022.TransactOpts, _strategy, _debtLimit)
}

// UpdateStrategyDebtLimit is a paid mutator transaction binding the contract method 0xcd7d8f4f.
//
// Solidity: function updateStrategyDebtLimit(address _strategy, uint256 _debtLimit) returns()
func (_Yvault022 *Yvault022TransactorSession) UpdateStrategyDebtLimit(_strategy common.Address, _debtLimit *big.Int) (*types.Transaction, error) {
	return _Yvault022.Contract.UpdateStrategyDebtLimit(&_Yvault022.TransactOpts, _strategy, _debtLimit)
}

// UpdateStrategyPerformanceFee is a paid mutator transaction binding the contract method 0xd0194ed6.
//
// Solidity: function updateStrategyPerformanceFee(address _strategy, uint256 _performanceFee) returns()
func (_Yvault022 *Yvault022Transactor) UpdateStrategyPerformanceFee(opts *bind.TransactOpts, _strategy common.Address, _performanceFee *big.Int) (*types.Transaction, error) {
	return _Yvault022.contract.Transact(opts, "updateStrategyPerformanceFee", _strategy, _performanceFee)
}

// UpdateStrategyPerformanceFee is a paid mutator transaction binding the contract method 0xd0194ed6.
//
// Solidity: function updateStrategyPerformanceFee(address _strategy, uint256 _performanceFee) returns()
func (_Yvault022 *Yvault022Session) UpdateStrategyPerformanceFee(_strategy common.Address, _performanceFee *big.Int) (*types.Transaction, error) {
	return _Yvault022.Contract.UpdateStrategyPerformanceFee(&_Yvault022.TransactOpts, _strategy, _performanceFee)
}

// UpdateStrategyPerformanceFee is a paid mutator transaction binding the contract method 0xd0194ed6.
//
// Solidity: function updateStrategyPerformanceFee(address _strategy, uint256 _performanceFee) returns()
func (_Yvault022 *Yvault022TransactorSession) UpdateStrategyPerformanceFee(_strategy common.Address, _performanceFee *big.Int) (*types.Transaction, error) {
	return _Yvault022.Contract.UpdateStrategyPerformanceFee(&_Yvault022.TransactOpts, _strategy, _performanceFee)
}

// UpdateStrategyRateLimit is a paid mutator transaction binding the contract method 0x62fdbc9f.
//
// Solidity: function updateStrategyRateLimit(address _strategy, uint256 _rateLimit) returns()
func (_Yvault022 *Yvault022Transactor) UpdateStrategyRateLimit(opts *bind.TransactOpts, _strategy common.Address, _rateLimit *big.Int) (*types.Transaction, error) {
	return _Yvault022.contract.Transact(opts, "updateStrategyRateLimit", _strategy, _rateLimit)
}

// UpdateStrategyRateLimit is a paid mutator transaction binding the contract method 0x62fdbc9f.
//
// Solidity: function updateStrategyRateLimit(address _strategy, uint256 _rateLimit) returns()
func (_Yvault022 *Yvault022Session) UpdateStrategyRateLimit(_strategy common.Address, _rateLimit *big.Int) (*types.Transaction, error) {
	return _Yvault022.Contract.UpdateStrategyRateLimit(&_Yvault022.TransactOpts, _strategy, _rateLimit)
}

// UpdateStrategyRateLimit is a paid mutator transaction binding the contract method 0x62fdbc9f.
//
// Solidity: function updateStrategyRateLimit(address _strategy, uint256 _rateLimit) returns()
func (_Yvault022 *Yvault022TransactorSession) UpdateStrategyRateLimit(_strategy common.Address, _rateLimit *big.Int) (*types.Transaction, error) {
	return _Yvault022.Contract.UpdateStrategyRateLimit(&_Yvault022.TransactOpts, _strategy, _rateLimit)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(uint256)
func (_Yvault022 *Yvault022Transactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yvault022.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(uint256)
func (_Yvault022 *Yvault022Session) Withdraw() (*types.Transaction, error) {
	return _Yvault022.Contract.Withdraw(&_Yvault022.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(uint256)
func (_Yvault022 *Yvault022TransactorSession) Withdraw() (*types.Transaction, error) {
	return _Yvault022.Contract.Withdraw(&_Yvault022.TransactOpts)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _shares) returns(uint256)
func (_Yvault022 *Yvault022Transactor) Withdraw0(opts *bind.TransactOpts, _shares *big.Int) (*types.Transaction, error) {
	return _Yvault022.contract.Transact(opts, "withdraw0", _shares)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _shares) returns(uint256)
func (_Yvault022 *Yvault022Session) Withdraw0(_shares *big.Int) (*types.Transaction, error) {
	return _Yvault022.Contract.Withdraw0(&_Yvault022.TransactOpts, _shares)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _shares) returns(uint256)
func (_Yvault022 *Yvault022TransactorSession) Withdraw0(_shares *big.Int) (*types.Transaction, error) {
	return _Yvault022.Contract.Withdraw0(&_Yvault022.TransactOpts, _shares)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 _shares, address _recipient) returns(uint256)
func (_Yvault022 *Yvault022Transactor) Withdraw1(opts *bind.TransactOpts, _shares *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _Yvault022.contract.Transact(opts, "withdraw1", _shares, _recipient)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 _shares, address _recipient) returns(uint256)
func (_Yvault022 *Yvault022Session) Withdraw1(_shares *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _Yvault022.Contract.Withdraw1(&_Yvault022.TransactOpts, _shares, _recipient)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 _shares, address _recipient) returns(uint256)
func (_Yvault022 *Yvault022TransactorSession) Withdraw1(_shares *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _Yvault022.Contract.Withdraw1(&_Yvault022.TransactOpts, _shares, _recipient)
}

// Yvault022ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Yvault022 contract.
type Yvault022ApprovalIterator struct {
	Event *Yvault022Approval // Event containing the contract specifics and raw log

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
func (it *Yvault022ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault022Approval)
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
		it.Event = new(Yvault022Approval)
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
func (it *Yvault022ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault022ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault022Approval represents a Approval event raised by the Yvault022 contract.
type Yvault022Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Yvault022 *Yvault022Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Yvault022ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Yvault022.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Yvault022ApprovalIterator{contract: _Yvault022.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Yvault022 *Yvault022Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Yvault022Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Yvault022.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault022Approval)
				if err := _Yvault022.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Yvault022 *Yvault022Filterer) ParseApproval(log types.Log) (*Yvault022Approval, error) {
	event := new(Yvault022Approval)
	if err := _Yvault022.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault022StrategyAddedIterator is returned from FilterStrategyAdded and is used to iterate over the raw logs and unpacked data for StrategyAdded events raised by the Yvault022 contract.
type Yvault022StrategyAddedIterator struct {
	Event *Yvault022StrategyAdded // Event containing the contract specifics and raw log

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
func (it *Yvault022StrategyAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault022StrategyAdded)
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
		it.Event = new(Yvault022StrategyAdded)
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
func (it *Yvault022StrategyAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault022StrategyAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault022StrategyAdded represents a StrategyAdded event raised by the Yvault022 contract.
type Yvault022StrategyAdded struct {
	Strategy       common.Address
	DebtLimit      *big.Int
	RateLimit      *big.Int
	PerformanceFee *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterStrategyAdded is a free log retrieval operation binding the contract event 0x5ec27a4fa537fc86d0d17d84e0ee3172c9d253c78cc4ab5c69ee99c5f7084f51.
//
// Solidity: event StrategyAdded(address indexed strategy, uint256 debtLimit, uint256 rateLimit, uint256 performanceFee)
func (_Yvault022 *Yvault022Filterer) FilterStrategyAdded(opts *bind.FilterOpts, strategy []common.Address) (*Yvault022StrategyAddedIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault022.contract.FilterLogs(opts, "StrategyAdded", strategyRule)
	if err != nil {
		return nil, err
	}
	return &Yvault022StrategyAddedIterator{contract: _Yvault022.contract, event: "StrategyAdded", logs: logs, sub: sub}, nil
}

// WatchStrategyAdded is a free log subscription operation binding the contract event 0x5ec27a4fa537fc86d0d17d84e0ee3172c9d253c78cc4ab5c69ee99c5f7084f51.
//
// Solidity: event StrategyAdded(address indexed strategy, uint256 debtLimit, uint256 rateLimit, uint256 performanceFee)
func (_Yvault022 *Yvault022Filterer) WatchStrategyAdded(opts *bind.WatchOpts, sink chan<- *Yvault022StrategyAdded, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault022.contract.WatchLogs(opts, "StrategyAdded", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault022StrategyAdded)
				if err := _Yvault022.contract.UnpackLog(event, "StrategyAdded", log); err != nil {
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
// Solidity: event StrategyAdded(address indexed strategy, uint256 debtLimit, uint256 rateLimit, uint256 performanceFee)
func (_Yvault022 *Yvault022Filterer) ParseStrategyAdded(log types.Log) (*Yvault022StrategyAdded, error) {
	event := new(Yvault022StrategyAdded)
	if err := _Yvault022.contract.UnpackLog(event, "StrategyAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault022StrategyReportedIterator is returned from FilterStrategyReported and is used to iterate over the raw logs and unpacked data for StrategyReported events raised by the Yvault022 contract.
type Yvault022StrategyReportedIterator struct {
	Event *Yvault022StrategyReported // Event containing the contract specifics and raw log

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
func (it *Yvault022StrategyReportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault022StrategyReported)
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
		it.Event = new(Yvault022StrategyReported)
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
func (it *Yvault022StrategyReportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault022StrategyReportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault022StrategyReported represents a StrategyReported event raised by the Yvault022 contract.
type Yvault022StrategyReported struct {
	Strategy  common.Address
	Gain      *big.Int
	Loss      *big.Int
	TotalGain *big.Int
	TotalLoss *big.Int
	TotalDebt *big.Int
	DebtAdded *big.Int
	DebtLimit *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStrategyReported is a free log retrieval operation binding the contract event 0x2fb611faf48b1d1b91edbba34cee10c6357adee410540e4a8f7a82b6b38673e4.
//
// Solidity: event StrategyReported(address indexed strategy, uint256 gain, uint256 loss, uint256 totalGain, uint256 totalLoss, uint256 totalDebt, uint256 debtAdded, uint256 debtLimit)
func (_Yvault022 *Yvault022Filterer) FilterStrategyReported(opts *bind.FilterOpts, strategy []common.Address) (*Yvault022StrategyReportedIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault022.contract.FilterLogs(opts, "StrategyReported", strategyRule)
	if err != nil {
		return nil, err
	}
	return &Yvault022StrategyReportedIterator{contract: _Yvault022.contract, event: "StrategyReported", logs: logs, sub: sub}, nil
}

// WatchStrategyReported is a free log subscription operation binding the contract event 0x2fb611faf48b1d1b91edbba34cee10c6357adee410540e4a8f7a82b6b38673e4.
//
// Solidity: event StrategyReported(address indexed strategy, uint256 gain, uint256 loss, uint256 totalGain, uint256 totalLoss, uint256 totalDebt, uint256 debtAdded, uint256 debtLimit)
func (_Yvault022 *Yvault022Filterer) WatchStrategyReported(opts *bind.WatchOpts, sink chan<- *Yvault022StrategyReported, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Yvault022.contract.WatchLogs(opts, "StrategyReported", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault022StrategyReported)
				if err := _Yvault022.contract.UnpackLog(event, "StrategyReported", log); err != nil {
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
// Solidity: event StrategyReported(address indexed strategy, uint256 gain, uint256 loss, uint256 totalGain, uint256 totalLoss, uint256 totalDebt, uint256 debtAdded, uint256 debtLimit)
func (_Yvault022 *Yvault022Filterer) ParseStrategyReported(log types.Log) (*Yvault022StrategyReported, error) {
	event := new(Yvault022StrategyReported)
	if err := _Yvault022.contract.UnpackLog(event, "StrategyReported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yvault022TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Yvault022 contract.
type Yvault022TransferIterator struct {
	Event *Yvault022Transfer // Event containing the contract specifics and raw log

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
func (it *Yvault022TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yvault022Transfer)
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
		it.Event = new(Yvault022Transfer)
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
func (it *Yvault022TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yvault022TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yvault022Transfer represents a Transfer event raised by the Yvault022 contract.
type Yvault022Transfer struct {
	Sender   common.Address
	Receiver common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_Yvault022 *Yvault022Filterer) FilterTransfer(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*Yvault022TransferIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Yvault022.contract.FilterLogs(opts, "Transfer", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &Yvault022TransferIterator{contract: _Yvault022.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_Yvault022 *Yvault022Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Yvault022Transfer, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Yvault022.contract.WatchLogs(opts, "Transfer", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yvault022Transfer)
				if err := _Yvault022.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Yvault022 *Yvault022Filterer) ParseTransfer(log types.Log) (*Yvault022Transfer, error) {
	event := new(Yvault022Transfer)
	if err := _Yvault022.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
