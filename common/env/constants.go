package env

import (
	"path"
	"path/filepath"
	"runtime"

	"github.com/ethereum/go-ethereum/common"
)

/**************************************************************************************************
** TContractData represents a comprehensive set of metadata for a blockchain contract.
** This structure provides not only the address but also contextual information like
** deployment block, version, and descriptive labels.
**
** @field Address The Ethereum address of the contract
** @field Block The block number at which the contract was deployed
** @field Version Optional numeric version identifier for the contract
** @field Tag Optional short identifier for categorizing the contract
** @field Label Optional longer description for the contract's purpose
**************************************************************************************************/
type TContractData struct {
	Address common.Address // Address of the contract
	Block   uint64         // Block number where the contract was deployed
	Version uint64         // Version of the contract. May be empty.
	Tag     string         // Tag of the contract. May be empty.
	Label   string         // Label of the contract. May be empty.
}

/**************************************************************************************************
** getCurrentPath is a helper function that returns the absolute path of the current package.
** It uses runtime reflection to determine where the code is being executed from.
**
** @return string The absolute path to the current package directory
**************************************************************************************************/
func getCurrentPath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

/**************************************************************************************************
** BASE_DATA_PATH is the absolute path to the data directory used by yDaemon.
** This path is calculated relative to the location of this file in the codebase.
** It serves as the root directory for various data files used by the application.
**************************************************************************************************/
var BASE_DATA_PATH, _ = filepath.Abs(getCurrentPath() + `../../../data/`)

// var BASE_ASSET_URL = `https://raw.githubusercontent.com/yearn/yearn-assets/master/icons/multichain-tokens/`
/**************************************************************************************************
** BASE_ASSET_URL is the base URL used to fetch token icons and other assets.
** This endpoint provides access to token logo images and other visual resources
** used to enhance the UI experience for applications using yDaemon.
**************************************************************************************************/
var BASE_ASSET_URL = `https://cdn.jsdelivr.net/gh/yearn/tokenassets@main/tokens`

/**************************************************************************************************
** GECKO_PRICE_URL contains the base URL for the CoinGecko token pricing API.
** This endpoint is used to retrieve current and historical token prices from CoinGecko.
**************************************************************************************************/
var GECKO_PRICE_URL = `https://api.coingecko.com/api/v3/simple/token_price/`

/**************************************************************************************************
** LLAMA_PRICE_URL contains the base URL for the DeFiLlama pricing API.
** This endpoint serves as an alternative or backup to CoinGecko for token price data.
** DeFiLlama often provides price data for tokens that may not be available on CoinGecko.
**************************************************************************************************/
var LLAMA_PRICE_URL = `https://coins.llama.fi/prices/current/`

/**************************************************************************************************
** CG_DEMO_KEYS stores an array of CoinGecko API keys that can be used for API requests.
** Having multiple keys allows for distribution of requests to avoid rate limiting.
** This array is populated from the CG_DEMO_KEYS environment variable during initialization.
**************************************************************************************************/
var CG_DEMO_KEYS = []string{}

/**************************************************************************************************
** CMS_ROOT_URL contains the base URL for the CMS (Content Management System) API.
** This endpoint is used to fetch vault metadata and other CMS-related data.
** When CMS_ROOT_URL ``, yDaemon will fetch metadata locally from
** ../webops-sdk/packages/cms/cdn/content
**************************************************************************************************/
var CMS_ROOT_URL = ``

/**************************************************************************************************
** RISK_CDN_URL contains the base URL for the risk score CDN.
** This endpoint is used to fetch vault risk scores from the Yearn risk assessment system.
** The risk scores are stored as JSON files organized by chain ID and vault address.
** Default value points to the production CDN, but can be overridden via RISK_CDN_URL env variable.
**************************************************************************************************/
var RISK_CDN_URL = `https://risk.yearn.fi/cdn/`

/**************************************************************************************************
** KONG_API_URL contains the GraphQL endpoint for Kong, the single source of truth for vault
** and strategy discovery. Defaults to https://kong.yearn.farm/api/gql
**************************************************************************************************/
var KONG_API_URL = `https://kong.yearn.farm/api/gql`
