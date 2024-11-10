package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"

	"github.com/gin-contrib/cache"
	goCache "github.com/patrickmn/go-cache"

	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/external/prices"
	"github.com/yearn/ydaemon/external/strategies"
	"github.com/yearn/ydaemon/external/tokens"
	"github.com/yearn/ydaemon/external/utils"
	"github.com/yearn/ydaemon/external/vaults"
)

var cachingStore *goCache.Cache

func init() {
	cachingStore = goCache.New(1*time.Minute, 5*time.Minute)
}

/**************************************************************************************************
** NewRouter create the routes and setup the server
**************************************************************************************************/
func NewRouter() *gin.Engine {
	gin.EnableJsonDecoderDisallowUnknownFields()
	gin.SetMode(gin.ReleaseMode)
	store := persistence.NewInMemoryStore(10 * time.Minute)
	// gin.DefaultWriter = nil
	router := gin.New()
	// pprof.Register(router)
	router.Use(gin.Recovery())
	corsConf := cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "HEAD", "POST", "OPTIONS"},
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

	router.GET(`/health`, func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "ok", "timestamp": time.Now().Format(time.RFC3339)})
	})

	// Vaults section
	{
		v := vaults.Controller{}
		// Retrieve the vaults for all chains
		// router.GET(`vaults`, c.GetIsYearn)
		router.GET(`vaults/detected`, cache.CachePage(store, time.Minute, func(ctx *gin.Context) {
			if result, err := v.GetAll(ctx); err == nil {
				ctx.JSON(http.StatusOK, result)
			}
		}))
		router.GET(`vaults`, cache.CachePage(store, time.Minute, func(ctx *gin.Context) {
			if result, err := v.GetIsYearn(ctx); err == nil {
				ctx.JSON(http.StatusOK, result)
			}
		}))
		router.GET(`vaults/all`, cache.CachePage(store, time.Minute, func(ctx *gin.Context) {
			if result, err := v.GetIsYearn(ctx); err == nil {
				ctx.JSON(http.StatusOK, result)
			}
		}))
		router.GET(`vaults/underthesea/v2`, cache.CachePage(store, time.Minute, func(ctx *gin.Context) {
			if result, err := v.GetV2(ctx); err == nil {
				ctx.JSON(http.StatusOK, result)
			}
		}))
		router.GET(`vaults/v2`, cache.CachePage(store, time.Minute, func(ctx *gin.Context) {
			if result, err := v.GetV2IsYearn(ctx); err == nil {
				ctx.JSON(http.StatusOK, result)
			}
		}))
		router.GET(`vaults/underthesea/v3`, cache.CachePage(store, time.Minute, func(ctx *gin.Context) {
			if result, err := v.GetV3(ctx); err == nil {
				ctx.JSON(http.StatusOK, result)
			}
		}))
		router.GET(`vaults/v3`, cache.CachePage(store, time.Minute, func(ctx *gin.Context) {
			if result, err := v.GetV3IsYearn(ctx); err == nil {
				ctx.JSON(http.StatusOK, result)
			}
		}))
		router.GET(`vaults/juiced`, cache.CachePage(store, time.Minute, func(ctx *gin.Context) {
			if result, err := v.GetIsYearnJuiced(ctx); err == nil {
				ctx.JSON(http.StatusOK, result)
			}
		}))
		router.GET(`vaults/gimme`, cache.CachePage(store, time.Minute, func(ctx *gin.Context) {
			if result, err := v.GetIsGimme(ctx); err == nil {
				ctx.JSON(http.StatusOK, result)
			}
		}))
		router.GET(`vaults/retired`, cache.CachePage(store, time.Minute, func(ctx *gin.Context) {
			if result, err := v.GetRetired(ctx); err == nil {
				ctx.JSON(http.StatusOK, result)
			}
		}))
		router.GET(`vaults/pendle`, cache.CachePage(store, time.Minute, func(ctx *gin.Context) {
			if result, err := v.GetIsYearnPendle(ctx); err == nil {
				ctx.JSON(http.StatusOK, result)
			}
		}))
		router.GET(`vaults/optimism`, cache.CachePage(store, time.Minute, func(ctx *gin.Context) {
			if result, err := v.GetIsOptimism(ctx); err == nil {
				ctx.JSON(http.StatusOK, result)
			}
		}))
		router.GET(`vaults/pooltogether`, cache.CachePage(store, time.Minute, func(ctx *gin.Context) {
			if result, err := v.GetIsYearnPoolTogether(ctx); err == nil {
				ctx.JSON(http.StatusOK, result)
			}
		}))
		router.GET(`vaults/ajna`, cache.CachePage(store, time.Minute, func(ctx *gin.Context) {
			if result, err := v.GetIsAjna(ctx); err == nil {
				ctx.JSON(http.StatusOK, result)
			}
		}))
		router.GET(`vaults/velodrome`, cache.CachePage(store, time.Minute, func(ctx *gin.Context) {
			if result, err := v.GetIsVelodrome(ctx); err == nil {
				ctx.JSON(http.StatusOK, result)
			}
		}))
		router.GET(`vaults/aerodrome`, cache.CachePage(store, time.Minute, func(ctx *gin.Context) {
			if result, err := v.GetIsAerodrome(ctx); err == nil {
				ctx.JSON(http.StatusOK, result)
			}
		}))
		router.GET(`vaults/curve`, cache.CachePage(store, time.Minute, func(ctx *gin.Context) {
			if result, err := v.GetIsCurve(ctx); err == nil {
				ctx.JSON(http.StatusOK, result)
			}
		}))

		/******************************************************************************************
		** Retrieve some/all vaults based on some specific criteria. This is chain specific and
		** will return the vaults for a specific chain.
		******************************************************************************************/
		router.GET(`:chainID/vaults/all`, CacheLegacyVaults(cachingStore, 10*time.Minute, v.GetLegacyIsYearn))
		router.GET(`:chainID/vaults/v2/all`, CacheLegacyVaults(cachingStore, 10*time.Minute, v.GetLegacyV2IsYearn))
		router.GET(`:chainID/vaults/v3/all`, CacheLegacyVaults(cachingStore, 10*time.Minute, v.GetLegacyV3IsYearn))
		router.GET(`:chainID/vaults/juiced/all`, CacheLegacyVaults(cachingStore, 10*time.Minute, v.GetLegacyIsYearnJuiced))
		router.GET(`:chainID/vaults/gimme/all`, CacheLegacyVaults(cachingStore, 10*time.Minute, v.GetLegacyIsGimme))
		router.GET(`:chainID/vaults/retired`, CacheLegacyVaults(cachingStore, 10*time.Minute, v.GetLegacyRetired))
		router.GET(`:chainID/vaults/some/:addresses`, v.GetLegacySomeVaults)

		/******************************************************************************************
		** Vaults for a custom integration
		******************************************************************************************/
		router.GET(`rotki/list/vaults`, CacheCustomVaults(cachingStore, 10*time.Minute, v.GetVaultsForRotki))
		router.GET(`rotki/count/vaults`, v.CountVaultsForRotki)

		/******************************************************************************************
		** Retrieve a specific vault based on the address. This is chain specific and will return
		** the vault for a specific chain.
		******************************************************************************************/
		router.GET(`:chainID/vaults/:address`, v.GetSimplifiedVault)
		router.GET(`:chainID/vault/:address`, v.GetSimplifiedVault)

		router.GET(`:chainID/vaults/harvests/:addresses`, v.GetHarvestsForVault)
		router.GET(`:chainID/earned/:address/:vaults`, v.GetEarnedPerVaultPerUser)
		router.GET(`:chainID/earned/:address`, v.GetEarnedPerUser)
		router.GET(`earned/:address`, v.GetEarnedPerUserForAllChains)

		// Retrieve the strategies for a specific chainID
		router.GET(`:chainID/strategies/all`, v.GetAllStrategies)
		router.GET(`:chainID/strategies/:address`, v.GetStrategy)
		router.GET(`:chainID/strategy/:address`, v.GetStrategy)

		// Retrieve the TVL
		router.GET(`vaults/tvl`, v.GetAllVaultsTVL)
		router.GET(`:chainID/vaults/tvl`, v.GetVaultsTVL)
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
