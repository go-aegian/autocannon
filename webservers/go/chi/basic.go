package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gsainz/autocannon/webservers/go"
)

func setupRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(weather.Predict(5))
	})

	return r
}

func main() {
	r := setupRouter()
	log.Fatal(http.ListenAndServe(":3001", r))
}
