package storage

import (
	"encoding/json"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
)

var _apySyncMap = make(map[uint64]*sync.Map)
var _apyJSONMetadataSyncMap = sync.Map{}
var _apyJSONMutexes = make(map[uint64]*sync.RWMutex)
var _apyJSONMutexesLock sync.Mutex // Protects access to _apyJSONMutexes map

type TJsonAPYStorage struct {
	TJsonMetadata
	APY map[common.Address]models.TVaultAPY `json:"apy"`
}

/** 🔵 - Yearn *************************************************************************************
** getAPYMutex safely gets or creates a mutex for a specific chainID
**************************************************************************************************/
func getAPYMutex(chainID uint64) *sync.RWMutex {
	_apyJSONMutexesLock.Lock()
	defer _apyJSONMutexesLock.Unlock()

	if mutex, exists := _apyJSONMutexes[chainID]; exists {
		return mutex
	}
	_apyJSONMutexes[chainID] = &sync.RWMutex{}
	return _apyJSONMutexes[chainID]
}

/** 🔵 - Yearn *************************************************************************************
** The function `loadAPYFromJson` is responsible for loading APY data from a JSON file.
**************************************************************************************************/
func loadAPYFromJson(chainID uint64) TJsonAPYStorage {
	var apyData TJsonAPYStorage
	chainIDStr := strconv.FormatUint(chainID, 10)

	// Load the JSON file
	file, err := os.Open(env.BASE_DATA_PATH + "/meta/apy/" + chainIDStr + ".json")
	if err != nil {
		return TJsonAPYStorage{}
	}
	defer file.Close()

	// Decode the JSON file into the map
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&apyData)
	if err != nil {
		logs.Error("Failed to decode APY JSON file: " + err.Error())
		return TJsonAPYStorage{}
	}

	return apyData
}

/** 🔵 - Yearn *************************************************************************************
** The function `storeAPYToJson` is responsible for storing APY data to a JSON file.
**************************************************************************************************/
func StoreAPYToJson(chainID uint64, apyData map[common.Address]models.TVaultAPY) {
	mutex := getAPYMutex(chainID)
	mutex.Lock()
	defer mutex.Unlock()

	chainIDStr := strconv.FormatUint(chainID, 10)
	previousAPY := loadAPYFromJson(chainID)
	version := detectVersionUpdate(chainID, previousAPY.Version, previousAPY.APY, apyData)

	data := TJsonAPYStorage{
		TJsonMetadata: TJsonMetadata{
			LastUpdate: time.Now(),
			Version:    version,
		},
		APY: apyData,
	}
	_apyJSONMetadataSyncMap.Store(chainID, TJsonMetadata{
		data.LastUpdate,
		data.Version,
		data.ShouldRefresh,
	})

	file, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		logs.Error("Failed to marshal APY JSON file: " + err.Error())
		return
	}
	if _, err := os.Stat(env.BASE_DATA_PATH + "/meta/apy"); os.IsNotExist(err) {
		os.MkdirAll(env.BASE_DATA_PATH+"/meta/apy", 0755)
	}
	err = os.WriteFile(env.BASE_DATA_PATH+"/meta/apy/"+chainIDStr+".json", file, 0644)
	if err != nil {
		logs.Error("Failed to write APY JSON file: " + err.Error())
	}
}

/**************************************************************************************************
** Retrieve the last time the APY data was updated for a specific chainID
**************************************************************************************************/
func GetAPYJsonMetadata(chainID uint64) TJsonMetadata {
	if jsonMetadata, ok := _apyJSONMetadataSyncMap.Load(chainID); ok {
		return jsonMetadata.(TJsonMetadata)
	}
	return TJsonMetadata{}
}

/**************************************************************************************************
** LoadAPY will retrieve all the APY data from the JSON file and store them in the
** _apySyncMap for fast access during that same execution.
**************************************************************************************************/
func LoadAPY(chainID uint64, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	mutex := getAPYMutex(chainID)
	mutex.RLock()
	defer mutex.RUnlock()

	file := loadAPYFromJson(chainID)
	_apyJSONMetadataSyncMap.Store(chainID, TJsonMetadata{
		file.LastUpdate,
		file.Version,
		file.ShouldRefresh,
	})
	for address, apy := range file.APY {
		safeSyncMap(_apySyncMap, chainID).Store(address, apy)
	}
}

/**************************************************************************************************
** StoreAPY will add a new APY entry in the _apySyncMap
**************************************************************************************************/
func StoreAPY(chainID uint64, vaultAddress common.Address, apy models.TVaultAPY) {
	safeSyncMap(_apySyncMap, chainID).Store(vaultAddress, apy)
}

/**************************************************************************************************
** GetAPY will return the APY data for a specific vault address on a given chainID
**************************************************************************************************/
func GetAPY(chainID uint64, vaultAddress common.Address) (models.TVaultAPY, bool) {
	apyFromSyncMap, ok := safeSyncMap(_apySyncMap, chainID).Load(vaultAddress)
	if !ok {
		return models.TVaultAPY{}, false
	}
	return apyFromSyncMap.(models.TVaultAPY), true
}

/**************************************************************************************************
** ListAPY will return a list of all the APY data stored in the caching system for a given
** chainID. Both a map and a slice are returned.
**************************************************************************************************/
func ListAPY(chainID uint64) (map[common.Address]models.TVaultAPY, []models.TVaultAPY) {
	apyMap := make(map[common.Address]models.TVaultAPY)
	apySlice := make([]models.TVaultAPY, 0)

	safeSyncMap(_apySyncMap, chainID).Range(func(key, value interface{}) bool {
		vaultAddress := key.(common.Address)
		apy := value.(models.TVaultAPY)
		apyMap[vaultAddress] = apy
		apySlice = append(apySlice, apy)
		return true
	})

	return apyMap, apySlice
}
