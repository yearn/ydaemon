package indexer

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/strategies"
	"github.com/yearn/ydaemon/internal/vaults"
)

func PostProcess(chainID uint64) {
	// Overwrite the st-yCRV strategy to set debtRatio to 100% and totalDebt to vault supply
	if chainID == 1 {
		styCRVStrategy, ok := strategies.FindStrategy(chainID, common.HexToAddress(`0xE7863292dd8eE5d215eC6D75ac00911D06E59B2d`))
		if ok {
			styCRVVault, ok := vaults.FindVault(chainID, styCRVStrategy.VaultAddress)
			if ok {
				styCRVStrategy.DebtRatio = bigNumber.NewUint64(10000)
				styCRVStrategy.TotalDebt = styCRVVault.LastTotalAssets
			}
		}
	}
}

func processNewVault(chainID uint64, vaultsList map[common.Address]models.TVaultsFromRegistry) {
	registries := IndexNewVaults(chainID)
	vaultMap := vaults.RetrieveAllVaults(chainID, registries)
	strategiesSlice := IndexNewStrategies(chainID, vaultMap)
	strategies.RetrieveAllStrategies(chainID, strategiesSlice)
	PostProcess(chainID)
}
