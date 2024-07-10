package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"github.com/yearn/ydaemon/common/helpers"
	"golang.org/x/time/rate"
)

var allowlist = []string{
	"https://yearn.finance",
	"https://yearn.fi",
	"https://ycrv.yearn.fi",
	"https://yeth.yearn.fi",
	"https://veyfi.yearn.fi",
	"https://juiced.yearn.fi",
	"https://buyback.yearn.fi",
	"https://seafood.yearn.watch",
	"https://juiced.app",
	"https://ajnafi.com",
	"https://tokenlistooor.com",
	"https://ybribe.com",
	"https://yearn.farm",
	"https://ape.tax",
}
var rootURI = []string{
	".yearn.fi",
	".yearn.finance",
	// ".yearn.farm",
	".juiced.app",
	".smold.app",
	"http://localhost:",
}

/**************************************************************************************************
** Rate limiting based on https://github.com/yangxikun/gin-limit-by-key/blob/master/limit.go and
** adapted to our needs.
**************************************************************************************************/
var limiterSet = cache.New(5*time.Minute, 10*time.Minute)
var accessPerOrigin = make(map[string][2]int64)

func NewRateLimiter(abort func(*gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		/******************************************************************************************
		** Retrieve the origin from the request header and use it as the key for the rate limiter.
		** If the origin is not present, we use an empty string as the key.
		******************************************************************************************/
		k := ``
		origin := c.Request.Header.Get("Origin")
		if len(origin) == 0 {
			k = ``
		} else {
			k = origin
		}

		/******************************************************************************************
		** Allows the requests from the allowlist without rate limiting. This is to allow us to
		** bypass the rate limiting for our own services.
		******************************************************************************************/
		if helpers.Contains(allowlist, origin) || helpers.EndsWithSubstring(rootURI, origin) {
			c.Next()
			return
		}

		/******************************************************************************************
		** Otherwise, we use the rate limiter to limit the requests to 10 qps/clientIp and permit
		******************************************************************************************/
		limiter, ok := limiterSet.Get(k)
		if !ok {
			var expire time.Duration
			// limit 25 query per second per origin and permit bursts of at most 25 tokens, and the limiter liveness time duration is 15 minutes
			limiter, expire = rate.NewLimiter(rate.Every(1*time.Second), 50), 15*time.Minute
			limiterSet.Set(k, limiter, expire)
		}
		ok = limiter.(*rate.Limiter).Allow()
		if !ok {
			abort(c)
			return
		}
		c.Next()
	}
}
