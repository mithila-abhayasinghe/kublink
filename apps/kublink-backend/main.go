package main

import (
	"fmt"
	"kublink-backend/internal/handlers"
	"kublink-backend/internal/health"
	"kublink-backend/internal/store"
	"log/slog"
	"net/http"
	"os"
	"time"

	_ "kublink-backend/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title			Kublink: URL Shortener
// @version		0.1.0
// @description	A URL shortener service in Go.
// @host			localhost:8000
// @BasePath		/
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
	logger := slog.NewJSONHandler(os.Stdout, nil)
	handler := &handlers.URLHandler{
		Store:   urlStore,
		BaseURL: "http://localhost" + port,
		Logger:  slog.New(logger),
	}

	router := http.NewServeMux()
	router.Handle("/swagger/", httpSwagger.WrapHandler)
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
