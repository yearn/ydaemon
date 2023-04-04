package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/bribes"
)

// GetRewardAdded returns the feed of reward added for the yBribeV3 contract
func GetRewardAdded(c *gin.Context) {
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}

	c.JSON(http.StatusOK, bribes.ListRewardAdded(chainID))
}
