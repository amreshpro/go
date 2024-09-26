package main

import (
	"fmt"
	"time"
)

func main() {

	// days 

	switch time.Now().Weekday(){
	case time.Saturday , time.Sunday :
		fmt.Println("Yoyo Its Weekend!")
	case time.Monday:
		fmt.Print("1st Working Day")
	case time.Friday:
		fmt.Println("Last Working day")
	default:
		fmt.Println("Working Day")		

	}


	whoAmI := func ( i interface {})  {
		switch t := i.(type){
		case string:
		 fmt.Println("I am string ",t)
		case int:
			fmt.Println("I am int ",t)
		case float64:
			fmt.Println("I am float ",t)
		case bool:
			fmt.Println("I am bool ",t)
	    default:
			fmt.Println("I am other type",t)
		}

	}

	
	
	whoAmI("golang")
	whoAmI(24)
	whoAmI(true)
	whoAmI(34.43)
}