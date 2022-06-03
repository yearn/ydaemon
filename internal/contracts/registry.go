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

// RegistryMetaData contains all meta data concerning the Registry contract.
var RegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"release_id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"template\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"api_version\",\"type\":\"string\"}],\"name\":\"NewRelease\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"vault_id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"api_version\",\"type\":\"string\"}],\"name\":\"NewVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"deployer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"api_version\",\"type\":\"string\"}],\"name\":\"NewExperimentalVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"governance\",\"type\":\"address\"}],\"name\":\"NewGovernance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tag\",\"type\":\"string\"}],\"name\":\"VaultTagged\",\"type\":\"event\"},{\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"gas\":36245,\"inputs\":[{\"name\":\"governance\",\"type\":\"address\"}],\"name\":\"setGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":37517,\"inputs\":[],\"name\":\"acceptGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":6831,\"inputs\":[],\"name\":\"latestRelease\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":2587,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"latestVault\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":82588,\"inputs\":[{\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"newRelease\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"guardian\",\"type\":\"address\"},{\"name\":\"rewards\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"newVault\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"guardian\",\"type\":\"address\"},{\"name\":\"rewards\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"releaseDelta\",\"type\":\"uint256\"}],\"name\":\"newVault\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"governance\",\"type\":\"address\"},{\"name\":\"guardian\",\"type\":\"address\"},{\"name\":\"rewards\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"newExperimentalVault\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"governance\",\"type\":\"address\"},{\"name\":\"guardian\",\"type\":\"address\"},{\"name\":\"rewards\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"releaseDelta\",\"type\":\"uint256\"}],\"name\":\"newExperimentalVault\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"endorseVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"vault\",\"type\":\"address\"},{\"name\":\"releaseDelta\",\"type\":\"uint256\"}],\"name\":\"endorseVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"tagger\",\"type\":\"address\"}],\"name\":\"setBanksy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"tagger\",\"type\":\"address\"},{\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"setBanksy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":186064,\"inputs\":[{\"name\":\"vault\",\"type\":\"address\"},{\"name\":\"tag\",\"type\":\"string\"}],\"name\":\"tagVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":1388,\"inputs\":[],\"name\":\"numReleases\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":1533,\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"name\":\"releases\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":1663,\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"name\":\"numVaults\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":1808,\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"},{\"name\":\"arg1\",\"type\":\"uint256\"}],\"name\":\"vaults\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":1623,\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"name\":\"tokens\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":1538,\"inputs\":[],\"name\":\"numTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":1783,\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"name\":\"isRegistered\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":1598,\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":1628,\"inputs\":[],\"name\":\"pendingGovernance\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":10229,\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"name\":\"tags\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":1903,\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"name\":\"banksy\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x336007556115eb56600436101561000d576110f9565b600035601c52600051341561002157600080fd5b63ab033ea98114156100535760043560a01c1561003d57600080fd5b600754331461004b57600080fd5b600435600855005b63238efcbc81141561009f57600854331461006d57600080fd5b3360075533610140527f4f386975ea1c2f7cf1845b08bee00626fbb624ecad16254d63d9bb9ba86526de6020610140a1005b637be0ca5e8114156101735760606101a0600463258294106101405261015c60016000546001808210156100d257600080fd5b8082039050905060e05260c052604060c020545afa6100f057600080fd5b603f3d116100fd57600080fd5b601d6101a06101a05101511061011257600080fd5b6000506101c08051602001806102408284600060045af161013257600080fd5b5050610240518061026001818260206001820306601f82010390500336823750506020610220526040610240510160206001820306601f8201039050610220f35b63e177dc708114156101db5760043560a01c1561018f57600080fd5b600360043560e05260c052604060c020600260043560e05260c052604060c020546001808210156101bf57600080fd5b8082039050905060e05260c052604060c0205460005260206000f35b6333990d4b8114156104315760043560a01c156101f757600080fd5b600754331461020557600080fd5b6000546101405260006101405111156102e1576060610400600463258294106103a0526103bc6004355afa61023957600080fd5b603f3d1161024657600080fd5b601d6104006104005101511061025b57600080fd5b600050610420602001516060610340600463258294106102e0526102fc60016101405160018082101561028d57600080fd5b8082039050905060e05260c052604060c020545afa6102ab57600080fd5b603f3d116102b857600080fd5b601d610340610340510151106102cd57600080fd5b60005061036060200151186102e157600080fd5b60043560016101405160e05260c052604060c0205561014051600181818301101561030b57600080fd5b8082019050905060005560606101c0600463258294106101605261017c6004355afa61033657600080fd5b603f3d1161034357600080fd5b601d6101c06101c05101511061035857600080fd5b6000506101e08051602001806102208284600060045af161037857600080fd5b50506004356102c052604061028052610280516102e052610220805160200180610280516102c0018284600060045af16103b157600080fd5b5050610280516102c00151806020610280516102c0010101818260206001820306601f82010390500336823750506020610280516102c0015160206001820306601f820103905061028051010161028052610140517fa6fbd216b6734f34092f1be6b7247e1551a6d4f2d5000c53721cfdc119a5b7cf610280516102c0a2005b63108ca11e81141561044857600061022052610469565b63b0b40fce81141561046457602060a461022037600050610469565b6106f4565b60043560a01c1561047957600080fd5b60243560a01c1561048957600080fd5b60443560a01c1561049957600080fd5b60606064356004016101403760406064356004013511156104b957600080fd5b60406084356004016101c03760206084356004013511156104d957600080fd5b60075433146104e757600080fd5b6000546001808210156104f957600080fd5b80820390509050610220518082101561051157600080fd5b80820390509050610240526101405161016051610180516101a0516101c0516101e0516102005161022051610240516102605160e060043561028052336102a0526044356102c0526024356102e05280610300526101408080516020018084610280018284600060045af161058557600080fd5b50508051820160206001820306601f820103905060200191505080610320526101c08080516020018084610280018284600060045af16105c457600080fd5b50505061024051610340525061032051806102800180518060206001820306601f82010390508201610440525050505b6103606104405110156106065761061b565b610440515160206104405103610440526105f4565b6103405161032051610300516102e0516102c0516102a05161028051600658016110ff565b61046052610260526102405261022052610200526101e0526101c0526101a05261018052610160526101405261046051610260526101405161016051610180516101a0516101c0516101e0516102005161022051610240516102605160043561028052610260516102a0526102a0516102805160065801611302565b610260526102405261022052610200526101e0526101c0526101a0526101805261016052610140526000506102605160005260206000f35b635b73aa0d81141561070b5760006102205261072c565b635bd4b0f281141561072757602060c46102203760005061072c565b610a64565b60043560a01c1561073c57600080fd5b60243560a01c1561074c57600080fd5b60443560a01c1561075c57600080fd5b60643560a01c1561076c57600080fd5b606060843560040161014037604060843560040135111561078c57600080fd5b604060a4356004016101c037602060a4356004013511156107ac57600080fd5b6000546001808210156107be57600080fd5b8082039050905061022051808210156107d657600080fd5b80820390509050610240526101405161016051610180516101a0516101c0516101e0516102005161022051610240516102605160e0600435610280526024356102a0526064356102c0526044356102e05280610300526101408080516020018084610280018284600060045af161084c57600080fd5b50508051820160206001820306601f820103905060200191505080610320526101c08080516020018084610280018284600060045af161088b57600080fd5b50505061024051610340525061032051806102800180518060206001820306601f82010390508201610440525050505b6103606104405110156108cd576108e2565b610440515160206104405103610440526108bb565b6103405161032051610300516102e0516102c0516102a05161028051600658016110ff565b61046052610260526102405261022052610200526101e0526101c0526101a052610180526101605261014052610460516102605260606102e0600463258294106102805261029c610260515afa61095d57600080fd5b603f3d1161096a57600080fd5b601d6102e06102e05101511061097f57600080fd5b6000506103008051602001806103408284600060045af161099f57600080fd5b5050610260516103e05260406103a0526103a051610400526103408051602001806103a0516103e0018284600060045af16109d957600080fd5b50506103a0516103e001518060206103a0516103e0010101818260206001820306601f820103905003368237505060206103a0516103e0015160206001820306601f82010390506103a05101016103a052336004357f57a9cdc2a05e05f66e76769bdbe88e21ec45d9ee0f97d4cb60395d4c75dcbcda6103a0516103e0a36102605160005260206000f35b6329b2e0c6811415610a7b57600061014052610a9c565b63b366a35c811415610a97576020602461014037600050610a9c565b610cac565b60043560a01c15610aac57600080fd5b6007543314610aba57600080fd5b3360206101c06004635aa6e6756101605261017c6004355afa610adc57600080fd5b601f3d11610ae957600080fd5b6000506101c05114610afa57600080fd5b600054600180821015610b0c57600080fd5b808203905090506101405180821015610b2457600080fd5b80820390509050610160526060610240600463258294106101e0526101fc60016101605160e05260c052604060c020545afa610b5f57600080fd5b603f3d11610b6c57600080fd5b601d61024061024051015110610b8157600080fd5b6000506102608051602001806101808284600060045af1610ba157600080fd5b50506101a0516060610300600463258294106102a0526102bc6004355afa610bc857600080fd5b603f3d11610bd557600080fd5b601d61030061030051015110610bea57600080fd5b6000506103206020015114610bfe57600080fd5b6020610240600463fc0c546a6101e0526101fc6004355afa610c1f57600080fd5b601f3d11610c2c57600080fd5b60005061024051610260526101405161016051610180516101a0516101c0516101e0516102005161022051610240516102605161026051610280526004356102a0526102a0516102805160065801611302565b610260526102405261022052610200526101e0526101c0526101a052610180526101605261014052600050005b632cad8f9f811415610cc357600161014052610cf4565b635e05f6af811415610cef5760243560011c15610cdf57600080fd5b6020602461014037600050610cf4565b610d29565b60043560a01c15610d0457600080fd5b6007543314610d1257600080fd5b61014051600a60043560e05260c052604060c02055005b6360bd68f8811415610ea25760043560a01c15610d4557600080fd5b6098602435600401610140376078602435600401351115610d6557600080fd5b600754331815610d8757600a3360e05260c052604060c02054610d8757600080fd5b61014080600960043560e05260c052604060c02060c052602060c020602082510161012060006005818352015b82610120516020021115610dc757610de9565b61012051602002850151610120518501555b8151600101808352811415610db4575b50505050505060043561024052604061020052610200516102605261014080516020018061020051610240018284600060045af1610e2657600080fd5b505061020051610240015180602061020051610240010101818260206001820306601f8201039050033682375050602061020051610240015160206001820306601f8201039050610200510101610200527f07bd58794c2ca0ae152f7719eb5f02c654476de972cadec0e8191ae8be42096d61020051610240a1005b6356e0a94b811415610eba5760005460005260206000f35b63b6a9f40f811415610ee057600160043560e05260c052604060c0205460005260206000f35b63f9c7bba5811415610f165760043560a01c15610efc57600080fd5b600260043560e05260c052604060c0205460005260206000f35b637bbfc69e811415610f5a5760043560a01c15610f3257600080fd5b600360043560e05260c052604060c02060243560e05260c052604060c0205460005260206000f35b634f64b2be811415610f8057600460043560e05260c052604060c0205460005260206000f35b638e499bcf811415610f985760055460005260206000f35b63c3c5a547811415610fce5760043560a01c15610fb457600080fd5b600660043560e05260c052604060c0205460005260206000f35b635aa6e675811415610fe65760075460005260206000f35b63f39c38a0811415610ffe5760085460005260206000f35b6358b8f8428114156110c15760043560a01c1561101a57600080fd5b600960043560e05260c052604060c0208060c052602060c020610180602082540161012060006005818352015b8261012051602002111561105a5761107c565b61012051850154610120516020028501525b8151600101808352811415611047575b50505050505061018051806101a001818260206001820306601f82010390500336823750506020610160526040610180510160206001820306601f8201039050610160f35b63ee711ed58114156110f75760043560a01c156110dd57600080fd5b600a60043560e05260c052604060c0205460005260206000f35b505b60006000fd5b610220526101405261016052610180526101a0526101c0526101e05261020052610240526000610320525b6102405160206001820306601f82010390506103205110151561114c57611165565b610320516102600152610320516020016103205261112a565b6102c0526000610320525b6102c05160206001820306601f820103905061032051101515611192576111ab565b610320516102e001526103205160200161032052611170565b60005060016102005160e05260c052604060c0205461034052600061034051186111d457600080fd5b7f602d3d8160093d39f3363d3d373d3d3d363d7300000000000000000000000000610380526103405160601b610393527f5af43d82803e903d91602b57fd5bf300000000000000000000000000000000006103a75260366103806000f061036052610360513b61124357600080fd5b6000600061016460c063a5b81fdf61038052610140516103a052610160516103c052610180516103e052806104005261024080805160200180846103a0018284600060045af161129257600080fd5b50508051820160206001820306601f820103905060200191505080610420526102c080805160200180846103a0018284600060045af16112d157600080fd5b5050506101a0516104405261039c90506000610360515af16112f257600080fd5b6103605160005260005161022051565b61018052610140526101605260026101405160e05260c052604060c020546101a05260006101a0511115611409576060610460600463258294106104005261041c610160515afa61135257600080fd5b603f3d1161135f57600080fd5b601d6104606104605101511061137457600080fd5b6000506104806020015160606103a0600463258294106103405261035c60036101405160e05260c052604060c0206101a0516001808210156113b557600080fd5b8082039050905060e05260c052604060c020545afa6113d357600080fd5b603f3d116113e057600080fd5b601d6103a06103a0510151106113f557600080fd5b6000506103c0602001511861140957600080fd5b6101605160036101405160e05260c052604060c0206101a05160e05260c052604060c020556101a051600181818301101561144357600080fd5b8082019050905060026101405160e05260c052604060c0205560066101405160e05260c052604060c0205415156114bd57600160066101405160e05260c052604060c0205561014051600460055460e05260c052604060c020556005805460018181830110156114b257600080fd5b808201905090508155505b6060610220600463258294106101c0526101dc610160515afa6114df57600080fd5b603f3d116114ec57600080fd5b601d6102206102205101511061150157600080fd5b6000506102408051602001806102808284600060045af161152157600080fd5b5050610160516103205260406102e0526102e051610340526102808051602001806102e051610320018284600060045af161155b57600080fd5b50506102e05161032001518060206102e051610320010101818260206001820306601f820103905003368237505060206102e051610320015160206001820306601f82010390506102e05101016102e0526101a051610140517fce089905ba4a4d622553bcb5646fd23e895c256f0376eee04e99e61cec1dc7e86102e051610320a361018051565b6100086115eb036100086000396100086115eb036000f3",
}

// RegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use RegistryMetaData.ABI instead.
var RegistryABI = RegistryMetaData.ABI

// RegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RegistryMetaData.Bin instead.
var RegistryBin = RegistryMetaData.Bin

// DeployRegistry deploys a new Ethereum contract, binding an instance of Registry to it.
func DeployRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Registry, error) {
	parsed, err := RegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// Registry is an auto generated Go binding around an Ethereum contract.
type Registry struct {
	RegistryCaller     // Read-only binding to the contract
	RegistryTransactor // Write-only binding to the contract
	RegistryFilterer   // Log filterer for contract events
}

// RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrySession struct {
	Contract     *Registry         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistryCallerSession struct {
	Contract *RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistryTransactorSession struct {
	Contract     *RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistryRaw struct {
	Contract *Registry // Generic contract binding to access the raw methods on
}

// RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistryCallerRaw struct {
	Contract *RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistryTransactorRaw struct {
	Contract *RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistry creates a new instance of Registry, bound to a specific deployed contract.
func NewRegistry(address common.Address, backend bind.ContractBackend) (*Registry, error) {
	contract, err := bindRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// NewRegistryCaller creates a new read-only instance of Registry, bound to a specific deployed contract.
func NewRegistryCaller(address common.Address, caller bind.ContractCaller) (*RegistryCaller, error) {
	contract, err := bindRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryCaller{contract: contract}, nil
}

// NewRegistryTransactor creates a new write-only instance of Registry, bound to a specific deployed contract.
func NewRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryTransactor, error) {
	contract, err := bindRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryTransactor{contract: contract}, nil
}

// NewRegistryFilterer creates a new log filterer instance of Registry, bound to a specific deployed contract.
func NewRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryFilterer, error) {
	contract, err := bindRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryFilterer{contract: contract}, nil
}

// bindRegistry binds a generic wrapper to an already deployed contract.
func bindRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transact(opts, method, params...)
}

// Banksy is a free data retrieval call binding the contract method 0xee711ed5.
//
// Solidity: function banksy(address arg0) view returns(bool)
func (_Registry *RegistryCaller) Banksy(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "banksy", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Banksy is a free data retrieval call binding the contract method 0xee711ed5.
//
// Solidity: function banksy(address arg0) view returns(bool)
func (_Registry *RegistrySession) Banksy(arg0 common.Address) (bool, error) {
	return _Registry.Contract.Banksy(&_Registry.CallOpts, arg0)
}

// Banksy is a free data retrieval call binding the contract method 0xee711ed5.
//
// Solidity: function banksy(address arg0) view returns(bool)
func (_Registry *RegistryCallerSession) Banksy(arg0 common.Address) (bool, error) {
	return _Registry.Contract.Banksy(&_Registry.CallOpts, arg0)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Registry *RegistryCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Registry *RegistrySession) Governance() (common.Address, error) {
	return _Registry.Contract.Governance(&_Registry.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Registry *RegistryCallerSession) Governance() (common.Address, error) {
	return _Registry.Contract.Governance(&_Registry.CallOpts)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address arg0) view returns(bool)
func (_Registry *RegistryCaller) IsRegistered(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "isRegistered", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address arg0) view returns(bool)
func (_Registry *RegistrySession) IsRegistered(arg0 common.Address) (bool, error) {
	return _Registry.Contract.IsRegistered(&_Registry.CallOpts, arg0)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address arg0) view returns(bool)
func (_Registry *RegistryCallerSession) IsRegistered(arg0 common.Address) (bool, error) {
	return _Registry.Contract.IsRegistered(&_Registry.CallOpts, arg0)
}

// LatestRelease is a free data retrieval call binding the contract method 0x7be0ca5e.
//
// Solidity: function latestRelease() view returns(string)
func (_Registry *RegistryCaller) LatestRelease(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "latestRelease")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LatestRelease is a free data retrieval call binding the contract method 0x7be0ca5e.
//
// Solidity: function latestRelease() view returns(string)
func (_Registry *RegistrySession) LatestRelease() (string, error) {
	return _Registry.Contract.LatestRelease(&_Registry.CallOpts)
}

// LatestRelease is a free data retrieval call binding the contract method 0x7be0ca5e.
//
// Solidity: function latestRelease() view returns(string)
func (_Registry *RegistryCallerSession) LatestRelease() (string, error) {
	return _Registry.Contract.LatestRelease(&_Registry.CallOpts)
}

// LatestVault is a free data retrieval call binding the contract method 0xe177dc70.
//
// Solidity: function latestVault(address token) view returns(address)
func (_Registry *RegistryCaller) LatestVault(opts *bind.CallOpts, token common.Address) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "latestVault", token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LatestVault is a free data retrieval call binding the contract method 0xe177dc70.
//
// Solidity: function latestVault(address token) view returns(address)
func (_Registry *RegistrySession) LatestVault(token common.Address) (common.Address, error) {
	return _Registry.Contract.LatestVault(&_Registry.CallOpts, token)
}

// LatestVault is a free data retrieval call binding the contract method 0xe177dc70.
//
// Solidity: function latestVault(address token) view returns(address)
func (_Registry *RegistryCallerSession) LatestVault(token common.Address) (common.Address, error) {
	return _Registry.Contract.LatestVault(&_Registry.CallOpts, token)
}

// NumReleases is a free data retrieval call binding the contract method 0x56e0a94b.
//
// Solidity: function numReleases() view returns(uint256)
func (_Registry *RegistryCaller) NumReleases(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "numReleases")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumReleases is a free data retrieval call binding the contract method 0x56e0a94b.
//
// Solidity: function numReleases() view returns(uint256)
func (_Registry *RegistrySession) NumReleases() (*big.Int, error) {
	return _Registry.Contract.NumReleases(&_Registry.CallOpts)
}

// NumReleases is a free data retrieval call binding the contract method 0x56e0a94b.
//
// Solidity: function numReleases() view returns(uint256)
func (_Registry *RegistryCallerSession) NumReleases() (*big.Int, error) {
	return _Registry.Contract.NumReleases(&_Registry.CallOpts)
}

// NumTokens is a free data retrieval call binding the contract method 0x8e499bcf.
//
// Solidity: function numTokens() view returns(uint256)
func (_Registry *RegistryCaller) NumTokens(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "numTokens")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumTokens is a free data retrieval call binding the contract method 0x8e499bcf.
//
// Solidity: function numTokens() view returns(uint256)
func (_Registry *RegistrySession) NumTokens() (*big.Int, error) {
	return _Registry.Contract.NumTokens(&_Registry.CallOpts)
}

// NumTokens is a free data retrieval call binding the contract method 0x8e499bcf.
//
// Solidity: function numTokens() view returns(uint256)
func (_Registry *RegistryCallerSession) NumTokens() (*big.Int, error) {
	return _Registry.Contract.NumTokens(&_Registry.CallOpts)
}

// NumVaults is a free data retrieval call binding the contract method 0xf9c7bba5.
//
// Solidity: function numVaults(address arg0) view returns(uint256)
func (_Registry *RegistryCaller) NumVaults(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "numVaults", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumVaults is a free data retrieval call binding the contract method 0xf9c7bba5.
//
// Solidity: function numVaults(address arg0) view returns(uint256)
func (_Registry *RegistrySession) NumVaults(arg0 common.Address) (*big.Int, error) {
	return _Registry.Contract.NumVaults(&_Registry.CallOpts, arg0)
}

// NumVaults is a free data retrieval call binding the contract method 0xf9c7bba5.
//
// Solidity: function numVaults(address arg0) view returns(uint256)
func (_Registry *RegistryCallerSession) NumVaults(arg0 common.Address) (*big.Int, error) {
	return _Registry.Contract.NumVaults(&_Registry.CallOpts, arg0)
}

// PendingGovernance is a free data retrieval call binding the contract method 0xf39c38a0.
//
// Solidity: function pendingGovernance() view returns(address)
func (_Registry *RegistryCaller) PendingGovernance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "pendingGovernance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingGovernance is a free data retrieval call binding the contract method 0xf39c38a0.
//
// Solidity: function pendingGovernance() view returns(address)
func (_Registry *RegistrySession) PendingGovernance() (common.Address, error) {
	return _Registry.Contract.PendingGovernance(&_Registry.CallOpts)
}

// PendingGovernance is a free data retrieval call binding the contract method 0xf39c38a0.
//
// Solidity: function pendingGovernance() view returns(address)
func (_Registry *RegistryCallerSession) PendingGovernance() (common.Address, error) {
	return _Registry.Contract.PendingGovernance(&_Registry.CallOpts)
}

// Releases is a free data retrieval call binding the contract method 0xb6a9f40f.
//
// Solidity: function releases(uint256 arg0) view returns(address)
func (_Registry *RegistryCaller) Releases(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "releases", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Releases is a free data retrieval call binding the contract method 0xb6a9f40f.
//
// Solidity: function releases(uint256 arg0) view returns(address)
func (_Registry *RegistrySession) Releases(arg0 *big.Int) (common.Address, error) {
	return _Registry.Contract.Releases(&_Registry.CallOpts, arg0)
}

// Releases is a free data retrieval call binding the contract method 0xb6a9f40f.
//
// Solidity: function releases(uint256 arg0) view returns(address)
func (_Registry *RegistryCallerSession) Releases(arg0 *big.Int) (common.Address, error) {
	return _Registry.Contract.Releases(&_Registry.CallOpts, arg0)
}

// Tags is a free data retrieval call binding the contract method 0x58b8f842.
//
// Solidity: function tags(address arg0) view returns(string)
func (_Registry *RegistryCaller) Tags(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "tags", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Tags is a free data retrieval call binding the contract method 0x58b8f842.
//
// Solidity: function tags(address arg0) view returns(string)
func (_Registry *RegistrySession) Tags(arg0 common.Address) (string, error) {
	return _Registry.Contract.Tags(&_Registry.CallOpts, arg0)
}

// Tags is a free data retrieval call binding the contract method 0x58b8f842.
//
// Solidity: function tags(address arg0) view returns(string)
func (_Registry *RegistryCallerSession) Tags(arg0 common.Address) (string, error) {
	return _Registry.Contract.Tags(&_Registry.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 arg0) view returns(address)
func (_Registry *RegistryCaller) Tokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "tokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 arg0) view returns(address)
func (_Registry *RegistrySession) Tokens(arg0 *big.Int) (common.Address, error) {
	return _Registry.Contract.Tokens(&_Registry.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 arg0) view returns(address)
func (_Registry *RegistryCallerSession) Tokens(arg0 *big.Int) (common.Address, error) {
	return _Registry.Contract.Tokens(&_Registry.CallOpts, arg0)
}

// Vaults is a free data retrieval call binding the contract method 0x7bbfc69e.
//
// Solidity: function vaults(address arg0, uint256 arg1) view returns(address)
func (_Registry *RegistryCaller) Vaults(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "vaults", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vaults is a free data retrieval call binding the contract method 0x7bbfc69e.
//
// Solidity: function vaults(address arg0, uint256 arg1) view returns(address)
func (_Registry *RegistrySession) Vaults(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Registry.Contract.Vaults(&_Registry.CallOpts, arg0, arg1)
}

// Vaults is a free data retrieval call binding the contract method 0x7bbfc69e.
//
// Solidity: function vaults(address arg0, uint256 arg1) view returns(address)
func (_Registry *RegistryCallerSession) Vaults(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Registry.Contract.Vaults(&_Registry.CallOpts, arg0, arg1)
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_Registry *RegistryTransactor) AcceptGovernance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "acceptGovernance")
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_Registry *RegistrySession) AcceptGovernance() (*types.Transaction, error) {
	return _Registry.Contract.AcceptGovernance(&_Registry.TransactOpts)
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_Registry *RegistryTransactorSession) AcceptGovernance() (*types.Transaction, error) {
	return _Registry.Contract.AcceptGovernance(&_Registry.TransactOpts)
}

// EndorseVault is a paid mutator transaction binding the contract method 0x29b2e0c6.
//
// Solidity: function endorseVault(address vault) returns()
func (_Registry *RegistryTransactor) EndorseVault(opts *bind.TransactOpts, vault common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "endorseVault", vault)
}

// EndorseVault is a paid mutator transaction binding the contract method 0x29b2e0c6.
//
// Solidity: function endorseVault(address vault) returns()
func (_Registry *RegistrySession) EndorseVault(vault common.Address) (*types.Transaction, error) {
	return _Registry.Contract.EndorseVault(&_Registry.TransactOpts, vault)
}

// EndorseVault is a paid mutator transaction binding the contract method 0x29b2e0c6.
//
// Solidity: function endorseVault(address vault) returns()
func (_Registry *RegistryTransactorSession) EndorseVault(vault common.Address) (*types.Transaction, error) {
	return _Registry.Contract.EndorseVault(&_Registry.TransactOpts, vault)
}

// EndorseVault0 is a paid mutator transaction binding the contract method 0xb366a35c.
//
// Solidity: function endorseVault(address vault, uint256 releaseDelta) returns()
func (_Registry *RegistryTransactor) EndorseVault0(opts *bind.TransactOpts, vault common.Address, releaseDelta *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "endorseVault0", vault, releaseDelta)
}

// EndorseVault0 is a paid mutator transaction binding the contract method 0xb366a35c.
//
// Solidity: function endorseVault(address vault, uint256 releaseDelta) returns()
func (_Registry *RegistrySession) EndorseVault0(vault common.Address, releaseDelta *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.EndorseVault0(&_Registry.TransactOpts, vault, releaseDelta)
}

// EndorseVault0 is a paid mutator transaction binding the contract method 0xb366a35c.
//
// Solidity: function endorseVault(address vault, uint256 releaseDelta) returns()
func (_Registry *RegistryTransactorSession) EndorseVault0(vault common.Address, releaseDelta *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.EndorseVault0(&_Registry.TransactOpts, vault, releaseDelta)
}

// NewExperimentalVault is a paid mutator transaction binding the contract method 0x5b73aa0d.
//
// Solidity: function newExperimentalVault(address token, address governance, address guardian, address rewards, string name, string symbol) returns(address)
func (_Registry *RegistryTransactor) NewExperimentalVault(opts *bind.TransactOpts, token common.Address, governance common.Address, guardian common.Address, rewards common.Address, name string, symbol string) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "newExperimentalVault", token, governance, guardian, rewards, name, symbol)
}

// NewExperimentalVault is a paid mutator transaction binding the contract method 0x5b73aa0d.
//
// Solidity: function newExperimentalVault(address token, address governance, address guardian, address rewards, string name, string symbol) returns(address)
func (_Registry *RegistrySession) NewExperimentalVault(token common.Address, governance common.Address, guardian common.Address, rewards common.Address, name string, symbol string) (*types.Transaction, error) {
	return _Registry.Contract.NewExperimentalVault(&_Registry.TransactOpts, token, governance, guardian, rewards, name, symbol)
}

// NewExperimentalVault is a paid mutator transaction binding the contract method 0x5b73aa0d.
//
// Solidity: function newExperimentalVault(address token, address governance, address guardian, address rewards, string name, string symbol) returns(address)
func (_Registry *RegistryTransactorSession) NewExperimentalVault(token common.Address, governance common.Address, guardian common.Address, rewards common.Address, name string, symbol string) (*types.Transaction, error) {
	return _Registry.Contract.NewExperimentalVault(&_Registry.TransactOpts, token, governance, guardian, rewards, name, symbol)
}

// NewExperimentalVault0 is a paid mutator transaction binding the contract method 0x5bd4b0f2.
//
// Solidity: function newExperimentalVault(address token, address governance, address guardian, address rewards, string name, string symbol, uint256 releaseDelta) returns(address)
func (_Registry *RegistryTransactor) NewExperimentalVault0(opts *bind.TransactOpts, token common.Address, governance common.Address, guardian common.Address, rewards common.Address, name string, symbol string, releaseDelta *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "newExperimentalVault0", token, governance, guardian, rewards, name, symbol, releaseDelta)
}

// NewExperimentalVault0 is a paid mutator transaction binding the contract method 0x5bd4b0f2.
//
// Solidity: function newExperimentalVault(address token, address governance, address guardian, address rewards, string name, string symbol, uint256 releaseDelta) returns(address)
func (_Registry *RegistrySession) NewExperimentalVault0(token common.Address, governance common.Address, guardian common.Address, rewards common.Address, name string, symbol string, releaseDelta *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.NewExperimentalVault0(&_Registry.TransactOpts, token, governance, guardian, rewards, name, symbol, releaseDelta)
}

// NewExperimentalVault0 is a paid mutator transaction binding the contract method 0x5bd4b0f2.
//
// Solidity: function newExperimentalVault(address token, address governance, address guardian, address rewards, string name, string symbol, uint256 releaseDelta) returns(address)
func (_Registry *RegistryTransactorSession) NewExperimentalVault0(token common.Address, governance common.Address, guardian common.Address, rewards common.Address, name string, symbol string, releaseDelta *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.NewExperimentalVault0(&_Registry.TransactOpts, token, governance, guardian, rewards, name, symbol, releaseDelta)
}

// NewRelease is a paid mutator transaction binding the contract method 0x33990d4b.
//
// Solidity: function newRelease(address vault) returns()
func (_Registry *RegistryTransactor) NewRelease(opts *bind.TransactOpts, vault common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "newRelease", vault)
}

// NewRelease is a paid mutator transaction binding the contract method 0x33990d4b.
//
// Solidity: function newRelease(address vault) returns()
func (_Registry *RegistrySession) NewRelease(vault common.Address) (*types.Transaction, error) {
	return _Registry.Contract.NewRelease(&_Registry.TransactOpts, vault)
}

// NewRelease is a paid mutator transaction binding the contract method 0x33990d4b.
//
// Solidity: function newRelease(address vault) returns()
func (_Registry *RegistryTransactorSession) NewRelease(vault common.Address) (*types.Transaction, error) {
	return _Registry.Contract.NewRelease(&_Registry.TransactOpts, vault)
}

// NewVault is a paid mutator transaction binding the contract method 0x108ca11e.
//
// Solidity: function newVault(address token, address guardian, address rewards, string name, string symbol) returns(address)
func (_Registry *RegistryTransactor) NewVault(opts *bind.TransactOpts, token common.Address, guardian common.Address, rewards common.Address, name string, symbol string) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "newVault", token, guardian, rewards, name, symbol)
}

