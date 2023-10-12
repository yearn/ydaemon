package apy

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/store"
	"github.com/yearn/ydaemon/internal/storage"
	"github.com/yearn/ydaemon/processes/initDailyBlock"
)

var COMPUTED_APR = make(map[uint64]map[common.Address]TVaultAPR)

func ComputeChainAPR(chainID uint64) {
	allVaults, _ := storage.ListVaults(chainID)
	gauges := retrieveCurveGauges(chainID)
	pools := retrieveCurveGetPools(chainID)
	subgraphData := retrieveCurveSubgraphData(chainID)
	fraxPools := retrieveFraxPools(chainID)

	if COMPUTED_APR[chainID] == nil {
		COMPUTED_APR[chainID] = make(map[common.Address]TVaultAPR)
	}

	for _, vault := range allVaults {
		if vault.IsRetired {
			continue
		}
		vaultToken, ok := storage.GetERC20(vault.ChainID, vault.Address)
		if !ok {
			continue
		}

		allStrategiesForVault, _ := storage.ListStrategiesForVault(chainID, vault.Address)
		ppsPerTime, _ := store.ListPricePerShare(chainID, vault.Address)
		ppsInception := bigNumber.NewFloat(1)
		ppsToday := helpers.GetToday(ppsPerTime, vaultToken.Decimals)
		if ppsToday == nil || ppsToday.IsZero() {
			ppsToday = ethereum.FetchPPSToday(chainID, vault.Address, vault.Decimals)
		}

		ppsWeekAgo := helpers.GetLastWeek(ppsPerTime, vault.Decimals)
		if ppsWeekAgo == nil || ppsWeekAgo.IsZero() {
			ppsWeekAgo = ethereum.FetchPPSLastWeek(chainID, vault.Address, vault.Decimals)
		}

		ppsMonthAgo := helpers.GetLastMonth(ppsPerTime, vault.Decimals)
		if ppsMonthAgo == nil || ppsMonthAgo.IsZero() {
			ppsMonthAgo = ethereum.FetchPPSLastMonth(chainID, vault.Address, vault.Decimals)
		}

		/**********************************************************************************************
		** Retrieve the vault performance fee and management fee, and calculate the net APR.
		** This can change from one vault to another and will be used to calculate the net APR.
		**********************************************************************************************/
		vaultPerformanceFee := helpers.ToNormalizedAmount(bigNumber.NewInt(int64(vault.PerformanceFee)), 4)
		vaultManagementFee := helpers.ToNormalizedAmount(bigNumber.NewInt(int64(vault.ManagementFee)), 4)
		oneMinusPerfFee := bigNumber.NewFloat(0).Sub(bigNumber.NewFloat(1), vaultPerformanceFee)
		_ = oneMinusPerfFee

		/**********************************************************************************************
		** The grossAPR is used by checking the price per share evolution over a 30 days period.
		** The netAPY is the grossAPR minus the performance fee and the management fee.
		**********************************************************************************************/
		netAPR := helpers.GetAPR(ppsToday, ppsMonthAgo, bigNumber.NewFloat(30))

		/**********************************************************************************************
		** Some vaults may have a staking rewards system. If so, we need to calculate the APR for
		** this staking rewards system and add it to the netAPY.
		**********************************************************************************************/
		stakingRewardAPR := computeStakingRewardsAPR(chainID, vault)

		/**********************************************************************************************
		** As we now have the base APR information we can init our structure. This base structure MUST
		** contains:
		** - The type of the APR (always v2:averaged for now)
		** - The grossAPR
		** - The netAPY
		** - The fees (performance and management)
		** - The points (PPS evolution over time, for one week, one month and since inception)
		**********************************************************************************************/
		vaultAPR := TVaultAPR{
			Type:   "v2:averaged",
			NetAPR: netAPR,
			Extra: TExtraRewards{
				StakingRewardsAPR: stakingRewardAPR,
			},
			Fees: TFees{
				Performance: vaultPerformanceFee,
				Management:  vaultManagementFee,
			},
			Points: THistoricalPoints{
				WeekAgo:   helpers.GetAPR(ppsToday, ppsWeekAgo, bigNumber.NewFloat(7)),
				MonthAgo:  helpers.GetAPR(ppsToday, ppsMonthAgo, bigNumber.NewFloat(30)),
				Inception: helpers.GetAPR(ppsToday, ppsInception, bigNumber.NewFloat(365)),
			},
		}

		/**********************************************************************************************
		** If it's a Curve Vault (has a Curve, Convex or Frax strategy), we can estimate the forward
		** APR, aka the expected APR we will get for the upcoming period.
		** We need to compute it and store it in our ForwardAPR structure.
		**********************************************************************************************/
		if isCurveVault(allStrategiesForVault) {
			vaultAPR.ForwardAPR = computeCurveLikeForwardAPR(
				vault,
				allStrategiesForVault,
				gauges,
				pools,
				subgraphData,
				fraxPools,
			)
		}
		if veloPool, ok := isVeloVault(chainID, vault); ok {
			vaultAPR.ForwardAPR = computeVeloLikeForwardAPR(
				vault,
				allStrategiesForVault,
				veloPool,
			)
		}

		if aeroPool, ok := isAeroVault(chainID, vault); ok {
			vaultAPR.ForwardAPR = computeVeloLikeForwardAPR(
				vault,
				allStrategiesForVault,
				aeroPool,
			)
		}

		COMPUTED_APR[chainID][vault.Address] = vaultAPR
	}
}

func Run(chainID uint64) {
	initYearnEcosystem(chainID)
	initDailyBlock.Run(chainID)
	ComputeChainAPR(chainID)
}
