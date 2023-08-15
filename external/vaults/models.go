package vaults

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/internal/meta"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/tokens"
	"github.com/yearn/ydaemon/internal/vaults"
	// "github.com/yearn/ydaemon/processes/apy"
)

// TExternalVaultHarvest is the struct containing the information about the harvest of a vault that can be used to compute the Gain/Loss and access the Transactions on the explorer.
type TExternalVaultHarvest struct {
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
	KeepVelo    float64 `json:"keep_velo"`
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
	Type              string                `json:"type"`
	GrossAPR          float64               `json:"gross_apr"`
	NetAPY            float64               `json:"net_apy"`
	StakingRewardsAPR float64               `json:"staking_rewards_apr"`
	Fees              TExternalAPYFees      `json:"fees"`
	Points            TExternalAPYPoints    `json:"points"`
	Composite         TExternalAPYComposite `json:"composite"`
	Error             string                `json:"error"`
}

// TExternalVaultMigration is the struct containing the information about the migration of a vault.
type TExternalVaultMigration struct {
	Available bool   `json:"available"`
	Address   string `json:"address"`
	Contract  string `json:"contract"`
}

// TExternalVaultStaking is the struct containing the information about the staking contract related to that vault.
type TExternalVaultStaking struct {
	Available bool    `json:"available"`
	Address   string  `json:"address"`
	TVL       float64 `json:"tvl"`
	Risk      int     `json:"risk"`
}

