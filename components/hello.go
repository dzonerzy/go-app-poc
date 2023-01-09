package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// hello is a component that displays a simple "Hello World!". A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type Hello struct {
	app.Compo
}

func (h *Hello) Render() app.UI {
	// call script to get the name
	var txt = ""
	app.Window().JSValue().Get("go").Get("main").Get("App").Get("Greet").Invoke("World").Then(func(v app.Value) {
		// set the name to element with id "name"
		app.Window().GetElementByID("name").Set("innerHTML", v.String())
	})
	return app.Div().Body(
		app.H1().ID("name").Text(txt).Style("color", "white"),
	)
}

func NewHello() *Hello {
	return &Hello{}
}
