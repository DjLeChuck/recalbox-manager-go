package main

import (
	"github.com/kataras/iris"

	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	app.Use(recover.New()) // recover from any http-relative panics
	app.Use(logger.New())  // log the requests to the terminal.

	app.StaticWeb("/css", "./assets/css")
	app.StaticWeb("/js", "./assets/js")

	tmpl := iris.Pug("./templates", ".pug").Layout("layouts/default.pug")
	tmpl.Reload(true) // reload templates on each request (development mode)

	app.RegisterView(tmpl)

	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.ViewData("Title", "Hi Page")
		ctx.ViewData("Name", "iris")
		ctx.View("views/home.pug")
	})

	app.Run(iris.Addr(":8888"), iris.WithoutServerError(iris.ErrServerClosed))
}