// NewVault is a paid mutator transaction binding the contract method 0x108ca11e.
//
// Solidity: function newVault(address token, address guardian, address rewards, string name, string symbol) returns(address)
func (_Registry *RegistrySession) NewVault(token common.Address, guardian common.Address, rewards common.Address, name string, symbol string) (*types.Transaction, error) {
	return _Registry.Contract.NewVault(&_Registry.TransactOpts, token, guardian, rewards, name, symbol)
}

// NewVault is a paid mutator transaction binding the contract method 0x108ca11e.
//
// Solidity: function newVault(address token, address guardian, address rewards, string name, string symbol) returns(address)
func (_Registry *RegistryTransactorSession) NewVault(token common.Address, guardian common.Address, rewards common.Address, name string, symbol string) (*types.Transaction, error) {
	return _Registry.Contract.NewVault(&_Registry.TransactOpts, token, guardian, rewards, name, symbol)
}

// NewVault0 is a paid mutator transaction binding the contract method 0xb0b40fce.
//
// Solidity: function newVault(address token, address guardian, address rewards, string name, string symbol, uint256 releaseDelta) returns(address)
func (_Registry *RegistryTransactor) NewVault0(opts *bind.TransactOpts, token common.Address, guardian common.Address, rewards common.Address, name string, symbol string, releaseDelta *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "newVault0", token, guardian, rewards, name, symbol, releaseDelta)
}

