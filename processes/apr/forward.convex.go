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
	apyTrace("forward.convex", chainID, args.vault.Address, args.strategy.Address, "input.gaugeAddress", args.gaugeAddress.Hex())
	apyTrace("forward.convex", chainID, args.vault.Address, args.strategy.Address, "input.baseAssetPrice", args.baseAssetPrice)
	apyTrace("forward.convex", chainID, args.vault.Address, args.strategy.Address, "input.poolPrice", args.poolPrice)
	apyTrace("forward.convex", chainID, args.vault.Address, args.strategy.Address, "input.baseAPY", args.baseAPY)
	apyTrace("forward.convex", chainID, args.vault.Address, args.strategy.Address, "input.rewardAPY", args.rewardAPY)
	apyTrace("forward.convex", chainID, args.vault.Address, args.strategy.Address, "input.poolWeeklyAPY", args.poolWeeklyAPY)

	/**********************************************************************************************
	** We first need to retrieve a bunch to be able to proceed:
	** - the cvxBoost, aka how much boost the convex voter has for this gauge
	** - the performanceFee and the managementFee for that vault
	** - the debtRatio for the strategy (aka the % of fund allocated to the strategy by the vault)
	**********************************************************************************************/
	cvxBoost := getCurveBoost(chainID, storage.CONVEX_VOTER_ADDRESS[chainID], args.gaugeAddress)
	apyTrace("forward.convex", chainID, args.vault.Address, args.strategy.Address, "cvxBoost", cvxBoost)
	keepCrv := determineConvexKeepCRV(args.strategy)
	apyTrace("forward.convex", chainID, args.vault.Address, args.strategy.Address, "keepCrv", keepCrv)
	debtRatio := helpers.ToNormalizedAmount(args.strategy.LastDebtRatio, 4)
	apyTrace("forward.convex", chainID, args.vault.Address, args.strategy.Address, "debtRatio", debtRatio)
	vaultPerformanceFee := helpers.ToNormalizedAmount(bigNumber.NewInt(int64(args.vault.PerformanceFee)), 4)
	apyTrace("forward.convex", chainID, args.vault.Address, args.strategy.Address, "vaultPerformanceFee", vaultPerformanceFee)
	vaultManagementFee := helpers.ToNormalizedAmount(bigNumber.NewInt(int64(args.vault.ManagementFee)), 4)
	apyTrace("forward.convex", chainID, args.vault.Address, args.strategy.Address, "vaultManagementFee", vaultManagementFee)
	oneMinusPerfFee := bigNumber.NewFloat(0).Sub(bigNumber.NewFloat(1), vaultPerformanceFee)
	apyTrace("forward.convex", chainID, args.vault.Address, args.strategy.Address, "oneMinusPerfFee", oneMinusPerfFee)

	/**********************************************************************************************
	** The CRV APR is simply the baseAPY (aka how much CRV we get from the gauge) not based on
	** curve API but directly from the formula used by convex subgraph (link below). We are using
	** this for th APR to match convex APR and avoid confusion. This should be very close to the
	** one we could compute.
	** The CVX APR is the APR we get from the CVX rewards, based on the CVX price and the amount
	** of CVX printed, based on the CRV rate for the given gauge. Tldr X crv = Y cvx and we do
	** something to gt an APR.
	**********************************************************************************************/
	crvAPR, cvxAPR, crvAPY, cvxAPY := getCVXPoolAPY(chainID, args.vault.Address, args.strategy.Address, args.baseAssetPrice)
	apyTrace("forward.convex", chainID, args.vault.Address, args.strategy.Address, "crvAPR", crvAPR)
	apyTrace("forward.convex", chainID, args.vault.Address, args.strategy.Address, "cvxAPR", cvxAPR)
	apyTrace("forward.convex", chainID, args.vault.Address, args.strategy.Address, "crvAPY", crvAPY)
	apyTrace("forward.convex", chainID, args.vault.Address, args.strategy.Address, "cvxAPY", cvxAPY)

	/**********************************************************************************************
	** Just like curve, Convex can have extra rewards which are incentives/bribes on top of the
	** base rewards. We need to retrieve them.
	**********************************************************************************************/
	_, rewardsAPY := getConvexRewardAPY(chainID, args.strategy, args.baseAssetPrice, args.poolPrice)
	apyTrace("forward.convex", chainID, args.vault.Address, args.strategy.Address, "rewardsAPY", rewardsAPY)

	/**********************************************************************************************
	** Calculate the CRV Gross APR:
	** 1. Taking the base APR and removing the percentage of CRV we want to keep
	** 2. Adding the rewards APR
	** 3. Adding the pool APY
	** 4. Adding the CVX APR
	**********************************************************************************************/
	keepCRVRatio := bigNumber.NewFloat(0).Sub(storage.ONE, keepCrv) // 1 - keepCRV
	apyTrace("forward.convex", chainID, args.vault.Address, args.strategy.Address, "keepCRVRatio", keepCRVRatio)
	grossAPY := bigNumber.NewFloat(0).Mul(crvAPY, keepCRVRatio)     // 1 - baseAPY * keepCRV
	apyTrace("forward.convex", chainID, args.vault.Address, args.strategy.Address, "grossAPY.afterKeepCRV", grossAPY)
	grossAPY = bigNumber.NewFloat(0).Add(grossAPY, rewardsAPY)      // 2 - (baseAPY * keepCRV) + rewardAPR
	apyTrace("forward.convex", chainID, args.vault.Address, args.strategy.Address, "grossAPY.afterRewards", grossAPY)
	grossAPY = bigNumber.NewFloat(0).Add(grossAPY, cvxAPY)          // 4 - (baseAPY * keepCRV) + rewardAPR + poolAPY + cvxAPR
	apyTrace("forward.convex", chainID, args.vault.Address, args.strategy.Address, "grossAPY.afterCvx", grossAPY)

	/**********************************************************************************************
	** Calculate the CRV Net APR:
	** Take the gross APR and remove the performance fee and the management fee
	**********************************************************************************************/
	netAPY := bigNumber.NewFloat(0).Mul(grossAPY, oneMinusPerfFee) // grossAPR * (1 - perfFee)
	apyTrace("forward.convex", chainID, args.vault.Address, args.strategy.Address, "netAPY.preMgmt", netAPY)
	if netAPY.Gt(vaultManagementFee) {
		netAPY = bigNumber.NewFloat(0).Sub(netAPY, vaultManagementFee) // (grossAPR * (1 - perfFee)) - managementFee
		apyTrace("forward.convex", chainID, args.vault.Address, args.strategy.Address, "netAPY.postMgmt", netAPY)
		netAPRFloat64, _ := netAPY.Float64()
		apyTrace("forward.convex", chainID, args.vault.Address, args.strategy.Address, "netAPY.netAPRFloat64", netAPRFloat64)
		netAPY = bigNumber.NewFloat(0).SetFloat64(convertFloatAPRToAPY(netAPRFloat64, 52))
		apyTrace("forward.convex", chainID, args.vault.Address, args.strategy.Address, "netAPY.compounded", netAPY)
		netAPY = bigNumber.NewFloat(0).Add(netAPY, args.poolWeeklyAPY)
		apyTrace("forward.convex", chainID, args.vault.Address, args.strategy.Address, "netAPY.withPool", netAPY)
	} else {
		netAPY = bigNumber.NewFloat(0).Add(bigNumber.NewFloat(0), args.poolWeeklyAPY)
		apyTrace("forward.convex", chainID, args.vault.Address, args.strategy.Address, "netAPY.poolOnly", netAPY)
	}

	netAPYWithDebtRatio := bigNumber.NewFloat(0).Mul(netAPY, debtRatio)
	apyTrace("forward.convex", chainID, args.vault.Address, args.strategy.Address, "result.netAPY", netAPY)
	apyTrace("forward.convex", chainID, args.vault.Address, args.strategy.Address, "result.netAPYWithDebtRatio", netAPYWithDebtRatio)
	boostWithDebtRatio := bigNumber.NewFloat(0).Mul(cvxBoost, debtRatio)
	apyTrace("forward.convex", chainID, args.vault.Address, args.strategy.Address, "result.boostWithDebtRatio", boostWithDebtRatio)
	poolAPYWithDebtRatio := bigNumber.NewFloat(0).Mul(args.poolWeeklyAPY, debtRatio)
	apyTrace("forward.convex", chainID, args.vault.Address, args.strategy.Address, "result.poolAPYWithDebtRatio", poolAPYWithDebtRatio)
	boostedAPRWithDebtRatio := bigNumber.NewFloat(0).Mul(crvAPR, debtRatio)
	apyTrace("forward.convex", chainID, args.vault.Address, args.strategy.Address, "result.boostedAPRWithDebtRatio", boostedAPRWithDebtRatio)
	baseAPRWithDebtRatio := bigNumber.NewFloat(0).Mul(args.baseAPY, debtRatio)
	apyTrace("forward.convex", chainID, args.vault.Address, args.strategy.Address, "result.baseAPRWithDebtRatio", baseAPRWithDebtRatio)
	cvxAPRWithDebtRatio := bigNumber.NewFloat(0).Mul(cvxAPR, debtRatio)
	apyTrace("forward.convex", chainID, args.vault.Address, args.strategy.Address, "result.cvxAPRWithDebtRatio", cvxAPRWithDebtRatio)
	rewardsAPYWithDebtRatio := bigNumber.NewFloat(0).Mul(args.rewardAPY, debtRatio)
	apyTrace("forward.convex", chainID, args.vault.Address, args.strategy.Address, "result.rewardsAPYWithDebtRatio", rewardsAPYWithDebtRatio)

	apyStruct := TStrategyAPY{
		Type:      "convex",
		DebtRatio: debtRatio,
		NetAPY:    netAPYWithDebtRatio,
		Composite: TCompositeData{
			Boost:      boostWithDebtRatio,
			PoolAPY:    poolAPYWithDebtRatio,
			BoostedAPR: boostedAPRWithDebtRatio,
			BaseAPR:    baseAPRWithDebtRatio,
			CvxAPR:     cvxAPRWithDebtRatio,
			RewardsAPY: rewardsAPYWithDebtRatio,
			KeepCRV:    keepCrv,
		},
	}
	apyTrace("forward.convex", chainID, args.vault.Address, args.strategy.Address, "result.apyStruct", apyStruct)
	return apyStruct
}
