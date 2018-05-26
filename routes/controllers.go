package routes

import (
	"github.com/kataras/iris"

	"github.com/djlechuck/recalbox-manager/store"
	"github.com/djlechuck/recalbox-manager/structs"
	"github.com/djlechuck/recalbox-manager/structs/forms"
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
			},
		},
		"Gpio": iris.Map{
			"Title": ctx.Translate("Controllers.Gpio.Title"),
			"Label": iris.Map{
				"Enabled": ctx.Translate("Controllers.Gpio.Label.Enabled"),
			},
		},
		"Ps3": iris.Map{
			"Title": ctx.Translate("Controllers.Ps3.Title"),
			"DriversList": []structs.SelectList{
				{Value: "official", Label: ctx.Translate("Officiel")},
				{Value: "shanwan", Label: ctx.Translate("Shanwan")},
				{Value: "bluez", Label: ctx.Translate("Bluez 5")},
			},
			"Label": iris.Map{
				"Enabled":     ctx.Translate("Controllers.Ps3.Label.Enabled"),
				"DriverToUse": ctx.Translate("Controllers.Ps3.Label.DriverToUse"),
			},
		},
	})

	ctx.View("views/controllers.pug")
}

// PostControllersHandler handles the POST requests on /controllers.
func PostControllersHandler(ctx iris.Context) {
	form := forms.Controllers{}
	err := ctx.ReadForm(&form)

	if err != nil {
		ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
		ctx.StatusCode(500)

		return
	}

	data := recalbox.FormatFormData(&form)
	err = recalbox.ProcessRecalboxSettingsForm(data)

	if err != nil {
		ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
		ctx.StatusCode(500)

		return
	}

	sess := store.Sessions.Start(ctx)
	sess.SetFlash("formSended", ctx.Translate("ConfigurationSaved"))

	ctx.Redirect("/controllers", 303)
}
