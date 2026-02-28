package middleware

import (
	"log/slog"
	"net/http"
)

func New(logger *slog.Logger) func(handler http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return requestLogging(logger, handler)
	}
}
