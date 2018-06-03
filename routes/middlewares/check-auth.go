package middlewares

import (
	"github.com/djlechuck/recalbox-manager/store"
	"github.com/djlechuck/recalbox-manager/structs"
	"github.com/kataras/iris"
)

// CheckAuthMiddlewareHandler checks if an authentification is required to
// access the routes.
func CheckAuthMiddlewareHandler(ctx iris.Context) {
	sess := store.Sessions.Start(ctx)
	auth := sess.Get("authentication").(*structs.Authentication)
	if !auth.Enabled || auth.IsAuthenticated {
		ctx.Next()
		return
	}

	ctx.StopExecution()
	ctx.Redirect("/login")
	return
}
