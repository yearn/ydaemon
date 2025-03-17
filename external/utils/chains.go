package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/env"
)

/**************************************************************************************************
** GetSupportedChains returns a list of all chains supported by the API.
**
** This function handles HTTP GET requests to the /info/chains endpoint, returning a JSON response
** with all chain IDs that are currently supported by the system.
**
** @param c *gin.Context - The Gin context for the HTTP request
**************************************************************************************************/
func GetSupportedChains(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"chains": env.CHAINS,
	})
}
