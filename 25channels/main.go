package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func CalculateValue(intChan chan int) {
	rand.Seed(time.Now().Unix())
	random := rand.Intn(20)
	// random := 5
	log.Println(random)
	intChan <- random
}

func main() {
	log.Println(("Hello Channels"))
	fmt.Println("HELLO")

	intChan := make(chan int)

	defer close(intChan)
	//important to close as there can be  limited no of channel open
	//go routines run parallelly

	go CalculateValue(intChan)
	go CalculateValue(intChan)

	numFromChannel := <-intChan
	log.Println(numFromChannel)
}
