package routes

import (
	"github.com/kataras/iris"

	"github.com/spf13/viper"

	"github.com/djlechuck/recalbox-manager/store"
	"github.com/djlechuck/recalbox-manager/structs/forms"
	"github.com/djlechuck/recalbox-manager/utils/errors"
	"github.com/djlechuck/recalbox-manager/utils/recalbox"
)

// GetConfigurationHandler handles the GET requests on /configuration.
func GetConfigurationHandler(ctx iris.Context) {
	sess := store.Sessions.Start(ctx)

	ctx.ViewData("PageTitle", ctx.Translate("Configuration"))
	ctx.ViewData("FormSended", sess.GetFlashString("formSended"))

	ctx.ViewData("SelectData", iris.Map{
		"EsMenus":              viper.Get("recalbox.esMenus"),
		"EmulatorsSpecialkeys": viper.Get("recalbox.emulatorsSpecialkeys"),
		"KeyboardLayouts":      viper.Get("recalbox.keyboardlayouts"),
		"SystemLocales":        viper.Get("recalbox.systemlocales"),
		"Timezones":            viper.Get("recalbox.timezones"),
		"UpdatesTypes": iris.Map{
			"stable":   ctx.Translate("Stable"),
			"beta":     ctx.Translate("Bêta"),
			"unstable": ctx.Translate("Instable"),
		},
	})

	ctx.ViewData("Tr", iris.Map{
		"Note":    ctx.Translate("Configuration.Note"),
		"BtnSave": ctx.Translate("BtnSave"),
	})

	ctx.View("views/configuration.pug")
}

// PostConfigurationHandler handles the POST requests on /configuration.
func PostConfigurationHandler(ctx iris.Context) {
	form := forms.Configuration{}
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

	ctx.Redirect("/configuration", 303)
}
