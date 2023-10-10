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

var _vaultsSyncMap = make(map[uint64]*sync.Map)

/** 🔵 - Yearn *************************************************************************************
** The function `loadVaultsFromJson` is responsible for loading vaults from a JSON file.
**************************************************************************************************/
func LoadVaultsFromJson(chainID uint64) map[common.Address]models.TVault {
	var vaults map[common.Address]models.TVault
	chainIDStr := strconv.FormatUint(chainID, 10)

	// Load the JSON file
	file, err := os.Open(env.BASE_DATA_PATH + "/meta/store/vaults/" + chainIDStr + ".json")
	if err != nil {
		return nil
	}
	defer file.Close()

	// Decode the JSON file into the map
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&vaults)
	if err != nil {
		logs.Error("Failed to decode vaults JSON file: " + err.Error())
		return nil
	}

	return vaults
}

/** 🔵 - Yearn *************************************************************************************
** The function `storeVaultsToJson` is responsible for storing vaults to a JSON file.
**************************************************************************************************/
func StoreVaultsToJson(chainID uint64, vaults map[common.Address]models.TVault) {
	chainIDStr := strconv.FormatUint(chainID, 10)

	file, _ := json.MarshalIndent(vaults, "", "\t")
	if _, err := os.Stat(env.BASE_DATA_PATH + "/meta/store/vaults"); os.IsNotExist(err) {
		os.MkdirAll(env.BASE_DATA_PATH+"/meta/store/vaults", 0755)
	}
	err := os.WriteFile(env.BASE_DATA_PATH+"/meta/store/vaults/"+chainIDStr+".json", file, 0644)
	if err != nil {
		logs.Error("Failed to write vaults JSON file: " + err.Error())
	}
}

/**************************************************************************************************
** LoadVaults will retrieve the all the vaults from the configured DB and store them in the
** _vaultsSyncMap for fast access during that same execution.
**************************************************************************************************/
func LoadVaults(chainID uint64, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	syncMap := _vaultsSyncMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_vaultsSyncMap[chainID] = syncMap
	}
	vaultMap := LoadVaultsFromJson(chainID)
	for _, vault := range vaultMap {
		syncMap.Store(vault.Address.Hex(), vault)
	}
}

/**************************************************************************************************
** StoreVault will add a new vault in the _vaultsSyncMap
**************************************************************************************************/
func StoreVault(chainID uint64, vault models.TVault) {
	syncMap := _vaultsSyncMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_vaultsSyncMap[chainID] = syncMap
	}
	syncMap.Store(vault.Address, vault)
}

/**************************************************************************************************
** ListVault will return a list of all the vaults stored in the caching system for a given
** chainID. Both a map and a slice are returned.
**************************************************************************************************/
func ListVaults(chainID uint64) (asMap map[common.Address]models.TVault, asSlice []models.TVault) {
	asMap = make(map[common.Address]models.TVault) // make to avoid nil map

	/**********************************************************************************************
	** We first retrieve the syncMap. This syncMap should be initialized first via the `LoadVaults`
	** function which take the data from the database/badger and store it in it.
	**********************************************************************************************/
	syncMap := _vaultsSyncMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_vaultsSyncMap[chainID] = syncMap
	}

	/**********************************************************************************************
	** We can just iterate over the syncMap and add the vaults to the map and slice.
	** As the stored vault data are only a subset of static, we need to use the actual structure
	** and not the DBVault one.
	**********************************************************************************************/
	syncMap.Range(func(key, value interface{}) bool {
		vault := value.(models.TVault)
		asMap[vault.Address] = vault
		asSlice = append(asSlice, vault)
		return true
	})

	return asMap, asSlice
}
