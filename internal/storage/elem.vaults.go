package storage

import (
	"database/sql"
	"encoding/json"
	"fmt"
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
var _vaultJSONMutexes = make(map[uint64]*sync.RWMutex)
var _vaultJSONMutexesLock sync.Mutex // Protects access to _vaultJSONMutexes map

/** 🔵 - Yearn *************************************************************************************
** getVaultMutex safely gets or creates a mutex for a specific chainID
**************************************************************************************************/
func getVaultMutex(chainID uint64) *sync.RWMutex {
	_vaultJSONMutexesLock.Lock()
	defer _vaultJSONMutexesLock.Unlock()

	if mutex, exists := _vaultJSONMutexes[chainID]; exists {
		return mutex
	}
	_vaultJSONMutexes[chainID] = &sync.RWMutex{}
	return _vaultJSONMutexes[chainID]
}

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

	retry := 0
	for {
		// Decode the JSON file into the map
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&vaults)
		if err != nil {
			logs.Error("Failed to decode vaults JSON file on chainID " + chainIDStr + ": " + err.Error())
			time.Sleep(1 * time.Second)
			retry++
			if retry > 5 {
				return TJsonVaultStorage{}
			}
			continue
		}
		if retry > 0 {
			logs.Success(`Succeed to decode vaults JSON file on chainID ` + chainIDStr + ` after ` + strconv.Itoa(retry) + ` retries`)
		}
		break
	}

	return vaults
}

