package web

import (
	"net/http"

	"github.com/cenwj/echo-docs/template"
	"github.com/labstack/echo"
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
