package prices

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/types/common"
	"github.com/yearn/ydaemon/internal/utils/helpers"
	"github.com/yearn/ydaemon/internal/utils/types/bigNumber"
)

//GetSomePrices will, for a list of tokens on a given chainID, return the price for them
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
		addressStr := address.String()

		humanizedPrices[addressStr] = 0
		rawPrices[addressStr] = bigNumber.NewInt()
		price, ok := Store.TokenPrices[chainID][address]
		if !ok {
			continue
		}
		humanizedPrice, _ := helpers.FormatAmount(price.String(), 6)
		humanizedPrices[addressStr] = humanizedPrice
		rawPrices[addressStr] = price
	}
	if humanized {
		c.JSON(http.StatusOK, humanizedPrices)
	} else {
		c.JSON(http.StatusOK, rawPrices)
	}
}
