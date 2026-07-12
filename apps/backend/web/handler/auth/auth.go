package auth

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/mennaverse/dotsync/apps/backend/consts"
	"github.com/mennaverse/dotsync/apps/backend/dto"
	"github.com/mennaverse/dotsync/apps/backend/service"
	"github.com/mennaverse/dotsync/apps/backend/web/handler/types"
)

func registerHandler(c *echo.Context) error {
	services := c.Get(consts.ServicesContextKey).(*service.Services)

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
		err = consts.ConvertToAppError(err)
		return err
	}

	return c.JSON(http.StatusOK, types.JsonResponse{
		"code":    consts.HttpSuccessCode,
		"message": "User registered successfully.",
		"data":    dto.UserResponseFromDB(user),
	})
}

func loginHandler(c *echo.Context) error {
	services := c.Get(consts.ServicesContextKey).(*service.Services)

	ctx := c.Request().Context()
	req := new(dto.LoginRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	pair, err := services.Authentication.Login(ctx, req)
	if err != nil {
		err = consts.ConvertToAppError(err)
		return err
	}

	c.SetCookie(&http.Cookie{
		Name:     consts.AccessTokenCookieName,
		Value:    pair.AccessToken,
		SameSite: http.SameSiteStrictMode,
		HttpOnly: true,
		Path:     "/",
	})

	c.SetCookie(&http.Cookie{
		Name:     consts.RefreshTokenCookieName,
		Value:    pair.RefreshToken,
		SameSite: http.SameSiteStrictMode,
		HttpOnly: true,
		Path:     "/",
	})

	return c.JSON(http.StatusOK, types.JsonResponse{
		"code":    consts.HttpSuccessCode,
		"message": "User logged in successfully.",
		"data":    pair,
	})
}

func logoutHandler(c *echo.Context) error {
	refreshToken := c.Get(consts.UserRefreshTokenContextKey).(string)
	services := c.Get(consts.ServicesContextKey).(*service.Services)

	ctx := c.Request().Context()
	if err := services.Authentication.Logout(ctx, refreshToken); err != nil {
		err = consts.ConvertToAppError(err)
		return err
	}

	c.SetCookie(&http.Cookie{
		Name:     consts.AccessTokenCookieName,
		Value:    "",
		SameSite: http.SameSiteStrictMode,
		HttpOnly: true,
		Path:     "/",
	})
	c.SetCookie(&http.Cookie{
		Name:     consts.RefreshTokenCookieName,
		Value:    "",
		SameSite: http.SameSiteStrictMode,
		HttpOnly: true,
		Path:     "/",
	})

	return c.JSON(http.StatusOK, types.JsonResponse{
		"code":    consts.HttpSuccessCode,
		"message": "User logged out successfully.",
	})
}

func logoutAllSessionsHandler(c *echo.Context) error {
	claims := c.Get(consts.UserClaimsContextKey).(*dto.Claims)
	services := c.Get(consts.ServicesContextKey).(*service.Services)

	ctx := c.Request().Context()
	if err := services.Authentication.LogoutAllSessions(ctx, claims.UserID); err != nil {
		err = consts.ConvertToAppError(err)
		return err
	}

	c.SetCookie(&http.Cookie{
		Name:     consts.AccessTokenCookieName,
		Value:    "",
		SameSite: http.SameSiteStrictMode,
		HttpOnly: true,
		Path:     "/",
	})
	c.SetCookie(&http.Cookie{
		Name:     consts.RefreshTokenCookieName,
		Value:    "",
		SameSite: http.SameSiteStrictMode,
		HttpOnly: true,
		Path:     "/",
	})

	return c.JSON(http.StatusOK, types.JsonResponse{
		"code":    consts.HttpSuccessCode,
		"message": "User logged out from all sessions successfully.",
	})
}

func verifyEmailHandler(c *echo.Context) error {
	services := c.Get(consts.ServicesContextKey).(*service.Services)
	token := c.QueryParam("token")

	if token == "" {
		return types.NewAppError(
			http.StatusBadRequest,
			consts.HttpMissingTokenCode,
			"Missing token in query parameters.",
			nil,
		)
	}

	ctx := c.Request().Context()
	if err := services.Authentication.VerifyEmail(ctx, token); err != nil {
		err = consts.ConvertToAppError(err)
		return err
	}

	return c.JSON(http.StatusOK, types.JsonResponse{
		"code":    consts.HttpSuccessCode,
		"message": "Email verified successfully.",
	})
}

type ResendVerificationEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}

func resendVerificationEmailHandler(c *echo.Context) error {
	services := c.Get(consts.ServicesContextKey).(*service.Services)

	ctx := c.Request().Context()
	req := new(ResendVerificationEmailRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	if err := services.Authentication.ResendVerificationEmail(ctx, req.Email); err != nil {
		err = consts.ConvertToAppError(err)
		return err
	}

	return c.JSON(http.StatusOK, types.JsonResponse{
		"code":    consts.HttpSuccessCode,
		"message": "Email sent successfully.",
	})
}

func forgotPasswordHandler(c *echo.Context) error {
	services := c.Get(consts.ServicesContextKey).(*service.Services)

	ctx := c.Request().Context()
	req := new(dto.ForgotPasswordRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	if err := services.Authentication.ForgotPassword(ctx, req.Email); err != nil {
		err = consts.ConvertToAppError(err)
		return err
	}

	return c.JSON(http.StatusOK, types.JsonResponse{
		"code":    consts.HttpSuccessCode,
		"message": "Password reset email sent successfully.",
	})
}

func resetPasswordHandler(c *echo.Context) error {
	services := c.Get(consts.ServicesContextKey).(*service.Services)
	token := c.QueryParam("token")

	if token == "" {
		return c.JSON(http.StatusBadRequest, types.JsonResponse{
			"code":    consts.HttpMissingTokenCode,
			"message": "Missing token in query parameters.",
		})
	}

	ctx := c.Request().Context()

	req := new(dto.ResetPasswordRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	if err := services.Authentication.ResetPassword(ctx, token, req.NewPassword); err != nil {
		err = consts.ConvertToAppError(err)
		return err
	}

	return c.JSON(http.StatusOK, types.JsonResponse{
		"code":    consts.HttpSuccessCode,
		"message": "Password reset successfully.",
	})
}

func validateAccessTokenHandler(c *echo.Context) error {
	claims := c.Get(consts.UserClaimsContextKey).(*dto.Claims)
	return c.JSON(http.StatusOK, types.JsonResponse{
		"code":    consts.HttpSuccessCode,
		"message": "Access token is valid.",
		"data":    claims,
	})
}
