package allocation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/internal/utils/helpers"
)

func (y Controller) GetMedianScore(c *gin.Context) {
	chainID, ok := helpers.AssertChainID(c.Query("chain_id"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}
	address, ok := helpers.AssertAddress(c.Query("address"), chainID)
	if !ok {
		c.String(http.StatusBadRequest, "invalid address")
		return
	}

	allocation := Store.StrategyGroupAllocation[chainID][address]
	c.JSON(http.StatusOK, gin.H{
		"network":       chainID,
		"address":       address,
		"name":          allocation.Strategy.Name,
		"riskGroup":     allocation.RiskGroup.Label,
		"currentTVL":    allocation.CurrentTVL,
		"availableTVL":  allocation.AvailableTVL,
		"currentUSDC":   allocation.CurrentUSDC,
		"availableUSDC": allocation.AvailableUSDC,
	})
}
