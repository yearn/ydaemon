package controllers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
	"github.com/majorfi/ydaemon/internal/ethereum"
	"github.com/majorfi/ydaemon/internal/logs"
	"github.com/majorfi/ydaemon/internal/models"
	"github.com/majorfi/ydaemon/internal/store"
	"github.com/majorfi/ydaemon/internal/utils"
)

//GetAllVaultsTVL will, for a given chainID, return the current TVL
func (y controller) GetAllVaultsTVL(c *gin.Context) {
	chainID, err := strconv.ParseUint(c.Param("chainID"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}

	client := graphql.NewClient(ethereum.GetGraphURI(chainID))
	request := graphQLRequestForAllVaults(c)
	var response models.TGraphQueryResponseForVaults
	if err := client.Run(context.Background(), request, &response); err != nil {
		logs.Error(err)
		c.String(http.StatusBadRequest, "invalid graphQL response")
		return
	}

	tvl := 0.0
	for _, vaultFromGraph := range response.Vaults {
		vaultAddress := common.HexToAddress(vaultFromGraph.Id)
		if utils.ContainsAddress(utils.BLACKLISTED_VAULTS[chainID], vaultAddress) {
			continue
		}
		vaultFromMeta, ok := store.VaultsFromMeta[chainID][vaultAddress]
		if ok && vaultFromMeta.HideAlways {
			continue
		}
		strategiesCondition := selectStrategiesCondition(c.Query("strategiesCondition"))
		withStrategiesDetails := c.Query("strategiesDetails") == "withDetails"
		withStrategiesRisk := c.Query("strategiesRisk") == "withRisk"
		currentVault := prepareVaultSchema(
			chainID,
			strategiesCondition,
			withStrategiesRisk,
			withStrategiesDetails,
			vaultFromGraph,
		)
		tvl += currentVault.TVL.TVL
	}

	c.JSON(http.StatusOK, tvl)
}
