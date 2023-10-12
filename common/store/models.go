package store

/**************************************************************************************************
** DBBaseSchema is the base of all schemas stored in the SQL database. It will store the UUID,
** aka a unique indentifier, the block number and the chainID. Every schema will inherit from this
** base schema.
**************************************************************************************************/
type DBBaseSchema struct {
	UUID    string `gorm:"primaryKey" json:"UUID,omitempty"`
	Block   uint64 `json:"block"`
	ChainID uint64 `json:"chainID"`
}

/**************************************************************************************************
** DBBlockTime corresponds to the db_block_times SQL table. It will store a simple schema of
** block number and timestamp. This is used to retrieve the block time of the current block.
**************************************************************************************************/
type DBBlockTime struct {
	DBBaseSchema
	Timestamp uint64 `json:"timestamp"`
}

/**************************************************************************************************
** DBVaultPricePerShare is a struct used to keep track of the pricePerShare for a vault, every day
** at noon UTC.
**************************************************************************************************/
type DBVaultPricePerShare struct {
	UUID                   string  `gorm:"primaryKey" json:"UUID,omitempty"` // UUID is the unique identifier of the entry
	Block                  uint64  `json:"block"`                            // Block is the block number of the entry
	Time                   uint64  `json:"time"`                             // Time is the timestamp of the entry
	ChainID                uint64  `json:"chainID"`                          // ChainID is the chain on which the vault is deployed
	Vault                  string  `json:"vault"`                            // Vault is the vault address
	PricePerShare          string  `json:"pricePerShare"`                    // PricePerShare is the price per share of the vault
	HumanizedPricePerShare float64 `json:"humanizedPricePerShare"`           // HumanizedPricePerShare is the price per share of the vault, humanized in base vaultDecimals
}
