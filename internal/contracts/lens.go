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

// OracleTokenAlias is an auto generated low-level Go binding around an user-defined struct.
type OracleTokenAlias struct {
	TokenAddress      common.Address
	TokenAliasAddress common.Address
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20 *IERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20 *IERC20Session) Decimals() (uint8, error) {
	return _IERC20.Contract.Decimals(&_IERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20 *IERC20CallerSession) Decimals() (uint8, error) {
	return _IERC20.Contract.Decimals(&_IERC20.CallOpts)
}

// ManageableMetaData contains all meta data concerning the Manageable contract.
var ManageableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_managementListAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"managementList\",\"outputs\":[{\"internalType\":\"contractManagementList\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"58443a3b": "managementList()",
	},
	Bin: "0x608060405234801561001057600080fd5b5060405161012138038061012183398101604081905261002f91610054565b600080546001600160a01b0319166001600160a01b0392909216919091179055610082565b600060208284031215610065578081fd5b81516001600160a01b038116811461007b578182fd5b9392505050565b6091806100906000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c806358443a3b14602d575b600080fd5b600054603f906001600160a01b031681565b6040516001600160a01b03909116815260200160405180910390f3fea264697066735822122006368884b3ed402e02b1d4b109ac11d5b0db15b6be7dcb3f7463a1051c4cf3c664736f6c63430008020033",
}

// ManageableABI is the input ABI used to generate the binding from.
// Deprecated: Use ManageableMetaData.ABI instead.
var ManageableABI = ManageableMetaData.ABI

// Deprecated: Use ManageableMetaData.Sigs instead.
// ManageableFuncSigs maps the 4-byte function signature to its string representation.
var ManageableFuncSigs = ManageableMetaData.Sigs

// ManageableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ManageableMetaData.Bin instead.
var ManageableBin = ManageableMetaData.Bin

// DeployManageable deploys a new Ethereum contract, binding an instance of Manageable to it.
func DeployManageable(auth *bind.TransactOpts, backend bind.ContractBackend, _managementListAddress common.Address) (common.Address, *types.Transaction, *Manageable, error) {
	parsed, err := ManageableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ManageableBin), backend, _managementListAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Manageable{ManageableCaller: ManageableCaller{contract: contract}, ManageableTransactor: ManageableTransactor{contract: contract}, ManageableFilterer: ManageableFilterer{contract: contract}}, nil
}

// Manageable is an auto generated Go binding around an Ethereum contract.
type Manageable struct {
	ManageableCaller     // Read-only binding to the contract
	ManageableTransactor // Write-only binding to the contract
	ManageableFilterer   // Log filterer for contract events
}

// ManageableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ManageableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManageableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ManageableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManageableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ManageableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManageableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ManageableSession struct {
	Contract     *Manageable       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ManageableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ManageableCallerSession struct {
	Contract *ManageableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ManageableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ManageableTransactorSession struct {
	Contract     *ManageableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ManageableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ManageableRaw struct {
	Contract *Manageable // Generic contract binding to access the raw methods on
}

// ManageableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ManageableCallerRaw struct {
	Contract *ManageableCaller // Generic read-only contract binding to access the raw methods on
}

// ManageableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ManageableTransactorRaw struct {
	Contract *ManageableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewManageable creates a new instance of Manageable, bound to a specific deployed contract.
func NewManageable(address common.Address, backend bind.ContractBackend) (*Manageable, error) {
	contract, err := bindManageable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Manageable{ManageableCaller: ManageableCaller{contract: contract}, ManageableTransactor: ManageableTransactor{contract: contract}, ManageableFilterer: ManageableFilterer{contract: contract}}, nil
}

// NewManageableCaller creates a new read-only instance of Manageable, bound to a specific deployed contract.
func NewManageableCaller(address common.Address, caller bind.ContractCaller) (*ManageableCaller, error) {
	contract, err := bindManageable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ManageableCaller{contract: contract}, nil
}

// NewManageableTransactor creates a new write-only instance of Manageable, bound to a specific deployed contract.
func NewManageableTransactor(address common.Address, transactor bind.ContractTransactor) (*ManageableTransactor, error) {
	contract, err := bindManageable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ManageableTransactor{contract: contract}, nil
}

// NewManageableFilterer creates a new log filterer instance of Manageable, bound to a specific deployed contract.
func NewManageableFilterer(address common.Address, filterer bind.ContractFilterer) (*ManageableFilterer, error) {
	contract, err := bindManageable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ManageableFilterer{contract: contract}, nil
}

// bindManageable binds a generic wrapper to an already deployed contract.
func bindManageable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ManageableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Manageable *ManageableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Manageable.Contract.ManageableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Manageable *ManageableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Manageable.Contract.ManageableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Manageable *ManageableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Manageable.Contract.ManageableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Manageable *ManageableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Manageable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Manageable *ManageableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Manageable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Manageable *ManageableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Manageable.Contract.contract.Transact(opts, method, params...)
}

