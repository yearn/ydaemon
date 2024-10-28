package vaults

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/sort"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

// getLegacyVaults will return a list of all vaults matching a provided filter function.
func getLegacyVaults(
	c *gin.Context,
	filterFunc func(vault models.TVault) bool,
) []TExternalVault {
	/** 🔵 - Yearn *************************************************************************************
	** page: A uint64 value that represents the page number for pagination. It is obtained from the
	** 'page' query parameter in the request. If the parameter is not provided, it defaults to 1.
	**
	** limit: A uint64 value that represents the number of vaults to be returned per page. It is
	** obtained from the 'limit' query parameter in the request. If the parameter is not provided,
	** it defaults to 200.
	**************************************************************************************************/
	page := helpers.SafeStringToUint64(getQuery(c, `page`), 1)
	limit := helpers.SafeStringToUint64(getQuery(c, `limit`), 200)

	/** 🔵 - Yearn *************************************************************************************
	** This function takes in a context from the gin framework. The context contains information
	** about the environment and request. From the context, the function extracts the following
	** parameters:
	**
	** hideAlways: A boolean value that determines whether to hide certain vaults. It is obtained
	** from the 'hideAlways' query parameter in the request. If the parameter is not provided,
	** it defaults to 'false'.
	**
	** orderBy: A string that determines the order in which the vaults are returned. It is obtained
	** from the 'orderBy' query parameter in the request. If the parameter is not provided,
	** it defaults to 'details.order'.
	**
	** orderDirection: A string that determines the direction of the ordering of the vaults. It is
	** obtained from the 'orderDirection' query parameter in the request. If the parameter is not
	** provided, it defaults to 'asc'.
	**
	** strategiesCondition: A string that determines the condition for selecting strategies. It is
	** obtained from the 'strategiesCondition' query parameter in the request.
	**
	** migrable: A string that determines the condition for selecting migrable vaults. It is
	** obtained from the 'migrable' query parameter in the request.
	**
	** chainID: A string that represents the chain ID of the vaults to be returned. It is obtained
	** from the 'chainID' path parameter in the request.
	**************************************************************************************************/
	hideAlways := helpers.StringToBool(helpers.SafeString(getQuery(c, `hideAlways`), `false`))
	orderBy := helpers.SafeString(getQuery(c, `orderBy`), `featuringScore`)
	orderDirection := helpers.SafeString(getQuery(c, `orderDirection`), `asc`)
	strategiesCondition := selectStrategiesCondition(getQuery(c, `strategiesCondition`))
	migrable := selectMigrableCondition(getQuery(c, `migrable`))
	chainID, ok := helpers.AssertChainID(c.Param(`chainID`))
	if !ok {
		c.String(http.StatusBadRequest, `invalid chainID`)
		return nil
	}
	if migrable != `none` && hideAlways {
		c.String(http.StatusBadRequest, `migrable and hideAlways cannot be true at the same time`)
		return nil
	}
	data := []TExternalVault{}
	allVaults, _ := storage.ListVaults(chainID)
	for _, currentVault := range allVaults {
		if !filterFunc(currentVault) {
			continue
		}
		chain, ok := env.GetChain(chainID)
		if !ok {
			continue
		}

		vaultAddress := currentVault.Address
		if helpers.Contains(chain.BlacklistedVaults, vaultAddress) {
			continue
		}

		newVault, err := NewVault().AssignTVault(currentVault)
		if err != nil {
			continue
		}
		if migrable == `none` && (newVault.Details.IsHidden || newVault.Details.IsRetired) && hideAlways {
			continue
		} else if migrable == `nodust` && (newVault.TVL.TVL < 100 || !newVault.Migration.Available) {
			continue
		} else if migrable == `all` && !newVault.Migration.Available {
			continue
		} else if migrable == `ignore` && (newVault.Migration.Available || newVault.Details.IsHidden || newVault.Details.IsRetired) {
			continue
		}

		vaultStrategies, _ := storage.ListStrategiesForVault(chainID, vaultAddress)
		newVault.Strategies = []TStrategy{}
		for _, strategy := range vaultStrategies {
			strategyWithDetails := NewStrategy().AssignTStrategy(strategy)
			if !strategyWithDetails.ShouldBeIncluded(strategiesCondition) {
				continue
			}

			newVault.Strategies = append(newVault.Strategies, strategyWithDetails)
		}

		data = append(data, newVault)
	}

	//Sort by details.order by default
	sort.SortBy(orderBy, orderDirection, data)
	start := (page - 1) * limit
	end := page * limit
	if end > uint64(len(data)) {
		end = uint64(len(data))
	}
	data = data[start:end]

	return data
}
