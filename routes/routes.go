package routes

import (
	"github.com/kataras/iris"

	"github.com/djlechuck/recalbox-manager/routes/middlewares"
)

// Configure registers the necessary routes to the app.
func Configure(app *iris.Application) {
	app.Get("/login", GetLoginHandler).Name = "login"
	app.Post("/login", PostLoginHandler).Name = "login_form"

	app.PartyFunc("/", func(manager iris.Party) {
		manager.Use(middlewares.CheckAuthMiddlewareHandler)

		manager.Get("/", GetHomeHandler).Name = "home"

		manager.Get("/logout", GetLogoutHandler).Name = "logout"

		manager.Get("/audio", GetAudioHandler).Name = "audio"
		manager.Post("/audio", PostAudioHandler).Name = "audio_form"

		manager.Get("/bios", GetBiosHandler).Name = "bios"
		manager.Get("/bios/check", GetBiosCheckHandler).Name = "bios_check"
		manager.Post("/bios/upload", iris.LimitRequestBodySize(256<<20), PostBiosUploadHandler).Name = "bios_form"
		manager.Get("/bios/delete/{file:string}", GetBiosDeleteHandler).Name = "bios_delete"

		manager.Get("/controllers", GetControllersHandler).Name = "controllers"
		manager.Post("/controllers", PostControllersHandler).Name = "controllers_form"

		manager.Get("/systems", GetSystemsHandler).Name = "systems"
		manager.Post("/systems", PostSystemsHandler).Name = "systems_form"

		manager.Get("/configuration", GetConfigurationHandler).Name = "configuration"
		manager.Post("/configuration", PostConfigurationHandler).Name = "configuration_form"

		manager.Get("/screenshots", GetScreenshotsHandler).Name = "screenshots"
		manager.Get("/screenshots/delete/{file:string}", GetScreenshotsDeleteHandler).Name = "screenshots_delete"
		manager.Get("/screenshots/take", GetScreenshotsTakeHandler).Name = "screenshots_take"

		manager.Get("/logs", GetLogsHandler).Name = "logs"
		manager.Post("/logs", PostLogsHandler).Name = "logs_form"

		manager.Get("/recalbox-conf", GetRecalboxConfHandler).Name = "recalbox-conf"
		manager.Post("/recalbox-conf", PostRecalboxConfHandler).Name = "recalbox-conf_form"

		manager.Get("/help", GetHelpHandler).Name = "help"

		manager.Get("/help/recalbox-support", GetLaunchRecalboxSupportHandler).Name = "launch-recalbox-support"

		manager.Get("/os/reboot", GetRebootOsHandler).Name = "os-reboot"
		manager.Get("/os/shutdown", GetShutdownOsHandler).Name = "os-shutdown"

		manager.Get("/es/action/{name:string}", GetActionEsHandler).Name = "es-action"
		manager.Get("/es/status", GetStatusEsHandler).Name = "es-status"

		manager.Get("/monitoring", GetMonitoringHandler).Name = "monitoring"

		manager.Get("/security", GetSecurityHandler).Name = "security"
		manager.Post("/security", PostSecurityHandler).Name = "security_form"
	})
}
