package tokens

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/models"
	"github.com/yearn/ydaemon/common/types/common"
	"github.com/yearn/ydaemon/external/meta"
	"github.com/yearn/ydaemon/internal/tokens"
)

type TAllTokens struct {
	Address       common.Address        `json:"address"`
	Name          string                `json:"name"`
	Symbol        string                `json:"symbol"`
	Decimals      uint64                `json:"decimals"`
	IsVault       bool                  `json:"isVault"`
	DisplayName   string                `json:"display_name,omitempty"`
	DisplaySymbol string                `json:"display_symbol,omitempty"`
	Description   string                `json:"description,omitempty"`
	Website       string                `json:"website,omitempty"`
	Categories    []string              `json:"categories,omitempty"`
	Localization  *models.TLocalization `json:"localization,omitempty"`
}

// GetAllTokens will return all the tokens informations, no matter the chainID.
func (y Controller) GetAllTokens(c *gin.Context) {
	localization := helpers.SafeString(c.Query("loc"), "en")
	allTokens := make(map[uint64]map[common.Address]TAllTokens)
	for _, chainID := range env.SUPPORTED_CHAIN_IDS {
		tokenList := tokens.ListTokens(chainID)
		for _, token := range tokenList {
			if _, ok := allTokens[chainID]; !ok {
				allTokens[chainID] = make(map[common.Address]TAllTokens)
			}
			metaToken, ok := meta.Store.TokensFromMeta[chainID][common.FromAddress(token.Address)]
			tokenDetails := TAllTokens{
				Address:  common.FromAddress(token.Address),
				Name:     token.Name,
				Symbol:   token.Symbol,
				Decimals: token.Decimals,
				IsVault:  token.IsVault(),
			}
			if ok {
				tokenDetails.DisplaySymbol = metaToken.Symbol
				tokenDetails.Website = metaToken.Website
				tokenDetails.Categories = metaToken.Categories

				if localization == "all" {
					tokenDetails.DisplayName = metaToken.Name
					tokenDetails.Description = metaToken.Description
					tokenDetails.Localization = metaToken.Localization
				} else {
					local := selectLocalizationFromString(localization, *metaToken.Localization)
					tokenDetails.DisplayName = local.Name
					tokenDetails.Description = local.Description
				}
			}
			allTokens[chainID][common.FromAddress(token.Address)] = tokenDetails
		}
	}
	c.JSON(http.StatusOK, allTokens)
}

//GetTokens will, for a given chainID, return a tokens list with the associated metadata
func (y Controller) GetTokens(c *gin.Context) {
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}
	localization := helpers.SafeString(c.Query("loc"), "en")

	allTokens := make(map[string]TAllTokens)
	tokenList := tokens.ListTokens(chainID)
	for _, token := range tokenList {
		metaToken, ok := meta.Store.TokensFromMeta[chainID][common.FromAddress(token.Address)]
		tokenDetails := TAllTokens{
			Address:  common.FromAddress(token.Address),
			Name:     token.Name,
			Symbol:   token.Symbol,
			Decimals: token.Decimals,
			IsVault:  token.IsVault(),
		}
		if ok {
			tokenDetails.DisplaySymbol = metaToken.Symbol
			tokenDetails.Website = metaToken.Website
			tokenDetails.Categories = metaToken.Categories

			if localization == "all" {
				tokenDetails.DisplayName = metaToken.Name
				tokenDetails.Description = metaToken.Description
				tokenDetails.Localization = metaToken.Localization
			} else {
				local := selectLocalizationFromString(localization, *metaToken.Localization)
				tokenDetails.DisplayName = local.Name
				tokenDetails.Description = local.Description
			}
		}
		allTokens[common.FromAddress(token.Address).Hex()] = tokenDetails
	}

	c.JSON(http.StatusOK, allTokens)
}
