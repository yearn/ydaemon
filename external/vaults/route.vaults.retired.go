package vaults

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/sort"
	"github.com/yearn/ydaemon/internal/vaults"
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
	allVaults := vaults.ListVaults(chainID)
	for _, currentVault := range allVaults {
		vaultAddress := currentVault.Address
		if helpers.Contains(env.BLACKLISTED_VAULTS[chainID], vaultAddress) {
			continue
		}

		newVault := NewVault().AssignTVault(currentVault)
		if !newVault.Details.Retired {
			continue
		}
		if newVault.Migration.Available {
			continue
		}
		if newVault.TVL.TVL < 1000 {
			continue
		}

		newVault.Strategies = []*TStrategy{}
		data = append(data, *newVault)
	}

	sort.SortBy(orderBy, orderDirection, data)
	c.JSON(http.StatusOK, data)
}
