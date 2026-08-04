package main

import (
	"errors"
	"fmt"
	"net/http"
	"text/template"
)

const addr = ":8080"

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("division by zero")
	}
	return x / y, nil
}

func Divide(w http.ResponseWriter, r *http.Request) {
	x := 100.0
	y := 10.0
	res, err := divide(x, y)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	} else {
		fmt.Fprintf(w, "%f divided by %f is %f", x, y, res)
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsed, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsed.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template: ", err)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Printf("Starting server on port %s...\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("WTF!", err)
	}
}
