package vaults

import (
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/types/common"
	"github.com/yearn/ydaemon/external/strategies"
)

// TToken holds the info about the underlying (or want) token of a vault
type TToken struct {
	Address     common.Address `json:"address"`
	Name        string         `json:"name"`
	DisplayName string         `json:"display_name"`
	Symbol      string         `json:"symbol"`
	Description string         `json:"description"`
	Decimals    uint64         `json:"decimals"`
	Icon        string         `json:"icon"`
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
	Type      string        `json:"type"`
	GrossAPR  float64       `json:"gross_apr"`
	NetAPY    float64       `json:"net_apy"`
	Fees      TAPYFees      `json:"fees"`
	Points    TAPYPoints    `json:"points"`
	Composite TAPYComposite `json:"composite"`
}

// TMigration helps us to know if a vault is in the process of being migrated.
type TMigration struct {
	Available bool           `json:"available"`
	Address   common.Address `json:"address"`
}

//TVault details holds some extra information about the vault.
type TVaultDetails struct {
	Management            common.Address `json:"management"`
	Governance            common.Address `json:"governance"`
	Guardian              common.Address `json:"guardian"`
	Rewards               common.Address `json:"rewards"`
	DepositLimit          *bigNumber.Int `json:"depositLimit"`
	AvailableDepositLimit *bigNumber.Int `json:"availableDepositLimit,omitempty"`
	Comment               string         `json:"comment"`
	APYTypeOverride       string         `json:"apyTypeOverride"`
	APYOverride           float64        `json:"apyOverride"`
	Order                 float32        `json:"-"`
	PerformanceFee        uint64         `json:"performanceFee"`
	ManagementFee         uint64         `json:"managementFee"`
	DepositsDisabled      bool           `json:"depositsDisabled"`
	WithdrawalsDisabled   bool           `json:"withdrawalsDisabled"`
	AllowZapIn            bool           `json:"allowZapIn"`
	AllowZapOut           bool           `json:"allowZapOut"`
	Retired               bool           `json:"retired"`
}

// TVault is the main structure returned by the API when trying to get all the vaults for a specific network
type TVault struct {
	Address            common.Address         `json:"address"`
	Symbol             string                 `json:"symbol"`
	DisplaySymbol      string                 `json:"display_symbol"`
	FormatedSymbol     string                 `json:"formated_symbol"`
	Name               string                 `json:"name"`
	DisplayName        string                 `json:"display_name"`
	FormatedName       string                 `json:"formated_name"`
	Icon               string                 `json:"icon"`
	Version            string                 `json:"version"`
	Type               string                 `json:"type"`
	Inception          uint64                 `json:"inception"`
	Decimals           uint64                 `json:"decimals"`
	Updated            uint64                 `json:"updated"`
	Endorsed           bool                   `json:"endorsed"`
	Emergency_shutdown bool                   `json:"emergency_shutdown"`
	Token              TToken                 `json:"token"`
	TVL                TTVL                   `json:"tvl"`
	APY                TAPY                   `json:"apy"`
	Strategies         []strategies.TStrategy `json:"strategies"`
	Migration          TMigration             `json:"migration"`
	Details            *TVaultDetails         `json:"details"`
}

// TVaultHarvest is the struct containing the information about the harvest of a vault that can be used to compute the Gain/Loss and access the Transactions on the explorer.
type TVaultHarvest struct {
	VaultAddress    string  `json:"vaultAddress,omitempty"`
	StrategyAddress string  `json:"strategyAddress"`
	TxHash          string  `json:"txHash"`
	Timestamp       string  `json:"timestamp"`
	Profit          string  `json:"profit"`
	ProfitValue     float64 `json:"profitValue"`
	Loss            string  `json:"loss"`
	LossValue       float64 `json:"lossValue"`
}
