package indexer

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/prices"
	"github.com/yearn/ydaemon/internal/strategies"
	"github.com/yearn/ydaemon/internal/tokens"
	"github.com/yearn/ydaemon/internal/utils"
	"github.com/yearn/ydaemon/internal/vaults"
)

func ProcessNewVault(chainID uint64, vaultsList map[common.Address]utils.TVaultsFromRegistry) {
	tokens.RetrieveAllTokens(chainID, vaultsList)
	prices.RetrieveAllPrices(chainID)
	vaults.RetrieveAllVaults(chainID, vaultsList)
	strategiesAddedList := strategies.RetrieveAllStrategiesAdded(chainID, vaultsList)
	strategies.RetrieveAllStrategies(chainID, strategiesAddedList)
}
