package main

import "fmt"

func main() {
	fmt.Println("STARTING LOOPS")

	days := []string{"SUN", "MON", "TUE", "WED", "THU", "FRI", "SAT"}
	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for index, day := range days {

		fmt.Printf("INDEX: %v  VALUE: %v   ", index, day)

	}

	roughValue := 2
	for roughValue < 10 {
		fmt.Println("VAlue is : ", roughValue)
		roughValue++
	}
}
