package routes

import (
	"io/ioutil"

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
