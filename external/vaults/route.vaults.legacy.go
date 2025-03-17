package vaults

import (
	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/internal/models"
)

/**************************************************************************************************
** GetLegacyIsYearn retrieves all Yearn vaults in the legacy format.
**
** This endpoint returns all vaults that are officially part of Yearn Finance in the original
** TExternalVault format rather than the newer simplified format. It filters vaults using the
** IsYearn inclusion flag, ensuring only official Yearn vaults are included.
**
** The function leverages the getLegacyVaults helper which supports comprehensive query parameters:
** - page & limit: For pagination control (default: page 1, limit 200)
** - hideAlways: Whether to hide vaults marked as hidden or retired (default: false)
** - orderBy: Field to sort results by (default: 'featuringScore')
** - orderDirection: Sort direction, 'asc' or 'desc' (default: 'asc')
** - strategiesCondition: Filter for strategies (default: 'debtRatio')
** - migrable: Filter for migrable vaults (default: 'none')
**
** Endpoint: GET /vaults/legacy/yearn
**
** @param c *gin.Context - The Gin context containing the HTTP request
** @return []TExternalVault - The list of Yearn vaults in legacy format
**************************************************************************************************/
func (y Controller) GetLegacyIsYearn(c *gin.Context) []TExternalVault {
	return getLegacyVaults(c, func(vault models.TVault) bool {
		return vault.Metadata.Inclusion.IsYearn
	})
}

/**************************************************************************************************
** GetLegacyV2IsYearn retrieves all V2 Yearn vaults in the legacy format.
**
** This endpoint returns vaults that use the V2 vault architecture and are officially part of
** Yearn Finance, formatted in the original TExternalVault structure. It combines the V2
** architecture filter and the IsYearn inclusion flag.
**
** The function leverages the getLegacyVaults helper which supports comprehensive query parameters:
** - page & limit: For pagination control (default: page 1, limit 200)
** - hideAlways: Whether to hide vaults marked as hidden or retired (default: false)
** - orderBy: Field to sort results by (default: 'featuringScore')
** - orderDirection: Sort direction, 'asc' or 'desc' (default: 'asc')
** - strategiesCondition: Filter for strategies (default: 'debtRatio')
** - migrable: Filter for migrable vaults (default: 'none')
**
** Endpoint: GET /vaults/legacy/v2/yearn
**
** @param c *gin.Context - The Gin context containing the HTTP request
** @return []TExternalVault - The list of V2 Yearn vaults in legacy format
**************************************************************************************************/
func (y Controller) GetLegacyV2IsYearn(c *gin.Context) []TExternalVault {
	return getLegacyVaults(c, func(vault models.TVault) bool {
		return isV2Vault(vault) && vault.Metadata.Inclusion.IsYearn
	})
}

/**************************************************************************************************
** GetLegacyV3IsYearn retrieves all V3 Yearn vaults in the legacy format.
**
** This endpoint returns vaults that use the V3 vault architecture and are officially part of
** Yearn Finance, formatted in the original TExternalVault structure. It combines the V3
** architecture filter and the IsYearn inclusion flag.
**
** The function leverages the getLegacyVaults helper which supports comprehensive query parameters:
** - page & limit: For pagination control (default: page 1, limit 200)
** - hideAlways: Whether to hide vaults marked as hidden or retired (default: false)
** - orderBy: Field to sort results by (default: 'featuringScore')
** - orderDirection: Sort direction, 'asc' or 'desc' (default: 'asc')
** - strategiesCondition: Filter for strategies (default: 'debtRatio')
** - migrable: Filter for migrable vaults (default: 'none')
**
** Endpoint: GET /vaults/legacy/v3/yearn
**
** @param c *gin.Context - The Gin context containing the HTTP request
** @return []TExternalVault - The list of V3 Yearn vaults in legacy format
**************************************************************************************************/
func (y Controller) GetLegacyV3IsYearn(c *gin.Context) []TExternalVault {
	return getLegacyVaults(c, func(vault models.TVault) bool {
		return isV3Vault(vault) && vault.Metadata.Inclusion.IsYearn
	})
}

