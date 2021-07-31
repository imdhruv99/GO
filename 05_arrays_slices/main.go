package main

import "fmt"

func main() {

	// Defining Array
	var fruits [2] string
	
	// Assigning values
	fruits[0] = "Apple"
	fruits[1] = "Orange"

	fmt.Println(fruits)
	fmt.Println(fruits[0])

	// Declaring and Assigning
	companies := [2] string {"Apple", "Amazon"}
	fmt.Println(companies)

	// Slice
	car := []string{"Tata", "Maruti", "Honda", "BMW"}
	fmt.Println(car)
	fmt.Println(len(car))
	fmt.Println(car[1:3])
	
}