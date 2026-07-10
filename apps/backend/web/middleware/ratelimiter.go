package middleware

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
	echomiddleware "github.com/labstack/echo/v5/middleware"
	"github.com/mennaverse/dotsync/apps/backend/consts"
)

var rateLimiterConfig = echomiddleware.RateLimiterConfig{
	Skipper: echomiddleware.DefaultSkipper,
	Store: echomiddleware.NewRateLimiterMemoryStoreWithConfig(
		echomiddleware.RateLimiterMemoryStoreConfig{
			Rate:      1,
			Burst:     2,
			ExpiresIn: 3 * time.Minute,
		},
	),
	IdentifierExtractor: func(c *echo.Context) (string, error) {
		return c.RealIP(), nil
	},
	ErrorHandler: func(c *echo.Context, err error) error {
		return c.JSON(http.StatusTooManyRequests, map[string]any{
			"code":  consts.HttpTooManyRequestsCode,
			"error": "Too many requests. Please try again later.",
		})
	},
	DenyHandler: func(c *echo.Context, identifier string, err error) error {
		return c.JSON(http.StatusTooManyRequests, map[string]any{
			"code":  consts.HttpTooManyRequestsCode,
			"error": "Too many requests. Please try again later.",
		})
	},
}

var RateLimiter = echomiddleware.RateLimiterWithConfig(rateLimiterConfig)
