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

// GetAllRetiredVaultsForAllChainsSimplified will, for a given chainID, return a list of all retired vaults
func (y Controller) GetAllRetiredVaultsForAllChainsSimplified(c *gin.Context) {
	orderBy := helpers.SafeString(getQuery(c, `orderBy`), `details.order`)
	orderDirection := helpers.SafeString(getQuery(c, `orderDirection`), `asc`)

	/** 🔵 - Yearn *************************************************************************************
	** chainsStr: A string that represents the chain IDs for which the vaults are to be returned. It is
	** obtained from the 'chainIDs' query parameter in the request. The string is split by commas to
	** obtain an array of chain IDs.
	**
	** chains: An array of uint64 values that represents the chain IDs for which the vaults are to be
	** returned. If the 'chains' query parameter is not provided or is empty, this array defaults to
	** all supported chain IDs.
	**
	** The 'chains' array is populated by iterating over the 'chainsStr' array and converting each
	** chain ID to a uint64 value. If the conversion is not successful, the chain ID is ignored.
	**
	** The 'chains' array is used to filter the vaults that are returned in the response.
	**************************************************************************************************/
	chainsStr := strings.Split(getQuery(c, `chainIDs`), `,`)
	chains := []uint64{}
	if len(chainsStr) == 0 || (len(chainsStr) == 1 && chainsStr[0] == ``) {
		chains = env.SUPPORTED_CHAIN_IDS
	} else {
		for _, chainStr := range chainsStr {
			chain, ok := helpers.AssertChainID(chainStr)
			if !ok {
				continue
			}
			chains = append(chains, chain)
		}
	}

	allVaults := []TSimplifiedExternalVault{}
	for _, chainID := range chains {
		vaultsForChain, _ := storage.ListVaults(chainID)
		for _, currentVault := range vaultsForChain {
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
			allVaults = append(allVaults, toSimplifiedVersion(newVault))
		}
	}

	sort.SortBy(orderBy, orderDirection, allVaults)
	c.JSON(http.StatusOK, allVaults)
}
