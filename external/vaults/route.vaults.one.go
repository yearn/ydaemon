package vaults

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/types/common"
	"github.com/yearn/ydaemon/internal/strategies"
	"github.com/yearn/ydaemon/internal/vaults"
)

//GetVault will, for a given chainID, return a list of all vaults
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

	// strategiesCondition := selectStrategiesCondition(c.Query("strategiesCondition"))
	withStrategiesDetails := c.Query("strategiesDetails") == "withDetails"
	withStrategiesRisk := c.Query("strategiesRisk") == "withRisk"
	currentVault, ok := vaults.FindVault(chainID, address)
	if !ok {
		c.String(http.StatusBadRequest, "invalid vault")
		return
	}
	vaultAddress := common.FromAddress(currentVault.Address)
	newVault := NewVault().AssignTVault(currentVault)

	vaultStrategies := strategies.ListStrategiesForVault(chainID, vaultAddress)
	for _, strategy := range vaultStrategies {
		var externalStrategy *TStrategy
		if withStrategiesDetails {
			externalStrategy = NewStrategy().AssignTStrategy(strategy)
		} else {
			externalStrategy = &TStrategy{
				Address:     common.FromAddress(strategy.Address),
				Name:        strategy.Name,
				Description: strategy.Description,
			}
		}
		if withStrategiesRisk {
			externalStrategy.Risk = NewRiskScore().AssignTStrategyFromRisk(strategy.BuildRiskScore())
		}
		newVault.Strategies = append(newVault.Strategies, externalStrategy)
	}

	c.JSON(http.StatusOK, newVault)
}
