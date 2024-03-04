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
	RewardsAPR            *bigNumber.Float `json:"rewardsAPR"`
	V3OracleCurrentAPR    *bigNumber.Float `json:"v3OracleCurrentAPR,omitempty"`
	V3OracleStratRatioAPR *bigNumber.Float `json:"v3OracleStratRatioAPR,omitempty"`
	KeepCRV               *bigNumber.Float `json:"keepCRV,omitempty"`
	KeepVelo              *bigNumber.Float `json:"keepVELO,omitempty"`
}
type TExtraRewards struct {
	StakingRewardsAPR *bigNumber.Float `json:"stakingRewardsAPR"`
}
type THistoricalPoints struct {
	WeekAgo   *bigNumber.Float `json:"weekAgo"`
	MonthAgo  *bigNumber.Float `json:"monthAgo"`
	Inception *bigNumber.Float `json:"inception"`
}
type TForwardAPR struct {
	Type      string           `json:"type"`
	NetAPR    *bigNumber.Float `json:"netAPR"`
	Composite TCompositeData   `json:"composite"`
}
type TVaultAPR struct {
	Type       string            `json:"type"`
	NetAPR     *bigNumber.Float  `json:"netAPR"`
	Fees       TFees             `json:"fees"`
	Points     THistoricalPoints `json:"points"`
	Extra      TExtraRewards     `json:"extra"`
	ForwardAPR TForwardAPR       `json:"forwardAPR"`
}

type TStrategyAPR struct {
	Type      string           `json:"type"`
	DebtRatio *bigNumber.Float `json:"debtRatio"`
	NetAPR    *bigNumber.Float `json:"netAPR"`
	Composite TCompositeData   `json:"composite"`
}
