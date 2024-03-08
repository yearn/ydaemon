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

// Struct4 is an auto generated low-level Go binding around an user-defined struct.
type Struct4 struct {
	Lp      common.Address
	Type    *big.Int
	Token0  common.Address
	Token1  common.Address
	Factory common.Address
	PoolFee *big.Int
}

// VeloSugarStruct0 is an auto generated low-level Go binding around an user-defined struct.
type VeloSugarStruct0 struct {
	TokenAddress   common.Address
	Symbol         string
	Decimals       uint8
	AccountBalance *big.Int
	Listed         bool
}

// Struct5 is an auto generated low-level Go binding around an user-defined struct.
type Struct5 struct {
	Lp             common.Address
	Symbol         string
	Decimals       uint8
	Liquidity      *big.Int
	Type           *big.Int
	Tick           *big.Int
	SqrtRatio      *big.Int
	Token0         common.Address
	Reserve0       *big.Int
	Staked0        *big.Int
	Token1         common.Address
	Reserve1       *big.Int
	Staked1        *big.Int
	Gauge          common.Address
	GaugeLiquidity *big.Int
	GaugeAlive     bool
	Fee            common.Address
	Bribe          common.Address
	Factory        common.Address
	Emissions      *big.Int
	EmissionsToken common.Address
	PoolFee        *big.Int
	UnstakedFee    *big.Int
	Token0Fees     *big.Int
	Token1Fees     *big.Int
}

// Struct1 is an auto generated low-level Go binding around an user-defined struct.
type Struct1 struct {
	Token  common.Address
	Amount *big.Int
}

// Struct3 is an auto generated low-level Go binding around an user-defined struct.
type Struct3 struct {
	VenftId *big.Int
	Lp      common.Address
	Amount  *big.Int
	Token   common.Address
	Fee     common.Address
	Bribe   common.Address
}

// Struct2 is an auto generated low-level Go binding around an user-defined struct.
type Struct2 struct {
	Ts        *big.Int
	Lp        common.Address
	Votes     *big.Int
	Emissions *big.Int
	Bribes    []Struct1
	Fees      []Struct1
}

// Struct6 is an auto generated low-level Go binding around an user-defined struct.
type Struct6 struct {
	Id              *big.Int
	Lp              common.Address
	Liquidity       *big.Int
	Staked          *big.Int
	Amount0         *big.Int
	Amount1         *big.Int
	Staked0         *big.Int
	Staked1         *big.Int
	UnstakedEarned0 *big.Int
	UnstakedEarned1 *big.Int
	EmissionsEarned *big.Int
	TickLower       *big.Int
	TickUpper       *big.Int
	SqrtRatioLower  *big.Int
	SqrtRatioUpper  *big.Int
}

