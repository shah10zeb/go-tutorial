package main

import (
	"fmt"
	"net/http"
	"servver/pkg/handlers"
)

const portNumber = ":8080"

func main() {

	http.HandleFunc("/about/", handlers.About)
	http.HandleFunc("/", handlers.Home)
	fmt.Printf("server listening to %s", portNumber)
	http.ListenAndServe(portNumber, nil)
}
