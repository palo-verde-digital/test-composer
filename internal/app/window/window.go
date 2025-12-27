package window

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/palo-verde-digital/test-composer/internal/pkg/project"
)

const (
	TemplateName        = "app"
	WelcomeTemplateName = "welcome"
	EditorTemplateName  = "editor"
)

func RegisterRoutes(e *echo.Echo) {

	e.GET("/", app)

}

func app(c echo.Context) error {

	log.Print("window.app")

	return c.Render(200, TemplateName, project.OpenProject)

}
