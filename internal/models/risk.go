package models

type TStrategyFromRiskRiskScores struct {
	TVLImpact           float64 `json:"TVLImpact"`
	AuditScore          float64 `json:"auditScore"`
	CodeReviewScore     float64 `json:"codeReviewScore"`
	ComplexityScore     float64 `json:"complexityScore"`
	LongevityImpact     float64 `json:"longevityImpact"`
	ProtocolSafetyScore float64 `json:"protocolSafetyScore"`
	TeamKnowledgeScore  float64 `json:"teamKnowledgeScore"`
	TestingScore        float64 `json:"testingScore"`
}

type TStrategyFromRiskRiskProfile struct {
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Median float64 `json:"median"`
}

type TStrategyFromRisk struct {
	RiskScores TStrategyFromRiskRiskScores `json:"riskScores"`
	// RiskProfile TStrategyFromRiskRiskProfile `json:"riskProfile"`
	// Tokens      []string                     `json:"tokens"`
}

type TStrategyGroupCritieria struct {
	NameLike   []string `json:"nameLike"`
	Strategies []string `json:"strategies"`
	Exclude    []string `json:"exclude"`
}

type TStrategyGroupFromRisk struct {
	GroupID             string                  `json:"id"`
	ChainID             uint64                  `json:"network"`
	AuditScore          float64                 `json:"auditScore"`
	CodeReviewScore     float64                 `json:"codeReviewScore"`
	ComplexityScore     float64                 `json:"complexityScore"`
	ProtocolSafetyScore float64                 `json:"protocolSafetyScore"`
	TeamKnowledgeScore  float64                 `json:"teamKnowledgeScore"`
	TestingScore        float64                 `json:"testingScore"`
	Criteria            TStrategyGroupCritieria `json:"criteria"`
}
