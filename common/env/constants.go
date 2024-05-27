package env

import (
	"path"
	"path/filepath"
	"runtime"

	"github.com/ethereum/go-ethereum/common"
)

type TContractData struct {
	Address common.Address // Address of the contract
	Block   uint64         // Block number where the contract was deployed
	Version uint64         // Version of the contract. May be empty.
	Tag     string         // Tag of the contract. May be empty.
	Label   string         // Label of the contract. May be empty.
}

func getCurrentPath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

// BASE_DATA_PATH is the base path to access the data informations
var BASE_DATA_PATH, _ = filepath.Abs(getCurrentPath() + `../../../data/`)

// var BASE_ASSET_URL = `https://raw.githubusercontent.com/yearn/yearn-assets/master/icons/multichain-tokens/`
// BASE_ASSET_URL is the base URL to access the tokens icons from the assets repo
var BASE_ASSET_URL = `https://assets.smold.app/api/token/`

// GECKO_PRICE_URL contains the URL for the CoinGecko API
var GECKO_PRICE_URL = `https://api.coingecko.com/api/v3/simple/token_price/`

// LLAMA_PRICE_URL contains the URL for the DeFiLlama pricing API
var LLAMA_PRICE_URL = `https://coins.llama.fi/prices/current/`

var CG_DEMO_KEYS = []string{}
