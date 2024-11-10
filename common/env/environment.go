package env

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/yearn/ydaemon/common/logs"
)

// SetEnv will init the environment variables based on the .env file
func SetEnv() {
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
	godotenv.Load(`.env`)
	ETHEREUM.SubgraphURI = os.Getenv("SUBGRAPGH_FOR_1")
	OPTIMISM.SubgraphURI = os.Getenv("SUBGRAPGH_FOR_10")
	ARBITRUM.SubgraphURI = os.Getenv("SUBGRAPGH_FOR_42161")
	CHAINS[1] = ETHEREUM
	CHAINS[10] = OPTIMISM
	CHAINS[100] = GNOSIS
	CHAINS[137] = POLYGON
	CHAINS[250] = FANTOM
	CHAINS[8453] = BASE
	CHAINS[42161] = ARBITRUM
	SetEnv()

	// Set them as supported
	for k := range CHAINS {
		SUPPORTED_CHAIN_IDS = append(SUPPORTED_CHAIN_IDS, k)
	}
	logs.Pretty(ETHEREUM.RpcURI)
}
