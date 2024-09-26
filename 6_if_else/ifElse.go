package main

import "fmt"

func main() {

	if isAdult := 18; isAdult >= 18 {
		fmt.Println("Person is adult")
	} else if isAdult >= 16 {
		fmt.Println("Persoon is 16 or 16+")
	} else {
		fmt.Print("Person is not adult")
	}
}
