package routes

import (
	"os"

	"github.com/kataras/iris"
	"github.com/spf13/viper"

	"github.com/djlechuck/recalbox-manager/utils/errors"
	"github.com/djlechuck/recalbox-manager/utils/screenshots"
)

// GetScreenshotsHandler handles the GET requests on /screenshots.
func GetScreenshotsHandler(ctx iris.Context) {
	hostname, err := os.Hostname()

	if err != nil {
		ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
		ctx.StatusCode(500)

		return
	}

	screenshotsPath := viper.GetString("recalbox.screenshotsPath")
	list, err := screenshots.ListImages(screenshotsPath)

	if err != nil {
		ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
		ctx.StatusCode(500)

		return
	}

	ctx.ViewData("PageTitle", ctx.Translate("Gestion des screenshots"))
	ctx.ViewData("Images", list)
	ctx.ViewData("Hostname", hostname)
	ctx.ViewData("Port", viper.GetInt("app.port"))

	ctx.View("views/screenshots.pug")
}

// GetScreenshotsDeleteHandler handles the GET requests on /screenshots/delete/:file.
func GetScreenshotsDeleteHandler(ctx iris.Context) {
	screenshotsPath := viper.GetString("recalbox.screenshotsPath")
	file := ctx.Params().Get("file")

	if file != "" {
		err := os.Remove(screenshotsPath + file)
		if err != nil {
			ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
			ctx.StatusCode(500)

			return
		}
	}

	ctx.Redirect("/screenshots")
}

// GetScreenshotsTakeHandler handles the GET requests on /screenshots/take.
func GetScreenshotsTakeHandler(ctx iris.Context) {
	err := screenshots.TakeScreenshot(viper.GetString("recalbox.screenshotsPath"))
	if err != nil {
		ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
		ctx.StatusCode(500)

		return
	}

	ctx.Redirect("/screenshots")
}
