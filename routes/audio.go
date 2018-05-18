package routes

import (
	"github.com/kataras/iris"
)

// AudioForm represents the form Audio.
type AudioForm struct {
	AudioBgmusic bool   `form:"audio-bgmusic"`
	AudioVolume  int    `form:"audio-volume"`
	AudioDevice  string `form:"audio-device"`
}

// GetAudioHandler handles the GET requests on /audio.
func GetAudioHandler(ctx iris.Context) {
	ctx.ViewData("PageTitle", ctx.Translate("Audio"))
	ctx.ViewData("Tr", map[string]interface{}{
		"Note":         ctx.Translate("SoundNote"),
		"SoundTitle":   ctx.Translate("SoundTitle"),
		"SoundWarning": ctx.Translate("SoundWarning"),
		"VolumeTitle":  ctx.Translate("Volume du son"),
		"DeviceTitle":  ctx.Translate("Sortie audio"),
		"DevicesList": map[string]string{
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
	formData := AudioForm{}
	err := ctx.ReadForm(&formData)

	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
	}
}
