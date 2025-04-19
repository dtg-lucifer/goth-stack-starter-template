package middlewares

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"runtime/debug"
)

// Recover is a middleware that recovers from panics
func Recover(next http.Handler) http.Handler {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelError,
		AddSource: true,
	}))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// Log the error and stack trace
				logger.Error("recovered from panic",
					"error", fmt.Sprintf("%v", err),
					"stack", string(debug.Stack()),
					"path", r.URL.Path,
					"method", r.Method,
				)

				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
