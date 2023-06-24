package main

import "github.com/yearn/ydaemon/common/env"

var STATUS_FOR_CHAINID = map[uint64]string{}

func init() {
	for _, chainID := range env.SUPPORTED_CHAIN_IDS {
		setStatusForChainID(chainID, "Not Started")
	}
}

func setStatusForChainID(chainID uint64, status string) {
	STATUS_FOR_CHAINID[chainID] = status
}

func getStatusForChainID(chainID uint64) string {
	return STATUS_FOR_CHAINID[chainID]
}
