package main

import (
	"fmt"
	"net/http"
	"servver/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about/", handlers.About)
	fmt.Printf("server listening to %s", portNumber)
	http.ListenAndServe(portNumber, nil)
}
