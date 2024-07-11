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
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
)

var _vaultsSyncMap = make(map[uint64]*sync.Map)
var _vaultJSONMetadataSyncMap = sync.Map{}

/** 🔵 - Yearn *************************************************************************************
** The function `loadVaultsFromJson` is responsible for loading vaults from a JSON file.
**************************************************************************************************/
func loadVaultsFromJson(chainID uint64) TJsonVaultStorage {
	var vaults TJsonVaultStorage
	chainIDStr := strconv.FormatUint(chainID, 10)

	// Load the JSON file
	file, err := os.Open(env.BASE_DATA_PATH + "/meta/vaults/" + chainIDStr + ".json")
	if err != nil {
		return TJsonVaultStorage{}
	}
	defer file.Close()

	// Decode the JSON file into the map
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&vaults)
	if err != nil {
		logs.Error("Failed to decode vaults JSON file: " + err.Error())
		return TJsonVaultStorage{}
	}

	return vaults
}

/** 🔵 - Yearn *************************************************************************************
** The function `storeVaultsToJson` is responsible for storing vaults to a JSON file.
**************************************************************************************************/
func StoreVaultsToJson(chainID uint64, vaults map[common.Address]models.TVault) {
	chainIDStr := strconv.FormatUint(chainID, 10)
	previousVaults := loadVaultsFromJson(chainID)
	version := detectVersionUpdate(chainID, previousVaults.Version, previousVaults.Vaults, vaults)

	allVaults := make(map[common.Address]models.TVault)
	for address, vault := range vaults {
		// @dev use this place if you need to update the json
		allVaults[address] = vault
	}

	data := TJsonVaultStorage{
		TJsonMetadata: TJsonMetadata{
			LastUpdate:    time.Now(),
			Version:       version,
			ShouldRefresh: false,
		},
		Vaults: allVaults,
	}
	_vaultJSONMetadataSyncMap.Store(chainID, TJsonMetadata{
		data.LastUpdate,
		data.Version,
		data.ShouldRefresh,
	})

	file, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		logs.Error("Failed to marshal vaults JSON file: " + err.Error())
		return
	}
	if _, err := os.Stat(env.BASE_DATA_PATH + "/meta/vaults"); os.IsNotExist(err) {
		os.MkdirAll(env.BASE_DATA_PATH+"/meta/vaults", 0755)
	}
	err = os.WriteFile(env.BASE_DATA_PATH+"/meta/vaults/"+chainIDStr+".json", file, 0644)
	if err != nil {
		logs.Error("Failed to write vaults JSON file: " + err.Error())
	}
}

/**************************************************************************************************
** Retrieve the last time the ERC20s were updated for a specific chainID
**************************************************************************************************/
func GetVaultsJsonMetadata(chainID uint64) TJsonMetadata {
	if jsonMetadata, ok := _vaultJSONMetadataSyncMap.Load(chainID); ok {
		return jsonMetadata.(TJsonMetadata)
	}
	return TJsonMetadata{}
}

/**************************************************************************************************
** LoadVaults will retrieve the all the vaults from the configured DB and store them in the
** _vaultsSyncMap for fast access during that same execution.
**************************************************************************************************/
func LoadVaults(chainID uint64, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}

	file := loadVaultsFromJson(chainID)
	_vaultJSONMetadataSyncMap.Store(chainID, TJsonMetadata{
		file.LastUpdate,
		file.Version,
		file.ShouldRefresh,
	})
	for _, vault := range file.Vaults {
		StoreVault(chainID, vault)
	}
}

/**************************************************************************************************
** StoreVault will add a new vault in the _vaultsSyncMap
**************************************************************************************************/
func StoreVault(chainID uint64, vault models.TVault) {
	if helpers.Contains(env.CHAINS[chainID].BlacklistedVaults, vault.Address) {
		return
	}
	safeSyncMap(_vaultsSyncMap, chainID).Store(vault.Address, vault)
}

/**************************************************************************************************
** ListVault will return a list of all the vaults stored in the caching system for a given
** chainID. Both a map and a slice are returned.
**************************************************************************************************/
func ListVaults(chainID uint64) (asMap map[common.Address]models.TVault, asSlice []models.TVault) {
	asMap = make(map[common.Address]models.TVault) // make to avoid nil map

	/**********************************************************************************************
	** We can just iterate over the syncMap and add the vaults to the map and slice.
	** As the stored vault data are only a subset of static, we need to use the actual structure
	**********************************************************************************************/
	safeSyncMap(_vaultsSyncMap, chainID).Range(func(key, value interface{}) bool {
		vault := value.(models.TVault)
		if addresses.Equals(vault.Address, common.Address{}) {
			return true
		}
		asMap[vault.Address] = vault
		asSlice = append(asSlice, vault)
		return true
	})

	return asMap, asSlice
}

/**************************************************************************************************
** GetVault will return a single vault stored in the caching system for a given pair of chainID
** and vault address.
**************************************************************************************************/
func GetVault(chainID uint64, vaultAddress common.Address) (models.TVault, bool) {
	vaultFromSyncMap, ok := safeSyncMap(_vaultsSyncMap, chainID).Load(vaultAddress)
	if !ok {
		return models.TVault{}, false
	}
	return vaultFromSyncMap.(models.TVault), true
}
