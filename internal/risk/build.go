package risk

import (
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/montanaflynn/stats"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
)

/**********************************************************************************************
** BuildRiskScore will build a TStrategyFromRisk from a TStrategy.
**********************************************************************************************/
func BuildRiskScore(t models.TStrategy) *models.TStrategyFromRisk {
	strategyGroup := retrieveStrategyGroup(t.ChainID, t)
	strategyRiskScore := getDefaultRiskGroup()
	strategyRiskScore.Address = t.Address
	strategyRiskScore.ChainID = t.ChainID
	if strategyGroup == nil {
		logs.Warning(`impossible to find strategy group for strategy `+t.Address.Hex(), `on chain `+strconv.FormatUint(t.ChainID, 10))
		return &strategyRiskScore
	}

	tvl, _, _ := computeTVL(t)

	/**********************************************************************************************
	** The following code block is responsible for assigning the strategy group's allocation and
	** label to the strategy risk score. It also assigns the group's shared scores and strategy
	** specific scores to the strategy risk score.
	**********************************************************************************************/
	strategyRiskScore.Allocation = strategyGroup.Allocation
	strategyRiskScore.RiskGroup = strategyGroup.Label
	// Group shared scores
	strategyRiskScore.RiskDetails.AuditScore = strategyGroup.AuditScore
	strategyRiskScore.RiskDetails.CodeReviewScore = strategyGroup.CodeReviewScore
	strategyRiskScore.RiskDetails.ComplexityScore = strategyGroup.ComplexityScore
	strategyRiskScore.RiskDetails.ProtocolSafetyScore = strategyGroup.ProtocolSafetyScore
	strategyRiskScore.RiskDetails.TeamKnowledgeScore = strategyGroup.TeamKnowledgeScore
	strategyRiskScore.RiskDetails.TestingScore = strategyGroup.TestingScore
	// Strategy specific scores
	strategyRiskScore.RiskDetails.TVLImpact = getTVLImpact(tvl)
	strategyRiskScore.RiskDetails.LongevityImpact = getLongevityImpact(t.TimeActivated)
	strategyRiskScore.RiskDetails.StackingRewardTVLScore = strategyGroup.StackingRewardTVLScore

	/**********************************************************************************************
	** Here we are initializing the scoreValues array with the risk details scores. If the
	** StackingRewardTVLScore is greater than 0, it is also appended to the scoreValues array.
	** The scores are then loaded into the stats package and the median is calculated. The
	** calculated median is then assigned to the RiskScore of the strategyRiskScore.
	**********************************************************************************************/
	scoreValues := []int{
		strategyRiskScore.RiskDetails.AuditScore,
		strategyRiskScore.RiskDetails.CodeReviewScore,
		strategyRiskScore.RiskDetails.ComplexityScore,
		strategyRiskScore.RiskDetails.ProtocolSafetyScore,
		strategyRiskScore.RiskDetails.TeamKnowledgeScore,
		strategyRiskScore.RiskDetails.TestingScore,
		strategyRiskScore.RiskDetails.LongevityImpact,
	}
	if strategyGroup.StackingRewardTVLScore > 0 {
		scoreValues = append(scoreValues, strategyRiskScore.RiskDetails.StackingRewardTVLScore)
	}
	scores := stats.LoadRawData(scoreValues)
	strategyRiskScore.RiskScore, _ = stats.Median(scores)

	return &strategyRiskScore
}

/**********************************************************************************************
** BuildVaultRiskScore will build a TStaking from a TVault.
**********************************************************************************************/
func BuildVaultStaking(t models.TVault) models.TStaking {
	staking := models.TStaking{
		Available: false,
	}
	stakingData := GetStakingData(t.ChainID, t.Address)
	if !addresses.Equals(stakingData.Address, common.Address{}) {
		staking = stakingData
	}
	return staking
}
