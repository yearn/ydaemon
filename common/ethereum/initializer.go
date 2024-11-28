package ethereum

import (
	"os"
	"strconv"

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
func Initialize() {
	// Create the RPC client for all the chains supported by yDaemon
	for _, chain := range env.GetChains() {
		client, err := ethclient.Dial(GetRPCURI(chain.ID))
		if err != nil {
			logs.Error(err, "Failed to connect to node")
			continue
		}
		RPC[chain.ID] = client
	}

	// Create the multicall client for all the chains supported by yDaemon
	for _, chain := range env.GetChains() {
		rpcToUse := GetRPCURI(chain.ID)
		multiCallURI, exists := os.LookupEnv("MULTICALL_RPC_URI_FOR_" + strconv.FormatUint(chain.ID, 10))
		if exists {
			rpcToUse = multiCallURI
		}

		MulticallClientForChainID[chain.ID] = NewMulticall(
			rpcToUse,
			chain.MulticallContract.Address,
		)
	}
}
