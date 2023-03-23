package tokens

import (
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/meta"
	"github.com/yearn/ydaemon/internal/tokens"
)

type TAllTokens struct {
	Address          common.MixedcaseAddress `json:"address"`
	Name             string                  `json:"name"`
	Symbol           string                  `json:"symbol"`
	Decimals         uint64                  `json:"decimals"`
	IsVault          bool                    `json:"isVault"`
	DisplayName      string                  `json:"display_name,omitempty"`
	DisplaySymbol    string                  `json:"display_symbol,omitempty"`
	Description      string                  `json:"description,omitempty"`
	Website          string                  `json:"website,omitempty"`
	Categories       []string                `json:"categories,omitempty"`
	UnderlyingTokens []string                `json:"underlyingTokens,omitempty"`
	Localization     *meta.TLocalization     `json:"localization,omitempty"`
}

// GetAllTokens will return all the tokens informations, no matter the chainID.
func (y Controller) GetAllTokens(c *gin.Context) {
	localization := helpers.SafeString(c.Query("loc"), "en")
	allTokens := make(map[uint64]map[string]TAllTokens)
	for _, chainID := range env.SUPPORTED_CHAIN_IDS {
		tokenList := tokens.ListTokens(chainID)
		for _, token := range tokenList {
			if _, ok := allTokens[chainID]; !ok {
				allTokens[chainID] = make(map[string]TAllTokens)
			}
			metaToken, ok := meta.GetMetaToken(chainID, token.Address)
			tokenDetails := TAllTokens{
				Address:  common.NewMixedcaseAddress(token.Address),
				Name:     token.Name,
				Symbol:   token.Symbol,
				Decimals: token.Decimals,
				IsVault:  token.IsVault(),
			}
			for _, addr := range token.UnderlyingTokensAddresses {
				tokenDetails.UnderlyingTokens = append(tokenDetails.UnderlyingTokens, addr.Hex())
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
					if metaToken.Localization != nil {
						local := selectLocalizationFromString(localization, *metaToken.Localization)
						tokenDetails.DisplayName = local.Name
						tokenDetails.Description = local.Description
					}
				}
			}
			allTokens[chainID][token.Address.Hex()] = tokenDetails
		}
	}
	c.JSON(http.StatusOK, allTokens)
}

// GetTokens will, for a given chainID, return a tokens list with the associated metadata
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
		metaToken, ok := meta.GetMetaToken(chainID, token.Address)
		tokenDetails := TAllTokens{
			Address:  common.NewMixedcaseAddress(token.Address),
			Name:     token.Name,
			Symbol:   token.Symbol,
			Decimals: token.Decimals,
			IsVault:  token.IsVault(),
		}
		for _, addr := range token.UnderlyingTokensAddresses {
			tokenDetails.UnderlyingTokens = append(tokenDetails.UnderlyingTokens, addr.Hex())
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
				if metaToken.Localization != nil {
					local := selectLocalizationFromString(localization, *metaToken.Localization)
					tokenDetails.DisplayName = local.Name
					tokenDetails.Description = local.Description
				}
			}
		}
		allTokens[token.Address.Hex()] = tokenDetails
	}

	c.JSON(http.StatusOK, allTokens)
}
