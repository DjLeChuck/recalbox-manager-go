package routes

import (
	"github.com/kataras/iris"

	"github.com/djlechuck/recalbox-manager/store"
	"github.com/djlechuck/recalbox-manager/structs"
	"github.com/djlechuck/recalbox-manager/structs/forms"
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
		"RatioList": []structs.SelectList{
			{Value: "auto", Label: ctx.Translate("Automatique")},
			{Value: "4/3", Label: "4/3"},
			{Value: "16/9", Label: "16/9"},
			{Value: "16/10", Label: "16/10"},
			{Value: "custom", Label: ctx.Translate("Personnalisé")},
		},
		"ShadersetList": []structs.SelectList{
			{Value: "none", Label: ctx.Translate("Aucun")},
			{Value: "retro", Label: ctx.Translate("Retro")},
			{Value: "scanlines", Label: ctx.Translate("Scanlines")},
		},
	})

	ctx.View("views/systems.pug")
}

// PostSystemsHandler handles the POST requests on /systems.
func PostSystemsHandler(ctx iris.Context) {
	form := forms.Systems{}
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

	ctx.Redirect("/systems", 303)
}
