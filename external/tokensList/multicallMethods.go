package tokensList

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
)

var ERC20ABI, _ = contracts.ERC20MetaData.GetAbi()

func getBalance(name string, contractAddress common.MixedcaseAddress, userAddress common.MixedcaseAddress) ethereum.Call {
	parsedData, err := ERC20ABI.Pack("balanceOf", userAddress)
	if err != nil {
		logs.Error(err)
	}
	return ethereum.Call{
		Target:   contractAddress.Address(),
		Abi:      ERC20ABI,
		Method:   `balanceOf`,
		CallData: parsedData,
		Name:     name,
	}
}
