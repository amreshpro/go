package main

//The const keyword declares the variable as "constant", which means that it is unchangeable and read-only.
/*
There are two types of constants:

Typed constants
const A int = 1
Untyped constants
const A = 1

Multiple constants can be grouped together into a block for readability:
const (
  A int = 1
  B = 3.14
  C = "Hi!"
)
*/

import "fmt"
 func ConstPro(){
	const PI float32 = 3.14;
	const FavNumber int32 = 999;
fmt.Println(PI)
fmt.Println(FavNumber)
 }