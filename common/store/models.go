package store

import (
	"github.com/yearn/ydaemon/internal/models"
)

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

/**************************************************************************************************
** DBVault corresponds to the db_vaults mySQL table. It will store the vaults information.
** Only STATIC informations (or supposed to be static) will be stored here. Dynamic informations
** should be fetched directly when needed.
**************************************************************************************************/
type DBVault struct {
	UUID           string            `gorm:"primaryKey" json:"UUID,omitempty"`
	Address        string            // Address is the vault address
	Management     string            // Management is the management address
	Governance     string            // Governance is the governance address
	Guardian       string            // Guardian is the guardian address
	Rewards        string            // Rewards is the rewards address
	Token          string            // Token is the underlying token address
	Type           models.TVaultType // Type of vault, can be Experimental, Standard or Automated
	Symbol         string            // Symbol of the vault
	DisplaySymbol  string            // DisplaySymbol is an overwrite symbol for display purpose
	FormatedSymbol string            // FormatedSymbol is an overwrite symbol yv + symbol
	Name           string            // Name of the vault
	DisplayName    string            // DisplayName is an overwrite name for display purpose
	FormatedName   string            // FormatedName is an overwrite name yearn vault + name
	Icon           string            // Icon URL to access the vault icon
	Version        string            // Version of the vault
	ChainID        uint64            // ChainID is the chain on which the vault is deployed
	Inception      uint64            // Inception indicates when the vault was deployed
	Activation     uint64            // Activation indicates when the vault was activated
	Decimals       uint64            // Decimals of the vault token
	PerformanceFee uint64            // PerformanceFee of the vault (/!\ semi-static)
	ManagementFee  uint64            // ManagementFee of the vault (/!\ semi-static)
	Endorsed       bool              // Endorsed indicates if the vault is endorsed by yearn
}
