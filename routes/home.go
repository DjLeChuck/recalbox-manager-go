package routes

import (
	"os"

	"github.com/kataras/iris"

	"github.com/djlechuck/recalbox-manager/structs"
	"github.com/djlechuck/recalbox-manager/utils/errors"
)

// GetHomeHandler handles the GET requests on /.
func GetHomeHandler(ctx iris.Context) {
	hostname, err := os.Hostname()

	if err != nil {
		ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
		ctx.StatusCode(500)

		return
	}

	ctx.ViewData("PageTitle", ctx.Translate("Accueil"))
	ctx.ViewData("Tiles", []structs.HomeTile{
		{
			Link:  "//" + hostname + ":8080/gamepad.html?analog",
			Image: "/static/img/gamepad.png",
			Label: ctx.Translate("Utiliser le gamepad virtuel"),
		},
		{
			Link:  "//" + hostname + ":8080/keyboard.html",
			Image: "/static/img/keyboard.png",
			Label: ctx.Translate("Utiliser le clavier virtuel"),
		},
		{
			Link:  "//" + hostname + ":8080/touchpad.html",
			Image: "/static/img/touchpad.png",
			Label: ctx.Translate("Utiliser le touchpad virtuel"),
		},
	})

	ctx.View("views/home.pug")
}
