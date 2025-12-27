package editor

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/palo-verde-digital/test-composer/internal/app/window"
	"github.com/palo-verde-digital/test-composer/internal/pkg/project"
)

func RegisterRoutes(e *echo.Echo) {

	e.GET("/project/close", closeProject)

}

func closeProject(c echo.Context) error {

	log.Print("start - editor.closeProject")
	defer log.Print("end - editor.closeProject")

	err := project.Write()
	if err != nil {
		return c.Render(200, window.EditorTemplateName, project.OpenProject)
	}

	return c.Render(200, window.WelcomeTemplateName, nil)

}
