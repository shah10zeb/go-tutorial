package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about/", About)
	fmt.Printf("server listening to %s", portNumber)
	http.ListenAndServe(portNumber, nil)
}
