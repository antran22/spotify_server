package template

import (
	"embed"
	"html/template"
	"io"
)

var (
	//go:embed "templates/*"
	embeddedTemplateFS embed.FS
)

type Renderer struct {
	template *template.Template
}

var templateRendererInstance *Renderer

func GetTemplateRenderer() *Renderer {
	if templateRendererInstance != nil {
		return templateRendererInstance
	}

	tmpl := template.Must(template.ParseFS(embeddedTemplateFS, "templates/*.gohtml"))

	templateRendererInstance = &Renderer{template: tmpl}
	return templateRendererInstance
}

func (r *Renderer) Render(w io.Writer, path string, p interface{}) error {
	return r.template.ExecuteTemplate(w, path, p)
}
