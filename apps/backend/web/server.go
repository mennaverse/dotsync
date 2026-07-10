package web

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/mennaverse/dotsync/apps/backend/web/handler"
	"github.com/mennaverse/dotsync/apps/backend/web/handler/types"
)

func StartServer(bindAddress string) error {
	e := echo.New()
	e.HTTPErrorHandler = errorHandler
	e.Validator = &customValidator{validator: validator.New()}

	handler.RegisterRoutes(e)

	fmt.Printf("Serving Dotsync in %s\n", bindAddress)
	return e.Start(bindAddress)
}

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		var validationErrors validator.ValidationErrors
		validationErrors = err.(validator.ValidationErrors)
		return validationErrors
	}
	return nil
}

func errorHandler(c *echo.Context, err error) {
	errString := fmt.Sprintf("Error: %s", err.Error())
	c.Logger().Error(errString)

	if resp, uErr := echo.UnwrapResponse(c.Response()); uErr == nil {
		if resp.Committed {
			return
		}
	}

	code := http.StatusInternalServerError
	var sc echo.HTTPStatusCoder
	if errors.As(err, &sc) {
		if tmp := sc.StatusCode(); tmp != 0 {
			code = tmp
		}
	}

	var cErr error
	if c.Request().Method == http.MethodHead {
		cErr = c.NoContent(code)
	} else {
		cErr = c.JSON(code, types.JsonResponse{
			"code":    code,
			"message": err.Error(),
		})
	}
	if cErr != nil {
		c.Logger().Error("failed to send error response", "error", errors.Join(err, cErr))
	}
}
