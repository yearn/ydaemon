package apy

import (
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/store"
	extVaults "github.com/yearn/ydaemon/external/vaults"
	"github.com/yearn/ydaemon/internal/meta"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/strategies"
	"github.com/yearn/ydaemon/internal/vaults"
	"github.com/yearn/ydaemon/processes/initDailyBlock"
)

var COMPUTED_APR = make(map[uint64]map[common.Address]TAPIV1APY)

func calculateGaugeBaseAPR(
	gauge models.CurveGauge,
	crvTokenPrice *bigNumber.Float,
	poolPrice *bigNumber.Float,
	baseAssetPrice *bigNumber.Float,
) *bigNumber.Float {
	// check if type of gauge.GaugeController.InflationRate is string or number and convert to bigNumber.Float

	inflationRate := bigNumber.NewFloat(0)
	switch gauge.GaugeController.InflationRate.(type) {
	case string:
		inflationRate = helpers.ToNormalizedAmount(bigNumber.NewInt(0).SetString(gauge.GaugeController.InflationRate.(string)), 18)
	case float64:
		inflationRate = bigNumber.NewFloat(0).SetFloat64(gauge.GaugeController.InflationRate.(float64))
	}

	gaugeWeight := helpers.ToNormalizedAmount(bigNumber.NewInt(0).SetString(gauge.GaugeController.GaugeRelativeWeight), 18)
	secondPerYear := bigNumber.NewFloat(0).SetFloat64(31556952)
	workingSupply := helpers.ToNormalizedAmount(bigNumber.NewInt(0).SetString(gauge.GaugeData.WorkingSupply), 18)
	perMaxBoost := bigNumber.NewFloat(0).SetFloat64(0.4)
	crvPrice := bigNumber.NewFloat(0).Clone(crvTokenPrice)

	// baseAPR = (inflationRate * gaugeWeight * (secondPerYear / workingSupply) * (perMaxBoost / poolPrice) * crvPrice) / baseAssetPrice
	baseAPR := bigNumber.NewFloat(0).Mul(inflationRate, gaugeWeight)
	baseAPR = bigNumber.NewFloat(0).Mul(baseAPR, bigNumber.NewFloat(0).Div(secondPerYear, workingSupply))
	baseAPR = bigNumber.NewFloat(0).Mul(baseAPR, bigNumber.NewFloat(0).Div(perMaxBoost, poolPrice))
	baseAPR = bigNumber.NewFloat(0).Mul(baseAPR, crvPrice)
	baseAPR = bigNumber.NewFloat(0).Div(baseAPR, baseAssetPrice)

	return baseAPR
}

func calculateStrategyAPR(
	vault models.TVault,
	strategy *models.TStrategy,
	gauge models.CurveGauge,
	pool models.CurvePool,
	fraxPool TFraxPool,
	subgraphItem models.CurveSubgraphData,
) TStrategyAPR {
	/**********************************************************************************************
	** First thing is to retrieve the data we need from curve API. This includes the pool, the
	** gauges and some APY data from their subgraph.
	**********************************************************************************************/
	chainID := vault.ChainID

	if subgraphItem.LatestWeeklyApy == 0 {
		logs.Error(`No APY data for vault ` + vault.Address.Hex())
		// logs.Pretty(subgraphItem)
		return TStrategyAPR{}
	}

	/**********************************************************************************************
	** We will need a bunch of prices to calculate the APY.
	** - We get the base asset price from the gauge LpTokenPrice field, which is in base 2.
	** - We get the CRV price from our price package, which is in base 6 but converted in 2.
	** - We get the pool price from the curve API, which is in base 18 but converted in 2.
	**********************************************************************************************/
	baseAssetPrice := bigNumber.NewFloat(0).SetFloat64(gauge.LpTokenPrice)
	crvPrice := getTokenPrice(chainID, CRV_TOKEN_ADDRESS[chainID])
	poolPrice := getPoolPrice(gauge)

	/**********************************************************************************************
	** With theses info and the gauges data, we cal calculate both the base APR and the curve
	** boost for the Yearn Voter.
	**********************************************************************************************/
	baseAPR := calculateGaugeBaseAPR(gauge, crvPrice, poolPrice, baseAssetPrice)

	/**********************************************************************************************
	** Once it's done, we need to calculate the rewards APR. Each gauge can get bribes, which are
	** extra rewards on top of the normal rewards.
	** An unknown amount of unknown tokens can be distributed to the gauge, but we still want to
	** calculate the APR for them.
	**********************************************************************************************/
	rewardAPR := getRewardsAPR(chainID, pool)

	/**********************************************************************************************
	** We need to know how is performing the pool in term of APY. We could calculate the whole what
	** is the price now, vs what it was one week ago, and cale it to one year etc. but the weekly
	** apy is already available in the Curve API.
	** So just use it.
	**********************************************************************************************/
	poolWeeklyAPY := getPoolWeeklyAPY(subgraphItem)
	poolDailyAPY := getPoolDailyAPY(subgraphItem)

	/**********************************************************************************************
	** Now, we need to know if we are in a standard curve strat, or a convex strat. If it's a
	** convex strat, we need to calculate the convex APY instead of the standard one.
	** If it's a Frax strategy, we need to calculate the Frax APY instead of the standard one.
	**********************************************************************************************/
	if isFraxStrategy(strategy) {
		return calculateFraxForwardAPR(
			TCalculateFraxAPYDataStruct{
				vault:          vault,
				gaugeAddress:   common.HexToAddress(gauge.Gauge),
				strategy:       strategy,
				baseAssetPrice: baseAssetPrice,
				poolPrice:      poolPrice,
				baseAPR:        baseAPR,
				rewardAPR:      rewardAPR,
				poolDailyAPY:   poolDailyAPY,
			},
			fraxPool,
		)
	}
	if isConvexStrategy(strategy) {
		return calculateConvexForwardAPR(
			TCalculateConvexAPYDataStruct{
				vault:          vault,
				gaugeAddress:   common.HexToAddress(gauge.Gauge),
				strategy:       strategy,
				baseAssetPrice: baseAssetPrice,
				poolPrice:      poolPrice,
				baseAPR:        baseAPR,
				rewardAPR:      rewardAPR,
				poolDailyAPY:   poolDailyAPY,
			},
		)
	}

	return calculateCurveForwardAPR(
		TCalculateCurveAPYDataStruct{
			vault:        vault,
			gaugeAddress: common.HexToAddress(gauge.Gauge),
			strategy:     strategy,
			baseAPR:      baseAPR,
			rewardAPR:    rewardAPR,
			poolAPY:      poolWeeklyAPY,
		},
	)
}

