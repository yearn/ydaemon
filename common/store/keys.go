package store

// TTables lists the keys used in the database
type TTables = struct {
	BLOCK_TIME        string
	PRICES            string
	HISTORICAL_PRICES string
	STRATEGIES        string
	TOKENS            string
	VAULTS            string
	VAULTS_LEGACY     string
}

var TABLES = TTables{
	BLOCK_TIME:        "blocktime",
	PRICES:            "prices",
	HISTORICAL_PRICES: "historical_prices",
	STRATEGIES:        "strategies",
	TOKENS:            "tokens",
	VAULTS:            "vaults",
	VAULTS_LEGACY:     "legacy_vaults",
}
