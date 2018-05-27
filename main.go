package main

import (
	"os"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/i18n"
	"github.com/kataras/iris/middleware/logger"

	"github.com/spf13/viper"

	"github.com/djlechuck/recalbox-manager/layouts"
	"github.com/djlechuck/recalbox-manager/routes"
	"github.com/djlechuck/recalbox-manager/structs"
	"github.com/djlechuck/recalbox-manager/utils/errors"
)

func main() {
	// Load configuration file
	viper.SetConfigName("recalbox")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./configs")

	err := viper.ReadInConfig() // Find and read the config file

	if err != nil { // Handle errors reading the config file
		panic(err)
	}

	if _, err := os.Stat("./configs/recalbox.dev.toml"); err == nil {
		viper.SetConfigName("recalbox.dev")
		err = viper.MergeInConfig()

		if err != nil { // Handle errors reading the config file
			panic(err)
		}
	}

	isDebug := viper.GetBool("app.debug")

	app := iris.New()
	app.Logger().SetLevel(viper.GetString("app.logLevel"))

	if !isDebug {
		f := errors.NewLogFile()

		defer f.Close()
		app.Logger().SetOutput(errors.NewLogFile())
	}

	if isDebug {
		app.Use(logger.New()) // log the requests to the terminal.
	}

	irisLanguages := make(map[string]string)
	languages := []structs.Language{}
	lErr := viper.UnmarshalKey("availableLanguages", &languages)

	if lErr != nil {
		panic(lErr)
	}

	for _, v := range languages {
		irisLanguages[v.Locale] = v.File
	}

	globalLocale := i18n.New(i18n.Config{
		Default:      "en-US",
		URLParameter: "lang",
		Languages:    irisLanguages})
	app.Use(globalLocale)

	app.StaticEmbeddedGzip("/static", "./assets", GzipAsset, GzipAssetNames)
	app.StaticWeb("/screenshots/viewer", viper.GetString("recalbox.screenshotsPath"))

	tmpl := iris.Pug("./templates", ".pug")
	tmpl.Binary(Asset, AssetNames)
	tmpl.Reload(isDebug) // reload templates on each request (development mode)

	app.RegisterView(tmpl)

	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewLayout(iris.NoLayout)

		err := ctx.Values().Get("error")
		errorMessage := "Something went wrong."

		if err != nil {
			errorMessage = err.(string)

			app.Logger().Error(err)
		}

		if ctx.IsAjax() {
			ctx.JSON(iris.Map{
				"success": false,
				"error":   errorMessage,
			})
		} else {
			ctx.ViewData("Is404", 404 == ctx.GetStatusCode())
			ctx.ViewData("ErrorMessage", errorMessage)
			ctx.View("layouts/error.pug")
		}
	})

	app.Configure(layouts.Configure, routes.Configure)
	app.Run(iris.Addr(":"+viper.GetString("app.port")), iris.WithoutServerError(iris.ErrServerClosed))
}
