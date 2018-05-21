package routes

import (
	"github.com/kataras/iris"

	"github.com/djlechuck/recalbox-manager/store"
	"github.com/djlechuck/recalbox-manager/utils/errors"
	"github.com/djlechuck/recalbox-manager/utils/recalbox"
)

// GetSystemsHandler handles the GET requests on /systems.
func GetSystemsHandler(ctx iris.Context) {
	sess := store.Sessions.Start(ctx)

	ctx.ViewData("PageTitle", ctx.Translate("Systems.Title"))
	ctx.ViewData("FormSended", sess.GetFlashString("formSended"))
	ctx.ViewData("Tr", iris.Map{
		"Note":    ctx.Translate("Controllers.Note"),
		"BtnSave": ctx.Translate("BtnSave"),
	})

	ctx.View("views/systems.pug")
}

// PostSystemsHandler handles the POST requests on /systems.
func PostSystemsHandler(ctx iris.Context) {
	formData := iris.Map{}
	err := ctx.ReadForm(&formData)

	if err != nil {
		ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
		ctx.StatusCode(500)

		return
	}

	err = recalbox.ProcessRecalboxSettingsForm(formData, []string{})

	if err != nil {
		ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
		ctx.StatusCode(500)

		return
	}

	sess := store.Sessions.Start(ctx)
	sess.SetFlash("formSended", ctx.Translate("ConfigurationSaved"))

	ctx.Redirect("/systems", 303)
}
