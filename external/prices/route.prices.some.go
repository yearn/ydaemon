package prices

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/storage"
)

/**************************************************************************************************
** GetSomePricesForChain will, for a list of tokens, return the price for them. It's chain
** specific. The expected parameter for addresses is a comma-separated list of token addresses.
**************************************************************************************************/
func (y Controller) GetSomePricesForChain(c *gin.Context) {
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

/**************************************************************************************************
** GetSomePrices will, for a list of tokens, return the price for them. It's chain agnostic.
** The expected parameter for addresses is a comma-separated list of token addresses prepended with
** the chain ID. For example:
** "1:0x6b175474e89094c44da98b954eedeac495271d0f,56:0x6b175474e89094c44da98b954eedeac495271d0f".
**************************************************************************************************/
func (y Controller) GetSomePrices(c *gin.Context) {
	humanized := helpers.StringToBool(helpers.SafeString(getQuery(c, "humanized"), "false"))
	rawPrices := make(map[uint64]map[string]*bigNumber.Int)
	humanizedPrices := make(map[uint64]map[string]*bigNumber.Float)
	addressesStr := strings.Split(c.Param("addresses"), ",")
	for _, addressStr := range addressesStr {
		splitted := strings.Split(addressStr, ":")
		if len(splitted) != 2 {
			continue
		}
		chainIDStr, addressStr := splitted[0], splitted[1]
		chainID, ok := helpers.AssertChainID(chainIDStr)
		if !ok {
			continue
		}
		address, ok := helpers.AssertAddress(addressStr, chainID)
		if !ok {
			continue
		}

		if _, ok := rawPrices[chainID]; !ok {
			rawPrices[chainID] = make(map[string]*bigNumber.Int)
			humanizedPrices[chainID] = make(map[string]*bigNumber.Float)
		}
		humanizedPrices[chainID][address.Hex()] = bigNumber.NewFloat()
		rawPrices[chainID][address.Hex()] = bigNumber.NewInt()
		price, ok := storage.GetPrice(chainID, address)
		if !ok {
			continue
		}
		humanizedPrices[chainID][address.Hex()] = price.HumanizedPrice
		rawPrices[chainID][address.Hex()] = price.Price
	}
	if humanized {
		c.JSON(http.StatusOK, humanizedPrices)
	} else {
		c.JSON(http.StatusOK, rawPrices)
	}
}

/**************************************************************************************************
** GetSomePostPrices will, for a list of tokens, return the price for them. It's chain agnostic.
** While the GetSomePrices expects the addresses to be in the path, this function expects them to
** be in the body of the request. The expected parameter for addresses is a comma-separated list of
** token addresses prepended with the chain ID. For example:
** "1:0x6b175474e89094c44da98b954eedeac495271d0f,56:0x6b175474e89094c44da98b954eedeac495271d0f".
**************************************************************************************************/
func (y Controller) GetSomePostPrices(c *gin.Context) {
	type expectedBody struct {
		Addresses string `json:"addresses"`
	}

	humanized := helpers.StringToBool(helpers.SafeString(getQuery(c, "humanized"), "false"))
	rawPrices := make(map[uint64]map[string]*bigNumber.Int)
	humanizedPrices := make(map[uint64]map[string]*bigNumber.Float)
	var body expectedBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	addresses := body.Addresses
	addressesStr := strings.Split(addresses, ",")
	for _, addressStr := range addressesStr {
		splitted := strings.Split(addressStr, ":")
		if len(splitted) != 2 {
			continue
		}
		chainIDStr, addressStr := splitted[0], splitted[1]
		chainID, ok := helpers.AssertChainID(chainIDStr)
		if !ok {
			continue
		}
		address, ok := helpers.AssertAddress(addressStr, chainID)
		if !ok {
			continue
		}

		if _, ok := rawPrices[chainID]; !ok {
			rawPrices[chainID] = make(map[string]*bigNumber.Int)
			humanizedPrices[chainID] = make(map[string]*bigNumber.Float)
		}
		humanizedPrices[chainID][address.Hex()] = bigNumber.NewFloat()
		rawPrices[chainID][address.Hex()] = bigNumber.NewInt()
		price, ok := storage.GetPrice(chainID, address)
		if !ok {
			continue
		}
		humanizedPrices[chainID][address.Hex()] = price.HumanizedPrice
		rawPrices[chainID][address.Hex()] = price.Price
	}
	if humanized {
		c.JSON(http.StatusOK, humanizedPrices)
	} else {
		c.JSON(http.StatusOK, rawPrices)
	}
}
