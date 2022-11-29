package main

import "fmt"

const EyeColor = "brown"

func main() {

	//Explicit type declare

	var name string = "Shaheb"
	fmt.Println(name)
	fmt.Printf("Hello FROM SHAHZEB : %T \n", name)

	//Implicit type declare
	var age = 25
	fmt.Println(age)
	fmt.Printf("Hello FROM SHAHZEB : %T \n", age)

	//no ar declare
	//:= walrus operator can only be used inside a function NOT outside
	is := true
	fmt.Println(is)
	fmt.Printf("Hello FROM SHAHZEB : %T \n", is)

	fmt.Println(EyeColor)
	fmt.Printf("Hello FROM SHAHZEB : %T \n", EyeColor)
}
