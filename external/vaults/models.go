package vaults

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/internal/fetcher"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
	"github.com/yearn/ydaemon/processes/apr"
	"github.com/yearn/ydaemon/processes/risks"
)

/**************************************************************************************************
** TExternalVaultHarvest represents harvest data for a vault strategy.
**
** This structure contains information about a strategy's harvest transaction, including profit/loss
** amounts and their USD values. It's primarily used in reporting vault performance and for
** analytics dashboards.
**************************************************************************************************/
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

/**************************************************************************************************
** TExternalVaultMigration contains migration information for a vault.
**
** When vaults need to be migrated to newer versions, this structure provides data about the
** migration target and availability status, helping users transition their funds.
**************************************************************************************************/
type TExternalVaultMigration struct {
	Available bool   `json:"available"`
	Address   string `json:"address"`
	Contract  string `json:"contract"`
}

/**************************************************************************************************
** TExternalVaultDetails contains metadata and classification information for a vault.
**
** This structure holds various flags and attributes that determine how a vault is displayed,
** categorized, and processed in the UI and API responses. It includes metadata like retirement
** status, boosting, automation, and categorization information.
**************************************************************************************************/
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

/**************************************************************************************************
** TExternalERC20Token represents token information for vault assets and rewards.
**
** This structure contains the basic information about an ERC20 token, including its address,
** name, symbol, decimals, and price data. It's used for vault assets, reward tokens, and
** other token representations in the API.
**************************************************************************************************/
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

/**************************************************************************************************
** TExternalVaultInfo contains risk and display metadata for vaults.
**
** This structure holds important information about a vault's risk profile, UI display guidelines,
** and special notices. It includes risk scores, UI notices, and flags that affect how the vault
** is presented to users.
**
** The risk scoring system uses a combination of individual risk factors (stored in the RiskScore array)
** and an overall RiskLevel that ranges from 1 (most secure) to 5 (least secure).
**************************************************************************************************/
type TExternalVaultInfo struct {
	SourceURL        string   `json:"sourceURL,omitempty"` // The vault might require some specific tokens that needs to be bought by a specific provider. It's the URL of the provider.
	RiskLevel        int8     `json:"riskLevel"`           // The risk level of the vault. The value is a calculated from the sum of all risk score from the object for Single Strategy Vaults. Multi-Strategy Vault, highest `riskLevel` of all strategies is set. 1 is the most secure and 5 is the least secure.
	UINotice         string   `json:"uiNotice,omitempty"`  // The notice to display in the UI
	IsRetired        bool     `json:"isRetired"`
	IsBoosted        bool     `json:"isBoosted"`
	IsHighlighted    bool     `json:"isHighlighted"`
	RiskScore        [11]int8 `json:"riskScore"`                  // All risk scores of the Single Strategy Vault. Multi-Strategy Vault won't have this object because its risk score is combination of multiple vaults. For risk value use `riskLevel`. (empty for Multi-Strategy Vault). Array of 11 integers: [review, testing, complexity, riskExposure, protocolIntegration, centralizationRisk, externalProtocolAudit, externalProtocolCentralisation, externalProtocolTvl, externalProtocolLongevity, externalProtocolType]
	RiskScoreComment string   `json:"riskScoreComment,omitempty"` // Comment for the risk score to the strategy. Can be empty.
}

