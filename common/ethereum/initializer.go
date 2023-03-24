package ethereum

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/logs"
)

/**************************************************************************************************
** The init function is a special function triggered directly on execution of the package.
** It is used to initialize the package.
** This init is responsible of creating the RPC clients for all the chains supported by yDaemon
** and storing them in the RPC map.
** Then it will create the multicall clients for each chain.
** Then, it will init the blockTimeSyncMap for all the chains.
***************************************************************************************************/
func init() {
	// Create the RPC client for all the chains supported by yDaemon
	for _, chainID := range env.SUPPORTED_CHAIN_IDS {
		client, err := ethclient.Dial(GetRPCURI(chainID))
		if err != nil {
			logs.Error(err, "Failed to connect to node")
			continue
		}
		RPC[chainID] = client
	}

	// Create the multicall client for all the chains supported by yDaemon
	for _, chainID := range env.SUPPORTED_CHAIN_IDS {
		MulticallClientForChainID[chainID] = NewMulticall(GetRPCURI(chainID), env.MULTICALL_ADDRESSES[chainID])
	}
}
