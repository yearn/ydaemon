package main

import (
	"net/http"
	"sync"
	"time"

	cache "github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
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
	"go.uber.org/ratelimit"
)

// RateLimitPerOrigin is a sync map if url -> ratelimit
var RateLimitPerOrigin = sync.Map{}

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
}

// NewRouter create the routes and setup the server
func NewRouter() *gin.Engine {
	gin.EnableJsonDecoderDisallowUnknownFields()
	gin.SetMode(gin.ReleaseMode)
	//disable logs
	gin.DefaultWriter = nil
	memoryStore := persist.NewMemoryStore(1 * time.Minute)

	router := gin.New()
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(gin.Recovery())
	router.Use(sentrygin.New(sentrygin.Options{Repanic: true}))
	router.Use(func(ctx *gin.Context) {
		if hub := sentrygin.GetHubFromContext(ctx); hub != nil {
			hub.Scope().SetTag("subsystem", "gin")
		}
		ctx.Next()
	})
	corsConf := cors.Config{
		// AllowAllOrigins: true,
		AllowMethods: []string{"GET", "HEAD"},
		AllowHeaders: []string{`Origin`, `Content-Length`, `Content-Type`, `Authorization`},
		AllowOriginFunc: func(origin string) bool {
			//Check if we have this origin in the RateLimitPerOrigin map
			//If we don't have it, we add it
			if helpers.Contains(allowlist, origin) {
				return true
			}
			//Allow our route URI
			if helpers.EndsWithSubstring(rootURI, origin) {
				return true
			}

			if rl, ok := RateLimitPerOrigin.Load(origin); !ok {
				rl := ratelimit.New(1) // per second
				RateLimitPerOrigin.Store(origin, rl)
				rl.Take()
			} else {
				rl.(ratelimit.Limiter).Take()
			}
			return true
		},
	}
	router.Use(cors.New(corsConf))
	router.GET(`/`, func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to yDaemon"})
	})

	// Vaults section
	{
		c := vaults.Controller{}
		// Retrieve the vaults for all chains
		router.GET(`vaults`, cache.CacheByRequestURI(memoryStore, 5*time.Minute), c.GetAllVaultsForAllChainsSimplified)                // Migrated to simplified ✅
		router.GET(`vaults/retired`, cache.CacheByRequestURI(memoryStore, 5*time.Minute), c.GetAllRetiredVaultsForAllChainsSimplified) // Migrated to simplified ✅
		router.GET(`vaults/all`, cache.CacheByRequestURI(memoryStore, 5*time.Minute), c.GetAllVaultsForAllChainsSimplified)            // Migrated to simplified ✅
		router.GET(`vaults/v3`, cache.CacheByRequestURI(memoryStore, 5*time.Minute), c.GetAllV3VaultsForAllChainsSimplified)           // Migrated to simplified ✅
		router.GET(`vaults/v2`, cache.CacheByRequestURI(memoryStore, 5*time.Minute), c.GetAllV2VaultsForAllChainsSimplified)           // Migrated to simplified ✅

		router.GET(`vaults/tvl`, cache.CacheByRequestURI(memoryStore, 5*time.Minute), c.GetAllVaultsTVL)

		// Retrieve the vaults for a specific chainID
		router.GET(`:chainID/vaults/tvl`, cache.CacheByRequestURI(memoryStore, 5*time.Minute), c.GetVaultsTVL)
		router.GET(`:chainID/vaults/all`, cache.CacheByRequestURI(memoryStore, 5*time.Minute), c.GetAllVaults)
		router.GET(`:chainID/vaults/retired`, cache.CacheByRequestURI(memoryStore, 5*time.Minute), c.GetRetiredVaults)
		router.GET(`:chainID/vaults/:address`, cache.CacheByRequestURI(memoryStore, 5*time.Minute), c.GetSimplifiedVault) // Migrated to simplified ✅
		router.GET(`:chainID/vault/:address`, cache.CacheByRequestURI(memoryStore, 5*time.Minute), c.GetSimplifiedVault)  // Migrated to simplified ✅

		router.GET(`info/vaults/blacklisted`, cache.CacheByRequestURI(memoryStore, 5*time.Minute), c.GetBlacklistedVaults)

		router.GET(`:chainID/vaults/harvests/:addresses`, cache.CacheByRequestURI(memoryStore, 5*time.Minute), c.GetHarvestsForVault)
		router.GET(`:chainID/earned/:address/:vaults`, cache.CacheByRequestURI(memoryStore, 5*time.Minute), c.GetEarnedPerVaultPerUser)
		router.GET(`:chainID/earned/:address`, cache.CacheByRequestURI(memoryStore, 5*time.Minute), c.GetEarnedPerUser)
		router.GET(`earned/:address`, cache.CacheByRequestURI(memoryStore, 5*time.Minute), c.GetEarnedPerUserForAllChains)

		// Retrieve the strategies for a specific chainID
		router.GET(`:chainID/strategies/all`, cache.CacheByRequestURI(memoryStore, 5*time.Minute), c.GetAllStrategies)
		router.GET(`:chainID/strategies/:address`, cache.CacheByRequestURI(memoryStore, 5*time.Minute), c.GetStrategy)
		router.GET(`:chainID/strategy/:address`, cache.CacheByRequestURI(memoryStore, 5*time.Minute), c.GetStrategy)
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
		router.GET(`prices/all`, cache.CacheByRequestURI(memoryStore, 5*time.Minute), c.GetAllPrices)
		router.GET(`:chainID/prices/all`, cache.CacheByRequestURI(memoryStore, 5*time.Minute), c.GetPrices)
		router.GET(`:chainID/prices/:address`, cache.CacheByRequestURI(memoryStore, 5*time.Minute), c.GetPrice)
		router.GET(`:chainID/prices/some/:addresses`, cache.CacheByRequestURI(memoryStore, 5*time.Minute), c.GetSomePrices)
		router.GET(`:chainID/prices/all/details`, cache.CacheByRequestURI(memoryStore, 5*time.Minute), c.GetAllPricesWithDetails)
	}

	// WARNING: DEPRECATED
	// yBribe API section
	{
		router.GET(`:chainID/bribes/newRewardFeed`, utils.GetRewardAdded)
	}

	return router
}
