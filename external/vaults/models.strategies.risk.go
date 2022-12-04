package vaults

import (
	"github.com/yearn/ydaemon/internal/strategies"
)

type TExternalStrategyRiskScore struct {
	RiskGroup           string                          `json:"riskGroup"`
	TVLImpact           int                             `json:"TVLImpact"`
	AuditScore          int                             `json:"auditScore"`
	CodeReviewScore     int                             `json:"codeReviewScore"`
	ComplexityScore     int                             `json:"complexityScore"`
	LongevityImpact     int                             `json:"longevityImpact"`
	ProtocolSafetyScore int                             `json:"protocolSafetyScore"`
	TeamKnowledgeScore  int                             `json:"teamKnowledgeScore"`
	TestingScore        int                             `json:"testingScore"`
	Allocation          *strategies.TStrategyAllocation `json:"allocation"`
}

func NewRiskScore() *TExternalStrategyRiskScore {
	return &TExternalStrategyRiskScore{}
}
func (v *TExternalStrategyRiskScore) AssignTStrategyFromRisk(internalStrategyRiskScore *strategies.TStrategyFromRisk) *TExternalStrategyRiskScore {
	v.RiskGroup = internalStrategyRiskScore.RiskGroup
	v.TVLImpact = int(internalStrategyRiskScore.RiskScores.TVLImpact)
	v.AuditScore = int(internalStrategyRiskScore.RiskScores.AuditScore)
	v.CodeReviewScore = int(internalStrategyRiskScore.RiskScores.CodeReviewScore)
	v.ComplexityScore = int(internalStrategyRiskScore.RiskScores.ComplexityScore)
	v.LongevityImpact = int(internalStrategyRiskScore.RiskScores.LongevityImpact)
	v.ProtocolSafetyScore = int(internalStrategyRiskScore.RiskScores.ProtocolSafetyScore)
	v.TeamKnowledgeScore = int(internalStrategyRiskScore.RiskScores.TeamKnowledgeScore)
	v.TestingScore = int(internalStrategyRiskScore.RiskScores.TestingScore)
	v.Allocation = internalStrategyRiskScore.Allocation
	return v
}
