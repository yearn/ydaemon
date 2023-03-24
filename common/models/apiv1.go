package models

// TAPIV1Fees holds the fees information about this vault.
type TAPIV1Fees struct {
	Performance float64 `json:"performance"`
	Withdrawal  float64 `json:"withdrawal"`
	Management  float64 `json:"management"`
	KeepCRV     float64 `json:"keep_crv"`
	CvxKeepCRV  float64 `json:"cvx_keep_crv"`
}

// TAPIV1APY contains all the information useful about the APY, APR, fees and breakdown.
type TAPIV1APY struct {
	Type              string     `json:"type"`
	GrossAPR          float64    `json:"gross_apr"`
	NetAPY            float64    `json:"net_apy"`
	StakingRewardsAPR float64    `json:"staking_rewards_apr"`
	Fees              TAPIV1Fees `json:"fees"`
	Points            struct {
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

// TAPIV1Vault is the structure of data we receive when calling api.yearn.finance/v1/chains/1/vaults/all
type TAPIV1Vault struct {
	Address string    `json:"address"`
	APY     TAPIV1APY `json:"apy"`
}
