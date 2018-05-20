package routes

import (
	"github.com/kataras/iris"

	"github.com/djlechuck/recalbox-manager/store"
	"github.com/djlechuck/recalbox-manager/utils/errors"
	"github.com/djlechuck/recalbox-manager/utils/recalbox"
)

// GetControllersHandler handles the GET requests on /controller.
func GetControllersHandler(ctx iris.Context) {
	sess := store.Sessions.Start(ctx)

	ctx.ViewData("PageTitle", ctx.Translate("Controllers.Title"))
	ctx.ViewData("FormSended", sess.GetFlashString("formSended"))

	ctx.ViewData("Tr", iris.Map{
		"Note":    ctx.Translate("Controllers.Note"),
		"BtnSave": ctx.Translate("BtnSave"),
		"Global": iris.Map{
			"Label": iris.Map{
				"Args": ctx.Translate("Global.Label.Args"),
			},
		},
		"Db9": iris.Map{
			"Title": ctx.Translate("Controllers.Db9.Title"),
			"Label": iris.Map{
				"Enabled": ctx.Translate("Controllers.Db9.Label.Enabled"),
			},
		},
		"Gamecon": iris.Map{
			"Title": ctx.Translate("Controllers.Gamecon.Title"),
			"Label": iris.Map{
				"Enabled": ctx.Translate("Controllers.Gamecon.Label.Enabled"),
				"Args":    ctx.Translate("Controllers.Gamecon.Label.Args"),
			},
		},
		"Gpio": iris.Map{
			"Title": ctx.Translate("Controllers.Gpio.Title"),
			"Label": iris.Map{
				"Enabled": ctx.Translate("Controllers.Gpio.Label.Enabled"),
				"Args":    ctx.Translate("Controllers.Gpio.Label.Args"),
			},
		},
		"Ps3": iris.Map{
			"Title": ctx.Translate("Controllers.Ps3.Title"),
			"DriversList": iris.Map{
				"":         "-",
				"official": ctx.Translate("Officiel"),
				"shanwan":  ctx.Translate("Shanwan"),
				"bluez":    ctx.Translate("Bluez 5"),
			},
			"Label": iris.Map{
				"Enabled": ctx.Translate("Controllers.Ps3.Label.Enabled"),
			},
		},
	})

	ctx.View("views/controllers.pug")
}

// PostControllersHandler handles the POST requests on /controllers.
func PostControllersHandler(ctx iris.Context) {
	formData := iris.Map{}
	err := ctx.ReadForm(&formData)

	if err != nil {
		ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
		ctx.StatusCode(500)

		return
	}

	err = recalbox.ProcessRecalboxSettingsForm(formData, []string{
		"controllers-db9-enabled",
		"controllers-gamecon-enabled",
		"controllers-gpio-enabled",
		"controllers-ps3-enabled",
	})

	if err != nil {
		ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
		ctx.StatusCode(500)

		return
	}

	sess := store.Sessions.Start(ctx)
	sess.SetFlash("formSended", ctx.Translate("ConfigurationSaved"))

	ctx.Redirect("/controllers", 303)
}
