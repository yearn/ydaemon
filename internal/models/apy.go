package models

import "github.com/yearn/ydaemon/common/bigNumber"

// TNewAPY contains the basic information of an ERC20 token
type TNewAPY struct {
	PricePerShare *bigNumber.Float `json:"pricePerShare"`
	WeekAPR       *bigNumber.Float `json:"weekAPR"`
	MonthAPR      *bigNumber.Float `json:"monthAPR"`
	YearAPR       *bigNumber.Float `json:"yearAPR"`
	InceptionAPR  *bigNumber.Float `json:"inceptionAPR"`
	Points        struct {
		WeekAgo  *bigNumber.Float `json:"weekAgo"`
		MonthAgo *bigNumber.Float `json:"monthAgo"`
		YearAgo  *bigNumber.Float `json:"yearAgo"`
	}
	Fees struct {
		PerformanceFee uint64 `json:"performanceFee"`
		ManagementFee  uint64 `json:"managementFee"`
	}
}
