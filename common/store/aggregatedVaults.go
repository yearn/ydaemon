package store

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/models"
)

type TStore struct {
	// AggregatedVault holds the raw data, aka pure unmodified from the queries, about a vault.
	AggregatedVault map[uint64]map[common.MixedcaseAddress]*models.TAggregatedVault
}

// Store holds the data for the partners for each chain.
var Store = TStore{}

func init() {
	// Store.VaultsFromAPIV1 = make(map[uint64]map[common.MixedcaseAddress]TLegacyAPI)
	Store.AggregatedVault = make(map[uint64]map[common.MixedcaseAddress]*models.TAggregatedVault)
}
