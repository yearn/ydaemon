package tokens

import (
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/internal/meta"
	"github.com/yearn/ydaemon/internal/utils/helpers"
	"github.com/yearn/ydaemon/internal/utils/models"
)

type TAllTokens struct {
	Address       common.Address        `json:"address"`
	Name          string                `json:"name"`
	Symbol        string                `json:"symbol"`
	Price         float64               `json:"price"`
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
	localization := helpers.ValueWithFallback(c.Query("loc"), "en")
	allTokens := make(map[uint64]map[common.Address]TAllTokens)
	for _, chainID := range helpers.SUPPORTED_CHAIN_IDS {
		tokenList := Store.Tokens[chainID]
		for _, token := range tokenList {
			if _, ok := allTokens[chainID]; !ok {
				allTokens[chainID] = make(map[common.Address]TAllTokens)
			}
			metaToken, ok := meta.Store.TokensFromMeta[chainID][token.Address]
			tokenDetails := TAllTokens{
				Address:  token.Address,
				Name:     token.Name,
				Symbol:   token.Symbol,
				Price:    token.Price,
				Decimals: token.Decimals,
				IsVault:  token.IsVault,
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
			allTokens[chainID][token.Address] = tokenDetails
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
	localization := helpers.ValueWithFallback(c.Query("loc"), "en")

	allTokens := make(map[common.Address]TAllTokens)
	tokenList := Store.Tokens[chainID]
	for _, token := range tokenList {
		metaToken, ok := meta.Store.TokensFromMeta[chainID][token.Address]
		tokenDetails := TAllTokens{
			Address:  token.Address,
			Name:     token.Name,
			Symbol:   token.Symbol,
			Price:    token.Price,
			Decimals: token.Decimals,
			IsVault:  token.IsVault,
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
		allTokens[token.Address] = tokenDetails
	}

	c.JSON(http.StatusOK, allTokens)
}
