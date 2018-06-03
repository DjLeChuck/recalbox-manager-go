package routes

import (
	"github.com/djlechuck/recalbox-manager/store"
	"github.com/djlechuck/recalbox-manager/structs"
	"github.com/djlechuck/recalbox-manager/structs/forms"
	"github.com/djlechuck/recalbox-manager/utils/errors"
	"github.com/kataras/iris"
)

// GetLoginHandler handles GET requests on /login.
func GetLoginHandler(ctx iris.Context) {
	sess := store.Sessions.Start(ctx)
	auth := sess.Get("authentication").(*structs.Authentication)
	if auth.IsAuthenticated {
		ctx.Redirect("/")
		return
	}

	ctx.View("views/login.pug")
}

// PostLoginHandler handles POST requests on /login.
func PostLoginHandler(ctx iris.Context) {
	form := forms.Login{}
	if err := ctx.ReadForm(&form); err != nil {
		ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
		ctx.StatusCode(500)

		return
	}

	sess := store.Sessions.Start(ctx)
	auth := sess.Get("authentication").(*structs.Authentication)
	credentials := auth.Credentials

	if credentials.Login == form.Username && credentials.Password == form.Password {
		auth.IsAuthenticated = true
		sess.Set("authentication", auth)
	}

	ctx.Redirect("/", 303)
}

// GetLogoutHandler handles GET requests on /logout.
func GetLogoutHandler(ctx iris.Context) {
	sess := store.Sessions.Start(ctx)
	auth := sess.Get("authentication").(*structs.Authentication)

	auth.IsAuthenticated = false
	sess.Set("authentication", auth)

	ctx.Redirect("/login")
}
