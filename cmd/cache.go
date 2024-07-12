package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"github.com/yearn/ydaemon/external/vaults"
)

type GetSimplifiedVaults func(c *gin.Context) []vaults.TSimplifiedExternalVault
type GetLegacyExternalVaults func(c *gin.Context) []vaults.TExternalVault

func CacheSimplifiedVaults(cachingStore *cache.Cache, expire time.Duration, handle GetSimplifiedVaults) gin.HandlerFunc {
	return func(c *gin.Context) {
		if result, found := cachingStore.Get(c.Request.URL.String()); found {
			c.JSON(http.StatusOK, result)
		} else {
			result := handle(c)
			cachingStore.Set(c.Request.URL.String(), result, expire)
			c.JSON(http.StatusOK, result)
		}
	}
}

func CacheLegacyVaults(cachingStore *cache.Cache, expire time.Duration, handle GetLegacyExternalVaults) gin.HandlerFunc {
	return func(c *gin.Context) {
		if result, found := cachingStore.Get(c.Request.URL.String()); found {
			c.JSON(http.StatusOK, result)
		} else {
			result := handle(c)
			cachingStore.Set(c.Request.URL.String(), result, expire)
			c.JSON(http.StatusOK, result)
		}
	}
}
