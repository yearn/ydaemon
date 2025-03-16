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
	// Additional fields used in route.harvests.go
	BlockNumber  string  `json:"blockNumber,omitempty"`
	HistAPR      float64 `json:"histAPR,omitempty"`
	VaultTVL     float64 `json:"vaultTVL,omitempty"`
	VaultName    string  `json:"vaultName,omitempty"`
	StrategyName string  `json:"strategyName,omitempty"`
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
** categorized, and processed in the UI and API responses. It includes:
**
** - Status flags: isRetired (no longer active), isHidden (not shown by default in UIs)
** - Classification: isPool (represents a liquidity pool), isAutomated (uses automation systems)
** - Display preferences: isHighlighted (promoted in UI listings)
** - Architecture flags: isAggregator (manages multiple strategies as a single vault)
** - Performance traits: isBoosted (has enhanced yield through external incentives)
** - Categorization: stability type, category, poolProvider, etc.
**
** API consumers can use these fields to filter, sort, and display vaults appropriately
** in user interfaces, dashboards, and analytics tools.
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
** TExternalVaultInfo contains risk assessment data and display metadata for vaults.
**
** This structure holds important information about a vault's risk profile, UI display guidelines,
** and special notices that help users make informed investment decisions.
**
** Risk Assessment System:
** - RiskLevel: Overall risk rating from 1 (safest) to 5 (highest risk)
** - RiskScore: Detailed 11-element array with individual risk factor scores:
**   [0] Code review quality (0-5)
**   [1] Testing coverage and quality (0-5)
**   [2] Code complexity assessment (0-5)
**   [3] Risk exposure to market conditions (0-5)
**   [4] Protocol integration complexity (0-5)
**   [5] Centralization risk of the strategy (0-5)
**   [6] External protocol audit status (0-5)
**   [7] External protocol centralization (0-5)
**   [8] External protocol TVL size (0-5)
**   [9] External protocol longevity (0-5)
**   [10] External protocol type risk (0-5)
**
** For Multi-Strategy Vaults, the overall RiskLevel is derived from the highest risk 
** strategy, while the detailed RiskScore array may not be populated (use the individual
** strategy risk scores instead).
**
** Display Guidelines:
** - UINotice: Optional message to display to users (warnings, information)
** - SourceURL: Link to relevant external resource (e.g., token purchase site)
** - Status flags: Control how the vault appears in listings
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
** TExternalVault represents a complete Yearn vault with all its associated data.
**
** This is the primary data structure used in API responses for detailed vault information.
** It provides a comprehensive view of a vault, including:
**
** Core Properties:
** - Identifiers: address, name, symbol, version
** - Classification: type, kind, category
** - Metadata: description, icon, display formats
**
** Financial Data:
** - TVL (Total Value Locked): assets under management and USD value
** - APR (Annual Percentage Rate): yield information across all sources
** - PricePerShare: token value growth over time
**
** Technical Components:
** - Underlying token details
** - Associated strategies (yield-generating components)
** - Staking options and rewards
** - Migration information (if a newer vault version exists)
**
** Risk and Presentation:
** - Risk assessment and scoring
** - UI display preferences and status flags
** - Featuring score (for sorting and highlighting)
**
** This structure is the most complete representation of a vault and should be used
** when detailed information is required. For list views or lightweight use cases,
** consider using TSimplifiedExternalVault instead.
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
** assignVaultAPR maps the internal TVaultAPY structure to the external TExternalVaultAPR structure.
**
** This function converts APY (Annual Percentage Yield) data from internal calculations to the 
** standardized APR (Annual Percentage Rate) format used in the API. The conversion preserves
** all component values needed for displaying comprehensive yield information to users.
**
** The field mapping is as follows:
** - Type: Yield calculation method (e.g., "crv", "compound")
** - NetAPR: Primary annualized rate value used for yield calculations and comparisons
** - Fees: All fee components (management, performance, etc.)
** - Points: Historical yield data points for trend analysis
** - PricePerShare: Token value growth data for verification
** - Extra: Additional yield sources (staking rewards, protocol rewards)
** - ForwardAPR: Projected future yield information
**
** @param vaultAPY apr.TVaultAPY - The internal APY structure to convert
** @return TExternalVaultAPR - The populated external APR structure for API responses
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

