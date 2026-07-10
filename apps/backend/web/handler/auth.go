package handler

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/mennaverse/dotsync/apps/backend/consts"
	"github.com/mennaverse/dotsync/apps/backend/dto"
	"github.com/mennaverse/dotsync/apps/backend/service"
)

func registerHandler(c *echo.Context) error {
	services := c.Get("services").(*service.Services)

	ctx := c.Request().Context()
	req := new(dto.RegisterRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	user, err := services.Authentication.Register(ctx, req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, jsonResponse{
		"code":    consts.HttpSuccessCode,
		"message": "User registered successfully.",
		"data":    dto.UserResponseFromDB(user),
	})
}

type ResendEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}

func resendEmailHandler(c *echo.Context) error {
	services := c.Get("services").(*service.Services)

	ctx := c.Request().Context()
	req := new(ResendEmailRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	if err := services.Authentication.ResendVerificationEmail(ctx, req.Email); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, jsonResponse{
		"code":    consts.HttpSuccessCode,
		"message": "Email sent successfully.",
	})
}
