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

// AeroStruct4 is an auto generated low-level Go binding around an user-defined struct.
type AeroStruct4 struct {
	Lp      common.Address
	Stable  bool
	Token0  common.Address
	Token1  common.Address
	Factory common.Address
}

// AeroStruct0 is an auto generated low-level Go binding around an user-defined struct.
type AeroStruct0 struct {
	Lp               common.Address
	Symbol           string
	Decimals         uint8
	Stable           bool
	TotalSupply      *big.Int
	Token0           common.Address
	Reserve0         *big.Int
	Claimable0       *big.Int
	Token1           common.Address
	Reserve1         *big.Int
	Claimable1       *big.Int
	Gauge            common.Address
	GaugeTotalSupply *big.Int
	GaugeAlive       bool
	Fee              common.Address
	Bribe            common.Address
	Factory          common.Address
	Emissions        *big.Int
	EmissionsToken   common.Address
	AccountBalance   *big.Int
	AccountEarned    *big.Int
	AccountStaked    *big.Int
	PoolFee          *big.Int
	Token0Fees       *big.Int
	Token1Fees       *big.Int
}

// AeroStruct5 is an auto generated low-level Go binding around an user-defined struct.
type AeroStruct5 struct {
	TokenAddress   common.Address
	Symbol         string
	Decimals       uint8
	AccountBalance *big.Int
	Listed         bool
}

// AeroStruct1 is an auto generated low-level Go binding around an user-defined struct.
type AeroStruct1 struct {
	Token  common.Address
	Amount *big.Int
}

// AeroStruct3 is an auto generated low-level Go binding around an user-defined struct.
type AeroStruct3 struct {
	VenftId *big.Int
	Lp      common.Address
	Amount  *big.Int
	Token   common.Address
	Fee     common.Address
	Bribe   common.Address
}

// AeroStruct2 is an auto generated low-level Go binding around an user-defined struct.
type AeroStruct2 struct {
	Ts        *big.Int
	Lp        common.Address
	Votes     *big.Int
	Emissions *big.Int
	Bribes    []AeroStruct1
	Fees      []AeroStruct1
}

