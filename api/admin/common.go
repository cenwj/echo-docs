package admin

import (
	"echo-docs/template"
	"github.com/labstack/echo"
	"net/http"
)

func Index(c echo.Context) error {
	return template.Render(c,
		http.StatusOK,
		"index",
		echo.Map{"title": "test"})
}


func Login(c echo.Context) error {
	return template.Render(c,
		http.StatusOK,
		"login",
		echo.Map{"title": "test"})
}