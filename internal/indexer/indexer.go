package indexer

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/fetcher"
	"github.com/yearn/ydaemon/internal/models"
)

// ChainID is the chain ID of the chain to process
// vaultsList is the list of vaults to process
// Method is either append or replace. Append will only look for the strats for the provided vaultsList. Replace will look for all the strats for the provided chainID.
func ProcessNewVault(chainID uint64, vaultsList map[common.Address]models.TVaultsFromRegistry, method fetcher.TProcessNewVaultMethod) (map[common.Address]models.TVault, map[string]models.TStrategy) {
	vaultMap := fetcher.RetrieveAllVaults(chainID, vaultsList, method)
	strategiesMap := IndexNewStrategies(chainID, vaultMap)
	fetcher.RetrieveAllStrategies(chainID, strategiesMap)
	return vaultMap, strategiesMap
}
