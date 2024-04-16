package apr

import (
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/fetcher"
	"github.com/yearn/ydaemon/internal/indexer"
	"github.com/yearn/ydaemon/processes/prices"
)

/**************************************************************************************************
** initYearnEcosystem is used to initialize the yearn ecosystem data, aka fetching all the vaults,
** strategies, tokens, prices, etc.
** Based on that, we have everything ready to compute the fees for each partner.
**************************************************************************************************/
func initYearnEcosystem(chainID uint64) {
	registries := indexer.IndexNewVaults(chainID)
	vaultMap := fetcher.RetrieveAllVaults(chainID, registries)
	strategiesMap := indexer.IndexNewStrategies(chainID, vaultMap)
	tokenMap := fetcher.RetrieveAllTokens(chainID, vaultMap)

	prices.RetrieveAllPrices(chainID, tokenMap)
	logs.Info(`loading staking pools...`)
	indexer.IndexStakingPools(chainID)
	indexer.IndexVeYFIStakingContract(chainID)
	logs.Info(`loading strategies...`)
	fetcher.RetrieveAllStrategies(chainID, strategiesMap)
}
