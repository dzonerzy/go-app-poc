/*
 * go-app-poc
 *	To use with linux change the following lines:
 *	//go:generate GOOS=js GOARCH=wasm go build -o frontend/src/web/app.wasm -ldflags=-w -ldflags=-s ./app
 */

//go:generate go run ./app

package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/src
var assets embed.FS

func main() {

	// Create an instance of the app structure
	WailsApp := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "go-app-poc",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        WailsApp.startup,
		Bind: []interface{}{
			WailsApp,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