// ManagementList is a free data retrieval call binding the contract method 0x58443a3b.
//
// Solidity: function managementList() view returns(address)
func (_Manageable *ManageableCaller) ManagementList(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Manageable.contract.Call(opts, &out, "managementList")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ManagementList is a free data retrieval call binding the contract method 0x58443a3b.
//
// Solidity: function managementList() view returns(address)
func (_Manageable *ManageableSession) ManagementList() (common.Address, error) {
	return _Manageable.Contract.ManagementList(&_Manageable.CallOpts)
}

// ManagementList is a free data retrieval call binding the contract method 0x58443a3b.
//
// Solidity: function managementList() view returns(address)
func (_Manageable *ManageableCallerSession) ManagementList() (common.Address, error) {
	return _Manageable.Contract.ManagementList(&_Manageable.CallOpts)
}

// ManagementListMetaData contains all meta data concerning the ManagementList contract.
var ManagementListMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"accountAddress\",\"type\":\"address\"}],\"name\":\"isManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"f3ae2415": "isManager(address)",
	},
}

// ManagementListABI is the input ABI used to generate the binding from.
// Deprecated: Use ManagementListMetaData.ABI instead.
var ManagementListABI = ManagementListMetaData.ABI

// Deprecated: Use ManagementListMetaData.Sigs instead.
// ManagementListFuncSigs maps the 4-byte function signature to its string representation.
var ManagementListFuncSigs = ManagementListMetaData.Sigs

// ManagementList is an auto generated Go binding around an Ethereum contract.
type ManagementList struct {
	ManagementListCaller     // Read-only binding to the contract
	ManagementListTransactor // Write-only binding to the contract
	ManagementListFilterer   // Log filterer for contract events
}

