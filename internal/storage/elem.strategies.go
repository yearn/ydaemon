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

var _strategiesSyncMap = make(map[uint64]*sync.Map)
var _strategiesMigratedSyncMap = make(map[uint64]*sync.Map) //This one is never stored but merged with the other one

/** 🔵 - Yearn *************************************************************************************
** The function `loaddStrategiesFromJson` is responsible for loading strategies from a JSON file.
**************************************************************************************************/
func loadStrategiesFromJson(chainID uint64) map[common.Address]models.TStrategyAdded {
	var strategies map[common.Address]models.TStrategyAdded
	chainIDStr := strconv.FormatUint(chainID, 10)

	// Load the JSON file
	file, err := os.Open(env.BASE_DATA_PATH + "/meta/store/strategies/" + chainIDStr + ".json")
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
func StoreStrategiesToJson(chainID uint64, strategies map[common.Address]models.TStrategyAdded) {
	chainIDStr := strconv.FormatUint(chainID, 10)

	file, _ := json.MarshalIndent(strategies, "", "\t")
	if _, err := os.Stat(env.BASE_DATA_PATH + "/meta/store/strategies"); os.IsNotExist(err) {
		os.MkdirAll(env.BASE_DATA_PATH+"/meta/store/strategies", 0755)
	}
	err := os.WriteFile(env.BASE_DATA_PATH+"/meta/store/strategies/"+chainIDStr+".json", file, 0644)
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
	syncMap := _strategiesSyncMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_strategiesSyncMap[chainID] = syncMap
	}
	strategyMap := loadStrategiesFromJson(chainID)
	for _, strategy := range strategyMap {
		syncMap.Store(strategy.Address.Hex(), strategy)
	}
}

/**************************************************************************************************
** StoreStrategy will add a new strategy in the _strategiesSyncMap
**************************************************************************************************/
func StoreStrategy(chainID uint64, strategy models.TStrategyAdded) {
	syncMap := _strategiesSyncMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_strategiesSyncMap[chainID] = syncMap
	}
	syncMap.Store(strategy.Address, strategy)
}

/**************************************************************************************************
** StoreStrategyMigrated will add a new strategy in the _strategiesMigratedSyncMap
**************************************************************************************************/
func StoreStrategyMigrated(chainID uint64, strategy models.TStrategyMigrated) {
	syncMap := _strategiesMigratedSyncMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_strategiesMigratedSyncMap[chainID] = syncMap
	}
	syncMap.Store(strategy.NewStrategyAddress, strategy)
}

/**************************************************************************************************
** ListVault will return a list of all the strategies stored in the caching system for a given
** chainID. Both a map and a slice are returned.
**************************************************************************************************/
func ListStrategies(chainID uint64) (asMap map[common.Address]models.TStrategyAdded, asSlice []models.TStrategyAdded) {
	asMap = make(map[common.Address]models.TStrategyAdded) // make to avoid nil map

	/**********************************************************************************************
	** We first retrieve the syncMap. This syncMap should be initialized first via the `LoadStrategies`
	** function which take the data from the database/badger and store it in it.
	**********************************************************************************************/
	syncMap := _strategiesSyncMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_strategiesSyncMap[chainID] = syncMap
	}

	/**********************************************************************************************
	** We can just iterate over the syncMap and add the strategies to the map and slice.
	** As the stored strategy data are only a subset of static, we need to use the actual structure
	** and not the DBVault one.
	**********************************************************************************************/
	syncMap.Range(func(key, value interface{}) bool {
		strategy := value.(models.TStrategyAdded)
		asMap[strategy.Address] = strategy
		asSlice = append(asSlice, strategy)
		return true
	})

	return asMap, asSlice
}

/**************************************************************************************************
** ListStrategiesMigrated will return a list of all the strategies migrated stored in the caching
** system for a given chainID. Both a map and a slice are returned.
**************************************************************************************************/
func ListStrategiesMigrated(chainID uint64) (
	asMap map[common.Address]models.TStrategyMigrated,
	asSlice []models.TStrategyMigrated,
) {
	asMap = make(map[common.Address]models.TStrategyMigrated) // make to avoid nil map

	/**********************************************************************************************
	** We first retrieve the syncMap. This syncMap should be initialized first via the `LoadStrategies`
	** function which take the data from the database/badger and store it in it.
	**********************************************************************************************/
	syncMap := _strategiesMigratedSyncMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_strategiesMigratedSyncMap[chainID] = syncMap
	}

	/**********************************************************************************************
	** We can just iterate over the syncMap and add the strategies to the map and slice.
	** As the stored strategy data are only a subset of static, we need to use the actual structure
	** and not the DBVault one.
	**********************************************************************************************/
	syncMap.Range(func(key, value interface{}) bool {
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
func GetStrategy(chainID uint64, strategyAddress common.Address) (models.TStrategyAdded, bool) {
	strategyFromSyncMap, ok := _strategiesSyncMap[chainID].Load(strategyAddress)
	if !ok {
		return models.TStrategyAdded{}, false
	}
	return strategyFromSyncMap.(models.TStrategyAdded), true
}

/**************************************************************************************************
** ListStrategiesForVault will return a list of all the strategies stored in the caching system for
** a given chainID and vault address. Both a map and a slice are returned.
**************************************************************************************************/
func ListStrategiesForVault(chainID uint64, vaultAddress common.Address) (asMap map[common.Address]models.TStrategyAdded, asSlice []models.TStrategyAdded) {
	asMap = make(map[common.Address]models.TStrategyAdded) // make to avoid nil map

	/**********************************************************************************************
	** We first retrieve the syncMap. This syncMap should be initialized first via the `LoadStrategies`
	** function which take the data from the database/badger and store it in it.
	**********************************************************************************************/
	syncMap := _strategiesSyncMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_strategiesSyncMap[chainID] = syncMap
	}

	/**********************************************************************************************
	** We can just iterate over the syncMap and add the strategies to the map and slice.
	** As the stored strategy data are only a subset of static, we need to use the actual structure
	** and not the DBVault one.
	**********************************************************************************************/
	syncMap.Range(func(key, value interface{}) bool {
		strategy := value.(models.TStrategyAdded)
		if strategy.VaultAddress != vaultAddress {
			return true
		}
		asMap[strategy.Address] = strategy
		asSlice = append(asSlice, strategy)
		return true
	})

	return asMap, asSlice
}
