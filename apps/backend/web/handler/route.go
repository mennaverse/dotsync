package handler

import (
	"github.com/labstack/echo/v5"
	"github.com/mennaverse/dotsync/apps/backend/web/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	g := e.Group("/api")

	g.POST(
		"/auth/register",
		registerHandler,
		middleware.RateLimiter,
		middleware.Services(),
	)

	g.POST(
		"/auth/resend-email",
		resendEmailHandler,
		middleware.RateLimiter,
		middleware.Services(),
	)
}
