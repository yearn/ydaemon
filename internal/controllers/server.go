package controllers

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// NewRouter create the routes and setup the server
func NewRouter() *gin.Engine {
	gin.EnableJsonDecoderDisallowUnknownFields()
	router := gin.New()
	router.Use(gin.Recovery())
	logger := logrus.StandardLogger()
	logger.SetFormatter(&logrus.JSONFormatter{})
	router.Use(ginrus.Ginrus(logger, time.RFC3339, true))
	corsConf := cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:    []string{`Origin`, `Content-Length`, `Content-Type`, `Authorization`},
	}
	router.Use(cors.New(corsConf))

	{
		c := controller{}
		router.GET(`something-something`, c.DoSomething)
	}

	return router
}
