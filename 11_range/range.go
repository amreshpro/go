package main

import "fmt"

func main() {
	num := []int{6, 7, 8, 9}
	// for i := 0; i < len(num); i++ {
	// 	fmt.Println(num[i])
	// }

	sum := 0
	for index, n := range num {
		sum = sum + n
		fmt.Println(index)
	}

	fmt.Println(sum)
}
