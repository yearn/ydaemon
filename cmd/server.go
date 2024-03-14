package main

import (
	"net/http"
	"sync"

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
	"http://localhost:",
}

// NewRouter create the routes and setup the server
func NewRouter() *gin.Engine {
	gin.EnableJsonDecoderDisallowUnknownFields()
	gin.SetMode(gin.ReleaseMode)
	//disable logs
	gin.DefaultWriter = nil

	router := gin.New()
	pprof.Register(router)
	router.Use(gin.Recovery())
	corsConf := cors.Config{
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
				rl := ratelimit.New(4) // per second
				RateLimitPerOrigin.Store(origin, rl)
				rl.Take()
			} else {
				rl.(ratelimit.Limiter).Take()
			}
			return true
		},
	}
	router.Use(cors.New(corsConf))
	router.Use(gzip.Gzip(gzip.DefaultCompression))
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
