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

// Multicall2Call is an auto generated low-level Go binding around an user-defined struct.
type Multicall2Call struct {
	Target   common.Address
	CallData []byte
}

// Multicall2Result is an auto generated low-level Go binding around an user-defined struct.
type Multicall2Result struct {
	Success    bool
	ReturnData []byte
}

// Multicall2MetaData contains all meta data concerning the Multicall2 contract.
var Multicall2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall2.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"aggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"returnData\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall2.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"blockAndAggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall2.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockCoinbase\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockDifficulty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gaslimit\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getEthBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"requireSuccess\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall2.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"tryAggregate\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall2.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"requireSuccess\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall2.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"tryBlockAndAggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall2.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"252dba42": "aggregate((address,bytes)[])",
		"c3077fa9": "blockAndAggregate((address,bytes)[])",
		"ee82ac5e": "getBlockHash(uint256)",
		"42cbb15c": "getBlockNumber()",
		"a8b0574e": "getCurrentBlockCoinbase()",
		"72425d9d": "getCurrentBlockDifficulty()",
		"86d516e8": "getCurrentBlockGasLimit()",
		"0f28c97d": "getCurrentBlockTimestamp()",
		"4d2301cc": "getEthBalance(address)",
		"27e86d6e": "getLastBlockHash()",
		"bce38bd7": "tryAggregate(bool,(address,bytes)[])",
		"399542e9": "tryBlockAndAggregate(bool,(address,bytes)[])",
	},
	Bin: "0x608060405234801561001057600080fd5b506109d3806100206000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c806372425d9d1161007157806372425d9d1461013a57806386d516e814610140578063a8b0574e14610146578063bce38bd714610154578063c3077fa914610174578063ee82ac5e14610187576100b4565b80630f28c97d146100b9578063252dba42146100ce57806327e86d6e146100ef578063399542e9146100f757806342cbb15c146101195780634d2301cc1461011f575b600080fd5b425b6040519081526020015b60405180910390f35b6100e16100dc3660046106e2565b610199565b6040516100c592919061084e565b6100bb610359565b61010a61010536600461071d565b61036c565b6040516100c5939291906108b6565b436100bb565b6100bb61012d3660046106c1565b6001600160a01b03163190565b446100bb565b456100bb565b6040514181526020016100c5565b61016761016236600461071d565b610384565b6040516100c5919061083b565b61010a6101823660046106e2565b610576565b6100bb61019536600461076f565b4090565b8051439060609067ffffffffffffffff8111156101c657634e487b7160e01b600052604160045260246000fd5b6040519080825280602002602001820160405280156101f957816020015b60608152602001906001900390816101e45790505b50905060005b83518110156103535760008085838151811061022b57634e487b7160e01b600052603260045260246000fd5b6020026020010151600001516001600160a01b031686848151811061026057634e487b7160e01b600052603260045260246000fd5b602002602001015160200151604051610279919061081f565b6000604051808303816000865af19150503d80600081146102b6576040519150601f19603f3d011682016040523d82523d6000602084013e6102bb565b606091505b5091509150816103125760405162461bcd60e51b815260206004820181905260248201527f4d756c746963616c6c206167677265676174653a2063616c6c206661696c656460448201526064015b60405180910390fd5b8084848151811061033357634e487b7160e01b600052603260045260246000fd5b60200260200101819052505050808061034b90610956565b9150506101ff565b50915091565b600061036660014361090f565b40905090565b438040606061037b8585610384565b90509250925092565b6060815167ffffffffffffffff8111156103ae57634e487b7160e01b600052604160045260246000fd5b6040519080825280602002602001820160405280156103f457816020015b6040805180820190915260008152606060208201528152602001906001900390816103cc5790505b50905060005b825181101561056f5760008084838151811061042657634e487b7160e01b600052603260045260246000fd5b6020026020010151600001516001600160a01b031685848151811061045b57634e487b7160e01b600052603260045260246000fd5b602002602001015160200151604051610474919061081f565b6000604051808303816000865af19150503d80600081146104b1576040519150601f19603f3d011682016040523d82523d6000602084013e6104b6565b606091505b5091509150851561051857816105185760405162461bcd60e51b815260206004820152602160248201527f4d756c746963616c6c32206167677265676174653a2063616c6c206661696c656044820152601960fa1b6064820152608401610309565b604051806040016040528083151581526020018281525084848151811061054f57634e487b7160e01b600052603260045260246000fd5b60200260200101819052505050808061056790610956565b9150506103fa565b5092915050565b600080606061058660018561036c565b9196909550909350915050565b80356001600160a01b03811681146105aa57600080fd5b919050565b600082601f8301126105bf578081fd5b8135602067ffffffffffffffff808311156105dc576105dc610987565b6105e982838502016108de565b83815282810190868401865b868110156106b357813589016040601f198181848f03011215610616578a8bfd5b61061f826108de565b61062a8a8501610593565b8152828401358981111561063c578c8dfd5b8085019450508d603f850112610650578b8cfd5b898401358981111561066457610664610987565b6106748b84601f840116016108de565b92508083528e84828701011115610689578c8dfd5b808486018c85013782018a018c9052808a01919091528652505092850192908501906001016105f5565b509098975050505050505050565b6000602082840312156106d2578081fd5b6106db82610593565b9392505050565b6000602082840312156106f3578081fd5b813567ffffffffffffffff811115610709578182fd5b610715848285016105af565b949350505050565b6000806040838503121561072f578081fd5b8235801515811461073e578182fd5b9150602083013567ffffffffffffffff811115610759578182fd5b610765858286016105af565b9150509250929050565b600060208284031215610780578081fd5b5035919050565b60008282518085526020808601955080818302840101818601855b848110156107e657858303601f19018952815180511515845284015160408585018190526107d2818601836107f3565b9a86019a94505050908301906001016107a2565b5090979650505050505050565b6000815180845261080b816020860160208601610926565b601f01601f19169290920160200192915050565b60008251610831818460208701610926565b9190910192915050565b6000602082526106db6020830184610787565b600060408201848352602060408185015281855180845260608601915060608382028701019350828701855b828110156108a857605f198887030184526108968683516107f3565b9550928401929084019060010161087a565b509398975050505050505050565b6000848252836020830152606060408301526108d56060830184610787565b95945050505050565b604051601f8201601f1916810167ffffffffffffffff8111828210171561090757610907610987565b604052919050565b60008282101561092157610921610971565b500390565b60005b83811015610941578181015183820152602001610929565b83811115610950576000848401525b50505050565b600060001982141561096a5761096a610971565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fdfea26469706673582212204075e7d2d0030b422fccd929ae796adedaac6482d0124593fbeb8aff91e2c0eb64736f6c63430008020033",
}

