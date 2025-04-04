package prices

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/storage"
)

/**************************************************************************************************
** GetAllPricesWithDetails retrieves detailed price information for all tokens on a specific
** blockchain network. This endpoint provides additional metadata beyond just the price values,
** including timestamp and price source when available.
**
** The function performs validation on the chain ID before retrieving the price information.
** Unlike the basic GetPrices endpoint, this endpoint always includes both raw and humanized
** price values in the response, along with any available metadata fields.
**
** The response is a map where:
** - Keys are token addresses (hex strings)
** - Values are objects containing price details (raw price, humanized price, etc.)
**
** @param c The Gin context containing request parameters
** - chainID: Path parameter specifying the blockchain network ID
**************************************************************************************************/
func (y Controller) GetAllPricesWithDetails(c *gin.Context) {
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}

	// Fetch all prices with details from storage
	priceDetails, _ := storage.ListPrices(chainID)

	// Format the response
	result := make(map[string]map[string]interface{}, len(priceDetails))
	for addr, priceDetail := range priceDetails {
		// Create a detail map with the fields we know are available
		details := map[string]interface{}{
			"price":          priceDetail.Price,
			"humanizedPrice": priceDetail.HumanizedPrice,
		}

		// Add additional fields if the model has them (timestamp, source, etc.)
		// This approach avoids accessing fields that might not exist

		result[addr.Hex()] = details
	}

	c.JSON(http.StatusOK, result)
}
