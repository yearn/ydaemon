package vaults

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

/************************************************************************************************
** StrategyCondition constants define the available filtering options for strategies.
** These constants are used as query parameters to determine which strategies will be included
** in API responses.
**
** Available filter conditions:
** - ALL: Include all strategies associated with a vault, regardless of status
** - IN_QUEUE: Include only strategies that are in the vault's withdrawal queue
** - DEBT_LIMIT: Include only strategies that have a non-zero debt limit allocation
** - DEBT_RATIO: Include only strategies that have a non-zero debt ratio allocation
** - ABSOLUTE: Include only strategies with actual debt > 0 (currently deployed funds)
**
** The default filter used when no condition is specified is DEBT_RATIO, which returns
** strategies that are allocated a portion of the vault's assets.
************************************************************************************************/
const (
	STRATEGY_CONDITION_ALL        = "all"
	STRATEGY_CONDITION_IN_QUEUE   = "inQueue"
	STRATEGY_CONDITION_DEBT_LIMIT = "debtLimit"
	STRATEGY_CONDITION_DEBT_RATIO = "debtRatio"
	STRATEGY_CONDITION_ABSOLUTE   = "absolute"
)

/************************************************************************************************
** MigrableCondition constants define the available filtering options for migrable vaults.
** These determine which migrable vaults will be included in API responses.
**
** - NONE: Do not include migrable vaults if they are hideAlways or retired
** - ALL: Include all migrable vaults
** - NO_DUST: Include migrable vaults with TVL > $100
** - IGNORE: Completely ignore any vault with a migration available
************************************************************************************************/
const (
	MIGRABLE_CONDITION_NONE    = "none"
	MIGRABLE_CONDITION_ALL     = "all"
	MIGRABLE_CONDITION_NO_DUST = "nodust"
	MIGRABLE_CONDITION_IGNORE  = "ignore"
)

/************************************************************************************************
** Common constants used across the vaults package.
** These include timeouts, default values, and array sizes.
************************************************************************************************/
const (
	// Default timeout for endpoint requests
	DEFAULT_REQUEST_TIMEOUT = 10 * time.Second

	// Default pagination values
	DEFAULT_PAGE_NUMBER = 1
	DEFAULT_PAGE_LIMIT  = 200
	MAX_PAGE_LIMIT      = 1000

	// Common TVL thresholds
	MIN_DUST_TVL = 100 // Minimum TVL in USD to not be considered "dust"

	// Multiplier values
	HIGHLIGHTING_MULTIPLIER = 1e18 // Used to boost featuring score for highlighted vaults
)

/************************************************************************************************
** Risk score indices for the standardized array of risk scores.
** These constants provide semantically meaningful names for the risk score array indices.
************************************************************************************************/
const (
	RISK_SCORE_REVIEW                      = 0
	RISK_SCORE_TESTING                     = 1
	RISK_SCORE_COMPLEXITY                  = 2
	RISK_SCORE_RISK_EXPOSURE               = 3
	RISK_SCORE_PROTOCOL_INTEGRATION        = 4
	RISK_SCORE_CENTRALIZATION_RISK         = 5
	RISK_SCORE_EXTERNAL_PROTOCOL_AUDIT     = 6
	RISK_SCORE_EXTERNAL_PROTOCOL_CENTRAL   = 7
	RISK_SCORE_EXTERNAL_PROTOCOL_TVL       = 8
	RISK_SCORE_EXTERNAL_PROTOCOL_LONGEVITY = 9
	RISK_SCORE_EXTERNAL_PROTOCOL_TYPE      = 10
	RISK_SCORE_ARRAY_LENGTH                = 11
)

