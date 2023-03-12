package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "about.page.html")
}

func RenderTemplate(w http.ResponseWriter, temp string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + temp)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing Template", err)
		return
	}
}
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about/", About)
	fmt.Printf("server listening to %s", portNumber)
	http.ListenAndServe(portNumber, nil)
}
