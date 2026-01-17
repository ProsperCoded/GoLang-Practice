package main
import 'fmt'

func main(){
	// Boolean
	var boolVar bool = true

	// Integer types
	var intVar int = 42
	var int8Var int8 = 127
	var int16Var int16 = 32767
	var int32Var int32 = 2147483647
	var int64Var int64 = 9223372036854775807

	// Unsigned integer types
	var uintVar uint = 42
	var uint8Var uint8 = 255
	var uint16Var uint16 = 65535
	var uint32Var uint32 = 4294967295
	var uint64Var uint64 = 18446744073709551615

	// Float types
	var float32Var float32 = 3.14
	var float64Var float64 = 3.141592653589793

	// Complex types
	var complex64Var complex64 = 1 + 2i
	var complex128Var complex128 = 3 + 4i

	// String
	var stringVar string = "Hello, Go!"

	// Byte (alias for uint8)
	var byteVar byte = 65

	// Rune (alias for int32)
	var runeVar rune = 'A'

	// Array
	var arrayVar [3]int = [3]int{1, 2, 3}

	// Slice
	var sliceVar []int = []int{1, 2, 3}

	// Map
	var mapVar map[string]int = map[string]int{"a": 1, "b": 2}

	// Struct
	type Person struct {
		name string
		age  int
	}
	var structVar Person = Person{"John", 30}

	// Interface
	var interfaceVar interface{} = "anything"
 
	// Pointer
	var pointerVar *int = &intVar

	// Function
	var funcVar func(int) int = func(x int) int { return x * 2 }

	var age, name  = 4, "prosper";
	fmt.Print(age)
	fmt.Print(name)
	// fmt.Println(boolVar, intVar, float64Var, stringVar, sliceVar, mapVar, structVar)
}