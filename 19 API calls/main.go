package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	fmt.Println("START WITH API")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("STATUS CODE", response.Status)
	fmt.Print("CONTENT LENGTH", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println("CONTENT", content)
	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)
	fmt.Println("BYTE COUNT", byteCount)
	fmt.Println("STRING =", responseString.String())
	//other way
	fmt.Println("CONTENT PARSED AS STRING", string(content))

}
