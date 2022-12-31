package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	fmt.Println("START WITH API")
	// PerformGetRequest()
	// PerformPostRequest()
	PerformPostFormRequest()
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

func PerformPostRequest() {
	const myurl = "http://localhost:8000/post"
	//fake json payload
	reBody := strings.NewReader(`
	{
		"name":"shahzeb",
		"age":"27",
		"money":"null"
	}`)
	response, err := http.Post(myurl, "application/json", reBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println("RESPONSE", response)
	fmt.Println("RESPONSE CONTENT", content)
	fmt.Println("RESPONSE CONTENT", string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/post"
	data := url.Values{}
	data.Add("Name", "SHAHZEB")
	data.Add("AGE", "27")
	data.Add("MONEY", "null")

	res, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	fmt.Println("RESPONSE", res)
	content, _ := ioutil.ReadAll(res.Body)
	fmt.Println("CONTENT ", content)
	fmt.Println("CONTENT STRING", string(content))

}
