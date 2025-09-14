package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	pingService := "http://localhost:8081"
	healthService := "http://localhost:8080"

	pingProxy, err := newReverseProxy(pingService)
	if err != nil {
		log.Fatalf("Failed to create ping proxy: %v", err)
	}
	healthProxy, err := newReverseProxy(healthService)
	if err != nil {
		log.Fatalf("Failed to create health proxy: %v", err)
	}

	r.PathPrefix("/products/").HandlerFunc(gatewayHandler(pingProxy))
	r.PathPrefix("/users/").HandlerFunc(gatewayHandler(healthProxy))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to the Gateway! Use /ping or /health endpoints."))
	})

	log.Println("Gateway server is running on port 8082")
	if err := http.ListenAndServe(":8082", r); err != nil {
		log.Fatalf("Failed to start gateway server: %v", err)
	}
}

func newReverseProxy(target string) (*httputil.ReverseProxy, error) {
	url, err := url.Parse(target)
	if err != nil {
		return nil, fmt.Errorf("invalid target URL: %w", err)
	}

	return httputil.NewSingleHostReverseProxy(url), nil
}

func gatewayHandler(proxy *httputil.ReverseProxy) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Proxying request to Host: %s, Path: %s", r.Host, r.URL.Path)
		proxy.ServeHTTP(w, r)
	}
}
