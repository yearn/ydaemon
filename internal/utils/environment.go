package utils

import (
	"os"

	"github.com/joho/godotenv"

	"github.com/majorfi/ydaemon/internal/logs"
)

var (
	// RPCURIFor1 : RPC we should use for the chain #1
	RPCURIFor1 string
	// RPCURIFor250 : RPC we should use for the chain #250
	RPCURIFor250 string
	// RPCURIFor42161 : RPC we should use for the chain #42161
	RPCURIFor42161 string
)

// SetEnv will init the environment variables based on the .env file
func SetEnv(path string) {
	var exists bool
	_ = godotenv.Load(path)

	RPCURIFor1, exists = os.LookupEnv("RPC_URI_FOR_1")
	if !exists {
		RPCURIFor1 = "https://eth.public-rpc.com"
		logs.Warning("RPC_URI_FOR_1 not set, using default value: [https://eth.public-rpc.com]")
	}

	RPCURIFor250, exists = os.LookupEnv("RPC_URI_FOR_250")
	if !exists {
		RPCURIFor250 = "https://rpc.ftm.tools"
		logs.Warning("RPC_URI_FOR_250 not set, using default value: [https://rpc.ftm.tools]")
	}

	RPCURIFor42161, exists = os.LookupEnv("RPC_URI_FOR_42161")
	if !exists {
		RPCURIFor42161 = "https://arbitrum.public-rpc.com"
		logs.Warning("RPC_URI_FOR_42161 not set, using default value: [https://arbitrum.public-rpc.com]")
	}
}

func init() {
	SetEnv(`.env`)
}
