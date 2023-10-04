package prices

import (
	"context"
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/store"
	"github.com/yearn/ydaemon/internal/prices"
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
func fetchPrices(
	chainID uint64,
	blockNumber *uint64,
	tokenList []common.Address,
) map[common.Address]*bigNumber.Int {
	newPriceMap := make(map[common.Address]*bigNumber.Int)

	/**********************************************************************************************
	** We now fill in the missing prices using the DeFiLlama and CoinGecko API.
	**********************************************************************************************/
	chanLlama := make(chan map[common.Address]*bigNumber.Int)
	go func() {
		chanLlama <- fetchPricesFromLlama(chainID, tokenList)
	}()
	chanGecko := make(chan map[common.Address]*bigNumber.Int)
	go func() {
		chanGecko <- fetchPricesFromGecko(chainID, tokenList)
	}()
	pricesLlama := <-chanLlama
	pricesGecko := <-chanGecko

	for _, token := range tokenList {
		if pricesLlama[token] != nil && !pricesLlama[token].IsZero() {
			newPriceMap[token] = pricesLlama[token]
			continue
		}
		if pricesGecko[token] != nil && !pricesGecko[token].IsZero() {
			newPriceMap[token] = pricesGecko[token]
		}
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
	** Once this is done, we will probably have some missing tokens. We can use the Velo API to
	** be able to calculate the price of some tokens. We will then add them to our map. Only on
	** optimism
	**********************************************************************************************/
	priceMapFromVeloPairsAPI := getPricesFromVeloPairsAPI(chainID)
	for token, price := range priceMapFromVeloPairsAPI {
		if !price.IsZero() && newPriceMap[token] == nil {
			newPriceMap[token] = price
		}
	}

	priceMapFromVeloOracle := fetchPricesFromSugar(chainID, blockNumber, tokenList)
	for token, price := range priceMapFromVeloOracle {
		if !price.IsZero() && newPriceMap[token] == nil {
			newPriceMap[token] = price
		}
	}

	/**********************************************************************************************
	** Once this is done, we will probably have some missing tokens. We can use the Aero API to
	** be able to calculate the price of some tokens. We will then add them to our map. Only on
	** Base
	**********************************************************************************************/
	priceMapFromAeroOracle := fetchPricesFromAeroSugar(chainID, blockNumber, tokenList)
	for token, price := range priceMapFromAeroOracle {
		if !price.IsZero() && newPriceMap[token] == nil {
			newPriceMap[token] = price
		}
	}

	/**********************************************************************************************
	** If we still have some missing prices, we will use the lens price oracle to fetch them.
	**********************************************************************************************/
	queryList := []common.Address{}
	for _, token := range tokenList {
		if newPriceMap[token] == nil || newPriceMap[token].IsZero() {
			queryList = append(queryList, token)
		}
	}
	priceMapLensOracle := fetchPricesFromLens(chainID, blockNumber, queryList)

	for token, price := range priceMapLensOracle {
		if !price.IsZero() && newPriceMap[token] == nil {
			newPriceMap[token] = price
		}
	}

	/**********************************************************************************************
	** With the ERC-4626 standard, the `price per share` is no longer relevant. We can use the new
	** `convertToAssets` function to get the value of the underlying asset for a given amount of
	** shares: 1 share will give me 1.23 asset for example.
	** Based on that, if we have the price of the underlying asset, we can calculate the price of
	** the share.
	**********************************************************************************************/
	sharesValue := fetchShareValueFromERC4626(chainID, blockNumber, queryList)
	for _, shareValue := range sharesValue {
		if shareValue.Value == nil || shareValue.Value.IsZero() {
			continue
		}
		if newPriceMap[shareValue.AssetAddress] == nil || newPriceMap[shareValue.AssetAddress].IsZero() {
			continue
		}

		token, ok := store.GetERC20(chainID, shareValue.AssetAddress)
		if !ok {
			continue
		}

		tokenDecimals := helpers.ToRawAmount(bigNumber.NewInt(1), token.Decimals)
		sharePrice := bigNumber.NewFloat(0).Quo(
			bigNumber.NewFloat(0).Mul(
				bigNumber.NewFloat(0).SetInt(shareValue.Value),
				bigNumber.NewFloat(0).SetInt(newPriceMap[shareValue.AssetAddress]),
			),
			bigNumber.NewFloat(0).SetInt(tokenDecimals),
		)
		newPriceMap[shareValue.VaultAddress] = sharePrice.Int()
	}

	/**********************************************************************************************
	** If the price is missing, check if it's a vault and try to compute the price from the
	** underlying tokens.
	**********************************************************************************************/
	for _, token := range tokenList {
		if newPriceMap[token] == nil || newPriceMap[token].IsZero() {
			if token, ok := store.GetERC20(chainID, token); ok {
				if store.IsVaultLike(token) {
					ppsPerTime, _ := store.ListPricePerShare(chainID, token.Address)
					underlyingToken := token.UnderlyingTokensAddresses[0]
					ppsToday := helpers.GetToday(ppsPerTime, token.Decimals)
					underlyingPrice := bigNumber.NewFloat(0).SetInt(newPriceMap[underlyingToken])
					vaultPrice := bigNumber.NewFloat(0).Mul(ppsToday, underlyingPrice)
					newPriceMap[token.Address] = vaultPrice.Int()
				}
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
			if tokenData, ok := store.GetERC20(chainID, token); ok {
				logs.Warning(`missing a valid price for ` + tokenData.Address.Hex() + ` (` + tokenData.Name + `)` + ` on chain ` + strconv.FormatUint(chainID, 10) + ` (` + string(tokenData.Type) + `)`)
			} else {
				logs.Warning(`missing a valid price for token ` + token.Hex() + ` on chain ` + strconv.FormatUint(chainID, 10))
			}
			priceErrorAlreadySent[chainID][token] = true
		}
	}

	itemsToUpsert := []store.DBHistoricalPrice{}
	for tokenAddress, price := range newPriceMap {
		/******************************************************************************************
		** To save some process, we will batch the saving to the database in one call.
		** The following code is creating the upsert array.
		******************************************************************************************/
		DBbaseSchema := store.DBBaseSchema{
			UUID:    store.GetUUID(strconv.FormatUint(chainID, 10) + strconv.FormatUint(*blockNumber, 10) + tokenAddress.Hex()),
			Block:   *blockNumber,
			ChainID: chainID,
		}
		humanizedPrice, _ := helpers.ToNormalizedAmount(price, 6).Float64()
		itemsToUpsert = append(itemsToUpsert, store.DBHistoricalPrice{
			DBBaseSchema:   DBbaseSchema,
			Token:          tokenAddress.Hex(),
			Price:          price.String(),
			HumanizedPrice: humanizedPrice,
		})
		store.AppendToHistoricalPrice(chainID, *blockNumber, tokenAddress, price)
	}

	store.StoreManyHistoricalPrice(itemsToUpsert)
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
func retrieveAllPrices(chainID uint64) map[common.Address]*bigNumber.Int {
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
	allTokens := store.ListERC20Addresses(chainID)

	/**********************************************************************************************
	** Somehow, some vaults are not in the registries, but we still need the price data for them.
	** We will add them manually here.
	**********************************************************************************************/
	allTokens = append(allTokens, env.CHAINS[chainID].ExtraTokens...)

	/**********************************************************************************************
	** Some tokens are just useless, errors or not wanted. We will remove them from the list.
	**********************************************************************************************/
	for _, tokenAddress := range env.CHAINS[chainID].IgnoredTokens {
		allTokens = helpers.RemoveFromArray(allTokens, tokenAddress)
	}
	allTokens = helpers.UniqueArrayAddress(allTokens)

	/**********************************************************************************************
	** Once we got out basic list, we will fetch the price for each of theses tokens, using lens
	** as primary source, and multiple other sources as fallback.
	**********************************************************************************************/
	currentBlockNumber, _ := ethereum.GetRPC(chainID).BlockNumber(context.Background())
	if len(allTokens) > 0 {
		allPrices := fetchPrices(chainID, &currentBlockNumber, allTokens)

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
					address.Hex(),
					price,
				)
			}(address, price)
		}
		wg.Wait()
		store.ListFromBadgerDB(chainID, store.TABLES.PRICES, &priceMap)
	}
	prices.StorePrices(chainID, priceMap)
	return priceMap
}

func Run(chainID uint64) {
	retrieveAllPrices(chainID)
}
