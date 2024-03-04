package apr

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/internal/models"
)

type TCalculatePrismaAPYDataStruct struct {
	vault          models.TVault
	gaugeAddress   common.Address
	strategy       models.TStrategy
	baseAssetPrice *bigNumber.Float
	poolPrice      *bigNumber.Float
	baseAPR        *bigNumber.Float
	rewardAPR      *bigNumber.Float
	poolDailyAPY   *bigNumber.Float
}

func calculatePrismaForwardAPR(args TCalculatePrismaAPYDataStruct) TStrategyAPR {
	/**********************************************************************************************
	** If the strategy is not a prisma strategy, we can't calculate the prisma APR.
	**********************************************************************************************/
	prismaReceiver := fetchPrismaReceiver(args.vault.ChainID, args.strategy.Address)
	if (prismaReceiver == common.Address{}) {
		return TStrategyAPR{}
	}

	/**********************************************************************************************
	** We will use the convexForwardAPR as a base for the Prisma APR. So our first step is to
	** calculate the convexForwardAPR.
	**********************************************************************************************/
	baseConvexStrategyData := calculateConvexForwardAPR(TCalculateConvexAPYDataStruct(args))

	/**********************************************************************************************
	** Finally, we need to add the prisma APR to the convexForwardAPR to get the strategy forward
	** APR.
	**********************************************************************************************/
	prismaAPR := getPrismaAPR(args.vault.ChainID, prismaReceiver)

	apyStruct := TStrategyAPR{
		Type:      "prisma",
		DebtRatio: baseConvexStrategyData.DebtRatio,
		NetAPR:    bigNumber.NewFloat(0).Add(baseConvexStrategyData.NetAPR, prismaAPR),
		Composite: TCompositeData{
			Boost:      baseConvexStrategyData.Composite.Boost,
			PoolAPY:    baseConvexStrategyData.Composite.PoolAPY,
			BoostedAPR: baseConvexStrategyData.Composite.BoostedAPR,
			BaseAPR:    baseConvexStrategyData.Composite.BaseAPR,
			CvxAPR:     baseConvexStrategyData.Composite.CvxAPR,
			RewardsAPR: bigNumber.NewFloat(0).Add(baseConvexStrategyData.Composite.RewardsAPR, prismaAPR),
			KeepCRV:    baseConvexStrategyData.Composite.KeepCRV,
		},
	}
	return apyStruct
}
