package models

import (
	"encoding/json"
	"strconv"

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
	VaultStabilityStable     TVaultStabilityType = "Stable"
	VaultStabilityVolatile   TVaultStabilityType = "Volatile"
	VaultStabilityUnknown    TVaultStabilityType = "Unknown"
	VaultStabilityCorrelated TVaultStabilityType = "Correlated"
	VaultStabilityUnstable   TVaultStabilityType = "Unstable"
)

type TVaultCategoryType string

const (
	VaultCategoryAutomatic TVaultCategoryType = "auto"
)

type TExtraProperties struct {
	YieldVaultAddress       string `json:"yieldVaultAddress,omitempty"`
	YearnVaultAsset         string `json:"yearnVaultAsset,omitempty"`
	YearnVault              string `json:"yearnVault,omitempty"`
	IsVaultV2               bool   `json:"isVaultV2,omitempty"`
	YearnGauge              string `json:"yearnGauge,omitempty"`
	CoveYearnStrategy       string `json:"coveYearnStrategy,omitempty"`
	AutoCompoundingGauge    string `json:"autoCompoundingGauge,omitempty"`
	NonAutoCompoundingGauge string `json:"nonAutoCompoundingGauge,omitempty"`
}

type TVaultsFromRegistry struct {
	Address         common.Address   `json:"address"`
	RegistryAddress common.Address   `json:"registryAddress"`
	TokenAddress    common.Address   `json:"tokenAddress"`
	Type            TTokenType       `json:"type"`
	Kind            TVaultKind       `json:"kind"`
	ExtraProperties TExtraProperties `json:"extraProperties,omitempty"`
	APIVersion      string           `json:"version"`
	ChainID         uint64           `json:"chainID"`
	BlockNumber     uint64           `json:"blockNumber"`
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

// TInclusion holds the inclusion of the vault
type TInclusion struct {
	IsSet           bool `json:"isSet"`           // If not set, automatic inclusion will be recomputed
	IsYearn         bool `json:"isYearn"`         // If the vault is a Yearn vault or not
	IsYearnJuiced   bool `json:"isYearnJuiced"`   // If the vault is a Yearn Juiced vault or not
	IsGimme         bool `json:"isGimme"`         // If the vault is a Gimme vault or not
	IsPoolTogether  bool `json:"isPoolTogether"`  // If the vault is a PoolTogether vault or not
	IsCove          bool `json:"isCove"`          // If the vault is a Cove related vault or not
	IsMorpho        bool `json:"isMorpho"`        // If the vault is a Morpho vault or not
	IsKatana        bool `json:"isKatana"`        // If the vault is a Katana vault or not
	IsPublicERC4626 bool `json:"isPublicERC4626"` // If the vault is from the public ERC4626 registry or not
}

// TRiskScore holds the risk score of the vault. All risk scores of the Single Strategy Vault. Multi-Strategy Vault won't have this object because its risk score is combination of multiple vaults. For risk value use `riskLevel`. (empty for Multi-Strategy Vault)
type TRiskScore struct {
	// Review score from Yearn security. 1 is for high trust to 5 low trust
	Review int8 `json:"review"`
	// Testing coverage of the strategy being evaluated. 5 -> 80% or less; 1 -> 100% or higher
	Testing int8 `json:"testing"`
	// Complexity: the sLOC count of the strategy being evaluated. 5 -> 750+ sLOC; 4 -> 450-600 sLOC; 3 -> 300-450 sLOC; 2 -> 150-300 sLOC; 1 -> 0-150 sLOC
	Complexity int8 `json:"complexity"`
	// RiskExposure score aims to find out how much and how often a strategy can be subject to losses. 5 -> Loss of funds or non recoverable funds up to 70-100% (Example, Leveraging cross assets and got liquidated, adding liquidity to volatile pairs single sided); 4 -> Loss of funds or non recoverable funds up to 15-70% (Example, adding liquidity to single sided curve stable pools); 3 -> Loss of funds or non recoverable funds up to 10-15% (Example, Protocol specific IL exposure, very high deposit/withdrawal fees); 2 -> Loss of funds or non recoverable funds up to 0-10% (Example, deposit/withdrawal fees or anything protocol specific); 1 -> Strategy has no lossable cases, only gains, up only.
	RiskExposure int8 `json:"riskExposure"`
	// ProtocolIntegration is the protocols that are integrated into the strategy that is being evaluated. 5 -> Strategy interacts with 5 external protocols; 4 -> Strategy interacts with 4 external protocols; 3 -> Strategy interacts with 3 external protocols; 2 -> Strategy interacts with 2 external protocols; 1 -> Strategy interacts with 1 external protocol
	ProtocolIntegration int8 `json:"protocolIntegration"`
	// CentralizationRisk is the centralization score of the strategy that is being evaluated. 5 -> Strategy heavily relies on off-chain management, potentially exposing user funds to rug possibilities by admins; 4 -> Strategy frequently depends on off-chain management but has safeguards against rug possibilities by admins; 3 -> Strategy involves privileged roles but less frequently and with less risk of rug possibilities; 2 -> Strategy has privileged roles but they are not vital for operations and pose minimal risk of rug possibilities; 1 -> Strategy operates without dependency on any privileged roles, ensuring full permissionlessness.`
	CentralizationRisk int8 `json:"centralizationRisk"`
	// ExternalProtocolAudit is the public audits count of the external protocols. 5 -> No audit conducted by a trusted firm or security researcher; 4 -> Audit conducted by 1 trusted firm or security researcher conducted; 3 -> Audit conducted by 2 trusted firm or security researcher conducted; 2 -> Audit conducted by 3 trusted firm or security researcher conducted; 1 -> Audit conducted by 4 or more trusted firm or security researcher conducted
	ExternalProtocolAudit int8 `json:"externalProtocolAudit"`
	// ExternalProtocolCentralisation is a measurement of the centralization score of the external protocols. 5 -> Contracts owner is an EOA or a multisig with less than 4 members OR Contracts are not verified OR Contracts owner can harm our strategy completely; 4 -> Contracts owner is a multisig but the addresses are not known/hidden OR Contracts owner can harm our strategy by setting parameters in external protocol contracts up to some degree; 3 -> Contracts owner is a multisig with known people but multisig threshold is very low; 2 -> Contracts owner is a multisig with known trusted people; 1 -> Contracts owner is a multisig with known trusted people with Timelock OR Contracts are governanceless, immutable OR Contracts owner can't do any harm to our strategy by setting parameters in external protocol contracts
	ExternalProtocolCentralisation int8 `json:"externalProtocolCentralisation"`
	// ExternalProtocolTvl is the active TVL that the external protocol holds. 5 -> TVL of $10M or less; 4 -> TVL between $10M and $40M;3 -> TVL between $40M and $120M; 2 -> TVL; between $120M and $480M; 1 -> TVL of $480M or morealisation"`
	ExternalProtocolTvl int8 `json:"externalProtocolTvl"`
	// ExternalProtocolLongevity is how long the external protocol contracts in scope have been deployed alive. 5 -> Less than 6 months; 4 -> Between 6 and 12 months; 3 -> Between 12 and 18 months; 2 -> Between 18 and 24 months; 1 -> 24 months or more
	ExternalProtocolLongevity int8 `json:"externalProtocolLongevity"`
	// ExternalProtocolType is a rough estimate of evaluating a protocol's purpose. 5 -> The main expertise of the protocol lies in off-chain operations, such as RWA protocols; 4 -> Cross-chain applications, like cross-chain bridges, cross-chain yield aggregators, and cross-chain lending/borrowing protocols; 3 -> AMM lending/borrowing protocols that are not forks of blue-chip protocols, leveraged farming protocols, as well as newly conceptualized protocols; 2 -> Slightly modified forked blue-chip protocols; 1 -> Blue-chip protocols such as AAVE, Compound, Uniswap, Curve, Convex, and Balancer.ity"`
	ExternalProtocolType int8 `json:"externalProtocolType"`
	// Comment is a comment for the risk score to the strategy. Can be empty.
	Comment string `json:"comment"`
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
	Protocols      []string           `json:"protocols"`      // The Protocols used by the vault
	Inclusion      TInclusion         `json:"inclusion"`      // Inclusion is a special field to know "where" the vault should be displayed.
	RiskLevel      int8               `json:"riskLevel"`      // The risk level of the vault (1 to 5, -1 if not set)
	RiskScore      TRiskScore         `json:"riskScore"`      // The risk score of the vault
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
	LastTotalAssets      *bigNumber.Int   `json:"lastTotalAssets"`      // Total assets locked in the vault (from blockchain or Kong)

	// Kong-sourced data (single source of truth for TVL and debts)
	KongTVL   string `json:"kongTvl,omitempty"`   // TVL from Kong API (tvl.close field)
	KongDebts string `json:"kongDebts,omitempty"` // JSON-encoded debts array from Kong API

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
	PricePerShare struct {
		WeekAgo   float64 `json:"week_ago"`
		MonthAgo  float64 `json:"month_ago"`
		Inception float64 `json:"inception"`
	} `json:"pricePerShare"`
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

// TCmsProtocolType represents the supported protocols
type TCmsProtocolType string

const (
	ProtocolCurve      TCmsProtocolType = "Curve"
	ProtocolBeethovenX TCmsProtocolType = "BeethovenX"
	ProtocolGamma      TCmsProtocolType = "Gamma"
	ProtocolBalancer   TCmsProtocolType = "Balancer"
	ProtocolYearn      TCmsProtocolType = "Yearn"
)

// TVaultCmsMetadataSchema represents the vault metadata structure from ycms
type TVaultCmsMetadataSchema struct {
	ChainID        uint64             `json:"chainId"`
	Address        common.Address     `json:"address"`
	Registry       *common.Address    `json:"registry,omitempty"`
	Kind           *TVaultKind        `json:"kind"`
	Type           *TTokenType        `json:"type"`
	IsRetired      bool               `json:"isRetired"`
	IsHidden       bool               `json:"isHidden"`
	IsAggregator   bool               `json:"isAggregator"`
	IsBoosted      bool               `json:"isBoosted"`
	IsAutomated    bool               `json:"isAutomated"`
	IsHighlighted  bool               `json:"isHighlighted"`
	IsPool         bool               `json:"isPool"`
	ShouldUseV2APR bool               `json:"shouldUseV2APR"`
	Migration      TMigration         `json:"migration"`
	Stability      TStability         `json:"stability"`
	Category       *string            `json:"category,omitempty"`
	DisplayName    *string            `json:"displayName,omitempty"`
	DisplaySymbol  *string            `json:"displaySymbol,omitempty"`
	Description    *string            `json:"description,omitempty"`
	SourceURI      *string            `json:"sourceURI,omitempty"`
	UINotice       *string            `json:"uiNotice,omitempty"`
	Protocols      []TCmsProtocolType `json:"protocols"`
	Inclusion      TInclusion         `json:"inclusion"`
}

type CoercibleUint64 struct {
	Value uint64
}

func (f *CoercibleUint64) UnmarshalJSON(data []byte) error {
	// Try to unmarshal as uint64 first
	if err := json.Unmarshal(data, &f.Value); err == nil {
		return nil
	}

	// Try to unmarshal as string
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	// Convert string to uint64
	if str == "" {
		f.Value = 0
		return nil
	}

	parsed, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return err
	}

	f.Value = parsed
	return nil
}

type TKongDebt struct {
	Strategy           string   `json:"strategy"`
	PerformanceFee     *string  `json:"performanceFee"`
	Activation         *string  `json:"activation"`
	DebtRatio          *string  `json:"debtRatio"`
	MinDebtPerHarvest  *string  `json:"minDebtPerHarvest"`
	MaxDebtPerHarvest  *string  `json:"maxDebtPerHarvest"`
	LastReport         *string  `json:"lastReport"`
	TotalDebt          *string  `json:"totalDebt"`
	TotalDebtUsd       *float64 `json:"totalDebtUsd"`
	TotalGain          *string  `json:"totalGain"`
	TotalGainUsd       *float64 `json:"totalGainUsd"`
	TotalLoss          *string  `json:"totalLoss"`
	TotalLossUsd       *float64 `json:"totalLossUsd"`
	CurrentDebt        *string  `json:"currentDebt"`
	CurrentDebtUsd     *float64 `json:"currentDebtUsd"`
	MaxDebt            *string  `json:"maxDebt"`
	MaxDebtUsd         *float64 `json:"maxDebtUsd"`
	TargetDebtRatio    *float64 `json:"targetDebtRatio"`
	MaxDebtRatio       *float64 `json:"maxDebtRatio"`
}

type KongAPY struct {
	PricePerShare        string   `json:"pricePerShare"`        // BigInt as string
	WeeklyNet            *float64 `json:"weeklyNet"`            // Float or null
	WeeklyPricePerShare  string   `json:"weeklyPricePerShare"`  // BigInt as string
	MonthlyNet           *float64 `json:"monthlyNet"`           // Float or null
	MonthlyPricePerShare string   `json:"monthlyPricePerShare"` // BigInt as string
	InceptionNet         *float64 `json:"inceptionNet"`         // Float or null
	BlockNumber          string   `json:"blockNumber"`          // BigInt as string
	BlockTime            string   `json:"blockTime"`            // BigInt as string
	Decimals             uint64   `json:"decimals"`             // Token decimals for PPS normalization
}

type TKongVaultSchema struct {
	Hook struct {
		Fees struct {
			ManagementFee  uint64 `json:"managementFee"`
			PerformanceFee uint64 `json:"performanceFee"`
		} `json:"fees"`
	} `json:"hook"`
	Snapshot struct {
		ManagementFee  CoercibleUint64 `json:"managementFee"`
		PerformanceFee CoercibleUint64 `json:"performanceFee"`
	} `json:"snapshot"`
	TVL          float64      `json:"tvl"`   // TVL from Kong (tvl.close field)
	Debts []TKongDebt  `json:"debts"` // Debts array from Kong
	APY   KongAPY      `json:"apy"`   // APY from Kong
}
