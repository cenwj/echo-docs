package web

import (
	"net/http"

	"github.com/labstack/echo"
	"echo-docs/template"
)

func Home(c echo.Context) error {

	return c.String(http.StatusOK, "Main hgjg Index")

}

func Index(c echo.Context) error {
	return template.Render(c,
		http.StatusOK,
		"socket",
		echo.Map{})
}
