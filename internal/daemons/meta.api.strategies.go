package daemons

import (
	"encoding/json"
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/majorfi/ydaemon/internal/logs"
	"github.com/majorfi/ydaemon/internal/models"
	"github.com/majorfi/ydaemon/internal/store"
	"github.com/majorfi/ydaemon/internal/utils"
)

// FetchStrategiesFromMeta fetches the strategies information from the data/meta folder for a given chainID
// and store the result to the global variable StrategiesFromMeta for later use.
func FetchStrategiesFromMeta(chainID uint64) {
	strategies := []models.TStrategyFromMeta{}
	chainIDStr := strconv.FormatUint(chainID, 10)
	content, _, err := utils.ReadAllFilesInDir(utils.META_BASE_PATH+`/strategies/`+chainIDStr+`/`, `.json`)
	if err != nil {
		logs.Error("Error fetching meta information from the Yearn Meta API", err)
		return
	}
	for _, content := range content {
		strategy := models.TStrategyFromMeta{}
		if err := json.Unmarshal(content, &strategy); err != nil {
			logs.Error("Error unmarshalling response body from the Yearn Meta API", err)
			continue
		}
		//Ensure address checksum
		allAddresses := []string{}
		for _, address := range strategy.Addresses {
			allAddresses = append(allAddresses, common.HexToAddress(address).String())
		}
		strategy.Addresses = allAddresses
		strategies = append(strategies, strategy)
	}
	store.RawMetaStrategies[chainID] = strategies

	// To provide faster access to the data, we index the mapping by the vault address, aka
	// {[vaultAddress]: TStrategyFromMeta} if we were working with JS/TS
	if store.StrategiesFromMeta[chainID] == nil {
		store.StrategiesFromMeta[chainID] = make(map[common.Address]models.TStrategyFromMeta)
	}
	for _, strategy := range strategies {
		for _, strategyAddress := range strategy.Addresses {
			store.StrategiesFromMeta[chainID][common.HexToAddress(strategyAddress)] = strategy
		}
	}
}

// LoadMetaStrategies is kept in order to have the same behavior everywhere, but as the data
// exists in the same directory as yDaemon, saving the data in the DB is not necessary.
func LoadMetaStrategies(chainID uint64, wg *sync.WaitGroup) {
	_ = chainID
	wg.Done()
}
