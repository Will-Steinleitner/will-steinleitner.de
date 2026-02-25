package backend

import (
	"will-steinleitner.de/internal/handlers"
	"will-steinleitner.de/internal/renderer"
)

type Application struct {
	home *handlers.Home
}

func NewApplication() *Application {
	r := renderer.NewRenderer()

	return &Application{
		home: handlers.NewHome(r),
	}
}

func (app *Application) GetHome() *handlers.Home {
	return app.home
}
