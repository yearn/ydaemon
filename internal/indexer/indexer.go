package indexer

import (
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/types/common"
	"github.com/yearn/ydaemon/internal/prices"
	"github.com/yearn/ydaemon/internal/strategies"
	"github.com/yearn/ydaemon/internal/tokens"
	"github.com/yearn/ydaemon/internal/utils"
	"github.com/yearn/ydaemon/internal/vaults"
)

func PostProcessStrategies(chainID uint64) {
	// Overwrite the st-yCRV strategy to set debtRatio to 100% and totalDebt to vault supply
	if chainID == 1 {
		styCRVStrategy, ok := strategies.FindStrategy(chainID, ethcommon.HexToAddress(`0xE7863292dd8eE5d215eC6D75ac00911D06E59B2d`))
		if ok {
			styCRVVault, ok := vaults.FindVault(chainID, common.FromAddress(styCRVStrategy.VaultAddress))
			if ok {
				styCRVStrategy.DebtRatio = bigNumber.NewUint64(10000)
				styCRVStrategy.TotalDebt = styCRVVault.TotalAssets
			}
		}
	}
}

func ProcessNewVault(chainID uint64, vaultsList map[ethcommon.Address]utils.TVaultsFromRegistry) {
	tokens.RetrieveAllTokens(chainID, vaultsList)
	prices.RetrieveAllPrices(chainID)
	vaults.RetrieveAllVaults(chainID, vaultsList)
	strategiesAddedList := strategies.RetrieveAllStrategiesAdded(chainID, vaultsList)
	strategies.RetrieveAllStrategies(chainID, strategiesAddedList)
	PostProcessStrategies(chainID)
}
