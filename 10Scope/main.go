// Scopes in Go.

package main

import "fmt"

// 2. Global Variables
// Global variable declaration
var g int
var x int = 25
var y int = 55

func main() {
	// fmt.Println("SCOPES in Go!")
	// 1. Local Variables
	// Local variable declaration
	var a, b int

	// var x int

	// Variable initialization
	a = 20
	b = 20
	g = a + b

	var x int = 100
	var y int = 1000
	// x := 100

	// Displaying the values of the variables
	fmt.Printf("a = %d, b = %d, global variable g = %d\n", a, b, g)
	fmt.Printf("x = %d\n", x)
	fmt.Printf("This y local variable shadows the global variable = %d\n", y)
}
