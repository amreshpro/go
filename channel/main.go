package main

import("fmt")
//<- the data flows in the direction of the arrow.
//Like maps and slices, channels must be created before use.
//ch:=make(chan int)
// ch<-v // send v to channel ch 
// v:=<-ch // receive from ch and assign value to v

func main() {

  ch:=make(chan int)
  ch<-6
  v:=<-ch
  fmt.Println(v)

}
