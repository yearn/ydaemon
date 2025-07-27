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

	// Convert array to map for easier lookup
	vaultsMap := make(map[common.Address]models.TVaultCmsMetadataSchema)
	for _, vault := range vaultsMetadata {
		vaultsMap[vault.Address] = vault
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
		vaultMeta, ok := meta[vault.Address]
		if ok {
			ApplyCmsVaultMeta(vaultMeta, &vault)
			logs.Info("Apply cms vault metadata", chainID, vault.Address)
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
		vaultMeta, ok := meta[address]
		if ok {
			ApplyCmsVaultMeta(vaultMeta, &vault)
			StoreVault(chainID, vault)
		}
	}
}
