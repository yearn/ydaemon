package vaults

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/sort"
	"github.com/yearn/ydaemon/internal/risk"
	"github.com/yearn/ydaemon/internal/storage"
)

// GetSomeVaults will, for a given chainID, return a list of all vaults
func (y Controller) GetSomeVaults(c *gin.Context) {
	orderBy := helpers.SafeString(getQuery(c, `orderBy`), `featuringScore`)
	orderDir := helpers.SafeString(getQuery(c, `orderDirection`), `asc`)
	stratDet := getQuery(c, `strategiesDetails`) == `withDetails`
	stratCon := selectStrategiesCondition(getQuery(c, `strategiesCondition`))
	chainID, ok := helpers.AssertChainID(c.Param(`chainID`))
	if !ok {
		c.String(http.StatusBadRequest, `invalid chainID`)
		return
	}

	addressesStr := strings.Split(strings.ToLower(c.Param(`addresses`)), `,`)
	data := []TExternalVault{}
	allVaults, _ := storage.ListVaults(chainID)
	for _, currentVault := range allVaults {
		if !helpers.Contains(addressesStr, strings.ToLower(currentVault.Address.Hex())) {
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
		vaultStrategies, _ := storage.ListStrategiesForVault(chainID, vaultAddress)
		newVault.Strategies = []TStrategy{}
		for _, strategy := range vaultStrategies {
			var externalStrategy TStrategy
			strategyWithDetails := NewStrategy().AssignTStrategy(strategy)
			if !strategyWithDetails.ShouldBeIncluded(stratCon) {
				continue
			}

			if stratDet {
				externalStrategy = strategyWithDetails
				externalStrategy.Risk = NewRiskScore().AssignTStrategyFromRisk(risk.BuildRiskScore(strategy))
			} else {
				externalStrategy = strategyWithDetails
			}
			newVault.Strategies = append(newVault.Strategies, externalStrategy)
		}
		if stratDet {
			newVault.RiskScore = newVault.ComputeRiskScore()
		}

		data = append(data, newVault)
	}

	//Sort by details.order by default
	sort.SortBy(orderBy, orderDir, data)
	c.JSON(http.StatusOK, data)
}
