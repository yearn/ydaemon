package apy

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/models"
)

type TCalculateCurveAPYDataStruct struct {
	vault        models.TVault
	gaugeAddress common.Address
	strategy     *models.TStrategy
	baseAPR      *bigNumber.Float
	rewardAPR    *bigNumber.Float
	poolAPY      *bigNumber.Float
}

func calculateCurveForwardAPR(args TCalculateCurveAPYDataStruct) TStrategyAPR {
	chainID := args.vault.ChainID

	/**********************************************************************************************
	** We first need to retrieve a bunch to be able to proceed:
	** - the yBoost, aka how much boost the Yearn voter has for this gauge
	** - the performanceFee and the managementFee for that vault
	** - the debtRatio for the strategy (aka the % of fund allocated to the strategy by the vault)
	**********************************************************************************************/
	yBoost := getCurveBoost(chainID, YEARN_VOTER_ADDRESS[chainID], args.gaugeAddress)
	keepCrv := determineCurveKeepCRV(args.strategy)
	debtRatio := helpers.ToNormalizedAmount(args.strategy.DebtRatio, 4)
	vaultPerformanceFee := helpers.ToNormalizedAmount(bigNumber.NewInt(int64(args.vault.PerformanceFee)), 4)
	vaultManagementFee := helpers.ToNormalizedAmount(bigNumber.NewInt(int64(args.vault.ManagementFee)), 4)
	oneMinusPerfFee := bigNumber.NewFloat(0).Sub(bigNumber.NewFloat(1), vaultPerformanceFee)

	/**********************************************************************************************
	** The CRV APR is simply the baseAPR (aka how much CRV we get from the gauge) scaled by the
	** yBoost. We then add the extraRewards which are incentives/bribes on top of the base rewards.
	**********************************************************************************************/
	crvAPR := bigNumber.NewFloat(0).Mul(args.baseAPR, yBoost)  // baseAPR * yBoost
	crvAPR = bigNumber.NewFloat(0).Add(crvAPR, args.rewardAPR) // (baseAPR * yBoost) + rewardAPR

	/**********************************************************************************************
	** Calculate the CRV Gross APR:
	** 1. Taking the base APR, scaling it with the boost and removing the percentage of CRV we want
	** to keep
	** 2. Adding the rewards APR
	** 3. Adding the pool APY
	**********************************************************************************************/
	keepCRVRatio := bigNumber.NewFloat(0).Sub(ONE, keepCrv)        // 1 - keepCRV
	grossAPR := bigNumber.NewFloat(0).Mul(args.baseAPR, yBoost)    // 1 - baseAPR * yBoost
	grossAPR = bigNumber.NewFloat(0).Mul(grossAPR, keepCRVRatio)   // 1 - baseAPR * yBoost * keepCRV
	grossAPR = bigNumber.NewFloat(0).Add(grossAPR, args.rewardAPR) // 2 - (baseAPR * yBoost * keepCRV) + rewardAPR
	grossAPR = bigNumber.NewFloat(0).Add(grossAPR, args.poolAPY)   // 3 - (baseAPR * yBoost * keepCRV) + rewardAPR + poolAPY

	/**********************************************************************************************
	** Calculate the CRV Net APR:
	** Take the gross APR and remove the performance fee and the management fee
	**********************************************************************************************/
	netAPR := bigNumber.NewFloat(0).Mul(grossAPR, oneMinusPerfFee) // grossAPR * (1 - perfFee)
	netAPR = bigNumber.NewFloat(0).Sub(netAPR, vaultManagementFee) // (grossAPR * (1 - perfFee)) - managementFee

	/**********************************************************************************************
	** Calculate the final boost
	**********************************************************************************************/
	boost := bigNumber.NewFloat(0).Mul(yBoost, debtRatio) // yBoost * debtRatioCurve

	apyStruct := TStrategyAPR{
		Type:      "crv",
		DebtRatio: debtRatio,
		NetAPR:    bigNumber.NewFloat(0).Mul(netAPR, debtRatio),
		Composite: TCompositeData{
			Boost:      bigNumber.NewFloat(0).Mul(boost, debtRatio),
			PoolAPY:    bigNumber.NewFloat(0).Mul(args.poolAPY, debtRatio),
			BoostedAPR: bigNumber.NewFloat(0).Mul(crvAPR, debtRatio),
			BaseAPR:    bigNumber.NewFloat(0).Mul(args.baseAPR, debtRatio),
			RewardsAPR: bigNumber.NewFloat(0).Mul(args.rewardAPR, debtRatio),
		},
	}
	return apyStruct
}
