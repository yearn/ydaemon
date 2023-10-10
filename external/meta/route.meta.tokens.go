package meta

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/storage"
)

// GetMetaTokens will, for a given chainID, return all the meta informations for the tokens.
// The data will be resolved as an object where the key is the checksummed address
// and the value the token metadata.
func (y Controller) GetMetaTokens(c *gin.Context) {
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}
	_, asList := storage.ListERC20(chainID)
	c.JSON(http.StatusOK, asList)
}

// GetMetaToken will, for a given address on given chainID, return the meta informations for the token.
// The data will be resolved as an object corresponding to the token models.
func (y Controller) GetMetaToken(c *gin.Context) {
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
	token, ok := storage.GetERC20(chainID, address)
	if !ok {
		c.String(http.StatusNotFound, "token not found")
		return
	}
	c.JSON(http.StatusOK, token)
}
