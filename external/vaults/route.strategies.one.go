package vaults

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/types/common"
	"github.com/yearn/ydaemon/internal/strategies"
)

// GetVault will, for a given chainID, return a list of all vaults
func (y Controller) GetStrategy(c *gin.Context) {
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

	withStrategiesDetails := c.Query("strategiesDetails") == "withDetails"
	strategy, ok := strategies.FindStrategy(chainID, address.ToAddress())
	if !ok {
		c.String(http.StatusBadRequest, "invalid strategy")
		return
	}
	var newStrategy *TStrategy
	if withStrategiesDetails {
		newStrategy = NewStrategy().AssignTStrategy(strategy)
		newStrategy.Risk = NewRiskScore().AssignTStrategyFromRisk(strategy.BuildRiskScore())
	} else {
		newStrategy = &TStrategy{
			Address:     common.FromAddress(strategy.Address),
			Name:        strategy.Name,
			Description: strategy.Description,
		}
	}
	c.JSON(http.StatusOK, *newStrategy)
}