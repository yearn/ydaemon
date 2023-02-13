package vaults

import (
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/types/common"
	"github.com/yearn/ydaemon/internal/meta"
	"github.com/yearn/ydaemon/internal/tokens"
	"github.com/yearn/ydaemon/internal/utils"
	"github.com/yearn/ydaemon/internal/vaults"
)

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

// TExternalVaultTVL is the struct containing the information about the TVL of a vault.
type TExternalVaultTVL struct {
	TotalAssets          *bigNumber.Int `json:"total_assets"`
	TotalDelegatedAssets *bigNumber.Int `json:"total_delegated_assets"`
	TVLDeposited         float64        `json:"tvl_deposited"`
	TVLDelegated         float64        `json:"tvl_delegated"`
	TVL                  float64        `json:"tvl"`
	Price                float64        `json:"price"`
}

// TExternalAPYFees holds the fees information about this vault.
type TExternalAPYFees struct {
	Performance float64 `json:"performance"`
	Withdrawal  float64 `json:"withdrawal"`
	Management  float64 `json:"management"`
	KeepCRV     float64 `json:"keep_crv"`
	CvxKeepCRV  float64 `json:"cvx_keep_crv"`
}

// TExternalAPYPoints holds the points information about this vault.
type TExternalAPYPoints struct {
	WeekAgo   float64 `json:"week_ago"`
	MonthAgo  float64 `json:"month_ago"`
	Inception float64 `json:"inception"`
}

// TExternalAPYComposite holds the points information about this vault.
type TExternalAPYComposite struct {
	Boost      float64 `json:"boost"`
	PoolAPY    float64 `json:"pool_apy"`
	BoostedAPR float64 `json:"boosted_apr"`
	BaseAPR    float64 `json:"base_apr"`
	CvxAPR     float64 `json:"cvx_apr"`
	RewardsAPR float64 `json:"rewards_apr"`
}

// TExternalVaultAPY holds the APY information about this vault.
type TExternalVaultAPY struct {
	Type      string                `json:"type"`
	GrossAPR  float64               `json:"gross_apr"`
	NetAPY    float64               `json:"net_apy"`
	Fees      TExternalAPYFees      `json:"fees"`
	Points    TExternalAPYPoints    `json:"points"`
	Composite TExternalAPYComposite `json:"composite"`
}

// TExternalVaultMigration is the struct containing the information about the migration of a vault.
type TExternalVaultMigration struct {
	Available bool           `json:"available"`
	Address   common.Address `json:"address"`
	Contract  common.Address `json:"contract"`
}

// TExternalVaultDetails is the struct containing the information about a vault.
type TExternalVaultDetails struct {
	Management            common.Address `json:"management"`
	Governance            common.Address `json:"governance"`
	Guardian              common.Address `json:"guardian"`
	Rewards               common.Address `json:"rewards"`
	DepositLimit          *bigNumber.Int `json:"depositLimit"`
	AvailableDepositLimit *bigNumber.Int `json:"availableDepositLimit,omitempty"`
	Comment               string         `json:"comment"`
	APYTypeOverride       string         `json:"apyTypeOverride"`
	APYOverride           float64        `json:"apyOverride"`
	Order                 float32        `json:"order"`
	PerformanceFee        uint64         `json:"performanceFee"`
	ManagementFee         uint64         `json:"managementFee"`
	DepositsDisabled      bool           `json:"depositsDisabled"`
	WithdrawalsDisabled   bool           `json:"withdrawalsDisabled"`
	AllowZapIn            bool           `json:"allowZapIn"`
	AllowZapOut           bool           `json:"allowZapOut"`
	Retired               bool           `json:"retired"`
	HideAlways            bool           `json:"hideAlways"`
}

// TExternalVault is the struct containing the information about a vault.
type TExternalVault struct {
	Address           common.Address          `json:"address"`
	Type              utils.TVaultType        `json:"type"`
	Symbol            string                  `json:"symbol"`
	DisplaySymbol     string                  `json:"display_symbol"`
	FormatedSymbol    string                  `json:"formated_symbol"`
	Name              string                  `json:"name"`
	DisplayName       string                  `json:"display_name"`
	FormatedName      string                  `json:"formated_name"`
	Icon              string                  `json:"icon"`
	Version           string                  `json:"version"`
	Category          string                  `json:"category"`
	Inception         uint64                  `json:"inception"`
	Decimals          uint64                  `json:"decimals"`
	ChainID           uint64                  `json:"chainID"`
	RiskScore         float64                 `json:"riskScore"`
	Endorsed          bool                    `json:"endorsed"`
	EmergencyShutdown bool                    `json:"emergency_shutdown"`
	Token             tokens.TERC20Token      `json:"token"`
	TVL               TExternalVaultTVL       `json:"tvl"`
	APY               TExternalVaultAPY       `json:"apy"`
	Details           *TExternalVaultDetails  `json:"details"`
	Strategies        []*TStrategy            `json:"strategies"`
	Migration         TExternalVaultMigration `json:"migration"`
}

