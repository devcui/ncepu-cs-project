package main

import (
	"fmt"
	"io"
	"text/template"

	"github.com/devcui/ncepu-cs-project/config"
	"github.com/devcui/ncepu-cs-project/mysql"
	"github.com/devcui/ncepu-cs-project/oauth2"
	"github.com/devcui/ncepu-cs-project/redis"
	"github.com/devcui/ncepu-cs-project/routes"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	var additional map[string]interface{}
	if data != nil {
		additional = data.(map[string]interface{})
	} else {
		additional = map[string]interface{}{}
	}
	additional["Path"] = config.Value.Server.Path
	return t.templates.ExecuteTemplate(w, name, additional)
}

func main() {
	if err := config.Setup(); err != nil {
		log.Fatalf("config.Setup() error: %v", err)
		return

	}
	if err := mysql.Setup(); err != nil {
		log.Fatalf("mysql.Setup() error: %v", err)
		return
	}
	redis.SetupSession()
	if err := oauth2.SetupServer(); err != nil {
		log.Fatal("oauth2.SetupServer() error: %v", err)
		return
	}
	app := echo.New()
	app.Logger.SetLevel(log.DEBUG)
	app.Renderer = &Template{
		templates: template.Must(template.ParseGlob("static/template/**/*")),
	}
	api := app.Group(config.Value.Server.Path)
	api.Static(config.Value.Server.Path+"/static/asset", "static/asset")
	api.Static(config.Value.Server.Path+"/page", "static/page")
	routes.Routes(api)
	app.Logger.Debug(app.Start(fmt.Sprintf("%s:%d", config.Value.Server.Host, config.Value.Server.Port)))
}
