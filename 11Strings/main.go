// Strings
package main

import (
	"fmt"
	"strings"
)

func main() {
	// Creating Strings
	var greeting = "Hello, World!"
	fmt.Printf("normal string: ")
	fmt.Printf("%s", greeting)
	fmt.Printf("\n")

	// Display the hexadecimal byte values of the string
	fmt.Printf("HEX bytes: ")
	for i := 0; i < len(greeting); i++ {
		fmt.Printf("%x", greeting[i])
	}
	fmt.Printf("\n")

	// Measuring String Length
	fmt.Printf("The length of this string is : ")
	fmt.Println(len(greeting))

	// String concatentation
	// create a slice of strings for concat.
	fruits := []string{"apples", "oranges", "and bananas"}
	// Concat. using strings.Join
	fmt.Println(strings.Join(fruits, " "))
}
