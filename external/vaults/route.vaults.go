package vaults

import (
	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/internal/models"
)

/**************************************************************************************************
** GetIsYearn is a gin handler function to retrieve all the vaults matching the
** inclusion.IsYearn filter.
**************************************************************************************************/
func (y Controller) GetIsYearn(c *gin.Context) []TSimplifiedExternalVault {
	return getVaults(c, func(vault models.TVault) bool {
		return vault.Metadata.Inclusion.IsYearn
	})
}

/**************************************************************************************************
** GetV3 is a gin handler function to retrieve all the v3 vaults
**************************************************************************************************/
func (y Controller) GetV3(c *gin.Context) []TSimplifiedExternalVault {
	return getVaults(c, func(vault models.TVault) bool {
		return isV3Vault(vault)
	})
}

/**************************************************************************************************
** GetV2 is a gin handler function to retrieve all the v2 vaults
**************************************************************************************************/
func (y Controller) GetV2(c *gin.Context) []TSimplifiedExternalVault {
	return getVaults(c, func(vault models.TVault) bool {
		return isV2Vault(vault)
	})
}

/**************************************************************************************************
** GetV2 is a gin handler function to retrieve all the retired vaults
**************************************************************************************************/
func (y Controller) GetRetired(c *gin.Context) []TSimplifiedExternalVault {
	return getVaults(c, func(vault models.TVault) bool {
		return vault.Metadata.IsRetired
	})
}

/**************************************************************************************************
** GetV2IsYearn is a gin handler function to retrieve all the V2 vaults matching the
** inclusion.IsYearn filter.
**************************************************************************************************/
func (y Controller) GetV2IsYearn(c *gin.Context) []TSimplifiedExternalVault {
	return getVaults(c, func(vault models.TVault) bool {
		return isV2Vault(vault) && vault.Metadata.Inclusion.IsYearn
	})
}

/**************************************************************************************************
** GetV3IsYearn is a gin handler function to retrieve all the V3 vaults matching the
** inclusion.IsYearn filter.
**************************************************************************************************/
func (y Controller) GetV3IsYearn(c *gin.Context) []TSimplifiedExternalVault {
	return getVaults(c, func(vault models.TVault) bool {
		return isV3Vault(vault) && vault.Metadata.Inclusion.IsYearn
	})
}

/**************************************************************************************************
** GetIsYearnJuiced is a gin handler function to retrieve all the vaults matching the
** inclusion.IsYearnJuiced filter.
**************************************************************************************************/
func (y Controller) GetIsYearnJuiced(c *gin.Context) []TSimplifiedExternalVault {
	return getVaults(c, func(vault models.TVault) bool {
		return vault.Metadata.Inclusion.IsYearnJuiced
	})
}

/**************************************************************************************************
** GetIsGimme is a gin handler function to retrieve all the vaults matching the inclusion.IsGimme
** filter.
**************************************************************************************************/
func (y Controller) GetIsGimme(c *gin.Context) []TSimplifiedExternalVault {
	return getVaults(c, func(vault models.TVault) bool {
		return vault.Metadata.Inclusion.IsGimme
	})
}

/**************************************************************************************************
** GetIsYearnPendle is a gin handler function to retrieve all the vaults matching the pendle
** category and the inclusion.IsYearn filter.
**************************************************************************************************/
func (y Controller) GetIsYearnPendle(c *gin.Context) []TSimplifiedExternalVault {
	return getVaults(c, func(vault models.TVault) bool {
		return vault.Metadata.Inclusion.IsYearn && vault.Metadata.Category == "Pendle" || vault.Metadata.Category == "Pendle Autorollover"
	})
}

/**************************************************************************************************
** GetIsOptimism is a gin handler function to retrieve all the optimism vaults matching the
** inclusion.IsYearn filter.
**************************************************************************************************/
func (y Controller) GetIsOptimism(c *gin.Context) []TSimplifiedExternalVault {
	return getVaults(c, func(vault models.TVault) bool {
		return vault.Metadata.Inclusion.IsYearn && vault.ChainID == 10
	})
}

/**************************************************************************************************
** GetIsYearnPoolTogether is a gin handler function to retrieve all the vaults matching the
** inclusion.isPoolTogether filter.
**************************************************************************************************/
func (y Controller) GetIsYearnPoolTogether(c *gin.Context) []TSimplifiedExternalVault {
	return getVaults(c, func(vault models.TVault) bool {
		return vault.Metadata.Inclusion.IsPoolTogether
	})
}
