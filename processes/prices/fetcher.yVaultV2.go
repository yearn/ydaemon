package prices

import (
	"math/big"

	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/multicalls"
	"github.com/yearn/ydaemon/internal/storage"
)

/**************************************************************************************************
** fetchPricePerShareFromVault will try to get the price per share for a vault type token
** It will return an array of struct with vault/asset/value
**************************************************************************************************/
func fetchPricePerShareFromVault(chainID uint64, blockNumber *uint64, tokens []models.TERC20Token) []TVaultToAsset {
	vaultToAsset := []TVaultToAsset{}

	/**********************************************************************************************
	** The first step is to prepare the multicall, connecting to the multicall instance and
	** preparing the array of calls to send. All calls for all tokens will be send in a single
	** multicall and will later be accessible via a concatenated string `tokenAddress + methodName`.
	**********************************************************************************************/
	calls := []ethereum.Call{}
	for _, token := range tokens {
		if _, ok := storage.GetERC20(chainID, token.Address); ok {
			calls = append(calls, multicalls.GetPricePerShare(token.Address.Hex(), token.Address))
			calls = append(calls, multicalls.GetToken(token.Address.Hex(), token.Address))
		}
	}

	/**********************************************************************************************
	** Then we can proceed the responses. We loop over the responses and check if the price is
	** available. If it is, we add it to the map. If it's not, we try to fetch it from an external
	** API.
	**********************************************************************************************/
	response := multicalls.Perform(chainID, calls, nil)

	for _, token := range tokens {
		rawPricePerShare := response[token.Address.Hex()+`pricePerShare`]
		rawToken := response[token.Address.Hex()+`token`]
		if len(rawPricePerShare) == 0 || len(rawToken) == 0 {
			continue
		}
		tokenPrice := bigNumber.SetInt(rawPricePerShare[0].(*big.Int))
		if tokenPrice.IsZero() {
			continue
		}

		vaultToAsset = append(vaultToAsset, TVaultToAsset{
			VaultAddress: token.Address,
			AssetAddress: helpers.DecodeAddress(rawToken),
			Value:        helpers.DecodeBigInt(rawPricePerShare),
		})
	}
	return vaultToAsset
}
