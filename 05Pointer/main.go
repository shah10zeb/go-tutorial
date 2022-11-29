package main

import "fmt"

func main() {

	fmt.Println("POINTER START")

	var pointer1 *int
	fmt.Printf("TYPE OF POINTER1  :%T ", pointer1)

	fmt.Println(pointer1)

	var num int = 10
	var pointer2 = &num
	fmt.Println(&num)
	fmt.Printf("TYPE OF &num  :%T ", &num)
	fmt.Printf("TYPE OF POINTER2  :%T ", pointer2)
	fmt.Printf("TYPE OF POINTER2  :%T ", *pointer2)

	fmt.Println(pointer2)

	*pointer2 = *pointer2 + 5

	fmt.Println("After pointer operation num is", num)
}
