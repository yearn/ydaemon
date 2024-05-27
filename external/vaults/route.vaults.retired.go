package vaults

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/sort"
	"github.com/yearn/ydaemon/internal/storage"
)

// GetRetiredVaults will, for a given chainID, return a list of all retired vaults
func (y Controller) GetRetiredVaults(c *gin.Context) {
	orderBy := helpers.SafeString(getQuery(c, `orderBy`), `details.order`)
	orderDirection := helpers.SafeString(getQuery(c, `orderDirection`), `asc`)
	chainID, ok := helpers.AssertChainID(c.Param(`chainID`))
	if !ok {
		c.String(http.StatusBadRequest, `invalid chainID`)
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
		vaultAddress := currentVault.Address
		if helpers.Contains(env.CHAINS[chainID].BlacklistedVaults, vaultAddress) {
			continue
		}

		newVault, err := NewVault().AssignTVault(currentVault)
		if err != nil {
			continue
		}
		if !newVault.Details.IsRetired {
			continue
		}
		if newVault.Migration.Available {
			continue
		}
		if newVault.TVL.TVL < 1000 {
			continue
		}

		newVault.Strategies = []TStrategy{}
		data = append(data, newVault)
	}

	sort.SortBy(orderBy, orderDirection, data)
	c.JSON(http.StatusOK, data)
}
