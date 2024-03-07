package main

//interface are named collection of method signature

import( 
  "fmt"
)



const Pi float64 = 3.14

type Geometry interface {
  Area() float64
  Perimeter( ) float64
}


type Rect struct{
  length,breadth float64
}

type Circle struct {
  radius float64
}

func (r Rect) Area() float64{
  return r.length * r.breadth
}

func (r Rect) Perimeter() float64 {
  
return 2 * (r.length+r.breadth) 

}


func (c Circle) Area() float64{

  return Pi * c.radius *c.radius
}


func (c Circle) Perimeter() float64{

return 2*Pi*c.radius

}

func measure(g Geometry)  {
  
fmt.Println(g)
fmt.Println(g.Area())
fmt.Println(g.Perimeter())  

}


func main() {

  r:=Rect{
    length:3,
    breadth:4,
  }

  c:=Circle{
    radius:5,
  }

 measure(r)
 measure(c) 
}
