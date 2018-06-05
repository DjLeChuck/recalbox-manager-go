package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/kataras/iris"
	"github.com/spf13/viper"
)

// GetRomsHandler handles GET requests on /roms.
func GetRomsHandler(ctx iris.Context) {
	ctx.ViewData("PageTitle", ctx.Translate("ROMs"))

	// /v1/systems
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://"+viper.GetString("api.url")+":"+viper.GetString("api.port")+"/v1/systems", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", "Bearer "+ctx.Values().GetString("ApiToken"))
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var systems []string
	err = json.Unmarshal(body, &systems)
	if err != nil {
		panic(err)
	}

	ctx.ViewData("Systems", systems)

	ctx.View("views/roms.pug")
}
