package meta

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/helpers"
)

// GetMetaVaultsLegacy will, for a given chainID, return all the meta informations for the vaults.
// The data will be resolved as-is, aka as an unorganized array of vaults metadata.
func (y Controller) GetMetaVaultsLegacy(c *gin.Context) {
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}

	vaultsFromMeta, ok := Store.RawMetaVaults[chainID]
	if !ok {
		c.String(http.StatusBadRequest, "no data available")
		return
	}

	c.JSON(http.StatusOK, vaultsFromMeta)
}

// GetMetaVaults will, for a given chainID, return all the meta informations for the vaults.
// The data will be resolved as an object where the key is the checksummed address
// and the value the vault metadata.
func (y Controller) GetMetaVaults(c *gin.Context) {
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}

	vaultsFromMeta, ok := Store.VaultsFromMeta[chainID]
	if !ok {
		c.String(http.StatusBadRequest, "no data available")
		return
	}

	c.JSON(http.StatusOK, vaultsFromMeta)
}

// GetMetaVault will, for a given address on given chainID, return the meta informations for the vault.
// The data will be resolved as an object corresponding to the vault models.
func (y Controller) GetMetaVault(c *gin.Context) {
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}
	address, ok := helpers.AssertAddress(c.Param("address"), chainID)
	if !ok {
		c.String(http.StatusBadRequest, "invalid address")
		return
	}
	vaultFromMeta, ok := Store.VaultsFromMeta[chainID][address]
	if !ok {
		c.String(http.StatusBadRequest, "no data available")
		return
	}

	c.JSON(http.StatusOK, vaultFromMeta)
}
