package tokensList

import (
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/types/common"
)

var ERC20ABI, _ = contracts.ERC20MetaData.GetAbi()

func getBalance(name string, contractAddress ethcommon.Address, userAddress common.Address) ethereum.Call {
	parsedData, err := ERC20ABI.Pack("balanceOf", userAddress.ToAddress())
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
