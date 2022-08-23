package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/majorfi/ydaemon/internal/models"
	"github.com/majorfi/ydaemon/internal/store"
)

// GetMetaProtocolsLegacy will, for a given chainID, return all the meta informations for the protocols.
// The data will be resolved as-is, aka as an unorganized array of protocols metadata.
func (y controller) GetMetaProtocolsLegacy(c *gin.Context) {
	chainID, err := strconv.ParseUint(c.Param("chainID"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}
	localization := valueWithFallback(c.Query("loc"), "en")
	protocolsFromMeta, ok := store.RawMetaProtocols[chainID]
	if !ok {
		c.String(http.StatusNoContent, "no data available")
		return
	}

	if localization == "all" {
		c.JSON(http.StatusOK, protocolsFromMeta)
		return
	}
	localizedProtocolsFromMeta := []models.TProtocolsFromMeta{}
	for _, protocol := range protocolsFromMeta {
		local := selectLocalizationFromString(localization, *protocol.Localization)
		protocol.Name = local.Name
		protocol.Description = local.Description
		protocol.Localization = nil
		localizedProtocolsFromMeta = append(localizedProtocolsFromMeta, protocol)
	}
	c.JSON(http.StatusOK, localizedProtocolsFromMeta)
}

// GetMetaProtocols will, for a given chainID, return all the meta informations for the protocols.
// The data will be resolved as an object where the key is the checksummed address
// and the value the protocol metadata.
func (y controller) GetMetaProtocols(c *gin.Context) {
	chainID, err := strconv.ParseUint(c.Param("chainID"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}
	localization := valueWithFallback(c.Query("loc"), "en")
	protocolsFromMeta, ok := store.ProtocolsFromMeta[chainID]
	if !ok {
		c.String(http.StatusNoContent, "no data available")
		return
	}
	if localization == "all" {
		c.JSON(http.StatusOK, protocolsFromMeta)
		return
	}
	localizedProtocolsFromMeta := []models.TProtocolsFromMeta{}
	for _, protocol := range protocolsFromMeta {
		local := selectLocalizationFromString(localization, *protocol.Localization)
		protocol.Name = local.Name
		protocol.Description = local.Description
		protocol.Localization = nil
		localizedProtocolsFromMeta = append(localizedProtocolsFromMeta, protocol)
	}
	c.JSON(http.StatusOK, localizedProtocolsFromMeta)
}

// GetMetaProtocol will, for a given address on given chainID, return the meta informations for the protocol.
// The data will be resolved as an object corresponding to the protocol models.
func (y controller) GetMetaProtocol(c *gin.Context) {
	chainID, err := strconv.ParseUint(c.Param("chainID"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}
	localization := valueWithFallback(c.Query("loc"), "en")
	protocolName := c.Param("name")
	if protocolName == `` {
		c.String(http.StatusBadRequest, "invalid name")
		return
	}
	protocolFromMeta, ok := store.ProtocolsFromMeta[chainID][protocolName]
	if !ok {
		c.String(http.StatusNoContent, "no data available")
		return
	}

	if localization == "all" {
		c.JSON(http.StatusOK, protocolFromMeta)
		return
	}
	local := selectLocalizationFromString(localization, *protocolFromMeta.Localization)
	protocolFromMeta.Name = local.Name
	protocolFromMeta.Description = local.Description
	protocolFromMeta.Localization = nil
	c.JSON(http.StatusOK, protocolFromMeta)
}
