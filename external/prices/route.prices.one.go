package prices

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/storage"
)

/**************************************************************************************************
** GetOnePrice retrieves price information for a specific token on a specific blockchain network.
**
** The function performs validation on:
** - Chain ID (must be a valid positive integer)
** - Token address (must be a valid Ethereum address for the specified chain)
**
** The endpoint supports two response formats based on the 'humanized' query parameter:
** - Raw price as a big integer with full precision (default)
** - Humanized price as a floating-point number for human readability (when humanized=true)
**
** @param c The Gin context containing request parameters
** - chainID: Path parameter specifying the blockchain network ID
** - address: Path parameter specifying the token address
** - humanized: Optional query parameter to format the price for human readability
**************************************************************************************************/
func (y Controller) GetOnePrice(c *gin.Context) {
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

	// Get price data
	price, ok := storage.GetPrice(chainID, address)
	if !ok {
		c.String(http.StatusNotFound, "price not found")
		return
	}

	// Determine response format based on 'humanized' parameter
	humanized := helpers.StringToBool(helpers.SafeString(getQuery(c, "humanized"), "false"))

	// Use helper to format response
	formatSinglePrice(c, price.Price, price.HumanizedPrice, humanized)
}

/**************************************************************************************************
** GetPrice returns the price for a single token on a specific chain.
**
** This function handles HTTP GET requests to the /:chainID/prices/:address endpoint.
** It retrieves the price for the specified token address on the specified chain.
**
** @param c *gin.Context - The Gin context for the HTTP request
**************************************************************************************************/
func (c *Controller) GetPrice(ctx *gin.Context) {
	chainID, ok := helpers.AssertChainID(ctx.Param("chainID"))
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid chainID"})
		return
	}

	// Get the token address from the URL
	rawAddress := ctx.Param("address")
	if rawAddress == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "address is required"})
		return
	}

	// Convert the address to checksummed format
	tokenAddress := addresses.ToAddress(strings.ToLower(rawAddress))
	if tokenAddress.Hex() == "0x0000000000000000000000000000000000000000" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid token address"})
		return
	}

	// Check if we want humanized prices
	humanized := ctx.Query("humanized") == "true"

	// Fetch price from storage
	price, ok := storage.GetPrice(chainID, tokenAddress)
	if !ok {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "price not found"})
		return
	}

	formatSinglePrice(ctx, price.Price, price.HumanizedPrice, humanized)
}