func computeForwardAPR(
	vault models.TVault,
	allStrategiesForVault []*models.TStrategy,
	gauges []models.CurveGauge,
	pools []models.CurvePool,
	subgraphData []models.CurveSubgraphData,
	fraxPools []TFraxPool,
) TForwardAPR {
	gauge := findGaugeForVault(vault.Token.Address, gauges)
	pool := findPoolForVault(vault.Token.Address, pools)
	fraxPool := findFraxPoolForVault(vault.Token.Address, fraxPools)
	subgraphItem := findSubgraphItemForVault(common.HexToAddress(gauge.Swap), subgraphData)

	TypeOf := ``
	GrossAPR := bigNumber.NewFloat(0)
	NetAPY := bigNumber.NewFloat(0)
	Boost := bigNumber.NewFloat(0)
	PoolAPY := bigNumber.NewFloat(0)
	BoostedAPR := bigNumber.NewFloat(0)
	BaseAPR := bigNumber.NewFloat(0)
	CvxAPR := bigNumber.NewFloat(0)
	RewardsAPR := bigNumber.NewFloat(0)
	for _, strategy := range allStrategiesForVault {
		if strategy.DebtRatio == nil || strategy.DebtRatio.IsZero() {
			logs.Info("Skipping strategy " + strategy.Address.Hex() + " for vault " + vault.Address.Hex() + " because debt ratio is zero")
			continue
		}
		if strategy.TotalDebt == nil || strategy.TotalDebt.IsZero() {
			logs.Info("Skipping strategy " + strategy.Address.Hex() + " for vault " + vault.Address.Hex() + " because total debt is zero")
			continue
		}

		strategyAPR := calculateStrategyAPR(
			vault,
			strategy,
			gauge,
			pool,
			fraxPool,
			subgraphItem,
		)
		TypeOf += strings.TrimSpace(` ` + strategyAPR.Type)
		GrossAPR = bigNumber.NewFloat(0).Add(GrossAPR, strategyAPR.GrossAPR)
		NetAPY = bigNumber.NewFloat(0).Add(NetAPY, strategyAPR.NetAPY)
		Boost = bigNumber.NewFloat(0).Add(Boost, strategyAPR.Composite.Boost)
		PoolAPY = bigNumber.NewFloat(0).Add(PoolAPY, strategyAPR.Composite.PoolAPY)
		BoostedAPR = bigNumber.NewFloat(0).Add(BoostedAPR, strategyAPR.Composite.BoostedAPR)
		BaseAPR = bigNumber.NewFloat(0).Add(BaseAPR, strategyAPR.Composite.BaseAPR)
		CvxAPR = bigNumber.NewFloat(0).Add(CvxAPR, strategyAPR.Composite.CvxAPR)
		RewardsAPR = bigNumber.NewFloat(0).Add(RewardsAPR, strategyAPR.Composite.RewardsAPR)
	}
	return TForwardAPR{
		Type:     strings.TrimSpace(TypeOf),
		GrossAPR: GrossAPR,
		NetAPY:   NetAPY,
		Composite: TAPIV1Composite{
			Boost:      Boost,
			PoolAPY:    PoolAPY,
			BoostedAPR: BoostedAPR,
			BaseAPR:    BaseAPR,
			CvxAPR:     CvxAPR,
			RewardsAPR: RewardsAPR,
		},
	}
}

