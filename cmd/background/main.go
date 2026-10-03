package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/cakmakfatih/chattered-background/internal/health"
	"github.com/cakmakfatih/chattered-background/internal/worker"
)

const httpAddr = ":8080"

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		slog.Error("REDIS_ADDR is required")
		os.Exit(1)
	}

	redisPassword := os.Getenv("REDIS_PASSWORD")
	if redisPassword == "" {
		slog.Error("REDIS_PASSWORD is required")
		os.Exit(1)
	}

	go func() {
		slog.Info("health HTTP server listening", "address", httpAddr)
		if err := http.ListenAndServe(httpAddr, health.NewHandler()); err != nil {
			slog.Error("health HTTP server stopped", "error", err)
			os.Exit(1)
		}
	}()

	if err := worker.Run(redisAddr, redisPassword); err != nil {
		slog.Error("background worker stopped", "error", err)
		os.Exit(1)
	}
}
