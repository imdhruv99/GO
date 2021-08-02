package main

import (
	"fmt"
	"time")


func process(c chan string) {
	time.Sleep(time.Millisecond * 10500)
	c <- "Process is Successful"
}

func main() {
	c := make(chan string)

	go process(c)

	for {
		time.Sleep(time.Millisecond * 1000)

		select {

		case v := <- c:
			fmt.Println("Value Received:", v)
			return
		
		default:
			fmt.Println("No Value Received")
		}
	}
}