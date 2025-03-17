package prices

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/storage"
)

/**************************************************************************************************
** GetSomePricesForChain retrieves prices for a specific list of tokens on a single blockchain
** network. This endpoint is optimized for cases where clients need price data for a predefined
** subset of tokens rather than all tokens on a chain.
**
** The function performs validation on:
** - Chain ID (must be a valid positive integer)
** - Token addresses (each must be a valid Ethereum address for the specified chain)
**
** Invalid addresses are skipped, and the function returns data only for valid addresses.
** This endpoint supports two response formats based on the 'humanized' query parameter:
** - Raw prices as big integers with full precision (default)
** - Humanized prices as floating-point numbers for human readability (when humanized=true)
**
** @param c The Gin context containing request parameters
** - chainID: Path parameter specifying the blockchain network ID
** - addresses: Path parameter containing comma-separated token addresses
** - humanized: Optional query parameter to format prices for human readability
**************************************************************************************************/
func (y Controller) GetSomePricesForChain(c *gin.Context) {
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}
	humanized := helpers.StringToBool(helpers.SafeString(getQuery(c, "humanized"), "false"))
	rawPrices := make(map[string]*bigNumber.Int)
	humanizedPrices := make(map[string]*bigNumber.Float)
	addressesStr := strings.Split(c.Param("addresses"), ",")
	for _, addressStr := range addressesStr {
		address, ok := helpers.AssertAddress(addressStr, chainID)
		if !ok {
			continue
		}

		humanizedPrices[address.Hex()] = bigNumber.NewFloat()
		rawPrices[address.Hex()] = bigNumber.NewInt()
		price, ok := storage.GetPrice(chainID, address)
		if !ok {
			continue
		}
		humanizedPrices[address.Hex()] = price.HumanizedPrice
		rawPrices[address.Hex()] = price.Price
	}
	if humanized {
		c.JSON(http.StatusOK, humanizedPrices)
	} else {
		c.JSON(http.StatusOK, rawPrices)
	}
}

/**************************************************************************************************
** GetSomePrices retrieves price information for a specific list of token addresses on a
** particular blockchain network. This allows clients to fetch prices for exactly the tokens
** they're interested in.
**
** The function performs validation on:
** - Chain ID (must be a valid positive integer)
** - Address list (must contain valid Ethereum addresses for the specified chain)
**
** The endpoint supports two response formats based on the 'humanized' query parameter:
** - Raw prices as big integers with full precision (default)
** - Humanized prices as floating-point numbers for human readability (when humanized=true)
**
** @param c The Gin context containing request parameters
** - chainID: Path parameter specifying the blockchain network ID
** - addressList: Path parameter with comma-separated list of token addresses
** - humanized: Optional query parameter to format prices for human readability
**************************************************************************************************/
func (y Controller) GetSomePrices(c *gin.Context) {
	// Validate chain ID
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}

	// Get and validate address list
	addressesStr := c.Param("addressList")
	if addressesStr == "" {
		c.String(http.StatusNotFound, "not found")
		return
	}

	// Parse the humanized flag
	humanized := helpers.StringToBool(helpers.SafeString(getQuery(c, "humanized"), "false"))

	// Validate addresses
	addressList := splitAndTrim(addressesStr, ",")
	validAddresses, invalidAddresses := validateAndParseAddressList(addressList, chainID)

	if len(validAddresses) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":            "No valid addresses provided",
			"invalidAddresses": invalidAddresses,
		})
		return
	}

	// Initialize results only for the format we need
	var rawPrices map[string]*bigNumber.Int
	var humanizedPrices map[string]*bigNumber.Float

	if humanized {
		humanizedPrices = make(map[string]*bigNumber.Float, len(validAddresses))
	} else {
		rawPrices = make(map[string]*bigNumber.Int, len(validAddresses))
	}

	// Fetch prices for each valid address
	for _, addr := range validAddresses {
		price, ok := storage.GetPrice(chainID, addr)
		if !ok {
			continue // Skip addresses with errors
		}

		addrStr := addr.Hex()
		if humanized {
			humanizedPrices[addrStr] = price.HumanizedPrice
		} else {
			rawPrices[addrStr] = price.Price
		}
	}

	// Use helper to format the response
	formatPriceMap(c, rawPrices, humanizedPrices, humanized)
}

/**************************************************************************************************
** GetSomePostPrices retrieves prices for multiple tokens across different blockchain networks.
** This endpoint accepts addresses in the request body rather than in the URL path, which allows
** handling larger batches of tokens that might exceed URL length limitations.
**
** The function performs validation on:
** - JSON request body structure (must have an "addresses" field)
** - Chain IDs and token addresses format (each must be valid in "chainID:address" format)
**
** The function expects a JSON body with an "addresses" field containing a comma-separated list
** of token addresses with chain IDs in the format "chainID:address", for example:
** {"addresses": "1:0x6b175474e89094c44da98b954eedeac495271d0f,10:0xDA10009cBd5D07dd0CeCc66161FC93D7c9000da1"}
**
** The response is organized as a nested map:
** - First level: Chain ID → token prices map
** - Second level: Token address (hex) → price value
**
** This endpoint supports two response formats based on the 'humanized' query parameter:
** - Raw prices as big integers with full precision (default)
** - Humanized prices as floating-point numbers for human readability (when humanized=true)
**
** @param c The Gin context containing request parameters and body
** - Request body: JSON object with "addresses" field containing comma-separated "chainID:address" pairs
** - humanized: Optional query parameter to format prices for human readability
**************************************************************************************************/
func (y Controller) GetSomePostPrices(c *gin.Context) {
	type expectedBody struct {
		Addresses string `json:"addresses"`
	}

	humanized := helpers.StringToBool(helpers.SafeString(getQuery(c, "humanized"), "false"))
	rawPrices := make(map[uint64]map[string]*bigNumber.Int)
	humanizedPrices := make(map[uint64]map[string]*bigNumber.Float)
	var body expectedBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	addresses := body.Addresses
	addressesStr := strings.Split(addresses, ",")
	for _, addressStr := range addressesStr {
		splitted := strings.Split(addressStr, ":")
		if len(splitted) != 2 {
			continue
		}
		chainIDStr, addressStr := splitted[0], splitted[1]
		chainID, ok := helpers.AssertChainID(chainIDStr)
		if !ok {
			continue
		}
		address, ok := helpers.AssertAddress(addressStr, chainID)
		if !ok {
			continue
		}

		if _, ok := rawPrices[chainID]; !ok {
			rawPrices[chainID] = make(map[string]*bigNumber.Int)
			humanizedPrices[chainID] = make(map[string]*bigNumber.Float)
		}
		humanizedPrices[chainID][address.Hex()] = bigNumber.NewFloat()
		rawPrices[chainID][address.Hex()] = bigNumber.NewInt()
		price, ok := storage.GetPrice(chainID, address)
		if !ok {
			continue
		}
		humanizedPrices[chainID][address.Hex()] = price.HumanizedPrice
		rawPrices[chainID][address.Hex()] = price.Price
	}
	if humanized {
		c.JSON(http.StatusOK, humanizedPrices)
	} else {
		c.JSON(http.StatusOK, rawPrices)
	}
}
