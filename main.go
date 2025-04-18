package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"

	"github.com/dtg-lucifer/goth-stack-starter/handlers"
	"github.com/dtg-lucifer/goth-stack-starter/middlewares"
)

func main() {
	if err := godotenv.Load(); err != nil {
		slog.Error("Error loading .env file", "file", ".env")
		slog.Info("Loading system env variables")
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	listenAddr := os.Getenv("LISTEN_ADDR")
	router := chi.NewMux()

	router.Use(middlewares.Recover)
	router.Use(middlewares.Logger(logger))

	router.Handle("/*", public())
	router.Get("/", middlewares.HandleError(handlers.Index))

	// Counter API endpoints
	router.Post("/api/counter/increment", middlewares.HandleError(handlers.CounterIncrement))
	router.Post("/api/counter/decrement", middlewares.HandleError(handlers.CounterDecrement))
	router.Post("/api/counter/reset", middlewares.HandleError(handlers.CounterReset))

	if err := http.ListenAndServe(listenAddr, router); err != nil {
		slog.Error("Error starting server", "ListenAddr", listenAddr)
	}
}
