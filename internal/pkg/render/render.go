package render

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	Templates *template.Template
}

func New(templateDir string) (*TemplateRenderer, error) {

	templatePattern := templateDir + "/**/*.html"

	if templates, err := template.ParseGlob(templatePattern); err != nil {
		return nil, err
	} else {
		return &TemplateRenderer{
			Templates: templates,
		}, nil
	}

}

func (t *TemplateRenderer) Render(w io.Writer, name string, data any, c echo.Context) error {

	return t.Templates.ExecuteTemplate(w, name, data)

}
