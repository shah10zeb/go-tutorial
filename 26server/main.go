package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "THIS IS HOME PAGE")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum, err := AddValues(2, 2)
	if err == nil {
		fmt.Printf("SUM %d", sum)
		fmt.Fprintf(w, "THIS IS ABOUT PAGE")
	}
}

func AddValues(x, y int) (int, error) {
	var sum int
	sum = x + y
	return sum, nil
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	http.ListenAndServe(portNumber, nil)
}
