package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/env"
)

/**************************************************************************************************
** Controller is the main handler for utility API endpoints. It provides access to system-wide
** configuration information and helper functions that don't fit into more specific domains.
**
** This struct follows the standard API handler pattern in the yDaemon codebase, where controller
** methods are registered as HTTP handlers in the router configuration.
**************************************************************************************************/
type Controller struct{}

/**************************************************************************************************
** GetSupportedChains returns a list of chain IDs supported by the yDaemon API.
**
** This endpoint is useful for client applications that need to discover which blockchain
** networks are available for querying data. The response is a simple array of integers
** representing the chain IDs.
**
** No validation or error handling is necessary for this endpoint as it simply returns
** a pre-defined list from the environment configuration.
**
** @param c The Gin context for handling the HTTP request and response
**************************************************************************************************/
func (y Controller) GetSupportedChains(c *gin.Context) {
	c.JSON(http.StatusOK, env.SUPPORTED_CHAIN_IDS)
}

/**************************************************************************************************
** GetAPIVersion returns the current version of the yDaemon API.
**
** This endpoint provides clients with information about the API version, which is useful for
** feature detection and compatibility checks. The response includes the version number and
** any relevant build information.
**
** @param c The Gin context for handling the HTTP request and response
**************************************************************************************************/
func (y Controller) GetAPIVersion(c *gin.Context) {
	version := map[string]string{
		"version": "1.0.0",
		"name":    "yDaemon",
	}
	c.JSON(http.StatusOK, version)
}

/**************************************************************************************************
** HealthCheck provides a simple endpoint for monitoring systems to verify that the API
** is running and responsive. This is useful for load balancers, container orchestration
** systems, and other monitoring tools.
**
** The response is a 200 OK status with a simple JSON object indicating the service is healthy.
**
** @param c The Gin context for handling the HTTP request and response
**************************************************************************************************/
func (y Controller) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
	})
}
