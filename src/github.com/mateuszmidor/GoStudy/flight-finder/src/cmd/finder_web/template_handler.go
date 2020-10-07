package main

import (
	"net/http"
	"path/filepath"
	"text/template"
)

type templateHandler struct {
	templ *template.Template
}

func NewTemplateHandler(filename string) *templateHandler {
	template := template.Must(template.ParseFiles(filepath.Join("data", filename)))
	return &templateHandler{templ: template}
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Host": r.Host,
	}

	t.templ.Execute(w, data)
}
