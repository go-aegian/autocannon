package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func setupRouter() *echo.Echo {
	r := echo.New()

	r.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Echo router running...")
	})

	return r
}

func main() {
	r := setupRouter()
	r.Logger.Fatal(r.Start(":3003"))
}
