package handlers

import (
	"net/http"

	"will-steinleitner.de/internal/renderer"
)

type Home struct {
	renderer *renderer.Renderer
}

func NewHome(renderer *renderer.Renderer) *Home {
	return &Home{
		renderer,
	}
}

func (h Home) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	switch {
	case request.Method == http.MethodGet && request.URL.Path == "/":
		h.renderer.RenderHTML(writer, "index.html")
		return
	}
}
