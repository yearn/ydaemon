package store

// TTables lists the keys used in the database
type TTables = struct {
	BLOCK_TIME    string
	PRICES        string
	STRATEGIES    string
	TOKENS        string
	VAULTS        string
	VAULTS_LEGACY string
}

var TABLES = TTables{
	BLOCK_TIME:    "blocktime",
	PRICES:        "prices",
	STRATEGIES:    "strategies",
	TOKENS:        "tokens",
	VAULTS:        "vaults",
	VAULTS_LEGACY: "legacy_vaults",
}

// TKeys lists the keys used in the database
type TKeys = struct {
	TOKENS_SUBGRAPH_DATA string
	TOKENS_SUBGRAPH_LIST string
}

var KEYS = TKeys{
	TOKENS_SUBGRAPH_DATA: "tokens_subgraph_data",
	TOKENS_SUBGRAPH_LIST: "tokens_subgraph_list",
}
