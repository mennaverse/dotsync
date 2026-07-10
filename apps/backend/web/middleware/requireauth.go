package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v5"
	"github.com/mennaverse/dotsync/apps/backend/consts"
	"github.com/mennaverse/dotsync/apps/backend/service"
)

func RequireAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			services := c.Get(consts.ServicesContextKey).(*service.Services)
			ctx := c.Request().Context()
			var accessToken string
			isFromCookie := false

			cookie, err := c.Cookie(consts.AccessTokenCookieName)
			if err == nil {
				accessToken = cookie.Value
				isFromCookie = true
			}

			if accessToken == "" {
				authHeader := c.Request().Header.Get("Authorization")
				if authHeader != "" {
					// Bearer <token>
					parts := strings.Split(authHeader, " ")
					if len(parts) == 2 && strings.ToLower(parts[0]) == "bearer" {
						accessToken = parts[1]
					}
				}
			}

			if accessToken == "" {
				return c.JSON(http.StatusUnauthorized, map[string]any{
					"code":  consts.HttpMissingAuthHeaderCode,
					"error": "Missing access token in cookie or Authorization header",
				})
			}

			claims, err := services.Authentication.ValidateAccessToken(ctx, accessToken)
			if err != nil && errors.Is(err, consts.ErrTokenInvalid) {
				var refreshToken string

				if isFromCookie {
					refreshCookie, err := c.Cookie(consts.RefreshTokenCookieName)
					if err == nil {
						refreshToken = refreshCookie.Value
					}
				} else {
					refreshToken = c.Request().Header.Get("X-Refresh-Token")
				}

				if refreshToken == "" {
					return c.JSON(http.StatusUnauthorized, map[string]any{
						"code":  consts.HttpInvalidAccessTokenCode,
						"error": "Invalid access token and missing refresh token",
					})
				}

				newTokens, err := services.Authentication.RefreshToken(ctx, refreshToken)
				if err != nil {
					return c.JSON(http.StatusUnauthorized, map[string]any{
						"code":  consts.ErrHttpInvalidRefreshTokenCode,
						"error": "Invalid refresh token",
					})
				}
				// Set the new access token in the cookie if it was from the cookie
				if isFromCookie {
					c.SetCookie(&http.Cookie{
						Name:     consts.AccessTokenCookieName,
						Value:    newTokens.AccessToken,
						HttpOnly: true,
						Path:     "/",
					})
					c.SetCookie(&http.Cookie{
						Name:     consts.RefreshTokenCookieName,
						Value:    newTokens.RefreshToken,
						HttpOnly: true,
						Path:     "/",
					})
				} else {
					// If the access token was from the Authorization header, set the new tokens in the response headers
					c.Response().Header().Set("X-New-Access-Token", newTokens.AccessToken)
					c.Response().Header().Set("X-New-Refresh-Token", newTokens.RefreshToken)
				}

				claims, err = services.Authentication.ValidateAccessToken(ctx, newTokens.AccessToken)
				if err != nil {
					return c.JSON(http.StatusUnauthorized, map[string]any{
						"code":  consts.HttpInvalidAccessTokenCode,
						"error": "Invalid access token after refresh",
					})
				}

				c.Set(consts.UserRefreshTokenContextKey, newTokens.RefreshToken)
			} else if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]any{
					"code":  consts.HttpInvalidAccessTokenCode,
					"error": "Invalid access token",
				})
			}

			c.Set(consts.UserClaimsContextKey, claims)

			return next(c)
		}
	}
}
