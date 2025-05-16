package vaults

import (
	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/internal/models"
)

/**************************************************************************************************
** GetAll is a gin handler function to retrieve all vaults across all supported chains.
**
** This endpoint returns all vaults in the system without filtering, though it still respects
** query parameters for sorting, pagination, and strategy conditions.
**
** Supported query parameters:
** - orderBy: Field to sort results by (default: 'featuringScore')
** - orderDirection: Sort direction, 'asc' or 'desc' (default: 'asc')
** - strategiesCondition: Strategy filter condition (default: 'debtRatio')
** - hideAlways: Whether to hide certain vaults (default: false)
** - migrable: Condition for including migrable vaults (default: 'none')
** - page/limit: Pagination controls (defaults: page 1, limit 200)
** - chainIDs: Comma-separated list of chain IDs to include
**
** Endpoint: GET /vaults
**
** @param c *gin.Context - The Gin context containing the HTTP request
** @return []TSimplifiedExternalVault - The list of all vaults
** @return error - Any error encountered during processing
**************************************************************************************************/
func (y Controller) GetAll(c *gin.Context) ([]TSimplifiedExternalVault, error) {
	return getVaults(c, func(vault models.TVault) bool {
		return true
	})
}

/**************************************************************************************************
** GetIsYearn retrieves only vaults that are officially part of Yearn Finance.
**
** This endpoint returns vaults marked with the 'isYearn' flag, filtering out third-party
** or non-Yearn vaults that might be in the system.
**
** Endpoint: GET /vaults/yearn
**
** @param c *gin.Context - The Gin context containing the HTTP request
** @return []TSimplifiedExternalVault - The list of Yearn vaults
** @return error - Any error encountered during processing
**************************************************************************************************/
func (y Controller) GetIsYearn(c *gin.Context) ([]TSimplifiedExternalVault, error) {
	return getVaults(c, func(vault models.TVault) bool {
		return vault.Metadata.Inclusion.IsYearn
	})
}

/**************************************************************************************************
** GetV3 retrieves only vaults that use the V3 vault architecture.
**
** This endpoint returns vaults that are built using the V3 vault architecture, which
** includes both single-strategy and multi-strategy vaults.
**
** Endpoint: GET /vaults/v3
**
** @param c *gin.Context - The Gin context containing the HTTP request
** @return []TSimplifiedExternalVault - The list of V3 vaults
** @return error - Any error encountered during processing
**************************************************************************************************/
func (y Controller) GetV3(c *gin.Context) ([]TSimplifiedExternalVault, error) {
	return getVaults(c, func(vault models.TVault) bool {
		return isV3Vault(vault)
	})
}

/**************************************************************************************************
** GetV2 retrieves only vaults that use the V2 vault architecture.
**
** This endpoint returns vaults that are built using the V2 vault architecture, which
** preceded the V3 architecture and is still widely used.
**
** Endpoint: GET /vaults/v2
**
** @param c *gin.Context - The Gin context containing the HTTP request
** @return []TSimplifiedExternalVault - The list of V2 vaults
** @return error - Any error encountered during processing
**************************************************************************************************/
func (y Controller) GetV2(c *gin.Context) ([]TSimplifiedExternalVault, error) {
	return getVaults(c, func(vault models.TVault) bool {
		return isV2Vault(vault)
	})
}

/**************************************************************************************************
** GetRetired retrieves vaults that have been marked as retired.
**
** This endpoint returns vaults that are no longer actively maintained or supported,
** but may still contain user funds or be of historical interest.
**
** Endpoint: GET /vaults/retired
**
** @param c *gin.Context - The Gin context containing the HTTP request
** @return []TSimplifiedExternalVault - The list of retired vaults
** @return error - Any error encountered during processing
**************************************************************************************************/
func (y Controller) GetRetired(c *gin.Context) ([]TSimplifiedExternalVault, error) {
	return getVaults(c, func(vault models.TVault) bool {
		return vault.Metadata.IsRetired
	})
}

/**************************************************************************************************
** GetV2IsYearn retrieves V2 vaults that are officially part of Yearn Finance.
**
** This endpoint combines the V2 and IsYearn filters to return only V2 vaults
** that are also marked as official Yearn vaults.
**
** Endpoint: GET /vaults/v2/yearn
**
** @param c *gin.Context - The Gin context containing the HTTP request
** @return []TSimplifiedExternalVault - The list of V2 Yearn vaults
** @return error - Any error encountered during processing
**************************************************************************************************/
func (y Controller) GetV2IsYearn(c *gin.Context) ([]TSimplifiedExternalVault, error) {
	return getVaults(c, func(vault models.TVault) bool {
		return isV2Vault(vault) && vault.Metadata.Inclusion.IsYearn
	})
}

/**************************************************************************************************
** GetV3IsYearn retrieves V3 vaults that are officially part of Yearn Finance.
**
** This endpoint combines the V3 and IsYearn filters to return only V3 vaults
** that are also marked as official Yearn vaults.
**
** Endpoint: GET /vaults/v3/yearn
**
** @param c *gin.Context - The Gin context containing the HTTP request
** @return []TSimplifiedExternalVault - The list of V3 Yearn vaults
** @return error - Any error encountered during processing
**************************************************************************************************/
func (y Controller) GetV3IsYearn(c *gin.Context) ([]TSimplifiedExternalVault, error) {
	return getVaults(c, func(vault models.TVault) bool {
		return isV3Vault(vault) && vault.Metadata.Inclusion.IsYearn
	})
}

