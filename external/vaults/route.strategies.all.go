package vaults

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/sort"
	"github.com/yearn/ydaemon/internal/storage"
)

// GetAllStrategies will, for a given chainID, return a list of all strategies
func (y Controller) GetAllStrategies(c *gin.Context) {
	orderBy := helpers.SafeString(getQuery(c, `orderBy`), `address`)
	orderDirection := helpers.SafeString(getQuery(c, `orderDirection`), `asc`)
	strategiesCondition := selectStrategiesCondition(getQuery(c, `strategiesCondition`))
	chainID, ok := helpers.AssertChainID(c.Param(`chainID`))
	if !ok {
		c.String(http.StatusBadRequest, `invalid chainID`)
		return
	}

	data := []TStrategy{}
	allVaults, _ := storage.ListVaults(chainID)
	for _, currentVault := range allVaults {
		chain, ok := env.GetChain(chainID)
		if !ok {
			continue
		}
		if helpers.Contains(chain.BlacklistedVaults, currentVault.Address) {
			continue
		}
		vaultStrategies, _ := storage.ListStrategiesForVault(chainID, currentVault.Address)
		for _, strategy := range vaultStrategies {
			strategyWithDetails := NewStrategy().AssignTStrategy(strategy)
			if !strategyWithDetails.ShouldBeIncluded(strategiesCondition) {
				continue
			}
			data = append(data, strategyWithDetails)
		}
	}

	sort.SortBy(orderBy, orderDirection, data)
	c.JSON(http.StatusOK, data)
}
