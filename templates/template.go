package templates

import (
	"html/template"
	"net/http"
	"path"
)

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	lp := path.Join("templates", "base.html")
	fp := path.Join("templates", tmpl+".html")

	// Parse both the specific template file and the base layout together
	t, err := template.ParseFiles(lp, fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = t.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
