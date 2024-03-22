package main

import (
	"fmt"
	"os"
)

// range -> index,value

func main() {

	for _, value := range os.Args[1:] {
		fmt.Println("Value: ", value)

	}

}
