package server

import (
	"net/http"

	"will-steinleitner.de/internal/server/handlers"
)

func RegisterRoutes(
	home *handlers.Home,
	middleware func(http.Handler) http.Handler,
) http.Handler {
	//fs := http.FileServer(http.Dir("web/static")) // Serves all files from the ./web/static directory over HTTP.
	fs := http.FileServer(http.Dir("../web/static")) // Serves all files from the ./web/static directory over HTTP.

	mux := http.NewServeMux()

	//pipeline: request  -> logging -> home -> response
	mux.Handle(http.MethodGet+" /static/", http.StripPrefix("/static/", fs))
	mux.Handle(http.MethodGet+" /", middleware(home)) // Forces the implementation of ServeHTTP.

	return mux
}
