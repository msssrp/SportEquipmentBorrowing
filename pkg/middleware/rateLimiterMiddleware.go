package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var (
	mu    sync.Mutex
	limit = rate.Every(15 * time.Minute)
	ips   = make(map[string]*rate.Limiter)
)

func RateLimiterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		mu.Lock()
		limiter, exists := ips[ip]
		if !exists {
			limiter = rate.NewLimiter(limit, 5) // Allow 5 requests every 15 minutes
			ips[ip] = limiter
		}
		mu.Unlock()

		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
			c.Abort()
			return
		}

		c.Next()
	}
}
