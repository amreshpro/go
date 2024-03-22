package main

import (
	"fmt"
	"os"
	"strings"
)
func main() {
	fmt.Println("Args[0]: ",os.Args[0])
	// for i := 1; i < len(os.Args); i++ {
	// 	fmt.Println("Value : ",os.Args[i])
	// }

	fmt.Println(strings.Join(os.Args[1:]," "))

}