package indexer

import (
	"math/big"

	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

func handleV2Vault(chainID uint64, value *contracts.YRegistryV2NewVault) models.TVaultsFromRegistry {
	newVault := models.TVaultsFromRegistry{
		ChainID:         chainID,
		RegistryAddress: value.Raw.Address,
		Address:         value.Vault,
		TokenAddress:    value.Token,
		APIVersion:      value.ApiVersion,
		BlockNumber:     value.Raw.BlockNumber,
		Type:            models.TokenTypeStandardVault,
	}
	return newVault
}

func handleV3Vault(chainID uint64, value *contracts.YRegistryV3NewVault) models.TVaultsFromRegistry {
	newVault := models.TVaultsFromRegistry{
		ChainID:         chainID,
		RegistryAddress: value.Raw.Address,
		Address:         value.Vault,
		TokenAddress:    value.Token,
		APIVersion:      value.ApiVersion,
		BlockNumber:     value.Raw.BlockNumber,
		Type:            models.TokenTypeStandardVault,
	}
	if value.VaultType.Cmp(big.NewInt(2)) == 0 {
		newVault.Type = models.TokenTypeAutomatedVault
	}
	return newVault
}

func handleV4Vault(chainID uint64, value *contracts.YRegistryV4NewEndorsedVault) models.TVaultsFromRegistry {
	newVault := models.TVaultsFromRegistry{
		ChainID:         chainID,
		RegistryAddress: value.Raw.Address,
		Address:         value.Vault,
		TokenAddress:    value.Asset,
		APIVersion:      value.ReleaseVersion.Text(10),
		BlockNumber:     value.Raw.BlockNumber,
		Type:            models.TokenTypeStandardVault,
	}
	logs.Info(`Got V4 vault ` + value.Vault.Hex() + ` from registry ` + value.Raw.Address.Hex())
	return newVault
}

func handleV2Strategies(chainID uint64, vaultVersion string, value *contracts.Yvault022StrategyAdded) models.TStrategyAdded {
	newStrategy := models.TStrategyAdded{
		Address:      value.Strategy,
		ChainID:      chainID,
		VaultVersion: vaultVersion,
		VaultAddress: value.Raw.Address,
		BlockNumber:  value.Raw.BlockNumber,
	}
	return newStrategy
}

func handleV3Strategies(chainID uint64, vaultVersion string, value *contracts.Yvault030StrategyAdded) models.TStrategyAdded {
	newStrategy := models.TStrategyAdded{
		Address:      value.Strategy,
		ChainID:      chainID,
		VaultVersion: vaultVersion,
		VaultAddress: value.Raw.Address,
		BlockNumber:  value.Raw.BlockNumber,
	}
	return newStrategy
}

func handleV3StrategiesMigration(chainID uint64, value *contracts.Yvault030StrategyMigrated) models.TStrategyMigrated {
	newStrategy := models.TStrategyMigrated{
		ChainID:            chainID,
		VaultAddress:       value.Raw.Address,
		OldStrategyAddress: value.OldVersion,
		NewStrategyAddress: value.NewVersion,
		BlockNumber:        value.Raw.BlockNumber,
	}
	return newStrategy
}

func handleV4Strategies(chainID uint64, vaultVersion string, value *contracts.Yvault043StrategyAdded) models.TStrategyAdded {
	newStrategy := models.TStrategyAdded{
		Address:      value.Strategy,
		ChainID:      chainID,
		VaultVersion: vaultVersion,
		VaultAddress: value.Raw.Address,
		BlockNumber:  value.Raw.BlockNumber,
	}
	return newStrategy
}

func handleV4StrategiesMigration(chainID uint64, value *contracts.Yvault043StrategyMigrated) models.TStrategyMigrated {
	newStrategy := models.TStrategyMigrated{
		ChainID:            chainID,
		VaultAddress:       value.Raw.Address,
		OldStrategyAddress: value.OldVersion,
		NewStrategyAddress: value.NewVersion,
		BlockNumber:        value.Raw.BlockNumber,
	}
	return newStrategy
}

func processMigrations(chainID uint64) {
	limit := uint64(1000)
	migratedHistoricalStrategies, _ := storage.ListStrategiesMigrated(chainID)
	shouldLoop := true
	for shouldLoop {
		limit++
		shouldLoop = false
		for _, newStrategy := range migratedHistoricalStrategies {
			if _, ok := storage.GetStrategy(chainID, newStrategy.NewStrategyAddress); ok {
				continue
			}
			if oldStrategy, ok := storage.GetStrategy(chainID, newStrategy.OldStrategyAddress); ok {
				migratedStrategy := models.TStrategyAdded{
					Address:      newStrategy.NewStrategyAddress,
					VaultAddress: newStrategy.VaultAddress,
					VaultVersion: oldStrategy.VaultVersion,
					BlockNumber:  newStrategy.BlockNumber,
					ChainID:      chainID,
				}
				storage.StoreStrategy(chainID, migratedStrategy)
			} else {
				shouldLoop = true
			}
		}

		if limit > 1000 {
			break
		}
	}
}
