package vaults

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/fetcher"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/risk"
	"github.com/yearn/ydaemon/internal/storage"
	"github.com/yearn/ydaemon/processes/apr"
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

// TExternalAPYFees holds the fees information about this vault.
type TExternalAPYFees struct {
	Performance float64 `json:"performance"`
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
	IsRetired       bool   `json:"isRetired"`    // If the vault is retired or not
	IsHidden        bool   `json:"isHidden"`     // If the vault is hidden or not
	IsAggregator    bool   `json:"isAggregator"` // If the vault should be treated as an aggregator vault (aka multi-strategy) even if he only has one strategy
	IsAutomated     bool   `json:"isAutomated"`
	IsPool          bool   `json:"isPool"`
	PoolProvider    string `json:"poolProvider,omitempty"`
	Stability       string `json:"stability"`
	StableBaseAsset string `json:"stableBaseAsset,omitempty"`
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
	Address           string                  `json:"address"`
	Type              models.TTokenType       `json:"type"`
	Kind              models.TVaultKind       `json:"kind"`
	Symbol            string                  `json:"symbol"`
	DisplaySymbol     string                  `json:"display_symbol"`
	FormatedSymbol    string                  `json:"formated_symbol"`
	Name              string                  `json:"name"`
	DisplayName       string                  `json:"display_name"`
	FormatedName      string                  `json:"formated_name"`
	Description       string                  `json:"description,omitempty"`
	Icon              string                  `json:"icon"`
	Version           string                  `json:"version"`
	Category          string                  `json:"category"`
	Decimals          uint64                  `json:"decimals"`
	ChainID           uint64                  `json:"chainID"`
	RiskScore         float64                 `json:"riskScore"`
	Endorsed          bool                    `json:"endorsed"`
	EmergencyShutdown bool                    `json:"emergency_shutdown"`
	Token             TExternalERC20Token     `json:"token"`
	TVL               models.TTVL             `json:"tvl"`
	APR               apr.TVaultAPR           `json:"apr"`
	Details           TExternalVaultDetails   `json:"details"`
	Strategies        []TStrategy             `json:"strategies"`
	Migration         TExternalVaultMigration `json:"migration"`
	Staking           TExternalVaultStaking   `json:"staking"`
	// Computing only
	FeaturingScore float64 `json:"featuringScore"`
}

type TSimplifiedExternalVaultTVL struct {
	TotalAssets *bigNumber.Int `json:"totalAssets"`
	TVL         float64        `json:"tvl"`
	Price       float64        `json:"price"`
}

type TSimplifiedExternalERC20Token struct {
	Address     string `json:"address"`
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	Description string `json:"description"`
	Decimals    uint64 `json:"decimals"`
}

type TStakingData struct {
	Address   string `json:"address"`
	Available bool   `json:"available"`
}

// TSimplifiedExternalVault is the struct containing the information about a vault.
type TSimplifiedExternalVault struct {
	Address        string                        `json:"address"`
	Type           models.TTokenType             `json:"type"`
	Kind           models.TVaultKind             `json:"kind"`
	Symbol         string                        `json:"symbol"`
	Name           string                        `json:"name"`
	Category       string                        `json:"category"`
	Version        string                        `json:"version"`
	Description    string                        `json:"description,omitempty"`
	Decimals       uint64                        `json:"decimals"`
	ChainID        uint64                        `json:"chainID"`
	Retired        bool                          `json:"retired"`
	Token          TSimplifiedExternalERC20Token `json:"token"`
	TVL            TSimplifiedExternalVaultTVL   `json:"tvl"`
	APR            apr.TVaultAPR                 `json:"apr"`
	Strategies     []TStrategy                   `json:"strategies"`
	Staking        TStakingData                  `json:"staking,omitempty"`
	Migration      TExternalVaultMigration       `json:"migration,omitempty"`
	FeaturingScore float64                       `json:"featuringScore"`
}

func NewVault() TExternalVault {
	return TExternalVault{}
}
func (v TExternalVault) AssignTVault(vault models.TVault) (TExternalVault, error) {
	vaultToken, ok := storage.GetERC20(vault.ChainID, vault.Address)
	if !ok {
		return v, errors.New(`token not found`)
	}
	name, displayName, formatedName := fetcher.BuildVaultNames(vault, vault.DisplayName)
	symbol, displaySymbol, formatedSymbol := fetcher.BuildVaultSymbol(vault, vault.DisplaySymbol)

	v.Address = vault.Address.Hex()
	v.Version = vault.Version
	v.Endorsed = vault.Endorsed
	v.EmergencyShutdown = vault.EmergencyShutdown
	v.ChainID = vault.ChainID
	v.TVL = fetcher.BuildVaultTVL(vault)
	v.Migration = toTExternalVaultMigration(vault.Migration)
	v.Staking = toTExternalVaultStaking(risk.BuildVaultStaking(vault))
	v.Category = fetcher.BuildVaultCategory(vault)
	v.Symbol = symbol
	v.DisplaySymbol = displaySymbol
	v.FormatedSymbol = formatedSymbol
	v.Name = name
	v.DisplayName = displayName
	v.FormatedName = formatedName
	v.Icon = vaultToken.Icon
	v.Type = vaultToken.Type
	v.Kind = vault.Kind
	v.Decimals = vaultToken.Decimals
	v.Description = vault.Description

	underlyingToken, ok := storage.GetUnderlyingERC20(vault.ChainID, vault.Address)
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
		underlyingToken, ok = storage.GetERC20(vault.ChainID, vault.AssetAddress)
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
		}
	}
	if v.Token.UnderlyingTokensAddresses == nil {
		v.Token.UnderlyingTokensAddresses = []string{}
	}

	v.APR = apr.COMPUTED_APR[vault.ChainID][vault.Address]
	v.Details = TExternalVaultDetails{
		IsRetired:       vault.IsRetired,
		IsHidden:        vault.IsHidden,
		IsAggregator:    vault.IsAggregator,
		IsAutomated:     vault.Classification.IsAutomated,
		IsPool:          vault.Classification.IsPool,
		PoolProvider:    vault.Classification.PoolProvider,
		Stability:       helpers.SafeString(vault.Classification.Stability, `Unknown`),
		StableBaseAsset: vault.Classification.StableBaseAsset,
	}

	return v, nil
}

func (v TExternalVault) ComputeRiskScore() float64 {
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
		Address:   migration.Target.Hex(),
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
