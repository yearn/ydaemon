package store

// TTables lists the keys used in the database
type TTables = struct {
	PRICES                     string
	VAULTS_LEGACY              string
	VAULTS_FROM_REGISTRY_SYNC  string // For a given registry, sync the last block checked for a given vault
	STRATEGIES_FROM_VAULT_SYNC string // For a given vault, current last block checked for the strategies added
	REGISTRY_SYNC              string
}

var TABLES = TTables{
	PRICES:                     `prices`,
	VAULTS_LEGACY:              `legacy_vaults`,
	VAULTS_FROM_REGISTRY_SYNC:  `db_new_vaults_from_registries`,
	STRATEGIES_FROM_VAULT_SYNC: `db_new_strategies_from_vaults`,
	REGISTRY_SYNC:              `db_registry_sync`,
}
