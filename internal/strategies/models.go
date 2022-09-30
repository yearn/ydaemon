package strategies

import "github.com/ethereum/go-ethereum/common"

// TStrategyDetails contains the details about a strategy.
type TStrategyDetails struct {
	Keeper                  common.Address `json:"keeper"`
	Strategist              common.Address `json:"strategist"`
	Rewards                 common.Address `json:"rewards"`
	HealthCheck             common.Address `json:"healthCheck"`
	Protocols               []string       `json:"protocols"`
	Version                 string         `json:"version"`
	TotalDebt               string         `json:"totalDebt"`
	TotalLoss               string         `json:"totalLoss"`
	TotalGain               string         `json:"totalGain"`
	RateLimit               string         `json:"rateLimit,omitempty"`         // Only < 0.3.2
	MinDebtPerHarvest       string         `json:"minDebtPerHarvest,omitempty"` // Only >= 0.3.2
	MaxDebtPerHarvest       string         `json:"maxDebtPerHarvest,omitempty"` // Only >= 0.3.2
	EstimatedTotalAssets    string         `json:"estimatedTotalAssets"`
	CreditAvailable         string         `json:"creditAvailable"`
	DebtOutstanding         string         `json:"debtOutstanding"`
	ExpectedReturn          string         `json:"expectedReturn"`
	DelegatedAssets         string         `json:"delegatedAssets"`
	DelegatedValue          string         `json:"delegatedValue"`
	APR                     float64        `json:"apr"`
	PerformanceFee          uint64         `json:"performanceFee"`
	LastReport              uint64         `json:"lastReport"`
	Activation              uint64         `json:"activation"`
	KeepCRV                 uint64         `json:"keepCRV"`
	DebtRatio               uint64         `json:"debtRatio,omitempty"` // Only > 0.2.2
	DebtLimit               uint64         `json:"debtLimit"`
	WithdrawalQueuePosition int64          `json:"withdrawalQueuePosition"`
	DoHealthCheck           bool           `json:"doHealthCheck"`
	InQueue                 bool           `json:"inQueue"`
	EmergencyExit           bool           `json:"emergencyExit"`
	IsActive                bool           `json:"isActive"`
}

// TStrategy contains all the information useful about the strategies currently active in this vault.
type TStrategy struct {
	Address     common.Address `json:"address"`
	Name        string         `json:"name"`
	Description string         `json:"description"`

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
