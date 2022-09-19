package tokens

import (
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/internal/utils/helpers"
)

// GetAllTokens will return all the tokens informations, no matter the chainID.
func (y Controller) GetAllTokens(c *gin.Context) {
	allTokens := make(map[uint64]map[common.Address]*TERC20Token)
	for _, chainID := range helpers.SUPPORTED_CHAIN_IDS {
		tokens := Store.Tokens[chainID]
		for _, token := range tokens {
			if _, ok := allTokens[chainID]; !ok {
				allTokens[chainID] = make(map[common.Address]*TERC20Token)
			}
			allTokens[chainID][token.Address] = token
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

	c.JSON(http.StatusOK, Store.Tokens[chainID])
}
