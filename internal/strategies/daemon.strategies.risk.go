package strategies

import (
	"encoding/json"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/montanaflynn/stats"
	"github.com/yearn/ydaemon/internal/tokens"
	"github.com/yearn/ydaemon/internal/utils/helpers"
	"github.com/yearn/ydaemon/internal/utils/logs"
	"github.com/yearn/ydaemon/internal/utils/models"
)

func excludeNameLike(strat models.TStrategyList, group TStrategyGroupFromRisk) bool {
	if len(group.Criteria.Exclude) > 0 {
		for _, stratExclude := range group.Criteria.Exclude {
			if strings.Contains(strings.ToLower(strat.Name), strings.ToLower(stratExclude)) {
				return true
			}
		}
	}
	return false
}

func includeAddress(strat models.TStrategyList, group TStrategyGroupFromRisk) bool {
	return helpers.ContainsAddress(group.Criteria.Strategies, strat.Strategy.String())
}

func includeNameLike(strat models.TStrategyList, group TStrategyGroupFromRisk) bool {
	if len(group.Criteria.NameLike) > 0 {
		for _, nameLike := range group.Criteria.NameLike {
			if strings.Contains(strings.ToLower(strat.Name), strings.ToLower(nameLike)) {
				return true
			}
		}
	}
	return false
}

