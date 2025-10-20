package apr

import (
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

func computeCurrentV3VaultAPY(
	vault models.TVault,
) TVaultAPY {
	chainID := vault.ChainID

	/**********************************************************************************************
	** Retrieve the vault performance fee and management fee.
	** This can change from one vault to another and will be used in the final APR calculations.
	**********************************************************************************************/
	vaultPerformanceFee := helpers.ToNormalizedAmount(bigNumber.NewInt(int64(vault.PerformanceFee)), 4)
	vaultManagementFee := helpers.ToNormalizedAmount(bigNumber.NewInt(int64(vault.ManagementFee)), 4)

	/**********************************************************************************************
	** Fetch Kong APY data (single source of truth)
	** Kong provides pre-calculated historical APY values - no need for PPS fetching
	**********************************************************************************************/
	kongAPY, ok := storage.GetKongAPY(chainID, vault.Address)

	if !ok {
		logs.Error("CRITICAL: Kong APY missing for vault %s on chain %d - data not found. Check Kong data source.", vault.Address.Hex(), chainID)

		return TVaultAPY{
			Type:   `v3:kong_missing`,
			NetAPY: bigNumber.NewFloat(0),
			Fees: TFees{
				Performance: vaultPerformanceFee,
				Management:  vaultManagementFee,
			},
			Points: THistoricalPoints{
				WeekAgo:   bigNumber.NewFloat(0),
				MonthAgo:  bigNumber.NewFloat(0),
				Inception: bigNumber.NewFloat(0),
			},
			PricePerShare: TPricePerShare{
				Today:    bigNumber.NewFloat(0),
				WeekAgo:  bigNumber.NewFloat(0),
				MonthAgo: bigNumber.NewFloat(0),
			},
		}
	}

	/**********************************************************************************************
	** Parse Kong APY values (floats for APY, strings for PPS)
	** Kong returns null for vaults without sufficient history
	**********************************************************************************************/
	monthlyAPY := parseKongFloatAPY(kongAPY.MonthlyNet)
	weeklyAPY := parseKongFloatAPY(kongAPY.WeeklyNet)
	inceptionAPY := parseKongFloatAPY(kongAPY.InceptionNet)
	ppsToday := parseKongStringPPS(kongAPY.PricePerShare)
	ppsWeekAgo := parseKongStringPPS(kongAPY.WeeklyPricePerShare)
	ppsMonthAgo := parseKongStringPPS(kongAPY.MonthlyPricePerShare)

	/**********************************************************************************************
	** Determine APY type based on vault age
	** v3:new_averaged for vaults less than a week old, v3:averaged for others
	**********************************************************************************************/
	vaultAPRType := `v3:averaged`
	if vault.Activation > kongAPY.BlockNumber {
		vaultAPRType = `v3:new_averaged`
	}

	return TVaultAPY{
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
}
