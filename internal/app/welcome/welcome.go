package welcome

import (
	"io"
	"log"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/palo-verde-digital/test-composer/internal/app/window"
	"github.com/palo-verde-digital/test-composer/internal/pkg/project"
	"gopkg.in/yaml.v3"
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
	errData := errorData{}

	if projectName := c.FormValue(projectNameInput); strings.TrimSpace(projectName) == "" {
		errData.CreateError = "error: project name is required"
		log.Printf("welcome.createProject - %s", errData.CreateError)

		return c.Render(200, window.WelcomeTemplateName, errData)
	} else {
		project.OpenProject = &project.Project{
			Name: projectName,
		}

		log.Print("end - welcome.createProject")
		return c.Render(200, window.EditorTemplateName, project.OpenProject)
	}

}

func openProject(c echo.Context) error {

	log.Print("start - welcome.openProject")
	errData := errorData{}

	if fileInfo, err := c.FormFile(projectFileInput); err != nil {
		errData.OpenError = "error: unable to locate project file"
	} else if filename := fileInfo.Filename; !strings.HasSuffix(filename, ".yaml") || !strings.HasSuffix(filename, ".yml") {
		errData.OpenError = "error: project file must be .yml/.yaml"
	} else if projectFile, err := fileInfo.Open(); err != nil {
		errData.OpenError = "error: unable to open project file"
	} else if fileContent, err := io.ReadAll(projectFile); err != nil {
		errData.OpenError = "error: unable to read project file"
	} else if err = yaml.Unmarshal(fileContent, project.OpenProject); err != nil {
		errData.OpenError = "error: unable to read project data"
	} else {
		projectFile.Close()
	}

	if errData.OpenError != "" {
		log.Printf("welcome.openProject - %s", errData.OpenError)
		return c.Render(200, window.WelcomeTemplateName, errData)
	}

	log.Print("end - welcome.openProject")
	return c.Render(200, window.EditorTemplateName, project.OpenProject)

}
