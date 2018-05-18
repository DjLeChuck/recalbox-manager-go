package routes

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/kataras/iris"

	"github.com/spf13/viper"

	"github.com/djlechuck/recalbox-manager/store"
)

// FormData represents the submitted data of a form.
type FormData map[string]interface{}

// GetAudioHandler handles the GET requests on /audio.
func GetAudioHandler(ctx iris.Context) {
	sess := store.Sessions.Start(ctx)

	ctx.ViewData("PageTitle", ctx.Translate("Audio"))
	ctx.ViewData("FormSended", sess.GetFlashString("formSended"))
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
	formData := FormData{}
	err := ctx.ReadForm(&formData)

	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
	}

	// Specific case for bgMusic
	if _, ok := formData["audio-bgmusic"]; !ok {
		formData["audio-bgmusic"] = "0"
	}

	pythonFile := viper.GetString("recalbox.pythonSettingsFile")

	for k, v := range formData {
		normalizedKey := strings.Replace(k, "-", ".", -1)
		output, cErr := exec.Command("python", pythonFile, "-command", "save", "-key", normalizedKey, "-value", v.(string)).CombinedOutput()

		if cErr != nil {
			fmt.Println(cErr.Error())
		}

		fmt.Println(string(output))
	}

	configScript := viper.GetString("recalbox.configScript")
	output, err := exec.Command(configScript, "volume", formData["audio-volume"].(string)).CombinedOutput()
	fmt.Println(configScript)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(output))

	sess := store.Sessions.Start(ctx)
	sess.SetFlash("formSended", ctx.Translate("ConfigurationSaved"))

	ctx.Redirect("/audio", 303)
}