// Multicall2ABI is the input ABI used to generate the binding from.
// Deprecated: Use Multicall2MetaData.ABI instead.
var Multicall2ABI = Multicall2MetaData.ABI

// Deprecated: Use Multicall2MetaData.Sigs instead.
// Multicall2FuncSigs maps the 4-byte function signature to its string representation.
var Multicall2FuncSigs = Multicall2MetaData.Sigs

// Multicall2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Multicall2MetaData.Bin instead.
var Multicall2Bin = Multicall2MetaData.Bin

// DeployMulticall2 deploys a new Ethereum contract, binding an instance of Multicall2 to it.
func DeployMulticall2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Multicall2, error) {
	parsed, err := Multicall2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Multicall2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Multicall2{Multicall2Caller: Multicall2Caller{contract: contract}, Multicall2Transactor: Multicall2Transactor{contract: contract}, Multicall2Filterer: Multicall2Filterer{contract: contract}}, nil
}

// Multicall2 is an auto generated Go binding around an Ethereum contract.
type Multicall2 struct {
	Multicall2Caller     // Read-only binding to the contract
	Multicall2Transactor // Write-only binding to the contract
	Multicall2Filterer   // Log filterer for contract events
}

// Multicall2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Multicall2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Multicall2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Multicall2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Multicall2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Multicall2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Multicall2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Multicall2Session struct {
	Contract     *Multicall2       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Multicall2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Multicall2CallerSession struct {
	Contract *Multicall2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// Multicall2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Multicall2TransactorSession struct {
	Contract     *Multicall2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// Multicall2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Multicall2Raw struct {
	Contract *Multicall2 // Generic contract binding to access the raw methods on
}

