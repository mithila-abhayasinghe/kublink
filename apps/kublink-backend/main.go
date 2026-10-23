package main

import (
	"context"
	"kublink-backend/internal/handlers"
	"kublink-backend/internal/health"
	"kublink-backend/internal/store"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	baseURL := func(p string) string {
		baseURL := os.Getenv("BASE_URL")
		if baseURL == "" {
			baseURL = "http://localhost" + port
		}
		return baseURL
	}(port)

	// Dependencies
	urlStore := store.NewURLStore()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	handler := &handlers.URLHandler{
		Store:   urlStore,
		BaseURL: baseURL,
		Logger:  logger,
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

	go func() {
		logger.Info("server starting", "port", port)
		err := server.ListenAndServe()
		if err != nil {
			logger.Error("error starting server", "err", err.Error())
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	logger.Info("server is shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Info("server forced to shutdown", "err", err.Error())
	}
	logger.Info("server stopped")
}
