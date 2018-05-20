package routes

import (
	"os/exec"
	"strings"

	"github.com/kataras/iris"

	"github.com/spf13/viper"

	"github.com/djlechuck/recalbox-manager/store"
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

	err = ProcessRecalboxSettingsForm(formData, []string{"audio-bgmusic"})

	if err != nil {
		ctx.Values().Set("error", err)
		ctx.StatusCode(500)

		return
	}

	sess := store.Sessions.Start(ctx)
	sess.SetFlash("formSended", ctx.Translate("ConfigurationSaved"))

	ctx.Redirect("/audio", 303)
}

// ProcessRecalboxSettingsForm loops through the form data and update the
// config of Recalbox by calling the config script.
// checkboxes represents checkboxes on the page. When submitted uncheckek,
// they have no value, so we force one.
func ProcessRecalboxSettingsForm(data iris.Map, checkboxes []string) (err error) {
	for _, c := range checkboxes {
		if _, ok := data[c]; !ok {
			data[c] = "0"
		}
	}

	pythonFile := viper.GetString("recalbox.pythonSettingsFile")

	for k, v := range data {
		normalizedKey := strings.Replace(k, "-", ".", -1)
		_, err := exec.Command("python", pythonFile, "-command", "save", "-key", normalizedKey, "-value", v.(string)).CombinedOutput()

		if err != nil {
			return err
		}
	}

	configScript := viper.GetString("recalbox.configScript")
	_, err = exec.Command(configScript, "volume", data["audio-volume"].(string)).CombinedOutput()

	if err != nil {
		return err
	}

	return nil
}
