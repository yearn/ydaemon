package apr

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

func computeCurrentV3VaultAPY(
	vault models.TVault,
	vaultToken models.TERC20Token,
) TVaultAPY {
	chainID := vault.ChainID
	yieldVault := vault.Address
	registry, found := storage.GetVaultFromRegistry(chainID, vault.Address)
	if found && registry.ExtraProperties.YieldVaultAddress != `` {
		yieldVault = common.HexToAddress(registry.ExtraProperties.YieldVaultAddress)
	}

	estBlockToday := ethereum.GetBlockNumberByPeriod(chainID, 0)
	estBlockLastWeek := ethereum.GetBlockNumberByPeriod(chainID, 7)
	estBlockLastMonth := ethereum.GetBlockNumberByPeriod(chainID, 30)
	estBlockLastYear := ethereum.GetBlockNumberByPeriod(chainID, 365)
	blocksSinceDeployment := estBlockToday - vault.Activation
	ppsToday := ethereum.FetchPPSToday(chainID, yieldVault, vault.Activation, vaultToken.Decimals)
	ppsWeekAgo := bigNumber.NewFloat(1)
	ppsMonthAgo := bigNumber.NewFloat(1)
	ppsInception := bigNumber.NewFloat(1)
	weeklyAPY := bigNumber.NewFloat(0)
	monthlyAPY := bigNumber.NewFloat(0)
	inceptionAPY := bigNumber.NewFloat(0)

	isLessThanAWeekOld := vault.Activation > 0 && estBlockLastWeek < vault.Activation
	isLessThanAMonthOld := vault.Activation > 0 && estBlockLastMonth < vault.Activation

	/**************************************************************************
	** Switch function to gracefully handle APY calculations for vaults that
	** are less than a week or month old.
	** We calculate the number of days the vault has been active and use that
	** number to calculate the APY.
	** default case if for all vaults older than a month.
	**************************************************************************/
	switch {
	case isLessThanAWeekOld:
		// Calculate average blocks per day over the last 7 days
		numBlocksIn7Days := estBlockToday - estBlockLastWeek
		numBlocksPerDay := float64(numBlocksIn7Days) / 7
		daysSinceDeployment := float64(blocksSinceDeployment) / numBlocksPerDay
		if daysSinceDeployment < 1 {
			daysSinceDeployment = 1
		}
		weeklyAPY = ethereum.CalculateAPY(ppsToday, ppsWeekAgo, int(daysSinceDeployment))
		monthlyAPY = weeklyAPY
		inceptionAPY = monthlyAPY
	case isLessThanAMonthOld:
		ppsWeekAgo = ethereum.FetchPPSLastWeek(chainID, yieldVault, vault.Activation, vaultToken.Decimals)
		weeklyAPY = ethereum.CalculateWeeklyAPY(ppsToday, ppsWeekAgo)
		// Calculate average blocks per day over the last 30 days
		numBlocksIn30Days := estBlockToday - estBlockLastMonth
		numBlocksPerDay := float64(numBlocksIn30Days) / 30
		daysSinceDeployment := float64(blocksSinceDeployment) / numBlocksPerDay
		if daysSinceDeployment < 1 {
			daysSinceDeployment = 1
		}
		monthlyAPY = ethereum.CalculateAPY(ppsToday, ppsMonthAgo, int(daysSinceDeployment))
		inceptionAPY = monthlyAPY
	default:
		ppsWeekAgo = ethereum.FetchPPSLastWeek(chainID, yieldVault, vault.Activation, vaultToken.Decimals)
		weeklyAPY = ethereum.CalculateWeeklyAPY(ppsToday, ppsWeekAgo)
		ppsMonthAgo = ethereum.FetchPPSLastMonth(chainID, yieldVault, vault.Activation, vaultToken.Decimals)
		monthlyAPY = ethereum.CalculateMonthlyAPY(ppsToday, ppsMonthAgo)
		numBlocksIn365Days := estBlockToday - estBlockLastYear
		numBlocksPerDay := float64(numBlocksIn365Days) / 365
		daysSinceDeployment := float64(blocksSinceDeployment) / numBlocksPerDay
		inceptionAPY = ethereum.CalculateAPY(ppsToday, ppsInception, int(daysSinceDeployment))
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
	** As we now have the base APR information we can init our structure. This base structure MUST
	** contains:
	** - The type of the APR (always v2:averaged for now)
	** - The grossAPR
	** - The netAPY
	** - The fees (performance and management)
	** - The points (PPS evolution over time, for one week, one month and since inception)
	**********************************************************************************************/
	vaultAPRType := `v3:averaged`
	if vault.Activation > estBlockLastWeek {
		vaultAPRType = `v3:new_averaged`
	}

	vaultAPR := TVaultAPY{
		Type:   vaultAPRType,
		NetAPY: monthlyAPY,
		Fees: TFees{
			Performance: vaultPerformanceFee,
			Management:  vaultManagementFee,
		},
		Points: THistoricalPoints{
			WeekAgo:   weeklyAPY,
			MonthAgo:  monthlyAPY,
			Inception: inceptionAPY,
		},
		PricePerShare: TPricePerShare{
			Today:    ppsToday,
			WeekAgo:  ppsWeekAgo,
			MonthAgo: ppsMonthAgo,
		},
	}
	return vaultAPR
}
