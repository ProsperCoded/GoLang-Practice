package main

import (
	"encoding/json"
	"fmt"
)

// Note: lowercase will make struct private, same with the fields
type Employee struct {
	name     string
	age      int
	position string
	salrary   float64

	address Address
}

// struct with tags 
type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	ZipCode string  `json:"zip_code"`
}

// will serve as a function to mutate the struct
func (e *Employee) updateName(newName string) {
	e.name = newName
}

// function to print address
func (a Address) printAddress() {
	fmt.Printf("Address: %s, %s, %s\n", a.Street, a.City, a.ZipCode)
}

func main() {
	address := Address{
		Street:  "123 Main St",
		City:    "Metropolis",
		ZipCode: "12345",
	}
	employee1 := Employee{
		name:     "John Doe",
		age:      30,
		position: "Software Engineer",
		salrary:   60000.50,
		address:  address,
}

	fmt.Println("Employee Details:")
	fmt.Printf("Name: %s\n", employee1.name)
	fmt.Printf("Age: %d\n", employee1.age)
	fmt.Printf("Position: %s\n", employee1.position)
	fmt.Printf("Salrary: %.2f\n", employee1.salrary)
	employee1.address.printAddress()

	// Ananymous Struct
	employee2 := struct {
		name     string
		age      int
		position string
		salrary   float64
	}{
		name:     "Jane Smith",
		age:      28,
		position: "Data Analyst",
		salrary:   55000.75,
	}

	// Mutating struct with pointers

	employeePtr := &employee1;
	employeePtr.age = 31; 

	// Testing tags defined ab ove 
	jsonData, _ := json.Marshal(employee1.address)
	fmt.Println("JSON Address:", string(jsonData))
}
