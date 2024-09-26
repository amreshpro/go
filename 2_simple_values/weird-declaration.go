package main;
// way1  - syntax :- var variablename type = value
// way2 - with the := sign , variablename := value


import "fmt"


func Declaration(){
	var name string = "Amresh"
	surname := "Maurya"

	fmt.Println(name,surname)


	// variable without initial value
	var a string //empty ""
	var b int //0
	var c bool //false
  
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	var a2, b2, c2, d2 int = 1, 3, 5, 7
	fmt.Println(a2)//1
	fmt.Println(b2)//3
	fmt.Println(c2)//5
	fmt.Println(d2)//7

	// Note: If you use the type keyword, it is only possible to declare one type of variable per line.
	var a3, b3 = 6, "Hello"
	c3, d3 := 7, "World!"
	fmt.Println("-----")
  
	fmt.Println(a3)//6
	fmt.Println(b3)// "Hello"
	fmt.Println(c3)//7
	fmt.Println(d3)//World

}