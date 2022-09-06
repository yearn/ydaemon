package partnersController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/internal/store"
	"github.com/yearn/ydaemon/internal/utils"
)

// CountAllPartners will return the length of the partners per chainID
func (y Controller) CountAllPartners(c *gin.Context) {
	total := uint64(0)
	allPartnersCount := make(map[uint64]uint64)
	for _, chainID := range utils.SUPPORTED_CHAIN_IDS {
		allPartnersCount[chainID] = uint64(len(store.PartnersByAddress[chainID]))
		total += allPartnersCount[chainID]
	}

	c.JSON(http.StatusOK, gin.H{
		"total":  total,
		"chains": allPartnersCount,
	})
}
