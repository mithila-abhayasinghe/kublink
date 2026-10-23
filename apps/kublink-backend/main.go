package main

import (
	"fmt"
	"kublink-backend/internal/handlers"
	"kublink-backend/internal/health"
	"net/http"
	"os"
	"time"
)

func main() {

	router := http.NewServeMux()
	router.HandleFunc("GET /{shortUrl}", handlers.GetLongUrl)
	router.HandleFunc("POST /api/v1/urls", handlers.CreateShortUrl)
	router.HandleFunc("GET /healthz", health.HealthCheck)

	server := &http.Server{
		Addr:         ":8000",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	fmt.Println("Server starting on :8000")
	err := server.ListenAndServe()

	if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

}
