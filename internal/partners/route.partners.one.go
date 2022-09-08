package partners

import (
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/internal/utils/helpers"
)

// GetPartner will, for a given address on given chainID, return the meta informations for the partner.
func (y Controller) GetPartner(c *gin.Context) {
	chainID, ok := helpers.AssertChainID(c, c.Param("chainID"))
	if !ok {
		return
	}

	partnerAddressOrName := c.Param("addressOrName")
	partner, ok := Store.PartnersByName[chainID][partnerAddressOrName]
	if !ok {
		partner, ok = Store.PartnersByAddress[chainID][common.HexToAddress(partnerAddressOrName)]
		if !ok {
			c.String(http.StatusBadRequest, "no data available")
			return
		}
	}

	partner = partner.BalanceOf()

	c.JSON(http.StatusOK, partner)
}
