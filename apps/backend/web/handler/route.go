package handler

import (
	"github.com/labstack/echo/v5"
	"github.com/mennaverse/dotsync/apps/backend/web/handler/auth"
)

func RegisterRoutes(e *echo.Echo) {
	g := e.Group("/api")

	auth.RegisterRoutes(g)
}
