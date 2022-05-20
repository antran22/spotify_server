package routes

import (
	"embed"
	"errors"
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

	tmpl := template.Must(template.ParseFS(embeddedTemplateFS, "templates/*.gohtml", "templates/**/*.gohtml"))

	funcMap := template.FuncMap{}

	funcMap["dict"] = dict

	tmpl.Funcs(funcMap)

	templateRendererInstance = &Renderer{template: tmpl}
	return templateRendererInstance
}

func (r *Renderer) Render(w io.Writer, path string, p interface{}) error {
	return r.template.ExecuteTemplate(w, path, p)
}

func dict(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, errors.New("invalid dict call")
	}
	dict := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, errors.New("dict keys must be strings")
		}
		dict[key] = values[i+1]
	}
	return dict, nil
}
