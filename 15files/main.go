package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("FILESSSSSSSSSSSSS")
	content := "THIS NEEDS TO BE IN A FILE"
	file, err := os.Create("./myFile")
	if err != nil {
		fmt.Println("ERRROR", err)
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		fmt.Println("ERRROR", err)
		panic(err)
	}
	fmt.Println("LENGTH", length)
	file.Close()
	readFile("./myFile")

}

func readFile(fileName string) {
	valueBytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("ERRROR", err)
		panic(err)
	}
	fmt.Println("TEXT DATA IS \n", valueBytes)
	fmt.Println("TEXT DATA IS \n", string(valueBytes))
}
