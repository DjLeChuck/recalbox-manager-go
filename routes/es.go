package routes

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/djlechuck/recalbox-manager/utils/errors"
	"github.com/kataras/iris"
	"github.com/spf13/viper"
)

// GetActionEsHandler handles GET requests on /es/action/:name.
func GetActionEsHandler(ctx iris.Context) {
	a := ctx.Params().Get("name")

	switch a {
	case "restart":
	case "start":
	case "stop":
	default:
		ctx.Values().Set("error", fmt.Sprintf("Invalid action %s", a))
		ctx.StatusCode(500)

		return
	}

	esp := viper.GetString("recalbox.emulationStationPath")
	cmd := exec.Command(esp, a)
	err := cmd.Start()
	if err != nil {
		ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
		ctx.StatusCode(500)

		return
	}

	ctx.JSON(iris.Map{
		"success": true,
	})
}

// GetStatusEsHandler handles GET requests on /es/status.
func GetStatusEsHandler(ctx iris.Context) {
	esp := viper.GetString("recalbox.emulationStationPath")
	cmd := esp + " status | cut -d ' ' -f 3"
	out, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
		ctx.StatusCode(500)

		return
	}

	ctx.JSON(iris.Map{
		"success": true,
		"running": "running" == strings.Trim(string(out), "\n"),
	})
}
