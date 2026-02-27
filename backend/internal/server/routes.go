package server

import (
	"net/http"

	"will-steinleitner.de/internal/handlers"
)

func RegisterRoutes(home *handlers.Home) http.Handler {
	fs := http.FileServer(http.Dir("../web/static")) // Serves all files from the ./web/static directory over HTTP.

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.Handle("/", home) // Forces the implementation of ServeHTTP.

	return mux
}