// NewVault0 is a paid mutator transaction binding the contract method 0xb0b40fce.
//
// Solidity: function newVault(address token, address guardian, address rewards, string name, string symbol, uint256 releaseDelta) returns(address)
func (_Registry *RegistrySession) NewVault0(token common.Address, guardian common.Address, rewards common.Address, name string, symbol string, releaseDelta *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.NewVault0(&_Registry.TransactOpts, token, guardian, rewards, name, symbol, releaseDelta)
}

// NewVault0 is a paid mutator transaction binding the contract method 0xb0b40fce.
//
// Solidity: function newVault(address token, address guardian, address rewards, string name, string symbol, uint256 releaseDelta) returns(address)
func (_Registry *RegistryTransactorSession) NewVault0(token common.Address, guardian common.Address, rewards common.Address, name string, symbol string, releaseDelta *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.NewVault0(&_Registry.TransactOpts, token, guardian, rewards, name, symbol, releaseDelta)
}

// SetBanksy is a paid mutator transaction binding the contract method 0x2cad8f9f.
//
// Solidity: function setBanksy(address tagger) returns()
func (_Registry *RegistryTransactor) SetBanksy(opts *bind.TransactOpts, tagger common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setBanksy", tagger)
}

// SetBanksy is a paid mutator transaction binding the contract method 0x2cad8f9f.
//
// Solidity: function setBanksy(address tagger) returns()
func (_Registry *RegistrySession) SetBanksy(tagger common.Address) (*types.Transaction, error) {
	return _Registry.Contract.SetBanksy(&_Registry.TransactOpts, tagger)
}

