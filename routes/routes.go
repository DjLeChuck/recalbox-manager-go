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
	app.Get("/bios/check", GetBiosCheckHandler).Name = "bios_check"
	app.Post("/bios/upload", iris.LimitRequestBodySize(256<<20), PostBiosUploadHandler).Name = "bios_form"

	app.Get("/controllers", GetControllersHandler).Name = "controllers"
	app.Post("/controllers", PostControllersHandler).Name = "controllers_form"
}
