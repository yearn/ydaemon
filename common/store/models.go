package store

import "github.com/yearn/ydaemon/internal/models"

/**************************************************************************************************
** DBBaseSchema is the base of all schemas stored in the mySQL database. It will store the UUID,
** aka a unique indentifier, the block number and the chainID. Every schema will inherit from this
** base schema.
**************************************************************************************************/
type DBBaseSchema struct {
	UUID    string `gorm:"primaryKey" json:"UUID,omitempty"`
	Block   uint64 `json:"block"`
	ChainID uint64 `json:"chainID"`
}

/**************************************************************************************************
** DBBlockTime corresponds to the db_block_times mySQL table. It will store a simple schema of
** block number and timestamp. This is used to retrieve the block time of the current block.
**************************************************************************************************/
type DBBlockTime struct {
	DBBaseSchema
	Timestamp uint64 `json:"timestamp"`
}

/**************************************************************************************************
** DBHistoricalPrice corresponds to the db_historical_price mySQL table. It will store a simple
** schema of block number, tokenAddress and price. This is used to retrieve the price of a token
** at a specific block.
**************************************************************************************************/
type DBHistoricalPrice struct {
	DBBaseSchema
	Token          string  `json:"token"`
	Price          string  `json:"price"`
	HumanizedPrice float64 `json:"humanizedPrice"`
}

/**************************************************************************************************
** DBNewVaultsFromRegistry corresponds to the db_new_vaults_from_registry mySQL table.
** It will store the FilterNewVault and FilterNewExperimentalVaults from the registry contract to
** avoir requerrying all the logs to find them.
**************************************************************************************************/
type DBNewVaultsFromRegistry struct {
	DBBaseSchema
	RegistryAddress string
	VaultsAddress   string
	TokenAddress    string
	BlockHash       string
	Type            models.TVaultType
	APIVersion      string
	Activation      uint64
	ManagementFee   uint64
	TxIndex         uint
	LogIndex        uint
}
