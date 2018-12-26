package route

import (
	"echo-docs/api/admin"
	"echo-docs/api/web"

	"echo-docs/template"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"echo-docs/sockets"
)

func Init() *echo.Echo {

	e := echo.New()

	e.Debug = true

	e.Static("/", "public")

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Renderer = template.New(template.TemplateConfig{
		Root:      "template/web",
		Extension: ".html",
		Master:    "",
	})

	e.GET("/ws", sockets.WsQuote)
	e.GET("/", web.Index)

	e.GET("/users", web.GetUsers)
	e.GET("/createUsers", web.CreateUsers)

	// 后端
	am := template.NewMiddleware(template.TemplateConfig{
		Root:      "template/admin",
		Extension: ".html",
		Master:    "layout",
	})

	a := e.Group("/admin", am)
	a.GET("/index", admin.Index)
	a.GET("/login", admin.Login)

	return e
}
