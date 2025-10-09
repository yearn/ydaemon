package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/external/vaults"
	"golang.org/x/sync/singleflight"
)

type GetSimplifiedVaults func(c *gin.Context) ([]vaults.TSimplifiedExternalVault, error)
type GetLegacyExternalVaults func(c *gin.Context) []vaults.TExternalVault
type GetCustomVaults func(c *gin.Context) []vaults.TRotkiVaults

var simplifiedVaultsSingleflight singleflight.Group
var legacyVaultsSingleflight singleflight.Group
var customVaultsSingleflight singleflight.Group

func CacheSimplifiedVaults(cachingStore *cache.Cache, expire time.Duration, handle GetSimplifiedVaults) gin.HandlerFunc {
	return func(c *gin.Context) {
		cacheKey := c.Request.URL.String()

		// Check cache first
		if result, found := cachingStore.Get(cacheKey); found && result != nil {
			vaults, ok := result.([]vaults.TSimplifiedExternalVault)
			if ok && len(vaults) > 0 {
				c.JSON(http.StatusOK, vaults)
				return
			}
		}

		// Use singleflight to prevent thundering herd on cache miss
		// Only one goroutine will execute the handler, others will wait for the result
		result, err, shared := simplifiedVaultsSingleflight.Do(cacheKey, func() (interface{}, error) {
			vaults, err := handle(c)
			if err != nil {
				return nil, err
			}

			// Cache the result
			if len(vaults) > 0 {
				cachingStore.Set(cacheKey, vaults, expire)
				logs.Info(`Cache miss with`, len(vaults), `vaults`)
			}

			return vaults, nil
		})

		if err != nil {
			logs.Error(`Error while getting simplified vaults`, err)
			return
		}

		if shared {
			logs.Info(`Singleflight shared result with`, len(result.([]vaults.TSimplifiedExternalVault)), `vaults`)
		}

		c.JSON(http.StatusOK, result)
	}
}

func CacheLegacyVaults(cachingStore *cache.Cache, expire time.Duration, handle GetLegacyExternalVaults) gin.HandlerFunc {
	return func(c *gin.Context) {
		cacheKey := c.Request.URL.String()

		// Check cache first
		if result, found := cachingStore.Get(cacheKey); found && result != nil {
			vaults, ok := result.([]vaults.TExternalVault)
			if ok && len(vaults) > 0 {
				c.JSON(http.StatusOK, vaults)
				return
			}
		}

		// Use singleflight to prevent thundering herd on cache miss
		result, err, shared := legacyVaultsSingleflight.Do(cacheKey, func() (interface{}, error) {
			vaults := handle(c)

			// Cache the result
			if len(vaults) > 0 {
				cachingStore.Set(cacheKey, vaults, expire)
				logs.Info(`Cache miss with`, len(vaults), `legacy vaults`)
			}

			return vaults, nil
		})

		if err != nil {
			logs.Error(`Error while getting legacy vaults`, err)
			return
		}

		if shared {
			logs.Info(`Singleflight shared result with`, len(result.([]vaults.TExternalVault)), `legacy vaults`)
		}

		c.JSON(http.StatusOK, result)
	}
}

func CacheCustomVaults(cachingStore *cache.Cache, expire time.Duration, handle GetCustomVaults) gin.HandlerFunc {
	return func(c *gin.Context) {
		cacheKey := c.Request.URL.String()

		// Check cache first
		if result, found := cachingStore.Get(cacheKey); found && result != nil {
			vaults, ok := result.([]vaults.TRotkiVaults)
			if ok && len(vaults) > 0 {
				c.JSON(http.StatusOK, vaults)
				return
			}
		}

		// Use singleflight to prevent thundering herd on cache miss
		result, err, shared := customVaultsSingleflight.Do(cacheKey, func() (interface{}, error) {
			vaults := handle(c)

			// Cache the result
			if len(vaults) > 0 {
				cachingStore.Set(cacheKey, vaults, expire)
				logs.Info(`Cache miss with`, len(vaults), `custom vaults`)
			}

			return vaults, nil
		})

		if err != nil {
			logs.Error(`Error while getting custom vaults`, err)
			return
		}

		if shared {
			logs.Info(`Singleflight shared result with`, len(result.([]vaults.TRotkiVaults)), `custom vaults`)
		}

		c.JSON(http.StatusOK, result)
	}
}
