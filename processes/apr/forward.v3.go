package apr

import (
	"strings"

	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

func isV3Vault(vault models.TVault) bool {
	versionMajor := strings.Split(vault.Version, `.`)[0]
	return vault.Kind == models.VaultKindMultiple || vault.Kind == models.VaultKindSingle || versionMajor == `3` || versionMajor == `~3`
}

func computeVaultV3ForwardAPY(
	vault models.TVault,
	allStrategiesForVault map[string]models.TStrategy,
) TForwardAPY {
	chainID := vault.ChainID

	/**********************************************************************************************
	** Fetch Kong oracle APR/APY data (single source of truth)
	** Kong provides pre-calculated oracle APR and APY from daily timeseries hook
	**********************************************************************************************/
	oracleAPR, oracleAPY, ok := storage.GetKongOracleAPY(chainID, vault.Address)

	if !ok {
		logs.Error("Kong oracle data missing for vault %s on chain %d - Kong source unavailable", vault.Address.Hex(), chainID)
		return TForwardAPY{
			Type:   `v3:onchainOracle`,
			NetAPY: bigNumber.NewFloat(0),
			Composite: TCompositeData{
				V3OracleCurrentAPR:    bigNumber.NewFloat(0),
				V3OracleStratRatioAPR: bigNumber.NewFloat(0),
			},
		}
	}

	// Handle nil values (Kong returns null for vaults without oracle data)
	if oracleAPY == nil {
		logs.Warning("Kong oracle APY is null for vault %s on chain %d - vault may not have oracle configured", vault.Address.Hex(), chainID)
		return TForwardAPY{
			Type:   `v3:onchainOracle`,
			NetAPY: bigNumber.NewFloat(0),
			Composite: TCompositeData{
				V3OracleCurrentAPR:    bigNumber.NewFloat(0),
				V3OracleStratRatioAPR: bigNumber.NewFloat(0),
			},
		}
	}

	/**********************************************************************************************
	** Use Kong's pre-calculated oracle APY
	** Kong already converts APR to APY using weekly compounding (52 periods/year)
	**********************************************************************************************/
	primaryAPY := bigNumber.NewFloat(*oracleAPY)
	primaryAPR := bigNumber.NewFloat(0)
	if oracleAPR != nil {
		primaryAPR = bigNumber.NewFloat(*oracleAPR)
	}

	return TForwardAPY{
		Type:   `v3:onchainOracle`,
		NetAPY: primaryAPY,
		Composite: TCompositeData{
			V3OracleCurrentAPR:    primaryAPR,
			V3OracleStratRatioAPR: bigNumber.NewFloat(0), // Not used with Kong oracle
		},
	}
}
