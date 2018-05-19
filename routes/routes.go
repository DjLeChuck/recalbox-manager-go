package routes

import (
	"github.com/kataras/iris"
)

// Configure registers the necessary routes to the app.
func Configure(app *iris.Application) {
	app.Get("/", GetHomeHandler).Name = "home"

	app.Get("/audio", GetAudioHandler).Name = "audio"
	app.Post("/audio", PostAudioHandler).Name = "audio_form"

	app.Get("/bios", GetBiosHandler).Name = "bios"
	app.Post("/upload/bios", iris.LimitRequestBodySize(256<<20), PostBiosUploadHandler).Name = "bios_form"
}
