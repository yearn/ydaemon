package prices

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/store"
	"github.com/yearn/ydaemon/internal/multicalls"
)

/**************************************************************************************************
** fetchPricePerShareFromVault will try to get the price per share for a vault type token
** It will return an array of struct with vault/asset/value
**************************************************************************************************/
func fetchPricePerShareFromVault(chainID uint64, blockNumber *uint64, tokenList []common.Address) []TVaultToAsset {
	vaultToAsset := []TVaultToAsset{}

	/**********************************************************************************************
	** The first step is to prepare the multicall, connecting to the multicall instance and
	** preparing the array of calls to send. All calls for all tokenList will be send in a single
	** multicall and will later be accessible via a concatened string `tokenAddress + methodName`.
	**********************************************************************************************/
	calls := []ethereum.Call{}
	for _, tokenAddress := range tokenList {
		if _, ok := store.GetERC20(chainID, tokenAddress); ok {
			calls = append(calls, multicalls.GetPricePerShare(tokenAddress.Hex(), tokenAddress))
			calls = append(calls, multicalls.GetToken(tokenAddress.Hex(), tokenAddress))
		}
	}

	/**********************************************************************************************
	** Then we can proceed the responses. We loop over the responses and check if the price is
	** available. If it is, we add it to the map. If it's not, we try to fetch it from an external
	** API.
	**********************************************************************************************/
	var response map[string][]interface{}
	var blockNumberBigInt *big.Int

	if blockNumber == nil {
		currentBlockNumber, _ := ethereum.GetRPC(chainID).BlockNumber(context.Background())
		blockNumber = &currentBlockNumber
		response = multicalls.Perform(chainID, calls, nil)
	} else {
		blockNumberBigInt = big.NewInt(int64(*blockNumber))
		response = multicalls.Perform(chainID, calls, blockNumberBigInt)
	}

	for _, token := range tokenList {
		rawPricePerShare := response[token.Hex()+`pricePerShare`]
		rawToken := response[token.Hex()+`token`]
		if len(rawPricePerShare) == 0 || len(rawToken) == 0 {
			continue
		}
		tokenPrice := bigNumber.SetInt(rawPricePerShare[0].(*big.Int))
		if tokenPrice.IsZero() {
			continue
		}

		vaultToAsset = append(vaultToAsset, TVaultToAsset{
			VaultAddress: token,
			AssetAddress: helpers.DecodeAddress(rawToken),
			Value:        helpers.DecodeBigInt(rawPricePerShare),
		})
	}
	return vaultToAsset
}
