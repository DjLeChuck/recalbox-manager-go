package routes

import (
	"os"

	"github.com/kataras/iris"

	"github.com/djlechuck/recalbox-manager/structs"
)

// GetHomeHandler handles the GET requests on /.
func GetHomeHandler(ctx iris.Context) {
	hostname, err := os.Hostname()

	if err != nil {
		ctx.Values().Set("errorMessage", err.Error())
		ctx.StatusCode(500)

		return
	}

	ctx.ViewData("PageTitle", ctx.Translate("Accueil"))
	ctx.ViewData("Tiles", []structs.HomeTile{
		{
			Link:  "//" + hostname + ":8080/gamepad.html?analog",
			Image: "/img/gamepad.png",
			Label: ctx.Translate("Utiliser le gamepad virtuel"),
		},
		{
			Link:  "//" + hostname + ":8080/keyboard.html",
			Image: "/img/keyboard.png",
			Label: ctx.Translate("Utiliser le clavier virtuel"),
		},
		{
			Link:  "//" + hostname + ":8080/touchepad.html",
			Image: "/img/touchpad.png",
			Label: ctx.Translate("Utiliser le touchpad virtuel"),
		},
	})

	ctx.View("views/home.pug")
}
