package main

import  (
	"fmt" 
	"math"
	
	// importing our own package
	"github.com/imdhruv99/go_crash_course/03_packages/reverse"
)

func main() {
	fmt.Println("Hello Dhruv")
	fmt.Println(math.Abs(10-2.5356))

	fmt.Println(reverse.Reverse("hello"))

	fmt.Println(reverse.Reverse("VURHD"))
}