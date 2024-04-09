package env

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/yearn/ydaemon/common/logs"
)

// SetEnv will init the environment variables based on the .env file
func SetEnv(path string) {
	godotenv.Load(path)

	for _, chain := range CHAINS {
		baseKey := `RPC_URI_FOR_`
		chainID := strconv.FormatUint(chain.ID, 10)
		RPCURI, exists := os.LookupEnv(baseKey + chainID)
		if !exists {
			logs.Debug(baseKey + chainID + " not set, using default value")
		} else {
			chain := CHAINS[chain.ID]
			chain.RpcURI = RPCURI
			CHAINS[chain.ID] = chain
		}
	}

	/**********************************************************************************************
	** Array of Coingecko keys to use
	**********************************************************************************************/
	allCGKeys, _ := os.LookupEnv("CG_DEMO_KEYS")
	if allCGKeys != `` {
		splittedKeys := strings.Split(allCGKeys, ",")
		CG_DEMO_KEYS = append(CG_DEMO_KEYS, splittedKeys...)
	}

}

func init() {
	SetEnv(`.env`)
}