/**************************************************************************************************
** GetIsYearnJuiced is a gin handler function to retrieve all the vaults matching the
** inclusion.IsYearnJuiced filter.
**************************************************************************************************/
func (y Controller) GetIsYearnJuiced(c *gin.Context) ([]TSimplifiedExternalVault, error) {
	return getVaults(c, func(vault models.TVault) bool {
		return vault.Metadata.Inclusion.IsYearnJuiced
	})
}

/**************************************************************************************************
** GetIsGimme is a gin handler function to retrieve all the vaults matching the inclusion.IsGimme
** filter.
**************************************************************************************************/
func (y Controller) GetIsGimme(c *gin.Context) ([]TSimplifiedExternalVault, error) {
	return getVaults(c, func(vault models.TVault) bool {
		return vault.Metadata.Inclusion.IsGimme
	})
}

/**************************************************************************************************
** GetIsYearnPendle is a gin handler function to retrieve all the vaults matching the pendle
** category and the inclusion.IsYearn filter.
**************************************************************************************************/
func (y Controller) GetIsYearnPendle(c *gin.Context) ([]TSimplifiedExternalVault, error) {
	return getVaults(c, func(vault models.TVault) bool {
		return vault.Metadata.Inclusion.IsYearn && vault.Metadata.Category == "Pendle" || vault.Metadata.Category == "Pendle Autorollover"
	})
}

/**************************************************************************************************
** GetIsOptimism is a gin handler function to retrieve all the optimism vaults matching the
** inclusion.IsYearn filter.
**************************************************************************************************/
func (y Controller) GetIsOptimism(c *gin.Context) ([]TSimplifiedExternalVault, error) {
	return getVaults(c, func(vault models.TVault) bool {
		return vault.Metadata.Inclusion.IsYearn && vault.ChainID == 10
	})
}

/**************************************************************************************************
** GetIsYearnPoolTogether is a gin handler function to retrieve all the vaults matching the
** inclusion.isPoolTogether filter.
**************************************************************************************************/
func (y Controller) GetIsYearnPoolTogether(c *gin.Context) ([]TSimplifiedExternalVault, error) {
	return getVaults(c, func(vault models.TVault) bool {
		return vault.Metadata.Inclusion.IsPoolTogether
	})
}

/**************************************************************************************************
** GetIsYearnCove is a gin handler function to retrieve all the vaults matching the
** inclusion.isCove filter.
**************************************************************************************************/
func (y Controller) GetIsYearnCove(c *gin.Context) ([]TSimplifiedExternalVault, error) {
	return getVaults(c, func(vault models.TVault) bool {
		return vault.Metadata.Inclusion.IsCove
	})
}

/**************************************************************************************************
** GetIsMorpho is a gin handler function to retrieve all the vaults matching the
** inclusion.isMorpho filter.
**************************************************************************************************/
func (y Controller) GetIsMorpho(c *gin.Context) ([]TSimplifiedExternalVault, error) {
	return getVaults(c, func(vault models.TVault) bool {
		return vault.Metadata.Inclusion.IsMorpho
	})
}

/**************************************************************************************************
** GetIsKatana is a gin handler function to retrieve all the vaults matching the
** inclusion.isKatana filter.
**************************************************************************************************/
func (y Controller) GetIsKatana(c *gin.Context) ([]TSimplifiedExternalVault, error) {
	return getVaults(c, func(vault models.TVault) bool {
		return vault.Metadata.Inclusion.IsKatana
	})
}

/**************************************************************************************************
** GetIsAjna is a gin handler function to retrieve all the vaults with a name matching `AJNA` or
** with the inclusion.IsYearnJuiced filter.
**************************************************************************************************/
func (y Controller) GetIsAjna(c *gin.Context) ([]TSimplifiedExternalVault, error) {
	return getVaults(c, func(vault models.TVault) bool {
		return vault.Metadata.Inclusion.IsYearn && vault.Metadata.Category == "Ajna"
	})
}

/**************************************************************************************************
** GetIsVelodrome is a gin handler function to retrieve all the vaults with a name matching the
** Velodrome category
**************************************************************************************************/
func (y Controller) GetIsVelodrome(c *gin.Context) ([]TSimplifiedExternalVault, error) {
	return getVaults(c, func(vault models.TVault) bool {
		return vault.Metadata.Inclusion.IsYearn && vault.Metadata.Category == "Velodrome"
	})
}

/**************************************************************************************************
** GetIsAerodrome is a gin handler function to retrieve all the vaults with a name matching the
** Aerodrome category
**************************************************************************************************/
func (y Controller) GetIsAerodrome(c *gin.Context) ([]TSimplifiedExternalVault, error) {
	return getVaults(c, func(vault models.TVault) bool {
		return vault.Metadata.Inclusion.IsYearn && vault.Metadata.Category == "Aerodrome"
	})
}

/**************************************************************************************************
** GetIsCurve is a gin handler function to retrieve all the vaults with a name matching the
** Curve category
**************************************************************************************************/
func (y Controller) GetIsCurve(c *gin.Context) ([]TSimplifiedExternalVault, error) {
	return getVaults(c, func(vault models.TVault) bool {
		return vault.Metadata.Inclusion.IsYearn && vault.Metadata.Category == "Curve"
	})
}
