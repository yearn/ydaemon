package prices

import (
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/internal/utils/helpers"
)

// GetAllPrices will return all the prices informations, no matter the chainID.
func (y Controller) GetAllPrices(c *gin.Context) {
	humanized := helpers.StringToBool(helpers.SafeString(c.Query("humanized"), "false"))

	allPrices := make(map[uint64]map[common.Address]*big.Int)
	allPricesHumanized := make(map[uint64]map[common.Address]float64)
	for _, chainID := range helpers.SUPPORTED_CHAIN_IDS {
		prices := Store.TokenPrices[chainID]
		for key, price := range prices {
			if _, ok := allPrices[chainID]; !ok {
				allPrices[chainID] = make(map[common.Address]*big.Int)
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
