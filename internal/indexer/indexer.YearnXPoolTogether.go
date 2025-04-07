package indexer

import (
	"math/big"

	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

func IndexYearnXPoolTogetherVaults(chainID uint64) (vaults []models.TVaultsFromRegistry) {
	chain, ok := env.GetChain(chainID)
	if !ok {
		return vaults
	}
	if len(chain.YearnXRegistries) == 0 {
		return vaults
	}

	var registryToUse env.TContractData
	for _, registry := range chain.YearnXRegistries {
		if registry.Label == `POOL_TOGETHER` {
			registryToUse = registry
			break
		}
	}

	/** 🔵 - Yearn *********************************************************************************
	** Loop over all the known vaults for the specified chain ID.
	**********************************************************************************************/
	client := ethereum.GetRPC(chainID)
	registry, err := contracts.NewPoolTogetherRegistry(registryToUse.Address, client)
	if err != nil {
		return vaults
	}
	totalVaults, err := registry.TotalVaults(nil)
	if err != nil {
		return vaults
	}

	for i := uint64(0); i < totalVaults.Uint64(); i++ {
		vaultAddress, err := registry.AllVaults(nil, big.NewInt(int64(i)))
		if err != nil {
			continue
		}
		if addresses.Equals(vaultAddress, `0x8aB157b779C72e2348364b5F8148cC45f63a8724`) {
			continue
		}

		// if _, ok := storage.GetVaultFromRegistry(chainID, vaultAddress); ok {
		// 	continue
		// }
		vaultContract, err := contracts.NewPrizeVault(vaultAddress, client)
		if err != nil {
			continue
		}
		vaultAsset, _ := vaultContract.Asset(nil)
		yieldVault, err := vaultContract.YieldVault(nil)
		poolTogetherVault := models.TVaultsFromRegistry{
			ChainID:         chainID,
			Address:         vaultAddress,
			RegistryAddress: registryToUse.Address,
			TokenAddress:    vaultAsset,
			APIVersion:      `~3.0.2`,
			BlockNumber:     registryToUse.Block,
			Type:            models.TTokenType(models.VaultKindSingle),
			Kind:            models.VaultKindSingle,
			ExtraProperties: models.TExtraProperties{
				YieldVaultAddress: yieldVault.String(),
			},
		}
		vaults = append(vaults, poolTogetherVault)
		storage.StoreNewVaultToRegistry(chainID, poolTogetherVault)
	}

	return vaults
}
