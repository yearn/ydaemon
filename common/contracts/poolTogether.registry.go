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

// PoolTogetherRegistryMetaData contains all meta data concerning the PoolTogetherRegistry contract.
var PoolTogetherRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractPrizeVault\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC4626\",\"name\":\"yieldVault\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractPrizePool\",\"name\":\"prizePool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"NewPrizeVault\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allVaults\",\"outputs\":[{\"internalType\":\"contractPrizeVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"contractIERC4626\",\"name\":\"_yieldVault\",\"type\":\"address\"},{\"internalType\":\"contractPrizePool\",\"name\":\"_prizePool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_claimer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_yieldFeeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_yieldFeePercentage\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_yieldBuffer\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"deployVault\",\"outputs\":[{\"internalType\":\"contractPrizeVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"deployedVaults\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"deployedByFactory\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"deployer\",\"type\":\"address\"}],\"name\":\"deployerNonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalVaults\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PoolTogetherRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolTogetherRegistryMetaData.ABI instead.
var PoolTogetherRegistryABI = PoolTogetherRegistryMetaData.ABI

// PoolTogetherRegistry is an auto generated Go binding around an Ethereum contract.
type PoolTogetherRegistry struct {
	PoolTogetherRegistryCaller     // Read-only binding to the contract
	PoolTogetherRegistryTransactor // Write-only binding to the contract
	PoolTogetherRegistryFilterer   // Log filterer for contract events
}

// PoolTogetherRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolTogetherRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolTogetherRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolTogetherRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolTogetherRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolTogetherRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolTogetherRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolTogetherRegistrySession struct {
	Contract     *PoolTogetherRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PoolTogetherRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolTogetherRegistryCallerSession struct {
	Contract *PoolTogetherRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// PoolTogetherRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolTogetherRegistryTransactorSession struct {
	Contract     *PoolTogetherRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// PoolTogetherRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolTogetherRegistryRaw struct {
	Contract *PoolTogetherRegistry // Generic contract binding to access the raw methods on
}

// PoolTogetherRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolTogetherRegistryCallerRaw struct {
	Contract *PoolTogetherRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// PoolTogetherRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolTogetherRegistryTransactorRaw struct {
	Contract *PoolTogetherRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoolTogetherRegistry creates a new instance of PoolTogetherRegistry, bound to a specific deployed contract.
func NewPoolTogetherRegistry(address common.Address, backend bind.ContractBackend) (*PoolTogetherRegistry, error) {
	contract, err := bindPoolTogetherRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PoolTogetherRegistry{PoolTogetherRegistryCaller: PoolTogetherRegistryCaller{contract: contract}, PoolTogetherRegistryTransactor: PoolTogetherRegistryTransactor{contract: contract}, PoolTogetherRegistryFilterer: PoolTogetherRegistryFilterer{contract: contract}}, nil
}

// NewPoolTogetherRegistryCaller creates a new read-only instance of PoolTogetherRegistry, bound to a specific deployed contract.
func NewPoolTogetherRegistryCaller(address common.Address, caller bind.ContractCaller) (*PoolTogetherRegistryCaller, error) {
	contract, err := bindPoolTogetherRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolTogetherRegistryCaller{contract: contract}, nil
}

// NewPoolTogetherRegistryTransactor creates a new write-only instance of PoolTogetherRegistry, bound to a specific deployed contract.
func NewPoolTogetherRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolTogetherRegistryTransactor, error) {
	contract, err := bindPoolTogetherRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolTogetherRegistryTransactor{contract: contract}, nil
}

// NewPoolTogetherRegistryFilterer creates a new log filterer instance of PoolTogetherRegistry, bound to a specific deployed contract.
func NewPoolTogetherRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolTogetherRegistryFilterer, error) {
	contract, err := bindPoolTogetherRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolTogetherRegistryFilterer{contract: contract}, nil
}

