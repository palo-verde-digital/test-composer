package window

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/palo-verde-digital/test-composer/internal/pkg/project"
)

const (
	TemplateName        = "index"
	WelcomeTemplateName = "welcome"
	EditorTemplateName  = "editor"
	AppsTemplateName    = "apps"
	AppTemplateName     = "app"
)

func RegisterRoutes(e *echo.Echo) {

	e.GET("/", app)

}

func app(c echo.Context) error {

	log.Print("start - window.app")
	defer log.Print("end - window.app")

	if project.OpenProject != nil {
		return c.Render(200, TemplateName, project.OpenProject.Name)
	}

	return c.Render(200, TemplateName, nil)

}
