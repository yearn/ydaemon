package vaults

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/strategies"
)

// GetStrategy will, for a given chainID, return the strategy for a given address
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

	strategy, ok := strategies.FindStrategy(chainID, address)
	if !ok {
		c.String(http.StatusBadRequest, "invalid strategy")
		return
	}
	// Always show details
	newStrategy := NewStrategy().AssignTStrategy(strategy)
	newStrategy.Risk = NewRiskScore().AssignTStrategyFromRisk(strategies.BuildRiskScore(strategy))

	c.JSON(http.StatusOK, *newStrategy)
}
