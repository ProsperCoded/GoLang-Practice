// RECURSION
package main

import "fmt"

func main() {
	fmt.Println("Recursion")
	fmt.Println("Factorial of 5:", factorial(5))
	fmt.Println("Fibonacci of 6:", fibonacci(6))
	fmt.Println("Fibonacci of 7:", fibonacci(7))
}

// Basic structure fo recursive function:

// Base Case: Stops the recursion [preventing the infinite loop]
// Recursive Case:Function calls itself

// Factorial function

func factorial(n int) int {
	if n == 0 { // base case
		return 1
		// 5! = 5 * 4 * 3 * 2 * 1 = 120
	} else {
		return n * factorial(n-1) // recursive case
	}
}

// Fibonacci sequence
func fibonacci(n int) int {
	if n <= 1 { // base case
		return n
	} else {
		return fibonacci(n-1) + fibonacci(n-2) //recursive case
	}
}

// fibonacci(6) // 8
// 0 1 1 2 2 5 8 13 21 34 55 89 ....
// 0 1 2 3 4 5 6 7 8 9

/*
 F(6)
 ├── F(5) (calculates to 5)
 │   ├── F(4) (calculates to 3)
 │   │   ├── F(3) (calculates to 2)
 │   │   │   ├── F(2) (calculates to 1)
 │   │   │   │   ├── F(1) (returns 1)
 │   │   │   │   └── F(0) (returns 0)
 │   │   │   └── F(1) (returns 1)
 │   │   └── F(2) (calculates to 1)
 │   └── F(3) (calculates to 2)
 └── F(4) (calculates to 3)
*/
