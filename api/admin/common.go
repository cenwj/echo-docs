package admin

import (
	"net/http"

	"github.com/cenwj/echo-docs/template"
	"github.com/labstack/echo"
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
