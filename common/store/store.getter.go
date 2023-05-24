package store

import (
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/internal/models"
)

/**************************************************************************************************
** GetBlockTime will try, for a specific blockNumber on a specific chain, to find its execution
** timestamp in the _blockTimeSyncMap.
**************************************************************************************************/
func GetBlockTime(chainID uint64, blockNumber uint64) (blockTime uint64, ok bool) {
	syncMap := _blockTimeSyncMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_blockTimeSyncMap[chainID] = syncMap
	}

	blockTimeData, ok := syncMap.Load(blockNumber)
	if !ok {
		return 0, false
	}

	return blockTimeData.(uint64), true
}

/**************************************************************************************************
** GetBlockPrice will try, for a specific blockNumber on a specific chain, the price of the
** provided token address.
** If the price is missing, this will try to fetch it via the lens oracle contract.
**************************************************************************************************/
func GetBlockPrice(chainID uint64, blockNumber uint64, tokenAddress common.Address) (price *bigNumber.Int, ok bool) {
	key := strconv.FormatUint(blockNumber, 10) + "_" + tokenAddress.Hex()
	tokenPrice, ok := _historicalPriceSyncMap[chainID].Load(key)
	if !ok {
		return bigNumber.NewInt(0), false
	}
	return tokenPrice.(*bigNumber.Int), true
}

/**************************************************************************************************
** GetERC20 will try, for a specific chain, to find the provided token address as ERC20 struct.
**************************************************************************************************/
func GetERC20(chainID uint64, tokenAddress common.Address) (token models.TERC20Token, ok bool) {
	key := tokenAddress.Hex()
	tokenFromSyncMap, ok := _erc20SyncMap[chainID].Load(key)
	if !ok {
		return models.TERC20Token{}, false
	}
	return tokenFromSyncMap.(models.TERC20Token), true
}

/**************************************************************************************************
** ListAllERC20 will return a list of all the ERC20 stored in the caching system for a given
** chainID. Both a map and a slice are returned.
**************************************************************************************************/
func ListAllERC20(chainID uint64) (asMap map[common.Address]models.TERC20Token, asSlice []models.TERC20Token) {
	asMap = make(map[common.Address]models.TERC20Token) // make to avoid nil map

	/**********************************************************************************************
	** We first retrieve the syncMap. This syncMap should be initialized first via the `LoadERC20`
	** function which take the data from the database/badger and store it in it.
	**********************************************************************************************/
	syncMap := _erc20SyncMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_erc20SyncMap[chainID] = syncMap
	}

	/**********************************************************************************************
	** We can just iterate over the syncMap and add the vaults to the map and slice.
	** As the stored vault data are only a subset of static, we need to use the actual structure
	** and not the DBVault one.
	**********************************************************************************************/
	syncMap.Range(func(key, value interface{}) bool {
		token := value.(models.TERC20Token)
		asMap[token.Address] = token
		asSlice = append(asSlice, token)
		return true
	})

	return asMap, asSlice
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
** ListAllVault will return a list of all the vaults stored in the caching system for a given
** chainID. Both a map and a slice are returned.
**************************************************************************************************/
func ListAllVaults(chainID uint64) (asMap map[common.Address]models.TVault, asSlice []models.TVault) {
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

/**************************************************************************************************
** ListAllStrategies will return a list of all the strategies stored in the caching system for a
** given chainID. Both a map and a slice are returned.
** Note: It's for the TStrategyAdded structure.
**************************************************************************************************/
func ListAllStrategiess(chainID uint64) (
	asMap map[common.Address]models.TStrategyAdded,
	asSlice []models.TStrategyAdded,
) {
	asMap = make(map[common.Address]models.TStrategyAdded) // make to avoid nil map

	/**********************************************************************************************
	** We first retrieve the syncMap. This syncMap should be initialized first via the
	** `LoadStrategies` function which take the data from the database/badger and store it in it.
	**********************************************************************************************/
	syncMap := _strategiesSyncMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_strategiesSyncMap[chainID] = syncMap
	}

	/**********************************************************************************************
	** We can just iterate over the syncMap and add the strategies to the map and slice.
	** As the stored vault data are only a subset of static, we need to use the actual structure
	** and not the DBVault one.
	**********************************************************************************************/
	syncMap.Range(func(key, value interface{}) bool {
		vault := value.(models.TStrategyAdded)
		asMap[vault.StrategyAddress] = vault
		asSlice = append(asSlice, vault)
		return true
	})

	return asMap, asSlice
}
