package main

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/dtg-lucifer/goth-stack-starter/handlers"
	"github.com/dtg-lucifer/goth-stack-starter/middlewares"
	"github.com/dtg-lucifer/goth-stack-starter/state"
	"github.com/dtg-lucifer/goth-stack-starter/utils"
)

type Server struct {
	Addr   string
	Store  *state.Store
	Router *chi.Mux
	Logger *slog.Logger
}

func NewServer(addr string, logger *slog.Logger) *Server {
	return &Server{
		Addr:   addr,
		Store:  state.NewStore(),
		Router: chi.NewMux(),
		Logger: logger,
	}
}

func (s *Server) setupMiddlewares() {
	// @TODO setup cors
	// s.Router.Use(CorsMiddleware())

	fileLogger := utils.NewFileLogger()

	s.Logger.Info("Attaching stdout logger")
	s.Router.Use(middlewares.Logger(s.Logger))
	s.Logger.Info("Attaching file logger")

	s.Router.Use(middlewares.Logger(fileLogger))
	s.Router.Use(middlewares.Recover)
	s.Router.Use(middlewares.RequestID)
}

func (s *Server) setupRoutes() {
	// @INFO serve static files
	s.Router.Handle("/*", public())

	// @INFO serve pages
	s.Router.Get("/", middlewares.HandleError(handlers.Index))

	// @INFO api endpoints
	s.Router.Post("/api/counter/increment", middlewares.HandleError(handlers.CounterIncrement))
	s.Router.Post("/api/counter/decrement", middlewares.HandleError(handlers.CounterDecrement))
	s.Router.Post("/api/counter/reset", middlewares.HandleError(handlers.CounterReset))
}

func (s *Server) Start() {
	s.setupMiddlewares()
	s.setupRoutes()

	s.Logger.Info("Starting server", "addr", s.Addr)
	if err := http.ListenAndServe(s.Addr, s.Router); err != nil {
		s.Logger.Error("Failed to start server", "error", err)
	}
}
