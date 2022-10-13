package partners

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/env"
)

// CountAllPartners will return the length of the partners per chainID
func (y Controller) CountAllPartners(c *gin.Context) {
	total := uint64(0)
	allPartnersCount := make(map[uint64]uint64)
	for _, chainID := range env.SUPPORTED_CHAIN_IDS {
		allPartnersCount[chainID] = uint64(len(Store.PartnersByAddress[chainID]))
		total += allPartnersCount[chainID]
	}

	c.JSON(http.StatusOK, gin.H{
		"total":  total,
		"chains": allPartnersCount,
	})
}
