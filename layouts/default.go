package layouts

import (
	"github.com/kataras/iris"

	"github.com/spf13/viper"

	"github.com/djlechuck/recalbox-manager/structs"
)

// New returns a new handler which adds some headers and view data
// describing the application, i.e the owner, the startup time.
func New(app *iris.Application) iris.Handler {
	return func(ctx iris.Context) {
		menuEntries := []structs.MenuItem{
			{
				Link:  app.GetRoute("home").FormattedPath,
				Glyph: "home",
				Label: ctx.Translate("Accueil"),
			},
			{
				Link:  "/monitoring",
				Glyph: "signal",
				Label: ctx.Translate("Monitoring"),
			},
			{
				Link:  app.GetRoute("audio").FormattedPath,
				Glyph: "volume-up",
				Label: ctx.Translate("Audio"),
			},
			{
				Link:  app.GetRoute("bios").FormattedPath,
				Glyph: "compact-disc",
				Label: ctx.Translate("BIOS"),
			},
			{
				Link:  app.GetRoute("controllers").FormattedPath,
				Glyph: "gamepad",
				Label: ctx.Translate("Contrôleurs"),
			},
			{
				Link:  "/systems",
				Glyph: "hdd",
				Label: ctx.Translate("Systèmes"),
			},
			{
				Link:  "/configuration",
				Glyph: "cog",
				Label: ctx.Translate("Configuration"),
			},
			{
				Link:  "/roms",
				Glyph: "save",
				Label: ctx.Translate("ROMs"),
			},
			{
				Link:  "/screenshots",
				Glyph: "images",
				Label: ctx.Translate("Screenshots"),
			},
			{
				Link:  "/help",
				Glyph: "question-circle",
				Label: ctx.Translate("Dépannage"), Children: []structs.MenuItem{
					{
						Link:  "/logs",
						Glyph: "file",
						Label: ctx.Translate("Logs"),
					},
					{
						Link:  "/recalbox-conf",
						Glyph: "file",
						Label: "recalbox.conf"},
					{
						Link:  "/help",
						Glyph: "question-circle",
						Label: ctx.Translate("Dépannage"),
					},
				}},
		}

		menuLanguages := make(map[string]string)
		languages := []structs.Language{}
		err := viper.UnmarshalKey("availableLanguages", &languages)

		if err != nil {
			ctx.Values().Set("errorMessage", err.Error())
			ctx.StatusCode(500)

			return
		}

		for _, v := range languages {
			menuLanguages[v.Locale] = v.Name
		}

		currentLang := menuLanguages[ctx.Values().GetString(ctx.Application().ConfigurationReadOnly().GetTranslateLanguageContextKey())]

		ctx.ViewLayout("layouts/default.pug")
		ctx.ViewData("RecalboxManagerTitle", ctx.Translate("Recalbox Manager"))
		ctx.ViewData("MenuEntries", menuEntries)
		ctx.ViewData("CurrentLang", currentLang)
		ctx.ViewData("AvailableLang", menuLanguages)
		ctx.ViewData("AppVersion", viper.GetString("app.version"))

		ctx.Next()
	}
}

// Configure creates a new layout middleware and registers that to the app.
func Configure(app *iris.Application) {
	h := New(app)
	app.Use(h)
}
