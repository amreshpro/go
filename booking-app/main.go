package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello GO")

	var name string
fmt.Scan(&name)
	fmt.Println("Hello "+name)
}
