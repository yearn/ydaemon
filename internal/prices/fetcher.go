package prices

import (
	"encoding/json"
	"io/ioutil"
	"math"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"sync"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/store"
	"github.com/yearn/ydaemon/common/traces"
	"github.com/yearn/ydaemon/common/types/common"
	"github.com/yearn/ydaemon/internal/tokens"
)

var priceErrorAlreadySent = make(map[uint64]map[common.Address]bool)

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
		traces.
			Capture(`error`, `missing a valid Lens Address for chain `+strconv.FormatUint(chainID, 10)).
			SetEntity(`prices`).
			SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
			Send()
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
		if len(rawTokenPrice) == 0 {
			continue
		}
		tokenPrice := bigNumber.SetInt(rawTokenPrice[0].(*big.Int))
		if tokenPrice.IsZero() {
			continue
		}
		newPriceMap[token] = helpers.DecodeBigInt(rawTokenPrice)
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

	/**********************************************************************************************
	** We now fill in the missing prices using the DeFiLlama and CoinGecko API.
	**********************************************************************************************/
	queryList := []common.Address{}
	for _, token := range tokenList {
		if newPriceMap[token] == nil || newPriceMap[token].IsZero() {
			queryList = append(queryList, token)
		}
	}
	// Call the two API endpoints async if we are missing prices
	if len(queryList) > 0 {
		chanLlama := make(chan []*bigNumber.Int)
		go func() {
			chanLlama <- fetchPricesFromLlama(chainID, queryList)
		}()
		chanGecko := make(chan []*bigNumber.Int)
		go func() {
			chanGecko <- fetchPricesFromGecko(chainID, queryList)
		}()
		pricesLlama := <-chanLlama
		pricesGecko := <-chanGecko

		for index, token := range queryList {
			priceLlama := pricesLlama[index]
			if priceLlama != nil && !priceLlama.IsZero() {
				newPriceMap[token] = priceLlama
				continue
			}
			priceGecko := pricesGecko[index]
			if priceGecko != nil && !priceGecko.IsZero() {
				newPriceMap[token] = priceGecko
			}
		}
	}

	/**********************************************************************************************
	** Finally, we will list all the tokens that are still missing a price to log them to Sentry.
	**********************************************************************************************/
	if priceErrorAlreadySent[chainID] == nil {
		priceErrorAlreadySent[chainID] = make(map[common.Address]bool)
	}

	for _, token := range queryList {
		if (newPriceMap[token] == nil || newPriceMap[token].IsZero()) && !priceErrorAlreadySent[chainID][token] {
			traces.
				Capture(`error`, `missing a valid price for token `+token.String()+` on chain `+strconv.FormatUint(chainID, 10)).
				SetEntity(`prices`).
				SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
				Send()
			priceErrorAlreadySent[chainID][token] = true
		}
	}

	return newPriceMap
}

/**************************************************************************************************
** fetchPriceFromLlama tries to fetch the price for a given token from
** the DeFiLlama pricing API, returns nil if there is no data returned
**************************************************************************************************/
func fetchPricesFromLlama(chainID uint64, tokens []common.Address) []*bigNumber.Int {
	// Return nil pointers as default
	prices := make([]*bigNumber.Int, len(tokens))

	var tokenString []string
	for _, token := range tokens {
		tokenString = append(tokenString, env.LLAMA_CHAIN_NAMES[chainID]+`:`+token.String())
	}
	resp, err := http.Get(env.LLAMA_PRICE_URL + strings.Join(tokenString, ","))
	if err != nil || resp.StatusCode != 200 {
		logs.Warning("Error fetching prices from DeFiLlama for chain", chainID)
		return prices
	}
	// Defer the closing of the response body to avoid memory leaks
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Warning("Error unmarshalling response body from the pricing API of DeFiLlama for chain", chainID)
		return prices
	}
	priceData := TLlamaPrice{}
	if err := json.Unmarshal(body, &priceData); err != nil {
		logs.Warning("Error unmarshalling response body from the pricing API of DeFiLlama for chain", chainID)
		return prices
	}

	// Parse response
	decimalsUSDC := bigNumber.NewFloat(math.Pow10(6))
	for index, tokenStr := range tokenString {
		data, ok := priceData.Coins[tokenStr]
		// Convert price into USDC decimals
		if ok {
			price := bigNumber.NewFloat(data.Price)
			prices[index] = bigNumber.NewFloat().Mul(price, decimalsUSDC).Int()
		}
	}
	return prices
}

