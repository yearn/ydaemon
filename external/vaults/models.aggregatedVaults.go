package vaults

import (
	"github.com/yearn/ydaemon/internal/vaults"
)

var aggregatedVault map[uint64]map[string]*vaults.TAggregatedVault

func init() {
	aggregatedVault = make(map[uint64]map[string]*vaults.TAggregatedVault)
}

func GetAggregatedVault(chainID uint64, address string) (*vaults.TAggregatedVault, bool) {
	if aggregatedVault[chainID] == nil {
		return nil, false
	}
	vault, ok := aggregatedVault[chainID][address]
	return vault, ok
}
