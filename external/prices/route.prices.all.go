package prices

import (
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

// For testing purposes - allows tests to intercept calls to storage.ListPrices
var listPricesFunc = func(chainID uint64) (map[common.Address]models.TPrices, []models.TPrices) {
	return storage.ListPrices(chainID)
}

/**************************************************************************************************
** GetAllPrices retrieves price information for all tokens across all supported blockchain
** networks. This handler provides a complete view of available price data in the system.
**
** The function iterates through all supported chains and fetches price data for each token present
** in the storage system for that chain. It then formats this data into a nested map structure:
** - First level: Chain ID → token prices map
** - Second level: Token address (hex) → price value
**
** The endpoint supports two response formats based on the 'humanized' query parameter:
** - Raw prices as big integers with full precision (default)
** - Humanized prices as floating-point numbers for human readability (when humanized=true)
**
** @param c The Gin context containing the 'humanized' query parameter
** - humanized: Optional parameter to format prices for human readability
**************************************************************************************************/
func (y Controller) GetAllPrices(c *gin.Context) {
	humanized := helpers.StringToBool(helpers.SafeString(getQuery(c, "humanized"), "false"))
	rawPrices := make(map[uint64]map[string]*bigNumber.Int)
	humanizedPrices := make(map[uint64]map[string]*bigNumber.Float)

	chains := env.GetChains()
	for _, chain := range chains {
		// Only process chains that we successfully can access
		if chain.ID == 0 {
			continue
		}

		allChainPrices, _ := listPricesFunc(chain.ID)
		// If there are no prices for this chain, skip it
		if len(allChainPrices) == 0 {
			continue
		}

		rawPrices[chain.ID] = make(map[string]*bigNumber.Int)
		humanizedPrices[chain.ID] = make(map[string]*bigNumber.Float)

		for addr, price := range allChainPrices {
			if humanized {
				humanizedPrices[chain.ID][addr.Hex()] = price.HumanizedPrice
			} else {
				rawPrices[chain.ID][addr.Hex()] = price.Price
			}
		}
	}

	// Return response based on humanized flag
	formatChainPriceMap(c, rawPrices, humanizedPrices, humanized)
}

/**************************************************************************************************
** GetPrices retrieves price information for all tokens on a specific blockchain network. This
** handler is useful when clients need prices for all tokens, but only on a particular chain.
**
** The function performs validation on the chain ID to ensure it's a valid positive integer. If
** valid, it fetches all prices stored for that chain and formats them in a map where:
** - Keys are token addresses in hex format
** - Values are price data, either as raw big integers or humanized floats
**
** The endpoint supports two response formats based on the 'humanized' query parameter:
** - Raw prices as big integers with full precision (default)
** - Humanized prices as floating-point numbers for human readability (when humanized=true)
**
** @param c The Gin context containing request parameters
** - chainID: Path parameter specifying the blockchain network ID
** - humanized: Optional query parameter to format prices for human readability
**************************************************************************************************/
func (y Controller) GetPrices(c *gin.Context) {
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}

	// Retrieve all prices for the chain
	allPrices, _ := listPricesFunc(chainID)

	// Process the humanized flag
	humanized := helpers.StringToBool(helpers.SafeString(getQuery(c, "humanized"), "false"))

	// Initialize the appropriate container based on humanized flag
	rawPrices := make(map[string]*bigNumber.Int)
	humanizedPrices := make(map[string]*bigNumber.Float)

	// Fill the container with prices
	for addr, price := range allPrices {
		if humanized {
			humanizedPrices[addr.Hex()] = price.HumanizedPrice
		} else {
			rawPrices[addr.Hex()] = price.Price
		}
	}

	// Return the appropriate response format
	formatPriceMap(c, rawPrices, humanizedPrices, humanized)
}