// bindPoolTogetherRegistry binds a generic wrapper to an already deployed contract.
func bindPoolTogetherRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PoolTogetherRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolTogetherRegistry *PoolTogetherRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolTogetherRegistry.Contract.PoolTogetherRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolTogetherRegistry *PoolTogetherRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolTogetherRegistry.Contract.PoolTogetherRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolTogetherRegistry *PoolTogetherRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolTogetherRegistry.Contract.PoolTogetherRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolTogetherRegistry *PoolTogetherRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolTogetherRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolTogetherRegistry *PoolTogetherRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolTogetherRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolTogetherRegistry *PoolTogetherRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolTogetherRegistry.Contract.contract.Transact(opts, method, params...)
}

// AllVaults is a free data retrieval call binding the contract method 0x9094a91e.
//
// Solidity: function allVaults(uint256 ) view returns(address)
func (_PoolTogetherRegistry *PoolTogetherRegistryCaller) AllVaults(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PoolTogetherRegistry.contract.Call(opts, &out, "allVaults", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllVaults is a free data retrieval call binding the contract method 0x9094a91e.
//
// Solidity: function allVaults(uint256 ) view returns(address)
func (_PoolTogetherRegistry *PoolTogetherRegistrySession) AllVaults(arg0 *big.Int) (common.Address, error) {
	return _PoolTogetherRegistry.Contract.AllVaults(&_PoolTogetherRegistry.CallOpts, arg0)
}

// AllVaults is a free data retrieval call binding the contract method 0x9094a91e.
//
// Solidity: function allVaults(uint256 ) view returns(address)
func (_PoolTogetherRegistry *PoolTogetherRegistryCallerSession) AllVaults(arg0 *big.Int) (common.Address, error) {
	return _PoolTogetherRegistry.Contract.AllVaults(&_PoolTogetherRegistry.CallOpts, arg0)
}

// DeployedVaults is a free data retrieval call binding the contract method 0xfaf20d11.
//
// Solidity: function deployedVaults(address vault) view returns(bool deployedByFactory)
func (_PoolTogetherRegistry *PoolTogetherRegistryCaller) DeployedVaults(opts *bind.CallOpts, vault common.Address) (bool, error) {
	var out []interface{}
	err := _PoolTogetherRegistry.contract.Call(opts, &out, "deployedVaults", vault)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DeployedVaults is a free data retrieval call binding the contract method 0xfaf20d11.
//
// Solidity: function deployedVaults(address vault) view returns(bool deployedByFactory)
func (_PoolTogetherRegistry *PoolTogetherRegistrySession) DeployedVaults(vault common.Address) (bool, error) {
	return _PoolTogetherRegistry.Contract.DeployedVaults(&_PoolTogetherRegistry.CallOpts, vault)
}

// DeployedVaults is a free data retrieval call binding the contract method 0xfaf20d11.
//
// Solidity: function deployedVaults(address vault) view returns(bool deployedByFactory)
func (_PoolTogetherRegistry *PoolTogetherRegistryCallerSession) DeployedVaults(vault common.Address) (bool, error) {
	return _PoolTogetherRegistry.Contract.DeployedVaults(&_PoolTogetherRegistry.CallOpts, vault)
}

// DeployerNonces is a free data retrieval call binding the contract method 0x1f61805d.
//
// Solidity: function deployerNonces(address deployer) view returns(uint256 nonce)
func (_PoolTogetherRegistry *PoolTogetherRegistryCaller) DeployerNonces(opts *bind.CallOpts, deployer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PoolTogetherRegistry.contract.Call(opts, &out, "deployerNonces", deployer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeployerNonces is a free data retrieval call binding the contract method 0x1f61805d.
//
// Solidity: function deployerNonces(address deployer) view returns(uint256 nonce)
func (_PoolTogetherRegistry *PoolTogetherRegistrySession) DeployerNonces(deployer common.Address) (*big.Int, error) {
	return _PoolTogetherRegistry.Contract.DeployerNonces(&_PoolTogetherRegistry.CallOpts, deployer)
}

// DeployerNonces is a free data retrieval call binding the contract method 0x1f61805d.
//
// Solidity: function deployerNonces(address deployer) view returns(uint256 nonce)
func (_PoolTogetherRegistry *PoolTogetherRegistryCallerSession) DeployerNonces(deployer common.Address) (*big.Int, error) {
	return _PoolTogetherRegistry.Contract.DeployerNonces(&_PoolTogetherRegistry.CallOpts, deployer)
}

// TotalVaults is a free data retrieval call binding the contract method 0x8d654023.
//
// Solidity: function totalVaults() view returns(uint256)
func (_PoolTogetherRegistry *PoolTogetherRegistryCaller) TotalVaults(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoolTogetherRegistry.contract.Call(opts, &out, "totalVaults")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalVaults is a free data retrieval call binding the contract method 0x8d654023.
//
// Solidity: function totalVaults() view returns(uint256)
func (_PoolTogetherRegistry *PoolTogetherRegistrySession) TotalVaults() (*big.Int, error) {
	return _PoolTogetherRegistry.Contract.TotalVaults(&_PoolTogetherRegistry.CallOpts)
}

// TotalVaults is a free data retrieval call binding the contract method 0x8d654023.
//
// Solidity: function totalVaults() view returns(uint256)
func (_PoolTogetherRegistry *PoolTogetherRegistryCallerSession) TotalVaults() (*big.Int, error) {
	return _PoolTogetherRegistry.Contract.TotalVaults(&_PoolTogetherRegistry.CallOpts)
}

// DeployVault is a paid mutator transaction binding the contract method 0x781d653c.
//
// Solidity: function deployVault(string _name, string _symbol, address _yieldVault, address _prizePool, address _claimer, address _yieldFeeRecipient, uint32 _yieldFeePercentage, uint256 _yieldBuffer, address _owner) returns(address)
func (_PoolTogetherRegistry *PoolTogetherRegistryTransactor) DeployVault(opts *bind.TransactOpts, _name string, _symbol string, _yieldVault common.Address, _prizePool common.Address, _claimer common.Address, _yieldFeeRecipient common.Address, _yieldFeePercentage uint32, _yieldBuffer *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _PoolTogetherRegistry.contract.Transact(opts, "deployVault", _name, _symbol, _yieldVault, _prizePool, _claimer, _yieldFeeRecipient, _yieldFeePercentage, _yieldBuffer, _owner)
}

// DeployVault is a paid mutator transaction binding the contract method 0x781d653c.
//
// Solidity: function deployVault(string _name, string _symbol, address _yieldVault, address _prizePool, address _claimer, address _yieldFeeRecipient, uint32 _yieldFeePercentage, uint256 _yieldBuffer, address _owner) returns(address)
func (_PoolTogetherRegistry *PoolTogetherRegistrySession) DeployVault(_name string, _symbol string, _yieldVault common.Address, _prizePool common.Address, _claimer common.Address, _yieldFeeRecipient common.Address, _yieldFeePercentage uint32, _yieldBuffer *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _PoolTogetherRegistry.Contract.DeployVault(&_PoolTogetherRegistry.TransactOpts, _name, _symbol, _yieldVault, _prizePool, _claimer, _yieldFeeRecipient, _yieldFeePercentage, _yieldBuffer, _owner)
}

// DeployVault is a paid mutator transaction binding the contract method 0x781d653c.
//
// Solidity: function deployVault(string _name, string _symbol, address _yieldVault, address _prizePool, address _claimer, address _yieldFeeRecipient, uint32 _yieldFeePercentage, uint256 _yieldBuffer, address _owner) returns(address)
func (_PoolTogetherRegistry *PoolTogetherRegistryTransactorSession) DeployVault(_name string, _symbol string, _yieldVault common.Address, _prizePool common.Address, _claimer common.Address, _yieldFeeRecipient common.Address, _yieldFeePercentage uint32, _yieldBuffer *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _PoolTogetherRegistry.Contract.DeployVault(&_PoolTogetherRegistry.TransactOpts, _name, _symbol, _yieldVault, _prizePool, _claimer, _yieldFeeRecipient, _yieldFeePercentage, _yieldBuffer, _owner)
}

// PoolTogetherRegistryNewPrizeVaultIterator is returned from FilterNewPrizeVault and is used to iterate over the raw logs and unpacked data for NewPrizeVault events raised by the PoolTogetherRegistry contract.
type PoolTogetherRegistryNewPrizeVaultIterator struct {
	Event *PoolTogetherRegistryNewPrizeVault // Event containing the contract specifics and raw log

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
func (it *PoolTogetherRegistryNewPrizeVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolTogetherRegistryNewPrizeVault)
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
		it.Event = new(PoolTogetherRegistryNewPrizeVault)
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
func (it *PoolTogetherRegistryNewPrizeVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolTogetherRegistryNewPrizeVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolTogetherRegistryNewPrizeVault represents a NewPrizeVault event raised by the PoolTogetherRegistry contract.
type PoolTogetherRegistryNewPrizeVault struct {
	Vault      common.Address
	YieldVault common.Address
	PrizePool  common.Address
	Name       string
	Symbol     string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewPrizeVault is a free log retrieval operation binding the contract event 0xc56567f6a29fddd9d58e1b710306ab9af164da3f9df01fd50069df34040da911.
//
// Solidity: event NewPrizeVault(address indexed vault, address indexed yieldVault, address indexed prizePool, string name, string symbol)
func (_PoolTogetherRegistry *PoolTogetherRegistryFilterer) FilterNewPrizeVault(opts *bind.FilterOpts, vault []common.Address, yieldVault []common.Address, prizePool []common.Address) (*PoolTogetherRegistryNewPrizeVaultIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var yieldVaultRule []interface{}
	for _, yieldVaultItem := range yieldVault {
		yieldVaultRule = append(yieldVaultRule, yieldVaultItem)
	}
	var prizePoolRule []interface{}
	for _, prizePoolItem := range prizePool {
		prizePoolRule = append(prizePoolRule, prizePoolItem)
	}

	logs, sub, err := _PoolTogetherRegistry.contract.FilterLogs(opts, "NewPrizeVault", vaultRule, yieldVaultRule, prizePoolRule)
	if err != nil {
		return nil, err
	}
	return &PoolTogetherRegistryNewPrizeVaultIterator{contract: _PoolTogetherRegistry.contract, event: "NewPrizeVault", logs: logs, sub: sub}, nil
}

// WatchNewPrizeVault is a free log subscription operation binding the contract event 0xc56567f6a29fddd9d58e1b710306ab9af164da3f9df01fd50069df34040da911.
//
// Solidity: event NewPrizeVault(address indexed vault, address indexed yieldVault, address indexed prizePool, string name, string symbol)
func (_PoolTogetherRegistry *PoolTogetherRegistryFilterer) WatchNewPrizeVault(opts *bind.WatchOpts, sink chan<- *PoolTogetherRegistryNewPrizeVault, vault []common.Address, yieldVault []common.Address, prizePool []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var yieldVaultRule []interface{}
	for _, yieldVaultItem := range yieldVault {
		yieldVaultRule = append(yieldVaultRule, yieldVaultItem)
	}
	var prizePoolRule []interface{}
	for _, prizePoolItem := range prizePool {
		prizePoolRule = append(prizePoolRule, prizePoolItem)
	}

	logs, sub, err := _PoolTogetherRegistry.contract.WatchLogs(opts, "NewPrizeVault", vaultRule, yieldVaultRule, prizePoolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolTogetherRegistryNewPrizeVault)
				if err := _PoolTogetherRegistry.contract.UnpackLog(event, "NewPrizeVault", log); err != nil {
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

// ParseNewPrizeVault is a log parse operation binding the contract event 0xc56567f6a29fddd9d58e1b710306ab9af164da3f9df01fd50069df34040da911.
//
// Solidity: event NewPrizeVault(address indexed vault, address indexed yieldVault, address indexed prizePool, string name, string symbol)
func (_PoolTogetherRegistry *PoolTogetherRegistryFilterer) ParseNewPrizeVault(log types.Log) (*PoolTogetherRegistryNewPrizeVault, error) {
	event := new(PoolTogetherRegistryNewPrizeVault)
	if err := _PoolTogetherRegistry.contract.UnpackLog(event, "NewPrizeVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
