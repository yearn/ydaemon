package controllers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// NewRouter create the routes and setup the server
func NewRouter() *gin.Engine {
	gin.EnableJsonDecoderDisallowUnknownFields()
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(logger(logrus.New()))
	corsConf := cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "HEAD"},
		AllowHeaders:    []string{`Origin`, `Content-Length`, `Content-Type`, `Authorization`},
	}
	router.Use(cors.New(corsConf))

	// General section
	{
		c := controller{}
		// Retrieve the vaults for a specific chainID
		router.GET(`:chainID/vaults/tvl`, c.GetAllVaultsTVL)
		router.GET(`:chainID/vaults/all`, c.GetAllVaults)
		router.GET(`:chainID/vaults/:address`, c.GetVault)
		router.GET(`:chainID/vault/:address`, c.GetVault)

		// Retrieve the reports for a specific strategy
		router.GET(`:chainID/reports/:address`, c.GetReports)

		// Get some information about the API
		router.GET(`info/chains`, c.GetSupportedChains)
		router.GET(`info/vaults/blacklisted`, c.GetBlacklistedVaults)

		// Proxy subgraph
		router.POST(`:chainID/graph`, c.GetGraph)

		// Automatic webhook connected to github to trigger some actions
		router.POST(`webhook/meta/trigger`, c.TriggerMetaRefreshWebhook)
	}

	// Meta API section
	{
		meta := controller{}

		// Proxy meta strategies
		router.GET(`api/:chainID/strategies/all`, meta.GetMetaStrategiesLegacy)
		router.GET(`:chainID/meta/strategies`, meta.GetMetaStrategies)
		router.GET(`api/:chainID/strategies/:address`, meta.GetMetaStrategy)
		router.GET(`:chainID/meta/strategies/:address`, meta.GetMetaStrategy)
		router.GET(`:chainID/meta/strategy/:address`, meta.GetMetaStrategy)

		// Proxy meta tokens
		router.GET(`api/:chainID/tokens/all`, meta.GetMetaTokensLegacy)
		router.GET(`:chainID/meta/tokens`, meta.GetMetaTokens)
		router.GET(`api/:chainID/tokens/:address`, meta.GetMetaToken)
		router.GET(`:chainID/meta/tokens/:address`, meta.GetMetaToken)
		router.GET(`:chainID/meta/token/:address`, meta.GetMetaToken)

		// Proxy meta vaults
		router.GET(`api/:chainID/vaults/all`, meta.GetMetaVaultsLegacy)
		router.GET(`:chainID/meta/vaults`, meta.GetMetaVaults)
		router.GET(`api/:chainID/vaults/:address`, meta.GetMetaVault)
		router.GET(`:chainID/meta/vaults/:address`, meta.GetMetaVault)
		router.GET(`:chainID/meta/vault/:address`, meta.GetMetaVault)

		// Proxy meta protocols
		router.GET(`api/:chainID/protocols/all`, meta.GetMetaProtocolsLegacy)
		router.GET(`:chainID/meta/protocols`, meta.GetMetaProtocols)
		router.GET(`api/:chainID/protocols/:name`, meta.GetMetaProtocol)
		router.GET(`:chainID/meta/protocols/:name`, meta.GetMetaProtocol)
		router.GET(`:chainID/meta/protocol/:name`, meta.GetMetaProtocol)
	}

	return router
}
