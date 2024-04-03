package prices

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

/**************************************************************************************************
** getPendlePricesFromAPI will fetch the prices of the tokens from the Pendle API. We will use the
** RetrivePendleTokens function from the storage package to get the list of tokens available on
** Pendle, with the price, calculated in USD by Pendle, available in the struct.
**************************************************************************************************/
func getPendlePricesFromAPI(chainID uint64, blockNumber *uint64, tokens []models.TERC20Token) map[common.Address]models.TPrices {
	priceMap := make(map[common.Address]models.TPrices)
	pendleTokens, ok := storage.GetCachedPendleTokens(chainID)
	if !ok {
		pendleTokens, ok = storage.RetrievePendleTokens(chainID)
		if !ok {
			return priceMap
		}
	}

	for _, tokenData := range pendleTokens {
		tokenPrice := bigNumber.NewFloat(tokenData.Price.USD)
		rawPrice := bigNumber.NewFloat(0).Mul(tokenPrice, bigNumber.NewFloat(1e6)).Int()
		priceMap[common.HexToAddress(tokenData.Address)] = models.TPrices{
			Address:        common.HexToAddress(tokenData.Address),
			Price:          rawPrice,
			HumanizedPrice: tokenPrice,
			Source:         `pendleAPI`,
		}
	}
	return priceMap
}
