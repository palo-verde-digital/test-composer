package welcome

import (
	"log"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/palo-verde-digital/test-composer/internal/app/window"
	"github.com/palo-verde-digital/test-composer/internal/pkg/project"
)

type errorData struct {
	CreateError, OpenError string
}

const (
	projectNameInput = "project-name"
	projectFileInput = "project-file"
)

func RegisterRoutes(e *echo.Echo) {

	e.POST("/project/create", createProject)
	e.POST("/project/open", openProject)

}

func createProject(c echo.Context) error {

	log.Print("start - welcome.createProject")
	defer log.Print("end - welcome.createProject")

	projectName := c.FormValue(projectNameInput)
	if strings.TrimSpace(projectName) == "" {
		return c.Render(200, window.WelcomeTemplateName, errorData{CreateError: "project name is required"})
	}

	project.Create(projectName)

	return c.Render(200, window.EditorTemplateName, project.OpenProject)

}

func openProject(c echo.Context) error {

	log.Print("start - welcome.openProject")
	defer log.Print("end - welcome.openProject")

	fileInfo, err := c.FormFile(projectFileInput)
	if err != nil {
		return c.Render(200, window.WelcomeTemplateName, errorData{OpenError: "invalid file info"})
	}

	err = project.Read(fileInfo)
	if err != nil {
		return c.Render(200, window.WelcomeTemplateName, errorData{OpenError: err.Error()})
	}

	return c.Render(200, window.EditorTemplateName, project.OpenProject)

}
