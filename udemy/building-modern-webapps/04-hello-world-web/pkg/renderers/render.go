package renderers

import (
	"fmt"
	"net/http"
	"text/template"
)

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	tmpl, exists := tc[t]
	if !exists {
		err := createTemplateCache(t)
		if err != nil {
			fmt.Println("shitty template: ", err)
			return
		}
	}

	err := tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template: ", err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	tc[t] = tmpl
	return nil
}
