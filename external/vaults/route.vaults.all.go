package vaults

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/sort"
	"github.com/yearn/ydaemon/common/types/common"
	"github.com/yearn/ydaemon/internal/strategies"
	"github.com/yearn/ydaemon/internal/vaults"
)

// GetAllVaults will, for a given chainID, return a list of all vaults
func (y Controller) GetAllVaults(c *gin.Context) {
	hideAlways := helpers.StringToBool(helpers.SafeString(c.Query(`hideAlways`), `false`))
	orderBy := helpers.SafeString(c.Query(`orderBy`), `details.order`)
	orderDirection := helpers.SafeString(c.Query(`orderDirection`), `asc`)
	strategiesCondition := selectStrategiesCondition(c.Query(`strategiesCondition`))
	withStrategiesDetails := c.Query(`strategiesDetails`) == `withDetails`
	migrable := selectMigrableCondition(c.Query(`migrable`))
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
	allVaults := vaults.ListVaults(chainID)
	for _, currentVault := range allVaults {
		vaultAddress := common.FromAddress(currentVault.Address)
		if helpers.Contains(env.BLACKLISTED_VAULTS[chainID], vaultAddress) {
			continue
		}

		newVault := NewVault().AssignTVault(currentVault)
		if migrable == `none` && (newVault.Details.HideAlways || newVault.Details.Retired) && hideAlways {
			continue
		} else if migrable == `nodust` && (newVault.TVL.TVL < 100 || !newVault.Migration.Available) {
			continue
		} else if migrable == `all` && !newVault.Migration.Available {
			continue
		} else if migrable == `ignore` && (newVault.Migration.Available || newVault.Details.HideAlways || newVault.Details.Retired) {
			continue
		}

		vaultStrategies := strategies.ListStrategiesForVault(chainID, vaultAddress)
		newVault.Strategies = []*TStrategy{}
		for _, strategy := range vaultStrategies {
			var externalStrategy *TStrategy
			strategyWithDetails := NewStrategy().AssignTStrategy(strategy)
			if !strategyWithDetails.ShouldBeIncluded(strategiesCondition) {
				continue
			}

			if withStrategiesDetails {
				externalStrategy = strategyWithDetails
				externalStrategy.Risk = NewRiskScore().AssignTStrategyFromRisk(strategy.BuildRiskScore())
			} else {
				externalStrategy = &TStrategy{
					Address:     common.FromAddress(strategy.Address),
					Name:        strategy.Name,
					DisplayName: strategy.DisplayName,
					Description: strategy.Description,
				}
			}
			newVault.Strategies = append(newVault.Strategies, externalStrategy)
		}
		if withStrategiesDetails {
			newVault.RiskScore = newVault.ComputeRiskScore()
		}

		data = append(data, *newVault)
	}

	// Preparing the sort. This specifics steps are needed to avoid a panic
	var sortedData = make([]interface{}, len(data))
	for i, d := range data {
		sortedData[i] = d
	}
	sort.SortBy(sortedData, orderBy, orderDirection) //Sort by details.order by default

	c.JSON(http.StatusOK, sortedData)
}
