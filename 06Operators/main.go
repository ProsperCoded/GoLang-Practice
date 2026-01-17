// Lesson6: Operators in Go Language
// An operator is a symbol that performs specific mathematical manipulations.

// Arithmetic Operators
// Relational Operators
// Logical Operators
// Bitwise Operators
// Assignment Operators
// & Miscellaneous Operators

package main

import "fmt"

func main() {
	// Arithmetic Operators
	// A := 10
	// B := 20
	// C := 11
	// D := 3
	// fmt.Println("A + B =", A+B) // Addition 30
	// fmt.Println("A - B =", A-B) // Subtraction -10
	// fmt.Println("A * B =", A*B) // Multiplication 200
	// fmt.Println("A / B =", A/B) // Division 5.0
	// fmt.Println("C % D =", C%D) // Modulus 2
	// A++
	// fmt.Println("A++ =", A) //Increment  11
	// A--
	// fmt.Println("A-- =", A) //Decrement  10
	//********************************

	// Relational operators - Comparison Operators
	//Equal to
	// fmt.Println("A == B:", A == B) // Equals - false
	// //Not equal to
	// fmt.Println("A != B:", A != B) // Equals - true
	// // Greater than
	// fmt.Println("A > B:", A > B) // Equals - false
	// // Less than
	// fmt.Println("A < B:", A < B) // Equals - true
	// // Greater than or equal to
	// fmt.Println("A >= B:", A >= B) // Equals - false
	// // Less than or equal to
	// fmt.Println("A <= B:", A <= B) // Equals - true

	// // Logical operators - Boolean Operators
	// A := true
	// B := false

	// //Logical operator AND
	// fmt.Println("A && B:", A && B) // Equals false
	// //Logical operator OR
	// fmt.Println("A || B:", A || B) // Equals true
	// //Logical operator NOT
	// fmt.Println("!A:", !A) // Equals false
	// fmt.Println("!B:", !B) // Equals true
	//********************************

	// Assignment Operators
	// A := 10
	// B := 20

	// A += B                    // A = A + B
	// fmt.Println("A += B:", A) // Equals 30

	// A -= B                    // A = A - B
	// fmt.Println("A -= B:", A) // Equals 10

	// ************************************

	// Bitwise operators
	// C := 30
	// 30 = 11110
	// 2 = 00010

	// // Bitwise AND
	// fmt.Println("C & 2 =", C&2) // Output: 2

	// // Bitwise OR
	// fmt.Println("C | 2 =", C|2) // Output: 30

	// // Bitwise XOR
	// fmt.Println("C ^ 2 =", C^2) // Output: 28

	// // Bitwise NOT
	// fmt.Println("~C =", ^C) // Output: -31

	// // Bitwise left shift
	// fmt.Println("C << 1 =", C<<1) // Output: 60

	// // Bitwise right shift
	// fmt.Println("C >> 1 =", C>>1) // Output: 15

	// ************************************

	// Miscellaneous operators
	A := 10

	// Address-of operator
	ptr := &A
	fmt.Println("Address of A:", ptr) // Output: (address of A)

	// Pointer dereference operator
	fmt.Println("Value of *ptr:", *ptr) // Output: 10
}