/** 🔵 - Yearn *************************************************************************************
** The function `storeVaultsToJson` is responsible for storing vaults to a JSON file.
**************************************************************************************************/
func StoreVaultsToJson(chainID uint64, vaults map[common.Address]models.TVault) {
	mutex := getVaultMutex(chainID)
	mutex.Lock()
	defer mutex.Unlock()

	chainIDStr := strconv.FormatUint(chainID, 10)
	previousVaults := loadVaultsFromJson(chainID)
	version := detectVersionUpdate(chainID, previousVaults.Version, previousVaults.Vaults, vaults)

	allVaults := make(map[common.Address]models.TVault)
	for address, vault := range vaults {
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
** ApplyCmsVaultMeta applies CMS metadata to a vault's metadata field.
** This function updates the vault's metadata with values from the CMS metadata,
** applying field-by-field updates to the vault.Metadata struct.
**
** @param vaultMeta The CMS metadata to apply
** @param vault The vault to update (passed by reference)
**************************************************************************************************/
func ApplyCmsVaultMeta(vaultMeta models.TVaultCmsMetadataSchema, vault *models.TVault) {
	// Apply boolean fields
	vault.Metadata.IsRetired = vaultMeta.IsRetired
	vault.Metadata.IsHighlighted = vaultMeta.IsHighlighted
	vault.Metadata.IsAggregator = vaultMeta.IsAggregator
	vault.Metadata.IsBoosted = vaultMeta.IsBoosted
	vault.Metadata.IsPool = vaultMeta.IsPool
	vault.Metadata.ShouldUseV2APR = vaultMeta.ShouldUseV2APR

	// Apply struct fields
	vault.Metadata.Migration = vaultMeta.Migration
	vault.Metadata.Stability = vaultMeta.Stability

	// Apply string fields (handle nil pointers)
	if vaultMeta.Category != nil {
		vault.Metadata.Category = models.TVaultCategoryType(*vaultMeta.Category)
	}
	if vaultMeta.DisplayName != nil {
		vault.Metadata.DisplayName = *vaultMeta.DisplayName
	}
	if vaultMeta.DisplaySymbol != nil {
		vault.Metadata.DisplaySymbol = *vaultMeta.DisplaySymbol
	}
	if vaultMeta.Description != nil {
		vault.Metadata.Description = *vaultMeta.Description
	}
	if vaultMeta.SourceURI != nil {
		vault.Metadata.SourceURI = *vaultMeta.SourceURI
	}
	if vaultMeta.UINotice != nil {
		vault.Metadata.UINotice = *vaultMeta.UINotice
	}

	// Apply protocols array (convert TCmsProtocolType to string)
	if vaultMeta.Protocols != nil {
		protocols := make([]string, len(vaultMeta.Protocols))
		for i, protocol := range vaultMeta.Protocols {
			protocols[i] = string(protocol)
		}
		vault.Metadata.Protocols = protocols
	}
}

/** 🔵 - Yearn *************************************************************************************
** FetchCmsVaultsMeta fetches vault metadata from the CMS for a specific chain ID.
** The CMS returns an array of vault metadata following the TVaultCmsMetadataSchema structure.
**
** @param chainID The blockchain network ID
** @return map[common.Address]models.TVaultCmsMetadataSchema A mapping of vault addresses to their metadata
**************************************************************************************************/
func FetchCmsVaultsMeta(chainID uint64) map[common.Address]models.TVaultCmsMetadataSchema {
	cmsRoot := env.CMS_ROOT_URL
	var vaultsMetadata []models.TVaultCmsMetadataSchema

	if cmsRoot != "" {
		cmsURL := cmsRoot + "/vaults/" +
			strconv.FormatUint(chainID, 10) + ".json"
		vaultsMetadata = helpers.FetchJSON[[]models.TVaultCmsMetadataSchema](cmsURL)
		logs.Success("Fetch", len(vaultsMetadata), "vault metadata from cms, chain", chainID)

	} else {
		localPath := env.BASE_DATA_PATH + "/cdn-dev/vaults/" + strconv.FormatUint(chainID, 10) + ".json"

		file, err := os.Open(localPath)
		if err != nil {
			logs.Error("Failed to open local CMS file: " + localPath + " - " + err.Error())
			return make(map[common.Address]models.TVaultCmsMetadataSchema)
		}
		defer file.Close()

		decoder := json.NewDecoder(file)
		err = decoder.Decode(&vaultsMetadata)
		if err != nil {
			logs.Error("Failed to decode local CMS file: " + localPath + " - " + err.Error())
			return make(map[common.Address]models.TVaultCmsMetadataSchema)
		}
		logs.Success("Load", len(vaultsMetadata), "vault metadata from local cms, chain", chainID)
	}

	// Convert array to map for easier lookup with normalized addresses
	vaultsMap := make(map[common.Address]models.TVaultCmsMetadataSchema)
	for _, vault := range vaultsMetadata {
		// Normalize address to ensure consistent lookup (case-insensitive)
		normalizedAddress := common.HexToAddress(vault.Address.Hex())
		vaultsMap[normalizedAddress] = vault
	}

	return vaultsMap
}

/**************************************************************************************************
** LoadVaults will retrieve the all the vaults from the configured DB and store them in the
** _vaultsSyncMap for fast access during that same execution.
**************************************************************************************************/
func LoadVaults(chainID uint64, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	mutex := getVaultMutex(chainID)
	mutex.Lock()
	defer mutex.Unlock()

	file := loadVaultsFromJson(chainID)
	_vaultJSONMetadataSyncMap.Store(chainID, TJsonMetadata{
		file.LastUpdate,
		file.Version,
		file.ShouldRefresh,
	})

	meta := FetchCmsVaultsMeta(chainID)

	for _, vault := range file.Vaults {
		// Normalize address to ensure consistent lookup (case-insensitive)
		normalizedAddress := common.HexToAddress(vault.Address.Hex())
		vaultMeta, ok := meta[normalizedAddress]
		if ok {
			ApplyCmsVaultMeta(vaultMeta, &vault)
			// logs.Info("Apply cms vault metadata", chainID, vault.Address)
		}
	}

	for _, vault := range file.Vaults {
		StoreVault(vault.ChainID, vault)
	}
}

/**************************************************************************************************
** StoreVault will add a new vault in the _vaultsSyncMap
**************************************************************************************************/
func StoreVault(chainID uint64, vault models.TVault) {
	chain, ok := env.GetChain(chainID)
	if !ok {
		return
	}
	if helpers.Contains(chain.BlacklistedVaults, vault.Address) {
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

/**************************************************************************************************
** RefreshVaultMetadata will fetch the latest vault metadata from CMS and apply it to all
** currently loaded vaults. This function is designed to be called periodically by the scheduler.
**
** @param chainID The blockchain network ID
**************************************************************************************************/
func RefreshVaultMetadata(chainID uint64) {
	mutex := getVaultMutex(chainID)
	mutex.Lock()
	defer mutex.Unlock()

	meta := FetchCmsVaultsMeta(chainID)

	vaultsMap, _ := ListVaults(chainID)

	for address, vault := range vaultsMap {
		// Normalize address to ensure consistent lookup (case-insensitive)
		normalizedAddress := common.HexToAddress(address.Hex())
		vaultMeta, ok := meta[normalizedAddress]
		if ok {
			ApplyCmsVaultMeta(vaultMeta, &vault)
			StoreVault(chainID, vault)
		}
	}
}

/*****************************************************************************
** Kong Data Storage and Caching
*****************************************************************************/

var _kongVaultDataSyncMap = make(map[uint64]*sync.Map)
var _kongDataMutexes = make(map[uint64]*sync.RWMutex)
var _kongDataMutexesLock sync.Mutex // Protects access to _kongDataMutexes map

/**************************************************************************************************
** safeKongSyncMap ensures thread-safe access to the kong sync map for a given chain ID
**************************************************************************************************/
func safeKongSyncMap(syncMap map[uint64]*sync.Map, chainID uint64) *sync.Map {
	if syncMap[chainID] == nil {
		syncMap[chainID] = &sync.Map{}
	}
	return syncMap[chainID]
}

/**************************************************************************************************
** getKongDataMutex returns a mutex for the given chain ID, creating it if it doesn't exist
**************************************************************************************************/
func getKongDataMutex(chainID uint64) *sync.RWMutex {
	_kongDataMutexesLock.Lock()
	defer _kongDataMutexesLock.Unlock()

	if _kongDataMutexes[chainID] == nil {
		_kongDataMutexes[chainID] = &sync.RWMutex{}
	}
	return _kongDataMutexes[chainID]
}

/**************************************************************************************************
** GetKongVaultData retrieves kong data for a specific vault from the cache
** Uses checksummed address for case-insensitive lookup
**************************************************************************************************/
func GetKongVaultData(chainID uint64, address common.Address) (models.TKongVaultSchema, bool) {
	// Normalize address to ensure consistent lookup (case-insensitive)
	normalizedAddress := common.HexToAddress(address.Hex())
	kongDataFromSyncMap, ok := safeKongSyncMap(_kongVaultDataSyncMap, chainID).Load(normalizedAddress)
	if !ok {
		return models.TKongVaultSchema{}, false
	}
	return kongDataFromSyncMap.(models.TKongVaultSchema), true
}

/**************************************************************************************************
** StoreKongVaultData stores kong data for a specific vault in the cache
** Uses checksummed address for case-insensitive storage
**************************************************************************************************/
func StoreKongVaultData(chainID uint64, address common.Address, kongData models.TKongVaultSchema) {
	// Normalize address to ensure consistent storage (case-insensitive)
	normalizedAddress := common.HexToAddress(address.Hex())
	safeKongSyncMap(_kongVaultDataSyncMap, chainID).Store(normalizedAddress, kongData)
}

/**************************************************************************************************
** ListKongVaultData returns all kong vault data for a specific chain
**************************************************************************************************/
func ListKongVaultData(chainID uint64) map[common.Address]models.TKongVaultSchema {
	kongDataMap := make(map[common.Address]models.TKongVaultSchema)

	safeKongSyncMap(_kongVaultDataSyncMap, chainID).Range(func(key, value interface{}) bool {
		address := key.(common.Address)
		kongData := value.(models.TKongVaultSchema)
		kongDataMap[address] = kongData
		return true
	})

	return kongDataMap
}

/**************************************************************************************************
** fetchKongVaultDataFromDB retrieves vault data from Kong database in batches
**************************************************************************************************/
func fetchKongVaultDataFromDB(chainID uint64) map[common.Address]models.TKongVaultSchema {
	db := GetDB()
	if db == nil {
		logs.Error("Failed to connect to Kong database")
		return make(map[common.Address]models.TKongVaultSchema)
	}

	sqlDB, err := db.DB()
	if err != nil {
		logs.Error("Failed to get raw SQL connection: " + err.Error())
		return make(map[common.Address]models.TKongVaultSchema)
	}

	query := `SELECT snapshot.address, snapshot.snapshot, snapshot.hook 
FROM snapshot 
JOIN thing on snapshot.chain_id = thing.chain_id and snapshot.address = thing.address
WHERE snapshot.chain_id = $1 AND thing.label = 'vault'`
	rows, err := sqlDB.Query(query, chainID)
	if err != nil {
		logs.Error(fmt.Sprintf("Failed to query Kong database for chain %d: %s", chainID, err.Error()))
		return make(map[common.Address]models.TKongVaultSchema)
	}
	defer rows.Close()

	kongVaults := make(map[common.Address]models.TKongVaultSchema)

	for rows.Next() {
		var addressStr string
		var snapshotJSON sql.NullString
		var hookJSON sql.NullString

		if err := rows.Scan(&addressStr, &snapshotJSON, &hookJSON); err != nil {
			logs.Error("Failed to scan Kong row: " + err.Error())
			continue
		}

		address := common.HexToAddress(addressStr)
		kongVault := models.TKongVaultSchema{}

		// Parse snapshot JSON
		if snapshotJSON.Valid {
			if err := json.Unmarshal([]byte(snapshotJSON.String), &kongVault.Snapshot); err != nil {
				logs.Error("Failed to unmarshal snapshot JSON for address " + addressStr + ": " + err.Error())
			}
		}

		// Parse hook JSON
		if hookJSON.Valid {
			if err := json.Unmarshal([]byte(hookJSON.String), &kongVault.Hook); err != nil {
				logs.Error("Failed to unmarshal hook JSON for address " + addressStr + ": " + err.Error())
			}
		}

		kongVaults[address] = kongVault
	}

	if err := rows.Err(); err != nil {
		logs.Error("Error iterating Kong rows: " + err.Error())
	}

	logs.Success(fmt.Sprintf("Fetched %d Kong vault records for chain %d", len(kongVaults), chainID))
	return kongVaults
}

/**************************************************************************************************
** RefreshKongData refreshes vault data with Kong database information for a specific chain
** This function is designed to be called periodically by the scheduler.
**************************************************************************************************/
func RefreshKongData(chainID uint64) {
	mutex := getKongDataMutex(chainID)
	mutex.Lock()
	defer mutex.Unlock()

	logs.Info(fmt.Sprintf("Refreshing Kong vault data for chain %d", chainID))

	kongVaults := fetchKongVaultDataFromDB(chainID)
	if len(kongVaults) == 0 {
		logs.Warning(fmt.Sprintf("No Kong vault data found for chain %d", chainID))
		return
	}

	// Store all kong data in the cache
	for address, kongData := range kongVaults {
		StoreKongVaultData(chainID, address, kongData)
	}

	logs.Success(fmt.Sprintf("Refreshed %d Kong vault records for chain %d", len(kongVaults), chainID))
}
