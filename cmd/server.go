package main

import (
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/external/meta"
	"github.com/yearn/ydaemon/external/partners"
	"github.com/yearn/ydaemon/external/prices"
	"github.com/yearn/ydaemon/external/strategies"
	"github.com/yearn/ydaemon/external/tokens"
	"github.com/yearn/ydaemon/external/utils"
	"github.com/yearn/ydaemon/external/vaults"
)

func middleware(log *logrus.Logger) gin.HandlerFunc {
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

// NewRouter create the routes and setup the server
func NewRouter() *gin.Engine {
	gin.EnableJsonDecoderDisallowUnknownFields()
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
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
		router.GET(`vaults/tvl`, c.GetAllVaultsTVL)

		// Retrieve the vaults for a specific chainID
		router.GET(`:chainID/vaults/tvl`, c.GetVaultsTVL)
		router.GET(`:chainID/vaults/all`, c.GetAllVaults)
		router.GET(`:chainID/vaults/:address`, c.GetVault)
		router.GET(`:chainID/vault/:address`, c.GetVault)

		router.GET(`info/vaults/blacklisted`, c.GetBlacklistedVaults)

		router.GET(`:chainID/vaults/harvest/:addresses`, c.GetHarvestForVault)

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

		// Proxy subgraph
		router.POST(`:chainID/graph`, utils.GetGraph)
	}

	// Meta API section
	{
		c := meta.Controller{}
		// Proxy meta strategies
		router.GET(`api/:chainID/strategies/all`, c.GetMetaStrategies)
		router.GET(`:chainID/meta/strategies`, c.GetMetaStrategies)
		router.GET(`api/:chainID/strategies/:address`, c.GetMetaStrategy)
		router.GET(`:chainID/meta/strategies/:address`, c.GetMetaStrategy)
		router.GET(`:chainID/meta/strategy/:address`, c.GetMetaStrategy)

		// Proxy meta tokens
		router.GET(`api/:chainID/tokens/all`, c.GetMetaTokens)
		router.GET(`:chainID/meta/tokens`, c.GetMetaTokens)
		router.GET(`api/:chainID/tokens/:address`, c.GetMetaToken)
		router.GET(`:chainID/meta/tokens/:address`, c.GetMetaToken)
		router.GET(`:chainID/meta/token/:address`, c.GetMetaToken)

		// Proxy meta vaults
		router.GET(`api/:chainID/vaults/all`, c.GetMetaVaultsLegacy)
		router.GET(`:chainID/meta/vaults`, c.GetMetaVaults)
		router.GET(`api/:chainID/vaults/:address`, c.GetMetaVault)
		router.GET(`:chainID/meta/vaults/:address`, c.GetMetaVault)
		router.GET(`:chainID/meta/vault/:address`, c.GetMetaVault)

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
	}

	// Prices API section
	{
		c := prices.Controller{}
		router.GET(`prices/all`, c.GetAllPrices)
		router.GET(`:chainID/prices/all`, c.GetPrices)
		router.GET(`:chainID/prices/:address`, c.GetPrice)
		router.GET(`:chainID/prices/some/:addresses`, c.GetSomePrices)
	}

	{
		router.StaticFile("api/tokens/list", env.BASE_DATA_PATH+`/meta/tokens/tokenList.json`)
	}

	{
		//TEST
		router.GET(`core/harvests/:chainID/:address`, utils.GetHarvests)
	}

	return router
}
