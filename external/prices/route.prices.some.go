package prices

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/storage"
)

// GetSomePrices will, for a list of tokens on a given chainID, return the price for them
func (y Controller) GetSomePrices(c *gin.Context) {
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
