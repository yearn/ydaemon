package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/external/meta"
	"github.com/yearn/ydaemon/external/partners"
	"github.com/yearn/ydaemon/external/prices"
	"github.com/yearn/ydaemon/external/strategies"
	"github.com/yearn/ydaemon/external/tokens"
	"github.com/yearn/ydaemon/external/tokensList"
	"github.com/yearn/ydaemon/external/utils"
	"github.com/yearn/ydaemon/external/vaults"
	"golang.org/x/time/rate"
)

var allowlist = []string{
	"https://yearn.finance",
	"https://yearn.fi",
	"https://ycrv.yearn.fi",
	"https://yeth.yearn.fi",
	"https://veyfi.yearn.fi",
	"https://juiced.yearn.fi",
	"https://buyback.yearn.fi",
	"https://seafood.yearn.watch",
	"https://juiced.app",
	"https://ajnafi.com",
	"https://tokenlistooor.com",
	"https://ybribe.com",
	"https://yearn.farm",
	"https://ape.tax",
}
var rootURI = []string{
	".yearn.fi",
	".yearn.finance",
	".yearn.farm",
	".juiced.app",
	".smold.app",
	"http://localhost:",
}

/**************************************************************************************************
** Rate limiting based on https://github.com/yangxikun/gin-limit-by-key/blob/master/limit.go and
** adapted to our needs.
**************************************************************************************************/
var limiterSet = cache.New(5*time.Minute, 10*time.Minute)

func NewRateLimiter(abort func(*gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		/******************************************************************************************
		** Retrieve the origin from the request header and use it as the key for the rate limiter.
		** If the origin is not present, we use an empty string as the key.
		******************************************************************************************/
		k := ``
		origin := c.Request.Header.Get("Origin")
		if len(origin) == 0 {
			k = ``
		} else {
			k = origin
		}

		/******************************************************************************************
		** Allows the requests from the allowlist without rate limiting. This is to allow us to
		** bypass the rate limiting for our own services.
		******************************************************************************************/
		if helpers.Contains(allowlist, origin) || helpers.EndsWithSubstring(rootURI, origin) || origin == `` {
			c.Next()
			return
		}

		/******************************************************************************************
		** Otherwise, we use the rate limiter to limit the requests to 10 qps/clientIp and permit
		******************************************************************************************/
		limiter, ok := limiterSet.Get(k)
		if !ok {
			var expire time.Duration
			// limit 25 query per second per origin and permit bursts of at most 25 tokens, and the limiter liveness time duration is 15 minutes
			limiter, expire = rate.NewLimiter(rate.Every(1*time.Second), 50), 15*time.Minute
			limiterSet.Set(k, limiter, expire)
		}
		ok = limiter.(*rate.Limiter).Allow()
		if !ok {
			abort(c)
			return
		}
		c.Next()
	}
}

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
		router.GET(`vaults`, c.GetAllVaultsForAllChainsSimplified)                // Migrated to simplified ✅
		router.GET(`vaults/retired`, c.GetAllRetiredVaultsForAllChainsSimplified) // Migrated to simplified ✅
		router.GET(`vaults/all`, c.GetAllVaultsForAllChainsSimplified)            // Migrated to simplified ✅
		router.GET(`vaults/v3`, c.GetAllV3VaultsForAllChainsSimplified)           // Migrated to simplified ✅
		router.GET(`vaults/v2`, c.GetAllV2VaultsForAllChainsSimplified)           // Migrated to simplified ✅

		router.GET(`vaults/tvl`, c.GetAllVaultsTVL)

		// Retrieve the vaults for a specific chainID
		router.GET(`:chainID/vaults/tvl`, c.GetVaultsTVL)
		router.GET(`:chainID/vaults/all`, c.GetAllVaults)
		router.GET(`:chainID/vaults/retired`, c.GetRetiredVaults)
		router.GET(`:chainID/vaults/:address`, c.GetSimplifiedVault) // Migrated to simplified ✅
		router.GET(`:chainID/vault/:address`, c.GetSimplifiedVault)  // Migrated to simplified ✅

		router.GET(`info/vaults/blacklisted`, c.GetBlacklistedVaults)

		router.GET(`:chainID/vaults/harvests/:addresses`, c.GetHarvestsForVault)
		router.GET(`:chainID/earned/:address/:vaults`, c.GetEarnedPerVaultPerUser)
		router.GET(`:chainID/earned/:address`, c.GetEarnedPerUser)
		router.GET(`earned/:address`, c.GetEarnedPerUserForAllChains)

		// Retrieve the strategies for a specific chainID
		router.GET(`:chainID/strategies/all`, c.GetAllStrategies)
		router.GET(`:chainID/strategies/:address`, c.GetStrategy)
		router.GET(`:chainID/strategy/:address`, c.GetStrategy)
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
		router.GET(`:chainID/prices/some/:addresses`, c.GetSomePrices)
		router.GET(`:chainID/prices/all/details`, c.GetAllPricesWithDetails)
	}

	// WARNING: DEPRECATED
	// yBribe API section
	{
		router.GET(`:chainID/bribes/newRewardFeed`, utils.GetRewardAdded)
	}

	return router
}
