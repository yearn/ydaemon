package store

import (
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/env"
)

var _historicalPriceSyncMap = make(map[uint64]*sync.Map)

/**************************************************************************************************
** The init function is a special function triggered directly on execution of the package.
** It is used to initialize the package.
** This init is responsible of loading the store
***************************************************************************************************/
func init() {
	wg := &sync.WaitGroup{}
	for _, chainID := range env.SUPPORTED_CHAIN_IDS {
		wg.Add(1)
		if _historicalPriceSyncMap == nil {
			_historicalPriceSyncMap = make(map[uint64]*sync.Map)
		}
		if _historicalPriceSyncMap[chainID] == nil {
			_historicalPriceSyncMap[chainID] = &sync.Map{}
		}

		go LoadHistoricalPrice(chainID, wg)
	}
	wg.Wait()
}

/**************************************************************************************************
** LoadHistoricalPrice will retrieve the stored values from the local DB and load them in the
** _historicalPriceSyncMap.
**************************************************************************************************/
func LoadHistoricalPrice(chainID uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	temp := make(map[string]string)
	Iterate(chainID, TABLES.HISTORICAL_PRICES, &temp)
	if temp != nil && (len(temp) > 0) {
		for k, v := range temp {
			_historicalPriceSyncMap[chainID].Store(k, bigNumber.NewInt(0).SetString(v))
		}
	}
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
** SaveBlockPrice will save a tokenPrice in the _historicalPriceSyncMap.
**************************************************************************************************/
func SaveBlockPrice(chainID uint64, blockNumber uint64, tokenAddress common.Address, price *bigNumber.Int) {
	syncMap := _historicalPriceSyncMap[chainID]
	key := strconv.FormatUint(blockNumber, 10) + "_" + tokenAddress.Hex()

	syncMap.Store(key, price)
	SaveInDB(chainID, TABLES.HISTORICAL_PRICES, key, price.String())
}
