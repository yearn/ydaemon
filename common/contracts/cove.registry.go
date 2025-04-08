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

// CoveYearnGaugeFactoryGaugeInfo is an auto generated low-level Go binding around an user-defined struct.
type CoveYearnGaugeFactoryGaugeInfo struct {
	YearnVaultAsset         common.Address
	YearnVault              common.Address
	IsVaultV2               bool
	YearnGauge              common.Address
	CoveYearnStrategy       common.Address
	AutoCompoundingGauge    common.Address
	NonAutoCompoundingGauge common.Address
}

// CoveRegistryMetaData contains all meta data concerning the CoveRegistry contract.
var CoveRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"factoryAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ysd\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cove\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardForwarderImpl_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc20RewardsGaugeImpl_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ysdRewardsGaugeImpl_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gaugeAdmin_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gaugeManager_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gaugePauser_\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddressNotContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GaugeAlreadyDeployed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GaugeNotDeployed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"yearnGauge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"coveYearnStrategy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"autoCompoundingGauge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nonAutoCompoundingGauge\",\"type\":\"address\"}],\"name\":\"CoveGaugesDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"COVE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"YEARN_STAKING_DELEGATE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"coveYearnStrategy\",\"type\":\"address\"}],\"name\":\"deployCoveGauges\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"erc20RewardsGaugeImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gaugeAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gaugeManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gaugePauser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"getAllGaugeInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"yearnVaultAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"yearnVault\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isVaultV2\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"yearnGauge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"coveYearnStrategy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"autoCompoundingGauge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nonAutoCompoundingGauge\",\"type\":\"address\"}],\"internalType\":\"structCoveYearnGaugeFactory.GaugeInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"yearnGauge\",\"type\":\"address\"}],\"name\":\"getGaugeInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"yearnVaultAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"yearnVault\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isVaultV2\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"yearnGauge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"coveYearnStrategy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"autoCompoundingGauge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nonAutoCompoundingGauge\",\"type\":\"address\"}],\"internalType\":\"structCoveYearnGaugeFactory.GaugeInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numOfSupportedYearnGauges\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardForwarderImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"setERC20RewardsGaugeImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"setGaugeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"setGaugeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"setGaugePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"setRewardForwarderImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"setYsdRewardsGaugeImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"supportedYearnGauges\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"yearnGaugeInfoStored\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"coveYearnStrategy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"autoCompoundingGauge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nonAutoCompoundingGauge\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ysdRewardsGaugeImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CoveRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use CoveRegistryMetaData.ABI instead.
var CoveRegistryABI = CoveRegistryMetaData.ABI

// CoveRegistry is an auto generated Go binding around an Ethereum contract.
type CoveRegistry struct {
	CoveRegistryCaller     // Read-only binding to the contract
	CoveRegistryTransactor // Write-only binding to the contract
	CoveRegistryFilterer   // Log filterer for contract events
}

// CoveRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type CoveRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoveRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CoveRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoveRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CoveRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoveRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CoveRegistrySession struct {
	Contract     *CoveRegistry     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CoveRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CoveRegistryCallerSession struct {
	Contract *CoveRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CoveRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CoveRegistryTransactorSession struct {
	Contract     *CoveRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CoveRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type CoveRegistryRaw struct {
	Contract *CoveRegistry // Generic contract binding to access the raw methods on
}

// CoveRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CoveRegistryCallerRaw struct {
	Contract *CoveRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// CoveRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CoveRegistryTransactorRaw struct {
	Contract *CoveRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCoveRegistry creates a new instance of CoveRegistry, bound to a specific deployed contract.
func NewCoveRegistry(address common.Address, backend bind.ContractBackend) (*CoveRegistry, error) {
	contract, err := bindCoveRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CoveRegistry{CoveRegistryCaller: CoveRegistryCaller{contract: contract}, CoveRegistryTransactor: CoveRegistryTransactor{contract: contract}, CoveRegistryFilterer: CoveRegistryFilterer{contract: contract}}, nil
}

// NewCoveRegistryCaller creates a new read-only instance of CoveRegistry, bound to a specific deployed contract.
func NewCoveRegistryCaller(address common.Address, caller bind.ContractCaller) (*CoveRegistryCaller, error) {
	contract, err := bindCoveRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CoveRegistryCaller{contract: contract}, nil
}

// NewCoveRegistryTransactor creates a new write-only instance of CoveRegistry, bound to a specific deployed contract.
func NewCoveRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*CoveRegistryTransactor, error) {
	contract, err := bindCoveRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CoveRegistryTransactor{contract: contract}, nil
}

// NewCoveRegistryFilterer creates a new log filterer instance of CoveRegistry, bound to a specific deployed contract.
func NewCoveRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*CoveRegistryFilterer, error) {
	contract, err := bindCoveRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CoveRegistryFilterer{contract: contract}, nil
}

// bindCoveRegistry binds a generic wrapper to an already deployed contract.
func bindCoveRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CoveRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CoveRegistry *CoveRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CoveRegistry.Contract.CoveRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CoveRegistry *CoveRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoveRegistry.Contract.CoveRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CoveRegistry *CoveRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CoveRegistry.Contract.CoveRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CoveRegistry *CoveRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CoveRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CoveRegistry *CoveRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoveRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CoveRegistry *CoveRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CoveRegistry.Contract.contract.Transact(opts, method, params...)
}

