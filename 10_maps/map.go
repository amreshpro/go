package main

import "fmt"
//associative datastructure - map-->hash(c+),object(javascript),dict(python)
func main() {
// creating map[key type] Value type	
  m:=make(map[string]string) 

  //setting an element 
  m["name"] = "Amresh"
  m["age"] = "24"

  fmt.Println(m["name"],m["age"])


	
}
