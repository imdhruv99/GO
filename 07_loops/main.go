package main

import "fmt"

func main() {

	// Long Method
	i := 1
	for i <= 10 {
		fmt.Println(i)
		// i = i + 1
		i++
	}

	// Short Method
	for i:= 1; i <= 10; i++ {
		fmt.Printf("Number %d \n", i)
	}

	// FizzBuzz
	FizzBuzz := 0
	Fizz := 0
	Buzz := 0
	for i := 1; i <= 100; i++ {
		if i % 15 == 0 {
			fmt.Println("FizzBuzz")
			FizzBuzz = FizzBuzz + 1
		}	else if i % 3 == 0 {
			fmt.Println("Fizz")
			Fizz = Fizz + 1
		}	else if i % 5 == 0 {
			fmt.Println("Buzz")
			Buzz = Buzz + 1
		}	else {
			fmt.Println(i)
		}
	}
	fmt.Printf("FizzBuzz count is %d \n", FizzBuzz)
	fmt.Printf("Fizz count is %d \n", Fizz)
	fmt.Printf("Buzz count is %d", Buzz)
}