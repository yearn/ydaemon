package main

import (
	"sync"

	"github.com/yearn/ydaemon/common/env"
)

/**************************************************************************************************
** statusStore is a thread-safe map for storing chain status information
** It uses sync.Map to handle concurrent access safely
**************************************************************************************************/
var (
	statusStore sync.Map
)

/**************************************************************************************************
** init initializes the status store for all configured chains
** Sets initial status as "Not Started" for each chain
**************************************************************************************************/
func init() {
	for chainID := range env.GetChains() {
		setStatusForChainID(chainID, "Not Started")
	}
}

/**************************************************************************************************
** setStatusForChainID updates the status for a specific chain ID
** Thread-safe operation using sync.Map.Store
**
** @param chainID uint64 - The chain ID to update
** @param status string - The new status to set
**************************************************************************************************/
func setStatusForChainID(chainID uint64, status string) {
	statusStore.Store(chainID, status)
}

/**************************************************************************************************
** getStatusForChainID retrieves the status for a specific chain ID
** Thread-safe operation using sync.Map.Load
**
** @param chainID uint64 - The chain ID to query
** @return string - The current status of the chain
**************************************************************************************************/
func getStatusForChainID(chainID uint64) string {
	if status, ok := statusStore.Load(chainID); ok {
		return status.(string)
	}
	return "Not Started"
}
