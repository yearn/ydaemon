package models

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
)

type TVaultType string

const (
	VaultTypeExperimental TVaultType = "Experimental"
	VaultTypeStandard     TVaultType = "Standard"
	VaultTypeAutomated    TVaultType = "Automated"
)

type TVaultsFromRegistry struct {
	RegistryAddress common.Address
	VaultsAddress   common.Address
	TokenAddress    common.Address
	BlockHash       common.Hash
	Type            TVaultType
	APIVersion      string
	BlockNumber     uint64
	Activation      uint64
	ManagementFee   uint64
	TxIndex         uint
	LogIndex        uint
}

// TTVL holds the info about the value locked in a vault
type TTVL struct {
	TotalAssets          *bigNumber.Int `json:"total_assets"`
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
}

// TMigration helps us to know if a vault is in the process of being migrated.
type TMigration struct {
	Available bool           `json:"available"`
	Address   common.Address `json:"address"`
	Contract  common.Address `json:"contract"`
}

// TStaking holds some metadata about the staking contract.
type TStaking struct {
	Available bool           `json:"available"`
	Address   common.Address `json:"address"`
	Risk      int            `json:"risk"`
	TVL       float64        `json:"tvl"`
}

// TVault is the main structure returned by the API when trying to get all the vaults for a specific network
type TVault struct {
	Address               common.Address   `json:"address"`
	Management            common.Address   `json:"management"`
	Governance            common.Address   `json:"governance"`
	Guardian              common.Address   `json:"guardian"`
	Rewards               common.Address   `json:"rewards"`
	WithdrawalQueue       []common.Address `json:"withdrawalQueue"`
	PricePerShare         *bigNumber.Int   `json:"pricePerShare"`
	DepositLimit          *bigNumber.Int   `json:"depositLimit"`
	AvailableDepositLimit *bigNumber.Int   `json:"availableDepositLimit,omitempty"`
	TotalAssets           *bigNumber.Int   `json:"total_assets"`
	Type                  TVaultType       `json:"type"`
	Symbol                string           `json:"symbol"`
	DisplaySymbol         string           `json:"display_symbol"`
	FormatedSymbol        string           `json:"formated_symbol"`
	Name                  string           `json:"name"`
	DisplayName           string           `json:"display_name"`
	FormatedName          string           `json:"formated_name"`
	Icon                  string           `json:"icon"`
	Version               string           `json:"version"`
	ChainID               uint64           `json:"chainID"`
	Inception             uint64           `json:"inception"`  //Timestamp
	Activation            uint64           `json:"activation"` //BlockNumber
	Decimals              uint64           `json:"decimals"`
	PerformanceFee        uint64           `json:"performanceFee"`
	ManagementFee         uint64           `json:"managementFee"`
	Endorsed              bool             `json:"endorsed"`
	EmergencyShutdown     bool             `json:"emergency_shutdown"`
	Token                 TERC20Token      `json:"token"`
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

type TInternalVaultFromMetaClassification struct {
	IsAutomated     bool   `json:"isAutomated"`
	IsPool          bool   `json:"isPool"`
	PoolProvider    string `json:"poolProvider,omitempty"`
	Stability       string `json:"stability"`
	StableBaseAsset string `json:"stableBaseAsset,omitempty"`
}

// TInternalVaultFromMeta is the structure of data we receive when calling meta.yearn.finance/api/1/vaults/all
type TInternalVaultFromMeta struct {
	Address              common.Address                       `json:"address"`
	MigrationTargetVault common.Address                       `json:"migrationTargetVault"`
	MigrationContract    common.Address                       `json:"migrationContract"`
	DisplayName          string                               `json:"displayName"`
	Comment              string                               `json:"comment"`
	APYTypeOverride      string                               `json:"apyTypeOverride"`
	APYOverride          float64                              `json:"apyOverride"`
	Order                float32                              `json:"order"`
	ChainID              uint64                               `json:"chainID"`
	HideAlways           bool                                 `json:"hideAlways"`
	DepositsDisabled     bool                                 `json:"depositsDisabled"`
	WithdrawalsDisabled  bool                                 `json:"withdrawalsDisabled"`
	MigrationAvailable   bool                                 `json:"migrationAvailable"`
	AllowZapIn           bool                                 `json:"allowZapIn"`
	AllowZapOut          bool                                 `json:"allowZapOut"`
	Retired              bool                                 `json:"retired"`
	Classification       TInternalVaultFromMetaClassification `json:"classification"`
}
