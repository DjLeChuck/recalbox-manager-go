package routes

import (
	"io/ioutil"

	"github.com/djlechuck/recalbox-manager/structs/forms"
	"github.com/djlechuck/recalbox-manager/utils/errors"
	"github.com/kataras/iris"
	"github.com/spf13/viper"
)

// GetRecalboxConfHandler handles the GET requests on /recalbox-conf.
func GetRecalboxConfHandler(ctx iris.Context) {
	cp := viper.GetString("recalbox.confPath")
	cc, err := ioutil.ReadFile(cp)
	if err != nil {
		ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
		ctx.StatusCode(500)

		return
	}

	ctx.ViewData("ConfPath", cp)
	ctx.ViewData("ConfContent", string(cc))

	ctx.ViewData("Tr", iris.Map{
		"BtnSave": ctx.Translate("BtnSave"),
	})

	ctx.View("views/recalbox-conf.pug")
}

// PostRecalboxConfHandler handles the POST requests on /recalbox-conf.
func PostRecalboxConfHandler(ctx iris.Context) {
	form := forms.RecalboxConf{}
	err := ctx.ReadForm(&form)
	if err != nil {
		ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
		ctx.StatusCode(500)

		return
	}

	cp := viper.GetString("recalbox.confPath")
	err = ioutil.WriteFile(cp, []byte(form.Content), 0644)
	if err != nil {
		ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
		ctx.StatusCode(500)

		return
	}

	ctx.Redirect("/recalbox-conf", 303)
}
