package meta

import (
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/internal/utils/helpers"
)

// GetMetaTokensLegacy will, for a given chainID, return all the meta informations for the tokens.
// The data will be resolved as-is, aka as an unorganized array of tokens metadata.
func (y Controller) GetMetaTokensLegacy(c *gin.Context) {
	chainID, err := strconv.ParseUint(c.Param("chainID"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}
	localization := helpers.ValueWithFallback(c.Query("loc"), "en")
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
	chainID, err := strconv.ParseUint(c.Param("chainID"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}
	localization := helpers.ValueWithFallback(c.Query("loc"), "en")
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
	chainID, err := strconv.ParseUint(c.Param("chainID"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}
	localization := helpers.ValueWithFallback(c.Query("loc"), "en")
	tokenAddress := c.Param("address")
	if tokenAddress == `` {
		c.String(http.StatusBadRequest, "invalid address")
		return
	}
	tokenFromMeta, ok := Store.TokensFromMeta[chainID][common.HexToAddress(tokenAddress)]
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
