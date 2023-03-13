package render

import (
	"fmt"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, temp string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + temp)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing +", err)
		return
	}
}