// VeloSugarMetaData contains all meta data concerning the VeloSugar contract.
var VeloSugarMetaData = &bind.MetaData{
	ABI: "[{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"_voter\",\"type\":\"address\"},{\"name\":\"_registry\",\"type\":\"address\"},{\"name\":\"_convertor\",\"type\":\"address\"},{\"name\":\"_nfpm\",\"type\":\"address\"},{\"name\":\"_slipstream_helper\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"forSwaps\",\"inputs\":[{\"name\":\"_limit\",\"type\":\"uint256\"},{\"name\":\"_offset\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"lp\",\"type\":\"address\"},{\"name\":\"type\",\"type\":\"int24\"},{\"name\":\"token0\",\"type\":\"address\"},{\"name\":\"token1\",\"type\":\"address\"},{\"name\":\"factory\",\"type\":\"address\"},{\"name\":\"pool_fee\",\"type\":\"uint256\"}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"tokens\",\"inputs\":[{\"name\":\"_limit\",\"type\":\"uint256\"},{\"name\":\"_offset\",\"type\":\"uint256\"},{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_addresses\",\"type\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"token_address\",\"type\":\"address\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\"},{\"name\":\"account_balance\",\"type\":\"uint256\"},{\"name\":\"listed\",\"type\":\"bool\"}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"all\",\"inputs\":[{\"name\":\"_limit\",\"type\":\"uint256\"},{\"name\":\"_offset\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"lp\",\"type\":\"address\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\"},{\"name\":\"liquidity\",\"type\":\"uint256\"},{\"name\":\"type\",\"type\":\"int24\"},{\"name\":\"tick\",\"type\":\"int24\"},{\"name\":\"sqrt_ratio\",\"type\":\"uint160\"},{\"name\":\"token0\",\"type\":\"address\"},{\"name\":\"reserve0\",\"type\":\"uint256\"},{\"name\":\"staked0\",\"type\":\"uint256\"},{\"name\":\"token1\",\"type\":\"address\"},{\"name\":\"reserve1\",\"type\":\"uint256\"},{\"name\":\"staked1\",\"type\":\"uint256\"},{\"name\":\"gauge\",\"type\":\"address\"},{\"name\":\"gauge_liquidity\",\"type\":\"uint256\"},{\"name\":\"gauge_alive\",\"type\":\"bool\"},{\"name\":\"fee\",\"type\":\"address\"},{\"name\":\"bribe\",\"type\":\"address\"},{\"name\":\"factory\",\"type\":\"address\"},{\"name\":\"emissions\",\"type\":\"uint256\"},{\"name\":\"emissions_token\",\"type\":\"address\"},{\"name\":\"pool_fee\",\"type\":\"uint256\"},{\"name\":\"unstaked_fee\",\"type\":\"uint256\"},{\"name\":\"token0_fees\",\"type\":\"uint256\"},{\"name\":\"token1_fees\",\"type\":\"uint256\"}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"byIndex\",\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"components\":[{\"name\":\"lp\",\"type\":\"address\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\"},{\"name\":\"liquidity\",\"type\":\"uint256\"},{\"name\":\"type\",\"type\":\"int24\"},{\"name\":\"tick\",\"type\":\"int24\"},{\"name\":\"sqrt_ratio\",\"type\":\"uint160\"},{\"name\":\"token0\",\"type\":\"address\"},{\"name\":\"reserve0\",\"type\":\"uint256\"},{\"name\":\"staked0\",\"type\":\"uint256\"},{\"name\":\"token1\",\"type\":\"address\"},{\"name\":\"reserve1\",\"type\":\"uint256\"},{\"name\":\"staked1\",\"type\":\"uint256\"},{\"name\":\"gauge\",\"type\":\"address\"},{\"name\":\"gauge_liquidity\",\"type\":\"uint256\"},{\"name\":\"gauge_alive\",\"type\":\"bool\"},{\"name\":\"fee\",\"type\":\"address\"},{\"name\":\"bribe\",\"type\":\"address\"},{\"name\":\"factory\",\"type\":\"address\"},{\"name\":\"emissions\",\"type\":\"uint256\"},{\"name\":\"emissions_token\",\"type\":\"address\"},{\"name\":\"pool_fee\",\"type\":\"uint256\"},{\"name\":\"unstaked_fee\",\"type\":\"uint256\"},{\"name\":\"token0_fees\",\"type\":\"uint256\"},{\"name\":\"token1_fees\",\"type\":\"uint256\"}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"positions\",\"inputs\":[{\"name\":\"_limit\",\"type\":\"uint256\"},{\"name\":\"_offset\",\"type\":\"uint256\"},{\"name\":\"_account\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"lp\",\"type\":\"address\"},{\"name\":\"liquidity\",\"type\":\"uint256\"},{\"name\":\"staked\",\"type\":\"uint256\"},{\"name\":\"amount0\",\"type\":\"uint256\"},{\"name\":\"amount1\",\"type\":\"uint256\"},{\"name\":\"staked0\",\"type\":\"uint256\"},{\"name\":\"staked1\",\"type\":\"uint256\"},{\"name\":\"unstaked_earned0\",\"type\":\"uint256\"},{\"name\":\"unstaked_earned1\",\"type\":\"uint256\"},{\"name\":\"emissions_earned\",\"type\":\"uint256\"},{\"name\":\"tick_lower\",\"type\":\"int24\"},{\"name\":\"tick_upper\",\"type\":\"int24\"},{\"name\":\"sqrt_ratio_lower\",\"type\":\"uint160\"},{\"name\":\"sqrt_ratio_upper\",\"type\":\"uint160\"}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"epochsLatest\",\"inputs\":[{\"name\":\"_limit\",\"type\":\"uint256\"},{\"name\":\"_offset\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"ts\",\"type\":\"uint256\"},{\"name\":\"lp\",\"type\":\"address\"},{\"name\":\"votes\",\"type\":\"uint256\"},{\"name\":\"emissions\",\"type\":\"uint256\"},{\"name\":\"bribes\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}]},{\"name\":\"fees\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}]}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"epochsByAddress\",\"inputs\":[{\"name\":\"_limit\",\"type\":\"uint256\"},{\"name\":\"_offset\",\"type\":\"uint256\"},{\"name\":\"_address\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"ts\",\"type\":\"uint256\"},{\"name\":\"lp\",\"type\":\"address\"},{\"name\":\"votes\",\"type\":\"uint256\"},{\"name\":\"emissions\",\"type\":\"uint256\"},{\"name\":\"bribes\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}]},{\"name\":\"fees\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}]}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"rewards\",\"inputs\":[{\"name\":\"_limit\",\"type\":\"uint256\"},{\"name\":\"_offset\",\"type\":\"uint256\"},{\"name\":\"_venft_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"venft_id\",\"type\":\"uint256\"},{\"name\":\"lp\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"fee\",\"type\":\"address\"},{\"name\":\"bribe\",\"type\":\"address\"}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"rewardsByAddress\",\"inputs\":[{\"name\":\"_venft_id\",\"type\":\"uint256\"},{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"venft_id\",\"type\":\"uint256\"},{\"name\":\"lp\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"fee\",\"type\":\"address\"},{\"name\":\"bribe\",\"type\":\"address\"}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"registry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"voter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"convertor\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"nfpm\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"cl_helper\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]}]",
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

