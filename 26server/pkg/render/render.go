package render

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func RenderTemplateTEST(w http.ResponseWriter, temp string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+temp, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing +", err)
		return
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	//check if temp is already in map
	_, inMap := tc[t]
	if !inMap {
		log.Println("Creating temp and storing in cache")
		// need to create template
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		//we have in map
		log.Println("using cached template")
	}
	tmpl = tc[t]

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"templates/base.layout.tmpl",
	}
	//parse template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	//add temp to cache
	tc[t] = tmpl
	return nil
}
