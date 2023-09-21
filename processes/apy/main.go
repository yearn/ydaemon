package apy

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/store"
	"github.com/yearn/ydaemon/internal/meta"
	"github.com/yearn/ydaemon/internal/strategies"
	"github.com/yearn/ydaemon/internal/vaults"
	"github.com/yearn/ydaemon/processes/initDailyBlock"
)

var COMPUTED_APR = make(map[uint64]map[common.Address]TAPIV1APY)

func ComputeChainAPR(chainID uint64) {
	allVaults := vaults.ListVaults(chainID)
	gauges := retrieveCurveGauges(chainID)
	pools := retrieveCurveGetPools(chainID)
	subgraphData := retrieveCurveSubgraphData(chainID)
	fraxPools := retrieveFraxPools(chainID)

	if COMPUTED_APR[chainID] == nil {
		COMPUTED_APR[chainID] = make(map[common.Address]TAPIV1APY)
	}

	for _, vault := range allVaults {
		vaultFromMeta, ok := meta.GetMetaVault(chainID, vault.Address)
		if ok {
			if vaultFromMeta.Retired {
				continue
			}
		}

		allStrategiesForVault := strategies.ListStrategiesForVault(chainID, vault.Address)
		ppsPerTime, _ := store.ListPricePerShare(chainID, vault.Address)
		ppsInception := bigNumber.NewFloat(1)
		ppsToday := helpers.GetToday(ppsPerTime, vault.Decimals)
		if ppsToday == nil || ppsToday.IsZero() {
			vaultContract, err := contracts.NewYearnVaultCaller(vault.Address, ethereum.GetRPC(chainID))
			if err != nil {
				logs.Error("Could not get vault contract for " + vault.Address.Hex())
				continue
			}
			pps, _ := vaultContract.PricePerShare(nil)
			ppsToday = helpers.ToNormalizedAmount(bigNumber.SetInt(pps), vault.Decimals)
		}
		ppsWeekAgo := helpers.GetLastWeek(ppsPerTime, vault.Decimals)
		ppsMonthAgo := helpers.GetLastMonth(ppsPerTime, vault.Decimals)

		/**********************************************************************************************
		** Retrieve the vault performance fee and management fee, and calculate the net APR.
		** This can change from one vault to another and will be used to calculate the net APR.
		**********************************************************************************************/
		vaultPerformanceFee := helpers.ToNormalizedAmount(bigNumber.NewInt(int64(vault.PerformanceFee)), 4)
		vaultManagementFee := helpers.ToNormalizedAmount(bigNumber.NewInt(int64(vault.ManagementFee)), 4)
		oneMinusPerfFee := bigNumber.NewFloat(0).Sub(bigNumber.NewFloat(1), vaultPerformanceFee)

		/**********************************************************************************************
		** The grossAPR is used by checking the price per share evolution over a 30 days period.
		** The netAPY is the grossAPR minus the performance fee and the management fee.
		**********************************************************************************************/
		grossAPR := helpers.GetAPR(ppsToday, ppsMonthAgo, bigNumber.NewFloat(30))
		netAPR := bigNumber.NewFloat(0).Mul(grossAPR, oneMinusPerfFee)
		netAPR = bigNumber.NewFloat(0).Sub(netAPR, vaultManagementFee)
		if netAPR.Lte(ZERO) {
			netAPR = bigNumber.NewFloat(0)
		}

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
		vaultAPY := TAPIV1APY{
			Type:              "v2:averaged",
			GrossAPR:          grossAPR,
			NetAPY:            netAPR,
			StakingRewardsAPR: stakingRewardAPR,
			Fees: TAPIV1Fees{
				Performance: vaultPerformanceFee,
				Management:  vaultManagementFee,
			},
			Points: TAPIV1Points{
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
			vaultAPY.ForwardAPR = computeCurveLikeForwardAPR(
				vault,
				allStrategiesForVault,
				gauges,
				pools,
				subgraphData,
				fraxPools,
			)
		}
		if veloPool, ok := isVeloVault(chainID, vault); ok {
			vaultAPY.ForwardAPR = computeVeloLikeForwardAPR(
				vault,
				allStrategiesForVault,
				veloPool,
			)
		}

		if aeroPool, ok := isAeroVault(chainID, vault); ok {
			vaultAPY.ForwardAPR = computeVeloLikeForwardAPR(
				vault,
				allStrategiesForVault,
				aeroPool,
			)
		}

		COMPUTED_APR[chainID][vault.Address] = vaultAPY
		// aggregatedVault, okLegacyAPI := externalVaults.GetAggregatedVault(chainID, vault.Address.Hex())
		// legacyAPY := vaults.BuildAPY(vault, aggregatedVault, okLegacyAPI)
		// spew.Printf("\n%s (%s) - (%s)\n", vault.Name, vault.Address.Hex(), vaultAPY.Type)
		// checkDiff("GrossAPR      ", legacyAPY.GrossAPR, vaultAPY.ForwardAPR.GrossAPR)
		// checkDiff("NetAPY        ", legacyAPY.NetAPY, vaultAPY.ForwardAPR.NetAPY)
		// checkDiff("Mngt Fee      ", legacyAPY.Fees.Management, vaultAPY.Fees.Management)
		// checkDiff("Perf Fee      ", legacyAPY.Fees.Performance, vaultAPY.Fees.Performance)
		// checkDiff("PPS Week      ", legacyAPY.Points.WeekAgo, vaultAPY.Points.WeekAgo)
		// checkDiff("PPS Month     ", legacyAPY.Points.MonthAgo, vaultAPY.Points.MonthAgo)
		// checkDiff("PPS Inception ", legacyAPY.Points.Inception, vaultAPY.Points.Inception)
		// checkDiff("BaseAPR       ", legacyAPY.Composite.BaseAPR, vaultAPY.Composite.BaseAPR)
		// checkDiff("Boost         ", legacyAPY.Composite.Boost, vaultAPY.Composite.Boost)
		// checkDiff("BoostedAPR    ", legacyAPY.Composite.BoostedAPR, vaultAPY.Composite.BoostedAPR)
		// checkDiff("PoolAPY       ", legacyAPY.Composite.PoolAPY, vaultAPY.Composite.PoolAPY)
		// checkDiff("CvxAPR        ", legacyAPY.Composite.CvxAPR, vaultAPY.Composite.CvxAPR)
		// checkDiff("RewardsAPR    ", legacyAPY.Composite.RewardsAPR, vaultAPY.Composite.RewardsAPR)
		// spew.Printf("\n")
	}

	/**********************************************************************************************
	** Some logging to check if the computed APR is correct.
	** Debug
	**********************************************************************************************/
	// for _, vault := range allVaults {
	// 	vaultFromMeta, ok := meta.GetMetaVault(chainID, vault.Address)
	// 	if ok {
	// 		if vaultFromMeta.Retired {
	// 			continue
	// 		}
	// 	}

	// 	if _, ok := COMPUTED_APR[chainID]; !ok {
	// 		// logs.Warning("No computed APR for " + vault.Address.Hex() + " | " + vault.Name)
	// 		continue
	// 	}
	// 	if _, ok := COMPUTED_APR[chainID][vault.Address]; !ok {
	// 		// logs.Warning("No computed APR for " + vault.Address.Hex() + " | " + vault.Name)
	// 		continue
	// 	}
	// 	vaultAPY := COMPUTED_APR[chainID][vault.Address]
	// 	aggregatedVault, okLegacyAPI := extVaults.GetAggregatedVault(chainID, addresses.ToAddress(vault.Address).Hex())
	// 	internalAPY := vaults.BuildAPY(vault, aggregatedVault, okLegacyAPI)
	// 	if COMPUTED_APR[chainID][vault.Address].Points.MonthAgo.IsZero() {
	// 		if internalAPY.Points.MonthAgo == 0 {
	// 			// logs.Info("Zero APR on both sources for " + vault.Address.Hex() + " | " + vault.Name)
	// 		} else {
	// 			// logs.Warning("Zero APR for " + vault.Address.Hex() + " | " + vault.Name)
	// 		}
	// 	}
	// 	spew.Printf("\n%s (%s) - (%s)\n", vault.Name, vault.Address.Hex(), vaultAPY.Type)
	// 	checkDiff("PPS Week          ", internalAPY.Points.WeekAgo, vaultAPY.Points.WeekAgo)
	// 	checkDiff("PPS Month         ", internalAPY.Points.MonthAgo, vaultAPY.Points.MonthAgo)
	// 	checkDiff("PPS Inception     ", internalAPY.Points.Inception, vaultAPY.Points.Inception)
	// 	if vaultAPY.ForwardAPR.Type == "" {
	// 		checkDiff("Gross/Forward APR ", internalAPY.GrossAPR, vaultAPY.GrossAPR)
	// 	} else {
	// 		checkDiff("Gross/Forward APR ", internalAPY.GrossAPR, vaultAPY.ForwardAPR.GrossAPR)
	// 	}
	// 	checkDiff("Staking APR       ", internalAPY.StakingRewardsAPR, vaultAPY.StakingRewardsAPR)
	// 	spew.Printf("\n")
	// }

	//dump computed APR in a file
	// file, err := os.Create("computedAPR.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()
	// json.NewEncoder(file).Encode(COMPUTED_APR)

	//Create a CSV file with all the computed APR with cols vaultAddress, pointWeek, pointMonth, pointInception, grossAPR, netAPY
	// file, err = os.Create("computedAPR.csv")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()
	// file.WriteString("vaultAddress,pointWeek,pointWeekExporter,pointMonth,pointMonthExporter,pointInception,pointInceptionExporter,grossAPR,forwardAPR,grossAPRExporter,netAPY,netAPYExporter\n")
	// for _, vault := range allVaults {
	// 	vaultFromMeta, ok := meta.GetMetaVault(chainID, vault.Address)
	// 	if ok {
	// 		if vaultFromMeta.Retired {
	// 			continue
	// 		}
	// 	}

	// 	if _, ok := COMPUTED_APR[chainID]; !ok {
	// 		continue
	// 	}
	// 	if _, ok := COMPUTED_APR[chainID][vault.Address]; !ok {
	// 		continue
	// 	}
	// 	vaultAPY := COMPUTED_APR[chainID][vault.Address]
	// 	aggregatedVault, okLegacyAPI := extVaults.GetAggregatedVault(1, addresses.ToAddress(vault.Address).Hex())
	// 	internalAPY := vaults.BuildAPY(vault, aggregatedVault, okLegacyAPI)
	// 	file.WriteString(vault.Address.Hex() + "," +
	// 		vaultAPY.Points.WeekAgo.String() + "," +
	// 		strconv.FormatFloat(internalAPY.Points.WeekAgo, 'f', -1, 64) + "," +
	// 		vaultAPY.Points.MonthAgo.String() + "," +
	// 		strconv.FormatFloat(internalAPY.Points.MonthAgo, 'f', -1, 64) + "," +
	// 		vaultAPY.Points.Inception.String() + "," +
	// 		strconv.FormatFloat(internalAPY.Points.Inception, 'f', -1, 64) + "," +
	// 		vaultAPY.GrossAPR.String() + "," +
	// 		vaultAPY.ForwardAPR.GrossAPR.String() + "," +
	// 		strconv.FormatFloat(internalAPY.GrossAPR, 'f', -1, 64) + "," +
	// 		vaultAPY.NetAPY.String() + "," +
	// 		strconv.FormatFloat(internalAPY.NetAPY, 'f', -1, 64) + "," +
	// 		"\n")
	// }

}

func Run(chainID uint64) {
	initYearnEcosystem(chainID)
	initDailyBlock.Run(chainID)
	ComputeChainAPR(chainID)
}