/************************************************************************************************
** validateStrategyCondition validates the strategy condition parameter and returns the
** appropriate value to use.
**
** @param c *gin.Context - The Gin context containing the request
** @param paramName string - The name of the query parameter to validate
** @return string - The validated strategy condition or default "debtRatio"
************************************************************************************************/
func validateStrategyCondition(c *gin.Context, paramName string) string {
	conditionParam := getQueryParam(c, paramName)
	if conditionParam == "" {
		return STRATEGY_CONDITION_DEBT_RATIO
	}

	validConditions := map[string]bool{
		STRATEGY_CONDITION_ALL:        true,
		STRATEGY_CONDITION_IN_QUEUE:   true,
		STRATEGY_CONDITION_DEBT_LIMIT: true,
		STRATEGY_CONDITION_DEBT_RATIO: true,
		STRATEGY_CONDITION_ABSOLUTE:   true,
	}

	if validConditions[conditionParam] {
		return conditionParam
	}

	// If invalid, log a warning and return the default
	c.Error(fmt.Errorf("invalid strategy condition: %s, using default", conditionParam))
	return STRATEGY_CONDITION_DEBT_RATIO
}

/************************************************************************************************
** validateMigrableCondition validates the migrable condition parameter and returns the
** appropriate value to use.
**
** @param c *gin.Context - The Gin context containing the request
** @param paramName string - The name of the query parameter to validate
** @return string - The validated migrable condition or default "none"
************************************************************************************************/
func validateMigrableCondition(c *gin.Context, paramName string) string {
	conditionParam := getQueryParam(c, paramName)
	if conditionParam == "" {
		return MIGRABLE_CONDITION_NONE
	}

	validConditions := map[string]bool{
		MIGRABLE_CONDITION_NONE:    true,
		MIGRABLE_CONDITION_ALL:     true,
		MIGRABLE_CONDITION_NO_DUST: true,
		MIGRABLE_CONDITION_IGNORE:  true,
	}

	if validConditions[conditionParam] {
		return conditionParam
	}

	// If invalid, log a warning and return the default
	c.Error(fmt.Errorf("invalid migrable condition: %s, using default", conditionParam))
	return MIGRABLE_CONDITION_NONE
}

/************************************************************************************************
** ProcessStrategiesForVault processes and filters strategies for a vault based on the
** specified condition.
**
** This function centralizes the common pattern of:
** 1. Retrieving strategies for a vault
** 2. Converting them to external format
** 3. Filtering based on condition
** 4. Adding timeout checks
**
** @param ctx context.Context - The request context for timeout management
** @param c *gin.Context - The Gin context for error handling
** @param chainID uint64 - The chain ID of the vault
** @param vaultAddress common.Address - The address of the vault
** @param condition string - The strategy filtering condition
** @param funcName string - The name of the calling function for error context
** @return []TStrategy - The processed list of strategies
** @return bool - True if successful, false if a timeout occurred
************************************************************************************************/
func ProcessStrategiesForVault(
	ctx context.Context,
	c *gin.Context,
	chainID uint64,
	vaultAddress common.Address,
	condition string,
	funcName string,
) ([]TStrategy, bool) {
	_, vaultStrategies := storage.ListStrategiesForVault(chainID, vaultAddress)

	// Pre-allocate the slice to avoid reallocations
	results := make([]TStrategy, 0, len(vaultStrategies))

	for _, strategy := range vaultStrategies {
		// Check for context timeout
		select {
		case <-ctx.Done():
			handleError(c, fmt.Errorf("operation timed out while processing strategies"),
				http.StatusGatewayTimeout, "Request processing timed out", funcName)
			return nil, false
		default:
			// Continue processing
		}

		// Use a deferred recovery function to prevent panics during strategy processing
		var strategyWithDetails TStrategy
		func() {
			defer func() {
				if r := recover(); r != nil {
					c.Error(fmt.Errorf("panic while processing strategy %s: %v",
						strategy.Address.String(), r))
					// Continue with next strategy
				}
			}()

			strategyWithDetails = CreateExternalStrategy(strategy)
		}()

		// Skip invalid strategies or if conversion failed
		if strategyWithDetails.Address == "" {
			c.Error(fmt.Errorf("failed to convert strategy %s to external format",
				strategy.Address.String()))
			continue
		}

		if !strategyWithDetails.ShouldBeIncluded(condition) {
			continue
		}

		results = append(results, strategyWithDetails)
	}

	return results, true
}

