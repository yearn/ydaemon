package vaults

import (
	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/internal/models"
)

/**************************************************************************************************
** GetLegacyIsYearn is a gin handler function to retrieve all the vaults matching the
** inclusion.IsYearn filter.
**************************************************************************************************/
func (y Controller) GetLegacyIsYearn(c *gin.Context) []TExternalVault {
	return getLegacyVaults(c, func(vault models.TVault) bool {
		return vault.Metadata.Inclusion.IsYearn
	})
}

/**************************************************************************************************
** GetLegacyV2IsYearn is a gin handler function to retrieve all the V2 vaults matching the
** inclusion.IsYearn filter.
**************************************************************************************************/
func (y Controller) GetLegacyV2IsYearn(c *gin.Context) []TExternalVault {
	return getLegacyVaults(c, func(vault models.TVault) bool {
		return isV2Vault(vault) && vault.Metadata.Inclusion.IsYearn
	})
}

/**************************************************************************************************
** GetLegacyV3IsYearn is a gin handler function to retrieve all the V3 vaults matching the
** inclusion.IsYearn filter.
**************************************************************************************************/
func (y Controller) GetLegacyV3IsYearn(c *gin.Context) []TExternalVault {
	return getLegacyVaults(c, func(vault models.TVault) bool {
		return isV3Vault(vault) && vault.Metadata.Inclusion.IsYearn
	})
}

/**************************************************************************************************
** GetLegacyIsYearnJuiced is a gin handler function to retrieve all the vaults matching the
** inclusion.IsYearnJuiced filter.
**************************************************************************************************/
func (y Controller) GetLegacyIsYearnJuiced(c *gin.Context) []TExternalVault {
	return getLegacyVaults(c, func(vault models.TVault) bool {
		return vault.Metadata.Inclusion.IsYearnJuiced
	})
}

/**************************************************************************************************
** GetLegacyIsGimme is a gin handler function to retrieve all the vaults matching the inclusion.IsGimme
** filter.
**************************************************************************************************/
func (y Controller) GetLegacyIsGimme(c *gin.Context) []TExternalVault {
	return getLegacyVaults(c, func(vault models.TVault) bool {
		return vault.Metadata.Inclusion.IsGimme
	})
}

/**************************************************************************************************
** GetLegacyRetired is a gin handler function to retrieve all the retired vaults
**************************************************************************************************/
func (y Controller) GetLegacyRetired(c *gin.Context) []TExternalVault {
	return getLegacyVaults(c, func(vault models.TVault) bool {
		return vault.Metadata.IsRetired
	})
}
