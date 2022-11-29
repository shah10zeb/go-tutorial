package main

import "fmt"

func main() {
	fmt.Println("Starting ARRAY")
	var fruitList [3]string
	fruitList[0] = "apple"
	fruitList[2] = "orange"

	//fruitList[1] is blank space

	fmt.Println("Fruitlist is", fruitList)
	//len
	fmt.Println("length of Fruitlist is", len(fruitList))
	var names = [3]string{"a", "b", "c"}
}
