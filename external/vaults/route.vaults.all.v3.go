package vaults

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/sort"
	"github.com/yearn/ydaemon/internal/risk"
	"github.com/yearn/ydaemon/internal/storage"
)

// GetAllV3Vaults will, for a given chainID, return a list of all vaults
func (y Controller) GetAllV3Vaults(c *gin.Context) {
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
	** withStrategiesDetails: A boolean value that determines whether to include details of the
	** strategies in the response. It is obtained from the 'strategiesDetails' query parameter in
	** the request. If the parameter is 'withDetails', this value is true.
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
	withStrategiesDetails := getQuery(c, `strategiesDetails`) == `withDetails`
	migrable := selectMigrableCondition(getQuery(c, `migrable`))
	chainID, ok := helpers.AssertChainID(c.Param(`chainID`))
	if !ok {
		c.String(http.StatusBadRequest, `invalid chainID`)
		return
	}
	if migrable != `none` && hideAlways {
		c.String(http.StatusBadRequest, `migrable and hideAlways cannot be true at the same time`)
		return
	}
	data := []TExternalVault{}
	allVaults, _ := storage.ListVaults(chainID)
	for _, currentVault := range allVaults {
		/******************************************************************************************
		** We want to ignore all non Yearn vaults
		******************************************************************************************/
		if !currentVault.Metadata.Inclusion.IsYearn {
			continue
		}
		if !isV3Vault(currentVault) {
			continue
		}
		vaultAddress := currentVault.Address
		if helpers.Contains(env.CHAINS[chainID].BlacklistedVaults, vaultAddress) {
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
			var externalStrategy TStrategy
			strategyWithDetails := NewStrategy().AssignTStrategy(strategy)
			if !strategyWithDetails.ShouldBeIncluded(strategiesCondition) {
				continue
			}

			if withStrategiesDetails {
				externalStrategy = strategyWithDetails
				externalStrategy.Risk = NewRiskScore().AssignTStrategyFromRisk(risk.BuildRiskScore(strategy))
			} else {
				externalStrategy = strategyWithDetails
			}
			newVault.Strategies = append(newVault.Strategies, externalStrategy)
		}
		if withStrategiesDetails {
			newVault.RiskScore = newVault.ComputeRiskScore()
		}

		data = append(data, newVault)
	}

	//Sort by details.order by default
	sort.SortBy(orderBy, orderDirection, data)
	c.JSON(http.StatusOK, data)
}
