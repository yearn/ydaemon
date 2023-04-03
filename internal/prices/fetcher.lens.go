package prices

import (
	"context"
	"math"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/store"
	"github.com/yearn/ydaemon/common/traces"
)

/**************************************************************************************************
** fetchPriceFromLlama tries to fetch the price for a given token from
** the DeFiLlama pricing API, returns nil if there is no data returned
**************************************************************************************************/
func fetchPricesFromLens(chainID uint64, blockNumber *uint64, tokens []common.Address) map[common.Address]*bigNumber.Int {
	newPriceMap := make(map[common.Address]*bigNumber.Int)

	/**********************************************************************************************
	** The first step is to prepare the multicall, connecting to the multicall instance and
	** preparing the array of calls to send. All calls for all tokens will be send in a single
	** multicall and will later be accessible via a concatened string `tokenAddress + methodName`.
	**********************************************************************************************/
	caller := ethereum.MulticallClientForChainID[chainID]
	lensAddress, ok := env.LENS_ADDRESSES[chainID]
	if !ok {
		traces.
			Capture(`error`, `missing a valid Lens Address for chain `+strconv.FormatUint(chainID, 10)).
			SetEntity(`prices`).
			SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
			Send()
		return newPriceMap
	}

	calls := []ethereum.Call{}
	for _, token := range tokens {
		calls = append(calls, getPriceUsdcRecommendedCall(token.String(), lensAddress, token))
	}

	/**********************************************************************************************
	** Regular fix for Fantom's RPC, which limit the number of calls in a multicall to a very low
	** number. We split the multicall in multiple calls of 3 calls each.
	** Otherwise, we just send the multicall as is.
	**********************************************************************************************/
	maxBatch := math.MaxInt64

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
		response = caller.ExecuteByBatch(calls, maxBatch, nil)
	} else {
		blockNumberBigInt = big.NewInt(int64(*blockNumber))
		response = caller.ExecuteByBatch(calls, maxBatch, blockNumberBigInt)

	}

	for _, token := range tokens {
		rawTokenPrice := response[token.String()+`getPriceUsdcRecommended`]
		if len(rawTokenPrice) == 0 {
			continue
		}
		tokenPrice := bigNumber.SetInt(rawTokenPrice[0].(*big.Int))
		if tokenPrice.IsZero() {
			continue
		}
		newPriceMap[token] = helpers.DecodeBigInt(rawTokenPrice)
		store.StoreHistoricalPrice(chainID, *blockNumber, token, tokenPrice)
	}
	return newPriceMap
}
