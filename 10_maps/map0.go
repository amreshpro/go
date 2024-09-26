package main

import "fmt"

// map[KeyType]ValueType
// var m map[string] int
/*  Map types are reference types, like pointers or slices,
and so the value of m above is nil; it doesn’t point to an initialized map.
A nil map behaves like an empty map when reading,
but attempts to write to a nil map will cause a runtime panic;
don’t do that.
To initialize a map, use the built in make function: */

// m := make(map[string]int)
/*The make function allocates and initializes a hash map data structure and returns a map value that points to it. The specifics of that data structure are an implementation detail of the runtime and are not specified by the language itself. In this article we will focus on the use of maps, not their implementation.*/
func Maps() {
	m := make(map[string]string)
	m["name"] = "Amresh"
	m["surname"] = "Maurya"
	fmt.Println(m)
	fmt.Println(m["isNotExist"])// 0,empty
	fmt.Println(len(m))//2
	delete(m,"surname")
	fmt.Println(m)
}
