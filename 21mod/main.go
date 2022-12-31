package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//not a good way to run go run main.go

	//also use go mod init github.com/shahzeb/mymodules

	//go get -u github.com/gorilla/mux
	greeter()

	fmt.Println("HELLO MODDD")
	r := mux.NewRouter()
	r.HandleFunc("/", serverHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))

}

func greeter() {
	fmt.Println("HELLO GREETINGS")
}

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<h1>HELLO HELLO HELLO<h1>`))
}
