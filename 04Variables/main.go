// Understanding Variables Declarations in Go
package main

import "fmt"

func main() {
	// way1: Declare and assign on 2 different lines
	var mango string = "This is a big mango!"
	var weight int = 54

	// way2: Declare and assign on the same line
	var height int = 23

	fmt.Println("Mango:", mango)
	fmt.Println("weight:", weight)
	fmt.Println("height:", height)

	// way3 (shorthand)
	// shorthand notation :=
	// Type inference 
	// var age = 23
	age := 54
	city := "Washington"
	fmt.Println("My age is:", age) 
	fmt.Println("My city is:", city)
	
	// Multiple declaration and init. on same line
	var apples, oranges int = 23,78
	fmt.Println("I have", apples, "apples and", oranges, "oranges")

	var fruits = apples + oranges
	fmt.Println("fruits:", fruits)

	var windows, mac, linux string = "Windows is ok\n", "Mac is meh\n", "Linux is the GOAT!\n"
	print(windows, mac, linux)


	// Static Type declaration
	var x float64 = 20.5
	fmt.Println(x)
	fmt.Printf("x is of type: %T\n", x)
	
	
	// Dynamic Type declaration
	y := 89
	fmt.Println(y)
	fmt.Printf("y is of type: %T\n", y)
	
	// Mixed Variable Declaration
	var a, b, c = 758.52, 8, "foobar"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("a is of type: %T\n", a)
	fmt.Printf("b is of type: %T\n", b)
	fmt.Printf("c is of type: %T\n", c)
}
