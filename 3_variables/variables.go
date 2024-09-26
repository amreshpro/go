package main

import "fmt"

// declared but unused variables through error in golang
func main() {
	var name string = "Amresh"
	var edu = "MCA"
	// shorthand
	age := 24 // better when declare + initialize
	fmt.Println("Hello, ", name)
	fmt.Println(edu)
	fmt.Println(age)

}