// Multicall2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Multicall2CallerRaw struct {
	Contract *Multicall2Caller // Generic read-only contract binding to access the raw methods on
}

// Multicall2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Multicall2TransactorRaw struct {
	Contract *Multicall2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMulticall2 creates a new instance of Multicall2, bound to a specific deployed contract.
func NewMulticall2(address common.Address, backend bind.ContractBackend) (*Multicall2, error) {
	contract, err := bindMulticall2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Multicall2{Multicall2Caller: Multicall2Caller{contract: contract}, Multicall2Transactor: Multicall2Transactor{contract: contract}, Multicall2Filterer: Multicall2Filterer{contract: contract}}, nil
}

// NewMulticall2Caller creates a new read-only instance of Multicall2, bound to a specific deployed contract.
func NewMulticall2Caller(address common.Address, caller bind.ContractCaller) (*Multicall2Caller, error) {
	contract, err := bindMulticall2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Multicall2Caller{contract: contract}, nil
}

// NewMulticall2Transactor creates a new write-only instance of Multicall2, bound to a specific deployed contract.
func NewMulticall2Transactor(address common.Address, transactor bind.ContractTransactor) (*Multicall2Transactor, error) {
	contract, err := bindMulticall2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Multicall2Transactor{contract: contract}, nil
}

// NewMulticall2Filterer creates a new log filterer instance of Multicall2, bound to a specific deployed contract.
func NewMulticall2Filterer(address common.Address, filterer bind.ContractFilterer) (*Multicall2Filterer, error) {
	contract, err := bindMulticall2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Multicall2Filterer{contract: contract}, nil
}

