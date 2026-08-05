package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsed, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsed.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template: ", err)
	}
}
