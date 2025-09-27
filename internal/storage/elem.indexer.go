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
)

// The _indexerSyncMap contains the last indexed blocks for each vault
var _indexerSyncMap = make(map[uint64]*sync.Map)
var _indexerJSONMutexes = make(map[uint64]*sync.RWMutex)
var _indexerJSONMutexesLock sync.Mutex // Protects access to _indexerJSONMutexes map

/** 🔵 - Yearn *************************************************************************************
** TJsonIndexerStorage structure for persisting indexer state to JSON files
**************************************************************************************************/
type TJsonIndexerStorage struct {
	TJsonMetadata
	LastIndexedBlocks map[string]uint64 `json:"lastIndexedBlocks"`
}

/** 🔵 - Yearn *************************************************************************************
** getIndexerMutex safely gets or creates a mutex for a specific chainID
**************************************************************************************************/
func getIndexerMutex(chainID uint64) *sync.RWMutex {
	_indexerJSONMutexesLock.Lock()
	defer _indexerJSONMutexesLock.Unlock()

	if mutex, exists := _indexerJSONMutexes[chainID]; exists {
		return mutex
	}
	_indexerJSONMutexes[chainID] = &sync.RWMutex{}
	return _indexerJSONMutexes[chainID]
}

/** 🔵 - Yearn *************************************************************************************
** loadIndexerFromJson loads the last indexed blocks from a JSON file
**************************************************************************************************/
func loadIndexerFromJson(chainID uint64) TJsonIndexerStorage {
	var indexerFile TJsonIndexerStorage
	chainIDStr := strconv.FormatUint(chainID, 10)

	// Load the JSON file
	file, err := os.Open(env.BASE_DATA_PATH + "/meta/indexer/" + chainIDStr + ".json")
	if err != nil {
		return TJsonIndexerStorage{
			TJsonMetadata: TJsonMetadata{
				LastUpdate:    time.Now(),
				Version:       TVersion{Major: 0, Minor: 1, Patch: 0},
				ShouldRefresh: false,
			},
			LastIndexedBlocks: make(map[string]uint64),
		}
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&indexerFile)
	if err != nil {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Duration(100*(i+1)) * time.Millisecond)
			file, err = os.Open(env.BASE_DATA_PATH + "/meta/indexer/" + chainIDStr + ".json")
			if err != nil {
				continue
			}
			defer file.Close()
			decoder := json.NewDecoder(file)
			err = decoder.Decode(&indexerFile)
			if err == nil {
				break
			}
		}
		if err != nil {
			logs.Error(`Failed to decode indexer json file for chain ` + chainIDStr + `: ` + err.Error())
			return TJsonIndexerStorage{
				TJsonMetadata: TJsonMetadata{
					LastUpdate:    time.Now(),
					Version:       TVersion{Major: 0, Minor: 1, Patch: 0},
					ShouldRefresh: false,
				},
				LastIndexedBlocks: make(map[string]uint64),
			}
		}
	}

	// Store in memory map
	if _, ok := _indexerSyncMap[chainID]; !ok {
		_indexerSyncMap[chainID] = &sync.Map{}
	}
	for vaultAddr, blockNum := range indexerFile.LastIndexedBlocks {
		_indexerSyncMap[chainID].Store(vaultAddr, blockNum)
	}

	return indexerFile
}

/** 🔵 - Yearn *************************************************************************************
** StoreIndexerToJson saves the last indexed blocks to a JSON file
**************************************************************************************************/
func StoreIndexerToJson(chainID uint64) {
	chainIDStr := strconv.FormatUint(chainID, 10)

	// Get mutex for this chain
	mutex := getIndexerMutex(chainID)
	mutex.Lock()
	defer mutex.Unlock()

	// Create the directory if it doesn't exist
	if _, err := os.Stat(env.BASE_DATA_PATH + "/meta/indexer"); os.IsNotExist(err) {
		os.MkdirAll(env.BASE_DATA_PATH+"/meta/indexer", 0755)
	}

	// Convert sync.Map to regular map for JSON serialization
	lastIndexedBlocks := make(map[string]uint64)
	if chainMap, ok := _indexerSyncMap[chainID]; ok {
		chainMap.Range(func(key, value interface{}) bool {
			if addr, ok := key.(string); ok {
				if blockNum, ok := value.(uint64); ok {
					lastIndexedBlocks[addr] = blockNum
				}
			}
			return true
		})
	}

	// Prepare the data structure
	data := TJsonIndexerStorage{
		TJsonMetadata: TJsonMetadata{
			LastUpdate:    time.Now(),
			Version:       TVersion{Major: 0, Minor: 1, Patch: 0},
			ShouldRefresh: false,
		},
		LastIndexedBlocks: lastIndexedBlocks,
	}

	// Marshal with indentation
	file, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		logs.Error(`Failed to marshal indexer json file for chain ` + chainIDStr + `: ` + err.Error())
		return
	}

	// Write to file
	err = os.WriteFile(env.BASE_DATA_PATH+"/meta/indexer/"+chainIDStr+".json", file, 0644)
	if err != nil {
		logs.Error(`Failed to write indexer json file for chain ` + chainIDStr + `: ` + err.Error())
	}
}

/** 🔵 - Yearn *************************************************************************************
** GetLastIndexedBlock retrieves the last indexed block for a specific vault
**************************************************************************************************/
func GetLastIndexedBlock(chainID uint64, vault common.Address) uint64 {
	// Ensure the map is initialized
	if _, ok := _indexerSyncMap[chainID]; !ok {
		_indexerSyncMap[chainID] = &sync.Map{}
		loadIndexerFromJson(chainID)
	}

	// Get the value from the sync.Map
	vaultAddr := vault.Hex()
	if value, ok := _indexerSyncMap[chainID].Load(vaultAddr); ok {
		if blockNum, ok := value.(uint64); ok {
			return blockNum
		}
	}

	return 0
}

/** 🔵 - Yearn *************************************************************************************
** SetLastIndexedBlock stores the last indexed block for a specific vault
**************************************************************************************************/
func SetLastIndexedBlock(chainID uint64, vault common.Address, block uint64) {
	// Ensure the map is initialized
	if _, ok := _indexerSyncMap[chainID]; !ok {
		_indexerSyncMap[chainID] = &sync.Map{}
	}

	// Store in memory
	vaultAddr := vault.Hex()
	_indexerSyncMap[chainID].Store(vaultAddr, block)

	// Persist to disk (you might want to throttle this for performance)
	// For now, we save immediately but you could batch these
	StoreIndexerToJson(chainID)
}

/** 🔵 - Yearn *************************************************************************************
** InitializeIndexerStorage loads the persisted indexer state from disk
**************************************************************************************************/
func InitializeIndexerStorage(chainID uint64) {
	loadIndexerFromJson(chainID)
}