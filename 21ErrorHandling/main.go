// LAST LESSON: ERROR HANDLING!
package main

import (
	"errors" // imported to create error messages
	"fmt"
	"math"
)

// The error interface
type error interface {
	Error() string
}

// Sqrt calculate the square root of a function
func Sqrt(value float64) (float64, error) {
	if value < 0 {
		return 0, errors.New("Negative number passed to Sqrt!")
	}
	// Return the square root, and a nil error
	return math.Sqrt(value), nil
}

func main() {

	// Test the Sqrt function with a negative value

	// result, err := Sqrt(-5)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(result)
	// }

	// Test the Sqrt function with a negative value

	result, err := Sqrt(5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
