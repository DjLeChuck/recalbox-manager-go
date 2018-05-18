package routes

import (
	"github.com/kataras/iris"
)

// Configure registers the necessary routes to the app.
func Configure(app *iris.Application) {
	app.Get("/", GetHomeHandler)

	app.Get("/audio", GetAudioHandler)
	app.Post("/audio", PostAudioHandler)
}
