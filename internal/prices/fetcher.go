package prices

import (
	"context"
	"math"
	"math/big"
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/store"
	"github.com/yearn/ydaemon/common/traces"
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
func fetchPrices(chainID uint64, blockNumber *uint64, tokenList []common.Address) map[common.Address]*bigNumber.Int {
	newPriceMap := fetchPricesFromLens(chainID, blockNumber, tokenList)

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
		chanLlama := make(chan map[common.Address]*bigNumber.Int)
		go func() {
			chanLlama <- fetchPricesFromLlama(chainID, queryList)
		}()
		chanGecko := make(chan map[common.Address]*bigNumber.Int)
		go func() {
			chanGecko <- fetchPricesFromGecko(chainID, queryList)
		}()
		pricesLlama := <-chanLlama
		pricesGecko := <-chanGecko

		for _, token := range queryList {
			if pricesLlama[token] != nil && !pricesLlama[token].IsZero() {
				newPriceMap[token] = pricesLlama[token]
				store.StoreHistoricalPrice(chainID, *blockNumber, token, pricesLlama[token])
				continue
			}
			if pricesGecko[token] != nil && !pricesGecko[token].IsZero() {
				newPriceMap[token] = pricesGecko[token]
				store.StoreHistoricalPrice(chainID, *blockNumber, token, pricesGecko[token])
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
	blockNumber *uint64,
	tokenList []common.Address,
) map[common.Address]*bigNumber.Int {
	newPriceMap := fetchPrices(chainID, blockNumber, tokenList)

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
func RetrieveAllPrices(chainID uint64) map[common.Address]*bigNumber.Int {
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
	priceMap := make(map[common.Address]*bigNumber.Int)
	store.ListFromBadgerDB(chainID, store.TABLES.PRICES, &priceMap)

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
	currentBlockNumber, _ := ethereum.GetRPC(chainID).BlockNumber(context.Background())
	if len(allTokens) > 0 {
		allPrices := findAllPrices(chainID, &currentBlockNumber, allTokens)

		/**********************************************************************************************
		** Once everything is setup, we will store each token in the DB. The storage is set as a map
		** of tokenAddress -> USDC Price. All tokens will be retrievable from the store.Interate() func.
		**********************************************************************************************/
		wg := sync.WaitGroup{}
		wg.Add(len(allPrices))
		for address, price := range allPrices {
			go func(address common.Address, price *bigNumber.Int) {
				defer wg.Done()
				store.SaveInBadgerDB(
					chainID,
					store.TABLES.PRICES,
					address.String(),
					price,
				)
			}(address, price)
		}
		wg.Wait()
		store.ListFromBadgerDB(chainID, store.TABLES.PRICES, &priceMap)
	}
	_priceMap[chainID] = priceMap
	return priceMap
}

/*
*************************************************************************************************

*************************************************************************************************
 */
func FetchPricesOnBlock(chainID uint64, blockNumber uint64, tokenList []common.Address) map[common.Address]*bigNumber.Int {
	newPriceMap := make(map[common.Address]*bigNumber.Int)

	/**********************************************************************************************
	** The first step is to prepare the multicall, connecting to the multicall instance and
	** preparing the array of calls to send. All calls for all tokens will be send in a single
	** multicall and will later be accessible via a concatened string `tokenAddress + methodName`.
	**********************************************************************************************/
	caller := ethereum.MulticallClientForChainID[chainID]
	lensAddress := env.LENS_ADDRESSES[chainID]
	calls := []ethereum.Call{}
	for _, tokenAddress := range tokenList {
		if tokenPrice, ok := store.GetBlockPrice(chainID, blockNumber, tokenAddress); ok {
			newPriceMap[tokenAddress] = tokenPrice
			continue
		}
		calls = append(calls, getPriceUsdcRecommendedCall(tokenAddress.String(), lensAddress, tokenAddress))
	}

	/**********************************************************************************************
	** Then we can proceed the responses. We loop over the responses and check if the price is
	** available. If it is, we add it to the map. If it's not, we try to fetch it from an external
	** API.
	**********************************************************************************************/
	response := caller.ExecuteByBatch(calls, math.MaxInt64, big.NewInt(int64(blockNumber)))
	for _, tokenAddress := range tokenList {
		rawTokenPrice := response[tokenAddress.String()+`getPriceUsdcRecommended`]
		if len(rawTokenPrice) == 0 {
			continue
		}
		tokenPrice := bigNumber.SetInt(rawTokenPrice[0].(*big.Int))
		if tokenPrice.IsZero() {
			continue
		}
		newPriceMap[tokenAddress] = helpers.DecodeBigInt(rawTokenPrice)
		store.StoreHistoricalPrice(chainID, blockNumber, tokenAddress, newPriceMap[tokenAddress])
	}

	return newPriceMap
}

func GetPricesOnBlock(chainID uint64, blockNumber uint64, tokenAddress common.Address) (*bigNumber.Int, bool) {
	tokenPrice, ok := store.GetBlockPrice(chainID, blockNumber, tokenAddress)
	if ok {
		return tokenPrice, true
	}

	client := ethereum.GetRPC(chainID)
	contractAddress := env.LENS_ADDRESSES[chainID]
	oracleContract, _ := contracts.NewOracleCaller(contractAddress, client)
	tokenPriceFromOracle, err := oracleContract.GetPriceUsdcRecommended(&bind.CallOpts{BlockNumber: big.NewInt(int64(blockNumber))}, tokenAddress)
	if err != nil {
		return bigNumber.NewInt(0), false
	}
	store.StoreHistoricalPrice(chainID, blockNumber, tokenAddress, bigNumber.NewInt(0).Set(tokenPriceFromOracle))
	return bigNumber.NewInt(0).Set(tokenPriceFromOracle), true
}
