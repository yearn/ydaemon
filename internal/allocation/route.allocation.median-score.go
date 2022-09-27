package allocation

import (
	"math/big"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/internal/utils/helpers"
)

// BigFloatNumberJSON exists so the output in JSON is pretty and not using exponent notation
type BigFloatNumberJSON struct{ *big.Float }

func (bfn BigFloatNumberJSON) MarshalJSON() ([]byte, error) {
	return []byte(bfn.String()), nil
}

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
	if c.Query("debtToAdd") != "" {
		debtToAdd, debtToAddErr := strconv.ParseFloat(c.Query("debtToAdd"), 64)

		if debtToAddErr == nil {
			c.String(http.StatusOK, strconv.FormatBool(big.NewFloat(debtToAdd).Cmp(allocation.AvailableTVL) < 0))
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"network":       chainID,
		"address":       address,
		"name":          allocation.Strategy.Name,
		"riskGroup":     allocation.RiskGroup.Label,
		"currentTVL":    BigFloatNumberJSON{allocation.CurrentTVL},
		"availableTVL":  BigFloatNumberJSON{allocation.AvailableTVL},
		"currentUSDC":   BigFloatNumberJSON{allocation.CurrentUSDC},
		"availableUSDC": BigFloatNumberJSON{allocation.AvailableUSDC},
	})
}
