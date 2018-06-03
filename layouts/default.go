package layouts

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/kataras/iris"

	"github.com/spf13/viper"

	"github.com/djlechuck/recalbox-manager/store"
	"github.com/djlechuck/recalbox-manager/structs"
	"github.com/djlechuck/recalbox-manager/utils/errors"
)

// New returns a new handler which adds some view data.
func New(app *iris.Application) iris.Handler {
	// Check if authentication is required
	var credentials structs.Credentials
	path := viper.GetString("auth.path") + viper.GetString("auth.file")
	authentication := &structs.Authentication{
		Enabled:         true,
		Credentials:     credentials,
		IsAuthenticated: false,
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		authentication.Enabled = false
	} else {
		content, err := ioutil.ReadFile(path)
		if err == nil {
			if err := json.Unmarshal(content, &credentials); err != nil {
				panic(err)
			}

			authentication.Credentials = credentials
		}
	}

	return func(ctx iris.Context) {
		currentPath := ctx.Path()
		menuEntries := []structs.MenuItem{
			{
				Link:     app.GetRoute("home").FormattedPath,
				Glyph:    "home",
				Label:    ctx.Translate("Accueil"),
				IsActive: app.GetRoute("home").FormattedPath == currentPath,
			},
			{
				Link:     app.GetRoute("monitoring").FormattedPath,
				Glyph:    "signal",
				Label:    ctx.Translate("Monitoring"),
				IsActive: app.GetRoute("monitoring").FormattedPath == currentPath,
			},
			{
				Link:     app.GetRoute("audio").FormattedPath,
				Glyph:    "volume-up",
				Label:    ctx.Translate("Audio"),
				IsActive: app.GetRoute("audio").FormattedPath == currentPath,
			},
			{
				Link:     app.GetRoute("bios").FormattedPath,
				Glyph:    "compact-disc",
				Label:    ctx.Translate("BIOS"),
				IsActive: app.GetRoute("bios").FormattedPath == currentPath,
			},
			{
				Link:     app.GetRoute("controllers").FormattedPath,
				Glyph:    "gamepad",
				Label:    ctx.Translate("Contrôleurs"),
				IsActive: app.GetRoute("controllers").FormattedPath == currentPath,
			},
			{
				Link:     app.GetRoute("systems").FormattedPath,
				Glyph:    "hdd",
				Label:    ctx.Translate("Systèmes"),
				IsActive: app.GetRoute("systems").FormattedPath == currentPath,
			},
			{
				Link:     app.GetRoute("configuration").FormattedPath,
				Glyph:    "cog",
				Label:    ctx.Translate("Configuration"),
				IsActive: app.GetRoute("configuration").FormattedPath == currentPath,
			},
			{
				Link:     "/roms",
				Glyph:    "save",
				Label:    ctx.Translate("ROMs"),
				IsActive: false,
			},
			{
				Link:     app.GetRoute("screenshots").FormattedPath,
				Glyph:    "images",
				Label:    ctx.Translate("Screenshots"),
				IsActive: app.GetRoute("screenshots").FormattedPath == currentPath,
			},
			{
				Link:  app.GetRoute("help").FormattedPath,
				Glyph: "question-circle",
				Label: ctx.Translate("Dépannage"),
				Children: []structs.MenuItem{
					{
						Link:     app.GetRoute("logs").FormattedPath,
						Glyph:    "file",
						Label:    ctx.Translate("Logs"),
						IsActive: app.GetRoute("logs").FormattedPath == currentPath,
					},
					{
						Link:     app.GetRoute("recalbox-conf").FormattedPath,
						Glyph:    "file",
						Label:    "recalbox.conf",
						IsActive: app.GetRoute("recalbox-conf").FormattedPath == currentPath,
					},
					{
						Link:     app.GetRoute("help").FormattedPath,
						Glyph:    "question-circle",
						Label:    ctx.Translate("Dépannage"),
						IsActive: app.GetRoute("help").FormattedPath == currentPath,
					},
				},
			},
			{
				Link:     app.GetRoute("security").FormattedPath,
				Glyph:    "lock",
				Label:    ctx.Translate("Sécurité"),
				IsActive: app.GetRoute("security").FormattedPath == currentPath,
			},
		}

		for k, v := range menuEntries {
			if 0 < len(v.Children) {
				menuEntries[k].ActiveClass = " dropdown"
			}

			for ck, cv := range v.Children {
				if cv.IsActive {
					menuEntries[k].ActiveClass = menuEntries[k].ActiveClass + " active"
					menuEntries[k].Children[ck].ActiveClass = "active"
				}
			}

			if v.IsActive {
				menuEntries[k].ActiveClass = menuEntries[k].ActiveClass + " active"
			}
		}

		var menuLanguages []structs.AvailableLanguage
		var languages []structs.Language
		err := viper.UnmarshalKey("availableLanguages", &languages)

		if err != nil {
			ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
			ctx.StatusCode(500)

			return
		}

		locale := ctx.Values().GetString(ctx.Application().ConfigurationReadOnly().GetTranslateLanguageContextKey())
		currentLang := ""

		for _, v := range languages {
			menuLanguages = append(menuLanguages, structs.AvailableLanguage{
				Locale: v.Locale,
				Name:   v.Name,
			})

			if v.Locale == locale {
				currentLang = v.Name
			}
		}

		ctx.ViewData("RecalboxManagerTitle", ctx.Translate("Recalbox Manager"))
		ctx.ViewData("MenuEntries", menuEntries)
		ctx.ViewData("CurrentLang", currentLang)
		ctx.ViewData("AvailableLang", menuLanguages)
		ctx.ViewData("AppVersion", viper.GetString("app.version"))
		ctx.ViewData("NeedAuth", authentication.Enabled)

		ctx.Gzip(true)

		sess := store.Sessions.Start(ctx)
		sess.Set("authentication", authentication)

		ctx.Next()
	}
}

// Configure creates a new layout middleware and registers itself to the app.
func Configure(app *iris.Application) {
	h := New(app)
	app.Use(h)
}
