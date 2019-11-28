package config

import (
	"../controllers"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"html/template"
)

// application struct
type App struct {
	Echo *echo.Echo
}
// initilize app and set dependencies .
func (app *App) init() {
	e := echo.New()
	app.Echo = e
	e.Static("/", "public")
	e.Renderer = &TemplateRegistry{
		Templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	homeController := controllers.HomeController{
		Echo: e,
	}
	wsController := controllers.WsController{
		Echo: e,
	}
	wsController.Init()
	homeController.Init()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

}

// init
func (app *App) Run(port int) {
	app.init()
	address := fmt.Sprintf(":%d", port)
	fmt.Println(app.Echo.Start(address))
}
