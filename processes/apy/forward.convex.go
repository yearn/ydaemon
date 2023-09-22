package apy

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/models"
)

type TCalculateConvexAPYDataStruct struct {
	vault          models.TVault
	gaugeAddress   common.Address
	strategy       *models.TStrategy
	baseAssetPrice *bigNumber.Float
	poolPrice      *bigNumber.Float
	baseAPR        *bigNumber.Float
	rewardAPR      *bigNumber.Float
	poolDailyAPY   *bigNumber.Float
}

func calculateConvexForwardAPR(args TCalculateConvexAPYDataStruct) TStrategyAPR {
	chainID := args.vault.ChainID

	/**********************************************************************************************
	** We first need to retrieve a bunch to be able to proceed:
	** - the cvxBoost, aka how much boost the convex voter has for this gauge
	** - the performanceFee and the managementFee for that vault
	** - the debtRatio for the strategy (aka the % of fund allocated to the strategy by the vault)
	**********************************************************************************************/
	cvxBoost := getCurveBoost(chainID, CONVEX_VOTER_ADDRESS[chainID], args.gaugeAddress)
	keepCrv := determineConvexKeepCRV(args.strategy)
	debtRatio := helpers.ToNormalizedAmount(args.strategy.DebtRatio, 4)
	vaultPerformanceFee := helpers.ToNormalizedAmount(bigNumber.NewInt(int64(args.vault.PerformanceFee)), 4)
	vaultManagementFee := helpers.ToNormalizedAmount(bigNumber.NewInt(int64(args.vault.ManagementFee)), 4)
	oneMinusPerfFee := bigNumber.NewFloat(0).Sub(bigNumber.NewFloat(1), vaultPerformanceFee)

	/**********************************************************************************************
	** The CRV APR is simply the baseAPR (aka how much CRV we get from the gauge) not based on
	** curve API but directly from the formula used by convex subgraph (link below). We are using
	** this for th APR to match convex APR and avoid confusion. This should be very close to the
	** one we could compute.
	** The CVX APR is the APR we get from the CVX rewards, based on the CVX price and the amount
	** of CVX printed, based on the CRV rate for the given gauge. Tldr X crv = Y cvx and we do
	** something to gt an APR.
	**********************************************************************************************/
	crvAPR, cvxAPR := getCVXPoolAPR(chainID, args.strategy.Address, args.baseAssetPrice)

	/**********************************************************************************************
	** Just like curve, Convex can have extra rewards which are incentives/bribes on top of the
	** base rewards. We need to retrieve them.
	**********************************************************************************************/
	rewardsAPR := getConvexRewardAPR(chainID, args.strategy, args.baseAssetPrice, args.poolPrice)

	/**********************************************************************************************
	** Calculate the CRV Gross APR:
	** 1. Taking the base APR and removing the percentage of CRV we want to keep
	** 2. Adding the rewards APR
	** 3. Adding the pool APY
	** 4. Adding the CVX APR
	**********************************************************************************************/
	keepCRVRatio := bigNumber.NewFloat(0).Sub(ONE, keepCrv)           // 1 - keepCRV
	grossAPR := bigNumber.NewFloat(0).Mul(crvAPR, keepCRVRatio)       // 1 - baseAPR * keepCRV
	grossAPR = bigNumber.NewFloat(0).Add(grossAPR, rewardsAPR)        // 2 - (baseAPR * keepCRV) + rewardAPR
	grossAPR = bigNumber.NewFloat(0).Add(grossAPR, args.poolDailyAPY) // 3 - (baseAPR * keepCRV) + rewardAPR + poolAPY
	grossAPR = bigNumber.NewFloat(0).Add(grossAPR, cvxAPR)            // 4 - (baseAPR * keepCRV) + rewardAPR + poolAPY + cvxAPR

	/**********************************************************************************************
	** Calculate the CRV Net APR:
	** Take the gross APR and remove the performance fee and the management fee
	**********************************************************************************************/
	netAPR := bigNumber.NewFloat(0).Mul(grossAPR, oneMinusPerfFee)
	netAPR = bigNumber.NewFloat(0).Sub(netAPR, vaultManagementFee)

	apyStruct := TStrategyAPR{
		Type:      "convex",
		DebtRatio: debtRatio,
		NetAPR:    bigNumber.NewFloat(0).Mul(netAPR, debtRatio),
		Composite: TCompositeData{
			Boost:      bigNumber.NewFloat(0).Mul(cvxBoost, debtRatio),
			PoolAPY:    bigNumber.NewFloat(0).Mul(args.poolDailyAPY, debtRatio),
			BoostedAPR: bigNumber.NewFloat(0).Mul(crvAPR, debtRatio),
			BaseAPR:    bigNumber.NewFloat(0).Mul(args.baseAPR, debtRatio),
			CvxAPR:     bigNumber.NewFloat(0).Mul(cvxAPR, debtRatio),
			RewardsAPR: bigNumber.NewFloat(0).Mul(args.rewardAPR, debtRatio),
		},
	}
	return apyStruct
}
