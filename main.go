package main

import (
	"github.com/kataras/iris"

	"github.com/kataras/iris/middleware/i18n"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

type MenuItem struct {
	Link, Glyph, Label string
}

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	app.Use(recover.New()) // recover from any http-relative panics
	app.Use(logger.New())  // log the requests to the terminal.

	availableLanguages := map[string]string{
		"en-US": "./locales/en.ini",
		"fr-FR": "./locales/fr.ini",
		"zh-CN": "./locales/zh.ini",
	}

	globalLocale := i18n.New(i18n.Config{
		Default:      "zh-CN",
		URLParameter: "lang",
		Languages:    availableLanguages})
	app.Use(globalLocale)

	app.StaticWeb("/css", "./assets/css")
	app.StaticWeb("/js", "./assets/js")

	tmpl := iris.Pug("./templates", ".pug")
	tmpl.Reload(true) // reload templates on each request (development mode)

	app.RegisterView(tmpl)

	app.Use(func(ctx iris.Context) {
		menuEntries := []MenuItem{
			{"/monitoring", "signal", ctx.Translate("Monitoring")},
			{"/audio", "volume-up", ctx.Translate("Audio")},
			{"/bios", "cd", ctx.Translate("BIOS")},
			{"/controllers", "phone", ctx.Translate("Contrôleurs")},
			{"/systems", "hdd", ctx.Translate("Systèmes")},
			{"/configuration", "cog", ctx.Translate("Configuration")},
			{"/roms", "floppy-disk", ctx.Translate("ROMs")},
			{"/screenshots", "picture", ctx.Translate("Screenshots")},

			// const secondMenuEntries = [{
			//   link: '/logs',
			//   glyph: 'file',
			//   label: t('Logs'),
			// }, {
			//   link: '/recalbox-conf',
			//   glyph: 'file',
			//   label: 'recalbox.conf',
			// }, {
			//   link: '/help',
			//   glyph: 'question-sign',
			//   label: t('Dépannage'),
			// }];
		}

		ctx.ViewLayout("layouts/default.pug")
		ctx.ViewData("RecalboxManagerTitle", ctx.Translate("Recalbox Manager"))
		ctx.ViewData("MenuEntries", menuEntries)
		ctx.ViewData("CurrentLang", ctx.Values().GetString(ctx.Application().ConfigurationReadOnly().GetTranslateLanguageContextKey()))
		ctx.ViewData("AvailableLang", availableLanguages)

		ctx.Next()
	})

	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.View("views/home.pug")
	})

	app.Run(iris.Addr(":8888"), iris.WithoutServerError(iris.ErrServerClosed))
}
