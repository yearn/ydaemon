package vaults

import (
	"math"

	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/types/common"
	"github.com/yearn/ydaemon/external/prices"
)

// Get the price of the underlying asset. This is tricky because of the decimals. The prices are fetched
// using the lens oracle daemon, stored in the TokenPrices map, and based on the USDC token, aka with 6
// decimals. We first need to parse the Int Price to a float64, then divide by 10^6 to get the price
// in an human readable USDC format.
func buildTokenPrice(chainID uint64, tokenAddress common.Address) (*bigNumber.Float, float64) {
	prices := prices.Store.TokenPrices[chainID]
	fPrice := new(bigNumber.Float)
	price, ok := prices[tokenAddress]
	if ok {
		fPrice.SetInt(price)
		humanizedPrice := new(bigNumber.Float).Quo(fPrice, bigNumber.NewFloat(math.Pow10(int(6))))
		fHumanizedPrice, _ := humanizedPrice.Float64()
		return humanizedPrice, fHumanizedPrice
	}
	return bigNumber.NewFloat(), 0.0
}

// Get the total assets locked in this vault. This is tricky because of the decimals. The total asset value
// is a string which will be formated as a float64 and scaled with the underlying token decimals. With that
// we should have a human readable Total Asset value, and we should be able to get the Total Value Locked
// in the vault thanks to the above humanizedPrice value.
func buildTVL(balanceToken *bigNumber.Int, decimals int, humanizedPrice *bigNumber.Float) float64 {
	_, humanizedTVL := helpers.FormatAmount(balanceToken.String(), decimals)
	fHumanizedTVLPrice, _ := bigNumber.NewFloat().Mul(humanizedTVL, humanizedPrice).Float64()
	return fHumanizedTVLPrice
}

// func prepareVaultSchema(
// 	chainID uint64,
// 	strategiesCondition string,
// 	withStrategiesRisk bool,
// 	withStrategiesDetails bool,
// 	currentVault *vaults.TVault,
// ) *vaults.TVault {
// 	chainIDAsString := strconv.FormatUint(chainID, 10)
// 	vaultAddress := common.FromAddress(currentVault.Address)
// 	tokenAddress := common.FromAddress(currentVault.Token.Address)
// 	tokenFromMeta := meta.Store.TokensFromMeta[chainID][tokenAddress]
// 	updated := 0    //helpers.FormatUint64(vaultFromGraph.LatestUpdate.Timestamp, 0)
// 	activation := 0 //helpers.FormatUint64(vaultFromGraph.Activation, 0)
// 	vaultFromMeta, ok := meta.Store.VaultsFromMeta[chainID][vaultAddress]
// 	if !ok {
// 		// If the vault file is missing, we set the default values for its fields
// 		vaultFromMeta = meta.TVaultFromMeta{
// 			Order:               1000000000,
// 			HideAlways:          false,
// 			DepositsDisabled:    false,
// 			WithdrawalsDisabled: false,
// 			MigrationAvailable:  false,
// 			AllowZapIn:          true,
// 			AllowZapOut:         true,
// 			Retired:             false,
// 		}
// 	}

// 	humanizedPrice, fHumanizedPrice := buildTokenPrice(
// 		chainID,
// 		tokenAddress,
// 	)

// 	strategies := strategies.BuildStrategies(
// 		chainID,
// 		withStrategiesDetails,
// 		withStrategiesRisk,
// 		strategiesCondition,
// 		humanizedPrice,
// 		vaultFromGraph,
// 	)

// 	fHumanizedTVLPrice := buildTVL(
// 		vaultFromGraph.BalanceTokens,
// 		int(vaultFromGraph.Token.Decimals),
// 		humanizedPrice,
// 	)
// 	delegatedTokenAsBN := bigNumber.NewInt(0)
// 	fDelegatedValue := 0.0

// 	for _, strat := range strategies {
// 		stratDelegatedValueAsFloat, err := strconv.ParseFloat(strat.DelegatedValue, 64)
// 		if err == nil {
// 			delegatedTokenAsBN = delegatedTokenAsBN.Add(delegatedTokenAsBN, strat.DelegatedAssets)
// 			fDelegatedValue += stratDelegatedValueAsFloat
// 		}
// 	}

// 	return currentVault

// vault := &TVault{
// 	Inception:      activation,
// 	TVL: TTVL{
// 		TotalAssets:          vaultFromGraph.BalanceTokens,
// 		TotalDelegatedAssets: delegatedTokenAsBN,
// 		TVL:                  fHumanizedTVLPrice - fDelegatedValue,
// 		TVLDeposited:         fHumanizedTVLPrice,
// 		TVLDelegated:         fDelegatedValue,
// 		Price:                fHumanizedPrice,
// 	},
// 	Strategies: strategies,
// }

// return vault
// }
