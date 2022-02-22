package main

import (
	"net/http"

	"github.com/gsainz/airborne"
	"github.com/gsainz/autocannon/webservers/go"
)

func setupServer() *airborne.HttpServer {
	server := airborne.NewServer("./go/airborne/config/", nil)

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
