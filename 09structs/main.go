package main

import "fmt"

func main() {
	fmt.Println("STarting Structs")
	//there are no inheritance
	//there is no hidden parent
	//structs==== Classes

	shahzeb := User{"Shahzeb", "abc.com", 25, true}
	fmt.Println(shahzeb)
	fmt.Printf("SHAHEB DETAILS are %+v \n", shahzeb)
	fmt.Printf("Name is %v \n", shahzeb.Name)

}

//Defining Struct
type User struct {
	//PUBLIC Fist Capital Letter
	Name   string
	Email  string
	Age    int
	Status bool
}
