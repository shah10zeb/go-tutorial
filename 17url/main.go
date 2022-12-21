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
	//.Scheme = protocol
	//.Port() = port number
	fmt.Println(result.Port())
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	qparam := result.Query()
	fmt.Printf("TYPE of QUERRY PARAMS ARE : %T \n", qparam)
	// url.Values fancy name for key values
	fmt.Println(qparam["coursename"])

	for _, val := range qparam {
		fmt.Println("PARAM : ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "abc.com",
		Path:    "/pathA",
		RawPath: "user=Shahzeb",
	}
	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
