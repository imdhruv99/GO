package main

import "fmt"

func main() {

	// Define map
	emails := make(map[string]string)

	// Assign
	emails["Oggy"] = "oggy@gmail.com"
	emails["Tom"] = "tom@gmail.com"
	emails["Jerry"] = "jerry@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Tom"])

	// Delete from map
	delete(emails, "Oggy")
	fmt.Println(emails)

	//	Declare and Assign
	anime := map[string]string {"1": "Dragon Ball Super", "2": "Naruto"}
	fmt.Println(anime)

}