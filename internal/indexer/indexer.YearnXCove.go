package indexer

import (
	"math/big"

	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

func IndexYearnXCoveVaults(chainID uint64) (vaults []models.TVaultsFromRegistry) {
	chain, ok := env.GetChain(chainID)
	if !ok {
		return vaults
	}
	if len(chain.YearnXRegistries) == 0 {
		return vaults
	}

	var registryToUse env.TContractData
	for _, registry := range chain.YearnXRegistries {
		if registry.Label == `COVE` {
			registryToUse = registry
			break
		}
	}

	/** 🔵 - Yearn *********************************************************************************
	** Loop over all the known vaults for the specified chain ID.
	**********************************************************************************************/
	client := ethereum.GetRPC(chainID)
	registry, err := contracts.NewCoveRegistry(registryToUse.Address, client)
	if err != nil {
		return vaults
	}
	totalVaults, err := registry.NumOfSupportedYearnGauges(nil)
	if err != nil {
		return vaults
	}

	allGaugesVaults, err := registry.GetAllGaugeInfo(nil, big.NewInt(int64(totalVaults.Uint64())), big.NewInt(0))
	if err != nil {
		return vaults
	}

	for _, gauge := range allGaugesVaults {
		vaultAddress := gauge.CoveYearnStrategy
		if _, ok := storage.GetVaultFromRegistry(chainID, vaultAddress); !ok {

			coveVault := models.TVaultsFromRegistry{
				ChainID:         chainID,
				Address:         vaultAddress,
				RegistryAddress: registryToUse.Address,
				TokenAddress:    gauge.YearnGauge,
				APIVersion:      `3.0.2`,
				BlockNumber:     registryToUse.Block,
				Type:            models.TTokenType(models.VaultKindSingle),
				Kind:            models.VaultKindSingle,
				ExtraProperties: models.TExtraProperties{
					YearnVaultAsset:         gauge.YearnVaultAsset.String(),
					YearnVault:              gauge.YearnGauge.String(),
					IsVaultV2:               gauge.IsVaultV2,
					YearnGauge:              gauge.YearnGauge.String(),
					CoveYearnStrategy:       gauge.CoveYearnStrategy.String(),
					AutoCompoundingGauge:    gauge.AutoCompoundingGauge.String(),
					NonAutoCompoundingGauge: gauge.NonAutoCompoundingGauge.String(),
				},
			}
			vaults = append(vaults, coveVault)
			storage.StoreNewVaultToRegistry(chainID, coveVault)
		}

		autoCompoundingVaultAddress := gauge.AutoCompoundingGauge
		if _, ok := storage.GetVaultFromRegistry(chainID, autoCompoundingVaultAddress); !ok {

			coveVault := models.TVaultsFromRegistry{
				ChainID:         chainID,
				Address:         autoCompoundingVaultAddress,
				RegistryAddress: registryToUse.Address,
				TokenAddress:    gauge.YearnGauge,
				APIVersion:      `3.0.2`,
				BlockNumber:     registryToUse.Block,
				Type:            models.TTokenType(models.VaultKindSingle),
				Kind:            models.VaultKindSingle,
				ExtraProperties: models.TExtraProperties{
					YearnVaultAsset:         gauge.YearnVaultAsset.String(),
					YearnVault:              gauge.YearnGauge.String(),
					IsVaultV2:               gauge.IsVaultV2,
					YearnGauge:              gauge.YearnGauge.String(),
					CoveYearnStrategy:       gauge.CoveYearnStrategy.String(),
					AutoCompoundingGauge:    gauge.AutoCompoundingGauge.String(),
					NonAutoCompoundingGauge: gauge.NonAutoCompoundingGauge.String(),
				},
			}
			vaults = append(vaults, coveVault)
			storage.StoreNewVaultToRegistry(chainID, coveVault)
		}

		nonAutoCompoundingVaultAddress := gauge.NonAutoCompoundingGauge
		if _, ok := storage.GetVaultFromRegistry(chainID, nonAutoCompoundingVaultAddress); !ok {
			coveVault := models.TVaultsFromRegistry{
				ChainID:         chainID,
				Address:         nonAutoCompoundingVaultAddress,
				RegistryAddress: registryToUse.Address,
				TokenAddress:    gauge.YearnGauge,
				APIVersion:      `3.0.2`,
				BlockNumber:     registryToUse.Block,
				Type:            models.TTokenType(models.VaultKindSingle),
				Kind:            models.VaultKindSingle,
				ExtraProperties: models.TExtraProperties{
					YearnVaultAsset:         gauge.YearnVaultAsset.String(),
					YearnVault:              gauge.YearnGauge.String(),
					IsVaultV2:               gauge.IsVaultV2,
					YearnGauge:              gauge.YearnGauge.String(),
					CoveYearnStrategy:       gauge.CoveYearnStrategy.String(),
					AutoCompoundingGauge:    gauge.AutoCompoundingGauge.String(),
					NonAutoCompoundingGauge: gauge.NonAutoCompoundingGauge.String(),
				},
			}
			vaults = append(vaults, coveVault)
			storage.StoreNewVaultToRegistry(chainID, coveVault)
		}
	}

	return vaults
}