// SetBanksy is a paid mutator transaction binding the contract method 0x2cad8f9f.
//
// Solidity: function setBanksy(address tagger) returns()
func (_Registry *RegistryTransactorSession) SetBanksy(tagger common.Address) (*types.Transaction, error) {
	return _Registry.Contract.SetBanksy(&_Registry.TransactOpts, tagger)
}

// SetBanksy0 is a paid mutator transaction binding the contract method 0x5e05f6af.
//
// Solidity: function setBanksy(address tagger, bool allowed) returns()
func (_Registry *RegistryTransactor) SetBanksy0(opts *bind.TransactOpts, tagger common.Address, allowed bool) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setBanksy0", tagger, allowed)
}

// SetBanksy0 is a paid mutator transaction binding the contract method 0x5e05f6af.
//
// Solidity: function setBanksy(address tagger, bool allowed) returns()
func (_Registry *RegistrySession) SetBanksy0(tagger common.Address, allowed bool) (*types.Transaction, error) {
	return _Registry.Contract.SetBanksy0(&_Registry.TransactOpts, tagger, allowed)
}

// SetBanksy0 is a paid mutator transaction binding the contract method 0x5e05f6af.
//
// Solidity: function setBanksy(address tagger, bool allowed) returns()
func (_Registry *RegistryTransactorSession) SetBanksy0(tagger common.Address, allowed bool) (*types.Transaction, error) {
	return _Registry.Contract.SetBanksy0(&_Registry.TransactOpts, tagger, allowed)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address governance) returns()
