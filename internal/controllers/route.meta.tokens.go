package controllers

import (
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/majorfi/ydaemon/internal/store"
)

// GetMetaTokensLegacy will, for a given chainID, return all the meta informations for the tokens.
// The data will be resolved as-is, aka as an unorganized array of tokens metadata.
func (y controller) GetMetaTokensLegacy(c *gin.Context) {
	chainID, err := strconv.ParseUint(c.Param("chainID"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}
	tokensFromMeta, ok := store.RawMetaTokens[chainID]
	if !ok {
		c.String(http.StatusNoContent, "no data available")
		return
	}

	c.JSON(http.StatusOK, tokensFromMeta)
}

// GetMetaTokens will, for a given chainID, return all the meta informations for the tokens.
// The data will be resolved as an object where the key is the checksummed address
// and the value the token metadata.
func (y controller) GetMetaTokens(c *gin.Context) {
	chainID, err := strconv.ParseUint(c.Param("chainID"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}
	tokensFromMeta, ok := store.TokensFromMeta[chainID]
	if !ok {
		c.String(http.StatusNoContent, "no data available")
		return
	}

	c.JSON(http.StatusOK, tokensFromMeta)
}

// GetMetaToken will, for a given address on given chainID, return the meta informations for the token.
// The data will be resolved as an object corresponding to the token models.
func (y controller) GetMetaToken(c *gin.Context) {
	chainID, err := strconv.ParseUint(c.Param("chainID"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}
	tokenAddress := c.Param("address")
	if tokenAddress == `` {
		c.String(http.StatusBadRequest, "invalid address")
		return
	}
	tokenFromMeta, ok := store.TokensFromMeta[chainID][common.HexToAddress(tokenAddress)]
	if !ok {
		c.String(http.StatusNoContent, "no data available")
		return
	}

	c.JSON(http.StatusOK, tokenFromMeta)
}
