//Notes
// var a [10]int
// a[0:10]==a[:10]==a[0:]==a[:]

package main

import "fmt"

func slice0() {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 23}
	slice := primes[1:3] //[indexStart(included),endIndex(excluded)]
	fmt.Println(slice)
	names := []string{
		"Amresh", "Ojas", "Arun",
	}

	fmt.Println(names)
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a)
	fmt.Println(b)

}
