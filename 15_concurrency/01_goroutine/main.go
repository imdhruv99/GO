package main

import (
	"fmt"
	"time")


func numbers() {
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Millisecond * 200)
		fmt.Printf("%d \n", i)
	}
}


func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(time.Millisecond * 400)
		fmt.Printf("%c \n", i)		
	}
}

func main() {

	// running first go routine
	go numbers()

	// running second fo routine
	go alphabets()

	time.Sleep(time.Millisecond * 5000)

	fmt.Println("Main Terminated!!!")

}