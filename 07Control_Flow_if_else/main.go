//  Control Flow:
// Conditionals (If Else statement)

package main

import "fmt"

func main() {
	// 1- If Statement
	// Executes a block of code if the condition is true.
	// age := 18
	// if age >= 18 {
	// 	fmt.Println("You are eligible to vote")
	// 2- Else
	// } else {
	// 	fmt.Println("You are not eligible to vote")
	// }

	// 3- Nested Statements
	score := 95
	if score >= 90 {
		fmt.Println("Grade:A")
	} else if score >= 80 {
		if score >= 85 {
			fmt.Println("Grade:B+")
		} else {
			fmt.Println("Grade:B")
		}
	} else {
		fmt.Println("Grade:C")
	}
}
