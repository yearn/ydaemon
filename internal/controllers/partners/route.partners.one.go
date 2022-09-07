package partnersController

import (
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/internal/store"
)

// GetPartner will, for a given address on given chainID, return the meta informations for the partner.
func (y Controller) GetPartner(c *gin.Context) {
	chainID, err := strconv.ParseUint(c.Param("chainID"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}
	partnerAddressOrName := c.Param("addressOrName")
	if partnerAddressOrName == `` {
		c.String(http.StatusBadRequest, "invalid address or name")
		return
	}
	partner, ok := store.PartnersByName[chainID][partnerAddressOrName]
	if !ok {
		partner, ok = store.PartnersByAddress[chainID][common.HexToAddress(partnerAddressOrName)]
		if !ok {
			c.String(http.StatusBadRequest, "no data available")
			return
		}
	}

	c.JSON(http.StatusOK, partner)
}
