package indexer

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/fetcher"
	"github.com/yearn/ydaemon/internal/models"
)

func processNewVault(chainID uint64, vaultsList map[common.Address]models.TVaultsFromRegistry) {
	vaultMap := fetcher.RetrieveAllVaults(chainID, vaultsList)
	strategiesMap := IndexNewStrategies(chainID, vaultMap)
	fetcher.RetrieveAllStrategies(chainID, strategiesMap)
}
