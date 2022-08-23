package controllers

import (
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/majorfi/ydaemon/internal/store"
)

// GetMetaStrategiesLegacy will, for a given chainID, return all the meta informations for the strategies.
// The data will be resolved as-is, aka as an unorganized array of strategies metadata.
func (y controller) GetMetaStrategiesLegacy(c *gin.Context) {
	chainID, err := strconv.ParseUint(c.Param("chainID"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}
	strategiesFromMeta, ok := store.RawMetaStrategies[chainID]
	if !ok {
		c.String(http.StatusNoContent, "no data available")
		return
	}

	c.JSON(http.StatusOK, strategiesFromMeta)
}

// GetMetaStrategies will, for a given chainID, return all the meta informations for the strategies.
// The data will be resolved as an object where the key is the checksummed address
// and the value the strategy metadata.
func (y controller) GetMetaStrategies(c *gin.Context) {
	chainID, err := strconv.ParseUint(c.Param("chainID"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}
	strategiesFromMeta, ok := store.StrategiesFromMeta[chainID]
	if !ok {
		c.String(http.StatusNoContent, "no data available")
		return
	}

	c.JSON(http.StatusOK, strategiesFromMeta)
}

// GetMetaStrategy will, for a given address on given chainID, return the meta informations for the strategy.
// The data will be resolved as an object corresponding to the strategy models.
func (y controller) GetMetaStrategy(c *gin.Context) {
	chainID, err := strconv.ParseUint(c.Param("chainID"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}
	strategyAddress := c.Param("address")
	if strategyAddress == `` {
		c.String(http.StatusBadRequest, "invalid address")
		return
	}
	strategyFromMeta, ok := store.StrategiesFromMeta[chainID][common.HexToAddress(strategyAddress)]
	if !ok {
		c.String(http.StatusNoContent, "no data available")
		return
	}

	c.JSON(http.StatusOK, strategyFromMeta)
}