// AeroSugarMetaData contains all meta data concerning the AeroSugar contract.
var AeroSugarMetaData = &bind.MetaData{
	ABI: "[{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"_voter\",\"type\":\"address\"},{\"name\":\"_registry\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"forSwaps\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"lp\",\"type\":\"address\"},{\"name\":\"stable\",\"type\":\"bool\"},{\"name\":\"token0\",\"type\":\"address\"},{\"name\":\"token1\",\"type\":\"address\"},{\"name\":\"factory\",\"type\":\"address\"}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"tokens\",\"inputs\":[{\"name\":\"_limit\",\"type\":\"uint256\"},{\"name\":\"_offset\",\"type\":\"uint256\"},{\"name\":\"_account\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"token_address\",\"type\":\"address\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\"},{\"name\":\"account_balance\",\"type\":\"uint256\"},{\"name\":\"listed\",\"type\":\"bool\"}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"all\",\"inputs\":[{\"name\":\"_limit\",\"type\":\"uint256\"},{\"name\":\"_offset\",\"type\":\"uint256\"},{\"name\":\"_account\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"lp\",\"type\":\"address\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\"},{\"name\":\"stable\",\"type\":\"bool\"},{\"name\":\"total_supply\",\"type\":\"uint256\"},{\"name\":\"token0\",\"type\":\"address\"},{\"name\":\"reserve0\",\"type\":\"uint256\"},{\"name\":\"claimable0\",\"type\":\"uint256\"},{\"name\":\"token1\",\"type\":\"address\"},{\"name\":\"reserve1\",\"type\":\"uint256\"},{\"name\":\"claimable1\",\"type\":\"uint256\"},{\"name\":\"gauge\",\"type\":\"address\"},{\"name\":\"gauge_total_supply\",\"type\":\"uint256\"},{\"name\":\"gauge_alive\",\"type\":\"bool\"},{\"name\":\"fee\",\"type\":\"address\"},{\"name\":\"bribe\",\"type\":\"address\"},{\"name\":\"factory\",\"type\":\"address\"},{\"name\":\"emissions\",\"type\":\"uint256\"},{\"name\":\"emissions_token\",\"type\":\"address\"},{\"name\":\"account_balance\",\"type\":\"uint256\"},{\"name\":\"account_earned\",\"type\":\"uint256\"},{\"name\":\"account_staked\",\"type\":\"uint256\"},{\"name\":\"pool_fee\",\"type\":\"uint256\"},{\"name\":\"token0_fees\",\"type\":\"uint256\"},{\"name\":\"token1_fees\",\"type\":\"uint256\"}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"byIndex\",\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"},{\"name\":\"_account\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"components\":[{\"name\":\"lp\",\"type\":\"address\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\"},{\"name\":\"stable\",\"type\":\"bool\"},{\"name\":\"total_supply\",\"type\":\"uint256\"},{\"name\":\"token0\",\"type\":\"address\"},{\"name\":\"reserve0\",\"type\":\"uint256\"},{\"name\":\"claimable0\",\"type\":\"uint256\"},{\"name\":\"token1\",\"type\":\"address\"},{\"name\":\"reserve1\",\"type\":\"uint256\"},{\"name\":\"claimable1\",\"type\":\"uint256\"},{\"name\":\"gauge\",\"type\":\"address\"},{\"name\":\"gauge_total_supply\",\"type\":\"uint256\"},{\"name\":\"gauge_alive\",\"type\":\"bool\"},{\"name\":\"fee\",\"type\":\"address\"},{\"name\":\"bribe\",\"type\":\"address\"},{\"name\":\"factory\",\"type\":\"address\"},{\"name\":\"emissions\",\"type\":\"uint256\"},{\"name\":\"emissions_token\",\"type\":\"address\"},{\"name\":\"account_balance\",\"type\":\"uint256\"},{\"name\":\"account_earned\",\"type\":\"uint256\"},{\"name\":\"account_staked\",\"type\":\"uint256\"},{\"name\":\"pool_fee\",\"type\":\"uint256\"},{\"name\":\"token0_fees\",\"type\":\"uint256\"},{\"name\":\"token1_fees\",\"type\":\"uint256\"}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"epochsLatest\",\"inputs\":[{\"name\":\"_limit\",\"type\":\"uint256\"},{\"name\":\"_offset\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"ts\",\"type\":\"uint256\"},{\"name\":\"lp\",\"type\":\"address\"},{\"name\":\"votes\",\"type\":\"uint256\"},{\"name\":\"emissions\",\"type\":\"uint256\"},{\"name\":\"bribes\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}]},{\"name\":\"fees\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}]}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"epochsByAddress\",\"inputs\":[{\"name\":\"_limit\",\"type\":\"uint256\"},{\"name\":\"_offset\",\"type\":\"uint256\"},{\"name\":\"_address\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"ts\",\"type\":\"uint256\"},{\"name\":\"lp\",\"type\":\"address\"},{\"name\":\"votes\",\"type\":\"uint256\"},{\"name\":\"emissions\",\"type\":\"uint256\"},{\"name\":\"bribes\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}]},{\"name\":\"fees\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}]}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"rewards\",\"inputs\":[{\"name\":\"_limit\",\"type\":\"uint256\"},{\"name\":\"_offset\",\"type\":\"uint256\"},{\"name\":\"_venft_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"venft_id\",\"type\":\"uint256\"},{\"name\":\"lp\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"fee\",\"type\":\"address\"},{\"name\":\"bribe\",\"type\":\"address\"}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"rewardsByAddress\",\"inputs\":[{\"name\":\"_venft_id\",\"type\":\"uint256\"},{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"venft_id\",\"type\":\"uint256\"},{\"name\":\"lp\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"fee\",\"type\":\"address\"},{\"name\":\"bribe\",\"type\":\"address\"}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"registry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"voter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]}]",
}

// AeroSugarABI is the input ABI used to generate the binding from.
// Deprecated: Use AeroSugarMetaData.ABI instead.
var AeroSugarABI = AeroSugarMetaData.ABI

// AeroSugar is an auto generated Go binding around an Ethereum contract.
type AeroSugar struct {
	AeroSugarCaller     // Read-only binding to the contract
	AeroSugarTransactor // Write-only binding to the contract
	AeroSugarFilterer   // Log filterer for contract events
}

