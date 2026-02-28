package backend

import (
	"will-steinleitner.de/cmd/server/config"
	"will-steinleitner.de/internal/handlers"
	"will-steinleitner.de/internal/renderer"
)

type Application struct {
	home *handlers.Home
	cfg  *config.Config
}

func NewApplication() *Application {
	cfg := config.NewConfig()
	cfg.Load("config/config.yml")
	//cfg.Load("cmd/server/config/config.yml")

	r := renderer.NewRenderer()

	return &Application{
		home: handlers.NewHome(r),
		cfg:  cfg,
	}
}

func (app *Application) GetHome() *handlers.Home {
	return app.home
}

func (app *Application) GetConfig() *config.Config {
	return app.cfg
}
