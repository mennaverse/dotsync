package web

import (
	"fmt"

	"github.com/labstack/echo/v5"
)

func StartServer(bindAddress string) error {
	e := echo.New()

	fmt.Printf("Serving Dotsync in %s\n", bindAddress)
	return e.Start(bindAddress)
}
