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

func getTVLImpact(tvlUSDC *bigNumber.Float) int {
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

func getLongevityImpact(activation *bigNumber.Int) int {
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

func getMedianScore(group TStrategyGroupFromRisk) int {
	scores := stats.LoadRawData([]int{
		group.AuditScore,
		group.CodeReviewScore,
		group.ComplexityScore,
		group.ProtocolSafetyScore,
		group.TeamKnowledgeScore,
		group.TestingScore,
	})
	median, _ := stats.Median(scores)
	return int(median)
}

func getMedianAllocation(group TStrategyGroupFromRisk) *bigNumber.Float {
	median := getMedianScore(group)
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

func getAllocationStatus(group TStrategyGroupFromRisk) string {
	median := getMedianScore(group)
	medianAllocation := getMedianAllocation(group)
	tvl := group.Allocation.CurrentTVL
	diff := bigNumber.NewFloat(0).Sub(medianAllocation, tvl)

	// under allocated - Green
	if diff.Sign() > 0 {
		return "Green"
	}
	diffAllowed := bigNumber.NewFloat(0)
	switch {
	case median <= 1:
		diffAllowed = bigNumber.NewFloat(100_000_000)
	case median <= 2:
		diffAllowed = bigNumber.NewFloat(50_000_000)
	case median <= 3:
		diffAllowed = bigNumber.NewFloat(40_000_000)
	case median <= 4:
		diffAllowed = bigNumber.NewFloat(9_000_000)
	case median <= 5:
		diffAllowed = bigNumber.NewFloat(1_000_000)
	}
	// over allocated but within allowed range - Yellow
	if diff.Sub(diffAllowed, diff).Sign() > 0 {
		return "Yellow"
	}
	// over allocated and out of range - Red
	return "Red"
}

func getStrategyGroup(chainID uint64, strategy *TStrategy) *TStrategyGroupFromRisk {
	toLowerName := strings.ToLower(strategy.Name)
	groups := ListStrategiesRiskGroups(chainID)
	for _, group := range groups {
		// check if nameLike and exclude intersect
		if helpers.Intersects(group.Criteria.NameLike, group.Criteria.Exclude) {
			toLowerNameLike := helpers.ToLower(group.Criteria.NameLike)
			toLowerExclude := helpers.ToLower(group.Criteria.Exclude)
			for _, nameLike := range toLowerNameLike {
				// if the nameLike is more specific
				if strings.Contains(toLowerName, nameLike) &&
					helpers.ContainsSubString(toLowerExclude, nameLike) {
					return group
				}
			}
		}
		// check address
		if helpers.Contains(group.Criteria.Strategies, strategy.Address.String()) {
			return group
		}

		// check exclude
		toLowerExclude := helpers.ToLower(group.Criteria.Exclude)
		if helpers.ContainsSubString(toLowerExclude, toLowerName) {
			continue
		}

		// check nameLike
		toLowerNameLike := helpers.ToLower(group.Criteria.NameLike)
		if helpers.ContainsSubString(toLowerNameLike, toLowerName) {
			return group
		}
	}
	return nil
}

func getDefaultRiskGroup() TStrategyFromRisk {
	return TStrategyFromRisk{
		RiskGroup: "Others",
		RiskScore: 5,
		RiskDetails: TStrategyFromRiskRiskDetails{
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
			Status:          "Green",
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
			traces.
				Capture(`warn`, `impossible to find stratGroup for group `+strategy.Name).
				SetEntity(`strategy`).
				SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
				SetTag(`rpcURI`, ethereum.GetRPCURI(chainID)).
				SetTag(`strategyAddress`, strategy.Address.Hex()).
				SetTag(`strategyName`, strategy.Name).
				Send()
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
		strategyGroup.Allocation.Status = getAllocationStatus(*strategyGroup)
		setRiskGroupInMap(chainID, strategyGroup)
	}
}
