package tokens

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

/**************************************************************************************************
** Controller is the main handler for token-related API endpoints. It provides access to token
** data across all supported chains or on specific chains.
**
** This struct follows the standard API handler pattern in the yDaemon codebase, where controller
** methods are registered as HTTP handlers in the router configuration.
**************************************************************************************************/
type Controller struct{}

/**************************************************************************************************
** TAllTokens represents a token with all its associated metadata for API responses.
** This structure contains both basic ERC20 token information (address, name, symbol, decimals)
** and extended metadata that enhances the token representation (display name, category, etc.).
**
** @field Address The Ethereum address of the token in hexadecimal format
** @field Name The official name of the token as defined in its contract
** @field Symbol The official symbol of the token as defined in its contract
** @field Decimals The number of decimal places the token uses
** @field IsVault Whether this token is a vault-like token with additional functionality
** @field DisplayName A human-friendly name that may differ from the contract name
** @field DisplaySymbol A human-friendly symbol that may differ from the contract symbol
** @field Description A brief explanation of the token's purpose or characteristics
** @field Category The classification of the token (e.g., "stablecoin", "governance", etc.)
** @field UnderlyingTokens A list of token addresses that this token wraps or represents
**************************************************************************************************/
type TAllTokens struct {
	Address          string   `json:"address"`
	Name             string   `json:"name"`
	Symbol           string   `json:"symbol"`
	Decimals         uint64   `json:"decimals"`
	IsVault          bool     `json:"isVault"`
	DisplayName      string   `json:"display_name,omitempty"`
	DisplaySymbol    string   `json:"display_symbol,omitempty"`
	Description      string   `json:"description,omitempty"`
	Category         string   `json:"category,omitempty"`
	UnderlyingTokens []string `json:"underlyingTokens,omitempty"`
}

/**************************************************************************************************
** convertToTokenResponse transforms a storage token model into the API response format.
** This helper function encapsulates the conversion logic to reduce code duplication across
** different API handlers.
**
** @param token The internal token model from storage
** @return TAllTokens The formatted token data for API responses
**************************************************************************************************/
func convertToTokenResponse(token models.TERC20Token) TAllTokens {
	// Use DisplaySymbol if available, otherwise fall back to Symbol
	displaySymbol := token.DisplaySymbol
	if displaySymbol == "" {
		displaySymbol = token.Symbol
	}

	tokenDetails := TAllTokens{
		Address:       token.Address.Hex(),
		Name:          token.Name,
		DisplayName:   token.DisplayName,
		Symbol:        token.Symbol,
		Decimals:      token.Decimals,
		DisplaySymbol: displaySymbol,
		Category:      token.Category,
		Description:   token.Description,
		IsVault:       token.IsVaultLike(),
	}

	// Add underlying tokens if present
	tokenDetails.UnderlyingTokens = make([]string, 0, len(token.UnderlyingTokensAddresses))
	for _, addr := range token.UnderlyingTokensAddresses {
		tokenDetails.UnderlyingTokens = append(tokenDetails.UnderlyingTokens, addr.Hex())
	}

	return tokenDetails
}

/**************************************************************************************************
** GetAllTokens retrieves information about all tokens across all supported blockchain networks.
** This comprehensive endpoint provides a complete view of token data in the system, organized by
** chain ID and token address.
**
** The function iterates through all supported chains and populates a nested map structure:
** - First level: Chain ID → token map
** - Second level: Token address (hex) → token details
**
** Each token entry includes basic ERC20 information (name, symbol, decimals) as well as extended
** metadata such as display names, categories, and underlying token relationships.
**
** @param c The Gin context for handling the HTTP request and response
** @return A JSON response with a map of chain IDs to token maps
**************************************************************************************************/
func (y Controller) GetAllTokens(c *gin.Context) {
	allTokens := make(map[uint64]map[string]TAllTokens)

	// Iterate through all supported chains
	for chainID := range env.GetChains() {
		tokens, _ := storage.ListERC20(chainID)

		// Skip if no tokens found for this chain
		if len(tokens) == 0 {
			continue
		}

		// Initialize the map for this chain
		allTokens[chainID] = make(map[string]TAllTokens, len(tokens))

		// Convert each token to the response format
		for _, token := range tokens {
			tokenDetails := convertToTokenResponse(token)
			allTokens[chainID][token.Address.Hex()] = tokenDetails
		}
	}

	c.JSON(http.StatusOK, allTokens)
}

/**************************************************************************************************
** GetTokens retrieves information about all tokens on a specific blockchain network.
** This endpoint is useful when clients need token data for a particular chain rather than
** across all supported networks.
**
** The function performs validation on:
** - Chain ID (must be a valid positive integer)
**
** If the chain ID is valid, it fetches all token data for that chain and formats it in a map where:
** - Keys are token addresses in hex format
** - Values are token details objects with metadata
**
** If an invalid chain ID is provided, the function returns a 400 Bad Request response.
**
** @param c The Gin context containing request parameters
** - chainID: Path parameter specifying the blockchain network ID
** @return A JSON response with a map of token addresses to token details, or an error message
**************************************************************************************************/
func (y Controller) GetTokens(c *gin.Context) {
	// Validate chain ID
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}

	// Fetch tokens for the specified chain
	tokenMap, _ := storage.ListERC20(chainID)

	// If no tokens found, return an empty map rather than null
	allTokens := make(map[string]TAllTokens, len(tokenMap))

	// Convert each token to the response format
	for _, token := range tokenMap {
		tokenDetails := convertToTokenResponse(token)
		allTokens[token.Address.Hex()] = tokenDetails
	}

	c.JSON(http.StatusOK, allTokens)
}

/**************************************************************************************************
** GetToken retrieves information about a specific token on a specific blockchain network.
** This endpoint allows direct access to detailed information about a single token.
**
** The function performs validation on:
** - Chain ID (must be a valid positive integer)
** - Token address (must be a valid Ethereum address for the specified chain)
**
** If both validators pass, the function attempts to fetch the token data. If found, it returns
** a detailed token profile with metadata; otherwise, it returns a 404 Not Found response.
**
** @param c The Gin context containing request parameters
** - chainID: Path parameter specifying the blockchain network ID
** - address: Path parameter specifying the token address
** @return A JSON response with token details or an appropriate error message
**************************************************************************************************/
func (y Controller) GetToken(c *gin.Context) {
	// Validate chain ID
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}

	// Validate token address
	address, ok := helpers.AssertAddress(c.Param("address"), chainID)
	if !ok {
		c.String(http.StatusBadRequest, "invalid address")
		return
	}

	// Get token data from storage
	token, ok := storage.GetERC20(chainID, address)
	if !ok {
		c.String(http.StatusNotFound, "token not found")
		return
	}

	// Convert to response format and return
	tokenDetails := convertToTokenResponse(token)
	c.JSON(http.StatusOK, tokenDetails)
}
