package prices

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/prices"
)

// GetPrice will, for a given token on a given chainID, return the price of it
func (y Controller) GetPrice(c *gin.Context) {
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}
	address, ok := helpers.AssertAddress(c.Param("address"), chainID)
	if !ok {
		c.String(http.StatusBadRequest, "invalid address")
		return
	}

	humanized := helpers.StringToBool(helpers.SafeString(c.Query("humanized"), "false"))
	price, ok := prices.FindPrice(chainID, address)
	if !ok {
		c.String(http.StatusNotFound, "token not found")
		return
	}

	if humanized {
		humanizedPrice, _ := helpers.FormatAmount(price.String(), 6)
		c.JSON(http.StatusOK, humanizedPrice)
	} else {
		c.JSON(http.StatusOK, price)
	}
}
