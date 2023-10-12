package storage

import (
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
)

var _pricesSyncMap = make(map[uint64]*sync.Map)

/**************************************************************************************************
** StorePrice will add a new price in the _pricesSyncMap
**************************************************************************************************/
func StorePrice(chainID uint64, price models.TPrices) {
	safeSyncMap(_pricesSyncMap, chainID).Store(price.Address, price)
}

/**************************************************************************************************
** ListVault will return a list of all the prices stored in the caching system for a given
** chainID. Both a map and a slice are returned.
**************************************************************************************************/
func ListPrices(chainID uint64) (
	asMap map[common.Address]models.TPrices,
	asSlice []models.TPrices,
) {
	asMap = make(map[common.Address]models.TPrices) // make to avoid nil map

	/**********************************************************************************************
	** We can just iterate over the syncMap and add the prices to the map and slice.
	** As the stored price data are only a subset of static, we need to use the actual structure
	**********************************************************************************************/
	safeSyncMap(_pricesSyncMap, chainID).Range(func(key, value interface{}) bool {
		price := value.(models.TPrices)
		asMap[price.Address] = price
		asSlice = append(asSlice, price)
		return true
	})

	return asMap, asSlice
}

/**************************************************************************************************
** GetPrice will return a single price stored in the caching system for a given pair of chainID
** and price address.
**************************************************************************************************/
func GetPrice(chainID uint64, tokenAddress common.Address) (models.TPrices, bool) {
	vaultFromSyncMap, ok := safeSyncMap(_pricesSyncMap, chainID).Load(tokenAddress)
	if !ok {
		logs.Warning(`unable to find price for token ` + tokenAddress.Hex() + ` on chain ` + strconv.FormatUint(chainID, 10))
		return models.TPrices{}, false
	}
	return vaultFromSyncMap.(models.TPrices), true
}
