package prices

import (
	"context"
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
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
	extraTokens := map[uint64][]string{
		1: {
			`0x34fe2a45D8df28459d7705F37eD13d7aE4382009`, // yvWBTC
			`0xD533a949740bb3306d119CC777fa900bA034cd52`, // CRV - used by yBribe UI
			`0x090185f2135308BaD17527004364eBcC2D37e5F6`, // Spell - used by yBribe UI
			`0xCdF7028ceAB81fA0C6971208e83fa7872994beE5`, // TNT - used by yBribe UI
			`0xba100000625a3754423978a60c9317c58a424e3D`, // BAL - used by yBAL UI
			`0x5c6Ee304399DBdB9C8Ef030aB642B10820DB8F56`, // BALWETH - used by yBAL UI
			`0x30D20208d987713f46DFD34EF128Bb16C404D10f`, // Stader
			`0xa2E3356610840701BDf5611a53974510Ae27E2e1`, // wBETH
			`0xac3E018457B222d93114458476f3E3416Abbe38F`, // Staked Frax Ether
			`0xf951E335afb289353dc249e82926178EaC7DEd78`, // Swell Network Ether
			`0x7f39C581F595B53c5cb19bD0b3f8dA6c935E2Ca0`, // Wrapped liquid staked Ether 2.0
			`0xA35b1B31Ce002FBF2058D22F30f95D405200A15b`, // Stader ETHx
			`0xBe9895146f7AF43049ca1c1AE358B0541Ea49704`, // Coinbase Wrapped Staked ETH
		},
		10:   {},
		250:  {},
		8453: {},
		42161: {
			`0x82e3A8F066a6989666b031d916c43672085b1582`, // YFI
			`0x11cDb42B0EB46D95f990BeDD4695A6e3fA034978`, // CRV
		},
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

			`0xC4C319E2D4d66CcA4464C0c2B32c9Bd23ebe784e`, // Hacked alETH
			`0x718AbE90777F5B778B52D553a5aBaa148DD0dc5D`, // Hacked vault
		},
		10:   {},
		250:  {},
		8453: {},
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
