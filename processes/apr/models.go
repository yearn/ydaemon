package apr

import "github.com/yearn/ydaemon/common/bigNumber"

type TFees struct {
	Performance *bigNumber.Float `json:"performance"`
	Management  *bigNumber.Float `json:"management"`
}
type TCompositeData struct {
	Boost                 *bigNumber.Float `json:"boost"`
	PoolAPY               *bigNumber.Float `json:"poolAPY"`
	BoostedAPR            *bigNumber.Float `json:"boostedAPR"`
	BaseAPR               *bigNumber.Float `json:"baseAPR"`
	CvxAPR                *bigNumber.Float `json:"cvxAPR"`
	RewardsAPY            *bigNumber.Float `json:"rewardsAPY"`
	V3OracleCurrentAPR    *bigNumber.Float `json:"v3OracleCurrentAPR,omitempty"`
	V3OracleStratRatioAPR *bigNumber.Float `json:"v3OracleStratRatioAPR,omitempty"`
	KeepCRV               *bigNumber.Float `json:"keepCRV,omitempty"`
	KeepVelo              *bigNumber.Float `json:"keepVELO,omitempty"`
}
type TExtraRewards struct {
	StakingRewardsAPY *bigNumber.Float `json:"stakingRewardsAPY"`
	GammaRewardAPY    *bigNumber.Float `json:"gammaRewardAPY"`
}
type THistoricalPoints struct {
	WeekAgo   *bigNumber.Float `json:"weekAgo"`
	MonthAgo  *bigNumber.Float `json:"monthAgo"`
	Inception *bigNumber.Float `json:"inception"`
}
type TPricePerShare struct {
	Today    *bigNumber.Float `json:"today"`
	WeekAgo  *bigNumber.Float `json:"weekAgo"`
	MonthAgo *bigNumber.Float `json:"monthAgo"`
}
type TForwardAPY struct {
	Type      string           `json:"type"`
	NetAPY    *bigNumber.Float `json:"netAPY"`
	Composite TCompositeData   `json:"composite"`
}
type TVaultAPY struct {
	Type          string            `json:"type"`
	NetAPY        *bigNumber.Float  `json:"netAPY"`
	Fees          TFees             `json:"fees"`
	Points        THistoricalPoints `json:"points"`
	PricePerShare TPricePerShare    `json:"pricePerShare"`
	Extra         TExtraRewards     `json:"extra"`
	ForwardAPY    TForwardAPY       `json:"forwardAPY"`
}

type TStrategyAPY struct {
	Type      string           `json:"type"`
	DebtRatio *bigNumber.Float `json:"debtRatio"`
	NetAPY    *bigNumber.Float `json:"netAPY"`
	Composite TCompositeData   `json:"composite"`
}
