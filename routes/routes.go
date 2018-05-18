package routes

import (
	"github.com/kataras/iris"
)

// Configure registers the necessary routes to the app.
func Configure(app *iris.Application) {
	app.Get("/", GetHomeHandler).Name = "home"

	app.Get("/audio", GetAudioHandler).Name = "audio"
	app.Post("/audio", PostAudioHandler)
}
