package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/harvests"
)

// GetSupportedChains returns a list of supported chains by the API
func GetSupportedChains(c *gin.Context) {
	c.JSON(http.StatusOK, env.SUPPORTED_CHAIN_IDS)
}

// GetHarvests returns a list of harvests
func GetHarvests(c *gin.Context) {
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}
	vaultAddress, ok := helpers.AssertAddress(c.Param("address"), chainID)
	if !ok {
		c.String(http.StatusBadRequest, "invalid address")
		return
	}

	allHarvest := harvests.AllHarvests[vaultAddress]

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
