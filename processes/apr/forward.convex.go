package apr

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

type TCalculateConvexAPYDataStruct struct {
	vault          models.TVault
	gaugeAddress   common.Address
	strategy       models.TStrategy
	baseAssetPrice *bigNumber.Float
	poolPrice      *bigNumber.Float
	baseAPY        *bigNumber.Float
	rewardAPY      *bigNumber.Float
	poolWeeklyAPY  *bigNumber.Float
}

func calculateConvexForwardAPY(args TCalculateConvexAPYDataStruct) TStrategyAPY {
	chainID := args.vault.ChainID

	/**********************************************************************************************
	** We first need to retrieve a bunch to be able to proceed:
	** - the cvxBoost, aka how much boost the convex voter has for this gauge
	** - the performanceFee and the managementFee for that vault
	** - the debtRatio for the strategy (aka the % of fund allocated to the strategy by the vault)
	**********************************************************************************************/
	cvxBoost := getCurveBoost(chainID, storage.CONVEX_VOTER_ADDRESS[chainID], args.gaugeAddress)
	keepCrv := determineConvexKeepCRV(args.strategy)
	debtRatio := helpers.ToNormalizedAmount(args.strategy.LastDebtRatio, 4)
	vaultPerformanceFee := helpers.ToNormalizedAmount(bigNumber.NewInt(int64(args.vault.PerformanceFee)), 4)
	vaultManagementFee := helpers.ToNormalizedAmount(bigNumber.NewInt(int64(args.vault.ManagementFee)), 4)
	oneMinusPerfFee := bigNumber.NewFloat(0).Sub(bigNumber.NewFloat(1), vaultPerformanceFee)

	/**********************************************************************************************
	** The CRV APR is simply the baseAPY (aka how much CRV we get from the gauge) not based on
	** curve API but directly from the formula used by convex subgraph (link below). We are using
	** this for th APR to match convex APR and avoid confusion. This should be very close to the
	** one we could compute.
	** The CVX APR is the APR we get from the CVX rewards, based on the CVX price and the amount
	** of CVX printed, based on the CRV rate for the given gauge. Tldr X crv = Y cvx and we do
	** something to gt an APR.
	**********************************************************************************************/
	crvAPR, cvxAPR, crvAPY, cvxAPY := getCVXPoolAPY(chainID, args.strategy.Address, args.baseAssetPrice)

	/**********************************************************************************************
	** Just like curve, Convex can have extra rewards which are incentives/bribes on top of the
	** base rewards. We need to retrieve them.
	**********************************************************************************************/
	_, rewardsAPY := getConvexRewardAPY(chainID, args.strategy, args.baseAssetPrice, args.poolPrice)

	/**********************************************************************************************
	** Calculate the CRV Gross APR:
	** 1. Taking the base APR and removing the percentage of CRV we want to keep
	** 2. Adding the rewards APR
	** 3. Adding the pool APY
	** 4. Adding the CVX APR
	**********************************************************************************************/
	keepCRVRatio := bigNumber.NewFloat(0).Sub(storage.ONE, keepCrv)    // 1 - keepCRV
	grossAPY := bigNumber.NewFloat(0).Mul(crvAPY, keepCRVRatio)        // 1 - baseAPY * keepCRV
	grossAPY = bigNumber.NewFloat(0).Add(grossAPY, rewardsAPY)         // 2 - (baseAPY * keepCRV) + rewardAPR
	grossAPY = bigNumber.NewFloat(0).Add(grossAPY, args.poolWeeklyAPY) // 3 - (baseAPY * keepCRV) + rewardAPR + poolAPY
	grossAPY = bigNumber.NewFloat(0).Add(grossAPY, cvxAPY)             // 4 - (baseAPY * keepCRV) + rewardAPR + poolAPY + cvxAPR

	/**********************************************************************************************
	** Calculate the CRV Net APR:
	** Take the gross APR and remove the performance fee and the management fee
	**********************************************************************************************/
	netAPY := bigNumber.NewFloat(0).Mul(grossAPY, oneMinusPerfFee) // grossAPR * (1 - perfFee)
	if netAPY.Gt(vaultManagementFee) {
		netAPY = bigNumber.NewFloat(0).Sub(netAPY, vaultManagementFee) // (grossAPR * (1 - perfFee)) - managementFee
	} else {
		netAPY = bigNumber.NewFloat(0)
	}

	apyStruct := TStrategyAPY{
		Type:      "convex",
		DebtRatio: debtRatio,
		NetAPY:    bigNumber.NewFloat(0).Mul(netAPY, debtRatio),
		Composite: TCompositeData{
			Boost:      bigNumber.NewFloat(0).Mul(cvxBoost, debtRatio),
			PoolAPY:    bigNumber.NewFloat(0).Mul(args.poolWeeklyAPY, debtRatio),
			BoostedAPR: bigNumber.NewFloat(0).Mul(crvAPR, debtRatio),
			BaseAPR:    bigNumber.NewFloat(0).Mul(args.baseAPY, debtRatio),
			CvxAPR:     bigNumber.NewFloat(0).Mul(cvxAPR, debtRatio),
			RewardsAPY: bigNumber.NewFloat(0).Mul(args.rewardAPY, debtRatio),
			KeepCRV:    keepCrv,
		},
	}
	return apyStruct
}
