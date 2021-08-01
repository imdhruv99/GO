package main

import "fmt"

func main() {

	x := 10
	y := 15

	// If else
	if x <= y {
		fmt.Printf("%d is less than or equal to %d", x, y)
	} else {
		fmt.Printf("%d is less than or equal to %d", y, x)
	}

	fmt.Printf("\n")

	// If Else If
	color := "blue"

	if color == "red" {
		fmt.Printf("color is red")
	} else if color == "blue" {
		fmt.Printf("color is blue")
	} else {
		fmt.Printf("color is not blue or red")
	}

	fmt.Printf("\n")

	// switch
	switch color {
		case "red":
			fmt.Println("color is red")
		case "blue":
			fmt.Println("color is blue")
		default:
			fmt.Println("color is not blue or red")
	}
}