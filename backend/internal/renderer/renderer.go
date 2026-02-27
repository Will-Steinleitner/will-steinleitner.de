package renderer

import (
	"bytes"
	"html/template"
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

	partials, err := filepath.Glob("../web/partials/*")
	if err != nil {
		log.Fatal(err)
	}

	for _, page := range partials {
		name := "partials/" + filepath.Base(page) // => "partials/header.html"
		log.Println(rendererTAG, ": caching partial html -", name)

		content, err := os.ReadFile(page)
		if err != nil {
			log.Fatal(rendererTAG, err)
		}
		cache[name] = content
	}
	return cache, err
}

func (r *Renderer) RenderHTML(w http.ResponseWriter, html string, data interface{}) {
	content, ok := r.htmlCache[html]
	if !ok {
		http.Error(w, "html is missing: "+html, 500)
		return
	}

	tpl, err := template.New(html).
		Funcs(template.FuncMap{"partial": r.partial}).
		Parse(string(content))
	if err != nil {
		http.Error(w, "template parse error: "+err.Error(), 500)
		return
	}

	var buf bytes.Buffer
	if err := tpl.Execute(&buf, data); err != nil {
		http.Error(w, "template execute error: "+err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write(buf.Bytes())
}

func (r *Renderer) partial(name string, data interface{}) template.HTML {
	content, ok := r.htmlCache[name]

	if !ok {
		return template.HTML("partial not found: " + name)
	}

	tpl, err := template.New(name).Parse(string(content))
	if err != nil {
		return template.HTML("partial parse error: " + err.Error())
	}

	var buf bytes.Buffer
	if err := tpl.Execute(&buf, data); err != nil {
		return template.HTML("partial execute error: " + err.Error())
	}

	return template.HTML(buf.String())
}
