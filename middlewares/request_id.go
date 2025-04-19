package middlewares

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/google/uuid"
)

func RequestID(next http.Handler) http.Handler {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelError,
		AddSource: true,
	}))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Generate a unique request ID
		requestId, err := uuid.NewV7()
		if err != nil {
			logger.Error("Failed to generate request ID", "error", err)
			panic(err)
		}

		// Set the request ID in the response header
		w.Header().Set("X-Request-ID", requestId.String())

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