/**************************************************************************************************
** GetLegacyIsYearnJuiced retrieves all Juiced Yearn vaults in the legacy format.
**
** This endpoint returns vaults that are part of the Yearn Juice system in the original
** TExternalVault format. Juiced vaults participate in the Yearn Juice tokenomics system
** which provides additional incentives for participants.
**
** The function leverages the getLegacyVaults helper which supports comprehensive query parameters:
** - page & limit: For pagination control (default: page 1, limit 200)
** - hideAlways: Whether to hide vaults marked as hidden or retired (default: false)
** - orderBy: Field to sort results by (default: 'featuringScore')
** - orderDirection: Sort direction, 'asc' or 'desc' (default: 'asc')
** - strategiesCondition: Filter for strategies (default: 'debtRatio')
** - migrable: Filter for migrable vaults (default: 'none')
**
** Endpoint: GET /vaults/legacy/juiced
**
** @param c *gin.Context - The Gin context containing the HTTP request
** @return []TExternalVault - The list of Juiced Yearn vaults in legacy format
**************************************************************************************************/
func (y Controller) GetLegacyIsYearnJuiced(c *gin.Context) []TExternalVault {
	return getLegacyVaults(c, func(vault models.TVault) bool {
		return vault.Metadata.Inclusion.IsYearnJuiced
	})
}

/**************************************************************************************************
** GetLegacyIsGimme retrieves all Gimme-branded vaults in the legacy format.
**
** This endpoint returns vaults that are part of the Gimme partnership or branding in the
** original TExternalVault format. Gimme vaults are a specific type of product within the
** Yearn ecosystem.
**
** The function leverages the getLegacyVaults helper which supports comprehensive query parameters:
** - page & limit: For pagination control (default: page 1, limit 200)
** - hideAlways: Whether to hide vaults marked as hidden or retired (default: false)
** - orderBy: Field to sort results by (default: 'featuringScore')
** - orderDirection: Sort direction, 'asc' or 'desc' (default: 'asc')
** - strategiesCondition: Filter for strategies (default: 'debtRatio')
** - migrable: Filter for migrable vaults (default: 'none')
**
** Endpoint: GET /vaults/legacy/gimme
**
** @param c *gin.Context - The Gin context containing the HTTP request
** @return []TExternalVault - The list of Gimme vaults in legacy format
**************************************************************************************************/
func (y Controller) GetLegacyIsGimme(c *gin.Context) []TExternalVault {
	return getLegacyVaults(c, func(vault models.TVault) bool {
		return vault.Metadata.Inclusion.IsGimme
	})
}

/**************************************************************************************************
** GetLegacyRetired retrieves all retired vaults in the legacy format.
**
** This endpoint returns vaults that have been officially retired or deprecated but are still
** accessible for historical or withdrawal purposes. The vaults are returned in the original
** TExternalVault format rather than the newer simplified format.
**
** Retired vaults typically represent older investment strategies that are no longer actively
** maintained or recommended for new deposits, but may still contain user funds that can be
** withdrawn.
**
** The function leverages the getLegacyVaults helper which supports comprehensive query parameters:
** - page & limit: For pagination control (default: page 1, limit 200)
** - hideAlways: Whether to hide vaults marked as hidden (default: false)
** - orderBy: Field to sort results by (default: 'featuringScore')
** - orderDirection: Sort direction, 'asc' or 'desc' (default: 'asc')
** - strategiesCondition: Filter for strategies (default: 'debtRatio')
** - migrable: Filter for migrable vaults (default: 'none')
**
** Endpoint: GET /vaults/legacy/retired
**
** @param c *gin.Context - The Gin context containing the HTTP request
** @return []TExternalVault - The list of retired vaults in legacy format
**************************************************************************************************/
func (y Controller) GetLegacyRetired(c *gin.Context) []TExternalVault {
	return getLegacyVaults(c, func(vault models.TVault) bool {
		return vault.Metadata.IsRetired
	})
}
