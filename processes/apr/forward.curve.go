package apr

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

type TCalculateCurveAPYDataStruct struct {
	vault        models.TVault
	gaugeAddress common.Address
	strategy     models.TStrategy
	baseAPY      *bigNumber.Float
	rewardAPY    *bigNumber.Float
	poolAPY      *bigNumber.Float
}

func calculateCurveForwardAPY(args TCalculateCurveAPYDataStruct) TStrategyAPY {
	chainID := args.vault.ChainID

	/**********************************************************************************************
	** We first need to retrieve a bunch to be able to proceed:
	** - the yBoost, aka how much boost the Yearn voter has for this gauge
	** - the performanceFee and the managementFee for that vault
	** - the debtRatio for the strategy (aka the % of fund allocated to the strategy by the vault)
	**********************************************************************************************/
	yBoost := getCurveBoost(chainID, storage.YEARN_VOTER_ADDRESS[chainID], args.gaugeAddress)
	keepCrv := determineCurveKeepCRV(args.strategy)
	debtRatio := helpers.ToNormalizedAmount(args.strategy.LastDebtRatio, 4)
	vaultPerformanceFee := helpers.ToNormalizedAmount(bigNumber.NewInt(int64(args.vault.PerformanceFee)), 4)
	vaultManagementFee := helpers.ToNormalizedAmount(bigNumber.NewInt(int64(args.vault.ManagementFee)), 4)
	oneMinusPerfFee := bigNumber.NewFloat(0).Sub(bigNumber.NewFloat(1), vaultPerformanceFee)

	/**********************************************************************************************
	** The CRV APR is simply the baseAPR (aka how much CRV we get from the gauge) scaled by the
	** yBoost. We then add the extraRewards which are incentives/bribes on top of the base rewards.
	**********************************************************************************************/
	crvAPY := bigNumber.NewFloat(0).Mul(args.baseAPY, yBoost)  // baseAPR * yBoost
	crvAPY = bigNumber.NewFloat(0).Add(crvAPY, args.rewardAPY) // (baseAPR * yBoost) + rewardAPY

	/**********************************************************************************************
	** Calculate the CRV Gross APR:
	** 1. Taking the base APR, scaling it with the boost and removing the percentage of CRV we want
	** to keep
	** 2. Adding the rewards APR
	** 3. Adding the pool APY
	**********************************************************************************************/
	keepCRVRatio := bigNumber.NewFloat(0).Sub(storage.ONE, keepCrv) // 1 - keepCRV
	grossAPY := bigNumber.NewFloat(0).Mul(args.baseAPY, yBoost)     // 1 - baseAPR * yBoost
	grossAPY = bigNumber.NewFloat(0).Mul(grossAPY, keepCRVRatio)    // 1 - baseAPR * yBoost * keepCRV
	grossAPY = bigNumber.NewFloat(0).Add(grossAPY, args.rewardAPY)  // 2 - (baseAPR * yBoost * keepCRV) + rewardAPY
	grossAPY = bigNumber.NewFloat(0).Add(grossAPY, args.poolAPY)    // 3 - (baseAPR * yBoost * keepCRV) + rewardAPY + poolAPY

	/**********************************************************************************************
	** Calculate the CRV Net APR:
	** Take the gross APR and remove the performance fee and the management fee
	**********************************************************************************************/
	netAPY := bigNumber.NewFloat(0).Mul(grossAPY, oneMinusPerfFee) // grossAPY * (1 - perfFee)
	if netAPY.Gt(vaultManagementFee) {
		netAPY = bigNumber.NewFloat(0).Sub(netAPY, vaultManagementFee) // (grossAPY * (1 - perfFee)) - managementFee
	} else {
		netAPY = bigNumber.NewFloat(0)
	}

	apyStruct := TStrategyAPY{
		Type:      "crv",
		DebtRatio: debtRatio,
		NetAPY:    bigNumber.NewFloat(0).Mul(netAPY, debtRatio),
		Composite: TCompositeData{
			Boost:      bigNumber.NewFloat(0).Mul(bigNumber.NewFloat(0).Mul(yBoost, debtRatio), debtRatio),
			PoolAPY:    bigNumber.NewFloat(0).Mul(args.poolAPY, debtRatio),
			BoostedAPR: bigNumber.NewFloat(0).Mul(crvAPY, debtRatio),
			BaseAPR:    bigNumber.NewFloat(0).Mul(args.baseAPY, debtRatio),
			RewardsAPY: bigNumber.NewFloat(0).Mul(args.rewardAPY, debtRatio),
			KeepCRV:    keepCrv,
		},
	}
	return apyStruct
}
