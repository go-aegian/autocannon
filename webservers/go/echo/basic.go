package main

import (
	"net/http"

	"github.com/gsainz/autocannon/webservers/go"
	"github.com/labstack/echo/v4"
)

func setupRouter() *echo.Echo {
	r := echo.New()

	r.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, weather.Predict(5))
	})

	return r
}

func main() {
	r := setupRouter()
	r.Logger.Fatal(r.Start(":3003"))
}
