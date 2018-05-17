package main

import (
	"os"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/i18n"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"

	"github.com/spf13/viper"
)

// MenuItem represent an entry of the menu.
type MenuItem struct {
	Link, Glyph, Label string
	Children           []MenuItem
}

// HomeTile represent a tile on the homepage.
type HomeTile struct {
	Link, Image, Label string
}

func main() {
	// Load configuration file
	viper.SetConfigName("recalbox")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./configs")

	err := viper.ReadInConfig() // Find and read the config file

	if err != nil { // Handle errors reading the config file
		panic(err.Error())
	}

	isDebug := viper.GetBool("app.debug")

	app := iris.New()
	app.Logger().SetLevel(viper.GetString("app.logLevel"))

	app.Use(recover.New()) // recover from any http-relative panics

	if isDebug {
		app.Use(logger.New()) // log the requests to the terminal.
	}

	availableLanguages := map[string]string{
		"en-US": "./locales/en.ini",
		"fr-FR": "./locales/fr.ini",
		"zh-CN": "./locales/zh.ini",
	}

	globalLocale := i18n.New(i18n.Config{
		Default:      "en-US",
		URLParameter: "lang",
		Languages:    availableLanguages})
	app.Use(globalLocale)

	app.StaticWeb("/css", "./assets/css")
	app.StaticWeb("/js", "./assets/js")
	app.StaticWeb("/img", "./assets/img")

	tmpl := iris.Pug("./templates", ".pug")
	tmpl.Reload(isDebug) // reload templates on each request (development mode)

	app.RegisterView(tmpl)

	app.Use(func(ctx iris.Context) {
		menuEntries := []MenuItem{
			{Link: "/", Glyph: "home", Label: ctx.Translate("Accueil")},
			{Link: "/monitoring", Glyph: "signal", Label: ctx.Translate("Monitoring")},
			{Link: "/audio", Glyph: "volume-up", Label: ctx.Translate("Audio")},
			{Link: "/bios", Glyph: "cd", Label: ctx.Translate("BIOS")},
			{Link: "/controllers", Glyph: "phone", Label: ctx.Translate("Contrôleurs")},
			{Link: "/systems", Glyph: "hdd", Label: ctx.Translate("Systèmes")},
			{Link: "/configuration", Glyph: "cog", Label: ctx.Translate("Configuration")},
			{Link: "/roms", Glyph: "floppy-disk", Label: ctx.Translate("ROMs")},
			{Link: "/screenshots", Glyph: "picture", Label: ctx.Translate("Screenshots")},
			{Link: "/help", Glyph: "question-sign", Label: ctx.Translate("Dépannage"), Children: []MenuItem{
				{Link: "/logs", Glyph: "file", Label: ctx.Translate("Logs")},
				{Link: "/recalbox-conf", Glyph: "file", Label: "recalbox.conf"},
				{Link: "/help", Glyph: "question-sign", Label: ctx.Translate("Dépannage")},
			}},
		}

		ctx.ViewLayout("layouts/default.pug")
		ctx.ViewData("RecalboxManagerTitle", ctx.Translate("Recalbox Manager"))
		ctx.ViewData("MenuEntries", menuEntries)
		ctx.ViewData("CurrentLang", ctx.Values().GetString(ctx.Application().ConfigurationReadOnly().GetTranslateLanguageContextKey()))
		ctx.ViewData("AvailableLang", availableLanguages)

		ctx.Next()
	})

	app.Handle("GET", "/", func(ctx iris.Context) {
		hostname, err := os.Hostname()

		if err != nil {
			panic(err)
		}

		ctx.ViewData("PageTitle", ctx.Translate("Accueil"))
		ctx.ViewData("Tiles", []HomeTile{
			{"//" + hostname + ":8080/gamepad.html?analog", "/img/gamepad.png", ctx.Translate("Utiliser le gamepad virtuel")},
			{"//" + hostname + ":8080/keyboard.html", "/img/keyboard.png", ctx.Translate("Utiliser le clavier virtuel")},
			{"//" + hostname + ":8080/touchepad.html", "/img/touchpad.png", ctx.Translate("Utiliser le touchpad virtuel")},
		})

		ctx.View("views/home.pug")
	})

	app.Handle("GET", "/audio", func(ctx iris.Context) {
		ctx.ViewData("PageTitle", ctx.Translate("Audio"))
		ctx.ViewData("Tr", map[string]string{
			"Note":         ctx.Translate("SoundNote"),
			"SoundTitle":   ctx.Translate("SoundTitle"),
			"SoundWarning": ctx.Translate("SoundWarning"),
			"VolumeTitle":  ctx.Translate("Volume du son"),
			"DeviceTitle":  ctx.Translate("Sortie audio"),
			"BtnSave":      ctx.Translate("BtnSave"),
		})

		ctx.View("views/audio.pug")
	})

	app.Configure(iris.WithConfiguration(iris.TOML("./configs/iris.tml")))
	app.Run(iris.Addr(":8888"), iris.WithoutServerError(iris.ErrServerClosed))
}
