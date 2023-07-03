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

// Struct2 is an auto generated low-level Go binding around an user-defined struct.
type Struct2 struct {
	PairAddress      common.Address
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
	WrappedBribe     common.Address
	Emissions        *big.Int
	EmissionsToken   common.Address
	AccountBalance   *big.Int
	AccountEarned    *big.Int
	AccountStaked    *big.Int
}

// Struct3 is an auto generated low-level Go binding around an user-defined struct.
type Struct3 struct {
	TokenAddress   common.Address
	Symbol         string
	Decimals       uint8
	AccountBalance *big.Int
	Listed         bool
}

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	Token  common.Address
	Amount *big.Int
}

// Struct1 is an auto generated low-level Go binding around an user-defined struct.
type Struct1 struct {
	Ts          *big.Int
	PairAddress common.Address
	Votes       *big.Int
	Emissions   *big.Int
	Bribes      []Struct0
	Fees        []Struct0
}

// VeloSugarMetaData contains all meta data concerning the VeloSugar contract.
var VeloSugarMetaData = &bind.MetaData{
	ABI: "[{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"setup\",\"inputs\":[{\"name\":\"_voter\",\"type\":\"address\"},{\"name\":\"_wrapped_bribe_factory\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"tokens\",\"inputs\":[{\"name\":\"_limit\",\"type\":\"uint256\"},{\"name\":\"_offset\",\"type\":\"uint256\"},{\"name\":\"_account\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"token_address\",\"type\":\"address\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\"},{\"name\":\"account_balance\",\"type\":\"uint256\"},{\"name\":\"listed\",\"type\":\"bool\"}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"all\",\"inputs\":[{\"name\":\"_limit\",\"type\":\"uint256\"},{\"name\":\"_offset\",\"type\":\"uint256\"},{\"name\":\"_account\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"pair_address\",\"type\":\"address\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\"},{\"name\":\"stable\",\"type\":\"bool\"},{\"name\":\"total_supply\",\"type\":\"uint256\"},{\"name\":\"token0\",\"type\":\"address\"},{\"name\":\"reserve0\",\"type\":\"uint256\"},{\"name\":\"claimable0\",\"type\":\"uint256\"},{\"name\":\"token1\",\"type\":\"address\"},{\"name\":\"reserve1\",\"type\":\"uint256\"},{\"name\":\"claimable1\",\"type\":\"uint256\"},{\"name\":\"gauge\",\"type\":\"address\"},{\"name\":\"gauge_total_supply\",\"type\":\"uint256\"},{\"name\":\"gauge_alive\",\"type\":\"bool\"},{\"name\":\"fee\",\"type\":\"address\"},{\"name\":\"bribe\",\"type\":\"address\"},{\"name\":\"wrapped_bribe\",\"type\":\"address\"},{\"name\":\"emissions\",\"type\":\"uint256\"},{\"name\":\"emissions_token\",\"type\":\"address\"},{\"name\":\"account_balance\",\"type\":\"uint256\"},{\"name\":\"account_earned\",\"type\":\"uint256\"},{\"name\":\"account_staked\",\"type\":\"uint256\"}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"byIndex\",\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"},{\"name\":\"_account\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"components\":[{\"name\":\"pair_address\",\"type\":\"address\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\"},{\"name\":\"stable\",\"type\":\"bool\"},{\"name\":\"total_supply\",\"type\":\"uint256\"},{\"name\":\"token0\",\"type\":\"address\"},{\"name\":\"reserve0\",\"type\":\"uint256\"},{\"name\":\"claimable0\",\"type\":\"uint256\"},{\"name\":\"token1\",\"type\":\"address\"},{\"name\":\"reserve1\",\"type\":\"uint256\"},{\"name\":\"claimable1\",\"type\":\"uint256\"},{\"name\":\"gauge\",\"type\":\"address\"},{\"name\":\"gauge_total_supply\",\"type\":\"uint256\"},{\"name\":\"gauge_alive\",\"type\":\"bool\"},{\"name\":\"fee\",\"type\":\"address\"},{\"name\":\"bribe\",\"type\":\"address\"},{\"name\":\"wrapped_bribe\",\"type\":\"address\"},{\"name\":\"emissions\",\"type\":\"uint256\"},{\"name\":\"emissions_token\",\"type\":\"address\"},{\"name\":\"account_balance\",\"type\":\"uint256\"},{\"name\":\"account_earned\",\"type\":\"uint256\"},{\"name\":\"account_staked\",\"type\":\"uint256\"}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"byAddress\",\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_account\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"components\":[{\"name\":\"pair_address\",\"type\":\"address\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\"},{\"name\":\"stable\",\"type\":\"bool\"},{\"name\":\"total_supply\",\"type\":\"uint256\"},{\"name\":\"token0\",\"type\":\"address\"},{\"name\":\"reserve0\",\"type\":\"uint256\"},{\"name\":\"claimable0\",\"type\":\"uint256\"},{\"name\":\"token1\",\"type\":\"address\"},{\"name\":\"reserve1\",\"type\":\"uint256\"},{\"name\":\"claimable1\",\"type\":\"uint256\"},{\"name\":\"gauge\",\"type\":\"address\"},{\"name\":\"gauge_total_supply\",\"type\":\"uint256\"},{\"name\":\"gauge_alive\",\"type\":\"bool\"},{\"name\":\"fee\",\"type\":\"address\"},{\"name\":\"bribe\",\"type\":\"address\"},{\"name\":\"wrapped_bribe\",\"type\":\"address\"},{\"name\":\"emissions\",\"type\":\"uint256\"},{\"name\":\"emissions_token\",\"type\":\"address\"},{\"name\":\"account_balance\",\"type\":\"uint256\"},{\"name\":\"account_earned\",\"type\":\"uint256\"},{\"name\":\"account_staked\",\"type\":\"uint256\"}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"epochsLatest\",\"inputs\":[{\"name\":\"_limit\",\"type\":\"uint256\"},{\"name\":\"_offset\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"ts\",\"type\":\"uint256\"},{\"name\":\"pair_address\",\"type\":\"address\"},{\"name\":\"votes\",\"type\":\"uint256\"},{\"name\":\"emissions\",\"type\":\"uint256\"},{\"name\":\"bribes\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}]},{\"name\":\"fees\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}]}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"epochsByAddress\",\"inputs\":[{\"name\":\"_limit\",\"type\":\"uint256\"},{\"name\":\"_offset\",\"type\":\"uint256\"},{\"name\":\"_address\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"ts\",\"type\":\"uint256\"},{\"name\":\"pair_address\",\"type\":\"address\"},{\"name\":\"votes\",\"type\":\"uint256\"},{\"name\":\"emissions\",\"type\":\"uint256\"},{\"name\":\"bribes\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}]},{\"name\":\"fees\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}]}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"pair_factory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"voter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"wrapped_bribe_factory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]}]",
}

