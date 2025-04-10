package storage

import (
	"encoding/json"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
)

// The _strategiesSyncMap is in JSON and contains the strategies informations
var _strategiesSyncMap = make(map[uint64]*sync.Map)
var _strategiesMigratedSyncMap = make(map[uint64]*sync.Map)
var _strategiesJSONMetadataSyncMap = sync.Map{}
var _strategyJSONMutexes = make(map[uint64]*sync.RWMutex)
var _strategyJSONMutexesLock sync.Mutex // Protects access to _strategyJSONMutexes map

/** 🔵 - Yearn *************************************************************************************
** getStrategyMutex safely gets or creates a mutex for a specific chainID
**************************************************************************************************/
func getStrategyMutex(chainID uint64) *sync.RWMutex {
	_strategyJSONMutexesLock.Lock()
	defer _strategyJSONMutexesLock.Unlock()

	if mutex, exists := _strategyJSONMutexes[chainID]; exists {
		return mutex
	}
	_strategyJSONMutexes[chainID] = &sync.RWMutex{}
	return _strategyJSONMutexes[chainID]
}

/** 🔵 - Yearn *************************************************************************************
** The function `loaddStrategiesFromJson` is responsible for loading strategies from a JSON file.
**************************************************************************************************/
func loadStrategiesFromJson(chainID uint64) TJsonStrategyStorage {
	var strategyFile TJsonStrategyStorage
	chainIDStr := strconv.FormatUint(chainID, 10)

	// Load the JSON file
	file, err := os.Open(env.BASE_DATA_PATH + "/meta/strategies/" + chainIDStr + ".json")
	if err != nil {
		return TJsonStrategyStorage{}
	}
	defer file.Close()

	retry := 0
	for {
		// Decode the JSON file into the map
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&strategyFile)
		if err != nil {
			logs.Error("Failed to decode strategies JSON file: " + err.Error())
			time.Sleep(1 * time.Second)
			retry++
			if retry > 5 {
				return TJsonStrategyStorage{}
			}
			continue
		}
		if retry > 0 {
			logs.Success(`Succeed to decode vaults JSON file on chainID ` + chainIDStr + ` after ` + strconv.Itoa(retry) + ` retries`)
		}
		break
	}
	return strategyFile
}

/** 🔵 - Yearn *************************************************************************************
** The function `storedStrategiesToJson` is responsible for storing strategies to a JSON file.
**************************************************************************************************/
func StoreStrategiesToJson(chainID uint64, strategies map[string]models.TStrategy) {
	mutex := getStrategyMutex(chainID)
	mutex.Lock()
	defer mutex.Unlock()

	chainIDStr := strconv.FormatUint(chainID, 10)
	previousStrategies := loadStrategiesFromJson(chainID)
	version := detectStrVersionUpdate(chainID, previousStrategies.Version, previousStrategies.Strategies, strategies)

	data := TJsonStrategyStorage{
		TJsonMetadata: TJsonMetadata{
			LastUpdate: time.Now(),
			Version:    version,
		},
		Strategies: strategies,
	}
	_strategiesJSONMetadataSyncMap.Store(chainID, TJsonMetadata{
		data.LastUpdate,
		data.Version,
		data.ShouldRefresh,
	})

	file, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		logs.Error("Failed to marshal strategies JSON file: " + err.Error())
	}
	if _, err := os.Stat(env.BASE_DATA_PATH + "/meta/strategies"); os.IsNotExist(err) {
		os.MkdirAll(env.BASE_DATA_PATH+"/meta/strategies", 0755)
	}
	if err := os.WriteFile(env.BASE_DATA_PATH+"/meta/strategies/"+chainIDStr+".json", file, 0644); err != nil {
		logs.Error("Failed to write strategies JSON file: " + err.Error())
	}
}

/**************************************************************************************************
** Retrieve the last time the strategies were updated for a specific chainID
**************************************************************************************************/
func GetStrategiesJsonMetadata(chainID uint64) TJsonMetadata {
	if jsonMetadata, ok := _strategiesJSONMetadataSyncMap.Load(chainID); ok {
		return jsonMetadata.(TJsonMetadata)
	}
	return TJsonMetadata{}
}

/**************************************************************************************************
** LoadStrategies will retrieve the all the strategies from the configured DB and store them in the
** _strategiesSyncMap for fast access during that same execution.
**************************************************************************************************/
func LoadStrategies(chainID uint64, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	mutex := getStrategyMutex(chainID)
	mutex.RLock()
	defer mutex.RUnlock()

	file := loadStrategiesFromJson(chainID)
	_strategiesJSONMetadataSyncMap.Store(chainID, TJsonMetadata{
		file.LastUpdate,
		file.Version,
		file.ShouldRefresh,
	})
	for _, strategy := range file.Strategies {
		strategyKey := strategy.Address.Hex() + `_` + strategy.VaultAddress.Hex()
		safeSyncMap(_strategiesSyncMap, chainID).Store(strategyKey, strategy)
	}
}

