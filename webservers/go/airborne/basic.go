package main

import (
	"net/http"

	"github.com/gsainz/airborne"
	"github.com/gsainz/airborne/config"
	"github.com/gsainz/autocannon/webservers/go"
)

func setupServer() *airborne.HttpServer {
	cfg, err := config.FromFile("./go/airborne/config/config.json")
	if err != nil {
		panic(err)
	}

	server := airborne.NewServer(cfg, airborne.NewMux())

	server.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(weather.Predict(5))
	})

	return server
}

func main() {
	server := setupServer()
	server.Run()
}
