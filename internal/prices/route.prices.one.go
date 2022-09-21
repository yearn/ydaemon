package prices

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/internal/utils/helpers"
)

//GetPrice will, for a given token on a given chainID, return the price of it
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

	humanized := helpers.ValueWithFallback(c.Query("humanized"), "false")
	price, ok := Store.TokenPrices[chainID][address]
	if !ok {
		c.String(http.StatusBadRequest, "invalid address")
		return
	}

	if humanized == "true" {
		humanizedPrice, _ := helpers.FormatAmount(price.String(), 6)
		c.JSON(http.StatusOK, humanizedPrice)
	} else {
		c.JSON(http.StatusOK, price)
	}
}
