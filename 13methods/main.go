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
	shahzeb.GetStatus()
	shahzeb.newMail()
	fmt.Printf("SHAHEB DETAILS are %+v \n", shahzeb)

}

//Defining Struct
type User struct {
	//PUBLIC Fist Capital Letter
	Name   string
	Email  string
	Age    int
	Status bool
	//Private
	// oneAge int
}

//Method
//this passes as a copy
func (u User) GetStatus() {
	fmt.Println("IS USER ACTIVE: ", u.Status)
}

func (u User) newMail() {
	u.Email = "test@Go.dev"
	fmt.Println("USER MAIL: ", u.Email)
}

//to pass an original object we need to pass pointer or reference
