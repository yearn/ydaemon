package prices

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/env"
)

/**********************************************************************************************
** Set of functions to store and retrieve the prices from the cache and/or database and being
** able to access them from the rest of the application.
** The _priceMap variable is not exported and is only used internally by the functions below.
**********************************************************************************************/
var _priceMap = make(map[uint64]*sync.Map)

func initOrGetPriceMap(chainID uint64) *sync.Map {
	syncMap := _priceMap[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		_priceMap[chainID] = syncMap
	}
	return syncMap
}

func init() {
	for _, chainID := range env.SUPPORTED_CHAIN_IDS {
		if _, ok := _priceMap[chainID]; !ok {
			_priceMap[chainID] = &sync.Map{}
		}
	}
}

/**********************************************************************************************
** The _historicalPriceMap variable is not exported and is only used internally by the
** functions below. It is used to store the historical prices of the tokens:
** map[chainID]map[blockNumber]map[tokenAddress]price
**********************************************************************************************/
var _historicalPriceMap = sync.Map{}

/**********************************************************************************************
** MapPrices will, for a given chainID, return a map of prices.
** The map will be of the form: map[vaultAddress]bigNumber.Int
**********************************************************************************************/
func MapPrices(chainID uint64) map[common.Address]*bigNumber.Int {
	syncMap := initOrGetPriceMap(chainID)
	prices := make(map[common.Address]*bigNumber.Int)
	syncMap.Range(func(key, price interface{}) bool {
		prices[key.(common.Address)] = price.(*bigNumber.Int)
		return true
	})
	return prices
}

/**********************************************************************************************
** FindPrice will, for a given chainID, try to find the price of the provided tokenAddress
** stored in _priceMap.
** It will return the price if found, and a boolean indicating if the token was found or not.
**********************************************************************************************/
func FindPrice(chainID uint64, tokenAddress common.Address) (*bigNumber.Int, bool) {
	syncMap := initOrGetPriceMap(chainID)
	price, ok := syncMap.Load(tokenAddress)
	if !ok {
		return nil, false
	}
	return price.(*bigNumber.Int), true
}

func StorePrices(chainID uint64, prices map[common.Address]*bigNumber.Int) {
	syncMap := initOrGetPriceMap(chainID)
	for addr, strategy := range prices {
		syncMap.Store(addr.Hex(), strategy)
	}
}
