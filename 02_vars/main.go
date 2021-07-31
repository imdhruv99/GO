package main

import "fmt"

func main() {

	// MAIN TYPES
	// string
	// bool
	// int
	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Creating variable using var keyword
	var name string = "Dhruv Prajapati"
	var age int32 = 22
	const isHappy = true


	fmt.Println(name, age, isHappy)

	// Getting Type
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isHappy)

	// Creating variable using shorthand method, It only works when you wanna define inside the function
	work, interest := "DevOps Engineer", "Machine Learning Deep Learning"
	fmt.Println(work, interest)

	height := 1.778
	fmt.Printf("%T", height)

}