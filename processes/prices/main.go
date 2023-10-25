package prices

import (
	"context"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/store"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

var priceErrorAlreadySent = make(map[uint64]map[common.Address]bool)

func listMissingPrices(chainID uint64, tokenMap map[common.Address]models.TERC20Token, newPriceMap map[common.Address]models.TPrices) []models.TERC20Token {
	tokenSlice := []models.TERC20Token{}
	for _, token := range tokenMap {
		if price, ok := newPriceMap[token.Address]; (ok && price.Price.IsZero()) || !ok {
			tokenSlice = append(tokenSlice, token)
		}
	}
	return tokenSlice
}

/**************************************************************************************************
** fetchPrices will, for a list of addresses, try to fetch all the prices from the lens price
** oracle. If the price is not available, it will try to fetch it from some external API.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - tokenMap: a list of addresses of the tokens we want to fetch the information for
**
** Returns:
** - a map of tokenAddress -> *bigNumber.Int
**************************************************************************************************/
func fetchPrices(
	chainID uint64,
	blockNumber *uint64,
	tokenMap map[common.Address]models.TERC20Token,
) map[common.Address]models.TPrices {
	newPriceMap := make(map[common.Address]models.TPrices)
	tokenSlice := storage.ERC20MapToSlice(tokenMap)

	/**********************************************************************************************
	** We now fill in the missing prices using the DeFiLlama and CoinGecko API.
	**********************************************************************************************/
	pricesLlama := fetchPricesFromLlama(chainID, tokenSlice)
	for _, token := range tokenMap {
		if price, ok := pricesLlama[token.Address]; ok {
			if !price.Price.IsZero() {
				newPriceMap[token.Address] = price
				continue
			}
		}
	}

	tokenSlice = listMissingPrices(chainID, tokenMap, newPriceMap)
	pricesGecko := fetchPricesFromGecko(chainID, tokenSlice)
	for _, token := range tokenMap {
		if price, ok := pricesGecko[token.Address]; ok {
			if !price.Price.IsZero() {
				newPriceMap[token.Address] = price
				continue
			}
		}
	}

	/**********************************************************************************************
	** Once this is done, we will probably have some missing tokens. We can use the Curve API to
	** be able to calculate the price of some tokens. We will then add them to our map.
	**********************************************************************************************/
	priceMapFromCurveFactoryAPI := getPricesFromCurveFactoriesAPI(chainID)
	for _, price := range priceMapFromCurveFactoryAPI {
		if !price.Price.IsZero() {
			newPriceMap[price.Address] = price
			continue
		}
	}

	/**********************************************************************************************
	** Once this is done, we will probably have some missing tokens. We can use the Velo API to
	** be able to calculate the price of some tokens. We will then add them to our map. Only on
	** optimism
	**********************************************************************************************/
	tokenSlice = listMissingPrices(chainID, tokenMap, newPriceMap)
	priceMapFromVeloOracle := fetchPricesFromSugar(chainID, blockNumber, tokenSlice)
	for _, price := range priceMapFromVeloOracle {
		if !price.Price.IsZero() {
			newPriceMap[price.Address] = price
			continue
		}
	}

	/**********************************************************************************************
	** Once this is done, we will probably have some missing tokens. We can use the Aero API to
	** be able to calculate the price of some tokens. We will then add them to our map. Only on
	** Base
	**********************************************************************************************/
	tokenSlice = listMissingPrices(chainID, tokenMap, newPriceMap)
	priceMapFromAeroOracle := fetchPricesFromAeroSugar(chainID, blockNumber, tokenSlice)
	for _, price := range priceMapFromAeroOracle {
		if !price.Price.IsZero() {
			newPriceMap[price.Address] = price
			continue
		}
	}

	/**********************************************************************************************
	** With the new version of the Curve LP Token, we can use the contract itself to get the price
	** of the LP token. We will then add them to our map.
	**********************************************************************************************/
	tokenSlice = listMissingPrices(chainID, tokenMap, newPriceMap)
	priceFromCurveAMM := fetchPricesFromCurveAMM(chainID, blockNumber, tokenSlice)
	for _, price := range priceFromCurveAMM {
		if !price.Price.IsZero() {
			newPriceMap[price.Address] = price
			continue
		}
	}

	/**********************************************************************************************
	** If we still have some missing prices, we will use the lens price oracle to fetch them.
	**********************************************************************************************/
	tokenSlice = listMissingPrices(chainID, tokenMap, newPriceMap)
	priceMapLensOracle := fetchPricesFromLens(chainID, blockNumber, tokenSlice)
	for _, price := range priceMapLensOracle {
		if price, ok := newPriceMap[price.Address]; ok && price.Price.IsZero() {
			newPriceMap[price.Address] = price
		}
	}

	/**********************************************************************************************
	** With the ERC-4626 standard, the `price per share` is no longer relevant. We can use the new
	** `convertToAssets` function to get the value of the underlying asset for a given amount of
	** shares: 1 share will give me 1.23 asset for example.
	** Based on that, if we have the price of the underlying asset, we can calculate the price of
	** the share.
	**********************************************************************************************/
	tokenSlice = listMissingPrices(chainID, tokenMap, newPriceMap)
	sharesValue := fetchShareValueFromERC4626(chainID, blockNumber, tokenSlice)
	for _, shareValue := range sharesValue {
		if shareValue.Value == nil || shareValue.Value.IsZero() {
			continue
		}
		currentPrice, ok := newPriceMap[shareValue.AssetAddress]
		if ok && !currentPrice.Price.IsZero() {
			continue
		}

		token, ok := storage.GetERC20(chainID, shareValue.AssetAddress)
		if !ok {
			continue
		}

		tokenDecimals := helpers.ToRawAmount(bigNumber.NewInt(1), token.Decimals)
		sharePrice := bigNumber.NewFloat(0).Quo(
			bigNumber.NewFloat(0).Mul(
				bigNumber.NewFloat(0).SetInt(shareValue.Value),
				bigNumber.NewFloat(0).SetInt(currentPrice.Price),
			),
			bigNumber.NewFloat(0).SetInt(tokenDecimals),
		)
		humanizedPrice := helpers.ToNormalizedAmount(sharePrice.Int(), 6)
		newPriceMap[shareValue.VaultAddress] = models.TPrices{
			Address:        shareValue.VaultAddress,
			Price:          sharePrice.Int(),
			HumanizedPrice: humanizedPrice,
			Source:         `ERC4626-convertToAssets`,
		}
	}

	/**********************************************************************************************
	** If the price is missing, check if it's a vault and try to compute the price from the
	** underlying tokens.
	**********************************************************************************************/
	logs.Error(
		newPriceMap[common.HexToAddress(`0x667002f9dc61ebcba8ee1cbeb2ad04060388f223`)],
		newPriceMap[common.HexToAddress(`0x78e20312105a85b7fe86db119cfd5cd8de81fdaa`)],
	)

	for _, token := range tokenMap {
		if addresses.Equals(token.Address, `0x78e20312105a85b7fe86db119cfd5cd8de81fdaa`) {
			logs.Pretty(`SAAME BEFORE`)
		}
		if value, ok := newPriceMap[token.Address]; !ok || value.Price.IsZero() {
			if addresses.Equals(token.Address, `0x78e20312105a85b7fe86db119cfd5cd8de81fdaa`) {
				logs.Pretty(`SAAME`)
			}
			if token.IsVaultLike() {
				ppsPerTime, _ := store.ListPricePerShare(chainID, token.Address)
				underlyingToken := token.UnderlyingTokensAddresses[0]
				ppsToday := helpers.GetToday(ppsPerTime, token.Decimals)
				if ppsToday == nil || ppsToday.IsZero() {
					ppsToday = ethereum.FetchPPSToday(chainID, token.Address, token.Decimals)
				}
				underlyingPrice := bigNumber.NewFloat(0).SetInt(newPriceMap[underlyingToken].Price)
				vaultPrice := bigNumber.NewFloat(0).Mul(ppsToday, underlyingPrice)
				humanizedPrice := helpers.ToNormalizedAmount(vaultPrice.Int(), 6)
				newPriceMap[token.Address] = models.TPrices{
					Address:        token.Address,
					Price:          vaultPrice.Int(),
					HumanizedPrice: humanizedPrice,
					Source:         `yVaultV2-pps`,
				}
			}
		}
	}
	logs.Error(newPriceMap[common.HexToAddress(`0x667002f9dc61ebcba8ee1cbeb2ad04060388f223`)])

	/**********************************************************************************************
	** If the price is still missing, we can fallback to try to retreive the price per share of the
	** vaults (vs the underlying token). Just like with the ERC-4626 standard, we can then
	** calculate the price of one share (underlyingPrice * pricePerShare)
	**********************************************************************************************/
	tokenSlice = listMissingPrices(chainID, tokenMap, newPriceMap)
	pricePerShareValue := fetchPricePerShareFromVault(chainID, blockNumber, tokenSlice)
	for _, ppsValue := range pricePerShareValue {
		if ppsValue.Value == nil || ppsValue.Value.IsZero() {
			continue
		}

		currentPrice, ok := newPriceMap[ppsValue.AssetAddress]
		if ok && !currentPrice.Price.IsZero() {
			continue
		}

		token, ok := storage.GetERC20(chainID, ppsValue.AssetAddress)
		if !ok {
			continue
		}

		tokenDecimals := helpers.ToRawAmount(bigNumber.NewInt(1), token.Decimals)
		sharePrice := bigNumber.NewFloat(0).Quo(
			bigNumber.NewFloat(0).Mul(
				bigNumber.NewFloat(0).SetInt(ppsValue.Value),
				bigNumber.NewFloat(0).SetInt(currentPrice.Price),
			),
			bigNumber.NewFloat(0).SetInt(tokenDecimals),
		)
		humanizedPrice := helpers.ToNormalizedAmount(sharePrice.Int(), 6)
		newPriceMap[ppsValue.VaultAddress] = models.TPrices{
			Address:        ppsValue.VaultAddress,
			Price:          sharePrice.Int(),
			HumanizedPrice: humanizedPrice,
			Source:         `yVaultV2-pps`,
		}
	}

	logs.Pretty(`DONE`)
	/**********************************************************************************************
	** Finally, we will list all the tokens that are still missing a price to log them to Sentry.
	**********************************************************************************************/
	if priceErrorAlreadySent[chainID] == nil {
		priceErrorAlreadySent[chainID] = make(map[common.Address]bool)
	}

	for _, token := range tokenMap {
		tokenPrice, ok := newPriceMap[token.Address]
		if !ok || tokenPrice.Price.IsZero() {
			if !priceErrorAlreadySent[chainID][token.Address] {
				logs.Warning(`missing a valid price for ` + token.Address.Hex() + ` (` + token.Name + `)` + ` on chain ` + strconv.FormatUint(chainID, 10) + ` (` + string(token.Type) + `)`)
			}
			priceErrorAlreadySent[chainID][token.Address] = true
		}
	}

	for _, price := range newPriceMap {
		storage.StorePrice(chainID, price)
	}
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
func RetrieveAllPrices(chainID uint64, tokenMap map[common.Address]models.TERC20Token) map[common.Address]models.TPrices {
	/**********************************************************************************************
	** Once we got out basic list, we will fetch the price for each of theses tokens, using lens
	** as primary source, and multiple other sources as fallback.
	**********************************************************************************************/
	currentBlockNumber, _ := ethereum.GetRPC(chainID).BlockNumber(context.Background())
	if len(tokenMap) > 0 {
		fetchPrices(chainID, &currentBlockNumber, tokenMap)
	}
	priceMap, _ := storage.ListPrices(chainID)
	return priceMap
}
