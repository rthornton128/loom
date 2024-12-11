package main

import (
	"net/http"

	"github.com/rthornton128/loom/app/config"
	"github.com/rthornton128/loom/pkg/core"
)

func (a *Application) Run() {
	router := (*core.Application)(a).Router()
	config.Draw(router)

	http.ListenAndServe(":8080", router)
}
