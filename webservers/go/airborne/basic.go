package main

import (
	"context"
	"net/http"

	"github.com/sainzg/autocannon/webservers/go"
	"github.com/sainzg/httpmux"
	"github.com/sainzg/logger"
	"github.com/sainzg/server/config"
	"github.com/sainzg/server/web"
)

func setupServer(log logger.Logger) web.Server {
	mux := httpmux.New()

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(weather.Predict(5))
	})

	cfg, err := config.FromFile("./config/config.yml")
	if err != nil {
		panic(err)
	}

	server, err := web.NewFromConfig(&cfg.Http, mux, log)

	return server
}

func main() {
	log := logger.NewZap()
	defer log.Sync()
	server := setupServer(log)
	server.Start(context.Background(), nil)
}
