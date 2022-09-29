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
		// Fetch activation and tvl from multicall
		data, ok := Store.StrategyMultiCallData[chainID][strat.Strategy]
		if !ok {
			logs.Error("Error fetching strategy data from multicall")
			return
		}

		// Fetch tvl of strategy
		tokenAddress, ok := tokens.Store.VaultToToken[chainID][strat.Vault]
		if !ok {
			logs.Error("Error fetching token address from vault")
			continue
		}
		tokenData, ok := tokens.Store.Tokens[chainID][tokenAddress]
		if !ok {
			logs.Error("Error fetching token data for token")
			continue
		}
		_, amount := helpers.FormatAmount(data.EstimatedTotalAssets.String(), int(tokenData.Decimals))
		tvl := big.NewFloat(0).Mul(big.NewFloat(0).Set(amount), big.NewFloat(tokenData.Price))

		// Fetch risk group
		stratGroup := getStrategyGroup(chainID, strat)

		// Send to Others group if no group was found
		var strategy models.TStrategyFromRisk
		if stratGroup == nil {
			strategy = models.TStrategyFromRisk{
				RiskGroup: "Others",
				RiskScores: models.TStrategyFromRiskRiskScores{
					TVLImpact:           getTVLImpact(tvl),
					AuditScore:          5,
					CodeReviewScore:     5,
					ComplexityScore:     5,
					LongevityImpact:     getLongevityImpact(data.Activation),
					ProtocolSafetyScore: 5,
					TeamKnowledgeScore:  5,
					TestingScore:        5,
				},
			}
			Store.StrategiesFromRisk[chainID][strat.Strategy] = strategy
			continue
		}

		// Update tvl of group
		if stratGroup.Allocation.CurrentTVL == nil {
			stratGroup.Allocation.CurrentTVL = big.NewFloat(0)
		}
		if stratGroup.Allocation.CurrentAmount == nil {
			stratGroup.Allocation.CurrentAmount = big.NewFloat(0)
		}
		stratGroup.Allocation.CurrentTVL.Add(stratGroup.Allocation.CurrentTVL, tvl)
		stratGroup.Allocation.CurrentAmount.Add(stratGroup.Allocation.CurrentAmount, amount)

		// Add strategy risk info
		strategy = models.TStrategyFromRisk{
			RiskGroup: stratGroup.Label,
			RiskScores: models.TStrategyFromRiskRiskScores{
				TVLImpact:           getTVLImpact(tvl),
				AuditScore:          stratGroup.AuditScore,
				CodeReviewScore:     stratGroup.CodeReviewScore,
				ComplexityScore:     stratGroup.ComplexityScore,
				LongevityImpact:     getLongevityImpact(data.Activation),
				ProtocolSafetyScore: stratGroup.ProtocolSafetyScore,
				TeamKnowledgeScore:  stratGroup.TeamKnowledgeScore,
				TestingScore:        stratGroup.TestingScore,
			},
		}
		Store.StrategiesFromRisk[chainID][strat.Strategy] = strategy
	}

	for _, strat := range strategies {
		stratGroup := getStrategyGroup(chainID, strat)
		if stratGroup == nil {
			continue
		}
		if stratGroup.Allocation.CurrentTVL == nil {
			stratGroup.Allocation.CurrentTVL = big.NewFloat(0)
		}
		if stratGroup.Allocation.CurrentAmount == nil {
			stratGroup.Allocation.CurrentAmount = big.NewFloat(0)
		}

		// Fetch median score allocation
		medianTVL := getMedianAllocation(*stratGroup)
		availableTVL := big.NewFloat(0).Sub(medianTVL, stratGroup.Allocation.CurrentTVL)
		if availableTVL.Sign() < 0 {
			availableTVL = big.NewFloat(0)
		}

		tokenAddress, ok := tokens.Store.VaultToToken[chainID][strat.Vault]
		if !ok {
			logs.Error("Error fetching token address from vault")
			continue
		}
		tokenData, ok := tokens.Store.Tokens[chainID][tokenAddress]
		if !ok {
			logs.Error("Error fetching token data for token")
			continue
		}

		var availableAmount *big.Float
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
