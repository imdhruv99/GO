package main

import "fmt"

func main() {

	// Assigning B to pointer of A
	a := 5
	b := &a
	fmt.Println("Value:", a, "Memory Adderess:", b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	// Reading value from Memory Address
	fmt.Println(*b)
	fmt.Println(*&a)

	// change value with pointer
	*b = 10
	fmt.Println(*b)
}