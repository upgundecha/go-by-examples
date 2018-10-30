package main

import (
	"fmt"
)

func main() {

	// with variable

	var name string
	fmt.Println("What is the input string?")
	fmt.Scanln(&name)
	fmt.Println(name, "has", len(name), "characters")
}
