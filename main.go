package main

import (
	"os"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/i18n"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"

	"github.com/spf13/viper"

	"github.com/djlechuck/recalbox-manager/layouts"
	"github.com/djlechuck/recalbox-manager/routes"
	"github.com/djlechuck/recalbox-manager/structs"
)

func main() {
	// Load configuration file
	viper.SetConfigName("recalbox")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./configs")

	err := viper.ReadInConfig() // Find and read the config file

	if err != nil { // Handle errors reading the config file
		panic(err.Error())
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

	app.Use(recover.New()) // recover from any http-relative panics

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

	app.StaticWeb("/css", "./assets/css")
	app.StaticWeb("/js", "./assets/js")
	app.StaticWeb("/img", "./assets/img")

	tmpl := iris.Pug("./templates", ".pug")
	tmpl.Reload(isDebug) // reload templates on each request (development mode)

	app.RegisterView(tmpl)

	app.Favicon("./assets/favicon.png")

	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.View("views/error.pug")
	})

	app.Configure(iris.WithConfiguration(iris.TOML("./configs/iris.tml")), layouts.Configure, routes.Configure)
	app.Run(iris.Addr(":"+viper.GetString("app.port")), iris.WithoutServerError(iris.ErrServerClosed))
}
