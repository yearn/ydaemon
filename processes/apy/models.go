package apy

import "github.com/yearn/ydaemon/common/bigNumber"

type TAPIV1Fees struct {
	Performance *bigNumber.Float `json:"performance"`
	Withdrawal  *bigNumber.Float `json:"withdrawal"`
	Management  *bigNumber.Float `json:"management"`
	KeepCRV     *bigNumber.Float `json:"keep_crv"`
	KeepVelo    *bigNumber.Float `json:"keep_velo"`
	CvxKeepCRV  *bigNumber.Float `json:"cvx_keep_crv"`
}
type TAPIV1Composite struct {
	Boost      *bigNumber.Float `json:"boost"`
	PoolAPY    *bigNumber.Float `json:"pool_apy"`
	BoostedAPR *bigNumber.Float `json:"boosted_apr"`
	BaseAPR    *bigNumber.Float `json:"base_apr"`
	CvxAPR     *bigNumber.Float `json:"cvx_apr"`
	RewardsAPR *bigNumber.Float `json:"rewards_apr"`
}
type TAPIV1Points struct {
	WeekAgo   *bigNumber.Float `json:"week_ago"`
	MonthAgo  *bigNumber.Float `json:"month_ago"`
	Inception *bigNumber.Float `json:"inception"`
}
type TForwardAPR struct {
	Type      string           `json:"type"`
	GrossAPR  *bigNumber.Float `json:"gross_apr"`
	NetAPY    *bigNumber.Float `json:"net_apy"`
	Composite TAPIV1Composite  `json:"composite"`
}
type TAPIV1APY struct {
	Type              string           `json:"type"`
	GrossAPR          *bigNumber.Float `json:"gross_apr"`
	NetAPY            *bigNumber.Float `json:"net_apy"`
	StakingRewardsAPR *bigNumber.Float `json:"staking_rewards_apr"`
	Fees              TAPIV1Fees       `json:"fees"`
	Points            TAPIV1Points     `json:"points"`
	Composite         TAPIV1Composite  `json:"composite"`
	ForwardAPR        TForwardAPR      `json:"forward_apr"`
}

type TStrategyAPR struct {
	Type      string           `json:"type"`
	DebtRatio *bigNumber.Float `json:"debt_ratio"`
	GrossAPR  *bigNumber.Float `json:"gross_apr"`
	NetAPY    *bigNumber.Float `json:"net_apy"`
	Composite TAPIV1Composite  `json:"composite"`
}
