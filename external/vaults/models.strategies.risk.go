package vaults

import "github.com/yearn/ydaemon/internal/models"

type TExternalStrategyRiskDetails struct {
	TVLImpact           int `json:"TVLImpact"`
	AuditScore          int `json:"auditScore"`
	CodeReviewScore     int `json:"codeReviewScore"`
	ComplexityScore     int `json:"complexityScore"`
	LongevityImpact     int `json:"longevityImpact"`
	ProtocolSafetyScore int `json:"protocolSafetyScore"`
	TeamKnowledgeScore  int `json:"teamKnowledgeScore"`
	TestingScore        int `json:"testingScore"`
}

type TExternalStrategyRiskScore struct {
	RiskScore   float64                      `json:"riskScore"`
	RiskGroup   string                       `json:"riskGroup"`
	RiskDetails TExternalStrategyRiskDetails `json:"riskDetails"`
	Allocation  *models.TStrategyAllocation  `json:"allocation"`
}

func NewRiskScore() *TExternalStrategyRiskScore {
	return &TExternalStrategyRiskScore{}
}
func (v *TExternalStrategyRiskScore) AssignTStrategyFromRisk(internalStrategyRiskScore *models.TStrategyFromRisk) *TExternalStrategyRiskScore {
	v.RiskScore = internalStrategyRiskScore.RiskScore
	v.RiskGroup = internalStrategyRiskScore.RiskGroup
	v.RiskDetails.TVLImpact = internalStrategyRiskScore.RiskDetails.TVLImpact
	v.RiskDetails.AuditScore = internalStrategyRiskScore.RiskDetails.AuditScore
	v.RiskDetails.CodeReviewScore = internalStrategyRiskScore.RiskDetails.CodeReviewScore
	v.RiskDetails.ComplexityScore = internalStrategyRiskScore.RiskDetails.ComplexityScore
	v.RiskDetails.LongevityImpact = internalStrategyRiskScore.RiskDetails.LongevityImpact
	v.RiskDetails.ProtocolSafetyScore = internalStrategyRiskScore.RiskDetails.ProtocolSafetyScore
	v.RiskDetails.TeamKnowledgeScore = internalStrategyRiskScore.RiskDetails.TeamKnowledgeScore
	v.RiskDetails.TestingScore = internalStrategyRiskScore.RiskDetails.TestingScore
	v.Allocation = internalStrategyRiskScore.Allocation
	return v
}
