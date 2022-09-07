package vaultsController

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
	"github.com/yearn/ydaemon/internal/ethereum"
	"github.com/yearn/ydaemon/internal/logs"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/store"
	"github.com/yearn/ydaemon/internal/utils"
)

func graphQLRequestForOneVault(vaultAddress string, c *gin.Context) *graphql.Request {
	withDetails := utils.ValueWithFallback(c.Query("strategiesDetails"), "noDetails") == "withDetails"

	return graphql.NewRequest(`{
		vault(id: "` + strings.ToLower(vaultAddress) + `") {
			` + utils.GetGraphRequestVault() + `
			` + utils.GetGraphRequestStrategies(40, withDetails) + `
		}
	}`)
}

//GetVault will, for a given chainID, return a list of all vaults
func (y Controller) GetVault(c *gin.Context) {
	chainID, err := strconv.ParseUint(c.Param("chainID"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}
	vaultAddress := c.Param("address")
	if vaultAddress == `` {
		c.String(http.StatusBadRequest, "invalid address")
		return
	}

	if utils.ContainsAddress(utils.BLACKLISTED_VAULTS[chainID], common.HexToAddress(vaultAddress)) {
		c.String(http.StatusBadRequest, "ignored address")
		return
	}
	vaultFromMeta, ok := store.VaultsFromMeta[chainID][common.HexToAddress(vaultAddress)]
	if ok && vaultFromMeta.HideAlways {
		c.String(http.StatusBadRequest, "ignored address")
		return
	}

	client := graphql.NewClient(ethereum.GetGraphURI(chainID))
	request := graphQLRequestForOneVault(vaultAddress, c)
	var response models.TGraphQueryResponseForVault
	if err := client.Run(context.Background(), request, &response); err != nil {
		logs.Error(err)
		c.String(http.StatusInternalServerError, "Impossible to fetch subgraph")
		return
	}

	strategiesCondition := selectStrategiesCondition(c.Query("strategiesCondition"))
	withStrategiesDetails := c.Query("strategiesDetails") == "withDetails"
	withStrategiesRisk := c.Query("strategiesRisk") == "withRisk"
	vaultFromGraph := response.Vault

	data := prepareVaultSchema(
		chainID,
		strategiesCondition,
		withStrategiesRisk,
		withStrategiesDetails,
		vaultFromGraph,
	)
	c.JSON(http.StatusOK, data)
}