// bindMulticall2 binds a generic wrapper to an already deployed contract.
func bindMulticall2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Multicall2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multicall2 *Multicall2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multicall2.Contract.Multicall2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multicall2 *Multicall2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multicall2.Contract.Multicall2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multicall2 *Multicall2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multicall2.Contract.Multicall2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multicall2 *Multicall2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multicall2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multicall2 *Multicall2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multicall2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multicall2 *Multicall2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multicall2.Contract.contract.Transact(opts, method, params...)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_Multicall2 *Multicall2Caller) GetBlockHash(opts *bind.CallOpts, blockNumber *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Multicall2.contract.Call(opts, &out, "getBlockHash", blockNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_Multicall2 *Multicall2Session) GetBlockHash(blockNumber *big.Int) ([32]byte, error) {
	return _Multicall2.Contract.GetBlockHash(&_Multicall2.CallOpts, blockNumber)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_Multicall2 *Multicall2CallerSession) GetBlockHash(blockNumber *big.Int) ([32]byte, error) {
	return _Multicall2.Contract.GetBlockHash(&_Multicall2.CallOpts, blockNumber)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 blockNumber)
func (_Multicall2 *Multicall2Caller) GetBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Multicall2.contract.Call(opts, &out, "getBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 blockNumber)
func (_Multicall2 *Multicall2Session) GetBlockNumber() (*big.Int, error) {
	return _Multicall2.Contract.GetBlockNumber(&_Multicall2.CallOpts)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 blockNumber)
func (_Multicall2 *Multicall2CallerSession) GetBlockNumber() (*big.Int, error) {
	return _Multicall2.Contract.GetBlockNumber(&_Multicall2.CallOpts)
}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_Multicall2 *Multicall2Caller) GetCurrentBlockCoinbase(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Multicall2.contract.Call(opts, &out, "getCurrentBlockCoinbase")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_Multicall2 *Multicall2Session) GetCurrentBlockCoinbase() (common.Address, error) {
	return _Multicall2.Contract.GetCurrentBlockCoinbase(&_Multicall2.CallOpts)
}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_Multicall2 *Multicall2CallerSession) GetCurrentBlockCoinbase() (common.Address, error) {
	return _Multicall2.Contract.GetCurrentBlockCoinbase(&_Multicall2.CallOpts)
}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_Multicall2 *Multicall2Caller) GetCurrentBlockDifficulty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Multicall2.contract.Call(opts, &out, "getCurrentBlockDifficulty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_Multicall2 *Multicall2Session) GetCurrentBlockDifficulty() (*big.Int, error) {
	return _Multicall2.Contract.GetCurrentBlockDifficulty(&_Multicall2.CallOpts)
}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_Multicall2 *Multicall2CallerSession) GetCurrentBlockDifficulty() (*big.Int, error) {
	return _Multicall2.Contract.GetCurrentBlockDifficulty(&_Multicall2.CallOpts)
}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_Multicall2 *Multicall2Caller) GetCurrentBlockGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Multicall2.contract.Call(opts, &out, "getCurrentBlockGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_Multicall2 *Multicall2Session) GetCurrentBlockGasLimit() (*big.Int, error) {
	return _Multicall2.Contract.GetCurrentBlockGasLimit(&_Multicall2.CallOpts)
}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_Multicall2 *Multicall2CallerSession) GetCurrentBlockGasLimit() (*big.Int, error) {
	return _Multicall2.Contract.GetCurrentBlockGasLimit(&_Multicall2.CallOpts)
}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_Multicall2 *Multicall2Caller) GetCurrentBlockTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Multicall2.contract.Call(opts, &out, "getCurrentBlockTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_Multicall2 *Multicall2Session) GetCurrentBlockTimestamp() (*big.Int, error) {
	return _Multicall2.Contract.GetCurrentBlockTimestamp(&_Multicall2.CallOpts)
}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_Multicall2 *Multicall2CallerSession) GetCurrentBlockTimestamp() (*big.Int, error) {
	return _Multicall2.Contract.GetCurrentBlockTimestamp(&_Multicall2.CallOpts)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_Multicall2 *Multicall2Caller) GetEthBalance(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Multicall2.contract.Call(opts, &out, "getEthBalance", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_Multicall2 *Multicall2Session) GetEthBalance(addr common.Address) (*big.Int, error) {
	return _Multicall2.Contract.GetEthBalance(&_Multicall2.CallOpts, addr)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_Multicall2 *Multicall2CallerSession) GetEthBalance(addr common.Address) (*big.Int, error) {
	return _Multicall2.Contract.GetEthBalance(&_Multicall2.CallOpts, addr)
}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_Multicall2 *Multicall2Caller) GetLastBlockHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Multicall2.contract.Call(opts, &out, "getLastBlockHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_Multicall2 *Multicall2Session) GetLastBlockHash() ([32]byte, error) {
	return _Multicall2.Contract.GetLastBlockHash(&_Multicall2.CallOpts)
}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_Multicall2 *Multicall2CallerSession) GetLastBlockHash() ([32]byte, error) {
	return _Multicall2.Contract.GetLastBlockHash(&_Multicall2.CallOpts)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes[] returnData)
func (_Multicall2 *Multicall2Transactor) Aggregate(opts *bind.TransactOpts, calls []Multicall2Call) (*types.Transaction, error) {
	return _Multicall2.contract.Transact(opts, "aggregate", calls)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes[] returnData)
func (_Multicall2 *Multicall2Session) Aggregate(calls []Multicall2Call) (*types.Transaction, error) {
	return _Multicall2.Contract.Aggregate(&_Multicall2.TransactOpts, calls)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes[] returnData)
func (_Multicall2 *Multicall2TransactorSession) Aggregate(calls []Multicall2Call) (*types.Transaction, error) {
	return _Multicall2.Contract.Aggregate(&_Multicall2.TransactOpts, calls)
}