/**************************************************************************************************
** TExternalCompositeData contains additional yield data for composite vaults.
**
** This structure is used for vaults that incorporate multiple yield sources or have boosted
** rewards. It includes information about boost multipliers and underlying pool APY values that
** contribute to the overall vault yield.
**************************************************************************************************/
type TExternalCompositeData struct {
	Boost                 *bigNumber.Float `json:"boost"`
	PoolAPY               *bigNumber.Float `json:"poolAPY"`
	BoostedAPR            *bigNumber.Float `json:"boostedAPR"`
	BaseAPR               *bigNumber.Float `json:"baseAPR"`
	CvxAPR                *bigNumber.Float `json:"cvxAPR"`
	RewardsAPR            *bigNumber.Float `json:"rewardsAPR"`
	V3OracleCurrentAPR    *bigNumber.Float `json:"v3OracleCurrentAPR,omitempty"`
	V3OracleStratRatioAPR *bigNumber.Float `json:"v3OracleStratRatioAPR,omitempty"`
	KeepCRV               *bigNumber.Float `json:"keepCRV,omitempty"`
	KeepVelo              *bigNumber.Float `json:"keepVELO,omitempty"`
}

/**************************************************************************************************
** TExternalExtraRewards represents additional yield sources for vaults.
**
** This structure contains information about extra rewards that may be available beyond
** the standard yield, such as staking rewards and gamma (protocol-specific) rewards.
**************************************************************************************************/
type TExternalExtraRewards struct {
	StakingRewardsAPR *bigNumber.Float `json:"stakingRewardsAPR"`
	GammaRewardAPR    *bigNumber.Float `json:"gammaRewardAPR"`
}

/**************************************************************************************************
** TExternalForwardAPR contains projected yield information for vaults.
**
** This structure holds forward-looking APR data which represents the expected future
** performance of the vault. It includes a netAPR value and composite data showing the
** breakdown of yield sources.
**************************************************************************************************/
type TExternalForwardAPR struct {
	Type      string                 `json:"type"`
	NetAPR    *bigNumber.Float       `json:"netAPR"`
	Composite TExternalCompositeData `json:"composite"`
}

/**************************************************************************************************
** TExternalVaultAPR contains comprehensive yield information for a vault.
**
** This structure aggregates all yield-related data for a vault, including current APR,
** historical performance, fee information, price per share data, and forward-looking APR
** projections. It serves as the central source for all performance metrics.
**************************************************************************************************/
type TExternalVaultAPR struct {
	Type          string                `json:"type"`
	NetAPR        *bigNumber.Float      `json:"netAPR"`
	Fees          apr.TFees             `json:"fees"`
	Points        apr.THistoricalPoints `json:"points"`
	PricePerShare apr.TPricePerShare    `json:"pricePerShare"`
	Extra         TExternalExtraRewards `json:"extra"`
	ForwardAPR    TExternalForwardAPR   `json:"forwardAPR"`
}

/**************************************************************************************************
** TExternalVault represents a complete vault entity with all associated data.
**
** This is the primary data structure for representing a Yearn vault. It contains comprehensive
** information about the vault, including its fundamental properties, token details, TVL,
** strategies, migration options, APR data, and metadata. This structure is used for the
** main vault API endpoints.
**************************************************************************************************/
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
	APR               TExternalVaultAPR       `json:"apr"`
	Details           TExternalVaultDetails   `json:"details"`
	Strategies        []TStrategy             `json:"strategies"`
	Migration         TExternalVaultMigration `json:"migration"`
	Staking           TStakingData            `json:"staking"`
	Info              TExternalVaultInfo      `json:"info,omitempty"`
	FeaturingScore    float64                 `json:"featuringScore"` // Computing only
	PricePerShare     *bigNumber.Int          `json:"pricePerShare"`
}

/**************************************************************************************************
** TSimplifiedExternalVaultTVL represents total value locked data for simplified vault responses.
**
** This structure contains essential TVL information for a vault in a more compact format,
** including the total assets in raw form, the calculated TVL in USD, and the token price.
**************************************************************************************************/
type TSimplifiedExternalVaultTVL struct {
	TotalAssets *bigNumber.Int `json:"totalAssets"`
	TVL         float64        `json:"tvl"`
	Price       float64        `json:"price"`
}

