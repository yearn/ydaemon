package prices

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/storage"
)

// GetAllPrices will return all the prices informations, no matter the chainID.
func (y Controller) GetAllPricesWithDetails(c *gin.Context) {
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}
	allChainPrices, _ := storage.ListPrices(chainID)
	c.JSON(http.StatusOK, allChainPrices)
}