/**************************************************************************************************
** fetchPriceFromGecko tries to fetch the price for a given token from
** the CoinGecko API, returns nil if there is no data returned
**************************************************************************************************/
func fetchPricesFromGecko(chainID uint64, tokens []common.Address) []*bigNumber.Int {
	// Return nil pointers as default
	prices := make([]*bigNumber.Int, len(tokens))

	var tokenString []string
	for _, token := range tokens {
		tokenString = append(tokenString, strings.ToLower(token.String()))
	}
	req, err := http.NewRequest("GET", env.GECKO_PRICE_URL+env.GECKO_CHAIN_NAMES[chainID], nil)
	if err != nil {
		logs.Warning("Error fetching prices from CoinGecko for chain", chainID)
		return prices
	}
	q := req.URL.Query()
	q.Add("contract_addresses", strings.Join(tokenString, ","))
	q.Add("vs_currencies", "usd")
	req.URL.RawQuery = q.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil || resp.StatusCode != 200 {
		logs.Warning("Error fetching prices from DeFiLlama for chain", chainID)
		return prices
	}
	// Defer the closing of the response body to avoid memory leaks
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Warning("Error unmarshalling response body from the API of CoinGecko for chain", chainID)
		return prices
	}
	priceData := TGeckoPrice{}
	if err := json.Unmarshal(body, &priceData); err != nil {
		logs.Warning("Error unmarshalling response body from the API of CoinGecko for chain", chainID)
		return prices
	}

	// Parse response
	decimalsUSDC := bigNumber.NewFloat(math.Pow10(6))
	for index, tokenStr := range tokenString {
		data, ok := priceData[tokenStr]
		// Convert price into USDC decimals
		if ok {
			price := bigNumber.NewFloat(data.USDPrice)
			prices[index] = bigNumber.NewFloat().Mul(price, decimalsUSDC).Int()
		}
	}
	return prices
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
	trace := traces.Init(`app.indexer.prices`).
		SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
		SetTag(`rpcURI`, ethereum.GetRPCURI(chainID)).
		SetTag(`entity`, `prices`).
		SetTag(`subsystem`, `daemon`)
	defer trace.Finish()

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
	extraTokens := map[uint64][]string{
		1: {
			`0x34fe2a45D8df28459d7705F37eD13d7aE4382009`, // yvWBTC
			`0xD533a949740bb3306d119CC777fa900bA034cd52`, // CRV - used by yBribe UI
			`0x090185f2135308BaD17527004364eBcC2D37e5F6`, // Spell - used by yBribe UI
			`0xCdF7028ceAB81fA0C6971208e83fa7872994beE5`, // TNT - used by yBribe UI
		},
		10:    {},
		250:   {},
		42161: {},
	}
	for _, tokenAddress := range extraTokens[chainID] {
		allTokens = append(allTokens, common.HexToAddress(tokenAddress))
	}

	/**********************************************************************************************
	** Some tokens are just useless, errors or not wanted. We will remove them from the list.
	**********************************************************************************************/
	ignoredTokens := map[uint64][]string{
		1: {
			`0x7AB4a7BE740131BdE216521B54ADddD672F44A05`, // nothing
			`0x61f46C65E403429266e8b569F23f70dD75d9BeE7`, // Old lp-yCRV
			`0x8a0889d47f9Aa0Fac1cC718ba34E26b867437880`, // Old st-yCRV
			`0x4c1317326fD8EFDeBdBE5e1cd052010D97723bd6`, // Another old st-yCRV
			`0x2E919d27D515868f3D5Bc9110fa738f9449FC6ad`, // Old yvCurve-yveCRV pool
			`0x7E46fd8a30869aa9ed55af031067Df666EfE87da`, // Old yvecrv-f
			`0x5904BAcE7a9cCab585242e9d22f67C9f2F1BF7E2`, // nothing
			`0x0309A528bBa0394dC4A2Ce59123C52E317A54604`, // Old yCRV-f
			`0xBF7AA989192b020a8d3e1C65a558e123834325cA`, // Irrelevant HBTC yVault
			`0xe92AE2cF5b373c1713eB5855D4D3aF81D8a8aCAE`, // Curve Stax Frax/Temple xLP + LP yVault - Unlisted
			`0x3883f5e181fccaF8410FA61e12b59BAd963fb645`, // Theta: Old Token
		},
		10:  {},
		250: {},
		42161: {
			`0x976a1C749cd8153909e0B04EebE931eF8957b15b`, // PHPTest
			`0xFa247d0D55a324ca19985577a2cDcFC383D87953`, // Philippine Peso (PHP)
		},
	}
	for _, tokenAddress := range ignoredTokens[chainID] {
		allTokens = helpers.RemoveFromArray(allTokens, common.HexToAddress(tokenAddress))
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
	_priceMap[chainID] = priceMap
	return priceMap
}