// VeloSugarABI is the input ABI used to generate the binding from.
// Deprecated: Use VeloSugarMetaData.ABI instead.
var VeloSugarABI = VeloSugarMetaData.ABI

// VeloSugar is an auto generated Go binding around an Ethereum contract.
type VeloSugar struct {
	VeloSugarCaller     // Read-only binding to the contract
	VeloSugarTransactor // Write-only binding to the contract
	VeloSugarFilterer   // Log filterer for contract events
}

// VeloSugarCaller is an auto generated read-only Go binding around an Ethereum contract.
type VeloSugarCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VeloSugarTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VeloSugarTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VeloSugarFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VeloSugarFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VeloSugarSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VeloSugarSession struct {
	Contract     *VeloSugar        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VeloSugarCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VeloSugarCallerSession struct {
	Contract *VeloSugarCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// VeloSugarTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VeloSugarTransactorSession struct {
	Contract     *VeloSugarTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// VeloSugarRaw is an auto generated low-level Go binding around an Ethereum contract.
type VeloSugarRaw struct {
	Contract *VeloSugar // Generic contract binding to access the raw methods on
}

// VeloSugarCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VeloSugarCallerRaw struct {
	Contract *VeloSugarCaller // Generic read-only contract binding to access the raw methods on
}

// VeloSugarTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VeloSugarTransactorRaw struct {
	Contract *VeloSugarTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVeloSugar creates a new instance of VeloSugar, bound to a specific deployed contract.
func NewVeloSugar(address common.Address, backend bind.ContractBackend) (*VeloSugar, error) {
	contract, err := bindVeloSugar(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VeloSugar{VeloSugarCaller: VeloSugarCaller{contract: contract}, VeloSugarTransactor: VeloSugarTransactor{contract: contract}, VeloSugarFilterer: VeloSugarFilterer{contract: contract}}, nil
}

// NewVeloSugarCaller creates a new read-only instance of VeloSugar, bound to a specific deployed contract.
func NewVeloSugarCaller(address common.Address, caller bind.ContractCaller) (*VeloSugarCaller, error) {
	contract, err := bindVeloSugar(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VeloSugarCaller{contract: contract}, nil
}

// NewVeloSugarTransactor creates a new write-only instance of VeloSugar, bound to a specific deployed contract.
func NewVeloSugarTransactor(address common.Address, transactor bind.ContractTransactor) (*VeloSugarTransactor, error) {
	contract, err := bindVeloSugar(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VeloSugarTransactor{contract: contract}, nil
}

// NewVeloSugarFilterer creates a new log filterer instance of VeloSugar, bound to a specific deployed contract.
func NewVeloSugarFilterer(address common.Address, filterer bind.ContractFilterer) (*VeloSugarFilterer, error) {
	contract, err := bindVeloSugar(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VeloSugarFilterer{contract: contract}, nil
}

// bindVeloSugar binds a generic wrapper to an already deployed contract.
func bindVeloSugar(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VeloSugarMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VeloSugar *VeloSugarRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VeloSugar.Contract.VeloSugarCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VeloSugar *VeloSugarRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VeloSugar.Contract.VeloSugarTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VeloSugar *VeloSugarRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VeloSugar.Contract.VeloSugarTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VeloSugar *VeloSugarCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VeloSugar.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VeloSugar *VeloSugarTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VeloSugar.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VeloSugar *VeloSugarTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VeloSugar.Contract.contract.Transact(opts, method, params...)
}

// All is a free data retrieval call binding the contract method 0xc0d0bf32.
//
// Solidity: function all(uint256 _limit, uint256 _offset, address _account) view returns((address,string,uint8,bool,uint256,address,uint256,uint256,address,uint256,uint256,address,uint256,bool,address,address,address,uint256,address,uint256,uint256,uint256)[])
func (_VeloSugar *VeloSugarCaller) All(opts *bind.CallOpts, _limit *big.Int, _offset *big.Int, _account common.Address) ([]Struct2, error) {
	var out []interface{}
	err := _VeloSugar.contract.Call(opts, &out, "all", _limit, _offset, _account)

	if err != nil {
		return *new([]Struct2), err
	}

	out0 := *abi.ConvertType(out[0], new([]Struct2)).(*[]Struct2)

	return out0, err

}

// All is a free data retrieval call binding the contract method 0xc0d0bf32.
//
// Solidity: function all(uint256 _limit, uint256 _offset, address _account) view returns((address,string,uint8,bool,uint256,address,uint256,uint256,address,uint256,uint256,address,uint256,bool,address,address,address,uint256,address,uint256,uint256,uint256)[])
func (_VeloSugar *VeloSugarSession) All(_limit *big.Int, _offset *big.Int, _account common.Address) ([]Struct2, error) {
	return _VeloSugar.Contract.All(&_VeloSugar.CallOpts, _limit, _offset, _account)
}

// All is a free data retrieval call binding the contract method 0xc0d0bf32.
//
// Solidity: function all(uint256 _limit, uint256 _offset, address _account) view returns((address,string,uint8,bool,uint256,address,uint256,uint256,address,uint256,uint256,address,uint256,bool,address,address,address,uint256,address,uint256,uint256,uint256)[])
func (_VeloSugar *VeloSugarCallerSession) All(_limit *big.Int, _offset *big.Int, _account common.Address) ([]Struct2, error) {
	return _VeloSugar.Contract.All(&_VeloSugar.CallOpts, _limit, _offset, _account)
}

// ByAddress is a free data retrieval call binding the contract method 0x6e917f2e.
//
// Solidity: function byAddress(address _address, address _account) view returns((address,string,uint8,bool,uint256,address,uint256,uint256,address,uint256,uint256,address,uint256,bool,address,address,address,uint256,address,uint256,uint256,uint256))
func (_VeloSugar *VeloSugarCaller) ByAddress(opts *bind.CallOpts, _address common.Address, _account common.Address) (Struct2, error) {
	var out []interface{}
	err := _VeloSugar.contract.Call(opts, &out, "byAddress", _address, _account)

	if err != nil {
		return *new(Struct2), err
	}

	out0 := *abi.ConvertType(out[0], new(Struct2)).(*Struct2)

	return out0, err

}

// ByAddress is a free data retrieval call binding the contract method 0x6e917f2e.
//
// Solidity: function byAddress(address _address, address _account) view returns((address,string,uint8,bool,uint256,address,uint256,uint256,address,uint256,uint256,address,uint256,bool,address,address,address,uint256,address,uint256,uint256,uint256))
func (_VeloSugar *VeloSugarSession) ByAddress(_address common.Address, _account common.Address) (Struct2, error) {
	return _VeloSugar.Contract.ByAddress(&_VeloSugar.CallOpts, _address, _account)
}

// ByAddress is a free data retrieval call binding the contract method 0x6e917f2e.
//
// Solidity: function byAddress(address _address, address _account) view returns((address,string,uint8,bool,uint256,address,uint256,uint256,address,uint256,uint256,address,uint256,bool,address,address,address,uint256,address,uint256,uint256,uint256))
func (_VeloSugar *VeloSugarCallerSession) ByAddress(_address common.Address, _account common.Address) (Struct2, error) {
	return _VeloSugar.Contract.ByAddress(&_VeloSugar.CallOpts, _address, _account)
}

// ByIndex is a free data retrieval call binding the contract method 0xfbb70183.
//
// Solidity: function byIndex(uint256 _index, address _account) view returns((address,string,uint8,bool,uint256,address,uint256,uint256,address,uint256,uint256,address,uint256,bool,address,address,address,uint256,address,uint256,uint256,uint256))
func (_VeloSugar *VeloSugarCaller) ByIndex(opts *bind.CallOpts, _index *big.Int, _account common.Address) (Struct2, error) {
	var out []interface{}
	err := _VeloSugar.contract.Call(opts, &out, "byIndex", _index, _account)

	if err != nil {
		return *new(Struct2), err
	}

	out0 := *abi.ConvertType(out[0], new(Struct2)).(*Struct2)

	return out0, err

}

// ByIndex is a free data retrieval call binding the contract method 0xfbb70183.
//
// Solidity: function byIndex(uint256 _index, address _account) view returns((address,string,uint8,bool,uint256,address,uint256,uint256,address,uint256,uint256,address,uint256,bool,address,address,address,uint256,address,uint256,uint256,uint256))
func (_VeloSugar *VeloSugarSession) ByIndex(_index *big.Int, _account common.Address) (Struct2, error) {
	return _VeloSugar.Contract.ByIndex(&_VeloSugar.CallOpts, _index, _account)
}

// ByIndex is a free data retrieval call binding the contract method 0xfbb70183.
//
// Solidity: function byIndex(uint256 _index, address _account) view returns((address,string,uint8,bool,uint256,address,uint256,uint256,address,uint256,uint256,address,uint256,bool,address,address,address,uint256,address,uint256,uint256,uint256))
func (_VeloSugar *VeloSugarCallerSession) ByIndex(_index *big.Int, _account common.Address) (Struct2, error) {
	return _VeloSugar.Contract.ByIndex(&_VeloSugar.CallOpts, _index, _account)
}

// EpochsByAddress is a free data retrieval call binding the contract method 0x8878d06c.
//
// Solidity: function epochsByAddress(uint256 _limit, uint256 _offset, address _address) view returns((uint256,address,uint256,uint256,(address,uint256)[],(address,uint256)[])[])
func (_VeloSugar *VeloSugarCaller) EpochsByAddress(opts *bind.CallOpts, _limit *big.Int, _offset *big.Int, _address common.Address) ([]Struct1, error) {
	var out []interface{}
	err := _VeloSugar.contract.Call(opts, &out, "epochsByAddress", _limit, _offset, _address)

	if err != nil {
		return *new([]Struct1), err
	}

	out0 := *abi.ConvertType(out[0], new([]Struct1)).(*[]Struct1)

	return out0, err

}

// EpochsByAddress is a free data retrieval call binding the contract method 0x8878d06c.
//
// Solidity: function epochsByAddress(uint256 _limit, uint256 _offset, address _address) view returns((uint256,address,uint256,uint256,(address,uint256)[],(address,uint256)[])[])
func (_VeloSugar *VeloSugarSession) EpochsByAddress(_limit *big.Int, _offset *big.Int, _address common.Address) ([]Struct1, error) {
	return _VeloSugar.Contract.EpochsByAddress(&_VeloSugar.CallOpts, _limit, _offset, _address)
}

// EpochsByAddress is a free data retrieval call binding the contract method 0x8878d06c.
//
// Solidity: function epochsByAddress(uint256 _limit, uint256 _offset, address _address) view returns((uint256,address,uint256,uint256,(address,uint256)[],(address,uint256)[])[])
func (_VeloSugar *VeloSugarCallerSession) EpochsByAddress(_limit *big.Int, _offset *big.Int, _address common.Address) ([]Struct1, error) {
	return _VeloSugar.Contract.EpochsByAddress(&_VeloSugar.CallOpts, _limit, _offset, _address)
}

// EpochsLatest is a free data retrieval call binding the contract method 0xd94b9bc6.
//
// Solidity: function epochsLatest(uint256 _limit, uint256 _offset) view returns((uint256,address,uint256,uint256,(address,uint256)[],(address,uint256)[])[])
func (_VeloSugar *VeloSugarCaller) EpochsLatest(opts *bind.CallOpts, _limit *big.Int, _offset *big.Int) ([]Struct1, error) {
	var out []interface{}
	err := _VeloSugar.contract.Call(opts, &out, "epochsLatest", _limit, _offset)

	if err != nil {
		return *new([]Struct1), err
	}

	out0 := *abi.ConvertType(out[0], new([]Struct1)).(*[]Struct1)

	return out0, err

}

// EpochsLatest is a free data retrieval call binding the contract method 0xd94b9bc6.
//
// Solidity: function epochsLatest(uint256 _limit, uint256 _offset) view returns((uint256,address,uint256,uint256,(address,uint256)[],(address,uint256)[])[])
func (_VeloSugar *VeloSugarSession) EpochsLatest(_limit *big.Int, _offset *big.Int) ([]Struct1, error) {
	return _VeloSugar.Contract.EpochsLatest(&_VeloSugar.CallOpts, _limit, _offset)
}

// EpochsLatest is a free data retrieval call binding the contract method 0xd94b9bc6.
//
// Solidity: function epochsLatest(uint256 _limit, uint256 _offset) view returns((uint256,address,uint256,uint256,(address,uint256)[],(address,uint256)[])[])
func (_VeloSugar *VeloSugarCallerSession) EpochsLatest(_limit *big.Int, _offset *big.Int) ([]Struct1, error) {
	return _VeloSugar.Contract.EpochsLatest(&_VeloSugar.CallOpts, _limit, _offset)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VeloSugar *VeloSugarCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VeloSugar.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VeloSugar *VeloSugarSession) Owner() (common.Address, error) {
	return _VeloSugar.Contract.Owner(&_VeloSugar.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VeloSugar *VeloSugarCallerSession) Owner() (common.Address, error) {
	return _VeloSugar.Contract.Owner(&_VeloSugar.CallOpts)
}

// PairFactory is a free data retrieval call binding the contract method 0xc606504f.
//
// Solidity: function pair_factory() view returns(address)
func (_VeloSugar *VeloSugarCaller) PairFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VeloSugar.contract.Call(opts, &out, "pair_factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PairFactory is a free data retrieval call binding the contract method 0xc606504f.
//
// Solidity: function pair_factory() view returns(address)
func (_VeloSugar *VeloSugarSession) PairFactory() (common.Address, error) {
	return _VeloSugar.Contract.PairFactory(&_VeloSugar.CallOpts)
}

// PairFactory is a free data retrieval call binding the contract method 0xc606504f.
//
// Solidity: function pair_factory() view returns(address)
func (_VeloSugar *VeloSugarCallerSession) PairFactory() (common.Address, error) {
	return _VeloSugar.Contract.PairFactory(&_VeloSugar.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_VeloSugar *VeloSugarCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VeloSugar.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_VeloSugar *VeloSugarSession) Token() (common.Address, error) {
	return _VeloSugar.Contract.Token(&_VeloSugar.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_VeloSugar *VeloSugarCallerSession) Token() (common.Address, error) {
	return _VeloSugar.Contract.Token(&_VeloSugar.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0x5cc33187.
//
// Solidity: function tokens(uint256 _limit, uint256 _offset, address _account) view returns((address,string,uint8,uint256,bool)[])
func (_VeloSugar *VeloSugarCaller) Tokens(opts *bind.CallOpts, _limit *big.Int, _offset *big.Int, _account common.Address) ([]Struct3, error) {
	var out []interface{}
	err := _VeloSugar.contract.Call(opts, &out, "tokens", _limit, _offset, _account)

	if err != nil {
		return *new([]Struct3), err
	}

	out0 := *abi.ConvertType(out[0], new([]Struct3)).(*[]Struct3)

	return out0, err

}

// Tokens is a free data retrieval call binding the contract method 0x5cc33187.
//
// Solidity: function tokens(uint256 _limit, uint256 _offset, address _account) view returns((address,string,uint8,uint256,bool)[])
func (_VeloSugar *VeloSugarSession) Tokens(_limit *big.Int, _offset *big.Int, _account common.Address) ([]Struct3, error) {
	return _VeloSugar.Contract.Tokens(&_VeloSugar.CallOpts, _limit, _offset, _account)
}

// Tokens is a free data retrieval call binding the contract method 0x5cc33187.
//
// Solidity: function tokens(uint256 _limit, uint256 _offset, address _account) view returns((address,string,uint8,uint256,bool)[])
func (_VeloSugar *VeloSugarCallerSession) Tokens(_limit *big.Int, _offset *big.Int, _account common.Address) ([]Struct3, error) {
	return _VeloSugar.Contract.Tokens(&_VeloSugar.CallOpts, _limit, _offset, _account)
}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_VeloSugar *VeloSugarCaller) Voter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VeloSugar.contract.Call(opts, &out, "voter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_VeloSugar *VeloSugarSession) Voter() (common.Address, error) {
	return _VeloSugar.Contract.Voter(&_VeloSugar.CallOpts)
}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_VeloSugar *VeloSugarCallerSession) Voter() (common.Address, error) {
	return _VeloSugar.Contract.Voter(&_VeloSugar.CallOpts)
}

// WrappedBribeFactory is a free data retrieval call binding the contract method 0x5aac4ed3.
//
// Solidity: function wrapped_bribe_factory() view returns(address)
func (_VeloSugar *VeloSugarCaller) WrappedBribeFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VeloSugar.contract.Call(opts, &out, "wrapped_bribe_factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WrappedBribeFactory is a free data retrieval call binding the contract method 0x5aac4ed3.
//
// Solidity: function wrapped_bribe_factory() view returns(address)
func (_VeloSugar *VeloSugarSession) WrappedBribeFactory() (common.Address, error) {
	return _VeloSugar.Contract.WrappedBribeFactory(&_VeloSugar.CallOpts)
}

// WrappedBribeFactory is a free data retrieval call binding the contract method 0x5aac4ed3.
//
// Solidity: function wrapped_bribe_factory() view returns(address)
func (_VeloSugar *VeloSugarCallerSession) WrappedBribeFactory() (common.Address, error) {
	return _VeloSugar.Contract.WrappedBribeFactory(&_VeloSugar.CallOpts)
}

// Setup is a paid mutator transaction binding the contract method 0x2d34ba79.
//
// Solidity: function setup(address _voter, address _wrapped_bribe_factory) returns()
func (_VeloSugar *VeloSugarTransactor) Setup(opts *bind.TransactOpts, _voter common.Address, _wrapped_bribe_factory common.Address) (*types.Transaction, error) {
	return _VeloSugar.contract.Transact(opts, "setup", _voter, _wrapped_bribe_factory)
}

// Setup is a paid mutator transaction binding the contract method 0x2d34ba79.
//
// Solidity: function setup(address _voter, address _wrapped_bribe_factory) returns()
func (_VeloSugar *VeloSugarSession) Setup(_voter common.Address, _wrapped_bribe_factory common.Address) (*types.Transaction, error) {
	return _VeloSugar.Contract.Setup(&_VeloSugar.TransactOpts, _voter, _wrapped_bribe_factory)
}

// Setup is a paid mutator transaction binding the contract method 0x2d34ba79.
//
// Solidity: function setup(address _voter, address _wrapped_bribe_factory) returns()
func (_VeloSugar *VeloSugarTransactorSession) Setup(_voter common.Address, _wrapped_bribe_factory common.Address) (*types.Transaction, error) {
	return _VeloSugar.Contract.Setup(&_VeloSugar.TransactOpts, _voter, _wrapped_bribe_factory)
}
