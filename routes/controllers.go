package routes

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/kataras/iris"

	"github.com/spf13/viper"

	"github.com/djlechuck/recalbox-manager/store"
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

	sess := store.Sessions.Start(ctx)
	sess.SetFlash("formSended", ctx.Translate("ConfigurationSaved"))

	ctx.Redirect("/controllers", 303)
}