// All is a free data retrieval call binding the contract method 0xb10daf7b.
//
// Solidity: function all(uint256 _limit, uint256 _offset) view returns((address,string,uint8,uint256,int24,int24,uint160,address,uint256,uint256,address,uint256,uint256,address,uint256,bool,address,address,address,uint256,address,uint256,uint256,uint256,uint256)[])
func (_VeloSugar *VeloSugarCaller) All(opts *bind.CallOpts, _limit *big.Int, _offset *big.Int) ([]Struct5, error) {
	var out []interface{}
	err := _VeloSugar.contract.Call(opts, &out, "all", _limit, _offset)

	if err != nil {
		return *new([]Struct5), err
	}

	out0 := *abi.ConvertType(out[0], new([]Struct5)).(*[]Struct5)

	return out0, err

}

// All is a free data retrieval call binding the contract method 0xb10daf7b.
//
// Solidity: function all(uint256 _limit, uint256 _offset) view returns((address,string,uint8,uint256,int24,int24,uint160,address,uint256,uint256,address,uint256,uint256,address,uint256,bool,address,address,address,uint256,address,uint256,uint256,uint256,uint256)[])
func (_VeloSugar *VeloSugarSession) All(_limit *big.Int, _offset *big.Int) ([]Struct5, error) {
	return _VeloSugar.Contract.All(&_VeloSugar.CallOpts, _limit, _offset)
}

// All is a free data retrieval call binding the contract method 0xb10daf7b.
//
// Solidity: function all(uint256 _limit, uint256 _offset) view returns((address,string,uint8,uint256,int24,int24,uint160,address,uint256,uint256,address,uint256,uint256,address,uint256,bool,address,address,address,uint256,address,uint256,uint256,uint256,uint256)[])
func (_VeloSugar *VeloSugarCallerSession) All(_limit *big.Int, _offset *big.Int) ([]Struct5, error) {
	return _VeloSugar.Contract.All(&_VeloSugar.CallOpts, _limit, _offset)
}

// ByIndex is a free data retrieval call binding the contract method 0x1f342dd6.
//
// Solidity: function byIndex(uint256 _index) view returns((address,string,uint8,uint256,int24,int24,uint160,address,uint256,uint256,address,uint256,uint256,address,uint256,bool,address,address,address,uint256,address,uint256,uint256,uint256,uint256))
func (_VeloSugar *VeloSugarCaller) ByIndex(opts *bind.CallOpts, _index *big.Int) (Struct5, error) {
	var out []interface{}
	err := _VeloSugar.contract.Call(opts, &out, "byIndex", _index)

	if err != nil {
		return *new(Struct5), err
	}

	out0 := *abi.ConvertType(out[0], new(Struct5)).(*Struct5)

	return out0, err

}

