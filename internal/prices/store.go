package prices

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type TStore struct {
	// VaultPricePerShare holds the details about the price for one share of a vault, based on a multicall
	VaultPricePerShare map[uint64]map[common.Address]*big.Int

	// TokenPrices contains, for each network, the prices for each token we fetched
	// the prices for. Price is accessible through TokenPrices[chainID][tokenAddress]
	TokenPrices map[uint64]map[common.Address]*big.Int
}

// Store holds the data for the partners for each chain.
var Store = TStore{}

func init() {
	Store.VaultPricePerShare = make(map[uint64]map[common.Address]*big.Int)
	Store.TokenPrices = make(map[uint64]map[common.Address]*big.Int)
}
