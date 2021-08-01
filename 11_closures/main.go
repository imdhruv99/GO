package main

import "fmt"

func addition() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	sum := addition()
	for i := 0; i <= 20; i++ {
		fmt.Println(sum(i))
	}
}