// ByIndex is a free data retrieval call binding the contract method 0x1f342dd6.
//
// Solidity: function byIndex(uint256 _index) view returns((address,string,uint8,uint256,int24,int24,uint160,address,uint256,uint256,address,uint256,uint256,address,uint256,bool,address,address,address,uint256,address,uint256,uint256,uint256,uint256))
func (_VeloSugar *VeloSugarSession) ByIndex(_index *big.Int) (Struct5, error) {
	return _VeloSugar.Contract.ByIndex(&_VeloSugar.CallOpts, _index)
}

// ByIndex is a free data retrieval call binding the contract method 0x1f342dd6.
//
// Solidity: function byIndex(uint256 _index) view returns((address,string,uint8,uint256,int24,int24,uint160,address,uint256,uint256,address,uint256,uint256,address,uint256,bool,address,address,address,uint256,address,uint256,uint256,uint256,uint256))
func (_VeloSugar *VeloSugarCallerSession) ByIndex(_index *big.Int) (Struct5, error) {
	return _VeloSugar.Contract.ByIndex(&_VeloSugar.CallOpts, _index)
}

// ClHelper is a free data retrieval call binding the contract method 0xc954a389.
//
// Solidity: function cl_helper() view returns(address)
func (_VeloSugar *VeloSugarCaller) ClHelper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VeloSugar.contract.Call(opts, &out, "cl_helper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ClHelper is a free data retrieval call binding the contract method 0xc954a389.
//
// Solidity: function cl_helper() view returns(address)
func (_VeloSugar *VeloSugarSession) ClHelper() (common.Address, error) {
	return _VeloSugar.Contract.ClHelper(&_VeloSugar.CallOpts)
}

// ClHelper is a free data retrieval call binding the contract method 0xc954a389.
//
// Solidity: function cl_helper() view returns(address)
func (_VeloSugar *VeloSugarCallerSession) ClHelper() (common.Address, error) {
	return _VeloSugar.Contract.ClHelper(&_VeloSugar.CallOpts)
}

// Convertor is a free data retrieval call binding the contract method 0xb5030306.
//
// Solidity: function convertor() view returns(address)
func (_VeloSugar *VeloSugarCaller) Convertor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VeloSugar.contract.Call(opts, &out, "convertor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Convertor is a free data retrieval call binding the contract method 0xb5030306.
//
// Solidity: function convertor() view returns(address)
func (_VeloSugar *VeloSugarSession) Convertor() (common.Address, error) {
	return _VeloSugar.Contract.Convertor(&_VeloSugar.CallOpts)
}

// Convertor is a free data retrieval call binding the contract method 0xb5030306.
//
// Solidity: function convertor() view returns(address)
func (_VeloSugar *VeloSugarCallerSession) Convertor() (common.Address, error) {
	return _VeloSugar.Contract.Convertor(&_VeloSugar.CallOpts)
}

// EpochsByAddress is a free data retrieval call binding the contract method 0x8878d06c.
//
// Solidity: function epochsByAddress(uint256 _limit, uint256 _offset, address _address) view returns((uint256,address,uint256,uint256,(address,uint256)[],(address,uint256)[])[])
func (_VeloSugar *VeloSugarCaller) EpochsByAddress(opts *bind.CallOpts, _limit *big.Int, _offset *big.Int, _address common.Address) ([]Struct2, error) {
	var out []interface{}
	err := _VeloSugar.contract.Call(opts, &out, "epochsByAddress", _limit, _offset, _address)

	if err != nil {
		return *new([]Struct2), err
	}

	out0 := *abi.ConvertType(out[0], new([]Struct2)).(*[]Struct2)

	return out0, err

}

// EpochsByAddress is a free data retrieval call binding the contract method 0x8878d06c.
//
// Solidity: function epochsByAddress(uint256 _limit, uint256 _offset, address _address) view returns((uint256,address,uint256,uint256,(address,uint256)[],(address,uint256)[])[])
func (_VeloSugar *VeloSugarSession) EpochsByAddress(_limit *big.Int, _offset *big.Int, _address common.Address) ([]Struct2, error) {
	return _VeloSugar.Contract.EpochsByAddress(&_VeloSugar.CallOpts, _limit, _offset, _address)
}

// EpochsByAddress is a free data retrieval call binding the contract method 0x8878d06c.
//
// Solidity: function epochsByAddress(uint256 _limit, uint256 _offset, address _address) view returns((uint256,address,uint256,uint256,(address,uint256)[],(address,uint256)[])[])
func (_VeloSugar *VeloSugarCallerSession) EpochsByAddress(_limit *big.Int, _offset *big.Int, _address common.Address) ([]Struct2, error) {
	return _VeloSugar.Contract.EpochsByAddress(&_VeloSugar.CallOpts, _limit, _offset, _address)
}

// EpochsLatest is a free data retrieval call binding the contract method 0xd94b9bc6.
//
// Solidity: function epochsLatest(uint256 _limit, uint256 _offset) view returns((uint256,address,uint256,uint256,(address,uint256)[],(address,uint256)[])[])
func (_VeloSugar *VeloSugarCaller) EpochsLatest(opts *bind.CallOpts, _limit *big.Int, _offset *big.Int) ([]Struct2, error) {
	var out []interface{}
	err := _VeloSugar.contract.Call(opts, &out, "epochsLatest", _limit, _offset)

	if err != nil {
		return *new([]Struct2), err
	}

	out0 := *abi.ConvertType(out[0], new([]Struct2)).(*[]Struct2)

	return out0, err

}

// EpochsLatest is a free data retrieval call binding the contract method 0xd94b9bc6.
//
// Solidity: function epochsLatest(uint256 _limit, uint256 _offset) view returns((uint256,address,uint256,uint256,(address,uint256)[],(address,uint256)[])[])
func (_VeloSugar *VeloSugarSession) EpochsLatest(_limit *big.Int, _offset *big.Int) ([]Struct2, error) {
	return _VeloSugar.Contract.EpochsLatest(&_VeloSugar.CallOpts, _limit, _offset)
}

// EpochsLatest is a free data retrieval call binding the contract method 0xd94b9bc6.
//
// Solidity: function epochsLatest(uint256 _limit, uint256 _offset) view returns((uint256,address,uint256,uint256,(address,uint256)[],(address,uint256)[])[])
func (_VeloSugar *VeloSugarCallerSession) EpochsLatest(_limit *big.Int, _offset *big.Int) ([]Struct2, error) {
	return _VeloSugar.Contract.EpochsLatest(&_VeloSugar.CallOpts, _limit, _offset)
}

// ForSwaps is a free data retrieval call binding the contract method 0xb224fcb5.
//
// Solidity: function forSwaps(uint256 _limit, uint256 _offset) view returns((address,int24,address,address,address,uint256)[])
func (_VeloSugar *VeloSugarCaller) ForSwaps(opts *bind.CallOpts, _limit *big.Int, _offset *big.Int) ([]Struct4, error) {
	var out []interface{}
	err := _VeloSugar.contract.Call(opts, &out, "forSwaps", _limit, _offset)

	if err != nil {
		return *new([]Struct4), err
	}

	out0 := *abi.ConvertType(out[0], new([]Struct4)).(*[]Struct4)

	return out0, err

}

// ForSwaps is a free data retrieval call binding the contract method 0xb224fcb5.
//
// Solidity: function forSwaps(uint256 _limit, uint256 _offset) view returns((address,int24,address,address,address,uint256)[])
func (_VeloSugar *VeloSugarSession) ForSwaps(_limit *big.Int, _offset *big.Int) ([]Struct4, error) {
	return _VeloSugar.Contract.ForSwaps(&_VeloSugar.CallOpts, _limit, _offset)
}

// ForSwaps is a free data retrieval call binding the contract method 0xb224fcb5.
//
// Solidity: function forSwaps(uint256 _limit, uint256 _offset) view returns((address,int24,address,address,address,uint256)[])
func (_VeloSugar *VeloSugarCallerSession) ForSwaps(_limit *big.Int, _offset *big.Int) ([]Struct4, error) {
	return _VeloSugar.Contract.ForSwaps(&_VeloSugar.CallOpts, _limit, _offset)
}

// Nfpm is a free data retrieval call binding the contract method 0x7303e913.
//
// Solidity: function nfpm() view returns(address)
func (_VeloSugar *VeloSugarCaller) Nfpm(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VeloSugar.contract.Call(opts, &out, "nfpm")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Nfpm is a free data retrieval call binding the contract method 0x7303e913.
//
// Solidity: function nfpm() view returns(address)
func (_VeloSugar *VeloSugarSession) Nfpm() (common.Address, error) {
	return _VeloSugar.Contract.Nfpm(&_VeloSugar.CallOpts)
}

// Nfpm is a free data retrieval call binding the contract method 0x7303e913.
//
// Solidity: function nfpm() view returns(address)
func (_VeloSugar *VeloSugarCallerSession) Nfpm() (common.Address, error) {
	return _VeloSugar.Contract.Nfpm(&_VeloSugar.CallOpts)
}

// Positions is a free data retrieval call binding the contract method 0xedbd33bf.
//
// Solidity: function positions(uint256 _limit, uint256 _offset, address _account) view returns((uint256,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int24,int24,uint160,uint160)[])
func (_VeloSugar *VeloSugarCaller) Positions(opts *bind.CallOpts, _limit *big.Int, _offset *big.Int, _account common.Address) ([]Struct6, error) {
	var out []interface{}
	err := _VeloSugar.contract.Call(opts, &out, "positions", _limit, _offset, _account)

	if err != nil {
		return *new([]Struct6), err
	}

	out0 := *abi.ConvertType(out[0], new([]Struct6)).(*[]Struct6)

	return out0, err

}

// Positions is a free data retrieval call binding the contract method 0xedbd33bf.
//
// Solidity: function positions(uint256 _limit, uint256 _offset, address _account) view returns((uint256,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int24,int24,uint160,uint160)[])
func (_VeloSugar *VeloSugarSession) Positions(_limit *big.Int, _offset *big.Int, _account common.Address) ([]Struct6, error) {
	return _VeloSugar.Contract.Positions(&_VeloSugar.CallOpts, _limit, _offset, _account)
}

// Positions is a free data retrieval call binding the contract method 0xedbd33bf.
//
// Solidity: function positions(uint256 _limit, uint256 _offset, address _account) view returns((uint256,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int24,int24,uint160,uint160)[])
func (_VeloSugar *VeloSugarCallerSession) Positions(_limit *big.Int, _offset *big.Int, _account common.Address) ([]Struct6, error) {
	return _VeloSugar.Contract.Positions(&_VeloSugar.CallOpts, _limit, _offset, _account)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_VeloSugar *VeloSugarCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VeloSugar.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_VeloSugar *VeloSugarSession) Registry() (common.Address, error) {
	return _VeloSugar.Contract.Registry(&_VeloSugar.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_VeloSugar *VeloSugarCallerSession) Registry() (common.Address, error) {
	return _VeloSugar.Contract.Registry(&_VeloSugar.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0xa9c57fee.
//
// Solidity: function rewards(uint256 _limit, uint256 _offset, uint256 _venft_id) view returns((uint256,address,uint256,address,address,address)[])
func (_VeloSugar *VeloSugarCaller) Rewards(opts *bind.CallOpts, _limit *big.Int, _offset *big.Int, _venft_id *big.Int) ([]Struct3, error) {
	var out []interface{}
	err := _VeloSugar.contract.Call(opts, &out, "rewards", _limit, _offset, _venft_id)

	if err != nil {
		return *new([]Struct3), err
	}

	out0 := *abi.ConvertType(out[0], new([]Struct3)).(*[]Struct3)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0xa9c57fee.
//
// Solidity: function rewards(uint256 _limit, uint256 _offset, uint256 _venft_id) view returns((uint256,address,uint256,address,address,address)[])
func (_VeloSugar *VeloSugarSession) Rewards(_limit *big.Int, _offset *big.Int, _venft_id *big.Int) ([]Struct3, error) {
	return _VeloSugar.Contract.Rewards(&_VeloSugar.CallOpts, _limit, _offset, _venft_id)
}

// Rewards is a free data retrieval call binding the contract method 0xa9c57fee.
//
// Solidity: function rewards(uint256 _limit, uint256 _offset, uint256 _venft_id) view returns((uint256,address,uint256,address,address,address)[])
func (_VeloSugar *VeloSugarCallerSession) Rewards(_limit *big.Int, _offset *big.Int, _venft_id *big.Int) ([]Struct3, error) {
	return _VeloSugar.Contract.Rewards(&_VeloSugar.CallOpts, _limit, _offset, _venft_id)
}

// RewardsByAddress is a free data retrieval call binding the contract method 0xcd824fb4.
//
// Solidity: function rewardsByAddress(uint256 _venft_id, address _pool) view returns((uint256,address,uint256,address,address,address)[])
func (_VeloSugar *VeloSugarCaller) RewardsByAddress(opts *bind.CallOpts, _venft_id *big.Int, _pool common.Address) ([]Struct3, error) {
	var out []interface{}
	err := _VeloSugar.contract.Call(opts, &out, "rewardsByAddress", _venft_id, _pool)

	if err != nil {
		return *new([]Struct3), err
	}

	out0 := *abi.ConvertType(out[0], new([]Struct3)).(*[]Struct3)

	return out0, err

}

// RewardsByAddress is a free data retrieval call binding the contract method 0xcd824fb4.
//
// Solidity: function rewardsByAddress(uint256 _venft_id, address _pool) view returns((uint256,address,uint256,address,address,address)[])
func (_VeloSugar *VeloSugarSession) RewardsByAddress(_venft_id *big.Int, _pool common.Address) ([]Struct3, error) {
	return _VeloSugar.Contract.RewardsByAddress(&_VeloSugar.CallOpts, _venft_id, _pool)
}

// RewardsByAddress is a free data retrieval call binding the contract method 0xcd824fb4.
//
// Solidity: function rewardsByAddress(uint256 _venft_id, address _pool) view returns((uint256,address,uint256,address,address,address)[])
func (_VeloSugar *VeloSugarCallerSession) RewardsByAddress(_venft_id *big.Int, _pool common.Address) ([]Struct3, error) {
	return _VeloSugar.Contract.RewardsByAddress(&_VeloSugar.CallOpts, _venft_id, _pool)
}

// Tokens is a free data retrieval call binding the contract method 0x295212be.
//
// Solidity: function tokens(uint256 _limit, uint256 _offset, address _account, address[] _addresses) view returns((address,string,uint8,uint256,bool)[])
func (_VeloSugar *VeloSugarCaller) Tokens(opts *bind.CallOpts, _limit *big.Int, _offset *big.Int, _account common.Address, _addresses []common.Address) ([]VeloSugarStruct0, error) {
	var out []interface{}
	err := _VeloSugar.contract.Call(opts, &out, "tokens", _limit, _offset, _account, _addresses)

	if err != nil {
		return *new([]VeloSugarStruct0), err
	}

	out0 := *abi.ConvertType(out[0], new([]VeloSugarStruct0)).(*[]VeloSugarStruct0)

	return out0, err

}

// Tokens is a free data retrieval call binding the contract method 0x295212be.
//
// Solidity: function tokens(uint256 _limit, uint256 _offset, address _account, address[] _addresses) view returns((address,string,uint8,uint256,bool)[])
func (_VeloSugar *VeloSugarSession) Tokens(_limit *big.Int, _offset *big.Int, _account common.Address, _addresses []common.Address) ([]VeloSugarStruct0, error) {
	return _VeloSugar.Contract.Tokens(&_VeloSugar.CallOpts, _limit, _offset, _account, _addresses)
}

// Tokens is a free data retrieval call binding the contract method 0x295212be.
//
// Solidity: function tokens(uint256 _limit, uint256 _offset, address _account, address[] _addresses) view returns((address,string,uint8,uint256,bool)[])
func (_VeloSugar *VeloSugarCallerSession) Tokens(_limit *big.Int, _offset *big.Int, _account common.Address, _addresses []common.Address) ([]VeloSugarStruct0, error) {
	return _VeloSugar.Contract.Tokens(&_VeloSugar.CallOpts, _limit, _offset, _account, _addresses)
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
