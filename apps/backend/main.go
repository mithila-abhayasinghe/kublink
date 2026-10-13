package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {

	router := http.NewServeMux()
	router.HandleFunc("GET /", getRoot)

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

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("GET / Root Request\n")
	io.WriteString(w, "Kublink: URL Shortener\n")
}
