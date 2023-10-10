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

var _newVaultsFromRegistrySyncMap = make(map[uint64]*sync.Map)

/** 🔵 - Yearn *************************************************************************************
** The function `loadHistoricalVaultsFromJson` is responsible for loading historical vaults from
** a JSON file. It reads the JSON file, parses the data into a map of vault addresses to
** `TVaultsFromRegistry` objects, and returns this map along with the highest block number found
** among all vaults. This function is used to initialize the state of the vaults from a previously
** saved state.
**************************************************************************************************/
func LoadHistoricalVaultsFromJson(chainID uint64) (map[common.Address]models.TVaultsFromRegistry, uint64) {
	var historicalVaults map[common.Address]models.TVaultsFromRegistry
	var highestBlockNumber uint64
	chainIDStr := strconv.FormatUint(chainID, 10)

	// Load the JSON file
	file, err := os.Open(env.BASE_DATA_PATH + "/meta/store/registries/" + chainIDStr + ".json")
	if err != nil {
		return nil, 0
	}
	defer file.Close()

	// Decode the JSON file into the map
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&historicalVaults)
	if err != nil {
		logs.Error("Failed to decode vaults JSON file: " + err.Error())
		return nil, 0
	}

	// Find the highest block number
	for _, vault := range historicalVaults {
		if vault.BlockNumber > highestBlockNumber {
			highestBlockNumber = vault.BlockNumber
		}
	}

	return historicalVaults, highestBlockNumber
}

/** 🔵 - Yearn *************************************************************************************
** The function `storeHistoricalVaultsToJson` is responsible for storing historical vaults to a
** JSON file. It takes a map of vault addresses to `TVaultsFromRegistry` objects, and writes this
** map to a JSON file. This function is used to save the state of the vaults for later use.
**************************************************************************************************/
func StoreHistoricalVaultsToJson(chainID uint64, historicalVaults map[common.Address]models.TVaultsFromRegistry) {
	chainIDStr := strconv.FormatUint(chainID, 10)

	file, _ := json.MarshalIndent(historicalVaults, "", "\t")
	if _, err := os.Stat(env.BASE_DATA_PATH + "/meta/store/registries"); os.IsNotExist(err) {
		os.MkdirAll(env.BASE_DATA_PATH+"/meta/store/registries", 0755)
	}
	err := os.WriteFile(env.BASE_DATA_PATH+"/meta/store/registries/"+chainIDStr+".json", file, 0644)
	if err != nil {
		logs.Error("Failed to write vaults JSON file: " + err.Error())
	}
}

/**************************************************************************************************
** LoadNewVaultsFromRegistry will retrieve the all the vaults added to the registries from the
** local DB and store them in the _newVaultsFromRegistrySyncMap for fast access during that same
** execution.
** Use local DB to not screw up the DB with the same data over and over again.
**************************************************************************************************/
func LoadNewVaultsFromRegistry(chainID uint64, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	syncMap := _newVaultsFromRegistrySyncMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_newVaultsFromRegistrySyncMap[chainID] = syncMap
	}

	historicalVaultsMap, _ := LoadHistoricalVaultsFromJson(chainID)
	for _, vault := range historicalVaultsMap {
		syncMap.Store(vault.Address.Hex(), vault)
	}
}

/**************************************************************************************************
** StoreNewVaultsFromRegistry will add a new vault in the _vaultsSyncMap
**************************************************************************************************/
func StoreNewVaultsFromRegistry(chainID uint64, vault models.TVaultsFromRegistry) {
	syncMap := _newVaultsFromRegistrySyncMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_newVaultsFromRegistrySyncMap[chainID] = syncMap
	}
	syncMap.Store(vault.Address.Hex(), vault)
}

/**************************************************************************************************
** ListLastNewVaultEventForRegistries will return a map of the last blockNumber for each registry
** that has been registered. This will allow us to know when to start listening for new events
** instead of going through the whole history.
**************************************************************************************************/
func ListLastNewVaultEventForRegistries(chainID uint64) map[string]uint64 {
	lastRegisteredEventForRegistry := make(map[string]uint64)

	/**********************************************************************************************
	** We first retrieve the syncMap. This syncMap should be initialized first via the
	** `LoadNewVaultsFromRegistry` function which take the data from the database/badger and store
	** it in it.
	**********************************************************************************************/
	syncMap := _newVaultsFromRegistrySyncMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_newVaultsFromRegistrySyncMap[chainID] = syncMap
	}

	syncMap.Range(func(key, value interface{}) bool {
		registry := value.(models.TVaultsFromRegistry).RegistryAddress
		blockNumber := value.(models.TVaultsFromRegistry).BlockNumber

		if _, ok := lastRegisteredEventForRegistry[registry.Hex()]; !ok {
			lastRegisteredEventForRegistry[registry.Hex()] = blockNumber
		} else if lastRegisteredEventForRegistry[registry.Hex()] < blockNumber {
			lastRegisteredEventForRegistry[registry.Hex()] = blockNumber
		}
		return true
	})

	return lastRegisteredEventForRegistry
}

/**************************************************************************************************
** ListVaultsFromRegistry will return a list of all the vaults stored in the caching system for a
** given chainID. Both a map and a slice are returned.
** This is for the registry vault version.
**************************************************************************************************/
func ListVaultsFromRegistry(chainID uint64) (asMap map[common.Address]models.TVaultsFromRegistry, asSlice []models.TVaultsFromRegistry) {
	asMap = make(map[common.Address]models.TVaultsFromRegistry) // make to avoid nil map

	/**********************************************************************************************
	** We first retrieve the syncMap. This syncMap should be initialized first via the `LoadNewVaultsFromRegistry`
	** function which take the data from the database/badger and store it in it.
	**********************************************************************************************/
	syncMap := _newVaultsFromRegistrySyncMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_newVaultsFromRegistrySyncMap[chainID] = syncMap
	}

	/**********************************************************************************************
	** We can just iterate over the syncMap and add the vaults to the map and slice.
	** As the stored vault data are only a subset of static, we need to use the actual structure
	** and not the DBVault one.
	**********************************************************************************************/
	syncMap.Range(func(key, value interface{}) bool {
		vault := value.(models.TVaultsFromRegistry)
		asMap[vault.Address] = vault
		asSlice = append(asSlice, vault)
		return true
	})

	return asMap, asSlice
}

/**************************************************************************************************
** GetVaultFromRegistry
**************************************************************************************************/
func GetVaultFromRegistry(chainID uint64, vaultAddress common.Address) (models.TVaultsFromRegistry, bool) {
	/**********************************************************************************************
	** We first retrieve the syncMap. This syncMap should be initialized first via the `LoadNewVaultsFromRegistry`
	** function which take the data from the database/badger and store it in it.
	**********************************************************************************************/
	syncMap := _newVaultsFromRegistrySyncMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_newVaultsFromRegistrySyncMap[chainID] = syncMap
	}

	/**********************************************************************************************
	** Here we are trying to load the vault from the syncMap using the vaultAddress. The loaded
	** vault is then type asserted to models.TVaultsFromRegistry. If the vault is not found in the
	** syncMap, we return an empty models.TVaultsFromRegistry and false.
	**********************************************************************************************/
	vault := models.TVaultsFromRegistry{}
	found := false
	syncMap.Range(func(key, value interface{}) bool {
		element := value.(models.TVaultsFromRegistry)
		if element.Address == vaultAddress {
			found = true
			vault = element
			return false
		}
		return true
	})
	if !found {
		return models.TVaultsFromRegistry{}, false
	}
	return vault, true
}
