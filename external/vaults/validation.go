package vaults

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

/************************************************************************************************
** StrategyCondition constants define the available filtering options for strategies.
** These determine which strategies will be included in API responses.
**
** - ALL: Include all strategies associated with a vault
** - IN_QUEUE: Include only strategies in the vault's withdrawal queue
** - DEBT_LIMIT: Include only strategies with an allocated debt limit
** - DEBT_RATIO: Include only strategies with an allocated debt ratio
** - ABSOLUTE: Include only active strategies with debt > 0
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
** ValidateStrategyCondition validates the strategy condition parameter and returns the
** appropriate value to use.
**
** @param c *gin.Context - The Gin context containing the request
** @param paramName string - The name of the query parameter to validate
** @return string - The validated strategy condition or default "debtRatio"
************************************************************************************************/
func ValidateStrategyCondition(c *gin.Context, paramName string) string {
	conditionParam := getQuery(c, paramName)
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
** ValidateMigrableCondition validates the migrable condition parameter and returns the
** appropriate value to use.
**
** @param c *gin.Context - The Gin context containing the request
** @param paramName string - The name of the query parameter to validate
** @return string - The validated migrable condition or default "none"
************************************************************************************************/
func ValidateMigrableCondition(c *gin.Context, paramName string) string {
	conditionParam := getQuery(c, paramName)
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
			HandleError(c, fmt.Errorf("operation timed out while processing strategies"),
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

			strategyWithDetails = NewStrategy().AssignTStrategy(strategy)
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
** ValidateAddressesParam validates a comma-separated list of addresses.
**
** @param c *gin.Context - The Gin context containing the request
** @param addressesParam string - The comma-separated list of addresses
** @param chainID uint64 - The chain ID to validate addresses against
** @param funcName string - The name of the calling function for error context
** @return []string - The list of validated addresses
** @return bool - True if validation succeeded, false if it failed
************************************************************************************************/
func ValidateAddressesParam(
	c *gin.Context,
	addressesParam string,
	chainID uint64,
	funcName string,
) ([]string, bool) {
	if addressesParam == "" {
		HandleError(c, fmt.Errorf("addresses parameter cannot be empty"),
			http.StatusBadRequest, "Missing required parameter", funcName)
		return nil, false
	}

	// Process addresses
	addressesStr := strings.Split(strings.ToLower(addressesParam), ",")
	if len(addressesStr) == 0 {
		HandleError(c, fmt.Errorf("at least one address must be provided"),
			http.StatusBadRequest, "Invalid parameter value", funcName)
		return nil, false
	}

	// Validate each address format
	for i, addr := range addressesStr {
		if !strings.HasPrefix(addr, "0x") || len(addr) != 42 {
			HandleError(c, fmt.Errorf("invalid address format at position %d: %s", i, addr),
				http.StatusBadRequest, "Invalid address format", funcName)
			return nil, false
		}
	}

	return addressesStr, true
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
** IsValidV3Vault checks if a vault is a valid v3 vault based on criteria.
**
** @param vault models.TVault - The vault to check
** @return bool - True if valid v3 vault, false otherwise
************************************************************************************************/
func IsValidV3Vault(vault models.TVault) bool {
	return vault.Kind == models.VaultKindMultiple ||
		vault.Kind == models.VaultKindSingle ||
		strings.HasPrefix(vault.Version, "3") ||
		vault.Version == "v3"
}

/************************************************************************************************
** IsValidV2Vault checks if a vault is a valid v2 vault based on criteria.
**
** @param vault models.TVault - The vault to check
** @return bool - True if valid v2 vault, false otherwise
************************************************************************************************/
func IsValidV2Vault(vault models.TVault) bool {
	// First check if it's a v3 vault (which takes precedence)
	if IsValidV3Vault(vault) {
		return false
	}

	// If it's a v1 vault, it's not v2
	if strings.HasPrefix(vault.Version, "1.") || vault.Version == "v1" || vault.Version == "1" {
		return false
	}

	// It's a v2 vault if it starts with "0." or "2." or equals "v2" or "2"
	return strings.HasPrefix(vault.Version, "0.") ||
		strings.HasPrefix(vault.Version, "2.") ||
		vault.Version == "v2" ||
		vault.Version == "2"
}