// BlockAndAggregate is a paid mutator transaction binding the contract method 0xc3077fa9.
//
// Solidity: function blockAndAggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_Multicall2 *Multicall2Transactor) BlockAndAggregate(opts *bind.TransactOpts, calls []Multicall2Call) (*types.Transaction, error) {
	return _Multicall2.contract.Transact(opts, "blockAndAggregate", calls)
}

// BlockAndAggregate is a paid mutator transaction binding the contract method 0xc3077fa9.
//
// Solidity: function blockAndAggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_Multicall2 *Multicall2Session) BlockAndAggregate(calls []Multicall2Call) (*types.Transaction, error) {
	return _Multicall2.Contract.BlockAndAggregate(&_Multicall2.TransactOpts, calls)
}

// BlockAndAggregate is a paid mutator transaction binding the contract method 0xc3077fa9.
//
// Solidity: function blockAndAggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_Multicall2 *Multicall2TransactorSession) BlockAndAggregate(calls []Multicall2Call) (*types.Transaction, error) {
	return _Multicall2.Contract.BlockAndAggregate(&_Multicall2.TransactOpts, calls)
}

// TryAggregate is a paid mutator transaction binding the contract method 0xbce38bd7.
//
// Solidity: function tryAggregate(bool requireSuccess, (address,bytes)[] calls) returns((bool,bytes)[] returnData)
func (_Multicall2 *Multicall2Transactor) TryAggregate(opts *bind.TransactOpts, requireSuccess bool, calls []Multicall2Call) (*types.Transaction, error) {
	return _Multicall2.contract.Transact(opts, "tryAggregate", requireSuccess, calls)
}

// TryAggregate is a paid mutator transaction binding the contract method 0xbce38bd7.
//
// Solidity: function tryAggregate(bool requireSuccess, (address,bytes)[] calls) returns((bool,bytes)[] returnData)
func (_Multicall2 *Multicall2Session) TryAggregate(requireSuccess bool, calls []Multicall2Call) (*types.Transaction, error) {
	return _Multicall2.Contract.TryAggregate(&_Multicall2.TransactOpts, requireSuccess, calls)
}

// TryAggregate is a paid mutator transaction binding the contract method 0xbce38bd7.
//
// Solidity: function tryAggregate(bool requireSuccess, (address,bytes)[] calls) returns((bool,bytes)[] returnData)
func (_Multicall2 *Multicall2TransactorSession) TryAggregate(requireSuccess bool, calls []Multicall2Call) (*types.Transaction, error) {
	return _Multicall2.Contract.TryAggregate(&_Multicall2.TransactOpts, requireSuccess, calls)
}

// TryBlockAndAggregate is a paid mutator transaction binding the contract method 0x399542e9.
//
// Solidity: function tryBlockAndAggregate(bool requireSuccess, (address,bytes)[] calls) returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_Multicall2 *Multicall2Transactor) TryBlockAndAggregate(opts *bind.TransactOpts, requireSuccess bool, calls []Multicall2Call) (*types.Transaction, error) {
	return _Multicall2.contract.Transact(opts, "tryBlockAndAggregate", requireSuccess, calls)
}

// TryBlockAndAggregate is a paid mutator transaction binding the contract method 0x399542e9.
//
// Solidity: function tryBlockAndAggregate(bool requireSuccess, (address,bytes)[] calls) returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_Multicall2 *Multicall2Session) TryBlockAndAggregate(requireSuccess bool, calls []Multicall2Call) (*types.Transaction, error) {
	return _Multicall2.Contract.TryBlockAndAggregate(&_Multicall2.TransactOpts, requireSuccess, calls)
}

// TryBlockAndAggregate is a paid mutator transaction binding the contract method 0x399542e9.
//
// Solidity: function tryBlockAndAggregate(bool requireSuccess, (address,bytes)[] calls) returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_Multicall2 *Multicall2TransactorSession) TryBlockAndAggregate(requireSuccess bool, calls []Multicall2Call) (*types.Transaction, error) {
	return _Multicall2.Contract.TryBlockAndAggregate(&_Multicall2.TransactOpts, requireSuccess, calls)
}