func ComputeChainAPR(chainID uint64) {
	if chainID != 1 && chainID != 10 {
		return
	}
	allVaults := vaults.ListVaults(chainID)
	gauges := retrieveCurveGauges(chainID)
	pools := retrieveCurveGetPools(chainID)
	subgraphData := retrieveCurveSubgraphData(chainID)
	fraxPools := retrieveFraxPools(chainID)

	if COMPUTED_APR[chainID] == nil {
		COMPUTED_APR[chainID] = make(map[common.Address]TAPIV1APY)
	}

	for _, vault := range allVaults {
		computeStakingRewardsAPR(chainID, vault)
		// if !addresses.Equals(vault.Address, common.HexToAddress("0x7788A5492bc948e1d8c2caa53b2e0a60ed5403b0")) {
		// 	continue
		// }
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
			Type:     "v2:averaged",
			GrossAPR: grossAPR,
			NetAPY:   netAPR,
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
			vaultAPY.ForwardAPR = computeForwardAPR(
				vault,
				allStrategiesForVault,
				gauges,
				pools,
				subgraphData,
				fraxPools,
			)
		}

		COMPUTED_APR[chainID][vault.Address] = vaultAPY

		// aggregatedVault, okLegacyAPI := externalVaults.GetAggregatedVault(chainID, vault.Address.Hex())
		// legacyAPY := vaults.BuildAPY(vault, aggregatedVault, okLegacyAPI)
		// spew.Printf("\n%s (%s) - (%s)\n", vault.Name, vault.Address.Hex(), vaultAPY.Type)
		// checkDiff("GrossAPR      ", legacyAPY.GrossAPR, vaultAPY.GrossAPR)
		// checkDiff("NetAPY        ", legacyAPY.NetAPY, vaultAPY.NetAPY)
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
	logs.Success(`Computing APR done`)

	/**********************************************************************************************
	** Some logging to check if the computed APR is correct.
	** Debug
	**********************************************************************************************/
	for _, vault := range allVaults {
		vaultFromMeta, ok := meta.GetMetaVault(chainID, vault.Address)
		if ok {
			if vaultFromMeta.Retired {
				continue
			}
		}

		if _, ok := COMPUTED_APR[chainID]; !ok {
			// logs.Warning("No computed APR for " + vault.Address.Hex() + " | " + vault.Name)
			continue
		}
		if _, ok := COMPUTED_APR[chainID][vault.Address]; !ok {
			// logs.Warning("No computed APR for " + vault.Address.Hex() + " | " + vault.Name)
			continue
		}
		vaultAPY := COMPUTED_APR[chainID][vault.Address]
		aggregatedVault, okLegacyAPI := extVaults.GetAggregatedVault(chainID, addresses.ToAddress(vault.Address).Hex())
		internalAPY := vaults.BuildAPY(vault, aggregatedVault, okLegacyAPI)
		if COMPUTED_APR[chainID][vault.Address].Points.MonthAgo.IsZero() {
			if internalAPY.Points.MonthAgo == 0 {
				// logs.Info("Zero APR on both sources for " + vault.Address.Hex() + " | " + vault.Name)
			} else {
				// logs.Warning("Zero APR for " + vault.Address.Hex() + " | " + vault.Name)
			}
		}
		spew.Printf("\n%s (%s) - (%s)\n", vault.Name, vault.Address.Hex(), vaultAPY.Type)
		checkDiff("PPS Week          ", internalAPY.Points.WeekAgo, vaultAPY.Points.WeekAgo)
		checkDiff("PPS Month         ", internalAPY.Points.MonthAgo, vaultAPY.Points.MonthAgo)
		checkDiff("PPS Inception     ", internalAPY.Points.Inception, vaultAPY.Points.Inception)
		if vaultAPY.ForwardAPR.Type == "" {
			checkDiff("Gross/Forward APR ", internalAPY.GrossAPR, vaultAPY.GrossAPR)
		} else {
			checkDiff("Gross/Forward APR ", internalAPY.GrossAPR, vaultAPY.ForwardAPR.GrossAPR)
		}
		spew.Printf("\n")
	}

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
