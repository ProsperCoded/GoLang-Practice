// Interfaces
package main

import (
	"fmt"
	"math"
)

// Define an interface
type Shape interface {
	area() float64
}

// Define a Circle struct
type Circle struct {
	radius float64
}

// Define a Rectangle struct
type Rectangle struct {
	width, height float64
}

// Implement area method for Circle
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Implement area method for Rectangle
func (r Rectangle) area() float64 {
	return r.width * r.height
}

// Function to get area of any shape!
func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	fmt.Println("Interfaces")

	// // Define an interface
	// type thisisaninterface interface {
	// 	// methods
	// 	// MethodName1 ReturnType
	// 	// MethodName2 ReturnType
	// 	// MethodName3 ReturnType
	// }

	// //Define a struct
	// type thisisastruct struct {
	// 	// variables
	// }

	circle := Circle{radius: 5}
	rectangle := Rectangle{width: 10, height: 5}

	fmt.Printf("Circle area: %f\n", getArea(circle))
	fmt.Printf("Rectangle area: %f\n", getArea(rectangle))
}
