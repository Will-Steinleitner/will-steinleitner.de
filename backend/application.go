package backend

import (
	"log/slog"
	"net/http"
	"os"

	"will-steinleitner.de/cmd/server/config"
	"will-steinleitner.de/internal/renderer"
	"will-steinleitner.de/internal/server/middleware"
)

type Application struct {
	renderer   *renderer.Renderer
	cfg        *config.Config
	middleware func(handler http.Handler) http.Handler
}

func NewApplication() *Application {
	cfg := config.NewConfig()
	cfg.Load("config/config.yml")
	//cfg.Load("cmd/server/config/config.yml")

	r := renderer.NewRenderer()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	mw := middleware.New(logger)

	return &Application{
		renderer:   r,
		cfg:        cfg,
		middleware: mw,
	}
}

func (app *Application) GetRenderer() *renderer.Renderer {
	return app.renderer
}

func (app *Application) GetConfig() *config.Config {
	return app.cfg
}

func (app *Application) GetMiddleware() func(handler http.Handler) http.Handler {
	return app.middleware
}
