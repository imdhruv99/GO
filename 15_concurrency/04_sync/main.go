package main

import (
	"fmt"
	"sync")


var x = 0
func increment(wg *sync.WaitGroup, c chan bool) {
	c <- true
	x = x + 1
	<- c
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	c := make(chan bool, 1)

	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, c)
	}

	w.Wait()

	fmt.Println("Final value of X", x)

}