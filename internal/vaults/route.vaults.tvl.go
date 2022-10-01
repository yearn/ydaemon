package vaults

import (
	"context"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
	"github.com/yearn/ydaemon/internal/meta"
	"github.com/yearn/ydaemon/internal/utils/env"
	"github.com/yearn/ydaemon/internal/utils/helpers"
	"github.com/yearn/ydaemon/internal/utils/logs"
	"github.com/yearn/ydaemon/internal/utils/models"
)

func computeChainTVL(chainID uint64, c *gin.Context) float64 {
	graphQLEndpoint, ok := env.GRAPH_URI[chainID]
	if !ok {
		logs.Error("No graph endpoint for chainID", chainID)
		return 0.0
	}

	client := graphql.NewClient(graphQLEndpoint)
	request := graphQLRequestForAllVaults(c)
	var response models.TGraphQueryResponseForVaults
	if err := client.Run(context.Background(), request, &response); err != nil {
		logs.Error(err)
		return 0.0
	}

	tvl := 0.0
	for _, vaultFromGraph := range response.Vaults {
		vaultAddress := vaultFromGraph.Id
		if helpers.ContainsAddress(env.BLACKLISTED_VAULTS[chainID], vaultAddress) {
			continue
		}
		vaultFromMeta, ok := meta.Store.VaultsFromMeta[chainID][vaultAddress]
		if ok && vaultFromMeta.HideAlways {
			continue
		}
		vaultTVL := prepareTVL(chainID, vaultFromGraph)
		tvl += vaultTVL
	}
	return tvl
}

//GetAllVaultsTVL will, for a all supported chains, return the current TVL
func (y Controller) GetAllVaultsTVL(c *gin.Context) {
	total := 0.0
	var wg sync.WaitGroup
	var tvl = make(map[uint64]float64)
	for _, chainID := range env.SUPPORTED_CHAIN_IDS {
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

//GetVaultsTVL will, for a given chainID, return the current TVL
func (y Controller) GetVaultsTVL(c *gin.Context) {
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}

	c.JSON(http.StatusOK, computeChainTVL(chainID, c))
}
