// SLICES
package main

import "fmt"

func main() {
	fmt.Println("Slices")
	// declaration of slice of integers
	// var numbers []int
	// numbers = make([]int, 5)
	// numbers := make([]int, 5)
	// initialization of slice of integers with length 5 filled with [0,0,0,0,0]
	// fmt.Println("Slice is : ", numbers)

	// Sublicing
	// numbers2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	// subSlice1 := numbers2[1:4]
	// subSlice2 := numbers2[:3]
	// subSlice3 := numbers2[4:]

	// fmt.Println("subSlice1:", subSlice1)
	// fmt.Println("subSlice2:", subSlice2)
	// fmt.Println("subSlice3:", subSlice3)

	// Append() and Copy()

	numbers := []int{}
	// Appending elements
	numbers = append(numbers, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	numbers1 := make([]int, len(numbers), cap(numbers)*2)
	// Copying contents
	copy(numbers1, numbers)

	fmt.Println("numbers slice:", numbers)
	fmt.Println("numbers1 slice after copy:", numbers1)

	// Nil Slice
	var num []int // nil slice
	fmt.Println("Number:", num)
	if num == nil {
		fmt.Println("Slice is nil")
	}
}
