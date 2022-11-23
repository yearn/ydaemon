package strategies

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/montanaflynn/stats"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/traces"
	"github.com/yearn/ydaemon/common/types/common"
	"github.com/yearn/ydaemon/internal/prices"
	"github.com/yearn/ydaemon/internal/tokens"
)

var stratGroupErrorAlreadySent = make(map[uint64]map[string]bool)

func excludeNameLike(strat *TStrategy, group TStrategyGroupFromRisk) bool {
	if len(group.Criteria.Exclude) > 0 {
		for _, stratExclude := range group.Criteria.Exclude {
			if strings.Contains(strings.ToLower(strat.Name), strings.ToLower(stratExclude)) {
				return true
			}
		}
	}
	return false
}

func includeAddress(strat *TStrategy, group TStrategyGroupFromRisk) bool {
	return helpers.Contains(group.Criteria.Strategies, common.FromAddress(strat.Address))
}

func includeNameLike(strat *TStrategy, group TStrategyGroupFromRisk) bool {
	if len(group.Criteria.NameLike) > 0 {
		for _, nameLike := range group.Criteria.NameLike {
			if strings.Contains(strings.ToLower(strat.Name), strings.ToLower(nameLike)) {
				return true
			}
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

func getStrategyGroup(chainID uint64, strategy *TStrategy) *TStrategyGroupFromRisk {
	groups := ListStrategiesRiskGroups(chainID)
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

/**************************************************************************************************
** RetrieveAllRisksGroupsFromFiles will read all files in the /risks/network directory for a given
** chainID and store them in the _strategyRiskGroupMap global variable.
** The goal of this function is to get a list of all the group risk information for a given
** chainID.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
**************************************************************************************************/
func RetrieveAllRisksGroupsFromFiles(chainID uint64) {
	chainIDStr := strconv.FormatUint(chainID, 10)
	content, _, err := helpers.ReadAllFilesInDir(env.BASE_DATA_PATH+`/risks/networks/`+chainIDStr+`/`, `.json`)
	if err != nil {
		traces.
			Capture(`warn`, `impossible to read risks files on chain `+chainIDStr).
			SetEntity(`strategy`).
			SetExtra(`error`, err.Error()).
			SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
			Send()
		return
	}
	for _, content := range content {
		riskGroup := TStrategyGroupFromRisk{}
		if err := json.Unmarshal(content, &riskGroup); err != nil {
			traces.
				Capture(`warn`, `impossible to unmarshall risks files response body `+chainIDStr).
				SetEntity(`strategy`).
				SetExtra(`error`, err.Error()).
				SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
				SetExtra(`content`, string(content)).
				Send()
			return
		}
		riskGroup.ChainID = chainID
		setRiskGroupInMap(chainID, &riskGroup)
	}
}

func ComputeRiskGroupAllocation(chainID uint64) {
	if stratGroupErrorAlreadySent[chainID] == nil {
		stratGroupErrorAlreadySent[chainID] = map[string]bool{}
	}
	//This will ensure we are working with clean data
	groups := ListStrategiesRiskGroups(chainID)
	for _, group := range groups {
		group.Allocation = &TStrategyAllocation{}
		group.Allocation.AvailableTVL = getMedianAllocation(*group)
		setRiskGroupInMap(chainID, group)
	}

	strategies := ListStrategies(chainID)
	for _, strategy := range strategies {
		strategyGroup := getStrategyGroup(chainID, strategy)
		if strategyGroup == nil {
			if !stratGroupErrorAlreadySent[chainID][strategy.Name] {
				traces.
					Capture(`warn`, `impossible to find stratGroup for group `+strategy.Name).
					SetEntity(`strategy`).
					SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
					SetTag(`rpcURI`, ethereum.GetRPCURI(chainID)).
					SetTag(`strategyAddress`, strategy.Address.Hex()).
					SetTag(`strategyName`, strategy.Name).
					Send()
				stratGroupErrorAlreadySent[chainID][strategy.Name] = true
			}
			continue
		}

		tokenData, ok := tokens.FindUnderlyingForVault(chainID, common.FromAddress(strategy.VaultAddress))
		if !ok {
			traces.
				Capture(`warn`, `impossible to find token for vault `+strategy.VaultAddress.Hex()).
				SetEntity(`strategy`).
				SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
				SetTag(`rpcURI`, ethereum.GetRPCURI(chainID)).
				SetTag(`strategyAddress`, strategy.Address.Hex()).
				SetTag(`strategyName`, strategy.Name).
				Send()
			continue
		}

		tokenPrice, ok := prices.FindPrice(chainID, common.FromAddress(tokenData.Address))
		if !ok {
			traces.
				Capture(`warn`, `impossible to find tokenPrice for token `+tokenData.Address.Hex()).
				SetEntity(`strategy`).
				SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
				SetTag(`rpcURI`, ethereum.GetRPCURI(chainID)).
				SetTag(`strategyAddress`, strategy.Address.Hex()).
				SetTag(`strategyName`, strategy.Name).
				Send()
			tokenPrice = bigNumber.NewInt(0)
		}
		_, price := helpers.FormatAmount(tokenPrice.String(), 6)
		_, amount := helpers.FormatAmount(strategy.EstimatedTotalAssets.String(), int(tokenData.Decimals))
		tvl := bigNumber.NewFloat(0).Mul(amount, price)

		strategyGroup.Allocation.CurrentTVL = bigNumber.NewFloat(0).Add(strategyGroup.Allocation.CurrentTVL, tvl)
		strategyGroup.Allocation.CurrentAmount = bigNumber.NewFloat(0).Add(strategyGroup.Allocation.CurrentAmount, amount)
		if tokenPrice.Sign() <= 0 {
			strategyGroup.Allocation.AvailableTVL = bigNumber.NewFloat(0)
			strategyGroup.Allocation.AvailableAmount = bigNumber.NewFloat(0)
		} else {
			strategyGroup.Allocation.AvailableTVL = bigNumber.NewFloat(0).Sub(strategyGroup.Allocation.AvailableTVL, tvl)
			strategyGroup.Allocation.AvailableAmount = bigNumber.NewFloat(0).Quo(strategyGroup.Allocation.AvailableTVL, price)
		}

		setRiskGroupInMap(chainID, strategyGroup)
	}
}
