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
** The function `loadRegistriesFromJson` is responsible for loading historical vaults from
** a JSON file. It reads the JSON file, parses the data into a map of vault addresses to
** `TVaultsFromRegistry` objects, and returns this map along with the highest block number found
** among all vaults. This function is used to initialize the state of the vaults from a previously
** saved state.
**************************************************************************************************/
func loadRegistriesFromJson(chainID uint64) (map[common.Address]models.TVaultsFromRegistry, uint64) {
	var historicalVaults map[common.Address]models.TVaultsFromRegistry
	var highestBlockNumber uint64
	chainIDStr := strconv.FormatUint(chainID, 10)

	// Load the JSON file
	file, err := os.Open(env.BASE_DATA_PATH + "/meta/registries/" + chainIDStr + ".json")
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
** The function `storeRegistriesToJson` is responsible for storing historical vaults to a
** JSON file. It takes a map of vault addresses to `TVaultsFromRegistry` objects, and writes this
** map to a JSON file. This function is used to save the state of the vaults for later use.
**************************************************************************************************/
func StoreRegistriesToJson(chainID uint64, historicalVaults map[common.Address]models.TVaultsFromRegistry) {
	chainIDStr := strconv.FormatUint(chainID, 10)

	file, _ := json.MarshalIndent(historicalVaults, "", "\t")
	if _, err := os.Stat(env.BASE_DATA_PATH + "/meta/registries"); os.IsNotExist(err) {
		os.MkdirAll(env.BASE_DATA_PATH+"/meta/registries", 0755)
	}
	err := os.WriteFile(env.BASE_DATA_PATH+"/meta/registries/"+chainIDStr+".json", file, 0644)
	if err != nil {
		logs.Error("Failed to write vaults JSON file: " + err.Error())
	}
}

/**************************************************************************************************
** LoadRegistries will retrieve the all the vaults added to the registries from the
** local DB and store them in the _newVaultsFromRegistrySyncMap for fast access during that same
** execution.
** Use local DB to not screw up the DB with the same data over and over again.
**************************************************************************************************/
func LoadRegistries(chainID uint64, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	historicalVaultsMap, _ := loadRegistriesFromJson(chainID)
	for _, vault := range historicalVaultsMap {
		safeSyncMap(_newVaultsFromRegistrySyncMap, chainID).Store(vault.Address, vault)
	}
}

/**************************************************************************************************
** StoreNewVaultToRegistry will add a new vault in the _vaultsSyncMap
**************************************************************************************************/
func StoreNewVaultToRegistry(chainID uint64, vault models.TVaultsFromRegistry) {
	safeSyncMap(_newVaultsFromRegistrySyncMap, chainID).Store(vault.Address, vault)
}

/**************************************************************************************************
** ListVaultsFromRegistries will retrieve the all the vaults added to the registries from the
** given chainID
**************************************************************************************************/
func ListVaultsFromRegistries(chainID uint64) (asMap map[common.Address]models.TVaultsFromRegistry, asSlice []models.TVaultsFromRegistry) {
	asMap = make(map[common.Address]models.TVaultsFromRegistry) // make to avoid nil map

	/**********************************************************************************************
	** We can just iterate over the syncMap and add the vaults to the map and slice.
	** As the stored vault data are only a subset of static, we need to use the actual structure
	**********************************************************************************************/
	safeSyncMap(_newVaultsFromRegistrySyncMap, chainID).Range(func(key, value interface{}) bool {
		vault := value.(models.TVaultsFromRegistry)
		asMap[vault.Address] = vault
		asSlice = append(asSlice, vault)
		return true
	})

	return asMap, asSlice
}

/**************************************************************************************************
** ListVaultsFromRegistry will return a list of all the vaults stored in the caching system for a
** given chainID and registry. Both a map and a slice are returned.
** This is for the registry vault version.
**************************************************************************************************/
func ListVaultsFromRegistry(chainID uint64, registryAddress common.Address) (asMap map[common.Address]models.TVaultsFromRegistry, asSlice []models.TVaultsFromRegistry) {
	asMap = make(map[common.Address]models.TVaultsFromRegistry) // make to avoid nil map

	/**********************************************************************************************
	** We can just iterate over the syncMap and add the vaults to the map and slice.
	** As the stored vault data are only a subset of static, we need to use the actual structure
	**********************************************************************************************/
	safeSyncMap(_newVaultsFromRegistrySyncMap, chainID).Range(func(key, value interface{}) bool {
		vault := value.(models.TVaultsFromRegistry)
		if vault.RegistryAddress != registryAddress {
			return true
		}
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
	** Here we are trying to load the vault from the syncMap using the vaultAddress. The loaded
	** vault is then type asserted to models.TVaultsFromRegistry. If the vault is not found in the
	** syncMap, we return an empty models.TVaultsFromRegistry and false.
	**********************************************************************************************/
	vault := models.TVaultsFromRegistry{}
	found := false
	safeSyncMap(_newVaultsFromRegistrySyncMap, chainID).Range(func(key, value interface{}) bool {
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
