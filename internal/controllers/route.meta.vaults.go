package controllers

import (
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/majorfi/ydaemon/internal/store"
)

// GetMetaVaultsLegacy will, for a given chainID, return all the meta informations for the vaults.
// The data will be resolved as-is, aka as an unorganized array of vaults metadata.
func (y controller) GetMetaVaultsLegacy(c *gin.Context) {
	chainID, err := strconv.ParseUint(c.Param("chainID"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}
	vaultsFromMeta, ok := store.RawMetaVaults[chainID]
	if !ok {
		c.String(http.StatusNoContent, "no data available")
		return
	}

	c.JSON(http.StatusOK, vaultsFromMeta)
}

// GetMetaVaults will, for a given chainID, return all the meta informations for the vaults.
// The data will be resolved as an object where the key is the checksummed address
// and the value the vault metadata.
func (y controller) GetMetaVaults(c *gin.Context) {
	chainID, err := strconv.ParseUint(c.Param("chainID"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}
	vaultsFromMeta, ok := store.VaultsFromMeta[chainID]
	if !ok {
		c.String(http.StatusNoContent, "no data available")
		return
	}

	c.JSON(http.StatusOK, vaultsFromMeta)
}

// GetMetaVault will, for a given address on given chainID, return the meta informations for the vault.
// The data will be resolved as an object corresponding to the vault models.
func (y controller) GetMetaVault(c *gin.Context) {
	chainID, err := strconv.ParseUint(c.Param("chainID"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}
	vaultAddress := c.Param("address")
	if vaultAddress == `` {
		c.String(http.StatusBadRequest, "invalid address")
		return
	}
	vaultFromMeta, ok := store.VaultsFromMeta[chainID][common.HexToAddress(vaultAddress)]
	if !ok {
		c.String(http.StatusNoContent, "no data available")
		return
	}

	c.JSON(http.StatusOK, vaultFromMeta)
}
