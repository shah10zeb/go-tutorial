package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	//alias for NAME
	Name     string `json:"coursename"`
	Price    int
	Platform string `json:"website"`
	//dont list platform
	Password string `json:"-"`
	//omit null nil
	Tags []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("JSON JSON JSON")
	// EncodeJSON()
	DecodeJSON()

}

func EncodeJSON() {
	lcoCourses := []course{
		{"go", 200, "abc.com", "123", []string{"go", "basics"}},
		{"go1", 200, "abc.com", "1234", []string{"go", "medium"}},
		{"go2", 200, "abc.com", "1235", nil},
	}

	finalJSON, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s \n", finalJSON)
}
func DecodeJSON() {
	jsonDataFromWeb := []byte(`{
		"coursename": "go",
		"Price": 200,
		"website": "abc.com",
		"tags": [
				"go",
				"basics"
		]
}`)

	var lcoCourse course
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("VALID ", checkValid)
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		//print interfces use %#v
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("NOT VALID")
	}

	//some case: only key value

	var myOnlineData map[string]interface{}

	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("KEY IS %v and value is %v and TYPE is %T", k, v, v)
	}
}
