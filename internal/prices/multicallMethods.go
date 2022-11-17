package prices

import (
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/types/common"
)

/**************************************************************************************************
** The multicall require a specific format for the call data. The following functions are helpers
** used to build them for some specific methods.
**************************************************************************************************/

var LensABI, _ = contracts.OracleMetaData.GetAbi()

// getPriceUsdcRecommendedCall will pack the transaction with it's argument and return the
// ethereum.Call object that can be used to execute the transaction.
func getPriceUsdcRecommendedCall(name string, contractAddress common.Address, tokenAddress common.Address) ethereum.Call {
	parsedData, _ := LensABI.Pack("getPriceUsdcRecommended", tokenAddress.Address)
	return ethereum.Call{
		Target:   contractAddress.Address,
		Abi:      LensABI,
		Method:   `getPriceUsdcRecommended`,
		CallData: parsedData,
		Name:     name,
	}
}
