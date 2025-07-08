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

var _pricesJSONMetadataSyncMap = sync.Map{}
var _pricesJSONMutexes = make(map[uint64]*sync.RWMutex)
var _pricesJSONMutexesLock sync.Mutex // Protects access to _pricesJSONMutexes map

type TJsonPricesStorage struct {
	TJsonMetadata
	Prices map[common.Address]models.TPrices `json:"prices"`
}

/** 🔵 - Yearn *************************************************************************************
** getPricesMutex safely gets or creates a mutex for a specific chainID
**************************************************************************************************/
func getPricesMutex(chainID uint64) *sync.RWMutex {
	_pricesJSONMutexesLock.Lock()
	defer _pricesJSONMutexesLock.Unlock()

	if mutex, exists := _pricesJSONMutexes[chainID]; exists {
		return mutex
	}
	_pricesJSONMutexes[chainID] = &sync.RWMutex{}
	return _pricesJSONMutexes[chainID]
}

/** 🔵 - Yearn *************************************************************************************
** The function `loadPricesFromJson` is responsible for loading prices from a JSON file.
**************************************************************************************************/
func loadPricesFromJson(chainID uint64) TJsonPricesStorage {
	var pricesData TJsonPricesStorage
	chainIDStr := strconv.FormatUint(chainID, 10)

	// Load the JSON file
	file, err := os.Open(env.BASE_DATA_PATH + "/meta/prices/" + chainIDStr + ".json")
	if err != nil {
		return TJsonPricesStorage{}
	}
	defer file.Close()

	// Decode the JSON file into the map
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&pricesData)
	if err != nil {
		logs.Error("Failed to decode prices JSON file: " + err.Error())
		return TJsonPricesStorage{}
	}

	return pricesData
}

/** 🔵 - Yearn *************************************************************************************
** The function `storePricesToJson` is responsible for storing prices to a JSON file.
**************************************************************************************************/
func StorePricesToJson(chainID uint64, pricesData map[common.Address]models.TPrices) {
	mutex := getPricesMutex(chainID)
	mutex.Lock()
	defer mutex.Unlock()

	chainIDStr := strconv.FormatUint(chainID, 10)
	previousPrices := loadPricesFromJson(chainID)
	version := detectVersionUpdate(chainID, previousPrices.Version, previousPrices.Prices, pricesData)

	data := TJsonPricesStorage{
		TJsonMetadata: TJsonMetadata{
			LastUpdate: time.Now(),
			Version:    version,
		},
		Prices: pricesData,
	}
	_pricesJSONMetadataSyncMap.Store(chainID, TJsonMetadata{
		data.LastUpdate,
		data.Version,
		data.ShouldRefresh,
	})

	file, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		logs.Error("Failed to marshal prices JSON file: " + err.Error())
		return
	}
	if _, err := os.Stat(env.BASE_DATA_PATH + "/meta/prices"); os.IsNotExist(err) {
		os.MkdirAll(env.BASE_DATA_PATH+"/meta/prices", 0755)
	}
	err = os.WriteFile(env.BASE_DATA_PATH+"/meta/prices/"+chainIDStr+".json", file, 0644)
	if err != nil {
		logs.Error("Failed to write prices JSON file: " + err.Error())
	}
}

/**************************************************************************************************
** Retrieve the last time the prices data was updated for a specific chainID
**************************************************************************************************/
func GetPricesJsonMetadata(chainID uint64) TJsonMetadata {
	if jsonMetadata, ok := _pricesJSONMetadataSyncMap.Load(chainID); ok {
		return jsonMetadata.(TJsonMetadata)
	}
	return TJsonMetadata{}
}

/**************************************************************************************************
** LoadPrices will retrieve all the prices data from the JSON file and store them in the
** _pricesSyncMap for fast access during that same execution.
**************************************************************************************************/
func LoadPrices(chainID uint64, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	mutex := getPricesMutex(chainID)
	mutex.RLock()
	defer mutex.RUnlock()

	file := loadPricesFromJson(chainID)
	_pricesJSONMetadataSyncMap.Store(chainID, TJsonMetadata{
		file.LastUpdate,
		file.Version,
		file.ShouldRefresh,
	})
	for address, price := range file.Prices {
		safeSyncMap(_pricesSyncMap, chainID).Store(address, price)
	}
}
