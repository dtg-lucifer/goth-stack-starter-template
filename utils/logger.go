package utils

import (
	"log/slog"
	"os"
)

func NewFileLogger() *slog.Logger {
	// @INFO file logger, stores logs in logs/app.log

	// Create logs directory if it doesn't exist
	// and create the log file
	// check if the directory exists
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		// create the directory
		err := os.Mkdir("logs", os.ModePerm)
		if err != nil {
			slog.Error("Failed to create logs directory", "error", err)
			return nil
		}
	}

	if _, err := os.Stat("logs/app.log"); os.IsNotExist(err) {
		// create the log file
		_, err := os.Create("logs/app.log")
		if err != nil {
			slog.Error("Failed to create log file", "error", err)
			return nil
		}
	}

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
