package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func setupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Gorilla router running..."))
	})

	return r
}

func main() {
	r := setupRouter()
	srv := &http.Server{
		Handler:      r,
		Addr:         ":3005",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
