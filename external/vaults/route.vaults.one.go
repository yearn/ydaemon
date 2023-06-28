package vaults

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/strategies"
	"github.com/yearn/ydaemon/internal/vaults"
)

// GetVault will, for a given chainID, return a list of all vaults
func (y Controller) GetVault(c *gin.Context) {
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}
	address, ok := helpers.AssertAddress(c.Param("address"), chainID)
	if !ok {
		c.String(http.StatusBadRequest, "invalid address")
		return
	}

	strategiesCondition := selectStrategiesCondition(getQuery(c, "strategiesCondition"))
	withStrategiesDetails := strings.EqualFold(getQuery(c, "strategiesDetails"), "withDetails")
	currentVault, ok := vaults.FindVault(chainID, address)
	if !ok {
		c.String(http.StatusBadRequest, "invalid vault")
		return
	}
	newVault := NewVault().AssignTVault(currentVault)

	vaultStrategies := strategies.ListStrategiesForVault(chainID, currentVault.Address)
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

	c.JSON(http.StatusOK, newVault)
}
