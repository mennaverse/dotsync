package auth

import (
	"github.com/labstack/echo/v5"
	"github.com/mennaverse/dotsync/apps/backend/web/middleware"
)

func RegisterRoutes(e *echo.Group) {
	e.POST(
		"/auth/register",
		registerHandler,
		middleware.RateLimiter,
		middleware.Services(),
	)

	e.POST(
		"/auth/login",
		loginHandler,
		middleware.RateLimiter,
		middleware.Services(),
	)

	e.POST(
		"/auth/logout",
		logoutHandler,
		middleware.RateLimiter,
		middleware.Services(),
		middleware.RequireAuth(),
	)

	e.POST(
		"/auth/logout-all",
		logoutAllSessionsHandler,
		middleware.RateLimiter,
		middleware.Services(),
		middleware.RequireAuth(),
	)

	e.POST(
		"/auth/verify-email",
		verifyEmailHandler,
		middleware.RateLimiter,
		middleware.Services(),
	)

	e.POST(
		"/auth/resend-verification-email",
		resendVerificationEmailHandler,
		middleware.RateLimiter,
		middleware.Services(),
	)

	e.POST(
		"/auth/forgot-password",
		forgotPasswordHandler,
		middleware.RateLimiter,
		middleware.Services(),
	)

	e.POST(
		"/auth/reset-password",
		resetPasswordHandler,
		middleware.RateLimiter,
		middleware.Services(),
	)

	e.GET(
		"/auth/validate-access-token",
		validateAccessTokenHandler,
		middleware.Services(),
		middleware.RequireAuth(),
	)
}
