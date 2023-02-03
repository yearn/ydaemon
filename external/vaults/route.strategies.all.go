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

// GetAllStrategies will, for a given chainID, return a list of all strategies
func (y Controller) GetAllStrategies(c *gin.Context) {
	orderBy := helpers.SafeString(c.Query(`orderBy`), `details.order`)
	orderDirection := helpers.SafeString(c.Query(`orderDirection`), `asc`)
	strategiesCondition := selectStrategiesCondition(c.Query(`strategiesCondition`))
	chainID, ok := helpers.AssertChainID(c.Param(`chainID`))
	if !ok {
		c.String(http.StatusBadRequest, `invalid chainID`)
		return
	}

	data := []TStrategy{}
	allVaults := vaults.ListVaults(chainID)
	for _, currentVault := range allVaults {
		vaultAddress := common.FromAddress(currentVault.Address)
		if helpers.Contains(env.BLACKLISTED_VAULTS[chainID], vaultAddress) {
			continue
		}
		vaultStrategies := strategies.ListStrategiesForVault(chainID, vaultAddress)
		for _, strategy := range vaultStrategies {
			strategyWithDetails := NewStrategy().AssignTStrategy(strategy)
			if !strategyWithDetails.ShouldBeIncluded(strategiesCondition) {
				continue
			}
			// Always show details
			strategyWithDetails.Risk = NewRiskScore().AssignTStrategyFromRisk(strategy.BuildRiskScore())
			data = append(data, *strategyWithDetails)
		}
	}

	// Preparing the sort. This specifics steps are needed to avoid a panic
	var sortedData = make([]interface{}, len(data))
	for i, d := range data {
		sortedData[i] = d
	}
	sort.SortBy(sortedData, orderBy, orderDirection) //Sort by details.order by default

	c.JSON(http.StatusOK, sortedData)

}