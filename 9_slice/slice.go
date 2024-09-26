package main

import "fmt"

// slices are dynamic arrays
// most used construct in go
// many useful method
func main() {
	// not specified size because it is dynamic
	//unintialized slice is nil
	var nums []int // nil
	var num2 [2]int

	nums = append(nums, 5)
	num2[0] = 12
	num2[1] = 34

	fmt.Println(nums)
	fmt.Println(num2)

	// famous way
	var numspro = make([]int, 2) // make(type of slice,size,capacity)
	fmt.Println(numspro)         // [0,0] not nil
	// default size == capcity
	// capacity  --> maximum numbers of elements can fit 
	fmt.Println(cap(numspro)) //2
	numspro = append(numspro, 45)
	numspro = append(numspro, 57)

fmt.Println(numspro)




}
