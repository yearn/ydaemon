package strategies

import (
	"strconv"

	"github.com/montanaflynn/stats"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/traces"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/prices"
	"github.com/yearn/ydaemon/internal/tokens"
)

/**********************************************************************************************
** BuildRiskScore will build a TStrategyFromRisk from a TStrategy.
**********************************************************************************************/
func BuildRiskScore(t *models.TStrategy) *models.TStrategyFromRisk {
	strategyGroup := getStrategyGroup(t.ChainID, t)
	strategyRiskScore := getDefaultRiskGroup()
	strategyRiskScore.Address = t.Address
	strategyRiskScore.ChainID = t.ChainID
	if strategyGroup == nil {
		traces.
			Capture(`warn`, `impossible to find strategyGroup for strategy `+t.Name).
			SetEntity(`strategy`).
			SetTag(`chainID`, strconv.FormatUint(t.ChainID, 10)).
			SetTag(`rpcURI`, ethereum.GetRPCURI(t.ChainID)).
			SetTag(`strategyAddress`, t.Address.Hex()).
			SetTag(`strategyName`, t.Name).
			Send()
		return &strategyRiskScore
	}

	// Fetch activation and tvl from multicall
	strategyRiskScore.RiskDetails.LongevityImpact = getLongevityImpact(t.Activation)

	// Fetch tvl of strategy
	tokenData, ok := tokens.FindUnderlyingForVault(t.ChainID, t.VaultAddress)
	if !ok {
		traces.
			Capture(`warn`, `impossible to find token for vault `+t.VaultAddress.Hex()).
			SetEntity(`strategy`).
			SetTag(`chainID`, strconv.FormatUint(t.ChainID, 10)).
			SetTag(`rpcURI`, ethereum.GetRPCURI(t.ChainID)).
			SetTag(`strategyAddress`, t.Address.Hex()).
			SetTag(`strategyName`, t.Name).
			Send()
		return &strategyRiskScore
	}

	_tokenPrice, ok := prices.FindPrice(t.ChainID, tokenData.Address)
	if !ok {
		traces.
			Capture(`warn`, `impossible to find tokenPrice for token `+tokenData.Address.Hex()+` on chain `+strconv.FormatUint(t.ChainID, 10)).
			SetEntity(`strategy`).
			SetTag(`chainID`, strconv.FormatUint(t.ChainID, 10)).
			SetTag(`rpcURI`, ethereum.GetRPCURI(t.ChainID)).
			SetTag(`strategyAddress`, t.Address.Hex()).
			SetTag(`strategyName`, t.Name).
			Send()
	}
	_, price := helpers.FormatAmount(_tokenPrice.String(), 6)
	_, amount := helpers.FormatAmount(t.EstimatedTotalAssets.String(), int(tokenData.Decimals))
	tvl := bigNumber.NewFloat(0).Mul(amount, price)

	// Updating the default schema
	strategyRiskScore.RiskGroup = strategyGroup.Label
	strategyRiskScore.RiskDetails.AuditScore = strategyGroup.AuditScore
	strategyRiskScore.RiskDetails.CodeReviewScore = strategyGroup.CodeReviewScore
	strategyRiskScore.RiskDetails.ComplexityScore = strategyGroup.ComplexityScore
	strategyRiskScore.RiskDetails.ProtocolSafetyScore = strategyGroup.ProtocolSafetyScore
	strategyRiskScore.RiskDetails.TeamKnowledgeScore = strategyGroup.TeamKnowledgeScore
	strategyRiskScore.RiskDetails.TestingScore = strategyGroup.TestingScore
	strategyRiskScore.RiskDetails.StackingRewardTVLScore = strategyGroup.StackingRewardTVLScore
	strategyRiskScore.RiskDetails.TVLImpact = getTVLImpact(tvl)
	strategyRiskScore.Allocation = strategyGroup.Allocation

	// Calculate median score for strategy
	if strategyGroup.StackingRewardTVLScore > 0 {
		scores := stats.LoadRawData([]int{
			strategyGroup.AuditScore,
			strategyGroup.CodeReviewScore,
			strategyGroup.ComplexityScore,
			strategyGroup.ProtocolSafetyScore,
			strategyGroup.TeamKnowledgeScore,
			strategyGroup.TestingScore,
			strategyGroup.StackingRewardTVLScore,
			strategyRiskScore.RiskDetails.LongevityImpact, // Add LongevityImpact
		})
		strategyRiskScore.RiskScore, _ = stats.Median(scores)
	} else {
		scores := stats.LoadRawData([]int{
			strategyGroup.AuditScore,
			strategyGroup.CodeReviewScore,
			strategyGroup.ComplexityScore,
			strategyGroup.ProtocolSafetyScore,
			strategyGroup.TeamKnowledgeScore,
			strategyGroup.TestingScore,
			strategyRiskScore.RiskDetails.LongevityImpact, // Add LongevityImpact
		})
		strategyRiskScore.RiskScore, _ = stats.Median(scores)
	}

	return &strategyRiskScore
}
