package prices

import (
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/multicalls"
)

var NO_MULTICALLS = true

/**************************************************************************************************
** fetchPriceFromLlama tries to fetch the price for a given token from
** the DeFiLlama pricing API, returns nil if there is no data returned
**************************************************************************************************/
func fetchPricesFromLens(chainID uint64, blockNumber *uint64, tokens []models.TERC20Token) map[common.Address]models.TPrices {
	priceMap := make(map[common.Address]models.TPrices)
	chain, ok := env.GetChain(chainID)
	if !ok {
		return priceMap
	}

	lensAddress := chain.LensContract.Address
	if (lensAddress == common.Address{}) {
		return priceMap
	}

	if true {
		return priceMap
	}

	/**********************************************************************************************
	** The lens contract is HEAVY. Like really heavy. On optimism, it can takes up to 10 minutes
	** to fetch all the ~250 prices we need. Using the following code reduces the time to a few s
	** for the same amount of data.
	**********************************************************************************************/
	if NO_MULTICALLS {
		client := ethereum.GetRPC(chainID)
		lensContract, err := contracts.NewOracleCaller(lensAddress, client)
		if err != nil {
			logs.Error(`error fetching lens contract`, err)
			return priceMap
		}

		/******************************************************************************************
		** To avoid DDOS the node, we are grouping the tokens by group of 40 tokens and running the
		** batch call for each group. This is a bit slower than all as goroutines, but it's safer.
		******************************************************************************************/
		grouped := [][]models.TERC20Token{}
		for i := 0; i < len(tokens); i += 500 {
			end := i + 500
			if end > len(tokens) {
				end = len(tokens)
			}
			grouped = append(grouped, tokens[i:end])
		}

		var allPriceMap sync.Map
		// var wg sync.WaitGroup
		// for _, tokens := range grouped {
		// 	wg.Add(1)
		// 	go func(tokens []models.TERC20Token) {
		// 		defer wg.Done()

		// 		for _, token := range tokens {
		// 			price, err := lensContract.GetPriceUsdcRecommended(nil, token.Address)
		// 			if err != nil {
		// 				logs.Error(err, ethereum.GetRPCURI(chainID))
		// 				return
		// 			}
		// 			if price.Cmp(big.NewInt(0)) > 0 {
		// 				humanizedPrice := helpers.ToNormalizedAmount(bigNumber.SetInt(price), 6)
		// 				allPriceMap.Store(token.Address, models.TPrices{
		// 					Address:        token.Address,
		// 					Price:          bigNumber.SetInt(price),
		// 					HumanizedPrice: humanizedPrice,
		// 					Source:         `lens`,
		// 				})
		// 			}
		// 		}
		// 	}(tokens)
		// }
		// wg.Wait()
		for _, tokens := range grouped {
			for _, token := range tokens {
				price, err := lensContract.GetPriceUsdcRecommended(nil, token.Address)
				if err != nil {
					logs.Error(`error fetching lens contract`, err)
					continue
				}
				if price.Cmp(big.NewInt(0)) > 0 {
					humanizedPrice := helpers.ToNormalizedAmount(bigNumber.SetInt(price), 6)
					allPriceMap.Store(token.Address, models.TPrices{
						Address:        token.Address,
						Price:          bigNumber.SetInt(price),
						HumanizedPrice: humanizedPrice,
						Source:         `lens`,
					})
				}
			}
		}

		allPriceMap.Range(func(key, value interface{}) bool {
			priceMap[key.(common.Address)] = value.(models.TPrices)
			return true
		})
	}

	/**********************************************************************************************
	** The first step is to prepare the multicall, connecting to the multicall instance and
	** preparing the array of calls to send. All calls for all tokens will be send in a single
	** multicall and will later be accessible via a concatened string `tokenAddress + methodName`.
	**********************************************************************************************/
	calls := []ethereum.Call{}
	if !NO_MULTICALLS {
		for _, token := range tokens {
			calls = append(calls, multicalls.GetPriceUsdcRecommendedCall(token.Address.Hex(), lensAddress, token.Address))
		}
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
