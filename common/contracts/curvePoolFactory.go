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

// CurvePoolFactoryMetaData contains all meta data concerning the CurvePoolFactory contract.
var CurvePoolFactoryMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"CryptoPoolDeployed\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false},{\"name\":\"coins\",\"type\":\"address[2]\",\"indexed\":false},{\"name\":\"A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"gamma\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"mid_fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"out_fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"allowed_extra_profit\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"fee_gamma\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"adjustment_step\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"admin_fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"ma_half_time\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"initial_price\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"deployer\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"LiquidityGaugeDeployed\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"indexed\":false},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false},{\"name\":\"gauge\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateFeeReceiver\",\"inputs\":[{\"name\":\"_old_fee_receiver\",\"type\":\"address\",\"indexed\":false},{\"name\":\"_new_fee_receiver\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdatePoolImplementation\",\"inputs\":[{\"name\":\"_old_pool_implementation\",\"type\":\"address\",\"indexed\":false},{\"name\":\"_new_pool_implementation\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateTokenImplementation\",\"inputs\":[{\"name\":\"_old_token_implementation\",\"type\":\"address\",\"indexed\":false},{\"name\":\"_new_token_implementation\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateGaugeImplementation\",\"inputs\":[{\"name\":\"_old_gauge_implementation\",\"type\":\"address\",\"indexed\":false},{\"name\":\"_new_gauge_implementation\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"TransferOwnership\",\"inputs\":[{\"name\":\"_old_owner\",\"type\":\"address\",\"indexed\":false},{\"name\":\"_new_owner\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"_fee_receiver\",\"type\":\"address\"},{\"name\":\"_pool_implementation\",\"type\":\"address\"},{\"name\":\"_token_implementation\",\"type\":\"address\"},{\"name\":\"_gauge_implementation\",\"type\":\"address\"},{\"name\":\"_weth\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deploy_pool\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_coins\",\"type\":\"address[2]\"},{\"name\":\"A\",\"type\":\"uint256\"},{\"name\":\"gamma\",\"type\":\"uint256\"},{\"name\":\"mid_fee\",\"type\":\"uint256\"},{\"name\":\"out_fee\",\"type\":\"uint256\"},{\"name\":\"allowed_extra_profit\",\"type\":\"uint256\"},{\"name\":\"fee_gamma\",\"type\":\"uint256\"},{\"name\":\"adjustment_step\",\"type\":\"uint256\"},{\"name\":\"admin_fee\",\"type\":\"uint256\"},{\"name\":\"ma_half_time\",\"type\":\"uint256\"},{\"name\":\"initial_price\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deploy_gauge\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_fee_receiver\",\"inputs\":[{\"name\":\"_fee_receiver\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_pool_implementation\",\"inputs\":[{\"name\":\"_pool_implementation\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_token_implementation\",\"inputs\":[{\"name\":\"_token_implementation\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_gauge_implementation\",\"inputs\":[{\"name\":\"_gauge_implementation\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"commit_transfer_ownership\",\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"accept_transfer_ownership\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"find_pool_for_coins\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"find_pool_for_coins\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_coins\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[2]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_decimals\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_balances\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_coin_indices\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_gauge\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_eth_index\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_token\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"fee_receiver\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"pool_implementation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"token_implementation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"gauge_implementation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"pool_count\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"pool_list\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]}]",
}

// CurvePoolFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use CurvePoolFactoryMetaData.ABI instead.
var CurvePoolFactoryABI = CurvePoolFactoryMetaData.ABI

// CurvePoolFactory is an auto generated Go binding around an Ethereum contract.
type CurvePoolFactory struct {
	CurvePoolFactoryCaller     // Read-only binding to the contract
	CurvePoolFactoryTransactor // Write-only binding to the contract
	CurvePoolFactoryFilterer   // Log filterer for contract events
}

// CurvePoolFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurvePoolFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurvePoolFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurvePoolFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurvePoolFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurvePoolFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurvePoolFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurvePoolFactorySession struct {
	Contract     *CurvePoolFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CurvePoolFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurvePoolFactoryCallerSession struct {
	Contract *CurvePoolFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// CurvePoolFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurvePoolFactoryTransactorSession struct {
	Contract     *CurvePoolFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// CurvePoolFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurvePoolFactoryRaw struct {
	Contract *CurvePoolFactory // Generic contract binding to access the raw methods on
}

// CurvePoolFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurvePoolFactoryCallerRaw struct {
	Contract *CurvePoolFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// CurvePoolFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurvePoolFactoryTransactorRaw struct {
	Contract *CurvePoolFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurvePoolFactory creates a new instance of CurvePoolFactory, bound to a specific deployed contract.
func NewCurvePoolFactory(address common.Address, backend bind.ContractBackend) (*CurvePoolFactory, error) {
	contract, err := bindCurvePoolFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CurvePoolFactory{CurvePoolFactoryCaller: CurvePoolFactoryCaller{contract: contract}, CurvePoolFactoryTransactor: CurvePoolFactoryTransactor{contract: contract}, CurvePoolFactoryFilterer: CurvePoolFactoryFilterer{contract: contract}}, nil
}

// NewCurvePoolFactoryCaller creates a new read-only instance of CurvePoolFactory, bound to a specific deployed contract.
func NewCurvePoolFactoryCaller(address common.Address, caller bind.ContractCaller) (*CurvePoolFactoryCaller, error) {
	contract, err := bindCurvePoolFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurvePoolFactoryCaller{contract: contract}, nil
}

// NewCurvePoolFactoryTransactor creates a new write-only instance of CurvePoolFactory, bound to a specific deployed contract.
func NewCurvePoolFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*CurvePoolFactoryTransactor, error) {
	contract, err := bindCurvePoolFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurvePoolFactoryTransactor{contract: contract}, nil
}

// NewCurvePoolFactoryFilterer creates a new log filterer instance of CurvePoolFactory, bound to a specific deployed contract.
func NewCurvePoolFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*CurvePoolFactoryFilterer, error) {
	contract, err := bindCurvePoolFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurvePoolFactoryFilterer{contract: contract}, nil
}

