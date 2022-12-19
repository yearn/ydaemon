package meta

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/meta"
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

	localization := helpers.SafeString(c.Query("loc"), "en")
	tokensFromMeta := meta.ListMetaTokens(chainID)
	if localization == "all" {
		c.JSON(http.StatusOK, tokensFromMeta)
		return
	}
	localizedTokensFromMeta := []*meta.TTokenFromMeta{}
	for _, token := range tokensFromMeta {
		local := selectLocalizationFromString(localization, *token.Localization)
		token.Name = local.Name
		token.Description = local.Description
		token.Localization = nil
		localizedTokensFromMeta = append(localizedTokensFromMeta, token)
	}
	c.JSON(http.StatusOK, localizedTokensFromMeta)
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

	localization := helpers.SafeString(c.Query("loc"), "en")
	tokenFromMeta, ok := meta.GetMetaToken(chainID, address)
	if !ok {
		c.String(http.StatusNotFound, "no data available")
		return
	}

	if localization == "all" {
		c.JSON(http.StatusOK, tokenFromMeta)
		return
	}
	local := selectLocalizationFromString(localization, *tokenFromMeta.Localization)
	tokenFromMeta.Name = local.Name
	tokenFromMeta.Description = local.Description
	tokenFromMeta.Localization = nil
	c.JSON(http.StatusOK, tokenFromMeta)
}
