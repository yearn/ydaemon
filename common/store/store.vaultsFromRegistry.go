package store

import (
	"encoding/json"
	"strconv"
	"sync"

	"github.com/dgraph-io/badger/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/internal/models"
)

var _newVaultsFromRegistrySyncMap = make(map[uint64]*sync.Map)

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

	temp := make(map[string]DBNewVaultsFromRegistry)
	ListFromBadgerDB(chainID, TABLES.VAULTS_FROM_REGISTRY_SYNC, &temp)

	if temp != nil && (len(temp) > 0) {
		for _, vaultFromDB := range temp {
			key := strconv.FormatUint(vaultFromDB.Block, 10) + "_" + addresses.ToAddress(vaultFromDB.RegistryAddress).Hex() + "_" + addresses.ToAddress(vaultFromDB.VaultsAddress).Hex() + "_" + addresses.ToAddress(vaultFromDB.TokenAddress).Hex() + "_" + vaultFromDB.APIVersion
			vaultFromRegistry := models.TVaultsFromRegistry{
				Address:         addresses.ToAddress(vaultFromDB.VaultsAddress),
				RegistryAddress: addresses.ToAddress(vaultFromDB.RegistryAddress),
				TokenAddress:    addresses.ToAddress(vaultFromDB.TokenAddress),
				BlockHash:       common.HexToHash(vaultFromDB.BlockHash),
				Type:            vaultFromDB.Type,
				APIVersion:      vaultFromDB.APIVersion,
				ChainID:         vaultFromDB.ChainID,
				BlockNumber:     vaultFromDB.Block,
				Activation:      vaultFromDB.Activation,
				ManagementFee:   vaultFromDB.ManagementFee,
				TxIndex:         vaultFromDB.TxIndex,
				LogIndex:        vaultFromDB.LogIndex,
			}
			syncMap.Store(key, vaultFromRegistry)
		}
	}
}

/**************************************************************************************************
** AppendToNewVaultsFromRegistry will add a new vault in the _vaultsSyncMap
**************************************************************************************************/
func AppendToNewVaultsFromRegistry(chainID uint64, vault models.TVaultsFromRegistry) {
	syncMap := _newVaultsFromRegistrySyncMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_newVaultsFromRegistrySyncMap[chainID] = syncMap
	}

	key := strconv.FormatUint(vault.BlockNumber, 10) + "_" + vault.RegistryAddress.Hex() + "_" + vault.Address.Hex() + "_" + vault.TokenAddress.Hex() + "_" + vault.APIVersion
	syncMap.Store(key, vault)
}

/**************************************************************************************************
** StoreNewVaultsFromRegistry will store a new vault in the _newVaultsFromRegistrySyncMap for fast
** access during that same execution, and will store it in the local DB for future executions.
** We are using the local DB because we don't want to trust the shared DB for this data.
**************************************************************************************************/
func StoreNewVaultsFromRegistry(chainID uint64, vault models.TVaultsFromRegistry) {
	AppendToNewVaultsFromRegistry(chainID, vault)

	key := strconv.FormatUint(vault.BlockNumber, 10) + "_" + vault.RegistryAddress.Hex() + "_" + vault.Address.Hex() + "_" + vault.TokenAddress.Hex() + "_" + vault.APIVersion

	go OpenBadgerDB(chainID, TABLES.VAULTS_FROM_REGISTRY_SYNC).Update(func(txn *badger.Txn) error {
		DBbaseSchema := DBBaseSchema{
			UUID:    getUUID(key),
			Block:   vault.BlockNumber,
			ChainID: chainID,
		}
		data := &DBNewVaultsFromRegistry{
			DBbaseSchema,
			vault.RegistryAddress.Hex(),
			vault.Address.Hex(),
			vault.TokenAddress.Hex(),
			vault.BlockHash.Hex(),
			vault.APIVersion,
			vault.Activation,
			vault.ManagementFee,
			vault.TxIndex,
			vault.LogIndex,
			vault.Type,
		}
		dataByte, err := json.Marshal(data)
		if err != nil {
			return err
		}
		return txn.Set([]byte(key), dataByte)
	})
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
