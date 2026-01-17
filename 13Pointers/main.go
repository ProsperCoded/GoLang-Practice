// Understanding Pointers
// A pointer stores the memory address of another variable.

// Declaring Pointers
// var ip *int
// var fp *float32

// Basic Pointer Operations
// Storing the address: & to get the address of a variable.
// Dereferencing: * to access the value at the pointer's address

package main

import "fmt"

func main() {
	var a int = 20
	var ip *int
	ip = &a
	fmt.Printf("Address of a: %x\n", &a)
	fmt.Printf("Address stored in ip: %x\n", ip)
	fmt.Printf("Value at *ip: %d\n", *ip)

	// Nil Pointers
	var ptr *int
	fmt.Printf("The value of ptr is : %x\n", ptr)
}

// Key Takeaways
// 1 Pointers store memory addresses, enabling direct access to data.
// 2 Use & to get an address and * to access the value at that address.
// 3 Nil pointers represent uninitialized pointers.
