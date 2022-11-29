package main

import "fmt"

//main acts as an entry point
func main() {
	fmt.Println("WELCOME TO FUNCTIONS")
	hello()
	// //NOT ALLOWED TO WRITE FUNTION INSIDE FUNCTION
	// func hi()  {
	// 	fmt.Println("HI")

	// }
	func() {
		fmt.Println("IMMEDIATELY INVOKED FUNTION can only be defined inside a function")
	}()
	sum := adder(5, 6)
	fmt.Println("SUM", sum)

	fmt.Println("SUM", addALL(1, 2, 3, 4))
	res, greet := getSumAndGreeting(1, 2, 3, 4)
	fmt.Printf("RES %v  GREET %v", res, greet)
}

func hello() {
	fmt.Println("NAMASTE GOOO")
}

//returning multiple values
func getSumAndGreeting(values ...int) (int, string) {
	total := 0
	for _, i := range values {
		total += i
	}
	return total, "HELLO I AM SHAHZEB"
}

func addALL(values ...int) int {
	total := 0
	for _, i := range values {
		total += i
	}
	return total

}
func adder(n1 int, n2 int) int {
	return n1 + n2
}
