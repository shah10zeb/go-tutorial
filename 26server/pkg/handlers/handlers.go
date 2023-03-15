package handlers

import (
	"fmt"
	"net/http"
	"servver/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("HELLO ")
	render.RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
