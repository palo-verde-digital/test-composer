package editor

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/palo-verde-digital/test-composer/internal/app/window"
	"github.com/palo-verde-digital/test-composer/internal/pkg/project"
)

func RegisterRoutes(e *echo.Echo) {

	e.GET("/project/close", closeProject)

	e.GET("/project/app", readApps)
	e.GET("/project/app/:appId", readApp)
	e.GET("/project/app/create", createApp)
	e.GET("/project/infrastructure", readInfrastucture)
	e.GET("/project/infrastructure/postgres", readPostgres)
	e.GET("/project/infrastructure/kafka", readKafka)
	e.GET("/project/infrastructure/redis", readRedis)

	e.POST("/project/app/:appId", updateApp)
	e.POST("/project/app/:appId/:envId", updateAppEnv)
	e.POST("/project/infrastructure/postgres", updatePostgres)
	e.POST("/project/infrastructure/kafka", updateKafka)
	e.POST("/project/infrastructure/kafka/:topicId", updateKafkaTopic)
	e.POST("/project/infrastructure/redis", updateRedis)

	e.DELETE("/project/app/:appId", deleteApp)
	e.DELETE("/project/app/:appId/:envId", deleteAppEnv)
	e.DELETE("/project/infrastructure/kafka/:topicId", deleteKafkaTopic)

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
