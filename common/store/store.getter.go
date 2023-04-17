package store

import (
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
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
** ListLastNewVaultEventForRegistries will return a map of the last blockNumber for each registry that
** has been registered. This will allow us to know when to start listening for new events instead
** of going through the whole history.
**************************************************************************************************/
func ListLastNewVaultEventForRegistries(chainID uint64) map[string]uint64 {
	lastRegisteredEventForRegistry := make(map[string]uint64)
	registryForChainID := _newVaultsFromRegistrySyncMap[chainID]
	if registryForChainID == nil {
		registryForChainID = &sync.Map{}
		_newVaultsFromRegistrySyncMap[chainID] = registryForChainID
	}
	registryForChainID.Range(func(key, value interface{}) bool {
		registry := value.(DBNewVaultsFromRegistry).RegistryAddress
		blockNumber := value.(DBNewVaultsFromRegistry).Block

		if _, ok := lastRegisteredEventForRegistry[registry]; !ok {
			lastRegisteredEventForRegistry[registry] = blockNumber
		} else if lastRegisteredEventForRegistry[registry] < blockNumber {
			lastRegisteredEventForRegistry[registry] = blockNumber
		}
		return true
	})

	return lastRegisteredEventForRegistry
}
