package daemons

import (
	"context"
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/machinebox/graphql"
	"github.com/majorfi/ydaemon/internal/ethereum"
	"github.com/majorfi/ydaemon/internal/logs"
	"github.com/majorfi/ydaemon/internal/models"
	"github.com/majorfi/ydaemon/internal/store"
)

// FetchStrategiesList is an utility function that will query the subgraph in order to
// extract the list of strategies used by the Yearn system in order to be able to get
// complementary information about the strategies, directly from the contracts
func FetchStrategiesList(chainID uint64) {
	strategyList := make(map[common.Address]models.TStrategyList)
	client := graphql.NewClient(ethereum.GetGraphURI(chainID))
	request := graphql.NewRequest(`
        {
			vaults(first: 1000) {
				id
				apiVersion
				strategies(first: 40) {address}
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
			strategyList[common.HexToAddress(strat.Address)] = models.TStrategyList{
				Strategy:     common.HexToAddress(strat.Address),
				Vault:        common.HexToAddress(vault.Id),
				VaultVersion: vault.ApiVersion,
			}
		}
	}
	store.StrategyList[chainID] = strategyList
	store.SaveInDBForChainID(`StrategyList`, chainID, store.StrategyList[chainID])
}

// LoadStrategyList will reload the strategyList data store from the last state of the local Badger Database
func LoadStrategyList(chainID uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	temp := make(map[common.Address]models.TStrategyList)
	err := store.LoadFromDBForChainID(`StrategyList`, chainID, &temp)
	if err != nil {
		if err.Error() == "Key not found" {
			logs.Warning("No StrategyList data found for chainID: " + strconv.FormatUint(chainID, 10))
			return
		}
		logs.Error(err)
		return
	}
	if temp != nil && (len(temp) > 0) {
		store.StrategyList[chainID] = temp
		logs.Success("Data loaded for the strategyList store for chainID: " + strconv.FormatUint(chainID, 10))
	} else {
		logs.Warning("No strategyList data found for chainID: " + strconv.FormatUint(chainID, 10))
	}
}
