package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gerdreiss/go-course/pkg/config"
	"github.com/gerdreiss/go-course/pkg/handlers"
	"github.com/gerdreiss/go-course/pkg/render"
)

const addr = ":8080"

func main() {
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache: ", err)
	}

	app := config.AppConfig{
		UseCache:      false,
		TemplateCache: tc,
		InfoLog:       log.Default(),
	}

	render.NewTemplates(&app)
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	fmt.Printf("Starting server on port %s...\n", addr)
	srv := &http.Server{
		Addr:    addr,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