// bindCurvePoolFactory binds a generic wrapper to an already deployed contract.
func bindCurvePoolFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CurvePoolFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurvePoolFactory *CurvePoolFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurvePoolFactory.Contract.CurvePoolFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurvePoolFactory *CurvePoolFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurvePoolFactory.Contract.CurvePoolFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurvePoolFactory *CurvePoolFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurvePoolFactory.Contract.CurvePoolFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurvePoolFactory *CurvePoolFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurvePoolFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurvePoolFactory *CurvePoolFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurvePoolFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurvePoolFactory *CurvePoolFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurvePoolFactory.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurvePoolFactory *CurvePoolFactoryCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurvePoolFactory.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurvePoolFactory *CurvePoolFactorySession) Admin() (common.Address, error) {
	return _CurvePoolFactory.Contract.Admin(&_CurvePoolFactory.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurvePoolFactory *CurvePoolFactoryCallerSession) Admin() (common.Address, error) {
	return _CurvePoolFactory.Contract.Admin(&_CurvePoolFactory.CallOpts)
}

// FeeReceiver is a free data retrieval call binding the contract method 0xcab4d3db.
//
// Solidity: function fee_receiver() view returns(address)
func (_CurvePoolFactory *CurvePoolFactoryCaller) FeeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurvePoolFactory.contract.Call(opts, &out, "fee_receiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeReceiver is a free data retrieval call binding the contract method 0xcab4d3db.
//
// Solidity: function fee_receiver() view returns(address)
func (_CurvePoolFactory *CurvePoolFactorySession) FeeReceiver() (common.Address, error) {
	return _CurvePoolFactory.Contract.FeeReceiver(&_CurvePoolFactory.CallOpts)
}

// FeeReceiver is a free data retrieval call binding the contract method 0xcab4d3db.
//
// Solidity: function fee_receiver() view returns(address)
func (_CurvePoolFactory *CurvePoolFactoryCallerSession) FeeReceiver() (common.Address, error) {
	return _CurvePoolFactory.Contract.FeeReceiver(&_CurvePoolFactory.CallOpts)
}

// FindPoolForCoins is a free data retrieval call binding the contract method 0xa87df06c.
//
// Solidity: function find_pool_for_coins(address _from, address _to) view returns(address)
func (_CurvePoolFactory *CurvePoolFactoryCaller) FindPoolForCoins(opts *bind.CallOpts, _from common.Address, _to common.Address) (common.Address, error) {
	var out []interface{}
	err := _CurvePoolFactory.contract.Call(opts, &out, "find_pool_for_coins", _from, _to)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FindPoolForCoins is a free data retrieval call binding the contract method 0xa87df06c.
//
// Solidity: function find_pool_for_coins(address _from, address _to) view returns(address)
func (_CurvePoolFactory *CurvePoolFactorySession) FindPoolForCoins(_from common.Address, _to common.Address) (common.Address, error) {
	return _CurvePoolFactory.Contract.FindPoolForCoins(&_CurvePoolFactory.CallOpts, _from, _to)
}

// FindPoolForCoins is a free data retrieval call binding the contract method 0xa87df06c.
//
// Solidity: function find_pool_for_coins(address _from, address _to) view returns(address)
func (_CurvePoolFactory *CurvePoolFactoryCallerSession) FindPoolForCoins(_from common.Address, _to common.Address) (common.Address, error) {
	return _CurvePoolFactory.Contract.FindPoolForCoins(&_CurvePoolFactory.CallOpts, _from, _to)
}

// FindPoolForCoins0 is a free data retrieval call binding the contract method 0x6982eb0b.
//
// Solidity: function find_pool_for_coins(address _from, address _to, uint256 i) view returns(address)
func (_CurvePoolFactory *CurvePoolFactoryCaller) FindPoolForCoins0(opts *bind.CallOpts, _from common.Address, _to common.Address, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurvePoolFactory.contract.Call(opts, &out, "find_pool_for_coins0", _from, _to, i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FindPoolForCoins0 is a free data retrieval call binding the contract method 0x6982eb0b.
//
// Solidity: function find_pool_for_coins(address _from, address _to, uint256 i) view returns(address)
func (_CurvePoolFactory *CurvePoolFactorySession) FindPoolForCoins0(_from common.Address, _to common.Address, i *big.Int) (common.Address, error) {
	return _CurvePoolFactory.Contract.FindPoolForCoins0(&_CurvePoolFactory.CallOpts, _from, _to, i)
}

// FindPoolForCoins0 is a free data retrieval call binding the contract method 0x6982eb0b.
//
// Solidity: function find_pool_for_coins(address _from, address _to, uint256 i) view returns(address)
func (_CurvePoolFactory *CurvePoolFactoryCallerSession) FindPoolForCoins0(_from common.Address, _to common.Address, i *big.Int) (common.Address, error) {
	return _CurvePoolFactory.Contract.FindPoolForCoins0(&_CurvePoolFactory.CallOpts, _from, _to, i)
}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_CurvePoolFactory *CurvePoolFactoryCaller) FutureAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurvePoolFactory.contract.Call(opts, &out, "future_admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_CurvePoolFactory *CurvePoolFactorySession) FutureAdmin() (common.Address, error) {
	return _CurvePoolFactory.Contract.FutureAdmin(&_CurvePoolFactory.CallOpts)
}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_CurvePoolFactory *CurvePoolFactoryCallerSession) FutureAdmin() (common.Address, error) {
	return _CurvePoolFactory.Contract.FutureAdmin(&_CurvePoolFactory.CallOpts)
}

// GaugeImplementation is a free data retrieval call binding the contract method 0x8df24207.
//
// Solidity: function gauge_implementation() view returns(address)
func (_CurvePoolFactory *CurvePoolFactoryCaller) GaugeImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurvePoolFactory.contract.Call(opts, &out, "gauge_implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GaugeImplementation is a free data retrieval call binding the contract method 0x8df24207.
//
// Solidity: function gauge_implementation() view returns(address)
func (_CurvePoolFactory *CurvePoolFactorySession) GaugeImplementation() (common.Address, error) {
	return _CurvePoolFactory.Contract.GaugeImplementation(&_CurvePoolFactory.CallOpts)
}

// GaugeImplementation is a free data retrieval call binding the contract method 0x8df24207.
//
// Solidity: function gauge_implementation() view returns(address)
func (_CurvePoolFactory *CurvePoolFactoryCallerSession) GaugeImplementation() (common.Address, error) {
	return _CurvePoolFactory.Contract.GaugeImplementation(&_CurvePoolFactory.CallOpts)
}

// GetBalances is a free data retrieval call binding the contract method 0x92e3cc2d.
//
// Solidity: function get_balances(address _pool) view returns(uint256[2])
func (_CurvePoolFactory *CurvePoolFactoryCaller) GetBalances(opts *bind.CallOpts, _pool common.Address) ([2]*big.Int, error) {
	var out []interface{}
	err := _CurvePoolFactory.contract.Call(opts, &out, "get_balances", _pool)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetBalances is a free data retrieval call binding the contract method 0x92e3cc2d.
//
// Solidity: function get_balances(address _pool) view returns(uint256[2])
func (_CurvePoolFactory *CurvePoolFactorySession) GetBalances(_pool common.Address) ([2]*big.Int, error) {
	return _CurvePoolFactory.Contract.GetBalances(&_CurvePoolFactory.CallOpts, _pool)
}

// GetBalances is a free data retrieval call binding the contract method 0x92e3cc2d.
//
// Solidity: function get_balances(address _pool) view returns(uint256[2])
func (_CurvePoolFactory *CurvePoolFactoryCallerSession) GetBalances(_pool common.Address) ([2]*big.Int, error) {
	return _CurvePoolFactory.Contract.GetBalances(&_CurvePoolFactory.CallOpts, _pool)
}

// GetCoinIndices is a free data retrieval call binding the contract method 0xeb85226d.
//
// Solidity: function get_coin_indices(address _pool, address _from, address _to) view returns(uint256, uint256)
func (_CurvePoolFactory *CurvePoolFactoryCaller) GetCoinIndices(opts *bind.CallOpts, _pool common.Address, _from common.Address, _to common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _CurvePoolFactory.contract.Call(opts, &out, "get_coin_indices", _pool, _from, _to)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetCoinIndices is a free data retrieval call binding the contract method 0xeb85226d.
//
// Solidity: function get_coin_indices(address _pool, address _from, address _to) view returns(uint256, uint256)
func (_CurvePoolFactory *CurvePoolFactorySession) GetCoinIndices(_pool common.Address, _from common.Address, _to common.Address) (*big.Int, *big.Int, error) {
	return _CurvePoolFactory.Contract.GetCoinIndices(&_CurvePoolFactory.CallOpts, _pool, _from, _to)
}

// GetCoinIndices is a free data retrieval call binding the contract method 0xeb85226d.
//
// Solidity: function get_coin_indices(address _pool, address _from, address _to) view returns(uint256, uint256)
func (_CurvePoolFactory *CurvePoolFactoryCallerSession) GetCoinIndices(_pool common.Address, _from common.Address, _to common.Address) (*big.Int, *big.Int, error) {
	return _CurvePoolFactory.Contract.GetCoinIndices(&_CurvePoolFactory.CallOpts, _pool, _from, _to)
}

// GetCoins is a free data retrieval call binding the contract method 0x9ac90d3d.
//
// Solidity: function get_coins(address _pool) view returns(address[2])
func (_CurvePoolFactory *CurvePoolFactoryCaller) GetCoins(opts *bind.CallOpts, _pool common.Address) ([2]common.Address, error) {
	var out []interface{}
	err := _CurvePoolFactory.contract.Call(opts, &out, "get_coins", _pool)

	if err != nil {
		return *new([2]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([2]common.Address)).(*[2]common.Address)

	return out0, err

}

// GetCoins is a free data retrieval call binding the contract method 0x9ac90d3d.
//
// Solidity: function get_coins(address _pool) view returns(address[2])
func (_CurvePoolFactory *CurvePoolFactorySession) GetCoins(_pool common.Address) ([2]common.Address, error) {
	return _CurvePoolFactory.Contract.GetCoins(&_CurvePoolFactory.CallOpts, _pool)
}

// GetCoins is a free data retrieval call binding the contract method 0x9ac90d3d.
//
// Solidity: function get_coins(address _pool) view returns(address[2])
func (_CurvePoolFactory *CurvePoolFactoryCallerSession) GetCoins(_pool common.Address) ([2]common.Address, error) {
	return _CurvePoolFactory.Contract.GetCoins(&_CurvePoolFactory.CallOpts, _pool)
}

// GetDecimals is a free data retrieval call binding the contract method 0x52b51555.
//
// Solidity: function get_decimals(address _pool) view returns(uint256[2])
func (_CurvePoolFactory *CurvePoolFactoryCaller) GetDecimals(opts *bind.CallOpts, _pool common.Address) ([2]*big.Int, error) {
	var out []interface{}
	err := _CurvePoolFactory.contract.Call(opts, &out, "get_decimals", _pool)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetDecimals is a free data retrieval call binding the contract method 0x52b51555.
//
// Solidity: function get_decimals(address _pool) view returns(uint256[2])
func (_CurvePoolFactory *CurvePoolFactorySession) GetDecimals(_pool common.Address) ([2]*big.Int, error) {
	return _CurvePoolFactory.Contract.GetDecimals(&_CurvePoolFactory.CallOpts, _pool)
}

// GetDecimals is a free data retrieval call binding the contract method 0x52b51555.
//
// Solidity: function get_decimals(address _pool) view returns(uint256[2])
func (_CurvePoolFactory *CurvePoolFactoryCallerSession) GetDecimals(_pool common.Address) ([2]*big.Int, error) {
	return _CurvePoolFactory.Contract.GetDecimals(&_CurvePoolFactory.CallOpts, _pool)
}

// GetEthIndex is a free data retrieval call binding the contract method 0xccb15605.
//
// Solidity: function get_eth_index(address _pool) view returns(uint256)
func (_CurvePoolFactory *CurvePoolFactoryCaller) GetEthIndex(opts *bind.CallOpts, _pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolFactory.contract.Call(opts, &out, "get_eth_index", _pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthIndex is a free data retrieval call binding the contract method 0xccb15605.
//
// Solidity: function get_eth_index(address _pool) view returns(uint256)
func (_CurvePoolFactory *CurvePoolFactorySession) GetEthIndex(_pool common.Address) (*big.Int, error) {
	return _CurvePoolFactory.Contract.GetEthIndex(&_CurvePoolFactory.CallOpts, _pool)
}

// GetEthIndex is a free data retrieval call binding the contract method 0xccb15605.
//
// Solidity: function get_eth_index(address _pool) view returns(uint256)
func (_CurvePoolFactory *CurvePoolFactoryCallerSession) GetEthIndex(_pool common.Address) (*big.Int, error) {
	return _CurvePoolFactory.Contract.GetEthIndex(&_CurvePoolFactory.CallOpts, _pool)
}

// GetGauge is a free data retrieval call binding the contract method 0xdaf297b9.
//
// Solidity: function get_gauge(address _pool) view returns(address)
func (_CurvePoolFactory *CurvePoolFactoryCaller) GetGauge(opts *bind.CallOpts, _pool common.Address) (common.Address, error) {
	var out []interface{}
	err := _CurvePoolFactory.contract.Call(opts, &out, "get_gauge", _pool)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGauge is a free data retrieval call binding the contract method 0xdaf297b9.
//
// Solidity: function get_gauge(address _pool) view returns(address)
func (_CurvePoolFactory *CurvePoolFactorySession) GetGauge(_pool common.Address) (common.Address, error) {
	return _CurvePoolFactory.Contract.GetGauge(&_CurvePoolFactory.CallOpts, _pool)
}

// GetGauge is a free data retrieval call binding the contract method 0xdaf297b9.
//
// Solidity: function get_gauge(address _pool) view returns(address)
func (_CurvePoolFactory *CurvePoolFactoryCallerSession) GetGauge(_pool common.Address) (common.Address, error) {
	return _CurvePoolFactory.Contract.GetGauge(&_CurvePoolFactory.CallOpts, _pool)
}

// GetToken is a free data retrieval call binding the contract method 0x977d9122.
//
// Solidity: function get_token(address _pool) view returns(address)
func (_CurvePoolFactory *CurvePoolFactoryCaller) GetToken(opts *bind.CallOpts, _pool common.Address) (common.Address, error) {
	var out []interface{}
	err := _CurvePoolFactory.contract.Call(opts, &out, "get_token", _pool)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetToken is a free data retrieval call binding the contract method 0x977d9122.
//
// Solidity: function get_token(address _pool) view returns(address)
func (_CurvePoolFactory *CurvePoolFactorySession) GetToken(_pool common.Address) (common.Address, error) {
	return _CurvePoolFactory.Contract.GetToken(&_CurvePoolFactory.CallOpts, _pool)
}

// GetToken is a free data retrieval call binding the contract method 0x977d9122.
//
// Solidity: function get_token(address _pool) view returns(address)
func (_CurvePoolFactory *CurvePoolFactoryCallerSession) GetToken(_pool common.Address) (common.Address, error) {
	return _CurvePoolFactory.Contract.GetToken(&_CurvePoolFactory.CallOpts, _pool)
}

// PoolCount is a free data retrieval call binding the contract method 0x956aae3a.
//
// Solidity: function pool_count() view returns(uint256)
func (_CurvePoolFactory *CurvePoolFactoryCaller) PoolCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolFactory.contract.Call(opts, &out, "pool_count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolCount is a free data retrieval call binding the contract method 0x956aae3a.
//
// Solidity: function pool_count() view returns(uint256)
func (_CurvePoolFactory *CurvePoolFactorySession) PoolCount() (*big.Int, error) {
	return _CurvePoolFactory.Contract.PoolCount(&_CurvePoolFactory.CallOpts)
}

// PoolCount is a free data retrieval call binding the contract method 0x956aae3a.
//
// Solidity: function pool_count() view returns(uint256)
func (_CurvePoolFactory *CurvePoolFactoryCallerSession) PoolCount() (*big.Int, error) {
	return _CurvePoolFactory.Contract.PoolCount(&_CurvePoolFactory.CallOpts)
}

// PoolImplementation is a free data retrieval call binding the contract method 0x2489a2c3.
//
// Solidity: function pool_implementation() view returns(address)
func (_CurvePoolFactory *CurvePoolFactoryCaller) PoolImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurvePoolFactory.contract.Call(opts, &out, "pool_implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolImplementation is a free data retrieval call binding the contract method 0x2489a2c3.
//
// Solidity: function pool_implementation() view returns(address)
func (_CurvePoolFactory *CurvePoolFactorySession) PoolImplementation() (common.Address, error) {
	return _CurvePoolFactory.Contract.PoolImplementation(&_CurvePoolFactory.CallOpts)
}

// PoolImplementation is a free data retrieval call binding the contract method 0x2489a2c3.
//
// Solidity: function pool_implementation() view returns(address)
func (_CurvePoolFactory *CurvePoolFactoryCallerSession) PoolImplementation() (common.Address, error) {
	return _CurvePoolFactory.Contract.PoolImplementation(&_CurvePoolFactory.CallOpts)
}

// PoolList is a free data retrieval call binding the contract method 0x3a1d5d8e.
//
// Solidity: function pool_list(uint256 arg0) view returns(address)
func (_CurvePoolFactory *CurvePoolFactoryCaller) PoolList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurvePoolFactory.contract.Call(opts, &out, "pool_list", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolList is a free data retrieval call binding the contract method 0x3a1d5d8e.
//
// Solidity: function pool_list(uint256 arg0) view returns(address)
func (_CurvePoolFactory *CurvePoolFactorySession) PoolList(arg0 *big.Int) (common.Address, error) {
	return _CurvePoolFactory.Contract.PoolList(&_CurvePoolFactory.CallOpts, arg0)
}

// PoolList is a free data retrieval call binding the contract method 0x3a1d5d8e.
//
// Solidity: function pool_list(uint256 arg0) view returns(address)
func (_CurvePoolFactory *CurvePoolFactoryCallerSession) PoolList(arg0 *big.Int) (common.Address, error) {
	return _CurvePoolFactory.Contract.PoolList(&_CurvePoolFactory.CallOpts, arg0)
}

// TokenImplementation is a free data retrieval call binding the contract method 0x35214d81.
//
// Solidity: function token_implementation() view returns(address)
func (_CurvePoolFactory *CurvePoolFactoryCaller) TokenImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurvePoolFactory.contract.Call(opts, &out, "token_implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenImplementation is a free data retrieval call binding the contract method 0x35214d81.
//
// Solidity: function token_implementation() view returns(address)
func (_CurvePoolFactory *CurvePoolFactorySession) TokenImplementation() (common.Address, error) {
	return _CurvePoolFactory.Contract.TokenImplementation(&_CurvePoolFactory.CallOpts)
}

// TokenImplementation is a free data retrieval call binding the contract method 0x35214d81.
//
// Solidity: function token_implementation() view returns(address)
func (_CurvePoolFactory *CurvePoolFactoryCallerSession) TokenImplementation() (common.Address, error) {
	return _CurvePoolFactory.Contract.TokenImplementation(&_CurvePoolFactory.CallOpts)
}

// AcceptTransferOwnership is a paid mutator transaction binding the contract method 0xe5ea47b8.
//
// Solidity: function accept_transfer_ownership() returns()
func (_CurvePoolFactory *CurvePoolFactoryTransactor) AcceptTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurvePoolFactory.contract.Transact(opts, "accept_transfer_ownership")
}

// AcceptTransferOwnership is a paid mutator transaction binding the contract method 0xe5ea47b8.
//
// Solidity: function accept_transfer_ownership() returns()
func (_CurvePoolFactory *CurvePoolFactorySession) AcceptTransferOwnership() (*types.Transaction, error) {
	return _CurvePoolFactory.Contract.AcceptTransferOwnership(&_CurvePoolFactory.TransactOpts)
}

// AcceptTransferOwnership is a paid mutator transaction binding the contract method 0xe5ea47b8.
//
// Solidity: function accept_transfer_ownership() returns()
func (_CurvePoolFactory *CurvePoolFactoryTransactorSession) AcceptTransferOwnership() (*types.Transaction, error) {
	return _CurvePoolFactory.Contract.AcceptTransferOwnership(&_CurvePoolFactory.TransactOpts)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _addr) returns()
func (_CurvePoolFactory *CurvePoolFactoryTransactor) CommitTransferOwnership(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _CurvePoolFactory.contract.Transact(opts, "commit_transfer_ownership", _addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _addr) returns()
func (_CurvePoolFactory *CurvePoolFactorySession) CommitTransferOwnership(_addr common.Address) (*types.Transaction, error) {
	return _CurvePoolFactory.Contract.CommitTransferOwnership(&_CurvePoolFactory.TransactOpts, _addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _addr) returns()
func (_CurvePoolFactory *CurvePoolFactoryTransactorSession) CommitTransferOwnership(_addr common.Address) (*types.Transaction, error) {
	return _CurvePoolFactory.Contract.CommitTransferOwnership(&_CurvePoolFactory.TransactOpts, _addr)
}

// DeployGauge is a paid mutator transaction binding the contract method 0x96bebb34.
//
// Solidity: function deploy_gauge(address _pool) returns(address)
func (_CurvePoolFactory *CurvePoolFactoryTransactor) DeployGauge(opts *bind.TransactOpts, _pool common.Address) (*types.Transaction, error) {
	return _CurvePoolFactory.contract.Transact(opts, "deploy_gauge", _pool)
}

// DeployGauge is a paid mutator transaction binding the contract method 0x96bebb34.
//
// Solidity: function deploy_gauge(address _pool) returns(address)
func (_CurvePoolFactory *CurvePoolFactorySession) DeployGauge(_pool common.Address) (*types.Transaction, error) {
	return _CurvePoolFactory.Contract.DeployGauge(&_CurvePoolFactory.TransactOpts, _pool)
}

// DeployGauge is a paid mutator transaction binding the contract method 0x96bebb34.
//
// Solidity: function deploy_gauge(address _pool) returns(address)
func (_CurvePoolFactory *CurvePoolFactoryTransactorSession) DeployGauge(_pool common.Address) (*types.Transaction, error) {
	return _CurvePoolFactory.Contract.DeployGauge(&_CurvePoolFactory.TransactOpts, _pool)
}

// DeployPool is a paid mutator transaction binding the contract method 0xc955fa04.
//
// Solidity: function deploy_pool(string _name, string _symbol, address[2] _coins, uint256 A, uint256 gamma, uint256 mid_fee, uint256 out_fee, uint256 allowed_extra_profit, uint256 fee_gamma, uint256 adjustment_step, uint256 admin_fee, uint256 ma_half_time, uint256 initial_price) returns(address)
func (_CurvePoolFactory *CurvePoolFactoryTransactor) DeployPool(opts *bind.TransactOpts, _name string, _symbol string, _coins [2]common.Address, A *big.Int, gamma *big.Int, mid_fee *big.Int, out_fee *big.Int, allowed_extra_profit *big.Int, fee_gamma *big.Int, adjustment_step *big.Int, admin_fee *big.Int, ma_half_time *big.Int, initial_price *big.Int) (*types.Transaction, error) {
	return _CurvePoolFactory.contract.Transact(opts, "deploy_pool", _name, _symbol, _coins, A, gamma, mid_fee, out_fee, allowed_extra_profit, fee_gamma, adjustment_step, admin_fee, ma_half_time, initial_price)
}

// DeployPool is a paid mutator transaction binding the contract method 0xc955fa04.
//
// Solidity: function deploy_pool(string _name, string _symbol, address[2] _coins, uint256 A, uint256 gamma, uint256 mid_fee, uint256 out_fee, uint256 allowed_extra_profit, uint256 fee_gamma, uint256 adjustment_step, uint256 admin_fee, uint256 ma_half_time, uint256 initial_price) returns(address)
func (_CurvePoolFactory *CurvePoolFactorySession) DeployPool(_name string, _symbol string, _coins [2]common.Address, A *big.Int, gamma *big.Int, mid_fee *big.Int, out_fee *big.Int, allowed_extra_profit *big.Int, fee_gamma *big.Int, adjustment_step *big.Int, admin_fee *big.Int, ma_half_time *big.Int, initial_price *big.Int) (*types.Transaction, error) {
	return _CurvePoolFactory.Contract.DeployPool(&_CurvePoolFactory.TransactOpts, _name, _symbol, _coins, A, gamma, mid_fee, out_fee, allowed_extra_profit, fee_gamma, adjustment_step, admin_fee, ma_half_time, initial_price)
}

// DeployPool is a paid mutator transaction binding the contract method 0xc955fa04.
//
// Solidity: function deploy_pool(string _name, string _symbol, address[2] _coins, uint256 A, uint256 gamma, uint256 mid_fee, uint256 out_fee, uint256 allowed_extra_profit, uint256 fee_gamma, uint256 adjustment_step, uint256 admin_fee, uint256 ma_half_time, uint256 initial_price) returns(address)
func (_CurvePoolFactory *CurvePoolFactoryTransactorSession) DeployPool(_name string, _symbol string, _coins [2]common.Address, A *big.Int, gamma *big.Int, mid_fee *big.Int, out_fee *big.Int, allowed_extra_profit *big.Int, fee_gamma *big.Int, adjustment_step *big.Int, admin_fee *big.Int, ma_half_time *big.Int, initial_price *big.Int) (*types.Transaction, error) {
	return _CurvePoolFactory.Contract.DeployPool(&_CurvePoolFactory.TransactOpts, _name, _symbol, _coins, A, gamma, mid_fee, out_fee, allowed_extra_profit, fee_gamma, adjustment_step, admin_fee, ma_half_time, initial_price)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xe41ab771.
//
// Solidity: function set_fee_receiver(address _fee_receiver) returns()
func (_CurvePoolFactory *CurvePoolFactoryTransactor) SetFeeReceiver(opts *bind.TransactOpts, _fee_receiver common.Address) (*types.Transaction, error) {
	return _CurvePoolFactory.contract.Transact(opts, "set_fee_receiver", _fee_receiver)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xe41ab771.
//
// Solidity: function set_fee_receiver(address _fee_receiver) returns()
func (_CurvePoolFactory *CurvePoolFactorySession) SetFeeReceiver(_fee_receiver common.Address) (*types.Transaction, error) {
	return _CurvePoolFactory.Contract.SetFeeReceiver(&_CurvePoolFactory.TransactOpts, _fee_receiver)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xe41ab771.
//
// Solidity: function set_fee_receiver(address _fee_receiver) returns()
func (_CurvePoolFactory *CurvePoolFactoryTransactorSession) SetFeeReceiver(_fee_receiver common.Address) (*types.Transaction, error) {
	return _CurvePoolFactory.Contract.SetFeeReceiver(&_CurvePoolFactory.TransactOpts, _fee_receiver)
}

// SetGaugeImplementation is a paid mutator transaction binding the contract method 0x8f03182c.
//
// Solidity: function set_gauge_implementation(address _gauge_implementation) returns()
func (_CurvePoolFactory *CurvePoolFactoryTransactor) SetGaugeImplementation(opts *bind.TransactOpts, _gauge_implementation common.Address) (*types.Transaction, error) {
	return _CurvePoolFactory.contract.Transact(opts, "set_gauge_implementation", _gauge_implementation)
}

// SetGaugeImplementation is a paid mutator transaction binding the contract method 0x8f03182c.
//
// Solidity: function set_gauge_implementation(address _gauge_implementation) returns()
func (_CurvePoolFactory *CurvePoolFactorySession) SetGaugeImplementation(_gauge_implementation common.Address) (*types.Transaction, error) {
	return _CurvePoolFactory.Contract.SetGaugeImplementation(&_CurvePoolFactory.TransactOpts, _gauge_implementation)
}

// SetGaugeImplementation is a paid mutator transaction binding the contract method 0x8f03182c.
//
// Solidity: function set_gauge_implementation(address _gauge_implementation) returns()
func (_CurvePoolFactory *CurvePoolFactoryTransactorSession) SetGaugeImplementation(_gauge_implementation common.Address) (*types.Transaction, error) {
	return _CurvePoolFactory.Contract.SetGaugeImplementation(&_CurvePoolFactory.TransactOpts, _gauge_implementation)
}

// SetPoolImplementation is a paid mutator transaction binding the contract method 0x9ed796d0.
//
// Solidity: function set_pool_implementation(address _pool_implementation) returns()
func (_CurvePoolFactory *CurvePoolFactoryTransactor) SetPoolImplementation(opts *bind.TransactOpts, _pool_implementation common.Address) (*types.Transaction, error) {
	return _CurvePoolFactory.contract.Transact(opts, "set_pool_implementation", _pool_implementation)
}

// SetPoolImplementation is a paid mutator transaction binding the contract method 0x9ed796d0.
//
// Solidity: function set_pool_implementation(address _pool_implementation) returns()
func (_CurvePoolFactory *CurvePoolFactorySession) SetPoolImplementation(_pool_implementation common.Address) (*types.Transaction, error) {
	return _CurvePoolFactory.Contract.SetPoolImplementation(&_CurvePoolFactory.TransactOpts, _pool_implementation)
}

// SetPoolImplementation is a paid mutator transaction binding the contract method 0x9ed796d0.
//
// Solidity: function set_pool_implementation(address _pool_implementation) returns()
func (_CurvePoolFactory *CurvePoolFactoryTransactorSession) SetPoolImplementation(_pool_implementation common.Address) (*types.Transaction, error) {
	return _CurvePoolFactory.Contract.SetPoolImplementation(&_CurvePoolFactory.TransactOpts, _pool_implementation)
}

// SetTokenImplementation is a paid mutator transaction binding the contract method 0x653023c2.
//
// Solidity: function set_token_implementation(address _token_implementation) returns()
func (_CurvePoolFactory *CurvePoolFactoryTransactor) SetTokenImplementation(opts *bind.TransactOpts, _token_implementation common.Address) (*types.Transaction, error) {
	return _CurvePoolFactory.contract.Transact(opts, "set_token_implementation", _token_implementation)
}

// SetTokenImplementation is a paid mutator transaction binding the contract method 0x653023c2.
//
// Solidity: function set_token_implementation(address _token_implementation) returns()
func (_CurvePoolFactory *CurvePoolFactorySession) SetTokenImplementation(_token_implementation common.Address) (*types.Transaction, error) {
	return _CurvePoolFactory.Contract.SetTokenImplementation(&_CurvePoolFactory.TransactOpts, _token_implementation)
}

// SetTokenImplementation is a paid mutator transaction binding the contract method 0x653023c2.
//
// Solidity: function set_token_implementation(address _token_implementation) returns()
func (_CurvePoolFactory *CurvePoolFactoryTransactorSession) SetTokenImplementation(_token_implementation common.Address) (*types.Transaction, error) {
	return _CurvePoolFactory.Contract.SetTokenImplementation(&_CurvePoolFactory.TransactOpts, _token_implementation)
}

// CurvePoolFactoryCryptoPoolDeployedIterator is returned from FilterCryptoPoolDeployed and is used to iterate over the raw logs and unpacked data for CryptoPoolDeployed events raised by the CurvePoolFactory contract.
type CurvePoolFactoryCryptoPoolDeployedIterator struct {
	Event *CurvePoolFactoryCryptoPoolDeployed // Event containing the contract specifics and raw log

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
func (it *CurvePoolFactoryCryptoPoolDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolFactoryCryptoPoolDeployed)
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
		it.Event = new(CurvePoolFactoryCryptoPoolDeployed)
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
func (it *CurvePoolFactoryCryptoPoolDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolFactoryCryptoPoolDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolFactoryCryptoPoolDeployed represents a CryptoPoolDeployed event raised by the CurvePoolFactory contract.
type CurvePoolFactoryCryptoPoolDeployed struct {
	Token              common.Address
	Coins              [2]common.Address
	A                  *big.Int
	Gamma              *big.Int
	MidFee             *big.Int
	OutFee             *big.Int
	AllowedExtraProfit *big.Int
	FeeGamma           *big.Int
	AdjustmentStep     *big.Int
	AdminFee           *big.Int
	MaHalfTime         *big.Int
	InitialPrice       *big.Int
	Deployer           common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterCryptoPoolDeployed is a free log retrieval operation binding the contract event 0x0394cb40d7dbe28dad1d4ee890bdd35bbb0d89e17924a80a542535e83d54ba14.
//
// Solidity: event CryptoPoolDeployed(address token, address[2] coins, uint256 A, uint256 gamma, uint256 mid_fee, uint256 out_fee, uint256 allowed_extra_profit, uint256 fee_gamma, uint256 adjustment_step, uint256 admin_fee, uint256 ma_half_time, uint256 initial_price, address deployer)
func (_CurvePoolFactory *CurvePoolFactoryFilterer) FilterCryptoPoolDeployed(opts *bind.FilterOpts) (*CurvePoolFactoryCryptoPoolDeployedIterator, error) {

	logs, sub, err := _CurvePoolFactory.contract.FilterLogs(opts, "CryptoPoolDeployed")
	if err != nil {
		return nil, err
	}
	return &CurvePoolFactoryCryptoPoolDeployedIterator{contract: _CurvePoolFactory.contract, event: "CryptoPoolDeployed", logs: logs, sub: sub}, nil
}

// WatchCryptoPoolDeployed is a free log subscription operation binding the contract event 0x0394cb40d7dbe28dad1d4ee890bdd35bbb0d89e17924a80a542535e83d54ba14.
//
// Solidity: event CryptoPoolDeployed(address token, address[2] coins, uint256 A, uint256 gamma, uint256 mid_fee, uint256 out_fee, uint256 allowed_extra_profit, uint256 fee_gamma, uint256 adjustment_step, uint256 admin_fee, uint256 ma_half_time, uint256 initial_price, address deployer)
func (_CurvePoolFactory *CurvePoolFactoryFilterer) WatchCryptoPoolDeployed(opts *bind.WatchOpts, sink chan<- *CurvePoolFactoryCryptoPoolDeployed) (event.Subscription, error) {

	logs, sub, err := _CurvePoolFactory.contract.WatchLogs(opts, "CryptoPoolDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolFactoryCryptoPoolDeployed)
				if err := _CurvePoolFactory.contract.UnpackLog(event, "CryptoPoolDeployed", log); err != nil {
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

// ParseCryptoPoolDeployed is a log parse operation binding the contract event 0x0394cb40d7dbe28dad1d4ee890bdd35bbb0d89e17924a80a542535e83d54ba14.
//
// Solidity: event CryptoPoolDeployed(address token, address[2] coins, uint256 A, uint256 gamma, uint256 mid_fee, uint256 out_fee, uint256 allowed_extra_profit, uint256 fee_gamma, uint256 adjustment_step, uint256 admin_fee, uint256 ma_half_time, uint256 initial_price, address deployer)
func (_CurvePoolFactory *CurvePoolFactoryFilterer) ParseCryptoPoolDeployed(log types.Log) (*CurvePoolFactoryCryptoPoolDeployed, error) {
	event := new(CurvePoolFactoryCryptoPoolDeployed)
	if err := _CurvePoolFactory.contract.UnpackLog(event, "CryptoPoolDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolFactoryLiquidityGaugeDeployedIterator is returned from FilterLiquidityGaugeDeployed and is used to iterate over the raw logs and unpacked data for LiquidityGaugeDeployed events raised by the CurvePoolFactory contract.
type CurvePoolFactoryLiquidityGaugeDeployedIterator struct {
	Event *CurvePoolFactoryLiquidityGaugeDeployed // Event containing the contract specifics and raw log

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
func (it *CurvePoolFactoryLiquidityGaugeDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolFactoryLiquidityGaugeDeployed)
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
		it.Event = new(CurvePoolFactoryLiquidityGaugeDeployed)
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
func (it *CurvePoolFactoryLiquidityGaugeDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolFactoryLiquidityGaugeDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolFactoryLiquidityGaugeDeployed represents a LiquidityGaugeDeployed event raised by the CurvePoolFactory contract.
type CurvePoolFactoryLiquidityGaugeDeployed struct {
	Pool  common.Address
	Token common.Address
	Gauge common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLiquidityGaugeDeployed is a free log retrieval operation binding the contract event 0x1d6247eae69b5feb96b30be78552f35de45f61fdb6d6d7e1b08aae159b6226af.
//
// Solidity: event LiquidityGaugeDeployed(address pool, address token, address gauge)
func (_CurvePoolFactory *CurvePoolFactoryFilterer) FilterLiquidityGaugeDeployed(opts *bind.FilterOpts) (*CurvePoolFactoryLiquidityGaugeDeployedIterator, error) {

	logs, sub, err := _CurvePoolFactory.contract.FilterLogs(opts, "LiquidityGaugeDeployed")
	if err != nil {
		return nil, err
	}
	return &CurvePoolFactoryLiquidityGaugeDeployedIterator{contract: _CurvePoolFactory.contract, event: "LiquidityGaugeDeployed", logs: logs, sub: sub}, nil
}

// WatchLiquidityGaugeDeployed is a free log subscription operation binding the contract event 0x1d6247eae69b5feb96b30be78552f35de45f61fdb6d6d7e1b08aae159b6226af.
//
// Solidity: event LiquidityGaugeDeployed(address pool, address token, address gauge)
func (_CurvePoolFactory *CurvePoolFactoryFilterer) WatchLiquidityGaugeDeployed(opts *bind.WatchOpts, sink chan<- *CurvePoolFactoryLiquidityGaugeDeployed) (event.Subscription, error) {

	logs, sub, err := _CurvePoolFactory.contract.WatchLogs(opts, "LiquidityGaugeDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolFactoryLiquidityGaugeDeployed)
				if err := _CurvePoolFactory.contract.UnpackLog(event, "LiquidityGaugeDeployed", log); err != nil {
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

// ParseLiquidityGaugeDeployed is a log parse operation binding the contract event 0x1d6247eae69b5feb96b30be78552f35de45f61fdb6d6d7e1b08aae159b6226af.
//
// Solidity: event LiquidityGaugeDeployed(address pool, address token, address gauge)
func (_CurvePoolFactory *CurvePoolFactoryFilterer) ParseLiquidityGaugeDeployed(log types.Log) (*CurvePoolFactoryLiquidityGaugeDeployed, error) {
	event := new(CurvePoolFactoryLiquidityGaugeDeployed)
	if err := _CurvePoolFactory.contract.UnpackLog(event, "LiquidityGaugeDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolFactoryTransferOwnershipIterator is returned from FilterTransferOwnership and is used to iterate over the raw logs and unpacked data for TransferOwnership events raised by the CurvePoolFactory contract.
type CurvePoolFactoryTransferOwnershipIterator struct {
	Event *CurvePoolFactoryTransferOwnership // Event containing the contract specifics and raw log

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
func (it *CurvePoolFactoryTransferOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolFactoryTransferOwnership)
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
		it.Event = new(CurvePoolFactoryTransferOwnership)
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
func (it *CurvePoolFactoryTransferOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolFactoryTransferOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolFactoryTransferOwnership represents a TransferOwnership event raised by the CurvePoolFactory contract.
type CurvePoolFactoryTransferOwnership struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferOwnership is a free log retrieval operation binding the contract event 0x5c486528ec3e3f0ea91181cff8116f02bfa350e03b8b6f12e00765adbb5af85c.
//
// Solidity: event TransferOwnership(address _old_owner, address _new_owner)
func (_CurvePoolFactory *CurvePoolFactoryFilterer) FilterTransferOwnership(opts *bind.FilterOpts) (*CurvePoolFactoryTransferOwnershipIterator, error) {

	logs, sub, err := _CurvePoolFactory.contract.FilterLogs(opts, "TransferOwnership")
	if err != nil {
		return nil, err
	}
	return &CurvePoolFactoryTransferOwnershipIterator{contract: _CurvePoolFactory.contract, event: "TransferOwnership", logs: logs, sub: sub}, nil
}

// WatchTransferOwnership is a free log subscription operation binding the contract event 0x5c486528ec3e3f0ea91181cff8116f02bfa350e03b8b6f12e00765adbb5af85c.
//
// Solidity: event TransferOwnership(address _old_owner, address _new_owner)
func (_CurvePoolFactory *CurvePoolFactoryFilterer) WatchTransferOwnership(opts *bind.WatchOpts, sink chan<- *CurvePoolFactoryTransferOwnership) (event.Subscription, error) {

	logs, sub, err := _CurvePoolFactory.contract.WatchLogs(opts, "TransferOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolFactoryTransferOwnership)
				if err := _CurvePoolFactory.contract.UnpackLog(event, "TransferOwnership", log); err != nil {
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

// ParseTransferOwnership is a log parse operation binding the contract event 0x5c486528ec3e3f0ea91181cff8116f02bfa350e03b8b6f12e00765adbb5af85c.
//
// Solidity: event TransferOwnership(address _old_owner, address _new_owner)
func (_CurvePoolFactory *CurvePoolFactoryFilterer) ParseTransferOwnership(log types.Log) (*CurvePoolFactoryTransferOwnership, error) {
	event := new(CurvePoolFactoryTransferOwnership)
	if err := _CurvePoolFactory.contract.UnpackLog(event, "TransferOwnership", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolFactoryUpdateFeeReceiverIterator is returned from FilterUpdateFeeReceiver and is used to iterate over the raw logs and unpacked data for UpdateFeeReceiver events raised by the CurvePoolFactory contract.
type CurvePoolFactoryUpdateFeeReceiverIterator struct {
	Event *CurvePoolFactoryUpdateFeeReceiver // Event containing the contract specifics and raw log

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
func (it *CurvePoolFactoryUpdateFeeReceiverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolFactoryUpdateFeeReceiver)
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
		it.Event = new(CurvePoolFactoryUpdateFeeReceiver)
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
func (it *CurvePoolFactoryUpdateFeeReceiverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolFactoryUpdateFeeReceiverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolFactoryUpdateFeeReceiver represents a UpdateFeeReceiver event raised by the CurvePoolFactory contract.
type CurvePoolFactoryUpdateFeeReceiver struct {
	OldFeeReceiver common.Address
	NewFeeReceiver common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpdateFeeReceiver is a free log retrieval operation binding the contract event 0x2861448678f0be67f11bfb5481b3e3b4cfeb3acc6126ad60a05f95bfc6530666.
//
// Solidity: event UpdateFeeReceiver(address _old_fee_receiver, address _new_fee_receiver)
func (_CurvePoolFactory *CurvePoolFactoryFilterer) FilterUpdateFeeReceiver(opts *bind.FilterOpts) (*CurvePoolFactoryUpdateFeeReceiverIterator, error) {

	logs, sub, err := _CurvePoolFactory.contract.FilterLogs(opts, "UpdateFeeReceiver")
	if err != nil {
		return nil, err
	}
	return &CurvePoolFactoryUpdateFeeReceiverIterator{contract: _CurvePoolFactory.contract, event: "UpdateFeeReceiver", logs: logs, sub: sub}, nil
}

// WatchUpdateFeeReceiver is a free log subscription operation binding the contract event 0x2861448678f0be67f11bfb5481b3e3b4cfeb3acc6126ad60a05f95bfc6530666.
//
// Solidity: event UpdateFeeReceiver(address _old_fee_receiver, address _new_fee_receiver)
func (_CurvePoolFactory *CurvePoolFactoryFilterer) WatchUpdateFeeReceiver(opts *bind.WatchOpts, sink chan<- *CurvePoolFactoryUpdateFeeReceiver) (event.Subscription, error) {

	logs, sub, err := _CurvePoolFactory.contract.WatchLogs(opts, "UpdateFeeReceiver")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolFactoryUpdateFeeReceiver)
				if err := _CurvePoolFactory.contract.UnpackLog(event, "UpdateFeeReceiver", log); err != nil {
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

// ParseUpdateFeeReceiver is a log parse operation binding the contract event 0x2861448678f0be67f11bfb5481b3e3b4cfeb3acc6126ad60a05f95bfc6530666.
//
// Solidity: event UpdateFeeReceiver(address _old_fee_receiver, address _new_fee_receiver)
func (_CurvePoolFactory *CurvePoolFactoryFilterer) ParseUpdateFeeReceiver(log types.Log) (*CurvePoolFactoryUpdateFeeReceiver, error) {
	event := new(CurvePoolFactoryUpdateFeeReceiver)
	if err := _CurvePoolFactory.contract.UnpackLog(event, "UpdateFeeReceiver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolFactoryUpdateGaugeImplementationIterator is returned from FilterUpdateGaugeImplementation and is used to iterate over the raw logs and unpacked data for UpdateGaugeImplementation events raised by the CurvePoolFactory contract.
type CurvePoolFactoryUpdateGaugeImplementationIterator struct {
	Event *CurvePoolFactoryUpdateGaugeImplementation // Event containing the contract specifics and raw log

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
func (it *CurvePoolFactoryUpdateGaugeImplementationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolFactoryUpdateGaugeImplementation)
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
		it.Event = new(CurvePoolFactoryUpdateGaugeImplementation)
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
func (it *CurvePoolFactoryUpdateGaugeImplementationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolFactoryUpdateGaugeImplementationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolFactoryUpdateGaugeImplementation represents a UpdateGaugeImplementation event raised by the CurvePoolFactory contract.
type CurvePoolFactoryUpdateGaugeImplementation struct {
	OldGaugeImplementation common.Address
	NewGaugeImplementation common.Address
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterUpdateGaugeImplementation is a free log retrieval operation binding the contract event 0x1fd705f9c77053962a503f2f2f57f0862b4c3af687c25615c13817a86946c359.
//
// Solidity: event UpdateGaugeImplementation(address _old_gauge_implementation, address _new_gauge_implementation)
func (_CurvePoolFactory *CurvePoolFactoryFilterer) FilterUpdateGaugeImplementation(opts *bind.FilterOpts) (*CurvePoolFactoryUpdateGaugeImplementationIterator, error) {

	logs, sub, err := _CurvePoolFactory.contract.FilterLogs(opts, "UpdateGaugeImplementation")
	if err != nil {
		return nil, err
	}
	return &CurvePoolFactoryUpdateGaugeImplementationIterator{contract: _CurvePoolFactory.contract, event: "UpdateGaugeImplementation", logs: logs, sub: sub}, nil
}

// WatchUpdateGaugeImplementation is a free log subscription operation binding the contract event 0x1fd705f9c77053962a503f2f2f57f0862b4c3af687c25615c13817a86946c359.
//
// Solidity: event UpdateGaugeImplementation(address _old_gauge_implementation, address _new_gauge_implementation)
func (_CurvePoolFactory *CurvePoolFactoryFilterer) WatchUpdateGaugeImplementation(opts *bind.WatchOpts, sink chan<- *CurvePoolFactoryUpdateGaugeImplementation) (event.Subscription, error) {

	logs, sub, err := _CurvePoolFactory.contract.WatchLogs(opts, "UpdateGaugeImplementation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolFactoryUpdateGaugeImplementation)
				if err := _CurvePoolFactory.contract.UnpackLog(event, "UpdateGaugeImplementation", log); err != nil {
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

// ParseUpdateGaugeImplementation is a log parse operation binding the contract event 0x1fd705f9c77053962a503f2f2f57f0862b4c3af687c25615c13817a86946c359.
//
// Solidity: event UpdateGaugeImplementation(address _old_gauge_implementation, address _new_gauge_implementation)
func (_CurvePoolFactory *CurvePoolFactoryFilterer) ParseUpdateGaugeImplementation(log types.Log) (*CurvePoolFactoryUpdateGaugeImplementation, error) {
	event := new(CurvePoolFactoryUpdateGaugeImplementation)
	if err := _CurvePoolFactory.contract.UnpackLog(event, "UpdateGaugeImplementation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolFactoryUpdatePoolImplementationIterator is returned from FilterUpdatePoolImplementation and is used to iterate over the raw logs and unpacked data for UpdatePoolImplementation events raised by the CurvePoolFactory contract.
type CurvePoolFactoryUpdatePoolImplementationIterator struct {
	Event *CurvePoolFactoryUpdatePoolImplementation // Event containing the contract specifics and raw log

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
func (it *CurvePoolFactoryUpdatePoolImplementationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolFactoryUpdatePoolImplementation)
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
		it.Event = new(CurvePoolFactoryUpdatePoolImplementation)
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
func (it *CurvePoolFactoryUpdatePoolImplementationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolFactoryUpdatePoolImplementationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolFactoryUpdatePoolImplementation represents a UpdatePoolImplementation event raised by the CurvePoolFactory contract.
type CurvePoolFactoryUpdatePoolImplementation struct {
	OldPoolImplementation common.Address
	NewPoolImplementation common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterUpdatePoolImplementation is a free log retrieval operation binding the contract event 0x0617fd31aa5ab95ec80eefc1eb61a2c477aa419d1d761b4e46f5f077e47852aa.
//
// Solidity: event UpdatePoolImplementation(address _old_pool_implementation, address _new_pool_implementation)
func (_CurvePoolFactory *CurvePoolFactoryFilterer) FilterUpdatePoolImplementation(opts *bind.FilterOpts) (*CurvePoolFactoryUpdatePoolImplementationIterator, error) {

	logs, sub, err := _CurvePoolFactory.contract.FilterLogs(opts, "UpdatePoolImplementation")
	if err != nil {
		return nil, err
	}
	return &CurvePoolFactoryUpdatePoolImplementationIterator{contract: _CurvePoolFactory.contract, event: "UpdatePoolImplementation", logs: logs, sub: sub}, nil
}

// WatchUpdatePoolImplementation is a free log subscription operation binding the contract event 0x0617fd31aa5ab95ec80eefc1eb61a2c477aa419d1d761b4e46f5f077e47852aa.
//
// Solidity: event UpdatePoolImplementation(address _old_pool_implementation, address _new_pool_implementation)
func (_CurvePoolFactory *CurvePoolFactoryFilterer) WatchUpdatePoolImplementation(opts *bind.WatchOpts, sink chan<- *CurvePoolFactoryUpdatePoolImplementation) (event.Subscription, error) {

	logs, sub, err := _CurvePoolFactory.contract.WatchLogs(opts, "UpdatePoolImplementation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolFactoryUpdatePoolImplementation)
				if err := _CurvePoolFactory.contract.UnpackLog(event, "UpdatePoolImplementation", log); err != nil {
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

// ParseUpdatePoolImplementation is a log parse operation binding the contract event 0x0617fd31aa5ab95ec80eefc1eb61a2c477aa419d1d761b4e46f5f077e47852aa.
//
// Solidity: event UpdatePoolImplementation(address _old_pool_implementation, address _new_pool_implementation)
func (_CurvePoolFactory *CurvePoolFactoryFilterer) ParseUpdatePoolImplementation(log types.Log) (*CurvePoolFactoryUpdatePoolImplementation, error) {
	event := new(CurvePoolFactoryUpdatePoolImplementation)
	if err := _CurvePoolFactory.contract.UnpackLog(event, "UpdatePoolImplementation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolFactoryUpdateTokenImplementationIterator is returned from FilterUpdateTokenImplementation and is used to iterate over the raw logs and unpacked data for UpdateTokenImplementation events raised by the CurvePoolFactory contract.
type CurvePoolFactoryUpdateTokenImplementationIterator struct {
	Event *CurvePoolFactoryUpdateTokenImplementation // Event containing the contract specifics and raw log

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
func (it *CurvePoolFactoryUpdateTokenImplementationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolFactoryUpdateTokenImplementation)
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
		it.Event = new(CurvePoolFactoryUpdateTokenImplementation)
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
func (it *CurvePoolFactoryUpdateTokenImplementationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolFactoryUpdateTokenImplementationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolFactoryUpdateTokenImplementation represents a UpdateTokenImplementation event raised by the CurvePoolFactory contract.
type CurvePoolFactoryUpdateTokenImplementation struct {
	OldTokenImplementation common.Address
	NewTokenImplementation common.Address
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterUpdateTokenImplementation is a free log retrieval operation binding the contract event 0x1cc4f8e20b0cd3e5109eb156cadcfd3a5496ac0794c6085ca02afd7d710dd566.
//
// Solidity: event UpdateTokenImplementation(address _old_token_implementation, address _new_token_implementation)
func (_CurvePoolFactory *CurvePoolFactoryFilterer) FilterUpdateTokenImplementation(opts *bind.FilterOpts) (*CurvePoolFactoryUpdateTokenImplementationIterator, error) {

	logs, sub, err := _CurvePoolFactory.contract.FilterLogs(opts, "UpdateTokenImplementation")
	if err != nil {
		return nil, err
	}
	return &CurvePoolFactoryUpdateTokenImplementationIterator{contract: _CurvePoolFactory.contract, event: "UpdateTokenImplementation", logs: logs, sub: sub}, nil
}

// WatchUpdateTokenImplementation is a free log subscription operation binding the contract event 0x1cc4f8e20b0cd3e5109eb156cadcfd3a5496ac0794c6085ca02afd7d710dd566.
//
// Solidity: event UpdateTokenImplementation(address _old_token_implementation, address _new_token_implementation)
func (_CurvePoolFactory *CurvePoolFactoryFilterer) WatchUpdateTokenImplementation(opts *bind.WatchOpts, sink chan<- *CurvePoolFactoryUpdateTokenImplementation) (event.Subscription, error) {

	logs, sub, err := _CurvePoolFactory.contract.WatchLogs(opts, "UpdateTokenImplementation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolFactoryUpdateTokenImplementation)
				if err := _CurvePoolFactory.contract.UnpackLog(event, "UpdateTokenImplementation", log); err != nil {
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

// ParseUpdateTokenImplementation is a log parse operation binding the contract event 0x1cc4f8e20b0cd3e5109eb156cadcfd3a5496ac0794c6085ca02afd7d710dd566.
//
// Solidity: event UpdateTokenImplementation(address _old_token_implementation, address _new_token_implementation)
func (_CurvePoolFactory *CurvePoolFactoryFilterer) ParseUpdateTokenImplementation(log types.Log) (*CurvePoolFactoryUpdateTokenImplementation, error) {
	event := new(CurvePoolFactoryUpdateTokenImplementation)
	if err := _CurvePoolFactory.contract.UnpackLog(event, "UpdateTokenImplementation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
