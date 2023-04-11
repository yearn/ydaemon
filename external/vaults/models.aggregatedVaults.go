package vaults

import "github.com/yearn/ydaemon/internal/models"

var aggregatedVault map[uint64]map[string]*models.TAggregatedVault

func init() {
	aggregatedVault = make(map[uint64]map[string]*models.TAggregatedVault)
}

func GetAggregatedVault(chainID uint64, address string) (*models.TAggregatedVault, bool) {
	if aggregatedVault[chainID] == nil {
		return nil, false
	}
	vault, ok := aggregatedVault[chainID][address]
	return vault, ok
}
