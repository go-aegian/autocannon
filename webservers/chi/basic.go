package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func setupRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Chi router running..."))
	})

	return r
}

func main() {
	r := setupRouter()
	http.ListenAndServe(":3001", r)
}
