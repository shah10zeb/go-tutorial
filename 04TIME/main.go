package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("LETS LEARN TIME")
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 "))
	fmt.Println(presentTime.Format("01-02-2006 Monday"))
	fmt.Println(presentTime.Format("01-02-2006 Monday 15:04:05"))
	fmt.Println(presentTime.Format("01-02-2006 Monday 03:04:05"))

	//REERTING STRING TO DATE

	createdDate := time.Date(2020, time.August, 20, 1, 28, 30, 0, time.UTC)
	fmt.Println(createdDate)

}
