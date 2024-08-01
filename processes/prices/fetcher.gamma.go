package prices

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

/**************************************************************************************************
** getGammaLPPricesFromAPI will fetch the prices of the tokens from the Gamma API. We will use the
** RetrieveGammaAllData function from the storage package to get the list of pools available on
** Gamma, with enough data to calculate the price of the tokens in USD.
**************************************************************************************************/
func getGammaLPPricesFromAPI(chainID uint64, blockNumber *uint64, tokens []models.TERC20Token) map[common.Address]models.TPrices {
	priceMap := make(map[common.Address]models.TPrices)

	if chainID != 137 {
		return priceMap
	}

	pools, ok := storage.GetCachedGammaAllData(chainID)
	if !ok {
		pools, ok = storage.RetrieveGammaAllData(chainID)
		if !ok {
			return priceMap
		}
	}

	for addr, poolData := range pools {
		totalTVLFloat := bigNumber.NewFloat(0).SetString(poolData.TvlUSD)
		totalSupplyFloat := bigNumber.NewFloat(0)
		switch poolData.TotalSupply.(type) {
		case string:
			totalSupply := bigNumber.NewInt(0).SetString(poolData.TotalSupply.(string))
			totalSupplyFloat = helpers.ToNormalizedAmount(totalSupply, 18)
		case *big.Int:
			totalSupply := bigNumber.NewInt(0).Set(poolData.TotalSupply.(*big.Int))
			totalSupplyFloat = helpers.ToNormalizedAmount(totalSupply, 18)
		case float64:
			totalSupply := bigNumber.NewFloat(0).SetFloat64(poolData.TotalSupply.(float64)).Int()
			totalSupplyFloat = helpers.ToNormalizedAmount(totalSupply, 18)
		}

		tokenPrice := bigNumber.NewFloat(0).Div(totalTVLFloat, totalSupplyFloat)
		rawPrice := bigNumber.NewFloat(0).Mul(tokenPrice, bigNumber.NewFloat(1e6)).Int()

		priceMap[common.HexToAddress(addr)] = models.TPrices{
			Address:        common.HexToAddress(addr),
			Price:          rawPrice,
			HumanizedPrice: tokenPrice,
			Source:         `gammaAPI`,
		}
	}
	return priceMap
}
