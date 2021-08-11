package main

import (
	"fmt"
	"strings"
)

func main() {

	var num int
	fmt.Println("Enter number of lines:")
	fmt.Scanln(&num)

	for i := num; i >= 1; i-- {
		fmt.Printf("%s\n", strings.Repeat("*", i))
	}

}