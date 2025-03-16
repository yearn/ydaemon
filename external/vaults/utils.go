package vaults

import (
	"fmt"
	"net/http"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
)

/************************************************************************************************
** Controller is the main struct for handling vault-related API endpoints.
** It provides methods for retrieving various types of vaults based on different filters.
************************************************************************************************/
type Controller struct{}

/************************************************************************************************
** getQuery retrieves a query parameter from a Gin context in a case-insensitive manner.
**
** @param c *gin.Context - The Gin context containing the request
** @param targetKey string - The name of the query parameter to retrieve
** @return string - The value of the query parameter or an empty string if not found
************************************************************************************************/
func getQuery(c *gin.Context, targetKey string) string {
	queryParams := c.Request.URL.Query()
	for key, values := range queryParams {
		if strings.EqualFold(targetKey, key) {
			return strings.Join(values, ",")
		}
	}
	return ""
}

/************************************************************************************************
** ValidateChainID validates a chainID parameter from the request context and sends appropriate
** error responses if validation fails.
**
** This utility function centralizes the common chain ID validation logic used across multiple
** endpoints. It performs the following validations:
** 1. Checks if the chainID parameter exists
** 2. Validates that the chainID is in the correct format
** 3. Verifies that the chainID is supported by the system
** 4. Confirms that a chain configuration exists for the chainID
**
** If any validation fails, it automatically sends an appropriate error response to the client.
**
** @param c *gin.Context - The Gin context containing the request
** @param paramName string - The name of the parameter to extract (defaults to "chainID" if empty)
** @return uint64 - The validated chain ID
** @return bool - True if validation succeeded, false if it failed (and response was sent)
************************************************************************************************/
func ValidateChainID(c *gin.Context, paramName string) (uint64, bool) {
	// Default to "chainID" if no parameter name is provided
	if paramName == "" {
		paramName = "chainID"
	}

	// Validate chainID parameter exists
	chainIDParam := c.Param(paramName)
	if chainIDParam == "" {
		c.String(http.StatusBadRequest, paramName+" parameter is required")
		return 0, false
	}

	// Validate chainID format and value
	chainID, ok := helpers.AssertChainID(chainIDParam)
	if !ok {
		c.String(http.StatusBadRequest, "invalid "+paramName+": "+chainIDParam)
		return 0, false
	}

	// Check if chain is supported
	if !helpers.Contains(env.SUPPORTED_CHAIN_IDS, chainID) {
		c.String(http.StatusBadRequest, fmt.Sprintf("chain %d is not supported", chainID))
		return 0, false
	}

	// Get chain configuration
	_, ok = env.GetChain(chainID)
	if !ok {
		c.String(http.StatusInternalServerError, fmt.Sprintf("chain configuration not found for %s %d", paramName, chainID))
		return 0, false
	}

	return chainID, true
}

/************************************************************************************************
** ValidateAddress validates an address parameter from the request context and sends appropriate
** error responses if validation fails.
**
** This utility function centralizes the common address validation logic used across multiple
** endpoints. It performs the following validations:
** 1. Checks if the address parameter exists
** 2. Validates that the address is in the correct format for the given chain
**
** If any validation fails, it automatically sends an appropriate error response to the client.
**
** @param c *gin.Context - The Gin context containing the request
** @param paramName string - The name of the parameter to extract (defaults to "address" if empty)
** @param chainID uint64 - The chain ID to validate the address against
** @return common.Address - The validated Ethereum address
** @return bool - True if validation succeeded, false if it failed (and response was sent)
************************************************************************************************/
func ValidateAddress(c *gin.Context, paramName string, chainID uint64) (common.Address, bool) {
	// Default to "address" if no parameter name is provided
	if paramName == "" {
		paramName = "address"
	}

	// Validate address parameter exists
	addressParam := c.Param(paramName)
	if addressParam == "" {
		c.String(http.StatusBadRequest, paramName+" parameter is required")
		return common.Address{}, false
	}

	// Validate address format
	address, ok := helpers.AssertAddress(addressParam, chainID)
	if !ok {
		c.String(http.StatusBadRequest, "invalid "+paramName+": "+addressParam)
		return common.Address{}, false
	}

	return address, true
}

/************************************************************************************************
** HandleError provides standardized error handling for API endpoints.
**
** This utility function ensures consistent error handling throughout the vaults package by:
** 1. Logging the error with appropriate context to Gin's error log
** 2. Formatting a user-friendly error message with relevant details
** 3. Setting the appropriate HTTP status code based on the error type
** 4. Returning a properly formatted JSON or string response to the client
**
** The function automatically adds the file and line number where the error occurred to aid in
** debugging. It also preserves any context that was added to the error using fmt.Errorf.
**
** @param c *gin.Context - The Gin context for the request
** @param err error - The error that occurred
** @param statusCode int - The HTTP status code to return (e.g., http.StatusBadRequest)
** @param userMessage string - A user-friendly message explaining what went wrong
** @param logContext string - Additional context for the log (e.g., function name or operation)
************************************************************************************************/
func HandleError(c *gin.Context, err error, statusCode int, userMessage string, logContext string) {
	// Skip this if there's no actual error
	if err == nil {
		return
	}

	// Get file and line number for better debugging
	_, file, line, ok := runtime.Caller(1)
	fileInfo := ""
	if ok {
		// Extract just the filename without the full path
		fileInfo = fmt.Sprintf(" [%s:%d]", filepath.Base(file), line)
	}

	// Format a detailed error message for the logs
	logMessage := fmt.Sprintf("%s: %v%s", logContext, err, fileInfo)

	// Log the error to Gin's internal error logger
	c.Error(fmt.Errorf("%s", logMessage))

	// If user message is empty, use a generic one
	if userMessage == "" {
		userMessage = "An error occurred while processing your request"
	}

	// Return appropriate response based on what the client would expect
	// If it's a JSON API endpoint, return JSON, otherwise return string
	if strings.Contains(c.GetHeader("Accept"), "application/json") {
		c.JSON(statusCode, gin.H{
			"error":   userMessage,
			"details": err.Error(),
		})
	} else {
		c.String(statusCode, fmt.Sprintf("%s: %v", userMessage, err.Error()))
	}
}

