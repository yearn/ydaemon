package models

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
)

type TVaultKind string

const (
	VaultKindLegacy   TVaultKind = "Legacy"
	VaultKindMultiple TVaultKind = "Multi Strategy"
	VaultKindSingle   TVaultKind = "Single Strategy"
)

type TVaultStabilityType string

const (
	VaultStabilityStable   TVaultStabilityType = "Stable"
	VaultStabilityVolatile TVaultStabilityType = "Volatile"
	VaultStabilityUnknown  TVaultStabilityType = "Unknown"
)

type TVaultCategoryType string

const (
	VaultCategoryAutomatic TVaultCategoryType = "auto"
)

type TVaultsFromRegistry struct {
	Address         common.Address `json:"address"`
	RegistryAddress common.Address `json:"registryAddress"`
	TokenAddress    common.Address `json:"tokenAddress"`
	Type            TTokenType     `json:"type"`
	Kind            TVaultKind     `json:"kind"`
	APIVersion      string         `json:"version"`
	ChainID         uint64         `json:"chainID"`
	BlockNumber     uint64         `json:"blockNumber"`
}

// TTVL holds the info about the value locked in a vault
type TTVL struct {
	TotalAssets *bigNumber.Int `json:"totalAssets"`
	TVL         float64        `json:"tvl"`
	Price       float64        `json:"price"`
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

// TInclusion holds the inclusion of the vault
type TInclusion struct {
	IsSet           bool `json:"isSet"`           // If not set, automatic inclusion will be recomputed
	IsYearn         bool `json:"isYearn"`         // If the vault is a Yearn vault or not
	IsYearnJuiced   bool `json:"isYearnJuiced"`   // If the vault is a Yearn Juiced vault or not
	IsGimme         bool `json:"isGimme"`         // If the vault is a Gimme vault or not
	IsPoolTogether  bool `json:"isPoolTogether"`  // If the vault is a Gimme vault or not
	IsPublicERC4626 bool `json:"isPublicERC4626"` // If the vault is from the public ERC4626 registry or not
}

type TVaultMetadata struct {
	IsRetired      bool               `json:"isRetired"`      // If the vault is retired or not
	IsHidden       bool               `json:"isHidden"`       // If the vault is hidden or not
	IsAggregator   bool               `json:"isAggregator"`   // If the vault is an aggregator or not
	IsBoosted      bool               `json:"isBoosted"`      // If the vault is boosted or not. Act as a fallback if yDaemon has no way to know it
	IsAutomated    bool               `json:"isAutomated"`    // If the vault is automated or not
	IsHighlighted  bool               `json:"isHighlighted"`  // If the vault is highlighted or not
	IsPool         bool               `json:"isPool"`         // If the vault is a pool or not
	ShouldUseV2APR bool               `json:"shouldUseV2APR"` // If the vault should use the V2 APR or not (only for V3 vaults)
	Migration      TMigration         `json:"migration"`      // If the vault is in the process of being migrated
	Stability      TStability         `json:"stability"`      // The stability of the vault
	Category       TVaultCategoryType `json:"category"`       // The category of the vault
	DisplayName    string             `json:"displayName"`    // The name of the vault
	DisplaySymbol  string             `json:"displaySymbol"`  // The symbol of the vault
	Description    string             `json:"description"`    // The description of the vault
	SourceURI      string             `json:"sourceURI"`      // The source URI of the vault
	UINotice       string             `json:"uiNotice"`       // The notice to display in the UI
	RiskLevel      int8               `json:"riskLevel"`      // The risk level of the vault (1 to 5, -1 if not set)
	Protocols      []string           `json:"protocols"`      // The Protocols used by the vault
	Inclusion      TInclusion         `json:"inclusion"`      // Inclusion is a special field to know "where" the vault should be displayed.
}

// TVault is the main structure returned by the API when trying to get all the vaults for a specific network
type TVault struct {
	// Immutable elements. They won't change
	Address         common.Address  `json:"address"`              // Address of the vault
	AssetAddress    common.Address  `json:"token"`                // Address of the underlying token
	RegistryAddress common.Address  `json:"registry"`             // Address of the registry
	Accountant      *common.Address `json:"accountant,omitempty"` // The address of the accountant
	Type            TTokenType      `json:"type"`                 // The type of the vault
	Kind            TVaultKind      `json:"kind"`                 // The kind of the vault (legacy, multi, single)
	Version         string          `json:"version"`              // The version of the vault
	Activation      uint64          `json:"activation"`           // When the vault was activated
	ChainID         uint64          `json:"chainID"`              // The chainID of the vault
	Endorsed        bool            `json:"endorsed"`             // If the vault is endorsed by Yearn

	// Semi-mutable eelements. They can change but rarely
	PerformanceFee    uint64 `json:"performanceFee"`    // The performance fee of the vault
	ManagementFee     uint64 `json:"managementFee"`     // The management fee of the vault
	EmergencyShutdown bool   `json:"emergencyShutdown"` // If the vault is in emergency shutdown

	// Mutable elements. They will often change
	LastActiveStrategies []common.Address `json:"lastActiveStrategies"` // The list of "active" strategies via their withdrawal queue
	LastPricePerShare    *bigNumber.Int   `json:"lastPricePerShare"`    // Price per share of the vault
	LastTotalAssets      *bigNumber.Int   `json:"lastTotalAssets"`      // Total assets locked in the vault

	// Manual elements. They are manually set by the team
	Metadata TVaultMetadata `json:"metadata"` // The metadata of the vault
}

type TLegacyAPIAPY struct {
	Type              string  `json:"type"`
	GrossAPR          float64 `json:"gross_apr"`
	NetAPY            float64 `json:"net_apy"`
	StakingRewardsAPR float64 `json:"staking_rewards_apr"`
	Fees              struct {
		Performance float64 `json:"performance"`
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

type TStability struct {
	Stability       TVaultStabilityType `json:"stability"`
	StableBaseAsset string              `json:"stableBaseAsset,omitempty"`
}
