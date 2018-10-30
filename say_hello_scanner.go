package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// without using variable
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("What is your name?")
	scanner.Scan()
	fmt.Println("Hello, ", scanner.Text(), ", nice to meet you!")
}
