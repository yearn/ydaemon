package indexer

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/fetcher"
	"github.com/yearn/ydaemon/internal/models"
)

// ChainID is the chain ID of the chain to process
// vaultsList is the list of vaults to process
// strategiesMapOrNil is an optional pre-fetched strategies map to avoid duplicate Kong API calls (pass nil to fetch)
// Method is either append or replace. Append will only look for the strats for the provided vaultsList. Replace will look for all the strats for the provided chainID.
func ProcessNewVault(chainID uint64, vaultsList map[common.Address]models.TVaultsFromRegistry, strategiesMapOrNil map[string]models.TStrategy, method fetcher.TProcessNewVaultMethod) (map[common.Address]models.TVault, map[string]models.TStrategy) {
	vaultMap := fetcher.RetrieveAllVaults(chainID, vaultsList, method)

	// If strategiesMap was provided (Kong-based flow), use it; otherwise fetch from Kong
	var strategiesMap map[string]models.TStrategy
	if strategiesMapOrNil != nil && len(strategiesMapOrNil) > 0 {
		strategiesMap = strategiesMapOrNil
	} else {
		// Legacy flow or empty strategies: fetch from Kong
		strategiesMap = IndexNewStrategies(chainID, vaultMap)
	}

	fetcher.RetrieveAllStrategies(chainID, strategiesMap)
	return vaultMap, strategiesMap
}
