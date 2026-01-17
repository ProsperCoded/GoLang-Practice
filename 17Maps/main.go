// MAPS
package main

import "fmt"

func main() {
	fmt.Println("Maps")
	// Initialize the map in Go
	// var myMap map[string]string
	// myMap = make(map[string]string)
	// fmt.Println("Map value is:", myMap)

	// Declare and initialize a map
	countryCapitalMap := make(map[string]string)
	// Insert key-value pairs into the map
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"
	countryCapitalMap["USA"] = "Washington"

	// Print capitals using keys
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}
	capital, ok := countryCapitalMap["USA"]
	if ok {
		fmt.Println("capital of USA is ", capital)
	} else {
		fmt.Println("Does not Exist!")
	}

	// Delete an entry
	delete(countryCapitalMap, "France")
	fmt.Println("Entry for France is deleted!")

	// Update the map state
	fmt.Println("Updated map:")
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}
}
