package prices

import (
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/types/common"
)

type TStore struct {
	// TokenPrices contains, for each network, the prices for each token we fetched
	// the prices for. Price is accessible through TokenPrices[chainID][tokenAddress]
	TokenPrices map[uint64]map[common.Address]*bigNumber.Int
}

// Store holds the data for the partners for each chain.
var Store = TStore{}

func init() {
	Store.TokenPrices = make(map[uint64]map[common.Address]*bigNumber.Int)
}
