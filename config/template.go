package config

import (
	"github.com/labstack/echo"
	"html/template"
	"io"
)

type TemplateRegistry struct {
	Templates *template.Template
}

// Implement e.Renderer interface
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}