/************************************************************************************************
** validateAddressesParam validates a comma-separated list of addresses.
**
** @param c *gin.Context - The Gin context containing the request
** @param addressesParam string - The comma-separated list of addresses
** @param chainID uint64 - The chain ID to validate addresses against
** @param funcName string - The name of the calling function for error context
** @return []string - The list of validated addresses
** @return bool - True if validation succeeded, false if it failed
************************************************************************************************/
func validateAddressesParam(
	c *gin.Context,
	addressesParam string,
	chainID uint64,
	funcName string,
) ([]string, bool) {
	if addressesParam == "" {
		handleError(c, fmt.Errorf("addresses parameter cannot be empty"),
			http.StatusBadRequest, "Missing required parameter", funcName)
		return nil, false
	}

	// Pre-validate the string to avoid unnecessary splits
	if len(addressesParam) < 42 { // Minimum length for a single address
		handleError(c, fmt.Errorf("address parameter is too short to be valid"),
			http.StatusBadRequest, "Invalid parameter value", funcName)
		return nil, false
	}

	// Process addresses - split first without lowercasing to improve performance
	addressesStr := strings.Split(addressesParam, ",")
	if len(addressesStr) == 0 {
		handleError(c, fmt.Errorf("at least one address must be provided"),
			http.StatusBadRequest, "Invalid parameter value", funcName)
		return nil, false
	}

	// Pre-allocate result slice for better performance
	result := make([]string, 0, len(addressesStr))

	// Validate each address format
	for i, addr := range addressesStr {
		// Use direct byte comparison to check prefix (more efficient than HasPrefix)
		if len(addr) != 42 || addr[0] != '0' || addr[1] != 'x' {
			handleError(c, fmt.Errorf("invalid address format at position %d: %s", i, addr),
				http.StatusBadRequest, "Invalid address format", funcName)
			return nil, false
		}

		// Convert to lowercase and add to result
		result = append(result, strings.ToLower(addr))
	}

	return result, true
}

/************************************************************************************************
** IsVaultBlacklisted checks if a vault is blacklisted on the specified chain.
**
** @param chainID uint64 - The chain ID to check
** @param address common.Address - The vault address to check
** @return bool - True if the vault is blacklisted, false otherwise
************************************************************************************************/
func IsVaultBlacklisted(chainID uint64, address common.Address) bool {
	chain, ok := env.GetChain(chainID)
	if !ok {
		return false
	}
	return helpers.Contains(chain.BlacklistedVaults, address)
}

/************************************************************************************************
** VaultVersionChecks contains centralized logic for identifying Yearn vault versions.
**
** Yearn has multiple generations of vaults with different implementations:
** - V1: First generation vaults (legacy)
** - V2: Second generation vaults with strategy separation (primary version until 2023)
** - V3: Third generation vaults with enhanced security, functionality, and gas efficiency
**
** These helper functions ensure consistent version detection across the codebase, which is
** critical for determining available functionality, compatible strategies, and API behavior.
************************************************************************************************/

// VaultVersionChecks provides a unified interface for vault version classification.
// Using this structure ensures consistent version detection throughout the application.
var VaultVersionChecks = struct {
	// IsV3 checks if a vault is a v3 vault based on version or kind
	//
	// A vault is considered V3 if:
	// - It has kind VaultKindMultiple or VaultKindSingle
	// - Its version string starts with "3" or equals "v3"
	//
	// @param vault models.TVault - The vault to check
	// @return bool - True if the vault is a v3 vault, false otherwise
	IsV3 func(vault models.TVault) bool

	// IsV2 checks if a vault is a v2 vault based on version
	//
	// A vault is considered V2 if:
	// - It's NOT a V3 vault AND
	// - It's NOT a V1 vault AND
	// - Its version starts with "0." or "2." or equals "v2" or "2"
	//
	// @param vault models.TVault - The vault to check
	// @return bool - True if the vault is a v2 vault, false otherwise
	IsV2 func(vault models.TVault) bool

	// IsV1 checks if a vault is a v1 vault based on version
	//
	// A vault is considered V1 if:
	// - Its version string starts with "1." or equals "v1" or "1"
	//
	// @param vault models.TVault - The vault to check
	// @return bool - True if the vault is a v1 vault, false otherwise
	IsV1 func(vault models.TVault) bool
}{
	IsV3: func(vault models.TVault) bool {
		return vault.Kind == models.VaultKindMultiple ||
			vault.Kind == models.VaultKindSingle ||
			strings.HasPrefix(vault.Version, "3") ||
			vault.Version == "v3"
	},
	IsV1: func(vault models.TVault) bool {
		return strings.HasPrefix(vault.Version, "1.") ||
			vault.Version == "v1" ||
			vault.Version == "1"
	},
}

