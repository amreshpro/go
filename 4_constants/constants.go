package main

import "fmt"

const name = "Amresh"

// edu := "MCA" // not allowed this syntax only work inside the functions
func main() {

	const PI = 3.14
	fmt.Println("PI = ", PI)

	// variable grouping
	const (
		name     = "mongodb"
		host     = "localhost"
		port     = 3456
		password = "flsdjlfwqepd"
	)
	fmt.Println(name) //mongodb
	fmt.Println(host, ":", port)

}
