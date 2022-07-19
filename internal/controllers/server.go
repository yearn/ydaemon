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

	{
		c := controller{}
		// Retrieve the vaults for a specific chainID
		router.GET(`:chainID/vaults/all`, c.GetAllVaults)
		router.GET(`:chainID/vaults/:address`, c.GetVault)

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

	return router
}
