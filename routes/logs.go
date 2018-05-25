package routes

import (
	"io/ioutil"

	"github.com/kataras/iris"

	"github.com/spf13/viper"

	"github.com/djlechuck/recalbox-manager/store"
	"github.com/djlechuck/recalbox-manager/structs/forms"
	"github.com/djlechuck/recalbox-manager/utils/errors"
)

// GetLogsHandler handles the GET requests on /logs.
func GetLogsHandler(ctx iris.Context) {
	sess := store.Sessions.Start(ctx)
	list := viper.GetStringSlice("recalbox.logsPaths")
	log := sess.GetString("AskedLog")
	sess.Delete("AskedLog")

	ctx.ViewData("LogContent", "")

	if log != "" {
		content, err := ioutil.ReadFile(log)
		if err != nil {
			ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
			ctx.StatusCode(500)

			return
		}

		ctx.ViewData("LogContent", string(content))
	}

	ctx.ViewData("PageTitle", ctx.Translate("Logs"))
	ctx.ViewData("LogName", log)
	ctx.ViewData("LogsList", list)

	ctx.View("views/logs.pug")
}

// PostLogsHandler handles the POST requests on /logs.
func PostLogsHandler(ctx iris.Context) {
	sess := store.Sessions.Start(ctx)
	form := forms.Logs{}
	err := ctx.ReadForm(&form)

	if err != nil {
		ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
		ctx.StatusCode(500)

		return
	}

	sess.Set("AskedLog", form.File)

	ctx.Redirect("/logs", 303)
}
