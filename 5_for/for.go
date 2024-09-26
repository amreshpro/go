package main

import "fmt"

// for -> only construct in go for looping
func main() {
	//no while keyword

	// while loop style in go
	i := 1
	for i <= 10 {
		fmt.Print(i, " ")
		i = i + 1
	}

	// classic loop
	for i := 0; i < 10; i++ {
		// println(i," ") // this println is for dev , not confirm to stay in go lang or removed
	
		fmt.Print(i*i, " ")
	}

	// infinite loop
	// for {
	// 	println("infinte loop")
	// }

}
