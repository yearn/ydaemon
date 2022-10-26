package strategies

import (
	"context"
	"strconv"
	"sync"

	"github.com/machinebox/graphql"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/models"
	"github.com/yearn/ydaemon/common/store"
	"github.com/yearn/ydaemon/common/types/common"
)

// FetchStrategiesList is an utility function that will query the subgraph in order to
// extract the list of strategies used by the Yearn system in order to be able to get
// complementary information about the strategies, directly from the contracts
func FetchStrategiesList(chainID uint64) {
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

	if Store.AggregatedStrategies[chainID] == nil {
		Store.AggregatedStrategies[chainID] = make(map[common.Address]*TAggregatedStrategy)
	}

	for _, vault := range response.Vaults {
		for _, strat := range vault.Strategies {
			Store.AggregatedStrategies[chainID][strat.Address] = &TAggregatedStrategy{
				Strategy:     strat.Address,
				Vault:        vault.Id,
				VaultVersion: vault.ApiVersion,
				Name:         strat.Name,
			}
			go store.SaveInDB(
				chainID,
				store.TABLES.STRATEGIES,
				strat.Address.String(),
				Store.AggregatedStrategies[chainID][strat.Address],
			)
		}
	}
}

// LoadAggregatedStrategies will reload the all the aggregatedStrategies from the DB for a given chainID
func LoadAggregatedStrategies(chainID uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	temp := make(map[common.Address]*TAggregatedStrategy)
	store.Interate(chainID, store.TABLES.STRATEGIES, &temp)
	if temp != nil && (len(temp) > 0) {
		Store.AggregatedStrategies[chainID] = temp
		logs.Success("Data loaded for the AggregatedStrategies store for chainID: " + strconv.FormatUint(chainID, 10))
	}
}
