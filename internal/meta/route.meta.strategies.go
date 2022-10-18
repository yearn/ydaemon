package meta

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/internal/utils/helpers"
)

// GetMetaStrategies will, for a given chainID, return all the meta informations for the strategies.
// The data will be resolved as-is, aka as an unorganized array of strategies metadata.
func (y Controller) GetMetaStrategies(c *gin.Context) {
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}

	localization := helpers.SafeString(c.Query("loc"), "en")
	strategiesFromMeta, ok := Store.RawMetaStrategies[chainID]
	if !ok {
		c.String(http.StatusBadRequest, "no data available")
		return
	}

	if localization == "all" {
		c.JSON(http.StatusOK, strategiesFromMeta)
		return
	}
	localizedStrategiesFromMeta := []TStrategyFromMeta{}
	for _, strategy := range strategiesFromMeta {
		local := selectLocalizationFromString(localization, *strategy.Localization)
		strategy.Name = local.Name
		strategy.Description = local.Description
		strategy.Localization = nil
		localizedStrategiesFromMeta = append(localizedStrategiesFromMeta, strategy)
	}
	c.JSON(http.StatusOK, localizedStrategiesFromMeta)
}

// GetMetaStrategy will, for a given address on given chainID, return the meta informations for the strategy.
// The data will be resolved as an object corresponding to the strategy models.
func (y Controller) GetMetaStrategy(c *gin.Context) {
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

	localization := helpers.SafeString(c.Query("loc"), "en")
	strategyFromMeta, ok := Store.StrategiesFromMeta[chainID][address]
	if !ok {
		c.String(http.StatusBadRequest, "no data available")
		return
	}

	if localization == "all" {
		c.JSON(http.StatusOK, strategyFromMeta)
		return
	}
	local := selectLocalizationFromString(localization, *strategyFromMeta.Localization)
	strategyFromMeta.Name = local.Name
	strategyFromMeta.Description = local.Description
	strategyFromMeta.Localization = nil
	c.JSON(http.StatusOK, strategyFromMeta)
}
