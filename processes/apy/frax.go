package apy

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/internal/models"
)

type TCalculateFraxAPYDataStruct struct {
	vault          models.TVault
	gaugeAddress   common.Address
	strategy       *models.TStrategy
	baseAssetPrice *bigNumber.Float
	poolPrice      *bigNumber.Float
	baseAPR        *bigNumber.Float
	rewardAPR      *bigNumber.Float
	poolDailyAPY   *bigNumber.Float
}

func calculateFraxForwardAPR(args TCalculateFraxAPYDataStruct, fraxPool TFraxPool) TStrategyAPR {
	baseConvexStrategyData := calculateConvexForwardAPR(
		TCalculateConvexAPYDataStruct(args),
	)
	minRewardsAPR := bigNumber.NewFloat(0).SetString(fraxPool.TotalRewardAPRs.Min)
	apyStruct := TStrategyAPR{
		Type:      "frax",
		DebtRatio: baseConvexStrategyData.DebtRatio,
		GrossAPR:  bigNumber.NewFloat(0).Add(baseConvexStrategyData.GrossAPR, minRewardsAPR),
		NetAPY:    bigNumber.NewFloat(0).Add(baseConvexStrategyData.NetAPY, minRewardsAPR),
		Composite: TAPIV1Composite{
			Boost:      baseConvexStrategyData.Composite.Boost,
			PoolAPY:    baseConvexStrategyData.Composite.PoolAPY,
			BoostedAPR: baseConvexStrategyData.Composite.BoostedAPR,
			BaseAPR:    baseConvexStrategyData.Composite.BaseAPR,
			CvxAPR:     baseConvexStrategyData.Composite.CvxAPR,
			RewardsAPR: bigNumber.NewFloat(0).Add(baseConvexStrategyData.Composite.RewardsAPR, minRewardsAPR),
		},
	}
	return apyStruct
}
