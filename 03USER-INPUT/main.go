package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//COMMA OK SYNTAX COMMA ERROR SYNTAX
	//BUFF IO package
	welcome := "WELCOME SHAHZEB"
	fmt.Println(welcome)

	//creating a reader
	//from here to read is ++ os library Stdin
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("ENTER AGE")

	//where ever this reader is reading from
	//keep reading until a /n is encountered
	input, _ := reader.ReadString('\n')

	fmt.Println("YOu TYPED, ", input)
	//type of input will be string even if we enter numer
	fmt.Printf("YOur TYPE , %T", input)

	//converting type

	// age, err := strconv.ParseInt(input, 10, 64)
	//the above statementcannot conert 25\n to number
	// we will need to trim \n
	age, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)

	if err != nil {
		fmt.Println("ERR: ", err)
	} else {
		fmt.Println("NExt year you will turn  ", age+1)
	}
}
