package vaults

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/sort"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/strategies"
	"github.com/yearn/ydaemon/internal/vaults"
)

// GetAllVaultsForAllChains will return a list of all vaults for all chains
func (y Controller) GetAllVaultsForAllChains(c *gin.Context) {
	hideAlways := helpers.StringToBool(helpers.SafeString(getQuery(c, `hideAlways`), `false`))
	orderBy := helpers.SafeString(getQuery(c, `orderBy`), `chainID`)
	orderDirection := helpers.SafeString(getQuery(c, `orderDirection`), `asc`)
	strategiesCondition := selectStrategiesCondition(getQuery(c, `strategiesCondition`))
	withStrategiesDetails := strings.EqualFold(getQuery(c, `strategiesDetails`), `withDetails`)
	migrable := selectMigrableCondition(getQuery(c, `migrable`))
	if migrable != `none` && hideAlways {
		c.String(http.StatusBadRequest, `migrable and hideAlways cannot be true at the same time`)
		return
	}

	data := []TExternalVault{}
	allVaults := []models.TVault{}
	for _, chainID := range env.SUPPORTED_CHAIN_IDS {
		vaultsForChain := vaults.ListVaults(chainID)
		for _, currentVault := range vaultsForChain {
			if helpers.Contains(env.BLACKLISTED_VAULTS[chainID], currentVault.Address) {
				continue
			}
			allVaults = append(allVaults, currentVault)
		}
	}
	for _, currentVault := range allVaults {
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

		vaultStrategies := strategies.ListStrategiesForVault(currentVault.ChainID, currentVault.Address)
		newVault.Strategies = []*TStrategy{}
		for _, strategy := range vaultStrategies {
			var externalStrategy *TStrategy
			strategyWithDetails := NewStrategy().AssignTStrategy(strategy)
			if !strategyWithDetails.ShouldBeIncluded(strategiesCondition) {
				continue
			}

			if withStrategiesDetails {
				externalStrategy = strategyWithDetails
				externalStrategy.Risk = NewRiskScore().AssignTStrategyFromRisk(strategies.BuildRiskScore(strategy))
			} else {
				externalStrategy = &TStrategy{
					Address:     strategy.Address.Hex(),
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

	//Sort by details.order by default
	sort.SortBy(orderBy, orderDirection, data)
	c.JSON(http.StatusOK, data)
}