func (_Registry *RegistryTransactor) SetGovernance(opts *bind.TransactOpts, governance common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setGovernance", governance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address governance) returns()
func (_Registry *RegistrySession) SetGovernance(governance common.Address) (*types.Transaction, error) {
	return _Registry.Contract.SetGovernance(&_Registry.TransactOpts, governance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address governance) returns()
func (_Registry *RegistryTransactorSession) SetGovernance(governance common.Address) (*types.Transaction, error) {
	return _Registry.Contract.SetGovernance(&_Registry.TransactOpts, governance)
}

// TagVault is a paid mutator transaction binding the contract method 0x60bd68f8.
//
// Solidity: function tagVault(address vault, string tag) returns()
func (_Registry *RegistryTransactor) TagVault(opts *bind.TransactOpts, vault common.Address, tag string) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "tagVault", vault, tag)
}

// TagVault is a paid mutator transaction binding the contract method 0x60bd68f8.
//
// Solidity: function tagVault(address vault, string tag) returns()
func (_Registry *RegistrySession) TagVault(vault common.Address, tag string) (*types.Transaction, error) {
	return _Registry.Contract.TagVault(&_Registry.TransactOpts, vault, tag)
}

// TagVault is a paid mutator transaction binding the contract method 0x60bd68f8.
//
// Solidity: function tagVault(address vault, string tag) returns()
func (_Registry *RegistryTransactorSession) TagVault(vault common.Address, tag string) (*types.Transaction, error) {
	return _Registry.Contract.TagVault(&_Registry.TransactOpts, vault, tag)
}

