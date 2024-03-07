package main


//type *T is a pointer to a T value .Its zero value is nil
// & operator generates a pointer to its operand


import ("fmt")

func main() {
 var p *int
  num:=45
  p = &num 

  fmt.Println(p) //&num
  fmt.Println(*p)//45
  fmt.Println(&p)// address of p  
  fmt.Println(&num)//p
  fmt.Println(num) //45
}
