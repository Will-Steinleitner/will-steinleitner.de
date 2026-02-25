package renderer

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const rendererTAG = "Renderer"

type Renderer struct {
	htmlCache map[string][]byte
}

func NewRenderer() *Renderer {
	log.Println(rendererTAG, ": building renderer..")
	htmlCache, err := newHTMLCache()
	if err != nil {
		log.Fatal(rendererTAG, err)
	}
	log.Println(rendererTAG, ": renderer built")

	return &Renderer{
		htmlCache: htmlCache,
	}
}
func newHTMLCache() (map[string][]byte, error) {
	cache := make(map[string][]byte)

	pages, err := filepath.Glob("../web/templates/*")
	if err != nil {
		log.Fatal(rendererTAG, err)
	}

	for _, page := range pages {
		name := filepath.Base(page)
		log.Println(rendererTAG, ": caching html -", name)
		content, err := os.ReadFile(page)
		if err != nil {
			log.Fatal(rendererTAG, err)
		}

		cache[name] = content
	}
	return cache, err
}
func (renderer *Renderer) RenderHTML(writer http.ResponseWriter, html string) {
	content, exists := renderer.htmlCache[html]
	if !exists {
		http.Error(writer, "html is missing", http.StatusInternalServerError)
		return
	}

	if _, err := writer.Write(content); err != nil {
		log.Println(rendererTAG, "write:", err)
	}
}
