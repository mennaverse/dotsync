package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v5"
	"github.com/mennaverse/dotsync/apps/backend/consts"
	"github.com/mennaverse/dotsync/apps/backend/manager"
)

func RequireAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]any{
					"code":  consts.HttpMissingAuthHeaderCode,
					"error": "Authorization header is missing",
				})
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				return c.JSON(http.StatusUnauthorized, map[string]any{
					"code":  consts.HttpMissingAuthHeaderCode,
					"error": "Authorization header format must be Bearer {token}",
				})
			}

			accessToken := parts[1]

			secrets := manager.NewSecretManager()
			cryptoManager := manager.NewCryptoManager(secrets)

			claims, err := cryptoManager.ParseJwtToken(accessToken)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]any{
					"code":  consts.HttpInvalidAccessTokenCode,
					"error": "Invalid access token",
				})
			}

			c.Set("user_claims", claims)

			return next(c)
		}
	}
}
