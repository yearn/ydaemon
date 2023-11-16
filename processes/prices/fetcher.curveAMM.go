package prices

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/multicalls"
)

/**************************************************************************************************
** fetchPricesFromCurveAMM
**************************************************************************************************/
func fetchPricesFromCurveAMM(chainID uint64, blockNumber *uint64, tokens []models.TERC20Token) map[common.Address]models.TPrices {
	priceMap := make(map[common.Address]models.TPrices)

	calls := []ethereum.Call{}
	for _, token := range tokens {
		calls = append(calls, multicalls.GetLPPrice(token.Address.Hex(), token.Address))
		calls = append(calls, multicalls.GetDecimals(token.Address.Hex(), token.Address))
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

	for _, token := range tokens {
		rawTokenPrice := response[token.Address.Hex()+`lp_price`]
		rawDecimals := response[token.Address.Hex()+`decimals`]
		if len(rawTokenPrice) == 0 {
			continue
		}
		decimals := helpers.DecodeUint64(rawDecimals)
		bigTokenPrice := helpers.DecodeBigInt(rawTokenPrice)
		if bigTokenPrice.IsZero() {
			continue
		}

		//todo: fix this
		tokenPriceUSD := helpers.ToNormalizedAmount(bigTokenPrice, decimals)
		// tokenPrice := bigNumber.NewFloat(0).Mul(tokenPriceUSD, bigNumber.NewFloat(1e6)).Int()
		// humanizedPrice := helpers.ToNormalizedAmount(bigTokenPrice, 6)
		priceMap[token.Address] = models.TPrices{
			Address:        token.Address,
			Price:          bigTokenPrice,
			HumanizedPrice: tokenPriceUSD,
			Source:         `curveAMM`,
		}
	}
	return priceMap
}
