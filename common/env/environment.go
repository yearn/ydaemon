package env

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/yearn/ydaemon/common/logs"
)

// SetEnv will init the environment variables based on the .env file
func SetEnv(path string) {
	_ = godotenv.Load(path)

	RPCURIFor1, exists := os.LookupEnv("RPC_URI_FOR_1")
	if !exists {
		logs.Warning("RPC_URI_FOR_1 not set, using default value: [https://eth.public-rpc.com]")
	} else {
		RPC_ENDPOINTS[1] = RPCURIFor1
	}

	RPCURIFor10, exists := os.LookupEnv("RPC_URI_FOR_10")
	if !exists {
		logs.Warning("RPC_URI_FOR_10 not set, using default value: [https://mainnet.optimism.io]")
	} else {
		RPC_ENDPOINTS[10] = RPCURIFor10
	}

	RPCURIFor250, exists := os.LookupEnv("RPC_URI_FOR_250")
	if !exists {
		logs.Warning("RPC_URI_FOR_250 not set, using default value: [https://rpc.ftm.tools]")
	} else {
		RPC_ENDPOINTS[250] = RPCURIFor250
	}

	RPCURIFor42161, exists := os.LookupEnv("RPC_URI_FOR_42161")
	if !exists {
		logs.Warning("RPC_URI_FOR_42161 not set, using default value: [https://arbitrum.public-rpc.com]")
	} else {
		RPC_ENDPOINTS[42161] = RPCURIFor42161
	}

	GraphAPIURI, exists := os.LookupEnv("GRAPH_API_URI")
	if !exists {
		logs.Warning("GRAPH_API_URI not set, using default value: [https://api.thegraph.com/subgraphs/name/rareweasel/yearn-vaults-v2-subgraph-mainnet]")
	} else {
		THEGRAPH_ENDPOINTS[1] = GraphAPIURI
	}
}

func init() {
	SetEnv(`.env`)
}
