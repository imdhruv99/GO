package main

import (
	"fmt"
	"strconv"
)

// two ways to define struct
type Person struct {
	// Name string
	// City string
	// Gender string
	// Age int
	
	Name, City, Gender string
	Age int
}

// Value Reciever
func (p Person) greetings() string {
	return "Hello, My Name is " + p.Name + ". My Age is " + strconv.Itoa(p.Age) + ". and I am from " + p.City 
}

// Pointer Reciever
func (p *Person) hasBirthday() {
	p.Age++

}

func main() {

	// Init Person struct with using Property Name
	person1 := Person{
		Name: "Dhruv",
		City: "Ahmedabad",
		Gender: "Male",
		Age: 22 }
	
	fmt.Println(person1)
	fmt.Println(person1.Name)

	// Init Person struct without using Property Name
	person2 := Person{"Hetvi", "Ahmedabad", "Female", 22}
	
	fmt.Println(person2)

	// Pointer Reciever
	person1.hasBirthday()
	
	// Value Reciever
	fmt.Println(person1.greetings())
}