package strategies

// TStrategyDetails contains the details about a strategy.
type TStrategyDetails struct {
	Protocols            []string `json:"protocols"`
	Version              string   `json:"version"`
	Keeper               string   `json:"keeper"`
	Strategist           string   `json:"strategist"`
	Rewards              string   `json:"rewards"`
	HealthCheck          string   `json:"healthCheck"`
	TotalDebt            string   `json:"totalDebt"`
	TotalLoss            string   `json:"totalLoss"`
	TotalGain            string   `json:"totalGain"`
	RateLimit            string   `json:"rateLimit,omitempty"`         // Only < 0.3.2
	MinDebtPerHarvest    string   `json:"minDebtPerHarvest,omitempty"` // Only >= 0.3.2
	MaxDebtPerHarvest    string   `json:"maxDebtPerHarvest,omitempty"` // Only >= 0.3.2
	EstimatedTotalAssets string   `json:"estimatedTotalAssets"`
	CreditAvailable      string   `json:"creditAvailable"`
	DebtOutstanding      string   `json:"debtOutstanding"`
	ExpectedReturn       string   `json:"expectedReturn"`
	DelegatedAssets      string   `json:"delegatedAssets"`
	DelegatedValue       string   `json:"delegatedValue"`
	APR                  float64  `json:"apr"`
	PerformanceFee       uint64   `json:"performanceFee"`
	LastReport           uint64   `json:"lastReport"`
	Activation           uint64   `json:"activation"`
	KeepCRV              uint64   `json:"keepCRV"`
	DebtRatio            uint64   `json:"debtRatio,omitempty"` // Only > 0.2.2
	DebtLimit            uint64   `json:"debtLimit"`
	DoHealthCheck        bool     `json:"doHealthCheck"`
	InQueue              bool     `json:"inQueue"`
	EmergencyExit        bool     `json:"emergencyExit"`
	IsActive             bool     `json:"isActive"`
}

// TStrategyRisk contains the details on the risk about a strategy.
type TStrategyRisk struct {
	TVLImpact           int `json:"TVLImpact"`
	AuditScore          int `json:"auditScore"`
	CodeReviewScore     int `json:"codeReviewScore"`
	ComplexityScore     int `json:"complexityScore"`
	LongevityImpact     int `json:"longevityImpact"`
	ProtocolSafetyScore int `json:"protocolSafetyScore"`
	TeamKnowledgeScore  int `json:"teamKnowledgeScore"`
	TestingScore        int `json:"testingScore"`
}

// TStrategy contains all the information useful about the strategies currently active in this vault.
type TStrategy struct {
	Address     string `json:"address"`
	Name        string `json:"name"`
	Description string `json:"description"`

	//The following fields are used for internal computation
	DelegatedAssets string `json:"-"`
	DelegatedValue  string `json:"-"`
	TotalDebt       string `json:"-"`
	InQueue         bool   `json:"-"`
	IsActive        bool   `json:"-"`
	//End of internal computation fields

	Details *TStrategyDetails `json:"details,omitempty"`
	Risk    *TStrategyRisk    `json:"risk,omitempty"`
}
