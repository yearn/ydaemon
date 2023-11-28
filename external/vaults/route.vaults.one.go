package vaults

import (
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/risk"
	"github.com/yearn/ydaemon/internal/storage"
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
	currentVault, ok := storage.GetVault(chainID, address)
	if !ok {
		c.String(http.StatusBadRequest, "invalid vault")
		return
	}

	newVault, err := NewVault().AssignTVault(currentVault)
	if err != nil {
		c.String(http.StatusBadRequest, "vault not found")
		return
	}
	vaultStrategies, _ := storage.ListStrategiesForVault(chainID, currentVault.Address)
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

	vaultAsStrategy, ok := storage.GetStrategy(newVault.ChainID, common.HexToAddress(newVault.Address))
	if ok {
		simplified := toSimplifiedVersion(newVault, vaultAsStrategy)
		simplified.Description = newVault.Description
		if simplified.Description == "" {
			simplified.Description = vaultAsStrategy.Description
		}
		c.JSON(http.StatusOK, simplified)
		return
	}
	simplified := toSimplifiedVersion(newVault, models.TStrategy{})
	simplified.Description = newVault.Description
	c.JSON(http.StatusOK, simplified)
}
