package editor

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/palo-verde-digital/test-composer/internal/app/window"
	"github.com/palo-verde-digital/test-composer/internal/pkg/project"
)

type EditorData struct {
	Errors  map[string]string
	Project *project.Project
}

func RegisterRoutes(e *echo.Echo) {

	e.GET("/project/close", closeProject)

	e.GET("/project/app", readApps)
	e.GET("/project/app/:appId", readApp)
	e.GET("/project/app/create", createApp)

	e.POST("/project/app/:appId/image", updateAppImage)
	e.POST("/project/app/:appId/env", createAppEnv)
	e.POST("/project/app/:appId/env/:envId", updateAppEnv)

	e.DELETE("/project/app/:appId", deleteApp)
	e.DELETE("/project/app/:appId/env/:envId", deleteAppEnv)

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