/**************************************************************************************************
** TSimplifiedExternalERC20Token represents basic token information for simplified vault responses.
**
** This structure contains the essential information about ERC20 tokens used in the simplified
** vault responses, focusing on identification, display, and basic properties.
**************************************************************************************************/
type TSimplifiedExternalERC20Token struct {
	Address     string `json:"address"`
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	Description string `json:"description"`
	Decimals    uint64 `json:"decimals"`
}

/**************************************************************************************************
** TStakingRewardsData represents information about token rewards from staking.
**
** This structure contains detailed information about the rewards available from staking
** vault tokens, including reward token details, rate information, and APR data.
**************************************************************************************************/
type TStakingRewardsData struct {
	Address    string           `json:"address"`
	Name       string           `json:"name"`
	Symbol     string           `json:"symbol"`
	Decimals   uint64           `json:"decimals"`
	Price      float64          `json:"price"`
	IsFinished bool             `json:"isFinished"`
	FinishedAt uint64           `json:"finishedAt"`
	APR        *bigNumber.Float `json:"apr"`
	PerWeek    *bigNumber.Float `json:"perWeek"`
}

/**************************************************************************************************
** TStakingData represents the overall staking capabilities for a vault.
**
** This structure contains information about the staking options available for a vault,
** including the staking contract address, availability status, and detailed reward information.
**************************************************************************************************/
type TStakingData struct {
	Address   string                `json:"address"`
	Available bool                  `json:"available"`
	Source    string                `json:"source"`
	Rewards   []TStakingRewardsData `json:"rewards"`
}

/**************************************************************************************************
** TSimplifiedExternalVault provides a streamlined representation of vault data.
**
** This structure contains a more compact version of vault information suitable for list views
** and scenarios where the full vault details aren't necessary. It includes essential properties,
** token information, TVL, APR, strategies, and metadata.
**************************************************************************************************/
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
	APR            TExternalVaultAPR             `json:"apr"`
	Strategies     []TStrategy                   `json:"strategies"`
	Staking        TStakingData                  `json:"staking,omitempty"`
	Migration      TExternalVaultMigration       `json:"migration,omitempty"`
	FeaturingScore float64                       `json:"featuringScore"`
	PricePerShare  *bigNumber.Int                `json:"pricePerShare"`
	Info           TExternalVaultInfo            `json:"info,omitempty"`
}

/************************************************************************************************
** Function to map the internal TVaultAPY structure to the external TExternalVaultAPR structure.
** This function takes a TVaultAPY object as input and returns a TExternalVaultAPR object.
** The mapping includes:
** - Type: The type of APR.
** - NetAPR: The net annual percentage rate.
** - Fees: The associated fees.
** - Points: Historical points data.
** - Extra: Extra rewards data.
** - ForwardAPR: Forward-looking APR data, including type, net APR, and composite data.
************************************************************************************************/
func assignVaultAPR(vaultAPY apr.TVaultAPY) TExternalVaultAPR {
	return TExternalVaultAPR{
		Type:          vaultAPY.Type,
		NetAPR:        vaultAPY.NetAPY,
		Fees:          vaultAPY.Fees,
		Points:        vaultAPY.Points,
		PricePerShare: vaultAPY.PricePerShare,
		Extra: TExternalExtraRewards{
			StakingRewardsAPR: vaultAPY.Extra.StakingRewardsAPY,
			GammaRewardAPR:    vaultAPY.Extra.GammaRewardAPY,
		},
		ForwardAPR: TExternalForwardAPR{
			Type:   vaultAPY.ForwardAPY.Type,
			NetAPR: vaultAPY.ForwardAPY.NetAPY,
			Composite: TExternalCompositeData{
				Boost:                 vaultAPY.ForwardAPY.Composite.Boost,
				PoolAPY:               vaultAPY.ForwardAPY.Composite.PoolAPY,
				BoostedAPR:            vaultAPY.ForwardAPY.Composite.BoostedAPR,
				BaseAPR:               vaultAPY.ForwardAPY.Composite.BaseAPR,
				CvxAPR:                vaultAPY.ForwardAPY.Composite.CvxAPR,
				RewardsAPR:            vaultAPY.ForwardAPY.Composite.RewardsAPY,
				V3OracleCurrentAPR:    vaultAPY.ForwardAPY.Composite.V3OracleCurrentAPR,
				V3OracleStratRatioAPR: vaultAPY.ForwardAPY.Composite.V3OracleStratRatioAPR,
			},
		},
	}
}

