package env

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/yearn/ydaemon/common/logs"
)

/**************************************************************************************************
** SetEnv initializes and updates environment-specific configuration values by reading from
** environment variables. This function is called during application startup by the init() function.
**
** It performs two primary tasks:
** 1. Updates RPC URIs for all supported chains from environment variables (RPC_URI_FOR_[chainID])
** 2. Loads CoinGecko API keys from the CG_DEMO_KEYS environment variable
**
** The function allows for runtime configuration of network endpoints and API services without
** requiring code changes or recompilation.
**************************************************************************************************/
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

	/**********************************************************************************************
	** CMS root URL configuration
	**********************************************************************************************/
	if cmsRoot, exists := os.LookupEnv("CMS_ROOT_URL"); exists {
		CMS_ROOT_URL = cmsRoot
	}

	/**********************************************************************************************
	** Risk score CDN URL configuration
	**********************************************************************************************/
	if riskCDN, exists := os.LookupEnv("RISK_CDN_URL"); exists {
		RISK_CDN_URL = riskCDN
	}

	/**********************************************************************************************
	** Kong API URL configuration
	**********************************************************************************************/
	if kongURL, exists := os.LookupEnv("KONG_API_URL"); exists {
		KONG_API_URL = kongURL
	}
}

/**************************************************************************************************
** init is automatically called when the package is imported and performs the initial setup of
** the environment configuration. This function is responsible for:
**
** 1. Loading variables from the .env file into the environment
** 2. Setting subgraph URIs for various chains from environment variables
** 3. Populating the CHAINS map with all supported blockchain networks
** 4. Calling SetEnv to apply any environment-specific overrides
** 5. Building the list of supported chain IDs for easy reference
**
** This initialization ensures that all chain configurations are properly loaded before
** any other parts of the application attempt to use them.
**************************************************************************************************/
func init() {
	godotenv.Load(`.env`)
	ETHEREUM.SubgraphURI = os.Getenv("SUBGRAPGH_FOR_1")
	OPTIMISM.SubgraphURI = os.Getenv("SUBGRAPGH_FOR_10")
	ARBITRUM.SubgraphURI = os.Getenv("SUBGRAPGH_FOR_42161")
	CHAINS[1] = ETHEREUM
	CHAINS[10] = OPTIMISM
	CHAINS[100] = GNOSIS
	CHAINS[137] = POLYGON
	CHAINS[146] = SONIC
	CHAINS[250] = FANTOM
	CHAINS[8453] = BASE
	CHAINS[42161] = ARBITRUM
	CHAINS[747474] = KATANA
	SetEnv()

	// Set them as supported
	for k := range CHAINS {
		SUPPORTED_CHAIN_IDS = append(SUPPORTED_CHAIN_IDS, k)
	}
}
