package main

import (
	"changeme/components"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func main() {
	app.Route("/", components.NewHello())
	app.RunWhenOnBrowser()
	h := &app.Handler{
		Name:        "go-app-poc",
		Title:       "go-app-poc",
		Description: "go-app-poc",
	}
	app.GenerateStaticWebsite("frontend/src", h)
}
