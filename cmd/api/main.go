package main

import (
	"github.com/xyz2200/go-app/app/hello/entrypoint"
	"log"
	"log/slog"
	"net/http"
)

var logger = slog.Default()

func main() {

	handler := entrypoint.NewHelloHandler()
	handler.Register()

	// Start server on port 8080 (adjust if needed)
	logger.Info("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
