package controllers

import (
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/majorfi/ydaemon/internal/models"
	"github.com/majorfi/ydaemon/internal/store"
	"github.com/majorfi/ydaemon/internal/utils"
)

// CountAllPartners will return the length of the partners per chainID
func (y controller) CountAllPartners(c *gin.Context) {
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

// GetAllPartners will return all the partners informations, no matter the chainID.
func (y controller) GetAllPartners(c *gin.Context) {
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
func (y controller) GetPartners(c *gin.Context) {
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

// GetPartner will, for a given address on given chainID, return the meta informations for the partner.
func (y controller) GetPartner(c *gin.Context) {
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
