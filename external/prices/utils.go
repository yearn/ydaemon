package prices

import (
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/helpers"
)

/**************************************************************************************************
** The prices package provides HTTP handlers for retrieving token price data across different
** blockchain networks. It acts as an external API layer for the yDaemon price data system,
** exposing various endpoints to query token prices in multiple formats and scopes.
**
** This package supports retrieving prices for individual tokens, batches of tokens, or all tokens
** across either specific chains or all supported networks. Responses can be formatted as either
** raw price data (with full precision) or humanized values for readability.
**************************************************************************************************/

/**************************************************************************************************
** Controller is the main handler struct for price-related routes. It follows the standard pattern
** of a method receiver for all route handlers in the package. This empty struct serves as a
** namespace for grouping related API endpoints and ensures consistent method signatures.
**************************************************************************************************/
type Controller struct{}

/**************************************************************************************************
** getQuery retrieves a query parameter from the Gin context in a case-insensitive manner.
** This helper function normalizes query parameter retrieval across all route handlers,
** providing consistent behavior for optional parameters like the 'humanized' flag.
**
** @param c The Gin context from which to extract the query parameter
** @param targetKey The name of the query parameter to retrieve
** @return string The value of the query parameter, or an empty string if not found
**************************************************************************************************/
func getQuery(c *gin.Context, targetKey string) string {
	queryParams := c.Request.URL.Query()
	for key, values := range queryParams {
		if strings.EqualFold(targetKey, key) {
			return strings.Join(values, ",")
		}
	}
	return ""
}

/**************************************************************************************************
** formatSinglePrice formats a price response for a single token based on the 'humanized' parameter.
** This function handles the common pattern of returning either a raw price or a humanized price
** based on the query parameter.
**
** @param c The Gin context containing the request parameters
** @param price The raw price value
** @param humanizedPrice The humanized price value
** @param humanizedOverride Optional explicit value to override the query parameter
**************************************************************************************************/
func formatSinglePrice(
	c *gin.Context,
	price *bigNumber.Int,
	humanizedPrice *bigNumber.Float,
	humanizedOverride ...bool,
) {
	humanized := false
	if len(humanizedOverride) > 0 {
		humanized = humanizedOverride[0]
	} else {
		humanized = helpers.StringToBool(helpers.SafeString(getQuery(c, "humanized"), "false"))
	}

	if humanized {
		c.JSON(http.StatusOK, humanizedPrice)
	} else {
		c.JSON(http.StatusOK, price)
	}
}

/**************************************************************************************************
** formatPriceMap formats a map of token addresses to prices based on the 'humanized' parameter.
** This function handles the common pattern across multiple endpoints of returning either raw or
** humanized prices for a collection of tokens.
**
** @param c The Gin context containing the request parameters
** @param rawPrices Map of addresses to raw price values
** @param humanizedPrices Map of addresses to humanized price values
** @param humanizedOverride Optional explicit value to override the query parameter
**************************************************************************************************/
func formatPriceMap(
	c *gin.Context,
	rawPrices map[string]*bigNumber.Int,
	humanizedPrices map[string]*bigNumber.Float,
	humanizedOverride ...bool,
) {
	humanized := false
	if len(humanizedOverride) > 0 {
		humanized = humanizedOverride[0]
	} else {
		humanized = helpers.StringToBool(helpers.SafeString(getQuery(c, "humanized"), "false"))
	}

	if humanized {
		c.JSON(http.StatusOK, humanizedPrices)
	} else {
		c.JSON(http.StatusOK, rawPrices)
	}
}

/**************************************************************************************************
** formatChainPriceMap formats a nested map of chain IDs to token prices based on the 'humanized'
** parameter. This function supports the multi-chain price endpoints.
**
** @param c The Gin context containing the request parameters
** @param rawPrices Nested map of chain IDs to addresses to raw price values
** @param humanizedPrices Nested map of chain IDs to addresses to humanized price values
** @param humanizedOverride Optional explicit value to override the query parameter
**************************************************************************************************/
func formatChainPriceMap(
	c *gin.Context,
	rawPrices map[uint64]map[string]*bigNumber.Int,
	humanizedPrices map[uint64]map[string]*bigNumber.Float,
	humanizedOverride ...bool,
) {
	humanized := false
	if len(humanizedOverride) > 0 {
		humanized = humanizedOverride[0]
	} else {
		humanized = helpers.StringToBool(helpers.SafeString(getQuery(c, "humanized"), "false"))
	}

	if humanized {
		c.JSON(http.StatusOK, humanizedPrices)
	} else {
		c.JSON(http.StatusOK, rawPrices)
	}
}

/**************************************************************************************************
** validateAndParseAddressList validates a list of addresses for a specific chain and returns a
** slice of valid addresses. This function centralizes the common pattern of validating multiple
** addresses and filtering out invalid ones.
**
** @param addressStrings The slice of address strings to validate
** @param chainID The chain ID to validate the addresses against
** @return A slice of valid Ethereum addresses
** @return A map of invalid addresses to error messages
**************************************************************************************************/
func validateAndParseAddressList(
	addressStrings []string,
	chainID uint64,
) ([]common.Address, map[string]string) {
	validAddresses := make([]common.Address, 0, len(addressStrings))
	invalidAddresses := make(map[string]string)

	for _, addressStr := range addressStrings {
		address, ok := helpers.AssertAddress(addressStr, chainID)
		if !ok {
			invalidAddresses[addressStr] = "invalid address format"
			continue
		}
		validAddresses = append(validAddresses, address)
	}

	return validAddresses, invalidAddresses
}

/**************************************************************************************************
** BatchLimitMiddleware creates a middleware that limits the number of items in batch requests.
** This function helps prevent abuse by limiting the maximum number of tokens that can be requested
** in a single call.
**
** @param maxItems The maximum number of items allowed in a batch request
** @return A Gin middleware function that enforces the limit
** INFO: UNUSED RIGHT NOW
**************************************************************************************************/
func BatchLimitMiddleware(maxItems int) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the addresses parameter and check the length
		addressesStr := c.Param("addresses")
		if addressesStr != "" {
			addresses := splitAndTrim(addressesStr, ",")
			if len(addresses) > maxItems {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Too many addresses requested",
					"limit": maxItems,
				})
				c.Abort()
				return
			}
		}

		// Continue to the next middleware/handler
		c.Next()
	}
}

/**************************************************************************************************
** splitAndTrim splits a string by a separator and trims whitespace from each element.
** This is a convenience function to help with parsing lists of items from request parameters.
**
** @param s The string to split
** @param sep The separator to split by
** @return A slice of strings with whitespace trimmed
**************************************************************************************************/
func splitAndTrim(s, sep string) []string {
	if s == "" {
		return []string{}
	}

	parts := strings.Split(s, sep)
	result := make([]string, 0, len(parts))

	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}

	return result
}
