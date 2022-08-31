package models

// TToken holds the info about the underlying (or want) token of a vault
type TToken struct {
	Address     string `json:"address"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Symbol      string `json:"symbol"`
	Description string `json:"description"`
	Decimals    uint64 `json:"decimals"`
	Icon        string `json:"icon"`
}

// TTVL holds the info about the value locked in a vault
type TTVL struct {
	TotalAssets          string  `json:"total_assets"`
	TotalDelegatedAssets string  `json:"total_delegated_assets"`
	TVLDeposited         float64 `json:"tvl_deposited"`
	TVLDelegated         float64 `json:"tvl_delegated"`
	TVL                  float64 `json:"tvl"`
	Price                float64 `json:"price"`
}

// TAPYFees holds the fees information about this vault.
type TAPYFees struct {
	Performance float64 `json:"performance"`
	Withdrawal  float64 `json:"withdrawal"`
	Management  float64 `json:"management"`
	KeepCRV     float64 `json:"keep_crv"`
	CvxKeepCRV  float64 `json:"cvx_keep_crv"`
}

// TAPYPoints holds the points information about this vault.
type TAPYPoints struct {
	WeekAgo   float64 `json:"week_ago"`
	MonthAgo  float64 `json:"month_ago"`
	Inception float64 `json:"inception"`
}

// TAPYComposite holds the points information about this vault.
type TAPYComposite struct {
	Boost      float64 `json:"boost"`
	PoolAPY    float64 `json:"pool_apy"`
	BoostedAPR float64 `json:"boosted_apr"`
	BaseAPR    float64 `json:"base_apr"`
	CvxAPR     float64 `json:"cvx_apr"`
	RewardsAPR float64 `json:"rewards_apr"`
}

// TAPY contains all the information useful about the APY, APR, fees and breakdown.
type TAPY struct {
	Type      string        `json:"type"`
	GrossAPR  float64       `json:"gross_apr"`
	NetAPY    float64       `json:"net_apy"`
	Fees      TAPYFees      `json:"fees"`
	Points    TAPYPoints    `json:"points"`
	Composite TAPYComposite `json:"composite"`
}

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
	Address         string            `json:"address"`
	Name            string            `json:"name"`
	Description     string            `json:"description"`
	DelegatedAssets string            `json:"-"`
	Details         *TStrategyDetails `json:"details,omitempty"`
	Risk            *TStrategyRisk    `json:"risk,omitempty"`
}

// TMigration helps us to know if a vault is in the process of being migrated.
type TMigration struct {
	Available bool   `json:"available"`
	Address   string `json:"address"`
}

//TVault details holds some extra information about the vault.
type TVaultDetails struct {
	Management            string  `json:"management"`
	Governance            string  `json:"governance"`
	Guardian              string  `json:"guardian"`
	Rewards               string  `json:"rewards"`
	DepositLimit          string  `json:"depositLimit"`
	Comment               string  `json:"comment"`
	AvailableDepositLimit string  `json:"availableDepositLimit,omitempty"`
	Order                 float32 `json:"-"`
	PerformanceFee        uint64  `json:"performanceFee"`
	ManagementFee         uint64  `json:"managementFee"`
	DepositsDisabled      bool    `json:"depositsDisabled"`
	WithdrawalsDisabled   bool    `json:"withdrawalsDisabled"`
	AllowZapIn            bool    `json:"allowZapIn"`
	AllowZapOut           bool    `json:"allowZapOut"`
	Retired               bool    `json:"retired"`
}

// TVault is the main structure returned by the API when trying to get all the vaults for a specific network
type TVault struct {
	Address            string         `json:"address"`
	Symbol             string         `json:"symbol"`
	DisplaySymbol      string         `json:"display_symbol"`
	FormatedSymbol     string         `json:"formated_symbol"`
	Name               string         `json:"name"`
	DisplayName        string         `json:"display_name"`
	FormatedName       string         `json:"formated_name"`
	Icon               string         `json:"icon"`
	Version            string         `json:"version"`
	Type               string         `json:"type"`
	Inception          uint64         `json:"inception"`
	Decimals           uint64         `json:"decimals"`
	Updated            uint64         `json:"updated"`
	Endorsed           bool           `json:"endorsed"`
	Emergency_shutdown bool           `json:"emergency_shutdown"`
	Token              TToken         `json:"token"`
	TVL                TTVL           `json:"tvl"`
	APY                TAPY           `json:"apy"`
	Strategies         []TStrategy    `json:"strategies"`
	Migration          TMigration     `json:"migration"`
	Details            *TVaultDetails `json:"details,omitempty"`
}
