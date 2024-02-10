
### Go Notes

[x]  Staically typed
[x] Fast run time 
[x] Compiled
[x] Fast Compiled time 
[x] Supports concurrency through goroutines and channel.
[x] Has automatic garbage collection
[x] Does not support classes and objects
[x] Does not supports inheritance


#### first.go 

```go 
package main
import "fmt"
func main()  {
  fmt.Println("Hello Golang")
}
```


fmt.Println("Hello World!") is a statement.

In Go, statements are separated by ending a line (hitting the Enter key) or by a semicolon ";".

Hitting the Enter key adds ";" to the end of the line implicitly (does not show up in the source code).

> The left curly bracket { cannot come at the start of a line.

```go
package main
import ("fmt")

func main()//error
{
  fmt.Println("Hello World!")
}
```

 way1  - syntax :- var variablename type = value

 way2 - with the := sign , variablename := value

> Note: In this case, the type of the variable is inferred from the value (means that the compiler decides the type of the variable, based on the value).

> Note: It is not possible to declare a variable using :=, without assigning a value to it.


### var    vs    :=
var --> used inside or outside of function

:= --> only inside functions

var --> variable declaration and value assignment can done separately.

:= --> Variable declaration and value assignment **cannot be done separately** (must be done in the same line)


> Note: If you use the type keyword, it is only possible to declare one type of variable per line.
> 
> If the type keyword is not specified, you can declare different types of variables in the same line:
```go
 var a3, b3 = 6, "Hello"
	c3, d3 := 7, "World!"
	fmt.Println("-----")
  
	fmt.Println(a3)//6
	fmt.Println(b3)// "Hello"
	fmt.Println(c3)//7
	fmt.Println(d3)//World
```

#### Multiple variable declarations can also be grouped together into a block for greater readability:



```go
func main() {
   var (
     a int
     b int = 1
     c string = "hello"
   )

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}
```

Multiple constants can be grouped together into a block for readability:
```go
const (
  A int = 1
  B = 3.14
  C = "Hi!"
)
```