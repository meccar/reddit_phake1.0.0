package web

import (
	"fmt"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("./templates/*.html"))

func Render(w http.ResponseWriter, tmpl string, data interface{}) {
	fmt.Printf("\n logging r %v\n", tmpl)
	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