// Initialize IsV2 separately because it depends on IsV3 and IsV1
func init() {
	VaultVersionChecks.IsV2 = func(vault models.TVault) bool {
		// First check if it's a v3 vault (which takes precedence)
		if VaultVersionChecks.IsV3(vault) {
			return false
		}

		// If it's a v1 vault, it's not v2
		if VaultVersionChecks.IsV1(vault) {
			return false
		}

		// It's a v2 vault if it starts with "0." or "2." or equals "v2" or "2"
		return strings.HasPrefix(vault.Version, "0.") ||
			strings.HasPrefix(vault.Version, "2.") ||
			vault.Version == "v2" ||
			vault.Version == "2"
	}
}

/************************************************************************************************
** IsValidV3Vault checks if a vault is a valid v3 vault based on criteria.
**
** @param vault models.TVault - The vault to check
** @return bool - True if valid v3 vault, false otherwise
************************************************************************************************/
func IsValidV3Vault(vault models.TVault) bool {
	return VaultVersionChecks.IsV3(vault)
}

/************************************************************************************************
** IsValidV2Vault checks if a vault is a valid v2 vault based on criteria.
**
** @param vault models.TVault - The vault to check
** @return bool - True if valid v2 vault, false otherwise
************************************************************************************************/
func IsValidV2Vault(vault models.TVault) bool {
	return VaultVersionChecks.IsV2(vault)
}

/************************************************************************************************
** PopulateRiskScoreArray creates a standardized risk score array from a risk score struct.
**
** This function centralizes the logic for converting a TRiskScore struct to the fixed-size
** array used in external vault representations. Using the constant indices ensures consistency
** and makes the code more readable and maintainable.
**
** @param riskScore models.TRiskScore - The risk score struct to convert
** @return [RISK_SCORE_ARRAY_LENGTH]int8 - The populated risk score array
************************************************************************************************/
func PopulateRiskScoreArray(riskScore models.TRiskScore) [RISK_SCORE_ARRAY_LENGTH]int8 {
	var riskScoreArray [RISK_SCORE_ARRAY_LENGTH]int8

	riskScoreArray[RISK_SCORE_REVIEW] = riskScore.Review
	riskScoreArray[RISK_SCORE_TESTING] = riskScore.Testing
	riskScoreArray[RISK_SCORE_COMPLEXITY] = riskScore.Complexity
	riskScoreArray[RISK_SCORE_RISK_EXPOSURE] = riskScore.RiskExposure
	riskScoreArray[RISK_SCORE_PROTOCOL_INTEGRATION] = riskScore.ProtocolIntegration
	riskScoreArray[RISK_SCORE_CENTRALIZATION_RISK] = riskScore.CentralizationRisk
	riskScoreArray[RISK_SCORE_EXTERNAL_PROTOCOL_AUDIT] = riskScore.ExternalProtocolAudit
	riskScoreArray[RISK_SCORE_EXTERNAL_PROTOCOL_CENTRAL] = riskScore.ExternalProtocolCentralisation
	riskScoreArray[RISK_SCORE_EXTERNAL_PROTOCOL_TVL] = riskScore.ExternalProtocolTvl
	riskScoreArray[RISK_SCORE_EXTERNAL_PROTOCOL_LONGEVITY] = riskScore.ExternalProtocolLongevity
	riskScoreArray[RISK_SCORE_EXTERNAL_PROTOCOL_TYPE] = riskScore.ExternalProtocolType

	return riskScoreArray
}
