package middleware

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/upinmcSE/goshop/internal/utils"
	"golang.org/x/time/rate"
)

type Client struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var (
	mu      sync.Mutex
	clients = make(map[string]*Client)
)

func getClientIP(ctx *gin.Context) string {
	ip := ctx.ClientIP()
	if ip == "" {
		ip = ctx.Request.RemoteAddr
	}

	return ip
}

func getRateLimiter(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	client, exists := clients[ip]
	if !exists {

		requestSecStr := utils.GetEnv("RATE_LIMITER_REQUEST_SEC", "5")
		burstStr := utils.GetEnv("RATE_LIMITER_REQUEST_BURST", "10")

		requestSec, err := strconv.Atoi(requestSecStr)
		if err != nil {
			panic(""+err.Error())
		}

		burst, err := strconv.Atoi(burstStr)
		if err != nil {
			panic(""+err.Error())
		}

		limiter := rate.NewLimiter(rate.Limit(requestSec), burst) // 5 request/sec, brust 10
		newClient := &Client{limiter, time.Now()}
		clients[ip] = newClient
		return limiter
	}

	client.lastSeen = time.Now()
	return client.limiter
}

func CleanupClients() {
	for {
		time.Sleep(time.Minute)
		mu.Lock()
		for ip, client := range clients {
			if time.Since(client.lastSeen) > 3*time.Minute {
				delete(clients, ip)
			}
		}
		mu.Unlock()
	}
}

// ab -n 20 -c 1 -H "X-API-Key:ab2a7a8a-d601-4bf7-b0e2-dd00e5459392" http://localhost:8080/api/v1/categories/golang
func RateLimiterMiddleware(rateLimterLogger *zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ip := getClientIP(ctx)

		limiter := getRateLimiter(ip)

		if !limiter.Allow() {
			if shoudLogRateLimit(ip) {
				rateLimterLogger.Warn().
					Str("method", ctx.Request.Method).
					Str("path", ctx.Request.URL.Path).
					Str("query", ctx.Request.URL.RawQuery).
					Str("client_ip", ctx.ClientIP()).
					Str("user_agent", ctx.Request.UserAgent()).
					Str("referer", ctx.Request.Referer()).
					Str("protocol", ctx.Request.Proto).
					Str("host", ctx.Request.Host).
					Str("remote_addr", ctx.Request.RemoteAddr).
					Str("request_uri", ctx.Request.RequestURI).
					Interface("headers", ctx.Request.Header).
					Msg("rate limiter exceeded")
			}

			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error":   "Too many request",
				"message": "Bạn đã gửi quá nhiêu request. Hãy thử lại sau",
			})

			return
		}

		ctx.Next()
	}
}

var rateLimitLogCache = sync.Map{}

const rateLimitLogTTL = 20 * time.Second

func shoudLogRateLimit(ip string) bool {
	now := time.Now()

	if val, ok := rateLimitLogCache.Load(ip); ok {
		if t, ok := val.(time.Time); ok && now.Sub(t) < rateLimitLogTTL {
			return false
		}
	}

	rateLimitLogCache.Store(ip, now)
	return true
}
