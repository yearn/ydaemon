package tokensList

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
)

var ERC20ABI, _ = contracts.ERC20MetaData.GetAbi()
var MulticallABI, _ = contracts.Multicall3MetaData.GetAbi()

func getBalance(name string, contractAddress common.Address, userAddress common.Address) ethereum.Call {
	parsedData, err := ERC20ABI.Pack("balanceOf", userAddress)
	if err != nil {
		logs.Error(err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      ERC20ABI,
		Method:   `balanceOf`,
		CallData: parsedData,
		Name:     name,
	}
}

func getEthBalance(name string, contractAddress common.Address, userAddress common.Address) ethereum.Call {
	parsedData, err := MulticallABI.Pack("getEthBalance", userAddress)
	if err != nil {
		logs.Error(err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      MulticallABI,
		Method:   `getEthBalance`,
		CallData: parsedData,
		Name:     name,
	}
}
