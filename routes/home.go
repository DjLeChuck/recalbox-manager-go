package routes

import (
	"os"

	"github.com/kataras/iris"
)

// HomeTile represents a tile on the homepage.
type HomeTile struct {
	Link, Image, Label string
}

// GetHomeHandler handles the GET requests on /.
func GetHomeHandler(ctx iris.Context) {
	hostname, err := os.Hostname()

	if err != nil {
		panic(err)
	}

	ctx.ViewData("PageTitle", ctx.Translate("Accueil"))
	ctx.ViewData("Tiles", []HomeTile{
		{"//" + hostname + ":8080/gamepad.html?analog", "/img/gamepad.png", ctx.Translate("Utiliser le gamepad virtuel")},
		{"//" + hostname + ":8080/keyboard.html", "/img/keyboard.png", ctx.Translate("Utiliser le clavier virtuel")},
		{"//" + hostname + ":8080/touchepad.html", "/img/touchpad.png", ctx.Translate("Utiliser le touchpad virtuel")},
	})

	ctx.View("views/home.pug")
}
