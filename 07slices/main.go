package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Start slices")
	// when you define size its array
	var fruits = []string{}
	fmt.Printf("type of fruits is %T", fruits)

	fruits = append(fruits, "apple")
	fruits = append(fruits, "banana")
	fruits = append(fruits, "apple")
	fruits = append(fruits, "banana")
	fmt.Println(fruits)
	var fruits2 = append(fruits[1:], fruits...)
	fmt.Println(fruits2)

	//make keyword
	highScores := make([]int, 4)
	highScores[0] = 10
	highScores[1] = 20
	highScores[2] = 30
	highScores[3] = 50

	fmt.Println(highScores)

	// if we add another it will out of bound
	//panic: runtime error: index out of range [4] with length 4
	//but append will reallocate memory

	//highScores[4] = 00
	highScores = append(highScores, 00)

	fmt.Println(highScores)

	//sort

	sort.Ints(highScores)
	fmt.Println(highScores)
	// remove value of slice on index

	//remoe value on index 2

	highScores = append(highScores[:2], highScores[3:]...)
	fmt.Println(highScores)
}
