package store

import (
	"github.com/yearn/ydaemon/internal/models"
)

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
** DBHistoricalPrice corresponds to the db_historical_price SQL table. It will store a simple
** schema of block number, tokenAddress and price. This is used to retrieve the price of a token
** at a specific block.
**************************************************************************************************/
type DBHistoricalPrice struct {
	DBBaseSchema
	Token          string  // Token is the address of the token
	Price          string  // Price is the price of the token as raw string
	HumanizedPrice float64 // HumanizedPrice is the price of the token in human readable format
}

/**************************************************************************************************
** DBNewVaultsFromRegistry corresponds to the db_new_vaults_from_registry SQL table.
** It will store the FilterNewVault and FilterNewExperimentalVaults from the registry contract to
** avoir requerrying all the logs to find them.
**************************************************************************************************/
type DBNewVaultsFromRegistry struct {
	DBBaseSchema
	RegistryAddress string            // RegistryAddress is the address of the registry contract
	VaultsAddress   string            // VaultAddress is the address of the vault contract
	TokenAddress    string            // TokenAddress is the address of the underlying token
	APIVersion      string            // APIVersion is the version of the vault API
	Type            models.TTokenType // Type of vault, can be Experimental, Standard or Automated
}

/**************************************************************************************************
** DBVault corresponds to the db_vaults SQL table. It will store the vaults information.
** Only STATIC informations (or suppposed to be static) will be stored here. Dynamic informations
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
	Type           models.TTokenType // Type of vault, can be Experimental, Standard or Automated
	Symbol         string            // Symbol of the vault
	DisplaySymbol  string            // DisplaySymbol is an overwrite symbol for display purpose
	FormatedSymbol string            // FormatedSymbol is an overwrite symbol yv + symbol
	Name           string            // Name of the vault
	DisplayName    string            // DisplayName is an overwrite name for display purpose
	FormatedName   string            // FormatedName is an overwrite name yearn vault + name
	Version        string            // Version of the vault
	ChainID        uint64            // ChainID is the chain on which the vault is deployed
	Inception      uint64            // Inception indicates when the vault was deployed
	Activation     uint64            // Activation indicates when the vault was activated
	Decimals       uint64            // Decimals of the vault token
	PerformanceFee uint64            // PerformanceFee of the vault (/!\ semi-static)
	ManagementFee  uint64            // ManagementFee of the vault (/!\ semi-static)
	Endorsed       bool              // Endorsed indicates if the vault is endorsed by yearn
}

/**************************************************************************************************
** DBERC20 corresponds to the db_erc20 SQL table. It will store an ERC20 minimal information.
**************************************************************************************************/
type DBERC20 struct {
	UUID                      string            `gorm:"primaryKey" json:"UUID,omitempty"`
	Address                   string            // Address is the token address
	Name                      string            // Name of the token
	Symbol                    string            // Symbol of the token
	DisplayName               string            // DisplayName is an overwrite name for display purpose
	DisplaySymbol             string            // DisplaySymbol is an overwrite symbol for display purpose
	Description               string            // Description of the token
	UnderlyingTokensAddresses string            // UnderlyingTokensAddresses is the list of underlying tokens addresses, as a single string separated by a comma
	Type                      models.TTokenType // Type of the token, Vault, Native, etc...
	Decimals                  uint64            // Decimals of the token
	ChainID                   uint64            // ChainID is the chain on which the token is deployed
}

/**************************************************************************************************
** DBRegistrySync is a struct used to keep track of the latest block synced for a specific chain
** for the registry new vaults events.
**************************************************************************************************/
type DBRegistrySync struct {
	UUID     string `gorm:"primaryKey" json:"UUID,omitempty"`
	Block    uint64 `json:"block"`
	ChainID  uint64 `json:"chainID"`
	Registry string `json:"registry"`
}

/**************************************************************************************************
** DBStrategy corresponds to the db_strategies SQL table. It will store the strategies information
** from the events added or migrated.
** Only STATIC informations (or suppposed to be static) will be stored here. Dynamic informations
** should be fetched directly when needed.
**************************************************************************************************/
type DBStrategy struct {
	UUID              string `gorm:"primaryKey" json:"UUID,omitempty"`
	VaultAddress      string // VaultAddress is the address of the vault contract for this strategy
	StrategyAddress   string // StrategyAddress is the address of the strategy contract
	TxHash            string // TxHash is the hash of the transaction where the strategy has been added
	DebtRatio         string // DebtRatio is the debt ratio of the strategy (>= 0.3.0)
	MaxDebtPerHarvest string // MaxDebtPerHarvest is the max debt per harvest of the strategy (>= 0.3.2)
	MinDebtPerHarvest string // MinDebtPerHarvest is the min debt per harvest of the strategy (>= 0.3.2)
	PerformanceFee    string // PerformanceFee is the performance fee of the strategy  (>= 0.2.2)
	DebtLimit         string // DebtLimit is the debt limit of the strategy (== 0.2.2)
	RateLimit         string // RateLimit is the rate limit of the strategy (== 0.2.2 - 0.3.0)
	VaultVersion      string // VaultVersion is the version of the vault at the time of the strategy addition
	ChainID           uint64 // ChainID is the chain on which the strategy is deployed
	BlockNumber       uint64 // BlockNumber is the block number where the strategy has been added
	TxIndex           uint   // TxIndex is the index of the transaction in the block
	LogIndex          uint   // LogIndex is the index of the log in the transaction
}

/**************************************************************************************************
** DBStrategyAddedSync is a struct used to keep track of the latest block synced for a specific
** chain when it comes to the strategy added events.
**************************************************************************************************/
type DBStrategyAddedSync struct {
	UUID    string `gorm:"primaryKey" json:"UUID,omitempty"`
	Block   uint64 `json:"block"`
	ChainID uint64 `json:"chainID"`
	Vault   string `json:"vault"`
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

/**************************************************************************************************
** DBVaultActivation is a struct used to keep track of the activation for a vault
**************************************************************************************************/
type DBVaultActivation struct {
	UUID    string `gorm:"primaryKey" json:"UUID,omitempty"` // UUID is the unique identifier of the entry
	Block   uint64 `json:"block"`                            // Block is the block number of the entry
	ChainID uint64 `json:"chainID"`                          // ChainID is the chain on which the vault is deployed
	Vault   string `json:"vault"`                            // Vault is the vault address
}
