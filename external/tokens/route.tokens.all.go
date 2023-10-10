package tokens

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/storage"
)

type Controller struct{}

type TAllTokens struct {
	Address          string   `json:"address"`
	Name             string   `json:"name"`
	Symbol           string   `json:"symbol"`
	Decimals         uint64   `json:"decimals"`
	IsVault          bool     `json:"isVault"`
	DisplayName      string   `json:"display_name,omitempty"`
	DisplaySymbol    string   `json:"display_symbol,omitempty"`
	Description      string   `json:"description,omitempty"`
	Category         string   `json:"category,omitempty"`
	UnderlyingTokens []string `json:"underlyingTokens,omitempty"`
}

// GetAllTokens will return all the tokens informations, no matter the chainID.
func (y Controller) GetAllTokens(c *gin.Context) {
	allTokens := make(map[uint64]map[string]TAllTokens)
	for chainID := range env.CHAINS {
		tokenList, _ := storage.ListERC20(chainID)
		for _, token := range tokenList {
			if _, ok := allTokens[chainID]; !ok {
				allTokens[chainID] = make(map[string]TAllTokens)
			}
			tokenDetails := TAllTokens{
				Address:       token.Address.Hex(),
				Name:          token.Name,
				DisplayName:   token.DisplayName,
				Symbol:        token.Symbol,
				Decimals:      token.Decimals,
				DisplaySymbol: token.Symbol,
				Category:      token.Category,
				Description:   token.Description,
				IsVault:       token.IsVaultLike(),
			}
			for _, addr := range token.UnderlyingTokensAddresses {
				tokenDetails.UnderlyingTokens = append(tokenDetails.UnderlyingTokens, addr.Hex())
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
	allTokens := make(map[string]TAllTokens)
	tokenMap, _ := storage.ListERC20(chainID)
	for _, token := range tokenMap {
		tokenDetails := TAllTokens{
			Address:       token.Address.Hex(),
			Name:          token.Name,
			DisplayName:   token.DisplayName,
			Symbol:        token.Symbol,
			Decimals:      token.Decimals,
			DisplaySymbol: token.Symbol,
			Category:      token.Category,
			Description:   token.Description,
			IsVault:       token.IsVaultLike(),
		}
		for _, addr := range token.UnderlyingTokensAddresses {
			tokenDetails.UnderlyingTokens = append(tokenDetails.UnderlyingTokens, addr.Hex())
		}
		allTokens[token.Address.Hex()] = tokenDetails
	}

	c.JSON(http.StatusOK, allTokens)
}
