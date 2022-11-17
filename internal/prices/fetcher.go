package prices

import (
	"math"
	"math/big"
	"sync"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/store"
	"github.com/yearn/ydaemon/common/types/common"
	"github.com/yearn/ydaemon/internal/tokens"
)

/**************************************************************************************************
** fetchPrices will, for a list of addresses, try to fetch all the prices from the lens price
** oracle. If the price is not available, it will try to fetch it from some external API.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - tokenList: a list of addresses of the tokens we want to fetch the information for
**
** Returns:
** - a map of tokenAddress -> *bigNumber.Int
**************************************************************************************************/
func fetchPrices(chainID uint64, tokenList []common.Address) map[common.Address]*bigNumber.Int {
	newPriceMap := make(map[common.Address]*bigNumber.Int)

	/**********************************************************************************************
	** The first step is to prepare the multicall, connecting to the multicall instance and
	** preparing the array of calls to send. All calls for all tokens will be send in a single
	** multicall and will later be accessible via a concatened string `tokenAddress + methodName`.
	**********************************************************************************************/
	caller := ethereum.MulticallClientForChainID[chainID]
	lensAddress, ok := env.LENS_ADDRESSES[chainID]
	if !ok {
		logs.Error(`Lens address not found for chainID: `, chainID)
		return newPriceMap
	}

	calls := []ethereum.Call{}
	for _, token := range tokenList {
		calls = append(calls, getPriceUsdcRecommendedCall(token.String(), lensAddress, token))
	}

	/**********************************************************************************************
	** Regular fix for Fantom's RPC, which limit the number of calls in a multicall to a very low
	** number. We split the multicall in multiple calls of 3 calls each.
	** Otherwise, we just send the multicall as is.
	**********************************************************************************************/
	maxBatch := math.MaxInt64
	if chainID == 250 {
		maxBatch = 3
	}

	/**********************************************************************************************
	** Then we can proceed the responses. We loop over the responses and check if the price is
	** available. If it is, we add it to the map. If it's not, we try to fetch it from an external
	** API.
	**********************************************************************************************/
	response := caller.ExecuteByBatch(calls, maxBatch, nil)
	for _, token := range tokenList {
		rawTokenPrice := response[token.String()+`getPriceUsdcRecommended`]
		tokenPrice := bigNumber.SetInt(rawTokenPrice[0].(*big.Int))
		if tokenPrice.IsZero() {
			continue
		}
		newPriceMap[token] = bigNumber.SetInt(rawTokenPrice[0].(*big.Int))
	}

	/**********************************************************************************************
	** Once this is done, we will probably have some missing tokens. We can use the Curve API to
	** be able to calculate the price of some tokens. We will then add them to our map.
	**********************************************************************************************/
	priceMapFromCurveFactoryAPI := getPricesFromCurveFactoriesAPI(chainID)
	for token, price := range priceMapFromCurveFactoryAPI {
		if !price.IsZero() && newPriceMap[token] == nil {
			newPriceMap[token] = price
		}
	}

	return newPriceMap
}

/**************************************************************************************************
** findAllPrices is simply a wrapper around our fetchPrices function to make it easier
** to read. It will call the fetchPrices function to get the prices, before assigning them to a new
** map.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - tokenList: our list of tokenAddress
**
** Returns:
** - a map of tokenAddress -> *bigNumber.Int
**************************************************************************************************/
func findAllPrices(
	chainID uint64,
	tokenList []common.Address,
) map[common.Address]*bigNumber.Int {
	newPriceMap := fetchPrices(chainID, tokenList)

	return newPriceMap
}

/**************************************************************************************************
** In order to correctly play with financial data, we need the prices. This is what this internal
** package is for. It will fetch the prices of the tokens and store them for later use.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - vaults: the array of TVaultsFromRegistry
**
** Returns:
** - a map of tokenAddress -> USDC Price
**************************************************************************************************/
func RetrieveAllPrices(chainID uint64) map[ethcommon.Address]*bigNumber.Int {
	timeBefore := time.Now()

	/**********************************************************************************************
	** First, try to retrieve the list of prices from the database to exclude the one existing
	** from the upcoming calls
	**********************************************************************************************/
	priceMap := make(map[ethcommon.Address]*bigNumber.Int)
	store.Iterate(chainID, store.TABLES.PRICES, &priceMap)

	/**********************************************************************************************
	** From the vault registry we could build the list of tokens used by our ecosystem. We will
	** use this list to fetch the prices of the tokens, using different sources.
	**********************************************************************************************/
	allTokens := tokens.ListTokensAddresses(chainID)

	/**********************************************************************************************
	** Somehow, some vaults are not in the registries, but we still need the price data for them.
	** We will add them manually here.
	**********************************************************************************************/
	extraTokens := []string{
		`0x34fe2a45D8df28459d7705F37eD13d7aE4382009`, // yvWBTC
		`0xD533a949740bb3306d119CC777fa900bA034cd52`, // CRV - used by yBribe UI
		`0x090185f2135308BaD17527004364eBcC2D37e5F6`, // Spell - used by yBribe UI
		`0xCdF7028ceAB81fA0C6971208e83fa7872994beE5`, // TNT - used by yBribe UI
	}
	for _, tokenAddress := range extraTokens {
		allTokens = append(allTokens, common.HexToAddress(tokenAddress))
	}
	allTokens = helpers.UniqueArrayAddress(allTokens)

	/**********************************************************************************************
	** Once we got out basic list, we will fetch the price for each of theses tokens, using lens
	** as primary source, and multiple other sources as fallback.
	**********************************************************************************************/
	if len(allTokens) > 0 {
		allPrices := findAllPrices(chainID, allTokens)

		/**********************************************************************************************
		** Once everything is setup, we will store each token in the DB. The storage is set as a map
		** of tokenAddress -> USDC Price. All tokens will be retrievable from the store.Interate() func.
		**********************************************************************************************/
		wg := sync.WaitGroup{}
		wg.Add(len(allPrices))
		for address, price := range allPrices {
			go func(address common.Address, price *bigNumber.Int) {
				defer wg.Done()
				store.SaveInDB(
					chainID,
					store.TABLES.PRICES,
					address.String(),
					price,
				)
			}(address, price)
		}
		wg.Wait()
		store.Iterate(chainID, store.TABLES.PRICES, &priceMap)
	}

	logs.Success(`It took`, time.Since(timeBefore), `to retrieve`, len(priceMap), `prices`)
	_priceMap[chainID] = priceMap
	return priceMap
}
