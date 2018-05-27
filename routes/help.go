package routes

import (
	"os"
	"os/exec"

	"github.com/kataras/iris"

	"github.com/spf13/viper"

	"github.com/djlechuck/recalbox-manager/structs"
	"github.com/djlechuck/recalbox-manager/utils/errors"
	"github.com/djlechuck/recalbox-manager/utils/recalbox"
	"github.com/djlechuck/recalbox-manager/utils/smartfile"
)

// GetHelpHandler handles the GET requests on /help.
func GetHelpHandler(ctx iris.Context) {
	ctx.ViewData("PageTitle", ctx.Translate("Dépannage"))

	ctx.ViewData("Tr", iris.Map{
		"Links": [3]structs.HelpLink{
			{
				Label: ctx.Translate("Le forum :"),
				Link:  ctx.Translate("https://forum.recalbox.com/"),
			}, {
				Label: ctx.Translate("Le chan IRC :"),
				Link:  ctx.Translate("https://kiwiirc.com/client/irc.freenode.net/#recalbox"),
			}, {
				Label: ctx.Translate("Le wiki :"),
				Link:  ctx.Translate("https://github.com/recalbox/recalbox-os/wiki/Home-(FR)"),
			},
		},
	})

	ctx.View("views/help.pug")
}

// GetLaunchRecalboxSupportHandler handles the GET requests on /help/recalbox-support.
func GetLaunchRecalboxSupportHandler(ctx iris.Context) {
	ss := viper.GetString("recalbox.supportScript")
	sp := viper.GetString("recalbox.savesPath")
	uuid, err := recalbox.PseudoUUID()
	if err != nil {
		ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
		ctx.StatusCode(500)

		return
	}

	// Generate the archive.
	ap := sp + "/recalbox-support-" + uuid + ".tar.gz"
	_, err = exec.Command(ss, ap).CombinedOutput()
	if err != nil {
		ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
		ctx.StatusCode(500)

		return
	}

	// Upload archive on SmartFile.
	path, err := smartfile.UploadArchive(ap)
	if err != nil {
		ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
		ctx.StatusCode(500)

		return
	}

	// Create the public link to the archive
	url, err := smartfile.GetLink(path)
	if err != nil {
		ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
		ctx.StatusCode(500)

		return
	}

	// Delete local archive.
	err = os.Remove(ap)
	if err != nil {
		ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
		ctx.StatusCode(500)

		return
	}

	ctx.JSON(map[string]interface{}{
		"success": true,
		"url":     url,
	})
}