/************************************************************************************************
** ValidateNumericQuery validates a numeric query parameter against a specified range.
**
** This utility function handles the extraction and validation of numeric query parameters.
** It performs the following operations:
** 1. Gets the query parameter value using the getQuery function
** 2. Converts the string value to uint64 using the provided default if empty
** 3. Validates that the value falls within the specified min and max bounds
** 4. If the value is outside bounds, logs an error and returns the default value
**
** @param c *gin.Context - The Gin context containing the request
** @param paramName string - The name of the query parameter to validate
** @param defaultValue uint64 - The default value to use if parameter is missing or invalid
** @param minValue uint64 - The minimum allowed value for the parameter
** @param maxValue uint64 - The maximum allowed value for the parameter
** @param logContext string - Additional context for error logging
** @return uint64 - The validated numeric value
************************************************************************************************/
func ValidateNumericQuery(c *gin.Context, paramName string, defaultValue, minValue, maxValue uint64, logContext string) uint64 {
	// Get the query parameter
	paramValue := getQuery(c, paramName)
	if paramValue == "" {
		return defaultValue
	}

	// Convert to uint64
	value, err := strconv.ParseUint(paramValue, 10, 64)
	if err != nil {
		c.Error(fmt.Errorf("%s: invalid %s parameter: %s, using default value %d",
			logContext, paramName, paramValue, defaultValue))
		return defaultValue
	}

	// Validate range
	if value < minValue {
		c.Error(fmt.Errorf("%s: %s parameter %d is below minimum value %d, using minimum value",
			logContext, paramName, value, minValue))
		return minValue
	}
	if value > maxValue {
		c.Error(fmt.Errorf("%s: %s parameter %d exceeds maximum value %d, using maximum value",
			logContext, paramName, value, maxValue))
		return maxValue
	}

	return value
}

/************************************************************************************************
** ValidateStringChoiceQuery validates a string query parameter against a list of valid options.
**
** This utility function handles the extraction and validation of string query parameters where
** the value must be one of a predefined set of options.
** It performs the following operations:
** 1. Gets the query parameter value using the getQuery function
** 2. If the value is empty, returns the provided default
** 3. Validates that the value is in the list of valid options
** 4. If the value is invalid, logs an error and returns the default value
**
** @param c *gin.Context - The Gin context containing the request
** @param paramName string - The name of the query parameter to validate
** @param defaultValue string - The default value to use if parameter is missing or invalid
** @param validOptions []string - The list of valid options for this parameter
** @param logContext string - Additional context for error logging
** @return string - The validated string value
************************************************************************************************/
func ValidateStringChoiceQuery(c *gin.Context, paramName, defaultValue string, validOptions []string, logContext string) string {
	// Get the query parameter
	paramValue := getQuery(c, paramName)
	if paramValue == "" {
		return defaultValue
	}

	// Check if the value is in the list of valid options
	for _, option := range validOptions {
		if strings.EqualFold(paramValue, option) {
			return option // Return the properly cased version from the valid options
		}
	}

	// Value is not valid, log an error and return the default
	c.Error(fmt.Errorf("%s: invalid %s parameter: %s, valid options are %v, using default value %s",
		logContext, paramName, paramValue, validOptions, defaultValue))
	return defaultValue
}

/************************************************************************************************
** ValidateBooleanQuery validates a boolean query parameter.
**
** This utility function handles the extraction and validation of boolean query parameters.
** It performs the following operations:
** 1. Gets the query parameter value using the getQuery function
** 2. If the value is empty, returns the provided default
** 3. Validates that the value is a valid boolean string
** 4. If the value is invalid, logs an error and returns the default value
**
** @param c *gin.Context - The Gin context containing the request
** @param paramName string - The name of the query parameter to validate
** @param defaultValue bool - The default value to use if parameter is missing or invalid
** @param logContext string - Additional context for error logging
** @return bool - The validated boolean value
************************************************************************************************/
func ValidateBooleanQuery(c *gin.Context, paramName string, defaultValue bool, logContext string) bool {
	// Get the query parameter
	paramValue := getQuery(c, paramName)
	if paramValue == "" {
		return defaultValue
	}

	// Convert to boolean
	switch strings.ToLower(paramValue) {
	case "true", "1", "yes", "y":
		return true
	case "false", "0", "no", "n":
		return false
	default:
		c.Error(fmt.Errorf("%s: invalid %s parameter: %s, valid values are true/false, using default value %v",
			logContext, paramName, paramValue, defaultValue))
		return defaultValue
	}
}
