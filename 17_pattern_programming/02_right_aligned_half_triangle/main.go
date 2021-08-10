package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var num int
	fmt.Println("Enter number of lines:")
	fmt.Scanln(&num)

	fmtString := "%" +strconv.Itoa(num) + "s\n"

	for i := 1; i <= num; i++ {
		fmt.Printf(fmtString, strings.Repeat("*", i))
	}

}