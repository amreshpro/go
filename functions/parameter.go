package main 

import (
  "fmt"

)


// Note  int x, int y  ==> x,y int   shorthand

func add(x,y int) int  {

  return x+y

}

func main(){

fmt.Println(add(5,45))
fmt.Println(add(22,34))


}
