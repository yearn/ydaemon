package models

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
)

type TVaultsFromRegistry struct {
	Address         common.Address `json:"address"`
	RegistryAddress common.Address `json:"registryAddress"`
	TokenAddress    common.Address `json:"tokenAddress"`
	Type            TTokenType     `json:"type"`
	APIVersion      string         `json:"version"`
	ChainID         uint64         `json:"chainID"`
	BlockNumber     uint64         `json:"blockNumber"`
}

// TTVL holds the info about the value locked in a vault
type TTVL struct {
	TotalAssets          *bigNumber.Int `json:"totalAssets"`
	TotalDelegatedAssets *bigNumber.Int `json:"total_delegated_assets"`
	TVLDeposited         float64        `json:"tvl_deposited"`
	TVLDelegated         float64        `json:"tvl_delegated"`
	TVL                  float64        `json:"tvl"`
	Price                float64        `json:"price"`
}

// TAPYFees holds the fees information about this vault.
type TAPYFees struct {
	Performance float64 `json:"performance"`
	Withdrawal  float64 `json:"withdrawal"`
	Management  float64 `json:"management"`
	KeepCRV     float64 `json:"keep_crv"`
	KeepVelo    float64 `json:"keep_velo"`
	CvxKeepCRV  float64 `json:"cvx_keep_crv"`
}

// TAPYPoints holds the points information about this vault.
type TAPYPoints struct {
	WeekAgo   float64 `json:"week_ago"`
	MonthAgo  float64 `json:"month_ago"`
	Inception float64 `json:"inception"`
}

// TAPYComposite holds the points information about this vault.
type TAPYComposite struct {
	Boost      float64 `json:"boost"`
	PoolAPY    float64 `json:"pool_apy"`
	BoostedAPR float64 `json:"boosted_apr"`
	BaseAPR    float64 `json:"base_apr"`
	CvxAPR     float64 `json:"cvx_apr"`
	RewardsAPR float64 `json:"rewards_apr"`
}

// TAPY contains all the information useful about the APY, APR, fees and breakdown.
type TAPY struct {
	Type              string        `json:"type"`
	GrossAPR          float64       `json:"gross_apr"`
	NetAPY            float64       `json:"net_apy"`
	StakingRewardsAPR float64       `json:"staking_rewards_apr"`
	Fees              TAPYFees      `json:"fees"`
	Points            TAPYPoints    `json:"points"`
	Composite         TAPYComposite `json:"composite"`
	Error             string        `json:"error_reason"`
}

// TMigration helps us to know if a vault is in the process of being migrated.
type TMigration struct {
	Available bool           `json:"available"`
	Target    common.Address `json:"target"`
	Contract  common.Address `json:"contract,omitempty"`
}

// TStaking holds some metadata about the staking contract.
type TStaking struct {
	Address      common.Address   `json:"address"`
	VaultAddress common.Address   `json:"vaultAddress"`
	Available    bool             `json:"available"`
	Risk         int              `json:"risk"`
	TVL          float64          `json:"tvl"`
	Amount       *bigNumber.Float `json:"amount"`
	Price        *bigNumber.Float `json:"price"`
}

// TVault is the main structure returned by the API when trying to get all the vaults for a specific network
type TVault struct {
	// Immutable elements. They won't change
	Address      common.Address `json:"address"`    // Address of the vault
	AssetAddress common.Address `json:"token"`      // Address of the underlying token
	Management   common.Address `json:"management"` // Address of the management contract
	Governance   common.Address `json:"governance"` // Address of the governance contract
	Guardian     common.Address `json:"guardian"`   // Address of the guardian contract
	Rewards      common.Address `json:"rewards"`    // Address of the rewards contract
	Type         TTokenType     `json:"type"`
	Version      string         `json:"version"`    // The version of the vault
	Activation   uint64         `json:"activation"` // When the vault was activated
	ChainID      uint64         `json:"chainID"`    // The chainID of the vault
	Endorsed     bool           `json:"endorsed"`   // If the vault is endorsed by Yearn

	// Semi-mutable eelements. They can change but rarely
	PerformanceFee    uint64 `json:"performanceFee"`    // The performance fee of the vault
	ManagementFee     uint64 `json:"managementFee"`     // The management fee of the vault
	EmergencyShutdown bool   `json:"emergencyShutdown"` // If the vault is in emergency shutdown

	// Mutable elements. They will often change
	LastActiveStrategies []common.Address `json:"lastActiveStrategies"` // The list of "active" strategies via their withdrawal queue
	LastPricePerShare    *bigNumber.Int   `json:"lastPricePerShare"`    // Price per share of the vault
	LastTotalAssets      *bigNumber.Int   `json:"lastTotalAssets"`      // Total assets locked in the vault
	LastUpdate           time.Time        `json:"lastUpdate"`           // When the vault was last updated

	// Manual elements. They are manually set by the team
	IsRetired      bool            `json:"isRetired"`      // If the vault is retired or not
	IsHidden       bool            `json:"isHidden"`       // If the vault is hidden or not
	Migration      TMigration      `json:"migration"`      // If the vault is in the process of being migrated
	Classification TClassification `json:"classification"` // The classification of the vault
}

type TLegacyAPIAPY struct {
	Type              string  `json:"type"`
	GrossAPR          float64 `json:"gross_apr"`
	NetAPY            float64 `json:"net_apy"`
	StakingRewardsAPR float64 `json:"staking_rewards_apr"`
	Fees              struct {
		Performance float64 `json:"performance"`
		Withdrawal  float64 `json:"withdrawal"`
		Management  float64 `json:"management"`
		KeepCRV     float64 `json:"keep_crv"`
		KeepVelo    float64 `json:"keep_velo"`
		CvxKeepCRV  float64 `json:"cvx_keep_crv"`
	} `json:"fees"`
	Points struct {
		WeekAgo   float64 `json:"week_ago"`
		MonthAgo  float64 `json:"month_ago"`
		Inception float64 `json:"inception"`
	} `json:"points"`
	Composite struct {
		Boost      float64 `json:"boost"`
		PoolAPY    float64 `json:"pool_apy"`
		BoostedAPR float64 `json:"boosted_apr"`
		BaseAPR    float64 `json:"base_apr"`
		CvxAPR     float64 `json:"cvx_apr"`
		RewardsAPR float64 `json:"rewards_apr"`
	} `json:"composite"`
	Error string `json:"error_reason"`
}
type TLegacyAPI struct {
	Address common.MixedcaseAddress
	APY     TLegacyAPIAPY
}

type TAggregatedVault struct {
	Address       common.MixedcaseAddress
	LegacyAPY     TLegacyAPIAPY
	PricePerShare bigNumber.Int
}

type TClassification struct {
	IsAutomated     bool   `json:"isAutomated"`
	IsPool          bool   `json:"isPool"`
	PoolProvider    string `json:"poolProvider,omitempty"`
	Stability       string `json:"stability"`
	StableBaseAsset string `json:"stableBaseAsset,omitempty"`
}
