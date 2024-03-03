package main

import (
	"fmt"
)

var (
	foo              = "Foo"
	firstName string = "Amresh"
	lastName  string
)

const (
	Pi = 3.14
)
//global scope
func main() {
	//local scope
	version := 1            //infer int
	updatedVersion := "V12" //infer string
	fmt.Println(version)
	fmt.Println(updatedVersion)
}
