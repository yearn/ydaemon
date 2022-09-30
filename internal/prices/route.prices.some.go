package prices

import (
	"math/big"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/internal/types/common"
	"github.com/yearn/ydaemon/internal/utils/helpers"
)

//GetSomePrices will, for a list of tokens on a given chainID, return the price for them
func (y Controller) GetSomePrices(c *gin.Context) {
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}
	humanized := helpers.StringToBool(helpers.SafeString(c.Query("humanized"), "false"))
	rawPrices := make(map[common.Address]*big.Int)
	humanizedPrices := make(map[common.Address]float64)
	addressesStr := strings.Split(c.Param("addresses"), ",")
	for _, addressStr := range addressesStr {
		address, ok := helpers.AssertAddress(addressStr, chainID)
		if !ok {
			continue
		}

		humanizedPrices[address] = 0
		rawPrices[address] = big.NewInt(0)
		price, ok := Store.TokenPrices[chainID][address]
		if !ok {
			continue
		}
		humanizedPrice, _ := helpers.FormatAmount(price.String(), 6)
		humanizedPrices[address] = humanizedPrice
		rawPrices[address] = price
	}
	if humanized {
		c.JSON(http.StatusOK, humanizedPrices)
	} else {
		c.JSON(http.StatusOK, rawPrices)
	}
}
