package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
)

// RateLimitConfig defines the config for RateLimit middleware.
type RateLimitConfig struct {
	// Limit is the maximum number of requests to allow per period.
	Limit int
	// Period is the duration in which the limit is enforced.
	Period time.Duration
}

// RateLimit creates a rate limiting middleware.
func RateLimit(config RateLimitConfig) echo.MiddlewareFunc {
	var mu sync.Mutex
	var hits = make(map[string]int)

	// Reset the hits map every "period".
	go func() {
		for {
			time.Sleep(config.Period)
			mu.Lock()
			hits = make(map[string]int)
			mu.Unlock()
		}
	}()

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			mu.Lock()
			defer mu.Unlock()

			ip := c.RealIP()
			if hits[ip] >= config.Limit {
				return c.String(http.StatusTooManyRequests, "too many requests")
			}

			hits[ip]++
			return next(c)
		}
	}
}
