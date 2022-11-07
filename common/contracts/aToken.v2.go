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

// ATokenV2MetaData contains all meta data concerning the ATokenV2 contract.
var ATokenV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractILendingPool\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlyingAssetAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reserveTreasuryAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"incentivesController\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BalanceTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlyingAsset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"treasury\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"incentivesController\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"aTokenDecimals\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"aTokenName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"aTokenSymbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ATOKEN_REVISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EIP712_REVISION\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POOL\",\"outputs\":[{\"internalType\":\"contractILendingPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RESERVE_TREASURY_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UINT_MAX_VALUE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNDERLYING_ASSET_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiverOfUnderlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIncentivesController\",\"outputs\":[{\"internalType\":\"contractIAaveIncentivesController\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getScaledUserBalanceAndSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"underlyingAssetDecimals\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"mintToTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"scaledBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scaledTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferOnLiquidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferUnderlyingTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ATokenV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use ATokenV2MetaData.ABI instead.
var ATokenV2ABI = ATokenV2MetaData.ABI

// ATokenV2 is an auto generated Go binding around an Ethereum contract.
type ATokenV2 struct {
	ATokenV2Caller     // Read-only binding to the contract
	ATokenV2Transactor // Write-only binding to the contract
	ATokenV2Filterer   // Log filterer for contract events
}

// ATokenV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type ATokenV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ATokenV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ATokenV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ATokenV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ATokenV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ATokenV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ATokenV2Session struct {
	Contract     *ATokenV2         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ATokenV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ATokenV2CallerSession struct {
	Contract *ATokenV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ATokenV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ATokenV2TransactorSession struct {
	Contract     *ATokenV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ATokenV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type ATokenV2Raw struct {
	Contract *ATokenV2 // Generic contract binding to access the raw methods on
}

// ATokenV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ATokenV2CallerRaw struct {
	Contract *ATokenV2Caller // Generic read-only contract binding to access the raw methods on
}

// ATokenV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ATokenV2TransactorRaw struct {
	Contract *ATokenV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewATokenV2 creates a new instance of ATokenV2, bound to a specific deployed contract.
func NewATokenV2(address common.Address, backend bind.ContractBackend) (*ATokenV2, error) {
	contract, err := bindATokenV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ATokenV2{ATokenV2Caller: ATokenV2Caller{contract: contract}, ATokenV2Transactor: ATokenV2Transactor{contract: contract}, ATokenV2Filterer: ATokenV2Filterer{contract: contract}}, nil
}

// NewATokenV2Caller creates a new read-only instance of ATokenV2, bound to a specific deployed contract.
func NewATokenV2Caller(address common.Address, caller bind.ContractCaller) (*ATokenV2Caller, error) {
	contract, err := bindATokenV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ATokenV2Caller{contract: contract}, nil
}

// NewATokenV2Transactor creates a new write-only instance of ATokenV2, bound to a specific deployed contract.
func NewATokenV2Transactor(address common.Address, transactor bind.ContractTransactor) (*ATokenV2Transactor, error) {
	contract, err := bindATokenV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ATokenV2Transactor{contract: contract}, nil
}

// NewATokenV2Filterer creates a new log filterer instance of ATokenV2, bound to a specific deployed contract.
func NewATokenV2Filterer(address common.Address, filterer bind.ContractFilterer) (*ATokenV2Filterer, error) {
	contract, err := bindATokenV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ATokenV2Filterer{contract: contract}, nil
}

// bindATokenV2 binds a generic wrapper to an already deployed contract.
func bindATokenV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ATokenV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ATokenV2 *ATokenV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ATokenV2.Contract.ATokenV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ATokenV2 *ATokenV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ATokenV2.Contract.ATokenV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ATokenV2 *ATokenV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ATokenV2.Contract.ATokenV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ATokenV2 *ATokenV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ATokenV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ATokenV2 *ATokenV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ATokenV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ATokenV2 *ATokenV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ATokenV2.Contract.contract.Transact(opts, method, params...)
}

