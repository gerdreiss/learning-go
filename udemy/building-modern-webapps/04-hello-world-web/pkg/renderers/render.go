package renderers

import (
	"fmt"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsed, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsed.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template: ", err)
	}
}
