package main

import (
	"github.com/aerogo/aero"
	"github.com/gsainz/autocannon/webservers/go"
)

func setupServer() *aero.Application {
	server := aero.New()
	server.Config.Ports.HTTP = 3009

	server.Get("/", func(ctx aero.Context) error {
		return ctx.JSON(string(weather.Predict(5)))
	})

	return server
}

func main() {
	server := setupServer()
	server.Run()
}
