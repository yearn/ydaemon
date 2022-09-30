package partners

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/internal/types/common"
	"github.com/yearn/ydaemon/internal/utils/helpers"
)

// GetAllPartners will return all the partners informations, no matter the chainID.
func (y Controller) GetAllPartners(c *gin.Context) {
	allPartners := make(map[uint64]map[common.Address]*TPartners)
	for _, chainID := range helpers.SUPPORTED_CHAIN_IDS {
		partners := Store.PartnersByAddress[chainID]
		for _, partner := range partners {
			if _, ok := allPartners[chainID]; !ok {
				allPartners[chainID] = make(map[common.Address]*TPartners)
			}
			allPartners[chainID][partner.Treasury] = partner
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

	partners, ok := Store.PartnersByAddress[chainID]
	if !ok {
		c.String(http.StatusBadRequest, "no data available")
		return
	}

	c.JSON(http.StatusOK, partners)
}
