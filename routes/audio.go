package routes

import (
	"github.com/kataras/iris"

	"github.com/djlechuck/recalbox-manager/store"
	"github.com/djlechuck/recalbox-manager/structs"
	"github.com/djlechuck/recalbox-manager/structs/forms"
	"github.com/djlechuck/recalbox-manager/utils/errors"
	"github.com/djlechuck/recalbox-manager/utils/recalbox"
)

// GetAudioHandler handles the GET requests on /audio.
func GetAudioHandler(ctx iris.Context) {
	sess := store.Sessions.Start(ctx)

	ctx.ViewData("PageTitle", ctx.Translate("Audio.Title"))
	ctx.ViewData("FormSended", sess.GetFlashString("formSended"))
	ctx.ViewData("Tr", iris.Map{
		"Note":        ctx.Translate("Audio.Note"),
		"BgMusic":     ctx.Translate("Audio.BgMusic"),
		"Warning":     ctx.Translate("Audio.Warning"),
		"VolumeTitle": ctx.Translate("Volume du son"),
		"DeviceTitle": ctx.Translate("Sortie audio"),
		"DevicesList": []structs.SelectList{
			{Value: "automatic", Label: ctx.Translate("Automatique")},
			{Value: "hdmi", Label: ctx.Translate("Prise HDMI")},
			{Value: "jack", Label: ctx.Translate("Prise Jack")},
		},
		"BtnSave": ctx.Translate("BtnSave"),
	})

	ctx.View("views/audio.pug")
}

// PostAudioHandler handles the POST requests on /audio.
func PostAudioHandler(ctx iris.Context) {
	form := forms.Audio{}
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

	ctx.Redirect("/audio", 303)
}
