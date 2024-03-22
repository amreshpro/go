package main

import (
	"fmt"
)

// Control Structure
// - for loop
// - range
// - break
// - continue

func main() {
	// while loop in go - no while keyword in go
	var t = 1
	for t <= 10 {
		fmt.Println("Value: ", t)
		t++
	}

	// infinite loop 
	for {
		fmt.Println("I Love Go!")
	}
}
