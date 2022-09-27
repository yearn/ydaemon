package allocation

import (
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/montanaflynn/stats"
	"github.com/yearn/ydaemon/internal/prices"
	"github.com/yearn/ydaemon/internal/strategies"
	"github.com/yearn/ydaemon/internal/tokens"
	"github.com/yearn/ydaemon/internal/utils/logs"
	"github.com/yearn/ydaemon/internal/utils/models"
)

func median_to_tvl(group strategies.TStrategyGroupFromRisk) *big.Float {
	scores := []float64{
		group.AuditScore,
		group.CodeReviewScore,
		group.ComplexityScore,
		group.ProtocolSafetyScore,
		group.TeamKnowledgeScore,
		group.TestingScore,
	}

	median, _ := stats.Median(scores)
	if median > 4 {
		return big.NewFloat(0)
	} else if median > 3 {
		return big.NewFloat(1_000_000)
	} else if median > 2 {
		return big.NewFloat(10_000_000)
	} else if median > 1 {
		return big.NewFloat(50_000_000)
	} else {
		return big.NewFloat(100_000_000)
	}

}

func calculateTvlUsd(chainID uint64, strat models.TStrategyList) *big.Float {
	total_assets := big.NewFloat(0)
	token := tokens.Store.VaultToToken[chainID][strat.Vault]
	estimatedAssets := new(big.Float).SetInt(strategies.Store.StrategyMultiCallData[chainID][strat.Strategy].EstimatedTotalAssets)
	num_tokens := new(big.Float).Quo(estimatedAssets, big.NewFloat(math.Pow(10, float64(token.Decimals))))

	if prices.Store.TokenPrices[chainID][token.Address] != nil {
		token_address := tokens.Store.VaultToToken[chainID][strat.Vault].Address
		token_price := big.NewFloat(tokens.Store.Tokens[chainID][token_address].Price) // humanized price
		total_assets.Mul(num_tokens, token_price)
	}
	return total_assets
}

func FetchAllocations(chainID uint64) {
	strats, ok := strategies.Store.StrategyList[chainID]
	if !ok {
		logs.Warning("Error reading strategyList information")
		return
	}
	groups := strategies.Store.StrategyGroupFromRisk[chainID]
	strategyGroupEstimateTotalAssets := make(map[string]*big.Float)
	if Store.StrategyGroupAllocation[chainID] == nil {
		Store.StrategyGroupAllocation[chainID] = make(map[common.Address]*TStrategyAllocation)
	}
	allocations := Store.StrategyGroupAllocation[chainID]
	for _, strat := range strats {
		strategyGroups := strategies.GetStrategyGroupsFromStrategy(strat, groups)
		for _, group := range strategyGroups {
			if strategyGroupEstimateTotalAssets[group.Label] == nil {
				strategyGroupEstimateTotalAssets[group.Label] = big.NewFloat(0)
			}
			strategyGroupEstimateTotalAssets[group.Label].Add(strategyGroupEstimateTotalAssets[group.Label], calculateTvlUsd(chainID, strat))
		}

	}
	for _, strat := range strats {
		strategyGroups := strategies.GetStrategyGroupsFromStrategy(strat, groups)
		for _, group := range strategyGroups {
			max_tvl := median_to_tvl(*group)
			max_tvl.Sub(max_tvl, strategyGroupEstimateTotalAssets[group.Label])
			available_tvl := big.NewFloat(0)
			if max_tvl.Cmp(big.NewFloat(0)) > 0 {
				available_tvl = max_tvl
			}
			token_address := tokens.Store.VaultToToken[chainID][strat.Vault].Address
			token_price := big.NewFloat(tokens.Store.Tokens[chainID][token_address].Price) // humanized price
			if token_price.Cmp(big.NewFloat(0)) == 0 {
				logs.Warning(fmt.Sprintf("Failed to fetch the USDC price for token {%s}", token_address))
				continue
			}

			strategy_tvl := calculateTvlUsd(chainID, strat)
			allocation, exist := allocations[strat.Strategy]
			strategyAllocation := &TStrategyAllocation{
				&strat, group, new(big.Float).Quo(strategy_tvl, token_price), new(big.Float).Quo(available_tvl, token_price), strategy_tvl, available_tvl,
			}
			if exist {
				if available_tvl.Cmp(allocation.AvailableTVL) == -1 {
					allocations[strat.Strategy] = strategyAllocation
				}
			} else {
				allocations[strat.Strategy] = strategyAllocation
			}

		}
	}

}
