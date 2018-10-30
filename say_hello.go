package main

import (
	"fmt"
)

func main() {

	// with variable

	var name string
	fmt.Println("What is your name?")
	fmt.Scanln(&name)
	fmt.Println("Hello, ", name, ", nice to meet you!")
}
