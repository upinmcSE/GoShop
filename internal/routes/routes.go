package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/upinmcSE/goshop/internal/middleware"
	"github.com/upinmcSE/goshop/internal/utils"
)

type Route interface {
	Register(r *gin.RouterGroup)
}

func RegisterRoutes(r *gin.Engine, routes ...Route) {

	httpLogger := utils.NewLoggerWithPath("http.log", "info")
	recoveryLogger := utils.NewLoggerWithPath("recovery.log", "warning")
	ratelimitLogger := utils.NewLoggerWithPath("ratelimit.log", "warning")

	r.Use(
		middleware.RateLimiterMiddleware(ratelimitLogger),
		middleware.LoggerMiddleware(httpLogger),
		middleware.RecoveryMiddleware(recoveryLogger),
		middleware.ApiKeyMiddleware(),
		middleware.AuthMiddleware(),
	)

	api := r.Group("/api/v1")

	for _, route := range routes {
		route.Register(api)
	}
}
