package store

import (
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
)

var _historicalPriceSyncMap = make(map[uint64]*sync.Map)

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
