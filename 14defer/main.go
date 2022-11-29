package main

import "fmt"

func main() {
	defer fmt.Println("ONE")
	defer fmt.Println("TWO")
	//LIFO for multiple defer
	fmt.Println("Three")
	myDefer()

	//
}

func myDefer() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}
