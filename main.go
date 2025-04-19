package main

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		slog.Error("Error loading .env file", "file", ".env")
		slog.Info("Loading system env variables")
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	}))

	slog.SetDefault(logger)
	slog.Info("Starting goth-stack-starter")

	listenAddr := os.Getenv("LISTEN_ADDR")

	server := NewServer(listenAddr, logger)
	server.Start()
}
