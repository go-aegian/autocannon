package main

import (
	"net/http"

	"github.com/gsainz/airborne"
)

func setupServer() *airborne.Mux {
	r := airborne.NewMux()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Airborne router running..."))
	})

	return r
}

func main() {
	r := setupServer()
	http.ListenAndServe(":3002", r)
}
