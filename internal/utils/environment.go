package utils

import (
	"os"

	"github.com/joho/godotenv"

	"github.com/majorfi/ydaemon/internal/logs"
)

var (
	// RPCURIFor1 : RPC we should use for the chain #1 (Ethereum)
	RPCURIFor1 string
	// RPCURIFor10 : RPC we should use for the chain #10 (Optimism)
	RPCURIFor10 string
	// RPCURIFor250 : RPC we should use for the chain #250 (Fantom)
	RPCURIFor250 string
	// RPCURIFor42161 : RPC we should use for the chain #42161 (Arbitrum)
	RPCURIFor42161 string
	// WebhookSecret: is a shared secret between Github Webook system and the Daemon to trigger some data refresh
	WebhookSecret string
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

	RPCURIFor10, exists = os.LookupEnv("RPC_URI_FOR_10")
	if !exists {
		RPCURIFor10 = "https://mainnet.optimism.io"
		logs.Warning("RPC_URI_FOR_10 not set, using default value: [https://mainnet.optimism.io]")
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

	WebhookSecret, exists = os.LookupEnv("WEBHOOK_SECRET")
	if !exists {
		logs.Warning("WEBHOOK_SECRET not set")
	}
}

func init() {
	SetEnv(`.env`)
}
