package strategies

import (
	"context"
	"strconv"
	"sync"

	"github.com/machinebox/graphql"
	"github.com/yearn/ydaemon/internal/utils/env"
	"github.com/yearn/ydaemon/internal/utils/logs"
	"github.com/yearn/ydaemon/internal/utils/models"
	"github.com/yearn/ydaemon/internal/utils/store"
	"github.com/yearn/ydaemon/internal/utils/types/common"
)

// FetchStrategiesList is an utility function that will query the subgraph in order to
// extract the list of strategies used by the Yearn system in order to be able to get
// complementary information about the strategies, directly from the contracts
func FetchStrategiesList(chainID uint64) {
	strategyList := make(map[common.Address]models.TStrategyList)
	graphQLEndpoint, ok := env.THEGRAPH_ENDPOINTS[chainID]
	if !ok {
		logs.Error("No graph endpoint for chainID", chainID)
		return
	}
	client := graphql.NewClient(graphQLEndpoint)
	request := graphql.NewRequest(`
        {
			vaults(first: 1000) {
				id
				apiVersion
				strategies(first: 40) {
					address
					name
				}
			}
        }
    `)
	var response models.TGraphQueryResponseForVaults
	if err := client.Run(context.Background(), request, &response); err != nil {
		logs.Error(`Error fetching token list from the graph: `, err)
		return
	}

	for _, vault := range response.Vaults {
		for _, strat := range vault.Strategies {
			strategyList[strat.Address] = models.TStrategyList{
				Strategy:     strat.Address,
				Vault:        vault.Id,
				VaultVersion: vault.ApiVersion,
				Name:         strat.Name,
			}
		}
	}
	Store.StrategyList[chainID] = strategyList
	store.SaveInDBForChainID(store.KEYS.StrategyList, chainID, Store.StrategyList[chainID])
}

// LoadStrategyList will reload the strategyList data store from the last state of the local Badger Database
func LoadStrategyList(chainID uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	temp := make(map[common.Address]models.TStrategyList)
	if err := store.LoadFromDBForChainID(store.KEYS.StrategyList, chainID, &temp); err != nil {
		return
	}
	if temp != nil && (len(temp) > 0) {
		Store.StrategyList[chainID] = temp
		logs.Success("Data loaded for the strategyList store for chainID: " + strconv.FormatUint(chainID, 10))
	}
}
