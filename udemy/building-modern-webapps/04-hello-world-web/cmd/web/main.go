package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gerdreiss/go-course/pkg/config"
	"github.com/gerdreiss/go-course/pkg/handlers"
	"github.com/gerdreiss/go-course/pkg/render"
	"github.com/gerdreiss/go-course/pkg/utils"
)

const addr = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache: ", err)
	}

	app.UseCache = false
	app.TemplateCache = tc
	render.NewTemplates(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	http.HandleFunc("/divide/{x}/{y}", utils.Divide)

	fmt.Printf("Starting server on port %s...\n", addr)
	_ = http.ListenAndServe(addr, nil)
}
