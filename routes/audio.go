package routes

import (
	"github.com/kataras/iris"

	"github.com/djlechuck/recalbox-manager/store"
	"github.com/djlechuck/recalbox-manager/utils/recalbox"
)

// GetAudioHandler handles the GET requests on /audio.
func GetAudioHandler(ctx iris.Context) {
	sess := store.Sessions.Start(ctx)

	ctx.ViewData("PageTitle", ctx.Translate("Audio.Title"))
	ctx.ViewData("FormSended", sess.GetFlashString("formSended"))
	ctx.ViewData("Tr", map[string]interface{}{
		"Note":        ctx.Translate("Audio.Note"),
		"BgMusic":     ctx.Translate("Audio.BgMusic"),
		"Warning":     ctx.Translate("Audio.Warning"),
		"VolumeTitle": ctx.Translate("Volume du son"),
		"DeviceTitle": ctx.Translate("Sortie audio"),
		"DevicesList": map[string]string{
			"":          "-",
			"automatic": ctx.Translate("Automatique"),
			"hdmi":      ctx.Translate("Prise HDMI"),
			"jack":      ctx.Translate("Prise Jack"),
		},
		"BtnSave": ctx.Translate("BtnSave"),
	})

	ctx.View("views/audio.pug")
}

// PostAudioHandler handles the POST requests on /audio.
func PostAudioHandler(ctx iris.Context) {
	formData := iris.Map{}
	err := ctx.ReadForm(&formData)

	if err != nil {
		ctx.Values().Set("error", err)
		ctx.StatusCode(500)

		return
	}

	err = recalbox.ProcessRecalboxSettingsForm(formData, []string{"audio-bgmusic"})

	if err != nil {
		ctx.Values().Set("error", err)
		ctx.StatusCode(500)

		return
	}

	sess := store.Sessions.Start(ctx)
	sess.SetFlash("formSended", ctx.Translate("ConfigurationSaved"))

	ctx.Redirect("/audio", 303)
}
