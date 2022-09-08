package vaults

import (
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/strategies"
	"github.com/yearn/ydaemon/internal/utils/helpers"
	"github.com/yearn/ydaemon/internal/utils/models"
)

func buildDelegated(delegatedBalanceToken string, decimals int, humanizedPrice *big.Float) float64 {
	_, delegatedTVL := helpers.FormatAmount(delegatedBalanceToken, decimals)
	fHumanizedTVLPrice, _ := big.NewFloat(0).Mul(delegatedTVL, humanizedPrice).Float64()
	return fHumanizedTVLPrice
}

func prepareTVL(
	chainID uint64,
	vaultFromGraph models.TVaultFromGraph,
) float64 {
	tokenAddress := common.HexToAddress(vaultFromGraph.Token.Id)
	delegatedTokenAsBN := big.NewInt(0)
	fDelegatedValue := 0.0

	humanizedPrice, _ := buildTokenPrice(chainID, tokenAddress)
	fHumanizedTVLPrice := buildTVL(
		vaultFromGraph.BalanceTokens,
		int(vaultFromGraph.Token.Decimals),
		humanizedPrice,
	)

	strategiesFromMulticall := strategies.Store.StrategyMultiCallData[chainID]
	for _, strat := range vaultFromGraph.Strategies {
		multicallData, ok := strategiesFromMulticall[common.HexToAddress(strat.Address)]
		if !ok {
			multicallData = models.TStrategyMultiCallData{}
		}

		delegatedAssets := helpers.BValueWithFallbackString(multicallData.DelegatedAssets, `0`)
		delegatedValue := strconv.FormatFloat(
			buildDelegated(
				delegatedAssets,
				int(vaultFromGraph.Token.Decimals),
				humanizedPrice,
			), 'f', -1, 64,
		)
		stratDelegatedValueAsFloat, err := strconv.ParseFloat(delegatedValue, 64)
		if err == nil {
			stratDelegatedTokenAsBN, ok := big.NewInt(0).SetString(delegatedAssets, 10)
			if ok {
				delegatedTokenAsBN = delegatedTokenAsBN.Add(delegatedTokenAsBN, stratDelegatedTokenAsBN)
				fDelegatedValue += stratDelegatedValueAsFloat
			}
		}
	}

	return fHumanizedTVLPrice - fDelegatedValue
}
