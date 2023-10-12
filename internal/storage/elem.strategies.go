package storage

import (
	"encoding/json"
	"os"
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
)

// The _strategiesSyncMap is in JSON and contains the strategies informations
var _strategiesSyncMap = make(map[uint64]*sync.Map)

// The _strategiesMigratedSyncMap is never stored in JSON but used for internal caching
var _strategiesMigratedSyncMap = make(map[uint64]*sync.Map)

/** 🔵 - Yearn *************************************************************************************
** The function `loaddStrategiesFromJson` is responsible for loading strategies from a JSON file.
**************************************************************************************************/
func loadStrategiesFromJson(chainID uint64) map[common.Address]models.TStrategy {
	var strategies map[common.Address]models.TStrategy
	chainIDStr := strconv.FormatUint(chainID, 10)

	// Load the JSON file
	file, err := os.Open(env.BASE_DATA_PATH + "/meta/strategies/" + chainIDStr + ".json")
	if err != nil {
		return nil
	}
	defer file.Close()

	// Decode the JSON file into the map
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&strategies)
	if err != nil {
		logs.Error("Failed to decode strategies JSON file: " + err.Error())
		return nil
	}

	return strategies
}

/** 🔵 - Yearn *************************************************************************************
** The function `storedStrategiesToJson` is responsible for storing strategies to a JSON file.
**************************************************************************************************/
func StoreStrategiesToJson(chainID uint64, strategies map[common.Address]models.TStrategy) {
	chainIDStr := strconv.FormatUint(chainID, 10)

	file, _ := json.MarshalIndent(strategies, "", "\t")
	if _, err := os.Stat(env.BASE_DATA_PATH + "/meta/strategies"); os.IsNotExist(err) {
		os.MkdirAll(env.BASE_DATA_PATH+"/meta/strategies", 0755)
	}
	err := os.WriteFile(env.BASE_DATA_PATH+"/meta/strategies/"+chainIDStr+".json", file, 0644)
	if err != nil {
		logs.Error("Failed to write strategies JSON file: " + err.Error())
	}
}

/**************************************************************************************************
** LoadStrategies will retrieve the all the strategies from the configured DB and store them in the
** _strategiesSyncMap for fast access during that same execution.
**************************************************************************************************/
func LoadStrategies(chainID uint64, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	strategyMap := loadStrategiesFromJson(chainID)
	for _, strategy := range strategyMap {
		safeSyncMap(_strategiesSyncMap, chainID).Store(strategy.Address, strategy)
	}
}

/**************************************************************************************************
** StoreStrategy will add a new strategy in the _strategiesSyncMap
** StoreStrategyMigrated will add a new strategy in the _strategiesMigratedSyncMap
** StoreStrategyIfMissing is used to only add a strategy if it doesn't already exist. This is
** usefull as the strategy loading is in two steps: first the indexing, then the data.
**************************************************************************************************/
func StoreStrategy(chainID uint64, strategy models.TStrategy) {
	safeSyncMap(_strategiesSyncMap, chainID).Store(strategy.Address, strategy)
}
func StoreStrategyMigrated(chainID uint64, strategy models.TStrategyMigrated) {
	safeSyncMap(_strategiesMigratedSyncMap, chainID).Store(strategy.NewStrategyAddress, strategy)
}
func StoreStrategyIfMissing(chainID uint64, strategy models.TStrategy) bool {
	syncMap := safeSyncMap(_strategiesSyncMap, chainID)
	if _, ok := syncMap.Load(strategy.Address); !ok {
		syncMap.Store(strategy.Address, strategy)
		return true
	}
	return false
}

/**************************************************************************************************
** ListStrategies will return a list of all the strategies stored
** ListStrategiesMigrated will return a list of all the strategies migrated stored
** Each returns both a map and a slice.
**************************************************************************************************/
func ListStrategies(chainID uint64) (
	asMap map[common.Address]models.TStrategy,
	asSlice []models.TStrategy,
) {
	asMap = make(map[common.Address]models.TStrategy) // make to avoid nil map

	/**********************************************************************************************
	** We can just iterate over the syncMap and add the strategies to the map and slice.
	** As the stored strategy data are only a subset of static, we need to use the actual structure
	**********************************************************************************************/
	safeSyncMap(_strategiesSyncMap, chainID).Range(func(key, value interface{}) bool {
		strategy := value.(models.TStrategy)
		asMap[strategy.Address] = strategy
		asSlice = append(asSlice, strategy)
		return true
	})

	return asMap, asSlice
}
func ListStrategiesMigrated(chainID uint64) (
	asMap map[common.Address]models.TStrategyMigrated,
	asSlice []models.TStrategyMigrated,
) {
	asMap = make(map[common.Address]models.TStrategyMigrated) // make to avoid nil map

	/**********************************************************************************************
	** We can just iterate over the syncMap and add the strategies to the map and slice.
	** As the stored strategy data are only a subset of static, we need to use the actual structure
	**********************************************************************************************/
	safeSyncMap(_strategiesMigratedSyncMap, chainID).Range(func(key, value interface{}) bool {
		strategy := value.(models.TStrategyMigrated)
		asMap[strategy.NewStrategyAddress] = strategy
		asSlice = append(asSlice, strategy)
		return true
	})

	return asMap, asSlice
}

/**************************************************************************************************
** GetStrategy will return a single strategy stored in the caching system for a given pair of
** chainID and strategy address.
**************************************************************************************************/
func GetStrategy(chainID uint64, strategyAddress common.Address) (models.TStrategy, bool) {
	strategyFromSyncMap, ok := safeSyncMap(_strategiesSyncMap, chainID).Load(strategyAddress)
	if !ok {
		return models.TStrategy{}, false
	}
	return strategyFromSyncMap.(models.TStrategy), true
}

/**************************************************************************************************
** ListStrategiesForVault will return a list of all the strategies stored in the caching system for
** a given chainID and vault address. Both a map and a slice are returned.
**************************************************************************************************/
func ListStrategiesForVault(chainID uint64, vaultAddress common.Address) (
	asMap map[common.Address]models.TStrategy,
	asSlice []models.TStrategy,
) {
	asMap = make(map[common.Address]models.TStrategy) // make to avoid nil map

	/**********************************************************************************************
	** We can just iterate over the syncMap and add the strategies to the map and slice.
	** As the stored strategy data are only a subset of static, we need to use the actual structure
	**********************************************************************************************/
	safeSyncMap(_strategiesSyncMap, chainID).Range(func(key, value interface{}) bool {
		strategy := value.(models.TStrategy)
		if strategy.VaultAddress != vaultAddress {
			return true
		}
		asMap[strategy.Address] = strategy
		asSlice = append(asSlice, strategy)
		return true
	})

	return asMap, asSlice
}
