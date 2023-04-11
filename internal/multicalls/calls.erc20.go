package multicalls

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
)

var ERC20ABI, _ = contracts.ERC20MetaData.GetAbi()

func GetName(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := ERC20ABI.Pack("name")
	if err != nil {
		logs.Error("Error packing ERC20ABI name", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      ERC20ABI,
		Method:   `name`,
		CallData: parsedData,
		Name:     name,
	}
}

func GetSymbol(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := ERC20ABI.Pack("symbol")
	if err != nil {
		logs.Error("Error packing ERC20ABI symbol", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      ERC20ABI,
		Method:   `symbol`,
		CallData: parsedData,
		Name:     name,
	}
}

func GetDecimals(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := ERC20ABI.Pack("decimals")
	if err != nil {
		logs.Error("Error packing ERC20ABI decimals", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      ERC20ABI,
		Method:   `decimals`,
		CallData: parsedData,
		Name:     name,
	}
}

func GetTotalSupply(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := ERC20ABI.Pack("totalSupply")
	if err != nil {
		logs.Error("Error packing ERC20ABI totalSupply", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      ERC20ABI,
		Method:   `totalSupply`,
		CallData: parsedData,
		Name:     name,
	}
}
