package meta

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/helpers"
)

// GetMetaProtocolsLegacy will, for a given chainID, return all the meta informations for the protocols.
// The data will be resolved as-is, aka as an unorganized array of protocols metadata.
func (y Controller) GetMetaProtocolsLegacy(c *gin.Context) {
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}

	localization := helpers.SafeString(c.Query("loc"), "en")
	protocolsFromMeta, ok := Store.RawMetaProtocols[chainID]
	if !ok {
		c.String(http.StatusBadRequest, "no data available")
		return
	}

	if localization == "all" {
		c.JSON(http.StatusOK, protocolsFromMeta)
		return
	}
	localizedProtocolsFromMeta := []TProtocolsFromMeta{}
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
func (y Controller) GetMetaProtocols(c *gin.Context) {
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}

	localization := helpers.SafeString(c.Query("loc"), "en")
	protocolsFromMeta, ok := Store.ProtocolsFromMeta[chainID]
	if !ok {
		c.String(http.StatusBadRequest, "no data available")
		return
	}
	if localization == "all" {
		c.JSON(http.StatusOK, protocolsFromMeta)
		return
	}
	localizedProtocolsFromMeta := []TProtocolsFromMeta{}
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
func (y Controller) GetMetaProtocol(c *gin.Context) {
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}

	localization := helpers.SafeString(c.Query("loc"), "en")
	protocolName := c.Param("name")
	protocolFromMeta, ok := Store.ProtocolsFromMeta[chainID][protocolName]
	if !ok {
		c.String(http.StatusBadRequest, "no data available")
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
