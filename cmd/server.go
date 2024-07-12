package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/external/prices"
	"github.com/yearn/ydaemon/external/strategies"
	"github.com/yearn/ydaemon/external/tokens"
	"github.com/yearn/ydaemon/external/utils"
	"github.com/yearn/ydaemon/external/vaults"
)

/**************************************************************************************************
** NewRouter create the routes and setup the server
**************************************************************************************************/
func NewRouter() *gin.Engine {
	cachingStore := cache.New(1*time.Minute, 5*time.Minute)

	gin.EnableJsonDecoderDisallowUnknownFields()
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = nil
	router := gin.New()
	// pprof.Register(router)
	router.Use(gin.Recovery())
	corsConf := cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "HEAD", "POST"},
		AllowHeaders:    []string{`Origin`, `Content-Length`, `Content-Type`, `Authorization`},
	}
	router.Use(cors.New(corsConf))
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	// router.Use(NewRateLimiter(func(c *gin.Context) {
	// 	c.AbortWithStatus(http.StatusTooManyRequests)
	// }))

	// Standard basic route for hello
	router.GET(`/`, func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to yDaemon"})
	})

	// Vaults section
	{
		c := vaults.Controller{}
		// Retrieve the vaults for all chains
		// router.GET(`vaults`, c.GetIsYearn)
		router.GET(`vaults`, CacheSimplifiedVaults(cachingStore, 10*time.Minute, c.GetIsYearn))
		router.GET(`vaults/all`, CacheSimplifiedVaults(cachingStore, 10*time.Minute, c.GetIsYearn))
		router.GET(`vaults/underthesea/v2`, CacheSimplifiedVaults(cachingStore, 10*time.Minute, c.GetV2))
		router.GET(`vaults/v2`, CacheSimplifiedVaults(cachingStore, 10*time.Minute, c.GetV2IsYearn))
		router.GET(`vaults/underthesea/v3`, CacheSimplifiedVaults(cachingStore, 10*time.Minute, c.GetV3))
		router.GET(`vaults/v3`, CacheSimplifiedVaults(cachingStore, 10*time.Minute, c.GetV3IsYearn))
		router.GET(`vaults/juiced`, CacheSimplifiedVaults(cachingStore, 10*time.Minute, c.GetIsYearnJuiced))
		router.GET(`vaults/gimme`, CacheSimplifiedVaults(cachingStore, 10*time.Minute, c.GetIsGimme))
		router.GET(`vaults/retired`, CacheSimplifiedVaults(cachingStore, 10*time.Minute, c.GetRetired))
		router.GET(`vaults/pendle`, CacheSimplifiedVaults(cachingStore, 10*time.Minute, c.GetIsYearnPendle))
		router.GET(`vaults/optimism`, CacheSimplifiedVaults(cachingStore, 10*time.Minute, c.GetIsOptimism))
		router.GET(`vaults/pooltogether`, CacheSimplifiedVaults(cachingStore, 10*time.Minute, c.GetIsYearnPoolTogether))

		/******************************************************************************************
		** Retrieve some/all vaults based on some specific criteria. This is chain specific and
		** will return the vaults for a specific chain.
		******************************************************************************************/
		router.GET(`:chainID/vaults/all`, CacheLegacyVaults(cachingStore, 10*time.Minute, c.GetLegacyIsYearn))
		router.GET(`:chainID/vaults/v2/all`, CacheLegacyVaults(cachingStore, 10*time.Minute, c.GetLegacyV2IsYearn))
		router.GET(`:chainID/vaults/v3/all`, CacheLegacyVaults(cachingStore, 10*time.Minute, c.GetLegacyV3IsYearn))
		router.GET(`:chainID/vaults/juiced/all`, CacheLegacyVaults(cachingStore, 10*time.Minute, c.GetLegacyIsYearnJuiced))
		router.GET(`:chainID/vaults/gimme/all`, CacheLegacyVaults(cachingStore, 10*time.Minute, c.GetLegacyIsGimme))
		router.GET(`:chainID/vaults/retired`, CacheLegacyVaults(cachingStore, 10*time.Minute, c.GetLegacyRetired))
		router.GET(`:chainID/vaults/some/:addresses`, c.GetLegacySomeVaults)

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

	// Tokens API section
	{
		c := tokens.Controller{}
		router.GET(`tokens/all`, c.GetAllTokens)
		router.GET(`:chainID/tokens/all`, c.GetTokens)
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

	return router
}
