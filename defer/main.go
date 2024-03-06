package main

import "fmt"

func main() {
  
defer fmt.Println("Hello")

  for v:=0;v<10;v++ {
  
    defer fmt.Println(v)

}

fmt.Println("world")

}
