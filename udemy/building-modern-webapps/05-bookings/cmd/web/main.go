package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/gerdreiss/bookings/pkg/config"
	"github.com/gerdreiss/bookings/pkg/handlers"
	"github.com/gerdreiss/bookings/pkg/render"
)

func main() {

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache: ", err)
	}

	app = config.AppConfig{
		InProduction:  false,
		UseCache:      false,
		TemplateCache: tc,
		InfoLog:       log.Default(),
	}

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = !app.InProduction
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

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
