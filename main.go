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

	listenAddr := os.Getenv("LISTEN_ADDR")
	router := chi.NewMux()

	router.Handle("/*", public())
	router.Get("/", middlewares.HandleError(handlers.Index))

	if err := http.ListenAndServe(listenAddr, router); err != nil {
		slog.Error("Error starting server", "ListenAddr", listenAddr)
	}
}
