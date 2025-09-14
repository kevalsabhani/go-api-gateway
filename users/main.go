package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	userRouter := r.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("/healthcheck", HealthCheckHandler).Methods("GET")

	healthServer := http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Println("users server is running on port 8080")
	if err := healthServer.ListenAndServe(); err != nil {
		panic(err)
	}
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"service": "users",
		"status":  "OK",
	})
}