/**************************************************************************************************
** StoreStrategy will add a new strategy in the _strategiesSyncMap
** StoreStrategyMigrated will add a new strategy in the _strategiesMigratedSyncMap
** StoreStrategyIfMissing is used to only add a strategy if it doesn't already exist. This is
** usefull as the strategy loading is in two steps: first the indexing, then the data.
**************************************************************************************************/
func StoreStrategy(chainID uint64, strategy models.TStrategy) {
	strategyKey := strategy.Address.Hex() + `_` + strategy.VaultAddress.Hex()
	safeSyncMap(_strategiesSyncMap, chainID).Store(strategyKey, strategy)
}
func StoreStrategyMigrated(chainID uint64, strategy models.TStrategyMigrated) {
	strategyKey := strategy.NewStrategyAddress.Hex() + `_` + strategy.VaultAddress.Hex()
	safeSyncMap(_strategiesMigratedSyncMap, chainID).Store(strategyKey, strategy)
}
func StoreStrategyIfMissing(chainID uint64, strategy models.TStrategy) bool {
	syncMap := safeSyncMap(_strategiesSyncMap, chainID)
	strategyKey := strategy.Address.Hex() + `_` + strategy.VaultAddress.Hex()
	if _, ok := syncMap.Load(strategyKey); !ok {
		syncMap.Store(strategyKey, strategy)
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
	asMap map[string]models.TStrategy,
	asSlice []models.TStrategy,
) {
	asMap = make(map[string]models.TStrategy) // make to avoid nil map

	/**********************************************************************************************
	** We can just iterate over the syncMap and add the strategies to the map and slice.
	** As the stored strategy data are only a subset of static, we need to use the actual structure
	**********************************************************************************************/
	safeSyncMap(_strategiesSyncMap, chainID).Range(func(_, value interface{}) bool {
		strategy := value.(models.TStrategy)
		strategyKey := strategy.Address.Hex() + `_` + strategy.VaultAddress.Hex()
		asMap[strategyKey] = strategy
		asSlice = append(asSlice, strategy)
		return true
	})

	return asMap, asSlice
}
func ListStrategiesMigrated(chainID uint64) (
	asMap map[string]models.TStrategyMigrated,
	asSlice []models.TStrategyMigrated,
) {
	asMap = make(map[string]models.TStrategyMigrated) // make to avoid nil map

	/**********************************************************************************************
	** We can just iterate over the syncMap and add the strategies to the map and slice.
	** As the stored strategy data are only a subset of static, we need to use the actual structure
	**********************************************************************************************/
	safeSyncMap(_strategiesMigratedSyncMap, chainID).Range(func(_, value interface{}) bool {
		strategy := value.(models.TStrategyMigrated)
		strategyKey := strategy.NewStrategyAddress.Hex() + `_` + strategy.VaultAddress.Hex()
		asMap[strategyKey] = strategy
		asSlice = append(asSlice, strategy)
		return true
	})

	return asMap, asSlice
}

/**************************************************************************************************
** GetStrategy will return a single strategy stored in the caching system for a given pair of
** chainID and strategyKey.
**************************************************************************************************/
func GetStrategy(
	chainID uint64,
	strategyAddress common.Address,
	vaultAddress common.Address,
) (models.TStrategy, bool) {
	strategyKey := strategyAddress.Hex() + `_` + vaultAddress.Hex()
	strategyFromSyncMap, ok := safeSyncMap(_strategiesSyncMap, chainID).Load(strategyKey)
	if !ok {
		return models.TStrategy{}, false
	}
	return strategyFromSyncMap.(models.TStrategy), true
}

/**************************************************************************************************
** GuessStrategy will return a single strategy stored in the caching system for a given pair of
** chainID and strategy address.
**************************************************************************************************/
func GuessStrategy(
	chainID uint64,
	strategyAddress common.Address,
) (models.TStrategy, bool) {
	strat := models.TStrategy{}
	safeSyncMap(_strategiesSyncMap, chainID).Range(func(_, value interface{}) bool {
		strategy := value.(models.TStrategy)
		if strategy.Address == strategyAddress {
			strat = strategy
			return true
		}
		return true
	})

	if addresses.Equals(strat.Address, common.Address{}) {
		return models.TStrategy{}, false
	}
	return strat, true
}

/**************************************************************************************************
** ListStrategiesForVault will return a list of all the strategies stored in the caching system for
** a given chainID and vault address. Both a map and a slice are returned.
**************************************************************************************************/
func ListStrategiesForVault(chainID uint64, vaultAddress common.Address) (
	asMap map[string]models.TStrategy,
	asSlice []models.TStrategy,
) {
	asMap = make(map[string]models.TStrategy) // make to avoid nil map

	/**********************************************************************************************
	** We can just iterate over the syncMap and add the strategies to the map and slice.
	** As the stored strategy data are only a subset of static, we need to use the actual structure
	**********************************************************************************************/
	safeSyncMap(_strategiesSyncMap, chainID).Range(func(_, value interface{}) bool {
		strategy := value.(models.TStrategy)
		if !addresses.Equals(strategy.VaultAddress, vaultAddress) {
			return true
		}
		strategyKey := strategy.Address.Hex() + `_` + strategy.VaultAddress.Hex()
		asMap[strategyKey] = strategy
		asSlice = append(asSlice, strategy)
		return true
	})

	return asMap, asSlice
}