// RegistryNewExperimentalVaultIterator is returned from FilterNewExperimentalVault and is used to iterate over the raw logs and unpacked data for NewExperimentalVault events raised by the Registry contract.
type RegistryNewExperimentalVaultIterator struct {
	Event *RegistryNewExperimentalVault // Event containing the contract specifics and raw log

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
func (it *RegistryNewExperimentalVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryNewExperimentalVault)
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
		it.Event = new(RegistryNewExperimentalVault)
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
func (it *RegistryNewExperimentalVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryNewExperimentalVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryNewExperimentalVault represents a NewExperimentalVault event raised by the Registry contract.
type RegistryNewExperimentalVault struct {
	Token      common.Address
	Deployer   common.Address
	Vault      common.Address
	ApiVersion string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewExperimentalVault is a free log retrieval operation binding the contract event 0x57a9cdc2a05e05f66e76769bdbe88e21ec45d9ee0f97d4cb60395d4c75dcbcda.
//
// Solidity: event NewExperimentalVault(address indexed token, address indexed deployer, address vault, string api_version)
func (_Registry *RegistryFilterer) FilterNewExperimentalVault(opts *bind.FilterOpts, token []common.Address, deployer []common.Address) (*RegistryNewExperimentalVaultIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var deployerRule []interface{}
	for _, deployerItem := range deployer {
		deployerRule = append(deployerRule, deployerItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "NewExperimentalVault", tokenRule, deployerRule)
	if err != nil {
		return nil, err
	}
	return &RegistryNewExperimentalVaultIterator{contract: _Registry.contract, event: "NewExperimentalVault", logs: logs, sub: sub}, nil
}

// WatchNewExperimentalVault is a free log subscription operation binding the contract event 0x57a9cdc2a05e05f66e76769bdbe88e21ec45d9ee0f97d4cb60395d4c75dcbcda.
//
// Solidity: event NewExperimentalVault(address indexed token, address indexed deployer, address vault, string api_version)
func (_Registry *RegistryFilterer) WatchNewExperimentalVault(opts *bind.WatchOpts, sink chan<- *RegistryNewExperimentalVault, token []common.Address, deployer []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var deployerRule []interface{}
	for _, deployerItem := range deployer {
		deployerRule = append(deployerRule, deployerItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "NewExperimentalVault", tokenRule, deployerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryNewExperimentalVault)
				if err := _Registry.contract.UnpackLog(event, "NewExperimentalVault", log); err != nil {
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

// ParseNewExperimentalVault is a log parse operation binding the contract event 0x57a9cdc2a05e05f66e76769bdbe88e21ec45d9ee0f97d4cb60395d4c75dcbcda.
//
// Solidity: event NewExperimentalVault(address indexed token, address indexed deployer, address vault, string api_version)
func (_Registry *RegistryFilterer) ParseNewExperimentalVault(log types.Log) (*RegistryNewExperimentalVault, error) {
	event := new(RegistryNewExperimentalVault)
	if err := _Registry.contract.UnpackLog(event, "NewExperimentalVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryNewGovernanceIterator is returned from FilterNewGovernance and is used to iterate over the raw logs and unpacked data for NewGovernance events raised by the Registry contract.
type RegistryNewGovernanceIterator struct {
	Event *RegistryNewGovernance // Event containing the contract specifics and raw log

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
func (it *RegistryNewGovernanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryNewGovernance)
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
		it.Event = new(RegistryNewGovernance)
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
func (it *RegistryNewGovernanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryNewGovernanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryNewGovernance represents a NewGovernance event raised by the Registry contract.
type RegistryNewGovernance struct {
	Governance common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewGovernance is a free log retrieval operation binding the contract event 0x4f386975ea1c2f7cf1845b08bee00626fbb624ecad16254d63d9bb9ba86526de.
//
// Solidity: event NewGovernance(address governance)
func (_Registry *RegistryFilterer) FilterNewGovernance(opts *bind.FilterOpts) (*RegistryNewGovernanceIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "NewGovernance")
	if err != nil {
		return nil, err
	}
	return &RegistryNewGovernanceIterator{contract: _Registry.contract, event: "NewGovernance", logs: logs, sub: sub}, nil
}

// WatchNewGovernance is a free log subscription operation binding the contract event 0x4f386975ea1c2f7cf1845b08bee00626fbb624ecad16254d63d9bb9ba86526de.
//
// Solidity: event NewGovernance(address governance)
func (_Registry *RegistryFilterer) WatchNewGovernance(opts *bind.WatchOpts, sink chan<- *RegistryNewGovernance) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "NewGovernance")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryNewGovernance)
				if err := _Registry.contract.UnpackLog(event, "NewGovernance", log); err != nil {
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

// ParseNewGovernance is a log parse operation binding the contract event 0x4f386975ea1c2f7cf1845b08bee00626fbb624ecad16254d63d9bb9ba86526de.
//
// Solidity: event NewGovernance(address governance)
func (_Registry *RegistryFilterer) ParseNewGovernance(log types.Log) (*RegistryNewGovernance, error) {
	event := new(RegistryNewGovernance)
	if err := _Registry.contract.UnpackLog(event, "NewGovernance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryNewReleaseIterator is returned from FilterNewRelease and is used to iterate over the raw logs and unpacked data for NewRelease events raised by the Registry contract.
type RegistryNewReleaseIterator struct {
	Event *RegistryNewRelease // Event containing the contract specifics and raw log

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
func (it *RegistryNewReleaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryNewRelease)
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
		it.Event = new(RegistryNewRelease)
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
func (it *RegistryNewReleaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryNewReleaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryNewRelease represents a NewRelease event raised by the Registry contract.
type RegistryNewRelease struct {
	ReleaseId  *big.Int
	Template   common.Address
	ApiVersion string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewRelease is a free log retrieval operation binding the contract event 0xa6fbd216b6734f34092f1be6b7247e1551a6d4f2d5000c53721cfdc119a5b7cf.
//
// Solidity: event NewRelease(uint256 indexed release_id, address template, string api_version)
func (_Registry *RegistryFilterer) FilterNewRelease(opts *bind.FilterOpts, release_id []*big.Int) (*RegistryNewReleaseIterator, error) {

	var release_idRule []interface{}
	for _, release_idItem := range release_id {
		release_idRule = append(release_idRule, release_idItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "NewRelease", release_idRule)
	if err != nil {
		return nil, err
	}
	return &RegistryNewReleaseIterator{contract: _Registry.contract, event: "NewRelease", logs: logs, sub: sub}, nil
}

// WatchNewRelease is a free log subscription operation binding the contract event 0xa6fbd216b6734f34092f1be6b7247e1551a6d4f2d5000c53721cfdc119a5b7cf.
//
// Solidity: event NewRelease(uint256 indexed release_id, address template, string api_version)
func (_Registry *RegistryFilterer) WatchNewRelease(opts *bind.WatchOpts, sink chan<- *RegistryNewRelease, release_id []*big.Int) (event.Subscription, error) {

	var release_idRule []interface{}
	for _, release_idItem := range release_id {
		release_idRule = append(release_idRule, release_idItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "NewRelease", release_idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryNewRelease)
				if err := _Registry.contract.UnpackLog(event, "NewRelease", log); err != nil {
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

// ParseNewRelease is a log parse operation binding the contract event 0xa6fbd216b6734f34092f1be6b7247e1551a6d4f2d5000c53721cfdc119a5b7cf.
//
// Solidity: event NewRelease(uint256 indexed release_id, address template, string api_version)
func (_Registry *RegistryFilterer) ParseNewRelease(log types.Log) (*RegistryNewRelease, error) {
	event := new(RegistryNewRelease)
	if err := _Registry.contract.UnpackLog(event, "NewRelease", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryNewVaultIterator is returned from FilterNewVault and is used to iterate over the raw logs and unpacked data for NewVault events raised by the Registry contract.
type RegistryNewVaultIterator struct {
	Event *RegistryNewVault // Event containing the contract specifics and raw log

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
func (it *RegistryNewVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryNewVault)
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
		it.Event = new(RegistryNewVault)
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
func (it *RegistryNewVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryNewVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryNewVault represents a NewVault event raised by the Registry contract.
type RegistryNewVault struct {
	Token      common.Address
	VaultId    *big.Int
	Vault      common.Address
	ApiVersion string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewVault is a free log retrieval operation binding the contract event 0xce089905ba4a4d622553bcb5646fd23e895c256f0376eee04e99e61cec1dc7e8.
//
// Solidity: event NewVault(address indexed token, uint256 indexed vault_id, address vault, string api_version)
func (_Registry *RegistryFilterer) FilterNewVault(opts *bind.FilterOpts, token []common.Address, vault_id []*big.Int) (*RegistryNewVaultIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var vault_idRule []interface{}
	for _, vault_idItem := range vault_id {
		vault_idRule = append(vault_idRule, vault_idItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "NewVault", tokenRule, vault_idRule)
	if err != nil {
		return nil, err
	}
	return &RegistryNewVaultIterator{contract: _Registry.contract, event: "NewVault", logs: logs, sub: sub}, nil
}

// WatchNewVault is a free log subscription operation binding the contract event 0xce089905ba4a4d622553bcb5646fd23e895c256f0376eee04e99e61cec1dc7e8.
//
// Solidity: event NewVault(address indexed token, uint256 indexed vault_id, address vault, string api_version)
func (_Registry *RegistryFilterer) WatchNewVault(opts *bind.WatchOpts, sink chan<- *RegistryNewVault, token []common.Address, vault_id []*big.Int) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var vault_idRule []interface{}
	for _, vault_idItem := range vault_id {
		vault_idRule = append(vault_idRule, vault_idItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "NewVault", tokenRule, vault_idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryNewVault)
				if err := _Registry.contract.UnpackLog(event, "NewVault", log); err != nil {
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

// ParseNewVault is a log parse operation binding the contract event 0xce089905ba4a4d622553bcb5646fd23e895c256f0376eee04e99e61cec1dc7e8.
//
// Solidity: event NewVault(address indexed token, uint256 indexed vault_id, address vault, string api_version)
func (_Registry *RegistryFilterer) ParseNewVault(log types.Log) (*RegistryNewVault, error) {
	event := new(RegistryNewVault)
	if err := _Registry.contract.UnpackLog(event, "NewVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryVaultTaggedIterator is returned from FilterVaultTagged and is used to iterate over the raw logs and unpacked data for VaultTagged events raised by the Registry contract.
type RegistryVaultTaggedIterator struct {
	Event *RegistryVaultTagged // Event containing the contract specifics and raw log

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
func (it *RegistryVaultTaggedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryVaultTagged)
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
		it.Event = new(RegistryVaultTagged)
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
func (it *RegistryVaultTaggedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryVaultTaggedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryVaultTagged represents a VaultTagged event raised by the Registry contract.
type RegistryVaultTagged struct {
	Vault common.Address
	Tag   string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterVaultTagged is a free log retrieval operation binding the contract event 0x07bd58794c2ca0ae152f7719eb5f02c654476de972cadec0e8191ae8be42096d.
//
// Solidity: event VaultTagged(address vault, string tag)
func (_Registry *RegistryFilterer) FilterVaultTagged(opts *bind.FilterOpts) (*RegistryVaultTaggedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "VaultTagged")
	if err != nil {
		return nil, err
	}
	return &RegistryVaultTaggedIterator{contract: _Registry.contract, event: "VaultTagged", logs: logs, sub: sub}, nil
}

// WatchVaultTagged is a free log subscription operation binding the contract event 0x07bd58794c2ca0ae152f7719eb5f02c654476de972cadec0e8191ae8be42096d.
//
// Solidity: event VaultTagged(address vault, string tag)
func (_Registry *RegistryFilterer) WatchVaultTagged(opts *bind.WatchOpts, sink chan<- *RegistryVaultTagged) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "VaultTagged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryVaultTagged)
				if err := _Registry.contract.UnpackLog(event, "VaultTagged", log); err != nil {
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

// ParseVaultTagged is a log parse operation binding the contract event 0x07bd58794c2ca0ae152f7719eb5f02c654476de972cadec0e8191ae8be42096d.
//
// Solidity: event VaultTagged(address vault, string tag)
func (_Registry *RegistryFilterer) ParseVaultTagged(log types.Log) (*RegistryVaultTagged, error) {
	event := new(RegistryVaultTagged)
	if err := _Registry.contract.UnpackLog(event, "VaultTagged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

