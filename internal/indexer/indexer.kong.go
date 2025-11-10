package indexer

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/kong"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

func IndexNewVaults(chainID uint64) map[common.Address]models.TVaultsFromRegistry {
	logs.Info(chainID, `-`, `Fetching all vaults and strategies from Kong GraphQL API (single source of truth)`)

	// Fetch strategies first so we can pass them to FetchVaultsFromKong
	strategiesByVault, err := kong.FetchStrategiesFromKong(chainID)
	if err != nil {
		logs.Error(chainID, `-`, `CRITICAL: Failed to fetch strategies from Kong: %v`, err)
		logs.Error(chainID, `-`, `Cannot start yDaemon without Kong data - failing fast`)
		panic(fmt.Sprintf("Kong GraphQL API unavailable for chain %d: %v", chainID, err))
	}

	kongVaultData, err := kong.FetchVaultsFromKong(chainID, strategiesByVault)
	if err != nil {
		logs.Error(chainID, `-`, `CRITICAL: Failed to fetch vaults from Kong: %v`, err)
		logs.Error(chainID, `-`, `Cannot start yDaemon without Kong data - failing fast`)
		panic(fmt.Sprintf("Kong GraphQL API unavailable for chain %d: %v", chainID, err))
	}
	
	vaultsFromKong := make(map[common.Address]models.TVaultsFromRegistry)

	for vaultAddr, data := range kongVaultData {
		// Type and Kind will be populated by CMS metadata refresh
		vault := models.TVaultsFromRegistry{
			Address:         vaultAddr,
			RegistryAddress: data.Vault.GetRegistry(),
			TokenAddress:    data.Vault.GetAssetAddress(), // From Kong asset field
			Type:            models.TokenTypeStandardVault, // Default, overridden by CMS
			Kind:            models.VaultKindMultiple,      // Default, overridden by CMS
			APIVersion:      data.Vault.GetAPIVersion(),         // From Kong apiVersion field
			ChainID:         chainID,
			BlockNumber:     data.Vault.GetBlockNumber(),
			ExtraProperties: models.TExtraProperties{},
		}

		vaultsFromKong[vaultAddr] = vault
		storage.StoreNewVaultToRegistry(chainID, vault)

		// Store Kong APY data from GraphQL API (single source of truth for APY calculations)
		// Convert KongDebt to TKongDebt
		var debts []models.TKongDebt
		for _, debt := range data.Debts {
			debts = append(debts, models.TKongDebt{
				Strategy:           debt.Strategy,
				PerformanceFee:     debt.PerformanceFee,
				Activation:         debt.Activation,
				DebtRatio:          debt.DebtRatio,
				MinDebtPerHarvest:  debt.MinDebtPerHarvest,
				MaxDebtPerHarvest:  debt.MaxDebtPerHarvest,
				LastReport:         debt.LastReport,
				TotalDebt:          debt.TotalDebt,
				TotalDebtUsd:       debt.TotalDebtUsd,
				TotalGain:          debt.TotalGain,
				TotalGainUsd:        debt.TotalGainUsd,
				TotalLoss:          debt.TotalLoss,
				TotalLossUsd:       debt.TotalLossUsd,
				CurrentDebt:        debt.CurrentDebt,
				CurrentDebtUsd:     debt.CurrentDebtUsd,
				MaxDebt:            debt.MaxDebt,
				MaxDebtUsd:         debt.MaxDebtUsd,
				TargetDebtRatio:    debt.TargetDebtRatio,
				MaxDebtRatio:       debt.MaxDebtRatio,
			})
		}
		
		kongSchema := models.TKongVaultSchema{
			ManagementFee:  data.Vault.GetManagementFee(),
			PerformanceFee: data.Vault.GetPerformanceFee(),
			APY: data.APY,
			Debts: debts,
			TVL: data.TVL,
			Strategies: data.Strategies,
			TotalAssets: data.TotalAssets,
		}
		storage.StoreKongVaultData(chainID, vaultAddr, kongSchema)
	}
	
	logs.Success(chainID, `-`, `Indexed %d vaults from Kong (complete replacement)`, len(vaultsFromKong))
	return vaultsFromKong
}

func IndexNewStrategies(chainID uint64, vaultMap map[common.Address]models.TVault) map[string]models.TStrategy {
	logs.Info(chainID, `-`, `Fetching all strategies from Kong GraphQL API (single source of truth)`)

	strategiesByVault, err := kong.FetchStrategiesFromKong(chainID)
	if err != nil {
		logs.Error(chainID, `-`, `CRITICAL: Failed to fetch strategies from Kong: %v`, err)
		logs.Error(chainID, `-`, `Cannot start yDaemon without Kong data - failing fast`)
		panic(fmt.Sprintf("Kong GraphQL API unavailable for chain %d: %v", chainID, err))
	}

	strategiesMap := make(map[string]models.TStrategy)
	totalStrategies := 0

	// Iterate through strategies mapped by vault
	for vaultAddr, kongStrategies := range strategiesByVault {
		for _, kongStrategy := range kongStrategies {
			strategyAddr := kongStrategy.GetAddress()

			// Store Kong strategy data directly (single source of truth)
			storage.StoreKongStrategyData(chainID, strategyAddr, vaultAddr, kongStrategy)

			// Use combination of strategy and vault address as key to handle
			// strategies that may be used by multiple vaults
			key := strategyAddr.Hex() + "_" + vaultAddr.Hex()
			strategiesMap[key] = kongStrategy.ToTStrategy()
			totalStrategies++
		}
	}

	logs.Success(chainID, `-`, `Indexed %d strategies from Kong (complete replacement)`, totalStrategies)
	return strategiesMap
}

