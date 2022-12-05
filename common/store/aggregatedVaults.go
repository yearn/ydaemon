package store

import (
	"github.com/yearn/ydaemon/common/models"
	"github.com/yearn/ydaemon/common/types/common"
)

type TStore struct {
	// AggregatedVault holds the raw data, aka pure unmodified from the queries, about a vault.
	AggregatedVault map[uint64]map[common.Address]*models.TAggregatedVault
}

// Store holds the data for the partners for each chain.
var Store = TStore{}

func init() {
	// Store.VaultsFromAPIV1 = make(map[uint64]map[common.Address]TLegacyAPI)
	Store.AggregatedVault = make(map[uint64]map[common.Address]*models.TAggregatedVault)
}