// ATOKENREVISION is a free data retrieval call binding the contract method 0x0bd7ad3b.
//
// Solidity: function ATOKEN_REVISION() view returns(uint256)
func (_ATokenV2 *ATokenV2Caller) ATOKENREVISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ATokenV2.contract.Call(opts, &out, "ATOKEN_REVISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ATOKENREVISION is a free data retrieval call binding the contract method 0x0bd7ad3b.
//
// Solidity: function ATOKEN_REVISION() view returns(uint256)
func (_ATokenV2 *ATokenV2Session) ATOKENREVISION() (*big.Int, error) {
	return _ATokenV2.Contract.ATOKENREVISION(&_ATokenV2.CallOpts)
}

// ATOKENREVISION is a free data retrieval call binding the contract method 0x0bd7ad3b.
//
// Solidity: function ATOKEN_REVISION() view returns(uint256)
func (_ATokenV2 *ATokenV2CallerSession) ATOKENREVISION() (*big.Int, error) {
	return _ATokenV2.Contract.ATOKENREVISION(&_ATokenV2.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_ATokenV2 *ATokenV2Caller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ATokenV2.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_ATokenV2 *ATokenV2Session) DOMAINSEPARATOR() ([32]byte, error) {
	return _ATokenV2.Contract.DOMAINSEPARATOR(&_ATokenV2.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_ATokenV2 *ATokenV2CallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _ATokenV2.Contract.DOMAINSEPARATOR(&_ATokenV2.CallOpts)
}

// EIP712REVISION is a free data retrieval call binding the contract method 0x78160376.
//
// Solidity: function EIP712_REVISION() view returns(bytes)
func (_ATokenV2 *ATokenV2Caller) EIP712REVISION(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _ATokenV2.contract.Call(opts, &out, "EIP712_REVISION")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EIP712REVISION is a free data retrieval call binding the contract method 0x78160376.
//
// Solidity: function EIP712_REVISION() view returns(bytes)
func (_ATokenV2 *ATokenV2Session) EIP712REVISION() ([]byte, error) {
	return _ATokenV2.Contract.EIP712REVISION(&_ATokenV2.CallOpts)
}

// EIP712REVISION is a free data retrieval call binding the contract method 0x78160376.
//
// Solidity: function EIP712_REVISION() view returns(bytes)
func (_ATokenV2 *ATokenV2CallerSession) EIP712REVISION() ([]byte, error) {
	return _ATokenV2.Contract.EIP712REVISION(&_ATokenV2.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_ATokenV2 *ATokenV2Caller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ATokenV2.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_ATokenV2 *ATokenV2Session) PERMITTYPEHASH() ([32]byte, error) {
	return _ATokenV2.Contract.PERMITTYPEHASH(&_ATokenV2.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_ATokenV2 *ATokenV2CallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _ATokenV2.Contract.PERMITTYPEHASH(&_ATokenV2.CallOpts)
}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_ATokenV2 *ATokenV2Caller) POOL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ATokenV2.contract.Call(opts, &out, "POOL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_ATokenV2 *ATokenV2Session) POOL() (common.Address, error) {
	return _ATokenV2.Contract.POOL(&_ATokenV2.CallOpts)
}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_ATokenV2 *ATokenV2CallerSession) POOL() (common.Address, error) {
	return _ATokenV2.Contract.POOL(&_ATokenV2.CallOpts)
}

// RESERVETREASURYADDRESS is a free data retrieval call binding the contract method 0xae167335.
//
// Solidity: function RESERVE_TREASURY_ADDRESS() view returns(address)
func (_ATokenV2 *ATokenV2Caller) RESERVETREASURYADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ATokenV2.contract.Call(opts, &out, "RESERVE_TREASURY_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RESERVETREASURYADDRESS is a free data retrieval call binding the contract method 0xae167335.
//
// Solidity: function RESERVE_TREASURY_ADDRESS() view returns(address)
func (_ATokenV2 *ATokenV2Session) RESERVETREASURYADDRESS() (common.Address, error) {
	return _ATokenV2.Contract.RESERVETREASURYADDRESS(&_ATokenV2.CallOpts)
}

// RESERVETREASURYADDRESS is a free data retrieval call binding the contract method 0xae167335.
//
// Solidity: function RESERVE_TREASURY_ADDRESS() view returns(address)
func (_ATokenV2 *ATokenV2CallerSession) RESERVETREASURYADDRESS() (common.Address, error) {
	return _ATokenV2.Contract.RESERVETREASURYADDRESS(&_ATokenV2.CallOpts)
}

// UINTMAXVALUE is a free data retrieval call binding the contract method 0xd0fc81d2.
//
// Solidity: function UINT_MAX_VALUE() view returns(uint256)
func (_ATokenV2 *ATokenV2Caller) UINTMAXVALUE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ATokenV2.contract.Call(opts, &out, "UINT_MAX_VALUE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UINTMAXVALUE is a free data retrieval call binding the contract method 0xd0fc81d2.
//
// Solidity: function UINT_MAX_VALUE() view returns(uint256)
func (_ATokenV2 *ATokenV2Session) UINTMAXVALUE() (*big.Int, error) {
	return _ATokenV2.Contract.UINTMAXVALUE(&_ATokenV2.CallOpts)
}

// UINTMAXVALUE is a free data retrieval call binding the contract method 0xd0fc81d2.
//
// Solidity: function UINT_MAX_VALUE() view returns(uint256)
func (_ATokenV2 *ATokenV2CallerSession) UINTMAXVALUE() (*big.Int, error) {
	return _ATokenV2.Contract.UINTMAXVALUE(&_ATokenV2.CallOpts)
}

// UNDERLYINGASSETADDRESS is a free data retrieval call binding the contract method 0xb16a19de.
//
// Solidity: function UNDERLYING_ASSET_ADDRESS() view returns(address)
func (_ATokenV2 *ATokenV2Caller) UNDERLYINGASSETADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ATokenV2.contract.Call(opts, &out, "UNDERLYING_ASSET_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UNDERLYINGASSETADDRESS is a free data retrieval call binding the contract method 0xb16a19de.
//
// Solidity: function UNDERLYING_ASSET_ADDRESS() view returns(address)
func (_ATokenV2 *ATokenV2Session) UNDERLYINGASSETADDRESS() (common.Address, error) {
	return _ATokenV2.Contract.UNDERLYINGASSETADDRESS(&_ATokenV2.CallOpts)
}

// UNDERLYINGASSETADDRESS is a free data retrieval call binding the contract method 0xb16a19de.
//
// Solidity: function UNDERLYING_ASSET_ADDRESS() view returns(address)
func (_ATokenV2 *ATokenV2CallerSession) UNDERLYINGASSETADDRESS() (common.Address, error) {
	return _ATokenV2.Contract.UNDERLYINGASSETADDRESS(&_ATokenV2.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0xb9844d8d.
//
// Solidity: function _nonces(address ) view returns(uint256)
func (_ATokenV2 *ATokenV2Caller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ATokenV2.contract.Call(opts, &out, "_nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0xb9844d8d.
//
// Solidity: function _nonces(address ) view returns(uint256)
func (_ATokenV2 *ATokenV2Session) Nonces(arg0 common.Address) (*big.Int, error) {
	return _ATokenV2.Contract.Nonces(&_ATokenV2.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0xb9844d8d.
//
// Solidity: function _nonces(address ) view returns(uint256)
func (_ATokenV2 *ATokenV2CallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _ATokenV2.Contract.Nonces(&_ATokenV2.CallOpts, arg0)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ATokenV2 *ATokenV2Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ATokenV2.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ATokenV2 *ATokenV2Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ATokenV2.Contract.Allowance(&_ATokenV2.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ATokenV2 *ATokenV2CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ATokenV2.Contract.Allowance(&_ATokenV2.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address user) view returns(uint256)
func (_ATokenV2 *ATokenV2Caller) BalanceOf(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ATokenV2.contract.Call(opts, &out, "balanceOf", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address user) view returns(uint256)
func (_ATokenV2 *ATokenV2Session) BalanceOf(user common.Address) (*big.Int, error) {
	return _ATokenV2.Contract.BalanceOf(&_ATokenV2.CallOpts, user)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address user) view returns(uint256)
func (_ATokenV2 *ATokenV2CallerSession) BalanceOf(user common.Address) (*big.Int, error) {
	return _ATokenV2.Contract.BalanceOf(&_ATokenV2.CallOpts, user)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ATokenV2 *ATokenV2Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ATokenV2.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ATokenV2 *ATokenV2Session) Decimals() (uint8, error) {
	return _ATokenV2.Contract.Decimals(&_ATokenV2.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ATokenV2 *ATokenV2CallerSession) Decimals() (uint8, error) {
	return _ATokenV2.Contract.Decimals(&_ATokenV2.CallOpts)
}

// GetIncentivesController is a free data retrieval call binding the contract method 0x75d26413.
//
// Solidity: function getIncentivesController() view returns(address)
func (_ATokenV2 *ATokenV2Caller) GetIncentivesController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ATokenV2.contract.Call(opts, &out, "getIncentivesController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetIncentivesController is a free data retrieval call binding the contract method 0x75d26413.
//
// Solidity: function getIncentivesController() view returns(address)
func (_ATokenV2 *ATokenV2Session) GetIncentivesController() (common.Address, error) {
	return _ATokenV2.Contract.GetIncentivesController(&_ATokenV2.CallOpts)
}

// GetIncentivesController is a free data retrieval call binding the contract method 0x75d26413.
//
// Solidity: function getIncentivesController() view returns(address)
func (_ATokenV2 *ATokenV2CallerSession) GetIncentivesController() (common.Address, error) {
	return _ATokenV2.Contract.GetIncentivesController(&_ATokenV2.CallOpts)
}

// GetScaledUserBalanceAndSupply is a free data retrieval call binding the contract method 0x0afbcdc9.
//
// Solidity: function getScaledUserBalanceAndSupply(address user) view returns(uint256, uint256)
func (_ATokenV2 *ATokenV2Caller) GetScaledUserBalanceAndSupply(opts *bind.CallOpts, user common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ATokenV2.contract.Call(opts, &out, "getScaledUserBalanceAndSupply", user)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetScaledUserBalanceAndSupply is a free data retrieval call binding the contract method 0x0afbcdc9.
//
// Solidity: function getScaledUserBalanceAndSupply(address user) view returns(uint256, uint256)
func (_ATokenV2 *ATokenV2Session) GetScaledUserBalanceAndSupply(user common.Address) (*big.Int, *big.Int, error) {
	return _ATokenV2.Contract.GetScaledUserBalanceAndSupply(&_ATokenV2.CallOpts, user)
}

// GetScaledUserBalanceAndSupply is a free data retrieval call binding the contract method 0x0afbcdc9.
//
// Solidity: function getScaledUserBalanceAndSupply(address user) view returns(uint256, uint256)
func (_ATokenV2 *ATokenV2CallerSession) GetScaledUserBalanceAndSupply(user common.Address) (*big.Int, *big.Int, error) {
	return _ATokenV2.Contract.GetScaledUserBalanceAndSupply(&_ATokenV2.CallOpts, user)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ATokenV2 *ATokenV2Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ATokenV2.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ATokenV2 *ATokenV2Session) Name() (string, error) {
	return _ATokenV2.Contract.Name(&_ATokenV2.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ATokenV2 *ATokenV2CallerSession) Name() (string, error) {
	return _ATokenV2.Contract.Name(&_ATokenV2.CallOpts)
}

// ScaledBalanceOf is a free data retrieval call binding the contract method 0x1da24f3e.
//
// Solidity: function scaledBalanceOf(address user) view returns(uint256)
func (_ATokenV2 *ATokenV2Caller) ScaledBalanceOf(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ATokenV2.contract.Call(opts, &out, "scaledBalanceOf", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ScaledBalanceOf is a free data retrieval call binding the contract method 0x1da24f3e.
//
// Solidity: function scaledBalanceOf(address user) view returns(uint256)
func (_ATokenV2 *ATokenV2Session) ScaledBalanceOf(user common.Address) (*big.Int, error) {
	return _ATokenV2.Contract.ScaledBalanceOf(&_ATokenV2.CallOpts, user)
}

// ScaledBalanceOf is a free data retrieval call binding the contract method 0x1da24f3e.
//
// Solidity: function scaledBalanceOf(address user) view returns(uint256)
func (_ATokenV2 *ATokenV2CallerSession) ScaledBalanceOf(user common.Address) (*big.Int, error) {
	return _ATokenV2.Contract.ScaledBalanceOf(&_ATokenV2.CallOpts, user)
}

// ScaledTotalSupply is a free data retrieval call binding the contract method 0xb1bf962d.
//
// Solidity: function scaledTotalSupply() view returns(uint256)
func (_ATokenV2 *ATokenV2Caller) ScaledTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ATokenV2.contract.Call(opts, &out, "scaledTotalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ScaledTotalSupply is a free data retrieval call binding the contract method 0xb1bf962d.
//
// Solidity: function scaledTotalSupply() view returns(uint256)
func (_ATokenV2 *ATokenV2Session) ScaledTotalSupply() (*big.Int, error) {
	return _ATokenV2.Contract.ScaledTotalSupply(&_ATokenV2.CallOpts)
}

// ScaledTotalSupply is a free data retrieval call binding the contract method 0xb1bf962d.
//
// Solidity: function scaledTotalSupply() view returns(uint256)
func (_ATokenV2 *ATokenV2CallerSession) ScaledTotalSupply() (*big.Int, error) {
	return _ATokenV2.Contract.ScaledTotalSupply(&_ATokenV2.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ATokenV2 *ATokenV2Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ATokenV2.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ATokenV2 *ATokenV2Session) Symbol() (string, error) {
	return _ATokenV2.Contract.Symbol(&_ATokenV2.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ATokenV2 *ATokenV2CallerSession) Symbol() (string, error) {
	return _ATokenV2.Contract.Symbol(&_ATokenV2.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ATokenV2 *ATokenV2Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ATokenV2.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ATokenV2 *ATokenV2Session) TotalSupply() (*big.Int, error) {
	return _ATokenV2.Contract.TotalSupply(&_ATokenV2.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ATokenV2 *ATokenV2CallerSession) TotalSupply() (*big.Int, error) {
	return _ATokenV2.Contract.TotalSupply(&_ATokenV2.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ATokenV2 *ATokenV2Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ATokenV2.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ATokenV2 *ATokenV2Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ATokenV2.Contract.Approve(&_ATokenV2.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ATokenV2 *ATokenV2TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ATokenV2.Contract.Approve(&_ATokenV2.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xd7020d0a.
//
// Solidity: function burn(address user, address receiverOfUnderlying, uint256 amount, uint256 index) returns()
func (_ATokenV2 *ATokenV2Transactor) Burn(opts *bind.TransactOpts, user common.Address, receiverOfUnderlying common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _ATokenV2.contract.Transact(opts, "burn", user, receiverOfUnderlying, amount, index)
}

// Burn is a paid mutator transaction binding the contract method 0xd7020d0a.
//
// Solidity: function burn(address user, address receiverOfUnderlying, uint256 amount, uint256 index) returns()
func (_ATokenV2 *ATokenV2Session) Burn(user common.Address, receiverOfUnderlying common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _ATokenV2.Contract.Burn(&_ATokenV2.TransactOpts, user, receiverOfUnderlying, amount, index)
}

// Burn is a paid mutator transaction binding the contract method 0xd7020d0a.
//
// Solidity: function burn(address user, address receiverOfUnderlying, uint256 amount, uint256 index) returns()
func (_ATokenV2 *ATokenV2TransactorSession) Burn(user common.Address, receiverOfUnderlying common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _ATokenV2.Contract.Burn(&_ATokenV2.TransactOpts, user, receiverOfUnderlying, amount, index)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ATokenV2 *ATokenV2Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ATokenV2.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ATokenV2 *ATokenV2Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ATokenV2.Contract.DecreaseAllowance(&_ATokenV2.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ATokenV2 *ATokenV2TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ATokenV2.Contract.DecreaseAllowance(&_ATokenV2.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ATokenV2 *ATokenV2Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ATokenV2.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ATokenV2 *ATokenV2Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ATokenV2.Contract.IncreaseAllowance(&_ATokenV2.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ATokenV2 *ATokenV2TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ATokenV2.Contract.IncreaseAllowance(&_ATokenV2.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x3118724e.
//
// Solidity: function initialize(uint8 underlyingAssetDecimals, string tokenName, string tokenSymbol) returns()
func (_ATokenV2 *ATokenV2Transactor) Initialize(opts *bind.TransactOpts, underlyingAssetDecimals uint8, tokenName string, tokenSymbol string) (*types.Transaction, error) {
	return _ATokenV2.contract.Transact(opts, "initialize", underlyingAssetDecimals, tokenName, tokenSymbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x3118724e.
//
// Solidity: function initialize(uint8 underlyingAssetDecimals, string tokenName, string tokenSymbol) returns()
func (_ATokenV2 *ATokenV2Session) Initialize(underlyingAssetDecimals uint8, tokenName string, tokenSymbol string) (*types.Transaction, error) {
	return _ATokenV2.Contract.Initialize(&_ATokenV2.TransactOpts, underlyingAssetDecimals, tokenName, tokenSymbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x3118724e.
//
// Solidity: function initialize(uint8 underlyingAssetDecimals, string tokenName, string tokenSymbol) returns()
func (_ATokenV2 *ATokenV2TransactorSession) Initialize(underlyingAssetDecimals uint8, tokenName string, tokenSymbol string) (*types.Transaction, error) {
	return _ATokenV2.Contract.Initialize(&_ATokenV2.TransactOpts, underlyingAssetDecimals, tokenName, tokenSymbol)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address user, uint256 amount, uint256 index) returns(bool)
func (_ATokenV2 *ATokenV2Transactor) Mint(opts *bind.TransactOpts, user common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _ATokenV2.contract.Transact(opts, "mint", user, amount, index)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address user, uint256 amount, uint256 index) returns(bool)
func (_ATokenV2 *ATokenV2Session) Mint(user common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _ATokenV2.Contract.Mint(&_ATokenV2.TransactOpts, user, amount, index)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address user, uint256 amount, uint256 index) returns(bool)
func (_ATokenV2 *ATokenV2TransactorSession) Mint(user common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _ATokenV2.Contract.Mint(&_ATokenV2.TransactOpts, user, amount, index)
}

// MintToTreasury is a paid mutator transaction binding the contract method 0x7df5bd3b.
//
// Solidity: function mintToTreasury(uint256 amount, uint256 index) returns()
func (_ATokenV2 *ATokenV2Transactor) MintToTreasury(opts *bind.TransactOpts, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _ATokenV2.contract.Transact(opts, "mintToTreasury", amount, index)
}

// MintToTreasury is a paid mutator transaction binding the contract method 0x7df5bd3b.
//
// Solidity: function mintToTreasury(uint256 amount, uint256 index) returns()
func (_ATokenV2 *ATokenV2Session) MintToTreasury(amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _ATokenV2.Contract.MintToTreasury(&_ATokenV2.TransactOpts, amount, index)
}

// MintToTreasury is a paid mutator transaction binding the contract method 0x7df5bd3b.
//
// Solidity: function mintToTreasury(uint256 amount, uint256 index) returns()
func (_ATokenV2 *ATokenV2TransactorSession) MintToTreasury(amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _ATokenV2.Contract.MintToTreasury(&_ATokenV2.TransactOpts, amount, index)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_ATokenV2 *ATokenV2Transactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ATokenV2.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_ATokenV2 *ATokenV2Session) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ATokenV2.Contract.Permit(&_ATokenV2.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_ATokenV2 *ATokenV2TransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ATokenV2.Contract.Permit(&_ATokenV2.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ATokenV2 *ATokenV2Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ATokenV2.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ATokenV2 *ATokenV2Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ATokenV2.Contract.Transfer(&_ATokenV2.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ATokenV2 *ATokenV2TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ATokenV2.Contract.Transfer(&_ATokenV2.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ATokenV2 *ATokenV2Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ATokenV2.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ATokenV2 *ATokenV2Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ATokenV2.Contract.TransferFrom(&_ATokenV2.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ATokenV2 *ATokenV2TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ATokenV2.Contract.TransferFrom(&_ATokenV2.TransactOpts, sender, recipient, amount)
}

// TransferOnLiquidation is a paid mutator transaction binding the contract method 0xf866c319.
//
// Solidity: function transferOnLiquidation(address from, address to, uint256 value) returns()
func (_ATokenV2 *ATokenV2Transactor) TransferOnLiquidation(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ATokenV2.contract.Transact(opts, "transferOnLiquidation", from, to, value)
}

// TransferOnLiquidation is a paid mutator transaction binding the contract method 0xf866c319.
//
// Solidity: function transferOnLiquidation(address from, address to, uint256 value) returns()
func (_ATokenV2 *ATokenV2Session) TransferOnLiquidation(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ATokenV2.Contract.TransferOnLiquidation(&_ATokenV2.TransactOpts, from, to, value)
}

// TransferOnLiquidation is a paid mutator transaction binding the contract method 0xf866c319.
//
// Solidity: function transferOnLiquidation(address from, address to, uint256 value) returns()
func (_ATokenV2 *ATokenV2TransactorSession) TransferOnLiquidation(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ATokenV2.Contract.TransferOnLiquidation(&_ATokenV2.TransactOpts, from, to, value)
}

// TransferUnderlyingTo is a paid mutator transaction binding the contract method 0x4efecaa5.
//
// Solidity: function transferUnderlyingTo(address target, uint256 amount) returns(uint256)
func (_ATokenV2 *ATokenV2Transactor) TransferUnderlyingTo(opts *bind.TransactOpts, target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ATokenV2.contract.Transact(opts, "transferUnderlyingTo", target, amount)
}

// TransferUnderlyingTo is a paid mutator transaction binding the contract method 0x4efecaa5.
//
// Solidity: function transferUnderlyingTo(address target, uint256 amount) returns(uint256)
func (_ATokenV2 *ATokenV2Session) TransferUnderlyingTo(target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ATokenV2.Contract.TransferUnderlyingTo(&_ATokenV2.TransactOpts, target, amount)
}

// TransferUnderlyingTo is a paid mutator transaction binding the contract method 0x4efecaa5.
//
// Solidity: function transferUnderlyingTo(address target, uint256 amount) returns(uint256)
func (_ATokenV2 *ATokenV2TransactorSession) TransferUnderlyingTo(target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ATokenV2.Contract.TransferUnderlyingTo(&_ATokenV2.TransactOpts, target, amount)
}

// ATokenV2ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ATokenV2 contract.
type ATokenV2ApprovalIterator struct {
	Event *ATokenV2Approval // Event containing the contract specifics and raw log

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
func (it *ATokenV2ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ATokenV2Approval)
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
		it.Event = new(ATokenV2Approval)
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
func (it *ATokenV2ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ATokenV2ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ATokenV2Approval represents a Approval event raised by the ATokenV2 contract.
type ATokenV2Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ATokenV2 *ATokenV2Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ATokenV2ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ATokenV2.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ATokenV2ApprovalIterator{contract: _ATokenV2.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ATokenV2 *ATokenV2Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ATokenV2Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ATokenV2.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ATokenV2Approval)
				if err := _ATokenV2.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ATokenV2 *ATokenV2Filterer) ParseApproval(log types.Log) (*ATokenV2Approval, error) {
	event := new(ATokenV2Approval)
	if err := _ATokenV2.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ATokenV2BalanceTransferIterator is returned from FilterBalanceTransfer and is used to iterate over the raw logs and unpacked data for BalanceTransfer events raised by the ATokenV2 contract.
type ATokenV2BalanceTransferIterator struct {
	Event *ATokenV2BalanceTransfer // Event containing the contract specifics and raw log

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
func (it *ATokenV2BalanceTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ATokenV2BalanceTransfer)
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
		it.Event = new(ATokenV2BalanceTransfer)
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
func (it *ATokenV2BalanceTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ATokenV2BalanceTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ATokenV2BalanceTransfer represents a BalanceTransfer event raised by the ATokenV2 contract.
type ATokenV2BalanceTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBalanceTransfer is a free log retrieval operation binding the contract event 0x4beccb90f994c31aced7a23b5611020728a23d8ec5cddd1a3e9d97b96fda8666.
//
// Solidity: event BalanceTransfer(address indexed from, address indexed to, uint256 value, uint256 index)
func (_ATokenV2 *ATokenV2Filterer) FilterBalanceTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ATokenV2BalanceTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ATokenV2.contract.FilterLogs(opts, "BalanceTransfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ATokenV2BalanceTransferIterator{contract: _ATokenV2.contract, event: "BalanceTransfer", logs: logs, sub: sub}, nil
}

// WatchBalanceTransfer is a free log subscription operation binding the contract event 0x4beccb90f994c31aced7a23b5611020728a23d8ec5cddd1a3e9d97b96fda8666.
//
// Solidity: event BalanceTransfer(address indexed from, address indexed to, uint256 value, uint256 index)
func (_ATokenV2 *ATokenV2Filterer) WatchBalanceTransfer(opts *bind.WatchOpts, sink chan<- *ATokenV2BalanceTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ATokenV2.contract.WatchLogs(opts, "BalanceTransfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ATokenV2BalanceTransfer)
				if err := _ATokenV2.contract.UnpackLog(event, "BalanceTransfer", log); err != nil {
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

// ParseBalanceTransfer is a log parse operation binding the contract event 0x4beccb90f994c31aced7a23b5611020728a23d8ec5cddd1a3e9d97b96fda8666.
//
// Solidity: event BalanceTransfer(address indexed from, address indexed to, uint256 value, uint256 index)
func (_ATokenV2 *ATokenV2Filterer) ParseBalanceTransfer(log types.Log) (*ATokenV2BalanceTransfer, error) {
	event := new(ATokenV2BalanceTransfer)
	if err := _ATokenV2.contract.UnpackLog(event, "BalanceTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ATokenV2BurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the ATokenV2 contract.
type ATokenV2BurnIterator struct {
	Event *ATokenV2Burn // Event containing the contract specifics and raw log

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
func (it *ATokenV2BurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ATokenV2Burn)
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
		it.Event = new(ATokenV2Burn)
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
func (it *ATokenV2BurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ATokenV2BurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ATokenV2Burn represents a Burn event raised by the ATokenV2 contract.
type ATokenV2Burn struct {
	From   common.Address
	Target common.Address
	Value  *big.Int
	Index  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0x5d624aa9c148153ab3446c1b154f660ee7701e549fe9b62dab7171b1c80e6fa2.
//
// Solidity: event Burn(address indexed from, address indexed target, uint256 value, uint256 index)
func (_ATokenV2 *ATokenV2Filterer) FilterBurn(opts *bind.FilterOpts, from []common.Address, target []common.Address) (*ATokenV2BurnIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _ATokenV2.contract.FilterLogs(opts, "Burn", fromRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &ATokenV2BurnIterator{contract: _ATokenV2.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0x5d624aa9c148153ab3446c1b154f660ee7701e549fe9b62dab7171b1c80e6fa2.
//
// Solidity: event Burn(address indexed from, address indexed target, uint256 value, uint256 index)
func (_ATokenV2 *ATokenV2Filterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *ATokenV2Burn, from []common.Address, target []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _ATokenV2.contract.WatchLogs(opts, "Burn", fromRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ATokenV2Burn)
				if err := _ATokenV2.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0x5d624aa9c148153ab3446c1b154f660ee7701e549fe9b62dab7171b1c80e6fa2.
//
// Solidity: event Burn(address indexed from, address indexed target, uint256 value, uint256 index)
func (_ATokenV2 *ATokenV2Filterer) ParseBurn(log types.Log) (*ATokenV2Burn, error) {
	event := new(ATokenV2Burn)
	if err := _ATokenV2.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ATokenV2InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ATokenV2 contract.
type ATokenV2InitializedIterator struct {
	Event *ATokenV2Initialized // Event containing the contract specifics and raw log

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
func (it *ATokenV2InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ATokenV2Initialized)
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
		it.Event = new(ATokenV2Initialized)
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
func (it *ATokenV2InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ATokenV2InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ATokenV2Initialized represents a Initialized event raised by the ATokenV2 contract.
type ATokenV2Initialized struct {
	UnderlyingAsset      common.Address
	Pool                 common.Address
	Treasury             common.Address
	IncentivesController common.Address
	ATokenDecimals       uint8
	ATokenName           string
	ATokenSymbol         string
	Params               []byte
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xb19e051f8af41150ccccb3fc2c2d8d15f4a4cf434f32a559ba75fe73d6eea20b.
//
// Solidity: event Initialized(address indexed underlyingAsset, address indexed pool, address treasury, address incentivesController, uint8 aTokenDecimals, string aTokenName, string aTokenSymbol, bytes params)
func (_ATokenV2 *ATokenV2Filterer) FilterInitialized(opts *bind.FilterOpts, underlyingAsset []common.Address, pool []common.Address) (*ATokenV2InitializedIterator, error) {

	var underlyingAssetRule []interface{}
	for _, underlyingAssetItem := range underlyingAsset {
		underlyingAssetRule = append(underlyingAssetRule, underlyingAssetItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _ATokenV2.contract.FilterLogs(opts, "Initialized", underlyingAssetRule, poolRule)
	if err != nil {
		return nil, err
	}
	return &ATokenV2InitializedIterator{contract: _ATokenV2.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xb19e051f8af41150ccccb3fc2c2d8d15f4a4cf434f32a559ba75fe73d6eea20b.
//
// Solidity: event Initialized(address indexed underlyingAsset, address indexed pool, address treasury, address incentivesController, uint8 aTokenDecimals, string aTokenName, string aTokenSymbol, bytes params)
func (_ATokenV2 *ATokenV2Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ATokenV2Initialized, underlyingAsset []common.Address, pool []common.Address) (event.Subscription, error) {

	var underlyingAssetRule []interface{}
	for _, underlyingAssetItem := range underlyingAsset {
		underlyingAssetRule = append(underlyingAssetRule, underlyingAssetItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _ATokenV2.contract.WatchLogs(opts, "Initialized", underlyingAssetRule, poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ATokenV2Initialized)
				if err := _ATokenV2.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xb19e051f8af41150ccccb3fc2c2d8d15f4a4cf434f32a559ba75fe73d6eea20b.
//
// Solidity: event Initialized(address indexed underlyingAsset, address indexed pool, address treasury, address incentivesController, uint8 aTokenDecimals, string aTokenName, string aTokenSymbol, bytes params)
func (_ATokenV2 *ATokenV2Filterer) ParseInitialized(log types.Log) (*ATokenV2Initialized, error) {
	event := new(ATokenV2Initialized)
	if err := _ATokenV2.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ATokenV2MintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the ATokenV2 contract.
type ATokenV2MintIterator struct {
	Event *ATokenV2Mint // Event containing the contract specifics and raw log

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
func (it *ATokenV2MintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ATokenV2Mint)
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
		it.Event = new(ATokenV2Mint)
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
func (it *ATokenV2MintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ATokenV2MintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ATokenV2Mint represents a Mint event raised by the ATokenV2 contract.
type ATokenV2Mint struct {
	From  common.Address
	Value *big.Int
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed from, uint256 value, uint256 index)
func (_ATokenV2 *ATokenV2Filterer) FilterMint(opts *bind.FilterOpts, from []common.Address) (*ATokenV2MintIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _ATokenV2.contract.FilterLogs(opts, "Mint", fromRule)
	if err != nil {
		return nil, err
	}
	return &ATokenV2MintIterator{contract: _ATokenV2.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed from, uint256 value, uint256 index)
func (_ATokenV2 *ATokenV2Filterer) WatchMint(opts *bind.WatchOpts, sink chan<- *ATokenV2Mint, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _ATokenV2.contract.WatchLogs(opts, "Mint", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ATokenV2Mint)
				if err := _ATokenV2.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed from, uint256 value, uint256 index)
func (_ATokenV2 *ATokenV2Filterer) ParseMint(log types.Log) (*ATokenV2Mint, error) {
	event := new(ATokenV2Mint)
	if err := _ATokenV2.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ATokenV2TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ATokenV2 contract.
type ATokenV2TransferIterator struct {
	Event *ATokenV2Transfer // Event containing the contract specifics and raw log

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
func (it *ATokenV2TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ATokenV2Transfer)
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
		it.Event = new(ATokenV2Transfer)
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
func (it *ATokenV2TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ATokenV2TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ATokenV2Transfer represents a Transfer event raised by the ATokenV2 contract.
type ATokenV2Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ATokenV2 *ATokenV2Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ATokenV2TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ATokenV2.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ATokenV2TransferIterator{contract: _ATokenV2.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ATokenV2 *ATokenV2Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ATokenV2Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ATokenV2.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ATokenV2Transfer)
				if err := _ATokenV2.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ATokenV2 *ATokenV2Filterer) ParseTransfer(log types.Log) (*ATokenV2Transfer, error) {
	event := new(ATokenV2Transfer)
	if err := _ATokenV2.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