/**************************************************************************************************
** NewVault creates a new empty TExternalVault instance.
**
** This constructor provides an initialized external vault structure that can be further
** populated with data from the internal models.
**
** @return TExternalVault - An initialized empty vault structure
**************************************************************************************************/
func NewVault() TExternalVault {
	return TExternalVault{}
}

/**************************************************************************************************
** AssignTVault populates an external vault structure with data from an internal vault model.
**
** This method converts an internal TVault structure to the external TExternalVault format,
** enriching it with additional data from related sources like token information, strategy data,
** risk scores, and APR calculations.
**
** @param vault models.TVault - The internal vault structure to convert
** @return TExternalVault - The populated external vault structure
** @return error - Any error encountered during the conversion process
**************************************************************************************************/
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

	asyncAPR, ok := apr.GetComputedAPY(vault.ChainID, vault.Address)
	if ok {
		v.APR = assignVaultAPR(asyncAPR.(apr.TVaultAPY))
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
	cachedRiskScore, err := risks.GetCachedRiskScore(vault.ChainID, vault.Address)
	if err == nil {
		v.Info.RiskLevel = cachedRiskScore.RiskLevel
		v.Info.RiskScore = [11]int8{
			cachedRiskScore.RiskScore.Review,
			cachedRiskScore.RiskScore.Testing,
			cachedRiskScore.RiskScore.Complexity,
			cachedRiskScore.RiskScore.RiskExposure,
			cachedRiskScore.RiskScore.ProtocolIntegration,
			cachedRiskScore.RiskScore.CentralizationRisk,
			cachedRiskScore.RiskScore.ExternalProtocolAudit,
			cachedRiskScore.RiskScore.ExternalProtocolCentralisation,
			cachedRiskScore.RiskScore.ExternalProtocolTvl,
			cachedRiskScore.RiskScore.ExternalProtocolLongevity,
			cachedRiskScore.RiskScore.ExternalProtocolType,
		}
		v.Info.RiskScoreComment = cachedRiskScore.RiskScore.Comment
	} else if (vault.Metadata.RiskScore != models.TRiskScore{}) {
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
		v.Info.RiskScoreComment = vault.Metadata.RiskScore.Comment
	}

	v.Info.UINotice = vault.Metadata.UINotice

	return v, nil
}

/**************************************************************************************************
** toTExternalVaultMigration converts internal migration data to the external format.
**
** This function transforms the internal TMigration structure to the external TExternalVaultMigration
** format, converting addresses to their hexadecimal string representation.
**
** @param migration models.TMigration - The internal migration data
** @return TExternalVaultMigration - The converted external migration structure
**************************************************************************************************/
func toTExternalVaultMigration(migration models.TMigration) TExternalVaultMigration {
	return TExternalVaultMigration{
		Available: migration.Available,
		Address:   migration.Target.Hex(),
		Contract:  migration.Contract.Hex(),
	}
}

/**************************************************************************************************
** toArrTMixedcaseAddress converts an array of common.Address to string addresses.
**
** This utility function transforms an array of Ethereum addresses from their binary representation
** to hexadecimal string format suitable for JSON serialization and API responses.
**
** @param addresses []common.Address - The array of addresses to convert
** @return []string - The array of hexadecimal address strings
**************************************************************************************************/
func toArrTMixedcaseAddress(addresses []common.Address) []string {
	arr := make([]string, len(addresses))
	for i, address := range addresses {
		arr[i] = address.Hex()
	}
	return arr
}
