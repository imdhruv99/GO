package main

import "fmt"

func digits(num int, c chan int) {
	for num != 0 {
		digit := num % 10
		c <- digit
		num /= 10
	}
	close(c)
}

func calcSquares(num int, c chan int) {
	sum := 0
	ch := make(chan int)

	go digits(num, ch)

	for digit := range ch {
		sum += digit * digit
	}
	c <- sum
}

func calcCubes(num int, c chan int) {
	sum := 0
	ch := make(chan int)

	go digits(num, ch)

	for digit := range ch {
		sum += digit * digit * digit
	}
	c <- sum
}

func main() {
	fmt.Println("Enter Number...")
	var num int
	fmt.Scanln(&num)

	sqrt := make(chan int)
	cube := make(chan int)

	go calcSquares(num, sqrt)
	go calcCubes(num, cube)

	s, c := <-sqrt, <-cube

	fmt.Println("Final Output: ", s + c)
}