/************************************************************************************************
** CreateExternalVault transforms an internal vault model into the external API representation.
**
** This function handles the complete conversion process from the internal data model to the 
** external API format, including:
**
** 1. Retrieving associated token information (vault token, underlying token)
** 2. Building display names and symbols with proper formatting
** 3. Calculating and populating TVL (Total Value Locked) data
** 4. Fetching and formatting APR (Annual Percentage Rate) information
** 5. Loading risk assessment data and metadata
** 6. Setting up migration status and highlighting information
** 7. Populating staking-related information if available
**
** The function will return an error in the following cases:
** - If the vault token information cannot be found
** - If required data for conversion is missing or invalid
**
** @param vault models.TVault - The internal vault structure to convert
** @return TExternalVault - The fully populated external vault structure for API responses
** @return error - Error if conversion fails (typically due to missing token data)
************************************************************************************************/
func CreateExternalVault(vault models.TVault) (TExternalVault, error) {
	// Get vault token
	vaultToken, ok := storage.GetERC20(vault.ChainID, vault.Address)
	if !ok {
		return TExternalVault{}, errors.New(`token not found`)
	}

	// Get vault name and symbol
	name, displayName, formatedName := fetcher.BuildVaultNames(vault, vault.Metadata.DisplayName)
	symbol, displaySymbol, formatedSymbol := fetcher.BuildVaultSymbol(vault, vault.Metadata.DisplaySymbol)
	strategies, _ := storage.ListStrategiesForVault(vault.ChainID, vault.Address)

	// Create the vault directly without intermediate objects
	externalVault := TExternalVault{
		Address:           vault.Address.Hex(),
		Version:           vault.Version,
		Endorsed:          vault.Endorsed,
		Boosted:           vault.Metadata.IsBoosted,
		EmergencyShutdown: vault.EmergencyShutdown,
		ChainID:           vault.ChainID,
		TVL:               fetcher.BuildVaultTVL(vault),
		Migration:         toTExternalVaultMigration(vault.Metadata.Migration),
		Symbol:            symbol,
		DisplaySymbol:     displaySymbol,
		FormatedSymbol:    formatedSymbol,
		Name:              name,
		DisplayName:       displayName,
		FormatedName:      formatedName,
		Icon:              vaultToken.Icon,
		Type:              vaultToken.Type,
		Kind:              vault.Kind,
		Decimals:          vaultToken.Decimals,
		Description:       vault.Metadata.Description,
		Category:          fetcher.BuildVaultCategory(vault, strategies),
		PricePerShare:     vault.LastPricePerShare,
		Details: TExternalVaultDetails{
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
		},
		Info: TExternalVaultInfo{
			SourceURL: vault.Metadata.SourceURI,
			RiskLevel: vault.Metadata.RiskLevel,
			RiskScore: [11]int8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	// Set staking data
	externalVault.Staking = assignStakingData(vault.ChainID, vault.Address)

	// Set token data
	if underlyingToken, ok := storage.GetUnderlyingERC20(vault.ChainID, vault.Address); ok {
		externalVault.Token = TExternalERC20Token{
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
		externalVault.Token = TExternalERC20Token{
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

	// Initialize empty token array if needed
	if externalVault.Token.UnderlyingTokensAddresses == nil {
		externalVault.Token.UnderlyingTokensAddresses = []string{}
	}

	// Set APR data
	asyncAPR, ok := apr.GetComputedAPY(vault.ChainID, vault.Address)
	if ok {
		externalVault.APR = assignVaultAPR(asyncAPR.(apr.TVaultAPY))
	}

	// Set stability defaults
	if externalVault.Details.Stability == `` {
		externalVault.Details.Stability = models.VaultStabilityUnknown
	}
	if externalVault.Details.Category == `` {
		externalVault.Details.Category = models.VaultCategoryAutomatic
	}
	if len(vault.Metadata.Protocols) > 0 {
		externalVault.Details.PoolProvider = vault.Metadata.Protocols[0]
	}

	// Set source URL from gauge if available
	poolResult, found := storage.GetGauge(vault.ChainID, vault.AssetAddress)
	if found && poolResult.PoolURLs.Deposit != nil && len(poolResult.PoolURLs.Deposit) > 0 {
		externalVault.Info.SourceURL = poolResult.PoolURLs.Deposit[0]
	}

	// Set risk data from cached risk score if available
	cachedRiskScore, err := risks.GetCachedRiskScore(vault.ChainID, vault.Address)
	if err == nil {
		externalVault.Info.RiskLevel = cachedRiskScore.RiskLevel
		externalVault.Info.RiskScore = PopulateRiskScoreArray(cachedRiskScore.RiskScore)
		externalVault.Info.RiskScoreComment = cachedRiskScore.RiskScore.Comment
	} else if (vault.Metadata.RiskScore != models.TRiskScore{}) {
		externalVault.Info.RiskScore = PopulateRiskScoreArray(vault.Metadata.RiskScore)
		externalVault.Info.RiskScoreComment = vault.Metadata.RiskScore.Comment
	}

	externalVault.Info.UINotice = vault.Metadata.UINotice

	return externalVault, nil
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
