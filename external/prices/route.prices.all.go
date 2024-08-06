package prices

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/storage"
)

// GetAllPrices will return all the prices informations, no matter the chainID.
func (y Controller) GetAllPrices(c *gin.Context) {
	humanized := helpers.StringToBool(helpers.SafeString(getQuery(c, "humanized"), "false"))

	allPrices := make(map[uint64]map[string]*bigNumber.Int)
	allPricesHumanized := make(map[uint64]map[string]*bigNumber.Float)
	for chainID := range env.GetChains() {
		allChainPrices, _ := storage.ListPrices(chainID)
		for key, price := range allChainPrices {
			if _, ok := allPrices[chainID]; !ok {
				allPrices[chainID] = make(map[string]*bigNumber.Int)
				allPricesHumanized[chainID] = make(map[string]*bigNumber.Float)
			}

			if humanized {
				allPricesHumanized[chainID][key.Hex()] = price.HumanizedPrice
			} else {
				allPrices[chainID][key.Hex()] = price.Price
			}
		}
	}
	if humanized {
		c.JSON(http.StatusOK, allPricesHumanized)
	} else {
		c.JSON(http.StatusOK, allPrices)
	}
}

// GetPrices will, for a given chainID, return a tokens list with the associated prices
func (y Controller) GetPrices(c *gin.Context) {
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}
	humanized := helpers.StringToBool(helpers.SafeString(getQuery(c, "humanized"), "false"))

	allPrices, _ := storage.ListPrices(chainID)
	if humanized {
		humanizedPrices := make(map[string]*bigNumber.Float)
		for key, price := range allPrices {
			humanizedPrices[key.Hex()] = price.HumanizedPrice
		}
		c.JSON(http.StatusOK, humanizedPrices)
	} else {
		prices := make(map[string]*bigNumber.Int)
		for key, price := range allPrices {
			prices[key.Hex()] = price.Price
		}
		c.JSON(http.StatusOK, prices)
	}
}
