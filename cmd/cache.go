package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/external/vaults"
)

type GetSimplifiedVaults func(c *gin.Context) ([]vaults.TSimplifiedExternalVault, error)
type GetLegacyExternalVaults func(c *gin.Context) []vaults.TExternalVault
type GetCustomVaults func(c *gin.Context) []vaults.TRotkiVaults

func CacheSimplifiedVaults(cachingStore *cache.Cache, expire time.Duration, handle GetSimplifiedVaults) gin.HandlerFunc {
	return func(c *gin.Context) {
		if result, found := cachingStore.Get(c.Request.URL.String()); found && result != nil {
			vaults, ok := result.([]vaults.TSimplifiedExternalVault)
			if ok && len(vaults) > 0 {
				logs.Info(`Cache hit with`, len(vaults), `vaults`)
				c.JSON(http.StatusOK, vaults)
				return
			}
		}

		result, err := handle(c)
		if err != nil {
			//Json was already sent
			return
		}
		if len(result) > 0 {
			cachingStore.Set(c.Request.URL.String(), result, expire)
			logs.Info(`Cache miss with`, len(result), `vaults`)
		}
		c.JSON(http.StatusOK, result)
	}
}

func CacheLegacyVaults(cachingStore *cache.Cache, expire time.Duration, handle GetLegacyExternalVaults) gin.HandlerFunc {
	return func(c *gin.Context) {
		if result, found := cachingStore.Get(c.Request.URL.String()); found && result != nil {
			vaults, ok := result.([]vaults.TExternalVault)
			if ok && len(vaults) > 0 {
				c.JSON(http.StatusOK, vaults)
				return
			}
		}

		result := handle(c)
		cachingStore.Set(c.Request.URL.String(), result, expire)
		c.JSON(http.StatusOK, result)
	}
}

func CacheCustomVaults(cachingStore *cache.Cache, expire time.Duration, handle GetCustomVaults) gin.HandlerFunc {
	return func(c *gin.Context) {
		if result, found := cachingStore.Get(c.Request.URL.String()); found && result != nil {
			vaults, ok := result.([]vaults.TRotkiVaults)
			if ok && len(vaults) > 0 {
				c.JSON(http.StatusOK, vaults)
				return
			}
		}

		result := handle(c)
		cachingStore.Set(c.Request.URL.String(), result, expire)
		c.JSON(http.StatusOK, result)
	}
}
