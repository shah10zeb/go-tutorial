package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("IF ELSE START")

	loginCount := 3
	var result string
	if loginCount > 10 {
		result = "REGULAR USER"
	} else if loginCount > 5 {
		result = "USER"
	} else {
		result = "NOT A USER"
	}
	fmt.Println(result)

	if loginCount%2 == 0 {
		fmt.Println("NUMBER IS EVEN")
	}

	if num := 3; num < 10 {
		fmt.Println("LESS THAN 10")
	} else {
		fmt.Println("NOT LESS THAN 10")

	}
	// comma OK syntax (error handling)
	// if err != nil {

	// }

	fmt.Println("Starting SWITCH")
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("ONE")
	case 2:
		fmt.Println("TWO")
	case 3:
		fmt.Println("THREE")
		fallthrough
		//will also execute the next case
	case 4:
		fmt.Println("FOUR")
	case 5:
		fmt.Println("FIVE")
	case 6:
		fmt.Println("SIX")
	default:
		fmt.Println("WWWWWWWWWWHATTTTTT")
		//case "abc"
	}
}
