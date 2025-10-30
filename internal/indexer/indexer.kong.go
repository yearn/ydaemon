package indexer

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/kong"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

func IndexNewVaults(chainID uint64) map[common.Address]models.TVaultsFromRegistry {
	logs.Info(chainID, `-`, `Fetching all vaults from Kong GraphQL API (single source of truth)`)
	
	kongVaultData, err := kong.FetchVaultsFromKong(chainID)
	if err != nil {
		logs.Error(chainID, `-`, `CRITICAL: Failed to fetch vaults from Kong: %v`, err)
		logs.Error(chainID, `-`, `Cannot start yDaemon without Kong data - failing fast`)
		panic(fmt.Sprintf("Kong GraphQL API unavailable for chain %d: %v", chainID, err))
	}
	
	vaultsFromKong := make(map[common.Address]models.TVaultsFromRegistry)
	
	for vaultAddr, data := range kongVaultData {
		// Create vault entry with Kong metadata
		// Type and Kind will be populated by CMS metadata refresh
		vault := models.TVaultsFromRegistry{
			Address:         vaultAddr,
			RegistryAddress: data.Vault.GetRegistry(),
			TokenAddress:    data.Vault.GetAssetAddress(), // From Kong asset field
			Type:            models.TokenTypeStandardVault, // Default, overridden by CMS
			Kind:            models.VaultKindMultiple,      // Default, overridden by CMS
			APIVersion:      data.Vault.APIVersion,         // From Kong apiVersion field
			ChainID:         chainID,
			BlockNumber:     data.Vault.GetBlockNumber(),
			ExtraProperties: models.TExtraProperties{},
		}
		
		vaultsFromKong[vaultAddr] = vault
		storage.StoreNewVaultToRegistry(chainID, vault)

		// Store Kong APY and fee data from GraphQL API (single source of truth)
		kongSchema := models.TKongVaultSchema{
			APY:            data.APY,
			ManagementFee:  data.Vault.GetManagementFee(),
			PerformanceFee: data.Vault.GetPerformanceFee(),
		}
		storage.StoreKongVaultData(chainID, vaultAddr, kongSchema)
	}
	
	logs.Success(chainID, `-`, `Indexed %d vaults from Kong (complete replacement)`, len(vaultsFromKong))
	return vaultsFromKong
}

func IndexNewStrategies(chainID uint64, vaultMap map[common.Address]models.TVault) map[string]models.TStrategy {
	logs.Info(chainID, `-`, `Fetching all strategies from Kong GraphQL API (single source of truth)`)
	
	kongVaultData, err := kong.FetchVaultsFromKong(chainID)
	if err != nil {
		logs.Error(chainID, `-`, `CRITICAL: Failed to fetch strategies from Kong: %v`, err)
		logs.Error(chainID, `-`, `Cannot start yDaemon without Kong data - failing fast`)
		panic(fmt.Sprintf("Kong GraphQL API unavailable for chain %d: %v", chainID, err))
	}
	
	strategiesMap := make(map[string]models.TStrategy)
	totalStrategies := 0
	
	for vaultAddr, data := range kongVaultData {
		vault, exists := vaultMap[vaultAddr]
		if !exists {
			// If vault doesn't exist in our map, skip its strategies
			continue
		}
		
		for _, strategyAddr := range data.Strategies {
			if addresses.Equals(strategyAddr, common.Address{}) {
				continue
			}
			
			strategy := models.TStrategy{
				Address:      strategyAddr,
				ChainID:      chainID,
				VaultVersion: vault.Version,
				VaultAddress: vaultAddr,
				Activation:   vault.Activation,
			}
			
			// Use combination of strategy and vault address as key to handle 
			// strategies that may be used by multiple vaults
			key := strategyAddr.Hex() + "_" + vaultAddr.Hex()
			strategiesMap[key] = strategy
			totalStrategies++
		}
	}
	
	logs.Success(chainID, `-`, `Indexed %d strategies from Kong (complete replacement)`, totalStrategies)
	return strategiesMap
}

