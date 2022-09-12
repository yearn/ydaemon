package strategies

// TStrategyGroupCriteria is a helper type for parsing the risk framework json
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
