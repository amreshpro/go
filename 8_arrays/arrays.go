package main

import "fmt"

func main() {
	var nums [4]int 
	nums[0] = 10
	nums[1] = 20
	
	fmt.Println("nums[0]: ",nums[0])
	fmt.Println("nums[1]: ",nums[1])
	fmt.Println("nums[2]: ",nums[2])  //default value = 0
	fmt.Println("length: ",len(nums))
	
// 2d array
matrix:=[2][2]int{{3,4},{6,7}}
fmt.Println(matrix)

}