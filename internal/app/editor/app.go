package editor

import (
	"log"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/palo-verde-digital/test-composer/internal/app/window"
	"github.com/palo-verde-digital/test-composer/internal/pkg/project"
)

type appData struct {
	Id   string
	Errs map[string]string
	App  project.Application
}

func readApps(c echo.Context) error {

	log.Print("start - app.readApps")
	defer log.Print("end - app.readApps")

	return c.Render(200, window.AppsTemplateName, project.OpenProject.Apps)

}

func createApp(c echo.Context) error {

	log.Print("start - app.createApp")
	defer log.Print("end - app.createApp")

	project.OpenProject.Apps[uuid.NewString()] = project.CreateApplication()

	return readApps(c)

}

func readApp(c echo.Context) error {

	log.Print("start - app.readApp")
	defer log.Print("end - app.readApp")

	appId := c.Param("appId")

	return c.Render(200, window.AppTemplateName, appData{
		Id:  appId,
		App: project.OpenProject.Apps[appId],
	})

}

func updateApp(c echo.Context) error {

	log.Print("start - app.updateApp")
	defer log.Print("end - app.updateApp")

	errs := map[string]string{}

	appId := c.Param("appId")
	app := project.OpenProject.Apps[appId]

	app.Name = c.FormValue("app-" + appId + "-name")
	errs["name"] = project.ValidateApplicationName(app.Name)

	app.Image = c.FormValue("app-" + appId + "-image")
	errs["image"] = project.ValidateApplicationImage(appId, app.Image)

	app.IsApi = c.FormValue("app-"+appId+"-isApi") == "on"

	if app.IsApi {
		app.ApiPort = c.FormValue("app-" + appId + "-apiPort")
		errs["apiPort"] = project.ValidateApplicationApiPort(app.ApiPort)
	}

	envKey, envVal := c.FormValue("app-"+appId+"-env-key"), c.FormValue("app-"+appId+"-env-val")
	if envKey != "" && envVal != "" {
		app.Env[uuid.NewString()] = project.Variable{
			Key:   envKey,
			Value: envVal,
		}
	}

	project.OpenProject.Apps[appId] = app
	return c.Render(200, window.AppTemplateName, appData{
		Id:   appId,
		Errs: errs,
		App:  app,
	})

}

func deleteApp(c echo.Context) error {

	log.Print("start - app.deleteApp")
	defer log.Print("end - app.deleteApp")

	appId := c.Param("appId")
	delete(project.OpenProject.Apps, appId)

	return readApps(c)

}

func updateAppEnv(c echo.Context) error {

	log.Print("start - app.updateAppEnv")
	defer log.Print("end - app.updateAppEnv")

	appId := c.Param("appId")
	envId := c.Param("envId")

	envKey := c.FormValue("app-" + appId + "-env-" + envId + "-key")
	envValue := c.FormValue("app-" + appId + "-env-" + envId + "-val")

	project.OpenProject.Apps[appId].Env[envId] = project.Variable{
		Key:   envKey,
		Value: envValue,
	}

	return c.Render(200, window.AppTemplateName, appData{
		Id:  appId,
		App: project.OpenProject.Apps[appId],
	})

}

func deleteAppEnv(c echo.Context) error {

	log.Print("start - app.deleteAppEnv")
	defer log.Print("end - app.deleteAppEnv")

	appId := c.Param("appId")
	envId := c.Param("envId")

	delete(project.OpenProject.Apps[appId].Env, envId)

	return c.Render(200, window.AppTemplateName, appData{
		Id:  appId,
		App: project.OpenProject.Apps[appId],
	})

}
