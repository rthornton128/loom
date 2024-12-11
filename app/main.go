package main

import (
	"github.com/rthornton128/loom/pkg/core"
)

type Application core.Application

func main() {
	app := (*Application)(core.NewApplication())
	app.Run()
}
