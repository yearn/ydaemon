package utils

import (
	"bytes"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal"
)

//GetSupportedChains returns a list of supported chains by the API
func GetSupportedChains(c *gin.Context) {
	c.JSON(http.StatusOK, env.SUPPORTED_CHAIN_IDS)
}

//GetHarvests returns a list of harvests
func GetHarvests(c *gin.Context) {
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

	allHarvest := internal.AllHarvests[address.Address]

	// sumOfAllGains := bigNumber.NewFloat(0)
	// sumOfAllFees := bigNumber.NewFloat(0)
	// for _, v := range allHarvest {
	// 	sumOfAllGains = sumOfAllGains.Add(sumOfAllGains, bigNumber.NewFloat().SetInt(v.Gain))
	// 	sumOfAllFees = sumOfAllFees.Add(sumOfAllFees, bigNumber.NewFloat().SetInt(v.Fees.TotalCollectedFee))
	// 	logs.Info(`Harvest fee ratio: `, bigNumber.NewFloat().Mul(v.Fees.TotalFeeRatio, bigNumber.NewFloat(100)).String()+`%`)
	// }

	// logs.Success(`Total gains: `, sumOfAllGains.String())
	// logs.Success(`Total fees: `, sumOfAllFees.String())
	// logs.Success(`Total fee ratio: `, bigNumber.NewFloat().Mul(bigNumber.NewFloat().Div(sumOfAllFees, sumOfAllGains), bigNumber.NewFloat(100)).String()+`%`)

	c.JSON(http.StatusOK, allHarvest)
}

//GetGraph returns a list of blacklisted vaults by the API
func GetGraph(c *gin.Context) {
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}

	// Read the body as string
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)

	// Proxy the request to the subgraph
	graphQLEndpoint, ok := env.THEGRAPH_ENDPOINTS[chainID]
	if !ok {
		c.String(http.StatusInternalServerError, "Impossible to fetch subgraph")
		return
	}

	client := graphql.NewClient(graphQLEndpoint)
	request := graphql.NewRequest(buf.String())
	var response interface{}
	if err := client.Run(context.Background(), request, &response); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{`error`: `bad request`})
		return
	}
	c.JSON(http.StatusOK, response)
}
