// Data Types
// Integers, floating points, booleans, complex numbers, and strings.
package main
import "fmt"
func main() {

	// Section1: Integers
	// Section1-A: signed integers
	var i int
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64

	// Assigning values to integer variables
	i = -128
	i8 = 127
	i16 = -32768
	i32 = 2147483647
	i64 = -9223372036854775808
	
	// Section1-B: unsigned integers
	var u uint
	var u8 uint8
	var u16 uint16
	var u32 uint32
	var u64 uint64

	// Assigning values to unsigned integer variables
	u = 255
	u8 = 255
	u16 = 65535
	u32 = 4294967295
	u64 = 18446744073709551615

	// Printing the signed and unsigned integer variables to the console
	fmt.Println("Signed integers:", i, i8, i16, i32, i64)
	fmt.Println("Unsigned integers:", u, u8, u16, u32, u64)
	
	// Section2: Floating Point
	// Section2-A : Float32
	// This is used for less precise calculations
	var f32 float32 = 10.6
	// Section2-B : Float64
	var f64 float64 = 10.6
	// This is used for more precise calculations

	// Printing the floating point values
	fmt.Println("FLOAT32:", f32)
	fmt.Println("FLOAT64:", f64)

	// Demonstrating the diff in precision between the f32 and f64
	var HP float64 = 10123456789012345
	var LP float32 = 10123456789012345
	fmt.Println("High precision float64:", HP)
	fmt.Println("Low precision float32:", LP)

	// Section3: Boolean Data Type
	var isActive bool = true
	var isOn bool = false

	fmt.Println("Is Active:", isActive)
	fmt.Println("Is On:", isOn)

	// Section4: Complex Data Type
	var CN1 complex128 = complex(5,10)
	var CN2 complex64 = complex(2,7)
	// print(CN1)
	// print(CN2)
	fmt.Println("CN1: ", CN1)
	fmt.Println("CN2: ", CN2)

	// Section5: String Data Type
	var name string = "Kermet the frog!"
	fmt.Println("My name is:", name)
}