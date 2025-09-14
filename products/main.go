package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	productRouter := r.PathPrefix("/products").Subrouter()
	productRouter.HandleFunc("/healthcheck", HealthCheckHandler).Methods("GET")

	pingServer := http.Server{
		Addr:    ":8081",
		Handler: r,
	}

	log.Println("products server is running on port 8081")
	if err := pingServer.ListenAndServe(); err != nil {
		panic(err)
	}
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"service": "products",
		"status":  "OK",
	})
}