// ManagementListCaller is an auto generated read-only Go binding around an Ethereum contract.
type ManagementListCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManagementListTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ManagementListTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManagementListFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ManagementListFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManagementListSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ManagementListSession struct {
	Contract     *ManagementList   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ManagementListCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ManagementListCallerSession struct {
	Contract *ManagementListCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ManagementListTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ManagementListTransactorSession struct {
	Contract     *ManagementListTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ManagementListRaw is an auto generated low-level Go binding around an Ethereum contract.
type ManagementListRaw struct {
	Contract *ManagementList // Generic contract binding to access the raw methods on
}

// ManagementListCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ManagementListCallerRaw struct {
	Contract *ManagementListCaller // Generic read-only contract binding to access the raw methods on
}

// ManagementListTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ManagementListTransactorRaw struct {
	Contract *ManagementListTransactor // Generic write-only contract binding to access the raw methods on
}

// NewManagementList creates a new instance of ManagementList, bound to a specific deployed contract.
func NewManagementList(address common.Address, backend bind.ContractBackend) (*ManagementList, error) {
	contract, err := bindManagementList(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ManagementList{ManagementListCaller: ManagementListCaller{contract: contract}, ManagementListTransactor: ManagementListTransactor{contract: contract}, ManagementListFilterer: ManagementListFilterer{contract: contract}}, nil
}

// NewManagementListCaller creates a new read-only instance of ManagementList, bound to a specific deployed contract.
func NewManagementListCaller(address common.Address, caller bind.ContractCaller) (*ManagementListCaller, error) {
	contract, err := bindManagementList(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ManagementListCaller{contract: contract}, nil
}

// NewManagementListTransactor creates a new write-only instance of ManagementList, bound to a specific deployed contract.
func NewManagementListTransactor(address common.Address, transactor bind.ContractTransactor) (*ManagementListTransactor, error) {
	contract, err := bindManagementList(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ManagementListTransactor{contract: contract}, nil
}

// NewManagementListFilterer creates a new log filterer instance of ManagementList, bound to a specific deployed contract.
func NewManagementListFilterer(address common.Address, filterer bind.ContractFilterer) (*ManagementListFilterer, error) {
	contract, err := bindManagementList(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ManagementListFilterer{contract: contract}, nil
}

// bindManagementList binds a generic wrapper to an already deployed contract.
func bindManagementList(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ManagementListABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ManagementList *ManagementListRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ManagementList.Contract.ManagementListCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ManagementList *ManagementListRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ManagementList.Contract.ManagementListTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ManagementList *ManagementListRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ManagementList.Contract.ManagementListTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ManagementList *ManagementListCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ManagementList.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ManagementList *ManagementListTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ManagementList.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ManagementList *ManagementListTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ManagementList.Contract.contract.Transact(opts, method, params...)
}

// IsManager is a paid mutator transaction binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address accountAddress) returns(bool)
func (_ManagementList *ManagementListTransactor) IsManager(opts *bind.TransactOpts, accountAddress common.Address) (*types.Transaction, error) {
	return _ManagementList.contract.Transact(opts, "isManager", accountAddress)
}

// IsManager is a paid mutator transaction binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address accountAddress) returns(bool)
func (_ManagementList *ManagementListSession) IsManager(accountAddress common.Address) (*types.Transaction, error) {
	return _ManagementList.Contract.IsManager(&_ManagementList.TransactOpts, accountAddress)
}

// IsManager is a paid mutator transaction binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address accountAddress) returns(bool)
func (_ManagementList *ManagementListTransactorSession) IsManager(accountAddress common.Address) (*types.Transaction, error) {
	return _ManagementList.Contract.IsManager(&_ManagementList.TransactOpts, accountAddress)
}

// OracleMetaData contains all meta data concerning the Oracle contract.
var OracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_managementListAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_usdcAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAliasAddress\",\"type\":\"address\"}],\"name\":\"TokenAliasAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenAliasRemoved\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAliasAddress\",\"type\":\"address\"}],\"name\":\"addTokenAlias\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAliasAddress\",\"type\":\"address\"}],\"internalType\":\"structOracle.TokenAlias[]\",\"name\":\"_tokenAliases\",\"type\":\"tuple[]\"}],\"name\":\"addTokenAliases\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculations\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getNormalizedValueUsdc\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceUsdc\",\"type\":\"uint256\"}],\"name\":\"getNormalizedValueUsdc\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"getPriceUsdcRecommended\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"managementList\",\"outputs\":[{\"internalType\":\"contractManagementList\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"removeTokenAlias\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"calculationAddresses\",\"type\":\"address[]\"}],\"name\":\"setCalculations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenAliases\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdcAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"6a4cb36c": "addTokenAlias(address,address)",
		"88dea52a": "addTokenAliases((address,address)[])",
		"3d71473b": "calculations()",
		"8a7f6680": "getNormalizedValueUsdc(address,uint256)",
		"d02d20aa": "getNormalizedValueUsdc(address,uint256,uint256)",
		"482ba306": "getPriceUsdcRecommended(address)",
		"58443a3b": "managementList()",
		"f3ae3077": "removeTokenAlias(address)",
		"255aacf1": "setCalculations(address[])",
		"41394ced": "tokenAliases(address)",
		"02d45457": "usdcAddress()",
	},
	Bin: "0x608060405234801561001057600080fd5b50604051610fd9380380610fd983398101604081905261002f9161007c565b600080546001600160a01b039384166001600160a01b031991821617909155600280549290931691161790556100ae565b80516001600160a01b038116811461007757600080fd5b919050565b6000806040838503121561008e578182fd5b61009783610060565b91506100a560208401610060565b90509250929050565b610f1c806100bd6000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c806358443a3b1161007157806358443a3b1461022c5780636a4cb36c1461023f57806388dea52a146102525780638a7f668014610265578063d02d20aa14610278578063f3ae30771461028b576100a9565b806302d4545714610188578063255aacf1146101b85780633d71473b146101cd57806341394ced146101e2578063482ba3061461020b575b60005b600154811015610125576000600182815481106100d957634e487b7160e01b600052603260045260246000fd5b60009182526020822001546001600160a01b03169150819036908037600080366000845afa90503d6000803e8015610110573d6000f35b5050808061011d90610e9f565b9150506100ac565b5060405162461bcd60e51b815260206004820152602c60248201527f4f7261636c653a2046616c6c6261636b2070726f7879206661696c656420746f60448201526b2072657475726e206461746160a01b60648201526084015b60405180910390fd5b60025461019b906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b6101cb6101c6366004610a41565b61029e565b005b6101d5610353565b6040516101af9190610c34565b61019b6101f036600461099a565b6003602052600090815260409020546001600160a01b031681565b61021e61021936600461099a565b6103b5565b6040519081526020016101af565b60005461019b906001600160a01b031681565b6101cb61024d3660046109b4565b6104a5565b6101cb610260366004610add565b6105ae565b61021e6102733660046109e6565b6106c9565b61021e610286366004610a0f565b6106ea565b6101cb61029936600461099a565b61080e565b6000805460405163f3ae241560e01b81523360048201526001600160a01b039091169063f3ae241590602401602060405180830381600087803b1580156102e457600080fd5b505af11580156102f8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061031c9190610ba2565b90508061033b5760405162461bcd60e51b815260040161017f90610c81565b815161034e906001906020850190610909565b505050565b606060018054806020026020016040519081016040528092919081815260200182805480156103ab57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161038d575b5050505050905090565b6001600160a01b038082166000908152600360205260408120549091168281156103dc5750805b6040516001600160a01b03821660248201526000908190309060440160408051601f198184030181529181526020820180516001600160e01b031663d61a784760e01b1790525161042d9190610bfb565b600060405180830381855afa9150503d8060008114610468576040519150601f19603f3d011682016040523d82523d6000602084013e61046d565b606091505b50915091508115610497578080602001905181019061048c9190610bc2565b9450505050506104a0565b60009450505050505b919050565b6000805460405163f3ae241560e01b81523360048201526001600160a01b039091169063f3ae241590602401602060405180830381600087803b1580156104eb57600080fd5b505af11580156104ff573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105239190610ba2565b9050806105425760405162461bcd60e51b815260040161017f90610c81565b6001600160a01b0383811660008181526003602090815260409182902080546001600160a01b031916948716948517905581519283528201929092527fa98e06f16867476bbb2764587292fb84dbd09b5502001113bc12b8359cb954f1910160405180910390a1505050565b6000805460405163f3ae241560e01b81523360048201526001600160a01b039091169063f3ae241590602401602060405180830381600087803b1580156105f457600080fd5b505af1158015610608573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061062c9190610ba2565b90508061064b5760405162461bcd60e51b815260040161017f90610c81565b60005b825181101561034e576106b783828151811061067a57634e487b7160e01b600052603260045260246000fd5b6020026020010151600001518483815181106106a657634e487b7160e01b600052603260045260246000fd5b6020026020010151602001516104a5565b806106c181610e9f565b91505061064e565b6000806106d5846103b5565b90506106e28484836106ea565b949350505050565b6000808490506000816001600160a01b031663313ce5676040518163ffffffff1660e01b815260040160206040518083038186803b15801561072b57600080fd5b505afa15801561073f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107639190610bda565b60ff169050600660008183106107845761077d8284610e88565b9050610791565b61078e8383610e88565b90505b600081156107de576107a38483610d1d565b6107ae90600a610d9b565b6107b983600a610d9b565b6107c3898b610e69565b6107cd9190610e69565b6107d79190610d35565b9050610800565b6107e983600a610d9b565b6107f3888a610e69565b6107fd9190610d35565b90505b9450505050505b9392505050565b6000805460405163f3ae241560e01b81523360048201526001600160a01b039091169063f3ae241590602401602060405180830381600087803b15801561085457600080fd5b505af1158015610868573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061088c9190610ba2565b9050806108ab5760405162461bcd60e51b815260040161017f90610c81565b6001600160a01b03821660008181526003602090815260409182902080546001600160a01b031916905590519182527f1262aa3d0f0c652f1a55888d946c631f1a234f587d9e3e9f8ee39974b628fe7f910160405180910390a15050565b82805482825590600052602060002090810192821561095e579160200282015b8281111561095e57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190610929565b5061096a92915061096e565b5090565b5b8082111561096a576000815560010161096f565b80356001600160a01b03811681146104a057600080fd5b6000602082840312156109ab578081fd5b61080782610983565b600080604083850312156109c6578081fd5b6109cf83610983565b91506109dd60208401610983565b90509250929050565b600080604083850312156109f8578182fd5b610a0183610983565b946020939093013593505050565b600080600060608486031215610a23578081fd5b610a2c84610983565b95602085013595506040909401359392505050565b60006020808385031215610a53578182fd5b823567ffffffffffffffff811115610a69578283fd5b8301601f81018513610a79578283fd5b8035610a8c610a8782610cf9565b610cc8565b8181528381019083850185840285018601891015610aa8578687fd5b8694505b83851015610ad157610abd81610983565b835260019490940193918501918501610aac565b50979650505050505050565b60006020808385031215610aef578182fd5b823567ffffffffffffffff811115610b05578283fd5b8301601f81018513610b15578283fd5b8035610b23610a8782610cf9565b818152838101908385016040808502860187018a1015610b41578788fd5b8795505b84861015610b945780828b031215610b5b578788fd5b610b6481610cc8565b610b6d83610983565b8152610b7a888401610983565b818901528452600195909501949286019290810190610b45565b509098975050505050505050565b600060208284031215610bb3578081fd5b81518015158114610807578182fd5b600060208284031215610bd3578081fd5b5051919050565b600060208284031215610beb578081fd5b815160ff81168114610807578182fd5b60008251815b81811015610c1b5760208186018101518583015201610c01565b81811115610c295782828501525b509190910192915050565b6020808252825182820181905260009190848201906040850190845b81811015610c755783516001600160a01b031683529284019291840191600101610c50565b50909695505050505050565b60208082526027908201527f4d616e6167656d656e744c6973743a2063616c6c6572206973206e6f7420612060408201526636b0b730b3b2b960c91b606082015260800190565b604051601f8201601f1916810167ffffffffffffffff81118282101715610cf157610cf1610ed0565b604052919050565b600067ffffffffffffffff821115610d1357610d13610ed0565b5060209081020190565b60008219821115610d3057610d30610eba565b500190565b600082610d5057634e487b7160e01b81526012600452602481fd5b500490565b80825b6001808611610d675750610d92565b818704821115610d7957610d79610eba565b80861615610d8657918102915b9490941c938002610d58565b94509492505050565b60006108076000198484600082610db457506001610807565b81610dc157506000610807565b8160018114610dd75760028114610de157610e0e565b6001915050610807565b60ff841115610df257610df2610eba565b6001841b915084821115610e0857610e08610eba565b50610807565b5060208310610133831016604e8410600b8410161715610e41575081810a83811115610e3c57610e3c610eba565b610807565b610e4e8484846001610d55565b808604821115610e6057610e60610eba565b02949350505050565b6000816000190483118215151615610e8357610e83610eba565b500290565b600082821015610e9a57610e9a610eba565b500390565b6000600019821415610eb357610eb3610eba565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fdfea2646970667358221220df720f279bbfba533cbf13a9f9912d7b8c1c2d5015c9cb2be17278944025dab364736f6c63430008020033",
}

// OracleABI is the input ABI used to generate the binding from.
// Deprecated: Use OracleMetaData.ABI instead.
var OracleABI = OracleMetaData.ABI

// Deprecated: Use OracleMetaData.Sigs instead.
// OracleFuncSigs maps the 4-byte function signature to its string representation.
var OracleFuncSigs = OracleMetaData.Sigs

// OracleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OracleMetaData.Bin instead.
var OracleBin = OracleMetaData.Bin

// DeployOracle deploys a new Ethereum contract, binding an instance of Oracle to it.
func DeployOracle(auth *bind.TransactOpts, backend bind.ContractBackend, _managementListAddress common.Address, _usdcAddress common.Address) (common.Address, *types.Transaction, *Oracle, error) {
	parsed, err := OracleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OracleBin), backend, _managementListAddress, _usdcAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// Oracle is an auto generated Go binding around an Ethereum contract.
type Oracle struct {
	OracleCaller     // Read-only binding to the contract
	OracleTransactor // Write-only binding to the contract
	OracleFilterer   // Log filterer for contract events
}

// OracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleSession struct {
	Contract     *Oracle           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleCallerSession struct {
	Contract *OracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleTransactorSession struct {
	Contract     *OracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleRaw struct {
	Contract *Oracle // Generic contract binding to access the raw methods on
}

// OracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleCallerRaw struct {
	Contract *OracleCaller // Generic read-only contract binding to access the raw methods on
}

// OracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleTransactorRaw struct {
	Contract *OracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracle creates a new instance of Oracle, bound to a specific deployed contract.
func NewOracle(address common.Address, backend bind.ContractBackend) (*Oracle, error) {
	contract, err := bindOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// NewOracleCaller creates a new read-only instance of Oracle, bound to a specific deployed contract.
func NewOracleCaller(address common.Address, caller bind.ContractCaller) (*OracleCaller, error) {
	contract, err := bindOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleCaller{contract: contract}, nil
}

// NewOracleTransactor creates a new write-only instance of Oracle, bound to a specific deployed contract.
func NewOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleTransactor, error) {
	contract, err := bindOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleTransactor{contract: contract}, nil
}

// NewOracleFilterer creates a new log filterer instance of Oracle, bound to a specific deployed contract.
func NewOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleFilterer, error) {
	contract, err := bindOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleFilterer{contract: contract}, nil
}

// bindOracle binds a generic wrapper to an already deployed contract.
func bindOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.OracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transact(opts, method, params...)
}

// Calculations is a free data retrieval call binding the contract method 0x3d71473b.
//
// Solidity: function calculations() view returns(address[])
func (_Oracle *OracleCaller) Calculations(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "calculations")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// Calculations is a free data retrieval call binding the contract method 0x3d71473b.
//
// Solidity: function calculations() view returns(address[])
func (_Oracle *OracleSession) Calculations() ([]common.Address, error) {
	return _Oracle.Contract.Calculations(&_Oracle.CallOpts)
}

// Calculations is a free data retrieval call binding the contract method 0x3d71473b.
//
// Solidity: function calculations() view returns(address[])
func (_Oracle *OracleCallerSession) Calculations() ([]common.Address, error) {
	return _Oracle.Contract.Calculations(&_Oracle.CallOpts)
}

// GetNormalizedValueUsdc is a free data retrieval call binding the contract method 0x8a7f6680.
//
// Solidity: function getNormalizedValueUsdc(address tokenAddress, uint256 amount) view returns(uint256)
func (_Oracle *OracleCaller) GetNormalizedValueUsdc(opts *bind.CallOpts, tokenAddress common.Address, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getNormalizedValueUsdc", tokenAddress, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNormalizedValueUsdc is a free data retrieval call binding the contract method 0x8a7f6680.
//
// Solidity: function getNormalizedValueUsdc(address tokenAddress, uint256 amount) view returns(uint256)
func (_Oracle *OracleSession) GetNormalizedValueUsdc(tokenAddress common.Address, amount *big.Int) (*big.Int, error) {
	return _Oracle.Contract.GetNormalizedValueUsdc(&_Oracle.CallOpts, tokenAddress, amount)
}

// GetNormalizedValueUsdc is a free data retrieval call binding the contract method 0x8a7f6680.
//
// Solidity: function getNormalizedValueUsdc(address tokenAddress, uint256 amount) view returns(uint256)
func (_Oracle *OracleCallerSession) GetNormalizedValueUsdc(tokenAddress common.Address, amount *big.Int) (*big.Int, error) {
	return _Oracle.Contract.GetNormalizedValueUsdc(&_Oracle.CallOpts, tokenAddress, amount)
}

// GetNormalizedValueUsdc0 is a free data retrieval call binding the contract method 0xd02d20aa.
//
// Solidity: function getNormalizedValueUsdc(address tokenAddress, uint256 amount, uint256 priceUsdc) view returns(uint256)
func (_Oracle *OracleCaller) GetNormalizedValueUsdc0(opts *bind.CallOpts, tokenAddress common.Address, amount *big.Int, priceUsdc *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getNormalizedValueUsdc0", tokenAddress, amount, priceUsdc)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNormalizedValueUsdc0 is a free data retrieval call binding the contract method 0xd02d20aa.
//
// Solidity: function getNormalizedValueUsdc(address tokenAddress, uint256 amount, uint256 priceUsdc) view returns(uint256)
func (_Oracle *OracleSession) GetNormalizedValueUsdc0(tokenAddress common.Address, amount *big.Int, priceUsdc *big.Int) (*big.Int, error) {
	return _Oracle.Contract.GetNormalizedValueUsdc0(&_Oracle.CallOpts, tokenAddress, amount, priceUsdc)
}

// GetNormalizedValueUsdc0 is a free data retrieval call binding the contract method 0xd02d20aa.
//
// Solidity: function getNormalizedValueUsdc(address tokenAddress, uint256 amount, uint256 priceUsdc) view returns(uint256)
func (_Oracle *OracleCallerSession) GetNormalizedValueUsdc0(tokenAddress common.Address, amount *big.Int, priceUsdc *big.Int) (*big.Int, error) {
	return _Oracle.Contract.GetNormalizedValueUsdc0(&_Oracle.CallOpts, tokenAddress, amount, priceUsdc)
}

// GetPriceUsdcRecommended is a free data retrieval call binding the contract method 0x482ba306.
//
// Solidity: function getPriceUsdcRecommended(address tokenAddress) view returns(uint256)
func (_Oracle *OracleCaller) GetPriceUsdcRecommended(opts *bind.CallOpts, tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getPriceUsdcRecommended", tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPriceUsdcRecommended is a free data retrieval call binding the contract method 0x482ba306.
//
// Solidity: function getPriceUsdcRecommended(address tokenAddress) view returns(uint256)
func (_Oracle *OracleSession) GetPriceUsdcRecommended(tokenAddress common.Address) (*big.Int, error) {
	return _Oracle.Contract.GetPriceUsdcRecommended(&_Oracle.CallOpts, tokenAddress)
}

// GetPriceUsdcRecommended is a free data retrieval call binding the contract method 0x482ba306.
//
// Solidity: function getPriceUsdcRecommended(address tokenAddress) view returns(uint256)
func (_Oracle *OracleCallerSession) GetPriceUsdcRecommended(tokenAddress common.Address) (*big.Int, error) {
	return _Oracle.Contract.GetPriceUsdcRecommended(&_Oracle.CallOpts, tokenAddress)
}

// ManagementList is a free data retrieval call binding the contract method 0x58443a3b.
//
// Solidity: function managementList() view returns(address)
func (_Oracle *OracleCaller) ManagementList(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "managementList")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ManagementList is a free data retrieval call binding the contract method 0x58443a3b.
//
// Solidity: function managementList() view returns(address)
func (_Oracle *OracleSession) ManagementList() (common.Address, error) {
	return _Oracle.Contract.ManagementList(&_Oracle.CallOpts)
}

// ManagementList is a free data retrieval call binding the contract method 0x58443a3b.
//
// Solidity: function managementList() view returns(address)
func (_Oracle *OracleCallerSession) ManagementList() (common.Address, error) {
	return _Oracle.Contract.ManagementList(&_Oracle.CallOpts)
}

// TokenAliases is a free data retrieval call binding the contract method 0x41394ced.
//
// Solidity: function tokenAliases(address ) view returns(address)
func (_Oracle *OracleCaller) TokenAliases(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "tokenAliases", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenAliases is a free data retrieval call binding the contract method 0x41394ced.
//
// Solidity: function tokenAliases(address ) view returns(address)
func (_Oracle *OracleSession) TokenAliases(arg0 common.Address) (common.Address, error) {
	return _Oracle.Contract.TokenAliases(&_Oracle.CallOpts, arg0)
}

// TokenAliases is a free data retrieval call binding the contract method 0x41394ced.
//
// Solidity: function tokenAliases(address ) view returns(address)
func (_Oracle *OracleCallerSession) TokenAliases(arg0 common.Address) (common.Address, error) {
	return _Oracle.Contract.TokenAliases(&_Oracle.CallOpts, arg0)
}

// UsdcAddress is a free data retrieval call binding the contract method 0x02d45457.
//
// Solidity: function usdcAddress() view returns(address)
func (_Oracle *OracleCaller) UsdcAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "usdcAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UsdcAddress is a free data retrieval call binding the contract method 0x02d45457.
//
// Solidity: function usdcAddress() view returns(address)
func (_Oracle *OracleSession) UsdcAddress() (common.Address, error) {
	return _Oracle.Contract.UsdcAddress(&_Oracle.CallOpts)
}

// UsdcAddress is a free data retrieval call binding the contract method 0x02d45457.
//
// Solidity: function usdcAddress() view returns(address)
func (_Oracle *OracleCallerSession) UsdcAddress() (common.Address, error) {
	return _Oracle.Contract.UsdcAddress(&_Oracle.CallOpts)
}

// AddTokenAlias is a paid mutator transaction binding the contract method 0x6a4cb36c.
//
// Solidity: function addTokenAlias(address tokenAddress, address tokenAliasAddress) returns()
func (_Oracle *OracleTransactor) AddTokenAlias(opts *bind.TransactOpts, tokenAddress common.Address, tokenAliasAddress common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "addTokenAlias", tokenAddress, tokenAliasAddress)
}

// AddTokenAlias is a paid mutator transaction binding the contract method 0x6a4cb36c.
//
// Solidity: function addTokenAlias(address tokenAddress, address tokenAliasAddress) returns()
func (_Oracle *OracleSession) AddTokenAlias(tokenAddress common.Address, tokenAliasAddress common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.AddTokenAlias(&_Oracle.TransactOpts, tokenAddress, tokenAliasAddress)
}

// AddTokenAlias is a paid mutator transaction binding the contract method 0x6a4cb36c.
//
// Solidity: function addTokenAlias(address tokenAddress, address tokenAliasAddress) returns()
func (_Oracle *OracleTransactorSession) AddTokenAlias(tokenAddress common.Address, tokenAliasAddress common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.AddTokenAlias(&_Oracle.TransactOpts, tokenAddress, tokenAliasAddress)
}

// AddTokenAliases is a paid mutator transaction binding the contract method 0x88dea52a.
//
// Solidity: function addTokenAliases((address,address)[] _tokenAliases) returns()
func (_Oracle *OracleTransactor) AddTokenAliases(opts *bind.TransactOpts, _tokenAliases []OracleTokenAlias) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "addTokenAliases", _tokenAliases)
}

// AddTokenAliases is a paid mutator transaction binding the contract method 0x88dea52a.
//
// Solidity: function addTokenAliases((address,address)[] _tokenAliases) returns()
func (_Oracle *OracleSession) AddTokenAliases(_tokenAliases []OracleTokenAlias) (*types.Transaction, error) {
	return _Oracle.Contract.AddTokenAliases(&_Oracle.TransactOpts, _tokenAliases)
}

// AddTokenAliases is a paid mutator transaction binding the contract method 0x88dea52a.
//
// Solidity: function addTokenAliases((address,address)[] _tokenAliases) returns()
func (_Oracle *OracleTransactorSession) AddTokenAliases(_tokenAliases []OracleTokenAlias) (*types.Transaction, error) {
	return _Oracle.Contract.AddTokenAliases(&_Oracle.TransactOpts, _tokenAliases)
}

// RemoveTokenAlias is a paid mutator transaction binding the contract method 0xf3ae3077.
//
// Solidity: function removeTokenAlias(address tokenAddress) returns()
func (_Oracle *OracleTransactor) RemoveTokenAlias(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "removeTokenAlias", tokenAddress)
}

// RemoveTokenAlias is a paid mutator transaction binding the contract method 0xf3ae3077.
//
// Solidity: function removeTokenAlias(address tokenAddress) returns()
func (_Oracle *OracleSession) RemoveTokenAlias(tokenAddress common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.RemoveTokenAlias(&_Oracle.TransactOpts, tokenAddress)
}

// RemoveTokenAlias is a paid mutator transaction binding the contract method 0xf3ae3077.
//
// Solidity: function removeTokenAlias(address tokenAddress) returns()
func (_Oracle *OracleTransactorSession) RemoveTokenAlias(tokenAddress common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.RemoveTokenAlias(&_Oracle.TransactOpts, tokenAddress)
}

// SetCalculations is a paid mutator transaction binding the contract method 0x255aacf1.
//
// Solidity: function setCalculations(address[] calculationAddresses) returns()
func (_Oracle *OracleTransactor) SetCalculations(opts *bind.TransactOpts, calculationAddresses []common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "setCalculations", calculationAddresses)
}

// SetCalculations is a paid mutator transaction binding the contract method 0x255aacf1.
//
// Solidity: function setCalculations(address[] calculationAddresses) returns()
func (_Oracle *OracleSession) SetCalculations(calculationAddresses []common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.SetCalculations(&_Oracle.TransactOpts, calculationAddresses)
}

// SetCalculations is a paid mutator transaction binding the contract method 0x255aacf1.
//
// Solidity: function setCalculations(address[] calculationAddresses) returns()
func (_Oracle *OracleTransactorSession) SetCalculations(calculationAddresses []common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.SetCalculations(&_Oracle.TransactOpts, calculationAddresses)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Oracle *OracleTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Oracle.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Oracle *OracleSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Oracle.Contract.Fallback(&_Oracle.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Oracle *OracleTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Oracle.Contract.Fallback(&_Oracle.TransactOpts, calldata)
}

// OracleTokenAliasAddedIterator is returned from FilterTokenAliasAdded and is used to iterate over the raw logs and unpacked data for TokenAliasAdded events raised by the Oracle contract.
type OracleTokenAliasAddedIterator struct {
	Event *OracleTokenAliasAdded // Event containing the contract specifics and raw log

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
func (it *OracleTokenAliasAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleTokenAliasAdded)
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
		it.Event = new(OracleTokenAliasAdded)
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
func (it *OracleTokenAliasAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleTokenAliasAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleTokenAliasAdded represents a TokenAliasAdded event raised by the Oracle contract.
type OracleTokenAliasAdded struct {
	TokenAddress      common.Address
	TokenAliasAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTokenAliasAdded is a free log retrieval operation binding the contract event 0xa98e06f16867476bbb2764587292fb84dbd09b5502001113bc12b8359cb954f1.
//
// Solidity: event TokenAliasAdded(address tokenAddress, address tokenAliasAddress)
func (_Oracle *OracleFilterer) FilterTokenAliasAdded(opts *bind.FilterOpts) (*OracleTokenAliasAddedIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "TokenAliasAdded")
	if err != nil {
		return nil, err
	}
	return &OracleTokenAliasAddedIterator{contract: _Oracle.contract, event: "TokenAliasAdded", logs: logs, sub: sub}, nil
}

// WatchTokenAliasAdded is a free log subscription operation binding the contract event 0xa98e06f16867476bbb2764587292fb84dbd09b5502001113bc12b8359cb954f1.
//
// Solidity: event TokenAliasAdded(address tokenAddress, address tokenAliasAddress)
func (_Oracle *OracleFilterer) WatchTokenAliasAdded(opts *bind.WatchOpts, sink chan<- *OracleTokenAliasAdded) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "TokenAliasAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleTokenAliasAdded)
				if err := _Oracle.contract.UnpackLog(event, "TokenAliasAdded", log); err != nil {
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

// ParseTokenAliasAdded is a log parse operation binding the contract event 0xa98e06f16867476bbb2764587292fb84dbd09b5502001113bc12b8359cb954f1.
//
// Solidity: event TokenAliasAdded(address tokenAddress, address tokenAliasAddress)
func (_Oracle *OracleFilterer) ParseTokenAliasAdded(log types.Log) (*OracleTokenAliasAdded, error) {
	event := new(OracleTokenAliasAdded)
	if err := _Oracle.contract.UnpackLog(event, "TokenAliasAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleTokenAliasRemovedIterator is returned from FilterTokenAliasRemoved and is used to iterate over the raw logs and unpacked data for TokenAliasRemoved events raised by the Oracle contract.
type OracleTokenAliasRemovedIterator struct {
	Event *OracleTokenAliasRemoved // Event containing the contract specifics and raw log

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
func (it *OracleTokenAliasRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleTokenAliasRemoved)
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
		it.Event = new(OracleTokenAliasRemoved)
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
func (it *OracleTokenAliasRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleTokenAliasRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleTokenAliasRemoved represents a TokenAliasRemoved event raised by the Oracle contract.
type OracleTokenAliasRemoved struct {
	TokenAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenAliasRemoved is a free log retrieval operation binding the contract event 0x1262aa3d0f0c652f1a55888d946c631f1a234f587d9e3e9f8ee39974b628fe7f.
//
// Solidity: event TokenAliasRemoved(address tokenAddress)
func (_Oracle *OracleFilterer) FilterTokenAliasRemoved(opts *bind.FilterOpts) (*OracleTokenAliasRemovedIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "TokenAliasRemoved")
	if err != nil {
		return nil, err
	}
	return &OracleTokenAliasRemovedIterator{contract: _Oracle.contract, event: "TokenAliasRemoved", logs: logs, sub: sub}, nil
}

// WatchTokenAliasRemoved is a free log subscription operation binding the contract event 0x1262aa3d0f0c652f1a55888d946c631f1a234f587d9e3e9f8ee39974b628fe7f.
//
// Solidity: event TokenAliasRemoved(address tokenAddress)
func (_Oracle *OracleFilterer) WatchTokenAliasRemoved(opts *bind.WatchOpts, sink chan<- *OracleTokenAliasRemoved) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "TokenAliasRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleTokenAliasRemoved)
				if err := _Oracle.contract.UnpackLog(event, "TokenAliasRemoved", log); err != nil {
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

// ParseTokenAliasRemoved is a log parse operation binding the contract event 0x1262aa3d0f0c652f1a55888d946c631f1a234f587d9e3e9f8ee39974b628fe7f.
//
// Solidity: event TokenAliasRemoved(address tokenAddress)
func (_Oracle *OracleFilterer) ParseTokenAliasRemoved(log types.Log) (*OracleTokenAliasRemoved, error) {
	event := new(OracleTokenAliasRemoved)
	if err := _Oracle.contract.UnpackLog(event, "TokenAliasRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
