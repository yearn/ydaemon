package allocation

import (
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/montanaflynn/stats"
	"github.com/yearn/ydaemon/internal/prices"
	"github.com/yearn/ydaemon/internal/strategies"
	"github.com/yearn/ydaemon/internal/tokens"
	"github.com/yearn/ydaemon/internal/utils/logs"
)

func median_to_tvl(group strategies.TStrategyGroupFromRisk) *big.Int {
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
		return big.NewInt(0)
	} else if median > 3 {
		return big.NewInt(1_000_000)
	} else if median > 2 {
		return big.NewInt(10_000_000)
	} else if median > 1 {
		return big.NewInt(50_000_000)
	} else {
		return big.NewInt(100_000_000)
	}

}

func FetchAllocations(chainID uint64) {
	strats, ok := strategies.Store.StrategyList[chainID]
	if !ok {
		logs.Warning("Error reading strategyList information")
		return
	}
	groups := strategies.Store.StrategyGroupFromRisk[chainID]
	strategyGroupEstimateTotalAssets := make(map[string]*big.Int)
	if Store.StrategyGroupAllocation[chainID] == nil {
		Store.StrategyGroupAllocation[chainID] = make(map[common.Address]*TStrategyAllocation)
	}
	allocations := Store.StrategyGroupAllocation[chainID]
	for _, strat := range strats {
		strategyGroups := strategies.GetStrategyGroupsFromStrategy(strat, groups)
		for _, group := range strategyGroups {
			if strategyGroupEstimateTotalAssets[group.Label] == nil {
				strategyGroupEstimateTotalAssets[group.Label] = big.NewInt(0)
			}

			token := tokens.Store.VaultToToken[chainID][strat.Vault]
			total_assets := big.NewInt(0)
			total_assets.Div(strategies.Store.StrategyMultiCallData[chainID][strat.Strategy].EstimatedTotalAssets, big.NewInt(int64(math.Pow(10, float64(token.Decimals)))))
			total_assets.Mul(total_assets, prices.Store.TokenPrices[chainID][token.Address])
			strategyGroupEstimateTotalAssets[group.Label].Add(strategyGroupEstimateTotalAssets[group.Label], total_assets)
		}

	}
	for _, strat := range strats {
		strategyGroups := strategies.GetStrategyGroupsFromStrategy(strat, groups)
		for _, group := range strategyGroups {
			max_tvl := median_to_tvl(*group)
			max_tvl.Sub(max_tvl, strategyGroupEstimateTotalAssets[group.Label])
			available_tvl := big.NewInt(0)
			if max_tvl.Cmp(big.NewInt(0)) > 0 {
				available_tvl = max_tvl
			}
			token_address := tokens.Store.VaultToToken[chainID][strat.Vault].Address
			token_price := prices.Store.TokenPrices[chainID][token_address]
			strategy_tvl := big.NewInt(0)
			if token_price.Cmp(big.NewInt(0)) == 0 {
				continue
			}

			token := tokens.Store.VaultToToken[chainID][strat.Vault]
			strategy_tvl.Div(strategies.Store.StrategyMultiCallData[chainID][strat.Strategy].EstimatedTotalAssets, big.NewInt(int64(math.Pow(10, float64(token.Decimals)))))
			strategy_tvl.Mul(strategy_tvl, prices.Store.TokenPrices[chainID][token.Address])
			allocation, exist := allocations[strat.Strategy]

			strategyAllocation := &TStrategyAllocation{
				&strat, group, big.NewInt(0).Div(strategy_tvl, token_price), big.NewInt(0).Div(available_tvl, token_price), strategy_tvl, available_tvl,
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
