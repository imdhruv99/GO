package main

import "fmt"

// func name(parameter type) return type of func {}
func greetings(s string) string {
	return "Hello " + s
}

func addition(num1, num2 int) int {
	return num1 + num2
	
}

func main() {
	s := "Dhruv"
	fmt.Println(greetings(s))

	num1 := 10
	num2 := 20
	fmt.Println(addition(num1, num2))
}