package env

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/yearn/ydaemon/common/logs"
)

// SetEnv will init the environment variables based on the .env file
func SetEnv(path string) {
	godotenv.Load(path)

	RPCURIFor1, exists := os.LookupEnv("RPC_URI_FOR_1")
	if !exists {
		logs.Debug("RPC_URI_FOR_1 not set, using default value: [https://eth.public-rpc.com]")
	} else {
		chain := CHAINS[1]
		chain.RpcURI = RPCURIFor1
		CHAINS[1] = chain
	}

	RPCURIFor10, exists := os.LookupEnv("RPC_URI_FOR_10")
	if !exists {
		logs.Debug("RPC_URI_FOR_10 not set, using default value: [https://mainnet.optimism.io]")
	} else {
		chain := CHAINS[10]
		chain.RpcURI = RPCURIFor10
		CHAINS[10] = chain
	}

	RPCURIFor137, exists := os.LookupEnv("RPC_URI_FOR_137")
	if !exists {
		logs.Debug("RPC_URI_FOR_137 not set, using default value: [https://polygon.llamarpc.com]")
	} else {
		chain := CHAINS[137]
		chain.RpcURI = RPCURIFor137
		CHAINS[137] = chain
	}

	RPCURIFor250, exists := os.LookupEnv("RPC_URI_FOR_250")
	if !exists {
		logs.Debug("RPC_URI_FOR_250 not set, using default value: [https://rpc.ftm.tools]")
	} else {
		chain := CHAINS[250]
		chain.RpcURI = RPCURIFor250
		CHAINS[250] = chain
	}

	RPCURIFor8453, exists := os.LookupEnv("RPC_URI_FOR_8453")
	if !exists {
		logs.Debug("RPC_URI_FOR_8453 not set, using default value: [https://developer-access-mainnet.base.org]")
	} else {
		chain := CHAINS[8453]
		chain.RpcURI = RPCURIFor8453
		CHAINS[8453] = chain
	}

	RPCURIFor42161, exists := os.LookupEnv("RPC_URI_FOR_42161")
	if !exists {
		logs.Debug("RPC_URI_FOR_42161 not set, using default value: [https://arbitrum.public-rpc.com]")
	} else {
		chain := CHAINS[42161]
		chain.RpcURI = RPCURIFor42161
		CHAINS[42161] = chain
	}

	CG_DEMO_KEY, _ = os.LookupEnv("CG_DEMO_KEY")
}

func init() {
	SetEnv(`.env`)
}
