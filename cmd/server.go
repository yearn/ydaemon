package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/external/meta"
	"github.com/yearn/ydaemon/external/partners"
	"github.com/yearn/ydaemon/external/prices"
	"github.com/yearn/ydaemon/external/strategies"
	"github.com/yearn/ydaemon/external/tokens"
	"github.com/yearn/ydaemon/external/tokensList"
	"github.com/yearn/ydaemon/external/utils"
	"github.com/yearn/ydaemon/external/vaults"
)

/**************************************************************************************************
** NewRouter create the routes and setup the server
**************************************************************************************************/
func NewRouter() *gin.Engine {
	gin.EnableJsonDecoderDisallowUnknownFields()
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = nil
	router := gin.New()
	pprof.Register(router)
	router.Use(gin.Recovery())
	corsConf := cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "HEAD"},
		AllowHeaders:    []string{`Origin`, `Content-Length`, `Content-Type`, `Authorization`},
	}
	router.Use(cors.New(corsConf))
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(NewRateLimiter(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusTooManyRequests)
	}))

	// Standard basic route for hello
	router.GET(`/`, func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to yDaemon"})
	})

	// Vaults section
	{
		c := vaults.Controller{}
		// Retrieve the vaults for all chains
		router.GET(`vaults`, c.GetAllVaultsForAllChainsSimplified)

		/******************************************************************************************
		** Retrieve some/all vaults based on some specific criteria. This is chain agnostic and
		** will return the vaults for all chains.
		******************************************************************************************/
		router.GET(`vaults/all`, c.GetAllVaultsForAllChainsSimplified)
		router.GET(`vaults/v3`, c.GetAllV3VaultsForAllChainsSimplified)
		router.GET(`vaults/v2`, c.GetAllV2VaultsForAllChainsSimplified)
		router.GET(`vaults/juiced`, c.GetAllJuicedVaultsForAllChainsSimplified)
		router.GET(`vaults/gimme`, c.GetAllGimmeVaultsForAllChainsSimplified)
		router.GET(`vaults/retired`, c.GetAllRetiredVaultsForAllChainsSimplified)

		/******************************************************************************************
		** Retrieve all vaults for seasolver
		******************************************************************************************/
		router.GET(`vaults/underthesea/v2`, c.GetAllV2VaultsForAllChainsSeaSimplified)
		router.GET(`vaults/underthesea/v3`, c.GetAllV3VaultsForAllChainsSeaSimplified)

		/******************************************************************************************
		** Retrieve some/all vaults based on some specific criteria. This is chain specific and
		** will return the vaults for a specific chain.
		******************************************************************************************/
		router.GET(`:chainID/vaults/all`, c.GetAllVaults)
		router.GET(`:chainID/vaults/v2/all`, c.GetAllV2Vaults)
		router.GET(`:chainID/vaults/v3/all`, c.GetAllV3Vaults)
		router.GET(`:chainID/vaults/juiced/all`, c.GetAllJuicedVaults)
		router.GET(`:chainID/vaults/gimme/all`, c.GetAllGimmeVaults)
		router.GET(`:chainID/vaults/retired`, c.GetRetiredVaults)
		router.GET(`:chainID/vaults/some/:addresses`, c.GetSomeVaults)

		/******************************************************************************************
		** Retrieve a specific vault based on the address. This is chain specific and will return
		** the vault for a specific chain.
		******************************************************************************************/
		router.GET(`:chainID/vaults/:address`, c.GetSimplifiedVault)
		router.GET(`:chainID/vault/:address`, c.GetSimplifiedVault)

		router.GET(`:chainID/vaults/harvests/:addresses`, c.GetHarvestsForVault)
		router.GET(`:chainID/earned/:address/:vaults`, c.GetEarnedPerVaultPerUser)
		router.GET(`:chainID/earned/:address`, c.GetEarnedPerUser)
		router.GET(`earned/:address`, c.GetEarnedPerUserForAllChains)

		// Retrieve the strategies for a specific chainID
		router.GET(`:chainID/strategies/all`, c.GetAllStrategies)
		router.GET(`:chainID/strategies/:address`, c.GetStrategy)
		router.GET(`:chainID/strategy/:address`, c.GetStrategy)

		// Retrieve the TVL
		router.GET(`vaults/tvl`, c.GetAllVaultsTVL)
		router.GET(`:chainID/vaults/tvl`, c.GetVaultsTVL)
	}

	// Strategies section
	{
		c := strategies.Controller{}
		// Retrieve the reports for a specific strategy
		router.GET(`:chainID/reports/:address`, c.GetReports)
	}

	// General section
	{
		// Get some information about the API
		vController := vaults.Controller{}
		router.GET(`info/vaults/blacklisted`, vController.GetBlacklistedVaults)
		router.GET(`info/chains`, utils.GetSupportedChains)
		router.GET(`:chainID/status`, func(ctx *gin.Context) {
			chainID, ok := helpers.AssertChainID(ctx.Param("chainID"))
			if !ok {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid chainID"})
				return
			}
			ctx.JSON(http.StatusOK, getStatusForChainID(chainID))
		})
	}

	// Meta API section
	{
		c := meta.Controller{}
		// Proxy meta protocols
		router.GET(`api/:chainID/protocols/all`, c.GetMetaProtocols)
		router.GET(`:chainID/meta/protocols`, c.GetMetaProtocols)
		router.GET(`api/:chainID/protocols/:name`, c.GetMetaProtocol)
		router.GET(`:chainID/meta/protocols/:name`, c.GetMetaProtocol)
		router.GET(`:chainID/meta/protocol/:name`, c.GetMetaProtocol)
	}

	// Partners API section
	{
		c := partners.Controller{}
		router.GET(`partners/count`, c.CountAllPartners)
		router.GET(`partners/all`, c.GetAllPartners)
		router.GET(`:chainID/partners/all`, c.GetPartners)
		router.GET(`:chainID/partners/:addressOrName`, c.GetPartner)
	}

	// Tokens API section
	{
		c := tokens.Controller{}
		router.GET(`tokens/all`, c.GetAllTokens)
		router.GET(`:chainID/tokens/all`, c.GetTokens)
		router.GET(`:chainID/tokenlistbalances/:address`, tokensList.GetYearnTokenList)
		router.GET(`:chainID/balances/:address`, tokensList.GetYearnTokenList)
		router.GET(`balances/:address`, tokensList.GetUserBalance)
		router.GET(`balancesN/:address`, tokensList.GetNewUserBalance)
	}

	// Prices API section
	{
		c := prices.Controller{}
		router.GET(`prices/all`, c.GetAllPrices)
		router.GET(`:chainID/prices/all`, c.GetPrices)
		router.GET(`:chainID/prices/:address`, c.GetPrice)
		router.GET(`:chainID/prices/some/:addresses`, c.GetSomePricesForChain)
		router.GET(`:chainID/prices/all/details`, c.GetAllPricesWithDetails)

		/******************************************************************************************
		** Retrieve some/all prices based on some specific criteria. This is chain agnostic and
		** will return the prices for all chains.
		******************************************************************************************/
		router.GET(`prices/some/:addresses`, c.GetSomePrices)
		router.POST(`prices/some`, c.GetSomePostPrices)

	}

	// WARNING: DEPRECATED
	// yBribe API section
	{
		router.GET(`:chainID/bribes/newRewardFeed`, utils.GetRewardAdded)
	}

	return router
}
