package vaults

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/fetcher"
	"github.com/yearn/ydaemon/internal/storage"
)

func computeChainTVL(chainID uint64, c *gin.Context) float64 {
	tvl := 0.0
	vaultsList, _ := storage.ListVaults(chainID)
	for _, currentVault := range vaultsList {
		/******************************************************************************************
		** We want to ignore all non Yearn vaults
		******************************************************************************************/
		if !currentVault.Metadata.Inclusion.IsYearn {
			continue
		}
		chain, ok := env.GetChain(chainID)
		if !ok {
			continue
		}
		if helpers.Contains(chain.BlacklistedVaults, currentVault.Address) {
			continue
		}
		vaultTVL := fetcher.BuildVaultTVL(currentVault)
		tvl += vaultTVL.TVL
	}
	return tvl
}

// GetAllVaultsTVL will, for a all supported chains, return the current TVL
func (y Controller) GetAllVaultsTVL(c *gin.Context) {
	total := 0.0
	var wg sync.WaitGroup
	var tvl = make(map[uint64]float64)
	for chainID := range env.GetChains() {
		wg.Add(1)
		go func(chainID uint64) {
			defer wg.Done()
			tvl[chainID] = computeChainTVL(chainID, c)
			total += tvl[chainID]
		}(uint64(chainID))
	}
	wg.Wait()
	c.JSON(http.StatusOK, gin.H{
		"total":  total,
		"chains": tvl,
	})
}

// GetVaultsTVL will, for a given chainID, return the current TVL
func (y Controller) GetVaultsTVL(c *gin.Context) {
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}
	c.JSON(http.StatusOK, computeChainTVL(chainID, c))
}
