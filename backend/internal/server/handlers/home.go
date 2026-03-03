package handlers

import (
	"net/http"

	"will-steinleitner.de/internal/renderer"
)

func HandleHome(r *renderer.Renderer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		if req.Method == http.MethodGet && req.URL.Path == "/" {
			r.RenderHTML(w, "index.html", struct {
				Title string
			}{
				Title: "Home",
			})
			return
		}

		http.NotFound(w, req)
	})
}
