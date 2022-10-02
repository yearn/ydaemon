package vaults

import (
	"github.com/yearn/ydaemon/internal/utils/models"
	"github.com/yearn/ydaemon/internal/utils/types/common"
)

type TStore struct {
	// VaultsFromAPIV1 holds the data about the Vaults from the Yearn API for each chain.
	VaultsFromAPIV1 map[uint64]map[common.Address]models.TAPIV1Vault
}

// Store holds the data for the partners for each chain.
var Store = TStore{}

func init() {
	Store.VaultsFromAPIV1 = make(map[uint64]map[common.Address]models.TAPIV1Vault)
}
