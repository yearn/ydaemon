package store

import (
	"context"

	"golang.org/x/time/rate"
)

var storeRateLimiter = rate.NewLimiter(4, 8)

func wait(name string) {
	// logs.Warning(`Limiter for ` + name + `: ` + strconv.FormatFloat(storeRateLimiter.Tokens(), 'f', 2, 64))
	storeRateLimiter.Wait(context.Background())
}

func StoreRateLimiter() *rate.Limiter {
	return storeRateLimiter
}
