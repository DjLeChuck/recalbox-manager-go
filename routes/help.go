package routes

import (
	"github.com/djlechuck/recalbox-manager/structs"
	"github.com/kataras/iris"
)

// GetHelpHandler handles the GET requests on /help.
func GetHelpHandler(ctx iris.Context) {
	ctx.ViewData("PageTitle", ctx.Translate("Dépannage"))

	ctx.ViewData("Tr", iris.Map{
		"Links": [3]structs.HelpLink{
			{
				Label: ctx.Translate("Le forum :"),
				Link:  ctx.Translate("https://forum.recalbox.com/"),
			}, {
				Label: ctx.Translate("Le chan IRC :"),
				Link:  ctx.Translate("https://kiwiirc.com/client/irc.freenode.net/#recalbox"),
			}, {
				Label: ctx.Translate("Le wiki :"),
				Link:  ctx.Translate("https://github.com/recalbox/recalbox-os/wiki/Home-(FR)"),
			},
		},
	})

	ctx.View("views/help.pug")
}
