package main

import (
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/traces"
	"github.com/yearn/ydaemon/external/meta"
	"github.com/yearn/ydaemon/external/partners"
	"github.com/yearn/ydaemon/external/prices"
	"github.com/yearn/ydaemon/external/strategies"
	"github.com/yearn/ydaemon/external/tokens"
	"github.com/yearn/ydaemon/external/utils"
	"github.com/yearn/ydaemon/external/vaults"
)

func middleware(log *logrus.Logger) gin.HandlerFunc {
	LEVEL, exists := os.LookupEnv("LOG_LEVEL")
	if exists {
		levels := map[string]logrus.Level{
			"DEBUG":   logrus.DebugLevel,
			"INFO":    logrus.InfoLevel,
			"WARNING": logrus.WarnLevel,
			"SUCCESS": logrus.WarnLevel,
			"ERROR":   logrus.ErrorLevel,
		}
		log.SetLevel(levels[LEVEL])
	}

	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.RequestURI, "/api/tokens/list") {
			c.Header("Cache-Control", "max-age=86400")
		} else {
			log.WithFields(logrus.Fields{
				"path":   c.Request.RequestURI,
				"method": c.Request.Method,
				"status": c.Writer.Status(),
			}).Info(time.Now().Format(time.RFC3339))
		}
	}
}

func GET(router *gin.Engine, path string, handler gin.HandlerFunc) {
	router.GET(path, func(c *gin.Context) {
		trace := traces.Init(`http.get`).
			SetTag(`subsystem`, `gin`).
			SetTag(path, path)

		if chainID, ok := helpers.AssertChainID(c.Param("chainID")); ok {
			trace.SetTag(`chainID`, strconv.Itoa(int(chainID)))
		}
		defer trace.Finish()

		handler(c)
	})
}

// NewRouter create the routes and setup the server
func NewRouter() *gin.Engine {
	gin.EnableJsonDecoderDisallowUnknownFields()
	gin.SetMode(gin.ReleaseMode)

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
	router.Use(middleware(logrus.New()))
	corsConf := cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "HEAD"},
		AllowHeaders:    []string{`Origin`, `Content-Length`, `Content-Type`, `Authorization`},
	}
	router.Use(cors.New(corsConf))

	// Vaults section
	{
		c := vaults.Controller{}
		// Retrieve the vaults for all chains
		GET(router, `vaults/tvl`, c.GetAllVaultsTVL)

		// Retrieve the vaults for a specific chainID
		GET(router, `:chainID/vaults/tvl`, c.GetVaultsTVL)
		GET(router, `:chainID/vaults/all`, c.GetAllVaults)
		GET(router, `:chainID/vaults/:address`, c.GetVault)
		GET(router, `:chainID/vault/:address`, c.GetVault)

		GET(router, `info/vaults/blacklisted`, c.GetBlacklistedVaults)

		GET(router, `:chainID/vaults/harvests/:addresses`, c.GetHarvestsForVault)
		GET(router, `:chainID/vaults/apy/:address`, c.GetVaultsVisionAPY)
		GET(router, `:chainID/earned/:address/:vaults`, c.GetEarnedPerVaultPerUser)
		GET(router, `:chainID/earned/:address`, c.GetEarnedPerUser)
	}

	// Strategies section
	{
		c := strategies.Controller{}
		// Retrieve the reports for a specific strategy
		GET(router, `:chainID/reports/:address`, c.GetReports)
	}

	// General section
	{
		// Get some information about the API
		GET(router, `info/chains`, utils.GetSupportedChains)

		// Proxy subgraph
		router.POST(`:chainID/graph`, utils.GetGraph)
	}

	// Meta API section
	{
		c := meta.Controller{}
		// Proxy meta strategies
		GET(router, `api/:chainID/strategies/all`, c.GetMetaStrategies)
		GET(router, `:chainID/meta/strategies`, c.GetMetaStrategies)
		GET(router, `api/:chainID/strategies/:address`, c.GetMetaStrategy)
		GET(router, `:chainID/meta/strategies/:address`, c.GetMetaStrategy)
		GET(router, `:chainID/meta/strategy/:address`, c.GetMetaStrategy)

		// Proxy meta tokens
		GET(router, `api/:chainID/tokens/all`, c.GetMetaTokens)
		GET(router, `:chainID/meta/tokens`, c.GetMetaTokens)
		GET(router, `api/:chainID/tokens/:address`, c.GetMetaToken)
		GET(router, `:chainID/meta/tokens/:address`, c.GetMetaToken)
		GET(router, `:chainID/meta/token/:address`, c.GetMetaToken)

		// Proxy meta vaults
		GET(router, `api/:chainID/vaults/all`, c.GetMetaVaultsLegacy)
		GET(router, `:chainID/meta/vaults`, c.GetMetaVaults)
		GET(router, `api/:chainID/vaults/:address`, c.GetMetaVault)
		GET(router, `:chainID/meta/vaults/:address`, c.GetMetaVault)
		GET(router, `:chainID/meta/vault/:address`, c.GetMetaVault)

		// Proxy meta protocols
		GET(router, `api/:chainID/protocols/all`, c.GetMetaProtocols)
		GET(router, `:chainID/meta/protocols`, c.GetMetaProtocols)
		GET(router, `api/:chainID/protocols/:name`, c.GetMetaProtocol)
		GET(router, `:chainID/meta/protocols/:name`, c.GetMetaProtocol)
		GET(router, `:chainID/meta/protocol/:name`, c.GetMetaProtocol)
	}

	// Partners API section
	{
		c := partners.Controller{}
		GET(router, `partners/count`, c.CountAllPartners)
		GET(router, `partners/all`, c.GetAllPartners)
		GET(router, `:chainID/partners/all`, c.GetPartners)
		GET(router, `:chainID/partners/:addressOrName`, c.GetPartner)
	}

	// Tokens API section
	{
		c := tokens.Controller{}
		GET(router, `tokens/all`, c.GetAllTokens)
		GET(router, `:chainID/tokens/all`, c.GetTokens)
	}

	// Prices API section
	{
		c := prices.Controller{}
		GET(router, `prices/all`, c.GetAllPrices)
		GET(router, `:chainID/prices/all`, c.GetPrices)
		GET(router, `:chainID/prices/:address`, c.GetPrice)
		GET(router, `:chainID/prices/some/:addresses`, c.GetSomePrices)
	}

	// yBribe API section
	{
		GET(router, `:chainID/bribes/newRewardFeed`, utils.GetRewardAdded)
	}

	{
		router.StaticFile("api/tokens/list", env.BASE_DATA_PATH+`/meta/tokens/tokenList.json`)
	}

	{
		//TEST
		GET(router, `core/harvests/:chainID/:address`, utils.GetHarvests)
	}

	// Make sure Sentry is working
	{
		router.GET("sentry-message-test", func(ctx *gin.Context) {
			if hub := sentrygin.GetHubFromContext(ctx); hub != nil {
				hub.WithScope(func(scope *sentry.Scope) {
					hub.CaptureMessage("sentry-message-test")
				})
			}
			ctx.Status(http.StatusOK)
		})

		router.GET("/sentry-panic-test", func(ctx *gin.Context) {
			panic("sentry-panic-test")
		})
	}

	return router
}
