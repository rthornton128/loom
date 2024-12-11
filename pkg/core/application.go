package core

import "github.com/rthornton128/loom/pkg/routing"

type Application struct {
	router *routing.Router
}

func NewApplication() *Application {
	return &Application{
		router: routing.NewRouter(),
	}
}

func (a *Application) Router() *routing.Router {
	return a.router
}
