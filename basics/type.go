package main

import "fmt"


func main() {
var isLoggedIn bool  = false
var small int8 = 23
var small2 int16 = 278
var med int32 = 23232
var large int64 = 343424242
var defaultValue int 
fmt.Println(defaultValue) // 0 defaultValue is zero in int 

fmt.Println(small)
fmt.Println(small2)
fmt.Println(med)
fmt.Println(large)

fmt.Println(isLoggedIn)
fmt.Printf("Variable type : %T \n",isLoggedIn)
}
