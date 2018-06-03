package routes

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/kataras/iris"
	"github.com/spf13/viper"

	"github.com/djlechuck/recalbox-manager/store"
	"github.com/djlechuck/recalbox-manager/structs"
	"github.com/djlechuck/recalbox-manager/structs/forms"
	"github.com/djlechuck/recalbox-manager/utils/errors"
)

// GetSecurityHandler handles GET requests on /security.
func GetSecurityHandler(ctx iris.Context) {
	ctx.ViewData("PageTitle", ctx.Translate("Securité"))

	ctx.View("/views/security.pug")
}

// PostSecurityHandler handles POST requests on /security.
func PostSecurityHandler(ctx iris.Context) {
	form := forms.Security{}
	if err := ctx.ReadForm(&form); err != nil {
		ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
		ctx.StatusCode(500)

		return
	}

	path := viper.GetString("auth.path") + viper.GetString("auth.file")
	sess := store.Sessions.Start(ctx)
	auth := sess.Get("authentication").(*structs.Authentication)

	if !form.NeedAuth {
		// If no auth required, delete the auth file.
		err := os.Remove(path)
		if err != nil {
			ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
			ctx.StatusCode(500)

			return
		}

		auth.Enabled = false
	} else {
		// If auth required, store credentials in auth file.
		cred := structs.Credentials{
			Login:    form.Username,
			Password: form.Password,
		}
		j, err := json.Marshal(cred)
		if err != nil {
			ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
			ctx.StatusCode(500)

			return
		}

		err = ioutil.WriteFile(path, j, 0644)
		if err != nil {
			ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
			ctx.StatusCode(500)

			return
		}

		// Reset authentication
		auth.Reset(cred)
	}

	sess.Set("authentication", auth)

	ctx.Redirect("/security", 303)
}
