package apr

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

func computeCurrentV2VaultAPY(
	vault models.TVault,
	vaultToken models.TERC20Token,
) TVaultAPY {
	chainID := vault.ChainID
	yieldVault := vault.Address
	registry, found := storage.GetVaultFromRegistry(chainID, vault.Address)
	if found && registry.ExtraProperties.YieldVaultAddress != `` {
		yieldVault = common.HexToAddress(registry.ExtraProperties.YieldVaultAddress)
	}

	ppsInception := bigNumber.NewFloat(1)
	ppsToday := ethereum.FetchPPSToday(chainID, yieldVault, vault.Activation, vaultToken.Decimals)
	ppsWeekAgo := ethereum.FetchPPSLastWeek(chainID, yieldVault, vault.Activation, vaultToken.Decimals)
	ppsMonthAgo := ethereum.FetchPPSLastMonth(chainID, yieldVault, vault.Activation, vaultToken.Decimals)

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
	vaultAPRType := `v2:averaged`
	monthAgoBlockNumber := ethereum.GetBlockNumberByPeriod(chainID, 30)
	if vault.Activation > monthAgoBlockNumber {
		vaultAPRType = `v2:new_averaged`
	}

	vaultAPR := TVaultAPY{
		Type:   vaultAPRType,
		NetAPY: ethereum.CalculateMonthlyAPY(ppsToday, ppsMonthAgo),
		Fees: TFees{
			Performance: vaultPerformanceFee,
			Management:  vaultManagementFee,
		},
		Points: THistoricalPoints{
			WeekAgo:   ethereum.CalculateWeeklyAPY(ppsToday, ppsWeekAgo),
			MonthAgo:  ethereum.CalculateMonthlyAPY(ppsToday, ppsMonthAgo),
			Inception: ethereum.CalculateYearlyAPY(ppsToday, ppsInception),
		},
		PricePerShare: TPricePerShare{
			Today:    ppsToday,
			WeekAgo:  ppsWeekAgo,
			MonthAgo: ppsMonthAgo,
		},
	}
	return vaultAPR
}
