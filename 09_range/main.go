package main

import "fmt"

func main() {

	// slice
	ids := []int{1, 3, 5, 7, 9, 11, 13, 15}

	// range
	for i, id := range ids {
		fmt.Printf("%d - ID: %d \n", i, id)
	}

	// `_` is used when we do not want to use variable
	for _, id := range ids {
		fmt.Printf("ID: %d \n", id)
	}

	// Sum of ids
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum is:", sum)


	// Range with map
	// Define map
	emails := make(map[string]string)

	// Assign
	emails["Oggy"] = "oggy@gmail.com"
	emails["Tom"] = "tom@gmail.com"
	emails["Jerry"] = "jerry@gmail.com"

	for key, email := range emails {
		fmt.Printf("%s: %s\n", key, email)
	}

	for k := range emails {
		fmt.Println("Key:", k)
	}
}