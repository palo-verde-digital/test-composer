package editor

import (
	"log"
	"strings"
	"unicode"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/palo-verde-digital/test-composer/internal/app/window"
	"github.com/palo-verde-digital/test-composer/internal/pkg/project"
	"github.com/palo-verde-digital/test-composer/pkg/stringutil"
)

type appData struct {
	Id     string
	Errors map[string]string
	App    project.Application
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

func deleteApp(c echo.Context) error {

	log.Print("start - app.deleteApp")
	defer log.Print("end - app.deleteApp")

	appId := c.Param("appId")
	delete(project.OpenProject.Apps, appId)

	return readApps(c)

}

func updateAppImage(c echo.Context) error {

	log.Print("start - app.updateAppImage")
	defer log.Print("end - app.updateAppImage")

	errors := map[string]string{}

	appId := c.Param("appId")
	app := project.OpenProject.Apps[appId]

	appImage := c.FormValue("app-" + appId + "-image")
	app.Image = appImage

	project.OpenProject.Apps[appId] = app

	if !strings.Contains(appImage, ":") {
		errors["image"] = "must be of format IMAGE_NAME:TAG"
	} else if !unicode.IsLetter(rune(appImage[0])) {
		errors["image"] = "must start w/ a lowercase letter"
	} else if stringutil.ContainsUpper(appImage) {
		errors["image"] = "cannot contain uppercase letters"
	}

	return c.Render(200, window.AppTemplateName, appData{
		Id:     appId,
		Errors: errors,
		App:    app,
	})

}

func createAppEnv(c echo.Context) error {

	log.Print("start - app.createAppEnv")
	defer log.Print("end - app.createAppEnv")

	appId := c.Param("appId")

	envKey := c.FormValue("app-" + appId + "-env-key")
	envVal := c.FormValue("app-" + appId + "-env-val")

	project.OpenProject.Apps[appId].Env[uuid.NewString()] = project.Variable{
		Key:   envKey,
		Value: envVal,
	}

	return c.Render(200, window.AppTemplateName, appData{
		Id:  appId,
		App: project.OpenProject.Apps[appId],
	})

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
