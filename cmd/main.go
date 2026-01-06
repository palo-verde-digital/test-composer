package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/palo-verde-digital/test-composer/internal/app/editor"
	"github.com/palo-verde-digital/test-composer/internal/app/welcome"
	"github.com/palo-verde-digital/test-composer/internal/app/window"
	"github.com/palo-verde-digital/test-composer/internal/pkg/render"
)

func main() {

	e := initServer()
	e.Start(":8080")

}

func initServer() *echo.Echo {

	e := echo.New()

	if renderer, err := render.New(".templates"); err != nil {
		log.Panicf("failed to initialize template renderer:\n\t%s", err.Error())
	} else {
		e.Renderer = renderer
	}

	return registerRoutes(e)

}

func registerRoutes(e *echo.Echo) *echo.Echo {

	e.Static("/dist", ".dist")

	window.RegisterRoutes(e)
	welcome.RegisterRoutes(e)
	editor.RegisterRoutes(e)

	return e

}
