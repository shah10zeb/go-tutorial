package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://lco.dev:3000/learn?coursename=go&payment=done"

func main() {
	fmt.Println("WELCOME TO GO URLS")
	fmt.Println(myurl)
	result, _ := url.Parse(myurl)
	fmt.Println(result)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	qparam := result.Query()
	fmt.Printf("TYPE of QUERRY PARAMS ARE : %T ", qparam)
	fmt.Println(qparam["coursename"])

	for _, val := range qparam {
		fmt.Println("PARAM : ", val)
	}

}
