package prices

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/types/common"
)

// GetAllPrices will return all the prices informations, no matter the chainID.
func (y Controller) GetAllPrices(c *gin.Context) {
	humanized := helpers.StringToBool(helpers.SafeString(c.Query("humanized"), "false"))

	allPrices := make(map[uint64]map[common.Address]*bigNumber.Int)
	allPricesHumanized := make(map[uint64]map[common.Address]float64)
	for _, chainID := range env.SUPPORTED_CHAIN_IDS {
		prices := Store.TokenPrices[chainID]
		for key, price := range prices {
			if _, ok := allPrices[chainID]; !ok {
				allPrices[chainID] = make(map[common.Address]*bigNumber.Int)
				allPricesHumanized[chainID] = make(map[common.Address]float64)
			}

			if humanized {
				humanizedPrice, _ := helpers.FormatAmount(price.String(), 6)
				allPricesHumanized[chainID][key] = humanizedPrice
			} else {
				allPrices[chainID][key] = price
			}
		}
	}
	if humanized {
		c.JSON(http.StatusOK, allPricesHumanized)
	} else {
		c.JSON(http.StatusOK, allPrices)
	}
}

//GetPrices will, for a given chainID, return a tokens list with the associated prices
func (y Controller) GetPrices(c *gin.Context) {
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}

	humanized := helpers.StringToBool(helpers.SafeString(c.Query("humanized"), "false"))
	if humanized {
		allPricesHumanized := make(map[common.Address]float64)
		prices := Store.TokenPrices[chainID]
		for key, price := range prices {
			humanizedPrice, _ := helpers.FormatAmount(price.String(), 6)
			allPricesHumanized[key] = humanizedPrice
		}
		c.JSON(http.StatusOK, allPricesHumanized)
	} else {
		c.JSON(http.StatusOK, Store.TokenPrices[chainID])
	}
}
