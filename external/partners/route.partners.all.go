package partners

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
)

// GetAllPartners will return all the partners informations, no matter the chainID.
func (y Controller) GetAllPartners(c *gin.Context) {
	allPartners := make(map[uint64]map[string]*TPartners)
	for _, chainID := range env.SUPPORTED_CHAIN_IDS {
		partners := partnersByAddress[chainID]
		for _, partner := range partners {
			if _, ok := allPartners[chainID]; !ok {
				allPartners[chainID] = make(map[string]*TPartners)
			}
			allPartners[chainID][partner.Treasury.Address().Hex()] = partner
		}
	}
	c.JSON(http.StatusOK, allPartners)
}

// GetPartners will, for a given chainID, return all the partners informations.
func (y Controller) GetPartners(c *gin.Context) {
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}

	partners, ok := partnersByAddress[chainID]
	if !ok {
		c.String(http.StatusBadRequest, "no data available")
		return
	}
	allPartnersAsHex := make(map[string]*TPartners)
	for address, partner := range partners {
		allPartnersAsHex[address.Address().Hex()] = partner
	}
	c.JSON(http.StatusOK, allPartnersAsHex)
}
