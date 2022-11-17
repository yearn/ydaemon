package strategies

import (
	"encoding/json"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/montanaflynn/stats"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/types/common"
	"github.com/yearn/ydaemon/internal/prices"
	"github.com/yearn/ydaemon/internal/tokens"
)

func excludeNameLike(strat *TAggregatedStrategy, group TStrategyGroupFromRisk) bool {
	if len(group.Criteria.Exclude) > 0 {
		for _, stratExclude := range group.Criteria.Exclude {
			// temporary fix to handle substring inclusion
			for _, nameLike := range group.Criteria.NameLike {
				if strings.Contains(strings.ToLower(nameLike), strings.ToLower(stratExclude)) && includeNameLike(strat, group) {
					return false
				}
			}
			if strings.Contains(strings.ToLower(strat.Name), strings.ToLower(stratExclude)) {
				return true
			}
		}
	}
	return false
}

func includeAddress(strat *TAggregatedStrategy, group TStrategyGroupFromRisk) bool {
	return helpers.ContainsAddress(group.Criteria.Strategies, strat.Strategy)
}

func includeNameLike(strat models.TStrategyList, group TStrategyGroupFromRisk) bool {
	for _, nameLike := range group.Criteria.NameLike {
		if strings.Contains(strings.ToLower(strat.Name), strings.ToLower(nameLike)) {
			return true
		}
	}
	return false
}

func getTVLImpact(tvlUSDC *bigNumber.Float) float64 {
	tvl, _ := tvlUSDC.Float32()
	switch {
	case tvl == 0:
		return 0
	case tvl < 1_000_000:
		return 1
	case tvl < 10_000_000:
		return 2
	case tvl < 50_000_000:
		return 3
	case tvl < 100_000_000:
		return 4
	default:
		return 5
	}
}

func getLongevityImpact(activation *bigNumber.Int) float64 {
	if activation == nil || activation.IsZero() {
		return 5
	}
	activationUnix := time.Unix(activation.Int64(), 0)
	longevity := time.Since(activationUnix)
	days := time.Hour * 24
	switch {
	case longevity > 240*days:
		return 1
	case longevity > 120*days:
		return 2
	case longevity > 30*days:
		return 3
	case longevity > 7*days:
		return 4
	default:
		return 5
	}
}

func getMedianAllocation(group TStrategyGroupFromRisk) *bigNumber.Float {
	scores := []float64{
		group.AuditScore,
		group.CodeReviewScore,
		group.ComplexityScore,
		group.ProtocolSafetyScore,
		group.TeamKnowledgeScore,
		group.TestingScore,
	}
	median, _ := stats.Median(scores)
	switch {
	case median <= 1:
		return bigNumber.NewFloat(100_000_000)
	case median <= 2:
		return bigNumber.NewFloat(50_000_000)
	case median <= 3:
		return bigNumber.NewFloat(10_000_000)
	case median <= 4:
		return bigNumber.NewFloat(1_000_000)
	default:
		return bigNumber.NewFloat(0)
	}
}

func getStrategyGroup(chainID uint64, strategy *TAggregatedStrategy) *TStrategyGroupFromRisk {
	groups := Store.StrategyGroupFromRisk[chainID]
	var stratGroup *TStrategyGroupFromRisk
	for _, group := range groups {
		if excludeNameLike(strategy, *group) {
			continue
		}
		if includeAddress(strategy, *group) || includeNameLike(strategy, *group) {
			stratGroup = group
			break
		}
	}
	return stratGroup
}

func getDefaultRiskGroup() TStrategyFromRisk {
	return TStrategyFromRisk{
		RiskGroup: "Others",
		RiskScores: TStrategyFromRiskRiskScores{
			TVLImpact:           5,
			AuditScore:          5,
			CodeReviewScore:     5,
			ComplexityScore:     5,
			LongevityImpact:     5,
			ProtocolSafetyScore: 5,
			TeamKnowledgeScore:  5,
			TestingScore:        5,
		},
		Allocation: &TStrategyAllocation{
			CurrentTVL:      bigNumber.NewFloat(0),
			AvailableTVL:    bigNumber.NewFloat(0),
			CurrentAmount:   bigNumber.NewFloat(0),
			AvailableAmount: bigNumber.NewFloat(0),
		},
	}
}

