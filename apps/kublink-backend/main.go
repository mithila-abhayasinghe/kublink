package main

import (
	"fmt"
	"kublink-backend/internal/handlers"
	"kublink-backend/internal/health"
	"kublink-backend/internal/store"
	"net/http"
	"os"
	"time"
)

func main() {

	port := func() string {
		port := os.Getenv("PORT")
		if port == "" {
			port = "8000"
		}
		return ":" + port
	}()

	// Dependencies
	urlStore := store.NewURLStore()
	handler := &handlers.URLHandler{
		Store:      urlStore,
		CurrentURL: "http://localhost" + port,
	}

	router := http.NewServeMux()
	router.HandleFunc("GET /{shortUrl}", handler.GetLongUrl)
	router.HandleFunc("POST /api/v1/urls", handler.CreateShortUrl)
	router.HandleFunc("GET /healthz", health.HealthCheck)

	server := &http.Server{
		Addr:         port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	fmt.Println("Server starting on", port)
	err := server.ListenAndServe()

	if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

}
