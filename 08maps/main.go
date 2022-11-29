package main

import "fmt"

func main() {
	fmt.Println("maps in GOOOOOOOOOOOOOOOO")
	//map[key]value
	nums := make(map[string]string)
	nums["zero"] = "ZERO"
	nums["one"] = "ONE"
	nums["two"] = "TWO"
	fmt.Println("NUMS", nums)
	//delete
	delete(nums, "one")
	fmt.Println(nums)
	//loops

	for key, value := range nums {
		fmt.Printf("Key : %v, Value :%v", key, value)
	}

	for _, value := range nums {
		fmt.Printf("Value :%v", value)
	}

}
