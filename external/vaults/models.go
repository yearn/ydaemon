package vaults

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/internal/fetcher"
	"github.com/yearn/ydaemon/internal/models"
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

// TExternalVaultMigration is the struct containing the information about the migration of a vault.
type TExternalVaultMigration struct {
	Available bool   `json:"available"`
	Address   string `json:"address"`
	Contract  string `json:"contract"`
}

// TExternalVaultDetails is the struct containing the information about a vault.
type TExternalVaultDetails struct {
	IsRetired       bool                       `json:"isRetired"`    // If the vault is retired or not
	IsHidden        bool                       `json:"isHidden"`     // If the vault is hidden or not
	IsAggregator    bool                       `json:"isAggregator"` // If the vault should be treated as an aggregator vault (aka multi-strategy) even if he only has one strategy
	IsBoosted       bool                       `json:"isBoosted"`    // If the vault is boosted or not. Humanly set.
	IsAutomated     bool                       `json:"isAutomated"`
	IsHighlighted   bool                       `json:"isHighlighted"`
	IsPool          bool                       `json:"isPool"`
	PoolProvider    string                     `json:"poolProvider,omitempty"`
	Stability       models.TVaultStabilityType `json:"stability"`
	Category        models.TVaultCategoryType  `json:"category"` // The category of the vault
	StableBaseAsset string                     `json:"stableBaseAsset,omitempty"`
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

// TExternalVaultInfo is the struct containing the information about a vault.
type TExternalVaultInfo struct {
	SourceURL     string   `json:"sourceURL,omitempty"` // The vault might require some specific tokens that needs to be bought by a specific provider. It's the URL of the provider.
	RiskLevel     int8     `json:"riskLevel"`           // The risk level of the vault. The value is a calculated from the sum of all risk score from the object for Single Strategy Vaults. Multi-Strategy Vault, highest `riskLevel` of all strategies is set. 1 is the most secure and 5 is the least secure.
	UINotice      string   `json:"uiNotice,omitempty"`  // The notice to display in the UI
	IsRetired     bool     `json:"isRetired"`
	IsBoosted     bool     `json:"isBoosted"`
	IsHighlighted bool     `json:"isHighlighted"`
	RiskScore     [11]int8 `json:"riskScore"` // All risk scores of the Single Strategy Vault. Multi-Strategy Vault won't have this object because its risk score is combination of multiple vaults. For risk value use `riskLevel`. (empty for Multi-Strategy Vault). Array of 11 integers: [review, testing, complexity, riskExposure, protocolIntegration, centralizationRisk, externalProtocolAudit, externalProtocolCentralisation, externalProtocolTvl, externalProtocolLongevity, externalProtocolType]
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
	Endorsed          bool                    `json:"endorsed"`
	Boosted           bool                    `json:"boosted"`
	EmergencyShutdown bool                    `json:"emergency_shutdown"`
	Token             TExternalERC20Token     `json:"token"`
	TVL               models.TTVL             `json:"tvl"`
	APR               apr.TVaultAPR           `json:"apr"`
	Details           TExternalVaultDetails   `json:"details"`
	Strategies        []TStrategy             `json:"strategies"`
	Migration         TExternalVaultMigration `json:"migration"`
	Staking           TStakingData            `json:"staking"`
	Info              TExternalVaultInfo      `json:"info,omitempty"`
	FeaturingScore    float64                 `json:"featuringScore"` // Computing only
	PricePerShare     *bigNumber.Int          `json:"pricePerShare"`
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

type TStakingRewardsData struct {
	Address    string           `json:"address"`
	Name       string           `json:"name"`
	Symbol     string           `json:"symbol"`
	Decimals   uint64           `json:"decimals"`
	Price      float64          `json:"price"`
	IsFinished bool             `json:"isFinished"`
	APR        *bigNumber.Float `json:"apr"`
	PerWeek    *bigNumber.Float `json:"perWeek"`
}
type TStakingData struct {
	Address   string                `json:"address"`
	Available bool                  `json:"available"`
	Source    string                `json:"source"`
	Rewards   []TStakingRewardsData `json:"rewards"`
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
	Token          TSimplifiedExternalERC20Token `json:"token"`
	TVL            TSimplifiedExternalVaultTVL   `json:"tvl"`
	APR            apr.TVaultAPR                 `json:"apr"`
	Strategies     []TStrategy                   `json:"strategies"`
	Staking        TStakingData                  `json:"staking,omitempty"`
	Migration      TExternalVaultMigration       `json:"migration,omitempty"`
	FeaturingScore float64                       `json:"featuringScore"`
	PricePerShare  *bigNumber.Int                `json:"pricePerShare"`
	Info           TExternalVaultInfo            `json:"info,omitempty"`
}

func NewVault() TExternalVault {
	return TExternalVault{}
}
func (v TExternalVault) AssignTVault(vault models.TVault) (TExternalVault, error) {
	vaultToken, ok := storage.GetERC20(vault.ChainID, vault.Address)
	if !ok {
		return v, errors.New(`token not found`)
	}
	name, displayName, formatedName := fetcher.BuildVaultNames(vault, vault.Metadata.DisplayName)
	symbol, displaySymbol, formatedSymbol := fetcher.BuildVaultSymbol(vault, vault.Metadata.DisplaySymbol)
	strategies, _ := storage.ListStrategiesForVault(vault.ChainID, vault.Address)

	v.Address = vault.Address.Hex()
	v.Version = vault.Version
	v.Endorsed = vault.Endorsed
	v.Boosted = vault.Metadata.IsBoosted
	v.EmergencyShutdown = vault.EmergencyShutdown
	v.ChainID = vault.ChainID
	v.TVL = fetcher.BuildVaultTVL(vault)
	v.Migration = toTExternalVaultMigration(vault.Metadata.Migration)
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
	v.Description = vault.Metadata.Description
	v.Category = fetcher.BuildVaultCategory(vault, strategies)
	v.PricePerShare = vault.LastPricePerShare
	v.Staking = assignStakingData(vault.ChainID, vault.Address)

	if underlyingToken, ok := storage.GetUnderlyingERC20(vault.ChainID, vault.Address); ok {
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
	} else if underlyingToken, ok = storage.GetERC20(vault.ChainID, vault.AssetAddress); ok {
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

	if v.Token.UnderlyingTokensAddresses == nil {
		v.Token.UnderlyingTokensAddresses = []string{}
	}

	asyncAPR, ok := apr.GetComputedAPR(vault.ChainID, vault.Address)
	if ok {
		v.APR = asyncAPR.(apr.TVaultAPR)
	}
	v.Details = TExternalVaultDetails{
		IsRetired:       vault.Metadata.IsRetired,
		IsHidden:        vault.Metadata.IsHidden,
		IsAggregator:    vault.Metadata.IsAggregator,
		IsBoosted:       vault.Metadata.IsBoosted,
		IsAutomated:     vault.Metadata.IsAutomated,
		IsHighlighted:   vault.Metadata.IsHighlighted,
		IsPool:          vault.Metadata.IsPool,
		Category:        vault.Metadata.Category,
		Stability:       vault.Metadata.Stability.Stability,
		StableBaseAsset: vault.Metadata.Stability.StableBaseAsset,
	}
	if v.Details.Stability == `` {
		v.Details.Stability = models.VaultStabilityUnknown
	}
	if v.Details.Category == `` {
		v.Details.Category = models.VaultCategoryAutomatic
	}
	if len(vault.Metadata.Protocols) > 0 {
		v.Details.PoolProvider = vault.Metadata.Protocols[0]
	}

	poolResult, found := storage.GetGauge(vault.ChainID, vault.AssetAddress)
	if found && poolResult.PoolURLs.Deposit != nil && len(poolResult.PoolURLs.Deposit) > 0 {
		v.Info.SourceURL = poolResult.PoolURLs.Deposit[0]
	}

	if v.Info.SourceURL == `` {
		v.Info.SourceURL = vault.Metadata.SourceURI
	}

	v.Info.RiskLevel = vault.Metadata.RiskLevel
	v.Info.RiskScore = [11]int8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	if (vault.Metadata.RiskScore != models.TRiskScore{}) {
		v.Info.RiskScore[0] = vault.Metadata.RiskScore.Review
		v.Info.RiskScore[1] = vault.Metadata.RiskScore.Testing
		v.Info.RiskScore[2] = vault.Metadata.RiskScore.Complexity
		v.Info.RiskScore[3] = vault.Metadata.RiskScore.RiskExposure
		v.Info.RiskScore[4] = vault.Metadata.RiskScore.ProtocolIntegration
		v.Info.RiskScore[5] = vault.Metadata.RiskScore.CentralizationRisk
		v.Info.RiskScore[6] = vault.Metadata.RiskScore.ExternalProtocolAudit
		v.Info.RiskScore[7] = vault.Metadata.RiskScore.ExternalProtocolCentralisation
		v.Info.RiskScore[8] = vault.Metadata.RiskScore.ExternalProtocolTvl
		v.Info.RiskScore[9] = vault.Metadata.RiskScore.ExternalProtocolLongevity
		v.Info.RiskScore[10] = vault.Metadata.RiskScore.ExternalProtocolType
	}
	v.Info.UINotice = vault.Metadata.UINotice

	return v, nil
}

func toTExternalVaultMigration(migration models.TMigration) TExternalVaultMigration {
	return TExternalVaultMigration{
		Available: migration.Available,
		Address:   migration.Target.Hex(),
		Contract:  migration.Contract.Hex(),
	}
}
func toArrTMixedcaseAddress(addresses []common.Address) []string {
	arr := make([]string, len(addresses))
	for i, address := range addresses {
		arr[i] = address.Hex()
	}
	return arr
}
