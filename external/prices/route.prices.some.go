package prices

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/prices"
)

// GetSomePrices will, for a list of tokens on a given chainID, return the price for them
func (y Controller) GetSomePrices(c *gin.Context) {
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}
	humanized := helpers.StringToBool(helpers.SafeString(c.Query("humanized"), "false"))
	rawPrices := make(map[string]*bigNumber.Int)
	humanizedPrices := make(map[string]float64)
	addressesStr := strings.Split(c.Param("addresses"), ",")
	for _, addressStr := range addressesStr {
		address, ok := helpers.AssertAddress(addressStr, chainID)
		if !ok {
			continue
		}

		humanizedPrices[address.Hex()] = 0
		rawPrices[address.Hex()] = bigNumber.NewInt()
		price, ok := prices.FindPrice(chainID, address)
		if !ok {
			continue
		}
		humanizedPrice, _ := helpers.FormatAmount(price.String(), 6)
		humanizedPrices[address.Hex()] = humanizedPrice
		rawPrices[address.Hex()] = price
	}
	if humanized {
		c.JSON(http.StatusOK, humanizedPrices)
	} else {
		c.JSON(http.StatusOK, rawPrices)
	}
}