// AeroSugarCaller is an auto generated read-only Go binding around an Ethereum contract.
type AeroSugarCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AeroSugarTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AeroSugarTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AeroSugarFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AeroSugarFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AeroSugarSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AeroSugarSession struct {
	Contract     *AeroSugar        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AeroSugarCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AeroSugarCallerSession struct {
	Contract *AeroSugarCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// AeroSugarTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AeroSugarTransactorSession struct {
	Contract     *AeroSugarTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AeroSugarRaw is an auto generated low-level Go binding around an Ethereum contract.
type AeroSugarRaw struct {
	Contract *AeroSugar // Generic contract binding to access the raw methods on
}

// AeroSugarCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AeroSugarCallerRaw struct {
	Contract *AeroSugarCaller // Generic read-only contract binding to access the raw methods on
}

// AeroSugarTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AeroSugarTransactorRaw struct {
	Contract *AeroSugarTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAeroSugar creates a new instance of AeroSugar, bound to a specific deployed contract.
func NewAeroSugar(address common.Address, backend bind.ContractBackend) (*AeroSugar, error) {
	contract, err := bindAeroSugar(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AeroSugar{AeroSugarCaller: AeroSugarCaller{contract: contract}, AeroSugarTransactor: AeroSugarTransactor{contract: contract}, AeroSugarFilterer: AeroSugarFilterer{contract: contract}}, nil
}

// NewAeroSugarCaller creates a new read-only instance of AeroSugar, bound to a specific deployed contract.
func NewAeroSugarCaller(address common.Address, caller bind.ContractCaller) (*AeroSugarCaller, error) {
	contract, err := bindAeroSugar(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AeroSugarCaller{contract: contract}, nil
}

// NewAeroSugarTransactor creates a new write-only instance of AeroSugar, bound to a specific deployed contract.
func NewAeroSugarTransactor(address common.Address, transactor bind.ContractTransactor) (*AeroSugarTransactor, error) {
	contract, err := bindAeroSugar(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AeroSugarTransactor{contract: contract}, nil
}

// NewAeroSugarFilterer creates a new log filterer instance of AeroSugar, bound to a specific deployed contract.
func NewAeroSugarFilterer(address common.Address, filterer bind.ContractFilterer) (*AeroSugarFilterer, error) {
	contract, err := bindAeroSugar(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AeroSugarFilterer{contract: contract}, nil
}

// bindAeroSugar binds a generic wrapper to an already deployed contract.
func bindAeroSugar(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AeroSugarABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AeroSugar *AeroSugarRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AeroSugar.Contract.AeroSugarCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AeroSugar *AeroSugarRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AeroSugar.Contract.AeroSugarTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AeroSugar *AeroSugarRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AeroSugar.Contract.AeroSugarTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AeroSugar *AeroSugarCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AeroSugar.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AeroSugar *AeroSugarTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AeroSugar.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AeroSugar *AeroSugarTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AeroSugar.Contract.contract.Transact(opts, method, params...)
}

// All is a free data retrieval call binding the contract method 0xc0d0bf32.
//
// Solidity: function all(uint256 _limit, uint256 _offset, address _account) view returns((address,string,uint8,bool,uint256,address,uint256,uint256,address,uint256,uint256,address,uint256,bool,address,address,address,uint256,address,uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_AeroSugar *AeroSugarCaller) All(opts *bind.CallOpts, _limit *big.Int, _offset *big.Int, _account common.Address) ([]AeroStruct0, error) {
	var out []interface{}
	err := _AeroSugar.contract.Call(opts, &out, "all", _limit, _offset, _account)

	if err != nil {
		return *new([]AeroStruct0), err
	}

	out0 := *abi.ConvertType(out[0], new([]AeroStruct0)).(*[]AeroStruct0)

	return out0, err

}

// All is a free data retrieval call binding the contract method 0xc0d0bf32.
//
// Solidity: function all(uint256 _limit, uint256 _offset, address _account) view returns((address,string,uint8,bool,uint256,address,uint256,uint256,address,uint256,uint256,address,uint256,bool,address,address,address,uint256,address,uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_AeroSugar *AeroSugarSession) All(_limit *big.Int, _offset *big.Int, _account common.Address) ([]AeroStruct0, error) {
	return _AeroSugar.Contract.All(&_AeroSugar.CallOpts, _limit, _offset, _account)
}

// All is a free data retrieval call binding the contract method 0xc0d0bf32.
//
// Solidity: function all(uint256 _limit, uint256 _offset, address _account) view returns((address,string,uint8,bool,uint256,address,uint256,uint256,address,uint256,uint256,address,uint256,bool,address,address,address,uint256,address,uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_AeroSugar *AeroSugarCallerSession) All(_limit *big.Int, _offset *big.Int, _account common.Address) ([]AeroStruct0, error) {
	return _AeroSugar.Contract.All(&_AeroSugar.CallOpts, _limit, _offset, _account)
}

// ByIndex is a free data retrieval call binding the contract method 0xfbb70183.
//
// Solidity: function byIndex(uint256 _index, address _account) view returns((address,string,uint8,bool,uint256,address,uint256,uint256,address,uint256,uint256,address,uint256,bool,address,address,address,uint256,address,uint256,uint256,uint256,uint256,uint256,uint256))
func (_AeroSugar *AeroSugarCaller) ByIndex(opts *bind.CallOpts, _index *big.Int, _account common.Address) (AeroStruct0, error) {
	var out []interface{}
	err := _AeroSugar.contract.Call(opts, &out, "byIndex", _index, _account)

	if err != nil {
		return *new(AeroStruct0), err
	}

	out0 := *abi.ConvertType(out[0], new(AeroStruct0)).(*AeroStruct0)

	return out0, err

}

// ByIndex is a free data retrieval call binding the contract method 0xfbb70183.
//
// Solidity: function byIndex(uint256 _index, address _account) view returns((address,string,uint8,bool,uint256,address,uint256,uint256,address,uint256,uint256,address,uint256,bool,address,address,address,uint256,address,uint256,uint256,uint256,uint256,uint256,uint256))
func (_AeroSugar *AeroSugarSession) ByIndex(_index *big.Int, _account common.Address) (AeroStruct0, error) {
	return _AeroSugar.Contract.ByIndex(&_AeroSugar.CallOpts, _index, _account)
}

// ByIndex is a free data retrieval call binding the contract method 0xfbb70183.
//
// Solidity: function byIndex(uint256 _index, address _account) view returns((address,string,uint8,bool,uint256,address,uint256,uint256,address,uint256,uint256,address,uint256,bool,address,address,address,uint256,address,uint256,uint256,uint256,uint256,uint256,uint256))
func (_AeroSugar *AeroSugarCallerSession) ByIndex(_index *big.Int, _account common.Address) (AeroStruct0, error) {
	return _AeroSugar.Contract.ByIndex(&_AeroSugar.CallOpts, _index, _account)
}

// EpochsByAddress is a free data retrieval call binding the contract method 0x8878d06c.
//
// Solidity: function epochsByAddress(uint256 _limit, uint256 _offset, address _address) view returns((uint256,address,uint256,uint256,(address,uint256)[],(address,uint256)[])[])
func (_AeroSugar *AeroSugarCaller) EpochsByAddress(opts *bind.CallOpts, _limit *big.Int, _offset *big.Int, _address common.Address) ([]AeroStruct2, error) {
	var out []interface{}
	err := _AeroSugar.contract.Call(opts, &out, "epochsByAddress", _limit, _offset, _address)

	if err != nil {
		return *new([]AeroStruct2), err
	}

	out0 := *abi.ConvertType(out[0], new([]AeroStruct2)).(*[]AeroStruct2)

	return out0, err

}

// EpochsByAddress is a free data retrieval call binding the contract method 0x8878d06c.
//
// Solidity: function epochsByAddress(uint256 _limit, uint256 _offset, address _address) view returns((uint256,address,uint256,uint256,(address,uint256)[],(address,uint256)[])[])
func (_AeroSugar *AeroSugarSession) EpochsByAddress(_limit *big.Int, _offset *big.Int, _address common.Address) ([]AeroStruct2, error) {
	return _AeroSugar.Contract.EpochsByAddress(&_AeroSugar.CallOpts, _limit, _offset, _address)
}

// EpochsByAddress is a free data retrieval call binding the contract method 0x8878d06c.
//
// Solidity: function epochsByAddress(uint256 _limit, uint256 _offset, address _address) view returns((uint256,address,uint256,uint256,(address,uint256)[],(address,uint256)[])[])
func (_AeroSugar *AeroSugarCallerSession) EpochsByAddress(_limit *big.Int, _offset *big.Int, _address common.Address) ([]AeroStruct2, error) {
	return _AeroSugar.Contract.EpochsByAddress(&_AeroSugar.CallOpts, _limit, _offset, _address)
}

// EpochsLatest is a free data retrieval call binding the contract method 0xd94b9bc6.
//
// Solidity: function epochsLatest(uint256 _limit, uint256 _offset) view returns((uint256,address,uint256,uint256,(address,uint256)[],(address,uint256)[])[])
func (_AeroSugar *AeroSugarCaller) EpochsLatest(opts *bind.CallOpts, _limit *big.Int, _offset *big.Int) ([]AeroStruct2, error) {
	var out []interface{}
	err := _AeroSugar.contract.Call(opts, &out, "epochsLatest", _limit, _offset)

	if err != nil {
		return *new([]AeroStruct2), err
	}

	out0 := *abi.ConvertType(out[0], new([]AeroStruct2)).(*[]AeroStruct2)

	return out0, err

}

// EpochsLatest is a free data retrieval call binding the contract method 0xd94b9bc6.
//
// Solidity: function epochsLatest(uint256 _limit, uint256 _offset) view returns((uint256,address,uint256,uint256,(address,uint256)[],(address,uint256)[])[])
func (_AeroSugar *AeroSugarSession) EpochsLatest(_limit *big.Int, _offset *big.Int) ([]AeroStruct2, error) {
	return _AeroSugar.Contract.EpochsLatest(&_AeroSugar.CallOpts, _limit, _offset)
}

// EpochsLatest is a free data retrieval call binding the contract method 0xd94b9bc6.
//
// Solidity: function epochsLatest(uint256 _limit, uint256 _offset) view returns((uint256,address,uint256,uint256,(address,uint256)[],(address,uint256)[])[])
func (_AeroSugar *AeroSugarCallerSession) EpochsLatest(_limit *big.Int, _offset *big.Int) ([]AeroStruct2, error) {
	return _AeroSugar.Contract.EpochsLatest(&_AeroSugar.CallOpts, _limit, _offset)
}

// ForSwaps is a free data retrieval call binding the contract method 0xcc259a0e.
//
// Solidity: function forSwaps() view returns((address,bool,address,address,address)[])
func (_AeroSugar *AeroSugarCaller) ForSwaps(opts *bind.CallOpts) ([]AeroStruct4, error) {
	var out []interface{}
	err := _AeroSugar.contract.Call(opts, &out, "forSwaps")

	if err != nil {
		return *new([]AeroStruct4), err
	}

	out0 := *abi.ConvertType(out[0], new([]AeroStruct4)).(*[]AeroStruct4)

	return out0, err

}

// ForSwaps is a free data retrieval call binding the contract method 0xcc259a0e.
//
// Solidity: function forSwaps() view returns((address,bool,address,address,address)[])
func (_AeroSugar *AeroSugarSession) ForSwaps() ([]AeroStruct4, error) {
	return _AeroSugar.Contract.ForSwaps(&_AeroSugar.CallOpts)
}

// ForSwaps is a free data retrieval call binding the contract method 0xcc259a0e.
//
// Solidity: function forSwaps() view returns((address,bool,address,address,address)[])
func (_AeroSugar *AeroSugarCallerSession) ForSwaps() ([]AeroStruct4, error) {
	return _AeroSugar.Contract.ForSwaps(&_AeroSugar.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_AeroSugar *AeroSugarCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AeroSugar.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_AeroSugar *AeroSugarSession) Registry() (common.Address, error) {
	return _AeroSugar.Contract.Registry(&_AeroSugar.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_AeroSugar *AeroSugarCallerSession) Registry() (common.Address, error) {
	return _AeroSugar.Contract.Registry(&_AeroSugar.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0xa9c57fee.
//
// Solidity: function rewards(uint256 _limit, uint256 _offset, uint256 _venft_id) view returns((uint256,address,uint256,address,address,address)[])
func (_AeroSugar *AeroSugarCaller) Rewards(opts *bind.CallOpts, _limit *big.Int, _offset *big.Int, _venft_id *big.Int) ([]AeroStruct3, error) {
	var out []interface{}
	err := _AeroSugar.contract.Call(opts, &out, "rewards", _limit, _offset, _venft_id)

	if err != nil {
		return *new([]AeroStruct3), err
	}

	out0 := *abi.ConvertType(out[0], new([]AeroStruct3)).(*[]AeroStruct3)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0xa9c57fee.
//
// Solidity: function rewards(uint256 _limit, uint256 _offset, uint256 _venft_id) view returns((uint256,address,uint256,address,address,address)[])
func (_AeroSugar *AeroSugarSession) Rewards(_limit *big.Int, _offset *big.Int, _venft_id *big.Int) ([]AeroStruct3, error) {
	return _AeroSugar.Contract.Rewards(&_AeroSugar.CallOpts, _limit, _offset, _venft_id)
}

// Rewards is a free data retrieval call binding the contract method 0xa9c57fee.
//
// Solidity: function rewards(uint256 _limit, uint256 _offset, uint256 _venft_id) view returns((uint256,address,uint256,address,address,address)[])
func (_AeroSugar *AeroSugarCallerSession) Rewards(_limit *big.Int, _offset *big.Int, _venft_id *big.Int) ([]AeroStruct3, error) {
	return _AeroSugar.Contract.Rewards(&_AeroSugar.CallOpts, _limit, _offset, _venft_id)
}

// RewardsByAddress is a free data retrieval call binding the contract method 0xcd824fb4.
//
// Solidity: function rewardsByAddress(uint256 _venft_id, address _pool) view returns((uint256,address,uint256,address,address,address)[])
func (_AeroSugar *AeroSugarCaller) RewardsByAddress(opts *bind.CallOpts, _venft_id *big.Int, _pool common.Address) ([]AeroStruct3, error) {
	var out []interface{}
	err := _AeroSugar.contract.Call(opts, &out, "rewardsByAddress", _venft_id, _pool)

	if err != nil {
		return *new([]AeroStruct3), err
	}

	out0 := *abi.ConvertType(out[0], new([]AeroStruct3)).(*[]AeroStruct3)

	return out0, err

}

// RewardsByAddress is a free data retrieval call binding the contract method 0xcd824fb4.
//
// Solidity: function rewardsByAddress(uint256 _venft_id, address _pool) view returns((uint256,address,uint256,address,address,address)[])
func (_AeroSugar *AeroSugarSession) RewardsByAddress(_venft_id *big.Int, _pool common.Address) ([]AeroStruct3, error) {
	return _AeroSugar.Contract.RewardsByAddress(&_AeroSugar.CallOpts, _venft_id, _pool)
}

// RewardsByAddress is a free data retrieval call binding the contract method 0xcd824fb4.
//
// Solidity: function rewardsByAddress(uint256 _venft_id, address _pool) view returns((uint256,address,uint256,address,address,address)[])
func (_AeroSugar *AeroSugarCallerSession) RewardsByAddress(_venft_id *big.Int, _pool common.Address) ([]AeroStruct3, error) {
	return _AeroSugar.Contract.RewardsByAddress(&_AeroSugar.CallOpts, _venft_id, _pool)
}

// Tokens is a free data retrieval call binding the contract method 0x5cc33187.
//
// Solidity: function tokens(uint256 _limit, uint256 _offset, address _account) view returns((address,string,uint8,uint256,bool)[])
func (_AeroSugar *AeroSugarCaller) Tokens(opts *bind.CallOpts, _limit *big.Int, _offset *big.Int, _account common.Address) ([]AeroStruct5, error) {
	var out []interface{}
	err := _AeroSugar.contract.Call(opts, &out, "tokens", _limit, _offset, _account)

	if err != nil {
		return *new([]AeroStruct5), err
	}

	out0 := *abi.ConvertType(out[0], new([]AeroStruct5)).(*[]AeroStruct5)

	return out0, err

}

// Tokens is a free data retrieval call binding the contract method 0x5cc33187.
//
// Solidity: function tokens(uint256 _limit, uint256 _offset, address _account) view returns((address,string,uint8,uint256,bool)[])
func (_AeroSugar *AeroSugarSession) Tokens(_limit *big.Int, _offset *big.Int, _account common.Address) ([]AeroStruct5, error) {
	return _AeroSugar.Contract.Tokens(&_AeroSugar.CallOpts, _limit, _offset, _account)
}

// Tokens is a free data retrieval call binding the contract method 0x5cc33187.
//
// Solidity: function tokens(uint256 _limit, uint256 _offset, address _account) view returns((address,string,uint8,uint256,bool)[])
func (_AeroSugar *AeroSugarCallerSession) Tokens(_limit *big.Int, _offset *big.Int, _account common.Address) ([]AeroStruct5, error) {
	return _AeroSugar.Contract.Tokens(&_AeroSugar.CallOpts, _limit, _offset, _account)
}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_AeroSugar *AeroSugarCaller) Voter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AeroSugar.contract.Call(opts, &out, "voter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_AeroSugar *AeroSugarSession) Voter() (common.Address, error) {
	return _AeroSugar.Contract.Voter(&_AeroSugar.CallOpts)
}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_AeroSugar *AeroSugarCallerSession) Voter() (common.Address, error) {
	return _AeroSugar.Contract.Voter(&_AeroSugar.CallOpts)
}
