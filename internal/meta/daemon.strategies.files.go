package meta

import (
	"encoding/json"
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/utils/helpers"
	"github.com/yearn/ydaemon/internal/utils/logs"
)

// FetchStrategiesFromMeta fetches the strategies information from the data/meta folder for a given chainID
// and store the result to the global variable StrategiesFromMeta for later use.
func FetchStrategiesFromMeta(chainID uint64) {
	strategies := []TStrategyFromMeta{}
	chainIDStr := strconv.FormatUint(chainID, 10)
	content, _, err := helpers.ReadAllFilesInDir(helpers.BASE_DATA_PATH+`/meta/strategies/`+chainIDStr+`/`, `.json`)
	if err != nil {
		logs.Warning("Error fetching meta information from the Yearn Meta API for chain", chainID)
		return
	}
	for _, content := range content {
		strategy := TStrategyFromMeta{}
		if err := json.Unmarshal(content, &strategy); err != nil {
			logs.Warning("Error unmarshalling response body from the Yearn Meta API for chain", chainID)
			continue
		}
		//Ensure address checksum
		strategy.Addresses = append([]common.Address{}, strategy.Addresses...)
		strategies = append(strategies, strategy)
	}
	Store.RawMetaStrategies[chainID] = strategies

	// To provide faster access to the data, we index the mapping by the vault address, aka
	// {[vaultAddress]: TStrategyFromMeta} if we were working with JS/TS
	if Store.StrategiesFromMeta[chainID] == nil {
		Store.StrategiesFromMeta[chainID] = make(map[common.Address]TStrategyFromMeta)
	}
	for _, strategy := range strategies {
		for _, strategyAddress := range strategy.Addresses {
			Store.StrategiesFromMeta[chainID][strategyAddress] = strategy
		}
	}
}

// LoadMetaStrategies is kept in order to have the same behavior everywhere, but as the data
// exists in the same directory as yDaemon, saving the data in the DB is not necessary.
func LoadMetaStrategies(chainID uint64, wg *sync.WaitGroup) {
	_ = chainID
	wg.Done()
}