// TExternalVaultDetails is the struct containing the information about a vault.
type TExternalVaultDetails struct {
	Management            string         `json:"management"`
	Governance            string         `json:"governance"`
	Guardian              string         `json:"guardian"`
	Rewards               string         `json:"rewards"`
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

// TExternalERC20Token contains the basic information of an ERC20 token
type TExternalERC20Token struct {
	Address                   string            `json:"address"`
	UnderlyingTokensAddresses []string          `json:"underlyingTokensAddresses"`
	Name                      string            `json:"name"`
	Symbol                    string            `json:"symbol"`
	Type                      models.TTokenType `json:"type"`
	DisplayName               string            `json:"display_name"`
	DisplaySymbol             string            `json:"display_symbol"`
	Description               string            `json:"description"`
	Icon                      string            `json:"icon"`
	Decimals                  uint64            `json:"decimals"`
}

// TExternalVault is the struct containing the information about a vault.
type TExternalVault struct {
	Address           string              `json:"address"`
	Type              models.TVaultType   `json:"type"`
	Symbol            string              `json:"symbol"`
	DisplaySymbol     string              `json:"display_symbol"`
	FormatedSymbol    string              `json:"formated_symbol"`
	Name              string              `json:"name"`
	DisplayName       string              `json:"display_name"`
	FormatedName      string              `json:"formated_name"`
	Icon              string              `json:"icon"`
	Version           string              `json:"version"`
	Category          string              `json:"category"`
	Inception         uint64              `json:"inception"`
	Decimals          uint64              `json:"decimals"`
	ChainID           uint64              `json:"chainID"`
	RiskScore         float64             `json:"riskScore"`
	Endorsed          bool                `json:"endorsed"`
	EmergencyShutdown bool                `json:"emergency_shutdown"`
	Token             TExternalERC20Token `json:"token"`
	TVL               TExternalVaultTVL   `json:"tvl"`
	APY               TExternalVaultAPY   `json:"apy"`
	// NewAPY            apy.TAPIV1APY           `json:"newApy"`
	Details    *TExternalVaultDetails  `json:"details"`
	Strategies []*TStrategy            `json:"strategies"`
	Migration  TExternalVaultMigration `json:"migration"`
	Staking    TExternalVaultStaking   `json:"staking"`
}

func NewVault() *TExternalVault {
	return &TExternalVault{}
}
func (v *TExternalVault) AssignTVault(internalVault models.TVault) *TExternalVault {
	vaultFromMeta, ok := meta.GetMetaVault(internalVault.ChainID, internalVault.Address)
	if !ok {
		vaultFromMeta = &models.TInternalVaultFromMeta{
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

	v.Address = internalVault.Address.Hex()
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
	v.TVL = TExternalVaultTVL(vaults.BuildTVL(internalVault))
	v.Migration = toTExternalVaultMigration(vaults.BuildMigration(internalVault))
	v.Staking = toTExternalVaultStaking(vaults.BuildStaking(internalVault))
	v.Category = vaults.BuildCategory(internalVault)

	underlyingToken, ok := tokens.FindUnderlyingForVault(internalVault.ChainID, internalVault.Address)
	if ok {
		v.Token = TExternalERC20Token{
			Address:                   underlyingToken.Address.Hex(),
			UnderlyingTokensAddresses: toArrTMixedcaseAddress(underlyingToken.UnderlyingTokensAddresses),
			Name:                      underlyingToken.Name,
			Symbol:                    underlyingToken.Symbol,
			Type:                      underlyingToken.Type,
			DisplayName:               underlyingToken.DisplayName,
			DisplaySymbol:             underlyingToken.DisplaySymbol,
			Description:               underlyingToken.Description,
			Icon:                      underlyingToken.Icon,
			Decimals:                  underlyingToken.Decimals,
		}
	} else {
		v.Token = TExternalERC20Token{
			Address:                   internalVault.Token.Address.Hex(),
			UnderlyingTokensAddresses: toArrTMixedcaseAddress(internalVault.Token.UnderlyingTokensAddresses),
			Name:                      internalVault.Token.Name,
			Symbol:                    internalVault.Token.Symbol,
			Type:                      internalVault.Token.Type,
			DisplayName:               internalVault.Token.DisplayName,
			DisplaySymbol:             internalVault.Token.DisplaySymbol,
			Description:               internalVault.Token.Description,
			Icon:                      internalVault.Token.Icon,
			Decimals:                  internalVault.Token.Decimals,
		}
	}
	if v.Token.UnderlyingTokensAddresses == nil {
		v.Token.UnderlyingTokensAddresses = []string{}
	}

	aggregatedVault, okLegacyAPI := GetAggregatedVault(v.ChainID, addresses.ToAddress(v.Address).Hex())
	internalAPY := vaults.BuildAPY(internalVault, aggregatedVault, okLegacyAPI)
	v.APY = TExternalVaultAPY{
		Type:              internalAPY.Type,
		GrossAPR:          internalAPY.GrossAPR,
		NetAPY:            internalAPY.NetAPY,
		StakingRewardsAPR: internalAPY.StakingRewardsAPR,
		Fees:              TExternalAPYFees(internalAPY.Fees),
		Points:            TExternalAPYPoints(internalAPY.Points),
		Composite:         TExternalAPYComposite(internalAPY.Composite),
		Error:             internalAPY.Error,
	}
	// v.NewAPY = apy.COMPUTED_APR[internalVault.ChainID][internalVault.Address]

	v.Details = &TExternalVaultDetails{
		Management:            internalVault.Management.Hex(),
		Governance:            internalVault.Governance.Hex(),
		Guardian:              internalVault.Guardian.Hex(),
		Rewards:               internalVault.Rewards.Hex(),
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
		if addresses.Equals(strat.Address, `0xE7863292dd8eE5d215eC6D75ac00911D06E59B2d`) {
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

func toTExternalVaultMigration(migration models.TMigration) TExternalVaultMigration {
	return TExternalVaultMigration{
		Available: migration.Available,
		Address:   migration.Address.Hex(),
		Contract:  migration.Contract.Hex(),
	}
}
func toTExternalVaultStaking(staking models.TStaking) TExternalVaultStaking {
	return TExternalVaultStaking{
		Available: staking.Available,
		Address:   staking.Address.Hex(),
		Risk:      staking.Risk,
		TVL:       staking.TVL,
	}
}
func toArrTMixedcaseAddress(addresses []common.Address) []string {
	arr := make([]string, len(addresses))
	for i, address := range addresses {
		arr[i] = address.Hex()
	}
	return arr
}
