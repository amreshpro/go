package main

import("fmt")
//<- the data flows in the direction of the arrow.
//Like maps and slices, channels must be created before use.
//ch:=make(chan int)
// ch<-v // send v to channel ch 
// v:=<-ch // receive from ch and assign value to v





func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

