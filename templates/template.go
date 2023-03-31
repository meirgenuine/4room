package templates

import (
	"html/template"
	"net/http"
	"path"
)

var (
	basePath = "templates"
)

func init() {
	// You can set a different base path for the templates here if needed
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tmplPath := path.Join(basePath, tmpl+".html")
	baseTemplatePath := path.Join(basePath, "base.html")

	t, err := template.ParseFiles(baseTemplatePath, tmplPath)
	if err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
