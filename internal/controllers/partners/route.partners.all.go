package partnersController

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/store"
	"github.com/yearn/ydaemon/internal/utils"
)

// GetAllPartners will return all the partners informations, no matter the chainID.
func (y Controller) GetAllPartners(c *gin.Context) {
	allPartners := make(map[uint64][]models.TPartners)
	for _, chainID := range utils.SUPPORTED_CHAIN_IDS {
		partners := store.PartnersByAddress[chainID]
		for _, partner := range partners {
			allPartners[chainID] = append(allPartners[chainID], partner)
		}
	}
	c.JSON(http.StatusOK, allPartners)
}

// GetPartners will, for a given chainID, return all the partners informations.
func (y Controller) GetPartners(c *gin.Context) {
	chainID, err := strconv.ParseUint(c.Param("chainID"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}
	partners, ok := store.PartnersByAddress[chainID]
	if !ok {
		c.String(http.StatusBadRequest, "no data available")
		return
	}

	c.JSON(http.StatusOK, partners)
}
