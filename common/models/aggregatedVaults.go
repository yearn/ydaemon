package models

import (
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/types/common"
)

// TLegacyAPIAPY contains all the information useful about the APY, APR, fees and breakdown.
type TLegacyAPIAPY struct {
	Type              string  `json:"type"`
	GrossAPR          float64 `json:"gross_apr"`
	NetAPY            float64 `json:"net_apy"`
	StakingRewardsAPR float64 `json:"staking_rewards_apr"`
	Fees struct {
		Performance float64 `json:"performance"`
		Withdrawal  float64 `json:"withdrawal"`
		Management  float64 `json:"management"`
		KeepCRV     float64 `json:"keep_crv"`
		CvxKeepCRV  float64 `json:"cvx_keep_crv"`
	} `json:"fees"`
	Points struct {
		WeekAgo   float64 `json:"week_ago"`
		MonthAgo  float64 `json:"month_ago"`
		Inception float64 `json:"inception"`
	} `json:"points"`
	Composite struct {
		Boost      float64 `json:"boost"`
		PoolAPY    float64 `json:"pool_apy"`
		BoostedAPR float64 `json:"boosted_apr"`
		BaseAPR    float64 `json:"base_apr"`
		CvxAPR     float64 `json:"cvx_apr"`
		RewardsAPR float64 `json:"rewards_apr"`
	} `json:"composite"`
}
type TLegacyAPI struct {
	Address common.Address
	APY     TLegacyAPIAPY
}

type TAggregatedVault struct {
	Address       common.Address
	LegacyAPY     TLegacyAPIAPY
	PricePerShare bigNumber.Int
}

func (v *TAggregatedVault) SetPricePerShare(pricePerShare *bigNumber.Int) {
	v.PricePerShare = *pricePerShare
}
