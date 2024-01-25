package prices

import (
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/multicalls"
)

/**************************************************************************************************
** fetchPriceFromLlama tries to fetch the price for a given token from
** the DeFiLlama pricing API, returns nil if there is no data returned
**************************************************************************************************/
func fetchPricesFromLens(chainID uint64, blockNumber *uint64, tokens []models.TERC20Token) map[common.Address]models.TPrices {
	priceMap := make(map[common.Address]models.TPrices)

	/**********************************************************************************************
	** The first step is to prepare the multicall, connecting to the multicall instance and
	** preparing the array of calls to send. All calls for all tokens will be send in a single
	** multicall and will later be accessible via a concatened string `tokenAddress + methodName`.
	**********************************************************************************************/
	lensAddress := env.CHAINS[chainID].LensContract.Address
	if (lensAddress == common.Address{}) {
		logs.Error(`missing a valid Lens Address for chain ` + strconv.FormatUint(chainID, 10))
		return priceMap
	}

	calls := []ethereum.Call{}
	for _, token := range tokens {
		calls = append(calls, multicalls.GetPriceUsdcRecommendedCall(token.Address.Hex(), lensAddress, token.Address))
	}
	if chainID == 1 {
		crvUSD := common.HexToAddress(`0xe5Afcf332a5457E8FafCD668BcE3dF953762Dfe7`)
		calls = append(calls, multicalls.GetPriceCrvUsdcCall(crvUSD.Hex(), crvUSD))
	}

	/**********************************************************************************************
	** Then we can proceed the responses. We loop over the responses and check if the price is
	** available. If it is, we add it to the map. If it's not, we try to fetch it from an external
	** API.
	**********************************************************************************************/
	response := multicalls.Perform(chainID, calls, nil)

	for _, token := range tokens {
		rawTokenPrice := response[token.Address.Hex()+`getPriceUsdcRecommended`]
		if len(rawTokenPrice) == 0 {
			continue
		}
		tokenPrice := bigNumber.SetInt(rawTokenPrice[0].(*big.Int))
		if tokenPrice.IsZero() {
			continue
		}

		price := helpers.DecodeBigInt(rawTokenPrice)
		humanizedPrice := helpers.ToNormalizedAmount(price, 6)
		priceMap[token.Address] = models.TPrices{
			Address:        token.Address,
			Price:          price,
			HumanizedPrice: humanizedPrice,
			Source:         `lens`,
		}
	}

	if chainID == 1 {
		crvUSD := common.HexToAddress(`0xe5Afcf332a5457E8FafCD668BcE3dF953762Dfe7`)
		rawTokenPrice := response[crvUSD.Hex()+`price`]
		if len(rawTokenPrice) > 0 {
			tokenPrice := bigNumber.SetInt(rawTokenPrice[0].(*big.Int))
			if tokenPrice.Gt(bigNumber.Zero) {
				tokenPrice = tokenPrice.Div(tokenPrice, bigNumber.NewInt(1e12))
				humanizedPrice := helpers.ToNormalizedAmount(tokenPrice, 6)
				priceMap[crvUSD] = models.TPrices{
					Address:        crvUSD,
					Price:          tokenPrice,
					HumanizedPrice: humanizedPrice,
					Source:         `lens`,
				}

			}
		}
	}
	return priceMap
}
