package server

import (
	"net/http"

	"will-steinleitner.de/internal/renderer"
	"will-steinleitner.de/internal/server/handlers"
)

func RegisterRoutes(
	renderer *renderer.Renderer,
	middleware func(http.Handler) http.Handler,
) http.Handler {
	fs := http.FileServer(http.Dir("web/static")) // Serves all files from the ./web/static directory over HTTP.
	//fs := http.FileServer(http.Dir("../web/static")) // Serves all files from the ./web/static directory over HTTP.
	mux := http.NewServeMux()

	//pipeline: request  -> logging -> home -> response
	mux.Handle(http.MethodGet+" /static/", http.StripPrefix("/static/", fs))
	mux.Handle(http.MethodGet+" /", middleware(handlers.HandleHome(renderer))) // Forces the implementation of ServeHTTP.
	mux.Handle(http.MethodGet+" /healthz", middleware(handlers.HandleHealthz()))

	return mux
}
