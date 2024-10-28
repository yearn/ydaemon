package apr

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/store"
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

	ppsPerTime, _ := store.ListPricePerShare(chainID, yieldVault)
	ppsInception := bigNumber.NewFloat(1)
	ppsToday := helpers.GetPPSToday(ppsPerTime, vaultToken.Decimals)
	if ppsToday == nil || ppsToday.IsZero() {
		ppsToday = ethereum.FetchPPSToday(chainID, yieldVault, vaultToken.Decimals)
	}

	ppsWeekAgo := helpers.GetPPSLastWeek(ppsPerTime, vaultToken.Decimals)
	if ppsWeekAgo == nil || ppsWeekAgo.IsZero() {
		ppsWeekAgo = ethereum.FetchPPSLastWeek(chainID, yieldVault, vaultToken.Decimals)
	}

	ppsMonthAgo := helpers.GetPPSLastMonth(ppsPerTime, vaultToken.Decimals)
	if ppsMonthAgo == nil || ppsMonthAgo.IsZero() {
		ppsMonthAgo = ethereum.FetchPPSLastMonth(chainID, yieldVault, vaultToken.Decimals)
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
	twoWeeksAgoBlockNumber := ethereum.GetBlockNumberXDaysAgo(chainID, 14)
	if vault.Activation > twoWeeksAgoBlockNumber {
		vaultAPRType = `v3:new_averaged`
	}

	vaultAPR := TVaultAPY{
		Type:   vaultAPRType,
		NetAPY: helpers.GetEvolution(ppsToday, ppsMonthAgo, bigNumber.NewFloat(30)),
		Fees: TFees{
			Performance: vaultPerformanceFee,
			Management:  vaultManagementFee,
		},
		Points: THistoricalPoints{
			WeekAgo:   helpers.GetEvolution(ppsToday, ppsWeekAgo, bigNumber.NewFloat(7)),
			MonthAgo:  helpers.GetEvolution(ppsToday, ppsMonthAgo, bigNumber.NewFloat(30)),
			Inception: helpers.GetEvolution(ppsToday, ppsInception, bigNumber.NewFloat(365)),
		},
		PricePerShare: TPricePerShare{
			Today:    ppsToday,
			WeekAgo:  ppsWeekAgo,
			MonthAgo: ppsMonthAgo,
		},
	}
	return vaultAPR
}
