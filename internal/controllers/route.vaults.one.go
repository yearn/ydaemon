package controllers

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
	"github.com/majorfi/ydaemon/internal/ethereum"
	"github.com/majorfi/ydaemon/internal/logs"
	"github.com/majorfi/ydaemon/internal/models"
	"github.com/majorfi/ydaemon/internal/utils"
)

func graphQLRequestForOneVault(vaultAddress string, c *gin.Context) *graphql.Request {
	withDetails := valueWithFallback(c.Query("strategiesDetails"), "noDetails") == "withDetails"

	return graphql.NewRequest(`{
		vault(id: "` + strings.ToLower(vaultAddress) + `") {
			` + utils.GetGraphRequestVault() + `
			` + utils.GetGraphRequestStrategies(40, withDetails) + `
		}
	}`)
}

//GetVault will, for a given chainID, return a list of all vaults
func (y controller) GetVault(c *gin.Context) {
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
		c.String(http.StatusBadRequest, "invalid address")
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