// FetchStrategiesFromRisk fetches the strategies information from the Risk Framework for a given chainID
// and store the result to the global variable StrategiesFromRisk for later use.
func FetchStrategiesFromRisk(chainID uint64) {
	// Read data from the risk framework json file
	chainIDStr := strconv.FormatUint(chainID, 10)
	content, _, err := helpers.ReadAllFilesInDir(env.BASE_DATA_PATH+`/risks/networks/`+chainIDStr+`/`, `.json`)
	if err != nil {
		logs.Warning("Error fetching information from the Risk Framework")
		return
	}
	for _, content := range content {
		group := &TStrategyGroupFromRisk{}
		if err := json.Unmarshal(content, group); err != nil {
			logs.Warning("Error unmarshalling response body from the Risk Framework")
			return
		}
		Store.StrategyGroupFromRisk[chainID] = append(Store.StrategyGroupFromRisk[chainID], group)
	}

	// Init the store if empty
	if Store.StrategiesFromRisk[chainID] == nil {
		Store.StrategiesFromRisk[chainID] = make(map[common.Address]TStrategyFromRisk)
	}

	strategies, ok := Store.AggregatedStrategies[chainID]
	if !ok {
		logs.Warning("Error reading strategyList information")
		return
	}

	// Refresh the tvl of groups
	groups := Store.StrategyGroupFromRisk[chainID]
	for _, group := range groups {
		group.Allocation = &TStrategyAllocation{}
		group.Allocation.AvailableTVL = getMedianAllocation(*group)
	}

	for _, strat := range strategies {
		stratGroup := getStrategyGroup(chainID, strat)
		strategy := getDefaultRiskGroup()

		// Fetch activation and tvl from multicall
		data, ok := Store.AggregatedStrategies[chainID][strat.Strategy]
		if !ok {
			logs.Warning("Error fetching strategy data from multicall")
			continue
		}
		strategy.RiskScores.LongevityImpact = getLongevityImpact(data.Activation)

		// Fetch tvl of strategy
		tokenData, ok := tokens.FindUnderlyingForVault(chainID, strat.Vault)
		if !ok {
			logs.Warning("Impossible to find token for vault ", strat.Vault.String())
			Store.StrategiesFromRisk[chainID][strat.Strategy] = strategy
			continue
		}

		_tokenPrice, ok := prices.FindPrice(chainID, common.FromAddress(tokenData.Address))
		if !ok {
			logs.Warning("Impossible to find tokenPrice for token ", tokenData.Address.String())
		}
		humanizedPrice, _ := helpers.FormatAmount(_tokenPrice.String(), 6)
		tokenPrice := bigNumber.NewFloat(humanizedPrice)

		_, amount := helpers.FormatAmount(data.EstimatedTotalAssets.String(), int(tokenData.Decimals))
		tvl := bigNumber.NewFloat(0).Mul(amount, tokenPrice)
		strategy.RiskScores.TVLImpact = getTVLImpact(tvl)

		// Send to Others group if no group was found
		if stratGroup == nil {
			logs.Warning("Impossible to find stratGroup for group ", strat.Name)
			Store.StrategiesFromRisk[chainID][strat.Strategy] = strategy
			continue
		}

		// Update tvl of group
		stratGroup.Allocation.CurrentTVL = bigNumber.NewFloat(0).Add(stratGroup.Allocation.CurrentTVL, tvl)
		stratGroup.Allocation.CurrentAmount = bigNumber.NewFloat(0).Add(stratGroup.Allocation.CurrentAmount, amount)
		if tokenPrice.Sign() <= 0 {
			stratGroup.Allocation.AvailableTVL = bigNumber.NewFloat(0)
			stratGroup.Allocation.AvailableAmount = bigNumber.NewFloat(0)
		} else {
			stratGroup.Allocation.AvailableTVL = bigNumber.NewFloat(0).Sub(stratGroup.Allocation.AvailableTVL, tvl)
			stratGroup.Allocation.AvailableAmount = bigNumber.NewFloat(0).Quo(stratGroup.Allocation.AvailableTVL, tokenPrice)
		}

		// Updating the default schema
		strategy.RiskGroup = stratGroup.Label
		strategy.RiskScores.AuditScore = stratGroup.AuditScore
		strategy.RiskScores.CodeReviewScore = stratGroup.CodeReviewScore
		strategy.RiskScores.ComplexityScore = stratGroup.ComplexityScore
		strategy.RiskScores.ProtocolSafetyScore = stratGroup.ProtocolSafetyScore
		strategy.RiskScores.TeamKnowledgeScore = stratGroup.TeamKnowledgeScore
		strategy.RiskScores.TestingScore = stratGroup.TestingScore
		strategy.Allocation = stratGroup.Allocation
		Store.StrategiesFromRisk[chainID][strat.Strategy] = strategy
	}
}

// LoadRiskStrategies is kept in order to have the same behavior everywhere, but as the data
// exists in the same directory as yDaemon, saving the data in the DB is not necessary.
func LoadRiskStrategies(chainID uint64, wg *sync.WaitGroup) {
	_ = chainID
	wg.Done()
}
