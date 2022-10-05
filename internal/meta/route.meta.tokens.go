package meta

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/internal/utils/env"
	"github.com/yearn/ydaemon/internal/utils/helpers"
)

// GetMetaTokensLegacy will, for a given chainID, return all the meta informations for the tokens.
// The data will be resolved as-is, aka as an unorganized array of tokens metadata.
func (y Controller) GetMetaTokensLegacy(c *gin.Context) {
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}

	localization := helpers.SafeString(c.Query("loc"), "en")
	tokensFromMeta, ok := Store.RawMetaTokens[chainID]
	if !ok {
		c.String(http.StatusBadRequest, "no data available")
		return
	}

	if localization == "all" {
		c.JSON(http.StatusOK, tokensFromMeta)
		return
	}
	localizedTokensFromMeta := []TTokenFromMeta{}
	for _, token := range tokensFromMeta {
		local := selectLocalizationFromString(localization, *token.Localization)
		token.Name = local.Name
		token.Description = local.Description
		token.Localization = nil
		localizedTokensFromMeta = append(localizedTokensFromMeta, token)
	}
	c.JSON(http.StatusOK, localizedTokensFromMeta)
}

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
	tokensFromMeta, ok := Store.TokensFromMeta[chainID]
	if !ok {
		c.String(http.StatusBadRequest, "no data available")
		return
	}

	if localization == "all" {
		c.JSON(http.StatusOK, tokensFromMeta)
		return
	}
	localizedTokensFromMeta := []TTokenFromMeta{}
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
	tokenFromMeta, ok := Store.TokensFromMeta[chainID][address]
	if !ok {
		c.String(http.StatusBadRequest, "no data available")
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

// GetTokenList is used to retrieve Yearn's Token List
func (y Controller) GetTokenList(c *gin.Context) {
	content, err := ioutil.ReadFile(env.BASE_DATA_PATH + `/meta/tokens/tokenList.json`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	var marshalDest map[string]interface{}
	if err := json.Unmarshal(content, &marshalDest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}

	delete(marshalDest, "$schema")

	c.JSON(http.StatusOK, marshalDest)
}
