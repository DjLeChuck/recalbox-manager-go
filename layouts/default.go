package layouts

import (
	"github.com/kataras/iris"

	"github.com/spf13/viper"
)

// MenuItem represents an entry of the menu.
type MenuItem struct {
	Link, Glyph, Label string
	Children           []MenuItem
}

// New returns a new handler which adds some headers and view data
// describing the application, i.e the owner, the startup time.
func New(app *iris.Application) iris.Handler {
	return func(ctx iris.Context) {
		menuEntries := []MenuItem{
			{Link: app.GetRoute("home").FormattedPath, Glyph: "home", Label: ctx.Translate("Accueil")},
			{Link: "/monitoring", Glyph: "signal", Label: ctx.Translate("Monitoring")},
			{Link: app.GetRoute("audio").FormattedPath, Glyph: "volume-up", Label: ctx.Translate("Audio")},
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
		ctx.ViewData("AvailableLang", viper.GetStringMapString("availableLanguages"))

		ctx.Next()
	}
}

// Configure creates a new layout middleware and registers that to the app.
func Configure(app *iris.Application) {
	h := New(app)
	app.Use(h)
}
