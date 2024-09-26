// [n]T  --> array of n values of type T
// var a [10]int
package main

import "fmt"

func Array2() {
	var a [10]int
	fmt.Println("a(default): ", a[0]) // 0 -> default value is zero in go array
	a[0] = 10
	a[1] = 20
	fmt.Println(a)
	fmt.Println(a[1])

	var str [10]string
	fmt.Println("a(default): ", str[0]) // empty -> default value of string in array
	str[0] = "Hello"
	str[1] = "Amresh"
	fmt.Println(str)
	fmt.Println(str[1])

}
