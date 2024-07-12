package vaults

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/sort"
	"github.com/yearn/ydaemon/internal/storage"
)

// GetSomeVaults will, for a given chainID, return a list of all vaults
func (y Controller) GetLegacySomeVaults(c *gin.Context) {
	orderBy := helpers.SafeString(getQuery(c, `orderBy`), `featuringScore`)
	orderDir := helpers.SafeString(getQuery(c, `orderDirection`), `asc`)
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
			strategyWithDetails := NewStrategy().AssignTStrategy(strategy)
			if !strategyWithDetails.ShouldBeIncluded(stratCon) {
				continue
			}

			newVault.Strategies = append(newVault.Strategies, strategyWithDetails)
		}

		data = append(data, newVault)
	}

	//Sort by details.order by default
	sort.SortBy(orderBy, orderDir, data)
	c.JSON(http.StatusOK, data)
}