func NewVault() *TExternalVault {
	return &TExternalVault{}
}
func (v *TExternalVault) AssignTVault(internalVault *vaults.TVault) *TExternalVault {
	vaultFromMeta, ok := meta.GetMetaVault(internalVault.ChainID, common.FromAddress(internalVault.Address))
	if !ok {
		vaultFromMeta = &meta.TVaultFromMeta{
			Order:               1000000000,
			HideAlways:          false,
			DepositsDisabled:    false,
			WithdrawalsDisabled: false,
			MigrationAvailable:  false,
			AllowZapIn:          true,
			AllowZapOut:         true,
			Retired:             false,
		}
	}

	v.Address = common.FromAddress(internalVault.Address)
	v.Symbol = internalVault.Symbol
	v.DisplaySymbol = internalVault.DisplaySymbol
	v.FormatedSymbol = internalVault.FormatedSymbol
	v.Name = internalVault.Name
	v.DisplayName = internalVault.DisplayName
	v.FormatedName = internalVault.FormatedName
	v.Icon = internalVault.Icon
	v.Version = internalVault.Version
	v.Type = internalVault.Type
	v.Inception = internalVault.Inception
	v.Decimals = internalVault.Decimals
	v.Endorsed = internalVault.Endorsed
	v.EmergencyShutdown = internalVault.EmergencyShutdown
	v.ChainID = internalVault.ChainID
	v.TVL = TExternalVaultTVL(internalVault.BuildTVL())
	v.Migration = TExternalVaultMigration(internalVault.BuildMigration())
	v.Category = internalVault.BuildCategory()

	underlyingToken, ok := tokens.FindUnderlyingForVault(internalVault.ChainID, common.FromAddress(internalVault.Address))
	if ok {
		v.Token = *underlyingToken
	} else {
		v.Token = internalVault.Token
	}
	if v.Token.UnderlyingTokensAddresses == nil {
		v.Token.UnderlyingTokensAddresses = []ethcommon.Address{}
	}

	internalAPY := internalVault.BuildAPY()
	v.APY = TExternalVaultAPY{
		Type:      internalAPY.Type,
		GrossAPR:  internalAPY.GrossAPR,
		NetAPY:    internalAPY.NetAPY,
		Fees:      TExternalAPYFees(internalAPY.Fees),
		Points:    TExternalAPYPoints(internalAPY.Points),
		Composite: TExternalAPYComposite(internalAPY.Composite),
	}

	v.Details = &TExternalVaultDetails{
		Management:            internalVault.Management,
		Governance:            internalVault.Governance,
		Guardian:              internalVault.Guardian,
		Rewards:               internalVault.Rewards,
		DepositLimit:          internalVault.DepositLimit,
		AvailableDepositLimit: internalVault.AvailableDepositLimit,
		PerformanceFee:        internalVault.PerformanceFee,
		ManagementFee:         internalVault.ManagementFee,
		Comment:               vaultFromMeta.Comment,
		Order:                 vaultFromMeta.Order,
		DepositsDisabled:      vaultFromMeta.DepositsDisabled,
		WithdrawalsDisabled:   vaultFromMeta.WithdrawalsDisabled,
		AllowZapIn:            vaultFromMeta.AllowZapIn,
		AllowZapOut:           vaultFromMeta.AllowZapOut,
		Retired:               vaultFromMeta.Retired,
		APYTypeOverride:       vaultFromMeta.APYTypeOverride,
		APYOverride:           vaultFromMeta.APYOverride,
		HideAlways:            vaultFromMeta.HideAlways,
	}

	return v
}

func (v *TExternalVault) ComputeRiskScore() float64 {
	totalRiskScore := bigNumber.NewFloat(0)
	for _, strat := range v.Strategies {
		totalDebt := bigNumber.NewFloat(0).SetInt(strat.Details.TotalDebt)
		//Specific overwrite for st-yCRV strategy
		if strat.Address == common.HexToAddress(`0xE7863292dd8eE5d215eC6D75ac00911D06E59B2d`) {
			totalDebt = bigNumber.NewFloat(0)
		}

		totalRiskScore = bigNumber.NewFloat(0).Add(
			totalRiskScore,
			bigNumber.NewFloat(0).Mul(
				bigNumber.NewFloat(strat.Risk.RiskScore),
				totalDebt,
			),
		)
	}
	vaultRiskScore, _ := bigNumber.NewFloat(0).Div(
		totalRiskScore,
		bigNumber.NewFloat().SetInt(v.TVL.TotalAssets),
	).Float64()

	return vaultRiskScore
}
