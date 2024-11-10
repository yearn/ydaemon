package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/external/vaults"
)

type GetSimplifiedVaults func(c *gin.Context) []vaults.TSimplifiedExternalVault
type GetLegacyExternalVaults func(c *gin.Context) []vaults.TExternalVault
type GetCustomVaults func(c *gin.Context) []vaults.TRotkiVaults

func CacheSimplifiedVaults(cachingStore *cache.Cache, expire time.Duration, handle GetSimplifiedVaults) gin.HandlerFunc {
	return func(c *gin.Context) {
		if result, found := cachingStore.Get(c.Request.URL.String()); found && result != nil && len(result.([]vaults.TSimplifiedExternalVault)) > 0 {
			logs.Info(`Cache hit with`, len(result.([]vaults.TSimplifiedExternalVault)), `vaults`)
			c.JSON(http.StatusOK, result)
		} else {
			result := handle(c)
			cachingStore.Set(c.Request.URL.String(), result, expire)
			logs.Info(`Cache miss with`, len(result), `vaults`)
			c.JSON(http.StatusOK, result)
		}
	}
}

func CacheLegacyVaults(cachingStore *cache.Cache, expire time.Duration, handle GetLegacyExternalVaults) gin.HandlerFunc {
	return func(c *gin.Context) {
		if result, found := cachingStore.Get(c.Request.URL.String()); found && result != nil && len(result.([]vaults.TExternalVault)) > 0 {
			c.JSON(http.StatusOK, result)
		} else {
			result := handle(c)
			cachingStore.Set(c.Request.URL.String(), result, expire)
			c.JSON(http.StatusOK, result)
		}
	}
}

func CacheCustomVaults(cachingStore *cache.Cache, expire time.Duration, handle GetCustomVaults) gin.HandlerFunc {
	return func(c *gin.Context) {
		if result, found := cachingStore.Get(c.Request.URL.String()); found && result != nil && len(result.([]vaults.TRotkiVaults)) > 0 {
			c.JSON(http.StatusOK, result)
		} else {
			result := handle(c)
			cachingStore.Set(c.Request.URL.String(), result, expire)
			c.JSON(http.StatusOK, result)
		}
	}
}
