package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	tc, err := createTemplateCache()
	if err != nil {
		log.Fatal("error creating template cache: ", err)
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("no template for ", tmpl)
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache() (map[string]*template.Template, error) {
	tc := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return tc, err
	}

	layouts, err := filepath.Glob("./templates/*.layout.tmpl")
	if err != nil {
		return tc, err
	}
	foundLayouts := len(layouts) > 0

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return tc, err
		}

		if foundLayouts {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return tc, err
			}
		}

		tc[name] = ts
	}

	return tc, nil
}

// /////////////////////////////////////////////////////////////////////////////////
// template cache
var localTemplateCache = make(map[string]*template.Template)

// render using manual cache creation and management
func RenderTemplateManual(w http.ResponseWriter, t string) {
	tmpl, exists := localTemplateCache[t]
	if !exists {
		err := createTemplateCacheManual(t)
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

func createTemplateCacheManual(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	localTemplateCache[t] = tmpl
	return nil
}
