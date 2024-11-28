package apr

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/internal/models"
)

type TCalculateFraxAPYDataStruct struct {
	vault          models.TVault
	gaugeAddress   common.Address
	strategy       models.TStrategy
	baseAssetPrice *bigNumber.Float
	poolPrice      *bigNumber.Float
	baseAPY        *bigNumber.Float
	rewardAPY      *bigNumber.Float
	poolDailyAPY   *bigNumber.Float
}

func calculateFraxForwardAPR(args TCalculateFraxAPYDataStruct, fraxPool TFraxPool) TStrategyAPY {
	/**********************************************************************************************
	** We will use the convexForwardAPR as a base for the Frax APR. So our first step is to
	** calculate the convexForwardAPR.
	**********************************************************************************************/
	baseConvexStrategyData := calculateConvexForwardAPY(TCalculateConvexAPYDataStruct(args))

	/**********************************************************************************************
	** We then need to add the minimum rewards APR to the convexForwardAPR to get the Frax APR.
	** The minimum rewards APR is the minimum amount of rewards we get from the Frax pool.
	**********************************************************************************************/
	minRewardsAPR := bigNumber.NewFloat(0)
	switch fraxPool.TotalRewardAPRs.Min.(type) {
	case string:
		minRewardsAPR = bigNumber.NewFloat(0).SetString(fraxPool.TotalRewardAPRs.Min.(string))
	case float64:
		minRewardsAPR = bigNumber.NewFloat(0).SetFloat64(fraxPool.TotalRewardAPRs.Min.(float64))
	}
	minRewardsAPR = bigNumber.NewFloat(0).Div(minRewardsAPR, bigNumber.NewFloat(100))
	minRewardsAPRFloat64, _ := minRewardsAPR.Float64()
	minRewardsAPY := bigNumber.NewFloat(0).SetFloat64(convertFloatAPRToAPY(minRewardsAPRFloat64, 365/15))

	apyStruct := TStrategyAPY{
		Type:      "frax",
		DebtRatio: baseConvexStrategyData.DebtRatio,
		NetAPY:    bigNumber.NewFloat(0).Add(baseConvexStrategyData.NetAPY, minRewardsAPY),
		Composite: TCompositeData{
			Boost:      baseConvexStrategyData.Composite.Boost,
			PoolAPY:    baseConvexStrategyData.Composite.PoolAPY,
			BoostedAPR: baseConvexStrategyData.Composite.BoostedAPR,
			BaseAPR:    baseConvexStrategyData.Composite.BaseAPR,
			CvxAPR:     baseConvexStrategyData.Composite.CvxAPR,
			RewardsAPY: bigNumber.NewFloat(0).Add(baseConvexStrategyData.Composite.RewardsAPY, minRewardsAPY),
			KeepCRV:    baseConvexStrategyData.Composite.KeepCRV,
		},
	}
	return apyStruct
}
