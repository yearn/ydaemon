package prices

// LLAMA_CHAIN_NAMES contains the chain identifiers for the DeFiLlama API
var LLAMA_CHAIN_NAMES = map[uint64]string{
	1:     `ethereum`,
	10:    `optimism`,
	250:   `fantom`,
	42161: `arbitrum`,
}

// LLAMA_PRICE_URL contains the URL for the DeFiLlama pricing API
var LLAMA_PRICE_URL = `https://coins.llama.fi/prices/current/`

// GECKO_CHAIN_NAMES contains the chain identifiers for the CoinGecko API
var GECKO_CHAIN_NAMES = map[uint64]string{
	1:     `ethereum`,
	10:    `optimistic-ethereum`,
	250:   `fantom`,
	42161: `arbitrum-one`,
}

// GECKO_PRICE_URL contains the URL for the CoinGecko API
var GECKO_PRICE_URL = `https://api.coingecko.com/api/v3/simple/token_price/`
