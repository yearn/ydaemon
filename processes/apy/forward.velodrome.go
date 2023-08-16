package apy

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/internal/models"
)

type TCalculateVeloAPYDataStruct struct {
	vault          models.TVault
	gaugeAddress   common.Address
	strategy       *models.TStrategy
	baseAssetPrice *bigNumber.Float
	poolPrice      *bigNumber.Float
	baseAPR        *bigNumber.Float
	rewardAPR      *bigNumber.Float
	poolDailyAPY   *bigNumber.Float
}

func calculateVeloForwardAPR(args TCalculateVeloAPYDataStruct, fraxPool TFraxPool) TStrategyAPR {
	/**********************************************************************************************
	** We will use the convexForwardAPR as a base for the Frax APR. So our first step is to
	** calculate the convexForwardAPR.
	**********************************************************************************************/
	baseConvexStrategyData := calculateConvexForwardAPR(TCalculateConvexAPYDataStruct(args))

	/**********************************************************************************************
	** We then need to add the minimum rewards APR to the convexForwardAPR to get the Frax APR.
	** The minimum rewards APR is the minimum amount of rewards we get from the Frax pool.
	**********************************************************************************************/
	minRewardsAPR := bigNumber.NewFloat(0).SetString(fraxPool.TotalRewardAPRs.Min)
	minRewardsAPR = bigNumber.NewFloat(0).Div(minRewardsAPR, bigNumber.NewFloat(100))

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
