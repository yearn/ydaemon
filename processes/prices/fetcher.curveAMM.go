package prices

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/multicalls"
)

/**************************************************************************************************
** fetchPricesFromCurveAMM
**************************************************************************************************/
func fetchPricesFromCurveAMM(chainID uint64, blockNumber *uint64, tokens []common.Address) map[common.Address]*bigNumber.Int {
	newPriceMap := make(map[common.Address]*bigNumber.Int)

	calls := []ethereum.Call{}
	for _, token := range tokens {
		calls = append(calls, multicalls.GetLPPrice(token.Hex(), token))
		calls = append(calls, multicalls.GetDecimals(token.Hex(), token))
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
		rawTokenPrice := response[token.Hex()+`lp_price`]
		rawDecimals := response[token.Hex()+`decimals`]
		if len(rawTokenPrice) == 0 {
			continue
		}
		decimals := helpers.DecodeUint64(rawDecimals)
		bigTokenPrice := bigNumber.SetInt(rawTokenPrice[0].(*big.Int))
		if bigTokenPrice.IsZero() {
			continue
		}
		tokenPriceUSD := helpers.ToNormalizedAmount(bigTokenPrice, decimals)
		tokenPrice := bigNumber.NewFloat(0).Mul(tokenPriceUSD, bigNumber.NewFloat(1e6)).Int()

		newPriceMap[token] = tokenPrice
	}
	return newPriceMap
}
