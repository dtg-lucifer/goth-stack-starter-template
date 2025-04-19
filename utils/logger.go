package utils

import (
	"log/slog"
	"os"
)

func NewFileLogger() *slog.Logger {
	// @TODO Create a new file logger
	file, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		slog.Error("Failed to open log file", "error", err)
		return nil
	}

	return slog.New(slog.NewTextHandler(file, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	}))
}
