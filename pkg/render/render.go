package render

import (
	"fmt"
	"html/template"
	"net/http"
)

//map of functions can use in the templates
var functions = template.FuncMap{}

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}
