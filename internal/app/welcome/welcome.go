package welcome

import (
	"log"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/palo-verde-digital/test-composer/internal/app/window"
	"github.com/palo-verde-digital/test-composer/internal/pkg/project"
)

type welcomeData struct {
	CreateError, OpenError string
}

func RegisterRoutes(e *echo.Echo) {

	e.POST("/project/create", createProject)
	e.POST("/project/open", openProject)

}

func createProject(c echo.Context) error {

	log.Print("start - welcome.createProject")
	defer log.Print("end - welcome.createProject")

	projectName := c.FormValue("project-name")
	if strings.TrimSpace(projectName) == "" {
		return c.Render(200, window.WelcomeTemplateName, welcomeData{CreateError: "project name is required"})
	}

	project.Create(projectName)

	return c.Render(200, window.EditorTemplateName, projectName)

}

func openProject(c echo.Context) error {

	log.Print("start - welcome.openProject")
	defer log.Print("end - welcome.openProject")

	fileInfo, err := c.FormFile("project-file")
	if err != nil {
		return c.Render(200, window.WelcomeTemplateName, welcomeData{OpenError: "invalid file info"})
	}

	err = project.Read(fileInfo)
	if err != nil {
		return c.Render(200, window.WelcomeTemplateName, welcomeData{OpenError: err.Error()})
	}

	return c.Render(200, window.EditorTemplateName, project.OpenProject.Name)

}