func getTVLImpact(tvlUSDC *big.Float) float64 {
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

func getLongevityImpact(activation *big.Int) float64 {
	if activation == nil || activation == big.NewInt(0) {
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

func getMedianAllocation(group TStrategyGroupFromRisk) *big.Float {
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
		return big.NewFloat(100_000_000)
	case median <= 2:
		return big.NewFloat(50_000_000)
	case median <= 3:
		return big.NewFloat(10_000_000)
	case median <= 4:
		return big.NewFloat(1_000_000)
	default:
		return big.NewFloat(0)
	}
}

func getStrategyGroup(chainID uint64, strategy models.TStrategyList) *TStrategyGroupFromRisk {
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

func getDefaultRiskGroup() models.TStrategyFromRisk {
	return models.TStrategyFromRisk{
		RiskGroup: "Others",
		RiskScores: models.TStrategyFromRiskRiskScores{
			TVLImpact:           5,
			AuditScore:          5,
			CodeReviewScore:     5,
			ComplexityScore:     5,
			LongevityImpact:     5,
			ProtocolSafetyScore: 5,
			TeamKnowledgeScore:  5,
			TestingScore:        5,
		},
	}
}

// FetchStrategiesFromRisk fetches the strategies information from the Risk Framework for a given chainID
// and store the result to the global variable StrategiesFromRisk for later use.
func FetchStrategiesFromRisk(chainID uint64) {
	// Read data from the risk framework json file
	chainIDStr := strconv.FormatUint(chainID, 10)
	content, _, err := helpers.ReadAllFilesInDir(helpers.BASE_DATA_PATH+`/risks/networks/`+chainIDStr+`/`, `.json`)
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
		Store.StrategiesFromRisk[chainID] = make(map[common.Address]models.TStrategyFromRisk)
	}

	strategies, ok := Store.StrategyList[chainID]
	if !ok {
		logs.Warning("Error reading strategyList information")
		return
	}

	// Refresh the tvl of groups
	groups := Store.StrategyGroupFromRisk[chainID]
	for _, group := range groups {
		group.Allocation = TStrategyGroupAllocation{}
	}

	for _, strat := range strategies {
		stratGroup := getStrategyGroup(chainID, strat)
		strategy := getDefaultRiskGroup()

		// Fetch activation and tvl from multicall
		data, ok := Store.StrategyMultiCallData[chainID][strat.Strategy]
		if !ok {
			logs.Warning("Error fetching strategy data from multicall")
			continue
		}
		strategy.RiskScores.LongevityImpact = getLongevityImpact(data.Activation)

		// Fetch tvl of strategy
		var tokenData *tokens.TERC20Token
		if tokenAddress, ok := tokens.Store.VaultToToken[chainID][strat.Vault]; ok {
			if tokenData, ok = tokens.Store.Tokens[chainID][tokenAddress]; !ok {
				logs.Warning("Impossible to find tokenData for token ", tokenAddress.String())
				Store.StrategiesFromRisk[chainID][strat.Strategy] = strategy
				continue
			}
		} else {
			logs.Warning("Impossible to find token for vault ", strat.Vault.String())
			Store.StrategiesFromRisk[chainID][strat.Strategy] = strategy
			continue
		}

		_, amount := helpers.FormatAmount(data.EstimatedTotalAssets.String(), int(tokenData.Decimals))
		tvl := big.NewFloat(0).Mul(big.NewFloat(0).Set(amount), big.NewFloat(tokenData.Price))
		strategy.RiskScores.TVLImpact = getTVLImpact(tvl)

		// Send to Others group if no group was found
		if stratGroup == nil {
			logs.Warning("Impossible to find stratGroup for group ", strat.Name)
			Store.StrategiesFromRisk[chainID][strat.Strategy] = strategy
			continue
		}

		// Update tvl of group
		stratGroup.Allocation.CurrentTVL = big.NewFloat(0).Add(helpers.SafeBigFloat(stratGroup.Allocation.CurrentTVL), tvl)
		stratGroup.Allocation.CurrentAmount = big.NewFloat(0).Add(helpers.SafeBigFloat(stratGroup.Allocation.CurrentAmount), amount)

		// Updating the default schema
		strategy.RiskGroup = stratGroup.Label
		strategy.RiskScores.AuditScore = stratGroup.AuditScore
		strategy.RiskScores.CodeReviewScore = stratGroup.CodeReviewScore
		strategy.RiskScores.ComplexityScore = stratGroup.ComplexityScore
		strategy.RiskScores.ProtocolSafetyScore = stratGroup.ProtocolSafetyScore
		strategy.RiskScores.TeamKnowledgeScore = stratGroup.TeamKnowledgeScore
		strategy.RiskScores.TestingScore = stratGroup.TestingScore
		Store.StrategiesFromRisk[chainID][strat.Strategy] = strategy
	}

	for _, strat := range strategies {
		stratGroup := getStrategyGroup(chainID, strat)
		if stratGroup == nil {
			continue
		}

		stratGroup.Allocation.CurrentTVL = helpers.SafeBigFloat(stratGroup.Allocation.CurrentTVL)
		stratGroup.Allocation.CurrentAmount = helpers.SafeBigFloat(stratGroup.Allocation.CurrentAmount)

		// Fetch median score allocation
		medianTVL := getMedianAllocation(*stratGroup)
		availableTVL := big.NewFloat(0).Sub(medianTVL, stratGroup.Allocation.CurrentTVL)
		if availableTVL.Sign() < 0 {
			availableTVL = big.NewFloat(0)
		}

		var tokenData *tokens.TERC20Token
		if tokenAddress, ok := tokens.Store.VaultToToken[chainID][strat.Vault]; ok {
			if tokenData, ok = tokens.Store.Tokens[chainID][tokenAddress]; !ok {
				logs.Warning("Impossible to find tokenData for token ", tokenAddress.String())
				continue
			}
		} else {
			logs.Warning("Impossible to find token for vault ", strat.Vault.String())
			continue
		}

		availableAmount := big.NewFloat(0)
		if tokenData.Price > 0 {
			availableAmount = big.NewFloat(0).Quo(availableTVL, big.NewFloat(tokenData.Price))
		}

		// Assign values from risk group
		if availableAmount != nil {
			stratRisk := Store.StrategiesFromRisk[chainID][strat.Strategy]
			stratRisk.Allocation.CurrentTVL = stratGroup.Allocation.CurrentTVL.String()
			stratRisk.Allocation.AvailableTVL = availableTVL.String()
			stratRisk.Allocation.CurrentAmount = stratGroup.Allocation.CurrentAmount.String()
			stratRisk.Allocation.AvailableAmount = availableAmount.String()
			Store.StrategiesFromRisk[chainID][strat.Strategy] = stratRisk
		}
	}
}

// LoadRiskStrategies is kept in order to have the same behavior everywhere, but as the data
// exists in the same directory as yDaemon, saving the data in the DB is not necessary.
func LoadRiskStrategies(chainID uint64, wg *sync.WaitGroup) {
	_ = chainID
	wg.Done()
}