// COVE is a free data retrieval call binding the contract method 0x7ea27a6f.
//
// Solidity: function COVE() view returns(address)
func (_CoveRegistry *CoveRegistryCaller) COVE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CoveRegistry.contract.Call(opts, &out, "COVE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// COVE is a free data retrieval call binding the contract method 0x7ea27a6f.
//
// Solidity: function COVE() view returns(address)
func (_CoveRegistry *CoveRegistrySession) COVE() (common.Address, error) {
	return _CoveRegistry.Contract.COVE(&_CoveRegistry.CallOpts)
}

// COVE is a free data retrieval call binding the contract method 0x7ea27a6f.
//
// Solidity: function COVE() view returns(address)
func (_CoveRegistry *CoveRegistryCallerSession) COVE() (common.Address, error) {
	return _CoveRegistry.Contract.COVE(&_CoveRegistry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CoveRegistry *CoveRegistryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CoveRegistry.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CoveRegistry *CoveRegistrySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CoveRegistry.Contract.DEFAULTADMINROLE(&_CoveRegistry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CoveRegistry *CoveRegistryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CoveRegistry.Contract.DEFAULTADMINROLE(&_CoveRegistry.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_CoveRegistry *CoveRegistryCaller) MANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CoveRegistry.contract.Call(opts, &out, "MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_CoveRegistry *CoveRegistrySession) MANAGERROLE() ([32]byte, error) {
	return _CoveRegistry.Contract.MANAGERROLE(&_CoveRegistry.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_CoveRegistry *CoveRegistryCallerSession) MANAGERROLE() ([32]byte, error) {
	return _CoveRegistry.Contract.MANAGERROLE(&_CoveRegistry.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_CoveRegistry *CoveRegistryCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CoveRegistry.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_CoveRegistry *CoveRegistrySession) PAUSERROLE() ([32]byte, error) {
	return _CoveRegistry.Contract.PAUSERROLE(&_CoveRegistry.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_CoveRegistry *CoveRegistryCallerSession) PAUSERROLE() ([32]byte, error) {
	return _CoveRegistry.Contract.PAUSERROLE(&_CoveRegistry.CallOpts)
}

// YEARNSTAKINGDELEGATE is a free data retrieval call binding the contract method 0x63748763.
//
// Solidity: function YEARN_STAKING_DELEGATE() view returns(address)
func (_CoveRegistry *CoveRegistryCaller) YEARNSTAKINGDELEGATE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CoveRegistry.contract.Call(opts, &out, "YEARN_STAKING_DELEGATE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// YEARNSTAKINGDELEGATE is a free data retrieval call binding the contract method 0x63748763.
//
// Solidity: function YEARN_STAKING_DELEGATE() view returns(address)
func (_CoveRegistry *CoveRegistrySession) YEARNSTAKINGDELEGATE() (common.Address, error) {
	return _CoveRegistry.Contract.YEARNSTAKINGDELEGATE(&_CoveRegistry.CallOpts)
}

// YEARNSTAKINGDELEGATE is a free data retrieval call binding the contract method 0x63748763.
//
// Solidity: function YEARN_STAKING_DELEGATE() view returns(address)
func (_CoveRegistry *CoveRegistryCallerSession) YEARNSTAKINGDELEGATE() (common.Address, error) {
	return _CoveRegistry.Contract.YEARNSTAKINGDELEGATE(&_CoveRegistry.CallOpts)
}

// Erc20RewardsGaugeImpl is a free data retrieval call binding the contract method 0x77ea46a2.
//
// Solidity: function erc20RewardsGaugeImpl() view returns(address)
func (_CoveRegistry *CoveRegistryCaller) Erc20RewardsGaugeImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CoveRegistry.contract.Call(opts, &out, "erc20RewardsGaugeImpl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc20RewardsGaugeImpl is a free data retrieval call binding the contract method 0x77ea46a2.
//
// Solidity: function erc20RewardsGaugeImpl() view returns(address)
func (_CoveRegistry *CoveRegistrySession) Erc20RewardsGaugeImpl() (common.Address, error) {
	return _CoveRegistry.Contract.Erc20RewardsGaugeImpl(&_CoveRegistry.CallOpts)
}

// Erc20RewardsGaugeImpl is a free data retrieval call binding the contract method 0x77ea46a2.
//
// Solidity: function erc20RewardsGaugeImpl() view returns(address)
func (_CoveRegistry *CoveRegistryCallerSession) Erc20RewardsGaugeImpl() (common.Address, error) {
	return _CoveRegistry.Contract.Erc20RewardsGaugeImpl(&_CoveRegistry.CallOpts)
}

// GaugeAdmin is a free data retrieval call binding the contract method 0xb6e70906.
//
// Solidity: function gaugeAdmin() view returns(address)
func (_CoveRegistry *CoveRegistryCaller) GaugeAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CoveRegistry.contract.Call(opts, &out, "gaugeAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GaugeAdmin is a free data retrieval call binding the contract method 0xb6e70906.
//
// Solidity: function gaugeAdmin() view returns(address)
func (_CoveRegistry *CoveRegistrySession) GaugeAdmin() (common.Address, error) {
	return _CoveRegistry.Contract.GaugeAdmin(&_CoveRegistry.CallOpts)
}

// GaugeAdmin is a free data retrieval call binding the contract method 0xb6e70906.
//
// Solidity: function gaugeAdmin() view returns(address)
func (_CoveRegistry *CoveRegistryCallerSession) GaugeAdmin() (common.Address, error) {
	return _CoveRegistry.Contract.GaugeAdmin(&_CoveRegistry.CallOpts)
}

// GaugeManager is a free data retrieval call binding the contract method 0xcbda3cee.
//
// Solidity: function gaugeManager() view returns(address)
func (_CoveRegistry *CoveRegistryCaller) GaugeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CoveRegistry.contract.Call(opts, &out, "gaugeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GaugeManager is a free data retrieval call binding the contract method 0xcbda3cee.
//
// Solidity: function gaugeManager() view returns(address)
func (_CoveRegistry *CoveRegistrySession) GaugeManager() (common.Address, error) {
	return _CoveRegistry.Contract.GaugeManager(&_CoveRegistry.CallOpts)
}

// GaugeManager is a free data retrieval call binding the contract method 0xcbda3cee.
//
// Solidity: function gaugeManager() view returns(address)
func (_CoveRegistry *CoveRegistryCallerSession) GaugeManager() (common.Address, error) {
	return _CoveRegistry.Contract.GaugeManager(&_CoveRegistry.CallOpts)
}

// GaugePauser is a free data retrieval call binding the contract method 0x95a8816f.
//
// Solidity: function gaugePauser() view returns(address)
func (_CoveRegistry *CoveRegistryCaller) GaugePauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CoveRegistry.contract.Call(opts, &out, "gaugePauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GaugePauser is a free data retrieval call binding the contract method 0x95a8816f.
//
// Solidity: function gaugePauser() view returns(address)
func (_CoveRegistry *CoveRegistrySession) GaugePauser() (common.Address, error) {
	return _CoveRegistry.Contract.GaugePauser(&_CoveRegistry.CallOpts)
}

// GaugePauser is a free data retrieval call binding the contract method 0x95a8816f.
//
// Solidity: function gaugePauser() view returns(address)
func (_CoveRegistry *CoveRegistryCallerSession) GaugePauser() (common.Address, error) {
	return _CoveRegistry.Contract.GaugePauser(&_CoveRegistry.CallOpts)
}

// GetAllGaugeInfo is a free data retrieval call binding the contract method 0x0005c412.
//
// Solidity: function getAllGaugeInfo(uint256 limit, uint256 offset) view returns((address,address,bool,address,address,address,address)[])
func (_CoveRegistry *CoveRegistryCaller) GetAllGaugeInfo(opts *bind.CallOpts, limit *big.Int, offset *big.Int) ([]CoveYearnGaugeFactoryGaugeInfo, error) {
	var out []interface{}
	err := _CoveRegistry.contract.Call(opts, &out, "getAllGaugeInfo", limit, offset)

	if err != nil {
		return *new([]CoveYearnGaugeFactoryGaugeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]CoveYearnGaugeFactoryGaugeInfo)).(*[]CoveYearnGaugeFactoryGaugeInfo)

	return out0, err

}

// GetAllGaugeInfo is a free data retrieval call binding the contract method 0x0005c412.
//
// Solidity: function getAllGaugeInfo(uint256 limit, uint256 offset) view returns((address,address,bool,address,address,address,address)[])
func (_CoveRegistry *CoveRegistrySession) GetAllGaugeInfo(limit *big.Int, offset *big.Int) ([]CoveYearnGaugeFactoryGaugeInfo, error) {
	return _CoveRegistry.Contract.GetAllGaugeInfo(&_CoveRegistry.CallOpts, limit, offset)
}

// GetAllGaugeInfo is a free data retrieval call binding the contract method 0x0005c412.
//
// Solidity: function getAllGaugeInfo(uint256 limit, uint256 offset) view returns((address,address,bool,address,address,address,address)[])
func (_CoveRegistry *CoveRegistryCallerSession) GetAllGaugeInfo(limit *big.Int, offset *big.Int) ([]CoveYearnGaugeFactoryGaugeInfo, error) {
	return _CoveRegistry.Contract.GetAllGaugeInfo(&_CoveRegistry.CallOpts, limit, offset)
}

// GetGaugeInfo is a free data retrieval call binding the contract method 0x5f31696d.
//
// Solidity: function getGaugeInfo(address yearnGauge) view returns((address,address,bool,address,address,address,address))
func (_CoveRegistry *CoveRegistryCaller) GetGaugeInfo(opts *bind.CallOpts, yearnGauge common.Address) (CoveYearnGaugeFactoryGaugeInfo, error) {
	var out []interface{}
	err := _CoveRegistry.contract.Call(opts, &out, "getGaugeInfo", yearnGauge)

	if err != nil {
		return *new(CoveYearnGaugeFactoryGaugeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(CoveYearnGaugeFactoryGaugeInfo)).(*CoveYearnGaugeFactoryGaugeInfo)

	return out0, err

}

// GetGaugeInfo is a free data retrieval call binding the contract method 0x5f31696d.
//
// Solidity: function getGaugeInfo(address yearnGauge) view returns((address,address,bool,address,address,address,address))
func (_CoveRegistry *CoveRegistrySession) GetGaugeInfo(yearnGauge common.Address) (CoveYearnGaugeFactoryGaugeInfo, error) {
	return _CoveRegistry.Contract.GetGaugeInfo(&_CoveRegistry.CallOpts, yearnGauge)
}

// GetGaugeInfo is a free data retrieval call binding the contract method 0x5f31696d.
//
// Solidity: function getGaugeInfo(address yearnGauge) view returns((address,address,bool,address,address,address,address))
func (_CoveRegistry *CoveRegistryCallerSession) GetGaugeInfo(yearnGauge common.Address) (CoveYearnGaugeFactoryGaugeInfo, error) {
	return _CoveRegistry.Contract.GetGaugeInfo(&_CoveRegistry.CallOpts, yearnGauge)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CoveRegistry *CoveRegistryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _CoveRegistry.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CoveRegistry *CoveRegistrySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CoveRegistry.Contract.GetRoleAdmin(&_CoveRegistry.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CoveRegistry *CoveRegistryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CoveRegistry.Contract.GetRoleAdmin(&_CoveRegistry.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_CoveRegistry *CoveRegistryCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CoveRegistry.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_CoveRegistry *CoveRegistrySession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _CoveRegistry.Contract.GetRoleMember(&_CoveRegistry.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_CoveRegistry *CoveRegistryCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _CoveRegistry.Contract.GetRoleMember(&_CoveRegistry.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_CoveRegistry *CoveRegistryCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _CoveRegistry.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_CoveRegistry *CoveRegistrySession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _CoveRegistry.Contract.GetRoleMemberCount(&_CoveRegistry.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_CoveRegistry *CoveRegistryCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _CoveRegistry.Contract.GetRoleMemberCount(&_CoveRegistry.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CoveRegistry *CoveRegistryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _CoveRegistry.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CoveRegistry *CoveRegistrySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CoveRegistry.Contract.HasRole(&_CoveRegistry.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CoveRegistry *CoveRegistryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CoveRegistry.Contract.HasRole(&_CoveRegistry.CallOpts, role, account)
}

// NumOfSupportedYearnGauges is a free data retrieval call binding the contract method 0x5899ef66.
//
// Solidity: function numOfSupportedYearnGauges() view returns(uint256)
func (_CoveRegistry *CoveRegistryCaller) NumOfSupportedYearnGauges(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CoveRegistry.contract.Call(opts, &out, "numOfSupportedYearnGauges")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumOfSupportedYearnGauges is a free data retrieval call binding the contract method 0x5899ef66.
//
// Solidity: function numOfSupportedYearnGauges() view returns(uint256)
func (_CoveRegistry *CoveRegistrySession) NumOfSupportedYearnGauges() (*big.Int, error) {
	return _CoveRegistry.Contract.NumOfSupportedYearnGauges(&_CoveRegistry.CallOpts)
}

// NumOfSupportedYearnGauges is a free data retrieval call binding the contract method 0x5899ef66.
//
// Solidity: function numOfSupportedYearnGauges() view returns(uint256)
func (_CoveRegistry *CoveRegistryCallerSession) NumOfSupportedYearnGauges() (*big.Int, error) {
	return _CoveRegistry.Contract.NumOfSupportedYearnGauges(&_CoveRegistry.CallOpts)
}

// RewardForwarderImpl is a free data retrieval call binding the contract method 0x4a16de57.
//
// Solidity: function rewardForwarderImpl() view returns(address)
func (_CoveRegistry *CoveRegistryCaller) RewardForwarderImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CoveRegistry.contract.Call(opts, &out, "rewardForwarderImpl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardForwarderImpl is a free data retrieval call binding the contract method 0x4a16de57.
//
// Solidity: function rewardForwarderImpl() view returns(address)
func (_CoveRegistry *CoveRegistrySession) RewardForwarderImpl() (common.Address, error) {
	return _CoveRegistry.Contract.RewardForwarderImpl(&_CoveRegistry.CallOpts)
}

// RewardForwarderImpl is a free data retrieval call binding the contract method 0x4a16de57.
//
// Solidity: function rewardForwarderImpl() view returns(address)
func (_CoveRegistry *CoveRegistryCallerSession) RewardForwarderImpl() (common.Address, error) {
	return _CoveRegistry.Contract.RewardForwarderImpl(&_CoveRegistry.CallOpts)
}

// SupportedYearnGauges is a free data retrieval call binding the contract method 0x5f2e2ca2.
//
// Solidity: function supportedYearnGauges(uint256 ) view returns(address)
func (_CoveRegistry *CoveRegistryCaller) SupportedYearnGauges(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CoveRegistry.contract.Call(opts, &out, "supportedYearnGauges", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SupportedYearnGauges is a free data retrieval call binding the contract method 0x5f2e2ca2.
//
// Solidity: function supportedYearnGauges(uint256 ) view returns(address)
func (_CoveRegistry *CoveRegistrySession) SupportedYearnGauges(arg0 *big.Int) (common.Address, error) {
	return _CoveRegistry.Contract.SupportedYearnGauges(&_CoveRegistry.CallOpts, arg0)
}

// SupportedYearnGauges is a free data retrieval call binding the contract method 0x5f2e2ca2.
//
// Solidity: function supportedYearnGauges(uint256 ) view returns(address)
func (_CoveRegistry *CoveRegistryCallerSession) SupportedYearnGauges(arg0 *big.Int) (common.Address, error) {
	return _CoveRegistry.Contract.SupportedYearnGauges(&_CoveRegistry.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CoveRegistry *CoveRegistryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CoveRegistry.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CoveRegistry *CoveRegistrySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CoveRegistry.Contract.SupportsInterface(&_CoveRegistry.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CoveRegistry *CoveRegistryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CoveRegistry.Contract.SupportsInterface(&_CoveRegistry.CallOpts, interfaceId)
}

// YearnGaugeInfoStored is a free data retrieval call binding the contract method 0x440cf170.
//
// Solidity: function yearnGaugeInfoStored(address ) view returns(address coveYearnStrategy, address autoCompoundingGauge, address nonAutoCompoundingGauge)
func (_CoveRegistry *CoveRegistryCaller) YearnGaugeInfoStored(opts *bind.CallOpts, arg0 common.Address) (struct {
	CoveYearnStrategy       common.Address
	AutoCompoundingGauge    common.Address
	NonAutoCompoundingGauge common.Address
}, error) {
	var out []interface{}
	err := _CoveRegistry.contract.Call(opts, &out, "yearnGaugeInfoStored", arg0)

	outstruct := new(struct {
		CoveYearnStrategy       common.Address
		AutoCompoundingGauge    common.Address
		NonAutoCompoundingGauge common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CoveYearnStrategy = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.AutoCompoundingGauge = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.NonAutoCompoundingGauge = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// YearnGaugeInfoStored is a free data retrieval call binding the contract method 0x440cf170.
//
// Solidity: function yearnGaugeInfoStored(address ) view returns(address coveYearnStrategy, address autoCompoundingGauge, address nonAutoCompoundingGauge)
func (_CoveRegistry *CoveRegistrySession) YearnGaugeInfoStored(arg0 common.Address) (struct {
	CoveYearnStrategy       common.Address
	AutoCompoundingGauge    common.Address
	NonAutoCompoundingGauge common.Address
}, error) {
	return _CoveRegistry.Contract.YearnGaugeInfoStored(&_CoveRegistry.CallOpts, arg0)
}

// YearnGaugeInfoStored is a free data retrieval call binding the contract method 0x440cf170.
//
// Solidity: function yearnGaugeInfoStored(address ) view returns(address coveYearnStrategy, address autoCompoundingGauge, address nonAutoCompoundingGauge)
func (_CoveRegistry *CoveRegistryCallerSession) YearnGaugeInfoStored(arg0 common.Address) (struct {
	CoveYearnStrategy       common.Address
	AutoCompoundingGauge    common.Address
	NonAutoCompoundingGauge common.Address
}, error) {
	return _CoveRegistry.Contract.YearnGaugeInfoStored(&_CoveRegistry.CallOpts, arg0)
}

// YsdRewardsGaugeImpl is a free data retrieval call binding the contract method 0xdaa6a0e8.
//
// Solidity: function ysdRewardsGaugeImpl() view returns(address)
func (_CoveRegistry *CoveRegistryCaller) YsdRewardsGaugeImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CoveRegistry.contract.Call(opts, &out, "ysdRewardsGaugeImpl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// YsdRewardsGaugeImpl is a free data retrieval call binding the contract method 0xdaa6a0e8.
//
// Solidity: function ysdRewardsGaugeImpl() view returns(address)
func (_CoveRegistry *CoveRegistrySession) YsdRewardsGaugeImpl() (common.Address, error) {
	return _CoveRegistry.Contract.YsdRewardsGaugeImpl(&_CoveRegistry.CallOpts)
}

// YsdRewardsGaugeImpl is a free data retrieval call binding the contract method 0xdaa6a0e8.
//
// Solidity: function ysdRewardsGaugeImpl() view returns(address)
func (_CoveRegistry *CoveRegistryCallerSession) YsdRewardsGaugeImpl() (common.Address, error) {
	return _CoveRegistry.Contract.YsdRewardsGaugeImpl(&_CoveRegistry.CallOpts)
}

// DeployCoveGauges is a paid mutator transaction binding the contract method 0x073048de.
//
// Solidity: function deployCoveGauges(address coveYearnStrategy) returns()
func (_CoveRegistry *CoveRegistryTransactor) DeployCoveGauges(opts *bind.TransactOpts, coveYearnStrategy common.Address) (*types.Transaction, error) {
	return _CoveRegistry.contract.Transact(opts, "deployCoveGauges", coveYearnStrategy)
}

// DeployCoveGauges is a paid mutator transaction binding the contract method 0x073048de.
//
// Solidity: function deployCoveGauges(address coveYearnStrategy) returns()
func (_CoveRegistry *CoveRegistrySession) DeployCoveGauges(coveYearnStrategy common.Address) (*types.Transaction, error) {
	return _CoveRegistry.Contract.DeployCoveGauges(&_CoveRegistry.TransactOpts, coveYearnStrategy)
}

// DeployCoveGauges is a paid mutator transaction binding the contract method 0x073048de.
//
// Solidity: function deployCoveGauges(address coveYearnStrategy) returns()
func (_CoveRegistry *CoveRegistryTransactorSession) DeployCoveGauges(coveYearnStrategy common.Address) (*types.Transaction, error) {
	return _CoveRegistry.Contract.DeployCoveGauges(&_CoveRegistry.TransactOpts, coveYearnStrategy)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CoveRegistry *CoveRegistryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CoveRegistry.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CoveRegistry *CoveRegistrySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CoveRegistry.Contract.GrantRole(&_CoveRegistry.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CoveRegistry *CoveRegistryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CoveRegistry.Contract.GrantRole(&_CoveRegistry.TransactOpts, role, account)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_CoveRegistry *CoveRegistryTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _CoveRegistry.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_CoveRegistry *CoveRegistrySession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _CoveRegistry.Contract.Multicall(&_CoveRegistry.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_CoveRegistry *CoveRegistryTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _CoveRegistry.Contract.Multicall(&_CoveRegistry.TransactOpts, data)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_CoveRegistry *CoveRegistryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CoveRegistry.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_CoveRegistry *CoveRegistrySession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CoveRegistry.Contract.RenounceRole(&_CoveRegistry.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_CoveRegistry *CoveRegistryTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CoveRegistry.Contract.RenounceRole(&_CoveRegistry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CoveRegistry *CoveRegistryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CoveRegistry.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CoveRegistry *CoveRegistrySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CoveRegistry.Contract.RevokeRole(&_CoveRegistry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CoveRegistry *CoveRegistryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CoveRegistry.Contract.RevokeRole(&_CoveRegistry.TransactOpts, role, account)
}

// SetERC20RewardsGaugeImplementation is a paid mutator transaction binding the contract method 0x4b945f94.
//
// Solidity: function setERC20RewardsGaugeImplementation(address impl) returns()
func (_CoveRegistry *CoveRegistryTransactor) SetERC20RewardsGaugeImplementation(opts *bind.TransactOpts, impl common.Address) (*types.Transaction, error) {
	return _CoveRegistry.contract.Transact(opts, "setERC20RewardsGaugeImplementation", impl)
}

// SetERC20RewardsGaugeImplementation is a paid mutator transaction binding the contract method 0x4b945f94.
//
// Solidity: function setERC20RewardsGaugeImplementation(address impl) returns()
func (_CoveRegistry *CoveRegistrySession) SetERC20RewardsGaugeImplementation(impl common.Address) (*types.Transaction, error) {
	return _CoveRegistry.Contract.SetERC20RewardsGaugeImplementation(&_CoveRegistry.TransactOpts, impl)
}

// SetERC20RewardsGaugeImplementation is a paid mutator transaction binding the contract method 0x4b945f94.
//
// Solidity: function setERC20RewardsGaugeImplementation(address impl) returns()
func (_CoveRegistry *CoveRegistryTransactorSession) SetERC20RewardsGaugeImplementation(impl common.Address) (*types.Transaction, error) {
	return _CoveRegistry.Contract.SetERC20RewardsGaugeImplementation(&_CoveRegistry.TransactOpts, impl)
}

// SetGaugeAdmin is a paid mutator transaction binding the contract method 0xcc9b4a5a.
//
// Solidity: function setGaugeAdmin(address admin) returns()
func (_CoveRegistry *CoveRegistryTransactor) SetGaugeAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _CoveRegistry.contract.Transact(opts, "setGaugeAdmin", admin)
}

// SetGaugeAdmin is a paid mutator transaction binding the contract method 0xcc9b4a5a.
//
// Solidity: function setGaugeAdmin(address admin) returns()
func (_CoveRegistry *CoveRegistrySession) SetGaugeAdmin(admin common.Address) (*types.Transaction, error) {
	return _CoveRegistry.Contract.SetGaugeAdmin(&_CoveRegistry.TransactOpts, admin)
}

// SetGaugeAdmin is a paid mutator transaction binding the contract method 0xcc9b4a5a.
//
// Solidity: function setGaugeAdmin(address admin) returns()
func (_CoveRegistry *CoveRegistryTransactorSession) SetGaugeAdmin(admin common.Address) (*types.Transaction, error) {
	return _CoveRegistry.Contract.SetGaugeAdmin(&_CoveRegistry.TransactOpts, admin)
}

// SetGaugeManager is a paid mutator transaction binding the contract method 0xd8cfc7f0.
//
// Solidity: function setGaugeManager(address manager) returns()
func (_CoveRegistry *CoveRegistryTransactor) SetGaugeManager(opts *bind.TransactOpts, manager common.Address) (*types.Transaction, error) {
	return _CoveRegistry.contract.Transact(opts, "setGaugeManager", manager)
}

// SetGaugeManager is a paid mutator transaction binding the contract method 0xd8cfc7f0.
//
// Solidity: function setGaugeManager(address manager) returns()
func (_CoveRegistry *CoveRegistrySession) SetGaugeManager(manager common.Address) (*types.Transaction, error) {
	return _CoveRegistry.Contract.SetGaugeManager(&_CoveRegistry.TransactOpts, manager)
}

// SetGaugeManager is a paid mutator transaction binding the contract method 0xd8cfc7f0.
//
// Solidity: function setGaugeManager(address manager) returns()
func (_CoveRegistry *CoveRegistryTransactorSession) SetGaugeManager(manager common.Address) (*types.Transaction, error) {
	return _CoveRegistry.Contract.SetGaugeManager(&_CoveRegistry.TransactOpts, manager)
}

// SetGaugePauser is a paid mutator transaction binding the contract method 0x806ed492.
//
// Solidity: function setGaugePauser(address pauser) returns()
func (_CoveRegistry *CoveRegistryTransactor) SetGaugePauser(opts *bind.TransactOpts, pauser common.Address) (*types.Transaction, error) {
	return _CoveRegistry.contract.Transact(opts, "setGaugePauser", pauser)
}

// SetGaugePauser is a paid mutator transaction binding the contract method 0x806ed492.
//
// Solidity: function setGaugePauser(address pauser) returns()
func (_CoveRegistry *CoveRegistrySession) SetGaugePauser(pauser common.Address) (*types.Transaction, error) {
	return _CoveRegistry.Contract.SetGaugePauser(&_CoveRegistry.TransactOpts, pauser)
}

// SetGaugePauser is a paid mutator transaction binding the contract method 0x806ed492.
//
// Solidity: function setGaugePauser(address pauser) returns()
func (_CoveRegistry *CoveRegistryTransactorSession) SetGaugePauser(pauser common.Address) (*types.Transaction, error) {
	return _CoveRegistry.Contract.SetGaugePauser(&_CoveRegistry.TransactOpts, pauser)
}

// SetRewardForwarderImplementation is a paid mutator transaction binding the contract method 0x3032fedb.
//
// Solidity: function setRewardForwarderImplementation(address impl) returns()
func (_CoveRegistry *CoveRegistryTransactor) SetRewardForwarderImplementation(opts *bind.TransactOpts, impl common.Address) (*types.Transaction, error) {
	return _CoveRegistry.contract.Transact(opts, "setRewardForwarderImplementation", impl)
}

// SetRewardForwarderImplementation is a paid mutator transaction binding the contract method 0x3032fedb.
//
// Solidity: function setRewardForwarderImplementation(address impl) returns()
func (_CoveRegistry *CoveRegistrySession) SetRewardForwarderImplementation(impl common.Address) (*types.Transaction, error) {
	return _CoveRegistry.Contract.SetRewardForwarderImplementation(&_CoveRegistry.TransactOpts, impl)
}

// SetRewardForwarderImplementation is a paid mutator transaction binding the contract method 0x3032fedb.
//
// Solidity: function setRewardForwarderImplementation(address impl) returns()
func (_CoveRegistry *CoveRegistryTransactorSession) SetRewardForwarderImplementation(impl common.Address) (*types.Transaction, error) {
	return _CoveRegistry.Contract.SetRewardForwarderImplementation(&_CoveRegistry.TransactOpts, impl)
}

// SetYsdRewardsGaugeImplementation is a paid mutator transaction binding the contract method 0xd84df686.
//
// Solidity: function setYsdRewardsGaugeImplementation(address impl) returns()
func (_CoveRegistry *CoveRegistryTransactor) SetYsdRewardsGaugeImplementation(opts *bind.TransactOpts, impl common.Address) (*types.Transaction, error) {
	return _CoveRegistry.contract.Transact(opts, "setYsdRewardsGaugeImplementation", impl)
}

// SetYsdRewardsGaugeImplementation is a paid mutator transaction binding the contract method 0xd84df686.
//
// Solidity: function setYsdRewardsGaugeImplementation(address impl) returns()
func (_CoveRegistry *CoveRegistrySession) SetYsdRewardsGaugeImplementation(impl common.Address) (*types.Transaction, error) {
	return _CoveRegistry.Contract.SetYsdRewardsGaugeImplementation(&_CoveRegistry.TransactOpts, impl)
}

// SetYsdRewardsGaugeImplementation is a paid mutator transaction binding the contract method 0xd84df686.
//
// Solidity: function setYsdRewardsGaugeImplementation(address impl) returns()
func (_CoveRegistry *CoveRegistryTransactorSession) SetYsdRewardsGaugeImplementation(impl common.Address) (*types.Transaction, error) {
	return _CoveRegistry.Contract.SetYsdRewardsGaugeImplementation(&_CoveRegistry.TransactOpts, impl)
}

// CoveRegistryCoveGaugesDeployedIterator is returned from FilterCoveGaugesDeployed and is used to iterate over the raw logs and unpacked data for CoveGaugesDeployed events raised by the CoveRegistry contract.
type CoveRegistryCoveGaugesDeployedIterator struct {
	Event *CoveRegistryCoveGaugesDeployed // Event containing the contract specifics and raw log

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
func (it *CoveRegistryCoveGaugesDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoveRegistryCoveGaugesDeployed)
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
		it.Event = new(CoveRegistryCoveGaugesDeployed)
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
func (it *CoveRegistryCoveGaugesDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoveRegistryCoveGaugesDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoveRegistryCoveGaugesDeployed represents a CoveGaugesDeployed event raised by the CoveRegistry contract.
type CoveRegistryCoveGaugesDeployed struct {
	YearnGauge              common.Address
	CoveYearnStrategy       common.Address
	AutoCompoundingGauge    common.Address
	NonAutoCompoundingGauge common.Address
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterCoveGaugesDeployed is a free log retrieval operation binding the contract event 0xf96f1963c1950958459fa57e3589a9795744fceff64938b229cdb647fb51e09e.
//
// Solidity: event CoveGaugesDeployed(address yearnGauge, address coveYearnStrategy, address autoCompoundingGauge, address nonAutoCompoundingGauge)
func (_CoveRegistry *CoveRegistryFilterer) FilterCoveGaugesDeployed(opts *bind.FilterOpts) (*CoveRegistryCoveGaugesDeployedIterator, error) {

	logs, sub, err := _CoveRegistry.contract.FilterLogs(opts, "CoveGaugesDeployed")
	if err != nil {
		return nil, err
	}
	return &CoveRegistryCoveGaugesDeployedIterator{contract: _CoveRegistry.contract, event: "CoveGaugesDeployed", logs: logs, sub: sub}, nil
}

// WatchCoveGaugesDeployed is a free log subscription operation binding the contract event 0xf96f1963c1950958459fa57e3589a9795744fceff64938b229cdb647fb51e09e.
//
// Solidity: event CoveGaugesDeployed(address yearnGauge, address coveYearnStrategy, address autoCompoundingGauge, address nonAutoCompoundingGauge)
func (_CoveRegistry *CoveRegistryFilterer) WatchCoveGaugesDeployed(opts *bind.WatchOpts, sink chan<- *CoveRegistryCoveGaugesDeployed) (event.Subscription, error) {

	logs, sub, err := _CoveRegistry.contract.WatchLogs(opts, "CoveGaugesDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoveRegistryCoveGaugesDeployed)
				if err := _CoveRegistry.contract.UnpackLog(event, "CoveGaugesDeployed", log); err != nil {
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

// ParseCoveGaugesDeployed is a log parse operation binding the contract event 0xf96f1963c1950958459fa57e3589a9795744fceff64938b229cdb647fb51e09e.
//
// Solidity: event CoveGaugesDeployed(address yearnGauge, address coveYearnStrategy, address autoCompoundingGauge, address nonAutoCompoundingGauge)
func (_CoveRegistry *CoveRegistryFilterer) ParseCoveGaugesDeployed(log types.Log) (*CoveRegistryCoveGaugesDeployed, error) {
	event := new(CoveRegistryCoveGaugesDeployed)
	if err := _CoveRegistry.contract.UnpackLog(event, "CoveGaugesDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoveRegistryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the CoveRegistry contract.
type CoveRegistryRoleAdminChangedIterator struct {
	Event *CoveRegistryRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *CoveRegistryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoveRegistryRoleAdminChanged)
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
		it.Event = new(CoveRegistryRoleAdminChanged)
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
func (it *CoveRegistryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoveRegistryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoveRegistryRoleAdminChanged represents a RoleAdminChanged event raised by the CoveRegistry contract.
type CoveRegistryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CoveRegistry *CoveRegistryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*CoveRegistryRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _CoveRegistry.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &CoveRegistryRoleAdminChangedIterator{contract: _CoveRegistry.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CoveRegistry *CoveRegistryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *CoveRegistryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _CoveRegistry.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoveRegistryRoleAdminChanged)
				if err := _CoveRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CoveRegistry *CoveRegistryFilterer) ParseRoleAdminChanged(log types.Log) (*CoveRegistryRoleAdminChanged, error) {
	event := new(CoveRegistryRoleAdminChanged)
	if err := _CoveRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoveRegistryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the CoveRegistry contract.
type CoveRegistryRoleGrantedIterator struct {
	Event *CoveRegistryRoleGranted // Event containing the contract specifics and raw log

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
func (it *CoveRegistryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoveRegistryRoleGranted)
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
		it.Event = new(CoveRegistryRoleGranted)
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
func (it *CoveRegistryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoveRegistryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoveRegistryRoleGranted represents a RoleGranted event raised by the CoveRegistry contract.
type CoveRegistryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CoveRegistry *CoveRegistryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CoveRegistryRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CoveRegistry.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CoveRegistryRoleGrantedIterator{contract: _CoveRegistry.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CoveRegistry *CoveRegistryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *CoveRegistryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CoveRegistry.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoveRegistryRoleGranted)
				if err := _CoveRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CoveRegistry *CoveRegistryFilterer) ParseRoleGranted(log types.Log) (*CoveRegistryRoleGranted, error) {
	event := new(CoveRegistryRoleGranted)
	if err := _CoveRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoveRegistryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the CoveRegistry contract.
type CoveRegistryRoleRevokedIterator struct {
	Event *CoveRegistryRoleRevoked // Event containing the contract specifics and raw log

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
func (it *CoveRegistryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoveRegistryRoleRevoked)
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
		it.Event = new(CoveRegistryRoleRevoked)
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
func (it *CoveRegistryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoveRegistryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoveRegistryRoleRevoked represents a RoleRevoked event raised by the CoveRegistry contract.
type CoveRegistryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CoveRegistry *CoveRegistryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CoveRegistryRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CoveRegistry.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CoveRegistryRoleRevokedIterator{contract: _CoveRegistry.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CoveRegistry *CoveRegistryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *CoveRegistryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CoveRegistry.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoveRegistryRoleRevoked)
				if err := _CoveRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CoveRegistry *CoveRegistryFilterer) ParseRoleRevoked(log types.Log) (*CoveRegistryRoleRevoked, error) {
	event := new(CoveRegistryRoleRevoked)
	if err := _CoveRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
