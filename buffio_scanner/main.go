package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Buffered IO Scanner Example")

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter text (type 'exit' to quit):")

	scanner.Scan()
	input := scanner.Text()

	fmt.Printf("You entered: %s\n", input)
}