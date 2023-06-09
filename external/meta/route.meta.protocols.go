package meta

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/meta"
	"github.com/yearn/ydaemon/internal/models"
)

// GetMetaProtocols will, for a given chainID, return all the meta informations for the protocols.
// The data will be resolved as an object where the key is the checksummed address
// and the value the protocol metadata.
func (y Controller) GetMetaProtocols(c *gin.Context) {
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}

	localization := helpers.SafeString(getQuery(c, "loc"), "en")
	protocolsFromMeta := meta.ListMetaProtocol(chainID)
	if localization == "all" {
		c.JSON(http.StatusOK, protocolsFromMeta)
		return
	}
	localizedProtocolsFromMeta := []*models.TProtocolsFromMeta{}
	for _, protocol := range protocolsFromMeta {
		if protocol.Localization == nil {
			protocol.Localization = nil
			localizedProtocolsFromMeta = append(localizedProtocolsFromMeta, protocol)
			continue
		}
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

	localization := helpers.SafeString(getQuery(c, "loc"), "en")
	protocolName := c.Param("name")
	protocolFromMeta, ok := meta.GetMetaProtocol(chainID, protocolName)
	if !ok {
		c.String(http.StatusNotFound, "no data available")
		return
	}

	if localization == "all" {
		c.JSON(http.StatusOK, protocolFromMeta)
		return
	}
	if protocolFromMeta.Localization == nil {
		protocolFromMeta.Localization = nil
		c.JSON(http.StatusOK, protocolFromMeta)
		return
	}
	local := selectLocalizationFromString(localization, *protocolFromMeta.Localization)
	protocolFromMeta.Name = local.Name
	protocolFromMeta.Description = local.Description
	protocolFromMeta.Localization = nil
	c.JSON(http.StatusOK, protocolFromMeta)
}
