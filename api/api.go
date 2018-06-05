package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/kataras/iris"
	"github.com/spf13/viper"
)

// New returns a new handler which configures the API for the app.
func New(app *iris.Application) iris.Handler {
	// Connect to the API
	url := "http://" + viper.GetString("api.url") + ":" + viper.GetString("api.port") + "/login"
	jsonStr := []byte(`{"user":"recaluser", "pass": "recalboxrox"}`)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var data API
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}

	return func(ctx iris.Context) {
		ctx.Values().Set("ApiToken", data.Token)

		ctx.Next()
	}
}

// Configure registers the API to the app.
func Configure(app *iris.Application) {
	h := New(app)
	app.Use(h)
}

// API represents the connection to the recalbox API.
type API struct {
	Token string `json:"token"`
}
